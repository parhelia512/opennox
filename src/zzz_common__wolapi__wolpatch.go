package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
)

func sub_41E5D0() *libc.WChar {
	var result *libc.WChar
	switch *memmap.PtrUint32(0x587000, 59168) {
	case 1:
		result = strMan.GetStringInFile("NoSuchServer", "C:\\NoxPost\\src\\common\\WolAPI\\wolpatch.c")
	case 2:
		result = strMan.GetStringInFile("CouldNotConnect", "C:\\NoxPost\\src\\common\\WolAPI\\wolpatch.c")
	case 3:
		result = strMan.GetStringInFile("LoginFailed", "C:\\NoxPost\\src\\common\\WolAPI\\wolpatch.c")
	case 4:
		result = strMan.GetStringInFile("NoSuchFile", "C:\\NoxPost\\src\\common\\WolAPI\\wolpatch.c")
	case 5:
		result = strMan.GetStringInFile("LocalFileOpenFailed", "C:\\NoxPost\\src\\common\\WolAPI\\wolpatch.c")
	case 6:
		result = strMan.GetStringInFile("InternalError", "C:\\NoxPost\\src\\common\\WolAPI\\wolpatch.c")
	case 7:
		result = strMan.GetStringInFile("disconnecterror", "C:\\NoxPost\\src\\common\\WolAPI\\wolpatch.c")
	default:
		result = strMan.GetStringInFile("UnknownError", "C:\\NoxPost\\src\\common\\WolAPI\\wolpatch.c")
	}
	return result
}
