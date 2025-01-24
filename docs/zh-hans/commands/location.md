# Location

在 `PixelFS` 中，`location` 是一个核心概念。只有将本地文件系统中的某个目录添加为 `PixelFS` 的 `location`，才能对该目录进行管理和操作。

## add

添加存储位置。

```shell
pixelfs location add
```

**选项**

- `--block-duration`: 指定分隔每个块的时间, 在生成 `m3u8` 文件时使用。默认值为 `20s`。
- `--block-size`: 指定每个块的大小, 在大部分场景下都会使用。默认值为 `4MB`。
- `--name`: 指定存储位置的名称。在同一个节点下，`name` 不能重复。
- `--node-id`: 指定节点的 `id`。也就是存储位置所在的节点。
- `--path`: 指定存储位置的路径。需要是本地文件系统中存在的目录。

## ls

列出所有存储位置。

```shell
pixelfs location ls
```

## rm

删除存储位置。

```shell
pixelfs location rm <location-id>
```

- `location-id`: 指定要删除的存储位置的 `id`。
