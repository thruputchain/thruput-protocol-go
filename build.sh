#!/bin/sh

set -e

CHECK_GENERATED=0
if [ "${1:-}" = "--check" ]; then
    CHECK_GENERATED=1
    shift
fi
if [ "$#" -ne 0 ]; then
    echo "usage: $0 [--check]" >&2
    exit 2
fi

# Files to exclude from XDR generation (space-separated list of filenames, e.g., "file1.x file2.x")
EXCLUDE_FILES="transaction_trace.x ffi_enums.x invoke_flag.x invoke_args.x executor_to_db_rpc.x groundhog_contract_spec.x repair2.x view.x view_gateway.x"

# generate pkg.go

# goxdr emits declarations in the order it receives files, so the order below
# decides the layout of generated.go. A bare glob takes that order from the
# shell's collation, which differs by machine -- glibc's en_US.UTF-8 ignores
# punctuation and puts dogwood_kit.x before dogwood.x, while C compares bytes
# and puts dogwood.x first -- so regenerating elsewhere reshuffled thousands of
# lines for no semantic change.
#
# Sort explicitly instead, on the name with punctuation removed, which is the
# order generated.go has always been in. Doing it here rather than by setting a
# locale keeps it reproducible on a machine that has no en_US.UTF-8 installed,
# where the setting would silently fall back to C. The key is emitted ahead of
# the path so ties fall back to comparing the path itself.
X_FILES=$(
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
            printf '%s\t%s\n' "$(printf '%s' "$filename" | tr -d '._')" "$file"
        fi
    done | LC_ALL=C sort | cut -f2- | tr '\n' ' '
)
X_FILES=$(echo "$X_FILES" | sed 's/ *$//')
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

if [ "$CHECK_GENERATED" -eq 1 ]; then
    generation_status=$(git status --porcelain --untracked-files=all)
    if [ -n "$generation_status" ]; then
        echo "Generated files are out of date:" >&2
        printf '%s\n' "$generation_status" >&2
        echo "Tracked diff summary:" >&2
        git diff --stat -- >&2
        exit 1
    fi
fi
