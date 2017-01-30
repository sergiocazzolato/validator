#!/bin/bash
set -x

MODELS_DIR="./models"
IMAGES_DIR="./images"
CHANNEL=${1:-edge}

create(){
    rm -rf "$IMAGES_DIR" && mkdir -p "$IMAGES_DIR"
    for platform in dragonboard pc pc-i386 pi3 pi2; do
        if [[ "$platform" == pc* ]]; then
            image_option="--image-size 3G"
        else
            image_option=""
        fi
        output="${IMAGES_DIR}/${platform}-${CHANNEL}.img"
        sudo /snap/bin/ubuntu-image "$image_option" -c "$CHANNEL" -o "$output" "$MODELS_DIR/${platform}.model"
    done
}

create
