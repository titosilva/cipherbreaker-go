#!/bin/bash

SRC="github.com/titosilva/cipherbreaker-go/"

if [[ "$1" == "bytestring" ]]
then
    echo "Compiling $1..."
    # Compile bytestring package
    go build "$SRC/pkg/bytestring"
    echo "Done"
fi

if [[ "$1" == "coder" ]]
then
    echo "Compiling $1..."
    # Compile bytestring package
    go build "$SRC/pkg/bytestring/coder"
    echo "Done"
fi