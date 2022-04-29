package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"math"
	"unsafe"
)

type nox_light_3 struct {
	r int32
	g int32
	b int32
}

func sub_484F90(a1 int32) {
	var (
		v1  int32
		v2  int32
		v3  float64
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
		v15 float64
		v16 *uint8
		v17 int32
		v18 float64
		v19 float64
		v20 float64
		v21 float64
		v22 int32
		v23 int32
		v24 int32
		v25 int32
		v26 int32
		v27 int32
		v28 float32
		v29 float32
		v30 float32
		v31 float32
		a5  float32
		a5a int32
		v34 int32
		v35 int2
		a3  int2
		v37 float32
		v38 float32
		v39 float32
		a1a int2
		a4  int32
		v42 int32
		a2  int2
		v44 int32
		v45 int32
	)
	v1 = a1
	if sub_45A840((*uint32)(unsafe.Pointer(uintptr(a1)))) != 0 || *(*uint32)(unsafe.Pointer(uintptr(a1 + 112)))&0x80000 != 0 && (func() uint32 {
		v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 120))))
		return uint32(v2) & 0x1000000
	}()) != 0 && v2&4 != 0 && *(*uint32)(unsafe.Pointer(uintptr(a1 + 144))) > 0 {
		if nox_xxx_get_57AF20() == 0 || uint32(a1) == *memmap.PtrUint32(0x852978, 8) || cgoFuncAddr(cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(a1 + 300))), (*func(*int32, int32) int32)(nil))) == cgoFuncAddr(nox_thing_glow_orb_draw) {
			if *(*uint32)(unsafe.Pointer(uintptr(a1 + 120)))&0x20000000 != 0 && nox_video_modeXxx_3801780 == 1 {
				a5 = float32(nox_common_randomFloatXxx_416090(0.89999998, 1.1) * float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 140)))))
				v3 = float64(sub_484C60(a5))
			} else {
				v3 = float64(*(*int32)(unsafe.Pointer(uintptr(a1 + 144))))
				a5 = *(*float32)(unsafe.Pointer(uintptr(a1 + 140)))
			}
			*(*float32)(unsafe.Pointer(&v44)) = float32(v3)
			if float64(a5) > float64(*(*float32)(unsafe.Pointer(&dword_587000_154968))) {
				if float64(a5) <= 31.0 {
					v37 = a5
				} else {
					v37 = 31.0
				}
				v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 16))))
				v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 168))))
				v6 = int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 12))) - dword_5d4594_2650676)
				v39 = a5 * a5
				v7 = int32(uint32(v4) - dword_5d4594_2650680)
				a4 = v6
				v42 = v7
				if v5 == math.MaxUint16 {
					v38 = *(*float32)(unsafe.Pointer(&v44)) * *(*float32)(unsafe.Pointer(&v44))
					v28 = float32((float64(a4)-float64(*(*float32)(unsafe.Pointer(&v44))))*0.043478262 + 0.5)
					v8 = int(v28)
					a5a = v8
					if v8 < 0 {
						a5a = 0
						v8 = 0
					}
					v29 = float32((float64(a4)+float64(*(*float32)(unsafe.Pointer(&v44))))*0.043478262 + 0.5)
					v9 = int(v29)
					v34 = v9
					if v9 > 56 {
						v34 = 56
						v9 = 56
					}
					v30 = float32((float64(v42)-float64(*(*float32)(unsafe.Pointer(&v44))))*0.043478262 + 0.5)
					v10 = int(v30)
					if v10 < 0 {
						v10 = 0
					}
					v31 = float32((float64(v42)+float64(*(*float32)(unsafe.Pointer(&v44))))*0.043478262 + 0.5)
					v11 = int(v31)
					v35.field_0 = v11
					if v11 > 44 {
						v35.field_0 = 44
						v11 = 44
					}
					a3.field_0 = v10
					v12 = v10 * 23
					a2.field_4 = v10 * 23
					if v10 <= v11 {
						v13 = v8 * 23
						a1a.field_0 = v8 * 23
						for {
							if v8 <= v9 {
								v14 = a4 - v13
								v45 = a4 - v13
								v15 = float64(v42-v12) * float64(v42-v12)
								v16 = (*uint8)(memmap.PtrOff(0x8531A0, uint32(v8*+40+v8*5+a3.field_0)))
								v17 = v9 - a5a + 1
								for {
									v18 = float64(v45)*float64(v45) + v15
									if v18 <= float64(v38) {
										v19 = float64(*mem_getFloatPtr(0x587000, 0x25D74)+v37) / (float64(*mem_getFloatPtr(0x587000, 0x25D70)**mem_getFloatPtr(0x587000, 0x25D5C))*v18/float64(v39) + 1.0)
										if v19 > float64(*mem_getFloatPtr(0x587000, 0x25D74)) {
											v20 = v19 - float64(*mem_getFloatPtr(0x587000, 0x25D74))
											if *(*uint32)(unsafe.Pointer(uintptr(v1 + 172))) != 0 {
												v21 = float64(*v16) - v20
											} else {
												v21 = v20 + float64(*v16)
											}
											if v21 >= 0.0 {
												if v21 > 31.0 {
													v21 = 31.0
												}
											} else {
												v21 = 0.0
											}
											v12 = a2.field_4
											v13 = a1a.field_0
											v14 -= 23
											*v16 = uint8(int8(int64(v21)))
											v45 = v14
										} else {
											v14 -= 23
											v45 = v14
										}
									} else {
										v14 -= 23
										v45 = v14
									}
									v16 = (*uint8)(unsafe.Add(unsafe.Pointer(v16), 45))
									v17--
									if v17 == 0 {
										break
									}
								}
								v8 = a5a
								v9 = v34
								v11 = v35.field_0
							}
							v12 += 23
							a2.field_4 = v12
							a3.field_0++
							if a3.field_0 > v11 {
								break
							}
						}
					}
				} else {
					a1a.field_0 = (v6 << 16) / 23
					a1a.field_4 = (v7 << 16) / 23
					v22 = int32(int64(*(*float32)(unsafe.Pointer(&v44)))<<16) / 23
					v23 = sub_4C1C60(v22, int32(*memmap.PtrUint32(0x85B3FC, uint32((int32(uint16(int16(int32(*(*uint16)(unsafe.Pointer(uintptr(v1 + 164))))+0x4000)))>>4)*4+12260))*16))
					v24 = sub_4C1C60(v22, int32(*memmap.PtrUint32(0x85B3FC, uint32((int32(*(*uint16)(unsafe.Pointer(uintptr(v1 + 164))))>>4)*4+12260))*16))
					a3.field_0 = a1a.field_0 + v23
					*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v23))), unsafe.Sizeof(uint16(0))*0)) = uint16(int16(int32(*(*uint16)(unsafe.Pointer(uintptr(v1 + 164)))) + int32(*(*uint16)(unsafe.Pointer(uintptr(v1 + 166))))))
					a3.field_4 = a1a.field_4 + v24
					v25 = sub_4C1C60(v22, int32(*memmap.PtrUint32(0x85B3FC, uint32((int32(uint16(int16(v23+0x4000)))>>4)*4+12260))*16))
					v26 = sub_4C1C60(v22, int32(*memmap.PtrUint32(0x85B3FC, uint32((int32(uint16(int16(v23)))>>4)*4+12260))*16))
					*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v23))), unsafe.Sizeof(uint16(0))*0)) = uint16(int16(int32(*(*uint16)(unsafe.Pointer(uintptr(v1 + 164)))) - int32(*(*uint16)(unsafe.Pointer(uintptr(v1 + 166))))))
					a2.field_4 = a1a.field_4 + v26
					a2.field_0 = a1a.field_0 + v25
					v27 = sub_4C1C60(v22, int32(*memmap.PtrUint32(0x85B3FC, uint32((int32(uint16(int16(v23+0x4000)))>>4)*4+12260))*16))
					v35.field_4 = a1a.field_4 + sub_4C1C60(v22, int32(*memmap.PtrUint32(0x85B3FC, uint32((int32(uint16(int16(v23)))>>4)*4+12260))*16))
					v35.field_0 = a1a.field_0 + v27
					sub_4854D0(&a1a, &a2, &a3, int32(uintptr(unsafe.Pointer(&a4))), a5)
					sub_4854D0(&a1a, &a3, &v35, int32(uintptr(unsafe.Pointer(&a4))), a5)
				}
			}
		}
	}
}
func sub_484BD0() float64 {
	return float64(*(*float32)(unsafe.Pointer(&dword_587000_154968)))
}
func sub_4C1C70(a1 int32, a2 int32) int32 {
	return int32((int64(a1) << 16) / int64(a2))
}
func sub_4697C0(a1 *int32, a2 *int32, a3 *int2, a4 int32, a5 *int32) int32 {
	var (
		result int32
		v6     *int32
		v7     int32
		v8     int32
		v9     int32
		v10    *int32
		v11    *byte
		v12    int32
		v13    int32
		v14    int32
		v15    int32
		v16    int32
		v17    uint32
		v18    int32
		i      *byte
		j      int32
		v21    uint32
		v22    *int32
		v23    int32
		v24    int32
	)
	result = *a1
	v6 = a2
	v7 = *a1 + *(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*1))
	v18 = *a1 + *(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*1))
	if *a1 < *a2 {
		result = *a2
	}
	if v7 > *(*int32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(int32(0))*1))+*a2 {
		v18 = *(*int32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(int32(0))*1)) + *a2
		v7 = *(*int32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(int32(0))*1)) + *a2
	}
	if v7-result > 1 {
		v23 = int32(*memmap.PtrUint32(0x587000, 142308))
		if uint32(a4) <= uint32(*memmap.PtrInt32(0x587000, 142308)) {
			v23 = a4
		}
		v21 = uint32((a4 >> 16) * (a4 >> 16))
		v8 = result*23 - a3.field_4
		v24 = result*23 - a3.field_4
		v9 = result
		if result < v7 {
			v10 = (*int32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(int32(0))*uintptr(result+2)))
			v11 = (*byte)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer((*byte)(unsafe.Pointer(a1)))) - uintptr(unsafe.Pointer((*byte)(unsafe.Pointer(v6))))))))
			v12 = int32(dword_587000_142316)
			v22 = v10
			for i = v11; ; v11 = i {
				v13 = *(*int32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(v10))), uint32(uintptr(unsafe.Pointer(v11)))))))
				v14 = v13*23 - a3.field_0
				v15 = v13
				if v13 < *v10 {
					v16 = v8 * v8
					for j = v8 * v8; ; v16 = j {
						v17 = uint32(sub_4C1C70(v12+v23, int32(*memmap.PtrUint32(0x587000, 142312)*66*uint32(v16+v14*v14)/v21+0x10000)))
						v12 = int32(dword_587000_142316)
						if v17 > uint32(*(*int32)(unsafe.Pointer(&dword_587000_142316))) {
							sub_4695E0(v15, v9, a5, int32((v17-dword_587000_142316)*8), 0)
							v12 = int32(dword_587000_142316)
						}
						v14 += 23
						if func() int32 {
							p := &v15
							*p++
							return *p
						}() >= *v22 {
							break
						}
					}
					v8 = v24
					v10 = v22
				}
				result = v18
				v8 += 23
				v9++
				v10 = (*int32)(unsafe.Add(unsafe.Pointer(v10), unsafe.Sizeof(int32(0))*1))
				v24 = v8
				v22 = v10
				if v9 >= v18 {
					break
				}
			}
		}
	}
	return result
}
func sub_4696B0(a1 *int2, a2 *int2, a3 *int2, a4 *int2, a5 int32, a6 *int32) int32 {
	var (
		v6     int32
		v7     int32
		v8     *int2
		v9     *int2
		v10    *int2
		result int32
		v12    int32
		v13    [184]byte
		v14    [184]byte
		v15    [184]byte
	)
	v6 = a1.field_4
	v7 = a2.field_4
	if v6 > v7 {
		v10 = a3
		if v7 <= a3.field_4 {
			v9 = a2
			v8 = a1
			goto LABEL_9
		}
		v9 = a3
	} else {
		v8 = a3
		if v6 <= a3.field_4 {
			v9 = a1
			v10 = a2
			goto LABEL_9
		}
		v9 = a3
	}
	v10 = a1
	v8 = a2
LABEL_9:
	sub_484DC0(v9, v10, (*int32)(unsafe.Pointer(&v15[0])))
	sub_484DC0(v9, v8, (*int32)(unsafe.Pointer(&v14[0])))
	sub_4697C0((*int32)(unsafe.Pointer(&v15[0])), (*int32)(unsafe.Pointer(&v14[0])), a4, a5, a6)
	result = v10.field_4
	v12 = v8.field_4
	if result >= v12 {
		if result > v12 {
			sub_484DC0(v8, v10, (*int32)(unsafe.Pointer(&v13[0])))
			result = sub_4697C0((*int32)(unsafe.Pointer(&v15[0])), (*int32)(unsafe.Pointer(&v13[0])), a4, a5, a6)
		}
	} else {
		sub_484DC0(v10, v8, (*int32)(unsafe.Pointer(&v13[0])))
		result = sub_4697C0((*int32)(unsafe.Pointer(&v13[0])), (*int32)(unsafe.Pointer(&v14[0])), a4, a5, a6)
	}
	return result
}
func nox_xxx_cliLight16_469140(dr *nox_drawable, vp *nox_draw_viewport_t) {
	var (
		arg0 int32 = int32(uintptr(unsafe.Pointer(dr)))
		v1   int32
		v2   int32
		v3   int32
		v4   int32
		v5   int32
		v6   int32
		v7   int32
		v16  uint32
		v17  int16
		v18  int32
		v19  int32
		v20  int32
		v21  int32
		v22  int32
		v23  int32
		v24  int32
		v25  int32
		v26  *int32
		v27  float32
		v33  int32
		v36  uint32
		a1   int2
		a4   int2
		a3   int2
		a2   int2
		v42  int2
		v44  int32
		v45  int32
	)
	v1 = arg0
	if !(sub_45A840((*uint32)(unsafe.Pointer(uintptr(arg0)))) != 0 || *(*uint32)(unsafe.Pointer(uintptr(arg0 + 112)))&0x80000 != 0 && (func() uint32 {
		v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(arg0 + 120))))
		return uint32(v2) & 0x1000000
	}()) != 0 && *(*uint32)(unsafe.Pointer(uintptr(arg0 + 144))) > 0 && v2&4 != 0) {
		return
	}
	if !(nox_xxx_get_57AF20() == 0 || uint32(arg0) == *memmap.PtrUint32(0x852978, 8) || cgoFuncAddr(cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(arg0 + 300))), (*func(*int32, int32) int32)(nil))) == cgoFuncAddr(nox_thing_glow_orb_draw)) {
		return
	}
	v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(arg0 + 148))))
	v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(arg0 + 144))))
	if *(*uint32)(unsafe.Pointer(uintptr(arg0 + 120)))&0x20000000 != 0 {
		v3 += randomIntMinMax(0, int32(*(*uint32)(unsafe.Pointer(uintptr(arg0 + 148)))>>18)) << 16
		v27 = float32(float64(v3) * 1.5258789e-05)
		v4 = sub_484C60(v27)
	}
	if v3 <= int32(sub_484BD0()*65536.0) {
		return
	}
	v33 = v3
	if uint32(v3) > uint32(*memmap.PtrInt32(0x587000, 142320)) {
		v33 = int32(*memmap.PtrUint32(0x587000, 142320))
	}
	v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(arg0 + 12))) - dword_5d4594_2650676)
	v7 = int32(*(*uint32)(unsafe.Pointer(uintptr(arg0 + 16))) - dword_5d4594_2650680)
	v36 = uint32((v3 >> 16) * (v3 >> 16))
	v6 = int32(*(*uint32)(unsafe.Pointer(uintptr(arg0 + 168))))
	a4.field_0 = v5
	a4.field_4 = v7
	if v6 == math.MaxUint16 {
		var (
			dlimit uint32 = uint32(v4 * v4)
			xmin   int32  = (v5 - v4) / 23
		)
		if xmin < 0 {
			xmin = 0
		}
		var xmax int32 = (v5 + v4) / 23
		if xmax > 56 {
			xmax = 56
		}
		var ymin int32 = (v7 - v4) / 23
		if ymin < 0 {
			ymin = 0
		}
		var ymax int32 = (v7 + v4) / 23
		if ymax > 44 {
			ymax = 44
		}
		var v11 int32 = int32(dword_587000_142328)
		for y := int32(ymin); y <= ymax; y++ {
			var (
				dy  int32 = v7 - y*23
				dy2 int32 = dy * dy
			)
			for x := int32(xmin); x <= xmax; x++ {
				var (
					dx   int32  = v5 - x*23
					dx2  int32  = dx * dx
					dist uint32 = uint32(dx2 + dy2)
				)
				if dist <= dlimit {
					v16 = uint32(sub_4C1C70(v33+v11, int32(dist*66**memmap.PtrUint32(0x587000, 142324)/v36+0x10000)))
					v11 = int32(dword_587000_142328)
					if v16 > uint32(v11) {
						sub_4695E0(x, y, (*int32)(unsafe.Pointer(uintptr(v1+152))), int32((v16-dword_587000_142328)*8), int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 172)))))
					}
					v5 = a4.field_0
					v7 = a4.field_4
				}
			}
		}
	} else {
		a1.field_0 = (v5 << 16) / 23
		v17 = int16(*(*uint16)(unsafe.Pointer(uintptr(arg0 + 164))))
		v18 = v4 << 16
		a1.field_4 = (v7 << 16) / 23
		v19 = v18 / 23
		v20 = sub_4C1C60(v18/23, int32(*memmap.PtrUint32(0x85B3FC, uint32((int32(uint16(int16(int32(v17)+0x4000)))>>4)*4+12260))*16))
		v21 = v18 / 23
		v22 = v20
		v23 = sub_4C1C60(v21, int32(*memmap.PtrUint32(0x85B3FC, uint32((int32(*(*uint16)(unsafe.Pointer(uintptr(arg0 + 164))))>>4)*4+12260))*16))
		a3.field_0 = a1.field_0 + v22
		*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v22))), unsafe.Sizeof(uint16(0))*0)) = uint16(int16(int32(*(*uint16)(unsafe.Pointer(uintptr(arg0 + 164)))) + int32(*(*uint16)(unsafe.Pointer(uintptr(arg0 + 166))))))
		a3.field_4 = a1.field_4 + v23
		v44 = sub_4C1C60(v19, int32(*memmap.PtrUint32(0x85B3FC, uint32((int32(uint16(int16(v22+0x4000)))>>4)*4+12260))*16))
		v24 = sub_4C1C60(v19, int32(*memmap.PtrUint32(0x85B3FC, uint32((int32(uint16(int16(v22)))>>4)*4+12260))*16))
		*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v22))), unsafe.Sizeof(uint16(0))*0)) = uint16(int16(int32(*(*uint16)(unsafe.Pointer(uintptr(v1 + 164)))) - int32(*(*uint16)(unsafe.Pointer(uintptr(v1 + 166))))))
		a2.field_0 = a1.field_0 + v44
		a2.field_4 = a1.field_4 + v24
		v45 = sub_4C1C60(v19, int32(*memmap.PtrUint32(0x85B3FC, uint32((int32(uint16(int16(v22+0x4000)))>>4)*4+12260))*16))
		v25 = sub_4C1C60(v19, int32(*memmap.PtrUint32(0x85B3FC, uint32((int32(uint16(int16(v22)))>>4)*4+12260))*16))
		v26 = (*int32)(unsafe.Pointer(uintptr(v1 + 152)))
		v42.field_0 = a1.field_0 + v45
		v42.field_4 = a1.field_4 + v25
		sub_4696B0(&a1, &a2, &a3, &a4, v3, v26)
		sub_4696B0(&a1, &a3, &v42, &a4, v3, v26)
	}
}
func nox_xxx___cfltcvt_init_430CC0() int32 {
	*memmap.PtrUint32(0x973F18, 7696) = 1
	*memmap.PtrUint32(0x973F18, 7700) = uint32(cgoFuncAddr(nox_xxx_someEdgeProcessing_480EF0))
	*memmap.PtrUint32(0x973F18, 7720) = uint32(cgoFuncAddr(sub_4814F0))
	dword_5d4594_3807156 = uint32(cgoFuncAddr(sub_469920))
	dword_5d4594_805836 = 0
	dword_5d4594_3805484 = uint32(cgoFuncAddr(sub_480BE0))
	dword_5d4594_3805492 = uint32(cgoFuncAddr(sub_480860))
	return 0
}
