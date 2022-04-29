package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"unsafe"
)

func sub_455C30() int32 {
	var (
		result int32
		v1     int32
		v2     *libc.WChar
		v3     *libc.WChar
	)
	if dword_5d4594_1045604 != 0 {
		return 1
	}
	result = int32(uintptr(unsafe.Pointer(newWindowFromFile("GUI_CTF.wnd", nil))))
	dword_5d4594_1045604 = uint32(result)
	if result != 0 {
		v1 = 8811
		for {
			v2 = (*libc.WChar)(unsafe.Pointer((*nox_window)(unsafe.Pointer(uintptr(result))).ChildByID(v1)))
			(*nox_window)(unsafe.Pointer(v2)).SetAllFuncs(nil, func(arg1 *nox_window, arg2 unsafe.Pointer) int32 {
				return sub_455CD0((*uint8)(unsafe.Pointer(arg1)), (*uint32)(arg2))
			}, nil)
			v3 = strMan.GetStringInFile("FlagHomeTT", "C:\\NoxPost\\src\\client\\Gui\\GUI_CTF.c")
			nox_xxx_wndWddSetTooltip_46B000((*nox_window_data)(unsafe.Pointer((*libc.WChar)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(libc.WChar(0))*18)))), v3)
			if func() int32 {
				p := &v1
				*p++
				return *p
			}() > 8826 {
				break
			}
			result = int32(dword_5d4594_1045604)
		}
		sub_455A00(0)
		*memmap.PtrUint32(6112660, 1045632) = uint32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg("FlagTeamBorder"))))
		result = 1
	}
	return result
}
func sub_455D80(a1 uint8, a2 int8) {
	var (
		result *libc.WChar
		v3     *libc.WChar
		v4     *libc.WChar
	)
	*memmap.PtrUint8(6112660, uint32(a1)+0xFF46B) = uint8(a2)
	result = (*libc.WChar)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045604)))).ChildByID(int32(a1) + 8810)))
	v3 = result
	if result != nil {
		if *(*libc.WChar)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(libc.WChar(0))*2))&32 != 0 {
			if int32(a2) != 0 {
				if int32(a2) == 1 {
					v4 = strMan.GetStringInFile("YourFlagCarriedTT", "C:\\NoxPost\\src\\client\\Gui\\GUI_CTF.c")
				} else {
					if int32(a2) != 2 {
						return
					}
					v4 = strMan.GetStringInFile("FlagAwayTT", "C:\\NoxPost\\src\\client\\Gui\\GUI_CTF.c")
				}
			} else {
				v4 = strMan.GetStringInFile("FlagHomeTT", "C:\\NoxPost\\src\\client\\Gui\\GUI_CTF.c")
			}
		} else if int32(a2) != 0 {
			if int32(a2) == 1 {
				v4 = strMan.GetStringInFile("TheirFlagCarriedTT", "C:\\NoxPost\\src\\client\\Gui\\GUI_CTF.c")
			} else {
				if int32(a2) != 2 {
					return
				}
				v4 = strMan.GetStringInFile("FlagAwayTT", "C:\\NoxPost\\src\\client\\Gui\\GUI_CTF.c")
			}
		} else {
			v4 = strMan.GetStringInFile("FlagHomeTT", "C:\\NoxPost\\src\\client\\Gui\\GUI_CTF.c")
		}
		nox_xxx_wndWddSetTooltip_46B000((*nox_window_data)(unsafe.Pointer((*libc.WChar)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(libc.WChar(0))*18)))), v4)
	}
}
