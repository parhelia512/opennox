package opennox

import "github.com/gotranspile/cxgo/runtime/libc"

var nox_loaded_thing_bin *nox_memfile = nil

func nox_open_thing_bin() *nox_memfile {
	var things *nox_memfile = nox_loaded_thing_bin
	if things == nil {
		things = nox_memfile_load(CString("thing.bin"), 7)
	}
	return things
}
func nox_ensure_thing_bin() bool {
	nox_loaded_thing_bin = nox_open_thing_bin()
	return nox_loaded_thing_bin != nil
}
func nox_free_thing_bin() {
	if nox_loaded_thing_bin != nil {
		nox_memfile_free(nox_loaded_thing_bin)
		nox_loaded_thing_bin = nil
	}
}
