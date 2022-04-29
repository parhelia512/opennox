package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"unsafe"
)

func sub_415B60(a1 int32) *libc.WChar {
	var (
		v1 int32
		v2 int32
		i  *uint8
		v4 int32
	)
	v1 = 0
	if *memmap.PtrUint32(0x587000, 34848) == 0 {
		return strMan.GetStringInFile("result:ERROR", "C:\\NoxPost\\src\\common\\Object\\ArmrLook.c")
	}
	*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v2))), unsafe.Sizeof(uint16(0))*1)) = 0
	for i = (*uint8)(memmap.PtrOff(0x587000, 34848)); ; i = (*uint8)(unsafe.Add(unsafe.Pointer(i), 24)) {
		*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v2))), unsafe.Sizeof(uint16(0))*0)) = *(*uint16)(unsafe.Pointer(uintptr(a1 + 4)))
		if uint32(v2) == *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(i))), unsafe.Sizeof(uint32(0))*2))) {
			break
		}
		v4 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(i))), unsafe.Sizeof(uint32(0))*6))))
		v1++
		if v4 == 0 {
			return strMan.GetStringInFile("result:ERROR", "C:\\NoxPost\\src\\common\\Object\\ArmrLook.c")
		}
	}
	return strMan.GetStringInFile(*(**byte)(memmap.PtrOff(0x587000, uint32(v1*24+0x8824))), "C:\\NoxPost\\src\\common\\Object\\ArmrLook.c")
}
func nox_xxx_loadLook_415D50() **byte {
	var (
		result **byte
		v1     *uint8
		v2     int32
	)
	result = *(***byte)(memmap.PtrOff(6112660, 371256))
	if *memmap.PtrUint32(6112660, 371256) == 0 {
		result = *(***byte)(memmap.PtrOff(0x587000, 35500))
		if *memmap.PtrUint32(0x587000, 35500) != 0 {
			result = (**byte)(memmap.PtrOff(0x587000, 35500))
			v1 = (*uint8)(memmap.PtrOff(0x587000, 35500))
			for {
				*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v1))), -int(unsafe.Sizeof(uint32(0))*1)))) = uint32(uintptr(unsafe.Pointer(strMan.GetStringInFile(*result, "C:\\NoxPost\\src\\common\\Object\\ArmrLook.c"))))
				v2 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v1))), unsafe.Sizeof(uint32(0))*3))))
				v1 = (*uint8)(unsafe.Add(unsafe.Pointer(v1), 12))
				result = (**byte)(unsafe.Pointer(v1))
				if v2 == 0 {
					break
				}
			}
		}
		*memmap.PtrUint32(6112660, 371256) = 1
	}
	return result
}
