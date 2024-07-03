#!/usr/bin/env bash

CURDIR=$(dirname "$(realpath "$0")")

SRC_DIR=$CURDIR/src
BIN_DIR=$CURDIR/bin

mkdir -p "$BIN_DIR"

cd "$SRC_DIR"
find . -name main.go -print0 | xargs -0 -I "{}" bash -c 'DIR=$(dirname "{}"); EXEC_NAME=$(basename "$DIR"); go build -C "$DIR" -o "'"$BIN_DIR/$EXEC_NAME"'"'