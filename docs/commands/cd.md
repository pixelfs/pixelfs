# Cd

`cd` command is used to change the current working directory.

```shell
pixelfs cd [node:/location/path/to]
```

- `node`: Unique identifier of the target device node.
- `/location/path/to`: Target path to switch to.

::: tip
`pixelfs cd ::` will switch to the root directory.
:::

## Examples

1. Switch to a specific path on a device:

```shell
pixelfs cd 0x29e3abdb587207dc4ac9c708670eefde717ef307:/home
```

This command will switch the working directory to the `/home` path on device `0x29e3abdb587207dc4ac9c708670eefde717ef307`.

2. Switch to the root directory:

```shell
pixelfs cd
# or
pixelfs cd ::
```

3. Switch to a subdirectory:

If the current directory is `/home`, the following command switches to `/home/user`:

```shell
pixelfs cd user
```

## Notes

1. Path validation: If the specified path does not exist or access is denied, an error message will be displayed.
2. Relative path support: Relative paths are supported, for example:

```shell
pixelfs cd ../documents
```
