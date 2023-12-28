#!/bin/bash
set -eo pipefail
#ARTIFACT_BUCKET=$(cat bucket-name.txt)
ARTIFACT_BUCKET=br.com.likwi.artifacts.apps.backend-site-dev
STACK_NAME=backend-site-stack
go mod tidy
CGO_ENABLED=0 GOOS=linux go build
pwd
aws cloudformation package --template-file template.yaml --s3-bucket $ARTIFACT_BUCKET --output-template-file out.yaml
aws cloudformation deploy --template-file out.yaml --stack-name $STACK_NAME --capabilities CAPABILITY_NAMED_IAM
