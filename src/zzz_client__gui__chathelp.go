package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	noxflags "github.com/noxworld-dev/opennox/v1/common/flags"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"image"
	"unsafe"
)

func nox_xxx_cliShowHelpGui_49C560() *uint32 {
	var (
		v0     int32
		result *uint32
		v2     *uint32
		v3     *libc.WChar
		v4     *libc.WChar
		v5     *libc.WChar
		v6     *byte
		v7     *byte
		v8     *byte
	)
	v0 = strMan.Lang()
	if nox_xxx_guiFontHeightMB_43F320(nil) > 10 {
		v0 = 2
	}
	result = (*uint32)(unsafe.Pointer(newWindowFromFile(*(**byte)(memmap.PtrOff(0x587000, uint32(v0*4)+0x282A0)), unsafe.Pointer(cgoFuncAddr(nox_xxx_wnd_49C760)))))
	dword_5d4594_1305680 = uint32(uintptr(unsafe.Pointer(result)))
	if result != nil {
		sub_46B120((*nox_window)(unsafe.Pointer(result)), nil)
		nox_xxx_wndShowModalMB((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1305680))))))
		sub_46C690((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1305680))))))
		guiFocus((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1305680))))))
		(*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1305680)))).SetPos(image.Pt((nox_win_width-*(*int32)(unsafe.Pointer(uintptr(dword_5d4594_1305680 + 8))))/2, (nox_win_height-*(*int32)(unsafe.Pointer(uintptr(dword_5d4594_1305680 + 12))))/2))
		if noxflags.HasGame(noxflags.GameHost) {
			v2 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1305680)))).ChildByID(4102)))
			v6 = (*byte)(unsafe.Pointer(sub_42E8E0(45, 1)))
			v3 = strMan.GetStringInFile("Sanchlp.wnd:Help", "C:\\NoxPost\\src\\client\\Gui\\chathelp.c")
			nox_swprintf((*libc.WChar)(memmap.PtrOff(6112660, 1304656)), v3, v6)
			nox_wcscat((*libc.WChar)(memmap.PtrOff(6112660, 1304656)), libc.CWString(" "))
			v7 = (*byte)(unsafe.Pointer(sub_42E8E0(8, 1)))
			v4 = strMan.GetStringInFile("cdecode.c:KeyToChat", "C:\\NoxPost\\src\\client\\Gui\\chathelp.c")
		} else {
			v2 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1305680)))).ChildByID(4102)))
			v8 = (*byte)(unsafe.Pointer(sub_42E8E0(45, 1)))
			v5 = strMan.GetStringInFile("Sanchlp.wnd:ClientHelp", "C:\\NoxPost\\src\\client\\Gui\\chathelp.c")
			nox_swprintf((*libc.WChar)(memmap.PtrOff(6112660, 1304656)), v5, v8)
			nox_wcscat((*libc.WChar)(memmap.PtrOff(6112660, 1304656)), libc.CWString(" "))
			v7 = (*byte)(unsafe.Pointer(sub_42E8E0(8, 1)))
			v4 = strMan.GetStringInFile("cdecode.c:KeyToChat", "C:\\NoxPost\\src\\client\\Gui\\chathelp.c")
		}
		nox_swprintf((*libc.WChar)(memmap.PtrOff(6112660, 1304400)), v4, v7)
		nox_wcscat((*libc.WChar)(memmap.PtrOff(6112660, 1304656)), (*libc.WChar)(memmap.PtrOff(6112660, 1304400)))
		(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v2)))))).Func94(asWindowEvent(0x4001, int32(uintptr(memmap.PtrOff(6112660, 1304656))), 0))
		if noxflags.HasGame(noxflags.GameHost) {
			if sub_459DA0() == 0 {
				nox_xxx_guiServerOptsLoad_457500()
			}
			sub_459D80(1)
		}
		if noxflags.HasGame(noxflags.GameModeQuest) || (func() *uint32 {
			result = (*uint32)(unsafe.Pointer(uintptr(nox_xxx_isQuest_4D6F50())))
			return result
		}()) != nil {
			result = (*uint32)(unsafe.Pointer(uintptr(sub_49C7A0())))
		}
	}
	return result
}
