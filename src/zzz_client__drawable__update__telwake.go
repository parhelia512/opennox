package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"math"
	"unsafe"
)

func nox_xxx_updDrawTeleportWake_4CD8D0(a1 int32, a2 int32) int32 {
	var (
		v2 int32
		v3 int32
		v4 int32
		v5 int32
	)
	if *memmap.PtrUint32(6112660, 1522980) == 0 {
		*memmap.PtrUint32(6112660, 1522980) = uint32(nox_xxx_getTTByNameSpriteMB_44CFC0("BlueSpark"))
	}
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 12))) + uint32(randomIntMinMax(-5, 5)))
	v3 = randomIntMinMax(-5, 5)
	v4 = int32(uintptr(unsafe.Pointer(nox_xxx_spriteLoadAdd_45A360_drawable(*memmap.PtrInt32(6112660, 1522980), v2, int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 16)))+uint32(v3))))))
	v5 = v4
	if v4 != 0 {
		*(*uint32)(unsafe.Pointer(uintptr(v4 + 432))) = *(*uint32)(unsafe.Pointer(uintptr(v4 + 12))) << 12
		*(*uint32)(unsafe.Pointer(uintptr(v4 + 436))) = *(*uint32)(unsafe.Pointer(uintptr(v4 + 16))) << 12
		*(*uint8)(unsafe.Pointer(uintptr(v4 + 299))) = uint8(int8(randomIntMinMax(0, math.MaxUint8)))
		*(*uint32)(unsafe.Pointer(uintptr(v5 + 440))) = uint32(randomIntMinMax(1, 100))
		*(*uint32)(unsafe.Pointer(uintptr(v5 + 448))) = nox_frame_xxx_2598000 + uint32(randomIntMinMax(10, 32))
		*(*uint32)(unsafe.Pointer(uintptr(v5 + 444))) = nox_frame_xxx_2598000
		*(*uint16)(unsafe.Pointer(uintptr(v5 + 104))) = 0
		*(*uint8)(unsafe.Pointer(uintptr(v5 + 296))) = uint8(int8(randomIntMinMax(3, 8)))
		nox_xxx_sprite_45A110_drawable((*nox_drawable)(unsafe.Pointer(uintptr(v5))))
	}
	return 1
}
