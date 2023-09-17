FROM golang:alpine as builder

WORKDIR /go/src/github.com/sxz799/job-log/server

COPY . .

RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go env \
    && go mod tidy \
    && go build -o server .

FROM alpine:latest


WORKDIR /go/src/github.com/sxz799/job-log/server

COPY --from=0 /go/src/github.com/sxz799/job-log/server/server ./
COPY --from=0 /go/src/github.com/sxz799/job-log/server/conf.yaml ./


EXPOSE 6000
ENTRYPOINT ./server
