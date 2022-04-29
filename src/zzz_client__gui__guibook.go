package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	noxflags "github.com/noxworld-dev/opennox/v1/common/flags"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"image"
	"unsafe"
)

var obj_5d4594_1046620 float2 = float2{}

func nox_xxx_bookShowMB_45AD70(a1 int32) {
	var (
		result uint32
		v2     *libc.WChar
	)
	result = uint32(nox_xxx_guiCursor_477600())
	if result != 0 {
		return
	}
	if !(nox_xxx_playerAnimCheck_4372B0() == 0 || (func() uint32 {
		result = uint32(bool2int(noxflags.HasGame(noxflags.GameModeCoop)))
		return result
	}()) == 0) {
		return
	}
	if *memmap.PtrUint32(0x8531A0, 2576) == 0 || nox_xxx_guiSpellSortList_45ADF0(int32(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(0x8531A0, 2576) + 2251))))) != 0 {
		nox_xxx_book_45B010(a1)
	} else {
		v2 = strMan.GetStringInFile("EmptyBook", "C:\\NoxPost\\src\\Client\\Gui\\guibook.c")
		nox_xxx_printCentered_445490(v2)
		clientPlaySoundSpecial(925, 100)
	}
}
func nox_xxx_bookDrawList_45BD40(a1 int32) int32 {
	var (
		v1          int32
		v2          int32
		v3          int32
		v4          int32
		v5          int32
		v6          int32
		v7          int32
		v8          int32
		v9          int8
		v10         *uint16
		v11         *int16
		v12         *uint16
		v13         *uint16
		v14         int32
		v15         *libc.WChar
		v16         int32
		v17         int32
		v18         int32
		v19         *libc.WChar
		v20         int32
		v21         *libc.WChar
		v22         *libc.WChar
		v23         int32
		v24         *uint16
		v25         int32
		v26         *uint16
		v27         int32
		v28         *uint16
		v29         *uint16
		v30         int32
		v31         int32
		v32         int32
		v33         *uint16
		v34         int32
		v35         *uint16
		v36         *uint16
		v37         int32
		v38         int16
		v39         *uint16
		v40         int32
		v41         int32
		v42         *uint16
		v43         int32
		v44         int32
		v45         *libc.WChar
		v46         int32
		v47         int32
		v48         *uint16
		v49         *uint16
		v50         *uint16
		v51         *uint16
		v52         *libc.WChar
		v53         *uint16
		v55         int32
		v56         int32
		v57         int32
		v58         int32
		v59         int32
		v60         int32
		v61         int32
		v62         int32
		v63         int32
		v64         int32
		v65         int32
		v66         int32
		v67         int32
		v68         int32
		WideCharStr [4]libc.WChar
		v70         int32
		v71         [256]libc.WChar
	)
	nox_gui_getWindowOffs_46AA20((*nox_window)(unsafe.Pointer(uintptr(a1))), (*uint32)(unsafe.Pointer(&v66)), (*uint32)(unsafe.Pointer(&v67)))
	v1 = v66 - 24
	v2 = v67 - 76
	v3 = nox_xxx_guiFontHeightMB_43F320(nil)
	dword_5d4594_1046656 = uint32(v3 + 2)
	nox_xxx_drawSetTextColor_434390(*memmap.PtrInt32(6112660, 1046880))
	nox_client_drawImageAt_47D2C0((*nox_video_bag_image_t)(unsafe.Pointer(uintptr(*memmap.PtrInt32(6112660, 1046856)))), v1, v2)
	if dword_5d4594_1046872 != 0 {
		nox_client_drawImageAt_47D2C0((*nox_video_bag_image_t)(unsafe.Pointer(uintptr(*memmap.PtrInt32(6112660, 1046660)))), v1, v2)
	} else {
		nox_client_drawImageAt_47D2C0((*nox_video_bag_image_t)(unsafe.Pointer(uintptr(*memmap.PtrInt32(6112660, 1046644)))), v1, v2)
	}
	if dword_5d4594_1046868 == 3 {
		nox_video_drawAnimatedImageOrCursorAt_4BE6D0(*(*int32)(unsafe.Pointer(&dword_5d4594_1046928)), v1, v2)
		goto LABEL_75
	}
	if dword_5d4594_1046868 == 2 {
		nox_video_drawAnimatedImageOrCursorAt_4BE6D0(*(*int32)(unsafe.Pointer(&dword_5d4594_1046924)), v1, v2)
		goto LABEL_75
	}
	if nox_xxx_aNox_cfg_0_587000_132132 != 0 {
		v4 = int32(dword_5d4594_1046936)
		*(*uint32)(unsafe.Pointer(&WideCharStr[0])) = 141/dword_5d4594_1046656 - 1
		if dword_5d4594_1046936 == 0 {
			(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1046944))))).Hide()
			v4 = int32(dword_5d4594_1046936)
		}
		v5 = 0
		v6 = v66 + 78
		v7 = v67 + 19
		if *(*uint32)(unsafe.Pointer(&WideCharStr[0]))*2 > 0 {
			for {
				v8 = int32(uint32(v5) + *(*uint32)(unsafe.Pointer(&WideCharStr[0]))*2*uint32(v4))
				if uint32(v8) >= uint32(*memmap.PtrInt32(6112660, 1047508))-dword_5d4594_1047512 {
					break
				}
				if uint32(v5) == *(*uint32)(unsafe.Pointer(&WideCharStr[0])) {
					v6 = v66 + 199
					v7 = v67 + 19
				}
				nox_xxx_drawSetTextColor_434390(*memmap.PtrInt32(6112660, 1046880))
				v9 = int8(*(*uint8)(unsafe.Pointer(uintptr(dword_5d4594_1047516 + 2251))))
				if dword_5d4594_1046868 == 1 {
					if int32(v9) == 2 && !nox_xxx_spellIsEnabled_424B70(int32(*memmap.PtrUint32(6112660, uint32(v8*4)+0xFF9B0)+74)) {
						nox_xxx_drawSetTextColor_434390(*memmap.PtrInt32(6112660, 1046884))
					}
					v10 = (*uint16)(unsafe.Pointer(uintptr(nox_xxx_guiCreatureGetName_427240(int32(*memmap.PtrUint32(6112660, uint32(v8*4)+0xFF9B0))))))
					nox_xxx_drawGetStringSize_43F840(nil, (*libc.WChar)(unsafe.Pointer(v10)), &v65, nil, 128)
					v62 = v7
					v60 = v6 - v65/2
					v11 = (*int16)(unsafe.Pointer(uintptr(nox_xxx_guiCreatureGetName_427240(int32(*memmap.PtrUint32(6112660, uint32(v8*4)+0xFF9B0))))))
				} else if int32(v9) != 0 {
					if !nox_xxx_spellIsEnabled_424B70(int32(*memmap.PtrUint32(6112660, uint32(v8*4)+0xFF9B0))) {
						nox_xxx_drawSetTextColor_434390(*memmap.PtrInt32(6112660, 1046884))
					}
					v13 = (*uint16)(unsafe.Pointer(nox_xxx_spellTitle_424930(int32(*memmap.PtrUint32(6112660, uint32(v8*4)+0xFF9B0)))))
					nox_xxx_drawGetStringSize_43F840(nil, (*libc.WChar)(unsafe.Pointer(v13)), &v65, nil, 128)
					v62 = v7
					v60 = v6 - v65/2
					v11 = (*int16)(unsafe.Pointer(nox_xxx_spellTitle_424930(int32(*memmap.PtrUint32(6112660, uint32(v8*4)+0xFF9B0)))))
				} else {
					v12 = (*uint16)(unsafe.Pointer(uintptr(nox_xxx_abilityGetName_0_425260(int32(*memmap.PtrUint32(6112660, uint32(v8*4)+0xFF9B0))))))
					nox_xxx_drawGetStringSize_43F840(nil, (*libc.WChar)(unsafe.Pointer(v12)), &v65, nil, 128)
					v62 = v7
					v60 = v6 - v65/2
					v11 = (*int16)(unsafe.Pointer(uintptr(nox_xxx_abilityGetName_0_425260(int32(*memmap.PtrUint32(6112660, uint32(v8*4)+0xFF9B0))))))
				}
				noxrend.DrawString(nil, (*libc.WChar)(unsafe.Pointer(v11)), image.Pt(v60, v62))
				v7 += int32(dword_5d4594_1046656)
				if uint32(func() int32 {
					p := &v5
					*p++
					return *p
				}()) >= *(*uint32)(unsafe.Pointer(&WideCharStr[0]))*2 {
					break
				}
				v4 = int32(dword_5d4594_1046936)
			}
		}
		goto LABEL_75
	}
	if dword_5d4594_1046868 == 1 {
		v14 = v67 + 51
		v15 = strMan.GetStringInFile("Size", "C:\\NoxPost\\src\\Client\\Gui\\guibook.c")
		nox_swprintf(&v71[0], libc.CWString("%s "), v15)
		v16 = int32(nox_xxx_guideGetUnitSize_427460(int32(*memmap.PtrUint32(6112660, dword_5d4594_1046932*4+0xFF9B0)))) - 1
		if v16 != 0 {
			v17 = v16 - 1
			if v17 != 0 {
				if v17 == 2 {
					v18 = 76
					v19 = strMan.GetStringInFile("Large", "C:\\NoxPost\\src\\Client\\Gui\\guibook.c")
					nox_wcscat(&v71[0], v19)
					v20 = 0
				} else {
					v18 = int32(*(*uint32)(unsafe.Pointer(&WideCharStr[0])))
					v20 = int32(*(*uint32)(unsafe.Pointer(&WideCharStr[0])))
				}
			} else {
				v18 = 38
				v21 = strMan.GetStringInFile("Medium", "C:\\NoxPost\\src\\Client\\Gui\\guibook.c")
				nox_wcscat(&v71[0], v21)
				v20 = 0
			}
		} else {
			v18 = 38
			v22 = strMan.GetStringInFile("Small", "C:\\NoxPost\\src\\Client\\Gui\\guibook.c")
			nox_wcscat(&v71[0], v22)
			v20 = 19
		}
		if int32(*(*uint8)(unsafe.Pointer(uintptr(dword_5d4594_1047516 + 2251)))) == 2 && (*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1047516 + 4232))) != 0 || noxflags.HasGame(noxflags.GameOnline) && !noxflags.HasGame(noxflags.GameModeQuest)) {
			nox_xxx_drawGetStringSize_43F840(nil, &v71[0], &v70, nil, 0)
			nox_xxx_drawStringWrap_43FAF0(nil, &v71[0], (108-v70)/2+v66+24, v14, 128, 0)
		}
		v23 = v14 + v3 + 2
		v24 = (*uint16)(unsafe.Pointer(uintptr(nox_xxx_guiCreatureGetName_427240(int32(*memmap.PtrUint32(6112660, dword_5d4594_1046932*4+0xFF9B0))))))
		nox_xxx_bookGetStringSize_43FA80(nil, (*libc.WChar)(unsafe.Pointer(v24)), &v70, &v68, 108)
		if v68 <= v3 {
			v25 = (108-v70)/2 + v66 + 24
		} else {
			v25 = v66 + 24
		}
		v55 = v25
		v26 = (*uint16)(unsafe.Pointer(uintptr(nox_xxx_guiCreatureGetName_427240(int32(*memmap.PtrUint32(6112660, dword_5d4594_1046932*4+0xFF9B0))))))
		nox_xxx_bookDrawString_43FA80_43FD80(nil, (*libc.WChar)(unsafe.Pointer(v26)), v55, v23, 128, 0)
		v63 = v23 + v20 + v68 + 2
		v61 = (108-v18)/2 + v66 + 24
		v27 = nox_xxx_bookGetCreatureImg_427400(int32(*memmap.PtrUint32(6112660, dword_5d4594_1046932*4+0xFF9B0)))
		nox_client_drawImageAt_47D2C0((*nox_video_bag_image_t)(unsafe.Pointer(uintptr(v27))), v61, v63)
		v28 = (*uint16)(unsafe.Pointer(uintptr(nox_xxx_guideGetDescById_4272E0(int32(*memmap.PtrUint32(6112660, dword_5d4594_1046932*4+0xFF9B0))))))
		v29 = v28
		v30 = v67 + 52
		if v28 != nil {
			nox_xxx_drawGetStringSize_43F840(nil, (*libc.WChar)(unsafe.Pointer(v28)), nil, (*int32)(unsafe.Pointer(&WideCharStr[0])), 92)
			v31 = int32(*(*uint32)(unsafe.Pointer(&WideCharStr[0])))
			v32 = v66 + 153
			goto LABEL_53
		}
	} else {
		if int32(*(*uint8)(unsafe.Pointer(uintptr(dword_5d4594_1047516 + 2251)))) == 0 {
			sub_425450(int32(*memmap.PtrUint32(6112660, dword_5d4594_1046932*4+0xFF9B0)))
			v33 = (*uint16)(unsafe.Pointer(uintptr(nox_xxx_abilityGetName_0_425260(int32(*memmap.PtrUint32(6112660, dword_5d4594_1046932*4+0xFF9B0))))))
			nox_xxx_bookGetStringSize_43FA80(nil, (*libc.WChar)(unsafe.Pointer(v33)), &v70, &v68, 108)
			if v68 <= v3 {
				v34 = (108-v70)/2 + v66 + 24
			} else {
				v34 = v66 + 24
			}
			v58 = v67 + 53
			v56 = v34
			v35 = (*uint16)(unsafe.Pointer(uintptr(nox_xxx_abilityGetName_0_425260(int32(*memmap.PtrUint32(6112660, dword_5d4594_1046932*4+0xFF9B0))))))
			nox_xxx_bookDrawString_43FA80_43FD80(nil, (*libc.WChar)(unsafe.Pointer(v35)), v56, v58, 128, 0)
			v36 = (*uint16)(unsafe.Pointer(uintptr(sub_4252F0(int32(*memmap.PtrUint32(6112660, dword_5d4594_1046932*4+0xFF9B0))))))
			v29 = v36
			v30 = v67 + 52
			if v36 == nil {
				goto LABEL_75
			}
			nox_xxx_drawGetStringSize_43F840(nil, (*libc.WChar)(unsafe.Pointer(v36)), nil, (*int32)(unsafe.Pointer(&WideCharStr[0])), 92)
			v31 = int32(*(*uint32)(unsafe.Pointer(&WideCharStr[0])))
			v32 = v66 + 153
		LABEL_53:
			v37 = (141-v31)/2 + v67 + 17
			if v37 > v30 {
				goto LABEL_73
			}
			goto LABEL_74
		}
		v38 = int16(uint16(nox_xxx_spellFlags_424A70(int32(*memmap.PtrUint32(6112660, dword_5d4594_1046932*4+0xFF9B0)))))
		v39 = (*uint16)(unsafe.Pointer(nox_xxx_spellTitle_424930(int32(*memmap.PtrUint32(6112660, dword_5d4594_1046932*4+0xFF9B0)))))
		nox_xxx_bookGetStringSize_43FA80(nil, (*libc.WChar)(unsafe.Pointer(v39)), &v65, &v68, 108)
		if v68 <= v3 {
			v40 = (108-v65)/2 + v66 + 24
		} else {
			v40 = v66 + 24
		}
		v41 = v67 + 53
		v59 = v67 + 53
		v57 = v40
		v42 = (*uint16)(unsafe.Pointer(nox_xxx_spellTitle_424930(int32(*memmap.PtrUint32(6112660, dword_5d4594_1046932*4+0xFF9B0)))))
		nox_xxx_bookDrawString_43FA80_43FD80(nil, (*libc.WChar)(unsafe.Pointer(v42)), v57, v59, 128, 0)
		v43 = v41 + v68 + 2
		v44 = nox_xxx_spellManaCost_4249A0(int32(*memmap.PtrUint32(6112660, dword_5d4594_1046932*4+0xFF9B0)), 1)
		v45 = strMan.GetStringInFile("ManaCost", "C:\\NoxPost\\src\\Client\\Gui\\guibook.c")
		nox_swprintf(&v71[0], libc.CWString("%s "), v45)
		if v44 != 0 {
			compat_itow(v44, &WideCharStr[0], 10)
			nox_wcscat(&v71[0], &WideCharStr[0])
		} else if nox_xxx_spellHasFlags_424A50(int32(*memmap.PtrUint32(6112660, dword_5d4594_1046932*4+0xFF9B0)), 0x800000) {
			nox_wcscat(&v71[0], libc.CWString("0"))
		} else {
			nox_wcscat(&v71[0], libc.CWString("*"))
		}
		nox_xxx_drawGetStringSize_43F840(nil, &v71[0], &v65, nil, 0)
		v46 = v43 + 2
		nox_xxx_drawStringWrap_43FAF0(nil, &v71[0], (108-v65)/2+v66+24, v46, 128, 0)
		v47 = v46 + v3 + 2
		if int32(v38)&256 != 0 {
			v48 = (*uint16)(unsafe.Pointer(strMan.GetStringInFile("SpellInstant", "C:\\NoxPost\\src\\Client\\Gui\\guibook.c")))
			nox_xxx_drawGetStringSize_43F840(nil, (*libc.WChar)(unsafe.Pointer(v48)), &v65, nil, 0)
			nox_xxx_drawStringWrap_43FAF0(nil, (*libc.WChar)(unsafe.Pointer(v48)), (108-v65)/2+v66+24, v47, 128, 0)
			v47 += v3
		}
		if int32(v38)&4 != 0 {
			v49 = (*uint16)(unsafe.Pointer(strMan.GetStringInFile("SpellTargeted", "C:\\NoxPost\\src\\Client\\Gui\\guibook.c")))
			nox_xxx_drawGetStringSize_43F840(nil, (*libc.WChar)(unsafe.Pointer(v49)), &v65, nil, 0)
			nox_xxx_drawStringWrap_43FAF0(nil, (*libc.WChar)(unsafe.Pointer(v49)), (108-v65)/2+v66+24, v47, 128, 0)
			v47 += v3
		}
		if int32(v38)&8 != 0 {
			v50 = (*uint16)(unsafe.Pointer(strMan.GetStringInFile("SpellAtLocation", "C:\\NoxPost\\src\\Client\\Gui\\guibook.c")))
			nox_xxx_drawGetStringSize_43F840(nil, (*libc.WChar)(unsafe.Pointer(v50)), &v65, nil, 0)
			nox_xxx_drawStringWrap_43FAF0(nil, (*libc.WChar)(unsafe.Pointer(v50)), (108-v65)/2+v66+24, v47, 128, 0)
			v47 += v3
		}
		if int32(v38)&32 != 0 {
			v51 = (*uint16)(unsafe.Pointer(strMan.GetStringInFile("SpellHostile", "C:\\NoxPost\\src\\Client\\Gui\\guibook.c")))
			nox_xxx_drawGetStringSize_43F840(nil, (*libc.WChar)(unsafe.Pointer(v51)), &v65, nil, 0)
			nox_xxx_drawStringWrap_43FAF0(nil, (*libc.WChar)(unsafe.Pointer(v51)), (108-v65)/2+v66+24, v47, 128, 0)
			v47 += v3
		}
		v64 = int32(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1047516 + *memmap.PtrUint32(6112660, dword_5d4594_1046932*4+0xFF9B0)*4 + 3696))))
		v52 = strMan.GetStringInFile("PowerLevel", "C:\\NoxPost\\src\\Client\\Gui\\guibook.c")
		nox_swprintf(&v71[0], v52, v64)
		nox_xxx_drawGetStringSize_43F840(nil, &v71[0], &v65, nil, 0)
		nox_xxx_drawStringWrap_43FAF0(nil, &v71[0], (108-v65)/2+v66+24, v3+v47, 128, 0)
		v53 = (*uint16)(unsafe.Pointer(nox_xxx_spellDescription_424A30(int32(*memmap.PtrUint32(6112660, dword_5d4594_1046932*4+0xFF9B0)))))
		v29 = v53
		v30 = v67 + 52
		if v53 != nil {
			nox_xxx_drawGetStringSize_43F840(nil, (*libc.WChar)(unsafe.Pointer(v53)), nil, (*int32)(unsafe.Pointer(&WideCharStr[0])), 92)
			v32 = v66 + 153
			v37 = int32((141-*(*uint32)(unsafe.Pointer(&WideCharStr[0])))/2 + uint32(v67) + 17)
			if v37 > v30 {
			LABEL_73:
				v37 = v30
			}
		LABEL_74:
			nox_xxx_drawStringWrap_43FAF0(nil, (*libc.WChar)(unsafe.Pointer(v29)), v32, v37, 92, 0)
			goto LABEL_75
		}
	}
LABEL_75:
	if nox_xxx_aNox_cfg_0_587000_132132 == 0 && *(*int32)(unsafe.Pointer(&dword_5d4594_1046932)) >= *memmap.PtrInt32(6112660, 1047508)-*(*int32)(unsafe.Pointer(&dword_5d4594_1047512))-1 {
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1046948))))).Hide()
	}
	return 1
}
func nox_xxx_book_45CF00(a1 *uint32) int32 {
	var (
		v1 *libc.WChar
		v3 *libc.WChar
		v4 *libc.WChar
	)
	if *a1 == 1310 {
		if int32(*(*uint8)(unsafe.Pointer(uintptr(dword_5d4594_1047516 + 2251)))) == 0 {
			v3 = strMan.GetStringInFile("ToolTipAbilityTab", "C:\\NoxPost\\src\\Client\\Gui\\guibook.c")
			nox_xxx_cursorSetTooltip_4776B0(v3)
			return 1
		}
		v4 = strMan.GetStringInFile("ToolTipSpellTab", "C:\\NoxPost\\src\\Client\\Gui\\guibook.c")
		nox_xxx_cursorSetTooltip_4776B0(v4)
	} else if *a1 == 1320 {
		v1 = strMan.GetStringInFile("ToolTipGuideTab", "C:\\NoxPost\\src\\Client\\Gui\\guibook.c")
		nox_xxx_cursorSetTooltip_4776B0(v1)
		return 1
	}
	return 1
}
func nox_xxx_bookDrawFn_45C7D0(a1 *uint32) int32 {
	var (
		v1  int32
		v3  int32
		v4  int32
		v5  int32
		v6  int32
		v7  int32
		v8  int32
		v9  int32
		v10 int32
		v11 float64
		v12 int32
		v13 int32
		v14 int32
		v15 int32
		v16 int32
		v17 int8
		v18 int8
		v19 int8
		v20 int8
		v21 int32
		v22 int32
		v23 int32
		v24 int32
		v25 int32
	)
	if dword_5d4594_1046652 != 0 {
		v1 = 3
	} else {
		v1 = 0
	}
	if dword_5d4594_1047520 == 0 {
		return 1
	}
	if dword_5d4594_1046648 != 0 && uint32(nox_xxx_bookGet_430B40_get_mouse_prev_seq())-dword_5d4594_1046648 < (nox_gameFPS*2) {
		return 1
	}
	nox_client_wndGetPosition_46AA60((*nox_window)(unsafe.Pointer(a1)), (*uint32)(unsafe.Pointer(&v25)), (*uint32)(unsafe.Pointer(&v24)))
	if dword_5d4594_1046648 != 0 {
		v3 = 50
		for {
			v19 = int8(randomIntMinMax(3, 6))
			v17 = int8(randomIntMinMax(2, 5))
			v16 = randomIntMinMax(-10, -1)
			v15 = randomIntMinMax(-10, 10)
			v4 = randomIntMinMax(0, 30)
			v13 = v24 + v4
			v5 = randomIntMinMax(0, 30)
			nox_client_newScreenParticle_431540(v1, v25+v5, v13, v15, v16, 1, v17, v19, 2, 1)
			v3--
			if v3 == 0 {
				break
			}
		}
		clientPlaySoundSpecial(795, 100)
		dword_5d4594_1046648 = 0
	}
	v6 = 2
	for {
		v20 = int8(randomIntMinMax(2, 4))
		v18 = int8(randomIntMinMax(1, 2))
		v7 = randomIntMinMax(0, 30)
		v14 = v24 + v7
		v8 = randomIntMinMax(0, 30)
		nox_client_newScreenParticle_431540(v1, v25+v8, v14, 0, 0, 0, v18, v20, 1, 1)
		v6--
		if v6 == 0 {
			break
		}
	}
	v22 = v24
	v21 = v25
	if dword_5d4594_1046652 == 1 {
		v9 = int32(uintptr(nox_xxx_spellGetAbilityIcon_425310(*(*int32)(unsafe.Pointer(&dword_5d4594_1047524)), 0)))
	} else {
		v9 = int32(uintptr(nox_xxx_spellIcon_424A90(*(*int32)(unsafe.Pointer(&dword_5d4594_1047524)))))
	}
	nox_client_drawImageAt_47D2C0((*nox_video_bag_image_t)(unsafe.Pointer(uintptr(v9))), v21, v22)
	*(*float32)(unsafe.Pointer(&dword_5d4594_1046636)) = *(*float32)(unsafe.Pointer(&dword_5d4594_1046636)) + obj_5d4594_1046620.field_0
	*(*float32)(unsafe.Pointer(&dword_5d4594_1046640)) = *(*float32)(unsafe.Pointer(&dword_5d4594_1046640)) + obj_5d4594_1046620.field_4
	if float64(*memmap.PtrInt32(6112660, 1046668)) <= float64(*(*float32)(unsafe.Pointer(&dword_5d4594_1046636))) && float64(*memmap.PtrInt32(6112660, 1046672)) <= float64(*(*float32)(unsafe.Pointer(&dword_5d4594_1046640))) {
	LABEL_26:
		nox_xxx_book_45DBE0(*(*unsafe.Pointer)(memmap.PtrOff(6112660, 1046676)), *(*int32)(unsafe.Pointer(&dword_5d4594_1047524)), *(*int32)(unsafe.Pointer(&dword_5d4594_1046852)))
		sub_45D810()
		goto LABEL_27
	}
	if float64(*(*float32)(unsafe.Pointer(&dword_5d4594_1046636))) > float64(*mem_getFloatPtr(6112660, *memmap.PtrUint32(6112660, 1046628)*8+0xFF8A4)) {
		v10 = int32(*memmap.PtrUint32(6112660, 1046628) + 1)
		*memmap.PtrUint32(6112660, 1046628) = uint32(v10)
		if v10 < *memmap.PtrInt32(6112660, 1046680) {
			if v10 <= *memmap.PtrInt32(6112660, 1046680)-1 {
				obj_5d4594_1046620.field_0 = *mem_getFloatPtr(6112660, uint32(v10*8)+0xFF8A4) - *mem_getFloatPtr(6112660, uint32(v10*8)+0xFF89C)
				obj_5d4594_1046620.field_4 = *mem_getFloatPtr(6112660, uint32(v10*8)+0xFF8A8) - *mem_getFloatPtr(6112660, uint32(v10*8)+0xFF8A0)
				nox_xxx_utilNormalizeVector_509F20(&obj_5d4594_1046620)
				if nox_win_width < 1000 {
					if nox_win_width < 750 {
						v11 = 6.0
					} else {
						v11 = 8.0
					}
				} else {
					v11 = 10.0
				}
				obj_5d4594_1046620.field_0 = float32(float64(obj_5d4594_1046620.field_0) * v11)
				obj_5d4594_1046620.field_4 = float32(float64(obj_5d4594_1046620.field_4) * v11)
			}
			goto LABEL_27
		}
		goto LABEL_26
	}
LABEL_27:
	v23 = int(*(*float32)(unsafe.Pointer(&dword_5d4594_1046640)))
	v12 = int(*(*float32)(unsafe.Pointer(&dword_5d4594_1046636)))
	(*nox_window)(unsafe.Pointer(a1)).SetPos(image.Pt(v12, v23))
	return 1
}
func sub_45D870() {
	var (
		v1  int32
		v3  int32
		v4  int32
		v5  int32
		v6  int32
		v7  int32
		v8  int8
		v9  int32
		v10 int32
	)
	if dword_5d4594_1047520 != 0 {
		if dword_5d4594_1046652 != 1 {
			v1 = 0
		} else {
			v1 = 3
		}
		v9 = (int32(*memmap.PtrUint32(6112660, 1046668)) - int(*(*float32)(unsafe.Pointer(&dword_5d4594_1046636)))) / 50
		v10 = (int32(*memmap.PtrUint32(6112660, 1046672)) - int(*(*float32)(unsafe.Pointer(&dword_5d4594_1046640)))) / 50
		v3 = int(*(*float32)(unsafe.Pointer(&dword_5d4594_1046636)))
		v4 = int(*(*float32)(unsafe.Pointer(&dword_5d4594_1046640)))
		v5 = 50
		for {
			v8 = int8(randomIntMinMax(3, 4))
			v7 = v4 - randomIntMinMax(0, 30)
			v6 = randomIntMinMax(0, 30)
			nox_client_newScreenParticle_431540(v1, v3+v6, v7, 0, 0, 1, v8, 0, 0, 1)
			v3 += v9
			v4 += v10
			v5--
			if v5 == 0 {
				break
			}
		}
		nox_xxx_book_45DBE0(*(*unsafe.Pointer)(memmap.PtrOff(6112660, 1046676)), *(*int32)(unsafe.Pointer(&dword_5d4594_1047524)), *(*int32)(unsafe.Pointer(&dword_5d4594_1046852)))
		sub_45D810()
	}
}
