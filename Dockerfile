FROM golang:1.12
ARG CODEGEN_VERSION="1.14.2"

RUN apt-get update && \
    apt-get install -y \
    git 

# Code generator stuff
RUN wget http://github.com/kubernetes/code-generator/archive/kubernetes-${CODEGEN_VERSION}.tar.gz && \
    mkdir -p /go/src/k8s.io/code-generator/ && \
    tar zxvf kubernetes-${CODEGEN_VERSION}.tar.gz --strip 1 -C /go/src/k8s.io/code-generator/ && \
    rm kubernetes-${CODEGEN_VERSION}.tar.gz && \
    \
    wget http://github.com/kubernetes/apimachinery/archive/kubernetes-${CODEGEN_VERSION}.tar.gz && \
    mkdir -p /go/src/k8s.io/apimachinery/ && \
    tar zxvf kubernetes-${CODEGEN_VERSION}.tar.gz --strip 1 -C /go/src/k8s.io/apimachinery/ && \
    rm kubernetes-${CODEGEN_VERSION}.tar.gz && \
    \
    wget http://github.com/kubernetes/api/archive/kubernetes-${CODEGEN_VERSION}.tar.gz && \
    mkdir -p /go/src/k8s.io/api/ && \
    tar zxvf kubernetes-${CODEGEN_VERSION}.tar.gz --strip 1 -C /go/src/k8s.io/api/ && \
    rm kubernetes-${CODEGEN_VERSION}.tar.gz && \
    \
    go get k8s.io/kube-openapi/cmd/openapi-gen


# Create user
ARG uid=1000
ARG gid=1000
RUN addgroup --gid $gid codegen && \
    adduser --gecos "First Last,RoomNumber,WorkPhone,HomePhone" --disabled-password --uid $uid --ingroup codegen codegen && \
    chown codegen:codegen -R /go

COPY hack /hack
RUN chown codegen:codegen -R /hack


USER codegen

WORKDIR /hack
CMD ["./update-codegen.sh"]