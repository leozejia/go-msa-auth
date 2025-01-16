# 使用官方 Go 镜像作为基础镜像
FROM golang:1.20

# 设置工作目录
WORKDIR /app

# 复制 go.mod 和 go.sum
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制项目文件
COPY . .

# 构建可执行文件
RUN go build -o go-msa-auth ./cmd/main.go

# 暴露端口
EXPOSE 8080

# 启动服务
CMD ["./go-msa-auth"]
