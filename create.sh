#!/bin/bash
set -x

CHANNEL=${1:-edge}
PLATFORMS=${2:-"dragonboard pc-amd64 pc-i386 pi3 pi2"}
VERSION=${3:-"16"}

for platform in $PLATFORMS; do
    if [[ "$platform" == pc* ]]; then
        image_option="--image-size 3G"
    else
        image_option=""
    fi
    output="./images/${platform}-${CHANNEL}"
    sudo rm -rf "$output" && mkdir -p "$output"
    sudo /usr/bin/ubuntu-image "$image_option" \
         -c "$CHANNEL" \
         -O "$output" \
         "./models/${platform}-${VERSION}.model"
done
