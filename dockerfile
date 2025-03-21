# 第一阶段：构建 Go 二进制
FROM golang:1.23 AS builder

WORKDIR /app
 # 当前上下文已是 backend 目录
COPY . . 
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# 第二阶段：运行环境
FROM alpine:latest

# 安装时区数据
RUN apk add --no-cache tzdata 

# 设置时区（根据实际需求调整，如 Asia/Shanghai、Europe/London 等）
ENV TZ=Asia/Shanghai

# 将时区文件链接到系统目录（Alpine 需要此操作）
RUN ln -sf /usr/share/zoneinfo/${TZ} /etc/localtime

WORKDIR /app
COPY --from=builder /app/main .
COPY ./settings.yaml .
COPY ./uploads/ ./uploads/
COPY ./models/esmodels/article_mapper.json ./models/esmodels/article_mapper.json
COPY ./models/esmodels/fulltext_mapper.json ./models/esmodels/fulltext_mapper.json

EXPOSE 8080
CMD ["./main"]

