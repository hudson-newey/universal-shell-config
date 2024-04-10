#!/bin/bash
set -eou pipefail

if [ "$#" -lt 1 ]; then
  echo "Usage: $0 <file> [options]"
  exit 1
fi

pkl-gen-go $1 --base-path github.com/myorg/myteam > /dev/null
go run main.go $@
