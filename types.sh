#!/bin/bash

read -r -d '' header <<'EOF'
// +build ignore
// generated with types.sh --- do not edit

package fs

/*
#include "/usr2/fs/include/params.h"
#include "/usr2/fs/include/fs_types.h"
#include "/usr2/fs/include/fscom.h"
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
        | sed -r 's/^(.+)$/\t\1\tC.struct_\1/' -
    echo ")"
};
f | gofmt > types.go
