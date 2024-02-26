#!/bin/bash

# 设置目标操作系统为 Linux
GOOS=linux

# 设置目标架构为 amd64
GOARCH=amd64

# 定义输出的可执行文件名
OUTPUT_NAME=SeeInLinux

# 编译项目
echo "开始编译项目..."
go build -o $OUTPUT_NAME

# 检查是否编译成功
if [ -f "$OUTPUT_NAME" ]; then
    echo "编译成功，输出文件：$OUTPUT_NAME"
else
    echo "编译失败."
fi
