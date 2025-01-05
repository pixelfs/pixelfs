package webdav

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/pixelfs/pixelfs/config"
	"github.com/pixelfs/pixelfs/log"
	cors "github.com/rs/cors/wrapper/gin"
	"golang.org/x/net/webdav"
)

const (
	httpRequestKey string = "http-request"
)

var (
	missingMethods = []string{
		"PROPFIND", "PROPPATCH", "MKCOL", "COPY", "MOVE", "LOCK", "UNLOCK",
	}
)

type PixelFS struct {
	cfg   *config.Config
	fs    *fileSystem
	users map[string]*config.User
}

func NewPixelFS() (*PixelFS, error) {
	cfg, err := config.GetConfig()
	if err != nil {
		return nil, err
	}

	if cfg.Token == "" {
		return nil, fmt.Errorf("token is empty")
	}

	users := make(map[string]*config.User)
	for _, user := range cfg.Webdav.Users {
		users[user.Username] = &user
	}

	return &PixelFS{
		cfg:   cfg,
		fs:    newFileSystem(cfg),
		users: users,
	}, nil
}

func (p *PixelFS) handler(c *gin.Context) {
	handler := webdav.Handler{
		Prefix:     p.cfg.Webdav.Prefix,
		FileSystem: p.fs,
		LockSystem: webdav.NewMemLS(),
	}

	handler.ServeHTTP(
		c.Writer,
		c.Request.WithContext(
			context.WithValue(c.Request.Context(), httpRequestKey, c.Request),
		),
	)
}

func (p *PixelFS) auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if len(p.cfg.Webdav.Users) > 0 {
			ctx.Writer.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)

			username, pwd, ok := ctx.Request.BasicAuth()
			if !ok {
				ctx.AbortWithStatus(http.StatusUnauthorized)
				return
			}

			user, ok := p.users[username]
			if !ok || user.Password != pwd {
				ctx.AbortWithStatus(http.StatusUnauthorized)
				return
			}

			allowed := user.Allowed(ctx.Request, func(filename string) bool {
				_, err := p.fs.Stat(ctx.Request.Context(), filename)
				return !os.IsNotExist(err)
			})

			if !allowed {
				ctx.AbortWithStatus(http.StatusForbidden)
				return
			}
		}
	}
}

func (p *PixelFS) cors() gin.HandlerFunc {
	return cors.New(cors.Options{
		AllowedOrigins:     p.cfg.Webdav.CORS.AllowOrigin,
		AllowedMethods:     p.cfg.Webdav.CORS.AllowMethods,
		AllowedHeaders:     p.cfg.Webdav.CORS.AllowHeaders,
		ExposedHeaders:     p.cfg.Webdav.CORS.ExposeHeaders,
		AllowCredentials:   p.cfg.Webdav.CORS.Credentials,
		MaxAge:             p.cfg.Webdav.CORS.MaxAge,
		OptionsPassthrough: false,
	})
}

func (p *PixelFS) Serve() error {
	gin.SetMode(gin.ReleaseMode)

	engine := gin.Default()
	group := engine.Group(p.cfg.Webdav.Prefix)

	group.Use(p.auth())
	group.Use(p.cors())
	group.Use(gin.Recovery())

	group.Any("/*webdav", p.handler)
	for _, v := range missingMethods {
		group.Handle(v, "/*webdav", p.handler)
	}

	log.Info().Str("listen", p.cfg.Webdav.Listen).Msg("pixelfs webdav is running")

	if len(p.cfg.Webdav.Users) == 0 {
		log.Warn().Msg("no users configured for webdav, authentication will be skipped")
	}

	// clean webdav cache
	cleanWebdavCache(p.cfg)

	return engine.Run(p.cfg.Webdav.Listen)
}
