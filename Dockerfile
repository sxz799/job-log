# 使用官方 Golang 镜像作为基础镜像
FROM golang:1.20 as builder

# 设置工作目录
WORKDIR /go/src/github.com/sxz799/job-log/server

RUN apk --no-cache add gcc musl-dev

# 将应用的代码复制到容器中
COPY ./server/ .


# 编译应用程序
RUN go env -w GO111MODULE=on \
    && go env \
    && go mod tidy \
    && go build -o app .


FROM node:16


WORKDIR /go/src/github.com/sxz799/job-log/web
COPY ./web/ .

RUN rm vite.config.ts
RUN mv vite.config-docker.ts vite.config.ts

RUN echo 'VITE_BASE_PATH=/' > .env.production


RUN yarn && yarn build


FROM alpine:latest

WORKDIR /home

RUN apk --no-cache add bash

RUN apk update && apk add tzdata 
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime 
RUN echo "Asia/Shanghai" > /etc/timezone


COPY --from=0 /go/src/github.com/sxz799/job-log/server/app ./
COPY --from=0 /go/src/github.com/sxz799/job-log/server/conf.yaml ./
COPY --from=1 /go/src/github.com/sxz799/job-log/web/dist/ ./dist

RUN mkdir "cert"
RUN mkdir "conf"

EXPOSE 3000

# 运行应用程序
CMD ["./app"]
