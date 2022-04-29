package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"math"
	"unsafe"
)

func nox_xxx_updDrawVortexSource_4CC950(a1 int32, a2 int32) int32 {
	var (
		v2 int32
		v3 int32
		v4 int32
		v5 int8
		v6 int32
		v7 int32
		v9 int32
	)
	if *memmap.PtrUint32(6112660, 1522952) == 0 {
		*memmap.PtrUint32(6112660, 1522952) = uint32(nox_xxx_getTTByNameSpriteMB_44CFC0("WhiteVortexOrb"))
		*memmap.PtrUint32(6112660, 1522944) = nox_color_rgb_4344A0(200, 200, 200)
		*memmap.PtrUint32(6112660, 1522948) = nox_color_rgb_4344A0(math.MaxUint8, math.MaxUint8, math.MaxUint8)
	}
	v2 = int32(*memmap.PtrUint32(6112660, 1522944))
	v3 = int32(*memmap.PtrUint32(6112660, 1522952))
	v9 = int32(*memmap.PtrUint32(6112660, 1522948))
	v4 = randomIntMinMax(0, math.MaxUint8)
	v5 = int8(v4)
	v6 = int32(uintptr(unsafe.Pointer(nox_xxx_spriteLoadAdd_45A360_drawable(v3, int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 12)))+uint32(*memmap.PtrInt32(0x587000, uint32(v4*8)+0x2EE58)*50/16)), int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 16)))+uint32(*memmap.PtrInt32(0x587000, uint32(v4*8)+0x2EE5C)*50/16))))))
	v7 = v6
	if v6 != 0 {
		*(*uint8)(unsafe.Pointer(uintptr(v6 + 448))) = uint8(v5)
		*(*uint16)(unsafe.Pointer(uintptr(a2 + 104))) = 0
		*(*uint8)(unsafe.Pointer(uintptr(v6 + 449))) = uint8(int8(randomIntMinMax(2, 3)))
		if randomIntMinMax(0, 100) > 50 {
			*(*uint8)(unsafe.Pointer(uintptr(v7 + 449))) = -*(*uint8)(unsafe.Pointer(uintptr(v7 + 449)))
		}
		*(*uint8)(unsafe.Pointer(uintptr(v7 + 451))) = 1
		*(*uint8)(unsafe.Pointer(uintptr(v7 + 450))) = 50
		*(*uint32)(unsafe.Pointer(uintptr(v7 + 440))) = *(*uint32)(unsafe.Pointer(uintptr(a2 + 12)))
		*(*uint32)(unsafe.Pointer(uintptr(v7 + 444))) = *(*uint32)(unsafe.Pointer(uintptr(a2 + 16)))
		*(*uint32)(unsafe.Pointer(uintptr(v7 + 432))) = uint32(v2)
		*(*uint32)(unsafe.Pointer(uintptr(v7 + 436))) = uint32(v9)
		nox_xxx_sprite_45A110_drawable((*nox_drawable)(unsafe.Pointer(uintptr(v7))))
		nox_xxx_spriteToSightDestroyList_49BAB0_drawable((*uint32)(unsafe.Pointer(uintptr(v7))))
	}
	return 1
}
