# 快速入门

你可以从 [Releases](https://github.com/pixelfs/pixelfs/releases) 下载 `PixelFS` 的二进制文件，选择适合你的操作系统和架构的版本。

## 安装 PixelFS

::: details Ubuntu/Dedian
```shell
# 下载 PixelFS 二进制文件 (请根据需要选择版本)
wget https://github.com/pixelfs/pixelfs/releases/download/v1.0.0/pixelfs_1.0.0_linux_amd64.deb

# 安装 PixelFS
sudo dpkg -i pixelfs_1.0.0_linux_amd64.deb

# 启动 PixelFS 服务 (以用户模式运行)
systemctl --user start pixelfs
systemctl --user enable pixelfs

# 启用长时间会话，使得即使用户退出登录后，用户的服务仍然保持运行
loginctl enable-linger
```
:::

::: details CentOS/RHEL
```shell
# 下载 PixelFS 二进制文件 (请根据需要选择版本)
wget https://github.com/pixelfs/pixelfs/releases/download/v1.0.0/pixelfs_1.0.0_linux_amd64.rpm

# 安装 PixelFS
sudo rpm -i pixelfs_1.0.0_linux_amd64.rpm

# 启动 PixelFS 服务 (以用户模式运行)
systemctl --user start pixelfs
systemctl --user enable pixelfs

# 启用长时间会话，使得即使用户退出登录后，用户的服务仍然保持运行
loginctl enable-linger
```
:::

::: details MacOS
```shell
# 使用 Homebrew 安装 PixelFS
brew tap pixelfs/tap
brew install pixelfs

# 启动 PixelFS 服务
brew services start pixelfs
```
:::

### 手动安装

1. 前往 [Releases](https://github.com/pixelfs/pixelfs/releases)，下载适合你操作系统和架构的 `PixelFS` 二进制文件。
2. 解压下载的文件，并将解压后的二进制文件移动到系统路径（如 `/usr/local/bin`）或将所在目录添加到系统的 `PATH` 环境变量中。

```shell
mv pixelfs /usr/local/bin/
chmod +x /usr/local/bin/pixelfs
```

3. 启动 `PixelFS` 后台服务（可选）：

```shell
pixelfs daemon
```

::: warning
如果未启动 `daemon` 服务，`PixelFS` 将不会对该设备上的文件进行管理。
:::

## 登录

命令：`pixelfs auth login`

运行命令后，终端会输出一个登录地址，例如：

```text
To authenticate, Please visit:

    https://www.pixelfs.io/auth/cli/d9e5ccb055924bc4d0801a56524766d52f0c26397e9f431abb19ada6be9c16df

Waiting for session...
```

登录步骤：

1. 复制生成的登录地址。
2. 在浏览器中打开该地址，并按照页面提示完成登录。
3. 登录成功后，终端会自动完成认证并准备就绪。

## 添加存储位置

命令：`pixelfs location add`

```shell
pixelfs location add \
    --node-id=0x29e3abdb587207dc4ac9c708670eefde717ef307 \
    --path=/path/to/data \
    --name=location-data
```

参数说明：

- `--node-id`: 节点 ID，标识存储节点。
- `--path`: 存储数据的路径，`PixelFS` 将在该路径中管理文件。
- `--name`: 存储位置的名称，需在同一节点内唯一，便于识别和管理。

::: tip
你可以使用 `pixelfs id` 查看节点 ID。
:::

## 添加 S3 存储信息

命令：`pixelfs storage add`

```shell
pixelfs storage add \
	--type=s3 \
	--region='auto' \
	--endpoint='https://xxxxxx.r2.cloudflarestorage.com' \
	--bucket='pixelfs' \
	--access-key='accessKeyId' \
	--secret-key='secretAccessKey'
```

参数说明：

- `--type`: 存储类型，目前仅支持 `s3`。
- `--region`: 存储区域，`auto` 表示自动选择。
- `--endpoint`: 存储节点的访问地址。
- `--bucket`: 存储桶名称。
- `--access-key`: 访问密钥 ID。
- `--secret-key`: 访问密钥密钥。

## 关联 S3 存储到节点

命令：`pixelfs storage link add`

```shell
pixelfs storage link add \
    --node-id=0x29e3abdb587207dc4ac9c708670eefde717ef307 \
    --storage-id=fe0dc5d1-da9f-41e0-a243-3b2582fc3501
```

::: tip
你可以使用 `pixelfs storage ls` 查看已添加的存储 ID。
:::

## 常用命令

命令：`pixelfs ls`

命令：`pixelfs cd 0x29e3abdb587207dc4ac9c708670eefde717ef307`

```shell
$ pixelfs ls
drw-------    - pixelfs 14 Jan 03:01 0x29e3abdb587207dc4ac9c708670eefde717ef307 ONLINE PIXELFS-NODE
```

## 完成指南！

恭喜你完成了 `PixelFS` 的快速上手 🎉 你现在可以管理你的文件了。

`PixelFS` 还有更多命令需要熟悉，你可以通过运行 `pixelfs --help` 或者 `pixelfs` 来查看它们。

