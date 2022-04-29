package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	noxflags "github.com/noxworld-dev/opennox/v1/common/flags"
	"image"
	"unsafe"
)

func sub_44B010() {
	var (
		v0 *libc.WChar
		v1 *libc.WChar
	)
	nox_game_checkStateWol_43C260()
	sub_41E300(9)
	sub_44E040()
	if noxflags.HasGame(noxflags.GameFlag26) {
		v0 = strMan.GetStringInFile("LostConnection", "C:\\NoxPost\\src\\client\\Gui\\woldisc.c")
		nox_xxx_networkLog_printf_413D30(CString("%S"), v0)
		if sub_43C710() != 0 {
			nox_xxx_reconStart_41E400()
		} else {
			nox_client_quit_4460C0()
		}
	} else if noxflags.HasGame(noxflags.GameHost) {
		dword_5d4594_830292 = uint32(uintptr(unsafe.Pointer(newWindowFromFile("woldisc.wnd", unsafe.Pointer(cgoFuncAddr(sub_44B0F0))))))
		(*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_830292)))).SetPos(image.Pt(nox_win_width/2-*(*int32)(unsafe.Pointer(uintptr(dword_5d4594_830292 + 24)))/2, nox_win_height/2-*(*int32)(unsafe.Pointer(uintptr(dword_5d4594_830292 + 28)))/2))
	} else {
		v1 = strMan.GetStringInFile("LostConnection", "C:\\NoxPost\\src\\client\\Gui\\woldisc.c")
		nox_xxx_printCentered_445490(v1)
	}
}
