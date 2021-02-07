FROM golang:1.15-alpine AS STAGE_BUILD

ARG SERVICE

RUN apk update && apk add make git

WORKDIR /go/src/github.com/freemiumvpn/${SERVICE}/

ADD go.sum go.mod Makefile /go/src/github.com/freemiumvpn/${SERVICE}/
RUN make install

ADD . /go/src/github.com/freemiumvpn/${SERVICE}/

RUN make test
RUN make build
RUN mv ./${SERVICE} /${SERVICE}
RUN mv ./setup /setup/


FROM alpine:3.12  AS STAGE_SERVE

RUN apk update && apk add curl \
    openssl \
    easy-rsa \
    openvpn \
    iptables \
    bash && \
    rm -rf /tmp/* \
    /var/tmp/* \
    /var/cache/apk/*

ARG SERVICE

WORKDIR /work_dir

COPY --from=STAGE_BUILD /${SERVICE} /work_dir/${SERVICE}
COPY --from=STAGE_BUILD /setup /etc/openvpn/setup/

RUN mkdir -p /dev/net && \
    mknod /dev/net/tun c 10 200
