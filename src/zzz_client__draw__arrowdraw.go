package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"unsafe"
)

func nox_thing_arrow_draw(a1 *int32, dr *nox_drawable) int32 {
	var (
		v2 int32
		v3 int32
		v4 *uint32
		v5 int32
		a2 *uint32 = &dr.field_0
	)
	v2 = int32(*memmap.PtrUint32(6112660, 1313720))
	if *memmap.PtrUint32(6112660, 1313720) == 0 {
		v2 = nox_xxx_getTTByNameSpriteMB_44CFC0("ArrowTailLink")
		*memmap.PtrUint32(6112660, 1313720) = uint32(v2)
	}
	v3 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*81)))
	if (*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*3))-uint32(v3))*(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*3))-uint32(v3))+(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*4))-*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*82)))*(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*4))-*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*82))) > 200 {
		v4 = &nox_xxx_spriteLoadAdd_45A360_drawable(v2, v3, int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*82)))).field_0
		*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*108)) = *(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*3))
		*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*109)) = *(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*4))
		nox_xxx_sprite_45A110_drawable((*nox_drawable)(unsafe.Pointer(v4)))
		v5 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*4)))
		*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*81)) = *(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*3))
		*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*82)) = uint32(v5)
		nox_xxx_spriteTransparentDecay_49B950(v4, int32(nox_gameFPS/3))
	}
	return nox_thing_slave_draw(a1, dr)
}
func nox_thing_weak_arrow_draw(a1 *int32, dr *nox_drawable) int32 {
	var (
		v2 int32
		v3 int32
		v4 *uint32
		v5 int32
		a2 *uint32 = &dr.field_0
	)
	v2 = int32(*memmap.PtrUint32(6112660, 1313724))
	if *memmap.PtrUint32(6112660, 1313724) == 0 {
		v2 = nox_xxx_getTTByNameSpriteMB_44CFC0("WeakArrowTailLink")
		*memmap.PtrUint32(6112660, 1313724) = uint32(v2)
	}
	v3 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*81)))
	if (*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*3))-uint32(v3))*(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*3))-uint32(v3))+(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*4))-*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*82)))*(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*4))-*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*82))) > 200 {
		v4 = &nox_xxx_spriteLoadAdd_45A360_drawable(v2, v3, int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*82)))).field_0
		*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*108)) = *(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*3))
		*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*109)) = *(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*4))
		nox_xxx_sprite_45A110_drawable((*nox_drawable)(unsafe.Pointer(v4)))
		v5 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*4)))
		*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*81)) = *(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*3))
		*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*82)) = uint32(v5)
		nox_xxx_spriteTransparentDecay_49B950(v4, int32(nox_gameFPS/3))
	}
	return nox_thing_slave_draw(a1, dr)
}
func nox_thing_arrow_tail_link_draw(a1 *uint32, dr *nox_drawable) int32 {
	var (
		v2 int32
		v3 int32
		v4 int32
		v5 int32
		v6 int32
		v7 int32
		a2 int32 = int32(uintptr(unsafe.Pointer(dr)))
	)
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 16))))
	v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 12))) + *a1 - *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*4)))
	v4 = int32(uint32(v2) + *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*1)) - uint32(*(*int16)(unsafe.Pointer(uintptr(a2 + 106)))) - uint32(*(*int16)(unsafe.Pointer(uintptr(a2 + 104)))) - *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*5)))
	v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 436))) - uint32(v2))
	v6 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 432))) + *a1 - *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*4)))
	if *(*uint32)(unsafe.Pointer(uintptr(a2 + 356)))-nox_frame_xxx_2598000 > 0 {
		v7 = int32(((*(*uint32)(unsafe.Pointer(uintptr(a2 + 356))) - nox_frame_xxx_2598000) << 6) / uint32(int32(nox_gameFPS/3)))
		if v7 >= 64 {
			v7 = 63
		}
		nox_client_drawSetColor_434460(int32(*memmap.PtrUint32(6112660, uint32(v7*4)+0x1408F4)))
		nox_client_drawEnableAlpha_434560(1)
		nox_client_drawSetAlpha_434580(128)
		nox_client_drawAddPoint_49F500(v3, v4-4)
		nox_client_drawAddPoint_49F500(v6, v5+v4-4)
		nox_client_drawLineFromPoints_49E4B0()
		nox_client_drawEnableAlpha_434560(0)
	}
	return 1
}
func nox_thing_weak_arrow_tail_link_draw(a1 *uint32, dr *nox_drawable) int32 {
	var (
		v2 int32
		v3 int32
		v4 int32
		v5 int32
		v6 int32
		v7 int32
		a2 int32 = int32(uintptr(unsafe.Pointer(dr)))
	)
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 16))))
	v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 12))) + *a1 - *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*4)))
	v4 = int32(uint32(v2) + *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*1)) - uint32(*(*int16)(unsafe.Pointer(uintptr(a2 + 106)))) - uint32(*(*int16)(unsafe.Pointer(uintptr(a2 + 104)))) - *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*5)))
	v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 436))) - uint32(v2))
	v6 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 432))) + *a1 - *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*4)))
	if *(*uint32)(unsafe.Pointer(uintptr(a2 + 356)))-nox_frame_xxx_2598000 > 0 {
		v7 = int32(((*(*uint32)(unsafe.Pointer(uintptr(a2 + 356))) - nox_frame_xxx_2598000) << 6) / uint32(int32(nox_gameFPS/3)))
		if v7 >= 64 {
			v7 = 63
		}
		nox_client_drawSetColor_434460(int32(*memmap.PtrUint32(6112660, uint32(v7*4)+0x1409F4)))
		nox_client_drawEnableAlpha_434560(1)
		nox_client_drawSetAlpha_434580(128)
		nox_client_drawAddPoint_49F500(v3, v4-4)
		nox_client_drawAddPoint_49F500(v6, v5+v4-4)
		nox_client_drawLineFromPoints_49E4B0()
		nox_client_drawEnableAlpha_434560(0)
	}
	return 1
}
