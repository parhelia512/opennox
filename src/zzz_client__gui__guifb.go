package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"unsafe"
)

func sub_456140(a1 uint8) {
	var (
		v2 int32
		v3 *libc.WChar
		v4 *libc.WChar
		v5 *libc.WChar
		v6 *libc.WChar
	)
	*memmap.PtrUint8(6112660, 1045644) = a1
	v2 = int32(dword_5d4594_1045636 + 36)
	switch a1 {
	case 0:
		*(*uint32)(unsafe.Pointer(uintptr(v2 + 24))) = uint32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg("BallAtHome"))))
		v3 = strMan.GetStringInFile("BallHomeTT", "C:\\NoxPost\\src\\client\\Gui\\guifb.c")
		nox_xxx_wndWddSetTooltip_46B000((*nox_window_data)(unsafe.Pointer(uintptr(v2))), v3)
	case 1:
		*(*uint32)(unsafe.Pointer(uintptr(v2 + 24))) = uint32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg("BallAway"))))
		v4 = strMan.GetStringInFile("BallAwayTT", "C:\\NoxPost\\src\\client\\Gui\\guifb.c")
		nox_xxx_wndWddSetTooltip_46B000((*nox_window_data)(unsafe.Pointer(uintptr(v2))), v4)
	case 2:
		*(*uint32)(unsafe.Pointer(uintptr(v2 + 24))) = uint32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg("BallRed"))))
		v5 = strMan.GetStringInFile("BallRedTT", "C:\\NoxPost\\src\\client\\Gui\\guifb.c")
		nox_xxx_wndWddSetTooltip_46B000((*nox_window_data)(unsafe.Pointer(uintptr(v2))), v5)
	case 4:
		*(*uint32)(unsafe.Pointer(uintptr(v2 + 24))) = uint32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg("BallBlue"))))
		v6 = strMan.GetStringInFile("BallBlueTT", "C:\\NoxPost\\src\\client\\Gui\\guifb.c")
		nox_xxx_wndWddSetTooltip_46B000((*nox_window_data)(unsafe.Pointer(uintptr(v2))), v6)
	}
}
