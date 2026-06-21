# kube-code-generator

A CLI tool and Docker image that wraps `kubernetes/code-generator` and
`kubernetes-sigs/controller-tools` to generate CRD Go client code and CRD
manifests from annotated Go API types. It abstracts the underlying bash
scripts (`kube_codegen.sh`) and binary invocations (`controller-gen`) behind
a simple CLI with flags like `--apis-in`, `--go-gen-out`, `--crd-gen-out`.

The Docker image bundles both upstream tools at pinned versions, providing a
hermetic and reproducible code-generation environment.

# Upgrade to a new Kubernetes version

This document describes the process to upgrade kube-code-generator to support a new Kubernetes version (e.g. v1.36 → v1.37).

## 1. Research upstream versions

Find the target tags for the new Kubernetes version:

- **code-generator**: https://github.com/kubernetes/code-generator/tags
  - Look for tag `kubernetes-<new-version>` (e.g. `kubernetes-1.37.0`)
- **controller-tools**: https://github.com/kubernetes-sigs/controller-tools/tags
  - Look for tag `v0.<xx>.0` that bumps k8s.io/* to the target version
  - Check the release notes to confirm which k8s version it supports

## 2. Check for breaking changes in `kube_codegen.sh`

Compare the `kube_codegen.sh` script between old and new versions:

```
https://raw.githubusercontent.com/kubernetes/code-generator/refs/tags/kubernetes-<old>/kube_codegen.sh
https://raw.githubusercontent.com/kubernetes/code-generator/refs/tags/kubernetes-<new>/kube_codegen.sh
```

Pay attention to:
- Changes to `kube::codegen::gen_client` function signature (flags, arguments)
- Changes to `kube::codegen::gen_helpers` function signature
- New flags that might be useful to expose (e.g. `--prefers-protobuf`)

## 3. Update `docker/prod/Dockerfile`

- Bump base image `golang:X.Y` if needed
- Change `CODEGEN_VERSION` to the new code-generator version
- Change `CONTROLLER_GEN_VERSION` to the new controller-tools version

## 4. Update `docker/dev/Dockerfile`

- Bump base image `golang:X.Y` if needed
- Update `GOLANGCI_LINT_VERSION` and `MOCKERY_VERSION` to latest stable

## 5. Update `.github/workflows/ci.yaml`

- Update `golangci/golangci-lint:v<old>-alpine` to match the new dev Dockerfile version

## 6. Update Go version in go.mod

- Change `go X.Y` to the new minimum Go version

## 7. Update `_example/go.mod`

- Bump `go` directive to match root go.mod
- Bump `k8s.io/apimachinery`, `k8s.io/client-go`, and `k8s.io/api` to the new Kubernetes version
- Run `cd _example && go mod tidy` to refresh indirect deps

## 8. Update `_example/Makefile`

- Change image tag to the new release version

## 9. Update `README.md`

- Bump the shield badge to the new Kubernetes version
- Add a new row to the compatibility table

## 10. Update `CHANGELOG.md`

- Add a new version section under `[Unreleased]`
- Add the version comparison links at the bottom

## 11. Verify

```bash
# Root project
go mod tidy
go test -race ./...
go build ./cmd/kube-code-generator/

# Example
cd _example
go mod tidy
go build ./...

# Docker build
make build-image

# Full generation test (using locally built image)
cd _example
# Update the IMAGE tag in Makefile to point to the local build, then:
make generate
# Verify generated client code compiles with new k8s deps
go build ./...
```
