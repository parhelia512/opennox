package opennox

import (
	noxflags "github.com/noxworld-dev/opennox/v1/common/flags"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"math"
	"unsafe"
)

func nox_thing_glyph_draw(a1 *int32, dr *nox_drawable) int32 {
	var (
		v3 int8
		v4 int32
		v5 int32
		v6 int32
		v7 int32
		a2 *uint32 = &dr.field_0
	)
	if !noxflags.HasGame(noxflags.GameClient) || *memmap.PtrUint32(0x852978, 8) == 0 {
		goto LABEL_10
	}
	if *(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*30))&0x40000000 != 0 {
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a2))), 0)) = math.MaxUint8
	LABEL_10:
		nox_client_drawEnableAlpha_434560(1)
		nox_client_drawSetAlpha_434580(uint8(uintptr(unsafe.Pointer(a2))))
		v7 = nox_thing_animate_draw((*uint32)(unsafe.Pointer(a1)), dr)
		nox_client_drawEnableAlpha_434560(0)
		nox_xxx_draw_434600(0)
		return v7
	}
	if nox_xxx_spriteCheckFlag31_4356C0((*nox_drawable)(unsafe.Pointer(uintptr(*memmap.PtrInt32(0x852978, 8)))), 21) {
		nox_xxx_draw_434600(1)
		sub_433E40(*memmap.PtrInt32(0x8531A0, 2572))
		v3 = -1
	LABEL_9:
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a2))), 0)) = uint8(v3)
		goto LABEL_10
	}
	v4 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*3)) - *(*uint32)(unsafe.Pointer(uintptr(*memmap.PtrUint32(0x852978, 8) + 12))))
	v5 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*4)) - *(*uint32)(unsafe.Pointer(uintptr(*memmap.PtrUint32(0x852978, 8) + 16))))
	v6 = v4*v4 + v5*v5
	if v6 < 22500 {
		v3 = int8(int32(-56 - v6*200/22500))
		goto LABEL_9
	}
	return 1
}
