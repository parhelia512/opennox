package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"unsafe"
)

func nox_xxx_guiSummonCreatureLoad_4C1D80() int32 {
	var (
		v0 *libc.WChar
		v1 *uint32
		v2 *byte
		v3 *uint32
		v4 *uint32
		v5 int32
		v6 *uint8
	)
	*memmap.PtrUint32(6112660, 1321004) = 0
	*memmap.PtrUint32(6112660, 1321000) = 0xFFFFFF6F
	dword_5d4594_1320988 = uint32(nox_win_width - 95)
	dword_5d4594_1320992 = 0xFFFFFF6F
	dword_5d4594_1321032 = uint32(uintptr(unsafe.Pointer(nox_window_new(nil, 8, nox_win_width-95, -145, 87, 115, nil))))
	(*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1321032)))).SetAllFuncs(func(arg1 int32, arg2 int32, arg3 int32, arg4 int32) int32 {
		return sub_4C2BD0()
	}, func(arg1 *nox_window, arg2 unsafe.Pointer) int32 {
		return sub_4C24A0()
	}, nil)
	dword_5d4594_1321036 = uint32(uintptr(unsafe.Pointer(nox_window_new((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1321032))))), 136, 5, 38, 76, 76, nil))))
	(*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1321036)))).SetAllFuncs(func(arg1 int32, arg2 int32, arg3 int32, arg4 int32) int32 {
		return nox_xxx_wndSummonProc_4C2B10((*uint32)(unsafe.Pointer(uintptr(arg1))), uint32(arg2), uint32(arg3))
	}, func(arg1 *nox_window, arg2 unsafe.Pointer) int32 {
		return nox_xxx_guiDrawSummonBox_4C1FE0((*uint32)(unsafe.Pointer(arg1)))
	}, unsafe.Pointer(cgoFuncAddr(sub_4C2C20)))
	v0 = strMan.GetStringInFile("ToolTipSummon", "C:\\NoxPost\\src\\Client\\Gui\\guisumn.c")
	nox_xxx_wndWddSetTooltip_46B000((*nox_window_data)(unsafe.Pointer(uintptr(dword_5d4594_1321036+36))), v0)
	*memmap.PtrUint32(6112660, 1320996) = uint32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg("CreatureCageBottom"))))
	v1 = (*uint32)(unsafe.Pointer(nox_window_new((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1321036))))), 160, 0, 0, 1, 1, nil)))
	v2 = (*byte)(unsafe.Pointer(nox_xxx_gLoadImg("CreatureCageTop")))
	nox_xxx_wndSetIcon_46AE60(int32(uintptr(unsafe.Pointer(v1))), int32(uintptr(unsafe.Pointer(v2))))
	nox_xxx_wndSetOffsetMB_46AE40(int32(uintptr(unsafe.Pointer(v1))), -5, -38)
	v3 = (*uint32)(unsafe.Pointer(nox_window_new((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1321032))))), 8, 19, 0, 48, 39, nil)))
	(*nox_window)(unsafe.Pointer(v3)).SetAllFuncs(func(arg1 int32, arg2 int32, arg3 int32, arg4 int32) int32 {
		return sub_4C2BE0()
	}, func(arg1 *nox_window, arg2 unsafe.Pointer) int32 {
		return sub_4C24A0()
	}, nil)
	*memmap.PtrUint32(6112660, 1321008) = uint32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg("CreatureCageHuntButtonLit"))))
	*memmap.PtrUint32(6112660, 1321012) = uint32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg("CreatureCageHuntButton"))))
	*memmap.PtrUint32(6112660, 1321016) = uint32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg("CreatureCageGuardButtonLit"))))
	*memmap.PtrUint32(6112660, 1321020) = uint32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg("CreatureCageGuardButton"))))
	dword_5d4594_1321024 = uint32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg("CreatureCageEscortButtonLit"))))
	*memmap.PtrUint32(6112660, 1321028) = uint32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg("CreatureCageEscortButton"))))
	v4 = (*uint32)(unsafe.Pointer(nox_window_new(nil, 168, int32(dword_5d4594_1320988+27), int32(dword_5d4594_1320992+12), 34, 34, nil)))
	dword_5d4594_1321040 = uint32(uintptr(unsafe.Pointer(v4)))
	v5 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*11)))
	*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v5))), 1)) |= 1
	*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*11)) = uint32(v5)
	(*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1321040)))).SetAllFuncs(func(arg1 int32, arg2 int32, arg3 int32, arg4 int32) int32 {
		return nox_xxx_wndSummonBigButtonProc_4C24B0(arg1, arg2, uint32(arg3))
	}, nil, unsafe.Pointer(cgoFuncAddr(sub_4C2CE0)))
	nox_xxx_wndSetIcon_46AE60(*(*int32)(unsafe.Pointer(&dword_5d4594_1321040)), *memmap.PtrInt32(6112660, 1321028))
	sub_46AEC0(*(*int32)(unsafe.Pointer(&dword_5d4594_1321040)), *(*int32)(unsafe.Pointer(&dword_5d4594_1321024)))
	nox_xxx_wndSetIconLit_46AEA0(*(*int32)(unsafe.Pointer(&dword_5d4594_1321040)), *(*int32)(unsafe.Pointer(&dword_5d4594_1321024)))
	nox_xxx_wndSetOffsetMB_46AE40(*(*int32)(unsafe.Pointer(&dword_5d4594_1321040)), -27, -12)
	*memmap.PtrUint8(6112660, 1321200) = 0
	(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1321032))))).Hide()
	(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1321040))))).Hide()
	v6 = (*uint8)(memmap.PtrOff(6112660, 1321060))
	for {
		*(*uint32)(unsafe.Pointer(v6)) = 0
		v6 = (*uint8)(unsafe.Add(unsafe.Pointer(v6), 32))
		if int32(uintptr(unsafe.Pointer(v6))) >= int32(uintptr(memmap.PtrOff(6112660, 1321188))) {
			break
		}
	}
	sub_4C2BF0()
	dword_5d4594_1321044 = 0
	dword_5d4594_1321204 = 0
	dword_5d4594_1321196 = 0
	return 1
}
func nox_xxx_wndSummonCreateList_4C2560(a1 *int2) {
	var (
		v1  **byte
		v2  *uint16
		v3  int32
		v4  bool
		v5  int32
		v6  int32
		v7  int32
		v8  int32
		i   int32
		v10 *uint32
		v12 int32
		v13 int32
	)
	nox_xxx_screenWidth_587000_184452 = 0
	v1 = (**byte)(memmap.PtrOff(0x587000, 184344))
	for {
		if v1 != (**byte)(memmap.PtrOff(0x587000, 184352)) {
			v2 = (*uint16)(unsafe.Pointer(strMan.GetStringInFile(*v1, "C:\\NoxPost\\src\\Client\\Gui\\guisumn.c")))
			nox_xxx_drawGetStringSize_43F840(nil, (*libc.WChar)(unsafe.Pointer(v2)), &v12, &v13, nox_win_width)
			if nox_xxx_screenWidth_587000_184452 < uint32(v12) {
				nox_xxx_screenWidth_587000_184452 = uint32(v12)
			}
		}
		v1 = (**byte)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof((*byte)(nil))*1))
		if int32(uintptr(unsafe.Pointer(v1))) >= int32(uintptr(memmap.PtrOff(0x587000, 184368))) {
			break
		}
	}
	nox_xxx_screenWidth_587000_184452 += 8
	v3 = nox_xxx_guiFontHeightMB_43F320(nil) + 2
	v5 = int32(uint32(a1.field_0) - nox_xxx_screenWidth_587000_184452/2)
	v4 = v5 < 0
	v12 = int32(uint32(a1.field_0) - nox_xxx_screenWidth_587000_184452/2)
	v6 = v3*5 + 12
	if v4 {
		v5 = 0
	} else {
		if nox_xxx_screenWidth_587000_184452+uint32(v5) < uint32(nox_win_width) {
			goto LABEL_11
		}
		v5 = int32(uint32(nox_win_width) - nox_xxx_screenWidth_587000_184452 - 1)
	}
	v12 = v5
LABEL_11:
	v7 = a1.field_4 - v6/2
	v13 = v7
	if v7 < 0 {
		v7 = 0
	LABEL_15:
		v13 = v7
		goto LABEL_16
	}
	if v6+v7 >= nox_win_height {
		v7 = nox_win_width - v6 - 1
		goto LABEL_15
	}
LABEL_16:
	dword_5d4594_1321044 = uint32(uintptr(unsafe.Pointer(nox_window_new(nil, 40, v5, v7, *(*int32)(unsafe.Pointer(&nox_xxx_screenWidth_587000_184452)), v3*5+12, nil))))
	(*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1321044)))).SetAllFuncs(nil, func(arg1 *nox_window, arg2 unsafe.Pointer) int32 {
		return sub_4C26F0(unsafe.Pointer(arg1))
	}, nil)
	nox_xxx_wndShowModalMB((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1321044))))))
	v8 = 0
	for i = 0; i < 6; i++ {
		if i != 2 {
			v10 = (*uint32)(unsafe.Pointer(nox_window_new((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1321044))))), 8, 0, v8, *(*int32)(unsafe.Pointer(&nox_xxx_screenWidth_587000_184452)), v3+1, nil)))
			(*nox_window)(unsafe.Pointer(v10)).SetAllFuncs(func(arg1 int32, arg2 int32, arg3 int32, arg4 int32) int32 {
				return nox_xxx_clientOrderCreature_4C2A60(arg1, uint32(arg2))
			}, func(arg1 *nox_window, arg2 unsafe.Pointer) int32 {
				return sub_4C27F0((*uint32)(unsafe.Pointer(arg1)))
			}, nil)
			*(*uint32)(unsafe.Add(unsafe.Pointer(v10), unsafe.Sizeof(uint32(0))*8)) = uint32(i)
			v8 += v3 + 2
		}
	}
	clientPlaySoundSpecial(791, 100)
}
func sub_4C27F0(a1 *uint32) int32 {
	var (
		result int32
		v2     *uint16
		v4     int32
		v5     int32
		v6     int32
		v7     int32
		v8     int32
		v9     *int16
		v10    int32
		v11    int32
		v12    int32
	)
	if dword_5d4594_1321208 == 0 {
		dword_5d4594_1321208 = uint32(nox_xxx_getNameId_4E3AA0(CString("CarnivorousPlant")))
	}
	if dword_5d4594_1321204 != 0 || (func() bool {
		result = 1
		return *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*8)) != 1
	}()) {
		v2 = (*uint16)(unsafe.Pointer(strMan.GetStringInFile(*(**byte)(memmap.PtrOff(0x587000, *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*8))*4+0x2D018)), "C:\\NoxPost\\src\\Client\\Gui\\guisumn.c")))
		nox_client_wndGetPosition_46AA60((*nox_window)(unsafe.Pointer(a1)), (*uint32)(unsafe.Pointer(&v11)), (*uint32)(unsafe.Pointer(&v10)))
		nox_xxx_drawGetStringSize_43F840(nil, (*libc.WChar)(unsafe.Pointer(v2)), &v12, nil, 0)
		var mpos nox_point = getMousePos()
		nox_xxx_guiFontHeightMB_43F320(nil)
		v4 = int32((nox_xxx_screenWidth_587000_184452-uint32(v12))/2 + 1)
		if nox_xxx_wndPointInWnd_46AAB0(a1, mpos.x, mpos.y) {
			sub_4C2A00(v11+v4, v10+3, int32(nox_color_yellow_2589772), int32(nox_color_black_2650656), (*int16)(unsafe.Pointer(v2)))
			if *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*8)) != *memmap.PtrUint32(0x587000, 184552) {
				*memmap.PtrUint32(0x587000, 184552) = *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*8))
				clientPlaySoundSpecial(920, 100)
				return 1
			}
			return 1
		}
		if dword_5d4594_1321204 != 0 {
			if sub_4C2DD0(*(*int32)(unsafe.Pointer(&dword_5d4594_1321204))) != 0 {
				sub_4C2A00(v11+v4, v10+3, int32(nox_color_white_2523948), int32(nox_color_black_2650656), (*int16)(unsafe.Pointer(v2)))
				return 1
			}
			v5 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*8)))
			if v5 != 4 && v5 != 5 {
				sub_4C2A00(v11+v4, v10+3, int32(nox_color_white_2523948), int32(nox_color_black_2650656), (*int16)(unsafe.Pointer(v2)))
				return 1
			}
			v6 = int32(*memmap.PtrUint32(0x85B3FC, 956))
			v9 = (*int16)(unsafe.Pointer(v2))
			v8 = int32(nox_color_black_2650656)
		} else {
			v7 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*8)))
			if v7 != 4 && v7 != 5 {
				sub_4C2A00(v11+v4, v10+3, int32(nox_color_blue_2650684), int32(nox_color_black_2650656), (*int16)(unsafe.Pointer(v2)))
				return 1
			}
			v9 = (*int16)(unsafe.Pointer(v2))
			if sub_4C2E00() == 0 {
				sub_4C2A00(v11+v4, v10+3, *memmap.PtrInt32(0x85B3FC, 956), int32(nox_color_black_2650656), (*int16)(unsafe.Pointer(v2)))
				return 1
			}
			v6 = int32(nox_color_blue_2650684)
			v8 = int32(nox_color_black_2650656)
		}
		sub_4C2A00(v11+v4, v10+3, v6, v8, v9)
		return 1
	}
	return result
}
func sub_4C2CE0() int32 {
	var (
		v0 *libc.WChar
		v2 *libc.WChar
		v3 *libc.WChar
	)
	switch *memmap.PtrUint32(0x587000, 184448) {
	case 3:
		v3 = strMan.GetStringInFile("ccs:GUARD", "C:\\NoxPost\\src\\Client\\Gui\\guisumn.c")
		nox_xxx_cursorSetTooltip_4776B0(v3)
	case 4:
		v2 = strMan.GetStringInFile("ccs:ESCORT", "C:\\NoxPost\\src\\Client\\Gui\\guisumn.c")
		nox_xxx_cursorSetTooltip_4776B0(v2)
		return 1
	case 5:
		v0 = strMan.GetStringInFile("ccs:HUNT", "C:\\NoxPost\\src\\Client\\Gui\\guisumn.c")
		nox_xxx_cursorSetTooltip_4776B0(v0)
		return 1
	}
	return 1
}
