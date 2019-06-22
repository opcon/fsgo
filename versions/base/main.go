// Copyright 2019 NVI Inc. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package fs

import (
	"unsafe"

	"github.com/nvi-inc/fs-go/versions"
	"github.com/nvi-inc/svipc"
)

type FieldSystem struct {
	Version string `json:"version"`
	Fscom   *Fscom `json:"shm"`
	// TODO(dh): Allow different installation paths
}

func (f *FieldSystem) Attach() (err error) {
	key, err := svipc.Ftok(SHM_PATH, SHM_ID)
	if err != nil {
		return
	}
	id, err := svipc.Shmget(key, unsafe.Sizeof(&Fscom{}), 0666)
	if err != nil {
		return
	}
	//TODO(dh): test out writing
	ptr, err := svipc.Shmat(id, 0, svipc.SHM_RDONLY)
	if err != nil {
		return
	}
	f.Fscom = (*Fscom)(ptr)
	return
}

func init() {
	versions.Add(FieldSystemVersion, func() versions.FieldSystem {
		return &FieldSystem{Version: FieldSystemVersion}
	})
}
