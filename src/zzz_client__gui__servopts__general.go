package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	noxflags "github.com/noxworld-dev/opennox/v1/common/flags"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"unsafe"
)

func nox_xxx_gui_4AD320(a1 int32) int32 {
	var (
		v1  int32
		v3  *uint32
		v4  *uint32
		v5  *uint32
		v6  *uint32
		v7  **byte
		v8  *libc.WChar
		v9  *uint32
		v10 *uint32
	)
	v1 = strMan.Lang()
	if nox_xxx_guiFontHeightMB_43F320(nil) > 10 {
		v1 = 2
	}
	if dword_5d4594_1309812 != 0 {
		return 0
	}
	if nox_xxx_check_flag_aaa_43AF70() == 1 {
		v3 = (*uint32)(unsafe.Pointer(newWindowFromFile(*(**byte)(memmap.PtrOff(0x587000, uint32(v1*4)+0x2A61C)), unsafe.Pointer(cgoFuncAddr(nox_xxx_windowServerOptionsGeneralProc_4AD5D0)))))
	} else {
		v3 = (*uint32)(unsafe.Pointer(newWindowFromFile(*(**byte)(memmap.PtrOff(0x587000, uint32(v1*4)+0x2A5F4)), unsafe.Pointer(cgoFuncAddr(nox_xxx_windowServerOptionsGeneralProc_4AD5D0)))))
	}
	dword_5d4594_1309812 = uint32(uintptr(unsafe.Pointer(v3)))
	sub_46B120((*nox_window)(unsafe.Pointer(v3)), (*nox_window)(unsafe.Pointer(uintptr(a1))))
	(*(*int32)(unsafe.Pointer(&dword_5d4594_1309812))).setDraw(func(arg1 int32, arg2 int32) int32 {
		return sub_4AD570()
	})
	v4 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1309812)))).ChildByID(0x2842)))
	if noxflags.HasGame(noxflags.GameModeCTF | noxflags.GameModeElimination) {
		nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v4)))))), 0)
	}
	nox_xxx_wndRetNULL_46A8A0()
	sub_4AD840()
	if nox_xxx_check_flag_aaa_43AF70() == 1 {
		sub_4AD4B0()
		v5 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1309812)))).ChildByID(10310)))
		nox_xxx_wnd_46B280(int32(uintptr(unsafe.Pointer(v5))), *(*int32)(unsafe.Pointer(&dword_5d4594_1309812)))
		int32(uintptr(unsafe.Pointer(v5))).setFunc94(func(arg1 int32, arg2 int32, arg3 int32, arg4 int32) int32 {
			return nox_xxx_windowServerOptionsGeneralProc_4AD5D0(arg1, arg2, (*int32)(unsafe.Pointer(uintptr(arg3))), arg4)
		})
		v6 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1309812)))).ChildByID(0x284D)))
		v7 = (**byte)(memmap.PtrOff(0x587000, 173540))
		for {
			v8 = strMan.GetStringInFile(*v7, "C:\\NoxPost\\src\\client\\Gui\\ServOpts\\general.c")
			(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v6)))))).Func94(asWindowEvent(0x400D, int32(uintptr(unsafe.Pointer(v8))), -1))
			v7 = (**byte)(unsafe.Add(unsafe.Pointer(v7), unsafe.Sizeof((*byte)(nil))*1))
			if int32(uintptr(unsafe.Pointer(v7))) >= int32(uintptr(memmap.PtrOff(0x587000, 173556))) {
				break
			}
		}
	} else {
		v9 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1309812)))).ChildByID(0x284F)))
		(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v9)))))).Hide()
	}
	if nox_common_getEngineFlag(nox_engine_flag(NOX_ENGINE_FLAG_DISABLE_GRAPHICS_RENDERING)) {
		v10 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1309812)))).ChildByID(0x2840)))
		nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v10)))))), 0)
	}
	return int32(dword_5d4594_1309812)
}
func sub_4AD4B0() int32 {
	var (
		v0     *uint32
		v1     *uint32
		v2     int32
		v3     int32
		v4     **byte
		v5     *uint16
		result int32
		v7     int32
		v8     int32
		v9     int32
	)
	v0 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1309812)))).ChildByID(0x284D)))
	v1 = v0
	v2 = nox_xxx_guiFontHeightMB_43F320(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v0), unsafe.Sizeof(uint32(0))*59))))) + 1
	v3 = 0
	v4 = (**byte)(memmap.PtrOff(0x587000, 173540))
	*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*7)) = *(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*5)) + uint32(v2*4) + 2
	*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*3)) = uint32(v2*4 + 2)
	for {
		v5 = (*uint16)(unsafe.Pointer(strMan.GetStringInFile(*v4, "C:\\NoxPost\\src\\client\\Gui\\ServOpts\\general.c")))
		nox_xxx_drawGetStringSize_43F840(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*59)))), (*libc.WChar)(unsafe.Pointer(v5)), &v9, nil, 0)
		if v9 > v3 {
			v3 = v9
		}
		v4 = (**byte)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof((*byte)(nil))*1))
		if int32(uintptr(unsafe.Pointer(v4))) >= int32(uintptr(memmap.PtrOff(0x587000, 173556))) {
			break
		}
	}
	result = (*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1309812)))).ChildByID(0x284C).width
	v7 = v3 + 7
	if v7 <= result {
		v7 = result
	}
	v8 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*6)))
	*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*2)) = uint32(v7)
	*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*4)) = uint32(v8 - v7)
	return result
}
func sub_4AD840() int32 {
	var (
		result int32
		v1     *uint32
		v2     *uint32
		v3     *uint32
		v4     *uint32
		v5     *uint32
		v6     *uint32
		v7     *libc.WChar
		v8     *uint32
		v9     int32
	)
	result = int32(dword_5d4594_1309812)
	if dword_5d4594_1309812 != 0 {
		if nox_server_doPlayersAutoRespawn_40A5F0() != 0 {
			v1 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1309812)))).ChildByID(0x283D)))
			*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*9)) |= 4
			if noxflags.HasGame(noxflags.GameModeElimination) {
				nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v1)))))), 0)
			}
		}
		if nox_server_sendMotd_108752 != 0 {
			v2 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1309812)))).ChildByID(0x283E)))
			*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*9)) |= 4
		}
		if sub_4D0D70() != 0 {
			v3 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1309812)))).ChildByID(0x2840)))
			*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*9)) |= 4
		}
		if sub_409F40(2) != 0 {
			v4 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1309812)))).ChildByID(0x2841)))
			*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*9)) |= 4
		}
		if sub_409F40(8192) != 0 {
			v5 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1309812)))).ChildByID(0x2842)))
			*(*uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(uint32(0))*9)) |= 4
		}
		result = nox_xxx_check_flag_aaa_43AF70()
		if result == 1 {
			v6 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1309812)))).ChildByID(0x284C)))
			v7 = strMan.GetStringInFile(*(**byte)(memmap.PtrOff(0x587000, nox_server_connectionType_3596*4+0x2A5E0)), "C:\\NoxPost\\src\\client\\Gui\\ServOpts\\general.c")
			(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v6)))))).Func94(asWindowEvent(0x4001, int32(uintptr(unsafe.Pointer(v7))), -1))
			v8 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1309812)))).ChildByID(0x2848)))
			v9 = nox_xxx_rateGet_40A6C0()
			result = (*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v8)))))).Func94(asWindowEvent(0x400A, 4-v9, 0))
		}
	}
	return result
}
