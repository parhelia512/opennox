package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	"unsafe"
)

func nox_xxx_spriteNPCInfo_49A4B0(a1 *uint32, a2 int32, a3 int32) int32 {
	var (
		result int32
		v4     int32
	)
	result = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*108)))
	v4 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*69)) - 1)
	*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*108)) = 0
	switch v4 {
	case 0:
		fallthrough
	case 1:
		fallthrough
	case 2:
		fallthrough
	case 3:
		if (uint32(a2) & 0xFFFFFFFC) == 0 {
			if result == 0 || *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*77)) == 0 {
				result = randomIntMinMax(23, 24)
			}
			goto LABEL_12
		}
		result = sub_4FA280(int32(uint32(a2) & 0xFFFFFFFC))
	case 4:
		fallthrough
	case 5:
		if a2&1024 != 0 {
			if result != 0 && *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*77)) != 0 {
			LABEL_12:
				*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*108)) = uint32(result)
			} else {
				result = randomIntMinMax(47, 49)
				*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*108)) = uint32(result)
			}
		} else if uint32(a2)&0x7FF8000 != 0 {
			result = 30
		} else {
			if (uint32(a3) & 0x3000000) == 0 {
				goto LABEL_22
			}
			result = 40
		}
	case 6:
		result = 21
	case 7:
	LABEL_22:
		if (a2 & 1024) != 0 {
			result = 38
		} else {
			result = 0
		}
	case 8:
		result = 1
	case 9:
		result = 2
	case 11:
		result = 4
	case 12:
		result = 6
	case 13:
		result = 50
	default:
		result = 0
	}
	return result
}
