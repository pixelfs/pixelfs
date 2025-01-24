# WebDAV

`PixelFS` provides support for the `WebDAV` protocol, allowing you to start a `WebDAV` service with a simple command. You can use any standard `WebDAV` client (such as Windows Explorer, macOS Finder, Linux file managers, etc.) to access and manage files in `PixelFS`.

## Starting the WebDAV Service

Run the following command to start the `WebDAV` service:

```shell
pixelfs webdav
```

::: tip
The `WebDAV` service listens on `0.0.0.0:5233` by default and is accessible over the network. If you need to change the port or bind address, modify the relevant settings in the configuration file.
:::

## Configuration

All `WebDAV` configuration details are stored in the `$HOME/.pixelfs/config.toml` file. For more detailed configuration, refer to [Configuration](/configuration).

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

### User Configuration

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

::: tip Permissions
permissions: User permissions, composed of the following characters:

- C: Create permission.
- R: Read permission.
- U: Update permission.
- D: Delete permission.

For example, permissions = "CR" means the user can only create and read files.
:::
