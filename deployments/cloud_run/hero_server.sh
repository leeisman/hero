#!/usr/bin/env bash
. ./deployments/cloud_run/env/hero_server.env

gcloud run deploy "${SERVICE_NAME}" \
      --image "${IMAGE}" \
      --allow-unauthenticated \
      --cpu 2000m \
      --memory 1024Mi \
      --max-instances 13 \
      --platform managed \
      --region  asia-northeast1\
      --project "${PROJECT_NAME}" \
      --set-env-vars  ^$$^ENV="${ENV}" \

