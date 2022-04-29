package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	noxflags "github.com/noxworld-dev/opennox/v1/common/flags"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"math"
	"unsafe"
)

func sub_446970_wol_chat() *byte {
	var (
		result *byte
		v1     int32
		v2     *uint8
	)
	result = *(**byte)(memmap.PtrOff(6112660, 829568))
	if *memmap.PtrUint32(6112660, 829568) == 0 {
		result = *(**byte)(memmap.PtrOff(0x587000, 108852))
		v1 = 0
		if *memmap.PtrUint32(0x587000, 108852) != 0 {
			v2 = (*uint8)(memmap.PtrOff(0x587000, 108848))
			for {
				*(*uint32)(unsafe.Pointer(v2)) = uint32(uintptr(unsafe.Pointer(strMan.GetStringInFile(*((**byte)(unsafe.Add(unsafe.Pointer((**byte)(unsafe.Pointer(v2))), unsafe.Sizeof((*byte)(nil))*1))), "C:\\NoxPost\\src\\client\\shell\\WolApi\\wolchat.c"))))
				result = (*byte)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), unsafe.Sizeof(uint32(0))*3))))))
				v2 = (*uint8)(unsafe.Add(unsafe.Pointer(v2), 8))
				v1++
				if result == nil {
					break
				}
			}
		}
		*memmap.PtrUint32(6112660, 826092) = uint32(v1)
		*memmap.PtrUint32(6112660, 829568) = 1
	}
	return result
}
func sub_446A20_wol_chat(a1 *libc.WChar) int32 {
	var (
		v1 *libc.WChar
		v3 *libc.WChar
	)
	sub_446CC0(0)
	sub_41F140(a1)
	sub_44A400()
	v3 = strMan.GetStringInFile("UpdatingChannels", "C:\\NoxPost\\src\\client\\shell\\WolApi\\wolchat.c")
	v1 = strMan.GetStringInFile("PleaseWait", "C:\\NoxPost\\src\\client\\shell\\WolApi\\wolchat.c")
	NewDialogWindow((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_829480))))), (*libc.WChar)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v1)))))), (*libc.WChar)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v3)))))), 288, nil, nil)
	sub_44A4B0()
	return 0
}
func sub_446AE0_wol_chat(a1 int32) *byte {
	var (
		result *byte
		v2     *byte
		v3     [108]byte
	)
	result = sub_446BC0(a1)
	v2 = result
	if result != nil {
		libc.StrCpy(&v3[36], result)
		if sub_41F9E0(int32(uintptr(unsafe.Pointer(&v3[0])))) != 0 {
			result = (*byte)(unsafe.Pointer(strMan.GetStringInFile("cantignore", "C:\\NoxPost\\src\\client\\shell\\WolApi\\wolchat.c")))
			if result != nil {
				result = (*byte)(unsafe.Pointer(uintptr(sub_447310(0, int32(uintptr(unsafe.Pointer(result)))))))
			}
		} else if sub_427880(v2) != 0 {
			sub_427850(int32(uintptr(unsafe.Pointer(&v3[0]))))
			sub_41F980(v2, 0)
			result = (*byte)(unsafe.Pointer(uintptr(sub_446C10(a1, 0))))
		} else {
			sub_427820(int32(uintptr(unsafe.Pointer(&v3[0]))))
			sub_41F980(v2, 1)
			result = (*byte)(unsafe.Pointer(uintptr(sub_446C10(a1, 1))))
		}
	}
	return result
}
func sub_446F80_wol_chat(a1 *libc.WChar) {
	var (
		v1 libc.WChar
		v2 *int32
		v3 *int32
		i  int32
		v5 *byte
		v6 *byte
		v7 *byte
		v8 *libc.WChar
	)
	if a1 == nil {
		return
	}
	v1 = *a1
	if *a1 == 0 || v1 == 10 || v1 == 13 {
		return
	}
	if sub_41EC00() != 0 {
		if sub_446C90() != 0 {
			v2 = (*int32)(unsafe.Pointer(uintptr((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_829488))))).Func94(asWindowEvent(0x4014, 0, 0)))))
			v3 = v2
			for i = *v2; i != -1; v3 = (*int32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(int32(0))*1)) {
				v5 = sub_446F40(i)
				v6 = v5
				if v5 != nil {
					if sub_41F860(v5, a1) != 0 {
						sub_4471A0(v6, int32(uintptr(unsafe.Pointer(a1))), 1, 1)
					}
				}
				i = *(*int32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(int32(0))*1))
			}
			return
		}
		sub_448680(a1, (*byte)(memmap.PtrOff(6112660, 827176)))
		if sub_40D3E0(int32(uintptr(memmap.PtrOff(6112660, 827176)))) != 0 {
			v7 = sub_41FA40()
			sub_4471A0(v7, int32(uintptr(unsafe.Pointer(a1))), 0, 0)
			return
		}
	}
	v8 = strMan.GetStringInFile("NotJoined", "C:\\NoxPost\\src\\client\\shell\\WolApi\\wolchat.c")
	if v8 != nil {
		sub_447310(0, int32(uintptr(unsafe.Pointer(v8))))
	}
}
func sub_447090_wol_chat(a1 *libc.WChar) {
	var (
		v1 libc.WChar
		v2 *int32
		v3 *int32
		i  int32
		v5 *byte
		v6 *byte
		v7 *byte
		v8 *libc.WChar
	)
	if a1 == nil {
		return
	}
	v1 = *a1
	if *a1 == 0 || v1 == 10 || v1 == 13 {
		return
	}
	if sub_41EC00() != 0 {
		if sub_446C90() != 0 {
			v2 = (*int32)(unsafe.Pointer(uintptr((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_829488))))).Func94(asWindowEvent(0x4014, 0, 0)))))
			v3 = v2
			for i = *v2; i != -1; v3 = (*int32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(int32(0))*1)) {
				v5 = sub_446F40(i)
				v6 = v5
				if v5 != nil {
					if sub_41F8F0(v5, a1) != 0 {
						sub_447250(v6, int32(uintptr(unsafe.Pointer(a1))), 1, 1)
					}
				}
				i = *(*int32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(int32(0))*1))
			}
			return
		}
		sub_448680(a1, (*byte)(memmap.PtrOff(6112660, 827176)))
		if sub_40D4D0(int32(uintptr(memmap.PtrOff(6112660, 827176)))) != 0 {
			v7 = sub_41FA40()
			sub_447250(v7, int32(uintptr(unsafe.Pointer(a1))), 0, 0)
			return
		}
	}
	v8 = strMan.GetStringInFile("NotJoined", "C:\\NoxPost\\src\\client\\shell\\WolApi\\wolchat.c")
	if v8 != nil {
		sub_447310(0, int32(uintptr(unsafe.Pointer(v8))))
	}
}
func sub_447470_wol_chat() int32 {
	var (
		v0     *libc.WChar
		result int32
	)
	v0 = strMan.GetStringInFile("joined", "C:\\NoxPost\\src\\client\\shell\\WolApi\\wolchat.c")
	result = sub_41EC00()
	if result != 0 {
		nox_swprintf((*libc.WChar)(memmap.PtrOff(6112660, 827176)), v0, result)
		result = (*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_829492))))).Func94(asWindowEvent(0x400D, int32(uintptr(memmap.PtrOff(6112660, 827176))), 9))
	}
	return result
}
func sub_4474C0_wol_chat(a1 int32, a2 int32) int32 {
	var (
		v2 *libc.WChar
		v3 *libc.WChar
	)
	if a1 != 0 {
		v2 = strMan.GetStringInFile("kicked", "C:\\NoxPost\\src\\client\\shell\\WolApi\\wolchat.c")
		nox_swprintf((*libc.WChar)(memmap.PtrOff(6112660, 827176)), v2, a2, a1)
	} else {
		v3 = strMan.GetStringInFile("youwerekicked", "C:\\NoxPost\\src\\client\\shell\\WolApi\\wolchat.c")
		nox_swprintf((*libc.WChar)(memmap.PtrOff(6112660, 827176)), v3, a2)
	}
	return (*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_829492))))).Func94(asWindowEvent(0x400D, int32(uintptr(memmap.PtrOff(6112660, 827176))), 9))
}
func sub_447540_wol_chat(a1 int32) int32 {
	var v1 *libc.WChar
	v1 = strMan.GetStringInFile("banned", "C:\\NoxPost\\src\\client\\shell\\WolApi\\wolchat.c")
	nox_swprintf((*libc.WChar)(memmap.PtrOff(6112660, 827176)), v1, a1)
	return (*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_829492))))).Func94(asWindowEvent(0x400D, int32(uintptr(memmap.PtrOff(6112660, 827176))), 9))
}
func sub_447590_wol_chat() int32 {
	var (
		v0     *libc.WChar
		result int32
	)
	v0 = strMan.GetStringInFile("left", "C:\\NoxPost\\src\\client\\shell\\WolApi\\wolchat.c")
	result = sub_41EC00()
	if result != 0 {
		nox_swprintf((*libc.WChar)(memmap.PtrOff(6112660, 827176)), v0, result)
		result = (*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_829492))))).Func94(asWindowEvent(0x400D, int32(uintptr(memmap.PtrOff(6112660, 827176))), 9))
	}
	return result
}
func nox_game_showWolChat_447620() int32 {
	var (
		result int32
		v1     **byte
		v2     *uint8
		v3     *libc.WChar
		v4     *uint8
		v5     *uint32
		v6     *byte
		v7     *byte
		v8     *uint32
		v9     *uint32
	)
	gameAddStateCode(1900)
	gameSetCliDrawFunc(unsafe.Pointer(cgoFuncAddr(sub_41E210)))
	nox_draw_setTabWidth_43FE20(150)
	dword_5d4594_829480 = uint32(uintptr(unsafe.Pointer(newWindowFromFile("wolchat.wnd", unsafe.Pointer(cgoFuncAddr(sub_447CC0))))))
	sub_446970_wol_chat()
	result = int32(dword_5d4594_829480)
	if dword_5d4594_829480 != 0 {
		nox_xxx_wndShowModalMB((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_829480))))))
		sub_44A4B0()
		(*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_829480)))).SetAllFuncs(sub_447C70, nil, nil)
		(*(*int32)(unsafe.Pointer(&dword_5d4594_829480))).setDraw(sub_4483A0)
		result = int32(uintptr(unsafe.Pointer(nox_gui_makeAnimation((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_829480)))), 0, 0, 0, -480, 0, 20, 0, -40))))
		nox_wnd_xxx_829520 = (*nox_gui_animation)(unsafe.Pointer(uintptr(result)))
		if nox_wnd_xxx_829520 != nil {
			nox_wnd_xxx_829520.field_0 = 1900
			nox_wnd_xxx_829520.field_12 = sub_446A90
			nox_wnd_xxx_829520.fnc_done_out = sub_447BD0
			dword_5d4594_829484 = uint32(uintptr(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_829480)))).ChildByID(1908))))
			dword_5d4594_829488 = uint32(uintptr(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_829480)))).ChildByID(1909))))
			dword_5d4594_829492 = uint32(uintptr(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_829480)))).ChildByID(1907))))
			*memmap.PtrUint32(6112660, 829496) = uint32(uintptr(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_829480)))).ChildByID(1916))))
			dword_5d4594_829500 = uint32(uintptr(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_829480)))).ChildByID(1910))))
			*memmap.PtrUint32(6112660, 829528) = uint32(uintptr(unsafe.Pointer(strMan.GetStringInFile("messageTitle", "C:\\NoxPost\\src\\client\\shell\\WolApi\\wolchat.c"))))
			dword_5d4594_829532 = uint32(uintptr(unsafe.Pointer(strMan.GetStringInFile("NoChannel", "C:\\NoxPost\\src\\client\\shell\\WolApi\\wolchat.c"))))
			*memmap.PtrUint32(6112660, 829536) = uint32(uintptr(memmap.PtrOff(0x587000, 110992)))
			*memmap.PtrUint32(6112660, 829540) = uint32(uintptr(unsafe.Pointer(strMan.GetStringInFile("pageMsg", "C:\\NoxPost\\src\\client\\shell\\WolApi\\wolchat.c"))))
			dword_5d4594_829544 = uint32(uintptr(memmap.PtrOff(0x587000, 111056)))
			*memmap.PtrUint32(6112660, 829548) = uint32(uintptr(unsafe.Pointer(strMan.GetStringInFile("PrivateMessage", "C:\\NoxPost\\src\\client\\shell\\WolApi\\wolchat.c"))))
			*memmap.PtrUint32(6112660, 829552) = uint32(uintptr(unsafe.Pointer(strMan.GetStringInFile("PrivateToMessage", "C:\\NoxPost\\src\\client\\shell\\WolApi\\wolchat.c"))))
			*memmap.PtrUint32(6112660, 829556) = uint32(uintptr(memmap.PtrOff(0x587000, 111204)))
			*memmap.PtrUint32(6112660, 829560) = uint32(uintptr(unsafe.Pointer(strMan.GetStringInFile("PrivateAction", "C:\\NoxPost\\src\\client\\shell\\WolApi\\wolchat.c"))))
			*memmap.PtrUint32(6112660, 829564) = uint32(uintptr(unsafe.Pointer(strMan.GetStringInFile("PrivateToAction", "C:\\NoxPost\\src\\client\\shell\\WolApi\\wolchat.c"))))
			*memmap.PtrUint32(6112660, 829524) = uint32(uintptr(unsafe.Pointer(strMan.GetStringInFile("Stats", "C:\\NoxPost\\src\\client\\shell\\WolApi\\wolchat.c"))))
			dword_587000_109280 = math.MaxUint32
			sub_446D50()
			nox_xxx_wndRetNULL_46A8A0()
			sub_447600()
			sub_41EBB0(0)
			sub_446CC0(0)
			sub_41F100(CString("Chat Channels"))
			if int32(**(**uint8)(memmap.PtrOff(0x587000, 108828))) != 0 {
				v1 = (**byte)(memmap.PtrOff(0x587000, 108828))
				v2 = (*uint8)(memmap.PtrOff(0x587000, 108828))
				for {
					v3 = strMan.GetStringInFile(*v1, "C:\\NoxPost\\src\\client\\shell\\WolApi\\wolchat.c")
					*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), -int(unsafe.Sizeof(uint32(0))*1)))) = uint32(uintptr(unsafe.Pointer(v3)))
					nox_wcslen(v3)
					v4 = (*uint8)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), unsafe.Sizeof(uint32(0))*3))))))
					v2 = (*uint8)(unsafe.Add(unsafe.Pointer(v2), 12))
					v1 = (**byte)(unsafe.Pointer(v2))
					if int32(*v4) == 0 {
						break
					}
				}
			}
			sub_41FFF0()
			dword_5d4594_829508 = uint32(uintptr(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_829480)))).ChildByID(1917))))
			dword_5d4594_829512 = uint32(uintptr(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_829480)))).ChildByID(1918))))
			dword_5d4594_829516 = uint32(uintptr(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_829480)))).ChildByID(1919))))
			v5 = *(**uint32)(unsafe.Pointer(uintptr(dword_5d4594_829492 + 32)))
			v6 = (*byte)(unsafe.Pointer(nox_xxx_gLoadImg("ShopInventorySlider")))
			v7 = (*byte)(unsafe.Pointer(nox_xxx_gLoadImg("ShopInventorySliderSelected")))
			sub_4B5700(*(*int32)(unsafe.Pointer(&dword_5d4594_829508)), 0, 0, int32(uintptr(unsafe.Pointer(v6))), int32(uintptr(unsafe.Pointer(v7))), int32(uintptr(unsafe.Pointer(v7))))
			nox_xxx_wnd_46B280(*(*int32)(unsafe.Pointer(&dword_5d4594_829508)), *(*int32)(unsafe.Pointer(&dword_5d4594_829492)))
			nox_xxx_wnd_46B280(*(*int32)(unsafe.Pointer(&dword_5d4594_829512)), *(*int32)(unsafe.Pointer(&dword_5d4594_829492)))
			nox_xxx_wnd_46B280(*(*int32)(unsafe.Pointer(&dword_5d4594_829516)), *(*int32)(unsafe.Pointer(&dword_5d4594_829492)))
			(*(*int32)(unsafe.Pointer(&dword_5d4594_829492))).setFunc93(func(arg1 int32, arg2 int32, arg3 int32, arg4 int32) int32 {
				return sub_447BF0((*uint32)(unsafe.Pointer(uintptr(arg1))), arg2, uint32(arg3), arg4)
			})
			*(*uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(uint32(0))*9)) = dword_5d4594_829508
			*(*uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(uint32(0))*7)) = dword_5d4594_829512
			*(*uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(uint32(0))*8)) = dword_5d4594_829516
			*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_829508 + 400))) + 8))) = 16
			*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_829508 + 400))) + 12))) = 46
			dword_5d4594_829508 = uint32(uintptr(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_829480)))).ChildByID(1920))))
			dword_5d4594_829512 = uint32(uintptr(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_829480)))).ChildByID(1921))))
			dword_5d4594_829516 = uint32(uintptr(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_829480)))).ChildByID(1922))))
			v8 = *(**uint32)(unsafe.Pointer(uintptr(dword_5d4594_829484 + 32)))
			sub_4B5700(*(*int32)(unsafe.Pointer(&dword_5d4594_829508)), 0, 0, int32(uintptr(unsafe.Pointer(v6))), int32(uintptr(unsafe.Pointer(v7))), int32(uintptr(unsafe.Pointer(v7))))
			nox_xxx_wnd_46B280(*(*int32)(unsafe.Pointer(&dword_5d4594_829508)), *(*int32)(unsafe.Pointer(&dword_5d4594_829484)))
			nox_xxx_wnd_46B280(*(*int32)(unsafe.Pointer(&dword_5d4594_829512)), *(*int32)(unsafe.Pointer(&dword_5d4594_829484)))
			nox_xxx_wnd_46B280(*(*int32)(unsafe.Pointer(&dword_5d4594_829516)), *(*int32)(unsafe.Pointer(&dword_5d4594_829484)))
			(*(*int32)(unsafe.Pointer(&dword_5d4594_829484))).setFunc93(func(arg1 int32, arg2 int32, arg3 int32, arg4 int32) int32 {
				return sub_447BF0((*uint32)(unsafe.Pointer(uintptr(arg1))), arg2, uint32(arg3), arg4)
			})
			*(*uint32)(unsafe.Add(unsafe.Pointer(v8), unsafe.Sizeof(uint32(0))*9)) = dword_5d4594_829508
			*(*uint32)(unsafe.Add(unsafe.Pointer(v8), unsafe.Sizeof(uint32(0))*7)) = dword_5d4594_829512
			*(*uint32)(unsafe.Add(unsafe.Pointer(v8), unsafe.Sizeof(uint32(0))*8)) = dword_5d4594_829516
			*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_829508 + 400))) + 8))) = 16
			*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_829508 + 400))) + 12))) = 46
			dword_5d4594_829508 = uint32(uintptr(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_829480)))).ChildByID(1923))))
			dword_5d4594_829512 = uint32(uintptr(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_829480)))).ChildByID(1924))))
			dword_5d4594_829516 = uint32(uintptr(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_829480)))).ChildByID(1925))))
			v9 = *(**uint32)(unsafe.Pointer(uintptr(dword_5d4594_829488 + 32)))
			sub_4B5700(*(*int32)(unsafe.Pointer(&dword_5d4594_829508)), 0, 0, int32(uintptr(unsafe.Pointer(v6))), int32(uintptr(unsafe.Pointer(v7))), int32(uintptr(unsafe.Pointer(v7))))
			nox_xxx_wnd_46B280(*(*int32)(unsafe.Pointer(&dword_5d4594_829508)), *(*int32)(unsafe.Pointer(&dword_5d4594_829488)))
			nox_xxx_wnd_46B280(*(*int32)(unsafe.Pointer(&dword_5d4594_829512)), *(*int32)(unsafe.Pointer(&dword_5d4594_829488)))
			nox_xxx_wnd_46B280(*(*int32)(unsafe.Pointer(&dword_5d4594_829516)), *(*int32)(unsafe.Pointer(&dword_5d4594_829488)))
			(*(*int32)(unsafe.Pointer(&dword_5d4594_829488))).setFunc93(func(arg1 int32, arg2 int32, arg3 int32, arg4 int32) int32 {
				return sub_447BF0((*uint32)(unsafe.Pointer(uintptr(arg1))), arg2, uint32(arg3), arg4)
			})
			*(*uint32)(unsafe.Add(unsafe.Pointer(v9), unsafe.Sizeof(uint32(0))*9)) = dword_5d4594_829508
			*(*uint32)(unsafe.Add(unsafe.Pointer(v9), unsafe.Sizeof(uint32(0))*7)) = dword_5d4594_829512
			*(*uint32)(unsafe.Add(unsafe.Pointer(v9), unsafe.Sizeof(uint32(0))*8)) = dword_5d4594_829516
			*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_829508 + 400))) + 8))) = 16
			*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_829508 + 400))) + 12))) = 46
			sub_41F370(0)
			if noxflags.HasGame(noxflags.GameFlag26) {
				sub_448450()
			} else {
				sub_41F3A0(-99999, 1)
			}
			result = 1
		}
	}
	return result
}
