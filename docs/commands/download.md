# Download

`download` command is used to download files.

```shell
pixelfs download <file/dir>
```

- `file/dir`: The file or directory to be downloaded.

**Options**

- `-o, --output`: The output path for the downloaded file.
- `-t, --thread`: The number of download threads.

## Example

```shell
pixelfs download 0x29e3abdb587207dc4ac9c708670eefde717ef307:/home/file.txt
```
