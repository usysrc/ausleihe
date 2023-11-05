#!/bin/sh
air --build.cmd "go build -o ./bin/api ./cmd/app/" --build.bin "./bin/api"
