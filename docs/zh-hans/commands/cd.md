# Cd

`cd` 命令用于切换当前工作目录。

```shell
pixelfs cd [node:/location/path/to]
```

- `node`: 指定目标设备节点的唯一标识。
- `/location/path/to`: 指定切换到的目标路径。

::: tip
`pixelfs cd ::` 会切换到根目录。
:::

## 示例

1. 切换到指定设备的路径：

```shell
pixelfs cd 0x29e3abdb587207dc4ac9c708670eefde717ef307:/home
```

此命令将工作目录切换到设备 `0x29e3abdb587207dc4ac9c708670eefde717ef307` 的 `/home` 路径。

2. 切换到根目录：

```shell
pixelfs cd
# or
pixelfs cd ::
```

3. 切换到子目录：

假设当前目录为 `/home`，执行以下命令可切换到 `/home/user`：

```shell
pixelfs cd user
```

## 注意事项

1. 路径验证: 如果指定路径不存在或无权限访问，会提示错误信息。
2. 相对路径支持: 支持相对路径，例如：

```shell
pixelfs cd ../documents
```
