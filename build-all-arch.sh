#!/bin/sh

GOOS=linux GOARCH=arm GOARM=5 make
mkdir -p dist/linux-arm7/
mv dist/dspa-client dist/linux-arm7/
mv dist/dspa-broadcaster dist/linux-arm7/
mv dist/dspa-speaker dist/linux-arm7/

GOOS=linux GOARCH=amd64 make
mkdir -p dist/linux-amd64/
mv dist/dspa-client dist/linux-amd64/
mv dist/dspa-broadcaster dist/linux-amd64/
mv dist/dspa-speaker dist/linux-amd64/

GOOS=darwin GOARCH=amd64 make
mkdir -p dist/darwin-amd64/
mv dist/dspa-client dist/darwin-amd64/
mv dist/dspa-broadcaster dist/darwin-amd64/
mv dist/dspa-speaker dist/darwin-amd64/

#GOOS=windows GOARCH=amd64 make
#mkdir -p dist/windows-amd64/
#mv dist/dspa-client.exe dist/windows-amd64/
#mv dist/dspa-broadcaster.exe dist/windows-amd64/
#mv dist/dspa-speaker.exe dist/windows-amd64/
