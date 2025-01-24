# Storage Link

`storage link` 命令用于关联存储到指定的 `node` 或者 `location`。

## add

关联存储到指定的 `node` 或者 `location`。

```shell
pixelfs storage link add \
    --storage-id=<storage-id> \
    --node-id=<node-id> \
    --limit-size=128MB
```

::: warning
`--node-id` 和 `--location-id` 二选一。`location` 优先级高于 `node`。
:::

::: tip
如果你需要切换 `node` 关联的 `storage`，需要先执行 `clean` 命令，然后再执行 `add` 命令。
:::

**选项**

- `--storage-id`：存储 ID。
- `--node-id`：节点 ID。
- `--location-id`：存储位置 ID。
- `--limit-size`：关联存储的限制大小。如果超过限制大小，将会自动清理 `block`。默认为 `128MB`。

## ls

列出所有关联存储。

```shell
pixelfs storage link ls
```

## rm

删除一个关联存储。

```shell
pixelfs storage link rm <link-id>
```

## clean

清除释放关联存储中的 `block`。

```shell
pixelfs storage link clean <link-id>
```
