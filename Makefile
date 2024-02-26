# Makefile for building the SeeInLinux project

# 设置目标操作系统和架构
GOOS=linux
GOARCH=amd64

# 定义输出的可执行文件名
OUTPUT_NAME=SeeInLinux

# 默认目标
all: build

# 编译项目
build:
	@echo "开始编译项目..."
	@GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $(OUTPUT_NAME)
	@if [ -f "$(OUTPUT_NAME)" ]; then \
		echo "编译成功，输出文件：$(OUTPUT_NAME)"; \
	else \
		echo "编译失败."; \
	fi

# 清理生成的文件
clean:
	@rm -f $(OUTPUT_NAME)
	@echo "清理完成."

.PHONY: all build clean
