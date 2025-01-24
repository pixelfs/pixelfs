# Daemon

`daemon` 命令用于启动一个后台进程，该进程会在后台持续运行，监听并处理来自服务端的请求。

```shell
pixelfs daemon
```

::: tip 注意
`daemon` 服务默认监听在 `0.0.0.0:15233`。如果需要更改端口或绑定地址，请在配置文件中修改相关设置。
:::

::: warning
如果未使用 `pixelfs auth login` 登录，`daemon` 服务将无法正常工作。
:::
