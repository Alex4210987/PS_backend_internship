# 使用官方的 Go 镜像作为基础镜像
FROM golang:latest

# 设置工作目录
WORKDIR /app

# 复制项目代码到镜像中
COPY . .

# 构建应用程序
RUN go build -o myapp

# 设置容器启动时运行的命令
CMD ["./myapp"]

