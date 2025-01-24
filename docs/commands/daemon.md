# Daemon

`daemon` command is used to start a background process that will continuously run in the background, listening and handling requests from the server.

```shell
pixelfs daemon
```

::: tip 注意
`daemon` service listens on `0.0.0.0:15233` by default. If you need to change the port or bind address, modify the relevant settings in the configuration file.
:::

::: warning
If you have not logged in using `pixelfs auth login`, the daemon service will not function properly.
:::
