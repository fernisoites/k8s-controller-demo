#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

REPO_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd -P)"

cd "${REPO_ROOT}/hack"
go build -o "${REPO_ROOT}/_bin/controller-gen" "sigs.k8s.io/controller-tools/cmd/controller-gen"
cd "${REPO_ROOT}"

./_bin/controller-gen \
  crd:crdVersions=v1 \
  paths=./pkg/apis/foo/v1 \
  output:stdout \
  > ./manifest/foo_crd.yaml
