#!/usr/bin/env bash

response=$(
    curl -s -X POST \
        -H 'Content-Type: application/json' \
        -d '{"code": "ABCDE-12345-FGIJH"}' \
        "$(terragrunt output -json | jq -r '.runner_url.value')"
)
test "$(echo $response | jq -r '.code')" = "ABCDE-12345-FGIJH"
test "$(echo $response | jq -r '.link')" = "https://reallycoolwebsite.com/download/abcde-12345-fgijh"
test "$(echo $response | jq -r '.success')" = "true"
test "$(echo $response | jq -r '.message')" = "Success"
