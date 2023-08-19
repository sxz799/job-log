# 使用官方 Golang 镜像作为基础镜像
FROM golang:1.20-alpine as builder

# 添加gcc环境 以支持sqlite
RUN apk --no-cache add gcc musl-dev

# 设置工作目录
WORKDIR /go/src/github.com/sxz799/todo-demo

# 将应用的代码复制到容器中
COPY ./server/ .


# 编译应用程序
RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go env \
    && go mod tidy \
    && go build -o app .

# 使用alpine镜像
FROM nginx:latest

COPY ./web/dist/ /user/share/nginx/html
WORKDIR /home

COPY --from=0 /go/src/github.com/sxz799/todo-demo/app ./


# 开放端口
EXPOSE 80

# 运行应用程序
CMD ["./app"]
