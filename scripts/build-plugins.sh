#!/bin/bash

cwd="${PWD}"

function finish {
    cd "$cwd"
}

trap finish EXIT ERR

find plugins -d 1 -type d | while read line; do
    name=$(basename $line)
    
    GO111MODULE=auto go build -buildmode=plugin -o "${name}.so" "$cwd/$line"
done
