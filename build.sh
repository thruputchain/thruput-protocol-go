#!/bin/sh

set -e

X_FILES=$(ls -1 protocol/*.x | tr '\n' ' ')
PKG_NAME="proto"

cat <<EOF > proto.go
package $PKG_NAME

// XDR man page: https://xdrpp.github.io/goxdr/pkg/github.com/xdrpp/goxdr/cmd/goxdr/goxdr.1.html

//go:generate goxdr -o generated.go -p $PKG_NAME $X_FILES
EOF

go mod tidy
go generate
go build
