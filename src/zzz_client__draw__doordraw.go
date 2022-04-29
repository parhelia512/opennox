package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	noxflags "github.com/noxworld-dev/opennox/v1/common/flags"
	"unsafe"
)

func nox_thing_door_draw(a1 *uint32, dr *nox_drawable) int32 {
	var (
		v2  *uint8
		v3  int32
		v4  int32
		v5  int32
		v6  int8
		v7  int8
		v9  int8
		v10 int8
		v11 int8
		v12 int32
		v13 int32
		v14 int32
		v15 int32
		a3  int2
	)
	nox_xxx_drawObject_4C4770_draw((*int32)(unsafe.Pointer(a1)), dr, int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(dr.field_76)), 4)))) + uint32(int32(dr.field_74_4)*4))))))
	if !noxflags.HasGame(noxflags.GameModeQuest) || int32(dr.field_108_1) != 1 {
		return 1
	}
	if dr.field_109 == 0 {
		dr.field_110 = uint32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg("DoorLockSilverSW"))))
		dr.field_109 = uint32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg("DoorLockSilverSE"))))
		dr.field_111 = uint32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg("DoorLockGoldSW"))))
		dr.field_112 = uint32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg("DoorLockGoldSE"))))
	}
	sub_4739E0(a1, (*int2)(unsafe.Pointer(&dr.pos)), &a3)
	a3.field_0 -= 64
	a3.field_4 -= 79
	v2 = (*uint8)(unsafe.Pointer(uintptr((cgoAsFunc(*(*uint32)(unsafe.Pointer(&dword_5d4594_3807156)), (*func(*int2) int32)(nil)))((*int2)(unsafe.Pointer(&dr.pos))))))
	sub_4345F0(1)
	*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v4))), 0)) = *(*uint8)(unsafe.Add(unsafe.Pointer(v2), 8))
	v3 = int32(*(*uint8)(unsafe.Add(unsafe.Pointer(v2), 4)))
	*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v5))), 0)) = *v2
	sub_433CD0(uint8(int8(v5)), uint8(int8(v3)), uint8(int8(v4)))
	v6 = int8(dr.field_74_4)
	if int32(v6) == 0 {
		v7 = int8(dr.field_108_2)
		v12 = a3.field_4 - 20
		a3.field_0 -= 15
		a3.field_4 -= 20
		if int32(v7) == 1 {
			nox_client_drawImageAt_47D2C0((*nox_video_bag_image_t)(unsafe.Pointer(uintptr(dr.field_110))), a3.field_0, v12)
		} else {
			nox_client_drawImageAt_47D2C0((*nox_video_bag_image_t)(unsafe.Pointer(uintptr(dr.field_111))), a3.field_0, v12)
		}
		return 1
	} else if int32(v6) == 8 {
		v9 = int8(dr.field_108_2)
		v13 = a3.field_4 - 20
		a3.field_0 += 15
		a3.field_4 -= 20
		if int32(v9) == 1 {
			nox_client_drawImageAt_47D2C0((*nox_video_bag_image_t)(unsafe.Pointer(uintptr(dr.field_109))), a3.field_0, v13)
		} else {
			nox_client_drawImageAt_47D2C0((*nox_video_bag_image_t)(unsafe.Pointer(uintptr(dr.field_112))), a3.field_0, v13)
		}
		return 1
	} else if int32(v6) == 16 {
		v10 = int8(dr.field_108_2)
		v14 = a3.field_4 + 2
		a3.field_0 += 8
		a3.field_4 += 2
		if int32(v10) == 1 {
			nox_client_drawImageAt_47D2C0((*nox_video_bag_image_t)(unsafe.Pointer(uintptr(dr.field_110))), a3.field_0, v14)
		} else {
			nox_client_drawImageAt_47D2C0((*nox_video_bag_image_t)(unsafe.Pointer(uintptr(dr.field_111))), a3.field_0, v14)
		}
		return 1
	} else {
		v11 = int8(dr.field_108_2)
		v15 = a3.field_4 + 2
		a3.field_0 -= 8
		a3.field_4 += 2
		if int32(v11) == 1 {
			nox_client_drawImageAt_47D2C0((*nox_video_bag_image_t)(unsafe.Pointer(uintptr(dr.field_109))), a3.field_0, v15)
		} else {
			nox_client_drawImageAt_47D2C0((*nox_video_bag_image_t)(unsafe.Pointer(uintptr(dr.field_112))), a3.field_0, v15)
		}
		return 1
	}
}
func nox_things_door_draw_parse(obj *nox_thing, f *nox_memfile, attr_value *byte) bool {
	obj.draw_func = nox_thing_door_draw
	var v3 unsafe.Pointer = nox_xxx_spriteLoadStaticRandomData_44C000(attr_value, f)
	obj.field_5c = v3
	return v3 != nil
}
