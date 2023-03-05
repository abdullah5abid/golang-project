#!/usr/bin/env bash

source ./scripts/env.sh

while getopts "b" opt; do
  case $opt in
  b) SHOULD_BUILD=true ;;
  esac
done

if [ ! -f $BINARY ] || [ -n "$SHOULD_BUILD" ]; then
  echo "Building $PROJECT_NAME..."
  ./scripts/build.sh || exit 1
fi

echo "Press CTRL+C to exit..."

$BINARY
