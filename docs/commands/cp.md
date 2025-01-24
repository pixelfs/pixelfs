# Cp

`cp` command is used to copy files or directories.

```shell
pixelfs cp <source> <destination>
```

- `source`: source file or directory.
- `destination`: target file or directory.

## Example

```shell
pixelfs cp /home/file.txt 0x29e3abdb587207dc4ac9c708670eefde717ef307:/home/file.txt
```

This command copies `/home/file.txt` to `/home/file.txt` on the device `0x29e3abdb587207dc4ac9c708670eefde717ef307`.
