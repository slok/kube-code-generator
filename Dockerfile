FROM golang:1.9-alpine

ARG CODEGEN_VERSION="1.9.1"

RUN apk --no-cache add \
    bash \
    git \
    openssl

# Code generator stuff
# Check: https://github.com/kubernetes/kubernetes/pull/57656
RUN wget http://github.com/kubernetes/code-generator/archive/kubernetes-${CODEGEN_VERSION}.tar.gz && \
    mkdir -p /go/src/k8s.io/code-generator/ && \
    tar zxvf kubernetes-${CODEGEN_VERSION}.tar.gz --strip 1 -C /go/src/k8s.io/code-generator/ && \
    mkdir -p /go/src/k8s.io/kubernetes/hack/boilerplate/ && \
    touch /go/src/k8s.io/kubernetes/hack/boilerplate/boilerplate.go.txt

# Create user
ARG uid=1000
ARG gid=1000
RUN addgroup -g $gid codegen && \
    adduser -D -u $uid -G codegen codegen && \
    chown codegen:codegen -R /go


USER codegen