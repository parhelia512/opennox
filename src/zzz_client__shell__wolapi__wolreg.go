package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"unsafe"
)

func sub_468890_wol_reg(a1 *byte, a2 *byte, a3 *byte, a4 *uint8, a5 *uint8, a6 *libc.WChar) int32 {
	var (
		v6     *uint32
		v7     int32
		v8     *uint32
		v9     int32
		v10    *uint32
		v11    int32
		v12    *uint32
		v13    int32
		v14    int8
		v15    *libc.WChar
		result int32
		v17    *libc.WChar
		v18    *libc.WChar
		v19    *libc.WChar
		v20    *libc.WChar
		v21    *libc.WChar
		v22    *libc.WChar
		v23    *libc.WChar
		v24    *libc.WChar
		v25    *libc.WChar
		v26    [80]byte
	)
	*a4 = uint8(((*nox_window)(dword_5d4594_1064816).ChildByID(1766).draw_data.field_0 >> 2) & 1)
	*a5 = uint8(((*nox_window)(dword_5d4594_1064816).ChildByID(1767).draw_data.field_0 >> 2) & 1)
	v6 = (*uint32)(unsafe.Pointer((*nox_window)(dword_5d4594_1064816).ChildByID(1762)))
	v7 = (*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v6)))))).Func94(asWindowEvent(0x401D, 0, 0))
	nox_sprintf(a1, CString("%S"), v7)
	v8 = (*uint32)(unsafe.Pointer((*nox_window)(dword_5d4594_1064816).ChildByID(1763)))
	v9 = (*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v8)))))).Func94(asWindowEvent(0x401D, 0, 0))
	nox_sprintf(a2, CString("%S"), v9)
	v10 = (*uint32)(unsafe.Pointer((*nox_window)(dword_5d4594_1064816).ChildByID(1764)))
	v11 = (*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v10)))))).Func94(asWindowEvent(0x401D, 0, 0))
	nox_sprintf(&v26[0], CString("%S"), v11)
	v12 = (*uint32)(unsafe.Pointer((*nox_window)(dword_5d4594_1064816).ChildByID(1765)))
	v13 = (*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v12)))))).Func94(asWindowEvent(0x401D, 0, 0))
	nox_sprintf(a3, CString("%S"), v13)
	v14 = int8(*a1)
	if *a1 != 0 {
		if int32(uint8(v14)) < 49 || int32(uint8(v14)) > 57 {
			if libc.StrLen(a2) == 8 {
				if libc.StrCmp(a2, &v26[0]) == 0 {
					if dword_5d4594_1064300 == 0 || *a3 != 0 {
						result = 1
					} else {
						v25 = strMan.GetStringInFile("TryAgain", "C:\\NoxPost\\src\\client\\shell\\WolApi\\wolreg.c")
						v20 = strMan.GetStringInFile("ErrNeedParentsEmail", "C:\\NoxPost\\src\\client\\shell\\WolApi\\wolreg.c")
						nox_swprintf(a6, libc.CWString("%s %s"), v20, v25)
						result = 0
					}
				} else {
					v24 = strMan.GetStringInFile("TryAgain", "C:\\NoxPost\\src\\client\\shell\\WolApi\\wolreg.c")
					v19 = strMan.GetStringInFile("ErrPassDontMatch", "C:\\NoxPost\\src\\client\\shell\\WolApi\\wolreg.c")
					nox_swprintf(a6, libc.CWString("%s %s"), v19, v24)
					result = 0
				}
			} else {
				v23 = strMan.GetStringInFile("TryAgain", "C:\\NoxPost\\src\\client\\shell\\WolApi\\wolreg.c")
				v18 = strMan.GetStringInFile("ErrPasswordLength", "C:\\NoxPost\\src\\client\\shell\\WolApi\\wolreg.c")
				nox_swprintf(a6, libc.CWString("%s %s"), v18, v23)
				result = 0
			}
		} else {
			v22 = strMan.GetStringInFile("TryAgain", "C:\\NoxPost\\src\\client\\shell\\WolApi\\wolreg.c")
			v17 = strMan.GetStringInFile("ErrNickFirstLetter", "C:\\NoxPost\\src\\client\\shell\\WolApi\\wolreg.c")
			nox_swprintf(a6, libc.CWString("%s %s"), v17, v22)
			result = 0
		}
	} else {
		v21 = strMan.GetStringInFile("TryAgain", "C:\\NoxPost\\src\\client\\shell\\WolApi\\wolreg.c")
		v15 = strMan.GetStringInFile("ErrNoNick", "C:\\NoxPost\\src\\client\\shell\\WolApi\\wolreg.c")
		nox_swprintf(a6, libc.CWString("%s %s"), v15, v21)
		result = 0
	}
	return result
}
func sub_468BB0_wol_reg(a1 *int32, a2 *int32, a3 *int32, a4 *libc.WChar) int32 {
	var (
		v4     *uint32
		v5     *libc.WChar
		v6     *uint32
		v7     *libc.WChar
		v8     *uint32
		v9     *libc.WChar
		v10    *uint32
		v11    int32
		v12    *libc.WChar
		result int32
		v14    *libc.WChar
		v15    *libc.WChar
		v16    *libc.WChar
		v17    *libc.WChar
		v18    *libc.WChar
	)
	*a4 = 0
	v4 = (*uint32)(unsafe.Pointer((*nox_window)(dword_5d4594_1064816).ChildByID(1758)))
	v5 = (*libc.WChar)(unsafe.Pointer(uintptr((*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v4)))))).Func94(asWindowEvent(0x401D, 0, 0)))))
	v6 = (*uint32)(unsafe.Pointer((*nox_window)(dword_5d4594_1064816).ChildByID(1759)))
	v7 = (*libc.WChar)(unsafe.Pointer(uintptr((*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v6)))))).Func94(asWindowEvent(0x401D, 0, 0)))))
	v8 = (*uint32)(unsafe.Pointer((*nox_window)(dword_5d4594_1064816).ChildByID(1760)))
	v9 = (*libc.WChar)(unsafe.Pointer(uintptr((*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v8)))))).Func94(asWindowEvent(0x401D, 0, 0)))))
	v10 = (*uint32)(unsafe.Pointer((*nox_window)(dword_5d4594_1064816).ChildByID(1761)))
	v11 = (*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v10)))))).Func94(asWindowEvent(0x401D, 0, 0))
	nox_sprintf((*byte)(memmap.PtrOff(6112660, 1064196)), CString("%S"), v11)
	if v5 != nil && v7 != nil && nox_wcslen(v9) == 4 {
		if int32(*memmap.PtrUint8(6112660, 1064196)) != 0 {
			*a1 = nox_wcstol(v5, nil, 10)
			*a2 = nox_wcstol(v7, nil, 10)
			*a3 = nox_wcstol(v9, nil, 10)
			if *a1 < 1 || *a1 > 12 || *a2 < 1 || *a2 > 31 {
				v17 = strMan.GetStringInFile("TryAgain", "C:\\NoxPost\\src\\client\\shell\\WolApi\\wolreg.c")
				v14 = strMan.GetStringInFile("ErrAge", "C:\\NoxPost\\src\\client\\shell\\WolApi\\wolreg.c")
				nox_swprintf(a4, libc.CWString("%s %s"), v14, v17)
				result = 0
			} else {
				result = 1
			}
		} else {
			v16 = strMan.GetStringInFile("TryAgain", "C:\\NoxPost\\src\\client\\shell\\WolApi\\wolreg.c")
			v12 = strMan.GetStringInFile("ErrEmail", "C:\\NoxPost\\src\\client\\shell\\WolApi\\wolreg.c")
			nox_swprintf(a4, libc.CWString("%s %s"), v12, v16)
			result = 0
		}
	} else {
		v18 = strMan.GetStringInFile("TryAgain", "C:\\NoxPost\\src\\client\\shell\\WolApi\\wolreg.c")
		v15 = strMan.GetStringInFile("ErrAge", "C:\\NoxPost\\src\\client\\shell\\WolApi\\wolreg.c")
		nox_swprintf(a4, libc.CWString("%s %s"), v15, v18)
		result = 0
	}
	return result
}
func sub_468F30_wol_reg() int32 {
	var v0 *libc.WChar
	v0 = strMan.GetStringInFile("Wolreg.c:ErrNoConsent", "C:\\NoxPost\\src\\client\\shell\\WolApi\\wolreg.c")
	sub_468840(v0)
	dword_5d4594_1064296 = 4
	(*nox_window)(dword_5d4594_1064820).Hide()
	return nox_xxx_wnd_46ABB0((*nox_window)(dword_5d4594_1064816), 1)
}
