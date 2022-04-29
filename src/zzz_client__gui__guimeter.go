package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	noxflags "github.com/noxworld-dev/opennox/v1/common/flags"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"math"
	"unsafe"
)

func sub_4710B0() int32 {
	var (
		v0     *libc.WChar
		v1     *libc.WChar
		v2     *libc.WChar
		v3     int32
		v4     *libc.WChar
		v5     *libc.WChar
		result int32
		v7     *libc.WChar
	)
	v0 = (*libc.WChar)(unsafe.Pointer(uintptr(sub_4615C0())))
	v1 = v0
	if v0 != nil {
		v2 = nox_xxx_clientAskInfoMb_4BF050((*nox_drawable)(unsafe.Pointer(v0)))
		nox_wcscpy((*libc.WChar)(memmap.PtrOff(6112660, 1091968)), v2)
		if *(*libc.WChar)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(libc.WChar(0))*58))&12 != 0 {
			v3 = int32(*memmap.PtrUint32(6112660, 1096292))
			if *memmap.PtrUint32(6112660, 1096292) == 0 {
				v3 = nox_xxx_getNameId_4E3AA0(CString("Quiver"))
				*memmap.PtrUint32(6112660, 1096292) = uint32(v3)
			}
			v4 = (*libc.WChar)(unsafe.Pointer(uintptr(sub_461600(v3))))
			if v4 != nil {
				nox_wcscat((*libc.WChar)(memmap.PtrOff(6112660, 1091968)), libc.CWString("\n"))
				v5 = nox_xxx_clientAskInfoMb_4BF050((*nox_drawable)(unsafe.Pointer(v4)))
				nox_wcscat((*libc.WChar)(memmap.PtrOff(6112660, 1091968)), v5)
			}
		}
		nox_xxx_cursorSetTooltip_4776B0((*libc.WChar)(memmap.PtrOff(6112660, 1091968)))
		result = 1
	} else {
		v7 = strMan.GetStringInFile("ToolTipCurWeapon", "C:\\NoxPost\\src\\Client\\Gui\\guimeter.c")
		nox_xxx_cursorSetTooltip_4776B0(v7)
		result = 1
	}
	return result
}
func sub_471160(a1 int32, a2 int32, a3 int32, a4 int32, a5 int32) int32 {
	var (
		v5     *libc.WChar
		v6     *libc.WChar
		result int32
	)
	nox_windows_arr_1093036[5].win = nox_window_new((*nox_window)(unsafe.Pointer(uintptr(a1))), 1032, a2, a3, a4, a5, nil)
	nox_windows_arr_1093036[6].win = nox_window_new((*nox_window)(unsafe.Pointer(uintptr(a1))), 1032, a2-17, a3-15, 15, 15, nil)
	nox_windows_arr_1093036[5].win.SetAllFuncs(nil, func(arg1 *nox_window, arg2 unsafe.Pointer) int32 {
		return sub_471250((*uint32)(unsafe.Pointer(arg1)))
	}, nil)
	v5 = strMan.GetStringInFile("ToolTipCharges", "C:\\NoxPost\\src\\Client\\Gui\\guimeter.c")
	nox_xxx_wndWddSetTooltip_46B000((*nox_window_data)(unsafe.Pointer((*libc.WChar)(unsafe.Pointer(&nox_windows_arr_1093036[5].win.draw_data)))), v5)
	nox_windows_arr_1093036[6].win.SetAllFuncs(nil, func(arg1 *nox_window, arg2 unsafe.Pointer) int32 {
		return sub_471450((*uint32)(unsafe.Pointer(arg1)))
	}, nil)
	v6 = strMan.GetStringInFile("ToolTipCharges", "C:\\NoxPost\\src\\Client\\Gui\\guimeter.c")
	nox_xxx_wndWddSetTooltip_46B000((*nox_window_data)(unsafe.Pointer((*libc.WChar)(unsafe.Pointer(&nox_windows_arr_1093036[6].win.draw_data)))), v6)
	result = int32(uintptr(unsafe.Pointer(nox_windows_arr_1093036[5].win)))
	nox_windows_arr_1093036[5].win.widget_data = unsafe.Pointer(uintptr(5))
	nox_windows_arr_1093036[6].win.widget_data = unsafe.Pointer(uintptr(6))
	return result
}
func nox_xxx_guiHealthManaInit_4714E0() int32 {
	var (
		v1  int32
		v2  *uint8
		v3  *libc.WChar
		v5  *libc.WChar
		v6  *libc.WChar
		v8  *uint32
		v9  *byte
		v10 *libc.WChar
		v11 *libc.WChar
		v12 *uint32
		v13 *byte
		v14 *libc.WChar
		v15 [32]byte
	)
	if *memmap.PtrUint32(0x8531A0, 2576) == 0 {
		return 0
	}
	dword_5d4594_1096288 = uint32(uintptr(guiFontPtrByName(CString("small"))))
	dword_5d4594_1096264 = 0
	dword_5d4594_1096256 = 0
	dword_5d4594_1096260 = 0
	*memmap.PtrUint32(6112660, 1091908) = 0
	*memmap.PtrUint32(6112660, 1091912) = 0
	*memmap.PtrUint32(6112660, 1091916) = uint32(nox_win_width)
	*memmap.PtrUint32(6112660, 1091920) = uint32(nox_win_height)
	*memmap.PtrUint32(6112660, 1091940) = uint32(nox_win_width)
	*memmap.PtrUint32(6112660, 1091944) = uint32(nox_win_height)
	*memmap.PtrUint32(6112660, 1091924) = 0
	*memmap.PtrUint32(6112660, 1091928) = 0
	v1 = 0
	v2 = (*uint8)(memmap.PtrOff(6112660, 1092996))
	var v2i int32 = 0
	for {
		nox_sprintf(&v15[0], CString("HealthMana%d"), func() int32 {
			p := &v1
			*p++
			return *p
		}())
		*(*uint32)(unsafe.Pointer(v2)) = uint32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg(&v15[0]))))
		v2 = (*uint8)(unsafe.Add(unsafe.Pointer(v2), 4))
		v2i++
		if v2i >= 10 {
			break
		}
	}
	if *memmap.PtrUint32(6112660, 1096268) == 0 {
		*memmap.PtrUint32(6112660, 1096268) = uint32(nox_xxx_getTTByNameSpriteMB_44CFC0("RedPotion"))
		dword_5d4594_1096272 = uint32(nox_xxx_getTTByNameSpriteMB_44CFC0("BluePotion"))
		dword_5d4594_1096276 = uint32(nox_xxx_getTTByNameSpriteMB_44CFC0("CurePoisonPotion"))
		dword_5d4594_1096280 = uint32(nox_xxx_getTTByNameSpriteMB_44CFC0("RedApple"))
		dword_5d4594_1096284 = uint32(nox_xxx_getTTByNameSpriteMB_44CFC0("Meat"))
	}
	dword_5d4594_1090276 = uint32(uintptr(unsafe.Pointer(nox_window_new(nil, 136, nox_win_width-91, nox_win_height-201, 91, 201, nil))))
	nox_xxx_wndSetIcon_46AE60(*(*int32)(unsafe.Pointer(&dword_5d4594_1090276)), *memmap.PtrInt32(6112660, 1092996))
	dword_5d4594_1091364 = uint32(uintptr(unsafe.Pointer(nox_window_new((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1090276))))), 8, 6, 166, 28, 30, nil))))
	(*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1091364)))).SetAllFuncs(func(arg1 int32, arg2 int32, arg3 int32, arg4 int32) int32 {
		return nox_xxx_guiBottleSlotProc_471B90(arg1, arg2)
	}, func(arg1 *nox_window, arg2 unsafe.Pointer) int32 {
		return nox_xxx_guiBottleSlotDrawFn_471A80((*uint32)(unsafe.Pointer(arg1)))
	}, nil)
	v3 = strMan.GetStringInFile("CurePoisonSlotTT", "C:\\NoxPost\\src\\Client\\Gui\\guimeter.c")
	nox_xxx_wndWddSetTooltip_46B000((*nox_window_data)(unsafe.Pointer(uintptr(dword_5d4594_1091364+36))), v3)
	*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1091364 + 32))) = 2
	*memmap.PtrUint16(6112660, 1091384) = 0
	var t4 *nox_thing = nox_get_thing(*(*int32)(unsafe.Pointer(&dword_5d4594_1096276)))
	if t4 != nil {
		nox_drawable_link_thing((*nox_drawable)(memmap.PtrOff(6112660, 1091388)), t4.field_1c)
		*memmap.PtrUint32(6112660, 1091368) = uint32(uintptr(memmap.PtrOff(6112660, 1091388)))
		if true {
			*memmap.PtrUint32(6112660, 1091508) |= 0x40000000
		}
	} else {
		*memmap.PtrUint32(6112660, 1091368) = 0
	}
	*memmap.PtrUint32(6112660, 1091380) = dword_5d4594_1096276
	dword_5d4594_1090292 = uint32(uintptr(unsafe.Pointer(nox_window_new((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1090276))))), 8, 34, 166, 28, 30, nil))))
	(*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1090292)))).SetAllFuncs(func(arg1 int32, arg2 int32, arg3 int32, arg4 int32) int32 {
		return nox_xxx_guiBottleSlotProc_471B90(arg1, arg2)
	}, func(arg1 *nox_window, arg2 unsafe.Pointer) int32 {
		return nox_xxx_guiBottleSlotDrawFn_471A80((*uint32)(unsafe.Pointer(arg1)))
	}, nil)
	v5 = strMan.GetStringInFile("HealthSlotTT", "C:\\NoxPost\\src\\Client\\Gui\\guimeter.c")
	nox_xxx_wndWddSetTooltip_46B000((*nox_window_data)(unsafe.Pointer(uintptr(dword_5d4594_1090292+36))), v5)
	*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1090292 + 32))) = 0
	*memmap.PtrUint16(6112660, 1090312) = 0
	*memmap.PtrUint32(6112660, 1090296) = 0
	*memmap.PtrUint32(6112660, 1090308) = 0
	if int32(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(0x8531A0, 2576) + 2251)))) != 0 {
		dword_5d4594_1090828 = uint32(uintptr(unsafe.Pointer(nox_window_new((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1090276))))), 8, 62, 166, 28, 30, nil))))
		(*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1090828)))).SetAllFuncs(func(arg1 int32, arg2 int32, arg3 int32, arg4 int32) int32 {
			return nox_xxx_guiBottleSlotProc_471B90(arg1, arg2)
		}, func(arg1 *nox_window, arg2 unsafe.Pointer) int32 {
			return nox_xxx_guiBottleSlotDrawFn_471A80((*uint32)(unsafe.Pointer(arg1)))
		}, nil)
		v6 = strMan.GetStringInFile("ManaSlotTT", "C:\\NoxPost\\src\\Client\\Gui\\guimeter.c")
		nox_xxx_wndWddSetTooltip_46B000((*nox_window_data)(unsafe.Pointer(uintptr(dword_5d4594_1090828+36))), v6)
		*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1090828 + 32))) = 1
		*memmap.PtrUint16(6112660, 1090848) = 0
		var t7 *nox_thing = nox_get_thing(*(*int32)(unsafe.Pointer(&dword_5d4594_1096272)))
		if t7 != nil {
			nox_drawable_link_thing((*nox_drawable)(memmap.PtrOff(6112660, 1090852)), t7.field_1c)
			*memmap.PtrUint32(6112660, 1090832) = uint32(uintptr(memmap.PtrOff(6112660, 1090852)))
			if true {
				*memmap.PtrUint32(6112660, 1090972) |= 0x40000000
			}
		} else {
			*memmap.PtrUint32(6112660, 1090832) = 0
		}
		*memmap.PtrUint32(6112660, 1090844) = dword_5d4594_1096272
		*memmap.PtrUint32(6112660, 1091900) = uint32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg("PoisonTube"))))
		v8 = (*uint32)(unsafe.Pointer(nox_window_new((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1090276))))), 136, 0, 0, 91, 159, nil)))
		v9 = (*byte)(unsafe.Pointer(nox_xxx_gLoadImg("HealthManaTubes")))
		nox_xxx_wndSetIcon_46AE60(int32(uintptr(unsafe.Pointer(v8))), int32(uintptr(unsafe.Pointer(v9))))
		nox_windows_arr_1093036[1].win = nox_window_new((*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v8)))))), 8, 60, 34, 25, 125, nil)
		nox_windows_arr_1093036[1].win.SetAllFuncs(func(arg1 int32, arg2 int32, arg3 int32, arg4 int32) int32 {
			return nox_xxx_guiHealthManaTubeProc_472100(arg1, arg2)
		}, func(arg1 *nox_window, arg2 unsafe.Pointer) int32 {
			return nox_xxx_guiHealthManaTubeDraw_471D10(int32(uintptr(unsafe.Pointer(arg1))))
		}, nil)
		v10 = strMan.GetStringInFile("ToolTipMana", "C:\\NoxPost\\src\\Client\\Gui\\guimeter.c")
		nox_xxx_wndWddSetTooltip_46B000((*nox_window_data)(unsafe.Pointer((*libc.WChar)(unsafe.Pointer(&nox_windows_arr_1093036[1].win.draw_data)))), v10)
		nox_windows_arr_1093036[1].win.widget_data = unsafe.Pointer(uintptr(1))
		nox_windows_arr_1093036[0].win = nox_window_new((*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v8)))))), 8, 34, 34, 25, 125, nil)
		nox_windows_arr_1093036[0].win.SetAllFuncs(func(arg1 int32, arg2 int32, arg3 int32, arg4 int32) int32 {
			return nox_xxx_guiHealthManaTubeProc_472100(arg1, arg2)
		}, func(arg1 *nox_window, arg2 unsafe.Pointer) int32 {
			return nox_xxx_guiHealthManaTubeDraw_471D10(int32(uintptr(unsafe.Pointer(arg1))))
		}, nil)
		v11 = strMan.GetStringInFile("ToolTipHealth", "C:\\NoxPost\\src\\Client\\Gui\\guimeter.c")
		nox_xxx_wndWddSetTooltip_46B000((*nox_window_data)(unsafe.Pointer((*libc.WChar)(unsafe.Pointer(&nox_windows_arr_1093036[0].win.draw_data)))), v11)
		nox_windows_arr_1093036[0].win.widget_data = nil
		nox_windows_arr_1093036[2].win = nox_window_new(nil, 8, 0, 0, 0, 0, nil)
		nox_windows_arr_1093036[2].win.SetAllFuncs(nil, func(arg1 *nox_window, arg2 unsafe.Pointer) int32 {
			return nox_xxx_drawHealthManaBar_471C00(int32(uintptr(unsafe.Pointer(arg1))))
		}, nil)
		nox_windows_arr_1093036[2].win.widget_data = nil
		nox_windows_arr_1093036[3].win = nox_window_new(nil, 8, 0, 0, 0, 0, nil)
		nox_windows_arr_1093036[3].win.SetAllFuncs(nil, func(arg1 *nox_window, arg2 unsafe.Pointer) int32 {
			return nox_xxx_drawHealthManaBar_471C00(int32(uintptr(unsafe.Pointer(arg1))))
		}, nil)
		nox_windows_arr_1093036[3].win.widget_data = unsafe.Pointer(uintptr(1))
		*memmap.PtrUint32(6112660, 1093176) = 1
	} else {
		*memmap.PtrUint32(6112660, 1091900) = uint32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg("WarriorPoisonTube"))))
		v12 = (*uint32)(unsafe.Pointer(nox_window_new((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1090276))))), 136, 0, 0, 91, 159, nil)))
		v13 = (*byte)(unsafe.Pointer(nox_xxx_gLoadImg("WarriorHealthTube")))
		nox_xxx_wndSetIcon_46AE60(int32(uintptr(unsafe.Pointer(v12))), int32(uintptr(unsafe.Pointer(v13))))
		nox_windows_arr_1093036[0].win = nox_window_new((*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v12)))))), 8, 34, 34, 25, 125, nil)
		nox_windows_arr_1093036[0].win.SetAllFuncs(func(arg1 int32, arg2 int32, arg3 int32, arg4 int32) int32 {
			return nox_xxx_guiHealthManaTubeProc_472100(arg1, arg2)
		}, func(arg1 *nox_window, arg2 unsafe.Pointer) int32 {
			return nox_xxx_guiHealthManaTubeDraw_471D10(int32(uintptr(unsafe.Pointer(arg1))))
		}, nil)
		v14 = strMan.GetStringInFile("ToolTipHealth", "C:\\NoxPost\\src\\Client\\Gui\\guimeter.c")
		nox_xxx_wndWddSetTooltip_46B000((*nox_window_data)(unsafe.Pointer((*libc.WChar)(unsafe.Pointer(&nox_windows_arr_1093036[0].win.draw_data)))), v14)
		nox_windows_arr_1093036[0].win.widget_data = nil
		nox_windows_arr_1093036[2].win = nox_window_new(nil, 24, 0, 0, 0, 0, nil)
		nox_windows_arr_1093036[2].win.SetAllFuncs(nil, func(arg1 *nox_window, arg2 unsafe.Pointer) int32 {
			return nox_xxx_drawHealthManaBar_471C00(int32(uintptr(unsafe.Pointer(arg1))))
		}, nil)
		nox_windows_arr_1093036[2].win.widget_data = nil
	}
	sub_472280()
	nox_xxx_guiHealthManaColorInit_470B00()
	if noxflags.HasGame(noxflags.GameModeQuest) {
		dword_5d4594_1096252 = 1
	} else {
		dword_5d4594_1096252 = 0
		nox_windows_arr_1093036[2].win.Hide()
		nox_windows_arr_1093036[3].win.Hide()
	}
	return 1
}
func nox_xxx_guiHealthManaTubeDraw_471D10(xLeft int32) int32 {
	var (
		v1     *uint32
		v2     int32
		v3     *uint8
		v4     int32
		result int32
		v6     int32
		v7     *int32
		v8     int32
		v9     *uint8
		v10    int32
		v11    int32
		v12    int32
		v13    int32
		v14    int32
		v15    int32
		v16    int32
		yTop   int32
		v18    int32
		v19    int32
		v20    *uint8
		v21    *uint8
	)
	v1 = (*uint32)(unsafe.Pointer(uintptr(xLeft)))
	v18 = int32(*(*uint32)(unsafe.Pointer(uintptr(xLeft + 32))))
	v2 = v18
	v3 = (*uint8)(unsafe.Pointer(&nox_windows_arr_1093036[v18]))
	v20 = v3
	if v18 == 0 && dword_5d4594_1096264 != 0 {
		nox_client_wndGetPosition_46AA60((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1090276)))), (*uint32)(unsafe.Pointer(&xLeft)), (*uint32)(unsafe.Pointer(&yTop)))
		nox_client_drawImageAt_47D2C0((*nox_video_bag_image_t)(unsafe.Pointer(uintptr(*memmap.PtrInt32(6112660, 1091900)))), xLeft, yTop)
	}
	nox_client_wndGetPosition_46AA60((*nox_window)(unsafe.Pointer(v1)), (*uint32)(unsafe.Pointer(&xLeft)), (*uint32)(unsafe.Pointer(&yTop)))
	xLeft += 5
	if dword_5d4594_3799468 != 0 {
		nox_client_drawSetColor_434460(int32(nox_color_black_2650656))
		nox_client_drawRectFilledOpaque_49CE30(xLeft, yTop, 15, 125)
	}
	v4 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v3))), unsafe.Sizeof(uint32(0))*2))))
	if v4 != 0 {
		v6 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v3))), unsafe.Sizeof(uint32(0))*1))) * 125 / uint32(v4))
		nox_client_drawRectFilledAlpha_49CF10(xLeft, yTop, 15, 125-v6)
		nox_client_drawSetColor_434460(int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v3))), unsafe.Sizeof(uint32(0))*3)))))
		nox_client_drawEnableAlpha_434560(1)
		nox_client_drawRectFilledOpaque_49CE30(xLeft, yTop-v6+125, 15, v6)
		nox_client_drawEnableAlpha_434560(0)
		nox_client_drawAddPoint_49F500(xLeft, yTop-v6+125)
		nox_xxx_rasterPointRel_49F570(14, 0)
		nox_client_drawLineFromPoints_49E4B0()
		if v2 < 2 && nox_client_renderBubbles_80844 == 1 {
			v19 = 64
			v7 = memmap.PtrInt32(6112660, uint32(v2*1536)+0x10AE44)
			v21 = (*uint8)(memmap.PtrOff(6112660, uint32(v2*1536)+0x10AE44))
			for {
				if *(*int32)(unsafe.Add(unsafe.Pointer(v7), unsafe.Sizeof(int32(0))*2)) != 0 {
					v8 = *((*int32)(unsafe.Add(unsafe.Pointer(v7), -int(unsafe.Sizeof(int32(0))*1)))) >> 4
					if v8 >= 125-v6 {
						if dword_5d4594_1096264 != 0 {
							nox_client_drawSetColor_434460(int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v20))), unsafe.Sizeof(uint32(0))*3)))))
						} else {
							nox_client_drawSetColor_434460(*(*int32)(unsafe.Add(unsafe.Pointer(v7), unsafe.Sizeof(int32(0))*3)))
						}
						if *v7 <= 2 {
							nox_client_drawRectFilledOpaque_49CE30(xLeft+*((*int32)(unsafe.Add(unsafe.Pointer(v7), -int(unsafe.Sizeof(int32(0))*2)))), yTop+v8, *v7, *v7)
						} else {
							sub_4720C0(xLeft+*((*int32)(unsafe.Add(unsafe.Pointer(v7), -int(unsafe.Sizeof(int32(0))*2)))), yTop+v8)
						}
						*((*int32)(unsafe.Add(unsafe.Pointer(v7), -int(unsafe.Sizeof(int32(0))*1)))) -= *(*int32)(unsafe.Add(unsafe.Pointer(v7), unsafe.Sizeof(int32(0))*1))
					} else {
						*(*int32)(unsafe.Add(unsafe.Pointer(v7), unsafe.Sizeof(int32(0))*2)) = 0
					}
					v2 = v18
				}
				v7 = (*int32)(unsafe.Add(unsafe.Pointer(v7), unsafe.Sizeof(int32(0))*6))
				v19--
				if v19 == 0 {
					break
				}
			}
			if v6 > 1 {
				v9 = v21
				v10 = 64
				for {
					if *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v9))), unsafe.Sizeof(uint32(0))*2))) == 0 {
						v11 = randomIntMinMax(1, 100)
						if v11 >= 80 {
							*(*uint32)(unsafe.Pointer(v9)) = uint32(bool2int(v11 >= 95) + 2)
						} else {
							*(*uint32)(unsafe.Pointer(v9)) = 1
						}
						v12 = randomIntMinMax(0, 14)
						v13 = int32(*(*uint32)(unsafe.Pointer(v9)))
						*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v9))), -int(unsafe.Sizeof(uint32(0))*2)))) = uint32(v12)
						if v13+v12 > 15 {
							*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v9))), -int(unsafe.Sizeof(uint32(0))*2)))) = uint32(15 - v13)
						}
						*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v9))), -int(unsafe.Sizeof(uint32(0))*1)))) = uint32((125 - v13) * 16)
						*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v9))), unsafe.Sizeof(uint32(0))*1))) = uint32(randomIntMinMax(4, 48))
						*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v9))), unsafe.Sizeof(uint32(0))*2))) = 1
						*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v19))), 0)) = uint8(int8(randomIntMinMax(0, 64)))
						if v18 != 0 {
							v14 = int32(nox_color_rgb_4344A0(v19, v19, math.MaxUint8))
						} else {
							v14 = int32(nox_color_rgb_4344A0(math.MaxUint8, v19, v19))
						}
						*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v9))), unsafe.Sizeof(uint32(0))*3))) = uint32(v14)
					}
					v9 = (*uint8)(unsafe.Add(unsafe.Pointer(v9), 24))
					v10--
					if v10 == 0 {
						break
					}
				}
				v2 = v18
			}
		}
		if v2 == 0 {
			if *(*int32)(unsafe.Pointer(&dword_5d4594_1096260)) <= 0 {
				v16 = int32(dword_5d4594_1096256 >> 3)
			} else {
				v15 = int32(dword_5d4594_1096260 + dword_5d4594_1096256)
				dword_5d4594_1096256 = uint32(v15)
				dword_5d4594_1096260--
				v16 = v15 >> 3
				if v16 >= 10 {
					v16 = 0
					dword_5d4594_1096256 = 0
				}
			}
			nox_xxx_wndSetIcon_46AE60(*(*int32)(unsafe.Pointer(&dword_5d4594_1090276)), int32(*memmap.PtrUint32(6112660, uint32(v16*4)+0x10AD84)))
		}
		sub_472080()
		result = 1
	} else {
		nox_client_drawRectFilledAlpha_49CF10(xLeft, yTop, 15, 125)
		result = 1
	}
	return result
}
