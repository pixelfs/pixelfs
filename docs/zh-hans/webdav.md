# WebDAV

`PixelFS` 提供了对 `WebDAV` 协议的支持，可以通过简单的命令启动 `WebDAV` 服务，使用通用 `WebDAV `客户端（例如 Windows 资源管理器、macOS Finder、Linux 文件管理器等）访问和管理 `PixelFS` 中的文件。

## 启动 WebDAV 服务

运行以下命令即可启动 `WebDAV` 服务：

```shell
pixelfs webdav
```

::: tip 注意
`WebDAV` 服务默认监听在 `0.0.0.0:5233`，可通过网络访问。如果需要更改端口或绑定地址，请在配置文件中修改相关设置。
:::

## 配置

`WebDAV` 所有配置信息都存放在 `$HOME/.pixelfs/config.toml` 文件中, 更详细的配置可以查看 [配置](/zh-hans/configuration)。

```toml
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

### 用户管理

```toml
[[webdav.users]]
username = 'admin'
password = '123456'
permissions = 'none'

[[webdav.users]]
username = "guest"
password = "123456"
permissions = "R"
```

::: tip 权限
permissions: 用户权限，由以下字符组成：

- C: 创建权限。
- R: 读取权限。
- U: 更新权限。
- D: 删除权限。

例如，permissions = "CR" 表示该用户只能创建和读取文件。
:::
