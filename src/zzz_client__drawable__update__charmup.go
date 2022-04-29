package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	"unsafe"
)

func sub_4CD150(a1 int32, a2 *uint32, a3 int32, a4 int32) {
	var (
		v4  int32
		v5  int32
		v6  int32
		v7  int32
		v8  int32
		v9  int32
		v10 int32
		v11 int32
		v12 *uint32
		v13 int32
		v14 *uint32
		v15 int32
		v16 int32
		v17 int32
		v18 int32
		v19 int32
		v20 int32
		v21 int32
		v22 int32
		v23 int8
		v25 int32
		v26 int32
		v27 int32
	)
	v4 = randomIntMinMax(0, 100)
	if v4 >= 50 {
		return
	}
	*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v4))), 0)) = *(*uint8)(unsafe.Pointer(uintptr(a3 + 432)))
	if int32(uint8(int8(v4))) == 0 {
		if a3 == -432 {
			return
		}
		if a4 != 0 {
			v5 = int32(*(*uint16)(unsafe.Pointer(uintptr(a3 + 441)))) + randomIntMinMax(-20, 20)
			v6 = randomIntMinMax(-20, 20)
			*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v7))), unsafe.Sizeof(uint16(0))*0)) = *(*uint16)(unsafe.Pointer(uintptr(a3 + 437)))
			v8 = int32(*(*uint16)(unsafe.Pointer(uintptr(a3 + 443)))) + v6
			*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v9))), unsafe.Sizeof(uint16(0))*0)) = *(*uint16)(unsafe.Pointer(uintptr(a3 + 439)))
		} else {
			v5 = int32(*(*uint16)(unsafe.Pointer(uintptr(a3 + 437)))) + randomIntMinMax(-20, 20)
			v10 = randomIntMinMax(-20, 20)
			*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v7))), unsafe.Sizeof(uint16(0))*0)) = *(*uint16)(unsafe.Pointer(uintptr(a3 + 441)))
			v8 = int32(*(*uint16)(unsafe.Pointer(uintptr(a3 + 439)))) + v10
			*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v9))), unsafe.Sizeof(uint16(0))*0)) = *(*uint16)(unsafe.Pointer(uintptr(a3 + 443)))
		}
		goto LABEL_19
	}
	if a3 == -432 {
		return
	}
	if nox_xxx_netTestHighBit_578B70(*(*uint32)(unsafe.Pointer(uintptr(a3 + 437)))) != 0 {
		v11 = nox_xxx_netClearHighBit_578B30(int16(uint16(*(*uint32)(unsafe.Pointer(uintptr(a3 + 437))))))
		v12 = nox_xxx_netSpriteByCodeStatic_45A720(v11)
	} else {
		v13 = nox_xxx_netClearHighBit_578B30(int16(uint16(*(*uint32)(unsafe.Pointer(uintptr(a3 + 437))))))
		v12 = nox_xxx_netSpriteByCodeDynamic_45A6F0(v13)
	}
	v14 = v12
	if nox_xxx_netTestHighBit_578B70(*(*uint32)(unsafe.Pointer(uintptr(a3 + 441)))) != 0 {
		v15 = nox_xxx_netClearHighBit_578B30(int16(uint16(*(*uint32)(unsafe.Pointer(uintptr(a3 + 441))))))
		v4 = int32(uintptr(unsafe.Pointer(nox_xxx_netSpriteByCodeStatic_45A720(v15))))
	} else {
		v16 = nox_xxx_netClearHighBit_578B30(int16(uint16(*(*uint32)(unsafe.Pointer(uintptr(a3 + 441))))))
		v4 = int32(uintptr(unsafe.Pointer(nox_xxx_netSpriteByCodeDynamic_45A6F0(v16))))
	}
	v17 = v4
	if v14 == nil || v4 == 0 {
		return
	}
	if a4 == 0 {
		v5 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v14), unsafe.Sizeof(uint32(0))*3)) + uint32(randomIntMinMax(-20, 20)))
		v19 = randomIntMinMax(-20, 20)
		v7 = int32(*(*uint32)(unsafe.Pointer(uintptr(v17 + 12))))
		v8 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v14), unsafe.Sizeof(uint32(0))*4)) + uint32(v19))
		v9 = int32(*(*uint32)(unsafe.Pointer(uintptr(v17 + 16))))
	LABEL_19:
		*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v25))), unsafe.Sizeof(uint16(0))*0)) = uint16(int16(v9))
		goto LABEL_20
	}
	v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(v4 + 12))) + uint32(randomIntMinMax(-20, 20)))
	v18 = randomIntMinMax(-20, 20)
	v7 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v14), unsafe.Sizeof(uint32(0))*3)))
	v8 = int32(*(*uint32)(unsafe.Pointer(uintptr(v17 + 16))) + uint32(v18))
	v25 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v14), unsafe.Sizeof(uint32(0))*4)))
LABEL_20:
	v20 = int32(uint32(v5) + *a2 - *(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*4)))
	v27 = int32(*(*int16)(unsafe.Pointer(uintptr(a3 + 104))))
	v26 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*5)))
	v21 = int32(uint32(v8) + *(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*1)) - uint32(v26) - uint32(v27))
	if v20 < 0 {
		v5 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*4)) + *a2 + 1)
	}
	if v21 < 0 {
		v8 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*1)) + uint32(v26) - uint32(v27) + 1)
	}
	if uint32(v20) >= *(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*8)) {
		v5 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*2)) + *(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*4)) - 1)
	}
	if uint32(v21) >= *(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*9)) {
		v8 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*3)) - uint32(v27) + uint32(v26) - 1)
	}
	v4 = int32(uintptr(unsafe.Pointer(nox_xxx_spriteLoadAdd_45A360_drawable(a1, v5, v8))))
	v22 = v4
	if v4 != 0 {
		v23 = int8(randomIntMinMax(6, 12))
		*(*uint16)(unsafe.Pointer(uintptr(v22 + 432))) = uint16(int16(v7))
		*(*uint16)(unsafe.Pointer(uintptr(v22 + 434))) = uint16(int16(v25))
		*(*uint8)(unsafe.Pointer(uintptr(v22 + 443))) = uint8(v23)
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v4))), 0)) = uint8(int8(randomIntMinMax(3, 10)))
		*(*uint8)(unsafe.Pointer(uintptr(v22 + 444))) = uint8(int8(v4))
	}
	return
}
