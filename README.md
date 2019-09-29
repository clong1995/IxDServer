# IxD相关服务
## 用于windows、linux、mac
## 编译
```bash
#在目标机器上编译
go build -ldflags "-w -s" -o bin/main
```
```bash
#交叉编译 unix
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags "-w -s" -o bin/main_darwin
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "-w -s" -o bin/main_win.exe
```
```bash
#交叉编译 win
SET CGO_ENABLED=0
SET GOOS=darwin
SET GOARCH=amd64
go build -ldflags "-w -s" -o bin/main

SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build -ldflags "-w -s" -o bin/main
```
```
#使用代理
#先开启本机代理
set all_proxy=http://127.0.0.1:9999
#然后开启go包的代理
set GOPROXY=https://goproxy.io
```
```
#升级
#杀死程序
ps -ef|gep xxx
#更换程序
#启动程序
nohup ./xxx &
```
## 运行，所有参数均有默认值，非必要无需覆盖
```
./main 
```
## 查看参数
```text
./main -h
```
## 升级
### 接口发布地址
```bash
scp -r assets clong@192.168.20.134:/tmp/
```
### 程序
```bash
scp main_linux clong@192.168.20.134:/tmp/
```
### 获取当前时间
```bash
date +%Y%m%d%H%M%S
```