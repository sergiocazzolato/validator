#!/bin/sh
set -x

MODELS_DIR="$HOME/Desktop/validate_image/models"
IMAGES_DIR="./images"

create(){
    rm -rf "$IMAGES_DIR" && mkdir -p "$IMAGES_DIR"
    for platform in pc pi3 pi2 dragonboard; do
        output="${IMAGES_DIR}/${platform}-edge.img"
        sudo /snap/bin/ubuntu-image -c edge -o "$output" "$MODELS_DIR/${platform}.model"
    done
}

create
