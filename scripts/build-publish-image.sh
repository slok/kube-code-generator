#!/usr/bin/env bash

set -o errexit
set -o nounset

[ -z "$VERSION" ] && echo "VERSION env var is required." && exit 1
[ -z "$IMAGE" ] && echo "IMAGE env var is required." && exit 1

./scripts/build-image.sh
./scripts/publish-image.sh
