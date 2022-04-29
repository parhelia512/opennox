package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	"github.com/gotranspile/cxgo/runtime/stdio"
	"github.com/noxworld-dev/opennox/v1/common/alloc"
	noxflags "github.com/noxworld-dev/opennox/v1/common/flags"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"image"
	"math"
	"maze.io/x/math32"
	"unicode"
	"unsafe"
)

var nox_wnd_xxx_1307732 *nox_gui_animation = nil
var nox_wnd_xxx_1308092 *nox_gui_animation = nil
var nox_wnd_xxx_1309740 *nox_gui_animation = nil
var dword_5d4594_3798632 unsafe.Pointer = nil
var dword_5d4594_3798644 *byte = nil
var dword_5d4594_1307292 unsafe.Pointer = nil

func sub_4A1A40(a1 int32) int32 {
	var v1 *uint32
	v1 = (*uint32)(unsafe.Pointer((*nox_window)(dword_5d4594_1307292).ChildByID(151)))
	return nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v1)))))), a1)
}
func sub_4A2560(a1 *uint32, a2 int32) int32 {
	var (
		v2 float64
		v3 float64
	)
	v2 = float64(uint32(*(*int16)(unsafe.Pointer(uintptr(a2 + 44)))) - *a1)
	v3 = float64(uint32(*(*int16)(unsafe.Pointer(uintptr(a2 + 46)))) - *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*1)))
	return bool2int(math.Sqrt(v3*v3+v2*v2) <= *mem_getDoublePtr(0x581450, 9720))
}
func sub_4A25C0(a1 *uint32, a2 *int32) int32 {
	var (
		v2 int32
		v3 *int32
	)
	v2 = 0
	v3 = (*int32)(unsafe.Pointer(nox_common_list_getFirstSafe_425890((*nox_list_item_t)(unsafe.Pointer(a2)))))
	if v3 == nil {
		return 0
	}
	for {
		if sub_4A2560(a1, int32(uintptr(unsafe.Pointer(v3)))) != 0 {
			v2++
		}
		v3 = (*int32)(unsafe.Pointer(nox_common_list_getNextSafe_4258A0((*nox_list_item_t)(unsafe.Pointer(v3)))))
		if v3 == nil {
			break
		}
	}
	return v2
}
func sub_4A2610(a1 int32, a2 *uint32, a3 *int32) int32 {
	var (
		i   *int32
		v4  int32
		v5  *uint32
		v6  *uint32
		v7  *uint32
		v8  *uint32
		v9  *byte
		v10 int32
		v11 *uint8
		v13 *uint32
		v14 *byte
		v15 [2]int32
		v16 [64]byte
		v17 [128]libc.WChar
	)
	dword_5d4594_1307720 = 0
	for i = (*int32)(unsafe.Pointer(nox_common_list_getFirstSafe_425890((*nox_list_item_t)(unsafe.Pointer(a3))))); i != nil; i = (*int32)(unsafe.Pointer(nox_common_list_getNextSafe_4258A0((*nox_list_item_t)(unsafe.Pointer(i))))) {
		if sub_4A2560(a2, int32(uintptr(unsafe.Pointer(i)))) != 0 {
			v4 = int32(dword_5d4594_1307720)
			*memmap.PtrUint32(6112660, dword_5d4594_1307720*4+0x13F2B4) = uint32(uintptr(unsafe.Pointer(i)))
			dword_5d4594_1307720 = uint32(v4 + 1)
		}
	}
	if dword_5d4594_1307720 > 0 {
		dword_5d4594_1307716 = uint32(uintptr(unsafe.Pointer(newWindowFromFile("proxlist.wnd", unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 376)))))))))
		sub_4A2830(int32(*a2+216), int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*1))+27), (*uint32)(unsafe.Pointer(&v15[0])))
		(*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1307716)))).SetPos(image.Pt(v15[0], v15[1]))
		nox_xxx_wnd_46B280(*(*int32)(unsafe.Pointer(&dword_5d4594_1307716)), a1)
		v5 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1307716)))).ChildByID(0x2750)))
		v6 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1307716)))).ChildByID(0x274E)))
		v13 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1307716)))).ChildByID(0x274F)))
		v7 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1307716)))).ChildByID(0x274D)))
		v8 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v7), unsafe.Sizeof(uint32(0))*8)))))
		v14 = (*byte)(unsafe.Pointer(nox_xxx_gLoadImg("UISlider")))
		v9 = (*byte)(unsafe.Pointer(nox_xxx_gLoadImg("UISliderLit")))
		sub_4B5700(int32(uintptr(unsafe.Pointer(v5))), 0, 0, int32(uintptr(unsafe.Pointer(v14))), int32(uintptr(unsafe.Pointer(v9))), int32(uintptr(unsafe.Pointer(v9))))
		nox_xxx_wnd_46B280(int32(uintptr(unsafe.Pointer(v5))), int32(uintptr(unsafe.Pointer(v7))))
		nox_xxx_wnd_46B280(int32(uintptr(unsafe.Pointer(v6))), int32(uintptr(unsafe.Pointer(v7))))
		nox_xxx_wnd_46B280(int32(uintptr(unsafe.Pointer(v13))), int32(uintptr(unsafe.Pointer(v7))))
		*(*uint32)(unsafe.Add(unsafe.Pointer(v8), unsafe.Sizeof(uint32(0))*9)) = uint32(uintptr(unsafe.Pointer(v5)))
		*(*uint32)(unsafe.Add(unsafe.Pointer(v8), unsafe.Sizeof(uint32(0))*7)) = uint32(uintptr(unsafe.Pointer(v6)))
		*(*uint32)(unsafe.Add(unsafe.Pointer(v8), unsafe.Sizeof(uint32(0))*8)) = uint32(uintptr(unsafe.Pointer(v13)))
		*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(uint32(0))*100)) + 8))) = 16
		*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(uint32(0))*100)) + 12))) = 10
		v10 = 0
		if dword_5d4594_1307720 > 0 {
			v11 = (*uint8)(memmap.PtrOff(6112660, 1307316))
			for {
				if int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(v11)) + 120)))) != 0 {
					libc.StrNCpy(&v16[0], (*byte)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(v11))+120))), 15)
					v16[15] = 0
				} else {
					nox_sprintAddrPort_43BC80((*byte)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(v11))+12))), *(*uint16)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(v11)) + 109))), &v16[0])
				}
				nox_swprintf(&v17[0], libc.CWString("%S   %dms"), &v16[0], *(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(v11)) + 96))))
				(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v7)))))).Func94(asWindowEvent(0x400D, int32(uintptr(unsafe.Pointer(&v17[0]))), -1))
				v10++
				v11 = (*uint8)(unsafe.Add(unsafe.Pointer(v11), 4))
				if v10 >= *(*int32)(unsafe.Pointer(&dword_5d4594_1307720)) {
					break
				}
			}
		}
	}
	return int32(dword_5d4594_1307716)
}
func sub_4A2830(a1 int32, a2 int32, a3 *uint32) *uint32 {
	var result *uint32
	result = a3
	*a3 = uint32(a1 - 100)
	*(*uint32)(unsafe.Add(unsafe.Pointer(a3), unsafe.Sizeof(uint32(0))*1)) = uint32(a2 - 20)
	if a1-100+200 > 600 {
		*a3 = 400
	}
	if a2-20+200 > 451 {
		*(*uint32)(unsafe.Add(unsafe.Pointer(a3), unsafe.Sizeof(uint32(0))*1)) = 251
	}
	if *(*uint32)(unsafe.Add(unsafe.Pointer(a3), unsafe.Sizeof(uint32(0))*1)) < 27 {
		*(*uint32)(unsafe.Add(unsafe.Pointer(a3), unsafe.Sizeof(uint32(0))*1)) = 27
	}
	if int32(*a3) < 216 {
		*a3 = 216
	}
	return result
}
func sub_4A2890() int32 {
	var result int32
	result = int32(dword_5d4594_1307716)
	if dword_5d4594_1307716 != 0 {
		result = (*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1307716)))).Destroy()
		dword_5d4594_1307716 = 0
	}
	return result
}
func sub_4A28B0() int32 {
	return bool2int(dword_5d4594_1307716 != 0)
}
func sub_4A28C0(a1 int32) int32 {
	var result int32
	if a1 < *(*int32)(unsafe.Pointer(&dword_5d4594_1307720)) {
		result = int32(*memmap.PtrUint32(6112660, uint32(a1*4)+0x13F2B4))
	} else {
		result = 0
	}
	return result
}
func nox_xxx_wndListboxProcWithoutData10_4A28E0(a1 *uint32, a2 int32, a3 uint32, a4 int32) int32 {
	var (
		v4     *uint32
		v5     int32
		result int32
		v7     int32
		v8     int32
		v9     int32
		v10    int32
		v11    int32
		v12    uint32
		v13    uint32
		v14    int32
		v15    int32
		v16    int32
		v17    uint32
		v18    uint32
		v19    int32
		v20    int32
		v21    int32
		v22    int32
	)
	v4 = a1
	v5 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*8)))
	switch a2 {
	case 5:
		fallthrough
	case 17:
		fallthrough
	case 18:
		return 1
	case 6:
		fallthrough
	case 7:
		v12 = a3
		a3 = *(*uint32)(unsafe.Pointer(uintptr(v5 + 48)))
		v13 = v12 >> 16
		nox_client_wndGetPosition_46AA60((*nox_window)(unsafe.Pointer(a1)), (*uint32)(unsafe.Pointer(&a1)), (*uint32)(unsafe.Pointer(&a4)))
		if int32(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v4))), unsafe.Sizeof(uint16(0))*54)))) != 0 {
			v14 = nox_xxx_guiFontHeightMB_43F320(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*59)))))
			a4 += v14 + 1
		}
		v15 = 0
		*(*uint32)(unsafe.Pointer(uintptr(v5 + 48))) = 0xFFFFFFFE
		v16 = 0
		for 2 != 0 {
			if v16 > 0 && *(*int32)(unsafe.Pointer(uintptr(uint32(v16) + *(*uint32)(unsafe.Pointer(uintptr(v5 + 24))) - 524))) > int32(*(*int16)(unsafe.Pointer(uintptr(v5 + 52))))+int32(*(*int16)(unsafe.Pointer(uintptr(v5 + 54)))) || v15 == int32(*(*int16)(unsafe.Pointer(uintptr(v5 + 44)))) {
				*(*uint32)(unsafe.Pointer(uintptr(v5 + 48))) = math.MaxUint32
			} else if *(*int32)(unsafe.Pointer(uintptr(uint32(v16) + *(*uint32)(unsafe.Pointer(uintptr(v5 + 24)))))) <= int32(v13+uint32(*(*int16)(unsafe.Pointer(uintptr(v5 + 54))))-uint32(a4)) {
				v15++
				v16 += 524
				continue
			}
			break
		}
		v17 = a3
		if uint32(v15) == a3 && *(*uint32)(unsafe.Pointer(uintptr(v5 + 20))) == 0 {
			*(*uint32)(unsafe.Pointer(uintptr(v5 + 48))) = math.MaxUint32
		}
		if *(*int32)(unsafe.Pointer(uintptr(v5 + 48))) == -2 && v15 < int32(*(*int16)(unsafe.Pointer(uintptr(v5 + 44)))) {
			*(*uint32)(unsafe.Pointer(uintptr(v5 + 48))) = uint32(v15)
		}
		if *(*int32)(unsafe.Pointer(uintptr(v5 + 48))) < 0 {
			if *(*uint32)(unsafe.Pointer(uintptr(v5 + 20))) != 0 {
				*(*uint32)(unsafe.Pointer(uintptr(v5 + 48))) = v17
			}
		}
		(*nox_window)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*13))))).Func94(asWindowEvent(0x4010, int32(uintptr(unsafe.Pointer(v4))), int32(*(*uint32)(unsafe.Pointer(uintptr(v5 + 48))))))
		return 1
	case 8:
		v22 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*11)))
		if v22&256 != 0 {
			(*nox_window)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*13))))).Func94(asWindowEvent(0x4000, int32(uintptr(unsafe.Pointer(a1))), 0))
		}
		return 1
	case 10:
		fallthrough
	case 11:
		v18 = a3 >> 16
		nox_client_wndGetPosition_46AA60((*nox_window)(unsafe.Pointer(a1)), &a3, (*uint32)(unsafe.Pointer(&a4)))
		if int32(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v4))), unsafe.Sizeof(uint16(0))*54)))) != 0 {
			v19 = nox_xxx_guiFontHeightMB_43F320(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*59)))))
			a4 += v19 + 1
		}
		v20 = 0
		v21 = 0
	case 19:
	LABEL_9:
		v8 = int32(*(*uint32)(unsafe.Pointer(uintptr(v5 + 48))))
		if v8 == -1 {
			*(*uint32)(unsafe.Pointer(uintptr(v5 + 48))) = 0
			nox_xxx_wndListBox_4A2D10(int32(uintptr(unsafe.Pointer(v4))), 0, 1)
		} else if v8 > 0 {
			v11 = v8 - 1
			*(*uint32)(unsafe.Pointer(uintptr(v5 + 48))) = uint32(v11)
			if *(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v5 + 24))) + uint32(v11*524)))) < uint32(*(*int16)(unsafe.Pointer(uintptr(v5 + 54)))) {
				nox_xxx_wndListBox_4A2D10(int32(uintptr(unsafe.Pointer(v4))), -1, 1)
			}
		}
		goto LABEL_21
	case 20:
	LABEL_6:
		v7 = int32(*(*uint32)(unsafe.Pointer(uintptr(v5 + 48))))
		if v7 == -1 {
			*(*uint32)(unsafe.Pointer(uintptr(v5 + 48))) = 0
			nox_xxx_wndListBox_4A2D10(int32(uintptr(unsafe.Pointer(v4))), 0, 1)
		} else if v7 < int32(*(*int16)(unsafe.Pointer(uintptr(v5 + 44))))-1 {
			v9 = v7 + 1
			v10 = int32(*(*int16)(unsafe.Pointer(uintptr(v5 + 52)))) + int32(*(*int16)(unsafe.Pointer(uintptr(v5 + 54))))
			*(*uint32)(unsafe.Pointer(uintptr(v5 + 48))) = uint32(v9)
			if *(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v5 + 24))) + uint32(v9*524)))) > uint32(v10) {
				nox_xxx_wndListBox_4A2D10(int32(uintptr(unsafe.Pointer(v4))), 1, 1)
			}
		}
	LABEL_21:
		(*nox_window)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*13))))).Func94(asWindowEvent(0x4010, int32(uintptr(unsafe.Pointer(v4))), int32(*(*uint32)(unsafe.Pointer(uintptr(v5 + 48))))))
		return 1
	case 21:
		switch a3 {
		case 15:
			fallthrough
		case 205:
			if a4 != 2 {
				return 1
			}
			nox_xxx_wndRetNULL_46A8A0()
			result = 1
		case 28:
			fallthrough
		case 57:
			if a4 != 1 {
				return 1
			}
			(*nox_window)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*13))))).Func94(asWindowEvent(0x4010, int32(uintptr(unsafe.Pointer(a1))), int32(*(*uint32)(unsafe.Pointer(uintptr(v5 + 48))))))
			result = 1
		case 200:
			if a4 == 2 {
				goto LABEL_9
			}
			return 1
		case 203:
			if a4 != 2 {
				return 1
			}
			nox_xxx_wndRetNULL_0_46A8B0()
			result = 1
		case 208:
			if a4 == 2 {
				goto LABEL_6
			}
			return 1
		default:
			return 0
		}
		return result
	default:
		return 0
	}
	for (v21 <= 0 || *(*uint32)(unsafe.Pointer(uintptr(uint32(v21) + *(*uint32)(unsafe.Pointer(uintptr(v5 + 24))) - 524))) <= uint32(int32(*(*int16)(unsafe.Pointer(uintptr(v5 + 52))))+int32(*(*int16)(unsafe.Pointer(uintptr(v5 + 54)))))) && v20 != int32(*(*int16)(unsafe.Pointer(uintptr(v5 + 44)))) {
		if *(*uint32)(unsafe.Pointer(uintptr(uint32(v21) + *(*uint32)(unsafe.Pointer(uintptr(v5 + 24)))))) > uint32(int32(v18+uint32(*(*int16)(unsafe.Pointer(uintptr(v5 + 54))))-uint32(a4))) {
			goto LABEL_50
		}
		v20++
		v21 += 524
	}
	v20 = -1
LABEL_50:
	(*nox_window)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*13))))).Func94(asWindowEvent(0x4011, int32(uintptr(unsafe.Pointer(v4))), v20))
	return 1
}
func nox_xxx_wndListBox_4A2D10(a1 int32, a2 int32, a3 int32) int32 {
	var (
		v3     int32
		v4     int32
		v5     int32
		result int32
		v7     int32
	)
	v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 32))))
	v4 = a2 + sub_4A4800(int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 32)))))
	if v4 >= 0 {
		v5 = int32(*(*int16)(unsafe.Pointer(uintptr(v3 + 44))))
		if v4 >= v5 {
			v4 = v5 - 1
		}
	} else {
		v4 = 0
	}
	if a3 != 0 {
		if v4 <= 0 {
			*(*uint16)(unsafe.Pointer(uintptr(v3 + 54))) = 0
		} else {
			*(*uint16)(unsafe.Pointer(uintptr(v3 + 54))) = uint16(int16(int32(*(*uint16)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v3 + 24))) + uint32(v4*524) - 524)))) + 1))
		}
	}
	result = int32(*(*uint32)(unsafe.Pointer(uintptr(v3 + 36))))
	if result != 0 {
		result = int32(*(*uint32)(unsafe.Pointer(uintptr(result + 32))))
		v7 = int32(*(*uint32)(unsafe.Pointer(uintptr(v3 + 40))) - uint32(*(*int16)(unsafe.Pointer(uintptr(v3 + 52)))) + 3)
		*(*uint32)(unsafe.Pointer(uintptr(result + 4))) = uint32(v7)
		if v7 < 0 {
			*(*uint32)(unsafe.Pointer(uintptr(result + 4))) = 0
		}
		*(*float32)(unsafe.Pointer(uintptr(result + 8))) = float32(float64(int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v3 + 36))) + 12)))-*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v3 + 36))) + 400))) + 12))))) / float64(*(*int32)(unsafe.Pointer(uintptr(result + 4)))))
		if a3 != 0 {
			result = (*nox_window)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v3 + 36)))))).Func94(asWindowEvent(0x400A, int32(*(*uint32)(unsafe.Pointer(uintptr(result + 4)))-uint32(*(*int16)(unsafe.Pointer(uintptr(v3 + 54))))), 0))
		}
	}
	return result
}
func nox_xxx_wndListboxProcWithData10_4A2DE0(a1 int32, a2 int32, a3 uint32, a4 int32) int32 {
	var (
		v4     int32
		v5     int32
		result int32
		v7     uint32
		v8     int32
		v9     uint32
		v10    int32
		v11    int32
		v12    *int32
		v13    int32
		v14    int32
		v15    int32
		v16    uint32
		v17    int32
		v18    uint32
		v19    int32
		v20    int32
		v21    int32
	)
	v4 = a1
	v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 32))))
	switch a2 {
	case 5:
		fallthrough
	case 17:
		fallthrough
	case 18:
		return 1
	case 6:
		fallthrough
	case 7:
		v7 = a3 >> 16
		nox_client_wndGetPosition_46AA60((*nox_window)(unsafe.Pointer(uintptr(a1))), (*uint32)(unsafe.Pointer(&a2)), &a3)
		if int32(*(*uint16)(unsafe.Pointer(uintptr(a1 + 108)))) != 0 {
			v8 = nox_xxx_guiFontHeightMB_43F320(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 236))))))
			v9 = a3 + uint32(v8) + 1
			a3 += uint32(v8 + 1)
		} else {
			v9 = a3
		}
		v10 = 0
		v11 = 0
		for 2 != 0 {
			if v11 <= 0 {
				goto LABEL_12
			}
			if *(*uint32)(unsafe.Pointer(uintptr(uint32(v11) + *(*uint32)(unsafe.Pointer(uintptr(v5 + 24))) - 524))) > uint32(int32(*(*int16)(unsafe.Pointer(uintptr(v5 + 52))))+int32(*(*int16)(unsafe.Pointer(uintptr(v5 + 54))))) {
				v10 = -1
			LABEL_16:
				v4 = a1
				goto LABEL_17
			}
			v4 = a1
		LABEL_12:
			if v10 != int32(*(*int16)(unsafe.Pointer(uintptr(v5 + 44)))) {
				if *(*uint32)(unsafe.Pointer(uintptr(uint32(v11) + *(*uint32)(unsafe.Pointer(uintptr(v5 + 24)))))) <= uint32(int32(v7+uint32(*(*int16)(unsafe.Pointer(uintptr(v5 + 54))))-v9)) {
					v4 = a1
					v10++
					v11 += 524
					continue
				}
				goto LABEL_16
			}
			break
		}
		v10 = -1
	LABEL_17:
		v12 = *(**int32)(unsafe.Pointer(uintptr(v5 + 48)))
		v13 = 0
		v14 = *v12
		if *v12 < 0 {
		LABEL_24:
			*(*int32)(unsafe.Add(unsafe.Pointer(v12), unsafe.Sizeof(int32(0))*uintptr(v13))) = v10
			*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v5 + 48))) + uint32(v13*4) + 4))) = math.MaxUint32
			(*nox_window)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v4 + 52)))))).Func94(asWindowEvent(0x4010, v4, v10))
			result = 1
		} else {
			v15 = int32(*(*uint32)(unsafe.Pointer(uintptr(v5 + 48))))
			for v14 != v10 {
				v14 = int32(*(*uint32)(unsafe.Pointer(uintptr(v15 + 4))))
				v15 += 4
				v13++
				if v14 < 0 {
					goto LABEL_24
				}
			}
			sub_4A3090((*int16)(unsafe.Pointer(uintptr(v5))), v13)
			(*nox_window)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v4 + 52)))))).Func94(asWindowEvent(0x4010, v4, v10))
			result = 1
		}
		return result
	case 8:
		v21 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 44))))
		if v21&256 != 0 {
			(*nox_window)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 52)))))).Func94(asWindowEvent(0x4000, a1, 0))
		}
		return 1
	case 10:
		fallthrough
	case 11:
		v16 = a3 >> 16
		nox_client_wndGetPosition_46AA60((*nox_window)(unsafe.Pointer(uintptr(a1))), (*uint32)(unsafe.Pointer(&a2)), &a3)
		if int32(*(*uint16)(unsafe.Pointer(uintptr(a1 + 108)))) != 0 {
			v17 = nox_xxx_guiFontHeightMB_43F320(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 236))))))
			v18 = a3 + uint32(v17) + 1
			a3 += uint32(v17 + 1)
		} else {
			v18 = a3
		}
		v19 = 0
		v20 = 0
		for 2 != 0 {
			if v20 <= 0 {
				goto LABEL_32
			}
			if *(*uint32)(unsafe.Pointer(uintptr(uint32(v20) + *(*uint32)(unsafe.Pointer(uintptr(v5 + 24))) - 524))) > uint32(int32(*(*int16)(unsafe.Pointer(uintptr(v5 + 52))))+int32(*(*int16)(unsafe.Pointer(uintptr(v5 + 54))))) {
				v19 = -1
			LABEL_36:
				(*nox_window)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 52)))))).Func94(asWindowEvent(0x4011, a1, v19))
				return 1
			}
			v4 = a1
		LABEL_32:
			if v19 != int32(*(*int16)(unsafe.Pointer(uintptr(v5 + 44)))) {
				if *(*uint32)(unsafe.Pointer(uintptr(uint32(v20) + *(*uint32)(unsafe.Pointer(uintptr(v5 + 24)))))) <= uint32(int32(v16+uint32(*(*int16)(unsafe.Pointer(uintptr(v5 + 54))))-v18)) {
					v4 = a1
					v19++
					v20 += 524
					continue
				}
				goto LABEL_36
			}
			break
		}
		(*nox_window)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v4 + 52)))))).Func94(asWindowEvent(0x4011, v4, -1))
		return 1
	case 19:
		if *(*uint32)(unsafe.Pointer(uintptr(v5 + 28))) == 0 || int32(*(*int16)(unsafe.Pointer(uintptr(v5 + 54)))) <= 0 {
			return 1
		}
		nox_xxx_wndListBox_4A2D10(a1, -1, 1)
		return 1
	case 20:
		if *(*uint32)(unsafe.Pointer(uintptr(v5 + 32))) == 0 || int32(*(*int16)(unsafe.Pointer(uintptr(v5 + 52))))+int32(*(*int16)(unsafe.Pointer(uintptr(v5 + 54)))) > *(*int32)(unsafe.Pointer(uintptr(v5 + 40))) {
			return 1
		}
		nox_xxx_wndListBox_4A2D10(a1, 1, 1)
		return 1
	case 21:
		if a3 != 15 {
			return 0
		}
		if a4 != 2 {
			return 1
		}
		nox_xxx_wndRetNULL_46A8A0()
		return 1
	default:
		return 0
	}
}
func sub_4A3090(a1 *int16, a2 int32) *int16 {
	var result *int16
	result = a1
	libc.MemCpy(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a1))), unsafe.Sizeof(uint32(0))*12)))+uint32(a2*4))), unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a1))), unsafe.Sizeof(uint32(0))*12)))+uint32(a2*4)+4)), int((int32(*a1)-a2)*4))
	*(*uint32)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a1))), unsafe.Sizeof(uint32(0))*12))) + uint32(int32(*a1)*4) - 4))) = math.MaxUint32
	return result
}
func nox_xxx_wndListboxProcPre_4A30D0(win *nox_window, ev uint32, a3 uint32, a4 int32) int32 {
	var (
		v9    *uint32
		v10   *uint32
		v11   *uint32
		v12   int32
		v13   int16
		v14   *uint32
		v15   int32
		v16   *libc.WChar
		v17   *libc.WChar
		v18   int16
		v19   int16
		v20   int32
		v21   int32
		v22   int32
		v23   int32
		v24   int32
		v25   int32
		v26   *uint32
		v27   int32
		v28   int32
		v29   int16
		v30   int16
		v31   int32
		v32   int16
		v33   **libc.WChar
		v34   int32
		v35   *libc.WChar
		v36   int32
		v37   int32
		v38   int32
		v39   int16
		v40   int32
		v41   *uint32
		v42   int32
		v43   int32
		v44   int32
		v45   int16
		v46   int32
		v4    *uint32                 = (*uint32)(unsafe.Pointer(win))
		sdata *nox_scrollListBox_data = (*nox_scrollListBox_data)(win.widget_data)
		ind   int32                   = 0
		wstr  *libc.WChar             = nil
	)
	if ev > 0x4012 {
		switch ev {
		case 0x4013:
			ind = int32(a3)
			if ind < 0 || ind >= int32(sdata.count) {
				if sdata.field_4 != 0 {
					libc.MemSet(unsafe.Pointer(sdata.field_12), math.MaxUint8, int(int32(sdata.count)*4))
					return 0
				}
				sdata.field_12 = (*uint32)(unsafe.Pointer(uintptr(math.MaxUint32)))
				return 0
			}
			v37 = int32(uintptr(unsafe.Pointer((*nox_scrollListBox_item)(unsafe.Add(unsafe.Pointer(sdata.items), unsafe.Sizeof(nox_scrollListBox_item{})*uintptr(ind))))))
			if int32(*(*uint16)(unsafe.Pointer(uintptr(v37 + 4)))) == 0 {
				return 0
			}
			if sdata.field_4 != 0 {
				*(*uint32)(unsafe.Add(unsafe.Pointer(sdata.field_12), unsafe.Sizeof(uint32(0))*0)) = uint32(ind)
				*(*uint32)(unsafe.Add(unsafe.Pointer(sdata.field_12), unsafe.Sizeof(uint32(0))*1)) = math.MaxUint32
				return 0
			}
			v38 = int32(sdata.field_13_1)
			sdata.field_12 = (*uint32)(unsafe.Pointer(uintptr(ind)))
			if *(*uint32)(unsafe.Pointer(uintptr(v37))) < uint32(v38) {
				win.Func94(asWindowEvent(0x401C, ind, 0))
				return 0
			}
			v39 = int16(sdata.field_13_0)
			if *(*uint32)(unsafe.Pointer(uintptr(v37))) > uint32(v38+int32(v39)) {
				if ind <= 0 {
					sdata.field_13_1 = 0
				} else {
					sdata.field_13_1 = uint16(int16(int32(*(*uint16)(unsafe.Pointer(uintptr(v37)))) - int32(v39)))
				}
				nox_xxx_wndListBox_4A2D10(int32(uintptr(unsafe.Pointer(win))), 0, 1)
				return 0
			}
			return 0
		case 0x4014:
			return int32(uintptr(unsafe.Pointer(sdata.field_12)))
		case 0x4015:
			ind = int32(a3)
			if ind < 0 || ind >= int32(sdata.count) {
				if sdata.field_4 == 0 {
					return 0
				}
				libc.MemSet(unsafe.Pointer(sdata.field_12), math.MaxUint8, int(int32(sdata.count)*4))
				return 0
			}
			if (*(*nox_scrollListBox_item)(unsafe.Add(unsafe.Pointer(sdata.items), unsafe.Sizeof(nox_scrollListBox_item{})*uintptr(ind)))).text[0] == 0 || sdata.field_4 == 0 {
				return 0
			}
			v33 = (**libc.WChar)(unsafe.Pointer(sdata.field_12))
			v34 = 0
			v35 = *v33
			if int32(uintptr(unsafe.Pointer(*v33))) < 0 {
				*(**libc.WChar)(unsafe.Add(unsafe.Pointer(v33), unsafe.Sizeof((*libc.WChar)(nil))*uintptr(v34))) = (*libc.WChar)(unsafe.Pointer(uintptr(ind)))
				*(*uint32)(unsafe.Add(unsafe.Pointer(sdata.field_12), unsafe.Sizeof(uint32(0))*uintptr(v34+1))) = math.MaxUint32
				return 0
			}
			v36 = int32(uintptr(unsafe.Pointer(sdata.field_12)))
			for unsafe.Pointer(v35) != unsafe.Pointer(uintptr(ind)) {
				v35 = *(**libc.WChar)(unsafe.Pointer(uintptr(v36 + 4)))
				v36 += 4
				v34++
				if int32(uintptr(unsafe.Pointer(v35))) < 0 {
					*(**libc.WChar)(unsafe.Add(unsafe.Pointer(v33), unsafe.Sizeof((*libc.WChar)(nil))*uintptr(v34))) = (*libc.WChar)(unsafe.Pointer(uintptr(ind)))
					*(*uint32)(unsafe.Add(unsafe.Pointer(sdata.field_12), unsafe.Sizeof(uint32(0))*uintptr(v34+1))) = math.MaxUint32
					return 0
				}
			}
			sub_4A3090((*int16)(unsafe.Pointer(sdata)), v34)
			return 0
		case 0x4016:
			ind = int32(a3)
			if ind < 0 || ind >= int32(sdata.count) {
				return 0
			}
			return int32(uintptr(unsafe.Pointer(&(*(*nox_scrollListBox_item)(unsafe.Add(unsafe.Pointer(sdata.items), unsafe.Sizeof(nox_scrollListBox_item{})*uintptr(ind)))).text[0])))
		case 0x4017:
			ind = a4
			wstr = (*libc.WChar)(unsafe.Pointer(uintptr(a3)))
			if ind < 0 || ind >= int32(sdata.count) {
				return 0
			}
			nox_wcsncpy(&(*(*nox_scrollListBox_item)(unsafe.Add(unsafe.Pointer(sdata.items), unsafe.Sizeof(nox_scrollListBox_item{})*uintptr(ind)))).text[0], wstr, math.MaxUint8)
			(*(*nox_scrollListBox_item)(unsafe.Add(unsafe.Pointer(sdata.items), unsafe.Sizeof(nox_scrollListBox_item{})*uintptr(ind)))).text[nox_wcslen(wstr)] = 0
			return 0
		case 0x4018:
			sdata.field_7 = unsafe.Pointer(uintptr(a3))
			return 0
		case 0x4019:
			sdata.field_8 = unsafe.Pointer(uintptr(a3))
			return 0
		case 16410:
			sdata.field_9 = unsafe.Pointer(uintptr(a3))
			return 0
		case 0x401B:
			ind = int32(a3)
			v40 = int32(sdata.field_11_0)
			if v40 < ind {
				return 0
			}
			libc.MemMove(unsafe.Pointer(sdata.items), unsafe.Pointer((*nox_scrollListBox_item)(unsafe.Add(unsafe.Pointer(sdata.items), unsafe.Sizeof(nox_scrollListBox_item{})*uintptr(ind)))), int(uint32(v40-ind)*uint32(unsafe.Sizeof(nox_scrollListBox_item{}))))
			sdata.field_11_0 -= uint16(int16(ind))
			sdata.field_11_1 = sdata.field_11_0
			if sdata.field_4 != 0 {
				v41 = sdata.field_12
				v42 = 0
				if int32(*v41) >= 0 {
					for {
						v43 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v41), unsafe.Sizeof(uint32(0))*uintptr(v42))))
						if ind < v43 {
							sub_4A3090((*int16)(unsafe.Pointer(sdata)), func() int32 {
								p := &v42
								x := *p
								*p--
								return x
							}())
						} else {
							*(*uint32)(unsafe.Add(unsafe.Pointer(v41), unsafe.Sizeof(uint32(0))*uintptr(v42))) = uint32(v43 - ind)
						}
						v41 = sdata.field_12
						v42++
						if int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v41), unsafe.Sizeof(uint32(0))*uintptr(v42)))) < 0 {
							break
						}
					}
				}
			} else {
				v44 = int32(uintptr(unsafe.Pointer(sdata.field_12)))
				if v44 > 0 {
					sdata.field_12 = (*uint32)(unsafe.Pointer(uintptr(v44 - ind)))
				}
			}
			if int32(sdata.field_13_1) > 0 {
				nox_xxx_wndListBox_4A2D10(int32(uintptr(unsafe.Pointer(win))), -ind, 1)
			}
			nox_xxx_wndListBox_4A3A70(int32(uintptr(unsafe.Pointer(win))))
			return 1
		case 0x401C:
			ind = int32(a3)
			if ind-1 >= 0 && ind-1 < int32(sdata.count) {
				sdata.field_13_1 = uint16((*(*nox_scrollListBox_item)(unsafe.Add(unsafe.Pointer(sdata.items), unsafe.Sizeof(nox_scrollListBox_item{})*uintptr(ind-1)))).field_0 + 1)
			} else {
				sdata.field_13_1 = 0
			}
			v45 = int16(sdata.field_13_0)
			v46 = int32(sdata.field_10)
			if int32(sdata.field_13_1)+int32(v45) >= v46 {
				sdata.field_13_1 = uint16(int16(v46 - int32(v45)))
			}
			nox_xxx_wndListBox_4A2D10(int32(uintptr(unsafe.Pointer(win))), 0, 1)
			return 0
		default:
			return 0
		}
		return 0
	}
	if ev == 0x4012 {
		if int32(a3) >= 0 {
			v32 = int16(sdata.field_11_0)
			if int32(a3) <= int32(v32) {
				sdata.field_11_1 = uint16(int16(uint16(a3)))
			} else {
				sdata.field_11_1 = uint16(v32)
			}
		} else {
			sdata.field_11_1 = 0
		}
		return 0
	}
	if ev > 0x4007 {
		switch ev {
		case 0x4009:
			v29 = int16(sdata.field_13_0)
			v30 = int16(int32(*(*uint16)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(sdata.field_9)) + 32))) + 4)))) - a4)
			v31 = int32(sdata.field_10 - uint32(v29) + 1)
			sdata.field_13_1 = uint16(v30)
			if int32(v30) > v31 {
				sdata.field_13_1 = uint16(int16(int32(sdata.field_10) - int32(v29) + 1))
			}
			if int32(sdata.field_13_1) < 0 {
				sdata.field_13_1 = 0
			}
			nox_xxx_wndListBox_4A2D10(int32(uintptr(unsafe.Pointer(win))), 0, 0)
			return 0
		case 0x400D:
			v18 = int16(sdata.field_11_1)
			v19 = int16(sdata.field_11_0)
			if int32(v18) != int32(v19) {
				if int32(v19) < int32(sdata.count) {
					v20 = int32(v19) - 1
					if v20 >= int32(v18) {
						v21 = v20 * 524
						for {
							v20--
							v22 = int32(uintptr(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer(&sdata.items.field_0), unsafe.Sizeof(uint32(0))*uintptr(v21))))))
							v21 -= 524
							libc.MemCpy(unsafe.Pointer(uintptr(v22+524)), unsafe.Pointer(uintptr(v22)), 524)
							if v20 < int32(sdata.field_11_1) {
								break
							}
						}
						v4 = (*uint32)(unsafe.Pointer(win))
					}
					nox_xxx_wndListBoxAddLine_4A3AC0((*libc.WChar)(unsafe.Pointer(uintptr(a3))), a4, v4)
				} else {
					if sdata.field_2 == 0 {
						return 0
					}
					win.Func94(asWindowEvent(0x401B, 1, 0))
					v23 = int32(sdata.field_11_0) - 1
					if v23 >= int32(sdata.field_11_1) {
						v24 = v23 * 524
						for {
							v23--
							v25 = int32(uint32(uintptr(unsafe.Pointer(sdata.items))) + uint32(v24))
							v24 -= 524
							libc.MemCpy(unsafe.Pointer(uintptr(v25+524)), unsafe.Pointer(uintptr(v25)), 524)
							if v23 < int32(sdata.field_11_1) {
								break
							}
						}
						v4 = (*uint32)(unsafe.Pointer(win))
					}
					nox_xxx_wndListBoxAddLine_4A3AC0((*libc.WChar)(unsafe.Pointer(uintptr(a3))), a4, v4)
				}
			} else if int32(v18) >= int32(sdata.count) {
				if sdata.field_2 == 0 {
					return 0
				}
				win.Func94(asWindowEvent(0x401B, 1, 0))
				nox_xxx_wndListBoxAddLine_4A3AC0((*libc.WChar)(unsafe.Pointer(uintptr(a3))), a4, (*uint32)(unsafe.Pointer(win)))
			} else {
				nox_xxx_wndListBoxAddLine_4A3AC0((*libc.WChar)(unsafe.Pointer(uintptr(a3))), a4, (*uint32)(unsafe.Pointer(win)))
			}
			if sdata.field_1 != 0 {
				for (*(*nox_scrollListBox_item)(unsafe.Add(unsafe.Pointer(sdata.items), unsafe.Sizeof(nox_scrollListBox_item{})*uintptr(int32(sdata.field_11_1)-1)))).field_0 >= uint32(int32(sdata.field_13_1)+int32(sdata.field_13_0)) {
					nox_xxx_wndListBox_4A2D10(int32(uintptr(unsafe.Pointer(v4))), 1, 1)
				}
			}
			v26 = sdata.field_12
			if sdata.field_4 != 0 {
				if int32(*v26) >= 0 {
					v27 = 0
					for {
						v28 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v26), unsafe.Sizeof(uint32(0))*uintptr(v27))))
						if int32(sdata.field_11_1) < v28 {
							*(*uint32)(unsafe.Add(unsafe.Pointer(v26), unsafe.Sizeof(uint32(0))*uintptr(v27))) = uint32(v28 + 1)
						}
						v26 = sdata.field_12
						v27++
						if int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v26), unsafe.Sizeof(uint32(0))*uintptr(v27)))) < 0 {
							break
						}
					}
					return 1
				}
			} else if int32(sdata.field_11_1) < int32(uintptr(unsafe.Pointer(v26))) {
				sdata.field_12 = (*uint32)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(v26))), 1))
			}
			return 1
		case 0x400E:
			ind = int32(a3)
			v12 = int32(sdata.field_11_0)
			if ind < 0 || ind >= v12 {
				return 0
			}
			for i := int32(ind); i < v12-1; i++ {
				libc.MemCpy(unsafe.Pointer((*nox_scrollListBox_item)(unsafe.Add(unsafe.Pointer(sdata.items), unsafe.Sizeof(nox_scrollListBox_item{})*uintptr(i)))), unsafe.Pointer((*nox_scrollListBox_item)(unsafe.Add(unsafe.Pointer(sdata.items), unsafe.Sizeof(nox_scrollListBox_item{})*uintptr(i+1)))), int(unsafe.Sizeof(nox_scrollListBox_item{})))
			}
			v13 = int16(func() uint16 {
				p := &sdata.field_11_0
				*p--
				return *p
			}())
			*(*nox_scrollListBox_item)(unsafe.Add(unsafe.Pointer(sdata.items), unsafe.Sizeof(nox_scrollListBox_item{})*uintptr(v13))) = nox_scrollListBox_item{}
			sdata.field_11_1 = uint16(v13)
			(*(*nox_scrollListBox_item)(unsafe.Add(unsafe.Pointer(sdata.items), unsafe.Sizeof(nox_scrollListBox_item{})*uintptr(v13)))).text[0] = 0
			if sdata.field_4 != 0 {
				v14 = sdata.field_12
				v15 = 0
				if int32(*v14) >= 0 {
					for {
						v16 = (*libc.WChar)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v14), unsafe.Sizeof(uint32(0))*uintptr(v15))))))
						if ind >= int32(uintptr(unsafe.Pointer(v16))) {
							if unsafe.Pointer(uintptr(ind)) == unsafe.Pointer(v16) {
								sub_4A3090((*int16)(unsafe.Pointer(sdata)), func() int32 {
									p := &v15
									x := *p
									*p--
									return x
								}())
							}
						} else {
							*(*uint32)(unsafe.Add(unsafe.Pointer(v14), unsafe.Sizeof(uint32(0))*uintptr(v15))) = uint32(uintptr(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(v16))), -1)))))
						}
						v14 = sdata.field_12
						v15++
						if int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v14), unsafe.Sizeof(uint32(0))*uintptr(v15)))) < 0 {
							break
						}
					}
					nox_xxx_wndListBox_4A3A70(int32(uintptr(unsafe.Pointer(win))))
					return 0
				}
			} else {
				v17 = (*libc.WChar)(unsafe.Pointer(sdata.field_12))
				if ind < int32(uintptr(unsafe.Pointer(v17))) {
					sdata.field_12 = (*uint32)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(v17))), -1))
					nox_xxx_wndListBox_4A3A70(int32(uintptr(unsafe.Pointer(win))))
					return 0
				}
				if unsafe.Pointer(uintptr(ind)) == unsafe.Pointer(v17) {
					sdata.field_12 = (*uint32)(unsafe.Pointer(uintptr(math.MaxUint32)))
				}
			}
			nox_xxx_wndListBox_4A3A70(int32(uintptr(unsafe.Pointer(win))))
			return 0
		case 0x400F:
			libc.MemSet(unsafe.Pointer(sdata.items), 0, int(int32(sdata.count)*int32(unsafe.Sizeof(nox_scrollListBox_item{}))))
			if a3 != 1 {
				sdata.field_13_1 = 0
			}
			if sdata.field_4 != 0 {
				libc.MemSet(unsafe.Pointer(sdata.field_12), math.MaxUint8, int(int32(sdata.count)*4))
			} else {
				sdata.field_12 = (*uint32)(unsafe.Pointer(uintptr(math.MaxUint32)))
			}
			sdata.field_11_1 = 0
			sdata.field_11_0 = 0
			sdata.field_10 = 0
			nox_xxx_wndListBox_4A2D10(int32(uintptr(unsafe.Pointer(win))), 0, 1)
			return 0
		default:
			return 0
		}
	}
	if ev == 0x4007 {
		if unsafe.Pointer(uintptr(a3)) == sdata.field_7 {
			if int32(sdata.field_13_1) > 0 {
				nox_xxx_wndListBox_4A2D10(int32(uintptr(unsafe.Pointer(win))), -1, 1)
				return 0
			}
		} else if unsafe.Pointer(uintptr(a3)) == sdata.field_8 && uint32(int32(sdata.field_13_0)+int32(sdata.field_13_1)) <= sdata.field_10 {
			nox_xxx_wndListBox_4A2D10(int32(uintptr(unsafe.Pointer(win))), 1, 1)
			return 0
		}
		return 0
	}
	if ev == 0x4001 {
		wstr = (*libc.WChar)(unsafe.Pointer(uintptr(a3)))
		nox_wcsncpy(&win.draw_data.text[0], wstr, 63)
		win.draw_data.text[nox_wcslen(wstr)] = 0
		return 0
	}
	if ev == 0x4004 {
		v9 = (*uint32)(sdata.field_7)
		if v9 != nil {
			(*nox_window)(unsafe.Pointer(v9)).SetPos(image.Pt(int32(uint32(int32(a3))-*(*uint32)(unsafe.Add(unsafe.Pointer(v9), unsafe.Sizeof(uint32(0))*2))), 0))
		}
		v10 = (*uint32)(sdata.field_8)
		if v10 != nil {
			(*nox_window)(unsafe.Pointer(v10)).SetPos(image.Pt(int32(uint32(int32(a3))-*(*uint32)(unsafe.Add(unsafe.Pointer(v10), unsafe.Sizeof(uint32(0))*2))), int32(uint32(a4)-*(*uint32)(unsafe.Add(unsafe.Pointer(v10), unsafe.Sizeof(uint32(0))*3)))))
		}
		v11 = (*uint32)(sdata.field_9)
		if v11 != nil {
			(*nox_window)(unsafe.Pointer(v11)).SetPos(image.Pt(int32(uint32(int32(a3))-*(*uint32)(unsafe.Add(unsafe.Pointer(v11), unsafe.Sizeof(uint32(0))*2))), int32(*(*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(sdata.field_7)) + 12))))))
			sub_46AB20((*uint32)(sdata.field_9), int32(*(*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(sdata.field_9)) + 8)))), int32(uint32(a4)-*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(sdata.field_9)) + 400))) + 12)))*2))
		}
		sdata.field_13_0 = uint16(int16(a4))
		if win.draw_data.text[0] != 0 {
			sdata.field_13_0 -= uint16(int16(nox_xxx_guiFontHeightMB_43F320(win.draw_data.font)))
		}
		return 0
	}
	if ev > 0x4000 {
		return 0
	}
	if ev == 0x4000 {
		if unsafe.Pointer(uintptr(a3)) == sdata.field_7 {
			if int32(sdata.field_13_1) > 0 {
				nox_xxx_wndListBox_4A2D10(int32(uintptr(unsafe.Pointer(win))), -1, 1)
				return 0
			}
		} else if unsafe.Pointer(uintptr(a3)) == sdata.field_8 && uint32(int32(sdata.field_13_0)+int32(sdata.field_13_1)) <= sdata.field_10 {
			nox_xxx_wndListBox_4A2D10(int32(uintptr(unsafe.Pointer(win))), 1, 1)
			return 0
		}
		return 0
	}
	if ev == 2 {
		if sdata != nil {
			alloc.Free(unsafe.Pointer(sdata.items))
			if sdata.field_4 != 0 {
				alloc.Free(unsafe.Pointer(sdata.field_12))
			}
			alloc.Free(unsafe.Pointer(sdata))
		}
		win.widget_data = nil
		return 0
	}
	if ev == 23 {
		if a3 != 0 {
			win.draw_data.field_0 |= 2
		} else {
			win.draw_data.field_0 &= 0xFFFFFFFD
		}
		var v7 int32 = win.ID()
		win.draw_data.win.Func94(asWindowEvent(0x4003, int32(a3), v7))
		return 1
	}
	return 0
}
func nox_xxx_wndListBox_4A3A70(a1 int32) int32 {
	var (
		v1 int32
		v2 int32
		v3 int32
		v4 int32
		v5 int32
	)
	v1 = 0
	v2 = 0
	v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 32))))
	if int32(*(*uint16)(unsafe.Pointer(uintptr(v3 + 44)))) > 0 {
		v4 = 0
		for {
			v5 = int32(uint32(v4) + *(*uint32)(unsafe.Pointer(uintptr(v3 + 24))))
			v1++
			v4 += 524
			v2 += int32(*(*uint8)(unsafe.Pointer(uintptr(v5 + 520)))) + 1
			*(*uint32)(unsafe.Pointer(uintptr(v5))) = uint32(v2)
			if v1 >= int32(*(*int16)(unsafe.Pointer(uintptr(v3 + 44)))) {
				break
			}
		}
	}
	*(*uint32)(unsafe.Pointer(uintptr(v3 + 40))) = uint32(v2)
	return nox_xxx_wndListBox_4A2D10(a1, 0, 1)
}
func nox_xxx_wndListBoxAddLine_4A3AC0(a1 *libc.WChar, a2 int32, a3 *uint32) int32 {
	var (
		v3 *uint32
		v4 int32
		v5 int32
		v6 *libc.WChar
		v7 *libc.WChar
	)
	v3 = a3
	v4 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a3), unsafe.Sizeof(uint32(0))*8)))
	v5 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a3), unsafe.Sizeof(uint32(0))*2)) - 7)
	if *(*uint32)(unsafe.Pointer(uintptr(v4 + 12))) != 0 {
		v5 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a3), unsafe.Sizeof(uint32(0))*2)) - 17)
	}
	if a2 >= 17 || a2 < 0 {
		*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v4 + 24))) + uint32(int32(*(*int16)(unsafe.Pointer(uintptr(v4 + 46))))*524) + 516))) = *(*uint32)(unsafe.Add(unsafe.Pointer(a3), unsafe.Sizeof(uint32(0))*26))
	} else {
		*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v4 + 24))) + uint32(int32(*(*int16)(unsafe.Pointer(uintptr(v4 + 46))))*524) + 516))) = **(**uint32)(memmap.PtrOff(0x85B3FC, uint32(a2*4+132)))
	}
	v6 = (*libc.WChar)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v4 + 24))) + uint32(int32(*(*int16)(unsafe.Pointer(uintptr(v4 + 46))))*524) + 4)))
	if a1 != nil {
		nox_wcsncpy(v6, a1, math.MaxUint8)
		*(*libc.WChar)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(libc.WChar(0))*math.MaxUint8)) = 0
		v7 = (*libc.WChar)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(libc.WChar(0))*uintptr(nox_wcslen(v6)-1)))
		if *v7 == 10 {
			*v7 = 0
		}
	} else {
		*v6 = 32
		*(*libc.WChar)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(libc.WChar(0))*1)) = 0
	}
	if (*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*1)) & 0x4000) == 0x4000 {
		*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v4 + 24))) + uint32(int32(*(*int16)(unsafe.Pointer(uintptr(v4 + 46))))*524) + 520))) = uint8(int8(nox_xxx_guiFontHeightMB_43F320(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*59)))))))
	} else {
		nox_xxx_drawGetStringSize_43F840(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*59)))), v6, nil, (*int32)(unsafe.Pointer(&a3)), v5)
		*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v4 + 24))) + uint32(int32(*(*int16)(unsafe.Pointer(uintptr(v4 + 46))))*524) + 520))) = uint8(uintptr(unsafe.Pointer(a3)))
	}
	*(*uint16)(unsafe.Pointer(uintptr(v4 + 46)))++
	*(*uint16)(unsafe.Pointer(uintptr(v4 + 44)))++
	return nox_xxx_wndListBox_4A3A70(int32(uintptr(unsafe.Pointer(v3))))
}
func nox_xxx_wndListboxInit_4A3C00(a1 int32, a2 int32) {
	if a1 != 0 {
		if int32(int8(*(*uint8)(unsafe.Pointer(uintptr(a1 + 4))))) >= 0 {
			a1.setDraw(func(arg1 int32, arg2 int32) int32 {
				return nox_xxx_wndListboxDrawNoImage_4A3C50((*uint32)(unsafe.Pointer(uintptr(arg1))), arg2)
			})
		} else {
			a1.setDraw(func(arg1 int32, arg2 int32) int32 {
				return nox_xxx_wndListboxDrawWithImage_4A3FC0((*uint32)(unsafe.Pointer(uintptr(arg1))), arg2)
			})
		}
		if *(*uint32)(unsafe.Pointer(uintptr(a2 + 16))) != 0 {
			a1.setFunc93(func(arg1 int32, arg2 int32, arg3 int32, arg4 int32) int32 {
				return nox_xxx_wndListboxProcWithData10_4A2DE0(arg1, arg2, uint32(arg3), arg4)
			})
		} else {
			a1.setFunc93(func(arg1 int32, arg2 int32, arg3 int32, arg4 int32) int32 {
				return nox_xxx_wndListboxProcWithoutData10_4A28E0((*uint32)(unsafe.Pointer(uintptr(arg1))), arg2, uint32(arg3), arg4)
			})
		}
	}
}
func nox_xxx_wndListboxDrawNoImage_4A3C50(a1 *uint32, a2 int32) int32 {
	var (
		v2    int32
		v3    int32
		v4    int32
		v5    int32
		v6    int32
		v7    int32
		v8    int32
		v9    int32
		v10   int32
		v11   int32
		v12   int32
		v13   *int32
		v14   int32
		v15   *int16
		v16   int32
		v17   int32
		v19   int32
		v20   int32
		v21   int32
		v22   int32
		v23   int32
		v24   int32
		yTop  int32
		xLeft int32
		v27   int32
		i     int32
		v29   int32
		v30   [256]int16
	)
	v2 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*8)))
	v20 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 28))))
	v23 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 20))))
	nox_client_wndGetPosition_46AA60((*nox_window)(unsafe.Pointer(a1)), (*uint32)(unsafe.Pointer(&xLeft)), (*uint32)(unsafe.Pointer(&yTop)))
	v3 = nox_xxx_guiFontHeightMB_43F320(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a2 + 200))))))
	v4 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*3)))
	v22 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*2)))
	if (*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*1)) & 8192) == 8192 {
		nox_draw_enableTextSmoothing_43F670(1)
	}
	nox_xxx_wndDraw_49F7F0()
	if *(*uint32)(unsafe.Pointer(uintptr(v2 + 12))) != 0 {
		v22 -= 10
	}
	if int32(*(*uint16)(unsafe.Pointer(uintptr(a2 + 72)))) != 0 {
		nox_xxx_drawSetTextColor_434390(int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 68)))))
		nox_xxx_drawStringWrap_43FAF0(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a2 + 200))))), (*libc.WChar)(unsafe.Pointer(uintptr(a2+72))), xLeft+1, yTop, v22, v3)
		v4 += int32(-1 - v3)
		yTop += v3 + 1
	}
	if *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*1))&8 != 0 {
		if int32(*(*uint8)(unsafe.Pointer(uintptr(a2))))&2 != 0 {
			v6 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 36))))
			v5 = v23
			goto LABEL_13
		}
		v5 = v23
	} else {
		v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 44))))
	}
	v6 = v20
LABEL_13:
	if uint32(v5) != 0x80000000 {
		nox_client_drawSetColor_434460(v5)
		nox_client_drawRectFilledOpaque_49CE30(xLeft, yTop, v22, v4)
	}
	if uint32(v6) != 0x80000000 {
		nox_client_drawSetColor_434460(v6)
		nox_client_drawBorderLines_49CC70(xLeft, yTop, v22, v4)
	}
	nox_client_copyRect_49F6F0(xLeft, yTop, v22, v4)
	v7 = int32(*(*int16)(unsafe.Pointer(uintptr(v2 + 54))))
	v8 = yTop - v7
	v21 = yTop - v7
	if *(*uint32)(unsafe.Pointer(uintptr(a2 + 68))) != 0x80000000 {
		v9 = 0
		v24 = 0
		for i = 0; ; i = v9 {
			if v9 > 0 && *(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 24))) + uint32(v9) - 524))) > uint32(int32(*(*int16)(unsafe.Pointer(uintptr(v2 + 54))))+int32(*(*int16)(unsafe.Pointer(uintptr(v2 + 52))))) || v24 == int32(*(*int16)(unsafe.Pointer(uintptr(v2 + 44)))) {
				break
			}
			v10 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 24))) + uint32(v9))
			if *(*uint32)(unsafe.Pointer(uintptr(v10))) < uint32(*(*int16)(unsafe.Pointer(uintptr(v2 + 54)))) {
				v21 = v8 + int32(*(*uint8)(unsafe.Pointer(uintptr(v10 + 520)))) + 1
				goto LABEL_38
			}
			v11 = int32(*(*uint8)(unsafe.Pointer(uintptr(v10 + 520))))
			v19 = int32(*(*uint32)(unsafe.Pointer(uintptr(v10 + 516))))
			v12 = v11 + 1
			v27 = v11 + 1
			nox_xxx_drawSetTextColor_434390(v19)
			if *(*uint32)(unsafe.Pointer(uintptr(v2 + 16))) != 0 {
				v13 = *(**int32)(unsafe.Pointer(uintptr(v2 + 48)))
				v14 = *v13
				if *v13 >= 0 {
					for v24 != v14 {
						v14 = *(*int32)(unsafe.Add(unsafe.Pointer(v13), unsafe.Sizeof(int32(0))*1))
						v13 = (*int32)(unsafe.Add(unsafe.Pointer(v13), unsafe.Sizeof(int32(0))*1))
						if v14 < 0 {
							goto LABEL_32
						}
					}
				LABEL_30:
					if *(*uint32)(unsafe.Pointer(uintptr(a2 + 52))) != 0x80000000 {
						nox_client_drawSetColor_434460(int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 52)))))
						nox_client_drawRectFilledOpaque_49CE30(xLeft, v21, v22, v12)
					}
					goto LABEL_32
				}
			} else if uint32(v24) == *(*uint32)(unsafe.Pointer(uintptr(v2 + 48))) {
				goto LABEL_30
			}
		LABEL_32:
			nox_xxx_drawSetTextColor_434390(int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 24))) + uint32(v9) + 516)))))
			if (*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*1)) & 0x4000) == 0x4000 {
				nox_wcscpy((*libc.WChar)(unsafe.Pointer(&v30[0])), (*libc.WChar)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 24)))+uint32(v9)+4))))
				v15 = &v30[nox_wcslen((*libc.WChar)(unsafe.Pointer(&v30[0])))]
				v16 = v22 - 7
				for {
					v17 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 200))))
					*v15 = 0
					v15 = (*int16)(unsafe.Add(unsafe.Pointer(v15), -int(unsafe.Sizeof(int16(0))*1)))
					nox_xxx_drawGetStringSize_43F840(unsafe.Pointer(uintptr(v17)), (*libc.WChar)(unsafe.Pointer((*uint16)(unsafe.Pointer(&v30[0])))), &v29, nil, 0)
					if v29 <= v16 {
						break
					}
				}
				nox_xxx_drawStringWrap_43FAF0(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a2 + 200))))), (*libc.WChar)(unsafe.Pointer(&v30[0])), xLeft+5, v21+2, v16, v27)
				v9 = i
				v12 = v27
			} else {
				nox_xxx_drawStringWrap_43FAF0(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a2 + 200))))), (*libc.WChar)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 24)))+uint32(v9)+4))), xLeft+5, v21+2, v22-7, v12)
			}
			v21 += v12
		LABEL_38:
			v8 = v21
			v9 += 524
			v24++
		}
	}
	sub_49F860()
	nox_draw_enableTextSmoothing_43F670(0)
	return 1
}
func nox_xxx_wndListboxDrawWithImage_4A3FC0(a1 *uint32, a2 int32) int32 {
	var (
		v2    int32
		v3    int32
		v4    int32
		v5    int32
		v6    int32
		v7    int32
		v8    int32
		v9    int32
		v10   int32
		v11   int32
		v12   int32
		v13   int32
		v14   *int32
		v15   int32
		v16   *int16
		v17   int32
		v18   int32
		v20   int32
		v21   int32
		v22   int32
		yTop  int32
		xLeft int32
		v25   int32
		v26   int32
		i     int32
		v28   int32
		v29   int32
		v30   [256]int16
	)
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 24))))
	v3 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*8)))
	v28 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 28))))
	i = v2
	nox_client_wndGetPosition_46AA60((*nox_window)(unsafe.Pointer(a1)), (*uint32)(unsafe.Pointer(&xLeft)), (*uint32)(unsafe.Pointer(&yTop)))
	v4 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*3)))
	v22 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*2)))
	if (*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*1)) & 8192) == 8192 {
		nox_draw_enableTextSmoothing_43F670(1)
	}
	nox_xxx_wndDraw_49F7F0()
	if *(*uint32)(unsafe.Pointer(uintptr(v3 + 12))) != 0 {
		v22 -= 10
	}
	if *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*1))&8 != 0 {
		v5 = i
	} else {
		v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 48))))
	}
	if v5 != 0 {
		nox_client_drawImageAt_47D2C0((*nox_video_bag_image_t)(unsafe.Pointer(uintptr(v5))), xLeft, yTop)
	}
	if int32(*(*uint16)(unsafe.Pointer(uintptr(a2 + 72)))) != 0 {
		nox_xxx_drawSetTextColor_434390(int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 68)))))
		nox_xxx_drawStringWrap_43FAF0(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a2 + 200))))), (*libc.WChar)(unsafe.Pointer(uintptr(a2+72))), xLeft+1, yTop, v22, 0)
		v6 = nox_xxx_guiFontHeightMB_43F320(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a2 + 200))))))
		v20 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 200))))
		yTop += v6 + 1
		v4 += int32(-1 - nox_xxx_guiFontHeightMB_43F320(unsafe.Pointer(uintptr(v20))))
	}
	if uint32(v28) != 0x80000000 {
		nox_client_drawSetColor_434460(v28)
		nox_client_drawBorderLines_49CC70(xLeft, yTop, v22, v4)
	}
	nox_client_copyRect_49F6F0(xLeft, yTop, v22, v4)
	v7 = int32(*(*int16)(unsafe.Pointer(uintptr(v3 + 54))))
	v8 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 68))))
	v9 = yTop - v7
	v25 = yTop - v7
	if uint32(v8) != 0x80000000 {
		v10 = 0
		v26 = 0
		for i = 0; ; i = v10 {
			if v10 > 0 && *(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v3 + 24))) + uint32(v10) - 524))) > uint32(int32(*(*int16)(unsafe.Pointer(uintptr(v3 + 54))))+int32(*(*int16)(unsafe.Pointer(uintptr(v3 + 52))))) || v26 == int32(*(*int16)(unsafe.Pointer(uintptr(v3 + 44)))) {
				break
			}
			v11 = int32(uint32(v10) + *(*uint32)(unsafe.Pointer(uintptr(v3 + 24))))
			if *(*uint32)(unsafe.Pointer(uintptr(v11))) < uint32(*(*int16)(unsafe.Pointer(uintptr(v3 + 54)))) {
				v25 = v9 + int32(*(*uint8)(unsafe.Pointer(uintptr(v11 + 520)))) + 1
				goto LABEL_35
			}
			v12 = int32(*(*uint8)(unsafe.Pointer(uintptr(v11 + 520))))
			v21 = int32(*(*uint32)(unsafe.Pointer(uintptr(v11 + 516))))
			v13 = v12 + 1
			v28 = v12 + 1
			nox_xxx_drawSetTextColor_434390(v21)
			if *(*uint32)(unsafe.Pointer(uintptr(v3 + 16))) != 0 {
				v14 = *(**int32)(unsafe.Pointer(uintptr(v3 + 48)))
				v15 = *v14
				if *v14 >= 0 {
					for v26 != v15 {
						v15 = *(*int32)(unsafe.Add(unsafe.Pointer(v14), unsafe.Sizeof(int32(0))*1))
						v14 = (*int32)(unsafe.Add(unsafe.Pointer(v14), unsafe.Sizeof(int32(0))*1))
						if v15 < 0 {
							goto LABEL_29
						}
					}
				LABEL_27:
					if *(*uint32)(unsafe.Pointer(uintptr(a2 + 52))) != 0x80000000 {
						nox_client_drawSetColor_434460(int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 52)))))
						nox_client_drawRectFilledOpaque_49CE30(xLeft, v25, v22, v13)
					}
					goto LABEL_29
				}
			} else if uint32(v26) == *(*uint32)(unsafe.Pointer(uintptr(v3 + 48))) {
				goto LABEL_27
			}
		LABEL_29:
			nox_xxx_drawSetTextColor_434390(int32(*(*uint32)(unsafe.Pointer(uintptr(uint32(v10) + *(*uint32)(unsafe.Pointer(uintptr(v3 + 24))) + 516)))))
			if (*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*1)) & 0x4000) == 0x4000 {
				nox_wcscpy((*libc.WChar)(unsafe.Pointer(&v30[0])), (*libc.WChar)(unsafe.Pointer(uintptr(uint32(v10)+*(*uint32)(unsafe.Pointer(uintptr(v3 + 24)))+4))))
				v16 = &v30[nox_wcslen((*libc.WChar)(unsafe.Pointer(&v30[0])))]
				v17 = v22 - 7
				for {
					v18 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 200))))
					*v16 = 0
					v16 = (*int16)(unsafe.Add(unsafe.Pointer(v16), -int(unsafe.Sizeof(int16(0))*1)))
					nox_xxx_drawGetStringSize_43F840(unsafe.Pointer(uintptr(v18)), (*libc.WChar)(unsafe.Pointer((*uint16)(unsafe.Pointer(&v30[0])))), &v29, nil, 0)
					if v29 <= v17 {
						break
					}
				}
				nox_xxx_drawStringWrap_43FAF0(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a2 + 200))))), (*libc.WChar)(unsafe.Pointer(&v30[0])), xLeft+5, v25+2, v17, v28)
				v10 = i
				v13 = v28
			} else {
				nox_xxx_drawStringWrap_43FAF0(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a2 + 200))))), (*libc.WChar)(unsafe.Pointer(uintptr(uint32(v10)+*(*uint32)(unsafe.Pointer(uintptr(v3 + 24)))+4))), xLeft+5, v25+2, v22-7, v13)
			}
			v25 += v13
		LABEL_35:
			v9 = v25
			v10 += 524
			v26++
		}
	}
	sub_49F860()
	nox_draw_enableTextSmoothing_43F670(0)
	return 1
}
func sub_4A4800(a1 int32) int32 {
	var (
		result int32
		v2     *uint32
		v3     int32
		v4     int32
	)
	result = 0
	v2 = *(**uint32)(unsafe.Pointer(uintptr(a1 + 24)))
	v3 = int32(*(*int16)(unsafe.Pointer(uintptr(a1 + 54))))
	if *v2 <= uint32(v3) {
		for result < int32(*(*int16)(unsafe.Pointer(uintptr(a1 + 44)))) {
			v4 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*131)))
			v2 = (*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*131))
			result++
			if v4 > v3 {
				return result
			}
		}
		result = 0
	}
	return result
}
func nox_game_showSelClass_4A4840() int32 {
	var (
		result int32
		v1     *uint32
		v2     *uint32
		v3     *uint32
	)
	sub_5007E0(CString("*:*"))
	sub_4A1BE0(1)
	dword_5d4594_1307724 = uint32(uintptr(unsafe.Pointer(nox_xxx_getHostInfoPtr_431770())))
	gameAddStateCode(600)
	result = int32(uintptr(unsafe.Pointer(newWindowFromFile("SelClass.wnd", unsafe.Pointer(cgoFuncAddr(sub_4A4A20))))))
	dword_5d4594_1307736 = uint32(result)
	if result != 0 {
		result.setFunc93(sub_4A18E0)
		result = int32(uintptr(unsafe.Pointer(nox_gui_makeAnimation((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1307736)))), 0, 0, 0, -460, 0, 20, 0, -40))))
		nox_wnd_xxx_1307732 = (*nox_gui_animation)(unsafe.Pointer(uintptr(result)))
		if result != 0 {
			nox_wnd_xxx_1307732.field_0 = 600
			nox_wnd_xxx_1307732.field_12 = sub_4A4970
			nox_wnd_xxx_1307732.fnc_done_out = sub_4A49A0
			v1 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1307736)))).ChildByID(601)))
			int32(uintptr(unsafe.Pointer(v1))).setDraw(sub_4A49D0)
			v2 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1307736)))).ChildByID(603)))
			int32(uintptr(unsafe.Pointer(v2))).setDraw(sub_4A49D0)
			v3 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1307736)))).ChildByID(602)))
			int32(uintptr(unsafe.Pointer(v3))).setDraw(sub_4A49D0)
			*memmap.PtrUint32(6112660, 1307728) = uint32(uintptr(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1307736)))).ChildByID(610))))
			nox_xxx_wndRetNULL_46A8A0()
			*memmap.PtrUint32(6112660, 1307740) = 0
			sub_4A19F0(CString("OptsBack.wnd:Back"))
			sub_4602F0()
			result = 1
		}
	}
	return result
}
func sub_4A4970() int32 {
	nox_wnd_xxx_1307732.state = nox_gui_anim_state(NOX_GUI_ANIM_OUT)
	sub_43BE40(2)
	clientPlaySoundSpecial(923, 100)
	return 1
}
func sub_4A49A0() int32 {
	var v0 func() int32
	v0 = nox_wnd_xxx_1307732.field_13
	nox_gui_freeAnimation_43C570(nox_wnd_xxx_1307732)
	(*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1307736)))).Destroy()
	v0()
	return 1
}
func sub_4A49D0(yTop int32, a2 int32) int32 {
	var (
		v1    *uint32
		xLeft int32
	)
	v1 = (*uint32)(unsafe.Pointer(uintptr(yTop)))
	if *memmap.PtrUint32(6112660, 1307740) != *(*uint32)(unsafe.Pointer(uintptr(yTop))) {
		nox_client_wndGetPosition_46AA60((*nox_window)(unsafe.Pointer(uintptr(yTop))), (*uint32)(unsafe.Pointer(&xLeft)), (*uint32)(unsafe.Pointer(&yTop)))
		nox_client_drawRectFilledAlpha_49CF10(xLeft, yTop, int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*6))-*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*4))), int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*7))-*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*5))))
	}
	return 1
}
func nox_xxx_clearAdminFlag() unsafe.Pointer {
	noxflags.UnsetEngine(nox_engine_flag(NOX_ENGINE_FLAG_ADMIN))
	return sub_4602F0()
}
func nox_xxx_setAdminFlag() {
	noxflags.SetEngine(nox_engine_flag(NOX_ENGINE_FLAG_ADMIN))
}
func sub_4A5E90() int32 {
	var (
		i      **byte
		v1     *uint32
		v2     *uint32
		v3     *uint32
		v4     *uint32
		v5     *uint32
		v6     *uint32
		v7     *uint32
		v8     *uint32
		v9     *uint32
		v10    *uint32
		result int32
	)
	if *memmap.PtrUint32(0x587000, 171384) != 0 {
		sub_4A62B0()
	}
	for i = (**byte)(sub_413390()); i != nil; i = (**byte)(unsafe.Pointer(uintptr(sub_4133A0(int32(uintptr(unsafe.Pointer(i))))))) {
		if libc.StrCmp(*i, CString("StreetPants")) == 0 {
			dword_5d4594_1308156 = uint32(uintptr(unsafe.Pointer(i)))
		} else if libc.StrCmp(*i, CString("StreetShirt")) == 0 {
			dword_5d4594_1308160 = uint32(uintptr(unsafe.Pointer(i)))
		} else if libc.StrCmp(*i, CString("StreetSneakers")) == 0 {
			dword_5d4594_1308164 = uint32(uintptr(unsafe.Pointer(i)))
		}
	}
	v1 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1308084)))).ChildByID(720)))
	dword_5d4594_1308096 = uint32(uintptr(unsafe.Pointer(v1)))
	*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*8)) = 0x20002
	v2 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1308084)))).ChildByID(721)))
	dword_5d4594_1308100 = uint32(uintptr(unsafe.Pointer(v2)))
	*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*8)) = 0x90001
	v3 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1308084)))).ChildByID(722)))
	dword_5d4594_1308104 = uint32(uintptr(unsafe.Pointer(v3)))
	*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*8)) = 0x90001
	v4 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1308084)))).ChildByID(723)))
	dword_5d4594_1308108 = uint32(uintptr(unsafe.Pointer(v4)))
	*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*8)) = 0x90001
	v5 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1308084)))).ChildByID(724)))
	dword_5d4594_1308112 = uint32(uintptr(unsafe.Pointer(v5)))
	*(*uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(uint32(0))*8)) = 0x90001
	v6 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1308084)))).ChildByID(725)))
	dword_5d4594_1308116 = uint32(uintptr(unsafe.Pointer(v6)))
	*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*8)) = uint32(int32(*memmap.PtrUint16(0x587000, 171372)) << 16)
	v7 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1308084)))).ChildByID(726)))
	dword_5d4594_1308120 = uint32(uintptr(unsafe.Pointer(v7)))
	*(*uint32)(unsafe.Add(unsafe.Pointer(v7), unsafe.Sizeof(uint32(0))*8)) = uint32(int32(*memmap.PtrUint16(0x587000, 171374)) << 16)
	v8 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1308084)))).ChildByID(727)))
	dword_5d4594_1308124 = uint32(uintptr(unsafe.Pointer(v8)))
	*(*uint32)(unsafe.Add(unsafe.Pointer(v8), unsafe.Sizeof(uint32(0))*8)) = uint32(int32(*memmap.PtrUint16(0x587000, 171376)) << 16)
	v9 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1308084)))).ChildByID(728)))
	dword_5d4594_1308128 = uint32(uintptr(unsafe.Pointer(v9)))
	*(*uint32)(unsafe.Add(unsafe.Pointer(v9), unsafe.Sizeof(uint32(0))*8)) = uint32(int32(*memmap.PtrUint16(0x587000, 171378)) << 16)
	v10 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1308084)))).ChildByID(729)))
	dword_5d4594_1308132 = uint32(uintptr(unsafe.Pointer(v10)))
	*(*uint32)(unsafe.Add(unsafe.Pointer(v10), unsafe.Sizeof(uint32(0))*8)) = uint32(int32(*memmap.PtrUint16(0x587000, 171380)) << 16)
	dword_5d4594_1308136 = uint32(uintptr(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1308084)))).ChildByID(711))))
	dword_5d4594_1308140 = uint32(uintptr(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1308084)))).ChildByID(712))))
	dword_5d4594_1308144 = uint32(uintptr(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1308084)))).ChildByID(713))))
	dword_5d4594_1308148 = uint32(uintptr(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1308084)))).ChildByID(714))))
	dword_5d4594_1308152 = uint32(uintptr(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1308084)))).ChildByID(751))))
	result = int32(dword_587000_171388)
	if dword_587000_171388 == 0 {
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1308152))))).Func94(asWindowEvent(0x401E, *(*int32)(unsafe.Pointer(&dword_5d4594_1307784)), 0))
		sub_4A61E0(*(**uint32)(unsafe.Pointer(&dword_5d4594_1308096)), 2, (*uint8)(unsafe.Pointer(uintptr(dword_5d4594_1307784+71))))
		sub_4A61E0(*(**uint32)(unsafe.Pointer(&dword_5d4594_1308100)), 1, (*uint8)(unsafe.Pointer(uintptr(dword_5d4594_1307784+68))))
		sub_4A61E0(*(**uint32)(unsafe.Pointer(&dword_5d4594_1308104)), 1, (*uint8)(unsafe.Pointer(uintptr(dword_5d4594_1307784+74))))
		sub_4A61E0(*(**uint32)(unsafe.Pointer(&dword_5d4594_1308108)), 1, (*uint8)(unsafe.Pointer(uintptr(dword_5d4594_1307784+77))))
		sub_4A61E0(*(**uint32)(unsafe.Pointer(&dword_5d4594_1308112)), 1, (*uint8)(unsafe.Pointer(uintptr(dword_5d4594_1307784+80))))
		*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1308116 + 32))) = uint32(int32(*(*uint8)(unsafe.Pointer(uintptr(dword_5d4594_1307784 + 83)))) << 16)
		*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1308120 + 32))) = uint32(int32(*(*uint8)(unsafe.Pointer(uintptr(dword_5d4594_1307784 + 84)))) << 16)
		*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1308124 + 32))) = uint32(int32(*(*uint8)(unsafe.Pointer(uintptr(dword_5d4594_1307784 + 85)))) << 16)
		*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1308128 + 32))) = uint32(int32(*(*uint8)(unsafe.Pointer(uintptr(dword_5d4594_1307784 + 86)))) << 16)
		result = int32(*(*uint8)(unsafe.Pointer(uintptr(dword_5d4594_1307784 + 87)))) << 16
		*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1308132 + 32))) = uint32(result)
	}
	return result
}
func sub_4A61E0(a1 *uint32, a2 int32, a3 *uint8) *uint8 {
	var (
		v3     uint32
		result *uint8
		v5     *uint32
		v6     *uint32
	)
	v3 = 0
	result = (*uint8)(memmap.PtrOff(6112660, uint32(a2*96)+0x13F495))
	for int32(*((*uint8)(unsafe.Add(unsafe.Pointer(result), -1)))) != int32(*a3) || int32(*result) != int32(*(*uint8)(unsafe.Add(unsafe.Pointer(a3), 1))) || int32(*(*uint8)(unsafe.Add(unsafe.Pointer(result), 1))) != int32(*(*uint8)(unsafe.Add(unsafe.Pointer(a3), 2))) {
		v3++
		result = (*uint8)(unsafe.Add(unsafe.Pointer(result), 3))
		if v3 >= 32 {
			goto LABEL_9
		}
	}
	if a2 == 1 {
		v5 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1308084)))).ChildByID(int32(*a1 - 10))))
		nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v5)))))), 1)
		result = (*uint8)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1308084)))).ChildByID(int32(*a1 + 10))))
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(result))), unsafe.Sizeof(uint32(0))*9))) |= 6
	}
LABEL_9:
	if v3 == 32 && a2 == 1 {
		v6 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1308084)))).ChildByID(int32(*a1 - 10))))
		result = (*uint8)(unsafe.Pointer(uintptr(nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v6)))))), 0))))
		*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v3))), unsafe.Sizeof(uint16(0))*0)) = 9
	}
	*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*8)) = uint32(int32(uint16(int16(a2))) | int32(uint16(v3))<<16)
	return result
}
func sub_4A62B0() {
	var (
		v0 *uint8
		v1 int32
		v3 int32
		v5 *uint8
		v6 int32
		v7 [16]byte
	)
	v0 = (*uint8)(memmap.PtrOff(6112660, 1307796))
	v1 = 0
	v5 = (*uint8)(memmap.PtrOff(6112660, 1307796))
	for {
		{
			*v0 = 0
			v6 = v1 + 1
			*(*uint8)(unsafe.Add(unsafe.Pointer(v0), 1)) = 0
			*(*uint8)(unsafe.Add(unsafe.Pointer(v0), 2)) = 0
			nox_sprintf(&v7[0], CString("UserColor%d"), v1+1)
			var v2 *obj_412ae0_t = nox_xxx_modifGetModifListByType_4133B0(2)
			if v2 != nil {
				for libc.StrCmp(&v7[0], v2.field_0) != 0 {
					v2 = nox_xxx_modifNext_4133C0(v2)
					if v2 == nil {
						goto LABEL_7
					}
				}
				v3 = int32(uintptr(unsafe.Pointer(&v2.field_6)))
				*(*uint16)(unsafe.Pointer(v5)) = *(*uint16)(unsafe.Pointer(uintptr(v3)))
				*(*uint8)(unsafe.Add(unsafe.Pointer(v5), 2)) = *(*uint8)(unsafe.Pointer(uintptr(v3 + 2)))
			LABEL_7:
				v0 = v5
			}
			v1 = v6
			v0 = (*uint8)(unsafe.Add(unsafe.Pointer(v0), 3))
			v5 = v0
		}
		if int32(uintptr(unsafe.Pointer(v0))) >= int32(uintptr(memmap.PtrOff(6112660, 1307892))) {
			break
		}
	}
	*memmap.PtrUint8(6112660, 1307989) = 222
	*memmap.PtrUint8(6112660, 1307991) = 222
	*memmap.PtrUint8(6112660, 1307995) = 154
	*memmap.PtrUint8(6112660, 1308000) = 154
	*memmap.PtrUint8(6112660, 1308038) = 154
	*memmap.PtrUint8(6112660, 1307988) = 233
	*memmap.PtrUint8(6112660, 1307990) = 205
	*memmap.PtrUint8(6112660, 1307992) = 182
	*memmap.PtrUint8(6112660, 1307993) = 168
	*memmap.PtrUint8(6112660, 1307994) = 218
	*memmap.PtrUint8(6112660, 1307996) = 110
	*memmap.PtrUint8(6112660, 1307997) = 236
	*memmap.PtrUint8(6112660, 1307998) = 202
	*memmap.PtrUint8(6112660, 1307999) = 178
	*memmap.PtrUint8(6112660, 1308001) = 131
	*memmap.PtrUint8(6112660, 1308002) = 105
	*memmap.PtrUint8(6112660, 1308003) = 117
	*memmap.PtrUint8(6112660, 1308004) = 95
	*memmap.PtrUint8(6112660, 1308005) = 74
	*memmap.PtrUint8(6112660, 1308006) = 239
	*memmap.PtrUint8(6112660, 1308007) = 233
	*memmap.PtrUint8(6112660, 1308008) = 193
	*memmap.PtrUint8(6112660, 1308009) = 228
	*memmap.PtrUint8(6112660, 1308010) = 228
	*memmap.PtrUint8(6112660, 1308011) = 193
	*memmap.PtrUint8(6112660, 1308012) = 203
	*memmap.PtrUint8(6112660, 1308013) = 185
	*memmap.PtrUint8(6112660, 1308014) = 156
	*memmap.PtrUint8(6112660, 1308015) = 200
	*memmap.PtrUint8(6112660, 1308016) = 164
	*memmap.PtrUint8(6112660, 1308017) = 110
	*memmap.PtrUint8(6112660, 1308018) = 169
	*memmap.PtrUint8(6112660, 1308019) = 131
	*memmap.PtrUint8(6112660, 1308020) = 79
	*memmap.PtrUint8(6112660, 1308021) = 243
	*memmap.PtrUint8(6112660, 1308022) = 158
	*memmap.PtrUint8(6112660, 1308023) = 119
	*memmap.PtrUint8(6112660, 1308024) = 115
	*memmap.PtrUint8(6112660, 1308025) = 77
	*memmap.PtrUint8(6112660, 1308026) = 34
	*memmap.PtrUint8(6112660, 1308027) = 91
	*memmap.PtrUint8(6112660, 1308028) = 55
	*memmap.PtrUint8(6112660, 1308029) = 20
	*memmap.PtrUint8(6112660, 1308030) = 249
	*memmap.PtrUint8(6112660, 1308031) = 205
	*memmap.PtrUint8(6112660, 1308032) = 138
	*memmap.PtrUint8(6112660, 1308033) = 240
	*memmap.PtrUint8(6112660, 1308034) = 235
	*memmap.PtrUint8(6112660, 1308035) = 171
	*memmap.PtrUint8(6112660, 1308036) = 134
	*memmap.PtrUint8(6112660, 1308037) = 143
	*memmap.PtrUint8(6112660, 1308039) = 135
	*memmap.PtrUint8(6112660, 1308040) = 102
	*memmap.PtrUint8(6112660, 1308041) = 67
	*memmap.PtrUint8(6112660, 1308042) = 115
	*memmap.PtrUint8(6112660, 1308043) = 80
	*memmap.PtrUint8(6112660, 1308044) = 46
	*memmap.PtrUint8(6112660, 1308045) = 165
	*memmap.PtrUint8(6112660, 1308046) = 93
	*memmap.PtrUint8(6112660, 1308047) = 70
	*memmap.PtrUint8(6112660, 1308048) = 99
	*memmap.PtrUint8(6112660, 1308049) = 65
	*memmap.PtrUint8(6112660, 1308050) = 37
	*memmap.PtrUint8(6112660, 1308051) = 83
	*memmap.PtrUint8(6112660, 1308052) = 52
	*memmap.PtrUint8(6112660, 1308053) = 29
	*memmap.PtrUint8(6112660, 1308054) = 168
	*memmap.PtrUint8(6112660, 1308055) = math.MaxInt8
	*memmap.PtrUint8(6112660, 1308056) = 84
	*memmap.PtrUint8(6112660, 1308057) = 171
	*memmap.PtrUint8(6112660, 1308058) = 160
	*memmap.PtrUint8(6112660, 1308059) = 96
	*memmap.PtrUint8(6112660, 1308060) = 141
	*memmap.PtrUint8(6112660, 1308061) = 136
	*memmap.PtrUint8(6112660, 1308062) = 118
	*memmap.PtrUint8(6112660, 1308063) = math.MaxInt8
	*memmap.PtrUint8(6112660, 1308065) = 82
	*memmap.PtrUint8(6112660, 1308064) = 116
	*memmap.PtrUint8(6112660, 1308080) = 116
	*memmap.PtrUint8(6112660, 1308082) = 116
	*memmap.PtrUint8(6112660, 1307898) = math.MaxUint8
	*memmap.PtrUint8(6112660, 1307904) = math.MaxUint8
	*memmap.PtrUint8(6112660, 1308066) = 96
	*memmap.PtrUint8(6112660, 1307908) = 68
	*memmap.PtrUint8(6112660, 1307909) = 68
	*memmap.PtrUint8(6112660, 1307926) = 68
	*memmap.PtrUint8(6112660, 1307936) = 68
	*memmap.PtrUint8(6112660, 1308067) = 84
	*memmap.PtrUint8(6112660, 1307937) = 214
	*memmap.PtrUint8(6112660, 1307938) = 214
	*memmap.PtrUint8(6112660, 1307939) = 214
	*memmap.PtrUint8(6112660, 1308068) = 51
	*memmap.PtrUint8(6112660, 1308069) = 162
	*memmap.PtrUint8(6112660, 1308070) = 129
	*memmap.PtrUint8(6112660, 1308071) = 100
	*memmap.PtrUint8(6112660, 1308072) = 78
	*memmap.PtrUint8(6112660, 1308073) = 64
	*memmap.PtrUint8(6112660, 1308074) = 46
	*memmap.PtrUint8(6112660, 1308075) = 65
	*memmap.PtrUint8(6112660, 1308076) = 50
	*memmap.PtrUint8(6112660, 1308077) = 45
	*memmap.PtrUint8(6112660, 1308078) = 171
	*memmap.PtrUint8(6112660, 1308079) = 158
	*memmap.PtrUint8(6112660, 1308081) = 175
	*memmap.PtrUint8(6112660, 1308083) = 131
	*memmap.PtrUint8(6112660, 1307892) = 243
	*memmap.PtrUint8(6112660, 1307893) = 183
	*memmap.PtrUint8(6112660, 1307894) = 159
	*memmap.PtrUint8(6112660, 1307895) = 199
	*memmap.PtrUint8(6112660, 1307896) = 132
	*memmap.PtrUint8(6112660, 1307897) = 58
	*memmap.PtrUint8(6112660, 1307899) = 170
	*memmap.PtrUint8(6112660, 1307900) = 86
	*memmap.PtrUint8(6112660, 1307901) = 146
	*memmap.PtrUint8(6112660, 1307902) = 121
	*memmap.PtrUint8(6112660, 1307903) = 106
	*memmap.PtrUint8(6112660, 1307905) = 251
	*memmap.PtrUint8(6112660, 1307906) = 186
	*memmap.PtrUint8(6112660, 1307907) = 247
	*memmap.PtrUint8(6112660, 1307910) = 63
	*memmap.PtrUint8(6112660, 1307911) = 187
	*memmap.PtrUint8(6112660, 1307912) = 239
	*memmap.PtrUint8(6112660, 1307913) = 196
	*memmap.PtrUint8(6112660, 1307914) = 194
	*memmap.PtrUint8(6112660, 1307915) = 177
	*memmap.PtrUint8(6112660, 1307916) = 253
	*memmap.PtrUint8(6112660, 1307917) = 246
	*memmap.PtrUint8(6112660, 1307918) = 102
	*memmap.PtrUint8(6112660, 1307919) = 115
	*memmap.PtrUint8(6112660, 1307920) = 77
	*memmap.PtrUint8(6112660, 1307921) = 34
	*memmap.PtrUint8(6112660, 1307922) = 229
	*memmap.PtrUint8(6112660, 1307923) = 121
	*memmap.PtrUint8(6112660, 1307924) = 25
	*memmap.PtrUint8(6112660, 1307925) = 82
	*memmap.PtrUint8(6112660, 1307927) = 60
	*memmap.PtrUint8(6112660, 1307928) = 235
	*memmap.PtrUint8(6112660, 1307929) = 229
	*memmap.PtrUint8(6112660, 1307930) = 137
	*memmap.PtrUint8(6112660, 1307931) = 147
	*memmap.PtrUint8(6112660, 1307932) = 0
	*memmap.PtrUint8(6112660, 1307933) = 0
	*memmap.PtrUint8(6112660, 1307934) = 0
	*memmap.PtrUint8(6112660, 1307935) = 37
	*memmap.PtrUint8(6112660, 1307940) = 149
	*memmap.PtrUint8(6112660, 1307941) = 108
	*memmap.PtrUint8(6112660, 1307942) = 43
	*memmap.PtrUint8(6112660, 1307943) = 52
	*memmap.PtrUint8(6112660, 1307944) = 25
	*memmap.PtrUint8(6112660, 1307945) = 0
	*memmap.PtrUint8(6112660, 1307946) = 150
	*memmap.PtrUint8(6112660, 1307947) = 55
	*memmap.PtrUint8(6112660, 1307952) = 150
	*memmap.PtrUint8(6112660, 1307955) = 131
	*memmap.PtrUint8(6112660, 1307956) = 104
	*memmap.PtrUint8(6112660, 1307961) = 117
	*memmap.PtrUint8(6112660, 1307962) = 117
	*memmap.PtrUint8(6112660, 1307963) = 117
	*memmap.PtrUint8(6112660, 1307981) = 104
	*memmap.PtrUint8(6112660, 1307948) = 0
	*memmap.PtrUint8(6112660, 1307950) = 0
	*memmap.PtrUint8(6112660, 1307984) = 0
	*memmap.PtrUint32(0x587000, 171384) = 0
	*memmap.PtrUint8(6112660, 1307949) = 9
	*memmap.PtrUint8(6112660, 1307951) = 12
	*memmap.PtrUint8(6112660, 1307953) = 146
	*memmap.PtrUint8(6112660, 1307954) = 81
	*memmap.PtrUint8(6112660, 1307957) = 179
	*memmap.PtrUint8(6112660, 1307958) = 121
	*memmap.PtrUint8(6112660, 1307959) = 165
	*memmap.PtrUint8(6112660, 1307960) = 66
	*memmap.PtrUint8(6112660, 1307964) = 100
	*memmap.PtrUint8(6112660, 1307965) = 130
	*memmap.PtrUint8(6112660, 1307966) = 125
	*memmap.PtrUint8(6112660, 1307967) = 56
	*memmap.PtrUint8(6112660, 1307968) = 48
	*memmap.PtrUint8(6112660, 1307969) = 40
	*memmap.PtrUint8(6112660, 1307970) = 108
	*memmap.PtrUint8(6112660, 1307971) = 81
	*memmap.PtrUint8(6112660, 1307972) = 58
	*memmap.PtrUint8(6112660, 1307973) = 52
	*memmap.PtrUint8(6112660, 1307974) = 40
	*memmap.PtrUint8(6112660, 1307975) = 39
	*memmap.PtrUint8(6112660, 1307976) = 143
	*memmap.PtrUint8(6112660, 1307977) = 142
	*memmap.PtrUint8(6112660, 1307978) = 124
	*memmap.PtrUint8(6112660, 1307979) = 56
	*memmap.PtrUint8(6112660, 1307980) = 29
	*memmap.PtrUint8(6112660, 1307982) = 15
	*memmap.PtrUint8(6112660, 1307983) = 88
	*memmap.PtrUint8(6112660, 1307985) = 87
	*memmap.PtrUint8(6112660, 1307986) = 53
	*memmap.PtrUint8(6112660, 1307987) = 53
}
func sub_4A6890() int32 {
	nox_wnd_xxx_1308092.state = nox_gui_anim_state(NOX_GUI_ANIM_OUT)
	sub_43BE40(2)
	clientPlaySoundSpecial(923, 100)
	sub_4A68C0()
	return 1
}
func sub_4A6B50(a1 *libc.WChar) int32 {
	var (
		v1  *libc.WChar
		v2  int32
		v3  int32
		v4  int32
		v5  *int16
		v6  libc.WChar
		v7  int32
		i   *libc.WChar
		v10 int32
		v11 int32
		v12 [26]int16
	)
	v1 = a1
	v2 = 0
	v3 = 1
	v10 = 0
	v4 = int32(nox_wcslen(a1))
	if v4 >= 1 {
		v5 = &v12[0]
		v11 = v4
		for {
			if unicode.IsSpace(rune(*v1)) {
				if v3 == 0 {
					*v5 = int16(*v1)
					v5 = (*int16)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(int16(0))*1))
					v10++
				}
			} else {
				if v3 == 1 {
					v6 = *v1
					if *v1 == 42 || v6 == 63 || v6 == 60 || v6 == 62 || v6 == 92 || v6 == 47 || v6 == 58 || v6 == 34 || v6 == 124 {
						*v1 = 45
					}
				}
				*v5 = int16(*v1)
				v5 = (*int16)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(int16(0))*1))
				v10++
				v3 = 0
				v2 = 1
			}
			v1 = (*libc.WChar)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(libc.WChar(0))*1))
			v11--
			if v11 == 0 {
				break
			}
		}
		v12[v10] = 0
		if v2 != 0 {
			nox_wcscpy(a1, (*libc.WChar)(unsafe.Pointer(&v12[0])))
			v7 = int32(nox_wcslen((*libc.WChar)(unsafe.Pointer(&v12[0]))) - 1)
			if v7 >= 0 {
				for i = (*libc.WChar)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(libc.WChar(0))*uintptr(v7))); unicode.IsSpace(rune(*i)); i = (*libc.WChar)(unsafe.Add(unsafe.Pointer(i), -int(unsafe.Sizeof(libc.WChar(0))*1))) {
					if func() int32 {
						p := &v7
						*p--
						return *p
					}() < 0 {
						return v2
					}
				}
				*(*libc.WChar)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(libc.WChar(0))*uintptr(v7+1))) = 0
			}
		}
	}
	return v2
}
func sub_4A6C90() int32 {
	var v0 func() int32
	v0 = nox_wnd_xxx_1308092.field_13
	nox_gui_freeAnimation_43C570(nox_wnd_xxx_1308092)
	(*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1308084)))).Destroy()
	(*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1308088)))).Destroy()
	if v0 != nil {
		v0()
	} else {
		nox_client_resetScreenParticles_431510()
		nox_gui_draw()
		if !noxflags.HasGame(noxflags.GameOnline) {
			nox_client_guiXxxDestroy_4A24A0()
			return 1
		}
		if nox_xxx_check_flag_aaa_43AF70() == 0 {
			nox_xxx_serverHost_43B4D0()
			return 1
		}
		if nox_xxx_check_flag_aaa_43AF70() == 1 {
			sub_43B670()
			return 1
		}
	}
	return 1
}
func sub_4A6D20(a1 int32, a2 int32) int32 {
	var (
		v1    int32
		v2    uint16
		v3    int32
		v4    int32
		v5    int32
		v6    int32
		v7    int32
		v8    int32
		v10   int32
		yTop  int32
		xLeft int32
		v13   int32
	)
	v1 = a1
	v2 = *(*uint16)(unsafe.Pointer(uintptr(a1 + 32)))
	v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 32))) >> 16)
	nox_client_wndGetPosition_46AA60((*nox_window)(unsafe.Pointer(uintptr(a1))), (*uint32)(unsafe.Pointer(&xLeft)), (*uint32)(unsafe.Pointer(&yTop)))
	v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 20))))
	v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 16))))
	v6 = int32(uint16(int16(v3))) + int32(v2)*32
	v13 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 28))) - uint32(v4))
	v7 = int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 24))) - uint32(v5))
	*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v5))), 0)) = *memmap.PtrUint8(6112660, uint32(v6*3)+0x13F496)
	*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v4))), 0)) = *memmap.PtrUint8(6112660, uint32(v6*3)+0x13F495)
	v10 = v5
	*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v5))), 0)) = *memmap.PtrUint8(6112660, uint32(v6*3)+0x13F494)
	v8 = int32(nox_color_rgb_4344A0(v5, v4, v10))
	nox_client_drawSetColor_434460(v8)
	nox_client_drawRectFilledOpaque_49CE30(xLeft, yTop, v7, v13)
	return 1
}
func sub_4A6DC0(a1 *uint32, a2 int32) int32 {
	var (
		v1  int32
		v2  int32
		v3  int32
		v4  int32
		v5  int32
		v6  int32
		v7  int32
		v8  int32
		v9  int32
		v10 int32
		v11 int32
		v12 int32
		v13 int32
		v14 int32
		v15 int32
		v16 int32
		v17 int32
		v18 int32
		v19 int32
		v20 int32
		v21 int32
		v22 int32
		v23 int32
		v24 *int32
		v25 int32
		v26 int32
		v27 int32
		i   int32
		v29 int32
		v30 int32
		v31 int32
		v32 int32
		v33 int32
		j   int32
		v35 int32
		v36 int32
		v37 int32
		v38 int32
		v39 int32
		v40 int32
		v41 int32
		v42 int32
		k   int32
		v44 int32
		v45 int32
		v46 int32
		v47 int32
		v48 int32
		v49 int32
		v50 int32
		v52 int32
		v53 int32
		v54 int32
		v55 int32
	)
	nox_client_wndGetPosition_46AA60((*nox_window)(unsafe.Pointer(a1)), (*uint32)(unsafe.Pointer(&v54)), (*uint32)(unsafe.Pointer(&v55)))
	v1 = int32(uint16(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1308096 + 32)))>>16)) + int32(*(*uint16)(unsafe.Pointer(uintptr(dword_5d4594_1308096 + 32))))*32
	v2 = int32(*memmap.PtrUint8(6112660, uint32(v1*3)+0x13F496))
	*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v3))), 0)) = *memmap.PtrUint8(6112660, uint32(v1*3)+0x13F495)
	*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v1))), 0)) = *memmap.PtrUint8(6112660, uint32(v1*3)+0x13F494)
	v4 = int32(nox_color_rgb_4344A0(v1, v3, v2))
	nox_xxx_drawPlayer_4341D0(6, v4)
	v5 = int32(uint16(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1308100 + 32)))>>16)) + int32(*(*uint16)(unsafe.Pointer(uintptr(dword_5d4594_1308100 + 32))))*32
	v6 = int32(*memmap.PtrUint8(6112660, uint32(v5*3)+0x13F496))
	*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v7))), 0)) = *memmap.PtrUint8(6112660, uint32(v5*3)+0x13F495)
	*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v5))), 0)) = *memmap.PtrUint8(6112660, uint32(v5*3)+0x13F494)
	v8 = int32(nox_color_rgb_4344A0(v5, v7, v6))
	nox_xxx_drawPlayer_4341D0(1, v8)
	v9 = int32(dword_5d4594_1308144)
	if int32(*(*uint8)(unsafe.Pointer(uintptr(dword_5d4594_1308144 + 4))))&8 != 0 {
		v10 = int32(uint16(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1308108 + 32))) >> 16))
		v11 = v10 + int32(*(*uint16)(unsafe.Pointer(uintptr(dword_5d4594_1308108 + 32))))*32
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v9))), 0)) = *memmap.PtrUint8(6112660, uint32(v11*3)+0x13F495)
		v12 = int32(nox_color_rgb_4344A0(int32(*memmap.PtrUint8(6112660, uint32(v11*3)+0x13F494)), v9, int32(*memmap.PtrUint8(6112660, uint32((v10+int32(*(*uint16)(unsafe.Pointer(uintptr(dword_5d4594_1308108 + 32))))*32)*3)+0x13F496))))
		nox_xxx_drawPlayer_4341D0(2, v12)
	} else {
		nox_xxx_drawPlayer_4341D0(2, v4)
	}
	nox_xxx_drawPlayer_4341D0(3, v4)
	v13 = int32(dword_5d4594_1308148)
	if int32(*(*uint8)(unsafe.Pointer(uintptr(dword_5d4594_1308148 + 4))))&8 != 0 {
		v14 = int32(uint16(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1308112 + 32))) >> 16))
		v15 = v14 + int32(*(*uint16)(unsafe.Pointer(uintptr(dword_5d4594_1308112 + 32))))*32
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v13))), 0)) = *memmap.PtrUint8(6112660, uint32(v15*3)+0x13F495)
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v15))), 0)) = *memmap.PtrUint8(6112660, uint32(v15*3)+0x13F494)
		v16 = int32(nox_color_rgb_4344A0(v15, v13, int32(*memmap.PtrUint8(6112660, uint32((v14+int32(*(*uint16)(unsafe.Pointer(uintptr(dword_5d4594_1308112 + 32))))*32)*3)+0x13F496))))
		nox_xxx_drawPlayer_4341D0(4, v16)
	} else {
		nox_xxx_drawPlayer_4341D0(4, v4)
	}
	v17 = int32(dword_5d4594_1308140)
	if int32(*(*uint8)(unsafe.Pointer(uintptr(dword_5d4594_1308140 + 4))))&8 != 0 {
		v18 = int32(uint16(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1308104 + 32))) >> 16))
		v19 = v18 + int32(*(*uint16)(unsafe.Pointer(uintptr(dword_5d4594_1308104 + 32))))*32
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v17))), 0)) = *memmap.PtrUint8(6112660, uint32(v19*3)+0x13F495)
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v19))), 0)) = *memmap.PtrUint8(6112660, uint32(v19*3)+0x13F494)
		v20 = int32(nox_color_rgb_4344A0(v19, v17, int32(*memmap.PtrUint8(6112660, uint32((v18+int32(*(*uint16)(unsafe.Pointer(uintptr(dword_5d4594_1308104 + 32))))*32)*3)+0x13F496))))
		nox_xxx_drawPlayer_4341D0(5, v20)
	} else {
		nox_xxx_drawPlayer_4341D0(5, v4)
	}
	if int32(*(*uint8)(unsafe.Pointer(uintptr(dword_5d4594_1308136 + 4))))&8 != 0 {
		nox_client_drawImageAt_47D2C0((*nox_video_bag_image_t)(unsafe.Pointer(uintptr(*memmap.PtrUint32(0x973A20, uint32(int32(*(*uint8)(unsafe.Pointer(uintptr(dword_5d4594_1307784 + 67))))*4+16))))), v54, v55)
	} else {
		nox_client_drawImageAt_47D2C0((*nox_video_bag_image_t)(unsafe.Pointer(uintptr(*memmap.PtrUint32(0x973A20, uint32(int32(*(*uint8)(unsafe.Pointer(uintptr(dword_5d4594_1307784 + 67))))*4+24))))), v54, v55)
	}
	v21 = 0
	v22 = int32(*(*uint8)(unsafe.Pointer(uintptr(dword_5d4594_1307784 + 67))))
	v23 = v22 * 3
	v24 = memmap.PtrInt32(0x973A20, uint32(v22*104+32))
	for {
		v25 = v21
		v26 = 1 << v21
		if 1<<v21 == 4 {
			v27 = 1
			for i = 3; i < 21; i += 3 {
				v29 = int32(dword_5d4594_1308156)
				*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v25))), 0)) = *(*uint8)(unsafe.Pointer(uintptr(uint32(i) + dword_5d4594_1308156 + 14)))
				*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v23))), 0)) = *(*uint8)(unsafe.Pointer(uintptr(uint32(i) + dword_5d4594_1308156 + 13)))
				*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v29))), 0)) = *(*uint8)(unsafe.Pointer(uintptr(uint32(i) + dword_5d4594_1308156 + 12)))
				sub_4340A0(func() int32 {
					p := &v27
					x := *p
					*p++
					return x
				}(), v29, v23, v25)
			}
			v30 = int32(uint16(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1308116 + 32)))>>16)) + int32(*(*uint16)(unsafe.Pointer(uintptr(dword_5d4594_1308116 + 32))))*32
			v31 = int32(*memmap.PtrUint8(6112660, uint32(v30*3)+0x13F496))
			*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v23))), 0)) = *memmap.PtrUint8(6112660, uint32(v30*3)+0x13F495)
			*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v30))), 0)) = *memmap.PtrUint8(6112660, uint32(v30*3)+0x13F494)
			v32 = int32(nox_color_rgb_4344A0(v30, v23, v31))
			nox_xxx_drawPlayer_4341D0(int32(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1308156 + 40)))), v32)
		} else if v26 == 1024 {
			v33 = 1
			for j = 3; j < 21; j += 3 {
				v35 = int32(dword_5d4594_1308160)
				*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v25))), 0)) = *(*uint8)(unsafe.Pointer(uintptr(dword_5d4594_1308160 + uint32(j) + 14)))
				*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v23))), 0)) = *(*uint8)(unsafe.Pointer(uintptr(dword_5d4594_1308160 + uint32(j) + 13)))
				*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v35))), 0)) = *(*uint8)(unsafe.Pointer(uintptr(dword_5d4594_1308160 + uint32(j) + 12)))
				sub_4340A0(func() int32 {
					p := &v33
					x := *p
					*p++
					return x
				}(), v35, v23, v25)
			}
			v36 = int32(uint16(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1308120 + 32)))>>16)) + int32(*(*uint16)(unsafe.Pointer(uintptr(dword_5d4594_1308120 + 32))))*32
			v37 = int32(*memmap.PtrUint8(6112660, uint32(v36*3)+0x13F496))
			*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v23))), 0)) = *memmap.PtrUint8(6112660, uint32(v36*3)+0x13F495)
			*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v36))), 0)) = *memmap.PtrUint8(6112660, uint32(v36*3)+0x13F494)
			v38 = int32(nox_color_rgb_4344A0(v36, v23, v37))
			nox_xxx_drawPlayer_4341D0(int32(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1308160 + 40)))), v38)
			v39 = int32(uint16(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1308124 + 32)))>>16)) + int32(*(*uint16)(unsafe.Pointer(uintptr(dword_5d4594_1308124 + 32))))*32
			v40 = int32(*memmap.PtrUint8(6112660, uint32(v39*3)+0x13F496))
			*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v41))), 0)) = *memmap.PtrUint8(6112660, uint32(v39*3)+0x13F495)
			*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v39))), 0)) = *memmap.PtrUint8(6112660, uint32(v39*3)+0x13F494)
			v52 = int32(nox_color_rgb_4344A0(v39, v41, v40))
			nox_xxx_drawPlayer_4341D0(int32(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1308160 + 44)))), v52)
		} else {
			if v26 != 1 {
				goto LABEL_27
			}
			v42 = 1
			for k = 3; k < 21; k += 3 {
				v44 = int32(dword_5d4594_1308164)
				*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v25))), 0)) = *(*uint8)(unsafe.Pointer(uintptr(uint32(k) + dword_5d4594_1308164 + 14)))
				*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v23))), 0)) = *(*uint8)(unsafe.Pointer(uintptr(uint32(k) + dword_5d4594_1308164 + 13)))
				*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v44))), 0)) = *(*uint8)(unsafe.Pointer(uintptr(uint32(k) + dword_5d4594_1308164 + 12)))
				sub_4340A0(func() int32 {
					p := &v42
					x := *p
					*p++
					return x
				}(), v44, v23, v25)
			}
			v45 = int32(uint16(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1308128 + 32)))>>16)) + int32(*(*uint16)(unsafe.Pointer(uintptr(dword_5d4594_1308128 + 32))))*32
			v46 = int32(*memmap.PtrUint8(6112660, uint32(v45*3)+0x13F496))
			*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v23))), 0)) = *memmap.PtrUint8(6112660, uint32(v45*3)+0x13F495)
			*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v45))), 0)) = *memmap.PtrUint8(6112660, uint32(v45*3)+0x13F494)
			v47 = int32(nox_color_rgb_4344A0(v45, v23, v46))
			nox_xxx_drawPlayer_4341D0(int32(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1308164 + 40)))), v47)
			v48 = int32(uint16(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1308132 + 32)))>>16)) + int32(*(*uint16)(unsafe.Pointer(uintptr(dword_5d4594_1308132 + 32))))*32
			v49 = int32(*memmap.PtrUint8(6112660, uint32(v48*3)+0x13F496))
			*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v50))), 0)) = *memmap.PtrUint8(6112660, uint32(v48*3)+0x13F495)
			*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v48))), 0)) = *memmap.PtrUint8(6112660, uint32(v48*3)+0x13F494)
			v53 = int32(nox_color_rgb_4344A0(v48, v50, v49))
			nox_xxx_drawPlayer_4341D0(int32(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1308164 + 36)))), v53)
		}
		nox_client_drawImageAt_47D2C0((*nox_video_bag_image_t)(unsafe.Pointer(uintptr(*v24))), v54, v55)
	LABEL_27:
		v21++
		v24 = (*int32)(unsafe.Add(unsafe.Pointer(v24), unsafe.Sizeof(int32(0))*1))
		if v21 >= 26 {
			break
		}
	}
	return 1
}
func sub_4A7270(a1 int32, a2 int32, a3 uint32, a4 int32) int32 {
	var (
		v3 int32
		v5 int2
	)
	if a2 != 5 {
		return 0
	}
	v5.field_4 = int32(a3 >> 16)
	v5.field_0 = int32(uint16(a3))
	v3 = sub_4281F0(&v5, (*int4)(unsafe.Pointer(uintptr(a1+16))))
	if v3 != 0 {
		return 0
	}
	sub_4A72D0(0xDEAD)
	return 1
}
func sub_4A72D0(a1 uint16) *uint32 {
	var result *uint32
	nox_xxx_wnd_46C6E0((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1308088))))))
	result = (*uint32)(unsafe.Pointer(uintptr((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1308088))))).Hide())))
	if int32(a1) < 32 {
		result = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1308088)))).ChildByID(*(*int32)(unsafe.Pointer(&dword_5d4594_1307792)))))
		if result != nil {
			*(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*8)) = uint32(int32(*memmap.PtrUint16(6112660, 1307788)) | int32(a1)<<16)
		}
	}
	return result
}
func sub_4A7330(a1 int32, a2 int32, a3 *int32, a4 uint32) int32 {
	var (
		result int32
		v5     int32
		v6     int32
		v7     *uint32
	)
	if a2 == 0x4005 {
		clientPlaySoundSpecial(920, 100)
		result = 1
	} else if a2 == 0x4007 {
		v5 = int32(a4 >> 16)
		v6 = (*nox_window)(unsafe.Pointer(a3)).ID()
		switch v6 {
		case 720:
			dword_5d4594_1307792 = uint32(v6)
			sub_4A7530(2)
			sub_4A7580(int32(uint16(a4)), v5)
		case 721:
			fallthrough
		case 722:
			fallthrough
		case 723:
			fallthrough
		case 724:
			dword_5d4594_1307792 = uint32(v6)
			sub_4A7530(1)
			sub_4A7580(int32(uint16(a4)), v5)
		case 725:
			fallthrough
		case 726:
			fallthrough
		case 727:
			fallthrough
		case 728:
			fallthrough
		case 729:
			dword_5d4594_1307792 = uint32(v6)
			sub_4A7530(0)
			sub_4A7580(int32(uint16(a4)), v5)
		case 731:
			fallthrough
		case 732:
			fallthrough
		case 733:
			fallthrough
		case 734:
			v7 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1308084)))).ChildByID(v6 - 20)))
			if v7 != nil {
				nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v7)))))), int32((^*(*uint32)(unsafe.Add(unsafe.Pointer(v7), unsafe.Sizeof(uint32(0))*1))>>3)&1))
			}
		case 761:
			fallthrough
		case 762:
			fallthrough
		case 763:
			fallthrough
		case 764:
			fallthrough
		case 765:
			fallthrough
		case 766:
			fallthrough
		case 767:
			fallthrough
		case 768:
			fallthrough
		case 769:
			fallthrough
		case 770:
			fallthrough
		case 771:
			fallthrough
		case 772:
			fallthrough
		case 773:
			fallthrough
		case 774:
			fallthrough
		case 775:
			fallthrough
		case 776:
			fallthrough
		case 777:
			fallthrough
		case 778:
			fallthrough
		case 779:
			fallthrough
		case 780:
			fallthrough
		case 781:
			fallthrough
		case 782:
			fallthrough
		case 783:
			fallthrough
		case 784:
			fallthrough
		case 785:
			fallthrough
		case 786:
			fallthrough
		case 787:
			fallthrough
		case 788:
			fallthrough
		case 789:
			fallthrough
		case 790:
			fallthrough
		case 791:
			fallthrough
		case 792:
			sub_4A72D0(uint16(int16(v6 - 761)))
		case 799:
			if *memmap.PtrUint32(6112660, 1308168) == 1 {
				nox_game_decStateInd_43BDC0()
			}
			nox_game_decStateInd_43BDC0()
			nox_game_decStateInd_43BDC0()
			dword_587000_171388 = 1
			if sub_4A75C0() != 0 {
				if int32(*(*uint8)(unsafe.Pointer(uintptr(dword_5d4594_1307784 + 66)))) == 0 {
					nox_xxx_gameSetMapPath_409D70(CString("war01a.map"))
				} else if int32(*(*uint8)(unsafe.Pointer(uintptr(dword_5d4594_1307784 + 66)))) == 1 {
					nox_xxx_gameSetMapPath_409D70(CString("wiz01a.map"))
				} else if int32(*(*uint8)(unsafe.Pointer(uintptr(dword_5d4594_1307784 + 66)))) == 2 {
					nox_xxx_gameSetMapPath_409D70(CString("con01a.map"))
				}
				sub_4A24C0(0)
				sub_4A6890()
				nox_wnd_xxx_1308092.field_13 = nil
			}
		default:
		}
		clientPlaySoundSpecial(921, 100)
		result = 1
	} else {
		result = 0
	}
	return result
}
func sub_4A7530(a1 uint16) *uint32 {
	var (
		v1     int32
		result *uint32
	)
	v1 = 761
	*memmap.PtrUint32(6112660, 1307788) = uint32(a1)
	for {
		result = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1308088)))).ChildByID(v1)))
		if result != nil {
			*(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*8)) = uint32(int32(a1) | int32(uint16(int16(v1-761)))<<16)
		}
		v1++
		if v1 > 792 {
			break
		}
	}
	return result
}
func sub_4A7580(a1 int32, a2 int32) int32 {
	nox_xxx_wndShowModalMB((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1308088))))))
	sub_46C690((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1308088))))))
	return (*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1308088)))).SetPos(image.Pt(int32(uint32(a1)-*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1308088 + 8)))), int32(uint32(a2)-*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1308088 + 12)))/2)))
}
func sub_4A7A60(a1 int32) int32 {
	var result int32
	result = a1
	dword_587000_171388 = uint32(a1)
	return result
}
func sub_4A7A70(a1 int32) int32 {
	var result int32
	result = a1
	*memmap.PtrUint32(6112660, 1308168) = uint32(a1)
	return result
}
func sub_4A7A80(a1 *byte) int32 {
	var (
		result int32
		v2     uint32
		v3     int8
		v4     *uint8
		v5     *byte
	)
	if a1 == nil {
		return 0
	}
	v2 = uint32(libc.StrLen(a1) + 1)
	v3 = int8(uint8(v2))
	v2 >>= 2
	libc.MemCpy(memmap.PtrOff(6112660, 1308644), unsafe.Pointer(a1), int(v2*4))
	v5 = (*byte)(unsafe.Add(unsafe.Pointer(a1), v2*4))
	v4 = (*uint8)(memmap.PtrOff(6112660, v2*4+0x13F7E4))
	*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v2))), 0)) = uint8(v3)
	result = 1
	libc.MemCpy(unsafe.Pointer(v4), unsafe.Pointer(v5), int(v2&3))
	return result
}
func sub_4A7AC0(a1 *byte) int32 {
	var (
		result int32
		v2     uint32
		v3     int8
		v4     *uint8
		v5     *byte
	)
	if a1 == nil {
		return 0
	}
	v2 = uint32(libc.StrLen(a1) + 1)
	v3 = int8(uint8(v2))
	v2 >>= 2
	libc.MemCpy(memmap.PtrOff(6112660, 1308172), unsafe.Pointer(a1), int(v2*4))
	v5 = (*byte)(unsafe.Add(unsafe.Pointer(a1), v2*4))
	v4 = (*uint8)(memmap.PtrOff(6112660, v2*4+0x13F60C))
	*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v2))), 0)) = uint8(v3)
	result = 1
	libc.MemCpy(unsafe.Pointer(v4), unsafe.Pointer(v5), int(v2&3))
	return result
}
func sub_4A7B00(a1 *byte) int32 {
	var (
		result int32
		v2     uint32
		v3     int8
		v4     *uint8
		v5     *byte
	)
	if a1 == nil {
		return 0
	}
	v2 = uint32(libc.StrLen(a1) + 1)
	v3 = int8(uint8(v2))
	v2 >>= 2
	libc.MemCpy(memmap.PtrOff(6112660, 1308352), unsafe.Pointer(a1), int(v2*4))
	v5 = (*byte)(unsafe.Add(unsafe.Pointer(a1), v2*4))
	v4 = (*uint8)(memmap.PtrOff(6112660, v2*4+0x13F6C0))
	*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v2))), 0)) = uint8(v3)
	result = 1
	libc.MemCpy(unsafe.Pointer(v4), unsafe.Pointer(v5), int(v2&3))
	return result
}
func sub_4A7B40(a1 *byte) int32 {
	var (
		v2 *byte
		v3 int32
		v4 *uint8
	)
	if a1 == nil {
		return 0
	}
	v2 = *(**byte)(memmap.PtrOff(0x587000, 171856))
	v3 = 0
	if *memmap.PtrUint32(0x587000, 171856) != 0 {
		v4 = (*uint8)(memmap.PtrOff(0x587000, 171856))
		for libc.StrCaseCmp(v2, a1) != 0 {
			v2 = (*byte)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v4))), unsafe.Sizeof(uint32(0))*2))))))
			v4 = (*uint8)(unsafe.Add(unsafe.Pointer(v4), 8))
			v3++
			if v2 == nil {
				return 1
			}
		}
		*memmap.PtrUint32(6112660, 1308184) = *memmap.PtrUint32(0x587000, uint32(v3*8)+171860)
	}
	return 1
}
func sub_4A7BA0(a1 *byte) int32 {
	var result int32
	result = int32(uintptr(unsafe.Pointer(a1)))
	if a1 != nil {
		*memmap.PtrUint32(6112660, 1308740) = uint32(libc.Atoi(libc.GoString(a1)))
		result = 1
	}
	return result
}
func sub_4A7BC0(a1 *byte) int32 {
	var (
		result int32
		v2     uint32
		v3     int8
		v4     *uint8
		v5     *byte
	)
	if a1 == nil {
		return 0
	}
	v2 = uint32(libc.StrLen(a1) + 1)
	v3 = int8(uint8(v2))
	v2 >>= 2
	libc.MemCpy(memmap.PtrOff(6112660, 1308324), unsafe.Pointer(a1), int(v2*4))
	v5 = (*byte)(unsafe.Add(unsafe.Pointer(a1), v2*4))
	v4 = (*uint8)(memmap.PtrOff(6112660, v2*4+0x13F6A4))
	*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v2))), 0)) = uint8(v3)
	result = 1
	libc.MemCpy(unsafe.Pointer(v4), unsafe.Pointer(v5), int(v2&3))
	return result
}
func sub_4A7C00(a1 *byte) int32 {
	var (
		result int32
		v2     uint32
		v3     int8
		v4     *uint8
		v5     *byte
	)
	if a1 == nil {
		return 0
	}
	v2 = uint32(libc.StrLen(a1) + 1)
	v3 = int8(uint8(v2))
	v2 >>= 2
	libc.MemCpy(memmap.PtrOff(6112660, 1308364), unsafe.Pointer(a1), int(v2*4))
	v5 = (*byte)(unsafe.Add(unsafe.Pointer(a1), v2*4))
	v4 = (*uint8)(memmap.PtrOff(6112660, v2*4+0x13F6CC))
	*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v2))), 0)) = uint8(v3)
	result = 1
	libc.MemCpy(unsafe.Pointer(v4), unsafe.Pointer(v5), int(v2&3))
	return result
}
func sub_4A7C40(a1 *byte) int32 {
	var result int32
	result = int32(uintptr(unsafe.Pointer(a1)))
	if a1 != nil {
		*memmap.PtrUint32(6112660, 1308188) = uint32(libc.Atoi(libc.GoString(a1)))
		result = 1
	}
	return result
}
func sub_4A7C60(a1 *byte) int32 {
	var (
		v1     *byte
		v2     *byte
		result int32
	)
	if a1 == nil {
		goto LABEL_12
	}
	*memmap.PtrUint32(6112660, 1308736) = 0
	*memmap.PtrUint32(6112660, 1308732) = 0
	v1 = libc.StrTok(a1, CString(",\t\r\n"))
	if v1 != nil {
		*memmap.PtrUint32(6112660, 1308732) = uint32(libc.Atoi(libc.GoString(v1)))
	}
	v2 = libc.StrTok(nil, CString(" \t\r\n"))
	if v2 != nil {
		*memmap.PtrUint32(6112660, 1308736) = uint32(libc.Atoi(libc.GoString(v2)))
	}
	if *memmap.PtrUint32(6112660, 1308732) != 0 && *memmap.PtrUint32(6112660, 1308736) != 0 {
		result = 1
	} else {
	LABEL_12:
		result = 0
	}
	return result
}
func sub_4A7CE0(a1 *byte) int32 {
	var result int32
	result = int32(uintptr(unsafe.Pointer(a1)))
	if a1 != nil {
		*memmap.PtrUint32(6112660, 1308728) = uint32(libc.Atoi(libc.GoString(a1)))
		result = 1
	}
	return result
}
func sub_4A7D00(a1 *byte) int32 {
	var (
		result int32
		v2     uint32
		v3     int8
		v4     *uint8
		v5     *byte
	)
	if a1 == nil {
		return 0
	}
	result = 0
	if libc.StrLen(a1) <= 128 {
		v2 = uint32(libc.StrLen(a1) + 1)
		v3 = int8(uint8(v2))
		v2 >>= 2
		libc.MemCpy(memmap.PtrOff(6112660, 1308192), unsafe.Pointer(a1), int(v2*4))
		v5 = (*byte)(unsafe.Add(unsafe.Pointer(a1), v2*4))
		v4 = (*uint8)(memmap.PtrOff(6112660, v2*4+0x13F620))
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v2))), 0)) = uint8(v3)
		result = 1
		libc.MemCpy(unsafe.Pointer(v4), unsafe.Pointer(v5), int(v2&3))
	}
	return result
}
func sub_4A7D50(a1 *byte) int32 {
	var result int32
	result = int32(uintptr(unsafe.Pointer(a1)))
	if a1 != nil {
		*memmap.PtrUint32(6112660, 1308348) = uint32(libc.Atoi(libc.GoString(a1)))
		result = 1
	}
	return result
}
func sub_4A7D70(a1 *byte) int32 {
	var (
		result int32
		v2     *File
		v3     *File
	)
	result = int32(uintptr(unsafe.Pointer(a1)))
	if a1 != nil {
		v2 = nox_fs_open_text(a1)
		v3 = v2
		if v2 == nil {
			return 0
		}
		nox_fs_fgets(v2, (*byte)(memmap.PtrOff(6112660, 1308388)), 256)
		if !nox_fs_feof(v3) {
			for sub_4A7DF0((*byte)(memmap.PtrOff(6112660, 1308388))) != 0 {
				nox_fs_fgets(v3, (*byte)(memmap.PtrOff(6112660, 1308388)), 256)
				if nox_fs_feof(v3) {
					break
				}
			}
			return 0
		}
		nox_fs_close(v3)
		result = 1
	}
	return result
}
func sub_4A7DF0(a1 *byte) int32 {
	var (
		v1 *byte
		v2 int32
		v3 *byte
		v5 *byte
	)
	v1 = libc.StrTok(a1, CString(" \t\r\n"))
	if v1 == nil {
		return 0
	}
	v2 = 0
	v3 = CString("MAP")
	for libc.StrCaseCmp(v1, v3) != 0 || libc.StrTok(nil, CString(" \t\r\n")) == nil {
		v3 = (*byte)(unsafe.Add(unsafe.Pointer(v3), 20))
		v2++
		if int32(uintptr(unsafe.Pointer(v3))) >= int32(uintptr(memmap.PtrOff(0x587000, 172172))) {
			return 0
		}
	}
	v5 = libc.StrTok(nil, CString("\t\r\n"))
	(cgoAsFunc(*(*uint32)(memmap.PtrOff(0x587000, uint32(v2*20)+0x29F98)), (*func(*byte))(nil)))(v5)
	return 1
}
func sub_4A7E70() *byte {
	return (*byte)(memmap.PtrOff(6112660, 1308644))
}
func sub_4A7E80() *byte {
	return (*byte)(memmap.PtrOff(6112660, 1308172))
}
func sub_4A7E90() *byte {
	return (*byte)(memmap.PtrOff(6112660, 1308352))
}
func sub_4A7EA0() int32 {
	return int32(*memmap.PtrUint32(6112660, 1308184))
}
func sub_4A7EB0() int32 {
	return int32(*memmap.PtrUint32(6112660, 1308740))
}
func sub_4A7EC0() *byte {
	return (*byte)(memmap.PtrOff(6112660, 1308324))
}
func sub_4A7ED0() *byte {
	return (*byte)(memmap.PtrOff(6112660, 1308364))
}
func sub_4A7EE0() int32 {
	return int32(*memmap.PtrUint32(6112660, 1308188))
}
func sub_4A7EF0() *byte {
	return (*byte)(memmap.PtrOff(6112660, 1308732))
}
func sub_4A7F00() int32 {
	return int32(*memmap.PtrUint32(6112660, 1308728))
}
func sub_4A7F10() *byte {
	return (*byte)(memmap.PtrOff(6112660, 1308192))
}
func sub_4A7F20() int32 {
	return int32(*memmap.PtrUint32(6112660, 1308348))
}
func sub_4A7F30() int32 {
	return int32(*memmap.PtrUint32(6112660, 1308744))
}
func sub_4A7F40(a1 int32) int32 {
	var result int32
	result = a1
	*memmap.PtrUint32(6112660, 1308744) = uint32(a1)
	return result
}
func nox_xxx_wndButtonProc_4A7F50(win *nox_window, ev int32, a3 int32, a4 int32) int32 {
	switch ev {
	case 5:
		win.draw_data.field_0 |= 4
		return 1
	case 6:
		fallthrough
	case 7:
		if (win.draw_data.field_0 & 4) == 0 {
			return 0
		}
		win.draw_data.win.Func94(asWindowEvent(0x4000|7, int32(uintptr(unsafe.Pointer(win))), a3))
		win.draw_data.field_0 &= 0xFFFFFFFB
		return 1
	case 8:
		win.draw_data.win.Func94(asWindowEvent(0x4000, int32(uintptr(unsafe.Pointer(win))), a3))
		return 1
	case 17:
		if win.draw_data.style&256 != 0 {
			win.draw_data.field_0 |= 2
			win.draw_data.win.Func94(asWindowEvent(0x4000|5, int32(uintptr(unsafe.Pointer(win))), a3))
			guiFocus(win)
		}
		return 1
	case 18:
		if win.draw_data.style&256 != 0 {
			win.draw_data.field_0 &= 0xFFFFFFFD
			win.draw_data.win.Func94(asWindowEvent(0x4000|6, int32(uintptr(unsafe.Pointer(win))), a3))
		}
		if win.draw_data.field_0&4 != 0 {
			win.draw_data.field_0 &= 0xFFFFFFFB
		}
		return 1
	case 21:
		switch a3 {
		case 15:
			fallthrough
		case 205:
			fallthrough
		case 208:
			if a4 == 2 {
				nox_xxx_wndRetNULL_46A8A0()
			}
			return 1
		case 28:
			fallthrough
		case 57:
			if a4 != 1 {
				win.draw_data.field_0 |= 4
				return 1
			}
			if win.draw_data.field_0&4 != 0 {
				win.draw_data.win.Func94(asWindowEvent(0x4000|7, int32(uintptr(unsafe.Pointer(win))), 0))
				win.draw_data.field_0 &= 0xFFFFFFFB
			}
			return 1
		case 200:
			fallthrough
		case 203:
			if a4 == 2 {
				nox_xxx_wndRetNULL_0_46A8B0()
			}
			return 1
		default:
			return 0
		}
		fallthrough
	default:
		return 0
	}
}
func nox_xxx_wndButtonDrawNoImg_4A81D0(a1 int32, a2 int32) int32 {
	var (
		v2    int32
		v3    int32
		v4    int32
		v5    int32
		v6    int32
		v7    int32
		v9    int32
		xLeft int32
		yTop  int32
	)
	v2 = a2
	v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 28))))
	v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 20))))
	nox_client_wndGetPosition_46AA60((*nox_window)(unsafe.Pointer(uintptr(a1))), (*uint32)(unsafe.Pointer(&xLeft)), (*uint32)(unsafe.Pointer(&yTop)))
	if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 4))))&8 != 0 {
		if *(*uint32)(unsafe.Pointer(uintptr(v2)))&4 != 0 {
			v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 52))))
		} else if *(*uint32)(unsafe.Pointer(uintptr(v2)))&2 != 0 {
			v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 36))))
		}
	} else {
		v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 44))))
	}
	if uint32(v3) != 0x80000000 {
		nox_client_drawSetColor_434460(v3)
		nox_client_drawBorderLines_49CC70(xLeft, yTop, int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 8)))), int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 12)))))
	}
	if uint32(v4) != 0x80000000 {
		nox_client_drawSetColor_434460(v4)
		nox_client_drawRectFilledOpaque_49CE30(xLeft+1, yTop+1, int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 8)))-2), int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 12)))-2))
	}
	if int32(*(*uint16)(unsafe.Pointer(uintptr(v2 + 72)))) != 0 {
		v5 = int32(uint32(xLeft) + *(*uint32)(unsafe.Pointer(uintptr(a1 + 8)))/2)
		v6 = int32(uint32(yTop) + *(*uint32)(unsafe.Pointer(uintptr(a1 + 12)))/2)
		if (*(*uint32)(unsafe.Pointer(uintptr(a1 + 4))) & 8192) == 8192 {
			nox_draw_enableTextSmoothing_43F670(1)
		}
		nox_xxx_drawGetStringSize_43F840(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 200))))), (*libc.WChar)(unsafe.Pointer(uintptr(v2+72))), &a2, nil, 0)
		if *(*uint32)(unsafe.Pointer(uintptr(v2 + 68))) != 0x80000000 {
			nox_xxx_drawSetTextColor_434390(int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 68)))))
			v9 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 8))))
			v7 = nox_xxx_guiFontHeightMB_43F320(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 200))))))
			nox_xxx_drawStringWrap_43FAF0(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 200))))), (*libc.WChar)(unsafe.Pointer(uintptr(v2+72))), v5-a2/2, v6-v7/2, v9, 0)
		}
		nox_draw_enableTextSmoothing_43F670(0)
	}
	return 1
}
func nox_xxx_wndButtonInit_4A8340(a1 int32) int32 {
	var result int32
	if int32(int8(*(*uint8)(unsafe.Pointer(uintptr(a1 + 4))))) >= 0 {
		result = (*nox_window)(unsafe.Pointer(uintptr(a1))).SetAllFuncs(func(arg1 int32, arg2 int32, arg3 int32, arg4 int32) int32 {
			return nox_xxx_wndButtonProc_4A7F50((*nox_window)(unsafe.Pointer(uintptr(arg1))), arg2, arg3, arg4)
		}, func(arg1 *nox_window, arg2 unsafe.Pointer) int32 {
			return nox_xxx_wndButtonDrawNoImg_4A81D0(int32(uintptr(unsafe.Pointer(arg1))), int32(uintptr(arg2)))
		}, nil)
	} else {
		result = (*nox_window)(unsafe.Pointer(uintptr(a1))).SetAllFuncs(func(arg1 int32, arg2 int32, arg3 int32, arg4 int32) int32 {
			return nox_xxx_wndButtonProc_4A7F50((*nox_window)(unsafe.Pointer(uintptr(arg1))), arg2, arg3, arg4)
		}, func(arg1 *nox_window, arg2 unsafe.Pointer) int32 {
			return nox_xxx_wndButtonDraw_4A8380((*uint32)(unsafe.Pointer(arg1)), int32(uintptr(arg2)))
		}, nil)
	}
	return result
}
func nox_xxx_wndButtonDraw_4A8380(a1 *uint32, a2 int32) int32 {
	var (
		v2  int32
		v3  *uint32
		v4  int32
		v5  int32
		v6  int32
		v7  int32
		v8  int32
		v10 int32
		v11 int32
	)
	v2 = a2
	v3 = a1
	v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 32))))
	v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 24))))
	nox_client_wndGetPosition_46AA60((*nox_window)(unsafe.Pointer(a1)), (*uint32)(unsafe.Pointer(&a1)), (*uint32)(unsafe.Pointer(&a2)))
	if *(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*1))&8 != 0 {
		if *(*uint32)(unsafe.Pointer(uintptr(v2)))&4 != 0 {
			v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 56))))
		} else if *(*uint32)(unsafe.Pointer(uintptr(v2)))&2 != 0 {
			v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 40))))
		}
	} else {
		v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 48))))
	}
	if v5 != 0 {
		nox_client_drawImageAt_47D2C0((*nox_video_bag_image_t)(unsafe.Pointer(uintptr(v5))), int32(uint32(int32(uintptr(unsafe.Pointer(a1))))+*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*24))), int32(uint32(a2)+*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*25))))
	}
	if v4 != 0 {
		nox_client_drawImageAt_47D2C0((*nox_video_bag_image_t)(unsafe.Pointer(uintptr(v4))), int32(uint32(int32(uintptr(unsafe.Pointer(a1))))+*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*24))), int32(uint32(a2)+*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*25))))
	}
	if int32(*(*uint16)(unsafe.Pointer(uintptr(v2 + 72)))) != 0 {
		v6 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*3)))
		a1 = (*uint32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(a1))), *(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*2))/2))))
		v7 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*1)) & 8192)
		a2 += v6 / 2
		if v7 == 8192 {
			nox_draw_enableTextSmoothing_43F670(1)
		}
		nox_xxx_drawGetStringSize_43F840(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 200))))), (*libc.WChar)(unsafe.Pointer(uintptr(v2+72))), &v11, nil, 0)
		if *(*uint32)(unsafe.Pointer(uintptr(v2 + 68))) != 0x80000000 {
			nox_xxx_drawSetTextColor_434390(int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 68)))))
			v10 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*2)))
			v8 = nox_xxx_guiFontHeightMB_43F320(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 200))))))
			nox_xxx_drawStringWrap_43FAF0(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 200))))), (*libc.WChar)(unsafe.Pointer(uintptr(v2+72))), int32(uintptr(unsafe.Pointer(a1)))-v11/2, a2-v8/2, v10, 0)
		}
		nox_draw_enableTextSmoothing_43F670(0)
	}
	return 1
}
func nox_xxx_wndRadioButtonProc_4A84E0(a1 *uint32, a2 int32, a3 int32, a4 int32) int32 {
	var (
		v4     int32
		v5     int32
		result int32
		v7     int32
		v8     int32
		v9     int32
		v10    int32
		i      *uint32
		v12    int32
		v13    int32
		j      *uint32
	)
	switch a2 {
	case 5:
		goto LABEL_31
	case 6:
		fallthrough
	case 7:
		v9 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*9)))
		v10 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*99)))
		if v9&4 != 0 {
			if v9&2 != 0 {
			LABEL_31:
				result = 1
			} else {
			LABEL_16:
				result = 0
			}
		} else {
			(*nox_window)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*13))))).Func94(asWindowEvent(0x4007, int32(uintptr(unsafe.Pointer(a1))), a3))
			if v10 != 0 {
				for i = *(**uint32)(unsafe.Pointer(uintptr(v10 + 400))); i != nil; i = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(i), unsafe.Sizeof(uint32(0))*97))))) {
					if i != a1 && *(*uint32)(unsafe.Add(unsafe.Pointer(i), unsafe.Sizeof(uint32(0))*10)) == *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*10)) {
						*(*uint32)(unsafe.Add(unsafe.Pointer(i), unsafe.Sizeof(uint32(0))*9)) &= 0xFFFFFFFB
					}
				}
			}
			v12 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*9)))
			*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v12))), 0)) = uint8(int8(v12 | 4))
			*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*9)) = uint32(v12)
			result = 1
		}
	case 8:
		(*nox_window)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*13))))).Func94(asWindowEvent(0x4000, int32(uintptr(unsafe.Pointer(a1))), a3))
		return 1
	case 17:
		v4 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*11)))
		if (v4 & 256) == 0 {
			goto LABEL_31
		}
		v5 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*9)))
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v5))), 0)) = uint8(int8(v5 | 2))
		*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*9)) = uint32(v5)
		(*nox_window)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*13))))).Func94(asWindowEvent(0x4005, int32(uintptr(unsafe.Pointer(a1))), a3))
		guiFocus((*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1)))))))
		result = 1
	case 18:
		v7 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*11)))
		if (v7 & 256) == 0 {
			goto LABEL_31
		}
		v8 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*13)))
		*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*9)) &= 0xFFFFFFFD
		(*nox_window)(unsafe.Pointer(uintptr(v8))).Func94(asWindowEvent(0x4006, int32(uintptr(unsafe.Pointer(a1))), a3))
		result = 1
	case 21:
		switch a3 {
		case 15:
			fallthrough
		case 205:
			fallthrough
		case 208:
			if a4 != 2 {
				goto LABEL_31
			}
			nox_xxx_wndRetNULL_46A8A0()
			result = 1
		case 28:
			fallthrough
		case 57:
			if a4 != 2 {
				goto LABEL_31
			}
			v13 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*99)))
			if *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*9))&4 != 0 {
				goto LABEL_31
			}
			(*nox_window)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*13))))).Func94(asWindowEvent(0x4007, int32(uintptr(unsafe.Pointer(a1))), 0))
			if v13 != 0 {
				for j = *(**uint32)(unsafe.Pointer(uintptr(v13 + 400))); j != nil; j = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(j), unsafe.Sizeof(uint32(0))*97))))) {
					if j != a1 && *(*uint32)(unsafe.Add(unsafe.Pointer(j), unsafe.Sizeof(uint32(0))*10)) == *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*10)) {
						*(*uint32)(unsafe.Add(unsafe.Pointer(j), unsafe.Sizeof(uint32(0))*9)) &= 0xFFFFFFFB
					}
				}
			}
			*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*9)) ^= 4
			result = 1
		case 200:
			fallthrough
		case 203:
			if a4 == 2 {
				nox_xxx_wndRetNULL_0_46A8B0()
			}
			goto LABEL_31
		default:
			goto LABEL_16
		}
		return result
	default:
		goto LABEL_16
	}
	return result
}
func nox_xxx_wndRadioButtonSetAllFn_4A87E0(a1 int32) int32 {
	var result int32
	if int32(int8(*(*uint8)(unsafe.Pointer(uintptr(a1 + 4))))) >= 0 {
		result = (*nox_window)(unsafe.Pointer(uintptr(a1))).SetAllFuncs(func(arg1 int32, arg2 int32, arg3 int32, arg4 int32) int32 {
			return nox_xxx_wndRadioButtonProc_4A84E0((*uint32)(unsafe.Pointer(uintptr(arg1))), arg2, arg3, arg4)
		}, func(arg1 *nox_window, arg2 unsafe.Pointer) int32 {
			return nox_xxx_wndRadioButtonDrawNoImg_4A8820(int32(uintptr(unsafe.Pointer(arg1))), int32(uintptr(arg2)))
		}, nil)
	} else {
		result = (*nox_window)(unsafe.Pointer(uintptr(a1))).SetAllFuncs(func(arg1 int32, arg2 int32, arg3 int32, arg4 int32) int32 {
			return nox_xxx_wndRadioButtonProc_4A84E0((*uint32)(unsafe.Pointer(uintptr(arg1))), arg2, arg3, arg4)
		}, func(arg1 *nox_window, arg2 unsafe.Pointer) int32 {
			return nox_xxx_wndRadioButtonDraw_4A8A20(int32(uintptr(unsafe.Pointer(arg1))), int32(uintptr(arg2)))
		}, nil)
	}
	return result
}
func nox_xxx_wndRadioButtonDrawNoImg_4A8820(a1 int32, a2 int32) int32 {
	var (
		v2  int32
		v3  int32
		v4  int32
		v5  int32
		v6  int32
		v7  int32
		v8  int32
		v9  int32
		v11 int32
		v12 int32
		v13 int32
		v14 int32
		v15 int32
	)
	v2 = a2
	v3 = a1
	v12 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 28))))
	a2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 20))))
	nox_client_wndGetPosition_46AA60((*nox_window)(unsafe.Pointer(uintptr(a1))), (*uint32)(unsafe.Pointer(&v14)), (*uint32)(unsafe.Pointer(&v15)))
	if int32(*(*uint8)(unsafe.Pointer(uintptr(v3 + 4))))&8 != 0 {
		if int32(*(*uint8)(unsafe.Pointer(uintptr(v2))))&2 != 0 {
			v12 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 36))))
		}
	} else {
		a2 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 44))))
	}
	v4 = v14 + 4
	v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(v3 + 12)))/2 + uint32(v15) - 5)
	v13 = nox_xxx_guiFontHeightMB_43F320(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 200))))))
	if uint32(a2) != 0x80000000 {
		nox_client_drawSetColor_434460(a2)
		nox_client_drawRectFilledOpaque_49CE30(v4, v5, 10, 10)
	}
	if uint32(v12) != 0x80000000 {
		nox_client_drawSetColor_434460(v12)
		nox_client_drawBorderLines_49CC70(v4, v5, 10, 10)
	}
	if int32(*(*uint8)(unsafe.Pointer(uintptr(v2))))&4 != 0 && *(*uint32)(unsafe.Pointer(uintptr(v2 + 52))) != 0x80000000 {
		nox_client_drawSetColor_434460(int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 52)))))
		nox_client_drawRectFilledOpaque_49CE30(v4+1, v5+1, 8, 8)
	}
	if (*(*uint32)(unsafe.Pointer(uintptr(v3 + 4))) & 8192) == 8192 {
		nox_draw_enableTextSmoothing_43F670(1)
	}
	if **(**uint32)(unsafe.Pointer(uintptr(v3 + 32))) != 0 {
		v6 = int32(*(*uint32)(unsafe.Pointer(uintptr(v3 + 12))))
		v14 += int32(*(*uint32)(unsafe.Pointer(uintptr(v3 + 8))) / 2)
		v7 = v6/2 + v15
		v8 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 200))))
		v15 = v7
		nox_xxx_drawGetStringSize_43F840(unsafe.Pointer(uintptr(v8)), (*libc.WChar)(unsafe.Pointer(uintptr(v2+72))), &a1, nil, 0)
		if *(*uint32)(unsafe.Pointer(uintptr(v2 + 68))) != 0x80000000 {
			nox_xxx_drawSetTextColor_434390(int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 68)))))
			v11 = int32(*(*uint32)(unsafe.Pointer(uintptr(v3 + 8))))
			v9 = nox_xxx_guiFontHeightMB_43F320(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 200))))))
			nox_xxx_drawStringWrap_43FAF0(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 200))))), (*libc.WChar)(unsafe.Pointer(uintptr(v2+72))), v14-a1/2, v15-v9/2, v11, 0)
		}
	} else if *(*uint32)(unsafe.Pointer(uintptr(v2 + 68))) != 0x80000000 {
		nox_xxx_drawSetTextColor_434390(int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 68)))))
		nox_xxx_drawStringWrap_43FAF0(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 200))))), (*libc.WChar)(unsafe.Pointer(uintptr(v2+72))), v4+14, v5-v13/2+5, int32(*(*uint32)(unsafe.Pointer(uintptr(v3 + 8)))), 0)
	}
	nox_draw_enableTextSmoothing_43F670(0)
	return 1
}
func nox_xxx_wndRadioButtonDraw_4A8A20(a1 int32, a2 int32) int32 {
	var (
		v2  int32
		v3  int32
		v4  int32
		v5  int32
		v6  int32
		v7  int32
		v8  int32
		v9  int32
		v10 int32
		v12 int32
		v13 int32
		v14 int32
	)
	v2 = a2
	v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 32))))
	v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 24))))
	nox_client_wndGetPosition_46AA60((*nox_window)(unsafe.Pointer(uintptr(a1))), (*uint32)(unsafe.Pointer(&v13)), (*uint32)(unsafe.Pointer(&v14)))
	if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 4))))&8 != 0 {
		if int32(*(*uint8)(unsafe.Pointer(uintptr(v2))))&2 != 0 {
			v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 40))))
		}
	} else {
		v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 48))))
	}
	if v4 != 0 {
		nox_client_drawImageAt_47D2C0((*nox_video_bag_image_t)(unsafe.Pointer(uintptr(v4))), int32(uint32(v13)+*(*uint32)(unsafe.Pointer(uintptr(v2 + 60)))), int32(uint32(v14)+*(*uint32)(unsafe.Pointer(uintptr(v2 + 64)))))
	}
	if int32(*(*uint8)(unsafe.Pointer(uintptr(v2))))&4 != 0 {
		v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 56))))
		if v5 != 0 {
			nox_client_drawImageAt_47D2C0((*nox_video_bag_image_t)(unsafe.Pointer(uintptr(v5))), int32(uint32(v13)+*(*uint32)(unsafe.Pointer(uintptr(v2 + 60)))), int32(uint32(v14)+*(*uint32)(unsafe.Pointer(uintptr(v2 + 64)))))
		}
	} else if v3 != 0 {
		nox_client_drawImageAt_47D2C0((*nox_video_bag_image_t)(unsafe.Pointer(uintptr(v3))), int32(uint32(v13)+*(*uint32)(unsafe.Pointer(uintptr(v2 + 60)))), int32(uint32(v14)+*(*uint32)(unsafe.Pointer(uintptr(v2 + 64)))))
	}
	if (*(*uint32)(unsafe.Pointer(uintptr(a1 + 4))) & 8192) == 8192 {
		nox_draw_enableTextSmoothing_43F670(1)
	}
	if **(**uint32)(unsafe.Pointer(uintptr(a1 + 32))) != 0 {
		v6 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 8))))
		v7 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 200))))
		v14 += int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 12))) / 2)
		v13 += v6 / 2
		nox_xxx_drawGetStringSize_43F840(unsafe.Pointer(uintptr(v7)), (*libc.WChar)(unsafe.Pointer(uintptr(v2+72))), &a2, nil, 0)
		if *(*uint32)(unsafe.Pointer(uintptr(v2 + 68))) != 0x80000000 {
			nox_xxx_drawSetTextColor_434390(int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 68)))))
			v12 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 8))))
			v8 = nox_xxx_guiFontHeightMB_43F320(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 200))))))
			nox_xxx_drawStringWrap_43FAF0(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 200))))), (*libc.WChar)(unsafe.Pointer(uintptr(v2+72))), v13-a2/2, v14-v8/2, v12, 0)
		}
	} else {
		v9 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 12))) - uint32(nox_xxx_guiFontHeightMB_43F320(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 200))))))))
		v10 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 68))))
		v14 += v9 / 2
		v13 += 28
		if uint32(v10) != 0x80000000 {
			nox_xxx_drawSetTextColor_434390(v10)
			nox_xxx_drawStringWrap_43FAF0(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 200))))), (*libc.WChar)(unsafe.Pointer(uintptr(v2+72))), v13, v14, int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 8)))), 0)
		}
	}
	nox_draw_enableTextSmoothing_43F670(0)
	return 1
}
func nox_xxx_wndCheckBoxProc_4A8C00(a1 int32, a2 int32, a3 int32, a4 int32) int32 {
	var (
		v4     int32
		v5     int32
		result int32
		v7     int32
		v8     int32
		v9     int32
	)
	switch a2 {
	case 5:
		goto LABEL_17
	case 6:
		fallthrough
	case 7:
		v8 = a1
		if (int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 36)))) & 2) == 0 {
			goto LABEL_18
		}
		(*nox_window)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 52)))))).Func94(asWindowEvent(0x4007, a1, a3))
	LABEL_9:
		*(*uint32)(unsafe.Pointer(uintptr(v8 + 36))) ^= 4
		result = 1
	case 8:
		(*nox_window)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 52)))))).Func94(asWindowEvent(0x4000, a1, a3))
		return 1
	case 17:
		v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 44))))
		if (v4 & 256) == 0 {
			goto LABEL_17
		}
		v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 52))))
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 36))) |= 2
		(*nox_window)(unsafe.Pointer(uintptr(v5))).Func94(asWindowEvent(0x4005, a1, a3))
		guiFocus((*nox_window)(unsafe.Pointer(uintptr(a1))))
		result = 1
	case 18:
		v7 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 44))))
		if v7&256 != 0 {
			v9 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 52))))
			*(*uint32)(unsafe.Pointer(uintptr(a1 + 36))) &= 0xFFFFFFFD
			(*nox_window)(unsafe.Pointer(uintptr(v9))).Func94(asWindowEvent(0x4006, a1, a3))
			result = 1
		} else {
		LABEL_17:
			result = 1
		}
	case 21:
		switch a3 {
		case 15:
			fallthrough
		case 205:
			fallthrough
		case 208:
			if a4 != 2 {
				goto LABEL_17
			}
			nox_xxx_wndRetNULL_46A8A0()
			return 1
		case 28:
			fallthrough
		case 57:
			if a4 != 2 {
				goto LABEL_17
			}
			v8 = a1
			(*nox_window)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v8 + 52)))))).Func94(asWindowEvent(0x4007, v8, 0))
			goto LABEL_9
		case 200:
			fallthrough
		case 203:
			if a4 == 2 {
				nox_xxx_wndRetNULL_0_46A8B0()
			}
			goto LABEL_17
		default:
			goto LABEL_18
		}
		fallthrough
	default:
	LABEL_18:
		result = 0
	}
	return result
}
func nox_xxx_wndCheckBoxInit_4A8E60(a1 int32) int32 {
	var result int32
	if int32(int8(*(*uint8)(unsafe.Pointer(uintptr(a1 + 4))))) >= 0 {
		result = (*nox_window)(unsafe.Pointer(uintptr(a1))).SetAllFuncs(nox_xxx_wndCheckBoxProc_4A8C00, func(arg1 *nox_window, arg2 unsafe.Pointer) int32 {
			return nox_xxx_wndDrawCheckBoxNoImg_4A8EA0(int32(uintptr(unsafe.Pointer(arg1))), int32(uintptr(arg2)))
		}, nil)
	} else {
		result = (*nox_window)(unsafe.Pointer(uintptr(a1))).SetAllFuncs(nox_xxx_wndCheckBoxProc_4A8C00, func(arg1 *nox_window, arg2 unsafe.Pointer) int32 {
			return nox_xxx_wndDrawCheckBox_4A9050((*uint32)(unsafe.Pointer(arg1)), int32(uintptr(arg2)))
		}, nil)
	}
	return result
}
func nox_xxx_wndDrawCheckBoxNoImg_4A8EA0(a1 int32, a2 int32) int32 {
	var (
		v2    int32
		v3    int32
		v4    int32
		v5    int32
		v6    int32
		v7    int32
		v9    int32
		xLeft int32
		yTop  int32
		v12   int32
	)
	v2 = a2
	v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 28))))
	v12 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 20))))
	nox_client_wndGetPosition_46AA60((*nox_window)(unsafe.Pointer(uintptr(a1))), (*uint32)(unsafe.Pointer(&xLeft)), (*uint32)(unsafe.Pointer(&yTop)))
	if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 4))))&8 != 0 {
		if int32(*(*uint8)(unsafe.Pointer(uintptr(v2))))&2 != 0 {
			v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 36))))
		}
	} else {
		v12 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 44))))
	}
	v4 = int32(uint32(yTop) + *(*uint32)(unsafe.Pointer(uintptr(a1 + 12)))/2 - 5)
	v5 = xLeft + 4
	if uint32(v12) != 0x80000000 {
		nox_client_drawSetColor_434460(v12)
		nox_client_drawRectFilledOpaque_49CE30(xLeft, yTop, int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 8)))), int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 12)))))
	}
	if uint32(v3) != 0x80000000 {
		nox_client_drawSetColor_434460(v3)
		nox_client_drawBorderLines_49CC70(v5, v4, 10, 10)
	}
	if int32(*(*uint8)(unsafe.Pointer(uintptr(v2))))&4 != 0 {
		if *(*uint32)(unsafe.Pointer(uintptr(v2 + 52))) != 0x80000000 {
			nox_client_drawSetColor_434460(int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 52)))))
			nox_client_drawRectFilledOpaque_49CE30(v5+1, v4+1, 8, 8)
			if uint32(v3) != 0x80000000 {
				nox_client_drawSetColor_434460(v3)
				nox_client_drawAddPoint_49F500(v5, v4)
				nox_xxx_rasterPointRel_49F570(9, 9)
				nox_client_drawLineFromPoints_49E4B0()
				nox_client_drawAddPoint_49F500(v5, v4+9)
				nox_xxx_rasterPointRel_49F570(9, -9)
				nox_client_drawLineFromPoints_49E4B0()
			}
		}
	}
	v6 = 5 - nox_xxx_guiFontHeightMB_43F320(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 200))))))/2 + v4
	v7 = v5 + 14
	nox_xxx_drawGetStringSize_43F840(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 200))))), (*libc.WChar)(unsafe.Pointer(uintptr(v2+72))), &v9, nil, 0)
	if *(*uint32)(unsafe.Pointer(uintptr(v2 + 68))) != 0x80000000 {
		if (*(*uint32)(unsafe.Pointer(uintptr(a1 + 4))) & 8192) == 8192 {
			nox_draw_enableTextSmoothing_43F670(1)
		}
		nox_xxx_drawSetTextColor_434390(int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 68)))))
		nox_xxx_drawStringWrap_43FAF0(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 200))))), (*libc.WChar)(unsafe.Pointer(uintptr(v2+72))), v7, v6, int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 8)))), 0)
		nox_draw_enableTextSmoothing_43F670(0)
	}
	return 1
}
func nox_xxx_wndDrawCheckBox_4A9050(a1 *uint32, a2 int32) int32 {
	var (
		v2  int32
		v3  *uint32
		v4  int32
		v5  int32
		v6  int32
		v7  int32
		v8  int32
		v9  int32
		v11 int32
		v12 int32
		v13 int32
	)
	v2 = a2
	v3 = a1
	v11 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 32))))
	a2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 24))))
	nox_client_wndGetPosition_46AA60((*nox_window)(unsafe.Pointer(a1)), (*uint32)(unsafe.Pointer(&v12)), (*uint32)(unsafe.Pointer(&v13)))
	v4 = int32(uint32(v13) + *(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*3))/2)
	if *(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*1))&8 != 0 {
		if int32(*(*uint8)(unsafe.Pointer(uintptr(v2))))&2 != 0 {
			v11 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 40))))
		}
	} else {
		a2 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 48))))
	}
	v5 = v4 - 5
	v6 = v12 + 4
	if a2 != 0 {
		nox_client_drawImageAt_47D2C0((*nox_video_bag_image_t)(unsafe.Pointer(uintptr(a2))), int32(uint32(v6)+*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*24))), int32(uint32(v5)+*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*25))))
	}
	if int32(*(*uint8)(unsafe.Pointer(uintptr(v2))))&4 != 0 {
		v7 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 56))))
		if v7 == 0 {
			goto LABEL_12
		}
	} else {
		v7 = v11
		if v11 == 0 {
			goto LABEL_12
		}
	}
	nox_client_drawImageAt_47D2C0((*nox_video_bag_image_t)(unsafe.Pointer(uintptr(v7))), int32(uint32(v6)+*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*24))), int32(uint32(v5)+*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*25))))
LABEL_12:
	v8 = 5 - nox_xxx_guiFontHeightMB_43F320(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 200))))))/2 + v5
	v9 = v6 + 16
	nox_xxx_drawGetStringSize_43F840(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 200))))), (*libc.WChar)(unsafe.Pointer(uintptr(v2+72))), (*int32)(unsafe.Pointer(&a1)), nil, 0)
	if *(*uint32)(unsafe.Pointer(uintptr(v2 + 68))) != 0x80000000 {
		if (*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*1)) & 8192) == 8192 {
			nox_draw_enableTextSmoothing_43F670(1)
		}
		nox_xxx_drawSetTextColor_434390(int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 68)))))
		nox_xxx_drawStringWrap_43FAF0(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 200))))), (*libc.WChar)(unsafe.Pointer(uintptr(v2+72))), v9, v8, int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*2))), 0)
		nox_draw_enableTextSmoothing_43F670(0)
	}
	return 1
}
func nox_gui_newButtonOrCheckbox_4A91A0(parent *nox_window, a2 int32, a3 int32, a4 int32, a5 int32, a6 int32, draw *nox_window_data) *nox_window {
	if draw.style&1 != 0 {
		var btn *nox_window = nox_window_new(parent, a2, a3, a4, a5, a6, func(arg1 int32, arg2 int32, arg3 int32, arg4 int32) int32 {
			return nox_xxx_wndButtonProcPre_4A9250(arg1, arg2, (*libc.WChar)(unsafe.Pointer(uintptr(arg3))), arg4)
		})
		if btn == nil {
			return nil
		}
		nox_xxx_wndButtonInit_4A8340(int32(uintptr(unsafe.Pointer(btn))))
		if draw.win == nil {
			draw.win = btn
		}
		nox_gui_windowCopyDrawData_46AF80(btn, unsafe.Pointer(draw))
		return btn
	}
	if draw.style&4 != 0 {
		var btn *nox_window = nox_window_new(parent, a2, a3, a4, a5, a6, func(arg1 int32, arg2 int32, arg3 int32, arg4 int32) int32 {
			return nox_xxx_wndCheckboxProcMB_4A92C0(arg1, arg2, (*libc.WChar)(unsafe.Pointer(uintptr(arg3))), arg4)
		})
		if btn == nil {
			return nil
		}
		nox_xxx_wndCheckBoxInit_4A8E60(int32(uintptr(unsafe.Pointer(btn))))
		if draw.win == nil {
			draw.win = btn
		}
		nox_gui_windowCopyDrawData_46AF80(btn, unsafe.Pointer(draw))
		return btn
	}
	return nil
}
func nox_xxx_wndButtonProcPre_4A9250(a1 int32, a2 int32, a3 *libc.WChar, a4 int32) int32 {
	var (
		result int32
		v4     int32
		v5     int32
	)
	if a2 == 23 {
		if a3 == nil {
			v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 36))))
			*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v4))), 0)) = uint8(int8(v4 & 253))
			*(*uint32)(unsafe.Pointer(uintptr(a1 + 36))) = uint32(v4)
		}
		v5 = (*nox_window)(unsafe.Pointer(uintptr(a1))).ID()
		(*nox_window)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 52)))))).Func94(asWindowEvent(0x4003, int32(uintptr(unsafe.Pointer(a3))), v5))
		result = 1
	} else {
		if a2 == 0x4001 {
			nox_wcsncpy((*libc.WChar)(unsafe.Pointer(uintptr(a1+108))), a3, 63)
			*(*uint16)(unsafe.Pointer(uintptr(a1 + 234))) = 0
		}
		result = 0
	}
	return result
}
func nox_xxx_wndCheckboxProcMB_4A92C0(a1 int32, a2 int32, a3 *libc.WChar, a4 int32) int32 {
	var (
		result int32
		v4     int32
		v5     int32
	)
	if a2 == 23 {
		if a3 == nil {
			v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 36))))
			*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v4))), 0)) = uint8(int8(v4 & 253))
			*(*uint32)(unsafe.Pointer(uintptr(a1 + 36))) = uint32(v4)
		}
		v5 = (*nox_window)(unsafe.Pointer(uintptr(a1))).ID()
		(*nox_window)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 52)))))).Func94(asWindowEvent(0x4003, int32(uintptr(unsafe.Pointer(a3))), v5))
		result = 1
	} else {
		if a2 == 0x4001 {
			nox_wcsncpy((*libc.WChar)(unsafe.Pointer(uintptr(a1+108))), a3, 63)
			*(*uint16)(unsafe.Pointer(uintptr(a1 + 234))) = 0
		}
		result = 0
	}
	return result
}
func nox_gui_newRadioButton_4A9330(parent *nox_window, a2 int32, a3 int32, a4 int32, a5 int32, a6 int32, draw *nox_window_data, data *nox_radioButton_data) *nox_window {
	if (draw.style & 2) == 0 {
		return nil
	}
	var win *nox_window = nox_window_new(parent, a2, a3, a4, a5, a6, func(arg1 int32, arg2 int32, arg3 int32, arg4 int32) int32 {
		return nox_xxx_wndRadioButtonProcPre_4A93C0(arg1, arg2, (*libc.WChar)(unsafe.Pointer(uintptr(arg3))), arg4)
	})
	if win == nil {
		return nil
	}
	nox_xxx_wndRadioButtonSetAllFn_4A87E0(int32(uintptr(unsafe.Pointer(win))))
	if draw.win == nil {
		draw.win = win
	}
	var d *nox_radioButton_data = new(nox_radioButton_data)
	d.field_0 = 0
	if data != nil {
		d.field_0 = data.field_0
	}
	win.widget_data = unsafe.Pointer(d)
	nox_gui_windowCopyDrawData_46AF80(win, unsafe.Pointer(draw))
	return win
}
func nox_xxx_wndRadioButtonProcPre_4A93C0(a1 int32, a2 int32, a3 *libc.WChar, a4 int32) int32 {
	var (
		v3 int32
		i  *uint32
		v5 int32
		v7 int32
		v8 int32
	)
	if a2 != 23 {
		if a2 == 0x4001 {
			nox_wcsncpy((*libc.WChar)(unsafe.Pointer(uintptr(a1+108))), a3, 63)
			*(*uint16)(unsafe.Pointer(uintptr(a1 + 234))) = 0
		} else if a2 == 0x4008 {
			v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 396))))
			if (int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 36)))) & 4) == 0 {
				if uintptr(unsafe.Pointer(a3)) == uintptr(1) {
					(*nox_window)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 52)))))).Func94(asWindowEvent(0x4007, a1, 0))
				}
				if v3 != 0 {
					for i = *(**uint32)(unsafe.Pointer(uintptr(v3 + 400))); i != nil; i = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(i), unsafe.Sizeof(uint32(0))*97))))) {
						if i != (*uint32)(unsafe.Pointer(uintptr(a1))) && *(*uint32)(unsafe.Add(unsafe.Pointer(i), unsafe.Sizeof(uint32(0))*10)) == *(*uint32)(unsafe.Pointer(uintptr(a1 + 40))) {
							*(*uint32)(unsafe.Add(unsafe.Pointer(i), unsafe.Sizeof(uint32(0))*9)) &= 0xFFFFFFFB
						}
					}
				}
				v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 36))))
				*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v5))), 0)) = uint8(int8(v5 | 4))
				*(*uint32)(unsafe.Pointer(uintptr(a1 + 36))) = uint32(v5)
				return 0
			}
		}
		return 0
	}
	if a3 == nil {
		v7 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 36))))
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v7))), 0)) = uint8(int8(v7 & 253))
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 36))) = uint32(v7)
	}
	v8 = (*nox_window)(unsafe.Pointer(uintptr(a1))).ID()
	(*nox_window)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 52)))))).Func94(asWindowEvent(0x4003, int32(uintptr(unsafe.Pointer(a3))), v8))
	return 1
}
func nox_xxx_loadDefColor_4A94A0() int32 {
	nox_color_black_2650656 = nox_color_rgb_4344A0(0, 0, 0)
	*memmap.PtrUint32(0x852978, 4) = nox_color_rgb_4344A0(8, 8, 8)
	*memmap.PtrUint32(0x85B3FC, 956) = nox_color_rgb_4344A0(115, 115, 115)
	*memmap.PtrUint32(6112660, 2597996) = nox_color_rgb_4344A0(212, 212, 212)
	nox_color_white_2523948 = nox_color_rgb_4344A0(math.MaxUint8, math.MaxUint8, math.MaxUint8)
	nox_color_violet_2598268 = nox_color_rgb_4344A0(100, 0, 0)
	*memmap.PtrUint32(0x85B3FC, 940) = nox_color_rgb_4344A0(math.MaxUint8, 0, 0)
	nox_color_red_2589776 = nox_color_rgb_4344A0(math.MaxUint8, 128, 128)
	*memmap.PtrUint32(0x85B3FC, 984) = nox_color_rgb_4344A0(0, 100, 0)
	*memmap.PtrUint32(0x8531A0, 2572) = nox_color_rgb_4344A0(0, math.MaxUint8, 0)
	nox_color_green_2614268 = nox_color_rgb_4344A0(128, math.MaxUint8, 128)
	*memmap.PtrUint32(0x85B3FC, 944) = nox_color_rgb_4344A0(0, 0, 140)
	nox_color_cyan_2649820 = nox_color_rgb_4344A0(0, 0, math.MaxUint8)
	nox_color_blue_2650684 = nox_color_rgb_4344A0(0, 160, math.MaxUint8)
	nox_color_orange_2614256 = nox_color_rgb_4344A0(240, 180, 42)
	nox_color_yellow_2589772 = nox_color_rgb_4344A0(math.MaxUint8, math.MaxUint8, 0)
	*memmap.PtrUint32(0x852978, 0) = nox_color_rgb_4344A0(math.MaxUint8, math.MaxUint8, 128)
	*memmap.PtrUint32(0x85B3FC, 132) = uint32(uintptr(unsafe.Pointer(&nox_color_black_2650656)))
	*memmap.PtrUint32(0x85B3FC, 136) = uint32(uintptr(memmap.PtrOff(0x852978, 4)))
	*memmap.PtrUint32(0x85B3FC, 140) = uint32(uintptr(memmap.PtrOff(0x85B3FC, 956)))
	*memmap.PtrUint32(0x85B3FC, 144) = uint32(uintptr(memmap.PtrOff(6112660, 2597996)))
	*memmap.PtrUint32(0x85B3FC, 148) = uint32(uintptr(unsafe.Pointer(&nox_color_white_2523948)))
	*memmap.PtrUint32(0x85B3FC, 152) = uint32(uintptr(unsafe.Pointer(&nox_color_violet_2598268)))
	*memmap.PtrUint32(0x85B3FC, 156) = uint32(uintptr(memmap.PtrOff(0x85B3FC, 940)))
	*memmap.PtrUint32(0x85B3FC, 160) = uint32(uintptr(unsafe.Pointer(&nox_color_red_2589776)))
	*memmap.PtrUint32(0x85B3FC, 164) = uint32(uintptr(memmap.PtrOff(0x85B3FC, 984)))
	*memmap.PtrUint32(0x85B3FC, 168) = uint32(uintptr(memmap.PtrOff(0x8531A0, 2572)))
	*memmap.PtrUint32(0x85B3FC, 172) = uint32(uintptr(unsafe.Pointer(&nox_color_green_2614268)))
	*memmap.PtrUint32(0x85B3FC, 176) = uint32(uintptr(memmap.PtrOff(0x85B3FC, 944)))
	*memmap.PtrUint32(0x85B3FC, 180) = uint32(uintptr(unsafe.Pointer(&nox_color_cyan_2649820)))
	*memmap.PtrUint32(0x85B3FC, 184) = uint32(uintptr(unsafe.Pointer(&nox_color_blue_2650684)))
	*memmap.PtrUint32(0x85B3FC, 188) = uint32(uintptr(unsafe.Pointer(&nox_color_orange_2614256)))
	*memmap.PtrUint32(0x85B3FC, 192) = uint32(uintptr(unsafe.Pointer(&nox_color_yellow_2589772)))
	*memmap.PtrUint32(0x85B3FC, 196) = uint32(uintptr(memmap.PtrOff(0x852978, 0)))
	return 1
}
func nox_xxx_loadPal_4A96C0_video_read_palette(a1 *byte) *File {
	var (
		result *File
		v2     *File
		v3     [8]byte
		v4     [12]byte
		v5     [12]byte
		v6     [32]byte
	)
	libc.StrCpy(&v4[0], CString("SHADEMAP"))
	libc.StrCpy(&v5[0], CString("TRANSMAP"))
	libc.StrCpy(&v3[0], CString("PALETTE"))
	if nox_video_modeXxx_3801780 != 1 {
		result = nox_fs_open(a1)
		v2 = result
		if result == nil {
			return result
		}
		if nox_binfile_fread_raw_40ADD0(&v6[0], 1, uint32(libc.StrLen(&v3[0])), result) != int32(libc.StrLen(&v3[0])) {
			nox_fs_close(v2)
			return nil
		}
		v6[libc.StrLen(&v3[0])] = 0
		if libc.StrCmp(&v6[0], &v3[0]) != 0 {
			nox_fs_close(v2)
			return nil
		}
		if nox_binfile_fread_raw_40ADD0((*byte)(memmap.PtrOff(6112660, 1308748)), 1, 768, v2) != 768 {
			nox_fs_close(v2)
			return nil
		}
		sub_4347F0((*byte)(memmap.PtrOff(6112660, 1308748)), 256)
		sub_4348C0()
		if sub_42FFF0(v2) == 0 {
			nox_fs_close(v2)
			return nil
		}
		if nox_binfile_fread_raw_40ADD0(&v6[0], 1, uint32(libc.StrLen(&v4[0])), v2) != int32(libc.StrLen(&v4[0])) {
			nox_fs_close(v2)
			return nil
		}
		v6[libc.StrLen(&v4[0])] = 0
		if libc.StrCmp(&v6[0], &v4[0]) != 0 {
			nox_fs_close(v2)
			return nil
		}
		if nox_binfile_fread_raw_40ADD0((*byte)(memmap.PtrOff(6112660, 2589804)), 1, 8192, v2) != 8192 {
			nox_fs_close(v2)
			return nil
		}
		sub_434610(int32(uintptr(memmap.PtrOff(6112660, 2589804))))
		if nox_binfile_fread_raw_40ADD0(&v6[0], 1, uint32(libc.StrLen(&v5[0])), v2) != int32(libc.StrLen(&v5[0])) {
			nox_fs_close(v2)
			return nil
		}
		v6[libc.StrLen(&v5[0])] = 0
		if libc.StrCmp(&v6[0], &v5[0]) != 0 {
			nox_fs_close(v2)
			return nil
		}
		if uint32(nox_binfile_fread_raw_40ADD0((*byte)(memmap.PtrOff(6112660, 2524236)), 1, 0x10000, v2)) != 0x10000 {
			nox_fs_close(v2)
			return nil
		}
		sub_434620(int32(uintptr(memmap.PtrOff(6112660, 2524236))))
		nox_fs_close(v2)
	}
	nox_xxx_loadDefColor_4A94A0()
	sub_4A9A10()
	return (*File)(unsafe.Pointer(uintptr(1)))
}
func sub_4A9A10() int32 {
	var i int32
	for i = 0; i < 6; i++ {
		nox_xxx_drawPlayer_4341D0(i, int32(nox_color_white_2523948))
	}
	return 1
}
func sub_4A9A30(a1 uint8, a2 uint8, a3 uint8) int32 {
	var (
		result int32
		v4     int32
		v5     *uint8
		v6     int32
		v7     int32
		v8     uint8
		v9     uint8
		v10    uint8
		v11    int32
	)
	result = 0
	v4 = 0x1000000
	v11 = 0
	v5 = (*uint8)(memmap.PtrOff(6112660, 1308749))
	for int32(*((*uint8)(unsafe.Add(unsafe.Pointer(v5), -1)))) != int32(a1) || int32(*v5) != int32(a2) || int32(*(*uint8)(unsafe.Add(unsafe.Pointer(v5), 1))) != int32(a3) {
		v6 = int32(*((*uint8)(unsafe.Add(unsafe.Pointer(v5), -1))))
		if v6-int32(a1) >= 0 {
			v10 = uint8(int8(v6 - int32(a1)))
		} else {
			v10 = uint8(int8(int32(a1) - v6))
		}
		if int32(*v5)-int32(a2) >= 0 {
			v9 = uint8(int8(int32(*v5) - int32(a2)))
		} else {
			v9 = uint8(int8(int32(a2) - int32(*v5)))
		}
		if int32(*(*uint8)(unsafe.Add(unsafe.Pointer(v5), 1)))-int32(a3) >= 0 {
			v8 = uint8(int8(int32(*(*uint8)(unsafe.Add(unsafe.Pointer(v5), 1))) - int32(a3)))
		} else {
			v8 = uint8(int8(int32(a3) - int32(*(*uint8)(unsafe.Add(unsafe.Pointer(v5), 1)))))
		}
		v7 = int32(v9) * int32(v9)
		if int32(v8)*int32(v8)+v7+int32(v10)*int32(v10) < v4 {
			v4 = int32(v8)*int32(v8) + v7 + int32(v10)*int32(v10)
			v11 = result
		}
		v5 = (*uint8)(unsafe.Add(unsafe.Pointer(v5), 3))
		result++
		if int32(uintptr(unsafe.Pointer(v5))) >= int32(uintptr(memmap.PtrOff(6112660, 1309517))) {
			return v11
		}
	}
	return result
}
func sub_4A9B20(a1 int32) int16 {
	var v1 int32
	if nox_video_modeXxx_3801780 == 1 {
		*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v1))), unsafe.Sizeof(uint16(0))*0)) = 0
	} else {
		v1 = (int32(*memmap.PtrUint8(6112660, uint32(a1*3)+1308750)) >> 3) | (int32(*memmap.PtrUint8(6112660, uint32(a1*3)+0x13F84D))&248|(int32(*memmap.PtrUint8(6112660, uint32(a1*3)+0x13F84C))&248)*32)*4
	}
	return int16(v1)
}
func sub_4A9C50(a1 int32) int32 {
	var v1 uint16
	*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v1))), unsafe.Sizeof(uint16(0))-1)) = *memmap.PtrUint8(6112660, uint32(a1*3)+1308750)
	*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v1))), 0)) = *memmap.PtrUint8(6112660, uint32(a1*3)+0x13F84D)
	return int32(*memmap.PtrUint8(6112660, uint32(a1*3)+0x13F84C)) | int32(v1)<<8
}
func nox_xxx_compassGenStrings_4A9C80() int32 {
	var (
		v0 int32
		v1 *uint8
		v2 int32
		v3 *uint8
		v5 [64]byte
	)
	*memmap.PtrUint32(6112660, 1309664) = 0
	v0 = 0
	v1 = (*uint8)(memmap.PtrOff(6112660, 1309644))
	for {
		nox_sprintf(&v5[0], CString("Compass%d"), func() int32 {
			p := &v0
			*p++
			return *p
		}())
		*(*uint32)(unsafe.Pointer(v1)) = uint32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg(&v5[0]))))
		v1 = (*uint8)(unsafe.Add(unsafe.Pointer(v1), 4))
		if int32(uintptr(unsafe.Pointer(v1))) >= int32(uintptr(memmap.PtrOff(6112660, 1309660))) {
			break
		}
	}
	v2 = 0
	v3 = (*uint8)(memmap.PtrOff(6112660, 1309516))
	for {
		nox_sprintf(&v5[0], CString("CompassMainArrow%d"), func() int32 {
			p := &v2
			*p++
			return *p
		}())
		*(*uint32)(unsafe.Pointer(v3)) = uint32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg(&v5[0]))))
		v3 = (*uint8)(unsafe.Add(unsafe.Pointer(v3), 4))
		if int32(uintptr(unsafe.Pointer(v3))) >= int32(uintptr(memmap.PtrOff(6112660, 1309644))) {
			break
		}
	}
	return 1
}
func nox_game_showOnlineOrLAN_413800() int32 {
	var (
		v3 *uint32
		v4 *uint32
		v5 *uint32
		v6 int32
		v7 *uint32
	)
	sub_44A400()
	gameAddStateCode(400)
	sub_4A1A40(1)
	if sub_40E0C0() != 0 && nox_xxx_check_flag_aaa_43AF70() == 1 {
		sub_41D4C0()
	}
	sub_43AF50(2)
	nox_win_onlineOrLAN_1309716 = uint32(uintptr(unsafe.Pointer(newWindowFromFile("ArnaMain.wnd", unsafe.Pointer(cgoFuncAddr(nox_xxx_windowArenaSub_4AA4D0))))))
	if nox_win_onlineOrLAN_1309716 == 0 {
		return 0
	}
	int32(nox_win_onlineOrLAN_1309716).setFunc93(sub_4A18E0)
	var wtop unsafe.Pointer = unsafe.Pointer((*nox_window)(unsafe.Pointer(uintptr(nox_win_onlineOrLAN_1309716))).ChildByID(410))
	int32(uintptr(wtop)).setFunc94(func(arg1 int32, arg2 int32, arg3 int32, arg4 int32) int32 {
		return nox_xxx_windowArenaSub_4AA4D0(arg1, uint32(arg2), (*int32)(unsafe.Pointer(uintptr(arg3))), arg4)
	})
	dword_5d4594_1309708 = uint32(uintptr(unsafe.Pointer(nox_gui_makeAnimation((*nox_window)(wtop), 0, 0, 0, -240, 0, 20, 0, -40))))
	if dword_5d4594_1309708 == 0 {
		return 0
	}
	*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1309708 + 0))) = 400
	*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1309708 + 48))) = uint32(cgoFuncAddr(sub_4AA450))
	*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1309708 + 56))) = uint32(cgoFuncAddr(sub_4AA490))
	var wbtn unsafe.Pointer = unsafe.Pointer((*nox_window)(unsafe.Pointer(uintptr(nox_win_onlineOrLAN_1309716))).ChildByID(420))
	int32(uintptr(wbtn)).setFunc94(func(arg1 int32, arg2 int32, arg3 int32, arg4 int32) int32 {
		return nox_xxx_windowArenaSub_4AA4D0(arg1, uint32(arg2), (*int32)(unsafe.Pointer(uintptr(arg3))), arg4)
	})
	*memmap.PtrUint32(6112660, 1309712) = uint32(uintptr(unsafe.Pointer(nox_gui_makeAnimation((*nox_window)(wbtn), 0, 240, 0, 480, 0, -20, 0, 40))))
	if *memmap.PtrUint32(6112660, 1309712) == 0 {
		return 0
	}
	sub_4A19F0(CString("OptsBack.wnd:Back"))
	nox_xxx_wndRetNULL_46A8A0()
	if noxflags.HasGame(noxflags.GameFlag26) {
		v3 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(uintptr(nox_win_onlineOrLAN_1309716))).ChildByID(411)))
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&nox_win_onlineOrLAN_1309716))))).Func94(asWindowEvent(0x4007, int32(uintptr(unsafe.Pointer(v3))), 0))
	} else if sub_4D6FA0() == 1 {
		v4 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(uintptr(nox_win_onlineOrLAN_1309716))).ChildByID(421)))
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&nox_win_onlineOrLAN_1309716))))).Func94(asWindowEvent(0x4007, int32(uintptr(unsafe.Pointer(v4))), 0))
	} else if sub_4D6FA0() == 2 {
		v5 = (*uint32)(unsafe.Pointer(GUIChildByID(150)))
		v6 = int32(uintptr(unsafe.Pointer(v5)))
		v7 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(v5)).ChildByID(152)))
		(*nox_window)(unsafe.Pointer(uintptr(v6))).Func94(asWindowEvent(0x4007, int32(uintptr(unsafe.Pointer(v7))), 0))
	}
	return 1
}
func sub_4AA450() int32 {
	*(*uint8)(unsafe.Pointer(uintptr(dword_5d4594_1309708 + 64))) = 2
	*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, 1309712) + 64))) = 2
	sub_43BE40(2)
	clientPlaySoundSpecial(923, 100)
	return 1
}
func sub_4AA490() int32 {
	var v0 func() int32
	v0 = cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1309708 + 52))), (*func() int32)(nil))
	nox_gui_freeAnimation_43C570((*nox_gui_animation)(*(*unsafe.Pointer)(unsafe.Pointer(&dword_5d4594_1309708))))
	nox_gui_freeAnimation_43C570((*nox_gui_animation)(*(*unsafe.Pointer)(memmap.PtrOff(6112660, 1309712))))
	(*nox_window)(unsafe.Pointer(uintptr(nox_win_onlineOrLAN_1309716))).Destroy()
	v0()
	return 1
}
func nox_game_showOptions_4AA6B0() int32 {
	var (
		v1  *uint32
		v2  *byte
		v3  int32
		v4  int32
		v5  uint32
		v6  *uint32
		v7  *byte
		v8  int32
		v9  int32
		v10 uint32
		v11 *uint32
		v12 *byte
		v13 int32
		v14 int32
		v15 uint32
		v16 *byte
		v17 *byte
		v18 *byte
		v19 *byte
		v20 *byte
		v21 *byte
	)
	gameAddStateCode(300)
	dword_5d4594_1309720 = uint32(uintptr(unsafe.Pointer(newWindowFromFile("Options.wnd", unsafe.Pointer(cgoFuncAddr(sub_4AABE0))))))
	if dword_5d4594_1309720 == 0 {
		return 0
	}
	if nox_client_advVideoOpts_New_4CB590((*nox_window)(unsafe.Pointer(uintptr(dword_5d4594_1309720)))) == 0 {
		return 0
	}
	(*(*int32)(unsafe.Pointer(&dword_5d4594_1309720))).setFunc93(sub_4A18E0)
	nox_draw_setTabWidth_43FE20(15)
	nox_wnd_xxx_1309740 = nox_gui_makeAnimation((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1309720)))), 0, 0, 0, -480, 0, 20, 0, -40)
	if nox_wnd_xxx_1309740 == nil {
		return 0
	}
	nox_wnd_xxx_1309740.field_0 = 300
	nox_wnd_xxx_1309740.field_12 = sub_4AA9C0
	nox_wnd_xxx_1309740.fnc_done_out = sub_4AAA10
	v1 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1309720)))).ChildByID(351)))
	*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*100)) + 8))) = 24
	*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*100)) + 12))) = 20
	v19 = (*byte)(unsafe.Pointer(nox_xxx_gLoadImg("OptionsVolumeSliderLit")))
	v16 = (*byte)(unsafe.Pointer(nox_xxx_gLoadImg("OptionsVolumeSliderLit")))
	v2 = (*byte)(unsafe.Pointer(nox_xxx_gLoadImg("OptionsVolumeSlider")))
	sub_4B5700(int32(uintptr(unsafe.Pointer(v1))), 0, 0, int32(uintptr(unsafe.Pointer(v2))), int32(uintptr(unsafe.Pointer(v16))), int32(uintptr(unsafe.Pointer(v19))))
	(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v1)))))).Func94(asWindowEvent(0x400B, 0, 0x4000))
	(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v1)))))).Func94(asWindowEvent(0x400A, int32(*(*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(dword_587000_127004)) + 4)))>>16), 0))
	dword_5d4594_1309728 = uint32(uintptr(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1309720)))).ChildByID(361))))
	v3 = sub_453070()
	v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1309728 + 36))))
	if v3 == 1 {
		v5 = uint32(v4 | 4)
	} else {
		v5 = uint32(v4) & 0xFFFFFFFB
	}
	*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1309728 + 36))) = v5
	v6 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1309720)))).ChildByID(352)))
	*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*100)) + 8))) = 24
	*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*100)) + 12))) = 20
	v20 = (*byte)(unsafe.Pointer(nox_xxx_gLoadImg("OptionsVolumeSliderLit")))
	v17 = (*byte)(unsafe.Pointer(nox_xxx_gLoadImg("OptionsVolumeSliderLit")))
	v7 = (*byte)(unsafe.Pointer(nox_xxx_gLoadImg("OptionsVolumeSlider")))
	sub_4B5700(int32(uintptr(unsafe.Pointer(v6))), 0, 0, int32(uintptr(unsafe.Pointer(v7))), int32(uintptr(unsafe.Pointer(v17))), int32(uintptr(unsafe.Pointer(v20))))
	(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v6)))))).Func94(asWindowEvent(0x400B, 0, 0x4000))
	(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v6)))))).Func94(asWindowEvent(0x400A, int32(*(*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(dword_587000_122852)) + 4)))>>16), 0))
	dword_5d4594_1309732 = uint32(uintptr(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1309720)))).ChildByID(362))))
	v8 = sub_44D990()
	v9 = int32(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1309732 + 36))))
	if v8 == 1 {
		v10 = uint32(v9 | 4)
	} else {
		v10 = uint32(v9) & 0xFFFFFFFB
	}
	*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1309732 + 36))) = v10
	v11 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1309720)))).ChildByID(353)))
	*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v11), unsafe.Sizeof(uint32(0))*100)) + 8))) = 24
	*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v11), unsafe.Sizeof(uint32(0))*100)) + 12))) = 20
	v21 = (*byte)(unsafe.Pointer(nox_xxx_gLoadImg("OptionsVolumeSliderLit")))
	v18 = (*byte)(unsafe.Pointer(nox_xxx_gLoadImg("OptionsVolumeSliderLit")))
	v12 = (*byte)(unsafe.Pointer(nox_xxx_gLoadImg("OptionsVolumeSlider")))
	sub_4B5700(int32(uintptr(unsafe.Pointer(v11))), 0, 0, int32(uintptr(unsafe.Pointer(v12))), int32(uintptr(unsafe.Pointer(v18))), int32(uintptr(unsafe.Pointer(v21))))
	(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v11)))))).Func94(asWindowEvent(0x400B, 0, 0x4000))
	(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v11)))))).Func94(asWindowEvent(0x400A, int32(*(*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(dword_587000_93164)) + 4)))>>16), 0))
	dword_5d4594_1309736 = uint32(uintptr(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1309720)))).ChildByID(363))))
	v13 = sub_43DC30()
	v14 = int32(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1309736 + 36))))
	if v13 == 1 {
		v15 = uint32(v14 | 4)
	} else {
		v15 = uint32(v14) & 0xFFFFFFFB
	}
	*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1309736 + 36))) = v15
	nox_xxx_wndRetNULL_46A8A0()
	sub_4A19F0(CString("OptsBack.wnd:Back"))
	sub_4A1A40(0)
	sub_4AAA70()
	return 1
}
func sub_4AAA70() *uint32 {
	var v0 *uint32
	_ = v0
	var v1 *uint32
	var v4 int32
	var result *uint32
	var v6 int32
	_ = v6
	var v7 int32
	nox_video_setMenuOptions((*nox_window)(unsafe.Pointer(uintptr(dword_5d4594_1309720))))
	if nox_video_getFullScreen() == 0 {
		v1 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1309720)))).ChildByID(331)))
	} else {
		v1 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1309720)))).ChildByID(332)))
	}
	(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v1)))))).Func94(asWindowEvent(0x4008, 1, 0))
	v4 = nox_video_getCutSize_4766D0()
	*memmap.PtrUint32(0x587000, 172884) = uint32(v4)
	if v4 > 69 {
		if v4 <= 79 {
			result = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1309720)))).ChildByID(312)))
			goto LABEL_22
		}
		if v4 <= 89 {
			result = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1309720)))).ChildByID(313)))
			goto LABEL_22
		}
		v7 = 314
	} else {
		v7 = 311
	}
	result = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1309720)))).ChildByID(v7)))
LABEL_22:
	if result != nil {
		result = (*uint32)(unsafe.Pointer(uintptr((*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(result)))))).Func94(asWindowEvent(0x4008, 1, 0)))))
	}
	return result
}
func sub_4AABE0(a1 int32, a2 int32, a3 *int32, a4 int32) int32 {
	var (
		v4     int32
		v5     int32
		result int32 = 0
		v7     int32
		v8     int32
		v9     func(uint32, int32, int32, int32)
		v10    func(uint32, int32, int32, int32)
		v11    func(uint32, int32, int32, int32)
		v12    unsafe.Pointer
		v13    func(unsafe.Pointer, int32, int32, int32)
		v14    func(uint32, int32, int32, int32)
		v15    int32
		v16    func(int32, int32, int32, int32)
		v17    int32
		v18    func(uint32, int32, int32, int32)
		v19    func(uint32, int32, int32, int32)
	)
	switch a2 {
	case 0x4005:
		clientPlaySoundSpecial(920, 100)
		result = 1
	case 0x4007:
		v4 = (*nox_window)(unsafe.Pointer(a3)).ID()
		if v4 >= 380 {
			return nox_gui_menu_proc_ext(v4)
		}
		if v4 > 341 {
			if v4 > 363 {
				if v4 == 2099 {
					sub_4AAA70()
				}
			} else if v4 == 363 {
				if sub_43DC30() == 1 {
					sub_43DC00()
				} else {
					sub_43DC10()
					sub_486320(*(**uint32)(unsafe.Pointer(&dword_587000_93164)), int32(*(*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(dword_587000_93164)) + 4)))>>16))
				}
			} else {
				v5 = v4 - 361
				if v5 != 0 {
					if v5 == 1 {
						if sub_44D990() == 1 {
							sub_44D960()
						} else {
							sub_44D970()
						}
					}
				} else if sub_453070() == 1 {
					sub_453050()
				} else {
					nox_xxx____setargv_9_453060()
				}
			}
		} else if v4 == 341 {
			sub_4AA9C0()
			nox_wnd_xxx_1309740.fnc_done_out = sub_4AB0C0
			nox_wnd_xxx_1309740.field_13 = sub_4CB880
		} else {
			switch v4 {
			case 311:
				nox_video_setCutSize_4766A0(65)
			case 312:
				nox_video_setCutSize_4766A0(75)
			case 313:
				nox_video_setCutSize_4766A0(85)
			case 314:
				nox_video_setCutSize_4766A0(100)
			case 332:
				nox_xxx_normalWndBits_587000_172880 = 16
				fallthrough
			case 333:
				nox_video_setFullScreen(1)
			case 331:
				nox_xxx_normalWndBits_587000_172880 = 8
				fallthrough
			case 334:
				nox_video_setFullScreen(0)
			default:
			}
		}
		clientPlaySoundSpecial(921, 100)
		result = 1
	case 0x4009:
		v7 = (*nox_window)(unsafe.Pointer(a3)).ID()
		v8 = int32(uintptr(unsafe.Pointer(nox_client_getWin1064916_46C720())))
		switch v7 {
		case 351:
			sub_486320(*(**uint32)(unsafe.Pointer(&dword_587000_127004)), a4)
			if v8 != 0 && *(**int32)(unsafe.Pointer(uintptr(v8 + 396))) == a3 {
				goto LABEL_76
			}
			if a4 != 0 {
				if sub_453070() != 0 {
					goto LABEL_75
				}
				v15 = int32(dword_5d4594_1309728)
				v16 = cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1309728 + 372))), (*func(int32, int32, int32, int32))(nil))
				if v16 == nil {
					goto LABEL_75
				}
				goto LABEL_74
			}
			if sub_453070() != 1 {
				goto LABEL_76
			}
			v14 = cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1309728 + 372))), (*func(uint32, int32, int32, int32))(nil))
			if v14 == nil {
				goto LABEL_76
			}
			v14(dword_5d4594_1309728, 21, 28, 2)
			result = 0
		case 352:
			sub_486320(*(**uint32)(unsafe.Pointer(&dword_587000_122852)), a4)
			if a4 != 0 {
				if sub_44D990() != 0 {
					goto LABEL_49
				}
				v12 = *(*unsafe.Pointer)(unsafe.Pointer(&dword_5d4594_1309732))
				v13 = cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1309732 + 372))), (*func(unsafe.Pointer, int32, int32, int32))(nil))
				if v13 == nil {
					goto LABEL_49
				}
				goto LABEL_48
			}
			if sub_44D990() != 1 {
				goto LABEL_76
			}
			v11 = cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1309732 + 372))), (*func(uint32, int32, int32, int32))(nil))
			if v11 == nil {
				goto LABEL_76
			}
			v11(dword_5d4594_1309732, 21, 28, 2)
			result = 0
		case 353:
			sub_486320(*(**uint32)(unsafe.Pointer(&dword_587000_93164)), a4)
			if a4 != 0 {
				if sub_43DC30() != 0 {
					goto LABEL_76
				}
				v10 = cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1309736 + 372))), (*func(uint32, int32, int32, int32))(nil))
				if v10 == nil {
					goto LABEL_76
				}
				v10(dword_5d4594_1309736, 21, 28, 2)
				result = 0
			} else {
				if sub_43DC30() != 1 {
					goto LABEL_76
				}
				v9 = cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1309736 + 372))), (*func(uint32, int32, int32, int32))(nil))
				if v9 == nil {
					goto LABEL_76
				}
				v9(dword_5d4594_1309736, 21, 28, 2)
				result = 0
			}
		case 316:
			nox_video_setGamma(float32(float64(a4)/50.0 + 0.5))
		case 318:
			nox_input_setSensitivity(math32.Pow(10.0, float32(float64(a4)/50.0-1.0)))
		default:
			goto LABEL_76
		}
	case 0x400C:
		v17 = (*nox_window)(unsafe.Pointer(a3)).ID() - 351
		if v17 != 0 {
			if v17 != 1 {
				goto LABEL_76
			}
			sub_486320(*(**uint32)(unsafe.Pointer(&dword_587000_122852)), a4)
			if a4 != 0 {
				if sub_44D990() == 0 {
					v12 = *(*unsafe.Pointer)(unsafe.Pointer(&dword_5d4594_1309732))
					v13 = cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1309732 + 372))), (*func(unsafe.Pointer, int32, int32, int32))(nil))
					if v13 != nil {
					LABEL_48:
						v13(v12, 21, 28, 2)
					}
				}
			LABEL_49:
				sub_4AA650()
				result = 0
			} else {
				if sub_44D990() != 1 {
					goto LABEL_76
				}
				v18 = cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1309732 + 372))), (*func(uint32, int32, int32, int32))(nil))
				if v18 == nil {
					goto LABEL_76
				}
				v18(dword_5d4594_1309732, 21, 28, 2)
				result = 0
			}
		} else {
			sub_486320(*(**uint32)(unsafe.Pointer(&dword_587000_127004)), a4)
			if a4 != 0 {
				if sub_453070() == 0 {
					v15 = int32(dword_5d4594_1309728)
					v16 = cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1309728 + 372))), (*func(int32, int32, int32, int32))(nil))
					if v16 != nil {
					LABEL_74:
						v16(v15, 21, 28, 2)
					}
				}
			LABEL_75:
				clientPlaySoundSpecial(768, 100)
				goto LABEL_76
			}
			if sub_453070() != 1 {
				goto LABEL_76
			}
			v19 = cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1309728 + 372))), (*func(uint32, int32, int32, int32))(nil))
			if v19 == nil {
				goto LABEL_76
			}
			v19(dword_5d4594_1309728, 21, 28, 2)
			result = 0
		}
	default:
	LABEL_76:
		result = 0
	}
	return result
}
func sub_4AB0C0() int32 {
	var v0 func() int32
	v0 = nox_wnd_xxx_1309740.field_13
	nox_gui_freeAnimation_43C570(nox_wnd_xxx_1309740)
	(*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1309720)))).Destroy()
	v0()
	return 1
}
func sub_4AB260() int32 {
	*memmap.PtrUint32(6112660, 1309752) = uint32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg("DisconnectIcon"))))
	dword_5d4594_1309756 = uint32(uintptr(unsafe.Pointer(nox_window_new(nil, 136, nox_win_width-50, nox_win_height/2+3, 50, 50, nil))))
	nox_xxx_wndSetIcon_46AE60(*(*int32)(unsafe.Pointer(&dword_5d4594_1309756)), *memmap.PtrInt32(6112660, 1309752))
	(*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1309756)))).SetAllFuncs(nil, func(arg1 *nox_window, arg2 unsafe.Pointer) int32 {
		return sub_4AB420(&arg1.id)
	}, nil)
	dword_5d4594_1309748 = uint32(uintptr(unsafe.Pointer(newWindowFromFile("discon.wnd", unsafe.Pointer(cgoFuncAddr(sub_4AB390))))))
	(*(*int32)(unsafe.Pointer(&dword_5d4594_1309748))).setFunc93(sub_4AB340)
	sub_46B120((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1309748)))), nil)
	(*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1309748)))).SetPos(image.Pt(int32(uint32(nox_win_width/2)-*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1309748 + 24)))/2), int32(uint32(nox_win_height/2)-*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1309748 + 28)))/2)))
	return 1
}
func sub_4AB340(a1 int32, a2 int32, a3 int32, a4 int32) int32 {
	if a2 != 21 {
		return 0
	}
	if a3 == 1 {
		return 1
	}
	if a3 == 57 {
		var mpos nox_point = getMousePos()
		(*nox_window)(unsafe.Pointer(uintptr(a1))).Func93(asWindowEvent(0x5, mpos.x|mpos.y<<16, 0))
	}
	return 0
}
func sub_4AB390(a1 int32, a2 int32, a3 *int32, a4 int32) int32 {
	var (
		v3     int32
		result int32
	)
	if a2 == 23 {
		return 1
	}
	if a2 != 0x4007 {
		return 0
	}
	v3 = (*nox_window)(unsafe.Pointer(a3)).ID() - 576
	if v3 == 0 {
		sub_43CF40()
		return 0
	}
	if v3 != 1 {
		return 0
	}
	sub_446380()
	if dword_5d4594_2650652 != 0 && sub_41E2F0() == 9 {
		sub_41F4B0()
		sub_41EC30()
		sub_446490(0)
		nox_xxx____setargv_4_44B000()
		sub_4AB4D0(0)
		result = 0
	} else {
		sub_43B750()
		sub_4AB4D0(0)
		result = 0
	}
	return result
}
func sub_4AB420(a1 *int32) int32 {
	var (
		v1 *int32
		v2 int32
		v4 int32
	)
	v1 = a1
	nox_client_wndGetPosition_46AA60((*nox_window)(unsafe.Pointer(a1)), (*uint32)(unsafe.Pointer(&a1)), (*uint32)(unsafe.Pointer(&v4)))
	v2 = *(*int32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(int32(0))*25))
	a1 = (*int32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(a1))), *(*int32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(int32(0))*24))))))
	nox_client_drawImageAt_47D2C0((*nox_video_bag_image_t)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(int32(0))*15))))), int32(uintptr(unsafe.Pointer(a1))), v2+v4)
	return 1
}
func sub_4AB470() int32 {
	var result int32
	(*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1309748)))).Destroy()
	(*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1309756)))).Destroy()
	result = 0
	dword_5d4594_1309756 = 0
	dword_5d4594_1309748 = 0
	return result
}
func sub_4AB4A0(a1 int32) int32 {
	var result int32
	if a1 != 0 {
		result = (*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1309756))))).Show()
	} else {
		result = (*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1309756))))).Hide()
	}
	return result
}
func sub_4AB4D0(a1 int32) int32 {
	var result int32
	if a1 != 0 {
		sub_44E040()
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1309748))))).Show()
		nox_xxx_wndShowModalMB((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1309748))))))
		sub_46C690((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1309748))))))
		guiFocus((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1309748))))))
		result = nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1309748))))), 1)
	} else {
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1309748))))).Hide()
		nox_xxx_wnd_46C6E0((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1309748))))))
		guiFocus(nil)
		result = nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1309748))))), 0)
	}
	return result
}
func sub_4ABDA0(a1 int32, a2 int16, a3 int16, a4 *uint32) int32 {
	var (
		v4     int32
		result int32
		v6     int32
		v7     uint32
		v8     int32
		v9     float32
		v10    int8
		v11    float32
		v12    int32
		v13    [2]int32
		v14    float32
		v15    int32
		v16    int32
		v17    float32
		v18    int32
		v19    [256]byte
	)
	cryptFileReadWrite((*uint8)(unsafe.Pointer(&v15))[:4])
	*a4 += 4
	cryptFileReadWrite((*uint8)(unsafe.Pointer(&v16))[:4])
	*a4 += 4
	if int32(a3) < 40 || int32(a2) < 4 {
		cryptFileReadWrite((*uint8)(unsafe.Pointer(&v13[0]))[:8])
		*a4 += 8
	} else {
		cryptFileReadWrite((*uint8)(unsafe.Pointer(&v11))[:4])
		*a4 += 4
		cryptFileReadWrite((*uint8)(unsafe.Pointer(&v17))[:4])
		v9 = v11
		*a4 += 4
		v13[0] = int(v9)
		v13[1] = int(v17)
	}
	if int32(a3) >= 10 {
		cryptFileReadWrite((*uint8)(unsafe.Pointer(&v12))[:1])
		*a4++
		cryptFileReadWrite((*uint8)(unsafe.Pointer(&v19[0]))[:uint32(uint8(int8(v12)))])
		*a4 += uint32(uint8(int8(v12)))
	}
	if int32(a3) >= 20 {
		cryptFileReadWrite((*uint8)(unsafe.Pointer(&v10))[:1])
		*a4++
	}
	if int32(a3) >= 30 {
		cryptFileReadWrite((*uint8)(unsafe.Pointer(&v12))[:1])
		*a4++
	}
	if int32(a3) >= 40 {
		nox_xxx_cryptSeekCur(4)
		*a4 += 4
		if int32(a2) >= 2 {
			if int32(a2) < 5 {
				cryptFileReadWrite((*uint8)(unsafe.Pointer(&v14))[:4])
				*a4 += 4
				v11 = v14
			} else {
				cryptFileReadWrite((*uint8)(unsafe.Pointer(&v11))[:2])
				*a4 += 2
			}
			v4 = int32(uint16(int16(int32(*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v11))), unsafe.Sizeof(uint16(0))*0))) * 4)))
			nox_xxx_cryptSeekCur(v4)
			*a4 += uint32(v4)
		}
		if int32(a2) >= 3 {
			cryptFileReadWrite((*uint8)(unsafe.Pointer(&v18))[:4])
			*a4 += 4
		}
	}
	result = int32(uintptr(unsafe.Pointer(nox_xxx_spriteLoadAdd_45A360_drawable(a1, v13[0], v13[1]))))
	v6 = result
	if result != 0 {
		v7 = *(*uint32)(unsafe.Pointer(uintptr(result + 120))) & 0xEEBF7E9D
		*(*uint32)(unsafe.Pointer(uintptr(result + 120))) = v7
		*(*uint32)(unsafe.Pointer(uintptr(result + 120))) = uint32(v16) | v7
		v8 = int32(*(*uint32)(unsafe.Pointer(uintptr(result + 280))))
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v8))), 0)) = uint8(int8(v8 & 161))
		*(*uint32)(unsafe.Pointer(uintptr(v6 + 280))) = uint32(v8)
		*(*uint32)(unsafe.Pointer(uintptr(v6 + 280))) = uint32(v18 | v8)
		*(*uint32)(unsafe.Pointer(uintptr(v6 + 128))) = uint32(v15)
		*(*uint8)(unsafe.Pointer(uintptr(v6 + 28))) = uint8(v10)
		if !noxflags.HasGame(noxflags.GameHost) && int32(*(*uint8)(unsafe.Pointer(uintptr(v6 + 28)))) != 0 && ((*(*uint32)(unsafe.Pointer(uintptr(v6 + 112)))&0x10000000) == 0 || !noxflags.HasGame(noxflags.GameModeChat)) {
			nox_xxx_createAtImpl_4191D0(*(*uint8)(unsafe.Pointer(uintptr(v6 + 28))), unsafe.Pointer(uintptr(v6+24)), 0, int32(*(*uint32)(unsafe.Pointer(uintptr(v6 + 128)))), 0)
		}
		result = v6
	}
	return result
}
func nox_xxx_spriteLoadFromMap_4AC020(thingInd int32, a2 int16, a3 *uint32) int32 {
	var (
		v3  *uint32
		v4  int16
		v5  int16
		v6  bool
		v7  int32
		v10 bool
		v11 int32
		v12 int32
		v13 int8
		v14 int32
		v15 int32
		v16 int32
		v17 int32
		v18 uint8
		v19 int32
		v20 float32
		v21 int32
		v22 int32
		v23 int32
		v24 int32
		v25 int32
		v26 int32
		v27 int32
		v28 int32
		v29 int32
		v30 float32
		v31 float32
	)
	v3 = a3
	v4 = a2
	v5 = 1
	v23 = 1
	if int32(a2) < 40 {
		return sub_4ABDA0(thingInd, v5, v4, v3)
	}
	cryptFileReadWrite((*uint8)(unsafe.Pointer(&v23))[:2])
	v5 = int16(v23)
	v6 = int32(int16(v23)) < 61
	*v3 += 2
	if v6 {
		return sub_4ABDA0(thingInd, v5, v4, v3)
	}
	cryptFileReadWrite((*uint8)(unsafe.Pointer(&v25))[:4])
	*v3 += 4
	cryptFileReadWrite((*uint8)(unsafe.Pointer(&v29))[:4])
	*v3 += 4
	cryptFileReadWrite((*uint8)(unsafe.Pointer(&v30))[:4])
	*v3 += 4
	cryptFileReadWrite((*uint8)(unsafe.Pointer(&v31))[:4])
	v20 = v31
	*v3 += 4
	v22 = int(v20)
	v7 = int(v30)
	var v9 int32 = int32(uintptr(unsafe.Pointer(nox_xxx_spriteLoadAdd_45A360_drawable(thingInd, v7, v22))))
	if v9 == 0 {
		return 0
	}
	*(*uint32)(unsafe.Pointer(uintptr(v9 + 128))) = uint32(v25)
	cryptFileReadWrite((*uint8)(unsafe.Pointer(&a2))[:1])
	v10 = int32(uint8(int8(a2))) == 0
	*v3++
	if !v10 {
		cryptFileReadWrite((*uint8)(unsafe.Pointer(&v26))[:4])
		v11 = v26
		*v3 += 4
		*(*uint32)(unsafe.Pointer(uintptr(v9 + 120))) = uint32(v11) | (*(*uint32)(unsafe.Pointer(uintptr(v9 + 120))) & 0xEEBF7E9D)
		cryptFileReadWrite((*uint8)(unsafe.Pointer(&a3))[:1])
		v12 = int32(uint8(uintptr(unsafe.Pointer(a3))))
		v21 = int32(uint8(uintptr(unsafe.Pointer(a3))))
		*v3++
		nox_xxx_cryptSeekCur(v21)
		*v3 += uint32(v12)
		cryptFileReadWrite((*uint8)(unsafe.Pointer(&thingInd))[:1])
		v13 = int8(thingInd)
		*v3++
		*(*uint8)(unsafe.Pointer(uintptr(v9 + 28))) = uint8(v13)
		cryptFileReadWrite((*uint8)(unsafe.Pointer(&a3))[:1])
		*v3++
		cryptFileReadWrite((*uint8)(unsafe.Pointer(&v27))[:2])
		v14 = int32(uint16(int16(v27))) * 4
		v19 = int32(uint16(int16(v27))) * 4
		*v3 += 2
		nox_xxx_cryptSeekCur(v19)
		*v3 += uint32(v14)
		cryptFileReadWrite((*uint8)(unsafe.Pointer(&v28))[:4])
		v15 = v28
		*v3 += 4
		v16 = int32(*(*uint32)(unsafe.Pointer(uintptr(v9 + 280))))
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v16))), 0)) = uint8(int8(v16 & 161))
		*(*uint32)(unsafe.Pointer(uintptr(v9 + 280))) = uint32(v15 | v16)
		if int32(int16(v23)) >= 63 {
			nox_xxx_cryptSeekCur(2)
			*v3 += 2
			cryptFileReadWrite((*uint8)(unsafe.Pointer(&v24))[:4])
			*v3 += 4
			nox_xxx_cryptSeekCur(v24)
			*v3 += uint32(v24)
			nox_xxx_cryptSeekCur(4)
			*v3 += 4
		}
		if int32(int16(v23)) >= 64 {
			nox_xxx_cryptSeekCur(4)
			*v3 += 4
		}
	}
	if !noxflags.HasGame(noxflags.GameHost) {
		v17 = int32(*memmap.PtrUint32(6112660, 1309788))
		if *memmap.PtrUint32(6112660, 1309788) == 0 {
			v17 = nox_xxx_getTTByNameSpriteMB_44CFC0("FlagMarker")
			*memmap.PtrUint32(6112660, 1309788) = uint32(v17)
		}
		v18 = *(*uint8)(unsafe.Pointer(uintptr(v9 + 28)))
		if int32(v18) != 0 {
			if *(*uint32)(unsafe.Pointer(uintptr(v9 + 108))) != uint32(v17) {
				nox_xxx_createAtImpl_4191D0(v18, unsafe.Pointer(uintptr(v9+24)), 0, int32(*(*uint32)(unsafe.Pointer(uintptr(v9 + 128)))), 0)
			}
		}
	}
	return v9
}
func nox_client_mapSpecialRWObjectData_4AC610() int32 {
	var (
		v1 uint16
		v2 uint16
		v3 int32
		v4 int32
		v5 int32
		v6 int32
	)
	v6 = 1
	cryptFileReadWrite((*uint8)(unsafe.Pointer(&v6))[:2])
	if int32(int16(v6)) > 1 {
		return 0
	}
	if nox_xxx_cryptGetXxx() == 0 {
		return 0
	}
	cryptFileReadWrite((*uint8)(unsafe.Pointer(&v4))[:2])
	if int32(uint16(int16(v4))) != 0 {
		for {
			cryptFileReadMaybeAlign((*uint8)(unsafe.Pointer(&v5))[:4])
			v1 = uint16(nox_xxx_objectTOCgetTT_42C2B0(int16(v4)))
			v2 = v1
			if int32(v1) == 0 {
				break
			}
			if sub_44D090(int32(v1)) != 0 {
				v3 = nox_xxx_clientLoadSomeObject_4AC6E0(v2)
				v5 -= v3
			}
			if v5 > 0 {
				nox_xxx_cryptSeekCur(v5)
			}
			cryptFileReadWrite((*uint8)(unsafe.Pointer(&v4))[:2])
			if int32(uint16(int16(v4))) == 0 {
				return 1
			}
		}
		return 0
	}
	return 1
}
func nox_xxx_clientLoadSomeObject_4AC6E0(a1 uint16) int32 {
	if *memmap.PtrUint32(6112660, 1309792) == 0 {
		*memmap.PtrUint32(6112660, 1309792) = uint32(nox_xxx_getTTByNameSpriteMB_44CFC0("ColorLight"))
		*memmap.PtrUint32(6112660, 1309796) = uint32(nox_xxx_getTTByNameSpriteMB_44CFC0("ColorLightMovable"))
		*memmap.PtrUint32(6112660, 1309800) = uint32(nox_xxx_getTTByNameSpriteMB_44CFC0("TeamBase"))
		*memmap.PtrUint32(6112660, 1309804) = uint32(nox_xxx_getTTByNameSpriteMB_44CFC0("PressurePlate"))
	}
	if uint32(a1) == *memmap.PtrUint32(6112660, 1309792) || uint32(a1) == *memmap.PtrUint32(6112660, 1309796) {
		return nox_xxx_colorLightClientLoad_4AC980(int32(a1))
	}
	if uint32(a1) == *memmap.PtrUint32(6112660, 1309800) {
		return nox_xxx_cliLoadTeamBase_4ACE00(int32(a1))
	}
	if uint32(a1) == *memmap.PtrUint32(6112660, 1309804) {
		return sub_4ACEF0(int32(a1))
	}
	if sub_44D060(int32(a1)) != 0 {
		return sub_4AD040(int32(a1))
	}
	if sub_44D040(int32(a1)) != 0 {
		return sub_4AC7B0(int32(a1))
	}
	return 0
}
func sub_4AC7B0(a1 int32) int32 {
	var (
		v1     int32
		v2     int32
		result int32
		v4     int32
		v5     int8
		v6     float32
		v7     int32
		v9     float64
		v10    float64
		v11    int32
		v12    int32
		v13    float32
		v14    int32
	)
	cryptFileReadWrite((*uint8)(unsafe.Pointer(&v12))[:2])
	v11 = 2
	v1 = nox_xxx_spriteLoadFromMap_4AC020(a1, int16(v12), (*uint32)(unsafe.Pointer(&v11)))
	v2 = v1
	if v1 != 0 {
		v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 112))))
		if (v4 & 128) == 0 {
			if v4&512 != 0 {
				cryptFileReadWrite((*uint8)(unsafe.Pointer(&a1))[:4])
				cryptFileReadWrite((*uint8)(unsafe.Pointer(&v14))[:4])
				v9 = float64(a1)
				v11 += 8
				v13 = float32(v9)
				*(*float32)(unsafe.Pointer(uintptr(v2 + 56))) = float32(v9)
				v10 = float64(v14)
				*(*float32)(unsafe.Pointer(uintptr(v2 + 60))) = float32(v10)
				if float64(v13) > 60.0 {
					*(*uint32)(unsafe.Pointer(uintptr(v2 + 56))) = 0x42700000
				}
				if v10 > 60.0 {
					*(*uint32)(unsafe.Pointer(uintptr(v2 + 60))) = 0x42700000
				}
				nox_shape_box_calc((*nox_shape)(unsafe.Pointer(uintptr(v2 + 44))))
			}
		} else {
			cryptFileReadWrite((*uint8)(unsafe.Pointer(&v13))[:4])
			v11 += 4
			cryptFileReadWrite((*uint8)(unsafe.Pointer(&v14))[:4])
			v5 = int8(v14)
			v11 += 4
			*(*uint8)(unsafe.Pointer(uintptr(v2 + 433))) = uint8(int8(v14))
			*(*uint8)(unsafe.Pointer(uintptr(v2 + 432))) = 0
			if int32(v5) != 0 {
				*(*uint8)(unsafe.Pointer(uintptr(v2 + 432))) = 1
			}
			v6 = v13
			a1 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v13))), unsafe.Sizeof(uint32(0))*0)))
			if int32(int16(v12)) >= 41 {
				cryptFileReadWrite((*uint8)(unsafe.Pointer(&a1))[:4])
				v6 = *(*float32)(unsafe.Pointer(&a1))
				v11 += 4
			}
			*(*uint8)(unsafe.Pointer(uintptr(v2 + 299))) = *(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v13))), 0))
			v7 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 12))) + uint32(*memmap.PtrInt32(0x587000, (*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v6))), unsafe.Sizeof(uint32(0))*0)))*8+0x2FE58)/2))
			sub_410390(v2, v7/23, int32((*(*uint32)(unsafe.Pointer(uintptr(v2 + 16)))+uint32(*memmap.PtrInt32(0x587000, (*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v6))), unsafe.Sizeof(uint32(0))*0)))*8+0x2FE5C)/2))/23))
		}
		*(*uint32)(unsafe.Pointer(uintptr(v2 + 288))) = 0
		nox_xxx_spriteSetActiveMB_45A990_drawable(v2)
		result = v11
	} else {
		nox_xxx_spriteLoadError_4356E0()
		result = 0
	}
	return result
}
func nox_xxx_colorLightClientLoad_4AC980(a1 int32) int32 {
	var (
		v1 int32
		v2 int32
		v4 *uint16
		v5 *uint8
		v6 *uint8
		v7 int32
		v8 int32
	)
	cryptFileReadWrite((*uint8)(unsafe.Pointer(&v8))[:2])
	v7 = 2
	v1 = nox_xxx_spriteLoadFromMap_4AC020(a1, int16(v8), (*uint32)(unsafe.Pointer(&v7)))
	v2 = v1
	if v1 == 0 {
		nox_xxx_spriteLoadError_4356E0()
		return 0
	}
	if int32(int16(v8)) >= 2 {
		cryptFileReadWrite((*uint8)(unsafe.Pointer(uintptr(v1 + 136)))[:4])
		v7 += 4
		cryptFileReadWrite((*uint8)(unsafe.Pointer(uintptr(v2 + 140)))[:4])
		v7 += 4
		cryptFileReadWrite((*uint8)(unsafe.Pointer(uintptr(v2 + 144)))[:4])
		v7 += 4
		cryptFileReadWrite((*uint8)(unsafe.Pointer(uintptr(v2 + 148)))[:4])
		v7 += 4
		cryptFileReadWrite((*uint8)(unsafe.Pointer(uintptr(v2 + 152)))[:12])
		v4 = (*uint16)(unsafe.Pointer(uintptr(v2 + 164)))
		v7 += 12
		cryptFileReadWrite((*uint8)(unsafe.Pointer(uintptr(v2 + 164)))[:2])
		v7 += 2
		cryptFileReadWrite((*uint8)(unsafe.Pointer(uintptr(v2 + 166)))[:2])
		v7 += 2
		cryptFileReadWrite((*uint8)(unsafe.Pointer(uintptr(v2 + 168)))[:4])
		v7 += 4
		cryptFileReadWrite((*uint8)(unsafe.Pointer(uintptr(v2 + 176)))[:2])
		v7 += 2
		cryptFileReadWrite((*uint8)(unsafe.Pointer(uintptr(v2 + 178)))[:48])
		v7 += 48
		cryptFileReadWrite((*uint8)(unsafe.Pointer(uintptr(v2 + 226)))[:16])
		v7 += 16
		cryptFileReadWrite((*uint8)(unsafe.Pointer(uintptr(v2 + 242)))[:16])
		v7 += 16
		cryptFileReadWrite((*uint8)(unsafe.Pointer(uintptr(v2 + 258)))[:2])
		v7 += 2
		cryptFileReadWrite((*uint8)(unsafe.Pointer(uintptr(v2 + 260)))[:2])
		v7 += 2
		cryptFileReadWrite((*uint8)(unsafe.Pointer(uintptr(v2 + 262)))[:2])
		v7 += 2
		cryptFileReadWrite((*uint8)(unsafe.Pointer(uintptr(v2 + 264)))[:4])
		v7 += 4
		cryptFileReadWrite((*uint8)(unsafe.Pointer(uintptr(v2 + 270)))[:2])
		v7 += 2
		cryptFileReadWrite((*uint8)(unsafe.Pointer(uintptr(v2 + 272)))[:2])
		v7 += 2
		cryptFileReadWrite((*uint8)(unsafe.Pointer(uintptr(v2 + 274)))[:1])
		v7++
		if int32(int16(v8)) > 40 {
			if int32(int16(v8)) >= 42 {
				cryptFileReadWrite((*uint8)(unsafe.Pointer(uintptr(v2 + 172)))[:4])
				v7 += 4
			} else {
				cryptFileReadWrite((*uint8)(unsafe.Pointer(&a1))[:1])
				v7++
				*(*uint32)(unsafe.Pointer(uintptr(v2 + 172))) = uint32(uint8(int8(a1)))
			}
		}
	} else {
		cryptFileReadWrite((*uint8)(unsafe.Pointer(uintptr(v1 + 136)))[:4])
		v7 += 4
		cryptFileReadWrite((*uint8)(unsafe.Pointer(uintptr(v2 + 140)))[:4])
		v7 += 4
		cryptFileReadWrite((*uint8)(unsafe.Pointer(uintptr(v2 + 144)))[:4])
		v7 += 4
		cryptFileReadWrite((*uint8)(unsafe.Pointer(uintptr(v2 + 148)))[:4])
		v7 += 4
		cryptFileReadWrite((*uint8)(unsafe.Pointer(uintptr(v2 + 152)))[:12])
		v4 = (*uint16)(unsafe.Pointer(uintptr(v2 + 164)))
		v7 += 12
		cryptFileReadWrite((*uint8)(unsafe.Pointer(uintptr(v2 + 164)))[:2])
		v7 += 2
		cryptFileReadWrite((*uint8)(unsafe.Pointer(uintptr(v2 + 166)))[:2])
		v7 += 2
		cryptFileReadWrite((*uint8)(unsafe.Pointer(uintptr(v2 + 168)))[:4])
		v7 += 4
		*(*uint16)(unsafe.Pointer(uintptr(v2 + 176))) = 0
		*(*uint16)(unsafe.Pointer(uintptr(v2 + 258))) = 0
		*(*uint16)(unsafe.Pointer(uintptr(v2 + 260))) = 0
		*(*uint16)(unsafe.Pointer(uintptr(v2 + 262))) = 0
		*(*uint32)(unsafe.Pointer(uintptr(v2 + 264))) = 0
		*(*uint16)(unsafe.Pointer(uintptr(v2 + 270))) = 0
		*(*uint8)(unsafe.Pointer(uintptr(v2 + 274))) = 128
		if nox_xxx_cryptGetXxx() != 1 {
			goto LABEL_24
		}
		if float64(*(*float32)(unsafe.Pointer(uintptr(v2 + 140)))) <= 63.0 && float64(*(*int32)(unsafe.Pointer(uintptr(v2 + 148))))**mem_getDoublePtr(0x581450, 9752) <= *mem_getDoublePtr(0x581450, 9744) {
		LABEL_13:
			*(*uint8)(unsafe.Pointer(uintptr(v2 + 432))) = 0
			*(*uint8)(unsafe.Pointer(uintptr(v2 + 433))) = 0
			*(*uint8)(unsafe.Pointer(uintptr(v2 + 434))) = 0
			v5 = (*uint8)(unsafe.Pointer(uintptr(v2 + 242)))
			v6 = (*uint8)(unsafe.Pointer(uintptr(v2 + 179)))
			for {
				if int32(*((*uint8)(unsafe.Add(unsafe.Pointer(v6), -1)))) != 0 || int32(*v6) != 0 || int32(*(*uint8)(unsafe.Add(unsafe.Pointer(v6), 1))) != 0 {
					*(*uint8)(unsafe.Pointer(uintptr(v2 + 432)))++
				}
				if int32(*((*uint8)(unsafe.Add(unsafe.Pointer(v5), -16)))) != 0 {
					*(*uint8)(unsafe.Pointer(uintptr(v2 + 433)))++
				}
				if int32(*v5) != 0 {
					*(*uint8)(unsafe.Pointer(uintptr(v2 + 434)))++
				}
				v6 = (*uint8)(unsafe.Add(unsafe.Pointer(v6), 3))
				v5 = (*uint8)(unsafe.Add(unsafe.Pointer(v5), 1))
				if int32(uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v5), int32(-242-v2)))))) >= 16 {
					break
				}
			}
			a1 = int32(*v4)
			*(*uint16)(unsafe.Pointer(uintptr(v2 + 268))) = uint16(int16(int64(float64(a1) * *mem_getDoublePtr(0x581450, 9752) * *mem_getDoublePtr(0x581450, 9736))))
			goto LABEL_24
		}
		sub_484CE0(v2+136, 63.0)
	}
	if nox_xxx_cryptGetXxx() == 1 {
		goto LABEL_13
	}
LABEL_24:
	*(*uint32)(unsafe.Pointer(uintptr(v2 + 288))) = 0
	nox_xxx_spriteSetActiveMB_45A990_drawable(v2)
	return v7
}
func nox_xxx_cliLoadTeamBase_4ACE00(a1 int32) int32 {
	var (
		v1     int32
		v2     int32
		result int32
		i      int32
		v5     int32
		v6     int32
		v7     int32
		v8     int32
		v9     [256]byte
	)
	cryptFileReadWrite((*uint8)(unsafe.Pointer(&v8))[:2])
	v6 = 2
	v1 = nox_xxx_spriteLoadFromMap_4AC020(a1, int16(v8), (*uint32)(unsafe.Pointer(&v6)))
	v2 = v1
	if v1 != 0 {
		*(*uint32)(unsafe.Pointer(uintptr(v1 + 288))) = 0
		nox_xxx_spriteSetActiveMB_45A990_drawable(v1)
		for i = 0; i < 16; i += 4 {
			cryptFileReadWrite((*uint8)(unsafe.Pointer(&v7))[:1])
			v6++
			cryptFileReadWrite((*uint8)(unsafe.Pointer(&v9[0]))[:uint32(uint8(int8(v7)))])
			v6 += int32(uint8(int8(v7)))
			v9[uint8(int8(v7))] = 0
			v5 = nox_xxx_modifGetIdByName_413290(&v9[0])
			*(*uint32)(unsafe.Pointer(uintptr(i + v2 + 432))) = uint32(uintptr(unsafe.Pointer(nox_xxx_modifGetDescById_413330(v5))))
			*(*uint16)(unsafe.Pointer(uintptr(v2 + 448))) = math.MaxUint16
			*(*uint16)(unsafe.Pointer(uintptr(v2 + 450))) = math.MaxUint16
		}
		result = v6
	} else {
		nox_xxx_spriteLoadError_4356E0()
		result = 0
	}
	return result
}
func sub_4ACEF0(a1 int32) int32 {
	var (
		v1     int32
		v2     int32
		result int32
		v4     *uint8
		v5     int32
		v6     int32
		v7     int32
		v8     *uint8
	)
	cryptFileReadWrite((*uint8)(unsafe.Pointer(&v7))[:2])
	v5 = 2
	v1 = nox_xxx_spriteLoadFromMap_4AC020(a1, int16(v7), (*uint32)(unsafe.Pointer(&v5)))
	v2 = v1
	if v1 != 0 {
		v4 = (*uint8)(unsafe.Pointer(uintptr(v1 + 432)))
		*(*uint32)(unsafe.Pointer(uintptr(v1 + 288))) = 0
		nox_xxx_spriteSetActiveMB_45A990_drawable(v1)
		cryptFileReadWrite((*uint8)(unsafe.Pointer(&v6))[:4])
		*(*float32)(unsafe.Pointer(uintptr(v2 + 56))) = float32(float64(v6))
		v5 += 4
		cryptFileReadWrite((*uint8)(unsafe.Pointer(&v6))[:4])
		*(*float32)(unsafe.Pointer(uintptr(v2 + 60))) = float32(float64(v6))
		v5 += 4
		nox_shape_box_calc((*nox_shape)(unsafe.Pointer(uintptr(v2 + 44))))
		*(*uint8)(unsafe.Add(unsafe.Pointer(v4), 4)) = 10
		*v4 = 90
		*(*uint8)(unsafe.Add(unsafe.Pointer(v4), 1)) = 90
		*(*uint8)(unsafe.Add(unsafe.Pointer(v4), 2)) = 90
		*(*uint8)(unsafe.Add(unsafe.Pointer(v4), 3)) = 10
		*(*uint8)(unsafe.Add(unsafe.Pointer(v4), 5)) = 10
		v8 = (*uint8)(unsafe.Add(unsafe.Pointer(v4), 5))
		if int32(int16(v7)) >= 41 {
			cryptFileReadWrite(v4[:1])
			v5++
			cryptFileReadWrite((*uint8)(unsafe.Add(unsafe.Pointer(v4), 1))[:1])
			v5++
			cryptFileReadWrite((*uint8)(unsafe.Add(unsafe.Pointer(v4), 2))[:1])
			v5++
			cryptFileReadWrite((*uint8)(unsafe.Add(unsafe.Pointer(v4), 3))[:1])
			v5++
			cryptFileReadWrite((*uint8)(unsafe.Add(unsafe.Pointer(v4), 4))[:1])
			v5++
			cryptFileReadWrite(v8[:1])
			v5++
		}
		result = v5
	} else {
		nox_xxx_spriteLoadError_4356E0()
		result = 0
	}
	return result
}
func sub_4AD040(a1 int32) int32 {
	var (
		v1     int32
		result int32
		v3     int32
		v4     int32
		v5     int32
	)
	cryptFileReadWrite((*uint8)(unsafe.Pointer(&v4))[:2])
	v3 = 2
	v1 = nox_xxx_spriteLoadFromMap_4AC020(a1, int16(v4), (*uint32)(unsafe.Pointer(&v3)))
	if v1 != 0 {
		if int32(int16(v4)) >= 61 {
			*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a1))), 0)) = 0
			cryptFileReadWrite((*uint8)(unsafe.Pointer(&v5))[:4])
			v3 += 4
			cryptFileReadWrite((*uint8)(unsafe.Pointer(&a1))[:1])
			v3++
			if int32(uint8(int8(a1))) == 1 {
				sub_459DD0((*nox_drawable)(unsafe.Pointer(uintptr(v1))), 1)
			}
		}
		*(*uint32)(unsafe.Pointer(uintptr(v1 + 288))) = 0
		nox_xxx_spriteSetActiveMB_45A990_drawable(v1)
		result = v3
	} else {
		nox_xxx_spriteLoadError_4356E0()
		result = 0
	}
	return result
}
func sub_4AD570() int32 {
	var (
		v1   *uint32
		mpos nox_point = getMousePos()
	)
	if nox_xxx_check_flag_aaa_43AF70() == 1 {
		v1 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1309812)))).ChildByID(0x284D)))
		if (*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v1)))))).Flags().IsHidden() == 0 && !nox_xxx_wndPointInWnd_46AAB0(v1, mpos.x, mpos.y) {
			nox_xxx_wndClearCaptureMain((*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v1)))))))
			(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v1)))))).Hide()
		}
	}
	return 1
}
func nox_xxx_windowServerOptionsGeneralProc_4AD5D0(a1 int32, a2 int32, a3 *int32, a4 int32) int32 {
	var (
		v4     *uint32
		v5     int32
		v6     int32
		v7     int32
		v8     int32
		v9     int32
		v10    *uint32
		result int32
		v12    int32
		v13    *uint32
	)
	if a2 == 0x4007 {
		v12 = (*nox_window)(unsafe.Pointer(a3)).ID()
		clientPlaySoundSpecial(766, 100)
		switch v12 {
		case 0x283D:
			if nox_server_doPlayersAutoRespawn_40A5F0() != 0 {
				nox_xxx_ruleSetNoRespawn_40A5E0(0)
			} else {
				nox_xxx_ruleSetNoRespawn_40A5E0(1)
			}
			result = 0
		case 0x283E:
			nox_server_sendMotd_108752 ^= 1
			result = 0
		case 0x2840:
			if sub_4D0D70() != 0 {
				sub_4D0D90(0)
			} else {
				sub_4D0D90(1)
			}
			result = 0
		case 0x2841:
			sub_409EF0(2)
			result = 0
		case 0x2842:
			sub_409EF0(8192)
			if sub_409F40(8192) != 0 {
				goto LABEL_24
			}
			sub_4D7EA0()
			result = 0
		case 0x284C:
			v13 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1309812)))).ChildByID(0x284D)))
			(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v13)))))).Show()
			nox_xxx_wndShowModalMB((*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v13)))))))
			nox_xxx_wndSetCaptureMain((*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v13)))))))
			result = 0
		case 0x284F:
			sub_4BDFD0()
			goto LABEL_24
		default:
			goto LABEL_24
		}
	} else if a2 == 0x4009 {
		nox_xxx_rateUpdate_40A6D0(4 - a4)
		result = 0
	} else if a2 != 16400 || (*nox_window)(unsafe.Pointer(a3)).ID() != 0x284D || (func() bool {
		v4 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1309812)))).ChildByID(0x284C)))
		v5 = (*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a3)))))).Func94(asWindowEvent(0x4014, 0, 0))
		v6 = v5
		return v5 < 0
	}()) || v5 >= int32(*(*int16)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(a3), unsafe.Sizeof(int32(0))*8)) + 44)))) {
	LABEL_24:
		result = 0
	} else {
		v7 = (*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a3)))))).Func94(asWindowEvent(0x4016, a4, 0))
		(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v4)))))).Func94(asWindowEvent(0x4001, v7, -1))
		nox_server_connectionType_3596 = uint32(v6 + 1)
		v8 = sub_40A710(v6 + 1)
		nox_xxx_rateUpdate_40A6D0(v8)
		v9 = 4 - nox_xxx_rateGet_40A6C0()
		v10 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1309812)))).ChildByID(0x2848)))
		(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v10)))))).Func94(asWindowEvent(0x400A, v9, 0))
		(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a3)))))).Hide()
		nox_xxx_wndClearCaptureMain((*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a3)))))))
		result = 0
	}
	return result
}
func sub_4AD820() int32 {
	(*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1309812)))).Destroy()
	dword_5d4594_1309812 = 0
	return sub_4BE610()
}
func sub_4AD9B0(a1 int32) int32 {
	if (*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1309820))))).Flags().IsHidden() != 0 {
		return 0
	}
	sub_413A00(0)
	sub_44D8F0()
	if a1 != 0 {
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1309820))))).Hide()
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1309824))))).Hide()
	} else {
		nox_common_writecfgfile(CString("nox.cfg"))
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1309820))))).Hide()
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1309824))))).Hide()
		sub_445C40()
	}
	return 1
}
func sub_4ADA40() int32 {
	var (
		v0 int32
		v1 *uint32
		v3 int32
	)
	sub_413A00(1)
	nox_xxx_wndShowModalMB((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1309820))))))
	nox_xxx_wndShowModalMB((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1309824))))))
	nox_client_advVideoOptsLoad_4CB330()
	v0 = nox_video_getCutSize_4766D0()
	if v0 > 69 {
		if v0 <= 79 {
			v1 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1309820)))).ChildByID(312)))
			goto LABEL_9
		}
		if v0 <= 89 {
			v1 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1309820)))).ChildByID(313)))
			goto LABEL_9
		}
		v3 = 314
	} else {
		v3 = 311
	}
	v1 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1309820)))).ChildByID(v3)))
LABEL_9:
	if v1 != nil {
		(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v1)))))).Func94(asWindowEvent(0x4008, 0, 0))
	}
	return nox_draw_setTabWidth_43FE20(15)
}
func nox_game_initOptionsInGame_4ADAD0() int32 {
	var (
		v1  *uint32
		v2  *byte
		v3  int32
		v4  int32
		v5  uint32
		v6  *uint32
		v7  *byte
		v8  int32
		v9  int32
		v10 uint32
		v11 *uint32
		v12 *byte
		v13 int32
		v14 int32
		v15 uint32
		v16 *uint32
	)
	_ = v16
	var v17 *uint32
	var i int32
	var v19 *uint32
	var v20 int32
	var v21 *int32
	var v22 *uint32
	var v23 int32
	var v24 *uint32
	var v25 *byte
	var v26 *byte
	var v27 *byte
	var v28 *byte
	var v29 *byte
	var v30 *byte
	var v31 int32
	_ = v31
	dword_5d4594_1309820 = uint32(uintptr(unsafe.Pointer(newWindowFromFile("Options.wnd", unsafe.Pointer(cgoFuncAddr(nox_xxx_windowOptionsProc_4ADF30))))))
	if dword_5d4594_1309820 == 0 {
		return 0
	}
	if nox_client_advVideoOpts_New_4CB590((*nox_window)(unsafe.Pointer(uintptr(dword_5d4594_1309820)))) == 0 {
		return 0
	}
	v1 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1309820)))).ChildByID(351)))
	*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*100)) + 8))) = 24
	*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*100)) + 12))) = 20
	v28 = (*byte)(unsafe.Pointer(nox_xxx_gLoadImg("OptionsVolumeSliderLit")))
	v25 = (*byte)(unsafe.Pointer(nox_xxx_gLoadImg("OptionsVolumeSliderLit")))
	v2 = (*byte)(unsafe.Pointer(nox_xxx_gLoadImg("OptionsVolumeSlider")))
	sub_4B5700(int32(uintptr(unsafe.Pointer(v1))), 0, 0, int32(uintptr(unsafe.Pointer(v2))), int32(uintptr(unsafe.Pointer(v25))), int32(uintptr(unsafe.Pointer(v28))))
	(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v1)))))).Func94(asWindowEvent(0x400B, 0, 0x4000))
	(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v1)))))).Func94(asWindowEvent(0x400A, int32(*(*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(dword_587000_127004)) + 4)))>>16), 0))
	dword_5d4594_1309828 = uint32(uintptr(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1309820)))).ChildByID(361))))
	v3 = sub_453070()
	v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1309828 + 36))))
	if v3 == 1 {
		v5 = uint32(v4 | 4)
	} else {
		v5 = uint32(v4) & 0xFFFFFFFB
	}
	*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1309828 + 36))) = v5
	v6 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1309820)))).ChildByID(352)))
	*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*100)) + 8))) = 24
	*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*100)) + 12))) = 20
	v29 = (*byte)(unsafe.Pointer(nox_xxx_gLoadImg("OptionsVolumeSliderLit")))
	v26 = (*byte)(unsafe.Pointer(nox_xxx_gLoadImg("OptionsVolumeSliderLit")))
	v7 = (*byte)(unsafe.Pointer(nox_xxx_gLoadImg("OptionsVolumeSlider")))
	sub_4B5700(int32(uintptr(unsafe.Pointer(v6))), 0, 0, int32(uintptr(unsafe.Pointer(v7))), int32(uintptr(unsafe.Pointer(v26))), int32(uintptr(unsafe.Pointer(v29))))
	(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v6)))))).Func94(asWindowEvent(0x400B, 0, 0x4000))
	(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v6)))))).Func94(asWindowEvent(0x400A, int32(*(*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(dword_587000_122852)) + 4)))>>16), 0))
	dword_5d4594_1309836 = uint32(uintptr(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1309820)))).ChildByID(362))))
	v8 = sub_44D990()
	v9 = int32(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1309836 + 36))))
	if v8 == 1 {
		v10 = uint32(v9 | 4)
	} else {
		v10 = uint32(v9) & 0xFFFFFFFB
	}
	*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1309836 + 36))) = v10
	v11 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1309820)))).ChildByID(353)))
	*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v11), unsafe.Sizeof(uint32(0))*100)) + 8))) = 24
	*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v11), unsafe.Sizeof(uint32(0))*100)) + 12))) = 20
	v30 = (*byte)(unsafe.Pointer(nox_xxx_gLoadImg("OptionsVolumeSliderLit")))
	v27 = (*byte)(unsafe.Pointer(nox_xxx_gLoadImg("OptionsVolumeSliderLit")))
	v12 = (*byte)(unsafe.Pointer(nox_xxx_gLoadImg("OptionsVolumeSlider")))
	sub_4B5700(int32(uintptr(unsafe.Pointer(v11))), 0, 0, int32(uintptr(unsafe.Pointer(v12))), int32(uintptr(unsafe.Pointer(v27))), int32(uintptr(unsafe.Pointer(v30))))
	(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v11)))))).Func94(asWindowEvent(0x400B, 0, 0x4000))
	(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v11)))))).Func94(asWindowEvent(0x400A, int32(*(*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(dword_587000_93164)) + 4)))>>16), 0))
	dword_5d4594_1309832 = uint32(uintptr(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1309820)))).ChildByID(363))))
	v13 = sub_43DC30()
	v14 = int32(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1309832 + 36))))
	if v13 == 1 {
		v15 = uint32(v14 | 4)
	} else {
		v15 = uint32(v14) & 0xFFFFFFFB
	}
	*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1309832 + 36))) = v15
	nox_video_setMenuOptions((*nox_window)(unsafe.Pointer(uintptr(dword_5d4594_1309820))))
	if nox_video_getFullScreen() != 0 {
		v17 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1309820)))).ChildByID(333)))
	} else {
		v17 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1309820)))).ChildByID(334)))
	}
	(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v17)))))).Func94(asWindowEvent(0x4008, 1, 0))
	for i = 320; i <= 332; i++ {
		v19 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1309820)))).ChildByID(i)))
		if v19 != nil && (*(*uint32)(unsafe.Add(unsafe.Pointer(v19), unsafe.Sizeof(uint32(0))*9))&4) == 0 {
			nox_xxx_wndClearFlag_46AD80(int32(uintptr(unsafe.Pointer(v19))), 8)
		}
	}
	dword_5d4594_1309824 = uint32(uintptr(unsafe.Pointer(nox_window_new(nil, 32, 0, 0, 1, 1, nil))))
	v20 = int32(*memmap.PtrUint32(0x587000, 174072))
	if *memmap.PtrInt32(0x587000, 174072) != -1 {
		v21 = memmap.PtrInt32(0x587000, 174080)
		for {
			v22 = (*uint32)(unsafe.Pointer(nox_window_new((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1309824))))), 0, v20, *((*int32)(unsafe.Add(unsafe.Pointer(v21), -int(unsafe.Sizeof(int32(0))*1)))), *v21, *(*int32)(unsafe.Add(unsafe.Pointer(v21), unsafe.Sizeof(int32(0))*1)), nil)))
			int32(uintptr(unsafe.Pointer(v22))).setDraw(func(arg1 int32, arg2 int32) int32 {
				return sub_4ADEF0((*uint32)(unsafe.Pointer(uintptr(arg1))), arg2)
			})
			v20 = *(*int32)(unsafe.Add(unsafe.Pointer(v21), unsafe.Sizeof(int32(0))*2))
			v21 = (*int32)(unsafe.Add(unsafe.Pointer(v21), unsafe.Sizeof(int32(0))*4))
			if v20 == -1 {
				break
			}
		}
	}
	v23 = int32(uint32(nox_win_width) - *(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1309820 + 8))))
	(*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1309820)))).SetPos(image.Pt(v23/2, 0))
	(*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1309824)))).SetPos(image.Pt(v23/2, 0))
	v24 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1309820)))).ChildByID(371)))
	(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v24)))))).Show()
	(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1309820))))).Hide()
	(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1309824))))).Hide()
	return 1
}
func sub_4ADEF0(a1 *uint32, a2 int32) int32 {
	var (
		xLeft int32
		yTop  int32
	)
	nox_client_wndGetPosition_46AA60((*nox_window)(unsafe.Pointer(a1)), (*uint32)(unsafe.Pointer(&xLeft)), (*uint32)(unsafe.Pointer(&yTop)))
	nox_client_drawRectFilledAlpha_49CF10(xLeft, yTop, int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*2))), int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*3))))
	return 1
}
func nox_xxx_windowOptionsProc_4ADF30(a1 int32, a2 int32, a3 *int32, a4 int32) int32 {
	var (
		result int32 = 0
		v5     int32
		v6     int32
		v7     func(uint32, int32, int32, int32)
		v8     func(uint32, int32, int32, int32)
		v9     func(uint32, int32, int32, int32)
		v10    unsafe.Pointer
		v11    func(unsafe.Pointer, int32, int32, int32)
		v12    func(uint32, int32, int32, int32)
		v13    int32
		v14    func(int32, int32, int32, int32)
		v15    int32
		v16    func(uint32, int32, int32, int32)
		v17    func(uint32, int32, int32, int32)
	)
	switch a2 {
	case 0x4005:
		clientPlaySoundSpecial(920, 100)
		result = 1
	case 0x4007:
		switch (*nox_window)(unsafe.Pointer(a3)).ID() {
		case 311:
			nox_draw_setCutSize_476700(65, 0)
		case 312:
			nox_draw_setCutSize_476700(75, 0)
		case 313:
			nox_draw_setCutSize_476700(85, 0)
		case 314:
			nox_draw_setCutSize_476700(100, 0)
		case 341:
			sub_4AD9B0(0)
			sub_445C40()
			sub_4C4260()
		case 361:
			if sub_453070() == 1 {
				sub_453050()
			} else {
				nox_xxx____setargv_9_453060()
			}
		case 362:
			if sub_44D990() == 1 {
				sub_44D960()
			} else {
				sub_44D970()
			}
		case 363:
			if sub_43DC30() == 1 {
				sub_43DC00()
			} else {
				sub_43DC10()
				sub_486320(*(**uint32)(unsafe.Pointer(&dword_587000_93164)), int32(*(*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(dword_587000_93164)) + 4)))>>16))
			}
		case 371:
			sub_4AD9B0(0)
		default:
		}
		clientPlaySoundSpecial(921, 100)
		result = 1
	case 0x4009:
		v5 = (*nox_window)(unsafe.Pointer(a3)).ID()
		v6 = int32(uintptr(unsafe.Pointer(nox_client_getWin1064916_46C720())))
		switch v5 {
		case 351:
			sub_486320(*(**uint32)(unsafe.Pointer(&dword_587000_127004)), a4)
			if v6 != 0 && *(**int32)(unsafe.Pointer(uintptr(v6 + 396))) == a3 {
				goto LABEL_64
			}
			if a4 != 0 {
				if sub_453070() != 0 {
					goto LABEL_63
				}
				v13 = int32(dword_5d4594_1309828)
				v14 = cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1309828 + 372))), (*func(int32, int32, int32, int32))(nil))
				if v14 == nil {
					goto LABEL_63
				}
				goto LABEL_62
			}
			if sub_453070() != 1 {
				goto LABEL_64
			}
			v12 = cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1309828 + 372))), (*func(uint32, int32, int32, int32))(nil))
			if v12 == nil {
				goto LABEL_64
			}
			v12(dword_5d4594_1309828, 21, 28, 2)
			result = 0
		case 352:
			sub_486320(*(**uint32)(unsafe.Pointer(&dword_587000_122852)), a4)
			if a4 != 0 {
				if sub_44D990() != 0 {
					goto LABEL_37
				}
				v10 = *(*unsafe.Pointer)(unsafe.Pointer(&dword_5d4594_1309836))
				v11 = cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1309836 + 372))), (*func(unsafe.Pointer, int32, int32, int32))(nil))
				if v11 == nil {
					goto LABEL_37
				}
				goto LABEL_36
			}
			if sub_44D990() != 1 {
				goto LABEL_64
			}
			v9 = cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1309836 + 372))), (*func(uint32, int32, int32, int32))(nil))
			if v9 == nil {
				goto LABEL_64
			}
			v9(dword_5d4594_1309836, 21, 28, 2)
			result = 0
		case 353:
			sub_486320(*(**uint32)(unsafe.Pointer(&dword_587000_93164)), a4)
			if a4 != 0 {
				if sub_43DC30() != 0 {
					goto LABEL_64
				}
				v8 = cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1309832 + 372))), (*func(uint32, int32, int32, int32))(nil))
				if v8 == nil {
					goto LABEL_64
				}
				v8(dword_5d4594_1309832, 21, 28, 2)
				result = 0
			} else {
				if sub_43DC30() != 1 {
					goto LABEL_64
				}
				v7 = cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1309832 + 372))), (*func(uint32, int32, int32, int32))(nil))
				if v7 == nil {
					goto LABEL_64
				}
				v7(dword_5d4594_1309832, 21, 28, 2)
				result = 0
			}
		case 316:
			nox_video_setGamma(float32(float64(a4)/50.0 + 0.5))
		case 318:
			nox_input_setSensitivity(math32.Pow(10.0, float32(float64(a4)/50.0-1.0)))
		default:
			goto LABEL_64
		}
	case 0x400C:
		v15 = (*nox_window)(unsafe.Pointer(a3)).ID() - 351
		if v15 != 0 {
			if v15 != 1 {
				goto LABEL_64
			}
			sub_486320(*(**uint32)(unsafe.Pointer(&dword_587000_122852)), a4)
			if a4 != 0 {
				if sub_44D990() == 0 {
					v10 = *(*unsafe.Pointer)(unsafe.Pointer(&dword_5d4594_1309836))
					v11 = cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1309836 + 372))), (*func(unsafe.Pointer, int32, int32, int32))(nil))
					if v11 != nil {
					LABEL_36:
						v11(v10, 21, 28, 2)
					}
				}
			LABEL_37:
				sub_4AA650()
				result = 0
			} else {
				if sub_44D990() != 1 {
					goto LABEL_64
				}
				v16 = cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1309836 + 372))), (*func(uint32, int32, int32, int32))(nil))
				if v16 == nil {
					goto LABEL_64
				}
				v16(dword_5d4594_1309836, 21, 28, 2)
				result = 0
			}
		} else {
			sub_486320(*(**uint32)(unsafe.Pointer(&dword_587000_127004)), a4)
			if a4 != 0 {
				if sub_453070() == 0 {
					v13 = int32(dword_5d4594_1309828)
					v14 = cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1309828 + 372))), (*func(int32, int32, int32, int32))(nil))
					if v14 != nil {
					LABEL_62:
						v14(v13, 21, 28, 2)
					}
				}
			LABEL_63:
				clientPlaySoundSpecial(768, 100)
				goto LABEL_64
			}
			if sub_453070() != 1 {
				goto LABEL_64
			}
			v17 = cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1309828 + 372))), (*func(uint32, int32, int32, int32))(nil))
			if v17 == nil {
				goto LABEL_64
			}
			v17(dword_5d4594_1309828, 21, 28, 2)
			result = 0
		}
	default:
	LABEL_64:
		result = 0
	}
	return result
}
func sub_4AE3B0() int32 {
	var result int32
	result = (*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1309820)))).Destroy()
	dword_5d4594_1309820 = 0
	return result
}
func sub_4AE3D0() int32 {
	var result int32
	if (*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1309820))))).Flags().IsHidden() != 0 {
		result = sub_4C4280()
	} else {
		result = 1
	}
	return result
}
func sub_4AE420() int32 {
	var result int32
	if *memmap.PtrUint32(0x973A20, 536) >= 3 {
		nox_client_drawAddPoint_49F500(int32(*(*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(dword_5d4594_3798696)) + *memmap.PtrUint32(0x973A20, 536)*8 - 24)))), int32(*(*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(dword_5d4594_3798696)) + *memmap.PtrUint32(0x973A20, 536)*8 - 20)))))
	}
	result = (*(*funcuint32(int32))(unsafe.Pointer(&dword_5d4594_3798716)))(1)
	if result != 0 {
		result = (*(*funcuint32(int32))(unsafe.Pointer(&dword_5d4594_3798716)))(1)
		if result != 0 {
			result = (*(*funcuint32(int32))(unsafe.Pointer(&dword_5d4594_3798716)))(0)
		}
	}
	return result
}
func sub_4AE470() int32 {
	var (
		result int32
		v1     int32
		v2     *uint16
		v3     *uint16
	)
	sub_4AEBD0()
	if *memmap.PtrUint32(0x973A20, 536) >= 3 {
		nox_client_drawAddPoint_49F500(int32(*(*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(dword_5d4594_3798696)) + *memmap.PtrUint32(0x973A20, 536)*8 - 24)))), int32(*(*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(dword_5d4594_3798696)) + *memmap.PtrUint32(0x973A20, 536)*8 - 20)))))
	}
	result = sub_4AEC20(1, 0)
	if result != 0 {
		result = sub_4AEC20(1, 0)
		if result != 0 {
			result = sub_4AEC20(0, 1)
			if result != 0 {
				v1 = int32(dword_5d4594_3798636)
				for result = int32(dword_5d4594_3798640); v1 <= *(*int32)(unsafe.Pointer(&dword_5d4594_3798640)); dword_5d4594_3798636 = uint32(v1) {
					v2 = *(**uint16)(unsafe.Pointer(uintptr(uint32(uintptr(dword_5d4594_3798632)) + uint32(v1*4))))
					v3 = (*uint16)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), unsafe.Sizeof(uint32(0))*1))))))
					if *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v3))), unsafe.Sizeof(uint32(0))*1))) != 0 {
						(cgoAsFunc(*(*uint32)(unsafe.Pointer(&dword_5d4594_3798720)), (*func(uint32, uint32, uint32))(nil)))(uint32(*v2), uint32(v1), uint32(*v3))
						v1 = int32(dword_5d4594_3798636)
					}
					result = int32(dword_5d4594_3798640)
					v1++
				}
			}
		}
	}
	return result
}
func sub_4AE580(a1 int32) int32 {
	if *memmap.PtrUint32(0x973A20, 536) >= uint32(a1) {
		nox_client_drawAddPoint_49F500(int32(*(*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(dword_5d4594_3798696)) + (*memmap.PtrUint32(0x973A20, 536)-uint32(a1))*8)))), int32(*(*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(dword_5d4594_3798696)) + (*memmap.PtrUint32(0x973A20, 536)-uint32(a1))*8 + 4)))))
	}
	return sub_49E510(a1)
}
func sub_4AE5C0(a1 int32) int32 {
	var (
		v1     int32
		result int32
		v3     int32
		v4     int32
		v5     *byte
		v6     *byte
		v7     uint32
		v8     uint16
		v9     uint16
		v10    [16]byte
	)
	sub_4AE5E0(a1, 1)
	v1 = int32(dword_5d4594_3798636)
	for result = int32(dword_5d4594_3798640); v1 <= *(*int32)(unsafe.Pointer(&dword_5d4594_3798640)); dword_5d4594_3798636 = uint32(v1) {
		v3 = 0
		v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(dword_5d4594_3798632)) + uint32(v1*4)))))
		if v4 != 0 {
			v5 = &v10[0]
			for {
				v3++
				*(*uint16)(unsafe.Pointer(v5)) = *(*uint16)(unsafe.Pointer(uintptr(v4)))
				v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(v4 + 4))))
				v5 = (*byte)(unsafe.Add(unsafe.Pointer(v5), 2))
				if v4 == 0 {
					break
				}
			}
			if v3 >= 2 {
				sub_48C5B0((*int16)(unsafe.Pointer(&v10[0])), v3)
				v6 = &v10[0]
				v7 = uint32(v3) >> 1
				for {
					v8 = *(*uint16)(unsafe.Pointer(v6))
					v9 = *((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v6))), unsafe.Sizeof(uint16(0))*1)))
					v6 = (*byte)(unsafe.Add(unsafe.Pointer(v6), 4))
					(cgoAsFunc(*(*uint32)(unsafe.Pointer(&dword_5d4594_3798720)), (*func(uint32, uint32, uint32))(nil)))(uint32(v8), dword_5d4594_3798636, uint32(v9))
					v7--
					if v7 == 0 {
						break
					}
				}
				v1 = int32(dword_5d4594_3798636)
			}
		}
		result = int32(dword_5d4594_3798640)
		v1++
	}
	return result
}
func sub_4AE5E0(a1 int32, a2 int32) int32 {
	var v2 int32
	v2 = a1
	if *memmap.PtrUint32(0x973A20, 536) >= uint32(a1) {
		nox_client_drawAddPoint_49F500(int32(*(*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(dword_5d4594_3798696)) + (*memmap.PtrUint32(0x973A20, 536)-uint32(a1))*8)))), int32(*(*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(dword_5d4594_3798696)) + (*memmap.PtrUint32(0x973A20, 536)-uint32(a1))*8 + 4)))))
	}
	if a2 != 0 {
		sub_4AEBD0()
	}
	if a1 <= 1 {
		return bool2int(sub_4AEC20(0, 1) != 0)
	}
	for sub_4AEC20(1, 0) != 0 {
		if func() int32 {
			p := &v2
			*p--
			return *p
		}() <= 1 {
			return bool2int(sub_4AEC20(0, 1) != 0)
		}
	}
	return 0
}
func sub_4AE6F0(a1 int32, a2 int32, a3 int32, a4 int32, a5 int32) int32 {
	var (
		result int32
		v6     int32
		v7     int32
		v8     int32
		v9     int32
		v10    func(int32, int32, int32)
		v11    int32
		v12    int32
		v13    int32
		v14    *uint16
		v15    *uint16
		v16    uint16
		v17    *uint16
		v18    *uint16
		v19    uint16
		v20    *uint16
		v21    *uint16
		v22    uint16
		v23    int32
		v24    *uint16
		v25    *uint16
		v26    uint16
		v27    bool
		v28    int32
		v29    int32
		v30    int32
		v31    int32
		v32    int32
		v33    int32
		v34    int32
		v35    int32
		v36    int32
		v37    int32
		v38    int32
		v39    int32
		v40    int32
		v41    int32
	)
	v31 = 1 - a3
	v35 = 0
	v37 = a3
	v34 = 3
	v32 = 5 - a3*2
	nox_client_drawSetColor_434460(a5)
	result = 1
	v6 = a4
	if a4 < 0 {
		v6 = 0
	LABEL_4:
		v7 = (a3 * *memmap.PtrInt32(6112660, uint32((256-v6)*4)+0x13FD90)) >> 16
		v30 = (a3 * *memmap.PtrInt32(6112660, uint32((256-v6)*4)+0x13FD90)) >> 16
		v8 = (a3 * *memmap.PtrInt32(6112660, uint32((256-v6)*4)+0x13FC90)) >> 16
		sub_4AEBD0()
		v9 = a2
		v41 = a2 - a3
		nox_client_drawAddPoint_49F500(a1, a2-a3)
		nox_client_drawAddPoint_49F500(a1, a2)
		sub_4AEC20(0, 0)
		nox_client_drawAddPoint_49F500(a1, a2)
		nox_client_drawAddPoint_49F500(a1+v7, a2+v8)
		sub_4AEC20(0, 1)
		v10 = cgoAsFunc(*(*uint32)(unsafe.Pointer(&dword_5d4594_3798708)), (*func(int32, int32, int32))(nil))
		if nox_draw_curDrawData_3799572.field_13 == 0 {
			v10 = cgoAsFunc(*(*uint32)(unsafe.Pointer(&dword_5d4594_3798720)), (*func(int32, int32, int32))(nil))
		}
		v11 = a2 + a3
		v12 = a2 - a3
		v38 = 0
		v39 = v9
		v40 = v9
		v36 = v9 * 4
		v33 = v41 * 4
		v13 = a1 - v9
		for {
			if v38 != 0 {
				if v12 < *(*int32)(unsafe.Pointer(&dword_5d4594_3798636)) || v12 > *(*int32)(unsafe.Pointer(&dword_5d4594_3798640)) {
					if v30 <= 0 {
						v10(v13+v40, v12, v13+v39)
					}
				} else {
					v14 = *(**uint16)(unsafe.Pointer(uintptr(uint32(v33) + uint32(uintptr(dword_5d4594_3798632)))))
					v15 = (*uint16)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v14))), unsafe.Sizeof(uint32(0))*1))))))
					if v15 != nil {
						if int32(*v14) > int32(*v15) {
							v16 = *v14
							*v14 = *v15
							*v15 = v16
						}
						if v30 > 0 {
							v10(int32(*v14), v41, int32(**((**uint16)(unsafe.Add(unsafe.Pointer((**uint16)(unsafe.Pointer(v14))), unsafe.Sizeof((*uint16)(nil))*1)))))
						} else {
							v10(v13+v40, v41, int32(*v14))
							v10(int32(**((**uint16)(unsafe.Add(unsafe.Pointer((**uint16)(unsafe.Pointer(v14))), unsafe.Sizeof((*uint16)(nil))*1)))), v41, v13+v39)
						}
					} else {
						v10(int32(*v14), v12, v13+v39)
					}
				}
			}
			if v35 != 0 {
				if v40 < *(*int32)(unsafe.Pointer(&dword_5d4594_3798636)) || v40 >= *(*int32)(unsafe.Pointer(&dword_5d4594_3798640)) {
					if v30 <= 0 {
						v10(v13+v41, v40, v13+v11)
					}
				} else {
					v17 = *(**uint16)(unsafe.Pointer(uintptr(uint32(v36) + uint32(uintptr(dword_5d4594_3798632)))))
					v18 = (*uint16)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v17))), unsafe.Sizeof(uint32(0))*1))))))
					if v18 != nil {
						if int32(*v17) > int32(*v18) {
							v19 = *v17
							*v17 = *v18
							*v18 = v19
						}
						if v30 > 0 {
							v10(int32(*v17), v40, int32(**((**uint16)(unsafe.Add(unsafe.Pointer((**uint16)(unsafe.Pointer(v17))), unsafe.Sizeof((*uint16)(nil))*1)))))
						} else {
							v10(v13+v41, v40, int32(*v17))
							v10(int32(**((**uint16)(unsafe.Add(unsafe.Pointer((**uint16)(unsafe.Pointer(v17))), unsafe.Sizeof((*uint16)(nil))*1)))), v40, v13+v11)
						}
					} else {
						v10(int32(*v17), v40, v13+v11)
					}
				}
			}
			if v39 >= *(*int32)(unsafe.Pointer(&dword_5d4594_3798636)) && v39 < *(*int32)(unsafe.Pointer(&dword_5d4594_3798640)) {
				break
			}
			if v30 <= 0 {
				v29 = v13 + v11
				v28 = v39
				v23 = v13 + v41
			LABEL_42:
				v10(v23, v28, v29)
			}
		LABEL_43:
			if v38 != 0 {
				if v11 < *(*int32)(unsafe.Pointer(&dword_5d4594_3798636)) || v11 > *(*int32)(unsafe.Pointer(&dword_5d4594_3798640)) {
					if v30 <= 0 {
						v10(v13+v40, v11, v13+v39)
					}
				} else {
					v24 = *(**uint16)(unsafe.Pointer(uintptr(uint32(uintptr(dword_5d4594_3798632)) + uint32(v11*4))))
					v25 = (*uint16)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v24))), unsafe.Sizeof(uint32(0))*1))))))
					if v25 != nil {
						if int32(*v24) > int32(*v25) {
							v26 = *v24
							*v24 = *v25
							*v25 = v26
						}
						if v30 > 0 {
							v10(int32(*v24), v11, int32(**((**uint16)(unsafe.Add(unsafe.Pointer((**uint16)(unsafe.Pointer(v24))), unsafe.Sizeof((*uint16)(nil))*1)))))
						} else {
							v10(v13+v40, v11, int32(*v24))
							v10(int32(**((**uint16)(unsafe.Add(unsafe.Pointer((**uint16)(unsafe.Pointer(v24))), unsafe.Sizeof((*uint16)(nil))*1)))), v11, v13+v39)
						}
					} else {
						v10(int32(*v24), v11, v13+v39)
					}
				}
			}
			if v31 >= 0 {
				v31 += v32
				v33 += 4
				v12 = v41 + 1
				v32 += 4
				v38 = 1
				v37--
				v41++
				v11--
			} else {
				v31 += v34
				v12 = v41
				v32 += 2
				v38 = 0
			}
			v34 += 2
			result = v35 + 1
			v39++
			v36 -= 4
			v27 = v37 < func() int32 {
				p := &v35
				*p++
				return *p
			}()
			v40--
			if v27 {
				return result
			}
		}
		v20 = *(**uint16)(unsafe.Pointer(uintptr(uint32(uintptr(dword_5d4594_3798632)) + uint32(v39*4))))
		v21 = (*uint16)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v20))), unsafe.Sizeof(uint32(0))*1))))))
		if v21 != nil {
			if int32(*v20) > int32(*v21) {
				v22 = *v20
				*v20 = *v21
				*v21 = v22
			}
			if v30 > 0 {
				v10(int32(*v20), v39, int32(**((**uint16)(unsafe.Add(unsafe.Pointer((**uint16)(unsafe.Pointer(v20))), unsafe.Sizeof((*uint16)(nil))*1)))))
			} else {
				v10(v13+v41, v39, int32(*v20))
				v10(int32(**((**uint16)(unsafe.Add(unsafe.Pointer((**uint16)(unsafe.Pointer(v20))), unsafe.Sizeof((*uint16)(nil))*1)))), v39, v13+v11)
			}
			goto LABEL_43
		}
		v29 = v13 + v11
		v28 = v39
		v23 = int32(*v20)
		goto LABEL_42
	}
	if a4 < 256 {
		goto LABEL_4
	}
	return result
}
func sub_4AEC20(a1 int32, a2 int32) int32 {
	var (
		result int32
		v3     int32
		v4     int32
		v5     int32
		v6     int32
		v7     int32
		v8     int32
		v9     int32
		v10    int32
		v11    int32
		v12    int32
		v13    int32
	)
	result = sub_49F5B0((*uint32)(unsafe.Pointer(&v12)), (*uint32)(unsafe.Pointer(&v13)), 0)
	if result != 0 {
		result = sub_49F5B0((*uint32)(unsafe.Pointer(&v10)), (*uint32)(unsafe.Pointer(&v11)), a1)
		if result != 0 {
			if sub_49FC20(&v10, &v11, &v12, &v13) != 0 {
				v3 = v12 - v10
				v10 <<= 16
				v4 = v11
				v5 = v3 << 16
				if v11 > v13 {
					v6 = -1
					v7 = v11 - v13
					v5 /= v11 - v13
					if v13 < *(*int32)(unsafe.Pointer(&dword_5d4594_3798636)) {
						dword_5d4594_3798636 = uint32(v13)
					}
					if v11 > *(*int32)(unsafe.Pointer(&dword_5d4594_3798640)) {
						dword_5d4594_3798640 = uint32(v11)
					}
				} else {
					v6 = 1
					v7 = v13 - v11
					if v13 != v11 {
						v5 /= v7
					}
					if v11 < *(*int32)(unsafe.Pointer(&dword_5d4594_3798636)) {
						dword_5d4594_3798636 = uint32(v11)
					}
					if v13 > *(*int32)(unsafe.Pointer(&dword_5d4594_3798640)) {
						dword_5d4594_3798640 = uint32(v13)
					}
				}
				if v7 > 0 {
					for {
						v8 = int32(*(*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(dword_5d4594_3798632)) + uint32(v4*4)))))
						*(*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(dword_5d4594_3798632)) + uint32(v4*4)))) = uint32(uintptr(unsafe.Pointer(dword_5d4594_3798648)))
						**(**uint16)(unsafe.Pointer(&dword_5d4594_3798648)) = *(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v10))), unsafe.Sizeof(uint16(0))*1))
						*(*uint32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(dword_5d4594_3798648), 4)))) = uint32(v8)
						dword_5d4594_3798648 = (*byte)(unsafe.Add(unsafe.Pointer(dword_5d4594_3798648), 8))
						v4 = v6 + v11
						v7--
						v10 += v5
						v11 += v6
						if v7 == 0 {
							break
						}
					}
				}
				if a2 != 0 {
					v9 = int32(*(*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(dword_5d4594_3798632)) + uint32(v4*4)))))
					*(*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(dword_5d4594_3798632)) + uint32(v4*4)))) = uint32(uintptr(unsafe.Pointer(dword_5d4594_3798648)))
					**(**uint16)(unsafe.Pointer(&dword_5d4594_3798648)) = uint16(int16(v12))
					*(*uint32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(dword_5d4594_3798648), 4)))) = uint32(v9)
					dword_5d4594_3798648 = (*byte)(unsafe.Add(unsafe.Pointer(dword_5d4594_3798648), 8))
				}
				result = 1
			} else {
				result = 1
			}
		}
	}
	return result
}
func sub_4AEDA0(a1 *int32, a2 *int32, a3 int32, a4 int32) int32 {
	var result int32
	result = a3
	if a3 >= 0 {
		if a3 > 256 {
			result = 256
		}
	} else {
		result = 0
	}
	if a1 != nil {
		*a1 = (a4 * *memmap.PtrInt32(6112660, uint32(result*4)+0x13FD90)) >> 16
	}
	if a2 != nil {
		result = (a4 * *memmap.PtrInt32(6112660, uint32(result*4)+0x13FC90)) >> 16
		*a2 = result
	}
	return result
}
func sub_4AEE30() int64 {
	var (
		v0     int32
		v1     *uint8
		result int64
	)
	v0 = 0
	v1 = (*uint8)(memmap.PtrOff(6112660, 1309840))
	for {
		result = int64(math.Sin(float64(v0+192)**mem_getDoublePtr(0x581450, 9768)**mem_getDoublePtr(0x581450, 9760)) * *(*float64)(unsafe.Pointer(&qword_581450_9552)))
		*(*uint32)(unsafe.Pointer(v1)) = uint32(int32(result))
		v1 = (*uint8)(unsafe.Add(unsafe.Pointer(v1), 4))
		v0++
		if int32(uintptr(unsafe.Pointer(v1))) >= int32(uintptr(memmap.PtrOff(6112660, 1311120))) {
			break
		}
	}
	return result
}
func nox_xxx_ParticleFxT0_4AEEA0(a1 **uint8) *uint8 {
	var result *uint8
	sub_4AF8C0((*uint32)(unsafe.Pointer(a1)))
	result = *(**uint8)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof((*uint8)(nil))*1))
	if result != nil {
		sub_4AF890(a1)
		result = (*uint8)(unsafe.Pointer(nox_xxx_partfxSwitch_4AF690((*uint32)(unsafe.Pointer(a1)), func(arg1 *uint32, arg2 *uint32, arg3 int32) {
			sub_4AEED0((*int32)(unsafe.Pointer(arg1)), int32(uintptr(unsafe.Pointer(arg2))), arg3)
		})))
	}
	return result
}
func sub_4AEED0(a1 *int32, a2 int32, a3 int32) *byte {
	var (
		v3     int32
		v4     *int32
		v5     int32
		v6     int32
		result *byte
		v8     *byte
		v9     int32
	)
	nox_xxx_drawMakeRGB_433F10(math.MaxUint8, 128, 32)
	v3 = sub_48C5E0(2, 3)
	sub_434080(v3)
	v4 = (*int32)(sub_4B0680(0, math.MaxUint8))
	v5 = sub_48C5E0(0, 40)
	v6 = sub_48C5E0(0, *(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*7)))
	result = nox_xxx_drawPartFx2_4AF990(a2, a3+v5, v5, v6)
	v8 = result
	if result != nil {
		sub_4AFB00(int32(uintptr(unsafe.Pointer(result))), int32(uintptr(unsafe.Pointer(v4))))
		sub_4AFB10(int32(uintptr(unsafe.Pointer(v8))), *a1)
		sub_4AFB90(int32(uintptr(unsafe.Pointer(v8))), 50)
		*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v9))), unsafe.Sizeof(uint16(0))*0)) = uint16(sub_48C610())
		sub_4AFBD0(int32(uintptr(unsafe.Pointer(v8))), v9)
		sub_4AFC50(int32(uintptr(unsafe.Pointer(v8))), math.MinInt16)
		result = (*byte)(unsafe.Pointer(uintptr(sub_4AFC60(int32(uintptr(unsafe.Pointer(v8))), 0xB333, 1))))
	}
	return result
}
func nox_xxx_ParticleFxT1_4AEF80(a1 *int32) int32 {
	var (
		v1     *int32
		v2     int32
		v3     int32
		result int32
		v5     *byte
		v6     int32
		v7     int32
		v8     int32
		v9     int32
	)
	v1 = a1
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(*a1 + 12))))
	v8 = int32(*(*uint32)(unsafe.Pointer(uintptr(*a1 + 16))))
	nox_set_color_rgb_434430(math.MaxUint8, math.MaxUint8, math.MaxUint8)
	v3 = *(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*5))
	result = v3 - 1
	if v3 != 0 {
		v9 = *(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*5))
		for {
			v5 = nox_xxx_drawPartFx2_4AF990(v2, v8, 0, 128)
			if v5 != nil {
				*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v5))), unsafe.Sizeof(uint32(0))*17))) = uint32(sub_48C5E0(-256, 256))
				*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v5))), unsafe.Sizeof(uint32(0))*18))) = 0
				*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v5))), unsafe.Sizeof(uint32(0))*19))) = 0
				sub_4AFB10(int32(uintptr(unsafe.Pointer(v5))), *v1)
				v6 = sub_48C5E0(0, 128)
				sub_4AFB90(int32(uintptr(unsafe.Pointer(v5))), v6)
				sub_4AFB50(int32(uintptr(unsafe.Pointer(v5))), 2)
				*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v7))), unsafe.Sizeof(uint16(0))*0)) = uint16(sub_48C610())
				sub_4AFBE0(int32(uintptr(unsafe.Pointer(v5))), int32(-16384-(v7>>2)))
				sub_4AFB60(int32(uintptr(unsafe.Pointer(v5))), int32(cgoFuncAddr(sub_4AF060)))
			}
			result = func() int32 {
				p := &v9
				*p--
				return *p
			}()
			if v9 == 0 {
				break
			}
		}
	}
	return result
}
func sub_4AF060(a1 int32) int32 {
	var (
		result int32
		v2     int32
		v3     int32
		v4     int32
		v5     int32
		v6     int32
		v7     int32
		v8     int32
		v9     int32
		v10    int32
		v11    int32
		v12    int32
	)
	result = a1
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 48))))
	if v2 >= 0 {
		v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 44))))
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 40)))++
		if v3 != 0 {
			*(*uint32)(unsafe.Pointer(uintptr(a1 + 44))) = uint32(v3 - 1)
		} else {
			v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 84))))
			v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 96))))
			v6 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 72))))
			if int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 68)))) >= 0 {
				v7 = v6 - 4
			} else {
				v7 = v6 + 4
			}
			v8 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 68))))
			*(*uint32)(unsafe.Pointer(uintptr(a1 + 72))) = uint32(v7)
			v9 = v7 + v8
			*(*uint32)(unsafe.Pointer(uintptr(a1 + 68))) = uint32(v9)
			v10 = int32(uint32(v9) * *(*uint32)(unsafe.Pointer(uintptr(a1 + 76))))
			v11 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 28))))
			*(*uint32)(unsafe.Pointer(uintptr(a1 + 76))) += 256
			v12 = (v11 << 16) + (v10 >> 2)
			*(*uint32)(unsafe.Pointer(uintptr(a1 + 48))) = uint32(v2 - 1)
			if v2-1 >= 0 {
				*(*uint32)(unsafe.Pointer(uintptr(a1 + 84))) = uint32(v4 + v5)
				*(*uint32)(unsafe.Pointer(uintptr(a1 + 80))) = uint32(v12)
			} else {
				result = (cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(a1 + 132))), (*func(int32) int32)(nil)))(a1)
			}
		}
	}
	return result
}
func nox_xxx_ParticleFxT2_4AF0F0(a1 **uint8) *uint8 {
	var (
		v1     **uint8
		result *uint8
		v3     **uint8
		v4     int64
		v5     int32
		v6     *uint8
		v7     int32
		v8     int32
		v9     int32
		v10    *byte
		v11    int32
		v12    *byte
		v13    int32
		v14    *uint8
		v15    int32
		v16    int32
	)
	v1 = a1
	sub_4AF8C0((*uint32)(unsafe.Pointer(a1)))
	result = *(**uint8)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof((*uint8)(nil))*1))
	if result != nil {
		sub_4AF890(v1)
		v3 = (**uint8)(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof((*uint8)(nil))*4))))
		v14 = *(**uint8)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof((*uint8)(nil))*3))
		a1 = v3
		sub_47D5C0(int32(uintptr(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof((*uint8)(nil))*1))))), (*uint32)(unsafe.Pointer(&v14)), (*uint32)(unsafe.Pointer(&a1)), (*uint32)(unsafe.Pointer(&v15)), (*uint32)(unsafe.Pointer(&v16)))
		v4 = int64(v15)
		v5 = int32(v4)
		result = (*uint8)(unsafe.Pointer(uintptr(v4 - int64(*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v4))), unsafe.Sizeof(uint32(0))*1))))))
		v6 = *(**uint8)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof((*uint8)(nil))*5))
		v7 = int32(uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v6), -1)))))
		if v6 != nil {
			for {
				if v7&1 != 0 {
					nox_client_drawSetColor_434460(int32(nox_color_white_2523948))
				} else {
					nox_client_drawSetColor_434460(int32(nox_color_yellow_2589772))
				}
				v13 = int32(uintptr(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof((*uint8)(nil))*7)))))
				v8 = sub_48C5E0(-8, -2)
				v12 = (*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(a1))), v8))
				v9 = sub_48C5E0(int32(uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v14), v5/2))))), int32(uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v14), v15))))))
				v10 = nox_xxx_drawPartFx2_4AF990(v9, int32(uintptr(unsafe.Pointer(v12))), 0, v13)
				v11 = int32(uintptr(unsafe.Pointer(v10)))
				if v10 != nil {
					sub_4AFB50(int32(uintptr(unsafe.Pointer(v10))), 2)
					sub_4AFB10(v11, int32(uintptr(unsafe.Pointer(*v1))))
					*(*uint32)(unsafe.Pointer(uintptr(v11 + 68))) = uint32(sub_48C5E0(v5/(-2), v5/2))
					*(*uint32)(unsafe.Pointer(uintptr(v11 + 72))) = 0
					sub_4AFB60(v11, int32(cgoFuncAddr(sub_4AF200)))
				}
				result = (*uint8)(unsafe.Pointer(uintptr(func() int32 {
					p := &v7
					x := *p
					*p--
					return x
				}())))
				if result == nil {
					break
				}
			}
		}
	}
	return result
}
func sub_4AF200(a1 int32) int32 {
	var (
		result int32
		v2     int32
		v3     int32
		v4     *uint32
		v5     int32
		v6     int32
		v7     int32
		v8     int32
		v9     int32
		v10    int32
	)
	result = a1
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 48))))
	if v2 >= 0 {
		v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 44))))
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 40)))++
		if v3 != 0 {
			*(*uint32)(unsafe.Pointer(uintptr(a1 + 44))) = uint32(v3 - 1)
		} else {
			if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 56))))&4 != 0 {
				v4 = *(**uint32)(unsafe.Pointer(uintptr(a1 + 8)))
				if v4 != nil {
					v5 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*4)) - *(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*9)))
					*(*uint32)(unsafe.Pointer(uintptr(a1 + 28))) += *(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*3)) - *(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*8))
					*(*uint32)(unsafe.Pointer(uintptr(a1 + 32))) += uint32(v5)
				}
			}
			v6 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 72))))
			if int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 68)))) >= 0 {
				v7 = v6 - 1
			} else {
				v7 = v6 + 1
			}
			v8 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 68))))
			*(*uint32)(unsafe.Pointer(uintptr(a1 + 72))) = uint32(v7)
			v9 = v7 + v8
			*(*uint32)(unsafe.Pointer(uintptr(a1 + 68))) = uint32(v9)
			*(*uint32)(unsafe.Pointer(uintptr(a1 + 48))) = uint32(v2 - 1)
			if v2-1 >= 0 {
				v10 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 32))) << 16)
				*(*uint32)(unsafe.Pointer(uintptr(a1 + 80))) = (uint32(v9) + *(*uint32)(unsafe.Pointer(uintptr(a1 + 28)))) << 16
				*(*uint32)(unsafe.Pointer(uintptr(a1 + 84))) = uint32(v10)
			} else {
				result = (cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(a1 + 132))), (*func(int32) int32)(nil)))(a1)
			}
		}
	}
	return result
}
func nox_xxx_ParticleFxT3_4AF2A0(a1 **uint8) *uint8 {
	var result *uint8
	sub_4AF8C0((*uint32)(unsafe.Pointer(a1)))
	result = *(**uint8)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof((*uint8)(nil))*1))
	if result != nil {
		sub_4AF890(a1)
		result = (*uint8)(unsafe.Pointer(nox_xxx_partfxSwitch_4AF690((*uint32)(unsafe.Pointer(a1)), func(arg1 *uint32, arg2 *uint32, arg3 int32) {
			sub_4AF2D0((*int32)(unsafe.Pointer(arg1)), int32(uintptr(unsafe.Pointer(arg2))), arg3)
		})))
	}
	return result
}
func sub_4AF2D0(a1 *int32, a2 int32, a3 int32) {
	var (
		v3 int32
		v4 *int32
		v5 int32
		v7 *byte
		v8 int32
		v9 int32
	)
	if *memmap.PtrUint32(6112660, 1311132) == 0 {
		*memmap.PtrUint32(6112660, 1311124) = nox_color_rgb_4344A0(64, 200, 64)
		*memmap.PtrUint32(6112660, 1311128) = nox_color_rgb_4344A0(64, 64, math.MaxUint8)
		*memmap.PtrUint32(6112660, 1311132) = 1
	}
	sub_434040(int32(*memmap.PtrUint32(6112660, uint32(*(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*6))*4)+0x140194)))
	v3 = sub_48C5E0(2, 5)
	sub_434080(v3)
	v4 = (*int32)(sub_4B0680(0, math.MaxUint8))
	v9 = a2
	v8 = a3
	sub_48C650(0, 256, 120, (*uint32)(unsafe.Pointer(&v9)), (*uint32)(unsafe.Pointer(&v8)))
	v5 = sub_48C5E0(1, *(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*7)))
	v7 = nox_xxx_drawPartFx2_4AF990(v9, v8, 0, v5)
	if v7 != nil {
		nox_xxx_partfxLoadParticle_4AFE20((*uint32)(unsafe.Pointer(v7)), *(**byte)(memmap.PtrOff(0x973A20, 472)))
		sub_4AFB10(int32(uintptr(unsafe.Pointer(v7))), *a1)
		sub_4AFB00(int32(uintptr(unsafe.Pointer(v7))), int32(uintptr(unsafe.Pointer(v4))))
		sub_4AFC90((*uint32)(unsafe.Pointer(v7)), a2, a3, 0)
	}
}
func nox_xxx_ParticleFxT4_4AF3D0(a1 *uint32) *uint32 {
	nox_set_color_rgb_434430(math.MaxUint8, math.MaxUint8, math.MaxUint8)
	return nox_xxx_partfxSwitch_4AF690(a1, func(arg1 *uint32, arg2 *uint32, arg3 int32) {
		sub_4AF400(int32(uintptr(unsafe.Pointer(arg1))), int32(uintptr(unsafe.Pointer(arg2))), arg3)
	})
}
func sub_4AF400(a1 int32, a2 int32, a3 int32) int32 {
	var (
		v3 int32
		v4 *byte
		v5 int32
		v6 int32
	)
	v3 = sub_48C5E0(0, 100)
	v4 = nox_xxx_drawPartFx2_4AF990(a2, a3, 0, v3)
	sub_4AFB50(int32(uintptr(unsafe.Pointer(v4))), 1)
	v5 = sub_48C5E0(0, 64)
	sub_4AFB90(int32(uintptr(unsafe.Pointer(v4))), v5)
	*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v6))), unsafe.Sizeof(uint16(0))*0)) = uint16(sub_48C610())
	return sub_4AFBE0(int32(uintptr(unsafe.Pointer(v4))), v6-6553)
}
func nox_xxx_ParticleFxT5_4AF450(a1 *int32) *byte {
	var (
		v1     *int32
		result *byte
		v3     *byte
	)
	nox_xxx_drawMakeRGB_433F10(50, 50, 50)
	sub_434080(10)
	v1 = (*int32)(sub_4B0680(0, math.MaxUint8))
	result = nox_xxx_drawPartFx2_4AF990(0, 0, 0, math.MaxInt32)
	v3 = result
	if result != nil {
		sub_4AFB10(int32(uintptr(unsafe.Pointer(result))), *a1)
		sub_4AFB00(int32(uintptr(unsafe.Pointer(v3))), int32(uintptr(unsafe.Pointer(v1))))
		result = (*byte)(unsafe.Pointer(uintptr(nox_xxx_registerParticleFx_4AFCF0(int32(cgoFuncAddr(sub_4AF4C0)), int32(uintptr(unsafe.Pointer(v3))), 1, *(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*7))))))
	}
	return result
}
func sub_4AF4C0(a1 int32) int32 {
	var (
		v1  int32
		v2  int32
		v3  int32
		v4  int32
		v5  *byte
		v6  int32
		v7  int32
		v8  int32
		v9  int32
		v11 int32
	)
	v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 20))) + 8))))
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 12))))
	v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 16))))
	nox_client_drawSetColor_434460(int32(nox_color_yellow_2589772))
	v11 = sub_48C5E0(v3-2, v3+2)
	v4 = sub_48C5E0(v2-2, v2+2)
	v5 = nox_xxx_drawPartFx2_4AF990(v4, v11, 2, 10)
	v6 = int32(uintptr(unsafe.Pointer(v5)))
	if v5 == nil {
		return 0
	}
	nox_xxx_partfxLoadParticle_4AFE20((*uint32)(unsafe.Pointer(v5)), *(**byte)(memmap.PtrOff(0x973A20, 472)))
	sub_4AFB00(v6, int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 20))) + 4)))))
	sub_4AFB50(v6, 1)
	sub_4AFC50(v6, math.MinInt16)
	sub_4AFC60(v6, 0xB333, 1)
	*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v7))), unsafe.Sizeof(uint16(0))*0)) = uint16(sub_48C610())
	sub_4AFBD0(v6, v7)
	*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v8))), unsafe.Sizeof(uint16(0))*0)) = uint16(sub_48C610())
	sub_4AFBE0(v6, v8)
	*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v9))), unsafe.Sizeof(uint16(0))*0)) = uint16(sub_48C610())
	sub_4AFBF0(v6, v9)
	return 1
}
func nox_xxx_ParticleFxT6_4AF5A0(a1 *int32) *byte {
	var (
		result *byte
		v2     *byte
	)
	result = nox_xxx_drawPartFx2_4AF990(0, 0, 0, math.MaxInt32)
	v2 = result
	if result != nil {
		sub_4AFB40(int32(uintptr(unsafe.Pointer(result))), *(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*2)))
		sub_4AFB10(int32(uintptr(unsafe.Pointer(v2))), *a1)
		sub_4AFB60(int32(uintptr(unsafe.Pointer(v2))), 0)
		sub_4AFB70(int32(uintptr(unsafe.Pointer(v2))), 0)
		result = (*byte)(unsafe.Pointer(uintptr(nox_xxx_registerParticleFx_4AFCF0(int32(cgoFuncAddr(nox_xxx_partFx_4AF600)), int32(uintptr(unsafe.Pointer(v2))), 1, 500))))
	}
	return result
}
func sub_4AF650(a1 *int32) int32 {
	nox_client_drawEnableAlpha_434560(1)
	nox_client_drawSetAlpha_434580(uint8(int8(*(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*12)) * math.MaxUint8 / *(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*13)))))
	nox_xxx_drawParticlefx_4AFEB0(a1)
	nox_client_drawEnableAlpha_434560(0)
	return 1
}
func nox_xxx_partFx_4AF600(a1 int32) int32 {
	var v1 *byte
	v1 = nox_xxx_drawPartFx2_4AF990(0, 0, 0, 5)
	if v1 == nil {
		return 0
	}
	sub_4AFB40(int32(uintptr(unsafe.Pointer(v1))), int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 20))) + 16)))))
	sub_4AFB10(int32(uintptr(unsafe.Pointer(v1))), int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 20))) + 12)))))
	sub_4AFB70(int32(uintptr(unsafe.Pointer(v1))), int32(cgoFuncAddr(sub_4AF650)))
	return 1
}
func nox_xxx_partfxSwitch_4AF690(a1 *uint32, a2 func(*uint32, *uint32, int32)) *uint32 {
	var (
		result *uint32
		v3     int32
		v4     int32
		v5     *int32
		v6     int32
		v7     int32
		v8     int32
		v9     int32
		v10    *uint8
		v11    int32
		v12    int32
		v13    int32
		v14    int32
		v15    *uint8
		v16    func(*uint32, *uint32, int32)
		v17    *uint32
		v18    int32
		v19    int32
		v20    *uint32
		v21    int32
		v22    int32
		v23    int32
		v24    int32
		v25    int32
		v26    int32
		v27    int32
		v28    int32
		v29    *uint32
		v30    int32
		v31    int32
		v32    *uint8
	)
	result = (*uint32)(nox_video_getImagePixdata_func((*nox_video_bag_image_t)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*1)))))))
	if result != nil {
		v3 = int32(*result)
		v4 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*1)))
		v5 = (*int32)(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*3))))
		result = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*2)) + *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*3)))))
		v22 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*5)))
		v30 = v3
		v6 = *v5
		v7 = int32(uintptr(unsafe.Pointer(v5))) + 5
		v8 = int32(uint32(v6) + *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*4)))
		v9 = 0
		v27 = v4
		v29 = result
		v10 = (*uint8)(unsafe.Pointer(uintptr(v7)))
		if v4 != 0 {
			v23 = v4
			for {
				v20 = result
				v11 = v3
				if v3 > 0 {
					for {
						v12 = int32(*(*uint8)(unsafe.Add(unsafe.Pointer(v10), 1)))
						v13 = int32(*v10) & 15
						v10 = (*uint8)(unsafe.Add(unsafe.Pointer(v10), 2))
						switch v13 {
						case 1:
							v20 = (*uint32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(v20))), v12))))
						case 2:
							fallthrough
						case 5:
							fallthrough
						case 6:
							v9 += v12
							v10 = (*uint8)(unsafe.Add(unsafe.Pointer(v10), v12*2))
						case 4:
							v9 += v12
							v10 = (*uint8)(unsafe.Add(unsafe.Pointer(v10), v12))
						default:
						}
						v11 -= v12
						if v11 <= 0 {
							break
						}
					}
					result = v29
					v3 = v30
				}
				v23--
				if v23 == 0 {
					break
				}
			}
			v4 = v27
		}
		v21 = v8
		v14 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*5)))
		v15 = (*uint8)(unsafe.Pointer(uintptr(v7)))
		if v4 != 0 {
			v31 = v4
			v16 = a2
			for {
				v17 = result
				v28 = v3
				if v3 > 0 {
					for {
						v18 = int32(*(*uint8)(unsafe.Add(unsafe.Pointer(v15), 1)))
						v19 = int32(*v15) & 15
						v15 = (*uint8)(unsafe.Add(unsafe.Pointer(v15), 2))
						v26 = v18
						v32 = v15
						switch v19 {
						case 1:
							v17 = (*uint32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(v17))), v18))))
						case 2:
							fallthrough
						case 5:
							fallthrough
						case 6:
							if v18 > 0 {
								v25 = v18
								for {
									v14 += v22
									if v14 >= v9 {
										v14 -= v9
										if v16 != nil {
											v16(a1, v17, v21)
											v15 = v32
											v18 = v26
										}
									}
									v17 = (*uint32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(v17))), 1))))
									v25--
									if v25 == 0 {
										break
									}
								}
							}
							v15 = (*uint8)(unsafe.Add(unsafe.Pointer(v15), v18*2))
						case 4:
							if v18 > 0 {
								v24 = v18
								for {
									v14 += v22
									if v14 >= v9 {
										v14 -= v9
										if v16 != nil {
											v16(a1, v17, v21)
											v15 = v32
											v18 = v26
										}
									}
									v17 = (*uint32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(v17))), 1))))
									v24--
									if v24 == 0 {
										break
									}
								}
							}
							v15 = (*uint8)(unsafe.Add(unsafe.Pointer(v15), v18))
						default:
						}
						v28 -= v18
						if v28 <= 0 {
							break
						}
					}
					result = v29
					v3 = v30
				}
				v21++
				v31--
				if v31 == 0 {
					break
				}
			}
		}
	}
	return result
}
func sub_4AF890(a1 **uint8) *uint8 {
	var result *uint8
	result = *a1
	*(**uint8)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof((*uint8)(nil))*3)) = (*uint8)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(*a1))), unsafe.Sizeof(uint32(0))*3))) - uint32(**a1))))
	*(**uint8)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof((*uint8)(nil))*4)) = (*uint8)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(result))), unsafe.Sizeof(uint32(0))*4))) - uint32(*((*int16)(unsafe.Add(unsafe.Pointer((*int16)(unsafe.Pointer(result))), unsafe.Sizeof(int16(0))*53)))) - uint32(*((*int16)(unsafe.Add(unsafe.Pointer((*int16)(unsafe.Pointer(result))), unsafe.Sizeof(int16(0))*52)))) - uint32(*(*uint8)(unsafe.Add(unsafe.Pointer(result), 1))))))
	return result
}
func sub_4AF8C0(a1 *uint32) int32 {
	var result int32
	result = int32(*(*uint32)(unsafe.Pointer(uintptr(*a1 + 8))))
	*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*1)) = uint32(result)
	return result
}
func sub_4AF8D0() int32 {
	var v1 *uint8
	*memmap.PtrUint32(0x973A20, 472) = uint32(uintptr(memmap.PtrOff(0x587000, 174348)))
	dword_5d4594_1311148 = 512
	dword_5d4594_1311140 = uint32(uintptr(alloc.Calloc(512, 136)))
	if dword_5d4594_1311140 == 0 {
		return 0
	}
	dword_5d4594_1311144 = uint32(uintptr(alloc.Calloc(512, 1)))
	if dword_5d4594_1311144 == 0 {
		return 0
	}
	v1 = (*uint8)(memmap.PtrOff(6112660, 1311160))
	for {
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v1))), -int(unsafe.Sizeof(uint32(0))*1)))) = 0
		*(*uint32)(unsafe.Pointer(v1)) = 0
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v1))), unsafe.Sizeof(uint32(0))*1))) = 0
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v1))), unsafe.Sizeof(uint32(0))*2))) = 0
		v1 = (*uint8)(unsafe.Add(unsafe.Pointer(v1), 24))
		if int32(uintptr(unsafe.Pointer(v1))) >= int32(uintptr(memmap.PtrOff(6112660, 1311928))) {
			break
		}
	}
	return 1
}
func sub_4AF950() {
	if dword_5d4594_1311140 != 0 {
		*(*unsafe.Pointer)(unsafe.Pointer(&dword_5d4594_1311140)) = nil
		dword_5d4594_1311140 = 0
	}
	if dword_5d4594_1311144 != 0 {
		*(*unsafe.Pointer)(unsafe.Pointer(&dword_5d4594_1311144)) = nil
		dword_5d4594_1311144 = 0
	}
	dword_5d4594_1311148 = 0
}
func nox_xxx_drawPartFx2_4AF990(a1 int32, a2 int32, a3 int32, a4 int32) *byte {
	var result *byte
	result = nox_xxx_partfxAllocSmth_4B01B0()
	if result != nil {
		*(*uint32)(unsafe.Pointer(result)) = 0
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(result))), unsafe.Sizeof(uint32(0))*1))) = 0
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(result))), unsafe.Sizeof(uint32(0))*4))) = 0
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(result))), unsafe.Sizeof(uint32(0))*3))) = 0
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(result))), unsafe.Sizeof(uint32(0))*5))) = 0
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(result))), unsafe.Sizeof(uint32(0))*6))) = nox_draw_curDrawData_3799572.field_61
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(result))), unsafe.Sizeof(uint32(0))*7))) = uint32(a1)
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(result))), unsafe.Sizeof(uint32(0))*8))) = uint32(a2)
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(result))), unsafe.Sizeof(uint32(0))*9))) = uint32(a3)
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(result))), unsafe.Sizeof(uint32(0))*22))) = uint32(a3 << 16)
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(result))), unsafe.Sizeof(uint32(0))*21))) = uint32(a2 << 16)
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(result))), unsafe.Sizeof(uint32(0))*12))) = uint32(a4)
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(result))), unsafe.Sizeof(uint32(0))*13))) = uint32(a4)
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(result))), unsafe.Sizeof(uint32(0))*10))) = 0
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(result))), unsafe.Sizeof(uint32(0))*11))) = 0
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(result))), unsafe.Sizeof(uint32(0))*14))) = 0
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(result))), unsafe.Sizeof(uint32(0))*20))) = uint32(a1 << 16)
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(result))), unsafe.Sizeof(uint32(0))*23))) = 0
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(result))), unsafe.Sizeof(uint32(0))*24))) = 0
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(result))), unsafe.Sizeof(uint32(0))*25))) = 0
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(result))), unsafe.Sizeof(uint32(0))*26))) = 0
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(result))), unsafe.Sizeof(uint32(0))*27))) = 0
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(result))), unsafe.Sizeof(uint32(0))*28))) = 0
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(result))), unsafe.Sizeof(uint32(0))*29))) = 0
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(result))), unsafe.Sizeof(uint32(0))*30))) = 0
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(result))), unsafe.Sizeof(uint32(0))*16))) = 0
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(result))), unsafe.Sizeof(uint32(0))*32))) = uint32(cgoFuncAddr(sub_4B0020))
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(result))), unsafe.Sizeof(uint32(0))*31))) = uint32(cgoFuncAddr(nox_xxx_drawParticlefx_4AFEB0))
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(result))), unsafe.Sizeof(uint32(0))*33))) = uint32(cgoFuncAddr(sub_4B01A0))
	}
	return result
}
func sub_4AFA40(a1 int32) {
	var (
		v1 int32
		v2 *uint32
	)
	if a1 != 0 {
		v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 60))))
		if v1 >= 0 && v1 < *(*int32)(unsafe.Pointer(&dword_5d4594_1311148)) {
			*(*uint8)(unsafe.Pointer(uintptr(dword_5d4594_1311144 + uint32(v1)))) = 0
			v2 = *(**uint32)(unsafe.Pointer(uintptr(a1 + 20)))
			if v2 != nil {
				*v2 = 0
			}
		}
	}
}
func sub_4AFA70(a1 int32) *byte {
	var (
		result *byte
		i      *byte
		v3     int32
	)
	result = libc.MemChr((*byte)(*(*unsafe.Pointer)(unsafe.Pointer(&dword_5d4594_1311144))), math.MaxUint8, int(dword_5d4594_1311148))
	for i = result; result != nil; i = result {
		v3 = int32(uintptr(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(i), -int(dword_5d4594_1311144))))))
		if *(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1311140 + uint32(uintptr(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(i), -int(dword_5d4594_1311144))))))*136 + 8))) == uint32(a1) {
			sub_4AFA40(int32(dword_5d4594_1311140 + uint32(uintptr(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(i), -int(dword_5d4594_1311144))))))*136))
		}
		result = (*byte)(unsafe.Pointer(uintptr(dword_5d4594_1311148 - uint32(v3) - 1)))
		if dword_5d4594_1311148-uint32(v3) == 1 {
			break
		}
		result = libc.MemChr((*byte)(unsafe.Add(unsafe.Pointer(i), 1)), math.MaxUint8, int(dword_5d4594_1311148-uint32(v3)-1))
	}
	return result
}
func sub_4AFAF0(a1 *uint32, a2 int32) int32 {
	var result int32
	result = a2
	*a1 = uint32(a2)
	return result
}
func sub_4AFB00(a1 int32, a2 int32) int32 {
	var result int32
	result = a2
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 4))) = uint32(a2)
	return result
}
func sub_4AFB10(a1 int32, a2 int32) int32 {
	var (
		result int32
		v3     int32
		v4     uint32
	)
	result = a1
	v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 56))))
	if a2 != 0 {
		v4 = uint32(v3 | 4)
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 8))) = uint32(a2)
	} else {
		v4 = uint32(v3) & 0xFFFFFFFB
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 8))) = 0
	}
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 56))) = v4
	return result
}
func sub_4AFB40(a1 int32, a2 int32) int32 {
	var result int32
	result = a2
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 16))) = uint32(a2)
	return result
}
func sub_4AFB50(a1 int32, a2 int32) int32 {
	var result int32
	result = a2
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 64))) = uint32(a2)
	return result
}
func sub_4AFB60(a1 int32, a2 int32) int32 {
	var result int32
	result = a2
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 128))) = uint32(a2)
	return result
}
func sub_4AFB70(a1 int32, a2 int32) int32 {
	var result int32
	result = a2
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 124))) = uint32(a2)
	return result
}
func sub_4AFB80(a1 int32, a2 int32) int32 {
	var result int32
	result = a2
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 132))) = uint32(a2)
	return result
}
func sub_4AFB90(a1 int32, a2 int32) int32 {
	var result int32
	result = a2
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 44))) = uint32(a2)
	return result
}
func sub_4AFBA0(a1 int32, a2 int32) int32 {
	var result int32
	result = a2
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 48))) = uint32(a2)
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 52))) = uint32(a2)
	return result
}
func sub_4AFBB0(a1 *uint32, a2 int32, a3 int32, a4 int32) *uint32 {
	var result *uint32
	result = a1
	*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*23)) = uint32(a2)
	*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*24)) = uint32(a3)
	*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*25)) = uint32(a4)
	return result
}
func sub_4AFBD0(a1 int32, a2 int32) int32 {
	var result int32
	result = a2
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 92))) = uint32(a2)
	return result
}
func sub_4AFBE0(a1 int32, a2 int32) int32 {
	var result int32
	result = a2
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 96))) = uint32(a2)
	return result
}
func sub_4AFBF0(a1 int32, a2 int32) int32 {
	var result int32
	result = a2
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 100))) = uint32(a2)
	return result
}
func sub_4AFC00(a1 *uint32, a2 int32, a3 int32, a4 int32) *uint32 {
	var result *uint32
	result = a1
	*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*26)) = uint32(a2)
	*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*27)) = uint32(a3)
	*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*28)) = uint32(a4)
	return result
}
func sub_4AFC20(a1 int32, a2 int32) int32 {
	var result int32
	result = a2
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 104))) = uint32(a2)
	return result
}
func sub_4AFC30(a1 int32, a2 int32) int32 {
	var result int32
	result = a2
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 108))) = uint32(a2)
	return result
}
func sub_4AFC40(a1 int32, a2 int32) int32 {
	var result int32
	result = a2
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 112))) = uint32(a2)
	return result
}
func sub_4AFC50(a1 int32, a2 int32) int32 {
	var result int32
	result = a2
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 116))) = uint32(a2)
	return result
}
func sub_4AFC60(a1 int32, a2 int32, a3 int32) int32 {
	var (
		result int32
		v4     int32
	)
	result = a1
	v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 56))))
	if a3 != 0 {
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 56))) = uint32(v4 | 2)
	} else {
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 56))) = uint32(v4) & 0xFFFFFFFD
	}
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 120))) = uint32(a2)
	return result
}
func sub_4AFC90(a1 *uint32, a2 int32, a3 int32, a4 int32) {
	var (
		v4 int32
		v5 int32
	)
	v4 = a4
	if a4 == 0 {
		v4 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*13)))
	}
	if v4 != 0 {
		v5 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*8)))
		*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*23)) = ((uint32(a2) - *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*7))) << 16) / uint32(v4)
		*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*24)) = uint32(((a3 - v5) << 16) / v4)
	}
}
func sub_4AFCD0(a1 *uint32) *uint32 {
	var (
		result *uint32
		v2     int32
		v3     int32
	)
	result = a1
	v2 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*21)))
	*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*7)) = *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*20)) >> 16
	v3 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*22)) >> 16)
	*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*8)) = uint32(v2 >> 16)
	*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*9)) = uint32(v3)
	return result
}
func nox_xxx_registerParticleFx_4AFCF0(a1 int32, a2 int32, a3 int32, a4 int32) int32 {
	var (
		v4 *uint8
		v5 int32
	)
	v4 = (*uint8)(memmap.PtrOff(6112660, 1311156))
	v5 = 0
	for *(*uint32)(unsafe.Pointer(v4)) != 0 {
		v4 = (*uint8)(unsafe.Add(unsafe.Pointer(v4), 24))
		if func() int32 {
			p := &v5
			*p++
			return *p
		}() >= 32 {
			return 0
		}
	}
	*(*uint32)(unsafe.Pointer(v4)) = uint32(a1)
	*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v4))), unsafe.Sizeof(uint32(0))*5))) = uint32(a2)
	*(*uint32)(unsafe.Pointer(uintptr(a2 + 20))) = uint32(uintptr(unsafe.Pointer(v4)))
	*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v4))), unsafe.Sizeof(uint32(0))*2))) = uint32(a3)
	*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v4))), unsafe.Sizeof(uint32(0))*1))) = uint32(a3)
	*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v4))), unsafe.Sizeof(uint32(0))*3))) = 0
	*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v4))), unsafe.Sizeof(uint32(0))*4))) = uint32(a4)
	return 1
}
func sub_4AFD40() {
	var (
		i   *byte
		v1  int32
		v2  int32
		v3  func(int32)
		v4  func(int32)
		v5  *uint8
		v6  int32
		v7  *func(int32) int32
		v8  int32
		v9  int32
		v10 int32
		v11 int32
	)
	for i = libc.MemChr((*byte)(*(*unsafe.Pointer)(unsafe.Pointer(&dword_5d4594_1311144))), math.MaxUint8, int(dword_5d4594_1311148)); i != nil; i = libc.MemChr((*byte)(unsafe.Add(unsafe.Pointer(i), 1)), math.MaxUint8, int(dword_5d4594_1311148-uint32(v1)-1)) {
		v1 = int32(uintptr(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(i), -int(dword_5d4594_1311144))))))
		v2 = int32(dword_5d4594_1311140 + uint32(uintptr(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(i), -int(dword_5d4594_1311144))))))*136)
		if (int32(*(*uint8)(unsafe.Pointer(uintptr(v2 + 56)))) & 8) == 0 {
			v3 = cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(v2 + 124))), (*func(int32))(nil))
			if v3 != nil {
				v3(int32(dword_5d4594_1311140 + uint32(uintptr(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(i), -int(dword_5d4594_1311144))))))*136))
			}
		}
		v4 = cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(v2 + 128))), (*func(int32))(nil))
		if v4 != nil {
			v4(v2)
		}
		if dword_5d4594_1311148-uint32(v1) == 1 {
			break
		}
	}
	v5 = (*uint8)(memmap.PtrOff(6112660, 1311168))
	v6 = 32
	for {
		v7 = (*func(int32) int32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v5), -12))))
		if *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v5))), -int(unsafe.Sizeof(uint32(0))*3)))) != 0 {
			v8 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v5))), -int(unsafe.Sizeof(uint32(0))*2)))))
			if v8 != 0 {
				if v8 == 1 {
					v9 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v5))), unsafe.Sizeof(uint32(0))*1))))
					v10 = int32(*(*uint32)(unsafe.Pointer(v5)) + 1)
					*(*uint32)(unsafe.Pointer(v5)) = uint32(v10)
					if v10 <= v9 && (*v7)(int32(uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v5), -12)))))) != 0 {
						*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v5))), -int(unsafe.Sizeof(uint32(0))*2)))) = *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v5))), -int(unsafe.Sizeof(uint32(0))*1))))
					} else {
						v11 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v5))), unsafe.Sizeof(uint32(0))*2))))
						*v7 = nil
						sub_4AFA40(v11)
					}
				}
			}
		}
		v5 = (*uint8)(unsafe.Add(unsafe.Pointer(v5), 24))
		v6--
		if v6 == 0 {
			break
		}
	}
}
func nox_xxx_partfxLoadParticle_4AFE20(a1 *uint32, a2 *byte) int32 {
	var (
		v2     int32
		v3     *uint32
		v4     *uint32
		result int32
	)
	v2 = nox_xxx_getTTByNameSpriteMB_44CFC0(a2)
	v3 = &nox_xxx_spriteLoadAdd_45A360_drawable(v2, int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*20))>>16), int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*21))>>16)).field_0
	v4 = v3
	if v3 != nil {
		*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*3)) = uint32(uintptr(unsafe.Pointer(v3)))
		*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*108)) = uint32(uintptr(unsafe.Pointer(a1)))
		nox_xxx_spriteTransparentDecay_49B950(v3, int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*13))))
		nox_xxx_sprite_45A110_drawable((*nox_drawable)(unsafe.Pointer(v4)))
	}
	result = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*14)))
	*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&result))), 0)) = uint8(int8(result | 8))
	*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*14)) = uint32(result)
	return result
}
func nox_xxx_drawParticlefx_4AFEB0(a1 *int32) int32 {
	var (
		v1     *int32
		v2     int32
		v3     int32
		v4     int32
		v5     int32
		v6     int32
		v7     int32
		v8     int32
		result int32
		v10    int32
		a1a    int2
		a2     int2
		v13    int32
	)
	v1 = a1
	v2 = *(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*22)) >> 16
	v3 = *(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*21)) >> 16
	a1a.field_0 = *(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*20)) >> 16
	a1a.field_4 = v3 - v2
	sub_4739A0(&a1a, &a2)
	v4 = a2.field_0
	v5 = a2.field_4
	if *a1 != 0 {
		nox_client_drawImageAt_47D2C0((*nox_video_bag_image_t)(unsafe.Pointer(uintptr(*a1))), a2.field_0, a2.field_4)
	}
	v6 = *(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*1))
	if v6 != 0 {
		nox_video_drawImageAt2_4B0820(unsafe.Pointer(uintptr(v6)), v4, v5)
	}
	v7 = *(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*3))
	if v7 != 0 && *(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*4)) != 0 {
		v8 = int32(*(*uint32)(unsafe.Pointer(uintptr(v7 + 12))))
		v13 = int32(*(*uint32)(unsafe.Pointer(uintptr(v7 + 16))))
		*(*uint32)(unsafe.Pointer(uintptr(v7 + 12))) = uint32(a1a.field_0)
		*(*uint32)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(int32(0))*3)) + 16))) = uint32(a1a.field_4)
		(cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(int32(0))*3)) + 300))), (*func(int32, int32))(nil)))(*(*int32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(int32(0))*4)), *(*int32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(int32(0))*3)))
		*(*uint32)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(int32(0))*3)) + 12))) = uint32(v8)
		*(*uint32)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(int32(0))*3)) + 16))) = uint32(v13)
	}
	result = *(*int32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(int32(0))*16))
	if result != 0 {
		nox_client_drawSetColor_434460(*(*int32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(int32(0))*6)))
		switch *(*int32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(int32(0))*16)) {
		case 1:
			result = sub_49EFA0(v4, v5)
		case 2:
			sub_49EFA0(v4, v5)
			sub_49EFA0(v4+1, v5)
			v10 = v5 + 1
			sub_49EFA0(v4, v10)
			result = sub_49EFA0(v4+1, v10)
		case 3:
			sub_49EFA0(v4+1, v5)
			sub_49EFA0(v4, v5+1)
			sub_49EFA0(v4+1, v5+1)
			sub_49EFA0(v4+2, v5+1)
			result = sub_49EFA0(v4+1, v5+2)
		default:
			result = sub_4B0BC0(v4, v5, *(*int32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(int32(0))*16))/2)
		}
	}
	return result
}
func sub_4B0020(a1 *uint32) int32 {
	var (
		v1         *uint32
		result     int32
		v3         int32
		v4         int32
		v5         int32
		v6         *uint32
		v7         int32
		v8         int32
		v9         int32
		v10        int32
		v11        bool
		v12        int32
		v13        int32
		v14        int32
		v15        int32
		v16        int32
		v17        int32
		v18        int32
		v19        int32
		v20        int32
		v21        int32
		v22        int32
		nNumerator int32
		v24        int32
		v25        int32
	)
	v1 = a1
	result = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*12)))
	if result >= 0 {
		v3 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*20)))
		v4 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*22)))
		*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*10))++
		v20 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*21)))
		v5 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*14)))
		v24 = v3
		if v5&4 != 0 {
			v6 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*2)))))
			if v6 != nil {
				v24 += int32((*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*3)) - *(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*8))) << 16)
				v20 += int32((*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*4)) - *(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*9))) << 16)
			}
		}
		v7 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*11)))
		if v7 != 0 {
			result = v7 - 1
			*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*11)) = uint32(result)
		} else {
			v8 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*23)))
			v9 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*24)))
			v10 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*25)))
			v25 = v8 + v24
			v11 = v10+v4 < 0
			v12 = v10 + v4
			v21 = v9 + v20
			v22 = v12
			if v11 {
				if v5&2 != 0 {
					v13 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*30)))
					v14 = v12 - v10
					v10 = -v10
					v22 = v14
					nNumerator = v13
					if v13 != 0 {
						v8 = compatMulDiv(v8, v13, 0x10000)
						v9 = compatMulDiv(v9, nNumerator, 0x10000)
						v10 = compatMulDiv(v10, nNumerator, 0x10000)
					}
				} else {
					v10 = 0
					v22 = 0
				}
			}
			v15 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*26)))
			if v15 != 0 {
				v8 = compatMulDiv(v8, v15, 0x10000)
			}
			v16 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*27)))
			if v16 != 0 {
				v9 = compatMulDiv(v9, v16, 0x10000)
			}
			v17 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*28)))
			if v17 != 0 {
				v10 = compatMulDiv(v10, v17, 0x10000)
			}
			v18 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*29)) + uint32(v10))
			v19 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*12)) - 1)
			*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*12)) = uint32(v19)
			if v19 >= 0 {
				result = v22
				*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*25)) = uint32(v18)
				*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*20)) = uint32(v25)
				*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*21)) = uint32(v21)
				*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*22)) = uint32(v22)
				*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*23)) = uint32(v8)
				*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*24)) = uint32(v9)
			} else {
				result = (cgoAsFunc(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*33)), (*func(*uint32) int32)(nil)).(func(*uint32) int32))(v1)
			}
		}
	}
	return result
}
func sub_4B01A0(a1 int32) {
	sub_4AFA40(a1)
}
func nox_xxx_partfxAllocSmth_4B01B0() *byte {
	var (
		v0     *byte
		result *byte
		v2     int32
	)
	v0 = libc.MemChr((*byte)(*(*unsafe.Pointer)(unsafe.Pointer(&dword_5d4594_1311144))), 0, int(dword_5d4594_1311148))
	if v0 == nil {
		result = (*byte)(unsafe.Pointer(uintptr(sub_4B0220(dword_5d4594_1311148 + 512))))
		if result == nil {
			return result
		}
		v0 = libc.MemChr((*byte)(*(*unsafe.Pointer)(unsafe.Pointer(&dword_5d4594_1311144))), 0, int(dword_5d4594_1311148))
	}
	v2 = int32(uintptr(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v0), -int(dword_5d4594_1311144))))))
	*(*uint8)(unsafe.Pointer(uintptr(dword_5d4594_1311144 + uint32(v2)))) = math.MaxUint8
	result = (*byte)(unsafe.Pointer(uintptr(dword_5d4594_1311140 + uint32(v2*136))))
	*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(result))), unsafe.Sizeof(uint32(0))*15))) = uint32(v2)
	return result
}
func sub_4B0220(a1 uint32) int32 {
	var (
		v1     uint32
		result int32
	)
	v1 = a1 - dword_5d4594_1311148
	result = int32(uintptr(libc.Realloc(*(*unsafe.Pointer)(unsafe.Pointer(&dword_5d4594_1311140)), int(a1*136))))
	if result != 0 {
		dword_5d4594_1311140 = uint32(result)
		libc.MemSet(unsafe.Pointer(uintptr(uint32(result)+dword_5d4594_1311148*136)), 0, int(v1*136))
		result = int32(uintptr(libc.Realloc(*(*unsafe.Pointer)(unsafe.Pointer(&dword_5d4594_1311144)), int(a1))))
		if result != 0 {
			dword_5d4594_1311144 = uint32(result)
			libc.MemSet(unsafe.Pointer(uintptr(dword_5d4594_1311148+uint32(result))), 0, int(v1))
			dword_5d4594_1311148 = a1
			result = 1
		}
	}
	return result
}
func sub_4B0610(a1 int32) {
	if dword_5d4594_1311936 != 0 && (*memmap.PtrUint32(6112660, 1311932) == 0 || a1 == 1) {
		sub_555500(1)
	}
}
func sub_4B0870(a1 *int32) int32 {
	var (
		v1     int32
		v2     *uint32
		v3     int32
		v4     int32
		v5     int32
		v6     *uint8
		v7     int32
		result int32
		v9     int32
		v10    *uint8
		v11    *uint8
		v12    int32
		v13    int32
		v14    int32
		v15    *uint8
		v16    int32
		v17    bool
		v18    int32
		v19    int32
		v20    int32
		v21    int32
		v22    int32
		v23    int32
		v24    int32
		v25    uint8
		v26    int32
		v27    int32
		v28    int32
		v29    uint8
		v30    int32
		v31    int32
		v32    uint8
		v33    int32
		v34    int32
		v35    int32
		v36    int32
		v37    int32
	)
	v1 = *a1
	v2 = (*uint32)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*16)))))
	v3 = *(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*1))
	v4 = *(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*2))
	*v2 = uint32(*a1 * 2)
	v2 = (*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*1))
	*v2 = uint32(v1 * 2)
	v2 = (*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*1))
	*v2 = uint32(-*a1)
	v2 = (*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*1))
	*v2 = uint32(-*a1)
	v32 = uint8(int8(int32(uint16(int16(v3**(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*12))))) >> 8))
	v5 = *(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*13))
	v23 = *(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*14))
	v29 = uint8(int8(int32(uint16(int16(v3*v5))) >> 8))
	v25 = uint8(int8(int32(uint16(int16(v3*v23))) >> 8))
	v6 = (*uint8)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(v2))), 5))
	v34 = int32(v32) << 16
	v18 = int32(v32) << 16
	v7 = -v1
	v31 = -v1
	v27 = int32((((uint32(v4**(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*12)))>>8)&math.MaxUint8)-uint32(v32))<<16) / v1
	v35 = int32(v29) << 16
	v19 = int32(v29) << 16
	v28 = int32((((uint32(v4*v5)>>8)&math.MaxUint8)-uint32(v29))<<16) / v1
	v36 = int32(v25) << 16
	v20 = int32(v25) << 16
	v30 = int32((((uint32(v4*v23)>>8)&math.MaxUint8)-uint32(v25))<<16) / v1
	result = v1 * 2
	if v1*2 != 0 {
		v33 = v1 * 2
		for {
			v9 = int32(int64(math.Sqrt(float64(v1*v1 - v7*v7))))
			if v9 < v1 {
				*v6 = 1
				v10 = (*uint8)(unsafe.Add(unsafe.Pointer(v6), 1))
				*v10 = uint8(int8(v1 - v9))
				v6 = (*uint8)(unsafe.Add(unsafe.Pointer(v10), 1))
			}
			v26 = v9 * 2
			if v9*2 != 0 {
				*v6 = 7
				v11 = (*uint8)(unsafe.Add(unsafe.Pointer(v6), 1))
				v12 = v34
				v13 = v35
				*v11 = uint8(int8(v9 * 2))
				v6 = (*uint8)(unsafe.Add(unsafe.Pointer(v11), 1))
				v37 = v36
				v21 = v18 / v9
				v22 = v19 / v9
				v14 = -v9
				v24 = v20 / v9
				for {
					*(*uint16)(unsafe.Pointer(v6)) = uint16(int16((v37>>19)&31 | (int32(*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v13))), 2)))&252|(int32(*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v12))), unsafe.Sizeof(uint16(0))*1)))&0xFFF8)*32)*8))
					if v14 >= 0 {
						v12 -= v21
						v13 -= v22
						v37 -= v24
					} else {
						v12 += v21
						v13 += v22
						v37 += v24
					}
					v14++
					v6 = (*uint8)(unsafe.Add(unsafe.Pointer(v6), 2))
					v26--
					if v26 == 0 {
						break
					}
				}
				v7 = v31
			}
			if v9 < v1 {
				*v6 = 1
				v15 = (*uint8)(unsafe.Add(unsafe.Pointer(v6), 1))
				*v15 = uint8(int8(v1 - v9))
				v6 = (*uint8)(unsafe.Add(unsafe.Pointer(v15), 1))
			}
			if v7 >= 0 {
				v19 -= v28
				v18 -= v27
				v16 = v20 - v30
			} else {
				v19 += v28
				v18 += v27
				v16 = v30 + v20
			}
			v20 = v16
			v7++
			result = v33 - 1
			v17 = v33 == 1
			v31 = v7
			v33--
			if v17 {
				break
			}
		}
	}
	return result
}
func sub_4B0BC0(a1 int32, a2 int32, a3 int32) int32 {
	var (
		v3     int32
		v4     func(int32, int32, int32) int32
		v5     int32
		v6     int32
		result int32
		v8     int32
		v9     int32
		v10    int32
		v11    int32
		v12    int32
		v13    int32
		v14    int32
		v15    int32
		v16    int32
	)
	v3 = a3
	v16 = 1 - a3
	v4 = cgoAsFunc(*(*uint32)(unsafe.Pointer(&dword_5d4594_3798708)), (*func(int32, int32, int32) int32)(nil))
	v10 = 5 - v3*2
	v12 = 0
	v13 = v3
	v11 = 3
	if nox_draw_curDrawData_3799572.field_13 == 0 {
		v4 = cgoAsFunc(*(*uint32)(unsafe.Pointer(&dword_5d4594_3798720)), (*func(int32, int32, int32) int32)(nil))
	}
	v5 = a2
	v6 = a1
	v4(a1, a2+v3, a1)
	v14 = a1 + v3
	v15 = v6 - v3
	v4(v6-v3, v5, v6+v3)
	v4(v6-v3, v5, v14)
	result = v4(v6, v5-v3, v6)
	if v3 > 0 {
		v8 = v6
		v9 = v5 - v6
		for {
			if v16 >= 0 {
				v16 += v10
				v10 += 4
				v13--
				v14--
				v15++
			} else {
				v16 += v11
				v10 += 2
			}
			v6++
			v12++
			v8--
			v11 += 2
			v4(v8, v9+v14, v6)
			v4(v15, v9+v6, v14)
			v4(v15, v9+v8, v14)
			v4(v8, v9+v15, v6)
			result = v13
			if v13 <= v12 {
				break
			}
		}
	}
	return result
}
func sub_4B0F50(a1 int32, a2 int32, a3 int32) int32 {
	var (
		v3        int32
		v4        int8
		v5        int32
		result    int32
		v7        int32
		v8        int32
		v9        int32
		v10       int32
		v11       int32
		v12       int32
		v13       int32
		v14       int32
		v15       int32
		v16       int32
		v17       int32
		v18       int32
		v19       int32
		v20       int32
		v21       int32
		v22       int32
		v23       int32
		v24       int32
		v25       int32
		v26       int32
		v27       int32
		v28       int32
		v29       int32
		v30       int32
		v31       int8
		pixbuffer **uint8 = nox_pixbuffer_rows_3798784
	)
	v3 = a3
	v28 = 0
	v20 = 1 - a3
	v19 = 5 - a3*2
	v27 = a3
	v26 = 3
	v4 = int8(*(*uint8)(unsafe.Pointer(&nox_draw_curDrawData_3799572.field_61)))
	v17 = nox_draw_curDrawData_3799572.clip.left
	v5 = a2
	v15 = nox_draw_curDrawData_3799572.clip.bottom
	v31 = int8(*(*uint8)(unsafe.Pointer(&nox_draw_curDrawData_3799572.field_61)))
	result = a2 + v27
	v16 = nox_draw_curDrawData_3799572.clip.right
	v18 = nox_draw_curDrawData_3799572.clip.top
	v24 = a2 + v27
	if a1 >= v17 && a1 < nox_draw_curDrawData_3799572.clip.right && result >= nox_draw_curDrawData_3799572.clip.top && result < v15 {
		*(*uint8)(unsafe.Add(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(result)))), a1)) = uint8(v4)
	}
	v7 = a1 + v3
	v30 = a1 + v3
	if a1+v3 >= v17 {
		if v7 < v16 && v5 >= v18 && v5 < v15 {
			*(*uint8)(unsafe.Add(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(v5)))), v30)) = uint8(v4)
			v7 = a1 + v3
		}
		if v7 >= v17 && v7 < v16 && v5 >= v18 && v5 < v15 {
			*(*uint8)(unsafe.Add(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(v5)))), v30)) = uint8(v4)
		}
	}
	v8 = v5 - v3
	if a1 >= v17 {
		if a1 < v16 && v8 >= v18 && v8 < v15 {
			*(*uint8)(unsafe.Add(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(v8)))), a1)) = uint8(v4)
		}
		if a1 >= v17 && a1 < v16 && v8 >= v18 && v8 < v15 {
			*(*uint8)(unsafe.Add(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(v8)))), a1)) = uint8(v4)
		}
	}
	v9 = a1 - v3
	if a1-v3 >= v17 {
		if v9 < v16 && v5 >= v18 && v5 < v15 {
			*(*uint8)(unsafe.Add(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(v5)))), v9)) = uint8(v31)
			result = v24
		}
		if v9 >= v17 && v9 < v16 && v5 >= v18 && v5 < v15 {
			*(*uint8)(unsafe.Add(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(v5)))), v9)) = uint8(v31)
			result = v24
		}
	}
	if a1 >= v17 && a1 < v16 && result >= v18 && result < v15 {
		*(*uint8)(unsafe.Add(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(result)))), a1)) = uint8(v31)
		result = v24
	}
	if v3 > 0 {
		v10 = a1
		v23 = v5 * 4
		v22 = v5 * 4
		v25 = result * 4
		v11 = v8 * 4
		v12 = v18
		result = a1
		v13 = v5 - a1
		v14 = v20
		v29 = v11
		for {
			if v14 >= 0 {
				v21 = v19 + v14
				v19 += 4
				v27--
				v30--
				v9++
				v29 += 4
				v25 -= 4
			} else {
				v21 = v26 + v14
				v19 += 2
			}
			v26 += 2
			v10++
			v28++
			v23 += 4
			result--
			v22 -= 4
			if v10 >= v17 && v10 < v16 && v13+v30 >= v12 && v13+v30 < v15 {
				*(*uint8)(unsafe.Add(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(v25/4)))), v10)) = uint8(v31)
			}
			if v30 >= v17 {
				if v30 < v16 && v13+v10 >= v12 && v13+v10 < v15 {
					*(*uint8)(unsafe.Add(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(v23/4)))), v30)) = uint8(v31)
					v12 = v18
				}
				if v30 < v16 && v13+result >= v12 && v13+result < v15 {
					*(*uint8)(unsafe.Add(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(v22/4)))), v30)) = uint8(v31)
					v12 = v18
				}
			}
			if v10 >= v17 && v10 < v16 && v13+v9 >= v12 && v13+v9 < v15 {
				*(*uint8)(unsafe.Add(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(v29/4)))), v10)) = uint8(v31)
			}
			if result >= v17 && result < v16 && v13+v9 >= v12 && v13+v9 < v15 {
				*(*uint8)(unsafe.Add(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(v29/4)))), result)) = uint8(v31)
			}
			if v9 >= v17 && v9 < v16 {
				if v13+result >= v12 && v13+result < v15 {
					*(*uint8)(unsafe.Add(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(v22/4)))), v9)) = uint8(v31)
				}
				if v9 < v16 && v13+v10 >= v12 && v13+v10 < v15 {
					*(*uint8)(unsafe.Add(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(v23/4)))), v9)) = uint8(v31)
				}
			}
			if result >= v17 && result < v16 && v13+v30 >= v12 && v13+v30 < v15 {
				*(*uint8)(unsafe.Add(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(v25/4)))), result)) = uint8(v31)
			}
			if v27 <= v28 {
				break
			}
			v14 = v21
		}
	}
	return result
}
func nox_video_drawCircle16Opaque_4B1380(a1 int32, a2 int32, a3 int32) int16 {
	var (
		v3        int32
		v4        int32
		result    int16
		v6        int32
		v7        int32
		v8        int32
		v9        int32
		v10       int32
		v11       int32
		v12       int32
		v13       int32
		v14       int32
		v15       int32
		v16       int32
		v17       int32
		v18       int32
		v19       int32
		v20       int32
		v21       int32
		v22       int32
		v23       int32
		v24       int32
		v25       int32
		pixbuffer **uint8 = nox_pixbuffer_rows_3798784
	)
	v3 = a2
	v4 = a3
	if nox_draw_curDrawData_3799572.flag_0 != 0 && sub_49F8E0(a1, a2, a3) {
		return int16(sub_4B15E0(a1, a2, a3))
	}
	v15 = 1 - a3
	v25 = 5 - a3*2
	result = int16(*(*uint16)(unsafe.Pointer(&nox_draw_curDrawData_3799572.field_61)))
	v6 = a1 * 2
	v7 = (a2 + v4) * 4
	v18 = v6
	v20 = 0
	v21 = v4
	v16 = 3
	*(*uint16)(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(v7/4)))))) + uint32(v6)))) = uint16(result)
	v23 = a2 * 4
	v8 = (v4 + a1) * 2
	v17 = v8
	*(*uint16)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(uint32(v3*4) + uint32(uintptr(unsafe.Pointer(pixbuffer)))))) + uint32(v8)))) = uint16(result)
	*(*uint16)(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(v3)))))) + uint32(v8)))) = uint16(result)
	v9 = (v3 - v4) * 4
	*(*uint16)(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(v9/4)))))) + uint32(v18)))) = uint16(result)
	*(*uint16)(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(v9/4)))))) + uint32(a1*2)))) = uint16(result)
	v10 = (a1 - v4) * 2
	*(*uint16)(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(v3)))))) + uint32(v10)))) = uint16(result)
	*(*uint16)(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(v3)))))) + uint32(v10)))) = uint16(result)
	*(*uint16)(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(v7/4)))))) + uint32(a1*2)))) = uint16(result)
	if v4 > 0 {
		v11 = v3 * 4
		v12 = v18
		v22 = v7
		v13 = v18
		v19 = v10
		v14 = v23
		v24 = v9
		for {
			if v15 >= 0 {
				v15 += v25
				v22 -= 4
				v24 += 4
				v25 += 4
				v21--
				v17 -= 2
				v19 += 2
			} else {
				v15 += v16
				v25 += 2
			}
			v12 += 2
			v13 -= 2
			v16 += 2
			v11 += 4
			v14 -= 4
			v20++
			*(*uint16)(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(v22/4)))))) + uint32(v12)))) = uint16(result)
			*(*uint16)(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(v11/4)))))) + uint32(v17)))) = uint16(result)
			*(*uint16)(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(v14/4)))))) + uint32(v17)))) = uint16(result)
			*(*uint16)(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(v24/4)))))) + uint32(v12)))) = uint16(result)
			*(*uint16)(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(v24/4)))))) + uint32(v13)))) = uint16(result)
			*(*uint16)(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(v14/4)))))) + uint32(v19)))) = uint16(result)
			*(*uint16)(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(v11/4)))))) + uint32(v19)))) = uint16(result)
			*(*uint16)(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(v22/4)))))) + uint32(v13)))) = uint16(result)
			if v21 <= v20 {
				break
			}
		}
	}
	return result
}
func sub_4B15E0(a1 int32, a2 int32, a3 int32) int32 {
	var (
		v3        int32
		v4        int16
		v5        int32
		v6        int32
		result    int32
		v8        int32
		v9        int32
		v10       int32
		v11       int32
		v12       int32
		v13       int32
		v14       int32
		v15       int32
		v16       int32
		v17       int32
		v18       int32
		v19       int32
		v20       int32
		v21       int32
		v22       int32
		v23       int32
		v24       int16
		v25       int32
		v26       int32
		v27       int32
		v28       int32
		v29       int32
		v30       int32
		v31       int32
		v32       int32
		v33       int32
		v34       int32
		v35       int32
		v36       int32
		v37       int32
		v38       int32
		pixbuffer **uint8 = nox_pixbuffer_rows_3798784
	)
	v3 = a3
	v32 = 0
	v21 = 1 - a3
	v31 = a3
	v22 = 5 - a3*2
	v30 = 3
	v4 = int16(*(*uint16)(unsafe.Pointer(&nox_draw_curDrawData_3799572.field_61)))
	v5 = nox_draw_curDrawData_3799572.clip.left
	v38 = nox_draw_curDrawData_3799572.clip.top
	v6 = a2
	v18 = nox_draw_curDrawData_3799572.clip.bottom
	result = a1
	v24 = int16(*(*uint16)(unsafe.Pointer(&nox_draw_curDrawData_3799572.field_61)))
	v8 = a2 + v31
	v20 = nox_draw_curDrawData_3799572.clip.left
	v19 = nox_draw_curDrawData_3799572.clip.right
	v36 = a2 + v31
	if a1 >= v5 && a1 < nox_draw_curDrawData_3799572.clip.right && v8 >= v38 && v8 < v18 {
		*(*uint16)(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(v8)))))) + uint32(a1*2)))) = uint16(v4)
	}
	v9 = a1 + v3
	v33 = a1 + v3
	if a1+v3 >= v5 {
		if v9 < v19 && v6 >= v38 && v6 < v18 {
			*(*uint16)(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(v6)))))) + uint32(v9*2)))) = uint16(v4)
		}
		if v9 >= v20 && v9 < v19 && v6 >= v38 && v6 < v18 {
			*(*uint16)(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(v6)))))) + uint32(v9*2)))) = uint16(v4)
		}
	}
	v10 = v6 - v3
	if a1 >= v20 {
		if a1 < v19 && v10 >= v38 && v10 < v18 {
			*(*uint16)(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(v10)))))) + uint32(a1*2)))) = uint16(v4)
		}
		if a1 >= v20 && a1 < v19 && v10 >= v38 && v10 < v18 {
			*(*uint16)(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(v10)))))) + uint32(a1*2)))) = uint16(v4)
		}
	}
	v11 = a1 - v3
	if a1-v3 >= v20 {
		if v11 < v19 && v6 >= v38 && v6 < v18 {
			*(*uint16)(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(v6)))))) + uint32(v11*2)))) = uint16(v4)
			v8 = v36
		}
		if v11 >= v20 && v11 < v19 && v6 >= v38 && v6 < v18 {
			*(*uint16)(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(v6)))))) + uint32(v11*2)))) = uint16(v4)
			v8 = v36
		}
	}
	if a1 >= v20 && a1 < v19 && v8 >= v38 && v8 < v18 {
		*(*uint16)(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(v8)))))) + uint32(a1*2)))) = uint16(v4)
		v8 = v36
	}
	if v3 > 0 {
		v35 = a1 - v3
		v28 = v11 * 2
		v27 = v6 * 4
		v23 = v8
		v25 = v10 * 4
		v37 = result * 2
		v12 = result - v3
		v26 = v6 - result
		v13 = v33 * 2
		v29 = result * 2
		v34 = result - v6
		v14 = v8
		v15 = v6
		for {
			if v21 >= 0 {
				v13 -= 2
				v21 += v22
				v22 += 4
				v31--
				v25 += 4
				v14--
				v28 += 2
				v35 = v12 + 1
				v23 = v14
			} else {
				v21 += v30
				v22 += 2
			}
			v30 += 2
			v32++
			v15++
			v37 += 2
			v27 -= 4
			v29 -= 2
			v16 = v15 + v34
			result--
			if v15+v34 >= v20 && v16 < v19 && v14 >= v38 && v14 < v18 {
				*(*uint16)(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(v14)))))) + uint32(v37)))) = uint16(v24)
				v14 = v23
			}
			v17 = v34 + v14
			if v17 >= v20 {
				if v17 < v19 && v15 >= v38 && v15 < v18 {
					*(*uint16)(unsafe.Pointer(uintptr(uint32(v13) + uint32(uintptr(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(v15))))))))) = uint16(v24)
				}
				if v17 >= v20 && v17 < v19 && result+v26 >= v38 && result+v26 < v18 {
					*(*uint16)(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(v27/4)))))) + uint32(v13)))) = uint16(v24)
				}
			}
			if v16 < v20 || v16 >= v19 {
				v12 = v35
			} else {
				v12 = v35
				if v35+v26 >= v38 && v35+v26 < v18 {
					*(*uint16)(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(v25/4)))))) + uint32(v37)))) = uint16(v24)
				}
			}
			if result >= v20 && result < v19 && v26+v12 >= v38 && v26+v12 < v18 {
				*(*uint16)(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(v25/4)))))) + uint32(v29)))) = uint16(v24)
			}
			if v12 >= v20 && v12 < v19 {
				if result+v26 >= v38 && result+v26 < v18 {
					*(*uint16)(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(v27/4)))))) + uint32(v28)))) = uint16(v24)
				}
				if v12 < v19 && v15 >= v38 && v15 < v18 {
					*(*uint16)(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(v15)))))) + uint32(v28)))) = uint16(v24)
				}
			}
			if result < v20 || result >= v19 {
				v14 = v23
			} else {
				v14 = v23
				if v23 >= v38 && v23 < v18 {
					*(*uint16)(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(v23)))))) + uint32(v29)))) = uint16(v24)
					v12 = v35
				}
			}
			if v31 <= v32 {
				break
			}
		}
	}
	return result
}
func sub_4B1E30(a1 int32, a2 int32, a3 int32) int32 {
	var (
		result    int32
		v4        int32
		v5        int32
		v6        int32
		v7        bool
		v8        uint8
		v9        int32
		v10       int32
		v11       int32
		v12       int32
		v13       int32
		v14       int32
		v15       int32
		v16       int32
		v17       int32
		v18       int32
		v19       int32
		v20       int32
		v21       int32
		v22       int32
		v23       uint8
		v24       int32
		v25       int32
		v26       int32
		v27       int32
		v28       int32
		v29       int32
		v30       int32
		v31       int32
		v32       int32
		v33       int32
		v34       int32
		i         int32
		v36       int32
		pixbuffer **uint8 = nox_pixbuffer_rows_3798784
	)
	result = int32(dword_5d4594_810632)
	if dword_5d4594_810632 != 0 {
		v4 = a3
		v32 = 0
		v22 = 1 - a3
		v25 = 5 - a3*2
		v33 = a3
		v29 = 3
		v5 = nox_draw_curDrawData_3799572.clip.right
		v23 = *(*uint8)(unsafe.Pointer(&nox_draw_curDrawData_3799572.field_61))
		v6 = nox_draw_curDrawData_3799572.clip.left
		v18 = nox_draw_curDrawData_3799572.clip.bottom
		result = a2
		v21 = v6
		v7 = a1 < v6
		v8 = *(*uint8)(unsafe.Pointer(&nox_draw_curDrawData_3799572.field_61))
		v9 = a2 + a3
		v19 = nox_draw_curDrawData_3799572.clip.right
		v20 = nox_draw_curDrawData_3799572.clip.top
		v30 = a2 + a3
		if !v7 && a1 < v5 && v9 >= nox_draw_curDrawData_3799572.clip.top && v9 < v18 {
			v5 = nox_draw_curDrawData_3799572.clip.right
			*(*uint8)(unsafe.Pointer(uintptr(uint32(a1) + *(*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer(pixbuffer))) + uint32((a2+a3)*4))))))) = *(*uint8)(unsafe.Pointer(uintptr(uint32(int32(v23)+(int32(*(*uint8)(unsafe.Pointer(uintptr(uint32(a1) + *(*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer(pixbuffer))) + uint32((a2+a3)*4))))))))<<8)) + dword_5d4594_810632)))
		}
		v10 = a1 + a3
		if a1+a3 >= v21 {
			if v10 < v5 && a2 >= v20 && a2 < v18 {
				v5 = v19
				*(*uint8)(unsafe.Add(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(a2)))), a1+a3)) = *(*uint8)(unsafe.Pointer(uintptr(uint32(int32(v23)+(int32(*(*uint8)(unsafe.Pointer(uintptr(uint32(a1+a3) + uint32(uintptr(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(a2))))))))))<<8)) + dword_5d4594_810632)))
				v10 = a1 + a3
			}
			if v10 >= v21 && v10 < v5 && a2 >= v20 && a2 < v18 {
				v5 = v19
				*(*uint8)(unsafe.Add(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(a2)))), v10)) = *(*uint8)(unsafe.Pointer(uintptr(uint32(int32(v23)+(int32(*(*uint8)(unsafe.Pointer(uintptr(uint32(v10) + uint32(uintptr(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(a2))))))))))<<8)) + dword_5d4594_810632)))
			}
		}
		v11 = a2 - a3
		if a1 >= v21 {
			if a1 < v5 && v11 >= v20 && v11 < v18 {
				v5 = v19
				*(*uint8)(unsafe.Pointer(uintptr(uint32(a1) + *(*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer(pixbuffer))) + uint32((a2-a3)*4))))))) = *(*uint8)(unsafe.Pointer(uintptr(uint32(int32(v23)+(int32(*(*uint8)(unsafe.Pointer(uintptr(uint32(a1) + *(*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer(pixbuffer))) + uint32((a2-a3)*4))))))))<<8)) + dword_5d4594_810632)))
				v11 = a2 - a3
			}
			if a1 >= v21 && a1 < v5 && v11 >= v20 && v11 < v18 {
				*(*uint8)(unsafe.Add(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(v11)))), a1)) = *(*uint8)(unsafe.Pointer(uintptr(uint32(int32(v23)+(int32(*(*uint8)(unsafe.Pointer(uintptr(uint32(a1) + uint32(uintptr(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(v11))))))))))<<8)) + dword_5d4594_810632)))
				v11 = a2 - a3
			}
		}
		v12 = a1 - a3
		if a1-a3 >= v21 {
			if v12 < v19 && a2 >= v20 && a2 < v18 {
				*(*uint8)(unsafe.Add(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(a2)))), v12)) = *(*uint8)(unsafe.Pointer(uintptr(uint32(int32(v23)+(int32(*(*uint8)(unsafe.Pointer(uintptr(uint32(v12) + uint32(uintptr(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(a2))))))))))<<8)) + dword_5d4594_810632)))
				v8 = v23
				v11 = a2 - a3
			}
			if v12 >= v21 && v12 < v19 && a2 >= v20 && a2 < v18 {
				*(*uint8)(unsafe.Add(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(a2)))), v12)) = *(*uint8)(unsafe.Pointer(uintptr(uint32(int32(v8)+(int32(*(*uint8)(unsafe.Pointer(uintptr(uint32(v12) + uint32(uintptr(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(a2))))))))))<<8)) + dword_5d4594_810632)))
				v8 = v23
				v11 = a2 - a3
			}
		}
		if a1 >= v21 && a1 < v19 {
			if v30 >= v20 && v30 < v18 {
				*(*uint8)(unsafe.Add(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(v30)))), a1)) = *(*uint8)(unsafe.Pointer(uintptr(uint32(int32(v8)+(int32(*(*uint8)(unsafe.Pointer(uintptr(uint32(a1) + uint32(uintptr(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(v30))))))))))<<8)) + dword_5d4594_810632)))
				v8 = v23
			}
			v11 = a2 - a3
		}
		if a3 > 0 {
			v36 = a1 - a3
			v31 = v30 * 4
			v28 = a1 + v4
			v24 = v11 * 4
			v34 = a1 - a2
			v13 = a2 - a1
			result = v36
			v14 = a1
			v27 = a2 * 4
			v26 = a2 * 4
			for i = a2 - a1; ; v13 = i {
				if v22 >= 0 {
					v22 += v25
					v25 += 4
					v33--
					v31 -= 4
					v15 = v28 - 1
					result++
					v28--
					v24 += 4
					v36 = result
				} else {
					v22 += v29
					v15 = v28
					v25 += 2
				}
				v29 += 2
				v14++
				v32++
				v26 -= 4
				a2--
				v27 += 4
				if v14 < v21 || v14 >= v19 {
					goto LABEL_54
				}
				v16 = v13 + v15
				v17 = v20
				if v16 >= v20 && v16 < v18 {
					break
				}
			LABEL_55:
				if v15 >= v21 && v15 < v19 {
					if v14+i >= v17 && v14+i < v18 {
						v17 = v20
						*(*uint8)(unsafe.Add(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(v27/4)))), v15)) = *(*uint8)(unsafe.Pointer(uintptr(uint32(int32(v8)+(int32(*(*uint8)(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(v27/4)))))) + uint32(v15)))))<<8)) + dword_5d4594_810632)))
						result = v36
					}
					if v15 < v19 && a2 >= v17 && a2 < v18 {
						*(*uint8)(unsafe.Add(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(v26/4)))), v15)) = *(*uint8)(unsafe.Pointer(uintptr(uint32(int32(v8)+(int32(*(*uint8)(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(v26/4)))))) + uint32(v15)))))<<8)) + dword_5d4594_810632)))
						result = v36
					}
				}
				if v14 >= v21 && v14 < v19 && result+i >= v17 && result+i < v18 {
					*(*uint8)(unsafe.Add(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(v24/4)))), v14)) = *(*uint8)(unsafe.Pointer(uintptr(uint32(int32(v8)+(int32(*(*uint8)(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(v24/4)))))) + uint32(v14)))))<<8)) + dword_5d4594_810632)))
					result = v36
				}
				if v34+a2 >= v21 && v34+a2 < v19 && result+i >= v17 && result+i < v18 {
					*(*uint8)(unsafe.Add(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(v24/4)))), v34+a2)) = *(*uint8)(unsafe.Pointer(uintptr(uint32(int32(v8)+(int32(*(*uint8)(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(v24/4)))))) + uint32(v34) + uint32(a2)))))<<8)) + dword_5d4594_810632)))
					result = v36
				}
				if result >= v21 && result < v19 {
					if a2 >= v17 && a2 < v18 {
						*(*uint8)(unsafe.Add(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(v26/4)))), result)) = *(*uint8)(unsafe.Pointer(uintptr(uint32(int32(v8)+(int32(*(*uint8)(unsafe.Pointer(uintptr(uint32(result) + uint32(uintptr(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(v26/4))))))))))<<8)) + dword_5d4594_810632)))
						result = v36
					}
					if result < v19 && v14+i >= v17 && v14+i < v18 {
						*(*uint8)(unsafe.Add(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(v27/4)))), result)) = *(*uint8)(unsafe.Pointer(uintptr(uint32(int32(v8)+(int32(*(*uint8)(unsafe.Pointer(uintptr(uint32(result) + uint32(uintptr(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(v27/4))))))))))<<8)) + dword_5d4594_810632)))
					}
				}
				if v33 <= v32 {
					return result
				}
			}
			*(*uint8)(unsafe.Add(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(v31/4)))), v14)) = *(*uint8)(unsafe.Pointer(uintptr(uint32(int32(v8)+(int32(*(*uint8)(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(v31/4)))))) + uint32(v14)))))<<8)) + dword_5d4594_810632)))
			result = v36
		LABEL_54:
			v17 = v20
			goto LABEL_55
		}
	}
	return result
}
func nox_video_drawCircle16Alpha_4B2480(a1 int32, a2 int32, a3 int32) *uint16 {
	var (
		result     *uint16
		colourBase uint32
		v5         int32
		v6         *uint16
		redcol     int32
		bluecol2   int32
		grencol2   int32
		v10        *uint16
		v11        *uint16
		v12        int32
		v13        *uint16
		v14        *uint16
		v15        *uint16
		v16        *uint16
		v17        *uint16
		v19        int32
		v20        *uint16
		v21        *uint16
		v22        *uint16
		v23        *uint16
		v24        *uint16
		v25        *uint16
		v26        *uint16
		v27        *uint16
		v44        int32
		v45        int32
		v46        int32
		v47        int32
		v48        int32
		v49        int32
		v50        int32
		v51        int32
		v52        int32
		v53        int32
		v54        int32
		v55        int32
		v56        int32
		v57        int32
		grencol    uint8
		v59        int32
		bluecol    uint8
		v61        int32
		v62        int32
		v63        int32
		v64        int32
	)
	if nox_draw_curDrawData_3799572.flag_0 != 0 && sub_49F8E0(a1, a2, a3) {
		return (*uint16)(unsafe.Pointer(uintptr(sub_4B3450(a1, a2, a3))))
	}
	var pixbuffer **uint8 = nox_pixbuffer_rows_3798784
	v57 = 0
	v56 = a3
	colourBase = uint32(*(*uint16)(unsafe.Pointer(&nox_draw_curDrawData_3799572.field_61)))
	v55 = 3
	grencol = uint8((*((*uint32)(unsafe.Pointer(&byte_5D4594_3804364[0+4]))) & colourBase) >> uint32(byte_5D4594_3804364[0+16]))
	v44 = 1 - a3
	bluecol = uint8(int8((int32(uint8(*(*uint16)(unsafe.Pointer(&nox_draw_curDrawData_3799572.field_61)))) & int32(byte_5D4594_3804364[0+8])) << int32(byte_5D4594_3804364[0+20])))
	v45 = 5 - a3*2
	v5 = (a3 + a2) * 4
	v53 = v5
	v46 = a1 * 2
	v6 = (*uint16)(unsafe.Pointer(uintptr(uint32(a1*2) + uint32(uintptr(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(v5/4)))))))))
	redcol = int32(uint8((*((*uint32)(unsafe.Pointer(&byte_5D4594_3804364[0+0]))) & colourBase) >> uint32(byte_5D4594_3804364[0+12])))
	bluecol2 = int32(bluecol)
	grencol2 = int32(grencol)
	*v6 = uint16(int16(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(nox_draw_colors_b_3804664), int32(uint8(int8(((int32(uint8(*v6))&int32(byte_5D4594_3804364[0+8]))<<int32(byte_5D4594_3804364[0+20]))+((int32(bluecol)-int32(uint8(int8((int32(uint8(*v6))&int32(byte_5D4594_3804364[0+8]))<<int32(byte_5D4594_3804364[0+20])))))>>2))))*2))))) | int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(nox_draw_colors_g_3804656), int32(uint8(int8((int32(uint16(int16(int32(*((*uint16)(unsafe.Pointer(&byte_5D4594_3804364[0+4]))))&int32(*v6))))>>int32(byte_5D4594_3804364[0+16]))+((int32(grencol)-int32(uint8(int8(int32(uint16(int16(int32(*((*uint16)(unsafe.Pointer(&byte_5D4594_3804364[0+4]))))&int32(*v6))))>>int32(byte_5D4594_3804364[0+16])))))>>2))))*2))))) | int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(nox_draw_colors_r_3804672), int32(uint8(int8((int32(uint16(int16(int32(*((*uint16)(unsafe.Pointer(&byte_5D4594_3804364[0+0]))))&int32(*v6))))>>int32(byte_5D4594_3804364[0+12]))+((redcol-int32(uint8(int8(int32(uint16(int16(int32(*((*uint16)(unsafe.Pointer(&byte_5D4594_3804364[0+0]))))&int32(*v6))))>>int32(byte_5D4594_3804364[0+12])))))>>2))))*2)))))))
	v49 = a2 * 4
	v10 = (*uint16)(unsafe.Pointer(uintptr(uint32((a1+a3)*2) + *(*uint32)(unsafe.Pointer(uintptr(uint32(a2*4) + uint32(uintptr(unsafe.Pointer(pixbuffer)))))))))
	v59 = (a1 + a3) * 2
	*v10 = uint16(int16(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(nox_draw_colors_b_3804664), int32(uint8(int8(((int32(uint8(*v10))&int32(byte_5D4594_3804364[0+8]))<<int32(byte_5D4594_3804364[0+20]))+((int32(bluecol)-int32(uint8(int8((int32(uint8(*v10))&int32(byte_5D4594_3804364[0+8]))<<int32(byte_5D4594_3804364[0+20])))))>>2))))*2))))) | int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(nox_draw_colors_g_3804656), int32(uint8(int8((int32(uint16(int16(int32(*((*uint16)(unsafe.Pointer(&byte_5D4594_3804364[0+4]))))&int32(*v10))))>>int32(byte_5D4594_3804364[0+16]))+((grencol2-int32(uint8(int8(int32(uint16(int16(int32(*((*uint16)(unsafe.Pointer(&byte_5D4594_3804364[0+4]))))&int32(*v10))))>>int32(byte_5D4594_3804364[0+16])))))>>2))))*2))))) | int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(nox_draw_colors_r_3804672), int32(uint8(int8((int32(uint16(int16(int32(*((*uint16)(unsafe.Pointer(&byte_5D4594_3804364[0+0]))))&int32(*v10))))>>int32(byte_5D4594_3804364[0+12]))+((redcol-int32(uint8(int8(int32(uint16(int16(int32(*((*uint16)(unsafe.Pointer(&byte_5D4594_3804364[0+0]))))&int32(*v10))))>>int32(byte_5D4594_3804364[0+12])))))>>2))))*2)))))))
	v11 = (*uint16)(unsafe.Pointer(uintptr(uint32(v59) + uint32(uintptr(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(a2)))))))))
	*v11 = uint16(int16(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(nox_draw_colors_b_3804664), int32(uint8(int8(((int32(uint8(*v11))&int32(byte_5D4594_3804364[0+8]))<<int32(byte_5D4594_3804364[0+20]))+((int32(bluecol)-int32(uint8(int8((int32(uint8(*v11))&int32(byte_5D4594_3804364[0+8]))<<int32(byte_5D4594_3804364[0+20])))))>>2))))*2))))) | int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(nox_draw_colors_g_3804656), int32(uint8(int8((int32(uint16(int16(int32(*((*uint16)(unsafe.Pointer(&byte_5D4594_3804364[0+4]))))&int32(*v11))))>>int32(byte_5D4594_3804364[0+16]))+((grencol2-int32(uint8(int8(int32(uint16(int16(int32(*((*uint16)(unsafe.Pointer(&byte_5D4594_3804364[0+4]))))&int32(*v11))))>>int32(byte_5D4594_3804364[0+16])))))>>2))))*2))))) | int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(nox_draw_colors_r_3804672), int32(uint8(int8((int32(uint16(int16(int32(*((*uint16)(unsafe.Pointer(&byte_5D4594_3804364[0+0]))))&int32(*v11))))>>int32(byte_5D4594_3804364[0+12]))+((redcol-int32(uint8(int8(int32(uint16(int16(int32(*((*uint16)(unsafe.Pointer(&byte_5D4594_3804364[0+0]))))&int32(*v11))))>>int32(byte_5D4594_3804364[0+12])))))>>2))))*2)))))))
	v12 = (a2 - a3) * 4
	v61 = v12
	v13 = (*uint16)(unsafe.Pointer(uintptr(uint32(a1*2) + uint32(uintptr(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(v12/4)))))))))
	*v13 = uint16(int16(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(nox_draw_colors_b_3804664), int32(uint8(int8(((int32(uint8(*v13))&int32(byte_5D4594_3804364[0+8]))<<int32(byte_5D4594_3804364[0+20]))+((bluecol2-int32(uint8(int8((int32(uint8(*v13))&int32(byte_5D4594_3804364[0+8]))<<int32(byte_5D4594_3804364[0+20])))))>>2))))*2))))) | int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(nox_draw_colors_g_3804656), int32(uint8(int8((int32(uint16(int16(int32(*((*uint16)(unsafe.Pointer(&byte_5D4594_3804364[0+4]))))&int32(*v13))))>>int32(byte_5D4594_3804364[0+16]))+((grencol2-int32(uint8(int8(int32(uint16(int16(int32(*((*uint16)(unsafe.Pointer(&byte_5D4594_3804364[0+4]))))&int32(*v13))))>>int32(byte_5D4594_3804364[0+16])))))>>2))))*2))))) | int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(nox_draw_colors_r_3804672), int32(uint8(int8((int32(uint16(int16(int32(*((*uint16)(unsafe.Pointer(&byte_5D4594_3804364[0+0]))))&int32(*v13))))>>int32(byte_5D4594_3804364[0+12]))+((redcol-int32(uint8(int8(int32(uint16(int16(int32(*((*uint16)(unsafe.Pointer(&byte_5D4594_3804364[0+0]))))&int32(*v13))))>>int32(byte_5D4594_3804364[0+12])))))>>2))))*2)))))))
	v51 = a1 * 2
	v14 = (*uint16)(unsafe.Pointer(uintptr(uint32(a1*2) + uint32(uintptr(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(v61/4)))))))))
	*v14 = uint16(int16(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(nox_draw_colors_b_3804664), int32(uint8(int8(((int32(uint8(*v14))&int32(byte_5D4594_3804364[0+8]))<<int32(byte_5D4594_3804364[0+20]))+((bluecol2-int32(uint8(int8((int32(uint8(*v14))&int32(byte_5D4594_3804364[0+8]))<<int32(byte_5D4594_3804364[0+20])))))>>2))))*2))))) | int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(nox_draw_colors_g_3804656), int32(uint8(int8((int32(uint16(int16(int32(*((*uint16)(unsafe.Pointer(&byte_5D4594_3804364[0+4]))))&int32(*v14))))>>int32(byte_5D4594_3804364[0+16]))+((grencol2-int32(uint8(int8(int32(uint16(int16(int32(*((*uint16)(unsafe.Pointer(&byte_5D4594_3804364[0+4]))))&int32(*v14))))>>int32(byte_5D4594_3804364[0+16])))))>>2))))*2))))) | int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(nox_draw_colors_r_3804672), int32(uint8(int8((int32(uint16(int16(int32(*((*uint16)(unsafe.Pointer(&byte_5D4594_3804364[0+0]))))&int32(*v14))))>>int32(byte_5D4594_3804364[0+12]))+((redcol-int32(uint8(int8(int32(uint16(int16(int32(*((*uint16)(unsafe.Pointer(&byte_5D4594_3804364[0+0]))))&int32(*v14))))>>int32(byte_5D4594_3804364[0+12])))))>>2))))*2)))))))
	v62 = (a1 - a3) * 2
	v15 = (*uint16)(unsafe.Pointer(uintptr(uint32(v62) + uint32(uintptr(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(a2)))))))))
	*v15 = uint16(int16(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(nox_draw_colors_b_3804664), int32(uint8(int8(((int32(uint8(*v15))&int32(byte_5D4594_3804364[0+8]))<<int32(byte_5D4594_3804364[0+20]))+((bluecol2-int32(uint8(int8((int32(uint8(*v15))&int32(byte_5D4594_3804364[0+8]))<<int32(byte_5D4594_3804364[0+20])))))>>2))))*2))))) | int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(nox_draw_colors_g_3804656), int32(uint8(int8((int32(uint16(int16(int32(*((*uint16)(unsafe.Pointer(&byte_5D4594_3804364[0+4]))))&int32(*v15))))>>int32(byte_5D4594_3804364[0+16]))+((grencol2-int32(uint8(int8(int32(uint16(int16(int32(*((*uint16)(unsafe.Pointer(&byte_5D4594_3804364[0+4]))))&int32(*v15))))>>int32(byte_5D4594_3804364[0+16])))))>>2))))*2))))) | int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(nox_draw_colors_r_3804672), int32(uint8(int8((int32(uint16(int16(int32(*((*uint16)(unsafe.Pointer(&byte_5D4594_3804364[0+0]))))&int32(*v15))))>>int32(byte_5D4594_3804364[0+12]))+((redcol-int32(uint8(int8(int32(uint16(int16(int32(*((*uint16)(unsafe.Pointer(&byte_5D4594_3804364[0+0]))))&int32(*v15))))>>int32(byte_5D4594_3804364[0+12])))))>>2))))*2)))))))
	v16 = (*uint16)(unsafe.Pointer(uintptr(uint32(v62) + *(*uint32)(unsafe.Pointer(uintptr(uint32(a2*4) + uint32(uintptr(unsafe.Pointer(pixbuffer)))))))))
	*v16 = uint16(int16(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(nox_draw_colors_b_3804664), int32(uint8(int8(((int32(uint8(*v16))&int32(byte_5D4594_3804364[0+8]))<<int32(byte_5D4594_3804364[0+20]))+((bluecol2-int32(uint8(int8((int32(uint8(*v16))&int32(byte_5D4594_3804364[0+8]))<<int32(byte_5D4594_3804364[0+20])))))>>2))))*2))))) | int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(nox_draw_colors_g_3804656), int32(uint8(int8((int32(uint16(int16(int32(*((*uint16)(unsafe.Pointer(&byte_5D4594_3804364[0+4]))))&int32(*v16))))>>int32(byte_5D4594_3804364[0+16]))+((grencol2-int32(uint8(int8(int32(uint16(int16(int32(*((*uint16)(unsafe.Pointer(&byte_5D4594_3804364[0+4]))))&int32(*v16))))>>int32(byte_5D4594_3804364[0+16])))))>>2))))*2))))) | int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(nox_draw_colors_r_3804672), int32(uint8(int8((int32(uint16(int16(int32(*((*uint16)(unsafe.Pointer(&byte_5D4594_3804364[0+0]))))&int32(*v16))))>>int32(byte_5D4594_3804364[0+12]))+((redcol-int32(uint8(int8(int32(uint16(int16(int32(*((*uint16)(unsafe.Pointer(&byte_5D4594_3804364[0+0]))))&int32(*v16))))>>int32(byte_5D4594_3804364[0+12])))))>>2))))*2)))))))
	v17 = (*uint16)(unsafe.Pointer(uintptr(uint32(v51) + uint32(uintptr(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(v53/4)))))))))
	*v17 = uint16(int16(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(nox_draw_colors_b_3804664), int32(uint8(int8(((int32(uint8(*v17))&int32(byte_5D4594_3804364[0+8]))<<int32(byte_5D4594_3804364[0+20]))+((bluecol2-int32(uint8(int8((int32(uint8(*v17))&int32(byte_5D4594_3804364[0+8]))<<int32(byte_5D4594_3804364[0+20])))))>>2))))*2))))) | int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(nox_draw_colors_g_3804656), int32(uint8(int8((int32(uint16(int16(int32(*((*uint16)(unsafe.Pointer(&byte_5D4594_3804364[0+4]))))&int32(*v17))))>>int32(byte_5D4594_3804364[0+16]))+((grencol2-int32(uint8(int8(int32(uint16(int16(int32(*((*uint16)(unsafe.Pointer(&byte_5D4594_3804364[0+4]))))&int32(*v17))))>>int32(byte_5D4594_3804364[0+16])))))>>2))))*2))))) | int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(nox_draw_colors_r_3804672), int32(uint8(int8((int32(uint16(int16(int32(*((*uint16)(unsafe.Pointer(&byte_5D4594_3804364[0+0]))))&int32(*v17))))>>int32(byte_5D4594_3804364[0+12]))+((redcol-int32(uint8(int8(int32(uint16(int16(int32(*((*uint16)(unsafe.Pointer(&byte_5D4594_3804364[0+0]))))&int32(*v17))))>>int32(byte_5D4594_3804364[0+12])))))>>2))))*2)))))))
	result = v17
	if a3 > 0 {
		v52 = v53
		v19 = v46
		v64 = v59
		v48 = v49
		v47 = v49
		v50 = v62
		v54 = v19
		v63 = v61
		for {
			if v44 >= 0 {
				v44 += v45
				v45 += 4
				v56 -= 1
				v64 -= 2
				v52 -= 4
				v63 += 4
				v50 += 2
			} else {
				v44 += v55
				v45 += 2
			}
			v55 += 2
			v57++
			v48 += 4
			v47 -= 4
			v54 -= 2
			v20 = (*uint16)(unsafe.Pointer(uintptr(uint32(v19+2) + uint32(uintptr(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(v52/4)))))))))
			*v20 = uint16(int16(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(nox_draw_colors_b_3804664), int32(uint8(int8(((int32(uint8(*v20))&int32(byte_5D4594_3804364[0+8]))<<int32(byte_5D4594_3804364[0+20]))+((bluecol2-int32(uint8(int8((int32(uint8(*v20))&int32(byte_5D4594_3804364[0+8]))<<int32(byte_5D4594_3804364[0+20])))))>>2))))*2))))) | int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(nox_draw_colors_g_3804656), int32(uint8(int8((int32(uint16(int16(int32(*((*uint16)(unsafe.Pointer(&byte_5D4594_3804364[0+4]))))&int32(*v20))))>>int32(byte_5D4594_3804364[0+16]))+((grencol2-int32(uint8(int8(int32(uint16(int16(int32(*((*uint16)(unsafe.Pointer(&byte_5D4594_3804364[0+4]))))&int32(*v20))))>>int32(byte_5D4594_3804364[0+16])))))>>2))))*2))))) | int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(nox_draw_colors_r_3804672), int32(uint8(int8((int32(uint16(int16(int32(*((*uint16)(unsafe.Pointer(&byte_5D4594_3804364[0+0]))))&int32(*v20))))>>int32(byte_5D4594_3804364[0+12]))+((redcol-int32(uint8(int8(int32(uint16(int16(int32(*((*uint16)(unsafe.Pointer(&byte_5D4594_3804364[0+0]))))&int32(*v20))))>>int32(byte_5D4594_3804364[0+12])))))>>2))))*2)))))))
			v21 = (*uint16)(unsafe.Pointer(uintptr(uint32(v64) + uint32(uintptr(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(v48/4)))))))))
			*v21 = uint16(int16(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(nox_draw_colors_b_3804664), int32(uint8(int8(((int32(uint8(*v21))&int32(byte_5D4594_3804364[0+8]))<<int32(byte_5D4594_3804364[0+20]))+((bluecol2-int32(uint8(int8((int32(uint8(*v21))&int32(byte_5D4594_3804364[0+8]))<<int32(byte_5D4594_3804364[0+20])))))>>2))))*2))))) | int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(nox_draw_colors_g_3804656), int32(uint8(int8((int32(uint16(int16(int32(*((*uint16)(unsafe.Pointer(&byte_5D4594_3804364[0+4]))))&int32(*v21))))>>int32(byte_5D4594_3804364[0+16]))+((grencol2-int32(uint8(int8(int32(uint16(int16(int32(*((*uint16)(unsafe.Pointer(&byte_5D4594_3804364[0+4]))))&int32(*v21))))>>int32(byte_5D4594_3804364[0+16])))))>>2))))*2))))) | int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(nox_draw_colors_r_3804672), int32(uint8(int8((int32(uint16(int16(int32(*((*uint16)(unsafe.Pointer(&byte_5D4594_3804364[0+0]))))&int32(*v21))))>>int32(byte_5D4594_3804364[0+12]))+((redcol-int32(uint8(int8(int32(uint16(int16(int32(*((*uint16)(unsafe.Pointer(&byte_5D4594_3804364[0+0]))))&int32(*v21))))>>int32(byte_5D4594_3804364[0+12])))))>>2))))*2)))))))
			v22 = (*uint16)(unsafe.Pointer(uintptr(uint32(v64) + uint32(uintptr(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(v47/4)))))))))
			*v22 = uint16(int16(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(nox_draw_colors_b_3804664), int32(uint8(int8(((int32(uint8(*v22))&int32(byte_5D4594_3804364[0+8]))<<int32(byte_5D4594_3804364[0+20]))+((bluecol2-int32(uint8(int8((int32(uint8(*v22))&int32(byte_5D4594_3804364[0+8]))<<int32(byte_5D4594_3804364[0+20])))))>>2))))*2))))) | int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(nox_draw_colors_g_3804656), int32(uint8(int8((int32(uint16(int16(int32(*((*uint16)(unsafe.Pointer(&byte_5D4594_3804364[0+4]))))&int32(*v22))))>>int32(byte_5D4594_3804364[0+16]))+((grencol2-int32(uint8(int8(int32(uint16(int16(int32(*((*uint16)(unsafe.Pointer(&byte_5D4594_3804364[0+4]))))&int32(*v22))))>>int32(byte_5D4594_3804364[0+16])))))>>2))))*2))))) | int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(nox_draw_colors_r_3804672), int32(uint8(int8((int32(uint16(int16(int32(*((*uint16)(unsafe.Pointer(&byte_5D4594_3804364[0+0]))))&int32(*v22))))>>int32(byte_5D4594_3804364[0+12]))+((redcol-int32(uint8(int8(int32(uint16(int16(int32(*((*uint16)(unsafe.Pointer(&byte_5D4594_3804364[0+0]))))&int32(*v22))))>>int32(byte_5D4594_3804364[0+12])))))>>2))))*2)))))))
			v23 = (*uint16)(unsafe.Pointer(uintptr(uint32(v19+2) + uint32(uintptr(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(v63/4)))))))))
			*v23 = uint16(int16(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(nox_draw_colors_b_3804664), int32(uint8(int8(((int32(uint8(*v23))&int32(byte_5D4594_3804364[0+8]))<<int32(byte_5D4594_3804364[0+20]))+((bluecol2-int32(uint8(int8((int32(uint8(*v23))&int32(byte_5D4594_3804364[0+8]))<<int32(byte_5D4594_3804364[0+20])))))>>2))))*2))))) | int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(nox_draw_colors_g_3804656), int32(uint8(int8((int32(uint16(int16(int32(*((*uint16)(unsafe.Pointer(&byte_5D4594_3804364[0+4]))))&int32(*v23))))>>int32(byte_5D4594_3804364[0+16]))+((grencol2-int32(uint8(int8(int32(uint16(int16(int32(*((*uint16)(unsafe.Pointer(&byte_5D4594_3804364[0+4]))))&int32(*v23))))>>int32(byte_5D4594_3804364[0+16])))))>>2))))*2))))) | int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(nox_draw_colors_r_3804672), int32(uint8(int8((int32(uint16(int16(int32(*((*uint16)(unsafe.Pointer(&byte_5D4594_3804364[0+0]))))&int32(*v23))))>>int32(byte_5D4594_3804364[0+12]))+((redcol-int32(uint8(int8(int32(uint16(int16(int32(*((*uint16)(unsafe.Pointer(&byte_5D4594_3804364[0+0]))))&int32(*v23))))>>int32(byte_5D4594_3804364[0+12])))))>>2))))*2)))))))
			v24 = (*uint16)(unsafe.Pointer(uintptr(uint32(v54) + uint32(uintptr(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(v63/4)))))))))
			*v24 = uint16(int16(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(nox_draw_colors_b_3804664), int32(uint8(int8(((int32(uint8(*v24))&int32(byte_5D4594_3804364[0+8]))<<int32(byte_5D4594_3804364[0+20]))+((bluecol2-int32(uint8(int8((int32(uint8(*v24))&int32(byte_5D4594_3804364[0+8]))<<int32(byte_5D4594_3804364[0+20])))))>>2))))*2))))) | int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(nox_draw_colors_g_3804656), int32(uint8(int8((int32(uint16(int16(int32(*((*uint16)(unsafe.Pointer(&byte_5D4594_3804364[0+4]))))&int32(*v24))))>>int32(byte_5D4594_3804364[0+16]))+((grencol2-int32(uint8(int8(int32(uint16(int16(int32(*((*uint16)(unsafe.Pointer(&byte_5D4594_3804364[0+4]))))&int32(*v24))))>>int32(byte_5D4594_3804364[0+16])))))>>2))))*2))))) | int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(nox_draw_colors_r_3804672), int32(uint8(int8((int32(uint16(int16(int32(*((*uint16)(unsafe.Pointer(&byte_5D4594_3804364[0+0]))))&int32(*v24))))>>int32(byte_5D4594_3804364[0+12]))+((redcol-int32(uint8(int8(int32(uint16(int16(int32(*((*uint16)(unsafe.Pointer(&byte_5D4594_3804364[0+0]))))&int32(*v24))))>>int32(byte_5D4594_3804364[0+12])))))>>2))))*2)))))))
			v25 = (*uint16)(unsafe.Pointer(uintptr(uint32(v50) + uint32(uintptr(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(v47/4)))))))))
			*v25 = uint16(int16(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(nox_draw_colors_b_3804664), int32(uint8(int8(((int32(uint8(*v25))&int32(byte_5D4594_3804364[0+8]))<<int32(byte_5D4594_3804364[0+20]))+((bluecol2-int32(uint8(int8((int32(uint8(*v25))&int32(byte_5D4594_3804364[0+8]))<<int32(byte_5D4594_3804364[0+20])))))>>2))))*2))))) | int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(nox_draw_colors_g_3804656), int32(uint8(int8((int32(uint16(int16(int32(*((*uint16)(unsafe.Pointer(&byte_5D4594_3804364[0+4]))))&int32(*v25))))>>int32(byte_5D4594_3804364[0+16]))+((grencol2-int32(uint8(int8(int32(uint16(int16(int32(*((*uint16)(unsafe.Pointer(&byte_5D4594_3804364[0+4]))))&int32(*v25))))>>int32(byte_5D4594_3804364[0+16])))))>>2))))*2))))) | int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(nox_draw_colors_r_3804672), int32(uint8(int8((int32(uint16(int16(int32(*((*uint16)(unsafe.Pointer(&byte_5D4594_3804364[0+0]))))&int32(*v25))))>>int32(byte_5D4594_3804364[0+12]))+((redcol-int32(uint8(int8(int32(uint16(int16(int32(*((*uint16)(unsafe.Pointer(&byte_5D4594_3804364[0+0]))))&int32(*v25))))>>int32(byte_5D4594_3804364[0+12])))))>>2))))*2)))))))
			v26 = (*uint16)(unsafe.Pointer(uintptr(uint32(v50) + uint32(uintptr(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(v48/4)))))))))
			*v26 = uint16(int16(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(nox_draw_colors_b_3804664), int32(uint8(int8(((int32(uint8(*v26))&int32(byte_5D4594_3804364[0+8]))<<int32(byte_5D4594_3804364[0+20]))+((bluecol2-int32(uint8(int8((int32(uint8(*v26))&int32(byte_5D4594_3804364[0+8]))<<int32(byte_5D4594_3804364[0+20])))))>>2))))*2))))) | int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(nox_draw_colors_g_3804656), int32(uint8(int8((int32(uint16(int16(int32(*((*uint16)(unsafe.Pointer(&byte_5D4594_3804364[0+4]))))&int32(*v26))))>>int32(byte_5D4594_3804364[0+16]))+((grencol2-int32(uint8(int8(int32(uint16(int16(int32(*((*uint16)(unsafe.Pointer(&byte_5D4594_3804364[0+4]))))&int32(*v26))))>>int32(byte_5D4594_3804364[0+16])))))>>2))))*2))))) | int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(nox_draw_colors_r_3804672), int32(uint8(int8((int32(uint16(int16(int32(*((*uint16)(unsafe.Pointer(&byte_5D4594_3804364[0+0]))))&int32(*v26))))>>int32(byte_5D4594_3804364[0+12]))+((redcol-int32(uint8(int8(int32(uint16(int16(int32(*((*uint16)(unsafe.Pointer(&byte_5D4594_3804364[0+0]))))&int32(*v26))))>>int32(byte_5D4594_3804364[0+12])))))>>2))))*2)))))))
			v27 = (*uint16)(unsafe.Pointer(uintptr(uint32(v54) + uint32(uintptr(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(v52/4)))))))))
			*v27 = uint16(int16(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(nox_draw_colors_b_3804664), int32(uint8(int8(((int32(uint8(*v27))&int32(byte_5D4594_3804364[0+8]))<<int32(byte_5D4594_3804364[0+20]))+((bluecol2-int32(uint8(int8((int32(uint8(*v27))&int32(byte_5D4594_3804364[0+8]))<<int32(byte_5D4594_3804364[0+20])))))>>2))))*2))))) | int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(nox_draw_colors_g_3804656), int32(uint8(int8((int32(uint16(int16(int32(*((*uint16)(unsafe.Pointer(&byte_5D4594_3804364[0+4]))))&int32(*v27))))>>int32(byte_5D4594_3804364[0+16]))+((grencol2-int32(uint8(int8(int32(uint16(int16(int32(*((*uint16)(unsafe.Pointer(&byte_5D4594_3804364[0+4]))))&int32(*v27))))>>int32(byte_5D4594_3804364[0+16])))))>>2))))*2))))) | int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(nox_draw_colors_r_3804672), int32(uint8(int8((int32(uint16(int16(int32(*((*uint16)(unsafe.Pointer(&byte_5D4594_3804364[0+0]))))&int32(*v27))))>>int32(byte_5D4594_3804364[0+12]))+((redcol-int32(uint8(int8(int32(uint16(int16(int32(*((*uint16)(unsafe.Pointer(&byte_5D4594_3804364[0+0]))))&int32(*v27))))>>int32(byte_5D4594_3804364[0+12])))))>>2))))*2)))))))
			result = (*uint16)(unsafe.Pointer(uintptr(v56)))
			if v56 <= v57 {
				break
			}
			v19 += 2
		}
	}
	return result
}
func sub_4B3450(a1 int32, a2 int32, a3 int32) int32 {
	var (
		v3        int16
		v4        uint8
		v5        uint8
		v6        uint8
		v7        int8
		v8        uint8
		v9        uint8
		v10       int32
		v11       int32
		v12       int32
		v13       int32
		result    int32
		v15       int32
		v16       int32
		v17       int32
		v18       int32
		v19       int32
		v20       int32
		v21       int32
		v22       *uint16
		v23       uint32
		v24       int32
		v25       *uint16
		v26       uint32
		v27       *uint16
		v28       uint32
		v29       *uint16
		v30       uint32
		v31       *uint16
		v32       uint32
		v33       *uint16
		v34       uint32
		v35       *uint16
		v36       uint32
		v37       *uint16
		v38       uint32
		v39       *uint16
		v40       *uint16
		v41       *uint16
		v42       *uint16
		v43       *uint16
		v44       *uint16
		v45       *uint16
		v46       *uint16
		v47       int32
		v48       int32
		v49       int32
		v50       int32
		v51       uint32
		v52       uint32
		v53       uint32
		v54       uint32
		v55       uint32
		v56       uint32
		v57       uint32
		v58       uint32
		v59       int32
		v60       int32
		v61       int32
		v62       int32
		v63       int32
		v64       int32
		v65       int32
		v66       int32
		v67       int32
		v68       int32
		v69       int32
		v70       int32
		v71       uint8
		v72       int32
		v73       int32
		v74       int32
		v75       int32
		v76       uint8
		v77       int32
		v78       int32
		v79       int32
		v80       uint8
		i         int32
		v82       int32
		v83       int32
		pixbuffer **uint8 = nox_pixbuffer_rows_3798784
	)
	v79 = a3
	v65 = 1 - a3
	v66 = 5 - a3*2
	v60 = nox_draw_curDrawData_3799572.clip.left
	v48 = nox_draw_curDrawData_3799572.clip.right
	v3 = int16(*(*uint16)(unsafe.Pointer(&nox_draw_curDrawData_3799572.field_61)))
	v49 = nox_draw_curDrawData_3799572.clip.top
	v50 = nox_draw_curDrawData_3799572.clip.bottom
	v78 = 0
	v75 = 3
	v71 = uint8(int8(int32(uint16(int16(int32(*((*uint16)(unsafe.Pointer(&byte_5D4594_3804364[0+0]))))&int32(v3)))) >> int32(byte_5D4594_3804364[0+12])))
	v4 = uint8(int8(int32(uint16(int16(int32(*((*uint16)(unsafe.Pointer(&byte_5D4594_3804364[0+0]))))&int32(v3)))) >> int32(byte_5D4594_3804364[0+12])))
	v76 = uint8(int8(int32(uint16(int16(int32(*((*uint16)(unsafe.Pointer(&byte_5D4594_3804364[0+4]))))&int32(v3)))) >> int32(byte_5D4594_3804364[0+16])))
	v5 = byte_5D4594_3804364[0+8]
	v6 = uint8(int8(int32(uint16(int16(int32(*((*uint16)(unsafe.Pointer(&byte_5D4594_3804364[0+4]))))&int32(v3)))) >> int32(byte_5D4594_3804364[0+16])))
	v7 = int8(int32(v3) & int32(byte_5D4594_3804364[0+8]))
	v8 = byte_5D4594_3804364[0+20]
	v73 = a3 + a2
	v80 = uint8(int8(int32(v7) << int32(byte_5D4594_3804364[0+20])))
	v9 = uint8(int8(int32(v7) << int32(byte_5D4594_3804364[0+20])))
	if a1 >= v60 && a1 < v48 && v73 >= v49 && v73 < v50 {
		v39 = (*uint16)(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(v73)))))) + uint32(a1*2))))
		v51 = uint32(*v39)
		*v39 = uint16(int16(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(nox_draw_colors_r_3804672), int32(uint8(((*((*uint32)(unsafe.Pointer(&byte_5D4594_3804364[0+0])))&v51)>>uint32(byte_5D4594_3804364[0+12]))+uint32((int32(v71)-int32(uint8((*((*uint32)(unsafe.Pointer(&byte_5D4594_3804364[0+0])))&v51)>>uint32(byte_5D4594_3804364[0+12]))))>>2)))*2))))) | int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(nox_draw_colors_g_3804656), int32(uint8(((*((*uint32)(unsafe.Pointer(&byte_5D4594_3804364[0+4])))&v51)>>uint32(byte_5D4594_3804364[0+16]))+uint32((int32(v76)-int32(uint8((*((*uint32)(unsafe.Pointer(&byte_5D4594_3804364[0+4])))&v51)>>uint32(byte_5D4594_3804364[0+16]))))>>2)))*2))))) | int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(nox_draw_colors_b_3804664), int32(uint8(int8(((int32(*(*uint8)(unsafe.Pointer(v39)))&int32(byte_5D4594_3804364[0+8]))<<int32(byte_5D4594_3804364[0+20]))+((int32(v80)-int32(uint8(int8((int32(*(*uint8)(unsafe.Pointer(v39)))&int32(byte_5D4594_3804364[0+8]))<<int32(byte_5D4594_3804364[0+20])))))>>2))))*2)))))))
		v8 = byte_5D4594_3804364[0+20]
		v5 = byte_5D4594_3804364[0+8]
	}
	v10 = a3 + a1
	v63 = a3 + a1
	if a3+a1 < v60 {
		goto LABEL_17
	}
	if v10 < v48 {
		if a2 >= v49 && a2 < v50 {
			v40 = (*uint16)(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(a2)))))) + uint32(v63*2))))
			v52 = uint32(*v40)
			*v40 = uint16(int16(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(nox_draw_colors_r_3804672), int32(uint8(((*((*uint32)(unsafe.Pointer(&byte_5D4594_3804364[0+0])))&v52)>>uint32(byte_5D4594_3804364[0+12]))+uint32((int32(v71)-int32(uint8((*((*uint32)(unsafe.Pointer(&byte_5D4594_3804364[0+0])))&v52)>>uint32(byte_5D4594_3804364[0+12]))))>>2)))*2))))) | int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(nox_draw_colors_g_3804656), int32(uint8(((*((*uint32)(unsafe.Pointer(&byte_5D4594_3804364[0+4])))&v52)>>uint32(byte_5D4594_3804364[0+16]))+uint32((int32(v76)-int32(uint8((*((*uint32)(unsafe.Pointer(&byte_5D4594_3804364[0+4])))&v52)>>uint32(byte_5D4594_3804364[0+16]))))>>2)))*2))))) | int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(nox_draw_colors_b_3804664), int32(uint8(int8(((int32(*(*uint8)(unsafe.Pointer(v40)))&int32(v5))<<int32(v8))+((int32(v80)-int32(uint8(int8((int32(*(*uint8)(unsafe.Pointer(v40)))&int32(v5))<<int32(v8)))))>>2))))*2)))))))
			v8 = byte_5D4594_3804364[0+20]
			v5 = byte_5D4594_3804364[0+8]
		}
		v10 = a3 + a1
	}
	if v10 < v60 || v10 >= v48 {
		goto LABEL_17
	}
	v11 = a2
	if a2 >= v49 && a2 < v50 {
		v41 = (*uint16)(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(a2)))))) + uint32(v63*2))))
		v53 = uint32(*v41)
		*v41 = uint16(int16(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(nox_draw_colors_r_3804672), int32(uint8(((*((*uint32)(unsafe.Pointer(&byte_5D4594_3804364[0+0])))&v53)>>uint32(byte_5D4594_3804364[0+12]))+uint32((int32(v71)-int32(uint8((*((*uint32)(unsafe.Pointer(&byte_5D4594_3804364[0+0])))&v53)>>uint32(byte_5D4594_3804364[0+12]))))>>2)))*2))))) | int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(nox_draw_colors_g_3804656), int32(uint8(((*((*uint32)(unsafe.Pointer(&byte_5D4594_3804364[0+4])))&v53)>>uint32(byte_5D4594_3804364[0+16]))+uint32((int32(v76)-int32(uint8((*((*uint32)(unsafe.Pointer(&byte_5D4594_3804364[0+4])))&v53)>>uint32(byte_5D4594_3804364[0+16]))))>>2)))*2))))) | int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(nox_draw_colors_b_3804664), int32(uint8(int8(((int32(*(*uint8)(unsafe.Pointer(v41)))&int32(v5))<<int32(v8))+((int32(v80)-int32(uint8(int8((int32(*(*uint8)(unsafe.Pointer(v41)))&int32(v5))<<int32(v8)))))>>2))))*2)))))))
		v8 = byte_5D4594_3804364[0+20]
		v5 = byte_5D4594_3804364[0+8]
	LABEL_17:
		v11 = a2
	}
	v67 = v11 - a3
	if a1 >= v60 && a1 < v48 && v67 >= v49 && v67 < v50 {
		v42 = (*uint16)(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(v67)))))) + uint32(a1*2))))
		v54 = uint32(*v42)
		*v42 = uint16(int16(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(nox_draw_colors_r_3804672), int32(uint8(((*((*uint32)(unsafe.Pointer(&byte_5D4594_3804364[0+0])))&v54)>>uint32(byte_5D4594_3804364[0+12]))+uint32((int32(v71)-int32(uint8((*((*uint32)(unsafe.Pointer(&byte_5D4594_3804364[0+0])))&v54)>>uint32(byte_5D4594_3804364[0+12]))))>>2)))*2))))) | int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(nox_draw_colors_g_3804656), int32(uint8(((*((*uint32)(unsafe.Pointer(&byte_5D4594_3804364[0+4])))&v54)>>uint32(byte_5D4594_3804364[0+16]))+uint32((int32(v76)-int32(uint8((*((*uint32)(unsafe.Pointer(&byte_5D4594_3804364[0+4])))&v54)>>uint32(byte_5D4594_3804364[0+16]))))>>2)))*2))))) | int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(nox_draw_colors_b_3804664), int32(uint8(int8(((int32(*(*uint8)(unsafe.Pointer(v42)))&int32(v5))<<int32(v8))+((int32(v80)-int32(uint8(int8((int32(*(*uint8)(unsafe.Pointer(v42)))&int32(v5))<<int32(v8)))))>>2))))*2)))))))
		v8 = byte_5D4594_3804364[0+20]
		v5 = byte_5D4594_3804364[0+8]
	}
	if a1 >= v60 && a1 < v48 && v67 >= v49 && v67 < v50 {
		v43 = (*uint16)(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(v67)))))) + uint32(a1*2))))
		v55 = uint32(*v43)
		*v43 = uint16(int16(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(nox_draw_colors_r_3804672), int32(uint8(((*((*uint32)(unsafe.Pointer(&byte_5D4594_3804364[0+0])))&v55)>>uint32(byte_5D4594_3804364[0+12]))+uint32((int32(v71)-int32(uint8((*((*uint32)(unsafe.Pointer(&byte_5D4594_3804364[0+0])))&v55)>>uint32(byte_5D4594_3804364[0+12]))))>>2)))*2))))) | int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(nox_draw_colors_g_3804656), int32(uint8(((*((*uint32)(unsafe.Pointer(&byte_5D4594_3804364[0+4])))&v55)>>uint32(byte_5D4594_3804364[0+16]))+uint32((int32(v76)-int32(uint8((*((*uint32)(unsafe.Pointer(&byte_5D4594_3804364[0+4])))&v55)>>uint32(byte_5D4594_3804364[0+16]))))>>2)))*2))))) | int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(nox_draw_colors_b_3804664), int32(uint8(int8(((int32(*(*uint8)(unsafe.Pointer(v43)))&int32(v5))<<int32(v8))+((int32(v80)-int32(uint8(int8((int32(*(*uint8)(unsafe.Pointer(v43)))&int32(v5))<<int32(v8)))))>>2))))*2)))))))
		v8 = byte_5D4594_3804364[0+20]
		v5 = byte_5D4594_3804364[0+8]
	}
	v12 = a1 - a3
	v61 = a1 - a3
	if a1-a3 >= v60 {
		if v12 < v48 {
			if a2 >= v49 && a2 < v50 {
				v44 = (*uint16)(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(a2)))))) + uint32(v61*2))))
				v56 = uint32(*v44)
				*v44 = uint16(int16(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(nox_draw_colors_r_3804672), int32(uint8(((*((*uint32)(unsafe.Pointer(&byte_5D4594_3804364[0+0])))&v56)>>uint32(byte_5D4594_3804364[0+12]))+uint32((int32(v71)-int32(uint8((*((*uint32)(unsafe.Pointer(&byte_5D4594_3804364[0+0])))&v56)>>uint32(byte_5D4594_3804364[0+12]))))>>2)))*2))))) | int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(nox_draw_colors_g_3804656), int32(uint8(((*((*uint32)(unsafe.Pointer(&byte_5D4594_3804364[0+4])))&v56)>>uint32(byte_5D4594_3804364[0+16]))+uint32((int32(v76)-int32(uint8((*((*uint32)(unsafe.Pointer(&byte_5D4594_3804364[0+4])))&v56)>>uint32(byte_5D4594_3804364[0+16]))))>>2)))*2))))) | int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(nox_draw_colors_b_3804664), int32(uint8(int8(((int32(*(*uint8)(unsafe.Pointer(v44)))&int32(v5))<<int32(v8))+((int32(v80)-int32(uint8(int8((int32(*(*uint8)(unsafe.Pointer(v44)))&int32(v5))<<int32(v8)))))>>2))))*2)))))))
				v8 = byte_5D4594_3804364[0+20]
				v5 = byte_5D4594_3804364[0+8]
			}
			v12 = a1 - a3
		}
		if v12 >= v60 && v12 < v48 && a2 >= v49 && a2 < v50 {
			v45 = (*uint16)(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(a2)))))) + uint32(v61*2))))
			v57 = uint32(*v45)
			*v45 = uint16(int16(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(nox_draw_colors_r_3804672), int32(uint8(((*((*uint32)(unsafe.Pointer(&byte_5D4594_3804364[0+0])))&v57)>>uint32(byte_5D4594_3804364[0+12]))+uint32((int32(v71)-int32(uint8((*((*uint32)(unsafe.Pointer(&byte_5D4594_3804364[0+0])))&v57)>>uint32(byte_5D4594_3804364[0+12]))))>>2)))*2))))) | int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(nox_draw_colors_g_3804656), int32(uint8(((*((*uint32)(unsafe.Pointer(&byte_5D4594_3804364[0+4])))&v57)>>uint32(byte_5D4594_3804364[0+16]))+uint32((int32(v76)-int32(uint8((*((*uint32)(unsafe.Pointer(&byte_5D4594_3804364[0+4])))&v57)>>uint32(byte_5D4594_3804364[0+16]))))>>2)))*2))))) | int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(nox_draw_colors_b_3804664), int32(uint8(int8(((int32(*(*uint8)(unsafe.Pointer(v45)))&int32(v5))<<int32(v8))+((int32(v80)-int32(uint8(int8((int32(*(*uint8)(unsafe.Pointer(v45)))&int32(v5))<<int32(v8)))))>>2))))*2)))))))
			v8 = byte_5D4594_3804364[0+20]
			v5 = byte_5D4594_3804364[0+8]
		}
	}
	v13 = a1
	if a1 >= v60 && a1 < v48 {
		if v73 >= v49 && v73 < v50 {
			v46 = (*uint16)(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(v73)))))) + uint32(a1*2))))
			v58 = uint32(*v46)
			*v46 = uint16(int16(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(nox_draw_colors_r_3804672), int32(uint8(((*((*uint32)(unsafe.Pointer(&byte_5D4594_3804364[0+0])))&v58)>>uint32(byte_5D4594_3804364[0+12]))+uint32((int32(v71)-int32(uint8((*((*uint32)(unsafe.Pointer(&byte_5D4594_3804364[0+0])))&v58)>>uint32(byte_5D4594_3804364[0+12]))))>>2)))*2))))) | int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(nox_draw_colors_g_3804656), int32(uint8(((*((*uint32)(unsafe.Pointer(&byte_5D4594_3804364[0+4])))&v58)>>uint32(byte_5D4594_3804364[0+16]))+uint32((int32(v76)-int32(uint8((*((*uint32)(unsafe.Pointer(&byte_5D4594_3804364[0+4])))&v58)>>uint32(byte_5D4594_3804364[0+16]))))>>2)))*2))))) | int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(nox_draw_colors_b_3804664), int32(uint8(int8(((int32(*(*uint8)(unsafe.Pointer(v46)))&int32(v5))<<int32(v8))+((int32(v80)-int32(uint8(int8((int32(*(*uint8)(unsafe.Pointer(v46)))&int32(v5))<<int32(v8)))))>>2))))*2)))))))
		}
		v13 = a1
	}
	result = a3
	if a3 > 0 {
		v59 = a3 + a2
		v15 = a2
		v82 = a1 - a3
		v72 = v61 * 2
		v69 = v15 * 4
		v83 = v13
		v62 = v13 * 2
		v47 = v67 * 4
		v64 = v63 * 2
		v16 = v13 * 2
		v17 = v15 - v13
		v18 = v13 - v15
		v70 = v15
		v74 = v16
		v68 = v17
		for i = v18; ; v18 = i {
			if v65 >= 0 {
				v65 += v66
				v66 += 4
				v79--
				v64 -= 2
				v47 += 4
				v72 += 2
				v82++
				v19 = func() int32 {
					p := &v59
					*p--
					return *p
				}()
			} else {
				v65 += v75
				v19 = v59
				v66 += 2
			}
			v75 += 2
			v78++
			v62 += 2
			v69 -= 4
			v20 = v70 + 1
			v70 = v20
			v74 -= 2
			v21 = v18 + v20
			v83--
			v77 = v21
			if v21 >= v60 && v21 < v48 && v19 >= v49 && v19 < v50 {
				v22 = (*uint16)(unsafe.Pointer(uintptr(uint32(v62) + uint32(uintptr(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(v19)))))))))
				v23 = uint32(*v22)
				*v22 = uint16(int16(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(nox_draw_colors_r_3804672), int32(uint8(((*((*uint32)(unsafe.Pointer(&byte_5D4594_3804364[0+0])))&v23)>>uint32(byte_5D4594_3804364[0+12]))+uint32((int32(v4)-int32(uint8((*((*uint32)(unsafe.Pointer(&byte_5D4594_3804364[0+0])))&v23)>>uint32(byte_5D4594_3804364[0+12]))))>>2)))*2))))) | int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(nox_draw_colors_g_3804656), int32(uint8(((*((*uint32)(unsafe.Pointer(&byte_5D4594_3804364[0+4])))&v23)>>uint32(byte_5D4594_3804364[0+16]))+uint32((int32(v6)-int32(uint8((*((*uint32)(unsafe.Pointer(&byte_5D4594_3804364[0+4])))&v23)>>uint32(byte_5D4594_3804364[0+16]))))>>2)))*2))))) | int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(nox_draw_colors_b_3804664), int32(uint8(int8(((int32(uint8(*v22))&int32(byte_5D4594_3804364[0+8]))<<int32(byte_5D4594_3804364[0+20]))+((int32(v9)-int32(uint8(int8((int32(uint8(*v22))&int32(byte_5D4594_3804364[0+8]))<<int32(byte_5D4594_3804364[0+20])))))>>2))))*2)))))))
				v21 = v77
			}
			v24 = v59 + i
			if v59+i >= v60 {
				if v24 < v48 && v70 >= v49 && v70 < v50 {
					v25 = (*uint16)(unsafe.Pointer(uintptr(uint32(v64) + uint32(uintptr(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(v70)))))))))
					v26 = uint32(*v25)
					v24 = v59 + i
					*v25 = uint16(int16(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(nox_draw_colors_r_3804672), int32(uint8(((*((*uint32)(unsafe.Pointer(&byte_5D4594_3804364[0+0])))&v26)>>uint32(byte_5D4594_3804364[0+12]))+uint32((int32(v4)-int32(uint8((*((*uint32)(unsafe.Pointer(&byte_5D4594_3804364[0+0])))&v26)>>uint32(byte_5D4594_3804364[0+12]))))>>2)))*2))))) | int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(nox_draw_colors_g_3804656), int32(uint8(((*((*uint32)(unsafe.Pointer(&byte_5D4594_3804364[0+4])))&v26)>>uint32(byte_5D4594_3804364[0+16]))+uint32((int32(v6)-int32(uint8((*((*uint32)(unsafe.Pointer(&byte_5D4594_3804364[0+4])))&v26)>>uint32(byte_5D4594_3804364[0+16]))))>>2)))*2))))) | int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(nox_draw_colors_b_3804664), int32(uint8(int8(((int32(uint8(*v25))&int32(byte_5D4594_3804364[0+8]))<<int32(byte_5D4594_3804364[0+20]))+((int32(v9)-int32(uint8(int8((int32(uint8(*v25))&int32(byte_5D4594_3804364[0+8]))<<int32(byte_5D4594_3804364[0+20])))))>>2))))*2)))))))
					v21 = v77
				}
				if v24 >= v60 && v24 < v48 && v68+v83 >= v49 && v68+v83 < v50 {
					v27 = (*uint16)(unsafe.Pointer(uintptr(uint32(v64) + uint32(uintptr(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(v69/4)))))))))
					v28 = uint32(*v27)
					*v27 = uint16(int16(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(nox_draw_colors_r_3804672), int32(uint8(((*((*uint32)(unsafe.Pointer(&byte_5D4594_3804364[0+0])))&v28)>>uint32(byte_5D4594_3804364[0+12]))+uint32((int32(v4)-int32(uint8((*((*uint32)(unsafe.Pointer(&byte_5D4594_3804364[0+0])))&v28)>>uint32(byte_5D4594_3804364[0+12]))))>>2)))*2))))) | int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(nox_draw_colors_g_3804656), int32(uint8(((*((*uint32)(unsafe.Pointer(&byte_5D4594_3804364[0+4])))&v28)>>uint32(byte_5D4594_3804364[0+16]))+uint32((int32(v6)-int32(uint8((*((*uint32)(unsafe.Pointer(&byte_5D4594_3804364[0+4])))&v28)>>uint32(byte_5D4594_3804364[0+16]))))>>2)))*2))))) | int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(nox_draw_colors_b_3804664), int32(uint8(int8(((int32(uint8(*v27))&int32(byte_5D4594_3804364[0+8]))<<int32(byte_5D4594_3804364[0+20]))+((int32(v9)-int32(uint8(int8((int32(uint8(*v27))&int32(byte_5D4594_3804364[0+8]))<<int32(byte_5D4594_3804364[0+20])))))>>2))))*2)))))))
					v21 = v77
				}
			}
			if v21 >= v60 && v21 < v48 && v68+v82 >= v49 && v68+v82 < v50 {
				v29 = (*uint16)(unsafe.Pointer(uintptr(uint32(v62) + uint32(uintptr(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(v47/4)))))))))
				v30 = uint32(*v29)
				*v29 = uint16(int16(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(nox_draw_colors_r_3804672), int32(uint8(((*((*uint32)(unsafe.Pointer(&byte_5D4594_3804364[0+0])))&v30)>>uint32(byte_5D4594_3804364[0+12]))+uint32((int32(v4)-int32(uint8((*((*uint32)(unsafe.Pointer(&byte_5D4594_3804364[0+0])))&v30)>>uint32(byte_5D4594_3804364[0+12]))))>>2)))*2))))) | int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(nox_draw_colors_g_3804656), int32(uint8(((*((*uint32)(unsafe.Pointer(&byte_5D4594_3804364[0+4])))&v30)>>uint32(byte_5D4594_3804364[0+16]))+uint32((int32(v6)-int32(uint8((*((*uint32)(unsafe.Pointer(&byte_5D4594_3804364[0+4])))&v30)>>uint32(byte_5D4594_3804364[0+16]))))>>2)))*2))))) | int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(nox_draw_colors_b_3804664), int32(uint8(int8(((int32(uint8(*v29))&int32(byte_5D4594_3804364[0+8]))<<int32(byte_5D4594_3804364[0+20]))+((int32(v9)-int32(uint8(int8((int32(uint8(*v29))&int32(byte_5D4594_3804364[0+8]))<<int32(byte_5D4594_3804364[0+20])))))>>2))))*2)))))))
			}
			if v83 >= v60 && v83 < v48 && v68+v82 >= v49 && v68+v82 < v50 {
				v31 = (*uint16)(unsafe.Pointer(uintptr(uint32(v74) + uint32(uintptr(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(v47/4)))))))))
				v32 = uint32(*v31)
				*v31 = uint16(int16(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(nox_draw_colors_r_3804672), int32(uint8(((*((*uint32)(unsafe.Pointer(&byte_5D4594_3804364[0+0])))&v32)>>uint32(byte_5D4594_3804364[0+12]))+uint32((int32(v4)-int32(uint8((*((*uint32)(unsafe.Pointer(&byte_5D4594_3804364[0+0])))&v32)>>uint32(byte_5D4594_3804364[0+12]))))>>2)))*2))))) | int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(nox_draw_colors_g_3804656), int32(uint8(((*((*uint32)(unsafe.Pointer(&byte_5D4594_3804364[0+4])))&v32)>>uint32(byte_5D4594_3804364[0+16]))+uint32((int32(v6)-int32(uint8((*((*uint32)(unsafe.Pointer(&byte_5D4594_3804364[0+4])))&v32)>>uint32(byte_5D4594_3804364[0+16]))))>>2)))*2))))) | int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(nox_draw_colors_b_3804664), int32(uint8(int8(((int32(uint8(*v31))&int32(byte_5D4594_3804364[0+8]))<<int32(byte_5D4594_3804364[0+20]))+((int32(v9)-int32(uint8(int8((int32(uint8(*v31))&int32(byte_5D4594_3804364[0+8]))<<int32(byte_5D4594_3804364[0+20])))))>>2))))*2)))))))
			}
			if v82 >= v60 {
				if v82 < v48 && v68+v83 >= v49 && v68+v83 < v50 {
					v33 = (*uint16)(unsafe.Pointer(uintptr(uint32(v72) + uint32(uintptr(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(v69/4)))))))))
					v34 = uint32(*v33)
					*v33 = uint16(int16(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(nox_draw_colors_r_3804672), int32(uint8(((*((*uint32)(unsafe.Pointer(&byte_5D4594_3804364[0+0])))&v34)>>uint32(byte_5D4594_3804364[0+12]))+uint32((int32(v4)-int32(uint8((*((*uint32)(unsafe.Pointer(&byte_5D4594_3804364[0+0])))&v34)>>uint32(byte_5D4594_3804364[0+12]))))>>2)))*2))))) | int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(nox_draw_colors_g_3804656), int32(uint8(((*((*uint32)(unsafe.Pointer(&byte_5D4594_3804364[0+4])))&v34)>>uint32(byte_5D4594_3804364[0+16]))+uint32((int32(v6)-int32(uint8((*((*uint32)(unsafe.Pointer(&byte_5D4594_3804364[0+4])))&v34)>>uint32(byte_5D4594_3804364[0+16]))))>>2)))*2))))) | int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(nox_draw_colors_b_3804664), int32(uint8(int8(((int32(uint8(*v33))&int32(byte_5D4594_3804364[0+8]))<<int32(byte_5D4594_3804364[0+20]))+((int32(v9)-int32(uint8(int8((int32(uint8(*v33))&int32(byte_5D4594_3804364[0+8]))<<int32(byte_5D4594_3804364[0+20])))))>>2))))*2)))))))
				}
				if v82 < v48 && v70 >= v49 && v70 < v50 {
					v35 = (*uint16)(unsafe.Pointer(uintptr(uint32(v72) + uint32(uintptr(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(v70)))))))))
					v36 = uint32(*v35)
					*v35 = uint16(int16(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(nox_draw_colors_r_3804672), int32(uint8(((*((*uint32)(unsafe.Pointer(&byte_5D4594_3804364[0+0])))&v36)>>uint32(byte_5D4594_3804364[0+12]))+uint32((int32(v4)-int32(uint8((*((*uint32)(unsafe.Pointer(&byte_5D4594_3804364[0+0])))&v36)>>uint32(byte_5D4594_3804364[0+12]))))>>2)))*2))))) | int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(nox_draw_colors_g_3804656), int32(uint8(((*((*uint32)(unsafe.Pointer(&byte_5D4594_3804364[0+4])))&v36)>>uint32(byte_5D4594_3804364[0+16]))+uint32((int32(v6)-int32(uint8((*((*uint32)(unsafe.Pointer(&byte_5D4594_3804364[0+4])))&v36)>>uint32(byte_5D4594_3804364[0+16]))))>>2)))*2))))) | int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(nox_draw_colors_b_3804664), int32(uint8(int8(((int32(uint8(*v35))&int32(byte_5D4594_3804364[0+8]))<<int32(byte_5D4594_3804364[0+20]))+((int32(v9)-int32(uint8(int8((int32(uint8(*v35))&int32(byte_5D4594_3804364[0+8]))<<int32(byte_5D4594_3804364[0+20])))))>>2))))*2)))))))
				}
			}
			if v83 >= v60 && v83 < v48 && v59 >= v49 && v59 < v50 {
				v37 = (*uint16)(unsafe.Pointer(uintptr(uint32(v74) + uint32(uintptr(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(pixbuffer), unsafe.Sizeof((*uint8)(nil))*uintptr(v59)))))))))
				v38 = uint32(*v37)
				*v37 = uint16(int16(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(nox_draw_colors_r_3804672), int32(uint8(((*((*uint32)(unsafe.Pointer(&byte_5D4594_3804364[0+0])))&v38)>>uint32(byte_5D4594_3804364[0+12]))+uint32((int32(v4)-int32(uint8((*((*uint32)(unsafe.Pointer(&byte_5D4594_3804364[0+0])))&v38)>>uint32(byte_5D4594_3804364[0+12]))))>>2)))*2))))) | int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(nox_draw_colors_g_3804656), int32(uint8(((*((*uint32)(unsafe.Pointer(&byte_5D4594_3804364[0+4])))&v38)>>uint32(byte_5D4594_3804364[0+16]))+uint32((int32(v6)-int32(uint8((*((*uint32)(unsafe.Pointer(&byte_5D4594_3804364[0+4])))&v38)>>uint32(byte_5D4594_3804364[0+16]))))>>2)))*2))))) | int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(nox_draw_colors_b_3804664), int32(uint8(int8(((int32(uint8(*v37))&int32(byte_5D4594_3804364[0+8]))<<int32(byte_5D4594_3804364[0+20]))+((int32(v9)-int32(uint8(int8((int32(uint8(*v37))&int32(byte_5D4594_3804364[0+8]))<<int32(byte_5D4594_3804364[0+20])))))>>2))))*2)))))))
			}
			result = v79
			if v79 <= v78 {
				break
			}
		}
	}
	return result
}
func sub_4B4860(a1 int32, a2 int32, a3 int32, a4 int32) int32 {
	var (
		v4     int32
		v5     int32
		result int32
		v7     int32
		v8     int32
		v9     int32
		v10    int32
		v11    int32
		v12    int32
		v13    int32
		v14    int32
	)
	v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 32))))
	switch a2 {
	case 5:
		goto LABEL_22
	case 6:
		fallthrough
	case 7:
		v9 = a3
		v10 = int32(uint16(int16(a3)))
		nox_client_wndGetPosition_46AA60((*nox_window)(unsafe.Pointer(uintptr(a1))), (*uint32)(unsafe.Pointer(&a4)), (*uint32)(unsafe.Pointer(&a3)))
		if v10 > *(*int32)(unsafe.Pointer(uintptr(a1 + 8)))+a4-10 {
			v10 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 8))) + uint32(a4) - 10)
		}
		(*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(uintptr(a1 + 400))))).SetPos(image.Pt(v10-a4-5, 0))
		(*nox_window)(unsafe.Pointer(uintptr(a1))).Func94(asWindowEvent(0x4000, 0, v9))
		return 1
	case 8:
		v8 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 44))))
		if (v8 & 256) == 0 {
			goto LABEL_22
		}
		(*nox_window)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 52)))))).Func94(asWindowEvent(0x4000, a1, 0))
		result = 1
	case 17:
		v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 44))))
		if (v5 & 256) == 0 {
			goto LABEL_22
		}
		nox_xxx_wndSetRectColor2MB_46AFE0((*nox_window)(unsafe.Pointer(uintptr(a1))), int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 72)))))
		(*nox_window)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 52)))))).Func94(asWindowEvent(0x4005, a1, 0))
		guiFocus((*nox_window)(unsafe.Pointer(uintptr(a1))))
		result = 1
	case 18:
		v7 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 44))))
		if v7&256 != 0 {
			nox_xxx_wndSetRectColor2MB_46AFE0((*nox_window)(unsafe.Pointer(uintptr(a1))), int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 64)))))
			(*nox_window)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 52)))))).Func94(asWindowEvent(0x4006, a1, 0))
			result = 1
		} else {
		LABEL_22:
			result = 1
		}
	case 21:
		switch a3 {
		case 15:
			fallthrough
		case 208:
			if a4 != 2 {
				goto LABEL_22
			}
			nox_xxx_wndRetNULL_46A8A0()
			result = 1
		case 200:
			if a4 == 2 {
				nox_xxx_wndRetNULL_0_46A8B0()
			}
			goto LABEL_22
		case 203:
			if a4 != 2 {
				goto LABEL_22
			}
			v13 = int32(*(*uint32)(unsafe.Pointer(uintptr(v4 + 12))))
			if v13 >= *(*int32)(unsafe.Pointer(uintptr(v4 + 4)))-1 {
				goto LABEL_22
			}
			v14 = v13 + 2
			*(*uint32)(unsafe.Pointer(uintptr(v4 + 12))) = uint32(v14)
			(*nox_window)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 52)))))).Func94(asWindowEvent(0x4009, a1, v14))
			(*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(uintptr(a1 + 400))))).SetPos(image.Pt(int32(int64(float64(int32(*(*uint32)(unsafe.Pointer(uintptr(v4 + 12)))-*(*uint32)(unsafe.Pointer(uintptr(v4)))))*float64(*(*float32)(unsafe.Pointer(uintptr(v4 + 8)))))), 0))
			result = 1
		case 205:
			if a4 != 2 {
				goto LABEL_22
			}
			v11 = int32(*(*uint32)(unsafe.Pointer(uintptr(v4 + 12))))
			if v11 <= *(*int32)(unsafe.Pointer(uintptr(v4)))+1 {
				goto LABEL_22
			}
			v12 = v11 - 2
			*(*uint32)(unsafe.Pointer(uintptr(v4 + 12))) = uint32(v12)
			(*nox_window)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 52)))))).Func94(asWindowEvent(0x4009, a1, v12))
			(*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(uintptr(a1 + 400))))).SetPos(image.Pt(int32(int64(float64(int32(*(*uint32)(unsafe.Pointer(uintptr(v4 + 12)))-*(*uint32)(unsafe.Pointer(uintptr(v4)))))*float64(*(*float32)(unsafe.Pointer(uintptr(v4 + 8)))))), 0))
			result = 1
		default:
			goto LABEL_23
		}
	default:
	LABEL_23:
		result = 0
	}
	return result
}
func nox_xxx_wndScrollBoxDraw_4B4BA0(a1 int32, a2 int32, a3 uint32, a4 int32) int32 {
	var (
		v4     int32
		v5     int32
		result int32
		v7     int32
		v8     int32
		v9     int32
		v10    int32
		v11    int32
		v12    int32
		v13    int32
		v14    int32
	)
	v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 32))))
	switch a2 {
	case 5:
		goto LABEL_22
	case 6:
		fallthrough
	case 7:
		v9 = int32(a3)
		v10 = int32(a3 >> 16)
		nox_client_wndGetPosition_46AA60((*nox_window)(unsafe.Pointer(uintptr(a1))), &a3, (*uint32)(unsafe.Pointer(&a4)))
		if v10 > *(*int32)(unsafe.Pointer(uintptr(a1 + 12)))+a4-10 {
			v10 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 12))) + uint32(a4) - 10)
		}
		(*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(uintptr(a1 + 400))))).SetPos(image.Pt(0, v10-a4-5))
		(*nox_window)(unsafe.Pointer(uintptr(a1))).Func94(asWindowEvent(0x4000, 0, v9))
		return 1
	case 8:
		v8 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 44))))
		if (v8 & 256) == 0 {
			goto LABEL_22
		}
		(*nox_window)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 52)))))).Func94(asWindowEvent(0x4000, a1, 0))
		result = 1
	case 17:
		v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 44))))
		if (v5 & 256) == 0 {
			goto LABEL_22
		}
		nox_xxx_wndSetRectColor2MB_46AFE0((*nox_window)(unsafe.Pointer(uintptr(a1))), int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 72)))))
		(*nox_window)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 52)))))).Func94(asWindowEvent(0x4005, a1, 0))
		guiFocus((*nox_window)(unsafe.Pointer(uintptr(a1))))
		result = 1
	case 18:
		v7 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 44))))
		if v7&256 != 0 {
			nox_xxx_wndSetRectColor2MB_46AFE0((*nox_window)(unsafe.Pointer(uintptr(a1))), int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 64)))))
			(*nox_window)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 52)))))).Func94(asWindowEvent(0x4006, a1, 0))
			result = 1
		} else {
		LABEL_22:
			result = 1
		}
	case 21:
		switch a3 {
		case 15:
			fallthrough
		case 205:
			if a4 != 2 {
				goto LABEL_22
			}
			nox_xxx_wndRetNULL_46A8A0()
			result = 1
		case 200:
			if a4 != 2 {
				goto LABEL_22
			}
			v11 = int32(*(*uint32)(unsafe.Pointer(uintptr(v4 + 12))))
			if v11 >= *(*int32)(unsafe.Pointer(uintptr(v4 + 4)))-1 {
				goto LABEL_22
			}
			v12 = v11 + 2
			*(*uint32)(unsafe.Pointer(uintptr(v4 + 12))) = uint32(v12)
			(*nox_window)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 52)))))).Func94(asWindowEvent(0x4009, a1, v12))
			(*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(uintptr(a1 + 400))))).SetPos(image.Pt(0, int32(int64(float64(int32(*(*uint32)(unsafe.Pointer(uintptr(v4 + 4)))-*(*uint32)(unsafe.Pointer(uintptr(v4 + 12)))))*float64(*(*float32)(unsafe.Pointer(uintptr(v4 + 8))))))))
			result = 1
		case 203:
			if a4 == 2 {
				nox_xxx_wndRetNULL_0_46A8B0()
			}
			goto LABEL_22
		case 208:
			if a4 != 2 {
				goto LABEL_22
			}
			v13 = int32(*(*uint32)(unsafe.Pointer(uintptr(v4 + 12))))
			if v13 <= *(*int32)(unsafe.Pointer(uintptr(v4)))+1 {
				goto LABEL_22
			}
			v14 = v13 - 2
			*(*uint32)(unsafe.Pointer(uintptr(v4 + 12))) = uint32(v14)
			(*nox_window)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 52)))))).Func94(asWindowEvent(0x4009, a1, v14))
			(*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(uintptr(a1 + 400))))).SetPos(image.Pt(0, int32(int64(float64(int32(*(*uint32)(unsafe.Pointer(uintptr(v4 + 4)))-*(*uint32)(unsafe.Pointer(uintptr(v4 + 12)))))*float64(*(*float32)(unsafe.Pointer(uintptr(v4 + 8))))))))
			result = 1
		default:
			goto LABEL_23
		}
	default:
	LABEL_23:
		result = 0
	}
	return result
}
func nox_gui_newSlider_4B4EE0(a1 int32, a2 int32, a3 int32, a4 int32, a5 int32, a6 int32, a7 *uint32, a8 *float32) *nox_window {
	var (
		v8  int32
		v9  int32
		v10 *uint32
		v11 int32
		v12 float64
		v13 *float32
	)
	v8 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a7), unsafe.Sizeof(uint32(0))*2)))
	if v8&16 != 0 {
		v9 = a6
		v10 = (*uint32)(unsafe.Pointer(nox_window_new((*nox_window)(unsafe.Pointer(uintptr(a1))), a2|256, a3, a4, a5, a6, func(arg1 int32, arg2 int32, arg3 int32, arg4 int32) int32 {
			return sub_4B5010(arg1, uint32(arg2), arg3, arg4)
		})))
		sub_4B51A0(int32(uintptr(unsafe.Pointer(v10))))
	} else {
		if (v8 & 8) == 0 {
			return nil
		}
		v9 = a6
		v10 = (*uint32)(unsafe.Pointer(nox_window_new((*nox_window)(unsafe.Pointer(uintptr(a1))), a2|256, a3, a4, a5, a6, func(arg1 int32, arg2 int32, arg3 int32, arg4 int32) int32 {
			return nox_xxx_wndScrollBoxProc_4B5320(arg1, uint32(arg2), arg3, uint32(arg4))
		})))
		nox_xxx_wndScrollBoxSetAllFn_4B5500(int32(uintptr(unsafe.Pointer(v10))))
	}
	if v10 != nil {
		if *(*uint32)(unsafe.Add(unsafe.Pointer(a7), unsafe.Sizeof(uint32(0))*4)) == 0 {
			*(*uint32)(unsafe.Add(unsafe.Pointer(a7), unsafe.Sizeof(uint32(0))*4)) = uint32(uintptr(unsafe.Pointer(v10)))
		}
		nox_gui_windowCopyDrawData_46AF80((*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v10)))))), unsafe.Pointer(a7))
		nox_xxx_wndScrollBoxButtonCreate_4B5640(int32(uintptr(unsafe.Pointer(v10))), a2|256, int32(uintptr(unsafe.Pointer(a7))))
		v11 = int32(*(*uint32)(unsafe.Pointer(a8)))
		if *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a8))), unsafe.Sizeof(uint32(0))*1))) == *(*uint32)(unsafe.Pointer(a8)) {
			*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a8))), unsafe.Sizeof(uint32(0))*1))) = uint32(v11 + 1)
		}
		if *(*uint32)(unsafe.Add(unsafe.Pointer(a7), unsafe.Sizeof(uint32(0))*2))&16 != 0 {
			v12 = float64(a5-10) / float64(int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a8))), unsafe.Sizeof(uint32(0))*1)))-uint32(v11)))
		} else {
			v12 = float64(v9-10) / float64(int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a8))), unsafe.Sizeof(uint32(0))*1)))-uint32(v11)))
		}
		*(*float32)(unsafe.Add(unsafe.Pointer(a8), unsafe.Sizeof(float32(0))*2)) = float32(v12)
		v13 = (*float32)(alloc.Calloc(1, 16))
		*v13 = *a8
		*(*float32)(unsafe.Add(unsafe.Pointer(v13), unsafe.Sizeof(float32(0))*1)) = *(*float32)(unsafe.Add(unsafe.Pointer(a8), unsafe.Sizeof(float32(0))*1))
		*(*float32)(unsafe.Add(unsafe.Pointer(v13), unsafe.Sizeof(float32(0))*2)) = *(*float32)(unsafe.Add(unsafe.Pointer(a8), unsafe.Sizeof(float32(0))*2))
		*(*float32)(unsafe.Add(unsafe.Pointer(v13), unsafe.Sizeof(float32(0))*3)) = *(*float32)(unsafe.Add(unsafe.Pointer(a8), unsafe.Sizeof(float32(0))*3))
		*(*uint32)(unsafe.Add(unsafe.Pointer(v10), unsafe.Sizeof(uint32(0))*8)) = uint32(uintptr(unsafe.Pointer(v13)))
	}
	return (*nox_window)(unsafe.Pointer(v10))
}
func sub_4B5010(a1 int32, a2 uint32, a3 int32, a4 int32) int32 {
	var (
		v4  int32
		v5  int32
		v6  int32
		v7  int32
		v8  int32
		v10 int32
		v11 int32
		v12 int32
		v13 int64
		v14 int32
		v15 int32
		v16 int32
		v17 float64
		v18 *int32
	)
	v4 = a1
	v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 32))))
	if a2 > 0x4000 {
		if a2 == 0x400A {
			if a3 >= *(*int32)(unsafe.Pointer(uintptr(v5))) && a3 <= *(*int32)(unsafe.Pointer(uintptr(v5 + 4))) {
				v17 = float64(a3)
				*(*uint32)(unsafe.Pointer(uintptr(v5 + 12))) = uint32(a3)
				(*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(uintptr(v4 + 400))))).SetPos(image.Pt(int32(int64(v17*float64(*(*float32)(unsafe.Pointer(uintptr(v5 + 8)))))), 0))
			}
		} else if a2 == 0x400B {
			v14 = a3
			v15 = a4
			*(*uint32)(unsafe.Pointer(uintptr(v5))) = uint32(a3)
			*(*uint32)(unsafe.Pointer(uintptr(v5 + 4))) = uint32(v15)
			v16 = int32(*(*uint32)(unsafe.Pointer(uintptr(v4 + 8))) - 10)
			*(*uint32)(unsafe.Pointer(uintptr(v5 + 12))) = uint32(v14)
			a3 = v15 - v14
			*(*float32)(unsafe.Pointer(uintptr(v5 + 8))) = float32(float64(v16) / float64(v15-v14))
			(*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(uintptr(v4 + 400))))).SetPos(image.Pt(0, 0))
			return 0
		}
		return 0
	}
	if a2 == 0x4000 {
		v10 = int32(uint16(int16(a4)))
		nox_client_wndGetPosition_46AA60((*nox_window)(unsafe.Pointer(uintptr(a1))), (*uint32)(unsafe.Pointer(&a3)), (*uint32)(unsafe.Pointer(&a1)))
		if v10 < a3 {
			(*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(uintptr(v4 + 400))))).SetPos(image.Pt(0, 0))
			v11 = int32(*(*uint32)(unsafe.Pointer(uintptr(v5))))
		LABEL_15:
			*(*uint32)(unsafe.Pointer(uintptr(v5 + 12))) = uint32(v11)
			goto LABEL_16
		}
		v12 = int32(*(*uint32)(unsafe.Pointer(uintptr(v4 + 8))))
		if v10 < v12+a3 {
			a4 = v10 - a3
			v13 = int64(float64(v10-a3) / float64(*(*float32)(unsafe.Pointer(uintptr(v5 + 8)))))
			v11 = int32(*(*uint32)(unsafe.Pointer(uintptr(v5 + 4))))
			*(*uint32)(unsafe.Pointer(uintptr(v5 + 12))) = uint32(int32(v13))
			if int32(v13) > v11 {
				goto LABEL_15
			}
		} else {
			(*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(uintptr(v4 + 400))))).SetPos(image.Pt(int32(uint32(v12)-*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v4 + 400))) + 8)))), 0))
			*(*uint32)(unsafe.Pointer(uintptr(v5 + 12))) = *(*uint32)(unsafe.Pointer(uintptr(v5 + 4)))
		}
	LABEL_16:
		(*nox_window)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v4 + 52)))))).Func94(asWindowEvent(0x4009, v4, int32(*(*uint32)(unsafe.Pointer(uintptr(v5 + 12))))))
		return 0
	}
	if a2 == 2 {
		*(*unsafe.Pointer)(unsafe.Pointer(uintptr(a1 + 32))) = nil
		return 0
	}
	if a2 != 23 {
		return 0
	}
	v6 = a3
	v7 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 36))))
	if a3 != 0 {
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v7))), 0)) = uint8(int8(v7 | 2))
	} else {
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v7))), 0)) = uint8(int8(v7 & 253))
	}
	v18 = (*int32)(unsafe.Pointer(uintptr(a1)))
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 36))) = uint32(v7)
	v8 = (*nox_window)(unsafe.Pointer(v18)).ID()
	(*nox_window)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v4 + 52)))))).Func94(asWindowEvent(0x4003, v6, v8))
	return 1
}
func sub_4B51A0(a1 int32) int32 {
	var result int32
	result = a1
	if a1 != 0 {
		if int32(int8(*(*uint8)(unsafe.Pointer(uintptr(a1 + 4))))) >= 0 {
			result = (*nox_window)(unsafe.Pointer(uintptr(a1))).SetAllFuncs(sub_4B4860, func(arg1 *nox_window, arg2 unsafe.Pointer) int32 {
				return sub_4B51E0(int32(uintptr(unsafe.Pointer(arg1))), int32(uintptr(arg2)))
			}, nil)
		} else {
			result = (*nox_window)(unsafe.Pointer(uintptr(a1))).SetAllFuncs(sub_4B4860, func(arg1 *nox_window, arg2 unsafe.Pointer) int32 {
				return sub_4B52C0(int32(uintptr(unsafe.Pointer(arg1))), int32(uintptr(arg2)))
			}, nil)
		}
	}
	return result
}
func sub_4B51E0(a1 int32, a2 int32) int32 {
	var (
		v2    int32
		v3    int32
		v4    int32
		xLeft int32
		yTop  int32
	)
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 28))))
	v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 20))))
	nox_client_wndGetPosition_46AA60((*nox_window)(unsafe.Pointer(uintptr(a1))), (*uint32)(unsafe.Pointer(&xLeft)), (*uint32)(unsafe.Pointer(&yTop)))
	if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 4))))&8 != 0 {
		if int32(*(*uint8)(unsafe.Pointer(uintptr(a2))))&2 != 0 && *(*uint32)(unsafe.Pointer(uintptr(a2 + 36))) != 0x80000000 {
			nox_client_drawSetColor_434460(int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 36)))))
			nox_client_drawBorderLines_49CC70(xLeft, yTop, int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 8)))), int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 12)))))
		}
	} else {
		v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 44))))
	}
	if uint32(v3) != 0x80000000 {
		nox_client_drawSetColor_434460(v3)
		nox_client_drawRectFilledOpaque_49CE30(xLeft+1, yTop+1, int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 8)))-2), int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 12)))-2))
	}
	v4 = int32(uint32(yTop) + *(*uint32)(unsafe.Pointer(uintptr(a1 + 12)))/2)
	if uint32(v2) != 0x80000000 {
		nox_client_drawSetColor_434460(v2)
		nox_client_drawRectFilledOpaque_49CE30(xLeft, v4-1, int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 8)))), 3)
	}
	return 1
}
func sub_4B52C0(a1 int32, a2 int32) int32 {
	var (
		v2 int32
		v4 int32
		v5 int32
	)
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 24))))
	nox_client_wndGetPosition_46AA60((*nox_window)(unsafe.Pointer(uintptr(a1))), (*uint32)(unsafe.Pointer(&v4)), (*uint32)(unsafe.Pointer(&v5)))
	if (int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 4)))) & 8) == 0 {
		v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 48))))
	}
	if v2 != 0 {
		nox_client_drawImageAt_47D2C0((*nox_video_bag_image_t)(unsafe.Pointer(uintptr(v2))), v4, v5)
	}
	return 1
}
func nox_xxx_wndScrollBoxProc_4B5320(a1 int32, a2 uint32, a3 int32, a4 uint32) int32 {
	var (
		v4  int32
		v5  int32
		v6  int32
		v7  int32
		v8  int64
		v9  int32
		v11 int32
		v12 int32
		v13 int32
		v14 uint32
		v15 int32
		v16 int32
		v17 float64
		v18 int32
		v19 int32
		v20 *int32
	)
	v4 = a1
	v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 32))))
	if a2 > 0x4007 {
		if a2 == 0x400A {
			if a3 >= *(*int32)(unsafe.Pointer(uintptr(v5))) {
				v18 = int32(*(*uint32)(unsafe.Pointer(uintptr(v5 + 4))))
				if a3 <= v18 {
					v19 = v18 - a3
					*(*uint32)(unsafe.Pointer(uintptr(v5 + 12))) = uint32(a3)
					a3 = v19
					(*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(uintptr(v4 + 400))))).SetPos(image.Pt(0, int32(int64(float64(v19)*float64(*(*float32)(unsafe.Pointer(uintptr(v5 + 8))))))))
				}
			}
		} else if a2 == 0x400B {
			v14 = a4
			v15 = a3
			*(*uint32)(unsafe.Pointer(uintptr(v5 + 4))) = a4
			*(*uint32)(unsafe.Pointer(uintptr(v5))) = uint32(v15)
			a3 = int32(v14 - uint32(v15))
			v16 = int32(*(*uint32)(unsafe.Pointer(uintptr(v4 + 12))))
			*(*uint32)(unsafe.Pointer(uintptr(v5 + 12))) = uint32(v15)
			v17 = float64(a3)
			a3 = v16 - 10
			*(*float32)(unsafe.Pointer(uintptr(v5 + 8))) = float32(float64(v16-10) / v17)
			(*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(uintptr(v4 + 400))))).SetPos(image.Pt(0, int32(int64(v17*float64(*(*float32)(unsafe.Pointer(uintptr(v5 + 8))))))))
			return 0
		}
		return 0
	}
	switch a2 {
	case 0x4007:
		(*nox_window)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 52)))))).Func94(asWindowEvent(0x400C, a1, int32(*(*uint32)(unsafe.Pointer(uintptr(v5 + 12))))))
		return 0
	case 2:
		*(*unsafe.Pointer)(unsafe.Pointer(uintptr(a1 + 32))) = nil
		return 0
	case 23:
		v11 = a3
		v12 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 36))))
		if a3 != 0 {
			*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v12))), 0)) = uint8(int8(v12 | 2))
		} else {
			*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v12))), 0)) = uint8(int8(v12 & 253))
		}
		v20 = (*int32)(unsafe.Pointer(uintptr(a1)))
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 36))) = uint32(v12)
		v13 = (*nox_window)(unsafe.Pointer(v20)).ID()
		(*nox_window)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v4 + 52)))))).Func94(asWindowEvent(0x4003, v11, v13))
		return 1
	}
	if a2 != 0x4000 {
		return 0
	}
	v6 = int32(a4 >> 16)
	nox_client_wndGetPosition_46AA60((*nox_window)(unsafe.Pointer(uintptr(a1))), (*uint32)(unsafe.Pointer(&a1)), (*uint32)(unsafe.Pointer(&a3)))
	if v6 >= a3 {
		v7 = int32(*(*uint32)(unsafe.Pointer(uintptr(v4 + 12))))
		if v6 < v7+a3 {
			a4 = uint32(v6 - a3)
			v8 = int64(float64(v6-a3) / float64(*(*float32)(unsafe.Pointer(uintptr(v5 + 8)))))
			v9 = int32(*(*uint32)(unsafe.Pointer(uintptr(v5 + 4))))
			*(*uint32)(unsafe.Pointer(uintptr(v5 + 12))) = uint32(int32(v8))
			if int32(v8) > v9 {
				*(*uint32)(unsafe.Pointer(uintptr(v5 + 12))) = uint32(v9)
			}
			*(*uint32)(unsafe.Pointer(uintptr(v5 + 12))) = uint32(v9) - *(*uint32)(unsafe.Pointer(uintptr(v5 + 12)))
		} else {
			(*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(uintptr(v4 + 400))))).SetPos(image.Pt(0, int32(uint32(v7)-*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v4 + 400))) + 12))))))
			*(*uint32)(unsafe.Pointer(uintptr(v5 + 12))) = *(*uint32)(unsafe.Pointer(uintptr(v5)))
		}
	} else {
		(*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(uintptr(v4 + 400))))).SetPos(image.Pt(0, 0))
		*(*uint32)(unsafe.Pointer(uintptr(v5 + 12))) = *(*uint32)(unsafe.Pointer(uintptr(v5 + 4)))
	}
	(*nox_window)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v4 + 52)))))).Func94(asWindowEvent(0x4009, v4, int32(*(*uint32)(unsafe.Pointer(uintptr(v5 + 12))))))
	return 0
}
func nox_xxx_wndScrollBoxSetAllFn_4B5500(a1 int32) int32 {
	var result int32
	result = a1
	if a1 != 0 {
		if int32(int8(*(*uint8)(unsafe.Pointer(uintptr(a1 + 4))))) >= 0 {
			result = (*nox_window)(unsafe.Pointer(uintptr(a1))).SetAllFuncs(func(arg1 int32, arg2 int32, arg3 int32, arg4 int32) int32 {
				return nox_xxx_wndScrollBoxDraw_4B4BA0(arg1, arg2, uint32(arg3), arg4)
			}, func(arg1 *nox_window, arg2 unsafe.Pointer) int32 {
				return nox_xxx_wndScrollBoxDraw_4B5540(int32(uintptr(unsafe.Pointer(arg1))), int32(uintptr(arg2)))
			}, nil)
		} else {
			result = (*nox_window)(unsafe.Pointer(uintptr(a1))).SetAllFuncs(func(arg1 int32, arg2 int32, arg3 int32, arg4 int32) int32 {
				return nox_xxx_wndScrollBoxDraw_4B4BA0(arg1, arg2, uint32(arg3), arg4)
			}, func(arg1 *nox_window, arg2 unsafe.Pointer) int32 {
				return nox_xxx_wndScrollBoxDraw_4B5620((*uint32)(unsafe.Pointer(arg1)), int32(uintptr(arg2)))
			}, nil)
		}
	}
	return result
}
func nox_xxx_wndScrollBoxDraw_4B5540(a1 int32, a2 int32) int32 {
	var (
		v2    int32
		v3    int32
		v4    int32
		xLeft int32
		yTop  int32
	)
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 28))))
	v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 20))))
	nox_client_wndGetPosition_46AA60((*nox_window)(unsafe.Pointer(uintptr(a1))), (*uint32)(unsafe.Pointer(&xLeft)), (*uint32)(unsafe.Pointer(&yTop)))
	if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 4))))&8 != 0 {
		if int32(*(*uint8)(unsafe.Pointer(uintptr(a2))))&2 != 0 {
			v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 52))))
			if *(*uint32)(unsafe.Pointer(uintptr(a2 + 36))) != 0x80000000 {
				nox_client_drawSetColor_434460(int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 36)))))
				nox_client_drawBorderLines_49CC70(xLeft, yTop, int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 8)))), int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 12)))))
			}
		}
	} else {
		v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 44))))
	}
	if uint32(v3) != 0x80000000 {
		nox_client_drawSetColor_434460(v3)
		nox_client_drawRectFilledOpaque_49CE30(xLeft+1, yTop+1, int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 8)))-2), int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 12)))-2))
	}
	v4 = int32(uint32(xLeft) + *(*uint32)(unsafe.Pointer(uintptr(a1 + 8)))/2)
	if uint32(v2) != 0x80000000 {
		nox_client_drawSetColor_434460(v2)
		nox_client_drawRectFilledOpaque_49CE30(v4-1, yTop+4, 3, int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 12)))-8))
	}
	return 1
}
func nox_xxx_wndScrollBoxDraw_4B5620(a1 *uint32, a2 int32) int32 {
	var (
		v2 int32
		v3 int32
	)
	nox_client_wndGetPosition_46AA60((*nox_window)(unsafe.Pointer(a1)), (*uint32)(unsafe.Pointer(&v2)), (*uint32)(unsafe.Pointer(&v3)))
	return 1
}
func nox_xxx_wndScrollBoxButtonCreate_4B5640(a1 int32, a2 int32, a3 int32) int32 {
	var (
		v3     int32
		v4     int32
		result int32
		v6     [332]byte
	)
	v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(a3 + 68))))
	*(*uint32)(unsafe.Pointer(&v6[0])) = 0
	libc.MemSet(unsafe.Pointer(&v6[4]), 0, 328)
	v4 = a2
	*(*uint32)(unsafe.Pointer(&v6[16])) = uint32(a1)
	*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v4))), 0)) = uint8(int8(a2&239 | 12))
	*(*uint32)(unsafe.Pointer(&v6[8])) = 1
	*(*uint32)(unsafe.Pointer(&v6[68])) = uint32(v3)
	if (v4 & 128) == 0 {
		*(*uint32)(unsafe.Pointer(&v6[52])) = *(*uint32)(unsafe.Pointer(uintptr(a3 + 52)))
		*(*uint32)(unsafe.Pointer(&v6[28])) = *(*uint32)(unsafe.Pointer(uintptr(a3 + 20)))
		*(*uint32)(unsafe.Pointer(&v6[20])) = *(*uint32)(unsafe.Pointer(uintptr(a3 + 28)))
	} else {
		*(*uint32)(unsafe.Pointer(&v6[32])) = *(*uint32)(unsafe.Pointer(uintptr(a3 + 32)))
		*(*uint32)(unsafe.Pointer(&v6[48])) = *(*uint32)(unsafe.Pointer(uintptr(a3 + 48)))
		*(*uint32)(unsafe.Pointer(&v6[56])) = *(*uint32)(unsafe.Pointer(uintptr(a3 + 56)))
		*(*uint32)(unsafe.Pointer(&v6[24])) = *(*uint32)(unsafe.Pointer(uintptr(a3 + 24)))
		*(*uint32)(unsafe.Pointer(&v6[40])) = *(*uint32)(unsafe.Pointer(uintptr(a3 + 40)))
	}
	if int32(*(*uint8)(unsafe.Pointer(uintptr(a3 + 8))))&16 != 0 {
		result = int32(uintptr(unsafe.Pointer(nox_gui_newButtonOrCheckbox_4A91A0((*nox_window)(unsafe.Pointer(uintptr(a1))), v4, 0, 0, 10, int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 12)))), (*nox_window_data)(unsafe.Pointer(&v6[0]))))))
	} else {
		result = int32(uintptr(unsafe.Pointer(nox_gui_newButtonOrCheckbox_4A91A0((*nox_window)(unsafe.Pointer(uintptr(a1))), v4, 0, 0, int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 8)))), 10, (*nox_window_data)(unsafe.Pointer(&v6[0]))))))
	}
	return result
}
func sub_4B5700(a1 int32, a2 int32, a3 int32, a4 int32, a5 int32, a6 int32) {
	var (
		v6 int32
		v7 int32
	)
	if a1 != 0 {
		v6 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 400))))
		v7 = int32(*(*uint32)(unsafe.Pointer(uintptr(v6 + 4))))
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v7))), 0)) = uint8(int8(v7 | 128))
		*(*uint32)(unsafe.Pointer(uintptr(v6 + 4))) = uint32(v7)
		nox_xxx_wndButtonInit_4A8340(int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 400)))))
		*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 400))) + 60))) = uint32(a2)
		*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 400))) + 68))) = uint32(a4)
		*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 400))) + 84))) = uint32(a3)
		*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 400))) + 92))) = uint32(a5)
		*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 400))) + 76))) = uint32(a6)
	}
}
func sub_4B5990() unsafe.Pointer {
	var (
		v0     int32
		v1     int32
		v2     *uint8
		v3     int32
		result unsafe.Pointer
		v5     *int16
		v6     [16]int16
	)
	v0 = 0
	v6[0] = 0
	libc.MemSet(unsafe.Pointer(&v6[1]), 0, 28)
	v6[15] = 0
	v1 = 0
	dword_5d4594_1312472 = 0
	if *memmap.PtrUint32(0x587000, 174360) == 0 {
		return *(*unsafe.Pointer)(unsafe.Pointer(&dword_5d4594_1312476))
	}
	v2 = (*uint8)(memmap.PtrOff(0x587000, 174360))
	for {
		v3 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), unsafe.Sizeof(uint32(0))*2))))
		v2 = (*uint8)(unsafe.Add(unsafe.Pointer(v2), 8))
		v1++
		if v3 == 0 {
			break
		}
	}
	dword_5d4594_1312472 = uint32(v1)
	if v1 <= 0 {
		return *(*unsafe.Pointer)(unsafe.Pointer(&dword_5d4594_1312476))
	}
	result = alloc.Calloc(int(v1), 4)
	dword_5d4594_1312476 = uint32(uintptr(result))
	v5 = &v6[0]
	if dword_5d4594_1312472 > 0 {
		for {
			*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1312476 + uint32(v0*4)))) = uint32(uintptr(unsafe.Pointer(sub_4B5A30_wol_locale((*libc.WChar)(unsafe.Pointer(v5))))))
			result = *(*unsafe.Pointer)(unsafe.Pointer(&dword_5d4594_1312476))
			v5 = *(**int16)(unsafe.Pointer(uintptr(dword_5d4594_1312476 + uint32(func() int32 {
				p := &v0
				x := *p
				*p++
				return x
			}()*4))))
			if v0 >= *(*int32)(unsafe.Pointer(&dword_5d4594_1312472)) {
				break
			}
		}
	}
	return result
}
func sub_4B5AB0(a1 int32, a2 int32, a3 *int32, a4 int32) int32 {
	var (
		v4 *int32
		v5 int32
		v6 *libc.WChar
		v7 int32
		v8 int32
		v9 int32
	)
	if a2 != 0x4007 {
		return 0
	}
	v4 = a3
	if (*nox_window)(unsafe.Pointer(a3)).ID() != 1985 {
		return 0
	}
	a2 = 0
	v5 = (*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1312484))))).Func94(asWindowEvent(0x4014, 0, 0))
	v6 = (*libc.WChar)(unsafe.Pointer(uintptr((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1312484))))).Func94(asWindowEvent(0x4016, v5, 0)))))
	v7 = sub_4B5B70_wol_locale(v6)
	v8 = sub_41FC40()
	sub_40DB50(v8+1, int32(uintptr(unsafe.Pointer(&a2))))
	if v7 != a2 {
		v9 = sub_41FC40()
		sub_40DB80(v9+1, v7)
	}
	(*nox_window)(unsafe.Pointer(uintptr(*memmap.PtrInt32(6112660, 1312488)))).Func94(asWindowEvent(0x4007, int32(uintptr(unsafe.Pointer(v4))), a4))
	sub_4B5BF0()
	return 0
}
func sub_4B5BF0() {
	nox_xxx_wnd_46C6E0((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1312480))))))
	nox_xxx_wndClearCaptureMain((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1312480))))))
	(*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1312480)))).Destroy()
	dword_5d4594_1312480 = 0
	*memmap.PtrUint32(6112660, 1312488) = 0
	dword_5d4594_1312484 = 0
	*(*unsafe.Pointer)(unsafe.Pointer(&dword_5d4594_1312476)) = nil
	dword_5d4594_1312472 = 0
}
func sub_4B5CD0() int32 {
	var (
		v0     int8
		v1     *uint8
		v2     int32
		v3     *uint8
		v4     int32
		v5     int32
		v6     int32
		v7     *uint8
		v8     int32
		v9     *uint8
		result int32
		v12    int32
	)
	v0 = 0
	v1 = (*uint8)(memmap.PtrOff(6112660, 1312500))
	for {
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v12))), 0)) = uint8(int8(int32(v0) / 63))
		*(*uint32)(unsafe.Pointer(v1)) = nox_color_rgb_4344A0(int32(uint8(int8(int32(v0)/63)))/3, int32(uint8(int8(int32(v0)/63)))*3/5, v12)
		v1 = (*uint8)(unsafe.Add(unsafe.Pointer(v1), 4))
		v0--
		if int32(uintptr(unsafe.Pointer(v1))) >= int32(uintptr(memmap.PtrOff(6112660, 1312756))) {
			break
		}
	}
	v2 = 0
	v3 = (*uint8)(memmap.PtrOff(6112660, 1312756))
	for {
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v12))), 0)) = uint8(int8(v2 / 32))
		v4 = int32(nox_color_rgb_4344A0(v12, v12, 0))
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v5))), 0)) = uint8(int8(int32(-1 - v12)))
		*(*uint32)(unsafe.Pointer(v3)) = uint32(v4)
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v3))), unsafe.Sizeof(uint32(0))*32))) = nox_color_rgb_4344A0(math.MaxUint8, v5, 0)
		v3 = (*uint8)(unsafe.Add(unsafe.Pointer(v3), 4))
		v2 += math.MaxUint8
		if int32(uintptr(unsafe.Pointer(v3))) >= int32(uintptr(memmap.PtrOff(6112660, 1312884))) {
			break
		}
	}
	v6 = 0
	v7 = (*uint8)(memmap.PtrOff(6112660, 1313012))
	for {
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v12))), 0)) = uint8(int8(v6 / 63))
		*(*uint32)(unsafe.Pointer(v7)) = nox_color_rgb_4344A0(v12, v12, v12)
		v7 = (*uint8)(unsafe.Add(unsafe.Pointer(v7), 4))
		v6 += math.MaxUint8
		if int32(uintptr(unsafe.Pointer(v7))) >= int32(uintptr(memmap.PtrOff(6112660, 1313268))) {
			break
		}
	}
	v8 = 0
	v9 = (*uint8)(memmap.PtrOff(6112660, 1313268))
	for {
		result = int32(nox_color_rgb_4344A0(v8/63, 50, 50))
		*(*uint32)(unsafe.Pointer(v9)) = uint32(result)
		v9 = (*uint8)(unsafe.Add(unsafe.Pointer(v9), 4))
		v8 += math.MaxUint8
		if int32(uintptr(unsafe.Pointer(v9))) >= int32(uintptr(memmap.PtrOff(6112660, 1313524))) {
			break
		}
	}
	return result
}
func sub_4B63B0(a1 *int2, a2 *int2) int32 {
	var (
		v2 int32
		v3 int32
	)
	nox_client_drawSetColor_434460(*memmap.PtrInt32(6112660, 1312492))
	nox_client_drawAddPoint_49F500(a1.field_0, a1.field_4)
	nox_client_drawAddPoint_49F500(a2.field_0, a2.field_4)
	nox_client_drawLineFromPoints_49E4B0()
	v2 = a2.field_0 - a1.field_0
	v3 = a2.field_4 - a1.field_4
	nox_client_drawSetColor_434460(*memmap.PtrInt32(6112660, 1312496))
	if v2 < 0 {
		v2 = -v2
	}
	if v3 < 0 {
		v3 = -v3
	}
	if v2 <= v3 {
		nox_client_drawAddPoint_49F500(a1.field_0-1, a1.field_4)
		nox_client_drawAddPoint_49F500(a2.field_0-1, a2.field_4)
		nox_client_drawLineFromPoints_49E4B0()
		nox_client_drawAddPoint_49F500(a1.field_0+1, a1.field_4)
		nox_client_drawAddPoint_49F500(a2.field_0+1, a2.field_4)
	} else {
		nox_client_drawAddPoint_49F500(a1.field_0, a1.field_4-1)
		nox_client_drawAddPoint_49F500(a2.field_0, a2.field_4-1)
		nox_client_drawLineFromPoints_49E4B0()
		nox_client_drawAddPoint_49F500(a1.field_0, a1.field_4+1)
		nox_client_drawAddPoint_49F500(a2.field_0, a2.field_4+1)
	}
	return nox_client_drawLineFromPoints_49E4B0()
}
func sub_4B64C0() int32 {
	var (
		v0 uint8
		v1 *uint8
		v2 int32
		v3 int32
		v5 int32
		v6 int32
		v7 int32
		v8 int32
	)
	*memmap.PtrUint32(6112660, 1313524) = nox_color_rgb_4344A0(math.MaxUint8, math.MaxUint8, 0)
	*memmap.PtrUint32(6112660, 1313528) = nox_color_rgb_4344A0(math.MaxUint8, 100, 0)
	dword_5d4594_1313532 = nox_color_rgb_4344A0(math.MaxUint8, math.MaxUint8, 0)
	dword_5d4594_1313536 = nox_color_rgb_4344A0(0, 0, math.MaxUint8)
	dword_5d4594_1313540 = nox_color_rgb_4344A0(0, 200, math.MaxUint8)
	*memmap.PtrUint32(6112660, 1313544) = nox_color_rgb_4344A0(0, 200, 200)
	*memmap.PtrUint32(6112660, 1313548) = nox_color_rgb_4344A0(50, math.MaxUint8, math.MaxUint8)
	*memmap.PtrUint32(6112660, 1313552) = nox_color_rgb_4344A0(math.MaxUint8, 0, math.MaxUint8)
	*memmap.PtrUint32(6112660, 1313556) = nox_color_rgb_4344A0(math.MaxUint8, 200, math.MaxUint8)
	*memmap.PtrUint32(6112660, 1313560) = nox_color_rgb_4344A0(math.MaxUint8, 200, 0)
	dword_5d4594_1313564 = nox_color_rgb_4344A0(math.MaxUint8, math.MaxUint8, 100)
	*memmap.PtrUint32(6112660, 1313568) = nox_color_rgb_4344A0(100, math.MaxUint8, 50)
	*memmap.PtrUint32(6112660, 1313572) = nox_color_rgb_4344A0(150, math.MaxUint8, 150)
	*memmap.PtrUint32(6112660, 1313576) = nox_color_rgb_4344A0(math.MaxUint8, math.MaxUint8, 0)
	*memmap.PtrUint32(6112660, 1313580) = nox_color_rgb_4344A0(0, 220, 0)
	*memmap.PtrUint32(6112660, 1313584) = nox_color_rgb_4344A0(150, math.MaxUint8, 150)
	*memmap.PtrUint32(6112660, 1313588) = nox_color_rgb_4344A0(200, 200, 200)
	*memmap.PtrUint32(6112660, 1313592) = nox_color_rgb_4344A0(math.MaxUint8, math.MaxUint8, math.MaxUint8)
	*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v8))), 0)) = math.MaxUint8
	*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v7))), 0)) = math.MaxUint8
	*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v6))), 0)) = math.MaxUint8
	v0 = 0
	v1 = (*uint8)(memmap.PtrOff(6112660, 1313656))
	v2 = 1085
	v5 = 600
	v3 = 765
	for {
		if int32(v0) > 3 {
			if int32(v0) > 7 {
				*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v8))), 0)) = uint8(int8(v2/9 - 1))
			} else {
				*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v7))), 0)) = uint8(int8(v5/4 - 1))
			}
		} else {
			*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v6))), 0)) = uint8(int8(v3 / 3))
		}
		*(*uint32)(unsafe.Pointer(v1)) = nox_color_rgb_4344A0(v8, v7, v6)
		v0++
		v3 -= math.MaxUint8
		v2 -= 155
		v1 = (*uint8)(unsafe.Add(unsafe.Pointer(v1), -4))
		v5 -= 200
		if int32(v0) >= 16 {
			break
		}
	}
	return sub_4B5CD0()
}
func sub_4B6880(a1 *uint32, a2 int32, a3 int32, a4 int32) int32 {
	var (
		v4     int32
		v5     int32
		v6     int32
		result int32
		v8     int32
		v9     int32
		v10    int32
		v11    int32
		v12    bool
		v13    int32
		xLeft  int2
		v15    int32
	)
	v4 = a2
	v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 448))) - *(*uint32)(unsafe.Pointer(uintptr(a2 + 444))))
	v15 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 448))) - nox_frame_xxx_2598000)
	v6 = v15
	if v15 == v5 {
		v6 = func() int32 {
			p := &v15
			*p--
			return *p
		}()
	}
	if v6 > 0 {
		v8 = int32(*a1 + *(*uint32)(unsafe.Pointer(uintptr(v4 + 12))) - *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*4)))
		v9 = int32(*(*uint32)(unsafe.Pointer(uintptr(v4 + 16))) - uint32(*(*int16)(unsafe.Pointer(uintptr(v4 + 106)))) - uint32(*(*int16)(unsafe.Pointer(uintptr(v4 + 104)))) - *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*5)))
		v10 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*1)))
		v11 = v10 + v9
		v12 = uint32(v8-10) < *a1
		xLeft.field_0 = v8
		xLeft.field_4 = v11
		if !v12 && v11-10 >= v10 && uint32(v8+10) < *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*2)) && uint32(v11+10) < *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*3)) {
			v13 = v15 * 4 / v5
			sub_4B6720(&xLeft, a4, v13*2+1, int8(v15*5/v5))
			nox_client_drawSetColor_434460(a3)
			nox_xxx_drawPointMB_499B70(xLeft.field_0, xLeft.field_4, v13)
		}
		result = 1
	} else {
		nox_xxx_spriteDeleteStatic_45A4E0_drawable((*nox_drawable)(unsafe.Pointer(uintptr(v4))))
		result = 0
	}
	return result
}
func sub_4B6970(a1 *uint32, dr *nox_drawable, a3 int32, a4 int32) int32 {
	var (
		v4 int32
		v5 int32
		a2 int32 = int32(uintptr(unsafe.Pointer(dr)))
	)
	v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 440))))
	*(*uint32)(unsafe.Pointer(uintptr(a2 + 432))) += uint32(v4 * *memmap.PtrInt32(0x587000, uint32(int32(*(*uint8)(unsafe.Pointer(uintptr(a2 + 299))))*8)+0x2EE58))
	v5 = int32(uint32(v4**memmap.PtrInt32(0x587000, uint32(int32(*(*uint8)(unsafe.Pointer(uintptr(a2 + 299))))*8)+0x2EE5C)) + *(*uint32)(unsafe.Pointer(uintptr(a2 + 436))))
	*(*uint32)(unsafe.Pointer(uintptr(a2 + 436))) = uint32(v5)
	nox_xxx_updateSpritePosition_49AA90((*nox_drawable)(unsafe.Pointer(uintptr(a2))), int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 432)))>>12), v5>>12)
	sub_4B69F0(a2)
	return sub_4B6880(a1, a2, a3, a4)
}
func sub_4B69F0(a1 int32) int16 {
	var (
		v1     int8
		result int16
		v3     int32
	)
	v1 = int8(*(*uint8)(unsafe.Pointer(uintptr(a1 + 296))))
	*(*uint16)(unsafe.Pointer(uintptr(a1 + 104))) += uint16(v1)
	result = int16(*(*uint16)(unsafe.Pointer(uintptr(a1 + 104))))
	if int32(result) >= 0 {
		if int32(uint8(nox_frame_xxx_2598000))&1 != 0 {
			*(*uint8)(unsafe.Pointer(uintptr(a1 + 296))) = uint8(int8(int32(v1) - 1))
		}
	} else {
		*(*uint16)(unsafe.Pointer(uintptr(a1 + 104))) = uint16(int16(int32(-result)))
		result = int16(int32(v1) * 0x6661)
		v3 = int32(v1) * (-9) / 10
		*(*uint8)(unsafe.Pointer(uintptr(a1 + 296))) = uint8(int8(v3))
		if int32(int8(v3)) < 2 {
			*(*uint16)(unsafe.Pointer(uintptr(a1 + 104))) = 0
			*(*uint8)(unsafe.Pointer(uintptr(a1 + 296))) = 0
		}
	}
	return result
}
func sub_4B6B80(a1 *int32, dr *nox_drawable, a3 int32) int32 {
	var (
		v3    int32
		v4    int32
		v5    int32
		v6    int32
		v7    int32
		v8    int32
		v10   int32
		v11   int32
		v12   int32
		v13   int32
		v14   int32
		v15   int32
		v16   int8
		v17   int8
		v18   int8
		v19   int32
		xLeft int2
		v21   int32
		a2    int32 = int32(uintptr(unsafe.Pointer(dr)))
	)
	v3 = a2
	if *memmap.PtrUint32(6112660, 1313660) == 0 {
		*memmap.PtrUint32(6112660, 1313660) = uint32(nox_xxx_getTTByNameSpriteMB_44CFC0("DrainManaOrb"))
		*memmap.PtrUint32(6112660, 1313664) = uint32(nox_xxx_getTTByNameSpriteMB_44CFC0("HealOrb"))
		*memmap.PtrUint32(6112660, 1313668) = uint32(nox_xxx_getTTByNameSpriteMB_44CFC0("CharmOrb"))
		*memmap.PtrUint32(6112660, 1313672) = uint32(nox_xxx_getTTByNameSpriteMB_44CFC0("WhiteOrb"))
		*memmap.PtrUint32(6112660, 1313676) = uint32(nox_xxx_getTTByNameSpriteMB_44CFC0("ManaBombOrb"))
		*memmap.PtrUint32(6112660, 1313680) = uint32(nox_xxx_getTTByNameSpriteMB_44CFC0("WhiteMoveOrb"))
		*memmap.PtrUint32(6112660, 1313684) = uint32(nox_xxx_getTTByNameSpriteMB_44CFC0("BlueMoveOrb"))
	}
	v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 108))))
	if uint32(v4) == *memmap.PtrUint32(6112660, 1313660) || uint32(v4) == *memmap.PtrUint32(6112660, 1313684) {
		v5 = int32(dword_5d4594_1313540)
		v21 = int32(dword_5d4594_1313536)
		goto LABEL_13
	}
	if uint32(v4) == *memmap.PtrUint32(6112660, 1313668) {
		v5 = int32(*memmap.PtrUint32(6112660, 1313584))
		v21 = int32(*memmap.PtrUint32(6112660, 1313580))
	LABEL_13:
		v19 = v5
		goto LABEL_14
	}
	if uint32(v4) == *memmap.PtrUint32(6112660, 1313672) || uint32(v4) == *memmap.PtrUint32(6112660, 1313676) || uint32(v4) == *memmap.PtrUint32(6112660, 1313680) {
		v21 = int32(*memmap.PtrUint32(6112660, 1313588))
		v19 = int32(*memmap.PtrUint32(6112660, 1313592))
	} else {
		v21 = int32(*memmap.PtrUint32(6112660, 1313528))
		v19 = int32(dword_5d4594_1313532)
	}
LABEL_14:
	if a3 != 0 {
		v6 = int32(uint32(*(*uint16)(unsafe.Pointer(uintptr(v3 + 432)))) - *(*uint32)(unsafe.Pointer(uintptr(v3 + 12))))
		v7 = int32(uint32(*(*uint16)(unsafe.Pointer(uintptr(v3 + 434)))) - *(*uint32)(unsafe.Pointer(uintptr(v3 + 16))))
		v8 = int32(sub_48C6B0(v6, v7))
		if v8+1 <= 10 {
			nox_xxx_spriteDeleteStatic_45A4E0_drawable((*nox_drawable)(unsafe.Pointer(uintptr(v3))))
			return 0
		}
		nox_xxx_updateSpritePosition_49AA90((*nox_drawable)(unsafe.Pointer(uintptr(v3))), int32(*(*uint32)(unsafe.Pointer(uintptr(v3 + 12)))+uint32(v6*int32(*(*uint8)(unsafe.Pointer(uintptr(v3 + 443))))/(v8+1))), int32(*(*uint32)(unsafe.Pointer(uintptr(v3 + 16)))+uint32(v7*int32(*(*uint8)(unsafe.Pointer(uintptr(v3 + 443))))/(v8+1))))
	}
	v10 = *a1
	v11 = *(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*1))
	v12 = int32(*(*uint32)(unsafe.Pointer(uintptr(v3 + 16))) - uint32(*(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*5))))
	xLeft.field_0 = int32(*(*uint32)(unsafe.Pointer(uintptr(v3 + 12))) + uint32(*a1) - uint32(*(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*4))))
	xLeft.field_4 = v12 + v11 - 22
	v13 = int32(*(*uint8)(unsafe.Pointer(uintptr(v3 + 444))))
	if xLeft.field_0-v13 < v10 {
		return 1
	}
	if xLeft.field_4-v13 < v11 {
		return 1
	}
	if v13+xLeft.field_0 >= *(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*2)) {
		return 1
	}
	if v13+xLeft.field_4 >= *(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*3)) {
		return 1
	}
	sub_4B6720(&xLeft, v21, v13, 5)
	nox_client_drawSetColor_434460(v19)
	nox_xxx_drawPointMB_499B70(xLeft.field_0, xLeft.field_4, int32(*(*uint8)(unsafe.Pointer(uintptr(v3 + 444))))>>1)
	v14 = int32(*(*uint32)(unsafe.Pointer(uintptr(v3 + 32))) - *(*uint32)(unsafe.Pointer(uintptr(v3 + 12))))
	v15 = int32(*(*uint32)(unsafe.Pointer(uintptr(v3 + 36))) - *(*uint32)(unsafe.Pointer(uintptr(v3 + 16))))
	nox_client_drawSetColor_434460(v19)
	nox_client_drawAddPoint_49F500(xLeft.field_0, xLeft.field_4)
	nox_xxx_rasterPointRel_49F570(v14, v15)
	nox_client_drawLineFromPoints_49E4B0()
	v16 = int8(*(*uint8)(unsafe.Pointer(uintptr(v3 + 445))))
	if int32(v16) == 0 {
		return 1
	}
	v17 = int8(int32(*(*uint8)(unsafe.Pointer(uintptr(v3 + 446)))) - 1)
	*(*uint8)(unsafe.Pointer(uintptr(v3 + 446))) = uint8(v17)
	if int32(v17) != 0 {
		return 1
	}
	*(*uint8)(unsafe.Pointer(uintptr(v3 + 446))) = uint8(v16)
	v18 = int8(int32(*(*uint8)(unsafe.Pointer(uintptr(v3 + 444)))) - 1)
	*(*uint8)(unsafe.Pointer(uintptr(v3 + 444))) = uint8(v18)
	if int32(v18) != 0 {
		return 1
	}
	nox_xxx_spriteDeleteStatic_45A4E0_drawable((*nox_drawable)(unsafe.Pointer(uintptr(v3))))
	return 0
}
func sub_4B71A0(a1 *uint32, a2 int32) int32 {
	var (
		v2     int32
		v3     int32
		v4     int32
		v5     int32
		v6     int32
		result int32
		v8     int32
		v9     int32
		v10    int32
		v11    int2
		v12    int32
	)
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 16))))
	v3 = int32(*(*int16)(unsafe.Pointer(uintptr(a2 + 106))))
	v11.field_0 = int32(*a1 + *(*uint32)(unsafe.Pointer(uintptr(a2 + 12))) - *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*4)))
	v11.field_4 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*1)) + uint32(v2) - uint32(v3) - uint32(*(*int16)(unsafe.Pointer(uintptr(a2 + 104)))) - *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*5)))
	*(*uint32)(unsafe.Pointer(uintptr(a2 + 432))) += *(*uint32)(unsafe.Pointer(uintptr(a2 + 440))) * *memmap.PtrUint32(0x587000, uint32(int32(*(*uint8)(unsafe.Pointer(uintptr(a2 + 299))))*8)+0x2EE58)
	v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 440)))**memmap.PtrUint32(0x587000, uint32(int32(*(*uint8)(unsafe.Pointer(uintptr(a2 + 299))))*8)+0x2EE5C) + *(*uint32)(unsafe.Pointer(uintptr(a2 + 436))))
	*(*uint32)(unsafe.Pointer(uintptr(a2 + 436))) = uint32(v4)
	nox_xxx_updateSpritePosition_49AA90((*nox_drawable)(unsafe.Pointer(uintptr(a2))), int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 432)))>>12), v4>>12)
	v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 444))) - *(*uint32)(unsafe.Pointer(uintptr(a2 + 448))))
	v6 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 444))) - nox_frame_xxx_2598000)
	if v6 == v5 {
		v6--
	}
	if v6 > 0 {
		if uint32(v11.field_0-10) >= *a1 && uint32(v11.field_4-10) >= *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*1)) && uint32(v11.field_0+10) < *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*2)) && uint32(v11.field_4+10) < *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*3)) {
			sub_4B69F0(a2)
			v8 = (v6 * 16 / v5) * 4
			v12 = int32(*memmap.PtrUint32(6112660, uint32(v8)+0x140B3C))
			v9 = v8 / 16
			v10 = (v8 / 16) >> 1
			sub_4B6720(&v11, int32(*memmap.PtrUint32(6112660, uint32(v8)+0x140B3C)), (v8/16)*2+1, int8(v6*5/v5))
			nox_client_drawSetColor_434460(v12)
			nox_client_drawRectFilledOpaque_49CE30(v11.field_0-v10, v11.field_4-v10, v9, v9)
		}
		result = 1
	} else {
		nox_xxx_spriteDeleteStatic_45A4E0_drawable((*nox_drawable)(unsafe.Pointer(uintptr(a2))))
		result = 0
	}
	return result
}
func nox_xxx_netHandleSummonPacket_4B7C40(a1 int16, a2 *uint16, a3 uint16, a4 uint8, a5 int16) *uint32 {
	var (
		v5     int32
		result *uint32
		v7     *uint32
		v8     *uint32
		v9     int32
		v10    int32
		v11    int32
	)
	v10 = int32(*(*uint16)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint16(0))*1)))
	v9 = int32(*a2)
	v5 = nox_xxx_getTTByNameSpriteMB_44CFC0("SummonEffect")
	result = &nox_xxx_spriteLoadAdd_45A360_drawable(v5, v9, v10).field_0
	v7 = result
	if result != nil {
		result = &nox_new_drawable_for_thing(int32(a3)).field_0
		v8 = result
		if result != nil {
			*(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*3)) = uint32(*a2)
			*(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*4)) = uint32(*(*uint16)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint16(0))*1)))
			*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(result))), 297))) = uint8(int8(nox_xxx_math_509EA0(int32(a4))))
			*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v11))), unsafe.Sizeof(uint16(0))*1)) = uint16(a1)
			*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v11))), unsafe.Sizeof(uint16(0))*0)) = uint16(a5)
			*(*uint32)(unsafe.Add(unsafe.Pointer(v8), unsafe.Sizeof(uint32(0))*69)) = 8
			*(*uint32)(unsafe.Add(unsafe.Pointer(v7), unsafe.Sizeof(uint32(0))*108)) = uint32(uintptr(unsafe.Pointer(v8)))
			*(*uint32)(unsafe.Add(unsafe.Pointer(v7), unsafe.Sizeof(uint32(0))*109)) = uint32(v11)
			result = (*uint32)(unsafe.Pointer(uintptr(nox_frame_xxx_2598000)))
			*(*uint32)(unsafe.Add(unsafe.Pointer(v7), unsafe.Sizeof(uint32(0))*79)) = nox_frame_xxx_2598000
		}
	}
	return result
}
func sub_4B7EE0(a1 int16) {
	if *memmap.PtrUint32(6112660, 1313744) == 0 {
		*memmap.PtrUint32(6112660, 1313744) = uint32(nox_xxx_getTTByNameSpriteMB_44CFC0("SummonEffect"))
	}
	if dword_5d4594_1313740 == 0 {
		dword_5d4594_1313740 = uint32(nox_xxx_getTTByNameSpriteMB_44CFC0("BlueSpark"))
	}
	var v2 int32 = sub_45A060()
	if v2 == 0 {
		return
	}
	for *(*uint32)(unsafe.Pointer(uintptr(v2 + 108))) != *memmap.PtrUint32(6112660, 1313744) || int32(*(*uint16)(unsafe.Pointer(uintptr(v2 + 438)))) != int32(a1) {
		v2 = sub_45A070(v2)
		if v2 == 0 {
			return
		}
	}
	nox_xxx_makePointFxCli_499610(*(*int32)(unsafe.Pointer(&dword_5d4594_1313740)), 50, 1000, 30, int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 12)))), int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 16)))))
	nox_xxx_spriteDelete_45A4B0((*nox_drawable)(unsafe.Pointer(*(**uint64)(unsafe.Pointer(uintptr(v2 + 432))))))
	nox_xxx_spriteDeleteStatic_45A4E0_drawable((*nox_drawable)(unsafe.Pointer(uintptr(v2))))
}
func nox_xxx_spriteShieldLoad_4B7F90() int32 {
	var result int32
	*memmap.PtrUint32(6112660, 1313748) = uint32(nox_xxx_getTTByNameSpriteMB_44CFC0("SphericalShieldNW"))
	*memmap.PtrUint32(6112660, 1313752) = uint32(nox_xxx_getTTByNameSpriteMB_44CFC0("SphericalShieldN"))
	*memmap.PtrUint32(6112660, 1313756) = uint32(nox_xxx_getTTByNameSpriteMB_44CFC0("SphericalShieldNE"))
	*memmap.PtrUint32(6112660, 1313760) = uint32(nox_xxx_getTTByNameSpriteMB_44CFC0("SphericalShieldW"))
	*memmap.PtrUint32(6112660, 1313768) = uint32(nox_xxx_getTTByNameSpriteMB_44CFC0("SphericalShieldE"))
	*memmap.PtrUint32(6112660, 1313772) = uint32(nox_xxx_getTTByNameSpriteMB_44CFC0("SphericalShieldSW"))
	*memmap.PtrUint32(6112660, 1313776) = uint32(nox_xxx_getTTByNameSpriteMB_44CFC0("SphericalShieldS"))
	result = nox_xxx_getTTByNameSpriteMB_44CFC0("SphericalShieldSE")
	*memmap.PtrUint32(6112660, 1313780) = uint32(result)
	*memmap.PtrUint32(6112660, 1313764) = 0
	*memmap.PtrUint32(6112660, 1313784) = 1
	return result
}
func nox_xxx_fxShield_4B8090(a1 uint32, a2 int32) *uint32 {
	var (
		v2     int32
		v3     int32
		result *uint32
		v5     int32
		v6     *uint32
		v7     int4
	)
	if *memmap.PtrUint32(6112660, 1313784) == 0 {
		nox_xxx_spriteShieldLoad_4B7F90()
	}
	v2 = a2
	switch a2 {
	case 0:
		nox_xxx_getTTByNameSpriteMB_44CFC0("SphericalShieldNW")
	case 1:
		nox_xxx_getTTByNameSpriteMB_44CFC0("ShpericalShieldN")
	case 2:
		nox_xxx_getTTByNameSpriteMB_44CFC0("SphericalShieldNW")
	case 3:
		nox_xxx_getTTByNameSpriteMB_44CFC0("SphericalShieldW")
	case 5:
		nox_xxx_getTTByNameSpriteMB_44CFC0("SphericalShieldE")
	case 6:
		nox_xxx_getTTByNameSpriteMB_44CFC0("SphericalShieldSW")
	case 7:
		nox_xxx_getTTByNameSpriteMB_44CFC0("SphericalShieldS")
	case 8:
		nox_xxx_getTTByNameSpriteMB_44CFC0("SphericalShieldSE")
	default:
	}
	if nox_xxx_netTestHighBit_578B70(a1) != 0 {
		v3 = nox_xxx_netClearHighBit_578B30(int16(uint16(a1)))
		result = nox_xxx_netSpriteByCodeStatic_45A720(v3)
	} else {
		v5 = nox_xxx_netClearHighBit_578B30(int16(uint16(a1)))
		result = nox_xxx_netSpriteByCodeDynamic_45A6F0(v5)
	}
	v6 = result
	if result != nil {
		v7.field_0 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*3)) - 10)
		v7.field_4 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*4)) - 10)
		v7.field_8 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*3)) + 10)
		v7.field_C = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*4)) + 10)
		dword_5d4594_1313788 = 0
		nox_xxx_forEachSprite_49AB00(&v7, unsafe.Pointer(cgoFuncAddr(nox_xxx_spriteScanForShield_4B81E0)), int32(uintptr(unsafe.Pointer(&a1))))
		result = *(**uint32)(unsafe.Pointer(&dword_5d4594_1313788))
		if dword_5d4594_1313788 != 1 {
			result = &nox_xxx_spriteLoadAdd_45A360_drawable(int32(*memmap.PtrUint32(6112660, uint32(v2*4)+0x140BD4)), int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*3))), int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*4))+3)).field_0
			if result != nil {
				*(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*108)) = a1
			}
		}
	}
	return result
}
func nox_xxx_spriteScanForShield_4B81E0(a1 int32, a2 int32) {
	var v2 *uint8
	v2 = (*uint8)(memmap.PtrOff(6112660, 1313748))
	for {
		if *(*uint32)(unsafe.Pointer(uintptr(a1 + 108))) == *(*uint32)(unsafe.Pointer(v2)) && *(*uint32)(unsafe.Pointer(uintptr(a1 + 432))) == *(*uint32)(unsafe.Pointer(uintptr(a2))) {
			dword_5d4594_1313788 = 1
		}
		v2 = (*uint8)(unsafe.Add(unsafe.Pointer(v2), 4))
		if int32(uintptr(unsafe.Pointer(v2))) >= int32(uintptr(memmap.PtrOff(6112660, 1313784))) {
			break
		}
	}
}
func nox_xxx_gameDeleteSpiningCrownSkull_4B8220() int32 {
	var result int32
	if dword_5d4594_1313792 != 0 {
		nox_xxx_spriteDelete_45A4B0((*nox_drawable)(unsafe.Pointer(*(**uint64)(unsafe.Pointer(&dword_5d4594_1313792)))))
	}
	dword_5d4594_1313792 = 0
	if dword_5d4594_1313796 != 0 {
		nox_xxx_spriteDelete_45A4B0((*nox_drawable)(unsafe.Pointer(*(**uint64)(unsafe.Pointer(&dword_5d4594_1313796)))))
	}
	result = int32(dword_5d4594_1313800)
	if dword_5d4594_1313800 != 0 {
		result = nox_xxx_spriteDelete_45A4B0((*nox_drawable)(unsafe.Pointer(*(**uint64)(unsafe.Pointer(&dword_5d4594_1313800)))))
	}
	dword_5d4594_1313796 = 0
	dword_5d4594_1313800 = 0
	return result
}
func sub_4B8960(a1 *int32, dr *nox_drawable, a3 int32, a4 *uint32, a5 int32, a6 int32) int16 {
	var (
		a2  int32 = int32(uintptr(unsafe.Pointer(dr)))
		v6  *uint8
		v7  int32
		v8  int32
		v9  int8
		v10 int32
		v11 int32
		i   int32
		v13 int32
		v14 int32
		v15 int32
		v16 uint8
		v17 int32
		j   int32
		v19 int32
		v20 int32
		k   int32
		v22 int32
		v23 int32
		v24 int32
		l   int32
		v27 int32
		v28 int32
		v29 *int32
		v30 int32
		v31 int32
	)
	v6 = (*uint8)(unsafe.Pointer(dr))
	v28 = bool2int(nox_xxx_spriteCheckFlag31_4356C0(dr, 23))
	v7 = 0
	v30 = 0
	v31 = 0
	if int32(uint8(nox_frame_xxx_2598000))&1 != 0 {
		v27 = int32(nox_color_white_2523948)
	} else {
		v27 = int32(nox_color_blue_2650684)
	}
	v8 = a3
	if a3&2 != 0 {
		v9 = int8(*(*uint8)(unsafe.Pointer(uintptr(a2 + 297))))
		if int32(v9) == 6 || int32(v9) == 7 || int32(v9) == 8 {
			v10 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 276))))
			if v10 != 6 && v10 != 35 && v10 != 39 && v10 != 40 {
				for ((1 << v7) & 2) == 0 {
					if func() int32 {
						p := &v7
						*p++
						return *p
					}() >= 26 {
						goto LABEL_22
					}
				}
				v11 = int32(*(*uint32)(unsafe.Pointer(uintptr(a5 + v7*4 + 52))))
				if v11 != 0 {
					if v28 != 0 {
						for i = 1; i < 7; i++ {
							nox_xxx_drawPlayer_4341D0(i, v27)
						}
						v6 = (*uint8)(unsafe.Pointer(dr))
					} else {
						sub_4B8CA0(a4, (*byte)(unsafe.Pointer(uintptr(2))))
					}
					nox_xxx_drawObject_4C4770_draw(a1, (*nox_drawable)(unsafe.Pointer(v6)), int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v11 + int32(*(*uint8)(unsafe.Add(unsafe.Pointer(v6), 297)))*4 + 4))) + uint32(a6*4))))))
					v8 = a3
					v30 = 1
				}
			}
		}
	}
LABEL_22:
	v13 = 0
	v29 = (*int32)(unsafe.Pointer(uintptr(a5 + 52)))
	for {
		v14 = 1 << v13
		if ((1 << v13) & 2) == 0 {
			v15 = int32(uint32(v14) & 0x3000000)
			if (uint32(v14)&0x3000000) == 0 || (func() bool {
				v16 = *(*uint8)(unsafe.Add(unsafe.Pointer(v6), 297))
				return int32(v16) != 3
			}()) && int32(v16) != 0 && int32(v16) != 6 && int32(v16) != 7 {
				if v14&v8 != 0 && ((v14&3072) == 0 || (v8&0x4000) == 0) {
					v17 = *v29
					if *v29 != 0 {
						if v28 != 0 {
							for j = 1; j < 7; j++ {
								nox_xxx_drawPlayer_4341D0(j, v27)
							}
							v6 = (*uint8)(unsafe.Pointer(dr))
						} else {
							sub_4B8CA0(a4, (*byte)(unsafe.Pointer(uintptr(1<<v13))))
						}
						nox_xxx_drawObject_4C4770_draw(a1, (*nox_drawable)(unsafe.Pointer(v6)), int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v17 + int32(*(*uint8)(unsafe.Add(unsafe.Pointer(v6), 297)))*4 + 4))) + uint32(a6*4))))))
						if v15 != 0 {
							v31 = 1
						}
						v8 = a3
					}
				}
			}
		}
		v13++
		v29 = (*int32)(unsafe.Add(unsafe.Pointer(v29), unsafe.Sizeof(int32(0))*1))
		if v13 >= 26 {
			break
		}
	}
	if a3&2 != 0 && v30 == 0 {
		v19 = 0
		for ((1 << v19) & 2) == 0 {
			if func() int32 {
				p := &v19
				*p++
				return *p
			}() >= 26 {
				goto LABEL_55
			}
		}
		v20 = int32(*(*uint32)(unsafe.Pointer(uintptr(a5 + v19*4 + 52))))
		if v20 != 0 {
			if v28 != 0 {
				for k = 1; k < 7; k++ {
					nox_xxx_drawPlayer_4341D0(k, v27)
				}
				v6 = (*uint8)(unsafe.Pointer(dr))
			} else {
				sub_4B8CA0(a4, (*byte)(unsafe.Pointer(uintptr(2))))
			}
			nox_xxx_drawObject_4C4770_draw(a1, (*nox_drawable)(unsafe.Pointer(v6)), int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v20 + int32(*(*uint8)(unsafe.Add(unsafe.Pointer(v6), 297)))*4 + 4))) + uint32(a6*4))))))
			v8 = a3
		}
	}
LABEL_55:
	*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v22))), unsafe.Sizeof(uint16(0))*0)) = uint16(int16(v31))
	if v31 == 0 && uint32(v8)&0x3000000 != 0 {
		v23 = 0
		if (uint32(v8) & 0x1000000) != 0 {
			v22 = 0x1000000
		} else {
			v22 = 0x2000000
		}
		for ((1 << v23) & v22) == 0 {
			if func() int32 {
				p := &v23
				*p++
				return *p
			}() >= 26 {
				return int16(v22)
			}
		}
		v24 = int32(*(*uint32)(unsafe.Pointer(uintptr(a5 + v23*4 + 52))))
		if v24 != 0 {
			if v28 != 0 {
				for l = 1; l < 7; l++ {
					nox_xxx_drawPlayer_4341D0(l, v27)
				}
				v6 = (*uint8)(unsafe.Pointer(dr))
			} else {
				sub_4B8CA0(a4, (*byte)(unsafe.Pointer(uintptr(v22))))
			}
			*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v22))), unsafe.Sizeof(uint16(0))*0)) = uint16(nox_xxx_drawObject_4C4770_draw(a1, (*nox_drawable)(unsafe.Pointer(v6)), int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v24 + int32(*(*uint8)(unsafe.Add(unsafe.Pointer(v6), 297)))*4 + 4))) + uint32(a6*4)))))))
		}
	}
	return int16(v22)
}
func sub_4B8CA0(a1 *uint32, a2 *byte) *uint32 {
	var (
		result *uint32
		v3     int32
		v4     int32
		v5     int32
		v6     int32
		v7     *uint32
		v8     int32
		v9     *uint8
		v10    *uint32
		v11    int32
		v12    *int32
		v13    int32
		v14    **uint32
		v15    int32
	)
	result = a1
	v3 = 0
	for (*byte)(unsafe.Pointer(uintptr(*result))) != a2 {
		v3++
		result = (*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*6))
		if v3 >= 26 {
			return result
		}
	}
	v4 = sub_415CD0(a2)
	result = nox_xxx_equipClothFindDefByTT_413270(v4)
	v7 = result
	if result != nil {
		v8 = 1
		v9 = (*uint8)(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*4))))
		for {
			*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&result))), 0)) = *(*uint8)(unsafe.Add(unsafe.Pointer(v9), 1))
			*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v6))), 0)) = *v9
			*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v5))), 0)) = *((*uint8)(unsafe.Add(unsafe.Pointer(v9), -1)))
			sub_4340A0(func() int32 {
				p := &v8
				x := *p
				*p++
				return x
			}(), v5, v6, int32(uintptr(unsafe.Pointer(result))))
			v9 = (*uint8)(unsafe.Add(unsafe.Pointer(v9), 3))
			if v8 >= 7 {
				break
			}
		}
		v10 = a1
		v11 = v3 * 3
		v12 = (*int32)(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer(v7), unsafe.Sizeof(uint32(0))*9))))
		v13 = 4
		v14 = (**uint32)(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*uintptr(v11*2+1)))))
		for {
			result = *v14
			if *v14 != nil {
				*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v5))), 0)) = *((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(result))), 26)))
				*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v10))), 0)) = *((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(result))), 25)))
				v15 = v5
				*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v5))), 0)) = *((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(result))), 24)))
				sub_4340A0(*v12, v5, int32(uintptr(unsafe.Pointer(v10))), v15)
			}
			v14 = (**uint32)(unsafe.Add(unsafe.Pointer(v14), unsafe.Sizeof((*uint32)(nil))*1))
			v12 = (*int32)(unsafe.Add(unsafe.Pointer(v12), unsafe.Sizeof(int32(0))*1))
			v13--
			if v13 == 0 {
				break
			}
		}
	}
	return result
}
func sub_4B8D40(a1 *int32, dr *nox_drawable, a3 int32, a4 *uint32, a5 int32, a6 int32) int16 {
	var (
		v6  *uint8
		v7  int32
		v8  *int32
		v9  int32
		v10 int32
		i   int32
		v13 int32
		v14 int32
	)
	v6 = (*uint8)(unsafe.Pointer(dr))
	v14 = bool2int(nox_xxx_spriteCheckFlag31_4356C0(dr, 25))
	if int32(uint8(nox_frame_xxx_2598000))&1 != 0 {
		v13 = int32(nox_color_white_2523948)
	} else {
		v13 = int32(nox_color_blue_2650684)
	}
	v7 = 1
	v8 = (*int32)(unsafe.Pointer(uintptr(a5 + 160)))
	for {
		v9 = 1 << v7
		if (1<<v7)&a3 != 0 {
			v10 = *v8
			if *v8 != 0 {
				if v14 != 0 {
					for i = 1; i < 7; i++ {
						nox_xxx_drawPlayer_4341D0(i, v13)
					}
					v6 = (*uint8)(unsafe.Pointer(dr))
				} else {
					sub_4B8E10(a4, (*byte)(unsafe.Pointer(uintptr(1<<v7))))
				}
				*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v9))), unsafe.Sizeof(uint16(0))*0)) = uint16(nox_xxx_drawObject_4C4770_draw(a1, (*nox_drawable)(unsafe.Pointer(v6)), int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v10 + int32(*(*uint8)(unsafe.Add(unsafe.Pointer(v6), 297)))*4 + 4))) + uint32(a6*4)))))))
			}
		}
		v7++
		v8 = (*int32)(unsafe.Add(unsafe.Pointer(v8), unsafe.Sizeof(int32(0))*1))
		if v7 >= 27 {
			break
		}
	}
	return int16(v9)
}
func sub_4B8E10(a1 *uint32, a2 *byte) *uint32 {
	var (
		result *uint32
		v3     int32
		v4     int32
		v5     int32
		v6     int32
		v7     *uint32
		v8     int32
		v9     *uint8
		v10    *uint32
		v11    int32
		v12    *int32
		v13    int32
		v14    **uint32
		v15    int32
	)
	result = a1
	v3 = 0
	for (*byte)(unsafe.Pointer(uintptr(*result))) != a2 {
		v3++
		result = (*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*6))
		if v3 >= 27 {
			return result
		}
	}
	v4 = sub_415840(a2)
	result = nox_xxx_getProjectileClassById_413250(v4)
	v7 = result
	if result != nil {
		v8 = 1
		v9 = (*uint8)(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*4))))
		for {
			*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&result))), 0)) = *(*uint8)(unsafe.Add(unsafe.Pointer(v9), 1))
			*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v6))), 0)) = *v9
			*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v5))), 0)) = *((*uint8)(unsafe.Add(unsafe.Pointer(v9), -1)))
			sub_4340A0(func() int32 {
				p := &v8
				x := *p
				*p++
				return x
			}(), v5, v6, int32(uintptr(unsafe.Pointer(result))))
			v9 = (*uint8)(unsafe.Add(unsafe.Pointer(v9), 3))
			if v8 >= 7 {
				break
			}
		}
		v10 = a1
		v11 = v3 * 3
		v12 = (*int32)(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer(v7), unsafe.Sizeof(uint32(0))*9))))
		v13 = 4
		v14 = (**uint32)(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*uintptr(v11*2+1)))))
		for {
			result = *v14
			if *v14 != nil {
				*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v5))), 0)) = *((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(result))), 26)))
				*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v10))), 0)) = *((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(result))), 25)))
				v15 = v5
				*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v5))), 0)) = *((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(result))), 24)))
				sub_4340A0(*v12, v5, int32(uintptr(unsafe.Pointer(v10))), v15)
			}
			v14 = (**uint32)(unsafe.Add(unsafe.Pointer(v14), unsafe.Sizeof((*uint32)(nil))*1))
			v12 = (*int32)(unsafe.Add(unsafe.Pointer(v12), unsafe.Sizeof(int32(0))*1))
			v13--
			if v13 == 0 {
				break
			}
		}
	}
	return result
}
func nox_xxx_drawOtherPlayerHP_4B8EB0(a1 *uint32, dr *nox_drawable, a3 uint16, a4 int8) {
	var (
		a2 int32 = int32(uintptr(unsafe.Pointer(dr)))
		v4 int32
		v5 int32
		v6 int32
		v7 float32
	)
	if a2 != 0 {
		v4 = int32(*a1 + *(*uint32)(unsafe.Pointer(uintptr(a2 + 12))) - *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*4)))
		v7 = *(*float32)(unsafe.Pointer(uintptr(a2 + 48))) + *(*float32)(unsafe.Pointer(uintptr(a2 + 48)))
		v5 = int(v7) + v4
		v6 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 16))) + *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*1)) - uint32(*(*int16)(unsafe.Pointer(uintptr(a2 + 104)))) - *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*5)) - 48)
		nox_client_drawSetColor_434460(int32(nox_color_black_2650656))
		nox_client_drawRectFilledOpaque_49CE30(v5, v6, 2, 48)
		if int32(a4) != 0 {
			nox_client_drawSetColor_434460(*memmap.PtrInt32(0x8531A0, 2572))
		} else {
			nox_client_drawSetColor_434460(*memmap.PtrInt32(0x85B3FC, 940))
		}
		nox_client_drawRectFilledOpaque_49CE30(v5, v6-int32(a3)*48/100+48, 2, int32(a3)*48/100)
		if int32(a4) != 0 {
			nox_client_drawSetColor_434460(*memmap.PtrInt32(0x85B3FC, 984))
		} else {
			nox_client_drawSetColor_434460(int32(nox_color_violet_2598268))
		}
		nox_client_drawBorderLines_49CC70(v5-1, v6-1, 4, 50)
	}
}
func sub_4B8FA0(dr *nox_drawable, a2 *int32, a3 *int32) int32 {
	var (
		a1 int32 = int32(uintptr(unsafe.Pointer(dr)))
		v3 int32
		v4 int32
		v5 bool
		v6 int32
		v8 int32
		v9 int32
	)
	v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 304))))
	if *(*uint32)(unsafe.Pointer(uintptr(a1 + 276))) != 0 || sub_48D830(dr) == 0 {
		if *(*uint32)(unsafe.Pointer(uintptr(a1 + 276))) != 4 || (func() bool {
			v5 = !nox_xxx_spriteCheckFlag31_4356C0(dr, 31)
			v4 = 53
			return v5
		}()) {
			v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 276))))
		}
	} else {
		v4 = 19
	}
	v5 = int32(*(*uint16)(unsafe.Pointer(uintptr(v3 + v4*264 + 44)))) == 0
	v6 = v3 + v4*264 + 4
	if v5 {
		return 0
	}
	v8 = sub_4BC5D0(dr, v6)
	if v8 < 0 {
		return 0
	}
	v9 = int32(*(*int16)(unsafe.Pointer(uintptr(v6 + 40))))
	if v8 >= v9 {
		v8 = v9 - 1
	}
	if a2 != nil {
		*a2 = v6
	}
	if a3 != nil {
		*a3 = v8
	}
	return int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v6 + 48))) + uint32(int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 297))))*4) + 4))) + uint32(v8*4)))))
}
