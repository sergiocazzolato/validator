#!/bin/bash
set -x

CHANNEL=${1:-edge}
PLATFORMS=${2:-"dragonboard pc-amd64 pc-i386 pi3 pi2"}

create_image(){
    local platform=$1
    local output=$2

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
    sudo rm -rf "$output" && mkdir -p "$output"
    sudo /snap/bin/ubuntu-image "$image_option" -c "$CHANNEL" -O "$output" --extra-snaps "$kernel_snap" "./models/${platform}.model"
}

create_netplan_config(){
    local platform=$1
    local unpack=$2

    sudo mkdir -p "$unpack/etc/netplan"

    cat <<EOF | sudo tee "$unpack/etc/netplan/00-snapd-config.yaml"
network:
  ethernets:
    eth0:
      addresses: []
      dhcp4: true
  version: 2
EOF
}

create_test_user(){
    local unpack=$1

    sudo chroot "$unpack" adduser --extrausers --quiet --disabled-password --gecos '' test
    echo test:ubuntu | sudo chpasswd --root "$unpack"
    echo 'test ALL=(ALL) NOPASSWD:ALL' | sudo tee "$unpack/etc/sudoers.d/99-test-user"
}

bootstrap_image(){
    local platform=$1
    local output_base=$2

    if [[ "$platform" == pc* ]]; then
        image_file="pc.img"
    else
        image_file="${platform}.img"
    fi

    trap 'sudo umount $tmp && sudo rm -rf $tmp $(dirname $unpack) $core_dest && sudo kpartx -ds $output_base/$image_file' EXIT
    loops=$(sudo kpartx -avs "$output_base/$image_file" | cut -d' ' -f 3)

    tmp=$(mktemp -d)

    for loop in $loops; do
        sudo mount "/dev/mapper/$loop" "$tmp"

        if [ -d "$tmp/system-data" ]; then
            break
        fi
        sudo umount "$tmp"
    done

    unpack="$(mktemp -d)/core"

    core_file=$(ls $tmp/system-data/var/lib/snapd/snaps/core_*.snap)

    sudo unsquashfs -d "$unpack" "$core_file"

    create_netplan_config "$platform" "$unpack"

    create_test_user "$unpack"

    core_dest=$(mktemp -d)
    sudo /snap/test-snapd-snapbuild/current/bin/snapbuild "$unpack" "$core_dest"
    sudo mv -f $core_dest/core_*.snap "$core_file"
    sudo cp "$core_file" $tmp/system-data/var/lib/snapd/seed/snaps
}

for platform in $PLATFORMS; do
    output_base="./images/${platform}-${CHANNEL}"

    create_image "$platform" "$output_base"
    bootstrap_image "$platform" "$output_base"
done
