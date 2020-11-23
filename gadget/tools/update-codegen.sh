#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

SCRIPT_ROOT=$(dirname "${BASH_SOURCE[0]}")/..
CODEGEN_PKG=${CODEGEN_PKG:-$(cd "${SCRIPT_ROOT}"; ls -d -1 ./vendor/k8s.io/code-generator 2>/dev/null || echo ../code-generator)}

bash "${CODEGEN_PKG}"/generate-groups.sh "deepcopy" \
  github.com/burov/kubernetes-examples/gadget/pkg/generated github.com/burov/kubernetes-examples/gadget/pkg/apis \
  gadget:v1beta1 \
  --go-header-file "${SCRIPT_ROOT}"/tools/boilerplate.go.txt