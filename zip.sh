#!/bin/bash

rm -rf gctl-bin
mkdir gctl-bin

cp gctl gctl-bin/
cp install.sh gctl-bin/

rm gctl-bin.zip

zip gctl-bin.zip gctl-bin

rm -rf gctl-bin
