#!/bin/bash

go_packages=(
  auth
  components
)

for package in "${go_packages[@]}"; do
  protoc \
    --proto_path=${PROJECT_PATH}/proto \
    --go_out=paths=source_relative:. \
    ${PROJECT_PATH}/proto/${package}/*.proto;
done
