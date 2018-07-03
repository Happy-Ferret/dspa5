#!/bin/sh

if [ $# -eq 0 ]; then
    echo "Usage: $0 <tag>"
    echo "Release version required as argument"
    exit 1
fi

./build-all-arch.sh


# github does not allow releases with the same filename, so a workaround is
# required
mkdir -p dist/tmp
ln -s ../linux-arm7/dspa-server dist/tmp/dspa-server-linux-arm7
ln -s ../linux-arm7/dspa-client dist/tmp/dspa-client-linux-arm7
ln -s ../linux-amd64/dspa-server dist/tmp/dspa-server-linux-amd64
ln -s ../linux-amd64/dspa-client dist/tmp/dspa-client-linux-amd64
ln -s ../darwin-amd64/dspa-server dist/tmp/dspa-server-darwin-amd64
ln -s ../darwin-amd64/dspa-client dist/tmp/dspa-client-darwin-amd64


hub release create -d \
    -a dist/tmp/dspa-server-linux-arm7#"DSPA server linux-arm7" \
    -a dist/tmp/dspa-client-linux-arm7#"DSPA client linux-arm7" \
    -a dist/tmp/dspa-server-linux-amd64#"DSPA server linux-amd64" \
    -a dist/tmp/dspa-client-linux-amd64#"DSPA client linux-amd64" \
    -a dist/tmp/dspa-server-darwin-amd64#"DSPA server darwin-amd64" \
    -a dist/tmp/dspa-client-darwin-amd64#"DSPA client darwin-amd64" \
    $1

rm -rf dist/tmp
