#!/bin/bash

APP_NAME="awxhelper"

echo 'export PATH="$HOME/.awxhelper/bin:$PATH"' >> ${HOME}/.$(basename "$SHELL")rc

source ${HOME}/.$(basename "$SHELL")rc

mkdir -p ${HOME}/.${APP_NAME}/bin
