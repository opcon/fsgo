#!/bin/bash

read -r -d '' header <<'EOF'
// +build ignore
// generated with consts.sh --- do not edit

package fs

/*
#include "/usr2/fs/include/params.h"
*/
import "C"

EOF

f() {
	echo "$header"
	echo "const ("
	cpp -E -dM -fpreprocessed "/usr2/fs/include/params.h" \
		| grep "#define" \
		| cut -d" "  -f2 \
		| sort -u \
		| sed -r 's/^(.+)$/\t\1=C.\1/' -
	echo ")"
};
f | gofmt > consts.go
