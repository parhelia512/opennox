package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"math"
	"unsafe"
)

func nox_xxx_updDrawMagic_4CDD80(a1 int32, a2 *uint32) {
	var (
		v2  *uint32
		v3  int32
		v4  int32
		v5  *uint32
		v6  int32
		v7  int32
		v8  int32
		v9  int32
		v10 int32
		v11 int32
		v13 int32
		v14 int32
		v15 int32
	)
	v2 = a2
	v3 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*3)) - *(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*108)))
	if uint32(v3*v3)+(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*4))-*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*109)))*(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*4))-*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*109))) > 200 {
		v4 = int32(*memmap.PtrUint32(6112660, 1523000))
		if *memmap.PtrUint32(6112660, 1523000) == 0 {
			v4 = nox_xxx_getTTByNameSpriteMB_44CFC0("MagicTailLink")
			*memmap.PtrUint32(6112660, 1523000) = uint32(v4)
		}
		v5 = &nox_xxx_spriteLoadAdd_45A360_drawable(v4, int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*108))), int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*109)))).field_0
		*(*uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(uint32(0))*108)) = *(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*3))
		*(*uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(uint32(0))*109)) = *(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*4))
		nox_xxx_sprite_45A110_drawable((*nox_drawable)(unsafe.Pointer(v5)))
		*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*108)) = *(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*3))
		*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*109)) = *(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*4))
		nox_xxx_spriteTransparentDecay_49B950(v5, int32(nox_gameFPS))
	}
	v13 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*3)) - *(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*8)))
	v14 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*4)) - *(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*9)))
	if *memmap.PtrUint32(6112660, 1523004) == 0 {
		*memmap.PtrUint32(6112660, 1523004) = uint32(nox_xxx_getTTByNameSpriteMB_44CFC0("BlueSpark"))
	}
	v6 = 0
	v7 = 0
	v15 = 4
	for {
		v8 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*8)) + uint32(v7/4) + uint32(randomIntMinMax(-8, 8)))
		v9 = randomIntMinMax(-8, 8)
		v10 = int32(uintptr(unsafe.Pointer(nox_xxx_spriteLoadAdd_45A360_drawable(*memmap.PtrInt32(6112660, 1523004), v8, int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*9))+uint32(v6/4)+uint32(v9))))))
		v11 = v10
		if v10 != 0 {
			*(*uint32)(unsafe.Pointer(uintptr(v10 + 300))) = uint32(cgoFuncAddr(nox_thing_magic_sparkle_draw))
			nox_xxx_spriteChangeLightColor_484BE0((*uint32)(unsafe.Pointer(uintptr(v10+136))), 128, 128, math.MaxUint8)
			*(*uint32)(unsafe.Pointer(uintptr(v11 + 432))) = *(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*3)) << 12
			*(*uint32)(unsafe.Pointer(uintptr(v11 + 436))) = *(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*4)) << 12
			*(*uint8)(unsafe.Pointer(uintptr(v11 + 299))) = 0
			*(*uint32)(unsafe.Pointer(uintptr(v11 + 440))) = 0
			*(*uint32)(unsafe.Pointer(uintptr(v11 + 448))) = nox_frame_xxx_2598000 + uint32(randomIntMinMax(10, 20))
			*(*uint32)(unsafe.Pointer(uintptr(v11 + 444))) = nox_frame_xxx_2598000
			*(*uint16)(unsafe.Pointer(uintptr(v11 + 104))) = *((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v2))), unsafe.Sizeof(uint16(0))*52)))
			*(*uint16)(unsafe.Pointer(uintptr(v11 + 106))) = *((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v2))), unsafe.Sizeof(uint16(0))*53)))
			*(*uint8)(unsafe.Pointer(uintptr(v11 + 296))) = 0
			nox_xxx_sprite_45A110_drawable((*nox_drawable)(unsafe.Pointer(uintptr(v11))))
		}
		v7 += v13
		v6 += v14
		v15--
		if v15 == 0 {
			break
		}
	}
}
