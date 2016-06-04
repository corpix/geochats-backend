#!/usr/bin/env bash

set -e
set -x

cwd="$(pwd)"

boiler_package="github.com/corpix/go-boilerplate"
boiler_package_base="$(dirname "$boiler_package")"
boiler_package_name="$(basename "$boiler_package")"

new_package="${cwd/$GOPATH\/src\//}"
package_base="$(dirname "$new_package")"
package_name="$(basename "$new_package")"

git clone "https://${boiler_package}" .
rm -rf init.sh

find . -type f | \
    xargs perl -p -i -e "s|${boiler_package}|${new_package}|g"

perl -p -i -e "s|${boiler_package_base}|${package_base}|g" GNUmakefile
perl -p -i -e "s|${boiler_package_name}|${package_name}|g" GNUmakefile

mv "${boiler_package_name}" "${package_name}"

echo Done.
