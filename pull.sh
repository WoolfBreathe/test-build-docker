#!/usr/bin/env bash

IMAGE_CTL="/usr/bin/ctr -n k8s.io"

source env.sh

source ${IMAGE_ENV_PATH}/${IMAGE_ENV_NAME}/params.sh

bash -c "${IMAGE_CTL} image pull ${IMAGE_TO}"

bash -c "${IMAGE_CTL} image tag ${IMAGE_TO} ${IMAGE_FROM}"

