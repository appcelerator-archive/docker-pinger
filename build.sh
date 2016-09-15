#!/bin/bash
# invokes make.sh to build an alpine binary, then it builds an alpine-based image with the binary
#
set -e
./make.sh
docker build -t appcelerator/pinger .
