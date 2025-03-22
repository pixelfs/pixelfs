# Sync

`sync` 命令用于管理文件同步。

## add

添加同步任务。

```shell
pixelfs sync add \
    --name=example \
    --src=0x29e3abdb587207dc4ac9c708670eefde717ef307:/home/sync \
    --dest=0x1211abdb5839908jhuh9c708670eekljhmh08ndb:/home/sync
```

**选项**

- `--name`：同步任务的名称。
- `--src`：源存储位置，格式为 `node-id:/path`。
- `--dest`：目标存储位置，格式为 `node-id:/path`。
- `--duplex`：是否双向同步，如果不指定该参数，则默认为单向同步。
- `--interval`：同步间隔时间，单位为秒，默认为 3600 秒。
- `--enabled`：是否启用同步任务，默认为 `true`。
- `--id`：指定同步任务的 `id`，编辑同步任务时使用。

## ls

列出所有同步任务。

```shell
pixelfs sync ls
```

## rm

删除同步任务。

```shell
pixelfs sync rm <sync-id>
```

- `sync-id`：同步任务的 `id`。

## start

启动同步任务。

```shell
pixelfs sync start <sync-id>
```

- `sync-id`：同步任务的 `id`。

## stop

停止同步任务。

```shell
pixelfs sync stop <sync-id>
```

- `sync-id`：同步任务的 `id`。
