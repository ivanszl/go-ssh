# ssh client for Golang
本客户端用于解决使用macOS系统的，使用iTerm的重度用户使用
本人就是其一，由于没有很好的记住密码和实现自动登陆功能，而写的一个工具

>注： 当然可以采用脚本实现自动登陆，但是会导致`rz`和`sz`用不了


支持:
- 采用证书登陆
- 支持证书密码
- 账号密码自登陆，采用用`username`@`host`:`port` 方式
- 证书密码和用户密码改为`SSH_PASSWORD`环境变量模式

## 安装方式

#### 采用go安装模式
```shell
go install github.com:ivanszl/go-ssh.git
```

#### 下载编译
```shell
git clone git@github.com:ivanszl/go-ssh.git
cd go-ssh
make build
cp go-ssh path/go-ssh
```

## 使用方式

#### 采用账号密码登陆
```shell
go-ssh -h test:test@127.0.0.1
```
#### 采用证书登陆
```shell
go-ssh -h test@127.0.0.1 -k ~/.ssh_client.pem -p key_passwd
````
