package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	"github.com/gotranspile/cxgo/runtime/stdio"
	"github.com/noxworld-dev/opennox/v1/common/alloc"
	noxflags "github.com/noxworld-dev/opennox/v1/common/flags"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"math"
	"unsafe"
)

var nox_server_object_updateable_1556852 *nox_object_t = nil
var nox_server_objects_uninited_1556860 *nox_object_t = nil
var nox_server_objects_updatable2_1556848 *nox_object_t = nil
var nox_server_objects_1556844 *nox_object_t = nil
var nox_xxx_host_player_unit_3843628 *nox_object_t = nil
var nox_common_maplist nox_list_item_t = nox_list_item_t{}

func nox_xxx_updDrawDBall_4CDF80(a1 int32, a2 int32) int32 {
	nox_xxx_updDrawAddRndSpark_4CDFA0(a2, (*uint32)(unsafe.Pointer(uintptr(3))))
	return 1
}
func sub_4CE0A0(a1 int32, a2 int32) int32 {
	nox_xxx_updDrawAddRndSpark_4CDFA0(a2, (*uint32)(unsafe.Pointer(uintptr(1))))
	return 1
}
func nox_xxx_updDrawCloud_4CE1D0(a1 int32, a2 int32) int32 {
	if int32(uint8(nox_frame_xxx_2598000))&1 != 0 {
		sub_4CE200(a1, a2, 1, 75)
	}
	return 1
}
func sub_4CE340(a1 int32, a2 int32) int32 {
	*(*uint16)(unsafe.Pointer(uintptr(a2 + 104))) += uint16(*(*uint8)(unsafe.Pointer(uintptr(a2 + 432))))
	return 1
}
func sub_4CE360(a1 int32, a2 int32) int32 {
	if int32(uint8(nox_frame_xxx_2598000))&1 != 0 {
		sub_4CE200(a1, a2, 1, 35)
	}
	return 1
}
func nox_xxx_updDrawColorlight_4CE390(a1 *uint32, a2 int32) int32 {
	var (
		result int32
		v3     int32
		v4     int32
		v5     int32
		v6     int32
	)
	if int32(*(*uint8)(unsafe.Pointer(uintptr(a2 + 432)))) != 0 || int32(*(*uint8)(unsafe.Pointer(uintptr(a2 + 433)))) != 0 || int32(*(*uint8)(unsafe.Pointer(uintptr(a2 + 434)))) != 0 {
		v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 12))))
		v4 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*4)))
		if v3 >= v4-100 && uint32(v3) <= *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*8))+uint32(v4)+100 {
			v5 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*5)))
			v6 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 16))))
			if v6 >= v5-100 && uint32(v6) <= *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*9))+uint32(v5)+100 && !noxflags.HasGame(noxflags.GameFlag22) {
				nox_xxx_colorLightAlterB0ArrayColor_4CE440(a2)
				nox_xxx_colorLightAlterIntensity_4CE610(a2)
				nox_xxx_colorLightAlterRadius_4CE760(a2)
				sub_4CE960(a2)
				sub_4CE8C0(a2)
			}
		}
		result = 1
	} else {
		sub_49BCD0((*nox_drawable)(unsafe.Pointer(uintptr(a2))))
		result = 1
	}
	return result
}
func nox_xxx_colorLightAlterB0ArrayColor_4CE440(a1 int32) {
	var (
		v1  int32
		v2  int8
		v3  float64
		v4  int64
		v5  float64
		v6  float64
		v7  int8
		v8  int8
		v9  int32
		v10 float64
		v11 int32
		v12 int8
		v13 float32
		v14 int32
		v15 int8
	)
	v1 = a1
	v2 = int8(*(*uint8)(unsafe.Pointer(uintptr(a1 + 432))))
	v12 = int8(*(*uint8)(unsafe.Pointer(uintptr(a1 + 432))))
	if int32(v2) > 1 && int32(*(*uint16)(unsafe.Pointer(uintptr(a1 + 258)))) != 0 {
		v3 = float64(nox_frame_xxx_2598000) / float64(int32(*(*uint16)(unsafe.Pointer(uintptr(a1 + 258))))*int32(v2))
		v4 = int64(v3)
		v5 = v3 - float64(int32(int64(v3)))
		if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 176))))&1 != 0 {
			v6 = v5 * float64(v2)
			v15 = int8(int64(v6))
			v7 = int8(int32(v15) + 1)
			if int32(int8(int32(v15)+1)) >= int32(v12) {
				v7 = 0
			}
		} else {
			v6 = v5 * float64(v2)
			if v4&1 != 0 {
				v8 = int8(uint8(uint64(v2) - uint64(int64(v6)) - 1))
				v15 = v8
				v7 = int8(int32(v8) - 1)
				if int32(v7) < 0 {
					v7 = 0
				}
			} else {
				v15 = int8(int64(v6))
				v7 = int8(int32(v15) + 1)
				if int32(int8(int32(v15)+1)) >= int32(v12) {
					v7 = int8(int32(v12) - 1)
				}
			}
		}
		v13 = float32(float64(*(*uint16)(unsafe.Pointer(uintptr(v1 + 258)))))
		v9 = int32(v7)
		v10 = float64(uint8(int8(int64((v6 - float64(int32(int64(v6)))) * float64(v13)))))
		v11 = v1 + int32(v15)*2 + int32(v15)
		v14 = int32(*(*uint8)(unsafe.Pointer(uintptr(int32(v15) + v1 + int32(v15)*2 + 120 + 60))))
		nox_xxx_spriteChangeLightColor_484BE0((*uint32)(unsafe.Pointer(uintptr(v1+136))), int32(int64(float64(int32(*(*uint8)(unsafe.Pointer(uintptr(v9 + v1 + v9*2 + 178))))-int32(*(*uint8)(unsafe.Pointer(uintptr(v11 + 178)))))/float64(v13)*v10+float64(*(*uint8)(unsafe.Pointer(uintptr(v11 + 178)))))), int32(int64(float64(int32(*(*uint8)(unsafe.Pointer(uintptr(v9 + v1 + v9*2 + 179))))-int32(*(*uint8)(unsafe.Pointer(uintptr(v11 + 179)))))/float64(v13)*v10+float64(*(*uint8)(unsafe.Pointer(uintptr(v11 + 179)))))), int32(int64(float64(int32(*(*uint8)(unsafe.Pointer(uintptr(v9 + v1 + v9*2 + 120 + 60))))-v14)/float64(v13)*v10+float64(v14))))
	}
}
func nox_xxx_colorLightAlterIntensity_4CE610(a1 int32) {
	var (
		v1  int32
		v2  float64
		v3  int64
		v4  float64
		v5  float64
		v6  int64
		v7  int8
		v8  float32
		v9  int8
		v10 float32
	)
	v1 = a1
	v9 = int8(*(*uint8)(unsafe.Pointer(uintptr(a1 + 433))))
	if int32(v9) > 1 && int32(*(*uint16)(unsafe.Pointer(uintptr(v1 + 260)))) != 0 {
		v2 = float64(nox_frame_xxx_2598000) / float64(int32(*(*uint16)(unsafe.Pointer(uintptr(v1 + 260))))*int32(v9))
		v3 = int64(v2)
		v4 = v2 - float64(int32(int64(v2)))
		if int32(*(*uint8)(unsafe.Pointer(uintptr(v1 + 176))))&4 != 0 {
			v5 = v4 * float64(v9)
			v6 = int64(v5)
			v7 = int8(uint8(uint64(int64(v5)) + 1))
			if int32(v7) >= int32(v9) {
				v7 = 0
			}
		} else {
			v5 = v4 * float64(v9)
			if v3&1 != 0 {
				*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v6))), 0)) = uint8(uint64(v9) - uint64(int64(v5)) - 1)
				v7 = int8(uint8(uint64(v9) - uint64(int64(v5)) - 2))
				if int32(v7) < 0 {
					v7 = 0
				}
			} else {
				v6 = int64(v5)
				v7 = int8(uint8(uint64(int64(v5)) + 1))
				if int32(v7) >= int32(v9) {
					v7 = int8(int32(v9) - 1)
				}
			}
		}
		v10 = float32(float64(*(*uint16)(unsafe.Pointer(uintptr(v1 + 260)))))
		v8 = float32(float64(uint8(int8(int64((v5-float64(int32(int64(v5))))*float64(v10)))))*(float64(int32(*(*uint8)(unsafe.Pointer(uintptr(int32(v7) + v1 + 226))))-int32(*(*uint8)(unsafe.Pointer(uintptr(int32(int8(v6)) + v1 + 226)))))/float64(v10)) + float64(*(*uint8)(unsafe.Pointer(uintptr(int32(int8(v6)) + v1 + 226)))))
		nox_xxx_spriteChangeIntensity_484D70_light_intensity(v1+136, v8)
	}
}
func nox_xxx_colorLightAlterRadius_4CE760(a1 int32) int32 {
	var (
		v1     int32
		result int32
		v3     int8
		v4     float64
		v5     int64
		v6     float64
		v7     float64
		v8     int64
		v9     int8
		v10    int8
		v11    int8
		v12    float32
	)
	v1 = a1
	result = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 168))))
	v3 = int8(*(*uint8)(unsafe.Pointer(uintptr(a1 + 434))))
	v11 = int8(*(*uint8)(unsafe.Pointer(uintptr(a1 + 434))))
	if result == 0 && int32(*(*uint16)(unsafe.Pointer(uintptr(v1 + 262)))) > 0 && int32(v3) > 1 {
		v4 = float64(nox_frame_xxx_2598000) / float64(int32(*(*uint16)(unsafe.Pointer(uintptr(v1 + 262))))*int32(v3))
		v5 = int64(v4)
		v6 = v4 - float64(int32(int64(v4)))
		if int32(*(*uint8)(unsafe.Pointer(uintptr(v1 + 176))))&16 != 0 {
			v7 = v6 * float64(v3)
			v8 = int64(v7)
			v9 = int8(uint8(uint64(int64(v7)) + 1))
			if int32(v9) >= int32(v11) {
				v9 = 0
			}
		} else {
			v7 = v6 * float64(v3)
			if v5&1 != 0 {
				v10 = int8(uint8(uint64(v3) - uint64(int64(v7)) - 1))
				*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v8))), 0)) = uint8(v10)
				v9 = int8(int32(v10) - 1)
				if int32(v9) < 0 {
					v9 = 0
				}
			} else {
				v8 = int64(v7)
				v9 = int8(uint8(uint64(int64(v7)) + 1))
				if int32(v9) >= int32(v11) {
					v9 = int8(int32(v11) - 1)
				}
			}
		}
		v12 = float32(float64(*(*uint16)(unsafe.Pointer(uintptr(v1 + 262)))))
		result = int32(nox_xxx_spriteChangeLightSize_484C30(v1+136, int32(int64(float64(uint8(int8(int64((v7-float64(int32(int64(v7))))*float64(v12)))))*(float64(int32(*(*uint8)(unsafe.Pointer(uintptr(int32(v9) + v1 + 242))))-int32(*(*uint8)(unsafe.Pointer(uintptr(int32(int8(v8)) + v1 + 242)))))/float64(v12))+float64(*(*uint8)(unsafe.Pointer(uintptr(int32(int8(v8)) + v1 + 242))))))))
	}
	return result
}
func sub_4CE8C0(a1 int32) {
	var (
		v1 *uint32
		v2 int32
		v3 float64
		v4 int32
	)
	if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 176))))&64 != 0 {
		v1 = nox_xxx_netSpriteByCodeStatic_45A720(int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 264)))))
		if v1 != nil {
			v2 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*4)) - *(*uint32)(unsafe.Pointer(uintptr(a1 + 16))))
			v4 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*3)) - *(*uint32)(unsafe.Pointer(uintptr(a1 + 12))))
			v3 = float64(nox_double2float(math.Acos(float64(v4)/math.Sqrt(float64(v4*v4+v2*v2))))) * 57.295776
			if v2 < 0 {
				v3 = 360.0 - v3
			}
			sub_484C00(a1+136, int32(int64(v3)))
		}
	}
}
func sub_4CE960(a1 int32) {
	var (
		v2 int16
		v3 float64
		v4 float64
		v5 float64
	)
	if *(*uint32)(unsafe.Pointer(uintptr(a1 + 168))) == 0 {
		return
	}
	v2 = int16(*(*uint16)(unsafe.Pointer(uintptr(a1 + 176))))
	if (int32(v2) & 128) == 0 {
		return
	}
	var v1 int16 = int16(*(*uint16)(unsafe.Pointer(uintptr(a1 + 270))))
	if int32(v1) == 0 {
		return
	}
	if (int32(v2) & 256) == 256 {
		v5 = float64(int32(*(*uint16)(unsafe.Pointer(uintptr(a1 + 272)))) - int32(*(*uint16)(unsafe.Pointer(uintptr(a1 + 268)))))
	} else {
		v5 = 360.0
	}
	v4 = float64(nox_gameFPS)
	v3 = float64(v1)
	var v6 int32 = int32((float64(nox_frame_xxx_2598000)/(v5/v3*v4)-float64(int32(int64(float64(nox_frame_xxx_2598000)/(v5/v3*v4)))))*(v5/v3*v4)*(v3/v4) + float64(*(*uint16)(unsafe.Pointer(uintptr(a1 + 268)))))
	if v6 >= 0 {
		if v6 >= 360 {
			v6 -= 360
		}
	} else {
		v6 += 360
	}
	sub_484C00(a1+136, v6)
}
func sub_4CEA90(a1 *float32, a2 *int2, a3 int32) int32 {
	var (
		v3  int32
		v4  int32
		v5  int32
		v7  int32
		v8  int32
		v9  int32
		v10 int32
		v11 int32
	)
	v3 = a2.field_0 + int(*a1)
	v4 = a2.field_4 + int(*(*float32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(float32(0))*1)))
	v9 = a2.field_0 + int(*(*float32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(float32(0))*2)))
	v10 = a2.field_4 + int(*(*float32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(float32(0))*3)))
	v7 = a2.field_0 + int(*(*float32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(float32(0))*4)))
	v8 = a2.field_4 + int(*(*float32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(float32(0))*5)))
	v11 = a2.field_0 + int(*(*float32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(float32(0))*6)))
	v5 = a2.field_4 + int(*(*float32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(float32(0))*7)))
	nox_client_drawSetColor_434460(a3)
	nox_client_drawAddPoint_49F500(v3, v4)
	nox_client_drawAddPoint_49F500(v7, v8)
	nox_client_drawLineFromPoints_49E4B0()
	nox_client_drawAddPoint_49F500(v11, v5)
	nox_client_drawAddPoint_49F500(v7, v8)
	nox_client_drawLineFromPoints_49E4B0()
	nox_client_drawAddPoint_49F500(v3, v4)
	nox_client_drawAddPoint_49F500(v9, v10)
	nox_client_drawLineFromPoints_49E4B0()
	nox_client_drawAddPoint_49F500(v11, v5)
	nox_client_drawAddPoint_49F500(v9, v10)
	return nox_client_drawLineFromPoints_49E4B0()
}
func sub_4CEBA0(a1 int32, a2 *byte) int32 {
	var (
		v2 *uint32
		v3 *uint32
		v4 *byte
		v5 *uint32
		v6 *uint32
		v8 *uint32
		v9 *byte
	)
	dword_5d4594_1523024 = uint32(uintptr(unsafe.Pointer(newWindowFromFile("rulelist.wnd", unsafe.Pointer(cgoFuncAddr(sub_4CF060))))))
	dword_5d4594_1523028 = uint32(uintptr(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1523024)))).ChildByID(10170))))
	dword_5d4594_1523032 = uint32(uintptr(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1523024)))).ChildByID(0x27BB))))
	dword_5d4594_1523036 = uint32(uintptr(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1523024)))).ChildByID(0x27BC))))
	dword_5d4594_1523040 = uint32(uintptr(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1523024)))).ChildByID(0x27BD))))
	dword_5d4594_1523044 = uint32(uintptr(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1523024)))).ChildByID(0x27BE))))
	dword_5d4594_1523048 = uint32(uintptr(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1523024)))).ChildByID(0x27BF))))
	v2 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1523024)))).ChildByID(0x27C0)))
	int32(uintptr(unsafe.Pointer(v2))).setDraw(sub_4CEED0)
	sub_46B120((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1523024)))), (*nox_window)(unsafe.Pointer(uintptr(a1))))
	v3 = *(**uint32)(unsafe.Pointer(uintptr(dword_5d4594_1523028 + 32)))
	v9 = (*byte)(unsafe.Pointer(nox_xxx_gLoadImg("UISlider")))
	v4 = (*byte)(unsafe.Pointer(nox_xxx_gLoadImg("UISliderLit")))
	v5 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1523024)))).ChildByID(0x27C3)))
	v6 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1523024)))).ChildByID(0x27C1)))
	v8 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1523024)))).ChildByID(0x27C2)))
	*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(uint32(0))*100)) + 8))) = 16
	*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(uint32(0))*100)) + 12))) = 10
	sub_4B5700(int32(uintptr(unsafe.Pointer(v5))), 0, 0, int32(uintptr(unsafe.Pointer(v9))), int32(uintptr(unsafe.Pointer(v4))), int32(uintptr(unsafe.Pointer(v4))))
	nox_xxx_wnd_46B280(int32(uintptr(unsafe.Pointer(v5))), *(*int32)(unsafe.Pointer(&dword_5d4594_1523028)))
	nox_xxx_wnd_46B280(int32(uintptr(unsafe.Pointer(v6))), *(*int32)(unsafe.Pointer(&dword_5d4594_1523028)))
	nox_xxx_wnd_46B280(int32(uintptr(unsafe.Pointer(v8))), *(*int32)(unsafe.Pointer(&dword_5d4594_1523028)))
	*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*9)) = uint32(uintptr(unsafe.Pointer(v5)))
	*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*7)) = uint32(uintptr(unsafe.Pointer(v6)))
	*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*8)) = uint32(uintptr(unsafe.Pointer(v8)))
	sub_4CED40(a2)
	return int32(dword_5d4594_1523024)
}
func sub_4CED40(a1 *byte) unsafe.Pointer {
	var (
		result       HANDLE
		v2           HANDLE
		FindFileData _WIN32_FIND_DATAA
		FileName     [64]byte
		v5           [100]libc.WChar
	)
	(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1523028))))).Func94(asWindowEvent(0x400F, 0, 0))
	nox_sprintf(&FileName[0], CString("maps\\%s\\*.rul"), a1)
	result = compatFindFirstFileA(&FileName[0], &FindFileData)
	v2 = result
	if uintptr(result) != uintptr(math.MaxUint32) {
		FindFileData.cFileName[libc.StrLen(&FindFileData.cAlternateFileName[0])+256] = 0
		if libc.StrCaseCmp(a1, &FindFileData.cAlternateFileName[0]) != 0 && libc.StrCaseCmp(CString("user"), &FindFileData.cAlternateFileName[0]) != 0 {
			nox_swprintf(&v5[0], libc.CWString("%S"), &FindFileData.cAlternateFileName[0])
			(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1523028))))).Func94(asWindowEvent(0x400D, int32(uintptr(unsafe.Pointer(&v5[0]))), -1))
		}
		for compatFindNextFileA(v2, &FindFileData) != 0 {
			FindFileData.cFileName[libc.StrLen(&FindFileData.cAlternateFileName[0])+256] = 0
			if libc.StrCaseCmp(a1, &FindFileData.cAlternateFileName[0]) != 0 {
				if libc.StrCaseCmp(CString("user"), &FindFileData.cAlternateFileName[0]) != 0 {
					nox_swprintf(&v5[0], libc.CWString("%S"), &FindFileData.cAlternateFileName[0])
					(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1523028))))).Func94(asWindowEvent(0x400D, int32(uintptr(unsafe.Pointer(&v5[0]))), -1))
				}
			}
		}
		result = unsafe.Pointer(uintptr(compatFindClose(v2)))
	}
	return result
}
func sub_4CEED0(a1 int32, a2 int32) int32 {
	var (
		v2    int32
		v3    int8
		v4    *uint16
		v5    *uint32
		xLeft int32
		yTop  int32
	)
	nox_client_wndGetPosition_46AA60((*nox_window)(unsafe.Pointer(uintptr(a1))), (*uint32)(unsafe.Pointer(&xLeft)), (*uint32)(unsafe.Pointer(&yTop)))
	if int32(int8(*(*uint8)(unsafe.Pointer(uintptr(a1 + 4))))) >= 0 {
		if *(*uint32)(unsafe.Pointer(uintptr(a2 + 20))) != 0x80000000 {
			nox_client_drawRectFilledOpaque_49CE30(xLeft, yTop, int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 8)))), int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 12)))))
		}
	} else {
		nox_client_drawImageAt_47D2C0((*nox_video_bag_image_t)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a2 + 24)))))), xLeft, yTop)
	}
	v2 = (*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1523028))))).Func94(asWindowEvent(0x4014, 0, 0))
	v3 = int8(*(*uint8)(unsafe.Pointer(uintptr(dword_5d4594_1523040 + 4))))
	if v2 < 0 {
		if int32(v3)&8 != 0 {
			nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1523040))))), 0)
		}
		if int32(*(*uint8)(unsafe.Pointer(uintptr(dword_5d4594_1523044 + 4))))&8 != 0 {
			nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1523044))))), 0)
		}
		if int32(*(*uint8)(unsafe.Pointer(uintptr(dword_5d4594_1523048 + 4))))&8 != 0 {
			nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1523048))))), 0)
		}
	} else {
		if (int32(v3) & 8) == 0 {
			nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1523040))))), 1)
		}
		if (int32(*(*uint8)(unsafe.Pointer(uintptr(dword_5d4594_1523044 + 4))))&8) == 0 && !noxflags.HasGame(noxflags.GameFlag15|noxflags.GameFlag16) {
			nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1523044))))), 1)
		}
		if (int32(*(*uint8)(unsafe.Pointer(uintptr(dword_5d4594_1523048 + 4)))) & 8) == 0 {
			nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1523048))))), 1)
		}
	}
	v4 = (*uint16)(unsafe.Pointer(uintptr((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1523032))))).Func94(asWindowEvent(0x401D, 0, 0)))))
	v5 = (*uint32)(unsafe.Pointer(nox_xxx_wndGetFocus_46B4F0()))
	if v5 != nil && *v5 == 0x27BB {
		if v4 != nil && int32(*v4) != 0 {
			if (int32(*(*uint8)(unsafe.Pointer(uintptr(dword_5d4594_1523036 + 4)))) & 8) == 0 {
				nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1523036))))), 1)
				return 1
			}
		} else if int32(*(*uint8)(unsafe.Pointer(uintptr(dword_5d4594_1523036 + 4))))&8 != 0 {
			nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1523036))))), 0)
		}
	}
	return 1
}
func sub_4CF060(a1 int32, a2 uint32, a3 *int32, a4 int32) int32 {
	var (
		v4  *uint32
		v6  *byte
		v7  *byte
		v8  *byte
		v9  int32
		v10 int32
		v11 int32
		v12 *byte
		v13 *byte
		v14 *byte
		v15 int32
		v16 int32
		v17 *byte
		v18 *byte
		v19 int32
		v20 int32
		v21 int32
		v22 *libc.WChar
		v23 *byte
		v24 *byte
		v25 int32
		v26 int32
		v27 *libc.WChar
		v28 [16]byte
	)
	if a2 > 0x4007 {
		if a2 == 16400 {
			(*nox_window)(unsafe.Pointer(a3)).ID()
		}
		return 1
	}
	if a2 != 0x4007 {
		if a2 != 23 && a2 == 0x4003 {
			v4 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1523024)))).ChildByID(a4)))
			if v4 == nil {
				return 0
			}
			if int32(uint16(uintptr(unsafe.Pointer(a3)))) == 1 {
				return 0
			}
			v6 = (*byte)(unsafe.Pointer(uintptr((*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v4)))))).Func94(asWindowEvent(0x401D, 0, 0)))))
			v7 = v6
			if v6 != nil {
				if *v6 != 0 {
					libc.Atoi(libc.GoString(v6))
					if a4 == 0x27BB {
						v8 = sub_4165B0()
						if libc.StrCaseCmp(v7, v8) == 0 || libc.StrCaseCmp(v7, CString("user")) == 0 {
							nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1523036))))), 0)
							return 1
						}
					}
				}
			}
		}
		return 1
	}
	v9 = (*nox_window)(unsafe.Pointer(a3)).ID()
	clientPlaySoundSpecial(766, 100)
	switch v9 {
	case 0x27BC:
		sub_416580()
		v22 = (*libc.WChar)(unsafe.Pointer(uintptr((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1523032))))).Func94(asWindowEvent(0x401D, 0, 0)))))
		nox_sprintf(&v28[0], CString("%S%s"), v22, memmap.PtrOff(0x587000, 191640))
		v23 = sub_4165B0()
		sub_459AA0(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v23))))))
		v24 = sub_4165B0()
		sub_57AAA0(&v28[0], v24, nil)
		v25 = 0
		v26 = int32(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1523028 + 32))))
		if int32(*(*int16)(unsafe.Pointer(uintptr(v26 + 44)))) <= 0 {
			goto LABEL_22
		}
	case 0x27BD:
		sub_416580()
		v10 = (*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1523028))))).Func94(asWindowEvent(0x4014, 0, 0))
		v11 = (*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1523028))))).Func94(asWindowEvent(0x4016, v10, 0))
		nox_sprintf(&v28[0], CString("%S%s"), v11, memmap.PtrOff(0x587000, 191592))
		v12 = sub_4165B0()
		sub_459AA0(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v12))))))
		v13 = sub_4165B0()
		sub_57AAA0(&v28[0], v13, nil)
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1523028))))).Func94(asWindowEvent(0x4013, -1, 0))
		return 1
	case 0x27BE:
		sub_416580()
		v14 = sub_4165B0()
		v15 = (*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1523028))))).Func94(asWindowEvent(0x4014, 0, 0))
		v16 = (*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1523028))))).Func94(asWindowEvent(0x4016, v15, 0))
		nox_sprintf(&v28[0], CString("%S%s"), v16, memmap.PtrOff(0x587000, 191608))
		sub_57A1E0((*int32)(unsafe.Pointer(v14)), &v28[0], nil, 7, int16(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v14))), unsafe.Sizeof(uint16(0))*26)))))
		sub_453F70(unsafe.Add(unsafe.Pointer(v14), 24))
		sub_4535E0((*int32)(unsafe.Add(unsafe.Pointer((*int32)(unsafe.Pointer(v14))), unsafe.Sizeof(int32(0))*11)))
		sub_4535F0(int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v14))), unsafe.Sizeof(uint32(0))*12)))))
		v17 = sub_4165B0()
		sub_459880(int32(uintptr(unsafe.Pointer(v17))))
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1523028))))).Func94(asWindowEvent(0x4013, -1, 0))
		sub_459D50(1)
		return 1
	case 0x27BF:
		v18 = sub_4165B0()
		v19 = (*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1523028))))).Func94(asWindowEvent(0x4014, 0, 0))
		v20 = v19
		v21 = (*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1523028))))).Func94(asWindowEvent(0x4016, v19, 0))
		nox_sprintf(&v28[0], CString("%S%s"), v21, memmap.PtrOff(0x587000, 191624))
		sub_57A9F0(v18, &v28[0])
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1523028))))).Func94(asWindowEvent(0x400E, v20, 0))
		return 1
	default:
		return 1
	}
	for {
		v27 = (*libc.WChar)(unsafe.Pointer(uintptr((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1523028))))).Func94(asWindowEvent(0x4016, v25, 0)))))
		if _nox_wcsicmp(v22, v27) == 0 {
			break
		}
		if func() int32 {
			p := &v25
			*p++
			return *p
		}() >= int32(*(*int16)(unsafe.Pointer(uintptr(v26 + 44)))) {
		LABEL_22:
			(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1523028))))).Func94(asWindowEvent(0x400D, int32(uintptr(unsafe.Pointer(v22))), -1))
			(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1523032))))).Func94(asWindowEvent(0x401E, int32(uintptr(memmap.PtrOff(6112660, 1523056))), 0))
			nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1523036))))), 0)
			return 1
		}
	}
	(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1523032))))).Func94(asWindowEvent(0x401E, int32(uintptr(memmap.PtrOff(6112660, 1523052))), 0))
	return 1
}
func nox_xxx_mapValidateMB_4CF470(a1 *byte, a2 int32) int32 {
	var (
		v2       int32
		v4       int32
		v5       int32
		FileName [1024]byte
	)
	v2 = 0
	if a2 == 0 {
		return 6
	}
	if a1 != nil {
		if libc.StrChr(a1, byte('\\')) != nil {
			libc.StrCpy(&FileName[0], a1)
		} else {
			libc.StrCpy(&FileName[0], CString("maps\\"))
			libc.StrNCat(&FileName[0], a1, func() int {
				if libc.StrLen(a1)-4 < 1024 {
					return libc.StrLen(a1) - 4
				}
				return 1024
			}())
			*(*uint16)(unsafe.Pointer(&FileName[libc.StrLen(&FileName[0])])) = *memmap.PtrUint16(0x587000, 191672)
			libc.StrCat(&FileName[0], a1)
		}
		if nox_fs_access(&FileName[0], 0) != -1 {
			v4 = 0
			if nox_fs_access(&FileName[0], 2) == -1 {
				v2 = 1
			}
			if cryptFileOpen(&FileName[0], 1, 19) != 0 {
				v2 |= 2
				cryptFileReadWrite((*uint8)(unsafe.Pointer(&v5))[:4])
				if v5 != -86065425 && v5 == -86050098 {
					cryptFileReadMaybeAlign((*uint8)(unsafe.Pointer(&v4))[:4])
					if v4 == a2 {
						v2 |= 4
					}
				}
				cryptFileClose()
			}
		}
	}
	return v2
}
func nox_xxx_mapFindCrown_4CFC30() {
	var (
		v0 int32
		v1 int32
	)
	if *memmap.PtrUint32(6112660, 1523076) == 0 {
		*memmap.PtrUint32(6112660, 1523076) = uint32(nox_xxx_getNameId_4E3AA0(CString("Crown")))
	}
	v0 = int32(uintptr(unsafe.Pointer(noxServer.firstServerObject())))
	if v0 != 0 {
		for {
			v1 = int32(uintptr(unsafe.Pointer((*nox_object_t)(unsafe.Pointer(uintptr(v0))).Next())))
			if uint32(*(*uint16)(unsafe.Pointer(uintptr(v0 + 4)))) == *memmap.PtrUint32(6112660, 1523076) {
				nox_xxx_delayedDeleteObject_4E5CC0((*nox_object_t)(unsafe.Pointer(uintptr(v0))))
				sub_4EC6A0(v0)
			}
			v0 = v1
			if v1 == 0 {
				break
			}
		}
	}
}
func sub_4CFDF0(a1 int32) int32 {
	var result int32
	result = a1
	*memmap.PtrUint32(6112660, 1523072) = uint32(a1)
	return result
}
func sub_4CFE00() int32 {
	return int32(*memmap.PtrUint32(6112660, 1523072))
}
func nox_xxx_mapGetTypeMB_4CFFA0(a1 unsafe.Pointer) int32 {
	return nox_mapToGameFlags_4CFF50(int32(*(*uint32)(unsafe.Pointer(uintptr(int32(uintptr(a1)) + 1392)))))
}
func sub_4CFFC0(a1 int32) int32 {
	return nox_mapToGameFlags_4CFF50(int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 28)))))
}
func sub_4CFFE0(a1 int32) int32 {
	var result int32
	result = int32(uintptr(unsafe.Pointer(noxServer.firstServerObjectUninited())))
	if result == 0 {
		return 0
	}
	for *(*uint32)(unsafe.Pointer(uintptr(result + 44))) != uint32(a1) {
		result = int32(uintptr(unsafe.Pointer(nox_server_getNextObjectUninited_4DA880((*nox_object_t)(unsafe.Pointer(uintptr(result)))))))
		if result == 0 {
			return 0
		}
	}
	return result
}
func nox_xxx_interesting_xfer_4D0010(a1 *uint32, a2 int32) int32 {
	var (
		i   int32
		v3  *uint32
		v4  *byte
		v5  func(int32) int32
		v6  int32
		v7  int32
		v8  int32
		v9  int32
		v10 int32
		v11 int32
		v12 int32
		v13 int32
		v14 int32
		v15 *byte
		v16 int32
		v17 int32
		v18 *byte
		v19 int32
		v20 *uint32
		v21 int32
		v22 int32
		v23 *byte
	)
	for i = int32(uintptr(unsafe.Pointer(noxServer.firstServerObjectUninited()))); i != 0; i = int32(uintptr(unsafe.Pointer(nox_server_getNextObjectUninited_4DA880((*nox_object_t)(unsafe.Pointer(uintptr(i))))))) {
		*(*uint32)(unsafe.Pointer(uintptr(i + 44))) = *(*uint32)(unsafe.Pointer(uintptr(i + 40)))
		*(*uint32)(unsafe.Pointer(uintptr(i + 40))) = uint32(func() int32 {
			p := &a2
			x := *p
			*p++
			return x
		}())
	}
	v3 = (*uint32)(unsafe.Pointer(noxServer.firstServerObjectUninited()))
	if v3 != nil {
		for {
			v4 = (*byte)(unsafe.Pointer(uintptr(nox_xxx_getUnitName_4E39D0(int32(uintptr(unsafe.Pointer(v3)))))))
			v5 = cgoAsFunc(*(*uint32)(unsafe.Pointer((*nox_objectType_t)(unsafe.Add(unsafe.Pointer(nox_xxx_objectTypeByID_4E3B60(v4)), unsafe.Sizeof(nox_objectType_t{})*212)))), (*func(int32) int32)(nil))
			if cgoFuncAddr(v5) != cgoFuncAddr(nox_xxx_XFerElevator_4F53D0) {
				break
			}
			v6 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*187)))
			v7 = sub_4CFFE0(int32(*(*uint32)(unsafe.Pointer(uintptr(v6 + 8)))))
			if v7 == 0 {
				goto LABEL_7
			}
			v8 = int32(*(*uint32)(unsafe.Pointer(uintptr(v7 + 40))))
			*(*uint32)(unsafe.Pointer(uintptr(v6 + 4))) = uint32(v7)
			*(*uint32)(unsafe.Pointer(uintptr(v6 + 8))) = uint32(v8)
		LABEL_28:
			v3 = (*uint32)(unsafe.Pointer(nox_server_getNextObjectUninited_4DA880((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v3)))))))))
			if v3 == nil {
				return a2
			}
		}
		if cgoFuncAddr(v5) != cgoFuncAddr(nox_xxx_XFerElevatorShaft_4F54A0) {
			if cgoFuncAddr(v5) == cgoFuncAddr(nox_xxx_XFerTransporter_4F5300) {
				v11 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*187)))
				v12 = sub_4CFFE0(int32(*(*uint32)(unsafe.Pointer(uintptr(v11 + 16)))))
				if v12 != 0 {
					v13 = int32(*(*uint32)(unsafe.Pointer(uintptr(v12 + 40))))
					*(*uint32)(unsafe.Pointer(uintptr(v11 + 12))) = uint32(v12)
					*(*uint32)(unsafe.Pointer(uintptr(v11 + 16))) = uint32(v13)
				} else {
					*(*uint32)(unsafe.Pointer(uintptr(v11 + 16))) = 0
					*(*uint32)(unsafe.Pointer(uintptr(v11 + 12))) = 0
				}
			} else if cgoFuncAddr(v5) == cgoFuncAddr(nox_xxx_XFerHole_4F51D0) {
				v14 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*175)))
				v15 = nox_xxx_mapGetWallSize_426A70()
				v16 = int32(*(*uint32)(unsafe.Pointer(uintptr(v14 + 12))))
				*(*uint32)(unsafe.Pointer(uintptr(v14 + 8))) += *a1 - *(*uint32)(unsafe.Pointer(v15))*23
				*(*uint32)(unsafe.Pointer(uintptr(v14 + 12))) = *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*1)) - *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v15))), unsafe.Sizeof(uint32(0))*1)))*23 + uint32(v16)
			} else if cgoFuncAddr(v5) == cgoFuncAddr(nox_xxx_XFerExit_4F4B90) {
				v17 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*175)))
				v18 = nox_xxx_mapGetWallSize_426A70()
				*(*float32)(unsafe.Pointer(uintptr(v17 + 80))) = float32(float64(int32(*a1-*(*uint32)(unsafe.Pointer(v18))*23)) + float64(*(*float32)(unsafe.Pointer(uintptr(v17 + 80)))))
				*(*float32)(unsafe.Pointer(uintptr(v17 + 84))) = float32(float64(int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*1))-*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v18))), unsafe.Sizeof(uint32(0))*1)))*23)) + float64(*(*float32)(unsafe.Pointer(uintptr(v17 + 84)))))
			} else if cgoFuncAddr(v5) == cgoFuncAddr(nox_xxx_XFerMover_4F5730) {
				v19 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*187)))
				v20 = (*uint32)(unsafe.Pointer(uintptr(sub_579C60(int32(*(*uint32)(unsafe.Pointer(uintptr(v19 + 8))))))))
				if v20 != nil {
					*(*uint32)(unsafe.Pointer(uintptr(v19 + 8))) = *v20
				} else {
					*(*uint32)(unsafe.Pointer(uintptr(v19 + 8))) = 0
				}
				v21 = sub_4CFFE0(int32(*(*uint32)(unsafe.Pointer(uintptr(v19 + 32)))))
				if v21 != 0 {
					*(*uint32)(unsafe.Pointer(uintptr(v19 + 32))) = *(*uint32)(unsafe.Pointer(uintptr(v21 + 40)))
				} else {
					*(*uint32)(unsafe.Pointer(uintptr(v19 + 32))) = 0
				}
			} else if cgoFuncAddr(v5) == cgoFuncAddr(nox_xxx_XFerGlyph_4F5890) {
				v22 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*173)))
				v23 = nox_xxx_mapGetWallSize_426A70()
				*(*float32)(unsafe.Pointer(uintptr(v22 + 28))) = float32(float64(int32(*a1-*(*uint32)(unsafe.Pointer(v23))*23)) + float64(*(*float32)(unsafe.Pointer(uintptr(v22 + 28)))))
				*(*float32)(unsafe.Pointer(uintptr(v22 + 32))) = float32(float64(int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*1))-*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v23))), unsafe.Sizeof(uint32(0))*1)))*23)) + float64(*(*float32)(unsafe.Pointer(uintptr(v22 + 32)))))
			}
			goto LABEL_28
		}
		v6 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*187)))
		v9 = sub_4CFFE0(int32(*(*uint32)(unsafe.Pointer(uintptr(v6 + 8)))))
		if v9 != 0 {
			v10 = int32(*(*uint32)(unsafe.Pointer(uintptr(v9 + 40))))
			*(*uint32)(unsafe.Pointer(uintptr(v6 + 4))) = uint32(v9)
			*(*uint32)(unsafe.Pointer(uintptr(v6 + 8))) = uint32(v10)
			goto LABEL_28
		}
	LABEL_7:
		*(*uint32)(unsafe.Pointer(uintptr(v6 + 8))) = 0
		*(*uint32)(unsafe.Pointer(uintptr(v6 + 4))) = 0
		goto LABEL_28
	}
	return a2
}
func sub_4D0250(a1 *byte, a2 *byte, a3 int32, a4 int32) int32 {
	var (
		result int32
		v5     *File
		v6     *uint32
		v7     int32
		v8     *uint32
		i      int32
		j      int32
		v11    int32
		k      int32
		l      int32
		v14    int32
		v15    int32
		v16    int32
		v17    [2]int32
		v19    int32
		v20    int32
		v21    int32
		v22    int32
		v23    int32
		v24    int32
		v25    int32
		v26    int32
		v27    int32
		v28    int4
		v29    [256]byte
	)
	if a3 <= 0 {
		return 0
	}
	result = cryptFileOpen(a1, 1, -1)
	if result != 0 {
		v5 = nox_xxx_mapgenGetSomeFile_426A60()
		nox_fs_fread(v5, unsafe.Pointer(&v16), 4)
		if v16 == a4 {
			nox_fs_fread(v5, unsafe.Pointer(&v17[0]), 4)
			nox_fs_fread(v5, unsafe.Pointer(&v17[1]), 4)
			nox_xxx_mapWall_426A80(&v17[0])
			nox_fs_fread(v5, unsafe.Pointer(&v19), 4)
			nox_fs_fread(v5, unsafe.Pointer(&v20), 4)
			nox_fs_fread(v5, unsafe.Pointer(&v25), 4)
			nox_fs_fread(v5, unsafe.Pointer(&v26), 4)
			nox_fs_fread(v5, unsafe.Pointer(&v21), 4)
			nox_fs_fread(v5, unsafe.Pointer(&v22), 4)
			nox_fs_fread(v5, unsafe.Pointer(&v23), 4)
			nox_fs_fread(v5, unsafe.Pointer(&v24), 4)
			v6 = (*uint32)(unsafe.Pointer(a2))
			if a2 == nil {
				v6 = (*uint32)(unsafe.Pointer(&v19))
			}
			sub_428170(unsafe.Pointer(v6), &v28)
			v7 = nox_xxx_wallGet_426A30()
			sub_426A20(5)
			nox_xxx_cryptSetTypeMB_426A50(1)
			for {
				cryptFileReadWrite((*uint8)(unsafe.Pointer(&v15))[:1])
				if int32(uint8(int8(v15))) == 0 {
					break
				}
				cryptFileReadWrite((*uint8)(unsafe.Pointer(&v29[0]))[:uint32(uint8(int8(v15)))])
				cryptFileReadWrite((*uint8)(unsafe.Pointer(&v27))[:4])
				if nox_xxx_mapReadSection_426EA0(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v6))))), &v29[0], (*uint32)(unsafe.Pointer(&v14))) == 0 {
					if v14 == 1 {
						goto LABEL_20
					}
					v8 = (*uint32)(unsafe.Pointer(nox_xxx_newObjectByTypeID_4E3810(&v29[0])))
					if v8 == nil {
						nox_xxx_cryptSetTypeMB_426A50(0)
						return 0
					}
					if (cgoAsFunc(*(*uint32)(unsafe.Add(unsafe.Pointer(v8), unsafe.Sizeof(uint32(0))*176)), (*func(*uint32, *int4) int32)(nil)).(func(*uint32, *int4) int32))(v8, &v28) == 0 {
						sub_426A20(v7)
						nox_xxx_objectFreeMem_4E38A0(int32(uintptr(unsafe.Pointer(v8))))
						cryptFileClose()
						return 0
					}
					nox_xxx_servMapLoadPlaceObj_4F3F50((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v8)))))), 0, unsafe.Pointer(&v28.field_0))
				}
			}
			for i = int32(uintptr(unsafe.Pointer(noxServer.firstServerObjectUninited()))); i != 0; i = int32(uintptr(unsafe.Pointer(nox_server_getNextObjectUninited_4DA880((*nox_object_t)(unsafe.Pointer(uintptr(i))))))) {
				*(*uint32)(unsafe.Pointer(uintptr(i + 16))) |= 0x80000000
			}
			nox_xxx_cryptSetTypeMB_426A50(0)
			sub_426A20(v7)
			cryptFileClose()
			if sub_579D20() == 0 {
			LABEL_20:
				sub_426A20(v7)
				cryptFileClose()
				return 0
			}
			for j = sub_579890(); j != 0; j = sub_5798A0(j) {
				*(*uint32)(unsafe.Pointer(uintptr(j + 480))) |= 0x80000000
			}
			v11 = nox_xxx_interesting_xfer_4D0010((*uint32)(unsafe.Pointer(&v28)), a3)
			for k = sub_579890(); k != 0; k = sub_5798A0(k) {
				*(*uint32)(unsafe.Pointer(uintptr(k + 4))) = 0
			}
			for l = int32(uintptr(unsafe.Pointer(noxServer.firstServerObjectUninited()))); l != 0; l = int32(uintptr(unsafe.Pointer(nox_server_getNextObjectUninited_4DA880((*nox_object_t)(unsafe.Pointer(uintptr(l))))))) {
				*(*uint32)(unsafe.Pointer(uintptr(l + 44))) = 0
			}
			nox_xxx_waypoint_5799C0()
			nox_xxx_unitClearPendingMB_4DB030()
			result = v11
		} else {
			result = 0
		}
	}
	return result
}
func sub_4D0550(a1 *byte) int32 {
	var (
		result int32
		v2     uint32
		v3     int8
		v4     int32
		v5     uint8
		v6     *byte
		v7     *byte
		v8     uint8
		v9     uint8
		v10    [256]byte
	)
	result = 0
	if a1 != nil {
		v10[0] = 0
		libc.StrNCat(&v10[0], a1, func() int {
			if libc.StrLen(a1)-4 < 256 {
				return libc.StrLen(a1) - 4
			}
			return 256
		}())
		v2 = uint32(libc.StrLen(&v10[0]) + 1)
		v3 = int8(uint8(v2 - 1))
		v9 = uint8(v2 - 1)
		if int32(uint8(v2)) != 1 {
			for v10[v9] != 92 {
				v9 = uint8(func() int8 {
					p := &v3
					*p--
					return *p
				}())
				if int32(v3) == 0 {
					goto LABEL_7
				}
			}
			v10[int32(v9)+1] = 0
		}
	LABEL_7:
		v4 = int32(*memmap.PtrUint32(0x587000, 191752))
		v5 = *memmap.PtrUint8(0x587000, 191756)
		v6 = &v10[libc.StrLen(&v10[0])]
		*(*uint32)(unsafe.Pointer(v6)) = *memmap.PtrUint32(0x587000, 191748)
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v6))), unsafe.Sizeof(uint32(0))*1))) = uint32(v4)
		*(*byte)(unsafe.Add(unsafe.Pointer(v6), 8)) = byte(v5)
		if sub_4D0670(&v10[0]) == 0 {
			v10[0] = 0
			libc.StrNCat(&v10[0], a1, func() int {
				if libc.StrLen(a1)-4 < 256 {
					return libc.StrLen(a1) - 4
				}
				return 256
			}())
			v7 = &v10[libc.StrLen(&v10[0])+1]
			v8 = *memmap.PtrUint8(0x587000, 191764)
			*(*uint32)(unsafe.Pointer(func() *byte {
				p := &v7
				*p = (*byte)(unsafe.Add(unsafe.Pointer(*p), -1))
				return *p
			}())) = *memmap.PtrUint32(0x587000, 191760)
			*(*byte)(unsafe.Add(unsafe.Pointer(v7), 4)) = byte(v8)
			sub_4D0670(&v10[0])
		}
		if nox_xxx_check_flag_aaa_43AF70() == 1 {
			sub_4D0670((*byte)(memmap.PtrOff(0x587000, 191768)))
		}
		result = 1
	}
	return result
}
func sub_4D0670(a1 *byte) int32 {
	var (
		v1 int32
		v2 *File
		v3 *File
		v4 *byte
		v5 int32
		v7 [255]byte
		v8 [256]libc.WChar
	)
	v1 = 6128
	v2 = nox_fs_open_text(a1)
	v3 = v2
	if v2 == nil {
		return 0
	}
	if !nox_fs_feof(v3) {
		for {
			libc.MemSet(unsafe.Pointer(&v7[0]), 0, 252)
			*(*uint16)(unsafe.Pointer(&v7[252])) = 0
			v7[254] = 0
			nox_fs_fgets(v3, &v7[0], math.MaxUint8)
			v4 = libc.StrChr(&v7[0], 10)
			if v4 != nil {
				*v4 = 0
			}
			if v7[0] != 0 {
				nox_swprintf(&v8[0], libc.CWString("%S"), &v7[0])
				v5 = sub_57AE30(&v7[0])
				if v5 != 0 {
					v1 = v5
				} else if noxflags.HasGame(uint32(v1)) {
					nox_server_parseCmdText_443C80(&v8[0], 1)
				}
			}
			if nox_fs_feof(v3) {
				break
			}
		}
	}
	nox_fs_close(v3)
	return 1
}
func nox_common_maplist_add_4D0760(map_ *nox_map_list_item) {
	var it *nox_map_list_item = (*nox_map_list_item)(unsafe.Pointer(nox_common_list_getFirstSafe_425890(&nox_common_maplist)))
	if it == nil {
		nox_common_list_append_4258E0(&nox_common_maplist, &map_.list)
		return
	}
	for libc.StrCmp(&map_.name[0], &it.name[0]) > 0 {
		it = (*nox_map_list_item)(unsafe.Pointer(nox_common_list_getNextSafe_4258A0(&it.list)))
		if it == nil {
			nox_common_list_append_4258E0(&nox_common_maplist, &map_.list)
			return
		}
	}
	nox_common_list_append_4258E0(&it.list, &map_.list)
}
func nox_common_maplist_free_4D0970() {
	var (
		result *int32
		v1     *int32
		v2     *int32
	)
	result = (*int32)(unsafe.Pointer(nox_common_list_getFirstSafe_425890(&nox_common_maplist)))
	v1 = result
	if result != nil {
		for {
			v2 = (*int32)(unsafe.Pointer(nox_common_list_getNextSafe_4258A0((*nox_list_item_t)(unsafe.Pointer(v1)))))
			sub_425920((**uint32)(unsafe.Pointer(v1)))
			alloc.Free(unsafe.Pointer(v1))
			v1 = v2
			if v2 == nil {
				break
			}
		}
	}
}
func nox_common_maplist_first_4D09B0() *nox_map_list_item {
	return (*nox_map_list_item)(unsafe.Pointer(nox_common_list_getFirstSafe_425890(&nox_common_maplist)))
}
func nox_common_maplist_next_4D09C0(list *nox_map_list_item) *nox_map_list_item {
	return (*nox_map_list_item)(unsafe.Pointer(nox_common_list_getNextSafe_4258A0(&list.list)))
}
func nox_common_maplist_mapByName_4D09D0(a1 *byte) *nox_map_list_item {
	var v1 *int32
	v1 = (*int32)(unsafe.Pointer(nox_common_maplist_first_4D09B0()))
	if v1 == nil {
		return nil
	}
	for libc.StrCaseCmp(a1, (*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(v1))), 12))) != 0 || *(*int32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(int32(0))*6)) == 0 {
		v1 = (*int32)(unsafe.Pointer(nox_common_maplist_next_4D09C0((*nox_map_list_item)(unsafe.Pointer(v1)))))
		if v1 == nil {
			return nil
		}
	}
	return (*nox_map_list_item)(unsafe.Pointer(v1))
}
func nox_xxx_loadMapCycle_4D0A30() *File {
	var (
		v0     int32
		result *File
		v7     *File
		v8     int32
		v9     int32
		v10    int32
		v11    int32
		v12    *byte
		v13    int32
		v14    int32
		v15    *File
		v16    [128]byte
		v17    [128]byte
	)
	v0 = 1
	libc.MemSet(memmap.PtrOff(6112660, 1548428), 0, 24)
	libc.StrCpy((*byte)(memmap.PtrOff(6112660, 1524108)), nox_fs_root())
	libc.StrCat((*byte)(memmap.PtrOff(6112660, 1524108)), CString("\\mapcycle.txt"))
	result = nox_fs_open_text((*byte)(memmap.PtrOff(6112660, 1524108)))
	v7 = result
	v15 = result
	if result != nil {
		if nox_fs_fgets(result, &v16[0], math.MaxInt8) {
			sub_4D0CC0(&v16[0])
			v8 = sub_4D0C80(&v16[0])
			if v8 < 0 {
				v9 = int32(*memmap.PtrUint32(6112660, 1548432))
				libc.StrCpy((*byte)(memmap.PtrOff(6112660, *memmap.PtrUint32(6112660, 1548432)*128+0x17620C)), &v16[0])
				*memmap.PtrUint32(6112660, 1548432) = uint32(v9 + 1)
			} else {
				v0 = v8
			}
		}
		for !nox_fs_feof(v7) {
			if nox_fs_fgets(v7, &v16[0], math.MaxInt8) {
				sub_4D0CC0(&v16[0])
				v10 = sub_4D0C80(&v16[0])
				if v10 < 0 {
					if *memmap.PtrInt32(6112660, uint32(v0*4)+0x17A08C) < 25 && v16[0] != 0 {
						v11 = sub_4D0C70(v0)
						libc.StrCpy(&v17[0], &v16[0])
						v12 = libc.StrTok(&v17[0], CString(".\n"))
						if nox_common_checkMapFile_4CFE10(v12) != 0 {
							v13 = nox_xxx_mapGetTypeMB_4CFFA0(unsafe.Pointer(uintptr(int32(uintptr(memmap.PtrOff(0x973F18, 2408))))))
							if v13&v11 != 0 {
								v14 = int32(*memmap.PtrUint32(6112660, uint32(v0*4)+0x17A08C))
								libc.StrCpy((*byte)(memmap.PtrOff(6112660, uint32((v14+v0*20+v0*5)*128)+0x17558C)), &v16[0])
								*memmap.PtrUint32(6112660, uint32(v0*4)+0x17A08C) = uint32(v14 + 1)
							}
						}
						v7 = v15
					}
				} else {
					v0 = v10
				}
			}
		}
		nox_fs_close(v7)
		result = nil
	} else {
		*memmap.PtrUint32(6112660, 1548484) = 0
	}
	return result
}
func sub_4D0C70(a1 int32) int32 {
	return int32(*memmap.PtrUint32(0x587000, uint32(a1*8)+0x2ED5C))
}
func sub_4D0C80(a1 *byte) int32 {
	var (
		v1 int32
		v2 **byte
	)
	v1 = 0
	v2 = (**byte)(memmap.PtrOff(0x587000, 191832))
	for libc.StrCaseCmp(*v2, a1) != 0 {
		v2 = (**byte)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof((*byte)(nil))*2))
		v1++
		if int32(uintptr(unsafe.Pointer(v2))) >= int32(uintptr(memmap.PtrOff(0x587000, 191880))) {
			return -1
		}
	}
	return v1
}
func sub_4D0CC0(a1 *byte) {
	var (
		v1 *byte
		v2 *byte
	)
	if a1 != nil {
		v1 = libc.StrChr(a1, 13)
		if v1 != nil {
			*v1 = 0
		}
		v2 = libc.StrChr(a1, 10)
		if v2 != nil {
			*v2 = 0
		}
	}
}
func sub_4D0D50(a1 int32) int32 {
	var (
		result int32
		v2     *uint8
	)
	result = 0
	v2 = (*uint8)(memmap.PtrOff(0x587000, 191836))
	for (uint32(a1) & *(*uint32)(unsafe.Pointer(v2))) == 0 {
		v2 = (*uint8)(unsafe.Add(unsafe.Pointer(v2), 8))
		result++
		if int32(uintptr(unsafe.Pointer(v2))) >= int32(uintptr(memmap.PtrOff(0x587000, 191884))) {
			return 0
		}
	}
	return result
}
func sub_4D0D70() int32 {
	return bool2int(*memmap.PtrUint32(6112660, 1548484) != 0 || nox_common_getEngineFlag(nox_engine_flag(NOX_ENGINE_FLAG_DISABLE_GRAPHICS_RENDERING)))
}
func sub_4D0D90(a1 int32) int32 {
	var result int32
	result = a1
	*memmap.PtrUint32(6112660, 1548484) = uint32(a1)
	return result
}
func sub_4D0DA0() {
	libc.MemSet(memmap.PtrOff(6112660, 1548452), 0, 24)
	libc.MemSet(memmap.PtrOff(6112660, 1548428), 0, 24)
}
func sub_4D0DC0(a1 int32, a2 int32) int32 {
	var result int32
	result = sub_4D0D50(a1)
	*memmap.PtrUint32(6112660, uint32(result*4)+0x17A0A4) = uint32(a2)
	return result
}
func sub_4D0DE0(a1 int32) int32 {
	return int32(*memmap.PtrUint32(6112660, uint32(sub_4D0D50(a1)*4)+0x17A0A4))
}
func nox_xxx_mapSelectFirst_4D0E00() int32 {
	var (
		i  *int32
		v3 int32
		v4 uint8
		v5 *uint8
	)
	_ = v5
	var v6 int32
	var result int32
	var v8 int32
	var v9 int32
	var v10 *uint8
	var v11 int32
	var v12 int32
	var v13 int32
	nox_platform_srand_time()
	dword_5d4594_1548476 = 0
	for i = (*int32)(unsafe.Pointer(nox_common_maplist_first_4D09B0())); i != nil; i = (*int32)(unsafe.Pointer(nox_common_maplist_next_4D09C0((*nox_map_list_item)(unsafe.Pointer(i))))) {
		if *(*int32)(unsafe.Add(unsafe.Pointer(i), unsafe.Sizeof(int32(0))*6)) != 0 {
			if sub_4CFFC0(int32(uintptr(unsafe.Pointer(i))))&4096 != 0 {
				if *(*int32)(unsafe.Pointer(&dword_5d4594_1548476)) < 128 {
					v3 = int32(dword_5d4594_1548476 * 32)
					libc.StrCpy((*byte)(memmap.PtrOff(6112660, dword_5d4594_1548476*32+0x174590)), (*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(i))), 12)))
					v4 = *memmap.PtrUint8(0x587000, 192004)
					v5 = (*uint8)(memmap.PtrOff(6112660, uint32(v3)+0x174590+uint32(libc.StrLen((*byte)(memmap.PtrOff(6112660, uint32(v3)+0x174590))))))
					*(*uint32)(unsafe.Pointer(v5)) = *memmap.PtrUint32(0x587000, 192000)
					*(*uint8)(unsafe.Add(unsafe.Pointer(v5), 4)) = v4
					v6 = int32(dword_5d4594_1548476 + 1)
					*memmap.PtrUint32(6112660, uint32(v3)+0x17458C) = 0
					dword_5d4594_1548476 = uint32(v6)
				}
			}
		}
	}
	result = int32(dword_5d4594_1548476)
	v8 = 1
	if dword_5d4594_1548476 > 0 {
		v9 = 1
		v10 = (*uint8)(memmap.PtrOff(6112660, 1525132))
		for {
			if *(*uint32)(unsafe.Pointer(v10)) == 0 {
				*(*uint32)(unsafe.Pointer(v10)) = uint32(func() int32 {
					p := &v8
					x := *p
					*p++
					return x
				}())
				v13 = v8
				v11 = v9
				if v9 < result {
					v12 = int32(uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v10), 32)))))
					for {
						if libc.StrNCaseCmp((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(v10))), 4)), (*byte)(unsafe.Pointer(uintptr(v12+4))), 6) == 0 {
							*(*uint32)(unsafe.Pointer(uintptr(v12))) = *(*uint32)(unsafe.Pointer(v10))
						}
						result = int32(dword_5d4594_1548476)
						v11++
						v12 += 32
						if v11 >= *(*int32)(unsafe.Pointer(&dword_5d4594_1548476)) {
							break
						}
					}
					v8 = v13
				}
			}
			v9++
			v10 = (*uint8)(unsafe.Add(unsafe.Pointer(v10), 32))
			if v9-1 >= result {
				break
			}
		}
	}
	return result
}
func sub_4D0F30() {
	var (
		v0 int32
		v1 *uint8
	)
	v0 = int32(dword_5d4594_1548476)
	dword_5d4594_1548480 = 1000
	if dword_5d4594_1548476 > 0 {
		v1 = (*uint8)(memmap.PtrOff(6112660, 1525160))
		for {
			*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v1))), -int(unsafe.Sizeof(uint32(0))*1)))) = 0
			*(*uint32)(unsafe.Pointer(v1)) = 0
			v1 = (*uint8)(unsafe.Add(unsafe.Pointer(v1), 32))
			v0--
			if v0 == 0 {
				break
			}
		}
	}
}
func nox_xxx_getQuestMapFile_4D0F60() *byte {
	var (
		v1  int32
		v2  *uint8
		v3  int32
		v4  int32
		v5  *uint8
		v6  int32
		v7  *uint8
		v8  int32
		v9  int32
		v10 int32
		v11 *uint8
		v12 int32
		v13 int32
		v14 int32
		i   *uint8
		v16 int32
	)
	if dword_5d4594_1548476 == 0 {
		return nil
	}
	if dword_5d4594_1548476 == 1 {
		return (*byte)(memmap.PtrOff(6112660, 1525136))
	}
	v1 = 0
	v16 = 0
	if *(*int32)(unsafe.Pointer(&dword_5d4594_1548476)) <= 0 {
		return (*byte)(memmap.PtrOff(6112660, uint32(noxRndCounter1.IntClamp(0, int32(dword_5d4594_1548476-1))*32)+0x174590))
	}
	v2 = (*uint8)(memmap.PtrOff(6112660, 1525156))
	v3 = int32(dword_5d4594_1548476)
	for {
		if *(*uint32)(unsafe.Pointer(v2)) > uint32(v1) {
			v16 = int32(*(*uint32)(unsafe.Pointer(v2)))
			v1 = int32(*(*uint32)(unsafe.Pointer(v2)))
		}
		v2 = (*uint8)(unsafe.Add(unsafe.Pointer(v2), 32))
		v3--
		if v3 == 0 {
			break
		}
	}
	if v1 == 0 {
		return (*byte)(memmap.PtrOff(6112660, uint32(noxRndCounter1.IntClamp(0, int32(dword_5d4594_1548476-1))*32)+0x174590))
	}
	v4 = 1
	v5 = (*uint8)(memmap.PtrOff(6112660, 1525156))
	v6 = int32(dword_5d4594_1548476)
	for {
		if dword_5d4594_1548476 > 1 {
			v7 = (*uint8)(memmap.PtrOff(6112660, 1525188))
			v8 = int32(dword_5d4594_1548476 - 1)
			for {
				if *(*uint32)(unsafe.Pointer(v5)) != *(*uint32)(unsafe.Pointer(v7)) {
					v4 = 0
				}
				v7 = (*uint8)(unsafe.Add(unsafe.Pointer(v7), 32))
				v8--
				if v8 == 0 {
					break
				}
			}
		}
		v5 = (*uint8)(unsafe.Add(unsafe.Pointer(v5), 32))
		v6--
		if v6 == 0 {
			break
		}
	}
	if v4 == 1 {
		v16++
	}
	v9 = 0
	v10 = 0
	v11 = (*uint8)(memmap.PtrOff(6112660, 1525132))
	for {
		if *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v11))), unsafe.Sizeof(uint32(0))*6))) < uint32(v16) && uint32(v10) != *memmap.PtrUint32(0x587000, 191880) && *(*uint32)(unsafe.Pointer(v11)) != *memmap.PtrUint32(6112660, *memmap.PtrUint32(0x587000, 191880)*32+0x17458C) && dword_5d4594_1548480-*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v11))), unsafe.Sizeof(uint32(0))*7))) > 4 {
			v9++
		}
		v10++
		v11 = (*uint8)(unsafe.Add(unsafe.Pointer(v11), 32))
		if v10 >= *(*int32)(unsafe.Pointer(&dword_5d4594_1548476)) {
			break
		}
	}
	v12 = noxRndCounter1.IntClamp(0, v9-1)
	v13 = 0
	v14 = 0
	if *(*int32)(unsafe.Pointer(&dword_5d4594_1548476)) <= 0 {
		return (*byte)(memmap.PtrOff(6112660, uint32(v12*32)+0x174590))
	}
	for i = (*uint8)(memmap.PtrOff(6112660, 1525132)); ; i = (*uint8)(unsafe.Add(unsafe.Pointer(i), 32)) {
		if *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(i))), unsafe.Sizeof(uint32(0))*6))) >= uint32(v16) || uint32(v14) == *memmap.PtrUint32(0x587000, 191880) || *(*uint32)(unsafe.Pointer(i)) == *memmap.PtrUint32(6112660, *memmap.PtrUint32(0x587000, 191880)*32+0x17458C) || dword_5d4594_1548480-*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(i))), unsafe.Sizeof(uint32(0))*7))) <= 4 {
			goto LABEL_36
		}
		if v13 == v12 {
			break
		}
		v13++
	LABEL_36:
		if func() int32 {
			p := &v14
			*p++
			return *p
		}() >= *(*int32)(unsafe.Pointer(&dword_5d4594_1548476)) {
			return (*byte)(memmap.PtrOff(6112660, uint32(v12*32)+0x174590))
		}
	}
	return (*byte)(memmap.PtrOff(6112660, uint32(v14*32)+0x174590))
}
func sub_4D11A0() {
	if *memmap.PtrUint32(6112660, 1548504) == 0 {
		nox_common_list_clear_425760((*nox_list_item_t)(memmap.PtrOff(6112660, 1548492)))
		*memmap.PtrUint32(6112660, 1548504) = 1
	}
}
func sub_4D11D0() {
	var (
		result *int32
		v1     *int32
		v2     *int32
	)
	result = (*int32)(unsafe.Pointer(nox_common_list_getFirstSafe_425890((*nox_list_item_t)(unsafe.Pointer(memmap.PtrInt32(6112660, 1548492))))))
	v1 = result
	if result != nil {
		for {
			v2 = (*int32)(unsafe.Pointer(nox_common_list_getNextSafe_4258A0((*nox_list_item_t)(unsafe.Pointer(v1)))))
			sub_425920((**uint32)(unsafe.Pointer(v1)))
			alloc.Free(unsafe.Pointer(v1))
			v1 = v2
			if v2 == nil {
				break
			}
		}
	}
}
func sub_4D1210(a1 int32) {
	var (
		result unsafe.Pointer
		v2     unsafe.Pointer
		v3     *uint32
	)
	result = unsafe.Pointer(uintptr(sub_4D12A0(a1)))
	if result == nil {
		result = unsafe.Pointer(nox_common_playerInfoFromNum_417090(a1))
		v2 = result
		if result != nil {
			v3 = (*uint32)(alloc.Calloc(1, 16))
			*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*3)) = uint32(uintptr(v2))
			nox_common_list_append_4258E0((*nox_list_item_t)(unsafe.Pointer(uintptr(int32(uintptr(memmap.PtrOff(6112660, 1548492)))))), (*nox_list_item_t)(unsafe.Pointer(v3)))
		}
	}
}
func sub_4D1250(a1 int32) *int32 {
	var (
		result *int32
		v2     *int32
	)
	result = (*int32)(unsafe.Pointer(nox_common_list_getFirstSafe_425890((*nox_list_item_t)(unsafe.Pointer(memmap.PtrInt32(6112660, 1548492))))))
	v2 = result
	if result != nil {
		for int32(*(*uint8)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(int32(0))*3)) + 2064)))) != a1 {
			result = (*int32)(unsafe.Pointer(nox_common_list_getNextSafe_4258A0((*nox_list_item_t)(unsafe.Pointer(v2)))))
			v2 = result
			if result == nil {
				return result
			}
		}
		sub_425920((**uint32)(unsafe.Pointer(v2)))
		alloc.Free(unsafe.Pointer(v2))
	}
	return result
}
func sub_4D12A0(a1 int32) int32 {
	var v1 *int32
	v1 = (*int32)(unsafe.Pointer(nox_common_list_getFirstSafe_425890((*nox_list_item_t)(unsafe.Pointer(memmap.PtrInt32(6112660, 1548492))))))
	if v1 == nil {
		return 0
	}
	for int32(*(*uint8)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(int32(0))*3)) + 2064)))) != a1 {
		v1 = (*int32)(unsafe.Pointer(nox_common_list_getNextSafe_4258A0((*nox_list_item_t)(unsafe.Pointer(v1)))))
		if v1 == nil {
			return 0
		}
	}
	return 1
}
func nox_xxx_mapSwitchLevel_4D12E0(a1 int32) {
	var (
		v1  int32
		v2  int32
		v3  int32
		v4  *uint32
		v5  int32
		v6  *uint32
		v11 [3]int32
	)
	v11[0] = 25
	v11[1] = 25
	v11[2] = 25
	noxflags.SetGame(noxflags.GameFlag20)
	sub_516F30()
	sub_421B10()
	sub_469B90(&v11[0])
	if noxflags.HasGame(noxflags.GameClient) {
		sub_4349C0((*uint32)(unsafe.Pointer(&v11[0])))
	}
	sub_511E60()
	if noxflags.HasGame(noxflags.GameModeCoop) {
		v1 = a1
		sub_4FCEB0(a1)
	} else {
		sub_4FCEB0(0)
		v1 = a1
	}
	nox_xxx_mapWall_4FF790()
	v2 = int32(uintptr(unsafe.Pointer(nox_xxx_getFirstPlayerUnit_4DA7C0())))
	if v2 != 0 {
		for {
			v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 748))))
			sub_4F7950(v2)
			*(*uint32)(unsafe.Pointer(uintptr(v3 + 296))) = 0
			nox_xxx_unitUnFreeze_4E7A60((*nox_object_t)(unsafe.Pointer(uintptr(v2))), 1)
			v4 = *(**uint32)(unsafe.Pointer(uintptr(v3 + 280)))
			*(*uint16)(unsafe.Pointer(uintptr(v3 + 160))) = 0
			if v4 != nil {
				nox_xxx_shopCancelSession_510DC0(v4)
			}
			*(*uint32)(unsafe.Pointer(uintptr(v3 + 280))) = 0
			if cgoFuncAddr(cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(v2 + 744))), (*func(*uint32) int32)(nil))) == cgoFuncAddr(nox_xxx_updatePlayerMonsterBot_4FAB20) {
				nox_xxx_playerBotCreate_4FA700(v2)
			}
			v2 = int32(uintptr(unsafe.Pointer(nox_xxx_getNextPlayerUnit_4DA7F0((*nox_object_t)(unsafe.Pointer(uintptr(v2)))))))
			if v2 == 0 {
				break
			}
		}
		v1 = a1
	}
	for {
		nox_xxx_unitsNewAddToList_4DAC00()
		sub_4E5BF0(v1)
		nox_xxx_spellCastByPlayer_4FEEF0()
		nox_xxx_finalizeDeletingUnits_4E5EC0()
		if noxServer.firstServerObjectUninited() == nil {
			break
		}
	}
	v5 = int32(uintptr(unsafe.Pointer(noxServer.firstServerObject())))
	if v5 != 0 {
		for {
			*(*uint32)(unsafe.Pointer(uintptr(v5 + 520))) = 0
			if nox_xxx_isUnit_4E5B50((*nox_object_t)(unsafe.Pointer(uintptr(v5)))) != 0 && int32(*(*uint8)(unsafe.Pointer(uintptr(v5 + 8))))&2 != 0 {
				v6 = *(**uint32)(unsafe.Pointer(uintptr(v5 + 748)))
				*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*309)) = math.MaxUint32
				*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*307)) = math.MaxUint32
				*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*317)) = math.MaxUint32
				*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*311)) = math.MaxUint32
				*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*313)) = math.MaxUint32
				*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*315)) = math.MaxUint32
				*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*319)) = math.MaxUint32
				*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*321)) = math.MaxUint32
				*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*323)) = math.MaxUint32
				*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*325)) = math.MaxUint32
				*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*98)) = 0
				*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*101)) = 0
			}
			v5 = int32(uintptr(unsafe.Pointer((*nox_object_t)(unsafe.Pointer(uintptr(v5))).Next())))
			if v5 == 0 {
				break
			}
		}
		v1 = a1
	}
	sub_50D1C0()
	for obj := (*nox_object_t)(firstServerObjectUpdatable2()); obj != nil; obj = obj.Next() {
		if sub_4E5B80(obj) != 0 {
			sub_4E81D0(obj)
		}
	}
	sub_4ECFE0()
	sub_511E20()
	nox_xxx_wall_410160()
	if v1 != 0 {
		nox_xxx_Fn_4FCAC0(v1, 1)
	} else {
		nox_xxx_Fn_4FCAC0(0, 0)
	}
	for j := int32(0); j < ptr_5D4594_2650668_cap*44; j += 44 {
		for k := int32(0); k < ptr_5D4594_2650668_cap; k++ {
			*(*uint8)(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer(*(**obj_5D4594_2650668_t)(unsafe.Add(unsafe.Pointer(ptr_5D4594_2650668), unsafe.Sizeof((*obj_5D4594_2650668_t)(nil))*uintptr(k)))))) + uint32(j)))) = 0
			*(*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer(*(**obj_5D4594_2650668_t)(unsafe.Add(unsafe.Pointer(ptr_5D4594_2650668), unsafe.Sizeof((*obj_5D4594_2650668_t)(nil))*uintptr(k)))))) + uint32(j) + 4))) = math.MaxUint8
			*(*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer(*(**obj_5D4594_2650668_t)(unsafe.Add(unsafe.Pointer(ptr_5D4594_2650668), unsafe.Sizeof((*obj_5D4594_2650668_t)(nil))*uintptr(k)))))) + uint32(j) + 24))) = math.MaxUint8
			nox_xxx_tileFreeTile_422200(int32(uint32(uintptr(unsafe.Pointer(*(**obj_5D4594_2650668_t)(unsafe.Add(unsafe.Pointer(ptr_5D4594_2650668), unsafe.Sizeof((*obj_5D4594_2650668_t)(nil))*uintptr(k)))))) + uint32(j) + 4))
			nox_xxx_tileFreeTile_422200(int32(uint32(uintptr(unsafe.Pointer(*(**obj_5D4594_2650668_t)(unsafe.Add(unsafe.Pointer(ptr_5D4594_2650668), unsafe.Sizeof((*obj_5D4594_2650668_t)(nil))*uintptr(k)))))) + uint32(j) + 24))
		}
	}
	sub_410730()
	nox_xxx_wallBreackableListClear_410810()
	nox_xxx_waypointDeleteAll_579DD0()
	nox_xxx_j_allocHitArray_511840()
	nox_xxx_decayDestroy_5117B0()
	sub_5112F0()
	sub_57C440()
	sub_57C000()
	sub_510E50()
	sub_4D1610()
	sub_4EC5B0()
	sub_50E360()
	sub_50D7E0()
	sub_4E4F80()
	noxflags.UnsetGame(noxflags.GameFlag20)
}
func nox_xxx____setargv_13_4D15B0() {
	*memmap.PtrUint32(6112660, 1548508) = 1
}
func sub_4D15C0() {
	*memmap.PtrUint32(6112660, 1548508) = 0
}
func sub_4D15D0() int32 {
	return int32(*memmap.PtrUint32(6112660, 1548508))
}
func nox_xxx_scavengerTreasureMax_4D1600() int32 {
	return int32(*memmap.PtrUint32(6112660, 1548528))
}
func sub_4D1610() {
	*memmap.PtrUint32(6112660, 1548528) = 0
}
func sub_4D1620() int32 {
	var result int32
	if dword_5d4594_1548532 != nil {
		result = int32(uint32(uintptr(dword_5d4594_1548532)) + 56)
	} else {
		result = 0
	}
	return result
}
func nox_xxx_parseGamedataBinPre_4D1630() int32 {
	var result int32
	result = nox_xxx_parseGamedataBin_419B30()
	if result != 0 {
		nox_xxx_loadWariorParams_424DF0()
		nox_xxx_loadBaseValues_57B200()
		result = bool2int(nox_xxx_loadMonsterBin_517010() != 0)
	}
	return result
}
func nox_xxx_servResetPlayers_4D23C0() int32 {
	var (
		i  *byte
		v2 int32
	)
	for i = (*byte)(unsafe.Pointer(nox_common_playerInfoGetFirst_416EA0())); i != nil; i = (*byte)(unsafe.Pointer(nox_common_playerInfoGetNext_416EE0((*nox_playerInfo)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(i))))))))) {
		if *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(i))), unsafe.Sizeof(uint32(0))*514))) != 0 {
			dword_5d4594_2649712 &= uint32(int32(^(1 << *(*byte)(unsafe.Add(unsafe.Pointer(i), 2064)))))
			v2 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(i))), unsafe.Sizeof(uint32(0))*514))))
			*(*byte)(unsafe.Add(unsafe.Pointer(i), 3676)) = 2
			nox_xxx_playerMakeDefItems_4EF7D0(v2, 1, 0)
			*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(i))), unsafe.Sizeof(uint32(0))*535))) = 0
			*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(i))), unsafe.Sizeof(uint32(0))*534))) = 0
		}
	}
	sub_51A100()
	noxflags.UnsetGame(noxflags.GameFlag18)
	nox_xxx_netGameSettings_4DEF00()
	nox_server_gameUnsetMapLoad_40A690()
	return 1
}
func sub_4D2FF0() *byte {
	var (
		result *byte
		v1     int32
	)
	result = (*byte)(unsafe.Pointer(uintptr(sub_40AA70(nil))))
	if result != nil || (func() bool {
		result = (*byte)(unsafe.Pointer(nox_common_playerInfoGetFirst_416EA0()))
		return (func() int32 {
			v1 = int32(uintptr(unsafe.Pointer(result)))
			return v1
		}()) == 0
	}()) {
		dword_5d4594_1548704 = 1
	} else {
		for {
			if int32(*(*uint8)(unsafe.Pointer(uintptr(v1 + 3680))))&1 != 0 {
				nox_xxx_netNeedTimestampStatus_4174F0(v1, 256)
			}
			result = (*byte)(unsafe.Pointer(nox_common_playerInfoGetNext_416EE0((*nox_playerInfo)(unsafe.Pointer(uintptr(v1))))))
			v1 = int32(uintptr(unsafe.Pointer(result)))
			if result == nil {
				break
			}
		}
		dword_5d4594_1548704 = 1
	}
	return result
}
func nox_xxx_netReportAllLatency_4D3050() *byte {
	var (
		result *byte
		v1     bool
		i      int32
		v3     [5]byte
	)
	v3[0] = 215
	if dword_5d4594_1548700 == 0 || (func() bool {
		result = (*byte)(unsafe.Pointer(nox_common_playerInfoGetNext_416EE0((*nox_playerInfo)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1548700))))))))
		return (func() uint32 {
			dword_5d4594_1548700 = uint32(uintptr(unsafe.Pointer(result)))
			return dword_5d4594_1548700
		}()) == 0
	}()) {
		result = (*byte)(unsafe.Pointer(nox_common_playerInfoGetFirst_416EA0()))
		dword_5d4594_1548700 = uint32(uintptr(unsafe.Pointer(result)))
	}
	if result != nil {
		for k := int32(0); *(*byte)(unsafe.Add(unsafe.Pointer(result), 2064)) != 31 && k < 32; k++ {
			v1 = sub_554240(int32(uint8(*(*byte)(unsafe.Add(unsafe.Pointer(result), 2064))))) == 0
			result = *(**byte)(unsafe.Pointer(&dword_5d4594_1548700))
			if !v1 {
				break
			}
			result = (*byte)(unsafe.Pointer(nox_common_playerInfoGetNext_416EE0((*nox_playerInfo)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1548700))))))))
			dword_5d4594_1548700 = uint32(uintptr(unsafe.Pointer(result)))
			if result == nil {
				result = (*byte)(unsafe.Pointer(nox_common_playerInfoGetFirst_416EA0()))
				dword_5d4594_1548700 = uint32(uintptr(unsafe.Pointer(result)))
			}
		}
		if result != nil {
			*(*uint16)(unsafe.Pointer(&v3[1])) = *((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(result))), unsafe.Sizeof(uint16(0))*1030)))
			*(*uint16)(unsafe.Pointer(&v3[3])) = uint16(int16(sub_554240(int32(uint8(*(*byte)(unsafe.Add(unsafe.Pointer(result), 2064)))))))
			result = (*byte)(unsafe.Pointer(nox_common_playerInfoGetFirst_416EA0()))
			for i = int32(uintptr(unsafe.Pointer(result))); result != nil; i = int32(uintptr(unsafe.Pointer(result))) {
				nox_netlist_addToMsgListCli_40EBC0(int32(*(*uint8)(unsafe.Pointer(uintptr(i + 2064)))), 1, (*uint8)(unsafe.Pointer(&v3[0])), 5)
				result = (*byte)(unsafe.Pointer(nox_common_playerInfoGetNext_416EE0((*nox_playerInfo)(unsafe.Pointer(uintptr(i))))))
			}
		}
	}
	return result
}
func sub_4D3310(a1 int32) int32 {
	var result int32
	result = a1
	*memmap.PtrUint32(6112660, 1548716) = uint32(a1)
	return result
}
func sub_4D3320() int32 {
	return int32(*memmap.PtrUint32(6112660, 1548716))
}
func sub_4D39F0(a3 *byte) int32 {
	var (
		v1     uint32
		v2     int8
		v3     *uint8
		v4     *byte
		v5     int32
		v6     int32
		v7     *uint8
		v8     uint32
		v9     *uint8
		v10    *byte
		v11    *uint8
		v12    int32
		v13    int32
		v14    int32
		v15    *byte
		v16    uint8
		result int32
		v18    [2048]byte
	)
	*memmap.PtrUint64(6112660, 1549772) = uint64(platformTicks())
	libc.MemSet(memmap.PtrOff(0x973F18, 35912), 0, 72)
	*memmap.PtrUint32(0x973F18, 35912) = 0
	*memmap.PtrUint32(0x973F18, 35916) = 0
	dword_5d4594_3835348 = 0
	dword_5d4594_3835356 = math.MaxUint8
	dword_5d4594_3835352 = 0
	dword_5d4594_3835360 = 0
	dword_5d4594_3835364 = 1
	dword_5d4594_3835368 = 1
	dword_5d4594_3835372 = 1
	*memmap.PtrUint32(0x973F18, 35948) = 0
	*memmap.PtrUint32(0x973F18, 35952) = 0
	*memmap.PtrUint32(0x973F18, 35956) = 0
	dword_5d4594_3835388 = 0
	dword_5d4594_3835392 = 1
	dword_5d4594_3835396 = math.MaxUint32
	*memmap.PtrUint8(0x973F18, 35972) = 2
	*memmap.PtrUint32(0x973F18, 35976) = 0
	*memmap.PtrUint32(0x973F18, 35980) = 0
	sub_51D0E0()
	if a3 != nil {
		v1 = uint32(libc.StrLen(a3) + 1)
		v2 = int8(uint8(v1))
		v1 >>= 2
		libc.MemCpy(memmap.PtrOff(0x973F18, 42152), unsafe.Pointer(a3), int(v1*4))
		v4 = (*byte)(unsafe.Add(unsafe.Pointer(a3), v1*4))
		v3 = (*uint8)(memmap.PtrOff(0x973F18, v1*4+0xA4A8))
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v1))), 0)) = uint8(v2)
		v5 = int32(*memmap.PtrUint32(0x587000, 197560))
		libc.MemCpy(unsafe.Pointer(v3), unsafe.Pointer(v4), int(v1&3))
		libc.StrCpy((*byte)(memmap.PtrOff(0x973F18, 36008)), a3)
		v6 = int32(*memmap.PtrUint32(0x587000, 197564))
		v7 = (*uint8)(memmap.PtrOff(0x973F18, uint32(libc.StrLen((*byte)(memmap.PtrOff(0x973F18, 36008)))+0x8CA8)))
		*(*uint32)(unsafe.Pointer(v7)) = *memmap.PtrUint32(0x587000, 197556)
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v7))), unsafe.Sizeof(uint32(0))*1))) = uint32(v5)
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v7))), unsafe.Sizeof(uint32(0))*2))) = uint32(v6)
		v8 = uint32(libc.StrLen(a3) + 1)
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v5))), 0)) = uint8(v8)
		v8 >>= 2
		libc.MemCpy(memmap.PtrOff(0x973F18, 38056), unsafe.Pointer(a3), int(v8*4))
		v10 = (*byte)(unsafe.Add(unsafe.Pointer(a3), v8*4))
		v9 = (*uint8)(memmap.PtrOff(0x973F18, v8*4+0x94A8))
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v8))), 0)) = uint8(int8(v5))
		*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint16(0))*0)) = *memmap.PtrUint16(0x587000, 197576)
		libc.MemCpy(unsafe.Pointer(v9), unsafe.Pointer(v10), int(v8&3))
		v11 = (*uint8)(memmap.PtrOff(0x973F18, uint32(libc.StrLen((*byte)(memmap.PtrOff(0x973F18, 38056)))+0x94A9)))
		v12 = int32(*memmap.PtrUint32(0x587000, 197572))
		*(*uint32)(unsafe.Pointer(func() *uint8 {
			p := &v11
			*p = (*uint8)(unsafe.Add(unsafe.Pointer(*p), -1))
			return *p
		}())) = *memmap.PtrUint32(0x587000, 197568)
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v6))), 0)) = *memmap.PtrUint8(0x587000, 197578)
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v11))), unsafe.Sizeof(uint32(0))*1))) = uint32(v12)
		*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v11))), unsafe.Sizeof(uint16(0))*4))) = uint16(int16(v5))
		*(*uint8)(unsafe.Add(unsafe.Pointer(v11), 10)) = uint8(int8(v6))
		nox_fs_remove((*byte)(memmap.PtrOff(0x973F18, 36008)))
		nox_fs_remove((*byte)(memmap.PtrOff(0x973F18, 38056)))
	} else {
		*memmap.PtrUint8(0x973F18, 42152) = *memmap.PtrUint8(6112660, 1549780)
		*memmap.PtrUint8(0x973F18, 40104) = *memmap.PtrUint8(6112660, 1549784)
		*memmap.PtrUint8(0x973F18, 36008) = *memmap.PtrUint8(6112660, 1549788)
		*memmap.PtrUint8(0x973F18, 38056) = *memmap.PtrUint8(6112660, 1549792)
	}
	nox_xxx_mapReset_5028E0()
	v13 = int32(*memmap.PtrUint32(0x587000, 197584))
	libc.StrCpy(&v18[0], a3)
	v14 = int32(*memmap.PtrUint32(0x587000, 197588))
	v15 = &v18[libc.StrLen(&v18[0])]
	*(*uint32)(unsafe.Pointer(v15)) = *memmap.PtrUint32(0x587000, 197580)
	v16 = *memmap.PtrUint8(0x587000, 197592)
	*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v15))), unsafe.Sizeof(uint32(0))*1))) = uint32(v13)
	*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v15))), unsafe.Sizeof(uint32(0))*2))) = uint32(v14)
	*(*byte)(unsafe.Add(unsafe.Pointer(v15), 12)) = byte(v16)
	sub_502A50(&v18[0])
	sub_502AB0(&v18[0])
	result = sub_502B10()
	dword_5d4594_3835312 = 0
	*memmap.PtrUint32(0x973F18, 35880) = 0
	*memmap.PtrUint32(6112660, 1599580) = 0
	return result
}
func sub_4D3C30() int64 {
	nox_xxx_free_503F40()
	sub_51D0E0()
	sub_502DF0()
	return int64(platformTicks())
}
func nox_xxx_tileInitdataClear_4D3C50(a1 unsafe.Pointer) {
	libc.MemCpy(memmap.PtrOff(0x973F18, 35912), a1, 72)
}
func sub_4D3C70() *uint8 {
	return (*uint8)(memmap.PtrOff(0x973F18, 35912))
}
func sub_4D3C80(a1 *uint32) *uint32 {
	var (
		result *uint32
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
	v2 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*3)))
	v10 = int32(*a1)
	v3 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*1)))
	v11 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*1)))
	if v2 < v3 {
		v3 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*3)))
		v10 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*2)))
		v11 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*3)))
	}
	if *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*5)) < uint32(v3) {
		v10 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*4)))
		v3 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*5)))
		v11 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*5)))
	}
	if *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*7)) < uint32(v3) {
		v10 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*6)))
		v11 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*7)))
	}
	v4 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*2)))
	v12 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*3)))
	if *a1 < uint32(v4) {
		v4 = int32(*a1)
		v12 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*1)))
	}
	if *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*4)) < uint32(v4) {
		v4 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*4)))
		v12 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*5)))
	}
	v5 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*6)))
	if v5 < v4 {
		v4 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*6)))
		v12 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*7)))
	}
	v6 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*4)))
	v7 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*5)))
	if *a1 > uint32(v6) {
		v6 = int32(*a1)
		v7 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*1)))
	}
	if *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*2)) > uint32(v6) {
		v7 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*3)))
		v6 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*2)))
	}
	if v5 > v6 {
		v6 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*6)))
		v7 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*7)))
	}
	v8 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*7)))
	v9 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*6)))
	if *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*1)) > uint32(v8) {
		v9 = int32(*a1)
		v8 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*1)))
	}
	if v2 > v8 {
		v9 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*2)))
		v8 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*3)))
	}
	if *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*5)) > uint32(v8) {
		v9 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*4)))
		v8 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*5)))
	}
	*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*6)) = uint32(v9)
	*a1 = uint32(v10)
	*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*2)) = uint32(v4)
	*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*7)) = uint32(v8)
	*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*1)) = uint32(v11)
	*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*4)) = uint32(v6)
	*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*5)) = uint32(v7)
	*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*3)) = uint32(v12)
	return result
}
func nox_xxx_mapGenFixCoords_4D3D90(a1 *float2, a2 *float2) int32 {
	if a1 == nil || a2 == nil {
		return 0
	}
	a2.field_0 = float32(float64(a1.field_4+a1.field_0)*0.70710677 + 2957.0)
	a2.field_4 = float32(float64(a1.field_4-a1.field_0)*0.70710677 + 2956.0)
	if float64(a2.field_0) <= 80.5 {
		a2.field_0 = 82.5
	}
	if float64(a2.field_4) <= 80.5 {
		a2.field_4 = 81.5
	}
	if float64(a2.field_0) >= 5853.5 {
		a2.field_0 = 5851.5
	}
	if float64(a2.field_4) >= 5853.5 {
		a2.field_4 = 5852.5
	}
	return 1
}
func sub_4D3E30(a1 *float2, a2 *float2) int32 {
	var result int32
	if a1 == nil || a2 == nil {
		return 0
	}
	if float64(a1.field_0) <= 80.5 {
		a1.field_0 = 82.5
	}
	if float64(a1.field_4) <= 80.5 {
		a1.field_4 = 81.5
	}
	if float64(a1.field_0) >= 5853.5 {
		a1.field_0 = 5851.5
	}
	if float64(a1.field_4) >= 5853.5 {
		a1.field_4 = 5852.5
	}
	result = 1
	a2.field_0 = float32((float64(a1.field_0) - 1.0 - float64(a1.field_4)) * 0.70710677)
	a2.field_4 = float32((float64(a1.field_4+a1.field_0) - 5912.0) * 0.70710677)
	return result
}
func sub_4D3ED0(a1 *float2) int32 {
	var (
		result int32
		v2     float32
		v3     float32
	)
	result = 0
	if a1 != nil {
		v2 = float32(float64(a1.field_4+a1.field_0)*0.70710677 + 2957.0)
		v3 = float32(float64(a1.field_4-a1.field_0)*0.70710677 + 2956.0)
		if float64(v2) > 80.5 && float64(v3) > 80.5 && float64(v2) < 5853.5 && float64(v3) < 5853.5 {
			result = 1
		}
	}
	return result
}
func sub_4D3F60(a1 int32) int32 {
	if a1 != 1 && a1 != 0 {
		return 0
	}
	*memmap.PtrUint32(0x973F18, 35980) = 1
	return 1
}
func sub_4D3F80(a1 int32) int32 {
	var result int32
	switch a1 {
	case 0:
		result = 1
	case 1:
		result = 2
	case 2:
		result = 5
	case 3:
		result = 0
	case 5:
		result = 8
	case 6:
		result = 3
	case 7:
		result = 6
	case 8:
		result = 7
	default:
		result = -1
	}
	return result
}
func sub_4D3FF0(a1 int32) int32 {
	var result int32
	switch a1 {
	case 0:
		result = 3
	case 1:
		result = 0
	case 2:
		result = 1
	case 3:
		result = 6
	case 5:
		result = 2
	case 6:
		result = 7
	case 7:
		result = 8
	case 8:
		result = 5
	default:
		result = -1
	}
	return result
}
func sub_4D4060(a1 *float2, a2 *float2, a3 int32) *float2 {
	var (
		v3     float64
		v4     float64
		v5     int32
		v6     int32
		v7     int32
		v8     float64
		result *float2
		v10    float64
		v11    int32
		v12    float32
		v13    *float2
	)
	v3 = float64(a1.field_0) + 11.5
	v4 = float64(a1.field_4) + 11.5
	v13 = (*float2)(unsafe.Pointer(uintptr(int64(v3 * 0.021739131))))
	v12 = float32(v4)
	v11 = int32(int64(v4 * 0.021739131))
	v5 = int32(int64(v3)) % 46
	v6 = int32(int64(v12) % 46)
	v7 = a3 * 2
	if v5 <= v6 {
		if v7-v5 <= v6 {
			result = a2
			a2.field_0 = float32(float64(int32(uintptr(unsafe.Pointer(v13))))*46.0 + 11.0)
			v10 = float64(v11)
			goto LABEL_8
		}
		v8 = float64(int32(uint32(uint64(int64(v3*0.021739131))-1))) * 46.0
	LABEL_6:
		result = a2
		a2.field_0 = float32(v8 + 34.0)
		a2.field_4 = float32(float64(v11)*46.0 + 11.0)
		return result
	}
	v8 = float64(int32(uintptr(unsafe.Pointer(v13)))) * 46.0
	if v7-v5 <= v6 {
		goto LABEL_6
	}
	result = a2
	a2.field_0 = float32(v8 + 11.0)
	v10 = float64(int32(uint32(uint64(int64(v4*0.021739131)) - 1)))
LABEL_8:
	result.field_4 = float32(v10*46.0 + 34.0)
	return result
}
func sub_4D4160() {
	var (
		v0               int32
		v1               int32
		v2               int32
		v3               int32
		v4               int32
		v5               *byte
		v6               *byte
		v8               float2
		FileName         [1024]byte
		ExistingFileName [1024]byte
	)
	noxflags.SetGame(noxflags.GameFlag23)
	nox_xxx_tileGetDefByName_51D4D0(CString("BrickLight"))
	v8.field_4 = 0.0
	v8.field_0 = 0.0
	v0 = 20
	for {
		v1 = 20
		for {
			sub_51D5E0(&v8.field_0)
			v1--
			v8.field_0 = float32(float64(v8.field_0) + 32.526913)
			if v1 == 0 {
				break
			}
		}
		v0--
		v8.field_0 = 0.0
		v8.field_4 = float32(float64(v8.field_4) + 32.526913)
		if v0 == 0 {
			break
		}
	}
	v8.field_4 = 0.0
	v8.field_0 = 0.0
	v2 = sub_5029A0(CString("CBD-aTest-C"))
	nox_xxx_mapgenSaveMap_503830(v2)
	sub_503B30(&v8)
	v8.field_4 = 65.053825
	v8.field_0 = 65.053825
	v3 = sub_5029A0(CString("CBD-aTest-C"))
	nox_xxx_mapgenSaveMap_503830(v3)
	sub_503B30(&v8)
	v8.field_4 = 195.16147
	v8.field_0 = 195.16147
	v4 = sub_5029A0(CString("CBD-aTest2-C"))
	nox_xxx_mapgenSaveMap_503830(v4)
	sub_503B30(&v8)
	v5 = nox_fs_root()
	nox_sprintf(&FileName[0], CString("%s\\nc.obj"), v5)
	v6 = nox_fs_root()
	nox_sprintf(&ExistingFileName[0], CString("%s\\blend.obj"), v6)
	nox_fs_remove(&FileName[0])
	nox_fs_move(&ExistingFileName[0], &FileName[0])
	noxflags.UnsetGame(noxflags.GameFlag23)
}
func sub_4D42D0(a1 int32) int32 {
	var result int32
	result = a1
	dword_5d4594_1550916 = uint32(a1)
	return result
}
func sub_4D42E0(a1 *byte) uint32 {
	var result uint32
	result = uint32(libc.StrLen(a1) + 1)
	libc.MemCpy(memmap.PtrOff(0x587000, 197860), unsafe.Pointer(a1), int(result))
	return result
}
func nox_xxx_getRandMapName_4D4310() *byte {
	return (*byte)(memmap.PtrOff(0x587000, 197860))
}
func nox_xxx_mapGenStart_4D4320() int32 {
	var (
		v0               int32
		v1               int32
		v2               *byte
		v3               *byte
		result           int32
		i                int32
		v6               *byte
		v7               *byte
		v8               *byte
		v9               *byte
		v10              *byte
		PathName         [1024]byte
		FileName         [1024]byte
		ExistingFileName [1024]byte
	)
	v0 = 100
	nox_xxx_mapSwitchLevel_4D12E0(1)
	noxflags.SetGame(noxflags.GameFlag23)
	*memmap.PtrUint32(6112660, 1550924) = uint32(uintptr(unsafe.Pointer(nox_get_and_zero_server_objects_4DA3C0())))
	libc.MemSet(memmap.PtrOff(0x973F18, 2408), 0, 1464)
	for {
		v1 = nox_xxx_mapGenStep_4D44E0()
		if v1 == 1 {
			break
		}
		if v1 != 0 {
			if func() int32 {
				p := &v0
				*p--
				return *p
			}() != 0 {
				continue
			}
		}
		return 0
	}
	v2 = nox_fs_root()
	nox_sprintf(&FileName[0], CString("%s\\nc.obj"), v2)
	v3 = nox_fs_root()
	nox_sprintf(&ExistingFileName[0], CString("%s\\blend.obj"), v3)
	if nox_fs_access(&FileName[0], 0) == -1 || (func() int32 {
		result = bool2int(nox_fs_remove(&FileName[0]))
		return result
	}()) != 0 {
		if nox_fs_access(&ExistingFileName[0], 0) == -1 || (func() int32 {
			result = bool2int(nox_fs_move(&ExistingFileName[0], &FileName[0]))
			return result
		}()) != 0 {
			for i = int32(uintptr(unsafe.Pointer(noxServer.firstServerObject()))); i != 0; i = int32(uintptr(unsafe.Pointer((*nox_object_t)(unsafe.Pointer(uintptr(i))).Next()))) {
				*(*uint32)(unsafe.Pointer(uintptr(i + 44))) = 0
			}
			nox_xxx_mapGenMakeInfo_4D5DB0(int32(uintptr(memmap.PtrOff(0x973F18, 2408))))
			v9 = nox_xxx_getRandMapName_4D4310()
			v6 = nox_fs_root()
			nox_sprintf(&PathName[0], CString("%s\\Maps\\$%s"), v6, v9)
			nox_fs_mkdir(&PathName[0])
			v10 = nox_xxx_getRandMapName_4D4310()
			v8 = nox_xxx_getRandMapName_4D4310()
			v7 = nox_fs_root()
			nox_sprintf(&PathName[0], CString("%s\\Maps\\$%s\\$%s.map"), v7, v8, v10)
			if nox_xxx_mapSaveMap_51E010(&PathName[0], 1) == 0 {
				return 0
			}
			nox_xxx_mapSwitchLevel_4D12E0(1)
			nox_set_server_objects_4DA3E0(unsafe.Pointer(uintptr(*memmap.PtrInt32(6112660, 1550924))))
			*memmap.PtrUint32(6112660, 1550924) = 0
			noxflags.UnsetGame(noxflags.GameFlag23)
			result = 1
		}
	}
	return result
}
func nox_xxx_mapGenStep_4D44E0() int32 {
	var (
		v0     *byte
		result int32
		v2     int32
		v3     int32
		v4     *float32
		v5     float64
		i      *byte
		j      *int32
		k      *int32
		a2     float2
	)
	dword_5d4594_1550916 = 0
	v0 = (*byte)(unsafe.Pointer(uintptr(sub_57C490(CString("theme")))))
	if v0 != nil {
		sub_4D42E0(v0)
	}
	sub_526C40(0)
	sub_51D100(0)
	result = nox_xxx_mapGenReadTheme_51E260(memmap.PtrInt32(6112660, 1549796), int32(uintptr(memmap.PtrOff(0x587000, 197860))))
	if result != 0 {
		*memmap.PtrUint32(6112660, 1549864) = uint32(int32(int64(float64(*mem_getFloatPtr(6112660, 1549860)) * 0.030743772)))
		nox_xxx_mapGenSetRngSeed_526AB0(*memmap.PtrUint32(6112660, 1549872))
		sub_526950()
		result = nox_xxx_mapgenAllocBuffer_5213E0()
		if result != 0 {
			result = sub_520EA0(int32(uintptr(memmap.PtrOff(6112660, 1549796))))
			if result != 0 {
				nox_xxx_mapGenMkSmallRoom_4D4F40((*uint32)(memmap.PtrOff(6112660, 1549796)))
				if nox_xxx_mapGen_InPrefab1_525D20(int32(uintptr(memmap.PtrOff(6112660, 1549796)))) != 0 {
					sub_4D52F0()
					if nox_xxx_mapGen_InPrefab2_5266F0(int32(uintptr(memmap.PtrOff(6112660, 1549796)))) != 0 {
						if nox_xxx_mapGenPlacePrefabs_526830(int32(uintptr(memmap.PtrOff(6112660, 1549796)))) == 0 {
							v2 = 0
						LABEL_25:
							nox_xxx_mapGenFreeTopRoom_521A40()
							nox_xxx_mapgenFreeBuffer_521400()
							sub_520F80()
							sub_520D50((*uint32)(memmap.PtrOff(6112660, 1549796)))
							return v2
						}
						sub_5259F0(*(*int32)(unsafe.Pointer(&dword_5d4594_1550916)), 0, 0.0)
						sub_525AF0(*(*int32)(unsafe.Pointer(&dword_5d4594_1550916)))
						if *memmap.PtrUint32(6112660, 1549980) != 0 {
							v3 = int32(int64(float64(*mem_getFloatPtr(6112660, 1549860)) * 0.030743772))
							v4 = nox_xxx_mapGenMakeRoomStruct_521940(v3*2+1, v3*2+1)
							v5 = float64(-v3) * 32.526913
							a2.field_0 = float32(v5)
							a2.field_4 = float32(v5)
							nox_xxx_mapGenSetRoomPos_521880((*uint32)(unsafe.Pointer(v4)), &a2)
							for i = (*byte)(nox_xxx_mapGenGetTopRoom_521710()); i != nil; i = (*byte)(unsafe.Pointer(uintptr(sub_521720(int32(uintptr(unsafe.Pointer(i))))))) {
								sub_521BC0(int32(uintptr(unsafe.Pointer(v4))), (*float2)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(i), 20)))), *((*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(i))), unsafe.Sizeof(float32(0))*7))), *((*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(i))), unsafe.Sizeof(float32(0))*8))))
							}
							sub_524070(int32(uintptr(memmap.PtrOff(6112660, 1549796))), int32(uintptr(unsafe.Pointer(v4))))
							nox_xxx_gen_524E00(int32(uintptr(memmap.PtrOff(6112660, 1549796))), int32(uintptr(unsafe.Pointer(v4))))
							nox_xxx_mapgen_522340(int32(uintptr(memmap.PtrOff(6112660, 1549796))), int32(uintptr(unsafe.Pointer(v4))))
							sub_521A10(unsafe.Pointer(v4))
						}
						if nox_xxx_mapGenMakeRooms_524310(int32(uintptr(memmap.PtrOff(6112660, 1549796)))) != 0 {
							for j = (*int32)(nox_xxx_mapGenGetTopRoom_521710()); j != nil; j = (*int32)(unsafe.Pointer(uintptr(sub_521720(int32(uintptr(unsafe.Pointer(j))))))) {
								if nox_xxx_mapGenCheckRoomType_5238F0(j) != 0 {
									nox_xxx_mapGenSetFlags_5235F0(-100)
									nox_xxx_gen_524E00(int32(uintptr(memmap.PtrOff(6112660, 1549796))), int32(uintptr(unsafe.Pointer(j))))
								}
							}
							for k = (*int32)(nox_xxx_mapGenGetTopRoom_521710()); k != nil; k = (*int32)(unsafe.Pointer(uintptr(sub_521720(int32(uintptr(unsafe.Pointer(k))))))) {
								if nox_xxx_mapGenCheckRoomType_5238F0(k) == 0 {
									nox_xxx_mapGenSetFlags_5235F0(-100)
									nox_xxx_gen_524E00(int32(uintptr(memmap.PtrOff(6112660, 1549796))), int32(uintptr(unsafe.Pointer(k))))
								}
							}
							sub_522D30(int32(uintptr(memmap.PtrOff(6112660, 1549796))))
							nox_xxx_mapgen_Doors_4D4790()
							nox_xxx_mapGenTryNextRoom_522F40((*uint32)(memmap.PtrOff(6112660, 1549796)))
							nox_xxx_mapGenGetTopRoom_521710()
							nox_xxx_mapGenFinishPopulate_5228B0_mapgen_populate(int32(uintptr(memmap.PtrOff(6112660, 1549796))))
							v2 = 1
							goto LABEL_25
						}
					}
				}
				v2 = 2
				goto LABEL_25
			}
		}
	}
	return result
}
func nox_xxx_mapgen_Doors_4D4790() *float32 {
	var (
		result *float32
		v1     *float32
		v2     *int32
		v3     int32
		v4     int32
		v5     int32
		v6     *int32
		v7     int32
		v8     *float32
		v9     float64
		v10    float64
		v11    *float32
		v12    float64
		v13    float64
		v14    float64
		v15    float64
		v16    int32
		v17    int32
		v18    *int32
		v19    int32
		v20    *float32
		v21    float64
		v22    float64
		v23    float64
		v24    *float32
		v25    float64
		v26    float64
		v27    float64
		v28    int32
		v29    int32 = 0
		v30    int32
		v31    int32
		v32    int32
		v33    int32 = 0
		v34    float2
		a2     float2
		v36    float32
		v37    float32
		v38    float2
		v39    float2
		a1     int2
	)
	result = (*float32)(nox_xxx_mapGenGetTopRoom_521710())
	v1 = result
	v2 = nil
	if result == nil {
		return result
	}
	v3 = v33
	for {
		if *(*uint32)(unsafe.Pointer(v1)) != 1 {
			goto LABEL_117
		}
		nox_xxx_mapGenSetFlags_5235F0(-100)
		v4 = 0
		v32 = 0
		for {
			if int32(*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(v1))), v4))), 216)))) == 0 && nox_xxx_mapGenRandFunc_526AC0(1, 100) <= *memmap.PtrInt32(6112660, 1549848) {
				switch v4 {
				case 0:
					fallthrough
				case 1:
					a1.field_0 = int32(*(*float32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(float32(0))*1)))
					if v4 == 1 {
						a1.field_4 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v1))), unsafe.Sizeof(uint32(0))*2))) + *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v1))), unsafe.Sizeof(uint32(0))*4))))
					} else {
						a1.field_4 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v1))), unsafe.Sizeof(uint32(0))*2))) - 1)
					}
					v31 = 0
					if *((*int32)(unsafe.Add(unsafe.Pointer((*int32)(unsafe.Pointer(v1))), unsafe.Sizeof(int32(0))*3))) <= 0 {
						goto LABEL_115
					}
					for {
						v17 = sub_521290(&a1)
						v18 = (*int32)(unsafe.Pointer(uintptr(v17)))
						if v17 != 0 {
							if (*int32)(unsafe.Pointer(uintptr(v17))) != v2 {
								goto LABEL_72
							}
							if int32(*(*uint8)(unsafe.Pointer(uintptr(v17 + 52))))&2 != 0 {
								v29++
							}
						}
						if (*int32)(unsafe.Pointer(uintptr(v17))) != v2 || (func() bool {
							v19 = v31
							return uint32(v31) == *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v1))), unsafe.Sizeof(uint32(0))*3)))-1
						}()) {
						LABEL_72:
							if v2 != nil && v29 >= 3 {
								v34.field_0 = float32(float64((v3+v31)/2)*32.526913 + float64(*(*float32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(float32(0))*9))))
								if v32 == 1 {
									v34.field_4 = *(*float32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(float32(0))*12))
								} else {
									v34.field_4 = *(*float32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(float32(0))*10))
								}
								sub_527030(&v34)
								v34.field_0 = float32(float64(v34.field_0) - 16.263456)
								if v29 < 4 {
									nox_xxx_mapGenGetObjID_527940(CString("ArchedDoor"))
								} else {
									nox_xxx_mapGenGetObjID_527940(CString("ArchedHalfDoor"))
								}
								v20 = nox_xxx_mapGenPlaceObj_5279B0((*float2)(unsafe.Pointer(&v34.field_0)))
								if v20 != nil {
									nox_xxx_mapGenOrientObj_527C60(int32(uintptr(unsafe.Pointer(v20))), 5)
								}
								a2.field_0 = v34.field_0
								if v32 == 1 {
									a2.field_4 = float32(float64(v34.field_4) - 32.526913)
								} else {
									a2.field_4 = v34.field_4
								}
								sub_521BC0(int32(uintptr(unsafe.Pointer(v1))), &a2, 32.526913, 32.526913)
								if nox_xxx_mapGenCheckRoomType_5238F0(v2) != 0 {
									v21 = float64(a2.field_4)
									if v32 == 1 {
										v22 = v21 + 32.526913
									} else {
										v22 = v21 - 32.526913
									}
									a2.field_4 = float32(v22)
									sub_521BC0(int32(uintptr(unsafe.Pointer(v2))), &a2, 32.526913, 32.526913)
								}
								v23 = float64(v34.field_0) + 16.263456
								if v29 < 4 {
									v37 = v34.field_4
								} else {
									v34.field_0 = float32(v23 + 32.526913)
									sub_527030(&v34)
									v34.field_0 = float32(float64(v34.field_0) + 16.263456)
									nox_xxx_mapGenGetObjID_527940(CString("ArchedHalfDoor"))
									v24 = nox_xxx_mapGenPlaceObj_5279B0((*float2)(unsafe.Pointer(&v34.field_0)))
									if v24 != nil {
										nox_xxx_mapGenOrientObj_527C60(int32(uintptr(unsafe.Pointer(v24))), 3)
									}
									v25 = float64(v34.field_0) - 32.526913
									v37 = v34.field_4
									v36 = float32(v25)
									a2.field_0 = float32(v25)
									if v32 == 1 {
										a2.field_4 = float32(float64(v34.field_4) - 32.526913)
									} else {
										a2.field_4 = v34.field_4
									}
									sub_521BC0(int32(uintptr(unsafe.Pointer(v1))), &a2, 32.526913, 32.526913)
									if nox_xxx_mapGenCheckRoomType_5238F0(v2) != 0 {
										v26 = float64(a2.field_4)
										if v32 == 1 {
											v27 = v26 + 32.526913
										} else {
											v27 = v26 - 32.526913
										}
										a2.field_4 = float32(v27)
										sub_521BC0(int32(uintptr(unsafe.Pointer(v2))), &a2, 32.526913, 32.526913)
									}
									v23 = float64(v36)
								}
								v39.field_0 = float32(v23)
								v39.field_4 = float32(float64(v37) - 32.526913)
								v38.field_0 = float32(v23)
								v38.field_4 = float32(float64(v37) + 32.526913)
								sub_522C80(&v39.field_0)
								sub_522C80(&v38.field_0)
								sub_51D3F0(&v39, &v38)
								sub_51D3F0(&v38, &v39)
								if v32 == 1 {
									sub_522CA0(int32(uintptr(unsafe.Pointer(v1))), &v39.field_0)
									if *v2 == 1 {
										sub_522CA0(int32(uintptr(unsafe.Pointer(v2))), &v38.field_0)
									}
								} else {
									sub_522CA0(int32(uintptr(unsafe.Pointer(v1))), &v38.field_0)
									if *v2 == 1 {
										sub_522CA0(int32(uintptr(unsafe.Pointer(v2))), &v39.field_0)
									}
								}
								sub_521900(int32(uintptr(unsafe.Pointer(v1))), int32(uintptr(unsafe.Pointer(v2))), v32)
								v28 = sub_523960(v32)
								sub_521900(int32(uintptr(unsafe.Pointer(v2))), int32(uintptr(unsafe.Pointer(v1))), v28)
							}
							v19 = v31
							v3 = v31
							v29 = 1
							v2 = v18
							if v18 != nil && *v18 == 1 && v32 == 1 {
								v2 = nil
							}
						}
						a1.field_0++
						v31 = v19 + 1
						if uint32(v19+1) >= *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v1))), unsafe.Sizeof(uint32(0))*3))) {
							goto LABEL_114
						}
					}
					fallthrough
				case 2:
					fallthrough
				case 3:
					if v4 == 2 {
						a1.field_0 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v1))), unsafe.Sizeof(uint32(0))*1))) + *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v1))), unsafe.Sizeof(uint32(0))*3))))
					} else {
						a1.field_0 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v1))), unsafe.Sizeof(uint32(0))*1))) - 1)
					}
					v30 = 0
					a1.field_4 = int32(*(*float32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(float32(0))*2)))
					if *((*int32)(unsafe.Add(unsafe.Pointer((*int32)(unsafe.Pointer(v1))), unsafe.Sizeof(int32(0))*4))) <= 0 {
						goto LABEL_115
					}
				default:
					goto LABEL_115
				}
				for {
					v5 = sub_521290(&a1)
					v6 = (*int32)(unsafe.Pointer(uintptr(v5)))
					if v5 != 0 {
						if (*int32)(unsafe.Pointer(uintptr(v5))) != v2 {
							goto LABEL_19
						}
						if int32(*(*uint8)(unsafe.Pointer(uintptr(v5 + 52))))&2 != 0 {
							v29++
						}
					}
					if (*int32)(unsafe.Pointer(uintptr(v5))) != v2 || (func() bool {
						v7 = v30
						return uint32(v30) == *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v1))), unsafe.Sizeof(uint32(0))*4)))-1
					}()) {
					LABEL_19:
						if v2 != nil && v29 >= 3 {
							if v32 == 2 {
								v34.field_0 = *(*float32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(float32(0))*11))
							} else {
								v34.field_0 = *(*float32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(float32(0))*9))
							}
							v34.field_4 = float32(float64((v3+v30)/2)*32.526913 + float64(*(*float32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(float32(0))*10))))
							sub_527030(&v34)
							v34.field_4 = float32(float64(v34.field_4) - 16.263456)
							if v29 < 4 {
								nox_xxx_mapGenGetObjID_527940(CString("ArchedDoor"))
							} else {
								nox_xxx_mapGenGetObjID_527940(CString("ArchedHalfDoor"))
							}
							v8 = nox_xxx_mapGenPlaceObj_5279B0((*float2)(unsafe.Pointer(&v34.field_0)))
							if v8 != nil {
								nox_xxx_mapGenOrientObj_527C60(int32(uintptr(unsafe.Pointer(v8))), 7)
							}
							if v32 == 2 {
								a2.field_0 = float32(float64(v34.field_0) - 32.526913)
							} else {
								a2.field_0 = v34.field_0
							}
							a2.field_4 = v34.field_4
							sub_521BC0(int32(uintptr(unsafe.Pointer(v1))), &a2, 32.526913, 32.526913)
							if nox_xxx_mapGenCheckRoomType_5238F0(v2) != 0 {
								v9 = float64(a2.field_0)
								if v32 == 2 {
									v10 = v9 + 32.526913
								} else {
									v10 = v9 - 32.526913
								}
								a2.field_0 = float32(v10)
								sub_521BC0(int32(uintptr(unsafe.Pointer(v2))), &a2, 32.526913, 32.526913)
							}
							if v29 < 4 {
								v15 = float64(v34.field_4) + 16.263456
								v36 = v34.field_0
							} else {
								v34.field_4 = float32(float64(v34.field_4) + 16.263456 + 32.526913)
								sub_527030(&v34)
								v34.field_4 = float32(float64(v34.field_4) + 16.263456)
								nox_xxx_mapGenGetObjID_527940(CString("ArchedHalfDoor"))
								v11 = nox_xxx_mapGenPlaceObj_5279B0((*float2)(unsafe.Pointer(&v34.field_0)))
								if v11 != nil {
									nox_xxx_mapGenOrientObj_527C60(int32(uintptr(unsafe.Pointer(v11))), 1)
								}
								v12 = float64(v34.field_4) - 32.526913
								v36 = v34.field_0
								v37 = float32(v12)
								if v32 == 2 {
									a2.field_0 = float32(float64(v34.field_0) - 32.526913)
								} else {
									a2.field_0 = v34.field_0
								}
								a2.field_4 = float32(v12)
								sub_521BC0(int32(uintptr(unsafe.Pointer(v1))), &a2, 32.526913, 32.526913)
								if nox_xxx_mapGenCheckRoomType_5238F0(v2) != 0 {
									v13 = float64(a2.field_0)
									if v32 == 2 {
										v14 = v13 + 32.526913
									} else {
										v14 = v13 - 32.526913
									}
									a2.field_0 = float32(v14)
									sub_521BC0(int32(uintptr(unsafe.Pointer(v2))), &a2, 32.526913, 32.526913)
								}
								v15 = float64(v37)
							}
							v39.field_0 = float32(float64(v36) - 32.526913)
							v39.field_4 = float32(v15)
							v38.field_0 = float32(float64(v36) + 32.526913)
							v38.field_4 = float32(v15)
							sub_522C80(&v39.field_0)
							sub_522C80(&v38.field_0)
							sub_51D3F0(&v39, &v38)
							sub_51D3F0(&v38, &v39)
							if v32 == 2 {
								sub_522CA0(int32(uintptr(unsafe.Pointer(v1))), &v39.field_0)
								if *v2 == 1 {
									sub_522CA0(int32(uintptr(unsafe.Pointer(v2))), &v38.field_0)
								}
							} else {
								sub_522CA0(int32(uintptr(unsafe.Pointer(v1))), &v38.field_0)
								if *v2 == 1 {
									sub_522CA0(int32(uintptr(unsafe.Pointer(v2))), &v39.field_0)
								}
							}
							sub_521900(int32(uintptr(unsafe.Pointer(v1))), int32(uintptr(unsafe.Pointer(v2))), v32)
							v16 = sub_523960(v32)
							sub_521900(int32(uintptr(unsafe.Pointer(v2))), int32(uintptr(unsafe.Pointer(v1))), v16)
						}
						v7 = v30
						v29 = 1
						v3 = v30
						v2 = v6
						if v6 != nil && *v6 == 1 && v32 == 3 {
							v2 = nil
						}
					}
					a1.field_4++
					v30 = v7 + 1
					if uint32(v7+1) >= *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v1))), unsafe.Sizeof(uint32(0))*4))) {
					LABEL_114:
						v4 = v32
						break
					}
				}
			}
		LABEL_115:
			v32 = func() int32 {
				p := &v4
				*p++
				return *p
			}()
			if v4 >= 4 {
				break
			}
			v2 = nil
		}
		v2 = nil
	LABEL_117:
		result = (*float32)(unsafe.Pointer(uintptr(sub_521720(int32(uintptr(unsafe.Pointer(v1)))))))
		v1 = result
		if result == nil {
			break
		}
	}
	return result
}
func nox_xxx_mapGenMkSmallRoom_4D4F40(a1 *uint32) *float32 {
	var (
		v1     *float32
		result *float32
		v3     float64
		v4     int32
		v5     *int32
		v6     *float32
		v7     int32
		v8     float64
		v9     int32
		v10    int32
		v11    int32
		v12    int32
		v13    *int32
		v14    int32
		v15    int32
		v16    *float32
		v17    int32
		v18    int32
		v19    int32
		v20    int32
		v21    int32
		v22    float2
		v24    int32
		v25    *int32
		v26    *float32
		v27    int32
		v28    float2
		v30    [128]int32
	)
	v1 = nil
	dword_5d4594_1550912 = 0
	if *a1 == 1 {
		v26 = nil
		v24 = 0
		v22.field_0 = 0.0
		v22.field_4 = 0.0
		v25 = memmap.PtrInt32(0x587000, 197924)
	LABEL_5:
		v4 = 0
		v5 = &v30[v24]
		v27 = *v25
		for {
			result = (*float32)(alloc.Calloc(1, 376))
			v6 = result
			*v5 = int32(uintptr(unsafe.Pointer(result)))
			if result == nil {
				break
			}
			v7 = v27
			*(*uint32)(unsafe.Pointer(v6)) = uint32(v27)
			switch v7 {
			case 2:
				v22.field_4 = float32(float64(v22.field_4) - 162.63457)
				*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v6))), unsafe.Sizeof(uint32(0))*3))) = 4
				*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v6))), unsafe.Sizeof(uint32(0))*4))) = 5
				sub_521900(int32(uintptr(unsafe.Pointer(v6))), int32(uintptr(unsafe.Pointer(v1))), 1)
				sub_521900(int32(uintptr(unsafe.Pointer(v1))), int32(uintptr(unsafe.Pointer(v6))), 0)
			case 3:
				if *(*uint32)(unsafe.Pointer(v1)) == 4 {
					v22.field_0 = float32(float64(v22.field_0) + 32.526913)
					v8 = float64(v22.field_4) + 130.10765
				} else {
					v8 = float64(v22.field_4) + 162.63457
				}
				v22.field_4 = float32(v8)
				*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v6))), unsafe.Sizeof(uint32(0))*3))) = 4
				*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v6))), unsafe.Sizeof(uint32(0))*4))) = 5
				sub_521900(int32(uintptr(unsafe.Pointer(v6))), int32(uintptr(unsafe.Pointer(v1))), 0)
				sub_521900(int32(uintptr(unsafe.Pointer(v1))), int32(uintptr(unsafe.Pointer(v6))), 1)
			case 4:
				if v1 != nil {
					v22.field_0 = float32(float64(v22.field_0) + 162.63457)
					sub_521900(int32(uintptr(unsafe.Pointer(v6))), int32(uintptr(unsafe.Pointer(v1))), 3)
					sub_521900(int32(uintptr(unsafe.Pointer(v1))), int32(uintptr(unsafe.Pointer(v6))), 2)
				}
				*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v6))), unsafe.Sizeof(uint32(0))*3))) = 5
				*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v6))), unsafe.Sizeof(uint32(0))*4))) = 4
			case 5:
				if *(*uint32)(unsafe.Pointer(v1)) == 3 {
					v22.field_4 = float32(float64(v22.field_4) + 32.526913)
				}
				v22.field_0 = float32(float64(v22.field_0) - 162.63457)
				*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v6))), unsafe.Sizeof(uint32(0))*3))) = 5
				*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v6))), unsafe.Sizeof(uint32(0))*4))) = 4
				sub_521900(int32(uintptr(unsafe.Pointer(v6))), int32(uintptr(unsafe.Pointer(v1))), 2)
				sub_521900(int32(uintptr(unsafe.Pointer(v1))), int32(uintptr(unsafe.Pointer(v6))), 3)
				if v4 == 5 {
					v26 = v6
				}
			default:
			}
			v28.field_0 = float32(float64(v22.field_0) - 878.22662)
			v28.field_4 = float32(float64(v22.field_4) - 878.22662)
			*(*float32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(float32(0))*7)) = float32(float64(*((*int32)(unsafe.Add(unsafe.Pointer((*int32)(unsafe.Pointer(v6))), unsafe.Sizeof(int32(0))*3)))) * 32.526913)
			*(*float32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(float32(0))*8)) = float32(float64(*((*int32)(unsafe.Add(unsafe.Pointer((*int32)(unsafe.Pointer(v6))), unsafe.Sizeof(int32(0))*4)))) * 32.526913)
			nox_xxx_mapGenSetRoomPos_521880((*uint32)(unsafe.Pointer(v6)), &v28)
			nox_xxx_mapGenAddNewRoom_521730((*uint32)(unsafe.Pointer(v6)))
			v1 = v6
			v9 = v24 + 1
			v5 = (*int32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(int32(0))*1))
			v4++
			v24++
			if v4 >= 10 {
				v25 = (*int32)(unsafe.Add(unsafe.Pointer(v25), unsafe.Sizeof(int32(0))*1))
				if int32(uintptr(unsafe.Pointer(v25))) < int32(uintptr(memmap.PtrOff(0x587000, 197940))) {
					goto LABEL_5
				}
				v10 = v30[0]
				v11 = v9
				v12 = v30[v9-1]
				v13 = &v30[v11]
				sub_521900(v12, v30[0], 2)
				sub_521900(v10, v12, 3)
				v14 = int32(uintptr(unsafe.Pointer(v26)))
				v15 = 0
				v22.field_0 = *(*float32)(unsafe.Add(unsafe.Pointer(v26), unsafe.Sizeof(float32(0))*5))
				v22.field_4 = *(*float32)(unsafe.Add(unsafe.Pointer(v26), unsafe.Sizeof(float32(0))*8)) + *(*float32)(unsafe.Add(unsafe.Pointer(v26), unsafe.Sizeof(float32(0))*6))
				for {
					result = (*float32)(alloc.Calloc(1, 376))
					v16 = result
					*v13 = int32(uintptr(unsafe.Pointer(result)))
					if result == nil {
						break
					}
					*(*uint32)(unsafe.Pointer(result)) = 3
					*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(result))), unsafe.Sizeof(uint32(0))*3))) = 4
					*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(result))), unsafe.Sizeof(uint32(0))*4))) = 5
					sub_521900(int32(uintptr(unsafe.Pointer(result))), v14, 0)
					sub_521900(v14, int32(uintptr(unsafe.Pointer(v16))), 1)
					*(*float32)(unsafe.Add(unsafe.Pointer(v16), unsafe.Sizeof(float32(0))*7)) = float32(float64(*((*int32)(unsafe.Add(unsafe.Pointer((*int32)(unsafe.Pointer(v16))), unsafe.Sizeof(int32(0))*3)))) * 32.526913)
					*(*float32)(unsafe.Add(unsafe.Pointer(v16), unsafe.Sizeof(float32(0))*8)) = float32(float64(*((*int32)(unsafe.Add(unsafe.Pointer((*int32)(unsafe.Pointer(v16))), unsafe.Sizeof(int32(0))*4)))) * 32.526913)
					nox_xxx_mapGenSetRoomPos_521880((*uint32)(unsafe.Pointer(v16)), &v22)
					nox_xxx_mapGenAddNewRoom_521730((*uint32)(unsafe.Pointer(v16)))
					v14 = int32(uintptr(unsafe.Pointer(v16)))
					v17 = v24 + 1
					v13 = (*int32)(unsafe.Add(unsafe.Pointer(v13), unsafe.Sizeof(int32(0))*1))
					v15++
					v22.field_4 = float32(float64(v22.field_4) + 162.63457)
					v24++
					if v15 >= 8 {
						v18 = v17
						v19 = v17 / 5
						if v17/5 < 1 {
							v19 = 1
						}
						v20 = nox_xxx_mapGenRandFunc_526AC0(0, v19)
						for {
							v21 = v30[v20]
							*(*uint32)(unsafe.Pointer(uintptr(v21 + 84))) = dword_5d4594_1550912
							dword_5d4594_1550912 = uint32(v21)
							v20 += nox_xxx_mapGenRandFunc_526AC0(1, v19)
							if v20 >= v18 {
								break
							}
						}
						result = (*float32)(unsafe.Pointer(uintptr(v30[v18-1])))
						dword_5d4594_1550916 = uint32(v30[v18-1])
						return result
					}
				}
				return result
			}
		}
	} else {
		result = nox_xxx_mapGenPrepareRoom_521990(int32(uintptr(unsafe.Pointer(a1))))
		dword_5d4594_1550916 = uint32(uintptr(unsafe.Pointer(result)))
		if result != nil {
			v3 = float64(int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*17))))
			v22.field_0 = 0.0
			v22.field_4 = float32(v3*32.526913 - float64(*(*float32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(float32(0))*8))) + 97.580734)
			nox_xxx_mapGenSetRoomPos_521880((*uint32)(unsafe.Pointer(result)), &v22)
			nox_xxx_mapGenAddNewRoom_521730(*(**uint32)(unsafe.Pointer(&dword_5d4594_1550916)))
			result = *(**float32)(unsafe.Pointer(&dword_5d4594_1550916))
			*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1550916 + 84))) = dword_5d4594_1550912
			dword_5d4594_1550912 = dword_5d4594_1550916
		}
	}
	return result
}
func sub_4D52F0() {
	var v0 *uint32
	v0 = *(**uint32)(unsafe.Pointer(&dword_5d4594_1550912))
	if dword_5d4594_1550912 != 0 {
		for {
			switch *v0 {
			case 1:
				sub_4D5350(v0, 0, 0, 0, 0)
			case 2:
				fallthrough
			case 3:
				sub_4D5350(v0, 0, 0, int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v0), unsafe.Sizeof(uint32(0))*3))), 0)
			case 4:
				fallthrough
			case 5:
				sub_4D5350(v0, 0, 0, int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v0), unsafe.Sizeof(uint32(0))*4))), 0)
			default:
			}
			v0 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v0), unsafe.Sizeof(uint32(0))*21)))))
			if v0 == nil {
				break
			}
		}
	}
}
func sub_4D5350(a1 *uint32, a2 int32, a3 int32, a4 int32, a5 int32) int32 {
	var (
		v5     int32
		result int32
	)
	v5 = a2 + 1
	if a2+1 >= *memmap.PtrInt32(6112660, 1549868) {
		return 0
	}
	nox_xxx_mapGenSetFlags_5235F0(-101)
	if *a1 == 1 {
		result = nox_xxx_mapGenFillRoom_4D53B0(int32(uintptr(unsafe.Pointer(a1))), v5, a3, a4, a5)
	} else {
		result = sub_4D5630(int32(uintptr(unsafe.Pointer(a1))), v5, a3, a4, a5)
	}
	return result
}
func nox_xxx_mapGenFillRoom_4D53B0(a1 int32, a2 int32, a3 int32, a4 int32, a5 int32) int32 {
	var (
		v5  *float32
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
		v18 *float32
		v19 float64
		v20 float64
		v21 int32
		v22 int32
		v23 *int32
		v24 int32
		v25 int32
		v26 int32
		v28 int32
		a2a float2
		v30 int32
		v31 int32
	)
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	v5 = (*float32)(unsafe.Pointer(uintptr(a1)))
	v6 = sub_5218B0(a1, 0)
	if v6 != 0 {
		v7 = 0
	} else {
		v7 = 2
	}
	if v6 != 0 {
		v30 = 0
	} else {
		v30 = 2
	}
	v8 = sub_5218B0(a1, 1)
	if v8 != 0 {
		v9 = 0
	} else {
		v9 = 3
	}
	if v8 != 0 {
		v31 = 0
	} else {
		v31 = 3
	}
	v10 = sub_5218B0(a1, 2)
	if v10 != 0 {
		v11 = 0
	} else {
		v11 = 4
	}
	if v10 != 0 {
		v32 = 0
	} else {
		v32 = 4
	}
	v12 = -bool2int(sub_5218B0(a1, 3) != 0)
	*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v12))), 0)) = uint8(int8(v12 & 251))
	v13 = v12 + 5
	v33 = v13
	if v7 != 0 || v9 != 0 || v11 != 0 || v13 != 0 {
		v14 = nox_xxx_mapGenRandFunc_526AC0(0, 3)
		v34 = 0
		for {
			v15 = (v14 + 1) % 4
			v16 = *((*int32)(unsafe.Add(unsafe.Pointer(&v30), unsafe.Sizeof(int32(0))*uintptr(v15))))
			v28 = v15
			if v16 == 0 {
				goto LABEL_6
			}
			v17 = nox_xxx_mapGenRandFunc_526AC0(int32(*memmap.PtrUint32(6112660, 1549808)-*memmap.PtrUint32(6112660, 1549812)), int32(*memmap.PtrUint32(6112660, 1549812)+*memmap.PtrUint32(6112660, 1549808)))
			v18 = nox_xxx_mapGenMakeHall_523EC0(int32(uintptr(memmap.PtrOff(6112660, 1549796))), v16, v17)
			if v18 == nil {
				return 0
			}
			switch *(*uint32)(unsafe.Pointer(v18)) {
			case 2:
				a2a.field_0 = float32(sub_521B00(int32(uintptr(unsafe.Pointer(v5))), int32(uintptr(unsafe.Pointer(v18)))))
				v19 = float64(*(*float32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(float32(0))*6)) - *(*float32)(unsafe.Add(unsafe.Pointer(v18), unsafe.Sizeof(float32(0))*8)))
				goto LABEL_15
			case 3:
				a2a.field_0 = float32(sub_521B00(int32(uintptr(unsafe.Pointer(v5))), int32(uintptr(unsafe.Pointer(v18)))))
				v19 = float64(*(*float32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(float32(0))*8)) + *(*float32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(float32(0))*6)))
				goto LABEL_15
			case 4:
				v20 = float64(*(*float32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(float32(0))*7)) + *(*float32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(float32(0))*5)))
				goto LABEL_14
			case 5:
				v20 = float64(*(*float32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(float32(0))*5)) - *(*float32)(unsafe.Add(unsafe.Pointer(v18), unsafe.Sizeof(float32(0))*7)))
			LABEL_14:
				a2a.field_0 = float32(v20)
				v19 = sub_521B30(int32(uintptr(unsafe.Pointer(v5))), int32(uintptr(unsafe.Pointer(v18))))
			LABEL_15:
				a2a.field_4 = float32(v19)
			default:
			}
			nox_xxx_mapGenSetRoomPos_521880((*uint32)(unsafe.Pointer(v18)), &a2a)
			if sub_5217A0(int32(uintptr(memmap.PtrOff(6112660, 1549796))), int32(uintptr(unsafe.Pointer(v18)))) == 0 {
			LABEL_17:
				sub_521A10(unsafe.Pointer(v18))
				goto LABEL_26
			}
			v21 = sub_523920(int32(*(*uint32)(unsafe.Pointer(v18))))
			v22 = sub_523960(v21)
			sub_521900(int32(uintptr(unsafe.Pointer(v18))), int32(uintptr(unsafe.Pointer(v5))), v22)
			v23 = (*int32)(unsafe.Pointer(uintptr(sub_521200(int32(uintptr(unsafe.Pointer(v18)))))))
			v24 = int32(uintptr(unsafe.Pointer(v23)))
			if v23 != nil {
				break
			}
			nox_xxx_mapGenAddNewRoom_521730((*uint32)(unsafe.Pointer(v18)))
			if sub_4D5350((*uint32)(unsafe.Pointer(v18)), a2, 1, v17, int32(uintptr(unsafe.Pointer(v5)))) != 0 {
				goto LABEL_25
			}
			sub_521760(int32(uintptr(unsafe.Pointer(v18))))
			sub_521A10(unsafe.Pointer(v18))
		LABEL_26:
			if func() int32 {
				p := &v34
				*p++
				return *p
			}() >= 8 {
				return 1
			}
		LABEL_6:
			v14 = v28
		}
		if nox_xxx_mapGenCheckRoomType_5238F0(v23) != 0 || int32(*(*uint8)(unsafe.Pointer(uintptr(v24 + 52))))&2 != 0 || v24 == a5 || nox_xxx_mapGenRandFunc_526AC0(1, 100) > *(*int32)(unsafe.Pointer(&dword_5d4594_1549844)) || sub_523A10(int32(uintptr(unsafe.Pointer(v18))), (*float32)(unsafe.Pointer(uintptr(v24)))) == 0 {
			goto LABEL_17
		}
		nox_xxx_mapGenAddNewRoom_521730((*uint32)(unsafe.Pointer(v18)))
		v25 = sub_523920(int32(*(*uint32)(unsafe.Pointer(v18))))
		sub_521A70(int32(uintptr(unsafe.Pointer(v18))), v24, v25)
	LABEL_25:
		v26 = sub_523920(int32(*(*uint32)(unsafe.Pointer(v18))))
		sub_521900(int32(uintptr(unsafe.Pointer(v5))), int32(uintptr(unsafe.Pointer(v18))), v26)
		goto LABEL_26
	}
	return 1
}
func sub_4D5630(a1 int32, a2 int32, a3 int32, a4 int32, a5 int32) int32 {
	var (
		v5  int32
		v6  int32
		v7  int32
		v8  *float32
		v9  float64
		v10 float64
		v11 int32
		v13 *float32
		v14 *float32
		v15 float64
		v16 float64
		v17 int32
		v18 *float32
		v19 int32
		v20 int32
		v21 int32
		v22 *float32
		v23 float64
		v24 float64
		v25 float64
		v26 int32
		v27 *float32
		v28 int32
		v29 int32
		v30 int32
		v31 *float32
		v32 *int32
		v33 float64
		v34 float64
		v35 float64
		v36 float64
		v37 int32
		v38 int32
		v39 *float32
		v40 int32
		v41 int32
		v42 int32
		v43 int32
		v44 int32
		v45 int32
		a2a float2
		v47 int32
		v48 int32
		v49 int32
	)
	v5 = a1
	v43 = 0
	v44 = 0
	if a3 >= *memmap.PtrInt32(6112660, 1549816) || (func() bool {
		v6 = sub_4D5D20((*uint32)(unsafe.Pointer(uintptr(a1))))
		v45 = v6
		return v6 == 1
	}()) {
		v7 = 0
		for {
			v8 = nox_xxx_mapGenPrepareRoom_521990(int32(uintptr(memmap.PtrOff(6112660, 1549796))))
			if v8 == nil {
				return 0
			}
			switch *(*uint32)(unsafe.Pointer(uintptr(a1))) {
			case 2:
				a2a.field_0 = float32(sub_521B60(int32(uintptr(unsafe.Pointer(v8))), a1))
				v9 = float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 24))) - *(*float32)(unsafe.Add(unsafe.Pointer(v8), unsafe.Sizeof(float32(0))*8)))
				goto LABEL_11
			case 3:
				a2a.field_0 = float32(sub_521B60(int32(uintptr(unsafe.Pointer(v8))), a1))
				v9 = float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 32))) + *(*float32)(unsafe.Pointer(uintptr(a1 + 24))))
				goto LABEL_11
			case 4:
				v10 = float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 28))) + *(*float32)(unsafe.Pointer(uintptr(a1 + 20))))
				goto LABEL_10
			case 5:
				v10 = float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 20))) - *(*float32)(unsafe.Add(unsafe.Pointer(v8), unsafe.Sizeof(float32(0))*7)))
			LABEL_10:
				a2a.field_0 = float32(v10)
				v9 = sub_521B90(int32(uintptr(unsafe.Pointer(v8))), a1)
			LABEL_11:
				a2a.field_4 = float32(v9)
			default:
			}
			nox_xxx_mapGenSetRoomPos_521880((*uint32)(unsafe.Pointer(v8)), &a2a)
			if sub_521820(int32(uintptr(memmap.PtrOff(6112660, 1549796))), int32(uintptr(unsafe.Pointer(v8)))) == 0 {
				sub_521A10(unsafe.Pointer(v8))
				if func() int32 {
					p := &v7
					*p++
					return *p
				}() < 10 {
					continue
				}
			}
			if v7 == 10 {
				return 0
			}
			nox_xxx_mapGenAddNewRoom_521730((*uint32)(unsafe.Pointer(v8)))
			v11 = sub_523920(int32(*(*uint32)(unsafe.Pointer(uintptr(a1)))))
			sub_521A70(a1, int32(uintptr(unsafe.Pointer(v8))), v11)
			sub_4D5350((*uint32)(unsafe.Pointer(v8)), a2, 0, 0, int32(uintptr(unsafe.Pointer(v8))))
			return 1
		}
	}
	if v6 != 2 && v6 != 8 && v6 != 32 && v6 != 64 {
		goto LABEL_43
	}
	v13 = nox_xxx_mapGenMakeHall_523EC0(int32(uintptr(memmap.PtrOff(6112660, 1549796))), int32(*memmap.PtrUint32(0x587000, *(*uint32)(unsafe.Pointer(uintptr(a1)))*4+0x304B4)), a4)
	v14 = v13
	if v13 == nil {
		return 0
	}
	switch *(*uint32)(unsafe.Pointer(uintptr(a1))) {
	case 2:
		v15 = float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 20))) - *(*float32)(unsafe.Add(unsafe.Pointer(v13), unsafe.Sizeof(float32(0))*7)))
		a2a.field_4 = *(*float32)(unsafe.Pointer(uintptr(a1 + 24)))
		a2a.field_0 = float32(v15)
	case 3:
		a2a.field_0 = *(*float32)(unsafe.Pointer(uintptr(a1 + 28))) + *(*float32)(unsafe.Pointer(uintptr(a1 + 20)))
		v16 = float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 32))) + *(*float32)(unsafe.Pointer(uintptr(a1 + 24))) - *(*float32)(unsafe.Add(unsafe.Pointer(v13), unsafe.Sizeof(float32(0))*8)))
		goto LABEL_26
	case 4:
		a2a.field_0 = *(*float32)(unsafe.Pointer(uintptr(a1 + 28))) + *(*float32)(unsafe.Pointer(uintptr(a1 + 20))) - *(*float32)(unsafe.Add(unsafe.Pointer(v13), unsafe.Sizeof(float32(0))*7))
		v16 = float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 24))) - *(*float32)(unsafe.Add(unsafe.Pointer(v13), unsafe.Sizeof(float32(0))*8)))
		goto LABEL_26
	case 5:
		v16 = float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 32))) + *(*float32)(unsafe.Pointer(uintptr(a1 + 24))))
		a2a.field_0 = *(*float32)(unsafe.Pointer(uintptr(a1 + 20)))
	LABEL_26:
		a2a.field_4 = float32(v16)
	default:
	}
	nox_xxx_mapGenSetRoomPos_521880((*uint32)(unsafe.Pointer(v13)), &a2a)
	v47 = sub_5239B0(int32(*(*uint32)(unsafe.Pointer(uintptr(a1)))))
	sub_521900(int32(uintptr(unsafe.Pointer(v14))), v5, v47)
	if sub_5217A0(int32(uintptr(memmap.PtrOff(6112660, 1549796))), int32(uintptr(unsafe.Pointer(v14)))) == 0 {
		goto LABEL_100
	}
	v17 = sub_521200(int32(uintptr(unsafe.Pointer(v14))))
	v18 = (*float32)(unsafe.Pointer(uintptr(v17)))
	if v17 == 0 {
		v43 = 1
		v19 = 0
		goto LABEL_34
	}
	if *(*uint32)(unsafe.Pointer(uintptr(v17))) != 1 || int32(*(*uint8)(unsafe.Pointer(uintptr(v17 + 52))))&2 != 0 || v17 == a5 || nox_xxx_mapGenRandFunc_526AC0(1, 100) > *(*int32)(unsafe.Pointer(&dword_5d4594_1549844)) {
	LABEL_100:
		v43 = 0
		goto LABEL_41
	}
	v43 = sub_523A10(int32(uintptr(unsafe.Pointer(v14))), v18)
	v19 = 1
	if v43 == 0 {
		goto LABEL_41
	}
LABEL_34:
	nox_xxx_mapGenAddNewRoom_521730((*uint32)(unsafe.Pointer(v14)))
	if v19 == 0 {
		if sub_4D5350((*uint32)(unsafe.Pointer(v14)), a2, a3+1, a4, a5) != 0 {
			goto LABEL_39
		}
		sub_521760(int32(uintptr(unsafe.Pointer(v14))))
		v43 = 0
	LABEL_41:
		sub_521A10(unsafe.Pointer(v14))
		if v45 == 2 || v45 == 8 {
			return 0
		}
		goto LABEL_43
	}
	v20 = sub_523920(int32(*(*uint32)(unsafe.Pointer(v14))))
	sub_521A70(int32(uintptr(unsafe.Pointer(v14))), int32(uintptr(unsafe.Pointer(v18))), v20)
LABEL_39:
	v21 = sub_523960(v47)
	sub_521900(v5, int32(uintptr(unsafe.Pointer(v14))), v21)
LABEL_43:
	if v45 != 4 && v45 != 16 && v45 != 32 && v45 != 64 {
		goto LABEL_71
	}
	v22 = nox_xxx_mapGenMakeHall_523EC0(int32(uintptr(memmap.PtrOff(6112660, 1549796))), int32(*memmap.PtrUint32(0x587000, *(*uint32)(unsafe.Pointer(uintptr(v5)))*4+0x304CC)), a4)
	if v22 == nil {
		return 0
	}
	switch *(*uint32)(unsafe.Pointer(uintptr(v5))) {
	case 2:
		v23 = float64(*(*float32)(unsafe.Pointer(uintptr(v5 + 28))) + *(*float32)(unsafe.Pointer(uintptr(v5 + 20))))
		a2a.field_4 = *(*float32)(unsafe.Pointer(uintptr(v5 + 24)))
		a2a.field_0 = float32(v23)
	case 3:
		a2a.field_0 = *(*float32)(unsafe.Pointer(uintptr(v5 + 20))) - *(*float32)(unsafe.Add(unsafe.Pointer(v22), unsafe.Sizeof(float32(0))*7))
		v24 = float64(*(*float32)(unsafe.Pointer(uintptr(v5 + 32))) + *(*float32)(unsafe.Pointer(uintptr(v5 + 24))))
		goto LABEL_53
	case 4:
		a2a.field_0 = *(*float32)(unsafe.Pointer(uintptr(v5 + 28))) + *(*float32)(unsafe.Pointer(uintptr(v5 + 20))) - *(*float32)(unsafe.Add(unsafe.Pointer(v22), unsafe.Sizeof(float32(0))*7))
		v25 = float64(*(*float32)(unsafe.Pointer(uintptr(v5 + 32))) + *(*float32)(unsafe.Pointer(uintptr(v5 + 24))))
		goto LABEL_54
	case 5:
		v24 = float64(*(*float32)(unsafe.Pointer(uintptr(v5 + 24))))
		a2a.field_0 = *(*float32)(unsafe.Pointer(uintptr(v5 + 20)))
	LABEL_53:
		v25 = v24 - float64(*(*float32)(unsafe.Add(unsafe.Pointer(v22), unsafe.Sizeof(float32(0))*8)))
	LABEL_54:
		a2a.field_4 = float32(v25)
	default:
	}
	nox_xxx_mapGenSetRoomPos_521880((*uint32)(unsafe.Pointer(v22)), &a2a)
	v48 = sub_523970(int32(*(*uint32)(unsafe.Pointer(uintptr(v5)))))
	sub_521900(int32(uintptr(unsafe.Pointer(v22))), v5, v48)
	if sub_5217A0(int32(uintptr(memmap.PtrOff(6112660, 1549796))), int32(uintptr(unsafe.Pointer(v22)))) == 0 {
		goto LABEL_101
	}
	v26 = sub_521200(int32(uintptr(unsafe.Pointer(v22))))
	v27 = (*float32)(unsafe.Pointer(uintptr(v26)))
	if v26 == 0 {
		v44 = 1
		v28 = 0
		goto LABEL_62
	}
	if *(*uint32)(unsafe.Pointer(uintptr(v26))) != 1 || int32(*(*uint8)(unsafe.Pointer(uintptr(v26 + 52))))&2 != 0 || v26 == a5 || nox_xxx_mapGenRandFunc_526AC0(1, 100) > *(*int32)(unsafe.Pointer(&dword_5d4594_1549844)) {
	LABEL_101:
		v44 = 0
		goto LABEL_69
	}
	v44 = sub_523A10(int32(uintptr(unsafe.Pointer(v22))), v27)
	v28 = 1
	if v44 == 0 {
		goto LABEL_69
	}
LABEL_62:
	nox_xxx_mapGenAddNewRoom_521730((*uint32)(unsafe.Pointer(v22)))
	if v28 == 0 {
		if sub_4D5350((*uint32)(unsafe.Pointer(v22)), a2, a3+1, a4, a5) != 0 {
			goto LABEL_67
		}
		sub_521760(int32(uintptr(unsafe.Pointer(v22))))
		v44 = 0
	LABEL_69:
		sub_521A10(unsafe.Pointer(v22))
		if v45 == 4 || v45 == 16 {
			return 0
		}
		goto LABEL_71
	}
	v29 = sub_523920(int32(*(*uint32)(unsafe.Pointer(v22))))
	sub_521A70(int32(uintptr(unsafe.Pointer(v22))), int32(uintptr(unsafe.Pointer(v27))), v29)
LABEL_67:
	v30 = sub_523960(v48)
	sub_521900(v5, int32(uintptr(unsafe.Pointer(v22))), v30)
LABEL_71:
	if v43 != 0 || v44 != 0 {
		if v45 != 32 && v45 != 8 && v45 != 16 {
			return 1
		}
		v31 = nox_xxx_mapGenMakeHall_523EC0(int32(uintptr(memmap.PtrOff(6112660, 1549796))), int32(*(*uint32)(unsafe.Pointer(uintptr(v5)))), a4)
		v32 = (*int32)(unsafe.Pointer(v31))
		if v31 != nil {
			switch *(*uint32)(unsafe.Pointer(uintptr(v5))) {
			case 2:
				v33 = float64(*(*float32)(unsafe.Pointer(uintptr(v5 + 24))))
				a2a.field_0 = *(*float32)(unsafe.Pointer(uintptr(v5 + 20)))
				a2a.field_4 = float32(v33 - float64(*(*float32)(unsafe.Add(unsafe.Pointer(v31), unsafe.Sizeof(float32(0))*8))))
			case 3:
				v34 = float64(*(*float32)(unsafe.Pointer(uintptr(v5 + 32))) + *(*float32)(unsafe.Pointer(uintptr(v5 + 24))))
				a2a.field_0 = *(*float32)(unsafe.Pointer(uintptr(v5 + 20)))
				a2a.field_4 = float32(v34)
			case 4:
				v35 = float64(*(*float32)(unsafe.Pointer(uintptr(v5 + 28))) + *(*float32)(unsafe.Pointer(uintptr(v5 + 20))))
				a2a.field_4 = *(*float32)(unsafe.Pointer(uintptr(v5 + 24)))
				a2a.field_0 = float32(v35)
			case 5:
				v36 = float64(*(*float32)(unsafe.Pointer(uintptr(v5 + 20))) - *(*float32)(unsafe.Add(unsafe.Pointer(v31), unsafe.Sizeof(float32(0))*7)))
				a2a.field_4 = *(*float32)(unsafe.Pointer(uintptr(v5 + 24)))
				a2a.field_0 = float32(v36)
			default:
			}
			nox_xxx_mapGenSetRoomPos_521880((*uint32)(unsafe.Pointer(v31)), &a2a)
			v37 = sub_523920(int32(*(*uint32)(unsafe.Pointer(uintptr(v5)))))
			v49 = sub_523960(v37)
			sub_521900(int32(uintptr(unsafe.Pointer(v32))), v5, v49)
			if sub_5217A0(int32(uintptr(memmap.PtrOff(6112660, 1549796))), int32(uintptr(unsafe.Pointer(v32)))) != 0 {
				v38 = sub_521200(int32(uintptr(unsafe.Pointer(v32))))
				v39 = (*float32)(unsafe.Pointer(uintptr(v38)))
				if v38 == 0 {
					v40 = 0
				LABEL_89:
					nox_xxx_mapGenAddNewRoom_521730((*uint32)(unsafe.Pointer(v32)))
					if v40 != 0 {
						v41 = sub_523920(*v32)
						sub_521A70(int32(uintptr(unsafe.Pointer(v32))), int32(uintptr(unsafe.Pointer(v39))), v41)
					LABEL_96:
						v42 = sub_523960(v49)
						sub_521900(v5, int32(uintptr(unsafe.Pointer(v32))), v42)
						return 1
					}
					if sub_4D5350((*uint32)(unsafe.Pointer(v32)), a2, a3+1, a4, a5) != 0 {
						goto LABEL_96
					}
					sub_521760(int32(uintptr(unsafe.Pointer(v32))))
					goto LABEL_94
				}
				if *(*uint32)(unsafe.Pointer(uintptr(v38))) == 1 && (int32(*(*uint8)(unsafe.Pointer(uintptr(v38 + 52))))&2) == 0 && v38 != a5 && nox_xxx_mapGenRandFunc_526AC0(1, 100) <= *(*int32)(unsafe.Pointer(&dword_5d4594_1549844)) {
					v40 = 1
					if sub_523A10(int32(uintptr(unsafe.Pointer(v32))), v39) != 0 {
						goto LABEL_89
					}
				}
			}
		LABEL_94:
			sub_521A10(unsafe.Pointer(v32))
			return 1
		}
	}
	return 0
}
func sub_4D5D20(a1 *uint32) int32 {
	var result int32
	if *a1 == 2 || *a1 == 3 {
		if *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*3)) >= *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*4)) {
			return 1
		}
	} else if *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*4)) >= *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*3)) {
		return 1
	}
	if nox_xxx_mapGenRandFunc_526AC0(1, 100) > *memmap.PtrInt32(6112660, 1549820) {
		if nox_xxx_mapGenRandFunc_526AC0(1, 100) > *memmap.PtrInt32(6112660, 1549824) {
			if nox_xxx_mapGenRandFunc_526AC0(1, 100) >= 50 {
				result = 4
			} else {
				result = 2
			}
		} else {
			result = 1
		}
	} else {
		nox_xxx_mapGenRandFunc_526AC0(1, 100)
		result = 32
	}
	return result
}
func nox_xxx_mapGenMakeInfo_4D5DB0(a1 int32) uint32 {
	var (
		result     uint32
		SystemTime _SYSTEMTIME
		DateStr    [1024]byte
	)
	libc.StrCpy((*byte)(unsafe.Pointer(uintptr(a1))), CString("Generated Map"))
	libc.StrCpy((*byte)(unsafe.Pointer(uintptr(a1+64))), CString("Generated Map"))
	*(*uint16)(unsafe.Pointer(uintptr(a1 + 576))) = *memmap.PtrUint16(0x587000, 198380)
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 592))) = *memmap.PtrUint32(0x587000, 198384)
	libc.StrCpy((*byte)(unsafe.Pointer(uintptr(a1+656))), CString("http://www.westwood.com"))
	libc.StrCpy((*byte)(unsafe.Pointer(uintptr(a1+720))), CString("http://www.westwood.com"))
	libc.StrCpy((*byte)(unsafe.Pointer(uintptr(a1+848))), CString("http://www.westwood.com"))
	libc.StrCpy((*byte)(unsafe.Pointer(uintptr(a1+976))), CString("Generated Map"))
	libc.StrCpy((*byte)(unsafe.Pointer(uintptr(a1+1232))), CString("Westwood Studios"))
	compatGetLocalTime(&SystemTime)
	compatGetDateFormatA(2048, 0, (*SYSTEMTIME)(unsafe.Pointer(&SystemTime)), (*byte)(memmap.PtrOff(0x587000, 198496)), &DateStr[0], 1024)
	result = uint32(libc.StrLen(&DateStr[0]) + 1)
	libc.MemCpy(unsafe.Pointer(uintptr(a1+1360)), unsafe.Pointer(&DateStr[0]), int(result))
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 1392))) = 3
	return result
}
func nox_xxx_mapGenStartAlt_4D5F30() int32 {
	var (
		v0               int32
		v1               int32
		v2               *byte
		v3               *byte
		result           int32
		FileName         [1024]byte
		ExistingFileName [1024]byte
	)
	v0 = 100
	noxflags.SetGame(noxflags.GameFlag23)
	libc.MemSet(memmap.PtrOff(0x973F18, 2408), 0, 1464)
	for {
		v1 = nox_xxx_mapGenStep_4D44E0()
		if v1 == 1 {
			break
		}
		if v1 != 0 {
			if func() int32 {
				p := &v0
				*p--
				return *p
			}() != 0 {
				continue
			}
		}
		return 0
	}
	v2 = nox_fs_root()
	nox_sprintf(&FileName[0], CString("%s\\nc.obj"), v2)
	v3 = nox_fs_root()
	nox_sprintf(&ExistingFileName[0], CString("%s\\blend.obj"), v3)
	result = bool2int(nox_fs_remove(&FileName[0]))
	if result == 0 {
		return result
	}
	if !nox_fs_move(&ExistingFileName[0], &FileName[0]) {
		return 0
	}
	nox_xxx_mapGenMakeInfo_4D5DB0(int32(uintptr(memmap.PtrOff(0x973F18, 2408))))
	noxflags.UnsetGame(noxflags.GameFlag23)
	return 1
}
func sub_4D6000(a1p *nox_object_t) int32 {
	var (
		a1     int32 = int32(uintptr(unsafe.Pointer(a1p)))
		result int32
		v2     int32
	)
	result = 0
	if a1 != 0 {
		v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
		*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 276))) + 4652))) = 0
		*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 276))) + 4656))) = 0
		*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 276))) + 4660))) = 0
		*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 276))) + 4664))) = 0
		*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 276))) + 4668))) = 0
		*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 276))) + 4672))) = 0
		*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 276))) + 4676))) = 0
		*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 276))) + 4680))) = 0
		*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 276))) + 4684))) = 0
		*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 276))) + 4688))) = uint32(nox_game_getQuestStage_4E3CC0())
		result = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 276))))
		*(*uint32)(unsafe.Pointer(uintptr(result + 4692))) = 63
	}
	return result
}
func sub_4D60B0() int32 {
	var (
		result int32
		i      int32
	)
	result = int32(uintptr(unsafe.Pointer(nox_xxx_getFirstPlayerUnit_4DA7C0())))
	for i = result; result != 0; i = result {
		sub_4D6000((*nox_object_t)(unsafe.Pointer(uintptr(i))))
		result = int32(uintptr(unsafe.Pointer(nox_xxx_getNextPlayerUnit_4DA7F0((*nox_object_t)(unsafe.Pointer(uintptr(i)))))))
	}
	return result
}
func sub_4D60E0(a1 int32) *uint32 {
	var (
		result *uint32
		v2     int32
	)
	result = (*uint32)(unsafe.Pointer(uintptr(a1)))
	if (int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 16)))) & 32) != 32 {
		v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
		result = *(**uint32)(unsafe.Pointer(uintptr(v2 + 276)))
		if *(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*1198)) == 1 {
			*(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*1163))++
			result = *(**uint32)(unsafe.Pointer(uintptr(v2 + 276)))
			*(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*1173)) |= 1
		}
	}
	return result
}
func sub_4D6130(a1 int32) int32 {
	var (
		result int32
		v2     int32
	)
	result = a1
	if a1 != 0 {
		if (int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 16)))) & 32) != 32 {
			v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
			*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 276))) + 4660)))++
			result = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 276))))
			*(*uint32)(unsafe.Pointer(uintptr(result + 4692))) |= 2
		}
	}
	return result
}
func sub_4D6170(a1 int32) int32 {
	var (
		result int32
		v2     int32
	)
	result = a1
	if a1 != 0 {
		if (int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 16)))) & 32) != 32 {
			v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
			*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 276))) + 4664)))++
			result = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 276))))
			*(*uint32)(unsafe.Pointer(uintptr(result + 4692))) |= 4
		}
	}
	return result
}
func sub_4D61B0(a1 int32) {
	var v2 int32
	if a1 != 0 {
		if (int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 16)))) & 32) != 32 {
			v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
			*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 276))) + 4668)))++
			var result int32 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 276))))
			*(*uint32)(unsafe.Pointer(uintptr(result + 4692))) |= 8
		}
	}
}
func sub_4D61F0(a1 int32) int32 {
	var (
		result int32
		v2     int32
	)
	result = a1
	if a1 != 0 {
		if (int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 16)))) & 32) != 32 {
			v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
			*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 276))) + 4672)))++
			*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 276))) + 4676)))++
			result = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 276))))
			*(*uint32)(unsafe.Pointer(uintptr(result + 4692))) |= 16
		}
	}
	return result
}
func sub_4D6360(a1 int32) int32 {
	var (
		result int32
		v2     int32
	)
	result = a1
	if a1 != 0 {
		v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
		*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&a1))), unsafe.Sizeof(uint16(0))*0)) = 1520
		*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&a1))), unsafe.Sizeof(uint16(0))*1)) = *(*uint16)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 276))) + 4652)))
		result = nox_xxx_netSendPacket0_4E5420(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 276))) + 2064)))), unsafe.Pointer(&a1), 4, 0, 1)
	}
	return result
}
func sub_4D63B0(a1 int32) int32 {
	var (
		result int32
		v2     int32
	)
	result = a1
	if a1 != 0 {
		v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
		*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&a1))), unsafe.Sizeof(uint16(0))*0)) = 1776
		*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&a1))), unsafe.Sizeof(uint16(0))*1)) = *(*uint16)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 276))) + 4660)))
		result = nox_xxx_netSendPacket0_4E5420(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 276))) + 2064)))), unsafe.Pointer(&a1), 4, 0, 1)
	}
	return result
}
func sub_4D6540(a1 int32) uint32 {
	var (
		v1     int32
		v2     *uint32
		v3     *uint32
		result uint32
		v5     uint32
		i      int32
		v7     *uint32
		v8     uint32
		v9     int32
		v10    float32
		v11    uint32
		v12    uint32
		v13    float32
		v14    uint32
	)
	v1 = nox_xxx_player_4E3CE0()
	v2 = &nox_common_playerInfoFromNum_417090(a1).field_0
	v3 = v2
	if v2 == nil || *(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*1198)) == 0 {
		return 0
	}
	if v1 == 1 {
		result = uint32(sub_4D66E0(*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*1167)), *(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*1168)), *(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*1166)), *(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*1172))))
	} else {
		v5 = 0
		v11 = 0
		v14 = 0
		v12 = 1
		v13 = 1.0
		for i = int32(uintptr(unsafe.Pointer(nox_xxx_getFirstPlayerUnit_4DA7C0()))); i != 0; i = int32(uintptr(unsafe.Pointer(nox_xxx_getNextPlayerUnit_4DA7F0((*nox_object_t)(unsafe.Pointer(uintptr(i))))))) {
			v7 = *(**uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(i + 748))) + 276)))
			if *(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*1198)) != 0 {
				v8 = uint32(sub_4D66E0(*(*uint32)(unsafe.Add(unsafe.Pointer(v7), unsafe.Sizeof(uint32(0))*1167)), *(*uint32)(unsafe.Add(unsafe.Pointer(v7), unsafe.Sizeof(uint32(0))*1168)), 0, *(*uint32)(unsafe.Add(unsafe.Pointer(v7), unsafe.Sizeof(uint32(0))*1172))))
				if v8 > v5 {
					v5 = v8
				}
				v11 += *(*uint32)(unsafe.Add(unsafe.Pointer(v7), unsafe.Sizeof(uint32(0))*1167))
				v12 = *(*uint32)(unsafe.Add(unsafe.Pointer(v7), unsafe.Sizeof(uint32(0))*1172))
				v14 += *(*uint32)(unsafe.Add(unsafe.Pointer(v7), unsafe.Sizeof(uint32(0))*1168))
			}
		}
		v9 = sub_4D66E0(v11, v14, 0, v12)
		if v5 > 0 {
			v13 = float32(float64(uint32(v9)) / float64(v5))
		}
		v10 = float32(float64(uint32(sub_4D66E0(*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*1167)), *(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*1168)), *(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*1166)), *(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*1172))))) * float64(v13))
		result = uint32(int(v10))
	}
	if result > 0x3B9AC9FF {
		result = 0x3B9AC9FF
	}
	return result
}
func sub_4D66E0(a1 uint32, a2 uint32, a3 uint32, a4 uint32) int32 {
	var v5 float32
	v5 = float32(float64(nox_double2float(math.Pow(float64(a4), *(*float64)(memmap.PtrOff(0x581450, 10088))))) * (float64(a1)*10.0 + float64(a2)*35.0 + float64(a3)*0.1))
	return int(v5)
}
func sub_4D6770(a1 int32) int32 {
	var (
		v1 *byte
		v2 int32
		v3 int32
		v4 *byte
		v5 int32
		v7 [90]byte
	)
	v1 = (*byte)(unsafe.Pointer(nox_common_playerInfoFromNum_417090(a1)))
	libc.MemSet(unsafe.Pointer(&v7[0]), 0, 88)
	*(*uint16)(unsafe.Pointer(&v7[88])) = 0
	v7[0] = 240
	v7[1] = 12
	*(*uint16)(unsafe.Pointer(&v7[2])) = uint16(int16(sub_4D7300()))
	v2 = 0
	v3 = int32(uintptr(unsafe.Pointer(nox_xxx_getFirstPlayerUnit_4DA7C0())))
	if v3 != 0 {
		v4 = &v7[8]
		for {
			v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(v3 + 748))))
			if *(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v5 + 276))) + 4792))) == 1 && v2 < 6 {
				*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v4))), -int(unsafe.Sizeof(uint16(0))*1)))) = *(*uint16)(unsafe.Pointer(uintptr(v3 + 36)))
				*(*uint16)(unsafe.Pointer(v4)) = *(*uint16)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v5 + 276))) + 4668)))
				*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v4))), unsafe.Sizeof(uint16(0))*3))) = *(*uint16)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v5 + 276))) + 4664)))
				*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v4))), unsafe.Sizeof(uint16(0))*1))) = *(*uint16)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v5 + 276))) + 4672)))
				*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v4))), unsafe.Sizeof(uint16(0))*2))) = *(*uint16)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v5 + 276))) + 4680)))
				*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v4))), unsafe.Sizeof(uint32(0))*2))) = sub_4D6540(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v5 + 276))) + 2064)))))
				v2++
				v4 = (*byte)(unsafe.Add(unsafe.Pointer(v4), 14))
				*(*uint16)(unsafe.Pointer(&v7[4])) = *((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v1))), unsafe.Sizeof(uint16(0))*2344)))
			}
			v3 = int32(uintptr(unsafe.Pointer(nox_xxx_getNextPlayerUnit_4DA7F0((*nox_object_t)(unsafe.Pointer(uintptr(v3)))))))
			if v3 == 0 {
				break
			}
		}
	}
	return nox_xxx_netSendPacket1_4E5390(a1, int32(uintptr(unsafe.Pointer(&v7[0]))), 90, 0, 1)
}
func sub_4D6880(a1 int32, a2 int32) int32 {
	var v3 [69]byte
	libc.MemSet(unsafe.Pointer(&v3[0]), 0, 68)
	v3[68] = 0
	v3[0] = 240
	v3[1] = 13
	if a2 != 0 {
		v3[4] |= 1
	}
	if sub_51A950() != 0 {
		v3[4] |= 2
	}
	*(*uint16)(unsafe.Pointer(&v3[2])) = uint16(int16(nox_game_getQuestStage_4E3CC0()))
	nox_server_currentMapGetFilename_409B30()
	libc.StrCpy(&v3[5], sub_4D6940())
	nox_server_currentMapGetFilename_409B30()
	libc.StrCpy(&v3[37], sub_4D6950())
	return nox_xxx_netSendPacket1_4E5390(a1, int32(uintptr(unsafe.Pointer(&v3[0]))), 69, 0, 1)
}
func sub_4D6940() *byte {
	return (*byte)(memmap.PtrOff(0x973F18, 3838))
}
func sub_4D6950() *byte {
	return (*byte)(memmap.PtrOff(0x973F18, 3806))
}
func nox_game_sendQuestStage_4D6960(a1 int32) int32 {
	var v2 [69]byte = [69]byte{}
	v2[0] = 240
	v2[1] = 14
	v2[4] = 0
	if sub_51A950() != 0 {
		v2[4] |= 2
	}
	*(*uint16)(unsafe.Pointer(&v2[2])) = uint16(int16(nox_game_getQuestStage_4E3CC0()))
	nox_server_currentMapGetFilename_409B30()
	libc.StrCpy(&v2[5], sub_4D6940())
	nox_server_currentMapGetFilename_409B30()
	libc.StrCpy(&v2[37], sub_4D6950())
	return nox_xxx_netSendPacket1_4E5390(a1, int32(uintptr(unsafe.Pointer(&v2[0]))), 69, 0, 1)
}
func sub_4D6A20(a1 int32, a2 int32) int32 {
	var (
		v3 int16
		v5 int32
	)
	v3 = int16(*(*uint16)(unsafe.Pointer(uintptr(a2 + 40))))
	*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint16(0))*0)) = 4080
	*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint16(0))*1)) = uint16(v3)
	return nox_xxx_netSendPacket0_4E5420(a1, unsafe.Pointer(&v5), 4, 0, 1)
}
func sub_4D6A60() int32 {
	var (
		result int32
		i      int32
	)
	nox_xxx_warriorMaxHealth_587000_312784 = 0x40400000
	nox_xxx_warriorMaxMana_587000_312788 = 0x3F800000
	nox_xxx_warriorMaxStrength_587000_312792 = 0x3F800000
	nox_xxx_warriorMaxSpeed_587000_312796 = 0x3F800000
	nox_xxx_conjurerMaxHealth_587000_312800 = 0x40400000
	nox_xxx_conjurerMaxMana_587000_312804 = 0x40400000
	nox_xxx_conjurerStrength_587000_312808 = 0x3F800000
	nox_xxx_conjurerSpeed_587000_312812 = 0x3F800000
	nox_xxx_wizardMaxHealth_587000_312816 = 0x40400000
	nox_xxx_wizardMaximumMana_587000_312820 = 0x40400000
	nox_xxx_wizardStrength_587000_312824 = 0x3F800000
	nox_xxx_wizardSpeed_587000_312828 = 0x3F800000
	nox_xxx_loadBaseValues_57B200()
	result = int32(uintptr(unsafe.Pointer(nox_xxx_getFirstPlayerUnit_4DA7C0())))
	for i = result; result != 0; i = result {
		nox_xxx_plrReadVals_4EEDC0(i, 0)
		nox_xxx_netStatsMultiplier_4D9C20(i)
		result = int32(uintptr(unsafe.Pointer(nox_xxx_getNextPlayerUnit_4DA7F0((*nox_object_t)(unsafe.Pointer(uintptr(i)))))))
	}
	return result
}
func sub_4D6B10(a1 int32) int32 {
	var (
		result int32
		i      int32
	)
	nox_xxx_warriorMaxHealth_587000_312784 = *memmap.PtrUint32(6112660, 1556076)
	nox_xxx_warriorMaxMana_587000_312788 = *memmap.PtrUint32(6112660, 1556084)
	nox_xxx_warriorMaxStrength_587000_312792 = *memmap.PtrUint32(6112660, 1556064)
	nox_xxx_warriorMaxSpeed_587000_312796 = *memmap.PtrUint32(6112660, 1556072)
	nox_xxx_conjurerMaxHealth_587000_312800 = *memmap.PtrUint32(6112660, 1556060)
	nox_xxx_conjurerMaxMana_587000_312804 = *memmap.PtrUint32(6112660, 1556096)
	nox_xxx_conjurerStrength_587000_312808 = *memmap.PtrUint32(6112660, 1550932)
	nox_xxx_conjurerSpeed_587000_312812 = *memmap.PtrUint32(6112660, 1556080)
	nox_xxx_wizardMaxHealth_587000_312816 = *memmap.PtrUint32(6112660, 1556088)
	nox_xxx_wizardMaximumMana_587000_312820 = *memmap.PtrUint32(6112660, 1556068)
	nox_xxx_wizardStrength_587000_312824 = *memmap.PtrUint32(6112660, 1556100)
	nox_xxx_wizardSpeed_587000_312828 = *memmap.PtrUint32(6112660, 1556092)
	nox_xxx_loadBaseValues_57B200()
	result = int32(uintptr(unsafe.Pointer(nox_xxx_getFirstPlayerUnit_4DA7C0())))
	for i = result; result != 0; i = result {
		nox_xxx_plrReadVals_4EEDC0(i, 0)
		if a1 != 0 {
			nox_xxx_netStatsMultiplier_4D9C20(i)
		}
		result = int32(uintptr(unsafe.Pointer(nox_xxx_getNextPlayerUnit_4DA7F0((*nox_object_t)(unsafe.Pointer(uintptr(i)))))))
	}
	return result
}
func sub_4D6BE0() int32 {
	var result int32
	*memmap.PtrUint32(6112660, 1556076) = nox_xxx_warriorMaxHealth_587000_312784
	*memmap.PtrUint32(6112660, 1556084) = nox_xxx_warriorMaxMana_587000_312788
	*memmap.PtrUint32(6112660, 1556064) = nox_xxx_warriorMaxStrength_587000_312792
	*memmap.PtrUint32(6112660, 1556072) = nox_xxx_warriorMaxSpeed_587000_312796
	*memmap.PtrUint32(6112660, 1556060) = nox_xxx_conjurerMaxHealth_587000_312800
	*memmap.PtrUint32(6112660, 1556096) = nox_xxx_conjurerMaxMana_587000_312804
	*memmap.PtrUint32(6112660, 1550932) = nox_xxx_conjurerStrength_587000_312808
	result = int32(nox_xxx_wizardMaximumMana_587000_312820)
	*memmap.PtrUint32(6112660, 1556080) = nox_xxx_conjurerSpeed_587000_312812
	*memmap.PtrUint32(6112660, 1556088) = nox_xxx_wizardMaxHealth_587000_312816
	*memmap.PtrUint32(6112660, 1556068) = nox_xxx_wizardMaximumMana_587000_312820
	*memmap.PtrUint32(6112660, 1556100) = nox_xxx_wizardStrength_587000_312824
	*memmap.PtrUint32(6112660, 1556092) = nox_xxx_wizardSpeed_587000_312828
	return result
}
func nox_xxx_isQuest_4D6F50() int32 {
	return int32(*memmap.PtrUint32(6112660, 1556160))
}
func nox_xxx_setQuest_4D6F60(a1 int32) int32 {
	var result int32
	result = a1
	*memmap.PtrUint32(6112660, 1556160) = uint32(a1)
	return result
}
func sub_4D6F70() int32 {
	return int32(*memmap.PtrUint32(6112660, 1556164))
}
func sub_4D6F80(a1 int32) int32 {
	var result int32
	result = a1
	*memmap.PtrUint32(6112660, 1556164) = uint32(a1)
	return result
}
func sub_4D6F90(a1 int32) int32 {
	var result int32
	result = a1
	*memmap.PtrUint32(6112660, 1556104) = uint32(a1)
	return result
}
func sub_4D6FA0() int32 {
	return int32(*memmap.PtrUint32(6112660, 1556104))
}
func sub_4D70B0() *byte {
	var result *byte
	result = *(**byte)(memmap.PtrOff(6112660, 1556152))
	if *memmap.PtrUint32(6112660, 1556152) == 0 {
		result = sub_4169F0()
	}
	return result
}
func nox_xxx_bookCreatureTest_4D70C0(a1 int32) int32 {
	return bool2int(noxflags.HasGame(noxflags.GameModeQuest) || a1 != 37 && a1 != 38 && a1 != 40 && a1 != 39)
}
func sub_4D7100(a1 int32) int32 {
	return bool2int(noxflags.HasGame(noxflags.GameModeQuest) || a1 != 111 && a1 != 112 && a1 != 114 && a1 != 113)
}
func sub_4D7150() int32 {
	var (
		result int32
		i      int32
		v2     *int32
		v3     int32
	)
	result = int32(dword_5d4594_1556144)
	if dword_5d4594_1556144 != 0 {
		if nox_frame_xxx_2598000 > uint32(*(*int32)(unsafe.Pointer(&dword_5d4594_1556144))) {
			result = int32(uintptr(unsafe.Pointer(nox_xxx_getFirstPlayerUnit_4DA7C0())))
			for i = result; result != 0; i = result {
				v2 = *(**int32)(unsafe.Pointer(uintptr(i + 748)))
				v3 = *(*int32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(int32(0))*69))
				if (int32(*(*uint8)(unsafe.Pointer(uintptr(v3 + 3680))))&1) == 1 && *(*uint32)(unsafe.Pointer(uintptr(v3 + 4792))) == 1 && *(*int32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(int32(0))*78)) == 0 && *(*int32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(int32(0))*79)) == 0 && (*(*uint32)(unsafe.Pointer(uintptr(v3 + 3680)))&16) == 16 {
					sub_4DF3C0((*nox_playerInfo)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(int32(0))*69))))))
					nox_xxx_playerLeaveObserver_0_4E6AA0((*nox_playerInfo)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(int32(0))*69))))))
					nox_xxx_playerCameraUnlock_4E6040(i)
				}
				result = int32(uintptr(unsafe.Pointer(nox_xxx_getNextPlayerUnit_4DA7F0((*nox_object_t)(unsafe.Pointer(uintptr(i)))))))
			}
		}
	}
	return result
}
func sub_4D71E0(a1 int32) int32 {
	var result int32
	result = a1
	dword_5d4594_1556136 = uint32(a1)
	return result
}
func sub_4D71F0() uint32 {
	var (
		result uint32
		v1     int32
		v2     int8
	)
	result = dword_5d4594_1556136
	if dword_5d4594_1556136 != 0 {
		if (nox_frame_xxx_2598000 - dword_5d4594_1556136) >= 9000 {
			v1 = 0
			result = uint32(uintptr(unsafe.Pointer(nox_xxx_getFirstPlayerUnit_4DA7C0())))
			if result != 0 {
				for {
					if *(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(result + 748))) + 308))) != 0 {
						v1 = 1
					}
					result = uint32(uintptr(unsafe.Pointer(nox_xxx_getNextPlayerUnit_4DA7F0((*nox_object_t)(unsafe.Pointer(uintptr(result)))))))
					if result == 0 {
						break
					}
				}
				if v1 != 0 {
					result = uint32(nox_xxx_player_4E3CE0())
					if result > 1 {
						sub_4D71E0(0)
						result = uint32(sub_4D72C0())
						if result == 0 {
							sub_4D72B0(1)
							v2 = int8(sub_4D72C0())
							result = uint32(sub_4D7280(math.MaxUint8, v2))
						}
					}
				}
			}
		}
	}
	return result
}
func sub_4D7280(a1 int32, a2 int8) int32 {
	var v4 [3]byte
	v4[0] = 240
	v4[1] = 24
	v4[2] = byte(a2)
	return nox_xxx_netSendPacket1_4E5390(a1, int32(uintptr(unsafe.Pointer(&v4[0]))), 3, 0, 1)
}
func sub_4D72B0(a1 int32) int32 {
	var result int32
	result = a1
	*memmap.PtrUint32(6112660, 1556140) = uint32(a1)
	return result
}
func sub_4D72C0() int32 {
	return int32(*memmap.PtrUint32(6112660, 1556140))
}
func sub_4D72D0(a1 int32) int32 {
	var result int32
	result = int32(dword_5d4594_1556128)
	*memmap.PtrUint32(6112660, 1556132) = dword_5d4594_1556128
	dword_5d4594_1556128 = uint32(a1)
	return result
}
func sub_4D72F0() int32 {
	return int32(dword_5d4594_1556128)
}
func sub_4D7300() int32 {
	return int32(*memmap.PtrUint32(6112660, 1556132))
}
func sub_4D7310() int32 {
	var (
		v0 int32
		v1 int32
		v2 int32
		v3 int32
		i  int32
	)
	v0 = 0
	v1 = 99999
	v2 = int32(uintptr(unsafe.Pointer(nox_xxx_getFirstPlayerUnit_4DA7C0())))
	if v2 == 0 {
		return 0
	}
	for {
		if *(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 748))) + 276))) + 4792))) == 1 {
			v3 = 0
			for i = v2.FirstItem(); i != 0; i = nox_xxx_inventoryGetNext_4E7990(i) {
				if int32(*(*uint8)(unsafe.Pointer(uintptr(i + 8))))&64 != 0 {
					v3++
				}
			}
			if v3 <= v1 {
				v1 = v3
				v0 = v2
			}
		}
		v2 = int32(uintptr(unsafe.Pointer(nox_xxx_getNextPlayerUnit_4DA7F0((*nox_object_t)(unsafe.Pointer(uintptr(v2)))))))
		if v2 == 0 {
			break
		}
	}
	return v0
}
func sub_4D7390(a1 int32) {
	var (
		v1 int32
		v2 int32
		v3 int32
	)
	if a1 != 0 {
		if nox_xxx_player_4E3CE0() != 0 {
			v1 = a1.FirstItem()
			if v1 != 0 {
				for {
					v2 = nox_xxx_inventoryGetNext_4E7990(v1)
					if int32(*(*uint8)(unsafe.Pointer(uintptr(v1 + 8))))&64 != 0 {
						v3 = sub_4D7310()
						if v3 != 0 {
							sub_4ED0C0(a1, (*nox_object_t)(unsafe.Pointer(uintptr(v1))))
							nox_xxx_inventoryPutImpl_4F3070(v3, (*nox_object_t)(unsafe.Pointer(uintptr(v1))), 1)
							nox_xxx_netPriMsgToPlayer_4DA2C0((*nox_object_t)(unsafe.Pointer(uintptr(v3))), CString("GeneralPrint:GainedKey"), 0)
							nox_xxx_aud_501960(820, (*nox_object_t)(unsafe.Pointer(uintptr(v3))), 0, 0)
						}
					}
					v1 = v2
					if v2 == 0 {
						break
					}
				}
			}
			sub_4EDD00(a1, 64)
		} else {
			sub_4EDD00(a1, 64)
		}
	}
}
func sub_4D7430() int32 {
	return int32(*memmap.PtrUint32(6112660, 1556116))
}
func sub_4D7440(a1 int32) int32 {
	var result int32
	result = a1
	*memmap.PtrUint32(6112660, 1556116) = uint32(a1)
	return result
}
func sub_4D7450(a1 int32, a2 int16) int32 {
	var v3 [4]byte
	v3[0] = 240
	v3[1] = 29
	*(*uint16)(unsafe.Pointer(&v3[2])) = uint16(a2)
	return nox_xxx_netSendPacket0_4E5420(a1, unsafe.Pointer(&v3[0]), 4, 0, 1)
}
func sub_4D7480(a1 int32) {
	var (
		v1 int32
		v2 int32
		v3 *float2
	)
	if a1 != 0 && int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))&4 != 0 {
		v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
		v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 316))))
		if v2 != 0 {
			v3 = *(**float2)(unsafe.Pointer(uintptr(v2 + 700)))
			nox_xxx_playerLeaveObserver_0_4E6AA0((*nox_playerInfo)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v1 + 276)))))))
			nox_xxx_playerCameraUnlock_4E6040(a1)
			*(*uint32)(unsafe.Pointer(uintptr(v1 + 316))) = 0
			nox_xxx_unitMove_4E7010((*nox_object_t)(unsafe.Pointer(uintptr(a1))), (*float2)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(float2{})*10)))
			nox_xxx_aud_501960(312, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 2, int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 36)))))
			nox_xxx_netSendPointFx_522FF0(-127, (*float2)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(float2{})*10)))
		}
	}
}
func sub_4D7520(a1 int32) int8 {
	var (
		v1 int32
		i  int32
		v3 int32
		v4 int32
		v5 int32
	)
	*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v1))), 0)) = *memmap.PtrUint8(6112660, 1556120)
	if *memmap.PtrUint32(6112660, 1556120) == 1 {
		if a1 != 0 {
			goto LABEL_19
		}
		for i = int32(uintptr(unsafe.Pointer(nox_xxx_getFirstPlayerUnit_4DA7C0()))); i != 0; i = int32(uintptr(unsafe.Pointer(nox_xxx_getNextPlayerUnit_4DA7F0((*nox_object_t)(unsafe.Pointer(uintptr(i))))))) {
			v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(i + 748))))
			if *(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v3 + 276))) + 4792))) != 0 && *(*uint32)(unsafe.Pointer(uintptr(v3 + 316))) != 0 {
				sub_4D7480(i)
			}
		}
		v1 = int32(uintptr(unsafe.Pointer(noxServer.firstServerObject())))
		v4 = v1
		if v1 == 0 {
		LABEL_19:
			*memmap.PtrUint32(6112660, 1556120) = uint32(a1)
		} else {
			for {
				v5 = int32(uintptr(unsafe.Pointer((*nox_object_t)(unsafe.Pointer(uintptr(v4))).Next())))
				*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v1))), 0)) = *(*uint8)(unsafe.Pointer(uintptr(v4 + 8)))
				if v1&32 != 0 && int32(*(*uint8)(unsafe.Pointer(uintptr(v4 + 12))))&2 != 0 {
					*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v1))), 0)) = uint8(int8(nox_xxx_objectSetOff_4E7600((*nox_object_t)(unsafe.Pointer(uintptr(v4))))))
				}
				v4 = v5
				if v5 == 0 {
					break
				}
			}
			*memmap.PtrUint32(6112660, 1556120) = 0
		}
	} else {
		*memmap.PtrUint32(6112660, 1556120) = uint32(a1)
	}
	return int8(v1)
}
func sub_4D75E0() int32 {
	return int32(*memmap.PtrUint32(6112660, 1556120))
}
func sub_4D75F0(a1 int32) int32 {
	var result int32
	result = a1
	*memmap.PtrUint32(6112660, 1556108) = uint32(a1)
	return result
}
func nox_server_checkWarpGate_4D7600() {
	var exp int32 = nox_xxx_player_4E3CE0()
	if exp == 0 {
		return
	}
	if (nox_frame_xxx_2598000 - *memmap.PtrUint32(6112660, 1556108)) < 30 {
		return
	}
	var inGate int32 = 0
	for unit := unsafe.Pointer(unsafe.Pointer(nox_xxx_getFirstPlayerUnit_4DA7C0())); unit != nil; unit = unsafe.Pointer(nox_xxx_getNextPlayerUnit_4DA7F0((*nox_object_t)(unit))) {
		var v3 int32 = int32(*(*uint32)(unsafe.Pointer(uintptr(int32(uintptr(unit)) + 748))))
		if *(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v3 + 276))) + 4792))) != 0 && *(*uint32)(unsafe.Pointer(uintptr(v3 + 316))) != 0 {
			inGate++
		}
	}
	if exp != inGate {
		return
	}
	if !nox_server_questMaybeWarp_4E8F60() {
		for unit := unsafe.Pointer(unsafe.Pointer(nox_xxx_getFirstPlayerUnit_4DA7C0())); unit != nil; unit = unsafe.Pointer(nox_xxx_getNextPlayerUnit_4DA7F0((*nox_object_t)(unit))) {
			var v5 int32 = int32(*(*uint32)(unsafe.Pointer(uintptr(int32(uintptr(unit)) + 748))))
			if *(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v5 + 276))) + 4792))) != 0 && *(*uint32)(unsafe.Pointer(uintptr(v5 + 316))) != 0 {
				sub_4D7480(int32(uintptr(unit)))
				if exp <= 1 {
					nox_xxx_netPriMsgToPlayer_4DA2C0((*nox_object_t)(unit), CString("Gauntlet.c:WarpRestrictedSolo"), 0)
				} else {
					nox_xxx_netPriMsgToPlayer_4DA2C0((*nox_object_t)(unit), CString("Gauntlet.c:WarpRestrictedMulti"), 0)
				}
			}
		}
	}
}
func sub_4D76E0(a1 int32) int32 {
	var result int32
	result = a1
	*memmap.PtrUint32(6112660, 1556124) = uint32(a1)
	return result
}
func sub_4D76F0() int32 {
	return int32(*memmap.PtrUint32(6112660, 1556124))
}
func nox_client_checkQuestExp_SKU2_4D7700() bool {
	return true
}
func nox_xxx_checkAccessToGLava_4D7790() int32 {
	var (
		v0       *byte
		FileName [1024]byte
	)
	v0 = nox_fs_root()
	nox_sprintf(&FileName[0], CString("%s\\Maps\\G_Lava"), v0)
	return bool2int(nox_fs_access(&FileName[0], 0) == 0)
}
func sub_4D77D0(a1 int32) int32 {
	var (
		result    int32
		Data      [4]uint8
		phkResult HKEY
		SubKey    [128]byte
	)
	result = a1 - 4096
	switch a1 {
	case 4096:
		*(*uint32)(unsafe.Pointer(&Data[0])) = 9472
		goto LABEL_8
	case 4098:
		*(*uint32)(unsafe.Pointer(&Data[0])) = 9474
		goto LABEL_8
	case 4099:
		*(*uint32)(unsafe.Pointer(&Data[0])) = 9475
		goto LABEL_8
	case 4101:
		*(*uint32)(unsafe.Pointer(&Data[0])) = 9477
		goto LABEL_8
	case 4102:
		*(*uint32)(unsafe.Pointer(&Data[0])) = 9478
		goto LABEL_8
	case 4105:
		*(*uint32)(unsafe.Pointer(&Data[0])) = 9481
	LABEL_8:
		libc.StrCpy(&SubKey[0], CString("SOFTWARE\\Westwood\\Nox"))
		result = compatRegOpenKeyExA((HKEY)(unsafe.Pointer(uintptr(1))), &SubKey[0], 0, 0xF003F, &phkResult)
		if result == 0 {
			compatRegSetValueExA(phkResult, CString("SKU"), 0, 4, &Data[0], 4)
			result = compatRegCloseKey(phkResult)
		}
	default:
		return result
	}
	return result
}
func nox_common_readSKU_fromRegistry_4D78C0() int32 {
	var (
		result    int32
		Data      [4]uint8
		phkResult HKEY
		cbData    uint32
		Type      uint32
		SubKey    [128]byte
	)
	libc.StrCpy(&SubKey[0], CString("SOFTWARE\\Westwood\\Nox"))
	*(*uint32)(unsafe.Pointer(&Data[0])) = math.MaxUint32
	cbData = 4
	result = compatRegOpenKeyExA((HKEY)(unsafe.Pointer(uintptr(1))), &SubKey[0], 0, 0xF003F, &phkResult)
	if result == 0 {
		compatRegQueryValueExA(phkResult, CString("SKU"), nil, &Type, &Data[0], &cbData)
		result = compatRegCloseKey(phkResult)
	}
	if *(*int32)(unsafe.Pointer(&Data[0])) < 9472 {
		result = nox_xxx_checkAccessToGLava_4D7790()
		if result == 1 {
			result = sub_4D77D0(*(*int32)(unsafe.Pointer(&Data[0])))
		}
	}
	return result
}
func nox_xxx_player_4D7960(a1 int8) int32 {
	var result int32
	result = 1 << int32(a1)
	*memmap.PtrUint32(6112660, 1556300) |= uint32(1 << int32(a1))
	return result
}
func nox_xxx_player_4D7980(a1 int8) int32 {
	return bool2int((*memmap.PtrUint32(6112660, 1556300) & uint32(1<<int32(a1))) != 0)
}
func sub_4D79A0(a1 int8) int32 {
	var result int32
	result = ^(1 << int32(a1))
	*memmap.PtrUint32(6112660, 1556300) &= uint32(result)
	return result
}
func sub_4D79C0(a1 int32) int32 {
	var (
		v1     int32
		result int32
		v3     int32
	)
	v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
	sub_4D9D20(math.MaxUint8, (*nox_object_t)(unsafe.Pointer(uintptr(a1))))
	sub_4D6000((*nox_object_t)(unsafe.Pointer(uintptr(a1))))
	for result = int32(uintptr(unsafe.Pointer(nox_xxx_getFirstPlayerUnit_4DA7C0()))); result != 0; result = int32(uintptr(unsafe.Pointer(nox_xxx_getNextPlayerUnit_4DA7F0((*nox_object_t)(unsafe.Pointer(uintptr(result))))))) {
		v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(result + 748))))
		*(*uint8)(unsafe.Pointer(uintptr(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v1 + 276))) + 2064)))) + v3 + 452))) = 0
		*(*uint32)(unsafe.Pointer(uintptr(v3 + int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v1 + 276))) + 2064))))*4 + 324))) = 0
		*(*uint8)(unsafe.Pointer(uintptr(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v1 + 276))) + 2064)))) + v3 + 484))) = 0
		*(*uint8)(unsafe.Pointer(uintptr(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v1 + 276))) + 2064)))) + v3 + 516))) = 0
	}
	return result
}
func sub_4D7A60(a1 int32) int32 {
	var result int32
	result = a1
	*memmap.PtrUint32(6112660, uint32(a1*4)+0x17BECC) = nox_frame_xxx_2598000
	return result
}
func sub_4D7A80() int32 {
	var (
		v0     *uint8
		v1     int32
		v2     int32
		v3     *byte
		v4     int32
		v5     int32
		v6     int32
		result int32
	)
	v0 = (*uint8)(memmap.PtrOff(6112660, 1556172))
	v1 = int32(324 - uint32(uintptr(memmap.PtrOff(6112660, 1556172))))
	v2 = 484
	for {
		v3 = (*byte)(unsafe.Pointer(nox_common_playerInfoFromNum_417090(v2 - 484)))
		if v3 != nil && *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v3))), unsafe.Sizeof(uint32(0))*523))) != 0 && *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v3))), unsafe.Sizeof(uint32(0))*514))) != 0 && *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v3))), unsafe.Sizeof(uint32(0))*1198))) == 1 {
			goto LABEL_12
		}
		if *(*uint32)(unsafe.Pointer(v0)) != 0 && nox_frame_xxx_2598000-*(*uint32)(unsafe.Pointer(v0)) > (nox_gameFPS*30) {
			v4 = int32(uintptr(unsafe.Pointer(nox_xxx_getFirstPlayerUnit_4DA7C0())))
			if v4 != 0 {
				v5 = int32(uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v0), v1)))))
				for {
					v6 = int32(*(*uint32)(unsafe.Pointer(uintptr(v4 + 748))))
					*(*uint8)(unsafe.Pointer(uintptr(v2 + v6 - 32))) = 0
					*(*uint32)(unsafe.Pointer(uintptr(v5 + v6))) = 0
					*(*uint8)(unsafe.Pointer(uintptr(v2 + v6))) = 0
					*(*uint8)(unsafe.Pointer(uintptr(v2 + v6 + 32))) = 0
					v4 = int32(uintptr(unsafe.Pointer(nox_xxx_getNextPlayerUnit_4DA7F0((*nox_object_t)(unsafe.Pointer(uintptr(v4)))))))
					if v4 == 0 {
						break
					}
				}
				v1 = int32(324 - uint32(uintptr(memmap.PtrOff(6112660, 1556172))))
			}
		LABEL_12:
			*(*uint32)(unsafe.Pointer(v0)) = 0
		}
		v2++
		v0 = (*uint8)(unsafe.Add(unsafe.Pointer(v0), 4))
		result = v2 - 484
		if v2-484 >= 32 {
			break
		}
	}
	return result
}
func sub_4D7B40() int32 {
	var result int32
	result = 0
	libc.MemSet(memmap.PtrOff(6112660, 1556172), 0, 128)
	return result
}
func sub_4D7B60(a1 int32) int32 {
	var (
		result int32
		i      int32
		v3     [7]byte
	)
	v3[0] = 210
	*(*uint16)(unsafe.Pointer(&v3[1])) = uint16(nox_xxx_netGetUnitCodeServ_578AC0((*nox_object_t)(unsafe.Pointer(uintptr(a1)))))
	*(*uint16)(unsafe.Pointer(&v3[3])) = *(*uint16)(unsafe.Pointer(uintptr(a1 + 4)))
	v3[5] = 1
	v3[6] = 2
	result = int32(uintptr(unsafe.Pointer(nox_xxx_getFirstPlayerUnit_4DA7C0())))
	for i = result; result != 0; i = result {
		nox_xxx_netSendPacket0_4E5420(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(i + 748))) + 276))) + 2064)))), unsafe.Pointer(&v3[0]), 7, 0, 1)
		result = int32(uintptr(unsafe.Pointer(nox_xxx_getNextPlayerUnit_4DA7F0((*nox_object_t)(unsafe.Pointer(uintptr(i)))))))
	}
	return result
}
func nox_xxx_netSendInterestingId_4D7BE0(a1 int32) int32 {
	var (
		result int32
		i      int32
		v3     [7]byte
	)
	v3[0] = 210
	*(*uint16)(unsafe.Pointer(&v3[1])) = uint16(nox_xxx_netGetUnitCodeServ_578AC0((*nox_object_t)(unsafe.Pointer(uintptr(a1)))))
	*(*uint16)(unsafe.Pointer(&v3[3])) = *(*uint16)(unsafe.Pointer(uintptr(a1 + 4)))
	v3[5] = 2
	v3[6] = 2
	result = int32(uintptr(unsafe.Pointer(nox_xxx_getFirstPlayerUnit_4DA7C0())))
	for i = result; result != 0; i = result {
		nox_xxx_netSendPacket0_4E5420(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(i + 748))) + 276))) + 2064)))), unsafe.Pointer(&v3[0]), 7, 0, 1)
		result = int32(uintptr(unsafe.Pointer(nox_xxx_getNextPlayerUnit_4DA7F0((*nox_object_t)(unsafe.Pointer(uintptr(i)))))))
	}
	return result
}
func sub_4D7C60() int32 {
	var (
		result int32
		v1     float32
		v2     float32
	)
	dword_5d4594_1556316 = 0
	*mem_getFloatPtr(6112660, 0x17BF54) = float32(nox_xxx_gamedataGetFloat_419D40(CString("CamperRadiusSq")))
	v1 = float32(nox_xxx_gamedataGetFloat_419D40(CString("CamperStartTime")))
	*memmap.PtrUint32(6112660, 1556312) = uint32(int(v1))
	v2 = float32(nox_xxx_gamedataGetFloat_419D40(CString("CamperFadeTime")))
	result = int(v2)
	*memmap.PtrUint32(6112660, 1556304) = uint32(result)
	return result
}
func sub_4D7CC0() int32 {
	var (
		result int32
		i      int32
		v2     int32
		v3     float64
		v4     float64
		j      int32
		v6     *uint32
		v7     int32
	)
	platformTicks()
	result = sub_409F40(8192)
	if result != 0 {
		if (nox_frame_xxx_2598000 - dword_5d4594_1556316) > uint32(*memmap.PtrInt32(6112660, 1556312)) {
			for i = int32(uintptr(unsafe.Pointer(nox_xxx_getFirstPlayerUnit_4DA7C0()))); i != 0; i = int32(uintptr(unsafe.Pointer(nox_xxx_getNextPlayerUnit_4DA7F0((*nox_object_t)(unsafe.Pointer(uintptr(i))))))) {
				if (*(*uint32)(unsafe.Pointer(uintptr(i + 16))) & 32800) == 0 {
					v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(i + 748))))
					if (int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 276))) + 3680)))) & 1) == 0 {
						v3 = float64(*(*float32)(unsafe.Pointer(uintptr(i + 60))) - *(*float32)(unsafe.Pointer(uintptr(v2 + 252))))
						v4 = float64(*(*float32)(unsafe.Pointer(uintptr(i + 56))) - *(*float32)(unsafe.Pointer(uintptr(v2 + 248))))
						if v4*v4+v3*v3 < float64(*mem_getFloatPtr(6112660, 0x17BF54)) {
							*(*uint32)(unsafe.Pointer(uintptr(v2 + 256))) = nox_frame_xxx_2598000
						}
						*(*uint32)(unsafe.Pointer(uintptr(v2 + 248))) = *(*uint32)(unsafe.Pointer(uintptr(i + 56)))
						*(*uint32)(unsafe.Pointer(uintptr(v2 + 252))) = *(*uint32)(unsafe.Pointer(uintptr(i + 60)))
					}
				}
			}
			dword_5d4594_1556316 = nox_frame_xxx_2598000
		}
		result = int32(uintptr(unsafe.Pointer(nox_xxx_getFirstPlayerUnit_4DA7C0())))
		for j = result; result != 0; j = result {
			if (*(*uint32)(unsafe.Pointer(uintptr(j + 16))) & 32800) == 0 {
				v6 = *(**uint32)(unsafe.Pointer(uintptr(j + 748)))
				if (int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*69)) + 3680)))) & 1) == 0 {
					v7 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*65)))
					if (dword_5d4594_1556316 - *(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*64))) > uint32(*memmap.PtrInt32(6112660, 1556304)) {
						if v7 != 0 {
							nox_xxx_netSendInterestingId_4D7BE0(j)
							*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*65)) = 0
						}
					} else if v7 == 0 {
						if !noxflags.HasGame(noxflags.GameModeCTF) || int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*69)) + 4))))&1 != 0 {
							nox_xxx_aud_501960(774, (*nox_object_t)(unsafe.Pointer(uintptr(j))), 0, 0)
						}
						sub_4D7B60(j)
						*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*65)) = 1
					}
				}
			}
			result = int32(uintptr(unsafe.Pointer(nox_xxx_getNextPlayerUnit_4DA7F0((*nox_object_t)(unsafe.Pointer(uintptr(j)))))))
		}
	}
	return result
}
func sub_4D7E50(a1 int32) int32 {
	var (
		result int32
		v2     *uint32
	)
	result = a1
	if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))&4 != 0 {
		v2 = *(**uint32)(unsafe.Pointer(uintptr(a1 + 748)))
		*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*62)) = 0
		*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*63)) = 0
		*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*64)) = nox_frame_xxx_2598000
		if *(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*65)) != 0 {
			result = nox_xxx_netSendInterestingId_4D7BE0(a1)
		}
		*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*65)) = 0
	}
	return result
}
func sub_4D7EA0() int32 {
	var (
		result int32
		i      int32
	)
	result = int32(uintptr(unsafe.Pointer(nox_xxx_getFirstPlayerUnit_4DA7C0())))
	for i = result; result != 0; i = result {
		sub_4D7E50(i)
		result = int32(uintptr(unsafe.Pointer(nox_xxx_getNextPlayerUnit_4DA7F0((*nox_object_t)(unsafe.Pointer(uintptr(i)))))))
	}
	return result
}
func nox_xxx_netCreatureCmd_4D7EE0(player int32, orderType int8) int32 {
	var buf [2]byte
	buf[0] = 237
	buf[1] = byte(orderType)
	return nox_xxx_netSendPacket1_4E5390(player, int32(uintptr(unsafe.Pointer(&buf[0]))), 2, 0, 1)
}
func nox_xxx_netNotifyRate_4D7F10(a1 int32) int32 {
	var v2 [2]byte
	v2[0] = 236
	v2[1] = byte(int8(nox_xxx_rateGet_40A6C0()))
	return nox_xxx_netSendPacket1_4E5390(a1, int32(uintptr(unsafe.Pointer(&v2[0]))), 2, 0, 1)
}
func nox_xxx_netReportObjectPoison_4D7F40(a1 int32, a2 *uint32, a3 int8) int32 {
	var v3 int32
	v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
	*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a1))), 0)) = 218
	*(*uint16)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(&a1))), 1)))) = uint16(nox_xxx_netGetUnitCodeServ_578AC0((*nox_object_t)(unsafe.Pointer(a2))))
	*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a1))), unsafe.Sizeof(int32(0))-1)) = uint8(a3)
	return nox_xxx_netSendPacket0_4E5420(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v3 + 276))) + 2064)))), unsafe.Pointer(&a1), 4, 0, 1)
}
func sub_4D80C0(a1 int32, a2 int8) int32 {
	var (
		result int32
		v3     int32
	)
	result = a1
	if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))&4 != 0 {
		v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a1))), 0)) = 235
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a1))), 1)) = uint8(a2)
		result = nox_xxx_netSendPacket0_4E5420(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v3 + 276))) + 2064)))), unsafe.Pointer(&a1), 2, 0, 1)
	}
	return result
}
func nox_xxx_netAbilRepotState_4D8100(a1 int32, a2 int8, a3 int8) int32 {
	var (
		result int32
		v4     int32
	)
	result = a1
	if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))&4 != 0 {
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a1))), 2)) = uint8(a3)
		v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(result + 748))))
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a1))), 1)) = uint8(a2)
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a1))), 0)) = 206
		result = nox_xxx_netSendPacket0_4E5420(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v4 + 276))) + 2064)))), unsafe.Pointer(&a1), 3, 0, 1)
	}
	return result
}
func nox_xxx_netReportActiveAbils_4D8150(a1 int32, a2 int8, a3 int8) int32 {
	var (
		result int32
		v4     int32
	)
	result = a1
	if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))&4 != 0 {
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a1))), 2)) = uint8(a3)
		v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(result + 748))))
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a1))), 1)) = uint8(a2)
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a1))), 0)) = 207
		result = nox_xxx_netSendPacket0_4E5420(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v4 + 276))) + 2064)))), unsafe.Pointer(&a1), 3, 0, 1)
	}
	return result
}
func sub_4D81A0(a1 int32) {
	var (
		v1 float64
		v2 int32
		v3 [5]byte
	)
	if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))&4 != 0 {
		v1 = float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 28))))
		v3[0] = 110
		v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
		*(*uint32)(unsafe.Pointer(&v3[1])) = uint32(int32(int64(v1)))
		nox_xxx_netSendPacket0_4E5420(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 276))) + 2064)))), unsafe.Pointer(&v3[0]), 5, 0, 1)
	}
}
func nox_xxx_netReportAnimFrame_4D81F0(a1 int32, a2 *uint32) int32 {
	var v3 [7]byte
	v3[0] = 107
	*(*uint16)(unsafe.Pointer(&v3[1])) = uint16(nox_xxx_netGetUnitCodeServ_578AC0((*nox_object_t)(unsafe.Pointer(a2))))
	*(*uint32)(unsafe.Pointer(&v3[3])) = *(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*33))
	return nox_xxx_netSendPacket0_4E5420(a1, unsafe.Pointer(&v3[0]), 7, 0, 1)
}
func nox_xxx_netReportXStatus_4D8230(a1 int32, a2 *uint32) int32 {
	var v3 [7]byte
	v3[0] = 101
	*(*uint16)(unsafe.Pointer(&v3[1])) = uint16(nox_xxx_netGetUnitCodeServ_578AC0((*nox_object_t)(unsafe.Pointer(a2))))
	*(*uint32)(unsafe.Pointer(&v3[3])) = *(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*5))
	return nox_xxx_netSendPacket0_4E5420(a1, unsafe.Pointer(&v3[0]), 7, 0, 1)
}
func nox_xxx_netReportPlrStatus_4D8270(a1 int32) int32 {
	var (
		v1 int32
		v2 int32
		v4 [5]byte
	)
	v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 16))))
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
	v4[0] = 102
	*(*uint32)(unsafe.Pointer(&v4[1])) = uint32(v1)
	return nox_xxx_netSendPacket0_4E5420(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 276))) + 2064)))), unsafe.Pointer(&v4[0]), 5, 0, 1)
}
func nox_xxx_netReportCharges_4D82B0(a1 int32, item *nox_object_t, a3 int8, a4 int8) int32 {
	var v5 [5]byte
	v5[0] = 100
	*(*uint16)(unsafe.Pointer(&v5[1])) = uint16(nox_xxx_netGetUnitCodeServ_578AC0(item))
	v5[3] = byte(a3)
	v5[4] = byte(a4)
	return nox_xxx_netSendPacket1_4E5390(a1, int32(uintptr(unsafe.Pointer(&v5[0]))), 5, 0, 0)
}
func sub_4D82F0(a1 int32, a2 *uint32) *uint32 {
	var (
		result *uint32
		v3     int32
		v4     int32
		v5     int32
		v6     int32
		v7     *uint32
		v8     int32
		v9     int32
		v10    int32
		v11    int32
		v12    int32
		v13    int32
		v14    *uint32
		v15    int32
		v16    int32
		v17    int32
		v18    int32
		v19    [11]byte
	)
	result = a2
	v3 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*2)))
	v4 = 0
	if uint32(v3)&0x11001000 != 0 {
		v5 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*173)))
		v6 = 4
		v7 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*173)))))
		for {
			if *v7 != 0 {
				v4 = 1
			}
			v7 = (*uint32)(unsafe.Add(unsafe.Pointer(v7), unsafe.Sizeof(uint32(0))*1))
			v6--
			if v6 == 0 {
				break
			}
		}
		if v4 != 0 {
			v8 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*123)))
			v19[0] = 81
			*(*uint16)(unsafe.Pointer(&v19[1])) = *(*uint16)(unsafe.Pointer(uintptr(v8 + 36)))
			if int32(*(*uint8)(unsafe.Pointer(uintptr(v8 + 8))))&4 != 0 {
				v19[2] |= 128
			}
			*(*uint32)(unsafe.Pointer(&v19[3])) = uint32(nox_xxx_weaponInventoryEquipFlags_415820((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a2))))))))
			v9 = 0
			v10 = v5
			for {
				if *(*uint32)(unsafe.Pointer(uintptr(v10))) != 0 {
					v19[v9+7] = byte(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v10))) + 4))))
				} else {
					v19[v9+7] = math.MaxUint8
				}
				v9++
				v10 += 4
				if v9 >= 4 {
					break
				}
			}
			result = (*uint32)(unsafe.Pointer(uintptr(nox_xxx_netSendPacket1_4E5390(a1, int32(uintptr(unsafe.Pointer(&v19[0]))), 11, 0, 0))))
		} else {
			v11 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*123)))
			v19[0] = 80
			*(*uint16)(unsafe.Pointer(&v19[1])) = *(*uint16)(unsafe.Pointer(uintptr(v11 + 36)))
			if int32(*(*uint8)(unsafe.Pointer(uintptr(v11 + 8))))&4 != 0 {
				v19[2] |= 128
			}
			*(*uint32)(unsafe.Pointer(&v19[3])) = uint32(nox_xxx_weaponInventoryEquipFlags_415820((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a2))))))))
			result = (*uint32)(unsafe.Pointer(uintptr(nox_xxx_netSendPacket1_4E5390(a1, int32(uintptr(unsafe.Pointer(&v19[0]))), 7, 0, 0))))
		}
	} else if uint32(v3)&0x2000000 != 0 {
		v12 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*173)))
		v13 = 4
		v14 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*173)))))
		for {
			if *v14 != 0 {
				v4 = 1
			}
			v14 = (*uint32)(unsafe.Add(unsafe.Pointer(v14), unsafe.Sizeof(uint32(0))*1))
			v13--
			if v13 == 0 {
				break
			}
		}
		if v4 != 0 {
			v15 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*123)))
			v19[0] = 82
			*(*uint16)(unsafe.Pointer(&v19[1])) = *(*uint16)(unsafe.Pointer(uintptr(v15 + 36)))
			if int32(*(*uint8)(unsafe.Pointer(uintptr(v15 + 8))))&4 != 0 {
				v19[2] |= 128
			}
			*(*uint32)(unsafe.Pointer(&v19[3])) = uint32(nox_xxx_unitArmorInventoryEquipFlags_415C70((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a2))))))))
			v16 = 0
			v17 = v12
			for {
				if *(*uint32)(unsafe.Pointer(uintptr(v17))) != 0 {
					v19[v16+7] = byte(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v17))) + 4))))
				} else {
					v19[v16+7] = math.MaxUint8
				}
				v16++
				v17 += 4
				if v16 >= 4 {
					break
				}
			}
			result = (*uint32)(unsafe.Pointer(uintptr(nox_xxx_netSendPacket1_4E5390(a1, int32(uintptr(unsafe.Pointer(&v19[0]))), 11, 0, 0))))
		} else {
			v18 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*123)))
			v19[0] = 79
			*(*uint16)(unsafe.Pointer(&v19[1])) = *(*uint16)(unsafe.Pointer(uintptr(v18 + 36)))
			if int32(*(*uint8)(unsafe.Pointer(uintptr(v18 + 8))))&4 != 0 {
				v19[2] |= 128
			}
			*(*uint32)(unsafe.Pointer(&v19[3])) = uint32(nox_xxx_unitArmorInventoryEquipFlags_415C70((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a2))))))))
			result = (*uint32)(unsafe.Pointer(uintptr(nox_xxx_netSendPacket1_4E5390(a1, int32(uintptr(unsafe.Pointer(&v19[0]))), 7, 0, 0))))
		}
	}
	return result
}
func nox_xxx_netReportDequip_4D84C0(a1 int32, object *nox_object_t) int32 {
	var (
		result int32
		v3     int32
		v4     int32
		v5     int32
		v6     int32
		v7     [7]byte
	)
	result = int32(uintptr(unsafe.Pointer(object)))
	v3 = int32(object.obj_class)
	if uint32(v3)&0x11001000 != 0 {
		v4 = int32(object.field_123)
		v7[0] = 84
		*(*uint16)(unsafe.Pointer(&v7[1])) = *(*uint16)(unsafe.Pointer(uintptr(v4 + 36)))
		v5 = nox_xxx_weaponInventoryEquipFlags_415820(object)
	} else {
		if (uint32(v3) & 0x2000000) == 0 {
			return result
		}
		v6 = int32(object.field_123)
		v7[0] = 83
		*(*uint16)(unsafe.Pointer(&v7[1])) = *(*uint16)(unsafe.Pointer(uintptr(v6 + 36)))
		v5 = nox_xxx_unitArmorInventoryEquipFlags_415C70(object)
	}
	*(*uint32)(unsafe.Pointer(&v7[3])) = uint32(v5)
	return nox_xxx_netSendPacket1_4E5390(a1, int32(uintptr(unsafe.Pointer(&v7[0]))), 7, 0, 0)
}
func nox_xxx_netReportEquip_4D8540(a1 int32, a2 *uint32, a3 int32) *uint32 {
	var (
		result *uint32
		v4     [3]byte
	)
	v4[0] = 96
	*(*uint16)(unsafe.Pointer(&v4[1])) = uint16(nox_xxx_netGetUnitCodeServ_578AC0((*nox_object_t)(unsafe.Pointer(a2))))
	nox_xxx_netSendPacket1_4E5390(a1, int32(uintptr(unsafe.Pointer(&v4[0]))), 3, 0, 0)
	result = (*uint32)(unsafe.Pointer(uintptr(a3)))
	if a3 != 0 {
		result = sub_4D82F0(math.MaxUint8, a2)
	}
	return result
}
func nox_xxx_netReportDequip_4D8590(a1 int32, object *nox_object_t) int32 {
	var (
		a2 *uint32 = (*uint32)(unsafe.Pointer(object))
		v4 [3]byte
	)
	v4[0] = 97
	*(*uint16)(unsafe.Pointer(&v4[1])) = uint16(nox_xxx_netGetUnitCodeServ_578AC0((*nox_object_t)(unsafe.Pointer(a2))))
	return nox_xxx_netSendPacket1_4E5390(a1, int32(uintptr(unsafe.Pointer(&v4[0]))), 3, 0, 0)
}
func nox_xxx_netReportTotalHealth_4D85C0(a1 int32, a2 *uint32) int32 {
	var (
		result int32
		v3     *uint16
		v4     [7]byte
	)
	result = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*139)))
	if result != 0 {
		v4[0] = 221
		*(*uint16)(unsafe.Pointer(&v4[1])) = uint16(nox_xxx_netGetUnitCodeServ_578AC0((*nox_object_t)(unsafe.Pointer(a2))))
		v3 = (*uint16)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*139)))))
		*(*uint16)(unsafe.Pointer(&v4[3])) = *v3
		*(*uint16)(unsafe.Pointer(&v4[5])) = *(*uint16)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint16(0))*2))
		result = nox_xxx_netSendPacket1_4E5390(a1, int32(uintptr(unsafe.Pointer(&v4[0]))), 7, 0, 1)
	}
	return result
}
func nox_xxx_netReportUnitCurrentHP_4D8620(a1 int32, a2 *uint32) int32 {
	var (
		v2     *uint32
		result int32
		v4     uint32
		v5     *uint32
	)
	v2 = a2
	result = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*139)))
	if result != 0 {
		v5 = a2
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a2))), 0)) = 65
		*(*uint16)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(&a2))), 1)))) = uint16(nox_xxx_netGetUnitCodeServ_578AC0((*nox_object_t)(unsafe.Pointer(v5))))
		*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v4))), unsafe.Sizeof(uint16(0))*0)) = *(*uint16)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*139)))))
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a2))), unsafe.Sizeof((*uint32)(nil))-1)) = uint8(v4 >> 1)
		result = nox_xxx_netSendPacket1_4E5390(a1, int32(uintptr(unsafe.Pointer(&a2))), 4, 0, 1)
	}
	return result
}
func nox_xxx_netSendTeam_4D8670(a1 int32, a2 *uint32) int32 {
	var (
		result int32
		v3     int16
		v4     *uint16
		v5     [5]byte
	)
	result = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*139)))
	if result != 0 {
		v5[0] = 196
		v5[1] = 12
		v3 = int16(uint16(nox_xxx_netGetUnitCodeServ_578AC0((*nox_object_t)(unsafe.Pointer(a2)))))
		v4 = (*uint16)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*139)))))
		*(*uint16)(unsafe.Pointer(&v5[2])) = uint16(v3)
		v5[4] = byte(int8(int32(*v4) * 100 / int32(*(*uint16)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint16(0))*2)))))
		result = nox_xxx_netSendPacket1_4E5390(a1, int32(uintptr(unsafe.Pointer(&v5[0]))), 5, 0, 1)
	}
	return result
}
func nox_xxx_netSendPlrHealthToTeam_4D86E0(a1 int32) *byte {
	var (
		v1     int32
		result *byte
		v3     *byte
	)
	v1 = a1
	result = (*byte)(unsafe.Pointer(nox_common_playerInfoFromNum_417090(a1)))
	v3 = result
	if result != nil {
		result = (*byte)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(result))), unsafe.Sizeof(uint32(0))*514))))))
		if result != nil {
			if *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(result))), unsafe.Sizeof(uint32(0))*139))) != 0 {
				*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a1))), 0)) = 67
				*(*uint16)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(&a1))), 1)))) = **(**uint16)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v3))), unsafe.Sizeof(uint32(0))*514))) + 556)))
				nox_xxx_netSendPacket1_4E5390(v1, int32(uintptr(unsafe.Pointer(&a1))), 3, 0, 1)
				result = (*byte)(unsafe.Pointer(uintptr(bool2int(noxflags.HasGame(noxflags.GameModeQuest)))))
				if result != nil {
					result = (*byte)(unsafe.Pointer(uintptr(nox_xxx_netSendTeam_4D8670(v1|128, *((**uint32)(unsafe.Add(unsafe.Pointer((**uint32)(unsafe.Pointer(v3))), unsafe.Sizeof((*uint32)(nil))*514)))))))
				}
			}
		}
	}
	return result
}
func nox_xxx_netReportHealthDelta_4D8760(a1 int32, a2 int16, a3 int16) int16 {
	var (
		result int16
		v4     [5]byte
	)
	result = a3
	if int32(a3) < 0 {
		*(*uint16)(unsafe.Pointer(&v4[3])) = uint16(a3)
		v4[0] = 66
		*(*uint16)(unsafe.Pointer(&v4[1])) = uint16(a2)
		result = int16(nox_xxx_netSendPacket0_4E5420(a1, unsafe.Pointer(&v4[0]), 5, 0, 1))
	}
	return result
}
func nox_xxx_itemReportHealth_4D87A0(a1 int32, item *nox_object_t) int32 {
	var (
		result int32
		v3     *uint16
		v4     [7]byte
	)
	result = int32(uintptr(item.data_139))
	if result != 0 {
		if int32(*(*uint16)(unsafe.Pointer(uintptr(result + 4)))) != 0 {
			v4[0] = 68
			*(*uint16)(unsafe.Pointer(&v4[1])) = uint16(nox_xxx_netGetUnitCodeServ_578AC0(item))
			v3 = (*uint16)(item.data_139)
			*(*uint16)(unsafe.Pointer(&v4[3])) = *(*uint16)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint16(0))*0))
			*(*uint16)(unsafe.Pointer(&v4[5])) = *(*uint16)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint16(0))*2))
			result = nox_xxx_netSendPacket1_4E5390(a1, int32(uintptr(unsafe.Pointer(&v4[0]))), 7, 0, 0)
		}
	}
	return result
}
func nox_xxx_netReportStamina_4D8800(a1 int32, a2 int32) int32 {
	var (
		result int32
		v3     int32
	)
	result = a2
	if int32(*(*uint8)(unsafe.Pointer(uintptr(a2 + 8))))&4 != 0 {
		v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 748))))
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a2))), 0)) = 71
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a2))), 1)) = *(*uint8)(unsafe.Pointer(uintptr(v3 + 91)))
		result = nox_xxx_netSendPacket1_4E5390(a1, int32(uintptr(unsafe.Pointer(&a2))), 2, 0, 1)
	}
	return result
}
func sub_4D8840(a1 int32, a2 int32) int32 {
	var (
		v3 int8
		v5 int16
	)
	v3 = int8(*(*uint8)(unsafe.Pointer(uintptr(a2 + 440))))
	*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v5))), 0)) = 91
	*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v5))), unsafe.Sizeof(int16(0))-1)) = uint8(v3)
	return nox_xxx_netSendPacket0_4E5420(a1, unsafe.Pointer(&v5), 2, 0, 1)
}
func sub_4D8870(a1 int32, a2 int32) int32 {
	var (
		result int32
		v3     int32
		v4     [5]byte
	)
	result = a2
	if int32(*(*uint8)(unsafe.Pointer(uintptr(a2 + 8))))&4 != 0 {
		v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 748))))
		v4[0] = 74
		*(*uint32)(unsafe.Pointer(&v4[1])) = *(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v3 + 276))) + 2164)))
		result = nox_xxx_netSendPacket0_4E5420(a1, unsafe.Pointer(&v4[0]), 5, 0, 1)
	}
	return result
}
func nox_xxx_netReportTotalMana_4D88C0(a1 int32, a2 int32) int32 {
	var (
		result int32
		v3     int32
		v4     [7]byte
	)
	result = a2
	v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 748))))
	if int32(*(*uint8)(unsafe.Pointer(uintptr(a2 + 8))))&4 != 0 && (v3 == 0 || int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v3 + 276))) + 2251)))) != 0) {
		v4[0] = 222
		*(*uint16)(unsafe.Pointer(&v4[1])) = uint16(nox_xxx_netGetUnitCodeServ_578AC0((*nox_object_t)(unsafe.Pointer(uintptr(a2)))))
		*(*uint16)(unsafe.Pointer(&v4[3])) = *(*uint16)(unsafe.Pointer(uintptr(v3 + 4)))
		*(*uint16)(unsafe.Pointer(&v4[5])) = *(*uint16)(unsafe.Pointer(uintptr(v3 + 8)))
		result = sub_4E5450(a1, &v4[0], 7, 0, 1)
	}
	return result
}
func nox_xxx_netReportMana_4D8930(a1 int32, a2 int32) int32 {
	var (
		result int32
		v3     int32
		v4     [5]byte
	)
	result = a2
	v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 748))))
	if int32(*(*uint8)(unsafe.Pointer(uintptr(a2 + 8))))&4 != 0 && (v3 == 0 || int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v3 + 276))) + 2251)))) != 0) {
		v4[0] = 69
		*(*uint16)(unsafe.Pointer(&v4[1])) = uint16(nox_xxx_netGetUnitCodeServ_578AC0((*nox_object_t)(unsafe.Pointer(uintptr(a2)))))
		*(*uint16)(unsafe.Pointer(&v4[3])) = *(*uint16)(unsafe.Pointer(uintptr(v3 + 4)))
		result = sub_4E5450(a1, &v4[0], 5, 0, 1)
	}
	return result
}
func nox_xxx_servSendStats_4D8990(a1 int32, a2 int32, a3 int8) int8 {
	var (
		result int8
		v4     int32
		v5     int16
		v6     int32
		v7     int16
		v8     [14]byte
	)
	result = int8(*(*uint8)(unsafe.Pointer(uintptr(a2 + 8))))
	v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 748))))
	if int32(result)&4 != 0 {
		v8[0] = 72
		v5 = int16(uint16(nox_xxx_netGetUnitCodeServ_578AC0((*nox_object_t)(unsafe.Pointer(uintptr(a2))))))
		v6 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 556))))
		*(*uint16)(unsafe.Pointer(&v8[1])) = uint16(v5)
		*(*uint16)(unsafe.Pointer(&v8[5])) = *(*uint16)(unsafe.Pointer(uintptr(v4 + 8)))
		v7 = int16(*(*uint16)(unsafe.Pointer(uintptr(a2 + 490))))
		*(*uint16)(unsafe.Pointer(&v8[3])) = *(*uint16)(unsafe.Pointer(uintptr(v6 + 4)))
		*(*uint16)(unsafe.Pointer(&v8[7])) = uint16(v7)
		*(*uint16)(unsafe.Pointer(&v8[9])) = *(*uint16)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v4 + 276))) + 2235)))
		*(*uint16)(unsafe.Pointer(&v8[11])) = *(*uint16)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v4 + 276))) + 2239)))
		v8[13] = byte(a3)
		result = int8(nox_xxx_netSendPacket0_4E5420(a1, unsafe.Pointer(&v8[0]), 14, 0, 1))
	}
	return result
}
func nox_xxx_netReportArmorVal_4D8A30(a1 int32, a2 int32) int32 {
	var v3 [5]byte
	v3[0] = 73
	*(*uint32)(unsafe.Pointer(&v3[1])) = uint32(a2)
	return nox_xxx_netSendPacket0_4E5420(a1, unsafe.Pointer(&v3[0]), 5, 0, 1)
}
func nox_xxx_netReportPickup_4D8A60(a1 int32, item *nox_object_t) int32 {
	var (
		v3 int16
		v4 int16
		v5 [5]byte
	)
	if item.obj_class&0x13001000 != 0 {
		return nox_xxx_netReportModifiablePickup_4D8AD0(a1, item)
	}
	v5[0] = 75
	v3 = int16(uint16(nox_xxx_netGetUnitCodeServ_578AC0(item)))
	v4 = int16(item.typ_ind)
	*(*uint16)(unsafe.Pointer(&v5[1])) = uint16(v3)
	*(*uint16)(unsafe.Pointer(&v5[3])) = uint16(v4)
	nox_xxx_netSendPacket1_4E5390(a1, int32(uintptr(unsafe.Pointer(&v5[0]))), 5, 0, 0)
	return nox_xxx_itemReportHealth_4D87A0(a1, item)
}
func nox_xxx_netReportModifiablePickup_4D8AD0(a1 int32, item *nox_object_t) int32 {
	var (
		v2 int32
		v3 int32
		v4 int32
		v6 [9]byte
	)
	v6[0] = 76
	v2 = int32(uintptr(item.init_data))
	*(*uint16)(unsafe.Pointer(&v6[1])) = uint16(nox_xxx_netGetUnitCodeServ_578AC0(item))
	*(*uint16)(unsafe.Pointer(&v6[3])) = item.typ_ind
	v3 = 0
	v4 = v2
	for {
		if *(*uint32)(unsafe.Pointer(uintptr(v4))) != 0 {
			v6[v3+5] = byte(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v4))) + 4))))
		} else {
			v6[v3+5] = math.MaxUint8
		}
		v3++
		v4 += 4
		if v3 >= 4 {
			break
		}
	}
	nox_xxx_netSendPacket1_4E5390(a1, int32(uintptr(unsafe.Pointer(&v6[0]))), 9, 0, 0)
	return nox_xxx_itemReportHealth_4D87A0(a1, item)
}
func nox_xxx_netReportDrop_4D8B50(a1 int32, object *nox_object_t) int32 {
	var (
		a2 int32 = int32(uintptr(unsafe.Pointer(object)))
		v3 [5]byte
	)
	v3[0] = 77
	*(*uint16)(unsafe.Pointer(&v3[1])) = uint16(nox_xxx_netGetUnitCodeServ_578AC0((*nox_object_t)(unsafe.Pointer(uintptr(a2)))))
	*(*uint16)(unsafe.Pointer(&v3[3])) = *(*uint16)(unsafe.Pointer(uintptr(a2 + 4)))
	return nox_xxx_netSendPacket1_4E5390(a1, int32(uintptr(unsafe.Pointer(&v3[0]))), 5, 0, 0)
}
func nox_xxx_netSendDMWinner_4D8B90(a1 int32, a2 int8) int32 {
	var (
		result int32
		v3     [8]byte
	)
	result = a1
	if a1 != 0 {
		if (int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8)))) & 4) == 0 {
			return result
		}
		v3[0] = 88
		*(*uint16)(unsafe.Pointer(&v3[1])) = uint16(nox_xxx_netGetUnitCodeServ_578AC0((*nox_object_t)(unsafe.Pointer(uintptr(a1)))))
	} else {
		v3[0] = 88
		*(*uint16)(unsafe.Pointer(&v3[1])) = 0
	}
	v3[3] = byte(a2)
	*(*uint32)(unsafe.Pointer(&v3[4])) = nox_frame_xxx_2598000
	return nox_xxx_netSendPacket1_4E5390(math.MaxUint8, int32(uintptr(unsafe.Pointer(&v3[0]))), 8, 0, 1)
}
func nox_xxx_netSendDMTeamWinner_4D8BF0(a1 int32, a2 int8) int32 {
	var v3 [8]byte
	v3[0] = 89
	if a1 != 0 {
		*(*uint16)(unsafe.Pointer(&v3[1])) = uint16(*(*uint8)(unsafe.Pointer(uintptr(a1 + 57))))
	} else {
		*(*uint16)(unsafe.Pointer(&v3[1])) = 0
	}
	v3[3] = byte(a2)
	*(*uint32)(unsafe.Pointer(&v3[4])) = nox_frame_xxx_2598000
	return nox_xxx_netSendPacket1_4E5390(math.MaxUint8, int32(uintptr(unsafe.Pointer(&v3[0]))), 8, 0, 1)
}
func nox_xxx_netFlagballWinner_4D8C40(a1 int32) int32 {
	var (
		v1 int16
		v3 [8]byte
	)
	v1 = int16(*(*uint8)(unsafe.Pointer(uintptr(a1 + 57))))
	v3[0] = 86
	*(*uint16)(unsafe.Pointer(&v3[1])) = uint16(v1)
	v3[3] = 0
	*(*uint32)(unsafe.Pointer(&v3[4])) = nox_frame_xxx_2598000
	return nox_xxx_netSendPacket1_4E5390(math.MaxUint8, int32(uintptr(unsafe.Pointer(&v3[0]))), 8, 0, 1)
}
func nox_xxx_netFlagWinner_4D8C40_4D8C80(a1 int32, a2 int8) int32 {
	var v3 [8]byte
	v3[0] = 87
	if a1 != 0 {
		*(*uint16)(unsafe.Pointer(&v3[1])) = uint16(*(*uint8)(unsafe.Pointer(uintptr(a1 + 57))))
	} else {
		*(*uint16)(unsafe.Pointer(&v3[1])) = math.MaxUint16
	}
	v3[3] = byte(a2)
	*(*uint32)(unsafe.Pointer(&v3[4])) = nox_frame_xxx_2598000
	return nox_xxx_netSendPacket1_4E5390(math.MaxUint8, int32(uintptr(unsafe.Pointer(&v3[0]))), 8, 0, 1)
}
func nox_xxx_scavengerHuntReport_4D8CD0(a1 int32) int32 {
	var (
		result int32
		v2     int32
		v3     [7]byte
	)
	result = a1
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
	if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))&4 != 0 {
		v3[0] = 85
		*(*uint16)(unsafe.Pointer(&v3[1])) = uint16(nox_xxx_netGetUnitCodeServ_578AC0((*nox_object_t)(unsafe.Pointer(uintptr(a1)))))
		*(*uint16)(unsafe.Pointer(&v3[3])) = *(*uint16)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 276))) + 2152)))
		*(*uint16)(unsafe.Pointer(&v3[5])) = *(*uint16)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 276))) + 2156)))
		result = nox_xxx_netSendPacket0_4E5420(math.MaxUint8, unsafe.Pointer(&v3[0]), 7, 0, 1)
	}
	return result
}
func nox_xxx_playerIncrementElimDeath_4D8D40(a1 int32) *byte {
	var (
		result *byte
		i      *byte
		v3     *byte
		v4     float64
		v5     int32
		v6     int32
		v7     float32
		v8     *uint8
		v9     int32
	)
	result = (*byte)(unsafe.Pointer(uintptr(a1)))
	if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))&4 != 0 {
		*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))) + 276))) + 2140)))++
		result = (*byte)(unsafe.Pointer(uintptr(bool2int(noxflags.HasGame(noxflags.GameModeElimination)))))
		if result != nil {
			if sub_40AA00() != 0 && sub_40AA20() == 0 {
				for i = (*byte)(unsafe.Pointer(nox_common_playerInfoGetFirst_416EA0())); i != nil; i = (*byte)(unsafe.Pointer(nox_common_playerInfoGetNext_416EE0((*nox_playerInfo)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(i))))))))) {
					if *(*byte)(unsafe.Add(unsafe.Pointer(i), 3680))&1 != 0 {
						nox_xxx_netNeedTimestampStatus_4174F0(int32(uintptr(unsafe.Pointer(i))), 256)
					}
				}
				sub_40AA30(1)
			}
			result = (*byte)(unsafe.Pointer(uintptr(bool2int(noxflags.HasGame(noxflags.GameSuddenDeath)))))
			if result == nil {
				result = (*byte)(unsafe.Pointer(uintptr(sub_40A300())))
				if result == nil {
					result = (*byte)(unsafe.Pointer(uintptr(sub_40AA00())))
					if result != nil {
						if !nox_xxx_CheckGameplayFlags_417DA0(4) {
							v5 = sub_40A770()
							result = (*byte)(unsafe.Pointer(uintptr(sub_40AA40())))
							if v5 >= int32(uintptr(unsafe.Pointer(result))) {
								return result
							}
							v8 = (*uint8)(memmap.PtrOff(0x587000, 198928))
							v4 = nox_xxx_gamedataGetFloat_419D40(CString("SuddenDeathCountdown"))
							goto LABEL_22
						}
						v9 = int32(nox_xxx_getTeamCounter_417DD0())
						result = (*byte)(unsafe.Pointer(uintptr(sub_40AA40())))
						if v9 < int32(uintptr(unsafe.Pointer(result))) {
							result = (*byte)(unsafe.Pointer(nox_server_teamFirst_418B10()))
							v3 = result
							if result != nil {
								for nox_xxx_countNonEliminatedPlayersInTeam_40A830(int32(uintptr(unsafe.Pointer(v3)))) != 1 {
									result = (*byte)(unsafe.Pointer(nox_server_teamNext_418B60((*nox_team_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v3)))))))))
									v3 = result
									if result == nil {
										return result
									}
								}
								v8 = (*uint8)(memmap.PtrOff(0x587000, 198872))
								v4 = nox_xxx_gamedataGetFloat_419D40(CString("SuddenDeathCountdown"))
							LABEL_22:
								v7 = float32(v4)
								v6 = int(v7)
								return nox_xxx_servStartCountdown_40A2A0(v6, (*byte)(unsafe.Pointer(v8)))
							}
						}
					}
				}
			}
		}
	}
	return result
}
func nox_xxx_changeScore_4D8E90(a1 int32, a2 int32) int32 {
	var result int32
	result = a1
	if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))&4 != 0 {
		result = int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))) + 276))))
		*(*uint32)(unsafe.Pointer(uintptr(result + 2136))) += uint32(a2)
	}
	return result
}
func nox_xxx_playerSubLessons_4D8EC0(a1 int32, a2 int32) int32 {
	var result int32
	result = a1
	if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))&4 != 0 {
		result = int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))) + 276))))
		*(*uint32)(unsafe.Pointer(uintptr(result + 2136))) -= uint32(a2)
	}
	return result
}
func nox_xxx_netReportLesson_4D8EF0(a1p *nox_object_t) int32 {
	var (
		a1 int32 = int32(uintptr(unsafe.Pointer(a1p)))
		v1 int32
		v3 [11]byte
	)
	v3[0] = 78
	v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
	*(*uint16)(unsafe.Pointer(&v3[1])) = *(*uint16)(unsafe.Pointer(uintptr(a1 + 36)))
	*(*uint32)(unsafe.Pointer(&v3[3])) = *(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v1 + 276))) + 2136)))
	*(*uint32)(unsafe.Pointer(&v3[7])) = *(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v1 + 276))) + 2140)))
	return nox_xxx_netSendPacket1_4E5390(math.MaxUint8, int32(uintptr(unsafe.Pointer(&v3[0]))), 11, 0, 1)
}
func nox_xxx_netTimerStatus_4D8F50(a1 int32, a2 int32) int32 {
	var v3 [13]byte
	v3[0] = 211
	*(*uint32)(unsafe.Pointer(&v3[1])) = uint32(a2)
	*(*uint32)(unsafe.Pointer(&v3[5])) = uint32(sub_40A230())
	*(*uint32)(unsafe.Pointer(&v3[9])) = nox_frame_xxx_2598000
	return nox_xxx_netSendPacket1_4E5390(a1, int32(uintptr(unsafe.Pointer(&v3[0]))), 13, 0, 1)
}
func nox_xxx_netReportEnchant_4D8F90(a1 int32, a2 *uint32) int32 {
	var v3 [7]byte
	v3[0] = 90
	*(*uint16)(unsafe.Pointer(&v3[1])) = uint16(nox_xxx_netGetUnitCodeServ_578AC0((*nox_object_t)(unsafe.Pointer(a2))))
	*(*uint32)(unsafe.Pointer(&v3[3])) = *(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*85))
	return nox_xxx_netSendPacket1_4E5390(a1, int32(uintptr(unsafe.Pointer(&v3[0]))), 7, 0, 1)
}
func nox_xxx_netReportObjHidden_4D8FD0(a1 int32, a2 *uint32) {
	var v2 *uint32
	v2 = a2
	if *(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*2))&0x800000 != 0 {
		*(*uint16)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(&a2))), 1)))) = uint16(nox_xxx_netGetUnitCodeServ_578AC0((*nox_object_t)(unsafe.Pointer(a2))))
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a2))), 0)) = uint8(int8(56 - bool2int((*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*4))&0x1000000) != 0)))
		nox_xxx_netSendPacket0_4E5420(a1, unsafe.Pointer(&a2), 3, 0, 1)
	}
}
func nox_xxx_netReportUnitHeight_4D9020(a1 int32, a2 int32) int32 {
	var (
		v2     int32
		v3     int16
		v4     float64
		v5     int64
		v6     float64
		v7     int64
		v8     float64
		result int32
		v10    int16
		v11    float64
		v12    int16
		v13    float64
		v14    *uint32
		v15    [6]byte
	)
	v2 = a2
	if int32(*(*uint8)(unsafe.Pointer(uintptr(a2 + 20))))&32 != 0 {
		v3 = int16(*(*uint16)(unsafe.Pointer(uintptr(a2 + 36))))
		v15[0] = 159
		v4 = float64(*(*float32)(unsafe.Pointer(uintptr(a2 + 104))))
		*(*uint16)(unsafe.Pointer(&v15[1])) = uint16(v3)
		v5 = int64(v4)
		v6 = float64(*(*float32)(unsafe.Pointer(uintptr(a2 + 108))))
		v15[3] = byte(int8(v5))
		v7 = int64(v6)
		v8 = float64(*(*float32)(unsafe.Pointer(uintptr(a2 + 116))))
		v15[4] = byte(int8(v7))
		v15[5] = byte(int8(int64(v8)))
		result = nox_netlist_addToMsgListCli_40EBC0(a1, 1, (*uint8)(unsafe.Pointer(&v15[0])), 6)
	} else {
		v14 = (*uint32)(unsafe.Pointer(uintptr(a2)))
		if float64(*(*float32)(unsafe.Pointer(uintptr(a2 + 104)))) < 0.0 {
			*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a2))), 0)) = 95
			v12 = int16(uint16(nox_xxx_netGetUnitCodeServ_578AC0((*nox_object_t)(unsafe.Pointer(v14)))))
			v13 = float64(*(*float32)(unsafe.Pointer(uintptr(v2 + 104))))
			*(*uint16)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(&a2))), 1)))) = uint16(v12)
			v11 = -v13
		} else {
			*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a2))), 0)) = 94
			v10 = int16(uint16(nox_xxx_netGetUnitCodeServ_578AC0((*nox_object_t)(unsafe.Pointer(v14)))))
			v11 = float64(*(*float32)(unsafe.Pointer(uintptr(v2 + 104))))
			*(*uint16)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(&a2))), 1)))) = uint16(v10)
		}
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a2))), unsafe.Sizeof(int32(0))-1)) = uint8(int8(int64(v11)))
		result = nox_netlist_addToMsgListCli_40EBC0(a1, 1, (*uint8)(unsafe.Pointer(&a2)), 4)
	}
	return result
}
func sub_4D90E0(a1 int32, a2 int8) int32 {
	var v3 [2]byte
	v3[0] = 151
	v3[1] = byte(a2)
	return nox_netlist_addToMsgListCli_40EBC0(a1, 1, (*uint8)(unsafe.Pointer(&v3[0])), 2)
}
func nox_xxx_earthquakeSend_4D9110(a1 *float32, a2 int32) int32 {
	var (
		result int32
		i      int32
		v4     int32
		v5     float64
		v6     float64
		v7     float64
	)
	result = int32(uintptr(unsafe.Pointer(nox_xxx_getFirstPlayerUnit_4DA7C0())))
	for i = result; result != 0; i = result {
		v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(i + 748))) + 276))))
		v5 = float64(*a1 - *(*float32)(unsafe.Pointer(uintptr(v4 + 3632))))
		v6 = float64(*(*float32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(float32(0))*1)) - *(*float32)(unsafe.Pointer(uintptr(v4 + 3636))))
		v7 = v6*v6 + v5*v5
		if v7 < 90000.0 {
			sub_4D90E0(int32(*(*uint8)(unsafe.Pointer(uintptr(v4 + 2064)))), int8(int64((1.0-v7*1.1111111e-05)*float64(a2))))
		}
		result = int32(uintptr(unsafe.Pointer(nox_xxx_getNextPlayerUnit_4DA7F0((*nox_object_t)(unsafe.Pointer(uintptr(i)))))))
	}
	return result
}
func nox_xxx_netReportAcquireCreature_4D91A0(a1 int32, a2 int32) int32 {
	var v3 [5]byte
	v3[0] = 108
	*(*uint16)(unsafe.Pointer(&v3[1])) = uint16(nox_xxx_netGetUnitCodeServ_578AC0((*nox_object_t)(unsafe.Pointer(uintptr(a2)))))
	*(*uint16)(unsafe.Pointer(&v3[3])) = *(*uint16)(unsafe.Pointer(uintptr(a2 + 4)))
	if noxflags.HasGame(noxflags.GameFlag28) {
		v3[4] |= 128
	}
	nox_xxx_netSendPacket1_4E5390(a1, int32(uintptr(unsafe.Pointer(&v3[0]))), 5, 0, 1)
	return nox_xxx_netReportTotalHealth_4D85C0(a1, (*uint32)(unsafe.Pointer(uintptr(a2))))
}
func nox_xxx_netFxShield_0_4D9200(a1 int32, a2 int32) int32 {
	var v4 int32
	*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v4))), 0)) = 109
	*(*uint16)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(&v4))), 1)))) = *(*uint16)(unsafe.Pointer(uintptr(a2 + 36)))
	if noxflags.HasGame(noxflags.GameFlag28) {
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v4))), 2)) |= 128
	}
	return nox_xxx_netSendPacket1_4E5390(a1, int32(uintptr(unsafe.Pointer(&v4))), 3, 0, 1)
}
func nox_xxx_netMonitorCreature_4D9250(a1 int32, a2 int32) int32 {
	var v3 [5]byte
	v3[0] = 219
	*(*uint16)(unsafe.Pointer(&v3[1])) = uint16(nox_xxx_netGetUnitCodeServ_578AC0((*nox_object_t)(unsafe.Pointer(uintptr(a2)))))
	*(*uint16)(unsafe.Pointer(&v3[3])) = *(*uint16)(unsafe.Pointer(uintptr(a2 + 4)))
	nox_xxx_netSendPacket1_4E5390(a1, int32(uintptr(unsafe.Pointer(&v3[0]))), 5, 0, 1)
	return nox_xxx_netReportTotalHealth_4D85C0(a1, (*uint32)(unsafe.Pointer(uintptr(a2))))
}
func nox_xxx_netSendUnMonitorCrea_4D92A0(a1 int32, a2 *uint32) int32 {
	var v4 [3]byte
	v4[0] = 220
	*(*uint16)(unsafe.Pointer(&v4[1])) = uint16(nox_xxx_netGetUnitCodeServ_578AC0((*nox_object_t)(unsafe.Pointer(a2))))
	return nox_xxx_netSendPacket1_4E5390(a1, int32(uintptr(unsafe.Pointer(&v4[0]))), 3, 0, 1)
}
func nox_xxx_netReportTeamBase_4D92D0(a1 int32, a2 int32) int32 {
	var (
		result int32
		v3     int32
		v4     int32
		v5     int32
		v6     [7]byte
	)
	result = int32(*memmap.PtrUint32(6112660, 1556320))
	v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 692))))
	if *memmap.PtrUint32(6112660, 1556320) == 0 {
		result = nox_xxx_getNameId_4E3AA0(CString("TeamBase"))
		*memmap.PtrUint32(6112660, 1556320) = uint32(result)
	}
	if *(*uint32)(unsafe.Pointer(uintptr(a2 + 8)))&0x13001000 != 0 || int32(*(*uint16)(unsafe.Pointer(uintptr(a2 + 4)))) == result {
		v6[0] = 103
		*(*uint16)(unsafe.Pointer(&v6[1])) = uint16(nox_xxx_netGetUnitCodeServ_578AC0((*nox_object_t)(unsafe.Pointer(uintptr(a2)))))
		v4 = 0
		v5 = v3
		for {
			if *(*uint32)(unsafe.Pointer(uintptr(v5))) != 0 {
				v6[v4+3] = byte(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v5))) + 4))))
			} else {
				v6[v4+3] = math.MaxUint8
			}
			v4++
			v5 += 4
			if v4 >= 4 {
				break
			}
		}
		result = nox_xxx_netSendPacket0_4E5420(a1, unsafe.Pointer(&v6[0]), 7, 0, 1)
	}
	return result
}
func nox_xxx_netReportStatsSpeed_4D9360(a1 int32, a2 *uint32, a3 int8, a4 int32) int32 {
	var v5 [8]byte
	v5[0] = 104
	*(*uint16)(unsafe.Pointer(&v5[1])) = uint16(nox_xxx_netGetUnitCodeServ_578AC0((*nox_object_t)(unsafe.Pointer(a2))))
	v5[7] = byte(a3)
	*(*uint32)(unsafe.Pointer(&v5[3])) = uint32(a4)
	return nox_xxx_netSendPacket0_4E5420(a1, unsafe.Pointer(&v5[0]), 8, 0, 1)
}
func nox_xxx_netSendReportNPC_4D93A0(a1 int32, a2 int32) *uint32 {
	var (
		v2 int32
		v3 *byte
		v4 int32
		v5 int32
		v6 int32
		v7 *byte
	)
	_ = v7
	var result *uint32
	var i *uint32
	var v10 [21]byte
	v10[0] = 105
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 748))))
	*(*uint16)(unsafe.Pointer(&v10[1])) = *(*uint16)(unsafe.Pointer(uintptr(a2 + 36)))
	if int32(*(*uint8)(unsafe.Pointer(uintptr(a2 + 540)))) != 0 {
		v10[2] |= 128
	}
	v3 = &v10[3]
	v4 = v2 + 2076
	v5 = 6
	for {
		v6 = v4
		v7 = v3
		v4 += 3
		v3 = (*byte)(unsafe.Add(unsafe.Pointer(v3), 3))
		v5--
		*(*uint16)(unsafe.Pointer(v7)) = *(*uint16)(unsafe.Pointer(uintptr(v6)))
		*(*byte)(unsafe.Add(unsafe.Pointer(v7), 2)) = byte(*(*uint8)(unsafe.Pointer(uintptr(v6 + 2))))
		if v5 == 0 {
			break
		}
	}
	result = (*uint32)(unsafe.Pointer(uintptr(nox_xxx_netSendPacket1_4E5390(a1, int32(uintptr(unsafe.Pointer(&v10[0]))), 21, 0, 1))))
	for i = *(**uint32)(unsafe.Pointer(uintptr(a2 + 504))); i != nil; i = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(i), unsafe.Sizeof(uint32(0))*124))))) {
		if *(*uint32)(unsafe.Add(unsafe.Pointer(i), unsafe.Sizeof(uint32(0))*4))&256 != 0 {
			result = sub_4D82F0(a1, i)
		}
	}
	return result
}
func nox_xxx_netSendJournalAdd_4D9440(a1 int32, a2 int32) int32 {
	var v3 [68]byte
	v3[0] = 213
	v3[1] = 1
	libc.StrCpy(&v3[2], (*byte)(unsafe.Pointer(uintptr(a2))))
	*(*uint16)(unsafe.Pointer(&v3[66])) = *(*uint16)(unsafe.Pointer(uintptr(a2 + 72)))
	return nox_xxx_netSendPacket0_4E5420(a1, unsafe.Pointer(&v3[0]), 68, 0, 1)
}
func nox_xxx_netSendJournalRemove_4D94A0(a1 int32, a2 *byte) int32 {
	var v3 [68]byte
	v3[0] = 213
	v3[1] = 2
	libc.StrCpy(&v3[2], a2)
	return nox_xxx_netSendPacket0_4E5420(a1, unsafe.Pointer(&v3[0]), 68, 0, 1)
}
func nox_xxx_netSendJournalUpdate_4D9500(a1 int32, a2 int32) int32 {
	var v3 [68]byte
	v3[0] = 213
	v3[1] = 3
	libc.StrCpy(&v3[2], (*byte)(unsafe.Pointer(uintptr(a2))))
	*(*uint16)(unsafe.Pointer(&v3[66])) = *(*uint16)(unsafe.Pointer(uintptr(a2 + 72)))
	return nox_xxx_netSendPacket0_4E5420(a1, unsafe.Pointer(&v3[0]), 68, 0, 1)
}
func nox_xxx_netSendChapterEnd_4D9560(a1 int32, a2 int8, a3 int32) int32 {
	var v5 [3]byte
	v5[1] = byte(a2)
	v5[0] = 214
	v5[2] = byte(int8(bool2int(a3 == 1)))
	return nox_xxx_netSendPacket0_4E5420(a1, unsafe.Pointer(&v5[0]), 3, 0, 1)
}
func nox_xxx_netSendFlagStatus_4D95A0(a1 int32, a2 int8, a3 int8, a4 int8, a5 int16) int32 {
	var v6 [6]byte
	v6[1] = byte(a3)
	v6[3] = byte(a4)
	v6[2] = byte(a2)
	v6[0] = 216
	*(*uint16)(unsafe.Pointer(&v6[4])) = uint16(a5)
	return nox_xxx_netSendPacket1_4E5390(a1, int32(uintptr(unsafe.Pointer(&v6[0]))), 6, 0, 1)
}
func nox_xxx_netSendBallStatus_4D95F0(a1 int32, a2 int8, a3 int16) int32 {
	var v4 [4]byte
	v4[1] = byte(a2)
	v4[0] = 217
	*(*uint16)(unsafe.Pointer(&v4[2])) = uint16(a3)
	return nox_xxx_netSendPacket1_4E5390(a1, int32(uintptr(unsafe.Pointer(&v4[0]))), 4, 0, 1)
}
func nox_xxx_netReportSpellStat_4D9630(a1 int32, a2 int32, a3 int8) int32 {
	var v4 [6]byte
	*(*uint32)(unsafe.Pointer(&v4[1])) = uint32(a2)
	v4[0] = 223
	v4[5] = byte(a3)
	return nox_xxx_netSendPacket0_4E5420(a1, unsafe.Pointer(&v4[0]), 6, 0, 1)
}
func nox_xxx_netSendSecondaryWeapon_4D9670(a1 int32, a2 *uint32, a3 int8) int32 {
	var v4 [4]byte
	v4[0] = 224
	*(*uint16)(unsafe.Pointer(&v4[1])) = uint16(nox_xxx_netGetUnitCodeServ_578AC0((*nox_object_t)(unsafe.Pointer(a2))))
	v4[3] = byte(a3)
	return nox_xxx_netSendPacket1_4E5390(a1, int32(uintptr(unsafe.Pointer(&v4[0]))), 4, 0, 0)
}
func nox_xxx_netMsgLastQuiver_4D96B0(a1 int32, a2 *uint32) int32 {
	var v4 [3]byte
	v4[0] = 225
	*(*uint16)(unsafe.Pointer(&v4[1])) = uint16(nox_xxx_netGetUnitCodeServ_578AC0((*nox_object_t)(unsafe.Pointer(a2))))
	return nox_xxx_netSendPacket1_4E5390(a1, int32(uintptr(unsafe.Pointer(&v4[0]))), 3, 0, 0)
}
func nox_xxx_netMsgInventoryLoaded_4D96E0(a1 int32) int32 {
	var v2 [1]byte
	v2[0] = 113
	return nox_xxx_netSendPacket1_4E5390(a1, int32(uintptr(unsafe.Pointer(&v2[0]))), 1, 0, 0)
}
func sub_4D9700(a1 int32) int32 {
	var result int32
	result = a1
	switch a1 {
	case 0:
		result = nox_xxx_netPrintLineToAll_4DA390(CString("report.c:NoComp"))
	case 1:
		result = nox_xxx_netPrintLineToAll_4DA390(CString("report.c:MinComp"))
	case 2:
		result = nox_xxx_netPrintLineToAll_4DA390(CString("report.c:AveComp"))
	case 3:
		result = nox_xxx_netPrintLineToAll_4DA390(CString("report.c:UserComp"))
	default:
		return result
	}
	return result
}
func nox_xxx_netScriptMessageKill_4D9760(a1 int32) int32 {
	var (
		result int32
		v2     int32
	)
	result = a1
	if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))&4 != 0 {
		v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a1))), 0)) = 203
		result = nox_xxx_netSendPacket0_4E5420(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 276))) + 2064)))), unsafe.Pointer(&a1), 1, 0, 1)
	}
	return result
}
func nox_xxx_netFriendAddRemove_4D97A0(a1 int32, a2 *uint32, a3 int32) int32 {
	*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a3))), 0)) = uint8(int8(bool2int(a3 != 1) + 52))
	*(*uint16)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(&a3))), 1)))) = uint16(nox_xxx_netGetUnitCodeServ_578AC0((*nox_object_t)(unsafe.Pointer(a2))))
	return nox_xxx_netSendPacket1_4E5390(a1, int32(uintptr(unsafe.Pointer(&a3))), 3, 0, 1)
}
func sub_4D97E0(a1 int32) int32 {
	var v2 [1]byte
	v2[0] = 54
	return nox_xxx_netSendPacket1_4E5390(a1, int32(uintptr(unsafe.Pointer(&v2[0]))), 1, 0, 1)
}
func nox_xxx_netMsgFadeBegin_4D9800(a1 int32, a2 int32) int32 {
	var v4 [3]byte
	v4[0] = 228
	v4[1] = byte(int8(bool2int(a1 != 0)))
	v4[2] = byte(int8(bool2int(a2 != 0)))
	return nox_xxx_netSendPacket1_4E5390(math.MaxUint8, int32(uintptr(unsafe.Pointer(&v4[0]))), 3, 0, 1)
}
func nox_xxx_netMsgFadeBeginPlayer(pl int32, a1 int32, a2 int32) int32 {
	var v4 [3]byte
	v4[0] = 228
	v4[1] = byte(int8(bool2int(a1 != 0)))
	v4[2] = byte(int8(bool2int(a2 != 0)))
	return nox_xxx_netSendPacket1_4E5390(pl, int32(uintptr(unsafe.Pointer(&v4[0]))), 3, 0, 1)
}
func nox_xxx_netHarpoonAttach_4D9840(a1 *uint32, a2 *uint32) {
	var v3 [7]byte
	if a1 != nil && a2 != nil {
		v3[0] = 158
		v3[1] = 7
		v3[2] = 0
		*(*uint16)(unsafe.Pointer(&v3[3])) = uint16(nox_xxx_netGetUnitCodeServ_578AC0((*nox_object_t)(unsafe.Pointer(a1))))
		*(*uint16)(unsafe.Pointer(&v3[5])) = uint16(nox_xxx_netGetUnitCodeServ_578AC0((*nox_object_t)(unsafe.Pointer(a2))))
	}
	nox_xxx_netSendPacket1_4E5390(math.MaxUint8, int32(uintptr(unsafe.Pointer(&v3[0]))), 7, 0, 1)
}
func nox_xxx_netHarpoonBreak_4D98A0(a1 *uint32, a2 *uint32) int32 {
	var v3 [7]byte
	if a1 != nil && a2 != nil {
		v3[0] = 158
		v3[1] = 14
		v3[2] = 0
		*(*uint16)(unsafe.Pointer(&v3[3])) = uint16(nox_xxx_netGetUnitCodeServ_578AC0((*nox_object_t)(unsafe.Pointer(a1))))
		*(*uint16)(unsafe.Pointer(&v3[5])) = uint16(nox_xxx_netGetUnitCodeServ_578AC0((*nox_object_t)(unsafe.Pointer(a2))))
	}
	return nox_xxx_netSendPacket1_4E5390(math.MaxUint8, int32(uintptr(unsafe.Pointer(&v3[0]))), 7, 0, 1)
}
func nox_xxx_playerReportAnything_4D9900(a1 int32) {
	var (
		v1  int32
		v2  int32
		v3  *uint16
		v4  int32
		v5  int32
		i   int32
		v7  *byte
		j   int32
		v9  int32
		v10 int32
		k   int32
		v12 int32
		v13 int32
		v14 int32
		v15 *uint16
		v16 int32
	)
	v1 = a1
	if a1 != 0 && int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))&4 != 0 {
		v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
		v3 = *(**uint16)(unsafe.Pointer(uintptr(a1 + 556)))
		v15 = *(**uint16)(unsafe.Pointer(uintptr(a1 + 556)))
		v16 = *(*int32)(unsafe.Pointer(uintptr(v2 + 228)))
		if *(*float32)(unsafe.Pointer(uintptr(v2 + 232))) != *(*float32)(unsafe.Pointer(&v16)) {
			nox_xxx_netReportArmorVal_4D8A30(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 276))) + 2064)))), v16)
			*(*uint32)(unsafe.Pointer(uintptr(v2 + 232))) = *(*uint32)(unsafe.Pointer(uintptr(v2 + 228)))
		}
		v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 276))))
		if *(*uint32)(unsafe.Pointer(uintptr(v4 + 2168))) != *(*uint32)(unsafe.Pointer(uintptr(v4 + 2164))) {
			sub_4D8870(int32(*(*uint8)(unsafe.Pointer(uintptr(v4 + 2064)))), v1)
			*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 276))) + 2168))) = *(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 276))) + 2164)))
		}
		v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 276))))
		if int32(*(*uint8)(unsafe.Pointer(uintptr(v5 + 2172)))) != int32(*(*uint8)(unsafe.Pointer(uintptr(v1 + 440)))) {
			sub_4D8840(int32(*(*uint8)(unsafe.Pointer(uintptr(v5 + 2064)))), v1)
			*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 276))) + 2172))) = *(*uint8)(unsafe.Pointer(uintptr(v1 + 440)))
		}
		if noxflags.HasGame(noxflags.GameModeQuest) {
			for i = 0; i < 32; i++ {
				v7 = (*byte)(unsafe.Pointer(nox_common_playerInfoFromNum_417090(i)))
				if v7 != nil && *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v7))), unsafe.Sizeof(uint32(0))*514))) != 0 && *(*uint32)(unsafe.Pointer(uintptr(v2 + 320))) != uint32(*(*uint8)(unsafe.Pointer(uintptr(i + v2 + 452)))) {
					sub_4D9D60(i, v1)
					*(*uint8)(unsafe.Pointer(uintptr(i + v2 + 452))) = *(*uint8)(unsafe.Pointer(uintptr(v2 + 320)))
				}
			}
			for j = 0; j < 32; j++ {
				v9 = 0
				if *memmap.PtrUint32(6112660, 1556324) == 0 {
					*memmap.PtrUint32(6112660, 1556324) = uint32(nox_xxx_getNameId_4E3AA0(CString("SilverKey")))
				}
				if nox_common_playerInfoFromNum_417090(j) != nil {
					v10 = v1.FirstItem()
					if v10 != 0 {
						for uint32(*(*uint16)(unsafe.Pointer(uintptr(v10 + 4)))) != *memmap.PtrUint32(6112660, 1556324) {
							v10 = nox_xxx_inventoryGetNext_4E7990(v10)
							if v10 == 0 {
								goto LABEL_25
							}
						}
						v9 = 1
					}
				LABEL_25:
					if v9 != int32(*(*uint8)(unsafe.Pointer(uintptr(j + v2 + 484)))) {
						sub_4D9DF0(j, v1, int8(v9))
						*(*uint8)(unsafe.Pointer(uintptr(j + v2 + 484))) = uint8(int8(v9))
					}
				}
			}
			for k = 0; k < 32; k++ {
				v12 = 0
				if *memmap.PtrUint32(6112660, 1556328) == 0 {
					*memmap.PtrUint32(6112660, 1556328) = uint32(nox_xxx_getNameId_4E3AA0(CString("GoldKey")))
				}
				if nox_common_playerInfoFromNum_417090(k) != nil {
					v13 = v1.FirstItem()
					if v13 != 0 {
						for uint32(*(*uint16)(unsafe.Pointer(uintptr(v13 + 4)))) != *memmap.PtrUint32(6112660, 1556328) {
							v13 = nox_xxx_inventoryGetNext_4E7990(v13)
							if v13 == 0 {
								goto LABEL_37
							}
						}
						v12 = 1
					}
				LABEL_37:
					if v12 != int32(*(*uint8)(unsafe.Pointer(uintptr(k + v2 + 516)))) {
						sub_4D9E30(k, v1, int8(v12))
						*(*uint8)(unsafe.Pointer(uintptr(k + v2 + 516))) = uint8(int8(v12))
					}
				}
			}
			v3 = v15
		}
		v14 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 276))))
		if int32(*(*uint8)(unsafe.Pointer(uintptr(v14 + 2184)))) != 0 {
			nox_xxx_netReportTotalHealth_4D85C0(int32(*(*uint8)(unsafe.Pointer(uintptr(v14 + 2064)))), (*uint32)(unsafe.Pointer(uintptr(v1))))
			nox_xxx_netReportTotalMana_4D88C0(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 276))) + 2064)))), v1)
			nox_xxx_servSendStats_4D8990(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 276))) + 2064)))), v1, int8(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 276))) + 3684)))))
			*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 276))) + 2184))) = 0
		}
		if nox_xxx_check_flag_aaa_43AF70() == 1 {
			sub_528030(v1)
		} else {
			if v3 != nil && int32(*v3) != int32(*(*uint16)(unsafe.Pointer(uintptr(v2 + 10)))) {
				nox_xxx_netSendPlrHealthToTeam_4D86E0(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 276))) + 2064)))))
				*(*uint16)(unsafe.Pointer(uintptr(v2 + 10))) = *v3
			}
			if int32(*(*uint16)(unsafe.Pointer(uintptr(v2 + 4)))) != int32(*(*uint16)(unsafe.Pointer(uintptr(v2 + 6)))) {
				nox_xxx_netReportMana_4D8930(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 276))) + 2064)))), v1)
				*(*uint16)(unsafe.Pointer(uintptr(v2 + 6))) = *(*uint16)(unsafe.Pointer(uintptr(v2 + 4)))
			}
		}
	}
}
func nox_xxx_netStatsMultiplier_4D9C20(a1 int32) int32 {
	var (
		result int32
		v2     int32
		v3     int32
		v4     [17]byte
	)
	result = a1
	if a1 != 0 {
		v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
		v4[0] = 239
		if int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 276))) + 2251)))) != 0 {
			if int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 276))) + 2251)))) == 1 {
				*(*uint32)(unsafe.Pointer(&v4[1])) = nox_xxx_wizardMaxHealth_587000_312816
				*(*uint32)(unsafe.Pointer(&v4[5])) = nox_xxx_wizardMaximumMana_587000_312820
				v3 = int32(nox_xxx_wizardSpeed_587000_312828)
				*(*uint32)(unsafe.Pointer(&v4[9])) = nox_xxx_wizardStrength_587000_312824
			} else {
				result = int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 276))) + 2251)))) - 2
				if int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 276))) + 2251)))) != 2 {
					return result
				}
				*(*uint32)(unsafe.Pointer(&v4[1])) = nox_xxx_conjurerMaxHealth_587000_312800
				*(*uint32)(unsafe.Pointer(&v4[5])) = nox_xxx_conjurerMaxMana_587000_312804
				v3 = int32(nox_xxx_conjurerSpeed_587000_312812)
				*(*uint32)(unsafe.Pointer(&v4[9])) = nox_xxx_conjurerStrength_587000_312808
			}
		} else {
			*(*uint32)(unsafe.Pointer(&v4[1])) = nox_xxx_warriorMaxHealth_587000_312784
			*(*uint32)(unsafe.Pointer(&v4[5])) = nox_xxx_warriorMaxMana_587000_312788
			v3 = int32(nox_xxx_warriorMaxSpeed_587000_312796)
			*(*uint32)(unsafe.Pointer(&v4[9])) = nox_xxx_warriorMaxStrength_587000_312792
		}
		*(*uint32)(unsafe.Pointer(&v4[13])) = uint32(v3)
		result = nox_xxx_netSendPacket0_4E5420(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 276))) + 2064)))), unsafe.Pointer(&v4[0]), 17, 0, 1)
	}
	return result
}
func sub_4D9CF0(a1 int32) int32 {
	var v2 [2]byte
	v2[0] = 240
	v2[1] = 0
	return nox_xxx_netSendPacket1_4E5390(a1, int32(uintptr(unsafe.Pointer(&v2[0]))), 2, 0, 1)
}
func sub_4D9D20(a1 int32, a2p *nox_object_t) int32 {
	var (
		a2 int32 = int32(uintptr(unsafe.Pointer(a2p)))
		v3 int16
		v5 [4]byte
	)
	v3 = int16(*(*uint16)(unsafe.Pointer(uintptr(a2 + 36))))
	*(*uint16)(unsafe.Pointer(&v5[0])) = 496
	*(*uint16)(unsafe.Pointer(&v5[2])) = uint16(v3)
	return nox_xxx_netSendPacket1_4E5390(a1, int32(uintptr(unsafe.Pointer(&v5[0]))), 4, 0, 1)
}
func sub_4D9D60(a1 int32, a2 int32) int32 {
	var (
		v2 int32
		v4 [5]byte
	)
	v4[0] = 240
	v4[1] = 4
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 748))))
	*(*uint16)(unsafe.Pointer(&v4[3])) = *(*uint16)(unsafe.Pointer(uintptr(a2 + 36)))
	v4[2] = byte(*(*uint8)(unsafe.Pointer(uintptr(v2 + 320))))
	return nox_xxx_netSendPacket0_4E5420(a1, unsafe.Pointer(&v4[0]), 5, 0, 1)
}
func sub_4D9DF0(a1 int32, a2 int32, a3 int8) int32 {
	var v4 [5]byte
	*(*uint16)(unsafe.Pointer(&v4[3])) = *(*uint16)(unsafe.Pointer(uintptr(a2 + 36)))
	v4[0] = 240
	v4[1] = 22
	v4[2] = byte(a3)
	return nox_xxx_netSendPacket0_4E5420(a1, unsafe.Pointer(&v4[0]), 5, 0, 1)
}
func sub_4D9E30(a1 int32, a2 int32, a3 int8) int32 {
	var v4 [5]byte
	*(*uint16)(unsafe.Pointer(&v4[3])) = *(*uint16)(unsafe.Pointer(uintptr(a2 + 36)))
	v4[0] = 240
	v4[1] = 23
	v4[2] = byte(a3)
	return nox_xxx_netSendPacket0_4E5420(a1, unsafe.Pointer(&v4[0]), 5, 0, 1)
}
func nox_xxx_netGauntlet_4D9E70(a1 int32) int32 {
	var v2 [2]byte
	v2[0] = 240
	v2[1] = 20
	return nox_xxx_netSendPacket0_4E5420(a1, unsafe.Pointer(&v2[0]), 2, 0, 1)
}
func nox_xxx_netSendLineMessage_4D9EB0(a1 int32, a2 *libc.WChar, _rest ...interface{}) int32 {
	var (
		result int32
		v3     int32
		v4     int8
		v5     int32
		v6     [516]libc.WChar
		va     libc.ArgList
	)
	va.Start(a2, _rest)
	result = a1
	if a1 != 0 && int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))&4 != 0 {
		v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
		nox_vswprintf(&v6[260], a2, va)
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v6[0]))), 0)) = 168
		*(*libc.WChar)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(&v6[0]))), 1)))) = 0
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v6[1]))), unsafe.Sizeof(libc.WChar(0))-1)) = 0
		if nox_xxx_cliCanTalkMB_4100F0((*int16)(unsafe.Pointer(&v6[260]))) != 0 {
			v4 = int8(int32(*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v6[1]))), unsafe.Sizeof(libc.WChar(0))-1))) | 2)
		} else {
			v4 = int8(int32(*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v6[1]))), unsafe.Sizeof(libc.WChar(0))-1))) | 4)
		}
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v6[1]))), unsafe.Sizeof(libc.WChar(0))-1)) = uint8(v4)
		v6[2] = 0
		v6[3] = 0
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v6[5]))), 0)) = 0
		v6[4] = libc.WChar(uint8(nox_wcslen(&v6[260]) + 1))
		if v6[1]&1024 != 0 {
			nox_wcscpy((*libc.WChar)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(&v6[5]))), 1)))), &v6[260])
			v5 = 2
		} else {
			nox_sprintf((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(&v6[5]))), 1)), CString("%S"), &v6[260])
			v5 = 1
		}
		result = nox_netlist_addToMsgListCli_40EBC0(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v3 + 276))) + 2064)))), 1, (*uint8)(unsafe.Pointer(&v6[0])), v5*int32(*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v6[4]))), 0)))+11)
	}
	return result
}
func nox_xxx_printToAll_4D9FD0(a1 int8, a2 *libc.WChar, _rest ...interface{}) int32 {
	var (
		v2     int8
		v3     int32
		result int32
		i      int32
		v6     [516]libc.WChar
		va     libc.ArgList
	)
	va.Start(a2, _rest)
	nox_vswprintf(&v6[260], a2, va)
	*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v6[0]))), 0)) = 168
	*(*libc.WChar)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(&v6[0]))), 1)))) = 0
	*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v6[1]))), unsafe.Sizeof(libc.WChar(0))-1)) = uint8(a1)
	if nox_xxx_cliCanTalkMB_4100F0((*int16)(unsafe.Pointer(&v6[260]))) != 0 {
		v2 = int8(int32(*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v6[1]))), unsafe.Sizeof(libc.WChar(0))-1))) | 2)
	} else {
		v2 = int8(int32(*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v6[1]))), unsafe.Sizeof(libc.WChar(0))-1))) | 4)
	}
	*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v6[1]))), unsafe.Sizeof(libc.WChar(0))-1)) = uint8(v2)
	v6[2] = 0
	v6[3] = 0
	*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v6[5]))), 0)) = 0
	v6[4] = libc.WChar(uint8(nox_wcslen(&v6[260]) + 1))
	if v6[1]&1024 != 0 {
		nox_wcscpy((*libc.WChar)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(&v6[5]))), 1)))), &v6[260])
		v3 = 2
	} else {
		nox_sprintf((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(&v6[5]))), 1)), CString("%S"), &v6[260])
		v3 = 1
	}
	result = int32(uintptr(unsafe.Pointer(nox_xxx_getFirstPlayerUnit_4DA7C0())))
	for i = result; result != 0; i = result {
		nox_netlist_addToMsgListCli_40EBC0(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(i + 748))) + 276))) + 2064)))), 1, (*uint8)(unsafe.Pointer(&v6[0])), v3*int32(*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v6[4]))), 0)))+11)
		result = int32(uintptr(unsafe.Pointer(nox_xxx_getNextPlayerUnit_4DA7F0((*nox_object_t)(unsafe.Pointer(uintptr(i)))))))
	}
	return result
}
func nox_xxx_netInformTextMsg_4DA0F0(a1 int32, a2 int32, a3 *int32) int32 {
	var (
		result int32
		v4     int32
		v5     [6]byte
	)
	result = a2
	switch a2 {
	case 0:
		fallthrough
	case 1:
		fallthrough
	case 2:
		fallthrough
	case 12:
		fallthrough
	case 13:
		fallthrough
	case 16:
		fallthrough
	case 20:
		fallthrough
	case 21:
		v5[1] = byte(int8(a2))
		v4 = *a3
		v5[0] = 169
		*(*uint32)(unsafe.Pointer(&v5[2])) = uint32(v4)
		result = nox_netlist_addToMsgListCli_40EBC0(a1, 1, (*uint8)(unsafe.Pointer(&v5[0])), 6)
	case 17:
		*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&a2))), unsafe.Sizeof(uint16(0))*0)) = 4521
		result = nox_netlist_addToMsgListCli_40EBC0(a1, 1, (*uint8)(unsafe.Pointer(&a2)), 2)
	default:
		return result
	}
	return result
}
func nox_xxx_netInformTextMsg2_4DA180(a1 int32, a2 *uint8) int32 {
	var (
		result int32
		i      int32
		j      int32
		k      int32
		v6     [6]byte
	)
	result = a1
	switch a1 {
	case 3:
		fallthrough
	case 4:
		fallthrough
	case 8:
		fallthrough
	case 18:
		fallthrough
	case 19:
		fallthrough
	case 21:
		v6[1] = byte(int8(a1))
		v6[0] = 169
		*(*uint32)(unsafe.Pointer(&v6[2])) = *(*uint32)(unsafe.Pointer(a2))
		result = int32(uintptr(unsafe.Pointer(nox_xxx_getFirstPlayerUnit_4DA7C0())))
		for i = result; result != 0; i = result {
			nox_netlist_addToMsgListCli_40EBC0(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(i + 748))) + 276))) + 2064)))), 1, (*uint8)(unsafe.Pointer(&v6[0])), 6)
			result = int32(uintptr(unsafe.Pointer(nox_xxx_getNextPlayerUnit_4DA7F0((*nox_object_t)(unsafe.Pointer(uintptr(i)))))))
		}
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
		*a2 = 169
		*(*uint8)(unsafe.Add(unsafe.Pointer(a2), 1)) = uint8(int8(a1))
		result = int32(uintptr(unsafe.Pointer(nox_xxx_getFirstPlayerUnit_4DA7C0())))
		for j = result; result != 0; j = result {
			nox_netlist_addToMsgListCli_40EBC0(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(j + 748))) + 276))) + 2064)))), 1, a2, 10)
			result = int32(uintptr(unsafe.Pointer(nox_xxx_getNextPlayerUnit_4DA7F0((*nox_object_t)(unsafe.Pointer(uintptr(j)))))))
		}
	case 14:
		*a2 = 169
		*(*uint8)(unsafe.Add(unsafe.Pointer(a2), 1)) = uint8(int8(a1))
		result = int32(uintptr(unsafe.Pointer(nox_xxx_getFirstPlayerUnit_4DA7C0())))
		for k = result; result != 0; k = result {
			nox_netlist_addToMsgListCli_40EBC0(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(k + 748))) + 276))) + 2064)))), 1, a2, 11)
			result = int32(uintptr(unsafe.Pointer(nox_xxx_getNextPlayerUnit_4DA7F0((*nox_object_t)(unsafe.Pointer(uintptr(k)))))))
		}
	default:
		return result
	}
	return result
}
func nox_xxx_netPriMsgToPlayer_4DA2C0(a1p *nox_object_t, a2 *byte, a3 int8) {
	var (
		a1 int32 = int32(uintptr(unsafe.Pointer(a1p)))
		v3 int32
		v4 [52]byte
	)
	if a1 != 0 && int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))&4 != 0 && a2 != nil && sub_419E60(a1) == 0 && libc.StrLen(a2) != 0 && libc.StrLen(a2) <= 48 {
		v4[2] = byte(a3)
		v4[0] = 169
		v4[1] = 15
		v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
		libc.StrCpy(&v4[3], a2)
		nox_netlist_addToMsgListCli_40EBC0(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v3 + 276))) + 2064)))), 1, (*uint8)(unsafe.Pointer(&v4[0])), int32(libc.StrLen(a2)+4))
	}
}
func nox_xxx_netPrintLineToAll_4DA390(a1 *byte) int32 {
	var (
		result int32
		i      int32
	)
	result = int32(uintptr(unsafe.Pointer(nox_xxx_getFirstPlayerUnit_4DA7C0())))
	for i = result; result != 0; i = result {
		nox_xxx_netPriMsgToPlayer_4DA2C0((*nox_object_t)(unsafe.Pointer(uintptr(i))), a1, 0)
		result = int32(uintptr(unsafe.Pointer(nox_xxx_getNextPlayerUnit_4DA7F0((*nox_object_t)(unsafe.Pointer(uintptr(i)))))))
	}
	return result
}
func nox_get_and_zero_server_objects_4DA3C0() *uint32 {
	var p *uint32 = (*uint32)(unsafe.Pointer(nox_server_objects_1556844))
	nox_server_objects_1556844 = nil
	return p
}
func nox_set_server_objects_4DA3E0(p unsafe.Pointer) {
	nox_server_objects_1556844 = (*nox_object_t)(p)
}
func nox_server_strcmpWithoutMapname_4DA3F0(a1 *byte, a2 *byte) int32 {
	var (
		v2     *byte
		v3     *byte
		v4     *byte
		v5     *byte
		v6     *byte
		v7     *byte
		result int32
	)
	libc.StrCpy((*byte)(memmap.PtrOff(6112660, 1556332)), a1)
	libc.StrCpy((*byte)(memmap.PtrOff(6112660, 1556588)), a2)
	v2 = libc.StrChr((*byte)(memmap.PtrOff(6112660, 1556332)), 58)
	if v2 != nil {
		v3 = (*byte)(memmap.PtrOff(6112660, 1556332))
		*v2 = 0
		v4 = (*byte)(unsafe.Add(unsafe.Pointer(v2), 1))
	} else {
		v3 = nil
		v4 = (*byte)(memmap.PtrOff(6112660, 1556332))
	}
	v5 = libc.StrChr((*byte)(memmap.PtrOff(6112660, 1556588)), 58)
	if v5 != nil {
		v6 = (*byte)(memmap.PtrOff(6112660, 1556588))
		*v5 = 0
		v7 = (*byte)(unsafe.Add(unsafe.Pointer(v5), 1))
	} else {
		v6 = nil
		v7 = (*byte)(memmap.PtrOff(6112660, 1556588))
	}
	if v3 != nil && v6 != nil && libc.StrCaseCmp(v3, v6) != 0 {
		result = 0
	} else {
		result = bool2int(libc.StrCmp(v4, v7) == 0)
	}
	return result
}
func nox_xxx_getObjectByScrName_4DA4F0(a1 *byte) int32 {
	var (
		i      int32
		result int32
		v3     int32
		v4     int32
		v5     int32
	)
	if libc.StrChr(a1, byte(':')) != nil {
		for i = int32(uintptr(unsafe.Pointer(noxServer.firstServerObject()))); i != 0; i = int32(uintptr(unsafe.Pointer((*nox_object_t)(unsafe.Pointer(uintptr(i))).Next()))) {
			result = sub_4DA5C0(i, a1)
			if result != 0 {
				return result
			}
		}
		v3 = int32(uintptr(unsafe.Pointer(noxServer.firstServerObjectUninited())))
		if v3 != 0 {
			for {
				result = sub_4DA5C0(v3, a1)
				if result != 0 {
					break
				}
				result = int32(uintptr(unsafe.Pointer(nox_server_getNextObjectUninited_4DA880((*nox_object_t)(unsafe.Pointer(uintptr(v3)))))))
				v3 = result
				if result == 0 {
					break
				}
			}
			return result
		}
		return 0
	}
	v4 = int32(uintptr(unsafe.Pointer(noxServer.firstServerObject())))
	if v4 != 0 {
		for {
			result = sub_4DA660(v4, a1)
			if result != 0 {
				break
			}
			v4 = int32(uintptr(unsafe.Pointer((*nox_object_t)(unsafe.Pointer(uintptr(v4))).Next())))
			if v4 == 0 {
				goto LABEL_12
			}
		}
	} else {
	LABEL_12:
		v5 = int32(uintptr(unsafe.Pointer(noxServer.firstServerObjectUninited())))
		if v5 == 0 {
			return 0
		}
		for {
			result = sub_4DA660(v5, a1)
			if result != 0 {
				break
			}
			v5 = int32(uintptr(unsafe.Pointer(nox_server_getNextObjectUninited_4DA880((*nox_object_t)(unsafe.Pointer(uintptr(v5)))))))
			if v5 == 0 {
				return 0
			}
		}
	}
	return result
}
func sub_4DA5C0(a1 int32, a2 *byte) int32 {
	var v2 int32
	v2 = a1
	if *(*uint32)(unsafe.Pointer(uintptr(a1))) != 0 && libc.StrCmp(*(**byte)(unsafe.Pointer(uintptr(a1))), a2) == 0 {
		return v2
	}
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 504))))
	if v2 != 0 {
		for *(*uint32)(unsafe.Pointer(uintptr(v2))) == 0 || libc.StrCmp(*(**byte)(unsafe.Pointer(uintptr(v2))), a2) != 0 {
			v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 496))))
			if v2 == 0 {
				return 0
			}
		}
		return v2
	}
	return 0
}
func sub_4DA660(a1 int32, a2 *byte) int32 {
	var (
		i   int32
		v3  *byte
		v4  *byte
		v5  *byte
		v6  bool
		v7  uint8
		v8  uint8
		v9  int32
		v10 *byte
		v11 uint8
		v12 uint8
		v14 *byte
	)
	i = a1
	if *(*uint32)(unsafe.Pointer(uintptr(a1))) != 0 {
		v3 = libc.StrChr(*(**byte)(unsafe.Pointer(uintptr(a1))), 58)
		v4 = a2
		if v3 != nil {
			v5 = (*byte)(unsafe.Add(unsafe.Pointer(v3), 1))
			for {
				v6 = uint32(*v5) < uint32(*v4)
				if *v5 != *v4 {
					break
				}
				if *v5 != 0 {
					v7 = uint8(*(*byte)(unsafe.Add(unsafe.Pointer(v5), 1)))
					v8 = uint8(*(*byte)(unsafe.Add(unsafe.Pointer(v4), 1)))
					v6 = int32(v7) < int32(v8)
					if int32(v7) != int32(v8) {
						break
					}
					v5 = (*byte)(unsafe.Add(unsafe.Pointer(v5), 2))
					v4 = (*byte)(unsafe.Add(unsafe.Pointer(v4), 2))
					if int32(v7) != 0 {
						continue
					}
				}
				v9 = 0
				goto LABEL_16
			}
		} else {
			v10 = *(**byte)(unsafe.Pointer(uintptr(a1)))
			for {
				v6 = uint32(*v10) < uint32(*v4)
				if *v10 != *v4 {
					break
				}
				if *v10 != 0 {
					v11 = uint8(*(*byte)(unsafe.Add(unsafe.Pointer(v10), 1)))
					v12 = uint8(*(*byte)(unsafe.Add(unsafe.Pointer(v4), 1)))
					v6 = int32(v11) < int32(v12)
					if int32(v11) != int32(v12) {
						break
					}
					v10 = (*byte)(unsafe.Add(unsafe.Pointer(v10), 2))
					v4 = (*byte)(unsafe.Add(unsafe.Pointer(v4), 2))
					if int32(v11) != 0 {
						continue
					}
				}
				v9 = 0
				goto LABEL_16
			}
		}
		v9 = -bool2int(v6) - (bool2int(v6) - 1)
	LABEL_16:
		if v9 == 0 {
			return i
		}
	}
	for i = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 504)))); i != 0; i = int32(*(*uint32)(unsafe.Pointer(uintptr(i + 496)))) {
		if *(*uint32)(unsafe.Pointer(uintptr(i))) != 0 {
			v14 = libc.StrChr(*(**byte)(unsafe.Pointer(uintptr(i))), 58)
			if v14 != nil {
				if libc.StrCmp((*byte)(unsafe.Add(unsafe.Pointer(v14), 1)), a2) == 0 {
					return i
				}
			} else if libc.StrCmp(*(**byte)(unsafe.Pointer(uintptr(i))), a2) == 0 {
				return i
			}
		}
	}
	return 0
}
func nox_server_getFirstObject_4DA790() *nox_object_t {
	return nox_server_objects_1556844
}
func nox_server_getNextObject_4DA7A0(obj *nox_object_t) *nox_object_t {
	var (
		a1     int32 = int32(uintptr(unsafe.Pointer(obj)))
		result int32
	)
	if a1 != 0 {
		result = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 444))))
	} else {
		result = 0
	}
	return (*nox_object_t)(unsafe.Pointer(uintptr(result)))
}
func nox_xxx_getFirstPlayerUnit_4DA7C0() *nox_object_t {
	for p := (*nox_playerInfo)(nox_common_playerInfoGetFirst_416EA0()); p != nil; p = nox_common_playerInfoGetNext_416EE0(p) {
		if p.playerUnit != nil {
			return p.playerUnit
		}
	}
	return nil
}
func nox_xxx_getNextPlayerUnit_4DA7F0(obj *nox_object_t) *nox_object_t {
	var (
		a1 int32 = int32(uintptr(unsafe.Pointer(obj)))
		v1 *byte
	)
	if a1 == 0 {
		return nil
	}
	if (int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8)))) & 4) == 0 {
		return nil
	}
	v1 = (*byte)(unsafe.Pointer(nox_common_playerInfoGetNext_416EE0((*nox_playerInfo)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))) + 276)))))))))
	if v1 == nil {
		return nil
	}
	for *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v1))), unsafe.Sizeof(uint32(0))*514))) == 0 {
		v1 = (*byte)(unsafe.Pointer(nox_common_playerInfoGetNext_416EE0((*nox_playerInfo)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v1)))))))))
		if v1 == nil {
			return nil
		}
	}
	return (*nox_object_t)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v1))), unsafe.Sizeof(uint32(0))*514))))))
}
func nox_xxx_getFirstUpdatable2Object_4DA840() *nox_object_t {
	return nox_server_objects_updatable2_1556848
}
func nox_xxx_getNextUpdatable2Object_4DA850(obj *nox_object_t) *nox_object_t {
	var (
		a1     int32 = int32(uintptr(unsafe.Pointer(obj)))
		result int32
	)
	if a1 != 0 {
		result = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 444))))
	} else {
		result = 0
	}
	return (*nox_object_t)(unsafe.Pointer(uintptr(result)))
}
func nox_server_getFirstObjectUninited_4DA870() *nox_object_t {
	return nox_server_objects_uninited_1556860
}
func nox_server_getNextObjectUninited_4DA880(obj *nox_object_t) *nox_object_t {
	var (
		a1     int32 = int32(uintptr(unsafe.Pointer(obj)))
		result int32
	)
	if a1 != 0 {
		result = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 444))))
	} else {
		result = 0
	}
	return (*nox_object_t)(unsafe.Pointer(uintptr(result)))
}
func nox_xxx_getFirstUpdatableObject_4DA8A0() *nox_object_t {
	return nox_server_object_updateable_1556852
}
func nox_xxx_getNextUpdatableObject_4DA8B0(obj *nox_object_t) *nox_object_t {
	var (
		a1     int32 = int32(uintptr(unsafe.Pointer(obj)))
		result int32
	)
	if a1 != 0 {
		result = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 476))))
	} else {
		result = 0
	}
	return (*nox_object_t)(unsafe.Pointer(uintptr(result)))
}
func nox_xxx_unitAddToUpdatable_4DA8D0(obj *nox_object_t) *nox_object_t {
	var (
		a1     int32 = int32(uintptr(unsafe.Pointer(obj)))
		result int32
	)
	result = a1
	if *(*uint32)(unsafe.Pointer(uintptr(a1 + 484))) == 0 && (int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))&1) == 0 {
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 480))) = 0
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 476))) = uint32(uintptr(unsafe.Pointer(nox_server_object_updateable_1556852)))
		if nox_server_object_updateable_1556852 != nil {
			nox_server_object_updateable_1556852.field_120 = uint32(a1)
		}
		nox_server_object_updateable_1556852 = (*nox_object_t)(unsafe.Pointer(uintptr(a1)))
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 484))) = 1
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 520))) = 0
	}
	return (*nox_object_t)(unsafe.Pointer(uintptr(result)))
}
func nox_xxx_unitRemoveFromUpdatable_4DA920(a1 *uint32) *uint32 {
	var (
		result *uint32
		v2     int32
		v3     int32
		v4     int32
	)
	result = a1
	if *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*121)) != 0 {
		v2 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*120)))
		if v2 != 0 {
			*(*uint32)(unsafe.Pointer(uintptr(v2 + 476))) = *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*119))
			v3 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*119)))
			if v3 != 0 {
				*(*uint32)(unsafe.Pointer(uintptr(v3 + 480))) = *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*120))
				*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*121)) = 0
				*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*130)) = 0
				return result
			}
		} else {
			nox_server_object_updateable_1556852 = (*nox_object_t)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*119)))))
			v4 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*119)))
			if v4 != 0 {
				*(*uint32)(unsafe.Pointer(uintptr(v4 + 480))) = 0
			}
		}
		*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*121)) = 0
		*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*130)) = 0
	}
	return result
}
func nox_xxx_unitNewAddShadow_4DA9A0(a1 *uint32) *uint32 {
	var (
		result *uint32
		v2     int32
	)
	result = a1
	v2 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*4)))
	if (uint32(v2) & 0x410000) == 0 {
		*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*118)) = 0
		*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*4)) = uint32(v2) | 0x10000
		*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*117)) = dword_5d4594_1556856
		if dword_5d4594_1556856 != 0 {
			*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1556856 + 472))) = uint32(uintptr(unsafe.Pointer(a1)))
		}
		dword_5d4594_1556856 = uint32(uintptr(unsafe.Pointer(a1)))
	}
	return result
}
func nox_xxx_action_4DA9F0(a1 *uint32) *uint32 {
	var (
		result *uint32
		v2     int32
		v3     int32
		v4     int32
	)
	result = a1
	v2 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*4)))
	if uint32(v2)&0x10000 != 0 {
		*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*4)) = uint32(v2) & 0xFFFEFFFF
		v3 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*118)))
		if v3 != 0 {
			*(*uint32)(unsafe.Pointer(uintptr(v3 + 468))) = *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*117))
		} else {
			dword_5d4594_1556856 = *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*117))
		}
		v4 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*117)))
		if v4 != 0 {
			*(*uint32)(unsafe.Pointer(uintptr(v4 + 472))) = *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*118))
		}
	}
	return result
}
func nox_xxx_createAt_4DAA50(obj *nox_object_t, owner *nox_object_t, a3 float32, a4 float32) {
	var (
		a1 int32 = int32(uintptr(unsafe.Pointer(obj)))
		v4 int32
		v5 int32
		v6 int32
		v7 int32
	)
	if *memmap.PtrUint32(6112660, 1556864) == 0 {
		*memmap.PtrUint32(6112660, 1556864) = uint32(nox_xxx_getNameId_4E3AA0(CString("Gold")))
		*memmap.PtrUint32(6112660, 1556868) = uint32(nox_xxx_getNameId_4E3AA0(CString("QuestGoldPile")))
		*memmap.PtrUint32(6112660, 1556872) = uint32(nox_xxx_getNameId_4E3AA0(CString("QuestGoldChest")))
	}
	v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 16))))
	if v4&36 != 0 {
		return
	}
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 16))) = uint32(v4) & 0x35E9FEDB
	*(*float32)(unsafe.Pointer(uintptr(a1 + 72))) = a3
	*(*float32)(unsafe.Pointer(uintptr(a1 + 56))) = a3
	*(*float32)(unsafe.Pointer(uintptr(a1 + 64))) = a3
	*(*float32)(unsafe.Pointer(uintptr(a1 + 76))) = a4
	*(*float32)(unsafe.Pointer(uintptr(a1 + 60))) = a4
	v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 56))))
	*(*float32)(unsafe.Pointer(uintptr(a1 + 68))) = a4
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 156))) = uint32(v5)
	*(*float32)(unsafe.Pointer(uintptr(a1 + 160))) = a4
	nox_xxx_objectUnkUpdateCoords_4E7290(a1)
	if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))&6 != 0 {
		nox_xxx_unitPostCreateNotify_4E7F10(a1)
	}
	if owner != nil {
		nox_xxx_unitSetOwner_4EC290(owner, (*nox_object_t)(unsafe.Pointer(uintptr(a1))))
	}
	v6 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 16))))
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 80))) = 0
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 84))) = 0
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 88))) = 0
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 92))) = 0
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 16))) = uint32(v6 | 4)
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 136))) = nox_frame_xxx_2598000
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 128))) = nox_frame_xxx_2598000
	if noxflags.HasGame(noxflags.GameOnline) && !noxflags.HasGame(noxflags.GameModeQuest) && (uint32(*(*uint16)(unsafe.Pointer(uintptr(a1 + 4)))) == *memmap.PtrUint32(6112660, 1556864) || *(*uint32)(unsafe.Pointer(uintptr(a1 + 8)))&0x3001110 != 0) && (int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))&1) == 0 {
		v7 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 16))))
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v7))), 0)) = uint8(int8(v7 | 64))
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 16))) = uint32(v7)
	}
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 448))) = 0
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 444))) = uint32(uintptr(unsafe.Pointer(nox_server_objects_uninited_1556860)))
	if nox_server_objects_uninited_1556860 != nil {
		*(*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer(nox_server_objects_uninited_1556860))) + 448))) = uint32(a1)
	}
	nox_server_objects_uninited_1556860 = (*nox_object_t)(unsafe.Pointer(uintptr(a1)))
	var v4b uint8 = *(*uint8)(unsafe.Pointer(uintptr(a1 + 52)))
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 16))) |= 0x2000000
	if int32(v4b) != 0 && ((*(*uint32)(unsafe.Pointer(uintptr(a1 + 8)))&0x10000000) == 0 || int32(*memmap.PtrUint32(0x973F18, 3800)) >= 0) {
		if noxflags.HasGame(noxflags.GameModeCoop) || nox_xxx_CheckGameplayFlags_417DA0(4) {
			nox_xxx_createAtImpl_4191D0(*(*uint8)(unsafe.Pointer(uintptr(a1 + 52))), unsafe.Pointer(uintptr(a1+48)), 0, int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 36)))), 0)
		}
	}
}
func nox_xxx_unitsNewAddToList_4DAC00() {
	var (
		i  int32
		v3 int32
		v4 func(int32, uint32)
		v5 uint32
		v6 int32
	)
	if nox_server_objects_uninited_1556860 == nil {
		nox_server_objects_uninited_1556860 = nil
		return
	}
	for v0 := int32(int32(uintptr(unsafe.Pointer(nox_server_objects_uninited_1556860)))); v0 != 0; v0 = i {
		i = int32(*(*uint32)(unsafe.Pointer(uintptr(v0 + 444))))
		for j := int32(int32(*(*uint32)(unsafe.Pointer(uintptr(v0 + 508))))); j != 0; j = int32(*(*uint32)(unsafe.Pointer(uintptr(v0 + 508)))) {
			if (int32(*(*uint8)(unsafe.Pointer(uintptr(j + 16)))) & 32) == 0 {
				break
			}
			nox_xxx_unitSetOwner_4EC290((*nox_object_t)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(j + 508)))))), (*nox_object_t)(unsafe.Pointer(uintptr(v0))))
		}
		if int32(*(*uint8)(unsafe.Pointer(uintptr(v0 + 8))))&1 != 0 {
			*(*uint32)(unsafe.Pointer(uintptr(v0 + 448))) = 0
			*(*uint32)(unsafe.Pointer(uintptr(v0 + 444))) = uint32(uintptr(unsafe.Pointer(nox_server_objects_updatable2_1556848)))
			if nox_server_objects_updatable2_1556848 != nil {
				*(*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer(nox_server_objects_updatable2_1556848))) + 448))) = uint32(v0)
			}
			nox_server_objects_updatable2_1556848 = (*nox_object_t)(unsafe.Pointer(uintptr(v0)))
		} else {
			v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(v0 + 16))))
			if uint32(v3)&0x10000 != 0 {
				*(*uint32)(unsafe.Pointer(uintptr(v0 + 16))) = uint32(v3) & 0xFFFEFFFF
				nox_xxx_unitNewAddShadow_4DA9A0((*uint32)(unsafe.Pointer(uintptr(v0))))
			}
			if *(*uint32)(unsafe.Pointer(uintptr(v0 + 16)))&0x80000 != 0 && !noxflags.HasGame(noxflags.GameModeQuest) {
				nox_xxx_respawnAdd_4EC5E0(v0)
			}
			if (*(*uint32)(unsafe.Pointer(uintptr(v0 + 744))) != 0 || float64(*(*float32)(unsafe.Pointer(uintptr(v0 + 80)))) != 0.0 || float64(*(*float32)(unsafe.Pointer(uintptr(v0 + 84)))) != 0.0) && int32(int8(*(*uint8)(unsafe.Pointer(uintptr(v0 + 8))))) >= 0 {
				nox_xxx_unitAddToUpdatable_4DA8D0((*nox_object_t)(unsafe.Pointer(uintptr(v0))))
			}
			*(*uint32)(unsafe.Pointer(uintptr(v0 + 448))) = 0
			*(*uint32)(unsafe.Pointer(uintptr(v0 + 444))) = uint32(uintptr(unsafe.Pointer(nox_server_objects_1556844)))
			if nox_server_objects_1556844 != nil {
				*(*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer(nox_server_objects_1556844))) + 448))) = uint32(v0)
			}
			nox_set_server_objects_4DA3E0(unsafe.Pointer(uintptr(v0)))
		}
		nox_xxx_unitCreateMissileSmth_517640((*nox_object_t)(unsafe.Pointer(uintptr(v0))))
		if *(*uint32)(unsafe.Pointer(uintptr(v0 + 696))) != 0 {
			sub_5117F0(v0)
		}
		v4 = cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(v0 + 688))), (*func(int32, uint32))(nil))
		if v4 != nil {
			v4(v0, 0)
		}
		v5 = *(*uint32)(unsafe.Pointer(uintptr(v0 + 8)))
		if v5&0x400000 != 0 {
			if *(*uint32)(unsafe.Pointer(uintptr(v0 + 12)))&24 != 0 {
				v6 = 0
			} else {
				v6 = int32(^*(*uint8)(unsafe.Pointer(uintptr(v0 + 12)))) >> 7
			}
		} else {
			v6 = int32((v5 >> 29) & 1)
		}
		if v5&0x800000 != 0 || v6 == 0 {
			nox_xxx_unitNeedSync_4E44F0(v0)
			if (*(*uint32)(unsafe.Pointer(uintptr(v0 + 8))) & 0x20400000) == 0 {
				*(*uint32)(unsafe.Pointer(uintptr(v0 + 148))) = 0
			}
		} else {
			nox_xxx_objectMakeUnseenByNoone_4E44E0(v0)
			sub_527E00(v0)
			*(*uint32)(unsafe.Pointer(uintptr(v0 + 148))) = math.MaxUint32
		}
		*(*uint32)(unsafe.Pointer(uintptr(v0 + 16))) &= 0xFDFFFFFF
	}
	nox_server_objects_uninited_1556860 = nil
}
func nox_xxx_servFinalizeDelObject_4DADE0(item *nox_object_t) int8 {
	var (
		a1 int32 = int32(uintptr(unsafe.Pointer(item)))
		v1 int32
	)
	v1 = int32(item.obj_flags)
	if v1&4 != 0 {
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v1))), 0)) = uint8(int8(v1 & 251))
		item.obj_flags = uint32(v1)
		nox_xxx_playerLeaveObsByObserved_4E60A0(a1)
		if !noxflags.HasGame(noxflags.GameFlag20) {
			nox_xxx_netReportDestroyObject_5289D0(a1)
		}
		nox_xxx_unit_511810(a1)
		nox_xxx_unitClearOwner_4EC300(item)
		nox_xxx_unitRemoveChild_4EC470(a1)
		sub_517870((*nox_object_t)(unsafe.Pointer(uintptr(a1))))
		sub_4DAE50(a1)
		sub_4ECFA0(a1)
		sub_511DE0(a1)
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v1))), 0)) = *(*uint8)(unsafe.Pointer(&item.obj_class))
		if v1&6 != 0 {
			*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v1))), 0)) = uint8(int8(sub_528990(a1)))
		}
	}
	return int8(v1)
}
func sub_4DAE50(a1 int32) int32 {
	var (
		v1     int32
		result int32
		v3     int32
		v4     int32
		v5     int32
	)
	nox_xxx_action_4DA9F0((*uint32)(unsafe.Pointer(uintptr(a1))))
	if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))&1 != 0 {
		v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 448))))
		if v1 != 0 {
			*(*uint32)(unsafe.Pointer(uintptr(v1 + 444))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 444)))
			result = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 444))))
			if result != 0 {
				*(*uint32)(unsafe.Pointer(uintptr(result + 448))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 448)))
			}
		} else {
			result = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 444))))
			nox_server_objects_updatable2_1556848 = (*nox_object_t)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 444))))))
			v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 444))))
			if v3 != 0 {
				*(*uint32)(unsafe.Pointer(uintptr(v3 + 448))) = 0
			}
		}
	} else {
		nox_xxx_unitRemoveFromUpdatable_4DA920((*uint32)(unsafe.Pointer(uintptr(a1))))
		v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 448))))
		if v4 != 0 {
			*(*uint32)(unsafe.Pointer(uintptr(v4 + 444))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 444)))
			result = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 444))))
			if result != 0 {
				*(*uint32)(unsafe.Pointer(uintptr(result + 448))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 448)))
			}
		} else {
			result = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 444))))
			nox_set_server_objects_4DA3E0(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 444))))))
			v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 444))))
			if v5 != 0 {
				*(*uint32)(unsafe.Pointer(uintptr(v5 + 448))) = 0
			}
		}
	}
	return result
}
func sub_4DAF10() {
	var (
		v0 *uint32
		v1 *uint32
		v2 int32
		v3 int32
		v4 *uint32
		v5 int32
		v6 int32
		v7 int32
		v8 int32
		v9 *uint32
	)
	v0 = *(**uint32)(unsafe.Pointer(&nox_server_objects_uninited_1556860))
	v1 = *(**uint32)(unsafe.Pointer(&nox_server_objects_uninited_1556860))
	if nox_server_objects_uninited_1556860 != nil {
		for {
			v2 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*2)))
			if v2&0x4000 != 0 {
				v3 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*187)))
				v4 = v0
				if v0 != nil {
					for {
						v5 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*2)))
						if (v5 & 0x8000) != 0 {
							v6 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*187)))
							if *(*uint32)(unsafe.Pointer(uintptr(v3 + 8))) == *(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*10)) {
								break
							}
						}
						v4 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*111)))))
						if v4 == nil {
							goto LABEL_9
						}
					}
					*(*uint32)(unsafe.Pointer(uintptr(v3 + 4))) = uint32(uintptr(unsafe.Pointer(v4)))
					*(*uint32)(unsafe.Pointer(uintptr(v6 + 4))) = uint32(uintptr(unsafe.Pointer(v1)))
					v0 = *(**uint32)(unsafe.Pointer(&nox_server_objects_uninited_1556860))
				}
			}
		LABEL_9:
			if *(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*2))&1024 != 0 {
				v7 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*187)))
				v8 = int32(*(*uint32)(unsafe.Pointer(uintptr(v7 + 16))))
				*(*uint32)(unsafe.Pointer(uintptr(v7 + 12))) = 0
				v0 = *(**uint32)(unsafe.Pointer(&nox_server_objects_uninited_1556860))
				if v8 != 0 {
					v9 = *(**uint32)(unsafe.Pointer(&nox_server_objects_uninited_1556860))
					if nox_server_objects_uninited_1556860 != nil {
						for (*(*uint32)(unsafe.Add(unsafe.Pointer(v9), unsafe.Sizeof(uint32(0))*2))&1024) == 0 || uint32(v8) != *(*uint32)(unsafe.Add(unsafe.Pointer(v9), unsafe.Sizeof(uint32(0))*10)) {
							v9 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v9), unsafe.Sizeof(uint32(0))*111)))))
							if v9 == nil {
								goto LABEL_17
							}
						}
						*(*uint32)(unsafe.Pointer(uintptr(v7 + 12))) = uint32(uintptr(unsafe.Pointer(v9)))
						v0 = *(**uint32)(unsafe.Pointer(&nox_server_objects_uninited_1556860))
					}
				}
			}
		LABEL_17:
			v1 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*111)))))
			if v1 == nil {
				break
			}
		}
	}
}
func sub_4DAFD0(a1 *uint32) {
	var (
		v2 int32
		v3 int32
		v4 int32
	)
	v2 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*4)))
	if (v2 & 4) == 0 {
		return
	}
	*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*4)) = uint32(v2) & 0xFFFFFFFB
	v3 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*112)))
	if v3 != 0 {
		*(*uint32)(unsafe.Pointer(uintptr(v3 + 444))) = *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*111))
		v4 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*111)))
		if v4 != 0 {
			*(*uint32)(unsafe.Pointer(uintptr(v4 + 448))) = uint32(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*112))))))))
		}
	} else {
		var p *uint32 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*111)))))
		nox_set_server_objects_4DA3E0(unsafe.Pointer(p))
		if p != nil {
			*(*uint32)(unsafe.Add(unsafe.Pointer(p), unsafe.Sizeof(uint32(0))*112)) = 0
		}
	}
}
func nox_xxx_unitClearPendingMB_4DB030() *uint32 {
	var (
		result *uint32
		v1     *uint32
		v2     int32
	)
	result = *(**uint32)(unsafe.Pointer(&nox_server_objects_uninited_1556860))
	if nox_server_objects_uninited_1556860 != nil {
		for {
			v1 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*111)))))
			*(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*4)) &= 0xFDFFFFFF
			v2 = int32(uintptr(unsafe.Pointer(nox_server_objects_1556844)))
			if nox_server_objects_1556844 != nil {
				*(*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer(nox_server_objects_1556844))) + 448))) = uint32(uintptr(unsafe.Pointer(result)))
				v2 = int32(uintptr(unsafe.Pointer(nox_server_objects_1556844)))
			}
			*(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*111)) = uint32(v2)
			*(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*112)) = 0
			nox_set_server_objects_4DA3E0(unsafe.Pointer(result))
			result = v1
			if v1 == nil {
				break
			}
		}
	}
	nox_server_objects_uninited_1556860 = nil
	return result
}
func sub_4DB100() int32 {
	var result int32
	result = 0
	dword_5d4594_1563080 = 0
	dword_5d4594_1563084 = 0
	dword_5d4594_1563096 = 0
	dword_5d4594_1563064 = 0
	dword_5d4594_1563092 = 0
	dword_5d4594_1563088 = 0
	*memmap.PtrUint32(6112660, 1563072) = 0
	*memmap.PtrUint32(6112660, 1563068) = 0
	return result
}
func sub_4DB130(a1 *byte) uint32 {
	var result uint32
	result = uint32(libc.StrLen(a1) + 1)
	libc.MemCpy(memmap.PtrOff(6112660, 1557900), unsafe.Pointer(a1), int(result))
	return result
}
func nox_xxx_gameGet_4DB1B0() int32 {
	return int32(dword_5d4594_1563080)
}
func sub_4DB1C0() int32 {
	return int32(dword_5d4594_1563084)
}
func nox_xxx_gameSetSwitchSolo_4DB220(a1 int32) int32 {
	var result int32
	result = a1
	*memmap.PtrUint32(6112660, 1563056) = uint32(a1)
	return result
}
func nox_xxx_gameSetNoMPFlag_4DB230(a1 int32) int32 {
	var result int32
	result = a1
	*memmap.PtrUint32(6112660, 1563060) = uint32(a1)
	return result
}
func nox_xxx_gameIsSwitchToSolo_4DB240() int32 {
	return int32(*memmap.PtrUint32(6112660, 1563056))
}
func nox_xxx_gameIsNotMultiplayer_4DB250() int32 {
	return int32(*memmap.PtrUint32(6112660, 1563060))
}
func nox_xxx_mapFilenameGetSolo_4DB260() *byte {
	return (*byte)(memmap.PtrOff(6112660, 1559960))
}
func nox_xxx_gameSetSoloSavePath_4DB270(a1 *byte) uint32 {
	var result uint32
	result = uint32(libc.StrLen(a1) + 1)
	libc.MemCpy(memmap.PtrOff(6112660, 1559960), unsafe.Pointer(a1), int(result))
	return result
}
func nox_client_countSaveFiles_4DC550() int32 {
	var (
		v0       int32
		v1       *byte
		v5       *byte
		v6       int32
		v7       *byte
		PathName [1024]byte
	)
	v0 = 0
	v1 = nox_fs_root()
	libc.StrCpy(&PathName[0], v1)
	libc.StrCat(&PathName[0], CString("\\Save\\"))
	nox_fs_mkdir(&PathName[0])
	v5 = nox_fs_root()
	nox_sprintf(&PathName[0], CString("%s\\Save\\AUTOSAVE\\Player.plr"), v5)
	if nox_fs_access(&PathName[0], 0) != -1 {
		v0 = 1
	}
	v6 = 13
	for {
		v7 = nox_fs_root()
		nox_sprintf(&PathName[0], CString("%s\\Save\\SAVE%04d\\Player.plr"), v7, v0)
		if nox_fs_access(&PathName[0], 0) != -1 {
			v0++
		}
		v6--
		if v6 == 0 {
			break
		}
	}
	return v0
}
func nox_client_countPlayerFiles02_4DC630() int32 {
	var (
		v0           int32
		v1           *byte
		v5           HANDLE
		v6           *byte
		FindFileData _WIN32_FIND_DATAA
		PathName     [1024]byte
		v10          [1280]byte
		v11          [1024]byte
	)
	v0 = 0
	v1 = nox_fs_root()
	libc.StrCpy(&PathName[0], v1)
	libc.StrCat(&PathName[0], CString("\\Save\\"))
	libc.StrCpy(&v11[0], &PathName[0])
	nox_fs_mkdir(&PathName[0])
	nox_fs_set_workdir(&PathName[0])
	v5 = compatFindFirstFileA(CString("*.plr"), &FindFileData)
	if uintptr(v5) != uintptr(math.MaxUint32) {
		if (FindFileData.dwFileAttributes & 16) == 0 {
			nox_sprintf(&PathName[0], CString("%s%s"), &v11[0], &FindFileData.cFileName[0])
			sub_41A000(&PathName[0], (*nox_savegame_xxx)(unsafe.Pointer(&v10[0])))
			if v10[0]&2 != 0 {
				v0 = 1
			}
		}
		for compatFindNextFileA(v5, &FindFileData) != 0 {
			if (FindFileData.dwFileAttributes & 16) == 0 {
				nox_sprintf(&PathName[0], CString("%s%s"), &v11[0], &FindFileData.cFileName[0])
				sub_41A000(&PathName[0], (*nox_savegame_xxx)(unsafe.Pointer(&v10[0])))
				if v10[0]&2 != 0 {
					v0++
				}
			}
		}
		compatFindClose(v5)
	}
	v6 = nox_fs_root()
	nox_fs_set_workdir(v6)
	return v0
}
func sub_4DC9B0(a1 *byte) *byte {
	var (
		result *byte
		v2     [12]byte
	)
	if a1 == nil {
		return (*byte)(unsafe.Pointer(uintptr(math.MaxUint32)))
	}
	result = libc.StrRChr(a1, 92)
	if result != nil {
		libc.StrNCpy(&v2[0], (*byte)(unsafe.Add(unsafe.Pointer(result), -4)), 4)
		v2[4] = 0
		result = (*byte)(unsafe.Pointer(uintptr(libc.Atoi(libc.GoString(&v2[0])))))
	}
	return result
}
func sub_4DCBE0() int32 {
	return int32(*memmap.PtrUint32(6112660, 1563076))
}
func sub_4DCBF0(a1 int32) int32 {
	var result int32
	result = a1
	dword_5d4594_1563064 = uint32(a1)
	return result
}
func sub_4DCC10(a1p *nox_object_t) int32 {
	var (
		a1     int32 = int32(uintptr(unsafe.Pointer(a1p)))
		v1     int32
		v2     int32
		result int32
	)
	v1 = 1
	if dword_5d4594_1563092 != 0 && (dword_5d4594_1563092+dword_5d4594_1563088) > nox_frame_xxx_2598000 {
		v1 = 0
	}
	if *(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))) + 284))) != 0 {
		v1 = 0
	}
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 16))))
	if v2&2 != 0 {
		v1 = 0
	}
	if (v2 & 0x8000) != 0 {
		v1 = 0
	}
	if sub_45D9B0() == 1 {
		result = 0
	} else {
		result = v1
	}
	return result
}
func nox_xxx_mapLoadOrSaveMB_4DCC70(a1 int32) int32 {
	var result int32
	result = a1
	*memmap.PtrUint32(6112660, 1563072) = uint32(a1)
	return result
}
func sub_4DCC90() int32 {
	var result int32
	result = 1
	if dword_5d4594_1563080 != 1 {
		result = bool2int(dword_5d4594_1563092 != 0)
	}
	return result
}
func nox_xxx_game_4DCCB0() int32 {
	var (
		result int32
		v1     *byte
		v2     int32
		v3     int32
	)
	if !noxflags.HasGame(noxflags.GameModeCoop) {
		return 1
	}
	v1 = (*byte)(unsafe.Pointer(nox_common_playerInfoFromNum_417090(31)))
	if v1 == nil || (func() int32 {
		v2 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v1))), unsafe.Sizeof(uint32(0))*514))))
		return v2
	}()) == 0 || sub_4DCC90() != 0 || sub_4139B0() != 0 || (nox_frame_xxx_2598000-*memmap.PtrUint32(6112660, 1563068)) < 30 || nox_xxx_guiCursor_477600() != 0 || (func() int32 {
		v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 16))))
		return int32(*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v3))), 1))) & 64
	}()) != 0 {
		result = 0
	} else {
		result = bool2int(sub_4DCC10((*nox_object_t)(unsafe.Pointer(uintptr(v2)))) != 0)
	}
	return result
}
func sub_4DCD40() int32 {
	var (
		v0       *byte
		result   int32
		i        int32
		v3       int32
		v4       int32
		FileName [1024]byte
	)
	*memmap.PtrUint32(6112660, 1563048) = sub_416A00()
	sub_4169E0()
	dword_5d4594_1563044 = 1
	v0 = nox_fs_root()
	nox_sprintf(&FileName[0], CString("%s\\Save\\_temp_.dat"), v0)
	result = int32(uintptr(unsafe.Pointer(nox_xxx_getFirstPlayerUnit_4DA7C0())))
	for i = result; result != 0; i = result {
		v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(i + 748))))
		v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(v3 + 276))))
		if *(*uint32)(unsafe.Pointer(uintptr(v4 + 4792))) != 0 && *(*uint32)(unsafe.Pointer(uintptr(v3 + 552))) != 1 {
			if nox_xxx_playerSaveToFile_41A140(&FileName[0], int32(*(*uint8)(unsafe.Pointer(uintptr(v4 + 2064))))) != 0 {
				sub_41CFA0(&FileName[0], int32(*(*uint8)(unsafe.Pointer(uintptr(v4 + 2064)))))
			}
			nox_fs_remove(&FileName[0])
		}
		result = int32(uintptr(unsafe.Pointer(nox_xxx_getNextPlayerUnit_4DA7F0((*nox_object_t)(unsafe.Pointer(uintptr(i)))))))
	}
	return result
}
func sub_4DCE00() int32 {
	var result int32
	result = int32(dword_5d4594_1563044)
	if dword_5d4594_1563044 != 0 {
		result = sub_419F00()
		if result == 0 {
			result = sub_4DCE30()
			dword_5d4594_1563044 = 0
		}
	}
	return result
}
func sub_4DCE30() int32 {
	if *memmap.PtrUint32(6112660, 1563048) == 0 {
		sub_4169F0()
	}
	return sub_446140()
}
func nox_xxx_sendGauntlet_4DCF80(a1 int32, a2 int8) int32 {
	var v4 [3]byte
	v4[0] = 240
	v4[1] = 28
	v4[2] = byte(a2)
	return nox_xxx_netSendPacket1_4E5390(a1, int32(uintptr(unsafe.Pointer(&v4[0]))), 3, 0, 0)
}
func sub_4DCFB0(a1p *nox_object_t) {
	var (
		a1       int32 = int32(uintptr(unsafe.Pointer(a1p)))
		result   int8
		v2       int32
		v3       int32
		v4       int32
		v5       *byte
		FileName [1024]byte
	)
	result = int8(a1)
	v2 = 1
	if a1 != 0 {
		v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
		v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(v3 + 276))))
		result = int8(*(*uint8)(unsafe.Pointer(uintptr(v4 + 2064))))
		if int32(result) != 31 {
			if *(*uint32)(unsafe.Pointer(uintptr(v4 + 4792))) != 0 && *(*uint32)(unsafe.Pointer(uintptr(v3 + 552))) != 1 {
				if sub_419EE0(int8(*(*uint8)(unsafe.Pointer(uintptr(v4 + 2064))))) != 0 {
					goto LABEL_13
				}
				v5 = nox_fs_root()
				nox_sprintf(&FileName[0], CString("%s\\Save\\_temp_.dat"), v5)
				if nox_xxx_playerSaveToFile_41A140(&FileName[0], int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v3 + 276))) + 2064))))) != 0 {
					v2 = sub_41CFA0(&FileName[0], int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v3 + 276))) + 2064)))))
				}
				nox_fs_remove(&FileName[0])
				if v2 != 0 {
				LABEL_13:
					nox_xxx_player_4D7960(int8(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v3 + 276))) + 2064)))))
				}
			} else {
				nox_xxx_playerCallDisconnect_4DEAB0(int32(*(*uint8)(unsafe.Pointer(uintptr(v4 + 2064)))), 4)
			}
		}
	}
}
func sub_4DD0B0(a1 int32) *byte {
	var (
		result *byte
		v2     int32
	)
	result = (*byte)(unsafe.Pointer(uintptr(a1)))
	if a1 != 0 {
		v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
		if nox_xxx_player_4D7980(int8(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 276))) + 2064))))) != 0 {
			result = nox_xxx_playerCallDisconnect_4DEAB0(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 276))) + 2064)))), 4)
		} else {
			sub_419EB0(int8(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 276))) + 2064)))), 0)
			result = (*byte)(unsafe.Pointer(uintptr(nox_xxx_sendGauntlet_4DCF80(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 276))) + 2064)))), 0))))
		}
	}
	return result
}
func nox_xxx_playerSendMOTD_4DD140(a1 int32) unsafe.Pointer {
	var (
		result unsafe.Pointer
		v3     int32
	)
	result = sub_4464D0(1, (*uint32)(unsafe.Pointer(&v3)))
	if result != nil {
		result = unsafe.Pointer(uintptr(sub_40BC60(a1, 1, CString("MOTD"), int32(uintptr(result)), v3, 1)))
	}
	return result
}
func nox_xxx_gameServerReadyMB_4DD180(a1 int32) *byte {
	var (
		result *byte
		v2     *byte
		v3     int32
		i      int32
		v5     int32
	)
	result = (*byte)(unsafe.Pointer(nox_common_playerInfoFromNum_417090(a1)))
	v2 = result
	if result != nil {
		nox_xxx_netNeedTimestampStatus_4174F0(int32(uintptr(unsafe.Pointer(result))), 16)
		if noxflags.HasGame(noxflags.GameOnline) && !noxflags.HasGame(noxflags.GameModeChat) {
			v3 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), unsafe.Sizeof(uint32(0))*514))))
			if v3 != 0 {
				nox_xxx_spellBuffOff_4FF5B0((*nox_object_t)(unsafe.Pointer(uintptr(v3))), 23)
				nox_xxx_buffApplyTo_4FF380((*nox_object_t)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), unsafe.Sizeof(uint32(0))*514)))))), 23, int16(int32(uint16(nox_gameFPS))*5), 5)
			}
			for i = int32(uintptr(unsafe.Pointer(noxServer.firstServerObject()))); i != 0; i = int32(uintptr(unsafe.Pointer((*nox_object_t)(unsafe.Pointer(uintptr(i))).Next()))) {
				if *(*uint32)(unsafe.Pointer(uintptr(i + 8)))&0x10000000 != 0 {
					nox_xxx_netMarkMinimapObject_417190(a1, i, 1)
				}
			}
		}
		if (*(*byte)(unsafe.Add(unsafe.Pointer(v2), 3680)) & 1) == 1 {
			v5 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), unsafe.Sizeof(uint32(0))*514))))
			*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), unsafe.Sizeof(uint32(0))*908))) = *(*uint32)(unsafe.Pointer(uintptr(v5 + 56)))
			*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), unsafe.Sizeof(uint32(0))*909))) = *(*uint32)(unsafe.Pointer(uintptr(v5 + 60)))
			if noxflags.HasGame(noxflags.GameModeSolo10) {
				nox_xxx_playerLeaveObserver_0_4E6AA0((*nox_playerInfo)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v2)))))))
			}
		}
		nox_xxx_wall_4DF1E0(a1)
		if noxflags.HasGame(noxflags.GameModeQuest) && sub_40A300() == 1 {
			nox_xxx_netGauntlet_4D9E70(a1)
		}
		if a1 == 31 && noxflags.HasGame(noxflags.GameModeChat) {
			if nox_server_connectionType_3596 == 0 && nox_xxx_check_flag_aaa_43AF70() == 1 {
				sub_49C820()
				return (*byte)(unsafe.Pointer(uintptr(nox_xxx_netStatsMultiplier_4D9C20(int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), unsafe.Sizeof(uint32(0))*514))))))))
			}
			if nox_server_sanctuaryHelp_54276 == 1 {
				nox_xxx_cliShowHelpGui_49C560()
				return (*byte)(unsafe.Pointer(uintptr(nox_xxx_netStatsMultiplier_4D9C20(int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), unsafe.Sizeof(uint32(0))*514))))))))
			}
			nox_xxx_guiServerOptsLoad_457500()
		}
		result = (*byte)(unsafe.Pointer(uintptr(nox_xxx_netStatsMultiplier_4D9C20(int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), unsafe.Sizeof(uint32(0))*514))))))))
	}
	return result
}
func nox_xxx____setargv_15_4DD310() {
	*memmap.PtrUint32(6112660, 1563280) = 1
}
func nox_xxx_netGuiGameSettings_4DD9B0(a1 int8, a2 unsafe.Pointer, a3 int32) int32 {
	var v4 [60]byte
	v4[0] = 177
	v4[1] = byte(a1)
	libc.MemCpy(unsafe.Pointer(&v4[2]), a2, 58)
	return nox_xxx_netSendPacket1_4E5390(a3, int32(uintptr(unsafe.Pointer(&v4[0]))), 60, 0, 0)
}
func nox_xxx_playerCheckName_4DDA00(pl *nox_playerInfo) {
	for i := int32(2); ; i++ {
		var ok bool = true
		for pl2 := (*nox_playerInfo)(nox_common_playerInfoGetFirst_416EA0()); pl2 != nil; pl2 = nox_common_playerInfoGetNext_416EE0(pl2) {
			if int32(pl.playerInd) == int32(pl2.playerInd) {
				continue
			}
			if nox_wcscmp(&pl.name_final[0], &pl2.name_final[0]) == 0 {
				ok = false
				break
			}
		}
		if ok {
			return
		}
		nox_swprintf(&pl.info.name_suff[0], libc.CWString(" %d"), i)
		nox_swprintf(&pl.name_final[0], libc.CWString("%s%s"), &pl.info.name[0], &pl.info.name_suff[0])
	}
}
func nox_xxx_netNewPlayerMakePacket_4DDA90(buf *uint8, pl *nox_playerInfo) {
	*(*uint8)(unsafe.Add(unsafe.Pointer(buf), 0)) = 45
	*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(buf), 1)))) = uint16(pl.netCode)
	*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(buf), 100)))) = uint16(pl.field_2136)
	*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(buf), 102)))) = uint16(pl.field_2140)
	*(*uint32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(buf), 104)))) = pl.field_0
	*(*uint32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(buf), 108)))) = pl.field_4
	*(*uint8)(unsafe.Add(unsafe.Pointer(buf), 116)) = uint8(pl.field_2152)
	*(*uint8)(unsafe.Add(unsafe.Pointer(buf), 117)) = uint8(pl.field_2156)
	*(*uint8)(unsafe.Add(unsafe.Pointer(buf), 118)) = uint8(int8(bool2int(int32(pl.field_3676) == 3)))
	*(*uint32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(buf), 112)))) = pl.field_3680 & 1059
	libc.MemCpy(unsafe.Add(unsafe.Pointer(buf), 119), unsafe.Pointer(&pl.field_2096[0]), libc.StrLen(&pl.field_2096[0])+1)
	libc.MemCpy(unsafe.Add(unsafe.Pointer(buf), 3), unsafe.Pointer(&pl.info), 97)
}
func nox_xxx_servSendSettings_4DDB40(punit *nox_object_t) {
	var (
		a1 *uint32 = (*uint32)(unsafe.Pointer(punit))
		v1 int32
		v2 int32
		v3 *byte
		v5 [256]byte
	)
	*(*uint32)(unsafe.Pointer(&v5[1])) = nox_frame_xxx_2598000
	v1 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*187)))
	v5[0] = 40
	nox_netlist_addToMsgListCli_40EBC0(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v1 + 276))) + 2064)))), 1, (*uint8)(unsafe.Pointer(&v5[0])), 5)
	v5[0] = 44
	*(*uint16)(unsafe.Pointer(&v5[1])) = uint16(nox_xxx_netGetUnitCodeServ_578AC0((*nox_object_t)(unsafe.Pointer(a1))))
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 276))))
	*(*uint32)(unsafe.Pointer(&v5[3])) = *(*uint32)(unsafe.Pointer(uintptr(v2 + 2068)))
	nox_netlist_addToMsgListCli_40EBC0(int32(*(*uint8)(unsafe.Pointer(uintptr(v2 + 2064)))), 1, (*uint8)(unsafe.Pointer(&v5[0])), 7)
	sub_4161E0()
	v3 = nox_xxx_cliGamedataGet_416590(0)
	v5[8] = 175
	*(*uint32)(unsafe.Pointer(&v5[9])) = nox_frame_xxx_2598000
	*(*uint32)(unsafe.Pointer(&v5[17])) = nox_common_gameFlags_getVal_40A5B0() & 0x7FFF0
	*(*uint32)(unsafe.Pointer(&v5[21])) = uint32(nox_xxx_getServerSubFlags_409E60())
	*(*uint32)(unsafe.Pointer(&v5[13])) = NOX_CLIENT_VERS_CODE
	v5[25] = byte(int8(nox_xxx_servGetPlrLimit_409FA0()))
	v5[26] = byte(int8(nox_xxx_servGamedataGet_40A020(int16(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v3))), unsafe.Sizeof(uint16(0))*26)))))))
	v5[27] = byte(sub_40A180(int16(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v3))), unsafe.Sizeof(uint16(0))*26))))))
	nox_netlist_addToMsgListCli_40EBC0(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v1 + 276))) + 2064)))), 1, (*uint8)(unsafe.Pointer(&v5[8])), 20)
	v5[28] = 176
	libc.StrCpy(&v5[29], nox_xxx_serverOptionsGetServername_40A4C0())
	libc.MemCpy(unsafe.Pointer(&v5[45]), unsafe.Add(unsafe.Pointer(v3), 24), 28)
	if sub_40A220() != 0 && (sub_40A300() != 0 || int32(sub_40A180(int16(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v3))), unsafe.Sizeof(uint16(0))*26)))))) != 0) {
		*(*uint32)(unsafe.Pointer(&v5[73])) = uint32(sub_40A230())
	} else {
		*(*uint32)(unsafe.Pointer(&v5[73])) = 0
	}
	nox_netlist_addToMsgListCli_40EBC0(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v1 + 276))) + 2064)))), 1, (*uint8)(unsafe.Pointer(&v5[28])), 49)
	nox_xxx_netNewPlayerMakePacket_4DDA90((*uint8)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(&v5[124])))))), (*nox_playerInfo)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v1 + 276)))))))
	nox_netlist_addToMsgListCli_40EBC0(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v1 + 276))) + 2064)))), 1, (*uint8)(unsafe.Pointer(&v5[124])), 129)
	nox_xxx_netSendBySock_4DDDC0(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v1 + 276))) + 2064)))))
	v5[80] = 43
	libc.StrCpy(&v5[81], nox_server_currentMapGetFilename_409B30())
	*(*uint32)(unsafe.Pointer(&v5[113])) = uint32(nox_xxx_mapCrcGetMB_409B00())
	*(*uint32)(unsafe.Pointer(&v5[117])) = nox_frame_xxx_2598000
	nox_xxx_netSendSock_552640(uint32(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v1 + 276))) + 2064))))+1), &v5[80], 41, int8(NOX_NET_SEND_NO_LOCK|NOX_NET_SEND_FLAG2))
	sub_4DDE10(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v1 + 276))) + 2064)))), int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 276)))))
}
func nox_xxx_netSendBySock_4DDDC0(a1 int32) {
	var (
		v2 int32
		v3 *uint8
	)
	v2 = a1
	if !noxflags.HasGame(noxflags.GameClient) || a1 != 31 {
		v3 = nox_netlist_copyPacketList_40ED60(a1, 1, (*uint32)(unsafe.Pointer(&a1)))
		if v3 != nil {
			nox_xxx_netSendSock_552640(uint32(v2+1), (*byte)(unsafe.Pointer(v3)), a1, int8(NOX_NET_SEND_NO_LOCK|NOX_NET_SEND_FLAG2))
		}
		nox_netlist_resetByInd_40ED10(v2, 1)
	}
}
func sub_4DDE10(a1 int32, a2 int32) *uint32 {
	var (
		result *uint32
		i      int32
		v4     int32
	)
	result = *(**uint32)(unsafe.Pointer(uintptr(a2 + 2056)))
	if result != nil {
		if dword_5d4594_1563276 == 0 {
			dword_5d4594_1563276 = uint32(nox_xxx_getNameId_4E3AA0(CString("Flag")))
		}
		result = *(**uint32)(unsafe.Pointer(uintptr(a2 + 2056)))
		for i = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*126))); i != 0; i = int32(*(*uint32)(unsafe.Pointer(uintptr(i + 496)))) {
			v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(i + 16))))
			if (v4 & 256) == 0 {
				result = *(**uint32)(unsafe.Pointer(&dword_5d4594_1563276))
				if uint32(*(*uint16)(unsafe.Pointer(uintptr(i + 4)))) != dword_5d4594_1563276 {
					continue
				}
			}
			result = sub_4D82F0(a1, (*uint32)(unsafe.Pointer(uintptr(i))))
		}
	}
	return result
}
func nox_xxx_playerObserveMonster_4DDE80(player int32, unit int32) int32 {
	var (
		v2     int32
		result int32
	)
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(player + 748))))
	if int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 276))) + 3680))))&1 != 0 {
		nox_xxx_playerLeaveObserver_0_4E6AA0((*nox_playerInfo)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 276)))))))
	}
	if int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 276))) + 3680))))&2 != 0 {
		nox_xxx_playerObserveClear_4DDEF0(player)
	}
	nox_xxx_netNeedTimestampStatus_4174F0(int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 276)))), 2)
	nox_xxx_playerCameraUnlock_4E6040(player)
	result = nox_xxx_playerCameraFollow_4E6060(player, unit)
	*(*uint32)(unsafe.Pointer(uintptr(player + 744))) = uint32(cgoFuncAddr(nox_xxx_updatePlayerObserver_4E62F0))
	return result
}
func nox_xxx_playerObserveClear_4DDEF0(player int32) int32 {
	var result int32
	result = int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(player + 748))) + 276))))
	if int32(*(*uint8)(unsafe.Pointer(uintptr(result + 3680))))&2 != 0 {
		nox_xxx_playerUnsetStatus_417530((*nox_playerInfo)(unsafe.Pointer(uintptr(result))), 2)
		result = nox_xxx_playerCameraUnlock_4E6040(player)
		*(*uint32)(unsafe.Pointer(uintptr(player + 744))) = uint32(cgoFuncAddr(nox_xxx_updatePlayer_4F8100))
	}
	return result
}
func nox_xxx_playerGetPossess_4DDF30(a1 int32) int32 {
	var (
		v1     int32
		result int32
	)
	v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))) + 276))))
	if int32(*(*uint8)(unsafe.Pointer(uintptr(v1 + 3680))))&2 != 0 {
		result = int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 3628))))
	} else {
		result = 0
	}
	return result
}
func nox_xxx_netPlayerIncomingServ_4DDF60(a1 int32) int32 {
	var (
		v1  int32
		v2  *byte
		v3  int32
		v4  int32
		v5  int32
		i   *byte
		v7  int32
		v8  *byte
		j   *byte
		v10 *uint8
		k   int32
		v13 int32
	)
	v1 = a1
	v2 = (*byte)(unsafe.Pointer(nox_common_playerInfoFromNum_417090(a1)))
	if v2 == nil {
		panic("abort")
	}
	if noxflags.HasGame(noxflags.GameModeQuest) {
		if a1 != 31 {
			if v2 != nil {
				v3 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), unsafe.Sizeof(uint32(0))*514))))
				if v3 != 0 {
					*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v3 + 748))) + 552))) = 1
				}
			}
		}
		sub_4D9CF0(a1)
		if v2 != nil && *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), unsafe.Sizeof(uint32(0))*514))) != 0 {
			sub_4D6000((*nox_object_t)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), unsafe.Sizeof(uint32(0))*514)))))))
		}
	}
	sub_57B920(unsafe.Add(unsafe.Pointer(v2), 16))
	v13 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), unsafe.Sizeof(uint32(0))*514))))
	dword_5d4594_2649712 |= uint32(1 << v1)
	v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(v13 + 56))))
	v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(v13 + 60))))
	nox_xxx_newPlayerSendAllPlayers_4DE300(v1)
	*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), unsafe.Sizeof(uint32(0))*1175))) = 0
	(cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(v13 + 688))), (*func(int32, uint32))(nil)))(v13, 0)
	*(*byte)(unsafe.Add(unsafe.Pointer(v2), 3676)) = 3
	if !noxflags.HasGame(noxflags.GameModeSolo10) {
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), unsafe.Sizeof(uint32(0))*908))) = uint32(v4)
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), unsafe.Sizeof(uint32(0))*909))) = uint32(v5)
	}
	if nox_server_sendMotd_108752 != 0 && noxflags.HasGame(noxflags.GameOnline) && !noxflags.HasGame(noxflags.GameModeQuest) {
		nox_xxx_playerSendMOTD_4DD140(v1)
	}
	for i = (*byte)(unsafe.Pointer(nox_common_playerInfoGetFirst_416EA0())); i != nil; i = (*byte)(unsafe.Pointer(nox_common_playerInfoGetNext_416EE0((*nox_playerInfo)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(i))))))))) {
		v7 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(i))), unsafe.Sizeof(uint32(0))*514))))
		if v7 != 0 {
			if i != v2 {
				nox_xxx_netMarkMinimapObject_417190(v1, v7, 1)
				nox_xxx_netMarkMinimapObject_417190(int32(uint8(*(*byte)(unsafe.Add(unsafe.Pointer(i), 2064)))), int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), unsafe.Sizeof(uint32(0))*514)))), 1)
				nox_xxx_netSendSimpleObject2_4DF360(int32(uint8(*(*byte)(unsafe.Add(unsafe.Pointer(i), 2064)))), int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), unsafe.Sizeof(uint32(0))*514)))))
				if noxflags.HasGame(noxflags.GameModeQuest) {
					nox_xxx_netSendTeam_4D8670(int32(uint8(*(*byte)(unsafe.Add(unsafe.Pointer(i), 2064)))), *((**uint32)(unsafe.Add(unsafe.Pointer((**uint32)(unsafe.Pointer(v2))), unsafe.Sizeof((*uint32)(nil))*514))))
					nox_xxx_netSendTeam_4D8670(v1, *((**uint32)(unsafe.Add(unsafe.Pointer((**uint32)(unsafe.Pointer(i))), unsafe.Sizeof((*uint32)(nil))*514))))
				}
			}
		}
	}
	nox_xxx_servMinimapRevealFlag_4DE380(v1)
	sub_4DF2E0(v1)
	if noxflags.HasGame(noxflags.GameModeElimination) && sub_40AA70((*nox_playerInfo)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v2))))))) == 0 {
		nox_xxx_netNeedTimestampStatus_4174F0(int32(uintptr(unsafe.Pointer(v2))), 256)
	}
	if nox_xxx_check_flag_aaa_43AF70() == 1 {
		sub_4161E0()
		sub_416690()
	}
	sub_4E8110(v1)
	if noxflags.HasGame(noxflags.GameModeFlagBall) {
		v8 = sub_4E8310()
		nox_xxx_netSendBallStatus_4D95F0(v1, int8(*v8), int16(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v8))), unsafe.Sizeof(uint16(0))*1)))))
	} else if noxflags.HasGame(noxflags.GameModeCTF) {
		for j = (*byte)(unsafe.Pointer(nox_server_teamFirst_418B10())); j != nil; j = (*byte)(unsafe.Pointer(nox_server_teamNext_418B60((*nox_team_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(j))))))))) {
			v10 = sub_4E8320(uint8(*(*byte)(unsafe.Add(unsafe.Pointer(j), 57))))
			nox_xxx_netSendFlagStatus_4D95A0(v1, int8(*v10), int8(*(*uint8)(unsafe.Add(unsafe.Pointer(v10), 2))), int8(*(*uint8)(unsafe.Add(unsafe.Pointer(v10), 1))), int16(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v10))), unsafe.Sizeof(uint16(0))*2)))))
		}
	}
	nox_xxx_sendAllClientStatus_4175C0(int32(uintptr(unsafe.Pointer(v2))))
	if sub_409F40(8192) != 0 {
		nox_xxx_sendAllPlayerIDs_4DE270(int32(uintptr(unsafe.Pointer(v2))))
	}
	if noxflags.HasGame(noxflags.GameModeQuest) {
		for k = int32(uintptr(unsafe.Pointer(nox_xxx_getFirstPlayerUnit_4DA7C0()))); k != 0; k = int32(uintptr(unsafe.Pointer(nox_xxx_getNextPlayerUnit_4DA7F0((*nox_object_t)(unsafe.Pointer(uintptr(k))))))) {
			if *(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(k + 748))) + 276))) + 4792))) == 1 {
				sub_4D9D20(v1, (*nox_object_t)(unsafe.Pointer(uintptr(k))))
			}
		}
	}
	if noxflags.HasGame(noxflags.GameModeQuest) {
		sub_4D7A60(v1)
	}
	return int32(*(*uint32)(unsafe.Pointer(uintptr(v13 + 36))))
}
func nox_xxx_sendAllPlayerIDs_4DE270(a1 int32) int32 {
	var (
		result int32
		i      int32
		v3     int32
		v4     [7]byte
	)
	result = int32(uintptr(unsafe.Pointer(nox_xxx_getFirstPlayerUnit_4DA7C0())))
	for i = result; result != 0; i = result {
		if *(*uint32)(unsafe.Pointer(uintptr(i + 36))) != *(*uint32)(unsafe.Pointer(uintptr(a1 + 2060))) {
			if *(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(i + 748))) + 260))) != 0 {
				v4[0] = 210
				*(*uint16)(unsafe.Pointer(&v4[1])) = uint16(nox_xxx_netGetUnitCodeServ_578AC0((*nox_object_t)(unsafe.Pointer(uintptr(i)))))
				v3 = int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 2064))))
				*(*uint16)(unsafe.Pointer(&v4[3])) = *(*uint16)(unsafe.Pointer(uintptr(i + 4)))
				v4[5] = 1
				v4[6] = 2
				nox_xxx_netSendPacket0_4E5420(v3, unsafe.Pointer(&v4[0]), 7, 0, 1)
			}
		}
		result = int32(uintptr(unsafe.Pointer(nox_xxx_getNextPlayerUnit_4DA7F0((*nox_object_t)(unsafe.Pointer(uintptr(i)))))))
	}
	return result
}
func nox_xxx_newPlayerSendAllPlayers_4DE300(a1 int32) *byte {
	var (
		result *byte
		i      int32
		v3     [132]byte
	)
	result = (*byte)(unsafe.Pointer(nox_common_playerInfoGetFirst_416EA0()))
	for i = int32(uintptr(unsafe.Pointer(result))); result != nil; i = int32(uintptr(unsafe.Pointer(result))) {
		if int32(*(*uint8)(unsafe.Pointer(uintptr(i + 2064)))) != a1 && (int32(*(*uint8)(unsafe.Pointer(uintptr(i + 2064)))) != 31 || !nox_common_getEngineFlag(nox_engine_flag(NOX_ENGINE_FLAG_DISABLE_GRAPHICS_RENDERING))) {
			nox_xxx_netNewPlayerMakePacket_4DDA90((*uint8)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(&v3[0])))))), (*nox_playerInfo)(unsafe.Pointer(uintptr(i))))
			nox_xxx_netSendPacket1_4E5390(a1, int32(uintptr(unsafe.Pointer(&v3[0]))), 129, 0, 0)
			sub_4DDE10(a1, i)
		}
		result = (*byte)(unsafe.Pointer(nox_common_playerInfoGetNext_416EE0((*nox_playerInfo)(unsafe.Pointer(uintptr(i))))))
	}
	return result
}
func nox_xxx_servMinimapRevealFlag_4DE380(a1 int32) int32 {
	var (
		result int32
		i      int32
		v3     int32
	)
	if *memmap.PtrUint32(6112660, 1563284) == 0 {
		*memmap.PtrUint32(6112660, 1563284) = uint32(nox_xxx_getNameId_4E3AA0(CString("GameBall")))
	}
	if *memmap.PtrUint32(6112660, 1563272) == 0 {
		*memmap.PtrUint32(6112660, 1563272) = uint32(nox_xxx_getNameId_4E3AA0(CString("Crown")))
	}
	result = int32(uintptr(unsafe.Pointer(noxServer.firstServerObject())))
	for i = result; result != 0; i = result {
		if int32(*(*uint8)(unsafe.Pointer(uintptr(i + 16))))&4 != 0 {
			v3 = int32(*(*uint16)(unsafe.Pointer(uintptr(i + 4))))
			if uint32(uint16(int16(v3))) == *memmap.PtrUint32(6112660, 1563284) || uint32(v3) == *memmap.PtrUint32(6112660, 1563272) {
				nox_xxx_netMarkMinimapObject_417190(a1, i, 1)
			}
		}
		result = int32(uintptr(unsafe.Pointer((*nox_object_t)(unsafe.Pointer(uintptr(i))).Next())))
	}
	return result
}
func sub_4DE410(a1 int32) *uint32 {
	var (
		v1     int32
		v2     int32
		result *uint32
		i      *uint32
		v5     int32
		v6     int32
		v7     uint32
		v8     int32
	)
	v1 = a1
	v2 = 1 << a1
	v8 = 1 << a1
	result = (*uint32)(unsafe.Pointer(noxServer.firstServerObject()))
	for i = result; result != nil; i = result {
		v5 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(i), unsafe.Sizeof(uint32(0))*2)))
		*(*uint32)(unsafe.Add(unsafe.Pointer(i), unsafe.Sizeof(uint32(0))*38)) |= uint32(v2)
		if (uint32(v5) & 0x20400000) == 0 {
			*(*uint32)(unsafe.Add(unsafe.Pointer(i), unsafe.Sizeof(uint32(0))*37)) &= uint32(^v2)
		}
		*(*uint32)(unsafe.Add(unsafe.Pointer(i), unsafe.Sizeof(uint32(0))*uintptr(v1+140))) &= 4095
		if *(*uint32)(unsafe.Add(unsafe.Pointer(i), unsafe.Sizeof(uint32(0))*2))&0x20400000 != 0 {
			v6 = 0x10000
			v7 = 1
			for {
				if sub_4E4C90(int32(uintptr(unsafe.Pointer(i))), v7) != 0 {
					*(*uint32)(unsafe.Add(unsafe.Pointer(i), unsafe.Sizeof(uint32(0))*uintptr(v1+140))) |= uint32(v6)
				}
				v7 *= 2
				v6 *= 2
				if v7 >= 0x10000 {
					break
				}
			}
			v2 = v8
		}
		result = (*uint32)(unsafe.Pointer((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(i)))))).Next()))
	}
	return result
}
func sub_4DE4D0(a1 int8) int32 {
	var (
		v1     int32
		result int32
		v3     int8
	)
	v1 = 1 << int32(a1)
	for result = int32(uintptr(unsafe.Pointer(noxServer.firstServerObject()))); result != 0; result = int32(uintptr(unsafe.Pointer((*nox_object_t)(unsafe.Pointer(uintptr(result))).Next()))) {
		v3 = int8(*(*uint8)(unsafe.Pointer(uintptr(result + 16))))
		*(*uint32)(unsafe.Pointer(uintptr(result + 152))) |= uint32(v1)
		if (int32(v3)&32) == 0 && (*(*uint32)(unsafe.Pointer(uintptr(result + 8)))&0x20400006) == 0 {
			*(*uint32)(unsafe.Pointer(uintptr(result + 148))) &= uint32(^v1)
		}
	}
	return result
}
func nox_xxx_playerDisconnFinish_4DE530(a1 int32, a2 int8) int8 {
	var (
		v2     int32
		v3     *byte
		v4     int32
		v5     *byte
		i      *byte
		v7     int32
		j      int32
		result int8
		v11    int32
		v13    int32
	)
	v2 = a1
	v13 = 1 << a1
	sub_4D79A0(int8(a1))
	sub_419EB0(int8(v2), 0)
	sub_4E80C0(int8(v2))
	v3 = (*byte)(unsafe.Pointer(nox_common_playerInfoFromNum_417090(v2)))
	v4 = int32(uintptr(unsafe.Pointer(v3)))
	if v3 != nil {
		if *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v3))), unsafe.Sizeof(uint32(0))*514))) != 0 {
			sub_506740(int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v3))), unsafe.Sizeof(uint32(0))*514)))))
		}
		if *(*uint32)(unsafe.Pointer(uintptr(v4 + 4792))) != 0 && *(*uint32)(unsafe.Pointer(uintptr(v4 + 2056))) != 0 && noxflags.HasGame(noxflags.GameModeQuest) {
			nox_xxx_aud_501960(1011, (*nox_object_t)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v4 + 2056)))))), 0, 0)
			*(*uint32)(unsafe.Pointer(uintptr(v4 + 4792))) = 0
			if nox_xxx_player_4E3CE0() == 0 {
				v5 = nox_xxx_getQuestMapFile_4D0F60()
				nox_game_setQuestStage_4E3CD0(0)
				nox_xxx_mapLoad_4D2450(v5)
			}
			sub_4D7390(int32(*(*uint32)(unsafe.Pointer(uintptr(v4 + 2056)))))
			sub_4DE790(int32(*(*uint32)(unsafe.Pointer(uintptr(v4 + 2056)))))
		}
	}
	if nox_xxx_check_flag_aaa_43AF70() == 1 && !noxflags.HasGame(noxflags.GameModeChat) {
		sub_425E90(v4, a2)
	}
	for i = (*byte)(unsafe.Pointer(nox_common_playerInfoGetFirst_416EA0())); i != nil; i = (*byte)(unsafe.Pointer(nox_common_playerInfoGetNext_416EE0((*nox_playerInfo)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(i))))))))) {
		v7 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(i))), unsafe.Sizeof(uint32(0))*514))))
		if v7 != 0 && i != (*byte)(unsafe.Pointer(uintptr(v4))) {
			nox_xxx_netUnmarkMinimapObj_417300(v2, v7, 3)
			nox_xxx_netUnmarkMinimapObj_417300(int32(uint8(*(*byte)(unsafe.Add(unsafe.Pointer(i), 2064)))), int32(*(*uint32)(unsafe.Pointer(uintptr(v4 + 2056)))), 3)
		}
	}
	if *(*uint32)(unsafe.Pointer(uintptr(v4 + 2056))) != 0 {
		for j = int32(uintptr(unsafe.Pointer(noxServer.firstServerObject()))); j != 0; j = int32(uintptr(unsafe.Pointer((*nox_object_t)(unsafe.Pointer(uintptr(j))).Next()))) {
			if (int32(*(*uint8)(unsafe.Pointer(uintptr(j + 8))))&128) != 0 && *(*uint32)(unsafe.Pointer(uintptr(j + 508))) == *(*uint32)(unsafe.Pointer(uintptr(v4 + 2056))) {
				*(*uint32)(unsafe.Pointer(uintptr(j + 508))) = 0
			}
		}
	}
	sub_40C0E0(v2 + 1)
	sub_4DE410(v2)
	if v4 != 0 {
		var v12 [3]byte = [3]byte{}
		v12[0] = 174
		if *(*uint32)(unsafe.Pointer(uintptr(v4 + 2056))) != 0 {
			*(*uint16)(unsafe.Pointer(&v12[1])) = *(*uint16)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v4 + 2056))) + 36)))
		}
		nox_xxx_netSendPacket0_4E5420(v2|128, unsafe.Pointer(&v12[0]), 3, 0, 0)
	}
	if int32(a2) == 4 {
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a1))), 0)) = 197
		nox_xxx_netSendSock_552640(uint32(v2+1), (*byte)(unsafe.Pointer(&a1)), 1, int8(NOX_NET_SEND_NO_LOCK|NOX_NET_SEND_FLAG2))
	}
	result = int8(sub_4E55F0(uint8(int8(v2))))
	if v4 != 0 {
		if *(*uint32)(unsafe.Pointer(uintptr(v4 + 2056))) != 0 {
			nox_xxx_playerRemoveSpawnedStuff_4E5AD0(int32(*(*uint32)(unsafe.Pointer(uintptr(v4 + 2056)))))
			if noxflags.HasGame(noxflags.GameModeQuest) {
				sub_4DE790(int32(*(*uint32)(unsafe.Pointer(uintptr(v4 + 2056)))))
			} else {
				nox_xxx_dropAllItems_4EDA40(*(**uint32)(unsafe.Pointer(uintptr(v4 + 2056))))
			}
		}
		v11 = v13
		*(*uint32)(unsafe.Pointer(uintptr(v4 + 2140))) = 0
		*(*uint32)(unsafe.Pointer(uintptr(v4 + 2136))) = 0
		dword_5d4594_2649712 &= uint32(^v11)
		*(*uint8)(unsafe.Pointer(uintptr(v4 + 3676))) = 2
		result = nox_xxx_playerUnsetStatus_417530((*nox_playerInfo)(unsafe.Pointer(uintptr(v4))), 16)
	}
	return result
}
func sub_4DE790(a1 int32) {
	var (
		v1 int32
		v2 int32
	)
	v1 = a1.FirstItem()
	if v1 != 0 {
		for {
			v2 = nox_xxx_inventoryGetNext_4E7990(v1)
			nox_xxx_delayedDeleteObject_4E5CC0((*nox_object_t)(unsafe.Pointer(uintptr(v1))))
			v1 = v2
			if v2 == 0 {
				break
			}
		}
	}
}
func nox_xxx_playerForceDisconnect_4DE7C0(ind int32) {
	var plr *nox_playerInfo = nox_common_playerInfoFromNum_417090(ind)
	nox_script_callByEvent_cgo(31, unsafe.Pointer(plr), nil)
	if sub_4D12A0(ind) != 0 {
		sub_4D1250(ind)
	}
	if plr.field_2068 != 0 {
		var v3 *int32 = sub_425A70(int32(plr.field_2068))
		if v3 != nil {
			sub_425B60(unsafe.Pointer(v3), ind)
		}
	}
	var v4 int32 = int32(*(*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer(plr.playerUnit))) + 748))))
	if *(*uint32)(unsafe.Pointer(uintptr(v4 + 280))) != 0 {
		nox_xxx_shopCancelSession_510DC0(*(**uint32)(unsafe.Pointer(uintptr(v4 + 280))))
	}
	*(*uint32)(unsafe.Pointer(uintptr(v4 + 280))) = 0
	sub_510E20(int32(plr.playerInd))
	sub_4FF990(1 << int32(plr.playerInd))
	if !noxflags.HasGame(noxflags.GameClient) {
		plr.active = 0
	}
	var pl *byte = (*byte)(unsafe.Pointer(plr))
	sub_56F4F0((*int32)(unsafe.Add(unsafe.Pointer((*int32)(unsafe.Pointer(pl))), unsafe.Sizeof(int32(0))*1146)))
	sub_56F4F0((*int32)(unsafe.Add(unsafe.Pointer((*int32)(unsafe.Pointer(pl))), unsafe.Sizeof(int32(0))*1148)))
	sub_56F4F0((*int32)(unsafe.Add(unsafe.Pointer((*int32)(unsafe.Pointer(pl))), unsafe.Sizeof(int32(0))*1149)))
	sub_56F4F0((*int32)(unsafe.Add(unsafe.Pointer((*int32)(unsafe.Pointer(pl))), unsafe.Sizeof(int32(0))*1150)))
	sub_56F4F0((*int32)(unsafe.Add(unsafe.Pointer((*int32)(unsafe.Pointer(pl))), unsafe.Sizeof(int32(0))*1151)))
	sub_56F4F0((*int32)(unsafe.Add(unsafe.Pointer((*int32)(unsafe.Pointer(pl))), unsafe.Sizeof(int32(0))*1152)))
	sub_56F4F0((*int32)(unsafe.Add(unsafe.Pointer((*int32)(unsafe.Pointer(pl))), unsafe.Sizeof(int32(0))*1153)))
	sub_56F4F0((*int32)(unsafe.Add(unsafe.Pointer((*int32)(unsafe.Pointer(pl))), unsafe.Sizeof(int32(0))*1154)))
	sub_56F4F0((*int32)(unsafe.Add(unsafe.Pointer((*int32)(unsafe.Pointer(pl))), unsafe.Sizeof(int32(0))*1155)))
	sub_56F4F0((*int32)(unsafe.Add(unsafe.Pointer((*int32)(unsafe.Pointer(pl))), unsafe.Sizeof(int32(0))*1156)))
	sub_56F4F0((*int32)(unsafe.Add(unsafe.Pointer((*int32)(unsafe.Pointer(pl))), unsafe.Sizeof(int32(0))*1157)))
	sub_56F4F0((*int32)(unsafe.Add(unsafe.Pointer((*int32)(unsafe.Pointer(pl))), unsafe.Sizeof(int32(0))*1158)))
	sub_56F4F0((*int32)(unsafe.Add(unsafe.Pointer((*int32)(unsafe.Pointer(pl))), unsafe.Sizeof(int32(0))*1159)))
	sub_56F4F0((*int32)(unsafe.Add(unsafe.Pointer((*int32)(unsafe.Pointer(pl))), unsafe.Sizeof(int32(0))*1147)))
	sub_56F4F0((*int32)(unsafe.Add(unsafe.Pointer((*int32)(unsafe.Pointer(pl))), unsafe.Sizeof(int32(0))*1160)))
	sub_56F4F0((*int32)(unsafe.Add(unsafe.Pointer((*int32)(unsafe.Pointer(pl))), unsafe.Sizeof(int32(0))*1161)))
	var buf [3]byte
	buf[0] = 46
	*(*uint16)(unsafe.Pointer(&buf[1])) = uint16(nox_xxx_netGetUnitCodeServ_578AC0(plr.playerUnit))
	nox_xxx_netSendPacket0_4E5420(ind|128, unsafe.Pointer(&buf[0]), 3, 0, 0)
	nox_xxx_delayedDeleteObject_4E5CC0(plr.playerUnit)
	plr.playerUnit = nil
	for i := int32(int32(uintptr(unsafe.Pointer(nox_xxx_getFirstPlayerUnit_4DA7C0())))); i != 0; i = int32(uintptr(unsafe.Pointer(nox_xxx_getNextPlayerUnit_4DA7F0((*nox_object_t)(unsafe.Pointer(uintptr(i))))))) {
		var v7 int32 = int32(*(*uint32)(unsafe.Pointer(uintptr(i + 748))))
		*(*uint8)(unsafe.Pointer(uintptr(ind + v7 + 452))) = 0
		*(*uint32)(unsafe.Pointer(uintptr(v7 + ind*4 + 324))) = 0
		*(*uint8)(unsafe.Pointer(uintptr(ind + v7 + 484))) = 0
		*(*uint8)(unsafe.Pointer(uintptr(ind + v7 + 516))) = 0
	}
	if nox_xxx_gamePlayIsAnyPlayers_40A8A0() != 0 {
		if noxflags.HasGame(noxflags.GameModeElimination) && nox_xxx_serverIsClosing_446180() == 0 && sub_40A770() == 1 {
			sub_5095E0()
		}
	} else {
		sub_40A1F0(0)
		nox_xxx_playerForceSendLessons_416E50(1)
		nox_server_teamsResetYyy_417D00()
		sub_40A970()
	}
	sub_4E55F0(uint8(int8(ind)))
	nox_xxx_playerResetImportantCtr_4E4F40(ind)
	sub_4E4F30(ind)
	if nox_xxx_check_flag_aaa_43AF70() == 1 {
		sub_4161E0()
		sub_416690()
	}
	if !noxflags.HasGame(noxflags.GameModeQuest) {
		return
	}
	if sub_4E9010() == 1 {
		nox_xxx_mapLoad_4D2450((*byte)(unsafe.Pointer(sub_4E8E50())))
	} else if nox_server_questMaybeWarp_4E8F60() {
		sub_4D60B0()
		sub_4D76E0(1)
		var v12 int32 = nox_game_getQuestStage_4E3CC0()
		var v13 int32 = nox_server_questNextStageThreshold_4D74F0(v12)
		nox_game_setQuestStage_4E3CD0(v13 - 1)
		nox_xxx_mapLoad_4D2450((*byte)(unsafe.Pointer(sub_4E8E50())))
	} else {
		var result *byte = (*byte)(unsafe.Pointer(nox_xxx_getFirstPlayerUnit_4DA7C0()))
		if result != nil {
			for *(*uint32)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(result))), unsafe.Sizeof(uint32(0))*187))) + 312))) == 0 {
				result = (*byte)(unsafe.Pointer(nox_xxx_getNextPlayerUnit_4DA7F0((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(result)))))))))
				if result == nil {
					return
				}
			}
			result = (*byte)(unsafe.Pointer(uintptr(sub_4E8E60())))
		}
	}
}
func nox_xxx_netGameSettings_4DEF00() int32 {
	var (
		v0 *byte
		v2 [20]byte
		v3 [49]byte
	)
	v0 = nox_xxx_cliGamedataGet_416590(0)
	v2[0] = 175
	*(*uint32)(unsafe.Pointer(&v2[1])) = nox_frame_xxx_2598000
	*(*uint32)(unsafe.Pointer(&v2[9])) = nox_common_gameFlags_getVal_40A5B0() & 0x7FFF0
	*(*uint32)(unsafe.Pointer(&v2[13])) = uint32(nox_xxx_getServerSubFlags_409E60())
	*(*uint32)(unsafe.Pointer(&v2[5])) = NOX_CLIENT_VERS_CODE
	v2[17] = byte(int8(nox_xxx_servGetPlrLimit_409FA0()))
	v2[18] = byte(int8(nox_xxx_servGamedataGet_40A020(int16(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v0))), unsafe.Sizeof(uint16(0))*26)))))))
	v2[19] = byte(sub_40A180(int16(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v0))), unsafe.Sizeof(uint16(0))*26))))))
	v3[0] = 176
	libc.StrCpy(&v3[1], nox_xxx_serverOptionsGetServername_40A4C0())
	libc.MemCpy(unsafe.Pointer(&v3[17]), unsafe.Add(unsafe.Pointer(v0), 24), 28)
	if sub_40A220() != 0 && (sub_40A300() != 0 || int32(sub_40A180(int16(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v0))), unsafe.Sizeof(uint16(0))*26)))))) != 0) {
		*(*uint32)(unsafe.Pointer(&v3[45])) = uint32(sub_40A230())
	} else {
		*(*uint32)(unsafe.Pointer(&v3[45])) = 0
	}
	nox_xxx_netSendPacket1_4E5390(159, int32(uintptr(unsafe.Pointer(&v2[0]))), 20, 0, 0)
	return nox_xxx_netSendPacket1_4E5390(159, int32(uintptr(unsafe.Pointer(&v3[0]))), 49, 0, 0)
}
func sub_4DF020() *byte {
	var (
		result *byte
		v1     int32
		v2     *uint8
		v3     *byte
		v4     bool
		i      *byte
		v6     [60]byte
	)
	result = sub_459AA0(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(&v6[0]))))))
	v1 = 29
	v2 = (*uint8)(memmap.PtrOff(6112660, 1563214))
	v3 = &v6[0]
	v4 = true
	for {
		if v1 == 0 {
			break
		}
		v4 = int32(*(*uint16)(unsafe.Pointer(v3))) == int32(*(*uint16)(unsafe.Pointer(v2)))
		v3 = (*byte)(unsafe.Add(unsafe.Pointer(v3), 2))
		v2 = (*uint8)(unsafe.Add(unsafe.Pointer(v2), 2))
		v1--
		if !v4 {
			break
		}
	}
	if !v4 {
		result = (*byte)(unsafe.Pointer(nox_common_playerInfoGetFirst_416EA0()))
		for i = result; result != nil; i = result {
			if *(*byte)(unsafe.Add(unsafe.Pointer(i), 2064)) != 31 {
				nox_xxx_netGuiGameSettings_4DD9B0(1, unsafe.Pointer(&v6[0]), int32(uint8(*(*byte)(unsafe.Add(unsafe.Pointer(i), 2064)))))
			}
			result = (*byte)(unsafe.Pointer(nox_common_playerInfoGetNext_416EE0((*nox_playerInfo)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(i)))))))))
		}
		libc.MemCpy(memmap.PtrOff(6112660, 1563214), unsafe.Pointer(&v6[0]), 56)
		*memmap.PtrUint16(6112660, 1563270) = *(*uint16)(unsafe.Pointer(&v6[56]))
	}
	return result
}
func nox_xxx_wallSendDestroyed_4DF0A0(a1 int32, a2 int32) int32 {
	var (
		result int32
		i      int32
		v4     [3]byte
	)
	v4[0] = 58
	*(*uint16)(unsafe.Pointer(&v4[1])) = *(*uint16)(unsafe.Pointer(uintptr(a1 + 10)))
	if a2 != 32 {
		return nox_xxx_netSendPacket0_4E5420(a2, unsafe.Pointer(&v4[0]), 3, 0, 1)
	}
	result = int32(uintptr(unsafe.Pointer(nox_xxx_getFirstPlayerUnit_4DA7C0())))
	for i = result; result != 0; i = result {
		nox_xxx_netSendPacket0_4E5420(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(i + 748))) + 276))) + 2064)))), unsafe.Pointer(&v4[0]), 3, 0, 1)
		result = int32(uintptr(unsafe.Pointer(nox_xxx_getNextPlayerUnit_4DA7F0((*nox_object_t)(unsafe.Pointer(uintptr(i)))))))
	}
	return result
}
func sub_4DF120(a1 int32) int32 {
	var (
		result int32
		i      int32
		v3     [3]byte
	)
	v3[0] = 59
	*(*uint16)(unsafe.Pointer(&v3[1])) = *(*uint16)(unsafe.Pointer(uintptr(a1 + 10)))
	result = int32(uintptr(unsafe.Pointer(nox_xxx_getFirstPlayerUnit_4DA7C0())))
	for i = result; result != 0; i = result {
		nox_xxx_netSendPacket0_4E5420(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(i + 748))) + 276))) + 2064)))), unsafe.Pointer(&v3[0]), 3, 0, 1)
		result = int32(uintptr(unsafe.Pointer(nox_xxx_getNextPlayerUnit_4DA7F0((*nox_object_t)(unsafe.Pointer(uintptr(i)))))))
	}
	return result
}
func sub_4DF180(a1 int32) int32 {
	var (
		result int32
		i      int32
		v3     [3]byte
	)
	v3[0] = 60
	*(*uint16)(unsafe.Pointer(&v3[1])) = *(*uint16)(unsafe.Pointer(uintptr(a1 + 10)))
	result = int32(uintptr(unsafe.Pointer(nox_xxx_getFirstPlayerUnit_4DA7C0())))
	for i = result; result != 0; i = result {
		nox_xxx_netSendPacket0_4E5420(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(i + 748))) + 276))) + 2064)))), unsafe.Pointer(&v3[0]), 3, 0, 1)
		result = int32(uintptr(unsafe.Pointer(nox_xxx_getNextPlayerUnit_4DA7F0((*nox_object_t)(unsafe.Pointer(uintptr(i)))))))
	}
	return result
}
func nox_xxx_wall_4DF1E0(a1 int32) int32 {
	var (
		i      *int32
		v2     int32
		result int32
		j      int32
		v5     int8
	)
	for i = (*int32)(nox_xxx_wallGetFirstBreakableCli_410870()); i != nil; i = (*int32)(unsafe.Pointer(uintptr(nox_xxx_wallGetNextBreakableCli_410880(i)))) {
		v2 = *(*int32)(unsafe.Add(unsafe.Pointer(i), unsafe.Sizeof(int32(0))*1))
		if int32(*(*uint8)(unsafe.Pointer(uintptr(v2 + 4))))&32 != 0 {
			nox_xxx_wallSendDestroyed_4DF0A0(v2, a1)
		}
	}
	sub_4DF2A0(int8(a1))
	if sub_5071C0() != 0 {
		sub_507190(a1, 1)
	}
	result = bool2int(noxflags.HasGame(noxflags.GameModeQuest))
	if result != 0 {
		for j = int32(uintptr(unsafe.Pointer(noxServer.firstServerObject()))); j != 0; j = int32(uintptr(unsafe.Pointer((*nox_object_t)(unsafe.Pointer(uintptr(j))).Next()))) {
			if (int32(*(*uint8)(unsafe.Pointer(uintptr(j + 8))))&128) != 0 && int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(j + 748))) + 48)))) != 0 {
				sub_4D6A20(a1, j)
			}
		}
		result = sub_4D72C0()
		if result == 1 {
			v5 = int8(sub_4D72C0())
			result = sub_4D7280(a1, v5)
		}
	}
	return result
}
func sub_4DF2A0(a1 int8) *int32 {
	var (
		v1     int32
		result *int32
	)
	v1 = 1 << int32(a1)
	for result = (*int32)(nox_xxx_wallSecretGetFirstWall_410780()); result != nil; result = (*int32)(unsafe.Pointer(uintptr(nox_xxx_wallSecretNext_410790(result)))) {
		if *(*int32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(int32(0))*5))&8 != 0 {
			*(*int32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(int32(0))*7)) |= v1
		} else {
			*(*int32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(int32(0))*7)) &= ^v1
		}
	}
	return result
}
func sub_4DF2E0(a1 int32) {
	var (
		i *byte
		j int32
	)
	if a1 != 31 {
		for i = (*byte)(unsafe.Pointer(nox_server_teamFirst_418B10())); i != nil; i = (*byte)(unsafe.Pointer(nox_server_teamNext_418B60((*nox_team_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(i))))))))) {
			sub_4197C0((*libc.WChar)(unsafe.Pointer(i)), a1)
			for j = int32(uintptr(unsafe.Pointer(nox_xxx_getFirstPlayerUnit_4DA7C0()))); j != 0; j = int32(uintptr(unsafe.Pointer(nox_xxx_getNextPlayerUnit_4DA7F0((*nox_object_t)(unsafe.Pointer(uintptr(j))))))) {
				if nox_xxx_teamCompare2_419180(j+48, uint8(*(*byte)(unsafe.Add(unsafe.Pointer(i), 57)))) != 0 {
					sub_4198A0(j+48, a1, int32(*(*uint32)(unsafe.Pointer(uintptr(j + 36)))))
				}
			}
		}
	}
}
func nox_xxx_netSendSimpleObject2_4DF360(a1 int32, a2 int32) int32 {
	var (
		v2 int16
		v3 float32
		v4 int16
		v5 float32
		v7 [9]byte
	)
	v7[0] = 47
	*(*uint16)(unsafe.Pointer(&v7[3])) = *(*uint16)(unsafe.Pointer(uintptr(a2 + 4)))
	v2 = int16(uint16(nox_xxx_netGetUnitCodeServ_578AC0((*nox_object_t)(unsafe.Pointer(uintptr(a2))))))
	v3 = *(*float32)(unsafe.Pointer(uintptr(a2 + 56)))
	*(*uint16)(unsafe.Pointer(&v7[1])) = uint16(v2)
	v4 = int16(int(v3))
	v5 = *(*float32)(unsafe.Pointer(uintptr(a2 + 60)))
	*(*uint16)(unsafe.Pointer(&v7[5])) = uint16(v4)
	*(*uint16)(unsafe.Pointer(&v7[7])) = uint16(int16(int(v5)))
	return nox_xxx_netSendPacket1_4E5390(a1, int32(uintptr(unsafe.Pointer(&v7[0]))), 9, 0, 1)
}
func sub_4DF3C0(pl *nox_playerInfo) {
	var (
		a1 int32 = int32(uintptr(unsafe.Pointer(pl)))
		v1 int32
		v2 *byte
		v3 *byte
		v4 *byte
		v5 *byte
		v6 int32
		v7 int32
		v8 int32
		v9 int32
	)
	v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 2056))))
	v2 = (*byte)(sub_416640())
	v3 = v2
	if v1 == 0 {
		return
	}
	if sub_40A740() == 0 && !noxflags.HasGame(noxflags.GameFlag16) {
		var v2b uint8 = nox_xxx_getTeamCounter_417DD0()
		if int32(v2b) != 0 {
			v2 = sub_4189D0()
			v4 = v2
			if v2 != nil {
				v2 = (*byte)(unsafe.Pointer(uintptr(nox_xxx_servObjectHasTeam_419130(v1 + 48))))
				if v2 == nil {
					nox_xxx_createAtImpl_4191D0(uint8(*(*byte)(unsafe.Add(unsafe.Pointer(v4), 57))), unsafe.Pointer(uintptr(v1+48)), 1, int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 36)))), 1)
				}
			}
		}
		return
	}
	v2 = *(**byte)(unsafe.Pointer(uintptr(a1 + 2068)))
	if v2 == nil {
		return
	}
	v5 = (*byte)(unsafe.Pointer(nox_server_teamByXxx_418AE0(int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 2068)))))))
	if v5 != nil {
		v6 = v1 + 48
		v7 = nox_xxx_servObjectHasTeam_419130(v1 + 48)
	} else {
		v8 = int32(uint8(*(*byte)(unsafe.Add(unsafe.Pointer(v3), 52))))
		if (noxflags.HasGame(noxflags.GameModeCTF|noxflags.GameModeFlagBall) || noxflags.HasGame(noxflags.GameModeKOTR) && nox_xxx_CheckGameplayFlags_417DA0(4)) && v8 > 2 {
			v8 = 2
		}
		v2 = (*byte)(unsafe.Pointer(uintptr(uint8(sub_417DE0()))))
		if int32(uintptr(unsafe.Pointer(v2))) >= v8 {
			return
		}
		if !noxflags.HasGame(noxflags.GameModeCTF|noxflags.GameModeFlagBall) || (func() bool {
			v9 = int32(uint8(sub_417DE0()))
			v2 = (*byte)(unsafe.Pointer(uintptr(sub_417DC0())))
			return v9 < int32(uintptr(unsafe.Pointer(v2)))
		}()) {
			v2 = sub_418A10()
			v5 = v2
			if v2 == nil {
				return
			}
			sub_418800((*libc.WChar)(unsafe.Pointer(v2)), (*libc.WChar)(unsafe.Pointer(uintptr(a1+2072))), 0)
			sub_418830(int32(uintptr(unsafe.Pointer(v5))), int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 2068)))))
			sub_4184D0((*nox_team_t)(unsafe.Pointer((*libc.WChar)(unsafe.Pointer(v5)))))
			v6 = v1 + 48
			v7 = nox_xxx_servObjectHasTeam_419130(v1 + 48)
		} else {
			return
		}
	}
	if v7 != 0 {
		sub_4196D0(v6, int32(uintptr(unsafe.Pointer(v5))), int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 36)))), 0)
	} else {
		nox_xxx_createAtImpl_4191D0(uint8(*(*byte)(unsafe.Add(unsafe.Pointer(v5), 57))), unsafe.Pointer(uintptr(v6)), 1, int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 36)))), 0)
	}
	return
}
func sub_4DF550() int32 {
	return int32(*memmap.PtrUint32(6112660, 1563148))
}
func sub_4DF580(a1 int32, a2 *uint32) int32 {
	var (
		v2 int32
		v3 int32
		v4 int32
		v5 int32
		v6 int32
		v7 int32
	)
	v2 = int32(*(*uint16)(unsafe.Pointer(uintptr(a1 + 5))))
	if int32(uint16(int16(v2)))-int32(*a2) >= 0 {
		v3 = int32(uint32(uint16(int16(v2))) - *a2)
	} else {
		v3 = int32(*a2 - uint32(v2))
	}
	v4 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*2)))
	if v3 >= v4 {
		return 0
	}
	v5 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*1)))
	v6 = int32(*(*uint16)(unsafe.Pointer(uintptr(a1 + 7))))
	if int32(uint16(int16(v6)))-v5 >= 0 {
		v7 = int32(uint16(int16(v6))) - v5
	} else {
		v7 = v5 - v6
	}
	if v7 >= v4 {
		return 0
	}
	*memmap.PtrUint32(6112660, 1563300) = uint32(a1)
	return 1
}
func sub_4DF5E0(a1 int32, a2 int32) int32 {
	var v3 [3]int32
	v3[0] = int32(*memmap.PtrUint32(6112660, 1563292))
	v3[1] = int32(*memmap.PtrUint32(6112660, 1563296))
	v3[2] = a2
	*memmap.PtrUint32(6112660, 1563300) = 0
	nox_netlist_forEach2_40F0F0(a1, func(arg1 uint32, arg2 int32) int32 {
		return sub_4DF580(int32(arg1), (*uint32)(unsafe.Pointer(uintptr(arg2))))
	}, int32(uintptr(unsafe.Pointer(&v3[0]))))
	return int32(*memmap.PtrUint32(6112660, 1563300))
}
func nox_xxx_netFn_UpdateStream_4DF630(a1 int32, a2 *byte, a3 uint32, a4 unsafe.Pointer) int32 {
	var (
		v3  *byte
		v4  int32
		v5  *byte
		v6  *byte
		v7  **int32
		v8  uint32
		v9  **int32
		v10 *byte
		v11 int32
		v13 *byte
		v14 *uint8
		i   **int32
		v16 int32
		j   **int32
		v18 int32
	)
	v3 = a2
	v4 = a1 - 1
	v5 = (*byte)(unsafe.Add(unsafe.Pointer(a2), a3))
	v6 = (*byte)(unsafe.Pointer(nox_common_playerInfoFromNum_417090(a1 - 1)))
	*memmap.PtrUint32(6112660, 1563308) = 0
	*memmap.PtrUint32(6112660, 1563312) = 0
	*memmap.PtrUint32(6112660, 1563304) = uint32(uintptr(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v6), 16)))))
	v7 = (**int32)(unsafe.Pointer(nox_netlist_getByInd2_40F080(a1-1, &a3)))
	if nox_xxx_chkIsMsgTimestamp_4DF7F0((*uint8)(unsafe.Pointer(v7))) != 0 {
		v8 = a3
		libc.MemCpy(unsafe.Pointer(a2), unsafe.Pointer(v7), int(a3))
		v3 = (*byte)(unsafe.Add(unsafe.Pointer(a2), v8))
		v9 = (**int32)(unsafe.Pointer(nox_netlist_getByInd2_40F080(v4, &a3)))
		if v9 != nil {
			*v3 = 164
			v10 = (*byte)(unsafe.Add(unsafe.Pointer(v3), 1))
			v11 = sub_4DF810(v10, int32(uintptr(unsafe.Pointer(v5))), int32(uintptr(unsafe.Pointer(v9))))
			if v11 == 0 {
				return int32(uintptr(unsafe.Pointer(v10))-uintptr(unsafe.Pointer(a2))) - 1
			}
			v13 = (*byte)(unsafe.Add(unsafe.Pointer(v10), v11))
			for {
				v14 = sub_4DF8F0(v4, (*uint8)(unsafe.Pointer(v13)), int32(uintptr(unsafe.Pointer(v5))))
				v13 = (*byte)(unsafe.Add(unsafe.Pointer(v13), uint32(uintptr(unsafe.Pointer(v14)))))
				if v14 == nil {
					break
				}
			}
			v3 = (*byte)(unsafe.Add(unsafe.Pointer(v13), sub_4DF8C0((*uint8)(unsafe.Pointer(v13)), int32(uintptr(unsafe.Pointer(v5))))))
		}
	}
	for i = (**int32)(unsafe.Pointer(nox_netlist_getInd_40EEB0(v4, 1, &a3))); i != nil; i = (**int32)(unsafe.Pointer(nox_netlist_getInd_40EEB0(v4, 1, &a3))) {
		if *(*byte)(unsafe.Pointer(i)) != 149 || *(*int32)(unsafe.Pointer(&dword_5d4594_2650652)) != 1 || (nox_frame_xxx_2598000%uint32(nox_xxx_rateGet_40A6C0())) == 0 {
			v16 = sub_4DFAF0(unsafe.Pointer(v3), int32(uintptr(unsafe.Pointer(v5))), unsafe.Pointer(i), int32(a3))
			if v16 == 0 {
				return int32(uintptr(unsafe.Pointer(v3)) - uintptr(unsafe.Pointer(a2)))
			}
			v3 = (*byte)(unsafe.Add(unsafe.Pointer(v3), v16))
		}
	}
	*memmap.PtrUint32(6112660, 1563312) = uint32(int32(uintptr(unsafe.Pointer(v3)) - uintptr(unsafe.Pointer(a2))))
	if dword_5d4594_2650652 == 0 || (nox_frame_xxx_2598000%uint32(nox_xxx_rateGet_40A6C0())) == 0 || noxflags.HasGame(noxflags.GameFlag4) {
		nox_xxx_netImportant_4E5770(uint8(int8(v4)), 1)
		for j = (**int32)(unsafe.Pointer(nox_netlist_getInd_40EEB0(v4, 1, &a3))); j != nil; j = (**int32)(unsafe.Pointer(nox_netlist_getInd_40EEB0(v4, 1, &a3))) {
			v18 = sub_4DFAF0(unsafe.Pointer(v3), int32(uintptr(unsafe.Pointer(v5))), unsafe.Pointer(j), int32(a3))
			if v18 == 0 {
				break
			}
			v3 = (*byte)(unsafe.Add(unsafe.Pointer(v3), v18))
		}
	}
	return int32(uintptr(unsafe.Pointer(v3)) - uintptr(unsafe.Pointer(a2)))
}
func nox_xxx_chkIsMsgTimestamp_4DF7F0(a1 *uint8) int32 {
	return bool2int(a1 != nil && (int32(*a1) == 39 || int32(*a1) == 40))
}
func sub_4DF810(a1 *byte, a2 int32, a3 int32) int32 {
	var (
		v4 int8
		v5 bool
		v6 *byte
		v7 int8
		v8 *uint8
		v9 *uint8
	)
	if *(*byte)(unsafe.Pointer(uintptr(a3))) != 195 {
		return 0
	}
	if a2-int32(uintptr(unsafe.Pointer(a1))) < 16 {
		return 0
	}
	*(*uint16)(unsafe.Pointer(uintptr(a3 + 5))) &= 0xFFFE
	*(*uint16)(unsafe.Pointer(uintptr(a3 + 7))) &= 0xFFFE
	v4 = sub_57B930(*memmap.PtrInt32(6112660, 1563304), int32(*(*uint16)(unsafe.Pointer(uintptr(a3 + 1)))), int32(*(*uint16)(unsafe.Pointer(uintptr(a3 + 3)))), nox_frame_xxx_2598000)
	v5 = int32(v4) == -1
	*a1 = byte(v4)
	v6 = (*byte)(unsafe.Add(unsafe.Pointer(a1), 1))
	if v5 {
		*(*uint32)(unsafe.Pointer(v6)) = *(*uint32)(unsafe.Pointer(uintptr(a3 + 1)))
		v6 = (*byte)(unsafe.Add(unsafe.Pointer(a1), 5))
	}
	*(*uint32)(unsafe.Pointer(v6)) = *(*uint32)(unsafe.Pointer(uintptr(a3 + 5)))
	v7 = int8(*(*uint8)(unsafe.Pointer(uintptr(a3 + 9))))
	v8 = (*uint8)(unsafe.Add(unsafe.Pointer(v6), 4))
	*v8 = uint8(v7)
	if *(*byte)(unsafe.Pointer(uintptr(a3 + 10))) != math.MaxUint8 {
		*func() *uint8 {
			p := &v8
			x := *p
			*p = (*uint8)(unsafe.Add(unsafe.Pointer(*p), 1))
			return x
		}() = uint8(int8(int32(v7) | 128))
		*v8 = *(*uint8)(unsafe.Pointer(uintptr(a3 + 10)))
	}
	v9 = (*uint8)(unsafe.Add(unsafe.Pointer(v8), 1))
	*v9 = *(*uint8)(unsafe.Pointer(uintptr(a3 + 11)))
	*memmap.PtrUint32(6112660, 1563292) = uint32(*(*uint16)(unsafe.Pointer(uintptr(a3 + 5))))
	*memmap.PtrUint32(6112660, 1563296) = uint32(*(*uint16)(unsafe.Pointer(uintptr(a3 + 7))))
	return int32(uintptr(unsafe.Pointer((*byte)(unsafe.Pointer(v9))))-uintptr(unsafe.Pointer(a1))) + 1
}
func sub_4DF8C0(a1 *uint8, a2 int32) int32 {
	if a2-int32(uintptr(unsafe.Pointer(a1))) < 3 {
		return 0
	}
	*a1 = 0
	*(*uint8)(unsafe.Add(unsafe.Pointer(a1), 1)) = 0
	*(*uint8)(unsafe.Add(unsafe.Pointer(a1), 2)) = 0
	return 3
}
func sub_4DF8F0(a1 int32, a2 *uint8, a3 int32) *uint8 {
	var (
		v3 *uint8
		v4 int32
		v6 int32
		v7 int32
		v8 int32
	)
	v3 = a2
	v4 = math.MaxInt8
	if nox_netlist_countByInd2_40F0B0(a1) == 0 {
		return nil
	}
	v6 = sub_4DF5E0(a1, math.MaxInt8)
	if v6 == 0 {
		for {
			v4 += math.MaxInt8
			if v4 > 4064 {
				return nil
			}
			v6 = sub_4DF5E0(a1, v4)
			if v6 != 0 {
				if v4 <= math.MaxInt8 {
					break
				}
				*a2 = 0
				v7 = 1
				v3 = (*uint8)(unsafe.Add(unsafe.Pointer(a2), 1))
				goto LABEL_9
			}
		}
	}
	v7 = 0
LABEL_9:
	v8 = sub_4DF9B0(int32(uintptr(unsafe.Pointer(v3))), a3, v6, v7)
	if v8 == -1 {
		return (*uint8)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v3)) - uintptr(unsafe.Pointer(a2))))))
	}
	nox_netlist_findAndFreeBuf_40F000(a1, (*uint8)(unsafe.Pointer(uintptr(v6))))
	return (*uint8)(unsafe.Add(unsafe.Pointer(v3), uint32(v8)-uint32(uintptr(unsafe.Pointer(a2)))))
}
func sub_4DF9B0(a1 int32, a2 int32, a3 int32, a4 int32) int32 {
	var (
		v4  int8
		v5  *uint16
		v6  uint16
		v7  *uint16
		v8  *uint8
		v10 *uint8
		v11 int8
		v12 float32
		v13 float32
	)
	if a2-a1 < 8 || a4 != 0 && a2-a1 < 10 {
		return -1
	}
	v4 = sub_57B930(*memmap.PtrInt32(6112660, 1563304), int32(*(*uint16)(unsafe.Pointer(uintptr(a3 + 1)))), int32(*(*uint16)(unsafe.Pointer(uintptr(a3 + 3)))), nox_frame_xxx_2598000)
	*(*uint8)(unsafe.Pointer(uintptr(a1))) = uint8(v4)
	v5 = (*uint16)(unsafe.Pointer(uintptr(a1 + 1)))
	if int32(v4) == -1 {
		*v5 = *(*uint16)(unsafe.Pointer(uintptr(a3 + 1)))
		*(*uint16)(unsafe.Pointer(uintptr(a1 + 3))) = *(*uint16)(unsafe.Pointer(uintptr(a3 + 3)))
		v5 = (*uint16)(unsafe.Pointer(uintptr(a1 + 5)))
	}
	if a4 != 0 {
		v6 = *(*uint16)(unsafe.Pointer(uintptr(a3 + 5)))
		if int32(v6) > 6000 || int32(*(*uint16)(unsafe.Pointer(uintptr(a3 + 7)))) > 6000 {
			return 0
		}
		*v5 = v6
		v7 = (*uint16)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(uint16(0))*1))
		*v7 = *(*uint16)(unsafe.Pointer(uintptr(a3 + 7)))
		v8 = (*uint8)(unsafe.Pointer((*uint16)(unsafe.Add(unsafe.Pointer(v7), unsafe.Sizeof(uint16(0))*1))))
	} else {
		v12 = float32(float64(*(*uint16)(unsafe.Pointer(uintptr(a3 + 5)))))
		*(*uint8)(unsafe.Pointer(v5)) = uint8(int8(int(v12) - int32(*memmap.PtrUint8(6112660, 1563292))))
		v10 = (*uint8)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(v5))), 1))
		v13 = float32(float64(*(*uint16)(unsafe.Pointer(uintptr(a3 + 7)))))
		*v10 = uint8(int8(int(v13) - int32(*memmap.PtrUint8(6112660, 1563296))))
		v8 = (*uint8)(unsafe.Add(unsafe.Pointer(v10), 1))
	}
	*memmap.PtrUint32(6112660, 1563292) = uint32(*(*uint16)(unsafe.Pointer(uintptr(a3 + 5))))
	*memmap.PtrUint32(6112660, 1563296) = uint32(*(*uint16)(unsafe.Pointer(uintptr(a3 + 7))))
	if *(*byte)(unsafe.Pointer(uintptr(a3))) == 48 || *(*byte)(unsafe.Pointer(uintptr(a3))) == 195 {
		v11 = int8(*(*uint8)(unsafe.Pointer(uintptr(a3 + 9))))
		*v8 = uint8(v11)
		if *(*byte)(unsafe.Pointer(uintptr(a3 + 10))) != math.MaxUint8 {
			*func() *uint8 {
				p := &v8
				x := *p
				*p = (*uint8)(unsafe.Add(unsafe.Pointer(*p), 1))
				return x
			}() = uint8(int8(int32(v11) | 128))
			*v8 = *(*uint8)(unsafe.Pointer(uintptr(a3 + 10)))
		}
		if *(*byte)(unsafe.Pointer(uintptr(a3))) == 195 {
			*func() *uint8 {
				p := &v8
				*p = (*uint8)(unsafe.Add(unsafe.Pointer(*p), 1))
				return *p
			}() = *(*uint8)(unsafe.Pointer(uintptr(a3 + 11)))
		}
		v8 = (*uint8)(unsafe.Add(unsafe.Pointer(v8), 1))
	}
	return int32(uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v8), -a1)))))
}
func sub_4DFAF0(a1 unsafe.Pointer, a2 int32, a3 unsafe.Pointer, a4 int32) int32 {
	var result int32
	result = a4
	if a2-int32(uintptr(a1)) < a4 {
		return 0
	}
	libc.MemCpy(a1, a3, int(a4))
	return result
}
func sub_4DFB50(a1 int32, a2 int32) *uint32 {
	*(*uint8)(unsafe.Pointer(uintptr(a2 + 440))) |= 8
	return nox_xxx_aud_501960(75, (*nox_object_t)(unsafe.Pointer(uintptr(a2))), 0, 0)
}
func sub_4DFB80(a1 int32, a2 int32) *uint32 {
	if nox_xxx_enchantItemTestInventory_4DFBB0(a2, 8) == 0 {
		*(*uint8)(unsafe.Pointer(uintptr(a2 + 440))) &= 247
	}
	return nox_xxx_aud_501960(76, (*nox_object_t)(unsafe.Pointer(uintptr(a2))), 0, 0)
}
func nox_xxx_enchantItemTestInventory_4DFBB0(a1 int32, a2 int8) int32 {
	var (
		v2 *uint32
		v3 int32
		v4 int32
		v5 int32
		v6 *uint8
	)
	if a1 == 0 {
		return 0
	}
	if int32(a2) == 0 {
		return 0
	}
	v2 = *(**uint32)(unsafe.Pointer(uintptr(a1 + 504)))
	if v2 == nil {
		return 0
	}
	for {
		v3 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*4)))
		if v3&256 != 0 {
			if *(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*2))&0x13001000 != 0 {
				break
			}
		}
	LABEL_13:
		v2 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*124)))))
		if v2 == nil {
			return 0
		}
	}
	v4 = 0
	v5 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*173)))
	for *(*uint32)(unsafe.Pointer(uintptr(v5))) == 0 {
	LABEL_12:
		v4++
		v5 += 4
		if v4 >= 4 {
			goto LABEL_13
		}
	}
	v6 = (*uint8)(memmap.PtrOff(0x587000, 200164))
	for *(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v5))) + 112))) != *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v6))), -int(unsafe.Sizeof(uint32(0))*1)))) || int32(a2) != int32(*v6) {
		v6 = (*uint8)(unsafe.Add(unsafe.Pointer(v6), 20))
		if int32(uintptr(unsafe.Pointer(v6))) >= int32(uintptr(memmap.PtrOff(0x587000, 200284))) {
			goto LABEL_12
		}
	}
	return 1
}
func nox_xxx_effectSpeedEngage_4DFC30(a1 int32, a2 int32) {
	var (
		v2 int32
		v3 int32
		v5 int32
	)
	v2 = a2
	if a2 != 0 {
		if int32(*(*uint8)(unsafe.Pointer(uintptr(a2 + 8))))&4 != 0 {
			v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 748))))
			*(*uint8)(unsafe.Pointer(uintptr(a2 + 440))) |= 16
			*(*float32)(unsafe.Pointer(&v5)) = *(*float32)(unsafe.Pointer(uintptr(a1 + 120))) + *(*float32)(unsafe.Pointer(uintptr(a2 + 552)))
			*(*float32)(unsafe.Pointer(uintptr(v2 + 552))) = *(*float32)(unsafe.Pointer(&v5))
			nox_xxx_netReportStatsSpeed_4D9360(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v3 + 276))) + 2064)))), (*uint32)(unsafe.Pointer(uintptr(v2))), 0, v5)
			nox_xxx_aud_501960(59, (*nox_object_t)(unsafe.Pointer(uintptr(v2))), 0, 0)
		}
	}
}
func nox_xxx_effectSpeedDisengage_4DFCA0(a1 int32, a2 int32) {
	var (
		v2 int32
		v3 int32
		v4 int32
	)
	v2 = a2
	if a2 != 0 && int32(*(*uint8)(unsafe.Pointer(uintptr(a2 + 8))))&4 != 0 {
		if nox_xxx_enchantItemTestInventory_4DFBB0(a2, 16) == 0 {
			*(*uint8)(unsafe.Pointer(uintptr(a2 + 440))) &= 239
		}
		v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 748))))
		*(*float32)(unsafe.Pointer(&v4)) = *(*float32)(unsafe.Pointer(uintptr(a2 + 552))) - *(*float32)(unsafe.Pointer(uintptr(a1 + 120)))
		*(*float32)(unsafe.Pointer(uintptr(v2 + 552))) = *(*float32)(unsafe.Pointer(&v4))
		nox_xxx_netReportStatsSpeed_4D9360(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v3 + 276))) + 2064)))), (*uint32)(unsafe.Pointer(uintptr(v2))), 0, v4)
		nox_xxx_aud_501960(60, (*nox_object_t)(unsafe.Pointer(uintptr(v2))), 0, 0)
	}
}
func sub_4DFD10(a1 int32, a2 int32) *uint32 {
	*(*uint8)(unsafe.Pointer(uintptr(a2 + 440))) |= 1
	return nox_xxx_aud_501960(102, (*nox_object_t)(unsafe.Pointer(uintptr(a2))), 0, 0)
}
func nox_xxx_modifFireProtection_4DFD40(a1 int32, a2 int32, a3 int32) {
	if a2 != 0 && a3 != 0 {
		if nox_xxx_enchantItemTestInventory_4DFBB0(a2, 1) == 0 {
			*(*uint8)(unsafe.Pointer(uintptr(a2 + 440))) &= 254
		}
		nox_xxx_aud_501960(103, (*nox_object_t)(unsafe.Pointer(uintptr(a2))), 0, 0)
	}
}
func nox_xxx_buff_4DFD80(a1 int32, a2 int32) *uint32 {
	*(*uint8)(unsafe.Pointer(uintptr(a2 + 440))) |= 4
	return nox_xxx_aud_501960(106, (*nox_object_t)(unsafe.Pointer(uintptr(a2))), 0, 0)
}
func sub_4DFDB0(a1 int32, a2 int32) *uint32 {
	if nox_xxx_enchantItemTestInventory_4DFBB0(a2, 4) == 0 {
		*(*uint8)(unsafe.Pointer(uintptr(a2 + 440))) &= 251
	}
	return nox_xxx_aud_501960(107, (*nox_object_t)(unsafe.Pointer(uintptr(a2))), 0, 0)
}
func nox_xxx_checkPoisonProtectEnch_4DFDE0(a1 int32, a2 int32) *uint32 {
	*(*uint8)(unsafe.Pointer(uintptr(a2 + 440))) |= 2
	return nox_xxx_aud_501960(110, (*nox_object_t)(unsafe.Pointer(uintptr(a2))), 0, 0)
}
func sub_4DFE10(a1 int32, a2 int32) *uint32 {
	if nox_xxx_enchantItemTestInventory_4DFBB0(a2, 2) == 0 {
		*(*uint8)(unsafe.Pointer(uintptr(a2 + 440))) &= 253
	}
	return nox_xxx_aud_501960(111, (*nox_object_t)(unsafe.Pointer(uintptr(a2))), 0, 0)
}
func nox_xxx_checkFireProtect_4DFE40(a1 *uint32) float64 {
	var (
		v1     float64
		result float64
		v3     *uint32
		v4     int32
		v5     *int32
		v6     int32
		v7     int32
		v8     *int32
		v9     int32
		v10    int32
		v11    uint8
		v12    float32
	)
	v1 = 0.0
	v12 = 0.0
	if a1 == nil {
		return 0.0
	}
	v3 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*126)))))
	if v3 != nil {
		for {
			v4 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*4)))
			if v4&256 != 0 && *(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*2))&0x13001000 != 0 {
				v5 = (*int32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*173)))))
				v6 = 4
				for {
					v7 = *v5
					if *v5 != 0 && cgoFuncAddr(cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(v7 + 112))), (*func(int32, int32) *uint32)(nil))) == cgoFuncAddr(sub_4DFD10) {
						v1 = v1 + float64(*(*float32)(unsafe.Pointer(uintptr(v7 + 120))))
					}
					v5 = (*int32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(int32(0))*1))
					v6--
					if v6 == 0 {
						break
					}
				}
			}
			v3 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*124)))))
			if v3 == nil {
				break
			}
		}
		v12 = float32(v1)
	}
	if *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*2))&0x13001000 != 0 {
		v8 = (*int32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*173)))))
		v9 = 4
		for {
			v10 = *v8
			if *v8 != 0 && cgoFuncAddr(cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(v10 + 112))), (*func(int32, int32) *uint32)(nil))) == cgoFuncAddr(sub_4DFD10) {
				v1 = v1 + float64(*(*float32)(unsafe.Pointer(uintptr(v10 + 120))))
			}
			v8 = (*int32)(unsafe.Add(unsafe.Pointer(v8), unsafe.Sizeof(int32(0))*1))
			v9--
			if v9 == 0 {
				break
			}
		}
		v12 = float32(v1)
	}
	if v1 > 0.5 {
		v12 = 0.5
	}
	if nox_xxx_testUnitBuffs_4FF350((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1)))))), 17) != 0 {
		v11 = uint8(nox_xxx_buffGetPower_4FF570(int32(uintptr(unsafe.Pointer(a1))), 17))
		result = nox_xxx_gamedataGetFloatTable_419D70(CString("FireSpellProtection"), int32(uint32(v11)-1)) + float64(v12)
	} else {
		result = float64(v12)
	}
	if result > 0.60000002 {
		result = 0.60000002
	}
	return result
}
func nox_xxx_checkElectrProtect_4DFF40(a1 *uint32) float64 {
	var (
		v1     float64
		result float64
		v3     *uint32
		v4     int32
		v5     *int32
		v6     int32
		v7     int32
		v8     *int32
		v9     int32
		v10    int32
		v11    uint8
		v12    float32
	)
	v1 = 0.0
	v12 = 0.0
	if a1 == nil {
		return 0.0
	}
	v3 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*126)))))
	if v3 != nil {
		for {
			v4 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*4)))
			if v4&256 != 0 && *(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*2))&0x13001000 != 0 {
				v5 = (*int32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*173)))))
				v6 = 4
				for {
					v7 = *v5
					if *v5 != 0 && cgoFuncAddr(cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(v7 + 112))), (*func(int32, int32) *uint32)(nil))) == cgoFuncAddr(nox_xxx_buff_4DFD80) {
						v1 = v1 + float64(*(*float32)(unsafe.Pointer(uintptr(v7 + 120))))
					}
					v5 = (*int32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(int32(0))*1))
					v6--
					if v6 == 0 {
						break
					}
				}
			}
			v3 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*124)))))
			if v3 == nil {
				break
			}
		}
		v12 = float32(v1)
	}
	if *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*2))&0x13001000 != 0 {
		v8 = (*int32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*173)))))
		v9 = 4
		for {
			v10 = *v8
			if *v8 != 0 && cgoFuncAddr(cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(v10 + 112))), (*func(int32, int32) *uint32)(nil))) == cgoFuncAddr(nox_xxx_buff_4DFD80) {
				v1 = v1 + float64(*(*float32)(unsafe.Pointer(uintptr(v10 + 120))))
			}
			v8 = (*int32)(unsafe.Add(unsafe.Pointer(v8), unsafe.Sizeof(int32(0))*1))
			v9--
			if v9 == 0 {
				break
			}
		}
		v12 = float32(v1)
	}
	if v1 >= 0.5 {
		v12 = 0.5
	}
	if nox_xxx_testUnitBuffs_4FF350((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1)))))), 20) != 0 {
		v11 = uint8(nox_xxx_buffGetPower_4FF570(int32(uintptr(unsafe.Pointer(a1))), 20))
		result = nox_xxx_gamedataGetFloatTable_419D70(CString("ElectricitySpellProtection"), int32(uint32(v11)-1)) + float64(v12)
	} else {
		result = float64(v12)
	}
	if result > 0.60000002 {
		result = 0.60000002
	}
	return result
}
func nox_xxx_getPoisonDmg_4E0040(a1 *uint32) float64 {
	var (
		v1     float64
		result float64
		v3     *uint32
		v4     int32
		v5     *int32
		v6     int32
		v7     int32
		v8     *int32
		v9     int32
		v10    int32
		v11    uint8
		v12    float32
	)
	v1 = 0.0
	v12 = 0.0
	if a1 == nil {
		return 0.0
	}
	v3 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*126)))))
	if v3 != nil {
		for {
			v4 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*4)))
			if v4&256 != 0 && *(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*2))&0x13001000 != 0 {
				v5 = (*int32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*173)))))
				v6 = 4
				for {
					v7 = *v5
					if *v5 != 0 && cgoFuncAddr(cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(v7 + 112))), (*func(int32, int32) *uint32)(nil))) == cgoFuncAddr(nox_xxx_checkPoisonProtectEnch_4DFDE0) {
						v1 = v1 + float64(*(*float32)(unsafe.Pointer(uintptr(v7 + 120))))
					}
					v5 = (*int32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(int32(0))*1))
					v6--
					if v6 == 0 {
						break
					}
				}
			}
			v3 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*124)))))
			if v3 == nil {
				break
			}
		}
		v12 = float32(v1)
	}
	if *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*2))&0x13001000 != 0 {
		v8 = (*int32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*173)))))
		v9 = 4
		for {
			v10 = *v8
			if *v8 != 0 && cgoFuncAddr(cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(v10 + 112))), (*func(int32, int32) *uint32)(nil))) == cgoFuncAddr(nox_xxx_checkPoisonProtectEnch_4DFDE0) {
				v1 = v1 + float64(*(*float32)(unsafe.Pointer(uintptr(v10 + 120))))
			}
			v8 = (*int32)(unsafe.Add(unsafe.Pointer(v8), unsafe.Sizeof(int32(0))*1))
			v9--
			if v9 == 0 {
				break
			}
		}
		v12 = float32(v1)
	}
	if v1 > 0.69999999 {
		v12 = 0.69999999
	}
	if nox_xxx_testUnitBuffs_4FF350((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1)))))), 18) != 0 {
		v11 = uint8(nox_xxx_buffGetPower_4FF570(int32(uintptr(unsafe.Pointer(a1))), 18))
		result = nox_xxx_gamedataGetFloatTable_419D70(CString("PoisonSpellProtection"), int32(uint32(v11)-1)) + float64(v12)
	} else {
		result = float64(v12)
	}
	if result > 0.89999998 {
		result = 0.89999998
	}
	return result
}
func sub_4E0140(a1 int32, a2 int32) *uint32 {
	*(*uint8)(unsafe.Pointer(uintptr(a2 + 440))) |= 32
	return nox_xxx_aud_501960(123, (*nox_object_t)(unsafe.Pointer(uintptr(a2))), 0, 0)
}
func sub_4E0170(a1 int32, a2 int32) {
	if a2 != 0 && int32(*(*uint8)(unsafe.Pointer(uintptr(a2 + 8))))&4 != 0 {
		if nox_xxx_enchantItemTestInventory_4DFBB0(a2, 32) == 0 {
			*(*uint8)(unsafe.Pointer(uintptr(a2 + 440))) &= 223
		}
		nox_xxx_aud_501960(124, (*nox_object_t)(unsafe.Pointer(uintptr(a2))), 0, 0)
	}
}
func nox_xxx_effectRegeneration_4E01D0(a1 int32, a2 int32) {
	var (
		v2 *uint32
		v3 uint16
		v4 uint32
		v5 int32
	)
	if a2 != 0 {
		v2 = *(**uint32)(unsafe.Pointer(uintptr(a2 + 492)))
		if v2 != nil {
			if *(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*139)) != 0 {
				if (nox_frame_xxx_2598000-*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*134))) >= uint32(int32(nox_gameFPS)) && (*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*4))&32800) == 0 {
					v3 = uint16(nox_xxx_unitGetMaxHP_4EE7A0(int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 492))))))
					if int32(uint16(nox_xxx_unitGetHP_4EE780((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v2))))))))) < int32(v3) {
						v4 = *(*uint32)(unsafe.Pointer(uintptr(a1 + 108)))
						if *(*uint32)(unsafe.Pointer(uintptr(a2 + 8)))&0x2000000 != 0 {
							v5 = nox_xxx_unitArmorInventoryEquipFlags_415C70((*nox_object_t)(unsafe.Pointer(uintptr(a2))))
							if v5&0x4000 != 0 {
								v4 /= 3
							}
						}
						if (nox_frame_xxx_2598000 % (v4 * nox_gameFPS / uint32(uint16(nox_xxx_unitGetMaxHP_4EE7A0(int32(uintptr(unsafe.Pointer(v2)))))))) == 0 {
							nox_xxx_unitAdjustHP_4EE460((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v2)))))), 1)
						}
					}
				}
			}
		}
	}
}
func nox_xxx_attribContinualReplen_4E02C0(a1 int32, a2 *uint32) {
	var (
		v2 int32
		v3 int32
		v4 uint8
		v5 uint8
		v6 int32
		v7 int32
		v8 int32
	)
	if a2 != nil {
		if (nox_frame_xxx_2598000 % *(*uint32)(unsafe.Pointer(uintptr(a1 + 108)))) == 0 {
			v2 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*2)))
			v3 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*184)))
			if v2&4096 != 0 {
				v4 = *(*uint8)(unsafe.Pointer(uintptr(v3 + 108)))
				v5 = *(*uint8)(unsafe.Pointer(uintptr(v3 + 109)))
				v6 = int32(func() uint8 {
					p := &v4
					x := *p
					*p++
					return x
				}())
				*(*uint8)(unsafe.Pointer(uintptr(v3 + 108))) = v4
				if int32(v4) > int32(v5) {
					*(*uint8)(unsafe.Pointer(uintptr(v3 + 108))) = v5
				}
				v7 = int32(*(*uint8)(unsafe.Pointer(uintptr(v3 + 108))))
				if v6 != v7 {
					v8 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*123)))
					if v8 != 0 {
						if int32(*(*uint8)(unsafe.Pointer(uintptr(v8 + 8))))&4 != 0 {
							nox_xxx_netReportCharges_4D82B0(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v8 + 748))) + 276))) + 2064)))), (*nox_object_t)(unsafe.Pointer(a2)), int8(v7), int8(v5))
						}
					}
				}
			}
		}
	}
}
func sub_4E0370(a1 int32, a2 int32, a3 int32, a4 int32, a5 int32, a6 *float32) *float32 {
	var result *float32
	result = a6
	*a6 = *(*float32)(unsafe.Pointer(uintptr(a1 + 80))) * *a6
	return result
}
func sub_4E0380(a1 int32, a2 int32, a3 int32, a4 int32, a5 int32, a6 *float32) *float32 {
	var result *float32
	result = a6
	*a6 = float32((1.0 - float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 80)))) + 1.0) * float64(*a6))
	return result
}
func nox_xxx_inversionEffect_4E03D0(a1 int32, a2 int32, a3 int32, a4 int32, a5 int32, a6 *int32) int32 {
	var result int32
	result = bool2int(*(*uint32)(unsafe.Pointer(uintptr(a1 + 96))) >= 1)
	*a6 = result
	return result
}
func nox_xxx_unusedCheckGripEffect_4E03F0(a1 int32, a2 int32, a3 int32, a4 int32) int32 {
	var (
		v4  int32
		v5  int32
		v6  *uint32
		v7  int32
		v8  func(int32, int32, int32, int32, int32, *int32) int32
		v9  bool
		i   *uint32
		v12 int32
	)
	v4 = a3
	if a3 == 0 || (*(*uint32)(unsafe.Pointer(uintptr(a3 + 8)))&0x3001000) == 0 {
		return 0
	}
	v5 = a4
	v6 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a3 + 692))) + 8)))
	v12 = 2
	for i = v6; ; i = (*uint32)(unsafe.Add(unsafe.Pointer(i), unsafe.Sizeof(uint32(0))*1)) {
		v7 = int32(*v6)
		if v7 != 0 {
			v8 = cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(v7 + 88))), (*func(int32, int32, int32, int32, int32, *int32) int32)(nil))
			if cgoFuncAddr(v8) == cgoFuncAddr(nox_xxx_gripEffect_4E0480) {
				v8(v7, v4, a2, v5, a1, &a4)
				if a4 == 0 {
					break
				}
			}
		}
		v6 = (*uint32)(unsafe.Add(unsafe.Pointer(i), unsafe.Sizeof(uint32(0))*1))
		v9 = func() int32 {
			p := &v12
			*p++
			return *p
		}() < 4
		if !v9 {
			return 0
		}
	}
	return 1
}
func nox_xxx_gripEffect_4E0480(a1 int32, a2 int32, a3 int32, a4 int32, a5 int32, a6 *int32) int32 {
	var result int32
	result = bool2int(*(*uint32)(unsafe.Pointer(uintptr(a1 + 96))) < 1)
	*a6 = result
	return result
}
func nox_xxx_effectDamageMultiplier_4E04C0(a1 int32, a2 int32, a3 int32, a4 int32, a5 *float32) *float32 {
	var result *float32
	result = a5
	*a5 = *(*float32)(unsafe.Pointer(uintptr(a1 + 44))) * *a5
	return result
}
func nox_xxx_stunEffect_4E04D0(a1 int32, a2 int32, a3 int32, a4 int32) {
	var (
		v4 int32
		v5 int32
		v6 int32
		v8 int32
		v9 [3]int32
	)
	v4 = a4
	if int32(*(*uint8)(unsafe.Pointer(uintptr(a4 + 8))))&6 != 0 {
		v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(a4 + 60))))
		v9[1] = int32(*(*uint32)(unsafe.Pointer(uintptr(a4 + 56))))
		v6 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 60))))
		v9[2] = v5
		v9[0] = a4
		nox_xxx_castStun_52C2C0(74, a3, a3, a3, &v9[0], int8(v6))
		if int32(*(*uint8)(unsafe.Pointer(uintptr(v4 + 8))))&4 != 0 {
			v8 = int32(*(*uint32)(unsafe.Pointer(uintptr(v4 + 748))))
			a4 = 0
			nox_xxx_netInformTextMsg_4DA0F0(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v8 + 276))) + 2064)))), 13, &a4)
		}
	}
}
func nox_xxx_fireEffect_4E0550(a1 int32, a2 int32, a3 int32, a4 int32) *uint32 {
	var (
		result *uint32
		v5     float32
	)
	result = (*uint32)(unsafe.Pointer(uintptr(a1)))
	v5 = *(*float32)(unsafe.Pointer(uintptr(a1 + 56)))
	if a4 != 0 {
		(cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(a4 + 716))), (*func(int32, int32, int32, uint32, int32))(nil)))(a4, a3, a2, uint32(int32(int64(v5))), 7)
		nox_xxx_netSparkExplosionFx_5231B0((*float32)(unsafe.Pointer(uintptr(a4+56))), int8(int64(float64(v5)*10.0)))
		result = nox_xxx_aud_501960(224, (*nox_object_t)(unsafe.Pointer(uintptr(a4))), 0, 0)
	}
	return result
}
func nox_xxx_fireRingEffect_4E05B0(a1 int32, a2 int32, a3 int32) int32 {
	var (
		result int32
		v4     int32
		v5     [3]int32
	)
	result = a3
	if a3 != 0 {
		v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(a3 + 60))))
		v5[1] = int32(*(*uint32)(unsafe.Pointer(uintptr(a3 + 56))))
		v5[2] = v4
		v5[0] = 0
		result = nox_xxx_spellCastCleansingFlame_52D5C0(10, a3, a3, a3, int32(uintptr(unsafe.Pointer(&v5[0]))), int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 48)))))
	}
	return result
}
func nox_xxx_blueFREffect_4E05F0(a1 int32, a2 int32, a3 int32) int32 {
	var (
		result int32
		v4     int32
		v5     [3]int32
	)
	result = a3
	if a3 != 0 {
		v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(a3 + 60))))
		v5[1] = int32(*(*uint32)(unsafe.Pointer(uintptr(a3 + 56))))
		v5[2] = v4
		v5[0] = 0
		result = nox_xxx_spellCastCleansingFlame_52D5C0(11, a3, a3, a3, int32(uintptr(unsafe.Pointer(&v5[0]))), int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 48)))))
	}
	return result
}
func nox_xxx_recoilEffect_4E0640(a1 int32, a2 int32, a3 int32, a4 int32) {
	if a2 != 0 {
		if a4 != 0 {
			nox_xxx_objectApplyForce_52DF80((*float32)(unsafe.Pointer(uintptr(a2+56))), (*nox_object_t)(unsafe.Pointer(uintptr(a4))), *(*float32)(unsafe.Pointer(uintptr(a1 + 56))))
		}
	}
}
func nox_xxx_confuseEffect_4E0670(a1 int32, a2 int32, a3 int32, a4 int32) {
	var (
		v4 int32
		v5 int32
		v6 int32
		v8 int32
		v9 [3]int32
	)
	v4 = a4
	if int32(*(*uint8)(unsafe.Pointer(uintptr(a4 + 8))))&6 != 0 {
		v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(a4 + 60))))
		v9[1] = int32(*(*uint32)(unsafe.Pointer(uintptr(a4 + 56))))
		v6 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 60))))
		v9[2] = v5
		v9[0] = a4
		nox_xxx_castConfuse_52C1E0(12, a3, a3, a3, &v9[0], int8(v6))
		if int32(*(*uint8)(unsafe.Pointer(uintptr(v4 + 8))))&4 != 0 {
			v8 = int32(*(*uint32)(unsafe.Pointer(uintptr(v4 + 748))))
			a4 = 1
			nox_xxx_netInformTextMsg_4DA0F0(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v8 + 276))) + 2064)))), 13, &a4)
		}
	}
}
func nox_xxx_lightngEffect_4E06F0(a1 int32, a2 int32, a3 int32, a4 int32) {
	if a4 != 0 {
		(cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(a4 + 716))), (*func(int32, int32, int32, uint32, int32))(nil)))(a4, a3, a2, uint32(int32(int64(*(*float32)(unsafe.Pointer(uintptr(a1 + 56)))))), 9)
		nox_xxx_netSendPointFx_522FF0(-127, (*float2)(unsafe.Pointer(uintptr(a4+56))))
		nox_xxx_aud_501960(225, (*nox_object_t)(unsafe.Pointer(uintptr(a4))), 0, 0)
	}
}
func nox_xxx_drainMEffect_4E0740(a1 int32, a2 int32, a3 int32, a4 int32, a5 *int32) {
	var (
		v5 int32
		v6 uint16
		v7 int16
	)
	if a3 != 0 && a4 != 0 && int32(*(*uint8)(unsafe.Pointer(uintptr(a4 + 8))))&6 != 0 {
		v5 = int32(int64(float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 68)))) * float64(*a5)))
		if v5 == 0 {
			v5 = 1
		}
		v6 = uint16(nox_xxx_unitGetOldMana_4EEC80(a4))
		v7 = int16(v6)
		if int32(v6) < v5 {
			nox_xxx_playerManaSub_4EEBF0(a4, int32(v6))
			nox_xxx_playerManaAdd_4EEB80((*nox_object_t)(unsafe.Pointer(uintptr(a3))), v7)
		} else {
			nox_xxx_playerManaSub_4EEBF0(a4, v5)
			nox_xxx_playerManaAdd_4EEB80((*nox_object_t)(unsafe.Pointer(uintptr(a3))), int16(v5))
		}
	}
}
func nox_xxx_vampirismEffect_4E07C0(a1 int32, a2 int32, a3 int32, a4 int32, a5 *int32) {
	var (
		v5 int32
		v6 int32
		v7 int32
	)
	if a3 != 0 {
		if a4 != 0 {
			v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(a4 + 8))))
			if v5&6 != 0 {
				if (v5&2) == 0 || (int32(*(*uint8)(unsafe.Pointer(uintptr(a4 + 12))))&64) == 0 {
					v6 = int32(int64(float64(*a5) * float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 68))))))
					if v6 == 0 {
						v6 = 1
					}
					v7 = int32(uint16(nox_xxx_unitGetHP_4EE780((*nox_object_t)(unsafe.Pointer(uintptr(a4))))))
					if v7 < v6 {
						nox_xxx_unitAdjustHP_4EE460((*nox_object_t)(unsafe.Pointer(uintptr(a3))), v7)
					} else {
						nox_xxx_unitAdjustHP_4EE460((*nox_object_t)(unsafe.Pointer(uintptr(a3))), v6)
					}
				}
			}
		}
	}
}
func nox_xxx_poisonEffect_4E0850(a1 int32, a2 int32, a3 int32, a4 int32) {
	var (
		v4 int32
		v5 int32
	)
	v4 = a4
	if ((int32(*(*uint8)(unsafe.Pointer(uintptr(a4 + 8))))&4) == 0 || int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a4 + 748))) + 88)))) != 16 || (nox_server_testTwoPointsAndDirection_4E6E50((*float2)(unsafe.Pointer(uintptr(a4+56))), int32(*(*int16)(unsafe.Pointer(uintptr(a4 + 124)))), (*float2)(unsafe.Pointer(uintptr(a3+56))))&1) == 0) && int32(*(*uint8)(unsafe.Pointer(uintptr(v4 + 8))))&6 != 0 && nox_xxx_activatePoison_4EE7E0(v4, 1, int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 72))))) != 0 {
		if int32(*(*uint8)(unsafe.Pointer(uintptr(v4 + 8))))&4 != 0 {
			v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(v4 + 748))))
			a4 = 2
			nox_xxx_netInformTextMsg_4DA0F0(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v5 + 276))) + 2064)))), 13, &a4)
		}
	}
}
func nox_xxx_sympathyEffect_4E08E0(a1 int32, a2 int32, a3 int32, a4 int32, a5 *int32) {
	var (
		v5  int32
		v6  int32
		v7  int32
		v8  int32
		v9  int32
		v10 float32
	)
	v5 = a3
	if a3 != 0 && a4 != 0 && int32(*(*uint8)(unsafe.Pointer(uintptr(a4 + 8))))&6 != 0 {
		v8 = a4
		v10 = *(*float32)(unsafe.Pointer(uintptr(a1 + 68)))
		v6 = *a5
		v9 = *a5
		v7 = int32(uint16(nox_xxx_unitGetHP_4EE780((*nox_object_t)(unsafe.Pointer(uintptr(v8))))))
		if v7 < v6 {
			v9 = v7
		}
		nox_xxx_unitDamageClear_4EE5E0((*nox_object_t)(unsafe.Pointer(uintptr(v5))), int32(int64(float64(v9)*float64(v10))))
	}
}
func nox_xxx_itemCheckReadinessEffect_4E0960(a1 int32) int32 {
	var (
		v1 int32
		v2 int32
		i  int32
	)
	v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 692))))
	if a1 == 0 || (*(*uint32)(unsafe.Pointer(uintptr(a1 + 8)))&0x13001000) == 0 {
		return 0
	}
	v2 = 2
	for i = v1 + 8; *(*uint32)(unsafe.Pointer(uintptr(i))) == 0 || cgoFuncAddr(cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(i))) + 40))), (*func() int32)(nil))) != cgoFuncAddr(nullsub_22); i += 4 {
		if func() int32 {
			p := &v2
			*p++
			return *p
		}() >= 4 {
			return 0
		}
	}
	return int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v1 + v2*4))) + 48))))
}
func nox_xxx_effectProjectileSpeed_4E09B0(a1 int32, a2 int32, a3 int32, a4 int32, a5 int32) int32 {
	var result int32
	result = a5
	*(*float32)(unsafe.Pointer(uintptr(a5 + 544))) = *(*float32)(unsafe.Pointer(uintptr(a1 + 44))) * *(*float32)(unsafe.Pointer(uintptr(a5 + 544)))
	return result
}
func nox_xxx_parseDamageTypeByName_4E0A00(a1 *byte) int32 {
	var (
		v1 int32
		v2 **byte
	)
	v1 = 0
	v2 = (**byte)(memmap.PtrOff(0x587000, 200728))
	for {
		if libc.StrCmp(a1, *v2) == 0 {
			break
		}
		v2 = (**byte)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof((*byte)(nil))*1))
		v1++
		if int32(uintptr(unsafe.Pointer(v2))) >= int32(uintptr(memmap.PtrOff(0x587000, 200800))) {
			break
		}
	}
	return v1
}
func nox_xxx_projectileReflect_4E0A70(a1 int32, a2 int32) int32 {
	var (
		result int32
		v3     float64
		v4     float64
		v5     int16
		v6     float64
		v7     float64
		v8     int32
		v9     float32
		v10    float32
		v11    float32
		v12    float32
		v13    float32
	)
	result = a1
	v3 = float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 56))) - *(*float32)(unsafe.Pointer(uintptr(a2 + 56))))
	v4 = float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 60))) - *(*float32)(unsafe.Pointer(uintptr(a2 + 60))))
	*(*uint16)(unsafe.Pointer(uintptr(a1 + 124))) += 128
	v5 = int16(*(*uint16)(unsafe.Pointer(uintptr(a1 + 124))))
	v11 = float32(-v4)
	v13 = float32(math.Sqrt(v4*v4+v3*v3) + 0.1)
	v6 = (v4*float64(*(*float32)(unsafe.Pointer(uintptr(result + 84)))) + v3*float64(*(*float32)(unsafe.Pointer(uintptr(result + 80))))) / float64(v13)
	v9 = float32(v6 * v3 / float64(v13))
	v10 = float32(v6 * v4 / float64(v13))
	v7 = (float64(v11**(*float32)(unsafe.Pointer(uintptr(result + 80)))) + v3*float64(*(*float32)(unsafe.Pointer(uintptr(result + 84))))) / float64(v13)
	v12 = float32(v7 * float64(v11) / float64(v13))
	*(*float32)(unsafe.Pointer(uintptr(result + 80))) = v12 - v9
	*(*float32)(unsafe.Pointer(uintptr(result + 84))) = float32(v3*v7/float64(v13) - float64(v10))
	if int32(v5) >= 256 {
		*(*uint16)(unsafe.Pointer(uintptr(result + 124))) = uint16(int16(int32(v5) - 256))
	}
	v8 = int32(*(*uint32)(unsafe.Pointer(uintptr(result + 76))))
	*(*uint32)(unsafe.Pointer(uintptr(result + 64))) = *(*uint32)(unsafe.Pointer(uintptr(result + 72)))
	*(*uint32)(unsafe.Pointer(uintptr(result + 68))) = uint32(v8)
	return result
}
func nox_xxx_damageDefaultProc_4E0B30(a1 int32, a2 int32, a3 int32, a4 int32, a5 int32) int32 {
	var (
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
		v16 float64
		v17 float64
		v18 int32
		v19 int32
		v20 int32
		v21 int32
		v22 int32
		v23 int8
		v24 *uint32
		v25 int32
		v26 int32
		v27 int32
		v28 int32
		v29 func(int32, int32)
		v30 uint8
		v31 float64
		v32 uint16
		v33 int32
		v34 float32
		v35 int32
		v36 float32
		v37 int8
		v38 int32
		v39 int32
		v40 int32
		v41 float32
		v42 float32
		v43 float32
		v44 float32
		v45 [4]int32
		v46 float32
		v47 float32
	)
	if a1 == 0 {
		return 1
	}
	if nox_xxx_testUnitBuffs_4FF350((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 23) != 0 {
		if (int32(uint8(nox_frame_xxx_2598000)) & 3) == 0 {
			nox_xxx_aud_501960(71, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
			return 1
		}
		return 1
	}
	v6 = a5
	if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))&2 != 0 {
		v7 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
		*(*uint32)(unsafe.Pointer(uintptr(v7 + 2188))) = 0
		if v6 == 1 || v6 == 12 || v6 == 7 || v6 == 14 || v6 == 6 {
			*(*uint32)(unsafe.Pointer(uintptr(v7 + 1440))) |= 0x80000
		}
	}
	v8 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 16))))
	if (v8 & 0x8000) == 0 {
		v10 = a2
		if !nox_xxx_CheckGameplayFlags_417DA0(1) {
			v11 = int32(uintptr(unsafe.Pointer(nox_xxx_findParentChainPlayer_4EC580((*nox_object_t)(unsafe.Pointer(uintptr(a2)))))))
			v12 = v11
			if v11 != 0 {
				if int32(*(*uint8)(unsafe.Pointer(uintptr(v11 + 8))))&6 != 0 && nox_xxx_unitIsEnemyTo_5330C0((*nox_object_t)(unsafe.Pointer(uintptr(a1))), (*nox_object_t)(unsafe.Pointer(uintptr(v11)))) == 0 && (a1 != v12 || noxflags.HasGame(noxflags.GameModeQuest)) {
					return 1
				}
			}
		}
		if a2 != 0 && a3 != 0 && nox_xxx_unitIsEnemyTo_5330C0((*nox_object_t)(unsafe.Pointer(uintptr(a1))), (*nox_object_t)(unsafe.Pointer(uintptr(a2)))) == 0 && int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))&6 != 0 && sub_4E1400(a2, (*uint32)(unsafe.Pointer(uintptr(a3)))) != 0 && sub_4E1470(a3) == 0 || int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 16))))&2 != 0 {
			return 1
		}
		if a2 != 0 {
			if nox_xxx_testUnitBuffs_4FF350((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 22) != 0 {
				if int32(*(*uint8)(unsafe.Pointer(uintptr(a2 + 8))))&6 != 0 {
					if a3 != 0 {
						if sub_4E1400(a2, (*uint32)(unsafe.Pointer(uintptr(a3)))) != 0 {
							nox_xxx_aud_501960(135, (*nox_object_t)(unsafe.Pointer(uintptr(a2))), 0, 0)
							nox_xxx_spellBuffOff_4FF5B0((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 22)
							v41 = float32(nox_xxx_gamedataGetFloatTable_419D70(CString("ShockDamage"), 4))
							v13 = int(v41)
							(cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(a2 + 716))), (*func(int32, int32, uint32, int32, int32))(nil)))(a2, a1, 0, v13, 9)
							if int32(*(*uint8)(unsafe.Pointer(uintptr(a2 + 8))))&4 != 0 {
								nox_xxx_playerSetState_4FA020((*nox_object_t)(unsafe.Pointer(uintptr(a2))), 23)
							}
						}
					}
				}
			}
		}
		if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))&2 != 0 {
			v14 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 12))))
			v15 = a5
			if v14&512 != 0 && a5 == 5 {
				return 1
			}
			if v14&1024 != 0 {
				if a5 == 1 || a5 == 12 {
					return 1
				}
				if a5 == 7 {
					a4 /= 2
				LABEL_53:
					v16 = nox_xxx_checkFireProtect_4DFE40((*uint32)(unsafe.Pointer(uintptr(a1))))
					if v16 != 0.0 && (int32(uint8(nox_frame_xxx_2598000))&3) == 0 {
						nox_xxx_aud_501960(104, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
					}
					v46 = float32(v16)
					v42 = float32((1.0 - float64(v46)) * float64(a4))
					a4 = int(v42)
					if a4 == 0 {
						a4 = 1
					}
				LABEL_58:
					if v15 == 9 || v15 == 17 {
						v17 = nox_xxx_checkElectrProtect_4DFF40((*uint32)(unsafe.Pointer(uintptr(a1))))
						if v17 != 0.0 && (int32(uint8(nox_frame_xxx_2598000))&3) == 0 {
							nox_xxx_aud_501960(108, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
						}
						v47 = float32(v17)
						v43 = float32((1.0 - float64(v47)) * float64(a4))
						a4 = int(v43)
						if a4 == 0 {
							a4 = 1
						}
						v18 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 8))))
						if v18&4 != 0 {
							*(*uint16)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))) + 160))) = 2
						} else if v18&2 != 0 && int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 12))))&16 != 0 {
							*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))) + 2094))) = 2
						}
					}
					if v10 == 0 {
						*(*uint32)(unsafe.Pointer(uintptr(a1 + 528))) = 0
						*(*uint32)(unsafe.Pointer(uintptr(a1 + 532))) = 0
						if v15 == 12 {
							nox_xxx_spellBuffOff_4FF5B0((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0)
						}
						v19 = a3
					LABEL_87:
						v22 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 8))))
						if v22&4 != 0 || v22&2 != 0 && int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 12))))&16 != 0 {
							nox_xxx_itemApplyDefendEffect2_4E1320(a1, v10, v19, &a4, v15)
						}
						if v19 != 0 {
							*(*uint32)(unsafe.Pointer(uintptr(a1 + 520))) = uint32(v19)
						} else {
							*(*uint32)(unsafe.Pointer(uintptr(a1 + 520))) = uint32(v10)
						}
						v23 = int8(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))
						*(*uint32)(unsafe.Pointer(uintptr(a1 + 524))) = uint32(v15)
						*(*uint32)(unsafe.Pointer(uintptr(a1 + 536))) = nox_frame_xxx_2598000
						if int32(v23)&2 != 0 {
							v24 = *(**uint32)(unsafe.Pointer(uintptr(a1 + 748)))
							v25 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v24), unsafe.Sizeof(uint32(0))*360)))
							v26 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v24), unsafe.Sizeof(uint32(0))*547)))
							*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v25))), 1)) |= 2
							*(*uint32)(unsafe.Add(unsafe.Pointer(v24), unsafe.Sizeof(uint32(0))*360)) = uint32(v25)
							if v26 == 0 {
								*(*uint32)(unsafe.Add(unsafe.Pointer(v24), unsafe.Sizeof(uint32(0))*547)) = 2
								*(*uint32)(unsafe.Add(unsafe.Pointer(v24), unsafe.Sizeof(uint32(0))*546)) = uint32(v15)
							}
						}
						if v19 != 0 && *(*uint32)(unsafe.Pointer(uintptr(v19 + 8)))&0x1001000 != 0 {
							nox_xxx_itemApplyPreDamageEffect_4E13B0(a1, v10, v19, int32(uintptr(unsafe.Pointer(&a4))))
						}
						if a1 != v19 || (*(*uint32)(unsafe.Pointer(uintptr(a1 + 8)))&0x1001000) == 0 {
							if v10 == 0 || (int32(*(*uint8)(unsafe.Pointer(uintptr(v10 + 8))))&2) == 0 || *(*uint32)(unsafe.Pointer(uintptr(v10 + 748))) == 0 || (func() int32 {
								v27 = nox_xxx_monsterGetSoundSet_424300(v10)
								return v27
							}()) == 0 || (func() int32 {
								v28 = int32(*(*uint32)(unsafe.Pointer(uintptr(v27 + 32))))
								return v28
							}()) == 0 || nox_xxx_getSevenDwords3_501940(v28) <= 0 {
								v29 = cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(a1 + 720))), (*func(int32, int32))(nil))
								if v29 != nil {
									if v19 != 0 {
										v29(a1, v19)
									} else {
										v29(a1, v10)
									}
								} else if v19 != 0 {
									nox_xxx_soundDefaultDamageSound_532E20(a1, v19)
								} else {
									nox_xxx_soundDefaultDamageSound_532E20(a1, v10)
								}
							}
						}
						if v10 != 0 && int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))&6 != 0 && nox_xxx_testUnitBuffs_4FF350((*nox_object_t)(unsafe.Pointer(uintptr(v10))), 13) != 0 {
							nox_xxx_aud_501960(163, (*nox_object_t)(unsafe.Pointer(uintptr(a3))), 0, 0)
							v30 = uint8(nox_xxx_buffGetPower_4FF570(v10, 13))
							v31 = nox_xxx_gamedataGetFloatTable_419D70(CString("VampirismCoeff"), int32(uint32(v30)-1))
							v44 = float32(v31 * float64(a4))
							v32 = uint16(int16(int(v44)))
							if int32(v32) < 1 {
								v32 = 1
							}
							nox_xxx_unitAdjustHP_4EE460((*nox_object_t)(unsafe.Pointer(uintptr(v10))), int32(v32))
							v45[0] = int(*(*float32)(unsafe.Pointer(uintptr(v10 + 56))))
							v33 = int(*(*float32)(unsafe.Pointer(uintptr(v10 + 60))))
							v34 = *(*float32)(unsafe.Pointer(uintptr(a1 + 56)))
							v45[1] = v33
							v35 = int(v34)
							v36 = *(*float32)(unsafe.Pointer(uintptr(a1 + 60)))
							v45[2] = v35
							v45[3] = int(v36)
							nox_xxx_netSendVampFx_523270(-94, (*int16)(unsafe.Pointer(&v45[0])), int16(v32))
						}
						nox_xxx_gameballOnPlayerDamage_4E1230(v10, a1, a4)
						if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))&4 != 0 {
							if a4 >= 20 {
								v37 = int8(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))) + 88))))
								if int32(v37) != 1 && int32(v37) != 15 {
									nox_xxx_playerSetState_4FA020((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 30)
								}
							}
						}
						if noxflags.HasGame(noxflags.GameModeCoop | noxflags.GameModeQuest) {
							sub_4FB050(v10, a1, &a4)
						}
						if v10 == 0 {
							goto LABEL_137
						}
						if int32(*(*uint8)(unsafe.Pointer(uintptr(v10 + 8))))&2 != 0 {
							v38 = v10
						} else {
							v39 = int32(*(*uint32)(unsafe.Pointer(uintptr(v10 + 508))))
							if v39 == 0 || (int32(*(*uint8)(unsafe.Pointer(uintptr(v39 + 8))))&2) == 0 {
							LABEL_137:
								if nox_xxx_testUnitBuffs_4FF350((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 26) != 0 && a5 != 5 {
									v40 = a3
									if a3 == 0 {
										v40 = v10
									}
									if a5 != 15 || v10 != a1 {
										nox_xxx_unitShieldReduceDamage_52F710(a1, &a4, a5, v40)
									}
									if a4 == 0 {
										return 0
									}
								}
								nox_xxx_unitDamageClear_4EE5E0((*nox_object_t)(unsafe.Pointer(uintptr(a1))), a4)
								return 1
							}
							v38 = int32(*(*uint32)(unsafe.Pointer(uintptr(v10 + 508))))
						}
						if v38 != 0 && nox_xxx_unitIsEnemyTo_5330C0((*nox_object_t)(unsafe.Pointer(uintptr(a1))), (*nox_object_t)(unsafe.Pointer(uintptr(v38)))) != 0 {
							sub_532880(v38)
						}
						goto LABEL_137
					}
					v19 = a3
					if a3 != 0 {
						if *(*uint32)(unsafe.Pointer(uintptr(a3 + 8)))&0x1001000 != 0 {
							*(*uint32)(unsafe.Pointer(uintptr(a1 + 528))) = *(*uint32)(unsafe.Pointer(uintptr(v10 + 72)))
							*(*uint32)(unsafe.Pointer(uintptr(a1 + 532))) = *(*uint32)(unsafe.Pointer(uintptr(v10 + 76)))
						} else {
							*(*uint32)(unsafe.Pointer(uintptr(a1 + 528))) = *(*uint32)(unsafe.Pointer(uintptr(a3 + 72)))
							*(*uint32)(unsafe.Pointer(uintptr(a1 + 532))) = *(*uint32)(unsafe.Pointer(uintptr(a3 + 76)))
						}
						if a3 == v10 || (int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))&2) == 0 {
							goto LABEL_83
						}
						v20 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
						*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v21))), unsafe.Sizeof(uint16(0))*1)) = 0
						*(*uint32)(unsafe.Pointer(uintptr(v20 + 2188))) = 1
						*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v21))), unsafe.Sizeof(uint16(0))*0)) = *(*uint16)(unsafe.Pointer(uintptr(a3 + 4)))
					} else {
						*(*uint32)(unsafe.Pointer(uintptr(a1 + 528))) = *(*uint32)(unsafe.Pointer(uintptr(v10 + 72)))
						*(*uint32)(unsafe.Pointer(uintptr(a1 + 532))) = *(*uint32)(unsafe.Pointer(uintptr(v10 + 76)))
						if v15 != 10 && v15 != 2 || (int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))&2) == 0 {
							goto LABEL_83
						}
						v20 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
						*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v21))), unsafe.Sizeof(uint16(0))*1)) = 0
						*(*uint32)(unsafe.Pointer(uintptr(v20 + 2188))) = 1
						*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v21))), unsafe.Sizeof(uint16(0))*0)) = *(*uint16)(unsafe.Pointer(uintptr(v10 + 4)))
					}
					*(*uint32)(unsafe.Pointer(uintptr(v20 + 2184))) = uint32(v21)
				LABEL_83:
					nox_xxx_spellBuffOff_4FF5B0((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0)
					goto LABEL_87
				}
			} else if v14&2048 != 0 {
				if a5 == 9 {
					return 1
				}
				if a5 == 17 {
					return 1
				}
			}
		} else {
			v15 = a5
		}
		if v15 != 1 && v15 != 12 && v15 != 7 {
			goto LABEL_58
		}
		goto LABEL_53
	}
	if nox_xxx_unitIsZombie_534A40(a1) == 0 {
		return 1
	}
	v9 = a3
	if a3 == 0 {
		v9 = a2
	}
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 520))) = uint32(v9)
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 524))) = uint32(v6)
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 536))) = nox_frame_xxx_2598000
	return 1
}
func nox_xxx_gameballOnPlayerDamage_4E1230(a1 int32, a2 int32, a3 int32) {
	var (
		v3 int32
		v4 int32
		v5 *byte
	)
	if int32(*(*uint8)(unsafe.Pointer(uintptr(a2 + 8))))&4 != 0 && a3 >= 30 {
		v3 = int32(*memmap.PtrUint32(6112660, 1563316))
		if *memmap.PtrUint32(6112660, 1563316) == 0 {
			v3 = nox_xxx_getNameId_4E3AA0(CString("GameBall"))
			*memmap.PtrUint32(6112660, 1563316) = uint32(v3)
		}
		v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 516))))
		if v4 != 0 {
			for int32(*(*uint16)(unsafe.Pointer(uintptr(v4 + 4)))) != v3 {
				v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(v4 + 512))))
				if v4 == 0 {
					return
				}
			}
			*(*uint32)(unsafe.Pointer(uintptr(v4 + 16))) &= 0xFFFFFFBF
			nox_xxx_objectApplyForce_52DF80((*float32)(unsafe.Pointer(uintptr(a2+56))), (*nox_object_t)(unsafe.Pointer(uintptr(v4))), 30.0)
			nox_xxx_unitClearOwner_4EC300((*nox_object_t)(unsafe.Pointer(uintptr(v4))))
			sub_4EB9B0(v4, a2)
			if nox_xxx_servObjectHasTeam_419130(v4+48) != 0 {
				v5 = (*byte)(unsafe.Pointer(nox_xxx_clientGetTeamColor_418AB0(int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 52)))))))
				if v5 != nil {
					sub_4196D0(v4+48, int32(uintptr(unsafe.Pointer(v5))), int32(*(*uint32)(unsafe.Pointer(uintptr(v4 + 36)))), 0)
				}
			} else {
				nox_xxx_createAtImpl_4191D0(*(*uint8)(unsafe.Pointer(uintptr(a1 + 52))), unsafe.Pointer(uintptr(v4+48)), 1, int32(*(*uint32)(unsafe.Pointer(uintptr(v4 + 36)))), 0)
			}
			nox_xxx_aud_501960(926, (*nox_object_t)(unsafe.Pointer(uintptr(a2))), 0, 0)
		}
	}
}
func nox_xxx_itemApplyDefendEffect2_4E1320(a1 int32, a2 int32, a3 int32, a4 *int32, a5 int32) int32 {
	var (
		result int32
		v6     *uint32
		v7     int32
		v8     *int32
		v9     int32
		v10    int2
		v11    int32
	)
	result = a1
	v6 = *(**uint32)(unsafe.Pointer(uintptr(a1 + 504)))
	if v6 != nil {
		v7 = a5
		for {
			result = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*4)))
			if result&256 != 0 {
				v11 = 2
				v8 = (*int32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*173)) + 8)))
				for {
					v9 = *v8
					if *v8 != 0 {
						if *(*uint32)(unsafe.Pointer(uintptr(v9 + 76))) != 0 {
							v10.field_0 = *a4
							v10.field_4 = v7
							(cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(v9 + 76))), (*func(int32, *uint32, int32, int32, int32, *int2))(nil)))(v9, v6, a1, a3, a2, &v10)
							*a4 = v10.field_0
						}
					}
					v8 = (*int32)(unsafe.Add(unsafe.Pointer(v8), unsafe.Sizeof(int32(0))*1))
					result = func() int32 {
						p := &v11
						*p--
						return *p
					}()
					if v11 == 0 {
						break
					}
				}
			}
			v6 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*124)))))
			if v6 == nil {
				break
			}
		}
	}
	return result
}
func nox_xxx_itemApplyPreDamageEffect_4E13B0(a1 int32, a2 int32, a3 int32, a4 int32) int32 {
	var (
		v4     int32
		v5     *int32
		v6     int32
		v7     func(int32, int32, int32, int32, int32)
		result int32
		v9     int32
	)
	v4 = a3
	v9 = 4
	v5 = *(**int32)(unsafe.Pointer(uintptr(v4 + 692)))
	for {
		v6 = *v5
		if *v5 != 0 {
			v7 = cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(v6 + 64))), (*func(int32, int32, int32, int32, int32))(nil))
			if v7 != nil {
				v7(v6, v4, a2, a1, a4)
			}
		}
		v5 = (*int32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(int32(0))*1))
		result = func() int32 {
			p := &v9
			*p--
			return *p
		}()
		if v9 == 0 {
			break
		}
	}
	return result
}
func sub_4E1400(a1 int32, a2 *uint32) int32 {
	var (
		v2     int32
		result int32
		v4     int32
	)
	if a2 != nil {
		v4 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*2)))
		if v4&4096 != 0 {
			if (*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*3)) & 0x47F0000) == 0 {
				return 1
			}
			if int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*184)) + 96))))&2 != 0 {
				return 1
			}
		} else if uint32(v4)&0x1000000 != 0 && (*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*3))&75432190) == 0 {
			return 1
		}
		result = (int32(uint8(int8(v4))) >> 1) & 1
	} else {
		v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 8))))
		if v2&4 != 0 {
			return 1
		}
		result = bool2int(v2&2 != 0 && int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 12))))&16 != 0)
	}
	return result
}
func sub_4E1470(a1 int32) int32 {
	var (
		v1     int32
		result int32
	)
	result = 0
	if a1 != 0 {
		if *(*uint32)(unsafe.Pointer(uintptr(a1 + 8)))&0x1000000 != 0 {
			v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 12))))
			if v1&0x4000 != 0 {
				result = 1
			}
		}
	}
	return result
}
func sub_4E14A0() int32 {
	return 0
}
func sub_4E14B0(a1 int32, a2 int32, a3 int32, a4 int32, a5 int32) int32 {
	var result int32
	if a1 != 0 && *(*uint32)(unsafe.Pointer(uintptr(a1 + 8)))&0x1001000 != 0 && (*(*uint32)(unsafe.Pointer(uintptr(a1 + 492))) != 0 || a5 == 12) {
		result = nox_xxx_damageDefaultProc_4E0B30(a1, a2, a3, a4, a5)
	} else {
		result = 0
	}
	return result
}
func nox_xxx_damageArmor_4E1500(a1 int32, a2 int32, a3 int32, a4 int32, a5 int32) int32 {
	var (
		v5     int32
		result int32
	)
	if *(*uint32)(unsafe.Pointer(uintptr(a1 + 8)))&0x2000000 != 0 && (*(*uint32)(unsafe.Pointer(uintptr(a1 + 492))) != 0 || a5 == 12) && (func() int32 {
		if a5 != 2 || (int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 24))))&16) == 0 {
			v5 = a4
		} else {
			v5 = a4 * 2
		}
		return v5
	}()) != 0 {
		result = nox_xxx_damageDefaultProc_4E0B30(a1, a2, a3, v5, a5)
	} else {
		result = 0
	}
	return result
}
func nox_xxx_playerDamageWeapon_4E1560(a1 int32, a2 int32, a3 int32, a4 int32, a5 float32, a6 int32) {
	var (
		v6  *float32
		v7  int32
		v8  func(int32, int32, int32, int32, int32, *float32)
		v9  int32
		v10 uint16
		v11 uint16
	)
	if a1 != 0 {
		if *(*uint32)(unsafe.Pointer(uintptr(a1 + 556))) != 0 {
			v6 = *(**float32)(unsafe.Pointer(uintptr(a1 + 748)))
			if (*(*uint32)(unsafe.Pointer(uintptr(a1 + 8)))&0x1000000) == 0 || (*(*uint32)(unsafe.Pointer(uintptr(a1 + 12)))&0x7800000) == 0 {
				v7 = int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 692))) + 4))))
				if v7 != 0 {
					v8 = cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(v7 + 76))), (*func(int32, int32, int32, int32, int32, *float32))(nil))
					if v8 != nil {
						v8(v7, a1, a2, a4, a3, &a5)
					}
				}
				a5 = a5 + *v6
				v9 = int(a5)
				*v6 = float32(float64(a5) - float64(v9))
				if v9 > 0 {
					v10 = **(**uint16)(unsafe.Pointer(uintptr(a1 + 556)))
					(cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(a1 + 716))), (*func(int32, int32, int32, int32, int32))(nil)))(a1, a3, a4, v9, a6)
					if a2 != 0 {
						if int32(*(*uint8)(unsafe.Pointer(uintptr(a2 + 8))))&4 != 0 {
							v11 = **(**uint16)(unsafe.Pointer(uintptr(a1 + 556)))
							if int32(v10) != int32(v11) {
								nox_xxx_itemDestroyed_4E1650(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a2 + 748))) + 276))) + 2064)))), (*uint32)(unsafe.Pointer(uintptr(a1))), v10, v11)
							}
						}
					}
				}
			}
		}
	}
}
func nox_xxx_itemDestroyed_4E1650(a1 int32, a2 *uint32, a3 uint16, a4 uint16) int32 {
	var (
		result int32
		v5     int32
	)
	result = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*139)))
	if result != 0 {
		if noxflags.HasGame(noxflags.GameModeCoop) {
			result = nox_xxx_itemReportHealth_4D87A0(a1, (*nox_object_t)(unsafe.Pointer(a2)))
		} else {
			v5 = sub_57B190(a3, *(*uint16)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*139)) + 4))))
			result = sub_57B190(a4, *(*uint16)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*139)) + 4))))
			if v5 != result {
				result = nox_xxx_itemReportHealth_4D87A0(a1, (*nox_object_t)(unsafe.Pointer(a2)))
			}
		}
	}
	return result
}
func nox_xxx_equipDamage_4E16D0(a1 int32, a2 int32, a3 int32, a4 int32, a5 float32, a6 int32) {
	var (
		v6  *float32
		v7  int32
		v8  func(int32, int32, int32, int32, int32, *float32)
		v9  int32
		v10 uint16
		v11 uint16
	)
	if a1 != 0 && *(*uint32)(unsafe.Pointer(uintptr(a1 + 556))) != 0 {
		v6 = *(**float32)(unsafe.Pointer(uintptr(a1 + 748)))
		v7 = int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 692))) + 4))))
		if v7 != 0 {
			v8 = cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(v7 + 76))), (*func(int32, int32, int32, int32, int32, *float32))(nil))
			if v8 != nil {
				v8(v7, a1, a2, a4, a3, &a5)
			}
		}
		a5 = a5 + *v6
		v9 = int(a5)
		*v6 = float32(float64(a5) - float64(v9))
		if v9 > 0 {
			v10 = **(**uint16)(unsafe.Pointer(uintptr(a1 + 556)))
			(cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(a1 + 716))), (*func(int32, int32, int32, int32, int32))(nil)))(a1, a3, a4, v9, a6)
			if int32(*(*uint8)(unsafe.Pointer(uintptr(a2 + 8))))&4 != 0 {
				v11 = **(**uint16)(unsafe.Pointer(uintptr(a1 + 556)))
				if int32(v10) != int32(v11) {
					nox_xxx_itemDestroyed_4E1650(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a2 + 748))) + 276))) + 2064)))), (*uint32)(unsafe.Pointer(uintptr(a1))), v10, v11)
				}
			}
		}
	}
}
