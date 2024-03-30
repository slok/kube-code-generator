# Set also `ARCH` ARG here so we can use it on all the `FROM`s. 
ARG ARCH

FROM golang:1.22 as build-stage

# Required by the built script for setting verion and cross-compiling.
ARG VERSION
ENV VERSION=${VERSION}
ARG ARCH
ENV GOARCH=${ARCH}

# Compile.
WORKDIR /src
COPY . .
RUN ./scripts/build/bin/build-raw.sh

FROM golang:1.22
ARG CODEGEN_VERSION="1.30.0-rc.0"
ARG CONTROLLER_GEN_VERSION="0.14.0"


COPY --from=build-stage /src/bin/kube-code-generator /usr/local/bin/kube-code-generator

RUN apt-get update && \
    apt-get install -y \
    git 

# Code generator stuff
RUN wget http://github.com/kubernetes/code-generator/archive/kubernetes-${CODEGEN_VERSION}.tar.gz && \
    mkdir -p /tmp/k8s-code-generator/ && \
    tar zxvf kubernetes-${CODEGEN_VERSION}.tar.gz --strip 1 -C /tmp/k8s-code-generator/ && \
    cd /tmp/k8s-code-generator/ && go mod tidy && cd - && \
    rm kubernetes-${CODEGEN_VERSION}.tar.gz && \
    \
    wget https://github.com/kubernetes-sigs/controller-tools/archive/v${CONTROLLER_GEN_VERSION}.tar.gz && \
    tar xvf ./v${CONTROLLER_GEN_VERSION}.tar.gz && \
    cd ./controller-tools-${CONTROLLER_GEN_VERSION}/ && \
    go mod tidy && \
    go build -o controller-gen  ./cmd/controller-gen/ && \
    mv ./controller-gen /usr/bin/ && \
    rm -rf ../v${CONTROLLER_GEN_VERSION}.tar.gz && \
    rm -rf ../controller-tools-${CONTROLLER_GEN_VERSION}


# Create user and workdir.
ARG uid=1000
ARG gid=1000
RUN addgroup --gid $gid app && \
    adduser --gecos "First Last,RoomNumber,WorkPhone,HomePhone" --disabled-password --uid $uid --ingroup app app && \
    chown app:app -R /go
RUN mkdir /app && chown app:app -R /app

USER app
WORKDIR /app

ENV KUBE_CODE_GENERATOR_CODEGEN_PATH=/tmp/k8s-code-generator
ENTRYPOINT ["kube-code-generator"]