package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"image"
	"unsafe"
)

var dword_5d4594_1321236 *nox_window = nil
var dword_5d4594_1321240 *nox_window = nil
var dword_5d4594_1321248 *nox_window = nil
var dword_5d4594_1321244 *nox_window = nil

func sub_4C3620() *byte {
	var (
		v0     int32
		i      int32
		v2     *libc.WChar
		v3     int32
		v4     *libc.WChar
		v5     *byte
		v6     *libc.WChar
		v7     *byte
		v8     *libc.WChar
		v9     int32
		v10    *libc.WChar
		result *byte
		v12    [256]byte
	)
	v0 = int32(*(*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer(dword_5d4594_1321240))) + 32))))
	sub_42CD90()
	for i = 0; i < int32(*(*int16)(unsafe.Pointer(uintptr(v0 + 44)))); i++ {
		v2 = (*libc.WChar)(unsafe.Pointer(uintptr((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1321240))))).Func94(asWindowEvent(0x4016, i, 0)))))
		v3 = int32(uintptr(unsafe.Pointer(nox_xxx_bindevent_bindNameByTitle_42EA40(v2))))
		v4 = (*libc.WChar)(unsafe.Pointer(uintptr((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1321244))))).Func94(asWindowEvent(0x4016, i, 0)))))
		v5 = nox_xxx_keybind_nameByTitle_42E960(v4)
		v6 = (*libc.WChar)(unsafe.Pointer(uintptr((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1321248))))).Func94(asWindowEvent(0x4016, i, 0)))))
		v7 = nox_xxx_keybind_nameByTitle_42E960(v6)
		if v7 != nil {
			nox_sprintf(&v12[0], CString("%s = %s"), v7, v3)
			nox_client_parseConfigHotkeysLine_42CF50(&v12[0])
		}
		if v5 != nil {
			nox_sprintf(&v12[0], CString("%s = %s"), v5, v3)
			nox_client_parseConfigHotkeysLine_42CF50(&v12[0])
		}
	}
	v8 = strMan.GetStringInFile("bindevent:ToggleQuitMenu", "C:\\NoxPost\\src\\client\\Gui\\GuiInput.c")
	v9 = int32(uintptr(unsafe.Pointer(nox_xxx_bindevent_bindNameByTitle_42EA40(v8))))
	v10 = strMan.GetStringInFile("keybind:Esc", "C:\\NoxPost\\src\\client\\Gui\\GuiInput.c")
	result = nox_xxx_keybind_nameByTitle_42E960(v10)
	if result != nil {
		nox_sprintf(&v12[0], CString("%s = %s"), result, v9)
		result = (*byte)(unsafe.Pointer(uintptr(nox_client_parseConfigHotkeysLine_42CF50(&v12[0]))))
	}
	return result
}
func sub_4C3760() int32 {
	var (
		result int32
		v1     **uint32
		v2     *uint32
		v3     *uint32
		v4     int32
		v5     int32
		v6     *uint32
		v7     int32
		v8     *uint32
		v9     *uint32
		v10    *uint32
		v11    *libc.WChar
	)
	_ = v11
	result = int32(uintptr(unsafe.Pointer(newWindowFromFile("InputCfg.wnd", unsafe.Pointer(cgoFuncAddr(sub_4C3A90))))))
	dword_5d4594_1321228 = uint32(result)
	if result != 0 {
		dword_5d4594_1321236 = (*nox_window)(unsafe.Pointer(uintptr(result))).ChildByID(910)
		dword_5d4594_1321240 = (*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1321228)))).ChildByID(911)
		dword_5d4594_1321244 = (*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1321228)))).ChildByID(912)
		dword_5d4594_1321248 = (*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1321228)))).ChildByID(913)
		result = int32(uintptr(unsafe.Pointer(dword_5d4594_1321236)))
		if dword_5d4594_1321236 != nil {
			v1 = *(***uint32)(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer(dword_5d4594_1321236))) + 32)))
			**(**uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof((*uint32)(nil))*7)) = 921
			**(**uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof((*uint32)(nil))*8)) = 922
			**(**uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof((*uint32)(nil))*9)) = 920
			(*(*int32)(unsafe.Pointer(&dword_5d4594_1321236))).setFunc94(func(arg1 int32, arg2 int32, arg3 int32, arg4 int32) int32 {
				return sub_4C3CD0(arg1, uint32(arg2), arg3, arg4)
			})
			sub_46B120((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1321240)))), (*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1321236))))))
			sub_46B120((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1321244)))), (*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1321236))))))
			sub_46B120((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1321248)))), (*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1321236))))))
			(*(*int32)(unsafe.Pointer(&dword_5d4594_1321244))).setFunc93(func(arg1 int32, arg2 int32, arg3 int32, arg4 int32) int32 {
				return sub_4C3A60((*uint32)(unsafe.Pointer(uintptr(arg1))), uint32(arg2), uint32(arg3), arg4)
			})
			(*(*int32)(unsafe.Pointer(&dword_5d4594_1321248))).setFunc93(func(arg1 int32, arg2 int32, arg3 int32, arg4 int32) int32 {
				return sub_4C3A60((*uint32)(unsafe.Pointer(uintptr(arg1))), uint32(arg2), uint32(arg3), arg4)
			})
			v2 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1321228)))).ChildByID(921)))
			(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1321240))))).Func94(asWindowEvent(0x4018, int32(uintptr(unsafe.Pointer(v2))), 0))
			(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1321244))))).Func94(asWindowEvent(0x4018, int32(uintptr(unsafe.Pointer(v2))), 0))
			(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1321248))))).Func94(asWindowEvent(0x4018, int32(uintptr(unsafe.Pointer(v2))), 0))
			v3 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1321228)))).ChildByID(922)))
			(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1321240))))).Func94(asWindowEvent(0x4019, int32(uintptr(unsafe.Pointer(v3))), 0))
			(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1321244))))).Func94(asWindowEvent(0x4019, int32(uintptr(unsafe.Pointer(v3))), 0))
			(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1321248))))).Func94(asWindowEvent(0x4019, int32(uintptr(unsafe.Pointer(v3))), 0))
			v4 = 971
			v5 = int32(sub_47DBC0()) + 971
			if v5 > 971 {
				for {
					v6 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1321228)))).ChildByID(v4)))
					nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v6)))))), 1)
					v4++
					if v4 >= v5 {
						break
					}
				}
			}
			v7 = nox_client_mousePriKey_430AF0()
			v8 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1321228)))).ChildByID(v7 + 971)))
			(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v8)))))).Func94(asWindowEvent(0x4008, 1, 0))
			(*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1321228)))).SetPos(image.Pt(int32((uint32(nox_win_width)-*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1321228 + 8))))/2), 0))
			dword_5d4594_1321232 = uint32(uintptr(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1321228)))).ChildByID(980))))
			sub_46B120((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1321232)))), nil)
			(*(*int32)(unsafe.Pointer(&dword_5d4594_1321232))).setFunc94(func(arg1 int32, arg2 int32, arg3 int32, arg4 int32) int32 {
				return sub_4C3A90(arg1, arg2, (*int32)(unsafe.Pointer(uintptr(arg3))), arg4)
			})
			(*(*int32)(unsafe.Pointer(&dword_5d4594_1321232))).setFunc93(func(arg1 int32, arg2 int32, arg3 int32, arg4 int32) int32 {
				return sub_4C3EB0(arg1, arg2, uint32(arg3), arg4)
			})
			(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1321232))))).Hide()
			(*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1321232)))).SetPos(image.Pt(int32((uint32(nox_win_width)-*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1321232 + 8))))/2), int32(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1321232 + 20))))))
			v9 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1321232)))).ChildByID(981)))
			sub_46AEE0(int32(uintptr(unsafe.Pointer(v9))), int32(uintptr(memmap.PtrOff(6112660, 1321256))))
			v10 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1321228)))).ChildByID(932)))
			nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v10)))))), 1)
			(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1321228))))).Hide()
			result = 1
		}
	}
	return result
}
func sub_4C3CD0(a1 int32, a2 uint32, a3 int32, a4 int32) int32 {
	var (
		v5 int32
		v6 *libc.WChar
		v7 int32
		v8 int32
		v9 int32
	)
	if a2 > 0x4007 {
		if a2 == 0x4009 {
			v7 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 32))))
			nox_xxx_wndListboxProcPre_4A30D0((*nox_window)(unsafe.Pointer(uintptr(a1))), 0x4009, uint32(uintptr(unsafe.Pointer((*libc.WChar)(unsafe.Pointer(uintptr(a3)))))), a4)
			v8 = sub_4A4800(v7)
			(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1321240))))).Func94(asWindowEvent(0x401C, v8, 0))
			(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1321244))))).Func94(asWindowEvent(0x401C, v8, 0))
			(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1321248))))).Func94(asWindowEvent(0x401C, v8, 0))
		} else if a2 == 16400 {
			v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(a3 + 32))))
			if int32(*(*uint32)(unsafe.Pointer(uintptr(v5 + 48)))) >= 0 {
				dword_5d4594_1321252 = uint32(a3)
				v9 = (*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1321240))))).Func94(asWindowEvent(0x4016, int32(*(*uint32)(unsafe.Pointer(uintptr(v5 + 48)))), 0))
				v6 = strMan.GetStringInFile("InputCfg.wnd:PressKey", "C:\\NoxPost\\src\\client\\Gui\\GuiInput.c")
				nox_swprintf((*libc.WChar)(memmap.PtrOff(6112660, 1321256)), libc.CWString("%s\n'%s'"), v6, v9)
				nox_xxx_wndShowModalMB((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1321232))))))
				guiFocus((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1321232))))))
				sub_46C690((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1321232))))))
				return nox_xxx_wndListboxProcPre_4A30D0((*nox_window)(unsafe.Pointer(uintptr(a1))), 16400, uint32(uintptr(unsafe.Pointer((*libc.WChar)(unsafe.Pointer(uintptr(a3)))))), a4)
			}
		}
	} else {
		if a2 != 0x4007 {
			if a2 == 23 {
				return 1
			}
			if a2 != 0x4000 {
				return nox_xxx_wndListboxProcPre_4A30D0((*nox_window)(unsafe.Pointer(uintptr(a1))), a2, uint32(uintptr(unsafe.Pointer((*libc.WChar)(unsafe.Pointer(uintptr(a3)))))), a4)
			}
		}
		if unsafe.Pointer((*uint32)(unsafe.Pointer(uintptr(a3)))) == unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1321228)))).ChildByID(921)) || unsafe.Pointer((*uint32)(unsafe.Pointer(uintptr(a3)))) == unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1321228)))).ChildByID(922)) {
			(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1321240))))).Func94(asWindowEvent(int32(a2), a3, 0))
			(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1321244))))).Func94(asWindowEvent(int32(a2), a3, 0))
			(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1321248))))).Func94(asWindowEvent(int32(a2), a3, 0))
			return nox_xxx_wndListboxProcPre_4A30D0((*nox_window)(unsafe.Pointer(uintptr(a1))), a2, uint32(uintptr(unsafe.Pointer((*libc.WChar)(unsafe.Pointer(uintptr(a3)))))), a4)
		}
	}
	return nox_xxx_wndListboxProcPre_4A30D0((*nox_window)(unsafe.Pointer(uintptr(a1))), a2, uint32(uintptr(unsafe.Pointer((*libc.WChar)(unsafe.Pointer(uintptr(a3)))))), a4)
}
