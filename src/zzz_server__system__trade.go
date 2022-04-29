package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	noxflags "github.com/noxworld-dev/opennox/v1/common/flags"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"math"
	"unsafe"
)

func sub_50E7A0(a1 *uint32, a2 int32) int32 {
	var (
		v2 *int32
		v4 int32
		v5 int32
		v6 int32
	)
	v2 = (*int32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*5)))))
	if v2 == nil {
		return 0
	}
	for *v2 != a2 {
		v2 = (*int32)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(int32(0))*2)))))
		if v2 == nil {
			return 0
		}
	}
	v4 = *(*int32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(int32(0))*2))
	if v4 != 0 {
		*(*uint32)(unsafe.Pointer(uintptr(v4 + 12))) = uint32(*(*int32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(int32(0))*3)))
	}
	v5 = *(*int32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(int32(0))*3))
	if v5 != 0 {
		*(*uint32)(unsafe.Pointer(uintptr(v5 + 8))) = uint32(*(*int32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(int32(0))*2)))
	}
	if v2 == (*int32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*5))))) {
		*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*5)) = uint32(*(*int32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(int32(0))*2)))
	}
	v6 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*2)))
	if (int32(*(*uint8)(unsafe.Pointer(uintptr(v6 + 8)))) & 4) == 0 {
		v6 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*3)))
	}
	sub_50E820(v6, *v2)
	nox_alloc_class_free_obj_first((*nox_alloc_class)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&nox_alloc_tradeItems_2386496)))), unsafe.Pointer(v2))
	return 1
}
func nox_xxx_createPlayerShopSession_50E8F0(a1 int32, a2 int32) *uint32 {
	var (
		v2     int32
		v3     int32
		v4     *uint32
		result *uint32
	)
	v2 = 0
	if noxflags.HasGame(noxflags.GameModeQuest) && (func() *uint32 {
		v3 = int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))) + 276))) + 2064))))
		v4 = *(**uint32)(memmap.PtrOff(6112660, uint32(v3*4)+0x2469BC))
		*memmap.PtrUint32(6112660, uint32(v3*4)+0x2469BC) = 0
		return v4
	}()) != nil {
		v2 = 1
	} else {
		result = nox_xxx_createShopStruct_50E870()
		v4 = result
		if result == nil {
			return result
		}
	}
	*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*2)) = uint32(a1)
	*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*3)) = uint32(a2)
	*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*4)) = 1
	if v2 == 0 {
		nox_xxx_loadShopItems_50E970(int32(uintptr(unsafe.Pointer(v4))))
	}
	return v4
}
func sub_50F0F0(a1 int32, a2 int32) int32 {
	var (
		v2 int32
		v3 int32
		v4 *libc.WChar
		v6 [86]byte
	)
	*(*uint16)(unsafe.Pointer(&v6[0])) = 3529
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 8))))
	if int32(*(*uint8)(unsafe.Pointer(uintptr(v2 + 8))))&4 != 0 {
		v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 8))))
		v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 12))))
	} else {
		v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 12))))
	}
	v4 = sub_4E39F0_obj_db((**byte)(unsafe.Pointer(uintptr(v2))))
	nox_wcsncpy((*libc.WChar)(unsafe.Pointer(&v6[4])), v4, 24)
	*(*uint16)(unsafe.Pointer(&v6[52])) = 0
	libc.StrCpy(&v6[54], (*byte)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 692)))+1684))))
	*(*uint16)(unsafe.Pointer(&v6[2])) = *(*uint16)(unsafe.Pointer(uintptr(v2 + 4)))
	return nox_xxx_netSendPacket1_4E5390(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v3 + 748))) + 276))) + 2064)))), int32(uintptr(unsafe.Pointer(&v6[0]))), 86, 0, 1)
}
func sub_50F1A0(a1 int32, a2 int32) int32 {
	var (
		v2 int32
		v3 int32
		v4 int32
		v6 *libc.WChar
		v8 [52]byte
	)
	*(*uint16)(unsafe.Pointer(&v8[0])) = 3273
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 8))))
	v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
	if a1 != v2 {
		if int32(*(*uint8)(unsafe.Pointer(uintptr(v2 + 8))))&4 != 0 {
			nox_wcscpy((*libc.WChar)(unsafe.Pointer(&v8[2])), (*libc.WChar)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 748))) + 276)))+4704))))
			return nox_xxx_netSendPacket1_4E5390(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v3 + 276))) + 2064)))), int32(uintptr(unsafe.Pointer(&v8[0]))), 52, 0, 1)
		}
		v6 = sub_4E39F0_obj_db(*(***byte)(unsafe.Pointer(uintptr(a2 + 12))))
	} else {
		v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 12))))
		if int32(*(*uint8)(unsafe.Pointer(uintptr(v4 + 8))))&4 != 0 {
			nox_wcscpy((*libc.WChar)(unsafe.Pointer(&v8[2])), (*libc.WChar)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v4 + 748))) + 276)))+4704))))
			return nox_xxx_netSendPacket1_4E5390(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v3 + 276))) + 2064)))), int32(uintptr(unsafe.Pointer(&v8[0]))), 52, 0, 1)
		}
		v6 = sub_4E39F0_obj_db((**byte)(unsafe.Pointer(uintptr(v4))))
	}
	nox_wcsncpy((*libc.WChar)(unsafe.Pointer(&v8[2])), v6, 24)
	*(*uint16)(unsafe.Pointer(&v8[50])) = 0
	return nox_xxx_netSendPacket1_4E5390(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v3 + 276))) + 2064)))), int32(uintptr(unsafe.Pointer(&v8[0]))), 52, 0, 1)
}
func nox_xxx_servSendShopItems_50F280(a1 int32, a2 int32) int32 {
	var (
		result int32
		i      *uint32
	)
	result = a2
	for i = *(**uint32)(unsafe.Pointer(uintptr(a2 + 20))); i != nil; i = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(i), unsafe.Sizeof(uint32(0))*2))))) {
		result = sub_50F2B0(a1, i)
	}
	return result
}
func nox_xxx_tradeSetPlayer_50F370(a1 *uint32, a2 int32) *uint32 {
	var result *uint32
	result = a1
	*a1 = 1
	if int32(*(*uint8)(unsafe.Pointer(uintptr(a2 + 8))))&4 != 0 {
		*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a2 + 748))) + 280))) = uint32(uintptr(unsafe.Pointer(a1)))
	}
	return result
}
func sub_50FAE0(a1 int32, a2 int32, a3 int32, a4 int32, a5 int32) int32 {
	var (
		v5  int32
		v6  int16
		v7  int32
		v8  int32
		i   int32
		v11 [15]byte
	)
	v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
	v11[0] = 201
	v11[2] = byte(int8(bool2int(a1 == a2)))
	v11[1] = 4
	v6 = int16(*(*uint16)(unsafe.Pointer(uintptr(a4 + 36))))
	*(*uint16)(unsafe.Pointer(&v11[3])) = *(*uint16)(unsafe.Pointer(uintptr(a4 + 4)))
	*(*uint32)(unsafe.Pointer(&v11[7])) = uint32(a5)
	v7 = int32(*(*uint32)(unsafe.Pointer(uintptr(a4 + 8))))
	*(*uint16)(unsafe.Pointer(&v11[5])) = uint16(v6)
	if uint32(v7)&0x13001000 != 0 {
		v8 = int32(*(*uint32)(unsafe.Pointer(uintptr(a4 + 692))))
		for i = 0; i < 4; i++ {
			if *(*uint32)(unsafe.Pointer(uintptr(v8))) != 0 {
				v11[i+11] = byte(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v8))) + 4))))
			} else {
				v11[i+11] = math.MaxUint8
			}
			v8 += 4
		}
	} else {
		*(*uint32)(unsafe.Pointer(&v11[11])) = math.MaxUint32
	}
	return nox_xxx_netSendPacket1_4E5390(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v5 + 276))) + 2064)))), int32(uintptr(unsafe.Pointer(&v11[0]))), 15, 0, 1)
}
func sub_50FD60(a1 *uint32, a2 int32) int32 {
	var (
		v2  *uint32
		v3  uint32
		v4  uint32
		v5  int32
		v6  *int32
		v7  int32
		v8  int32
		v9  int32
		i   *int32
		v12 [4]int32
		v13 [4]int32
	)
	v13[0] = 0
	v12[0] = 0
	v2 = a1
	v13[1] = 0
	v12[1] = 0
	v3 = 0
	v13[2] = 0
	v12[2] = 0
	v13[3] = 0
	for v12[3] = 0; v2 != nil; v2 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*2))))) {
		v4 = 0
		v5 = int32(*(*uint16)(unsafe.Pointer(uintptr(*v2 + 4))))
		v6 = &v13[0]
		for *v6 != v5 {
			v4++
			v6 = (*int32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(int32(0))*1))
			if v4 >= 4 {
				v7 = v12[v3]
				v13[v3] = v5
				v12[func() uint32 {
					p := &v3
					x := *p
					*p++
					return x
				}()] = v7 + 1
				goto LABEL_6
			}
		}
		v12[v4]++
	LABEL_6:
	}
	v8 = 0
	if v3 == 0 {
		return 1
	}
	*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v9))), unsafe.Sizeof(uint16(0))*1)) = 0
	for i = &v13[0]; ; i = (*int32)(unsafe.Add(unsafe.Pointer(i), unsafe.Sizeof(int32(0))*1)) {
		*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v9))), unsafe.Sizeof(uint16(0))*0)) = *(*uint16)(unsafe.Pointer(uintptr(a2 + 4)))
		if *i == v9 && (*(*uint32)(unsafe.Pointer(uintptr(a2 + 8)))&0x13001000) == 0 {
			break
		}
		if uint32(func() int32 {
			p := &v8
			*p++
			return *p
		}()) >= v3 {
			return bool2int(v3 < 4)
		}
	}
	return 1
}
func sub_510320(a1 int32, a2 int32) {
	var (
		v2 int32
		v3 *uint8
		v4 int32
		v5 int32
		v6 int8
		v7 int32
		v8 int32
	)
	if a1 != 0 && a2 != 0 {
		v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 8))))
		if int32(*(*uint8)(unsafe.Pointer(uintptr(v2 + 8))))&4 != 0 {
			v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 12))))
		}
		v3 = *(**uint8)(unsafe.Pointer(uintptr(v2 + 692)))
		v4 = nox_xxx_getSomeShopData_5103A0(a2, a1)
		if v4 != -1 {
			v5 = int32(uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v3), v4*28)))))
			v6 = int8(int32(*(*uint8)(unsafe.Add(unsafe.Pointer(v3), v4*28+8))) - 1)
			*(*uint8)(unsafe.Pointer(uintptr(v5 + 8))) = uint8(v6)
			if int32(v6) == 0 {
				v7 = v4
				if v4 < int32(*v3)-1 {
					v8 = v5 + 4
					for {
						v7++
						libc.MemCpy(unsafe.Pointer(uintptr(v8)), unsafe.Pointer(uintptr(v8+28)), 28)
						v8 += 28
						if v7 >= int32(*v3)-1 {
							break
						}
					}
				}
				*v3--
			}
		}
	}
}
func sub_5104F0(a1 int32, a2 int16) int32 {
	var (
		v2 int32
		v4 [4]byte
	)
	*(*uint16)(unsafe.Pointer(&v4[2])) = uint16(a2)
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
	v4[0] = 201
	v4[1] = 27
	return nox_xxx_netSendPacket0_4E5420(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 276))) + 2064)))), unsafe.Pointer(&v4[0]), 4, 0, 1)
}
func sub_510540(a1 int32) int32 {
	if !noxflags.HasGame(noxflags.GameModeQuest) {
		return 1
	}
	if *memmap.PtrUint32(6112660, 2386520) == 0 {
		*memmap.PtrUint32(6112660, 2386520) = uint32(nox_xxx_getNameId_4E3AA0(CString("Diamond")))
		*memmap.PtrUint32(6112660, 2386528) = uint32(nox_xxx_getNameId_4E3AA0(CString("Emerald")))
		*memmap.PtrUint32(6112660, 2386524) = uint32(nox_xxx_getNameId_4E3AA0(CString("Ruby")))
		*memmap.PtrUint32(6112660, 2386532) = uint32(nox_xxx_getNameId_4E3AA0(CString("AnkhTradable")))
	}
	var v1 int32 = int32(*(*uint16)(unsafe.Pointer(uintptr(a1 + 4))))
	if uint32(uint16(int16(v1))) == *memmap.PtrUint32(6112660, 2386520) || uint32(v1) == *memmap.PtrUint32(6112660, 2386524) || uint32(v1) == *memmap.PtrUint32(6112660, 2386528) || uint32(v1) == *memmap.PtrUint32(6112660, 2386532) {
		return 0
	}
	return 1
}
func sub_5105D0(a1 int32) int32 {
	var v1 int32
	if *memmap.PtrUint32(6112660, 2386536) == 0 {
		*memmap.PtrUint32(6112660, 2386536) = uint32(nox_xxx_getNameId_4E3AA0(CString("Diamond")))
		*memmap.PtrUint32(6112660, 2386544) = uint32(nox_xxx_getNameId_4E3AA0(CString("Emerald")))
		*memmap.PtrUint32(6112660, 2386540) = uint32(nox_xxx_getNameId_4E3AA0(CString("Ruby")))
	}
	v1 = int32(*(*uint16)(unsafe.Pointer(uintptr(a1 + 4))))
	return bool2int(uint32(uint16(int16(v1))) == *memmap.PtrUint32(6112660, 2386536) || uint32(v1) == *memmap.PtrUint32(6112660, 2386540) || uint32(v1) == *memmap.PtrUint32(6112660, 2386544))
}
func nox_xxx_servShopStart_50EF10_trade(a1 int32, a2 int32) *uint32 {
	var (
		v2  int32
		v3  int32
		v4  *libc.WChar
		v6  int32
		v7  int32
		v8  *libc.WChar
		v9  *uint32
		v10 int8
		v11 int32
		v12 int32
		v13 int32
		v14 [128]libc.WChar
	)
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 8))) & 4)
	if v2 != 0 {
		v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))) + 280))))
		if v3 != 0 {
			if *(*uint32)(unsafe.Pointer(uintptr(v3 + 12))) != uint32(a2) {
				v4 = strMan.GetStringInFile("StarterAlreadyTrading", "C:\\NoxPost\\src\\Server\\System\\Trade.c")
				nox_xxx_netSendLineMessage_4D9EB0(a1, v4)
			}
			return nil
		}
	}
	v6 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 8))) & 4)
	if v6 != 0 {
		v7 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 748))))
		if *(*uint32)(unsafe.Pointer(uintptr(v7 + 280))) != 0 {
			if v2 != 0 {
				v13 = int32(*(*uint32)(unsafe.Pointer(uintptr(v7 + 276))) + 4704)
				v8 = strMan.GetStringInFile("OtherAlreadyTrading", "C:\\NoxPost\\src\\Server\\System\\Trade.c")
				nox_swprintf(&v14[0], v8, v13)
				nox_xxx_netSendLineMessage_4D9EB0(a1, &v14[0])
			}
			return nil
		}
	}
	if v2 != 0 {
		if v6 == 0 {
			v9 = nox_xxx_createPlayerShopSession_50E8F0(a1, a2)
			goto LABEL_18
		}
	} else if v6 == 0 {
		goto LABEL_17
	}
	if v2 == 0 {
		v9 = nox_xxx_createPlayerShopSession_50E8F0(a2, a1)
		goto LABEL_18
	}
LABEL_17:
	v9 = nox_xxx_createShopStruct_50E870()
	*(*uint32)(unsafe.Add(unsafe.Pointer(v9), unsafe.Sizeof(uint32(0))*2)) = uint32(a1)
	*(*uint32)(unsafe.Add(unsafe.Pointer(v9), unsafe.Sizeof(uint32(0))*3)) = uint32(a2)
LABEL_18:
	if v9 == nil {
		return nil
	}
	*(*uint32)(unsafe.Add(unsafe.Pointer(v9), unsafe.Sizeof(uint32(0))*1)) = nox_frame_xxx_2598000
	nox_xxx_tradeSetPlayer_50F370(v9, a1)
	nox_xxx_tradeSetPlayer_50F370(v9, a2)
	v10 = int8(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))
	if *(*uint32)(unsafe.Add(unsafe.Pointer(v9), unsafe.Sizeof(uint32(0))*4)) != 0 {
		if int32(v10)&4 != 0 {
			sub_50F0F0(a1, int32(uintptr(unsafe.Pointer(v9))))
		}
		if int32(*(*uint8)(unsafe.Pointer(uintptr(a2 + 8))))&4 != 0 {
			sub_50F0F0(a2, int32(uintptr(unsafe.Pointer(v9))))
		}
	} else {
		if int32(v10)&4 != 0 {
			sub_50F1A0(a1, int32(uintptr(unsafe.Pointer(v9))))
		}
		if int32(*(*uint8)(unsafe.Pointer(uintptr(a2 + 8))))&4 != 0 {
			sub_50F1A0(a2, int32(uintptr(unsafe.Pointer(v9))))
		}
	}
	if *(*uint32)(unsafe.Add(unsafe.Pointer(v9), unsafe.Sizeof(uint32(0))*4)) != 0 {
		v11 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v9), unsafe.Sizeof(uint32(0))*2)))
		if int32(*(*uint8)(unsafe.Pointer(uintptr(v11 + 8))))&4 != 0 {
			nox_xxx_servSendShopItems_50F280(v11, int32(uintptr(unsafe.Pointer(v9))))
		} else {
			nox_xxx_servSendShopItems_50F280(int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v9), unsafe.Sizeof(uint32(0))*3))), int32(uintptr(unsafe.Pointer(v9))))
		}
		if *(*uint32)(unsafe.Add(unsafe.Pointer(v9), unsafe.Sizeof(uint32(0))*4)) != 0 && noxflags.HasGame(noxflags.GameModeCoop) {
			v12 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v9), unsafe.Sizeof(uint32(0))*2)))
			if (int32(*(*uint8)(unsafe.Pointer(uintptr(v12 + 8)))) & 4) == 0 {
				v12 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v9), unsafe.Sizeof(uint32(0))*3)))
			}
			nox_xxx_unitFreeze_4E79C0((*nox_object_t)(unsafe.Pointer(uintptr(v12))), 0)
		}
	}
	return v9
}
func nox_xxx_tradeP2PAddOffer2_50F820_trade(a1 int32, a2 int32, a3 float32) int32 {
	var (
		v3     float32
		result int32
		v5     *int32
		v6     *int32
		v7     *libc.WChar
		v8     *libc.WChar
		v9     int32
		v10    int32
		v11    int32
		v12    int32
	)
	if *memmap.PtrUint32(6112660, 2386516) == 0 {
		*memmap.PtrUint32(6112660, 2386516) = uint32(nox_xxx_getNameId_4E3AA0(CString("Gold")))
	}
	if *(*uint32)(unsafe.Pointer(uintptr(a1 + 8))) == uint32(a2) {
		v3 = a3
		if sub_50FD60(*(**uint32)(unsafe.Pointer(uintptr(a1 + 32))), *(*int32)(unsafe.Add(unsafe.Pointer((*int32)(unsafe.Pointer(&a3))), unsafe.Sizeof(int32(0))*0))) == 0 {
			return 0
		}
	} else {
		if *(*uint32)(unsafe.Pointer(uintptr(a1 + 12))) != uint32(a2) {
			return 0
		}
		v3 = a3
		if sub_50FD60(*(**uint32)(unsafe.Pointer(uintptr(a1 + 36))), *(*int32)(unsafe.Add(unsafe.Pointer((*int32)(unsafe.Pointer(&a3))), unsafe.Sizeof(int32(0))*0))) == 0 {
			return 0
		}
	}
	v5 = (*int32)(nox_alloc_class_new_obj_zero((*nox_alloc_class)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&nox_alloc_tradeItems_2386496))))))
	v6 = v5
	if v5 != nil {
		*(*float32)(unsafe.Pointer(v5)) = v3
		*(*int32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(int32(0))*2)) = 0
		*(*int32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(int32(0))*3)) = 0
		*(*int32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(int32(0))*1)) = nox_xxx_shopGetItemCost_50E3D0(1, a1, v3)
		if *(*uint32)(unsafe.Pointer(uintptr(a1 + 8))) == uint32(a2) {
			*(*int32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(int32(0))*2)) = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 32))))
			v9 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 32))))
			if v9 != 0 {
				*(*uint32)(unsafe.Pointer(uintptr(v9 + 12))) = uint32(uintptr(unsafe.Pointer(v6)))
			}
			*(*uint32)(unsafe.Pointer(uintptr(a1 + 32))) = uint32(uintptr(unsafe.Pointer(v6)))
		} else {
			*(*int32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(int32(0))*2)) = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 36))))
			v10 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 36))))
			if v10 != 0 {
				*(*uint32)(unsafe.Pointer(uintptr(v10 + 12))) = uint32(uintptr(unsafe.Pointer(v6)))
			}
			*(*uint32)(unsafe.Pointer(uintptr(a1 + 36))) = uint32(uintptr(unsafe.Pointer(v6)))
		}
		sub_50FB90((*uint32)(unsafe.Pointer(uintptr(a1))))
		v11 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 8))))
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 24))) = 0
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 28))) = 0
		if int32(*(*uint8)(unsafe.Pointer(uintptr(v11 + 8))))&4 != 0 {
			if (int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 12))) + 8))))&4) == 0 && *(*uint32)(unsafe.Pointer(uintptr(a1 + 44))) <= uint32(*(*int32)(unsafe.Pointer(uintptr(a1 + 40)))) {
				*(*uint32)(unsafe.Pointer(uintptr(a1 + 28))) = 1
			}
		} else if *(*uint32)(unsafe.Pointer(uintptr(a1 + 40))) <= uint32(*(*int32)(unsafe.Pointer(uintptr(a1 + 44)))) {
			*(*uint32)(unsafe.Pointer(uintptr(a1 + 24))) = 1
		}
		if int32(*(*uint8)(unsafe.Pointer(uintptr(v11 + 8))))&4 != 0 {
			nox_xxx_tradeP2PUpdStuff_50FA00(v11, (*uint32)(unsafe.Pointer(uintptr(a1))))
			sub_50FAE0(int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 8)))), a2, a1, *v6, *(*int32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(int32(0))*1)))
			sub_50F720(int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 8)))), (*uint32)(unsafe.Pointer(uintptr(a1))))
		}
		v12 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 12))))
		if int32(*(*uint8)(unsafe.Pointer(uintptr(v12 + 8))))&4 != 0 {
			nox_xxx_tradeP2PUpdStuff_50FA00(v12, (*uint32)(unsafe.Pointer(uintptr(a1))))
			sub_50FAE0(int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 12)))), a2, a1, *v6, *(*int32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(int32(0))*1)))
			sub_50F720(int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 12)))), (*uint32)(unsafe.Pointer(uintptr(a1))))
		}
		result = 1
	} else {
		if int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 8))) + 8))))&4 != 0 {
			v7 = strMan.GetStringInFile("TradeMaxObjectsReached", "C:\\NoxPost\\src\\Server\\System\\Trade.c")
			nox_xxx_netSendLineMessage_4D9EB0(int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 8)))), v7)
		}
		if (int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 12))) + 8)))) & 4) == 0 {
			return 0
		}
		v8 = strMan.GetStringInFile("TradeMaxObjectsReached", "C:\\NoxPost\\src\\Server\\System\\Trade.c")
		nox_xxx_netSendLineMessage_4D9EB0(int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 12)))), v8)
		result = 0
	}
	return result
}
func sub_5100C0_trade(a1 int32, a2 *uint32, a3 int32) *float32 {
	var (
		v3     int32
		v4     uint32
		result *float32
		v6     int32
		v7     uint32
		v8     uint32
		v9     int32
		v10    *libc.WChar
		v11    int32
		v12    int32
		v13    *uint32
		v14    func(int32, *uint32, int32, int32)
		v15    int32
		v16    float32
		v17    float32
		v18    int32
	)
	v3 = a1
	v15 = a1
	v18 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
	v4 = uint32(nox_xxx_playerGetGold_4FA6B0(v15))
	if dword_5d4594_2386548 == 0 {
		dword_5d4594_2386548 = uint32(nox_xxx_getNameId_4E3AA0(CString("AnkhTradable")))
	}
	result = (*float32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*5)))))
	if result != nil {
		for *(*uint32)(unsafe.Pointer(result)) == 0 || *(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(result)) + 36))) != uint32(a3) {
			result = (*float32)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(result))), unsafe.Sizeof(uint32(0))*2))))))
			if result == nil {
				return result
			}
		}
		v6 = int32(*(*uint32)(unsafe.Pointer(result)))
		if *(*uint32)(unsafe.Pointer(result)) != 0 {
			v7 = uint32(nox_xxx_shopGetItemCost_50E3D0(1, int32(uintptr(unsafe.Pointer(a2))), *result))
			v8 = v7
			if v7 > v4 {
				return (*float32)(unsafe.Pointer(uintptr(sub_5104F0(v3, int16(uint16(v7-v4))))))
			}
			if int32(*(*uint8)(unsafe.Pointer(uintptr(v6 + 8))))&16 != 0 {
				v9 = nox_xxx_inventoryCountObjects_4E7D30(v3, int32(*(*uint16)(unsafe.Pointer(uintptr(v6 + 4)))))
				if v9 >= (func() int32 {
					if noxflags.HasGame(noxflags.GameModeCoop | noxflags.GameModeQuest) {
						return 9
					}
					return 3
				}()) {
					v10 = strMan.GetStringInFile("pickup.c:MaxSameItem", "C:\\NoxPost\\src\\Server\\System\\Trade.c")
					return (*float32)(unsafe.Pointer(uintptr(nox_xxx_netSendLineMessage_4D9EB0(v3, v10))))
				}
			}
			if uint32(*(*uint16)(unsafe.Pointer(uintptr(v6 + 4)))) == dword_5d4594_2386548 {
				v16 = float32(nox_xxx_gamedataGetFloat_419D40(CString("MaxExtraLives")))
				if *(*uint32)(unsafe.Pointer(uintptr(v18 + 320))) >= uint32(int(v16)) {
					nox_xxx_netPriMsgToPlayer_4DA2C0((*nox_object_t)(unsafe.Pointer(uintptr(v3))), CString("pickup.c:MaxTradableAnkhsReached"), 0)
					return (*float32)(unsafe.Pointer(nox_xxx_aud_501960(925, (*nox_object_t)(unsafe.Pointer(uintptr(v3))), 0, 0)))
				}
			}
			if noxflags.HasGame(noxflags.GameModeQuest) {
				v11 = int32(*(*uint32)(unsafe.Pointer(uintptr(v6 + 8))))
				if v11&4096 != 0 {
					if *(*uint32)(unsafe.Pointer(uintptr(v6 + 12)))&0x200000 != 0 {
						v17 = float32(nox_xxx_gamedataGetFloat_419D40(CString("ForceOfNatureStaffLimit")))
						v12 = int(v17)
						if nox_xxx_inventoryCountObjects_4E7D30(v3, int32(*(*uint16)(unsafe.Pointer(uintptr(v6 + 4))))) >= v12 {
							nox_xxx_netPriMsgToPlayer_4DA2C0((*nox_object_t)(unsafe.Pointer(uintptr(v3))), CString("pickup.c:MaxSameItem"), 0)
							return (*float32)(unsafe.Pointer(nox_xxx_aud_501960(925, (*nox_object_t)(unsafe.Pointer(uintptr(v3))), 0, 0)))
						}
					}
				}
			}
			if noxflags.HasGame(noxflags.GameModeQuest) && (sub_5105D0(v6) != 0 || uint32(*(*uint16)(unsafe.Pointer(uintptr(v6 + 4)))) == dword_5d4594_2386548) {
				v13 = (*uint32)(unsafe.Pointer(nox_xxx_newObjectWithTypeInd_4E3450(int32(*(*uint16)(unsafe.Pointer(uintptr(v6 + 4)))))))
			} else {
				v13 = (*uint32)(unsafe.Pointer(uintptr(v6)))
			}
			if *(*uint32)(unsafe.Add(unsafe.Pointer(v13), unsafe.Sizeof(uint32(0))*2))&272 != 0 || (func() func(int32, *uint32, int32, int32) {
				v14 = cgoAsFunc(*(*uint32)(unsafe.Add(unsafe.Pointer(v13), unsafe.Sizeof(uint32(0))*177)), (*func(int32, *uint32, int32, int32))(nil)).(func(int32, *uint32, int32, int32))
				return v14
			}()) == nil {
				nox_xxx_inventoryPutImpl_4F3070(v3, (*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v13)))))), 1)
				nox_xxx_aud_501960(307, (*nox_object_t)(unsafe.Pointer(uintptr(v3))), 2, int32(*(*uint32)(unsafe.Pointer(uintptr(v3 + 36)))))
			} else {
				v14(v3, v13, 1, 1)
			}
			sub_510320(v6, int32(uintptr(unsafe.Pointer(a2))))
			if sub_510540(v6) != 0 {
				sub_50E7A0(a2, v6)
			}
			nox_xxx_playerSubGold_4FA5D0(v3, v8)
			result = (*float32)(unsafe.Pointer(uintptr(sub_4D8870(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v18 + 276))) + 2064)))), v3))))
		}
	}
	return result
}
func sub_510640_trade(a1 int32, a2 int32, a3 int32, a4 *float32) *float32 {
	var (
		v4     int32
		result *float32
		v6     *uint32
		v7     int32
		v8     uint32
		v9     int32
		v10    int32
		v11    int32
		v12    *uint32
		v13    func(int32, *uint32, int32, int32)
		v14    *libc.WChar
		v15    int32
		v16    float32
		v17    float32
		v18    uint32
		v19    uint32
		v20    int32
	)
	v4 = a1
	v15 = a1
	v20 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
	v19 = uint32(nox_xxx_playerGetGold_4FA6B0(v15))
	if dword_5d4594_2386552 == 0 {
		dword_5d4594_2386552 = uint32(nox_xxx_getNameId_4E3AA0(CString("AnkhTradable")))
	}
	result = a4
	v18 = 0
	if a4 != nil {
		for {
			v6 = (*uint32)(unsafe.Pointer(uintptr(a2)))
			result = *(**float32)(unsafe.Pointer(uintptr(a2 + 20)))
			if result == nil {
				return result
			}
			for int32(*(*uint16)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(result)) + 4)))) != a3 {
				result = (*float32)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(result))), unsafe.Sizeof(uint32(0))*2))))))
				if result == nil {
					return result
				}
			}
			v7 = int32(*(*uint32)(unsafe.Pointer(result)))
			if *(*uint32)(unsafe.Pointer(result)) == 0 {
				return result
			}
			v8 = uint32(nox_xxx_shopGetItemCost_50E3D0(1, a2, *result))
			if v8 > v19 {
				return (*float32)(unsafe.Pointer(uintptr(sub_5104F0(v4, int16(uint16(v8-v19))))))
			}
			if int32(*(*uint8)(unsafe.Pointer(uintptr(v7 + 8))))&16 != 0 {
				v9 = nox_xxx_inventoryCountObjects_4E7D30(v4, int32(*(*uint16)(unsafe.Pointer(uintptr(v7 + 4)))))
				if v9 >= (func() int32 {
					if noxflags.HasGame(noxflags.GameModeCoop | noxflags.GameModeQuest) {
						return 9
					}
					return 3
				}()) {
					v14 = strMan.GetStringInFile("pickup.c:MaxSameItem", "C:\\NoxPost\\src\\Server\\System\\Trade.c")
					return (*float32)(unsafe.Pointer(uintptr(nox_xxx_netSendLineMessage_4D9EB0(v4, v14))))
				}
				v6 = (*uint32)(unsafe.Pointer(uintptr(a2)))
			}
			if uint32(*(*uint16)(unsafe.Pointer(uintptr(v7 + 4)))) == dword_5d4594_2386552 {
				v16 = float32(nox_xxx_gamedataGetFloat_419D40(CString("MaxExtraLives")))
				if *(*uint32)(unsafe.Pointer(uintptr(v20 + 320))) >= uint32(int(v16)) {
					break
				}
			}
			if noxflags.HasGame(noxflags.GameModeQuest) {
				v10 = int32(*(*uint32)(unsafe.Pointer(uintptr(v7 + 8))))
				if v10&4096 != 0 {
					if *(*uint32)(unsafe.Pointer(uintptr(v7 + 12)))&0x200000 != 0 {
						v17 = float32(nox_xxx_gamedataGetFloat_419D40(CString("ForceOfNatureStaffLimit")))
						v11 = int(v17)
						if nox_xxx_inventoryCountObjects_4E7D30(v4, int32(*(*uint16)(unsafe.Pointer(uintptr(v7 + 4))))) >= v11 {
							nox_xxx_netPriMsgToPlayer_4DA2C0((*nox_object_t)(unsafe.Pointer(uintptr(v4))), CString("pickup.c:MaxSameItem"), 0)
							goto LABEL_36
						}
						v6 = (*uint32)(unsafe.Pointer(uintptr(a2)))
					}
				}
			}
			if noxflags.HasGame(noxflags.GameModeQuest) && (sub_5105D0(v7) != 0 || uint32(*(*uint16)(unsafe.Pointer(uintptr(v7 + 4)))) == dword_5d4594_2386552) {
				v12 = (*uint32)(unsafe.Pointer(nox_xxx_newObjectWithTypeInd_4E3450(int32(*(*uint16)(unsafe.Pointer(uintptr(v7 + 4)))))))
			} else {
				v12 = (*uint32)(unsafe.Pointer(uintptr(v7)))
			}
			v13 = cgoAsFunc(*(*uint32)(unsafe.Add(unsafe.Pointer(v12), unsafe.Sizeof(uint32(0))*177)), (*func(int32, *uint32, int32, int32))(nil)).(func(int32, *uint32, int32, int32))
			if v13 != nil {
				v13(v4, v12, 1, 1)
			} else {
				nox_xxx_inventoryPutImpl_4F3070(v4, (*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v12)))))), 1)
			}
			sub_510320(v7, int32(uintptr(unsafe.Pointer(v6))))
			if sub_510540(v7) != 0 {
				sub_50E7A0(v6, v7)
			}
			nox_xxx_playerSubGold_4FA5D0(v4, v8)
			sub_4D8870(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v20 + 276))) + 2064)))), v4)
			result = (*float32)(unsafe.Pointer(uintptr(func() uint32 {
				p := &v18
				*p++
				return *p
			}())))
			if v18 >= uint32(uintptr(unsafe.Pointer(a4))) {
				return result
			}
		}
		nox_xxx_netPriMsgToPlayer_4DA2C0((*nox_object_t)(unsafe.Pointer(uintptr(v4))), CString("pickup.c:MaxTradableAnkhsReached"), 0)
	LABEL_36:
		result = (*float32)(unsafe.Pointer(nox_xxx_aud_501960(925, (*nox_object_t)(unsafe.Pointer(uintptr(v4))), 0, 0)))
	}
	return result
}
func sub_5109C0_trade(a1 *int32, a2 int32, a3 *uint32) *uint32 {
	var (
		v3     int32
		v4     int32
		result *uint32
		v6     *libc.WChar
		v7     *libc.WChar
		v8     [8]byte
	)
	v3 = *(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*187))
	if *memmap.PtrUint32(6112660, 2386556) == 0 {
		*memmap.PtrUint32(6112660, 2386556) = uint32(nox_xxx_getNameId_4E3AA0(CString("Glyph")))
	}
	v4 = *(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*126))
	result = a3
	v8[0] = 201
	v8[1] = 29
	*(*uint16)(unsafe.Pointer(&v8[2])) = uint16(uintptr(unsafe.Pointer(a3)))
	if v4 != 0 {
		for *(**uint32)(unsafe.Pointer(uintptr(v4 + 36))) != a3 {
			v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(v4 + 496))))
			if v4 == 0 {
				return result
			}
		}
		if nox_xxx_ItemIsDroppable_53EBF0(v4) == 1 {
			v6 = strMan.GetStringInFile("CantSellQuestItem", "C:\\NoxPost\\src\\Server\\System\\Trade.c")
			nox_xxx_netSendLineMessage_4D9EB0(int32(uintptr(unsafe.Pointer(a1))), v6)
			result = nox_xxx_aud_501960(925, (*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1)))))), 2, *(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*9)))
		} else if uint32(*(*uint16)(unsafe.Pointer(uintptr(v4 + 4)))) == *memmap.PtrUint32(6112660, 2386556) {
			v7 = strMan.GetStringInFile("CantSellItem", "C:\\NoxPost\\src\\Server\\System\\Trade.c")
			nox_xxx_netSendLineMessage_4D9EB0(int32(uintptr(unsafe.Pointer(a1))), v7)
			result = nox_xxx_aud_501960(925, (*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1)))))), 2, *(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*9)))
		} else {
			*(*uint32)(unsafe.Pointer(&v8[4])) = uint32(nox_xxx_shopGetItemCost_50E3D0(0, a2, *(*float32)(unsafe.Pointer(&v4))))
			result = (*uint32)(unsafe.Pointer(uintptr(nox_xxx_netSendPacket0_4E5420(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v3 + 276))) + 2064)))), unsafe.Pointer(&v8[0]), 8, 0, 1))))
		}
	}
	return result
}
func sub_510BE0_trade(a1 *int32, a2 int32, a3 *uint32) *uint32 {
	var (
		result *uint32
		v4     int32
		v5     int32
		v6     *libc.WChar
		v7     *libc.WChar
		v8     int32
	)
	nox_xxx_playerGetGold_4FA6B0(int32(uintptr(unsafe.Pointer(a1))))
	result = *(**uint32)(unsafe.Pointer(&dword_5d4594_2386560))
	v4 = *(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*187))
	if dword_5d4594_2386560 == 0 {
		result = (*uint32)(unsafe.Pointer(uintptr(nox_xxx_getNameId_4E3AA0(CString("Glyph")))))
		dword_5d4594_2386560 = uint32(uintptr(unsafe.Pointer(result)))
	}
	v5 = *(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*126))
	if v5 != 0 {
		result = a3
		for *(**uint32)(unsafe.Pointer(uintptr(v5 + 36))) != a3 {
			v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(v5 + 496))))
			if v5 == 0 {
				return result
			}
		}
		if nox_xxx_ItemIsDroppable_53EBF0(v5) == 1 {
			v6 = strMan.GetStringInFile("CantSellQuestItem", "C:\\NoxPost\\src\\Server\\System\\Trade.c")
			nox_xxx_netSendLineMessage_4D9EB0(int32(uintptr(unsafe.Pointer(a1))), v6)
			result = nox_xxx_aud_501960(925, (*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1)))))), 2, *(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*9)))
		} else if uint32(*(*uint16)(unsafe.Pointer(uintptr(v5 + 4)))) == dword_5d4594_2386560 {
			v7 = strMan.GetStringInFile("CantSellItem", "C:\\NoxPost\\src\\Server\\System\\Trade.c")
			nox_xxx_netSendLineMessage_4D9EB0(int32(uintptr(unsafe.Pointer(a1))), v7)
			result = nox_xxx_aud_501960(925, (*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1)))))), 2, *(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*9)))
		} else {
			sub_4ED0C0(int32(uintptr(unsafe.Pointer(a1))), (*nox_object_t)(unsafe.Pointer(uintptr(v5))))
			nox_xxx_delayedDeleteObject_4E5CC0((*nox_object_t)(unsafe.Pointer(uintptr(v5))))
			v8 = nox_xxx_shopGetItemCost_50E3D0(0, a2, *(*float32)(unsafe.Pointer(&v5)))
			nox_xxx_playerAddGold_4FA590(int32(uintptr(unsafe.Pointer(a1))), v8)
			sub_4D8870(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v4 + 276))) + 2064)))), int32(uintptr(unsafe.Pointer(a1))))
			result = nox_xxx_aud_501960(307, (*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1)))))), 2, *(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*9)))
		}
	}
	return result
}
