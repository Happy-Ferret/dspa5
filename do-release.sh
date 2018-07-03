#!/bin/sh

if [ $# -eq 0 ]; then
    echo "Usage: $0 <tag>"
    echo "Release version required as argument"
fi

./build-all-arch.sh


hub release create -d \
    -a dist/linux-arm7/dspa-server#"DSPA server linux-arm7" \
    -a dist/linux-arm7/dspa-client#"DSPA client linux-arm7" \
    -a dist/linux-amd64/dspa-server#"DSPA server linux-amd64" \
    -a dist/linux-amd64/dspa-client#"DSPA client linux-amd64" \
    -a dist/darwin-amd64/dspa-server#"DSPA server darwin-amd64" \
    -a dist/darwin-amd64/dspa-client#"DSPA client darwin-amd64" \
    $1
