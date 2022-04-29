package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"unsafe"
)

func sub_448730_wol_dialogs() int32 {
	var result int32
	if dword_5d4594_830112 != 0 {
		return 0
	}
	if dword_5d4594_830116 != 0 {
		return 0
	}
	if dword_5d4594_830120 != 0 {
		return 0
	}
	result = int32(uintptr(unsafe.Pointer(newWindowFromFile("wolfind.wnd", unsafe.Pointer(cgoFuncAddr(sub_4489C0_wol_dialogs))))))
	dword_5d4594_830124 = uint32(result)
	if result != 0 {
		*(*uint32)(unsafe.Pointer(uintptr(result + 56))) = nox_color_black_2650656
		*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_830124 + 16))) = (uint32(nox_win_width) - *(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_830124 + 8)))) / 2
		*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_830124 + 24))) = *(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_830124 + 8))) + *(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_830124 + 16)))
		*memmap.PtrUint32(6112660, 830128) = uint32(uintptr(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_830124)))).ChildByID(1933))))
		(*memmap.PtrInt32(6112660, 830128)).setFunc93(func(arg1 int32, arg2 int32, arg3 int32, arg4 int32) int32 {
			return sub_448F00((*uint32)(unsafe.Pointer(uintptr(arg1))), arg2, arg3, arg4)
		})
		dword_5d4594_830136 = uint32(uintptr(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_830124)))).ChildByID(1934))))
		dword_5d4594_830132 = uint32(uintptr(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_830124)))).ChildByID(1935))))
		(*(*int32)(unsafe.Pointer(&dword_5d4594_830132))).setFunc93(func(arg1 int32, arg2 int32, arg3 int32, arg4 int32) int32 {
			return sub_448F60((*uint32)(unsafe.Pointer(uintptr(arg1))), arg2, arg3, arg4)
		})
		dword_5d4594_830152 = uint32(uintptr(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_830124)))).ChildByID(1931))))
		*memmap.PtrUint32(6112660, 830156) = uint32(uintptr(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_830124)))).ChildByID(1932))))
		nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_830152))))), 0)
		dword_5d4594_830140 = uint32(uintptr(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_830124)))).ChildByID(1938))))
		*memmap.PtrUint32(6112660, 830144) = uint32(uintptr(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_830124)))).ChildByID(1939))))
		dword_5d4594_830148 = uint32(uintptr(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_830124)))).ChildByID(1936))))
		guiFocus((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_830132))))))
		*memmap.PtrUint32(6112660, 830160) = uint32(uintptr(unsafe.Pointer(strMan.GetStringInFile("WolFind.wnd:Title", "C:\\NoxPost\\src\\client\\shell\\WolApi\\woldlgs.c"))))
		*memmap.PtrUint32(6112660, 830164) = uint32(uintptr(unsafe.Pointer(strMan.GetStringInFile("WolFind.wnd:pageTitle", "C:\\NoxPost\\src\\client\\shell\\WolApi\\woldlgs.c"))))
		*memmap.PtrUint32(6112660, 830168) = uint32(uintptr(unsafe.Pointer(strMan.GetStringInFile("WolFind.wnd:Loc", "C:\\NoxPost\\src\\client\\shell\\WolApi\\woldlgs.c"))))
		*memmap.PtrUint32(6112660, 830172) = uint32(uintptr(unsafe.Pointer(strMan.GetStringInFile("WolFind.wnd:Msg", "C:\\NoxPost\\src\\client\\shell\\WolApi\\woldlgs.c"))))
		*memmap.PtrUint32(6112660, 830176) = uint32(uintptr(unsafe.Pointer(strMan.GetStringInFile("WolFind.wnd:GoFind", "C:\\NoxPost\\src\\client\\shell\\WolApi\\woldlgs.c"))))
		*memmap.PtrUint32(6112660, 830180) = uint32(uintptr(unsafe.Pointer(strMan.GetStringInFile("WolFind.wnd:GoPage", "C:\\NoxPost\\src\\client\\shell\\WolApi\\woldlgs.c"))))
		*memmap.PtrUint32(6112660, 830184) = uint32(uintptr(unsafe.Pointer(strMan.GetStringInFile("searching", "C:\\NoxPost\\src\\client\\shell\\WolApi\\woldlgs.c"))))
		*memmap.PtrUint32(6112660, 830188) = uint32(uintptr(unsafe.Pointer(strMan.GetStringInFile("paging", "C:\\NoxPost\\src\\client\\shell\\WolApi\\woldlgs.c"))))
		dword_587000_111668 = 1
		dword_5d4594_830116 = 1
		sub_46B120((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_830124)))), nil)
		nox_xxx_wndShowModalMB((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_830124))))))
		sub_46C690((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_830124))))))
		guiFocus((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_830124))))))
		(*(*int32)(unsafe.Pointer(&dword_5d4594_830124))).setDraw(sub_448FC0)
		result = 1
	}
	return result
}
func sub_4489C0_wol_dialogs(a1 int32, a2 int32, a3 *int32, a4 int32) int32 {
	var (
		v4     *libc.WChar
		v5     *libc.WChar
		result int32
		v7     *libc.WChar
		v8     int32
	)
	if a2 == 0x4007 {
		v8 = (*nox_window)(unsafe.Pointer(a3)).ID()
		clientPlaySoundSpecial(766, 100)
		switch v8 {
		case 1931:
			(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_830136))))).Show()
			(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_830132))))).Hide()
			(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_830140))))).Func94(asWindowEvent(0x4001, *memmap.PtrInt32(6112660, 830160), 0))
			(*nox_window)(unsafe.Pointer(uintptr(*memmap.PtrInt32(6112660, 830144)))).Func94(asWindowEvent(0x4001, *memmap.PtrInt32(6112660, 830168), 0))
			(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_830148))))).Func94(asWindowEvent(0x4001, *memmap.PtrInt32(6112660, 830176), 0))
			nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_830152))))), 0)
			nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(*memmap.PtrInt32(6112660, 830156)))), 1)
			if dword_5d4594_830204 != 0 {
				nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_830148))))), 0)
			} else {
				nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_830148))))), 1)
			}
			dword_587000_111668 = 1
			result = 0
		case 1932:
			(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_830136))))).Hide()
			(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_830132))))).Show()
			(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_830140))))).Func94(asWindowEvent(0x4001, *memmap.PtrInt32(6112660, 830164), 0))
			(*nox_window)(unsafe.Pointer(uintptr(*memmap.PtrInt32(6112660, 830144)))).Func94(asWindowEvent(0x4001, *memmap.PtrInt32(6112660, 830172), 0))
			(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_830148))))).Func94(asWindowEvent(0x4001, *memmap.PtrInt32(6112660, 830180), 0))
			nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_830152))))), 1)
			nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(*memmap.PtrInt32(6112660, 830156)))), 0)
			dword_587000_111668 = 0
			if dword_5d4594_830208 != 0 {
				nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_830148))))), 0)
			} else {
				nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_830148))))), 1)
			}
			result = 0
		case 1936:
			sub_448CF0_wol_dialogs()
			goto LABEL_24
		case 1937:
			nox_xxx_wnd_46C6E0((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_830124))))))
			(*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_830124)))).Destroy()
			sub_447600()
			dword_5d4594_830116 = 0
			result = 0
		default:
			goto LABEL_24
		}
	} else if a2 == 16400 && (*nox_window)(unsafe.Pointer(a3)).ID() == 1934 && a4 != -1 && (func() libc.WChar {
		v4 = (*libc.WChar)(unsafe.Pointer(uintptr((*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a3)))))).Func94(asWindowEvent(0x4016, a4, 0)))))
		return *v4
	}()) != 0 && (func() *libc.WChar {
		v5 = (*libc.WChar)(unsafe.Pointer(uintptr(sub_41EC00())))
		return v4
	}()) != nil {
		if v5 != nil {
			if nox_wcscmp(v4, v5) == 0 {
				v7 = strMan.GetStringInFile("alreadyinchannel", "C:\\NoxPost\\src\\client\\shell\\WolApi\\woldlgs.c")
			} else {
				nox_xxx_wnd_46C6E0((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_830124))))))
				(*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_830124)))).Destroy()
				sub_447600()
				dword_5d4594_830116 = 0
				dword_5d4594_830124 = 0
				sub_40D380()
				sub_446A20_wol_chat(v4)
				v7 = strMan.GetStringInFile("changedchannel", "C:\\NoxPost\\src\\client\\shell\\WolApi\\woldlgs.c")
			}
			(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_830140))))).Func94(asWindowEvent(0x4001, int32(uintptr(unsafe.Pointer(v7))), 0))
			result = 0
		} else {
			sub_446A20_wol_chat(nil)
			result = 0
		}
	} else {
	LABEL_24:
		result = 0
	}
	return result
}
func sub_448CF0_wol_dialogs() int32 {
	var (
		v0     *uint16
		result int32
		v2     *libc.WChar
		v3     int32
		v4     *libc.WChar
		v6     [36]byte
		v7     [72]byte
		v8     [256]byte
	)
	v0 = (*uint16)(unsafe.Pointer(uintptr((*nox_window)(unsafe.Pointer(uintptr(*memmap.PtrInt32(6112660, 830128)))).Func94(asWindowEvent(0x401D, 0, 0)))))
	if dword_587000_111668 != 0 {
		if v0 != nil && int32(*v0) != 0 {
			nox_sprintf(&v7[0], CString("%S"), v0)
			sub_40D740(int32(uintptr(unsafe.Pointer(&v6[0]))))
			nox_swprintf((*libc.WChar)(memmap.PtrOff(6112660, 829840)), *(**libc.WChar)(memmap.PtrOff(6112660, 830184)), v0)
			(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_830140))))).Func94(asWindowEvent(0x4001, int32(uintptr(memmap.PtrOff(6112660, 829840))), 0))
			dword_5d4594_830204 = 1
			nox_xxx_setKeybTimeout_4160D0(18)
			return nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_830148))))), 0)
		}
		v2 = strMan.GetStringInFile("InvalidUser", "C:\\NoxPost\\src\\client\\shell\\WolApi\\woldlgs.c")
		return (*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_830140))))).Func94(asWindowEvent(0x4001, int32(uintptr(unsafe.Pointer(v2))), 0))
	}
	v3 = (*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_830132))))).Func94(asWindowEvent(0x401D, 0, 0))
	nox_sprintf(&v8[0], CString("%S"), v3)
	if v0 == nil || int32(*v0) == 0 {
		v2 = strMan.GetStringInFile("InvalidUser", "C:\\NoxPost\\src\\client\\shell\\WolApi\\woldlgs.c")
		return (*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_830140))))).Func94(asWindowEvent(0x4001, int32(uintptr(unsafe.Pointer(v2))), 0))
	}
	if v8[0] != 0 {
		nox_sprintf(&v7[0], CString("%S"), v0)
		sub_40D770(int32(uintptr(unsafe.Pointer(&v6[0]))), int32(uintptr(unsafe.Pointer(&v8[0]))))
		nox_swprintf((*libc.WChar)(memmap.PtrOff(6112660, 829840)), *(**libc.WChar)(memmap.PtrOff(6112660, 830188)), v0)
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_830140))))).Func94(asWindowEvent(0x4001, int32(uintptr(memmap.PtrOff(6112660, 829840))), 0))
		dword_5d4594_830208 = 1
		nox_xxx_setKeybTimeout_4160D0(19)
		nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_830148))))), 0)
		result = (*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_830132))))).Func94(asWindowEvent(0x401E, int32(uintptr(memmap.PtrOff(6112660, 830212))), 0))
	} else {
		v4 = strMan.GetStringInFile("InvalidMessage", "C:\\NoxPost\\src\\client\\shell\\WolApi\\woldlgs.c")
		result = (*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_830140))))).Func94(asWindowEvent(0x4001, int32(uintptr(unsafe.Pointer(v4))), 0))
	}
	return result
}
func sub_4490C0_wol_dialogs(a1 int32) {
	var (
		v1 int32
		v3 *byte
		v4 *byte
		v5 *uint32
		v6 *libc.WChar
	)
	v1 = 0
	if sub_43AF80() == 11 {
		sub_43AFC0(a1)
		return
	}
	if dword_5d4594_830116 == 0 {
		return
	}
	v3 = sub_420110()
	if v3 == nil {
		sub_4491B0_wol_dialogs()
		return
	}
	(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_830136))))).Func94(asWindowEvent(0x400F, 0, 0))
	if a1 != 0 {
		v4 = (*byte)(unsafe.Pointer(uintptr(a1 + 52)))
		for {
			if libc.StrNCmp(v4, v3, libc.StrLen(v3)) == 0 {
				v5 = sub_41E7A0(v4)
				(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_830136))))).Func94(asWindowEvent(0x400D, int32(uintptr(unsafe.Pointer(v5))), 9))
				v1++
			}
			if *(*uint32)(unsafe.Pointer(uintptr(a1 + 48))) == 0 {
				break
			}
		}
	}
	v6 = strMan.GetStringInFile("findsuccess", "C:\\NoxPost\\src\\client\\shell\\WolApi\\woldlgs.c")
	(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_830140))))).Func94(asWindowEvent(0x4001, int32(uintptr(unsafe.Pointer(v6))), 0))
	if v1 == 0 {
		sub_4491B0_wol_dialogs()
	}
}
func sub_4491B0_wol_dialogs() int32 {
	var (
		result int32
		v1     *libc.WChar
	)
	if sub_43AF80() == 11 {
		nox_client_setConnError_43AFA0(9)
		return 1
	}
	result = int32(dword_5d4594_830116)
	if dword_5d4594_830116 != 0 {
		v1 = strMan.GetStringInFile("usernoton", "C:\\NoxPost\\src\\client\\shell\\WolApi\\woldlgs.c")
		result = (*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_830140))))).Func94(asWindowEvent(0x4001, int32(uintptr(unsafe.Pointer(v1))), 0))
	}
	return result
}
func sub_449200_wol_dialogs() int32 {
	var (
		result int32
		v1     *libc.WChar
	)
	result = int32(dword_5d4594_830116)
	if dword_5d4594_830116 != 0 {
		v1 = strMan.GetStringInFile("pagesuccess", "C:\\NoxPost\\src\\client\\shell\\WolApi\\woldlgs.c")
		result = (*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_830140))))).Func94(asWindowEvent(0x4001, int32(uintptr(unsafe.Pointer(v1))), 0))
	}
	return result
}
func sub_449240_wol_dialogs() int32 {
	var (
		result int32
		v1     *libc.WChar
	)
	result = int32(dword_5d4594_830116)
	if dword_5d4594_830116 != 0 {
		v1 = strMan.GetStringInFile("usernoton", "C:\\NoxPost\\src\\client\\shell\\WolApi\\woldlgs.c")
		result = (*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_830140))))).Func94(asWindowEvent(0x4001, int32(uintptr(unsafe.Pointer(v1))), 0))
	}
	return result
}
func sub_4497D0_wol_dialogs(a1 int32, a2 int32, a3 *int32, a4 int32) int32 {
	var (
		v3     *uint32
		v4     int32
		result int32
		v6     int32
		v7     int32
		v8     int32
		v9     *uint32
		v10    int32
		v11    int32
		i      *byte
		v13    int32
		v14    *libc.WChar
		v15    *libc.WChar
	)
	if a2 == 0x4005 {
		clientPlaySoundSpecial(920, 100)
		result = 1
	} else if a2 == 0x4007 {
		v6 = (*nox_window)(unsafe.Pointer(a3)).ID() - 1952
		if v6 != 0 {
			v7 = v6 - 10
			if v7 != 0 {
				if v7 == 1 {
					nox_xxx_wnd_46C6E0((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_830108))))))
					(*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_830108)))).Destroy()
					dword_5d4594_830108 = 0
					if sub_41E2F0() != 7 {
						sub_468020()
						clientPlaySoundSpecial(921, 100)
						return 1
					}
				}
			} else {
				v8 = 0
				v9 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_830108)))).ChildByID(1961)))
				v10 = (*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v9)))))).Func94(asWindowEvent(0x4014, 0, 0))
				if v10 >= 0 {
					v11 = sub_4200F0()
					for i = sub_4205B0(v11); i != nil; i = (*byte)(unsafe.Pointer(uintptr(sub_4206B0(int32(uintptr(unsafe.Pointer(i))))))) {
						if v10 == 0 {
							break
						}
						v10--
					}
					sub_4207D0(int32(uintptr(unsafe.Pointer(i))))
					if sub_41E2F0() == 7 {
						v8 = 1
						sub_41E300(3)
					}
					v13 = sub_41E2F0()
					sub_41DA70(v13, 3)
				}
				nox_xxx_wnd_46C6E0((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_830108))))))
				(*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_830108)))).Destroy()
				dword_5d4594_830108 = 0
				if v8 != 0 {
					v15 = strMan.GetStringInFile("Noxworld.c:ConChatServ", "C:\\NoxPost\\src\\client\\shell\\WolApi\\woldlgs.c")
					v14 = strMan.GetStringInFile("wolchat.c:PleaseWait", "C:\\NoxPost\\src\\client\\shell\\WolApi\\woldlgs.c")
					NewDialogWindow(nil, (*libc.WChar)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v14)))))), (*libc.WChar)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v15)))))), 288, nil, nil)
					sub_44A4B0()
					clientPlaySoundSpecial(921, 100)
					return 1
				}
			}
		} else {
			nox_xxx_wnd_46C6E0((*nox_window)(unsafe.Pointer(uintptr(*memmap.PtrInt32(6112660, 830096)))))
			(*nox_window)(unsafe.Pointer(*(**uint32)(memmap.PtrOff(6112660, 830096)))).Destroy()
			dword_5d4594_830112 = 0
			sub_447600()
		}
		clientPlaySoundSpecial(921, 100)
		result = 1
	} else {
		if a2 == 16400 && (*nox_window)(unsafe.Pointer(a3)).ID() == 1961 {
			v3 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_830108)))).ChildByID(1962)))
			v4 = (*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a3)))))).Func94(asWindowEvent(0x4014, 0, 0))
			nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v3)))))), bool2int(v4 >= 0))
		}
		result = 0
	}
	return result
}
