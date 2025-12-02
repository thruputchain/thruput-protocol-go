#!/bin/sh

set -e

# generate pkg.go

X_FILES=$(ls -1 protocol/*.x | tr '\n' ' ')
PKG_NAME="proto"

cat <<EOF > pkg.go
package $PKG_NAME

// XDR man page: https://xdrpp.github.io/goxdr/pkg/github.com/xdrpp/goxdr/cmd/goxdr/goxdr.1.html

//go:generate goxdr -o generated.go -p $PKG_NAME $X_FILES
EOF

# generate hogp.go

SCRIPT_DIR=$(dirname $(readlink -f "$0"))
cd ${SCRIPT_DIR}

proto_commit=$(git submodule status --cached -- ./protocol | awk '{print $1}' | cut -c2-)
quoted=$(printf '"%s"\n' "$proto_commit")

cat > hogp.go <<EOF
package proto

func ProtocolGitCommit() string {
	return ${quoted}
}

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit()[:12]
}
EOF

# build go

go mod tidy
go generate
go build
