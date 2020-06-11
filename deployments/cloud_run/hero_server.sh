#!/usr/bin/env bash
. ./deployments/cloud_run/env/hero_server.env

gcloud run deploy "${SERVICE_NAME}" \
      --image "${IMAGE}" \
      --allow-unauthenticated \
      --cpu 1000m \
      --memory 128Mi \
      --max-instances 1 \
      --platform managed \
      --region  asia-northeast1\
      --project "${PROJECT_NAME}" \
      --set-env-vars  ^$$^ENV="${ENV}" \

