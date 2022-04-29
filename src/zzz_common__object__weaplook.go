package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"unsafe"
)

func nox_xxx_loadModifyers_4158C0() **byte {
	var (
		result **byte
		v1     *uint8
		v2     int32
	)
	result = *(***byte)(memmap.PtrOff(6112660, 371248))
	if *memmap.PtrUint32(6112660, 371248) == 0 {
		result = *(***byte)(memmap.PtrOff(0x587000, 33396))
		if *memmap.PtrUint32(0x587000, 33396) != 0 {
			result = (**byte)(memmap.PtrOff(0x587000, 33396))
			v1 = (*uint8)(memmap.PtrOff(0x587000, 33396))
			for {
				*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v1))), -int(unsafe.Sizeof(uint32(0))*1)))) = uint32(uintptr(unsafe.Pointer(strMan.GetStringInFile(*result, "C:\\NoxPost\\src\\common\\Object\\WeapLook.c"))))
				v2 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v1))), unsafe.Sizeof(uint32(0))*3))))
				v1 = (*uint8)(unsafe.Add(unsafe.Pointer(v1), 12))
				result = (**byte)(unsafe.Pointer(v1))
				if v2 == 0 {
					break
				}
			}
		}
		*memmap.PtrUint32(6112660, 371248) = 1
	}
	return result
}
