#!/usr/bin/env bash

set -e

# http://stackoverflow.com/a/21142256/2055281

echo "mode: atomic" > coverage.txt

for d in $(find ./* -maxdepth 10 -type d); do
    if ls $d/*.go &> /dev/null; then
        go test -covermode=atomic $d
    fi
done

