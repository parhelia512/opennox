package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"unsafe"
)

func nox_xxx_itemGetName_4E77E0_obj_util(a1 int32) *libc.WChar {
	var (
		v1  int32
		v2  *uint32
		v3  *uint32
		v4  *uint32
		v5  *libc.WChar
		v7  int32
		v8  *libc.WChar
		v9  int32
		v10 int32
		v11 *libc.WChar
		v12 int32
	)
	v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 8))))
	if uint32(v1)&0x13001000 != 0 {
		v2 = *(**uint32)(unsafe.Pointer(uintptr(a1 + 692)))
		if uint32(v1)&0x11001000 != 0 {
			v3 = nox_xxx_getProjectileClassById_413250(int32(*(*uint16)(unsafe.Pointer(uintptr(a1 + 4)))))
		} else {
			v3 = nox_xxx_equipClothFindDefByTT_413270(int32(*(*uint16)(unsafe.Pointer(uintptr(a1 + 4)))))
		}
		v4 = v3
		if v3 == nil {
			v12 = nox_xxx_getUnitName_4E39D0(a1)
			v5 = strMan.GetStringInFile("NoInfo", "C:\\NoxPost\\src\\Server\\Object\\objutil.c")
			nox_swprintf((*libc.WChar)(memmap.PtrOff(6112660, 1565660)), v5, v12)
			return (*libc.WChar)(memmap.PtrOff(6112660, 1565660))
		}
		nox_wcscpy((*libc.WChar)(memmap.PtrOff(6112660, 1565660)), (*libc.WChar)(memmap.PtrOff(6112660, 1567732)))
		if *v2 != 0 && *(*uint32)(unsafe.Pointer(uintptr(*v2 + 8))) != 0 {
			nox_wcscat((*libc.WChar)(memmap.PtrOff(6112660, 1565660)), *(**libc.WChar)(unsafe.Pointer(uintptr(*v2 + 8))))
			nox_wcscat((*libc.WChar)(memmap.PtrOff(6112660, 1565660)), libc.CWString(" "))
		}
		v7 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*1)))
		if v7 != 0 {
			v8 = *(**libc.WChar)(unsafe.Pointer(uintptr(v7 + 8)))
			if v8 != nil {
				nox_wcscat((*libc.WChar)(memmap.PtrOff(6112660, 1565660)), v8)
				nox_wcscat((*libc.WChar)(memmap.PtrOff(6112660, 1565660)), libc.CWString(" "))
			}
		}
		if *(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*2)) != 0 {
			nox_wcscat((*libc.WChar)(memmap.PtrOff(6112660, 1565660)), (*libc.WChar)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*2))))))
		}
		v9 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*2)))
		if v9 != 0 && *(*uint32)(unsafe.Pointer(uintptr(v9 + 8))) != 0 {
			nox_wcscat((*libc.WChar)(memmap.PtrOff(6112660, 1565660)), libc.CWString(" "))
			nox_wcscat((*libc.WChar)(memmap.PtrOff(6112660, 1565660)), *(**libc.WChar)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*2)) + 8))))
		}
		v10 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*3)))
		if v10 != 0 && *(*uint32)(unsafe.Pointer(uintptr(v10 + 12))) != 0 {
			nox_wcscat((*libc.WChar)(memmap.PtrOff(6112660, 1565660)), libc.CWString(" "))
			nox_wcscat((*libc.WChar)(memmap.PtrOff(6112660, 1565660)), *(**libc.WChar)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*3)) + 12))))
			return (*libc.WChar)(memmap.PtrOff(6112660, 1565660))
		}
	} else {
		v11 = strMan.GetStringInFile("NoDescription", "C:\\NoxPost\\src\\Server\\Object\\objutil.c")
		nox_wcscpy((*libc.WChar)(memmap.PtrOff(6112660, 1565660)), v11)
	}
	return (*libc.WChar)(memmap.PtrOff(6112660, 1565660))
}
