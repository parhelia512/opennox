package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"unsafe"
)

func sub_4E39F0_obj_db(a1 **byte) *libc.WChar {
	var (
		v1 *byte
		v2 *byte
		v3 *byte
		v4 *byte
		v5 uint8
		v6 *uint8
		i  *uint8
	)
	v1 = *a1
	if *a1 == nil {
		v1 = (*byte)(unsafe.Pointer(uintptr(nox_xxx_getUnitName_4E39D0(int32(uintptr(unsafe.Pointer(a1)))))))
	}
	v2 = libc.StrRChr(v1, 58)
	if v2 != nil {
		v3 = (*byte)(unsafe.Add(unsafe.Pointer(v2), 1))
	} else {
		v3 = v1
	}
	v4 = v3
	libc.StrCpy((*byte)(memmap.PtrOff(6112660, 1563460)), CString("NPC:"))
	v5 = uint8(*v3)
	v6 = (*uint8)(memmap.PtrOff(6112660, uint32(libc.StrLen((*byte)(memmap.PtrOff(6112660, 1563460)))+1563460)))
	for i = (*uint8)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v4), 1)))); int32(v5) != 0; i = (*uint8)(unsafe.Add(unsafe.Pointer(i), 1)) {
		if int32(v5) != 95 {
			*func() *uint8 {
				p := &v6
				x := *p
				*p = (*uint8)(unsafe.Add(unsafe.Pointer(*p), 1))
				return x
			}() = v5
		}
		v5 = *i
	}
	*v6 = 0
	return strMan.GetStringInFile((*byte)(memmap.PtrOff(6112660, 1563460)), "C:\\NoxPost\\src\\Server\\DBase\\objdb.c")
}
