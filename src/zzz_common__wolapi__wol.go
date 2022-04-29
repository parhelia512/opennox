package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	noxflags "github.com/noxworld-dev/opennox/v1/common/flags"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"unsafe"
)

func nox_xxx_wolApiError_41D1D0(a1 int32) *uint32 {
	var (
		v1     *libc.WChar
		v2     *libc.WChar
		v3     int32
		v4     int32
		v5     uint32
		v6     *uint8
		result *uint32
	)
	v1 = strMan.GetStringInFile("WolApiError", "C:\\NoxPost\\src\\common\\WolAPI\\wol.c")
	if *(*int32)(unsafe.Pointer(&dword_5d4594_2660652)) > -2147221247 {
		if dword_5d4594_2660652 != 0 {
			v2 = strMan.GetStringInFile("UnknownError", "C:\\NoxPost\\src\\common\\WolAPI\\wol.c")
		} else {
			v2 = strMan.GetStringInFile("Internalerror", "C:\\NoxPost\\src\\common\\WolAPI\\wol.c")
		}
		v6 = (*uint8)(unsafe.Pointer(v2))
	} else if *(*int32)(unsafe.Pointer(&dword_5d4594_2660652)) >= -2147221248 {
		v6 = (*uint8)(memmap.PtrOff(0x85B3FC, 10436))
	} else {
		switch uint32(*(*int32)(unsafe.Pointer(&dword_5d4594_2660652))) + 0x7FFBFF9C {
		case 0:
			v2 = strMan.GetStringInFile("NickInUse", "C:\\NoxPost\\src\\common\\WolAPI\\wol.c")
		case 1:
			v2 = strMan.GetStringInFile("InvalidPass", "C:\\NoxPost\\src\\common\\WolAPI\\wol.c")
		case 3:
			v2 = strMan.GetStringInFile("ConnectionNetDown", "C:\\NoxPost\\src\\common\\WolAPI\\wol.c")
		case 4:
			v2 = strMan.GetStringInFile("ConnectionLookupError", "C:\\NoxPost\\src\\common\\WolAPI\\wol.c")
		case 5:
			v2 = strMan.GetStringInFile("ConnectionError", "C:\\NoxPost\\src\\common\\WolAPI\\wol.c")
		case 6:
			v2 = strMan.GetStringInFile("Timeout", "C:\\NoxPost\\src\\common\\WolAPI\\wol.c")
		case 8:
			v2 = strMan.GetStringInFile("registryhosed", "C:\\NoxPost\\src\\common\\WolAPI\\wol.c")
		case 14:
			v3 = sub_4207E0()
			v4 = v3
			if v3 != 0 {
				v5 = uint32(libc.StrCSpn((*byte)(unsafe.Pointer(uintptr(v3+24))), CString(":")))
				if libc.StrNCmp(CString("Official Nox Games for New Players"), (*byte)(unsafe.Pointer(uintptr(v5+uint32(v4)+25))), 34) == 0 {
					v1 = strMan.GetStringInFile("Guiserv.c:Notice", "C:\\NoxPost\\src\\common\\WolAPI\\wol.c")
					v2 = strMan.GetStringInFile("TooSkilled", "C:\\NoxPost\\src\\common\\WolAPI\\wol.c")
				} else {
					v2 = strMan.GetStringInFile("banned", "C:\\NoxPost\\src\\common\\WolAPI\\wol.c")
				}
			} else {
				v2 = strMan.GetStringInFile("banned", "C:\\NoxPost\\src\\common\\WolAPI\\wol.c")
			}
		case 16:
			v2 = strMan.GetStringInFile("DisabledError", "C:\\NoxPost\\src\\common\\WolAPI\\wol.c")
		case 17:
			v2 = strMan.GetStringInFile("serialbanned", "C:\\NoxPost\\src\\common\\WolAPI\\wol.c")
		case 18:
			v2 = strMan.GetStringInFile("serialdupe", "C:\\NoxPost\\src\\common\\WolAPI\\wol.c")
		default:
			v2 = strMan.GetStringInFile("UnknownError", "C:\\NoxPost\\src\\common\\WolAPI\\wol.c")
		}
		v6 = (*uint8)(unsafe.Pointer(v2))
	}
	sub_44A400()
	if noxflags.HasGame(noxflags.GameFlag26) {
		nox_xxx_networkLog_printf_413D30(CString("%S %S"), v1, v6)
		nox_xxx_setContinueMenuOrHost_43DDD0(0)
		nox_game_exit_xxx_43DE60()
	}
	if a1 != 0 {
		result = sub_44AFB0(int32(uintptr(unsafe.Pointer(v1))), int32(uintptr(unsafe.Pointer(v6))), *(*int32)(unsafe.Pointer(&dword_5d4594_2660652)))
	} else {
		result = (*uint32)(NewDialogWindow(nil, (*libc.WChar)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v1)))))), (*libc.WChar)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v6)))))), 33, nil, nil))
	}
	return result
}
