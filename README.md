# Go VLBI Field System library
[![](https://godoc.org/github.com/nathany/looper?status.svg)](http://godoc.org/github.com/nvi-inc/fsgo)

This library implements pure Go access to the VLBI Field System shared memory and semaphores. 

Some high level interfaces are provided for supported all FS versions, but access for a particular version can be done via `reflect`. 

Write access is not tested and while should work, will require care and understanding of the FS idioms.

FS scheduling and class buffers not implemented yet.

This includes a modified version of `cgo` for generating the interface. This code is covered by its original license.
