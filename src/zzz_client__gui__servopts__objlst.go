package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	noxflags "github.com/noxworld-dev/opennox/v1/common/flags"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"unsafe"
)

func nox_xxx_guiObjlistLoad_4530C0(a1 int32, a2 int32) int32 {
	var (
		v2  int32
		v3  *libc.WChar
		v4  int32
		v5  int32
		v6  int32
		v7  *libc.WChar
		v8  int32
		v9  int32
		v10 int32
		v11 *uint32
		v12 *uint32
		v14 [66]libc.WChar
	)
	v2 = 0
	dword_5d4594_1045468 = uint32(uintptr(unsafe.Pointer(newWindowFromFile("objlst.wnd", unsafe.Pointer(cgoFuncAddr(sub_4533D0))))))
	(*(*int32)(unsafe.Pointer(&dword_5d4594_1045468))).setDraw(sub_453350)
	sub_46B120((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045468)))), (*nox_window)(unsafe.Pointer(uintptr(a1))))
	nox_xxx_wnd_46B280(*(*int32)(unsafe.Pointer(&dword_5d4594_1045468)), a1)
	dword_5d4594_1045464 = uint32(uintptr(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045468)))).ChildByID(1510))))
	sub_4532E0()
	(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045464))))).Func94(asWindowEvent(0x400F, 0, 0))
	if uint32(a2) == 0x1000000 {
		dword_5d4594_1045460 = 0
		v7 = strMan.GetStringInFile("Weapons", "C:\\NoxPost\\src\\client\\Gui\\ServOpts\\objlst.c")
		nox_wcscpy(&v14[0], v7)
		v8 = 4
		v9 = 25
		for {
			v10 = sub_4159F0(v8)
			if v10 != 0 {
				(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045464))))).Func94(asWindowEvent(0x400D, v10, -1))
				v2++
			}
			v8 *= 2
			v9--
			if v9 == 0 {
				break
			}
		}
	} else if uint32(a2) == 0x2000000 {
		dword_5d4594_1045460 = 1
		v3 = strMan.GetStringInFile("servopts.wnd:Armor", "C:\\NoxPost\\src\\client\\Gui\\ServOpts\\objlst.c")
		nox_wcscpy(&v14[0], v3)
		v4 = 1
		v5 = 26
		for {
			v6 = sub_415E80(v4)
			if v6 != 0 {
				(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045464))))).Func94(asWindowEvent(0x400D, v6, -1))
				v2++
			}
			v4 *= 2
			v5--
			if v5 == 0 {
				break
			}
		}
	}
	(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045464))))).Func94(asWindowEvent(0x4001, int32(uintptr(unsafe.Pointer(&v14[0]))), 0))
	v11 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045468)))).ChildByID(1513)))
	(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045464))))).Func94(asWindowEvent(0x4018, int32(uintptr(unsafe.Pointer(v11))), 0))
	v12 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045468)))).ChildByID(1514)))
	(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045464))))).Func94(asWindowEvent(0x4019, int32(uintptr(unsafe.Pointer(v12))), 0))
	*memmap.PtrUint32(6112660, dword_5d4594_1045460*4+0xFF3E0) = uint32(v2)
	sub_453750()
	if !noxflags.HasGame(noxflags.GameHost) || noxflags.HasGame(noxflags.GameFlag15|noxflags.GameFlag16) {
		sub_46AD20(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045468)), 1515, 1533, 0)
	}
	return int32(dword_5d4594_1045468)
}
