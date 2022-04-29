package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"math"
	"unsafe"
)

func sub_454640() int32 {
	var (
		v0     *uint32
		v1     *uint32
		v2     int32
		v3     *uint16
		v4     *uint16
		v5     *uint16
		v6     int32
		v7     int32
		result int32
		v9     int32
		v10    int32
	)
	v0 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045516)))).ChildByID(0x278B)))
	v1 = v0
	v2 = nox_xxx_guiFontHeightMB_43F320(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v0), unsafe.Sizeof(uint32(0))*59))))) + 1
	*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*7)) = *(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*5)) + uint32(v2*4) + 2
	*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*3)) = uint32(v2*4 + 2)
	v3 = (*uint16)(unsafe.Pointer(strMan.GetStringInFile("WARRIOR", "C:\\NoxPost\\src\\client\\Gui\\ServOpts\\access.c")))
	nox_xxx_drawGetStringSize_43F840(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*59)))), (*libc.WChar)(unsafe.Pointer(v3)), &v9, nil, 0)
	v4 = (*uint16)(unsafe.Pointer(strMan.GetStringInFile("WIZARD", "C:\\NoxPost\\src\\client\\Gui\\ServOpts\\access.c")))
	nox_xxx_drawGetStringSize_43F840(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*59)))), (*libc.WChar)(unsafe.Pointer(v4)), &v10, nil, 0)
	if v10 > v9 {
		v9 = v10
	}
	v5 = (*uint16)(unsafe.Pointer(strMan.GetStringInFile("CONJURER", "C:\\NoxPost\\src\\client\\Gui\\ServOpts\\access.c")))
	nox_xxx_drawGetStringSize_43F840(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*59)))), (*libc.WChar)(unsafe.Pointer(v5)), &v10, nil, 0)
	v6 = v9
	if v10 > v9 {
		v6 = v10
	}
	v7 = v6 + 7
	v9 = v7
	*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*2)) = uint32(v7)
	result = int32(uint32(v9) + *(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*4)))
	*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*6)) = uint32(result)
	return result
}
func sub_454740() *int32 {
	var (
		v0          *byte
		v1          *uint32
		v2          *libc.WChar
		v3          *uint32
		v4          *uint32
		v5          *uint32
		v6          *libc.WChar
		v7          *libc.WChar
		v8          *libc.WChar
		v9          int8
		v10         int32
		v11         int32
		i           *byte
		WideCharStr [18]libc.WChar
	)
	v0 = (*byte)(sub_416640())
	v1 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045516)))).ChildByID(0x2798)))
	v2 = nox_xxx_sysopGetPass_40A630()
	(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v1)))))).Func94(asWindowEvent(0x401E, int32(uintptr(unsafe.Pointer(v2))), 0))
	if int32(*(*int16)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v0), 105))))) != -1 {
		nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045584))))), 1)
		*(*uint32)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, 1045568) + 36))) |= 4
		compat_itow(int32(*(*uint16)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v0), 105))))), &WideCharStr[0], 10)
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045584))))).Func94(asWindowEvent(0x401E, int32(uintptr(unsafe.Pointer(&WideCharStr[0]))), 0))
	}
	if int32(*(*int16)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v0), 107))))) != -1 {
		nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045588))))), 1)
		*(*uint32)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, 1045572) + 36))) |= 4
		compat_itow(int32(*(*uint16)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v0), 107))))), &WideCharStr[0], 10)
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045588))))).Func94(asWindowEvent(0x401E, int32(uintptr(unsafe.Pointer(&WideCharStr[0]))), 0))
	}
	v3 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045516)))).ChildByID(0x278C)))
	if int32(*(*byte)(unsafe.Add(unsafe.Pointer(v0), 102))) < 0 {
		*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*9)) |= 4
	}
	if *(*byte)(unsafe.Add(unsafe.Pointer(v0), 100))&32 != 0 {
		nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045556))))), 1)
		*(*uint32)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, 1045524) + 36))) |= 4
	}
	(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045556))))).Func94(asWindowEvent(0x401E, int32(uintptr(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v0), 78))))), 0))
	if sub_4D6F30() != 0 {
		nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045520))))), 0)
	} else {
		nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045520))))), 1)
		if *(*byte)(unsafe.Add(unsafe.Pointer(v0), 100))&16 != 0 {
			*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1045520 + 36))) = 4
		} else {
			v4 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045516)))).ChildByID(0x27DE)))
			nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v4)))))), 0)
		}
	}
	v5 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045516)))).ChildByID(0x27DF)))
	*(*uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(uint32(0))*9)) |= 4
	dword_5d4594_1045596 = dword_5d4594_1045528
	v6 = strMan.GetStringInFile("WARRIOR", "C:\\NoxPost\\src\\client\\Gui\\ServOpts\\access.c")
	(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045552))))).Func94(asWindowEvent(0x400D, int32(uintptr(unsafe.Pointer(v6))), -1))
	v7 = strMan.GetStringInFile("WIZARD", "C:\\NoxPost\\src\\client\\Gui\\ServOpts\\access.c")
	(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045552))))).Func94(asWindowEvent(0x400D, int32(uintptr(unsafe.Pointer(v7))), -1))
	v8 = strMan.GetStringInFile("CONJURER", "C:\\NoxPost\\src\\client\\Gui\\ServOpts\\access.c")
	(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045552))))).Func94(asWindowEvent(0x400D, int32(uintptr(unsafe.Pointer(v8))), -1))
	if *(*byte)(unsafe.Add(unsafe.Pointer(v0), 100))&16 != 0 {
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045532))))).Show()
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045528))))).Hide()
	}
	v9 = int8(*(*byte)(unsafe.Add(unsafe.Pointer(v0), 100)))
	if int32(v9) != 0 {
		v10 = 0
		v11 = int32(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1045552 + 32))))
		if int32(v9)&1 != 0 {
			**(**uint32)(unsafe.Pointer(uintptr(v11 + 48))) = 0
			v10 = 1
			*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v11 + 48))) + 4))) = math.MaxUint32
		}
		if *(*byte)(unsafe.Add(unsafe.Pointer(v0), 100))&2 != 0 {
			*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v11 + 48))) + uint32(func() int32 {
				p := &v10
				*p++
				return *p
			}()*4) - 4))) = 1
			*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v11 + 48))) + uint32(v10*4)))) = math.MaxUint32
		}
		if *(*byte)(unsafe.Add(unsafe.Pointer(v0), 100))&4 != 0 {
			*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v11 + 48))) + uint32(v10*4)))) = 2
			*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v11 + 48))) + uint32(v10*4) + 4))) = math.MaxUint32
		}
	}
	compat_itow(int32(uint8(*(*byte)(unsafe.Add(unsafe.Pointer(v0), 104)))), &WideCharStr[0], 10)
	(*nox_window)(unsafe.Pointer(uintptr(*memmap.PtrInt32(6112660, 1045592)))).Func94(asWindowEvent(0x401E, int32(uintptr(unsafe.Pointer(&WideCharStr[0]))), 0))
	for i = (*byte)(unsafe.Pointer(nox_common_playerInfoGetFirst_416EA0())); i != nil; i = (*byte)(unsafe.Pointer(nox_common_playerInfoGetNext_416EE0((*nox_playerInfo)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(i))))))))) {
		if *(*byte)(unsafe.Add(unsafe.Pointer(i), 2064)) != 31 || !nox_common_getEngineFlag(nox_engine_flag(NOX_ENGINE_FLAG_DISABLE_GRAPHICS_RENDERING)) {
			sub_455920(int32(uintptr(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(i), 4704))))))
		}
	}
	return sub_455800()
}
