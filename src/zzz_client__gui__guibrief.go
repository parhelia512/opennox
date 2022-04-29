package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"image"
	"unsafe"
)

func sub_44E410() *libc.WChar {
	var (
		v0     int32
		v1     int32
		i      int32
		v3     int32
		v4     int32
		result *libc.WChar
		v6     int32
		v7     int32
		v8     [64]byte
	)
	v0 = 0
	v7 = 0
	v1 = 0
	for i = 1; ; i++ {
		for {
			if v0 != 0 {
				if v0 == 1 {
					if v1 != 0 {
						v3 = bool2int(v1 != 1) + 2
					} else {
						v3 = 4
					}
				} else {
					v3 = v0 + 3
				}
			} else {
				v3 = 1
			}
			nox_sprintf(&v8[0], CString("Briefing:%sChapterBegin%d"), *memmap.PtrUint32(0x587000, uint32(v1*4)+0x1E040), i)
			v4 = (v1 + v0 + v1*10) * 32
			*memmap.PtrUint32(6112660, uint32(v4)+831300) = uint32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg(&v8[9]))))
			*memmap.PtrUint32(6112660, uint32(v4)+0xCAF48) = uint32(uintptr(unsafe.Pointer(nox_strman_loadString_40F1D0(&v8[0], (**byte)(memmap.PtrOff(6112660, uint32(v4)+0xCAF4C)), CString("C:\\NoxPost\\src\\client\\Gui\\GUIBrief.c"), 1221))))
			v6 = int32(*memmap.PtrUint32(0x587000, uint32(v1*4)+0x1E040))
			*memmap.PtrUint32(6112660, uint32(v4)+0xCAF50) = uint32(v3)
			nox_sprintf(&v8[0], CString("Briefing:%sChapterLoss%d"), v6, i)
			*memmap.PtrUint32(6112660, uint32(v4)+0xCAF54) = uint32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg(&v8[9]))))
			v1++
			*memmap.PtrUint32(6112660, uint32(v4)+831320) = uint32(uintptr(unsafe.Pointer(nox_strman_loadString_40F1D0(&v8[0], (**byte)(memmap.PtrOff(6112660, uint32(v4)+0xCAF5C)), CString("C:\\NoxPost\\src\\client\\Gui\\GUIBrief.c"), 1227))))
			*memmap.PtrUint32(6112660, uint32(v4)+0xCAF60) = uint32(v3)
			if v1 >= 3 {
				break
			}
			v0 = v7
		}
		v7 = i
		if i >= 11 {
			break
		}
		v0 = i
		v1 = 0
	}
	*memmap.PtrUint32(6112660, 831264) = uint32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg("CreditsImage"))))
	result = nox_strman_loadString_40F1D0(CString("Nox:Credits"), (**byte)(memmap.PtrOff(6112660, 831272)), CString("C:\\NoxPost\\src\\client\\Gui\\GUIBrief.c"), 1233)
	*memmap.PtrUint32(6112660, 831268) = uint32(uintptr(unsafe.Pointer(result)))
	return result
}
func sub_44E8E0(a1 int32, a2 int32) int32 {
	var (
		v2          int32
		v3          int32
		v4          *libc.WChar
		v5          *libc.WChar
		v6          *uint8
		v7          *uint8
		v8          int32
		v9          int32
		v10         int32
		v11         int32
		v12         int32
		v13         int32
		v14         *libc.WChar
		v15         int32
		v16         *libc.WChar
		v17         int32
		v18         *libc.WChar
		v19         int32
		v20         *libc.WChar
		v21         *libc.WChar
		v22         int32
		v23         int32
		v24         int32
		v25         *libc.WChar
		v26         *libc.WChar
		v27         int32
		v28         *libc.WChar
		v29         *libc.WChar
		v30         int32
		v31         int32
		result      int32
		v33         int32
		v34         *uint16
		v35         int32
		v36         float32
		v37         int32
		v38         int32
		v39         int32
		v40         int32
		v41         int32
		v42         int32
		v43         int32
		v44         int32
		v45         int32
		v46         int32
		v47         int32
		v48         *uint8
		v49         int32
		v50         int32
		v51         int32
		v52         int32
		v53         int32
		WideCharStr [11]libc.WChar
		v55         [257]libc.WChar
		v56         [256]libc.WChar
		v57         [256]libc.WChar
	)
	v49 = (nox_win_width - NOX_DEFAULT_WIDTH) / 2
	v47 = 0
	v45 = 0
	v50 = (nox_win_height - NOX_DEFAULT_HEIGHT) / 2
	nox_xxx_drawSetTextColor_434390(int32(nox_color_white_2523948))
	v2 = nox_win_width / 2
	v3 = nox_win_height / 2
	v51 = nox_win_width / 2
	v52 = nox_win_height / 2
	v4 = strMan.GetStringInFile("GUIBrief.c:GauntletStatTitle", "C:\\NoxPost\\src\\client\\Gui\\GUIBrief.c")
	nox_wcscpy(&v55[1], v4)
	v5 = strMan.GetStringInFile("Noxworld.c:Stage", "C:\\NoxPost\\src\\client\\Gui\\GUIBrief.c")
	nox_wcscpy(&v56[0], v5)
	nox_wcscat(&v56[0], libc.CWString(" XX1 "))
	compat_itow(*memmap.PtrInt32(6112660, 831228), &WideCharStr[0], 10)
	nox_wcscat(&v56[0], &WideCharStr[0])
	nox_swprintf(&v57[0], libc.CWString("%s - %s"), &v55[1], &v56[0])
	nox_xxx_drawGetStringSize_43F840(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a2 + 200))))), &v57[0], &v39, &v38, 0)
	nox_xxx_drawSetTextColor_434390(int32(nox_color_white_2523948))
	noxrend.DrawString(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a2 + 200))))), (*libc.WChar)(unsafe.Pointer((*int16)(unsafe.Pointer(&v57[0])))), image.Pt(v2-v39/2, v38+v3-240))
	v40 = int32(*memmap.PtrUint32(0x587000, 122968) - *memmap.PtrUint32(0x587000, 122964))
	v36 = float32(float64(v38) * 1.5)
	v41 = int(v36)
	v6 = (*uint8)(memmap.PtrOff(0x587000, 122964))
	v43 = 0
	v42 = 0
	v7 = (*uint8)(memmap.PtrOff(6112660, 832364))
	v48 = (*uint8)(memmap.PtrOff(0x587000, 122964))
	for {
		v8 = int32(*(*uint32)(unsafe.Pointer(v6)) + uint32(v3) - 240)
		v9 = int32(*(*uint32)(unsafe.Pointer(v7)))
		v10 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v6))), -int(unsafe.Sizeof(uint32(0))*1)))) + uint32(v2) - 320)
		if *(*uint32)(unsafe.Pointer(v7)) != 0 {
			v43++
			if uint32(v9) == *memmap.PtrUint32(0x8531A0, 2576) {
				v47 = int32(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v7))), unsafe.Sizeof(uint16(0))*5))))
			} else {
				v45 += int32(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v7))), unsafe.Sizeof(uint16(0))*5))))
			}
			nox_xxx_drawSetTextColor_434390(int32(nox_color_orange_2614256))
			nox_swprintf(&v55[1], libc.CWString("%d) %s"), v42+1, *(*uint32)(unsafe.Pointer(v7))+4704)
			v11 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 200))))
			v46 = int32(*memmap.PtrUint32(0x587000, 122968) - *memmap.PtrUint32(0x587000, 122960) + uint32(v10) - 16)
			nox_xxx_drawGetStringSize_43F840(unsafe.Pointer(uintptr(v11)), &v55[1], &v44, &v53, 0)
			for v10+v44 >= v46 {
				v12 = int32(nox_wcslen(&v55[1]))
				if v12 <= 5 {
					break
				}
				v55[v12] = 0
				nox_xxx_drawGetStringSize_43F840(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a2 + 200))))), &v55[1], &v44, &v53, 0)
			}
			nox_xxx_drawStringWrap_43FAF0(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a2 + 200))))), &v55[1], v10, v8, v40-8, v38)
			v13 = v41 + v41/2 + v8
			nox_xxx_drawSetTextColor_434390(int32(nox_color_white_2523948))
			v14 = strMan.GetStringInFile("GUIBrief.c:GeneratorsDestroyed", "C:\\NoxPost\\src\\client\\Gui\\GUIBrief.c")
			nox_xxx_drawStringWrap_43FAF0(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a2 + 200))))), v14, v10, v13, *(*int32)(unsafe.Pointer(&dword_5d4594_832476)), v38)
			nox_swprintf(&v55[1], libc.CWString(" %d"), *((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v7))), unsafe.Sizeof(uint16(0))*3))))
			nox_xxx_drawSetTextColor_434390(int32(nox_color_green_2614268))
			nox_xxx_drawStringWrap_43FAF0(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a2 + 200))))), &v55[1], int32(uint32(v10)+dword_5d4594_832476), v13, int32(uint32(v40)-dword_5d4594_832476-8), v38)
			v15 = v41 + v13
			nox_xxx_drawSetTextColor_434390(int32(nox_color_white_2523948))
			v16 = strMan.GetStringInFile("GUIBrief.c:numSecretsFound", "C:\\NoxPost\\src\\client\\Gui\\GUIBrief.c")
			nox_xxx_drawStringWrap_43FAF0(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a2 + 200))))), v16, v10, v15, *(*int32)(unsafe.Pointer(&dword_5d4594_832476)), v38)
			nox_swprintf(&v55[1], libc.CWString(" %d"), *((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v7))), unsafe.Sizeof(uint16(0))*4))))
			nox_xxx_drawSetTextColor_434390(int32(nox_color_green_2614268))
			nox_xxx_drawStringWrap_43FAF0(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a2 + 200))))), &v55[1], int32(uint32(v10)+dword_5d4594_832476), v15, int32(uint32(v40)-dword_5d4594_832476-8), v38)
			v17 = v41 + v15
			nox_xxx_drawSetTextColor_434390(int32(nox_color_white_2523948))
			v18 = strMan.GetStringInFile("GUIBrief.c:Kills", "C:\\NoxPost\\src\\client\\Gui\\GUIBrief.c")
			nox_xxx_drawStringWrap_43FAF0(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a2 + 200))))), v18, v10, v17, *(*int32)(unsafe.Pointer(&dword_5d4594_832476)), v38)
			nox_swprintf(&v55[1], libc.CWString(" %d"), *((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v7))), unsafe.Sizeof(uint16(0))*2))))
			nox_xxx_drawSetTextColor_434390(int32(nox_color_green_2614268))
			nox_xxx_drawStringWrap_43FAF0(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a2 + 200))))), &v55[1], int32(uint32(v10)+dword_5d4594_832476), v17, int32(uint32(v40)-dword_5d4594_832476-8), v38)
			v19 = v41 + v17
			nox_xxx_drawSetTextColor_434390(int32(nox_color_white_2523948))
			v20 = strMan.GetStringInFile("GUIBrief.c:TotalScore", "C:\\NoxPost\\src\\client\\Gui\\GUIBrief.c")
			nox_xxx_drawStringWrap_43FAF0(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a2 + 200))))), v20, v10, v19, *(*int32)(unsafe.Pointer(&dword_5d4594_832476)), v38)
			nox_swprintf(&v55[1], libc.CWString(" %d"), *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v7))), unsafe.Sizeof(uint32(0))*3))))
			nox_xxx_drawSetTextColor_434390(int32(nox_color_blue_2650684))
			nox_xxx_drawStringWrap_43FAF0(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a2 + 200))))), &v55[1], int32(uint32(v10)+dword_5d4594_832476), v19, int32(uint32(v40)-dword_5d4594_832476-8), v38)
			v6 = v48
		}
		v6 = (*uint8)(unsafe.Add(unsafe.Pointer(v6), 8))
		v7 = (*uint8)(unsafe.Add(unsafe.Pointer(v7), 16))
		v42++
		v48 = v6
		if int32(uintptr(unsafe.Pointer(v6))) >= int32(uintptr(memmap.PtrOff(0x587000, 123012))) {
			break
		}
		v2 = v51
		v3 = v52
	}
	v21 = strMan.GetStringInFile("GeneralPrint:SecretsTotal", "C:\\NoxPost\\src\\client\\Gui\\GUIBrief.c")
	nox_swprintf(&v55[1], v21, *memmap.PtrUint32(6112660, 832356))
	nox_xxx_drawGetStringSize_43F840(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a2 + 200))))), &v55[1], &v39, &v38, 0)
	v22 = v49
	v23 = v49 - v39/2 + 320
	v24 = v50 + (150-v38)*2 + 150 - v38
	nox_xxx_drawSetTextColor_434390(int32(nox_color_orange_2614256))
	noxrend.DrawString(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a2 + 200))))), (*libc.WChar)(unsafe.Pointer((*int16)(unsafe.Pointer(&v55[1])))), image.Pt(v23, v24))
	if v47 != 0 {
		v37 = v47
		v25 = strMan.GetStringInFile("GeneralPrint:SecretsFound", "C:\\NoxPost\\src\\client\\Gui\\GUIBrief.c")
		nox_swprintf(&v55[1], v25, v37)
	} else {
		v26 = strMan.GetStringInFile("GeneralPrint:SecretsNoneFound", "C:\\NoxPost\\src\\client\\Gui\\GUIBrief.c")
		nox_swprintf(&v55[1], v26)
	}
	if v43 <= 1 {
		nox_wcscpy(&v57[0], &v55[1])
	} else {
		v27 = v45
		if v45 != 0 {
			v28 = strMan.GetStringInFile("GeneralPrint:SecretsFoundByFriends", "C:\\NoxPost\\src\\client\\Gui\\GUIBrief.c")
			if v28 != nil {
				nox_swprintf(&v56[0], v28, v27)
			}
		} else {
			v29 = strMan.GetStringInFile("GeneralPrint:SecretsNoneFoundByFriends", "C:\\NoxPost\\src\\client\\Gui\\GUIBrief.c")
			nox_wcscpy(&v56[0], v29)
		}
		nox_swprintf(&v57[0], libc.CWString("%s - %s"), &v55[1], &v56[0])
	}
	nox_xxx_drawGetStringSize_43F840(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a2 + 200))))), &v57[0], &v39, &v38, 0)
	v30 = v22 - v39/2 + 320
	v31 = v50 + (225-v38)*2
	nox_xxx_drawSetTextColor_434390(int32(nox_color_orange_2614256))
	noxrend.DrawString(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a2 + 200))))), (*libc.WChar)(unsafe.Pointer((*int16)(unsafe.Pointer(&v57[0])))), image.Pt(v30, v31))
	result = int32(nox_frame_xxx_2598000 / 30)
	if nox_frame_xxx_2598000%30 != 0 {
		if dword_587000_122956 != 1 {
			return result
		}
	} else {
		result = 1
		if dword_587000_122956 == 1 {
			dword_587000_122956 = nox_frame_xxx_2598000 % 30
			return result
		}
		dword_587000_122956 = 1
	}
	v33 = int32(nox_color_white_2523948)
	v34 = (*uint16)(unsafe.Pointer(strMan.GetStringInFile("GeneralPrint:QuestSplash12", "C:\\NoxPost\\src\\client\\Gui\\GUIBrief.c")))
	nox_xxx_drawGetStringSize_43F840(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a2 + 200))))), (*libc.WChar)(unsafe.Pointer(v34)), &v39, nil, 0)
	v35 = v22 - v39/2 + 320
	nox_xxx_drawSetTextColor_434390(v33)
	return noxrend.DrawString(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a2 + 200))))), (*libc.WChar)(unsafe.Pointer((*int16)(unsafe.Pointer(v34)))), image.Pt(v35, v50+450))
}
func sub_44F0F0(a1 int32, a2 int32) int32 {
	var (
		v2     int32
		v3     int32
		v4     *libc.WChar
		result int32
		v6     int32
		v7     int32
		v8     *uint16
		v9     int32
		v10    int32
		v11    int32
		v12    int32
		v13    [256]libc.WChar
		v14    [256]libc.WChar
	)
	v2 = nox_win_width / 2
	v3 = nox_win_height / 2
	nox_xxx_drawSetTextColor_434390(int32(nox_color_white_2523948))
	v4 = strMan.GetStringInFile("Noxworld.c:Stage", "C:\\NoxPost\\src\\client\\Gui\\GUIBrief.c")
	nox_wcscpy(&v13[0], v4)
	nox_wcscat(&v13[0], libc.CWString(" %d"))
	nox_swprintf(&v14[0], &v13[0], nox_gui_getQuestStage_450B10())
	nox_xxx_drawGetStringSize_43F840(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a2 + 200))))), &v14[0], &v10, &v11, 0)
	noxrend.DrawString(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a2 + 200))))), (*libc.WChar)(unsafe.Pointer((*int16)(unsafe.Pointer(&v14[0])))), image.Pt(v2-v10/2, v3+(v11-80)*2+v11-80))
	if *memmap.PtrUint32(6112660, 832464) != 0 {
		nox_xxx_drawGetStringSize_43F840(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a2 + 200))))), (*libc.WChar)(unsafe.Pointer(*(**uint16)(memmap.PtrOff(6112660, 832464)))), &v10, &v11, 0)
		noxrend.DrawString(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a2 + 200))))), (*libc.WChar)(unsafe.Pointer(*(**int16)(memmap.PtrOff(6112660, 832464)))), image.Pt(v2-v10/2, v3+(80-v11)*2+80-v11))
	}
	result = int32(nox_frame_xxx_2598000 / 30)
	if nox_frame_xxx_2598000%30 != 0 {
		if nox_xxx_aSpellphoneme_3_587000_123008 != 1 {
			return result
		}
	} else {
		result = 1
		if nox_xxx_aSpellphoneme_3_587000_123008 == 1 {
			nox_xxx_aSpellphoneme_3_587000_123008 = nox_frame_xxx_2598000 % 30
			return result
		}
		nox_xxx_aSpellphoneme_3_587000_123008 = 1
	}
	v12 = int32(nox_color_white_2523948)
	v6 = (nox_win_width - NOX_DEFAULT_WIDTH) / 2
	v7 = (nox_win_height - NOX_DEFAULT_HEIGHT) / 2
	v8 = (*uint16)(unsafe.Pointer(strMan.GetStringInFile("GeneralPrint:QuestSplash12", "C:\\NoxPost\\src\\client\\Gui\\GUIBrief.c")))
	nox_xxx_drawGetStringSize_43F840(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a2 + 200))))), (*libc.WChar)(unsafe.Pointer(v8)), &v10, nil, 0)
	v9 = v6 - v10/2 + 320
	nox_xxx_drawSetTextColor_434390(v12)
	return noxrend.DrawString(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a2 + 200))))), (*libc.WChar)(unsafe.Pointer((*int16)(unsafe.Pointer(v8)))), image.Pt(v9, v7+462))
}
func sub_44F300(a1 int32, a2 int32) int32 {
	var (
		v2     *byte
		v3     int32
		v4     int32
		v5     *uint16
		v6     *int16
		v7     *libc.WChar
		v8     int32
		v9     *int16
		v10    *libc.WChar
		v11    int32
		v12    *int16
		v13    *libc.WChar
		v14    int32
		v15    *int16
		v16    *libc.WChar
		v17    *uint16
		v18    *uint16
		v19    *uint16
		v20    *uint16
		v21    *uint16
		v22    *uint16
		result int32
		v24    int32
		v25    *uint16
		v26    int32
		v27    int32
		v28    int32
		v29    *uint16
		v30    *uint16
		v31    int32
		v32    *uint16
		v33    *uint16
		v34    int2
	)
	v2 = (*byte)(unsafe.Pointer(nox_draw_getViewport_437250()))
	sub_44E110()
	nox_client_resetScreenParticles_431510()
	nox_xxx_bookHideMB_45ACA0(1)
	sub_446780()
	v3 = (nox_win_width - NOX_DEFAULT_WIDTH) / 2
	v4 = (nox_win_height - NOX_DEFAULT_HEIGHT) / 2
	v5 = (*uint16)(unsafe.Pointer(strMan.GetStringInFile("GeneralPrint:QuestSplash1", "C:\\NoxPost\\src\\client\\Gui\\GUIBrief.c")))
	nox_xxx_drawGetStringSize_43F840(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a2 + 200))))), (*libc.WChar)(unsafe.Pointer(v5)), &v31, nil, 0)
	v29 = (*uint16)(unsafe.Pointer(uintptr(v3 - v31/2 + 320)))
	v30 = (*uint16)(unsafe.Pointer(uintptr(v4 + 20)))
	nox_xxx_drawSetTextColor_434390(int32(nox_color_black_2650656))
	v26 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 200))))
	v32 = (*uint16)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(v29))), -1))))
	noxrend.DrawString(unsafe.Pointer(uintptr(v26)), (*libc.WChar)(unsafe.Pointer((*int16)(unsafe.Pointer(v5)))), image.Pt(int32(uintptr(unsafe.Pointer(v29)))-1, v4+19))
	v33 = (*uint16)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(v29))), 1))))
	noxrend.DrawString(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a2 + 200))))), (*libc.WChar)(unsafe.Pointer((*int16)(unsafe.Pointer(v5)))), image.Pt(int32(uintptr(unsafe.Pointer(v29)))+1, v4+19))
	noxrend.DrawString(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a2 + 200))))), (*libc.WChar)(unsafe.Pointer((*int16)(unsafe.Pointer(v5)))), image.Pt(int32(uintptr(unsafe.Pointer(v32))), v4+21))
	noxrend.DrawString(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a2 + 200))))), (*libc.WChar)(unsafe.Pointer((*int16)(unsafe.Pointer(v5)))), image.Pt(int32(uintptr(unsafe.Pointer(v33))), v4+21))
	nox_xxx_drawSetTextColor_434390(int32(nox_color_orange_2614256))
	noxrend.DrawString(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a2 + 200))))), (*libc.WChar)(unsafe.Pointer((*int16)(unsafe.Pointer(v5)))), image.Pt(int32(uintptr(unsafe.Pointer(v29))), v4+20))
	v34.field_0 = v3 + 73
	v34.field_4 = v4 + 123
	sub_473A10((*uint32)(unsafe.Pointer(v2)), &v34, (*uint32)(unsafe.Pointer(uintptr(dword_5d4594_832492+12))))
	(cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_832492 + 300))), (*func(*byte, uint32))(nil)))(v2, dword_5d4594_832492)
	v34.field_0 = v3 + 109
	v34.field_4 = v4 + 76
	v6 = (*int16)(unsafe.Pointer(strMan.GetStringInFile("GeneralPrint:QuestSplash2a", "C:\\NoxPost\\src\\client\\Gui\\GUIBrief.c")))
	nox_xxx_drawSetTextColor_434390(int32(nox_color_orange_2614256))
	noxrend.DrawString(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_832484)))), (*libc.WChar)(unsafe.Pointer(v6)), image.Pt(v34.field_0, v34.field_4))
	nox_xxx_drawGetStringSize_43F840(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_832484)))), (*libc.WChar)(unsafe.Pointer((*uint16)(unsafe.Pointer(v6)))), &v27, nil, 0)
	v34.field_0 += v27 + 4
	nox_xxx_drawSetTextColor_434390(int32(nox_color_white_2523948))
	v7 = strMan.GetStringInFile("GeneralPrint:QuestSplash2b", "C:\\NoxPost\\src\\client\\Gui\\GUIBrief.c")
	nox_xxx_drawStringWrap_43FAF0(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_832484)))), v7, v34.field_0, v34.field_4, v3-v34.field_0+520, 0)
	v34.field_0 = v3 + 565
	v34.field_4 = v4 + 117
	sub_473A10((*uint32)(unsafe.Pointer(v2)), &v34, (*uint32)(unsafe.Pointer(uintptr(dword_5d4594_832496+12))))
	(cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_832496 + 300))), (*func(*byte, uint32))(nil)))(v2, dword_5d4594_832496)
	v29 = (*uint16)(unsafe.Pointer(strMan.GetStringInFile("GeneralPrint:QuestSplash3a", "C:\\NoxPost\\src\\client\\Gui\\GUIBrief.c")))
	nox_xxx_drawGetStringSize_43F840(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_832484)))), (*libc.WChar)(unsafe.Pointer(v29)), &v27, nil, 0)
	v30 = (*uint16)(unsafe.Pointer(strMan.GetStringInFile("GeneralPrint:QuestSplash3b", "C:\\NoxPost\\src\\client\\Gui\\GUIBrief.c")))
	nox_xxx_drawGetStringSize_43F840(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_832484)))), (*libc.WChar)(unsafe.Pointer(v30)), &v28, nil, 0)
	v8 = v3 + 520
	if v28+v27 <= 390 {
		v34.field_4 = v4 + 115
		v34.field_0 = v8 - v28 - v27 - 4
		nox_xxx_drawSetTextColor_434390(int32(nox_color_orange_2614256))
		noxrend.DrawString(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_832484)))), (*libc.WChar)(unsafe.Pointer((*int16)(unsafe.Pointer(v29)))), image.Pt(v34.field_0, v34.field_4))
		v34.field_0 = v8 - v28
		v34.field_4 = v4 + 115
		nox_xxx_drawSetTextColor_434390(int32(nox_color_white_2523948))
		noxrend.DrawString(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_832484)))), (*libc.WChar)(unsafe.Pointer((*int16)(unsafe.Pointer(v30)))), image.Pt(v34.field_0, v34.field_4))
	} else {
		v34.field_0 = v3 + 199
		v34.field_4 = v4 + 115
		nox_xxx_drawSetTextColor_434390(int32(nox_color_orange_2614256))
		noxrend.DrawString(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_832484)))), (*libc.WChar)(unsafe.Pointer((*int16)(unsafe.Pointer(v29)))), image.Pt(v34.field_0, v34.field_4))
		v34.field_0 += v27 + 4
		nox_xxx_drawSetTextColor_434390(int32(nox_color_white_2523948))
		nox_xxx_drawStringWrap_43FAF0(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_832484)))), (*libc.WChar)(unsafe.Pointer(v30)), v34.field_0, v34.field_4, v8-v34.field_0, 0)
	}
	v34.field_0 = v3 + 133
	v34.field_4 = v4 + 192
	sub_473A10((*uint32)(unsafe.Pointer(v2)), &v34, (*uint32)(unsafe.Pointer(uintptr(dword_5d4594_832504+12))))
	(cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_832504 + 300))), (*func(*byte, uint32))(nil)))(v2, dword_5d4594_832504)
	v34.field_0 = v3 + 157
	v34.field_4 = v4 + 156
	v9 = (*int16)(unsafe.Pointer(strMan.GetStringInFile("GeneralPrint:QuestSplash4a", "C:\\NoxPost\\src\\client\\Gui\\GUIBrief.c")))
	nox_xxx_drawSetTextColor_434390(int32(nox_color_orange_2614256))
	noxrend.DrawString(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_832484)))), (*libc.WChar)(unsafe.Pointer(v9)), image.Pt(v34.field_0, v34.field_4))
	nox_xxx_drawGetStringSize_43F840(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_832484)))), (*libc.WChar)(unsafe.Pointer((*uint16)(unsafe.Pointer(v9)))), &v27, nil, 0)
	v34.field_0 += v27 + 4
	nox_xxx_drawSetTextColor_434390(int32(nox_color_white_2523948))
	v10 = strMan.GetStringInFile("GeneralPrint:QuestSplash4b", "C:\\NoxPost\\src\\client\\Gui\\GUIBrief.c")
	nox_xxx_drawStringWrap_43FAF0(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_832484)))), v10, v34.field_0, v34.field_4, v4-v34.field_0+630, 0)
	v34.field_0 = v3 + 525
	v34.field_4 = v4 + 222
	sub_473A10((*uint32)(unsafe.Pointer(v2)), &v34, (*uint32)(unsafe.Pointer(uintptr(dword_5d4594_832500+12))))
	(cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_832500 + 300))), (*func(*byte, uint32))(nil)))(v2, dword_5d4594_832500)
	v29 = (*uint16)(unsafe.Pointer(strMan.GetStringInFile("GeneralPrint:QuestSplash7a", "C:\\NoxPost\\src\\client\\Gui\\GUIBrief.c")))
	nox_xxx_drawGetStringSize_43F840(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_832484)))), (*libc.WChar)(unsafe.Pointer(v29)), &v27, nil, 0)
	v30 = (*uint16)(unsafe.Pointer(strMan.GetStringInFile("GeneralPrint:QuestSplash7b", "C:\\NoxPost\\src\\client\\Gui\\GUIBrief.c")))
	nox_xxx_drawGetStringSize_43F840(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_832484)))), (*libc.WChar)(unsafe.Pointer(v30)), &v28, nil, 0)
	v11 = v3 + 500
	if v28+v27 <= 215 {
		v34.field_4 = v4 + 198
		v34.field_0 = v11 - v28 - v27 - 4
		nox_xxx_drawSetTextColor_434390(int32(nox_color_orange_2614256))
		noxrend.DrawString(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_832484)))), (*libc.WChar)(unsafe.Pointer((*int16)(unsafe.Pointer(v29)))), image.Pt(v34.field_0, v34.field_4))
		v34.field_0 = v11 - v28
		v34.field_4 = v4 + 198
		nox_xxx_drawSetTextColor_434390(int32(nox_color_white_2523948))
		noxrend.DrawString(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_832484)))), (*libc.WChar)(unsafe.Pointer((*int16)(unsafe.Pointer(v30)))), image.Pt(v34.field_0, v34.field_4))
	} else {
		v34.field_0 = v3 + 250
		v34.field_4 = v4 + 198
		nox_xxx_drawSetTextColor_434390(int32(nox_color_orange_2614256))
		noxrend.DrawString(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_832484)))), (*libc.WChar)(unsafe.Pointer((*int16)(unsafe.Pointer(v29)))), image.Pt(v34.field_0, v34.field_4))
		v34.field_0 += v27 + 4
		nox_xxx_drawSetTextColor_434390(int32(nox_color_white_2523948))
		nox_xxx_drawStringWrap_43FAF0(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_832484)))), (*libc.WChar)(unsafe.Pointer(v30)), v34.field_0, v34.field_4, v11-v34.field_0, 0)
	}
	v34.field_0 = v3 + 182
	v34.field_4 = v4 + 262
	sub_473A10((*uint32)(unsafe.Pointer(v2)), &v34, (*uint32)(unsafe.Pointer(uintptr(dword_5d4594_832528+12))))
	(cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_832528 + 300))), (*func(*byte, uint32))(nil)))(v2, dword_5d4594_832528)
	v34.field_0 = v3 + 201
	v34.field_4 = v4 + 251
	sub_473A10((*uint32)(unsafe.Pointer(v2)), &v34, (*uint32)(unsafe.Pointer(uintptr(dword_5d4594_832536+12))))
	(cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_832536 + 300))), (*func(*byte, uint32))(nil)))(v2, dword_5d4594_832536)
	v34.field_0 = v3 + 185
	v34.field_4 = v4 + 234
	sub_473A10((*uint32)(unsafe.Pointer(v2)), &v34, (*uint32)(unsafe.Pointer(uintptr(dword_5d4594_832532+12))))
	(cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_832532 + 300))), (*func(*byte, uint32))(nil)))(v2, dword_5d4594_832532)
	v34.field_0 = v3 + 221
	v34.field_4 = v4 + 240
	v12 = (*int16)(unsafe.Pointer(strMan.GetStringInFile("GeneralPrint:QuestSplash5a", "C:\\NoxPost\\src\\client\\Gui\\GUIBrief.c")))
	nox_xxx_drawSetTextColor_434390(int32(nox_color_orange_2614256))
	noxrend.DrawString(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_832484)))), (*libc.WChar)(unsafe.Pointer(v12)), image.Pt(v34.field_0, v34.field_4))
	nox_xxx_drawGetStringSize_43F840(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_832484)))), (*libc.WChar)(unsafe.Pointer((*uint16)(unsafe.Pointer(v12)))), &v27, nil, 0)
	v34.field_0 += v27 + 4
	nox_xxx_drawSetTextColor_434390(int32(nox_color_white_2523948))
	v13 = strMan.GetStringInFile("GeneralPrint:QuestSplash5b", "C:\\NoxPost\\src\\client\\Gui\\GUIBrief.c")
	nox_xxx_drawStringWrap_43FAF0(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_832484)))), v13, v34.field_0, v34.field_4, v3-v34.field_0+470, 0)
	v34.field_0 = v3 + 484
	v34.field_4 = v4 + 278
	sub_473A10((*uint32)(unsafe.Pointer(v2)), &v34, (*uint32)(unsafe.Pointer(uintptr(dword_5d4594_832516+12))))
	(cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_832516 + 300))), (*func(*byte, uint32))(nil)))(v2, dword_5d4594_832516)
	v34.field_0 = v3 + 503
	v34.field_4 = v4 + 303
	sub_473A10((*uint32)(unsafe.Pointer(v2)), &v34, (*uint32)(unsafe.Pointer(uintptr(dword_5d4594_832520+12))))
	(cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_832520 + 300))), (*func(*byte, uint32))(nil)))(v2, dword_5d4594_832520)
	v29 = (*uint16)(unsafe.Pointer(strMan.GetStringInFile("GeneralPrint:QuestSplash6a", "C:\\NoxPost\\src\\client\\Gui\\GUIBrief.c")))
	nox_xxx_drawGetStringSize_43F840(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_832484)))), (*libc.WChar)(unsafe.Pointer(v29)), &v27, nil, 0)
	v30 = (*uint16)(unsafe.Pointer(strMan.GetStringInFile("GeneralPrint:QuestSplash6b", "C:\\NoxPost\\src\\client\\Gui\\GUIBrief.c")))
	nox_xxx_drawGetStringSize_43F840(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_832484)))), (*libc.WChar)(unsafe.Pointer(v30)), &v28, nil, 0)
	v14 = v3 + 462
	if v28+v27 <= 350 {
		v34.field_4 = v4 + 286
		v34.field_0 = v14 - v28 - v27 - 4
		nox_xxx_drawSetTextColor_434390(int32(nox_color_orange_2614256))
		noxrend.DrawString(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_832484)))), (*libc.WChar)(unsafe.Pointer((*int16)(unsafe.Pointer(v29)))), image.Pt(v34.field_0, v34.field_4))
		v34.field_0 = v14 - v28
		v34.field_4 = v4 + 286
		nox_xxx_drawSetTextColor_434390(int32(nox_color_white_2523948))
		noxrend.DrawString(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_832484)))), (*libc.WChar)(unsafe.Pointer((*int16)(unsafe.Pointer(v30)))), image.Pt(v34.field_0, v34.field_4))
	} else {
		v34.field_0 = v3 + 113
		v34.field_4 = v4 + 286
		nox_xxx_drawSetTextColor_434390(int32(nox_color_orange_2614256))
		noxrend.DrawString(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_832484)))), (*libc.WChar)(unsafe.Pointer((*int16)(unsafe.Pointer(v29)))), image.Pt(v34.field_0, v34.field_4))
		v34.field_0 += v27 + 4
		nox_xxx_drawSetTextColor_434390(int32(nox_color_white_2523948))
		nox_xxx_drawStringWrap_43FAF0(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_832484)))), (*libc.WChar)(unsafe.Pointer(v30)), v34.field_0, v34.field_4, v14-v34.field_0, 0)
	}
	v34.field_0 = v3 + 186
	v34.field_4 = v4 + 333
	sub_473A10((*uint32)(unsafe.Pointer(v2)), &v34, (*uint32)(unsafe.Pointer(uintptr(dword_5d4594_832512+12))))
	(cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_832512 + 300))), (*func(*byte, uint32))(nil)))(v2, dword_5d4594_832512)
	v34.field_0 = v3 + 219
	v34.field_4 = v4 + 345
	sub_473A10((*uint32)(unsafe.Pointer(v2)), &v34, (*uint32)(unsafe.Pointer(uintptr(dword_5d4594_832508+12))))
	(cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_832508 + 300))), (*func(*byte, uint32))(nil)))(v2, dword_5d4594_832508)
	v34.field_0 = v3 + 220
	v34.field_4 = v4 + 322
	sub_473A10((*uint32)(unsafe.Pointer(v2)), &v34, (*uint32)(unsafe.Pointer(uintptr(dword_5d4594_832524+12))))
	(cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_832524 + 300))), (*func(*byte, uint32))(nil)))(v2, dword_5d4594_832524)
	v34.field_0 = v3 + 241
	v34.field_4 = v4 + 330
	v15 = (*int16)(unsafe.Pointer(strMan.GetStringInFile("GeneralPrint:QuestSplash8a", "C:\\NoxPost\\src\\client\\Gui\\GUIBrief.c")))
	nox_xxx_drawSetTextColor_434390(int32(nox_color_orange_2614256))
	noxrend.DrawString(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_832484)))), (*libc.WChar)(unsafe.Pointer(v15)), image.Pt(v34.field_0, v34.field_4))
	nox_xxx_drawGetStringSize_43F840(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_832484)))), (*libc.WChar)(unsafe.Pointer((*uint16)(unsafe.Pointer(v15)))), &v27, nil, 0)
	v34.field_0 += v27 + 4
	nox_xxx_drawSetTextColor_434390(int32(nox_color_white_2523948))
	v16 = strMan.GetStringInFile("GeneralPrint:QuestSplash8b", "C:\\NoxPost\\src\\client\\Gui\\GUIBrief.c")
	nox_xxx_drawStringWrap_43FAF0(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_832484)))), v16, v34.field_0, v34.field_4, v3-v34.field_0+550, 0)
	v17 = (*uint16)(unsafe.Pointer(strMan.GetStringInFile("GeneralPrint:QuestSplash9a", "C:\\NoxPost\\src\\client\\Gui\\GUIBrief.c")))
	v18 = (*uint16)(unsafe.Pointer(strMan.GetStringInFile("GeneralPrint:QuestSplash9b", "C:\\NoxPost\\src\\client\\Gui\\GUIBrief.c")))
	nox_xxx_drawGetStringSize_43F840(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_832484)))), (*libc.WChar)(unsafe.Pointer(v17)), &v27, nil, 0)
	nox_xxx_drawGetStringSize_43F840(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_832484)))), (*libc.WChar)(unsafe.Pointer(v18)), &v28, nil, 0)
	v34.field_0 = v3 - (v27+v28)/2 + 320
	v34.field_4 = v4 + 370
	nox_xxx_drawSetTextColor_434390(int32(nox_color_orange_2614256))
	noxrend.DrawString(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_832484)))), (*libc.WChar)(unsafe.Pointer((*int16)(unsafe.Pointer(v17)))), image.Pt(v34.field_0, v34.field_4))
	v34.field_0 += v27 + 4
	nox_xxx_drawSetTextColor_434390(int32(nox_color_white_2523948))
	noxrend.DrawString(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_832484)))), (*libc.WChar)(unsafe.Pointer((*int16)(unsafe.Pointer(v18)))), image.Pt(v34.field_0, v34.field_4))
	v19 = (*uint16)(unsafe.Pointer(strMan.GetStringInFile("GeneralPrint:QuestSplash10a", "C:\\NoxPost\\src\\client\\Gui\\GUIBrief.c")))
	v20 = (*uint16)(unsafe.Pointer(strMan.GetStringInFile("GeneralPrint:QuestSplash10b", "C:\\NoxPost\\src\\client\\Gui\\GUIBrief.c")))
	nox_xxx_drawGetStringSize_43F840(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_832484)))), (*libc.WChar)(unsafe.Pointer(v19)), &v27, nil, 0)
	nox_xxx_drawGetStringSize_43F840(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_832484)))), (*libc.WChar)(unsafe.Pointer(v20)), &v28, nil, 0)
	v34.field_0 = v3 - (v28+v27)/2 + 320
	v34.field_4 = v4 + 395
	nox_xxx_drawSetTextColor_434390(int32(nox_color_orange_2614256))
	noxrend.DrawString(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_832484)))), (*libc.WChar)(unsafe.Pointer((*int16)(unsafe.Pointer(v19)))), image.Pt(v34.field_0, v34.field_4))
	v34.field_0 += v27 + 4
	nox_xxx_drawSetTextColor_434390(int32(nox_color_white_2523948))
	noxrend.DrawString(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_832484)))), (*libc.WChar)(unsafe.Pointer((*int16)(unsafe.Pointer(v20)))), image.Pt(v34.field_0, v34.field_4))
	v21 = (*uint16)(unsafe.Pointer(strMan.GetStringInFile("GeneralPrint:QuestSplash11a", "C:\\NoxPost\\src\\client\\Gui\\GUIBrief.c")))
	v22 = (*uint16)(unsafe.Pointer(strMan.GetStringInFile("GeneralPrint:QuestSplash11b", "C:\\NoxPost\\src\\client\\Gui\\GUIBrief.c")))
	nox_xxx_drawGetStringSize_43F840(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_832484)))), (*libc.WChar)(unsafe.Pointer(v21)), &v27, nil, 0)
	nox_xxx_drawGetStringSize_43F840(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_832484)))), (*libc.WChar)(unsafe.Pointer(v22)), &v28, nil, 0)
	v34.field_0 = v3 - (v28+v27)/2 + 320
	v34.field_4 = v4 + 420
	nox_xxx_drawSetTextColor_434390(int32(nox_color_orange_2614256))
	noxrend.DrawString(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_832484)))), (*libc.WChar)(unsafe.Pointer((*int16)(unsafe.Pointer(v21)))), image.Pt(v34.field_0, v34.field_4))
	v34.field_0 += v27 + 4
	nox_xxx_drawSetTextColor_434390(int32(nox_color_white_2523948))
	noxrend.DrawString(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_832484)))), (*libc.WChar)(unsafe.Pointer((*int16)(unsafe.Pointer(v22)))), image.Pt(v34.field_0, v34.field_4))
	result = int32(nox_frame_xxx_2598000 / 30)
	if nox_frame_xxx_2598000%30 != 0 {
		if *memmap.PtrUint32(0x587000, 123012) != 1 {
			return result
		}
	} else {
		result = 1
		if *memmap.PtrUint32(0x587000, 123012) == 1 {
			*memmap.PtrUint32(0x587000, 123012) = nox_frame_xxx_2598000 % 30
			return result
		}
		*memmap.PtrUint32(0x587000, 123012) = 1
	}
	v24 = int32(nox_color_white_2523948)
	v25 = (*uint16)(unsafe.Pointer(strMan.GetStringInFile("GeneralPrint:QuestSplash12", "C:\\NoxPost\\src\\client\\Gui\\GUIBrief.c")))
	nox_xxx_drawGetStringSize_43F840(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_832484)))), (*libc.WChar)(unsafe.Pointer(v25)), &v27, nil, 0)
	v34.field_4 = v4 + 450
	v34.field_0 = v3 - v27/2 + 320
	nox_xxx_drawSetTextColor_434390(v24)
	return noxrend.DrawString(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a2 + 200))))), (*libc.WChar)(unsafe.Pointer((*int16)(unsafe.Pointer(v25)))), image.Pt(v34.field_0, v34.field_4))
}
func nox_xxx_clientQuestWinScreen_450770(a1 int32) int32 {
	var (
		v1 uint32
		v2 *uint8
		v3 *uint16
		v4 *uint32
		v5 *uint16
		v6 *uint16
		v7 *uint16
		v8 *uint16
		v9 int32
	)
	libc.MemSet(memmap.PtrOff(6112660, 832364), 0, 96)
	*memmap.PtrUint32(6112660, 832356) = 0
	v1 = 0
	*memmap.PtrUint32(6112660, 832356) = uint32(*(*uint16)(unsafe.Pointer(uintptr(a1 + 2))))
	*memmap.PtrUint32(6112660, 831228) = uint32(*(*uint16)(unsafe.Pointer(uintptr(a1 + 4))))
	v2 = (*uint8)(memmap.PtrOff(6112660, 832368))
	v3 = (*uint16)(unsafe.Pointer(uintptr(a1 + 6)))
	for {
		if int32(*v3) != 0 {
			*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), -int(unsafe.Sizeof(uint32(0))*1)))) = uint32(uintptr(unsafe.Pointer(noxServer.getPlayerByID(int32(*v3)))))
			*(*uint16)(unsafe.Pointer(v2)) = *(*uint16)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint16(0))*4))
			*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v2))), unsafe.Sizeof(uint16(0))*1))) = *(*uint16)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint16(0))*1))
			*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v2))), unsafe.Sizeof(uint16(0))*2))) = *(*uint16)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint16(0))*2))
			*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v2))), unsafe.Sizeof(uint16(0))*3))) = *(*uint16)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint16(0))*3))
			*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), unsafe.Sizeof(uint32(0))*2))) = *(*uint32)(unsafe.Pointer((*uint16)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint16(0))*5))))
			v1++
		}
		v2 = (*uint8)(unsafe.Add(unsafe.Pointer(v2), 16))
		v3 = (*uint16)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint16(0))*7))
		if int32(uintptr(unsafe.Pointer(v2))) >= int32(uintptr(memmap.PtrOff(6112660, 832464))) {
			break
		}
	}
	libc.Sort(memmap.PtrOff(6112660, 832364), v1, 16, sub_450960)
	if dword_5d4594_832476 == 0 {
		v4 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&nox_wnd_briefing_831232)))).ChildByID(1010)))
		v5 = (*uint16)(unsafe.Pointer(strMan.GetStringInFile("GUIBrief.c:GeneratorsDestroyed", "C:\\NoxPost\\src\\client\\Gui\\GUIBrief.c")))
		nox_xxx_drawGetStringSize_43F840(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*59)))), (*libc.WChar)(unsafe.Pointer(v5)), &a1, nil, 0)
		if a1 > *(*int32)(unsafe.Pointer(&dword_5d4594_832476)) {
			dword_5d4594_832476 = uint32(a1)
		}
		v6 = (*uint16)(unsafe.Pointer(strMan.GetStringInFile("GUIBrief.c:Kills", "C:\\NoxPost\\src\\client\\Gui\\GUIBrief.c")))
		nox_xxx_drawGetStringSize_43F840(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*59)))), (*libc.WChar)(unsafe.Pointer(v6)), &a1, nil, 0)
		if a1 > *(*int32)(unsafe.Pointer(&dword_5d4594_832476)) {
			dword_5d4594_832476 = uint32(a1)
		}
		v7 = (*uint16)(unsafe.Pointer(strMan.GetStringInFile("GUIBrief.c:numSecretsFound", "C:\\NoxPost\\src\\client\\Gui\\GUIBrief.c")))
		nox_xxx_drawGetStringSize_43F840(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*59)))), (*libc.WChar)(unsafe.Pointer(v7)), &a1, nil, 0)
		if a1 > *(*int32)(unsafe.Pointer(&dword_5d4594_832476)) {
			dword_5d4594_832476 = uint32(a1)
		}
		v8 = (*uint16)(unsafe.Pointer(strMan.GetStringInFile("GUIBrief.c:TotalScore", "C:\\NoxPost\\src\\client\\Gui\\GUIBrief.c")))
		nox_xxx_drawGetStringSize_43F840(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*59)))), (*libc.WChar)(unsafe.Pointer(v8)), &a1, nil, 0)
		v9 = int32(dword_5d4594_832476)
		if a1 > *(*int32)(unsafe.Pointer(&dword_5d4594_832476)) {
			v9 = a1
			dword_5d4594_832476 = uint32(a1)
		}
		if v9 > 85 {
			dword_5d4594_832476 = 85
		}
	}
	return nox_client_lockScreenBriefing_450160(254, 1, 1)
}
func nox_client_showQuestBriefing2_450980(a1 int32, a2 int32) int32 {
	var (
		v2     *byte
		v3     *libc.WChar
		result int32
	)
	dword_5d4594_832480 = 0
	nox_client_resetScreenParticles_431510()
	nox_xxx_bookHideMB_45ACA0(1)
	sub_446780()
	v2 = (*byte)(unsafe.Pointer(nox_xxx_gLoadImg((*byte)(unsafe.Pointer(uintptr(a1 + 5))))))
	sub_450AD0(v2)
	if libc.StrLen((*byte)(unsafe.Pointer(uintptr(a1+37)))) != 0 {
		v3 = strMan.GetStringInFile((*byte)(unsafe.Pointer(uintptr(a1+37))), "C:\\NoxPost\\src\\client\\Gui\\GUIBrief.c")
		sub_450AF0(int32(uintptr(unsafe.Pointer(v3))))
	} else {
		sub_450AF0(int32(uintptr(memmap.PtrOff(6112660, 832544))))
	}
	nox_gui_setQuestStage_450B00(int32(*(*uint16)(unsafe.Pointer(uintptr(a1 + 2)))))
	if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 4))))&2 != 0 {
		dword_5d4594_832480 = 1
	}
	result = a2
	if a2 != 0 {
		result = nox_client_lockScreenBriefing_450160(254, 1, 2)
	}
	return result
}
func nox_client_showQuestBriefing_450A30(a1 int32, a2 int32) int32 {
	var (
		v2     *byte
		v3     *libc.WChar
		result int32
	)
	dword_5d4594_832480 = 0
	nox_client_resetScreenParticles_431510()
	nox_xxx_bookHideMB_45ACA0(1)
	sub_446780()
	v2 = (*byte)(unsafe.Pointer(nox_xxx_gLoadImg((*byte)(unsafe.Pointer(uintptr(a1 + 5))))))
	sub_450AD0(v2)
	if libc.StrLen((*byte)(unsafe.Pointer(uintptr(a1+37)))) != 0 {
		v3 = strMan.GetStringInFile((*byte)(unsafe.Pointer(uintptr(a1+37))), "C:\\NoxPost\\src\\client\\Gui\\GUIBrief.c")
		sub_450AF0(int32(uintptr(unsafe.Pointer(v3))))
	} else {
		sub_450AF0(int32(uintptr(memmap.PtrOff(6112660, 832548))))
	}
	nox_gui_setQuestStage_450B00(int32(*(*uint16)(unsafe.Pointer(uintptr(a1 + 2)))))
	result = a2
	if a2 != 0 {
		result = nox_client_lockScreenBriefing_450160(254, 1, 4)
	}
	return result
}
