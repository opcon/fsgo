package fs

import (
	"fs/versions"
)

func init() {
	versions.Add(FieldSystemVersion, func() versions.FieldSystem {
		fs, err := AttachFS()
		if err != nil {
			panic(err)
		}
		return fs
	})
}
