# Location

In `PixelFS`, `location` is a core concept. Only by adding a directory from the local file system as a location in `PixelFS` can that directory be managed and operated.

## add

Add a storage location.

```shell
pixelfs location add
```

**Options**

- `--block-duration`: Time to separate each block, used when generating `m3u8` files. The default value is `20s`.
- `--block-size`: Size of each block, which is used in most scenarios. The default value is `4MB`.
- `--name`: Name of the storage location. The `name` must be unique within the same node.
- `--node-id`: Id of the node. This is the node where the storage location resides.
- `--path`: Path of the storage location. It must be an existing directory in the local file system.

## ls

List all storage locations.

```shell
pixelfs location ls
```

## rm

Remove a storage location.

```shell
pixelfs location rm <location-id>
```

- `location-id`: Id of the storage location to be removed.
