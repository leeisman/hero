#!/usr/bin/env bash
. ./deployments/cloud_run/env/hero_server.env

gcloud run deploy "${SERVICE_NAME}" \
      --image "${IMAGE}" \
      --allow-unauthenticated \
      --cpu 2000m \
      --memory 1024Mi \
      --max-instances 20 \
      --platform managed \
      --region  asia-northeast1\
      --project "${PROJECT_NAME}" \
      --set-env-vars  ^$$^ENV="${ENV}" \
      --set-env-vars  ^$$^MAX_IDLE_CONNS="${MAX_IDLE_CONNS}" \
      --set-env-vars  ^$$^MAX_OPEN_CONNS="${MAX_OPEN_CONNS}" \

