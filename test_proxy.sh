#!/usr/bin/env bash

PROMPT="Largest city in Finland? One word answer, please."

curl -s -X POST "http://localhost:3131/v1/chat/completions" \
  -H "Content-Type: application/json" \
  -d "{\"model\": \"phi3-fast:latest\", \"messages\": [{\"role\": \"user\", \"content\": \"${PROMPT}\"}] }"