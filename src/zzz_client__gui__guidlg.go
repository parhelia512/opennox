package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"image"
	"unsafe"
)

func sub_479D30(a1 *libc.WChar, a2 int32, a3 *byte, a4 *byte, a5 int8) int32 {
	var (
		v5     *uint32
		v6     *uint32
		v7     *uint32
		v8     *uint32
		v9     *uint32
		v10    *uint32
		v11    *libc.WChar
		v12    *uint32
		v13    *uint32
		v14    *uint32
		v15    *libc.WChar
		result int32
		v17    *byte
		v18    int32
		v19    int32
	)
	v5 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1123524)))).ChildByID(3901)))
	v6 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1123524)))).ChildByID(3910)))
	sub_445C20()
	(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v5)))))).Func94(asWindowEvent(0x400F, 0, 0))
	nox_window_get_size((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1123524))))), &v19, &v18)
	(*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1123524)))).SetPos(image.Pt(nox_win_width-v19, nox_win_height-v18))
	sub_47A020(a3)
	nox_wcscpy((*libc.WChar)(memmap.PtrOff(6112660, 1107056)), a1)
	v17 = (*byte)(unsafe.Pointer(nox_xxx_gLoadImg(a4)))
	v7 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1123524)))).ChildByID(3905)))
	nox_xxx_wndSetIcon_46AE60(int32(uintptr(unsafe.Pointer(v7))), int32(uintptr(unsafe.Pointer(v17))))
	(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v6)))))).Func94(asWindowEvent(0x4001, int32(uintptr(memmap.PtrOff(6112660, 1107056))), 0))
	v8 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1123524)))).ChildByID(3908)))
	if int32(a5)&1 != 0 {
		nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v8)))))), 1)
		(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v8)))))).Show()
		*(*uint32)(unsafe.Add(unsafe.Pointer(v8), unsafe.Sizeof(uint32(0))*5)) = 5
		v9 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1123524)))).ChildByID(3909)))
		nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v9)))))), 1)
		(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v9)))))).Show()
		*(*uint32)(unsafe.Add(unsafe.Pointer(v9), unsafe.Sizeof(uint32(0))*5)) = 35
		v10 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1123524)))).ChildByID(3906)))
		v11 = strMan.GetStringInFile("Dialog.wnd:Done", "C:\\NoxPost\\src\\client\\Gui\\GUIDlg.c")
	} else {
		nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v8)))))), 0)
		(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v8)))))).Hide()
		v12 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1123524)))).ChildByID(3909)))
		nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v12)))))), 0)
		(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v12)))))).Hide()
		v10 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1123524)))).ChildByID(3906)))
		v11 = strMan.GetStringInFile("Dialog.wnd:Done", "C:\\NoxPost\\src\\client\\Gui\\GUIDlg.c")
	}
	sub_46AEE0(int32(uintptr(unsafe.Pointer(v10))), int32(uintptr(unsafe.Pointer(v11))))
	nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v10)))))), 1)
	(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v10)))))).Show()
	*(*uint32)(unsafe.Add(unsafe.Pointer(v10), unsafe.Sizeof(uint32(0))*5)) = 95
	v13 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1123524)))).ChildByID(3907)))
	nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v13)))))), 1)
	(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v13)))))).Show()
	*(*uint32)(unsafe.Add(unsafe.Pointer(v13), unsafe.Sizeof(uint32(0))*5)) = 65
	if int32(a5)&2 != 0 {
		v14 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1123524)))).ChildByID(3906)))
		v15 = strMan.GetStringInFile("Dialog.wnd:Next", "C:\\NoxPost\\src\\client\\Gui\\GUIDlg.c")
		sub_46AEE0(int32(uintptr(unsafe.Pointer(v14))), int32(uintptr(unsafe.Pointer(v15))))
	}
	sub_467C10()
	nox_xxx_bookHideMB_45ACA0(0)
	*memmap.PtrUint32(0x587000, 153436) = nox_client_renderGUI_80828
	nox_client_renderGUI_80828 = 0
	nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1123524))))), 1)
	nox_xxx_wndShowModalMB((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1123524))))))
	*memmap.PtrUint32(6112660, 1123528) = uint32(a2)
	result = nox_xxx_playDialogFile_44D900(*memmap.PtrInt32(6112660, 1115312), 100)
	dword_5d4594_1123520 = 1
	return result
}
func sub_47A020(a1 *byte) *libc.WChar {
	var (
		v1     *uint32
		v2     int32
		v3     *libc.WChar
		v4     *libc.WChar
		result *libc.WChar
		v6     uint32
		v7     int32
		v8     *libc.WChar
		v9     int32
		v10    *libc.WChar
		v11    int32
		v12    *int16
		v13    [4]libc.WChar
	)
	*(*uint32)(unsafe.Pointer(&v13[0])) = *memmap.PtrUint32(0x587000, 153644)
	v13[2] = libc.WChar(*memmap.PtrUint16(0x587000, 153648))
	v1 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1123524)))).ChildByID(3901)))
	v2 = int32(uintptr(unsafe.Pointer(v1)))
	v3 = nil
	v12 = (*int16)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*8)))))
	v11 = 0
	v9 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*2)) - 10)
	v4 = nox_strman_loadString_40F1D0(a1, (**byte)(memmap.PtrOff(6112660, 1115312)), CString("C:\\NoxPost\\src\\client\\Gui\\GUIDlg.c"), 338)
	nox_wcscpy((*libc.WChar)(memmap.PtrOff(6112660, 1107120)), v4)
	result = nox_wcstok((*libc.WChar)(memmap.PtrOff(6112660, 1107120)), &v13[0])
	v10 = result
	if result != nil {
		for {
			if v3 == nil {
				v3 = result
			}
			v6 = nox_wcslen(v3)
			nox_xxx_drawGetStringSize_43F840(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 236))))), v3, (*int32)(unsafe.Pointer(&a1)), nil, 0)
			v7 = int32(v6)
			if int32(uintptr(unsafe.Pointer(a1))) > v9 {
				for {
					v6 = uint32(v7)
					if v7 == 0 {
						break
					}
					v8 = (*libc.WChar)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(libc.WChar(0))*uintptr(v7)))
					for *v8 != 32 {
						v6--
						v8 = (*libc.WChar)(unsafe.Add(unsafe.Pointer(v8), -int(unsafe.Sizeof(libc.WChar(0))*1)))
						if v6 == 0 {
							goto LABEL_15
						}
					}
					if v6 == 0 {
						break
					}
					v7 = int32(v6 - 1)
					nox_wcsncpy((*libc.WChar)(memmap.PtrOff(6112660, 1115324)), v3, v6)
					*memmap.PtrUint16(6112660, v6*2+0x1104BC) = 0
					nox_xxx_drawGetStringSize_43F840(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 236))))), (*libc.WChar)(unsafe.Pointer(memmap.PtrUint16(6112660, 1115324))), (*int32)(unsafe.Pointer(&a1)), nil, 0)
					if int32(uintptr(unsafe.Pointer(a1))) <= v9 {
						goto LABEL_16
					}
				}
			LABEL_15:
				v6 = uint32(nox_draw_getFontAdvance_43F9E0(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 236))))), v3, v9))
			}
		LABEL_16:
			nox_wcsncpy((*libc.WChar)(memmap.PtrOff(6112660, 1115324)), v3, v6)
			v3 = (*libc.WChar)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(libc.WChar(0))*uintptr(v6)))
			*memmap.PtrUint16(6112660, v6*2+0x1104BC) = 0
			if *v3 == 32 {
				v3 = (*libc.WChar)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(libc.WChar(0))*1))
			}
			result = (*libc.WChar)(unsafe.Pointer(uintptr(*v12)))
			if v11 >= int32(uintptr(unsafe.Pointer(result))) {
				break
			}
			(*nox_window)(unsafe.Pointer(uintptr(v2))).Func94(asWindowEvent(0x400D, int32(uintptr(memmap.PtrOff(6112660, 1115324))), -1))
			if nox_wcslen(v3) != 0 {
				result = v10
			} else {
				result = nox_wcstok(nil, &v13[0])
				v10 = result
				v3 = result
			}
			v11++
			if result == nil {
				break
			}
			result = v10
		}
	}
	return result
}
