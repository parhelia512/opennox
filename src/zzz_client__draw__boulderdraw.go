package opennox

import "unsafe"

func nox_thing_boulder_draw(a1 *int32, dr *nox_drawable) int32 {
	var (
		v2 int32
		v3 int32
		v4 int32
		v5 uint32
		v6 int32
		a2 int32 = int32(uintptr(unsafe.Pointer(dr)))
	)
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 304))))
	if *(*uint32)(unsafe.Pointer(uintptr(a2 + 432))) == 0 && *(*uint32)(unsafe.Pointer(uintptr(a2 + 436))) == 0 {
		*(*uint32)(unsafe.Pointer(uintptr(a2 + 432))) = *(*uint32)(unsafe.Pointer(uintptr(a2 + 12)))
		*(*uint32)(unsafe.Pointer(uintptr(a2 + 436))) = *(*uint32)(unsafe.Pointer(uintptr(a2 + 16)))
	}
	v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 12))) - *(*uint32)(unsafe.Pointer(uintptr(a2 + 432))))
	v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 16))) - *(*uint32)(unsafe.Pointer(uintptr(a2 + 436))))
	if v3*v3+v4*v4 >= 100 {
		if v3 <= 0 {
			if v4 > 0 {
				*(*uint32)(unsafe.Pointer(uintptr(a2 + 444))) = 0
			LABEL_8:
				v5 = *(*uint32)(unsafe.Pointer(uintptr(a2 + 440))) + 1
				*(*uint32)(unsafe.Pointer(uintptr(a2 + 440))) = v5
				if v5 >= 16 {
					*(*uint32)(unsafe.Pointer(uintptr(a2 + 440))) = 0
				}
				goto LABEL_17
			}
			*(*uint32)(unsafe.Pointer(uintptr(a2 + 444))) = 16
		} else {
			if v4 > 0 {
				*(*uint32)(unsafe.Pointer(uintptr(a2 + 444))) = 16
				goto LABEL_8
			}
			*(*uint32)(unsafe.Pointer(uintptr(a2 + 444))) = 0
		}
		v6 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 440))))
		if v6 != 0 {
			*(*uint32)(unsafe.Pointer(uintptr(a2 + 440))) = uint32(v6 - 1)
		} else {
			*(*uint32)(unsafe.Pointer(uintptr(a2 + 440))) = 15
		}
	LABEL_17:
		*(*uint32)(unsafe.Pointer(uintptr(a2 + 432))) = *(*uint32)(unsafe.Pointer(uintptr(a2 + 12)))
		*(*uint32)(unsafe.Pointer(uintptr(a2 + 436))) = *(*uint32)(unsafe.Pointer(uintptr(a2 + 16)))
	}
	nox_xxx_drawObject_4C4770_draw(a1, dr, int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 4))) + (*(*uint32)(unsafe.Pointer(uintptr(a2 + 440)))+*(*uint32)(unsafe.Pointer(uintptr(a2 + 444))))*4)))))
	return 1
}
