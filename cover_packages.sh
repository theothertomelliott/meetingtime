#!/bin/bash

# This script tests multiple packages and creates a consolidated cover profile
# Based on https://github.com/getlantern/flashlight-build/blob/devel/testandcover.bash

function die() {
  echo $*
  exit 1
}

export GOPATH=`pwd`:$GOPATH

# Initialize profile.cov
echo "mode: count" > profile.cov

# Initialize error tracking
ERROR=""

# Test each package and append coverage profile info to profile.cov
for pkg in `go list ./... | grep -v /vendor/ | grep -v /examples/`
do
    go test -v -covermode=count -coverprofile=profile_tmp.cov $pkg || ERROR="Error testing $pkg"
    tail -n +2 profile_tmp.cov >> profile.cov || die "Unable to append coverage for $pkg"
done

rm profile_tmp.cov

if [ ! -z "$ERROR" ]
then
    die "Encountered error, last error was: $ERROR"
fi