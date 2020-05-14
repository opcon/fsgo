package fs_test

import (
	"io/ioutil"
	"os"
	"os/exec"
	"testing"

	fs "github.com/nvi-inc/fsgo"
)

func TestInstalledVersionFromGit(t *testing.T) {
	t.Run("nonexistant path", func(t *testing.T) {
		_, err := fs.InstalledVersionFromGit("/thispathshouldnotexit")
		if !os.IsNotExist(err) {
			t.Error("unexpected error", err)
		}
	})

	must := func(s string, err error) {
		if err != nil {
			t.Fatal(s, err)
		}
	}

	path, err := ioutil.TempDir("", "fstest*")
	t.Cleanup(func() {
		os.RemoveAll(path)
	})

	const tagVersion = "10.0.0"

	must("setting up testdir", err)

	t.Run("not a git repo", func(t *testing.T) {
		_, err := fs.InstalledVersionFromGit(path)
		if err != fs.ErrNotGitDir {
			t.Error("unexpected error", err)
		}
	})

	must("git init", exec.Command("git", "init", path).Run())
	git := func(args ...string) {
		args = append([]string{
			"--git-dir=" + path + "/.git",
			"--work-tree=" + path,
		}, args...)
		if err := exec.Command("git", args...).Run(); err != nil {
			t.Fatalf("while calling git with args %v: %v", args, err)
		}
	}

	must("writing test file", ioutil.WriteFile(path+"/test.txt", []byte("test"), 0644))
	git("add", path+"/test.txt")
	git("commit", "-m", "test")
	git("tag", tagVersion)

	version, err := fs.InstalledVersionFromGit(path)

	if err != nil {
		t.Fatal("running InstalledVersionFromGit", err)
	}
	if version != tagVersion {
		t.Errorf("version %q does not match expected %q", version, tagVersion)

	}
}
