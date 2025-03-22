# 文件同步

`PixelFS` 支持多设备之间的文件同步，可以通过简单的命令启动文件同步服务，实现多设备之间的文件同步。

## 添加同步

命令：`pixelfs sync add`

```shell
pixelfs sync add \
    --name=example \
    --src=0x29e3abdb587207dc4ac9c708670eefde717ef307:/home/sync \
    --dest=0x1211abdb5839908jhuh9c708670eekljhmh08ndb:/home/sync \
    --duplex
```

参数说明：

- `--name`：同步任务的名称。
- `--src`：源存储位置，格式为 `node-id:/path`。
- `--dest`：目标存储位置，格式为 `node-id:/path`。
- `--duplex`：是否双向同步，如果不指定该参数，则默认为单向同步。

## 忽略文件

在同步过程中，可以通过 `.pixelfsignore` 文件忽略指定文件或目录。

> `.pixelfsignore` 文件的格式与 `.gitignore` 文件相同。

示例：

```text
# 忽略所有 .log 文件
*.log

# 忽略目录
logs/
```

## 启动同步

命令：`pixelfs sync start`

```shell
pixelfs sync start <sync-id>
```

::: tip
你可以使用 `pixelfs sync ls` 查看同步任务 ID。
:::

## 删除同步

命令：`pixelfs sync rm`

```shell
pixelfs sync rm <sync-id>
```

## 停止同步

命令：`pixelfs sync stop`

> 停止同步任务仅会停止当前同步服务，下次 `PixelFS` 服务启动时，同步任务将自动恢复。

```shell
pixelfs sync stop <sync-id>
```
