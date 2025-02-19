FROM golang:1.23.3-alpine3.20

WORKDIR /go/app

COPY ./ /go/app

RUN set -eux; \
    go env -w GOPROXY=https://goproxy.cn,direct; \
    go mod tidy; \
	go build main.go;

FROM alpine:3.20

ENV LANG=zh_CN.UTF-8 \
    TZ=Asia/Shanghai

# --from表示从其他阶段拷贝内容到本阶段，0表示从第一个阶段拷贝到本阶段
COPY --from=0 /go/app/main /root/main
COPY conf/local.ini /root/conf/local.ini

RUN set -eux; \
    sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories; \
    apk add --update --no-cache \
    tzdata \
    && rm -rf /var/cache/apk/* /tmp/* \
    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

COPY dockerfile.crontabs /etc/crontabs/root

CMD ["crond", "-f"]

# docker build -t load_azurlane_data:latest -f ./Dockerfile .