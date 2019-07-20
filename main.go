// Copyright 2019 NVI Inc. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package fs

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"

	versions "github.com/nvi-inc/fsgo/versions"
	_ "github.com/nvi-inc/fsgo/versions/all"
)

const basepath = "/usr2/fs"

type FieldSystem versions.FieldSystem
type Rdbe versions.Rdbe

// NewFieldSystemVersion opens the Field System shared memory using the memory layout of the verion specified, and
// returns an error if the specified version is not supported or the system call fails.
func NewFieldSystemVersion(version string) (fs FieldSystem, err error) {
	creator, ok := versions.Creators[version]
	if !ok {
		return nil, fmt.Errorf("error: version %s not supported", version)
	}
	fs = creator()
	err = fs.Attach()
	return fs, err
}

// NewFieldSystem opens the Field System shared memory by attempting to auto detect the installed version, and returns
// an error if the detected version is not supported or the system call fails.
func NewFieldSystem() (fs FieldSystem, err error) {
	version, err := InstalledVersion()
	if err != nil {
		return nil, err
	}

	creator, ok := versions.Creators[version]
	if !ok {
		return nil, fmt.Errorf("error: version %s not supported", version)
	}
	fs = creator()
	err = fs.Attach()
	return fs, err
}

// SupportedVersions list the versions of the Field System for which this library contains the shared memory layout
func SupportedVersions() []string {
	v := make([]string, 0, len(versions.Creators))
	for k := range versions.Creators {
		v = append(v, k)
	}
	return v
}

// InstalledVersion attempts to detect the installed version of the Field System by parsing
// the Makefile in "/usr2/fs" directory.
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
