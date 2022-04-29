package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"math"
	"unsafe"
)

func nox_thing_magic_draw(a1 *int32, dr *nox_drawable) int32 {
	var (
		v2  int32
		v3  int32
		v4  int32
		v5  int32
		v6  int32
		v7  int32
		v9  float32
		v10 int2
		a2  int32 = int32(uintptr(unsafe.Pointer(dr)))
	)
	if dword_5d4594_1313804 == 0 {
		*memmap.PtrUint32(6112660, 1313808) = nox_color_rgb_4344A0(0, 200, math.MaxUint8)
		*memmap.PtrUint32(6112660, 1313812) = nox_color_rgb_4344A0(math.MaxUint8, math.MaxUint8, 50)
		dword_5d4594_1313804 = 1
	}
	v2 = *a1
	v3 = *(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*1))
	v4 = int32(uint32(*a1) + *(*uint32)(unsafe.Pointer(uintptr(a2 + 12))) - uint32(*(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*4))))
	v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 16))) - uint32(*(*int16)(unsafe.Pointer(uintptr(a2 + 106)))) - uint32(*(*int16)(unsafe.Pointer(uintptr(a2 + 104)))) - uint32(*(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*5))))
	v10.field_0 = v4
	v6 = v3 + v5
	v10.field_4 = v6
	if v4-10 >= v2 && v6-10 >= v3 && v4+10 < *(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*2)) && v6+10 < *(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*3)) {
		v7 = randomIntMinMax(1, 4)
		sub_4B6720(&v10, *memmap.PtrInt32(6112660, 1313808), v7*2+1, int8((v7>>1)+3))
		nox_client_drawSetColor_434460(int32(nox_color_white_2523948))
		nox_client_drawRectFilledOpaque_49CE30(v10.field_0-(v7>>1), v10.field_4-(v7>>1), v7, v7)
		nox_xxx_spriteChangeLightColor_484BE0((*uint32)(unsafe.Pointer(uintptr(a2+136))), 200, 200, math.MaxUint8)
		v9 = float32(nox_common_randomFloatXxx_416090(0.0, 100.0))
		nox_xxx_spriteChangeIntensity_484D70_light_intensity(a2+136, v9)
	}
	return 1
}
func nox_thing_magic_missle_draw(a1 *int32, dr *nox_drawable) int32 {
	var (
		v2  int32
		v3  int32
		v4  int32
		v5  int32
		v6  int32
		v7  int32
		v9  float32
		v10 int2
		a2  int32 = int32(uintptr(unsafe.Pointer(dr)))
	)
	if dword_5d4594_1313804 == 0 {
		*memmap.PtrUint32(6112660, 1313808) = nox_color_rgb_4344A0(0, 200, math.MaxUint8)
		*memmap.PtrUint32(6112660, 1313812) = nox_color_rgb_4344A0(math.MaxUint8, math.MaxUint8, 50)
		dword_5d4594_1313804 = 1
	}
	v2 = *a1
	v3 = *(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*1))
	v4 = int32(uint32(*a1) + *(*uint32)(unsafe.Pointer(uintptr(a2 + 12))) - uint32(*(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*4))))
	v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 16))) - uint32(*(*int16)(unsafe.Pointer(uintptr(a2 + 106)))) - uint32(*(*int16)(unsafe.Pointer(uintptr(a2 + 104)))) - uint32(*(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*5))))
	v10.field_0 = v4
	v6 = v3 + v5
	v10.field_4 = v6
	if v4-10 >= v2 && v6-10 >= v3 && v4+10 < *(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*2)) && v6+10 < *(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*3)) {
		v7 = randomIntMinMax(1, 4)
		sub_4B6720(&v10, *memmap.PtrInt32(6112660, 1313812), v7*2+1, int8((v7>>1)+3))
		nox_client_drawSetColor_434460(int32(nox_color_white_2523948))
		nox_client_drawRectFilledOpaque_49CE30(v10.field_0-(v7>>1), v10.field_4-(v7>>1), v7, v7)
		nox_xxx_spriteChangeLightColor_484BE0((*uint32)(unsafe.Pointer(uintptr(a2+136))), math.MaxUint8, 180, 50)
		v9 = float32(nox_common_randomFloatXxx_416090(0.0, 100.0))
		nox_xxx_spriteChangeIntensity_484D70_light_intensity(a2+136, v9)
	}
	return 1
}
func nox_thing_magic_missle_tail_link_draw(a1 *uint32, dr *nox_drawable) int32 {
	var (
		v2  int32
		v3  int32
		v4  int32
		v5  int32
		v6  int32
		v7  int32
		v8  int32
		v10 float32
		v11 int32
		v12 int32
		v13 int32
		v14 int32
		a2  int32 = int32(uintptr(unsafe.Pointer(dr)))
	)
	v2 = a2
	v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 12))) + *a1 - *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*4)))
	v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 16))) + *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*1)) - uint32(*(*int16)(unsafe.Pointer(uintptr(a2 + 106)))) - uint32(*(*int16)(unsafe.Pointer(uintptr(a2 + 104)))) - *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*5)))
	v12 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*1)) - uint32(*(*int16)(unsafe.Pointer(uintptr(a2 + 106)))) - uint32(*(*int16)(unsafe.Pointer(uintptr(a2 + 104)))) - *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*5)) + *(*uint32)(unsafe.Pointer(uintptr(a2 + 436))))
	v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 356))) - nox_frame_xxx_2598000)
	v11 = int32(*a1 - *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*4)) + *(*uint32)(unsafe.Pointer(uintptr(a2 + 432))))
	v14 = v5
	if v5 > 0 {
		v6 = (v5 << 6) / int32(nox_gameFPS/3)
		v13 = int32(nox_gameFPS / 3)
		if v6 >= 64 {
			v6 = 63
		}
		v7 = int32(*memmap.PtrUint32(6112660, uint32(v6*4)+0x1407F4))
		v8 = v2 + 136
		nox_xxx_spriteChangeLightColor_484BE0((*uint32)(unsafe.Pointer(uintptr(v2+136))), math.MaxUint8, 128, 50)
		v10 = float32(float64(v14) * 20.0 / float64(v13))
		nox_xxx_spriteChangeIntensity_484D70_light_intensity(v8, v10)
		nox_client_drawSetColor_434460(v7)
		nox_client_drawEnableAlpha_434560(1)
		nox_client_drawSetAlpha_434580(128)
		nox_client_drawAddPoint_49F500(v3, v4)
		nox_client_drawAddPoint_49F500(v11, v12)
		nox_client_drawLineFromPoints_49E4B0()
		nox_client_drawEnableAlpha_434560(0)
	}
	return 1
}
func nox_thing_magic_tail_link_draw(a1 *uint32, dr *nox_drawable) int32 {
	var (
		v2  int32
		v3  int32
		v4  int32
		v5  int32
		v6  bool
		v7  int32
		v8  uint32
		v9  int32
		v10 int32
		v12 float32
		v13 int32
		v14 int32
		v15 int32
		a2  int32 = int32(uintptr(unsafe.Pointer(dr)))
	)
	v2 = a2
	v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 12))) + *a1 - *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*4)))
	v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 16))) + *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*1)) - uint32(*(*int16)(unsafe.Pointer(uintptr(a2 + 106)))) - uint32(*(*int16)(unsafe.Pointer(uintptr(a2 + 104)))) - *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*5)))
	v13 = int32(*a1 - *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*4)) + *(*uint32)(unsafe.Pointer(uintptr(a2 + 432))))
	v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 356))) - nox_frame_xxx_2598000)
	v6 = *(*uint32)(unsafe.Pointer(uintptr(a2 + 356))) == nox_frame_xxx_2598000
	v14 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*1)) - uint32(*(*int16)(unsafe.Pointer(uintptr(a2 + 106)))) - uint32(*(*int16)(unsafe.Pointer(uintptr(a2 + 104)))) - *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*5)) + *(*uint32)(unsafe.Pointer(uintptr(a2 + 436))))
	v15 = v5
	if v5 >= 0 && !v6 {
		v8 = uint32(v5 << 6)
		v7 = int32(uint32(v5<<6) / nox_gameFPS)
		if int32(v8/nox_gameFPS) >= 64 {
			v7 = 63
		}
		v9 = int32(*memmap.PtrUint32(6112660, uint32(v7*4)+1312500))
		v10 = v2 + 136
		nox_xxx_spriteChangeLightColor_484BE0((*uint32)(unsafe.Pointer(uintptr(v2+136))), 128, 128, math.MaxUint8)
		v12 = float32(float64(v15) * 20.0 / float64(int32(nox_gameFPS)))
		nox_xxx_spriteChangeIntensity_484D70_light_intensity(v10, v12)
		nox_client_drawSetColor_434460(v9)
		nox_client_drawEnableAlpha_434560(1)
		nox_client_drawSetAlpha_434580(128)
		nox_client_drawAddPoint_49F500(v3, v4)
		nox_client_drawAddPoint_49F500(v13, v14)
		nox_client_drawLineFromPoints_49E4B0()
		nox_client_drawEnableAlpha_434560(0)
	}
	return 1
}
func nox_thing_drain_mana_draw() int32 {
	return 1
}
