package opennox

import (
	"github.com/gotranspile/cxgo/runtime/cmath"
	"github.com/gotranspile/cxgo/runtime/libc"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"math"
	"unsafe"
)

func sub_4CCEA0(a1 *uint32, a2 int32) {
	var (
		v2  *uint32
		v3  int32
		v4  int32
		v5  int32
		v6  int32
		v8  int32
		v9  int32
		v11 int32
		v12 int32
		v13 int32
		v14 int32
		v15 int32
		v16 int32
		v17 float2
		v18 int32
	)
	if *memmap.PtrUint32(6112660, 1522964) == 0 {
		*memmap.PtrUint32(6112660, 1522964) = uint32(nox_xxx_getTTByNameSpriteMB_44CFC0("Spark"))
	}
	v2 = a1
	v3 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*4)) - *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*9)))
	v4 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*3)) - *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*8)))
	v5 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*4)) - *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*9)))
	v18 = v4
	v15 = v3
	v6 = int32(cmath.Abs(int64(v4)) + cmath.Abs(int64(v5)))
	if v6/7 > 0 {
		v16 = v6 / 7
		for {
			v8 = randomIntMinMax(0, 100)
			v9 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*8)) + uint32(v4*v8/100))
			v11 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*9)) + uint32((v3*v8)/100))
			v12 = int32(uintptr(unsafe.Pointer(nox_xxx_spriteLoadAdd_45A360_drawable(*memmap.PtrInt32(6112660, 1522964), v9, v11))))
			if v12 != 0 {
				*(*uint32)(unsafe.Pointer(uintptr(v12 + 432))) = uint32(v9 << 12)
				*(*uint32)(unsafe.Pointer(uintptr(v12 + 436))) = uint32(v11 << 12)
				v17.field_0 = float32(float64(-v18))
				v17.field_4 = float32(float64(-v15))
				v13 = randomIntMinMax(-25, 25)
				v14 = nox_xxx_math_509ED0(&v17) + v13
				if v14 < 0 {
					v14 += int32(uint32(math.MaxUint8-v14) >> 8 << 8)
				}
				*(*uint8)(unsafe.Pointer(uintptr(v12 + 299))) = uint8(int8(v14))
				*(*uint32)(unsafe.Pointer(uintptr(v12 + 440))) = uint32(a2 * randomIntMinMax(100, 300))
				*(*uint32)(unsafe.Pointer(uintptr(v12 + 448))) = nox_frame_xxx_2598000 + uint32(randomIntMinMax(30, 45))
				*(*uint32)(unsafe.Pointer(uintptr(v12 + 444))) = nox_frame_xxx_2598000
				*(*uint16)(unsafe.Pointer(uintptr(v12 + 104))) = 28
				*(*uint16)(unsafe.Pointer(uintptr(v12 + 106))) = 0
				*(*uint8)(unsafe.Pointer(uintptr(v12 + 296))) = uint8(int8(randomIntMinMax(-2, 4)))
				nox_xxx_sprite_45A110_drawable((*nox_drawable)(unsafe.Pointer(uintptr(v12))))
			}
			v16--
			if v16 == 0 {
				break
			}
			v3 = v15
			v4 = v18
		}
	}
}
