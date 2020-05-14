FSPATH=/usr2/fs
export FSPATH

# Needed to override included Makefile
default: version all/all.go

define ALLGOHEADER
// Copyright 2019 NVI Inc. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package all
endef
export ALLGOHEADER

all/all.go: $(shell ls) .
	echo "$$ALLGOHEADER" > $@
	for d in fs[0-9]*; do \
		echo "import _ \"github.com/nvi-inc/fsgo/versions/$$d\"" >>  all/all.go; \
	done
	gofmt -w $@

# Get version number
FS_VERSION := $(shell git --git-dir=$(FSPATH)/.git describe --always --tags 2>/dev/null)

files = consts.go main.go methods.go types_386.go types_amd64.go types.go utils.go version.go rdbe.go

version: $(addprefix fs$(FS_VERSION)/, $(files))

fs$(FS_VERSION)/%.go: base/%.go
	mkdir -p fs$(FS_VERSION)
	ln -sf ../base/$(notdir $@) $@

fs$(FS_VERSION)/version.go: $(FSPATH)/Makefile
	mkdir -p fs$(FS_VERSION)
	/bin/echo -n -e "package fs\nconst  FieldSystemVersion = \"$(FS_VERSION)\""  > $@

fs$(FS_VERSION)/types.go: ./utils/types.sh $(FSPATH)/include/fscom.h
	mkdir -p fs$(FS_VERSION)
	./utils/types.sh > $@

fs$(FS_VERSION)/types_%.go: fs$(FS_VERSION)/types.go utils/cgo/cgo Makefile
	mkdir -p fs$(FS_VERSION)
	export GOARCH=$(basename $*);\
		./utils/cgo/cgo -godefs $< |\
		sed -e 's/PadCgo/padCgo/' -e 's/uint8/byte/' -e 's/int8/byte/' |\
		gofmt > $@
	rm -r _obj

utils/cgo/cgo:
	cd utils/cgo; go build
