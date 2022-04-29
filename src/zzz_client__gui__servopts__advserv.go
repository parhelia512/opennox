package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"unsafe"
)

func sub_4BE2C0(a1 int32) int32 {
	var (
		v1 *libc.WChar
		v2 *uint32
	)
	*(*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(sub_416640())) + 74))) = uint32(a1)
	v1 = strMan.GetStringInFile("AudCullDesc", "C:\\NoxPost\\src\\client\\Gui\\ServOpts\\advserv.c")
	nox_swprintf((*libc.WChar)(memmap.PtrOff(6112660, 1316716)), v1, a1)
	v2 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1316972)))).ChildByID(2120)))
	return (*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v2)))))).Func94(asWindowEvent(0x4001, int32(uintptr(memmap.PtrOff(6112660, 1316716))), -1))
}
