package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	"unsafe"
)

func sub_41F3A0(a1 int32, a2 int32) bool {
	var (
		v2 *libc.WChar
		v4 *libc.WChar
	)
	if sub_41E2F0() == 7 {
		v4 = strMan.GetStringInFile("wolchat.c:LoadingChannels", "C:\\NoxPost\\src\\common\\WolAPI\\wolchnl.c")
		v2 = strMan.GetStringInFile("wolchat.c:PleaseWait", "C:\\NoxPost\\src\\common\\WolAPI\\wolchnl.c")
	} else {
		v4 = strMan.GetStringInFile("noxworld.c:LoadingGames", "C:\\NoxPost\\src\\common\\WolAPI\\wolchnl.c")
		v2 = strMan.GetStringInFile("wolchat.c:PleaseWait", "C:\\NoxPost\\src\\common\\WolAPI\\wolchnl.c")
	}
	NewDialogWindow(nil, (*libc.WChar)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v2)))))), (*libc.WChar)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v4)))))), 288, nil, nil)
	sub_44A4B0()
	return sub_40D2F0(a1, a2) != 0
}
