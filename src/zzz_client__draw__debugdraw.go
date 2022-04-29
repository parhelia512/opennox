package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"image"
	"unsafe"
)

func nox_thing_debug_draw(vp *nox_draw_viewport_t, dr *nox_drawable) int32 {
	var (
		a1  *uint32 = (*uint32)(unsafe.Pointer(vp))
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
		v16 float32
		a2a int2
		v18 int32
		v19 int32
	)
	_ = v19
	var v20 int32
	v2 = int32(*memmap.PtrUint32(0x8531A0, 2572))
	if dr.field_72 >= nox_frame_xxx_2598000 {
		v2 = int32(nox_color_yellow_2589772)
	}
	nox_client_drawSetColor_434460(v2)
	nox_xxx_drawSetTextColor_434390(int32(nox_color_white_2523948))
	v3 = dr.pos.y
	v4 = int32(*a1 + uint32(dr.pos.x) - *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*4)))
	v5 = int32(dr.flags28)
	v6 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*1)) - *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*5)))
	a2a.field_0 = int32(*a1 + uint32(dr.pos.x) - *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*4)))
	a2a.field_4 = v3 + v6
	if (v5 & 128) == 0 {
		if v5&2 != 0 {
			sub_4BD010(dr, &a2a, v2)
			nox_client_drawAddPoint_49F500(a2a.field_0, a2a.field_4)
			nox_xxx_rasterPointRel_49F570(int32(*memmap.PtrUint32(0x587000, uint32(int32(dr.field_74_2)*8)+179880)), int32(*memmap.PtrUint32(0x587000, uint32(int32(dr.field_74_2)*8)+0x2BEAC)))
			nox_client_drawLineFromPoints_49E4B0()
			nox_swprintf((*libc.WChar)(memmap.PtrOff(6112660, 1316540)), libc.CWString("%d"), dr.field_32)
			noxrend.DrawString(nil, (*libc.WChar)(unsafe.Pointer(memmap.PtrInt16(6112660, 1316540))), image.Pt(a2a.field_0, a2a.field_4-10))
			nox_swprintf((*libc.WChar)(memmap.PtrOff(6112660, 1316540)), libc.CWString("%S"), nox_get_thing_name(int32(dr.field_27)))
			noxrend.DrawString(nil, (*libc.WChar)(unsafe.Pointer(memmap.PtrInt16(6112660, 1316540))), image.Pt(a2a.field_0, a2a.field_4))
			nox_swprintf((*libc.WChar)(memmap.PtrOff(6112660, 1316540)), libc.CWString("%S"), *memmap.PtrUint32(0x587000, dr.field_69*4+178920))
			noxrend.DrawString(nil, (*libc.WChar)(unsafe.Pointer(memmap.PtrInt16(6112660, 1316540))), image.Pt(a2a.field_0, a2a.field_4+10))
		} else if v5&4 != 0 {
			sub_4BD010(dr, &a2a, v2)
			nox_client_drawAddPoint_49F500(a2a.field_0, a2a.field_4)
			nox_xxx_rasterPointRel_49F570(int32(*memmap.PtrUint32(0x587000, uint32(int32(dr.field_74_2)*8)+179880)), int32(*memmap.PtrUint32(0x587000, uint32(int32(dr.field_74_2)*8)+0x2BEAC)))
			nox_client_drawLineFromPoints_49E4B0()
			nox_swprintf((*libc.WChar)(memmap.PtrOff(6112660, 1316540)), libc.CWString("%d"), dr.field_32)
			noxrend.DrawString(nil, (*libc.WChar)(unsafe.Pointer(memmap.PtrInt16(6112660, 1316540))), image.Pt(a2a.field_0, a2a.field_4-10))
			nox_swprintf((*libc.WChar)(memmap.PtrOff(6112660, 1316540)), libc.CWString("%S"), nox_get_thing_name(int32(dr.field_27)))
			noxrend.DrawString(nil, (*libc.WChar)(unsafe.Pointer(memmap.PtrInt16(6112660, 1316540))), image.Pt(a2a.field_0, a2a.field_4))
			nox_swprintf((*libc.WChar)(memmap.PtrOff(6112660, 1316540)), libc.CWString("%S"), *memmap.PtrUint32(0x587000, dr.field_69*4+0x2BA08))
			noxrend.DrawString(nil, (*libc.WChar)(unsafe.Pointer(memmap.PtrInt16(6112660, 1316540))), image.Pt(a2a.field_0, a2a.field_4+10))
		} else {
			sub_4BD010(dr, &a2a, v2)
		}
	} else {
		v20 = v4
		v7 = int(dr.field_25)
		v16 = dr.field_24
		v8 = a2a.field_4 - v7
		v19 = a2a.field_0
		v9 = a2a.field_4 - int(v16)
		v10 = int32(dr.field_74_4) * 8
		v11 = *memmap.PtrInt32(0x587000, uint32(v10)+0x2FE58)
		v18 = *memmap.PtrInt32(0x587000, uint32(v10)+0x2FE5C)
		nox_client_drawAddPoint_49F500(a2a.field_0, v9)
		nox_xxx_rasterPointRel_49F570(v11, v18)
		nox_client_drawLineFromPoints_49E4B0()
		nox_client_drawAddPoint_49F500(v20, v8)
		nox_xxx_rasterPointRel_49F570(v11, v18)
		nox_client_drawLineFromPoints_49E4B0()
		nox_client_drawAddPoint_49F500(a2a.field_0, v9)
		nox_client_drawAddPoint_49F500(v20, v8)
		nox_client_drawLineFromPoints_49E4B0()
		nox_client_drawAddPoint_49F500(a2a.field_0+v11, v9+v18)
		nox_client_drawAddPoint_49F500(v20+v11, v8+v18)
		nox_client_drawLineFromPoints_49E4B0()
	}
	nox_swprintf((*libc.WChar)(memmap.PtrOff(6112660, 1316540)), libc.CWString("%d"), dr.field_32)
	noxrend.DrawString(nil, (*libc.WChar)(unsafe.Pointer(memmap.PtrInt16(6112660, 1316540))), image.Pt(a2a.field_0, a2a.field_4-10))
	nox_swprintf((*libc.WChar)(memmap.PtrOff(6112660, 1316540)), libc.CWString("%S"), nox_get_thing_name(int32(dr.field_27)))
	noxrend.DrawString(nil, (*libc.WChar)(unsafe.Pointer(memmap.PtrInt16(6112660, 1316540))), image.Pt(a2a.field_0, a2a.field_4))
	return 1
}
