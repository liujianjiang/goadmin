FROM golang:alpine as builder

WORKDIR /data/app/goadmin
COPY . .

RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go env -w CGO_ENABLED=0 \
    && go env \
    && go mod tidy \
    && go build -o server .

FROM alpine:latest

LABEL MAINTAINER="875394153@qq.com"

WORKDIR /data/app/goadmin

COPY --from=0 /data/app/goadmin/server ./
COPY --from=0 /data/app/goadmin/resource ./resource/

EXPOSE 8989
ENTRYPOINT ./server -c config.docker.yaml
