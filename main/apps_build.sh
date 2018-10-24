#!/usr/bin/env bash

export GOPATH=`pwd`/../../../../../
#export GOOS=linux
#export GOARCH=amd64

go build -o ${GOPATH}/src/main/plugins/apps.tea apps_plugin.go