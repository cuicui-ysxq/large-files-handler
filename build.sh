#!/usr/bin/env bash

CURDIR=$(dirname "$(realpath "$0")")

SRC_DIR=$CURDIR/src
BIN_DIR=$CURDIR/bin

mkdir -p "$BIN_DIR"

cd "$SRC_DIR"
find . -name main.go -print0 | xargs -0 -I "{}" bash -c 'go build -o "'$BIN_DIR'/$(basename "$(dirname "{}")")" "{}"'