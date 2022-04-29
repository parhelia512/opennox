package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	"github.com/noxworld-dev/opennox/v1/common/alloc"
	noxflags "github.com/noxworld-dev/opennox/v1/common/flags"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"unsafe"
)

func nox_xxx_guiServerPlayersLoad_456270(a1 int32) int32 {
	var (
		v1  int32
		v3  *uint32
		v4  *uint32
		v5  *byte
		v6  *uint32
		v7  *uint32
		v8  *uint32
		v9  *uint32
		v10 *libc.WChar
		v11 *uint32
		v12 *uint32
		v13 *byte
		v14 *uint32
		v15 *uint32
		v16 *uint32
	)
	v1 = strMan.Lang()
	if nox_xxx_guiFontHeightMB_43F320(nil) > 10 {
		v1 = 2
	}
	if dword_5d4594_1045684 != 0 {
		return 0
	}
	dword_5d4594_1045684 = uint32(uintptr(unsafe.Pointer(newWindowFromFile(*(**byte)(memmap.PtrOff(0x587000, uint32(v1*4)+0x1F818)), unsafe.Pointer(cgoFuncAddr(sub_4567C0))))))
	dword_5d4594_1045688 = uint32(uintptr(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045684)))).ChildByID(0x290B))))
	dword_5d4594_1045692 = uint32(uintptr(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045684)))).ChildByID(0x290D))))
	sub_46B120((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045684)))), (*nox_window)(unsafe.Pointer(uintptr(a1))))
	(*(*int32)(unsafe.Pointer(&dword_5d4594_1045684))).setDraw(sub_456640)
	nox_xxx_wndRetNULL_46A8A0()
	nox_common_list_clear_425760((*nox_list_item_t)(memmap.PtrOff(6112660, 1045652)))
	nox_common_list_clear_425760((*nox_list_item_t)(memmap.PtrOff(6112660, 1045668)))
	v3 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045684)))).ChildByID(0x2905)))
	v4 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045684)))).ChildByID(0x2906)))
	v14 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*8)))))
	v13 = (*byte)(unsafe.Pointer(nox_xxx_gLoadImg("UISlider")))
	v5 = (*byte)(unsafe.Pointer(nox_xxx_gLoadImg("UISliderLit")))
	v6 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045684)))).ChildByID(0x2915)))
	v15 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045684)))).ChildByID(0x2913)))
	v11 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045684)))).ChildByID(0x2914)))
	*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*100)) + 8))) = 16
	*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*100)) + 12))) = 10
	sub_4B5700(int32(uintptr(unsafe.Pointer(v6))), 0, 0, int32(uintptr(unsafe.Pointer(v13))), int32(uintptr(unsafe.Pointer(v5))), int32(uintptr(unsafe.Pointer(v5))))
	nox_xxx_wnd_46B280(int32(uintptr(unsafe.Pointer(v6))), int32(uintptr(unsafe.Pointer(v3))))
	nox_xxx_wnd_46B280(int32(uintptr(unsafe.Pointer(v15))), int32(uintptr(unsafe.Pointer(v3))))
	nox_xxx_wnd_46B280(int32(uintptr(unsafe.Pointer(v11))), int32(uintptr(unsafe.Pointer(v3))))
	*(*uint32)(unsafe.Add(unsafe.Pointer(v14), unsafe.Sizeof(uint32(0))*9)) = uint32(uintptr(unsafe.Pointer(v6)))
	*(*uint32)(unsafe.Add(unsafe.Pointer(v14), unsafe.Sizeof(uint32(0))*7)) = uint32(uintptr(unsafe.Pointer(v15)))
	*(*uint32)(unsafe.Add(unsafe.Pointer(v14), unsafe.Sizeof(uint32(0))*8)) = uint32(uintptr(unsafe.Pointer(v11)))
	v7 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*8)))))
	v8 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045684)))).ChildByID(10520)))
	v16 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045684)))).ChildByID(0x2916)))
	v12 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045684)))).ChildByID(0x2917)))
	*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v8), unsafe.Sizeof(uint32(0))*100)) + 8))) = 16
	*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v8), unsafe.Sizeof(uint32(0))*100)) + 12))) = 10
	sub_4B5700(int32(uintptr(unsafe.Pointer(v8))), 0, 0, int32(uintptr(unsafe.Pointer(v13))), int32(uintptr(unsafe.Pointer(v5))), int32(uintptr(unsafe.Pointer(v5))))
	nox_xxx_wnd_46B280(int32(uintptr(unsafe.Pointer(v8))), int32(uintptr(unsafe.Pointer(v4))))
	nox_xxx_wnd_46B280(int32(uintptr(unsafe.Pointer(v16))), int32(uintptr(unsafe.Pointer(v4))))
	nox_xxx_wnd_46B280(int32(uintptr(unsafe.Pointer(v12))), int32(uintptr(unsafe.Pointer(v4))))
	*(*uint32)(unsafe.Add(unsafe.Pointer(v7), unsafe.Sizeof(uint32(0))*9)) = uint32(uintptr(unsafe.Pointer(v8)))
	*(*uint32)(unsafe.Add(unsafe.Pointer(v7), unsafe.Sizeof(uint32(0))*7)) = uint32(uintptr(unsafe.Pointer(v16)))
	*(*uint32)(unsafe.Add(unsafe.Pointer(v7), unsafe.Sizeof(uint32(0))*8)) = uint32(uintptr(unsafe.Pointer(v12)))
	sub_456500()
	v9 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045684)))).ChildByID(0x2908)))
	if noxflags.HasGame(noxflags.GameModeChat) {
		v10 = strMan.GetStringInFile("Title1", "C:\\NoxPost\\src\\client\\Gui\\ServOpts\\playrlst.c")
		(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v9)))))).Func94(asWindowEvent(0x4001, int32(uintptr(unsafe.Pointer(v10))), -1))
	}
	return int32(dword_5d4594_1045684)
}
func sub_4567C0(a1 int32, a2 int32, a3 *int32, a4 int32) int32 {
	var (
		v3  *byte
		v4  int32
		v5  *uint32
		v6  int32
		v7  *uint32
		v8  int32
		v9  *int32
		v10 *int32
		i   int32
		v12 *libc.WChar
		v13 *byte
		v14 *byte
		v15 *uint32
		v16 int32
		v17 *uint32
		v18 *uint32
		v20 *uint32
		v21 int32
		v22 *byte
		v23 *libc.WChar
		v24 *uint32
		v25 int32
		v26 *byte
		v27 *libc.WChar
		v28 *libc.WChar
		v29 *byte
		v30 [56]libc.WChar
	)
	if a2 != 0x4007 {
		if a2 != 16400 {
			return 0
		}
		if (*nox_window)(unsafe.Pointer(a3)).ID() == 0x2906 {
			v3 = sub_4165B0()
			if (*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a3)))))).Func94(asWindowEvent(0x4014, 0, 0)) < 0 {
				nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045688))))), 0)
				nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045692))))), 0)
			} else if noxflags.HasGame(noxflags.GameFlag16) || *(*byte)(unsafe.Add(unsafe.Pointer(v3), 53)) < 0 {
				nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045688))))), 0)
				nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045692))))), 0)
			} else {
				if noxflags.HasGame(noxflags.GameHost) {
					nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045692))))), 1)
				}
				if noxflags.HasGame(noxflags.GameModeChat) || *memmap.PtrUint32(6112660, 1045696) == 0 {
					nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045688))))), 1)
				} else {
					nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045688))))), 0)
				}
			}
			if noxflags.HasGame(noxflags.GameHost) && nox_common_getEngineFlag(nox_engine_flag(NOX_ENGINE_FLAG_DISABLE_GRAPHICS_RENDERING)) {
				nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045688))))), 0)
			}
		}
	}
	v4 = (*nox_window)(unsafe.Pointer(a3)).ID()
	clientPlaySoundSpecial(766, 100)
	if v4 > 0x290B {
		if v4 == 0x290D {
			v28 = strMan.GetStringInFile("NewName", "C:\\NoxPost\\src\\client\\Gui\\ServOpts\\playrlst.c")
			v27 = strMan.GetStringInFile("Rename", "C:\\NoxPost\\src\\client\\Gui\\ServOpts\\playrlst.c")
			NewDialogWindow((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045684))))), (*libc.WChar)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v27)))))), (*libc.WChar)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v28)))))), 163, nil, nil)
		}
		return 0
	}
	if v4 == 0x290B {
		v24 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045684)))).ChildByID(0x2906)))
		v25 = (*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v24)))))).Func94(asWindowEvent(0x4014, 0, 0))
		sub_456D00(v25, &v30[0])
		v26 = sub_418A40(&v30[0])
		if v26 != nil {
			sub_456BB0(int32(uintptr(unsafe.Pointer(v26))))
			return 0
		}
		return 0
	}
	if v4 == 4001 {
		v20 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045684)))).ChildByID(0x2906)))
		v21 = (*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v20)))))).Func94(asWindowEvent(0x4014, 0, 0))
		sub_456D00(v21, &v30[0])
		v22 = sub_418A40(&v30[0])
		if v22 != nil {
			v23 = (*libc.WChar)(unsafe.Pointer(uintptr(sub_449E60(-88))))
			if sub_4190F0(v23) == 0 {
				nox_xxx_teamRenameMB_418CD0((*libc.WChar)(unsafe.Pointer(v22)), v23)
				return 0
			}
		}
		return 0
	}
	if v4 != 0x2907 {
		return 0
	}
	v5 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045684)))).ChildByID(0x2906)))
	v6 = (*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v5)))))).Func94(asWindowEvent(0x4014, 0, 0))
	if v6 >= 0 {
		sub_456D00(v6, &v30[0])
		v29 = sub_418A40(&v30[0])
		if v29 != nil {
			v7 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045684)))).ChildByID(0x2905)))
			v8 = int32(uintptr(unsafe.Pointer(v7)))
			v9 = (*int32)(unsafe.Pointer(uintptr((*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v7)))))).Func94(asWindowEvent(0x4014, 0, 0)))))
			v10 = v9
			for i = *v9; i >= 0; v10 = (*int32)(unsafe.Add(unsafe.Pointer(v10), unsafe.Sizeof(int32(0))*1)) {
				v12 = (*libc.WChar)(unsafe.Pointer(uintptr((*nox_window)(unsafe.Pointer(uintptr(v8))).Func94(asWindowEvent(0x4016, i, 0)))))
				v13 = nox_xxx_playerByName_4170D0(v12)
				v14 = v13
				if v13 != nil {
					if (*(*byte)(unsafe.Add(unsafe.Pointer(v13), 3680))&1) == 0 && (*(*byte)(unsafe.Add(unsafe.Pointer(v13), 4))&1) == 0 {
						v15 = nox_xxx_objGetTeamByNetCode_418C80(int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v13))), unsafe.Sizeof(uint32(0))*515)))))
						v16 = int32(uintptr(unsafe.Pointer(v15)))
						if v15 != nil {
							if nox_xxx_servObjectHasTeam_419130(int32(uintptr(unsafe.Pointer(v15)))) != 0 {
								sub_4196D0(v16, int32(uintptr(unsafe.Pointer(v29))), int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v14))), unsafe.Sizeof(uint32(0))*515)))), 1)
							} else {
								nox_xxx_createAtImpl_4191D0(uint8(*(*byte)(unsafe.Add(unsafe.Pointer(v29), 57))), unsafe.Pointer(uintptr(v16)), 1, int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v14))), unsafe.Sizeof(uint32(0))*515)))), 1)
							}
						}
					}
				}
				i = *(*int32)(unsafe.Add(unsafe.Pointer(v10), unsafe.Sizeof(int32(0))*1))
			}
		}
	}
	v17 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045684)))).ChildByID(0x2906)))
	(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v17)))))).Func94(asWindowEvent(0x4013, -1, 0))
	v18 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045684)))).ChildByID(0x2905)))
	(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v18)))))).Func94(asWindowEvent(0x4013, -1, 0))
	return 0
}
func sub_457010(a1 int32, a2 *libc.WChar) int32 {
	var (
		v2     *byte
		result int32
		v4     *uint32
		v5     int32
		v6     uint8
		v7     uint8
		v8     *libc.WChar
		v9     [56]libc.WChar
	)
	v2 = sub_4165B0()
	result = int32(dword_5d4594_1045684)
	if dword_5d4594_1045684 != 0 {
		v4 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045684)))).ChildByID(0x2906)))
		v5 = (*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v4)))))).Func94(asWindowEvent(0x4014, 0, 0))
		nox_wcscpy(&v9[0], a2)
		if noxflags.HasGame(noxflags.GameModeCTF|noxflags.GameModeFlagBall) || *(*byte)(unsafe.Add(unsafe.Pointer(v2), 52))&96 != 0 {
			v6 = *(*uint8)(unsafe.Pointer(uintptr(a1 + 57)))
			if int32(v6) < 3 {
				if int32(v6) == 1 {
					v8 = strMan.GetStringInFile("RedFlag", "C:\\NoxPost\\src\\client\\Gui\\ServOpts\\playrlst.c")
				} else {
					v8 = strMan.GetStringInFile("BlueFlag", "C:\\NoxPost\\src\\client\\Gui\\ServOpts\\playrlst.c")
				}
				nox_wcscat(&v9[0], v8)
			}
		}
		(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v4)))))).Func94(asWindowEvent(0x400E, v5, 0))
		(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v4)))))).Func94(asWindowEvent(0x4012, v5, 0))
		v7 = sub_457120(a1)
		(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v4)))))).Func94(asWindowEvent(0x400D, int32(uintptr(unsafe.Pointer(&v9[0]))), int32(v7)))
		result = (*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v4)))))).Func94(asWindowEvent(0x4013, v5, 0))
	}
	return result
}
func sub_457230(a1 *libc.WChar) *byte {
	var (
		v1     *byte
		result *byte
		v3     *uint32
		v4     *byte
		v5     int8
		v6     *uint32
		v7     *libc.WChar
		v8     [56]libc.WChar
	)
	v1 = sub_4165B0()
	result = *(**byte)(unsafe.Pointer(&dword_5d4594_1045684))
	if dword_5d4594_1045684 != 0 {
		v3 = (*uint32)(alloc.Calloc(1, 72))
		result = sub_418A40(a1)
		v4 = result
		if result != nil {
			*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*15)) = uint32(uint8(*(*byte)(unsafe.Add(unsafe.Pointer(result), 57))))
			*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*17)) = uint32(uint8(*(*byte)(unsafe.Add(unsafe.Pointer(result), 56))))
			*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(v3))), 64))) = sub_457120(int32(uintptr(unsafe.Pointer(result))))
			sub_425770(unsafe.Pointer(v3))
			nox_common_list_append_4258E0((*nox_list_item_t)(unsafe.Pointer(uintptr(int32(uintptr(memmap.PtrOff(6112660, 1045668)))))), (*nox_list_item_t)(unsafe.Pointer(v3)))
			nox_wcscpy((*libc.WChar)(unsafe.Add(unsafe.Pointer((*libc.WChar)(unsafe.Pointer(v3))), unsafe.Sizeof(libc.WChar(0))*6)), a1)
			nox_wcscpy(&v8[0], a1)
			if noxflags.HasGame(noxflags.GameModeCTF|noxflags.GameModeFlagBall) || *(*byte)(unsafe.Add(unsafe.Pointer(v1), 52))&96 != 0 {
				v5 = int8(*(*byte)(unsafe.Add(unsafe.Pointer(v4), 57)))
				if int32(uint8(v5)) < 3 {
					if int32(v5) == 1 {
						v7 = strMan.GetStringInFile("RedFlag", "C:\\NoxPost\\src\\client\\Gui\\ServOpts\\playrlst.c")
					} else {
						v7 = strMan.GetStringInFile("BlueFlag", "C:\\NoxPost\\src\\client\\Gui\\ServOpts\\playrlst.c")
					}
					nox_wcscat(&v8[0], v7)
				}
			}
			v6 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045684)))).ChildByID(0x2906)))
			result = (*byte)(unsafe.Pointer(uintptr((*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v6)))))).Func94(asWindowEvent(0x400D, int32(uintptr(unsafe.Pointer(&v8[0]))), int32(*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(v3))), 64)))))))))
		}
	}
	return result
}
