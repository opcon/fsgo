package fs_test

import (
	"io/ioutil"
	"os"
	"os/exec"
	"testing"

	fs "github.com/nvi-inc/fsgo"
)

func TestInstalledVersionFromMakefile(t *testing.T) {
	const testVersion = "9.11.19"
	const makefile = `VERSION = 9
SUBLEVEL = 11
PATCHLEVEL = 19
FS_VERSION = $(VERSION).$(SUBLEVEL).$(PATCHLEVEL)
export VERSION SUBLEVEL PATCHLEVEL FS_VERSION
	`
	must := func(s string, err error) {
		if err != nil {
			t.Fatal(s, err)
		}
	}

	root, err := ioutil.TempDir("", "fstest*")
	must("setup test root", err)

	t.Cleanup(func() {
		os.RemoveAll(root)
	})

	path := root + "/fs"
	must("setup dir", os.Mkdir(path, 0o700))
	must("writing test file", ioutil.WriteFile(path+"/Makefile", []byte(makefile), 0o600))

	version, err := fs.InstalledVersionFromMakefile(path)
	if err != nil {
		t.Error("unexpected error", err)
	}
	if version != testVersion {
		t.Errorf("expected version %q, got version %q", testVersion, version)
	}
}

func TestInstalledVersionFromPath(t *testing.T) {
	const testVersion = "10.0.0"
	must := func(s string, err error) {
		if err != nil {
			t.Fatal(s, err)
		}
	}

	root, err := ioutil.TempDir("", "fstest*")
	must("setup test root", err)

	t.Cleanup(func() {
		os.RemoveAll(root)
	})

	path := root + "/fs"

	t.Run("not a symbolic link", func(t *testing.T) {
		must("setup dir", os.Mkdir(path, 0o600))
		_, err := fs.InstalledVersionFromPath(path)
		if err == nil {
			t.Error("expected an error")
		}
		must("remove dir", os.Remove(path))
	})

	t.Run("reject extra junk", func(t *testing.T) {
		realpath := root + "/fs-" + testVersion + "-test"
		must("setup test path", os.Mkdir(realpath, 0o600))
		must("setup symlink", os.Symlink(realpath, path))

		version, err := fs.InstalledVersionFromPath(path)
		if err == nil {
			t.Errorf("expected an error instead got %q", version)
		}
		must("cleanup realpath", os.Remove(realpath))
		must("cleanup path", os.Remove(path))
	})

	t.Run("detect good version", func(t *testing.T) {
		realpath := root + "/fs-" + testVersion
		must("setup test path", os.Mkdir(realpath, 0o600))
		must("setup symlink", os.Symlink(realpath, path))

		version, err := fs.InstalledVersionFromPath(path)
		if err != nil {
			t.Error("unexpected error", err)
		}
		if version != testVersion {
			t.Errorf("expected version %q, got version %q", testVersion, version)
		}
	})

}

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

func TestSupportedVersions(t *testing.T) {
	vs := fs.SupportedVersions()
	if len(vs) == 0 {
		t.Errorf("no supported versions")
	}
}

func TestAttach(t *testing.T) {
	// Need to be run with the FS installed and setup
	_, err := fs.Attach()
	if err != nil {
		t.Fatal(err)
	}

}
