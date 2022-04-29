package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	noxflags "github.com/noxworld-dev/opennox/v1/common/flags"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"image"
	"unsafe"
)

func nox_xxx_spellPutInBox_45DEB0(a1 *int32, a2 int32, a3 int32, a4 int32) int32 {
	var (
		v4 int32
		v5 *libc.WChar
		v6 *uint32
		v7 int32
	)
	v4 = nox_xxx_spellBoxPointToWnd_45DE60(int32(uintptr(unsafe.Pointer(a1))), a3, a4)
	if v4 >= 0 {
		if a1 != memmap.PtrInt32(6112660, 1047940) {
		LABEL_8:
			nox_xxx_spellKeyPackSetSpell_45DC40(int32(uintptr(unsafe.Pointer(a1))), a2, v4)
			return 1
		}
		if nox_xxx_spellCanUseInTrap_424BF0(a2) {
			v6 = (*uint32)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*51)))))
			v7 = 0
			for *v6 != uint32(a2) {
				v7++
				v6 = (*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*2))
				if v7 >= 3 {
					goto LABEL_8
				}
			}
			v5 = strMan.GetStringInFile("OneSpellPerTrap", "C:\\NoxPost\\src\\Client\\Gui\\guispell.c")
		} else {
			v5 = strMan.GetStringInFile("RestrictedTrapSpell", "C:\\NoxPost\\src\\Client\\Gui\\guispell.c")
		}
		nox_xxx_printCentered_445490(v5)
		clientPlaySoundSpecial(925, 100)
	}
	return 0
}
func nox_client_buildTrap_45E040() {
	var (
		v0     **uint32
		v1     *uint32
		v2     int32
		i      int32
		result *uint32
		v5     *libc.WChar
		v6     *int32
		v7     int32
		v8     int8
		v9     [5]int32
	)
	v0 = *(***uint32)(memmap.PtrOff(6112660, 1048144))
	v1 = *(**uint32)(memmap.PtrOff(6112660, 1048144))
	v2 = 0
	for i = 0; i < 3; i++ {
		if *v1 != 0 {
			break
		}
		v1 = (*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*2))
	}
	if i == 3 {
		if *memmap.PtrUint32(0x8531A0, 2576) != 0 {
			if *(*uint32)(unsafe.Pointer(uintptr(*memmap.PtrUint32(0x8531A0, 2576) + 3832))) != 0 {
				v5 = strMan.GetStringInFile("TrapError", "C:\\NoxPost\\src\\Client\\Gui\\guispell.c")
				nox_xxx_printCentered_445490(v5)
				clientPlaySoundSpecial(925, 100)
			}
		}
	} else {
		*memmap.PtrUint8(6112660, 1049488) = 0
		v6 = &v9[0]
		v7 = 3
		for {
			result = *v0
			if *v0 != nil {
				*v6 = int32(uintptr(unsafe.Pointer(result)))
				v2++
				v6 = (*int32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(int32(0))*1))
			}
			v0 = (**uint32)(unsafe.Add(unsafe.Pointer(v0), unsafe.Sizeof((*uint32)(nil))*2))
			v7--
			if v7 == 0 {
				break
			}
		}
		if v2 < 5 {
			v8 = int8(*memmap.PtrUint8(6112660, 1047924))
			v9[v2] = 34
			nox_xxx_clientSendSpell_45DB20((*byte)(unsafe.Pointer(&v9[0])), v2+1, v8)
			*memmap.PtrUint32(6112660, 1047916) = 0
			*memmap.PtrUint32(6112660, 1049480) = 0
		}
	}
}
func nox_xxx_quickBarCreate_45E190() int32 {
	var (
		v0  int32
		v1  int32
		v2  int32
		v4  *uint8
		v5  *byte
		v6  *byte
		v7  *byte
		v8  *byte
		v9  *libc.WChar
		v10 *libc.WChar
		v11 *byte
		v12 *byte
		v13 *byte
		v14 *byte
		v15 *libc.WChar
		v16 *libc.WChar
		v17 *byte
		v18 *byte
		v19 *byte
		v20 *byte
		v21 *byte
		v22 int32
		v23 *uint32
		v24 *libc.WChar
		v25 *byte
		v26 *libc.WChar
		v27 *libc.WChar
		v28 *byte
		v29 *libc.WChar
		v30 *byte
		v31 *byte
		v32 *byte
		v33 *byte
		v34 *byte
		v35 *byte
		v36 *libc.WChar
		v37 *libc.WChar
		v38 *byte
		v39 *libc.WChar
		v40 *libc.WChar
		v41 *byte
		v42 *libc.WChar
		v43 *libc.WChar
		v44 *byte
		v45 *libc.WChar
		v46 *byte
		v47 *uint32
		v48 int32
		v49 int32
		v50 *int32
		v51 int32
		v52 int32
		v53 *uint32
		v54 *byte
		v55 *byte
		v56 *byte
		v57 *byte
		v58 *uint32
		v59 *uint32
		v60 int32
		v61 bool
		j   int32
		v63 int32
		v64 uint16
		v65 int32
		v66 int32
		v67 int32
		v68 int32
		v69 int32
		i   int32
		v71 *uint8
		v72 [32]byte
	)
	v0 = 0
	*memmap.PtrUint32(6112660, 1047916) = 0
	*memmap.PtrUint8(6112660, 1047920) = 0
	sub_416170(5)
	*memmap.PtrUint32(6112660, 1047924) = 0
	*memmap.PtrUint32(6112660, 1047928) = 0
	*memmap.PtrUint32(6112660, 1049480) = 0
	*memmap.PtrUint8(6112660, 1049488) = 0
	v68 = nox_xxx_guiFontHeightMB_43F320(nil)
	dword_5d4594_1047548 = uint32((nox_win_width - 320) / 2)
	dword_5d4594_1047552 = uint32(nox_win_height - 74)
	*memmap.PtrUint32(6112660, 1049684) = uint32(uintptr(guiFontPtrByName(CString("small"))))
	*memmap.PtrUint32(0x587000, 133656) = dword_5d4594_1047548
	v1 = int32(dword_5d4594_1047548 + 69)
	*memmap.PtrUint32(0x587000, 133660) = dword_5d4594_1047552 - 17
	v2 = int32(dword_5d4594_1047552 + 32)
	*memmap.PtrUint32(0x587000, 133664) = dword_5d4594_1047548 + 320
	*memmap.PtrUint32(0x587000, 133668) = uint32(nox_win_height)
	if *memmap.PtrUint32(0x8531A0, 2576) == 0 {
		return 0
	}
	if int32(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(0x8531A0, 2576) + 2251)))) != 0 {
		nox_xxx_quickBarInitWindow_4601F0(*(*int32)(unsafe.Pointer(&nox_xxx_aClosewoodengat_587000_133480)), int32(dword_5d4594_1047548+69), int32(dword_5d4594_1047552+32), 5, 0, int32(cgoFuncAddr(nox_xxx_quickBarWnd_45EF50)), int32(cgoFuncAddr(nox_xxx_quickBarWarriorDraw_45FDE0)))
	} else {
		nox_xxx_quickBarInitWindow_4601F0(*(*int32)(unsafe.Pointer(&nox_xxx_aClosewoodengat_587000_133480)), int32(dword_5d4594_1047548+69), int32(dword_5d4594_1047552+32), 5, 0, int32(cgoFuncAddr(nox_xxx_quickBarWnd_45EF50)), int32(cgoFuncAddr(nox_xxx_quickBarDrawFn_45FBD0)))
	}
	v4 = (*uint8)(memmap.PtrOff(6112660, 1048964))
	for {
		v2 -= 60
		nox_xxx_quickBarInitWindow_4601F0(int32(uintptr(unsafe.Pointer(v4))), v1, v2, 5, 0, int32(cgoFuncAddr(nox_xxx_quickBarWnd_45EF50)), int32(cgoFuncAddr(nox_xxx_quickBarWarriorDraw_45FDE0)))
		(*nox_window)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v4))), unsafe.Sizeof(uint32(0))*52)))))).Hide()
		*(*uint8)(unsafe.Add(unsafe.Pointer(v4), 200)) = 0
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v4))), unsafe.Sizeof(uint32(0))*51))) = uint32(uintptr(unsafe.Pointer(v4)))
		v4 = (*uint8)(unsafe.Add(unsafe.Pointer(v4), -256))
		if int32(uintptr(unsafe.Pointer(v4))) < int32(uintptr(memmap.PtrOff(6112660, 1048196))) {
			break
		}
	}
	if int32(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(0x8531A0, 2576) + 2251)))) != 0 {
		dword_5d4594_1049504 = uint32(uintptr(unsafe.Pointer(nox_window_new(nil, 1160, int32(dword_5d4594_1047548+260), *(*int32)(unsafe.Pointer(&dword_5d4594_1047552)), 45, 66, nil))))
		nox_xxx_wndSetOffsetMB_46AE40(*(*int32)(unsafe.Pointer(&dword_5d4594_1049504)), -263, 0)
		dword_5d4594_1049536 = uint32(nox_win_height - 74)
		(*(*int32)(unsafe.Pointer(&dword_5d4594_1049504))).setFunc93(nox_xxx_quickbar_45F8D0)
		dword_5d4594_1049520 = uint32(uintptr(unsafe.Pointer(nox_window_new((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1049504))))), 1032, 9, 33, 32, 32, nil))))
		(*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1049520)))).SetAllFuncs(func(arg1 int32, arg2 int32, arg3 int32, arg4 int32) int32 {
			return nox_xxx_quickbarTrapButtonProc_45F7A0((*uint32)(unsafe.Pointer(uintptr(arg1))), uint32(arg2), uint32(arg3))
		}, func(arg1 *nox_window, arg2 unsafe.Pointer) int32 {
			return nox_xxx_quickbarDrawFn_460000()
		}, nil)
		dword_5d4594_1049500 = uint32(uintptr(unsafe.Pointer(nox_window_new((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1049504))))), 1160, 0, 19, 12, 12, nil))))
		(*(*int32)(unsafe.Pointer(&dword_5d4594_1049500))).setFunc93(func(arg1 int32, arg2 int32, arg3 int32, arg4 int32) int32 {
			return nox_xxx_quickbarTrapProc_45FB90(arg1, uint32(arg2), arg3, arg4)
		})
		nox_xxx_wndSetOffsetMB_46AE40(*(*int32)(unsafe.Pointer(&dword_5d4594_1049500)), -265, -23)
		v5 = (*byte)(unsafe.Pointer(nox_xxx_gLoadImg("QuickBarTrapButton")))
		nox_xxx_wndSetIcon_46AE60(*(*int32)(unsafe.Pointer(&dword_5d4594_1049500)), int32(uintptr(unsafe.Pointer(v5))))
		v6 = (*byte)(unsafe.Pointer(nox_xxx_gLoadImg("QuickBarTrapButtonLit")))
		nox_xxx_wndSetIconLit_46AEA0(*(*int32)(unsafe.Pointer(&dword_5d4594_1049500)), int32(uintptr(unsafe.Pointer(v6))))
	} else {
		dword_5d4594_1049504 = uint32(uintptr(unsafe.Pointer(nox_window_new(nil, 1672, int32(dword_5d4594_1047548+260), *(*int32)(unsafe.Pointer(&dword_5d4594_1047552)), 45, 66, nil))))
		nox_xxx_wndSetOffsetMB_46AE40(*(*int32)(unsafe.Pointer(&dword_5d4594_1049504)), -263, 0)
		dword_5d4594_1049536 = uint32(nox_win_height - 74)
		(*(*int32)(unsafe.Pointer(&dword_5d4594_1049504))).setFunc93(nox_xxx_quickbar_45F8D0)
	}
	if int32(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(0x8531A0, 2576) + 2251)))) != 0 {
		if int32(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(0x8531A0, 2576) + 2251)))) == 1 {
			if (*memmap.PtrUint32(0x8531A0, 2576) == 0 || *(*uint32)(unsafe.Pointer(uintptr(*memmap.PtrUint32(0x8531A0, 2576) + 3832))) == 0) && (!noxflags.HasGame(noxflags.GameOnline) || noxflags.HasGame(noxflags.GameModeQuest) || nox_xxx_isQuest_4D6F50() != 0 || sub_4D6F70() != 0) {
				v17 = (*byte)(unsafe.Pointer(nox_xxx_gLoadImg("QuickBarWarriorRight")))
				nox_xxx_wndSetIcon_46AE60(*(*int32)(unsafe.Pointer(&dword_5d4594_1049504)), int32(uintptr(unsafe.Pointer(v17))))
				v18 = (*byte)(unsafe.Pointer(nox_xxx_gLoadImg("QuickBarWarriorRight")))
				nox_xxx_wndSetIconLit_46AEA0(*(*int32)(unsafe.Pointer(&dword_5d4594_1049504)), int32(uintptr(unsafe.Pointer(v18))))
				nox_xxx_wndClearFlag_46AD80(*(*int32)(unsafe.Pointer(&dword_5d4594_1049520)), 8)
				nox_xxx_wndClearFlag_46AD80(*(*int32)(unsafe.Pointer(&dword_5d4594_1049500)), 8)
				(*(*int32)(unsafe.Pointer(&dword_5d4594_1049500))).setDraw(func(arg1 int32, arg2 int32) int32 {
					return nox_xxx_quickbarButtonBookDraw_45EF30()
				})
			} else {
				v13 = (*byte)(unsafe.Pointer(nox_xxx_gLoadImg("QuickBarTrap")))
				nox_xxx_wndSetIcon_46AE60(*(*int32)(unsafe.Pointer(&dword_5d4594_1049504)), int32(uintptr(unsafe.Pointer(v13))))
				v14 = (*byte)(unsafe.Pointer(nox_xxx_gLoadImg("QuickBarTrapHit")))
				nox_xxx_wndSetIconLit_46AEA0(*(*int32)(unsafe.Pointer(&dword_5d4594_1049504)), int32(uintptr(unsafe.Pointer(v14))))
				v15 = strMan.GetStringInFile("ToolTipLayTrap", "C:\\NoxPost\\src\\Client\\Gui\\guispell.c")
				nox_xxx_wndWddSetTooltip_46B000((*nox_window_data)(unsafe.Pointer(uintptr(dword_5d4594_1049520+36))), v15)
				nox_xxx_wnd_46AD60(*(*int32)(unsafe.Pointer(&dword_5d4594_1049520)), 8)
				v16 = strMan.GetStringInFile("ToolTipTrapConstruct", "C:\\NoxPost\\src\\Client\\Gui\\guispell.c")
				nox_xxx_wndWddSetTooltip_46B000((*nox_window_data)(unsafe.Pointer(uintptr(dword_5d4594_1049500+36))), v16)
				nox_xxx_wnd_46AD60(*(*int32)(unsafe.Pointer(&dword_5d4594_1049500)), 8)
			}
		} else if int32(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(0x8531A0, 2576) + 2251)))) == 2 {
			if (*memmap.PtrUint32(0x8531A0, 2576) == 0 || *(*uint32)(unsafe.Pointer(uintptr(*memmap.PtrUint32(0x8531A0, 2576) + 3832))) == 0) && (!noxflags.HasGame(noxflags.GameOnline) || noxflags.HasGame(noxflags.GameModeQuest) || nox_xxx_isQuest_4D6F50() != 0 || sub_4D6F70() != 0) {
				v11 = (*byte)(unsafe.Pointer(nox_xxx_gLoadImg("QuickBarWarriorRight")))
				nox_xxx_wndSetIcon_46AE60(*(*int32)(unsafe.Pointer(&dword_5d4594_1049504)), int32(uintptr(unsafe.Pointer(v11))))
				v12 = (*byte)(unsafe.Pointer(nox_xxx_gLoadImg("QuickBarWarriorRight")))
				nox_xxx_wndSetIconLit_46AEA0(*(*int32)(unsafe.Pointer(&dword_5d4594_1049504)), int32(uintptr(unsafe.Pointer(v12))))
				nox_xxx_wndClearFlag_46AD80(*(*int32)(unsafe.Pointer(&dword_5d4594_1049520)), 8)
				nox_xxx_wndClearFlag_46AD80(*(*int32)(unsafe.Pointer(&dword_5d4594_1049500)), 8)
				(*(*int32)(unsafe.Pointer(&dword_5d4594_1049500))).setDraw(func(arg1 int32, arg2 int32) int32 {
					return nox_xxx_quickbarButtonBookDraw_45EF30()
				})
			} else {
				v7 = (*byte)(unsafe.Pointer(nox_xxx_gLoadImg("QuickBarBomber")))
				nox_xxx_wndSetIcon_46AE60(*(*int32)(unsafe.Pointer(&dword_5d4594_1049504)), int32(uintptr(unsafe.Pointer(v7))))
				v8 = (*byte)(unsafe.Pointer(nox_xxx_gLoadImg("QuickBarBomberHit")))
				nox_xxx_wndSetIconLit_46AEA0(*(*int32)(unsafe.Pointer(&dword_5d4594_1049504)), int32(uintptr(unsafe.Pointer(v8))))
				v9 = strMan.GetStringInFile("ToolTipSummonBomber", "C:\\NoxPost\\src\\Client\\Gui\\guispell.c")
				nox_xxx_wndWddSetTooltip_46B000((*nox_window_data)(unsafe.Pointer(uintptr(dword_5d4594_1049520+36))), v9)
				nox_xxx_wnd_46AD60(*(*int32)(unsafe.Pointer(&dword_5d4594_1049520)), 8)
				v10 = strMan.GetStringInFile("ToolTipTrapConstruct", "C:\\NoxPost\\src\\Client\\Gui\\guispell.c")
				nox_xxx_wndWddSetTooltip_46B000((*nox_window_data)(unsafe.Pointer(uintptr(dword_5d4594_1049500+36))), v10)
				nox_xxx_wnd_46AD60(*(*int32)(unsafe.Pointer(&dword_5d4594_1049500)), 8)
			}
		}
	} else {
		v19 = (*byte)(unsafe.Pointer(nox_xxx_gLoadImg("QuickBarWarriorRight")))
		nox_xxx_wndSetIcon_46AE60(*(*int32)(unsafe.Pointer(&dword_5d4594_1049504)), int32(uintptr(unsafe.Pointer(v19))))
		v20 = (*byte)(unsafe.Pointer(nox_xxx_gLoadImg("QuickBarWarriorRight")))
		nox_xxx_wndSetIconLit_46AEA0(*(*int32)(unsafe.Pointer(&dword_5d4594_1049504)), int32(uintptr(unsafe.Pointer(v20))))
	}
	if int32(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(0x8531A0, 2576) + 2251)))) != 0 {
		nox_xxx_quickBarInitWindow_4601F0(int32(uintptr(memmap.PtrOff(6112660, 1047940))), int32(dword_5d4594_1047548+122), int32(dword_5d4594_1047552-17), 3, 21, int32(cgoFuncAddr(nox_xxx_quickBarWnd_45EF50)), int32(cgoFuncAddr(nox_xxx_quickBarWarriorDraw_45FDE0)))
		v21 = (*byte)(unsafe.Pointer(nox_xxx_gLoadImg("QuickBarTrapTray")))
		nox_xxx_wndSetIcon_46AE60(*memmap.PtrInt32(6112660, 1048148), int32(uintptr(unsafe.Pointer(v21))))
		nox_xxx_wndSetOffsetMB_46AE40(*memmap.PtrInt32(6112660, 1048148), -40, -20)
		v22 = int32(*memmap.PtrUint32(6112660, 1048192))
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v22))), 0)) = uint8(int8(int32(*memmap.PtrUint8(6112660, 1048192)) | 1))
		*memmap.PtrUint32(6112660, 1048192) = uint32(v22)
		(*nox_window)(unsafe.Pointer(uintptr(*memmap.PtrInt32(6112660, 1048148)))).Hide()
		dword_5d4594_1049484 = 0
		v23 = (*uint32)(unsafe.Pointer(nox_window_new((*nox_window)(unsafe.Pointer(uintptr(*memmap.PtrInt32(6112660, 1048148)))), 1032, 20, -7, 110, v68, nil)))
		(*nox_window)(unsafe.Pointer(v23)).SetAllFuncs(nil, func(arg1 *nox_window, arg2 unsafe.Pointer) int32 {
			return nox_xxx_quickbarDraw_45FAC0((*uint32)(unsafe.Pointer(arg1)))
		}, nil)
		v24 = (*libc.WChar)(unsafe.Pointer(nox_window_new((*nox_window)(unsafe.Pointer(uintptr(*memmap.PtrInt32(6112660, 1048148)))), 1032, 15, 12, 10, 14, nil)))
		nox_xxx_wndSetIcon_46AE60(int32(uintptr(unsafe.Pointer(v24))), 0)
		v25 = (*byte)(unsafe.Pointer(nox_xxx_gLoadImg("QuickBarTrapTrayUpLit")))
		nox_xxx_wndSetIconLit_46AEA0(int32(uintptr(unsafe.Pointer(v24))), int32(uintptr(unsafe.Pointer(v25))))
		(*nox_window)(unsafe.Pointer(v24)).SetAllFuncs(func(arg1 int32, arg2 int32, arg3 int32, arg4 int32) int32 {
			return nox_xxx_quickbarTrapUpDownProc_45F630(arg1, uint32(arg2))
		}, func(arg1 *nox_window, arg2 unsafe.Pointer) int32 {
			return nox_xxx_quickbarTrapUpDownDraw_45F6F0((*uint32)(unsafe.Pointer(arg1)))
		}, nil)
		nox_xxx_wndSetOffsetMB_46AE40(int32(uintptr(unsafe.Pointer(v24))), -55, -32)
		v26 = strMan.GetStringInFile("ToolTipPrevTrap", "C:\\NoxPost\\src\\Client\\Gui\\guispell.c")
		nox_xxx_wndWddSetTooltip_46B000((*nox_window_data)(unsafe.Pointer((*libc.WChar)(unsafe.Add(unsafe.Pointer(v24), unsafe.Sizeof(libc.WChar(0))*18)))), v26)
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v24))), unsafe.Sizeof(uint32(0))*92))) = 3
		v27 = (*libc.WChar)(unsafe.Pointer(nox_window_new((*nox_window)(unsafe.Pointer(uintptr(*memmap.PtrInt32(6112660, 1048148)))), 1032, 15, 32, 10, 14, nil)))
		nox_xxx_wndSetIcon_46AE60(int32(uintptr(unsafe.Pointer(v27))), 0)
		v28 = (*byte)(unsafe.Pointer(nox_xxx_gLoadImg("QuickBarTrapTrayDownLit")))
		nox_xxx_wndSetIconLit_46AEA0(int32(uintptr(unsafe.Pointer(v27))), int32(uintptr(unsafe.Pointer(v28))))
		(*nox_window)(unsafe.Pointer(v27)).SetAllFuncs(func(arg1 int32, arg2 int32, arg3 int32, arg4 int32) int32 {
			return nox_xxx_quickbarTrapUpDownProc_45F630(arg1, uint32(arg2))
		}, func(arg1 *nox_window, arg2 unsafe.Pointer) int32 {
			return nox_xxx_quickbarTrapUpDownDraw_45F6F0((*uint32)(unsafe.Pointer(arg1)))
		}, nil)
		nox_xxx_wndSetOffsetMB_46AE40(int32(uintptr(unsafe.Pointer(v27))), -55, -52)
		v29 = strMan.GetStringInFile("ToolTipNextTrap", "C:\\NoxPost\\src\\Client\\Gui\\guispell.c")
		nox_xxx_wndWddSetTooltip_46B000((*nox_window_data)(unsafe.Pointer((*libc.WChar)(unsafe.Add(unsafe.Pointer(v27), unsafe.Sizeof(libc.WChar(0))*18)))), v29)
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v27))), unsafe.Sizeof(uint32(0))*92))) = 4
		dword_5d4594_1049508 = uint32(uintptr(unsafe.Pointer(nox_window_new(nil, 1032, int32(dword_5d4594_1047548-1), int32(dword_5d4594_1047552+26), 61, 48, nil))))
		nox_xxx_wndSetOffsetMB_46AE40(*(*int32)(unsafe.Pointer(&dword_5d4594_1049508)), 1, -26)
		v30 = (*byte)(unsafe.Pointer(nox_xxx_gLoadImg("QuickBarSpellSetBase")))
		nox_xxx_wndSetIcon_46AE60(*(*int32)(unsafe.Pointer(&dword_5d4594_1049508)), int32(uintptr(unsafe.Pointer(v30))))
		v31 = (*byte)(unsafe.Pointer(nox_xxx_gLoadImg("QuickBarSpellSetBase")))
		nox_xxx_wndSetIconLit_46AEA0(*(*int32)(unsafe.Pointer(&dword_5d4594_1049508)), int32(uintptr(unsafe.Pointer(v31))))
		*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1049508 + 368))) = 5
		(*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1049508)))).SetAllFuncs(nox_xxx_quickbar_45F8D0, func(arg1 *nox_window, arg2 unsafe.Pointer) int32 {
			return nox_xxx_quickbarTrapUpDownDraw_45F6F0((*uint32)(unsafe.Pointer(arg1)))
		}, nil)
	} else {
		dword_5d4594_1049508 = uint32(uintptr(unsafe.Pointer(nox_window_new(nil, 1032, int32(dword_5d4594_1047548-1), int32(dword_5d4594_1047552+26), 61, 48, nil))))
		nox_xxx_wndSetOffsetMB_46AE40(*(*int32)(unsafe.Pointer(&dword_5d4594_1049508)), 1, -26)
		v32 = (*byte)(unsafe.Pointer(nox_xxx_gLoadImg("QuickBarWarriorLeft")))
		nox_xxx_wndSetIcon_46AE60(*(*int32)(unsafe.Pointer(&dword_5d4594_1049508)), int32(uintptr(unsafe.Pointer(v32))))
		v33 = (*byte)(unsafe.Pointer(nox_xxx_gLoadImg("QuickBarWarriorLeft")))
		nox_xxx_wndSetIconLit_46AEA0(*(*int32)(unsafe.Pointer(&dword_5d4594_1049508)), int32(uintptr(unsafe.Pointer(v33))))
		(*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1049508)))).SetAllFuncs(nox_xxx_quickbar_45F8D0, func(arg1 *nox_window, arg2 unsafe.Pointer) int32 {
			return nox_xxx_quickbarTrapUpDownDraw_45F6F0((*uint32)(unsafe.Pointer(arg1)))
		}, nil)
	}
	dword_5d4594_1049524 = uint32(uintptr(unsafe.Pointer(nox_window_new((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1049508))))), 1160, 0, 9, 29, 30, nil))))
	v34 = (*byte)(unsafe.Pointer(nox_xxx_gLoadImg("SpellbookButton")))
	nox_xxx_wndSetIcon_46AE60(*(*int32)(unsafe.Pointer(&dword_5d4594_1049524)), int32(uintptr(unsafe.Pointer(v34))))
	v35 = (*byte)(unsafe.Pointer(nox_xxx_gLoadImg("SpellbookButtonLit")))
	nox_xxx_wndSetIconLit_46AEA0(*(*int32)(unsafe.Pointer(&dword_5d4594_1049524)), int32(uintptr(unsafe.Pointer(v35))))
	*memmap.PtrUint32(6112660, 1049528) = uint32(uintptr(unsafe.Pointer(nox_window_new((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1049524))))), 1064, 1, 2, 28, 28, nil))))
	(*nox_window)(unsafe.Pointer(*(**uint32)(memmap.PtrOff(6112660, 1049528)))).SetAllFuncs(func(arg1 int32, arg2 int32, arg3 int32, arg4 int32) int32 {
		return nox_xxx_quickbarButtonBookWnd_45F450(arg1, uint32(arg2))
	}, func(arg1 *nox_window, arg2 unsafe.Pointer) int32 {
		return nox_xxx_quickbarButtonBookDraw_45EF30()
	}, unsafe.Pointer(cgoFuncAddr(nox_xxx_quickbarButtonBook_45F3F0)))
	v36 = strMan.GetStringInFile("OpenSpellBookTT", "C:\\NoxPost\\src\\Client\\Gui\\guispell.c")
	nox_xxx_wndWddSetTooltip_46B000((*nox_window_data)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, 1049528)+36))), v36)
	if int32(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(0x8531A0, 2576) + 2251)))) != 0 {
		v37 = (*libc.WChar)(unsafe.Pointer(nox_window_new((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1049508))))), 1032, 30, 0, 15, 19, nil)))
		nox_xxx_wndSetOffsetMB_46AE40(int32(uintptr(unsafe.Pointer(v37))), -29, -26)
		nox_xxx_wndSetIcon_46AE60(int32(uintptr(unsafe.Pointer(v37))), 0)
		v38 = (*byte)(unsafe.Pointer(nox_xxx_gLoadImg("QuickBarSpellSetUpLit")))
		nox_xxx_wndSetIconLit_46AEA0(int32(uintptr(unsafe.Pointer(v37))), int32(uintptr(unsafe.Pointer(v38))))
		(*nox_window)(unsafe.Pointer(v37)).SetAllFuncs(func(arg1 int32, arg2 int32, arg3 int32, arg4 int32) int32 {
			return nox_xxx_quickbarTrapUpDownProc_45F630(arg1, uint32(arg2))
		}, func(arg1 *nox_window, arg2 unsafe.Pointer) int32 {
			return nox_xxx_quickbarTrapUpDownDraw_45F6F0((*uint32)(unsafe.Pointer(arg1)))
		}, nil)
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v37))), unsafe.Sizeof(uint32(0))*92))) = 0
		v39 = strMan.GetStringInFile("ToolTipPrevSpellSet", "C:\\NoxPost\\src\\Client\\Gui\\guispell.c")
		nox_xxx_wndWddSetTooltip_46B000((*nox_window_data)(unsafe.Pointer((*libc.WChar)(unsafe.Add(unsafe.Pointer(v37), unsafe.Sizeof(libc.WChar(0))*18)))), v39)
		v40 = (*libc.WChar)(unsafe.Pointer(nox_window_new((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1049508))))), 1032, 30, 29, 15, 19, nil)))
		nox_xxx_wndSetOffsetMB_46AE40(int32(uintptr(unsafe.Pointer(v40))), -29, -55)
		nox_xxx_wndSetIcon_46AE60(int32(uintptr(unsafe.Pointer(v40))), 0)
		v41 = (*byte)(unsafe.Pointer(nox_xxx_gLoadImg("QuickBarSpellSetDownLit")))
		nox_xxx_wndSetIconLit_46AEA0(int32(uintptr(unsafe.Pointer(v40))), int32(uintptr(unsafe.Pointer(v41))))
		(*nox_window)(unsafe.Pointer(v40)).SetAllFuncs(func(arg1 int32, arg2 int32, arg3 int32, arg4 int32) int32 {
			return nox_xxx_quickbarTrapUpDownProc_45F630(arg1, uint32(arg2))
		}, func(arg1 *nox_window, arg2 unsafe.Pointer) int32 {
			return nox_xxx_quickbarTrapUpDownDraw_45F6F0((*uint32)(unsafe.Pointer(arg1)))
		}, nil)
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v40))), unsafe.Sizeof(uint32(0))*92))) = 1
		v42 = strMan.GetStringInFile("ToolTipNextSpellSet", "C:\\NoxPost\\src\\Client\\Gui\\guispell.c")
		nox_xxx_wndWddSetTooltip_46B000((*nox_window_data)(unsafe.Pointer((*libc.WChar)(unsafe.Add(unsafe.Pointer(v40), unsafe.Sizeof(libc.WChar(0))*18)))), v42)
		v43 = (*libc.WChar)(unsafe.Pointer(nox_window_new((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1049508))))), 1032, 48, 16, 13, 17, nil)))
		nox_xxx_wndSetOffsetMB_46AE40(int32(uintptr(unsafe.Pointer(v43))), -47, -42)
		nox_xxx_wndSetIcon_46AE60(int32(uintptr(unsafe.Pointer(v43))), 0)
		v44 = (*byte)(unsafe.Pointer(nox_xxx_gLoadImg("QuickBarSpellSetMaxLit")))
		nox_xxx_wndSetIconLit_46AEA0(int32(uintptr(unsafe.Pointer(v43))), int32(uintptr(unsafe.Pointer(v44))))
		(*nox_window)(unsafe.Pointer(v43)).SetAllFuncs(func(arg1 int32, arg2 int32, arg3 int32, arg4 int32) int32 {
			return nox_xxx_quickbarTrapUpDownProc_45F630(arg1, uint32(arg2))
		}, func(arg1 *nox_window, arg2 unsafe.Pointer) int32 {
			return nox_xxx_quickbarTrapUpDownDraw_45F6F0((*uint32)(unsafe.Pointer(arg1)))
		}, nil)
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v43))), unsafe.Sizeof(uint32(0))*92))) = 2
		v45 = strMan.GetStringInFile("ToolTipAllSpellSets", "C:\\NoxPost\\src\\Client\\Gui\\guispell.c")
		nox_xxx_wndWddSetTooltip_46B000((*nox_window_data)(unsafe.Pointer((*libc.WChar)(unsafe.Add(unsafe.Pointer(v43), unsafe.Sizeof(libc.WChar(0))*18)))), v45)
		dword_5d4594_1049516 = uint32(uintptr(unsafe.Pointer(nox_window_new(nil, 1032, 0, 0, 1, 1, nil))))
		(*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1049516)))).SetAllFuncs(func(arg1 int32, arg2 int32, arg3 int32, arg4 int32) int32 {
			return sub_45EF40()
		}, func(arg1 *nox_window, arg2 unsafe.Pointer) int32 {
			return sub_45F8F0()
		}, nil)
		dword_5d4594_1049512 = uint32(uintptr(unsafe.Pointer(nox_window_new(nil, 1152, *(*int32)(unsafe.Pointer(&dword_5d4594_1047548)), *(*int32)(unsafe.Pointer(&dword_5d4594_1047552)), 2, 2, nil))))
		v46 = (*byte)(unsafe.Pointer(nox_xxx_gLoadImg("QuickBarTitle")))
		nox_xxx_wndSetIcon_46AE60(*(*int32)(unsafe.Pointer(&dword_5d4594_1049512)), int32(uintptr(unsafe.Pointer(v46))))
		v47 = (*uint32)(unsafe.Pointer(nox_window_new((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1049512))))), 8, 115, 6, 101, 14, nil)))
		(*nox_window)(unsafe.Pointer(v47)).SetAllFuncs(nil, func(arg1 *nox_window, arg2 unsafe.Pointer) int32 {
			return sub_45F9B0((*uint32)(unsafe.Pointer(arg1)))
		}, nil)
	}
	v64 = 0
	for {
		v65 = 0
		v66 = 0
		v69 = 5
		v48 = int32(v64) << 8
		v49 = int32(*memmap.PtrUint32(6112660, uint32(v48)+0xFFF54))
		v71 = (*uint8)(memmap.PtrOff(6112660, uint32(v48)+0xFFE84))
		v50 = memmap.PtrInt32(6112660, uint32(v48)+0xFFF6C)
		v51 = int32(*(*uint32)(unsafe.Pointer(uintptr(v49 + 16))) + 10)
		v52 = int32(*(*uint32)(unsafe.Pointer(uintptr(v49 + 20))) + 5)
		v67 = v51
		for i = v52; ; v52 = i {
			v53 = (*uint32)(unsafe.Pointer(nox_window_new(nil, 1160, v51, v52, 30, 10, nil)))
			v63 = v0 + 1
			if int32(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(0x8531A0, 2576) + 2251)))) != 0 {
				nox_sprintf(&v72[0], CString("QuickBarNugget%d"), v63)
				v54 = (*byte)(unsafe.Pointer(nox_xxx_gLoadImg(&v72[0])))
				nox_xxx_wndSetIcon_46AE60(int32(uintptr(unsafe.Pointer(v53))), int32(uintptr(unsafe.Pointer(v54))))
				*(*uint32)(unsafe.Pointer(&v72[libc.StrLen(&v72[0])])) = *memmap.PtrUint32(0x587000, 134996)
				v55 = (*byte)(unsafe.Pointer(nox_xxx_gLoadImg(&v72[0])))
				nox_xxx_wndSetIconLit_46AEA0(int32(uintptr(unsafe.Pointer(v53))), int32(uintptr(unsafe.Pointer(v55))))
				v51 = v67
			} else {
				nox_sprintf(&v72[0], CString("QuickBarWarriorNugget%d"), v63)
				v56 = (*byte)(unsafe.Pointer(nox_xxx_gLoadImg(&v72[0])))
				nox_xxx_wndSetIcon_46AE60(int32(uintptr(unsafe.Pointer(v53))), int32(uintptr(unsafe.Pointer(v56))))
				v57 = (*byte)(unsafe.Pointer(nox_xxx_gLoadImg(&v72[0])))
				nox_xxx_wndSetIconLit_46AEA0(int32(uintptr(unsafe.Pointer(v53))), int32(uintptr(unsafe.Pointer(v57))))
			}
			nox_xxx_wndSetOffsetMB_46AE40(int32(uintptr(unsafe.Pointer(v53))), int32(-70-v65), -23)
			(*nox_window)(unsafe.Pointer(v53)).SetAllFuncs(nox_xxx_quickbar_45F8D0, nil, nil)
			*v50 = int32(uintptr(unsafe.Pointer(v53)))
			nox_xxx_updateSpellIconDir_45DE10(v0, int32(uintptr(unsafe.Pointer(v71))))
			if int32(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(0x8531A0, 2576) + 2251)))) != 0 {
				v58 = (*uint32)(unsafe.Pointer(nox_window_new((*nox_window)(unsafe.Pointer(uintptr(*v50))), 1032, 12, 0, 10, 10, nil)))
				(*nox_window)(unsafe.Pointer(v58)).SetAllFuncs(func(arg1 int32, arg2 int32, arg3 int32, arg4 int32) int32 {
					return sub_45F520(arg1, uint32(arg2))
				}, func(arg1 *nox_window, arg2 unsafe.Pointer) int32 {
					return nox_xxx_quickbarButtonBookDraw_45EF30()
				}, unsafe.Pointer(cgoFuncAddr(sub_45F480)))
				*(*uint32)(unsafe.Add(unsafe.Pointer(v58), unsafe.Sizeof(uint32(0))*92)) = uint32(v0 | int32(v64)<<16)
			}
			if int32(v64) == 4 {
				v59 = (*uint32)(unsafe.Pointer(nox_window_new((*nox_window)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, uint32(v66)+0x10036C)))), 1032, 13, 39, 10, 10, nil)))
				(*nox_window)(unsafe.Pointer(v59)).SetAllFuncs(nil, func(arg1 *nox_window, arg2 unsafe.Pointer) int32 {
					return sub_45F5D0((*uint32)(unsafe.Pointer(arg1)))
				}, nil)
				*(*uint32)(unsafe.Add(unsafe.Pointer(v59), unsafe.Sizeof(uint32(0))*92)) = uint32(func() int32 {
					p := &v0
					x := *p
					*p++
					return x
				}())
			} else {
				(*nox_window)(unsafe.Pointer(uintptr(*v50))).Hide()
			}
			v50 = (*int32)(unsafe.Add(unsafe.Pointer(v50), unsafe.Sizeof(int32(0))*1))
			v60 = int32(*memmap.PtrUint32(0x587000, uint32(v66)+0x20970))
			v51 += v60
			v67 = v51
			v61 = v69 == 1
			v65 += v60
			v66 += 4
			v69--
			if v61 {
				break
			}
		}
		if int32(func() uint16 {
			p := &v64
			*p++
			return *p
		}()) >= 5 {
			break
		}
		v0 = 0
	}
	if int32(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(0x8531A0, 2576) + 2251)))) == 0 {
		for j = 0; j < 120; j += 24 {
			*memmap.PtrUint32(6112660, uint32(j)+0xFFCEC) = *memmap.PtrUint32(0x587000, uint32(j)+0x209A0)
			*memmap.PtrUint32(6112660, uint32(j)+0xFFCF0) = *memmap.PtrUint32(0x587000, uint32(j)+133540)
			*memmap.PtrUint32(6112660, uint32(j)+0xFFCF4) = *memmap.PtrUint32(0x587000, uint32(j)+0x209A8)
			*memmap.PtrUint32(6112660, uint32(j)+1047800) = *memmap.PtrUint32(0x587000, uint32(j)+0x209AC)
			*memmap.PtrUint32(6112660, uint32(j)+0xFFCFC) = *memmap.PtrUint32(0x587000, uint32(j)+0x209B0)
			*memmap.PtrUint32(6112660, uint32(j)+0xFFD00) = 0
		}
	}
	nox_xxx_clientUpdateButtonRow_45E110(0)
	return 1
}
func nox_xxx_quickbarButtonBook_45F3F0() int32 {
	var v0 *libc.WChar
	if sub_45CFC0() != 0 {
		v0 = strMan.GetStringInFile("CloseSpellbookTT", "C:\\NoxPost\\src\\Client\\Gui\\guispell.c")
	} else {
		v0 = strMan.GetStringInFile("OpenSpellbookTT", "C:\\NoxPost\\src\\Client\\Gui\\guispell.c")
	}
	nox_xxx_cursorSetTooltip_4776B0(v0)
	return 1
}
func sub_45F480(a1 int32) int32 {
	var v1 *libc.WChar
	if sub_45F500(int32(*(*uint16)(unsafe.Pointer(uintptr(a1 + 368)))), int32(uintptr(memmap.PtrOff(6112660, uint32(int32(uint16(*(*uint32)(unsafe.Pointer(uintptr(a1 + 368)))>>16))*256)+0xFFE84)))) != 0 {
		v1 = strMan.GetStringInFile("ToolTipCastOnMe", "C:\\NoxPost\\src\\Client\\Gui\\guispell.c")
	} else {
		v1 = strMan.GetStringInFile("ToolTipCastAtOther", "C:\\NoxPost\\src\\Client\\Gui\\guispell.c")
	}
	nox_xxx_cursorSetTooltip_4776B0(v1)
	return 1
}
func sub_45F9B0(a1 *uint32) int32 {
	var (
		v1  *uint32
		v2  *libc.WChar
		v3  *byte
		v5  int32
		v6  int32
		v7  int32
		v8  int32
		v9  int32
		v10 int32
		v11 [32]libc.WChar
	)
	if *memmap.PtrUint32(6112660, 1049476) == 0 {
		v1 = a1
		nox_client_wndGetPosition_46AA60((*nox_window)(unsafe.Pointer(a1)), (*uint32)(unsafe.Pointer(&v6)), (*uint32)(unsafe.Pointer(&a1)))
		nox_window_get_size((*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v1)))))), &v7, &v9)
		v5 = int32(*(*uint8)(unsafe.Pointer(uintptr(uint32(uintptr(nox_xxx_aClosewoodengat_587000_133480)) + 200)))) + 1
		v2 = strMan.GetStringInFile("SpellSet", "C:\\NoxPost\\src\\Client\\Gui\\guispell.c")
		nox_swprintf(&v11[0], v2, v5)
		nox_xxx_drawGetStringSize_43F840(nil, &v11[0], &v8, &v10, 0)
		v3 = (*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(a1))), -nox_xxx_guiFontHeightMB_43F320(nil)))
		a1 = (*uint32)(unsafe.Add(unsafe.Pointer(v3), nox_xxx_guiFontHeightMB_43F320(unsafe.Pointer(uintptr(*memmap.PtrInt32(6112660, 1049684))))+1))
		v6 += (v7 - v8) / 2
		nox_xxx_drawSetTextColor_434390(int32(nox_color_white_2523948))
		nox_xxx_drawSetColor_4343E0(*memmap.PtrInt32(0x852978, 4))
		nox_draw_drawStringHL_43F730(nil, (*libc.WChar)(unsafe.Pointer((*int16)(unsafe.Pointer(&v11[0])))), v6, int32(uintptr(unsafe.Pointer(a1))))
	}
	return 1
}
func nox_xxx_quickbarDraw_45FAC0(a1 *uint32) int32 {
	var (
		v1 *uint32
		v2 *libc.WChar
		v4 int32
		v5 int32
		v6 int32
		v7 int32
		v8 int32
		v9 [32]libc.WChar
	)
	v1 = a1
	nox_client_wndGetPosition_46AA60((*nox_window)(unsafe.Pointer(a1)), (*uint32)(unsafe.Pointer(&a1)), (*uint32)(unsafe.Pointer(&v7)))
	nox_window_get_size((*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v1)))))), &v5, &v8)
	v4 = int32(*memmap.PtrUint8(6112660, 1048140)) + 1
	v2 = strMan.GetStringInFile("TrapSet", "C:\\NoxPost\\src\\Client\\Gui\\guispell.c")
	nox_swprintf(&v9[0], v2, v4)
	nox_xxx_drawGetStringSize_43F840(nil, &v9[0], &v6, nil, 0)
	a1 = (*uint32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(a1))), (v5-v6)/2))))
	nox_xxx_drawSetTextColor_434390(int32(nox_color_white_2523948))
	nox_xxx_drawSetColor_4343E0(*memmap.PtrInt32(0x852978, 4))
	nox_draw_drawStringHL_43F730(nil, (*libc.WChar)(unsafe.Pointer((*int16)(unsafe.Pointer(&v9[0])))), int32(uintptr(unsafe.Pointer(a1))), v7)
	return 1
}
func nox_xxx_quickBarDrawFn_45FBD0(yTop int32) int32 {
	var (
		v1    int32
		v2    int32
		v3    *uint32
		v5    int32
		v6    int32
		v7    *int32
		v8    int32
		v9    int32
		v10   *libc.WChar
		v11   *int16
		v12   *libc.WChar
		v13   int32
		v14   uint32
		v15   int32
		v16   int32
		v17   int32
		xLeft int32
		v19   int32
	)
	v1 = yTop
	v2 = 0
	v3 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(yTop + 368))) + 212)))
	for {
		if uint32(yTop) == *v3 {
			break
		}
		v2++
		v3 = (*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*1))
		if v2 >= 5 {
			break
		}
	}
	if v2 == 5 {
		return 0
	}
	nox_client_wndGetPosition_46AA60((*nox_window)(unsafe.Pointer(uintptr(yTop))), (*uint32)(unsafe.Pointer(&xLeft)), (*uint32)(unsafe.Pointer(&yTop)))
	v5 = nox_xxx_guiFontHeightMB_43F320(nil)
	v6 = int32(*(*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(nox_xxx_aClosewoodengat_587000_133480)) + 204))))
	v7 = (*int32)(unsafe.Pointer(uintptr(v6 + v2*8)))
	v8 = int32(*(*uint32)(unsafe.Pointer(uintptr(v6 + v2*8))))
	if v8 != 0 {
		if nox_frame_xxx_2598000 <= 10 || uint32(v2) != *memmap.PtrUint32(0x587000, 133484) || (nox_frame_xxx_2598000-*memmap.PtrUint32(6112660, 1049540)) >= 10 {
			v9 = int32(uintptr(nox_xxx_spellGetAbilityIcon_425310(v8, 0)))
			v19 = 0
		} else {
			v9 = int32(uintptr(nox_xxx_spellGetAbilityIcon_425310(v8, 1)))
			v19 = 1
		}
		v16 = *v7
		if *memmap.PtrUint32(6112660, uint32(*v7*24)+0xFFCE0) != 0 {
			v10 = (*libc.WChar)(unsafe.Pointer(uintptr(nox_xxx_abilityGetName_0_425260(v16))))
			nox_xxx_wndWddSetTooltip_46B000((*nox_window_data)(unsafe.Pointer(uintptr(v1+36))), v10)
			if v9 != 0 {
			LABEL_14:
				nox_client_drawImageAt_47D2C0((*nox_video_bag_image_t)(unsafe.Pointer(uintptr(v9))), xLeft, yTop)
			LABEL_19:
				v13 = *v7 * 24
				if *memmap.PtrUint32(6112660, uint32(v13)+0xFFCE0) != 0 || *memmap.PtrUint32(6112660, uint32(v13)+0xFFCDC) == 0 || nox_xxx_playerAnimCheck_4372B0() != 0 {
					v14 = uint32(nox_xxx_abilityCooldown_4252D0(*v7))
					if v14/nox_gameFPS != 0 {
						if v19 == 0 {
							nox_client_drawRectFilledAlpha_49CF10(xLeft, yTop, 34, int32(34-(nox_frame_xxx_2598000-*memmap.PtrUint32(6112660, uint32(*v7*24)+0xFFCE8))/(v14/nox_gameFPS)))
						}
					}
				}
				return 1
			}
			nox_xxx_drawSetTextColor_434390(int32(nox_color_white_2523948))
			v17 = yTop + v5 + 2
			v15 = xLeft + 6
			v11 = (*int16)(unsafe.Pointer(strMan.GetStringInFile("NoIcon", "C:\\NoxPost\\src\\Client\\Gui\\guispell.c")))
		} else {
			v12 = (*libc.WChar)(unsafe.Pointer(uintptr(nox_xxx_abilityGetName_0_425260(v16))))
			nox_xxx_wndWddSetTooltip_46B000((*nox_window_data)(unsafe.Pointer(uintptr(v1+36))), v12)
			if v9 != 0 {
				goto LABEL_14
			}
			nox_xxx_drawSetTextColor_434390(int32(nox_color_white_2523948))
			v17 = yTop + v5 + 2
			v15 = xLeft + 6
			v11 = (*int16)(unsafe.Pointer(strMan.GetStringInFile("NoIcon", "C:\\NoxPost\\src\\Client\\Gui\\guispell.c")))
		}
		noxrend.DrawString(nil, (*libc.WChar)(unsafe.Pointer(v11)), image.Pt(v15, v17))
		goto LABEL_19
	}
	return 1
}
func nox_xxx_quickBarWarriorDraw_45FDE0(yTop int32) int32 {
	var (
		v1    int32
		v2    int32
		v3    *uint8
		v4    *int32
		v5    *uint32
		v7    int32
		v8    int32
		v9    int32
		v10   uint8
		v11   *libc.WChar
		v12   *int16
		v13   int32
		v14   int32
		v15   *int32
		v16   int32
		v17   int32
		xLeft int32
		v19   *uint8
		v20   int32
	)
	v1 = yTop
	v2 = 0
	v3 = *(**uint8)(unsafe.Pointer(uintptr(yTop + 368)))
	v19 = v3
	v4 = (*int32)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v3))), unsafe.Sizeof(uint32(0))*51))))))
	v5 = (*uint32)(unsafe.Add(unsafe.Pointer(v3), 212))
	for {
		if uint32(yTop) == *v5 {
			break
		}
		v2++
		v5 = (*uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(uint32(0))*1))
		if v2 >= 5 {
			break
		}
	}
	if v2 == 5 {
		return 0
	}
	nox_client_wndGetPosition_46AA60((*nox_window)(unsafe.Pointer(uintptr(yTop))), (*uint32)(unsafe.Pointer(&xLeft)), (*uint32)(unsafe.Pointer(&yTop)))
	if *(*int32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(int32(0))*uintptr(v2*2))) != 0 {
		if nox_frame_xxx_2598000 <= 10 || uint32(v2) != *memmap.PtrUint32(0x587000, 133484) || (nox_frame_xxx_2598000-*memmap.PtrUint32(6112660, 1049540)) >= 10 {
			v7 = int32(uintptr(nox_xxx_spellIcon_424A90(*(*int32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(int32(0))*uintptr(v2*2))))))
		} else {
			v7 = int32(uintptr(nox_xxx_spellIconHighlight_424AB0(*(*int32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(int32(0))*uintptr(v2*2))))))
		}
		v8 = *(*int32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(int32(0))*uintptr(v2*2)))
		v9 = v7
		v10 = *memmap.PtrUint8(6112660, uint32(v8)+0x1003C8)
		if int32(int8(v10)) > 0 {
			*memmap.PtrUint8(6112660, uint32(v8)+0x1003C8) = uint8(int8(int32(v10) - 1))
		}
		v11 = nox_xxx_spellTitle_424930(*(*int32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(int32(0))*uintptr(v2*2))))
		nox_xxx_wndWddSetTooltip_46B000((*nox_window_data)(unsafe.Pointer(uintptr(v1+36))), v11)
		if v9 != 0 {
			nox_client_drawImageAt_47D2C0((*nox_video_bag_image_t)(unsafe.Pointer(uintptr(v9))), xLeft, yTop)
		} else {
			nox_xxx_drawSetTextColor_434390(int32(nox_color_white_2523948))
			v17 = nox_xxx_guiFontHeightMB_43F320(nil) + yTop + 2
			v16 = xLeft + 6
			v12 = (*int16)(unsafe.Pointer(strMan.GetStringInFile("NoIcon", "C:\\NoxPost\\src\\Client\\Gui\\guispell.c")))
			noxrend.DrawString(nil, (*libc.WChar)(unsafe.Pointer(v12)), image.Pt(v16, v17))
		}
		v13 = nox_xxx_cliGetMana_470DD0()
		v14 = nox_xxx_spellManaCost_4249A0(*(*int32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(int32(0))*uintptr(v2*2))), 1)
		v20 = v14
		if unsafe.Pointer(v19) == memmap.PtrOff(6112660, 1047940) && v2 > 0 {
			v15 = v4
			v19 = (*uint8)(unsafe.Pointer(uintptr(v2)))
			for {
				if *v15 != 0 {
					v13 -= nox_xxx_spellManaCost_4249A0(*v15, 2)
				}
				v15 = (*int32)(unsafe.Add(unsafe.Pointer(v15), unsafe.Sizeof(int32(0))*2))
				v19 = (*uint8)(unsafe.Add(unsafe.Pointer(v19), -1))
				if v19 == nil {
					break
				}
			}
			v14 = v20
		}
		if !nox_xxx_spellIsEnabled_424B70(*(*int32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(int32(0))*uintptr(v2*2)))) || nox_xxx_playerAnimCheck_4372B0() != 0 || nox_xxx_spriteCheckFlag31_4356C0((*nox_drawable)(unsafe.Pointer(uintptr(*memmap.PtrInt32(0x852978, 8)))), 29) {
			nox_client_drawRectFilledAlpha_49CF10(xLeft, yTop, 30, 30)
			return 1
		}
		if v13 < v14 && v14 != 0 {
			nox_client_drawRectFilledAlpha_49CF10(xLeft, yTop, 30, (v14-v13)*30/v14)
			return 1
		}
	} else {
		nox_xxx_wndWddSetTooltip_46B000((*nox_window_data)(unsafe.Pointer(uintptr(v1+36))), (*libc.WChar)(memmap.PtrOff(6112660, 1049716)))
	}
	return 1
}
func sub_460070() int32 {
	var (
		result int32
		v1     *byte
		v2     *byte
		v3     *libc.WChar
		v4     *libc.WChar
		v5     *byte
		v6     *byte
		v7     *libc.WChar
		v8     *libc.WChar
	)
	result = int32(*memmap.PtrUint32(0x8531A0, 2576))
	if *memmap.PtrUint32(0x8531A0, 2576) != 0 {
		if int32(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(0x8531A0, 2576) + 2251)))) == 1 {
			v5 = (*byte)(unsafe.Pointer(nox_xxx_gLoadImg("QuickBarTrap")))
			nox_xxx_wndSetIcon_46AE60(*(*int32)(unsafe.Pointer(&dword_5d4594_1049504)), int32(uintptr(unsafe.Pointer(v5))))
			v6 = (*byte)(unsafe.Pointer(nox_xxx_gLoadImg("QuickBarTrapHit")))
			nox_xxx_wndSetIconLit_46AEA0(*(*int32)(unsafe.Pointer(&dword_5d4594_1049504)), int32(uintptr(unsafe.Pointer(v6))))
			v7 = strMan.GetStringInFile("ToolTipLayTrap", "C:\\NoxPost\\src\\Client\\Gui\\guispell.c")
			nox_xxx_wndWddSetTooltip_46B000((*nox_window_data)(unsafe.Pointer(uintptr(dword_5d4594_1049520+36))), v7)
			nox_xxx_wnd_46AD60(*(*int32)(unsafe.Pointer(&dword_5d4594_1049520)), 8)
			v8 = strMan.GetStringInFile("ToolTipTrapConstruct", "C:\\NoxPost\\src\\Client\\Gui\\guispell.c")
			nox_xxx_wndWddSetTooltip_46B000((*nox_window_data)(unsafe.Pointer(uintptr(dword_5d4594_1049500+36))), v8)
			nox_xxx_wnd_46AD60(*(*int32)(unsafe.Pointer(&dword_5d4594_1049500)), 8)
			result = (*(*int32)(unsafe.Pointer(&dword_5d4594_1049500))).setDraw(nil)
		} else {
			result = int32(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(0x8531A0, 2576) + 2251)))) - 2
			if int32(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(0x8531A0, 2576) + 2251)))) == 2 {
				v1 = (*byte)(unsafe.Pointer(nox_xxx_gLoadImg("QuickBarBomber")))
				nox_xxx_wndSetIcon_46AE60(*(*int32)(unsafe.Pointer(&dword_5d4594_1049504)), int32(uintptr(unsafe.Pointer(v1))))
				v2 = (*byte)(unsafe.Pointer(nox_xxx_gLoadImg("QuickBarBomberHit")))
				nox_xxx_wndSetIconLit_46AEA0(*(*int32)(unsafe.Pointer(&dword_5d4594_1049504)), int32(uintptr(unsafe.Pointer(v2))))
				v3 = strMan.GetStringInFile("ToolTipSummonBomber", "C:\\NoxPost\\src\\Client\\Gui\\guispell.c")
				nox_xxx_wndWddSetTooltip_46B000((*nox_window_data)(unsafe.Pointer(uintptr(dword_5d4594_1049520+36))), v3)
				nox_xxx_wnd_46AD60(*(*int32)(unsafe.Pointer(&dword_5d4594_1049520)), 8)
				v4 = strMan.GetStringInFile("ToolTipTrapConstruct", "C:\\NoxPost\\src\\Client\\Gui\\guispell.c")
				nox_xxx_wndWddSetTooltip_46B000((*nox_window_data)(unsafe.Pointer(uintptr(dword_5d4594_1049500+36))), v4)
				nox_xxx_wnd_46AD60(*(*int32)(unsafe.Pointer(&dword_5d4594_1049500)), 8)
				result = (*(*int32)(unsafe.Pointer(&dword_5d4594_1049500))).setDraw(nil)
			}
		}
	}
	return result
}
func nox_xxx_quickbarAddTrap_460EC0(a1 int32) *uint32 {
	var (
		result *uint32
		v2     int32
		v3     int32
		v4     int32
		v5     int32
		v6     int32
		v7     int8
		v8     int8
	)
	if a1 == 0 {
		return (*uint32)(unsafe.Pointer(uintptr(sub_460070())))
	}
	result = *(**uint32)(unsafe.Pointer(&dword_5d4594_1049500))
	if (int32(*(*uint8)(unsafe.Pointer(uintptr(dword_5d4594_1049500 + 4)))) & 8) == 0 {
		v2 = 50
		dword_5d4594_1049536 = uint32(nox_win_height + 1)
		for {
			v8 = int8(randomIntMinMax(4, 6))
			v7 = int8(randomIntMinMax(3, 6))
			v6 = randomIntMinMax(-20, -5)
			v5 = randomIntMinMax(-5, 5)
			v4 = int32(uint32(randomIntMinMax(0, 20)) + *(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1049504 + 20))) + 10)
			v3 = randomIntMinMax(0, 20)
			result = (*uint32)(unsafe.Pointer(nox_client_newScreenParticle_431540(0, int32(uint32(v3)+*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1049504 + 16)))+10), v4, v5, v6, 1, v7, v8, 2, 1)))
			v2--
			if v2 == 0 {
				break
			}
		}
	}
	return result
}
