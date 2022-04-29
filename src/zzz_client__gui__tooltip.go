package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"unsafe"
)

func nox_xxx_clientAskInfoMb_4BF050(a1p *nox_drawable) *libc.WChar {
	var (
		a1     *libc.WChar = (*libc.WChar)(unsafe.Pointer(a1p))
		v1     *int32
		v2     int32
		v3     *libc.WChar
		v4     *uint32
		v5     *libc.WChar
		result *libc.WChar
		v7     int32
		v8     int32
		v9     *libc.WChar
		v10    int32
		v11    *libc.WChar
		v12    *libc.WChar
		v13    int32
		v14    *libc.WChar
		v15    int32
		v16    *libc.WChar
		v17    *libc.WChar
		v18    int32
		v19    int32
		v20    *libc.WChar
		v21    *libc.WChar
		v22    *libc.WChar
		v23    *libc.WChar
		v24    int32
		v25    int32
		v26    *libc.WChar
		v27    *libc.WChar
		v28    *libc.WChar
		v29    int32
		v30    *libc.WChar
		v31    *libc.WChar
		v32    *libc.WChar
		v33    int32
		v34    *libc.WChar
		v35    *libc.WChar
	)
	*memmap.PtrUint16(6112660, 1317000) = 0
	nox_wcscpy((*libc.WChar)(memmap.PtrOff(6112660, 1317000)), (*libc.WChar)(memmap.PtrOff(6112660, 1319048)))
	v1 = (*int32)(unsafe.Pointer(a1))
	if a1 == nil {
		return (*libc.WChar)(memmap.PtrOff(6112660, 1317000))
	}
	v2 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a1))), unsafe.Sizeof(uint32(0))*28))))
	if (uint32(v2) & 0x13001000) == 0 {
		if (v2 & 256) == 0 {
			result = nox_get_thing_pretty_name(int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a1))), unsafe.Sizeof(uint32(0))*27)))))
			if result != nil {
				return result
			}
			return (*libc.WChar)(memmap.PtrOff(6112660, 1317000))
		}
		v18 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a1))), unsafe.Sizeof(uint32(0))*29))))
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a1))), 0)) = 243
		if v18&1 != 0 {
			v19 = *(*int32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(int32(0))*108))
			if v19 == 0 {
				*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a1))), 0)) = 226
				*(*uint16)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(&a1))), 1)))) = uint16(nox_xxx_netGetUnitCodeCli_578B00(int32(uintptr(unsafe.Pointer(v1)))))
				*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a1))), unsafe.Sizeof((*libc.WChar)(nil))-1)) = 1
				*(*int32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(int32(0))*108)) = 137
				goto LABEL_93
			}
			if v19 != 137 {
				if strMan.Lang() != 6 {
					v22 = strMan.GetStringInFile("BookOf", "C:\\NoxPost\\src\\client\\Gui\\ToolTip.c")
					nox_swprintf((*libc.WChar)(memmap.PtrOff(6112660, 1317000)), libc.CWString("%s "), v22)
					v23 = nox_xxx_spellTitle_424930(*(*int32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(int32(0))*108)))
					nox_wcscat((*libc.WChar)(memmap.PtrOff(6112660, 1317000)), v23)
					goto LABEL_93
				}
				v20 = nox_xxx_spellTitle_424930(*(*int32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(int32(0))*108)))
				nox_wcscat((*libc.WChar)(memmap.PtrOff(6112660, 1317000)), v20)
				nox_wcscat((*libc.WChar)(memmap.PtrOff(6112660, 1317000)), libc.CWString(" "))
				v21 = strMan.GetStringInFile("BookOf", "C:\\NoxPost\\src\\client\\Gui\\ToolTip.c")
				goto LABEL_90
			}
			return (*libc.WChar)(memmap.PtrOff(6112660, 1317000))
		}
		if v18&2 != 0 {
			v24 = *(*int32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(int32(0))*108))
			if v24 == 0 {
				*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a1))), 0)) = 226
				*(*uint16)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(&a1))), 1)))) = uint16(nox_xxx_netGetUnitCodeCli_578B00(int32(uintptr(unsafe.Pointer(v1)))))
				*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a1))), unsafe.Sizeof((*libc.WChar)(nil))-1)) = 2
				*(*int32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(int32(0))*108)) = 41
				goto LABEL_93
			}
			if v24 == 41 {
				return (*libc.WChar)(memmap.PtrOff(6112660, 1317000))
			}
			if strMan.Lang() == 3 || strMan.Lang() == 5 {
				v27 = strMan.GetStringInFile("LoreScroll", "C:\\NoxPost\\src\\client\\Gui\\ToolTip.c")
				nox_wcscat((*libc.WChar)(memmap.PtrOff(6112660, 1317000)), v27)
				nox_wcscat((*libc.WChar)(memmap.PtrOff(6112660, 1317000)), libc.CWString(" "))
				v28 = (*libc.WChar)(unsafe.Pointer(uintptr(nox_xxx_guiCreatureGetName_427240(*(*int32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(int32(0))*108))))))
				nox_wcscat((*libc.WChar)(memmap.PtrOff(6112660, 1317000)), v28)
			} else {
				v25 = nox_xxx_guiCreatureGetName_427240(*(*int32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(int32(0))*108)))
				nox_swprintf((*libc.WChar)(memmap.PtrOff(6112660, 1317000)), libc.CWString("%s "), v25)
				v26 = strMan.GetStringInFile("LoreScroll", "C:\\NoxPost\\src\\client\\Gui\\ToolTip.c")
				nox_wcscat((*libc.WChar)(memmap.PtrOff(6112660, 1317000)), v26)
			}
		} else if v18&4 != 0 {
			v29 = *(*int32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(int32(0))*108))
			if v29 == 0 {
				*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a1))), 0)) = 226
				*(*uint16)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(&a1))), 1)))) = uint16(nox_xxx_netGetUnitCodeCli_578B00(int32(uintptr(unsafe.Pointer(v1)))))
				*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a1))), unsafe.Sizeof((*libc.WChar)(nil))-1)) = 4
				*(*int32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(int32(0))*108)) = 6
				goto LABEL_93
			}
			if v29 == 6 {
				return (*libc.WChar)(memmap.PtrOff(6112660, 1317000))
			}
			if strMan.Lang() == 6 {
				v30 = (*libc.WChar)(unsafe.Pointer(uintptr(nox_xxx_abilityGetName_0_425260(*(*int32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(int32(0))*108))))))
				nox_wcscat((*libc.WChar)(memmap.PtrOff(6112660, 1317000)), v30)
				nox_wcscat((*libc.WChar)(memmap.PtrOff(6112660, 1317000)), libc.CWString(" "))
				v21 = strMan.GetStringInFile("BookOf", "C:\\NoxPost\\src\\client\\Gui\\ToolTip.c")
			LABEL_90:
				nox_wcscat((*libc.WChar)(memmap.PtrOff(6112660, 1317000)), v21)
				goto LABEL_93
			}
			v31 = strMan.GetStringInFile("BookOf", "C:\\NoxPost\\src\\client\\Gui\\ToolTip.c")
			nox_swprintf((*libc.WChar)(memmap.PtrOff(6112660, 1317000)), libc.CWString("%s "), v31)
			v32 = (*libc.WChar)(unsafe.Pointer(uintptr(nox_xxx_abilityGetName_0_425260(*(*int32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(int32(0))*108))))))
			nox_wcscat((*libc.WChar)(memmap.PtrOff(6112660, 1317000)), v32)
		} else {
			result = nox_get_thing_pretty_name(*(*int32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(int32(0))*27)))
			if result != nil {
				return result
			}
		}
	LABEL_93:
		if int32(uint8(uintptr(unsafe.Pointer(a1)))) != 243 {
			nox_xxx_netClientSend2_4E53C0(31, unsafe.Pointer(&a1), 4, 0, 1)
			return (*libc.WChar)(memmap.PtrOff(6112660, 1317000))
		}
		return (*libc.WChar)(memmap.PtrOff(6112660, 1317000))
	}
	v3 = nil
	v34 = nil
	v35 = nil
	a1 = nil
	if uint32(v2)&0x11001000 != 0 {
		v4 = nox_xxx_getProjectileClassById_413250(*(*int32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(int32(0))*27)))
	} else {
		v4 = nox_xxx_equipClothFindDefByTT_413270(*(*int32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(int32(0))*27)))
	}
	if v4 != nil {
		v7 = int32(uint32(*(*int32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(int32(0))*28))) & 0x1000000)
		if v7 == 0 || uint32(v7) == 0x1000000 && (uint32(*(*int32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(int32(0))*29)))&0x7800000) == 0 {
			v8 = *(*int32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(int32(0))*108))
			if v8 != 0 {
				v9 = *(**libc.WChar)(unsafe.Pointer(uintptr(v8 + 8)))
				if v9 != nil {
					v3 = v9
				}
			}
			v10 = *(*int32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(int32(0))*109))
			if v10 != 0 {
				v11 = *(**libc.WChar)(unsafe.Pointer(uintptr(v10 + 8)))
				if v11 != nil {
					v34 = v11
				}
			}
		}
		v12 = (*libc.WChar)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*2)))))
		if v7 == 0 || (uint32(*(*int32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(int32(0))*29)))&0x7800000) == 0 {
			v13 = *(*int32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(int32(0))*110))
			if v13 != 0 {
				v14 = *(**libc.WChar)(unsafe.Pointer(uintptr(v13 + 8)))
				if v14 != nil {
					v35 = v14
				}
			}
			v15 = *(*int32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(int32(0))*111))
			if v15 != 0 {
				v16 = *(**libc.WChar)(unsafe.Pointer(uintptr(v15 + 12)))
				if v16 != nil {
					a1 = v16
				}
			}
		}
		switch strMan.Lang() {
		case 2:
			nox_wcscat((*libc.WChar)(memmap.PtrOff(6112660, 1317000)), v12)
			if v35 != nil {
				nox_wcscat((*libc.WChar)(memmap.PtrOff(6112660, 1317000)), libc.CWString(" "))
				nox_wcscat((*libc.WChar)(memmap.PtrOff(6112660, 1317000)), v35)
			}
			v17 = a1
			if a1 != nil {
				nox_wcscat((*libc.WChar)(memmap.PtrOff(6112660, 1317000)), libc.CWString(" "))
				nox_wcscat((*libc.WChar)(memmap.PtrOff(6112660, 1317000)), v17)
			}
			if v34 != nil {
				nox_wcscat((*libc.WChar)(memmap.PtrOff(6112660, 1317000)), v34)
				nox_wcscat((*libc.WChar)(memmap.PtrOff(6112660, 1317000)), libc.CWString(" "))
			}
			if v3 == nil {
				return (*libc.WChar)(memmap.PtrOff(6112660, 1317000))
			}
			nox_wcscat((*libc.WChar)(memmap.PtrOff(6112660, 1317000)), libc.CWString(" "))
			nox_wcscat((*libc.WChar)(memmap.PtrOff(6112660, 1317000)), v3)
			result = (*libc.WChar)(memmap.PtrOff(6112660, 1317000))
		case 3:
			if v3 != nil {
				nox_wcscat((*libc.WChar)(memmap.PtrOff(6112660, 1317000)), v3)
				nox_wcscat((*libc.WChar)(memmap.PtrOff(6112660, 1317000)), libc.CWString(" "))
			}
			nox_wcscat((*libc.WChar)(memmap.PtrOff(6112660, 1317000)), v12)
			if v34 != nil {
				nox_wcscat((*libc.WChar)(memmap.PtrOff(6112660, 1317000)), libc.CWString(" "))
				nox_wcscat((*libc.WChar)(memmap.PtrOff(6112660, 1317000)), v34)
			}
			if v35 != nil {
				nox_wcscat((*libc.WChar)(memmap.PtrOff(6112660, 1317000)), libc.CWString(" "))
				nox_wcscat((*libc.WChar)(memmap.PtrOff(6112660, 1317000)), v35)
			}
			if a1 == nil {
				return (*libc.WChar)(memmap.PtrOff(6112660, 1317000))
			}
			nox_wcscat((*libc.WChar)(memmap.PtrOff(6112660, 1317000)), libc.CWString(" "))
			goto LABEL_67
		case 5:
			nox_wcscat((*libc.WChar)(memmap.PtrOff(6112660, 1317000)), v12)
			if v34 != nil {
				nox_wcscat((*libc.WChar)(memmap.PtrOff(6112660, 1317000)), libc.CWString(" "))
				nox_wcscat((*libc.WChar)(memmap.PtrOff(6112660, 1317000)), v34)
			}
			if v3 != nil {
				nox_wcscat((*libc.WChar)(memmap.PtrOff(6112660, 1317000)), libc.CWString(" "))
				nox_wcscat((*libc.WChar)(memmap.PtrOff(6112660, 1317000)), v3)
			}
			if v35 != nil {
				nox_wcscat((*libc.WChar)(memmap.PtrOff(6112660, 1317000)), libc.CWString(" "))
				nox_wcscat((*libc.WChar)(memmap.PtrOff(6112660, 1317000)), v35)
			}
			if a1 == nil {
				return (*libc.WChar)(memmap.PtrOff(6112660, 1317000))
			}
			nox_wcscat((*libc.WChar)(memmap.PtrOff(6112660, 1317000)), libc.CWString(" "))
			nox_wcscat((*libc.WChar)(memmap.PtrOff(6112660, 1317000)), a1)
			result = (*libc.WChar)(memmap.PtrOff(6112660, 1317000))
		case 6:
			if v3 != nil {
				nox_wcscat((*libc.WChar)(memmap.PtrOff(6112660, 1317000)), v3)
				nox_wcscat((*libc.WChar)(memmap.PtrOff(6112660, 1317000)), libc.CWString(" "))
			}
			if v35 != nil {
				nox_wcscat((*libc.WChar)(memmap.PtrOff(6112660, 1317000)), v35)
				nox_wcscat((*libc.WChar)(memmap.PtrOff(6112660, 1317000)), libc.CWString(" "))
			}
			if a1 != nil {
				nox_wcscat((*libc.WChar)(memmap.PtrOff(6112660, 1317000)), a1)
				nox_wcscat((*libc.WChar)(memmap.PtrOff(6112660, 1317000)), libc.CWString(" "))
			}
			if v34 != nil {
				nox_wcscat((*libc.WChar)(memmap.PtrOff(6112660, 1317000)), libc.CWString(" "))
				nox_wcscat((*libc.WChar)(memmap.PtrOff(6112660, 1317000)), v34)
			}
			nox_wcscat((*libc.WChar)(memmap.PtrOff(6112660, 1317000)), v12)
			result = (*libc.WChar)(memmap.PtrOff(6112660, 1317000))
		default:
			if v3 != nil {
				nox_wcscat((*libc.WChar)(memmap.PtrOff(6112660, 1317000)), v3)
				nox_wcscat((*libc.WChar)(memmap.PtrOff(6112660, 1317000)), libc.CWString(" "))
			}
			if v34 != nil {
				nox_wcscat((*libc.WChar)(memmap.PtrOff(6112660, 1317000)), v34)
				nox_wcscat((*libc.WChar)(memmap.PtrOff(6112660, 1317000)), libc.CWString(" "))
			}
			nox_wcscat((*libc.WChar)(memmap.PtrOff(6112660, 1317000)), v12)
			if v35 != nil {
				nox_wcscat((*libc.WChar)(memmap.PtrOff(6112660, 1317000)), libc.CWString(" "))
				nox_wcscat((*libc.WChar)(memmap.PtrOff(6112660, 1317000)), v35)
			}
			if a1 == nil {
				return (*libc.WChar)(memmap.PtrOff(6112660, 1317000))
			}
			nox_wcscat((*libc.WChar)(memmap.PtrOff(6112660, 1317000)), libc.CWString(" "))
		LABEL_67:
			nox_wcscat((*libc.WChar)(memmap.PtrOff(6112660, 1317000)), a1)
			result = (*libc.WChar)(memmap.PtrOff(6112660, 1317000))
		}
	} else {
		v33 = int32(uintptr(unsafe.Pointer(nox_get_thing_name(*(*int32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(int32(0))*27))))))
		v5 = strMan.GetStringInFile("NoArmsInfo", "C:\\NoxPost\\src\\client\\Gui\\ToolTip.c")
		nox_swprintf((*libc.WChar)(memmap.PtrOff(6112660, 1317000)), v5, v33)
		result = (*libc.WChar)(memmap.PtrOff(6112660, 1317000))
	}
	return result
}
