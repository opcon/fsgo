export FSPATH=/usr2/fs

default: version

# Get version number
include $(FSPATH)/Makefile

vgos: version fs$(FS_VERSION)/rdbe.go

files = consts.go main.go methods.go types_386.go types_amd64.go types.go utils.go version.go
# Excludes rdbe.go by default since 9.11.* doesn't have the structures

version: $(addprefix fs$(FS_VERSION)/, $(files))

fs$(FS_VERSION)/%.go: base/%.go
	mkdir -p fs$(FS_VERSION)
	ln -sf ../base/$(notdir $@) $@

fs$(FS_VERSION)/version.go: $(FSPATH)/Makefile
	mkdir -p fs$(FS_VERSION)
	/bin/echo -n -e "package fs\nconst  FieldSystemVersion = \"$(FS_VERSION)\""  > $@

fs$(FS_VERSION)/types.go: ./utils/types.sh 
	mkdir -p fs$(FS_VERSION)
	./utils/types.sh > $@

fs$(FS_VERSION)/types_%.go: fs$(FS_VERSION)/types.go utils/cgo/cgo 
	mkdir -p fs$(FS_VERSION)
	export GOARCH=$(basename $*);\
		./utils/cgo/cgo -godefs $< |\
		sed -e 's/Pad_cgo/pad_cgo/' -e 's/uint8/byte/' -e 's/int8/byte/' |\
		gofmt > $@
	rm -r _obj
