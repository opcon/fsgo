package versions

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func Attach() (FieldSystem, error) {
	vers, err := InstalledVersion()
	if err != nil {
		return nil, err
	}
	creator, ok := Creators[vers]
	if !ok {
		return nil, fmt.Errorf("error: version %s not supported", vers)
	}
	fs, err := creator()
	return fs, err

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
