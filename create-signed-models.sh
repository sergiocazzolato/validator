#!/bin/bash

MODELS_PATH="./models"

for model in $(find $MODELS_PATH/*.json); do
    cat "$model" | snap sign -k default >"${model/-model.json/.model}" 2>&1
done
