#!/bin/bash

echo 'export PATH="$HOME/.awxhelper/bin:$PATH"' >> ${HOME}/.$(basename "$SHELL")rc

source ${HOME}/.$(basename "$SHELL")rc

mkdir -p ${HOME}/.awxhelper/bin
