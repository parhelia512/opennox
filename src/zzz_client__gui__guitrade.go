package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"unsafe"
)

func sub_4C09D0() int32 {
	var (
		v0  *uint32
		v2  *libc.WChar
		v3  *libc.WChar
		v4  *libc.WChar
		v5  *libc.WChar
		v6  *libc.WChar
		v7  *uint32
		v8  *uint32
		v9  *libc.WChar
		v10 *libc.WChar
		v11 *libc.WChar
		v12 *libc.WChar
		v13 *libc.WChar
		v14 *libc.WChar
		v15 *uint8
		v16 *uint8
		v17 int32
		v18 *uint8
		v19 *uint8
		v20 int32
		v21 *libc.WChar
	)
	v0 = (*uint32)(unsafe.Pointer(newWindowFromFile("Trade.wnd", unsafe.Pointer(cgoFuncAddr(sub_4C0C90)))))
	dword_5d4594_1320940 = uint32(uintptr(unsafe.Pointer(v0)))
	if v0 == nil {
		return 0
	}
	(*nox_window)(unsafe.Pointer(v0)).SetAllFuncs(func(arg1 int32, arg2 int32, arg3 int32, arg4 int32) int32 {
		return sub_4C0630(arg1, uint32(arg2), uint32(arg3))
	}, func(arg1 *nox_window, arg2 unsafe.Pointer) int32 {
		return sub_4C0D00()
	}, nil)
	v2 = strMan.GetStringInFile("TradeMain", "C:\\NoxPost\\src\\client\\Gui\\GUITrade.c")
	nox_xxx_wndWddSetTooltip_46B000((*nox_window_data)(unsafe.Pointer(uintptr(dword_5d4594_1320940+36))), v2)
	v3 = (*libc.WChar)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1320940)))).ChildByID(3702)))
	v4 = strMan.GetStringInFile("TradePlayerName", "C:\\NoxPost\\src\\client\\Gui\\GUITrade.c")
	nox_xxx_wndWddSetTooltip_46B000((*nox_window_data)(unsafe.Pointer((*libc.WChar)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(libc.WChar(0))*18)))), v4)
	v5 = (*libc.WChar)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1320940)))).ChildByID(3703)))
	v6 = strMan.GetStringInFile("TradeVendorName", "C:\\NoxPost\\src\\client\\Gui\\GUITrade.c")
	nox_xxx_wndWddSetTooltip_46B000((*nox_window_data)(unsafe.Pointer((*libc.WChar)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(libc.WChar(0))*18)))), v6)
	v7 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1320940)))).ChildByID(3704)))
	(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v7)))))).setTooltipFunc(unsafe.Pointer(cgoFuncAddr(sub_4C1120)))
	v8 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1320940)))).ChildByID(3705)))
	(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v8)))))).setTooltipFunc(unsafe.Pointer(cgoFuncAddr(sub_4C1120)))
	v9 = (*libc.WChar)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1320940)))).ChildByID(3708)))
	v10 = strMan.GetStringInFile("TradePlayerAccept", "C:\\NoxPost\\src\\client\\Gui\\GUITrade.c")
	nox_xxx_wndWddSetTooltip_46B000((*nox_window_data)(unsafe.Pointer((*libc.WChar)(unsafe.Add(unsafe.Pointer(v9), unsafe.Sizeof(libc.WChar(0))*18)))), v10)
	v11 = (*libc.WChar)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1320940)))).ChildByID(3709)))
	v12 = strMan.GetStringInFile("TradeVendorAccept", "C:\\NoxPost\\src\\client\\Gui\\GUITrade.c")
	nox_xxx_wndWddSetTooltip_46B000((*nox_window_data)(unsafe.Pointer((*libc.WChar)(unsafe.Add(unsafe.Pointer(v11), unsafe.Sizeof(libc.WChar(0))*18)))), v12)
	v13 = (*libc.WChar)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1320940)))).ChildByID(3710)))
	v14 = strMan.GetStringInFile("TradeCancel", "C:\\NoxPost\\src\\client\\Gui\\GUITrade.c")
	nox_xxx_wndWddSetTooltip_46B000((*nox_window_data)(unsafe.Pointer((*libc.WChar)(unsafe.Add(unsafe.Pointer(v13), unsafe.Sizeof(libc.WChar(0))*18)))), v14)
	(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1320940))))).Hide()
	nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1320940))))), 0)
	v15 = (*uint8)(memmap.PtrOff(6112660, 1319288))
	for {
		v16 = v15
		v17 = 2
		for {
			*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v16))), -int(unsafe.Sizeof(uint32(0))*1)))) = 0
			*(*uint32)(unsafe.Pointer(v16)) = 0
			v16 = (*uint8)(unsafe.Add(unsafe.Pointer(v16), 280))
			v17--
			if v17 == 0 {
				break
			}
		}
		v15 = (*uint8)(unsafe.Add(unsafe.Pointer(v15), 140))
		if int32(uintptr(unsafe.Pointer(v15))) >= int32(uintptr(memmap.PtrOff(6112660, 1319568))) {
			break
		}
	}
	v18 = (*uint8)(memmap.PtrOff(6112660, 1320312))
	for {
		v19 = v18
		v20 = 2
		for {
			*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v19))), -int(unsafe.Sizeof(uint32(0))*1)))) = 0
			*(*uint32)(unsafe.Pointer(v19)) = 0
			v19 = (*uint8)(unsafe.Add(unsafe.Pointer(v19), 280))
			v20--
			if v20 == 0 {
				break
			}
		}
		v18 = (*uint8)(unsafe.Add(unsafe.Pointer(v18), 140))
		if int32(uintptr(unsafe.Pointer(v18))) >= int32(uintptr(memmap.PtrOff(6112660, 1320592))) {
			break
		}
	}
	v21 = strMan.GetStringInFile("TotalValueLabel", "C:\\NoxPost\\src\\client\\Gui\\GUITrade.c")
	nox_wcscpy((*libc.WChar)(memmap.PtrOff(6112660, 1319972)), v21)
	*memmap.PtrUint32(6112660, 1320188) = 0
	*memmap.PtrUint32(6112660, 1320192) = 0
	*memmap.PtrUint32(6112660, 1320196) = uint32(nox_win_width)
	*memmap.PtrUint32(6112660, 1320200) = uint32(nox_win_height)
	*memmap.PtrUint32(6112660, 1320220) = uint32(nox_win_width)
	*memmap.PtrUint32(6112660, 1320224) = uint32(nox_win_height)
	*memmap.PtrUint32(6112660, 1320204) = 0
	*memmap.PtrUint32(6112660, 1320208) = 0
	*memmap.PtrUint32(6112660, 1320164) = uint32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg("TradeBase"))))
	*memmap.PtrUint32(6112660, 1320168) = uint32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg("TradeLeftAcceptPushed"))))
	*memmap.PtrUint32(6112660, 1320172) = uint32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg("TradeLeftAcceptLit"))))
	*memmap.PtrUint32(6112660, 1320176) = uint32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg("TradeRightAcceptLit"))))
	*memmap.PtrUint32(6112660, 1320180) = uint32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg("TradeCancelLit"))))
	*memmap.PtrUint32(6112660, 1320184) = uint32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg("TradeGold"))))
	return 1
}
func sub_4C15D0(a1 int32) int32 {
	var (
		result int32
		v2     int32
		v3     *uint8
		v4     int32
		v5     int32
		v6     *uint8
		v7     int32
		v8     *uint8
		v9     int32
		v10    int32
		v11    *libc.WChar
		v12    *uint8
	)
	result = int32(dword_5d4594_1320964)
	v2 = 0
	v12 = nil
	if dword_5d4594_1320964 != 0 {
		v3 = (*uint8)(memmap.PtrOff(6112660, 1319284))
		for 2 != 0 {
			v4 = 0
			v5 = int32(uintptr(unsafe.Pointer(v3)))
			for {
				if sub_4C1760(v5, int32(*(*uint16)(unsafe.Pointer(uintptr(a1 + 2))))) != 0 {
					v6 = (*uint8)(memmap.PtrOff(6112660, uint32((v2+v4*2)*140)+0x142174))
					goto LABEL_18
				}
				v4++
				v5 += 280
				if v4 >= 2 {
					break
				}
			}
			v3 = (*uint8)(unsafe.Add(unsafe.Pointer(v3), 140))
			v2++
			if int32(uintptr(unsafe.Pointer(v3))) < int32(uintptr(memmap.PtrOff(6112660, 1319564))) {
				continue
			}
			break
		}
		v7 = 0
		v8 = (*uint8)(memmap.PtrOff(6112660, 1320308))
		for 2 != 0 {
			v9 = 0
			v10 = int32(uintptr(unsafe.Pointer(v8)))
			for {
				if sub_4C1760(v10, int32(*(*uint16)(unsafe.Pointer(uintptr(a1 + 2))))) != 0 {
					v12 = (*uint8)(memmap.PtrOff(6112660, uint32((v7+v9*2)*140)+0x142574))
					goto LABEL_17
				}
				v9++
				v10 += 280
				if v9 >= 2 {
					break
				}
			}
			v8 = (*uint8)(unsafe.Add(unsafe.Pointer(v8), 140))
			v7++
			if int32(uintptr(unsafe.Pointer(v8))) < int32(uintptr(memmap.PtrOff(6112660, 1320588))) {
				continue
			}
			break
		}
		v11 = strMan.GetStringInFile("TradeGUIItemNotFound", "C:\\NoxPost\\src\\client\\Gui\\GUITrade.c")
		nox_xxx_printCentered_445490(v11)
	LABEL_17:
		v6 = v12
	LABEL_18:
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v6))), unsafe.Sizeof(uint32(0))*34))) -= *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v6))), unsafe.Sizeof(uint32(0))*34))) / *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v6))), unsafe.Sizeof(uint32(0))*1)))
		sub_4C1710(int32(uintptr(unsafe.Pointer(v6))), int32(*(*uint16)(unsafe.Pointer(uintptr(a1 + 2)))))
		result = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v6))), unsafe.Sizeof(uint32(0))*1))) - 1)
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v6))), unsafe.Sizeof(uint32(0))*1))) = uint32(result)
		if result == 0 {
			result = nox_xxx_spriteDelete_45A4B0((*nox_drawable)(unsafe.Pointer(*(**uint64)(unsafe.Pointer(v6)))))
			*(*uint32)(unsafe.Pointer(v6)) = 0
		}
	}
	return result
}
