#!/bin/sh
APP=goinsta
VERSION="${VERSION:=0.0.0}"
RELEASES=release
BIN=bin
build(){
    os=$1
    arch=$2
    export GOOS=$os
    export GOARCH=$arch
    #Disable dynamic link to libraries
    export CGO_ENABLED=0
    echo "BUILDING FOR $os $arch"
    cd bin
    zip -r ../${RELEASES}/${APP}_${VERSION}_${GOOS}_${GOARCH}.zip *
    cd -
    go env GOOS GOARCH
}

rm -rf ${RELEASES}
mkdir ${RELEASES}
make ${BIN}
rm -rf  ${BIN}
mkdir  ${BIN}
mv ${APP} ./ ${BIN}

build linux amd64
build linux 386
build linux arm
build freebsd amd64
build freebsd 386
build freebsd arm
build openbsd amd64
build openbsd 386
build darwin amd64
build windows amd64
build windows 386
build solaris amd64

cd release
sha256sum * > ${APP}_${VERSION}_SHA256SUMS
cd -