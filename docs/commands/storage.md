# Storage

`storage` command is used to manage storage (add/remove storage), primarily serving the file transfer and copy functionality between devices. Each file must be associated with a specified `storage` before transfer.

## add

Add a storage.

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

**Options**

- `--name`: Storage name.
- `--type`: Storage type, currently only `s3` is supported.
- `--region`: `S3` storage region, auto means auto-select.
- `--endpoint`: `S3` storage endpoint.
- `--bucket`: `S3` bucket name.
- `--access-key`: `S3` access key ID.
- `--secret-key`: `S3` secret key.
- `--prefix`: Storage path prefix.
- `--path-style`: Whether to enable path-style access.
- `--network`: `S3` service network type, `public` or `private`.

## ls

List all storages.

```shell
pixelfs storage ls
```

## rm

Delete a storage.

```shell
pixelfs storage rm <storage-id>
```
