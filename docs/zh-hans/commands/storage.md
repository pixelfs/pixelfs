# Storage

`storage` 命令用于管理存储（添加/删除存储），主要服务于跨设备之间文件传输复制功能。每个文件在传输前，必须关联一个指定的 `storage`。

## add

添加一个存储。

```shell
pixelfs storage add \
    --name='cf-test' \
	--type=s3 \
	--region='auto' \
	--endpoint='https://xxxxxx.r2.cloudflarestorage.com' \
	--bucket='pixelfs' \
	--access-key='accessKeyId' \
	--secret-key='secretAccessKey'
```

**选项**

- `--name`：存储名称。
- `--type`：存储类型，目前仅支持支持 `s3`。
- `--region`：`S3` 存储区域，`auto` 表示自动选择。
- `--endpoint`：`S3` 存储服务端点。
- `--bucket`：`S3` 存储桶名称。
- `--access-key`：`S3` 存储访问密钥 ID。
- `--secret-key`：`S3` 存储访问密钥。
- `--prefix`：存储路径前缀。
- `--path-style`：是否启用路径样式访问。
- `--network`：`S3` 服务网络类型，`public` 或 `private`。

## ls

列出所有存储。

```shell
pixelfs storage ls
```

## rm

删除一个存储。

```shell
pixelfs storage rm <storage-id>
```
