package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"math"
	"unsafe"
)

const (
	DDSCAPS_OFFSCREENPLAIN = 1
	DDSCAPS_SYSTEMMEMORY   = 2
	DDSCAPS_VIDEOMEMORY    = 4
	DDSCAPS_PRIMARYSURFACE = 8
)

var byte_5D4594_3804364 [160]uint8 = [160]uint8{}
var dword_973C64 uint32
var dword_6F7C40 func() int16
var dword_6F7C34 func() int16
var nox_video_renderTargetFlags int32 = 0
var nox_video_windowsPlatformVersion int32 = 0
var dword_975240 func(uint32, *uint32, *uint32, *uint32)
var dword_975380 func(uint32, uint32, uint32) int32
var nox_draw_colors_r_3804672 *uint8 = nil
var nox_draw_colors_g_3804656 *uint8 = nil
var nox_draw_colors_b_3804664 *uint8 = nil

func sub_4340A0(a1 int32, a2 int32, a3 int32, a4 int32) {
	var v4 uint64
	_ = v4
	var v5 int32
	var v6 uint64
	_ = v6
	var v7 uint64
	_ = v7
	a2 = a2 & math.MaxUint8
	a3 = a3 & math.MaxUint8
	a4 = a4 & math.MaxUint8
	v4 = uint64(a1)
	if a1 >= 0 && a1 < 16 {
		v5 = int32(uint32(a1*48) + uint32(uintptr(unsafe.Pointer(&nox_draw_curDrawData_3799572.field_66))))
		*(*uint32)(unsafe.Pointer(uintptr(v5 + 32))) = uint32(uint8(int8(a4)))
		*(*uint32)(unsafe.Pointer(uintptr(v5 + 24))) = uint32(uint8(int8(a2)))
		*(*uint32)(unsafe.Pointer(uintptr(v5 + 28))) = uint32(uint8(int8(a3)))
		nox_color_rgb_func(uint8(int8(a2)), uint8(int8(a3)), uint8(int8(a4)), (*uint32)(unsafe.Pointer(uintptr(v5+40))))
	}
}
func sub_4347F0(a1 *byte, a2 int32) int32 {
	if a2 <= 256 {
		sub_435120((*uint8)(memmap.PtrOff(0x973F18, 3880)), a1)
		sub_4353F0()
		sub_435040()
		sub_434F00()
	}
	return bool2int(nox_xxx_makeFillerColor_48BDE0())
}
func sub_4348C0() int32 {
	return 0
}
func sub_434FB0() int32 {
	return 1
}
func sub_4352E0() {
}
func sub_444F90() HDC {
	panic("abort")
	return 0
}
func sub_444FC0(a1 HDC) {
	panic("abort")
}
func sub_4353C0() {
}
func sub_4353F0() {
}
func sub_4354F0() {
}
func sub_435550() {
}
func sub_433CD0(a1 uint8, a2 uint8, a3 uint8) int32 {
	var v5 int64
	_ = v5
	var v6 uint64
	_ = v6
	var v7 uint64
	_ = v7
	var result int32
	nox_draw_curDrawData_3799572.field_24 = uint32(a1)
	nox_draw_curDrawData_3799572.field_25 = uint32(a2)
	nox_draw_curDrawData_3799572.field_26 = uint32(a3)
	nox_draw_curDrawData_3799572.field_16 = uint32(bool2int(int32(a1) == math.MaxUint8 && int32(a2) == math.MaxUint8 && int32(a3) == math.MaxUint8))
	result = dword_975380(uint32(a1), uint32(a2), uint32(a3))
	*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&nox_draw_curDrawData_3799572.field_258))), unsafe.Sizeof(uint16(0))*1)) = uint16(int16(result))
	return result
}
func sub_433E40(a1 int32) int32 {
	var (
		v2 int32
		v3 int32
	)
	dword_975240(uint32(a1), (*uint32)(unsafe.Pointer(&v3)), (*uint32)(unsafe.Pointer(&v2)), (*uint32)(unsafe.Pointer(&a1)))
	return sub_433CD0(uint8(int8(v3)), uint8(int8(v2)), uint8(int8(a1)))
}
func sub_433ED0(a1 int32) {
	var (
		v2 int32
		v3 int32
	)
	dword_975240(uint32(a1), (*uint32)(unsafe.Pointer(&v3)), (*uint32)(unsafe.Pointer(&v2)), (*uint32)(unsafe.Pointer(&a1)))
	sub_433E80(uint8(int8(v3)), uint8(int8(v2)), uint8(int8(a1)))
}
func sub_434040(a1 int32) {
	var (
		v2 int32
		v3 int32
	)
	dword_975240(uint32(a1), (*uint32)(unsafe.Pointer(&v3)), (*uint32)(unsafe.Pointer(&v2)), (*uint32)(unsafe.Pointer(&a1)))
	nox_xxx_drawMakeRGB_433F10(uint8(int8(v3)), uint8(int8(v2)), uint8(int8(a1)))
}
func nox_xxx_drawPlayer_4341D0(a1 int32, a2 int32) int32 {
	var (
		result int32
		v3     int32
		v4     int32
		v5     int32
		v6     uint64
	)
	_ = v6
	var v7 uint64
	_ = v7
	var v8 int64
	_ = v8
	var v9 uint64
	_ = v9
	var v10 int32
	result = a1
	if a1 >= 0 && a1 < 16 {
		v3 = a2
		v4 = int32(uint32(a1*48) + uint32(uintptr(unsafe.Pointer(&nox_draw_curDrawData_3799572.field_66))))
		result = int32(*(*uint32)(unsafe.Pointer(uintptr(uint32(a1*48) + uint32(uintptr(unsafe.Pointer(&nox_draw_curDrawData_3799572.field_76)))))))
		if a2 != result {
			dword_975240(uint32(a2), (*uint32)(unsafe.Pointer(&a1)), (*uint32)(unsafe.Pointer(&a2)), (*uint32)(unsafe.Pointer(&v10)))
			*(*uint32)(unsafe.Pointer(uintptr(v4 + 24))) = uint32(uint8(int8(a1)))
			*(*uint32)(unsafe.Pointer(uintptr(v4 + 28))) = uint32(uint8(int8(a2)))
			v5 = int32(uint8(int8(v10)))
			*(*uint32)(unsafe.Pointer(uintptr(v4 + 40))) = uint32(v3)
			*(*uint32)(unsafe.Pointer(uintptr(v4 + 32))) = uint32(v5)
		}
	}
	return result
}
func sub_434480(a1 int32, a2 *int32, a3 *int32, a4 *int32) {
	dword_975240(uint32(a1), (*uint32)(unsafe.Pointer(a2)), (*uint32)(unsafe.Pointer(a3)), (*uint32)(unsafe.Pointer(a4)))
}
func sub_434AA0(a1 int32, a2 int32, a3 int32) int32 {
	return dword_975380(uint32(a1), uint32(a2), uint32(a3))
}
func sub_434AC0(a1 int32) int32 {
	var (
		v2 int32
		v3 int32
	)
	dword_975240(uint32(a1), (*uint32)(unsafe.Pointer(&v3)), (*uint32)(unsafe.Pointer(&v2)), (*uint32)(unsafe.Pointer(&a1)))
	return dword_975380(uint32(v3), uint32(v2), uint32(a1))
}
func sub_434B60() int32 {
	var (
		v0 int32
		v1 *pixel888
	)
	_ = v1
	var v2 int32
	var v3 *byte
	var v4 int64
	var v13 int32
	var v14 [1536]byte
	var v19 [256]pixel888
	v0 = nox_video_getGammaSetting_434B00()
	v1 = &v19[0]
	v2 = 0
	v3 = &v14[512]
	v13 = 0
	for {
		if v0 == 1 {
			*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v4))), unsafe.Sizeof(uint16(0))*0)) = uint16(int16(int32(uint16(int16(v2))) << 8))
		} else {
			v4 = int64(math.Pow(float64(v13)*0.00392156862745098, 1.0/(float64(v0-1)*0.1666666666666667+1.0)) * 65535.0)
		}
		*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v3))), unsafe.Sizeof(uint16(0))*256))) = uint16(int16(v4))
		*(*uint16)(unsafe.Pointer(v3)) = uint16(int16(v4))
		*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v3))), -int(unsafe.Sizeof(uint16(0))*256)))) = uint16(int16(v4))
		v2++
		v3 = (*byte)(unsafe.Add(unsafe.Pointer(v3), 2))
		v13 = v2
		if v2 >= 256 {
			break
		}
	}
	return 0
}
func sub_4B0300(a1 *byte) int32 {
	var result int32
	result = int32(*memmap.PtrUint32(6112660, 1311928))
	if *memmap.PtrInt32(6112660, 1311928) < 2 {
		libc.StrNCpy((*byte)(memmap.PtrOff(6112660, *memmap.PtrUint32(6112660, 1311928)*260+1311940)), a1, 260)
		result = int32(func() uint32 {
			p := memmap.PtrUint32(6112660, 1311928)
			*p++
			return *p
		}())
	}
	return result
}
