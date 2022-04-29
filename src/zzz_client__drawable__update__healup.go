package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"unsafe"
)

func sub_4CD450(a1 *uint32, a2 int32) int32 {
	var (
		v2  int32
		v3  int32
		v4  int32
		v5  int32
		v6  *uint32
		v7  int32
		v8  *uint32
		v9  int32
		v10 *uint32
		v11 int32
		v12 *uint32
		v13 *uint32
		v14 int32
		v15 int32
		v16 int32
		v17 int32
		v18 int32
		v19 int32
		v20 int8
		v22 int32
		v23 int16
		v24 int32
	)
	if *memmap.PtrUint32(6112660, 1522972) == 0 {
		*memmap.PtrUint32(6112660, 1522972) = uint32(nox_xxx_getTTByNameSpriteMB_44CFC0("HealOrb"))
	}
	if randomIntMinMax(0, 100) < 50 {
		if int32(*(*uint8)(unsafe.Pointer(uintptr(a2 + 432)))) != 0 {
			if a2 == -432 {
				return 1
			}
			if nox_xxx_netTestHighBit_578B70(*(*uint32)(unsafe.Pointer(uintptr(a2 + 437)))) != 0 {
				v5 = nox_xxx_netClearHighBit_578B30(int16(uint16(*(*uint32)(unsafe.Pointer(uintptr(a2 + 437))))))
				v6 = nox_xxx_netSpriteByCodeStatic_45A720(v5)
			} else {
				v7 = nox_xxx_netClearHighBit_578B30(int16(uint16(*(*uint32)(unsafe.Pointer(uintptr(a2 + 437))))))
				v6 = nox_xxx_netSpriteByCodeDynamic_45A6F0(v7)
			}
			v8 = v6
			if nox_xxx_netTestHighBit_578B70(*(*uint32)(unsafe.Pointer(uintptr(a2 + 441)))) != 0 {
				v9 = nox_xxx_netClearHighBit_578B30(int16(uint16(*(*uint32)(unsafe.Pointer(uintptr(a2 + 441))))))
				v10 = nox_xxx_netSpriteByCodeStatic_45A720(v9)
			} else {
				v11 = nox_xxx_netClearHighBit_578B30(int16(uint16(*(*uint32)(unsafe.Pointer(uintptr(a2 + 441))))))
				v10 = nox_xxx_netSpriteByCodeDynamic_45A6F0(v11)
			}
			v12 = v10
			if v8 == nil || v10 == nil {
				return 1
			}
			v2 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v10), unsafe.Sizeof(uint32(0))*3)) + uint32(randomIntMinMax(-20, 20)))
			v3 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v12), unsafe.Sizeof(uint32(0))*4)) + uint32(randomIntMinMax(-20, 20)))
			v4 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v8), unsafe.Sizeof(uint32(0))*4)))
			v22 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v8), unsafe.Sizeof(uint32(0))*3)))
		} else {
			if a2 == -432 {
				return 1
			}
			v2 = int32(*(*uint16)(unsafe.Pointer(uintptr(a2 + 441)))) + randomIntMinMax(-20, 20)
			v3 = int32(*(*uint16)(unsafe.Pointer(uintptr(a2 + 443)))) + randomIntMinMax(-20, 20)
			*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v22))), unsafe.Sizeof(uint16(0))*0)) = *(*uint16)(unsafe.Pointer(uintptr(a2 + 437)))
			*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v4))), unsafe.Sizeof(uint16(0))*0)) = *(*uint16)(unsafe.Pointer(uintptr(a2 + 439)))
		}
		v13 = a1
		v23 = int16(v4)
		v14 = int32(*a1)
		v15 = int32(*a1 - *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*4)))
		v16 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*5)))
		v24 = int32(*(*int16)(unsafe.Pointer(uintptr(a2 + 104))))
		v17 = v2 + v15
		v18 = int32(uint32(v3) + *(*uint32)(unsafe.Add(unsafe.Pointer(v13), unsafe.Sizeof(uint32(0))*1)) - uint32(v16) - uint32(v24))
		if v17 < 0 {
			v2 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v13), unsafe.Sizeof(uint32(0))*4)) + uint32(v14) + 1)
		}
		if v18 < 0 {
			v3 = int32(uint32(v16-v24) + *(*uint32)(unsafe.Add(unsafe.Pointer(v13), unsafe.Sizeof(uint32(0))*1)) + 1)
		}
		if uint32(v17) >= *(*uint32)(unsafe.Add(unsafe.Pointer(v13), unsafe.Sizeof(uint32(0))*8)) {
			v2 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v13), unsafe.Sizeof(uint32(0))*2)) + *(*uint32)(unsafe.Add(unsafe.Pointer(v13), unsafe.Sizeof(uint32(0))*4)) - 1)
		}
		if uint32(v18) >= *(*uint32)(unsafe.Add(unsafe.Pointer(v13), unsafe.Sizeof(uint32(0))*9)) {
			v3 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v13), unsafe.Sizeof(uint32(0))*3)) - uint32(v24) + uint32(v16) - 1)
		}
		v19 = int32(uintptr(unsafe.Pointer(nox_xxx_spriteLoadAdd_45A360_drawable(*memmap.PtrInt32(6112660, 1522972), v2, v3))))
		if v19 != 0 {
			v20 = int8(randomIntMinMax(6, 12))
			*(*uint16)(unsafe.Pointer(uintptr(v19 + 432))) = uint16(int16(v22))
			*(*uint16)(unsafe.Pointer(uintptr(v19 + 434))) = uint16(v23)
			*(*uint8)(unsafe.Pointer(uintptr(v19 + 443))) = uint8(v20)
			*(*uint8)(unsafe.Pointer(uintptr(v19 + 444))) = uint8(int8(randomIntMinMax(3, 10)))
		}
	}
	return 1
}
