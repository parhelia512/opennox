package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"math"
	"unsafe"
)

func nox_xxx_updDrawAddRndSpark_4CDFA0(a1 int32, a2 *uint32) *uint32 {
	var (
		result *uint32
		v3     *uint32
		v4     *uint32
	)
	if *memmap.PtrUint32(6112660, 1523008) == 0 {
		*memmap.PtrUint32(6112660, 1523008) = uint32(nox_xxx_getTTByNameSpriteMB_44CFC0("DeathBallSpark"))
	}
	result = a2
	if int32(uintptr(unsafe.Pointer(a2))) > 0 {
		v3 = a2
		for {
			result = &nox_xxx_spriteLoadAdd_45A360_drawable(*memmap.PtrInt32(6112660, 1523008), int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 12)))), int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 16))))).field_0
			v4 = result
			if result != nil {
				*(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*108)) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 12))) << 12
				*(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*109)) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 16))) << 12
				*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(result))), 299))) = uint8(int8(randomIntMinMax(0, math.MaxUint8)))
				*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*110)) = uint32(randomIntMinMax(1000, 3000))
				*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*112)) = nox_frame_xxx_2598000 + uint32(randomIntMinMax(10, 40))
				*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*111)) = nox_frame_xxx_2598000
				*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v4))), unsafe.Sizeof(uint16(0))*52))) = 22
				*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(v4))), 296))) = uint8(int8(randomIntMinMax(0, 4)))
				result = nox_xxx_sprite_45A110_drawable((*nox_drawable)(unsafe.Pointer(v4)))
			}
			v3 = (*uint32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(v3))), -1))))
			if v3 == nil {
				break
			}
		}
	}
	return result
}
func nox_xxx_updDrawDBallCharge_4CE0C0(a1 int32, a2 int32) int32 {
	var (
		v2  int32
		v3  int16
		v4  int32
		v5  int16
		v6  int16
		v7  int32
		v9  int32
		v10 int8
		v11 [4]int16
	)
	if *memmap.PtrUint32(6112660, 1523012) == 0 {
		*memmap.PtrUint32(6112660, 1523012) = uint32(nox_xxx_getTTByNameSpriteMB_44CFC0("CharmOrb"))
	}
	v2 = 10
	v3 = int16(*(*uint16)(unsafe.Pointer(uintptr(a2 + 16))))
	v11[0] = int16(*(*uint16)(unsafe.Pointer(uintptr(a2 + 12))))
	v11[1] = v3
	for {
		v4 = randomIntMinMax(0, math.MaxUint8)
		v5 = int16(randomIntMinMax(2, 8))
		v6 = int16(int32(*(*uint16)(unsafe.Pointer(uintptr(a2 + 16)))) + int32(v5)*int32(*memmap.PtrInt16(0x587000, uint32(v4*8)+0x2EE5C)))
		v11[2] = int16(int32(*(*uint16)(unsafe.Pointer(uintptr(a2 + 12)))) + int32(v5)*int32(*memmap.PtrInt16(0x587000, uint32(v4*8)+0x2EE58)))
		v11[3] = v6
		if randomIntMinMax(0, 100) < 50 {
			v10 = int8(randomIntMinMax(6, 10))
			v9 = randomIntMinMax(-20, 20)
			v7 = randomIntMinMax(-20, 20)
			sub_499490(*memmap.PtrInt32(6112660, 1523012), (*uint16)(unsafe.Pointer(&v11[0])), v7, v9, v10, 0)
		}
		v2--
		if v2 == 0 {
			break
		}
	}
	return 1
}
