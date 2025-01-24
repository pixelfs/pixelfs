# Upload

`upload` command is used to upload a file.

```shell
pixelfs upload <file> --input=<input-file>
```

- `file`：The path where the uploaded file will be saved.

**Options**

- `-i, --input`：Specify the input file.

## Example

```shell
pixelfs upload /home/test.txt --input=/home/test.txt
```
