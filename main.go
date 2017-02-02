package fs

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"

	versions "fs/versions"
	_ "fs/versions/all"
)

const basepath = "/usr2/fs"

type FieldSystem versions.FieldSystem
type Rdbe versions.Rdbe

func NewFieldSystem() (FieldSystem, error) {
	vers, err := InstalledVersion()
	if err != nil {
		return nil, err
	}
	creator, ok := versions.Creators[vers]
	if !ok {
		return nil, fmt.Errorf("error: version %s not supported", vers)
	}
	fs := creator()
	err = fs.Attach()

	return fs, err

}

func SupportedVersions() []string {
	v := make([]string, 0, len(versions.Creators))
	for k := range versions.Creators {
		v = append(v, k)
	}
	return v
}

// Parses Makefile in fs directory to find the installed FS version
func InstalledVersion() (string, error) {
	r, err := regexp.Compile(`^(\w+)\s*=\s*(\w+)$`)
	if err != nil {
		return "", err
	}

	vars := make(map[string]int)

	f, err := os.Open(basepath + "/Makefile")
	if err != nil {
		return "", err
	}

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		m := r.FindStringSubmatch(line)
		if m == nil {
			continue
		}
		i, err := strconv.Atoi(m[2])
		if err != nil {
			continue
		}
		vars[m[1]] = i
	}
	return fmt.Sprintf("%d.%d.%d", vars["VERSION"], vars["SUBLEVEL"], vars["PATCHLEVEL"]), nil
}
