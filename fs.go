//go:generate bash ./types.sh
//go:generate bash -c "cgo -godefs types.go | sed 's/Pad_cgo/pad_cgo/' | gofmt > types_$GOARCH.go"

package fs

/*
#cgo LDFLAGS: /usr2/fs/clib/clib.a /usr2/fs/rtelb/rtelb.a

#include <stdio.h>
#include <sys/types.h>
#include <sys/ipc.h>
#include <signal.h>
#include <unistd.h>
#include <stdlib.h>

#include "/usr2/fs/include/ipckeys.h"
#include "/usr2/fs/include/params.h"
#include "/usr2/fs/include/fs_types.h"
#include "/usr2/fs/include/fscom.h"


extern struct fscom *shm_addr;

void setup_ids();
int nsem_test(const char*);

// Don't use setup_ids as it turns off signals

void setup(){
    void sem_att(), skd_att(), shm_att(), cls_att(), brk_att();
    shm_att(SHM_KEY);
    cls_att(CLS_KEY);
    skd_att(SKD_KEY);
    sem_att(SEM_KEY);
    nsem_att(NSEM_KEY);
    brk_att(BRK_KEY);
    go_att(GO_KEY);
};


struct fscom *getshm() {
	if (nsem_test("fs   ") != 1) {
		return NULL;
	}
	return shm_addr;
}
*/
import "C"

import (
	"fmt"
	"strings"
	"unsafe"
)

func GetShm() *fscom {
	return (*fscom)(unsafe.Pointer(C.getshm()))
}

func SemIsLocked(key string) bool {
	ckey := C.CString(fmt.Sprintf("%-5s", key))
	return (C.nsem_test(ckey) >= 1)
}

func Str(str []int8) string {
	return strings.TrimSpace(
		C.GoStringN(
			(*C.char)(unsafe.Pointer(&str[0])),
			(C.int)(len(str))))
}

func init() {
	C.setup()
}
