package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	noxflags "github.com/noxworld-dev/opennox/v1/common/flags"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"image"
	"math"
	"unsafe"
)

func nox_xxx_clientDrawAll_436100_draw_A() {
	if sub_436550() == 0 {
		noxflags.UnsetEngine(nox_engine_flag(NOX_ENGINE_FLAG_9))
	} else {
		noxflags.SetEngine(nox_engine_flag(NOX_ENGINE_FLAG_9))
	}
	if *memmap.PtrUint32(6112660, 814540) == 0 {
		*memmap.PtrUint32(6112660, 814540) = uint32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg("MenuSystemBG"))))
	}
	if nox_common_getEngineFlag(nox_engine_flag(NOX_ENGINE_FLAG_9)) {
		var v10 unsafe.Pointer = guiFontPtrByName(CString("large"))
		nox_client_drawImageAt_47D2C0((*nox_video_bag_image_t)(unsafe.Pointer(uintptr(*memmap.PtrInt32(6112660, 814540)))), 0, 0)
		var v11 *libc.WChar = strMan.GetStringInFile("InProgress", "C:\\NoxPost\\src\\client\\System\\client.c")
		var v22 int32 = 0
		nox_xxx_drawGetStringSize_43F840(v10, v11, &v22, nil, 0)
		nox_xxx_drawSetTextColor_434390(int32(nox_color_white_2523948))
		var v21 int32 = nox_win_height / 2
		var v20 int32 = (nox_win_width - v22) / 2
		var v12 *libc.WChar = strMan.GetStringInFile("InProgress", "C:\\NoxPost\\src\\client\\System\\client.c")
		noxrend.DrawString(v10, v12, image.Pt(v20, v21))
	}
}
func nox_xxx_clientDrawAll_436100_draw_B() {
	var v25 [128]libc.WChar
	nox_wcscpy(&v25[0], (*libc.WChar)(memmap.PtrOff(6112660, 811376)))
	var v14 int32 = (nox_win_width - 310) / 2
	var v16 int32 = (nox_win_height - 200) / 2
	var v15 *int32 = memmap.PtrInt32(6112660, *memmap.PtrUint32(6112660, 811060)*4+0xC6370)
	if *v15 == 0 {
		*v15 = int32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg(*(**byte)(memmap.PtrOff(0x587000, *memmap.PtrUint32(6112660, 811060)*4+0x14ED0))))))
	}
	nox_client_drawImageAt_47D2C0((*nox_video_bag_image_t)(unsafe.Pointer(uintptr(*v15))), v14, v16)
	var v22 int32 = 0
	var v23 int32 = 0
	nox_xxx_drawGetStringSize_43F840(nil, &v25[0], &v22, &v23, 220)
	var v17 int32 = v14 + 45
	var v18 int32 = v16 + (49-v23)/2 + 143
	nox_client_drawSetColor_434460(int32(nox_color_white_2523948))
	for tok := (*libc.WChar)(nox_wcstok(&v25[0], libc.CWString("\n\r"))); tok != nil; tok = nox_wcstok(nil, libc.CWString("\n\r")) {
		nox_xxx_drawSetTextColor_434390(int32(nox_color_white_2523948))
		nox_xxx_drawGetStringSize_43F840(nil, tok, &v22, nil, 0)
		nox_xxx_drawStringWrap_43FAF0(nil, tok, v17+(220-v22)/2, v18, 220, 0)
		v18 += nox_xxx_guiFontHeightMB_43F320(nil)
	}
}
func nox_xxx_drawBandwith_436970(a1 int32) int32 {
	var (
		v1     int32
		v2     *int16
		v3     int32
		v4     int32
		v5     int32
		v6     int32
		result int32
		v8     int32
		v9     int32
		v10    int32
	)
	v9 = nox_win_height - 1
	v1 = nox_win_height - 31
	nox_xxx_drawSetTextColor_434390(int32(nox_color_white_2523948))
	v8 = v1 - nox_xxx_guiFontHeightMB_43F320(nil)
	v2 = (*int16)(unsafe.Pointer(strMan.GetStringInFile("Bandwidth", "C:\\NoxPost\\src\\client\\System\\client.c")))
	noxrend.DrawString(nil, (*libc.WChar)(unsafe.Pointer(v2)), image.Pt(0, v8))
	nox_client_drawSetColor_434460(*memmap.PtrInt32(0x85B3FC, 956))
	nox_client_drawBorderLines_49CC70(0, v1, nox_win_width, 31)
	v3 = int32(*memmap.PtrUint32(6112660, 812452))
	*memmap.PtrUint32(6112660, *memmap.PtrUint32(6112660, 812452)*4+811940) = uint32(a1 / 4)
	if a1/4 > 30 {
		*memmap.PtrUint32(6112660, uint32(v3*4)+811940) = 30
	}
	v4 = (v3 + 1) % 128
	*memmap.PtrUint32(6112660, 812452) = uint32((v3 + 1) % 128)
	v5 = 0
	v6 = nox_win_width / 128
	nox_client_drawSetColor_434460(int32(nox_color_yellow_2589772))
	v10 = math.MaxInt8
	for {
		nox_client_drawAddPoint_49F500(v5, int32(uint32(v9)-*memmap.PtrUint32(6112660, uint32(v4*4)+811940)))
		v5 += v6
		nox_client_drawAddPoint_49F500(v5, int32(uint32(v9)-*memmap.PtrUint32(6112660, uint32(((v4+1)%128)*4)+811940)))
		nox_client_drawLineFromPoints_49E4B0()
		v4 = (v4 + 1) % 128
		result = func() int32 {
			p := &v10
			*p--
			return *p
		}()
		if v10 == 0 {
			break
		}
	}
	return result
}
func sub_436AA0(a1 int32) int32 {
	var (
		v1     int32
		v2     int32
		v3     *int16
		v4     int32
		v5     int32
		v6     int32
		v7     int32
		v8     int32
		v9     int32
		v10    int32
		v11    int32
		v12    bool
		result int32
		v14    int32
		v15    int32
		v16    int32
	)
	v1 = nox_xxx_guiFontHeightMB_43F320(nil)
	v2 = v1 + 30
	nox_xxx_drawSetTextColor_434390(int32(nox_color_white_2523948))
	v14 = v1 - nox_xxx_guiFontHeightMB_43F320(nil)
	v3 = (*int16)(unsafe.Pointer(strMan.GetStringInFile("FPS", "C:\\NoxPost\\src\\client\\System\\client.c")))
	noxrend.DrawString(nil, (*libc.WChar)(unsafe.Pointer(v3)), image.Pt(0, v14))
	nox_client_drawSetColor_434460(*memmap.PtrInt32(0x85B3FC, 956))
	v4 = v1
	v5 = 4
	for {
		nox_client_drawAddPoint_49F500(0, v4)
		nox_client_drawAddPoint_49F500(nox_win_width, v4)
		nox_client_drawLineFromPoints_49E4B0()
		v4 += 10
		v5--
		if v5 == 0 {
			break
		}
	}
	nox_client_drawAddPoint_49F500(0, v1)
	nox_client_drawAddPoint_49F500(0, v1+30)
	nox_client_drawLineFromPoints_49E4B0()
	nox_client_drawAddPoint_49F500(nox_win_width-1, v1)
	nox_client_drawAddPoint_49F500(nox_win_width-1, v1+30)
	nox_client_drawLineFromPoints_49E4B0()
	v6 = int32(*memmap.PtrUint32(6112660, 812968) + 1)
	*memmap.PtrUint32(6112660, *memmap.PtrUint32(6112660, 812968)*4+0xC65A8) = uint32(a1 * 10 / 10)
	*memmap.PtrUint32(6112660, 812968) = uint32(v6 % 128)
	v7 = v6 % 128
	v16 = math.MaxInt8
	v8 = nox_win_width / 128
	v9 = 0
	for {
		v15 = (v7 + 1) % 128
		v10 = int32(*memmap.PtrUint32(6112660, uint32(v7*4)+0xC65A8))
		if v10 >= 10 {
			v12 = v10 < 20
			v11 = int32(nox_color_yellow_2589772)
			if !v12 {
				v11 = int32(*memmap.PtrUint32(0x8531A0, 2572))
			}
		} else {
			v11 = int32(*memmap.PtrUint32(0x85B3FC, 940))
		}
		nox_client_drawSetColor_434460(v11)
		nox_client_drawAddPoint_49F500(v9, int32(uint32(v2)-*memmap.PtrUint32(6112660, uint32(v7*4)+0xC65A8)))
		v7 = (v7 + 1) % 128
		v9 += v8
		nox_client_drawAddPoint_49F500(v9, int32(uint32(v2)-*memmap.PtrUint32(6112660, uint32(v15*4)+0xC65A8)))
		nox_client_drawLineFromPoints_49E4B0()
		result = func() int32 {
			p := &v16
			*p--
			return *p
		}()
		if v16 == 0 {
			break
		}
	}
	return result
}
func nox_xxx_drawTimingMB_436C40() int32 {
	var (
		v0     *int16
		v1     int32
		v2     int32
		v3     int32
		v4     int32
		v5     int32
		v6     int32
		v7     int32
		v8     int32
		v9     bool
		result int32
		v11    int32
		v12    int32
		v13    int32
	)
	nox_xxx_drawSetTextColor_434390(int32(nox_color_white_2523948))
	v11 = 60 - nox_xxx_guiFontHeightMB_43F320(nil)
	v0 = (*int16)(unsafe.Pointer(strMan.GetStringInFile("CSTiming", "C:\\NoxPost\\src\\client\\System\\client.c")))
	noxrend.DrawString(nil, (*libc.WChar)(unsafe.Pointer(v0)), image.Pt(0, v11))
	nox_client_drawSetColor_434460(*memmap.PtrInt32(0x85B3FC, 956))
	nox_client_drawBorderLines_49CC70(0, 60, nox_win_width, 31)
	v1 = int32(*memmap.PtrUint32(6112660, 813996))
	v2 = int32(*memmap.PtrUint32(6112660, 811928) * 30 / 100)
	*memmap.PtrUint32(6112660, *memmap.PtrUint32(6112660, 813996)*4+0xC67AC) = uint32(v2)
	if v2 > 30 {
		*memmap.PtrUint32(6112660, uint32(v1*4)+0xC67AC) = 30
	}
	v3 = int32(*memmap.PtrUint32(6112660, 811936) * 30 / 100)
	*memmap.PtrUint32(6112660, uint32(v1*4)+0xC69AC) = uint32(v3)
	if v3 > 30 {
		*memmap.PtrUint32(6112660, uint32(v1*4)+0xC69AC) = 30
	}
	v4 = (v1 + 1) % 128
	*memmap.PtrUint32(6112660, 813996) = uint32((v1 + 1) % 128)
	v12 = math.MaxInt8
	v5 = nox_win_width / 128
	v13 = nox_win_width / 128
	v6 = 0
	for {
		v7 = (v4 + 1) % 128
		nox_client_drawSetColor_434460(int32(nox_color_blue_2650684))
		nox_client_drawAddPoint_49F500(v6, int32(90-*memmap.PtrUint32(6112660, uint32(v4*4)+0xC67AC)))
		v8 = v6 + v5
		nox_client_drawAddPoint_49F500(v8, int32(90-*memmap.PtrUint32(6112660, uint32(v7*4)+0xC67AC)))
		nox_client_drawLineFromPoints_49E4B0()
		nox_client_drawSetColor_434460(*memmap.PtrInt32(0x85B3FC, 940))
		nox_client_drawAddPoint_49F500(v6, int32(90-*memmap.PtrUint32(6112660, uint32(v4*4)+0xC69AC)))
		nox_client_drawAddPoint_49F500(v8, int32(90-*memmap.PtrUint32(6112660, uint32(v7*4)+0xC69AC)))
		nox_client_drawLineFromPoints_49E4B0()
		v6 = v8
		result = v12 - 1
		v9 = v12 == 1
		v4 = (v4 + 1) % 128
		v12--
		if v9 {
			break
		}
		v5 = v13
	}
	return result
}
func nox_xxx_drawPing_436DF0(a1 int32) int32 {
	var (
		v1     int32
		v2     *int16
		v3     int32
		v4     int32
		v5     int32
		v6     int32
		v7     int32
		v8     int32
		v9     int32
		v10    bool
		v11    bool
		result int32
		v13    int32
		v14    int32
		v15    int32
	)
	v1 = nox_win_height - 80
	v14 = nox_win_height - 80 + 30
	nox_xxx_drawSetTextColor_434390(int32(nox_color_white_2523948))
	v13 = v1 - nox_xxx_guiFontHeightMB_43F320(nil)
	v2 = (*int16)(unsafe.Pointer(strMan.GetStringInFile("Ping", "C:\\NoxPost\\src\\client\\System\\client.c")))
	noxrend.DrawString(nil, (*libc.WChar)(unsafe.Pointer(v2)), image.Pt(0, v13))
	nox_client_drawSetColor_434460(*memmap.PtrInt32(0x85B3FC, 956))
	nox_client_drawBorderLines_49CC70(0, v1, nox_win_width, 31)
	v3 = int32(*memmap.PtrUint32(6112660, 814512))
	v4 = a1 * 30 / 500
	*memmap.PtrUint32(6112660, *memmap.PtrUint32(6112660, 814512)*4+814000) = uint32(v4)
	if v4 > 30 {
		*memmap.PtrUint32(6112660, uint32(v3*4)+814000) = 30
	}
	v5 = (v3 + 1) % 128
	*memmap.PtrUint32(6112660, 814512) = uint32((v3 + 1) % 128)
	v15 = math.MaxInt8
	v6 = nox_win_width / 128
	v7 = 0
	for {
		v8 = int32(*memmap.PtrUint32(6112660, uint32(v5*4)+814000))
		if v8 >= 100 {
			v10 = v8 < 350
			v9 = int32(nox_color_yellow_2589772)
			if !v10 {
				v9 = int32(*memmap.PtrUint32(0x85B3FC, 940))
			}
		} else {
			v9 = int32(*memmap.PtrUint32(0x8531A0, 2572))
		}
		nox_client_drawSetColor_434460(v9)
		nox_client_drawAddPoint_49F500(v7, int32(uint32(v14)-*memmap.PtrUint32(6112660, uint32(v5*4)+814000)))
		nox_client_drawAddPoint_49F500(v7+v6, int32(uint32(v14)-*memmap.PtrUint32(6112660, uint32(((v5+1)%128)*4)+814000)))
		nox_client_drawLineFromPoints_49E4B0()
		v7 += v6
		result = v15 - 1
		v11 = v15 == 1
		v5 = (v5 + 1) % 128
		v15--
		if v11 {
			break
		}
	}
	return result
}
func sub_436F50() int32 {
	var (
		v0     int32
		v2     int32
		v3     int32
		v4     *byte
		result int32
		v6     int32
		v7     *libc.WChar
		v8     int32
		v9     *libc.WChar
	)
	v0 = nox_xxx_guiFontHeightMB_43F320(nil)
	var rdr *nox_draw_viewport_t = nox_draw_getViewport_437250()
	v2 = rdr.x1 + 10
	v3 = rdr.y1 + 90
	nox_xxx_drawSetTextColor_434390(int32(nox_color_white_2523948))
	v4 = nox_server_currentMapGetFilename_409B30()
	nox_swprintf((*libc.WChar)(memmap.PtrOff(6112660, 811120)), libc.CWString("%S"), v4)
	noxrend.DrawString(nil, (*libc.WChar)(unsafe.Pointer(memmap.PtrInt16(6112660, 811120))), image.Pt(v2, v3))
	result = int32(*memmap.PtrUint32(0x852978, 8))
	v6 = v0 + v3
	if *memmap.PtrUint32(0x852978, 8) != 0 {
		if *memmap.PtrUint32(0x8531A0, 2576) != 0 {
			nox_swprintf((*libc.WChar)(memmap.PtrOff(6112660, 811120)), libc.CWString("X:%d\tY:%d"), *(*uint32)(unsafe.Pointer(uintptr(*memmap.PtrUint32(0x852978, 8) + 12))), *(*uint32)(unsafe.Pointer(uintptr(*memmap.PtrUint32(0x852978, 8) + 16))))
			noxrend.DrawString(nil, (*libc.WChar)(unsafe.Pointer(memmap.PtrInt16(6112660, 811120))), image.Pt(v2, v6))
			v9 = strMan.GetStringInFile(*(**byte)(memmap.PtrOff(0x587000, uint32(int32(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(0x8531A0, 2576) + 2251))))*4+0x7310))), "C:\\NoxPost\\src\\client\\System\\client.c")
			v8 = int32(*(*byte)(unsafe.Pointer(uintptr(*memmap.PtrUint32(0x8531A0, 2576) + 3684))))
			v7 = strMan.GetStringInFile("PlayerInfo", "C:\\NoxPost\\src\\client\\System\\client.c")
			nox_swprintf((*libc.WChar)(memmap.PtrOff(6112660, 811120)), v7, v8, v9)
			result = noxrend.DrawString(nil, (*libc.WChar)(unsafe.Pointer(memmap.PtrInt16(6112660, 811120))), image.Pt(v2, v0+v6))
		}
	}
	return result
}
func nox_xxx_getRandomName_4358A0() *byte {
	var (
		v0 int32
		v1 **byte
		v2 *byte
		v4 [34]*byte
	)
	v0 = int32(*memmap.PtrUint32(6112660, 814516))
	v4[0] = CString("Dweezle")
	v4[1] = CString("Glork")
	v4[2] = CString("Floogle")
	v4[3] = CString("Goombah")
	v4[4] = CString("Kraun")
	v4[5] = CString("Kloog")
	v4[6] = CString("Zurg")
	v4[7] = CString("Darg")
	v4[8] = CString("Arfingle")
	v4[9] = CString("Buurl")
	v4[10] = CString("Gurgin")
	v4[11] = CString("Grok")
	v4[12] = CString("Hurlong")
	v4[13] = CString("Luric")
	v4[14] = CString("Lupis")
	v4[15] = CString("Mallik")
	v4[16] = CString("Thrall")
	v4[17] = CString("Norwood")
	v4[18] = CString("Nulik")
	v4[19] = CString("Orin")
	v4[20] = CString("Olaf")
	v4[21] = CString("Orguk")
	v4[22] = CString("Pervis")
	v4[23] = CString("Paavik")
	v4[24] = CString("Qix")
	v4[25] = CString("Xevin")
	v4[26] = CString("Xurcon")
	v4[27] = CString("Markoan")
	v4[28] = CString("Yuric")
	v4[29] = CString("Yoovis")
	v4[30] = CString("Yalek")
	v4[31] = CString("Zug")
	v4[32] = CString("Zivik")
	v4[33] = nil
	if *memmap.PtrUint32(6112660, 814516) == 0 {
		v1 = &v4[0]
		for {
			v2 = *(**byte)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof((*byte)(nil))*1))
			v1 = (**byte)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof((*byte)(nil))*1))
			v0++
			if v2 == nil {
				break
			}
		}
		*memmap.PtrUint32(6112660, 814516) = uint32(v0)
	}
	return v4[randomIntMinMax(0, v0-1)]
}
