package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	noxflags "github.com/noxworld-dev/opennox/v1/common/flags"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"image"
	"unsafe"
)

var nox_wnd_quitMenu_825760 *nox_window = nil

func nox_xxx_wndLoadQuitMenu_445790() int32 {
	nox_wnd_quitMenu_825760 = newWindowFromFile("QuitMenu.wnd", unsafe.Pointer(cgoFuncAddr(nox_xxx_menuGameOnButton_445840)))
	if nox_wnd_quitMenu_825760 == nil {
		return 0
	}
	nox_wnd_quitMenu_825760.SetAllFuncs(func(arg1 int32, arg2 int32, arg3 int32, arg4 int32) int32 {
		return sub_445BB0()
	}, func(arg1 *nox_window, arg2 unsafe.Pointer) int32 {
		return nox_xxx_wndDrawQuitMenu_445BC0((*uint32)(unsafe.Pointer(arg1)))
	}, nil)
	nox_wnd_quitMenu_825760.draw_data.bg_color = nox_color_black_2650656
	nox_wnd_quitMenu_825760.off_x = (nox_win_width - nox_wnd_quitMenu_825760.width) / 2
	nox_wnd_quitMenu_825760.end_x = nox_wnd_quitMenu_825760.off_x + nox_wnd_quitMenu_825760.width
	if nox_win_height > 768 {
		nox_wnd_quitMenu_825760.off_y = (nox_win_height - nox_wnd_quitMenu_825760.height - nox_win_height/3) / 2
		nox_wnd_quitMenu_825760.end_y = nox_wnd_quitMenu_825760.off_y + nox_wnd_quitMenu_825760.height
	}
	sub_445C40()
	dword_5d4594_825752 = 0
	dword_5d4594_825768 = 0
	var v1 *libc.WChar = strMan.GetStringInFile("Vote", "C:\\NoxPost\\src\\client\\Gui\\guiquit.c")
	nox_swprintf((*libc.WChar)(memmap.PtrOff(6112660, 825772)), libc.CWString("%s"), v1)
	return 1
}
func nox_xxx_menuGameOnButton_445840(a1 *uint32, a2 int32, a3 *int32, a4 int32) int32 {
	var (
		v3     int32
		v4     int32
		result int32
		v6     int32
		v7     int32
		v8     *byte
		v9     int32
		v10    int32
		v12    int32
		v13    *libc.WChar
		v14    int32
		v15    *libc.WChar
		v16    int32
		v17    int32
		v18    int32
		v19    *uint32
		v20    int32
		v21    int32
		v22    *libc.WChar
		v23    *libc.WChar
	)
	if a2 != 0x4007 {
		return 0
	}
	v3 = (*nox_window)(unsafe.Pointer(a3)).ID()
	clientPlaySoundSpecial(766, 100)
	switch v3 {
	case 9001:
		sub_445C40()
		sub_413A00(1)
		if !noxflags.HasGame(noxflags.GameModeCoop) || nox_xxx_playerAnimCheck_4372B0() != 0 {
			sub_445B40()
			v17 = *(*int32)(unsafe.Add(unsafe.Pointer(a3), unsafe.Sizeof(int32(0))*9))
			*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v17))), 0)) = uint8(int8(v17 & 253))
			*(*int32)(unsafe.Add(unsafe.Pointer(a3), unsafe.Sizeof(int32(0))*9)) = v17
			result = 0
		} else {
			v23 = strMan.GetStringInFile("GUIQuit.c:ReallyLoadMessage", "C:\\NoxPost\\src\\client\\Gui\\guiquit.c")
			v15 = strMan.GetStringInFile("SelChar.c:LoadLabel", "C:\\NoxPost\\src\\client\\Gui\\guiquit.c")
			NewDialogWindow(nil, (*libc.WChar)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v15)))))), (*libc.WChar)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v23)))))), 56, unsafe.Pointer(cgoFuncAddr(sub_445B40)), unsafe.Pointer(cgoFuncAddr(sub_445BA0)))
			v16 = *(*int32)(unsafe.Add(unsafe.Pointer(a3), unsafe.Sizeof(int32(0))*9))
			*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v16))), 0)) = uint8(int8(v16 & 253))
			*(*int32)(unsafe.Add(unsafe.Pointer(a3), unsafe.Sizeof(int32(0))*9)) = v16
			result = 0
		}
	case 9002:
		v18 = int32(*(*uint32)(unsafe.Pointer(uintptr(*memmap.PtrUint32(0x852978, 8) + 120))))
		if (v18 & 0x8000) == 0 {
			sub_445C40()
			if noxflags.HasGame(noxflags.GameModeCoop) {
				sub_4DB130(CString("AUTOSAVE"))
				sub_4DB170(1, 0, 0)
			}
			goto LABEL_27
		}
		v19 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(a1)).ChildByID(v3)))
		nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v19)))))), 0)
		v20 = *(*int32)(unsafe.Add(unsafe.Pointer(a3), unsafe.Sizeof(int32(0))*9))
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v20))), 0)) = uint8(int8(v20 & 253))
		*(*int32)(unsafe.Add(unsafe.Pointer(a3), unsafe.Sizeof(int32(0))*9)) = v20
		result = 0
	case 9003:
		sub_445C40()
		if noxflags.HasGame(noxflags.GameModeCoop) {
			nox_savegame_sub_46D580()
		} else {
			nox_xxx_netSavePlayer_41CE00()
		}
		if sub_43C6E0() != 0 {
			goto LABEL_27
		}
		sub_43CF70()
		v12 = *(*int32)(unsafe.Add(unsafe.Pointer(a3), unsafe.Sizeof(int32(0))*9))
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v12))), 0)) = uint8(int8(v12 & 253))
		*(*int32)(unsafe.Add(unsafe.Pointer(a3), unsafe.Sizeof(int32(0))*9)) = v12
		result = 0
	case 9004:
		nox_xxx_wndClearCaptureMain(nox_wnd_quitMenu_825760)
		v22 = strMan.GetStringInFile("GUIQuit.c:ReallyQuitMessage", "C:\\NoxPost\\src\\client\\Gui\\guiquit.c")
		v13 = strMan.GetStringInFile("GUIQuit.c:ReallyQuitTitle", "C:\\NoxPost\\src\\client\\Gui\\guiquit.c")
		NewDialogWindow(nox_wnd_quitMenu_825760, (*libc.WChar)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v13)))))), (*libc.WChar)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v22)))))), 56, unsafe.Pointer(cgoFuncAddr(nox_xxx_quitDialogYes_445B20)), unsafe.Pointer(cgoFuncAddr(nox_xxx_quitDialogNo_445B30)))
		v14 = *(*int32)(unsafe.Add(unsafe.Pointer(a3), unsafe.Sizeof(int32(0))*9))
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v14))), 0)) = uint8(int8(v14 & 253))
		*(*int32)(unsafe.Add(unsafe.Pointer(a3), unsafe.Sizeof(int32(0))*9)) = v14
		result = 0
	case 9005:
		sub_445C40()
		sub_4ADA40()
		v4 = *(*int32)(unsafe.Add(unsafe.Pointer(a3), unsafe.Sizeof(int32(0))*9))
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v4))), 0)) = uint8(int8(v4 & 253))
		*(*int32)(unsafe.Add(unsafe.Pointer(a3), unsafe.Sizeof(int32(0))*9)) = v4
		result = 0
	case 9006:
		goto LABEL_11
	case 9007:
		if noxflags.HasGame(noxflags.GameHost) {
			v8 = (*byte)(unsafe.Pointer(noxServer.getPlayerByID(int32(nox_player_netCode_85319C))))
			nox_xxx_serverHandleClientConsole_443E90((*nox_playerInfo)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v8)))))), 0, nil)
		LABEL_11:
			sub_445C40()
			v9 = *(*int32)(unsafe.Add(unsafe.Pointer(a3), unsafe.Sizeof(int32(0))*9))
			*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v9))), 0)) = uint8(int8(v9 & 253))
			*(*int32)(unsafe.Add(unsafe.Pointer(a3), unsafe.Sizeof(int32(0))*9)) = v9
			result = 0
		} else {
			nox_xxx_netServerCmd_440950_empty()
			sub_445C40()
			v7 = *(*int32)(unsafe.Add(unsafe.Pointer(a3), unsafe.Sizeof(int32(0))*9))
			*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v7))), 0)) = uint8(int8(v7 & 253))
			*(*int32)(unsafe.Add(unsafe.Pointer(a3), unsafe.Sizeof(int32(0))*9)) = v7
			result = 0
		}
	case 9008:
		sub_445C40()
		nox_xxx_guiServerOptsLoad_457500()
		v10 = *(*int32)(unsafe.Add(unsafe.Pointer(a3), unsafe.Sizeof(int32(0))*9))
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v10))), 0)) = uint8(int8(v10 & 253))
		*(*int32)(unsafe.Add(unsafe.Pointer(a3), unsafe.Sizeof(int32(0))*9)) = v10
		result = 0
	case 9009:
		sub_445C40()
		if noxflags.HasGame(noxflags.GameModeQuest) {
			sub_48CB10(4)
		} else {
			sub_48CB10(0)
		}
		v6 = *(*int32)(unsafe.Add(unsafe.Pointer(a3), unsafe.Sizeof(int32(0))*9))
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v6))), 0)) = uint8(int8(v6 & 253))
		*(*int32)(unsafe.Add(unsafe.Pointer(a3), unsafe.Sizeof(int32(0))*9)) = v6
		result = 0
	default:
	LABEL_27:
		v21 = *(*int32)(unsafe.Add(unsafe.Pointer(a3), unsafe.Sizeof(int32(0))*9))
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v21))), 0)) = uint8(int8(v21 & 253))
		*(*int32)(unsafe.Add(unsafe.Pointer(a3), unsafe.Sizeof(int32(0))*9)) = v21
		result = 0
	}
	return result
}
func sub_445B40() unsafe.Pointer {
	var (
		result unsafe.Pointer
		v1     *libc.WChar
		v2     *libc.WChar
	)
	sub_413A00(0)
	result = unsafe.Pointer(uintptr(sub_4DB790(CString("AUTOSAVE"))))
	if result == nil {
		v2 = strMan.GetStringInFile("GUISave.c:SaveErrorTitle", "C:\\NoxPost\\src\\client\\Gui\\guiquit.c")
		v1 = strMan.GetStringInFile("GUISave.c:SaveErrorTitle", "C:\\NoxPost\\src\\client\\Gui\\guiquit.c")
		result = NewDialogWindow(nil, (*libc.WChar)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v1)))))), (*libc.WChar)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v2)))))), 33, nil, nil)
	}
	return result
}
func sub_445C40() {
	var (
		result int32
		v1     *uint32
		v2     *uint32
		v3     *uint32
		v4     *uint32
		v5     *uint32
		v6     *uint32
		v7     *uint32
		v8     *uint32
		v9     *uint32
		v10    *uint32
		v11    *uint32
		v12    *uint32
		v13    *uint32
		v14    *uint32
		v15    *uint32
		v16    *uint32
		v17    *uint32
		v18    *uint32
		v19    *uint32
		v20    *libc.WChar
		v21    *libc.WChar
	)
	if nox_xxx_wndGetFlags_46ADA0(int32(uintptr(unsafe.Pointer(nox_wnd_quitMenu_825760))))&16 != 0 {
		if *memmap.PtrUint32(0x852978, 8) == 0 || !noxflags.HasGame(noxflags.GameModeCoop) || (func() bool {
			result = int32(*(*uint32)(unsafe.Pointer(uintptr(*memmap.PtrUint32(0x852978, 8) + 276))))
			return result != 2
		}()) && result != 1 && result != 51 {
			if sub_45D9B0() != 1 {
				if nox_xxx_checkGameFlagPause_413A50() != 1 {
					clientPlaySoundSpecial(921, 100)
					nox_xxx_wndShowModalMB(nox_wnd_quitMenu_825760)
					nox_wnd_quitMenu_825760.flags |= 8
					nox_xxx_wndSetCaptureMain(nox_wnd_quitMenu_825760)
					if noxflags.HasGame(noxflags.GameModeCoop) {
						v20 = strMan.GetStringInFile("SoloSaveLabel", "C:\\NoxPost\\src\\client\\Gui\\guiquit.c")
						v1 = (*uint32)(unsafe.Pointer(nox_wnd_quitMenu_825760.ChildByID(9003)))
						sub_46AEE0(int32(uintptr(unsafe.Pointer(v1))), int32(uintptr(unsafe.Pointer(v20))))
						v2 = (*uint32)(unsafe.Pointer(nox_wnd_quitMenu_825760.ChildByID(9001)))
						(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v2)))))).Show()
						v3 = (*uint32)(unsafe.Pointer(nox_wnd_quitMenu_825760.ChildByID(9002)))
						(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v3)))))).Show()
						v4 = (*uint32)(unsafe.Pointer(nox_wnd_quitMenu_825760.ChildByID(9007)))
						(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v4)))))).Hide()
						v5 = (*uint32)(unsafe.Pointer(nox_wnd_quitMenu_825760.ChildByID(9008)))
						(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v5)))))).Hide()
						v6 = (*uint32)(unsafe.Pointer(nox_wnd_quitMenu_825760.ChildByID(9009)))
						(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v6)))))).Hide()
						v7 = (*uint32)(unsafe.Pointer(nox_wnd_quitMenu_825760.ChildByID(9004)))
						(*nox_window)(unsafe.Pointer(v7)).SetPos(image.Pt(int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*4))), int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*5)))))
						sub_413A00(1)
						sub_46AB20((*uint32)(unsafe.Pointer(nox_wnd_quitMenu_825760)), 220, 285)
					} else {
						v21 = strMan.GetStringInFile("MultiplayerSaveLabel", "C:\\NoxPost\\src\\client\\Gui\\guiquit.c")
						v8 = (*uint32)(unsafe.Pointer(nox_wnd_quitMenu_825760.ChildByID(9003)))
						sub_46AEE0(int32(uintptr(unsafe.Pointer(v8))), int32(uintptr(unsafe.Pointer(v21))))
						v9 = (*uint32)(unsafe.Pointer(nox_wnd_quitMenu_825760.ChildByID(9001)))
						(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v9)))))).Hide()
						v10 = (*uint32)(unsafe.Pointer(nox_wnd_quitMenu_825760.ChildByID(9002)))
						(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v10)))))).Hide()
						v11 = (*uint32)(unsafe.Pointer(nox_wnd_quitMenu_825760.ChildByID(9007)))
						(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v11)))))).Show()
						nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v11)))))), 1)
						v12 = (*uint32)(unsafe.Pointer(nox_wnd_quitMenu_825760.ChildByID(9008)))
						(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v12)))))).Show()
						v13 = (*uint32)(unsafe.Pointer(nox_wnd_quitMenu_825760.ChildByID(9009)))
						(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v13)))))).Show()
						sub_46AEE0(int32(uintptr(unsafe.Pointer(v13))), int32(uintptr(memmap.PtrOff(6112660, 825772))))
						if noxflags.HasGame(noxflags.GameFlag15|noxflags.GameFlag16) || int32(nox_xxx_getTeamCounter_417DD0()) == 0 {
							nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v13)))))), 0)
						} else {
							nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v13)))))), 1)
						}
						v14 = (*uint32)(unsafe.Pointer(nox_wnd_quitMenu_825760.ChildByID(9004)))
						(*nox_window)(unsafe.Pointer(v14)).SetPos(image.Pt(int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v13), unsafe.Sizeof(uint32(0))*4))), int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v13), unsafe.Sizeof(uint32(0))*5))+45)))
						sub_46AB20((*uint32)(unsafe.Pointer(nox_wnd_quitMenu_825760)), 220, 330)
						if noxflags.HasGame(noxflags.GameModeQuest) {
							v15 = (*uint32)(unsafe.Pointer(nox_wnd_quitMenu_825760.ChildByID(9007)))
							nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v15)))))), 0)
							v16 = (*uint32)(unsafe.Pointer(nox_wnd_quitMenu_825760.ChildByID(9003)))
							nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v16)))))), 0)
						}
						if nox_common_getEngineFlag(nox_engine_flag(NOX_ENGINE_FLAG_DISABLE_GRAPHICS_RENDERING)) {
							v17 = (*uint32)(unsafe.Pointer(nox_wnd_quitMenu_825760.ChildByID(9007)))
							nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v17)))))), 0)
							v18 = (*uint32)(unsafe.Pointer(nox_wnd_quitMenu_825760.ChildByID(9005)))
							nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v18)))))), 0)
							v19 = (*uint32)(unsafe.Pointer(nox_wnd_quitMenu_825760.ChildByID(9003)))
							nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v19)))))), 0)
						}
					}
				}
			}
		}
	} else {
		guiFocus(nil)
		nox_xxx_wndClearCaptureMain(nox_wnd_quitMenu_825760)
		nox_wnd_quitMenu_825760.Hide()
		nox_wnd_quitMenu_825760.flags &= -9
		sub_413A00(0)
	}
}
func sub_446190() {
	var (
		v1 *libc.WChar
		v2 *libc.WChar
		v3 *libc.WChar
		v4 int8
	)
	dword_5d4594_825764 = 0
	if noxflags.HasGame(noxflags.GameHost) {
		v1 = strMan.GetStringInFile("ServerManualShutdown", "C:\\NoxPost\\src\\client\\Gui\\guiquit.c")
		nox_xxx_networkLog_printf_413D30(CString("%S"), v1)
	}
	if noxflags.HasGame(noxflags.GameFlag26) {
		nox_game_checkStateWol_43C260()
		sub_41E300(9)
		nox_xxx____setargv_4_44B000()
	} else if nox_xxx_isQuest_4D6F50() != 0 {
		if sub_4D6F30() != 0 {
			nox_game_checkStateMenu_43C2F0()
		}
		sub_4D70B0()
		sub_4D6F40(0)
		sub_4D6F90(0)
	}
	if dword_5d4594_825768 == 0 {
		dword_5d4594_825752 = 0
		dword_5d4594_825768 = nox_frame_xxx_2598000
		if noxflags.HasGame(noxflags.GameHost) {
			if nox_xxx_check_flag_aaa_43AF70() == 1 && noxflags.HasGame(noxflags.GameFlag15|noxflags.GameFlag16) {
				sub_416150(15, 0)
			}
			sub_509CB0()
		}
		if noxflags.HasGame(noxflags.GameHost) {
			dword_5d4594_825752 = uint32(nox_common_playerInfoCount_416F40() - 1)
			sub_467440(0)
			if dword_5d4594_825752 > 0 {
				v4 = -57
				nox_xxx_netSendPacket0_4E5420(159, unsafe.Pointer(&v4), 1, 0, 1)
				v3 = strMan.GetStringInFile("ShuttingDown", "C:\\NoxPost\\src\\client\\Gui\\guiquit.c")
				v2 = strMan.GetStringInFile("Wolchat.c:PleaseWait", "C:\\NoxPost\\src\\client\\Gui\\guiquit.c")
				NewDialogWindow(nil, (*libc.WChar)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v2)))))), (*libc.WChar)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v3)))))), 0, nil, nil)
				sub_44A360(1)
				goto LABEL_20
			}
		} else if !noxflags.HasGame(noxflags.GameOnline) {
			goto LABEL_20
		}
		sub_446380()
	}
LABEL_20:
	if dword_5d4594_2650652 != 0 {
		if sub_41E2F0() == 9 {
			sub_41F4B0()
			sub_41EC30()
			sub_446490(0)
			nox_xxx____setargv_4_44B000()
		}
	}
}
