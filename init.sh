#!/usr/bin/env bash

set -e
set -x

package="github.com/corpix/go-boilerplate"
cwd="$(pwd)"

git clone "https://${package}" .
rm -rf init.sh

find . -type f | \
    xargs perl -p -i -e "s|${package}|${pwd/$GOPATH\/src\//}|g"

echo Done.
