# Download

`download` 命令用于下载文件。

```shell
pixelfs download <file/dir>
```

- `file/dir`: 指定要下载的文件或目录。

**选项**

- `-o, --output`: 指定下载文件的输出路径。
- `-t, --thread`: 指定下载线程数。

## 示例

```shell
pixelfs download 0x29e3abdb587207dc4ac9c708670eefde717ef307:/home/file.txt
```
