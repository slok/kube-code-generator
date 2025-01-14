# Kube code generator

![Kubernetes release](https://img.shields.io/badge/Kubernetes-v1.32-green?logo=Kubernetes&style=flat&color=326CE5&logoColor=white)

## Introduction

When we speak about Kubernetes operators or controllers, normally Go code, CR, CRDs... are required. To create all the autogenerated Kubernetes Go code (Clients, helpers...) and manifests (CRD), the process is a bit painful.

This small project tries making easy this process, by creating a small layer between Kubernetes official tooling that are used to get all this autogenerated stuff, and abstract options and infer some others, making a better UX for the user.

The projects that are used under the hood are:

- [code-generator](https://github.com/kubernetes/code-generator) for Go code autogeneration.
- [controller-tools](https://github.com/kubernetes-sigs/controller-tools) for CRD autogeneration.

## Why and when use this

- You don't like, need or use kubebuilder for your CRDs.
- You want simple tooling to generate Kubernetes CRD Go clients and manifests.
- You like safe standards and simple things.
- You use CRDs for more/other things than operators (e.g: generating CLIs, storing state on k8s as APIs...).
- You don't want to do hacky and ugly stuff to start creating Kubernetes tooling.

## Features

- Small API/configuration.
- Safe standards
- Ready to use Docker images.
- Generates CR client Go code (Used in controllers and operators).
- Generates CR informers Go code (Used in controllers and operators).
- Generates CR listers Go code (Used in controllers and operators).
- Generates CR ["apply configurations"](https://pkg.go.dev/k8s.io/client-go/applyconfigurations) Go code (Used in controllers and operators).
- Generates CRD manifests (Used for API registration on k8s clusters).

## How to use it

The easiest way is to use the provided Docker image as it has all the required upstream dependencies.

Here is an example that mounts the current directory (a Go project) and generates the Go code and the CRDs by providing the APIs input directory and the generation output directories:

```bash
docker run -it --rm -v ${PWD}:/app ghcr.io/slok/kube-code-generator \
 --apis-in ./apis \
 --go-gen-out ./gen \
 --crd-gen-out ./gen/manifests
```

However, the best way to know how to use it is with a full example, you have it in [_example](_example/) dir.

### Optional features

These are the list of features that can be enabled when generating Go code or CRDs:

- `--apply-configurations`: Generates [apply configurations](https://pkg.go.dev/k8s.io/client-go/applyconfigurations) for CRs require for [server-side-apply](https://kubernetes.io/docs/reference/using-api/server-side-apply/).

## Kubernetes versions

It's suggested that if you don't have an old Kubenretes version, you try the latest kube-code-generator
(`latest`), however the ones described on the table here are known to work correctly on the specific version.

| Kubernetes | Docker image                                           |
| ---------- | ------------------------------------------------------ |
|  v1.32     | `docker pull ghcr.io/slok/kube-code-generator:v0.5.0`  |
|  v1.31     | `docker pull ghcr.io/slok/kube-code-generator:v0.3.2`  |
|  v1.30     | `docker pull ghcr.io/slok/kube-code-generator:v0.2.0`  |
|  v1.29     | `docker pull ghcr.io/slok/kube-code-generator:v0.2.0`  |
|  v1.28     | `docker pull ghcr.io/slok/kube-code-generator:v0.2.0`  |
|  v1.27     | `docker pull ghcr.io/slok/kube-code-generator:v0.2.0`  |

### Versions <v1.27

With the release of Kubernetes v1.30, this app was rewritten from bash hacky scripts into a proper Go application, that is easier, more extendable and safer to use.

In case you tested this new refactor and isn't working for you, Inwould suggest to try the previous versions, check the previous [Readme](https://github.com/slok/kube-code-generator/tree/v1.27.0).
