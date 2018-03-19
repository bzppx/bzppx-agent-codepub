```
                    __                     __                                     __
  _____ ____   ____/ /___   ____   __  __ / /_       ____    ____   ___   ____   / /_
 / ___// __ \ / __  // _ \ / __ \ / / / // __ \     / __ \  / __ \ / _ \ / __ \ / __/
/ /__ / /_/ // /_/ //  __// /_/ // /_/ // /_/ / -- / /_/ / / /_/ //  __// / / // /_
\___/ \____/ \__,_/ \___ / .___/ \__,_//_.___/     \__/\_\ \___,/ \___//_/ /_/ \__/
                        /_/                               /____/
```

[![stable](https://img.shields.io/badge/stable-stable-green.svg)](https://github.com/bzppx/bzppx-agent-codepub/) [![license](https://img.shields.io/github/license/bzppx/bzppx-agent-codepub.svg?style=plastic)]() [![download_count](https://img.shields.io/github/downloads/bzppx/bzppx-agent-codepub/total.svg?style=plastic)](https://github.com/bzppx/bzppx-agent-codepub/releases) [![download](https://img.shields.io/github/release/bzppx/bzppx-agent-codepub.svg?style=plastic)](https://github.com/bzppx/bzppx-agent-codepub/releases)

# 前言
codepub-agent 是 codepub 代码发布系统的 agent 部署程序，需配合 codepub 后台使用

地址：https://github.com/bzppx/bzppx-codepub

# 安装

## 1. 添加节点
请先在 codepub 后台添加节点

## 2. 下载二进制程序
打开 https://github.com/bzppx/bzppx-agent-codepub/releases 找到对应平台的版本下载编译好的压缩包

```
# 创建目录
$ mkdir codepub-agent
$ cd codepub
# 以 linux amd64 为例，下载版本 0.8 压缩包
$ wget https://github.com/bzppx/bzppx-agent-codepub/releases/download/v0.8/bzppx-agent-codepub-linux-amd64.tar.gz
# 解压到当前目录
$ tar -zxvf bzppx-agent-codepub-linux-amd64.tar.gz
```

## 3. 配置 agent

打开 config.toml 配置文件

```
[access]
token = "53d86cc809090ee047cf343fd787888e" # codepub 后台添加节点后生成的 token

[rpc]
listen = ":9091" # codepub 后台添加节点的 port

[cert]
key_file = "cert/server.key"
crt_file = "cert/server.pem"

[log]
level = ["info","error","debug"]
dir = "log"
FileMaxSize = 102400000
MaxCount = 3
```

## 4. 后台启动 agent
```
nohub ./bzppx-agent-codepub --conf config.toml &
```

# 开发

环境要求：go 1.8
```
$ git clone https://github.com/bzppx/bzppx-agent-codepub.git
$ cd bzppx-agent-codepub
$ go get ./...
$ go build ./
```

# 反馈

欢迎提交意见和代码 https://github.com/bzppx/bzppx-codepub/issues

## License

MIT

谢谢
---
Create By BZPPX