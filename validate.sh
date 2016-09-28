#!/bin/sh
set -x

MODELS_DIR="$HOME/Desktop/validate_image"
IMAGES_DIR="$MODELS_DIR/images"

validate(){
    # only amd64 for now :)
    kvm -m 1024 -redir :8022::22 -snapshot "$IMAGES_DIR/pc-test.img" &
    export SPREAD_ADHOC_UC_ADDRESS=localhost:8022
    # need to setup user!
    #spread -v external-ubuntu-core:ubuntu-core-16-device-amd64
    #kill $!
}

create
