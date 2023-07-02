# Kube code generator

![Kubernetes release](https://img.shields.io/badge/Kubernetes-v1.27-green?logo=Kubernetes&style=flat&color=326CE5&logoColor=white)

A kubernetes code generator container that makes easier to generate CRD manifests and its Go clients.

Uses [official code-generator](https://github.com/kubernetes/code-generator) created by Kubernetes to autogenerate the code required for the CRDs.

## Generation targets

- CRD based Go code (clients, lib...).
- CRD manifest YAMLs to register your CRs on the cluster.

## Docker image versions

|                  | Docker image                                            |
| ---------------- | ------------------------------------------------------- |
| Kubernetes v1.27 | `docker pull ghcr.io/slok/kube-code-generator:v1.27.0`  |
| Kubernetes v1.26 | `docker pull ghcr.io/slok/kube-code-generator:v1.26.0`  |
| Kubernetes v1.25 | `docker pull ghcr.io/slok/kube-code-generator:v1.25.0`  |
| Kubernetes v1.24 | `docker pull ghcr.io/slok/kube-code-generator:v1.24.0`  |
| Kubernetes v1.23 | `docker pull ghcr.io/slok/kube-code-generator:v1.23.0`  |
| Kubernetes v1.22 | `docker pull ghcr.io/slok/kube-code-generator:v1.22.0`  |
| Kubernetes v1.21 | `docker pull ghcr.io/slok/kube-code-generator:v1.21.1`  |
| Kubernetes v1.20 | `docker pull ghcr.io/slok/kube-code-generator:v1.20.1`  |
| Kubernetes v1.19 | `docker pull ghcr.io/slok/kube-code-generator:v1.19.2`  |
| Kubernetes v1.18 | `docker pull ghcr.io/slok/kube-code-generator:v1.18.0`  |
| Kubernetes v1.17 | `docker pull ghcr.io/slok/kube-code-generator:v1.17.3`  |
| Kubernetes v1.16 | `docker pull ghcr.io/slok/kube-code-generator:v1.16.7`  |
| Kubernetes v1.15 | `docker pull ghcr.io/slok/kube-code-generator:v1.15.10` |
| Kubernetes v1.14 | `docker pull ghcr.io/slok/kube-code-generator:v1.14.2`  |
| Kubernetes v1.13 | `docker pull ghcr.io/slok/kube-code-generator:v1.13.5`  |
| Kubernetes v1.12 | `docker pull ghcr.io/slok/kube-code-generator:v1.12.4`  |
| Kubernetes v1.11 | `docker pull ghcr.io/slok/kube-code-generator:v1.11.3`  |
| Kubernetes v1.10 | `docker pull ghcr.io/slok/kube-code-generator:v1.10.0`  |
| Kubernetes v1.9  | `docker pull ghcr.io/slok/kube-code-generator:v1.9.1`   |

## Getting started

The best way to know how to use it is by checking the [example](example/) that will generate the required clients and CRD manifests.

### Optional settings

Some settings are optional so you can customize special cases:

- On CRD manifest YAML generation:
  - `CRD_FLAG` env var to overwrite CRD flag with a custom one. (E.g: Use `allowDangerousTypes=true` to allow `float64` on generation, [more info here][unsecure-float64])

[unsecure-float64]: https://github.com/kubernetes-sigs/controller-tools/issues/245
