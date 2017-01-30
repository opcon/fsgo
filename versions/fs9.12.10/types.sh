#!/bin/bash

FSPATH=${FSPATH:-"/usr2/fs"}

read -r -d '' header <<EOF
// +build ignore
// generated with types.sh --- do not edit

package fs

/*
#include "$FSPATH/include/params.h"
#include "$FSPATH/include/fs_types.h"
#include "$FSPATH/include/fscom.h"
*/
import "C"

EOF

f() {
    echo "$header"

    echo "type ("
    echo "$header"\
        | grep include \
        | cpp \
        | grep -o 'struct[[:space:]]\+\w\+' \
        | cut -d" "  -f2 \
        | sort -u \
        | sed -r 's/^(.+)$/\t\u\1\tC.struct_\1/' -
    echo ")"
};
f | gofmt > types.go
