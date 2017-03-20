#!/bin/bash
set -x

CHANNEL=${1:-edge}
PLATFORMS=${2:-"dragonboard pc-amd64 pc-i386 pi3 pi2"}

for platform in $PLATFORMS; do
    # download stable kernel
    if [[ "$platform" == pi* ]]; then
        kernel_snap_name="pi2-kernel"
        arch="armhf"
    elif [ "$platform" = dragonboard ]; then
        arch="arm64"
        kernel_snap_name="dragonboard-kernel"
    elif [ "$platform" = pc-amd64 ]; then
        arch="amd64"
        kernel_snap_name="pc-kernel"
    else
        arch="i386"
        kernel_snap_name="pc-kernel"
    fi
    download_url=$(curl -s -H "X-Ubuntu-Architecture: $arch" \
                        -H 'X-Ubuntu-Series: 16' \
                        https://search.apps.ubuntu.com/api/v1/snaps/details/${kernel_snap_name}?channel=stable | jq -j '.anon_download_url')
    kernel_snap="./snaps/${platform}-kernel.snap"
    sudo rm -f "$kernel_snap"
    curl -L "$download_url" -o "$kernel_snap"

    if [[ "$platform" == pc* ]]; then
        image_option="--image-size 3G"
    else
        image_option=""
    fi
    output="./images/${platform}-${CHANNEL}"
    sudo rm -rf "$output" && mkdir -p "$output"
    sudo /snap/bin/ubuntu-image "$image_option" -c "$CHANNEL" -O "$output" --extra-snaps "$kernel_snap" "./models/${platform}.model"
done
