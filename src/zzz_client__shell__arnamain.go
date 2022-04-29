package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	noxflags "github.com/noxworld-dev/opennox/v1/common/flags"
	"unsafe"
)

func nox_xxx_windowArenaSub_4AA4D0(a1 int32, a2 uint32, a3 *int32, a4 int32) int32 {
	if a2 == 0x4005 {
		clientPlaySoundSpecial(920, 100)
		return 1
	}
	if a2 != 0x4007 {
		return 0
	}
	if int32(*(*uint8)(unsafe.Pointer(uintptr(dword_5d4594_1309708 + 64)))) != 0 && !noxflags.HasGame(noxflags.GameFlag26) && sub_4D6F30() == 0 {
		clientPlaySoundSpecial(921, 100)
		return 1
	}
	var sid int32 = (*nox_window)(unsafe.Pointer(a3)).ID()
	if sid != 411 {
		if sid != 421 {
			clientPlaySoundSpecial(921, 100)
			return 1
		}
		sub_4AA450()
		*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1309708 + 52))) = uint32(cgoFuncAddr(nox_game_showGameSel_4379F0))
		sub_43AF50(0)
		noxflags.UnsetGame(noxflags.GameFlag25)
		clientPlaySoundSpecial(921, 100)
		return 1
	}
	if sub_41D440() == 0 {
		var (
			v7 *libc.WChar = strMan.GetStringInFile("WOLAPIINIT_Failed", "C:\\NoxPost\\src\\client\\shell\\ArnaMain.c")
			v5 *libc.WChar = strMan.GetStringInFile("Error", "C:\\NoxPost\\src\\client\\shell\\ArnaMain.c")
		)
		NewDialogWindow((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&nox_win_onlineOrLAN_1309716))))), (*libc.WChar)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v5)))))), (*libc.WChar)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v7)))))), 33, nil, nil)
		return 0
	}
	if sub_40D5B0(int32(uintptr(unsafe.Pointer(&a2)))) != 0 && a2 < 65545 {
		sub_41D4C0()
		var v8 *libc.WChar = strMan.GetStringInFile("Wol.c:RegistryHosed", "C:\\NoxPost\\src\\client\\shell\\ArnaMain.c")
		var v6 *libc.WChar = strMan.GetStringInFile("Error", "C:\\NoxPost\\src\\client\\shell\\ArnaMain.c")
		NewDialogWindow((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&nox_win_onlineOrLAN_1309716))))), (*libc.WChar)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v6)))))), (*libc.WChar)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v8)))))), 33, nil, nil)
		return 0
	}
	sub_4A1A40(0)
	sub_4AA450()
	*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1309708 + 52))) = uint32(cgoFuncAddr(nox_game_showWolLogin_44A560))
	noxflags.UnsetGame(noxflags.GameFlag25)
	clientPlaySoundSpecial(921, 100)
	return 1
}
