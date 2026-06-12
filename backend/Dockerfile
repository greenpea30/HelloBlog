# ===== Stage 1: Build Go binary =====
FROM golang:1.25-alpine AS builder

WORKDIR /app

# 安装构建依赖
RUN apk add --no-cache gcc musl-dev

# 先复制 go.mod/go.sum 缓存依赖下载
COPY go.mod go.sum ./
RUN go mod download

# 复制源码并编译
COPY . .
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o /app/helloblog-api ./cmd/api/main.go && \
    CGO_ENABLED=0 go build -ldflags="-s -w" -o /app/helloblog-mcp ./cmd/mcp/main.go

# ===== Stage 2: 最小运行镜像 =====
FROM alpine:3.19

RUN apk add --no-cache ca-certificates tzdata && \
    cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone

WORKDIR /app

# 从 builder 复制编译好的二进制
COPY --from=builder /app/helloblog-api .
COPY --from=builder /app/helloblog-mcp .

# 创建上传目录
RUN mkdir -p /app/uploads

EXPOSE 8080 8081

# 默认启动 API（MCP 可通过 docker-compose 另起一个容器）
CMD ["./helloblog-api"]
