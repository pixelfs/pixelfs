# Configuration

`PixelFS` allows you to modify certain configurations, and all configuration details are stored in the `$HOME/.pixelfs/config.toml` file.

::: tip
If you are running `PixelFS` for the first time, a `config.toml` file will be automatically created.
:::

## Configuration Parameters

| Parameter Name             | Type     |         Default Value         | Description                                         |
|----------------------------|----------|:-----------------------------:|-----------------------------------------------------|
| endpoint                   | string   |   `https://www.pixelfs.io`    | Server address                                      |
| token                      | string   |               -               | User TOKEN information                              |
| pwd                        | string   |               -               | Path information used with `pixelfs cd`             |
| debug                      | bool     |            `false`            | Enable debug logs                                   |
| daemon.listen              | string   |        `0.0.0.0:15233`        | Address and port for the `daemon` service           |
| ffmpeg.cache.expire        | int      |            `86400`            | `ffmpeg` cache expiration time (seconds)            |
| ffmpeg.cache.path          | string   | `$HOME/.pixelfs/cache/ffmpeg` | `ffmpeg` cache path                                 |
| webdav.listen              | string   |        `0.0.0.0:5233`         | Address and port for the `webdav` service           |
| webdav.cache.expire        | int      |            `86400`            | `webdav` cache expiration time (seconds)            |
| webdav.cache.path          | string   | `$HOME/.pixelfs/cache/webdav` | `webdav` cache path                                 |
| webdav.cors.credentials    | bool     |               -               | Support for CORS credentials                        |
| webdav.cors.allow_origin   | string[] |            `['*']`            | Allowed origins for CORS requests                   |
| webdav.cors.allow_headers  | string[] |            `['*']`            | Allowed headers for CORS requests                   |
| webdav.cors.allow_methods  | string[] |            `['*']`            | Allowed methods for CORS requests                   |
| webdav.cors.expose_headers | string[] |               -               | Response headers accessible by clients              |
| webdav.cors.max_age        | int      |               -               | Maximum cache time for preflight requests (seconds) |
| webdav.users.username      | string   |               -               | Username for the `webdav` service                   |
| webdav.users.password      | string   |               -               | Password for the `webdav` service                   |
| webdav.users.permissions   | string   |               -               | `webdav` user permissions, e.g., `CRUD`             |

## Configuration Example

```toml
endpoint = 'https://www.pixelfs.io'
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
