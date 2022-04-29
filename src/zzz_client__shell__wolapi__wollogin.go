package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	noxflags "github.com/noxworld-dev/opennox/v1/common/flags"
	"unsafe"
)

func nox_game_showWolLogin_44A560() int32 {
	var (
		result int32
		i      int32
		v2     int32
		v3     *uint32
		v4     *uint32
		v5     *uint32
		v6     *uint32
		v7     *uint32
		v8     *uint32
		v9     *byte
		v10    *libc.WChar
		v11    *libc.WChar
		v12    *uint32
		v13    int32
		v14    *uint32
		v15    *libc.WChar
		v16    *uint8
		v17    *uint8
		v18    *byte
		v19    [100]libc.WChar
	)
	noxflags.SetGame(noxflags.GameFlag3)
	sub_43AF50(1)
	sub_41E300(1)
	sub_4A1BE0(1)
	sub_4A1A40(0)
	sub_41D4C0()
	sub_41D440()
	if dword_5d4594_830248 != 0 {
		nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_830248))))), 1)
	} else {
		gameAddStateCode(1700)
		result = int32(uintptr(unsafe.Pointer(newWindowFromFile("wolapi.wnd", unsafe.Pointer(cgoFuncAddr(sub_44AB30))))))
		dword_5d4594_830248 = uint32(result)
		if result == 0 {
			return result
		}
		result = int32(uintptr(unsafe.Pointer(nox_gui_makeAnimation((*nox_window)(unsafe.Pointer(uintptr(result))), 0, 0, 0, -480, 0, 20, 0, -40))))
		nox_wnd_xxx_830244 = (*nox_gui_animation)(unsafe.Pointer(uintptr(result)))
		if result == 0 {
			return result
		}
		nox_wnd_xxx_830244.field_0 = 1700
		nox_wnd_xxx_830244.field_12 = sub_44AA40
		nox_wnd_xxx_830244.fnc_done_out = sub_44AA70
		(*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_830248)))).SetAllFuncs(func(arg1 int32, arg2 int32, arg3 int32, arg4 int32) int32 {
			return sub_44AAC0()
		}, nil, nil)
	}
	sub_41FCF0()
	dword_5d4594_830264 = uint32(uintptr(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_830248)))).ChildByID(1708))))
	(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_830264))))).Func94(asWindowEvent(0x400F, 0, 0))
	for i = 0; i < 128; i++ {
		sub_41FB90(i, (*uint32)(unsafe.Pointer(&v16)), (*uint32)(unsafe.Pointer(&v17)))
		nox_swprintf(&v19[0], libc.CWString("%S"), v16)
		if v19[0] == 0 {
			nox_swprintf(&v19[0], libc.CWString(" "))
		}
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_830264))))).Func94(asWindowEvent(0x400D, int32(uintptr(unsafe.Pointer(&v19[0]))), -1))
	}
	v2 = sub_41FC40()
	if v2 == -1 {
		sub_41FC20(0)
		v2 = 0
	}
	(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_830264))))).Func94(asWindowEvent(0x4013, v2, 0))
	sub_41FB90(v2, (*uint32)(unsafe.Pointer(&v16)), (*uint32)(unsafe.Pointer(&v17)))
	dword_5d4594_830256 = uint32(uintptr(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_830248)))).ChildByID(1701))))
	if v16 != nil && int32(*v16) != 0 {
		nox_swprintf(&v19[0], libc.CWString("%S"), v16)
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_830256))))).Func94(asWindowEvent(0x401E, int32(uintptr(unsafe.Pointer(&v19[0]))), 0))
	}
	dword_5d4594_830260 = uint32(uintptr(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_830248)))).ChildByID(1702))))
	(*(*int32)(unsafe.Pointer(&dword_5d4594_830260))).setFunc93(func(arg1 int32, arg2 int32, arg3 int32, arg4 int32) int32 {
		return sub_44AAD0((*uint32)(unsafe.Pointer(uintptr(arg1))), arg2, arg3, arg4)
	})
	v3 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_830248)))).ChildByID(1703)))
	dword_5d4594_830252 = uint32(uintptr(unsafe.Pointer(v3)))
	if v17 != nil && int32(*v17) != 0 {
		nox_swprintf(&v19[0], libc.CWString("%S"), v17)
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_830260))))).Func94(asWindowEvent(0x401E, int32(uintptr(unsafe.Pointer(&v19[0]))), 0))
		v4 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_830248)))).ChildByID(1703)))
		*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*9)) |= 4
		dword_5d4594_830276 = 1
	} else {
		*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*9)) &= 0xFFFFFFFB
		dword_5d4594_830276 = 0
	}
	v5 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_830248)))).ChildByID(1709)))
	v6 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_830248)))).ChildByID(1710)))
	v7 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_830248)))).ChildByID(1711)))
	v8 = *(**uint32)(unsafe.Pointer(uintptr(dword_5d4594_830264 + 32)))
	v18 = (*byte)(unsafe.Pointer(nox_xxx_gLoadImg("UISlider")))
	v9 = (*byte)(unsafe.Pointer(nox_xxx_gLoadImg("UISliderLit")))
	sub_4B5700(int32(uintptr(unsafe.Pointer(v5))), 0, 0, int32(uintptr(unsafe.Pointer(v18))), int32(uintptr(unsafe.Pointer(v9))), int32(uintptr(unsafe.Pointer(v9))))
	nox_xxx_wnd_46B280(int32(uintptr(unsafe.Pointer(v5))), *(*int32)(unsafe.Pointer(&dword_5d4594_830264)))
	nox_xxx_wnd_46B280(int32(uintptr(unsafe.Pointer(v6))), *(*int32)(unsafe.Pointer(&dword_5d4594_830264)))
	nox_xxx_wnd_46B280(int32(uintptr(unsafe.Pointer(v7))), *(*int32)(unsafe.Pointer(&dword_5d4594_830264)))
	*(*uint32)(unsafe.Add(unsafe.Pointer(v8), unsafe.Sizeof(uint32(0))*9)) = uint32(uintptr(unsafe.Pointer(v5)))
	*(*uint32)(unsafe.Add(unsafe.Pointer(v8), unsafe.Sizeof(uint32(0))*7)) = uint32(uintptr(unsafe.Pointer(v6)))
	*(*uint32)(unsafe.Add(unsafe.Pointer(v8), unsafe.Sizeof(uint32(0))*8)) = uint32(uintptr(unsafe.Pointer(v7)))
	*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(uint32(0))*100)) + 8))) = 16
	*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(uint32(0))*100)) + 12))) = 10
	if dword_5d4594_830272 == 1 || sub_4D3320() == 1 {
		if noxflags.HasGame(noxflags.GameFlag26) {
			if dword_5d4594_830272 == 1 {
				v10 = strMan.GetStringInFile("woldisc.c:LostConnection", "C:\\NoxPost\\src\\client\\shell\\WolApi\\wollogin.c")
				nox_xxx_networkLog_printf_413D30(CString("%S"), v10)
			} else {
				nox_xxx_networkLog_printf_413D30(CString("InitWolApiMenu: Quitting due to all players left"))
			}
			nox_xxx_setContinueMenuOrHost_43DDD0(0)
			nox_game_exit_xxx_43DE60()
			return 0
		}
		v15 = strMan.GetStringInFile("Discon", "C:\\NoxPost\\src\\client\\shell\\WolApi\\wollogin.c")
		v11 = strMan.GetStringInFile("Notice", "C:\\NoxPost\\src\\client\\shell\\WolApi\\wollogin.c")
		NewDialogWindow((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_830248))))), (*libc.WChar)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v11)))))), (*libc.WChar)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v15)))))), 33, nil, nil)
		sub_44A4B0()
		dword_5d4594_830272 = 0
	}
	if sub_44A4A0() == 0 {
		nox_xxx_wndRetNULL_46A8A0()
		guiFocus((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_830256))))))
	}
	dword_5d4594_830268 = 0
	if noxflags.HasGame(noxflags.GameFlag26) {
		v12 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_830248)))).ChildByID(1708)))
		v13 = sub_4A7F00()
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_830248))))).Func94(asWindowEvent(0x4010, int32(uintptr(unsafe.Pointer(v12))), v13-1))
		v14 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_830248)))).ChildByID(1706)))
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_830248))))).Func94(asWindowEvent(0x4007, int32(uintptr(unsafe.Pointer(v14))), 0))
	}
	return 1
}
