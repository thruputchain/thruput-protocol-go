#!/bin/sh

set -e

# Files to exclude from XDR generation (space-separated list of filenames, e.g., "file1.x file2.x")
EXCLUDE_FILES="transaction_trace.x ffi_enums.x invoke_flag.x executor_to_db_rpc.x groundhog_contract_spec.x"

# generate pkg.go

X_FILES=""
for file in protocol/*.x; do
    filename=$(basename "$file")
    exclude=0
    for excl in $EXCLUDE_FILES; do
        if [ "$filename" = "$excl" ]; then
            exclude=1
            break
        fi
    done
    if [ $exclude -eq 0 ]; then
        X_FILES="$X_FILES $file"
    fi
done
X_FILES=$(echo "$X_FILES" | sed 's/^ *//')
PKG_NAME="proto"

cat <<EOF > pkg.go
package $PKG_NAME

// XDR man page: https://xdrpp.github.io/goxdr/pkg/github.com/xdrpp/goxdr/cmd/goxdr/goxdr.1.html

//go:generate goxdr -o generated.go -p $PKG_NAME $X_FILES
EOF

# generate hogp.go

SCRIPT_DIR=$(dirname $(readlink -f "$0"))
cd ${SCRIPT_DIR}

proto_commit=$(git -C ./protocol rev-parse HEAD)
quoted=$(printf '"%s"\n' "$proto_commit")

cat > hogp.go <<EOF
package proto

const ProtocolGitCommit string = ${quoted}

func HogProtocol() string {
	return "hogp/" + ProtocolGitCommit[:12]
}
EOF

# build go

go mod tidy
GOWORK=off go generate
go build
