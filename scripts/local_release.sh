#!/bin/bash

APP="awxhelper"

go build -o bin/${APP} main.go

mv bin/${APP} "$HOME/.${APP}/bin/${APP}"
