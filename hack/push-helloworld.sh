#!/usr/bin/env bash

set -o nounset
set -o errexit
set -o pipefail
REPO_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd -P)"

KO_DOCKER_REPO=localhost:5001/demo ko publish --base-import-paths ${REPO_ROOT}/cmd/helloworld
