#!/bin/bash
set -x

CHANNEL=${1:-edge}
VERSION=${2:-"16"}
SNAPS=${3:-""}
PLATFORMS=${4:-"dragonboard pc-amd64 pc-i386 pi3 pi2"}

for platform in $PLATFORMS; do
    image_option=""
    if [[ "$platform" == pc* ]]; then
        image_option="--image-size 3G"
    fi

    snaps=""
    if [ -n "$SNAPS" ]; then
        for snap in $SNAPS; do
            snaps="$snaps --snap $snap"
        done
    fi

    output="./images/${platform}-${VERSION}-${CHANNEL}"
    sudo rm -rf "$output" && mkdir -p "$output"
    sudo ubuntu-image "$image_option" "$snaps" \
         -c "$CHANNEL" \
         -O "$output" \
         "./models/${platform}-${VERSION}.model"
done
