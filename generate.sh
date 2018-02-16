#!/bin/bash


GO=$(which go)
PACKR=$(which packr)

[ -z "$GO" -o -z "$PACKR" ] && echo "Missing required go binaries" && exit 1

if [ $# -ne 1 ]; then
    echo "Usage: ./generate.sh scripts/script_name.sh"
    exit 1
fi
[ ! -f $1 ] && echo "No such file $1" && exit 1

folder=$(dirname $1)
file=$(basename $1)


sed "s,PLACEHOLDER_FOLDER,${folder},g; s,PLACEHOLDER_SCRIPT,${file},g" \
src/main.src > main.go

packr
go build .
