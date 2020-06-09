#!/usr/bin/env bash
. ./builds/hero/docker.env

VERSION=$(git rev-parse HEAD)
IMAGE_NAME="gcr.io/${PROJECT_NAME}/${HERO_SERVICE}"
docker build . -f ./builds/hero/Dockerfile -t "${IMAGE_NAME}:${VERSION}" -t "${IMAGE_NAME}:latest"
docker push "${IMAGE_NAME}:latest"
