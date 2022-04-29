package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"unsafe"
)

func sub_468110_wol_prog() int32 {
	var (
		result int32
		v1     *uint32
		v2     *uint32
		v3     *libc.WChar
	)
	result = int32(dword_5d4594_1064192)
	if dword_5d4594_1064192 != 0 {
		v1 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1064192)))).ChildByID(1804)))
		(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v1)))))).Show()
		v2 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1064192)))).ChildByID(1805)))
		v3 = strMan.GetStringInFile("startingdownload", "C:\\NoxPost\\src\\client\\shell\\WolApi\\wolprog.c")
		result = (*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v2)))))).Func94(asWindowEvent(0x4001, int32(uintptr(unsafe.Pointer(v3))), 0))
	}
	return result
}
func sub_468170_wol_prog() int32 {
	var (
		result int32
		v1     *uint32
		v2     int32
		v3     *libc.WChar
		v4     *uint32
		v5     int32
		v6     int32
		v7     int32
		v8     int32
	)
	result = int32(dword_5d4594_1064192)
	if dword_5d4594_1064192 != 0 {
		v6 = 0
		v5 = 0
		v8 = 0
		v7 = 0
		v1 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1064192)))).ChildByID(1804)))
		sub_41E590((*uint32)(unsafe.Pointer(&v6)), (*uint32)(unsafe.Pointer(&v5)), (*uint32)(unsafe.Pointer(&v8)), (*uint32)(unsafe.Pointer(&v7)))
		if v6 != 0 {
			v2 = v6 * 100 / v5
		} else {
			v2 = 0
		}
		(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v1)))))).Func94(asWindowEvent(0x4020, v2, 0))
		v3 = strMan.GetStringInFile("byteprogress", "C:\\NoxPost\\src\\client\\shell\\WolApi\\wolprog.c")
		nox_swprintf((*libc.WChar)(memmap.PtrOff(6112660, 1063680)), v3, v6, v5, v8/3600, v8/60, v8%60, v7/3600, v7/60, v7%60)
		v4 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1064192)))).ChildByID(1805)))
		result = (*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v4)))))).Func94(asWindowEvent(0x4001, int32(uintptr(memmap.PtrOff(6112660, 1063680))), 0))
	}
	return result
}
func sub_4682B0_wol_prog() {
	var (
		v0 *uint32
		v1 *libc.WChar
		v3 *uint32
		v4 *uint32
	)
	if dword_5d4594_1064192 != 0 {
		v0 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1064192)))).ChildByID(1805)))
		v1 = strMan.GetStringInFile("PatchComplete", "C:\\NoxPost\\src\\client\\shell\\WolApi\\wolprog.c")
		if v1 != nil {
			(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v0)))))).Func94(asWindowEvent(0x4001, int32(uintptr(unsafe.Pointer(v1))), 0))
		}
	}
	if sub_41E540() == 0 {
		v3 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1064192)))).ChildByID(1806)))
		(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v3)))))).Hide()
		v4 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1064192)))).ChildByID(1807)))
		(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v4)))))).Show()
		nox_xxx_setContinueMenuOrHost_43DDD0(0)
		nox_game_exit_xxx_43DE60()
	}
}
