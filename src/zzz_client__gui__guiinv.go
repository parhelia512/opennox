package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	noxflags "github.com/noxworld-dev/opennox/v1/common/flags"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"unsafe"
)

func nox_xxx_spritePickup_461660(a1 int32, a2 int32, a3 unsafe.Pointer) int32 {
	var (
		v3 *int32
		v4 *libc.WChar
		v6 int32
		v7 int32
		v8 int32
		a4 int2
	)
	if uint32(a2) != dword_5d4594_1062560 && uint32(a2) != *memmap.PtrUint32(6112660, 1049728) && uint32(a2) != *memmap.PtrUint32(6112660, 1049724) && uint32(a2) != dword_5d4594_1062556 && uint32(a2) != dword_5d4594_1062564 {
		v3 = sub_461970(a1, a2)
		if v3 != nil {
			if int32(*(*uint8)(unsafe.Pointer(uintptr(*v3 + 112))))&16 != 0 {
				sub_472310()
			}
		} else {
			if sub_4617C0(a1, a2, a3, &a4) == 0 {
				v4 = strMan.GetStringInFile("InventoryFull", "C:\\NoxPost\\src\\Client\\Gui\\guiinv.c")
				nox_xxx_printCentered_445490(v4)
				return 0
			}
			v6 = a4.field_0
			v7 = a4.field_4
			if nox_client_inventory_grid_1050020[a4.field_4+NOX_INVENTORY_ROW_COUNT*a4.field_0].field_0.flags28&16 != 0 {
				sub_472310()
				v7 = a4.field_4
				v6 = a4.field_0
			}
			if nox_client_inventory_grid_1050020[v7+NOX_INVENTORY_ROW_COUNT*v6].field_0.flags28&0x3001000 != 0 {
				dword_5d4594_1062516 = 0
				if v7 >= 3 {
					dword_5d4594_1062516 = uint32((v7*5 - 10) * 10)
				}
			}
		}
		v8 = int32(uintptr(unsafe.Pointer(nox_get_thing(a2))))
		if v8 != 0 {
			if *(*uint32)(unsafe.Pointer(uintptr(v8 + 32)))&0x1001000 != 0 {
				sub_4673F0(a4.field_0, a4.field_4)
			}
		}
	}
	return 1
}
func sub_4617C0(a1 int32, a2 int32, a3 unsafe.Pointer, a4 *int2) int32 {
	var (
		v4     uint16
		i      uint16
		result int32
		v7     int32
		v8     int32
		v9     *uint8
		v10    *uint32
		v11    *libc.WChar
		v12    int32
		v13    int32
		v14    int32
		v15    *uint32
	)
	v4 = 0
	for {
		for i = 0; int32(i) < 4; i++ {
			if int32(nox_client_inventory_grid_1050020[int32(v4)+NOX_INVENTORY_ROW_COUNT*int32(i)].field_140) == 0 {
				v7 = int32(i)
				v8 = int32(v4)
				v9 = (*uint8)(unsafe.Pointer(&nox_client_inventory_grid_1050020[int32(v4)+NOX_INVENTORY_ROW_COUNT*int32(i)]))
				v10 = &nox_new_drawable_for_thing(a2).field_0
				*(*uint32)(unsafe.Pointer(v9)) = uint32(uintptr(unsafe.Pointer(v10)))
				if v10 != nil {
					*(*uint32)(unsafe.Add(unsafe.Pointer(v10), unsafe.Sizeof(uint32(0))*30)) |= 0x40000000
					v12 = int32(*(*uint32)(unsafe.Pointer(v9)))
					*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v9))), unsafe.Sizeof(uint32(0))*1))) = uint32(a1)
					*(*uint8)(unsafe.Add(unsafe.Pointer(v9), 140)) = 1
					if *(*uint32)(unsafe.Pointer(uintptr(v12 + 112)))&0x13001000 != 0 {
						libc.MemCpy(unsafe.Pointer(uintptr(v12+432)), a3, 20)
						v8 = int32(v4)
					}
					if a4 != nil {
						a4.field_0 = v7
						a4.field_4 = v8
					}
					if sub_461930() != 0 && dword_5d4594_1062480 == 0 {
						if (func() uint32 {
							v13 = int32(*(*uint32)(unsafe.Pointer(v9)))
							v14 = int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(v9)) + 112))))
							return uint32(v14) & 0x1000000
						}()) != 0 && (int32(*(*uint8)(unsafe.Pointer(uintptr(v13 + 116))))&2) == 0 || v14&4096 != 0 {
							v15 = nox_xxx_getProjectileClassById_413250(int32(*(*uint32)(unsafe.Pointer(uintptr(v13 + 108)))))
							if v15 == nil || int32(uint8(int8(1<<int32(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(0x8531A0, 2576) + 2251)))))))&int32(*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(v15))), 62)))) != 0 {
								nox_xxx_clientSetAltWeapon_461550(int32(uintptr(unsafe.Pointer(v9))))
								*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v9))), unsafe.Sizeof(uint32(0))*34))) = 1
							}
						}
					}
					result = 1
				} else {
					v11 = strMan.GetStringInFile("DrawablesExhausted", "C:\\NoxPost\\src\\Client\\Gui\\guiinv.c")
					nox_xxx_printCentered_445490(v11)
					result = 0
				}
				return result
			}
		}
		if int32(func() uint16 {
			p := &v4
			*p++
			return *p
		}()) < 20 {
			continue
		}
		break
	}
	return 0
}
func sub_461A80(a1 int32) {
	var (
		v1 int32
		v2 *byte
		v3 *uint64
		v4 *uint64
		v5 *libc.WChar
	)
	v1 = 0
	v2 = sub_461EF0(a1)
	if v2 != nil {
		if int32(*(*uint8)(unsafe.Pointer(uintptr(**(**uint32)(unsafe.Pointer(v2)) + 112))))&16 != 0 {
			v1 = 1
		}
		sub_461E60((***uint64)(unsafe.Pointer(v2)))
		*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(v2)) + 132))) = 0
		sub_461B50()
		v3 = (*uint64)(unsafe.Pointer(uintptr(sub_461F90(a1))))
		if v3 != nil {
			nox_xxx_spriteDelete_45A4B0((*nox_drawable)(unsafe.Pointer(v3)))
		}
		if v1 != 0 {
			sub_472310()
		}
	} else {
		v4 = *(**uint64)(memmap.PtrOff(6112660, 1049848))
		if *memmap.PtrUint32(6112660, 1049848) != 0 && *(*uint32)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, 1049848) + 128))) == uint32(a1) {
			if int32(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, 1049848) + 112))))&16 != 0 {
				sub_472310()
				v4 = *(**uint64)(memmap.PtrOff(6112660, 1049848))
			}
			nox_xxx_spriteDelete_45A4B0((*nox_drawable)(unsafe.Pointer(v4)))
			*memmap.PtrUint32(6112660, 1049848) = 0
			dword_5d4594_1049856 = 0
			nox_xxx_cursorResetDraggedItem_4776A0()
		} else {
			v5 = strMan.GetStringInFile("DroppedNotFound", "C:\\NoxPost\\src\\Client\\Gui\\guiinv.c")
			nox_xxx_printCentered_445490(v5)
		}
	}
}
func sub_462040(a1 int32) {
	var (
		v1  int32
		v2  *byte
		v3  int32
		v4  unsafe.Pointer
		v5  *uint32
		v6  *uint32
		v7  *libc.WChar
		v8  int32
		v9  int32
		v10 int32
		v11 *byte
		v12 *uint8
		v13 int32
		v14 *uint8
		v16 int16
		v17 *int32
		v18 int16
		v19 int16
		v20 int32
	)
	v1 = a1
	v2 = sub_461EF0(a1)
	if v2 != nil {
		v20 = sub_4622E0(int32(**(**uint32)(unsafe.Pointer(v2))))
		v3 = int32(**(**uint32)(unsafe.Pointer(v2)))
	} else {
		if *memmap.PtrUint32(6112660, 1049848) == 0 || *(*uint32)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, 1049848) + 128))) != uint32(a1) {
			v7 = strMan.GetStringInFile("EquippedNotFound", "C:\\NoxPost\\src\\Client\\Gui\\guiinv.c")
			nox_xxx_printCentered_445490(v7)
			return
		}
		v20 = sub_4622E0(*memmap.PtrInt32(6112660, 1049848))
		v3 = int32(*memmap.PtrUint32(6112660, 1049848))
	}
	v4 = unsafe.Pointer(uintptr(v3 + 432))
	v19 = int16(*(*uint16)(unsafe.Pointer(uintptr(v3 + 294))))
	v18 = int16(*(*uint16)(unsafe.Pointer(uintptr(v3 + 292))))
	if v20 == 9 {
		v7 = strMan.GetStringInFile("TooManyEquipped", "C:\\NoxPost\\src\\Client\\Gui\\guiinv.c")
		nox_xxx_printCentered_445490(v7)
		return
	}
	v5 = &nox_new_drawable_for_thing(int32(*(*uint32)(unsafe.Pointer(uintptr(v3 + 108))))).field_0
	v6 = v5
	if v5 == nil {
		v7 = strMan.GetStringInFile("DrawablesExhausted", "C:\\NoxPost\\src\\Client\\Gui\\guiinv.c")
		nox_xxx_printCentered_445490(v7)
		return
	}
	v8 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(uint32(0))*30)))
	*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*32)) = uint32(v1)
	*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*30)) = uint32(v8) | 0x40000000
	libc.MemCpy(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*108))), v4, 24)
	*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v6))), unsafe.Sizeof(uint16(0))*146))) = uint16(v18)
	*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v6))), unsafe.Sizeof(uint16(0))*147))) = uint16(v19)
	sub_4623E0(v6, v20)
	v9 = 0
	if v2 != nil {
		*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(v2)) + 132))) = 1
		if *(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(v2)) + 136))) != 0 {
			nox_xxx_clientSetAltWeapon_461550(0)
			*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(v2)) + 136))) = 0
		}
	}
	if *(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*28))&0x1000000 != 0 && *(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*29))&12 != 0 {
		v10 = 1
		if dword_5d4594_1062488 != 0 && (func() *byte {
			v11 = sub_461EF0(*(*int32)(unsafe.Pointer(&dword_5d4594_1062488)))
			return v11
		}()) != nil {
			*(*uint32)(unsafe.Pointer(uintptr(**(**uint32)(unsafe.Pointer(v11)) + 128))) = *(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(v11)) + 4)))
			nox_xxx_clientEquip_4623B0(int32(**(**uint32)(unsafe.Pointer(v11))))
		} else {
			v12 = (*uint8)(unsafe.Pointer(&nox_client_inventory_grid_1050020[0]))
			for {
				{
					if v10 == 0 {
						break
					}
					v13 = 0
					v14 = v12
					for int32(*(*uint8)(unsafe.Add(unsafe.Pointer(v14), 140))) == 0 || (int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(v14)) + 115))))&1) == 0 || *(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(v14)) + 116))) != 2 {
						v13++
						v14 = (*uint8)(unsafe.Add(unsafe.Pointer(v14), NOX_INVENTORY_ROW_COUNT*unsafe.Sizeof(nox_inventory_cell_t{})))
						if v13 >= 4 {
							goto LABEL_26
						}
					}
					var v15 int32 = v9 + NOX_INVENTORY_ROW_COUNT*v13
					nox_client_inventory_grid_1050020[v15].field_0.field_32 = nox_client_inventory_grid_1050020[v15].field_4
					nox_xxx_clientEquip_4623B0(int32(uintptr(unsafe.Pointer(nox_client_inventory_grid_1050020[v15].field_0))))
					v10 = 0
				LABEL_26:
					v12 = (*uint8)(unsafe.Add(unsafe.Pointer(v12), unsafe.Sizeof(nox_inventory_cell_t{})))
					v9++
				}
				if int32(uintptr(unsafe.Pointer(v12))) >= int32(uintptr(unsafe.Pointer(&nox_client_inventory_grid_1050020[NOX_INVENTORY_ROW_COUNT-1]))) {
					break
				}
			}
		}
	}
	if v20 == 0 {
		dword_5d4594_1062488 = *(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*32))
	}
	v16 = int16(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v6))), unsafe.Sizeof(uint16(0))*224))))
	if int32(v16) >= 0 {
		sub_470D90(int32(v16), int32(*((*int16)(unsafe.Add(unsafe.Pointer((*int16)(unsafe.Pointer(v6))), unsafe.Sizeof(int16(0))*225)))))
	}
	if dword_5d4594_1062496 != 0 {
		v17 = (*int32)(unsafe.Pointer(sub_461EF0(*(*int32)(unsafe.Pointer(&dword_5d4594_1062496)))))
		if v17 != nil {
			*(*uint32)(unsafe.Pointer(uintptr(*v17 + 136))) = 1
			nox_xxx_clientSetAltWeapon_461550(*v17)
			dword_5d4594_1062496 = 0
		}
	}
}
func sub_462740() int32 {
	var (
		v0 *libc.WChar
		v1 *uint32
	)
	if (*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1062476))))).Flags().IsHidden() != 0 {
		return 0
	}
	(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1062476))))).Hide()
	dword_5d4594_1063116 = 0
	dword_5d4594_1063120 = 0
	v0 = strMan.GetStringInFile("thing.db:IdentifyDescription", "C:\\NoxPost\\src\\Client\\Gui\\guiinv.c")
	nox_wcscpy((*libc.WChar)(memmap.PtrOff(6112660, 1063124)), v0)
	v1 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1062476)))).ChildByID(9156)))
	(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v1)))))).Func94(asWindowEvent(0x400F, 0, 0))
	nox_xxx_wndClearCaptureMain((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1062456))))))
	dword_5d4594_1049864 = 0
	nox_client_setCursorType_477610(0)
	return 1
}
func sub_4627F0(a1 *uint32) int32 {
	var (
		v2     int32
		v3     int32
		result int32
		v5     *libc.WChar
		v6     *uint32
		v7     *libc.WChar
		v8     *libc.WChar
		v9     *uint32
		v10    *uint32
		v11    *libc.WChar
		v12    *libc.WChar
		v13    *libc.WChar
		v14    *libc.WChar
		v15    int32
		v16    int32
		v17    *float32
		v18    float64
		v19    int32
		v20    int32
		v21    *libc.WChar
		v22    *libc.WChar
		v23    int32
		v24    int32
		v25    int32
		v26    *uint32
		v27    int32
		v28    float64
		v29    int32
		v30    float64
		v31    float64
		v32    float64
		v33    int32
		v34    *libc.WChar
		v35    *libc.WChar
		v36    *libc.WChar
		v37    *libc.WChar
		v38    *libc.WChar
		v39    int32
		v40    *uint32
		v41    *libc.WChar
		v42    *libc.WChar
		v43    int32
		v44    int32
		v45    *libc.WChar
		v46    int32
		v47    int32
		v48    *libc.WChar
		v49    int32
		v50    *uint32
		v51    *libc.WChar
		v52    *libc.WChar
		v53    *libc.WChar
		v54    *libc.WChar
		v55    *libc.WChar
		v56    int32
		v57    float64
		v58    float64
		v59    float64
		v60    float64
		v61    float64
		v62    float64
		v63    int32
		v64    int32
		v65    float32
		v66    int32
		v67    int32
		v68    float32
		v69    float32
		v70    float32
		v71    float32
		v72    float32
		v73    int32
		v74    int2
		v75    [256]libc.WChar
		v76    [256]libc.WChar
	)
	v73 = 1
	var mpos nox_point = getMousePos()
	nox_xxx_drawSetTextColor_434390(int32(nox_color_white_2523948))
	v2 = 0
	nox_client_drawSetColor_434460(int32(nox_color_black_2650656))
	nox_client_drawRectFilledOpaque_49CE30(int32(*a1+11), int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*1))+15), 200, 200)
	sub_463370(*(**uint32)(unsafe.Pointer(&dword_5d4594_1062456)), &mpos, (*uint32)(unsafe.Pointer(&v74)))
	if sub_4281F0(&v74, (*int4)(memmap.PtrOff(0x587000, 136352))) != 0 || sub_4281F0(&v74, (*int4)(memmap.PtrOff(0x587000, 136368))) != 0 {
		if sub_4281F0(&v74, (*int4)(memmap.PtrOff(0x587000, 136368))) != 0 && (v74.field_4-13)/50 != 1 {
			nox_client_setCursorType_477610(7)
			goto LABEL_14
		}
	} else {
		if sub_4281F0(&v74, (*int4)(memmap.PtrOff(0x587000, 136336))) != 0 {
			nox_client_setCursorType_477610(0)
			goto LABEL_14
		}
		if sub_478030() == 0 {
			nox_client_setCursorType_477610(7)
			goto LABEL_14
		}
		if sub_479870() == 0 || (func() bool {
			*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v3))), 0)) = uint8(int8(bool2int(sub_479880((*uint32)(unsafe.Pointer(&v74))))))
			return v3 == 0
		}()) {
			nox_client_setCursorType_477610(7)
			goto LABEL_14
		}
	}
	nox_client_setCursorType_477610(6)
LABEL_14:
	result = int32(dword_5d4594_1063116)
	if dword_5d4594_1063116 == 0 {
		if dword_5d4594_1063120 != 0 {
			dword_5d4594_1063120 = 0
			v5 = strMan.GetStringInFile("thing.db:IdentifyDescription", "C:\\NoxPost\\src\\Client\\Gui\\guiinv.c")
			nox_wcscpy((*libc.WChar)(memmap.PtrOff(6112660, 1063124)), v5)
			v6 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1062476)))).ChildByID(9156)))
			result = (*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v6)))))).Func94(asWindowEvent(0x400F, 0, 0))
		}
		return result
	}
	if dword_5d4594_1063120 == dword_5d4594_1063116 {
		return result
	}
	dword_5d4594_1063120 = dword_5d4594_1063116
	v7 = strMan.GetStringInFile("IdentifyItem", "C:\\NoxPost\\src\\Client\\Gui\\guiinv.c")
	nox_swprintf((*libc.WChar)(memmap.PtrOff(6112660, 1063124)), libc.CWString("%s "), v7)
	v8 = nox_xxx_clientAskInfoMb_4BF050((*nox_drawable)(unsafe.Pointer(*(**libc.WChar)(unsafe.Pointer(&dword_5d4594_1063116)))))
	nox_wcscpy(&v75[0], v8)
	if nox_wcscmp(&v75[0], (*libc.WChar)(memmap.PtrOff(6112660, 1063652))) == 0 {
		dword_5d4594_1063120 = 0
	}
	nox_wcscat((*libc.WChar)(memmap.PtrOff(6112660, 1063124)), &v75[0])
	v9 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1062476)))).ChildByID(9151)))
	sub_46AEE0(int32(uintptr(unsafe.Pointer(v9))), int32(uintptr(memmap.PtrOff(6112660, 1063124))))
	v10 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1062476)))).ChildByID(9156)))
	(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v10)))))).Func94(asWindowEvent(0x400F, 0, 0))
	if noxflags.HasGame(noxflags.GameModeCoop) {
		if int32(*(*uint16)(unsafe.Pointer(uintptr(dword_5d4594_1063116 + 294)))) != 0 {
			sub_4633B0(*(*int32)(unsafe.Pointer(&dword_5d4594_1063116)), &v71, &v68)
			v63 = int32(v68)
			v56 = int32(v71)
			v11 = strMan.GetStringInFile("IdentifyDurability", "C:\\NoxPost\\src\\Client\\Gui\\guiinv.c")
			nox_swprintf(&v75[0], v11, v56, v63)
		} else {
			v12 = strMan.GetStringInFile("IdentifyDurabilityIndestructable", "C:\\NoxPost\\src\\Client\\Gui\\guiinv.c")
			nox_wcscpy(&v75[0], v12)
		}
	} else {
		switch sub_57B190(*(*uint16)(unsafe.Pointer(uintptr(dword_5d4594_1063116 + 292))), *(*uint16)(unsafe.Pointer(uintptr(dword_5d4594_1063116 + 294)))) {
		case 0:
			v13 = strMan.GetStringInFile("IdentifyDurabilityNoDamage", "C:\\NoxPost\\src\\Client\\Gui\\guiinv.c")
			goto LABEL_30
		case 1:
			v52 = strMan.GetStringInFile("IdentifyDurabilitySlight", "C:\\NoxPost\\src\\Client\\Gui\\guiinv.c")
			nox_swprintf(&v75[0], v52)
		case 2:
			v53 = strMan.GetStringInFile("IdentifyDurabilityModerate", "C:\\NoxPost\\src\\Client\\Gui\\guiinv.c")
			nox_swprintf(&v75[0], v53)
		case 3:
			v13 = strMan.GetStringInFile("IdentifyDurabilitySevere", "C:\\NoxPost\\src\\Client\\Gui\\guiinv.c")
		LABEL_30:
			nox_swprintf(&v75[0], v13)
		case 4:
			v51 = strMan.GetStringInFile("IdentifyDurabilityIndestructable", "C:\\NoxPost\\src\\Client\\Gui\\guiinv.c")
			nox_swprintf(&v75[0], v51)
		default:
		}
	}
	(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v10)))))).Func94(asWindowEvent(0x400D, int32(uintptr(unsafe.Pointer(&v75[0]))), -1))
	(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v10)))))).Func94(asWindowEvent(0x400D, int32(uintptr(memmap.PtrOff(6112660, 1063656))), -1))
	v64 = int32(*(*uint8)(unsafe.Pointer(uintptr(dword_5d4594_1063116 + 298))))
	v14 = strMan.GetStringInFile("IdentifyWeight", "C:\\NoxPost\\src\\Client\\Gui\\guiinv.c")
	nox_swprintf(&v75[0], v14, v64)
	(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v10)))))).Func94(asWindowEvent(0x400D, int32(uintptr(unsafe.Pointer(&v75[0]))), -1))
	(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v10)))))).Func94(asWindowEvent(0x400D, int32(uintptr(memmap.PtrOff(6112660, 1063660))), -1))
	v15 = int32(dword_5d4594_1063116)
	v16 = int32(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1063116 + 112))))
	if (uint32(v16) & 0x2000000) == 0 {
		if (uint32(v16) & 0x1001000) == 0 {
			goto LABEL_72
		}
		v23 = int32(*memmap.PtrUint32(0x8531A0, 2576))
		*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v68))), unsafe.Sizeof(uint32(0))*0)) = dword_5d4594_1063116 + 432
		v69 = 1.0
		v70 = *mem_getFloatPtr(0x8531A0, 2576)
		if *memmap.PtrUint32(6112660, 1063644) == 0 {
			*memmap.PtrUint32(6112660, 1063644) = uint32(nox_xxx_getTTByNameSpriteMB_44CFC0("ArcherArrow"))
			v24 = nox_xxx_getTTByNameSpriteMB_44CFC0("ArcherBolt")
			v23 = int32(*memmap.PtrUint32(0x8531A0, 2576))
			*memmap.PtrUint32(6112660, 1063648) = uint32(v24)
			v15 = int32(dword_5d4594_1063116)
		}
		if v23 == 0 || (int32(*(*uint8)(unsafe.Pointer(uintptr(v15 + 116))))&2) == 0 {
			goto LABEL_50
		}
		v25 = int32(*(*uint32)(unsafe.Pointer(uintptr(v23 + 4))))
		if v25&4 != 0 {
			v26 = nox_xxx_getProjectileClassById_413250(*memmap.PtrInt32(6112660, 1063644))
			v2 = 4
		} else {
			if (v25 & 8) == 0 {
				goto LABEL_50
			}
			v26 = nox_xxx_getProjectileClassById_413250(*memmap.PtrInt32(6112660, 1063648))
			v2 = 8
		}
		if v26 != nil {
			goto LABEL_51
		}
		v15 = int32(dword_5d4594_1063116)
	LABEL_50:
		v26 = nox_xxx_getProjectileClassById_413250(int32(*(*uint32)(unsafe.Pointer(uintptr(v15 + 108)))))
	LABEL_51:
		v71 = float32(sub_4626C0(*(*int32)(unsafe.Pointer(&dword_5d4594_1063116))))
		v72 = float32(sub_462700(*(*int32)(unsafe.Pointer(&dword_5d4594_1063116))))
		v27 = int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v68))), unsafe.Sizeof(uint32(0))*0))))))
		if *(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v68))), unsafe.Sizeof(uint32(0))*0))))) != 0 && cgoFuncAddr(cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(v27 + 40))), (*func(int32, int32, int32, int32, *float32) *float32)(nil))) == cgoFuncAddr(nox_xxx_effectDamageMultiplier_4E04C0) {
			v69 = *(*float32)(unsafe.Pointer(uintptr(v27 + 44)))
		}
		v28 = nox_xxx_calcBoltDamage_4EF1E0(int32(*(*uint32)(unsafe.Pointer(uintptr((*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v70))), unsafe.Sizeof(uint32(0))*0))) + 2239)))), int32(uintptr(unsafe.Pointer(v26))))
		v29 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v26), unsafe.Sizeof(uint32(0))*1)))
		v70 = float32(v28*float64(v69) + float64(v71) + float64(v72))
		if uint32(v29) == *memmap.PtrUint32(6112660, 1063648) && noxflags.HasGame(noxflags.GameModeCoop) {
			v30 = nox_xxx_gamedataGetFloat_419D40((*byte)(memmap.PtrOff(0x587000, 137632)))
		} else {
			*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v68))), unsafe.Sizeof(uint32(0))*0)) = uint32(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v26))), unsafe.Sizeof(uint16(0))*36))))
			v30 = float64(*(*int32)(unsafe.Add(unsafe.Pointer((*int32)(unsafe.Pointer(&v68))), unsafe.Sizeof(int32(0))*0)))
		}
		v68 = float32(v30 * float64(v69))
		v31 = float64(v70 - v68 - v72 - v71)
		v69 = float32(v31)
		if v31 < 0.0 {
			v32 = float64(v70 - v69)
			v69 = 0.0
			v70 = float32(v32)
		}
		v33 = int32(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1063116 + 116))))
		if v33&12 != 0 {
			v34 = strMan.GetStringInFile("WeaponDamageLabelNA", "C:\\NoxPost\\src\\Client\\Gui\\guiinv.c")
			nox_swprintf(&v75[0], v34)
		} else if (v33&2) == 0 || v2 != 0 {
			v58 = float64(v70)
			v55 = strMan.GetStringInFile("WeaponDamageLabel", "C:\\NoxPost\\src\\Client\\Gui\\guiinv.c")
			nox_swprintf(&v75[0], v55, v58)
		} else {
			v57 = float64(v70)
			v54 = strMan.GetStringInFile("WeaponDamageLabelUnknownPlus", "C:\\NoxPost\\src\\Client\\Gui\\guiinv.c")
			nox_swprintf(&v75[0], v54, v57)
		}
		(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v10)))))).Func94(asWindowEvent(0x400D, int32(uintptr(unsafe.Pointer(&v75[0]))), -1))
		nox_wcscpy(&v75[0], libc.CWString("  "))
		v59 = float64(v68)
		v35 = strMan.GetStringInFile("BaseDamageLabel", "C:\\NoxPost\\src\\Client\\Gui\\guiinv.c")
		nox_swprintf(&v76[0], v35, v59)
		nox_wcscat(&v75[0], &v76[0])
		(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v10)))))).Func94(asWindowEvent(0x400D, int32(uintptr(unsafe.Pointer(&v75[0]))), -1))
		nox_wcscpy(&v75[0], libc.CWString("  "))
		v60 = float64(v69)
		v36 = strMan.GetStringInFile("StrengthDamageLabel", "C:\\NoxPost\\src\\Client\\Gui\\guiinv.c")
		nox_swprintf(&v76[0], v36, v60)
		nox_wcscat(&v75[0], &v76[0])
		(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v10)))))).Func94(asWindowEvent(0x400D, int32(uintptr(unsafe.Pointer(&v75[0]))), -1))
		if float64(v72) > 0.0 {
			nox_wcscpy(&v75[0], libc.CWString("  "))
			v61 = float64(v72)
			v37 = strMan.GetStringInFile("FireDamageLabel", "C:\\NoxPost\\src\\Client\\Gui\\guiinv.c")
			nox_swprintf(&v76[0], v37, v61)
			nox_wcscat(&v75[0], &v76[0])
			(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v10)))))).Func94(asWindowEvent(0x400D, int32(uintptr(unsafe.Pointer(&v75[0]))), -1))
		}
		if float64(v71) > 0.0 {
			nox_wcscpy(&v75[0], libc.CWString("  "))
			v62 = float64(v71)
			v38 = strMan.GetStringInFile("ElectricalDamageLabel", "C:\\NoxPost\\src\\Client\\Gui\\guiinv.c")
			nox_swprintf(&v76[0], v38, v62)
			nox_wcscat(&v75[0], &v76[0])
			(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v10)))))).Func94(asWindowEvent(0x400D, int32(uintptr(unsafe.Pointer(&v75[0]))), -1))
		}
		(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v10)))))).Func94(asWindowEvent(0x400D, int32(uintptr(memmap.PtrOff(6112660, 1063668))), -1))
		goto LABEL_71
	}
	v17 = (*float32)(unsafe.Pointer(nox_xxx_equipClothFindDefByTT_413270(int32(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1063116 + 108)))))))
	v18 = 1.0
	v19 = int32(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1063116 + 432))))
	if v19 != 0 && cgoFuncAddr(cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(v19 + 76))), (*func(int32, int32, int32, int32, int32, *float32) *float32)(nil))) == cgoFuncAddr(sub_4E0370) {
		v18 = float64(*(*float32)(unsafe.Pointer(uintptr(v19 + 80))))
	}
	v65 = float32(v18*float64(*(*float32)(unsafe.Add(unsafe.Pointer(v17), unsafe.Sizeof(float32(0))*16)))*1000.0 + 0.5)
	v20 = int(v65)
	if int32(*(*uint8)(unsafe.Pointer(uintptr(dword_5d4594_1063116 + 116))))&2 != 0 {
		v21 = strMan.GetStringInFile("ArmorValueLabelNA", "C:\\NoxPost\\src\\Client\\Gui\\guiinv.c")
		nox_swprintf(&v75[0], v21)
	} else {
		v66 = v20
		v22 = strMan.GetStringInFile("ArmorValueLabel", "C:\\NoxPost\\src\\Client\\Gui\\guiinv.c")
		nox_swprintf(&v75[0], v22, v66)
	}
	(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v10)))))).Func94(asWindowEvent(0x400D, int32(uintptr(unsafe.Pointer(&v75[0]))), -1))
	(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v10)))))).Func94(asWindowEvent(0x400D, int32(uintptr(memmap.PtrOff(6112660, 1063664))), -1))
LABEL_71:
	v15 = int32(dword_5d4594_1063116)
LABEL_72:
	v39 = int32(*(*uint32)(unsafe.Pointer(uintptr(v15 + 112))))
	if uint32(v39)&0x13001000 != 0 {
		if uint32(v39)&0x11001000 != 0 {
			v40 = nox_xxx_getProjectileClassById_413250(int32(*(*uint32)(unsafe.Pointer(uintptr(v15 + 108)))))
		} else {
			v40 = nox_xxx_equipClothFindDefByTT_413270(int32(*(*uint32)(unsafe.Pointer(uintptr(v15 + 108)))))
		}
		if v40 != nil {
			v15 = int32(dword_5d4594_1063116)
			v43 = int32(dword_5d4594_1063116 + 432)
			if *(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1063116 + 112)))&0x10000000 != 0 {
				goto LABEL_91
			}
			v44 = int32(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1063116 + 440))))
			if v44 != 0 && *(*uint32)(unsafe.Pointer(uintptr(v44 + 16))) != 0 {
				v45 = strMan.GetStringInFile("IdentifySpecialAttributes", "C:\\NoxPost\\src\\Client\\Gui\\guiinv.c")
				nox_swprintf(&v75[0], v45)
				(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v10)))))).Func94(asWindowEvent(0x400D, int32(uintptr(unsafe.Pointer(&v75[0]))), -1))
				v46 = 0
				nox_wcscpy(&v75[0], libc.CWString("  "))
				nox_swprintf(&v76[0], *(**libc.WChar)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v43 + 8))) + 16))))
				nox_wcscat(&v75[0], &v76[0])
				(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v10)))))).Func94(asWindowEvent(0x400D, int32(uintptr(unsafe.Pointer(&v75[0]))), -1))
				v15 = int32(dword_5d4594_1063116)
			} else {
				v46 = v73
			}
			v47 = int32(*(*uint32)(unsafe.Pointer(uintptr(v43 + 12))))
			if v47 != 0 && *(*uint32)(unsafe.Pointer(uintptr(v47 + 16))) != 0 {
				if v46 == 1 {
					v48 = strMan.GetStringInFile("IdentifySpecialAttributes", "C:\\NoxPost\\src\\Client\\Gui\\guiinv.c")
					nox_swprintf(&v75[0], v48)
					(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v10)))))).Func94(asWindowEvent(0x400D, int32(uintptr(unsafe.Pointer(&v75[0]))), -1))
					v73 = 0
					v46 = 0
				}
				nox_wcscpy(&v75[0], libc.CWString("  "))
				nox_swprintf(&v76[0], *(**libc.WChar)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v43 + 12))) + 16))))
				nox_wcscat(&v75[0], &v76[0])
				(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v10)))))).Func94(asWindowEvent(0x400D, int32(uintptr(unsafe.Pointer(&v75[0]))), -1))
				v15 = int32(dword_5d4594_1063116)
			}
			if v46 != 0 {
				goto LABEL_91
			}
			(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v10)))))).Func94(asWindowEvent(0x400D, int32(uintptr(memmap.PtrOff(6112660, 1063672))), -1))
		} else {
			v41 = strMan.GetStringInFile("IdentifySpecialAttributes", "C:\\NoxPost\\src\\Client\\Gui\\guiinv.c")
			nox_swprintf(&v75[0], v41)
			(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v10)))))).Func94(asWindowEvent(0x400D, int32(uintptr(unsafe.Pointer(&v75[0]))), -1))
			v42 = strMan.GetStringInFile("IdentifyUnknown", "C:\\NoxPost\\src\\Client\\Gui\\guiinv.c")
			nox_swprintf(&v75[0], v42)
			(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v10)))))).Func94(asWindowEvent(0x400D, int32(uintptr(unsafe.Pointer(&v75[0]))), -1))
		}
		v15 = int32(dword_5d4594_1063116)
	}
LABEL_91:
	v49 = int32(uintptr(unsafe.Pointer(nox_get_thing_desc(int32(*(*uint32)(unsafe.Pointer(uintptr(v15 + 108))))))))
	if v49 != 0 {
		(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v10)))))).Func94(asWindowEvent(0x400D, v49, -1))
	}
	v67 = nox_get_thing_pretty_image(int32(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1063116 + 108)))))
	v50 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1062476)))).ChildByID(9155)))
	return nox_xxx_wndSetIcon_46AE60(int32(uintptr(unsafe.Pointer(v50))), v67)
}
func nox_client_makePlayerStatsDlg_463880(a1 *int32) {
	var (
		v77 [256]libc.WChar
		v1  int32   = nox_xxx_guiFontHeightMB_43F320(nil)
		v2  int32   = nox_xxx_guiFontHeightMB_43F320(unsafe.Pointer(uintptr(dword_5d4594_1063636)))
		v68 int32   = v1 - v2
		v51 float32 = float32(float64(v1-v2)*0.5 + 0.5)
		v3  int32   = int(v51)
		v73 int32   = v3
		v72 int32   = int32(nox_color_white_2523948)
		v6  int32   = *(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*0))
		v7  int32   = *(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*1))
		v4  int32   = int32(*memmap.PtrUint32(0x8531A0, 2576))
	)
	if v4 == 0 {
		return
	}
	sub_57B350()
	var v70 *float32 = nox_xxx_plrGetMaxVarsPtr_57B360(int32(*(*uint8)(unsafe.Pointer(uintptr(v4 + 2251)))))
	var v71 *float32 = nox_xxx_plrGetMaxVarsPtr_57B360(0)
	var v8 int32 = v6 + 11
	var v9 int32 = v7 + 15
	nox_xxx_drawSetTextColor_434390(v72)
	nox_client_drawSetColor_434460(int32(nox_color_black_2650656))
	nox_client_drawRectFilledOpaque_49CE30(v8, v9, 200, 200)
	var v10 int32 = v8 + 2
	var v11 int32 = v9 + v1*2 + 3
	var v52 int32 = int32(*(*byte)(unsafe.Pointer(uintptr(v4 + 3684))))
	var v12 *libc.WChar = strMan.GetStringInFile("StatsLevel", "C:\\NoxPost\\src\\Client\\Gui\\guiinv.c")
	nox_swprintf(&v77[0], v12, v52)
	nox_xxx_drawStringWrap_43FAF0(nil, &v77[0], v10, v11, 200, 0)
	var v13 int32 = v11 + v1 + 1
	if noxflags.HasGame(noxflags.GameModeCoop) {
		var (
			v53 int32       = int32(int64(nox_xxx_gamedataGetFloatTable_419D70(CString("XPTable"), int32(*(*byte)(unsafe.Pointer(uintptr(v4 + 3684)))+1))))
			v41 int32       = int32(*memmap.PtrUint32(6112660, 1062544))
			v14 *libc.WChar = strMan.GetStringInFile("StatsEXP", "C:\\NoxPost\\src\\Client\\Gui\\guiinv.c")
		)
		nox_swprintf(&v77[0], v14, v41, v53)
		nox_xxx_drawStringWrap_43FAF0(nil, &v77[0], v10, v13, 200, 0)
	}
	var v15 int32 = v1*2 + 2 + v13
	var v16 *libc.WChar = strMan.GetStringInFile("StatsHealth", "C:\\NoxPost\\src\\Client\\Gui\\guiinv.c")
	nox_xxx_drawStringWrap_43FAF0(nil, v16, v10, v15, 200, 0)
	nox_client_drawSetColor_434460(int32(nox_color_violet_2598268))
	nox_client_drawRectFilledOpaque_49CE30(v10+60, v15, 90, v1)
	var v54 float32 = float32(float64(int32(*(*uint32)(unsafe.Pointer(uintptr(v4 + 2247)))*90)) / float64(*v70))
	var v67 int32 = int(v54)
	nox_client_drawSetColor_434460(int32(nox_color_red_2589776))
	nox_client_drawRectFilledOpaque_49CE30(v10+60, v15, v67, v1)
	v68 = sub_470CC0() * 90
	var v55 float32 = float32(float64(v68) / float64(*v70))
	v67 = int(v55)
	nox_client_drawSetColor_434460(*memmap.PtrInt32(0x85B3FC, 940))
	nox_client_drawRectFilledOpaque_49CE30(v10+60, v15, v67, v1)
	var v56 int32 = int(*v70)
	var v42 int32 = int32(*(*uint32)(unsafe.Pointer(uintptr(v4 + 2247))))
	var v17 *libc.WChar = strMan.GetStringInFile("MinMaxFormat", "C:\\NoxPost\\src\\Client\\Gui\\guiinv.c")
	nox_swprintf(&v77[0], v17, v42, v56)
	nox_xxx_drawGetStringSize_43F840(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1063636)))), &v77[0], &v67, nil, 0)
	var v69 float32 = 0
	*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v69))), unsafe.Sizeof(uint32(0))*0)) = uint32(v15 + v73)
	nox_xxx_drawStringWrap_43FAF0(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1063636)))), &v77[0], v10-v67+193, v15+v73, 200, 0)
	var v18 int32 = sub_470CC0()
	nox_swprintf(&v77[0], libc.CWString("%d"), v18)
	nox_xxx_drawStringWrap_43FAF0(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1063636)))), &v77[0], v10+45, *(*int32)(unsafe.Add(unsafe.Pointer((*int32)(unsafe.Pointer(&v69))), unsafe.Sizeof(int32(0))*0)), 200, 0)
	var v19 int32 = v15 + v1 + 1
	if int32(*(*uint8)(unsafe.Pointer(uintptr(v4 + 2251)))) != 0 {
		nox_client_drawSetColor_434460(*memmap.PtrInt32(0x85B3FC, 944))
		nox_client_drawRectFilledOpaque_49CE30(v10+60, v19, 90, v1)
		v68 = int32(*(*uint32)(unsafe.Pointer(uintptr(v4 + 2243))) * 90)
		var v57 float32 = float32(float64(v68) / float64(*(*float32)(unsafe.Add(unsafe.Pointer(v70), unsafe.Sizeof(float32(0))*1))))
		v67 = int(v57)
		var v20 *libc.WChar = strMan.GetStringInFile("StatsMana", "C:\\NoxPost\\src\\Client\\Gui\\guiinv.c")
		nox_xxx_drawStringWrap_43FAF0(nil, v20, v10, v19, 200, 0)
		nox_client_drawSetColor_434460(int32(nox_color_blue_2650684))
		nox_client_drawRectFilledOpaque_49CE30(v10+60, v19, v67, v1)
		v68 = nox_xxx_cliGetMana_470DD0() * 90
		var v58 float32 = float32(float64(v68) / float64(*(*float32)(unsafe.Add(unsafe.Pointer(v70), unsafe.Sizeof(float32(0))*1))))
		v67 = int(v58)
		nox_client_drawSetColor_434460(int32(nox_color_cyan_2649820))
		nox_client_drawRectFilledOpaque_49CE30(v10+60, v19, v67, v1)
		var v59 int32 = int(*(*float32)(unsafe.Add(unsafe.Pointer(v70), unsafe.Sizeof(float32(0))*1)))
		var v43 int32 = int32(*(*uint32)(unsafe.Pointer(uintptr(v4 + 2243))))
		var v21 *libc.WChar = strMan.GetStringInFile("MinMaxFormat", "C:\\NoxPost\\src\\Client\\Gui\\guiinv.c")
		nox_swprintf(&v77[0], v21, v43, v59)
		nox_xxx_drawGetStringSize_43F840(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1063636)))), &v77[0], &v67, nil, 0)
		nox_xxx_drawStringWrap_43FAF0(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1063636)))), &v77[0], v10-v67+193, v19+v73, 200, 0)
		var v22 int32 = nox_xxx_cliGetMana_470DD0()
		nox_swprintf(&v77[0], libc.CWString("%d"), v22)
		nox_xxx_drawStringWrap_43FAF0(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1063636)))), &v77[0], v10+45, v19+v73, 200, 0)
		v19 += v1 + 1
	}
	nox_client_drawSetColor_434460(*memmap.PtrInt32(0x85B3FC, 956))
	nox_client_drawRectFilledOpaque_49CE30(v10+60, v19, 90, v1)
	v68 = int32(*(*uint32)(unsafe.Pointer(uintptr(v4 + 2239))) * 90)
	var v60 float32 = float32(float64(v68) / float64(*(*float32)(unsafe.Add(unsafe.Pointer(v70), unsafe.Sizeof(float32(0))*3))))
	v67 = int(v60)
	var v23 *libc.WChar = strMan.GetStringInFile("StatsStrength", "C:\\NoxPost\\src\\Client\\Gui\\guiinv.c")
	nox_xxx_drawStringWrap_43FAF0(nil, v23, v10, v19, 200, 0)
	nox_client_drawSetColor_434460(*memmap.PtrInt32(6112660, 2597996))
	nox_client_drawRectFilledOpaque_49CE30(v10+60, v19, v67, v1)
	var v61 int32 = int(*(*float32)(unsafe.Add(unsafe.Pointer(v70), unsafe.Sizeof(float32(0))*3)))
	var v44 int32 = int32(*(*uint32)(unsafe.Pointer(uintptr(v4 + 2239))))
	var v24 *libc.WChar = strMan.GetStringInFile("MinMaxFormat", "C:\\NoxPost\\src\\Client\\Gui\\guiinv.c")
	nox_swprintf(&v77[0], v24, v44, v61)
	nox_xxx_drawGetStringSize_43F840(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1063636)))), &v77[0], &v67, nil, 0)
	nox_xxx_drawStringWrap_43FAF0(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1063636)))), &v77[0], v10-v67+193, v19+v73, 200, 0)
	nox_swprintf(&v77[0], libc.CWString("%d"), *(*uint32)(unsafe.Pointer(uintptr(v4 + 2239))))
	nox_xxx_drawStringWrap_43FAF0(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1063636)))), &v77[0], v10+45, v19+v73, 200, 0)
	var v25 int32 = v19 + v1 + 1
	nox_client_drawSetColor_434460(int32(nox_color_orange_2614256))
	nox_client_drawRectFilledOpaque_49CE30(v10+60, v25, 90, v1)
	v68 = int32(*(*uint32)(unsafe.Pointer(uintptr(v4 + 2235))) * 90)
	var v62 float32 = float32(float64(v68)/float64(*(*float32)(unsafe.Add(unsafe.Pointer(v70), unsafe.Sizeof(float32(0))*2))) + 0.5)
	v67 = int(v62)
	var v26 *libc.WChar = strMan.GetStringInFile("StatsSpeed", "C:\\NoxPost\\src\\Client\\Gui\\guiinv.c")
	nox_xxx_drawStringWrap_43FAF0(nil, v26, v10, v25, 200, 0)
	nox_client_drawSetColor_434460(int32(nox_color_yellow_2589772))
	nox_client_drawRectFilledOpaque_49CE30(v10+60, v25, v67, v1)
	nox_xxx_drawSetTextColor_434390(int32(nox_color_white_2523948))
	v69 = float32(float64(*mem_getFloatPtr(6112660, 1063100)) / (float64(*(*float32)(unsafe.Add(unsafe.Pointer(v71), unsafe.Sizeof(float32(0))*2))) * 1e-06))
	if int32(*memmap.PtrUint8(6112660, 1062541))&2 != 0 {
		v69 = float32((float64(v67)+float64(v69))*1.25 - float64(v67))
	}
	if int32(*memmap.PtrUint8(6112660, 1062540))&16 != 0 {
		v69 = float32((float64(v67)+float64(v69))*0.5 - float64(v67))
	}
	if float64(v69) >= 0.0 {
		if float64(v69) > 0.0 {
			*(*float32)(unsafe.Pointer(&v68)) = COERCE_FLOAT(uint32(int(v69)))
			if v67+v68 > 90 {
				v68 = 90 - v67
			}
			nox_client_drawSetColor_434460(int32(nox_color_yellow_2589772))
			nox_client_drawRectFilledOpaque_49CE30(v67+v10+60, v25, v68, v1)
			nox_xxx_drawSetTextColor_434390(int32(nox_color_blue_2650684))
		}
	} else {
		nox_client_drawSetColor_434460(*memmap.PtrInt32(0x85B3FC, 944))
		var v45 float32 = -v69
		var v46 int32 = int(v45)
		var v27 int32 = int(v69)
		nox_client_drawRectFilledOpaque_49CE30(v67+v27+v10+60, v25, v46, v1)
	}
	*(*float32)(unsafe.Pointer(&v68)) = float32(float64(v69) * 100.0 * 0.011111111)
	var v63 float32 = float32(float64(*(*float32)(unsafe.Add(unsafe.Pointer(v70), unsafe.Sizeof(float32(0))*2))) * 100.0 / float64(*(*float32)(unsafe.Add(unsafe.Pointer(v71), unsafe.Sizeof(float32(0))*2))))
	var v64 int32 = int(v63)
	var v47 float32 = float32(float64(*(*int32)(unsafe.Pointer(uintptr(v4 + 2235))))*100.0/float64(*(*float32)(unsafe.Add(unsafe.Pointer(v71), unsafe.Sizeof(float32(0))*2))) + float64(*(*float32)(unsafe.Pointer(&v68))) + 0.5)
	var v48 int32 = int(v47)
	var v28 *libc.WChar = strMan.GetStringInFile("MinMaxFormat", "C:\\NoxPost\\src\\Client\\Gui\\guiinv.c")
	nox_swprintf(&v77[0], v28, v48, v64)
	var v76 int32 = 0
	nox_xxx_drawGetStringSize_43F840(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1063636)))), &v77[0], &v76, nil, 0)
	*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v69))), unsafe.Sizeof(uint32(0))*0)) = uint32(v25 + v73)
	nox_xxx_drawStringWrap_43FAF0(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1063636)))), &v77[0], v10-v76+193, v25+v73, 200, 0)
	nox_xxx_drawSetTextColor_434390(int32(nox_color_white_2523948))
	var v65 float32 = float32(float64(*(*int32)(unsafe.Pointer(uintptr(v4 + 2235))))*100.0/float64(*(*float32)(unsafe.Add(unsafe.Pointer(v71), unsafe.Sizeof(float32(0))*2))) + float64(*(*float32)(unsafe.Pointer(&v68))) + 0.5)
	var v29 int32 = int(v65)
	nox_swprintf(&v77[0], libc.CWString("%d"), v29)
	nox_xxx_drawStringWrap_43FAF0(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1063636)))), &v77[0], v10+45, *(*int32)(unsafe.Add(unsafe.Pointer((*int32)(unsafe.Pointer(&v69))), unsafe.Sizeof(int32(0))*0)), 200, 0)
	nox_xxx_drawSetTextColor_434390(int32(nox_color_white_2523948))
	var v30 int32 = v1*2 + 2 + v25
	if strMan.Lang() == 6 || strMan.Lang() == 8 {
		v10 += 39
	}
	nox_xxx_drawSetTextColor_434390(v72)
	var v31 *libc.WChar = strMan.GetStringInFile("StatsArmorLabel", "C:\\NoxPost\\src\\Client\\Gui\\guiinv.c")
	nox_wcscpy(&v77[0], v31)
	var v75 int32 = 0
	nox_xxx_drawGetStringSize_43F840(nil, &v77[0], &v75, nil, 0)
	nox_xxx_drawStringWrap_43FAF0(nil, &v77[0], v10, v30, 0, 0)
	var v49 float32 = float32(float64(*mem_getFloatPtr(6112660, 0x103694))*1000.0 + 0.5)
	var v50 int32 = int(v49)
	var v32 *libc.WChar = strMan.GetStringInFile("MinMaxFormat", "C:\\NoxPost\\src\\Client\\Gui\\guiinv.c")
	nox_swprintf(&v77[0], v32, v50, 1000)
	nox_xxx_drawStringWrap_43FAF0(nil, &v77[0], v75+v10+5, v30, 0, 0)
	var v74 int32 = v30 + v1 + 1
	var itemsWeight int32 = 0
	for i := int32(0); i < NOX_INVENTORY_ROW_COUNT; i++ {
		var (
			v71a *nox_inventory_cell_t = &nox_client_inventory_grid_1050020[i]
			v33  *uint8                = &v71a.field_140
		)
		for j := int32(0); j < NOX_INVENTORY_COL_COUNT; j++ {
			if int32(*v33) != 0 {
				itemsWeight += int32(*v33) * int32(*(*uint8)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v33))), -int(unsafe.Sizeof(uint32(0))*35)))) + 298))))
			}
			v33 = (*uint8)(unsafe.Add(unsafe.Pointer(v33), NOX_INVENTORY_ROW_COUNT*unsafe.Sizeof(nox_inventory_cell_t{})))
		}
	}
	nox_xxx_drawSetTextColor_434390(v72)
	var v35 *libc.WChar = strMan.GetStringInFile("DollWeight", "C:\\NoxPost\\src\\Client\\Gui\\guiinv.c")
	nox_xxx_drawGetStringSize_43F840(nil, v35, &v67, nil, 0)
	var v36 int32 = v74
	var v40 int32 = v74
	var v39 int32 = v10 + v75 - v67
	var v37 *libc.WChar = strMan.GetStringInFile("DollWeight", "C:\\NoxPost\\src\\Client\\Gui\\guiinv.c")
	nox_xxx_drawStringWrap_43FAF0(nil, v37, v39, v40, 0, 0)
	if itemsWeight > int32(*(*uint16)(unsafe.Pointer(uintptr(v4 + 3652)))) {
		v72 = int32(*memmap.PtrUint32(0x85B3FC, 940))
	}
	nox_xxx_drawSetTextColor_434390(v72)
	var v66 int32 = int32(*(*uint16)(unsafe.Pointer(uintptr(v4 + 3652))))
	var v38 *libc.WChar = strMan.GetStringInFile("MinMaxFormat", "C:\\NoxPost\\src\\Client\\Gui\\guiinv.c")
	nox_swprintf(&v77[0], v38, itemsWeight, v66)
	nox_xxx_drawStringWrap_43FAF0(nil, &v77[0], v75+v10+5, v36, 0, 0)
}
func sub_4649B0(a1 int32, a2 int32, a3 int32) int32 {
	var (
		result int32
		v4     *uint8
		v5     uint8
		v6     *uint32
		v7     *libc.WChar
		v8     *uint8
		v9     int32
		v10    int32
	)
	result = sub_464B40(a2, a3)
	if result == 0 {
		return 0
	}
	v4 = (*uint8)(unsafe.Pointer(&nox_client_inventory_grid_1050020[a3+NOX_INVENTORY_ROW_COUNT*a2]))
	v5 = nox_client_inventory_grid_1050020[a3+NOX_INVENTORY_ROW_COUNT*a2].field_140
	if int32(v5) != 0 {
		if *(*uint32)(unsafe.Pointer(uintptr(a1 + 112)))&0x4000000 != 0 {
			return 0
		}
		if *(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(v4)) + 108))) != *(*uint32)(unsafe.Pointer(uintptr(a1 + 108))) {
			return 0
		}
	}
	if int32(v5) >= 32 {
		return 0
	}
	if int32(v5) == 0 {
		v6 = &nox_new_drawable_for_thing(int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 108))))).field_0
		*(*uint32)(unsafe.Pointer(v4)) = uint32(uintptr(unsafe.Pointer(v6)))
		if v6 == nil {
			v7 = strMan.GetStringInFile("DrawablesExhausted", "C:\\NoxPost\\src\\Client\\Gui\\guiinv.c")
			nox_xxx_printCentered_445490(v7)
			return 0
		}
		*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*30)) |= 0x40000000
		libc.MemCpy(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(v4))+432)), unsafe.Pointer(uintptr(a1+432)), 24)
		*(*uint16)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(v4)) + 292))) = *(*uint16)(unsafe.Pointer(uintptr(a1 + 292)))
		*(*uint16)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(v4)) + 294))) = *(*uint16)(unsafe.Pointer(uintptr(a1 + 294)))
	}
	v8 = (*uint8)(unsafe.Pointer(&array_5D4594_1049872[0]))
	*(*uint32)(unsafe.Add(unsafe.Pointer(v4), int32(func() uint8 {
		p := (*uint8)(unsafe.Add(unsafe.Pointer(v4), 140))
		x := *p
		*p++
		return x
	}())*4+4)) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 128)))
	*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v4))), unsafe.Sizeof(uint32(0))*33))) = 0
	for {
		v9 = int32(*(*uint32)(unsafe.Pointer(v8)))
		if *(*uint32)(unsafe.Pointer(v8)) != 0 {
			break
		}
	LABEL_17:
		v8 = (*uint8)(unsafe.Add(unsafe.Pointer(v8), 4))
		if uintptr(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v8)))))) >= uintptr(unsafe.Pointer(&array_5D4594_1049872[9])) {
			return 1
		}
	}
	for *(*uint32)(unsafe.Pointer(uintptr(v9 + 128))) != *(*uint32)(unsafe.Pointer(uintptr(a1 + 128))) {
		v9 = int32(*(*uint32)(unsafe.Pointer(uintptr(v9 + 368))))
		if v9 == 0 {
			goto LABEL_17
		}
	}
	v10 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v4))), unsafe.Sizeof(uint32(0))*34))))
	*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v4))), unsafe.Sizeof(uint32(0))*33))) = 1
	if v10 != 0 {
		nox_xxx_clientSetAltWeapon_461550(0)
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v4))), unsafe.Sizeof(uint32(0))*34))) = 0
	}
	return 1
}
func sub_464BD0(a1 int32, a2 int32, a3 uint32) int32 {
	var (
		v4  int32
		v5  int32
		v6  int32
		v7  int32
		v8  int32
		v9  int32
		v10 int32
		v14 int32
		v15 int32
		v16 int32
		v17 int32
		v19 int32
		v20 int32
		v21 int32
		v26 int32
		v28 *uint32
		v29 *libc.WChar
		v30 int32
		v31 int32
		v32 int32
		v33 int32
		v34 int32
		v36 int32
		v37 int32
		v38 int32
		v40 *uint32
		v41 int32
		v42 int32
		v45 int32
		v47 int32
		v48 int32
		v49 unsafe.Pointer
		v50 *libc.WChar
		v51 int2
		v52 int32
		v53 int32
		v54 int32
		v55 int32
		v56 int2
		v57 int2
		v58 int2
		v59 int2
	)
	v59.field_4 = int32(a3 >> 16)
	v59.field_0 = int32(uint16(a3))
	sub_463370(*(**uint32)(unsafe.Pointer(&dword_5d4594_1062456)), (*nox_point)(unsafe.Pointer(&v59)), (*uint32)(unsafe.Pointer(&v56)))
	if sub_45D9B0() != 0 || int32(*memmap.PtrUint8(6112660, 1049868)) != 2 {
		return 1
	}
	switch a2 {
	case 5:
		if nox_xxx_playerAnimCheck_4372B0() != 0 {
			return 1
		}
		if dword_5d4594_1049864 == 5 {
			v8 = sub_4281F0(&v56, (*int4)(memmap.PtrOff(0x587000, 136352)))
			if v8 != 0 {
				v9 = (v56.field_0 - 314) / 50
				v10 = int32((dword_5d4594_1062512 + uint32(v56.field_4) - 13) / 50)
				if sub_464B40(v9, v10) == 0 {
					return 1
				}
				var v11 int32 = v10 + NOX_INVENTORY_ROW_COUNT*v9
				if int32(nox_client_inventory_grid_1050020[v11].field_140) != 0 {
					var dr *nox_drawable = nox_client_inventory_grid_1050020[v11].field_0
					dword_5d4594_1063116 = uint32(uintptr(unsafe.Pointer(dr)))
					dr.field_32 = nox_client_inventory_grid_1050020[v11].field_4
				} else {
					dword_5d4594_1063116 = 0
				}
				return 1
			}
			if sub_478030() != 0 {
				if sub_479870() != 0 {
					*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v14))), 0)) = uint8(int8(bool2int(sub_479880((*uint32)(unsafe.Pointer(&v56))))))
					if v14 != 0 {
						dword_5d4594_1063116 = uint32(sub_4798A0((*uint32)(unsafe.Pointer(&v56))))
						return 1
					}
				}
			}
		} else if dword_5d4594_1049864 == 6 {
			v15 = sub_4281F0(&v56, (*int4)(memmap.PtrOff(0x587000, 136352)))
			if v15 != 0 {
				v16 = (v56.field_0 - 314) / 50
				v17 = int32((dword_5d4594_1062512 + uint32(v56.field_4) - 13) / 50)
				if sub_464B40(v16, v17) != 0 {
					var v18 int32 = v17 + NOX_INVENTORY_ROW_COUNT*v16
					if int32(nox_client_inventory_grid_1050020[v18].field_140) != 0 {
						v19 = int32(uintptr(unsafe.Pointer(nox_client_inventory_grid_1050020[v18].field_0)))
						v20 = int32(nox_client_inventory_grid_1050020[v18].field_4)
						*(*uint32)(unsafe.Pointer(uintptr(v19 + 128))) = uint32(v20)
						if v19 != 0 {
							nox_xxx_trade_4657B0(int16(v20))
							return 1
						}
					}
				}
			}
		} else if int32(*memmap.PtrUint8(6112660, 1049870)) != 1 || (func() bool {
			v21 = sub_4281F0(&v56, (*int4)(memmap.PtrOff(0x587000, 136336)))
			return v21 != 1
		}()) {
			if int32(*memmap.PtrUint8(6112660, 1049869)) != 0 || sub_4281F0(&v56, (*int4)(memmap.PtrOff(0x587000, 136384))) != 0 || sub_4281F0(&v56, (*int4)(memmap.PtrOff(0x587000, 136400))) != 0 {
				if sub_4281F0(&v56, (*int4)(memmap.PtrOff(0x587000, 136384))) == 1 {
					return 0
				}
				if sub_4281F0(&v56, (*int4)(memmap.PtrOff(0x587000, 136400))) == 1 {
					return 0
				}
			} else {
				nox_xxx_wndSetCaptureMain((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1062456))))))
				if sub_479590() == 3 {
					nox_xxx_clientTradeMB_4657E0((*uint32)(unsafe.Pointer(&v56)))
				} else {
					sub_4658A0(*(*int32)(unsafe.Pointer(&dword_5d4594_1062456)), &v56)
				}
				if *memmap.PtrUint32(6112660, 1049848) != 0 {
					nox_xxx_cursorSetDraggedItem_477690((*nox_drawable)(unsafe.Pointer(uintptr(*memmap.PtrInt32(6112660, 1049848)))))
					nox_xxx_setKeybTimeout_4160D0(0)
					*(*int2)(memmap.PtrOff(6112660, 1062572)) = v56
					clientPlaySoundSpecial(791, 100)
					return 1
				}
			}
		}
		return 1
	case 6:
		goto LABEL_53
	case 7:
		if nox_xxx_playerAnimCheck_4372B0() != 0 || dword_5d4594_1049864 == 6 {
			return 1
		}
		if int32(*memmap.PtrUint8(6112660, 1049869)) == 0 {
			v26 = sub_4281F0(&v56, (*int4)(memmap.PtrOff(0x587000, 136368)))
			if v26 != 0 {
				if (v56.field_4-13)/50 == 1 {
					if dword_5d4594_1049864 != 5 {
						sub_465CA0()
						return 1
					}
				LABEL_5:
					sub_462740()
					return 1
				}
			}
		}
	LABEL_53:
		if nox_xxx_playerAnimCheck_4372B0() != 0 || dword_5d4594_1049864 == 6 {
			return 1
		}
		if dword_5d4594_1049864 == 5 {
			if nox_xxx_cursorGetTypePrev_477630() == 7 {
				goto LABEL_5
			}
		} else {
			nox_xxx_wndClearCaptureMain((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1062456))))))
		}
		if dword_5d4594_1049864 == 4 {
			v58 = v59
			sub_473970(&v58, &v58)
			v28 = &nox_drawable_find_49ABF0((*nox_point)(unsafe.Pointer(&v58)), 20).field_0
			if v28 != nil {
				v57.field_0 = nox_win_width / 2
				v57.field_4 = nox_win_height / 2
				sub_473970(&v57, &v57)
				if (uint32(v57.field_0)-*(*uint32)(unsafe.Add(unsafe.Pointer(v28), unsafe.Sizeof(uint32(0))*3)))*(uint32(v57.field_0)-*(*uint32)(unsafe.Add(unsafe.Pointer(v28), unsafe.Sizeof(uint32(0))*3)))+(uint32(v57.field_4)-*(*uint32)(unsafe.Add(unsafe.Pointer(v28), unsafe.Sizeof(uint32(0))*4)))*(uint32(v57.field_4)-*(*uint32)(unsafe.Add(unsafe.Pointer(v28), unsafe.Sizeof(uint32(0))*4))) <= 5625 {
				LABEL_64:
					dword_5d4594_1049864 = 0
					return 1
				}
				v29 = strMan.GetStringInFile("ObjectTooFar", "C:\\NoxPost\\src\\Client\\Gui\\guiinv.c")
			} else {
				v29 = strMan.GetStringInFile("NoObject", "C:\\NoxPost\\src\\Client\\Gui\\guiinv.c")
			}
			nox_xxx_printCentered_445490(v29)
			goto LABEL_64
		}
		if *memmap.PtrUint32(6112660, 1049848) == 0 {
			return 1
		}
		if !nox_xxx_wndPointInWnd_46AAB0(*(**uint32)(unsafe.Pointer(&dword_5d4594_1062456)), v59.field_0, v59.field_4) || (func() int32 {
			v30 = sub_4281F0(&v56, (*int4)(memmap.PtrOff(0x587000, 136384)))
			return v30
		}()) != 0 || (func() int32 {
			v31 = sub_4281F0(&v56, (*int4)(memmap.PtrOff(0x587000, 136400)))
			return v31
		}()) != 0 {
			v58 = v59
			sub_473970(&v58, &v57)
			if dword_5d4594_1049856 == 1 {
				if sub_4C12C0() == 0 {
					nox_xxx_clientDrop_465BE0(&v57)
				}
			} else {
				v47 = int32(dword_5d4594_1049800_inventory_click_row_index + dword_5d4594_1049796_inventory_click_column_index*14 + dword_5d4594_1049796_inventory_click_column_index*7)
				v48 = int32(nox_client_inventory_grid_1050020[v47].field_140)
				if int32(nox_client_inventory_grid_1050020[v47].field_140) != 0 {
					v49 = nil
					nox_xxx_wndClearCaptureMain((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1062456))))))
					if *(*uint32)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, 1049848) + 112)))&0x13001000 != 0 {
						v49 = unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, 1049848) + 432))
					}
					sub_4C05F0(0, 0)
					v53 = int32(*(*uint32)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, 1049848) + 108))))
					v52 = int32(*(*uint32)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, 1049848) + 128))))
					v51 = v58
					v50 = strMan.GetStringInFile("DropLabel", "C:\\NoxPost\\src\\Client\\Gui\\guiinv.c")
					nox_gui_itemAmountDialog_4C0430((*libc.WChar)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v50)))))), v51.field_0, v51.field_4, v52, v53, v49, v48+1, 0, unsafe.Pointer(cgoFuncAddr(sub_465CD0)), nil)
				} else if sub_4C12C0() == 0 {
					nox_xxx_clientDrop_465BE0(&v57)
				}
			}
			if dword_5d4594_1049856 != 0 {
				goto LABEL_121
			}
		LABEL_119:
			v55 = int32(dword_5d4594_1049800_inventory_click_row_index)
			v54 = int32(dword_5d4594_1049796_inventory_click_column_index)
			goto LABEL_120
		}
		v32 = int32(*memmap.PtrUint32(6112660, 1062572) - uint32(v56.field_0))
		v33 = int32(*memmap.PtrUint32(6112660, 1062576) - uint32(v56.field_4))
		if !nox_xxx_checkKeybTimeout_4160F0(0, nox_gameFPS/3) && v32*v32+v33*v33 < 100 {
			v34 = sub_4281F0(&v56, (*int4)(memmap.PtrOff(0x587000, 136352)))
			if v34 == 0 {
				goto LABEL_121
			}
			if sub_4C12C0() == 0 {
				if *(*uint32)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, 1049848) + 112)))&0x3001000 != 0 {
					var v35 int32 = int32(dword_5d4594_1049800_inventory_click_row_index + NOX_INVENTORY_ROW_COUNT*dword_5d4594_1049796_inventory_click_column_index)
					if nox_client_inventory_grid_1050020[v35].field_136 != 0 {
						nox_xxx_clientSetAltWeapon_461550(0)
						nox_client_inventory_grid_1050020[v35].field_136 = 0
					} else if nox_client_inventory_grid_1050020[v35].field_132 != 0 {
						nox_xxx_clientDequip_464B70(*memmap.PtrInt32(6112660, 1049848))
					} else {
						nox_xxx_clientKeyEquip_465C30(*(*int32)(unsafe.Pointer(&dword_5d4594_1049796_inventory_click_column_index)), *(*int32)(unsafe.Pointer(&dword_5d4594_1049800_inventory_click_row_index)))
					}
				} else {
					nox_xxx_clientUse_465C70(*memmap.PtrInt32(6112660, 1049848))
				}
			}
			goto LABEL_80
		}
		v36 = sub_4281F0(&v56, (*int4)(memmap.PtrOff(0x587000, 136336)))
		if v36 != 0 && int32(*memmap.PtrUint8(6112660, 1049870)) == 0 {
			if dword_5d4594_1049856 == 0 {
				nox_xxx_clientEquip_4623B0(*memmap.PtrInt32(6112660, 1049848))
				sub_4649B0(*memmap.PtrInt32(6112660, 1049848), *(*int32)(unsafe.Pointer(&dword_5d4594_1049796_inventory_click_column_index)), *(*int32)(unsafe.Pointer(&dword_5d4594_1049800_inventory_click_row_index)))
			}
			goto LABEL_121
		}
		v37 = sub_4281F0(&v56, (*int4)(memmap.PtrOff(0x587000, 136352)))
		if v37 == 0 {
			goto LABEL_119
		}
		v38 = int32(*(*uint32)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, 1049848) + 108))))
		if uint32(v38) == dword_5d4594_1062560 || uint32(v38) == *memmap.PtrUint32(6112660, 1049728) || uint32(v38) == *memmap.PtrUint32(6112660, 1049724) || uint32(v38) == dword_5d4594_1062556 || uint32(v38) == dword_5d4594_1062564 {
			sub_4649B0(*memmap.PtrInt32(6112660, 1049848), *(*int32)(unsafe.Pointer(&dword_5d4594_1049796_inventory_click_column_index)), *(*int32)(unsafe.Pointer(&dword_5d4594_1049800_inventory_click_row_index)))
			goto LABEL_121
		}
		dword_5d4594_1049804 = uint32((v56.field_0 - 314) / 50)
		dword_5d4594_1049808 = (dword_5d4594_1062512 + uint32(v56.field_4) - 13) / 50
		if sub_464B40((v56.field_0-314)/50, int32((dword_5d4594_1062512+uint32(v56.field_4)-13)/50)) == 0 {
			goto LABEL_121
		}
		if dword_5d4594_1049856 != 0 {
			var v39 int32 = int32(dword_5d4594_1049808 + NOX_INVENTORY_ROW_COUNT*dword_5d4594_1049804)
			if int32(nox_client_inventory_grid_1050020[v39].field_140) != 0 && (func() *uint32 {
				v40 = &nox_client_inventory_grid_1050020[v39].field_0.field_0
				return v40
			}()) != nil && ((func() uint32 {
				v41 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v40), unsafe.Sizeof(uint32(0))*28)))
				return uint32(v41) & 0x2000000
			}()) != 0 && *(*uint32)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, 1049848) + 112)))&0x2000000 != 0 && *(*uint32)(unsafe.Add(unsafe.Pointer(v40), unsafe.Sizeof(uint32(0))*29)) == *(*uint32)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, 1049848) + 116))) || uint32(v41)&0x1001000 != 0 && *(*uint32)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, 1049848) + 112)))&0x1001000 != 0) {
				v42 = int32(nox_client_inventory_grid_1050020[v39].field_4)
				*memmap.PtrUint32(6112660, 1049860) = 1
				*(*uint32)(unsafe.Add(unsafe.Pointer(v40), unsafe.Sizeof(uint32(0))*32)) = uint32(v42)
				nox_xxx_clientEquip_4623B0(int32(uintptr(unsafe.Pointer(v40))))
			} else {
				*memmap.PtrUint32(6112660, 1049860) = 1
				nox_xxx_clientDequip_464B70(*memmap.PtrInt32(6112660, 1049848))
			}
			goto LABEL_121
		}
		if int32(nox_client_inventory_grid_1050020[dword_5d4594_1049800_inventory_click_row_index+NOX_INVENTORY_ROW_COUNT*dword_5d4594_1049796_inventory_click_column_index].field_140) != 0 {
			v55 = int32(dword_5d4594_1049800_inventory_click_row_index)
			v54 = int32(dword_5d4594_1049796_inventory_click_column_index)
		LABEL_120:
			sub_4649B0(*memmap.PtrInt32(6112660, 1049848), v54, v55)
			goto LABEL_121
		}
		if sub_4649B0(*memmap.PtrInt32(6112660, 1049848), *(*int32)(unsafe.Pointer(&dword_5d4594_1049804)), *(*int32)(unsafe.Pointer(&dword_5d4594_1049808))) == 0 {
		LABEL_80:
			sub_4649B0(*memmap.PtrInt32(6112660, 1049848), *(*int32)(unsafe.Pointer(&dword_5d4594_1049796_inventory_click_column_index)), *(*int32)(unsafe.Pointer(&dword_5d4594_1049800_inventory_click_row_index)))
			goto LABEL_121
		}
		clientPlaySoundSpecial(792, 100)
		var v43 int32 = int32(dword_5d4594_1049800_inventory_click_row_index + NOX_INVENTORY_ROW_COUNT*dword_5d4594_1049796_inventory_click_column_index)
		v45 = int32(nox_client_inventory_grid_1050020[v43].field_136)
		if v45 != 0 {
			var v46 int32 = int32(dword_5d4594_1049808 + NOX_INVENTORY_ROW_COUNT*dword_5d4594_1049804)
			nox_client_inventory_grid_1050020[v46].field_136 = uint32(v45)
			nox_client_inventory_grid_1050020[v43].field_136 = 0
			dword_5d4594_1062480 = uint32(uintptr(unsafe.Pointer(&nox_client_inventory_grid_1050020[v46])))
		}
		sub_461B50()
	LABEL_121:
		nox_xxx_cursorResetDraggedItem_4776A0()
		if dword_5d4594_1049856 == 0 {
			nox_xxx_spriteDelete_45A4B0((*nox_drawable)(unsafe.Pointer(*(**uint64)(memmap.PtrOff(6112660, 1049848)))))
		}
		*memmap.PtrUint32(6112660, 1049848) = 0
		dword_5d4594_1049856 = 0
		return 1
	case 8:
		return 1
	case 9:
		if dword_5d4594_1049864 == 5 {
			goto LABEL_5
		}
		return 0
	case 19:
		if nox_xxx_playerAnimCheck_4372B0() != 0 {
			return 1
		}
		v6 = sub_4281F0(&v56, (*int4)(memmap.PtrOff(0x587000, 136384)))
		if v6 != 0 {
			goto LABEL_124
		}
		v7 = sub_4281F0(&v56, (*int4)(memmap.PtrOff(0x587000, 136400)))
		if v7 != 0 {
			goto LABEL_124
		}
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1062456))))).Func94(asWindowEvent(0x4007, *memmap.PtrInt32(6112660, 1062500), 0))
		return 1
	case 20:
		if nox_xxx_playerAnimCheck_4372B0() != 0 {
			return 1
		}
		v4 = sub_4281F0(&v56, (*int4)(memmap.PtrOff(0x587000, 136384)))
		if v4 != 0 {
			goto LABEL_124
		}
		v5 = sub_4281F0(&v56, (*int4)(memmap.PtrOff(0x587000, 136400)))
		if v5 != 0 {
			goto LABEL_124
		}
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1062456))))).Func94(asWindowEvent(0x4007, *memmap.PtrInt32(6112660, 1062504), 0))
		return 1
	default:
	LABEL_124:
		if dword_5d4594_1049864 == 5 {
			return 1
		}
		return 0
	}
}
func nox_xxx_cliInventorySpriteUpd_465A30() {
	var inventory_item_idx int32 = int32(dword_5d4594_1049800_inventory_click_row_index + NOX_INVENTORY_ROW_COUNT*dword_5d4594_1049796_inventory_click_column_index)
	if int32(nox_client_inventory_grid_1050020[inventory_item_idx].field_140) != 0 {
		var v1 *uint32 = &nox_new_drawable_for_thing(int32(nox_client_inventory_grid_1050020[inventory_item_idx].field_0.field_27)).field_0
		*memmap.PtrUint32(6112660, 1049848) = uint32(uintptr(unsafe.Pointer(v1)))
		if v1 != nil {
			*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*30)) |= 0x40000000
			*(*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer(v1))) + 128))) = nox_client_inventory_grid_1050020[inventory_item_idx].field_4
			libc.MemCpy(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer(v1)))+432)), unsafe.Pointer(&nox_client_inventory_grid_1050020[inventory_item_idx].field_0.field_108_1), 24)
			*(*uint16)(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer(v1))) + 292))) = nox_client_inventory_grid_1050020[inventory_item_idx].field_0.field_73_1
			*(*uint16)(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer(v1))) + 294))) = nox_client_inventory_grid_1050020[inventory_item_idx].field_0.field_73_2
			var v3 [2]*int32
			v3[0] = (*int32)(unsafe.Pointer(&nox_client_inventory_grid_1050020[inventory_item_idx].field_0))
			v3[1] = nil
			sub_461E60((***uint64)(unsafe.Pointer(&v3[0])))
		} else {
			var v2 *libc.WChar = strMan.GetStringInFile("DrawablesExhausted", "C:\\NoxPost\\src\\Client\\Gui\\guiinv.c")
			nox_xxx_printCentered_445490(v2)
		}
	}
}
func sub_466160() int32 {
	var v0 *libc.WChar
	if int32(*memmap.PtrUint8(6112660, 1049868)) == 2 {
		v0 = strMan.GetStringInFile("CloseInventoryTT", "C:\\NoxPost\\src\\Client\\Gui\\guiinv.c")
	} else {
		v0 = strMan.GetStringInFile("OpenInventoryTT", "C:\\NoxPost\\src\\Client\\Gui\\guiinv.c")
	}
	nox_xxx_cursorSetTooltip_4776B0(v0)
	return 1
}
func sub_4661D0() int32 {
	var (
		v0 *libc.WChar
		v2 *libc.WChar
	)
	if dword_5d4594_1062480 != 0 {
		v0 = nox_xxx_clientAskInfoMb_4BF050((*nox_drawable)(unsafe.Pointer(**(***libc.WChar)(unsafe.Pointer(&dword_5d4594_1062480)))))
		nox_xxx_cursorSetTooltip_4776B0(v0)
	} else {
		v2 = strMan.GetStringInFile("ToolTipWeapon2Area", "C:\\NoxPost\\src\\Client\\Gui\\guiinv.c")
		nox_xxx_cursorSetTooltip_4776B0(v2)
	}
	return 1
}
func sub_466660(a1 int32, a2 *int2) *libc.WChar {
	var (
		v2     int32
		v3     int32
		result *libc.WChar
		v5     *libc.WChar
		v6     int32
		v7     int32
		v8     int32
		v9     int32
		v10    int32
		v13    int32
	)
	v2 = sub_4281F0(a2, (*int4)(memmap.PtrOff(0x587000, 136336)))
	if v2 == 0 {
		if int32(*memmap.PtrUint8(6112660, 1049869)) == 0 {
			v6 = sub_4281F0(a2, (*int4)(memmap.PtrOff(0x587000, 136368)))
			if v6 != 0 {
				v7 = a2.field_4 - 13
				v8 = v7 / 50
				v9 = 20
				dword_5d4594_1049796_inventory_click_column_index = uint32(v7 / 50)
			} else {
				v10 = sub_4281F0(a2, (*int4)(memmap.PtrOff(0x587000, 136352)))
				if v10 == 0 {
					return nil
				}
				v8 = (a2.field_0 - 314) / 50
				dword_5d4594_1049796_inventory_click_column_index = uint32((a2.field_0 - 314) / 50)
				v9 = int32((uint32(a2.field_4) + dword_5d4594_1062512 - 13) / 50)
			}
			dword_5d4594_1049800_inventory_click_row_index = uint32(v9)
			if sub_464B40(v8, v9) != 0 {
				var v12 int32 = int32(dword_5d4594_1049800_inventory_click_row_index + NOX_INVENTORY_ROW_COUNT*dword_5d4594_1049796_inventory_click_column_index)
				if int32(nox_client_inventory_grid_1050020[v12].field_140) != 0 {
					v13 = int32(uintptr(unsafe.Pointer(nox_client_inventory_grid_1050020[v12].field_0)))
					*(*uint32)(unsafe.Pointer(uintptr(v13 + 128))) = nox_client_inventory_grid_1050020[v12].field_4
					return nox_xxx_clientAskInfoMb_4BF050((*nox_drawable)(unsafe.Pointer(uintptr(v13))))
				}
			}
		}
		return nil
	}
	if int32(*memmap.PtrUint8(6112660, 1049870)) == 1 {
		return nil
	}
	v3 = sub_465990((*uint32)(unsafe.Pointer(a2)))
	if v3 == -1 {
		return strMan.GetStringInFile("DollRegionError", "C:\\NoxPost\\src\\Client\\Gui\\guiinv.c")
	}
	v5 = *(**libc.WChar)(unsafe.Pointer(&array_5D4594_1049872[v3]))
	if v5 != nil {
		result = nox_xxx_clientAskInfoMb_4BF050((*nox_drawable)(unsafe.Pointer(v5)))
	} else {
		result = strMan.GetStringInFile("ToolTipDrag", "C:\\NoxPost\\src\\Client\\Gui\\guiinv.c")
	}
	return result
}
func nox_xxx_inventroryOnHovewerSub_4667E0(a1 int32, a2 int32, a3 uint32) int32 {
	var (
		v3     int32
		v4     int32
		v5     int32
		v6     int32
		v7     *libc.WChar
		result int32
		v9     int32
		v10    *libc.WChar
		v11    int32
		v12    *libc.WChar
		v13    int32
		v14    *libc.WChar
		v15    int2
	)
	v3 = 40
	v15.field_0 = int32(uint16(a3))
	v15.field_4 = int32(a3 >> 16)
	v4 = 0
	for v3 <= int32(uint16(a3)) {
		v3 += 35
		v4++
	}
	v5 = 0
	for {
		if uint32(1<<v5)&*memmap.PtrUint32(6112660, 1062540) != 0 && v5 != 31 {
			v4--
		}
		if v4 < 0 {
			break
		}
		v5++
		if v5 >= 32 {
			break
		}
	}
	if v5 != 32 {
		v6 = nox_xxx_getEnchantSpell_424920(v5)
		v7 = nox_xxx_spellTitle_424930(v6)
		nox_xxx_cursorSetTooltip_4776B0(v7)
		return 1
	}
	v9 = 0
	for {
		if (1<<v9)&int32(*memmap.PtrUint8(6112660, 1062536)) != 0 {
			v4--
		}
		if v4 < 0 {
			break
		}
		v9++
		if v9 >= 6 {
			break
		}
	}
	if v9 != 6 {
		v10 = sub_413480(int8(1 << v9))
		nox_xxx_cursorSetTooltip_4776B0(v10)
		return 1
	}
	if !noxflags.HasGame(noxflags.GameModeQuest) {
		goto LABEL_28
	}
	v11 = sub_4281F0(&v15, (*int4)(memmap.PtrOff(6112660, 1049812)))
	if v11 == 1 {
		v12 = strMan.GetStringInFile("thing.db:AnkhGUI", "C:\\NoxPost\\src\\Client\\Gui\\guiinv.c")
		nox_xxx_cursorSetTooltip_4776B0(v12)
		return 1
	}
	v13 = sub_4281F0(&v15, (*int4)(memmap.PtrOff(6112660, 1049828)))
	if v13 == 1 && sub_4BFD30() == 1 {
		v14 = strMan.GetStringInFile("GeneralPrint:TooltipKeyIcon", "C:\\NoxPost\\src\\Client\\Gui\\guiinv.c")
		nox_xxx_cursorSetTooltip_4776B0(v14)
		result = 1
	} else {
	LABEL_28:
		nox_xxx_cursorSetTooltip_4776B0(nil)
		result = 1
	}
	return result
}
func sub_466E20(a1 *uint32) int32 {
	var (
		v1     *libc.WChar
		result int32
	)
	switch *a1 {
	case 9105:
		v1 = strMan.GetStringInFile("JournalModeTT", "C:\\NoxPost\\src\\Client\\Gui\\guiinv.c")
		goto LABEL_7
	case 9106:
		v1 = strMan.GetStringInFile("InventoryModeTT", "C:\\NoxPost\\src\\Client\\Gui\\guiinv.c")
		goto LABEL_7
	case 9107:
		v1 = strMan.GetStringInFile("StatsModeTT", "C:\\NoxPost\\src\\Client\\Gui\\guiinv.c")
		goto LABEL_7
	case 9108:
		v1 = strMan.GetStringInFile("PaperDollModeTT", "C:\\NoxPost\\src\\Client\\Gui\\guiinv.c")
		goto LABEL_7
	case 9111:
		v1 = strMan.GetStringInFile("CloseInventoryTT", "C:\\NoxPost\\src\\Client\\Gui\\guiinv.c")
	LABEL_7:
		nox_xxx_cursorSetTooltip_4776B0(v1)
		result = 1
	default:
		result = 0
	}
	return result
}
func nox_xxx_inventoryNameSignInit_4671E0() int32 {
	var (
		result int32
		v1     int32
		v2     *libc.WChar
		v3     int32
		v4     *libc.WChar
		v5     [100]byte
	)
	nox_wcscpy((*libc.WChar)(memmap.PtrOff(6112660, 1062588)), (*libc.WChar)(memmap.PtrOff(6112660, 1063676)))
	if noxflags.HasGame(noxflags.GameModeQuest) || nox_xxx_isQuest_4D6F50() != 0 || (func() int32 {
		result = sub_4D6F70()
		return result
	}()) != 0 {
		result = int32(dword_5d4594_1049844)
		if dword_5d4594_1049844 > NOX_PLAYER_MAX_LEVEL {
			result = NOX_PLAYER_MAX_LEVEL
		}
		v1 = int32(*memmap.PtrUint32(0x8531A0, 2576))
	} else {
		v1 = int32(*memmap.PtrUint32(0x8531A0, 2576))
		if *memmap.PtrUint32(0x8531A0, 2576) == 0 {
			return result
		}
		result = int32(*(*byte)(unsafe.Pointer(uintptr(*memmap.PtrUint32(0x8531A0, 2576) + 3684))))
	}
	if v1 != 0 {
		nox_sprintf(&v5[0], CString("experience:%s%d"), *memmap.PtrUint32(0x587000, uint32(int32(*(*uint8)(unsafe.Pointer(uintptr(v1 + 2251))))*4+0x7310)), result)
		v4 = strMan.GetStringInFile(&v5[0], "C:\\NoxPost\\src\\Client\\Gui\\guiinv.c")
		v3 = int32(*memmap.PtrUint32(0x8531A0, 2576) + 4704)
		v2 = strMan.GetStringInFile("ElaborateNameFormat", "C:\\NoxPost\\src\\Client\\Gui\\guiinv.c")
		result = nox_swprintf((*libc.WChar)(memmap.PtrOff(6112660, 1062588)), v2, v3, v4)
	}
	return result
}
func sub_467750(a1 int32, a2 int8) int32 {
	if a1 != 0 {
		var v2 *uint32 = (*uint32)(unsafe.Pointer(sub_461EF0(a1)))
		if v2 != nil {
			if dword_5d4594_1062480 != 0 {
				*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1062480 + 136))) = 0
			}
			dword_5d4594_1062480 = *v2
			*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1062480 + 136))) = 1
			return 1
		}
	} else {
		if dword_5d4594_1062480 != 0 {
			*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1062480 + 136))) = 0
			dword_5d4594_1062480 = 0
		}
	}
	if int32(a2) != 0 {
		if int32(a2) != 1 {
			return 0
		}
		var v5 *libc.WChar = strMan.GetStringInFile("Weapon2CantUse", "C:\\NoxPost\\src\\Client\\Gui\\guiinv.c")
		nox_xxx_printCentered_445490(v5)
		if dword_5d4594_1062484 == 0 {
			return 0
		}
		var v6 *int32 = (*int32)(unsafe.Pointer(sub_461EF0(*(*int32)(unsafe.Pointer(&dword_5d4594_1062484)))))
		if v6 != nil {
			nox_xxx_clientSetAltWeapon_461550(*v6)
			return 0
		}
	}
	dword_5d4594_1062484 = 0
	return 0
}
