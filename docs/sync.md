# File Sync

`PixelFS` supports file synchronization between multiple devices. You can start the file synchronization service with a simple command to achieve synchronization across devices.

## Add Sync

Command: `pixelfs sync add`

```shell
pixelfs sync add \
    --name=example \
    --src=0x29e3abdb587207dc4ac9c708670eefde717ef307:/home/sync \
    --dest=0x1211abdb5839908jhuh9c708670eekljhmh08ndb:/home/sync \
    --duplex
```

Parameter description:

- `--name`: The name of the synchronization task.
- `--src`: Source storage location, in the format `node-id:/path`.
- `--dest`: Destination storage location, in the format `node-id:/path`.
- `--duplex`: Whether to enable bidirectional synchronization. If this parameter is not specified, the default is one-way synchronization.

## Ignore Files

During synchronization, you can ignore specific files or directories using the `.pixelfsignore` file.

> The format of the `.pixelfsignore` file is the same as `.gitignore`.

Example:

```text
# Ignore all .log files
*.log  

# Ignore directories
logs/ 
```

## Start Sync

Command: `pixelfs sync start`

```shell
pixelfs sync start <sync-id>
```

::: tip
You can use `pixelfs sync ls` to view synchronization task IDs.
:::

## Delete Sync

Command: `pixelfs sync rm`

```shell
pixelfs sync rm <sync-id>
```

## Stop Sync

Command: `pixelfs sync stop`

> 	Stopping a synchronization task will only stop the current synchronization service. The synchronization task will automatically resume the next time `PixelFS` starts.

```shell
pixelfs sync stop <sync-id>
```
