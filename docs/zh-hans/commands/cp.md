# Cp

`cp` 命令用于复制文件或目录。

```shell
pixelfs cp <source> <destination>
```

- `source`: 指定源文件或目录。
- `destination`: 指定目标文件或目录。

## 示例

```shell
pixelfs cp /home/file.txt 0x29e3abdb587207dc4ac9c708670eefde717ef307:/home/file.txt
```

此命令将 `/home/file.txt` 复制到设备 `0x29e3abdb587207dc4ac9c708670eefde717ef307` 的 `/home/file.txt`。
