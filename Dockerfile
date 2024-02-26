# 使用官方 Go 镜像作为构建环境
FROM golang:alpine as builder

# 设置工作目录
WORKDIR /app

# 复制源代码到容器中
COPY . .

# 设置目标操作系统和架构环境变量
ENV GOOS=linux GOARCH=amd64

# 编译项目
RUN go build -o SeeInLinux

# 使用 scratch 作为运行环境
FROM scratch

# 从构建阶段复制可执行文件到当前目录
COPY --from=builder /app/SeeInLinux .

# 运行可执行文件
CMD ["./SeeInLinux"]
