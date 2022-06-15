#!/bin/bash

VERSION="0.0.1"
if [[ "$OSTYPE" == "darwin"* ]]; then
  OS="darwin"
else
  OS="linux"
fi

ARCH="arm64"
if [[ "$(uname -m)" == "aarch64" ]]; then
  ARCH="arm64"
elif [[ "$(uname -m)" != $ARCH ]]; then
  ARCH="amd64"
fi

URL="https://github.com/ivan-leschinsky/runapp/releases/download/v${VERSION}/runapp-${OS}-${ARCH}"
BIN_PATH="/usr/local/bin/runapp"

if [ -e $BIN_PATH ]
then
  echo "Removing already installed version"
  rm $BIN_PATH
fi

echo "Downloading and installing runapp v${VERSION} for ${OS} ${ARCH}"
curl -sL $URL -o $BIN_PATH && chmod +x $BIN_PATH
echo "Done. Installed into ${BIN_PATH}"
