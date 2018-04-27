#!/usr/bin/env bash

set -e

cd $(dirname $0)

PROTOC_VERSION="3.5.1"

PROTOC_OS="$(uname -s)"
PROTOC_ARCH="$(uname -m)"
case "${PROTOC_OS}" in
  Darwin) PROTOC_OS="osx" ;;
  Linux) PROTOC_OS="linux" ;;
  *)
    echo "Invalid value for uname -s: ${PROTOC_OS}" >&2
    exit 1
esac

PROTOC="./protoc/bin/protoc"

if [[ "$(${PROTOC} --version 2>/dev/null)" != "libprotoc ${PROTOC_VERSION}" ]]; then
  rm -rf ./protoc
  mkdir -p protoc
  curl -L "https://github.com/google/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-${PROTOC_OS}-${PROTOC_ARCH}.zip" > protoc/protoc.zip
  cd ./protoc && unzip protoc.zip && cd ..
fi

go install github.com/golang/protobuf/protoc-gen-go 

PROTODIR=github.com/gogo/protobuf/types
PKGMAP=Mgoogle/protobuf/descriptor.proto=github.com/gogo/protobuf/protoc-gen-gogo/descriptor
PKGMAP=${PKGMAP},Mgoogle/protobuf/any.proto=${PROTODIR}
PKGMAP=${PKGMAP},Mgoogle/protobuf/api.proto=${PROTODIR}
PKGMAP=${PKGMAP},Mgoogle/protobuf/duration.proto=${PROTODIR}
PKGMAP=${PKGMAP},Mgoogle/protobuf/empty.proto=${PROTODIR}
PKGMAP=${PKGMAP},Mgoogle/protobuf/field_mask.proto=${PROTODIR}
PKGMAP=${PKGMAP},Mgoogle/protobuf/struct.proto=${PROTODIR}
PKGMAP=${PKGMAP},Mgoogle/protobuf/timestamp.proto=${PROTODIR}
PKGMAP=${PKGMAP},Mgoogle/protobuf/wrappers.proto=${PROTODIR}

outdir="${GOPATH}/src"
${PROTOC} "--gogo_out=plugins=grpc,${PKGMAP}:$outdir" -I. *.proto
