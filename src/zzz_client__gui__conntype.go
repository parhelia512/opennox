package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"image"
	"unsafe"
)

func sub_49C820() int32 {
	var (
		v0 *uint32
		v1 **byte
		v2 *libc.WChar
	)
	dword_5d4594_1305684 = uint32(uintptr(unsafe.Pointer(newWindowFromFile("conntype.wnd", unsafe.Pointer(cgoFuncAddr(sub_49CA60))))))
	sub_46B120((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1305684)))), nil)
	nox_xxx_wndShowModalMB((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1305684))))))
	sub_46C690((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1305684))))))
	guiFocus((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1305684))))))
	sub_49C910()
	(*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1305684)))).SetPos(image.Pt(nox_win_width/2-*(*int32)(unsafe.Pointer(uintptr(dword_5d4594_1305684 + 8)))/2, nox_win_height/2-*(*int32)(unsafe.Pointer(uintptr(dword_5d4594_1305684 + 12)))/2))
	nox_xxx_guiServerOptsLoad_457500()
	sub_459D80(1)
	v0 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1305684)))).ChildByID(0x2870)))
	v1 = (**byte)(memmap.PtrOff(0x587000, 164928))
	for {
		v2 = strMan.GetStringInFile(*v1, "C:\\NoxPost\\src\\client\\Gui\\conntype.c")
		(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v0)))))).Func94(asWindowEvent(0x400D, int32(uintptr(unsafe.Pointer(v2))), -1))
		v1 = (**byte)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof((*byte)(nil))*1))
		if int32(uintptr(unsafe.Pointer(v1))) >= int32(uintptr(memmap.PtrOff(0x587000, 164944))) {
			break
		}
	}
	return (*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v0)))))).Func94(asWindowEvent(0x4013, 0, 0))
}
func sub_49C910() *uint32 {
	var (
		v0     *uint16
		v1     *uint16
		v2     **byte
		v3     int32
		v4     int32
		v5     *uint16
		v6     int32
		v7     int32
		result *uint32
		v9     int32
		v10    int32
		v11    int32
		v12    int32
	)
	v0 = (*uint16)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1305684)))).ChildByID(0x2870)))
	v1 = v0
	v2 = (**byte)(memmap.PtrOff(0x587000, 164928))
	v3 = (nox_xxx_guiFontHeightMB_43F320(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v0))), unsafe.Sizeof(uint32(0))*59)))))) + 1) * 5
	*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v1))), unsafe.Sizeof(uint32(0))*7))) = *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v1))), unsafe.Sizeof(uint32(0))*5))) + uint32(v3) + 2
	*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v1))), unsafe.Sizeof(uint32(0))*3))) = uint32(v3 + 2)
	v4 = 0
	for {
		v5 = (*uint16)(unsafe.Pointer(strMan.GetStringInFile(*v2, "C:\\NoxPost\\src\\client\\Gui\\conntype.c")))
		nox_xxx_drawGetStringSize_43F840(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v1))), unsafe.Sizeof(uint32(0))*59))))), (*libc.WChar)(unsafe.Pointer(v5)), &v11, nil, 0)
		if v11 > v4 {
			v4 = v11
		}
		v2 = (**byte)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof((*byte)(nil))*1))
		if int32(uintptr(unsafe.Pointer(v2))) >= int32(uintptr(memmap.PtrOff(0x587000, 164944))) {
			break
		}
	}
	v6 = v4 + 7
	nox_xxx_drawGetStringSize_43F840(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v1))), unsafe.Sizeof(uint32(0))*59))))), (*libc.WChar)(unsafe.Pointer((*uint16)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint16(0))*54)))), &v12, nil, 0)
	if v6 <= v12 {
		v6 = v12
	}
	v7 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v1))), unsafe.Sizeof(uint32(0))*4))))
	*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v1))), unsafe.Sizeof(uint32(0))*2))) = uint32(v6)
	*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v1))), unsafe.Sizeof(uint32(0))*6))) = uint32(v7 + v6)
	*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1305684 + 16))) = uint32(v7 - 40)
	*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1305684 + 20))) = *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v1))), unsafe.Sizeof(uint32(0))*5))) - 40
	*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1305684 + 24))) = *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v1))), unsafe.Sizeof(uint32(0))*6))) + 40
	*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1305684 + 28))) = *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v1))), unsafe.Sizeof(uint32(0))*7))) + 40
	*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1305684 + 8))) = *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v1))), unsafe.Sizeof(uint32(0))*2))) + 80
	*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1305684 + 12))) = *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v1))), unsafe.Sizeof(uint32(0))*3))) + 80
	result = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1305684)))).ChildByID(0x2871)))
	*(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*5)) = *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v1))), unsafe.Sizeof(uint32(0))*7))) + 10
	v9 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*2)))
	v10 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*5)))
	*(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*4)) = *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v1))), unsafe.Sizeof(uint32(0))*6))) - uint32(v9)
	*(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*7)) = uint32(v10) + *(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*3))
	*(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*6)) = uint32(v9) + *(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*4))
	return result
}
