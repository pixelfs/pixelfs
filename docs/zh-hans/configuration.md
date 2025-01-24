# 配置

`PixelFS` 允许你修改一些配置，所有配置信息都存放在 `$HOME/.pixelfs/config.toml` 文件中。

::: tip 注意
如果你是首次运行 `PixelFS`, 则会自动创建一个 `config.toml` 文件。
:::

## 参数配置说明

| 参数名                      | 类型      |                 默认值               | 说明                               |
|----------------------------|----------|:-----------------------------------:|------------------------------------|
| endpoint                   | string   |        `https://pixelfs.io`         | 服务端地址                           |
| token                      | string   |                  -                  | 用户 TOKEN 信息                     |
| pwd                        | string   |                  -                  | 使用 `pixelfs cd` 切换的路径信息     |
| debug                      | bool     |               `false`               | 是否开启调试日志                     |
| daemon.listen              | string   |           `0.0.0.0:15233`           | `daemon` 服务监听的地址和端口         |
| ffmpeg.cache.expire        | int      |               `86400`               | `ffmpeg` 缓存时间（秒）              |
| ffmpeg.cache.path          | string   |    `$HOME/.pixelfs/cache/ffmpeg`    | `ffmpeg` 缓存路径                    |
| webdav.listen              | string   |           `0.0.0.0:5233`            | `webdav` 服务监听的地址和端口         |
| webdav.cache.expire        | int      |               `86400`               | `webdav` 缓存时间（秒）              |
| webdav.cache.path          | string   |    `$HOME/.pixelfs/cache/webdav`    | `webdav` 缓存路径                    |
| webdav.cors.credentials    | bool     |                  -                  | 是否支持跨域请求凭据（CORS）          |
| webdav.cors.allow_origin   | string[] |               `['*']`               | 允许的跨域请求来源                   |
| webdav.cors.allow_headers  | string[] |               `['*']`               | 允许的跨域请求头                     |
| webdav.cors.allow_methods  | string[] |               `['*']`               | 允许的跨域请求方法                   |
| webdav.cors.expose_headers | string[] |                  -                  | 可被客户端访问的响应头                |
| webdav.cors.max_age        | int      |                  -                  | 预检请求的最大缓存时间（秒）           |
| webdav.users.username      | string   |                  -                  | `webdav` 服务用户名                  |
| webdav.users.password      | string   |                  -                  | `webdav` 服务密码                    |
| webdav.users.permissions   | string   |                  -                  | `webdav` 用户权限，例如：`CRUD`        |

## 示例配置

```toml
endpoint = 'https://pixelfs.io'
pwd = '/'
token = 'auth-token'

[daemon]
listen = '0.0.0.0:15233'

[ffmpeg]
[ffmpeg.cache]
expire = 86400
path = '$HOME/.pixelfs/cache/ffmpeg'

[webdav]
listen = '0.0.0.0:5233'

[webdav.cache]
expire = 86400
path = '$HOME/.pixelfs/cache/webdav'

[webdav.cors]
allow_headers = ['*']
allow_methods = ['*']
allow_origin = ['*']

[[webdav.users]]
username = 'admin'
password = '123456'
permissions = 'none'

[[webdav.users.rules]]
path = '/0x29e3abdb587207dc4ac9c708670eefde717ef307'
permissions = 'CRUD'
```
