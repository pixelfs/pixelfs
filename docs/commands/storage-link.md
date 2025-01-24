# Storage Link

`storage link` command is used to associate storage with a specified `node` or `location`.

## add

Associate storage with a specified node or location.

```shell
pixelfs storage link add \
    --storage-id=<storage-id> \
    --node-id=<node-id> \
    --limit-size=128MB
```

::: warning
You must use either `--node-id` or `--location-id`. `location` takes precedence over `node`.
:::

::: tip
If you need to change the `storage` associated with a `node`, you must first execute the `clean` command and then the `add` command.
:::

**Options**

- `--storage-id`：Storage ID.
- `--node-id`：Node ID.
- `--location-id`：Storage location ID.
- `--limit-size`：The size limit for the associated storage. If the limit is exceeded, the block will be automatically cleared. The default is `128MB`.

## ls

List all associated storages.

```shell
pixelfs storage link ls
```

## rm

Delete an associated storage.

```shell
pixelfs storage link rm <link-id>
```

## clean

Clear and release the `block` in the associated storage.

```shell
pixelfs storage link clean <link-id>
```
