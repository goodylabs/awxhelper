#!/bin/bash

APP="awxhelper"
RELEASE_URL="https://api.github.com/repos/goodylabs/awxhelper/releases/latest"

APP_DIR="${HOME}/.${APP}"
APP_BIN_DIR="${APP_DIR}/bin"
APP_BIN_PATH="${APP_BIN_DIR}/${APP}"

mkdir $APP_DIR
mkdir $APP_BIN_DIR

os_type=$(uname -s | tr '[:upper:]' '[:lower:]')
arch=$(uname -m)

if [ "$arch" = "x86_64" ]; then
    arch="amd64"
elif [ "$arch" = "aarch64" ] || [ "$arch" = "arm64" ]; then
    arch="arm64"
else
    echo "Unsupported architecture: $arch"
    exit 1
fi

artifact_url=$(curl -s "$RELEASE_URL" | jq -r ".assets[] | select(.name | test(\"${os_type}-${arch}\")) | .browser_download_url")

if [[ -z "$artifact_url" ]]; then
    echo "No compatible binary found for ${os_type}-${arch}"
    exit 1
fi

echo "Downloading ${APP} binary for ${os_type}-${arch}..."

curl --fail --location --progress-bar --compressed --retry 3 --retry-delay 5 \
  --max-time 10 -o "$APP_BIN_PATH" "$artifact_url"

chmod +x "$APP_BIN_PATH"

echo "${APP} has been installed successfully!"
