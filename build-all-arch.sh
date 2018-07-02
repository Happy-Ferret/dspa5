#!/bin/sh

GOOS=linux GOARCH=arm GOARM=5 make
mkdir -p dist/linux-arm7/
mv dist/dspa-server dist/linux-arm7/
mv dist/dspa-client dist/linux-arm7/

GOOS=linux GOARCH=amd64 make
mkdir -p dist/linux-amd64/
mv dist/dspa-server dist/linux-amd64/
mv dist/dspa-client dist/linux-amd64/

GOOS=darwin GOARCH=amd64 make
mkdir -p dist/darwin-amd64/
mv dist/dspa-server dist/darwin-amd64/
mv dist/dspa-client dist/darwin-amd64/
