package versions

import (
	"bufio"
	"fmt"
	"fsversions/versions"
	"os"
	"regexp"
	"strconv"
)

const (
	basepath = "/usr2/fs"
)

type Creator func() (FieldSystem, err)

var Creators = map[string]Creator{}

func Add(name string, creator Creator) {
	Creators[name] = creator
}

func GetInstalled() (Creator, error) {
	vers, err := InstalledVersion()
	if err != nil {
		return nil, err
	}
	creator, ok := versions.Creators[vers]
	if !ok {
		return nil, fmt.Errorf("error: version %s not supported", vers)
	}
	return creator, nil

}

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
