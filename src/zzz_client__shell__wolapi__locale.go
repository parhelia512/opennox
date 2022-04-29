package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"image"
	"math"
	"unsafe"
)

func sub_4B5770_wol_locale(a1 int32) int32 {
	var (
		i   int32
		v2  *libc.WChar
		v3  *libc.WChar
		v4  *uint32
		v5  *uint32
		v6  *uint32
		v7  *uint32
		v8  *byte
		v10 *byte
	)
	*memmap.PtrUint32(6112660, 1312488) = uint32(a1)
	dword_5d4594_1312480 = uint32(uintptr(unsafe.Pointer(newWindowFromFile("locale.wnd", unsafe.Pointer(cgoFuncAddr(sub_4B5AB0))))))
	(*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1312480)))).SetPos(image.Pt(nox_win_width/2-75, nox_win_height/2-77))
	sub_46B120((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1312480)))), nil)
	nox_xxx_wndShowModalMB((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1312480))))))
	sub_46C690((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1312480))))))
	guiFocus((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1312480))))))
	dword_5d4594_1312484 = uint32(uintptr(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1312480)))).ChildByID(1981))))
	sub_4B5990()
	for i = 0; i < *(*int32)(unsafe.Pointer(&dword_5d4594_1312472)); i++ {
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1312484))))).Func94(asWindowEvent(0x400D, int32(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1312476 + uint32(i*4))))), -1))
	}
	v2 = strMan.GetStringInFile("Other", "C:\\NoxPost\\src\\client\\shell\\WolApi\\locale.c")
	(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1312484))))).Func94(asWindowEvent(0x400D, int32(uintptr(unsafe.Pointer(v2))), -1))
	(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1312484))))).Func94(asWindowEvent(0x4013, 0, 0))
	v3 = strMan.GetStringInFile("Unknown", "C:\\NoxPost\\src\\client\\shell\\WolApi\\locale.c")
	(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1312484))))).Func94(asWindowEvent(0x400D, int32(uintptr(unsafe.Pointer(v3))), -1))
	(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1312484))))).Func94(asWindowEvent(0x4013, 0, 0))
	v4 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1312480)))).ChildByID(1982)))
	v5 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1312480)))).ChildByID(1983)))
	v6 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1312480)))).ChildByID(1984)))
	v7 = *(**uint32)(unsafe.Pointer(uintptr(dword_5d4594_1312484 + 32)))
	v10 = (*byte)(unsafe.Pointer(nox_xxx_gLoadImg("UISlider")))
	v8 = (*byte)(unsafe.Pointer(nox_xxx_gLoadImg("UISliderLit")))
	sub_4B5700(int32(uintptr(unsafe.Pointer(v4))), 0, 0, int32(uintptr(unsafe.Pointer(v10))), int32(uintptr(unsafe.Pointer(v8))), int32(uintptr(unsafe.Pointer(v8))))
	nox_xxx_wnd_46B280(int32(uintptr(unsafe.Pointer(v4))), *(*int32)(unsafe.Pointer(&dword_5d4594_1312484)))
	nox_xxx_wnd_46B280(int32(uintptr(unsafe.Pointer(v5))), *(*int32)(unsafe.Pointer(&dword_5d4594_1312484)))
	nox_xxx_wnd_46B280(int32(uintptr(unsafe.Pointer(v6))), *(*int32)(unsafe.Pointer(&dword_5d4594_1312484)))
	*(*uint32)(unsafe.Add(unsafe.Pointer(v7), unsafe.Sizeof(uint32(0))*9)) = uint32(uintptr(unsafe.Pointer(v4)))
	*(*uint32)(unsafe.Add(unsafe.Pointer(v7), unsafe.Sizeof(uint32(0))*7)) = uint32(uintptr(unsafe.Pointer(v5)))
	*(*uint32)(unsafe.Add(unsafe.Pointer(v7), unsafe.Sizeof(uint32(0))*8)) = uint32(uintptr(unsafe.Pointer(v6)))
	*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*100)) + 8))) = 16
	*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*100)) + 12))) = 10
	return 1
}
func sub_4B5A30_wol_locale(a1 *libc.WChar) *libc.WChar {
	var (
		v1 int32
		v2 *int16
		v3 **byte
		v4 *libc.WChar
		v6 [16]int16
	)
	v1 = 0
	libc.MemSet(unsafe.Pointer(&v6[0]), math.MaxUint8, int(unsafe.Sizeof([16]int16{})))
	v2 = &v6[0]
	if *(*int32)(unsafe.Pointer(&dword_5d4594_1312472)) <= 0 {
		return (*libc.WChar)(unsafe.Pointer(&v6[0]))
	}
	v3 = (**byte)(memmap.PtrOff(0x587000, 174360))
	for {
		v4 = strMan.GetStringInFile(*v3, "C:\\NoxPost\\src\\client\\shell\\WolApi\\locale.c")
		if nox_wcscmp(v4, (*libc.WChar)(unsafe.Pointer(v2))) < 0 && nox_wcscmp(v4, a1) > 0 {
			v2 = (*int16)(unsafe.Pointer(v4))
		}
		v1++
		v3 = (**byte)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof((*byte)(nil))*2))
		if v1 >= *(*int32)(unsafe.Pointer(&dword_5d4594_1312472)) {
			break
		}
	}
	return (*libc.WChar)(unsafe.Pointer(v2))
}
func sub_4B5B70_wol_locale(a1 *libc.WChar) int32 {
	var (
		v1     int32
		v2     **byte
		v3     *libc.WChar
		v4     *libc.WChar
		result int32
	)
	v1 = 0
	if *(*int32)(unsafe.Pointer(&dword_5d4594_1312472)) <= 0 {
	LABEL_5:
		v4 = strMan.GetStringInFile("Unknown", "C:\\NoxPost\\src\\client\\shell\\WolApi\\locale.c")
		result = bool2int(nox_wcscmp(a1, v4) != 0)
	} else {
		v2 = (**byte)(memmap.PtrOff(0x587000, 174360))
		for {
			v3 = strMan.GetStringInFile(*v2, "C:\\NoxPost\\src\\client\\shell\\WolApi\\locale.c")
			if nox_wcscmp(a1, v3) == 0 {
				break
			}
			v1++
			v2 = (**byte)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof((*byte)(nil))*2))
			if v1 >= *(*int32)(unsafe.Pointer(&dword_5d4594_1312472)) {
				goto LABEL_5
			}
		}
		result = int32(*memmap.PtrUint32(0x587000, uint32(v1*8)+0x2A91C))
	}
	return result
}
