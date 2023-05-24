#!/bin/bash
build() {
    sudo CGO_ENABLED=0 GOOS=$2  GOARCH=$1 go build -o keysbuilder main.go
}

if [ ! $1 ];then
     echo "./$0 [架构类型 操作系统 | 示例: ./$0 amd64 linux]"
     echo "支持的架构类型有arm64/amd64 支持的运行操作系统有windows/linux/darwin(macos)"
     exit 1
     else
     build $1 $2
fi
