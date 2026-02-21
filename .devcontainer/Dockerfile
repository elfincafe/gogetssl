FROM golang:1.26.0-alpine3.23

ENV TZ=Asia/Tokyo

RUN apk update && \
    apk add vim \
            git \
            tzdata \
            doas \
            openssh-client-default && \
    echo -e "permit nopass :root\npermit nopass :wheel" > /etc/doas.conf && \
    chmod 600 /etc/doas.conf && \
    go install golang.org/x/tools/cmd/goimports@latest

WORKDIR ${GOPATH}/src
CMD ["/bin/sh"]
