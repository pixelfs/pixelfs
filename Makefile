PIXELFSD_PID = $(shell ps aux | grep "pixelfs daemon" | grep -v grep | awk "{print $2}")
PIXELFS_WEBDAV_PID = $(shell ps aux | grep "pixelfs webdav" | grep -v grep | awk "{print $2}")

rwildcard=$(foreach d,$(wildcard $1*),$(call rwildcard,$d/,$2) $(filter $(subst *,%,$2),$d))

GO_SOURCES = $(call rwildcard,,*.go)
PROTO_SOURCES = $(call rwildcard,,*.proto)

build:
	nix build .#pixelfs

watch-daemon: kill-daemon
	@air --build.cmd "go build -o ./tmp/pixelfs -v ./cmd/pixelfs" --build.delay 1s --build.bin "tmp/pixelfs daemon" --build.exclude_dir "control,webdav"

kill-daemon:
	if [ -n "$(PIXELFSD_PID)" ]; then kill -9 $(PIXELFSD_PID); fi

watch-webdav: kill-webdav
	@air --build.cmd "go build -o ./tmp/pixelfs -v ./cmd/pixelfs" --build.delay 1s --build.bin "tmp/pixelfs webdav" --build.exclude_dir "control,pixelfsd"

kill-webdav:
	if [ -n "$(PIXELFS_WEBDAV_PID)" ]; then kill -9 $(PIXELFS_WEBDAV_PID); fi

lint:
	golangci-lint run --fix --timeout 10m

fmt:
	prettier --write '**/**.{ts,js,md,yaml,yml,sass,css,scss,html}'
	golines --max-len=88 --base-formatter=gofumpt -w $(GO_SOURCES)
	clang-format -style="{BasedOnStyle: Google, IndentWidth: 4, AlignConsecutiveDeclarations: true, AlignConsecutiveAssignments: true, ColumnLimit: 0}" -i $(PROTO_SOURCES)

proto-lint:
	cd proto/ && buf lint

generate:
	rm -rf gen
	buf generate proto
