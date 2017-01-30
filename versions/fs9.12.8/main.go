//go:generate bash ./types.sh
//go:generate bash -c "cgo -godefs types.go | sed -e 's/Pad_cgo/pad_cgo/' -e 's/uint8/byte/' -e 's/int8/byte/' | gofmt > types_$GOARCH.go"
package fs

import (
	"fs/versions"
)

func init() {
	versions.Add(FieldSystemVersion, func() (versions.FieldSystem, error) {
		return AttachFS()
	})
}
