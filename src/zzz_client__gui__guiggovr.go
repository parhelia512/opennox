package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"image"
	"unsafe"
)

func sub_49B4B0(a1 *uint16) int32 {
	var (
		v1     *libc.WChar
		v2     *libc.WChar
		v3     *libc.WChar
		v4     *uint32
		v5     *uint32
		v6     *uint32
		v7     *uint32
		v8     *uint32
		v9     *uint32
		result int32
		v11    int32
		v12    int32
		v13    int32
		v14    int32
		v15    int32
	)
	(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1303452))))).Show()
	nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1303452))))), 1)
	clientPlaySoundSpecial(1007, 100)
	nox_window_get_size((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1303452))))), &v15, &v14)
	(*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1303452)))).SetPos(image.Pt(nox_win_width/2-v15/2, nox_win_height/2-v14/2))
	v11 = int32(*(*uint16)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint16(0))*1)))
	v1 = strMan.GetStringInFile("GGOver.wnd:GeneratorsDestroyed", "C:\\NoxPost\\src\\client\\Gui\\GUIGGOvr.c")
	nox_swprintf((*libc.WChar)(memmap.PtrOff(6112660, 1302172)), v1, v11)
	v12 = int32(*(*uint16)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint16(0))*2)))
	v2 = strMan.GetStringInFile("GGOver.wnd:NumSecretsFound", "C:\\NoxPost\\src\\client\\Gui\\GUIGGOvr.c")
	nox_swprintf((*libc.WChar)(memmap.PtrOff(6112660, 1301916)), v2, v12)
	v13 = int32(*(*uint16)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint16(0))*3)))
	v3 = strMan.GetStringInFile("GGOver.wnd:Kills", "C:\\NoxPost\\src\\client\\Gui\\GUIGGOvr.c")
	nox_swprintf((*libc.WChar)(memmap.PtrOff(6112660, 1302428)), v3, v13)
	nox_swprintf((*libc.WChar)(memmap.PtrOff(6112660, 1303196)), (*libc.WChar)(memmap.PtrOff(6112660, 1303460)))
	v4 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1303452)))).ChildByID(10710)))
	sub_46AEE0(int32(uintptr(unsafe.Pointer(v4))), int32(uintptr(memmap.PtrOff(6112660, 1302172))))
	v5 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1303452)))).ChildByID(0x29D1)))
	sub_46AEE0(int32(uintptr(unsafe.Pointer(v5))), int32(uintptr(memmap.PtrOff(6112660, 1302940))))
	v6 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1303452)))).ChildByID(0x29D2)))
	sub_46AEE0(int32(uintptr(unsafe.Pointer(v6))), int32(uintptr(memmap.PtrOff(6112660, 1302684))))
	v7 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1303452)))).ChildByID(0x29D3)))
	sub_46AEE0(int32(uintptr(unsafe.Pointer(v7))), int32(uintptr(memmap.PtrOff(6112660, 1301916))))
	v8 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1303452)))).ChildByID(0x29D4)))
	sub_46AEE0(int32(uintptr(unsafe.Pointer(v8))), int32(uintptr(memmap.PtrOff(6112660, 1302428))))
	v9 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1303452)))).ChildByID(0x29D7)))
	sub_46AEE0(int32(uintptr(unsafe.Pointer(v9))), int32(uintptr(memmap.PtrOff(6112660, 1303196))))
	result = int32(nox_frame_xxx_2598000)
	*memmap.PtrUint32(6112660, 1303456) = nox_frame_xxx_2598000
	return result
}
func sub_49B6E0() int32 {
	var (
		result int32
		v1     int32
		v2     *libc.WChar
		v3     *uint32
		v4     int32
	)
	result = int32(dword_5d4594_1303452)
	if dword_5d4594_1303452 != 0 {
		result = (*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1303452))))).Flags().IsHidden()
		if result == 0 {
			v1 = int32(*memmap.PtrUint32(6112660, 1303456) + nox_gameFPS*30 - nox_frame_xxx_2598000)
			if v1 < 0 {
				v1 = 0
			}
			if *memmap.PtrUint32(0x8531A0, 2576) != 0 && int32(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(0x8531A0, 2576) + 2064)))) == 31 {
				nox_wcscpy((*libc.WChar)(memmap.PtrOff(6112660, 1301852)), (*libc.WChar)(memmap.PtrOff(6112660, 1303464)))
			} else {
				v4 = int32(uint32(v1) / nox_gameFPS)
				v2 = strMan.GetStringInFile("Rules.c:Time", "C:\\NoxPost\\src\\client\\Gui\\GUIGGOvr.c")
				nox_swprintf((*libc.WChar)(memmap.PtrOff(6112660, 1301852)), libc.CWString("%s - %d"), v2, v4)
			}
			v3 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1303452)))).ChildByID(0x29D8)))
			result = sub_46AEE0(int32(uintptr(unsafe.Pointer(v3))), int32(uintptr(memmap.PtrOff(6112660, 1301852))))
		}
	}
	return result
}
