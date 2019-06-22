// Copyright 2019 NVI Inc. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package versions

type Rdbe interface {
	RdbeNum() int
	RdbeMap(int) (map[string]interface{}, error)
	RdbeUpdatedFn(int) (func() bool, error)
}
