#!/bin/bash

CMD=$1
shift

TARGET="./cmd/$CMD"

if [ ! -d "$TARGET" ]; then
  echo "❌ Command '$CMD' does not exist."
  exit 1
fi

go run "$TARGET" "$@"
