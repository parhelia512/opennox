package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"math"
	"unsafe"
)

func nox_thing_undead_killer_draw(a1 *int32, dr *nox_drawable) int32 {
	var (
		result int32
		v3     int16
		v4     int32
		v5     int32
		v6     int32
		v7     int32
		v8     int32
		v9     int32
		v10    int32
		v11    int32
		v12    int32
		v13    int8
		v14    [4]int16
		v15    int2
		a2     int32 = int32(uintptr(unsafe.Pointer(dr)))
	)
	if *memmap.PtrUint32(6112660, 1313728) == 0 {
		*memmap.PtrUint32(6112660, 1313736) = uint32(nox_xxx_getTTByNameSpriteMB_44CFC0("WhiteOrb"))
		*memmap.PtrUint32(6112660, 1313732) = nox_color_rgb_4344A0(100, 100, math.MaxUint8)
		*memmap.PtrUint32(6112660, 1313728) = 1
	}
	if (nox_frame_xxx_2598000 - *(*uint32)(unsafe.Pointer(uintptr(a2 + 316)))) <= 70 {
		if randomIntMinMax(0, 100) > 85 {
			v3 = int16(*(*uint16)(unsafe.Pointer(uintptr(a2 + 328))))
			v14[0] = int16(*(*uint16)(unsafe.Pointer(uintptr(a2 + 324))))
			v14[1] = v3
			v14[2] = int16(int32(*(*uint16)(unsafe.Pointer(uintptr(a2 + 12)))) + randomIntMinMax(-5, 5))
			v14[3] = int16(int32(*(*uint16)(unsafe.Pointer(uintptr(a2 + 16)))) + randomIntMinMax(-5, 5))
			v13 = int8(randomIntMinMax(6, 10))
			v4 = randomIntMinMax(-5, 5)
			sub_499490(*memmap.PtrInt32(6112660, 1313736), (*uint16)(unsafe.Pointer(&v14[0])), 0, v4, v13, 0)
		}
		v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 12))))
		v6 = *(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*1)) - int32(*(*int16)(unsafe.Pointer(uintptr(a2 + 106)))) - int32(*(*int16)(unsafe.Pointer(uintptr(a2 + 104))))
		v7 = *(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*5))
		v8 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 328))))
		v15.field_0 = v5 + *a1 - *(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*4))
		v12 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 324))))
		v11 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 16))))
		v15.field_4 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 16))) + uint32(v6) - uint32(v7))
		v9 = int32(sub_48C690(v5, v11, v12, v8))
		v10 = 8 - v9/40
		if v10 < 0 {
			v10 = 1
		}
		sub_4B6720(&v15, *memmap.PtrInt32(6112660, 1313732), v10, 12)
		result = 1
	} else {
		nox_xxx_spriteDeleteStatic_45A4E0_drawable(dr)
		result = 0
	}
	return result
}
