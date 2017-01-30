
export FSPATH=/home/dhorsley/src/fs

.PHONY: version
version: $(FSPATH)/Makefile
	cp -r template fs$(FS_VERSION)
	echo "package fs" > fs$(FS_VERSION)/version.go
	echo >> fs$(FS_VERSION)/version.go
	echo "const  FieldSystemVersion = \"$(FS_VERSION)\"" >> fs$(FS_VERSION)/version.go
	cd fs$(FS_VERSION); go generate
	cd fs$(FS_VERSION); export GOARCH=386; go generate

include $(FSPATH)/Makefile
