package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"unsafe"
)

func sub_41DAC0(a1 int32) *libc.WChar {
	var result *libc.WChar
	switch uint32(a1) + 0x7FFBFF9A {
	case 0:
		result = strMan.GetStringInFile("nonesuch", "C:\\NoxPost\\src\\common\\WolAPI\\wolstate.c")
	case 8:
		result = strMan.GetStringInFile("channelfull", "C:\\NoxPost\\src\\common\\WolAPI\\wolstate.c")
	case 10:
		result = strMan.GetStringInFile("channeldoesnotexist", "C:\\NoxPost\\src\\common\\WolAPI\\wolstate.c")
	case 11:
		result = strMan.GetStringInFile("badpassword", "C:\\NoxPost\\src\\common\\WolAPI\\wolstate.c")
	case 12:
		result = strMan.GetStringInFile("bannedfromchannel", "C:\\NoxPost\\src\\common\\WolAPI\\wolstate.c")
	default:
		result = strMan.GetStringInFile("unknownerror", "C:\\NoxPost\\src\\common\\WolAPI\\wolstate.c")
	}
	if result != nil {
		result = (*libc.WChar)(unsafe.Pointer(uintptr(sub_447310(0, int32(uintptr(unsafe.Pointer(result)))))))
	}
	return result
}
func sub_41DDB0() int32 {
	var (
		v0 int32
		v1 int32
		v2 *libc.WChar
		v3 *libc.WChar
	)
	if *memmap.PtrUint32(0x587000, 58232) == 0 {
		sub_40D250()
		return 1
	}
	v0 = int32(*memmap.PtrUint32(0x587000, 58232))
	for {
		v1 = sub_41E2F0()
		switch sub_41DCC0(v1) {
		case 5:
			if dword_5d4594_2660652 == 0 {
				sub_468F30_wol_reg()
				break
			}
			if int32(*memmap.PtrUint16(0x85B3FC, 10436)) != 0 {
				sub_468DF0((*libc.WChar)(memmap.PtrOff(0x85B3FC, 10436)))
				break
			}
			switch dword_5d4594_2660652 + 0x7FFBFF9B {
			case 0:
				v2 = strMan.GetStringInFile("ServerBusy", "C:\\NoxPost\\src\\common\\WolAPI\\wolstate.c")
				nox_wcscpy((*libc.WChar)(memmap.PtrOff(0x85B3FC, 10436)), v2)
			case 1:
				v3 = strMan.GetStringInFile("Timeout", "C:\\NoxPost\\src\\common\\WolAPI\\wolstate.c")
				nox_swprintf((*libc.WChar)(memmap.PtrOff(0x85B3FC, 10436)), v3)
			case 155:
				v3 = strMan.GetStringInFile("InvalidField", "C:\\NoxPost\\src\\common\\WolAPI\\wolstate.c")
				nox_swprintf((*libc.WChar)(memmap.PtrOff(0x85B3FC, 10436)), v3)
			case 156:
				v3 = strMan.GetStringInFile("CantVerify", "C:\\NoxPost\\src\\common\\WolAPI\\wolstate.c")
				nox_swprintf((*libc.WChar)(memmap.PtrOff(0x85B3FC, 10436)), v3)
			default:
				v3 = strMan.GetStringInFile("NetError", "C:\\NoxPost\\src\\common\\WolAPI\\wolstate.c")
				nox_swprintf((*libc.WChar)(memmap.PtrOff(0x85B3FC, 10436)), v3)
			}
			sub_468DF0((*libc.WChar)(memmap.PtrOff(0x85B3FC, 10436)))
		case 20:
			sub_468E60(0)
		case 21:
			fallthrough
		case 22:
			sub_468E60(1)
		default:
		}
		v0--
		if v0 == 0 {
			break
		}
	}
	sub_40D250()
	return 1
}
