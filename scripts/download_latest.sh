#!/bin/bash
VER=0.2.6
ARCH=amd64

ARTIFACT_NAME="sidebike_${VER}_linux_${ARCH}"
FILE_URL=https://github.com/williamfzc/sidebike/releases/download/v${VER}/${ARTIFACT_NAME}
PROXY_FILE_URL=https://ghproxy.com/${FILE_URL}

echo ${PROXY_FILE_URL}
wget $PROXY_FILE_URL
echo "downloaded: ${ARTIFACT_NAME}"
