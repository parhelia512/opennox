package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	"unsafe"
)

func nox_client_showConnError_43D0A0(a1 int32) {
	var (
		v1 *libc.WChar
		v2 *libc.WChar
	)
	v1 = strMan.GetStringInFile("ConnectionError", "C:\\NoxPost\\src\\Client\\Network\\netclint.c")
	switch a1 + 23 {
	case 0:
		v2 = strMan.GetStringInFile("ConnectAckTimeout", "C:\\NoxPost\\src\\Client\\Network\\netclint.c")
	case 1:
		v2 = strMan.GetStringInFile("SocketCreate", "C:\\NoxPost\\src\\Client\\Network\\netclint.c")
	case 2:
		v2 = strMan.GetStringInFile("WinsockInit", "C:\\NoxPost\\src\\Client\\Network\\netclint.c")
	case 3:
		v2 = strMan.GetStringInFile("VersionConflict", "C:\\NoxPost\\src\\Client\\Network\\netclint.c")
	case 4:
		v2 = strMan.GetStringInFile("Timeout", "C:\\NoxPost\\src\\Client\\Network\\netclint.c")
	case 5:
		v2 = strMan.GetStringInFile("JoinConnTooManyUsers", "C:\\NoxPost\\src\\Client\\Network\\netclint.c")
	case 8:
		v2 = strMan.GetStringInFile("InvalidPort", "C:\\NoxPost\\src\\Client\\Network\\netclint.c")
	case 10:
		v2 = strMan.GetStringInFile("JoinConnUserNotAllowed", "C:\\NoxPost\\src\\Client\\Network\\netclint.c")
	case 11:
		v2 = strMan.GetStringInFile("JoinConnUserBanned", "C:\\NoxPost\\src\\Client\\Network\\netclint.c")
	case 17:
		v2 = strMan.GetStringInFile("JoinConnRefused", "C:\\NoxPost\\src\\Client\\Network\\netclint.c")
	case 19:
		v2 = strMan.GetStringInFile("InvalidAddress", "C:\\NoxPost\\src\\Client\\Network\\netclint.c")
	case 20:
		v2 = strMan.GetStringInFile("InvalidHandle", "C:\\NoxPost\\src\\Client\\Network\\netclint.c")
	default:
		v2 = strMan.GetStringInFile("UnknownConnError", "C:\\NoxPost\\src\\Client\\Network\\netclint.c")
	}
	NewDialogWindow(nil, (*libc.WChar)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v1)))))), (*libc.WChar)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v2)))))), 33, nil, nil)
	sub_44A360(1)
}
func sub_43D260(a1 int32, a2 int32) {
	var (
		v2 *libc.WChar
		v3 [256]libc.WChar
	)
	if a1 != 0 {
		if a2 != 0 {
			v2 = strMan.GetStringInFile("WolPage", "C:\\NoxPost\\src\\Client\\Network\\netclint.c")
			nox_swprintf(&v3[0], v2, a1, a2)
			nox_xxx_printCentered_445490(&v3[0])
		}
	}
}
