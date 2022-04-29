package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"unsafe"
)

func sub_4A22A0(a1 int32, a2 *int32) int32 {
	var (
		v2    int32
		v3    int32
		v4    *uint8
		v5    int32
		v6    int32
		v7    bool
		v8    int32
		v9    int32
		v10   *uint8
		v11   int32
		xLeft int32
		yTop  int32
	)
	nox_xxx_bookGet_430B40_get_mouse_prev_seq()
	nox_client_wndGetPosition_46AA60((*nox_window)(unsafe.Pointer(uintptr(a1))), (*uint32)(unsafe.Pointer(&xLeft)), (*uint32)(unsafe.Pointer(&yTop)))
	if int32(int8(*(*uint8)(unsafe.Pointer(uintptr(a1 + 4))))) >= 0 {
		if uint32(*(*int32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(int32(0))*5))) != 0x80000000 {
			nox_client_drawSetColor_434460(*(*int32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(int32(0))*5)))
			nox_client_drawRectFilledOpaque_49CE30(xLeft, yTop, int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 8)))), int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 12)))))
		}
	} else {
		v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 100))))
		xLeft += int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 96))))
		v3 = v2 + yTop
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v2))), 0)) = *(*uint8)(unsafe.Pointer(uintptr(a1 + 36)))
		yTop = v3
		if v2&2 != 0 {
			nox_client_drawImageAt_47D2C0((*nox_video_bag_image_t)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(int32(0))*10))))), xLeft, v3)
		} else {
			nox_client_drawImageAt_47D2C0((*nox_video_bag_image_t)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(int32(0))*6))))), xLeft, v3)
		}
	}
	if *memmap.PtrUint32(0x587000, 168836) != 0 {
		v4 = (*uint8)(memmap.PtrOff(0x587000, 168868))
		for {
			v5 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v4))), unsafe.Sizeof(uint32(0))*2))))
			if v5 != 0 {
				*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v4))), unsafe.Sizeof(uint32(0))*2))) = uint32(v5 - 1)
			}
			v6 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v4))), unsafe.Sizeof(uint32(0))*1))))
			if v6 != 0 {
				*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v4))), unsafe.Sizeof(uint32(0))*1))) = uint32(v6 - 1)
				*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v4))), unsafe.Sizeof(uint32(0))*2))) = uint32(randomIntMinMax(60, 120))
			}
			v7 = func() uint32 {
				p := (*uint32)(unsafe.Pointer(v4))
				x := *p
				*p--
				return x
			}() == 1
			v8 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v4))), -int(unsafe.Sizeof(uint32(0))*5)))))
			if v7 {
				if v8 != 0 {
					*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v4))), -int(unsafe.Sizeof(uint32(0))*5)))) = 0
					*(*uint32)(unsafe.Pointer(v4)) = uint32(randomIntMinMax(int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v4))), -int(unsafe.Sizeof(uint32(0))*4))))), int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v4))), -int(unsafe.Sizeof(uint32(0))*3)))))))
					*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v4))), unsafe.Sizeof(uint32(0))*2))) = uint32(randomIntMinMax(60, 90))
				} else {
					*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v4))), -int(unsafe.Sizeof(uint32(0))*5)))) = 1
					*(*uint32)(unsafe.Pointer(v4)) = uint32(randomIntMinMax(int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v4))), -int(unsafe.Sizeof(uint32(0))*2))))), int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v4))), -int(unsafe.Sizeof(uint32(0))*1)))))))
				}
			} else if v8 == 0 && *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v4))), unsafe.Sizeof(uint32(0))*2))) == 0 && *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v4))), unsafe.Sizeof(uint32(0))*1))) == 0 && randomIntMinMax(0, 100) > 75 {
				*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v4))), unsafe.Sizeof(uint32(0))*1))) = uint32(randomIntMinMax(4, 8))
			}
			v9 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v4))), unsafe.Sizeof(uint32(0))*4))))
			v4 = (*uint8)(unsafe.Add(unsafe.Pointer(v4), 48))
			if v9 == 0 {
				break
			}
		}
	}
	if *memmap.PtrUint32(0x587000, 168832) != 0 {
		v10 = (*uint8)(memmap.PtrOff(0x587000, 168872))
		for {
			if *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v10))), -int(unsafe.Sizeof(uint32(0))*6)))) == 0 && *(*uint32)(unsafe.Pointer(v10)) == 0 {
				nox_client_drawImageAt_47D2C0((*nox_video_bag_image_t)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v10))), -int(unsafe.Sizeof(uint32(0))*9))))))), int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v10))), -int(unsafe.Sizeof(uint32(0))*8))))), int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v10))), -int(unsafe.Sizeof(uint32(0))*7))))))
			}
			v11 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v10))), unsafe.Sizeof(uint32(0))*2))))
			v10 = (*uint8)(unsafe.Add(unsafe.Pointer(v10), 48))
			if v11 == 0 {
				break
			}
		}
	}
	return 1
}
