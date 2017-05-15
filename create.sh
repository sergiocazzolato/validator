#!/bin/bash
set -x

CHANNEL=${1:-edge}
PLATFORMS=${2:-"dragonboard pc-amd64 pc-i386 pi3 pi2"}
ROOT=/tmp/root
IMAGE="${ROOT}/image"
BASE_DIR="${IMAGE}/var/lib/snapd/seed"
SNAPS_DIR="${BASE_DIR}/snaps"

create_assert_file(){
    local snap_name="$1"
    local assertions_dir="${BASE_DIR}/assertions"

    account_key_file=$(grep -l "name: store" ${assertions_dir}/*.account-key)

    snap_declaration_file=$(grep -l "$snap_name" ${assertions_dir}/*.snap-declaration)

    snap_id=$(cat $snap_declaration_file | sed -n 's|snap-id: \(.*\)|\1|p')

    snap_revision_file=$(grep -l "snap-id: ${snap_id}" ${assertions_dir}/*.snap-revision)

    target_file=$(basename $(ls $(dirname $assertions_dir)/snaps/${snap_name}*))
    target_file_name="${target_file%.*}"

    for f in $account_key_file $snap_declaration_file $snap_revision_file; do (cat "${f}"; echo) >> "${SNAPS_DIR}/${target_file_name}.assert"; done
}

for platform in $PLATFORMS; do
    rm -rf $IMAGE; mkdir -p $IMAGE

    snap prepare-image --channel stable "./models/${platform}.model" $ROOT

    if [[ "$platform" == pi* ]]; then
        kernel_snap_name="pi2-kernel"
    elif [ "$platform" = dragonboard ]; then
        kernel_snap_name="dragonboard-kernel"
    else
        kernel_snap_name="pc-kernel"
    fi

    create_assert_file "$kernel_snap_name"

    if [[ "$platform" == pc* ]]; then
        image_option="--image-size 3G"
    else
        image_option=""
    fi
    output="./images/${platform}-${CHANNEL}"
    sudo rm -rf "$output" && mkdir -p "$output"
    sudo /snap/bin/ubuntu-image "$image_option" \
         -c "$CHANNEL" \
         -O "$output" \
         --extra-snaps ${SNAPS_DIR}/${kernel_snap_name}_*.snap \
         "./models/${platform}.model"
    rm -rf $IMAGE
done
