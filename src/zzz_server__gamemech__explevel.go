package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	noxflags "github.com/noxworld-dev/opennox/v1/common/flags"
	"unsafe"
)

func sub_56F980(a1 int32, a2 uint8) *uint32 {
	var (
		result *uint32
		v3     int32
	)
	result = (*uint32)(unsafe.Pointer(uintptr(a1)))
	if uint32(a1) >= 0x2734945F {
		result = sub_56F590(a1)
		if result != nil {
			dword_5d4594_2516328 ^= *(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*1))
			v3 = int32(dword_5d4594_2516348 ^ (uint32(a2) + (dword_5d4594_2516348 ^ *(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*1)))))
			*(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*1)) = uint32(v3)
			dword_5d4594_2516328 ^= uint32(v3)
			result = (*uint32)(unsafe.Pointer(uintptr(nox_xxx_protectData_56F5C0())))
		}
	}
	return result
}
func sub_4EF2E0_exp_level(a1 int32) {
	var (
		v1 int32
		v2 *libc.WChar
	)
	v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))) + 276))))
	if (nox_xxx_gameGet_4DB1B0() != 1 || sub_4DB1C0() == 0) && nox_xxx_gamedataGetFloatTable_419D70(CString("XPTable"), int32(*(*byte)(unsafe.Pointer(uintptr(v1 + 3684)))+1)) <= float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 28)))) {
		*(*uint8)(unsafe.Pointer(uintptr(v1 + 3684)))++
		sub_56F980(int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 4644)))), 1)
		nox_xxx_plrReadVals_4EEDC0(a1, 1)
		if noxflags.HasGame(noxflags.GameModeCoop) {
			sub_57AF30(a1, 0)
		} else {
			nox_xxx_aud_501960(902, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 2, int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 36)))))
			v2 = strMan.GetStringInFile("LevelUP", "C:\\NoxPost\\src\\Server\\GameMech\\explevel.c")
			nox_xxx_netSendLineMessage_4D9EB0(a1, v2)
		}
	}
}
func nox_xxx_plyrGiveExp_4EF3A0_exp_level(a1 int32, a2 float32) {
	var (
		v2 int32
		v3 *libc.WChar
	)
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
	*(*float32)(unsafe.Pointer(uintptr(a1 + 28))) = a2 + *(*float32)(unsafe.Pointer(uintptr(a1 + 28)))
	sub_56FA40(int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 276))) + 4604)))), a2)
	sub_4D81A0(a1)
	v3 = strMan.GetStringInFile("health.c:gainpoints", "C:\\NoxPost\\src\\Server\\GameMech\\explevel.c")
	nox_xxx_netSendLineMessage_4D9EB0(a1, v3, uint32(int32(int64(a2))))
	sub_4EF2E0_exp_level(a1)
}
