FROM golang:1.26.0-alpine3.23

ARG GOPATH

ENV GOPATH=${GOPATH:-/usr/local/gogetssl}

RUN apk update && \
    apk add vim \
            doas && \
    echo -e "permit nopass :root\npermit nopass :wheel" > /etc/doas.conf && \
    chmod 600 /etc/doas.conf && \
    go install golang.org/x/tools/cmd/goimports@latest

CMD [ "/bin/sh" ]
