package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	"github.com/gotranspile/cxgo/runtime/stdio"
	"github.com/noxworld-dev/opennox/v1/common/alloc"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"image"
	"math"
	"unicode"
	"unsafe"
)

var dword_5d4594_1189592 unsafe.Pointer = nil
var dword_5d4594_1189596 unsafe.Pointer = nil
var nox_arr_84EB20 [63840]uint8 = [63840]uint8{}
var nox_pixbuffer_rows_3798776 **uint8 = nil
var nox_pixbuffer_3798788 *uint8 = nil
var dword_5d4594_1096640 unsafe.Pointer = nil
var func_587000_154940 func(*int2, uint32, uint32) = func(arg1 *int2, arg2 uint32, arg3 uint32) {
	nox_xxx_tileDraw_4815E0((*uint32)(unsafe.Pointer(arg1)), int32(arg2))
}
var func_587000_154944 func(int32, int32) int32 = func(arg1 int32, arg2 int32) int32 {
	return int32(nox_xxx_drawTexEdgesProbably_481900((*uint32)(unsafe.Pointer(uintptr(arg1))), (*uint32)(unsafe.Pointer(uintptr(arg2)))))
}
var nox_client_drawAtFunc_3799500 func(*nox_video_bag_image_t, int32, int32) = nil
var nox_video_cutSize int32 = 100
var nox_client_spriteUnderCursorXxx_1096644 unsafe.Pointer = nil

func sub_476080(a1 *uint8) int32 {
	var (
		v1     int32
		v2     int32
		result int32
		v4     int32
		v5     int32
	)
	if *memmap.PtrUint32(0x852978, 8) == 0 {
		return int32(*(*uint8)(unsafe.Add(unsafe.Pointer(a1), 6)))*23 + 11
	}
	switch *a1 {
	case 0:
		fallthrough
	case 3:
		fallthrough
	case 11:
		v1 = -23
		v2 = int32(*(*uint8)(unsafe.Add(unsafe.Pointer(a1), 5)))*23 + 22
		result = int32(*(*uint8)(unsafe.Add(unsafe.Pointer(a1), 6))) * 23
		goto LABEL_5
	case 1:
		fallthrough
	case 4:
		fallthrough
	case 12:
		v1 = 23
		v2 = int32(*(*uint8)(unsafe.Add(unsafe.Pointer(a1), 5))) * 23
		result = int32(*(*uint8)(unsafe.Add(unsafe.Pointer(a1), 6))) * 23
	LABEL_5:
		v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(*memmap.PtrUint32(0x852978, 8) + 12))) - uint32(v2))
		v5 = int32(uint32(v1)*(*(*uint32)(unsafe.Pointer(uintptr(*memmap.PtrUint32(0x852978, 8) + 16)))-uint32(result)) - uint32(v4*23))
		if v1 < 0 {
			v5 = int32(uint32(v4*23) - uint32(v1)*(*(*uint32)(unsafe.Pointer(uintptr(*memmap.PtrUint32(0x852978, 8) + 16)))-uint32(result)))
		}
		if v5 < 0 {
			result += 22
		}
	default:
		result = int32(*(*uint8)(unsafe.Add(unsafe.Pointer(a1), 6)))*23 + 11
	}
	return result
}
func sub_476160(a1 unsafe.Pointer, a2 unsafe.Pointer) int32 {
	var (
		v2     int32
		result int32
	)
	if int32(int8(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(a1) + 112))))) >= 0 {
		v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(a1) + 16))) + uint32(*(*int16)(unsafe.Pointer(uintptr(*(*uint32)(a1) + 104)))))
	} else {
		v2 = sub_4761B0((*nox_drawable)(unsafe.Pointer(uintptr(*(*uint32)(a1)))))
	}
	if int32(int8(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(a2) + 112))))) >= 0 {
		result = int32(uint32(v2) - (*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(a2) + 16))) + uint32(*(*int16)(unsafe.Pointer(uintptr(*(*uint32)(a2) + 104))))))
	} else {
		result = v2 - sub_4761B0((*nox_drawable)(unsafe.Pointer(uintptr(*(*uint32)(a2)))))
	}
	return result
}
func sub_4761B0(a1p *nox_drawable) int32 {
	var (
		a1     int32 = int32(uintptr(unsafe.Pointer(a1p)))
		result int32
		v2     int32
		v3     int32
		v4     int32
	)
	if *memmap.PtrUint32(0x852978, 8) == 0 {
		return int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 16))) + uint32(*memmap.PtrInt32(0x587000, uint32(int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 299))))*8)+0x2FE5C)/2))
	}
	result = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 16))))
	v2 = int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 299)))) * 8
	v3 = int32((*(*uint32)(unsafe.Pointer(uintptr(*memmap.PtrUint32(0x852978, 8) + 16)))-uint32(result))*uint32(*memmap.PtrInt32(0x587000, uint32(v2)+0x2FE58)) - (*(*uint32)(unsafe.Pointer(uintptr(*memmap.PtrUint32(0x852978, 8) + 12)))-*(*uint32)(unsafe.Pointer(uintptr(a1 + 12))))*uint32(*memmap.PtrInt32(0x587000, uint32(v2)+0x2FE5C)))
	if *memmap.PtrInt32(0x587000, uint32(v2)+0x2FE58) < 0 {
		v3 = int32((*(*uint32)(unsafe.Pointer(uintptr(*memmap.PtrUint32(0x852978, 8) + 12)))-*(*uint32)(unsafe.Pointer(uintptr(a1 + 12))))*uint32(*memmap.PtrInt32(0x587000, uint32(int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 299))))*8)+0x2FE5C)) - (*(*uint32)(unsafe.Pointer(uintptr(*memmap.PtrUint32(0x852978, 8) + 16)))-uint32(result))*uint32(*memmap.PtrInt32(0x587000, uint32(int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 299))))*8)+0x2FE58)))
	}
	v4 = result + *memmap.PtrInt32(0x587000, uint32(int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 299))))*8)+0x2FE5C)
	if v3 >= 0 {
		if v4 <= result {
			result += *memmap.PtrInt32(0x587000, uint32(int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 299))))*8)+0x2FE5C)
		}
	} else if v4 > result {
		result += *memmap.PtrInt32(0x587000, uint32(int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 299))))*8)+0x2FE5C)
	}
	return result
}
func sub_476240(a1 unsafe.Pointer, a2 unsafe.Pointer) int32 {
	var v2 int32
	v2 = sub_476080(*(**uint8)(a1))
	return v2 - sub_476080(*(**uint8)(a2))
}
func sub_476270(a1p *nox_draw_viewport_t) {
	var (
		a1  *uint32 = (*uint32)(unsafe.Pointer(a1p))
		v1  *uint32
		v2  int32
		v3  int32
		v4  int32
		v5  int32
		v6  int32
		v7  int32
		v8  *float32
		v9  int32
		v10 int32
		v11 int32
		v12 int32
		v13 int32
		v15 int32
		v16 int32
		v17 int32
		v18 int32
		v19 int16
		v21 int32
		v22 int32
		v23 *float32
		v24 int32
		v25 *float32
		v26 int32
		v27 int32
		v28 int32
		v29 int32
		i   int32
		v31 int32
	)
	v27 = sub_50CB00()
	v24 = sub_50D210()
	v1 = a1
	v23 = (*float32)(sub_50CB10())
	if v27 >= 2 {
		v2 = 1
		v3 = 0
		for i = 0; ; v3 = i {
			v4 = int32(uint32(uint64(*v1) + uint64(int64(*(*float32)(unsafe.Add(unsafe.Pointer(v23), unsafe.Sizeof(float32(0))*uintptr(v3*2))))) - uint64(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*4)))))
			v25 = (*float32)(unsafe.Add(unsafe.Pointer(v23), unsafe.Sizeof(float32(0))*uintptr(v3*2+1)))
			v5 = int32(uint32(uint64(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*1))) + uint64(int64(*v25)) - uint64(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*5)))))
			if i == v24 {
				nox_client_drawSetColor_434460(*memmap.PtrInt32(0x85B3FC, 940))
			} else if float64(*(*float32)(unsafe.Add(unsafe.Pointer(v23), unsafe.Sizeof(float32(0))*uintptr(i*2)))) < 0.0 || float64(*v25) < 0.0 {
				nox_client_drawSetColor_434460(int32(nox_color_blue_2650684))
			} else {
				nox_client_drawSetColor_434460(int32(nox_color_yellow_2589772))
			}
			nox_client_drawBorderLines_49CC70(v4-1, v5-1, 3, 3)
			nox_client_drawSetColor_434460(*memmap.PtrInt32(0x8531A0, 2572))
			nox_client_drawAddPoint_49F500(v4, v5)
			v6 = int32(uint32(uint64(*v1) + uint64(int64(*(*float32)(unsafe.Add(unsafe.Pointer(v23), unsafe.Sizeof(float32(0))*uintptr(v2*2))))) - uint64(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*4)))))
			v7 = int32(uint32(uint64(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*1))) + uint64(int64(*(*float32)(unsafe.Add(unsafe.Pointer(v23), unsafe.Sizeof(float32(0))*uintptr(v2*2+1))))) - uint64(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*5)))))
			nox_client_drawAddPoint_49F500(v6, v7)
			nox_client_drawLineFromPoints_49E4B0()
			i = func() int32 {
				p := &v2
				x := *p
				*p++
				return x
			}()
			if v2 >= v27 {
				break
			}
		}
		nox_client_drawSetColor_434460(int32(nox_color_yellow_2589772))
		nox_client_drawBorderLines_49CC70(v6-1, v7-1, 3, 3)
	}
	v8 = (*float32)(unsafe.Pointer(sub_50CD30()))
	v9 = int32(uint32(uint64(*v1) + uint64(int64(*v8)) - uint64(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*4)))))
	v10 = int32(uint32(uint64(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*1))) + uint64(int64(*(*float32)(unsafe.Add(unsafe.Pointer(v8), unsafe.Sizeof(float32(0))*1)))) - uint64(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*5)))))
	nox_client_drawSetColor_434460(int32(nox_color_white_2523948))
	nox_client_drawBorderLines_49CC70(v9-2, v10-2, 5, 5)
	v11 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*4)))
	v12 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*5)))
	v26 = int32(int64(float64(v11/23) + 0.5))
	v13 = int32(int64(float64(v12/23) + 0.5))
	v28 = int32(int64(float64((nox_win_width+v11)/23) + 0.5))
	var v14 int32 = int32(int64(float64((nox_win_height+v12)/23) + 0.5))
	v15 = v14
	v29 = v14
	if v13 > v14 {
		return
	}
	v16 = v13 * 23
	v31 = v13 * 23
	for {
		v17 = v26
		if v26 > v28 {
			goto LABEL_37
		}
		v18 = v26 * 23
		for {
			if int32(sub_57B500(v17, v13, 64)) == -1 {
				v19 = sub_50AB50(v17, v13)
				if int32(v19)&64 != 0 {
					nox_set_color_rgb_434430(0, math.MaxUint8, 0)
				} else if int32(v19)&2 != 0 {
					nox_set_color_rgb_434430(0, math.MaxUint8, math.MaxUint8)
				} else {
					if int32(v19)&1 != 0 {
						nox_set_color_rgb_434430(0, 100, 100)
						goto LABEL_34
					}
					if int32(v19)&60 != 0 {
						v22 = math.MaxUint8
					LABEL_32:
						v21 = 0
					LABEL_33:
						nox_set_color_rgb_434430(math.MaxUint8, v21, v22)
						goto LABEL_34
					}
					if int32(v19)&1024 != 0 {
						v22 = 0
						goto LABEL_32
					}
					if int32(v19)&512 != 0 {
						v22 = 0
						v21 = math.MaxUint8
						goto LABEL_33
					}
					if (int32(v19) & 256) == 0 {
						goto LABEL_35
					}
					nox_set_color_rgb_434430(100, 100, 0)
				}
			} else {
				nox_set_color_rgb_434430(0, 0, math.MaxUint8)
			}
		LABEL_34:
			nox_client_drawBorderLines_49CC70(int32(uint32(v18)-*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*4))), int32(uint32(v31)-*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*5))), 23, 23)
		LABEL_35:
			v17++
			v18 += 23
			if v17 > v28 {
				break
			}
		}
		v15 = v29
		v16 = v31
	LABEL_37:
		v13++
		v16 += 23
		v31 = v16
		if v13 > v15 {
			break
		}
	}
}
func sub_4765F0(a1p *nox_draw_viewport_t) int32 {
	var (
		a1     int32 = int32(uintptr(unsafe.Pointer(a1p)))
		result int32
		v2     int32
	)
	*memmap.PtrUint32(6112660, 1096440) = nox_client_highResFloors_154952
	result = int32(*memmap.PtrUint32(0x852978, 8))
	*memmap.PtrUint32(6112660, 1096444) = nox_client_highResFrontWalls_80820
	if *memmap.PtrUint32(0x852978, 8) != 0 && nox_client_lockHighResFloors_1193152 == 0 {
		v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(*memmap.PtrUint32(0x852978, 8) + 276))))
		if (v2 == 3 || v2 == 6 || v2 == 45) && (*(*int32)(unsafe.Pointer(uintptr(a1 + 16)))-*memmap.PtrInt32(6112660, 1096556) >= 4 || *(*int32)(unsafe.Pointer(uintptr(a1 + 20)))-*memmap.PtrInt32(6112660, 1096560) >= 4) {
			nox_client_highResFloors_154952 = 0
			nox_client_highResFrontWalls_80820 = 0
		}
		result = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 16))))
		*memmap.PtrUint32(6112660, 1096556) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 16)))
		*memmap.PtrUint32(6112660, 1096560) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 20)))
	}
	return result
}
func sub_476680() int32 {
	var result int32
	result = int32(*memmap.PtrUint32(6112660, 1096440))
	nox_client_highResFloors_154952 = *memmap.PtrUint32(6112660, 1096440)
	nox_client_highResFrontWalls_80820 = *memmap.PtrUint32(6112660, 1096444)
	return result
}
func nox_video_setCutSize_4766A0(a1 int32) {
	if a1 >= 65 {
		if a1 > 100 {
			a1 = 100
		}
		nox_video_cutSize = a1
	} else {
		nox_video_cutSize = 65
	}
}
func nox_video_getCutSize_4766D0() int32 {
	return nox_video_cutSize
}
func sub_476850(vp *nox_draw_viewport_t, dr *nox_drawable) {
	var (
		a2  *uint8 = (*uint8)(unsafe.Pointer(dr))
		v2  int32
		v3  *int32
		v4  int32
		v5  int32
		v6  *uint32
		v7  int32
		v8  int32
		v9  int32
		v10 int32
		v11 int32
		v12 uint32
		v13 *uint8
		v14 uint32
		v15 *byte
		v16 uint32
		v17 int32
		v18 int32
		v19 uint32
		v20 uint32
		v21 int32
		v22 *uint8
		v23 *byte
		v24 *uint8
		v25 int32
		v26 int32
		i   int32
		v28 uint32
		v29 int32
		v30 int32
		v31 uint32
	)
	if cgoFuncAddr(*((*func(*int32, int32) int32)(unsafe.Add(unsafe.Pointer((*func(*int32, int32) int32)(unsafe.Pointer(a2))), unsafe.Sizeof(uintptr(0))*75)))) == cgoFuncAddr(nox_thing_static_draw) {
		if *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a2))), unsafe.Sizeof(uint32(0))*28)))&0x40000 != 0 && (*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a2))), unsafe.Sizeof(uint32(0))*30)))&0x1000000) == 0 {
			return
		}
		v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a2))), unsafe.Sizeof(uint32(0))*76))) + 4))))
	} else {
		v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a2))), unsafe.Sizeof(uint32(0))*76))) + 4))) + *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a2))), unsafe.Sizeof(uint32(0))*77)))*4))))
	}
	v3 = (*int32)(nox_video_getImagePixdata_func((*nox_video_bag_image_t)(unsafe.Pointer(uintptr(v2)))))
	if v3 != nil {
		v4 = *v3
		v5 = *(*int32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(int32(0))*1))
		v6 = (*uint32)(unsafe.Pointer((*int32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(int32(0))*2))))
		v26 = v5
		v7 = int32(dword_5d4594_3798820)
		v8 = int32(*v6 + *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a2))), unsafe.Sizeof(uint32(0))*3))) - uint32(*a2))
		v9 = int32(uintptr(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*2)))))
		v30 = v4
		v10 = int32(*(*uint32)(unsafe.Pointer(uintptr(v9 - 4))) + *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a2))), unsafe.Sizeof(uint32(0))*4))) - uint32(*((*int16)(unsafe.Add(unsafe.Pointer((*int16)(unsafe.Pointer(a2))), unsafe.Sizeof(int16(0))*53)))) - uint32(*((*int16)(unsafe.Add(unsafe.Pointer((*int16)(unsafe.Pointer(a2))), unsafe.Sizeof(int16(0))*52)))) - uint32(*(*uint8)(unsafe.Add(unsafe.Pointer(a2), 1))))
		if v8 < *(*int32)(unsafe.Pointer(&dword_5d4594_3798820)) || uint32(v8+v4) >= uint32(*(*int32)(unsafe.Pointer(&dword_5d4594_3798820)))+dword_5d4594_3798800 || v10 < *(*int32)(unsafe.Pointer(&dword_5d4594_3798824)) || uint32(v10+v26) >= uint32(*(*int32)(unsafe.Pointer(&dword_5d4594_3798824)))+dword_5d4594_3798808 {
			*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a2))), unsafe.Sizeof(uint32(0))*86))) = 0
		} else {
			v11 = int32(nox_xxx_waypointCounterMB_587000_154948)
			if *(*int32)(unsafe.Pointer(&nox_xxx_waypointCounterMB_587000_154948)) <= 0 {
				*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a2))), unsafe.Sizeof(uint32(0))*86))) = 0
				v7 = int32(dword_5d4594_3798820)
				v11 = int32(nox_xxx_waypointCounterMB_587000_154948)
			}
			if v11-*((*int32)(unsafe.Add(unsafe.Pointer((*int32)(unsafe.Pointer(a2))), unsafe.Sizeof(int32(0))*86))) > 1 || v11 <= 0 {
				v12 = dword_5d4594_3798844
				v13 = (*uint8)(unsafe.Pointer(uintptr(v9 + 1)))
				v28 = dword_5d4594_3798844
				v25 = int32(dword_5d4594_3798844 - dword_5d4594_3798796)
				v14 = uint32(v8) + dword_5d4594_3798796 + dword_5d4594_3798836 + dword_5d4594_3798804*(uint32(v10)+dword_5d4594_3798840-dword_5d4594_3798824) - uint32(v7)
				v31 = v14
				if v14 >= uint32(*(*int32)(unsafe.Pointer(&dword_5d4594_3798844))) {
					v14 -= uint32(v25)
					v31 = v14
				}
				if v26 != 0 {
					v29 = v26
					for {
						v15 = (*byte)(unsafe.Pointer(uintptr(v14)))
						for i = v30; i > 0; i -= int32(v19) {
							v16 = uint32(*(*uint8)(unsafe.Add(unsafe.Pointer(v13), 1)))
							v17 = int32(*v13) & 15
							v13 = (*uint8)(unsafe.Add(unsafe.Pointer(v13), 2))
							v18 = v17 - 1
							v19 = v16
							if v18 != 0 {
								if v18 == 1 {
									if uint32(uintptr(unsafe.Pointer(v15))) >= v12 || uint32(uintptr(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v15), v16))))) < v12 {
										libc.MemCpy(unsafe.Pointer(v15), unsafe.Pointer(v13), int((v16>>2)*4))
										v24 = (*uint8)(unsafe.Add(unsafe.Pointer(v13), (v16>>2)*4))
										v23 = (*byte)(unsafe.Add(unsafe.Pointer(v15), (v16>>2)*4))
										v13 = (*uint8)(unsafe.Add(unsafe.Pointer(v13), v16))
										v15 = (*byte)(unsafe.Add(unsafe.Pointer(v15), v16))
										libc.MemCpy(unsafe.Pointer(v23), unsafe.Pointer(v24), int(v16&3))
										v14 = v31
										v12 = v28
									} else {
										v20 = uint32(uintptr(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v15), v16-v12)))))
										v19 -= v20
										libc.MemCpy(unsafe.Pointer(v15), unsafe.Pointer(v13), int(v19))
										v21 = int32(dword_5d4594_3798796)
										v22 = (*uint8)(unsafe.Add(unsafe.Pointer(v13), v19))
										v13 = (*uint8)(unsafe.Add(unsafe.Pointer(v13), v20+v19))
										libc.MemCpy(*(*unsafe.Pointer)(unsafe.Pointer(&dword_5d4594_3798796)), unsafe.Pointer(v22), int(v20))
										v15 = (*byte)(unsafe.Pointer(uintptr(v20 + uint32(v21))))
										v14 = v31
										v12 = v28
									}
								}
							} else {
								v15 = (*byte)(unsafe.Add(unsafe.Pointer(v15), v16))
								if uint32(uintptr(unsafe.Pointer(v15))) >= v12 {
									v15 = (*byte)(unsafe.Add(unsafe.Pointer(v15), -v25))
								}
							}
						}
						v14 += dword_5d4594_3798804
						v31 = v14
						if v14 >= v12 {
							v14 -= uint32(v25)
							v31 = v14
						}
						v29--
						if v29 == 0 {
							break
						}
					}
				}
			} else {
				*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a2))), unsafe.Sizeof(uint32(0))*86))) = uint32(v11)
			}
		}
	}
}
func sub_476AE0(vp *nox_draw_viewport_t, dr *nox_drawable) unsafe.Pointer {
	var (
		a2      *uint8 = (*uint8)(unsafe.Pointer(dr))
		v2      *uint8
		result  int32
		result2 func(*int32, int32) int32
		v4      int32
		v5      int32
		v6      int32
		v7      int32
		v8      *byte
		v9      int32
		v10     int32
		v11     uint32
		v12     *uint8
		v13     int32
		v14     int32
		v15     uint32
		v16     uint32
		v17     int32
		v18     int32
		v19     int32
		v20     int32
		v21     int32
		v22     int32
		v23     int32
		v24     int32
		v25     int32
		v26     int32
		v27     int32
		v28     int32
		v29     uint32
		v30     int32
		v31     int32
		v32     int32
		v33     func(uint32, *uint8, int32)
	)
	v2 = a2
	result2 = cgoAsFunc(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a2))), unsafe.Sizeof(uint32(0))*75))), (*func(*int32, int32) int32)(nil)).(func(*int32, int32) int32)
	if cgoFuncAddr(result2) == cgoFuncAddr(nox_thing_static_draw) {
		if *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a2))), unsafe.Sizeof(uint32(0))*28)))&0x40000 != 0 && (*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a2))), unsafe.Sizeof(uint32(0))*30)))&0x1000000) == 0 {
			return unsafe.Pointer(cgoFuncAddr(result2))
		}
		v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a2))), unsafe.Sizeof(uint32(0))*76))) + 4))))
	} else {
		v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a2))), unsafe.Sizeof(uint32(0))*76))) + 4))) + *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a2))), unsafe.Sizeof(uint32(0))*77)))*4))))
	}
	result = int32(uintptr(nox_video_getImagePixdata_func((*nox_video_bag_image_t)(unsafe.Pointer(uintptr(v4))))))
	if result != 0 {
		v33 = func(arg1 uint32, arg2 *uint8, arg3 int32) {
			sub_476D70((*uint32)(unsafe.Pointer(uintptr(arg1))), (*int32)(unsafe.Pointer(arg2)), uint32(arg3))
		}
		v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(result))))
		v6 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(uintptr(result)))), unsafe.Sizeof(uint32(0))*1))))
		v27 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(uintptr(result)))), unsafe.Sizeof(uint32(0))*1))))
		v7 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(uintptr(result)))), unsafe.Sizeof(uint32(0))*2))) + *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), unsafe.Sizeof(uint32(0))*3))) - uint32(*v2))
		v8 = (*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(uintptr(result)))), 16))
		result = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(uintptr(result)))), unsafe.Sizeof(uint32(0))*3))) + *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), unsafe.Sizeof(uint32(0))*4))) - uint32(*((*int16)(unsafe.Add(unsafe.Pointer((*int16)(unsafe.Pointer(v2))), unsafe.Sizeof(int16(0))*53)))) - uint32(*((*int16)(unsafe.Add(unsafe.Pointer((*int16)(unsafe.Pointer(v2))), unsafe.Sizeof(int16(0))*52)))) - uint32(*(*uint8)(unsafe.Add(unsafe.Pointer(v2), 1))))
		v31 = v5
		if v7 < *(*int32)(unsafe.Pointer(&dword_5d4594_3798820)) || uint32(v7+v5) >= uint32(*(*int32)(unsafe.Pointer(&dword_5d4594_3798820)))+dword_5d4594_3798800 || (func() bool {
			v9 = int32(dword_5d4594_3798824)
			return result < *(*int32)(unsafe.Pointer(&dword_5d4594_3798824))
		}()) || uint32(result+v6) >= uint32(*(*int32)(unsafe.Pointer(&dword_5d4594_3798824)))+dword_5d4594_3798808 {
			*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), unsafe.Sizeof(uint32(0))*86))) = 0
		} else {
			v10 = int32(nox_xxx_waypointCounterMB_587000_154948)
			if *(*int32)(unsafe.Pointer(&nox_xxx_waypointCounterMB_587000_154948)) <= 0 {
				*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), unsafe.Sizeof(uint32(0))*86))) = 0
				v9 = int32(dword_5d4594_3798824)
				v10 = int32(nox_xxx_waypointCounterMB_587000_154948)
			}
			if v10-*((*int32)(unsafe.Add(unsafe.Pointer((*int32)(unsafe.Pointer(v2))), unsafe.Sizeof(int32(0))*86))) > 1 || v10 <= 0 {
				v11 = dword_5d4594_3798844
				v12 = (*uint8)(unsafe.Add(unsafe.Pointer(v8), 1))
				v29 = dword_5d4594_3798844
				v26 = int32(dword_5d4594_3798844 - dword_5d4594_3798796)
				v13 = int32(dword_5d4594_3798804 * (uint32(result) + dword_5d4594_3798840 - uint32(v9)))
				v14 = int32(uint32(v7) + dword_5d4594_3798836 - dword_5d4594_3798820)
				v15 = uint32(v13) + dword_5d4594_3798796 + uint32(v14*2)
				v25 = int32(uint32(v13) + dword_5d4594_3798796 + uint32(v14*2))
				if v15 >= uint32(*(*int32)(unsafe.Pointer(&dword_5d4594_3798844))) {
					v15 -= uint32(v26)
					v25 = int32(v15)
				}
				result = v27 - 1
				if v27 != 0 {
					v30 = v27
					for {
						v16 = v15
						v28 = v31
						if v31 > 0 {
							for {
								v17 = int32(*(*uint8)(unsafe.Add(unsafe.Pointer(v12), 1)))
								v18 = int32(*v12) & 15
								v12 = (*uint8)(unsafe.Add(unsafe.Pointer(v12), 2))
								v19 = v18 - 1
								v32 = v17
								v20 = v17 * 2
								if v19 != 0 {
									if v19 == 1 {
										if v16 >= v11 || v16+uint32(v20) < v11 {
											v33(v16, v12, v17*2)
											v12 = (*uint8)(unsafe.Add(unsafe.Pointer(v12), v20))
											v16 += uint32(v20)
										} else {
											v21 = int32(v16 + uint32(v20) - v29)
											v22 = int32(v29 - v16)
											v33(v16, v12, int32(v29-v16))
											v23 = int32(dword_5d4594_3798796)
											v24 = int32(uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v12), v22)))))
											v33(dword_5d4594_3798796, (*uint8)(unsafe.Pointer(uintptr(v24))), v21)
											v12 = (*uint8)(unsafe.Pointer(uintptr(v21 + v24)))
											v16 = uint32(v21 + v23)
											v11 = v29
										}
									}
								} else {
									v16 += uint32(v20)
									if v16 >= v11 {
										v16 -= uint32(v26)
									}
								}
								v28 -= v32
								if v28 <= 0 {
									break
								}
							}
							v15 = uint32(v25)
						}
						v15 += dword_5d4594_3798804
						v25 = int32(v15)
						if v15 >= v11 {
							v15 -= uint32(v26)
							v25 = int32(v15)
						}
						result = func() int32 {
							p := &v30
							*p--
							return *p
						}()
						if v30 == 0 {
							break
						}
					}
				}
			} else {
				*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), unsafe.Sizeof(uint32(0))*86))) = uint32(v10)
			}
		}
	}
	return unsafe.Pointer(uintptr(result))
}
func sub_476D70(a1 *uint32, a2 *int32, a3 uint32) int16 {
	var (
		v3 *uint32
		v4 int32
		v5 *int32
		v6 int32
	)
	v3 = a1
	v4 = int32(a3 >> 2)
	v5 = a2
	if a3>>2 != 0 {
		for {
			v6 = *v5
			v5 = (*int32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(int32(0))*1))
			*v3 = uint32(v6)
			v3 = (*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*1))
			if func() int32 {
				p := &v4
				x := *p
				*p--
				return x
			}() <= 1 {
				break
			}
		}
	}
	if a3&3 != 0 {
		*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v6))), unsafe.Sizeof(uint16(0))*0)) = *(*uint16)(unsafe.Pointer(v5))
		*(*uint16)(unsafe.Pointer(v3)) = *(*uint16)(unsafe.Pointer(v5))
	}
	return int16(v6)
}
func sub_476DA0(a1 *uint32, a2 *uint32, a3 uint32) {
	var (
		v3 *uint32
		v4 int32
		v5 *uint32
	)
	v3 = a1
	v4 = int32(a3 >> 2)
	v5 = a2
	if a3>>2 != 0 {
		for {
			*v3 = ((*v5 & 0xFFC0FFC0) >> 1) | (*v5 & 0x1F001F)
			v5 = (*uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(uint32(0))*1))
			v3 = (*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*1))
			if func() int32 {
				p := &v4
				x := *p
				*p--
				return x
			}() <= 1 {
				break
			}
		}
	}
	if a3&3 != 0 {
		*(*uint16)(unsafe.Pointer(v3)) = uint16(int16((int32(uint16(int16(int32(*(*uint16)(unsafe.Pointer(v5)))&0xFFC0))) >> 1) | (int32(*(*uint16)(unsafe.Pointer(v5))) & 31)))
	}
}
func nox_client_setPhonemeFrame_476E00(a1 int32) int32 {
	var result int32
	result = a1
	*memmap.PtrUint32(6112660, uint32(a1*4)+0x10BB94) = nox_frame_xxx_2598000
	return result
}
func sub_476E20() *uint32 {
	var (
		v0 int32
		v1 *byte
		v2 *uint32
	)
	v0 = 0
	for {
		v1 = (*byte)(unsafe.Pointer(nox_xxx_gLoadImg(*(**byte)(memmap.PtrOff(0x587000, uint32(v0)+0x24EE8)))))
		*memmap.PtrUint32(6112660, uint32(v0)+0x10BB74) = uint32(uintptr(unsafe.Pointer(v1)))
		if v1 == nil {
			break
		}
		v0 += 4
		if v0 >= 32 {
			v2 = (*uint32)(unsafe.Pointer(nox_window_new(nil, 64, (nox_win_width-100)/2, (nox_win_height-100)/2, 1, 1, nil)))
			(*nox_window)(unsafe.Pointer(v2)).SetAllFuncs(nil, func(arg1 *nox_window, arg2 unsafe.Pointer) int32 {
				return sub_476E90()
			}, nil)
			return v2
		}
	}
	return nil
}
func sub_476E90() int32 {
	var (
		v0 *uint8
		v1 int32
	)
	v0 = (*uint8)(memmap.PtrOff(0x587000, 151208))
	v1 = 0
	for {
		if *memmap.PtrUint32(6112660, uint32(v1)+0x10BB94) != 0 {
			nox_client_drawImageAt_47D2C0((*nox_video_bag_image_t)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, uint32(v1)+0x10BB74)))), int32(uint32(nox_win_width/2)+*(*uint32)(unsafe.Pointer(v0))-16), int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v0))), unsafe.Sizeof(uint32(0))*1)))+uint32(nox_win_height/2)-41))
			if (nox_frame_xxx_2598000 - *memmap.PtrUint32(6112660, uint32(v1)+0x10BB94)) > 3 {
				*memmap.PtrUint32(6112660, uint32(v1)+0x10BB94) = 0
			}
		}
		v0 = (*uint8)(unsafe.Add(unsafe.Pointer(v0), 8))
		v1 += 4
		if int32(uintptr(unsafe.Pointer(v0))) >= int32(uintptr(memmap.PtrOff(0x587000, 151272))) {
			break
		}
	}
	return 1
}
func nox_xxx_clientDeleteSprite_476F10_drawable(dr *nox_drawable) int32 {
	var (
		a1     int32 = int32(uintptr(unsafe.Pointer(dr)))
		result int32
	)
	result = a1
	if dword_5d4594_1096640 == unsafe.Pointer(uintptr(a1)) {
		dword_5d4594_1096640 = nil
	}
	if nox_client_spriteUnderCursorXxx_1096644 == unsafe.Pointer(uintptr(a1)) {
		nox_client_spriteUnderCursorXxx_1096644 = nil
	}
	return result
}
func nox_xxx_packetGetMarshall_476F40() uint32 {
	var result uint32
	if dword_5d4594_1096640 != nil {
		result = nox_xxx_netGetUnitCodeCli_578B00(*(*int32)(unsafe.Pointer(&dword_5d4594_1096640)))
	} else {
		result = 0
	}
	return result
}
func sub_476F60() int32 {
	var result int32
	if dword_5d4594_1096640 != nil {
		result = int32(*(*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(dword_5d4594_1096640)) + 128))))
	} else {
		result = 0
	}
	return result
}
func nox_xxx_spriteGetMB_476F80() unsafe.Pointer {
	return dword_5d4594_1096640
}
func nox_xxx_clientGetSpriteAtCursor_476F90() unsafe.Pointer {
	return nox_client_spriteUnderCursorXxx_1096644
}
func nox_xxx_clientEnumHover_476FA0() {
	var v2 int4
	if *memmap.PtrUint32(6112660, 1096632) == 0 {
		*memmap.PtrUint32(6112660, 1096632) = uint32(nox_xxx_getNameId_4E3AA0(CString("Glyph")))
	}
	var mpos nox_point = getMousePos()
	sub_473970((*int2)(unsafe.Pointer(&mpos)), (*int2)(unsafe.Pointer(&mpos)))
	dword_5d4594_1096640 = nil
	nox_client_spriteUnderCursorXxx_1096644 = nil
	*memmap.PtrUint32(6112660, 1096628) = 0
	v2.field_0 = mpos.x - 96
	v2.field_8 = mpos.x + 96
	v2.field_C = mpos.y + 96
	v2.field_4 = mpos.y - 96
	dword_5d4594_1096636 = 0
	nox_xxx_forEachSprite_49AB00(&v2, unsafe.Pointer(cgoFuncAddr(nox_xxx_clientOnCursorHover_477050)), int32(uintptr(unsafe.Pointer(&mpos))))
}
func nox_xxx_clientOnCursorHover_477050(arg0 int32, a2 int32) {
	var (
		v2  int32
		v3  int32
		v4  int32
		v5  int32
		v6  *byte
		v7  int32
		v8  float32
		v9  int32
		v10 int32
		v11 int32
		v12 int32
		v13 int32
		v14 int32
		v15 float32
		v16 int32
		v17 int32
		v18 int32
		v19 int32
		v20 int32
		v21 int32
		v22 int32
		v23 float32
		v24 float32
		v25 float32
		v26 float32
		a3  float2
		a1  float2
		v29 int32
	)
	if *memmap.PtrUint32(6112660, 1096648) == 0 {
		*memmap.PtrUint32(6112660, 1096648) = uint32(nox_xxx_getTTByNameSpriteMB_44CFC0("Polyp"))
	}
	v2 = arg0
	if uint32(arg0) == *memmap.PtrUint32(0x852978, 8) {
		return
	}
	v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(arg0 + 120))))
	if !((v3&0x8000) == 0 && (!nox_xxx_spriteCheckFlag31_4356C0((*nox_drawable)(unsafe.Pointer(uintptr(arg0))), 0) || nox_xxx_spriteCheckFlag31_4356C0((*nox_drawable)(unsafe.Pointer(uintptr(*memmap.PtrInt32(0x852978, 8)))), 21))) {
		return
	}
	v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(arg0 + 112))))
	if !((v4&2) == 0 || (func() bool {
		v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(arg0 + 116))))
		return (v5 & 0x4000) == 0
	}())) {
		return
	}
	if !(uint32(v4)&0x80400206 != 0 || *(*uint32)(unsafe.Pointer(uintptr(arg0 + 108))) == *memmap.PtrUint32(6112660, 1096648)) {
		return
	}
	if nox_xxx_client_4984B0_drawable((*nox_drawable)(unsafe.Pointer(uintptr(arg0)))) == 0 {
		return
	}
	if !((int32(*(*uint8)(unsafe.Pointer(uintptr(arg0 + 112))))&4) == 0 || (func() *byte {
		v6 = (*byte)(unsafe.Pointer(noxServer.getPlayerByID(int32(*(*uint32)(unsafe.Pointer(uintptr(arg0 + 128)))))))
		return v6
	}()) != nil && (*(*byte)(unsafe.Add(unsafe.Pointer(v6), 3680))&1) == 0) {
		return
	}
	v7 = int32(*(*uint32)(unsafe.Pointer(uintptr(arg0 + 112))))
	if !(((uint32(v7)&0x400000) == 0 || (int32(*(*uint8)(unsafe.Pointer(uintptr(arg0 + 116))))&128) != 0) && ((v7&2) == 0 || *(*uint32)(unsafe.Pointer(uintptr(arg0 + 276))) != 10)) {
		return
	}
	v23 = float32(float64(*(*int32)(unsafe.Pointer(uintptr(arg0 + 16)))) - float64(*(*float32)(unsafe.Pointer(uintptr(arg0 + 100)))) - float64(*(*int16)(unsafe.Pointer(uintptr(arg0 + 104)))))
	v29 = int(v23)
	v24 = float32(float64(*(*int32)(unsafe.Pointer(uintptr(v2 + 16)))) - float64(*(*float32)(unsafe.Pointer(uintptr(v2 + 96)))) - float64(*(*int16)(unsafe.Pointer(uintptr(v2 + 104)))))
	v8 = COERCE_FLOAT(uint32(int(v24)))
	a3.field_0 = v8
	if *(*uint32)(unsafe.Pointer(uintptr(v2 + 44))) == 2 {
		v25 = *(*float32)(unsafe.Pointer(uintptr(v2 + 48))) * *(*float32)(unsafe.Pointer(uintptr(v2 + 48)))
		*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&a3.field_0))), unsafe.Sizeof(uint32(0))*0)) = uint32(int(v25))
		v17 = int(*(*float32)(unsafe.Pointer(uintptr(v2 + 48))))
		v18 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 12))))
		v19 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 12))) - uint32(v17))
		v20 = v18 + int(*(*float32)(unsafe.Pointer(uintptr(v2 + 48))))
		v21 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 4))))
		if v21 <= (*(*int32)(unsafe.Add(unsafe.Pointer((*int32)(unsafe.Pointer(&v8))), unsafe.Sizeof(int32(0))*0))) {
			v8 = *(*float32)(unsafe.Pointer(&v29))
			if v21 >= v29 {
				if *(*int32)(unsafe.Pointer(uintptr(a2))) <= v19 || *(*int32)(unsafe.Pointer(uintptr(a2))) >= v20 {
					return
				}
				goto LABEL_38
			}
		}
		v15 = a3.field_0
		v16 = int32((*(*uint32)(unsafe.Pointer(uintptr(a2)))-uint32(v18))*(*(*uint32)(unsafe.Pointer(uintptr(a2)))-uint32(v18)) + (uint32(v21)-(*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v8))), unsafe.Sizeof(uint32(0))*0))))*(uint32(v21)-(*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v8))), unsafe.Sizeof(uint32(0))*0)))))
	} else {
		if *(*uint32)(unsafe.Pointer(uintptr(v2 + 44))) != 3 {
			return
		}
		a1.field_0 = float32(float64(*(*int32)(unsafe.Pointer(uintptr(v2 + 12)))))
		a1.field_4 = float32(float64(*(*int32)(unsafe.Add(unsafe.Pointer((*int32)(unsafe.Pointer(&a3.field_0))), unsafe.Sizeof(int32(0))*0))))
		a3.field_0 = float32(float64(*(*int32)(unsafe.Pointer(uintptr(a2)))))
		a3.field_4 = float32(float64(*(*int32)(unsafe.Pointer(uintptr(a2 + 4)))))
		if nox_xxx_map_57B850(&a1, (*float32)(unsafe.Pointer(uintptr(v2+44))), &a3) != 0 || (func() int32 {
			a1.field_4 = float32(float64(v29))
			return nox_xxx_map_57B850(&a1, (*float32)(unsafe.Pointer(uintptr(v2+44))), &a3)
		}()) != 0 || (func() bool {
			v9 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 12))) + uint32(int(*(*float32)(unsafe.Pointer(uintptr(v2 + 72))))))
			v10 = v29 + int(*(*float32)(unsafe.Pointer(uintptr(v2 + 76))))
			v11 = int32((*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v8))), unsafe.Sizeof(uint32(0))*0))) + uint32(int(*(*float32)(unsafe.Pointer(uintptr(v2 + 76))))))
			return *(*uint32)(unsafe.Pointer(uintptr(a2))) > uint32(v9)
		}()) && *(*uint32)(unsafe.Pointer(uintptr(a2))) < uint32(*(*int32)(unsafe.Pointer(uintptr(v2 + 12)))) && (func() bool {
			v12 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 4))))
			return v12 > v10
		}()) && v12 < v11 {
			goto LABEL_38
		}
		v13 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 12))) + uint32(int(*(*float32)(unsafe.Pointer(uintptr(v2 + 80))))))
		v14 = v29 + int(*(*float32)(unsafe.Pointer(uintptr(v2 + 84))))
		*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v15))), unsafe.Sizeof(uint32(0))*0)) = (*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v8))), unsafe.Sizeof(uint32(0))*0))) + uint32(int(*(*float32)(unsafe.Pointer(uintptr(v2 + 84)))))
		if *(*int32)(unsafe.Pointer(uintptr(a2))) < *(*int32)(unsafe.Pointer(uintptr(v2 + 12))) {
			return
		}
		if *(*int32)(unsafe.Pointer(uintptr(a2))) >= v13 {
			return
		}
		v16 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 4))))
		if v16 <= v14 {
			return
		}
	}
	if v16 >= (*(*int32)(unsafe.Add(unsafe.Pointer((*int32)(unsafe.Pointer(&v15))), unsafe.Sizeof(int32(0))*0))) {
		return
	}
LABEL_38:
	v26 = float32(float64(*(*int16)(unsafe.Pointer(uintptr(v2 + 104)))) + float64(*(*int32)(unsafe.Pointer(uintptr(v2 + 16)))) + float64(*(*float32)(unsafe.Pointer(uintptr(v2 + 96)))))
	v22 = int(v26)
	if v22 > *memmap.PtrInt32(6112660, 1096628) {
		*memmap.PtrUint32(6112660, 1096628) = uint32(v22)
		dword_5d4594_1096640 = unsafe.Pointer(uintptr(v2))
	}
	if uint32(v2) != *memmap.PtrUint32(0x852978, 8) && v22 > *(*int32)(unsafe.Pointer(&dword_5d4594_1096636)) && nox_xxx_client_57B400(v2) != 0 {
		if *memmap.PtrUint32(0x8531A0, 2576) != 0 && int32(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(0x8531A0, 2576) + 2251)))) == 1 && *(*uint32)(unsafe.Pointer(uintptr(v2 + 108))) == *memmap.PtrUint32(6112660, 1096632) {
			if nox_client_spriteUnderCursorXxx_1096644 == nil {
				nox_client_spriteUnderCursorXxx_1096644 = unsafe.Pointer(uintptr(v2))
				dword_5d4594_1096636 = 0
			}
		} else {
			dword_5d4594_1096636 = uint32(v22)
			nox_client_spriteUnderCursorXxx_1096644 = unsafe.Pointer(uintptr(v2))
		}
	}
}
func nox_video_drawCursorSelectCircle2_477470(a1 *uint32, a2 int32, a3 int32) {
	var (
		v3 int32
		v4 int32
		v5 int32
		v6 int32
		v7 float32
	)
	if a3 != 0 {
		if *(*uint32)(unsafe.Pointer(uintptr(a3 + 44))) == 2 {
			v4 = int(*(*float32)(unsafe.Pointer(uintptr(a3 + 48))))
		} else {
			if *(*uint32)(unsafe.Pointer(uintptr(a3 + 44))) != 3 {
				v3 = 6
			LABEL_8:
				v5 = int32(*a1 + *(*uint32)(unsafe.Pointer(uintptr(a3 + 12))) - *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*4)))
				v6 = int32(*(*uint32)(unsafe.Pointer(uintptr(a3 + 16))) + *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*1)) - *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*5)))
				nox_client_drawSetColor_434460(a2)
				nox_client_drawEnableAlpha_434560(1)
				nox_client_drawSetAlpha_434580(128)
				nox_video_drawCircle_4B0B90(v5, v6, v3-1)
				nox_video_drawCircle_4B0B90(v5, v6, v3+1)
				nox_client_drawEnableAlpha_434560(0)
				nox_video_drawCircle_4B0B90(v5, v6, v3)
				return
			}
			v7 = *(*float32)(unsafe.Pointer(uintptr(a3 + 60))) + *(*float32)(unsafe.Pointer(uintptr(a3 + 56)))
			v4 = int(v7) / 2
		}
		v3 = v4 + 6
		goto LABEL_8
	}
}
func nox_xxx_guiCursor_477600() int32 {
	return int32(*memmap.PtrUint32(6112660, 1096672))
}
func nox_xxx_cursorSetTooltip_4776B0(a1 *libc.WChar) {
	if a1 != nil {
		if int32(nox_wcslen(a1)) >= 256 {
			nox_wcsncpy((*libc.WChar)(memmap.PtrOff(6112660, 1096676)), a1, math.MaxUint8)
			*memmap.PtrUint16(6112660, 1097186) = 0
		} else {
			nox_wcscpy((*libc.WChar)(memmap.PtrOff(6112660, 1096676)), a1)
		}
	} else {
		*memmap.PtrUint16(6112660, 1096676) = 0
	}
}
func sub_478030() int32 {
	return int32(dword_5d4594_1098624)
}
func sub_478040() int32 {
	var v2 int16
	if dword_5d4594_1098624 == 0 {
		return 0
	}
	v2 = 4809
	nox_netlist_addToMsgListCli_40EBC0(31, 0, (*uint8)(unsafe.Pointer(&v2)), 2)
	sub_467680()
	return 1
}
func sub_478080(a1 int32) int32 {
	var (
		v1     *byte
		result int32
	)
	if dword_5d4594_1098624 != 0 && (func() *byte {
		v1 = sub_4780A0(a1)
		return v1
	}()) != nil {
		result = int32(*(*uint32)(unsafe.Pointer(v1)))
	} else {
		result = 0
	}
	return result
}
func sub_4780A0(a1 int32) *byte {
	var (
		v1 int32
		v2 *uint8
		v3 int32
		v4 *uint8
		v5 int32
		v6 *uint8
	)
	v1 = 0
	v2 = (*uint8)(memmap.PtrOff(6112660, 1098644))
	for {
		v3 = 0
		v4 = v2
		for {
			if *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v4))), -int(unsafe.Sizeof(uint32(0))*1)))) != 0 {
				v5 = 0
				v6 = v4
				for *(*uint32)(unsafe.Pointer(v6)) != uint32(a1) {
					v5++
					v6 = (*uint8)(unsafe.Add(unsafe.Pointer(v6), 4))
					if v5 >= 32 {
						goto LABEL_7
					}
				}
				return (*byte)(memmap.PtrOff(6112660, uint32((v1+v3*10)*140)+0x10C38C))
			}
		LABEL_7:
			v3++
			v4 = (*uint8)(unsafe.Add(unsafe.Pointer(v4), 1400))
			if v3 >= 6 {
				break
			}
		}
		v2 = (*uint8)(unsafe.Add(unsafe.Pointer(v2), 140))
		v1++
		if int32(uintptr(unsafe.Pointer(v2))) < int32(uintptr(memmap.PtrOff(6112660, 1100044))) {
			continue
		}
		return nil
	}
}
func sub_478110() int32 {
	var (
		v0     *uint32
		result int32
		v2     *uint32
		v3     *uint8
		v4     *uint8
		v5     int32
		v6     *uint32
		v7     int32
		v8     int32
	)
	*memmap.PtrUint32(6112660, 1098500) = uint32(nox_win_width)
	*memmap.PtrUint32(6112660, 1098504) = uint32(nox_win_height)
	*memmap.PtrUint32(6112660, 1098524) = uint32(nox_win_width)
	*memmap.PtrUint32(6112660, 1098492) = 0
	*memmap.PtrUint32(6112660, 1098496) = 0
	*memmap.PtrUint32(6112660, 1098528) = uint32(nox_win_height)
	*memmap.PtrUint32(6112660, 1098508) = 0
	*memmap.PtrUint32(6112660, 1098512) = 0
	v0 = (*uint32)(unsafe.Pointer(newWindowFromFile("Shop.wnd", unsafe.Pointer(cgoFuncAddr(sub_478480)))))
	dword_5d4594_1098576 = uint32(uintptr(unsafe.Pointer(v0)))
	if v0 == nil {
		return 0
	}
	(*nox_window)(unsafe.Pointer(v0)).SetAllFuncs(func(arg1 int32, arg2 int32, arg3 int32, arg4 int32) int32 {
		return sub_478650(arg1, arg2, uint32(arg3))
	}, func(arg1 *nox_window, arg2 unsafe.Pointer) int32 {
		return sub_478970()
	}, nil)
	v2 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1098576)))).ChildByID(3806)))
	(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v2)))))).setTooltipFunc(unsafe.Pointer(cgoFuncAddr(sub_478E50)))
	v3 = (*uint8)(memmap.PtrOff(6112660, 1098636))
	for {
		v4 = v3
		v5 = 6
		for {
			*(*uint32)(unsafe.Pointer(v4)) = 0
			v4 = (*uint8)(unsafe.Add(unsafe.Pointer(v4), 1400))
			v5--
			if v5 == 0 {
				break
			}
		}
		v3 = (*uint8)(unsafe.Add(unsafe.Pointer(v3), 140))
		if int32(uintptr(unsafe.Pointer(v3))) >= int32(uintptr(memmap.PtrOff(6112660, 1100036))) {
			break
		}
	}
	sub_478F10()
	nox_window_get_size((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1098576))))), &v8, &v7)
	(*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1098576)))).SetPos(image.Pt(nox_win_width-v8, nox_win_height-v7))
	(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1098576))))).Hide()
	nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1098576))))), 0)
	*memmap.PtrUint32(6112660, 1098400) = uint32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg("ShopBase"))))
	*memmap.PtrUint32(6112660, 1098404) = uint32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg("ShopTradeMode"))))
	*memmap.PtrUint32(6112660, 1098408) = uint32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg("ShopIdentifyMode"))))
	*memmap.PtrUint32(6112660, 1098412) = uint32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg("ShopRepairMode"))))
	*memmap.PtrUint32(6112660, 1098416) = uint32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg("ShopRepairMode"))))
	*memmap.PtrUint32(6112660, 1098420) = uint32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg("ShopExitMode"))))
	*memmap.PtrUint32(6112660, 1098424) = uint32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg("ShopInventoryBar1"))))
	*memmap.PtrUint32(6112660, 1098428) = uint32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg("ShopInventoryBar2"))))
	*memmap.PtrUint32(6112660, 1098432) = uint32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg("ShopInventorySlider"))))
	*memmap.PtrUint32(6112660, 1098436) = uint32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg("ShopInventorySliderSelected"))))
	*memmap.PtrUint32(6112660, 1098448) = uint32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg("ShopInventoryUp"))))
	*memmap.PtrUint32(6112660, 1098452) = uint32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg("ShopInventoryUpSelected"))))
	*memmap.PtrUint32(6112660, 1098440) = uint32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg("ShopInventorydown"))))
	*memmap.PtrUint32(6112660, 1098444) = uint32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg("ShopInventorydownSelected"))))
	dword_5d4594_1098456 = uint32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg("ShopTextBorder"))))
	*memmap.PtrUint32(6112660, 1098460) = uint32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg("ShopkeeperPic"))))
	*memmap.PtrUint32(6112660, 1098464) = uint32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg("ShopkeeperWarriorPic"))))
	*memmap.PtrUint32(6112660, 1098468) = uint32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg("ShopkeeperConjurerPic"))))
	*memmap.PtrUint32(6112660, 1098472) = uint32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg("ShopkeeperWizardPic"))))
	*memmap.PtrUint32(6112660, 1098476) = uint32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg("ShopkeeperLandOfTheDeadPic"))))
	*memmap.PtrUint32(6112660, 1098480) = uint32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg("ShopkeeperMagicShopPic"))))
	dword_5d4594_1098580 = uint32(uintptr(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1098576)))).ChildByID(3807))))
	*memmap.PtrUint32(6112660, 1098584) = uint32(uintptr(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1098576)))).ChildByID(3808))))
	*memmap.PtrUint32(6112660, 1098588) = uint32(uintptr(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1098576)))).ChildByID(3809))))
	v6 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1098576)))).ChildByID(3806)))
	nox_client_wndGetPosition_46AA60((*nox_window)(unsafe.Pointer(v6)), (*uint32)(memmap.PtrOff(6112660, 1098380)), (*uint32)(memmap.PtrOff(6112660, 1098384)))
	nox_window_get_size((*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v6)))))), (*int32)(memmap.PtrOff(6112660, 1098388)), (*int32)(memmap.PtrOff(6112660, 1098392)))
	*memmap.PtrUint32(6112660, 1098388) += *memmap.PtrUint32(6112660, 1098380)
	*memmap.PtrUint32(6112660, 1098392) += *memmap.PtrUint32(6112660, 1098384)
	*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1098580 + 400))) + 8))) = 16
	*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1098580 + 400))) + 12))) = 12
	nox_xxx_wndSetOffsetMB_46AE40(int32(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1098580 + 400)))), 0, -15)
	sub_4B5700(*(*int32)(unsafe.Pointer(&dword_5d4594_1098580)), 0, 0, *memmap.PtrInt32(6112660, 1098432), *memmap.PtrInt32(6112660, 1098436), *memmap.PtrInt32(6112660, 1098436))
	result = 1
	dword_5d4594_1098592 = *(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1098580 + 32))) + 4)))
	return result
}
func sub_478480(a1 int32, a2 int32, a3 *int32, a4 int32) int32 {
	var (
		result int32
		v5     int32
		v7     int32
		v8     int32
	)
	if a2 != 0x4007 {
		if a2 == 0x4009 {
			dword_5d4594_1107036 = dword_5d4594_1098592 - uint32(a4)
			return 0
		}
		return 0
	}
	if sub_45D9B0() != 0 {
		return 0
	}
	v5 = (*nox_window)(unsafe.Pointer(a3)).ID()
	clientPlaySoundSpecial(766, 100)
	switch v5 {
	case 3801:
		if dword_5d4594_1098624 == 0 {
			return 0
		}
		sub_478040()
		result = 0
	case 3802:
		if dword_5d4594_1098624 == 0 {
			return 0
		}
		if dword_5d4594_1098628 == 4 {
			sub_467680()
		}
		nox_client_setCursorType_477610(12)
		dword_5d4594_1098628 = 3
		if sub_467C80() != 0 {
			return 0
		}
		sub_467BB0()
		result = 0
	case 3803:
		if dword_5d4594_1098624 == 0 {
			return 0
		}
		sub_467650()
		dword_5d4594_1098628 = 4
		result = 0
	case 3804:
		if dword_5d4594_1098624 != 0 {
			if dword_5d4594_1098628 == 4 {
				sub_467680()
			}
			nox_client_setCursorType_477610(11)
			dword_5d4594_1098628 = 2
		}
		return 0
	case 3808:
		if *(*int32)(unsafe.Pointer(&dword_5d4594_1107036))-50 >= 0 {
			v8 = int32(dword_5d4594_1107036 - 50 - (dword_5d4594_1107036-50)%50)
		} else {
			v8 = 0
		}
		dword_5d4594_1107036 = uint32(v8)
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1098580))))).Func94(asWindowEvent(0x400A, int32(dword_5d4594_1098592-uint32(v8)), 0))
		result = 0
	case 3809:
		if *(*int32)(unsafe.Pointer(&dword_5d4594_1107036))+50 <= *(*int32)(unsafe.Pointer(&dword_5d4594_1098592)) {
			v7 = int32(dword_5d4594_1107036 + 50 - (dword_5d4594_1107036+50)%50)
			dword_5d4594_1107036 = dword_5d4594_1107036 + 50 - (dword_5d4594_1107036+50)%50
		} else {
			v7 = int32(dword_5d4594_1098592)
			dword_5d4594_1107036 = dword_5d4594_1098592
		}
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1098580))))).Func94(asWindowEvent(0x400A, int32(dword_5d4594_1098592-uint32(v7)), 0))
		result = 0
	default:
		return 0
	}
	return result
}
func sub_478650(a1 int32, a2 int32, a3 uint32) int32 {
	var (
		v4 int32
		v5 int2
	)
	v5.field_0 = int32(uint16(a3))
	v5.field_4 = int32(a3 >> 16)
	if sub_45D9B0() == 0 {
		switch a2 {
		case 5:
			v4 = sub_4281F0(&v5, (*int4)(memmap.PtrOff(6112660, 1098380)))
			if v4 != 0 {
				if dword_5d4594_1098628 == 2 {
					sub_478730(&v5.field_0)
				}
			}
		case 19:
			if dword_5d4594_1098628 == 2 {
				(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1098576))))).Func94(asWindowEvent(0x4007, *memmap.PtrInt32(6112660, 1098584), 0))
				return 1
			}
		case 20:
			if dword_5d4594_1098628 == 2 {
				(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1098576))))).Func94(asWindowEvent(0x4007, *memmap.PtrInt32(6112660, 1098588), 0))
				return 1
			}
		default:
			return 0
		}
	}
	return 1
}
func sub_478850(a1 int32, a2 int16, a3 int32, a4 int32) {
	if a4 != 0 {
		if a4 == 1 {
			nox_client_tradeXxxBuyAccept_478880(a3, a2)
		} else {
			sub_4788F0(a3, a4)
		}
	}
}
func sub_478970() int32 {
	var (
		result int32
		a1     int2
	)
	a1.field_0 = nox_win_width - NOX_DEFAULT_WIDTH
	a1.field_4 = nox_win_height - NOX_DEFAULT_HEIGHT
	nox_client_drawImageAt_47D2C0((*nox_video_bag_image_t)(unsafe.Pointer(uintptr(*memmap.PtrInt32(6112660, 1098400)))), nox_win_width-NOX_DEFAULT_WIDTH, nox_win_height-NOX_DEFAULT_HEIGHT)
	if dword_5d4594_1098628 == 2 {
		nox_client_drawImageAt_47D2C0((*nox_video_bag_image_t)(unsafe.Pointer(uintptr(*memmap.PtrInt32(6112660, 1098404)))), a1.field_0, a1.field_4)
		sub_478C80()
		result = 1
	} else if dword_5d4594_1098628 == 3 {
		nox_client_drawImageAt_47D2C0((*nox_video_bag_image_t)(unsafe.Pointer(uintptr(*memmap.PtrInt32(6112660, 1098408)))), a1.field_0, a1.field_4)
		sub_478B10(&a1)
		result = 1
	} else if dword_5d4594_1098628 == 1 {
		nox_client_drawImageAt_47D2C0((*nox_video_bag_image_t)(unsafe.Pointer(uintptr(*memmap.PtrInt32(6112660, 1098412)))), a1.field_0, a1.field_4)
		sub_478A70(&a1)
		result = 1
	} else {
		if dword_5d4594_1098628 == 4 {
			nox_client_drawImageAt_47D2C0((*nox_video_bag_image_t)(unsafe.Pointer(uintptr(*memmap.PtrInt32(6112660, 1098416)))), a1.field_0, a1.field_4)
			sub_478BC0(&a1.field_0)
		}
		result = 1
	}
	return result
}
func sub_478A70(a1 *int2) int32 {
	var (
		v1     *uint32
		result int32
		v3     int32
		v4     int32
		v5     int32
		v6     int32
	)
	v1 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1098576)))).ChildByID(3806)))
	nox_client_wndGetPosition_46AA60((*nox_window)(unsafe.Pointer(v1)), (*uint32)(unsafe.Pointer(&v5)), (*uint32)(unsafe.Pointer(&v6)))
	nox_window_get_size((*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v1)))))), &v4, &v3)
	nox_xxx_drawSetTextColor_434390(int32(nox_color_white_2523948))
	nox_client_drawImageAt_47D2C0((*nox_video_bag_image_t)(unsafe.Pointer(uintptr(dword_5d4594_1098456))), a1.field_0, a1.field_4)
	result = int32(dword_5d4594_1098604)
	if dword_5d4594_1098604 != 0 {
		result = nox_xxx_drawStringWrap_43FAF0(nil, (*libc.WChar)(unsafe.Pointer(*(**uint16)(unsafe.Pointer(&dword_5d4594_1098604)))), v5+8, v6+8, v4-16, v3-16)
	}
	return result
}
func sub_478C80() int32 {
	var (
		v0  int32
		v1  int32
		v2  *uint8
		v3  int32
		v4  int32
		v6  int32
		v7  *uint8
		v8  uint32
		v9  int32
		v10 int32
		v11 [32]libc.WChar
	)
	v8 = uint32(sub_4674A0())
	nox_xxx_wndDraw_49F7F0()
	nox_client_copyRect_49F6F0(*memmap.PtrInt32(6112660, 1098380), *memmap.PtrInt32(6112660, 1098384), int32(*memmap.PtrUint32(6112660, 1098388)-*memmap.PtrUint32(6112660, 1098380)), int32(*memmap.PtrUint32(6112660, 1098392)-*memmap.PtrUint32(6112660, 1098384)))
	v0 = int32(*memmap.PtrUint32(6112660, 1098380))
	v10 = int32(*memmap.PtrUint32(6112660, 1098380))
	v1 = int32(*memmap.PtrUint32(6112660, 1098384) - dword_5d4594_1107036)
	v2 = (*uint8)(memmap.PtrOff(6112660, 1098640))
	v9 = nox_xxx_guiFontHeightMB_43F320(nil)
	v6 = 0
	v7 = (*uint8)(memmap.PtrOff(6112660, 1098640))
	for {
		if v1 > *memmap.PtrInt32(6112660, 1098384)-50 {
			nox_client_drawImageAt_47D2C0((*nox_video_bag_image_t)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, uint32((v6%2)*4)+0x10C2B8)))), v0, v1)
			v3 = v0 + 5
			v4 = 6
			for {
				if *(*uint32)(unsafe.Pointer(v2)) != 0 {
					*(*uint32)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), -int(unsafe.Sizeof(uint32(0))*1)))) + 12))) = uint32(v3 + 20)
					*(*uint32)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), -int(unsafe.Sizeof(uint32(0))*1)))) + 16))) = uint32(v1 + 25)
					(cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), -int(unsafe.Sizeof(uint32(0))*1)))) + 300))), (*func(*uint8, uint32))(nil)))((*uint8)(memmap.PtrOff(6112660, 1098492)), *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), -int(unsafe.Sizeof(uint32(0))*1)))))
					if v8 < *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), unsafe.Sizeof(uint32(0))*33))) {
						nox_client_drawRectFilledAlpha_49CF10(v3-5, v1, 50, 50)
					}
					if *(*uint32)(unsafe.Pointer(v2)) > 1 {
						nox_swprintf(&v11[0], libc.CWString("%d"), *(*uint32)(unsafe.Pointer(v2)))
						nox_xxx_drawSetTextColor_434390(int32(nox_color_white_2523948))
						nox_xxx_drawStringWrap_43FAF0(nil, &v11[0], v3, v1+5, 320, 0)
					}
					nox_swprintf(&v11[0], libc.CWString("%d"), *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), unsafe.Sizeof(uint32(0))*33))))
					nox_xxx_drawSetTextColor_434390(int32(nox_color_yellow_2589772))
					nox_xxx_drawStringWrap_43FAF0(nil, &v11[0], v3, v1-v9+45, 320, 0)
				}
				v3 += 50
				v2 = (*uint8)(unsafe.Add(unsafe.Pointer(v2), 1400))
				v4--
				if v4 == 0 {
					break
				}
			}
			v2 = v7
			v0 = v10
		}
		v1 += 50
		if v1 >= *memmap.PtrInt32(6112660, 1098392) {
			break
		}
		v2 = (*uint8)(unsafe.Add(unsafe.Pointer(v2), 140))
		v6++
		v7 = v2
		if int32(uintptr(unsafe.Pointer(v2))) >= int32(uintptr(memmap.PtrOff(6112660, 1100040))) {
			break
		}
	}
	return sub_49F860()
}
func sub_478E50(a1 int32, a2 int32, a3 uint32) int32 {
	var (
		v3 int32
		v4 int32
		v5 int32
		v6 **libc.WChar
		v7 *libc.WChar
	)
	if dword_5d4594_1098628 == 2 {
		v3 = int32((uint32(uint16(a3)) - *memmap.PtrUint32(6112660, 1098380)) / 50)
		v4 = int32((a3>>16)-*memmap.PtrUint32(6112660, 1098384)+dword_5d4594_1107036) / 50
		if v3 >= 6 {
			v3 = 5
		}
		if v4 >= 10 {
			v4 = 9
		}
		v5 = (v4 + v3*10) * 35
		v6 = (**libc.WChar)(memmap.PtrOff(6112660, uint32(v5*4)+0x10C38C))
		if *memmap.PtrUint32(6112660, uint32(v5*4)+0x10C390) != 0 {
			*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(*v6))), unsafe.Sizeof(uint32(0))*32))) = uint32(uintptr(unsafe.Pointer(*(**libc.WChar)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof((*libc.WChar)(nil))*2)))))
			v7 = nox_xxx_clientAskInfoMb_4BF050((*nox_drawable)(unsafe.Pointer(*v6)))
			nox_xxx_cursorSetTooltip_4776B0(v7)
		}
	}
	return 1
}
func sub_478F10() *uint32 {
	var (
		v0     *uint8
		v1     int32
		v2     unsafe.Pointer
		result *uint32
		v4     *uint8
	)
	v4 = (*uint8)(memmap.PtrOff(6112660, 1098636))
	for {
		v0 = v4
		v1 = 6
		for {
			if *(*uint32)(unsafe.Pointer(v0)) != 0 {
				nox_xxx_spriteDelete_45A4B0((*nox_drawable)(unsafe.Pointer(*(**uint64)(unsafe.Pointer(v0)))))
			}
			*(*uint32)(unsafe.Pointer(v0)) = 0
			v2 = unsafe.Add(unsafe.Pointer(v0), 8)
			*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v0))), unsafe.Sizeof(uint32(0))*1))) = 0
			v0 = (*uint8)(unsafe.Add(unsafe.Pointer(v0), 1400))
			libc.MemSet(v2, 0, 128)
			*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v0))), -int(unsafe.Sizeof(uint32(0))*316)))) = 0
			v1--
			if v1 == 0 {
				break
			}
		}
		result = (*uint32)(unsafe.Add(unsafe.Pointer(v4), 140))
		v4 = (*uint8)(unsafe.Add(unsafe.Pointer(v4), 140))
		if int32(uintptr(unsafe.Pointer(v4))) >= int32(uintptr(memmap.PtrOff(6112660, 1100036))) {
			break
		}
	}
	dword_5d4594_1107036 = 0
	return result
}
func sub_478F80() int32 {
	var result int32
	sub_478F10()
	sub_44D8F0()
	(*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1098576)))).Destroy()
	result = 0
	dword_5d4594_1098576 = 0
	dword_5d4594_1098624 = 0
	dword_5d4594_1098628 = 1
	dword_5d4594_1098596 = 0
	dword_5d4594_1098600 = 0
	dword_5d4594_1098604 = 0
	*memmap.PtrUint32(6112660, 1098608) = 0
	dword_5d4594_1098616 = 0
	return result
}
func nox_xxx_getShopPic_4790F0(a1 int32) *byte {
	var (
		v1     *uint32
		result *byte
	)
	v1 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1098576)))).ChildByID(3805)))
	if *memmap.PtrUint32(6112660, 1107040) == 0 {
		*memmap.PtrUint32(6112660, 1098396) = uint32(nox_xxx_getTTByNameSpriteMB_44CFC0("Shopkeeper"))
		*memmap.PtrUint32(6112660, 1098560) = uint32(nox_xxx_getTTByNameSpriteMB_44CFC0("ShopkeeperWarriorsRealm"))
		*memmap.PtrUint32(6112660, 1098556) = uint32(nox_xxx_getTTByNameSpriteMB_44CFC0("ShopkeeperConjurerRealm"))
		*memmap.PtrUint32(6112660, 1098564) = uint32(nox_xxx_getTTByNameSpriteMB_44CFC0("ShopkeeperWizardRealm"))
		*memmap.PtrUint32(6112660, 1098572) = uint32(nox_xxx_getTTByNameSpriteMB_44CFC0("ShopkeeperLandOfTheDead"))
		*memmap.PtrUint32(6112660, 1098568) = uint32(nox_xxx_getTTByNameSpriteMB_44CFC0("ShopkeeperMagicShop"))
		*memmap.PtrUint32(6112660, 1098484) = uint32(nox_xxx_getTTByNameSpriteMB_44CFC0("ShopkeeperPurple"))
		*memmap.PtrUint32(6112660, 1097292) = uint32(nox_xxx_getTTByNameSpriteMB_44CFC0("ShopkeeperYellow"))
		*memmap.PtrUint32(6112660, 1107040) = 1
	}
	if uint32(a1) == *memmap.PtrUint32(6112660, 1098396) {
		result = (*byte)(unsafe.Pointer(nox_xxx_gLoadImg("ShopKeeperPic")))
		*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*15)) = uint32(uintptr(unsafe.Pointer(result)))
	} else if uint32(a1) == *memmap.PtrUint32(6112660, 1098560) {
		result = (*byte)(unsafe.Pointer(nox_xxx_gLoadImg("ShopKeeperWarriorPic")))
		*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*15)) = uint32(uintptr(unsafe.Pointer(result)))
	} else if uint32(a1) == *memmap.PtrUint32(6112660, 1098556) {
		result = (*byte)(unsafe.Pointer(nox_xxx_gLoadImg("ShopKeeperConjurerPic")))
		*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*15)) = uint32(uintptr(unsafe.Pointer(result)))
	} else if uint32(a1) == *memmap.PtrUint32(6112660, 1098564) {
		result = (*byte)(unsafe.Pointer(nox_xxx_gLoadImg("ShopKeeperWizardPic")))
		*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*15)) = uint32(uintptr(unsafe.Pointer(result)))
	} else if uint32(a1) == *memmap.PtrUint32(6112660, 1098572) {
		result = (*byte)(unsafe.Pointer(nox_xxx_gLoadImg("ShopKeeperLandOfTheDeadPic")))
		*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*15)) = uint32(uintptr(unsafe.Pointer(result)))
	} else if uint32(a1) == *memmap.PtrUint32(6112660, 1098568) {
		result = (*byte)(unsafe.Pointer(nox_xxx_gLoadImg("ShopKeeperMagicShopPic")))
		*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*15)) = uint32(uintptr(unsafe.Pointer(result)))
	} else if uint32(a1) == *memmap.PtrUint32(6112660, 1098484) {
		result = (*byte)(unsafe.Pointer(nox_xxx_gLoadImg("ShopKeeperPurplePic")))
		*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*15)) = uint32(uintptr(unsafe.Pointer(result)))
	} else {
		if uint32(a1) == *memmap.PtrUint32(6112660, 1097292) {
			result = (*byte)(unsafe.Pointer(nox_xxx_gLoadImg("ShopKeeperBrownPic")))
		} else {
			result = (*byte)(unsafe.Pointer(nox_xxx_gLoadImg("ShopKeeperPic")))
		}
		*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*15)) = uint32(uintptr(unsafe.Pointer(result)))
	}
	return result
}
func sub_479280() int32 {
	var result int32
	result = int32(dword_5d4594_1098624)
	if dword_5d4594_1098624 != 0 {
		sub_467680()
		dword_5d4594_1098624 = 0
		dword_5d4594_1098628 = 0
		sub_478F10()
		sub_44D8F0()
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1098576))))).Hide()
		nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1098576))))), 0)
		sub_467C10()
		nox_client_setCursorType_477610(0)
		result = int32(nox_client_renderGUI_80828)
		if nox_client_renderGUI_80828 == 0 && *memmap.PtrUint32(6112660, 1098612) == 1 {
			nox_client_renderGUI_80828 = 1
		}
	}
	return result
}
func sub_479300(a1 int32, a2 int32, a3 int32, a4 int16, a5 int32) *uint32 {
	var (
		result *uint32
		v6     *uint32
		v7     *int32
		i      int32
	)
	result = (*uint32)(unsafe.Pointer(sub_4793C0(a1)))
	v6 = result
	if result != nil {
		if *result == 0 {
			result = &nox_new_drawable_for_thing(a1).field_0
			*v6 = uint32(uintptr(unsafe.Pointer(result)))
			if result == nil {
				return result
			}
			*(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*30)) |= 0x40000000
			*(*uint16)(unsafe.Pointer(uintptr(*v6 + 294))) = uint16(a4)
			*(*uint16)(unsafe.Pointer(uintptr(*v6 + 292))) = uint16(a4)
			if *(*uint32)(unsafe.Pointer(uintptr(*v6 + 112)))&0x13001000 != 0 {
				v7 = (*int32)(unsafe.Pointer(uintptr(*v6 + 432)))
				for i = 0; i < 4; i++ {
					if *(*byte)(unsafe.Pointer(uintptr(i + a5))) == math.MaxUint8 {
						*v7 = 0
					} else {
						*v7 = int32(uintptr(unsafe.Pointer(nox_xxx_modifGetDescById_413330(int32(*(*uint8)(unsafe.Pointer(uintptr(i + a5))))))))
					}
					v7 = (*int32)(unsafe.Add(unsafe.Pointer(v7), unsafe.Sizeof(int32(0))*1))
				}
			}
			*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*1)) = 0
		}
		*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*1))+2))) = uint32(a2)
		result = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*1)) + 1)))
		*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*34)) = uint32(a3)
		*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*1)) = uint32(uintptr(unsafe.Pointer(result)))
	}
	return result
}
func sub_4793C0(a1 int32) *byte {
	var (
		v1 int32
		v2 *uint8
		v3 int32
		v4 *uint8
	)
	v1 = 0
	v2 = (*uint8)(memmap.PtrOff(6112660, 1098636))
	for {
		v3 = 0
		v4 = v2
		for {
			if *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v4))), unsafe.Sizeof(uint32(0))*1))) != 0 && *(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(v4)) + 108))) == uint32(a1) && (*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(v4)) + 112)))&0x4000000) == 0 {
				return (*byte)(memmap.PtrOff(6112660, uint32((v1+v3*10)*140)+0x10C38C))
			}
			v3++
			v4 = (*uint8)(unsafe.Add(unsafe.Pointer(v4), 1400))
			if v3 >= 6 {
				break
			}
		}
		v2 = (*uint8)(unsafe.Add(unsafe.Pointer(v2), 140))
		v1++
		if int32(uintptr(unsafe.Pointer(v2))) < int32(uintptr(memmap.PtrOff(6112660, 1100036))) {
			continue
		}
		break
	}
	return sub_479430()
}
func sub_479430() *byte {
	var (
		v0 int32
		v1 *uint8
		v2 int32
		v3 *uint8
	)
	v0 = 0
	v1 = (*uint8)(memmap.PtrOff(6112660, 1098640))
	for {
		v2 = 0
		v3 = v1
		for {
			if *(*uint32)(unsafe.Pointer(v3)) == 0 {
				return (*byte)(memmap.PtrOff(6112660, uint32((v0+v2*10)*140)+0x10C38C))
			}
			v2++
			v3 = (*uint8)(unsafe.Add(unsafe.Pointer(v3), 1400))
			if v2 >= 6 {
				break
			}
		}
		v1 = (*uint8)(unsafe.Add(unsafe.Pointer(v1), 140))
		v0++
		if int32(uintptr(unsafe.Pointer(v1))) < int32(uintptr(memmap.PtrOff(6112660, 1100040))) {
			continue
		}
		break
	}
	return nil
}
func sub_479480(a1 int32) *byte {
	var (
		result *byte
		v2     *byte
	)
	result = sub_4780A0(a1)
	v2 = result
	if result != nil {
		sub_4794D0(int32(uintptr(unsafe.Pointer(result))), a1)
		result = (*byte)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), unsafe.Sizeof(uint32(0))*1))) - 1)))
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), unsafe.Sizeof(uint32(0))*1))) = uint32(uintptr(unsafe.Pointer(result)))
		if result == nil {
			result = (*byte)(unsafe.Pointer(uintptr(nox_xxx_spriteDelete_45A4B0((*nox_drawable)(unsafe.Pointer(*(**uint64)(unsafe.Pointer(v2))))))))
			*(*uint32)(unsafe.Pointer(v2)) = 0
			*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), unsafe.Sizeof(uint32(0))*34))) = 0
		}
	}
	return result
}
func sub_4794D0(a1 int32, a2 int32) int32 {
	var (
		result int32
		i      *uint32
		v4     *uint32
		v5     int32
	)
	result = 0
	for i = (*uint32)(unsafe.Pointer(uintptr(a1 + 8))); *i != uint32(a2); i = (*uint32)(unsafe.Add(unsafe.Pointer(i), unsafe.Sizeof(uint32(0))*1)) {
		if func() int32 {
			p := &result
			*p++
			return *p
		}() >= 32 {
			return result
		}
	}
	if result < 31 {
		v4 = (*uint32)(unsafe.Pointer(uintptr(a1 + result*4 + 8)))
		v5 = 31 - result
		for {
			result = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*1)))
			*v4 = uint32(result)
			v4 = (*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*1))
			v5--
			if v5 == 0 {
				break
			}
		}
	}
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 132))) = 0
	return result
}
func sub_479590() int32 {
	return int32(dword_5d4594_1098628)
}
func sub_4795A0(a1 int32) {
	if dword_5d4594_1098628 == 4 {
		if a1 != 4 {
			sub_467680()
		}
		dword_5d4594_1098628 = uint32(a1)
	} else {
		dword_5d4594_1098628 = uint32(a1)
	}
}
func sub_479690(a1 int32, a2 int16, a3 int16, a4 int32) int32 {
	var result int32
	result = a4
	dword_5d4594_1098616 = 0
	if a4 != 0 {
		if a4 == 1 {
			result = nox_client_tradeXxxSellAccept_4796D0(a2)
		} else {
			result = sub_479700(a3, int8(a4))
		}
	}
	return result
}
func nox_client_tradeXxxSellAccept_4796D0(a1 int16) int32 {
	var v2 [4]byte
	v2[0] = 201
	v2[1] = 24
	*(*uint16)(unsafe.Pointer(&v2[2])) = uint16(a1)
	return nox_netlist_addToMsgListCli_40EBC0(31, 0, (*uint8)(unsafe.Pointer(&v2[0])), 4)
}
func sub_479700(a1 int16, a2 int8) int32 {
	var v3 [5]byte
	v3[0] = 201
	v3[1] = 25
	*(*uint16)(unsafe.Pointer(&v3[2])) = uint16(a1)
	v3[4] = byte(a2)
	return nox_netlist_addToMsgListCli_40EBC0(31, 0, (*uint8)(unsafe.Pointer(&v3[0])), 5)
}
func sub_479810() {
	dword_5d4594_1098620 = 0
}
func sub_479820(a1 int32, a2 int16) int32 {
	var result int32
	result = sub_479840(a2)
	dword_5d4594_1098620 = 0
	return result
}
func sub_479840(a1 int16) int32 {
	var v2 [4]byte
	v2[0] = 201
	v2[1] = 26
	*(*uint16)(unsafe.Pointer(&v2[2])) = uint16(a1)
	return nox_netlist_addToMsgListCli_40EBC0(31, 0, (*uint8)(unsafe.Pointer(&v2[0])), 4)
}
func sub_479870() int32 {
	return bool2int(dword_5d4594_1098628 == 2)
}
func sub_479880(a1 *uint32) bool {
	return sub_4281F0((*int2)(unsafe.Pointer(a1)), (*int4)(memmap.PtrOff(6112660, 1098380))) != 0
}
func sub_4798A0(a1 *uint32) int32 {
	var (
		v1 int32
		v2 int32
		v3 int32
		v4 int32
		v5 *uint8
	)
	v1 = sub_4281F0((*int2)(unsafe.Pointer(a1)), (*int4)(memmap.PtrOff(6112660, 1098380)))
	if v1 == 0 {
		return 0
	}
	v2 = int32((*a1 - *memmap.PtrUint32(6112660, 1098380)) / 50)
	v3 = int32((*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*1)) - *memmap.PtrUint32(6112660, 1098384) + dword_5d4594_1107036) / 50)
	if v2 >= 6 {
		v2 = 5
	}
	if v3 >= 10 {
		v3 = 9
	}
	v4 = (v3 + v2*10) * 35
	v5 = (*uint8)(memmap.PtrOff(6112660, uint32(v4*4)+0x10C38C))
	if *memmap.PtrUint32(6112660, uint32(v4*4)+0x10C390) == 0 {
		return 0
	}
	*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(v5)) + 128))) = *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v5))), unsafe.Sizeof(uint32(0))*2)))
	return int32(*(*uint32)(unsafe.Pointer(v5)))
}
func sub_479950() int32 {
	var v2 unsafe.Pointer
	if (*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1123524))))).Flags().IsHidden() == 1 {
		return 0
	}
	*memmap.PtrUint8(6112660, 1123516) = 0
	*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v2))), unsafe.Sizeof(uint16(0))*0)) = 720
	*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v2))), 2)) = 0
	nox_netlist_addToMsgListCli_40EBC0(31, 0, (*uint8)(unsafe.Pointer(&v2)), 3)
	return 1
}
func sub_4799A0() int32 {
	var (
		result int32
		v1     *uint32
		v2     *uint32
		v3     *uint32
		v4     int32
		v5     *uint32
		v6     *byte
		v7     *uint32
		v8     *byte
		v9     *byte
		v10    *uint32
	)
	*memmap.PtrUint32(6112660, 1107052) = nox_color_rgb_4344A0(240, 128, 64)
	result = int32(uintptr(unsafe.Pointer(newWindowFromFile("Dialog.wnd", unsafe.Pointer(cgoFuncAddr(nox_xxx_guiDialog_479B00))))))
	dword_5d4594_1123524 = uint32(result)
	if result != 0 {
		result.setFunc93(func(arg1 int32, arg2 int32, arg3 int32, arg4 int32) int32 {
			return sub_479BE0((*uint32)(unsafe.Pointer(uintptr(arg1))), arg2, uint32(arg3), arg4)
		})
		(*(*int32)(unsafe.Pointer(&dword_5d4594_1123524))).setDraw(sub_479CB0)
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1123524))))).setTooltipFunc(unsafe.Pointer(cgoFuncAddr(sub_479D00)))
		v1 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1123524)))).ChildByID(3904)))
		v2 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1123524)))).ChildByID(3903)))
		v10 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1123524)))).ChildByID(3902)))
		v3 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1123524)))).ChildByID(3901)))
		v4 = int32(uintptr(unsafe.Pointer(v3)))
		v5 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*8)))))
		v9 = (*byte)(unsafe.Pointer(nox_xxx_gLoadImg("UISliderLit")))
		v8 = (*byte)(unsafe.Pointer(nox_xxx_gLoadImg("UISliderLit")))
		v6 = (*byte)(unsafe.Pointer(nox_xxx_gLoadImg("UISlider")))
		sub_4B5700(int32(uintptr(unsafe.Pointer(v1))), 0, 0, int32(uintptr(unsafe.Pointer(v6))), int32(uintptr(unsafe.Pointer(v8))), int32(uintptr(unsafe.Pointer(v9))))
		nox_xxx_wnd_46B280(int32(uintptr(unsafe.Pointer(v1))), v4)
		nox_xxx_wnd_46B280(int32(uintptr(unsafe.Pointer(v2))), v4)
		nox_xxx_wnd_46B280(int32(uintptr(unsafe.Pointer(v10))), v4)
		*(*uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(uint32(0))*9)) = uint32(uintptr(unsafe.Pointer(v1)))
		*(*uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(uint32(0))*7)) = uint32(uintptr(unsafe.Pointer(v2)))
		*(*uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(uint32(0))*8)) = uint32(uintptr(unsafe.Pointer(v10)))
		v7 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1123524)))).ChildByID(3906)))
		int32(uintptr(unsafe.Pointer(v7))).setDraw(func(arg1 int32, arg2 int32) int32 {
			return sub_479C40((*uint32)(unsafe.Pointer(uintptr(arg1))), arg2)
		})
		nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1123524))))), 0)
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1123524))))).Hide()
		dword_5d4594_1123520 = 0
		result = 1
	}
	return result
}
func nox_xxx_guiDialog_479B00(a1 int32, a2 int32, a3 *int32, a4 int32) int32 {
	var (
		v3     int32
		result int32
	)
	if a2 != 0x4007 {
		return 0
	}
	v3 = (*nox_window)(unsafe.Pointer(a3)).ID()
	if sub_45D9B0() != 0 {
		return 0
	}
	clientPlaySoundSpecial(766, 100)
	switch v3 {
	case 3906:
		sub_479950()
		result = 0
	case 3907:
		nox_xxx_playDialogFile_44D900(*memmap.PtrInt32(6112660, 1115312), 100)
		result = 0
	case 3908:
		*memmap.PtrUint8(6112660, 1123516) = 1
		*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&a2))), unsafe.Sizeof(uint16(0))*0)) = 720
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a2))), 2)) = 1
		nox_netlist_addToMsgListCli_40EBC0(31, 0, (*uint8)(unsafe.Pointer(&a2)), 3)
		result = 0
	case 3909:
		*memmap.PtrUint8(6112660, 1123516) = 2
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a2))), 2)) = 2
		*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&a2))), unsafe.Sizeof(uint16(0))*0)) = 720
		nox_netlist_addToMsgListCli_40EBC0(31, 0, (*uint8)(unsafe.Pointer(&a2)), 3)
		return 0
	default:
		return 0
	}
	return result
}
func sub_479BE0(a1 *uint32, a2 int32, a3 uint32, a4 int32) int32 {
	switch a2 {
	case 5:
		fallthrough
	case 6:
		fallthrough
	case 7:
		fallthrough
	case 9:
		fallthrough
	case 10:
		fallthrough
	case 11:
		fallthrough
	case 13:
		fallthrough
	case 14:
		fallthrough
	case 15:
		nox_xxx_wndPointInWnd_46AAB0(a1, int32(uint16(a3)), int32(a3>>16))
	default:
		return 1
	}
	return 1
}
func sub_479C40(a1 *uint32, a2 int32) int32 {
	var (
		v2    int8
		yTop  int32
		xLeft int32
	)
	v2 = int8(nox_xxx_bookGet_430B40_get_mouse_prev_seq())
	if sub_44D930() == 0 && (int32(v2)&math.MaxInt8) < 30 && int32(v2)&8 != 0 {
		nox_client_wndGetPosition_46AA60((*nox_window)(unsafe.Pointer(a1)), (*uint32)(unsafe.Pointer(&xLeft)), (*uint32)(unsafe.Pointer(&yTop)))
		sub_49CD30(xLeft, yTop, int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*2))), int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*3))-2), *memmap.PtrInt32(6112660, 1107052), 4)
	}
	return nox_xxx_wndButtonDrawNoImg_4A81D0(int32(uintptr(unsafe.Pointer(a1))), a2)
}
func sub_479CB0(a1 int32, a2 int32) int32 {
	var (
		v2 int32
		v4 int32
		v5 int32
	)
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 24))))
	nox_client_wndGetPosition_46AA60((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1123524)))), (*uint32)(unsafe.Pointer(&v4)), (*uint32)(unsafe.Pointer(&v5)))
	nox_client_drawImageAt_47D2C0((*nox_video_bag_image_t)(unsafe.Pointer(uintptr(v2))), nox_win_width-NOX_DEFAULT_WIDTH, nox_win_height-NOX_DEFAULT_HEIGHT)
	return 1
}
func sub_479D00() int32 {
	return 1
}
func sub_479D10() int32 {
	var result int32
	(*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1123524)))).Destroy()
	result = 0
	dword_5d4594_1123524 = 0
	dword_5d4594_1123520 = 0
	return result
}
func sub_47A1F0() int32 {
	var (
		v0     *uint32
		result int32
	)
	v0 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1123524)))).ChildByID(3901)))
	(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v0)))))).Func94(asWindowEvent(0x400F, 0, 0))
	nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1123524))))), 0)
	(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1123524))))).Hide()
	sub_44D8F0()
	result = int32(nox_client_renderGUI_80828)
	dword_5d4594_1123520 = 0
	if nox_client_renderGUI_80828 == 0 && *memmap.PtrUint32(0x587000, 153436) == 1 {
		nox_client_renderGUI_80828 = 1
	}
	return result
}
func sub_47A260() int32 {
	return int32(dword_5d4594_1123520)
}
func sub_47A310(lpPathName *byte) uint32 {
	var (
		result       uint32
		v3           *byte
		v4           *byte
		v5           HANDLE
		v6           *byte
		v7           *byte
		v8           int32
		v9           HANDLE
		FindFileData _WIN32_FIND_DATAA
		Buffer       [512]byte
	)
	dword_5d4594_1189600 = 0
	v8 = 0
	if nox_fs_set_workdir(lpPathName) || (func() uint32 {
		result = compatGetLastError()
		return result
	}()) == 0 {
		nox_fs_workdir(&Buffer[0], 512)
		v3 = (*byte)(memmap.PtrOff(6112660, func() uint32 {
			p := &dword_5d4594_1189600
			x := *p
			*p++
			return x
		}()*512+0x1126D0))
		libc.StrCpy(v3, &Buffer[0])
		if *(*int32)(unsafe.Pointer(&dword_5d4594_1189600)) <= 0 {
		LABEL_27:
			nox_fs_set_workdir(&Buffer[0])
			result = uint32(bool2int(nox_fs_set_workdir(CString(".."))))
		} else {
			v4 = (*byte)(memmap.PtrOff(6112660, 1124048))
			for nox_fs_set_workdir(v4) || compatGetLastError() == 0 {
				v5 = compatFindFirstFileA((*byte)(memmap.PtrOff(0x587000, 153724)), &FindFileData)
				v9 = v5
				if uintptr(v5) != uintptr(math.MaxUint32) {
					if FindFileData.dwFileAttributes&16 != 0 && libc.StrCmp(&FindFileData.cFileName[0], CString(".")) != 0 && libc.StrCmp(&FindFileData.cFileName[0], CString("..")) != 0 {
						v6 = (*byte)(memmap.PtrOff(6112660, dword_5d4594_1189600*512+0x1126D0))
						libc.StrCpy(v6, v4)
						if *(*byte)(unsafe.Add(unsafe.Pointer(v4), libc.StrLen(v4)-1)) != 92 {
							*(*uint16)(unsafe.Add(unsafe.Pointer(v6), libc.StrLen(v6))) = *memmap.PtrUint16(0x587000, 153736)
						}
						libc.StrCat(v6, &FindFileData.cFileName[0])
						if libc.StrLen(v6) >= 512 {
							nox_exit(1)
						}
						if func() int32 {
							p := (*int32)(unsafe.Pointer(&dword_5d4594_1189600))
							*p++
							return *p
						}() >= 128 {
							nox_exit(1)
						}
					}
					for compatFindNextFileA(v5, &FindFileData) != 0 {
						if FindFileData.dwFileAttributes&16 != 0 && libc.StrCmp(&FindFileData.cFileName[0], CString(".")) != 0 && libc.StrCmp(&FindFileData.cFileName[0], CString("..")) != 0 {
							v7 = (*byte)(memmap.PtrOff(6112660, dword_5d4594_1189600*512+0x1126D0))
							libc.StrCpy(v7, v4)
							if *(*byte)(unsafe.Add(unsafe.Pointer(v4), libc.StrLen(v4)-1)) != 92 {
								*(*uint16)(unsafe.Add(unsafe.Pointer(v7), libc.StrLen(v7))) = *memmap.PtrUint16(0x587000, 153748)
							}
							libc.StrCat(v7, &FindFileData.cFileName[0])
							if libc.StrLen(v7) >= 512 {
								nox_exit(1)
							}
							if func() int32 {
								p := (*int32)(unsafe.Pointer(&dword_5d4594_1189600))
								*p++
								return *p
							}() >= 128 {
								nox_exit(1)
							}
							v5 = v9
						}
					}
					compatFindClose(v5)
				}
				v4 = (*byte)(unsafe.Add(unsafe.Pointer(v4), 512))
				if func() int32 {
					p := &v8
					*p++
					return *p
				}() >= *(*int32)(unsafe.Pointer(&dword_5d4594_1189600)) {
					goto LABEL_27
				}
			}
			nox_fs_set_workdir(&Buffer[0])
			result = uint32(bool2int(nox_fs_set_workdir(CString(".."))))
		}
	}
	return result
}
func nox_xxx_videoBag_LoadTile_47AD60(a1 int32, a2 int32, a3 *uint16) int32 {
	var (
		v3  int32
		v4  int32
		i   int32
		v6  int32
		v7  int32
		v8  int16
		v9  int32
		v10 int32
		v11 int32
		v12 int32
		v13 int32
		v14 int16
		v15 int16
		v17 int32
	)
	_ = v17
	var v18 int32
	v17 = 0
	if a1 != 46 || a2 != 46 {
		return 0
	}
	v3 = 0xC8000
	v4 = 1
	for i = 23; i > 0; i-- {
		v6 = 0
		if v4 > 0 {
			v17 += v4
			for {
				v7 = i + v3
				v8 = int16((int32(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, 1189588) + uint32(i) + uint32(v3) + uint32(v6) - 0x64000))))&248 | (int32(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, 1189588) + uint32(v7) + uint32(v6) - 0xC8000))))&248)*32) * 4)
				*a3 = uint16(int16((int32(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, 1189588) + uint32(i) + uint32(v3) + uint32(func() int32 {
					p := &v6
					x := *p
					*p++
					return x
				}()))))) >> 3) | int32(v8)))
				a3 = (*uint16)(unsafe.Add(unsafe.Pointer(a3), unsafe.Sizeof(uint16(0))*1))
				if v6 >= v4 {
					break
				}
			}
		}
		v4 += 2
		v3 += 640
	}
	v9 = 1
	v10 = 0xCB980
	v11 = 45
	v18 = 23
	for {
		v12 = 0
		if v11 > 0 {
			v17 += v11
			for {
				v13 = v10 + v9
				v14 = int16((int32(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, 1189588) + uint32(v10) + uint32(v9) + uint32(v12) - 0x64000))))&248 | (int32(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, 1189588) + uint32(v13) + uint32(v12) - 0xC8000))))&248)*32) * 4)
				a3 = (*uint16)(unsafe.Add(unsafe.Pointer(a3), unsafe.Sizeof(uint16(0))*1))
				v15 = int16(int32(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, 1189588) + uint32(v10) + uint32(v9) + uint32(func() int32 {
					p := &v12
					x := *p
					*p++
					return x
				}()))))) >> 3)
				*((*uint16)(unsafe.Add(unsafe.Pointer(a3), -int(unsafe.Sizeof(uint16(0))*1)))) = uint16(int16(int32(v15) | int32(v14)))
				if v12 >= v11 {
					break
				}
			}
		}
		v9++
		v11 -= 2
		v10 += 640
		v18--
		if v18 == 0 {
			break
		}
	}
	return 1
}
func nox_xxx_videoBag_LoadEdge_47AF30(a1 int32, a2 int32, a3 *uint8) int32 {
	var (
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
		i   int32
		v17 int32
		v18 int32
		v19 int32
		v20 int32
		v21 int32
		v22 int32
		v23 int32
		v24 int32
		v25 int32
		v26 int32
		j   int32
		v28 int32
	)
	_ = v28
	var v29 int32
	var v30 *uint8
	var v31 int32
	var v32 int32
	var v33 int32
	var v34 int32
	var v35 int32
	var v36 int32
	var v37 int32
	var v38 int32
	var v39 int32
	var v40 int16
	var v41 int32
	var v42 *uint16
	var v43 int32
	var v44 int32
	var v45 int8
	var v46 int32
	var v47 int32
	var v48 int16
	var v49 *uint8
	_ = v49
	var v50 int32
	var v51 *uint8
	_ = v51
	var v52 int32
	var v53 *uint8
	_ = v53
	var v54 *uint16
	var v55 bool
	var v56 int32
	var v57 int32
	var v58 int32
	var v59 int32
	var v60 int16
	var v61 int32
	var v62 *uint16
	var v63 int32
	var v64 int32
	var v65 int32
	var v66 int16
	var v67 *uint8
	_ = v67
	var v68 int32
	var v69 *uint8
	_ = v69
	var v70 int32
	var v71 *uint8
	_ = v71
	var v72 *uint16
	var v73 bool
	var v75 int32
	var v76 int32
	var v77 int32
	var v78 uint8
	var v79 int32
	var v80 uint8
	var v81 int16
	var v82 int32
	var v83 int32
	var v84 int16
	var v85 int32
	var v86 int8
	var v87 *uint8
	var v88 *uint8
	var v89 int8
	var v90 int8
	var v91 int32
	var v92 int32
	var v93 int32
	var v94 int32
	var v95 int32
	var v96 int32
	var v97 int32
	*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v4))), unsafe.Sizeof(uint16(0))*1)) = 0
	if a1 != 46 || a2 != 46 {
		return 0
	}
	*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint16(0))*1)) = *memmap.PtrUint16(6112660, 1189590)
	v86 = int8(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, 1189588) + 1))))
	v89 = int8(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, 1189588) + 0x64001))))
	v90 = int8(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, 1189588) + 0xC8001))))
	*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v4))), unsafe.Sizeof(uint16(0))*0)) = uint16(int16(int32(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, 1189588) + 1)))) & 248))
	*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v3))), unsafe.Sizeof(uint16(0))*0)) = uint16(int16(int32(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, 1189588) + 0x64001)))) & 248))
	v6 = (v3 | v4*32) * 4
	*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint16(0))*0)) = uint16(int16(int32(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, 1189588) + 0xC8001)))) >> 3))
	v7 = v5 | v6
	v81 = int16(v7)
	v8 = int32(*memmap.PtrUint32(6112660, 1189588) + 0xC8000)
	*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v7))), unsafe.Sizeof(uint16(0))*0)) = uint16(int16(int32(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, 1189588) + 0x64000)))) & 248))
	v78 = math.MaxUint8
	*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v3))), unsafe.Sizeof(uint16(0))*0)) = uint16(int16(int32(**(**uint8)(memmap.PtrOff(6112660, 1189588))) & 248))
	v9 = v3 * 32
	v10 = v9 | v7
	*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v9))), unsafe.Sizeof(uint16(0))*0)) = uint16(int16(int32(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, 1189588) + 0xC8000)))) >> 3))
	v11 = 1
	v12 = v9 | v10*4
	v13 = 0
	v84 = int16(v12)
	v91 = 0
	v14 = 0
	v15 = 23
	for v15 > 0 {
		for i = 0; i < v11; i++ {
			if int32(v86) != int32(*(*uint8)(unsafe.Pointer(uintptr(uint32(v8+v15+i) - 0xC8000)))) || int32(v89) != int32(*(*uint8)(unsafe.Pointer(uintptr(uint32(v8+v15+i) - 0x64000)))) || int32(v90) != int32(*(*uint8)(unsafe.Pointer(uintptr(v8 + v15 + i)))) {
				v13 = 1
				v78 = uint8(int8(v14))
				v91 = 1
				break
			}
			v13 = v91
		}
		v14++
		v15--
		v11 += 2
		v8 += 640
		if v13 != 0 {
			goto LABEL_27
		}
	}
	v17 = int32(*memmap.PtrUint32(6112660, 1189588) + uint32((v14*5+6400)<<7))
	v18 = 91 - v14*2
	for {
		if v14 >= 46 {
			break
		}
		v19 = 0
		if v18 > 0 {
			for int32(v86) == int32(*(*uint8)(unsafe.Pointer(uintptr(uint32(v19+v17+v14) - 0xC8016)))) && int32(v89) == int32(*(*uint8)(unsafe.Pointer(uintptr(uint32(v14+v17+v19) - 0x64016)))) && int32(v90) == int32(*(*uint8)(unsafe.Pointer(uintptr(v14 + v17 + v19 - 22)))) {
				if func() int32 {
					p := &v19
					*p++
					return *p
				}() >= v18 {
					goto LABEL_26
				}
			}
			v78 = uint8(int8(v14))
			v13 = 1
		}
	LABEL_26:
		v14++
		v18 -= 2
		v17 += 640
		if v13 != 0 {
			break
		}
	}
LABEL_27:
	v20 = 0
	v80 = math.MaxUint8
	v92 = 0
	v21 = 45
	v22 = int32(*memmap.PtrUint32(6112660, 1189588) + 848000)
	v23 = 1
	for v23 <= 45 {
		v24 = 0
		if v23 > 0 {
			for int32(v86) == int32(*(*uint8)(unsafe.Pointer(uintptr(uint32(v24+v22+v21) - 0xC8016)))) && int32(v89) == int32(*(*uint8)(unsafe.Pointer(uintptr(uint32(v22+v24+v21) - 0x64016)))) && int32(v90) == int32(*(*uint8)(unsafe.Pointer(uintptr(v22 + v24 + v21 - 22)))) {
				if func() int32 {
					p := &v24
					*p++
					return *p
				}() >= v23 {
					v20 = v92
					goto LABEL_36
				}
			}
			v20 = 1
			v80 = uint8(int8(v21))
			v92 = 1
		}
	LABEL_36:
		v21--
		v23 += 2
		v22 -= 640
		if v20 != 0 {
			goto LABEL_48
		}
	}
	v93 = v21*2 + 1
	v25 = 23 - v21
	v26 = int32(*memmap.PtrUint32(6112660, 1189588) + uint32((v21*5+6400)<<7))
	for {
		if v21 < 0 {
			break
		}
		for j = 0; j < v93; j++ {
			if int32(v86) != int32(*(*uint8)(unsafe.Pointer(uintptr(uint32(v26+v25+j) - 0xC8000)))) || int32(v89) != int32(*(*uint8)(unsafe.Pointer(uintptr(uint32(v26+v25+j) - 0x64000)))) || int32(v90) != int32(*(*uint8)(unsafe.Pointer(uintptr(v26 + v25 + j)))) {
				v80 = uint8(int8(v21))
				v20 = 1
				break
			}
		}
		v21--
		v25++
		v26 -= 640
		v93 -= 2
		if v20 != 0 {
			break
		}
	}
LABEL_48:
	v28 = 2
	*a3 = v78
	v29 = 2
	*(*uint8)(unsafe.Add(unsafe.Pointer(a3), 1)) = v80
	v30 = (*uint8)(unsafe.Add(unsafe.Pointer(a3), 2))
	if int32(v80) >= 23 {
		v31 = 22
	} else {
		v31 = int32(v80)
	}
	v32 = int32(v78)
	v77 = int32(v78)
	if int32(v78) <= v31 {
		v82 = int32(v78)*2 + 1
		v33 = 23 - int32(v78)
		v34 = (int32(v78)*5 + 6400) << 7
		v35 = v31 - int32(v78) + 1
		v75 = (int32(v78)*5 + 6400) << 7
		v36 = v35 + int32(v78)
		v79 = 23 - int32(v78)
		v85 = v35
		v77 = v36
		for {
			v37 = int32(uint32(v34+v33) - 0xC8000)
			v38 = int32(uint32(v34+v33) - 0x64000)
			*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v37))), unsafe.Sizeof(uint16(0))*0)) = uint16(int16(int32(*(*uint8)(unsafe.Pointer(uintptr(uint32(v37) + *memmap.PtrUint32(6112660, 1189588))))) & 248))
			*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v38))), unsafe.Sizeof(uint16(0))*0)) = uint16(int16(int32(*(*uint8)(unsafe.Pointer(uintptr(uint32(v38) + *memmap.PtrUint32(6112660, 1189588))))) & 248))
			v39 = (v38 | v37*32) * 4
			v40 = int16((int32(*(*uint8)(unsafe.Pointer(uintptr(uint32(v33+v75) + *memmap.PtrUint32(6112660, 1189588))))) >> 3) | v39)
			if int32(v40) == int32(v81) {
				v41 = 1
				v28 += 2
				*v30 = 1
				v87 = (*uint8)(unsafe.Add(unsafe.Pointer(v30), 1))
				v30 = (*uint8)(unsafe.Add(unsafe.Pointer(v30), 2))
				v94 = 1
				v29 += 2
			} else if int32(v40) == int32(v84) {
				*v30 = 3
				v87 = (*uint8)(unsafe.Add(unsafe.Pointer(v30), 1))
				v30 = (*uint8)(unsafe.Add(unsafe.Pointer(v30), 2))
				v28 += 2
				v94 = 3
				v41 = 1
				v29 += 2
			} else {
				*v30 = 2
				v42 = (*uint16)(unsafe.Add(unsafe.Pointer(v30), 1))
				v87 = (*uint8)(unsafe.Pointer(v42))
				v42 = (*uint16)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(v42))), 1))))
				v28 += 4
				v94 = 2
				*v42 = uint16(v40)
				v30 = (*uint8)(unsafe.Pointer((*uint16)(unsafe.Add(unsafe.Pointer(v42), unsafe.Sizeof(uint16(0))*1))))
				v41 = 1
				v29 += 3
			}
			v43 = 1
			v96 = 1
			if v82 > 1 {
				v44 = v75
				for {
					v45 = int8(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, 1189588) + uint32(v44) + uint32(v33) - 0xC8000 + uint32(v43)))))
					v46 = int32(*memmap.PtrUint32(6112660, 1189588) + uint32(v75) + uint32(v33) - 0x64000)
					*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v43))), unsafe.Sizeof(uint16(0))*0)) = uint16(int16(int32(v45) & 248))
					*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v46))), unsafe.Sizeof(uint16(0))*0)) = uint16(int16(int32(*(*uint8)(unsafe.Pointer(uintptr(v46 + v96)))) & 248))
					v47 = (v46 | v43*32) * 4
					v48 = int16((int32(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, 1189588) + uint32(v75) + uint32(v79) + uint32(v96))))) >> 3) | v47)
					if int32(v48) == int32(v81) {
						if v94 == 1 {
							v41++
						} else {
							v49 = v87
							v50 = int32(uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v30), 1)))))
							v87 = (*uint8)(unsafe.Pointer(uintptr(v50)))
							v30 = (*uint8)(unsafe.Pointer(uintptr(v50 + 1)))
							*v49 = uint8(int8(v41))
							*((*uint8)(unsafe.Add(unsafe.Pointer(v30), -2))) = 1
							v28 += 2
							v94 = 1
							v41 = 1
							v29 += 2
						}
					} else if int32(v48) == int32(v84) {
						if v94 == 3 {
							v41++
						} else {
							v51 = v87
							v52 = int32(uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v30), 1)))))
							v87 = (*uint8)(unsafe.Pointer(uintptr(v52)))
							v30 = (*uint8)(unsafe.Pointer(uintptr(v52 + 1)))
							*v51 = uint8(int8(v41))
							*((*uint8)(unsafe.Add(unsafe.Pointer(v30), -2))) = 3
							v28 += 2
							v94 = 3
							v41 = 1
							v29 += 2
						}
					} else if v94 == 2 {
						*(*uint16)(unsafe.Pointer(v30)) = uint16(v48)
						v30 = (*uint8)(unsafe.Add(unsafe.Pointer(v30), 2))
						v41++
						v28 += 2
						v29++
					} else {
						v53 = v87
						v54 = (*uint16)(unsafe.Add(unsafe.Pointer(v30), 1))
						v87 = (*uint8)(unsafe.Pointer(v54))
						v54 = (*uint16)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(v54))), 1))))
						*v53 = uint8(int8(v41))
						*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(v54))), -2))) = 2
						*v54 = uint16(v48)
						v30 = (*uint8)(unsafe.Pointer((*uint16)(unsafe.Add(unsafe.Pointer(v54), unsafe.Sizeof(uint16(0))*1))))
						v28 += 4
						v94 = 2
						v41 = 1
						v29 += 3
					}
					if func() int32 {
						p := &v96
						*p++
						return *p
					}() >= v82 {
						break
					}
					v44 = v75
					v33 = v79
					v43 = v96
				}
				v33 = v79
			}
			v32 = v77
			v33--
			*v87 = uint8(int8(v41))
			v82 += 2
			v55 = v85 == 1
			v79 = v33
			v75 += 640
			v85--
			if v55 {
				break
			}
			v34 = v75
		}
	}
	if v32 <= int32(v80) {
		v56 = (v32*5 + 6400) << 7
		v76 = (v32*5 + 6400) << 7
		v83 = 91 - v32*2
		for {
			v57 = int32(*memmap.PtrUint32(6112660, 1189588) + uint32(v56) - 0xC8000)
			v58 = int32(*memmap.PtrUint32(6112660, 1189588) + uint32(v56) - 0x64000)
			*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v57))), unsafe.Sizeof(uint16(0))*0)) = uint16(int16(int32(*(*uint8)(unsafe.Pointer(uintptr(v57 + v32 - 22)))) & 248))
			*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v58))), unsafe.Sizeof(uint16(0))*0)) = uint16(int16(int32(*(*uint8)(unsafe.Pointer(uintptr(v58 + v32 - 22)))) & 248))
			v59 = (v58 | v57*32) * 4
			v60 = int16((int32(*(*uint8)(unsafe.Pointer(uintptr(uint32(v56) + *memmap.PtrUint32(6112660, 1189588) + uint32(v32) - 22)))) >> 3) | v59)
			if int32(v60) == int32(v81) {
				*v30 = 1
				v61 = 1
				v88 = (*uint8)(unsafe.Add(unsafe.Pointer(v30), 1))
				v30 = (*uint8)(unsafe.Add(unsafe.Pointer(v30), 2))
				v28 += 2
				v95 = 1
				v29 += 2
			} else if int32(v60) == int32(v84) {
				*v30 = 3
				v88 = (*uint8)(unsafe.Add(unsafe.Pointer(v30), 1))
				v30 = (*uint8)(unsafe.Add(unsafe.Pointer(v30), 2))
				v28 += 2
				v95 = 3
				v61 = 1
				v29 += 2
			} else {
				*v30 = 2
				v62 = (*uint16)(unsafe.Add(unsafe.Pointer(v30), 1))
				v88 = (*uint8)(unsafe.Pointer(v62))
				v62 = (*uint16)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(v62))), 1))))
				v28 += 4
				v95 = 2
				*v62 = uint16(v60)
				v30 = (*uint8)(unsafe.Pointer((*uint16)(unsafe.Add(unsafe.Pointer(v62), unsafe.Sizeof(uint16(0))*1))))
				v61 = 1
				v29 += 3
			}
			v97 = 1
			if v83 > 1 {
				for {
					v63 = int32(uint32(v97) + *memmap.PtrUint32(6112660, 1189588) + uint32(v56) - 0xC8000)
					*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v63))), 0)) = *(*uint8)(unsafe.Pointer(uintptr(v63 + v32 - 22)))
					v64 = int32(uint32(v97) + *memmap.PtrUint32(6112660, 1189588) + uint32(v76) - 0x64000)
					*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v63))), unsafe.Sizeof(uint16(0))*0)) = uint16(int16(v63 & 248))
					*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v64))), unsafe.Sizeof(uint16(0))*0)) = uint16(int16(int32(*(*uint8)(unsafe.Pointer(uintptr(v64 + v77 - 22)))) & 248))
					v65 = (v64 | v63*32) * 4
					v66 = int16((int32(*(*uint8)(unsafe.Pointer(uintptr(uint32(v97) + *memmap.PtrUint32(6112660, 1189588) + uint32(v76) + uint32(v77) - 22)))) >> 3) | v65)
					if int32(v66) == int32(v81) {
						if v95 == 1 {
							v61++
						} else {
							v67 = v88
							v68 = int32(uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v30), 1)))))
							v88 = (*uint8)(unsafe.Pointer(uintptr(v68)))
							v30 = (*uint8)(unsafe.Pointer(uintptr(v68 + 1)))
							*v67 = uint8(int8(v61))
							*((*uint8)(unsafe.Add(unsafe.Pointer(v30), -2))) = 1
							v28 += 2
							v95 = 1
							v61 = 1
							v29 += 2
						}
					} else if int32(v66) == int32(v84) {
						if v95 == 3 {
							v61++
						} else {
							v69 = v88
							v70 = int32(uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v30), 1)))))
							v88 = (*uint8)(unsafe.Pointer(uintptr(v70)))
							v30 = (*uint8)(unsafe.Pointer(uintptr(v70 + 1)))
							*v69 = uint8(int8(v61))
							*((*uint8)(unsafe.Add(unsafe.Pointer(v30), -2))) = 3
							v28 += 2
							v95 = 3
							v61 = 1
							v29 += 2
						}
					} else if v95 == 2 {
						*(*uint16)(unsafe.Pointer(v30)) = uint16(v66)
						v30 = (*uint8)(unsafe.Add(unsafe.Pointer(v30), 2))
						v61++
						v28 += 2
						v29++
					} else {
						v71 = v88
						v72 = (*uint16)(unsafe.Add(unsafe.Pointer(v30), 1))
						v88 = (*uint8)(unsafe.Pointer(v72))
						v72 = (*uint16)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(v72))), 1))))
						*v71 = uint8(int8(v61))
						*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(v72))), -2))) = 2
						*v72 = uint16(v66)
						v30 = (*uint8)(unsafe.Pointer((*uint16)(unsafe.Add(unsafe.Pointer(v72), unsafe.Sizeof(uint16(0))*1))))
						v28 += 4
						v95 = 2
						v61 = 1
						v29 += 3
					}
					v56 = v76
					v73 = v97+1 < v83
					v32 = v77
					v97++
					if !v73 {
						break
					}
				}
			}
			v32++
			v56 += 640
			v77 = v32
			*v88 = uint8(int8(v61))
			v76 = v56
			v83 -= 2
			if v32 > int32(v80) {
				break
			}
		}
	}
	return 1
}
func nox_xxx_videoBag_LoadPXImg_47B7F0(a1 int32, a2 int32, a3 int32, a4 uint8) int32 {
	var (
		v4  int32
		v5  int32
		v6  uint8
		i   int32
		v8  int32
		v9  int32
		v10 int32
		v11 int32
		v12 int32
		v13 int32
		j   int32
		v15 int32
		v16 int32
		v17 int32
		v18 int32
		v19 int32
		v20 int32
		v21 *uint8
		v22 int32
		v23 int32
		v24 int32
		v25 int32
		v26 int32
		v27 int32
		v28 *uint8
		v29 int32
		v30 int32
		v31 int32
		v32 int32
		v34 *uint8
		v35 int32
		v36 int32
		v37 int32
		v38 int32
		v39 *uint8
		v40 int16
		v41 int16
		v42 *uint16
		v43 int32
		v44 int32
		v45 int32
		v46 bool
		v47 int32
		v48 int16
		v49 int16
		v50 *uint16
		v51 *uint8
		v52 int8
		v53 *uint8
		v54 int32
		v55 int16
		v56 int16
		v57 int16
		v58 *uint16
		v59 int32
		v60 int32
		v61 int8
		v62 *uint8
		v63 int16
		v64 int16
		v65 *uint16
		v66 int32
		v67 int8
		v68 *uint16
		v69 int32
		v70 int32
		v71 int32
		v72 *uint8
	)
	_ = v72
	var v73 int32
	var v74 int32
	var v75 int32
	var v76 int32
	var v77 int32
	var v78 int32
	var v79 *uint16
	var v80 *uint8
	_ = v80
	var v81 int32
	var v82 int32
	var v83 int8
	var v84 *uint8
	var v85 int32
	var v86 int8
	var v87 int8
	var v88 *uint16
	var v89 int32
	var v90 int8
	var v91 *uint8
	var v92 int32
	var v93 int32
	var v94 int32
	var v95 int32
	var v96 int32
	var v97 int32
	var v98 *uint16
	var v99 *uint8
	_ = v99
	var v100 int32
	var v101 int32
	var v102 int32
	var v103 int32
	var v104 int32
	var v105 int32
	var v106 int32
	var v107 int32
	var v108 int32
	var v109 int32
	var v110 uint8
	var v111 uint8
	var v112 uint8
	var v113 uint8
	var v114 int32
	var v115 uint8
	var v116 uint8
	var v117 uint8
	var v118 uint8
	var v119 int32
	var v120 int32
	var v121 uint8
	var v122 uint8
	var v123 int32
	var v124 int32
	var v125 int32
	var v126 int16
	var v127 int16
	var v128 uint8
	var v129 uint8
	var v130 uint8
	var v131 int32
	var v132 int32
	var v133 int32
	var v134 int32
	var v135 int32
	var v136 int32
	var v137 int32
	var v138 int32
	var v139 *uint8
	var v140 *uint8
	var v141 *uint8
	var v142 *uint8
	var v143 int32
	var v144 int32
	var v145 int32
	var v146 int32
	var v147 int32
	var v148 int32
	var v149 int32
	var v150 int32
	v114 = -1
	v4 = int32(*memmap.PtrUint32(6112660, 1189588) + 0xC8000)
	v128 = **(**uint8)(memmap.PtrOff(6112660, 1189588))
	v129 = *(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, 1189588) + 0x64000)))
	v130 = *(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, 1189588) + 0xC8000)))
	v5 = 0
	v6 = **(**uint8)(memmap.PtrOff(6112660, 1189588))
	v124 = 0
	for {
		if v124 >= a2 {
			break
		}
		for i = 0; i < a1; i++ {
			v8 = int32(*(*uint8)(unsafe.Pointer(uintptr(uint32(v4+i) - 0xC8000))))
			if int32(v128) <= v8+5 && int32(v6) >= v8-5 {
				v9 = int32(*(*uint8)(unsafe.Pointer(uintptr(uint32(v4+i) - 0x64000))))
				if int32(v129) <= v9+5 && int32(v129) >= v9-5 {
					v10 = int32(*(*uint8)(unsafe.Pointer(uintptr(v4 + i))))
					if int32(v130) <= v10+5 && int32(v130) >= v10-5 {
						continue
					}
				}
			}
			v5 = 1
			v114 = v124
			break
		}
		v4 += 640
		v124++
		if v5 != 0 {
			break
		}
	}
	v11 = 0
	v123 = -1
	v12 = a2 - 1
	v13 = int32(*memmap.PtrUint32(6112660, 1189588) + uint32(((a2-1)*5+6400)<<7))
	for {
		if v12 < 0 {
			break
		}
		for j = 0; j < a1; j++ {
			v15 = int32(*(*uint8)(unsafe.Pointer(uintptr(uint32(v13+j) - 0xC8000))))
			if int32(v6) > v15+5 || int32(v6) < v15-5 || (func() bool {
				v16 = int32(*(*uint8)(unsafe.Pointer(uintptr(uint32(v13+j) - 0x64000))))
				return int32(v129) > v16+5
			}()) || int32(v129) < v16-5 || (func() bool {
				v17 = int32(*(*uint8)(unsafe.Pointer(uintptr(v13 + j))))
				return int32(v130) > v17+5
			}()) || int32(v130) < v17-5 {
				v6 = **(**uint8)(memmap.PtrOff(6112660, 1189588))
				v123 = v12
				v11 = 1
				break
			}
			v6 = **(**uint8)(memmap.PtrOff(6112660, 1189588))
		}
		v12--
		v13 -= 640
		if v11 != 0 {
			break
		}
	}
	v18 = -1
	v119 = 0
	v19 = 0
	for {
		if v19 >= a1 {
			break
		}
		v20 = v114
		v21 = (*uint8)(unsafe.Pointer(uintptr(uint32(v19) + *memmap.PtrUint32(6112660, 1189588) + uint32((v114*5+6400)<<7))))
		for v20 <= v123 {
			v22 = int32(*((*uint8)(unsafe.Add(unsafe.Pointer(v21), -819200))))
			if int32(v128) > v22+5 || int32(v128) < v22-5 || (func() bool {
				v23 = int32(*((*uint8)(unsafe.Add(unsafe.Pointer(v21), -409600))))
				return int32(v129) > v23+5
			}()) || int32(v129) < v23-5 || (func() bool {
				v24 = int32(*v21)
				return int32(v130) > v24+5
			}()) || int32(v130) < v24-5 {
				v25 = 1
				v18 = v19
				v119 = 1
				goto LABEL_40
			}
			v20++
			v21 = (*uint8)(unsafe.Add(unsafe.Pointer(v21), 640))
		}
		v25 = v119
	LABEL_40:
		v19++
		if v25 != 0 {
			break
		}
	}
	v106 = v18
	v109 = -1
	v120 = 0
	v26 = a1 - 1
	for {
		if v26 < 0 {
			break
		}
		v27 = v114
		v28 = (*uint8)(unsafe.Pointer(uintptr(uint32(v26) + *memmap.PtrUint32(6112660, 1189588) + uint32((v114*5+6400)<<7))))
		for v27 <= v123 {
			v29 = int32(*((*uint8)(unsafe.Add(unsafe.Pointer(v28), -819200))))
			if int32(v128) > v29+5 || int32(v128) < v29-5 || (func() bool {
				v30 = int32(*((*uint8)(unsafe.Add(unsafe.Pointer(v28), -409600))))
				return int32(v129) > v30+5
			}()) || int32(v129) < v30-5 || (func() bool {
				v31 = int32(*v28)
				return int32(v130) > v31+5
			}()) || int32(v130) < v31-5 {
				v32 = 1
				v109 = v26
				v120 = 1
				goto LABEL_54
			}
			v27++
			v28 = (*uint8)(unsafe.Add(unsafe.Pointer(v28), 640))
		}
		v32 = v120
	LABEL_54:
		v26--
		if v32 != 0 {
			break
		}
	}
	if v114 == -1 {
		*(*uint32)(unsafe.Pointer(uintptr(a3))) = 0
		*(*uint32)(unsafe.Pointer(uintptr(a3 + 4))) = 0
		*(*uint32)(unsafe.Pointer(uintptr(a3 + 8))) = 0
		*(*uint32)(unsafe.Pointer(uintptr(a3 + 12))) = 0
		*(*uint8)(unsafe.Pointer(uintptr(a3 + 16))) = 0
		*(*uint8)(unsafe.Pointer(uintptr(a3 + 17))) = 0
		*(*uint8)(unsafe.Pointer(uintptr(a3 + 18))) = 0
		return 1
	}
	*(*uint32)(unsafe.Pointer(uintptr(a3))) = uint32(v109 - v18 + 1)
	v34 = (*uint8)(unsafe.Pointer(uintptr(a3 + 17)))
	*(*uint32)(unsafe.Pointer(uintptr(a3 + 4))) = uint32(v123 - v114 + 1)
	v35 = 17
	*(*uint32)(unsafe.Pointer(uintptr(a3 + 8))) = uint32(v18)
	*(*uint32)(unsafe.Pointer(uintptr(a3 + 12))) = uint32(v114)
	*(*uint8)(unsafe.Pointer(uintptr(a3 + 16))) = uint8(int8(int32(a4) | 128))
	v36 = 17
	if int32(a4) == 4 || int32(a4) == 6 {
		sub_47D090(a1, a2)
	}
	v37 = v114
	if v114 > v123 {
		goto LABEL_208
	}
	v38 = int32(a4) - 3
	v125 = int32(a4) - 3
	for {
		switch v38 {
		case 0:
			v110 = *(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, 1189588) + uint32(v37*640) + uint32(v18))))
			if int32(v128) > int32(v110)+5 || int32(v128) < int32(v110)-5 || (func() bool {
				v107 = (v37*5 + 3200) << 7
				v131 = int32(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, 1189588) + uint32(v107) + uint32(v18)))))
				return int32(v129) > v131+5
			}()) || int32(v129) < v131-5 || (func() bool {
				v104 = (v37*5 + 6400) << 7
				v132 = int32(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, 1189588) + uint32(v104) + uint32(v18)))))
				return int32(v130) > v132+5
			}()) || int32(v130) < v132-5 {
				v104 = (v37*5 + 6400) << 7
				v107 = (v37*5 + 3200) << 7
				v40 = int16((((int32(v110) & 248) * 32) | (int32(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, 1189588) + uint32((v37*5+3200)<<7) + uint32(v18))))) & 248)) * 4)
				v41 = int16((int32(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, 1189588) + uint32((v37*5+6400)<<7) + uint32(v18))))) >> 3) | int32(v40))
				*v34 = 2
				v42 = (*uint16)(unsafe.Add(unsafe.Pointer(v34), 1))
				v35 += 4
				v39 = (*uint8)(unsafe.Pointer(v42))
				v42 = (*uint16)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(v42))), 1))))
				v147 = 2
				v133 = 1
				*v42 = uint16(v41)
				v34 = (*uint8)(unsafe.Pointer((*uint16)(unsafe.Add(unsafe.Pointer(v42), unsafe.Sizeof(uint16(0))*1))))
				v139 = v39
				v36 += 3
			} else {
				v35 += 2
				*v34 = 1
				v147 = 1
				v133 = 1
				v39 = (*uint8)(unsafe.Add(unsafe.Pointer(v34), 1))
				v34 = (*uint8)(unsafe.Add(unsafe.Pointer(v34), 2))
				v139 = v39
				v36 += 2
			}
			v143 = v18 + 1
			if v18+1 > v109 {
				goto LABEL_97
			}
			for {
				v115 = *(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, 1189588) + uint32(v37*640) + uint32(v143))))
				if int32(v128) > int32(v115)+5 || int32(v128) < int32(v115)-5 {
					goto LABEL_88
				}
				v43 = int32(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, 1189588) + uint32(v107) + uint32(v143)))))
				if int32(v129) > v43+5 || int32(v129) < v43-5 || (func() bool {
					v44 = int32(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, 1189588) + uint32(v104) + uint32(v143)))))
					return int32(v130) > v44+5
				}()) || int32(v130) < v44-5 {
					v39 = v139
				LABEL_88:
					v48 = int16((((int32(v115) & 248) * 32) | (int32(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, 1189588) + uint32(v107) + uint32(v143))))) & 248)) * 4)
					v49 = int16((int32(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, 1189588) + uint32(v104) + uint32(v143))))) >> 3) | int32(v48))
					if v147 == 1 {
						v50 = (*uint16)(unsafe.Add(unsafe.Pointer(v34), 1))
						*v39 = uint8(int8(v133))
						*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(v50))), -1))) = 2
						v39 = (*uint8)(unsafe.Pointer(v50))
						v50 = (*uint16)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(v50))), 1))))
						v35 += 4
						v147 = 2
						*v50 = uint16(v49)
						v34 = (*uint8)(unsafe.Pointer((*uint16)(unsafe.Add(unsafe.Pointer(v50), unsafe.Sizeof(uint16(0))*1))))
						v139 = v39
						v133 = 1
						v36 += 3
					} else {
						if func() int32 {
							p := &v133
							*p++
							return *p
						}() == 254 {
							*v39 = 253
							*v34 = 2
							v35 += 2
							v39 = (*uint8)(unsafe.Add(unsafe.Pointer(v34), 1))
							v34 = (*uint8)(unsafe.Add(unsafe.Pointer(v34), 2))
							v133 = 1
							v139 = v39
							v36 += 2
						}
						*(*uint16)(unsafe.Pointer(v34)) = uint16(v49)
						v34 = (*uint8)(unsafe.Add(unsafe.Pointer(v34), 2))
						v35 += 2
						v36++
					}
					goto LABEL_96
				}
				if v147 != 1 {
					break
				}
				v45 = v133 + 1
				v133 = v45
				v46 = v45 == 254
				v39 = v139
				if v46 {
					*v139 = 253
					*v34 = 1
					v47 = int32(uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v34), 1)))))
				LABEL_85:
					v39 = (*uint8)(unsafe.Pointer(uintptr(v47)))
					v34 = (*uint8)(unsafe.Pointer(uintptr(v47 + 1)))
					v35 += 2
					v133 = 1
					v139 = v39
					v36 += 2
				}
			LABEL_96:
				if func() int32 {
					p := &v143
					*p++
					return *p
				}() > v109 {
				LABEL_97:
					*v39 = uint8(int8(v133))
					goto LABEL_206
				}
			}
			v47 = int32(uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v34), 1)))))
			v147 = 1
			*v139 = uint8(int8(v133))
			*(*uint8)(unsafe.Pointer(uintptr(v47 - 1))) = 1
			goto LABEL_85
		case 1:
			v101 = v37 * 640
			v111 = *(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, 1189588) + uint32(v37*640) + uint32(v18))))
			if int32(v128) > int32(v111)+5 || int32(v128) < int32(v111)-5 || int32(v129) > int32(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, 1189588) + uint32((v37*5+3200)<<7) + uint32(v18)))))+5 || int32(v129) < int32(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, 1189588) + uint32((v37*5+3200)<<7) + uint32(v18)))))-5 || int32(v130) > int32(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, 1189588) + uint32((v37*5+6400)<<7) + uint32(v18)))))+5 || int32(v130) < int32(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, 1189588) + uint32((v37*5+6400)<<7) + uint32(v18)))))-5 {
				v52 = int8(*(*uint8)(unsafe.Pointer(uintptr(uint32(v101) + uint32(uintptr(dword_5d4594_1189592)) + uint32(v18)))))
				if int32(v52) == -1 {
					*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v55))), unsafe.Sizeof(int16(0))-1)) = 0
					v56 = int16((((int32(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, 1189588) + uint32(v37*640) + uint32(v18))))) & 248) * 32) | (int32(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, 1189588) + uint32((v37*5+3200)<<7) + uint32(v18))))) & 248)) * 4)
					*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v55))), 0)) = uint8(int8(int32(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, 1189588) + uint32((v37*5+6400)<<7) + uint32(v18))))) >> 3))
					v57 = int16(int32(v55) | int32(v56))
					*v34 = 2
					v58 = (*uint16)(unsafe.Add(unsafe.Pointer(v34), 1))
					v148 = 2
					v51 = (*uint8)(unsafe.Pointer(v58))
					v58 = (*uint16)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(v58))), 1))))
					v134 = 1
					v140 = v51
					*v58 = uint16(v57)
					v34 = (*uint8)(unsafe.Pointer((*uint16)(unsafe.Add(unsafe.Pointer(v58), unsafe.Sizeof(uint16(0))*1))))
					v35 += 4
				} else {
					v134 = 1
					*v34 = uint8(int8((int32(v52) * 16) | 4))
					v53 = (*uint8)(unsafe.Add(unsafe.Pointer(v34), 1))
					v54 = int32(*(*uint8)(unsafe.Pointer(uintptr(uint32(uintptr(dword_5d4594_1189592)) + uint32(v101) + uint32(v18)))))
					v18 = v106
					v148 = (v54 * 16) | 4
					v51 = func() *uint8 {
						p := &v53
						x := *p
						*p = (*uint8)(unsafe.Add(unsafe.Pointer(*p), 1))
						return x
					}()
					v140 = v51
					*v53 = *(*uint8)(unsafe.Pointer(uintptr(uint32(v101) + *memmap.PtrUint32(6112660, 1189588) + uint32(v106))))
					v34 = (*uint8)(unsafe.Add(unsafe.Pointer(v53), 1))
					v35 += 3
				}
				v36 += 3
			} else {
				v35 += 2
				*v34 = 1
				v148 = 1
				v134 = 1
				v51 = (*uint8)(unsafe.Add(unsafe.Pointer(v34), 1))
				v34 = (*uint8)(unsafe.Add(unsafe.Pointer(v34), 2))
				v140 = v51
				v36 += 2
			}
			v144 = v18 + 1
			if v18+1 > v109 {
				goto LABEL_141
			}
			for {
				v116 = *(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, 1189588) + uint32(v101) + uint32(v144))))
				if int32(v128) <= int32(v116)+5 && int32(v128) >= int32(v116)-5 {
					if int32(v129) <= int32(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, 1189588) + uint32((v37*5+3200)<<7) + uint32(v144)))))+5 && int32(v129) >= int32(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, 1189588) + uint32((v37*5+3200)<<7) + uint32(v144)))))-5 && int32(v130) <= int32(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, 1189588) + uint32((v37*5+6400)<<7) + uint32(v144)))))+5 && int32(v130) >= int32(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, 1189588) + uint32((v37*5+6400)<<7) + uint32(v144)))))-5 {
						if v148 == 1 {
							v59 = v134 + 1
							v134 = v59
							v46 = v59 == 254
							v51 = v140
							if !v46 {
								goto LABEL_140
							}
							*v140 = 253
							*v34 = 1
							v60 = int32(uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v34), 1)))))
						} else {
							v60 = int32(uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v34), 1)))))
							v148 = 1
							*v140 = uint8(int8(v134))
							*(*uint8)(unsafe.Pointer(uintptr(v60 - 1))) = 1
						}
						v51 = (*uint8)(unsafe.Pointer(uintptr(v60)))
						v34 = (*uint8)(unsafe.Pointer(uintptr(v60 + 1)))
						v35 += 2
						v134 = 1
						v140 = v51
						v36 += 2
						goto LABEL_140
					}
					v51 = v140
				}
				v61 = int8(*(*uint8)(unsafe.Pointer(uintptr(uint32(v101) + uint32(uintptr(dword_5d4594_1189592)) + uint32(v144)))))
				if int32(v61) == -1 {
					v63 = int16((((int32(v116) & 248) * 32) | (int32(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, 1189588) + uint32((v37*5+3200)<<7) + uint32(v144))))) & 248)) * 4)
					v64 = int16((int32(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, 1189588) + uint32((v37*5+6400)<<7) + uint32(v144))))) >> 3) | int32(v63))
					if v148 != 2 {
						v65 = (*uint16)(unsafe.Add(unsafe.Pointer(v34), 1))
						*v51 = uint8(int8(v134))
						*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(v65))), -1))) = 2
						v51 = (*uint8)(unsafe.Pointer(v65))
						v65 = (*uint16)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(v65))), 1))))
						v35 += 4
						v140 = v51
						*v65 = uint16(v64)
						v34 = (*uint8)(unsafe.Pointer((*uint16)(unsafe.Add(unsafe.Pointer(v65), unsafe.Sizeof(uint16(0))*1))))
						v148 = 2
						v134 = 1
						v36 += 3
						goto LABEL_140
					}
					if func() int32 {
						p := &v134
						*p++
						return *p
					}() == 254 {
						*v51 = 253
						*v34 = 2
						v35 += 2
						v51 = (*uint8)(unsafe.Add(unsafe.Pointer(v34), 1))
						v34 = (*uint8)(unsafe.Add(unsafe.Pointer(v34), 2))
						v134 = 1
						v140 = v51
						v36 += 2
					}
					*(*uint16)(unsafe.Pointer(v34)) = uint16(v64)
					v34 = (*uint8)(unsafe.Add(unsafe.Pointer(v34), 2))
					v35 += 2
				} else {
					v121 = uint8(int8((int32(v61) * 16) | 4))
					if v148 != int32(v121) {
						v148 = int32(v121)
						*v51 = uint8(int8(v134))
						*v34 = v121
						v62 = (*uint8)(unsafe.Add(unsafe.Pointer(v34), 1))
						v51 = func() *uint8 {
							p := &v62
							x := *p
							*p = (*uint8)(unsafe.Add(unsafe.Pointer(*p), 1))
							return x
						}()
						v35 += 3
						*v62 = *(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, 1189588) + uint32(v101) + uint32(v144))))
						v34 = (*uint8)(unsafe.Add(unsafe.Pointer(v62), 1))
						v134 = 1
						v140 = v51
						v36 += 3
						goto LABEL_140
					}
					if func() int32 {
						p := &v134
						*p++
						return *p
					}() == 254 {
						*v51 = 253
						*v34 = v121
						v51 = (*uint8)(unsafe.Add(unsafe.Pointer(v34), 1))
						v34 = (*uint8)(unsafe.Add(unsafe.Pointer(v34), 2))
						v35 += 2
						v134 = 1
						v140 = v51
						v36 += 2
					}
					v34 = (*uint8)(unsafe.Add(unsafe.Pointer(v34), 1))
					v35++
					*((*uint8)(unsafe.Add(unsafe.Pointer(v34), -1))) = *(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, 1189588) + uint32(v101) + uint32(v144))))
				}
				v36++
			LABEL_140:
				v144++
				if v144 > v109 {
					break
				}
			}
		LABEL_141:
			*v51 = uint8(int8(v134))
		LABEL_206:
			v18 = v106
		LABEL_207:
			if func() int32 {
				p := &v37
				*p++
				return *p
			}() <= v123 {
				v38 = v125
				continue
			}
		LABEL_208:
			*v34 = 0
			*(*uint8)(unsafe.Add(unsafe.Pointer(v34), 1)) = 0
			return 1
		case 2:
			v102 = v37 * 640
			v112 = *(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, 1189588) + uint32(v37*640) + uint32(v18))))
			if int32(v128) > int32(v112)+5 || int32(v128) < int32(v112)-5 || (func() bool {
				v108 = (v37*5 + 3200) << 7
				v135 = int32(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, 1189588) + uint32(v108) + uint32(v18)))))
				return int32(v129) > v135+5
			}()) || int32(v129) < v135-5 || (func() bool {
				v105 = (v37*5 + 6400) << 7
				v136 = int32(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, 1189588) + uint32(v105) + uint32(v18)))))
				return int32(v130) > v136+5
			}()) || int32(v130) < v136-5 {
				v108 = (v37*5 + 3200) << 7
				v66 = (v37*5 + 6400) << 7
				v105 = v66
				v35 += 4
				*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v66))), 0)) = *(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, 1189588) + uint32(v66) + uint32(v18))))
				v18 = v106
				v149 = 5
				v137 = 1
				v126 = int16((int32(*(*uint8)(unsafe.Pointer(uintptr(uint32(v102) + uint32(uintptr(dword_5d4594_1189596)) + uint32(v106))))) >> 4) | (v66 & 240))
				v67 = int8(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, 1189588) + uint32(v108) + uint32(v106)))))
				*v34 = 5
				v68 = (*uint16)(unsafe.Add(unsafe.Pointer(v34), 1))
				v141 = (*uint8)(unsafe.Pointer(v68))
				v68 = (*uint16)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(v68))), 1))))
				*v68 = uint16(int16(((((int32(v112) & 240) * 16) | (int32(v67) & 240)) * 16) | int32(v126)))
				v34 = (*uint8)(unsafe.Pointer((*uint16)(unsafe.Add(unsafe.Pointer(v68), unsafe.Sizeof(uint16(0))*1))))
				v36 += 3
			} else {
				v35 += 2
				*v34 = 1
				v141 = (*uint8)(unsafe.Add(unsafe.Pointer(v34), 1))
				v34 = (*uint8)(unsafe.Add(unsafe.Pointer(v34), 2))
				v149 = 1
				v137 = 1
				v36 += 2
			}
			v69 = v18 + 1
			v145 = v18 + 1
			if v18+1 <= v109 {
				for {
					v117 = *(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, 1189588) + uint32(v102) + uint32(v69))))
					if int32(v128) > int32(v117)+5 || int32(v128) < int32(v117)-5 || (func() bool {
						v70 = int32(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, 1189588) + uint32(v108) + uint32(v145)))))
						return int32(v129) > v70+5
					}()) || int32(v129) < v70-5 || (func() bool {
						v71 = int32(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, 1189588) + uint32(v105) + uint32(v145)))))
						return int32(v130) > v71+5
					}()) || int32(v130) < v71-5 {
						v75 = int32(*memmap.PtrUint32(6112660, 1189588) + uint32(v105))
						v76 = int32(*memmap.PtrUint32(6112660, 1189588) + uint32(v108))
						*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v77))), unsafe.Sizeof(uint16(0))*1)) = *(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v145))), unsafe.Sizeof(uint16(0))*1))
						*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v75))), unsafe.Sizeof(uint16(0))*0)) = uint16(int16((int32(*(*uint8)(unsafe.Pointer(uintptr(uint32(uintptr(dword_5d4594_1189596)) + uint32(v102) + uint32(v145))))) >> 4) | (int32(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, 1189588) + uint32(v105) + uint32(v145))))) & 240)))
						*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v76))), unsafe.Sizeof(uint16(0))*0)) = uint16(int16(int32(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, 1189588) + uint32(v108) + uint32(v145))))) & 240))
						*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v77))), unsafe.Sizeof(uint16(0))*0)) = uint16(int16(int32(v117) & 240))
						v78 = (((v77 * 16) | v76) * 16) | v75
						if v149 == 5 {
							if func() int32 {
								p := &v137
								*p++
								return *p
							}() == 254 {
								v80 = v141
								v81 = int32(uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v34), 1)))))
								v141 = (*uint8)(unsafe.Pointer(uintptr(v81)))
								v34 = (*uint8)(unsafe.Pointer(uintptr(v81 + 1)))
								*v80 = 253
								*((*uint8)(unsafe.Add(unsafe.Pointer(v34), -2))) = 5
								v35 += 2
								v137 = 1
								v36 += 2
							}
							*(*uint16)(unsafe.Pointer(v34)) = uint16(int16(v78))
							v34 = (*uint8)(unsafe.Add(unsafe.Pointer(v34), 2))
							v35 += 2
							v36++
						} else {
							v79 = (*uint16)(unsafe.Add(unsafe.Pointer(v34), 1))
							v35 += 4
							*v141 = uint8(int8(v137))
							*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(v79))), -1))) = 5
							v141 = (*uint8)(unsafe.Pointer(v79))
							v79 = (*uint16)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(v79))), 1))))
							v149 = 2
							v137 = 1
							*v79 = uint16(int16(v78))
							v34 = (*uint8)(unsafe.Pointer((*uint16)(unsafe.Add(unsafe.Pointer(v79), unsafe.Sizeof(uint16(0))*1))))
							v36 += 3
						}
					} else if v149 == 1 {
						if func() int32 {
							p := &v137
							*p++
							return *p
						}() == 254 {
							v72 = v141
							v73 = int32(uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v34), 1)))))
							v141 = (*uint8)(unsafe.Pointer(uintptr(v73)))
							v34 = (*uint8)(unsafe.Pointer(uintptr(v73 + 1)))
							*v72 = 253
							*((*uint8)(unsafe.Add(unsafe.Pointer(v34), -2))) = 1
							v35 += 2
							v137 = 1
							v36 += 2
						}
					} else {
						v74 = int32(uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v34), 1)))))
						v35 += 2
						*v141 = uint8(int8(v137))
						*(*uint8)(unsafe.Pointer(uintptr(v74 - 1))) = 1
						v141 = (*uint8)(unsafe.Pointer(uintptr(v74)))
						v34 = (*uint8)(unsafe.Pointer(uintptr(v74 + 1)))
						v149 = 1
						v137 = 1
						v36 += 2
					}
					v69 = func() int32 {
						p := &v145
						*p++
						return *p
					}()
					if v145 > v109 {
						break
					}
				}
				v18 = v106
			}
			*v141 = uint8(int8(v137))
			goto LABEL_207
		case 3:
			v103 = v37 * 640
			v113 = *(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, 1189588) + uint32(v37*640) + uint32(v18))))
			if int32(v128) > int32(v113)+5 || int32(v128) < int32(v113)-5 || int32(v129) > int32(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, 1189588) + uint32((v37*5+3200)<<7) + uint32(v18)))))+5 || int32(v129) < int32(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, 1189588) + uint32((v37*5+3200)<<7) + uint32(v18)))))-5 || int32(v130) > int32(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, 1189588) + uint32((v37*5+6400)<<7) + uint32(v18)))))+5 || int32(v130) < int32(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, 1189588) + uint32((v37*5+6400)<<7) + uint32(v18)))))-5 {
				v83 = int8(*(*uint8)(unsafe.Pointer(uintptr(uint32(v103) + uint32(uintptr(dword_5d4594_1189592)) + uint32(v18)))))
				v35 += 4
				if int32(v83) == -1 {
					v150 = 5
					v86 = int8(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, 1189588) + uint32((v37*5+6400)<<7) + uint32(v18)))))
					v18 = v106
					v127 = int16((int32(*(*uint8)(unsafe.Pointer(uintptr(uint32(v103) + uint32(uintptr(dword_5d4594_1189596)) + uint32(v106))))) >> 4) | (int32(v86) & 240))
					v87 = int8(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, 1189588) + uint32((v37*5+3200)<<7) + uint32(v106)))))
					*v34 = 5
					v88 = (*uint16)(unsafe.Add(unsafe.Pointer(v34), 1))
					v142 = (*uint8)(unsafe.Pointer(v88))
					v88 = (*uint16)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(v88))), 1))))
					v82 = 1
					*v88 = uint16(int16(((((int32(v113) & 240) * 16) | (int32(v87) & 240)) * 16) | int32(v127)))
					v34 = (*uint8)(unsafe.Pointer((*uint16)(unsafe.Add(unsafe.Pointer(v88), unsafe.Sizeof(uint16(0))*1))))
					v138 = 1
					v36 += 3
				} else {
					*v34 = uint8(int8((int32(v83) * 16) | 6))
					v84 = (*uint8)(unsafe.Add(unsafe.Pointer(v34), 1))
					v85 = int32(*(*uint8)(unsafe.Pointer(uintptr(uint32(uintptr(dword_5d4594_1189592)) + uint32(v103) + uint32(v18)))))
					v142 = func() *uint8 {
						p := &v84
						x := *p
						*p = (*uint8)(unsafe.Add(unsafe.Pointer(*p), 1))
						return x
					}()
					*v84 = *(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, 1189588) + uint32(v103) + uint32(v106))))
					v18 = v106
					v84 = (*uint8)(unsafe.Add(unsafe.Pointer(v84), 1))
					v150 = (v85 * 16) | 6
					v82 = 1
					*v84 = *(*uint8)(unsafe.Pointer(uintptr(uint32(v103) + uint32(uintptr(dword_5d4594_1189596)) + uint32(v106))))
					v34 = (*uint8)(unsafe.Add(unsafe.Pointer(v84), 1))
					v138 = 1
					v36 += 4
				}
			} else {
				v82 = 1
				v35 += 2
				*v34 = 1
				v142 = (*uint8)(unsafe.Add(unsafe.Pointer(v34), 1))
				v34 = (*uint8)(unsafe.Add(unsafe.Pointer(v34), 2))
				v150 = 1
				v138 = 1
				v36 += 2
			}
			v146 = v18 + 1
			if v18+1 > v109 {
				goto LABEL_205
			}
			for {
				v118 = *(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, 1189588) + uint32(v103) + uint32(v146))))
				if int32(v128) > int32(v118)+5 || int32(v128) < int32(v118)-5 {
					goto LABEL_192
				}
				if int32(v129) > int32(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, 1189588) + uint32((v37*5+3200)<<7) + uint32(v146)))))+5 || int32(v129) < int32(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, 1189588) + uint32((v37*5+3200)<<7) + uint32(v146)))))-5 || int32(v130) > int32(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, 1189588) + uint32((v37*5+6400)<<7) + uint32(v146)))))+5 || int32(v130) < int32(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, 1189588) + uint32((v37*5+6400)<<7) + uint32(v146)))))-5 {
					break
				}
				if v150 != 1 {
					v89 = int32(uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v34), 1)))))
					v150 = 1
					*v142 = uint8(int8(v138))
					*(*uint8)(unsafe.Pointer(uintptr(v89 - 1))) = 1
					goto LABEL_189
				}
				v82 = func() int32 {
					p := &v138
					*p++
					return *p
				}()
				if v138 == 254 {
					v89 = int32(uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v34), 1)))))
					*v142 = 253
					*(*uint8)(unsafe.Pointer(uintptr(v89 - 1))) = 1
				LABEL_189:
					v82 = 1
					v142 = (*uint8)(unsafe.Pointer(uintptr(v89)))
					v34 = (*uint8)(unsafe.Pointer(uintptr(v89 + 1)))
					v35 += 2
					v138 = 1
					v36 += 2
				}
			LABEL_204:
				if func() int32 {
					p := &v146
					*p++
					return *p
				}() > v109 {
				LABEL_205:
					*v142 = uint8(int8(v82))
					goto LABEL_206
				}
			}
			v82 = v138
		LABEL_192:
			v90 = int8(*(*uint8)(unsafe.Pointer(uintptr(uint32(v103) + uint32(uintptr(dword_5d4594_1189592)) + uint32(v146)))))
			if int32(v90) == -1 {
				*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v94))), unsafe.Sizeof(uint16(0))*1)) = *memmap.PtrUint16(6112660, 1189590)
				*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v95))), unsafe.Sizeof(uint16(0))*1)) = *(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v146))), unsafe.Sizeof(uint16(0))*1))
				v96 = int32(*memmap.PtrUint32(6112660, 1189588) + uint32((v37*5+3200)<<7))
				*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v94))), unsafe.Sizeof(uint16(0))*0)) = uint16(int16((int32(*(*uint8)(unsafe.Pointer(uintptr(uint32(uintptr(dword_5d4594_1189596)) + uint32(v103) + uint32(v146))))) >> 4) | (int32(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, 1189588) + uint32((v37*5+6400)<<7) + uint32(v146))))) & 240)))
				*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v96))), unsafe.Sizeof(uint16(0))*0)) = uint16(int16(int32(*(*uint8)(unsafe.Pointer(uintptr(v96 + v146)))) & 240))
				*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v95))), unsafe.Sizeof(uint16(0))*0)) = uint16(int16(int32(v118) & 240))
				v97 = (((v95 * 16) | v96) * 16) | v94
				if v150 == 5 {
					v82 = func() int32 {
						p := &v138
						*p++
						return *p
					}()
					if v138 == 254 {
						v99 = v142
						v100 = int32(uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v34), 1)))))
						v142 = (*uint8)(unsafe.Pointer(uintptr(v100)))
						v34 = (*uint8)(unsafe.Pointer(uintptr(v100 + 1)))
						*v99 = 253
						*((*uint8)(unsafe.Add(unsafe.Pointer(v34), -2))) = 5
						v82 = 1
						v35 += 2
						v138 = 1
						v36 += 2
					}
					*(*uint16)(unsafe.Pointer(v34)) = uint16(int16(v97))
					v34 = (*uint8)(unsafe.Add(unsafe.Pointer(v34), 2))
					v35 += 2
					goto LABEL_203
				}
				v98 = (*uint16)(unsafe.Add(unsafe.Pointer(v34), 1))
				v35 += 4
				*v142 = uint8(int8(v138))
				*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(v98))), -1))) = 5
				v142 = (*uint8)(unsafe.Pointer(v98))
				v98 = (*uint16)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(v98))), 1))))
				v82 = 1
				v150 = 5
				*v98 = uint16(int16(v97))
				v34 = (*uint8)(unsafe.Pointer((*uint16)(unsafe.Add(unsafe.Pointer(v98), unsafe.Sizeof(uint16(0))*1))))
				v138 = 1
				v36 += 3
			} else {
				v122 = uint8(int8((int32(v90) * 16) | 6))
				if v150 == int32(v122) {
					v138 = func() int32 {
						p := &v82
						*p++
						return *p
					}()
					if v82 == 254 {
						v92 = int32(uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v34), 1)))))
						v35 += 2
						*v142 = 253
						*(*uint8)(unsafe.Pointer(uintptr(v92 - 1))) = v122
						v82 = 1
						v142 = (*uint8)(unsafe.Pointer(uintptr(v92)))
						v34 = (*uint8)(unsafe.Pointer(uintptr(v92 + 1)))
						v138 = 1
						v36 += 2
					}
					v93 = int32(uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v34), 1)))))
					*(*uint8)(unsafe.Pointer(uintptr(v93 - 1))) = *(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, 1189588) + uint32(v103) + uint32(v146))))
					v34 = (*uint8)(unsafe.Pointer(uintptr(v93 + 1)))
					v35++
					*((*uint8)(unsafe.Add(unsafe.Pointer(v34), -1))) = *(*uint8)(unsafe.Pointer(uintptr(uint32(uintptr(dword_5d4594_1189596)) + uint32(v103) + uint32(v146))))
				LABEL_203:
					v36++
					goto LABEL_204
				}
				v150 = int32(v122)
				v91 = (*uint8)(unsafe.Add(unsafe.Pointer(v34), 1))
				v35 += 4
				*v142 = uint8(int8(v82))
				*((*uint8)(unsafe.Add(unsafe.Pointer(v91), -1))) = v122
				v142 = func() *uint8 {
					p := &v91
					x := *p
					*p = (*uint8)(unsafe.Add(unsafe.Pointer(*p), 1))
					return x
				}()
				v82 = 1
				*func() *uint8 {
					p := &v91
					x := *p
					*p = (*uint8)(unsafe.Add(unsafe.Pointer(*p), 1))
					return x
				}() = *(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, 1189588) + uint32(v103) + uint32(v146))))
				v138 = 1
				*v91 = *(*uint8)(unsafe.Pointer(uintptr(uint32(uintptr(dword_5d4594_1189596)) + uint32(v103) + uint32(v146))))
				v34 = (*uint8)(unsafe.Add(unsafe.Pointer(v91), 1))
				v36 += 4
			}
			goto LABEL_204
		default:
			goto LABEL_207
		}
	}
}
func sub_47D090(a1 int32, a2 int32) int32 {
	var (
		result int32
		v3     int32
		v4     int32
		v5     int32
		v6     int32
		v7     int8
		v8     *uint8
		v9     int32
		v10    int32
	)
	result = a2
	if a2 > 0 {
		v3 = int32(uintptr(dword_5d4594_1189592))
		v4 = 0
		v9 = a2
		for {
			v5 = 0
			if a1 > 0 {
				v6 = v4
				for {
					v7 = 0
					v8 = (*uint8)(memmap.PtrOff(0x587000, 153689))
					v10 = 0
					for {
						if int32(*(*uint8)(unsafe.Pointer(uintptr(v3 + v6)))) == int32(*((*uint8)(unsafe.Add(unsafe.Pointer(v8), -1)))) && int32(*(*uint8)(unsafe.Pointer(uintptr(uint32(v4+v3) + 0xC8000 + uint32(v5) - 0x64000)))) == int32(*v8) && int32(*(*uint8)(unsafe.Pointer(uintptr(uint32(v4+v3) + 0xC8000 + uint32(v5))))) == int32(*(*uint8)(unsafe.Add(unsafe.Pointer(v8), 1))) {
							*(*uint8)(unsafe.Pointer(uintptr(v3 + v6))) = uint8(v7)
							v3 = int32(uintptr(dword_5d4594_1189592))
							v10 = 1
						}
						v8 = (*uint8)(unsafe.Add(unsafe.Pointer(v8), 3))
						v7++
						if int32(uintptr(unsafe.Pointer(v8))) >= int32(uintptr(memmap.PtrOff(0x587000, 153710))) {
							break
						}
					}
					if v10 == 0 {
						*(*uint8)(unsafe.Pointer(uintptr(v4 + v3 + v5))) = math.MaxUint8
						v3 = int32(uintptr(dword_5d4594_1189592))
					}
					v5++
					v6++
					if v5 >= a1 {
						break
					}
				}
			}
			v4 += 640
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
func sub_47D380(a1 int32, a2 int32) int32 {
	var (
		v2 int32
		v3 int32
		v4 int32
		v5 int32
	)
	v2 = a1
	v3 = a2
	if a1 > a2 {
		v2 = a2
		v3 = a1
	}
	v4 = nox_draw_curDrawData_3799572.clip.left
	if v2 >= v4 {
		if v2 >= nox_draw_curDrawData_3799572.clip.right {
			return 0
		}
	} else {
		v2 = nox_draw_curDrawData_3799572.clip.left
	}
	v5 = nox_draw_curDrawData_3799572.clip.right
	if v3 < v5 {
		if v3 < v4 {
			return 0
		}
	} else {
		v3 = nox_draw_curDrawData_3799572.clip.right
	}
	if v2 == v3 {
		return 0
	}
	if v2 != v4 || v3 != v5 {
		*memmap.PtrUint32(0x973F18, 52) = uint32(v2)
		*memmap.PtrUint32(0x973F18, 12) = uint32(v3)
		dword_5d4594_3799452 = 1
	}
	return 1
}
func sub_47D3F0(a1 int32) int32 {
	var result int32
	result = a1
	*memmap.PtrUint32(0x973F18, 44) = uint32(a1)
	return result
}
func nox_xxx_video_ReadTile_Real_47D4E0(a1 *uint32) int16 {
	var (
		v1     int32
		v2     *uint16
		result int16
	)
	v1 = int32(*a1 * *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*1)))
	v2 = (*uint16)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(a1))), 17))))
	for {
		result = int16(int32(*v2)&31 | ((int32(*v2) * 2) & 0xFFC0))
		*v2 = uint16(result)
		v2 = (*uint16)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint16(0))*1))
		if func() int32 {
			p := &v1
			x := *p
			*p--
			return x
		}() <= 1 {
			break
		}
	}
	return result
}
func nox_xxx_video_ReadSprite_Real_47D530(data unsafe.Pointer) {
	var (
		v2     int32 = 0
		v6     bool
		width  int32  = *((*int32)(unsafe.Add(unsafe.Pointer((*int32)(data)), unsafe.Sizeof(int32(0))*0)))
		height int32  = *((*int32)(unsafe.Add(unsafe.Pointer((*int32)(data)), unsafe.Sizeof(int32(0))*1)))
		pix    *int16 = (*int16)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(data)), 17))))
		h      int32  = height
	)
	for {
		var (
			op uint8 = *((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(pix))), 0)))
			b  uint8 = *((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(pix))), 1)))
		)
		pix = (*int16)(unsafe.Add(unsafe.Pointer(pix), unsafe.Sizeof(int16(0))*1))
		op &= 15
		if int32(op) == 2 {
			var n int32 = int32(b)
			for {
				*pix = int16(int32(*pix)&31 | (int32(*pix)&0x7FE0)*2)
				pix = (*int16)(unsafe.Add(unsafe.Pointer(pix), unsafe.Sizeof(int16(0))*1))
				v6 = func() int32 {
					p := &n
					x := *p
					*p--
					return x
				}() <= 1
				if v6 {
					break
				}
			}
			v6 = v2 <= int32(b)
			v2 -= int32(b)
			if v6 {
				v6 = func() int32 {
					p := &h
					x := *p
					*p--
					return x
				}() <= 1
				if v6 {
					return
				}
				v2 = width
			}
		} else if int32(op) == 4 {
			pix = (*int16)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(pix))), b))))
			v6 = v2 <= int32(b)
			v2 -= int32(b)
			if v6 {
				v6 = func() int32 {
					p := &h
					x := *p
					*p--
					return x
				}() <= 1
				if v6 {
					return
				}
				v2 = width
			}
		} else if int32(op) == 5 || int32(op) == 6 {
			pix = (*int16)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(pix))), b))))
			pix = (*int16)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(pix))), b))))
			v6 = v2 <= int32(b)
			v2 -= int32(b)
			if v6 {
				v6 = func() int32 {
					p := &h
					x := *p
					*p--
					return x
				}() <= 1
				if v6 {
					return
				}
				v2 = width
			}
		}
	}
}
func sub_47D5C0(a1 int32, a2 *uint32, a3 *uint32, a4 *uint32, a5 *uint32) int32 {
	var (
		v5 *uint32
		v6 *uint32
		v7 *uint32
	)
	if a4 != nil {
		*a4 = 0
	}
	if a5 != nil {
		*a5 = 0
	}
	if a1 == 0 {
		return 0
	}
	v5 = (*uint32)(nox_video_getImagePixdata_func((*nox_video_bag_image_t)(unsafe.Pointer(uintptr(a1)))))
	if v5 == nil {
		return 0
	}
	if a4 != nil {
		*a4 = *v5
	}
	v6 = (*uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(uint32(0))*1))
	if a5 != nil {
		*a5 = *v6
	}
	v7 = (*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*1))
	if a2 != nil {
		*a2 += *v7
	}
	if a3 != nil {
		*a3 += *(*uint32)(unsafe.Add(unsafe.Pointer(v7), unsafe.Sizeof(uint32(0))*1))
	}
	return 1
}
func sub_47DBC0() uint8 {
	return *memmap.PtrUint8(6112660, 1193128)
}
func sub_47FCE0(a1 *uint32, a2 int32) int32 {
	var (
		v2 int32
		v3 *uint8
		v4 int32
		v5 int32
	)
	v2 = 0
	if *(*int32)(unsafe.Pointer(&dword_5d4594_3804684)) <= 0 {
	LABEL_8:
		v4 = int32(dword_5d4594_3804684 * 16)
		v5 = int32(dword_5d4594_3804684 + 1)
		*memmap.PtrUint32(0x973F18, uint32(v4+6088)) = *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*3))
		*memmap.PtrUint32(0x973F18, uint32(v4+6092)) = *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*2))
		*memmap.PtrUint32(0x973F18, uint32(v4+6096)) = *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*21))
		*memmap.PtrUint32(0x973F18, uint32(v4+6100)) = *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*26))
		dword_5d4594_3804684 = uint32(v5)
	} else {
		v3 = (*uint8)(memmap.PtrOff(0x973F18, 6092))
		for *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v3))), -int(unsafe.Sizeof(uint32(0))*1)))) != *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*3)) || *(*uint32)(unsafe.Pointer(v3)) != *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*2)) || *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v3))), unsafe.Sizeof(uint32(0))*1))) != *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*21)) || *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v3))), unsafe.Sizeof(uint32(0))*2))) != *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*26)) {
			v2++
			v3 = (*uint8)(unsafe.Add(unsafe.Pointer(v3), 16))
			if v2 >= *(*int32)(unsafe.Pointer(&dword_5d4594_3804684)) {
				goto LABEL_8
			}
		}
	}
	return 1
}
func nox_xxx_setupSomeVideo_47FEF0() int32 {
	var (
		v0     *uint8
		result int32
		v2     *uint8
	)
	_ = v2
	*memmap.PtrUint32(0x973F18, 7716) = uint32(cgoFuncAddr(sub_480250))
	*memmap.PtrUint32(0x973F18, 5248) = uint32(cgoFuncAddr(sub_480220))
	v0 = (*uint8)(memmap.PtrOff(6112660, 2598060))
	for {
		*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v0))), -int(unsafe.Sizeof(uint16(0))*16)))) = 0x7C00
		*(*uint16)(unsafe.Pointer(v0)) = 992
		*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v0))), unsafe.Sizeof(uint16(0))*16))) = 31
		*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v0))), -int(unsafe.Sizeof(uint16(0))*12)))) = 0xF800
		*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v0))), unsafe.Sizeof(uint16(0))*4))) = 0xFC00
		*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v0))), unsafe.Sizeof(uint16(0))*20))) = 0xF800
		*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v0))), -int(unsafe.Sizeof(uint16(0))*8)))) = 0xF800
		*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v0))), unsafe.Sizeof(uint16(0))*8))) = 2016
		*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v0))), unsafe.Sizeof(uint16(0))*24))) = 31
		*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v0))), -int(unsafe.Sizeof(uint16(0))*4)))) = 248
		*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v0))), unsafe.Sizeof(uint16(0))*12))) = 252
		*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v0))), unsafe.Sizeof(uint16(0))*28))) = 248
		v0 = (*uint8)(unsafe.Add(unsafe.Pointer(v0), 2))
		if int32(uintptr(unsafe.Pointer(v0))) >= int32(uintptr(memmap.PtrOff(6112660, 2598068))) {
			break
		}
	}
	*memmap.PtrUint32(6112660, 2598128) = 0
	*memmap.PtrUint32(6112660, 2598160) = 0
	*memmap.PtrUint32(6112660, 2598132) = 0
	*memmap.PtrUint32(6112660, 2598136) = 0
	*memmap.PtrUint32(6112660, 2598168) = 0
	*memmap.PtrUint32(6112660, 2598144) = 0
	*memmap.PtrUint32(6112660, 2598176) = 0
	*memmap.PtrUint32(6112660, 2598152) = 0
	*memmap.PtrUint32(6112660, 2598184) = 0
	*memmap.PtrUint32(6112660, 2598124) = 7
	*memmap.PtrUint32(6112660, 2598156) = 2
	*memmap.PtrUint32(6112660, 2598164) = 5
	*memmap.PtrUint32(6112660, 2598140) = 8
	*memmap.PtrUint32(6112660, 2598172) = 3
	*memmap.PtrUint32(6112660, 2598148) = 8
	*memmap.PtrUint32(6112660, 2598180) = 3
	dword_5d4594_3807160 = uint32(cgoFuncAddr(sub_480EB0))
	result = 1
	return result
}
func sub_4801B0(a2 *uint8, a3 *uint8) int8 {
	var (
		v2     uint32
		v3     uint32
		result int8
	)
	*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v3))), unsafe.Sizeof(uint16(0))*1)) = *(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&a3))), unsafe.Sizeof(uint16(0))*1))
	*a2 = uint8(int8(int32(*a3) * 8))
	*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v2))), unsafe.Sizeof(uint16(0))*0)) = *(*uint16)(unsafe.Pointer(a3))
	*(*uint8)(unsafe.Add(unsafe.Pointer(a2), 1)) = uint8((v2 >> 2) & 248)
	*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v3))), unsafe.Sizeof(uint16(0))*0)) = *(*uint16)(unsafe.Pointer(a3))
	result = int8(uint8((v3 >> 7) & 248))
	*(*uint8)(unsafe.Add(unsafe.Pointer(a2), 2)) = uint8(result)
	return result
}
func sub_4801E0(a1 *uint8, a2 *uint16) *uint16 {
	var result *uint16
	result = a2
	*a2 = uint16(int16((int32(*a1) >> 3) | (int32(*(*uint8)(unsafe.Add(unsafe.Pointer(a1), 1)))&248|(int32(*(*uint8)(unsafe.Add(unsafe.Pointer(a1), 2)))&248)*32)*4))
	return result
}
func sub_480220(a1 *uint8, a2 *uint8) *uint8 {
	var (
		v2     uint32
		result *uint8
	)
	result = a2
	*a1 = uint8(int8(int32(*a2) * 8))
	*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v2))), unsafe.Sizeof(uint16(0))*0)) = *(*uint16)(unsafe.Pointer(a2))
	*(*uint8)(unsafe.Add(unsafe.Pointer(a1), 1)) = uint8((v2 >> 3) & 252)
	*(*uint8)(unsafe.Add(unsafe.Pointer(a1), 2)) = uint8(int8(int32(*(*uint8)(unsafe.Add(unsafe.Pointer(a2), 1))) & 248))
	return result
}
func sub_480250(a1 *uint8, a2 *uint16) *uint16 {
	var result *uint16
	result = a2
	*a2 = uint16(int16((int32(*a1) >> 3) | (int32(*(*uint8)(unsafe.Add(unsafe.Pointer(a1), 1)))&252|(int32(*(*uint8)(unsafe.Add(unsafe.Pointer(a1), 2)))&248)*32)*8))
	return result
}
func sub_480290(a1 unsafe.Pointer, a2 int32, a3 int32, a4 *uint32, a5 *uint32) uint32 {
	panic("abort")
	return 0
}
func sub_4805A0(a1 int32, a2 unsafe.Pointer, a3 int32, a4 *uint32, a5 *uint32) int16 {
	panic("abort")
	return 0
}
func sub_480700(a1 int32, a2 unsafe.Pointer, a3 int32, a4 *uint32, a5 *uint32) int16 {
	panic("abort")
	return 0
}
func sub_480860(a1 *uint16, a2 *uint16, a3 int32, a4 *int32, a5 *int32) {
	var (
		dword_975384 *uint16 = (*uint16)(unsafe.Pointer(nox_draw_colors_g_3804656))
		dword_97538C *uint16 = (*uint16)(unsafe.Pointer(nox_draw_colors_b_3804664))
		dword_975394 *uint16 = (*uint16)(unsafe.Pointer(nox_draw_colors_r_3804672))
		dword_975260 uint16  = *((*uint16)(unsafe.Pointer(&byte_5D4594_3804364[0+0])))
		dword_975264 uint16  = *((*uint16)(unsafe.Pointer(&byte_5D4594_3804364[0+4])))
		dword_975268 uint16  = *((*uint16)(unsafe.Pointer(&byte_5D4594_3804364[0+8])))
		dword_97526C uint16  = *((*uint16)(unsafe.Pointer(&byte_5D4594_3804364[0+12])))
		dword_975270 uint16  = *((*uint16)(unsafe.Pointer(&byte_5D4594_3804364[0+16])))
		dword_975274 uint16  = *((*uint16)(unsafe.Pointer(&byte_5D4594_3804364[0+20])))
	)
	for a3 > 0 {
		*a1 = uint16(int16(int32(*(*uint16)(unsafe.Add(unsafe.Pointer(dword_97538C), unsafe.Sizeof(uint16(0))*uintptr(((*(*int32)(unsafe.Add(unsafe.Pointer(a4), unsafe.Sizeof(int32(0))*2))*((int32(*a2)&int32(dword_975268))<<int32(dword_975274)))>>16)&math.MaxUint8)))) | int32(*(*uint16)(unsafe.Add(unsafe.Pointer(dword_975384), unsafe.Sizeof(uint16(0))*uintptr(((*(*int32)(unsafe.Add(unsafe.Pointer(a4), unsafe.Sizeof(int32(0))*1))*((int32(*a2)&int32(dword_975264))>>int32(dword_975270)))>>16)&math.MaxUint8)))) | int32(*(*uint16)(unsafe.Add(unsafe.Pointer(dword_975394), unsafe.Sizeof(uint16(0))*uintptr(((*(*int32)(unsafe.Add(unsafe.Pointer(a4), unsafe.Sizeof(int32(0))*0))*((int32(*a2)&int32(dword_975260))>>int32(dword_97526C)))>>16)&math.MaxUint8))))))
		*(*int32)(unsafe.Add(unsafe.Pointer(a4), unsafe.Sizeof(int32(0))*0)) += *(*int32)(unsafe.Add(unsafe.Pointer(a5), unsafe.Sizeof(int32(0))*0))
		*(*int32)(unsafe.Add(unsafe.Pointer(a4), unsafe.Sizeof(int32(0))*1)) += *(*int32)(unsafe.Add(unsafe.Pointer(a5), unsafe.Sizeof(int32(0))*1))
		*(*int32)(unsafe.Add(unsafe.Pointer(a4), unsafe.Sizeof(int32(0))*2)) += *(*int32)(unsafe.Add(unsafe.Pointer(a5), unsafe.Sizeof(int32(0))*2))
		a1 = (*uint16)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint16(0))*1))
		a2 = (*uint16)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint16(0))*1))
		a3--
	}
}
func sub_480950(a1 *int32, a2 *int32, a3 int32, a4 int32, a5 int32) int32 {
	var (
		v5     int32
		result int32
		v7     int32
		v8     uint32
		v9     int32
		v10    int32
		v11    int32
		v12    *int32
		v13    int32
		v14    int32
		v15    int32
		v16    int32
		v17    int32
		v18    unsafe.Pointer
		v19    int32
		v20    int32
		v21    [3]int32
		v22    [3]int32
		v23    unsafe.Pointer
	)
	v5 = *a1
	result = int32(uint32(*(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*1))) - dword_5d4594_3798824)
	v7 = int32(uint32(*a1) - dword_5d4594_3798820)
	if v7 >= 0 && result >= 0 && v7 < *(*int32)(unsafe.Pointer(&dword_5d4594_3798800)) && result < *(*int32)(unsafe.Pointer(&dword_5d4594_3798808)) {
		v8 = dword_5d4594_3798804*(dword_5d4594_3798840+uint32(result)) + dword_5d4594_3798796 + (dword_5d4594_3798836+uint32(v7))*2
		if v8 >= uint32(*(*int32)(unsafe.Pointer(&dword_5d4594_3798844))) {
			v8 += dword_5d4594_3798796 - dword_5d4594_3798844
		}
		if nox_common_getEngineFlag(nox_engine_flag(NOX_ENGINE_DISABLE_SOFT_LIGHTS)) {
			if uint32(a4*2)+v8 < uint32(*(*int32)(unsafe.Pointer(&dword_5d4594_3798844))) {
				result = (cgoAsFunc(*(*uint32)(unsafe.Pointer(&dword_5d4594_3807160)), (*func(uint32, uint32, uint32) int32)(nil)))(uint32(a5), v8, uint32(a4*2))
			} else {
				v9 = int32(dword_5d4594_3798844 - v8)
				(cgoAsFunc(*(*uint32)(unsafe.Pointer(&dword_5d4594_3807160)), (*func(uint32, uint32, uint32))(nil)))(uint32(a5), v8, dword_5d4594_3798844-v8)
				result = (cgoAsFunc(*(*uint32)(unsafe.Pointer(&dword_5d4594_3807160)), (*func(uint32, uint32, uint32) int32)(nil)))(uint32(a5+v9), dword_5d4594_3798796, uint32(a4*2-v9))
			}
		} else {
			v23 = unsafe.Pointer(uintptr(v8))
			v10 = int32(uint32(v5-*a2*23) - dword_5d4594_2650676)
			v11 = a3*12 + *a2*280
			v12 = (*int32)(unsafe.Pointer(&nox_arr_84EB20[v11+4]))
			v13 = int32(*memmap.PtrUint32(0x8529A0, uint32(((*((*int32)(unsafe.Pointer(&nox_arr_84EB20[v11+284])))-*((*int32)(unsafe.Pointer(&nox_arr_84EB20[v11+4]))))>>8)*4+1020)))
			v22[0] = int32(*memmap.PtrUint32(0x8529A0, uint32(((*((*int32)(unsafe.Pointer(&nox_arr_84EB20[v11+284])))-*((*int32)(unsafe.Pointer(&nox_arr_84EB20[v11+4]))))>>8)*4+1020)))
			v14 = int32(*memmap.PtrUint32(0x8529A0, uint32(((*((*int32)(unsafe.Pointer(&nox_arr_84EB20[v11+288])))-*((*int32)(unsafe.Pointer(&nox_arr_84EB20[v11+8]))))>>8)*4+1020)))
			v22[1] = int32(*memmap.PtrUint32(0x8529A0, uint32(((*((*int32)(unsafe.Pointer(&nox_arr_84EB20[v11+288])))-*((*int32)(unsafe.Pointer(&nox_arr_84EB20[v11+8]))))>>8)*4+1020)))
			v15 = int32(*memmap.PtrUint32(0x8529A0, uint32(((*((*int32)(unsafe.Pointer(&nox_arr_84EB20[v11+292])))-*((*int32)(unsafe.Pointer(&nox_arr_84EB20[v11+12]))))>>8)*4+1020)))
			v22[2] = v15
			v21[0] = *v12
			v16 = *(*int32)(unsafe.Add(unsafe.Pointer(v12), unsafe.Sizeof(int32(0))*1))
			v21[1] = *(*int32)(unsafe.Add(unsafe.Pointer(v12), unsafe.Sizeof(int32(0))*1))
			v21[2] = *(*int32)(unsafe.Add(unsafe.Pointer(v12), unsafe.Sizeof(int32(0))*2))
			if v10 <= 0 {
				v17 = 23
			} else {
				v21[2] += v10 * v15
				v21[0] += v10 * v13
				v21[1] = v10*v14 + v16
				v17 = 23 - v10
			}
			if v17 > a4 {
				v17 = a4
			}
			v18 = v23
			for {
				sub_480290(v18, a5, v17, (*uint32)(unsafe.Pointer(&v21[0])), (*uint32)(unsafe.Pointer(&v22[0])))
				v18 = unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(v18)), v17*2)))
				a5 += v17 * 2
				if uint32(uintptr(v18)) >= uint32(*(*int32)(unsafe.Pointer(&dword_5d4594_3798844))) {
					v18 = unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(v18)), -(((*(*int32)(unsafe.Pointer(&dword_5d4594_3798844)) - *(*int32)(unsafe.Pointer(&dword_5d4594_3798796))) >> 1) * 2))))
				}
				result = a4 - v17
				a4 -= v17
				if a4 <= 0 {
					break
				}
				v17 = 23
				if result < 23 {
					v17 = result
				}
				v19 = *a2 + 1
				*a2 = v19
				v20 = a3*12 + v19*280
				v21[0] = int32(*((*uint32)(unsafe.Pointer(&nox_arr_84EB20[v20+4]))))
				v21[1] = int32(*((*uint32)(unsafe.Pointer(&nox_arr_84EB20[v20+8]))))
				v21[2] = int32(*((*uint32)(unsafe.Pointer(&nox_arr_84EB20[v20+12]))))
				v22[0] = int32(*memmap.PtrUint32(0x8529A0, uint32(((*((*int32)(unsafe.Pointer(&nox_arr_84EB20[v20+284])))-*((*int32)(unsafe.Pointer(&nox_arr_84EB20[v20+4]))))>>8)*4+1020)))
				v22[1] = int32(*memmap.PtrUint32(0x8529A0, uint32(((*((*int32)(unsafe.Pointer(&nox_arr_84EB20[v20+288])))-*((*int32)(unsafe.Pointer(&nox_arr_84EB20[v20+8]))))>>8)*4+1020)))
				v22[2] = int32(*memmap.PtrUint32(0x8529A0, uint32(((*((*int32)(unsafe.Pointer(&nox_arr_84EB20[v20+292])))-*((*int32)(unsafe.Pointer(&nox_arr_84EB20[v20+12]))))>>8)*4+1020)))
			}
		}
	}
	return result
}
func sub_480BE0(a1 *uint32, a2 *int32, a3 int32, a4 int32, a5 int32) {
	var (
		v6  *uint8
		v7  int32
		v9  int32
		v10 int32
		v16 int32
		v17 int32
		v18 int32
		v21 int32
		v22 int32
		v23 int32
		v24 int32
		v25 int32
		v26 int32
		x1  *uint8 = *(**uint8)(unsafe.Pointer(&dword_5d4594_3798844))
		x2  *uint8 = *(**uint8)(unsafe.Pointer(&dword_5d4594_3798796))
	)
	v6 = (*uint8)(unsafe.Pointer(uintptr(dword_5d4594_3798804*(dword_5d4594_3798840+*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*1))-dword_5d4594_3798824) + dword_5d4594_3798796 + (dword_5d4594_3798836+*a1-dword_5d4594_3798820)*2)))
	if uintptr(unsafe.Pointer(v6)) >= uintptr(unsafe.Pointer(x1)) {
		v6 = (*uint8)(unsafe.Add(unsafe.Pointer(v6), int32(uintptr(unsafe.Pointer(x2))-uintptr(unsafe.Pointer(x1)))))
	}
	if nox_common_getEngineFlag(nox_engine_flag(NOX_ENGINE_DISABLE_SOFT_LIGHTS)) {
		if uintptr(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v6), a4*2)))))))) < uintptr(unsafe.Pointer(x1)) {
			(cgoAsFunc(*(*uint32)(unsafe.Pointer(&dword_5d4594_3807160)), (*func(uint32, uint32, uint32) int32)(nil)))(uint32(a5), uint32(uintptr(unsafe.Pointer(v6))), uint32(a4*2))
		} else {
			v7 = int32(uintptr(unsafe.Pointer(x1)) - uintptr(unsafe.Pointer(v6)))
			(cgoAsFunc(*(*uint32)(unsafe.Pointer(&dword_5d4594_3807160)), (*func(uint32, uint32, uint32))(nil)))(uint32(a5), uint32(uintptr(unsafe.Pointer(v6))), uint32(int32(uintptr(unsafe.Pointer(x1))-uintptr(unsafe.Pointer(v6)))))
			(cgoAsFunc(*(*uint32)(unsafe.Pointer(&dword_5d4594_3807160)), (*func(uint32, uint32, uint32) int32)(nil)))(uint32(a5+v7), dword_5d4594_3798796, uint32(a4*2-v7))
		}
	} else {
		v9 = int32(*a1 - uint32(*a2*23) - dword_5d4594_2650676)
		v21 = a3 * 12
		v10 = a3*12 + *a2*280
		var v11 int32 = *((*int32)(unsafe.Pointer(&nox_arr_84EB20[v10+4])))
		var v12 int32 = *((*int32)(unsafe.Pointer(&nox_arr_84EB20[v10+8])))
		var v13 int32 = *((*int32)(unsafe.Pointer(&nox_arr_84EB20[v10+12])))
		var v11b int32 = *((*int32)(unsafe.Pointer(&nox_arr_84EB20[v10+284])))
		var v12b int32 = *((*int32)(unsafe.Pointer(&nox_arr_84EB20[v10+288])))
		var v13b int32 = *((*int32)(unsafe.Pointer(&nox_arr_84EB20[v10+292])))
		v22 = int32(*memmap.PtrUint32(0x8529A0, uint32(((v11b-v11)>>8)*4+1020)))
		v23 = int32(*memmap.PtrUint32(0x8529A0, uint32(((v12b-v12)>>8)*4+1020)))
		v24 = int32(*memmap.PtrUint32(0x8529A0, uint32(((v13b-v13)>>8)*4+1020)))
		if v9 <= 0 {
			v25 = 23
		} else {
			v11 += v9 * v22
			v12 += v9 * v23
			v13 += v9 * v24
			v25 = 23 - v9
		}
		if v25 > a4 {
			v25 = a4
		}
		for {
			for v26 = v25; v26 != 0; v26-- {
				var (
					r uint8
					g uint8
					b uint8
					x uint16 = *(*uint16)(unsafe.Pointer(v6))
				)
				r = uint8(int8((v13 * (int32(x) & 31)) >> 13))
				g = uint8(int8((v12 * ((int32(x) >> 2) & 248)) >> 16))
				b = uint8(int8((v11 * ((int32(x) >> 7) & 248)) >> 16))
				*(*uint16)(unsafe.Pointer(uintptr(a5))) = uint16(int16((int32(r) >> 3) | (int32(g)&248)<<2 | (int32(b)&248)<<7))
				v6 = (*uint8)(unsafe.Add(unsafe.Pointer(v6), 2))
				a5 += 2
				if uintptr(unsafe.Pointer(v6)) >= uintptr(unsafe.Pointer(x1)) {
					v6 = x2
				}
				v11 += v22
				v12 += v23
				v13 += v24
			}
			v16 = a4 - v25
			a4 = v16
			if v16 <= 0 {
				break
			}
			v25 = 23
			if v16 < 23 {
				v25 = v16
			}
			v17 = *a2 + 1
			*a2 = v17
			v18 = v21 + v17*280
			v11 = *((*int32)(unsafe.Pointer(&nox_arr_84EB20[v18+4])))
			v12 = *((*int32)(unsafe.Pointer(&nox_arr_84EB20[v18+8])))
			v13 = *((*int32)(unsafe.Pointer(&nox_arr_84EB20[v18+12])))
			v11b = *((*int32)(unsafe.Pointer(&nox_arr_84EB20[v18+284])))
			v12b = *((*int32)(unsafe.Pointer(&nox_arr_84EB20[v18+288])))
			v13b = *((*int32)(unsafe.Pointer(&nox_arr_84EB20[v18+292])))
			v22 = int32(*memmap.PtrUint32(0x8529A0, uint32(((v11b-v11)>>8)*4+1020)))
			v23 = int32(*memmap.PtrUint32(0x8529A0, uint32(((v12b-v12)>>8)*4+1020)))
			v24 = int32(*memmap.PtrUint32(0x8529A0, uint32(((v13b-v13)>>8)*4+1020)))
		}
	}
}
func sub_480EB0(a1 *uint16, a2 *int16, a3 int32) *uint16 {
	var (
		v3 *int16
		v4 *uint16
		v5 int32
		v6 int16
	)
	v3 = a2
	v4 = a1
	v5 = a3
	for {
		v6 = *v3
		v3 = (*int16)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(int16(0))*1))
		*v4 = uint16(int16(int32(v6)&31 | ((int32(v6) * 2) & 0xFFC0)))
		v4 = (*uint16)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint16(0))*1))
		if func() int32 {
			p := &v5
			x := *p
			*p--
			return x
		}() <= 1 {
			break
		}
	}
	return v4
}
func nox_xxx_someEdgeProcessing_480EF0(a1 int32, a2 int32, a3 int32, a4 *int32, a5 *int32, a6 int32, a7 int32, a8 int32, a9 int32, a10 int32) *int32 {
	var (
		result *int32
		v10    int32
		v11    int8
		v12    int32
		v13    int32
		v14    *int32
		v15    int32
		v16    *byte
		v17    int32
		v18    int32
		v19    int32
		v20    int32
		v21    int32
		v22    int32
		v23    int32
		v24    int32
		v25    int32
		v26    int32
		v27    int32
		i      int32
		v29    int8
		v30    uint32
		v31    unsafe.Pointer
		v32    int32
		j      int32
		v34    int8
		v35    int8
		v36    int32
		v37    int32
		v38    uint8
		v39    int8
		v40    int32
		v41    int32
		v42    func(*byte, *byte, int32, *int32, *int32)
		v43    int32
		v44    int32
		v45    int2
		v46    int2
		v47    int32
		v48    int32
		v49    int32
		v50    [3]int32
		v51    [3]int32
		v52    *byte
		v53    int32
		v54    int32
		v55    int32
		v56    int32
		v57    uint8
		v58    uint8
		v59    uint8
		v60    uint8
		v61    uint8
	)
	result = (*int32)(unsafe.Pointer(uintptr(a1)))
	v10 = 0
	v44 = 0
	if result == nil {
		return nil
	}
	v11 = int8(nox_video_bag_image_type((*nox_video_bag_image_t)(unsafe.Pointer(uintptr(a1)))))
	v45.field_4 = 0
	v45.field_0 = 0
	v46.field_4 = 0
	v46.field_0 = 0
	if (int32(v11) & 63) != 3 {
		return result
	}
	v42 = func(arg1 *byte, arg2 *byte, arg3 int32, arg4 *int32, arg5 *int32) {
		sub_480860((*uint16)(unsafe.Pointer(arg1)), (*uint16)(unsafe.Pointer(arg2)), arg3, arg4, arg5)
	}
	result = (*int32)(nox_video_getImagePixdata_42FB30((*nox_video_bag_image_t)(unsafe.Pointer(uintptr(a1)))))
	if result == nil {
		return result
	}
	v12 = *result
	v13 = *(*int32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(int32(0))*1)) - a9
	v43 = *result
	v40 = *result
	v41 = v13
	if v13 <= 0 {
		return result
	}
	v14 = (*int32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(int32(0))*3))
	result = (*int32)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(int32(0))*2)) + a2)))
	v15 = *v14
	v16 = (*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(v14))), 5))
	v17 = v15 + a3
	v53 = v15 + a3
	if !(int32(uintptr(unsafe.Pointer(result))) <= *(*int32)(unsafe.Pointer(&dword_5d4594_3807116)) && v17 <= *(*int32)(unsafe.Pointer(&dword_5d4594_3807152))) {
		return result
	}
	if int32(uintptr(unsafe.Pointer(result))) < *(*int32)(unsafe.Pointer(&dword_5d4594_3807140)) {
		if int32(uintptr(unsafe.Pointer(result)))+v12 < *(*int32)(unsafe.Pointer(&dword_5d4594_3807140)) {
			return result
		}
		v10 = *(*int32)(unsafe.Pointer(&dword_5d4594_3807140)) - int32(uintptr(unsafe.Pointer(result)))
		v40 = v12 - (*(*int32)(unsafe.Pointer(&dword_5d4594_3807140)) - int32(uintptr(unsafe.Pointer(result))))
		result = *(**int32)(unsafe.Pointer(&dword_5d4594_3807140))
	}
	if int32(uintptr(unsafe.Pointer(result)))+v40 > *(*int32)(unsafe.Pointer(&dword_5d4594_3807116)) {
		v40 = *(*int32)(unsafe.Pointer(&dword_5d4594_3807116)) - int32(uintptr(unsafe.Pointer(result)))
	}
	if v17 < *(*int32)(unsafe.Pointer(&dword_5d4594_3807136)) {
		if v13+v17 < *(*int32)(unsafe.Pointer(&dword_5d4594_3807136)) {
			return result
		}
		v18 = *(*int32)(unsafe.Pointer(&dword_5d4594_3807136)) - v17
		v17 = *(*int32)(unsafe.Pointer(&dword_5d4594_3807136))
		v13 -= v18
		v53 = *(*int32)(unsafe.Pointer(&dword_5d4594_3807136))
		v41 = v13
		v44 = v18
	}
	if v13+v17 > *(*int32)(unsafe.Pointer(&dword_5d4594_3807152)) {
		v13 = *(*int32)(unsafe.Pointer(&dword_5d4594_3807152)) - v17
		v41 = *(*int32)(unsafe.Pointer(&dword_5d4594_3807152)) - v17
	}
	if !(v40 > 0 && v13 > 0) {
		return result
	}
	v19 = a6
	if a6 < v17 {
		v19 = v17
	}
	if v13+v17 > v19 {
		v41 = v19 - v17
	}
	if a7 <= int32(uintptr(unsafe.Pointer(result)))+v10 {
		v20 = 0
	} else {
		v20 = a7 - v10 - int32(uintptr(unsafe.Pointer(result)))
		v10 = a7 - int32(uintptr(unsafe.Pointer(result)))
		v40 -= v20
	}
	if a8 < int32(uintptr(unsafe.Pointer(result)))+v40+v10 {
		v40 = a8 - v10 - int32(uintptr(unsafe.Pointer(result)))
	}
	if v40 <= 0 {
		return result
	}
	v21 = int32(uintptr(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(nox_pixbuffer_rows_3798784), unsafe.Sizeof((*uint8)(nil))*uintptr(v17)))))) + (int32(uintptr(unsafe.Pointer(result)))+v20)*2
	v22 = *(*int32)(unsafe.Add(unsafe.Pointer(a4), unsafe.Sizeof(int32(0))*1))
	v52 = (*byte)(unsafe.Pointer(uintptr(v21)))
	v23 = *a4
	v49 = *(*int32)(unsafe.Add(unsafe.Pointer(a4), unsafe.Sizeof(int32(0))*2)) << 8
	v23 <<= 8
	v24 = (*a5 << 8) - v23
	v47 = v23
	v48 = v22 << 8
	v25 = *(*int32)(unsafe.Add(unsafe.Pointer(a5), unsafe.Sizeof(int32(0))*1)) << 8
	v26 = *(*int32)(unsafe.Add(unsafe.Pointer(a5), unsafe.Sizeof(int32(0))*2)) << 8
	v50[0] = v24 / v43
	v50[1] = (v25 - v48) / v43
	v50[2] = (v26 - v49) / v43
	if v44 > 0 {
		v27 = v44
		for {
			for i = v43; i > 0; i -= int32(v57) {
				v29 = int8(*v16)
				v57 = uint8(*(*byte)(unsafe.Add(unsafe.Pointer(v16), 1)))
				v16 = (*byte)(unsafe.Add(unsafe.Pointer(v16), 2))
				*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a1))), 0)) = uint8(v29)
				if int32(v29) == 2 {
					v16 = (*byte)(unsafe.Add(unsafe.Pointer(v16), int32(v57)*2))
				}
			}
			v27--
			if v27 == 0 {
				break
			}
		}
	}
	v46.field_4 = v53
	sub_473970(&v46, &v45)
	v30 = 0
	v31 = nil
	result = (*int32)(unsafe.Pointer(uintptr(v41 - 1)))
	if v41 == 0 {
		return result
	}
	v56 = v41
	var bpitch int32 = nox_getBackbufferPitch()
	for {
		v32 = 0
		v54 = 0
		v55 = 0
		if nox_client_highResFrontWalls_80820 != 0 || (v45.field_4&1) == 0 {
			if v10 <= 0 {
				v35 = int8(a1)
			} else {
				for {
					v35 = int8(*v16)
					v59 = uint8(*(*byte)(unsafe.Add(unsafe.Pointer(v16), 1)))
					v16 = (*byte)(unsafe.Add(unsafe.Pointer(v16), 2))
					*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a1))), 0)) = uint8(v35)
					if int32(v35) == 2 {
						v16 = (*byte)(unsafe.Add(unsafe.Pointer(v16), int32(v59)*2))
					}
					v32 += int32(v59)
					if v32 >= v10 {
						break
					}
				}
			}
			if v32 <= v10 {
				v36 = v40
			} else {
				v36 = v40
				v37 = v40
				if v32-v10 <= v40 {
					v37 = v32 - v10
				}
				if int32(v35) == 2 {
					v51[0] = v47 + v10*v50[0]
					v51[1] = v48 + v10*v50[1]
					v51[2] = v49 + v10*v50[2]
					v42(v52, (*byte)(unsafe.Add(unsafe.Pointer(v16), (v32-v10)*(-2))), v37, &v51[0], &v50[0])
					v36 = v40
					v54 = v37 * 2
				}
				v55 = 0
			}
			if v32 < v36+v10 {
				for {
					v38 = uint8(*(*byte)(unsafe.Add(unsafe.Pointer(v16), 1)))
					*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a1))), 0)) = uint8(*v16)
					v16 = (*byte)(unsafe.Add(unsafe.Pointer(v16), 2))
					v60 = v38
					if int32(uint8(int8(a1))) == 2 {
						if int32(v38) > v10+v36-v32 {
							v55 = v32 + int32(v38) - v36 - v10
							v60 = uint8(int8(int32(-int8(v32 + int32(-int8(v36)) - v10))))
						}
						v51[0] = v47 + v50[0]*v32
						v51[2] = v49 + v50[2]*v32
						v51[1] = v48 + v50[1]*v32
						v42((*byte)(unsafe.Add(unsafe.Pointer(v52), (v32-v10)*2)), v16, int32(v60), &v51[0], &v50[0])
						v36 = v40
						v16 = (*byte)(unsafe.Add(unsafe.Pointer(v16), int32(v60)*2))
					}
					v32 += int32(v60)
					v54 += int32(v60) * 2
					if v32 >= v36+v10 {
						break
					}
				}
				if v55 != 0 {
					if int32(uint8(int8(a1))) == 2 {
						v16 = (*byte)(unsafe.Add(unsafe.Pointer(v16), v55*2))
					}
					v32 += v55
				}
			}
			for ; v32 < v43; v32 += int32(v61) {
				v39 = int8(*v16)
				v61 = uint8(*(*byte)(unsafe.Add(unsafe.Pointer(v16), 1)))
				v16 = (*byte)(unsafe.Add(unsafe.Pointer(v16), 2))
				*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a1))), 0)) = uint8(v39)
				if int32(v39) == 2 {
					v16 = (*byte)(unsafe.Add(unsafe.Pointer(v16), int32(v61)*2))
				}
			}
		} else {
			if v31 != nil && v30 != 0 {
				libc.MemCpy(unsafe.Pointer(v52), v31, int(v30))
			}
			for j = v43; j > 0; j -= int32(v58) {
				v34 = int8(*v16)
				v58 = uint8(*(*byte)(unsafe.Add(unsafe.Pointer(v16), 1)))
				v16 = (*byte)(unsafe.Add(unsafe.Pointer(v16), 2))
				*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a1))), 0)) = uint8(v34)
				if int32(v34) == 2 {
					v16 = (*byte)(unsafe.Add(unsafe.Pointer(v16), int32(v58)*2))
				}
			}
		}
		v31 = unsafe.Pointer(v52)
		v52 = (*byte)(unsafe.Add(unsafe.Pointer(v52), bpitch))
		v45.field_4++
		result = (*int32)(unsafe.Pointer(uintptr(func() int32 {
			p := &v56
			*p--
			return *p
		}())))
		if v56 == 0 {
			break
		}
		v30 = uint32(v54)
	}
	return result
}
func sub_481410() {
	nox_xxx_waypointCounterMB_587000_154948 = math.MaxUint32
}
func nox_xxx_tileSetDrawFn_481420() int32 {
	var result int32
	result = int32(nox_client_texturedFloors_154956)
	if nox_client_texturedFloors_154956 != 0 {
		func_587000_154940 = func(arg1 *int2, arg2 uint32, arg3 uint32) {
			nox_xxx_tileDraw_4815E0((*uint32)(unsafe.Pointer(arg1)), int32(arg2))
		}
		func_587000_154944 = func(arg1 int32, arg2 int32) int32 {
			return int32(nox_xxx_drawTexEdgesProbably_481900((*uint32)(unsafe.Pointer(uintptr(arg1))), (*uint32)(unsafe.Pointer(uintptr(arg2)))))
		}
	} else {
		func_587000_154940 = func(arg1 *int2, arg2 uint32, arg3 uint32) {
			sub_481770((*uint32)(unsafe.Pointer(arg1)), int32(arg2), uint16(arg3))
		}
		func_587000_154944 = nullsub_8
	}
	dword_5d4594_1193188 = 1
	return result
}
func nox_xxx_tileDraw_4815E0(a1 *uint32, a2 int32) int8 {
	var (
		v2  uint32
		v3  int32
		v4  *byte
		v5  uint32
		v6  int32
		v7  int32
		v8  uint32
		v9  int32
		v10 *byte
		v11 *byte
		v12 int8
		v14 *byte
		i   int32
	)
	v2 = dword_5d4594_3798804*(dword_5d4594_3798840+*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*1))-dword_5d4594_3798824) + dword_5d4594_3798796 + ((dword_5d4594_3798836 + *a1 - dword_5d4594_3798820) << uint32(*memmap.PtrUint8(0x973F18, 7696)))
	if v2 >= uint32(*(*int32)(unsafe.Pointer(&dword_5d4594_3798844))) {
		v2 += dword_5d4594_3798796 - dword_5d4594_3798844
	}
	v3 = int32(uintptr(nox_video_getImagePixdata_42FB30((*nox_video_bag_image_t)(unsafe.Pointer(uintptr(a2))))))
	v4 = (*byte)(unsafe.Pointer(uintptr(v3)))
	v14 = (*byte)(unsafe.Pointer(uintptr(v3)))
	if v3 != 0 {
		v5 = dword_5d4594_3798844
		if v2+uint32(*memmap.PtrInt32(0x973CE0, 376)) >= uint32(*(*int32)(unsafe.Pointer(&dword_5d4594_3798844))) {
			v6 = 0
			for i = 0; ; v6 = i {
				v7 = int32(*memmap.PtrUint32(0x973CE0, uint32(v6+192)) << uint32(*memmap.PtrUint8(0x973F18, 7696)))
				v8 = *memmap.PtrUint32(0x973CE0, uint32(v6+384)) << uint32(*memmap.PtrUint8(0x973F18, 7696))
				if v2+uint32(v7)+v8 < v5 {
					libc.MemCpy(unsafe.Pointer(uintptr(uint32(v7)+v2)), unsafe.Pointer(v4), int((v8>>2)*4))
					v11 = (*byte)(unsafe.Add(unsafe.Pointer(v4), (v8>>2)*4))
					v10 = (*byte)(unsafe.Pointer(uintptr(uint32(v7) + v2 + (v8>>2)*4)))
					v12 = int8(uint8(v8))
				} else {
					v9 = int32(v5 - uint32(v7) - v2)
					if v9 > 0 {
						libc.MemCpy(unsafe.Pointer(uintptr(uint32(v7)+v2)), unsafe.Pointer(v4), int(v9))
						libc.MemCpy(*(*unsafe.Pointer)(unsafe.Pointer(&dword_5d4594_3798796)), unsafe.Add(unsafe.Pointer(v14), v9), int(v8-uint32(v9)))
						v5 = dword_5d4594_3798844
						v2 += dword_5d4594_3798796 - dword_5d4594_3798844
						goto LABEL_13
					}
					v2 += dword_5d4594_3798796 - v5
					libc.MemCpy(unsafe.Pointer(uintptr(uint32(v7)+v2)), unsafe.Pointer(v4), int((v8>>2)*4))
					v11 = (*byte)(unsafe.Add(unsafe.Pointer(v4), (v8>>2)*4))
					v10 = (*byte)(unsafe.Pointer(uintptr(uint32(v7) + v2 + (v8>>2)*4)))
					v12 = int8(uint8(v8))
				}
				libc.MemCpy(unsafe.Pointer(v10), unsafe.Pointer(v11), int(int32(v12)&3))
				v5 = dword_5d4594_3798844
			LABEL_13:
				v2 += dword_5d4594_3798804
				v14 = (*byte)(unsafe.Add(unsafe.Pointer(v14), v8))
				*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v3))), 0)) = uint8(int8(i + 4))
				i += 4
				if i >= 184 {
					return int8(v3)
				}
				v4 = v14
			}
		} else {
			*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v3))), 0)) = uint8(int8(sub_4831C0(v3, int32(v2))))
		}
	}
	return int8(v3)
}
func sub_481770(a1 *uint32, a2 int32, a3 uint16) *byte {
	var (
		v3     int32
		v4     uint8
		v5     uint32
		v6     uint32
		v7     int32
		v8     int32
		v9     int32
		v10    int32
		result *byte
		v12    int32
	)
	if nox_client_texturedFloors_154956 == 0 && (int32(*memmap.PtrUint8(0x85B3FC, uint32(int32(a3)*60+0x7F1E)))&1) == 1 {
		dword_5d4594_1193156 = 1
	}
	v3 = int32(*memmap.PtrUint32(0x85B3FC, uint32(int32(a3)*60+0x7F14)))
	v4 = *memmap.PtrUint8(0x973F18, 7696)
	v5 = dword_5d4594_3798844
	v6 = dword_5d4594_3798804*(dword_5d4594_3798840+*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*1))-dword_5d4594_3798824) + dword_5d4594_3798796 + ((dword_5d4594_3798836 + *a1 - dword_5d4594_3798820) << uint32(*memmap.PtrUint8(0x973F18, 7696)))
	if v6 >= uint32(*(*int32)(unsafe.Pointer(&dword_5d4594_3798844))) {
		v6 += dword_5d4594_3798796 - dword_5d4594_3798844
	}
	if v6+*memmap.PtrUint32(0x973CE0, 376) < uint32(*(*int32)(unsafe.Pointer(&dword_5d4594_3798844))) {
		result = (*byte)(unsafe.Pointer(uintptr(sub_484450(v3, int32(v6)))))
	} else {
		v7 = 0
		v12 = 0
		for {
			v8 = int32(*memmap.PtrUint32(0x973CE0, uint32(v7+192)) << uint32(v4))
			v9 = int32(*memmap.PtrUint32(0x973CE0, uint32(v7+384)) << uint32(v4))
			if v6+uint32(v9)+uint32(v8) < v5 {
				result = (*byte)(unsafe.Pointer(uintptr(sub_49D1C0(unsafe.Pointer(uintptr(v6+uint32(v8))), v3, v9))))
				v5 = dword_5d4594_3798844
			} else {
				v10 = int32(v5 - uint32(v8) - v6)
				if v10 <= 0 {
					v6 += dword_5d4594_3798796 - v5
					result = (*byte)(unsafe.Pointer(uintptr(sub_49D1C0(unsafe.Pointer(uintptr(v6+uint32(v8))), v3, v9))))
					v5 = dword_5d4594_3798844
					v7 = v12
				} else {
					sub_49D1C0(unsafe.Pointer(uintptr(v6+uint32(v8))), v3, int32(v5-uint32(v8)-v6))
					sub_49D1C0(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_3798796)))), v3, v9-v10)
					v5 = dword_5d4594_3798844
					v7 = v12
					result = (*byte)(unsafe.Pointer(uintptr(dword_5d4594_3798796 - dword_5d4594_3798844)))
					v6 += dword_5d4594_3798796 - dword_5d4594_3798844
				}
			}
			v7 += 4
			v6 += dword_5d4594_3798804
			v12 = v7
			if v7 >= 184 {
				break
			}
			v4 = *memmap.PtrUint8(0x973F18, 7696)
		}
	}
	return result
}
func nox_xxx_drawTexEdgesProbably_481900(a1 *uint32, a2 *uint32) int8 {
	var (
		v2   int32
		v3   int32
		v4   int32
		v5   int32
		v6   int32
		v7   int32
		v8   int32
		v9   int32
		v10  *uint8
		v11  int32
		v12  uint32
		v13  uint32
		v14  *byte
		v15  *byte
		v16  int32
		v17  *byte
		v18  uint32
		v19  *byte
		v20  uint32
		v22  int32
		v24  int32
		v25  uint8
		v26  *byte
		v27  *byte
		v28  int32
		v29  uint8
		v30  int32
		v31  uint32
		v32  uint8
		v33  *byte
		addr *int32
	)
	v2 = int32(dword_5d4594_3798824)
	v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(*memmap.PtrUint32(0x85B3FC, *a2*60+0x7F04) + (*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*1))+uint32(*memmap.PtrUint16(0x85B3FC, *a2*60+32530)))*4))))
	v4 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*2)))
	v5 = int32(dword_5d4594_3798836)
	addr = (*int32)(unsafe.Pointer(uintptr(*memmap.PtrUint32(0x85B3FC, uint32(v4*60+0x7004)) + (*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*3))+uint32(*memmap.PtrUint16(0x85B3FC, uint32(v4*60+28690))))*4)))
	v6 = int32(*(*uint32)(unsafe.Pointer(addr)))
	v7 = int32(dword_5d4594_3798840)
	*memmap.PtrUint32(6112660, uint32(v4*4)+2523980) = 1
	v8 = int32(dword_5d4594_3798804*(uint32(v7)+*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*1))-uint32(v2)) + dword_5d4594_3798796 + ((uint32(v5) + *a1 - dword_5d4594_3798820) << uint32(*memmap.PtrUint8(0x973F18, 7696))))
	v9 = int32(uintptr(nox_video_getImagePixdata_42FB30((*nox_video_bag_image_t)(unsafe.Pointer(uintptr(v6))))))
	if v9 != 0 {
		v32 = *(*uint8)(unsafe.Pointer(uintptr(v9)))
		v25 = *(*uint8)(unsafe.Pointer(uintptr(v9 + 1)))
		v10 = (*uint8)(unsafe.Pointer(uintptr(v9 + 2)))
		v9 = int32(uintptr(nox_video_getImagePixdata_42FB30((*nox_video_bag_image_t)(unsafe.Pointer(uintptr(v3))))))
		v11 = v9
		if v9 != 0 {
			v12 = dword_5d4594_3798844
			v9 = int32(v32)
			v13 = dword_5d4594_3798804*uint32(v32) + uint32(v8)
			v31 = v13
			v14 = (*byte)(unsafe.Pointer(uintptr((*memmap.PtrUint32(0x973CE0, uint32(int32(v32)*+4)) << uint32(*memmap.PtrUint8(0x973F18, 7696))) + uint32(v11))))
			v27 = v14
			if v13 >= uint32(*(*int32)(unsafe.Pointer(&dword_5d4594_3798844))) {
				v13 += dword_5d4594_3798796 - dword_5d4594_3798844
				v31 = v13
			}
			v28 = int32(v32)
			v30 = int32(v25)
			if int32(v32) <= int32(v25) {
				for {
					v24 = int32(*memmap.PtrUint32(0x973CE0, uint32(v9*4+384)))
					v26 = v14
					v15 = (*byte)(unsafe.Pointer(uintptr(v13 + (*memmap.PtrUint32(0x973CE0, uint32(v9*4+192)) << uint32(*memmap.PtrUint8(0x973F18, 7696))))))
					if v24 > 0 {
						for {
							*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v9))), 0)) = *v10
							v29 = *(*uint8)(unsafe.Add(unsafe.Pointer(v10), 1))
							v33 = (*byte)(unsafe.Add(unsafe.Pointer(v10), 2))
							v16 = int32(v29) << int32(*memmap.PtrUint8(0x973F18, 7696))
							switch uint8(int8(v9)) {
							case 2:
								if uint32(uintptr(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v15), v16))))) < v12 {
									v19 = v33
									v18 = uint32(int32(v29) << int32(*memmap.PtrUint8(0x973F18, 7696)))
									v17 = v15
								} else {
									libc.MemCpy(unsafe.Pointer(v15), unsafe.Pointer(v33), int(v12-uint32(uintptr(unsafe.Pointer(v15)))))
									v17 = *(**byte)(unsafe.Pointer(&dword_5d4594_3798796))
									v18 = uint32(v16) - (v12 - uint32(uintptr(unsafe.Pointer(v15))))
									v19 = (*byte)(unsafe.Add(unsafe.Pointer(v33), v12-uint32(uintptr(unsafe.Pointer(v15)))))
								}
								libc.MemCpy(unsafe.Pointer(v17), unsafe.Pointer(v19), int(v18))
								v12 = dword_5d4594_3798844
								v13 = v31
								v33 = (*byte)(unsafe.Add(unsafe.Pointer(v33), v16))
							case 3:
								if uint32(uintptr(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v15), v16))))) < v12 {
									libc.MemCpy(unsafe.Pointer(v15), unsafe.Pointer(v14), int(v16))
								} else {
									v20 = v12 - uint32(uintptr(unsafe.Pointer(v15)))
									libc.MemCpy(unsafe.Pointer(v15), unsafe.Pointer(v14), int(v20))
									v14 = v26
									libc.MemCpy(*(*unsafe.Pointer)(unsafe.Pointer(&dword_5d4594_3798796)), unsafe.Add(unsafe.Pointer(v26), v20), int(uint32(v16)-v20))
								}
								v12 = dword_5d4594_3798844
								v13 = v31
							case 1:
							default:
								return int8(v9)
							}
							v15 = (*byte)(unsafe.Add(unsafe.Pointer(v15), v16))
							v22 = v24 - int32(v29)
							v14 = (*byte)(unsafe.Add(unsafe.Pointer(v14), v16))
							v24 -= int32(v29)
							v26 = v14
							if uint32(uintptr(unsafe.Pointer(v15))) >= v12 {
								v15 = (*byte)(unsafe.Add(unsafe.Pointer(v15), dword_5d4594_3798796-v12))
							}
							v10 = (*uint8)(unsafe.Pointer(v33))
							if v22 <= 0 {
								break
							}
						}
						v9 = v28
					}
					v14 = (*byte)(unsafe.Add(unsafe.Pointer(v27), *memmap.PtrUint32(0x973CE0, uint32(v9*4+384))<<uint32(*memmap.PtrUint8(0x973F18, 7696))))
					v13 += dword_5d4594_3798804
					v27 = (*byte)(unsafe.Add(unsafe.Pointer(v27), *memmap.PtrUint32(0x973CE0, uint32(v9*4+384))<<uint32(*memmap.PtrUint8(0x973F18, 7696))))
					v31 = v13
					if v13 >= v12 {
						v13 += dword_5d4594_3798796 - v12
						v31 = v13
					}
					v28 = func() int32 {
						p := &v9
						*p++
						return *p
					}()
					if v9 > v30 {
						break
					}
				}
			}
		}
	}
	return int8(v9)
}
func nox_xxx_tileCallDrawEdges_481BF0(a1 int32, a2 int32) {
	var i int32
	for i = a2; i != 0; i = int32(*(*uint32)(unsafe.Pointer(uintptr(i + 16)))) {
		func_587000_154944(a1, i)
	}
}
func nox_xxx_tileDrawMB_481C20_A(vp *nox_draw_viewport_t, v3 int32) {
	var (
		v17 int32
		v63 int32
		v16 int32
		v62 int32
		v14 int32
		v15 int32
		v13 int32
		v12 int32
		v11 int32
		v4  int32
		v6  int32
		v10 int32
		v9  int32
		v7  int32
		v8  int32
		j   int32
		v68 int2
	)
	if v3 >= *(*int32)(unsafe.Pointer(&dword_5d4594_3798820))+23 {
		var v71 int32 = nox_getBackbufWidth() + v3
		if uint32(v71) <= uint32(*(*int32)(unsafe.Pointer(&dword_5d4594_3798800)))+dword_5d4594_3798820-46 || *(*int32)(unsafe.Pointer(&dword_5d4594_3798812))+*(*int32)(unsafe.Pointer(&dword_5d4594_3798828))-1 >= 128 {
			return
		}
		if v71 > *(*int32)(unsafe.Pointer(&dword_5d4594_3798800))+*(*int32)(unsafe.Pointer(&dword_5d4594_3798820)) {
			nox_xxx_tileDrawImpl_4826A0(vp)
			return
		}
		v7 = int32(dword_5d4594_3798828 + 1)
		dword_5d4594_3798828 = uint32(v7)
		v8 = int32(dword_5d4594_3798820 + 46)
		dword_5d4594_3798820 += 46
		j = int32(dword_5d4594_3798812 + uint32(v7) - 2)
		v9 = int32(dword_5d4594_3798836 + 46)
		dword_5d4594_3798836 += 46
		if *(*int32)(unsafe.Pointer(&dword_5d4594_3798836)) >= *(*int32)(unsafe.Pointer(&dword_5d4594_3798800)) {
			dword_5d4594_3798836 = uint32(v9) - dword_5d4594_3798800
			v10 = int32(func() uint32 {
				p := &dword_5d4594_3798840
				*p++
				return *p
			}())
			if *(*int32)(unsafe.Pointer(&dword_5d4594_3798840)) >= *(*int32)(unsafe.Pointer(&dword_5d4594_3798808)) {
				dword_5d4594_3798840 = uint32(v10) - dword_5d4594_3798808
			}
		}
		v4 = int32(dword_5d4594_3798800 + uint32(v8) - 92)
	} else {
		if *(*int32)(unsafe.Pointer(&dword_5d4594_3798828)) <= 0 {
			return
		}
		if v3 < *(*int32)(unsafe.Pointer(&dword_5d4594_3798820))-23 {
			nox_xxx_tileDrawImpl_4826A0(vp)
			return
		}
		v4 = int32(dword_5d4594_3798820 - 46)
		j = int32(func() uint32 {
			p := &dword_5d4594_3798828
			*p--
			return *p
		}())
		dword_5d4594_3798820 -= 46
		dword_5d4594_3798836 -= 46
		if *(*int32)(unsafe.Pointer(&dword_5d4594_3798836)) < 0 {
			v6 = *(*int32)(unsafe.Pointer(&dword_5d4594_3798840)) - 1
			var v5 bool = *(*int32)(unsafe.Pointer(&dword_5d4594_3798840))-1 < 0
			dword_5d4594_3798836 += dword_5d4594_3798800
			*(*int32)(unsafe.Pointer(&dword_5d4594_3798840))--
			if v5 {
				*(*int32)(unsafe.Pointer(&dword_5d4594_3798840)) = *(*int32)(unsafe.Pointer(&dword_5d4594_3798808)) + v6
			}
		}
	}
	var v76 int32 = v4
	sub_481410()
	v11 = int32(dword_5d4594_3798824)
	var v74 uint32 = dword_5d4594_3798832
	if *(*int32)(unsafe.Pointer(&dword_5d4594_3798832)) < *(*int32)(unsafe.Pointer(&dword_5d4594_3798832))+*(*int32)(unsafe.Pointer(&dword_5d4594_3798816)) {
		v12 = int32(dword_5d4594_3798832 * 44)
		for i := int32(int32(dword_5d4594_3798832 * 44)); ; v12 = i {
			*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v13))), unsafe.Sizeof(uint16(0))*1)) = *(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&ptr_5D4594_2650668))), unsafe.Sizeof(uint16(0))*1))
			v14 = int32(uint32(v12) + uint32(uintptr(unsafe.Pointer(*(**obj_5D4594_2650668_t)(unsafe.Add(unsafe.Pointer(ptr_5D4594_2650668), unsafe.Sizeof((*obj_5D4594_2650668_t)(nil))*uintptr(j)))))))
			if int32(*(*uint8)(unsafe.Pointer(uintptr(v14))))&2 != 0 {
				*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v13))), unsafe.Sizeof(uint16(0))*0)) = *(*uint16)(unsafe.Pointer(uintptr(v14 + 24)))
				v15 = int32(*(*uint16)(unsafe.Pointer(uintptr(v14 + 24))))
				v62 = int32(*(*uint32)(unsafe.Pointer(uintptr(*memmap.PtrUint32(0x85B3FC, uint32(v15*60+0x7F04)) + (*(*uint32)(unsafe.Pointer(uintptr(v14 + 28)))+uint32(*memmap.PtrUint16(0x85B3FC, uint32(v15*60+32530))))*4))))
				v68.field_0 = v76
				v68.field_4 = v11 + 23
				func_587000_154940(&v68, uint32(v62), uint32(v13))
				*memmap.PtrUint32(0x85B3FC, uint32(v15*4+228)) = 1
				if *(*uint32)(unsafe.Pointer(uintptr(v14 + 40))) != 0 {
					nox_xxx_tileCallDrawEdges_481BF0(int32(uintptr(unsafe.Pointer(&v68))), int32(*(*uint32)(unsafe.Pointer(uintptr(v14 + 40)))))
				}
			}
			if int32(*(*uint8)(unsafe.Pointer(uintptr(v14))))&1 != 0 {
				*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v13))), unsafe.Sizeof(uint16(0))*0)) = *(*uint16)(unsafe.Pointer(uintptr(v14 + 4)))
				v16 = int32(*(*uint16)(unsafe.Pointer(uintptr(v14 + 4))))
				v63 = int32(*(*uint32)(unsafe.Pointer(uintptr(*memmap.PtrUint32(0x85B3FC, uint32(v16*60+0x7F04)) + (*(*uint32)(unsafe.Pointer(uintptr(v14 + 8)))+uint32(*memmap.PtrUint16(0x85B3FC, uint32(v16*60+32530))))*4))))
				v68.field_0 = v76 + 23
				v68.field_4 = v11
				func_587000_154940(&v68, uint32(v63), uint32(v13))
				*memmap.PtrUint32(0x85B3FC, uint32(v16*4+228)) = 1
				v17 = int32(*(*uint32)(unsafe.Pointer(uintptr(v14 + 20))))
				if v17 != 0 {
					nox_xxx_tileCallDrawEdges_481BF0(int32(uintptr(unsafe.Pointer(&v68))), v17)
				}
			}
			v11 += 46
			var v18 bool = uint32(int32(func() uint32 {
				p := &v74
				*p++
				return *p
			}())) < uint32(*(*int32)(unsafe.Pointer(&dword_5d4594_3798832)))+dword_5d4594_3798816
			i += 44
			if !v18 {
				break
			}
		}
	}
}
func nox_xxx_tileDrawMB_481C20_B(vp *nox_draw_viewport_t, v78 int32) {
	var (
		v33 int32
		v65 int32
		v32 int32
		v68 int2
		v64 int32
		v31 int32
		v28 int32
		v29 int8
		v30 int32
		v27 int32
		v25 int32
		v26 int32
		v21 int32
		v20 int32
		v19 int32
		v23 int32
		v24 int32
		v22 int32
		v76 int32
	)
	if v78 >= *(*int32)(unsafe.Pointer(&dword_5d4594_3798824))+23 {
		if v78+nox_getBackbufHeight() <= *(*int32)(unsafe.Pointer(&dword_5d4594_3798824))+*(*int32)(unsafe.Pointer(&dword_5d4594_3798808)) {
			return
		}
		v22 = int32(dword_5d4594_3798832)
		if *(*int32)(unsafe.Pointer(&dword_5d4594_3798832))+*(*int32)(unsafe.Pointer(&dword_5d4594_3798816)) >= 128 {
			return
		}
		if v78+nox_getBackbufHeight() > *(*int32)(unsafe.Pointer(&dword_5d4594_3798824))+*(*int32)(unsafe.Pointer(&dword_5d4594_3798808))+46 {
			nox_xxx_tileDrawImpl_4826A0(vp)
			return
		}
		dword_5d4594_3798832++
		v23 = int32(dword_5d4594_3798824 + 46)
		v19 = int32(dword_5d4594_3798816 + uint32(v22))
		v24 = int32(dword_5d4594_3798840 + 46)
		dword_5d4594_3798824 += 46
		dword_5d4594_3798840 += 46
		if *(*int32)(unsafe.Pointer(&dword_5d4594_3798840)) >= *(*int32)(unsafe.Pointer(&dword_5d4594_3798808)) {
			dword_5d4594_3798840 = uint32(v24) - dword_5d4594_3798808
		}
		v76 = int32(uint32(v23) + dword_5d4594_3798808 - 46)
	} else {
		if *(*int32)(unsafe.Pointer(&dword_5d4594_3798832)) <= 0 {
			return
		}
		if v78 < *(*int32)(unsafe.Pointer(&dword_5d4594_3798824))-23 {
			nox_xxx_tileDrawImpl_4826A0(vp)
			return
		}
		v19 = int32(dword_5d4594_3798832 - 1)
		v20 = int32(dword_5d4594_3798824 - 46)
		v21 = int32(dword_5d4594_3798840 - 46)
		var v5 bool = *(*int32)(unsafe.Pointer(&dword_5d4594_3798840))-46 < 0
		dword_5d4594_3798832--
		dword_5d4594_3798824 -= 46
		dword_5d4594_3798840 -= 46
		if v5 {
			dword_5d4594_3798840 = dword_5d4594_3798808 + uint32(v21)
		}
		v76 = v20
	}
	sub_481410()
	v25 = int32(dword_5d4594_3798828)
	v26 = int32(dword_5d4594_3798820)
	var i int32 = int32(dword_5d4594_3798828)
	var j int32 = int32(dword_5d4594_3798812 + dword_5d4594_3798828 - 1)
	if dword_5d4594_3798828 < uint32(j) {
		v27 = v19 * 44
		var v71 int32 = v19 * 44
		for {
			v28 = int32(uintptr(unsafe.Pointer(*(**obj_5D4594_2650668_t)(unsafe.Add(unsafe.Pointer(ptr_5D4594_2650668), unsafe.Sizeof((*obj_5D4594_2650668_t)(nil))*uintptr(v25))))))
			v29 = int8(*(*uint8)(unsafe.Pointer(uintptr(v28 + v27))))
			v30 = v27 + v28
			if int32(v29)&2 != 0 {
				*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v27))), unsafe.Sizeof(uint16(0))*0)) = *(*uint16)(unsafe.Pointer(uintptr(v30 + 24)))
				v31 = int32(*(*uint16)(unsafe.Pointer(uintptr(v30 + 24))))
				v64 = int32(*(*uint32)(unsafe.Pointer(uintptr(*memmap.PtrUint32(0x85B3FC, uint32(v31*60+0x7F04)) + (*(*uint32)(unsafe.Pointer(uintptr(v30 + 28)))+uint32(*memmap.PtrUint16(0x85B3FC, uint32(v31*60+32530))))*4))))
				v68.field_0 = v26
				v68.field_4 = v76 + 23
				func_587000_154940(&v68, uint32(v64), uint32(v27))
				*memmap.PtrUint32(0x85B3FC, uint32(v31*4+228)) = 1
				if *(*uint32)(unsafe.Pointer(uintptr(v30 + 40))) != 0 {
					nox_xxx_tileCallDrawEdges_481BF0(int32(uintptr(unsafe.Pointer(&v68))), int32(*(*uint32)(unsafe.Pointer(uintptr(v30 + 40)))))
				}
			}
			if int32(*(*uint8)(unsafe.Pointer(uintptr(v30))))&1 != 0 {
				*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v27))), unsafe.Sizeof(uint16(0))*0)) = *(*uint16)(unsafe.Pointer(uintptr(v30 + 4)))
				v32 = int32(*(*uint16)(unsafe.Pointer(uintptr(v30 + 4))))
				v65 = int32(*(*uint32)(unsafe.Pointer(uintptr(*memmap.PtrUint32(0x85B3FC, uint32(v32*60+0x7F04)) + (*(*uint32)(unsafe.Pointer(uintptr(v30 + 8)))+uint32(*memmap.PtrUint16(0x85B3FC, uint32(v32*60+32530))))*4))))
				v68.field_0 = v26 + 23
				v68.field_4 = v76
				func_587000_154940(&v68, uint32(v65), uint32(v27))
				*memmap.PtrUint32(0x85B3FC, uint32(v32*4+228)) = 1
				v33 = int32(*(*uint32)(unsafe.Pointer(uintptr(v30 + 20))))
				if v33 != 0 {
					nox_xxx_tileCallDrawEdges_481BF0(int32(uintptr(unsafe.Pointer(&v68))), v33)
				}
			}
			v26 += 46
			if func() int32 {
				p := &i
				*p++
				return *p
			}() >= j {
				break
			}
			v27 = v71
			v25 = i
		}
	}
}
func nox_xxx_tileDrawMB_481C20_C_textured(vp *nox_draw_viewport_t, v72 int32, v78 int32) {
	var (
		v48 *uint8
		v49 int32
		v50 int32
		v51 *uint8
		v52 int32
		v53 *int32
		v54 int32
		v55 int32
		v56 uint32
		v57 uint32
		v58 int32
		v59 int32
		v60 int32
		v61 int32
		v66 int2
		v67 int2
		v69 int32
		v70 int32
		v75 *uint8
		v80 int32
	)
	v67.field_4 = int32(uint32(v78) + dword_5d4594_3679320)
	sub_4745F0(vp)
	v48 = &nox_arr_84EB20[0]
	for {
		*(*uint32)(unsafe.Pointer(v48)) = math.MaxUint32
		v48 = (*uint8)(unsafe.Add(unsafe.Pointer(v48), 280))
		if (int32(uintptr(unsafe.Pointer(v48)) - uintptr(unsafe.Pointer(&nox_arr_84EB20[0])))) >= int32(unsafe.Sizeof([63840]uint8{})) {
			break
		}
	}
	v66.field_4 = (v67.field_4 - *(*int32)(unsafe.Pointer(&dword_5d4594_2650680))) / 23
	v49 = int32(dword_5d4594_3679320)
	v50 = int32(uint32(v67.field_4-v66.field_4*23) - dword_5d4594_2650680)
	v78 = int32(uint32(v67.field_4-v66.field_4*23) - dword_5d4594_2650680)
	v80 = int32(dword_5d4594_3679320)
	if *(*int32)(unsafe.Pointer(&dword_5d4594_3679320)) < *(*int32)(unsafe.Pointer(&dword_5d4594_3798156)) {
		v51 = (*uint8)(unsafe.Pointer(&nox_arr_957820[dword_5d4594_3679320*128]))
		var v76 int32 = int32(dword_5d4594_3679320*4 - 4)
		var ishift uint8 = (*memmap.PtrUint8(0x973F18, 7696))
		var func1 unsafe.Pointer = memmap.PtrOff(0x973F18, 7720)
		for j := int32(int32(uintptr(unsafe.Pointer(&nox_arr_957820[dword_5d4594_3679320*128])))); ; v51 = (*uint8)(unsafe.Pointer(uintptr(j))) {
			var i int32 = int32(uint32(uintptr(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(nox_pixbuffer_rows_3798784), unsafe.Sizeof((*uint8)(nil))*uintptr(v49)))))))
			if v50 == 23 {
				v78 = 0
				v66.field_4++
			}
			v52 = int32(nox_arr_956A00[v76/4+1])
			v53 = (*int32)(unsafe.Pointer(v51))
			if nox_client_highResFloors_154952 != 0 || (v67.field_4&1) == 0 || v51 == &nox_arr_957820[0] {
				if v52 > 0 {
					var v71 int32 = int32(uint32(v52+1) >> 1)
					for {
						{
							v58 = *(*int32)(unsafe.Add(unsafe.Pointer(v53), unsafe.Sizeof(int32(0))*1))
							v69 = *v53
							v59 = v69 + v72
							v53 = (*int32)(unsafe.Add(unsafe.Pointer(v53), unsafe.Sizeof(int32(0))*2))
							v67.field_0 = v69 + v72
							v60 = (v69 + v72 - *(*int32)(unsafe.Pointer(&dword_5d4594_2650676))) / 23
							v66.field_0 = v60
							v70 = (v69 + v72 - *(*int32)(unsafe.Pointer(&dword_5d4594_2650676))) / 23
							v61 = int32(dword_5d4594_2650676 + uint32(v60*23))
							var v74 uint32 = uint32(v58 - v69)
							v75 = (*uint8)(unsafe.Pointer(&nox_arr_84EB20[v60*280]))
							if v61 < v58+v72+23 {
								for {
									if *(*uint32)(unsafe.Pointer(v75)) != uint32(v66.field_4) {
										(cgoAsFunc(*(*uint32)(func1), (*func(*int2) int32)(nil)))(&v66)
										v59 = v67.field_0
										v60 = v66.field_0
									}
									v61 += 23
									v75 = (*uint8)(unsafe.Add(unsafe.Pointer(v75), 280))
									v66.field_0 = func() int32 {
										p := &v60
										*p++
										return *p
									}()
									if v61 >= int32(v74+uint32(v59)+23) {
										break
									}
								}
							}
							v66.field_0 = v70
							(cgoAsFunc(*(*uint32)(unsafe.Pointer(&dword_5d4594_3805484)), (*func(*int2, *int32, uint32, uint32, uint32))(nil)))(&v67, &v66.field_0, uint32(v78), v74, uint32(i+(v69<<int32(ishift))))
							v71--
						}
						if v71 == 0 {
							break
						}
					}
				}
			} else if v52 > 0 {
				var v74 uint32 = uint32(v52+1) >> 1
				for {
					v54 = *v53
					v55 = *(*int32)(unsafe.Add(unsafe.Pointer(v53), unsafe.Sizeof(int32(0))*1))
					v53 = (*int32)(unsafe.Add(unsafe.Pointer(v53), unsafe.Sizeof(int32(0))*2))
					v56 = uint32((v55 - v54) << int32(ishift))
					v57 = v74
					libc.MemCpy(unsafe.Pointer(uintptr(i+(v54<<int32(ishift)))), unsafe.Pointer(uintptr(uint32(v54<<int32(ishift))+uint32(uintptr(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(nox_pixbuffer_rows_3798784), unsafe.Sizeof((*uint8)(nil))*uintptr(v76/4)))))))), int(v56))
					v74 = v57 - 1
					if v57 == 1 {
						break
					}
				}
			}
			v78++
			var v18 bool = func() int32 {
				p := &v80
				*p++
				return *p
			}() < *(*int32)(unsafe.Pointer(&dword_5d4594_3798156))
			j += 128
			v76 += 4
			v67.field_4++
			if !v18 {
				break
			}
			v49 = v80
			v50 = v78
		}
	}
}
func nox_xxx_tileDrawMB_481C20_C_solid(vp *nox_draw_viewport_t, v72 int32, v78 int32) {
	var (
		v35 *uint8
		v36 int32
		v37 int32
		v38 int32
		v39 *byte
		v40 uint32
		v41 int32
		v42 *byte
		v43 *byte
		v44 *byte
		v45 int8
		v46 bool
		v47 int32
		v67 int2
		v79 int32
	)
	v67.field_4 = int32(uint32(v78) + dword_5d4594_3679320)
	sub_4745F0(vp)
	v79 = int32(dword_5d4594_3679320)
	if *(*int32)(unsafe.Pointer(&dword_5d4594_3679320)) < *(*int32)(unsafe.Pointer(&dword_5d4594_3798156)) {
		var (
			v78a *uint8 = (*uint8)(unsafe.Pointer(&nox_arr_957820[dword_5d4594_3679320*128]))
			v18  bool
		)
		for {
			{
				v35 = v78a
				v36 = int32(nox_arr_956A00[v79])
				var i int32 = int32(uintptr(unsafe.Pointer(*(**uint8)(unsafe.Add(unsafe.Pointer(nox_pixbuffer_rows_3798784), unsafe.Sizeof((*uint8)(nil))*uintptr(v79))))))
				if v36 > 0 {
					var j int32 = int32(uint32(v36+1) >> 1)
					for {
						{
							v37 = int32(*(*uint32)(unsafe.Pointer(v35)))
							v38 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v35))), unsafe.Sizeof(uint32(0))*1))) - *(*uint32)(unsafe.Pointer(v35)))
							v67.field_0 = int32(*(*uint32)(unsafe.Pointer(v35)) + uint32(v72))
							v35 = (*uint8)(unsafe.Add(unsafe.Pointer(v35), 8))
							v39 = (*byte)(unsafe.Pointer(uintptr(dword_5d4594_3798804*(uint32(v67.field_4)+dword_5d4594_3798840-dword_5d4594_3798824) + dword_5d4594_3798796 + ((uint32(v37+v72) + dword_5d4594_3798836 - dword_5d4594_3798820) << uint32(*memmap.PtrUint8(0x973F18, 7696))))))
							if uint32(uintptr(unsafe.Pointer(v39))) >= uint32(*(*int32)(unsafe.Pointer(&dword_5d4594_3798844))) {
								v39 = (*byte)(unsafe.Add(unsafe.Pointer(v39), dword_5d4594_3798796-dword_5d4594_3798844))
							}
							var v76 int32 = i + (v37 << int32(*memmap.PtrUint8(0x973F18, 7696)))
							v40 = uint32(v38 << int32(*memmap.PtrUint8(0x973F18, 7696)))
							if uint32(uintptr(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v39), v38<<int32(*memmap.PtrUint8(0x973F18, 7696))))))) < uint32(*(*int32)(unsafe.Pointer(&dword_5d4594_3798844))) {
								v42 = (*byte)(unsafe.Pointer(uintptr(v76)))
								libc.MemCpy(unsafe.Pointer(uintptr(v76)), unsafe.Pointer(v39), int((v40>>2)*4))
							} else {
								v41 = int32(dword_5d4594_3798844 - uint32(uintptr(unsafe.Pointer(v39))))
								libc.MemCpy(unsafe.Pointer(uintptr(v76)), unsafe.Pointer(v39), int(dword_5d4594_3798844-uint32(uintptr(unsafe.Pointer(v39)))))
								v39 = *(**byte)(unsafe.Pointer(&dword_5d4594_3798796))
								v42 = (*byte)(unsafe.Pointer(uintptr(v41 + v76)))
								v40 = uint32((v38 << int32(*memmap.PtrUint8(0x973F18, 7696))) - v41)
								libc.MemCpy(unsafe.Pointer(v42), *(*unsafe.Pointer)(unsafe.Pointer(&dword_5d4594_3798796)), int((v40>>2)*4))
							}
							v44 = (*byte)(unsafe.Add(unsafe.Pointer(v39), (v40>>2)*4))
							v43 = (*byte)(unsafe.Add(unsafe.Pointer(v42), (v40>>2)*4))
							v45 = int8(uint8(v40))
							v47 = j - 1
							v46 = j == 1
							libc.MemCpy(unsafe.Pointer(v43), unsafe.Pointer(v44), int(int32(v45)&3))
							j = v47
						}
						if v46 {
							break
						}
					}
				}
				v18 = func() int32 {
					p := &v79
					*p++
					return *p
				}() < *(*int32)(unsafe.Pointer(&dword_5d4594_3798156))
				v78a = (*uint8)(unsafe.Add(unsafe.Pointer(v78a), 128))
				v67.field_4++
			}
			if !v18 {
				break
			}
		}
	}
}
func nox_xxx_tileCheckRedrawMB_482570(vp *nox_draw_viewport_t) int32 {
	var (
		a1  *uint32 = (*uint32)(unsafe.Pointer(vp))
		v1  int32
		v2  int32
		v3  int32
		v4  int32
		i   int32
		v6  int32
		v7  *uint32
		v8  int32
		v10 int32
		v11 int32
		v12 int32
	)
	v1 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*5)) - *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*1)))
	v2 = int32((*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*4)) - *a1 - 11) / 46)
	if v2 < 0 {
		v2 = 0
	}
	v12 = int32(dword_5d4594_3798812 + uint32(v2) - 1)
	if v12 >= 128 {
		v12 = math.MaxInt8
		v2 = int32(math.MaxInt8 - dword_5d4594_3798812)
	}
	v3 = (v1-11)/46 - 1
	if v3 < 0 {
		v3 = 0
	}
	v4 = int32(dword_5d4594_3798816 + uint32(v3))
	v10 = int32(dword_5d4594_3798816 + uint32(v3))
	if *(*int32)(unsafe.Pointer(&dword_5d4594_3798816))+v3 >= 128 {
		v10 = math.MaxInt8
		v3 = int32(math.MaxInt8 - dword_5d4594_3798816)
		v4 = math.MaxInt8
	}
	v11 = v3
	if v3 >= v4 {
		return 0
	}
	for i = v3 * 44; ; i += 44 {
		v6 = v2
		if v2 < v12 {
			break
		}
	LABEL_19:
		if func() int32 {
			p := &v11
			*p++
			return *p
		}() >= v4 {
			return 0
		}
	}
	v7 = (*uint32)(unsafe.Pointer((**obj_5D4594_2650668_t)(unsafe.Add(unsafe.Pointer(ptr_5D4594_2650668), unsafe.Sizeof((*obj_5D4594_2650668_t)(nil))*uintptr(v2)))))
	for {
		v8 = int32(uint32(i) + *v7)
		if int32(*(*uint8)(unsafe.Pointer(uintptr(v8))))&1 != 0 {
			if (int32(*memmap.PtrUint8(0x85B3FC, *(*uint32)(unsafe.Pointer(uintptr(v8 + 4)))*60+0x7F1E)) & 1) == 1 {
				return 1
			}
		}
		if int32(*(*uint8)(unsafe.Pointer(uintptr(v8))))&2 != 0 && (int32(*memmap.PtrUint8(0x85B3FC, *(*uint32)(unsafe.Pointer(uintptr(v8 + 24)))*60+0x7F1E))&1) == 1 {
			return 1
		}
		v6++
		v7 = (*uint32)(unsafe.Add(unsafe.Pointer(v7), unsafe.Sizeof(uint32(0))*1))
		if v6 >= v12 {
			v4 = v10
			goto LABEL_19
		}
	}
}
func nox_xxx_tileDrawImpl_4826A0(vp *nox_draw_viewport_t) int32 {
	var (
		a1     *uint32 = (*uint32)(unsafe.Pointer(vp))
		v1     int32
		v2     int32
		result int32
		v4     int32
		v5     int32
		v6     int32
		v7     int32
		v8     int32
		v9     int32
		v10    int32
		v11    int32
		v12    int8
		v13    int32
		v14    int32
		v15    int32
		v16    int32
		v17    bool
		v18    int32
		v19    int32
		v20    int32
		v21    int32
		v22    int32
		v23    int32
		v24    int2
		v25    int32
	)
	v1 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*4)) - *a1)
	v2 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*5)) - *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*1)))
	dword_5d4594_3798836 = 0
	dword_5d4594_3798840 = 0
	sub_481410()
	libc.MemSet(memmap.PtrOff(0x85B3FC, 228), 0, 704)
	libc.MemSet(memmap.PtrOff(6112660, 2523980), 0, 256)
	v25 = (v1 - 11) / 46
	if v25 < 0 {
		v25 = 0
	}
	v20 = int32(dword_5d4594_3798812 + uint32(v25) - 1)
	if v20 >= 128 {
		v20 = math.MaxInt8
		v25 = int32(math.MaxInt8 - dword_5d4594_3798812)
	}
	result = (v2-11)/46 - 1
	if result < 0 {
		result = 0
	}
	v4 = int32(dword_5d4594_3798816 + uint32(result))
	if *(*int32)(unsafe.Pointer(&dword_5d4594_3798816))+result >= 128 {
		v4 = math.MaxInt8
		result = int32(math.MaxInt8 - dword_5d4594_3798816)
	}
	v5 = v25
	dword_5d4594_3798824 = uint32(result*46 - 11)
	v6 = result*46 - 57
	dword_5d4594_3798828 = uint32(v25)
	dword_5d4594_3798820 = uint32(v25*46 - 11)
	dword_5d4594_3798832 = uint32(result)
	if result < v4 {
		v21 = result * 44
		v23 = v4 - result
		v7 = v20
		for {
			v8 = v25*46 - 57
			v6 += 46
			v9 = v5
			v22 = v5
			if v5 < v7 {
				for {
					*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v10))), unsafe.Sizeof(uint16(0))*1)) = *(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&ptr_5D4594_2650668))), unsafe.Sizeof(uint16(0))*1))
					v8 += 46
					v11 = int32(uintptr(unsafe.Pointer(*(**obj_5D4594_2650668_t)(unsafe.Add(unsafe.Pointer(ptr_5D4594_2650668), unsafe.Sizeof((*obj_5D4594_2650668_t)(nil))*uintptr(v9))))))
					v12 = int8(*(*uint8)(unsafe.Pointer(uintptr(v11 + v21))))
					v13 = v21 + v11
					if int32(v12) != 0 {
						if int32(v12)&2 != 0 {
							*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v10))), unsafe.Sizeof(uint16(0))*0)) = *(*uint16)(unsafe.Pointer(uintptr(v13 + 24)))
							v14 = int32(*(*uint16)(unsafe.Pointer(uintptr(v13 + 24))))
							v18 = int32(*(*uint32)(unsafe.Pointer(uintptr(*memmap.PtrUint32(0x85B3FC, uint32(v14*60+0x7F04)) + (*(*uint32)(unsafe.Pointer(uintptr(v13 + 28)))+uint32(*memmap.PtrUint16(0x85B3FC, uint32(v14*60+32530))))*4))))
							v24.field_0 = v8
							v24.field_4 = v6 + 23
							func_587000_154940(&v24, uint32(v18), uint32(v10))
							*memmap.PtrUint32(0x85B3FC, uint32(v14*4+228)) = 1
							if *(*uint32)(unsafe.Pointer(uintptr(v13 + 40))) != 0 {
								nox_xxx_tileCallDrawEdges_481BF0(int32(uintptr(unsafe.Pointer(&v24))), int32(*(*uint32)(unsafe.Pointer(uintptr(v13 + 40)))))
							}
						}
						if int32(*(*uint8)(unsafe.Pointer(uintptr(v13))))&1 != 0 {
							*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v10))), unsafe.Sizeof(uint16(0))*0)) = *(*uint16)(unsafe.Pointer(uintptr(v13 + 4)))
							v15 = int32(*(*uint16)(unsafe.Pointer(uintptr(v13 + 4))))
							v19 = int32(*(*uint32)(unsafe.Pointer(uintptr(*memmap.PtrUint32(0x85B3FC, uint32(v15*60+0x7F04)) + (*(*uint32)(unsafe.Pointer(uintptr(v13 + 8)))+uint32(*memmap.PtrUint16(0x85B3FC, uint32(v15*60+32530))))*4))))
							v24.field_0 = v8 + 23
							v24.field_4 = v6
							func_587000_154940(&v24, uint32(v19), uint32(v10))
							*memmap.PtrUint32(0x85B3FC, uint32(v15*4+228)) = 1
							v16 = int32(*(*uint32)(unsafe.Pointer(uintptr(v13 + 20))))
							if v16 != 0 {
								nox_xxx_tileCallDrawEdges_481BF0(int32(uintptr(unsafe.Pointer(&v24))), v16)
							}
						}
					}
					v7 = v20
					v9 = func() int32 {
						p := &v22
						*p++
						return *p
					}()
					if v22 >= v20 {
						break
					}
				}
				v5 = v25
			}
			result = v23 - 1
			v17 = v23 == 1
			v21 += 44
			v23--
			if v17 {
				break
			}
		}
	}
	return result
}
func sub_482910(a1 int32, a2 int32) int8 {
	var (
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
		v13    int32
		v14    int32
		v15    int32
		v16    int32
		v17    int32
		v18    int32
		v19    int32
		v20    int32
		v21    int32
		v22    int32
		v23    int32
		v24    int32
		v25    int32
		v26    int32
		v27    int32
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
		v42    int32
		v43    int32
		v44    int32
		v45    int32
		v46    int32
		result int8
	)
	v2 = int32(dword_5d4594_3798804)
	*(*uint8)(unsafe.Pointer(uintptr(a2 + 23))) = *(*uint8)(unsafe.Pointer(uintptr(a1)))
	v3 = v2 + a2
	*(*uint16)(unsafe.Pointer(uintptr(v3 + 22))) = *(*uint16)(unsafe.Pointer(uintptr(a1 + 1)))
	*(*uint8)(unsafe.Pointer(uintptr(v3 + 24))) = *(*uint8)(unsafe.Pointer(uintptr(a1 + 3)))
	v4 = v2 + v2 + a2
	*(*uint8)(unsafe.Pointer(uintptr(v4 + 21))) = *(*uint8)(unsafe.Pointer(uintptr(a1 + 4)))
	*(*uint16)(unsafe.Pointer(uintptr(v4 + 22))) = *(*uint16)(unsafe.Pointer(uintptr(a1 + 5)))
	*(*uint16)(unsafe.Pointer(uintptr(v4 + 24))) = *(*uint16)(unsafe.Pointer(uintptr(a1 + 7)))
	v5 = v2 + v4
	*(*uint32)(unsafe.Pointer(uintptr(v5 + 20))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 9)))
	*(*uint16)(unsafe.Pointer(uintptr(v5 + 24))) = *(*uint16)(unsafe.Pointer(uintptr(a1 + 13)))
	*(*uint8)(unsafe.Pointer(uintptr(v5 + 26))) = *(*uint8)(unsafe.Pointer(uintptr(a1 + 15)))
	v6 = v2 + v5
	*(*uint8)(unsafe.Pointer(uintptr(v6 + 19))) = *(*uint8)(unsafe.Pointer(uintptr(a1 + 16)))
	*(*uint32)(unsafe.Pointer(uintptr(v6 + 20))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 17)))
	*(*uint32)(unsafe.Pointer(uintptr(v6 + 24))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 21)))
	v7 = v2 + v6
	*(*uint16)(unsafe.Pointer(uintptr(v7 + 18))) = *(*uint16)(unsafe.Pointer(uintptr(a1 + 25)))
	*(*uint32)(unsafe.Pointer(uintptr(v7 + 20))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 27)))
	*(*uint32)(unsafe.Pointer(uintptr(v7 + 24))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 31)))
	*(*uint8)(unsafe.Pointer(uintptr(v7 + 28))) = *(*uint8)(unsafe.Pointer(uintptr(a1 + 35)))
	v8 = v2 + v7
	*(*uint8)(unsafe.Pointer(uintptr(v8 + 17))) = *(*uint8)(unsafe.Pointer(uintptr(a1 + 36)))
	*(*uint16)(unsafe.Pointer(uintptr(v8 + 18))) = *(*uint16)(unsafe.Pointer(uintptr(a1 + 37)))
	*(*uint32)(unsafe.Pointer(uintptr(v8 + 20))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 39)))
	*(*uint32)(unsafe.Pointer(uintptr(v8 + 24))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 43)))
	*(*uint16)(unsafe.Pointer(uintptr(v8 + 28))) = *(*uint16)(unsafe.Pointer(uintptr(a1 + 47)))
	v9 = v2 + v8
	*(*uint32)(unsafe.Pointer(uintptr(v9 + 16))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 49)))
	*(*uint32)(unsafe.Pointer(uintptr(v9 + 20))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 53)))
	*(*uint32)(unsafe.Pointer(uintptr(v9 + 24))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 57)))
	*(*uint16)(unsafe.Pointer(uintptr(v9 + 28))) = *(*uint16)(unsafe.Pointer(uintptr(a1 + 61)))
	*(*uint8)(unsafe.Pointer(uintptr(v9 + 30))) = *(*uint8)(unsafe.Pointer(uintptr(a1 + 63)))
	v10 = v2 + v9
	*(*uint8)(unsafe.Pointer(uintptr(v10 + 15))) = *(*uint8)(unsafe.Pointer(uintptr(a1 + 64)))
	*(*uint32)(unsafe.Pointer(uintptr(v10 + 16))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 65)))
	*(*uint32)(unsafe.Pointer(uintptr(v10 + 20))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 69)))
	*(*uint32)(unsafe.Pointer(uintptr(v10 + 24))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 73)))
	*(*uint32)(unsafe.Pointer(uintptr(v10 + 28))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 77)))
	v11 = v2 + v10
	*(*uint16)(unsafe.Pointer(uintptr(v11 + 14))) = *(*uint16)(unsafe.Pointer(uintptr(a1 + 81)))
	*(*uint32)(unsafe.Pointer(uintptr(v11 + 16))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 83)))
	*(*uint32)(unsafe.Pointer(uintptr(v11 + 20))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 87)))
	*(*uint32)(unsafe.Pointer(uintptr(v11 + 24))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 91)))
	*(*uint32)(unsafe.Pointer(uintptr(v11 + 28))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 95)))
	*(*uint8)(unsafe.Pointer(uintptr(v11 + 32))) = *(*uint8)(unsafe.Pointer(uintptr(a1 + 99)))
	v12 = v2 + v11
	*(*uint8)(unsafe.Pointer(uintptr(v12 + 13))) = *(*uint8)(unsafe.Pointer(uintptr(a1 + 100)))
	*(*uint16)(unsafe.Pointer(uintptr(v12 + 14))) = *(*uint16)(unsafe.Pointer(uintptr(a1 + 101)))
	*(*uint32)(unsafe.Pointer(uintptr(v12 + 16))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 103)))
	*(*uint32)(unsafe.Pointer(uintptr(v12 + 20))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 107)))
	*(*uint32)(unsafe.Pointer(uintptr(v12 + 24))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 111)))
	*(*uint32)(unsafe.Pointer(uintptr(v12 + 28))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 115)))
	*(*uint16)(unsafe.Pointer(uintptr(v12 + 32))) = *(*uint16)(unsafe.Pointer(uintptr(a1 + 119)))
	v13 = v2 + v12
	*(*uint32)(unsafe.Pointer(uintptr(v13 + 12))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 121)))
	*(*uint32)(unsafe.Pointer(uintptr(v13 + 16))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 125)))
	*(*uint32)(unsafe.Pointer(uintptr(v13 + 20))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 129)))
	*(*uint32)(unsafe.Pointer(uintptr(v13 + 24))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 133)))
	*(*uint32)(unsafe.Pointer(uintptr(v13 + 28))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 137)))
	*(*uint16)(unsafe.Pointer(uintptr(v13 + 32))) = *(*uint16)(unsafe.Pointer(uintptr(a1 + 141)))
	*(*uint8)(unsafe.Pointer(uintptr(v13 + 34))) = *(*uint8)(unsafe.Pointer(uintptr(a1 + 143)))
	v14 = v2 + v13
	*(*uint8)(unsafe.Pointer(uintptr(v14 + 11))) = *(*uint8)(unsafe.Pointer(uintptr(a1 + 144)))
	*(*uint32)(unsafe.Pointer(uintptr(v14 + 12))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 145)))
	*(*uint32)(unsafe.Pointer(uintptr(v14 + 16))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 149)))
	*(*uint32)(unsafe.Pointer(uintptr(v14 + 20))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 153)))
	*(*uint32)(unsafe.Pointer(uintptr(v14 + 24))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 157)))
	*(*uint32)(unsafe.Pointer(uintptr(v14 + 28))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 161)))
	*(*uint32)(unsafe.Pointer(uintptr(v14 + 32))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 165)))
	v15 = v2 + v14
	*(*uint16)(unsafe.Pointer(uintptr(v15 + 10))) = *(*uint16)(unsafe.Pointer(uintptr(a1 + 169)))
	*(*uint32)(unsafe.Pointer(uintptr(v15 + 12))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 171)))
	*(*uint32)(unsafe.Pointer(uintptr(v15 + 16))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 175)))
	*(*uint32)(unsafe.Pointer(uintptr(v15 + 20))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 179)))
	*(*uint32)(unsafe.Pointer(uintptr(v15 + 24))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 183)))
	*(*uint32)(unsafe.Pointer(uintptr(v15 + 28))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 187)))
	*(*uint32)(unsafe.Pointer(uintptr(v15 + 32))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 191)))
	*(*uint8)(unsafe.Pointer(uintptr(v15 + 36))) = *(*uint8)(unsafe.Pointer(uintptr(a1 + 195)))
	v16 = v2 + v15
	*(*uint8)(unsafe.Pointer(uintptr(v16 + 9))) = *(*uint8)(unsafe.Pointer(uintptr(a1 + 196)))
	*(*uint16)(unsafe.Pointer(uintptr(v16 + 10))) = *(*uint16)(unsafe.Pointer(uintptr(a1 + 197)))
	*(*uint32)(unsafe.Pointer(uintptr(v16 + 12))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 199)))
	*(*uint32)(unsafe.Pointer(uintptr(v16 + 16))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 203)))
	*(*uint32)(unsafe.Pointer(uintptr(v16 + 20))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 207)))
	*(*uint32)(unsafe.Pointer(uintptr(v16 + 24))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 211)))
	*(*uint32)(unsafe.Pointer(uintptr(v16 + 28))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 215)))
	*(*uint32)(unsafe.Pointer(uintptr(v16 + 32))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 219)))
	*(*uint16)(unsafe.Pointer(uintptr(v16 + 36))) = *(*uint16)(unsafe.Pointer(uintptr(a1 + 223)))
	v17 = v2 + v16
	*(*uint32)(unsafe.Pointer(uintptr(v17 + 8))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 225)))
	*(*uint32)(unsafe.Pointer(uintptr(v17 + 12))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 229)))
	*(*uint32)(unsafe.Pointer(uintptr(v17 + 16))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 233)))
	*(*uint32)(unsafe.Pointer(uintptr(v17 + 20))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 237)))
	*(*uint32)(unsafe.Pointer(uintptr(v17 + 24))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 241)))
	*(*uint32)(unsafe.Pointer(uintptr(v17 + 28))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 245)))
	*(*uint32)(unsafe.Pointer(uintptr(v17 + 32))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 249)))
	*(*uint16)(unsafe.Pointer(uintptr(v17 + 36))) = *(*uint16)(unsafe.Pointer(uintptr(a1 + 253)))
	*(*uint8)(unsafe.Pointer(uintptr(v17 + 38))) = *(*uint8)(unsafe.Pointer(uintptr(a1 + math.MaxUint8)))
	v18 = v2 + v17
	*(*uint8)(unsafe.Pointer(uintptr(v18 + 7))) = *(*uint8)(unsafe.Pointer(uintptr(a1 + 256)))
	*(*uint32)(unsafe.Pointer(uintptr(v18 + 8))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 257)))
	*(*uint32)(unsafe.Pointer(uintptr(v18 + 12))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 261)))
	*(*uint32)(unsafe.Pointer(uintptr(v18 + 16))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 265)))
	*(*uint32)(unsafe.Pointer(uintptr(v18 + 20))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 269)))
	*(*uint32)(unsafe.Pointer(uintptr(v18 + 24))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 273)))
	*(*uint32)(unsafe.Pointer(uintptr(v18 + 28))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 277)))
	*(*uint32)(unsafe.Pointer(uintptr(v18 + 32))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 281)))
	*(*uint32)(unsafe.Pointer(uintptr(v18 + 36))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 285)))
	v19 = v2 + v18
	*(*uint16)(unsafe.Pointer(uintptr(v19 + 6))) = *(*uint16)(unsafe.Pointer(uintptr(a1 + 289)))
	*(*uint32)(unsafe.Pointer(uintptr(v19 + 8))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 291)))
	*(*uint32)(unsafe.Pointer(uintptr(v19 + 12))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 295)))
	*(*uint32)(unsafe.Pointer(uintptr(v19 + 16))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 299)))
	*(*uint32)(unsafe.Pointer(uintptr(v19 + 20))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 303)))
	*(*uint32)(unsafe.Pointer(uintptr(v19 + 24))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 307)))
	*(*uint32)(unsafe.Pointer(uintptr(v19 + 28))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 311)))
	*(*uint32)(unsafe.Pointer(uintptr(v19 + 32))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 315)))
	*(*uint32)(unsafe.Pointer(uintptr(v19 + 36))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 319)))
	*(*uint8)(unsafe.Pointer(uintptr(v19 + 40))) = *(*uint8)(unsafe.Pointer(uintptr(a1 + 323)))
	v20 = v2 + v19
	*(*uint8)(unsafe.Pointer(uintptr(v20 + 5))) = *(*uint8)(unsafe.Pointer(uintptr(a1 + 324)))
	*(*uint16)(unsafe.Pointer(uintptr(v20 + 6))) = *(*uint16)(unsafe.Pointer(uintptr(a1 + 325)))
	*(*uint32)(unsafe.Pointer(uintptr(v20 + 8))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 327)))
	*(*uint32)(unsafe.Pointer(uintptr(v20 + 12))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 331)))
	*(*uint32)(unsafe.Pointer(uintptr(v20 + 16))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 335)))
	*(*uint32)(unsafe.Pointer(uintptr(v20 + 20))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 339)))
	*(*uint32)(unsafe.Pointer(uintptr(v20 + 24))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 343)))
	*(*uint32)(unsafe.Pointer(uintptr(v20 + 28))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 347)))
	*(*uint32)(unsafe.Pointer(uintptr(v20 + 32))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 351)))
	*(*uint32)(unsafe.Pointer(uintptr(v20 + 36))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 355)))
	*(*uint16)(unsafe.Pointer(uintptr(v20 + 40))) = *(*uint16)(unsafe.Pointer(uintptr(a1 + 359)))
	v21 = v2 + v20
	*(*uint32)(unsafe.Pointer(uintptr(v21 + 4))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 361)))
	*(*uint32)(unsafe.Pointer(uintptr(v21 + 8))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 365)))
	*(*uint32)(unsafe.Pointer(uintptr(v21 + 12))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 369)))
	*(*uint32)(unsafe.Pointer(uintptr(v21 + 16))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 373)))
	*(*uint32)(unsafe.Pointer(uintptr(v21 + 20))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 377)))
	*(*uint32)(unsafe.Pointer(uintptr(v21 + 24))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 381)))
	*(*uint32)(unsafe.Pointer(uintptr(v21 + 28))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 385)))
	*(*uint32)(unsafe.Pointer(uintptr(v21 + 32))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 389)))
	*(*uint32)(unsafe.Pointer(uintptr(v21 + 36))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 393)))
	*(*uint16)(unsafe.Pointer(uintptr(v21 + 40))) = *(*uint16)(unsafe.Pointer(uintptr(a1 + 397)))
	*(*uint8)(unsafe.Pointer(uintptr(v21 + 42))) = *(*uint8)(unsafe.Pointer(uintptr(a1 + 399)))
	v22 = v2 + v21
	*(*uint8)(unsafe.Pointer(uintptr(v22 + 3))) = *(*uint8)(unsafe.Pointer(uintptr(a1 + 400)))
	*(*uint32)(unsafe.Pointer(uintptr(v22 + 4))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 401)))
	*(*uint32)(unsafe.Pointer(uintptr(v22 + 8))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 405)))
	*(*uint32)(unsafe.Pointer(uintptr(v22 + 12))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 409)))
	*(*uint32)(unsafe.Pointer(uintptr(v22 + 16))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 413)))
	*(*uint32)(unsafe.Pointer(uintptr(v22 + 20))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 417)))
	*(*uint32)(unsafe.Pointer(uintptr(v22 + 24))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 421)))
	*(*uint32)(unsafe.Pointer(uintptr(v22 + 28))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 425)))
	*(*uint32)(unsafe.Pointer(uintptr(v22 + 32))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 429)))
	*(*uint32)(unsafe.Pointer(uintptr(v22 + 36))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 433)))
	*(*uint32)(unsafe.Pointer(uintptr(v22 + 40))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 437)))
	v23 = v2 + v22
	*(*uint16)(unsafe.Pointer(uintptr(v23 + 2))) = *(*uint16)(unsafe.Pointer(uintptr(a1 + 441)))
	*(*uint32)(unsafe.Pointer(uintptr(v23 + 4))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 443)))
	*(*uint32)(unsafe.Pointer(uintptr(v23 + 8))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 447)))
	*(*uint32)(unsafe.Pointer(uintptr(v23 + 12))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 451)))
	*(*uint32)(unsafe.Pointer(uintptr(v23 + 16))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 455)))
	*(*uint32)(unsafe.Pointer(uintptr(v23 + 20))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 459)))
	*(*uint32)(unsafe.Pointer(uintptr(v23 + 24))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 463)))
	*(*uint32)(unsafe.Pointer(uintptr(v23 + 28))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 467)))
	*(*uint32)(unsafe.Pointer(uintptr(v23 + 32))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 471)))
	*(*uint32)(unsafe.Pointer(uintptr(v23 + 36))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 475)))
	*(*uint32)(unsafe.Pointer(uintptr(v23 + 40))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 479)))
	*(*uint8)(unsafe.Pointer(uintptr(v23 + 44))) = *(*uint8)(unsafe.Pointer(uintptr(a1 + 483)))
	v24 = v2 + v23
	*(*uint8)(unsafe.Pointer(uintptr(v24 + 1))) = *(*uint8)(unsafe.Pointer(uintptr(a1 + 484)))
	*(*uint16)(unsafe.Pointer(uintptr(v24 + 2))) = *(*uint16)(unsafe.Pointer(uintptr(a1 + 485)))
	*(*uint32)(unsafe.Pointer(uintptr(v24 + 4))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 487)))
	*(*uint32)(unsafe.Pointer(uintptr(v24 + 8))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 491)))
	*(*uint32)(unsafe.Pointer(uintptr(v24 + 12))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 495)))
	*(*uint32)(unsafe.Pointer(uintptr(v24 + 16))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 499)))
	*(*uint32)(unsafe.Pointer(uintptr(v24 + 20))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 503)))
	*(*uint32)(unsafe.Pointer(uintptr(v24 + 24))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 507)))
	*(*uint32)(unsafe.Pointer(uintptr(v24 + 28))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 511)))
	*(*uint32)(unsafe.Pointer(uintptr(v24 + 32))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 515)))
	*(*uint32)(unsafe.Pointer(uintptr(v24 + 36))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 519)))
	*(*uint32)(unsafe.Pointer(uintptr(v24 + 40))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 523)))
	*(*uint16)(unsafe.Pointer(uintptr(v24 + 44))) = *(*uint16)(unsafe.Pointer(uintptr(a1 + 527)))
	v25 = v2 + v24
	*(*uint8)(unsafe.Pointer(uintptr(v25 + 1))) = *(*uint8)(unsafe.Pointer(uintptr(a1 + 529)))
	*(*uint16)(unsafe.Pointer(uintptr(v25 + 2))) = *(*uint16)(unsafe.Pointer(uintptr(a1 + 530)))
	*(*uint32)(unsafe.Pointer(uintptr(v25 + 4))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 532)))
	*(*uint32)(unsafe.Pointer(uintptr(v25 + 8))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 536)))
	*(*uint32)(unsafe.Pointer(uintptr(v25 + 12))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 540)))
	*(*uint32)(unsafe.Pointer(uintptr(v25 + 16))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 544)))
	*(*uint32)(unsafe.Pointer(uintptr(v25 + 20))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 548)))
	*(*uint32)(unsafe.Pointer(uintptr(v25 + 24))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 552)))
	*(*uint32)(unsafe.Pointer(uintptr(v25 + 28))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 556)))
	*(*uint32)(unsafe.Pointer(uintptr(v25 + 32))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 560)))
	*(*uint32)(unsafe.Pointer(uintptr(v25 + 36))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 564)))
	*(*uint32)(unsafe.Pointer(uintptr(v25 + 40))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 568)))
	*(*uint16)(unsafe.Pointer(uintptr(v25 + 44))) = *(*uint16)(unsafe.Pointer(uintptr(a1 + 572)))
	v26 = v2 + v25
	*(*uint16)(unsafe.Pointer(uintptr(v26 + 2))) = *(*uint16)(unsafe.Pointer(uintptr(a1 + 574)))
	*(*uint32)(unsafe.Pointer(uintptr(v26 + 4))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 576)))
	*(*uint32)(unsafe.Pointer(uintptr(v26 + 8))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 580)))
	*(*uint32)(unsafe.Pointer(uintptr(v26 + 12))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 584)))
	*(*uint32)(unsafe.Pointer(uintptr(v26 + 16))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 588)))
	*(*uint32)(unsafe.Pointer(uintptr(v26 + 20))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 592)))
	*(*uint32)(unsafe.Pointer(uintptr(v26 + 24))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 596)))
	*(*uint32)(unsafe.Pointer(uintptr(v26 + 28))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 600)))
	*(*uint32)(unsafe.Pointer(uintptr(v26 + 32))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 604)))
	*(*uint32)(unsafe.Pointer(uintptr(v26 + 36))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 608)))
	*(*uint32)(unsafe.Pointer(uintptr(v26 + 40))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 612)))
	*(*uint8)(unsafe.Pointer(uintptr(v26 + 44))) = *(*uint8)(unsafe.Pointer(uintptr(a1 + 616)))
	v27 = v2 + v26
	*(*uint8)(unsafe.Pointer(uintptr(v27 + 3))) = *(*uint8)(unsafe.Pointer(uintptr(a1 + 617)))
	*(*uint32)(unsafe.Pointer(uintptr(v27 + 4))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 618)))
	*(*uint32)(unsafe.Pointer(uintptr(v27 + 8))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 622)))
	*(*uint32)(unsafe.Pointer(uintptr(v27 + 12))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 626)))
	*(*uint32)(unsafe.Pointer(uintptr(v27 + 16))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 630)))
	*(*uint32)(unsafe.Pointer(uintptr(v27 + 20))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 634)))
	*(*uint32)(unsafe.Pointer(uintptr(v27 + 24))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 638)))
	*(*uint32)(unsafe.Pointer(uintptr(v27 + 28))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 642)))
	*(*uint32)(unsafe.Pointer(uintptr(v27 + 32))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 646)))
	*(*uint32)(unsafe.Pointer(uintptr(v27 + 36))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 650)))
	*(*uint32)(unsafe.Pointer(uintptr(v27 + 40))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 654)))
	v28 = v2 + v27
	*(*uint32)(unsafe.Pointer(uintptr(v28 + 4))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 658)))
	*(*uint32)(unsafe.Pointer(uintptr(v28 + 8))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 662)))
	*(*uint32)(unsafe.Pointer(uintptr(v28 + 12))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 666)))
	*(*uint32)(unsafe.Pointer(uintptr(v28 + 16))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 670)))
	*(*uint32)(unsafe.Pointer(uintptr(v28 + 20))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 674)))
	*(*uint32)(unsafe.Pointer(uintptr(v28 + 24))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 678)))
	*(*uint32)(unsafe.Pointer(uintptr(v28 + 28))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 682)))
	*(*uint32)(unsafe.Pointer(uintptr(v28 + 32))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 686)))
	*(*uint32)(unsafe.Pointer(uintptr(v28 + 36))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 690)))
	*(*uint16)(unsafe.Pointer(uintptr(v28 + 40))) = *(*uint16)(unsafe.Pointer(uintptr(a1 + 694)))
	*(*uint8)(unsafe.Pointer(uintptr(v28 + 42))) = *(*uint8)(unsafe.Pointer(uintptr(a1 + 696)))
	v29 = v2 + v28
	*(*uint8)(unsafe.Pointer(uintptr(v29 + 5))) = *(*uint8)(unsafe.Pointer(uintptr(a1 + 697)))
	*(*uint16)(unsafe.Pointer(uintptr(v29 + 6))) = *(*uint16)(unsafe.Pointer(uintptr(a1 + 698)))
	*(*uint32)(unsafe.Pointer(uintptr(v29 + 8))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 700)))
	*(*uint32)(unsafe.Pointer(uintptr(v29 + 12))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 704)))
	*(*uint32)(unsafe.Pointer(uintptr(v29 + 16))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 708)))
	*(*uint32)(unsafe.Pointer(uintptr(v29 + 20))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 712)))
	*(*uint32)(unsafe.Pointer(uintptr(v29 + 24))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 716)))
	*(*uint32)(unsafe.Pointer(uintptr(v29 + 28))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 720)))
	*(*uint32)(unsafe.Pointer(uintptr(v29 + 32))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 724)))
	*(*uint32)(unsafe.Pointer(uintptr(v29 + 36))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 728)))
	*(*uint16)(unsafe.Pointer(uintptr(v29 + 40))) = *(*uint16)(unsafe.Pointer(uintptr(a1 + 732)))
	v30 = v2 + v29
	*(*uint16)(unsafe.Pointer(uintptr(v30 + 6))) = *(*uint16)(unsafe.Pointer(uintptr(a1 + 734)))
	*(*uint32)(unsafe.Pointer(uintptr(v30 + 8))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 736)))
	*(*uint32)(unsafe.Pointer(uintptr(v30 + 12))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 740)))
	*(*uint32)(unsafe.Pointer(uintptr(v30 + 16))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 744)))
	*(*uint32)(unsafe.Pointer(uintptr(v30 + 20))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 748)))
	*(*uint32)(unsafe.Pointer(uintptr(v30 + 24))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 752)))
	*(*uint32)(unsafe.Pointer(uintptr(v30 + 28))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 756)))
	*(*uint32)(unsafe.Pointer(uintptr(v30 + 32))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 760)))
	*(*uint32)(unsafe.Pointer(uintptr(v30 + 36))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 764)))
	*(*uint8)(unsafe.Pointer(uintptr(v30 + 40))) = *(*uint8)(unsafe.Pointer(uintptr(a1 + 768)))
	v31 = v2 + v30
	*(*uint8)(unsafe.Pointer(uintptr(v31 + 7))) = *(*uint8)(unsafe.Pointer(uintptr(a1 + 769)))
	*(*uint32)(unsafe.Pointer(uintptr(v31 + 8))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 770)))
	*(*uint32)(unsafe.Pointer(uintptr(v31 + 12))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 774)))
	*(*uint32)(unsafe.Pointer(uintptr(v31 + 16))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 778)))
	*(*uint32)(unsafe.Pointer(uintptr(v31 + 20))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 782)))
	*(*uint32)(unsafe.Pointer(uintptr(v31 + 24))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 786)))
	*(*uint32)(unsafe.Pointer(uintptr(v31 + 28))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 790)))
	*(*uint32)(unsafe.Pointer(uintptr(v31 + 32))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 794)))
	*(*uint32)(unsafe.Pointer(uintptr(v31 + 36))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 798)))
	v32 = v2 + v31
	*(*uint32)(unsafe.Pointer(uintptr(v32 + 8))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 802)))
	*(*uint32)(unsafe.Pointer(uintptr(v32 + 12))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 806)))
	*(*uint32)(unsafe.Pointer(uintptr(v32 + 16))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 810)))
	*(*uint32)(unsafe.Pointer(uintptr(v32 + 20))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 814)))
	*(*uint32)(unsafe.Pointer(uintptr(v32 + 24))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 818)))
	*(*uint32)(unsafe.Pointer(uintptr(v32 + 28))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 822)))
	*(*uint32)(unsafe.Pointer(uintptr(v32 + 32))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 826)))
	*(*uint16)(unsafe.Pointer(uintptr(v32 + 36))) = *(*uint16)(unsafe.Pointer(uintptr(a1 + 830)))
	*(*uint8)(unsafe.Pointer(uintptr(v32 + 38))) = *(*uint8)(unsafe.Pointer(uintptr(a1 + 832)))
	v33 = v2 + v32
	*(*uint8)(unsafe.Pointer(uintptr(v33 + 9))) = *(*uint8)(unsafe.Pointer(uintptr(a1 + 833)))
	*(*uint16)(unsafe.Pointer(uintptr(v33 + 10))) = *(*uint16)(unsafe.Pointer(uintptr(a1 + 834)))
	*(*uint32)(unsafe.Pointer(uintptr(v33 + 12))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 836)))
	*(*uint32)(unsafe.Pointer(uintptr(v33 + 16))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 840)))
	*(*uint32)(unsafe.Pointer(uintptr(v33 + 20))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 844)))
	*(*uint32)(unsafe.Pointer(uintptr(v33 + 24))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 848)))
	*(*uint32)(unsafe.Pointer(uintptr(v33 + 28))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 852)))
	*(*uint32)(unsafe.Pointer(uintptr(v33 + 32))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 856)))
	*(*uint16)(unsafe.Pointer(uintptr(v33 + 36))) = *(*uint16)(unsafe.Pointer(uintptr(a1 + 860)))
	v34 = v2 + v33
	*(*uint16)(unsafe.Pointer(uintptr(v34 + 10))) = *(*uint16)(unsafe.Pointer(uintptr(a1 + 862)))
	*(*uint32)(unsafe.Pointer(uintptr(v34 + 12))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 864)))
	*(*uint32)(unsafe.Pointer(uintptr(v34 + 16))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 868)))
	*(*uint32)(unsafe.Pointer(uintptr(v34 + 20))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 872)))
	*(*uint32)(unsafe.Pointer(uintptr(v34 + 24))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 876)))
	*(*uint32)(unsafe.Pointer(uintptr(v34 + 28))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 880)))
	*(*uint32)(unsafe.Pointer(uintptr(v34 + 32))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 884)))
	*(*uint8)(unsafe.Pointer(uintptr(v34 + 36))) = *(*uint8)(unsafe.Pointer(uintptr(a1 + 888)))
	v35 = v2 + v34
	*(*uint8)(unsafe.Pointer(uintptr(v35 + 11))) = *(*uint8)(unsafe.Pointer(uintptr(a1 + 889)))
	*(*uint32)(unsafe.Pointer(uintptr(v35 + 12))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 890)))
	*(*uint32)(unsafe.Pointer(uintptr(v35 + 16))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 894)))
	*(*uint32)(unsafe.Pointer(uintptr(v35 + 20))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 898)))
	*(*uint32)(unsafe.Pointer(uintptr(v35 + 24))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 902)))
	*(*uint32)(unsafe.Pointer(uintptr(v35 + 28))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 906)))
	*(*uint32)(unsafe.Pointer(uintptr(v35 + 32))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 910)))
	v36 = v2 + v35
	*(*uint32)(unsafe.Pointer(uintptr(v36 + 12))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 914)))
	*(*uint32)(unsafe.Pointer(uintptr(v36 + 16))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 918)))
	*(*uint32)(unsafe.Pointer(uintptr(v36 + 20))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 922)))
	*(*uint32)(unsafe.Pointer(uintptr(v36 + 24))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 926)))
	*(*uint32)(unsafe.Pointer(uintptr(v36 + 28))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 930)))
	*(*uint16)(unsafe.Pointer(uintptr(v36 + 32))) = *(*uint16)(unsafe.Pointer(uintptr(a1 + 934)))
	*(*uint8)(unsafe.Pointer(uintptr(v36 + 34))) = *(*uint8)(unsafe.Pointer(uintptr(a1 + 936)))
	v37 = v2 + v36
	*(*uint8)(unsafe.Pointer(uintptr(v37 + 13))) = *(*uint8)(unsafe.Pointer(uintptr(a1 + 937)))
	*(*uint16)(unsafe.Pointer(uintptr(v37 + 14))) = *(*uint16)(unsafe.Pointer(uintptr(a1 + 938)))
	*(*uint32)(unsafe.Pointer(uintptr(v37 + 16))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 940)))
	*(*uint32)(unsafe.Pointer(uintptr(v37 + 20))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 944)))
	*(*uint32)(unsafe.Pointer(uintptr(v37 + 24))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 948)))
	*(*uint32)(unsafe.Pointer(uintptr(v37 + 28))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 952)))
	*(*uint16)(unsafe.Pointer(uintptr(v37 + 32))) = *(*uint16)(unsafe.Pointer(uintptr(a1 + 956)))
	v38 = v2 + v37
	*(*uint16)(unsafe.Pointer(uintptr(v38 + 14))) = *(*uint16)(unsafe.Pointer(uintptr(a1 + 958)))
	*(*uint32)(unsafe.Pointer(uintptr(v38 + 16))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 960)))
	*(*uint32)(unsafe.Pointer(uintptr(v38 + 20))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 964)))
	*(*uint32)(unsafe.Pointer(uintptr(v38 + 24))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 968)))
	*(*uint32)(unsafe.Pointer(uintptr(v38 + 28))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 972)))
	*(*uint8)(unsafe.Pointer(uintptr(v38 + 32))) = *(*uint8)(unsafe.Pointer(uintptr(a1 + 976)))
	v39 = v2 + v38
	*(*uint8)(unsafe.Pointer(uintptr(v39 + 15))) = *(*uint8)(unsafe.Pointer(uintptr(a1 + 977)))
	*(*uint32)(unsafe.Pointer(uintptr(v39 + 16))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 978)))
	*(*uint32)(unsafe.Pointer(uintptr(v39 + 20))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 982)))
	*(*uint32)(unsafe.Pointer(uintptr(v39 + 24))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 986)))
	*(*uint32)(unsafe.Pointer(uintptr(v39 + 28))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 990)))
	v40 = v2 + v39
	*(*uint32)(unsafe.Pointer(uintptr(v40 + 16))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 994)))
	*(*uint32)(unsafe.Pointer(uintptr(v40 + 20))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 998)))
	*(*uint32)(unsafe.Pointer(uintptr(v40 + 24))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1002)))
	*(*uint16)(unsafe.Pointer(uintptr(v40 + 28))) = *(*uint16)(unsafe.Pointer(uintptr(a1 + 1006)))
	*(*uint8)(unsafe.Pointer(uintptr(v40 + 30))) = *(*uint8)(unsafe.Pointer(uintptr(a1 + 1008)))
	v41 = v2 + v40
	*(*uint8)(unsafe.Pointer(uintptr(v41 + 17))) = *(*uint8)(unsafe.Pointer(uintptr(a1 + 1009)))
	*(*uint16)(unsafe.Pointer(uintptr(v41 + 18))) = *(*uint16)(unsafe.Pointer(uintptr(a1 + 1010)))
	*(*uint32)(unsafe.Pointer(uintptr(v41 + 20))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1012)))
	*(*uint32)(unsafe.Pointer(uintptr(v41 + 24))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1016)))
	*(*uint16)(unsafe.Pointer(uintptr(v41 + 28))) = *(*uint16)(unsafe.Pointer(uintptr(a1 + 1020)))
	v42 = v2 + v41
	*(*uint16)(unsafe.Pointer(uintptr(v42 + 18))) = *(*uint16)(unsafe.Pointer(uintptr(a1 + 1022)))
	*(*uint32)(unsafe.Pointer(uintptr(v42 + 20))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1024)))
	*(*uint32)(unsafe.Pointer(uintptr(v42 + 24))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1028)))
	*(*uint8)(unsafe.Pointer(uintptr(v42 + 28))) = *(*uint8)(unsafe.Pointer(uintptr(a1 + 1032)))
	v43 = v2 + v42
	*(*uint8)(unsafe.Pointer(uintptr(v43 + 19))) = *(*uint8)(unsafe.Pointer(uintptr(a1 + 1033)))
	*(*uint32)(unsafe.Pointer(uintptr(v43 + 20))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1034)))
	*(*uint32)(unsafe.Pointer(uintptr(v43 + 24))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1038)))
	v44 = v2 + v43
	*(*uint32)(unsafe.Pointer(uintptr(v44 + 20))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1042)))
	*(*uint16)(unsafe.Pointer(uintptr(v44 + 24))) = *(*uint16)(unsafe.Pointer(uintptr(a1 + 1046)))
	*(*uint8)(unsafe.Pointer(uintptr(v44 + 26))) = *(*uint8)(unsafe.Pointer(uintptr(a1 + 1048)))
	v45 = v2 + v44
	*(*uint8)(unsafe.Pointer(uintptr(v45 + 21))) = *(*uint8)(unsafe.Pointer(uintptr(a1 + 1049)))
	*(*uint16)(unsafe.Pointer(uintptr(v45 + 22))) = *(*uint16)(unsafe.Pointer(uintptr(a1 + 1050)))
	*(*uint16)(unsafe.Pointer(uintptr(v45 + 24))) = *(*uint16)(unsafe.Pointer(uintptr(a1 + 1052)))
	v46 = v2 + v45
	*(*uint16)(unsafe.Pointer(uintptr(v46 + 22))) = *(*uint16)(unsafe.Pointer(uintptr(a1 + 1054)))
	result = int8(*(*uint8)(unsafe.Pointer(uintptr(a1 + 1056))))
	*(*uint8)(unsafe.Pointer(uintptr(v46 + 24))) = uint8(result)
	*(*uint8)(unsafe.Pointer(uintptr(v2 + v46 + 23))) = *(*uint8)(unsafe.Pointer(uintptr(a1 + 1057)))
	return result
}
func sub_4831C0(a1 int32, a2 int32) int16 {
	var (
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
		v13    int32
		v14    int32
		v15    int32
		v16    int32
		v17    int32
		v18    int32
		v19    int32
		v20    int32
		v21    int32
		v22    int32
		v23    int32
		v24    int32
		v25    int32
		v26    int32
		v27    int32
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
		v42    int32
		v43    int32
		v44    int32
		v45    int32
		v46    int32
		result int16
	)
	v2 = int32(dword_5d4594_3798804)
	*(*uint16)(unsafe.Pointer(uintptr(a2 + 46))) = *(*uint16)(unsafe.Pointer(uintptr(a1)))
	v3 = v2 + a2
	*(*uint32)(unsafe.Pointer(uintptr(v3 + 44))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 2)))
	*(*uint16)(unsafe.Pointer(uintptr(v3 + 48))) = *(*uint16)(unsafe.Pointer(uintptr(a1 + 6)))
	v4 = v2 + v2 + a2
	*(*uint16)(unsafe.Pointer(uintptr(v4 + 42))) = *(*uint16)(unsafe.Pointer(uintptr(a1 + 8)))
	*(*uint32)(unsafe.Pointer(uintptr(v4 + 44))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 10)))
	*(*uint32)(unsafe.Pointer(uintptr(v4 + 48))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 14)))
	v5 = v2 + v4
	*(*uint32)(unsafe.Pointer(uintptr(v5 + 40))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 18)))
	*(*uint32)(unsafe.Pointer(uintptr(v5 + 44))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 22)))
	*(*uint32)(unsafe.Pointer(uintptr(v5 + 48))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 26)))
	*(*uint16)(unsafe.Pointer(uintptr(v5 + 52))) = *(*uint16)(unsafe.Pointer(uintptr(a1 + 30)))
	v6 = v2 + v5
	*(*uint16)(unsafe.Pointer(uintptr(v6 + 38))) = *(*uint16)(unsafe.Pointer(uintptr(a1 + 32)))
	*(*uint32)(unsafe.Pointer(uintptr(v6 + 40))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 34)))
	*(*uint32)(unsafe.Pointer(uintptr(v6 + 44))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 38)))
	*(*uint32)(unsafe.Pointer(uintptr(v6 + 48))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 42)))
	*(*uint32)(unsafe.Pointer(uintptr(v6 + 52))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 46)))
	v7 = v2 + v6
	*(*uint32)(unsafe.Pointer(uintptr(v7 + 36))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 50)))
	*(*uint32)(unsafe.Pointer(uintptr(v7 + 40))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 54)))
	*(*uint32)(unsafe.Pointer(uintptr(v7 + 44))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 58)))
	*(*uint32)(unsafe.Pointer(uintptr(v7 + 48))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 62)))
	*(*uint32)(unsafe.Pointer(uintptr(v7 + 52))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 66)))
	*(*uint16)(unsafe.Pointer(uintptr(v7 + 56))) = *(*uint16)(unsafe.Pointer(uintptr(a1 + 70)))
	v8 = v2 + v7
	*(*uint16)(unsafe.Pointer(uintptr(v8 + 34))) = *(*uint16)(unsafe.Pointer(uintptr(a1 + 72)))
	*(*uint32)(unsafe.Pointer(uintptr(v8 + 36))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 74)))
	*(*uint32)(unsafe.Pointer(uintptr(v8 + 40))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 78)))
	*(*uint32)(unsafe.Pointer(uintptr(v8 + 44))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 82)))
	*(*uint32)(unsafe.Pointer(uintptr(v8 + 48))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 86)))
	*(*uint32)(unsafe.Pointer(uintptr(v8 + 52))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 90)))
	*(*uint32)(unsafe.Pointer(uintptr(v8 + 56))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 94)))
	v9 = v2 + v8
	*(*uint32)(unsafe.Pointer(uintptr(v9 + 32))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 98)))
	*(*uint32)(unsafe.Pointer(uintptr(v9 + 36))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 102)))
	*(*uint32)(unsafe.Pointer(uintptr(v9 + 40))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 106)))
	*(*uint32)(unsafe.Pointer(uintptr(v9 + 44))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 110)))
	*(*uint32)(unsafe.Pointer(uintptr(v9 + 48))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 114)))
	*(*uint32)(unsafe.Pointer(uintptr(v9 + 52))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 118)))
	*(*uint32)(unsafe.Pointer(uintptr(v9 + 56))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 122)))
	*(*uint16)(unsafe.Pointer(uintptr(v9 + 60))) = *(*uint16)(unsafe.Pointer(uintptr(a1 + 126)))
	v10 = v2 + v9
	*(*uint16)(unsafe.Pointer(uintptr(v10 + 30))) = *(*uint16)(unsafe.Pointer(uintptr(a1 + 128)))
	*(*uint32)(unsafe.Pointer(uintptr(v10 + 32))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 130)))
	*(*uint32)(unsafe.Pointer(uintptr(v10 + 36))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 134)))
	*(*uint32)(unsafe.Pointer(uintptr(v10 + 40))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 138)))
	*(*uint32)(unsafe.Pointer(uintptr(v10 + 44))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 142)))
	*(*uint32)(unsafe.Pointer(uintptr(v10 + 48))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 146)))
	*(*uint32)(unsafe.Pointer(uintptr(v10 + 52))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 150)))
	*(*uint32)(unsafe.Pointer(uintptr(v10 + 56))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 154)))
	*(*uint32)(unsafe.Pointer(uintptr(v10 + 60))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 158)))
	v11 = v2 + v10
	*(*uint32)(unsafe.Pointer(uintptr(v11 + 28))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 162)))
	*(*uint32)(unsafe.Pointer(uintptr(v11 + 32))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 166)))
	*(*uint32)(unsafe.Pointer(uintptr(v11 + 36))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 170)))
	*(*uint32)(unsafe.Pointer(uintptr(v11 + 40))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 174)))
	*(*uint32)(unsafe.Pointer(uintptr(v11 + 44))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 178)))
	*(*uint32)(unsafe.Pointer(uintptr(v11 + 48))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 182)))
	*(*uint32)(unsafe.Pointer(uintptr(v11 + 52))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 186)))
	*(*uint32)(unsafe.Pointer(uintptr(v11 + 56))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 190)))
	*(*uint32)(unsafe.Pointer(uintptr(v11 + 60))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 194)))
	*(*uint16)(unsafe.Pointer(uintptr(v11 + 64))) = *(*uint16)(unsafe.Pointer(uintptr(a1 + 198)))
	v12 = v2 + v11
	*(*uint16)(unsafe.Pointer(uintptr(v12 + 26))) = *(*uint16)(unsafe.Pointer(uintptr(a1 + 200)))
	*(*uint32)(unsafe.Pointer(uintptr(v12 + 28))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 202)))
	*(*uint32)(unsafe.Pointer(uintptr(v12 + 32))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 206)))
	*(*uint32)(unsafe.Pointer(uintptr(v12 + 36))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 210)))
	*(*uint32)(unsafe.Pointer(uintptr(v12 + 40))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 214)))
	*(*uint32)(unsafe.Pointer(uintptr(v12 + 44))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 218)))
	*(*uint32)(unsafe.Pointer(uintptr(v12 + 48))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 222)))
	*(*uint32)(unsafe.Pointer(uintptr(v12 + 52))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 226)))
	*(*uint32)(unsafe.Pointer(uintptr(v12 + 56))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 230)))
	*(*uint32)(unsafe.Pointer(uintptr(v12 + 60))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 234)))
	*(*uint32)(unsafe.Pointer(uintptr(v12 + 64))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 238)))
	v13 = v2 + v12
	*(*uint32)(unsafe.Pointer(uintptr(v13 + 24))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 242)))
	*(*uint32)(unsafe.Pointer(uintptr(v13 + 28))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 246)))
	*(*uint32)(unsafe.Pointer(uintptr(v13 + 32))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 250)))
	*(*uint32)(unsafe.Pointer(uintptr(v13 + 36))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 254)))
	*(*uint32)(unsafe.Pointer(uintptr(v13 + 40))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 258)))
	*(*uint32)(unsafe.Pointer(uintptr(v13 + 44))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 262)))
	*(*uint32)(unsafe.Pointer(uintptr(v13 + 48))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 266)))
	*(*uint32)(unsafe.Pointer(uintptr(v13 + 52))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 270)))
	*(*uint32)(unsafe.Pointer(uintptr(v13 + 56))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 274)))
	*(*uint32)(unsafe.Pointer(uintptr(v13 + 60))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 278)))
	*(*uint32)(unsafe.Pointer(uintptr(v13 + 64))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 282)))
	*(*uint16)(unsafe.Pointer(uintptr(v13 + 68))) = *(*uint16)(unsafe.Pointer(uintptr(a1 + 286)))
	v14 = v2 + v13
	*(*uint16)(unsafe.Pointer(uintptr(v14 + 22))) = *(*uint16)(unsafe.Pointer(uintptr(a1 + 288)))
	*(*uint32)(unsafe.Pointer(uintptr(v14 + 24))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 290)))
	*(*uint32)(unsafe.Pointer(uintptr(v14 + 28))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 294)))
	*(*uint32)(unsafe.Pointer(uintptr(v14 + 32))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 298)))
	*(*uint32)(unsafe.Pointer(uintptr(v14 + 36))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 302)))
	*(*uint32)(unsafe.Pointer(uintptr(v14 + 40))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 306)))
	*(*uint32)(unsafe.Pointer(uintptr(v14 + 44))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 310)))
	*(*uint32)(unsafe.Pointer(uintptr(v14 + 48))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 314)))
	*(*uint32)(unsafe.Pointer(uintptr(v14 + 52))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 318)))
	*(*uint32)(unsafe.Pointer(uintptr(v14 + 56))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 322)))
	*(*uint32)(unsafe.Pointer(uintptr(v14 + 60))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 326)))
	*(*uint32)(unsafe.Pointer(uintptr(v14 + 64))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 330)))
	*(*uint32)(unsafe.Pointer(uintptr(v14 + 68))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 334)))
	v15 = v2 + v14
	*(*uint32)(unsafe.Pointer(uintptr(v15 + 20))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 338)))
	*(*uint32)(unsafe.Pointer(uintptr(v15 + 24))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 342)))
	*(*uint32)(unsafe.Pointer(uintptr(v15 + 28))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 346)))
	*(*uint32)(unsafe.Pointer(uintptr(v15 + 32))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 350)))
	*(*uint32)(unsafe.Pointer(uintptr(v15 + 36))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 354)))
	*(*uint32)(unsafe.Pointer(uintptr(v15 + 40))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 358)))
	*(*uint32)(unsafe.Pointer(uintptr(v15 + 44))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 362)))
	*(*uint32)(unsafe.Pointer(uintptr(v15 + 48))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 366)))
	*(*uint32)(unsafe.Pointer(uintptr(v15 + 52))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 370)))
	*(*uint32)(unsafe.Pointer(uintptr(v15 + 56))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 374)))
	*(*uint32)(unsafe.Pointer(uintptr(v15 + 60))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 378)))
	*(*uint32)(unsafe.Pointer(uintptr(v15 + 64))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 382)))
	*(*uint32)(unsafe.Pointer(uintptr(v15 + 68))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 386)))
	*(*uint16)(unsafe.Pointer(uintptr(v15 + 72))) = *(*uint16)(unsafe.Pointer(uintptr(a1 + 390)))
	v16 = v2 + v15
	*(*uint16)(unsafe.Pointer(uintptr(v16 + 18))) = *(*uint16)(unsafe.Pointer(uintptr(a1 + 392)))
	*(*uint32)(unsafe.Pointer(uintptr(v16 + 20))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 394)))
	*(*uint32)(unsafe.Pointer(uintptr(v16 + 24))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 398)))
	*(*uint32)(unsafe.Pointer(uintptr(v16 + 28))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 402)))
	*(*uint32)(unsafe.Pointer(uintptr(v16 + 32))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 406)))
	*(*uint32)(unsafe.Pointer(uintptr(v16 + 36))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 410)))
	*(*uint32)(unsafe.Pointer(uintptr(v16 + 40))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 414)))
	*(*uint32)(unsafe.Pointer(uintptr(v16 + 44))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 418)))
	*(*uint32)(unsafe.Pointer(uintptr(v16 + 48))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 422)))
	*(*uint32)(unsafe.Pointer(uintptr(v16 + 52))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 426)))
	*(*uint32)(unsafe.Pointer(uintptr(v16 + 56))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 430)))
	*(*uint32)(unsafe.Pointer(uintptr(v16 + 60))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 434)))
	*(*uint32)(unsafe.Pointer(uintptr(v16 + 64))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 438)))
	*(*uint32)(unsafe.Pointer(uintptr(v16 + 68))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 442)))
	*(*uint32)(unsafe.Pointer(uintptr(v16 + 72))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 446)))
	v17 = v2 + v16
	*(*uint32)(unsafe.Pointer(uintptr(v17 + 16))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 450)))
	*(*uint32)(unsafe.Pointer(uintptr(v17 + 20))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 454)))
	*(*uint32)(unsafe.Pointer(uintptr(v17 + 24))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 458)))
	*(*uint32)(unsafe.Pointer(uintptr(v17 + 28))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 462)))
	*(*uint32)(unsafe.Pointer(uintptr(v17 + 32))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 466)))
	*(*uint32)(unsafe.Pointer(uintptr(v17 + 36))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 470)))
	*(*uint32)(unsafe.Pointer(uintptr(v17 + 40))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 474)))
	*(*uint32)(unsafe.Pointer(uintptr(v17 + 44))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 478)))
	*(*uint32)(unsafe.Pointer(uintptr(v17 + 48))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 482)))
	*(*uint32)(unsafe.Pointer(uintptr(v17 + 52))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 486)))
	*(*uint32)(unsafe.Pointer(uintptr(v17 + 56))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 490)))
	*(*uint32)(unsafe.Pointer(uintptr(v17 + 60))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 494)))
	*(*uint32)(unsafe.Pointer(uintptr(v17 + 64))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 498)))
	*(*uint32)(unsafe.Pointer(uintptr(v17 + 68))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 502)))
	*(*uint32)(unsafe.Pointer(uintptr(v17 + 72))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 506)))
	*(*uint16)(unsafe.Pointer(uintptr(v17 + 76))) = *(*uint16)(unsafe.Pointer(uintptr(a1 + 510)))
	v18 = v2 + v17
	*(*uint16)(unsafe.Pointer(uintptr(v18 + 14))) = *(*uint16)(unsafe.Pointer(uintptr(a1 + 512)))
	*(*uint32)(unsafe.Pointer(uintptr(v18 + 16))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 514)))
	*(*uint32)(unsafe.Pointer(uintptr(v18 + 20))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 518)))
	*(*uint32)(unsafe.Pointer(uintptr(v18 + 24))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 522)))
	*(*uint32)(unsafe.Pointer(uintptr(v18 + 28))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 526)))
	*(*uint32)(unsafe.Pointer(uintptr(v18 + 32))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 530)))
	*(*uint32)(unsafe.Pointer(uintptr(v18 + 36))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 534)))
	*(*uint32)(unsafe.Pointer(uintptr(v18 + 40))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 538)))
	*(*uint32)(unsafe.Pointer(uintptr(v18 + 44))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 542)))
	*(*uint32)(unsafe.Pointer(uintptr(v18 + 48))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 546)))
	*(*uint32)(unsafe.Pointer(uintptr(v18 + 52))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 550)))
	*(*uint32)(unsafe.Pointer(uintptr(v18 + 56))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 554)))
	*(*uint32)(unsafe.Pointer(uintptr(v18 + 60))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 558)))
	*(*uint32)(unsafe.Pointer(uintptr(v18 + 64))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 562)))
	*(*uint32)(unsafe.Pointer(uintptr(v18 + 68))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 566)))
	*(*uint32)(unsafe.Pointer(uintptr(v18 + 72))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 570)))
	*(*uint32)(unsafe.Pointer(uintptr(v18 + 76))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 574)))
	v19 = v2 + v18
	*(*uint32)(unsafe.Pointer(uintptr(v19 + 12))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 578)))
	*(*uint32)(unsafe.Pointer(uintptr(v19 + 16))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 582)))
	*(*uint32)(unsafe.Pointer(uintptr(v19 + 20))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 586)))
	*(*uint32)(unsafe.Pointer(uintptr(v19 + 24))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 590)))
	*(*uint32)(unsafe.Pointer(uintptr(v19 + 28))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 594)))
	*(*uint32)(unsafe.Pointer(uintptr(v19 + 32))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 598)))
	*(*uint32)(unsafe.Pointer(uintptr(v19 + 36))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 602)))
	*(*uint32)(unsafe.Pointer(uintptr(v19 + 40))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 606)))
	*(*uint32)(unsafe.Pointer(uintptr(v19 + 44))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 610)))
	*(*uint32)(unsafe.Pointer(uintptr(v19 + 48))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 614)))
	*(*uint32)(unsafe.Pointer(uintptr(v19 + 52))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 618)))
	*(*uint32)(unsafe.Pointer(uintptr(v19 + 56))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 622)))
	*(*uint32)(unsafe.Pointer(uintptr(v19 + 60))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 626)))
	*(*uint32)(unsafe.Pointer(uintptr(v19 + 64))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 630)))
	*(*uint32)(unsafe.Pointer(uintptr(v19 + 68))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 634)))
	*(*uint32)(unsafe.Pointer(uintptr(v19 + 72))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 638)))
	*(*uint32)(unsafe.Pointer(uintptr(v19 + 76))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 642)))
	*(*uint16)(unsafe.Pointer(uintptr(v19 + 80))) = *(*uint16)(unsafe.Pointer(uintptr(a1 + 646)))
	v20 = v2 + v19
	*(*uint16)(unsafe.Pointer(uintptr(v20 + 10))) = *(*uint16)(unsafe.Pointer(uintptr(a1 + 648)))
	*(*uint32)(unsafe.Pointer(uintptr(v20 + 12))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 650)))
	*(*uint32)(unsafe.Pointer(uintptr(v20 + 16))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 654)))
	*(*uint32)(unsafe.Pointer(uintptr(v20 + 20))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 658)))
	*(*uint32)(unsafe.Pointer(uintptr(v20 + 24))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 662)))
	*(*uint32)(unsafe.Pointer(uintptr(v20 + 28))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 666)))
	*(*uint32)(unsafe.Pointer(uintptr(v20 + 32))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 670)))
	*(*uint32)(unsafe.Pointer(uintptr(v20 + 36))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 674)))
	*(*uint32)(unsafe.Pointer(uintptr(v20 + 40))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 678)))
	*(*uint32)(unsafe.Pointer(uintptr(v20 + 44))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 682)))
	*(*uint32)(unsafe.Pointer(uintptr(v20 + 48))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 686)))
	*(*uint32)(unsafe.Pointer(uintptr(v20 + 52))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 690)))
	*(*uint32)(unsafe.Pointer(uintptr(v20 + 56))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 694)))
	*(*uint32)(unsafe.Pointer(uintptr(v20 + 60))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 698)))
	*(*uint32)(unsafe.Pointer(uintptr(v20 + 64))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 702)))
	*(*uint32)(unsafe.Pointer(uintptr(v20 + 68))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 706)))
	*(*uint32)(unsafe.Pointer(uintptr(v20 + 72))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 710)))
	*(*uint32)(unsafe.Pointer(uintptr(v20 + 76))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 714)))
	*(*uint32)(unsafe.Pointer(uintptr(v20 + 80))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 718)))
	v21 = v2 + v20
	*(*uint32)(unsafe.Pointer(uintptr(v21 + 8))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 722)))
	*(*uint32)(unsafe.Pointer(uintptr(v21 + 12))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 726)))
	*(*uint32)(unsafe.Pointer(uintptr(v21 + 16))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 730)))
	*(*uint32)(unsafe.Pointer(uintptr(v21 + 20))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 734)))
	*(*uint32)(unsafe.Pointer(uintptr(v21 + 24))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 738)))
	*(*uint32)(unsafe.Pointer(uintptr(v21 + 28))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 742)))
	*(*uint32)(unsafe.Pointer(uintptr(v21 + 32))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 746)))
	*(*uint32)(unsafe.Pointer(uintptr(v21 + 36))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 750)))
	*(*uint32)(unsafe.Pointer(uintptr(v21 + 40))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 754)))
	*(*uint32)(unsafe.Pointer(uintptr(v21 + 44))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 758)))
	*(*uint32)(unsafe.Pointer(uintptr(v21 + 48))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 762)))
	*(*uint32)(unsafe.Pointer(uintptr(v21 + 52))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 766)))
	*(*uint32)(unsafe.Pointer(uintptr(v21 + 56))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 770)))
	*(*uint32)(unsafe.Pointer(uintptr(v21 + 60))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 774)))
	*(*uint32)(unsafe.Pointer(uintptr(v21 + 64))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 778)))
	*(*uint32)(unsafe.Pointer(uintptr(v21 + 68))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 782)))
	*(*uint32)(unsafe.Pointer(uintptr(v21 + 72))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 786)))
	*(*uint32)(unsafe.Pointer(uintptr(v21 + 76))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 790)))
	*(*uint32)(unsafe.Pointer(uintptr(v21 + 80))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 794)))
	*(*uint16)(unsafe.Pointer(uintptr(v21 + 84))) = *(*uint16)(unsafe.Pointer(uintptr(a1 + 798)))
	v22 = v2 + v21
	*(*uint16)(unsafe.Pointer(uintptr(v22 + 6))) = *(*uint16)(unsafe.Pointer(uintptr(a1 + 800)))
	*(*uint32)(unsafe.Pointer(uintptr(v22 + 8))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 802)))
	*(*uint32)(unsafe.Pointer(uintptr(v22 + 12))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 806)))
	*(*uint32)(unsafe.Pointer(uintptr(v22 + 16))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 810)))
	*(*uint32)(unsafe.Pointer(uintptr(v22 + 20))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 814)))
	*(*uint32)(unsafe.Pointer(uintptr(v22 + 24))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 818)))
	*(*uint32)(unsafe.Pointer(uintptr(v22 + 28))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 822)))
	*(*uint32)(unsafe.Pointer(uintptr(v22 + 32))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 826)))
	*(*uint32)(unsafe.Pointer(uintptr(v22 + 36))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 830)))
	*(*uint32)(unsafe.Pointer(uintptr(v22 + 40))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 834)))
	*(*uint32)(unsafe.Pointer(uintptr(v22 + 44))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 838)))
	*(*uint32)(unsafe.Pointer(uintptr(v22 + 48))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 842)))
	*(*uint32)(unsafe.Pointer(uintptr(v22 + 52))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 846)))
	*(*uint32)(unsafe.Pointer(uintptr(v22 + 56))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 850)))
	*(*uint32)(unsafe.Pointer(uintptr(v22 + 60))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 854)))
	*(*uint32)(unsafe.Pointer(uintptr(v22 + 64))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 858)))
	*(*uint32)(unsafe.Pointer(uintptr(v22 + 68))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 862)))
	*(*uint32)(unsafe.Pointer(uintptr(v22 + 72))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 866)))
	*(*uint32)(unsafe.Pointer(uintptr(v22 + 76))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 870)))
	*(*uint32)(unsafe.Pointer(uintptr(v22 + 80))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 874)))
	*(*uint32)(unsafe.Pointer(uintptr(v22 + 84))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 878)))
	v23 = v2 + v22
	*(*uint32)(unsafe.Pointer(uintptr(v23 + 4))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 882)))
	*(*uint32)(unsafe.Pointer(uintptr(v23 + 8))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 886)))
	*(*uint32)(unsafe.Pointer(uintptr(v23 + 12))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 890)))
	*(*uint32)(unsafe.Pointer(uintptr(v23 + 16))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 894)))
	*(*uint32)(unsafe.Pointer(uintptr(v23 + 20))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 898)))
	*(*uint32)(unsafe.Pointer(uintptr(v23 + 24))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 902)))
	*(*uint32)(unsafe.Pointer(uintptr(v23 + 28))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 906)))
	*(*uint32)(unsafe.Pointer(uintptr(v23 + 32))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 910)))
	*(*uint32)(unsafe.Pointer(uintptr(v23 + 36))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 914)))
	*(*uint32)(unsafe.Pointer(uintptr(v23 + 40))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 918)))
	*(*uint32)(unsafe.Pointer(uintptr(v23 + 44))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 922)))
	*(*uint32)(unsafe.Pointer(uintptr(v23 + 48))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 926)))
	*(*uint32)(unsafe.Pointer(uintptr(v23 + 52))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 930)))
	*(*uint32)(unsafe.Pointer(uintptr(v23 + 56))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 934)))
	*(*uint32)(unsafe.Pointer(uintptr(v23 + 60))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 938)))
	*(*uint32)(unsafe.Pointer(uintptr(v23 + 64))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 942)))
	*(*uint32)(unsafe.Pointer(uintptr(v23 + 68))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 946)))
	*(*uint32)(unsafe.Pointer(uintptr(v23 + 72))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 950)))
	*(*uint32)(unsafe.Pointer(uintptr(v23 + 76))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 954)))
	*(*uint32)(unsafe.Pointer(uintptr(v23 + 80))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 958)))
	*(*uint32)(unsafe.Pointer(uintptr(v23 + 84))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 962)))
	*(*uint16)(unsafe.Pointer(uintptr(v23 + 88))) = *(*uint16)(unsafe.Pointer(uintptr(a1 + 966)))
	v24 = v2 + v23
	*(*uint16)(unsafe.Pointer(uintptr(v24 + 2))) = *(*uint16)(unsafe.Pointer(uintptr(a1 + 968)))
	*(*uint32)(unsafe.Pointer(uintptr(v24 + 4))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 970)))
	*(*uint32)(unsafe.Pointer(uintptr(v24 + 8))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 974)))
	*(*uint32)(unsafe.Pointer(uintptr(v24 + 12))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 978)))
	*(*uint32)(unsafe.Pointer(uintptr(v24 + 16))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 982)))
	*(*uint32)(unsafe.Pointer(uintptr(v24 + 20))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 986)))
	*(*uint32)(unsafe.Pointer(uintptr(v24 + 24))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 990)))
	*(*uint32)(unsafe.Pointer(uintptr(v24 + 28))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 994)))
	*(*uint32)(unsafe.Pointer(uintptr(v24 + 32))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 998)))
	*(*uint32)(unsafe.Pointer(uintptr(v24 + 36))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1002)))
	*(*uint32)(unsafe.Pointer(uintptr(v24 + 40))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1006)))
	*(*uint32)(unsafe.Pointer(uintptr(v24 + 44))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1010)))
	*(*uint32)(unsafe.Pointer(uintptr(v24 + 48))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1014)))
	*(*uint32)(unsafe.Pointer(uintptr(v24 + 52))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1018)))
	*(*uint32)(unsafe.Pointer(uintptr(v24 + 56))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1022)))
	*(*uint32)(unsafe.Pointer(uintptr(v24 + 60))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1026)))
	*(*uint32)(unsafe.Pointer(uintptr(v24 + 64))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1030)))
	*(*uint32)(unsafe.Pointer(uintptr(v24 + 68))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1034)))
	*(*uint32)(unsafe.Pointer(uintptr(v24 + 72))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1038)))
	*(*uint32)(unsafe.Pointer(uintptr(v24 + 76))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1042)))
	*(*uint32)(unsafe.Pointer(uintptr(v24 + 80))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1046)))
	*(*uint32)(unsafe.Pointer(uintptr(v24 + 84))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1050)))
	*(*uint32)(unsafe.Pointer(uintptr(v24 + 88))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1054)))
	v25 = v2 + v24
	*(*uint16)(unsafe.Pointer(uintptr(v25 + 2))) = *(*uint16)(unsafe.Pointer(uintptr(a1 + 1058)))
	*(*uint32)(unsafe.Pointer(uintptr(v25 + 4))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1060)))
	*(*uint32)(unsafe.Pointer(uintptr(v25 + 8))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1064)))
	*(*uint32)(unsafe.Pointer(uintptr(v25 + 12))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1068)))
	*(*uint32)(unsafe.Pointer(uintptr(v25 + 16))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1072)))
	*(*uint32)(unsafe.Pointer(uintptr(v25 + 20))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1076)))
	*(*uint32)(unsafe.Pointer(uintptr(v25 + 24))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1080)))
	*(*uint32)(unsafe.Pointer(uintptr(v25 + 28))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1084)))
	*(*uint32)(unsafe.Pointer(uintptr(v25 + 32))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1088)))
	*(*uint32)(unsafe.Pointer(uintptr(v25 + 36))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1092)))
	*(*uint32)(unsafe.Pointer(uintptr(v25 + 40))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1096)))
	*(*uint32)(unsafe.Pointer(uintptr(v25 + 44))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1100)))
	*(*uint32)(unsafe.Pointer(uintptr(v25 + 48))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1104)))
	*(*uint32)(unsafe.Pointer(uintptr(v25 + 52))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1108)))
	*(*uint32)(unsafe.Pointer(uintptr(v25 + 56))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1112)))
	*(*uint32)(unsafe.Pointer(uintptr(v25 + 60))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1116)))
	*(*uint32)(unsafe.Pointer(uintptr(v25 + 64))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1120)))
	*(*uint32)(unsafe.Pointer(uintptr(v25 + 68))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1124)))
	*(*uint32)(unsafe.Pointer(uintptr(v25 + 72))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1128)))
	*(*uint32)(unsafe.Pointer(uintptr(v25 + 76))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1132)))
	*(*uint32)(unsafe.Pointer(uintptr(v25 + 80))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1136)))
	*(*uint32)(unsafe.Pointer(uintptr(v25 + 84))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1140)))
	*(*uint32)(unsafe.Pointer(uintptr(v25 + 88))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1144)))
	v26 = v2 + v25
	*(*uint32)(unsafe.Pointer(uintptr(v26 + 4))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1148)))
	*(*uint32)(unsafe.Pointer(uintptr(v26 + 8))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1152)))
	*(*uint32)(unsafe.Pointer(uintptr(v26 + 12))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1156)))
	*(*uint32)(unsafe.Pointer(uintptr(v26 + 16))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1160)))
	*(*uint32)(unsafe.Pointer(uintptr(v26 + 20))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1164)))
	*(*uint32)(unsafe.Pointer(uintptr(v26 + 24))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1168)))
	*(*uint32)(unsafe.Pointer(uintptr(v26 + 28))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1172)))
	*(*uint32)(unsafe.Pointer(uintptr(v26 + 32))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1176)))
	*(*uint32)(unsafe.Pointer(uintptr(v26 + 36))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1180)))
	*(*uint32)(unsafe.Pointer(uintptr(v26 + 40))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1184)))
	*(*uint32)(unsafe.Pointer(uintptr(v26 + 44))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1188)))
	*(*uint32)(unsafe.Pointer(uintptr(v26 + 48))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1192)))
	*(*uint32)(unsafe.Pointer(uintptr(v26 + 52))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1196)))
	*(*uint32)(unsafe.Pointer(uintptr(v26 + 56))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1200)))
	*(*uint32)(unsafe.Pointer(uintptr(v26 + 60))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1204)))
	*(*uint32)(unsafe.Pointer(uintptr(v26 + 64))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1208)))
	*(*uint32)(unsafe.Pointer(uintptr(v26 + 68))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1212)))
	*(*uint32)(unsafe.Pointer(uintptr(v26 + 72))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1216)))
	*(*uint32)(unsafe.Pointer(uintptr(v26 + 76))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1220)))
	*(*uint32)(unsafe.Pointer(uintptr(v26 + 80))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1224)))
	*(*uint32)(unsafe.Pointer(uintptr(v26 + 84))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1228)))
	*(*uint16)(unsafe.Pointer(uintptr(v26 + 88))) = *(*uint16)(unsafe.Pointer(uintptr(a1 + 1232)))
	v27 = v2 + v26
	*(*uint16)(unsafe.Pointer(uintptr(v27 + 6))) = *(*uint16)(unsafe.Pointer(uintptr(a1 + 1234)))
	*(*uint32)(unsafe.Pointer(uintptr(v27 + 8))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1236)))
	*(*uint32)(unsafe.Pointer(uintptr(v27 + 12))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1240)))
	*(*uint32)(unsafe.Pointer(uintptr(v27 + 16))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1244)))
	*(*uint32)(unsafe.Pointer(uintptr(v27 + 20))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1248)))
	*(*uint32)(unsafe.Pointer(uintptr(v27 + 24))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1252)))
	*(*uint32)(unsafe.Pointer(uintptr(v27 + 28))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1256)))
	*(*uint32)(unsafe.Pointer(uintptr(v27 + 32))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1260)))
	*(*uint32)(unsafe.Pointer(uintptr(v27 + 36))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1264)))
	*(*uint32)(unsafe.Pointer(uintptr(v27 + 40))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1268)))
	*(*uint32)(unsafe.Pointer(uintptr(v27 + 44))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1272)))
	*(*uint32)(unsafe.Pointer(uintptr(v27 + 48))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1276)))
	*(*uint32)(unsafe.Pointer(uintptr(v27 + 52))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1280)))
	*(*uint32)(unsafe.Pointer(uintptr(v27 + 56))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1284)))
	*(*uint32)(unsafe.Pointer(uintptr(v27 + 60))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1288)))
	*(*uint32)(unsafe.Pointer(uintptr(v27 + 64))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1292)))
	*(*uint32)(unsafe.Pointer(uintptr(v27 + 68))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1296)))
	*(*uint32)(unsafe.Pointer(uintptr(v27 + 72))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1300)))
	*(*uint32)(unsafe.Pointer(uintptr(v27 + 76))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1304)))
	*(*uint32)(unsafe.Pointer(uintptr(v27 + 80))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1308)))
	*(*uint32)(unsafe.Pointer(uintptr(v27 + 84))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1312)))
	v28 = v2 + v27
	*(*uint32)(unsafe.Pointer(uintptr(v28 + 8))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1316)))
	*(*uint32)(unsafe.Pointer(uintptr(v28 + 12))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1320)))
	*(*uint32)(unsafe.Pointer(uintptr(v28 + 16))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1324)))
	*(*uint32)(unsafe.Pointer(uintptr(v28 + 20))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1328)))
	*(*uint32)(unsafe.Pointer(uintptr(v28 + 24))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1332)))
	*(*uint32)(unsafe.Pointer(uintptr(v28 + 28))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1336)))
	*(*uint32)(unsafe.Pointer(uintptr(v28 + 32))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1340)))
	*(*uint32)(unsafe.Pointer(uintptr(v28 + 36))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1344)))
	*(*uint32)(unsafe.Pointer(uintptr(v28 + 40))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1348)))
	*(*uint32)(unsafe.Pointer(uintptr(v28 + 44))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1352)))
	*(*uint32)(unsafe.Pointer(uintptr(v28 + 48))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1356)))
	*(*uint32)(unsafe.Pointer(uintptr(v28 + 52))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1360)))
	*(*uint32)(unsafe.Pointer(uintptr(v28 + 56))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1364)))
	*(*uint32)(unsafe.Pointer(uintptr(v28 + 60))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1368)))
	*(*uint32)(unsafe.Pointer(uintptr(v28 + 64))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1372)))
	*(*uint32)(unsafe.Pointer(uintptr(v28 + 68))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1376)))
	*(*uint32)(unsafe.Pointer(uintptr(v28 + 72))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1380)))
	*(*uint32)(unsafe.Pointer(uintptr(v28 + 76))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1384)))
	*(*uint32)(unsafe.Pointer(uintptr(v28 + 80))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1388)))
	*(*uint16)(unsafe.Pointer(uintptr(v28 + 84))) = *(*uint16)(unsafe.Pointer(uintptr(a1 + 1392)))
	v29 = v2 + v28
	*(*uint16)(unsafe.Pointer(uintptr(v29 + 10))) = *(*uint16)(unsafe.Pointer(uintptr(a1 + 1394)))
	*(*uint32)(unsafe.Pointer(uintptr(v29 + 12))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1396)))
	*(*uint32)(unsafe.Pointer(uintptr(v29 + 16))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1400)))
	*(*uint32)(unsafe.Pointer(uintptr(v29 + 20))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1404)))
	*(*uint32)(unsafe.Pointer(uintptr(v29 + 24))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1408)))
	*(*uint32)(unsafe.Pointer(uintptr(v29 + 28))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1412)))
	*(*uint32)(unsafe.Pointer(uintptr(v29 + 32))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1416)))
	*(*uint32)(unsafe.Pointer(uintptr(v29 + 36))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1420)))
	*(*uint32)(unsafe.Pointer(uintptr(v29 + 40))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1424)))
	*(*uint32)(unsafe.Pointer(uintptr(v29 + 44))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1428)))
	*(*uint32)(unsafe.Pointer(uintptr(v29 + 48))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1432)))
	*(*uint32)(unsafe.Pointer(uintptr(v29 + 52))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1436)))
	*(*uint32)(unsafe.Pointer(uintptr(v29 + 56))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1440)))
	*(*uint32)(unsafe.Pointer(uintptr(v29 + 60))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1444)))
	*(*uint32)(unsafe.Pointer(uintptr(v29 + 64))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1448)))
	*(*uint32)(unsafe.Pointer(uintptr(v29 + 68))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1452)))
	*(*uint32)(unsafe.Pointer(uintptr(v29 + 72))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1456)))
	*(*uint32)(unsafe.Pointer(uintptr(v29 + 76))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1460)))
	*(*uint32)(unsafe.Pointer(uintptr(v29 + 80))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1464)))
	v30 = v2 + v29
	*(*uint32)(unsafe.Pointer(uintptr(v30 + 12))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1468)))
	*(*uint32)(unsafe.Pointer(uintptr(v30 + 16))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1472)))
	*(*uint32)(unsafe.Pointer(uintptr(v30 + 20))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1476)))
	*(*uint32)(unsafe.Pointer(uintptr(v30 + 24))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1480)))
	*(*uint32)(unsafe.Pointer(uintptr(v30 + 28))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1484)))
	*(*uint32)(unsafe.Pointer(uintptr(v30 + 32))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1488)))
	*(*uint32)(unsafe.Pointer(uintptr(v30 + 36))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1492)))
	*(*uint32)(unsafe.Pointer(uintptr(v30 + 40))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1496)))
	*(*uint32)(unsafe.Pointer(uintptr(v30 + 44))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1500)))
	*(*uint32)(unsafe.Pointer(uintptr(v30 + 48))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1504)))
	*(*uint32)(unsafe.Pointer(uintptr(v30 + 52))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1508)))
	*(*uint32)(unsafe.Pointer(uintptr(v30 + 56))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1512)))
	*(*uint32)(unsafe.Pointer(uintptr(v30 + 60))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1516)))
	*(*uint32)(unsafe.Pointer(uintptr(v30 + 64))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1520)))
	*(*uint32)(unsafe.Pointer(uintptr(v30 + 68))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1524)))
	*(*uint32)(unsafe.Pointer(uintptr(v30 + 72))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1528)))
	*(*uint32)(unsafe.Pointer(uintptr(v30 + 76))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1532)))
	*(*uint16)(unsafe.Pointer(uintptr(v30 + 80))) = *(*uint16)(unsafe.Pointer(uintptr(a1 + 1536)))
	v31 = v2 + v30
	*(*uint16)(unsafe.Pointer(uintptr(v31 + 14))) = *(*uint16)(unsafe.Pointer(uintptr(a1 + 1538)))
	*(*uint32)(unsafe.Pointer(uintptr(v31 + 16))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1540)))
	*(*uint32)(unsafe.Pointer(uintptr(v31 + 20))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1544)))
	*(*uint32)(unsafe.Pointer(uintptr(v31 + 24))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1548)))
	*(*uint32)(unsafe.Pointer(uintptr(v31 + 28))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1552)))
	*(*uint32)(unsafe.Pointer(uintptr(v31 + 32))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1556)))
	*(*uint32)(unsafe.Pointer(uintptr(v31 + 36))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1560)))
	*(*uint32)(unsafe.Pointer(uintptr(v31 + 40))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1564)))
	*(*uint32)(unsafe.Pointer(uintptr(v31 + 44))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1568)))
	*(*uint32)(unsafe.Pointer(uintptr(v31 + 48))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1572)))
	*(*uint32)(unsafe.Pointer(uintptr(v31 + 52))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1576)))
	*(*uint32)(unsafe.Pointer(uintptr(v31 + 56))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1580)))
	*(*uint32)(unsafe.Pointer(uintptr(v31 + 60))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1584)))
	*(*uint32)(unsafe.Pointer(uintptr(v31 + 64))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1588)))
	*(*uint32)(unsafe.Pointer(uintptr(v31 + 68))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1592)))
	*(*uint32)(unsafe.Pointer(uintptr(v31 + 72))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1596)))
	*(*uint32)(unsafe.Pointer(uintptr(v31 + 76))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1600)))
	v32 = v2 + v31
	*(*uint32)(unsafe.Pointer(uintptr(v32 + 16))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1604)))
	*(*uint32)(unsafe.Pointer(uintptr(v32 + 20))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1608)))
	*(*uint32)(unsafe.Pointer(uintptr(v32 + 24))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1612)))
	*(*uint32)(unsafe.Pointer(uintptr(v32 + 28))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1616)))
	*(*uint32)(unsafe.Pointer(uintptr(v32 + 32))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1620)))
	*(*uint32)(unsafe.Pointer(uintptr(v32 + 36))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1624)))
	*(*uint32)(unsafe.Pointer(uintptr(v32 + 40))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1628)))
	*(*uint32)(unsafe.Pointer(uintptr(v32 + 44))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1632)))
	*(*uint32)(unsafe.Pointer(uintptr(v32 + 48))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1636)))
	*(*uint32)(unsafe.Pointer(uintptr(v32 + 52))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1640)))
	*(*uint32)(unsafe.Pointer(uintptr(v32 + 56))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1644)))
	*(*uint32)(unsafe.Pointer(uintptr(v32 + 60))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1648)))
	*(*uint32)(unsafe.Pointer(uintptr(v32 + 64))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1652)))
	*(*uint32)(unsafe.Pointer(uintptr(v32 + 68))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1656)))
	*(*uint32)(unsafe.Pointer(uintptr(v32 + 72))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1660)))
	*(*uint16)(unsafe.Pointer(uintptr(v32 + 76))) = *(*uint16)(unsafe.Pointer(uintptr(a1 + 1664)))
	v33 = v2 + v32
	*(*uint16)(unsafe.Pointer(uintptr(v33 + 18))) = *(*uint16)(unsafe.Pointer(uintptr(a1 + 1666)))
	*(*uint32)(unsafe.Pointer(uintptr(v33 + 20))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1668)))
	*(*uint32)(unsafe.Pointer(uintptr(v33 + 24))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1672)))
	*(*uint32)(unsafe.Pointer(uintptr(v33 + 28))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1676)))
	*(*uint32)(unsafe.Pointer(uintptr(v33 + 32))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1680)))
	*(*uint32)(unsafe.Pointer(uintptr(v33 + 36))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1684)))
	*(*uint32)(unsafe.Pointer(uintptr(v33 + 40))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1688)))
	*(*uint32)(unsafe.Pointer(uintptr(v33 + 44))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1692)))
	*(*uint32)(unsafe.Pointer(uintptr(v33 + 48))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1696)))
	*(*uint32)(unsafe.Pointer(uintptr(v33 + 52))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1700)))
	*(*uint32)(unsafe.Pointer(uintptr(v33 + 56))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1704)))
	*(*uint32)(unsafe.Pointer(uintptr(v33 + 60))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1708)))
	*(*uint32)(unsafe.Pointer(uintptr(v33 + 64))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1712)))
	*(*uint32)(unsafe.Pointer(uintptr(v33 + 68))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1716)))
	*(*uint32)(unsafe.Pointer(uintptr(v33 + 72))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1720)))
	v34 = v2 + v33
	*(*uint32)(unsafe.Pointer(uintptr(v34 + 20))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1724)))
	*(*uint32)(unsafe.Pointer(uintptr(v34 + 24))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1728)))
	*(*uint32)(unsafe.Pointer(uintptr(v34 + 28))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1732)))
	*(*uint32)(unsafe.Pointer(uintptr(v34 + 32))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1736)))
	*(*uint32)(unsafe.Pointer(uintptr(v34 + 36))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1740)))
	*(*uint32)(unsafe.Pointer(uintptr(v34 + 40))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1744)))
	*(*uint32)(unsafe.Pointer(uintptr(v34 + 44))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1748)))
	*(*uint32)(unsafe.Pointer(uintptr(v34 + 48))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1752)))
	*(*uint32)(unsafe.Pointer(uintptr(v34 + 52))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1756)))
	*(*uint32)(unsafe.Pointer(uintptr(v34 + 56))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1760)))
	*(*uint32)(unsafe.Pointer(uintptr(v34 + 60))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1764)))
	*(*uint32)(unsafe.Pointer(uintptr(v34 + 64))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1768)))
	*(*uint32)(unsafe.Pointer(uintptr(v34 + 68))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1772)))
	*(*uint16)(unsafe.Pointer(uintptr(v34 + 72))) = *(*uint16)(unsafe.Pointer(uintptr(a1 + 1776)))
	v35 = v2 + v34
	*(*uint16)(unsafe.Pointer(uintptr(v35 + 22))) = *(*uint16)(unsafe.Pointer(uintptr(a1 + 1778)))
	*(*uint32)(unsafe.Pointer(uintptr(v35 + 24))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1780)))
	*(*uint32)(unsafe.Pointer(uintptr(v35 + 28))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1784)))
	*(*uint32)(unsafe.Pointer(uintptr(v35 + 32))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1788)))
	*(*uint32)(unsafe.Pointer(uintptr(v35 + 36))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1792)))
	*(*uint32)(unsafe.Pointer(uintptr(v35 + 40))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1796)))
	*(*uint32)(unsafe.Pointer(uintptr(v35 + 44))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1800)))
	*(*uint32)(unsafe.Pointer(uintptr(v35 + 48))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1804)))
	*(*uint32)(unsafe.Pointer(uintptr(v35 + 52))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1808)))
	*(*uint32)(unsafe.Pointer(uintptr(v35 + 56))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1812)))
	*(*uint32)(unsafe.Pointer(uintptr(v35 + 60))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1816)))
	*(*uint32)(unsafe.Pointer(uintptr(v35 + 64))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1820)))
	*(*uint32)(unsafe.Pointer(uintptr(v35 + 68))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1824)))
	v36 = v2 + v35
	*(*uint32)(unsafe.Pointer(uintptr(v36 + 24))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1828)))
	*(*uint32)(unsafe.Pointer(uintptr(v36 + 28))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1832)))
	*(*uint32)(unsafe.Pointer(uintptr(v36 + 32))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1836)))
	*(*uint32)(unsafe.Pointer(uintptr(v36 + 36))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1840)))
	*(*uint32)(unsafe.Pointer(uintptr(v36 + 40))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1844)))
	*(*uint32)(unsafe.Pointer(uintptr(v36 + 44))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1848)))
	*(*uint32)(unsafe.Pointer(uintptr(v36 + 48))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1852)))
	*(*uint32)(unsafe.Pointer(uintptr(v36 + 52))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1856)))
	*(*uint32)(unsafe.Pointer(uintptr(v36 + 56))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1860)))
	*(*uint32)(unsafe.Pointer(uintptr(v36 + 60))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1864)))
	*(*uint32)(unsafe.Pointer(uintptr(v36 + 64))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1868)))
	*(*uint16)(unsafe.Pointer(uintptr(v36 + 68))) = *(*uint16)(unsafe.Pointer(uintptr(a1 + 1872)))
	v37 = v2 + v36
	*(*uint16)(unsafe.Pointer(uintptr(v37 + 26))) = *(*uint16)(unsafe.Pointer(uintptr(a1 + 1874)))
	*(*uint32)(unsafe.Pointer(uintptr(v37 + 28))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1876)))
	*(*uint32)(unsafe.Pointer(uintptr(v37 + 32))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1880)))
	*(*uint32)(unsafe.Pointer(uintptr(v37 + 36))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1884)))
	*(*uint32)(unsafe.Pointer(uintptr(v37 + 40))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1888)))
	*(*uint32)(unsafe.Pointer(uintptr(v37 + 44))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1892)))
	*(*uint32)(unsafe.Pointer(uintptr(v37 + 48))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1896)))
	*(*uint32)(unsafe.Pointer(uintptr(v37 + 52))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1900)))
	*(*uint32)(unsafe.Pointer(uintptr(v37 + 56))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1904)))
	*(*uint32)(unsafe.Pointer(uintptr(v37 + 60))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1908)))
	*(*uint32)(unsafe.Pointer(uintptr(v37 + 64))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1912)))
	v38 = v2 + v37
	*(*uint32)(unsafe.Pointer(uintptr(v38 + 28))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1916)))
	*(*uint32)(unsafe.Pointer(uintptr(v38 + 32))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1920)))
	*(*uint32)(unsafe.Pointer(uintptr(v38 + 36))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1924)))
	*(*uint32)(unsafe.Pointer(uintptr(v38 + 40))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1928)))
	*(*uint32)(unsafe.Pointer(uintptr(v38 + 44))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1932)))
	*(*uint32)(unsafe.Pointer(uintptr(v38 + 48))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1936)))
	*(*uint32)(unsafe.Pointer(uintptr(v38 + 52))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1940)))
	*(*uint32)(unsafe.Pointer(uintptr(v38 + 56))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1944)))
	*(*uint32)(unsafe.Pointer(uintptr(v38 + 60))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1948)))
	*(*uint16)(unsafe.Pointer(uintptr(v38 + 64))) = *(*uint16)(unsafe.Pointer(uintptr(a1 + 1952)))
	v39 = v2 + v38
	*(*uint16)(unsafe.Pointer(uintptr(v39 + 30))) = *(*uint16)(unsafe.Pointer(uintptr(a1 + 1954)))
	*(*uint32)(unsafe.Pointer(uintptr(v39 + 32))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1956)))
	*(*uint32)(unsafe.Pointer(uintptr(v39 + 36))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1960)))
	*(*uint32)(unsafe.Pointer(uintptr(v39 + 40))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1964)))
	*(*uint32)(unsafe.Pointer(uintptr(v39 + 44))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1968)))
	*(*uint32)(unsafe.Pointer(uintptr(v39 + 48))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1972)))
	*(*uint32)(unsafe.Pointer(uintptr(v39 + 52))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1976)))
	*(*uint32)(unsafe.Pointer(uintptr(v39 + 56))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1980)))
	*(*uint32)(unsafe.Pointer(uintptr(v39 + 60))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1984)))
	v40 = v2 + v39
	*(*uint32)(unsafe.Pointer(uintptr(v40 + 32))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1988)))
	*(*uint32)(unsafe.Pointer(uintptr(v40 + 36))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1992)))
	*(*uint32)(unsafe.Pointer(uintptr(v40 + 40))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 1996)))
	*(*uint32)(unsafe.Pointer(uintptr(v40 + 44))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 2000)))
	*(*uint32)(unsafe.Pointer(uintptr(v40 + 48))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 2004)))
	*(*uint32)(unsafe.Pointer(uintptr(v40 + 52))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 2008)))
	*(*uint32)(unsafe.Pointer(uintptr(v40 + 56))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 2012)))
	*(*uint16)(unsafe.Pointer(uintptr(v40 + 60))) = *(*uint16)(unsafe.Pointer(uintptr(a1 + 2016)))
	v41 = v2 + v40
	*(*uint16)(unsafe.Pointer(uintptr(v41 + 34))) = *(*uint16)(unsafe.Pointer(uintptr(a1 + 2018)))
	*(*uint32)(unsafe.Pointer(uintptr(v41 + 36))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 2020)))
	*(*uint32)(unsafe.Pointer(uintptr(v41 + 40))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 2024)))
	*(*uint32)(unsafe.Pointer(uintptr(v41 + 44))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 2028)))
	*(*uint32)(unsafe.Pointer(uintptr(v41 + 48))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 2032)))
	*(*uint32)(unsafe.Pointer(uintptr(v41 + 52))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 2036)))
	*(*uint32)(unsafe.Pointer(uintptr(v41 + 56))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 2040)))
	v42 = v2 + v41
	*(*uint32)(unsafe.Pointer(uintptr(v42 + 36))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 2044)))
	*(*uint32)(unsafe.Pointer(uintptr(v42 + 40))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 2048)))
	*(*uint32)(unsafe.Pointer(uintptr(v42 + 44))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 2052)))
	*(*uint32)(unsafe.Pointer(uintptr(v42 + 48))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 2056)))
	*(*uint32)(unsafe.Pointer(uintptr(v42 + 52))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 2060)))
	*(*uint16)(unsafe.Pointer(uintptr(v42 + 56))) = *(*uint16)(unsafe.Pointer(uintptr(a1 + 2064)))
	v43 = v2 + v42
	*(*uint16)(unsafe.Pointer(uintptr(v43 + 38))) = *(*uint16)(unsafe.Pointer(uintptr(a1 + 2066)))
	*(*uint32)(unsafe.Pointer(uintptr(v43 + 40))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 2068)))
	*(*uint32)(unsafe.Pointer(uintptr(v43 + 44))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 2072)))
	*(*uint32)(unsafe.Pointer(uintptr(v43 + 48))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 2076)))
	*(*uint32)(unsafe.Pointer(uintptr(v43 + 52))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 2080)))
	v44 = v2 + v43
	*(*uint32)(unsafe.Pointer(uintptr(v44 + 40))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 2084)))
	*(*uint32)(unsafe.Pointer(uintptr(v44 + 44))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 2088)))
	*(*uint32)(unsafe.Pointer(uintptr(v44 + 48))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 2092)))
	*(*uint16)(unsafe.Pointer(uintptr(v44 + 52))) = *(*uint16)(unsafe.Pointer(uintptr(a1 + 2096)))
	v45 = v2 + v44
	*(*uint16)(unsafe.Pointer(uintptr(v45 + 42))) = *(*uint16)(unsafe.Pointer(uintptr(a1 + 2098)))
	*(*uint32)(unsafe.Pointer(uintptr(v45 + 44))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 2100)))
	*(*uint32)(unsafe.Pointer(uintptr(v45 + 48))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 2104)))
	v46 = v2 + v45
	*(*uint32)(unsafe.Pointer(uintptr(v46 + 44))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 2108)))
	result = int16(*(*uint16)(unsafe.Pointer(uintptr(a1 + 2112))))
	*(*uint16)(unsafe.Pointer(uintptr(v46 + 48))) = uint16(result)
	*(*uint16)(unsafe.Pointer(uintptr(v2 + v46 + 46))) = *(*uint16)(unsafe.Pointer(uintptr(a1 + 2114)))
	return result
}
func sub_483FE0(a1 int32, a2 int32) int32 {
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
		v14    int32
		v15    int32
		v16    int32
		v17    int32
		v18    int32
		v19    int32
		v20    int32
		v21    int32
		v22    int32
		v23    int32
		v24    int32
		v25    int32
		v26    int32
		v27    int32
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
		v42    int32
		v43    int32
		v44    int32
		v45    int32
		v46    int32
		v47    int32
	)
	result = a1
	v3 = int32(dword_5d4594_3798804)
	*(*uint8)(unsafe.Pointer(uintptr(a2 + 23))) = uint8(int8(a1))
	v4 = v3 + a2
	*(*uint16)(unsafe.Pointer(uintptr(v4 + 22))) = uint16(int16(a1))
	*(*uint8)(unsafe.Pointer(uintptr(v4 + 24))) = uint8(int8(a1))
	v5 = v3 + v3 + a2
	*(*uint8)(unsafe.Pointer(uintptr(v5 + 21))) = uint8(int8(a1))
	*(*uint16)(unsafe.Pointer(uintptr(v5 + 22))) = uint16(int16(a1))
	*(*uint16)(unsafe.Pointer(uintptr(v5 + 24))) = uint16(int16(a1))
	v6 = v3 + v5
	*(*uint32)(unsafe.Pointer(uintptr(v6 + 20))) = uint32(a1)
	*(*uint16)(unsafe.Pointer(uintptr(v6 + 24))) = uint16(int16(a1))
	*(*uint8)(unsafe.Pointer(uintptr(v6 + 26))) = uint8(int8(a1))
	v7 = v3 + v6
	*(*uint8)(unsafe.Pointer(uintptr(v7 + 19))) = uint8(int8(a1))
	*(*uint32)(unsafe.Pointer(uintptr(v7 + 20))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v7 + 24))) = uint32(a1)
	v8 = v3 + v7
	*(*uint16)(unsafe.Pointer(uintptr(v8 + 18))) = uint16(int16(a1))
	*(*uint32)(unsafe.Pointer(uintptr(v8 + 20))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v8 + 24))) = uint32(a1)
	*(*uint8)(unsafe.Pointer(uintptr(v8 + 28))) = uint8(int8(a1))
	v9 = v3 + v8
	*(*uint8)(unsafe.Pointer(uintptr(v9 + 17))) = uint8(int8(a1))
	*(*uint16)(unsafe.Pointer(uintptr(v9 + 18))) = uint16(int16(a1))
	*(*uint32)(unsafe.Pointer(uintptr(v9 + 20))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v9 + 24))) = uint32(a1)
	*(*uint16)(unsafe.Pointer(uintptr(v9 + 28))) = uint16(int16(a1))
	v10 = v3 + v9
	*(*uint32)(unsafe.Pointer(uintptr(v10 + 16))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v10 + 20))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v10 + 24))) = uint32(a1)
	*(*uint16)(unsafe.Pointer(uintptr(v10 + 28))) = uint16(int16(a1))
	*(*uint8)(unsafe.Pointer(uintptr(v10 + 30))) = uint8(int8(a1))
	v11 = v3 + v10
	*(*uint8)(unsafe.Pointer(uintptr(v11 + 15))) = uint8(int8(a1))
	*(*uint32)(unsafe.Pointer(uintptr(v11 + 16))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v11 + 20))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v11 + 24))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v11 + 28))) = uint32(a1)
	v12 = v3 + v11
	*(*uint16)(unsafe.Pointer(uintptr(v12 + 14))) = uint16(int16(a1))
	*(*uint32)(unsafe.Pointer(uintptr(v12 + 16))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v12 + 20))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v12 + 24))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v12 + 28))) = uint32(a1)
	*(*uint8)(unsafe.Pointer(uintptr(v12 + 32))) = uint8(int8(a1))
	v13 = v3 + v12
	*(*uint8)(unsafe.Pointer(uintptr(v13 + 13))) = uint8(int8(a1))
	*(*uint16)(unsafe.Pointer(uintptr(v13 + 14))) = uint16(int16(a1))
	*(*uint32)(unsafe.Pointer(uintptr(v13 + 16))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v13 + 20))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v13 + 24))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v13 + 28))) = uint32(a1)
	*(*uint16)(unsafe.Pointer(uintptr(v13 + 32))) = uint16(int16(a1))
	v14 = v3 + v13
	*(*uint32)(unsafe.Pointer(uintptr(v14 + 12))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v14 + 16))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v14 + 20))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v14 + 24))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v14 + 28))) = uint32(a1)
	*(*uint16)(unsafe.Pointer(uintptr(v14 + 32))) = uint16(int16(a1))
	*(*uint8)(unsafe.Pointer(uintptr(v14 + 34))) = uint8(int8(a1))
	v15 = v3 + v14
	*(*uint8)(unsafe.Pointer(uintptr(v15 + 11))) = uint8(int8(a1))
	*(*uint32)(unsafe.Pointer(uintptr(v15 + 12))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v15 + 16))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v15 + 20))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v15 + 24))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v15 + 28))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v15 + 32))) = uint32(a1)
	v16 = v3 + v15
	*(*uint16)(unsafe.Pointer(uintptr(v16 + 10))) = uint16(int16(a1))
	*(*uint32)(unsafe.Pointer(uintptr(v16 + 12))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v16 + 16))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v16 + 20))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v16 + 24))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v16 + 28))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v16 + 32))) = uint32(a1)
	*(*uint8)(unsafe.Pointer(uintptr(v16 + 36))) = uint8(int8(a1))
	v17 = v3 + v16
	*(*uint8)(unsafe.Pointer(uintptr(v17 + 9))) = uint8(int8(a1))
	*(*uint16)(unsafe.Pointer(uintptr(v17 + 10))) = uint16(int16(a1))
	*(*uint32)(unsafe.Pointer(uintptr(v17 + 12))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v17 + 16))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v17 + 20))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v17 + 24))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v17 + 28))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v17 + 32))) = uint32(a1)
	*(*uint16)(unsafe.Pointer(uintptr(v17 + 36))) = uint16(int16(a1))
	v18 = v3 + v17
	*(*uint32)(unsafe.Pointer(uintptr(v18 + 8))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v18 + 12))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v18 + 16))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v18 + 20))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v18 + 24))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v18 + 28))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v18 + 32))) = uint32(a1)
	*(*uint16)(unsafe.Pointer(uintptr(v18 + 36))) = uint16(int16(a1))
	*(*uint8)(unsafe.Pointer(uintptr(v18 + 38))) = uint8(int8(a1))
	v19 = v3 + v18
	*(*uint8)(unsafe.Pointer(uintptr(v19 + 7))) = uint8(int8(a1))
	*(*uint32)(unsafe.Pointer(uintptr(v19 + 8))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v19 + 12))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v19 + 16))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v19 + 20))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v19 + 24))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v19 + 28))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v19 + 32))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v19 + 36))) = uint32(a1)
	v20 = v3 + v19
	*(*uint16)(unsafe.Pointer(uintptr(v20 + 6))) = uint16(int16(a1))
	*(*uint32)(unsafe.Pointer(uintptr(v20 + 8))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v20 + 12))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v20 + 16))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v20 + 20))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v20 + 24))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v20 + 28))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v20 + 32))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v20 + 36))) = uint32(a1)
	*(*uint8)(unsafe.Pointer(uintptr(v20 + 40))) = uint8(int8(a1))
	v21 = v3 + v20
	*(*uint8)(unsafe.Pointer(uintptr(v21 + 5))) = uint8(int8(a1))
	*(*uint16)(unsafe.Pointer(uintptr(v21 + 6))) = uint16(int16(a1))
	*(*uint32)(unsafe.Pointer(uintptr(v21 + 8))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v21 + 12))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v21 + 16))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v21 + 20))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v21 + 24))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v21 + 28))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v21 + 32))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v21 + 36))) = uint32(a1)
	*(*uint16)(unsafe.Pointer(uintptr(v21 + 40))) = uint16(int16(a1))
	v22 = v3 + v21
	*(*uint32)(unsafe.Pointer(uintptr(v22 + 4))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v22 + 8))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v22 + 12))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v22 + 16))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v22 + 20))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v22 + 24))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v22 + 28))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v22 + 32))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v22 + 36))) = uint32(a1)
	*(*uint16)(unsafe.Pointer(uintptr(v22 + 40))) = uint16(int16(a1))
	*(*uint8)(unsafe.Pointer(uintptr(v22 + 42))) = uint8(int8(a1))
	v23 = v3 + v22
	*(*uint8)(unsafe.Pointer(uintptr(v23 + 3))) = uint8(int8(a1))
	*(*uint32)(unsafe.Pointer(uintptr(v23 + 4))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v23 + 8))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v23 + 12))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v23 + 16))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v23 + 20))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v23 + 24))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v23 + 28))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v23 + 32))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v23 + 36))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v23 + 40))) = uint32(a1)
	v24 = v3 + v23
	*(*uint16)(unsafe.Pointer(uintptr(v24 + 2))) = uint16(int16(a1))
	*(*uint32)(unsafe.Pointer(uintptr(v24 + 4))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v24 + 8))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v24 + 12))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v24 + 16))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v24 + 20))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v24 + 24))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v24 + 28))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v24 + 32))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v24 + 36))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v24 + 40))) = uint32(a1)
	*(*uint8)(unsafe.Pointer(uintptr(v24 + 44))) = uint8(int8(a1))
	v25 = v3 + v24
	*(*uint8)(unsafe.Pointer(uintptr(v25 + 1))) = uint8(int8(a1))
	*(*uint16)(unsafe.Pointer(uintptr(v25 + 2))) = uint16(int16(a1))
	*(*uint32)(unsafe.Pointer(uintptr(v25 + 4))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v25 + 8))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v25 + 12))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v25 + 16))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v25 + 20))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v25 + 24))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v25 + 28))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v25 + 32))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v25 + 36))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v25 + 40))) = uint32(a1)
	*(*uint16)(unsafe.Pointer(uintptr(v25 + 44))) = uint16(int16(a1))
	v26 = v3 + v25
	*(*uint8)(unsafe.Pointer(uintptr(v26 + 1))) = uint8(int8(a1))
	*(*uint16)(unsafe.Pointer(uintptr(v26 + 2))) = uint16(int16(a1))
	*(*uint32)(unsafe.Pointer(uintptr(v26 + 4))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v26 + 8))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v26 + 12))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v26 + 16))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v26 + 20))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v26 + 24))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v26 + 28))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v26 + 32))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v26 + 36))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v26 + 40))) = uint32(a1)
	*(*uint16)(unsafe.Pointer(uintptr(v26 + 44))) = uint16(int16(a1))
	v27 = v3 + v26
	*(*uint16)(unsafe.Pointer(uintptr(v27 + 2))) = uint16(int16(a1))
	*(*uint32)(unsafe.Pointer(uintptr(v27 + 4))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v27 + 8))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v27 + 12))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v27 + 16))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v27 + 20))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v27 + 24))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v27 + 28))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v27 + 32))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v27 + 36))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v27 + 40))) = uint32(a1)
	*(*uint8)(unsafe.Pointer(uintptr(v27 + 44))) = uint8(int8(a1))
	v28 = v3 + v27
	*(*uint8)(unsafe.Pointer(uintptr(v28 + 3))) = uint8(int8(a1))
	*(*uint32)(unsafe.Pointer(uintptr(v28 + 4))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v28 + 8))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v28 + 12))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v28 + 16))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v28 + 20))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v28 + 24))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v28 + 28))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v28 + 32))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v28 + 36))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v28 + 40))) = uint32(a1)
	v29 = v3 + v28
	*(*uint32)(unsafe.Pointer(uintptr(v29 + 4))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v29 + 8))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v29 + 12))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v29 + 16))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v29 + 20))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v29 + 24))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v29 + 28))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v29 + 32))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v29 + 36))) = uint32(a1)
	*(*uint16)(unsafe.Pointer(uintptr(v29 + 40))) = uint16(int16(a1))
	*(*uint8)(unsafe.Pointer(uintptr(v29 + 42))) = uint8(int8(a1))
	v30 = v3 + v29
	*(*uint8)(unsafe.Pointer(uintptr(v30 + 5))) = uint8(int8(a1))
	*(*uint16)(unsafe.Pointer(uintptr(v30 + 6))) = uint16(int16(a1))
	*(*uint32)(unsafe.Pointer(uintptr(v30 + 8))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v30 + 12))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v30 + 16))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v30 + 20))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v30 + 24))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v30 + 28))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v30 + 32))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v30 + 36))) = uint32(a1)
	*(*uint16)(unsafe.Pointer(uintptr(v30 + 40))) = uint16(int16(a1))
	v31 = v3 + v30
	*(*uint16)(unsafe.Pointer(uintptr(v31 + 6))) = uint16(int16(a1))
	*(*uint32)(unsafe.Pointer(uintptr(v31 + 8))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v31 + 12))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v31 + 16))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v31 + 20))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v31 + 24))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v31 + 28))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v31 + 32))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v31 + 36))) = uint32(a1)
	*(*uint8)(unsafe.Pointer(uintptr(v31 + 40))) = uint8(int8(a1))
	v32 = v3 + v31
	*(*uint8)(unsafe.Pointer(uintptr(v32 + 7))) = uint8(int8(a1))
	*(*uint32)(unsafe.Pointer(uintptr(v32 + 8))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v32 + 12))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v32 + 16))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v32 + 20))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v32 + 24))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v32 + 28))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v32 + 32))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v32 + 36))) = uint32(a1)
	v33 = v3 + v32
	*(*uint32)(unsafe.Pointer(uintptr(v33 + 8))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v33 + 12))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v33 + 16))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v33 + 20))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v33 + 24))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v33 + 28))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v33 + 32))) = uint32(a1)
	*(*uint16)(unsafe.Pointer(uintptr(v33 + 36))) = uint16(int16(a1))
	*(*uint8)(unsafe.Pointer(uintptr(v33 + 38))) = uint8(int8(a1))
	v34 = v3 + v33
	*(*uint8)(unsafe.Pointer(uintptr(v34 + 9))) = uint8(int8(a1))
	*(*uint16)(unsafe.Pointer(uintptr(v34 + 10))) = uint16(int16(a1))
	*(*uint32)(unsafe.Pointer(uintptr(v34 + 12))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v34 + 16))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v34 + 20))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v34 + 24))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v34 + 28))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v34 + 32))) = uint32(a1)
	*(*uint16)(unsafe.Pointer(uintptr(v34 + 36))) = uint16(int16(a1))
	v35 = v3 + v34
	*(*uint16)(unsafe.Pointer(uintptr(v35 + 10))) = uint16(int16(a1))
	*(*uint32)(unsafe.Pointer(uintptr(v35 + 12))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v35 + 16))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v35 + 20))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v35 + 24))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v35 + 28))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v35 + 32))) = uint32(a1)
	*(*uint8)(unsafe.Pointer(uintptr(v35 + 36))) = uint8(int8(a1))
	v36 = v3 + v35
	*(*uint8)(unsafe.Pointer(uintptr(v36 + 11))) = uint8(int8(a1))
	*(*uint32)(unsafe.Pointer(uintptr(v36 + 12))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v36 + 16))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v36 + 20))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v36 + 24))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v36 + 28))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v36 + 32))) = uint32(a1)
	v37 = v3 + v36
	*(*uint32)(unsafe.Pointer(uintptr(v37 + 12))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v37 + 16))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v37 + 20))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v37 + 24))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v37 + 28))) = uint32(a1)
	*(*uint16)(unsafe.Pointer(uintptr(v37 + 32))) = uint16(int16(a1))
	*(*uint8)(unsafe.Pointer(uintptr(v37 + 34))) = uint8(int8(a1))
	v38 = v3 + v37
	*(*uint8)(unsafe.Pointer(uintptr(v38 + 13))) = uint8(int8(a1))
	*(*uint16)(unsafe.Pointer(uintptr(v38 + 14))) = uint16(int16(a1))
	*(*uint32)(unsafe.Pointer(uintptr(v38 + 16))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v38 + 20))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v38 + 24))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v38 + 28))) = uint32(a1)
	*(*uint16)(unsafe.Pointer(uintptr(v38 + 32))) = uint16(int16(a1))
	v39 = v3 + v38
	*(*uint16)(unsafe.Pointer(uintptr(v39 + 14))) = uint16(int16(a1))
	*(*uint32)(unsafe.Pointer(uintptr(v39 + 16))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v39 + 20))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v39 + 24))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v39 + 28))) = uint32(a1)
	*(*uint8)(unsafe.Pointer(uintptr(v39 + 32))) = uint8(int8(a1))
	v40 = v3 + v39
	*(*uint8)(unsafe.Pointer(uintptr(v40 + 15))) = uint8(int8(a1))
	*(*uint32)(unsafe.Pointer(uintptr(v40 + 16))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v40 + 20))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v40 + 24))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v40 + 28))) = uint32(a1)
	v41 = v3 + v40
	*(*uint32)(unsafe.Pointer(uintptr(v41 + 16))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v41 + 20))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v41 + 24))) = uint32(a1)
	*(*uint16)(unsafe.Pointer(uintptr(v41 + 28))) = uint16(int16(a1))
	*(*uint8)(unsafe.Pointer(uintptr(v41 + 30))) = uint8(int8(a1))
	v42 = v3 + v41
	*(*uint8)(unsafe.Pointer(uintptr(v42 + 17))) = uint8(int8(a1))
	*(*uint16)(unsafe.Pointer(uintptr(v42 + 18))) = uint16(int16(a1))
	*(*uint32)(unsafe.Pointer(uintptr(v42 + 20))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v42 + 24))) = uint32(a1)
	*(*uint16)(unsafe.Pointer(uintptr(v42 + 28))) = uint16(int16(a1))
	v43 = v3 + v42
	*(*uint16)(unsafe.Pointer(uintptr(v43 + 18))) = uint16(int16(a1))
	*(*uint32)(unsafe.Pointer(uintptr(v43 + 20))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v43 + 24))) = uint32(a1)
	*(*uint8)(unsafe.Pointer(uintptr(v43 + 28))) = uint8(int8(a1))
	v44 = v3 + v43
	*(*uint8)(unsafe.Pointer(uintptr(v44 + 19))) = uint8(int8(a1))
	*(*uint32)(unsafe.Pointer(uintptr(v44 + 20))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v44 + 24))) = uint32(a1)
	v45 = v3 + v44
	*(*uint32)(unsafe.Pointer(uintptr(v45 + 20))) = uint32(a1)
	*(*uint16)(unsafe.Pointer(uintptr(v45 + 24))) = uint16(int16(a1))
	*(*uint8)(unsafe.Pointer(uintptr(v45 + 26))) = uint8(int8(a1))
	v46 = v3 + v45
	*(*uint8)(unsafe.Pointer(uintptr(v46 + 21))) = uint8(int8(a1))
	*(*uint16)(unsafe.Pointer(uintptr(v46 + 22))) = uint16(int16(a1))
	*(*uint16)(unsafe.Pointer(uintptr(v46 + 24))) = uint16(int16(a1))
	v47 = v3 + v46
	*(*uint16)(unsafe.Pointer(uintptr(v47 + 22))) = uint16(int16(a1))
	*(*uint8)(unsafe.Pointer(uintptr(v47 + 24))) = uint8(int8(a1))
	*(*uint8)(unsafe.Pointer(uintptr(v3 + v47 + 23))) = uint8(int8(a1))
	return result
}
func sub_484450(a1 int32, a2 int32) int32 {
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
		v14    int32
		v15    int32
		v16    int32
		v17    int32
		v18    int32
		v19    int32
		v20    int32
		v21    int32
		v22    int32
		v23    int32
		v24    int32
		v25    int32
		v26    int32
		v27    int32
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
		v42    int32
		v43    int32
		v44    int32
		v45    int32
		v46    int32
		v47    int32
	)
	result = a1
	v3 = int32(dword_5d4594_3798804)
	*(*uint16)(unsafe.Pointer(uintptr(a2 + 46))) = uint16(int16(a1))
	v4 = v3 + a2
	*(*uint32)(unsafe.Pointer(uintptr(v4 + 44))) = uint32(a1)
	*(*uint16)(unsafe.Pointer(uintptr(v4 + 48))) = uint16(int16(a1))
	v5 = v3 + v3 + a2
	*(*uint16)(unsafe.Pointer(uintptr(v5 + 42))) = uint16(int16(a1))
	*(*uint32)(unsafe.Pointer(uintptr(v5 + 44))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v5 + 48))) = uint32(a1)
	v6 = v3 + v5
	*(*uint32)(unsafe.Pointer(uintptr(v6 + 40))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v6 + 44))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v6 + 48))) = uint32(a1)
	*(*uint16)(unsafe.Pointer(uintptr(v6 + 52))) = uint16(int16(a1))
	v7 = v3 + v6
	*(*uint16)(unsafe.Pointer(uintptr(v7 + 38))) = uint16(int16(a1))
	*(*uint32)(unsafe.Pointer(uintptr(v7 + 40))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v7 + 44))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v7 + 48))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v7 + 52))) = uint32(a1)
	v8 = v3 + v7
	*(*uint32)(unsafe.Pointer(uintptr(v8 + 36))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v8 + 40))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v8 + 44))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v8 + 48))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v8 + 52))) = uint32(a1)
	*(*uint16)(unsafe.Pointer(uintptr(v8 + 56))) = uint16(int16(a1))
	v9 = v3 + v8
	*(*uint16)(unsafe.Pointer(uintptr(v9 + 34))) = uint16(int16(a1))
	*(*uint32)(unsafe.Pointer(uintptr(v9 + 36))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v9 + 40))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v9 + 44))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v9 + 48))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v9 + 52))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v9 + 56))) = uint32(a1)
	v10 = v3 + v9
	*(*uint32)(unsafe.Pointer(uintptr(v10 + 32))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v10 + 36))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v10 + 40))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v10 + 44))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v10 + 48))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v10 + 52))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v10 + 56))) = uint32(a1)
	*(*uint16)(unsafe.Pointer(uintptr(v10 + 60))) = uint16(int16(a1))
	v11 = v3 + v10
	*(*uint16)(unsafe.Pointer(uintptr(v11 + 30))) = uint16(int16(a1))
	*(*uint32)(unsafe.Pointer(uintptr(v11 + 32))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v11 + 36))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v11 + 40))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v11 + 44))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v11 + 48))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v11 + 52))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v11 + 56))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v11 + 60))) = uint32(a1)
	v12 = v3 + v11
	*(*uint32)(unsafe.Pointer(uintptr(v12 + 28))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v12 + 32))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v12 + 36))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v12 + 40))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v12 + 44))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v12 + 48))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v12 + 52))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v12 + 56))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v12 + 60))) = uint32(a1)
	*(*uint16)(unsafe.Pointer(uintptr(v12 + 64))) = uint16(int16(a1))
	v13 = v3 + v12
	*(*uint16)(unsafe.Pointer(uintptr(v13 + 26))) = uint16(int16(a1))
	*(*uint32)(unsafe.Pointer(uintptr(v13 + 28))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v13 + 32))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v13 + 36))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v13 + 40))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v13 + 44))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v13 + 48))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v13 + 52))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v13 + 56))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v13 + 60))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v13 + 64))) = uint32(a1)
	v14 = v3 + v13
	*(*uint32)(unsafe.Pointer(uintptr(v14 + 24))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v14 + 28))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v14 + 32))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v14 + 36))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v14 + 40))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v14 + 44))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v14 + 48))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v14 + 52))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v14 + 56))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v14 + 60))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v14 + 64))) = uint32(a1)
	*(*uint16)(unsafe.Pointer(uintptr(v14 + 68))) = uint16(int16(a1))
	v15 = v3 + v14
	*(*uint16)(unsafe.Pointer(uintptr(v15 + 22))) = uint16(int16(a1))
	*(*uint32)(unsafe.Pointer(uintptr(v15 + 24))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v15 + 28))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v15 + 32))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v15 + 36))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v15 + 40))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v15 + 44))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v15 + 48))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v15 + 52))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v15 + 56))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v15 + 60))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v15 + 64))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v15 + 68))) = uint32(a1)
	v16 = v3 + v15
	*(*uint32)(unsafe.Pointer(uintptr(v16 + 20))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v16 + 24))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v16 + 28))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v16 + 32))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v16 + 36))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v16 + 40))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v16 + 44))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v16 + 48))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v16 + 52))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v16 + 56))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v16 + 60))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v16 + 64))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v16 + 68))) = uint32(a1)
	*(*uint16)(unsafe.Pointer(uintptr(v16 + 72))) = uint16(int16(a1))
	v17 = v3 + v16
	*(*uint16)(unsafe.Pointer(uintptr(v17 + 18))) = uint16(int16(a1))
	*(*uint32)(unsafe.Pointer(uintptr(v17 + 20))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v17 + 24))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v17 + 28))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v17 + 32))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v17 + 36))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v17 + 40))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v17 + 44))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v17 + 48))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v17 + 52))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v17 + 56))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v17 + 60))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v17 + 64))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v17 + 68))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v17 + 72))) = uint32(a1)
	v18 = v3 + v17
	*(*uint32)(unsafe.Pointer(uintptr(v18 + 16))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v18 + 20))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v18 + 24))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v18 + 28))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v18 + 32))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v18 + 36))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v18 + 40))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v18 + 44))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v18 + 48))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v18 + 52))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v18 + 56))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v18 + 60))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v18 + 64))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v18 + 68))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v18 + 72))) = uint32(a1)
	*(*uint16)(unsafe.Pointer(uintptr(v18 + 76))) = uint16(int16(a1))
	v19 = v3 + v18
	*(*uint16)(unsafe.Pointer(uintptr(v19 + 14))) = uint16(int16(a1))
	*(*uint32)(unsafe.Pointer(uintptr(v19 + 16))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v19 + 20))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v19 + 24))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v19 + 28))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v19 + 32))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v19 + 36))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v19 + 40))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v19 + 44))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v19 + 48))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v19 + 52))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v19 + 56))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v19 + 60))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v19 + 64))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v19 + 68))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v19 + 72))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v19 + 76))) = uint32(a1)
	v20 = v3 + v19
	*(*uint32)(unsafe.Pointer(uintptr(v20 + 12))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v20 + 16))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v20 + 20))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v20 + 24))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v20 + 28))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v20 + 32))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v20 + 36))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v20 + 40))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v20 + 44))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v20 + 48))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v20 + 52))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v20 + 56))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v20 + 60))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v20 + 64))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v20 + 68))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v20 + 72))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v20 + 76))) = uint32(a1)
	*(*uint16)(unsafe.Pointer(uintptr(v20 + 80))) = uint16(int16(a1))
	v21 = v3 + v20
	*(*uint16)(unsafe.Pointer(uintptr(v21 + 10))) = uint16(int16(a1))
	*(*uint32)(unsafe.Pointer(uintptr(v21 + 12))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v21 + 16))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v21 + 20))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v21 + 24))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v21 + 28))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v21 + 32))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v21 + 36))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v21 + 40))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v21 + 44))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v21 + 48))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v21 + 52))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v21 + 56))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v21 + 60))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v21 + 64))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v21 + 68))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v21 + 72))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v21 + 76))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v21 + 80))) = uint32(a1)
	v22 = v3 + v21
	*(*uint32)(unsafe.Pointer(uintptr(v22 + 8))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v22 + 12))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v22 + 16))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v22 + 20))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v22 + 24))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v22 + 28))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v22 + 32))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v22 + 36))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v22 + 40))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v22 + 44))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v22 + 48))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v22 + 52))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v22 + 56))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v22 + 60))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v22 + 64))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v22 + 68))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v22 + 72))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v22 + 76))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v22 + 80))) = uint32(a1)
	*(*uint16)(unsafe.Pointer(uintptr(v22 + 84))) = uint16(int16(a1))
	v23 = v3 + v22
	*(*uint16)(unsafe.Pointer(uintptr(v23 + 6))) = uint16(int16(a1))
	*(*uint32)(unsafe.Pointer(uintptr(v23 + 8))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v23 + 12))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v23 + 16))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v23 + 20))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v23 + 24))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v23 + 28))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v23 + 32))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v23 + 36))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v23 + 40))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v23 + 44))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v23 + 48))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v23 + 52))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v23 + 56))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v23 + 60))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v23 + 64))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v23 + 68))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v23 + 72))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v23 + 76))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v23 + 80))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v23 + 84))) = uint32(a1)
	v24 = v3 + v23
	*(*uint32)(unsafe.Pointer(uintptr(v24 + 4))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v24 + 8))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v24 + 12))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v24 + 16))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v24 + 20))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v24 + 24))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v24 + 28))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v24 + 32))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v24 + 36))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v24 + 40))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v24 + 44))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v24 + 48))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v24 + 52))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v24 + 56))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v24 + 60))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v24 + 64))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v24 + 68))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v24 + 72))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v24 + 76))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v24 + 80))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v24 + 84))) = uint32(a1)
	*(*uint16)(unsafe.Pointer(uintptr(v24 + 88))) = uint16(int16(a1))
	v25 = v3 + v24
	*(*uint16)(unsafe.Pointer(uintptr(v25 + 2))) = uint16(int16(a1))
	*(*uint32)(unsafe.Pointer(uintptr(v25 + 4))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v25 + 8))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v25 + 12))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v25 + 16))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v25 + 20))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v25 + 24))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v25 + 28))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v25 + 32))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v25 + 36))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v25 + 40))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v25 + 44))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v25 + 48))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v25 + 52))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v25 + 56))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v25 + 60))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v25 + 64))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v25 + 68))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v25 + 72))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v25 + 76))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v25 + 80))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v25 + 84))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v25 + 88))) = uint32(a1)
	v26 = v3 + v25
	*(*uint16)(unsafe.Pointer(uintptr(v26 + 2))) = uint16(int16(a1))
	*(*uint32)(unsafe.Pointer(uintptr(v26 + 4))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v26 + 8))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v26 + 12))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v26 + 16))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v26 + 20))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v26 + 24))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v26 + 28))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v26 + 32))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v26 + 36))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v26 + 40))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v26 + 44))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v26 + 48))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v26 + 52))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v26 + 56))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v26 + 60))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v26 + 64))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v26 + 68))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v26 + 72))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v26 + 76))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v26 + 80))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v26 + 84))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v26 + 88))) = uint32(a1)
	v27 = v3 + v26
	*(*uint32)(unsafe.Pointer(uintptr(v27 + 4))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v27 + 8))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v27 + 12))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v27 + 16))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v27 + 20))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v27 + 24))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v27 + 28))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v27 + 32))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v27 + 36))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v27 + 40))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v27 + 44))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v27 + 48))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v27 + 52))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v27 + 56))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v27 + 60))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v27 + 64))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v27 + 68))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v27 + 72))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v27 + 76))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v27 + 80))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v27 + 84))) = uint32(a1)
	*(*uint16)(unsafe.Pointer(uintptr(v27 + 88))) = uint16(int16(a1))
	v28 = v3 + v27
	*(*uint16)(unsafe.Pointer(uintptr(v28 + 6))) = uint16(int16(a1))
	*(*uint32)(unsafe.Pointer(uintptr(v28 + 8))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v28 + 12))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v28 + 16))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v28 + 20))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v28 + 24))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v28 + 28))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v28 + 32))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v28 + 36))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v28 + 40))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v28 + 44))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v28 + 48))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v28 + 52))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v28 + 56))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v28 + 60))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v28 + 64))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v28 + 68))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v28 + 72))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v28 + 76))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v28 + 80))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v28 + 84))) = uint32(a1)
	v29 = v3 + v28
	*(*uint32)(unsafe.Pointer(uintptr(v29 + 8))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v29 + 12))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v29 + 16))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v29 + 20))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v29 + 24))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v29 + 28))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v29 + 32))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v29 + 36))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v29 + 40))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v29 + 44))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v29 + 48))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v29 + 52))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v29 + 56))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v29 + 60))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v29 + 64))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v29 + 68))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v29 + 72))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v29 + 76))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v29 + 80))) = uint32(a1)
	*(*uint16)(unsafe.Pointer(uintptr(v29 + 84))) = uint16(int16(a1))
	v30 = v3 + v29
	*(*uint16)(unsafe.Pointer(uintptr(v30 + 10))) = uint16(int16(a1))
	*(*uint32)(unsafe.Pointer(uintptr(v30 + 12))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v30 + 16))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v30 + 20))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v30 + 24))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v30 + 28))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v30 + 32))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v30 + 36))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v30 + 40))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v30 + 44))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v30 + 48))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v30 + 52))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v30 + 56))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v30 + 60))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v30 + 64))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v30 + 68))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v30 + 72))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v30 + 76))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v30 + 80))) = uint32(a1)
	v31 = v3 + v30
	*(*uint32)(unsafe.Pointer(uintptr(v31 + 12))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v31 + 16))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v31 + 20))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v31 + 24))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v31 + 28))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v31 + 32))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v31 + 36))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v31 + 40))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v31 + 44))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v31 + 48))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v31 + 52))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v31 + 56))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v31 + 60))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v31 + 64))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v31 + 68))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v31 + 72))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v31 + 76))) = uint32(a1)
	*(*uint16)(unsafe.Pointer(uintptr(v31 + 80))) = uint16(int16(a1))
	v32 = v3 + v31
	*(*uint16)(unsafe.Pointer(uintptr(v32 + 14))) = uint16(int16(a1))
	*(*uint32)(unsafe.Pointer(uintptr(v32 + 16))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v32 + 20))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v32 + 24))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v32 + 28))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v32 + 32))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v32 + 36))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v32 + 40))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v32 + 44))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v32 + 48))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v32 + 52))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v32 + 56))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v32 + 60))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v32 + 64))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v32 + 68))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v32 + 72))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v32 + 76))) = uint32(a1)
	v33 = v3 + v32
	*(*uint32)(unsafe.Pointer(uintptr(v33 + 16))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v33 + 20))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v33 + 24))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v33 + 28))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v33 + 32))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v33 + 36))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v33 + 40))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v33 + 44))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v33 + 48))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v33 + 52))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v33 + 56))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v33 + 60))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v33 + 64))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v33 + 68))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v33 + 72))) = uint32(a1)
	*(*uint16)(unsafe.Pointer(uintptr(v33 + 76))) = uint16(int16(a1))
	v34 = v3 + v33
	*(*uint16)(unsafe.Pointer(uintptr(v34 + 18))) = uint16(int16(a1))
	*(*uint32)(unsafe.Pointer(uintptr(v34 + 20))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v34 + 24))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v34 + 28))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v34 + 32))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v34 + 36))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v34 + 40))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v34 + 44))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v34 + 48))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v34 + 52))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v34 + 56))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v34 + 60))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v34 + 64))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v34 + 68))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v34 + 72))) = uint32(a1)
	v35 = v3 + v34
	*(*uint32)(unsafe.Pointer(uintptr(v35 + 20))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v35 + 24))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v35 + 28))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v35 + 32))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v35 + 36))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v35 + 40))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v35 + 44))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v35 + 48))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v35 + 52))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v35 + 56))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v35 + 60))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v35 + 64))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v35 + 68))) = uint32(a1)
	*(*uint16)(unsafe.Pointer(uintptr(v35 + 72))) = uint16(int16(a1))
	v36 = v3 + v35
	*(*uint16)(unsafe.Pointer(uintptr(v36 + 22))) = uint16(int16(a1))
	*(*uint32)(unsafe.Pointer(uintptr(v36 + 24))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v36 + 28))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v36 + 32))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v36 + 36))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v36 + 40))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v36 + 44))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v36 + 48))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v36 + 52))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v36 + 56))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v36 + 60))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v36 + 64))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v36 + 68))) = uint32(a1)
	v37 = v3 + v36
	*(*uint32)(unsafe.Pointer(uintptr(v37 + 24))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v37 + 28))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v37 + 32))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v37 + 36))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v37 + 40))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v37 + 44))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v37 + 48))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v37 + 52))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v37 + 56))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v37 + 60))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v37 + 64))) = uint32(a1)
	*(*uint16)(unsafe.Pointer(uintptr(v37 + 68))) = uint16(int16(a1))
	v38 = v3 + v37
	*(*uint16)(unsafe.Pointer(uintptr(v38 + 26))) = uint16(int16(a1))
	*(*uint32)(unsafe.Pointer(uintptr(v38 + 28))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v38 + 32))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v38 + 36))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v38 + 40))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v38 + 44))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v38 + 48))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v38 + 52))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v38 + 56))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v38 + 60))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v38 + 64))) = uint32(a1)
	v39 = v3 + v38
	*(*uint32)(unsafe.Pointer(uintptr(v39 + 28))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v39 + 32))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v39 + 36))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v39 + 40))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v39 + 44))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v39 + 48))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v39 + 52))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v39 + 56))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v39 + 60))) = uint32(a1)
	*(*uint16)(unsafe.Pointer(uintptr(v39 + 64))) = uint16(int16(a1))
	v40 = v3 + v39
	*(*uint16)(unsafe.Pointer(uintptr(v40 + 30))) = uint16(int16(a1))
	*(*uint32)(unsafe.Pointer(uintptr(v40 + 32))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v40 + 36))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v40 + 40))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v40 + 44))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v40 + 48))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v40 + 52))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v40 + 56))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v40 + 60))) = uint32(a1)
	v41 = v3 + v40
	*(*uint32)(unsafe.Pointer(uintptr(v41 + 32))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v41 + 36))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v41 + 40))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v41 + 44))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v41 + 48))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v41 + 52))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v41 + 56))) = uint32(a1)
	*(*uint16)(unsafe.Pointer(uintptr(v41 + 60))) = uint16(int16(a1))
	v42 = v3 + v41
	*(*uint16)(unsafe.Pointer(uintptr(v42 + 34))) = uint16(int16(a1))
	*(*uint32)(unsafe.Pointer(uintptr(v42 + 36))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v42 + 40))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v42 + 44))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v42 + 48))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v42 + 52))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v42 + 56))) = uint32(a1)
	v43 = v3 + v42
	*(*uint32)(unsafe.Pointer(uintptr(v43 + 36))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v43 + 40))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v43 + 44))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v43 + 48))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v43 + 52))) = uint32(a1)
	*(*uint16)(unsafe.Pointer(uintptr(v43 + 56))) = uint16(int16(a1))
	v44 = v3 + v43
	*(*uint16)(unsafe.Pointer(uintptr(v44 + 38))) = uint16(int16(a1))
	*(*uint32)(unsafe.Pointer(uintptr(v44 + 40))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v44 + 44))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v44 + 48))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v44 + 52))) = uint32(a1)
	v45 = v3 + v44
	*(*uint32)(unsafe.Pointer(uintptr(v45 + 40))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v45 + 44))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v45 + 48))) = uint32(a1)
	*(*uint16)(unsafe.Pointer(uintptr(v45 + 52))) = uint16(int16(a1))
	v46 = v3 + v45
	*(*uint16)(unsafe.Pointer(uintptr(v46 + 42))) = uint16(int16(a1))
	*(*uint32)(unsafe.Pointer(uintptr(v46 + 44))) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(v46 + 48))) = uint32(a1)
	v47 = v3 + v46
	*(*uint32)(unsafe.Pointer(uintptr(v47 + 44))) = uint32(a1)
	*(*uint16)(unsafe.Pointer(uintptr(v47 + 48))) = uint16(int16(a1))
	*(*uint16)(unsafe.Pointer(uintptr(v3 + v47 + 46))) = uint16(int16(a1))
	return result
}
func sub_484BA0(a1 float32) {
	if float64(a1) >= 1.0 {
		*(*float32)(unsafe.Pointer(&dword_587000_154968)) = a1
	} else {
		dword_587000_154968 = 0x3F800000
	}
}
func nox_xxx_spriteChangeLightColor_484BE0(a1 *uint32, a2 int32, a3 int32, a4 int32) *uint32 {
	var result *uint32
	result = a1
	*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*4)) = uint32(a2)
	*a1 = 2
	*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*5)) = uint32(a3)
	*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*6)) = uint32(a4)
	return result
}
func sub_484C00(a1 int32, a2 int32) int64 {
	var result int64
	result = int64(float64(a2)*0.0027777778**(*float64)(unsafe.Pointer(&qword_581450_9552)) + *(*float64)(unsafe.Pointer(&qword_581450_9544)))
	*(*uint16)(unsafe.Pointer(uintptr(a1 + 28))) = uint16(int16(result))
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 32))) = 0
	return result
}
func nox_xxx_spriteChangeLightSize_484C30(a1 int32, a2 int32) int64 {
	var result int64
	result = int64(float64(a2)*0.0027777778**(*float64)(unsafe.Pointer(&qword_581450_9552)) + *(*float64)(unsafe.Pointer(&qword_581450_9544)))
	*(*uint16)(unsafe.Pointer(uintptr(a1 + 30))) = uint16(int16(result))
	return result
}
func sub_484C60(a1 float32) int32 {
	var v2 float64
	if float64(a1) <= float64(*(*float32)(unsafe.Pointer(&dword_587000_154968))) {
		return 0
	}
	if float64(a1) <= 31.0 {
		v2 = float64(a1)
	} else {
		v2 = 31.0
	}
	return nox_double2int(math.Sqrt(((float64(*mem_getFloatPtr(0x587000, 154980))+v2)/float64(*mem_getFloatPtr(0x587000, 154980)+*(*float32)(unsafe.Pointer(&dword_587000_154968))) + 1.0) * float64(a1*a1/(*mem_getFloatPtr(0x587000, 0x25D60)**mem_getFloatPtr(0x587000, 0x25D5C)))))
}
func sub_484CE0(a1 int32, a2 float32) int32 {
	var result int32
	if float64(a2) > 63.0 {
		a2 = 63.0
	}
	*(*float32)(unsafe.Pointer(uintptr(a1 + 4))) = a2
	result = sub_484C60(a2)
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 8))) = uint32(result)
	return result
}
func sub_484D20(a1 int32, a2 float32) int32 {
	var result int32
	if float64(a2) > 63.0 {
		a2 = 63.0
	}
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 12))) = uint32(int32(int64(float64(a2)**(*float64)(unsafe.Pointer(&qword_581450_9552)) + *(*float64)(unsafe.Pointer(&qword_581450_9544)))))
	result = sub_484C60(a2)
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 8))) = uint32(result)
	return result
}
func nox_xxx_spriteChangeIntensity_484D70_light_intensity(a1 int32, a2 float32) int32 {
	var result int32
	if float64(a2) > 63.0 {
		a2 = 63.0
	}
	*(*float32)(unsafe.Pointer(uintptr(a1 + 4))) = a2
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 12))) = uint32(int32(int64(float64(a2)**(*float64)(unsafe.Pointer(&qword_581450_9552)) + *(*float64)(unsafe.Pointer(&qword_581450_9544)))))
	result = sub_484C60(a2)
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 8))) = uint32(result)
	return result
}
func sub_484DC0(a1 *int2, a2 *int2, a3 *int32) int32 {
	var (
		result int32
		v4     int32
		v5     int32
		v6     int32
		v7     int32
		v8     int32
		v9     *int32
		v10    int32
	)
	result = (a1.field_4 + 0x8000) >> 16
	*a3 = result
	v4 = ((a2.field_4 + 0x8000) >> 16) - result + 1
	*(*int32)(unsafe.Add(unsafe.Pointer(a3), unsafe.Sizeof(int32(0))*1)) = v4
	if v4 > 0 {
		v5 = a1.field_0
		result = (a2.field_0 - a1.field_0) / v4
		v6 = *a3 + v4
		if *a3 < 0 {
			*(*int32)(unsafe.Add(unsafe.Pointer(a3), unsafe.Sizeof(int32(0))*1)) = v6
			*a3 = 0
		}
		if v6 > 44 {
			*(*int32)(unsafe.Add(unsafe.Pointer(a3), unsafe.Sizeof(int32(0))*1)) += 44 - v6
			v6 = 44
		}
		v7 = *a3
		if *a3 < v6 {
			v8 = v5 + 0x8000
			v9 = (*int32)(unsafe.Add(unsafe.Pointer(a3), unsafe.Sizeof(int32(0))*uintptr(v7+2)))
			v10 = v6 - v7
			for {
				*v9 = v8 >> 16
				if v8>>16 < 0 {
					*v9 = 0
				}
				if *v9 >= 56 {
					*v9 = 55
				}
				v8 += result
				v9 = (*int32)(unsafe.Add(unsafe.Pointer(v9), unsafe.Sizeof(int32(0))*1))
				v10--
				if v10 == 0 {
					break
				}
			}
		}
	}
	return result
}
func sub_4854D0(a1 *int2, a2 *int2, a3 *int2, a4 int32, a5 float32) int32 {
	var (
		v5     int32
		v6     int32
		v7     *int2
		v8     *int2
		v9     *int2
		result int32
		v11    int32
		v12    [46]int32
		v13    [46]int32
		v14    [46]int32
	)
	v5 = a1.field_4
	v6 = a2.field_4
	if v5 > v6 {
		v9 = a3
		if v6 <= a3.field_4 {
			v8 = a2
			v7 = a1
			goto LABEL_9
		}
		v8 = a3
	} else {
		v7 = a3
		if v5 <= a3.field_4 {
			v8 = a1
			v9 = a2
			goto LABEL_9
		}
		v8 = a3
	}
	v9 = a1
	v7 = a2
LABEL_9:
	sub_484DC0(v8, v9, &v14[0])
	sub_484DC0(v8, v7, &v13[0])
	sub_4855D0(int32(uintptr(unsafe.Pointer(&v14[0]))), int32(uintptr(unsafe.Pointer(&v13[0]))), a4, a5)
	result = v9.field_4
	v11 = v7.field_4
	if result >= v11 {
		if result > v11 {
			sub_484DC0(v7, v9, &v12[0])
			result = sub_4855D0(int32(uintptr(unsafe.Pointer(&v14[0]))), int32(uintptr(unsafe.Pointer(&v12[0]))), a4, a5)
		}
	} else {
		sub_484DC0(v9, v7, &v12[0])
		result = sub_4855D0(int32(uintptr(unsafe.Pointer(&v12[0]))), int32(uintptr(unsafe.Pointer(&v13[0]))), a4, a5)
	}
	return result
}
func sub_4855D0(a1 int32, a2 int32, a3 int32, a4 float32) int32 {
	var (
		v4     int32
		v5     int32
		v6     int32
		result int32
		v8     float64
		v9     int32
		v10    float64
		v11    *uint32
		v12    int32
		v13    float64
		v14    float64
		v15    *uint8
		v16    float64
		v17    float64
		v18    int32
		v19    int32
	)
	v4 = a2
	v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1))))
	v6 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1))) + *(*uint32)(unsafe.Pointer(uintptr(a1 + 4))))
	if *(*uint32)(unsafe.Pointer(uintptr(a1))) < uint32(*(*int32)(unsafe.Pointer(uintptr(a2)))) {
		v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2))))
	}
	if uint32(v6) > uint32(*(*int32)(unsafe.Pointer(uintptr(a2 + 4))))+*(*uint32)(unsafe.Pointer(uintptr(a2))) {
		v6 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 4))) + *(*uint32)(unsafe.Pointer(uintptr(a2))))
	}
	result = v6 - v5
	if v6-v5 > 1 {
		if float64(a4) <= 31.0 {
			v8 = float64(a4)
		} else {
			v8 = 31.0
		}
		*(*float32)(unsafe.Pointer(&v18)) = a4 * a4
		result = int32(uint32(v5*23) - *(*uint32)(unsafe.Pointer(uintptr(a3 + 4))))
		v9 = v5
		v10 = float64(result)
		v19 = v5
		if v5 < v6 {
			v11 = (*uint32)(unsafe.Pointer(uintptr(v4 + v5*4 + 8)))
			result = a1 - v4
			for {
				v12 = int32(*(*uint32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(v11))), result)))))
				v13 = float64(int32(uint32(v12*23) - *(*uint32)(unsafe.Pointer(uintptr(a3)))))
				if uint32(v12) < *v11 {
					v14 = float64(*mem_getFloatPtr(0x587000, 0x25D6C)) + v8
					v15 = (*uint8)(memmap.PtrOff(0x8531A0, uint32(v12*+40+v12*5+v9)))
					for {
						v16 = v14 / ((v13*v13+v10*v10)*float64(*mem_getFloatPtr(0x587000, 0x25D68))*float64(*mem_getFloatPtr(0x587000, 0x25D5C))/float64(*(*float32)(unsafe.Pointer(&v18))) + 1.0)
						if v16 > float64(*mem_getFloatPtr(0x587000, 0x25D6C)) {
							v17 = v16 - float64(*mem_getFloatPtr(0x587000, 0x25D6C)) + float64(*v15)
							if v17 > 31.0 {
								v17 = 31.0
							}
							*v15 = uint8(int8(int64(v17)))
						}
						v12++
						v13 = v13 + 23.0
						v15 = (*uint8)(unsafe.Add(unsafe.Pointer(v15), 45))
						if uint32(v12) >= *v11 {
							break
						}
					}
					v9 = v19
					result = a1 - v4
				}
				v10 = v10 + 23.0
				v9++
				v11 = (*uint32)(unsafe.Add(unsafe.Pointer(v11), unsafe.Sizeof(uint32(0))*1))
				v19 = v9
				if v9 >= v6 {
					break
				}
			}
		}
	}
	return result
}
func sub_485AE0(a1 int32, a2 int32) uint32 {
	var (
		v2     *uint8
		v3     int32
		v4     *uint8
		v5     int32
		v6     int32
		result uint32
		v8     bool
	)
	v2 = *(**uint8)(unsafe.Pointer(&dword_5d4594_1193184))
	v3 = a1
	v4 = (*uint8)(unsafe.Pointer(uintptr(dword_5d4594_1193180 - 1)))
	v5 = int32(dword_5d4594_1193176)
	for {
		v6 = int32(*func() *uint8 {
			p := &v2
			x := *p
			*p = (*uint8)(unsafe.Add(unsafe.Pointer(*p), 1))
			return x
		}()) << 13
		v4 = (*uint8)(unsafe.Add(unsafe.Pointer(v4), 1))
		result = uint32(v5+v6) >> 8
		v5 += a2
		v8 = func() int32 {
			p := &v3
			x := *p
			*p--
			return x
		}() <= 1
		*v4 = *memmap.PtrUint8(6112660, result+0x27846C)
		if v8 {
			break
		}
	}
	dword_5d4594_1193176 = uint32(v5)
	dword_5d4594_1193184 = uint32(uintptr(unsafe.Pointer(v2)))
	dword_5d4594_1193180 = uint32(uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v4), 1)))))
	return result
}
func nox_thing_read_floor_485B30(a1 int32, a2 *byte) int32 {
	var (
		v2  int32
		v3  *uint8
		v4  int32
		v5  int32
		v6  *byte
		v7  int32
		v9  *uint8
		v10 int32
		v11 int32
		v12 *int32
		v13 int32
		v14 *byte
		v15 int8
		v16 int32
		v17 uint8
		i   int32
		v19 uint8
		v20 int32
		v21 *byte
		v22 [32]byte
		v23 uint8
	)
	v2 = a1
	v16 = a1
	v3 = (*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 8))) + 4)))
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 8))) = uint32(uintptr(unsafe.Pointer(v3)))
	*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a1))), 0)) = *v3
	*(*uint32)(unsafe.Pointer(uintptr(v2 + 8))) = uint32(uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v3), 1)))))
	nox_memfile_read(unsafe.Pointer(&v22[0]), 1, int32(uint8(int8(a1))), (*nox_memfile)(unsafe.Pointer(uintptr(v16))))
	v4 = int32(dword_5d4594_251568)
	v5 = 0
	v22[uint8(int8(a1))] = 0
	if v4 <= 0 {
	LABEL_5:
		v7 = a1
	} else {
		v6 = (*byte)(memmap.PtrOff(0x85B3FC, 32484))
		for libc.StrCmp(v6, &v22[0]) != 0 {
			v5++
			v6 = (*byte)(unsafe.Add(unsafe.Pointer(v6), 60))
			if v5 >= *(*int32)(unsafe.Pointer(&dword_5d4594_251568)) {
				goto LABEL_5
			}
		}
		v7 = v5
	}
	if uint32(v5) == dword_5d4594_251568 {
		return 0
	}
	v9 = (*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 8))) + 12)))
	*(*uint32)(unsafe.Pointer(uintptr(v2 + 8))) = uint32(uintptr(unsafe.Pointer(v9)))
	*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v21))), 0)) = *v9
	*(*uint32)(unsafe.Pointer(uintptr(v2 + 8))) = uint32(uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v9), 1)))))
	v19 = *(*uint8)(unsafe.Add(unsafe.Pointer(v9), 1))
	*(*uint32)(unsafe.Pointer(uintptr(v2 + 8))) = uint32(uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v9), 2)))))
	v17 = *(*uint8)(unsafe.Add(unsafe.Pointer(v9), 2))
	*(*uint32)(unsafe.Pointer(uintptr(v2 + 8))) = uint32(uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v9), 4)))))
	v10 = int32(uint8(uintptr(unsafe.Pointer(v21)))) * int32(v19) * int32(v17)
	v20 = v7 * 15
	*memmap.PtrUint32(0x85B3FC, uint32(v7*60+0x7F04)) = uint32(uintptr(alloc.Calloc(int(v10), 4)))
	v11 = 0
	for i = 0; v11 < v10; *(*uint32)(unsafe.Pointer(uintptr(*memmap.PtrUint32(0x85B3FC, uint32(v20*4+0x7F04)) + uint32(v11*4) - 4))) = uint32(uintptr(unsafe.Pointer(nox_xxx_readImgMB_42FAA0(v13, int8(uintptr(unsafe.Pointer(v21))), a2)))) {
		v12 = *(**int32)(unsafe.Pointer(uintptr(v2 + 8)))
		v13 = *v12
		*(*uint32)(unsafe.Pointer(uintptr(v2 + 8))) = uint32(uintptr(unsafe.Pointer((*int32)(unsafe.Add(unsafe.Pointer(v12), unsafe.Sizeof(int32(0))*1)))))
		*a2 = byte(*memmap.PtrUint8(6112660, 1193192))
		if v13 == -1 {
			v14 = *(**byte)(unsafe.Pointer(uintptr(v2 + 8)))
			v15 = int8(*func() *byte {
				p := &v14
				x := *p
				*p = (*byte)(unsafe.Add(unsafe.Pointer(*p), 1))
				return x
			}())
			*(*uint32)(unsafe.Pointer(uintptr(v2 + 8))) = uint32(uintptr(unsafe.Pointer(v14)))
			*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v21))), 0)) = uint8(v15)
			v23 = uint8(*v14)
			*(*uint32)(unsafe.Pointer(uintptr(v2 + 8))) = uint32(uintptr(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v14), 1)))))
			nox_memfile_read(unsafe.Pointer(a2), 1, int32(v23), (*nox_memfile)(unsafe.Pointer(uintptr(v2))))
			v13 = -1
			*(*byte)(unsafe.Add(unsafe.Pointer(a2), v23)) = 0
			v11 = i
		}
		i = func() int32 {
			p := &v11
			*p++
			return *p
		}()
	}
	return 1
}
func nox_thing_read_edge_485D40(a1 int32, a2 *byte) int32 {
	var (
		v2     int32
		v3     *uint8
		v4     int32
		v5     int32
		v6     *byte
		v7     int32
		result int32
		v9     *uint8
		v10    int8
		v11    uint8
		v12    int32
		v13    uint8
		v14    int32
		v15    int32
		v16    *int32
		v17    int32
		v18    *byte
		v19    int8
		v20    *int32
		v21    int32
		v22    int32
		i      int32
		v24    int32
		v25    uint32
		v26    [64]byte
		v27    uint8
	)
	v2 = a1
	v22 = a1
	v3 = (*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 8))) + 4)))
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 8))) = uint32(uintptr(unsafe.Pointer(v3)))
	*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a1))), 0)) = *v3
	*(*uint32)(unsafe.Pointer(uintptr(v2 + 8))) = uint32(uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v3), 1)))))
	nox_memfile_read(unsafe.Pointer(&v26[0]), 1, int32(uint8(int8(a1))), (*nox_memfile)(unsafe.Pointer(uintptr(v22))))
	v4 = int32(dword_5d4594_251572)
	v5 = 0
	v26[uint8(int8(a1))] = 0
	if v4 <= 0 {
	LABEL_5:
		v7 = a1
	} else {
		v6 = (*byte)(memmap.PtrOff(0x85B3FC, 28644))
		for libc.StrCmp(v6, &v26[0]) != 0 {
			v5++
			v6 = (*byte)(unsafe.Add(unsafe.Pointer(v6), 60))
			if v5 >= *(*int32)(unsafe.Pointer(&dword_5d4594_251572)) {
				goto LABEL_5
			}
		}
		v7 = v5
	}
	if uint32(v5) == dword_5d4594_251572 {
		return 0
	}
	v9 = (*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 8))) + 9)))
	*(*uint32)(unsafe.Pointer(uintptr(v2 + 8))) = uint32(uintptr(unsafe.Pointer(v9)))
	v25 = uint32(*v9)
	*(*uint32)(unsafe.Pointer(uintptr(v2 + 8))) = uint32(uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v9), 2)))))
	v10 = int8(*(*uint8)(unsafe.Add(unsafe.Pointer(v9), 2)))
	*(*uint32)(unsafe.Pointer(uintptr(v2 + 8))) = uint32(uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v9), 3)))))
	if int32(v10) == 1 {
		return 0
	}
	v11 = *(*uint8)(unsafe.Add(unsafe.Pointer(v9), 3))
	v12 = int32(uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v9), 4)))))
	*(*uint32)(unsafe.Pointer(uintptr(v2 + 8))) = uint32(uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v9), 4)))))
	v13 = *(*uint8)(unsafe.Add(unsafe.Pointer(v9), 4))
	*(*uint32)(unsafe.Pointer(uintptr(v2 + 8))) = uint32(v12 + 1)
	v14 = int32(v25 * 2 * uint32(int32(v11)+int32(v13)))
	result = int32(uintptr(alloc.Calloc(int(v14), 5)))
	libc.MemSet(unsafe.Pointer(uintptr(result)), 0, int(v14*5))
	v24 = v7 * 15
	*memmap.PtrUint32(0x85B3FC, uint32(v7*60+0x7004)) = uint32(result)
	if result != 0 {
		v15 = 0
		for i = 0; v15 < v14; *(*uint32)(unsafe.Pointer(uintptr(*memmap.PtrUint32(0x85B3FC, uint32(v24*4+0x7004)) + uint32(v15*4) - 4))) = uint32(uintptr(unsafe.Pointer(nox_xxx_readImgMB_42FAA0(v17, int8(uint8(v25)), a2)))) {
			v16 = *(**int32)(unsafe.Pointer(uintptr(v2 + 8)))
			v17 = *v16
			*(*uint32)(unsafe.Pointer(uintptr(v2 + 8))) = uint32(uintptr(unsafe.Pointer((*int32)(unsafe.Add(unsafe.Pointer(v16), unsafe.Sizeof(int32(0))*1)))))
			*a2 = byte(*memmap.PtrUint8(6112660, 1193196))
			if v17 == -1 {
				v18 = *(**byte)(unsafe.Pointer(uintptr(v2 + 8)))
				v19 = int8(*func() *byte {
					p := &v18
					x := *p
					*p = (*byte)(unsafe.Add(unsafe.Pointer(*p), 1))
					return x
				}())
				*(*uint32)(unsafe.Pointer(uintptr(v2 + 8))) = uint32(uintptr(unsafe.Pointer(v18)))
				*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v25))), 0)) = uint8(v19)
				v27 = uint8(*v18)
				*(*uint32)(unsafe.Pointer(uintptr(v2 + 8))) = uint32(uintptr(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v18), 1)))))
				nox_memfile_read(unsafe.Pointer(a2), 1, int32(v27), (*nox_memfile)(unsafe.Pointer(uintptr(v2))))
				v17 = -1
				*(*byte)(unsafe.Add(unsafe.Pointer(a2), v27)) = 0
				v15 = i
			}
			i = func() int32 {
				p := &v15
				*p++
				return *p
			}()
		}
		v20 = *(**int32)(unsafe.Pointer(uintptr(v2 + 8)))
		v21 = *v20
		*(*uint32)(unsafe.Pointer(uintptr(v2 + 8))) = uint32(uintptr(unsafe.Pointer((*int32)(unsafe.Add(unsafe.Pointer(v20), unsafe.Sizeof(int32(0))*1)))))
		result = bool2int(uint32(v21) == 0x454E4420)
	}
	return result
}
func nox_xxx_initSight_485F80() {
	var (
		v0 int32
		v1 int32
		v2 int32
		i  int32
		v4 int32
		v5 int32
		v6 int32
		v7 int32
	)
	v0 = 0
	v1 = 0
	v2 = 1
	for i = 23; i > 0; i-- {
		*memmap.PtrUint32(0x973CE0, uint32(v1*+4)) = uint32(v0)
		*memmap.PtrUint32(0x973CE0, uint32(v1*4+192)) = uint32(i)
		*memmap.PtrUint32(0x973CE0, uint32(v1*4+384)) = uint32(v2)
		v0 += v2
		v1++
		v2 += 2
	}
	if v1 < 46 {
		v4 = 91 - v1*2
		v5 = v1 - 22
		v6 = v1 * 4
		v7 = 46 - v1
		for {
			*memmap.PtrUint32(0x973CE0, uint32(+v6)) = uint32(v0)
			*memmap.PtrUint32(0x973CE0, uint32(v6+192)) = uint32(v5)
			*memmap.PtrUint32(0x973CE0, uint32(v6+384)) = uint32(v4)
			v0 += v4
			v5++
			v4 -= 2
			v6 += 4
			v7--
			if v7 == 0 {
				break
			}
		}
	}
	var v8 *uint32 = (*uint32)(memmap.PtrOff(0x8529A0, 0))
	for j := int32(0); j < 253; j++ {
		*(*uint32)(unsafe.Add(unsafe.Pointer(v8), unsafe.Sizeof(uint32(0))*uintptr(j))) = uint32(((j - 253) * 256) / 22)
	}
	*memmap.PtrUint32(0x8529A0, 1016) = 0
	var v12 *uint32 = (*uint32)(memmap.PtrOff(0x8529A0, 1020))
	for k := int32(0); k < 256; k++ {
		*(*uint32)(unsafe.Add(unsafe.Pointer(v12), unsafe.Sizeof(uint32(0))*uintptr(k))) = uint32((k * 256) / 22)
	}
}
func nox_xxx_tile_486060() int32 {
	var result int32
	dword_5d4594_1193188 = 1
	result = int32(dword_5d4594_3798804 * 45)
	*memmap.PtrUint32(0x973CE0, 376) = dword_5d4594_3798804*45 + uint32(46<<int32(*memmap.PtrUint8(0x973F18, 7696)))
	return result
}
func sub_4862E0(a3 int32, a4 int32) int32 {
	*(*uint32)(unsafe.Pointer(uintptr(a3))) = 0
	*(*uint64)(unsafe.Pointer(uintptr(a3 + 24))) = uint64(platformTicks())
	sub_486380((*uint32)(unsafe.Pointer(uintptr(a3))), 1000, 0, 0x4000)
	sub_486320((*uint32)(unsafe.Pointer(uintptr(a3))), a4)
	return sub_4863B0((*uint32)(unsafe.Pointer(uintptr(a3))))
}
func sub_486320(a1 *uint32, a2 int32) *uint32 {
	var result *uint32
	result = a1
	*a1 |= 1
	*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*2)) = uint32(a2 << 16)
	return result
}
func sub_486340(a1 *uint32) *uint32 {
	var result *uint32
	result = a1
	*a1 |= 1
	return result
}
func sub_486350(a1p unsafe.Pointer, a2 int32) int32 {
	var (
		a1 int32 = int32(uintptr(a1p))
		v2 int64
		v3 int32
	)
	*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v2))), unsafe.Sizeof(uint32(0))*0)) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 8)))
	v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 4))))
	*(*uint32)(unsafe.Pointer(uintptr(a1))) &= 0xFFFFFFFE
	if uint32(int32(v2)) == uint32(v3) {
		v2 = int64(platformTicks())
		*(*uint64)(unsafe.Pointer(uintptr(a1 + 24))) = uint64(v2)
	}
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 8))) = uint32(a2 << 16)
	return int32(v2)
}
func sub_486380(a1 *uint32, a2 uint32, a3 int32, a4 int32) int32 {
	var result int32
	*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*4)) = a2
	*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*3)) = uint32(a4<<16) / a2
	result = a3
	*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*5)) = uint32(a3)
	return result
}
func sub_4863B0(a2 *uint32) int32 {
	var (
		v1     int32
		v2     uint32
		result int32
		v4     int64
		v5     uint32
		v6     uint32
		v7     uint32
		v8     uint32
		v9     uint32
		v10    uint64
		v11    int32
		v12    uint32
		v13    uint32
		v14    uint32
		v15    uint32
	)
	v1 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*2)) - *(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*1)))
	if v1 == 0 {
		return 0
	}
	if int32(*(*uint8)(unsafe.Pointer(a2)))&1 != 0 {
		*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*1)) = *(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*2))
		v2 = *a2
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v2))), 0)) = uint8(*a2 | 2)
		*a2 = v2
		result = 1
	} else {
		v4 = int64(platformTicks())
		v5 = *(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*6))
		v6 = *(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*7))
		v7 = uint32(int32(v4))
		*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*6)) = uint32(int32(v4))
		*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v4))), unsafe.Sizeof(uint32(0))*0)) = *(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*5))
		*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*7)) = *(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v4))), unsafe.Sizeof(uint32(0))*1))
		v10 = (((uint64(*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v4))), unsafe.Sizeof(uint32(0))*1)))) << 32) | uint64(v7)) - (((uint64(v6)) << 32) | uint64(v5))
		v9 = uint32(((((uint64(*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v4))), unsafe.Sizeof(uint32(0))*1)))) << 32) | uint64(v7)) - (((uint64(v6)) << 32) | uint64(v5))) >> 32)
		v8 = uint32(v10)
		if (((uint64(v9)) << 32) | uint64(uint32(v10))) > (((uint64(v4)) << 32) | uint64(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*4)))) {
			v8 = *(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*4))
		}
		v11 = int32(v8 * *(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*3)))
		if v1 >= 0 {
			if v11 <= v1 {
				v15 = uint32(v11) + *(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*1))
				v14 = *a2
				*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v14))), 0)) = uint8(*a2 | 2)
				*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*1)) = v15
			} else {
				*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*1)) += uint32(v1)
				v14 = *a2
				*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v14))), 0)) = uint8(*a2 | 2)
			}
			*a2 = v14
			result = 1
		} else {
			if v11 <= -v1 {
				v13 = *(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*1)) - uint32(v11)
				v12 = *a2
				*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v12))), 0)) = uint8(*a2 | 2)
				*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*1)) = v13
			} else {
				*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*1)) += uint32(v1)
				v12 = *a2
				*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v12))), 0)) = uint8(*a2 | 2)
			}
			*a2 = v12
			result = 1
		}
	}
	return result
}
func sub_4864A0(a3 *uint32) *uint32 {
	sub_4862E0(int32(uintptr(unsafe.Pointer(a3))), 0x4000)
	sub_4862E0(int32(uintptr(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer(a3), unsafe.Sizeof(uint32(0))*8))))), 100)
	sub_4862E0(int32(uintptr(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer(a3), unsafe.Sizeof(uint32(0))*16))))), 8192)
	sub_486380((*uint32)(unsafe.Add(unsafe.Pointer(a3), unsafe.Sizeof(uint32(0))*8)), 1000, 0, 10)
	sub_486380(a3, 1000, 0, 0x4000)
	sub_486380((*uint32)(unsafe.Add(unsafe.Pointer(a3), unsafe.Sizeof(uint32(0))*16)), 1000, 0, 0x4000)
	return sub_486320(a3, 0x4000)
}
func sub_486520(a2 *uint32) int32 {
	sub_4863B0(a2)
	sub_4863B0((*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*8)))
	return sub_4863B0((*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*16)))
}
func sub_486550(a1 *uint8) int32 {
	return bool2int(int32(*a1)&2 != 0 || int32(*(*uint8)(unsafe.Add(unsafe.Pointer(a1), 32)))&2 != 0 || int32(*(*uint8)(unsafe.Add(unsafe.Pointer(a1), 64)))&2 != 0)
}
func sub_486570(a1 *uint32, a2 *uint32) int32 {
	var v2 int32
	sub_486320(a1, int32(((*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*1))>>16)*(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*1))>>16))>>14))
	sub_4863B0(a1)
	sub_486320((*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*8)), int32((*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*9))>>16)*(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*9))>>16))/100)
	sub_4863B0((*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*8)))
	if *(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*17))>>16 != 8192 {
		v2 = int32((*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*17)) >> 16) + (*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*17)) >> 16) - 8192)
		if v2 >= 0 {
			if uint32(v2) > 0x4000 {
				v2 = 0x4000
			}
		} else {
			v2 = 0
		}
		sub_486320((*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*16)), v2)
	}
	return sub_4863B0((*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*16)))
}
func sub_486620(a1 *uint32) *uint32 {
	var result *uint32
	result = a1
	*a1 &= 0xFFFFFFFD
	*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*16)) &= 0xFFFFFFFD
	*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*8)) &= 0xFFFFFFFD
	return result
}
func sub_486640(a1p unsafe.Pointer, a2 int32) int32 {
	var a1 int32 = int32(uintptr(a1p))
	return int32(uint32(a2) * (*(*uint32)(unsafe.Pointer(uintptr(a1 + 36))) >> 16) / 100)
}
func sub_486670(a1 int32, a2 int32) int32 {
	var result int32
	result = a2
	if a2 != 0 {
		result = a2 - 1
		if a2 == 1 {
			*memmap.PtrUint32(0x587000, 155052) = uint32(a1)
		} else {
			result = a2 - 2
			if a2 == 2 {
				result = a1
				*memmap.PtrUint32(0x587000, 155056) = uint32(a1)
			}
		}
	} else {
		*memmap.PtrUint32(0x587000, 155048) = uint32(a1)
	}
	return result
}
func sub_4866A0(a1 int32) int32 {
	switch a1 {
	case 0:
		return int32(*memmap.PtrUint32(0x587000, 155048))
	case 1:
		return int32(*memmap.PtrUint32(0x587000, 155052))
	case 2:
		return int32(*memmap.PtrUint32(0x587000, 155056))
	}
	return 0x4000
}
func sub_4866D0(a1 *uint32, a2 int32) int32 {
	return int32(*a1 + uint32(a2*36))
}
func sub_4866F0(a1 *byte, a2 *byte) *byte {
	var (
		v2 *byte
		v3 *byte
		v4 uint8
		v5 *byte
	)
	_ = v5
	var v6 *File
	var v7 *byte
	var v8 *byte
	var v9 uint8
	var v10 *File
	var v11 unsafe.Pointer
	var v12 int32
	var v13 int32
	var v14 int32
	var v15 HANDLE
	var v17 [12]byte
	var v20 [260]byte
	var FindFileData _WIN32_FIND_DATAA
	v2 = (*byte)(alloc.Calloc(1, 288))
	libc.MemSet(unsafe.Pointer(v2), 0, 288)
	libc.StrCpy(&v20[0], a1)
	v3 = libc.StrRChr(&v20[0], 46)
	if v3 != nil {
		*v3 = 0
	}
	v4 = *memmap.PtrUint8(0x587000, 155064)
	v5 = &v20[libc.StrLen(&v20[0])]
	*(*uint32)(unsafe.Pointer(v5)) = *memmap.PtrUint32(0x587000, 155060)
	*(*byte)(unsafe.Add(unsafe.Pointer(v5), 4)) = byte(v4)
	v6 = nox_fs_open(&v20[0])
	if v6 == nil {
		goto LABEL_27
	}
	v7 = libc.StrRChr(&v20[0], 46)
	if v7 != nil {
		*v7 = 0
	}
	v8 = &v20[libc.StrLen(&v20[0])+1]
	v9 = *memmap.PtrUint8(0x587000, 155076)
	*(*uint32)(unsafe.Pointer(func() *byte {
		p := &v8
		*p = (*byte)(unsafe.Add(unsafe.Pointer(*p), -1))
		return *p
	}())) = *memmap.PtrUint32(0x587000, 155072)
	*(*byte)(unsafe.Add(unsafe.Pointer(v8), 4)) = byte(v9)
	v10 = nox_fs_open(&v20[0])
	*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), unsafe.Sizeof(uint32(0))*67))) = uint32(uintptr(unsafe.Pointer(v10)))
	if v10 == nil || nox_binfile_fread_raw_40ADD0(&v17[0], 12, 1, v6) != 1 {
		goto LABEL_27
	}
	*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), unsafe.Sizeof(uint32(0))*1))) = *(*uint32)(unsafe.Pointer(&v17[8]))
	if *(*int32)(unsafe.Pointer(&v17[8])) <= 0 {
		goto LABEL_14
	}
	v11 = alloc.Calloc(int(*(*uint32)(unsafe.Pointer(&v17[8]))), 36)
	v12 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), unsafe.Sizeof(uint32(0))*1))))
	*(*uint32)(unsafe.Pointer(v2)) = uint32(uintptr(v11))
	libc.MemSet(v11, 0, int((uint32(v12*36)>>2)*4))
	if *(*uint32)(unsafe.Pointer(&v17[4])) != 1 {
		if nox_binfile_fread_raw_40ADD0(*(**byte)(unsafe.Pointer(v2)), *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), unsafe.Sizeof(uint32(0))*1)))*36, 1, v6) == 1 {
			goto LABEL_14
		}
	LABEL_27:
		if v2 != nil {
			sub_4869C0(unsafe.Pointer(v2))
		}
		if v6 != nil {
			nox_fs_close(v6)
		}
		return nil
	}
	v13 = 0
	if *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), unsafe.Sizeof(uint32(0))*1))) > 0 {
		v14 = 0
		for nox_binfile_fread_raw_40ADD0((*byte)(unsafe.Pointer(uintptr(uint32(v14)+*(*uint32)(unsafe.Pointer(v2))))), 32, 1, v6) == 1 {
			v13++
			*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(v2)) + uint32(v14) + 32))) = 0
			v14 += 36
			if uint32(v13) >= *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), unsafe.Sizeof(uint32(0))*1))) {
				goto LABEL_14
			}
		}
		goto LABEL_27
	}
LABEL_14:
	nox_fs_close(v6)
	libc.Sort(*(*unsafe.Pointer)(unsafe.Pointer(v2)), *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), unsafe.Sizeof(uint32(0))*1))), 36, func(arg1 unsafe.Pointer, arg2 unsafe.Pointer) int32 {
		return int32(libc.StrCaseCmp((*byte)(arg1), (*byte)(arg2)))
	})
	*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), unsafe.Sizeof(uint32(0))*69))) = 0
	if a2 != nil {
		libc.StrCpy((*byte)(unsafe.Add(unsafe.Pointer(v2), 8)), a2)
		if *(*byte)(unsafe.Add(unsafe.Pointer(v2), libc.StrLen((*byte)(unsafe.Add(unsafe.Pointer(v2), 8)))+8)) == 92 {
			*(*byte)(unsafe.Add(unsafe.Pointer(v2), libc.StrLen((*byte)(unsafe.Add(unsafe.Pointer(v2), 8)))+8)) = 0
		}
		v15 = compatFindFirstFileA((*byte)(unsafe.Add(unsafe.Pointer(v2), 8)), &FindFileData)
		if uintptr(v15) != uintptr(math.MaxUint32) {
			if FindFileData.dwFileAttributes&16 != 0 {
			LABEL_21:
				*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), unsafe.Sizeof(uint32(0))*69))) = 1
			} else {
				for compatFindNextFileA(v15, &FindFileData) != 0 {
					if FindFileData.dwFileAttributes&16 != 0 {
						goto LABEL_21
					}
				}
			}
			compatFindClose(v15)
		}
		if *(*byte)(unsafe.Add(unsafe.Pointer(v2), libc.StrLen((*byte)(unsafe.Add(unsafe.Pointer(v2), 8)))+8)) != 92 {
			*(*uint16)(unsafe.Add(unsafe.Pointer(v2), libc.StrLen((*byte)(unsafe.Add(unsafe.Pointer(v2), 8)))+8)) = *memmap.PtrUint16(0x587000, 155084)
		}
	}
	return v2
}
func sub_4869C0(lpMem unsafe.Pointer) {
	if *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(lpMem)), unsafe.Sizeof(uint32(0))*67))) != 0 {
		nox_fs_close(*((**File)(unsafe.Add(unsafe.Pointer((**File)(lpMem)), unsafe.Sizeof((*File)(nil))*67))))
	}
	if *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(lpMem)), unsafe.Sizeof(uint32(0))*68))) != 0 {
		nox_fs_close(*((**File)(unsafe.Add(unsafe.Pointer((**File)(lpMem)), unsafe.Sizeof((*File)(nil))*68))))
	}
	if *(*uint32)(lpMem) != 0 {
		*(*unsafe.Pointer)(lpMem) = nil
	}
	lpMem = nil
}
func sub_486A10(a1 int32, a2 unsafe.Pointer) uint32 {
	var (
		v2     unsafe.Pointer
		result uint32
	)
	v2 = libc.Search(a2, *(*unsafe.Pointer)(unsafe.Pointer(uintptr(a1))), *(*uint32)(unsafe.Pointer(uintptr(a1 + 4))), 36, func(arg1 unsafe.Pointer, arg2 unsafe.Pointer) int32 {
		return int32(libc.StrCaseCmp((*byte)(arg1), (*byte)(arg2)))
	})
	if v2 != nil {
		result = (uint32(uintptr(v2)) - *(*uint32)(unsafe.Pointer(uintptr(a1)))) / 36
	} else {
		result = math.MaxUint32
	}
	return result
}
func sub_486A50(a1 *uint32, a2 int32) *byte {
	var result *byte
	if a2 < 0 || uint32(a2) >= *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*1)) {
		result = (*byte)(memmap.PtrOff(0x587000, 155088))
	} else {
		result = (*byte)(unsafe.Pointer(uintptr(sub_4866D0(a1, a2))))
	}
	return result
}
func sub_486A80(a1 *uint32, a2 int32) int32 {
	var result int32
	if a2 < 0 || uint32(a2) >= *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*1)) {
		result = 0
	} else {
		result = int32(*(*uint32)(unsafe.Pointer(uintptr(*a1 + uint32(a2*36) + 20))))
	}
	return result
}
func sub_486AA0(a1 *uint32, a2 int32, a3 *uint32) int32 {
	var (
		v3     *uint32
		result int32
	)
	v3 = (*uint32)(unsafe.Pointer(uintptr(sub_4866D0(a1, a2))))
	*a3 = 4
	*(*uint32)(unsafe.Add(unsafe.Pointer(a3), unsafe.Sizeof(uint32(0))*2)) = *(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*6))
	*(*uint32)(unsafe.Add(unsafe.Pointer(a3), unsafe.Sizeof(uint32(0))*3)) = uint32(bool2int((*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*7))&1) != 0) + 1)
	*(*uint32)(unsafe.Add(unsafe.Pointer(a3), unsafe.Sizeof(uint32(0))*6)) = *(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*8))
	if *(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*7))&8 != 0 {
		result = 2
		*(*uint32)(unsafe.Add(unsafe.Pointer(a3), unsafe.Sizeof(uint32(0))*1)) = 2
		*(*uint32)(unsafe.Add(unsafe.Pointer(a3), unsafe.Sizeof(uint32(0))*4)) = 2
	} else {
		*(*uint32)(unsafe.Add(unsafe.Pointer(a3), unsafe.Sizeof(uint32(0))*1)) = 0
		result = bool2int((*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*7))&4) != 0) + 1
		*(*uint32)(unsafe.Add(unsafe.Pointer(a3), unsafe.Sizeof(uint32(0))*4)) = uint32(result)
	}
	return result
}
func sub_486B10(a1 *uint32, a2 int32, a3 *uint32) int32 {
	var (
		result int32
		v4     int32
	)
	result = sub_4866D0(a1, a2)
	v4 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a3), unsafe.Sizeof(uint32(0))*2)))
	*(*uint32)(unsafe.Pointer(uintptr(result + 28))) = 0
	*(*uint32)(unsafe.Pointer(uintptr(result + 24))) = uint32(v4)
	if *(*uint32)(unsafe.Add(unsafe.Pointer(a3), unsafe.Sizeof(uint32(0))*3)) == 2 {
		*(*uint32)(unsafe.Pointer(uintptr(result + 28))) = 1
	}
	if *(*uint32)(unsafe.Add(unsafe.Pointer(a3), unsafe.Sizeof(uint32(0))*4)) == 2 {
		*(*uint32)(unsafe.Pointer(uintptr(result + 28))) |= 4
	}
	return result
}
func sub_486B60(a1 int32, a2 int32) int32 {
	var (
		v2  int32
		v3  *File
		v6  *File
		v7  *File
		v8  int32
		v9  int32
		v10 int32
		v12 int32
		v13 [8]byte
		v14 [16]byte
		v15 [12]int32
	)
	v12 = 1
	v2 = sub_4866D0((*uint32)(unsafe.Pointer(uintptr(a1))), a2)
	sub_486E00(a1)
	v3 = *(**File)(unsafe.Pointer(uintptr(a1 + 268)))
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 280))) = uint32(uintptr(unsafe.Pointer(v3)))
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 284))) = *(*uint32)(unsafe.Pointer(uintptr(v2 + 20)))
	if nox_fs_fseek(v3, int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 16)))), stdio.SEEK_SET) != 0 {
		v12 = 0
	}
	if *(*uint32)(unsafe.Pointer(uintptr(v2 + 20))) == 0 {
		v12 = 0
	}
	if *(*uint32)(unsafe.Pointer(uintptr(a1 + 276))) == 0 {
		return v12
	}
	libc.StrCpy((*byte)(unsafe.Pointer(&v15[3])), (*byte)(unsafe.Pointer(uintptr(a1+8))))
	libc.StrCat((*byte)(unsafe.Pointer(&v15[3])), (*byte)(unsafe.Pointer(uintptr(v2))))
	libc.StrCat((*byte)(unsafe.Pointer(&v15[3])), CString(".wav"))
	v6 = nox_fs_open((*byte)(unsafe.Pointer(&v15[3])))
	v7 = v6
	v8 = 0
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 272))) = uint32(uintptr(unsafe.Pointer(v6)))
	if v6 == nil {
		return v12
	}
	if nox_binfile_fread_raw_40ADD0((*byte)(unsafe.Pointer(&v15[0])), 12, 1, v6) != 1 || uint32(v15[0]) != 0x46464952 || uint32(v15[2]) != 0x45564157 {
		stdio.Printf("error: '%s' is bad - cannot read\n", &v15[3])
		if *(*uint32)(unsafe.Pointer(uintptr(a1 + 272))) != 0 {
			nox_fs_close(*(**File)(unsafe.Pointer(uintptr(a1 + 272))))
			*(*uint32)(unsafe.Pointer(uintptr(a1 + 272))) = 0
		}
		return v12
	}
	if nox_binfile_fread_raw_40ADD0(&v13[0], 8, 1, v7) != 1 {
		goto LABEL_18
	}
	for {
		if *(*uint32)(unsafe.Pointer(&v13[0])) == 0x20746D66 {
			nox_binfile_fread_raw_40ADD0(&v14[0], 16, 1, v7)
			nox_fs_fseek(v7, int32(*(*uint32)(unsafe.Pointer(&v13[4]))-16), stdio.SEEK_CUR)
			goto LABEL_15
		}
		if *(*uint32)(unsafe.Pointer(&v13[0])) == 0x61746164 {
			break
		}
		nox_fs_fseek(v7, *(*int32)(unsafe.Pointer(&v13[4])), stdio.SEEK_CUR)
	LABEL_15:
		if nox_binfile_fread_raw_40ADD0(&v13[0], 8, 1, v7) != 1 {
			goto LABEL_18
		}
	}
	v8 = int32(*(*uint32)(unsafe.Pointer(&v13[4])))
LABEL_18:
	*(*uint32)(unsafe.Pointer(uintptr(v2 + 28))) = 2
	if int32(*(*uint16)(unsafe.Pointer(&v14[12])))/int32(*(*uint16)(unsafe.Pointer(&v14[2]))) == 2 {
		*(*uint32)(unsafe.Pointer(uintptr(v2 + 28))) = 6
	}
	if int32(*(*uint16)(unsafe.Pointer(&v14[2]))) == 2 {
		v9 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 28))))
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v9))), 0)) = uint8(int8(v9 | 1))
		*(*uint32)(unsafe.Pointer(uintptr(v2 + 28))) = uint32(v9)
	}
	*(*uint32)(unsafe.Pointer(uintptr(v2 + 24))) = *(*uint32)(unsafe.Pointer(&v14[4]))
	v10 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 272))))
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 284))) = uint32(v8)
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 280))) = uint32(v10)
	return 1
}
func sub_486DB0(a1 int32, a2 *byte, a3 int32) int32 {
	var (
		result int32
		v4     int32
	)
	if *(*uint32)(unsafe.Pointer(uintptr(a1 + 280))) == 0 {
		return 0
	}
	v4 = a3
	if a3 > *(*int32)(unsafe.Pointer(uintptr(a1 + 284))) {
		v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 284))))
	}
	if v4 <= 0 || (func() bool {
		result = nox_binfile_fread_raw_40ADD0(a2, 1, uint32(v4), *(**File)(unsafe.Pointer(uintptr(a1 + 280))))
		return result < 0
	}()) {
		result = 0
	}
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 284))) -= uint32(result)
	return result
}
func sub_486E00(a1 int32) *File {
	var result *File
	result = *(**File)(unsafe.Pointer(uintptr(a1 + 272)))
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 280))) = 0
	if result != nil {
		nox_fs_close(result)
		result = nil
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 272))) = 0
	}
	return result
}
func sub_486E30(a1 int32, a2 *uint32) int32 {
	var result int32
	*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*33)) = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 192)))++
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 212)))++
	nox_common_list_append_4258E0((*nox_list_item_t)(unsafe.Pointer(uintptr(a1+200))), (*nox_list_item_t)(unsafe.Pointer(a2)))
	result = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 212))) - 1)
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 212))) = uint32(result)
	if result < 0 {
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 212))) = 0
	}
	return result
}
func sub_486E90(a1 int32) int32 {
	var (
		v1     int32
		result int32
	)
	v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 132))))
	sub_425920((**uint32)(unsafe.Pointer(uintptr(a1))))
	*(*uint32)(unsafe.Pointer(uintptr(v1 + 192)))--
	*(*uint32)(unsafe.Pointer(uintptr(v1 + 212)))++
	sub_425920((**uint32)(unsafe.Pointer(uintptr(a1))))
	result = int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 212))) - 1)
	*(*uint32)(unsafe.Pointer(uintptr(v1 + 212))) = uint32(result)
	if result < 0 {
		*(*uint32)(unsafe.Pointer(uintptr(v1 + 212))) = 0
	}
	return result
}
func sub_486EF0() unsafe.Pointer {
	var (
		result unsafe.Pointer
		v1     *int32
		i      int32
	)
	result = *(*unsafe.Pointer)(unsafe.Pointer(&dword_5d4594_1193336))
	if dword_5d4594_1193336 != 0 {
		result = dword_587000_155144
		if *(*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(dword_587000_155144)) + 24))) == 0 {
			v1 = *(**int32)(unsafe.Pointer(uintptr(uint32(uintptr(dword_587000_155144)) + 12)))
			for i = int32(uint32(uintptr(dword_587000_155144)) + 12); v1 != (*int32)(unsafe.Pointer(uintptr(i))); v1 = (*int32)(unsafe.Pointer(uintptr(*v1))) {
				if (*(*int32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(int32(0))*3)) & 2) == 0 {
					result = unsafe.Pointer(uintptr((cgoAsFunc(*(*int32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(int32(0))*54)), (*func(*int32) int32)(nil)).(func(*int32) int32))(v1)))
				}
			}
		}
	}
	return result
}
func sub_486F30() int32 {
	nox_common_list_clear_425760((*nox_list_item_t)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_587000_155144)))))
	nox_common_list_clear_425760((*nox_list_item_t)(unsafe.Pointer(uintptr(uint32(uintptr(dword_587000_155144)) + 12))))
	*(*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(dword_587000_155144)) + 24))) = 0
	*memmap.PtrUint32(6112660, 1193340) = uint32(uintptr(dword_587000_155144)) + 32
	sub_4864A0((*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(dword_587000_155144)) + 32))))
	dword_5d4594_1193336 = 1
	return 0
}
func sub_486FA0(a1 int32) *uint32 {
	var (
		result *uint32
		v2     *uint32
		v3     int32
	)
	result = sub_486FE0(a1)
	v2 = result
	if result != nil {
		v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 12))))
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v3))), 0)) = uint8(int8(v3 | 1))
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 12))) = uint32(v3)
		sub_487050(v2)
		if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))&2 != 0 {
			*memmap.PtrUint32(6112660, 1193332) = 1
		}
		result = v2
	}
	return result
}
func sub_486FE0(a1 int32) *uint32 {
	var v1 *uint32
	v1 = (*uint32)(alloc.Calloc(1, 88))
	libc.MemSet(unsafe.Pointer(v1), 0, 88)
	sub_425770(unsafe.Pointer(v1))
	*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*4)) = 0
	*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*3)) = uint32(a1)
	if (cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(a1 + 20))), (*func(*uint32) int32)(nil)))(v1) == 0 {
		return v1
	}
	if v1 != nil {
		sub_487030(unsafe.Pointer(v1))
	}
	return nil
}
func sub_487030(lpMem unsafe.Pointer) {
	(cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(lpMem)), unsafe.Sizeof(uint32(0))*3))) + 24))), (*func(unsafe.Pointer))(nil)))(lpMem)
	*(*uint32)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(lpMem)), unsafe.Sizeof(uint32(0))*3))) + 12))) &= 0xFFFFFFFE
	lpMem = nil
}
func sub_487050(a1 *uint32) {
	nox_common_list_append_4258E0((*nox_list_item_t)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_587000_155144))))), (*nox_list_item_t)(unsafe.Pointer(a1)))
}
func sub_487070(lpMem unsafe.Pointer) {
	sub_487090((**uint32)(lpMem))
	sub_487030(lpMem)
	*memmap.PtrUint32(6112660, 1193332) = 0
}
func sub_487090(a1 **uint32) **uint32 {
	return sub_425920(a1)
}
func sub_4870A0() {
	var (
		v1 *int32
		v2 *int32
		v3 *int32
	)
	v1 = sub_4870E0((*int32)(unsafe.Pointer(&v3)))
	if v1 != nil {
		for {
			v2 = sub_487100(&v3)
			sub_487070(unsafe.Pointer(v1))
			v1 = v2
			if v2 == nil {
				break
			}
		}
	}
}
func sub_4870E0(a1 *int32) *int32 {
	var result *int32
	result = (*int32)(unsafe.Pointer(nox_common_list_getFirstSafe_425890((*nox_list_item_t)(unsafe.Pointer(*(**int32)(unsafe.Pointer(&dword_587000_155144)))))))
	*a1 = int32(uintptr(unsafe.Pointer(result)))
	return result
}
func sub_487100(a1 **int32) *int32 {
	if *a1 != nil {
		*a1 = (*int32)(unsafe.Pointer(nox_common_list_getNextSafe_4258A0((*nox_list_item_t)(unsafe.Pointer(*a1)))))
	}
	return *a1
}
func sub_487120() int32 {
	var (
		v0 int32
		i  *int32
		v3 *int32
	)
	v0 = 0
	for i = sub_4870E0((*int32)(unsafe.Pointer(&v3))); i != nil; i = sub_487100(&v3) {
		v0 += *(*int32)(unsafe.Add(unsafe.Pointer(i), unsafe.Sizeof(int32(0))*5))
	}
	return v0
}
func sub_487150(a1 int32, a2 unsafe.Pointer) *uint32 {
	var (
		v2 int32
		v3 *uint32
		v4 *uint32
		v6 int32
	)
	v2 = a1
	if a1 == -1 {
		v2 = 0
	}
	sub_487360(v2, (**int32)(unsafe.Pointer(&a1)), &v6)
	if a1 == 0 {
		return nil
	}
	v3 = *(**uint32)(unsafe.Pointer(uintptr(a1 + v6*4 + 24)))
	if v3 == nil {
		v4 = sub_4871C0(a1, v6, a2)
		v3 = v4
		if v4 == nil {
			return nil
		}
		*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*47)) = uint32(v2)
		sub_487310(v4)
	}
	*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*4))++
	return v3
}
func sub_4871C0(a1 int32, a2 int32, a3 unsafe.Pointer) *uint32 {
	var (
		v3 int32
		v4 *uint32
	)
	v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 12))))
	v4 = (*uint32)(alloc.Calloc(1, 264))
	libc.MemSet(unsafe.Pointer(v4), 0, 264)
	sub_425770(unsafe.Pointer(v4))
	*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*6)) = uint32(a2)
	*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*5)) = uint32(a1)
	*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*4)) = 0
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 16)))++
	*(*uint32)(unsafe.Pointer(uintptr(a1 + a2*4 + 24))) = uint32(uintptr(unsafe.Pointer(v4)))
	*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*64)) = *(*uint32)(unsafe.Pointer(uintptr(v3 + 36)))
	nox_common_list_clear_425760((*nox_list_item_t)(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*50)))))
	sub_4864A0((*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*22)))
	*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*53)) = 0
	*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*56)) = 33
	*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*60)) = 0
	*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*58)) = 0
	*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*62)) = 0
	*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*54)) = uint32(cgoFuncAddr(sub_4873C0))
	*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*57)) = 0
	*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*61)) = 0
	*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*59)) = 0
	*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*63)) = 0
	nullsub_10(uint32(uintptr(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*15))))))
	nullsub_10(uint32(uintptr(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*8))))))
	if a3 != nil {
		sub_487590(int32(uintptr(unsafe.Pointer(v4))), a3)
	}
	if (cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(v3 + 28))), (*func(*uint32) int32)(nil)))(v4) == 0 {
		return v4
	}
	if v4 != nil {
		sub_4872C0(unsafe.Pointer(v4))
	}
	return nil
}
func sub_4872C0(lpMem unsafe.Pointer) {
	var (
		v1 int32
		v2 int32
	)
	sub_487910(int32(uintptr(lpMem)), -1)
	(cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(lpMem)), unsafe.Sizeof(uint32(0))*5))) + 12))) + 32))), (*func(unsafe.Pointer))(nil)))(lpMem)
	*(*uint32)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(lpMem)), unsafe.Sizeof(uint32(0))*5))) + *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(lpMem)), unsafe.Sizeof(uint32(0))*6)))*4 + 24))) = 0
	v1 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(lpMem)), unsafe.Sizeof(uint32(0))*5))))
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 16))) - 1)
	*(*uint32)(unsafe.Pointer(uintptr(v1 + 16))) = uint32(v2)
	if v2 < 0 {
		*(*uint32)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(lpMem)), unsafe.Sizeof(uint32(0))*5))) + 16))) = 0
	}
	lpMem = nil
}
func sub_487310(a1 *uint32) int32 {
	var result int32
	*(*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(dword_587000_155144)) + 24)))++
	nox_common_list_append_4258E0((*nox_list_item_t)(unsafe.Pointer(uintptr(uint32(uintptr(dword_587000_155144))+12))), (*nox_list_item_t)(unsafe.Pointer(a1)))
	result = int32(*(*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(dword_587000_155144)) + 24))) - 1)
	*(*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(dword_587000_155144)) + 24))) = uint32(result)
	if result < 0 {
		*(*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(dword_587000_155144)) + 24))) = 0
	}
	return result
}
func sub_487360(a1 int32, a2 **int32, a3 *int32) *int32 {
	var (
		result *int32
		i      int32
		v5     int32
		v6     *int32
	)
	result = sub_4870E0((*int32)(unsafe.Pointer(&v6)))
	for i = a1; result != nil; result = sub_487100(&v6) {
		v5 = *(*int32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(int32(0))*5))
		if i < v5 {
			break
		}
		i -= v5
	}
	*a2 = result
	if result != nil {
		result = a3
		*a3 = i
	} else {
		*a3 = -1
	}
	return result
}
func sub_4873C0(a3 int32) int32 {
	var (
		v1  int32
		v3  int64
		v4  uint32
		v5  int32
		v6  bool
		v7  uint32
		v8  uint32
		v9  int32
		v10 uint32
		v11 *uint32
		v12 int32
		v13 int32
		v14 *uint32
		v15 int32
		v16 int32
		v17 int32
	)
	v1 = a3
	if *(*uint32)(unsafe.Pointer(uintptr(a3 + 212))) != 0 {
		return -2146304000
	}
	v3 = int64(platformTicks())
	v4 = *(*uint32)(unsafe.Pointer(uintptr(a3 + 248)))
	v5 = int32(v3)
	v6 = uint32(int32(v3)) < v4
	v7 = uint32(int32(v3 - int64(v4)))
	v8 = *(*uint32)(unsafe.Pointer(uintptr(a3 + 224)))
	v16 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v3))), unsafe.Sizeof(uint32(0))*1)))
	v9 = int32((*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v3))), unsafe.Sizeof(uint32(0))*1))) - (uint32(bool2int(v6)) + *(*uint32)(unsafe.Pointer(uintptr(a3 + 252)))))
	v10 = *(*uint32)(unsafe.Pointer(uintptr(a3 + 228)))
	if (((uint64(v9)) << 32) | uint64(v7)) >= (((uint64(v10)) << 32) | uint64(v8)) {
		*(*uint32)(unsafe.Pointer(uintptr(a3 + 232))) = v7
		*(*uint32)(unsafe.Pointer(uintptr(a3 + 236))) = uint32(v9)
		if *(*uint64)(unsafe.Pointer(uintptr(a3 + 240))) > (((uint64(v10))<<32)|uint64(v8))*10 {
			*(*uint32)(unsafe.Pointer(uintptr(a3 + 240))) = 0
			*(*uint32)(unsafe.Pointer(uintptr(a3 + 244))) = 0
		}
		if (((uint64(v9)) << 32) | uint64(v7)) > *(*uint64)(unsafe.Pointer(uintptr(a3 + 240))) {
			*(*uint32)(unsafe.Pointer(uintptr(a3 + 240))) = v7
			*(*uint32)(unsafe.Pointer(uintptr(a3 + 244))) = uint32(v9)
		}
		v11 = (*uint32)(unsafe.Pointer(uintptr(a3 + 88)))
		v15 = a3 + 88
		sub_486520((*uint32)(unsafe.Pointer(uintptr(a3 + 88))))
		if *(*uint32)(unsafe.Pointer(uintptr(a3 + 184))) != 0 {
			sub_486520(*(**uint32)(unsafe.Pointer(uintptr(a3 + 184))))
		}
		if *(*uint32)(unsafe.Pointer(uintptr(a3 + 184))) == 0 || (func() int32 {
			v17 = sub_486550(*(**uint8)(unsafe.Pointer(uintptr(a3 + 184))))
			return v17
		}()) == 0 {
			v17 = sub_486550((*uint8)(unsafe.Pointer(uintptr(v1 + 88))))
		}
		*(*uint32)(unsafe.Pointer(uintptr(v1 + 248))) = uint32(v5)
		*(*uint32)(unsafe.Pointer(uintptr(v1 + 252))) = uint32(v16)
		v12 = int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 200))))
		if v12 != v1+200 {
			for {
				v13 = int32(*(*uint32)(unsafe.Pointer(uintptr(v12))))
				if int32(*(*uint8)(unsafe.Pointer(uintptr(v12 + 124))))&1 != 0 && *(*uint32)(unsafe.Pointer(uintptr(v12 + 288))) != 0 {
					if (func() int32 {
						sub_486520((*uint32)(unsafe.Pointer(uintptr(v12 + 16))))
						return v17
					}()) != 0 || sub_486550((*uint8)(unsafe.Pointer(uintptr(v12+16)))) != 0 || *(*uint32)(unsafe.Pointer(uintptr(v12 + 116))) != 0 && sub_486550(*(**uint8)(unsafe.Pointer(uintptr(v12 + 116)))) != 0 || *(*uint32)(unsafe.Pointer(uintptr(v12 + 112))) != 0 && sub_486550(*(**uint8)(unsafe.Pointer(uintptr(v12 + 112)))) != 0 {
						sub_4BD840(v12)
						(cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v12 + 172))) + 32))), (*func(int32))(nil)))(v12)
					}
				}
				v12 = v13
				if v13 == v1+200 {
					break
				}
			}
			v11 = (*uint32)(unsafe.Pointer(uintptr(v15)))
		}
		v14 = *(**uint32)(unsafe.Pointer(uintptr(v1 + 184)))
		if v14 != nil {
			sub_486620(v14)
		}
		sub_486620(v11)
	}
	return 0
}
func sub_487590(a1 int32, a2 unsafe.Pointer) int32 {
	var result int32
	result = a1
	libc.MemCpy(unsafe.Pointer(uintptr(a1+60)), a2, 28)
	return result
}
func sub_4875B0(a1 *int32) *int32 {
	var result *int32
	result = (*int32)(unsafe.Pointer(nox_common_list_getFirstSafe_425890((*nox_list_item_t)(unsafe.Pointer(uintptr(uint32(uintptr(dword_587000_155144)) + 12))))))
	*a1 = int32(uintptr(unsafe.Pointer(result)))
	return result
}
func sub_4875D0(a1 **int32) *int32 {
	if *a1 != nil {
		*a1 = (*int32)(unsafe.Pointer(nox_common_list_getNextSafe_4258A0((*nox_list_item_t)(unsafe.Pointer(*a1)))))
	}
	return *a1
}
func sub_4875F0() int32 {
	var (
		v0     *int32
		v1     *int32
		result int32
		v3     *int32
	)
	*(*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(dword_587000_155144)) + 24)))++
	v0 = sub_4875B0((*int32)(unsafe.Pointer(&v3)))
	if v0 != nil {
		for {
			v1 = sub_4875D0(&v3)
			sub_487680(unsafe.Pointer(v0))
			v0 = v1
			if v1 == nil {
				break
			}
		}
	}
	result = int32(*(*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(dword_587000_155144)) + 24))) - 1)
	*(*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(dword_587000_155144)) + 24))) = uint32(result)
	if result < 0 {
		*(*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(dword_587000_155144)) + 24))) = 0
	}
	return result
}
func sub_487650(lpMem *uint32) {
	var v2 int32
	v2 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(lpMem), unsafe.Sizeof(uint32(0))*4)) - 1)
	*(*uint32)(unsafe.Add(unsafe.Pointer(lpMem), unsafe.Sizeof(uint32(0))*4)) = uint32(v2)
	if v2 < 0 {
		*(*uint32)(unsafe.Add(unsafe.Pointer(lpMem), unsafe.Sizeof(uint32(0))*4)) = 0
	}
	sub_487680(unsafe.Pointer(lpMem))
}
func sub_487680(lpMem unsafe.Pointer) {
	sub_4876A0((**uint32)(lpMem))
	sub_4872C0(lpMem)
}
func sub_4876A0(a1 **uint32) unsafe.Pointer {
	var result unsafe.Pointer
	*(*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(dword_587000_155144)) + 24)))++
	sub_425920(a1)
	result = unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(dword_587000_155144)) + 24))) - 1))
	*(*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(dword_587000_155144)) + 24))) = uint32(uintptr(result))
	if int32(uintptr(result)) < 0 {
		result = dword_587000_155144
		*(*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(dword_587000_155144)) + 24))) = 0
	}
	return result
}
func sub_4876E0() int32 {
	return -2147221504
}
func sub_4876F0(a1 int32, a2 int32) *uint32 {
	return sub_486320((*uint32)(unsafe.Pointer(uintptr(a1+88))), a2)
}
func sub_487710(a1 int32, a2 int32) int32 {
	return sub_486350(unsafe.Pointer(uintptr(a1+88)), a2)
}
func sub_487730(a1 int32) int32 {
	return int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 196))))
}
func sub_487740(a1 int32) int32 {
	return int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 192))))
}
func sub_487750(a1 int32) *uint32 {
	var (
		v1 *uint32
		v2 *uint32
	)
	if *(*int32)(unsafe.Pointer(uintptr(a1 + 192))) >= *(*int32)(unsafe.Pointer(uintptr(a1 + 196))) {
		return nil
	}
	v1 = sub_4BD720(a1)
	v2 = v1
	if v1 == nil {
		return nil
	}
	sub_486E30(a1, v1)
	return v2
}
func sub_487790(a1 int32, a2 int32) int32 {
	var (
		v2 int32
		v3 int32
	)
	v2 = 0
	if sub_487750(a1) != nil {
		v3 = a2
		for {
			v2++
			v3--
			if !(v3 != 0 && sub_487750(a1) != nil) {
				break
			}
		}
	}
	return v2
}
func sub_4877D0(a1 int32, a2 *int32) *int32 {
	var result *int32
	result = (*int32)(unsafe.Pointer(nox_common_list_getFirstSafe_425890((*nox_list_item_t)(unsafe.Pointer(uintptr(a1 + 200))))))
	*a2 = int32(uintptr(unsafe.Pointer(result)))
	return result
}
func sub_4877F0(a1 **int32) *int32 {
	if *a1 != nil {
		*a1 = (*int32)(unsafe.Pointer(nox_common_list_getNextSafe_4258A0((*nox_list_item_t)(unsafe.Pointer(*a1)))))
	}
	return *a1
}
func sub_487810(a1 int32, a2 int32) *int32 {
	var (
		v2     uint32
		v3     int32
		v4     *int32
		result *int32
		v6     int32
		v7     uint32
		v8     int32
		v9     *int32
	)
	v2 = math.MaxUint32
	if a2 == -1 {
		a2 = 1
	}
	v3 = math.MaxInt8
	v4 = nil
	v8 = math.MaxInt8
	v9 = nil
	for result = sub_4877D0(a1, &a1); result != nil; result = sub_4877F0((**int32)(unsafe.Pointer(&a1))) {
		if *(*int32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(int32(0))*3)) == a2 {
			if (*(*int32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(int32(0))*31)) & 21) == 0 {
				return result
			}
			v6 = *(*int32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(int32(0))*30))
			if *(*int32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(int32(0))*31))&1 != 0 {
				if v6 >= v3 {
					if v6 == v3 {
						v7 = uint32(*(*int32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(int32(0))*45)))
						if v7 < v2 && v2-v7 >= 1638 {
							v3 = *(*int32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(int32(0))*30))
							v4 = result
							v2 = uint32(*(*int32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(int32(0))*45)))
						}
					}
				} else {
					v2 = uint32(*(*int32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(int32(0))*45)))
					v3 = *(*int32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(int32(0))*30))
					v4 = result
				}
			} else if v6 < v8 {
				v8 = *(*int32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(int32(0))*30))
				v9 = result
			}
		}
	}
	result = v9
	if v9 == nil || v8 > v3 {
		result = v4
	}
	return result
}
func sub_4878E0(a1 int32, a2 int32) *int32 {
	var (
		v2 *int32
		v3 *int32
	)
	v2 = sub_487810(a1, 1)
	v3 = v2
	if v2 != nil {
		sub_4BDA10(int32(uintptr(unsafe.Pointer(v2))), a2)
	}
	return v3
}
func sub_487910(a1 int32, a2 int32) int32 {
	var (
		v2 *int32
		v3 int32
		v4 *int32
	)
	v2 = sub_4877D0(a1, &a1)
	if v2 == nil {
		return 0
	}
	v3 = a2
	for {
		v4 = sub_4877F0((**int32)(unsafe.Pointer(&a1)))
		if v3 == -1 || *(*int32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(int32(0))*3)) == v3 {
			sub_4BDA60(unsafe.Pointer(v2))
		}
		v2 = v4
		if v4 == nil {
			break
		}
	}
	return 0
}
func sub_487970(a1 int32, a2 int32) *int32 {
	var (
		result *int32
		v3     *int32
		v4     int32
		v5     *int32
	)
	result = sub_4877D0(a1, &a1)
	v3 = result
	if result != nil {
		v4 = a2
		for {
			result = sub_4877F0((**int32)(unsafe.Pointer(&a1)))
			v5 = result
			if v4 == -1 || *(*int32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(int32(0))*3)) == v4 {
				result = (*int32)(unsafe.Pointer(uintptr(sub_4BDA80(int32(uintptr(unsafe.Pointer(v3)))))))
			}
			v3 = v5
			if v5 == nil {
				break
			}
		}
	}
	return result
}
func sub_4879C0(a1 int32, a2 int32) *int32 {
	var (
		result *int32
		v3     int32
	)
	result = sub_4877D0(a1, &a1)
	if result != nil {
		v3 = a2
		for {
			if *(*int32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(int32(0))*31))&1 != 0 && (v3 == -1 || *(*int32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(int32(0))*3)) == v3) {
				sub_4BDAC0(int32(uintptr(unsafe.Pointer(result))))
			}
			result = sub_4877F0((**int32)(unsafe.Pointer(&a1)))
			if result == nil {
				break
			}
		}
	}
	return result
}
func sub_487A10(a1 int32, a2 int32) *int32 {
	var (
		result *int32
		v3     int32
	)
	result = sub_4877D0(a1, &a1)
	if result != nil {
		v3 = a2
		for {
			if *(*int32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(int32(0))*31))&4 != 0 && (v3 == -1 || *(*int32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(int32(0))*3)) == v3) {
				sub_4BDAF0(int32(uintptr(unsafe.Pointer(result))))
			}
			result = sub_4877F0((**int32)(unsafe.Pointer(&a1)))
			if result == nil {
				break
			}
		}
	}
	return result
}
func sub_487A60(a1 int32) int32 {
	return (cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(a1 + 216))), (*func(int32) int32)(nil)))(a1)
}
func sub_487BC0(a1 uint32) *uint32 {
	var (
		v1 *uint32
		v2 unsafe.Pointer
	)
	v1 = (*uint32)(alloc.Calloc(1, 28))
	libc.MemSet(unsafe.Pointer(v1), 0, 28)
	sub_487C30(v1)
	*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*1)) = a1
	v2 = alloc.Calloc(1, int(a1))
	*v1 = uint32(uintptr(v2))
	if v2 != nil {
		return v1
	}
	if v1 != nil {
		sub_487C10(unsafe.Pointer(v1))
	}
	return nil
}
func sub_487C10(lpMem unsafe.Pointer) {
	if *(*uint32)(lpMem) != 0 {
		*(*unsafe.Pointer)(lpMem) = nil
	}
	lpMem = nil
}
func sub_487C30(a1 *uint32) {
	*a1 = 0
	*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*1)) = 0
	*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*5)) = 0
	*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*6)) = 0
	nox_common_list_clear_425760((*nox_list_item_t)(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*2)))))
}
func sub_487C50(a1 int32, a2 *uint32) int32 {
	var result int32
	nox_common_list_append_4258E0((*nox_list_item_t)(unsafe.Pointer(uintptr(a1+8))), (*nox_list_item_t)(unsafe.Pointer(a2)))
	result = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*4)) + *(*uint32)(unsafe.Pointer(uintptr(a1 + 4))))
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 4))) = uint32(result)
	*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*5)) = uint32(a1)
	return result
}
func sub_487C80(a1 int32) int32 {
	return int32(uintptr(unsafe.Pointer(nox_common_list_getNext_425940((*nox_list_item_t)(unsafe.Pointer(uintptr(a1 + 8)))))))
}
func sub_487CB0(a1 int32, a2 int32) int32 {
	return int32(uint32(a2) * *(*uint32)(unsafe.Pointer(uintptr(a1 + 20))) / 1000)
}
func sub_487CD0(a1 int32, a2 int32) int64 {
	return int64(uint32(a2*1000) / *(*uint32)(unsafe.Pointer(uintptr(a1 + 20))))
}
func sub_487D00(a1 *uint32) int32 {
	var (
		v1     int32
		result int32
	)
	v1 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*1)))
	result = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*2)) * *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*3)) * *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*4)))
	*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*5)) = uint32(result)
	if v1 == 1 {
		result >>= 2
		*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*5)) = uint32(result)
	}
	return result
}
func sub_487D30(a1 *uint32, a2 int32, a3 int32) *uint32 {
	var result *uint32
	*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*4)) = uint32(a3)
	*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*3)) = uint32(a2)
	result = (*uint32)(sub_425770(unsafe.Pointer(a1)))
	*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*5)) = 0
	return result
}
func sub_487D60(a1 int32) int32 {
	var result int32
	result = a1
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 20))) = 0
	return result
}
func nox_xxx_wndEditProc_487D70_key(a1 *uint32, v4 int32, a3 int32, a4 int32) int32 {
	switch a3 {
	case 1:
		fallthrough
	case 58:
		fallthrough
	case 59:
		fallthrough
	case 60:
		fallthrough
	case 61:
		fallthrough
	case 62:
		fallthrough
	case 63:
		fallthrough
	case 64:
		fallthrough
	case 65:
		fallthrough
	case 66:
		fallthrough
	case 67:
		fallthrough
	case 68:
		fallthrough
	case 87:
		fallthrough
	case 88:
		fallthrough
	case 199:
		fallthrough
	case 201:
		fallthrough
	case 207:
		fallthrough
	case 209:
		return 0
	case 14:
		fallthrough
	case 211:
		if a3 != 14 && (strMan.Lang() == 6 || strMan.Lang() == 8) {
			break
		}
		if a4 != 2 {
			return 1
		}
		if int32(*(*uint16)(unsafe.Pointer(uintptr(v4 + 1054)))) == 0 {
			var v6 int16 = int16(*(*uint16)(unsafe.Pointer(uintptr(v4 + 1052))))
			if int32(v6) == 0 {
				return 1
			}
			var v7 uint16 = uint16(int16(int32(v6) - 1))
			*(*uint16)(unsafe.Pointer(uintptr(v4 + 1052))) = v7
			*(*uint16)(unsafe.Pointer(uintptr(v4 + int32(v7)*2))) = 0
			return 1
		}
	case 15:
		fallthrough
	case 205:
		fallthrough
	case 208:
		if a4 != 2 {
			return 1
		}
		nox_xxx_wndRetNULL_46A8A0()
		return 1
	case 28:
		fallthrough
	case 156:
		if a4 != 2 || *(*uint32)(unsafe.Pointer(uintptr(v4 + 1044))) != 0 {
			return 1
		}
		(*nox_window)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*13))))).Func94(asWindowEvent(0x401F, int32(uintptr(unsafe.Pointer(a1))), 0))
		return 1
	case 200:
		fallthrough
	case 203:
		if a4 != 2 {
			return 1
		}
		nox_xxx_wndRetNULL_0_46A8B0()
		return 1
	default:
	}
	if strMan.Lang() == 6 || strMan.Lang() == 8 {
		if *(*uint32)(unsafe.Pointer(uintptr(v4 + 1036))) == 0 && *(*uint32)(unsafe.Pointer(uintptr(v4 + 1032))) == 0 && *(*uint32)(unsafe.Pointer(uintptr(v4 + 1028))) == 0 {
			var v8 *libc.WChar = nox_input_getStringBuffer_57011C()
			nox_wcscpy((*libc.WChar)(unsafe.Pointer(uintptr(v4+512))), v8)
			nox_input_freeStringBuffer_57011C(v8)
			*(*uint16)(unsafe.Pointer(uintptr(v4 + 1054))) = uint16(nox_wcslen((*libc.WChar)(unsafe.Pointer(uintptr(v4 + 512)))))
			if false {
				return 1
			}
			(*nox_window)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v4 + 1048)))))).Hide()
			return 1
		}
	}
	var v9 uint16 = nox_input_scanCodeToAlpha_47F950(uint16(int16(a3)))
	if int32(v9) == 0 || a4 != 2 {
		return 1
	}
	var v10 int32 = 0
	if *(*uint32)(unsafe.Pointer(uintptr(v4 + 1028))) != 0 {
		v10 = bool2int(unicode.IsDigit(rune(v9)))
		if v10 == 0 {
			return 1
		}
	} else if *(*uint32)(unsafe.Pointer(uintptr(v4 + 1032))) != 0 {
		v10 = bool2int(libc.IsAlnum(rune(v9)))
		if v10 == 0 {
			return 1
		}
	}
	var v11 int32 = int32(*(*uint16)(unsafe.Pointer(uintptr(v4 + 1052))))
	if int32(uint16(int16(v11))) >= int32(*(*int16)(unsafe.Pointer(uintptr(v4 + 1040))))-1 {
		return 1
	}
	*(*uint16)(unsafe.Pointer(uintptr(v4 + v11*2))) = v9
	*(*uint16)(unsafe.Pointer(uintptr(v4 + int32(func() uint16 {
		p := (*uint16)(unsafe.Pointer(uintptr(v4 + 1052)))
		*p++
		return *p
	}())*2))) = 0
	return 1
}
func nox_xxx_wndEditProc_487D70(a1p *nox_window, a2 int32, a3 int32, a4 int32) int32 {
	var (
		a1  *uint32 = (*uint32)(unsafe.Pointer(a1p))
		v4  int32
		v12 int32
		v13 int32
		v14 int32
		v15 int32
		v16 int32
		v17 int32
	)
	v4 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*8)))
	switch a2 {
	case 7:
		*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*9)) |= 2
		guiFocus((*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1)))))))
		return 1
	case 8:
		v15 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*11)))
		if v15&256 != 0 {
			(*nox_window)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*13))))).Func94(asWindowEvent(0x4000, int32(uintptr(unsafe.Pointer(a1))), 0))
		}
		return 1
	case 17:
		v12 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*11)))
		if (v12 & 256) == 0 {
			return 1
		}
		v16 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*13)))
		*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*9)) |= 2
		(*nox_window)(unsafe.Pointer(uintptr(v16))).Func94(asWindowEvent(0x4005, int32(uintptr(unsafe.Pointer(a1))), 0))
		guiFocus((*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1)))))))
		return 1
	case 18:
		v13 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*11)))
		if v13&256 != 0 {
			v14 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*9)))
			*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v14))), 0)) = uint8(int8(v14 & 253))
			v17 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*13)))
			*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*9)) = uint32(v14)
			(*nox_window)(unsafe.Pointer(uintptr(v17))).Func94(asWindowEvent(0x4006, int32(uintptr(unsafe.Pointer(a1))), 0))
		}
		return 1
	case 21:
		return nox_xxx_wndEditProc_487D70_key(a1, v4, a3, a4)
	default:
		return 0
	}
}
func nox_xxx_wndEditDrawNoImage_488160(a1 int32, a2 int32) int32 {
	var (
		v2    int32
		v3    int32
		v4    int32
		v5    int32
		v6    int32
		v7    int32
		v8    *libc.WChar
		v9    int16
		v10   uint32
		v11   uint32
		v12   *int16
		i     int32
		v14   *uint32
		v15   int32
		v16   int32
		v17   bool
		v19   int32
		v20   int32
		v21   int32
		v22   int32
		v23   *libc.WChar
		v24   int32
		v25   *libc.WChar
		v26   *int16
		yTop  int32
		xLeft int32
		v29   int32
		v30   int32
		v31   [256]int16
	)
	v2 = a1
	v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 20))))
	v25 = *(**libc.WChar)(unsafe.Pointer(uintptr(a2 + 28)))
	v23 = *(**libc.WChar)(unsafe.Pointer(uintptr(a1 + 32)))
	v22 = 0
	v24 = 0
	*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v23))), unsafe.Sizeof(uint32(0))*261))) = 0
	v30 = v3
	v26 = &v31[0]
	nox_client_wndGetPosition_46AA60((*nox_window)(unsafe.Pointer(uintptr(a1))), (*uint32)(unsafe.Pointer(&xLeft)), (*uint32)(unsafe.Pointer(&yTop)))
	v4 = xLeft
	v21 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 8))))
	v5 = nox_xxx_guiFontHeightMB_43F320(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a2 + 200))))))
	v6 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 12))))
	v29 = v5
	v7 = yTop + v6/2 - v5/2
	if ((*(*uint32)(unsafe.Pointer(uintptr(a1 + 4))) >> 13) & 1) == 1 {
		nox_draw_enableTextSmoothing_43F670(1)
	}
	if (*(*uint32)(unsafe.Pointer(uintptr(a1 + 4)))>>3)&1 != 0 {
		if int32(*(*uint8)(unsafe.Pointer(uintptr(a2))))&2 != 0 {
			v25 = *(**libc.WChar)(unsafe.Pointer(uintptr(a2 + 36)))
		}
	} else {
		v30 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 44))))
	}
	if int32(*(*uint16)(unsafe.Pointer(uintptr(a2 + 72)))) != 0 {
		nox_xxx_drawGetStringSize_43F840(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a2 + 200))))), (*libc.WChar)(unsafe.Pointer(uintptr(a2+72))), &v22, nil, 0)
		nox_xxx_drawSetTextColor_434390(int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 68)))))
		nox_xxx_drawStringWrap_43FAF0(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a2 + 200))))), (*libc.WChar)(unsafe.Pointer(uintptr(a2+72))), v4+2, v7, v21, 0)
		v4 += v22 + 6
		v21 += int32(-6 - v22)
	}
	v8 = v23
	v9 = int16(*(*libc.WChar)(unsafe.Add(unsafe.Pointer(v23), unsafe.Sizeof(libc.WChar(0))*521)))
	if int32(v9) > 0 && v21 > int32(v9) {
		v21 = int32(v9)
		v4 = int32(uint32(xLeft) + *(*uint32)(unsafe.Pointer(uintptr(a1 + 8))) - uint32(v9))
	}
	if uint32(v30) != 0x80000000 {
		nox_client_drawSetColor_434460(v30)
		nox_client_drawRectFilledOpaque_49CE30(v4+1, yTop+1, v21-2, int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 12)))-2))
		v8 = v23
	}
	if uintptr(unsafe.Pointer(v25)) != uintptr(0x80000000) {
		nox_client_drawSetColor_434460(int32(uintptr(unsafe.Pointer(v25))))
		nox_client_drawBorderLines_49CC70(v4, yTop, v21, int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 12)))))
		v8 = v23
	}
	if *(*uint32)(unsafe.Pointer(uintptr(a2 + 68))) != 0x80000000 {
		if *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v8))), unsafe.Sizeof(uint32(0))*256))) != 0 {
			v10 = nox_wcslen(v8)
			v11 = 0
			if int32(v10) > 0 {
				memset32((*uint32)(unsafe.Pointer(&v31[0])), 0x2A002A, v10>>1)
				v12 = &v31[(v10>>1)*2]
				for i = int32(v10 & 1); i != 0; i-- {
					*v12 = 42
					v12 = (*int16)(unsafe.Add(unsafe.Pointer(v12), unsafe.Sizeof(int16(0))*1))
				}
				v2 = a1
				v11 = v10
			}
			v31[v11] = 0
		} else {
			nox_wcscpy((*libc.WChar)(unsafe.Pointer(&v31[0])), v8)
		}
		nox_xxx_drawGetStringSize_43F840(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a2 + 200))))), (*libc.WChar)(unsafe.Pointer((*uint16)(unsafe.Pointer(&v31[0])))), &v22, nil, 0)
		v19 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 200))))
		v25 = (*libc.WChar)(unsafe.Add(unsafe.Pointer(v23), unsafe.Sizeof(libc.WChar(0))*256))
		nox_xxx_drawGetStringSize_43F840(unsafe.Pointer(uintptr(v19)), (*libc.WChar)(unsafe.Add(unsafe.Pointer(v23), unsafe.Sizeof(libc.WChar(0))*256)), &v24, nil, 0)
		if ((*(*uint32)(unsafe.Pointer(uintptr(v2 + 4)))>>14)&1) == 1 && v24+v22 > 0 && v21 >= 10 && v24+v22+10 > v21 {
			for {
				v20 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 200))))
				v26 = (*int16)(unsafe.Add(unsafe.Pointer(v26), unsafe.Sizeof(int16(0))*1))
				nox_xxx_drawGetStringSize_43F840(unsafe.Pointer(uintptr(v20)), (*libc.WChar)(unsafe.Pointer((*uint16)(unsafe.Pointer(v26)))), &v22, nil, 0)
				if v24+v22+10 <= v21 {
					break
				}
			}
		}
		v14 = (*uint32)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v23))), unsafe.Sizeof(uint32(0))*262))))))
		if v14 != nil {
			(*nox_window)(unsafe.Pointer(v14)).SetPos(image.Pt(v4+v22, v29+v7))
		}
		nox_xxx_drawSetTextColor_434390(int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 68)))))
		nox_xxx_drawStringWrap_43FAF0(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a2 + 200))))), (*libc.WChar)(unsafe.Pointer(v26)), v4+5, v7, 0, 0)
		v15 = int32(nox_color_rgb_4344A0(192, 0, 192))
		nox_xxx_drawSetTextColor_434390(v15)
		nox_xxx_drawStringWrap_43FAF0(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a2 + 200))))), v25, v4+v22+5, v7, 0, 0)
		v16 = v4 + v22 + v24 + 5
		if unsafe.Pointer(uintptr(v2)) == unsafe.Pointer(nox_xxx_wndGetFocus_46B4F0()) {
			v17 = (int32(func() uint8 {
				p := memmap.PtrUint8(6112660, 1193344)
				x := *p
				*p++
				return x
			}()) & 8) == 0
			if !v17 {
				nox_client_drawSetColor_434460(int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 68)))))
				nox_client_drawRectFilledOpaque_49CE30(v16, v7, 2, v29)
			}
		}
	}
	nox_draw_enableTextSmoothing_43F670(0)
	return 1
}
func nox_gui_newEntryField_488500(a1p *nox_window, a2 int32, a3 int32, a4 int32, a5 int32, a6 int32, a7 int32, a8 *libc.WChar) *nox_window {
	var (
		a1     int32 = int32(uintptr(unsafe.Pointer(a1p)))
		v8     *uint32
		v9     bool
		v10    *int32
		result *uint32
		v12    *uint32
		v13    int32
		v14    [56]byte
		v15    [332]byte
	)
	if int32(*(*uint8)(unsafe.Pointer(uintptr(a7 + 8))))&128 != 0 {
		v8 = (*uint32)(unsafe.Pointer(nox_window_new((*nox_window)(unsafe.Pointer(uintptr(a1))), a2, a3, a4, a5, a6, func(arg1 int32, arg2 int32, arg3 int32, arg4 int32) int32 {
			return nox_xxx_wndEditProcPre_488710(arg1, uint32(arg2), (*libc.WChar)(unsafe.Pointer(uintptr(arg3))), arg4)
		})))
		nox_xxx_wndEdit_488830(int32(uintptr(unsafe.Pointer(v8))))
		if v8 == nil {
			return (*nox_window)(unsafe.Pointer(v8))
		}
		if *(*uint32)(unsafe.Pointer(uintptr(a7 + 16))) == 0 {
			*(*uint32)(unsafe.Pointer(uintptr(a7 + 16))) = uint32(uintptr(unsafe.Pointer(v8)))
		}
		nox_gui_windowCopyDrawData_46AF80((*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v8)))))), unsafe.Pointer(uintptr(a7)))
		libc.MemSet(unsafe.Pointer(a8), 0, 256)
		libc.MemSet(unsafe.Pointer((*libc.WChar)(unsafe.Add(unsafe.Pointer(a8), unsafe.Sizeof(libc.WChar(0))*256))), 0, 256)
		*(*libc.WChar)(unsafe.Add(unsafe.Pointer(a8), unsafe.Sizeof(libc.WChar(0))*526)) = libc.WChar(nox_wcslen(a8))
		v9 = int32(int16(*(*libc.WChar)(unsafe.Add(unsafe.Pointer(a8), unsafe.Sizeof(libc.WChar(0))*520)))) < 256
		*(*libc.WChar)(unsafe.Add(unsafe.Pointer(a8), unsafe.Sizeof(libc.WChar(0))*527)) = 0
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a8))), unsafe.Sizeof(uint32(0))*261))) = 0
		if !v9 {
			*(*libc.WChar)(unsafe.Add(unsafe.Pointer(a8), unsafe.Sizeof(libc.WChar(0))*520)) = 256
		}
		v10 = (*int32)(alloc.Calloc(1, 1056))
		libc.MemCpy(unsafe.Pointer(v10), unsafe.Pointer(a8), 1056)
		*(*uint32)(unsafe.Add(unsafe.Pointer(v8), unsafe.Sizeof(uint32(0))*8)) = uint32(uintptr(unsafe.Pointer(v10)))
		if strMan.Lang() != 8 && strMan.Lang() != 6 {
			result = v8
			*(*int32)(unsafe.Add(unsafe.Pointer(v10), unsafe.Sizeof(int32(0))*262)) = 0
			return (*nox_window)(unsafe.Pointer(result))
		}
		*(*[332]byte)(unsafe.Pointer(&v15[0])) = [332]byte{}
		*(*uint32)(unsafe.Pointer(&v15[24])) = 0
		*(*uint32)(unsafe.Pointer(&v15[28])) = *(*uint32)(unsafe.Pointer(uintptr(a7 + 68)))
		*(*uint32)(unsafe.Pointer(&v15[36])) = 0x80000000
		*(*uint32)(unsafe.Pointer(&v15[52])) = 0x80000000
		*(*uint32)(unsafe.Pointer(&v15[68])) = *(*uint32)(unsafe.Pointer(&v15[28]))
		*(*[56]byte)(unsafe.Pointer(&v14[0])) = [56]byte{}
		*(*uint32)(unsafe.Pointer(&v14[8])) = 1
		*(*uint32)(unsafe.Pointer(&v14[12])) = 1
		*(*uint32)(unsafe.Pointer(&v15[48])) = 0
		*(*uint16)(unsafe.Pointer(&v14[2])) = 10
		*(*uint16)(unsafe.Pointer(&v14[0])) = 128
		*(*uint32)(unsafe.Pointer(&v14[4])) = 0
		*(*uint32)(unsafe.Pointer(&v14[16])) = 0
		*(*uint16)(unsafe.Pointer(&v15[72])) = 0
		*(*uint32)(unsafe.Pointer(&v15[8])) = 288
		v12 = (*uint32)(unsafe.Pointer(nox_gui_newScrollListBox_4A4310(nil, 0x44B0, 0, a6, 110, 119, int32(uintptr(unsafe.Pointer(&v15[0]))), (*nox_scrollListBox_data)(unsafe.Pointer((*int16)(unsafe.Pointer(&v14[0])))))))
		*(*int32)(unsafe.Add(unsafe.Pointer(v10), unsafe.Sizeof(int32(0))*262)) = int32(uintptr(unsafe.Pointer(v12)))
		if v12 != nil {
			nox_xxx_wndClearFlag_46AD80(int32(uintptr(unsafe.Pointer(v12))), 128)
			nox_xxx_wndListboxInit_4A3C00(*(*int32)(unsafe.Add(unsafe.Pointer(v10), unsafe.Sizeof(int32(0))*262)), int32(*(*uint32)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(v10), unsafe.Sizeof(int32(0))*262)) + 32)))))
			v13 = int32(nox_color_rgb_4344A0(0, 0, 0))
			nox_xxx_wndSetRectColor2MB_46AFE0((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(v10), unsafe.Sizeof(int32(0))*262))))), v13)
			return (*nox_window)(unsafe.Pointer(v8))
		}
	}
	return nil
}
func nox_xxx_wndEditProcPre_488710(a1 int32, a2 uint32, a3 *libc.WChar, a4 int32) int32 {
	var (
		v3 int32
		v4 int32
		v6 int32
		v7 int32
	)
	v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 32))))
	if a2 > 0x401D {
		if a2 == 0x401E {
			nox_wcsncpy((*libc.WChar)(unsafe.Pointer(uintptr(v3))), a3, math.MaxUint8)
			*(*uint16)(unsafe.Pointer(uintptr(v3 + 510))) = 0
			*(*uint16)(unsafe.Pointer(uintptr(v3 + 1052))) = uint16(nox_wcslen((*libc.WChar)(unsafe.Pointer(uintptr(v3)))))
			*(*uint16)(unsafe.Pointer(uintptr(v3 + 512))) = 0
			*(*uint16)(unsafe.Pointer(uintptr(v3 + 1054))) = 0
		}
		return 0
	}
	if a2 == 0x401D {
		return int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 32))))
	}
	if a2 == 2 {
		(*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(uintptr(v3 + 1048))))).Destroy()
		*(*unsafe.Pointer)(unsafe.Pointer(uintptr(a1 + 32))) = nil
		return 0
	}
	if a2 != 23 {
		return 0
	}
	if a3 != nil {
		dword_5d4594_1193352 = uint32(a1)
		nox_input_enableTextEdit_5700CA()
		v6 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 36))))
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v6))), 0)) = uint8(int8(v6 | 6))
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 36))) = uint32(v6)
	} else {
		nox_input_disableTextEdit_5700F6()
		dword_5d4594_1193352 = 0
		v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 36))))
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v4))), 0)) = uint8(int8(v4 & 249))
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 36))) = uint32(v4)
		(*nox_window)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v3 + 1048)))))).Hide()
		*(*uint16)(unsafe.Pointer(uintptr(v3 + 512))) = 0
		*(*uint16)(unsafe.Pointer(uintptr(v3 + 1054))) = 0
	}
	v7 = (*nox_window)(unsafe.Pointer(uintptr(a1))).ID()
	(*nox_window)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 52)))))).Func94(asWindowEvent(0x4003, int32(uintptr(unsafe.Pointer(a3))), v7))
	return 1
}
func nox_xxx_wndEdit_488830(a1 int32) int32 {
	var result int32
	result = a1
	if a1 != 0 {
		if int32(int8(*(*uint8)(unsafe.Pointer(uintptr(a1 + 4))))) >= 0 {
			result = (*nox_window)(unsafe.Pointer(uintptr(a1))).SetAllFuncs(func(arg1 int32, arg2 int32, arg3 int32, arg4 int32) int32 {
				return nox_xxx_wndEditProc_487D70((*nox_window)(unsafe.Pointer(uintptr(arg1))), arg2, arg3, arg4)
			}, func(arg1 *nox_window, arg2 unsafe.Pointer) int32 {
				return nox_xxx_wndEditDrawNoImage_488160(int32(uintptr(unsafe.Pointer(arg1))), int32(uintptr(arg2)))
			}, nil)
		} else {
			result = (*nox_window)(unsafe.Pointer(uintptr(a1))).SetAllFuncs(func(arg1 int32, arg2 int32, arg3 int32, arg4 int32) int32 {
				return nox_xxx_wndEditProc_487D70((*nox_window)(unsafe.Pointer(uintptr(arg1))), arg2, arg3, arg4)
			}, func(arg1 *nox_window, arg2 unsafe.Pointer) int32 {
				return nox_xxx_wndEditDrawWithImage_488870(int32(uintptr(unsafe.Pointer(arg1))), int32(uintptr(arg2)))
			}, nil)
		}
	}
	return result
}
func nox_xxx_wndEditDrawWithImage_488870(a1 int32, a2 int32) int32 {
	var (
		v2    int32
		v3    int32
		v4    int32
		v5    int32
		v6    int32
		v7    int32
		v8    uint32
		v9    uint32
		v10   *int16
		i     int32
		v12   int32
		v13   bool
		v15   int32
		xLeft int32
		v17   int32
		v18   *int16
		v19   int32
		v20   *libc.WChar
		v21   int32
		v22   *libc.WChar
		v23   int32
		v24   [256]int16
	)
	v2 = a1
	v20 = *(**libc.WChar)(unsafe.Pointer(uintptr(a2 + 24)))
	v22 = *(**libc.WChar)(unsafe.Pointer(uintptr(a1 + 32)))
	*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v22))), unsafe.Sizeof(uint32(0))*261))) = 0
	v18 = &v24[0]
	nox_client_wndGetPosition_46AA60((*nox_window)(unsafe.Pointer(uintptr(a1))), (*uint32)(unsafe.Pointer(&xLeft)), (*uint32)(unsafe.Pointer(&v21)))
	v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 8))))
	v4 = nox_xxx_guiFontHeightMB_43F320(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a2 + 200))))))
	v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 12))))
	v23 = v4
	v6 = v21 + v5/2 - v4/2
	if ((*(*uint32)(unsafe.Pointer(uintptr(a1 + 4))) >> 13) & 1) == 1 {
		nox_draw_enableTextSmoothing_43F670(1)
	}
	if (*(*uint32)(unsafe.Pointer(uintptr(a1 + 4)))>>3)&1 != 0 {
		v7 = int32(uintptr(unsafe.Pointer(v20)))
	} else {
		v7 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 48))))
	}
	if v7 != 0 {
		nox_client_drawImageAt_47D2C0((*nox_video_bag_image_t)(unsafe.Pointer(uintptr(v7))), xLeft, v21)
	}
	nox_xxx_drawSetTextColor_434390(int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 68)))))
	if int32(*(*uint16)(unsafe.Pointer(uintptr(a2 + 72)))) != 0 {
		nox_xxx_drawGetStringSize_43F840(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a2 + 200))))), (*libc.WChar)(unsafe.Pointer(uintptr(a2+72))), &v17, nil, 0)
		nox_xxx_drawStringWrap_43FAF0(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a2 + 200))))), (*libc.WChar)(unsafe.Pointer(uintptr(a2+72))), xLeft+2, v6, v3, 0)
		v3 += int32(-6 - v17)
		xLeft += v17 + 6
	}
	if *(*uint32)(unsafe.Pointer(uintptr(a2 + 68))) != 0x80000000 {
		if *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v22))), unsafe.Sizeof(uint32(0))*256))) != 0 {
			v8 = nox_wcslen(v22)
			v9 = 0
			if int32(v8) > 0 {
				memset32((*uint32)(unsafe.Pointer(&v24[0])), 0x2A002A, v8>>1)
				v10 = &v24[(v8>>1)*2]
				for i = int32(v8 & 1); i != 0; i-- {
					*v10 = 42
					v10 = (*int16)(unsafe.Add(unsafe.Pointer(v10), unsafe.Sizeof(int16(0))*1))
				}
				v2 = a1
				v9 = v8
			}
			v24[v9] = 0
		} else {
			nox_wcscpy((*libc.WChar)(unsafe.Pointer(&v24[0])), v22)
		}
		nox_xxx_drawGetStringSize_43F840(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a2 + 200))))), (*libc.WChar)(unsafe.Pointer((*uint16)(unsafe.Pointer(&v24[0])))), &v17, nil, 0)
		v20 = (*libc.WChar)(unsafe.Add(unsafe.Pointer(v22), unsafe.Sizeof(libc.WChar(0))*256))
		nox_xxx_drawGetStringSize_43F840(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a2 + 200))))), (*libc.WChar)(unsafe.Add(unsafe.Pointer(v22), unsafe.Sizeof(libc.WChar(0))*256)), &v19, nil, 0)
		if ((*(*uint32)(unsafe.Pointer(uintptr(v2 + 4)))>>14)&1) == 1 && v17+v19 > 0 && v17+v19+10 > v3 {
			for {
				v15 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 200))))
				v18 = (*int16)(unsafe.Add(unsafe.Pointer(v18), unsafe.Sizeof(int16(0))*1))
				nox_xxx_drawGetStringSize_43F840(unsafe.Pointer(uintptr(v15)), (*libc.WChar)(unsafe.Pointer((*uint16)(unsafe.Pointer(v18)))), &v17, nil, 0)
				if v19+v17+10 <= v3 {
					break
				}
			}
		}
		nox_xxx_drawSetTextColor_434390(int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 68)))))
		nox_xxx_drawStringWrap_43FAF0(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a2 + 200))))), (*libc.WChar)(unsafe.Pointer(v18)), xLeft+5, v6, v3, 0)
		v12 = int32(nox_color_rgb_4344A0(192, 0, 192))
		nox_xxx_drawSetTextColor_434390(v12)
		nox_xxx_drawStringWrap_43FAF0(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a2 + 200))))), v20, v17+xLeft+5, v6, v3, 0)
		xLeft += v17 + v19 + 5
		if unsafe.Pointer(uintptr(v2)) == unsafe.Pointer(nox_xxx_wndGetFocus_46B4F0()) {
			v13 = (int32(func() uint8 {
				p := memmap.PtrUint8(6112660, 1193344)
				x := *p
				*p++
				return x
			}()) & 8) == 0
			if !v13 {
				nox_client_drawSetColor_434460(int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 68)))))
				nox_client_drawRectFilledOpaque_49CE30(xLeft, v6, 2, v23)
			}
		}
	}
	nox_draw_enableTextSmoothing_43F670(0)
	return 1
}
func sub_488B60() int32 {
	var v0 **int32
	v0 = (**int32)(operator_new(4))
	if v0 != nil {
		dword_5d4594_1193348 = uint32(uintptr(unsafe.Pointer(v0)))
	} else {
		dword_5d4594_1193348 = 0
	}
	return 1
}
func sub_488BA0() int32 {
	var v0 unsafe.Pointer
	v0 = *(*unsafe.Pointer)(unsafe.Pointer(&dword_5d4594_1193348))
	if dword_5d4594_1193348 != 0 {
		operator_delete(v0)
	}
	dword_5d4594_1193348 = 0
	return 1
}
func nox_xxx_onChar_488BD0(a1 uint16) {
	var (
		v2 int32
		v3 int32
	)
	if dword_5d4594_1193348 != 0 {
		if dword_5d4594_1193352 != 0 {
			v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1193352 + 32))))
			if strMan.Lang() == 6 || strMan.Lang() == 8 {
				if *(*uint32)(unsafe.Pointer(uintptr(v2 + 1036))) == 0 {
					if *(*uint32)(unsafe.Pointer(uintptr(v2 + 1032))) == 0 {
						if *(*uint32)(unsafe.Pointer(uintptr(v2 + 1028))) == 0 {
							*(*uint32)(unsafe.Pointer(uintptr(v2 + 1044))) = 1
							switch a1 {
							case 7:
								fallthrough
							case 8:
								fallthrough
							case 9:
								fallthrough
							case 11:
								fallthrough
							case 12:
							case 10:
								fallthrough
							case 13:
								*(*uint32)(unsafe.Pointer(uintptr(v2 + 1044))) = 0
							default:
								v3 = int32(*(*uint16)(unsafe.Pointer(uintptr(v2 + 1052))))
								if int32(uint16(int16(v3))) < int32(*(*int16)(unsafe.Pointer(uintptr(v2 + 1040))))-1 {
									*(*uint16)(unsafe.Pointer(uintptr(v2 + v3*2))) = a1
									*(*uint16)(unsafe.Pointer(uintptr(v2 + int32(func() uint16 {
										p := (*uint16)(unsafe.Pointer(uintptr(v2 + 1052)))
										*p++
										return *p
									}())*2))) = 0
								}
								var v4 *libc.WChar = nox_input_getStringBuffer_57011C()
								nox_wcscpy((*libc.WChar)(unsafe.Pointer(uintptr(v2+512))), v4)
								nox_input_freeStringBuffer_57011C(v4)
								*(*uint16)(unsafe.Pointer(uintptr(v2 + 1054))) = uint16(nox_wcslen((*libc.WChar)(unsafe.Pointer(uintptr(v2 + 512)))))
								(*nox_window)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 1048)))))).Hide()
							}
						}
					}
				}
			}
		}
	}
}
func sub_4896C0() int32 {
	return int32(*memmap.PtrUint32(6112660, 1193356))
}
func sub_4896D0(a1 int32) int32 {
	var result int32
	result = int32(*memmap.PtrUint32(6112660, 1193356))
	*memmap.PtrUint32(6112660, 1193356) = uint32(a1)
	return result
}
func sub_4896E0() int32 {
	if dword_5d4594_1193360 != 0 {
		*(*unsafe.Pointer)(unsafe.Pointer(&dword_5d4594_1193360)) = nil
	}
	return 1
}
func sub_489870() int32 {
	var (
		v0 int32
		v1 *uint8
		v2 *uint32
		v3 *libc.WChar
		v4 uint32
		v5 int32
		v6 int8
		v7 int32
	)
	v0 = nox_xxx_check_flag_aaa_43AF70()
	v1 = (*uint8)(memmap.PtrOff(6112660, uint32(v0*44)+0x1235AC))
	if *memmap.PtrUint32(6112660, uint32(v0*4)+0x12359C) == 2 {
		*(*uint32)(unsafe.Pointer(v1)) = ((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1193380)))).ChildByID(0x272C).draw_data.field_0 >> 2) & 1
		v2 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1193380)))).ChildByID(0x272F)))
		v3 = (*libc.WChar)(unsafe.Pointer(uintptr((*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v2)))))).Func94(asWindowEvent(0x401D, 0, 0)))))
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v1))), unsafe.Sizeof(uint32(0))*4))) = uint32(nox_wcstol(v3, nil, 10))
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v1))), unsafe.Sizeof(uint32(0))*1))) = ((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1193380)))).ChildByID(0x272D).draw_data.field_0 >> 2) & 1
		v4 = (*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1193380)))).ChildByID(10030).draw_data.field_0
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v1))), unsafe.Sizeof(uint32(0))*3))) = 0
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v1))), unsafe.Sizeof(uint32(0))*2))) = (v4 >> 2) & 1
		if (*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1193380)))).ChildByID(0x271F).draw_data.field_0&4 != 0 {
			v5 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v1))), unsafe.Sizeof(uint32(0))*3))))
			*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v5))), 0)) = uint8(int8(v5 | 128))
			*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v1))), unsafe.Sizeof(uint32(0))*3))) = uint32(v5)
			v6 = int8(*((*uint8)(unsafe.Pointer(&(*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1193380)))).ChildByID(0x2720).draw_data.field_0))))
			v7 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v1))), unsafe.Sizeof(uint32(0))*3))))
			if int32(v6)&4 != 0 {
				*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v7))), 0)) = uint8(int8(v7 | 1))
			} else {
				*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v7))), 0)) = uint8(int8(v7 | 2))
			}
			*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v1))), unsafe.Sizeof(uint32(0))*3))) = uint32(v7)
		}
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v1))), unsafe.Sizeof(uint32(0))*5))) = ((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1193380)))).ChildByID(0x271E).draw_data.field_0 >> 2) & 1
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v1))), unsafe.Sizeof(uint32(0))*10))) = ((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1193380)))).ChildByID(0x2722).draw_data.field_0 >> 2) & 1
	}
	return (*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1193380))))).Hide()
}
func nox_xxx_checkSomeFlagsOnJoin_4899C0(srv *nox_gui_server_ent_t) int32 {
	var (
		a1  int32 = int32(uintptr(unsafe.Pointer(srv)))
		v1  int32
		v2  int32
		v3  int32
		v4  *uint8
		v5  int32
		v6  int32
		v7  uint32
		v9  int8
		v10 uint8
		v11 uint8
		v12 int8
		v13 [15]int32
		v14 uint8
		v15 uint8
	)
	v1 = nox_xxx_check_flag_aaa_43AF70()
	v2 = v1 * 11
	v3 = int32(*memmap.PtrUint32(6112660, uint32(v1*4)+0x12359C))
	v4 = (*uint8)(memmap.PtrOff(6112660, uint32(v2*4)+0x1235AC))
	if v3 == 0 {
		return 1
	}
	v5 = v3 - 1
	if v5 != 0 {
		if v5 == 1 {
			v6 = a1
			if *(*uint32)(unsafe.Pointer(v4)) != 0 {
				v7 = *(*uint32)(unsafe.Pointer(uintptr(a1 + 96)))
				if v7 > *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v4))), unsafe.Sizeof(uint32(0))*4))) && v7 != 9999 {
					return 0
				}
			}
			if *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v4))), unsafe.Sizeof(uint32(0))*1))) != 0 && int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 100))))&16 != 0 {
				return 0
			}
			if *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v4))), unsafe.Sizeof(uint32(0))*2))) != 0 && int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 100))))&32 != 0 {
				return 0
			}
			v9 = int8(*(*uint8)(unsafe.Pointer(uintptr(a1 + 102))))
			if int32(v9) < 0 && *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v4))), unsafe.Sizeof(uint32(0))*3))) > uint32(int32(v9)&math.MaxInt8) {
				return 0
			}
			if *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v4))), unsafe.Sizeof(uint32(0))*5))) != 0 {
				libc.StrCpy((*byte)(unsafe.Pointer(&v13[0])), (*byte)(unsafe.Pointer(uintptr(a1+111))))
				sub_57A1E0(&v13[0], nil, nil, 5, int16(*(*uint16)(unsafe.Pointer(uintptr(a1 + 163)))))
				v10 = 0
				v14 = 0
				for uint32(v13[int32(v14)+6]) == *(*uint32)(unsafe.Pointer(uintptr(int32(v14)*4 + v6 + 135))) {
					v14 = func() uint8 {
						p := &v10
						*p++
						return *p
					}()
					if int32(v10) >= 5 {
						v11 = 0
						v15 = 0
						for int32(*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v13[11]))), v15)))) == int32(*(*uint8)(unsafe.Pointer(uintptr(int32(v15) + v6 + 155)))) {
							v15 = func() uint8 {
								p := &v11
								*p++
								return *p
							}()
							if int32(v11) >= 4 {
								if uint32(v13[12]) == *(*uint32)(unsafe.Pointer(uintptr(v6 + 159))) {
									goto LABEL_26
								}
								return 0
							}
						}
						return 0
					}
				}
				return 0
			}
		LABEL_26:
			if *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v4))), unsafe.Sizeof(uint32(0))*10))) != 0 && *(*uint32)(unsafe.Pointer(uintptr(v6 + 48))) != NOX_CLIENT_VERS_CODE {
				return 0
			}
		}
		return 1
	}
	v12 = int8(*(*uint8)(unsafe.Pointer(uintptr(a1 + 100))))
	if int32(v12)&16 != 0 {
		return 0
	}
	if int32(v12)&32 != 0 {
		return 0
	}
	return bool2int(*(*uint32)(unsafe.Pointer(uintptr(a1 + 48))) == NOX_CLIENT_VERS_CODE)
}
func sub_489B80(a1 int32) *uint32 {
	var (
		result *uint32
		v2     int32
		v3     *uint8
		v4     *uint32
		v5     *uint32
		v6     *uint32
		v7     *uint32
		v8     *uint32
		v9     *uint32
		v10    *uint32
		v11    int32
		v12    int32
		v13    [16]libc.WChar
	)
	result = (*uint32)(unsafe.Pointer(newWindowFromFile("filter.wnd", unsafe.Pointer(cgoFuncAddr(nox_xxx_windowMplayFilterProc_489E70)))))
	dword_5d4594_1193380 = uint32(uintptr(unsafe.Pointer(result)))
	if result != nil {
		dword_5d4594_1193384 = uint32(uintptr(unsafe.Pointer((*nox_window)(unsafe.Pointer(result)).ChildByID(0x271C))))
		v2 = nox_xxx_check_flag_aaa_43AF70()
		v3 = (*uint8)(memmap.PtrOff(6112660, uint32(v2*44)+0x1235AC))
		sub_46B120((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1193380)))), (*nox_window)(unsafe.Pointer(uintptr(a1))))
		sub_46B120((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1193384)))), (*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1193380))))))
		(*(*int32)(unsafe.Pointer(&dword_5d4594_1193384))).setFunc94(func(arg1 int32, arg2 int32, arg3 int32, arg4 int32) int32 {
			return nox_xxx_windowMplayFilterProc_489E70(arg1, arg2, (*int32)(unsafe.Pointer(uintptr(arg3))), arg4)
		})
		if *(*uint32)(unsafe.Pointer(v3)) != 0 {
			v4 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1193380)))).ChildByID(0x272C)))
			*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*9)) |= 4
		}
		nox_swprintf(&v13[0], libc.CWString("%d"), *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v3))), unsafe.Sizeof(uint32(0))*4))))
		v5 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1193380)))).ChildByID(0x272F)))
		(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v5)))))).Func94(asWindowEvent(0x401E, int32(uintptr(unsafe.Pointer(&v13[0]))), -1))
		if *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v3))), unsafe.Sizeof(uint32(0))*1))) != 0 {
			v6 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1193380)))).ChildByID(0x272D)))
			*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*9)) |= 4
		}
		if *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v3))), unsafe.Sizeof(uint32(0))*2))) != 0 {
			v7 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1193380)))).ChildByID(10030)))
			*(*uint32)(unsafe.Add(unsafe.Pointer(v7), unsafe.Sizeof(uint32(0))*9)) |= 4
		}
		if *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v3))), unsafe.Sizeof(uint32(0))*3))) != 0 {
			v8 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1193380)))).ChildByID(0x271F)))
			*(*uint32)(unsafe.Add(unsafe.Pointer(v8), unsafe.Sizeof(uint32(0))*9)) |= 4
		}
		if int32(*(*uint8)(unsafe.Add(unsafe.Pointer(v3), 12)))&2 != 0 {
			v9 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1193380)))).ChildByID(0x2721)))
		} else {
			v9 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1193380)))).ChildByID(0x2720)))
		}
		v10 = v9
		*(*uint32)(unsafe.Add(unsafe.Pointer(v9), unsafe.Sizeof(uint32(0))*9)) |= 4
		if *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v3))), unsafe.Sizeof(uint32(0))*5))) != 0 {
			v10 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1193380)))).ChildByID(0x271E)))
			*(*uint32)(unsafe.Add(unsafe.Pointer(v10), unsafe.Sizeof(uint32(0))*9)) |= 4
		}
		if *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v3))), unsafe.Sizeof(uint32(0))*10))) != 0 {
			v10 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1193380)))).ChildByID(0x2722)))
			*(*uint32)(unsafe.Add(unsafe.Pointer(v10), unsafe.Sizeof(uint32(0))*9)) |= 4
		}
		v11 = int32(*memmap.PtrUint32(6112660, uint32(v2*4)+0x12359C))
		if v11 != 0 {
			v12 = v11 - 1
			if v12 != 0 {
				if v12 == 1 {
					v10 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1193380)))).ChildByID(0x272A)))
					sub_489DC0()
				}
			} else {
				v10 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1193380)))).ChildByID(0x2729)))
				(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1193384))))).Hide()
			}
		} else {
			v10 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1193380)))).ChildByID(0x2728)))
			(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1193384))))).Hide()
		}
		*(*uint32)(unsafe.Add(unsafe.Pointer(v10), unsafe.Sizeof(uint32(0))*9)) |= 4
		result = *(**uint32)(unsafe.Pointer(&dword_5d4594_1193380))
	}
	return result
}
func sub_489DC0() {
	var (
		v0 *uint32
		v1 *uint32
	)
	(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1193384))))).Show()
	if (*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1193384)))).ChildByID(0x272C).draw_data.field_0&4 != 0 {
		v0 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1193384)))).ChildByID(0x272F)))
		nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v0)))))), 1)
	} else {
		v1 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1193384)))).ChildByID(0x272F)))
		nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v1)))))), 0)
	}
	if (*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1193384)))).ChildByID(0x271F).draw_data.field_0&4 != 0 {
		sub_46AD20(*(**uint32)(unsafe.Pointer(&dword_5d4594_1193384)), 0x2720, 0x2721, 1)
	} else {
		sub_46AD20(*(**uint32)(unsafe.Pointer(&dword_5d4594_1193384)), 0x2720, 0x2721, 0)
	}
}
func nox_xxx_windowMplayFilterProc_489E70(a1 int32, a2 int32, a3 *int32, a4 int32) int32 {
	var (
		v3     int32
		v4     int32
		result int32
		v6     *uint32
		v7     int32
		v8     int32
	)
	v3 = nox_xxx_check_flag_aaa_43AF70()
	if a2 == 23 {
		return 1
	}
	if a2 != 0x4007 {
		return 0
	}
	v4 = (*nox_window)(unsafe.Pointer(a3)).ID()
	clientPlaySoundSpecial(766, 100)
	switch v4 {
	case 0x271F:
		sub_46AD20(*(**uint32)(unsafe.Pointer(&dword_5d4594_1193380)), 0x2720, 0x2721, int32((uint32(^*(*int32)(unsafe.Add(unsafe.Pointer(a3), unsafe.Sizeof(int32(0))*9)))>>2)&1))
		return 0
	case 0x2728:
		v7 = int32(dword_5d4594_1193384)
		*memmap.PtrUint32(6112660, uint32(v3*4)+0x12359C) = 0
		(*nox_window)(unsafe.Pointer(uintptr(v7))).Hide()
		result = 0
	case 0x2729:
		v8 = int32(dword_5d4594_1193384)
		*memmap.PtrUint32(6112660, uint32(v3*4)+0x12359C) = 1
		(*nox_window)(unsafe.Pointer(uintptr(v8))).Hide()
		result = 0
	case 0x272A:
		*memmap.PtrUint32(6112660, uint32(v3*4)+0x12359C) = 2
		sub_489DC0()
		result = 0
	case 0x272C:
		v6 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1193380)))).ChildByID(0x272F)))
		nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v6)))))), int32((uint32(^*(*int32)(unsafe.Add(unsafe.Pointer(a3), unsafe.Sizeof(int32(0))*9)))>>2)&1))
		result = 0
	default:
		return 0
	}
	return result
}
func sub_489FB0() int32 {
	var result int32
	result = int32(dword_5d4594_1193380)
	if dword_5d4594_1193380 != 0 {
		sub_489870()
		nox_xxx_wndClearCaptureMain((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1193380))))))
		result = (*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1193380)))).Destroy()
		dword_5d4594_1193380 = 0
	}
	return result
}
func sub_489FF0(a1 int32, a2 int32, a3 unsafe.Pointer) int32 {
	var result int32
	*memmap.PtrUint32(6112660, uint32(a1*4)+0x12359C) = uint32(a2)
	result = a1 * 11
	libc.MemCpy(memmap.PtrOff(6112660, uint32(a1*44)+0x1235AC), a3, 44)
	return result
}
func sub_48A020(a1 int32, a2 *uint32) int32 {
	*a2 = uint32(uintptr(memmap.PtrOff(6112660, uint32(a1*44)+0x1235AC)))
	return int32(*memmap.PtrUint32(6112660, uint32(a1*4)+0x12359C))
}
func nox_video_checkIsWinNT_48A1D0(a1 int32) int32 {
	var v1 int32
	if nox_video_windowsPlatformVersion != 5 {
		v1 = int32(*(*uint32)(unsafe.Pointer(&nox_video_renderTargetFlags)))
		if a1 != 0 {
			*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v1))), 0)) = uint8(int8(nox_video_renderTargetFlags | 32))
		} else {
			*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v1))), 0)) = uint8(int8(nox_video_renderTargetFlags & 223))
		}
		*(*uint32)(unsafe.Pointer(&nox_video_renderTargetFlags)) = uint32(v1)
		sub_444D00()
	}
	return int32((*(*uint32)(unsafe.Pointer(&nox_video_renderTargetFlags)) >> 5) & 1)
}
func nox_xxx_setSomeFunc_48A210(a1 int32) int32 {
	var result int32
	result = a1
	*memmap.PtrUint32(6112660, 1193504) = uint32(a1)
	return result
}
func sub_48A5D0() int32 {
	return 0
}
func sub_48B3E0(a1 int32) int32 {
	var result int32
	result = int32(dword_5d4594_1193672)
	dword_5d4594_1193672 = uint32(a1)
	return result
}
func sub_48BD90(a1 int32) {
	nox_mutex_lock((*nox_mutex_t)(memmap.PtrOff(0x973F18, 168)))
	if a1 == 0 {
		compatSetRect((LPRECT)(memmap.PtrOff(6112660, 1193532)), 0, 0, 0, 0)
	}
	dword_5d4594_1193664 = 1
	dword_5d4594_1193668 = 1
	dword_5d4594_1193624 = nil
	nox_mutex_unlock((*nox_mutex_t)(memmap.PtrOff(0x973F18, 168)))
}
func sub_48C420() int8 {
	var (
		v0 *uint8
		v1 *uint8
		v2 uint8
		v3 int32
		v4 int32
		v6 int32
		v7 int32
	)
	v0 = *(**uint8)(unsafe.Pointer(&dword_5d4594_1193516))
	v1 = *(**uint8)(unsafe.Pointer(&dword_5d4594_1193584))
	v2 = *memmap.PtrUint8(6112660, 1193520)
	v3 = int32(*memmap.PtrUint32(6112660, 1193520) >> 2)
	if *memmap.PtrUint32(6112660, 1193520)>>2 != 0 {
		for {
			v4 = int32(*(*uint32)(unsafe.Pointer(v0)))
			v0 = (*uint8)(unsafe.Add(unsafe.Pointer(v0), 4))
			*(*uint32)(unsafe.Pointer(v1)) = uint32(v4)
			v1 = (*uint8)(unsafe.Add(unsafe.Pointer(v1), 4))
			if func() int32 {
				p := &v3
				x := *p
				*p--
				return x
			}() <= 1 {
				break
			}
		}
	}
	v6 = int32(v2) & 3
	if v6 != 0 {
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v4))), 0)) = *func() *uint8 {
			p := &v0
			x := *p
			*p = (*uint8)(unsafe.Add(unsafe.Pointer(*p), 1))
			return x
		}()
		*func() *uint8 {
			p := &v1
			x := *p
			*p = (*uint8)(unsafe.Add(unsafe.Pointer(*p), 1))
			return x
		}() = uint8(int8(v4))
		v7 = v6 - 1
		if v7 != 0 {
			*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v4))), 0)) = *func() *uint8 {
				p := &v0
				x := *p
				*p = (*uint8)(unsafe.Add(unsafe.Pointer(*p), 1))
				return x
			}()
			*func() *uint8 {
				p := &v1
				x := *p
				*p = (*uint8)(unsafe.Add(unsafe.Pointer(*p), 1))
				return x
			}() = uint8(int8(v4))
			if v7 != 1 {
				*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v4))), 0)) = *func() *uint8 {
					p := &v0
					x := *p
					*p = (*uint8)(unsafe.Add(unsafe.Pointer(*p), 1))
					return x
				}()
				*func() *uint8 {
					p := &v1
					x := *p
					*p = (*uint8)(unsafe.Add(unsafe.Pointer(*p), 1))
					return x
				}() = uint8(int8(v4))
			}
		}
	}
	dword_5d4594_1193516 = uint32(uintptr(unsafe.Pointer(v0)))
	dword_5d4594_1193584 = uint32(uintptr(unsafe.Pointer(v1)))
	return int8(v4)
}
func sub_48C480() int16 {
	var (
		v0 *int32
		v1 *uint32
		v2 int8
		v3 int32
		v4 int32
	)
	v0 = *(**int32)(unsafe.Pointer(&dword_5d4594_1193516))
	v1 = *(**uint32)(unsafe.Pointer(&dword_5d4594_1193584))
	v2 = int8(int32(*memmap.PtrUint8(6112660, 1193520)) & 1)
	v3 = int32(*memmap.PtrUint32(6112660, 1193520) >> 1)
	if *memmap.PtrUint32(6112660, 1193520)>>1 != 0 {
		for {
			v4 = *v0
			v0 = (*int32)(unsafe.Add(unsafe.Pointer(v0), unsafe.Sizeof(int32(0))*1))
			*v1 = uint32(v4)
			v1 = (*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*1))
			if func() int32 {
				p := &v3
				x := *p
				*p--
				return x
			}() <= 1 {
				break
			}
		}
	}
	if int32(v2) != 0 {
		*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v4))), unsafe.Sizeof(uint16(0))*0)) = *(*uint16)(unsafe.Pointer(v0))
		v0 = (*int32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(v0))), 2))))
		*(*uint16)(unsafe.Pointer(v1)) = uint16(int16(v4))
		v1 = (*uint32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(v1))), 2))))
	}
	dword_5d4594_1193516 = uint32(uintptr(unsafe.Pointer(v0)))
	dword_5d4594_1193584 = uint32(uintptr(unsafe.Pointer(v1)))
	return int16(v4)
}
