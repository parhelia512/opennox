package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"unsafe"
)

func nox_xxx_printCentered_445490(a1 *libc.WChar) {
	var (
		v1 int32
		v2 int32
		v3 *uint16
	)
	if a1 != nil {
		v1 = int32(func() uint32 {
			p := &dword_5d4594_825736
			*p++
			return *p
		}())
		if dword_5d4594_825736 == 3 {
			v1 = 0
			dword_5d4594_825736 = 0
		}
		nox_wcscpy((*libc.WChar)(unsafe.Pointer(memmap.PtrUint16(6112660, uint32(v1*644)+0xC91FC))), a1)
		v2 = int32(dword_5d4594_825736 * 644)
		*memmap.PtrUint32(6112660, uint32(v2)+824440) = nox_frame_xxx_2598000 + nox_gameFPS*4 + nox_gameFPS
		*memmap.PtrUint8(6112660, uint32(v2)+0xC947C) = 0
		v3 = (*uint16)(unsafe.Pointer(strMan.GetStringInFile("systemmsg", "C:\\NoxPost\\src\\Client\\Gui\\guimsg.c")))
		nox_gui_console_Printf_450C00(uint8(int8(NOX_CONSOLE_RED)), (*libc.WChar)(unsafe.Pointer(v3)), a1)
	}
}
