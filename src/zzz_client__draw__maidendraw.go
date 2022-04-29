package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	noxflags "github.com/noxworld-dev/opennox/v1/common/flags"
	"unsafe"
)

func nox_thing_maiden_draw(a1 *uint32, dr *nox_drawable) int32 {
	if !noxflags.HasGame(noxflags.GameFlag22) {
		var v9 *byte = (*byte)(unsafe.Pointer(nox_npc_by_id(int32(dr.field_32))))
		if v9 == nil {
			return 1
		}
		var v10 int32 = 0
		var v11 *int32 = (*int32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v9), 8))))
		var v12 int32
		for {
			v12 = v10 + 1
			nox_xxx_drawPlayer_4341D0(v10+1, *v11)
			v10 = v12
			v11 = (*int32)(unsafe.Add(unsafe.Pointer(v11), unsafe.Sizeof(int32(0))*1))
			if v12 >= 6 {
				break
			}
		}
		return nox_thing_monster_draw((*int32)(unsafe.Pointer(a1)), dr)
	}
	var v2 *nox_object_t = noxServer.firstServerObject()
	if v2 == nil {
		return nox_thing_monster_draw((*int32)(unsafe.Pointer(a1)), dr)
	}
	var v3 int32
	for {
		v3 = int32(v2.extent)
		if dr.field_32 == uint32(v3) {
			break
		}
		v2 = v2.Next()
		if v2 == nil {
			return nox_thing_monster_draw((*int32)(unsafe.Pointer(a1)), dr)
		}
	}
	var v5 int32 = int32(uintptr(v2.data_update))
	var v6 int32 = 0
	var v7 *uint8 = (*uint8)(unsafe.Pointer(uintptr(v5 + 2077)))
	var v8 int32
	for {
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v3))), 0)) = *(*uint8)(unsafe.Add(unsafe.Pointer(v7), 1))
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v5))), 0)) = *((*uint8)(unsafe.Add(unsafe.Pointer(v7), -1)))
		v8 = v6 + 1
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v6))), 0)) = *v7
		sub_4340A0(v8, v5, v6, v3)
		v6 = v8
		v7 = (*uint8)(unsafe.Add(unsafe.Pointer(v7), 3))
		if v8 >= 6 {
			break
		}
	}
	return nox_thing_monster_draw((*int32)(unsafe.Pointer(a1)), dr)
}
func nox_things_maiden_draw_parse(obj *nox_thing, f *nox_memfile, attr_value *byte) bool {
	var result int32 = bool2int(nox_things_monster_draw_parse(obj, f, attr_value))
	obj.draw_func = nox_thing_maiden_draw
	return result != 0
}
