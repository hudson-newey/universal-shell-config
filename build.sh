#!/bin/bash
set -eou pipefail

pkl-gen-go pkl/Config.pkl --base-path github.com/myorg/myteam > /dev/null
go run main.go
