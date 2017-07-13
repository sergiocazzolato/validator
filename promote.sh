#!/bin/sh
set -ex

ORIGIN_CHANNEL=${1:-beta}
TARGET_CHANNEL=${2:-candidate}

for rev in $(snapcraft list-revisions core | grep "${ORIGIN_CHANNEL}\*" | cut -d ' ' -f 1); do
    snapcraft release core "$rev" "${TARGET_CHANNEL}"
done
