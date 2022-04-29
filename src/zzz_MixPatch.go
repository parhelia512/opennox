package opennox

import "unsafe"

func sub_980523(unit *nox_object_t) {
	if unit == nil {
		return
	}
	for it := (*nox_object_t)(unit.field_126); it != nil; it = it.field_124 {
		if (it.obj_class&0x2000000) != 0 && (it.obj_flags&256) != 0 {
			if uint32(nox_xxx_unitArmorInventoryEquipFlags_415C70(it))&0x3000000 != 0 {
				*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr((uint32(uintptr(unit.data_update))) + 276))) + 2500))) = uint32(uintptr(unsafe.Pointer(it)))
			}
		}
	}
}
func sub_9805EB(unit *nox_object_t) *nox_object_t {
	if unit == nil {
		return nil
	}
	for it := (*nox_object_t)(unit.field_126); it != nil; it = it.field_124 {
		if (it.obj_class&0x2000000) != 0 && it.obj_flags == 16 {
			return it
		}
	}
	return nil
}
func mix_recvfrom(s nox_socket_t, buf *byte, len_ int32, from *nox_net_sockaddr) int32 {
	var result int32 = nox_net_recvfrom(s, unsafe.Pointer(buf), uint32(len_), (*nox_net_sockaddr_in)(unsafe.Pointer(from)))
	if int32(*(*uint16)(unsafe.Pointer(buf))) != 0xF13A {
		return result
	}
	return MixRecvFromReplacer(s, buf, len_, from)
}
