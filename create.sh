#!/bin/sh
set -x

MODELS_DIR="./models"
IMAGES_DIR="./images"
CHANNEL=${1:-edge}

create(){
    rm -rf "$IMAGES_DIR" && mkdir -p "$IMAGES_DIR"
    for platform in pc pi3 pi2 dragonboard; do
        output="${IMAGES_DIR}/${platform}-${CHANNEL}.img"
        sudo /snap/bin/ubuntu-image -c "$CHANNEL" -o "$output" "$MODELS_DIR/${platform}.model"
    done
}

create
