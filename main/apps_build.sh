#!/usr/bin/env bash

export GOPATH=`pwd`/../../../../../

go build -o ${GOPATH}/src/main/plugins/apps.tea apps.go