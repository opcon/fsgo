// Copyright 2019 NVI Inc. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package fs

import "math"
import "strings"

func IsNan32(f float32) bool {
	return f != f
}

func IsInf32(f float32) bool {
	return f < -math.MaxFloat32 || f > math.MaxFloat32
}

// Convert a null-terminated byte array to a native Go string
func cstr(str []byte) string {
	n := 0
	for ; str[n] != 0; n++ {
	}
	if n == 0 {
		return ""
	}
	return string(str[:n-1])
}

// Convert a space padded byte array to a native Go string
func fsstr(s []byte) string {
	return strings.TrimSpace(string(s))
}
