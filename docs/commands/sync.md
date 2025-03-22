# Sync

`sync` command is used to manage file sync.

## add

Add a sync task.

```shell
pixelfs sync add \
    --name=example \
    --src=0x29e3abdb587207dc4ac9c708670eefde717ef307:/home/sync \
    --dest=0x1211abdb5839908jhuh9c708670eekljhmh08ndb:/home/sync
```

**Options**

- `--name`: The name of the syn task.
- `--src`: Source storage location, in the format `node-id:/path`.
- `--dest`: Destination storage location, in the format `node-id:/path`.
- `--duplex`: Whether to enable bidirectional sync. If this parameter is not specified, the default is one-way sync.
- `--interval`: Sync interval time in seconds, default is 3600 seconds.
- `--enabled`: Whether to enable the sync task, default is true.
- `--id`: Specify the id of the sync task, used when editing a sync task.

## ls

List all sync tasks.

```shell
pixelfs sync ls
```

## rm

Delete a sync task.

```shell
pixelfs sync rm <sync-id>
```

- `sync-id`：The `id` of the sync task.

## start

Start the sync task.

```shell
pixelfs sync start <sync-id>
```

- `sync-id`：The `id` of the sync task.

## stop

Stop the sync task.

```shell
pixelfs sync stop <sync-id>
```

- `sync-id`：The `id` of the sync task.
