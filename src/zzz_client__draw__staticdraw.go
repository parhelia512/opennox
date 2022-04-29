package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	"github.com/noxworld-dev/opennox/v1/common/alloc"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"unsafe"
)

func nox_thing_static_draw(a1 *uint32, dr *nox_drawable) int32 {
	if (dr.flags28&0x40000) == 0 || dr.flags30&0x1000000 != 0 {
		nox_xxx_drawObject_4C4770_draw((*int32)(unsafe.Pointer(a1)), dr, int32(*(*uint32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(dr.field_76)), 4))))))
	}
	return 1
}
func nox_thing_static_random_draw(a1 *uint32, dr *nox_drawable) int32 {
	nox_xxx_drawObject_4C4770_draw((*int32)(unsafe.Pointer(a1)), dr, int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(dr.field_76)), 4)))) + dr.field_77*4)))))
	return 1
}
func nox_things_static_draw_parse(obj *nox_thing, f *nox_memfile, attr_value *byte) bool {
	var (
		a3   *uint8 = (*uint8)(unsafe.Pointer(attr_value))
		v6   *byte
		v8   int32
		v11  uint8
		data *uint32 = (*uint32)(alloc.Calloc(1, 8))
	)
	if data == nil {
		return false
	}
	*(*uint32)(unsafe.Add(unsafe.Pointer(data), unsafe.Sizeof(uint32(0))*0)) = 8
	v6 = (*byte)(unsafe.Pointer(a3))
	v8 = int32(f.ReadU32())
	*a3 = *memmap.PtrUint8(6112660, 830856)
	if v8 == -1 {
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a3))), 0)) = f.ReadU8()
		v11 = f.ReadU8()
		nox_memfile_read(unsafe.Pointer(v6), 1, int32(v11), f)
		*(*byte)(unsafe.Add(unsafe.Pointer(v6), v11)) = 0
	}
	*(*uint32)(unsafe.Add(unsafe.Pointer(data), unsafe.Sizeof(uint32(0))*1)) = uint32(uintptr(unsafe.Pointer(nox_xxx_readImgMB_42FAA0(v8, int8(uintptr(unsafe.Pointer(a3))), v6))))
	obj.draw_func = nox_thing_static_draw
	obj.field_5c = unsafe.Pointer(data)
	return true
}
func nox_things_static_random_draw_parse(obj *nox_thing, f *nox_memfile, attr_value *byte) bool {
	obj.draw_func = nox_thing_static_random_draw
	obj.field_5c = nox_xxx_spriteLoadStaticRandomData_44C000(attr_value, f)
	return obj.field_5c != nil
}
