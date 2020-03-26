FROM golang:1.14
ARG CODEGEN_VERSION="1.18.0"
ARG CONTROLLER_GEN_VERSION="0.2.5"

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
    wget https://github.com/openshift/kubernetes-sigs-controller-tools/releases/download/v${CONTROLLER_GEN_VERSION}/controller-gen-linux-amd64 && \
    mv controller-gen-linux-amd64 /usr/bin/controller-gen && \
    chmod a+x /usr/bin/controller-gen && \
    \
    go get k8s.io/kube-openapi/cmd/openapi-gen

# Create user
ARG uid=1000
ARG gid=1000
RUN addgroup --gid $gid codegen && \
    adduser --gecos "First Last,RoomNumber,WorkPhone,HomePhone" --disabled-password --uid $uid --ingroup codegen codegen && \
    chown codegen:codegen -R /go

COPY hack /hack
RUN chown codegen:codegen -R /hack && \
    mv /hack/* /usr/bin

USER codegen

WORKDIR /usr/bin

CMD ["update-codegen.sh"]