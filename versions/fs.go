// Copyright 2019 NVI Inc. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package versions

type FieldSystem interface {
	Attach() error
	Log() string
	Schedule() string
	Source() string
	SemLocked(string) (bool, error)
	Semaphores() []string
	Tracking() bool
	DataValid() bool
}
