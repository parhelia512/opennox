package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"unsafe"
)

func nox_xxx_guiChatIconLoad_445650() int32 {
	var v0 *uint16
	*memmap.PtrUint32(6112660, 825748) = uint32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg("ChatIcon"))))
	dword_5d4594_825744 = uint32(uintptr(unsafe.Pointer(nox_window_new(nil, 136, nox_win_width-50, nox_win_height/2-50, 50, 50, nil))))
	nox_xxx_wndSetIcon_46AE60(*(*int32)(unsafe.Pointer(&dword_5d4594_825744)), *memmap.PtrInt32(6112660, 825748))
	(*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_825744)))).SetAllFuncs(nil, func(arg1 *nox_window, arg2 unsafe.Pointer) int32 {
		return nox_xxx_guiChatMode_4456E0(&arg1.id)
	}, nil)
	v0 = (*uint16)(unsafe.Pointer(strMan.GetStringInFile("chatmode", "C:\\NoxPost\\src\\client\\Gui\\chaticon.c")))
	nox_xxx_wndWddSetTooltip_46B000((*nox_window_data)(unsafe.Pointer(uintptr(dword_5d4594_825744+36))), (*libc.WChar)(unsafe.Pointer(v0)))
	return 1
}
