package fs

import (
	"errors"
	"strings"
	"unsafe"

	"github.com/fatih/structs"

	"fs/svipc"
)

func AttachFS() (fs *Fscom, err error) {
	key, err := svipc.Ftok(SHM_PATH, SHM_ID)
	if err != nil {
		return
	}
	id, err := svipc.Shmget(key, unsafe.Sizeof(&Fscom{}), 0666)
	if err != nil {
		return
	}
	ptr, err := svipc.Shmat(id, 0, svipc.SHM_RDONLY)
	if err != nil {
		return
	}
	fs = (*Fscom)(ptr)
	return
}

func (_ *Fscom) Version() string {
	return FieldSystemVersion
}

func (fs *Fscom) Log() string {
	return fsstr(fs.LLOG[:])

}
func (fs *Fscom) Schedule() string {
	return fsstr(fs.LSKD[:])
}
func (fs *Fscom) Source() string {
	return fsstr(fs.Lsorna[:])
}
func (fs *Fscom) SemLocked(semname string) (locked bool, err error) {
	// FS stores a list of names for semephores in the NSEM group. The list is in
	// list in fs.Sem.  This function queies the semephones in that group by name.
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
	for i := 0; i < int(fs.Sem.Allocated); i++ {
		s := strings.TrimSpace(string(fs.Sem.Name[i][:]))
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

func (fs *Fscom) Semaphores() []string {
	sems := make([]string, 0, fs.Sem.Allocated)
	for i := 0; i < int(fs.Sem.Allocated); i++ {
		sems = append(sems, fsstr(fs.Sem.Name[i][:]))
	}
	return sems
}

func (fs *Fscom) Tracking() bool {
	return fs.Ionsor == 1
}
func (fs *Fscom) DataValid() bool {
	return fs.DataValid[0].UserDv == 1
}

func (fs *Fscom) Dump() map[string]interface{} {
	return structs.Map(fs)
}
