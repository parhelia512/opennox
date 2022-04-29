package opennox

import (
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"unsafe"
)

func nox_thing_vortex_draw(a1 *int32, dr *nox_drawable) int32 {
	var (
		v2     int32
		v3     int32
		v4     int32
		v5     bool
		v6     int8
		v7     int32
		v8     int32
		v9     int32
		result int32
		v11    float32
		xLeft  int2
		a2a    int2
		a3     int2
		a2     int32 = int32(uintptr(unsafe.Pointer(dr)))
	)
	if *memmap.PtrUint32(6112660, 1313820) == 0 {
		dword_5d4594_1313816 = nox_color_rgb_4344A0(170, 170, 170)
		*memmap.PtrUint32(6112660, 1313820) = 1
	}
	v2 = int32(*(*uint8)(unsafe.Pointer(uintptr(a2 + 448)))) * 8
	v3 = int32(*(*uint8)(unsafe.Pointer(uintptr(a2 + 450))))
	a2a.field_0 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 440))) + uint32(v3**memmap.PtrInt32(0x587000, uint32(v2)+0x2EE58)/16))
	a2a.field_4 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 444))) + uint32(v3**memmap.PtrInt32(0x587000, uint32(v2)+0x2EE5C)/16))
	sub_4739E0((*uint32)(unsafe.Pointer(a1)), &a2a, &xLeft)
	v4 = xLeft.field_4 - int32(*(*int16)(unsafe.Pointer(uintptr(a2 + 104))))
	v5 = xLeft.field_0 <= *a1
	xLeft.field_4 -= int32(*(*int16)(unsafe.Pointer(uintptr(a2 + 104))))
	if v5 || xLeft.field_0 >= *(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*2)) || v4 <= *(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*1)) || v4 >= *(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*3)) {
		goto LABEL_22
	}
	if a2a.field_4 >= *(*int32)(unsafe.Pointer(uintptr(a2 + 444))) {
		sub_4B6720(&xLeft, int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 432)))), 3, 5)
		nox_client_drawSetColor_434460(int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 436)))))
		nox_xxx_drawPointMB_499B70(xLeft.field_0, xLeft.field_4, 3)
	} else {
		sub_4B6720(&xLeft, *(*int32)(unsafe.Pointer(&dword_5d4594_1313816)), 2, 4)
		nox_client_drawSetColor_434460(*(*int32)(unsafe.Pointer(&dword_5d4594_1313816)))
		nox_xxx_drawPointMB_499B70(xLeft.field_0, xLeft.field_4, 2)
	}
	nox_client_drawAddPoint_49F500(xLeft.field_0, xLeft.field_4)
	v6 = int8(*(*uint8)(unsafe.Pointer(uintptr(a2 + 449))))
	v7 = int32(*(*uint8)(unsafe.Pointer(uintptr(a2 + 448)))) - int32(v6)*2
	if int32(v6) > 0 {
		if v7 < 0 {
			v7 += 256
		} else {
			v7 += 0
		}
	} else if v7 >= 256 {
		v7 -= 256
	}
	v8 = int32(*(*uint8)(unsafe.Pointer(uintptr(a2 + 450))))
	a2a.field_0 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 440))) + uint32(v8)**memmap.PtrUint32(0x587000, uint32(v7*8)+0x2EE58)/16)
	a2a.field_4 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 444))) + uint32(v8)**memmap.PtrUint32(0x587000, uint32(v7*8)+0x2EE5C)/16)
	sub_4739E0((*uint32)(unsafe.Pointer(a1)), &a2a, &xLeft)
	xLeft.field_4 -= int32(*(*int16)(unsafe.Pointer(uintptr(a2 + 104))))
	nox_client_drawAddPoint_49F500(xLeft.field_0, xLeft.field_4)
	if a2a.field_4 >= *(*int32)(unsafe.Pointer(uintptr(a2 + 444))) {
		nox_client_drawSetColor_434460(int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 436)))))
	} else {
		nox_client_drawSetColor_434460(*(*int32)(unsafe.Pointer(&dword_5d4594_1313816)))
	}
	nox_client_drawLineFromPoints_49E4B0()
	*(*uint8)(unsafe.Pointer(uintptr(a2 + 448))) += *(*uint8)(unsafe.Pointer(uintptr(a2 + 449)))
	*(*uint16)(unsafe.Pointer(uintptr(a2 + 104))) += uint16(*(*uint8)(unsafe.Pointer(uintptr(a2 + 451))))
	sub_4739E0((*uint32)(unsafe.Pointer(a1)), (*int2)(unsafe.Pointer(uintptr(a2+12))), &a3)
	v11 = float32(float64(*(*int16)(unsafe.Pointer(uintptr(a2 + 104)))) * 0.0024999999 * 50.0)
	v9 = int(v11)
	if 50-v9 <= 0 {
	LABEL_22:
		nox_xxx_spriteDeleteStatic_45A4E0_drawable(dr)
		result = 0
	} else {
		*(*uint8)(unsafe.Pointer(uintptr(a2 + 450))) = uint8(int8(50 - v9))
		result = 1
	}
	return result
}
