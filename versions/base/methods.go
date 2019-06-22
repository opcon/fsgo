// Copyright 2019 NVI Inc. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package fs

import (
	"errors"
	"strings"

	"fs/svipc"
)

func (f *FieldSystem) Log() string {
	return fsstr(f.Fscom.LLOG[:])

}
func (f *FieldSystem) Schedule() string {
	return fsstr(f.Fscom.LSKD[:])
}
func (f *FieldSystem) Source() string {
	return fsstr(f.Fscom.Lsorna[:])
}
func (f *FieldSystem) SemLocked(semname string) (locked bool, err error) {
	// FS stores a list of names for semephores in the NSEM group. The list is in
	// list in f.Fscom.Sem.  This function queies the semephones in that group by name.
	key, err := svipc.Ftok(NSEM_PATH, NSEM_ID)
	if err != nil {
		return
	}

	sid, err := svipc.Semget(key, 0, 0)
	if err != nil {
		return
	}

	semnum := -1
	semname = strings.TrimSpace(semname)
	for i := 0; i < int(f.Fscom.Sem.Allocated); i++ {
		s := strings.TrimSpace(string(f.Fscom.Sem.Name[i][:]))
		if s == semname {
			semnum = i
			break
		}
	}
	if semnum == -1 {
		err = errors.New("sem not found")
		return
	}

	ret, err := svipc.Semctl(sid, semnum, svipc.GETVAL)
	if err != nil {
		return
	}
	// FS uses sems in a strange way
	locked = (int(ret) == 0)
	return
}

func (f *FieldSystem) Semaphores() []string {
	sems := make([]string, 0, f.Fscom.Sem.Allocated)
	for i := 0; i < int(f.Fscom.Sem.Allocated); i++ {
		sems = append(sems, fsstr(f.Fscom.Sem.Name[i][:]))
	}
	return sems
}

func (f *FieldSystem) Tracking() bool {
	return f.Fscom.Ionsor == 1
}
func (f *FieldSystem) DataValid() bool {
	return f.Fscom.DataValid[0].UserDv == 1
}
