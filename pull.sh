#!/usr/bin/env sh

IMAGE_CTL="crictl -n k8s.io"

source env.sh

source ${IMAGE_ENV_PATH}/${IMAGE_ENV_NAME}/params.sh

$("${IMAGE_CTL} image pull ${IMAGE_TO}")

$("${IMAGE_CTL} image tag ${IMAGE_TO} ${IMAGE_FROM}")
