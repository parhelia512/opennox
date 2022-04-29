package opennox

import (
	"github.com/gotranspile/cxgo/runtime/cnet"
	"github.com/gotranspile/cxgo/runtime/libc"
	"github.com/gotranspile/cxgo/runtime/stdio"
	"github.com/noxworld-dev/opennox/v1/common/alloc"
	noxflags "github.com/noxworld-dev/opennox/v1/common/flags"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"math"
	"unsafe"
)

var nox_win_width int32 = 0
var nox_win_height int32 = 0
var ptr_5D4594_754088 *obj_5D4594_754088_t = nil
var ptr_5D4594_754088_cnt int32 = 0
var ptr_5D4594_754092 *obj_5D4594_754088_t = nil
var ptr_5D4594_754092_cnt int32 = 0
var obj_5D4594_754104_switch int32 = 0
var nox_draw_curDrawData_3799572 *nox_render_data_t = nil
var nox_screenParticles_head *nox_screenParticle = nil
var dword_5d4594_806052 *nox_screenParticle = nil
var nox_draw_colorTablesRev_3804668 unsafe.Pointer = nil
var dword_5d4594_814624 unsafe.Pointer = nil

func nox_xxx_wallMath_427F30(a1 *int2, a2 *int32) int32 {
	var (
		v2  int32
		v3  int32
		v4  int32
		v5  int32
		v6  int32
		v7  int32
		v9  int32
		v10 int32
		v12 int32
		v13 int32
	)
	v2 = a1.field_0
	if float64(a1.field_0) < 57.5 {
		return 0
	}
	v3 = a1.field_4
	if float64(v3) < 57.5 {
		return 0
	}
	if v2 > 5888 {
		return 0
	}
	if v3 > 5888 {
		return 0
	}
	v4 = v2 / 23
	v5 = v3 / 23
	v6 = *(*int32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(int32(0))*1)) / 23
	v12 = *(*int32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(int32(0))*2)) / 23
	v13 = *(*int32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(int32(0))*3)) / 23
	v7 = *(*int32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(int32(0))*4)) / 23
	v9 = *(*int32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(int32(0))*5)) / 23
	if v3/23 < v6 || v5 > *(*int32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(int32(0))*7))/23 || v4 < v12 || v4 > v7 {
		return 0
	}
	if v5 > v13 {
		if v4 < v5+v12-v13 {
			return 0
		}
		v10 = *a2 / 23
	} else {
		v10 = *a2 / 23
		if v4 < v10+v6-v5 {
			return 0
		}
	}
	if v5 > v9 {
		if v4 > v7+v9-v5 {
			return 0
		}
	} else if v4 > v5+v10-v6 {
		return 0
	}
	return 1
}
func sub_428170(a1p unsafe.Pointer, a2 *int4) int32 {
	var (
		a1 *uint32 = (*uint32)(a1p)
		v2 int32
		v3 int32
	)
	v2 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*1)))
	if uint32(v2) >= *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*7)) {
		a2.field_C = v2
		a2.field_4 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*7)))
	} else {
		a2.field_4 = v2
		a2.field_C = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*7)))
	}
	v3 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*2)))
	if uint32(v3) >= *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*4)) {
		a2.field_8 = v3
		a2.field_0 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*4)))
	} else {
		a2.field_0 = v3
		a2.field_8 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*4)))
	}
	if a2.field_0 < 0 {
		a2.field_0 = 0
	}
	if a2.field_4 < 0 {
		a2.field_4 = 0
	}
	if a2.field_8 >= 5888 {
		a2.field_8 = 5887
	}
	if a2.field_C >= 5888 {
		a2.field_C = 5887
	}
	return 0
}
func sub_4281F0(a1 *int2, a2 *int4) int32 {
	var (
		v2     int32
		result bool
	)
	result = false
	if a1.field_0 >= a2.field_0 && a1.field_0 <= a2.field_8 {
		v2 = a1.field_4
		if v2 >= a2.field_4 && v2 <= a2.field_C {
			result = true
		}
	}
	return bool2int(result)
}
func sub_428220(a1 *float2, a2 *float4) int32 {
	return bool2int(float64(a1.field_0) >= float64(a2.field_0) && float64(a1.field_0) <= float64(a2.field_8) && float64(a1.field_4) >= float64(a2.field_4) && float64(a1.field_4) <= float64(a2.field_C))
}
func nox_shape_box_calc(s *nox_shape) {
	var (
		mul float32 = 0.35354999
		px  float32 = s.box_w * mul
		py  float32 = s.box_h * mul
		v   float64 = 0.0
	)
	v = float64(-px + py)
	s.box_left_top = float32(v)
	s.box_left_top_2 = float32(v)
	v = float64(-px - py)
	s.box_left_bottom = float32(v)
	s.box_left_bottom_2 = float32(v)
	v = float64(+px + py)
	s.box_right_top = float32(v)
	s.box_right_top_2 = float32(v)
	v = float64(+px - py)
	s.box_right_bottom = float32(v)
	s.box_right_bottom_2 = float32(v)
}
func sub_4282D0(a1 *byte, a2 int32) *byte {
	var result *byte
	result = libc.StrNCpy((*byte)(memmap.PtrOff(6112660, 741316)), a1, 16)
	*memmap.PtrUint32(6112660, 741304) = uint32(a2)
	return result
}
func sub_4282F0(a1 int32, a2 int32, a3 uint32) unsafe.Pointer {
	var (
		v3     int32
		v4     uint32
		i      uint32
		v6     uint32
		result unsafe.Pointer
		v8     uint32
		v9     int32
		v10    int32
		v11    int32
		j      int32
	)
	v3 = a1
	v4 = 0
	if *(*uint32)(unsafe.Pointer(uintptr(a1 + 608))) != 0 {
		for i = 0; i < uint32(*(*int32)(unsafe.Pointer(&dword_5d4594_741332))); i++ {
			*(*unsafe.Pointer)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 608))) + i*4))) = nil
			*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 608))) + i*4))) = 0
		}
		*(*unsafe.Pointer)(unsafe.Pointer(uintptr(a1 + 608))) = nil
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 608))) = 0
	}
	if *(*uint32)(unsafe.Pointer(uintptr(a1 + 612))) != 0 {
		*(*unsafe.Pointer)(unsafe.Pointer(uintptr(a1 + 612))) = nil
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 612))) = 0
	}
	if *(*uint32)(unsafe.Pointer(uintptr(a1 + 616))) != 0 {
		*(*unsafe.Pointer)(unsafe.Pointer(uintptr(a1 + 616))) = nil
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 616))) = 0
	}
	if *(*uint32)(unsafe.Pointer(uintptr(a1 + 620))) != 0 {
		*(*unsafe.Pointer)(unsafe.Pointer(uintptr(a1 + 620))) = nil
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 620))) = 0
	}
	if *(*uint32)(unsafe.Pointer(uintptr(a1 + 624))) != 0 {
		*(*unsafe.Pointer)(unsafe.Pointer(uintptr(a1 + 624))) = nil
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 624))) = 0
	}
	if *(*uint32)(unsafe.Pointer(uintptr(a1 + 628))) != 0 {
		*(*unsafe.Pointer)(unsafe.Pointer(uintptr(a1 + 628))) = nil
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 628))) = 0
	}
	if *(*uint32)(unsafe.Pointer(uintptr(a1 + 632))) != 0 {
		*(*unsafe.Pointer)(unsafe.Pointer(uintptr(a1 + 632))) = nil
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 632))) = 0
	}
	v6 = a3 * 4
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 608))) = uint32(uintptr(alloc.Calloc(int(a3), 4)))
	if a3 != 0 {
		for {
			*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 608))) + func() uint32 {
				p := &v4
				x := *p
				*p++
				return x
			}()*4))) = uint32(uintptr(alloc.Calloc(1, 10)))
			if v4 >= a3 {
				break
			}
		}
	}
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 612))) = uint32(uintptr(alloc.Calloc(1, int(v6))))
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 616))) = uint32(uintptr(alloc.Calloc(1, int(v6))))
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 620))) = uint32(uintptr(alloc.Calloc(1, int(a3))))
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 624))) = uint32(uintptr(alloc.Calloc(1, int(a3))))
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 628))) = uint32(uintptr(alloc.Calloc(1, int(v6))))
	result = alloc.Calloc(1, int(a3))
	v8 = 0
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 632))) = uint32(uintptr(result))
	if a3 != 0 {
		v9 = a2 + 16
		for j = a2 + 16; ; v9 = j {
			libc.StrCpy(*(**byte)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v3 + 608))) + v8*4))), (*byte)(unsafe.Pointer(uintptr(v9-16))))
			*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v3 + 612))) + v8*4))) = *(*uint32)(unsafe.Pointer(uintptr(v9 - 4)))
			*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v3 + 616))) + v8*4))) = *(*uint32)(unsafe.Pointer(uintptr(v9)))
			*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v3 + 620))) + v8))) = *(*uint8)(unsafe.Pointer(uintptr(v9 + 4)))
			*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v3 + 624))) + v8))) = *(*uint8)(unsafe.Pointer(uintptr(v9 + 5)))
			*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v3 + 632))) + v8))) = *(*uint8)(unsafe.Pointer(uintptr(v9 + 12)))
			v10 = int32(*(*uint32)(unsafe.Pointer(uintptr(v9 + 8))))
			v11 = int32(libc.GetTime(nil) - libc.Time(v10))
			v8++
			*(*uint32)(unsafe.Pointer(uintptr(j + 8))) = uint32(v11)
			*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v3 + 628))) + v8*4 - 4))) = uint32(v11)
			result = unsafe.Pointer(uintptr(a3))
			j += 32
			if v8 >= a3 {
				break
			}
		}
		*(*uint16)(unsafe.Pointer(uintptr(v3 + 6))) = uint16(a3)
		dword_5d4594_741332 = a3
	} else {
		*(*uint16)(unsafe.Pointer(uintptr(a1 + 6))) = 0
		dword_5d4594_741332 = 0
	}
	return result
}
func sub_428540(a1 int32, a2 *byte, a3 int32) uint32 {
	var (
		result uint32
		v4     *byte
		v5     int8
	)
	if *(*uint32)(unsafe.Pointer(uintptr(a1 + 636))) != 0 {
		*(*unsafe.Pointer)(unsafe.Pointer(uintptr(a1 + 636))) = nil
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 636))) = 0
	}
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 636))) = uint32(uintptr(alloc.Calloc(int(a3), 2)))
	result = 0
	if a3*2 != 0 {
		v4 = a2
		for {
			v5 = int8(*v4)
			v4 = (*byte)(unsafe.Add(unsafe.Pointer(v4), 2))
			*(*uint8)(unsafe.Pointer(uintptr(result + *(*uint32)(unsafe.Pointer(uintptr(a1 + 636)))))) = uint8(v5)
			*(*uint8)(unsafe.Pointer(uintptr(result + *(*uint32)(unsafe.Pointer(uintptr(a1 + 636))) + 1))) = uint8(*((*byte)(unsafe.Add(unsafe.Pointer(v4), -1))))
			result += 2
			if result >= uint32(a3*2) {
				break
			}
		}
		*memmap.PtrUint32(6112660, 741308) = uint32(a3)
	} else {
		*memmap.PtrUint32(6112660, 741308) = uint32(a3)
	}
	return result
}
func sub_4285C0(a1 *int16) *byte {
	var (
		result *byte
		v2     *int16
		v3     unsafe.Pointer
		v4     uint32
		v5     bool
		v6     unsafe.Pointer
		v7     unsafe.Pointer
		v8     unsafe.Pointer
		v9     unsafe.Pointer
		v10    unsafe.Pointer
		v11    unsafe.Pointer
		v12    *byte
		v13    int32
		v14    uint32
		v15    uint32
		v16    uint32
		v17    uint32
		v18    uint32
		v19    uint32
		v20    int32
	)
	result = (*byte)(unsafe.Pointer(uintptr(nox_xxx_player_4E3CE0())))
	v2 = a1
	*a1 = int16(uintptr(unsafe.Pointer(result)))
	if int32(int16(uintptr(unsafe.Pointer(result)))) <= 0 {
		dword_5d4594_741332 = uint32(*a1)
	} else {
		v3 = alloc.Calloc(int(int16(uintptr(unsafe.Pointer(result)))), 4)
		v4 = 0
		v5 = int32(*a1) == 0
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a1))), unsafe.Sizeof(uint32(0))*134))) = uint32(uintptr(v3))
		if !v5 {
			for {
				*(*uint32)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a1))), unsafe.Sizeof(uint32(0))*134))) + func() uint32 {
					p := &v4
					x := *p
					*p++
					return x
				}()*4))) = uint32(uintptr(alloc.Calloc(1, 10)))
				if v4 >= uint32(*a1) {
					break
				}
			}
		}
		v6 = alloc.Calloc(int(*a1), 4)
		v19 = uint32(*a1)
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a1))), unsafe.Sizeof(uint32(0))*135))) = uint32(uintptr(v6))
		v7 = alloc.Calloc(1, int(v19))
		v18 = uint32(int32(*a1) * 4)
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a1))), unsafe.Sizeof(uint32(0))*144))) = uint32(uintptr(v7))
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a1))), unsafe.Sizeof(uint32(0))*136))) = uint32(uintptr(alloc.Calloc(1, int(v18))))
		v8 = alloc.Calloc(int(*a1), 4)
		v17 = uint32(int32(*a1) * 4)
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a1))), unsafe.Sizeof(uint32(0))*137))) = uint32(uintptr(v8))
		v9 = alloc.Calloc(1, int(v17))
		v16 = uint32(int32(*a1) * 4)
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a1))), unsafe.Sizeof(uint32(0))*138))) = uint32(uintptr(v9))
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a1))), unsafe.Sizeof(uint32(0))*139))) = uint32(uintptr(alloc.Calloc(1, int(v16))))
		v10 = alloc.Calloc(int(*a1), 4)
		v15 = uint32(int32(*a1) * 4)
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a1))), unsafe.Sizeof(uint32(0))*140))) = uint32(uintptr(v10))
		v11 = alloc.Calloc(1, int(v15))
		v14 = uint32(int32(*a1) * 4)
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a1))), unsafe.Sizeof(uint32(0))*141))) = uint32(uintptr(v11))
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a1))), unsafe.Sizeof(uint32(0))*142))) = uint32(uintptr(alloc.Calloc(1, int(v14))))
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a1))), unsafe.Sizeof(uint32(0))*143))) = uint32(uintptr(alloc.Calloc(int(*a1), 4)))
		v20 = 0
		result = (*byte)(unsafe.Pointer(nox_common_playerInfoGetFirst_416EA0()))
		v12 = result
		if result != nil {
			for {
				if *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v12))), unsafe.Sizeof(uint32(0))*1198))) != 0 {
					libc.StrCpy(*(**byte)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), unsafe.Sizeof(uint32(0))*134))) + uint32(v20*4)))), (*byte)(unsafe.Add(unsafe.Pointer(v12), 2096)))
					v13 = nox_xxx_net_getIP_554200(uint32(int32(uint8(*(*byte)(unsafe.Add(unsafe.Pointer(v12), 2064)))) + 1))
					*(*uint32)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), unsafe.Sizeof(uint32(0))*135))) + uint32(v20*4)))) = cnet.Htonl(uint32(v13))
					*(*uint8)(unsafe.Pointer(uintptr(uint32(v20) + *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), unsafe.Sizeof(uint32(0))*144)))))) = uint8(*(*byte)(unsafe.Add(unsafe.Pointer(v12), 2251)))
					*(*uint32)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), unsafe.Sizeof(uint32(0))*136))) + uint32(v20*4)))) = *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v12))), unsafe.Sizeof(uint32(0))*1172)))
					*(*uint32)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), unsafe.Sizeof(uint32(0))*137))) + uint32(v20*4)))) = *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v12))), unsafe.Sizeof(uint32(0))*1172)))
					*(*uint32)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), unsafe.Sizeof(uint32(0))*138))) + uint32(v20*4)))) = *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v12))), unsafe.Sizeof(uint32(0))*1166)))
					*(*uint32)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), unsafe.Sizeof(uint32(0))*139))) + uint32(v20*4)))) = *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v12))), unsafe.Sizeof(uint32(0))*1165)))
					*(*uint32)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), unsafe.Sizeof(uint32(0))*140))) + uint32(v20*4)))) = *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v12))), unsafe.Sizeof(uint32(0))*1167)))
					*(*uint32)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), unsafe.Sizeof(uint32(0))*141))) + uint32(v20*4)))) = *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v12))), unsafe.Sizeof(uint32(0))*1168)))
					*(*uint32)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), unsafe.Sizeof(uint32(0))*142))) + uint32(v20*4)))) = *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v12))), unsafe.Sizeof(uint32(0))*1171)))
					*(*uint32)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), unsafe.Sizeof(uint32(0))*143))) + uint32(func() int32 {
						p := &v20
						x := *p
						*p++
						return x
					}()*4)))) = sub_4D6540(int32(uint8(*(*byte)(unsafe.Add(unsafe.Pointer(v12), 2064)))))
				}
				v12 = (*byte)(unsafe.Pointer(nox_common_playerInfoGetNext_416EE0((*nox_playerInfo)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v12)))))))))
				if v12 == nil {
					break
				}
			}
			result = (*byte)(unsafe.Pointer(uintptr(*v2)))
			dword_5d4594_741332 = uint32(*v2)
		} else {
			dword_5d4594_741332 = uint32(*v2)
		}
	}
	return result
}
func sub_428810(a1 int32, a2 int32) int32 {
	var (
		v2 *uint16
		v3 *uint16
		v4 int32
		v6 [72]byte
	)
	v2 = sub_42ADA0(a1, a2, *memmap.PtrInt16(6112660, 741308), memmap.PtrUint32(6112660, 741312))
	v3 = sub_42A8B0((*uint8)(unsafe.Pointer(v2)), memmap.PtrInt32(6112660, 741312))
	alloc.Free(unsafe.Pointer(v2))
	v4 = 1
	if sub_420360(&v6[0], (*uint16)(unsafe.Pointer(&a2))) != 0 {
		v4 = sub_40DEA0(int32(uintptr(unsafe.Pointer(&v6[0]))), int32(int16(a2)), int32(uintptr(unsafe.Pointer(v3))), *memmap.PtrInt32(6112660, 741312))
	}
	alloc.Free(unsafe.Pointer(v3))
	return v4
}
func sub_428890(a1 *int16) int32 {
	var (
		v1 *uint16
		v2 *uint16
		v3 int32
		v5 [72]byte
	)
	v1 = sub_42B810(a1, memmap.PtrUint32(6112660, 741300))
	v2 = sub_42A8B0((*uint8)(unsafe.Pointer(v1)), memmap.PtrInt32(6112660, 741300))
	alloc.Free(unsafe.Pointer(v1))
	v3 = 1
	if sub_420360(&v5[0], (*uint16)(unsafe.Pointer(&a1))) != 0 {
		v3 = sub_40DEA0(int32(uintptr(unsafe.Pointer(&v5[0]))), int32(int16(uintptr(unsafe.Pointer(a1)))), int32(uintptr(unsafe.Pointer(v2))), *memmap.PtrInt32(6112660, 741300))
	}
	alloc.Free(unsafe.Pointer(v2))
	return v3
}
func sub_428910(a1 *unsafe.Pointer) unsafe.Pointer {
	var (
		i      uint32
		result unsafe.Pointer
	)
	if *(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(unsafe.Pointer(nil))*152)) != nil {
		for i = 0; i < uint32(*(*int32)(unsafe.Pointer(&dword_5d4594_741332))); i++ {
			*((*unsafe.Pointer)(unsafe.Add(unsafe.Pointer((*unsafe.Pointer)(*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(unsafe.Pointer(nil))*152)))), unsafe.Sizeof(unsafe.Pointer(nil))*uintptr(i)))) = nil
			*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(unsafe.Pointer(nil))*152)))), unsafe.Sizeof(uint32(0))*uintptr(i)))) = 0
		}
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(unsafe.Pointer(nil))*152)) = nil
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(unsafe.Pointer(nil))*152)) = nil
	}
	if *(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(unsafe.Pointer(nil))*153)) != nil {
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(unsafe.Pointer(nil))*153)) = nil
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(unsafe.Pointer(nil))*153)) = nil
	}
	if *(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(unsafe.Pointer(nil))*154)) != nil {
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(unsafe.Pointer(nil))*154)) = nil
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(unsafe.Pointer(nil))*154)) = nil
	}
	if *(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(unsafe.Pointer(nil))*155)) != nil {
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(unsafe.Pointer(nil))*155)) = nil
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(unsafe.Pointer(nil))*155)) = nil
	}
	result = *(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(unsafe.Pointer(nil))*159))
	if result != nil {
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(unsafe.Pointer(nil))*159)) = nil
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(unsafe.Pointer(nil))*159)) = nil
	}
	return result
}
func sub_4289D0(a1 *unsafe.Pointer) unsafe.Pointer {
	var (
		i      uint32
		result unsafe.Pointer
	)
	if *(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(unsafe.Pointer(nil))*134)) != nil {
		for i = 0; i < uint32(*(*int32)(unsafe.Pointer(&dword_5d4594_741332))); i++ {
			*((*unsafe.Pointer)(unsafe.Add(unsafe.Pointer((*unsafe.Pointer)(*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(unsafe.Pointer(nil))*134)))), unsafe.Sizeof(unsafe.Pointer(nil))*uintptr(i)))) = nil
			*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(unsafe.Pointer(nil))*134)))), unsafe.Sizeof(uint32(0))*uintptr(i)))) = 0
		}
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(unsafe.Pointer(nil))*134)) = nil
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(unsafe.Pointer(nil))*134)) = nil
	}
	if *(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(unsafe.Pointer(nil))*135)) != nil {
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(unsafe.Pointer(nil))*135)) = nil
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(unsafe.Pointer(nil))*135)) = nil
	}
	if *(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(unsafe.Pointer(nil))*144)) != nil {
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(unsafe.Pointer(nil))*144)) = nil
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(unsafe.Pointer(nil))*144)) = nil
	}
	if *(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(unsafe.Pointer(nil))*136)) != nil {
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(unsafe.Pointer(nil))*136)) = nil
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(unsafe.Pointer(nil))*136)) = nil
	}
	if *(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(unsafe.Pointer(nil))*137)) != nil {
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(unsafe.Pointer(nil))*137)) = nil
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(unsafe.Pointer(nil))*137)) = nil
	}
	if *(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(unsafe.Pointer(nil))*138)) != nil {
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(unsafe.Pointer(nil))*138)) = nil
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(unsafe.Pointer(nil))*138)) = nil
	}
	if *(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(unsafe.Pointer(nil))*139)) != nil {
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(unsafe.Pointer(nil))*139)) = nil
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(unsafe.Pointer(nil))*139)) = nil
	}
	if *(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(unsafe.Pointer(nil))*140)) != nil {
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(unsafe.Pointer(nil))*140)) = nil
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(unsafe.Pointer(nil))*140)) = nil
	}
	if *(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(unsafe.Pointer(nil))*141)) != nil {
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(unsafe.Pointer(nil))*141)) = nil
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(unsafe.Pointer(nil))*141)) = nil
	}
	if *(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(unsafe.Pointer(nil))*142)) != nil {
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(unsafe.Pointer(nil))*142)) = nil
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(unsafe.Pointer(nil))*142)) = nil
	}
	result = *(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(unsafe.Pointer(nil))*143))
	if result != nil {
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(unsafe.Pointer(nil))*143)) = nil
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(unsafe.Pointer(nil))*143)) = nil
	}
	return result
}
func nox_server_mapRWObjectTOC_428B30() int32 {
	var (
		v1  int32
		v2  uint16
		v3  int32
		v4  int32
		v5  int32
		i   int32
		v7  uint16
		v8  int32
		v9  int32
		v10 int32
		v11 int32
		v12 [256]byte
	)
	v11 = 1
	cryptFileReadWrite((*uint8)(unsafe.Pointer(&v11))[:2])
	if int32(int16(v11)) > 1 {
		return 0
	}
	sub_42BFB0()
	if nox_xxx_cryptGetXxx() != 0 {
		cryptFileReadWrite((*uint8)(unsafe.Pointer(&v9))[:2])
		for i = 0; int32(uint16(int16(i))) < int32(uint16(int16(v9))); i++ {
			cryptFileReadWrite((*uint8)(unsafe.Pointer(&v10))[:2])
			cryptFileReadWrite((*uint8)(unsafe.Pointer(&v8))[:1])
			cryptFileReadWrite((*uint8)(unsafe.Pointer(&v12[0]))[:uint32(uint8(int8(v8)))])
			v12[uint8(int8(v8))] = 0
			if !noxflags.HasGame(noxflags.GameClient) || noxflags.HasGame(noxflags.GameHost) {
				v7 = uint16(int16(nox_xxx_getNameId_4E3AA0(&v12[0])))
			} else {
				v7 = uint16(int16(nox_xxx_getTTByNameSpriteMB_44CFC0(&v12[0])))
			}
			sub_42C310(v7, int16(v10))
		}
		return 1
	}
	sub_42BFE0()
	*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v1))), unsafe.Sizeof(uint16(0))*0)) = uint16(sub_42C300())
	v9 = v1
	cryptFileReadWrite((*uint8)(unsafe.Pointer(&v9))[:2])
	v2 = 0
	if nox_xxx_unitDefGetCount_4E3AC0() == 0 {
		return 1
	}
	v3 = 0
	for {
		*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v4))), unsafe.Sizeof(uint16(0))*0)) = uint16(sub_42C2E0(v3))
		v10 = v4
		if int32(uint16(int16(v4))) != 0 {
			cryptFileReadWrite((*uint8)(unsafe.Pointer(&v10))[:2])
			v5 = int32(uintptr(unsafe.Pointer(nox_xxx_objectTypeByInd_4E3B70(v3))))
			*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v8))), 0)) = uint8(int8(libc.StrLen(*(**byte)(unsafe.Pointer(uintptr(v5 + 4))))))
			cryptFileReadWrite((*uint8)(unsafe.Pointer(&v8))[:1])
			cryptFileReadWrite((*(**uint8)(unsafe.Pointer(uintptr(v5 + 4))))[:uint32(uint8(int8(v8)))])
		}
		v3 = int32(func() uint16 {
			p := &v2
			*p++
			return *p
		}())
		if uint32(v2) >= uint32(nox_xxx_unitDefGetCount_4E3AC0()) {
			break
		}
	}
	return 1
}
func nox_server_mapRWPolygons_428CD0(a1 int32) int32 {
	var (
		i   *byte
		v3  *byte
		j   uint32
		k   *byte
		v6  *byte
		v7  *byte
		v8  *byte
		v9  uint32
		v10 *uint8
		v11 *uint16
		v12 *uint8
		v13 *byte
		v14 int32
		v15 uint32
		v16 bool
		v17 *byte
		v18 float64
		v19 float64
		v20 float64
		v21 float64
		v22 *byte
		v23 *byte
		v24 *uint8
		v25 uint32
		v26 int32
		v27 int32
		v28 uint32
		v29 [2]float32
		v30 int32
		v31 int32
		v32 [12]byte
	)
	v30 = 0
	if a1 != 0 {
		return 1
	}
	v26 = 4
	cryptFileReadWrite((*uint8)(unsafe.Pointer(&v26))[:2])
	if int32(int16(v26)) > 4 {
		return 0
	}
	if nox_xxx_cryptGetXxx() != 0 {
		cryptFileReadWrite((*uint8)(unsafe.Pointer(&v25))[:4])
		v9 = 1
		if v25 >= 1 {
			for {
				cryptFileReadWrite((*uint8)(unsafe.Pointer(&a1))[:4])
				cryptFileReadWrite((*uint8)(unsafe.Pointer(&v29[0]))[:8])
				if nox_xxx_polygonSetAngle_420D40(*(*int32)(unsafe.Add(unsafe.Pointer((*int32)(unsafe.Pointer(&v29[0]))), unsafe.Sizeof(int32(0))*0)), *(*int32)(unsafe.Add(unsafe.Pointer((*int32)(unsafe.Pointer(&v29[1]))), unsafe.Sizeof(int32(0))*0)), v9, a1) == nil {
					return 0
				}
				if int32(int16(v26)) < 3 {
					cryptFileReadWrite((*uint8)(unsafe.Pointer(&v31))[:4])
				}
				if func() uint32 {
					p := &v9
					*p++
					return *p
				}() > v25 {
					break
				}
			}
		}
		cryptFileReadWrite((*uint8)(unsafe.Pointer(&v25))[:4])
		v28 = 1
		if v25 >= 1 {
			for {
				v10 = sub_421230()
				if v10 == nil {
					return 0
				}
				cryptFileReadWrite((*uint8)(unsafe.Pointer(&v27))[:1])
				cryptFileReadWrite((*uint8)(unsafe.Add(unsafe.Pointer(v10), 4))[:uint32(uint8(int8(v27)))])
				*(*uint8)(unsafe.Add(unsafe.Pointer(v10), int32(uint8(int8(v27)))+4)) = 0
				if int32(int16(v26)) >= 3 || (func() bool {
					cryptFileReadWrite((*uint8)(unsafe.Pointer(&v31))[:4])
					cryptFileReadWrite((*uint8)(unsafe.Pointer(&v32[0]))[:12])
					*(*uint8)(unsafe.Add(unsafe.Pointer(v10), 104)) = uint8(v32[0])
					*(*uint8)(unsafe.Add(unsafe.Pointer(v10), 105)) = uint8(v32[4])
					*(*uint8)(unsafe.Add(unsafe.Pointer(v10), 106)) = uint8(v32[8])
					return int32(int16(v26)) >= 3
				}()) {
					cryptFileReadWrite((*uint8)(unsafe.Pointer(&a1))[:1])
					*(*uint8)(unsafe.Add(unsafe.Pointer(v10), 104)) = uint8(int8(a1))
					cryptFileReadWrite((*uint8)(unsafe.Pointer(&a1))[:1])
					*(*uint8)(unsafe.Add(unsafe.Pointer(v10), 105)) = uint8(int8(a1))
					cryptFileReadWrite((*uint8)(unsafe.Pointer(&a1))[:1])
					*(*uint8)(unsafe.Add(unsafe.Pointer(v10), 106)) = uint8(int8(a1))
				}
				cryptFileReadWrite((*uint8)(unsafe.Add(unsafe.Pointer(v10), 130))[:1])
				v11 = (*uint16)(unsafe.Add(unsafe.Pointer(v10), 128))
				cryptFileReadWrite((*uint8)(unsafe.Add(unsafe.Pointer(v10), 128))[:2])
				v12 = (*uint8)(alloc.Calloc(int(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v10))), unsafe.Sizeof(uint16(0))*64)))), 4))
				*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v10))), unsafe.Sizeof(uint32(0))*27))) = uint32(uintptr(unsafe.Pointer(v12)))
				if v12 == nil {
					return 0
				}
				cryptFileReadWrite(v12[:uint32(int32(*v11)*4)])
				sub_421040(int32(uintptr(unsafe.Pointer(v10))))
				v13 = nox_xxx_polygonGetAngle_421030(int32(**((**uint32)(unsafe.Add(unsafe.Pointer((**uint32)(unsafe.Pointer(v10))), unsafe.Sizeof((*uint32)(nil))*27)))))
				*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v10))), unsafe.Sizeof(uint32(0))*22))) = uint32(int(*((*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(v13))), unsafe.Sizeof(float32(0))*1)))))
				*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v10))), unsafe.Sizeof(uint32(0))*23))) = uint32(int(*((*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(v13))), unsafe.Sizeof(float32(0))*2)))))
				*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v10))), unsafe.Sizeof(uint32(0))*24))) = uint32(int(*((*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(v13))), unsafe.Sizeof(float32(0))*1)))))
				v14 = int(*((*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(v13))), unsafe.Sizeof(float32(0))*2))))
				v15 = 1
				v16 = int32(*v11) <= 1
				*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v10))), unsafe.Sizeof(uint32(0))*25))) = uint32(v14)
				if !v16 {
					for {
						v17 = nox_xxx_polygonGetAngle_421030(int32(*(*uint32)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v10))), unsafe.Sizeof(uint32(0))*27))) + v15*4)))))
						v18 = float64(*((*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(v17))), unsafe.Sizeof(float32(0))*1))))
						v29[0] = *((*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(v17))), unsafe.Sizeof(float32(0))*1)))
						if v18 >= float64(*((*int32)(unsafe.Add(unsafe.Pointer((*int32)(unsafe.Pointer(v10))), unsafe.Sizeof(int32(0))*22)))) {
							v19 = float64(*((*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(v17))), unsafe.Sizeof(float32(0))*1))))
							v29[0] = *((*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(v17))), unsafe.Sizeof(float32(0))*1)))
							if v19 > float64(*((*int32)(unsafe.Add(unsafe.Pointer((*int32)(unsafe.Pointer(v10))), unsafe.Sizeof(int32(0))*24)))) {
								*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v10))), unsafe.Sizeof(uint32(0))*24))) = uint32(int(v29[0]))
							}
						} else {
							*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v10))), unsafe.Sizeof(uint32(0))*22))) = uint32(int(v29[0]))
						}
						v20 = float64(*((*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(v17))), unsafe.Sizeof(float32(0))*2))))
						v29[0] = *((*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(v17))), unsafe.Sizeof(float32(0))*2)))
						if v20 >= float64(*((*int32)(unsafe.Add(unsafe.Pointer((*int32)(unsafe.Pointer(v10))), unsafe.Sizeof(int32(0))*23)))) {
							v21 = float64(*((*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(v17))), unsafe.Sizeof(float32(0))*2))))
							v29[0] = *((*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(v17))), unsafe.Sizeof(float32(0))*2)))
							if v21 > float64(*((*int32)(unsafe.Add(unsafe.Pointer((*int32)(unsafe.Pointer(v10))), unsafe.Sizeof(int32(0))*25)))) {
								*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v10))), unsafe.Sizeof(uint32(0))*25))) = uint32(int(v29[0]))
							}
						} else {
							*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v10))), unsafe.Sizeof(uint32(0))*23))) = uint32(int(v29[0]))
						}
						v15++
						if v15 >= uint32(*v11) {
							break
						}
					}
				}
				if int32(int16(v26)) >= 2 {
					v22 = *(**byte)(unsafe.Pointer(v10))
					nox_xxx_xferReadScriptHandler_4F5580(int32(uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v10), 112))))), *(**byte)(unsafe.Pointer(v10)))
					if v22 != nil {
						v23 = (*byte)(unsafe.Add(unsafe.Pointer(v22), 128))
					} else {
						v23 = nil
					}
					nox_xxx_xferReadScriptHandler_4F5580(int32(uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v10), 120))))), v23)
				}
				if int32(int16(v26)) >= 4 {
					v24 = (*uint8)(unsafe.Add(unsafe.Pointer(v10), 132))
					cryptFileReadWrite(v24[:4])
					if int32(*v24)&1 != 0 {
						v30++
					}
				}
				if func() uint32 {
					p := &v28
					*p++
					return *p
				}() > v25 {
					break
				}
			}
		}
		sub_4D72D0(v30)
		return 1
	}
	v25 = 0
	for i = nox_xxx_polygon_420CA0(); i != nil; v25++ {
		i = nox_xxx_polygon_420CD0((*uint32)(unsafe.Pointer(i)))
	}
	cryptFileReadWrite((*uint8)(unsafe.Pointer(&v25))[:4])
	v3 = nox_xxx_polygon_420CA0()
	for j = 0; j < v25; j++ {
		cryptFileReadWrite((*uint8)(unsafe.Pointer(v3))[:4])
		cryptFileReadWrite((*uint8)(unsafe.Add(unsafe.Pointer(v3), 4))[:8])
		v3 = nox_xxx_polygon_420CD0((*uint32)(unsafe.Pointer(v3)))
	}
	v25 = 0
	for k = nox_xxx_polygonGetNext_4210A0(); k != nil; v25++ {
		k = sub_4210E0(int32(uintptr(unsafe.Pointer(k))))
	}
	cryptFileReadWrite((*uint8)(unsafe.Pointer(&v25))[:4])
	v6 = nox_xxx_polygonGetNext_4210A0()
	v28 = 0
	if v25 <= 0 {
		return 1
	}
	for {
		v7 = *(**byte)(unsafe.Pointer(v6))
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v27))), 0)) = uint8(int8(libc.StrLen((*byte)(unsafe.Add(unsafe.Pointer(v6), 4)))))
		cryptFileReadWrite((*uint8)(unsafe.Pointer(&v27))[:1])
		cryptFileReadWrite((*uint8)(unsafe.Add(unsafe.Pointer(v6), 4))[:uint32(uint8(int8(v27)))])
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a1))), 0)) = uint8(*(*byte)(unsafe.Add(unsafe.Pointer(v6), 104)))
		cryptFileReadWrite((*uint8)(unsafe.Pointer(&a1))[:1])
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a1))), 0)) = uint8(*(*byte)(unsafe.Add(unsafe.Pointer(v6), 105)))
		cryptFileReadWrite((*uint8)(unsafe.Pointer(&a1))[:1])
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a1))), 0)) = uint8(*(*byte)(unsafe.Add(unsafe.Pointer(v6), 106)))
		cryptFileReadWrite((*uint8)(unsafe.Pointer(&a1))[:1])
		cryptFileReadWrite((*uint8)(unsafe.Add(unsafe.Pointer(v6), 130))[:1])
		cryptFileReadWrite((*uint8)(unsafe.Add(unsafe.Pointer(v6), 128))[:2])
		cryptFileReadWrite((*((**uint8)(unsafe.Add(unsafe.Pointer((**uint8)(unsafe.Pointer(v6))), unsafe.Sizeof((*uint8)(nil))*27))))[:uint32(int32(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v6))), unsafe.Sizeof(uint16(0))*64))))*4)])
		nox_xxx_xferReadScriptHandler_4F5580(int32(uintptr(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v6), 112))))), v7)
		if v7 != nil {
			v8 = (*byte)(unsafe.Add(unsafe.Pointer(v7), 128))
		} else {
			v8 = nil
		}
		nox_xxx_xferReadScriptHandler_4F5580(int32(uintptr(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v6), 120))))), v8)
		cryptFileReadWrite((*uint8)(unsafe.Add(unsafe.Pointer(v6), 132))[:4])
		v6 = sub_4210E0(int32(uintptr(unsafe.Pointer(v6))))
		v28++
		if v28 >= v25 {
			break
		}
	}
	return 1
}
func nox_server_mapRWAmbientData_429200() int32 {
	var (
		result int32
		v1     *byte
		v2     int32
		v3     [3]int32
	)
	v2 = 1
	cryptFileReadWrite((*uint8)(unsafe.Pointer(&v2))[:2])
	if int32(int16(v2)) < 1 {
		return 0
	}
	if nox_xxx_cryptGetXxx() != 0 {
		if nox_xxx_cryptGetXxx() == 1 {
			cryptFileReadWrite((*uint8)(unsafe.Pointer(&v3[0]))[:4])
			cryptFileReadWrite((*uint8)(unsafe.Pointer(&v3[1]))[:4])
			cryptFileReadWrite((*uint8)(unsafe.Pointer(&v3[2]))[:4])
			sub_469B90(&v3[0])
			if noxflags.HasGame(noxflags.GameClient | noxflags.GameFlag22) {
				sub_4349C0((*uint32)(unsafe.Pointer(&v3[0])))
			}
		}
		result = 1
	} else {
		v1 = nox_xxx_getAmbientColor_469BB0()
		cryptFileReadWrite((*uint8)(unsafe.Pointer(v1))[:4])
		cryptFileReadWrite((*uint8)(unsafe.Add(unsafe.Pointer(v1), 4))[:4])
		cryptFileReadWrite((*uint8)(unsafe.Add(unsafe.Pointer(v1), 8))[:4])
		result = 1
	}
	return result
}
func nox_server_mapRWWindowWalls_4292C0(a1 *uint32) int32 {
	var (
		result int32
		v2     *uint32
		v3     *byte
		v4     *uint32
		v5     int32
		v6     int32
		v7     int2
		v9     int4
	)
	v5 = 2
	cryptFileReadWrite((*uint8)(unsafe.Pointer(&v5))[:2])
	if int32(int16(v5)) > 2 {
		return 0
	}
	if nox_xxx_cryptGetXxx() != 0 {
		cryptFileReadWrite((*uint8)(memmap.PtrOff(6112660, 741336))[:2])
		v6 = 0
		if int32(*memmap.PtrUint16(6112660, 741336)) > 0 {
			v2 = a1
			for {
				cryptFileReadWrite((*uint8)(unsafe.Pointer(&v7))[:8])
				if a1 != nil {
					v3 = nox_xxx_mapGetWallSize_426A70()
					sub_428170(unsafe.Pointer(a1), &v9)
					v7.field_0 += int32(uint32(v9.field_0/23) - *(*uint32)(unsafe.Pointer(v3)))
					v7.field_4 += int32(uint32(v9.field_4/23) - *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v3))), unsafe.Sizeof(uint32(0))*1))))
				}
				if noxflags.HasGame(noxflags.GameFlag23) {
					v4 = nox_xxx_cliWallGet_5042F0(v7.field_0, v7.field_4)
					if v4 != nil {
						v2 = (*uint32)(unsafe.Pointer(uintptr(*v4)))
					}
				} else {
					v2 = (*uint32)(nox_server_getWallAtGrid_410580(v7.field_0, v7.field_4))
				}
				if v2 != nil {
					*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(v2))), 4))) |= 64
					if int32(int16(v5)) < 2 {
						*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(v2))), 2))) = 0
					}
				}
				v6++
				if v6 >= int32(*memmap.PtrInt16(6112660, 741336)) {
					break
				}
			}
		}
		result = 1
	} else {
		*memmap.PtrUint16(6112660, 741336) = 0
		nox_xxx_wallForeachFn_410640(func(arg1 int32, arg2 int32) {
			sub_429450((*uint8)(unsafe.Pointer(uintptr(arg1))), (*uint32)(unsafe.Pointer(uintptr(arg2))))
		}, int32(uintptr(unsafe.Pointer(a1))))
		cryptFileReadWrite((*uint8)(memmap.PtrOff(6112660, 741336))[:2])
		nox_xxx_wallForeachFn_410640(func(arg1 int32, arg2 int32) {
			sub_4294B0((*uint8)(unsafe.Pointer(uintptr(arg1))), (*uint32)(unsafe.Pointer(uintptr(arg2))))
		}, int32(uintptr(unsafe.Pointer(a1))))
		result = 1
	}
	return result
}
func sub_429450(a1 *uint8, a2 *uint32) {
	var (
		v2 int32
		v3 int2
	)
	if a2 == nil || (func() int32 {
		v2 = int32(*(*uint8)(unsafe.Add(unsafe.Pointer(a1), 6)))
		v3.field_0 = int32(*(*uint8)(unsafe.Add(unsafe.Pointer(a1), 5))) * 23
		v3.field_4 = v2 * 23
		return nox_xxx_wallMath_427F30(&v3, (*int32)(unsafe.Pointer(a2)))
	}()) != 0 {
		if int32(*(*uint8)(unsafe.Add(unsafe.Pointer(a1), 4)))&64 != 0 {
			*memmap.PtrUint16(6112660, 741336)++
		}
	}
}
func sub_4294B0(a1 *uint8, a2 *uint32) {
	var (
		v2 int32
		v3 int32
		v4 int2
	)
	if a2 == nil || (func() int32 {
		v2 = int32(*(*uint8)(unsafe.Add(unsafe.Pointer(a1), 6)))
		v4.field_0 = int32(*(*uint8)(unsafe.Add(unsafe.Pointer(a1), 5))) * 23
		v4.field_4 = v2 * 23
		return nox_xxx_wallMath_427F30(&v4, (*int32)(unsafe.Pointer(a2)))
	}()) != 0 {
		if int32(*(*uint8)(unsafe.Add(unsafe.Pointer(a1), 4)))&64 != 0 {
			v3 = int32(*(*uint8)(unsafe.Add(unsafe.Pointer(a1), 6)))
			v4.field_0 = int32(*(*uint8)(unsafe.Add(unsafe.Pointer(a1), 5)))
			v4.field_4 = v3
			cryptFileReadWrite((*uint8)(unsafe.Pointer(&v4))[:8])
		}
	}
}
func nox_xxx_wallBreakableCounterClear_429520() {
	*memmap.PtrUint32(6112660, 741344) = 0
}
func nox_server_mapRWDestructableWalls_429530(a1 *uint32) int32 {
	var (
		result int32
		v2     *uint32
		v3     *byte
		v4     *uint32
		v5     int32
		v6     int32
		v7     int2
		v9     int4
	)
	v5 = 1
	cryptFileReadWrite((*uint8)(unsafe.Pointer(&v5))[:2])
	if int32(int16(v5)) > 1 {
		return 0
	}
	if nox_xxx_cryptGetXxx() != 0 {
		cryptFileReadWrite((*uint8)(memmap.PtrOff(6112660, 741340))[:2])
		v6 = 0
		if int32(*memmap.PtrUint16(6112660, 741340)) > 0 {
			v2 = a1
			for {
				cryptFileReadWrite((*uint8)(unsafe.Pointer(&v7))[:8])
				if a1 != nil {
					v3 = nox_xxx_mapGetWallSize_426A70()
					sub_428170(unsafe.Pointer(a1), &v9)
					v7.field_0 += int32(uint32(v9.field_0/23) - *(*uint32)(unsafe.Pointer(v3)))
					v7.field_4 += int32(uint32(v9.field_4/23) - *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v3))), unsafe.Sizeof(uint32(0))*1))))
				}
				if noxflags.HasGame(noxflags.GameFlag23) {
					v4 = nox_xxx_cliWallGet_5042F0(v7.field_0, v7.field_4)
					if v4 != nil {
						v2 = (*uint32)(unsafe.Pointer(uintptr(*v4)))
					}
				} else {
					v2 = (*uint32)(nox_server_getWallAtGrid_410580(v7.field_0, v7.field_4))
				}
				if v2 != nil {
					*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(v2))), 4))) |= 8
					*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v2))), unsafe.Sizeof(uint16(0))*5))) = *memmap.PtrUint16(6112660, 741344)
					*memmap.PtrUint32(6112660, 741344)++
					if !noxflags.HasGame(noxflags.GameFlag23) {
						nox_xxx_wallBreackableListAdd_410840(int32(uintptr(unsafe.Pointer(v2))))
					}
				}
				v6++
				if v6 >= int32(*memmap.PtrInt16(6112660, 741340)) {
					break
				}
			}
		}
		result = 1
	} else {
		*memmap.PtrUint16(6112660, 741340) = 0
		nox_xxx_wallForeachFn_410640(func(arg1 int32, arg2 int32) {
			nox_xxx_wall_4296E0((*uint8)(unsafe.Pointer(uintptr(arg1))), (*uint32)(unsafe.Pointer(uintptr(arg2))))
		}, int32(uintptr(unsafe.Pointer(a1))))
		cryptFileReadWrite((*uint8)(memmap.PtrOff(6112660, 741340))[:2])
		nox_xxx_wallForeachFn_410640(func(arg1 int32, arg2 int32) {
			sub_429740((*uint8)(unsafe.Pointer(uintptr(arg1))), (*uint32)(unsafe.Pointer(uintptr(arg2))))
		}, int32(uintptr(unsafe.Pointer(a1))))
		result = 1
	}
	return result
}
func nox_xxx_wall_4296E0(a1 *uint8, a2 *uint32) {
	var (
		v2 int32
		v3 int2
	)
	if a2 == nil || (func() int32 {
		v2 = int32(*(*uint8)(unsafe.Add(unsafe.Pointer(a1), 6)))
		v3.field_0 = int32(*(*uint8)(unsafe.Add(unsafe.Pointer(a1), 5))) * 23
		v3.field_4 = v2 * 23
		return nox_xxx_wallMath_427F30(&v3, (*int32)(unsafe.Pointer(a2)))
	}()) != 0 {
		if int32(*(*uint8)(unsafe.Add(unsafe.Pointer(a1), 4)))&8 != 0 {
			*memmap.PtrUint16(6112660, 741340)++
		}
	}
}
func sub_429740(a1 *uint8, a2 *uint32) {
	var (
		v2 int32
		v3 int32
		v4 int2
	)
	if a2 == nil || (func() int32 {
		v2 = int32(*(*uint8)(unsafe.Add(unsafe.Pointer(a1), 6)))
		v4.field_0 = int32(*(*uint8)(unsafe.Add(unsafe.Pointer(a1), 5))) * 23
		v4.field_4 = v2 * 23
		return nox_xxx_wallMath_427F30(&v4, (*int32)(unsafe.Pointer(a2)))
	}()) != 0 {
		if int32(*(*uint8)(unsafe.Add(unsafe.Pointer(a1), 4)))&8 != 0 {
			v3 = int32(*(*uint8)(unsafe.Add(unsafe.Pointer(a1), 6)))
			v4.field_0 = int32(*(*uint8)(unsafe.Add(unsafe.Pointer(a1), 5)))
			v4.field_4 = v3
			cryptFileReadWrite((*uint8)(unsafe.Pointer(&v4))[:8])
		}
	}
}
func nox_xxx_wallSecretCounterClear_4297B0() {
	*memmap.PtrUint32(6112660, 741352) = 0
}
func nox_server_mapRWSecretWalls_4297C0(a1 *uint32) int32 {
	var (
		v2  *byte
		v3  *int32
		v4  *uint8
		v5  *byte
		v6  *int32
		v7  int32
		v8  int8
		v9  int32
		v10 int32 = 0
		v11 int32
		v12 int4
	)
	v9 = 2
	cryptFileReadWrite((*uint8)(unsafe.Pointer(&v9))[:2])
	if int32(int16(v9)) > 2 {
		return 0
	}
	if nox_xxx_cryptGetXxx() == 0 {
		*memmap.PtrUint16(6112660, 741348) = 0
		nox_xxx_wallForeachFn_410640(func(arg1 int32, arg2 int32) {
			sub_429A00((*uint8)(unsafe.Pointer(uintptr(arg1))), (*uint32)(unsafe.Pointer(uintptr(arg2))))
		}, int32(uintptr(unsafe.Pointer(a1))))
		cryptFileReadWrite((*uint8)(memmap.PtrOff(6112660, 741348))[:2])
		nox_xxx_wallForeachFn_410640(func(arg1 int32, arg2 int32) {
			sub_429A60(arg1, (*uint32)(unsafe.Pointer(uintptr(arg2))))
		}, int32(uintptr(unsafe.Pointer(a1))))
		return 1
	}
	cryptFileReadWrite((*uint8)(memmap.PtrOff(6112660, 741348))[:2])
	v11 = 0
	if int32(*memmap.PtrUint16(6112660, 741348)) > 0 {
		for {
			v2 = (*byte)(alloc.Calloc(1, 32))
			v3 = (*int32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v2), 4))))
			cryptFileReadWrite((*uint8)(unsafe.Add(unsafe.Pointer(v2), 4))[:8])
			cryptFileReadWrite((*uint8)(unsafe.Add(unsafe.Pointer(v2), 16))[:4])
			cryptFileReadWrite((*uint8)(unsafe.Add(unsafe.Pointer(v2), 20))[:1])
			v4 = (*uint8)(unsafe.Add(unsafe.Pointer(v2), 21))
			*(*byte)(unsafe.Add(unsafe.Pointer(v2), 21)) = 0
			if int32(int16(v9)) >= 2 {
				cryptFileReadWrite((*uint8)(unsafe.Add(unsafe.Pointer(v2), 21))[:1])
				cryptFileReadWrite((*uint8)(unsafe.Add(unsafe.Pointer(v2), 22))[:1])
				cryptFileReadWrite((*uint8)(unsafe.Add(unsafe.Pointer(v2), 24))[:4])
				cryptFileReadWrite((*uint8)(unsafe.Add(unsafe.Pointer(v2), 28))[:4])
			}
			if a1 != nil {
				v5 = nox_xxx_mapGetWallSize_426A70()
				sub_428170(unsafe.Pointer(a1), &v12)
				*v3 += int32(uint32(v12.field_0/23) - *(*uint32)(unsafe.Pointer(v5)))
				*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), unsafe.Sizeof(uint32(0))*2))) += uint32(v12.field_4/23) - *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v5))), unsafe.Sizeof(uint32(0))*1)))
			}
			if !noxflags.HasGame(noxflags.GameFlag23) {
				break
			}
			v6 = (*int32)(unsafe.Pointer(nox_xxx_cliWallGet_5042F0(*v3, int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), unsafe.Sizeof(uint32(0))*2)))))))
			if v6 == nil {
				goto LABEL_14
			}
			v7 = *v6
			v10 = v7
		LABEL_15:
			if v7 != 0 {
				v8 = int8(*(*uint8)(unsafe.Pointer(uintptr(v7 + 4))))
				*(*uint32)(unsafe.Pointer(uintptr(v7 + 28))) = uint32(uintptr(unsafe.Pointer(v2)))
				*(*uint8)(unsafe.Pointer(uintptr(v7 + 4))) = uint8(int8(int32(v8) | 4))
				*(*uint16)(unsafe.Pointer(uintptr(v7 + 10))) = *memmap.PtrUint16(6112660, 741352)
				*memmap.PtrUint32(6112660, 741352)++
				*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), unsafe.Sizeof(uint32(0))*3))) = uint32(v7)
				if int32(*v4) == 0 {
					if *(*byte)(unsafe.Add(unsafe.Pointer(v2), 20))&8 != 0 {
						*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), unsafe.Sizeof(uint32(0))*7))) = math.MaxUint32
						*v4 = 3
						*(*byte)(unsafe.Add(unsafe.Pointer(v2), 22)) = 23
					} else {
						*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), unsafe.Sizeof(uint32(0))*7))) = 0
						*v4 = 1
						*(*byte)(unsafe.Add(unsafe.Pointer(v2), 22)) = 0
					}
				}
				if !noxflags.HasGame(noxflags.GameFlag23) {
					nox_xxx_wallSecretBlock_410760((*uint32)(unsafe.Pointer(v2)))
				}
			}
			if func() int32 {
				p := &v11
				*p++
				return *p
			}() >= int32(*memmap.PtrInt16(6112660, 741348)) {
				return 1
			}
		}
		v10 = int32(uintptr(nox_server_getWallAtGrid_410580(*v3, int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), unsafe.Sizeof(uint32(0))*2)))))))
	LABEL_14:
		v7 = v10
		goto LABEL_15
	}
	return 1
}
func sub_429A00(a1 *uint8, a2 *uint32) {
	var (
		v2 int32
		v3 int2
	)
	if a2 == nil || (func() int32 {
		v2 = int32(*(*uint8)(unsafe.Add(unsafe.Pointer(a1), 6)))
		v3.field_0 = int32(*(*uint8)(unsafe.Add(unsafe.Pointer(a1), 5))) * 23
		v3.field_4 = v2 * 23
		return nox_xxx_wallMath_427F30(&v3, (*int32)(unsafe.Pointer(a2)))
	}()) != 0 {
		if int32(*(*uint8)(unsafe.Add(unsafe.Pointer(a1), 4)))&4 != 0 {
			*memmap.PtrUint16(6112660, 741348)++
		}
	}
}
func sub_429A60(a1 int32, a2 *uint32) {
	var (
		v2 int32
		v3 int32
		v4 *uint8
		v5 int2
	)
	if a2 == nil || (func() int32 {
		v2 = int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 6))))
		v5.field_0 = int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 5)))) * 23
		v5.field_4 = v2 * 23
		return nox_xxx_wallMath_427F30(&v5, (*int32)(unsafe.Pointer(a2)))
	}()) != 0 {
		if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 4))))&4 != 0 {
			v3 = int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 6))))
			v4 = *(**uint8)(unsafe.Pointer(uintptr(a1 + 28)))
			v5.field_0 = int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 5))))
			v5.field_4 = v3
			cryptFileReadWrite((*uint8)(unsafe.Pointer(&v5))[:8])
			cryptFileReadWrite((*uint8)(unsafe.Add(unsafe.Pointer(v4), 16))[:4])
			cryptFileReadWrite((*uint8)(unsafe.Add(unsafe.Pointer(v4), 20))[:1])
			cryptFileReadWrite((*uint8)(unsafe.Add(unsafe.Pointer(v4), 21))[:1])
			cryptFileReadWrite((*uint8)(unsafe.Add(unsafe.Pointer(v4), 22))[:1])
			cryptFileReadWrite((*uint8)(unsafe.Add(unsafe.Pointer(v4), 24))[:4])
			cryptFileReadWrite((*uint8)(unsafe.Add(unsafe.Pointer(v4), 28))[:4])
		}
	}
}
func nox_server_mapRWWallMap_429B20(a1 *uint32) int32 {
	var (
		v2  *uint32
		v3  int32
		v5  int32
		v6  int32
		v7  int32
		v8  int32
		v9  int32
		v10 *uint8
		v11 *uint8
		v12 *nox_player_polygon_check_data
		v13 int32
		v14 int32
		v15 int8
		v16 int8
		v17 *uint8
		v18 int32
		v19 *uint8
		v20 *uint8
		v21 *uint8
		v22 int8
		v23 int8
		v24 int32
		v25 int32
		v26 int32
		v27 int32
		v28 int32
		v29 int32
		v30 int32
		v31 int32
		v32 int2
		v33 int4
	)
	v31 = nox_xxx_wallGet_426A30()
	if int32(*memmap.PtrUint8(6112660, 741372)) == 0 {
		*memmap.PtrUint8(6112660, 741372) = uint8(int8(nox_xxx_wallTileByName_410D60(CString("MagicWallSystemUseOnly"))))
	}
	nox_xxx_wallSecretCounterClear_4297B0()
	nox_xxx_wallBreakableCounterClear_429520()
	v28 = 7
	cryptFileReadWrite((*uint8)(unsafe.Pointer(&v28))[:2])
	if int32(int16(v28)) > 7 {
		return 0
	}
	if int32(int16(v28)) < 6 {
		return sub_42A150(int16(v28), a1)
	}
	v2 = a1
	if nox_xxx_cryptGetXxx() == 0 {
		if a1 != nil {
			sub_428170(unsafe.Pointer(a1), &v33)
			*memmap.PtrUint32(6112660, 741360) = uint32(v33.field_0 / 23)
			v3 = v33.field_8 / 23
			dword_5d4594_741356 = uint32(v33.field_8 / 23)
			*memmap.PtrUint32(6112660, 741368) = uint32(v33.field_4 / 23)
			v5 = v33.field_C / 23
			dword_5d4594_741364 = uint32(v5)
		} else {
			*memmap.PtrUint32(6112660, 741368) = 256
			*memmap.PtrUint32(6112660, 741360) = 256
			dword_5d4594_741364 = 0
			dword_5d4594_741356 = 0
			nox_xxx_wallForeachFn_410640(func(arg1 int32, arg2 int32) {
				sub_42A0F0(arg1)
			}, 0)
			v3 = int32(dword_5d4594_741356)
			v5 = int32(dword_5d4594_741364)
		}
		v25 = int32(uint32(v3) - *memmap.PtrUint32(6112660, 741360) + 1)
		v27 = int32(uint32(v5) - *memmap.PtrUint32(6112660, 741368) + 1)
	}
	cryptFileReadWrite((*uint8)(memmap.PtrOff(6112660, 741360))[:4])
	cryptFileReadWrite((*uint8)(memmap.PtrOff(6112660, 741368))[:4])
	cryptFileReadWrite((*uint8)(unsafe.Pointer(&v25))[:4])
	cryptFileReadWrite((*uint8)(unsafe.Pointer(&v27))[:4])
	v26 = 0
	v29 = 0
	if nox_xxx_cryptGetXxx() != 0 {
		if v2 != nil {
			sub_428170(unsafe.Pointer(v2), &v33)
			v13 = int32(uint32(v33.field_0/23) - *memmap.PtrUint32(6112660, 741360))
			*memmap.PtrUint32(6112660, 741360) = uint32(v33.field_0 / 23)
			v29 = v13
			v14 = int32(uint32(v33.field_4/23) - *memmap.PtrUint32(6112660, 741368))
			*memmap.PtrUint32(6112660, 741368) = uint32(v33.field_4 / 23)
			v26 = v14
		}
		cryptFileReadWrite((*uint8)(unsafe.Pointer(&v22))[:1])
		v15 = v22
		if int32(v22) == -1 {
			return 1
		}
		for {
			cryptFileReadWrite((*uint8)(unsafe.Pointer(&a1))[:1])
			*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v30))), 0)) = uint8(int8(v29 + int32(v15)))
			*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a1))), 0)) = uint8(int8(v26 + int32(uint8(uintptr(unsafe.Pointer(a1))))))
			cryptFileReadWrite((*uint8)(unsafe.Pointer(&v24))[:1])
			v16 = int8(int32(uint8(int8(v24))) >> 7)
			*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v24))), 0)) = uint8(int8(v24 & math.MaxInt8))
			if noxflags.HasGame(noxflags.GameFlag23) {
				v17 = (*uint8)(unsafe.Pointer(uintptr(*sub_504290(int8(v30), int8(uintptr(unsafe.Pointer(a1)))))))
			} else {
				v18 = int32(uint8(int8(v30)))
				v19 = (*uint8)(nox_server_getWallAtGrid_410580(int32(uint8(int8(v30))), int32(uint8(uintptr(unsafe.Pointer(a1))))))
				v17 = v19
				if v19 != nil {
					if v31&1 != 0 {
						*v19 = nox_xxx_wall_42A6C0(*v19, uint8(int8(v24)))
					} else {
						*v19 = uint8(int8(v24))
					}
					goto LABEL_46
				}
				v17 = (*uint8)(nox_xxx_wallCreateAt_410250(v18, int32(uint8(uintptr(unsafe.Pointer(a1))))))
				if v17 == nil {
					return 0
				}
			}
			*v17 = uint8(int8(v24))
		LABEL_46:
			if int32(v16) != 0 {
				*(*uint8)(unsafe.Add(unsafe.Pointer(v17), 4)) |= 128
			}
			v20 = (*uint8)(unsafe.Add(unsafe.Pointer(v17), 1))
			cryptFileReadWrite((*uint8)(unsafe.Add(unsafe.Pointer(v17), 1))[:1])
			v21 = (*uint8)(unsafe.Add(unsafe.Pointer(v17), 2))
			cryptFileReadWrite((*uint8)(unsafe.Add(unsafe.Pointer(v17), 2))[:1])
			if v31&1 != 0 && int32(*v21) >= int32(nox_xxx_mapWallMaxVariation_410DD0(int32(*v20), int32(*v17), 0)) {
				*v21 = 0
			}
			*(*uint8)(unsafe.Add(unsafe.Pointer(v17), 7)) = nox_xxx_mapWallGetHpByTile_410E20(int32(*v20))
			if int32(uint16(int16(v28))) == 6 {
				cryptFileReadWrite((*uint8)(unsafe.Pointer(&v24))[:1])
				*(*uint8)(unsafe.Add(unsafe.Pointer(v17), 8)) = 1
				*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v17))), unsafe.Sizeof(uint32(0))*3))) = 0
			} else {
				cryptFileReadWrite((*uint8)(unsafe.Add(unsafe.Pointer(v17), 8))[:1])
				cryptFileReadWrite((*uint8)(unsafe.Pointer(&v24))[:1])
				*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v17))), unsafe.Sizeof(uint32(0))*3))) = uint32(uint8(int8(v24)))
			}
			cryptFileReadWrite((*uint8)(unsafe.Pointer(&v22))[:1])
			v15 = v22
			if int32(v22) == -1 {
				return 1
			}
		}
	}
	v6 = int32(*memmap.PtrUint32(6112660, 741368))
	v26 = int32(*memmap.PtrUint32(6112660, 741368))
	if *memmap.PtrUint32(6112660, 741368) <= uint32(*memmap.PtrInt32(6112660, 741368)+v27) {
		v7 = v25
		v8 = int32(*memmap.PtrUint32(6112660, 741360))
		for {
			v9 = v8
			v29 = v8
			if v8 <= v8+v7 {
				for {
					v10 = (*uint8)(nox_server_getWallAtGrid_410580(v9, v26))
					v11 = v10
					if v10 != nil {
						if int32(*(*uint8)(unsafe.Add(unsafe.Pointer(v10), 1))) != int32(*memmap.PtrUint8(6112660, 741372)) {
							if v2 == nil || (func() int32 {
								v32.field_0 = int32(*(*uint8)(unsafe.Add(unsafe.Pointer(v10), 5))) * 23
								v32.field_4 = int32(*(*uint8)(unsafe.Add(unsafe.Pointer(v10), 6))) * 23
								return nox_xxx_wallMath_427F30(&v32, (*int32)(unsafe.Pointer(v2)))
							}()) != 0 {
								cryptFileReadWrite((*uint8)(unsafe.Add(unsafe.Pointer(v11), 5))[:1])
								cryptFileReadWrite((*uint8)(unsafe.Add(unsafe.Pointer(v11), 6))[:1])
								if int32(*(*uint8)(unsafe.Add(unsafe.Pointer(v11), 4))) >= 0 {
									*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v24))), 0)) = *v11
								} else {
									*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v24))), 0)) = uint8(int8(int32(*v11) | 128))
								}
								cryptFileReadWrite((*uint8)(unsafe.Pointer(&v24))[:1])
								cryptFileReadWrite((*uint8)(unsafe.Add(unsafe.Pointer(v11), 1))[:1])
								cryptFileReadWrite((*uint8)(unsafe.Add(unsafe.Pointer(v11), 2))[:1])
								v32.field_0 = int32(*(*uint8)(unsafe.Add(unsafe.Pointer(v11), 5)))*23 + 11
								v32.field_4 = int32(*(*uint8)(unsafe.Add(unsafe.Pointer(v11), 6)))*23 + 11
								v12 = nox_xxx_polygonIsPlayerInPolygon_4217B0(&v32, 0)
								if v12 != nil || (func() *nox_player_polygon_check_data {
									v12 = (*nox_player_polygon_check_data)(unsafe.Pointer(sub_421990(&v32, 10.0, 0)))
									return v12
								}()) != nil {
									v23 = int8(*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v12.field_0[32]))), 2)))
								} else {
									v23 = 100
								}
								cryptFileReadWrite((*uint8)(unsafe.Pointer(&v23))[:1])
								if noxflags.HasGame(noxflags.GameFlag22) {
									*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v24))), 0)) = 0
								} else {
									*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v24))), 0)) = *(*uint8)(unsafe.Add(unsafe.Pointer(v11), 12))
								}
								cryptFileReadWrite((*uint8)(unsafe.Pointer(&v24))[:1])
								v2 = a1
								v9 = v29
							}
						}
					}
					v8 = int32(*memmap.PtrUint32(6112660, 741360))
					v7 = v25
					v29 = func() int32 {
						p := &v9
						*p++
						return *p
					}()
					if v9 > *memmap.PtrInt32(6112660, 741360)+v25 {
						break
					}
				}
				v6 = int32(*memmap.PtrUint32(6112660, 741368))
			}
			v26++
			if v26 > v6+v27 {
				break
			}
		}
	}
	v22 = -1
	cryptFileReadWrite((*uint8)(unsafe.Pointer(&v22))[:1])
	return 1
}
func sub_42A0F0(a1 int32) int32 {
	var result int32
	if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 5)))) < *memmap.PtrInt32(6112660, 741360) {
		*memmap.PtrUint32(6112660, 741360) = uint32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 5))))
	}
	if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 5)))) > *(*int32)(unsafe.Pointer(&dword_5d4594_741356)) {
		dword_5d4594_741356 = uint32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 5))))
	}
	if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 6)))) < *memmap.PtrInt32(6112660, 741368) {
		*memmap.PtrUint32(6112660, 741368) = uint32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 6))))
	}
	result = int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 6))))
	if result > *(*int32)(unsafe.Pointer(&dword_5d4594_741364)) {
		dword_5d4594_741364 = uint32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 6))))
	}
	return result
}
func sub_42A150(a1 int16, a2 *uint32) int32 {
	var (
		v2  int32
		v3  *uint32
		v4  int32
		v5  int32
		v6  int32
		v7  int32
		v8  int32
		v9  int32
		v10 int32
		v11 int32
		v12 int32
		v13 *uint8
		v14 *nox_player_polygon_check_data
		v16 int32
		v17 int32
		v18 uint8
		v19 int8
		v20 **uint8
		v21 *uint8
		v22 *uint8
		v23 uint8
		v24 int16
		v25 *uint8
		v26 int8
		v27 int32
		v28 int32
		v29 int32
		v30 int32
		v31 int32
		v32 int2
		v33 int4
	)
	v2 = nox_xxx_wallGet_426A30()
	v3 = a2
	v30 = v2
	v4 = 0
	if nox_xxx_cryptGetXxx() == 0 {
		if a2 != nil {
			sub_428170(unsafe.Pointer(a2), &v33)
			*memmap.PtrUint32(6112660, 741360) = uint32(v33.field_0 / 23)
			v5 = v33.field_8 / 23
			dword_5d4594_741356 = uint32(v33.field_8 / 23)
			*memmap.PtrUint32(6112660, 741368) = uint32(v33.field_4 / 23)
			v6 = v33.field_C / 23
			dword_5d4594_741364 = uint32(v33.field_C / 23)
		} else {
			*memmap.PtrUint32(6112660, 741368) = 256
			*memmap.PtrUint32(6112660, 741360) = 256
			dword_5d4594_741364 = 0
			dword_5d4594_741356 = 0
			nox_xxx_wallForeachFn_410640(func(arg1 int32, arg2 int32) {
				sub_42A0F0(arg1)
			}, 0)
			v5 = int32(dword_5d4594_741356)
			v6 = int32(dword_5d4594_741364)
		}
		v28 = int32(uint32(v5) - *memmap.PtrUint32(6112660, 741360) + 1)
		v29 = int32(uint32(v6) - *memmap.PtrUint32(6112660, 741368) + 1)
	}
	cryptFileReadWrite((*uint8)(memmap.PtrOff(6112660, 741360))[:4])
	cryptFileReadWrite((*uint8)(memmap.PtrOff(6112660, 741368))[:4])
	cryptFileReadWrite((*uint8)(unsafe.Pointer(&v28))[:4])
	cryptFileReadWrite((*uint8)(unsafe.Pointer(&v29))[:4])
	if nox_xxx_cryptGetXxx() == 0 {
		v7 = int32(*memmap.PtrUint32(6112660, 741368))
		v8 = int32(*memmap.PtrUint32(6112660, 741368))
		if *memmap.PtrUint32(6112660, 741368) > uint32(*memmap.PtrInt32(6112660, 741368)+v29) {
			return 1
		}
		v9 = v28
		v10 = int32(*memmap.PtrUint32(6112660, 741360))
		for {
			v11 = v10
			if v10 > v10+v9 {
				goto LABEL_27
			}
			for {
				v12 = int32(uintptr(nox_server_getWallAtGrid_410580(v11, v8)))
				v13 = (*uint8)(unsafe.Pointer(uintptr(v12)))
				if v3 != nil {
					if v12 == 0 {
						goto LABEL_15
					}
					v32.field_0 = int32(*(*uint8)(unsafe.Pointer(uintptr(v12 + 5)))) * 23
					v32.field_4 = int32(*(*uint8)(unsafe.Pointer(uintptr(v12 + 6)))) * 23
					if nox_xxx_wallMath_427F30(&v32, (*int32)(unsafe.Pointer(v3))) == 0 {
						v13 = nil
						*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v27))), 0)) = math.MaxUint8
						goto LABEL_19
					}
				}
				if v13 == nil {
				LABEL_15:
					*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v27))), 0)) = math.MaxUint8
					goto LABEL_19
				}
				if int32(*(*uint8)(unsafe.Add(unsafe.Pointer(v13), 4))) >= 0 {
					*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v27))), 0)) = *v13
				} else {
					*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v27))), 0)) = uint8(int8(int32(*v13) | 128))
				}
			LABEL_19:
				cryptFileReadWrite((*uint8)(unsafe.Pointer(&v27))[:1])
				if int32(uint8(int8(v27))) != math.MaxUint8 {
					cryptFileReadWrite((*uint8)(unsafe.Add(unsafe.Pointer(v13), 1))[:1])
					cryptFileReadWrite((*uint8)(unsafe.Add(unsafe.Pointer(v13), 2))[:1])
					v32.field_0 = int32(*(*uint8)(unsafe.Add(unsafe.Pointer(v13), 5)))*23 + 11
					v32.field_4 = int32(*(*uint8)(unsafe.Add(unsafe.Pointer(v13), 6)))*23 + 11
					v14 = nox_xxx_polygonIsPlayerInPolygon_4217B0(&v32, 0)
					if v14 != nil || (func() *nox_player_polygon_check_data {
						v14 = (*nox_player_polygon_check_data)(unsafe.Pointer(sub_421990(&v32, 10.0, 0)))
						return v14
					}()) != nil {
						v26 = int8(*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v14.field_0[32]))), 2)))
					} else {
						v26 = 1
					}
					cryptFileReadWrite((*uint8)(unsafe.Pointer(&v26))[:1])
					*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a1))), 0)) = *(*uint8)(unsafe.Add(unsafe.Pointer(v13), 12))
					cryptFileReadWrite((*uint8)(unsafe.Pointer(&a1))[:1])
				}
				v10 = int32(*memmap.PtrUint32(6112660, 741360))
				v9 = v28
				v11++
				if v11 > *memmap.PtrInt32(6112660, 741360)+v28 {
					break
				}
			}
			v7 = int32(*memmap.PtrUint32(6112660, 741368))
		LABEL_27:
			if func() int32 {
				p := &v8
				*p++
				return *p
			}() > v7+v29 {
				return 1
			}
		}
	}
	if v3 != nil {
		sub_428170(unsafe.Pointer(v3), &v33)
		*memmap.PtrUint32(6112660, 741360) = uint32(v33.field_0 / 23)
		*memmap.PtrUint32(6112660, 741368) = uint32(v33.field_4 / 23)
	}
	v31 = 0
	if v29 < 0 {
		return 1
	}
	v16 = v28
	for {
		v17 = 0
		if v16 >= 0 {
			break
		}
	LABEL_60:
		v31 = func() int32 {
			p := &v4
			*p++
			return *p
		}()
		if v4 > v29 {
			return 1
		}
	}
	for {
		cryptFileReadWrite((*uint8)(unsafe.Pointer(&v27))[:1])
		if int32(uint8(int8(v27))) != math.MaxUint8 {
			break
		}
	LABEL_59:
		v16 = v28
		if func() int32 {
			p := &v17
			*p++
			return *p
		}() > v28 {
			goto LABEL_60
		}
	}
	v18 = uint8(int8(v27))
	*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v27))), 0)) = uint8(int8(v27 & math.MaxInt8))
	v19 = int8(int32(v18) >> 7)
	if noxflags.HasGame(noxflags.GameFlag23) {
		v20 = (**uint8)(unsafe.Pointer(sub_504290(int8(v17+int32(*memmap.PtrUint8(6112660, 741360))), int8(v4+int32(*memmap.PtrUint8(6112660, 741368))))))
		v21 = *v20
		**v20 = uint8(int8(v27))
	LABEL_43:
		if int32(v19) != 0 {
			*(*uint8)(unsafe.Add(unsafe.Pointer(v21), 4)) |= 128
		}
		v24 = a1
		if int32(a1) < 2 {
			*(*uint8)(unsafe.Add(unsafe.Pointer(v21), 1)) = 0
			*(*uint8)(unsafe.Add(unsafe.Pointer(v21), 7)) = nox_xxx_mapWallGetHpByTile_410E20(0)
			*(*uint8)(unsafe.Add(unsafe.Pointer(v21), 8)) = 1
		} else {
			v25 = (*uint8)(unsafe.Add(unsafe.Pointer(v21), 1))
			cryptFileReadWrite((*uint8)(unsafe.Add(unsafe.Pointer(v21), 1))[:1])
			if int32(v24) >= 3 {
				cryptFileReadWrite((*uint8)(unsafe.Add(unsafe.Pointer(v21), 2))[:1])
			} else {
				sub_42A650(v21)
			}
			if v30&1 != 0 && int32(*(*uint8)(unsafe.Add(unsafe.Pointer(v21), 2))) >= int32(nox_xxx_mapWallMaxVariation_410DD0(int32(*v25), int32(*v21), 0)) {
				*(*uint8)(unsafe.Add(unsafe.Pointer(v21), 2)) = 0
			}
			*(*uint8)(unsafe.Add(unsafe.Pointer(v21), 7)) = nox_xxx_mapWallGetHpByTile_410E20(int32(*v25))
			if int32(v24) < 4 {
				*(*uint8)(unsafe.Add(unsafe.Pointer(v21), 8)) = 1
			} else {
				cryptFileReadWrite((*uint8)(unsafe.Add(unsafe.Pointer(v21), 8))[:1])
			}
			*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a2))), 0)) = 0
			if int32(v24) >= 5 {
				cryptFileReadWrite((*uint8)(unsafe.Pointer(&a2))[:1])
			}
			v4 = v31
			*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v21))), unsafe.Sizeof(uint32(0))*3))) = uint32(uint8(uintptr(unsafe.Pointer(a2))))
		}
		goto LABEL_59
	}
	v22 = (*uint8)(nox_server_getWallAtGrid_410580(int32(uint32(v17)+*memmap.PtrUint32(6112660, 741360)), int32(uint32(v4)+*memmap.PtrUint32(6112660, 741368))))
	v21 = v22
	if v22 != nil {
		if v30&1 != 0 {
			v23 = nox_xxx_wall_42A6C0(*v22, uint8(int8(v27)))
		LABEL_42:
			*v21 = v23
			goto LABEL_43
		}
	LABEL_41:
		v23 = uint8(int8(v27))
		goto LABEL_42
	}
	v21 = (*uint8)(nox_xxx_wallCreateAt_410250(int32(uint32(v17)+*memmap.PtrUint32(6112660, 741360)), int32(uint32(v4)+*memmap.PtrUint32(6112660, 741368))))
	if v21 != nil {
		goto LABEL_41
	}
	return 0
}
func sub_42A650(a1 *uint8) int32 {
	var (
		v1     uint8
		result int32
	)
	v1 = *a1
	*(*uint8)(unsafe.Add(unsafe.Pointer(a1), 2)) = 0
	if int32(v1) == 0 {
		*(*uint8)(unsafe.Add(unsafe.Pointer(a1), 2)) = uint8(int8(int32(*(*uint8)(unsafe.Add(unsafe.Pointer(a1), 5))) % 3))
	}
	if int32(v1) == 1 {
		*(*uint8)(unsafe.Add(unsafe.Pointer(a1), 2)) = uint8(int8(int32(*(*uint8)(unsafe.Add(unsafe.Pointer(a1), 5))) % 3))
	}
	result = nox_xxx_getWallSprite_46A3B0(int32(*(*uint8)(unsafe.Add(unsafe.Pointer(a1), 1))), int32(v1), int32(*(*uint8)(unsafe.Add(unsafe.Pointer(a1), 2))), (int32(*(*uint8)(unsafe.Add(unsafe.Pointer(a1), 4)))>>2)&2)
	if result == 0 {
		*(*uint8)(unsafe.Add(unsafe.Pointer(a1), 2)) = 0
	}
	return result
}
func nox_xxx_wall_42A6C0(a1 uint8, a2 uint8) uint8 {
	return *memmap.PtrUint8(0x587000, uint32(int32(a1)*13)+71276+uint32(a2))
}
func nox_server_mapRWMapInfo_42A6E0() int32 {
	var vers int32 = 3
	cryptFileReadWrite((*uint8)(unsafe.Pointer(&vers))[:2])
	if vers > 3 {
		return 0
	}
	if vers >= 1 {
		cryptFileReadWrite((*uint8)(memmap.PtrOff(0x973F18, 2408))[:64])
		cryptFileReadWrite((*uint8)(memmap.PtrOff(0x973F18, 2472))[:512])
		cryptFileReadWrite((*uint8)(memmap.PtrOff(0x973F18, 2984))[:16])
		cryptFileReadWrite((*uint8)(memmap.PtrOff(0x973F18, 3000))[:64])
		cryptFileReadWrite((*uint8)(memmap.PtrOff(0x973F18, 3064))[:64])
		cryptFileReadWrite((*uint8)(memmap.PtrOff(0x973F18, 3128))[:128])
		cryptFileReadWrite((*uint8)(memmap.PtrOff(0x973F18, 3256))[:128])
		cryptFileReadWrite((*uint8)(memmap.PtrOff(0x973F18, 3384))[:256])
		cryptFileReadWrite((*uint8)(memmap.PtrOff(0x973F18, 3640))[:128])
		cryptFileReadWrite((*uint8)(memmap.PtrOff(0x973F18, 3768))[:32])
		cryptFileReadWrite((*uint8)(memmap.PtrOff(0x973F18, 3800))[:4])
		if vers == 2 {
			cryptFileReadWrite((*uint8)(memmap.PtrOff(0x973F18, 3804))[:1])
			cryptFileReadWrite((*uint8)(memmap.PtrOff(0x973F18, 3805))[:1])
		} else if nox_xxx_cryptGetXxx() == 1 {
			*memmap.PtrUint8(0x973F18, 3804) = 2
			*memmap.PtrUint8(0x973F18, 3805) = 16
		}
	}
	if vers < 3 {
		*memmap.PtrUint8(0x973F18, 3806) = *memmap.PtrUint8(6112660, 741376)
		*memmap.PtrUint8(0x973F18, 3838) = *memmap.PtrUint8(6112660, 741380)
	} else {
		var v2 int32 = int32(libc.StrLen((*byte)(memmap.PtrOff(0x973F18, 3806))))
		cryptFileReadWrite((*uint8)(unsafe.Pointer(&v2))[:1])
		cryptFileReadWrite((*uint8)(memmap.PtrOff(0x973F18, 3806))[:uint32(v2)])
		*memmap.PtrUint8(0x973F18, uint32(v2+3806)) = 0
		v2 = int32(libc.StrLen((*byte)(memmap.PtrOff(0x973F18, 3838))))
		cryptFileReadWrite((*uint8)(unsafe.Pointer(&v2))[:1])
		cryptFileReadWrite((*uint8)(memmap.PtrOff(0x973F18, 3838))[:uint32(v2)])
		*memmap.PtrUint8(0x973F18, uint32(v2+3838)) = 0
	}
	return 1
}
func sub_42A8B0(a1 *uint8, a2 *int32) *uint16 {
	var (
		v2     int32
		v3     *uint8
		v4     *uint8
		result *uint16
		v6     int32
		v7     unsafe.Pointer
		v8     int32
		v9     *uint16
		v10    [12]uint8
	)
	*(*uint16)(unsafe.Pointer(&v10[0])) = 0
	*(*uint16)(unsafe.Pointer(&v10[2])) = 0
	v2 = *a2
	*(*uint32)(unsafe.Pointer(&v10[4])) = 0
	v3 = (*uint8)(alloc.Calloc(int(v2), 2))
	sub_42A970(a1, v3, a2)
	v4 = sub_42AC50(v3, (*uint32)(unsafe.Pointer(a2)))
	if v3 != nil {
		alloc.Free(unsafe.Pointer(v3))
	}
	if v4 != nil {
		v6 = *a2
		v7 = operator_new(16)
		if v7 != nil {
			v8 = sub_42C910(int32(uintptr(v7)), (*byte)(memmap.PtrOff(0x587000, 71480)), unsafe.Pointer(v4), uint16(int16(v6)))
		} else {
			v8 = 0
		}
		sub_42C360((*uint32)(unsafe.Pointer(&v10[0])), v8)
		alloc.Free(unsafe.Pointer(v4))
		v9 = sub_42C480((*uint32)(unsafe.Pointer(&v10[0])), (*uint32)(unsafe.Pointer(a2)))
		sub_42C330((*uint32)(unsafe.Pointer(&v10[0])))
		result = v9
	} else {
		sub_42C330((*uint32)(unsafe.Pointer(&v10[0])))
		result = nil
	}
	return result
}
func sub_42A970(a1 *uint8, a2 *uint8, a3 *int32) int32 {
	var (
		v3     int32
		v4     int8
		v5     int32
		v6     int32
		v7     int32
		v8     int32
		v9     *int32
		v10    *uint8
		v11    int32
		v12    int8
		result int32
		v14    int32
		v15    int32
		v16    int32
		v17    int8
		v18    uint8
		v19    [256]int32
	)
	v3 = 0
	v4 = 1
	*(*[256]int32)(unsafe.Pointer(&v19[0])) = [256]int32{}
	v18 = 1
	v5 = *a3
	if *a3 > 0 {
		for {
			v6 = int32(*(*uint8)(unsafe.Add(unsafe.Pointer(a1), func() int32 {
				p := &v3
				x := *p
				*p++
				return x
			}())))
			v19[v6]++
			if v3 >= v5 {
				break
			}
		}
	}
	v7 = v19[0]
	v17 = 0
	v8 = 0
	v9 = &v19[0]
	for *v9 != 0 {
		if *v9 < v7 {
			v7 = *v9
			v17 = int8(v8)
		}
		v8++
		v9 = (*int32)(unsafe.Add(unsafe.Pointer(v9), unsafe.Sizeof(int32(0))*1))
		if v8 >= 256 {
			goto LABEL_10
		}
	}
	v17 = int8(v8)
LABEL_10:
	v10 = a1
	v11 = 1
	*a2 = uint8(v17)
	v12 = int8(*a1)
	result = 1
	v14 = *a3
	if *a3 < 1 {
		*a3 = 1
	} else {
		for {
			if int32(*(*uint8)(unsafe.Add(unsafe.Pointer(v10), v11))) != int32(v12) || v11 >= v14 {
				if int32(uint8(v4)) <= 3 {
					if int32(v18) > 0 {
						v16 = int32(v18)
						for {
							if int32(v12) == int32(v17) {
								*(*uint8)(unsafe.Add(unsafe.Pointer(a2), func() int32 {
									p := &result
									x := *p
									*p++
									return x
								}())) = uint8(v12)
							}
							*(*uint8)(unsafe.Add(unsafe.Pointer(a2), func() int32 {
								p := &result
								x := *p
								*p++
								return x
							}())) = uint8(v12)
							v16--
							if v16 == 0 {
								break
							}
						}
					}
				} else {
					*(*uint8)(unsafe.Add(unsafe.Pointer(a2), result)) = uint8(v17)
					v15 = result + 1
					*(*uint8)(unsafe.Add(unsafe.Pointer(a2), func() int32 {
						p := &v15
						x := *p
						*p++
						return x
					}())) = uint8(v4)
					*(*uint8)(unsafe.Add(unsafe.Pointer(a2), v15)) = uint8(v12)
					result = v15 + 1
				}
				v10 = a1
				v4 = 1
				v12 = int8(*(*uint8)(unsafe.Add(unsafe.Pointer(a1), v11)))
			} else {
				v4++
			}
			v11++
			v18 = uint8(v4)
			v14 = *a3
			if v11 > *a3 {
				break
			}
		}
		*a3 = result
	}
	return result
}
func sub_42AAA0(a1 *int32) float64 {
	var (
		v1  int32
		v2  int32
		v3  int32
		v4  int32
		v5  int32
		v6  int32
		v7  *uint8
		v8  int32
		v9  int32
		v10 int32
		v11 int32
		v13 int32
	)
	if *a1 < 0 || *memmap.PtrUint32(6112660, 741656) == 0 {
		*memmap.PtrUint32(6112660, 741656) = 1
		v1 = *a1
		if *a1 < 0 {
			v1 = -v1
		}
		v2 = 1
		v3 = 21
		v4 = int32((0x9A4EC86 - uint32(v1)) % 1000000000)
		*memmap.PtrUint32(6112660, 741604) = (0x9A4EC86 - uint32(v1)) % 1000000000
		for {
			*memmap.PtrUint32(6112660, uint32((v3%55)*4)+0xB5008) = uint32(v2)
			v2 = int32(uint32(v4-v2) + (func() uint32 {
				if v4-v2 < 0 {
					return 1000000000
				}
				return 0
			}()))
			v4 = int32(*memmap.PtrUint32(6112660, uint32((v3%55)*4)+0xB5008))
			v3 += 21
			if v3 > 1134 {
				break
			}
		}
		v5 = 4
		for {
			v6 = 1
			v7 = (*uint8)(memmap.PtrOff(6112660, 741388))
			for {
				v8 = int32(*(*uint32)(unsafe.Pointer(v7)) - *memmap.PtrUint32(6112660, uint32(((v6+30)%55)*4)+0xB500C))
				*(*uint32)(unsafe.Pointer(v7)) = uint32(v8)
				if v8 < 0 {
					*(*uint32)(unsafe.Pointer(v7)) = uint32(v8) + 1000000000
				}
				v7 = (*uint8)(unsafe.Add(unsafe.Pointer(v7), 4))
				v6++
				if int32(uintptr(unsafe.Pointer(v7))) > int32(uintptr(memmap.PtrOff(6112660, 741604))) {
					break
				}
			}
			v5--
			if v5 == 0 {
				break
			}
		}
		dword_5d4594_741648 = 0
		dword_5d4594_741652 = 31
		*a1 = 1
	}
	v9 = int32(func() uint32 {
		p := &dword_5d4594_741648
		*p++
		return *p
	}())
	if dword_5d4594_741648 == 56 {
		v9 = 1
		dword_5d4594_741648 = 1
	}
	v10 = int32(func() uint32 {
		p := &dword_5d4594_741652
		*p++
		return *p
	}())
	if dword_5d4594_741652 == 56 {
		v10 = 1
		dword_5d4594_741652 = 1
	}
	v11 = int32(*memmap.PtrUint32(6112660, uint32(v9*4)+0xB5008) - *memmap.PtrUint32(6112660, uint32(v10*4)+0xB5008))
	v13 = v11
	if v11 < 0 {
		v11 += 1000000000
		v13 = v11
	}
	*memmap.PtrUint32(6112660, uint32(v9*4)+0xB5008) = uint32(v11)
	return float64(v13) * *mem_getDoublePtr(0x581450, 8368)
}
func sub_42ABF0(a1 int32, a2 int32, a3 int32) {
	var (
		v3 float64
		i  int32
	)
	if a3 > 0 {
		a3 = -a3
	}
	v3 = sub_42AAA0(&a3)
	for i = 0; i < a2; i++ {
		var v5 int32 = int32(v3 * 255.0)
		if v5 < 0 {
			v5 = -v5
		}
		*(*uint8)(unsafe.Pointer(uintptr(i + a1))) = uint8(int8(v5))
		v3 = sub_42AAA0(&a3)
	}
}
func sub_42AC50(a1 *uint8, a2 *uint32) *uint8 {
	var (
		v2     *uint32
		v3     unsafe.Pointer
		v4     uint32
		v5     *uint8
		result *uint8
		v7     libc.Time
		v8     int32
		v9     uint32
		v10    libc.Time
		v11    uint8
		v12    int32
		v13    int32
		v14    *uint8
		v15    int32
		v16    int8
		v17    int32
		v18    unsafe.Pointer
	)
	_ = v18
	var v19 libc.Time
	var lpMem unsafe.Pointer
	var v21 int32
	var v22 uint32
	var v23 float32
	var v24 uint8
	v2 = a2
	v19 = 0
	v3 = alloc.Calloc(1, int(*a2))
	v4 = *a2 + 5
	lpMem = v3
	v5 = (*uint8)(alloc.Calloc(1, int(v4)))
	if int32(*a2) >= 15 {
		v7 = libc.GetTime(nil)
		v8 = int32(v7)
		if v7 > 0 {
			v8 = int32(-v7)
		}
		v9 = cnet.Htonl(uint32(v8))
		v17 = int32(*a2)
		v22 = v9
		sub_42ABF0(int32(uintptr(v3)), v17, v8)
		v10 = libc.GetTime(nil)
		v19 = v10
		if v10 > 0 {
			v19 = -v10
		}
		if int32(*a2) >= 241 {
			v23 = 241.0
		} else {
			v23 = float32(float64(int32(*a2 - 14)))
		}
		v11 = uint8(uint64(int64(sub_42AAA0((*int32)(unsafe.Pointer(&v19)))*float64(v23))) + 10)
		v12 = 0
		v13 = 0
		v24 = v11
		if int32(*v2) > 0 {
			v21 = int32(v11)
			v14 = a1
			v15 = int32(uintptr(unsafe.Pointer((*uint8)(lpMem))) - uintptr(unsafe.Pointer(a1)))
			for {
				if v13 == 5 {
					*(*uint8)(unsafe.Add(unsafe.Pointer(v5), 5)) = v11
					v13 = 6
				}
				if v13 == v21 {
					*(*uint32)(unsafe.Add(unsafe.Pointer(v5), v13)) = v22
					v13 += 4
				}
				v16 = int8(int32(*v14) ^ int32(*(*uint8)(unsafe.Add(unsafe.Pointer(v14), v15))))
				v13++
				v12++
				v14 = (*uint8)(unsafe.Add(unsafe.Pointer(v14), 1))
				*(*uint8)(unsafe.Add(unsafe.Pointer(v5), v13-1)) = uint8(v16)
				if v12 >= int32(*v2) {
					break
				}
				v11 = v24
			}
		}
		v18 = lpMem
		*v2 += 5
		v18 = nil
		result = v5
	} else {
		*a2 = 0xFFFFFFFE
		v3 = nil
		alloc.Free(unsafe.Pointer(v5))
		result = nil
	}
	return result
}
func sub_42ADA0(a1 int32, a2 int32, a3 int16, a4 *uint32) *uint16 {
	var (
		v4   int32
		v5   int32
		v6   unsafe.Pointer
		v7   int32
		v8   int32
		v9   unsafe.Pointer
		v10  int32
		v11  int32
		v12  unsafe.Pointer
		v13  int32
		v14  int32
		v15  unsafe.Pointer
		v16  int32
		v17  unsafe.Pointer
		v18  int32
		v19  int32
		v20  unsafe.Pointer
		v21  int32
		v22  int32
		v23  unsafe.Pointer
		v24  int32
		v25  unsafe.Pointer
		v26  int32
		v27  int32
		v28  unsafe.Pointer
		v29  int32
		v30  int32
		v31  unsafe.Pointer
		v32  int32
		v33  int32
		v34  unsafe.Pointer
		v35  int32
		v36  int32
		v37  unsafe.Pointer
		v38  int32
		v39  unsafe.Pointer
		v40  int32
		v41  unsafe.Pointer
		v42  int32
		v43  unsafe.Pointer
		v44  int32
		v45  unsafe.Pointer
		v46  int32
		v47  unsafe.Pointer
		v48  int32
		v49  int32
		v50  unsafe.Pointer
		v51  int32
		v52  int32
		v53  unsafe.Pointer
		v54  int32
		v55  int32
		v56  unsafe.Pointer
		v57  int32
		v58  int32
		v59  unsafe.Pointer
		v60  int32
		v61  unsafe.Pointer
		v62  int32
		v63  unsafe.Pointer
		v64  int32
		v65  unsafe.Pointer
		v66  int32
		v67  int32
		v68  unsafe.Pointer
		v69  int32
		v70  int32
		v71  unsafe.Pointer
		v72  int32
		v73  unsafe.Pointer
		v74  int32
		v75  unsafe.Pointer
		v76  int32
		v77  unsafe.Pointer
		v78  int32
		v79  unsafe.Pointer
		v80  int32
		v81  unsafe.Pointer
		v82  int32
		v83  int32
		v84  int8
		v85  unsafe.Pointer
		v86  int32
		v87  unsafe.Pointer
		v88  int32
		v89  int16
		v90  unsafe.Pointer
		v91  int32
		v92  int32
		v93  *byte
		v94  unsafe.Pointer
		v95  int32
		v96  int32
		v97  unsafe.Pointer
		v98  int32
		v99  int32
		v100 unsafe.Pointer
		v101 int32
		v102 unsafe.Pointer
		v103 int32
		v104 unsafe.Pointer
		v105 int32
		v106 int32
		v107 unsafe.Pointer
		v108 int32
		v109 unsafe.Pointer
		v110 int32
		v111 unsafe.Pointer
		v112 unsafe.Pointer
		v113 int32
		v114 int32
		v115 *uint16
		v117 int8
		v118 [12]byte
		v119 int8
		v120 int8
		v121 int8
		v122 int8
		v123 int8
		v124 int8
		v125 int8
		v126 int8
		v127 int8
		v128 int8
		v129 int8
	)
	v4 = a1
	v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 8))))
	*(*uint16)(unsafe.Pointer(&v118[0])) = 0
	*(*uint16)(unsafe.Pointer(&v118[2])) = 0
	*(*uint32)(unsafe.Pointer(&v118[4])) = 0
	v6 = operator_new(16)
	if v6 != nil {
		v7 = sub_42C8B0(int32(uintptr(v6)), CString("MXPL"), int8(v5))
	} else {
		v7 = 0
	}
	sub_42C360((*uint32)(unsafe.Pointer(&v118[0])), v7)
	v8 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 12))))
	v9 = operator_new(16)
	if v9 != nil {
		v10 = sub_42C8B0(int32(uintptr(v9)), CString("IDNO"), int8(v8))
	} else {
		v10 = 0
	}
	sub_42C360((*uint32)(unsafe.Pointer(&v118[0])), v10)
	v11 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 16))))
	v12 = operator_new(16)
	if v12 != nil {
		v13 = sub_42C8B0(int32(uintptr(v12)), CString("GSKU"), int8(v11))
	} else {
		v13 = 0
	}
	sub_42C360((*uint32)(unsafe.Pointer(&v118[0])), v13)
	v14 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 20))))
	v15 = operator_new(16)
	if v15 != nil {
		v16 = sub_42C8B0(int32(uintptr(v15)), CString("GSTY"), int8(v14))
	} else {
		v16 = 0
	}
	sub_42C360((*uint32)(unsafe.Pointer(&v118[0])), v16)
	v119 = int8(*(*uint8)(unsafe.Pointer(uintptr(a1 + 24))))
	v17 = operator_new(16)
	if v17 != nil {
		v18 = sub_42C7F0(int32(uintptr(v17)), CString("CLGM"), v119)
	} else {
		v18 = 0
	}
	sub_42C360((*uint32)(unsafe.Pointer(&v118[0])), v18)
	v19 = int32(*(*uint32)(unsafe.Pointer(uintptr(v4 + 32))))
	v20 = operator_new(16)
	if v20 != nil {
		v21 = sub_42C8B0(int32(uintptr(v20)), CString("LIMT"), int8(v19))
	} else {
		v21 = 0
	}
	sub_42C360((*uint32)(unsafe.Pointer(&v118[0])), v21)
	v22 = int32(*(*uint32)(unsafe.Pointer(uintptr(v4 + 36))))
	v23 = operator_new(16)
	if v23 != nil {
		v24 = sub_42C8B0(int32(uintptr(v23)), CString("TLMT"), int8(v22))
	} else {
		v24 = 0
	}
	sub_42C360((*uint32)(unsafe.Pointer(&v118[0])), v24)
	v120 = int8(*(*uint8)(unsafe.Pointer(uintptr(v4 + 40))))
	v25 = operator_new(16)
	if v25 != nil {
		v26 = sub_42C7F0(int32(uintptr(v25)), CString("RSTC"), v120)
	} else {
		v26 = 0
	}
	sub_42C360((*uint32)(unsafe.Pointer(&v118[0])), v26)
	v27 = int32(*(*uint32)(unsafe.Pointer(uintptr(v4 + 44))))
	v28 = operator_new(16)
	if v28 != nil {
		v29 = sub_42C8B0(int32(uintptr(v28)), CString("MINE"), int8(v27))
	} else {
		v29 = 0
	}
	sub_42C360((*uint32)(unsafe.Pointer(&v118[0])), v29)
	v30 = int32(*(*uint32)(unsafe.Pointer(uintptr(v4 + 48))))
	v31 = operator_new(16)
	if v31 != nil {
		v32 = sub_42C8B0(int32(uintptr(v31)), CString("MAXE"), int8(v30))
	} else {
		v32 = 0
	}
	sub_42C360((*uint32)(unsafe.Pointer(&v118[0])), v32)
	v33 = int32(*(*uint32)(unsafe.Pointer(uintptr(v4 + 52))))
	v34 = operator_new(16)
	if v34 != nil {
		v35 = sub_42C8B0(int32(uintptr(v34)), CString("MINP"), int8(v33))
	} else {
		v35 = 0
	}
	sub_42C360((*uint32)(unsafe.Pointer(&v118[0])), v35)
	v36 = int32(*(*uint32)(unsafe.Pointer(uintptr(v4 + 56))))
	v37 = operator_new(16)
	if v37 != nil {
		v38 = sub_42C8B0(int32(uintptr(v37)), CString("MAXP"), int8(v36))
	} else {
		v38 = 0
	}
	sub_42C360((*uint32)(unsafe.Pointer(&v118[0])), v38)
	v121 = int8(*(*uint8)(unsafe.Pointer(uintptr(v4 + 60))))
	v39 = operator_new(16)
	if v39 != nil {
		v40 = sub_42C7F0(int32(uintptr(v39)), CString("VIDM"), v121)
	} else {
		v40 = 0
	}
	sub_42C360((*uint32)(unsafe.Pointer(&v118[0])), v40)
	v122 = int8(*(*uint8)(unsafe.Pointer(uintptr(v4 + 61))))
	v41 = operator_new(16)
	if v41 != nil {
		v42 = sub_42C7F0(int32(uintptr(v41)), CString("SVRS"), v122)
	} else {
		v42 = 0
	}
	sub_42C360((*uint32)(unsafe.Pointer(&v118[0])), v42)
	v123 = int8(*(*uint8)(unsafe.Pointer(uintptr(v4 + 25))))
	v43 = operator_new(16)
	if v43 != nil {
		v44 = sub_42C7F0(int32(uintptr(v43)), CString("NTMS"), v123)
	} else {
		v44 = 0
	}
	sub_42C360((*uint32)(unsafe.Pointer(&v118[0])), v44)
	v45 = operator_new(16)
	if v45 != nil {
		v46 = sub_42C8E0(int32(uintptr(v45)), CString("SCEN"), (*byte)(unsafe.Pointer(uintptr(v4+96))))
	} else {
		v46 = 0
	}
	sub_42C360((*uint32)(unsafe.Pointer(&v118[0])), v46)
	v47 = operator_new(16)
	if v47 != nil {
		v48 = sub_42C8E0(int32(uintptr(v47)), CString("GNAM"), (*byte)(unsafe.Pointer(uintptr(v4+352))))
	} else {
		v48 = 0
	}
	sub_42C360((*uint32)(unsafe.Pointer(&v118[0])), v48)
	v49 = int32(*(*uint32)(unsafe.Pointer(uintptr(v4 + 64))))
	v50 = operator_new(16)
	if v50 != nil {
		v51 = sub_42C8B0(int32(uintptr(v50)), CString("SPL1"), int8(v49))
	} else {
		v51 = 0
	}
	sub_42C360((*uint32)(unsafe.Pointer(&v118[0])), v51)
	v52 = int32(*(*uint32)(unsafe.Pointer(uintptr(v4 + 68))))
	v53 = operator_new(16)
	if v53 != nil {
		v54 = sub_42C8B0(int32(uintptr(v53)), CString("SPL2"), int8(v52))
	} else {
		v54 = 0
	}
	sub_42C360((*uint32)(unsafe.Pointer(&v118[0])), v54)
	v55 = int32(*(*uint32)(unsafe.Pointer(uintptr(v4 + 72))))
	v56 = operator_new(16)
	if v56 != nil {
		v57 = sub_42C8B0(int32(uintptr(v56)), CString("SPL3"), int8(v55))
	} else {
		v57 = 0
	}
	sub_42C360((*uint32)(unsafe.Pointer(&v118[0])), v57)
	v58 = int32(*(*uint32)(unsafe.Pointer(uintptr(v4 + 88))))
	v59 = operator_new(16)
	if v59 != nil {
		v60 = sub_42C8B0(int32(uintptr(v59)), CString("ARMR"), int8(v58))
	} else {
		v60 = 0
	}
	sub_42C360((*uint32)(unsafe.Pointer(&v118[0])), v60)
	v124 = int8(*(*uint8)(unsafe.Pointer(uintptr(v4 + 84))))
	v61 = operator_new(16)
	if v61 != nil {
		v62 = sub_42C7F0(int32(uintptr(v61)), CString("WPN1"), v124)
	} else {
		v62 = 0
	}
	sub_42C360((*uint32)(unsafe.Pointer(&v118[0])), v62)
	v125 = int8(*(*uint8)(unsafe.Pointer(uintptr(v4 + 85))))
	v63 = operator_new(16)
	if v63 != nil {
		v64 = sub_42C7F0(int32(uintptr(v63)), CString("WPN2"), v125)
	} else {
		v64 = 0
	}
	sub_42C360((*uint32)(unsafe.Pointer(&v118[0])), v64)
	v126 = int8(*(*uint8)(unsafe.Pointer(uintptr(v4 + 86))))
	v65 = operator_new(16)
	if v65 != nil {
		v66 = sub_42C7F0(int32(uintptr(v65)), CString("WPN3"), v126)
	} else {
		v66 = 0
	}
	sub_42C360((*uint32)(unsafe.Pointer(&v118[0])), v66)
	v67 = int32(*(*uint32)(unsafe.Pointer(uintptr(v4 + 92))))
	v68 = operator_new(16)
	if v68 != nil {
		v69 = sub_42C8B0(int32(uintptr(v68)), CString("STAF"), int8(v67))
	} else {
		v69 = 0
	}
	sub_42C360((*uint32)(unsafe.Pointer(&v118[0])), v69)
	v70 = int32(*(*uint32)(unsafe.Pointer(uintptr(v4 + 28))))
	v71 = operator_new(16)
	if v71 != nil {
		v72 = sub_42C8B0(int32(uintptr(v71)), CString("DURA"), int8(v70))
	} else {
		v72 = 0
	}
	sub_42C360((*uint32)(unsafe.Pointer(&v118[0])), v72)
	v73 = operator_new(16)
	if v73 != nil {
		v74 = sub_42C7F0(int32(uintptr(v73)), CString("FINI"), 1)
	} else {
		v74 = 0
	}
	sub_42C360((*uint32)(unsafe.Pointer(&v118[0])), v74)
	v127 = int8(*(*uint8)(unsafe.Pointer(uintptr(v4 + 26))))
	v75 = operator_new(16)
	if v75 != nil {
		v76 = sub_42C7F0(int32(uintptr(v75)), CString("TRNY"), v127)
	} else {
		v76 = 0
	}
	sub_42C360((*uint32)(unsafe.Pointer(&v118[0])), v76)
	switch a2 {
	case 0:
		*memmap.PtrUint32(6112660, 741668) = 0
		v77 = operator_new(16)
		if v77 != nil {
			v78 = sub_42C8B0(int32(uintptr(v77)), CString("SEQU"), 0)
		} else {
			v78 = 0
		}
		sub_42C360((*uint32)(unsafe.Pointer(&v118[0])), v78)
		v79 = operator_new(16)
		if v79 != nil {
			v80 = sub_42C7F0(int32(uintptr(v79)), CString("ENDF"), 0)
		} else {
			v80 = 0
		}
		sub_42C360((*uint32)(unsafe.Pointer(&v118[0])), v80)
		v81 = operator_new(16)
		if v81 != nil {
			v82 = sub_42C820(int32(uintptr(v81)), (*byte)(memmap.PtrOff(0x587000, 71872)), -1)
			sub_42C360((*uint32)(unsafe.Pointer(&v118[0])), v82)
			break
		}
		sub_42C360((*uint32)(unsafe.Pointer(&v118[0])), 0)
	case 2:
		v83 = int32(*memmap.PtrUint32(6112660, 741668) + 1)
		*memmap.PtrUint32(6112660, 741668) = uint32(v83)
		v84 = int8(v83)
		v85 = operator_new(16)
		if v85 != nil {
			v86 = sub_42C8B0(int32(uintptr(v85)), CString("SEQU"), v84)
		} else {
			v86 = 0
		}
		sub_42C360((*uint32)(unsafe.Pointer(&v118[0])), v86)
		v87 = operator_new(16)
		if v87 != nil {
			v88 = sub_42C7F0(int32(uintptr(v87)), CString("ENDF"), 1)
		} else {
			v88 = 0
		}
		sub_42C360((*uint32)(unsafe.Pointer(&v118[0])), v88)
		v89 = int16(*(*uint16)(unsafe.Pointer(uintptr(v4 + 6))))
		v90 = operator_new(16)
		if v90 != nil {
			v91 = sub_42C820(int32(uintptr(v90)), (*byte)(memmap.PtrOff(0x587000, 71896)), int8(v89))
		} else {
			v91 = 0
		}
		sub_42C360((*uint32)(unsafe.Pointer(&v118[0])), v91)
		v92 = 0
		for *memmap.PtrUint32(6112660, 741660) = 0; v92 < int32(*(*int16)(unsafe.Pointer(uintptr(v4 + 6)))); *memmap.PtrUint32(6112660, 741660) = uint32(v92) {
			*memmap.PtrUint8(0x587000, 71491) = uint8(int8(v92 + 48))
			v93 = *(**byte)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v4 + 608))) + uint32(v92*4))))
			v94 = operator_new(16)
			if v94 != nil {
				v95 = sub_42C8E0(int32(uintptr(v94)), CString("LGL?"), v93)
			} else {
				v95 = 0
			}
			sub_42C360((*uint32)(unsafe.Pointer(&v118[0])), v95)
			*memmap.PtrUint8(0x587000, 71499) = uint8(int8(int32(*memmap.PtrUint8(6112660, 741660)) + 48))
			v96 = int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v4 + 612))) + *memmap.PtrUint32(6112660, 741660)*4))))
			v97 = operator_new(16)
			if v97 != nil {
				v98 = sub_42C8B0(int32(uintptr(v97)), CString("IPL?"), int8(v96))
			} else {
				v98 = 0
			}
			sub_42C360((*uint32)(unsafe.Pointer(&v118[0])), v98)
			*memmap.PtrUint8(0x587000, 71515) = uint8(int8(int32(*memmap.PtrUint8(6112660, 741660)) + 48))
			v99 = int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v4 + 616))) + *memmap.PtrUint32(6112660, 741660)*4))))
			v100 = operator_new(16)
			if v100 != nil {
				v101 = sub_42C8B0(int32(uintptr(v100)), CString("CNL?"), int8(v99))
			} else {
				v101 = 0
			}
			sub_42C360((*uint32)(unsafe.Pointer(&v118[0])), v101)
			*memmap.PtrUint8(0x587000, 71507) = uint8(int8(int32(*memmap.PtrUint8(6112660, 741660)) + 48))
			v128 = int8(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, 741660) + *(*uint32)(unsafe.Pointer(uintptr(v4 + 620)))))))
			v102 = operator_new(16)
			if v102 != nil {
				v103 = sub_42C7F0(int32(uintptr(v102)), CString("CLL?"), v128)
			} else {
				v103 = 0
			}
			sub_42C360((*uint32)(unsafe.Pointer(&v118[0])), v103)
			*memmap.PtrUint8(0x587000, 71523) = uint8(int8(int32(*memmap.PtrUint8(6112660, 741660)) + 48))
			v129 = int8(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v4 + 624))) + *memmap.PtrUint32(6112660, 741660)))))
			v104 = operator_new(16)
			if v104 != nil {
				v105 = sub_42C7F0(int32(uintptr(v104)), CString("CMP?"), v129)
			} else {
				v105 = 0
			}
			sub_42C360((*uint32)(unsafe.Pointer(&v118[0])), v105)
			*memmap.PtrUint8(0x587000, 71531) = uint8(int8(int32(*memmap.PtrUint8(6112660, 741660)) + 48))
			v106 = int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v4 + 628))) + *memmap.PtrUint32(6112660, 741660)*4))))
			v107 = operator_new(16)
			if v107 != nil {
				v108 = sub_42C8B0(int32(uintptr(v107)), CString("DUR?"), int8(v106))
			} else {
				v108 = 0
			}
			sub_42C360((*uint32)(unsafe.Pointer(&v118[0])), v108)
			*memmap.PtrUint8(0x587000, 71539) = uint8(int8(int32(*memmap.PtrUint8(6112660, 741660)) + 48))
			v117 = int8(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v4 + 632))) + *memmap.PtrUint32(6112660, 741660)))))
			v109 = operator_new(16)
			if v109 != nil {
				v110 = sub_42C7F0(int32(uintptr(v109)), CString("PAR?"), v117)
			} else {
				v110 = 0
			}
			sub_42C360((*uint32)(unsafe.Pointer(&v118[0])), v110)
			v92 = int32(*memmap.PtrUint32(6112660, 741660) + 1)
		}
		v111 = *(*unsafe.Pointer)(unsafe.Pointer(uintptr(v4 + 636)))
		v112 = operator_new(16)
		if v112 != nil {
			v113 = sub_42C910(int32(uintptr(v112)), (*byte)(memmap.PtrOff(0x587000, 71904)), v111, uint16(int16(int32(a3)*2)))
			sub_42C360((*uint32)(unsafe.Pointer(&v118[0])), v113)
			break
		}
		sub_42C360((*uint32)(unsafe.Pointer(&v118[0])), 0)
	case 1:
		*memmap.PtrUint32(6112660, 741668)++
		sub_42BDC0((*uint32)(unsafe.Pointer(&v118[0])), CString("SEQU"), int8(*memmap.PtrUint8(6112660, 741668)))
		sub_42BCE0((*uint32)(unsafe.Pointer(&v118[0])), CString("ENDF"), 0)
		sub_42BD50((*uint32)(unsafe.Pointer(&v118[0])), (*byte)(memmap.PtrOff(0x587000, 71928)), int8(uint8(*(*uint16)(unsafe.Pointer(uintptr(v4 + 6))))))
		v114 = 0
		for *memmap.PtrUint32(6112660, 741660) = 0; v114 < int32(*(*int16)(unsafe.Pointer(uintptr(v4 + 6)))); *memmap.PtrUint32(6112660, 741660) = uint32(v114) {
			*memmap.PtrUint8(0x587000, 71491) = uint8(int8(v114 + 48))
			sub_42BE30((*uint32)(unsafe.Pointer(&v118[0])), (*byte)(memmap.PtrOff(0x587000, 71488)), *(**byte)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v4 + 608))) + uint32(v114*4)))))
			*memmap.PtrUint8(0x587000, 71499) = uint8(int8(int32(*memmap.PtrUint8(6112660, 741660)) + 48))
			sub_42BDC0((*uint32)(unsafe.Pointer(&v118[0])), CString("IPL?"), int8(uint8(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v4 + 612))) + *memmap.PtrUint32(6112660, 741660)*4))))))
			*memmap.PtrUint8(0x587000, 71515) = uint8(int8(int32(*memmap.PtrUint8(6112660, 741660)) + 48))
			sub_42BDC0((*uint32)(unsafe.Pointer(&v118[0])), CString("CNL?"), int8(uint8(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v4 + 616))) + *memmap.PtrUint32(6112660, 741660)*4))))))
			*memmap.PtrUint8(0x587000, 71507) = uint8(int8(int32(*memmap.PtrUint8(6112660, 741660)) + 48))
			sub_42BCE0((*uint32)(unsafe.Pointer(&v118[0])), CString("CLL?"), int8(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, 741660) + *(*uint32)(unsafe.Pointer(uintptr(v4 + 620))))))))
			*memmap.PtrUint8(0x587000, 71523) = uint8(int8(int32(*memmap.PtrUint8(6112660, 741660)) + 48))
			sub_42BCE0((*uint32)(unsafe.Pointer(&v118[0])), CString("CMP?"), int8(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v4 + 624))) + *memmap.PtrUint32(6112660, 741660))))))
			*memmap.PtrUint8(0x587000, 71531) = uint8(int8(int32(*memmap.PtrUint8(6112660, 741660)) + 48))
			sub_42BDC0((*uint32)(unsafe.Pointer(&v118[0])), CString("DUR?"), int8(uint8(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v4 + 628))) + *memmap.PtrUint32(6112660, 741660)*4))))))
			*memmap.PtrUint8(0x587000, 71539) = uint8(int8(int32(*memmap.PtrUint8(6112660, 741660)) + 48))
			sub_42BCE0((*uint32)(unsafe.Pointer(&v118[0])), CString("PAR?"), int8(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v4 + 632))) + *memmap.PtrUint32(6112660, 741660))))))
			v114 = int32(*memmap.PtrUint32(6112660, 741660) + 1)
		}
		sub_42BEA0((*uint32)(unsafe.Pointer(&v118[0])), (*byte)(memmap.PtrOff(0x587000, 71936)), *(*unsafe.Pointer)(unsafe.Pointer(uintptr(v4 + 636))), uint16(int16(int32(a3)*2)))
	}
	v115 = sub_42C480((*uint32)(unsafe.Pointer(&v118[0])), a4)
	sub_42C330((*uint32)(unsafe.Pointer(&v118[0])))
	return v115
}
func sub_42B810(a1 *int16, a2 *uint32) *uint16 {
	var (
		v2  *int16
		v3  int32
		v4  unsafe.Pointer
		v5  int32
		v6  int32
		v7  unsafe.Pointer
		v8  int32
		v9  int32
		v10 unsafe.Pointer
		v11 int32
		v12 unsafe.Pointer
		v13 int32
		v14 unsafe.Pointer
		v15 int32
		v16 int32
		v17 unsafe.Pointer
		v18 int32
		v19 unsafe.Pointer
		v20 int32
		v21 int16
		v22 unsafe.Pointer
		v23 int32
		v24 int8
		v25 unsafe.Pointer
		v26 int32
		v27 int32
		v28 *byte
		v29 unsafe.Pointer
		v30 int32
		v31 int32
		v32 unsafe.Pointer
		v33 int32
		v34 unsafe.Pointer
		v35 int32
		v36 int32
		v37 unsafe.Pointer
		v38 int32
		v39 int32
		v40 unsafe.Pointer
		v41 int32
		v42 int32
		v43 unsafe.Pointer
		v44 int32
		v45 int32
		v46 unsafe.Pointer
		v47 int32
		v48 int32
		v49 unsafe.Pointer
		v50 int32
		v51 int32
		v52 unsafe.Pointer
		v53 int32
		v54 int32
		v55 unsafe.Pointer
		v56 int32
		v57 *uint16
		v59 [12]byte
		v60 int8
		v61 int8
	)
	v2 = a1
	v3 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a1))), unsafe.Sizeof(uint32(0))*1))))
	*(*uint16)(unsafe.Pointer(&v59[0])) = 0
	*(*uint16)(unsafe.Pointer(&v59[2])) = 0
	*(*uint32)(unsafe.Pointer(&v59[4])) = 0
	v4 = operator_new(16)
	if v4 != nil {
		v5 = sub_42C8B0(int32(uintptr(v4)), CString("IDNO"), int8(v3))
	} else {
		v5 = 0
	}
	sub_42C360((*uint32)(unsafe.Pointer(&v59[0])), v5)
	v6 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a1))), unsafe.Sizeof(uint32(0))*2))))
	v7 = operator_new(16)
	if v7 != nil {
		v8 = sub_42C8B0(int32(uintptr(v7)), CString("GSKU"), int8(v6))
	} else {
		v8 = 0
	}
	sub_42C360((*uint32)(unsafe.Pointer(&v59[0])), v8)
	v9 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a1))), unsafe.Sizeof(uint32(0))*3))))
	v10 = operator_new(16)
	if v10 != nil {
		v11 = sub_42C8B0(int32(uintptr(v10)), CString("GSTY"), int8(v9))
	} else {
		v11 = 0
	}
	sub_42C360((*uint32)(unsafe.Pointer(&v59[0])), v11)
	v12 = operator_new(16)
	if v12 != nil {
		v13 = sub_42C8E0(int32(uintptr(v12)), CString("SCEN"), (*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(a1))), 24)))
	} else {
		v13 = 0
	}
	sub_42C360((*uint32)(unsafe.Pointer(&v59[0])), v13)
	v14 = operator_new(16)
	if v14 != nil {
		v15 = sub_42C8E0(int32(uintptr(v14)), CString("GNAM"), (*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(a1))), 280)))
	} else {
		v15 = 0
	}
	sub_42C360((*uint32)(unsafe.Pointer(&v59[0])), v15)
	v16 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a1))), unsafe.Sizeof(uint32(0))*5))))
	v17 = operator_new(16)
	if v17 != nil {
		v18 = sub_42C8B0(int32(uintptr(v17)), CString("DURA"), int8(v16))
	} else {
		v18 = 0
	}
	sub_42C360((*uint32)(unsafe.Pointer(&v59[0])), v18)
	v60 = int8(*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(a1))), 16))))
	v19 = operator_new(16)
	if v19 != nil {
		v20 = sub_42C7F0(int32(uintptr(v19)), CString("TRNY"), v60)
	} else {
		v20 = 0
	}
	sub_42C360((*uint32)(unsafe.Pointer(&v59[0])), v20)
	v21 = *v2
	v22 = operator_new(16)
	if v22 != nil {
		v23 = sub_42C820(int32(uintptr(v22)), (*byte)(memmap.PtrOff(0x587000, 72000)), int8(v21))
	} else {
		v23 = 0
	}
	sub_42C360((*uint32)(unsafe.Pointer(&v59[0])), v23)
	v24 = int8(*memmap.PtrUint8(6112660, 741672))
	v25 = operator_new(16)
	if v25 != nil {
		v26 = sub_42C8B0(int32(uintptr(v25)), CString("SEQU"), v24)
	} else {
		v26 = 0
	}
	sub_42C360((*uint32)(unsafe.Pointer(&v59[0])), v26)
	v27 = 0
	for *memmap.PtrUint32(6112660, 741664) = 0; v27 < int32(*v2); *memmap.PtrUint32(6112660, 741664) = uint32(v27) {
		*memmap.PtrUint8(0x587000, 71547) = uint8(int8(v27 + 48))
		v28 = *(**byte)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), unsafe.Sizeof(uint32(0))*134))) + uint32(v27*4))))
		v29 = operator_new(16)
		if v29 != nil {
			v30 = sub_42C8E0(int32(uintptr(v29)), CString("LGLS"), v28)
		} else {
			v30 = 0
		}
		sub_42C360((*uint32)(unsafe.Pointer(&v59[0])), v30)
		*memmap.PtrUint8(0x587000, 71555) = uint8(int8(int32(*memmap.PtrUint8(6112660, 741664)) + 48))
		v31 = int32(*(*uint32)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), unsafe.Sizeof(uint32(0))*135))) + *memmap.PtrUint32(6112660, 741664)*4))))
		v32 = operator_new(16)
		if v32 != nil {
			v33 = sub_42C8B0(int32(uintptr(v32)), CString("IPLS"), int8(v31))
		} else {
			v33 = 0
		}
		sub_42C360((*uint32)(unsafe.Pointer(&v59[0])), v33)
		*memmap.PtrUint8(0x587000, 71563) = uint8(int8(int32(*memmap.PtrUint8(6112660, 741664)) + 48))
		v61 = int8(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, 741664) + *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), unsafe.Sizeof(uint32(0))*144)))))))
		v34 = operator_new(16)
		if v34 != nil {
			v35 = sub_42C7F0(int32(uintptr(v34)), CString("CLLS"), v61)
		} else {
			v35 = 0
		}
		sub_42C360((*uint32)(unsafe.Pointer(&v59[0])), v35)
		*memmap.PtrUint8(0x587000, 71571) = uint8(int8(int32(*memmap.PtrUint8(6112660, 741664)) + 48))
		v36 = int32(*(*uint32)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), unsafe.Sizeof(uint32(0))*136))) + *memmap.PtrUint32(6112660, 741664)*4))))
		v37 = operator_new(16)
		if v37 != nil {
			v38 = sub_42C8B0(int32(uintptr(v37)), CString("CSTS"), int8(v36))
		} else {
			v38 = 0
		}
		sub_42C360((*uint32)(unsafe.Pointer(&v59[0])), v38)
		*memmap.PtrUint8(0x587000, 71579) = uint8(int8(int32(*memmap.PtrUint8(6112660, 741664)) + 48))
		v39 = int32(*(*uint32)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), unsafe.Sizeof(uint32(0))*137))) + *memmap.PtrUint32(6112660, 741664)*4))))
		v40 = operator_new(16)
		if v40 != nil {
			v41 = sub_42C8B0(int32(uintptr(v40)), CString("HSTS"), int8(v39))
		} else {
			v41 = 0
		}
		sub_42C360((*uint32)(unsafe.Pointer(&v59[0])), v41)
		*memmap.PtrUint8(0x587000, 71587) = uint8(int8(int32(*memmap.PtrUint8(6112660, 741664)) + 48))
		v42 = int32(*(*uint32)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), unsafe.Sizeof(uint32(0))*138))) + *memmap.PtrUint32(6112660, 741664)*4))))
		v43 = operator_new(16)
		if v43 != nil {
			v44 = sub_42C8B0(int32(uintptr(v43)), CString("MKLS"), int8(v42))
		} else {
			v44 = 0
		}
		sub_42C360((*uint32)(unsafe.Pointer(&v59[0])), v44)
		*memmap.PtrUint8(0x587000, 71595) = uint8(int8(int32(*memmap.PtrUint8(6112660, 741664)) + 48))
		v45 = int32(*(*uint32)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), unsafe.Sizeof(uint32(0))*139))) + *memmap.PtrUint32(6112660, 741664)*4))))
		v46 = operator_new(16)
		if v46 != nil {
			v47 = sub_42C8B0(int32(uintptr(v46)), CString("ANKS"), int8(v45))
		} else {
			v47 = 0
		}
		sub_42C360((*uint32)(unsafe.Pointer(&v59[0])), v47)
		*memmap.PtrUint8(0x587000, 71603) = uint8(int8(int32(*memmap.PtrUint8(6112660, 741664)) + 48))
		v48 = int32(*(*uint32)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), unsafe.Sizeof(uint32(0))*140))) + *memmap.PtrUint32(6112660, 741664)*4))))
		v49 = operator_new(16)
		if v49 != nil {
			v50 = sub_42C8B0(int32(uintptr(v49)), CString("GNDS"), int8(v48))
		} else {
			v50 = 0
		}
		sub_42C360((*uint32)(unsafe.Pointer(&v59[0])), v50)
		*memmap.PtrUint8(0x587000, 71611) = uint8(int8(int32(*memmap.PtrUint8(6112660, 741664)) + 48))
		v51 = int32(*(*uint32)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), unsafe.Sizeof(uint32(0))*141))) + *memmap.PtrUint32(6112660, 741664)*4))))
		v52 = operator_new(16)
		if v52 != nil {
			v53 = sub_42C8B0(int32(uintptr(v52)), CString("SECS"), int8(v51))
		} else {
			v53 = 0
		}
		sub_42C360((*uint32)(unsafe.Pointer(&v59[0])), v53)
		*memmap.PtrUint8(0x587000, 71619) = uint8(int8(int32(*memmap.PtrUint8(6112660, 741664)) + 48))
		v54 = int32(*(*uint32)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), unsafe.Sizeof(uint32(0))*142))) + *memmap.PtrUint32(6112660, 741664)*4))))
		v55 = operator_new(16)
		if v55 != nil {
			v56 = sub_42C8B0(int32(uintptr(v55)), CString("BPTS"), int8(v54))
		} else {
			v56 = 0
		}
		sub_42C360((*uint32)(unsafe.Pointer(&v59[0])), v56)
		*memmap.PtrUint8(0x587000, 71627) = uint8(int8(int32(*memmap.PtrUint8(6112660, 741664)) + 48))
		sub_42BDC0((*uint32)(unsafe.Pointer(&v59[0])), CString("SCRS"), int8(uint8(*(*uint32)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), unsafe.Sizeof(uint32(0))*143))) + *memmap.PtrUint32(6112660, 741664)*4))))))
		v27 = int32(*memmap.PtrUint32(6112660, 741664) + 1)
	}
	v57 = sub_42C480((*uint32)(unsafe.Pointer(&v59[0])), a2)
	*memmap.PtrUint32(6112660, 741672)++
	sub_42C330((*uint32)(unsafe.Pointer(&v59[0])))
	return v57
}
func sub_42BCE0(this *uint32, a2 *byte, a3 int8) int32 {
	var (
		v3 *uint32
		v4 unsafe.Pointer
		v5 int32
	)
	v3 = this
	v4 = operator_new(16)
	if v4 != nil {
		v5 = sub_42C7F0(int32(uintptr(v4)), a2, a3)
	} else {
		v5 = 0
	}
	return sub_42C360(v3, v5)
}
func sub_42BD50(this *uint32, a2 *byte, a3 int8) int32 {
	var (
		v3 *uint32
		v4 unsafe.Pointer
		v5 int32
	)
	v3 = this
	v4 = operator_new(16)
	if v4 != nil {
		v5 = sub_42C820(int32(uintptr(v4)), a2, a3)
	} else {
		v5 = 0
	}
	return sub_42C360(v3, v5)
}
func sub_42BDC0(this *uint32, a2 *byte, a3 int8) int32 {
	var (
		v3 *uint32
		v4 unsafe.Pointer
		v5 int32
	)
	v3 = this
	v4 = operator_new(16)
	if v4 != nil {
		v5 = sub_42C8B0(int32(uintptr(v4)), a2, a3)
	} else {
		v5 = 0
	}
	return sub_42C360(v3, v5)
}
func sub_42BE30(this *uint32, a2 *byte, a3 *byte) int32 {
	var (
		v3 *uint32
		v4 unsafe.Pointer
		v5 int32
	)
	v3 = this
	v4 = operator_new(16)
	if v4 != nil {
		v5 = sub_42C8E0(int32(uintptr(v4)), a2, a3)
	} else {
		v5 = 0
	}
	return sub_42C360(v3, v5)
}
func sub_42BEA0(this *uint32, a2 *byte, a3 unsafe.Pointer, a4 uint16) int32 {
	var (
		v4 *uint32
		v5 unsafe.Pointer
		v6 int32
	)
	v4 = this
	v5 = operator_new(16)
	if v5 != nil {
		v6 = sub_42C910(int32(uintptr(v5)), a2, a3, a4)
	} else {
		v6 = 0
	}
	return sub_42C360(v4, v6)
}
func sub_42BFB0() int32 {
	var result int32
	result = 0
	libc.MemSet(*(*unsafe.Pointer)(unsafe.Pointer(&dword_5d4594_741676)), 0, int(dword_5d4594_741680*2))
	*memmap.PtrUint16(6112660, 741684) = 0
	return result
}
func sub_42BFE0() int32 {
	var (
		result int32
		i      int32
		v2     uint16
		v3     bool
		v4     *uint16
	)
	_ = v4
	var j int32
	var v6 uint16
	var v7 *uint16
	_ = v7
	var k int32
	var v9 *uint16
	var v11 uint16
	var v12 *uint16
	_ = v12
	var m int32
	var v14 *uint16
	var v15 int32
	result = bool2int(noxflags.HasGame(noxflags.GameHost | noxflags.GameFlag22))
	if result != 0 {
		sub_42BFB0()
		v15 = 1
		for i = int32(uintptr(unsafe.Pointer(noxServer.firstServerObject()))); i != 0; i = int32(uintptr(unsafe.Pointer((*nox_object_t)(unsafe.Pointer(uintptr(i))).Next()))) {
			v2 = *(*uint16)(unsafe.Pointer(uintptr(i + 4)))
			v3 = int32(*(*uint16)(unsafe.Pointer(uintptr(dword_5d4594_741676 + uint32(int32(v2)*2))))) == 0
			v4 = (*uint16)(unsafe.Pointer(uintptr(dword_5d4594_741676 + uint32(int32(v2)*2))))
			if v3 {
				*v4 = uint16(int16(v15))
				*memmap.PtrUint16(6112660, 741684)++
				v15++
			}
			for j = int32(*(*uint32)(unsafe.Pointer(uintptr(i + 504)))); j != 0; j = int32(*(*uint32)(unsafe.Pointer(uintptr(j + 496)))) {
				v6 = *(*uint16)(unsafe.Pointer(uintptr(j + 4)))
				v3 = int32(*(*uint16)(unsafe.Pointer(uintptr(dword_5d4594_741676 + uint32(int32(v6)*2))))) == 0
				v7 = (*uint16)(unsafe.Pointer(uintptr(dword_5d4594_741676 + uint32(int32(v6)*2))))
				if v3 {
					*v7 = uint16(int16(v15))
					*memmap.PtrUint16(6112660, 741684)++
					v15++
				}
			}
			if *(*uint32)(unsafe.Pointer(uintptr(i + 8)))&0x20000 != 0 {
				sub_42C210(i, (*uint16)(unsafe.Pointer(&v15)), (*uint16)(memmap.PtrOff(6112660, 741684)))
			}
		}
		for k = int32(uintptr(unsafe.Pointer(noxServer.firstServerObjectUninited()))); k != 0; k = int32(uintptr(unsafe.Pointer(nox_server_getNextObjectUninited_4DA880((*nox_object_t)(unsafe.Pointer(uintptr(k))))))) {
			v9 = (*uint16)(unsafe.Pointer(uintptr(dword_5d4594_741676 + uint32(int32(*(*uint16)(unsafe.Pointer(uintptr(k + 4))))*2))))
			if int32(*v9) == 0 {
				*v9 = uint16(int16(v15))
				*memmap.PtrUint16(6112660, 741684)++
				v15++
			}
			if *(*uint32)(unsafe.Pointer(uintptr(k + 8)))&0x20000 != 0 {
				sub_42C210(k, (*uint16)(unsafe.Pointer(&v15)), (*uint16)(memmap.PtrOff(6112660, 741684)))
			}
		}
		for obj := (*nox_object_t)(firstServerObjectUpdatable2()); obj != nil; obj = obj.Next() {
			v11 = obj.typ_ind
			v3 = int32(*(*uint16)(unsafe.Pointer(uintptr(dword_5d4594_741676 + uint32(int32(v11)*2))))) == 0
			v12 = (*uint16)(unsafe.Pointer(uintptr(dword_5d4594_741676 + uint32(int32(v11)*2))))
			if v3 {
				*v12 = uint16(int16(v15))
				*memmap.PtrUint16(6112660, 741684)++
				v15++
			}
		}
		result = bool2int(noxflags.HasGame(noxflags.GameFlag22))
		if result == 0 {
			result = bool2int(noxflags.HasGame(noxflags.GameHost))
			if result == 1 {
				result = bool2int(noxflags.HasGame(noxflags.GameClient))
				if result == 1 {
					result = bool2int(noxflags.HasGame(noxflags.GameFlag23))
					if result == 0 {
						result = sub_45A060()
						for m = result; result != 0; m = result {
							if sub_4E3AD0(int32(*(*uint32)(unsafe.Pointer(uintptr(m + 108))))) == 0 && sub_4E3B80(int32(*(*uint32)(unsafe.Pointer(uintptr(m + 108))))) != 0 {
								v14 = (*uint16)(unsafe.Pointer(uintptr(dword_5d4594_741676 + *(*uint32)(unsafe.Pointer(uintptr(m + 108)))*2)))
								if int32(*v14) == 0 {
									*v14 = uint16(int16(v15))
									*memmap.PtrUint16(6112660, 741684)++
									v15++
								}
							}
							result = sub_45A070(m)
						}
					}
				}
			}
		}
	}
	return result
}
func sub_42C210(a1 int32, a2 *uint16, a3 *uint16) int32 {
	var (
		result int32
		v4     int32
		v5     int32
		v6     *uint16
		i      int32
		v8     uint16
		v9     bool
		v10    *uint16
	)
	_ = v10
	var v11 int32
	result = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
	v11 = 3
	for {
		v4 = result
		v5 = 4
		for {
			if *(*uint32)(unsafe.Pointer(uintptr(v4))) != 0 {
				v6 = (*uint16)(unsafe.Pointer(uintptr(dword_5d4594_741676 + uint32(int32(*(*uint16)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v4))) + 4))))*2))))
				if int32(*v6) == 0 {
					*v6 = func() uint16 {
						p := a2
						x := *p
						*p++
						return x
					}()
					*a3++
				}
				for i = int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v4))) + 504)))); i != 0; i = int32(*(*uint32)(unsafe.Pointer(uintptr(i + 496)))) {
					v8 = *(*uint16)(unsafe.Pointer(uintptr(i + 4)))
					v9 = int32(*(*uint16)(unsafe.Pointer(uintptr(dword_5d4594_741676 + uint32(int32(v8)*2))))) == 0
					v10 = (*uint16)(unsafe.Pointer(uintptr(dword_5d4594_741676 + uint32(int32(v8)*2))))
					if v9 {
						*v10 = func() uint16 {
							p := a2
							x := *p
							*p++
							return x
						}()
						*a3++
					}
				}
			}
			v4 += 4
			v5--
			if v5 == 0 {
				break
			}
		}
		result = v4
		v11--
		if v11 == 0 {
			break
		}
	}
	return result
}
func nox_xxx_objectTOCgetTT_42C2B0(a1 int16) int16 {
	var (
		v1 *uint16
		v2 uint32
	)
	v1 = *(**uint16)(unsafe.Pointer(&dword_5d4594_741676))
	if dword_5d4594_741676 != 0 && (func() uint32 {
		v2 = 0
		return dword_5d4594_741680
	}()) != 0 {
		for int32(*v1) != int32(a1) {
			v2++
			v1 = (*uint16)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint16(0))*1))
			if v2 >= uint32(*(*int32)(unsafe.Pointer(&dword_5d4594_741680))) {
				goto LABEL_5
			}
		}
	} else {
	LABEL_5:
		*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v2))), unsafe.Sizeof(uint16(0))*0)) = 0
	}
	return int16(uint16(v2))
}
func sub_42C2E0(a1 int32) int16 {
	var result int16
	if dword_5d4594_741676 != 0 {
		result = int16(*(*uint16)(unsafe.Pointer(uintptr(dword_5d4594_741676 + uint32(a1*2)))))
	} else {
		result = 0
	}
	return result
}
func sub_42C300() int16 {
	return int16(*memmap.PtrUint16(6112660, 741684))
}
func sub_42C310(a1 uint16, a2 int16) unsafe.Pointer {
	var result unsafe.Pointer
	result = *(*unsafe.Pointer)(unsafe.Pointer(&dword_5d4594_741676))
	if dword_5d4594_741676 != 0 {
		*(*uint16)(unsafe.Pointer(uintptr(dword_5d4594_741676 + uint32(int32(a1)*2)))) = uint16(a2)
	}
	return result
}
func sub_42C330(this *uint32) {
	var (
		v1 *uint32
		v2 *uint32
	)
	v1 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(this), unsafe.Sizeof(uint32(0))*1)))))
	if v1 != nil {
		for {
			v2 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*3)))))
			if v1 != nil {
				sub_42CC50((*unsafe.Pointer)(unsafe.Pointer(v1)))
				operator_delete(unsafe.Pointer(v1))
			}
			v1 = v2
			if v2 == nil {
				break
			}
		}
	}
}
func sub_42C360(this *uint32, a2 int32) int32 {
	var result int32
	result = a2
	*(*uint32)(unsafe.Pointer(uintptr(a2 + 12))) = *(*uint32)(unsafe.Add(unsafe.Pointer(this), unsafe.Sizeof(uint32(0))*1))
	*(*uint32)(unsafe.Add(unsafe.Pointer(this), unsafe.Sizeof(uint32(0))*1)) = uint32(a2)
	return result
}
func sub_42C370(this *uint16, a2 *uint16) *uint16 {
	var (
		v2  *uint16
		v3  uint16
		v4  uint16
		v5  *uint16
		v6  uint16
		v7  unsafe.Pointer
		v8  int32
		v9  int32
		v10 *uint32
		v11 uint16
		v12 uint32
		v13 unsafe.Pointer
		v14 int32
		v16 int32
		v17 *uint16
		v18 int32
	)
	v2 = this
	v3 = *a2
	v17 = this
	*this = *a2
	*this = cnet.Ntohs(v3)
	v4 = *(*uint16)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint16(0))*1))
	*(*uint16)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint16(0))*1)) = *(*uint16)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint16(0))*1))
	v5 = (*uint16)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint16(0))*2))
	*(*uint16)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint16(0))*1)) = cnet.Ntohs(v4)
	v6 = *v2
	*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), unsafe.Sizeof(uint32(0))*1))) = 0
	v16 = int32(v6) - 4
	if v16 <= 0 {
		return v2
	}
	for {
		v7 = operator_new(16)
		v8 = int32(*(*uint32)(unsafe.Pointer(v5)))
		v9 = int32(uintptr(v7))
		v10 = (*uint32)(unsafe.Pointer((*uint16)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(uint16(0))*4))))
		*(*uint32)(v7) = uint32(v8)
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(v7)), unsafe.Sizeof(uint32(0))*1))) = *((*uint32)(unsafe.Add(unsafe.Pointer(v10), -int(unsafe.Sizeof(uint32(0))*1))))
		v11 = cnet.Ntohs(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(v7)), unsafe.Sizeof(uint16(0))*3))))
		v12 = uint32(v11)
		v13 = operator_new(uint32(v11))
		*(*uint32)(unsafe.Pointer(uintptr(v9 + 8))) = uint32(uintptr(v13))
		libc.MemCpy(v13, unsafe.Pointer(v10), int(v12))
		v14 = -(int32(cnet.Ntohs(*(*uint16)(unsafe.Pointer(uintptr(v9 + 6))))) & 3) & 3
		v18 = int32(uint32(int32(uintptr(unsafe.Pointer(v10)))+v14) + v12)
		v16 += int32(uint32(int32(-8-v14)) - v12)
		sub_42CCE0((*uint16)(unsafe.Pointer(uintptr(v9))))
		sub_42C360((*uint32)(unsafe.Pointer(v17)), v9)
		if v16 <= 0 {
			break
		}
		v5 = (*uint16)(unsafe.Pointer(uintptr(v18)))
	}
	return v17
}
func sub_42C480(this *uint32, a2 *uint32) *uint16 {
	var (
		v2 *uint32
		i  int32
		v4 uint32
		v5 *uint16
		v6 int32
		v7 *uint32
		v8 *byte
		v9 int32
	)
	v2 = this
	*a2 = 4
	for i = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(this), unsafe.Sizeof(uint32(0))*1))); i != 0; i = int32(*(*uint32)(unsafe.Pointer(uintptr(i + 12)))) {
		v4 = *a2 + 8
		*a2 = v4
		*a2 = uint32(*(*uint16)(unsafe.Pointer(uintptr(i + 6)))) + v4 + uint32(-((int32(uint8(*(*uint16)(unsafe.Pointer(uintptr(i + 6)))))+int32(uint8(v4)))&3)&3)
	}
	v5 = (*uint16)(operator_new(*a2))
	*v5 = cnet.Htons(*(*uint16)(unsafe.Pointer(a2)))
	*(*uint16)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(uint16(0))*1)) = cnet.Htons(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v2))), unsafe.Sizeof(uint16(0))*1))))
	v6 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*1)))
	v7 = (*uint32)(unsafe.Pointer((*uint16)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(uint16(0))*2))))
	if v6 == 0 {
		return v5
	}
	for {
		sub_42CC70(v6)
		*v7 = *(*uint32)(unsafe.Pointer(uintptr(v6)))
		*(*uint32)(unsafe.Add(unsafe.Pointer(v7), unsafe.Sizeof(uint32(0))*1)) = *(*uint32)(unsafe.Pointer(uintptr(v6 + 4)))
		v8 = (*byte)(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer(v7), unsafe.Sizeof(uint32(0))*2))))
		libc.MemCpy(unsafe.Pointer(v8), *(*unsafe.Pointer)(unsafe.Pointer(uintptr(v6 + 8))), int(cnet.Ntohs(*(*uint16)(unsafe.Pointer(uintptr(v6 + 6))))))
		v9 = int32(uintptr(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v8), cnet.Ntohs(*(*uint16)(unsafe.Pointer(uintptr(v6 + 6)))))))))
		v7 = (*uint32)(unsafe.Pointer(uintptr((-(int32(cnet.Ntohs(*(*uint16)(unsafe.Pointer(uintptr(v6 + 6))))) & 3) & 3) + v9)))
		sub_42CCE0((*uint16)(unsafe.Pointer(uintptr(v6))))
		v6 = int32(*(*uint32)(unsafe.Pointer(uintptr(v6 + 12))))
		if v6 == 0 {
			break
		}
	}
	return v5
}
func sub_42C580(this *uint32, a2 *byte) int32 {
	var v2 int32
	v2 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(this), unsafe.Sizeof(uint32(0))*1)))
	if v2 == 0 {
		return 0
	}
	for libc.StrNCmp(a2, (*byte)(unsafe.Pointer(uintptr(v2))), 4) != 0 {
		v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 12))))
		if v2 == 0 {
			return 0
		}
	}
	return v2
}
func sub_42C5C0(this *uint32, a2 *byte, a3 *uint8) bool {
	var v3 int32
	v3 = sub_42C580(this, a2)
	if v3 != 0 {
		*a3 = **(**uint8)(unsafe.Pointer(uintptr(v3 + 8)))
	}
	return v3 != 0
}
func sub_42C5E0(this *uint32, a2 *byte, a3 *uint8) bool {
	var v3 int32
	v3 = sub_42C580(this, a2)
	if v3 != 0 {
		*a3 = **(**uint8)(unsafe.Pointer(uintptr(v3 + 8)))
	}
	return v3 != 0
}
func sub_42C600(this *uint32, a2 *byte, a3 *uint16) bool {
	var v3 int32
	v3 = sub_42C580(this, a2)
	if v3 != 0 {
		*a3 = **(**uint16)(unsafe.Pointer(uintptr(v3 + 8)))
	}
	return v3 != 0
}
func sub_42C630(this *uint32, a2 *byte, a3 *uint16) bool {
	var v3 int32
	v3 = sub_42C580(this, a2)
	if v3 != 0 {
		*a3 = **(**uint16)(unsafe.Pointer(uintptr(v3 + 8)))
	}
	return v3 != 0
}
func sub_42C660(this *uint32, a2 *byte, a3 *uint32) bool {
	var v3 int32
	v3 = sub_42C580(this, a2)
	if v3 != 0 {
		*a3 = **(**uint32)(unsafe.Pointer(uintptr(v3 + 8)))
	}
	return v3 != 0
}
func sub_42C680(this *uint32, a2 *byte, a3 *uint32) bool {
	var v3 int32
	v3 = sub_42C580(this, a2)
	if v3 != 0 {
		*a3 = **(**uint32)(unsafe.Pointer(uintptr(v3 + 8)))
	}
	return v3 != 0
}
func sub_42C6A0(this *uint32, a2 *byte, a3 *byte) bool {
	var (
		v3 int32
		v4 bool
	)
	v3 = sub_42C580(this, a2)
	v4 = v3 == 0
	if v3 != 0 {
		libc.StrCpy(a3, *(**byte)(unsafe.Pointer(uintptr(v3 + 8))))
		v4 = v3 == 0
	}
	return !v4
}
func sub_42C6E0(this *uint32, a2 *byte, a3 *uint32) bool {
	var v3 int32
	v3 = sub_42C580(this, a2)
	if v3 != 0 {
		*a3 = **(**uint32)(unsafe.Pointer(uintptr(v3 + 8)))
	}
	return v3 != 0
}
func sub_42C700(this *uint32, a2 *byte, a3 *uint32) bool {
	var v3 int32
	v3 = sub_42C580(this, a2)
	if v3 != 0 {
		*a3 = **(**uint32)(unsafe.Pointer(uintptr(v3 + 8)))
	}
	return v3 != 0
}
func sub_42C720(this *uint32, a2 *byte, a3 unsafe.Pointer, a4 *uint32) bool {
	var (
		v4 int32
		v5 bool
		v6 uint32
	)
	v4 = sub_42C580(this, a2)
	v5 = v4 == 0
	if v4 != 0 {
		v6 = uint32(*(*uint16)(unsafe.Pointer(uintptr(v4 + 6))))
		if int32(uint16(v6)) >= int32(*a4) {
			v6 = *a4
		}
		libc.MemCpy(a3, *(*unsafe.Pointer)(unsafe.Pointer(uintptr(v4 + 8))), int(v6))
		*a4 = uint32(*(*uint16)(unsafe.Pointer(uintptr(v4 + 6))))
		v5 = v4 == 0
	}
	return !v5
}
func sub_42C770(this *unsafe.Pointer) int32 {
	var (
		v1     *unsafe.Pointer
		result int32
	)
	v1 = this
	operator_delete(*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(this), unsafe.Sizeof(unsafe.Pointer(nil))*2)))
	result = 0
	libc.StrCpy((*byte)(unsafe.Pointer(v1)), (*byte)(memmap.PtrOff(6112660, 741688)))
	*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v1))), unsafe.Sizeof(uint16(0))*2))) = 0
	*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v1))), unsafe.Sizeof(uint16(0))*3))) = 0
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(unsafe.Pointer(nil))*2)) = nil
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(unsafe.Pointer(nil))*3)) = nil
	return result
}
func sub_42C7C0(this int32, a2 *byte, a3 int8) int32 {
	var v3 int32
	v3 = this
	*(*uint32)(unsafe.Pointer(uintptr(this + 8))) = 0
	sub_42C940((*unsafe.Pointer)(unsafe.Pointer(uintptr(this))), a2, a3)
	return v3
}
func sub_42C7F0(this int32, a2 *byte, a3 int8) int32 {
	var v3 int32
	v3 = this
	*(*uint32)(unsafe.Pointer(uintptr(this + 8))) = 0
	sub_42C9A0((*unsafe.Pointer)(unsafe.Pointer(uintptr(this))), a2, a3)
	return v3
}
func sub_42C820(this int32, a2 *byte, a3 int8) int32 {
	var v3 int32
	v3 = this
	*(*uint32)(unsafe.Pointer(uintptr(this + 8))) = 0
	sub_42CA00((*unsafe.Pointer)(unsafe.Pointer(uintptr(this))), a2, a3)
	return v3
}
func sub_42C850(this int32, a2 *byte, a3 int8) int32 {
	var v3 int32
	v3 = this
	*(*uint32)(unsafe.Pointer(uintptr(this + 8))) = 0
	sub_42CA60((*unsafe.Pointer)(unsafe.Pointer(uintptr(this))), a2, a3)
	return v3
}
func sub_42C880(this int32, a2 *byte, a3 int8) int32 {
	var v3 int32
	v3 = this
	*(*uint32)(unsafe.Pointer(uintptr(this + 8))) = 0
	sub_42CAC0((*unsafe.Pointer)(unsafe.Pointer(uintptr(this))), a2, a3)
	return v3
}
func sub_42C8B0(this int32, a2 *byte, a3 int8) int32 {
	var v3 int32
	v3 = this
	*(*uint32)(unsafe.Pointer(uintptr(this + 8))) = 0
	sub_42CB20((*unsafe.Pointer)(unsafe.Pointer(uintptr(this))), a2, a3)
	return v3
}
func sub_42C8E0(this int32, a2 *byte, a3 *byte) int32 {
	var v3 int32
	v3 = this
	*(*uint32)(unsafe.Pointer(uintptr(this + 8))) = 0
	sub_42CB80((*unsafe.Pointer)(unsafe.Pointer(uintptr(this))), a2, a3)
	return v3
}
func sub_42C910(this int32, a2 *byte, a3 unsafe.Pointer, a4 uint16) int32 {
	var v4 int32
	v4 = this
	*(*uint32)(unsafe.Pointer(uintptr(this + 8))) = 0
	sub_42CBF0((*unsafe.Pointer)(unsafe.Pointer(uintptr(this))), a2, a3, a4)
	return v4
}
func sub_42C940(this *unsafe.Pointer, a2 *byte, a3 int8) *byte {
	var (
		v3     *unsafe.Pointer
		result *byte
		v5     uint32
	)
	v3 = this
	sub_42C770(this)
	libc.StrNCpy((*byte)(unsafe.Pointer(v3)), a2, 4)
	*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v3))), unsafe.Sizeof(uint16(0))*2))) = 1
	*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v3))), unsafe.Sizeof(uint16(0))*3))) = 1
	result = (*byte)(operator_new(1))
	v5 = uint32(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v3))), unsafe.Sizeof(uint16(0))*3))))
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(unsafe.Pointer(nil))*2)) = unsafe.Pointer(result)
	libc.MemCpy(unsafe.Pointer(result), unsafe.Pointer(&a3), int(v5))
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(unsafe.Pointer(nil))*3)) = nil
	return result
}
func sub_42C9A0(this *unsafe.Pointer, a2 *byte, a3 int8) *byte {
	var (
		v3     *unsafe.Pointer
		result *byte
		v5     uint32
	)
	v3 = this
	sub_42C770(this)
	libc.StrNCpy((*byte)(unsafe.Pointer(v3)), a2, 4)
	*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v3))), unsafe.Sizeof(uint16(0))*2))) = 2
	*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v3))), unsafe.Sizeof(uint16(0))*3))) = 1
	result = (*byte)(operator_new(1))
	v5 = uint32(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v3))), unsafe.Sizeof(uint16(0))*3))))
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(unsafe.Pointer(nil))*2)) = unsafe.Pointer(result)
	libc.MemCpy(unsafe.Pointer(result), unsafe.Pointer(&a3), int(v5))
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(unsafe.Pointer(nil))*3)) = nil
	return result
}
func sub_42CA00(this *unsafe.Pointer, a2 *byte, a3 int8) *byte {
	var (
		v3     *unsafe.Pointer
		result *byte
		v5     uint32
	)
	v3 = this
	sub_42C770(this)
	libc.StrNCpy((*byte)(unsafe.Pointer(v3)), a2, 4)
	*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v3))), unsafe.Sizeof(uint16(0))*2))) = 3
	*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v3))), unsafe.Sizeof(uint16(0))*3))) = 2
	result = (*byte)(operator_new(2))
	v5 = uint32(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v3))), unsafe.Sizeof(uint16(0))*3))))
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(unsafe.Pointer(nil))*2)) = unsafe.Pointer(result)
	var a3i int32 = int32(a3)
	libc.MemCpy(unsafe.Pointer(result), unsafe.Pointer(&a3i), int(v5))
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(unsafe.Pointer(nil))*3)) = nil
	return result
}
func sub_42CA60(this *unsafe.Pointer, a2 *byte, a3 int8) *byte {
	var (
		v3     *unsafe.Pointer
		result *byte
		v5     uint32
	)
	v3 = this
	sub_42C770(this)
	libc.StrNCpy((*byte)(unsafe.Pointer(v3)), a2, 4)
	*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v3))), unsafe.Sizeof(uint16(0))*2))) = 4
	*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v3))), unsafe.Sizeof(uint16(0))*3))) = 2
	result = (*byte)(operator_new(2))
	v5 = uint32(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v3))), unsafe.Sizeof(uint16(0))*3))))
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(unsafe.Pointer(nil))*2)) = unsafe.Pointer(result)
	var a3i int32 = int32(a3)
	libc.MemCpy(unsafe.Pointer(result), unsafe.Pointer(&a3i), int(v5))
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(unsafe.Pointer(nil))*3)) = nil
	return result
}
func sub_42CAC0(this *unsafe.Pointer, a2 *byte, a3 int8) *byte {
	var (
		v3     *unsafe.Pointer
		result *byte
		v5     uint32
	)
	v3 = this
	sub_42C770(this)
	libc.StrNCpy((*byte)(unsafe.Pointer(v3)), a2, 4)
	*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v3))), unsafe.Sizeof(uint16(0))*2))) = 5
	*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v3))), unsafe.Sizeof(uint16(0))*3))) = 4
	result = (*byte)(operator_new(4))
	v5 = uint32(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v3))), unsafe.Sizeof(uint16(0))*3))))
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(unsafe.Pointer(nil))*2)) = unsafe.Pointer(result)
	var a3i int32 = int32(a3)
	libc.MemCpy(unsafe.Pointer(result), unsafe.Pointer(&a3i), int(v5))
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(unsafe.Pointer(nil))*3)) = nil
	return result
}
func sub_42CB20(this *unsafe.Pointer, a2 *byte, a3 int8) *byte {
	var (
		v3     *unsafe.Pointer
		result *byte
		v5     uint32
	)
	v3 = this
	sub_42C770(this)
	libc.StrNCpy((*byte)(unsafe.Pointer(v3)), a2, 4)
	*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v3))), unsafe.Sizeof(uint16(0))*2))) = 6
	*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v3))), unsafe.Sizeof(uint16(0))*3))) = 4
	result = (*byte)(operator_new(4))
	v5 = uint32(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v3))), unsafe.Sizeof(uint16(0))*3))))
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(unsafe.Pointer(nil))*2)) = unsafe.Pointer(result)
	var a3i int32 = int32(a3)
	libc.MemCpy(unsafe.Pointer(result), unsafe.Pointer(&a3i), int(v5))
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(unsafe.Pointer(nil))*3)) = nil
	return result
}
func sub_42CB80(this *unsafe.Pointer, a2 *byte, a3 *byte) *byte {
	var (
		v3     *unsafe.Pointer
		v4     uint32
		result *byte
		v6     uint16
	)
	v3 = this
	sub_42C770(this)
	libc.StrNCpy((*byte)(unsafe.Pointer(v3)), a2, 4)
	*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v3))), unsafe.Sizeof(uint16(0))*2))) = 7
	v4 = uint32(libc.StrLen(a3) + 1)
	*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v3))), unsafe.Sizeof(uint16(0))*3))) = uint16(v4)
	result = (*byte)(operator_new(uint32(uint16(v4))))
	v6 = *((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v3))), unsafe.Sizeof(uint16(0))*3)))
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(unsafe.Pointer(nil))*2)) = unsafe.Pointer(result)
	libc.MemCpy(unsafe.Pointer(result), unsafe.Pointer(a3), int(v6))
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(unsafe.Pointer(nil))*3)) = nil
	return result
}
func sub_42CBF0(this *unsafe.Pointer, a2 *byte, a3 unsafe.Pointer, a4 uint16) unsafe.Pointer {
	var (
		v4     *unsafe.Pointer
		result unsafe.Pointer
		v6     uint32
	)
	v4 = this
	sub_42C770(this)
	libc.StrNCpy((*byte)(unsafe.Pointer(v4)), a2, 4)
	*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v4))), unsafe.Sizeof(uint16(0))*2))) = 20
	*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v4))), unsafe.Sizeof(uint16(0))*3))) = a4
	result = operator_new(uint32(a4))
	v6 = uint32(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v4))), unsafe.Sizeof(uint16(0))*3))))
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(unsafe.Pointer(nil))*2)) = result
	libc.MemCpy(result, a3, int(v6))
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(unsafe.Pointer(nil))*3)) = nil
	return result
}
func sub_42CC60(this *uint16) int32 {
	return int32(*(*uint16)(unsafe.Add(unsafe.Pointer(this), unsafe.Sizeof(uint16(0))*2)))
}
func sub_42CC70(this int32) uint16 {
	var (
		v1     int32
		v2     uint16
		v3     uint16
		result uint16
	)
	v1 = this
	switch *(*uint16)(unsafe.Pointer(uintptr(this + 4))) {
	case 3:
		fallthrough
	case 4:
		**(**uint16)(unsafe.Pointer(uintptr(this + 8))) = cnet.Htons(**(**uint16)(unsafe.Pointer(uintptr(this + 8))))
	case 5:
		fallthrough
	case 6:
		**(**uint32)(unsafe.Pointer(uintptr(this + 8))) = cnet.Htonl(**(**uint32)(unsafe.Pointer(uintptr(this + 8))))
	default:
	}
	v2 = cnet.Htons(*(*uint16)(unsafe.Pointer(uintptr(v1 + 4))))
	v3 = *(*uint16)(unsafe.Pointer(uintptr(v1 + 6)))
	*(*uint16)(unsafe.Pointer(uintptr(v1 + 4))) = v2
	result = cnet.Htons(v3)
	*(*uint16)(unsafe.Pointer(uintptr(v1 + 6))) = result
	return result
}
func sub_42CCE0(this *uint16) int16 {
	var (
		v1 *uint16
		v2 uint16
		v3 uint16
		v4 uint32
	)
	v1 = this
	v2 = cnet.Ntohs(*(*uint16)(unsafe.Add(unsafe.Pointer(this), unsafe.Sizeof(uint16(0))*2)))
	v3 = *(*uint16)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint16(0))*3))
	*(*uint16)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint16(0))*2)) = v2
	*(*uint16)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint16(0))*3)) = cnet.Ntohs(v3)
	*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v4))), unsafe.Sizeof(uint16(0))*0)) = uint16(int16(int32(*(*uint16)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint16(0))*2))) - 3))
	switch *(*uint16)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint16(0))*2)) {
	case 3:
		fallthrough
	case 4:
		*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v4))), unsafe.Sizeof(uint16(0))*0)) = cnet.Ntohs(**((**uint16)(unsafe.Add(unsafe.Pointer((**uint16)(unsafe.Pointer(v1))), unsafe.Sizeof((*uint16)(nil))*2))))
		**((**uint16)(unsafe.Add(unsafe.Pointer((**uint16)(unsafe.Pointer(v1))), unsafe.Sizeof((*uint16)(nil))*2))) = uint16(v4)
	case 5:
		fallthrough
	case 6:
		v4 = cnet.Ntohl(**((**uint32)(unsafe.Add(unsafe.Pointer((**uint32)(unsafe.Pointer(v1))), unsafe.Sizeof((*uint32)(nil))*2))))
		**((**uint32)(unsafe.Add(unsafe.Pointer((**uint32)(unsafe.Pointer(v1))), unsafe.Sizeof((*uint32)(nil))*2))) = v4
	default:
		return int16(uint16(v4))
	}
	return int16(uint16(v4))
}
func nox_xxx_clientTalk_42E7B0(a1p *nox_drawable) {
	var (
		a1 int32 = int32(uintptr(unsafe.Pointer(a1p)))
		v1 int32
		v2 int16
	)
	v1 = a1
	if a1 != 0 && (*memmap.PtrUint32(0x8531A0, 2576) == 0 || (int32(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(0x8531A0, 2576) + 3680))))&3) == 0) && sub_478030() != 1 && nox_gui_xxx_check_446360() != 1 {
		v2 = int16(*(*uint16)(unsafe.Pointer(uintptr(v1 + 128))))
		*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&a1))), unsafe.Sizeof(uint16(0))*0)) = 464
		*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&a1))), unsafe.Sizeof(uint16(0))*1)) = uint16(v2)
		nox_netlist_addToMsgListCli_40EBC0(31, 0, (*uint8)(unsafe.Pointer(&a1)), 4)
	}
}
func nox_xxx_clientCollideOrUse_42E810(a1p *nox_drawable) {
	var (
		a1 int32 = int32(uintptr(unsafe.Pointer(a1p)))
		v1 int32
	)
	if a1 != 0 && (*memmap.PtrUint32(0x8531A0, 2576) == 0 || (int32(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(0x8531A0, 2576) + 3680))))&3) == 0) {
		v1 = a1
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a1))), 0)) = 123
		*(*uint16)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(&a1))), 1)))) = uint16(nox_xxx_netGetUnitCodeCli_578B00(v1))
		nox_netlist_addToMsgListCli_40EBC0(31, 0, (*uint8)(unsafe.Pointer(&a1)), 3)
	}
}
func nox_xxx_clientTrade_42E850(a1p *nox_drawable) {
	var (
		a1 int32 = int32(uintptr(unsafe.Pointer(a1p)))
		v1 int32
	)
	v1 = a1
	if a1 != 0 && (*memmap.PtrUint32(0x8531A0, 2576) == 0 || (int32(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(0x8531A0, 2576) + 3680))))&3) == 0) && sub_47A260() != 1 && nox_gui_xxx_check_446360() != 1 {
		*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&a1))), unsafe.Sizeof(uint16(0))*0)) = 5577
		*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&a1))), unsafe.Sizeof(uint16(0))*1)) = uint16(nox_xxx_netGetUnitCodeCli_578B00(v1))
		nox_netlist_addToMsgListCli_40EBC0(31, 0, (*uint8)(unsafe.Pointer(&a1)), 4)
	}
}
func sub_42EB90(a1 int32) int32 {
	var result int32
	result = a1
	*memmap.PtrUint32(6112660, 754052) = uint32(a1)
	return result
}
func sub_42EBA0() int32 {
	return int32(*memmap.PtrUint32(6112660, 754052))
}
func sub_42ED20() {
	if obj_5D4594_754104_switch == 0 {
		return
	}
	obj_5D4594_754104_switch = 0
	if ptr_5D4594_754088 == nil {
		return
	}
	for i := int32(0); i < ptr_5D4594_754088_cnt; i++ {
		var cur *obj_5D4594_754088_t = (*obj_5D4594_754088_t)(unsafe.Add(unsafe.Pointer(ptr_5D4594_754088), unsafe.Sizeof(obj_5D4594_754088_t{})*uintptr(i)))
		cur.fnc(cur.field_4)
	}
}
func sub_42ED70() {
	if obj_5D4594_754104_switch != 0 {
		return
	}
	obj_5D4594_754104_switch = 1
	if ptr_5D4594_754092 == nil {
		return
	}
	for i := int32(0); i < ptr_5D4594_754092_cnt; i++ {
		var cur *obj_5D4594_754088_t = (*obj_5D4594_754088_t)(unsafe.Add(unsafe.Pointer(ptr_5D4594_754092), unsafe.Sizeof(obj_5D4594_754088_t{})*uintptr(i)))
		cur.fnc(cur.field_4)
	}
}
func sub_42EDC0() {
	if ptr_5D4594_754088 != nil {
		alloc.Free(unsafe.Pointer(ptr_5D4594_754088))
		ptr_5D4594_754088 = nil
	}
	if ptr_5D4594_754092 != nil {
		alloc.Free(unsafe.Pointer(ptr_5D4594_754092))
		ptr_5D4594_754092 = nil
	}
}
func sub_42FFF0(a1 *File) int32 {
	var (
		v1     int32
		result int32
		v3     [12]byte
		v4     [32]byte
	)
	libc.StrCpy(&v3[0], CString("LOOKUPTABLE"))
	v1 = nox_binfile_fread_raw_40ADD0(&v4[0], 1, uint32(libc.StrLen(&v3[0])), a1)
	result = 0
	if v1 == int32(libc.StrLen(&v3[0])) {
		v4[libc.StrLen(&v3[0])] = 0
		if libc.StrCmp(&v4[0], &v3[0]) == 0 {
			result = bool2int(nox_binfile_fread_raw_40ADD0((*byte)(memmap.PtrOff(6112660, 754340)), 1, 0x8000, a1) == 0x8000)
		} else {
			result = 0
		}
	}
	return result
}
func sub_430880(a1 int32) int32 {
	var result int32
	result = a1
	*memmap.PtrUint32(6112660, 805808) = uint32(a1)
	return result
}
func nox_xxx_get_430890() int32 {
	return int32(*memmap.PtrUint32(6112660, 805808))
}
func sub_430AA0(a1 int32) int32 {
	var result int32
	result = a1 - 1
	if a1 == 1 {
		dword_5d4594_805820 = 1
		nox_xxx_useAudio_587000_80772 = 9
	} else {
		result = a1 - 2
		if a1 == 2 {
			dword_5d4594_805820 = 2
			nox_xxx_useAudio_587000_80772 = 13
		} else {
			dword_5d4594_805820 = 0
			nox_xxx_useAudio_587000_80772 = 5
		}
	}
	return result
}
func nox_client_mousePriKey_430AF0() int32 {
	return int32(dword_5d4594_805820)
}
func nox_xxx_cursor_430B00() int32 {
	return int32(nox_xxx_useAudio_587000_80772)
}
func nox_client_setMousePos_430B10(x int32, y int32) {
	nox_client_changeMousePos_430A00(x, y, true)
}
func sub_430B50(a1 int32, a2 int32, a3 int32, a4 int32) int32 {
	var result int32
	dword_5d4594_3807140 = uint32(a1)
	result = a4
	dword_5d4594_3807136 = uint32(a2)
	dword_5d4594_3807116 = uint32(a3)
	dword_5d4594_3807152 = uint32(a4)
	return result
}
func sub_430B80(a1 *uint32) *byte {
	*a1 = dword_5d4594_3804684
	return (*byte)(memmap.PtrOff(0x973F18, 6088))
}
func nox_xxx_tileInitBuf_430DB0(a1 int32, a2 int32) int32 {
	var (
		v2 uint32
		v3 *byte
	)
	dword_5d4594_3798812 = uint32(a1/46 + 4)
	dword_5d4594_3798820 = 0
	dword_5d4594_3798824 = 0
	dword_5d4594_3798800 = dword_5d4594_3798812 * 46
	dword_5d4594_3798816 = uint32(a2/46 + 3)
	dword_5d4594_3798828 = 0
	dword_5d4594_3798832 = 0
	dword_5d4594_3798808 = dword_5d4594_3798816 * 46
	v2 = ((dword_5d4594_3798812 * 46) << uint32(*memmap.PtrUint8(0x973F18, 7696))) * 46 * dword_5d4594_3798816
	dword_5d4594_3798804 = (dword_5d4594_3798812 * 46) << uint32(*memmap.PtrUint8(0x973F18, 7696))
	dword_5d4594_3798836 = 0
	dword_5d4594_3798840 = 0
	v3 = (*byte)(alloc.Calloc(1, int(v2)))
	dword_5d4594_3798796 = uint32(uintptr(unsafe.Pointer(v3)))
	if v3 == nil {
		return 0
	}
	dword_5d4594_3798844 = uint32(uintptr(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v3), v2)))))
	return 1
}
func nox_video_freeFloorBuffer_430EC0() int32 {
	if dword_5d4594_3798796 != 0 {
		*(*unsafe.Pointer)(unsafe.Pointer(&dword_5d4594_3798796)) = nil
		dword_5d4594_3798796 = 0
	}
	sub_444C50()
	return 1
}
func nox_xxx_freeFloorBuffer_430EF0() {
	nox_video_freeFloorBuffer_430EC0()
}

var screenshot_num int32 = 0

func nox_audio_initall(a3 int32) int32 {
	if nox_enable_audio == 0 {
		return 1
	}
	var v1 int32
	var v2 int32
	var v3 int32
	if a3 != 0 {
		sub_486F30()
		if sub_4311F0() != 0 {
			dword_587000_81128 = unsafe.Pointer(uintptr(dword_5d4594_805984 + 88))
			dword_5d4594_805980 = uint32(uintptr(unsafe.Pointer(sub_4866F0(CString("audio"), CString("audio")))))
		}
	}
	sub_4864A0((*uint32)(memmap.PtrOff(6112660, 805884)))
	sub_4864A0(*(**uint32)(unsafe.Pointer(&dword_587000_93164)))
	sub_4864A0(*(**uint32)(unsafe.Pointer(&dword_587000_122852)))
	sub_4864A0(*(**uint32)(unsafe.Pointer(&dword_587000_127004)))
	nox_xxx_WorkerHurt_44D810()
	sub_43D8E0()
	sub_451850(*(*int32)(unsafe.Pointer(&dword_5d4594_805984)), *(*int32)(unsafe.Pointer(&dword_5d4594_805980)))
	v1 = sub_4866A0(2)
	if v1 == 0 {
		sub_43DC00()
	}
	sub_486320(*(**uint32)(unsafe.Pointer(&dword_587000_93164)), v1)
	v2 = sub_4866A0(1)
	if v2 == 0 {
		sub_44D960()
	}
	sub_486320(*(**uint32)(unsafe.Pointer(&dword_587000_122852)), v2)
	v3 = sub_4866A0(0)
	if v3 == 0 {
		sub_453050()
	}
	sub_486320(*(**uint32)(unsafe.Pointer(&dword_587000_127004)), v3)
	return 1
}
func sub_4311F0() int32 {
	var (
		v0 *uint32
		v2 [7]int32
	)
	sub_486FA0(*memmap.PtrInt32(0x587000, 94032))
	nullsub_10(uint32(uintptr(unsafe.Pointer(&v2[0]))))
	v2[2] = 22050
	v2[1] = 0
	v2[3] = 2
	v2[4] = 2
	v2[0] = 4
	sub_487D00((*uint32)(unsafe.Pointer(&v2[0])))
	v0 = sub_487150(-1, unsafe.Pointer(&v2[0]))
	dword_5d4594_805984 = uint32(uintptr(unsafe.Pointer(v0)))
	return bool2int(v0 != nil && sub_487790(int32(uintptr(unsafe.Pointer(v0))), 16) == 16)
}
func sub_431270() {
	if dword_5d4594_805984 != 0 {
		sub_487680(*(*unsafe.Pointer)(unsafe.Pointer(&dword_5d4594_805984)))
		dword_5d4594_805984 = 0
	}
}
func sub_431290() *int32 {
	var result *int32
	result = *(**int32)(unsafe.Pointer(&dword_5d4594_805984))
	if dword_5d4594_805984 != 0 {
		result = sub_487970(*(*int32)(unsafe.Pointer(&dword_5d4594_805984)), -1)
	}
	return result
}
func sub_431330() int32 {
	return bool2int(dword_5d4594_805984 != 0)
}
func sub_431340() int32 {
	return 1
}
func sub_431370() int32 {
	return bool2int(sub_488B60() != 0)
}
func sub_431380() {
	sub_488BA0()
	sub_4896E0()
}
func nox_client_newScreenParticle_431540(a1 int32, a2 int32, a3 int32, a4 int32, a5 int32, a6 int32, a7 int8, a8 int8, a9 int8, a10 int8) *nox_screenParticle {
	var (
		v10 int32
		v11 int32
	)
	if nox_alloc_screenParticles_806044 == nil {
		return nil
	}
	switch a1 {
	case 0:
		v10 = int32(*memmap.PtrUint32(6112660, 806016))
		v11 = int32(*memmap.PtrUint32(6112660, 806036))
	case 1:
		v10 = int32(*memmap.PtrUint32(6112660, 806028))
		v11 = int32(*memmap.PtrUint32(6112660, 806004))
	case 2:
		v10 = int32(*memmap.PtrUint32(6112660, 806032))
		v11 = int32(*memmap.PtrUint32(6112660, 806040))
	case 3:
		v10 = int32(*memmap.PtrUint32(6112660, 806020))
		v11 = int32(*memmap.PtrUint32(6112660, 806012))
	case 4:
		v10 = int32(*memmap.PtrUint32(6112660, 806008))
		v11 = int32(*memmap.PtrUint32(6112660, 806024))
	default:
		return nil
	}
	var p *nox_screenParticle = (*nox_screenParticle)(nox_alloc_class_new_obj_zero((*nox_alloc_class)(nox_alloc_screenParticles_806044)))
	if p == nil {
		p = dword_5d4594_806052
		if p == nil {
			return nil
		}
		sub_4316C0(p)
	}
	p.field_24 = uint32(a2 << 16)
	p.field_28 = uint32(a3 << 16)
	p.field_40[0] = uint8(a7)
	p.field_40[1] = uint8(a8)
	p.field_40[2] = uint8(a9)
	p.field_40[3] = uint8(a8)
	p.draw_fnc = nox_client_screenParticleDraw_489700
	p.field_16 = uint32(a4 << 16)
	p.field_20 = uint32(a5 << 16)
	p.field_36 = uint32(a6 << 16)
	*(*uint8)(unsafe.Pointer(&p.field_32)) = uint8(a10)
	p.field_4 = uint32(a1)
	p.field_8 = uint32(v10)
	p.field_12 = uint32(v11)
	nox_client_addScreenParticle_431680(p)
	if p.field_36 == 0 && int32(p.field_40[1]) == 0 {
		p.field_40[1] = 3
		p.field_40[2] = 2
		p.field_40[3] = 3
	}
	return p
}
func nox_client_addScreenParticle_431680(p *nox_screenParticle) {
	p.field_44 = nox_screenParticles_head
	p.field_48 = nil
	if nox_screenParticles_head != nil {
		nox_screenParticles_head.field_48 = p
	} else {
		dword_5d4594_806052 = p
	}
	nox_screenParticles_head = p
}
func sub_4316C0(p *nox_screenParticle) {
	if p == dword_5d4594_806052 {
		dword_5d4594_806052 = p.field_48
	}
	var v2 *nox_screenParticle = p.field_44
	if v2 != nil {
		v2.field_48 = p.field_48
	}
	var v3 *nox_screenParticle = p.field_48
	if v3 != nil {
		v3.field_44 = p.field_44
	} else {
		nox_screenParticles_head = p.field_44
	}
}
func sub_431700(a1 *uint64) {
	sub_4316C0((*nox_screenParticle)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1)))))))
	nox_alloc_class_free_obj_first((*nox_alloc_class)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&nox_alloc_screenParticles_806044)))), unsafe.Pointer(a1))
}
func nox_client_screenParticlesDraw_431720(rdr *nox_draw_viewport_t) {
	if rdr == nil {
		return
	}
	sub_430B50(rdr.x1, rdr.y1, rdr.x2, rdr.y2)
	var p2 *nox_screenParticle = nil
	for p := (*nox_screenParticle)(nox_screenParticles_head); p != nil; p = p2 {
		dword_5d4594_3799524 = 1
		p2 = p.field_44
		p.draw_fnc(unsafe.Pointer(rdr), p)
	}
}
func nox_xxx_getHostInfoPtr_431770() *byte {
	return (*byte)(memmap.PtrOff(6112660, 807172))
}
func nox_xxx_copyServerIPAndPort_431790(a1 *byte) *byte {
	var result *byte
	result = a1
	if a1 != nil {
		result = libc.StrNCpy((*byte)(memmap.PtrOff(6112660, 806060)), a1, 23)
	}
	return result
}
func sub_434510(a1 *byte, a2 uint8, a3 uint8, a4 uint8) *byte {
	var (
		result *byte
		v5     [10]int32
	)
	v5[6] = int32(a2)
	result = a1
	v5[7] = int32(a3)
	v5[8] = int32(a4)
	libc.MemCpy(unsafe.Pointer(a1), unsafe.Pointer(&v5[0]), 40)
	return result
}
func sub_434610(a1 int32) int32 {
	var result int32
	result = a1
	dword_5d4594_810636 = uint32(a1)
	return result
}
func sub_434620(a1 int32) int32 {
	var result int32
	result = a1
	dword_5d4594_810632 = uint32(a1)
	return result
}
func sub_434630(a1 uint8, a2 uint8, a3 uint8) uint8 {
	var (
		v3     int32
		v4     *uint8
		v5     int32
		v6     uint32
		v7     int32
		v8     int32
		result uint8
		v10    *uint8
		v11    int32
		v12    int32
		v13    int32
		v14    *uint8
	)
	_ = v14
	var v15 int32
	var v16 int32
	var v17 uint32
	var v18 uint8
	v3 = 0
	v4 = (*uint8)(memmap.PtrOff(0x973F18, 5096))
	if *(*int32)(unsafe.Pointer(&dword_5d4594_808568)) <= 0 {
	LABEL_6:
		v5 = int32(a1)
		v6 = uint32(int32(a1) | (int32(a2)|int32(a3)<<8)<<8)
		v7 = 0
		v15 = math.MaxUint8
		for {
			v8 = (v7 + v15) / 2
			if v6 >= uint32(*memmap.PtrInt32(6112660, uint32(v8*4)+0xC5A84)) {
				if v6 == 0 {
					return *memmap.PtrUint8(6112660, uint32(v8)+0xC5570)
				}
				v7 = v8 + 1
			} else {
				v15 = v8 - 1
			}
			if v7 > v15 {
				break
			}
		}
		v17 = 0x2FA03
		v18 = 0
		v16 = 0
		v10 = (*uint8)(memmap.PtrOff(0x973F18, 3882))
		for {
			v11 = int32(*((*uint8)(unsafe.Add(unsafe.Pointer(v10), -2)))) - v5
			v12 = int32(*((*uint8)(unsafe.Add(unsafe.Pointer(v10), -1)))) - int32(a2)
			if uint32(v11*v11+v12*v12)+uint32(int32(*v10)-int32(a3))*(uint32(*v10)-uint32(a3)) < v17 {
				v17 = uint32(v11*v11 + v12*v12 + (int32(*v10)-int32(a3))*(int32(*v10)-int32(a3)))
				v18 = uint8(int8(v16))
			}
			v10 = (*uint8)(unsafe.Add(unsafe.Pointer(v10), 4))
			if func() int32 {
				p := &v16
				*p++
				return *p
			}() >= 256 {
				break
			}
			v5 = int32(a1)
		}
		v13 = int32(dword_5d4594_810628)
		*memmap.PtrUint8(0x973F18, dword_5d4594_810628*4+5096) = a1
		v14 = (*uint8)(memmap.PtrOff(0x973F18, uint32(v13*4+5096)))
		*(*uint8)(unsafe.Add(unsafe.Pointer(v14), 1)) = a2
		*(*uint8)(unsafe.Add(unsafe.Pointer(v14), 2)) = a3
		*(*uint8)(unsafe.Add(unsafe.Pointer(v14), 3)) = v18
		if *(*int32)(unsafe.Pointer(&dword_5d4594_808568)) < 32 {
			dword_5d4594_808568++
		}
		if func() int32 {
			p := (*int32)(unsafe.Pointer(&dword_5d4594_810628))
			*p++
			return *p
		}() >= 32 {
			dword_5d4594_810628 = 0
		}
		result = v18
	} else {
		for int32(a1) != int32(*v4) || int32(a2) != int32(*(*uint8)(unsafe.Add(unsafe.Pointer(v4), 1))) || int32(a3) != int32(*(*uint8)(unsafe.Add(unsafe.Pointer(v4), 2))) {
			v4 = (*uint8)(unsafe.Add(unsafe.Pointer(v4), 4))
			if func() int32 {
				p := &v3
				*p++
				return *p
			}() >= *(*int32)(unsafe.Pointer(&dword_5d4594_808568)) {
				goto LABEL_6
			}
		}
		result = *(*uint8)(unsafe.Add(unsafe.Pointer(v4), 3))
	}
	return result
}
func sub_4347C0(a1 int8) int8 {
	var (
		result int8
		v2     int32
	)
	result = a1
	v2 = 256
	for {
		v2--
		*memmap.PtrUint8(0x973F18, 3882) = uint8(a1)
		*memmap.PtrUint8(0x973F18, 3881) = uint8(a1)
		*memmap.PtrUint8(0x973F18, 3880) = uint8(a1)
		*memmap.PtrUint8(0x973F18, 3883) = 0
		if v2 == 0 {
			break
		}
	}
	return result
}
func sub_434820(a1 int32, a2 int8, a3 int8, a4 int8) int32 {
	if a1 < 0 || a1 >= 256 {
		return 0
	}
	*memmap.PtrUint8(0x973F18, uint32(a1*4+3880)) = uint8(a2)
	*memmap.PtrUint8(0x973F18, uint32(a1*4+3881)) = uint8(a3)
	*memmap.PtrUint8(0x973F18, uint32(a1*4+3882)) = uint8(a4)
	*memmap.PtrUint8(0x973F18, uint32(a1*4+3883)) = 4
	sub_435040()
	sub_434F00()
	nox_xxx_makeFillerColor_48BDE0()
	return 1
}
func sub_434870(a1 int32, a2 *uint8, a3 *uint8, a4 *uint8) int32 {
	if a1 < 0 || a1 >= 256 {
		return 0
	}
	if a2 != nil {
		*a2 = *memmap.PtrUint8(0x973F18, uint32(a1*4+3880))
	}
	if a3 != nil {
		*a3 = *memmap.PtrUint8(0x973F18, uint32(a1*4+3881))
	}
	if a4 != nil {
		*a4 = *memmap.PtrUint8(0x973F18, uint32(a1*4+3882))
	}
	return 1
}
func sub_434920() int32 {
	var result int32
	if dword_5d4594_808564 != 0 {
		return 0
	}
	result = 1
	libc.MemCpy(memmap.PtrOff(6112660, 808572), memmap.PtrOff(0x973F18, 3880), 1024)
	dword_5d4594_808564 = 1
	return result
}
func sub_434950() int32 {
	if dword_5d4594_808564 == 0 {
		return 0
	}
	dword_5d4594_808564 = 0
	libc.MemCpy(memmap.PtrOff(0x973F18, 3880), memmap.PtrOff(6112660, 808572), 1024)
	sub_435040()
	sub_434F00()
	return 1
}
func sub_434E80(a1 int8, a2 int8, a3 int8) int8 {
	return int8(*(*uint8)(unsafe.Pointer(uintptr(uint32(uintptr(nox_draw_colorTablesRev_3804668)) + uint32((int32(uint8(int8(int32(a3)&248)))>>3)|(int32(a2)&248)*4|(int32(a1)&248)<<7)))))
}
func sub_434EC0(a1 int8, a2 int8, a3 int8) int8 {
	return int8(*(*uint8)(unsafe.Pointer(uintptr(uint32(uintptr(nox_draw_colorTablesRev_3804668)) + uint32((int32(uint8(int8(int32(a3)&248)))>>3)|(int32(a2)&252)*8|(int32(a1)&248)<<8)))))
}
func sub_434F00() int32 {
	return 1
}
func sub_435040() {
	var (
		buf  [256]pixel8888
		data *uint8
	)
	data = (*uint8)(memmap.PtrOff(0x973F18, 3880))
	for i := int32(0); i < 256; i++ {
		buf[i].field_0 = int8(i)
		buf[i].field_1 = int8(*(*uint8)(unsafe.Add(unsafe.Pointer(data), i*4+0)))
		buf[i].field_2 = int8(*(*uint8)(unsafe.Add(unsafe.Pointer(data), i*4+1)))
		buf[i].field_3 = int8(*(*uint8)(unsafe.Add(unsafe.Pointer(data), i*4+2)))
	}
	sub_48C580(&buf[0], 256)
	data = (*uint8)(memmap.PtrOff(6112660, 809604))
	for i := int32(0); i < 256; i++ {
		*(*uint8)(unsafe.Add(unsafe.Pointer(data), i*4+0)) = uint8(buf[i].field_1)
		*(*uint8)(unsafe.Add(unsafe.Pointer(data), i*4+1)) = uint8(buf[i].field_2)
		*(*uint8)(unsafe.Add(unsafe.Pointer(data), i*4+2)) = uint8(buf[i].field_3)
		*(*uint8)(unsafe.Add(unsafe.Pointer(data), i*4+3)) = 0
		*memmap.PtrUint8(6112660, uint32(i)+0xC5570) = uint8(buf[i].field_0)
	}
	dword_5d4594_808568 = 0
	dword_5d4594_810628 = 0
}
func sub_4350E0(a1 *uint8, a2 *uint8) *uint8 {
	var (
		v2     *uint8
		result *uint8
		v4     int32
	)
	v2 = a2
	result = a1
	v4 = 256
	for {
		*result = *v2
		*(*uint8)(unsafe.Add(unsafe.Pointer(result), 1)) = *(*uint8)(unsafe.Add(unsafe.Pointer(v2), 1))
		*(*uint8)(unsafe.Add(unsafe.Pointer(result), 2)) = *(*uint8)(unsafe.Add(unsafe.Pointer(v2), 2))
		*(*uint8)(unsafe.Add(unsafe.Pointer(result), 3)) = 4
		result = (*uint8)(unsafe.Add(unsafe.Pointer(result), 4))
		v2 = (*uint8)(unsafe.Add(unsafe.Pointer(v2), 4))
		v4--
		if v4 == 0 {
			break
		}
	}
	return result
}
func sub_435120(a1 *uint8, a2 *byte) {
	var (
		result *byte
		v3     *uint8
		v4     int32
		v5     int8
		v6     *byte
		v7     int8
		v8     int8
	)
	result = a2
	v3 = a1
	v4 = 256
	for {
		v5 = int8(*result)
		v6 = (*byte)(unsafe.Add(unsafe.Pointer(result), 1))
		*v3 = uint8(v5)
		v7 = int8(*func() *byte {
			p := &v6
			x := *p
			*p = (*byte)(unsafe.Add(unsafe.Pointer(*p), 1))
			return x
		}())
		*(*uint8)(unsafe.Add(unsafe.Pointer(v3), 1)) = uint8(v7)
		v8 = int8(*v6)
		*(*uint8)(unsafe.Add(unsafe.Pointer(v3), 3)) = 4
		*(*uint8)(unsafe.Add(unsafe.Pointer(v3), 2)) = uint8(v8)
		result = (*byte)(unsafe.Add(unsafe.Pointer(v6), 1))
		v3 = (*uint8)(unsafe.Add(unsafe.Pointer(v3), 4))
		v4--
		if v4 == 0 {
			break
		}
	}
}
func sub_435150(a1 *uint8, a2 *byte) {
	var (
		v2     *byte
		result *uint8
		v4     int32
		v5     int8
		v6     *uint8
	)
	v2 = a2
	result = a1
	v4 = 256
	for {
		v5 = int8(*v2)
		v2 = (*byte)(unsafe.Add(unsafe.Pointer(v2), 4))
		*result = uint8(v5)
		v6 = (*uint8)(unsafe.Add(unsafe.Pointer(result), 1))
		*func() *uint8 {
			p := &v6
			x := *p
			*p = (*uint8)(unsafe.Add(unsafe.Pointer(*p), 1))
			return x
		}() = uint8(*((*byte)(unsafe.Add(unsafe.Pointer(v2), -3))))
		*v6 = uint8(*((*byte)(unsafe.Add(unsafe.Pointer(v2), -2))))
		result = (*uint8)(unsafe.Add(unsafe.Pointer(v6), 1))
		v4--
		if v4 == 0 {
			break
		}
	}
}
func sub_435240(a1 uint8, a2 *uint8, a3 *uint8, a4 *byte) {
	var (
		v4     int32
		result int8
	)
	v4 = int32(a1) * 4
	*a2 = *memmap.PtrUint8(0x973F18, uint32(v4+3880))
	*a3 = *memmap.PtrUint8(0x973F18, uint32(v4+3881))
	result = int8(*memmap.PtrUint8(0x973F18, uint32(v4+3882)))
	*a4 = byte(result)
}
func sub_435280(a1 int16, a2 *uint8, a3 *uint8, a4 *uint8) {
	var result int8
	*a2 = uint8(int8(int32(uint16(int16(int32(*((*uint16)(unsafe.Pointer(&byte_5D4594_3804364[0+0]))))&int32(a1)))) >> int32(byte_5D4594_3804364[0+12])))
	*a3 = uint8(int8(int32(uint16(int16(int32(*((*uint16)(unsafe.Pointer(&byte_5D4594_3804364[0+4]))))&int32(a1)))) >> int32(byte_5D4594_3804364[0+16])))
	result = int8((int32(uint8(int8(a1))) & int32(byte_5D4594_3804364[0+8])) << int32(byte_5D4594_3804364[0+20]))
	*a4 = uint8(result)
}
func nox_xxx_initTime_435570() int64 {
	var result int64
	result = int64(platformTicks())
	*memmap.PtrUint64(6112660, 811908) = uint64(result)
	return result
}
func sub_435590() int32 {
	return int32(*memmap.PtrUint32(6112660, 811916))
}
func sub_435690(a1 *uint32) *uint32 {
	var result *uint32
	result = a1
	*a1 = *memmap.PtrUint32(6112660, 811364)
	*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*1)) = *memmap.PtrUint32(6112660, 811368)
	return result
}
func nox_xxx_spriteCheckFlag31_4356C0(dr *nox_drawable, a2 int8) bool {
	var (
		a1     int32 = int32(uintptr(unsafe.Pointer(dr)))
		result int32
	)
	result = a1
	if a1 != 0 {
		result = bool2int((uint32(1<<int32(a2)) & *(*uint32)(unsafe.Pointer(uintptr(a1 + 124)))) != 0)
	}
	return result != 0
}
func nox_xxx_spriteLoadError_4356E0() {
	*memmap.PtrUint32(0x587000, 85720) = 0
}
func sub_435700(a1 *libc.WChar, a2 int32) *libc.WChar {
	var result *libc.WChar
	result = nox_wcscpy((*libc.WChar)(memmap.PtrOff(6112660, 811376)), a1)
	*memmap.PtrUint32(6112660, 811060) = uint32(a2)
	return result
}
func sub_435740() uint32 {
	var result uint32
	result = platformTicks()
	*memmap.PtrUint32(6112660, 811924) = result
	return result
}
func sub_435750() uint32 {
	var result uint32
	result = platformTicks() - *memmap.PtrUint32(6112660, 811924)
	*memmap.PtrUint32(6112660, 811928) = result
	return result
}
func nox_xxx_time_startProfile_435770() uint32 {
	var result uint32
	result = platformTicks()
	*memmap.PtrUint32(6112660, 811932) = result
	return result
}
func nox_xxx_time_endProfile_435780() uint32 {
	var result uint32
	result = platformTicks() - *memmap.PtrUint32(6112660, 811932)
	*memmap.PtrUint32(6112660, 811936) = result
	return result
}
func nox_xxx_cliToggleObsWindow_4357A0() int32 {
	var result int32
	if *memmap.PtrUint32(0x8531A0, 2576) != 0 && int32(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(0x8531A0, 2576) + 3680))))&1 != 0 {
		result = nox_xxx_showObserverWindow_48CA70(0)
	} else {
		result = nox_xxx_showObserverWindow_48CA70(1)
	}
	return result
}
func sub_435F60() int32 {
	var result int32
	result = int32(1 - dword_5d4594_811904)
	dword_5d4594_811904 = 1 - dword_5d4594_811904
	return result
}
func sub_436550() int32 {
	var v0 int32
	if sub_459DA0() != 0 || nox_gui_xxx_check_446360() != 0 || sub_49CB40() != 0 || sub_49C810() != 0 || sub_446950() != 0 || sub_4706A0() != 0 || nox_gui_console_flagXxx_451410() != 0 {
		v0 = int32(nox_frame_xxx_2598000)
	} else {
		v0 = int32(nox_frame_xxx_2598000)
		if nox_frame_xxx_2598000 != 2 {
			return bool2int(nox_frame_xxx_2598000-*memmap.PtrUint32(6112660, 811920) == 1)
		}
	}
	*memmap.PtrUint32(6112660, 811920) = uint32(v0)
	return 1
}
func sub_437100() int32 {
	var result int32
	result = int32(nox_client_renderGUI_80828)
	if *memmap.PtrUint32(6112660, 811064) != nox_client_renderGUI_80828 && !nox_common_getEngineFlag(nox_engine_flag(NOX_ENGINE_FLAG_DISABLE_GRAPHICS_RENDERING)) {
		*memmap.PtrUint32(6112660, 811064) = nox_client_renderGUI_80828
		sub_4721A0(*(*int32)(unsafe.Pointer(&nox_client_renderGUI_80828)))
		sub_460EA0(*(*int32)(unsafe.Pointer(&nox_client_renderGUI_80828)))
		nox_window_set_visible_unk5(*(*int32)(unsafe.Pointer(&nox_client_renderGUI_80828)))
		sub_45D500(*(*int32)(unsafe.Pointer(&nox_client_renderGUI_80828)))
		sub_455A00(*(*int32)(unsafe.Pointer(&nox_client_renderGUI_80828)))
		sub_455F10(*(*int32)(unsafe.Pointer(&nox_client_renderGUI_80828)))
		sub_4706C0(*(*int32)(unsafe.Pointer(&nox_client_renderGUI_80828)))
		result = int32(nox_client_renderGUI_80828)
		if nox_client_renderGUI_80828 == 0 {
			result = sub_478000()
		}
	}
	return result
}
func nox_xxx_playerAnimCheck_4372B0() int32 {
	var (
		v0     int32
		result int32
	)
	result = 1
	if *memmap.PtrUint32(0x852978, 8) != 0 {
		v0 = int32(*(*uint32)(unsafe.Pointer(uintptr(*memmap.PtrUint32(0x852978, 8) + 276))))
		if v0 != 1 && v0 != 2 && v0 != 51 {
			result = 0
		}
	}
	return result
}
func nox_xxx_clientIsObserver_4372E0() int32 {
	var result int32
	if *memmap.PtrUint32(0x8531A0, 2576) != 0 && *(*uint32)(unsafe.Pointer(uintptr(*memmap.PtrUint32(0x8531A0, 2576) + 2092))) == 1 {
		result = bool2int((*(*uint32)(unsafe.Pointer(uintptr(*memmap.PtrUint32(0x8531A0, 2576) + 3680))) & 3) != 0)
	} else {
		result = 0
	}
	return result
}
func sub_437320(a1 int32) int32 {
	var (
		v1     int32
		v2     *uint8
		v3     int32
		v4     int32
		result int32
	)
	v1 = 0
	v2 = (*uint8)(memmap.PtrOff(0x587000, 87484))
	for {
		if *(*uint32)(unsafe.Pointer(uintptr(a1 + 96))) <= uint32(*(*int32)(unsafe.Pointer(v2))) {
			break
		}
		v2 = (*uint8)(unsafe.Add(unsafe.Pointer(v2), 4))
		v1++
		if int32(uintptr(unsafe.Pointer(v2))) >= int32(uintptr(memmap.PtrOff(0x587000, 87496))) {
			break
		}
	}
	if v1 > 2 {
		v1 = 2
	}
	v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 28))) + 36)
	v4 = v1 * 16
	if *(*int32)(unsafe.Pointer(&dword_587000_87412)) == -1 {
		*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 28))) + 76))) = *memmap.PtrUint32(6112660, uint32(v4)+0xC6DE8)
		*(*uint32)(unsafe.Pointer(uintptr(v3 + 24))) = *memmap.PtrUint32(6112660, uint32(v4)+0xC6DE4)
		result = int32(*memmap.PtrUint32(6112660, uint32(v4)+0xC6DE4))
	} else {
		*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 28))) + 76))) = *memmap.PtrUint32(6112660, uint32(v4)+0xC6DE0)
		*(*uint32)(unsafe.Pointer(uintptr(v3 + 24))) = *memmap.PtrUint32(6112660, uint32(v4)+0xC6DDC)
		result = int32(*memmap.PtrUint32(6112660, uint32(v4)+0xC6DDC))
	}
	*(*uint32)(unsafe.Pointer(uintptr(v3 + 48))) = uint32(result)
	return result
}
func sub_4375C0(a1 int32) {
	if nox_wol_server_result_cnt_815088 != 0 {
		sub_46AD20(*(**uint32)(unsafe.Pointer(&nox_wol_wnd_world_814980)), 10070, int32(nox_wol_server_result_cnt_815088+0x2755), a1)
	}
}
func sub_437860(a1 int32, a2 int32) int32 {
	var (
		result int32
		v3     *uint8
	)
	result = 0
	v3 = (*uint8)(memmap.PtrOff(0x587000, 87532))
	for a1 <= int32(*((*int16)(unsafe.Add(unsafe.Pointer((*int16)(unsafe.Pointer(v3))), -int(unsafe.Sizeof(int16(0))*2))))) || a1 >= int32(*(*int16)(unsafe.Pointer(v3))) || a2 <= int32(*((*int16)(unsafe.Add(unsafe.Pointer((*int16)(unsafe.Pointer(v3))), -int(unsafe.Sizeof(int16(0))*1))))) || a2 >= int32(*((*int16)(unsafe.Add(unsafe.Pointer((*int16)(unsafe.Pointer(v3))), unsafe.Sizeof(int16(0))*1)))) {
		v3 = (*uint8)(unsafe.Add(unsafe.Pointer(v3), 8))
		result++
		if int32(uintptr(unsafe.Pointer(v3))) >= int32(uintptr(memmap.PtrOff(0x587000, 87564))) {
			return 0
		}
	}
	return result
}
func sub_4379C0() {
	if dword_587000_87408 == 1 {
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&nox_wol_wnd_gameList_815012))))).Func94(asWindowEvent(0x400F, 0, 0))
	}
}
func sub_438330() int32 {
	var v0 func() int32
	v0 = nox_wnd_xxx_815040.field_13
	nox_gui_freeAnimation_43C570(nox_wnd_xxx_815040)
	if !noxflags.HasGame(noxflags.GameFlag29) {
		nox_client_guiXxx_43A9D0()
	}
	if v0 != nil {
		v0()
	}
	return 1
}
func sub_438370() int32 {
	if nox_wnd_xxx_815040.state == nox_gui_anim_state(NOX_GUI_ANIM_OUT_DONE) {
		return sub_438330()
	}
	nox_wnd_xxx_815040.state = nox_gui_anim_state(NOX_GUI_ANIM_OUT)
	sub_43BE40(2)
	clientPlaySoundSpecial(923, 100)
	return 1
}
func sub_438480() int32 {
	(*(*int32)(unsafe.Pointer(&nox_wol_wnd_gameList_815012))).setFunc94(func(arg1 int32, arg2 int32, arg3 int32, arg4 int32) int32 {
		return sub_439050(arg1, uint32(arg2), (*int32)(unsafe.Pointer(uintptr(arg3))), uint32(arg4))
	})
	(*(*int32)(unsafe.Pointer(&nox_wol_wnd_gameList_815012))).setFunc93(func(arg1 int32, arg2 int32, arg3 int32, arg4 int32) int32 {
		return sub_438EF0((*uint32)(unsafe.Pointer(uintptr(arg1))), arg2, uint32(arg3), arg4)
	})
	sub_46B120((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_815016)))), (*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&nox_wol_wnd_gameList_815012))))))
	(*(*int32)(unsafe.Pointer(&dword_5d4594_815016))).setFunc93(func(arg1 int32, arg2 int32, arg3 int32, arg4 int32) int32 {
		return sub_438EF0((*uint32)(unsafe.Pointer(uintptr(arg1))), arg2, uint32(arg3), arg4)
	})
	sub_46B120((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_815020)))), (*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&nox_wol_wnd_gameList_815012))))))
	(*(*int32)(unsafe.Pointer(&dword_5d4594_815020))).setFunc93(func(arg1 int32, arg2 int32, arg3 int32, arg4 int32) int32 {
		return sub_438EF0((*uint32)(unsafe.Pointer(uintptr(arg1))), arg2, uint32(arg3), arg4)
	})
	sub_46B120((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_815024)))), (*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&nox_wol_wnd_gameList_815012))))))
	(*(*int32)(unsafe.Pointer(&dword_5d4594_815024))).setFunc93(func(arg1 int32, arg2 int32, arg3 int32, arg4 int32) int32 {
		return sub_438EF0((*uint32)(unsafe.Pointer(uintptr(arg1))), arg2, uint32(arg3), arg4)
	})
	sub_46B120((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_815028)))), (*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&nox_wol_wnd_gameList_815012))))))
	(*(*int32)(unsafe.Pointer(&dword_5d4594_815028))).setFunc93(func(arg1 int32, arg2 int32, arg3 int32, arg4 int32) int32 {
		return sub_438EF0((*uint32)(unsafe.Pointer(uintptr(arg1))), arg2, uint32(arg3), arg4)
	})
	sub_46B120((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_815032)))), (*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&nox_wol_wnd_gameList_815012))))))
	(*(*int32)(unsafe.Pointer(&dword_5d4594_815032))).setFunc93(func(arg1 int32, arg2 int32, arg3 int32, arg4 int32) int32 {
		return sub_438EF0((*uint32)(unsafe.Pointer(uintptr(arg1))), arg2, uint32(arg3), arg4)
	})
	**(**uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer(dword_5d4594_815004))) + 32))) + 28))) = 0x2733
	**(**uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer(dword_5d4594_815004))) + 32))) + 32))) = 0x2734
	**(**uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer(dword_5d4594_815004))) + 32))) + 36))) = 0x2730
	(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815016))))).Func94(asWindowEvent(0x4018, int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(nox_wol_wnd_gameList_815012 + 32))) + 28)))), 0))
	(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815020))))).Func94(asWindowEvent(0x4018, int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(nox_wol_wnd_gameList_815012 + 32))) + 28)))), 0))
	(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815024))))).Func94(asWindowEvent(0x4018, int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(nox_wol_wnd_gameList_815012 + 32))) + 28)))), 0))
	(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815028))))).Func94(asWindowEvent(0x4018, int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(nox_wol_wnd_gameList_815012 + 32))) + 28)))), 0))
	(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815032))))).Func94(asWindowEvent(0x4018, int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(nox_wol_wnd_gameList_815012 + 32))) + 28)))), 0))
	(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815016))))).Func94(asWindowEvent(0x4019, int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(nox_wol_wnd_gameList_815012 + 32))) + 32)))), 0))
	(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815020))))).Func94(asWindowEvent(0x4019, int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(nox_wol_wnd_gameList_815012 + 32))) + 32)))), 0))
	(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815024))))).Func94(asWindowEvent(0x4019, int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(nox_wol_wnd_gameList_815012 + 32))) + 32)))), 0))
	(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815028))))).Func94(asWindowEvent(0x4019, int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(nox_wol_wnd_gameList_815012 + 32))) + 32)))), 0))
	(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815032))))).Func94(asWindowEvent(0x4019, int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(nox_wol_wnd_gameList_815012 + 32))) + 32)))), 0))
	(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815016))))).Func94(asWindowEvent(0x401a, int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(nox_wol_wnd_gameList_815012 + 32))) + 36)))), 0))
	(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815020))))).Func94(asWindowEvent(0x401a, int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(nox_wol_wnd_gameList_815012 + 32))) + 36)))), 0))
	(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815024))))).Func94(asWindowEvent(0x401a, int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(nox_wol_wnd_gameList_815012 + 32))) + 36)))), 0))
	(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815028))))).Func94(asWindowEvent(0x401a, int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(nox_wol_wnd_gameList_815012 + 32))) + 36)))), 0))
	return (*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815032))))).Func94(asWindowEvent(0x401a, int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(nox_wol_wnd_gameList_815012 + 32))) + 36)))), 0))
}
func sub_438C80(a1 int32, a2 int32) int32 {
	var (
		v2   [404]byte
		mpos nox_point = getMousePos()
	)
	if (*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815000))))).Flags().IsHidden() == 0 {
		libc.MemCpy(unsafe.Pointer(&v2[0]), *(*unsafe.Pointer)(unsafe.Pointer(&dword_5d4594_815000)), int(unsafe.Sizeof([404]byte{})))
		*(*uint32)(unsafe.Pointer(&v2[16])) -= 32
		*(*uint32)(unsafe.Pointer(&v2[20])) -= 32
		*(*uint32)(unsafe.Pointer(&v2[8])) += 64
		*(*uint32)(unsafe.Pointer(&v2[12])) += 64
		if dword_5d4594_815044 == 0 && !nox_xxx_wndPointInWnd_46AAB0((*uint32)(unsafe.Pointer(&v2[0])), mpos.x, mpos.y) {
			(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815000))))).Hide()
			(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&nox_wol_wnd_gameList_815012))))).Func94(asWindowEvent(0x4013, -1, 0))
			dword_5d4594_815056 = 0
			nox_xxx_wnd_46C6E0((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815000))))))
			guiFocus((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&nox_wol_wnd_world_814980))))))
		}
	}
	if sub_4A28B0() != 0 && !nox_xxx_wndPointInWnd_46AAB0(*(**uint32)(memmap.PtrOff(6112660, 815036)), mpos.x, mpos.y) {
		sub_4A2890()
		guiFocus((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&nox_wol_wnd_world_814980))))))
	}
	if nox_game_createOrJoin_815048 != 0 && sub_438DD0(uint32(mpos.x), uint32(mpos.y)) != 0 {
		nox_client_setCursorType_477610(9)
	} else if sub_44A4A0() == 0 {
		nox_client_setCursorType_477610(0)
	}
	return 1
}
func sub_438DD0(a1 uint32, a2 uint32) int32 {
	if *(*int32)(unsafe.Pointer(&dword_587000_87412)) == -1 {
		if a1 > 216 && a1 < 600 && a2 > 27 && a2 < 451 {
			return 1
		}
	} else if a1 > 226 && a1 < 590 && a2 > 37 && a2 < 441 {
		return 1
	}
	return 0
}
func sub_438E30(a1 *uint32, a2 int32) int32 {
	var (
		v1 *uint32
		v2 int32
		v3 int32
		v4 **int16
		v6 int32
	)
	v1 = a1
	nox_client_wndGetPosition_46AA60((*nox_window)(unsafe.Pointer(a1)), (*uint32)(unsafe.Pointer(&v6)), (*uint32)(unsafe.Pointer(&a1)))
	v2 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*25)))
	if *(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*9))&6 != 0 {
		nox_client_drawImageAt_47D2C0((*nox_video_bag_image_t)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*19))))), int32(uint32(v6)+*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*24))), int32(uintptr(unsafe.Pointer(a1)))+v2)
	} else {
		nox_client_drawImageAt_47D2C0((*nox_video_bag_image_t)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*15))))), int32(uint32(v6)+*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*24))), int32(uintptr(unsafe.Pointer(a1)))+v2)
	}
	v3 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*100)))
	if v3 == 0 {
		return 1
	}
	for {
		if (int32(*(*uint8)(unsafe.Pointer(uintptr(v3 + 4))))&16) == 0 && *(*uint32)(unsafe.Pointer(uintptr(v3 + 44))) == 2048 {
			v4 = *(***int16)(unsafe.Pointer(uintptr(v3 + 32)))
			nox_xxx_drawSetTextColor_434390(int32(nox_color_white_2523948))
			nox_xxx_drawStringStyle_43F7B0(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v3 + 236))))), (*libc.WChar)(unsafe.Pointer(*v4)), int32(uint32(v6)+*(*uint32)(unsafe.Pointer(uintptr(v3 + 16)))), int32(uint32(int32(uintptr(unsafe.Pointer(a1))))+*(*uint32)(unsafe.Pointer(uintptr(v3 + 20)))))
		}
		v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(v3 + 388))))
		if v3 == 0 {
			break
		}
	}
	return 1
}
func sub_438EF0(a1 *uint32, a2 int32, a3 uint32, a4 int32) int32 {
	var (
		result int32
		v5     int32
		v6     int32
	)
	if a2 == 19 {
		v6 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*8)))
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&nox_wol_wnd_gameList_815012))))).Func94(asWindowEvent(0x4007, int32(*(*uint32)(unsafe.Pointer(uintptr(v6 + 28)))), 0))
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815016))))).Func94(asWindowEvent(0x4007, int32(*(*uint32)(unsafe.Pointer(uintptr(v6 + 28)))), 0))
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815020))))).Func94(asWindowEvent(0x4007, int32(*(*uint32)(unsafe.Pointer(uintptr(v6 + 28)))), 0))
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815024))))).Func94(asWindowEvent(0x4007, int32(*(*uint32)(unsafe.Pointer(uintptr(v6 + 28)))), 0))
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815028))))).Func94(asWindowEvent(0x4007, int32(*(*uint32)(unsafe.Pointer(uintptr(v6 + 28)))), 0))
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815032))))).Func94(asWindowEvent(0x4007, int32(*(*uint32)(unsafe.Pointer(uintptr(v6 + 28)))), 0))
		result = 0
	} else if a2 == 20 {
		v5 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*8)))
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&nox_wol_wnd_gameList_815012))))).Func94(asWindowEvent(0x4007, int32(*(*uint32)(unsafe.Pointer(uintptr(v5 + 32)))), 0))
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815016))))).Func94(asWindowEvent(0x4007, int32(*(*uint32)(unsafe.Pointer(uintptr(v5 + 32)))), 0))
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815020))))).Func94(asWindowEvent(0x4007, int32(*(*uint32)(unsafe.Pointer(uintptr(v5 + 32)))), 0))
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815024))))).Func94(asWindowEvent(0x4007, int32(*(*uint32)(unsafe.Pointer(uintptr(v5 + 32)))), 0))
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815028))))).Func94(asWindowEvent(0x4007, int32(*(*uint32)(unsafe.Pointer(uintptr(v5 + 32)))), 0))
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815032))))).Func94(asWindowEvent(0x4007, int32(*(*uint32)(unsafe.Pointer(uintptr(v5 + 32)))), 0))
		result = 0
	} else {
		result = nox_xxx_wndListboxProcWithoutData10_4A28E0(a1, a2, a3, a4)
	}
	return result
}
func sub_439050(a1 int32, a2 uint32, a3 *int32, a4 uint32) int32 {
	var (
		v4 int32
		v5 int32
		v7 int32
	)
	if a2 > 0x400F {
		if a2 == 16400 {
			v7 = (*nox_window)(unsafe.Pointer(a3)).ID()
			if v7 >= 0x2736 && v7 <= 0x273A {
				(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815016))))).Func94(asWindowEvent(0x4013, int32(a4), 0))
				(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815020))))).Func94(asWindowEvent(0x4013, int32(a4), 0))
				(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815024))))).Func94(asWindowEvent(0x4013, int32(a4), 0))
				(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815028))))).Func94(asWindowEvent(0x4013, int32(a4), 0))
				(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815032))))).Func94(asWindowEvent(0x4013, int32(a4), 0))
				if a4 < uint32(*(*int32)(unsafe.Pointer(&nox_wol_server_result_cnt_815088))) {
					var pos nox_point = getMousePos()
					dword_5d4594_814624 = unsafe.Pointer(sub_4A04C0(int32(a4)))
					sub_439370((*int2)(unsafe.Pointer(&pos)), int32(uintptr(dword_5d4594_814624)))
				}
			}
		} else if a2 == 0x4013 || a2 == 0x401C {
			(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815016))))).Func94(asWindowEvent(int32(a2), int32(uintptr(unsafe.Pointer(a3))), 0))
			(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815020))))).Func94(asWindowEvent(int32(a2), int32(uintptr(unsafe.Pointer(a3))), 0))
			(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815024))))).Func94(asWindowEvent(int32(a2), int32(uintptr(unsafe.Pointer(a3))), 0))
			(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815028))))).Func94(asWindowEvent(int32(a2), int32(uintptr(unsafe.Pointer(a3))), 0))
			(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815032))))).Func94(asWindowEvent(int32(a2), int32(uintptr(unsafe.Pointer(a3))), 0))
		}
	} else if a2 >= 0x400E {
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815016))))).Func94(asWindowEvent(int32(a2), int32(uintptr(unsafe.Pointer(a3))), int32(a4)))
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815020))))).Func94(asWindowEvent(int32(a2), int32(uintptr(unsafe.Pointer(a3))), int32(a4)))
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815024))))).Func94(asWindowEvent(int32(a2), int32(uintptr(unsafe.Pointer(a3))), int32(a4)))
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815028))))).Func94(asWindowEvent(int32(a2), int32(uintptr(unsafe.Pointer(a3))), int32(a4)))
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815032))))).Func94(asWindowEvent(int32(a2), int32(uintptr(unsafe.Pointer(a3))), int32(a4)))
	} else {
		switch a2 {
		case 23:
			return 1
		case 0x4000:
			if unsafe.Pointer(a3) == unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&nox_wol_wnd_gameList_815012)))).ChildByID(0x273B)) || unsafe.Pointer(a3) == unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&nox_wol_wnd_gameList_815012)))).ChildByID(0x273C)) {
				(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815016))))).Func94(asWindowEvent(0x4000, int32(uintptr(unsafe.Pointer(a3))), 0))
				(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815020))))).Func94(asWindowEvent(0x4000, int32(uintptr(unsafe.Pointer(a3))), 0))
				(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815024))))).Func94(asWindowEvent(0x4000, int32(uintptr(unsafe.Pointer(a3))), 0))
				(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815028))))).Func94(asWindowEvent(0x4000, int32(uintptr(unsafe.Pointer(a3))), 0))
				(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815032))))).Func94(asWindowEvent(0x4000, int32(uintptr(unsafe.Pointer(a3))), 0))
			}
		case 0x4009:
			v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 32))))
			nox_xxx_wndListboxProcPre_4A30D0((*nox_window)(unsafe.Pointer(uintptr(a1))), 0x4009, uint32(uintptr(unsafe.Pointer((*libc.WChar)(unsafe.Pointer(a3))))), int32(a4))
			v5 = sub_4A4800(v4)
			(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815016))))).Func94(asWindowEvent(0x401C, v5, 0))
			(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815020))))).Func94(asWindowEvent(0x401C, v5, 0))
			(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815024))))).Func94(asWindowEvent(0x401C, v5, 0))
			(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815028))))).Func94(asWindowEvent(0x401C, v5, 0))
			(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815032))))).Func94(asWindowEvent(0x401C, v5, 0))
		}
	}
	return nox_xxx_wndListboxProcPre_4A30D0((*nox_window)(unsafe.Pointer(uintptr(a1))), a2, uint32(uintptr(unsafe.Pointer((*libc.WChar)(unsafe.Pointer(a3))))), int32(a4))
}
func sub_439450(a1 int32, a2 int32, a3 *uint32) *uint32 {
	var (
		result *uint32
		v4     int32
	)
	result = a3
	*a3 = uint32(a1 - 65)
	*(*uint32)(unsafe.Add(unsafe.Pointer(a3), unsafe.Sizeof(uint32(0))*1)) = uint32(a2 - 20)
	if a1-65+130 > 600 {
		*a3 = 470
	}
	if a2-20+120 > 451 {
		*(*uint32)(unsafe.Add(unsafe.Pointer(a3), unsafe.Sizeof(uint32(0))*1)) = 331
	}
	if dword_587000_87408 == 1 {
		v4 = 55
		if *(*uint32)(unsafe.Add(unsafe.Pointer(a3), unsafe.Sizeof(uint32(0))*1)) >= 55 {
			goto LABEL_10
		}
	} else {
		v4 = 27
		if *(*uint32)(unsafe.Add(unsafe.Pointer(a3), unsafe.Sizeof(uint32(0))*1)) >= 27 {
			goto LABEL_10
		}
	}
	*(*uint32)(unsafe.Add(unsafe.Pointer(a3), unsafe.Sizeof(uint32(0))*1)) = uint32(v4)
LABEL_10:
	if int32(*a3) < 216 {
		*a3 = 216
	}
	return result
}
func sub_439CC0(a1 int32, a2 *byte) *byte {
	var (
		v2     uint32
		result *byte
	)
	v2 = uint32(uintptr(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(libc.StrStr((*byte)(unsafe.Pointer(uintptr(a1+52))), CString("'s_game"))), int32(-52-a1))))))
	result = libc.StrNCpy(a2, (*byte)(unsafe.Pointer(uintptr(a1+52))), int(v2))
	*(*byte)(unsafe.Add(unsafe.Pointer(a2), v2)) = 0
	return result
}
func sub_439D00(a1 *int32, a2 int32, a3 uint32, a4 int32) int32 {
	if a2 == 5 {
		if (*nox_window)(unsafe.Pointer(a1)).ID() == 10020 && nox_game_createOrJoin_815048 == 1 {
			sub_439D90(uint32(uint16(a3)), a3>>16)
			return 1
		}
		return 0
	}
	if a2 != 21 {
		return 0
	}
	if a3 != 1 {
		if a3 != 28 && a3 == 57 {
			var mpos nox_point = getMousePos()
			(*nox_window)(unsafe.Pointer(a1)).Func93(asWindowEvent(0x5, mpos.x|mpos.y<<16, 0))
		}
		return 0
	}
	if a4 == 2 {
		sub_4373A0()
	}
	return 1
}
func sub_439D90(a1 uint32, a2 uint32) int32 {
	var (
		result int32
		v3     int16
	)
	result = sub_438DD0(a1, a2)
	if result != 0 {
		v3 = int16(uint16(a2 + uint32(*memmap.PtrUint16(0x587000, dword_587000_87412*8+87530)) - 27))
		*memmap.PtrUint16(6112660, 814916) = uint16(a1 + uint32(*memmap.PtrUint16(0x587000, dword_587000_87412*8+87528)) - 216)
		*memmap.PtrUint16(6112660, 814918) = uint16(v3)
		sub_43B460()
		if sub_43BDB0()&2 != 0 {
			nox_xxx_setQuest_4D6F60(1)
		}
		if nox_xxx_isQuest_4D6F50() != 0 {
			if nox_client_countPlayerFiles04_4DC7D0() != 0 {
				sub_4A7A70(1)
				nox_game_showSelChar_4A4DB0()
				return nox_client_setCursorType_477610(0)
			}
		} else if nox_client_countPlayerFiles02_4DC630() != 0 {
			sub_4A7A70(1)
			nox_game_showSelChar_4A4DB0()
			return nox_client_setCursorType_477610(0)
		}
		sub_4A7A70(0)
		nox_game_showSelClass_4A4840()
		nox_client_setCursorType_477610(0)
	}
	return result
}
func sub_43A920() int32 {
	var result int32
	guiFocus((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&nox_wol_wnd_world_814980))))))
	if dword_587000_87404 == 1 {
		sub_554D10()
		gameSetCliDrawFunc(unsafe.Pointer(cgoFuncAddr(sub_41E210)))
	}
	if sub_43BE30() == 0 || *memmap.PtrUint32(6112660, 815084) == 0 {
		sub_44A400()
	}
	result = sub_43AF90(0)
	dword_5d4594_815044 = 0
	return result
}
func nox_client_guiXxx_43A9D0() int32 {
	nox_xxx_wndClearCaptureMain((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_814984))))))
	sub_489FB0()
	sub_4A2890()
	if dword_5d4594_815000 != 0 && *(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_815000 + 396))) == 0 {
		nox_xxx_wnd_46C6E0((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815000))))))
		(*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_815000)))).Destroy()
		dword_5d4594_815000 = 0
	}
	if nox_wol_wnd_world_814980 != nil {
		(*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&nox_wol_wnd_world_814980)))).Destroy()
		nox_wol_wnd_world_814980 = nil
	}
	sub_43A920()
	guiFocus(nil)
	nox_wol_server_result_cnt_815088 = 0
	sub_49FFA0(0)
	sub_554D10()
	nox_client_setCursorType_477610(0)
	return sub_43DE40(nil)
}
func sub_43AA70() *byte {
	var (
		v0     *byte
		v1     *byte
		v2     *byte
		v3     int8
		v4     int8
		v5     int8
		v6     *byte
		result *byte
		v8     int16
		v9     [32]byte
		v10    [268]byte
	)
	if dword_5d4594_528252 != 0 && dword_5d4594_528256 != 0 {
		nox_xxx_networkLog_printf_413D30(CString("RECON: Posting server to WOL"))
	}
	nox_game_createOrJoin_815048 = 0
	dword_5d4594_815052 = 1
	v0 = nox_xxx_cliGamedataGet_416590(0)
	v1 = (*byte)(sub_416640())
	libc.MemCpy(unsafe.Add(unsafe.Pointer(v1), 111), unsafe.Pointer(v0), 58)
	*(*uint16)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v1), 163)))) = uint16(nox_common_gameFlags_getVal_40A5B0())
	*(*uint32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v1), 135)))) = math.MaxUint32
	*(*uint32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v1), 139)))) = math.MaxUint32
	*(*uint32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v1), 143)))) = math.MaxUint32
	*(*uint32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v1), 147)))) = math.MaxUint32
	*(*uint32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v1), 151)))) = math.MaxUint32
	*(*uint32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v1), 155)))) = math.MaxUint32
	*(*uint32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v1), 159)))) = math.MaxUint32
	v2 = nox_xxx_serverOptionsGetServername_40A4C0()
	libc.StrNCpy((*byte)(unsafe.Add(unsafe.Pointer(v1), 120)), v2, 15)
	libc.StrCpy((*byte)(unsafe.Add(unsafe.Pointer(v1), 111)), nox_xxx_mapGetMapName_409B40())
	if nox_xxx_isQuest_4D6F50() != 0 {
		if dword_5d4594_528256 != 0 {
			*(*uint16)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v1), 165)))) = uint16(int16(nox_game_getQuestStage_4E3CC0()))
		} else {
			*(*uint16)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v1), 165)))) = 1
		}
	}
	*(*byte)(unsafe.Add(unsafe.Pointer(v1), 104)) = byte(int8(nox_xxx_servGetPlrLimit_409FA0()))
	v3 = int8(nox_common_playerInfoCount_416F40())
	*(*byte)(unsafe.Add(unsafe.Pointer(v1), 103)) = byte(v3)
	if nox_common_getEngineFlag(nox_engine_flag(NOX_ENGINE_FLAG_DISABLE_GRAPHICS_RENDERING)) {
		*(*byte)(unsafe.Add(unsafe.Pointer(v1), 103)) = byte(int8(int32(v3) - 1))
		*(*byte)(unsafe.Add(unsafe.Pointer(v1), 104))--
	}
	v4 = int8(sub_43BE50_get_video_mode_id())
	v5 = int8(*(*byte)(unsafe.Add(unsafe.Pointer(v1), 102)))
	*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v1))), unsafe.Sizeof(uint32(0))*12))) = NOX_CLIENT_VERS_CODE
	*(*byte)(unsafe.Add(unsafe.Pointer(v1), 102)) = byte(int8(int32(v5)&128 | int32(v4)))
	*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v1))), unsafe.Sizeof(uint32(0))*11))) = *memmap.PtrUint32(6112660, 814916)
	*(*uint16)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v1), 109)))) = uint16(int16(nox_xxx_servGetPort_40A430()))
	nox_client_setServerConnectAddr_435720(CString("localhost"))
	if dword_587000_87404 == 1 {
		*(*[268]byte)(unsafe.Pointer(&v10[0])) = [268]byte{}
		v6 = sub_41FA40()
		nox_sprintf(&v9[0], CString("%s%s"), v6, memmap.PtrOff(0x587000, 90752))
		libc.StrCpy(&v10[52], &v9[0])
		*(*uint32)(unsafe.Pointer(&v10[0])) = uint32(sub_420100())
		*(*uint32)(unsafe.Pointer(&v10[4])) = 1
		*(*uint32)(unsafe.Pointer(&v10[8])) = 32
		*(*uint32)(unsafe.Pointer(&v10[12])) = 0
		*(*uint32)(unsafe.Pointer(&v10[16])) = 0
		*(*uint32)(unsafe.Pointer(&v10[20])) = 1
		*(*uint32)(unsafe.Pointer(&v10[24])) = 1
		*(*uint32)(unsafe.Pointer(&v10[44])) = 0
		*(*uint32)(unsafe.Pointer(&v10[28])) = 0
		*(*uint32)(unsafe.Pointer(&v10[224])) = NOX_CLIENT_VERS_CODE
		*(*uint32)(unsafe.Pointer(&v10[32])) = *memmap.PtrUint32(6112660, 814916)
		v10[sub_425550((*uint8)(unsafe.Add(unsafe.Pointer(v1), 100)), (*uint8)(unsafe.Pointer(&v10[69])), 552)+69] = 0
		sub_40D320(int32(uintptr(unsafe.Pointer(&v10[0]))))
	}
	result = nox_xxx_cliGamedataGet_416590(1)
	v8 = int16(int32(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(result))), unsafe.Sizeof(uint16(0))*26)))) & 0xE90F)
	*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v8))), unsafe.Sizeof(int16(0))-1)) |= 1
	*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(result))), unsafe.Sizeof(uint16(0))*26))) = uint16(v8)
	return result
}
func sub_43ACC0() {
	var (
		v0  *byte
		i   int32
		v2  int32
		v4  *byte
		v5  int32
		v6  int32
		v7  *uint32
		v8  *uint32
		v9  [172]byte
		v10 [32]byte
		v11 [32]byte
	)
	if nox_wol_server_result_cnt_815088 == 0 && dword_5d4594_815044 == 0 {
		v0 = sub_41FA40()
		nox_sprintf(&v11[0], CString("%s%s"), v0, memmap.PtrOff(0x587000, 90768))
		for i = sub_41F0E0(); i != 0; i = sub_41F0C0() {
			if libc.StrCmp(&v11[0], (*byte)(unsafe.Pointer(uintptr(i+52)))) != 0 {
				libc.MemSet(unsafe.Pointer(&v9[0]), 0, 168)
				v9[168] = 0
				sub_4255F0((*uint8)(unsafe.Pointer(uintptr(i+69))), (*uint8)(unsafe.Pointer(&v9[100])), 552)
				*(*uint32)(unsafe.Pointer(&v9[96])) = 9999
				if int32(*(*uint32)(unsafe.Pointer(uintptr(i + 40)))) >= 0 {
					*(*uint32)(unsafe.Pointer(&v9[96])) = *(*uint32)(unsafe.Pointer(uintptr(i + 40)))
				}
				v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(i + 32))))
				*(*uint32)(unsafe.Pointer(&v9[44])) = *(*uint32)(unsafe.Pointer(uintptr(i + 32)))
				*(*uint32)(unsafe.Pointer(&v9[48])) = *(*uint32)(unsafe.Pointer(uintptr(i + 224)))
				if *(*int32)(unsafe.Pointer(&dword_587000_87412)) == -1 || sub_437860(int32(int16(v2)), int32(*(*int16)(unsafe.Pointer(&v9[46])))) == *(*int32)(unsafe.Pointer(&dword_587000_87412)) {
					if nox_xxx_checkSomeFlagsOnJoin_4899C0((*nox_gui_server_ent_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(&v9[0]))))))) != 0 {
						v4 = nox_net_ip2str(nox_net_in_addr(cnet.Htonl(*(*uint32)(unsafe.Pointer(uintptr(i + 36))))))
						*(*uint32)(unsafe.Pointer(&v9[12])) = *(*uint32)(unsafe.Pointer(v4))
						*(*uint32)(unsafe.Pointer(&v9[16])) = *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v4))), unsafe.Sizeof(uint32(0))*1)))
						*(*uint32)(unsafe.Pointer(&v9[20])) = *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v4))), unsafe.Sizeof(uint32(0))*2)))
						v5 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v4))), unsafe.Sizeof(uint32(0))*3))))
						*(*uint32)(unsafe.Pointer(&v9[32])) = uint32(i)
						*(*uint32)(unsafe.Pointer(&v9[24])) = uint32(v5)
						*(*uint32)(unsafe.Pointer(&v9[36])) = nox_wol_server_result_cnt_815088
						*(*uint32)(unsafe.Pointer(&v9[44])) = *(*uint32)(unsafe.Pointer(uintptr(i + 32)))
						nox_wol_servers_addResult_4A0030((*nox_gui_server_ent_t)(unsafe.Pointer(&v9[0])))
						if v9[120] != 0 {
							libc.StrNCpy(&v10[0], &v9[120], 15)
							v10[15] = 0
						} else {
							nox_sprintAddrPort_43BC80((*byte)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(&v9[12])))))), *(*uint16)(unsafe.Pointer(&v9[109])), &v10[0])
						}
						if func() uint32 {
							p := &nox_wol_server_result_cnt_815088
							*p++
							return *p
						}() >= 2500 {
							break
						}
					}
				}
			}
		}
		if nox_wol_server_result_cnt_815088 != 0 {
			sub_4A0360()
		}
		sub_44A400()
		if noxflags.HasGame(noxflags.GameFlag26) && *(*int32)(unsafe.Pointer(&dword_587000_87412)) == -1 {
			v6 = sub_4A7F20() + 0x2745
			v7 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&nox_wol_wnd_world_814980)))).ChildByID(0x2712)))
			(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&nox_wol_wnd_world_814980))))).Func94(asWindowEvent(0x4007, int32(uintptr(unsafe.Pointer(v7))), 0))
			v8 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&nox_wol_wnd_world_814980)))).ChildByID(v6)))
			(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&nox_wol_wnd_world_814980))))).Func94(asWindowEvent(0x4007, int32(uintptr(unsafe.Pointer(v8))), 0))
		}
		if nox_game_createOrJoin_815048 != 0 {
			sub_4375C0(0)
		}
	}
}
func sub_43AF30() int32 {
	return int32(dword_5d4594_815052)
}
func sub_43AF40() int32 {
	return int32(nox_game_createOrJoin_815048)
}
func sub_43AF50(a1 int32) int32 {
	var result int32
	result = a1
	dword_587000_87404 = uint32(a1)
	dword_5d4594_2650652 = uint32(bool2int(a1 == 1))
	return result
}
func nox_xxx_check_flag_aaa_43AF70() int32 {
	return int32(dword_587000_87404)
}
func sub_43AF80() int32 {
	return int32(dword_5d4594_814548)
}
func sub_43AF90(a1 int32) int32 {
	var result int32
	result = a1
	dword_5d4594_814548 = uint32(a1)
	return result
}
func nox_client_setConnError_43AFA0(a1 int32) {
	nox_client_connError_814552 = uint32(a1)
	sub_43AF90(2)
}
func sub_43AFC0(a1 int32) {
	var (
		v1  *int32
		v3  int32
		v4  int32
		v5  int32
		v6  int32
		v7  int32
		v8  int32
		v9  int32
		v10 int2
	)
	v1 = sub_4A04F0((*byte)(unsafe.Pointer(uintptr(a1 + 52))))
	if v1 == nil {
		nox_client_setConnError_43AFA0(9)
		return
	}
	if dword_587000_87408 == 0 {
		sub_43A920()
		v10.field_0 = int32(*((*int16)(unsafe.Add(unsafe.Pointer((*int16)(unsafe.Pointer(v1))), unsafe.Sizeof(int16(0))*22)))) + 216
		v10.field_4 = int32(*((*int16)(unsafe.Add(unsafe.Pointer((*int16)(unsafe.Pointer(v1))), unsafe.Sizeof(int16(0))*23)))) + 27
		dword_5d4594_814624 = unsafe.Pointer(v1)
		sub_439370(&v10, int32(uintptr(unsafe.Pointer(v1))))
		nox_client_setMousePos_430B10(v10.field_0, v10.field_4)
		return
	}
	sub_43A920()
	v3 = sub_4A0330(v1)
	(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&nox_wol_wnd_gameList_815012))))).Func94(asWindowEvent(0x4013, v3, 0))
	(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&nox_wol_wnd_gameList_815012))))).Func94(asWindowEvent(0x401C, v3, 0))
	v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_815016 + 32))))
	v10.field_0 = int32(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_815016 + 16))) + 216)
	v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_815016 + 20))))
	v6 = int32(*(*uint32)(unsafe.Pointer(uintptr(v4 + 48))) * 131)
	v7 = int32(*(*uint32)(unsafe.Pointer(uintptr(v4 + 24))))
	v8 = int32(*(*int16)(unsafe.Pointer(uintptr(v4 + 54))))
	v9 = int32(*(*uint32)(unsafe.Pointer(uintptr(v7 + v6*4))))
	dword_5d4594_814624 = unsafe.Pointer(v1)
	v10.field_4 = v9 - v8 + v5 + 27
	sub_439370(&v10, int32(uintptr(unsafe.Pointer(v1))))
	nox_client_setMousePos_430B10(v10.field_0, v10.field_4)
}
func nox_xxx_failconn_43B0E0(a1 int32) {
	var (
		i  int32
		v2 *int32
		v3 int32
		v4 int32
		v5 int32
		v6 int32
	)
	for i = a1; i != 0; i = int32(*(*uint32)(unsafe.Pointer(uintptr(i + 48)))) {
		v2 = sub_4A04F0((*byte)(unsafe.Pointer(uintptr(i + 52))))
		v3 = int32(uintptr(unsafe.Pointer(v2)))
		if v2 != nil {
			if *(*int32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(int32(0))*24)) == 9999 {
				v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(i + 40))))
				if v4 >= 0 {
					*(*uint32)(unsafe.Pointer(uintptr(v3 + 96))) = uint32(v4)
					v5 = sub_4A0330((*int32)(unsafe.Pointer(uintptr(v3))))
					if dword_587000_87408 == 1 {
						if nox_xxx_checkSomeFlagsOnJoin_4899C0((*nox_gui_server_ent_t)(unsafe.Pointer(uintptr(v3)))) != 0 {
							nox_swprintf((*libc.WChar)(memmap.PtrOff(6112660, 814756)), libc.CWString("%d"), *(*uint32)(unsafe.Pointer(uintptr(v3 + 96))))
							(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815028))))).Func94(asWindowEvent(0x4017, int32(uintptr(memmap.PtrOff(6112660, 814756))), v5))
							if nox_xxx_getConnResult_4A0560() == 6 || nox_xxx_getConnResult_4A0560() == 7 {
								sub_425920((**uint32)(unsafe.Pointer(uintptr(v3))))
								v6 = nox_wol_servers_addResult_4A0030((*nox_gui_server_ent_t)(unsafe.Pointer(uintptr(v3))))
								(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&nox_wol_wnd_gameList_815012))))).Func94(asWindowEvent(0x400E, v5, 0))
								sub_43B2A0(v6)
								nox_gui_wol_newServerLine_43B7C0((*nox_gui_server_ent_t)(unsafe.Pointer(uintptr(v3))))
							}
						} else {
							(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&nox_wol_wnd_gameList_815012))))).Func94(asWindowEvent(0x400E, v5, 0))
							sub_4A0540(unsafe.Pointer(uintptr(v3)))
						}
					} else if nox_xxx_checkSomeFlagsOnJoin_4899C0((*nox_gui_server_ent_t)(unsafe.Pointer(uintptr(v3)))) != 0 {
						if int32(*(*uint8)(unsafe.Pointer(uintptr(v3 + 120)))) != 0 {
							libc.StrNCpy((*byte)(memmap.PtrOff(6112660, 814920)), (*byte)(unsafe.Pointer(uintptr(v3+120))), 15)
							*memmap.PtrUint8(6112660, 814935) = 0
						} else {
							nox_sprintAddrPort_43BC80((*byte)(unsafe.Pointer(uintptr(v3+12))), *(*uint16)(unsafe.Pointer(uintptr(v3 + 109))), (*byte)(memmap.PtrOff(6112660, 814920)))
						}
						if *(*uint32)(unsafe.Pointer(uintptr(v3 + 96))) == 9999 {
							nox_swprintf((*libc.WChar)(memmap.PtrOff(6112660, 814628)), libc.CWString("%S -- ms"), memmap.PtrOff(6112660, 814920))
						} else {
							nox_swprintf((*libc.WChar)(memmap.PtrOff(6112660, 814628)), libc.CWString("%S %dms"), memmap.PtrOff(6112660, 814920), *(*uint32)(unsafe.Pointer(uintptr(v3 + 96))))
						}
						nox_xxx_wndWddSetTooltip_46B000((*nox_window_data)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v3 + 28)))+36))), (*libc.WChar)(memmap.PtrOff(6112660, 814628)))
						sub_437320(v3)
					} else {
						sub_4A0540(unsafe.Pointer(uintptr(v3)))
					}
				}
			}
		}
	}
}
func sub_43B2A0(a1 int32) int32 {
	var result int32
	result = a1
	*(*uint16)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(nox_wol_wnd_gameList_815012 + 32))) + 46))) = uint16(int16(a1))
	*(*uint16)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_815016 + 32))) + 46))) = uint16(int16(a1))
	*(*uint16)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_815020 + 32))) + 46))) = uint16(int16(a1))
	*(*uint16)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_815024 + 32))) + 46))) = uint16(int16(a1))
	*(*uint16)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_815028 + 32))) + 46))) = uint16(int16(a1))
	*(*uint16)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_815032 + 32))) + 46))) = uint16(int16(a1))
	return result
}
func nox_client_getServerAddr_43B300() uint32 {
	var result uint32
	if dword_5d4594_815056 != 0 {
		result = uint32(cnet.ParseAddr(libc.GoString((*byte)(unsafe.Pointer(uintptr(uint32(uintptr(dword_5d4594_814624)) + 12))))))
	} else {
		result = 0
	}
	return result
}
func nox_client_getServerPort_43B320() int32 {
	if dword_5d4594_815056 != 0 {
		return int32(*memmap.PtrUint32(6112660, 814604))
	}
	return 0
}
func sub_43B340() int32 {
	var result int32
	if dword_5d4594_815056 != 0 {
		result = int32(*(*uint16)(unsafe.Pointer(uintptr(uint32(uintptr(dword_5d4594_814624)) + 163))))
	} else {
		result = 0
	}
	return result
}
func sub_43B460() int32 {
	sub_438370()
	nox_wnd_xxx_815040.fnc_done_out = sub_43B490
	nox_xxx_wnd_46C6E0((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815000))))))
	return nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_814984))))), 0)
}
func sub_43B490() int32 {
	if gameGetStateCode() == 1700 {
		return sub_438330()
	}
	(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&nox_wol_wnd_world_814980))))).Hide()
	(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815000))))).Hide()
	nox_client_setCursorType_477610(0)
	return 1
}
