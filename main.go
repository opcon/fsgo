// Copyright 2019 NVI Inc. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package fs

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	versions "github.com/nvi-inc/fsgo/versions"
	_ "github.com/nvi-inc/fsgo/versions/all"
)

const DefaultPath = "/usr2/fs"

type FieldSystem versions.FieldSystem
type Rdbe versions.Rdbe

// NewFieldSystemVersion opens the Field System shared memory using the memory
// layout of the verion specified, and returns an error if the specified
// version is not supported or the system call fails.
func NewFieldSystemVersion(version string) (fs FieldSystem, err error) {
	creator, ok := versions.Creators[version]
	if !ok {
		return nil, fmt.Errorf("error: version %s not supported", version)
	}
	fs = creator()
	err = fs.Attach()
	return fs, err
}

// NewFieldSystem opens the Field System shared memory by attempting to auto
// detect the installed version, and returns an error if the detected version
// is not supported or the system call fails.
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

// SupportedVersions list the versions of the Field System for which this
// library contains the shared memory layout
func SupportedVersions() []string {
	v := make([]string, 0, len(versions.Creators))
	for k := range versions.Creators {
		v = append(v, k)
	}
	return v
}

// InstalledVersion attempts to detect the installed version of the Field System by
// checking:
// -   if path is a git repository, then using the tag.
// -   if path is a symlink to /usr2/fs-(version) and using (version)
// -   if path/Makefile contains the variables VERSION SUBLEVEL PATCHLEVEL
func InstalledVersion() (string, error) {

	return "", fmt.Errorf("")
}

var ErrNotGitDir = errors.New("not a git directory")

// InstalledVersionFromGit attemps to detect the FS version installed in path
// by using the git tags.  This requires git to be in the path.
func InstalledVersionFromGit(path string) (string, error) {
	var out bytes.Buffer

	if _, err := os.Stat(path); err != nil {
		return "", err
	}

	if _, err := os.Stat(path + "/.git"); os.IsNotExist(err) {
		return "", ErrNotGitDir
	}

	gitargs := []string{
		fmt.Sprintf("--git-dir=%s/.git", path),
		"describe",
		"--always",
		"--tags",
	}

	cmd := exec.Command("git", gitargs...)
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("calling git: %w", err)
	}

	return strings.TrimSpace(out.String()), nil
}

var pathRegex = regexp.MustCompile(`fs-(\d+\.\d+\.\d+)$`)
var ErrPathMatch = errors.New("not a git directory")

// InstalledVersionFromPath attemps to detect the version of the installed
// Field System by examining the sympolic link "path"
func InstalledVersionFromPath(path string) (string, error) {
	name, err := os.Readlink(path)
	if err != nil {
		return "", err
	}
	m := pathRegex.FindStringSubmatch(filepath.Base(name))
	if m == nil || len(m) == 0 {
		return "", ErrPathMatch
	}

	return m[1], nil
}

// InstalledVersionFromMakefile attempts to detect the installed version of the Field System by parsing
// the Makefile in "/usr2/fs" directory.
func InstalledVersionFromMakefile(path string) (string, error) {
	r, err := regexp.Compile(`^(\w+)\s*=\s*(\w+)$`)
	if err != nil {
		return "", err
	}

	vars := make(map[string]int)

	f, err := os.Open(path + "/Makefile")
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
	// TODO: check if these are digits
	return fmt.Sprintf("%d.%d.%d", vars["VERSION"], vars["SUBLEVEL"], vars["PATCHLEVEL"]), nil
}
