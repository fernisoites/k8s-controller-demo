#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

if [[ "${1:-}" == "--create-cluster" ]]; then
  gcloud beta container \
    --project "chao-play-project" \
    clusters create "knative-demo" \
    --zone "us-central1-c" \
    --no-enable-basic-auth
fi

# Make sure to install on demo cluster
CONTEXT="gke_chao-play-project_us-central1-c_knative-demo"

# Follow https://knative.dev/docs/install/yaml-install/serving/install-serving-with-yaml/#prerequisites
kubectl --context=$CONTEXT apply -f https://github.com/knative/serving/releases/download/knative-v1.3.0/serving-crds.yaml
kubectl --context=$CONTEXT apply -f https://github.com/knative/serving/releases/download/knative-v1.3.0/serving-core.yaml
kubectl --context=$CONTEXT apply -f https://github.com/knative/net-kourier/releases/download/knative-v1.3.0/kourier.yaml
kubectl --context=$CONTEXT patch configmap/config-network \
  --namespace knative-serving \
  --type merge \
  --patch '{"data":{"ingress.class":"kourier.ingress.networking.knative.dev"}}'
kubectl --context=$CONTEXT apply -f https://github.com/knative/serving/releases/download/knative-v1.3.0/serving-default-domain.yaml
