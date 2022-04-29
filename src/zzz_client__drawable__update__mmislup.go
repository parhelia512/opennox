package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"math"
	"unsafe"
)

func nox_xxx_updDrawMagicMissile_4CD9E0(a1 int32, a2 *uint32) int32 {
	var (
		v2  *uint32
		v3  int32
		v4  int32
		v5  int32
		v6  int32
		v8  int32
		v9  int32
		v10 int32
		v11 int32
		v12 *uint32
		v14 int32
		v15 int32
		v16 int32
	)
	v2 = a2
	if *memmap.PtrUint32(6112660, 1522984) == 0 {
		*memmap.PtrUint32(6112660, 1522984) = uint32(nox_xxx_getTTByNameSpriteMB_44CFC0("Spark"))
		*memmap.PtrUint32(6112660, 1522988) = uint32(nox_xxx_getTTByNameSpriteMB_44CFC0("Puff"))
		*memmap.PtrUint32(6112660, 1522992) = uint32(nox_xxx_getTTByNameSpriteMB_44CFC0("MagicMissileTailLink"))
	}
	v3 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*3)) - *(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*8)))
	v15 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*4)) - *(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*9)))
	v14 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*3)) - *(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*8)))
	v16 = 0
	if *memmap.PtrUint32(0x587000, 190108) > 0 {
		for {
			v4 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*8)))
			v5 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*9)))
			v6 = v3*randomIntMinMax(0, 100)/100 + v4
			v8 = (v15*randomIntMinMax(0, 100))/100 + v5
			v9 = int32(uintptr(unsafe.Pointer(nox_xxx_spriteLoadAdd_45A360_drawable(*memmap.PtrInt32(6112660, 1522984), v6, v8))))
			v10 = v9
			if v9 != 0 {
				*(*uint32)(unsafe.Pointer(uintptr(v9 + 432))) = uint32(v6 << 12)
				*(*uint32)(unsafe.Pointer(uintptr(v9 + 436))) = uint32(v8 << 12)
				*(*uint8)(unsafe.Pointer(uintptr(v9 + 299))) = uint8(int8(randomIntMinMax(0, math.MaxUint8)))
				*(*uint32)(unsafe.Pointer(uintptr(v10 + 440))) = 0
				*(*uint32)(unsafe.Pointer(uintptr(v10 + 448))) = nox_frame_xxx_2598000 + uint32(randomIntMinMax(3, 10))
				*(*uint32)(unsafe.Pointer(uintptr(v10 + 444))) = nox_frame_xxx_2598000
				*(*uint16)(unsafe.Pointer(uintptr(v10 + 104))) = 20
				*(*uint8)(unsafe.Pointer(uintptr(v10 + 296))) = uint8(int8(randomIntMinMax(0, 6)))
				nox_xxx_sprite_45A110_drawable((*nox_drawable)(unsafe.Pointer(uintptr(v10))))
			}
			if func() int32 {
				p := &v16
				*p++
				return *p
			}() >= *memmap.PtrInt32(0x587000, 190108) {
				break
			}
			v3 = v14
		}
	}
	v11 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*108)))
	if (*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*3))-uint32(v11))*(*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*3))-uint32(v11))+(*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*4))-*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*109)))*(*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*4))-*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*109))) > 200 {
		v12 = &nox_xxx_spriteLoadAdd_45A360_drawable(*memmap.PtrInt32(6112660, 1522992), v11, int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*109)))).field_0
		*(*uint32)(unsafe.Add(unsafe.Pointer(v12), unsafe.Sizeof(uint32(0))*108)) = *(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*3))
		*(*uint32)(unsafe.Add(unsafe.Pointer(v12), unsafe.Sizeof(uint32(0))*109)) = *(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*4))
		nox_xxx_sprite_45A110_drawable((*nox_drawable)(unsafe.Pointer(v12)))
		*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*108)) = *(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*3))
		*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*109)) = *(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*4))
		nox_xxx_spriteTransparentDecay_49B950(v12, int32(nox_gameFPS/3))
	}
	return 1
}
