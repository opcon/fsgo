//go:generate bash ./types.sh
//go:generate bash -c "cgo -godefs types.go > types_$GOARCH.go"

package fs

/*
#cgo LDFLAGS: /usr2/fs/clib/clib.a /usr2/fs/rtelb/rtelb.a

#include <stdlib.h>
extern struct fscom *shm_addr;

void setup_ids();
int nsem_test(const char*);

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
	C.setup_ids()
}
