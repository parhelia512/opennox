package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"unsafe"
)

func nox_xxx_loadGuides_427070() int32 {
	var (
		v0  **byte
		i   *uint8
		v2  int32
		v3  *libc.WChar
		v4  int32
		v5  int32
		v6  *byte
		v7  *byte
		v8  int32
		v9  int32
		v10 *byte
		v11 int32
		v13 [256]byte
	)
	v0 = (**byte)(memmap.PtrOff(0x587000, 70504))
	for i = (*uint8)(memmap.PtrOff(6112660, 740108)); ; i = (*uint8)(unsafe.Add(unsafe.Pointer(i), 28)) {
		v2 = sub_44D330(*v0)
		if v2 == 0 {
			break
		}
		nox_sprintf(&v13[0], CString("creature:%s"), *v0)
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(i))), -int(unsafe.Sizeof(uint32(0))*1)))) = uint32(uintptr(unsafe.Pointer(strMan.GetStringInFile(&v13[0], "C:\\NoxPost\\src\\common\\Magic\\ComGuide.c"))))
		if libc.StrCmp(*(**byte)(unsafe.Pointer(uintptr(v2))), CString("Bomber")) == 0 {
			*(*uint32)(unsafe.Pointer(i)) = 0
		} else {
			*(*uint32)(unsafe.Pointer(i)) = uint32(nox_xxx_getTTByNameSpriteMB_44CFC0(*v0))
		}
		nox_sprintf(&v13[0], CString("creature_desc:%s"), *v0)
		v3 = strMan.GetStringInFile(&v13[0], "C:\\NoxPost\\src\\common\\Magic\\ComGuide.c")
		v4 = int32(*memmap.PtrUint32(0x587000, 71248))
		v5 = int32(*memmap.PtrUint32(0x587000, 71252))
		v6 = *v0
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(i))), unsafe.Sizeof(uint32(0))*1))) = uint32(uintptr(unsafe.Pointer(v3)))
		*(*uint32)(unsafe.Pointer(&v13[0])) = uint32(v4)
		libc.StrCpy(&v13[8], CString("Cage"))
		*(*uint32)(unsafe.Pointer(&v13[4])) = uint32(v5)
		libc.StrCat(&v13[0], v6)
		v7 = (*byte)(unsafe.Pointer(nox_xxx_gLoadImg(&v13[0])))
		v8 = int32(*memmap.PtrUint32(0x587000, 71264))
		v9 = int32(*memmap.PtrUint32(0x587000, 71268))
		v10 = *v0
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(i))), unsafe.Sizeof(uint32(0))*2))) = uint32(uintptr(unsafe.Pointer(v7)))
		*(*uint32)(unsafe.Pointer(&v13[0])) = uint32(v8)
		libc.StrCpy(&v13[8], CString("k"))
		*(*uint32)(unsafe.Pointer(&v13[4])) = uint32(v9)
		libc.StrCat(&v13[0], v10)
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(i))), unsafe.Sizeof(uint32(0))*3))) = uint32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg(&v13[0]))))
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(i))), unsafe.Sizeof(uint32(0))*4))) = 0
		v11 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 36))))
		if v11&1 != 0 {
			*(*uint8)(unsafe.Add(unsafe.Pointer(i), 20)) = 1
		} else if v11&2 != 0 {
			*(*uint8)(unsafe.Add(unsafe.Pointer(i), 20)) = 2
		} else {
			*(*uint8)(unsafe.Add(unsafe.Pointer(i), 20)) = 4
		}
		v0 = (**byte)(unsafe.Add(unsafe.Pointer(v0), unsafe.Sizeof((*byte)(nil))*1))
		if int32(uintptr(unsafe.Pointer(v0))) >= int32(uintptr(memmap.PtrOff(0x587000, 70664))) {
			return 1
		}
	}
	return 0
}
