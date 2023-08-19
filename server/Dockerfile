FROM golang:alpine as builder

WORKDIR /go/src/github.com/sxz799/todo-demo/server



COPY . .

RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go env -w CGO_ENABLED=0 \
    && go env \
    && go mod tidy \
    && go build -o server .

FROM alpine:latest


WORKDIR /go/src/github.com/sxz799/todo-demo/server

COPY --from=0 /go/src/github.com/sxz799/todo-demo/server/server ./
COPY --from=0 /go/src/github.com/sxz799/todo-demo/server/conf.yaml ./


EXPOSE 4000
ENTRYPOINT ./server
