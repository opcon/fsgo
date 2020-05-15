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

var ErrVersionNotSupported = errors.New("version not supported")

// Attach attaches to the shared memory of the Field System installed at
// DefaultPath, attempting to automatically detect the version. Returns an
// error if the detected version is not supported or the system call fails.
func Attach() (fs FieldSystem, err error) {
	return AttachPath(DefaultPath)
}

// AttachPath attaches to the shared memory of the Field System installed at
// path, attempting to automatically detect the version. Returns an error if
// the detected version is not supported or the system call fails.
func AttachPath(path string) (fs FieldSystem, err error) {
	version, err := InstalledVersion(path)
	if err != nil {
		return nil, err
	}
	return AttachVersionPath(version, path)
}

// AttachPath attaches to the shared memory of the Field System installed at
// DefaultPath, using the shared memory version specified. Returns an error if
// the version is not supported or the system call fails.
func AttachVersion(version string) (fs FieldSystem, err error) {
	return AttachVersionPath(version, DefaultPath)
}

// AttachVersionPath attaches to the shared memory of the Field System installed at
// path, using the shared memory version specified. Returns an error if the
// version is not supported or the system call fails.
func AttachVersionPath(version, path string) (fs FieldSystem, err error) {
	creator, ok := versions.Creators[version]
	if !ok {
		return nil, fmt.Errorf("load version %s: %w", version, ErrVersionNotSupported)
	}
	fs = creator()
	err = fs.Attach()
	if err != nil {
		return nil, fmt.Errorf("attaching shared memory: %w", err)
	}
	return fs, nil
}

// SupportedVersions list the versions of the Field System supported
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
func InstalledVersion(path string) (string, error) {
	version, errgit := InstalledVersionFromGit(path)
	if errgit == nil {
		return version, nil
	}
	if os.IsNotExist(errgit) {
		return "", errgit
	}

	version, errpath := InstalledVersionFromPath(path)
	if errpath == nil {
		return version, nil
	}
	version, errmake := InstalledVersionFromMakefile(path)
	if errmake == nil {
		return version, nil
	}

	return "",
		fmt.Errorf("git method: %v, path method: %v, makefile method: %v",
			errgit, errpath, errmake)
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

// InstalledVersionFromMakefile attempts to detect the installed version of the
// Field System by parsing the Makefile in "/usr2/fs" directory.
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
