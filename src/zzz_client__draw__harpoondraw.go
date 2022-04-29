package opennox

import (
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"unsafe"
)

func nox_thing_harpoon_draw(a1 *int32, dr *nox_drawable) int32 {
	return nox_thing_slave_draw(a1, dr)
}
func nox_thing_harpoon_rope_draw(a1 *int32, dr *nox_drawable) int32 {
	var (
		v2  int32
		v3  int32
		v4  uint16
		v5  int32
		v6  int32
		v7  int32
		v8  int32
		v9  *uint32
		v10 *uint32
		v11 *uint32
		v12 int32
		v13 int32
		v14 int32
		a1a int2
		a2a int2
		v18 int32
		a2  int32 = int32(uintptr(unsafe.Pointer(dr)))
	)
	if int32(*(*uint8)(unsafe.Pointer(uintptr(a2 + 432)))) == 0 {
		v2 = *(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*4))
		v3 = *a1
		v4 = *(*uint16)(unsafe.Pointer(uintptr(a2 + 439)))
		a1a.field_0 = *a1 + int32(*(*uint16)(unsafe.Pointer(uintptr(a2 + 437)))) - v2
		v5 = *(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*5))
		v6 = *(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*1))
		a1a.field_4 = int32(v4) - v5 + v6 - 20
		v7 = int32(*(*uint16)(unsafe.Pointer(uintptr(a2 + 441)))) - v2
		v8 = int32(*(*uint16)(unsafe.Pointer(uintptr(a2 + 443)))) - v5
		a2a.field_0 = v3 + v7
		a2a.field_4 = v8 + v6 - 20
	LABEL_12:
		*memmap.PtrUint32(6112660, 1312492) = nox_color_rgb_4344A0(144, 104, 64)
		*memmap.PtrUint32(6112660, 1312496) = nox_color_rgb_4344A0(24, 16, 0)
		sub_4B63B0(&a1a, &a2a)
		return 1
	}
	if nox_xxx_netTestHighBit_578B70(*(*uint32)(unsafe.Pointer(uintptr(a2 + 437)))) != 0 {
		v9 = nox_xxx_netSpriteByCodeStatic_45A720(int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 437)))))
	} else {
		v9 = nox_xxx_netSpriteByCodeDynamic_45A6F0(int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 437)))))
	}
	v10 = v9
	if nox_xxx_netTestHighBit_578B70(*(*uint32)(unsafe.Pointer(uintptr(a2 + 441)))) != 0 {
		v11 = nox_xxx_netSpriteByCodeStatic_45A720(int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 441)))))
	} else {
		v11 = nox_xxx_netSpriteByCodeDynamic_45A6F0(int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 441)))))
	}
	if v10 != nil && v11 != nil {
		v12 = *(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*4))
		v13 = *a1
		v14 = *(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*1))
		v18 = *(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*5))
		a1a.field_0 = int32(uint32(v13) + *(*uint32)(unsafe.Add(unsafe.Pointer(v10), unsafe.Sizeof(uint32(0))*3)) - uint32(v12))
		a1a.field_4 = int32(uint32(v14) + *(*uint32)(unsafe.Add(unsafe.Pointer(v10), unsafe.Sizeof(uint32(0))*4)) - uint32(v18))
		a2a.field_0 = int32(uint32(v13) + *(*uint32)(unsafe.Add(unsafe.Pointer(v11), unsafe.Sizeof(uint32(0))*3)) - uint32(v12))
		a2a.field_4 = int32(uint32(v14) + *(*uint32)(unsafe.Add(unsafe.Pointer(v11), unsafe.Sizeof(uint32(0))*4)) - uint32(*((*int16)(unsafe.Add(unsafe.Pointer((*int16)(unsafe.Pointer(v11))), unsafe.Sizeof(int16(0))*53)))) - uint32(*((*int16)(unsafe.Add(unsafe.Pointer((*int16)(unsafe.Pointer(v11))), unsafe.Sizeof(int16(0))*52)))) - uint32(v18))
		a1a.field_0 += int32(*memmap.PtrUint32(0x587000, uint32(int32(*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(v10))), 297))))*8)+0x2AEF8))
		a1a.field_4 += int32(*memmap.PtrUint32(0x587000, uint32(int32(*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(v10))), 297))))*8)+0x2AEFC))
		a2a.field_4 -= 8
		goto LABEL_12
	}
	return 1
}
