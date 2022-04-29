package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	noxflags "github.com/noxworld-dev/opennox/v1/common/flags"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"unsafe"
)

func sub_409FB0_settings(a1 int16, a2 uint16) {
	var (
		v2 int32
		v3 uint16
		v4 *libc.WChar
	)
	v2 = sub_409A70(a1)
	v3 = a2
	if int32(*memmap.PtrUint16(6112660, uint32(v2*2+3488))) != int32(a2) {
		if int32(a2) > 999 {
			v3 = 999
		}
		*memmap.PtrUint16(6112660, uint32(v2*2+3488)) = v3
		nox_server_gameSettingsUpdated = 1
		if nox_client_isConnected_43C700() != 0 {
			v4 = strMan.GetStringInFile("parsecmd.c:lessonsset", "C:\\NoxPost\\src\\Common\\System\\settings.c")
			sub_440A20(v4, v3)
		}
	}
}
func sub_40A040_settings(a1 int16, a2 uint8) {
	var (
		v2 int32
		v3 *libc.WChar
		v4 int32
		v5 *libc.WChar
		v6 *libc.WChar
	)
	v2 = sub_409A70(a1)
	if int32(*memmap.PtrUint8(6112660, uint32(v2+3500))) == int32(a2) {
		return
	}
	if !(!noxflags.HasGame(noxflags.GameSuddenDeath) && dword_5d4594_3592 == 0) {
		v6 = strMan.GetStringInFile("NotInSuddenDeath", "C:\\NoxPost\\src\\Common\\System\\settings.c")
		nox_xxx_printCentered_445490(v6)
		return
	}
	nox_server_gameSettingsUpdated = 1
	if nox_client_isConnected_43C700() != 0 {
		if int32(a2) == 0 {
			v5 = strMan.GetStringInFile("parsecmd.c:timedisabled", "C:\\NoxPost\\src\\Common\\System\\settings.c")
			sub_440A20(v5)
		} else {
			v3 = strMan.GetStringInFile("parsecmd.c:timeset", "C:\\NoxPost\\src\\Common\\System\\settings.c")
			sub_440A20(v3, a2)
			*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v4))), 0)) = uint8(int8(nox_xxx_gamePlayIsAnyPlayers_40A8A0()))
			if v4 == 0 {
				v5 = strMan.GetStringInFile("TimeLimitDeferred", "C:\\NoxPost\\src\\Common\\System\\settings.c")
				sub_440A20(v5)
			}
		}
	}
	*memmap.PtrUint8(6112660, uint32(v2+3500)) = a2
	*memmap.PtrUint64(6112660, 3468) = uint64(uint32(int32(a2)*60000) + platformTicks())
}
