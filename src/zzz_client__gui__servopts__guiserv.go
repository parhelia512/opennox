package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	noxflags "github.com/noxworld-dev/opennox/v1/common/flags"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"image"
	"math"
	"unsafe"
)

type nox_gui_gamemode struct {
	name  *byte
	title *libc.WChar
	mode  uint32
	hide  bool
}

var nox_gui_gamemodes [8]nox_gui_gamemode = [8]nox_gui_gamemode{{name: CString("CTF"), title: nil, mode: 32, hide: false}, {name: CString("Arena"), title: nil, mode: 256, hide: false}, {name: CString("Highlander"), title: nil, mode: 1024, hide: false}, {name: CString("KotR"), title: nil, mode: 16, hide: false}, {name: CString("Flagball"), title: nil, mode: 64, hide: false}, {name: CString("Quest"), title: nil, mode: 4096, hide: true}, {name: CString("Noxworld.c:Chat"), title: nil, mode: 128, hide: false}, {}}
var nox_gui_gamemode_cnt int32 = int32(unsafe.Sizeof([8]nox_gui_gamemode{})/unsafe.Sizeof(nox_gui_gamemode{}) - 1)
var nox_gui_gamemode_loaded_1046548 int32 = 0

func nox_gui_gamemode_load_457410() {
	if nox_gui_gamemode_loaded_1046548 != 0 {
		return
	}
	for i := int32(0); i < nox_gui_gamemode_cnt; i++ {
		var p *nox_gui_gamemode = &nox_gui_gamemodes[i]
		if p.name == nil {
			break
		}
		p.title = strMan.GetStringInFile(p.name, "C:\\NoxPost\\src\\client\\Gui\\ServOpts\\guiserv.c")
	}
	nox_gui_gamemode_loaded_1046548 = 1
}
func nox_xxx_guiServerOptionsGetGametypeName_4573C0(mode int16) *libc.WChar {
	mode &= 6128
	if nox_gui_gamemode_loaded_1046548 == 0 {
		nox_gui_gamemode_load_457410()
	}
	for i := int32(0); i < nox_gui_gamemode_cnt; i++ {
		var p *nox_gui_gamemode = &nox_gui_gamemodes[i]
		if p.mode == uint32(mode) {
			return p.title
		}
	}
	return nox_gui_gamemodes[1].title
}
func sub_457A10() {
	var (
		v0 *uint32
		v1 *uint32
		v3 int32
		v6 int32
		v7 int32
	)
	v0 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1046492)))).ChildByID(10120)))
	v1 = v0
	var max int32 = 0
	var cnt int32 = 0
	for i := int32(0); i < nox_gui_gamemode_cnt; i++ {
		var p *nox_gui_gamemode = &nox_gui_gamemodes[i]
		if p.hide {
			continue
		}
		cnt++
		var w int32 = 0
		nox_xxx_drawGetStringSize_43F840(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*59)))), p.title, &w, nil, 0)
		if w > max {
			max = w
		}
	}
	v3 = cnt * (nox_xxx_guiFontHeightMB_43F320(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v0), unsafe.Sizeof(uint32(0))*59))))) + 1)
	*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*7)) = *(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*5)) + uint32(v3) + 2
	*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*3)) = uint32(v3 + 2)
	v6 = max + 7
	v7 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*6)) - uint32(v6))
	*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*2)) = uint32(v6)
	*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*4)) = uint32(v7)
}
func sub_459650(title *libc.WChar) int32 {
	for i := int32(0); i < nox_gui_gamemode_cnt; i++ {
		var p *nox_gui_gamemode = &nox_gui_gamemodes[i]
		if nox_wcscmp(p.title, title) == 0 {
			return int32(p.mode)
		}
	}
	return 0
}
func sub_459C10() int32 {
	var v0 *libc.WChar
	v0 = (*libc.WChar)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1046492)))).ChildByID(0x2787)))
	return sub_459650((*libc.WChar)(unsafe.Add(unsafe.Pointer(v0), unsafe.Sizeof(libc.WChar(0))*54)))
}
func nox_xxx_windowServerOptionsFillGametypeList_4596A0() {
	var v0 unsafe.Pointer = unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1046492)))).ChildByID(10120))
	(*nox_window)(v0).Func94(asWindowEvent(0x400F, 0, 0))
	for i := int32(0); i < nox_gui_gamemode_cnt; i++ {
		var p *nox_gui_gamemode = &nox_gui_gamemodes[i]
		if !p.hide {
			(*nox_window)(v0).Func94(asWindowEvent(0x400D, int32(uintptr(unsafe.Pointer(p.title))), -1))
		}
	}
}
func nox_xxx_guiServerOptsLoad_457500() int32 {
	var (
		v1  int32
		v2  *uint32
		v3  *uint32
		v4  *uint32
		v5  *byte
		v6  *uint32
		v7  *uint32
		v8  *uint32
		v9  *byte
		v10 *libc.WChar
		v11 *libc.WChar
		v12 *libc.WChar
		v13 *uint32
		v14 *byte
		v15 *uint32
	)
	if nox_gui_xxx_check_446360() == 0 {
		if dword_5d4594_1046492 != 0 {
			clientPlaySoundSpecial(231, 100)
			nox_xxx_guiServerOptionsTryHide_4574D0()
			return 1
		}
		if noxflags.HasGame(noxflags.GameHost) {
			sub_459D50(1)
			if dword_587000_129656 != 0 {
				nox_common_list_clear_425760((*nox_list_item_t)(memmap.PtrOff(6112660, 1045956)))
			}
		}
		v1 = strMan.Lang()
		if nox_xxx_guiFontHeightMB_43F320(nil) > 10 {
			v1 = 2
		}
		dword_5d4594_1046492 = uint32(uintptr(unsafe.Pointer(newWindowFromFile(*(**byte)(memmap.PtrOff(0x587000, uint32(v1*4)+0x1FAE0)), unsafe.Pointer(cgoFuncAddr(nox_xxx_guiServerOptionsProcPre_4585D0))))))
		nox_draw_setTabWidth_43FE20(100)
		(*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1046492)))).SetPos(image.Pt(int32(uint32(nox_win_width)-*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1046492 + 8)))-10), 0))
		(*(*int32)(unsafe.Pointer(&dword_5d4594_1046492))).setFunc93(nox_xxx_guiServerOptionsProc_458590)
		(*(*int32)(unsafe.Pointer(&dword_5d4594_1046492))).setDraw(func(arg1 int32, arg2 int32) int32 {
			return nox_xxx_windowServerOptionsDrawProc_458500((*uint32)(unsafe.Pointer(uintptr(arg1))), arg2)
		})
		dword_5d4594_1046512 = (*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1046492)))).ChildByID(0x2775)
		dword_5d4594_1046496 = uint32(uintptr(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1046492)))).ChildByID(0x2782))))
		dword_5d4594_1046500 = uint32(uintptr(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1046492)))).ChildByID(0x27C7))))
		dword_5d4594_1046504 = uint32(uintptr(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1046492)))).ChildByID(0x27D5))))
		dword_5d4594_1046508 = uint32(uintptr(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1046492)))).ChildByID(0x27D7))))
		dword_5d4594_1046524 = uint32(uintptr(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1046492)))).ChildByID(10150))))
		dword_5d4594_1046516 = uint32(uintptr(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1046492)))).ChildByID(0x2796))))
		dword_5d4594_1046520 = uint32(uintptr(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1046492)))).ChildByID(0x2797))))
		dword_5d4594_1046536 = uint32(uintptr(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1046492)))).ChildByID(0x27A9))))
		v2 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1046492)))).ChildByID(0x285B)))
		(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v2)))))).setTooltipFunc(unsafe.Pointer(cgoFuncAddr(nox_xxx_options_457AA0)))
		v3 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1046492)))).ChildByID(0x285D)))
		(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v3)))))).setTooltipFunc(unsafe.Pointer(cgoFuncAddr(nox_xxx_options_457B00)))
		nox_xxx_wnd_46B280(*(*int32)(unsafe.Pointer(&dword_5d4594_1046524)), *(*int32)(unsafe.Pointer(&dword_5d4594_1046492)))
		nox_xxx_wnd_46B280(*(*int32)(unsafe.Pointer(&dword_5d4594_1046532)), *(*int32)(unsafe.Pointer(&dword_5d4594_1046492)))
		nox_xxx_wnd_46B280(*(*int32)(unsafe.Pointer(&dword_5d4594_1046536)), *(*int32)(unsafe.Pointer(&dword_5d4594_1046492)))
		nox_xxx_wnd_46B280(*(*int32)(unsafe.Pointer(&dword_5d4594_1046500)), *(*int32)(unsafe.Pointer(&dword_5d4594_1046492)))
		nox_xxx_wnd_46B280(*(*int32)(unsafe.Pointer(&dword_5d4594_1046504)), *(*int32)(unsafe.Pointer(&dword_5d4594_1046492)))
		nox_xxx_wnd_46B280(*(*int32)(unsafe.Pointer(&dword_5d4594_1046508)), *(*int32)(unsafe.Pointer(&dword_5d4594_1046492)))
		(*(*int32)(unsafe.Pointer(&dword_5d4594_1046524))).setFunc94(func(arg1 int32, arg2 int32, arg3 int32, arg4 int32) int32 {
			return nox_xxx_guiServerOptionsProcPre_4585D0(arg1, uint32(arg2), arg3, arg4)
		})
		(*(*int32)(unsafe.Pointer(&dword_5d4594_1046532))).setFunc94(func(arg1 int32, arg2 int32, arg3 int32, arg4 int32) int32 {
			return nox_xxx_guiServerOptionsProcPre_4585D0(arg1, uint32(arg2), arg3, arg4)
		})
		(*(*int32)(unsafe.Pointer(&dword_5d4594_1046536))).setFunc94(func(arg1 int32, arg2 int32, arg3 int32, arg4 int32) int32 {
			return nox_xxx_guiServerOptionsProcPre_4585D0(arg1, uint32(arg2), arg3, arg4)
		})
		(*(*int32)(unsafe.Pointer(&dword_5d4594_1046500))).setFunc94(func(arg1 int32, arg2 int32, arg3 int32, arg4 int32) int32 {
			return nox_xxx_guiServerOptionsProcPre_4585D0(arg1, uint32(arg2), arg3, arg4)
		})
		(*(*int32)(unsafe.Pointer(&dword_5d4594_1046504))).setFunc94(func(arg1 int32, arg2 int32, arg3 int32, arg4 int32) int32 {
			return nox_xxx_guiServerOptionsProcPre_4585D0(arg1, uint32(arg2), arg3, arg4)
		})
		(*(*int32)(unsafe.Pointer(&dword_5d4594_1046508))).setFunc94(func(arg1 int32, arg2 int32, arg3 int32, arg4 int32) int32 {
			return nox_xxx_guiServerOptionsProcPre_4585D0(arg1, uint32(arg2), arg3, arg4)
		})
		*memmap.PtrUint32(6112660, 1046352) = uint32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg("UITabs1"))))
		dword_5d4594_1046356 = uint32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg("UITabs2"))))
		dword_5d4594_1046360 = uint32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg("UITabs3"))))
		v4 = *(**uint32)(unsafe.Pointer(uintptr(dword_5d4594_1046496 + 32)))
		v14 = (*byte)(unsafe.Pointer(nox_xxx_gLoadImg("UISlider")))
		v5 = (*byte)(unsafe.Pointer(nox_xxx_gLoadImg("UISliderLit")))
		v6 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1046500)))).ChildByID(0x27C6)))
		v7 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1046500)))).ChildByID(10180)))
		v15 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1046500)))).ChildByID(0x27C5)))
		*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*100)) + 8))) = 16
		*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*100)) + 12))) = 10
		sub_4B5700(int32(uintptr(unsafe.Pointer(v6))), 0, 0, int32(uintptr(unsafe.Pointer(v14))), int32(uintptr(unsafe.Pointer(v5))), int32(uintptr(unsafe.Pointer(v5))))
		nox_xxx_wnd_46B280(int32(uintptr(unsafe.Pointer(v6))), *(*int32)(unsafe.Pointer(&dword_5d4594_1046496)))
		nox_xxx_wnd_46B280(int32(uintptr(unsafe.Pointer(v7))), *(*int32)(unsafe.Pointer(&dword_5d4594_1046496)))
		nox_xxx_wnd_46B280(int32(uintptr(unsafe.Pointer(v15))), *(*int32)(unsafe.Pointer(&dword_5d4594_1046496)))
		*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*9)) = uint32(uintptr(unsafe.Pointer(v6)))
		*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*7)) = uint32(uintptr(unsafe.Pointer(v7)))
		*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*8)) = uint32(uintptr(unsafe.Pointer(v15)))
		v8 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1046492)))).ChildByID(0x27B0)))
		*(*uint32)(unsafe.Add(unsafe.Pointer(v8), unsafe.Sizeof(uint32(0))*9)) |= 4
		v9 = sub_4165D0(0)
		if noxflags.HasGame(noxflags.GameHost) {
			sub_4161E0()
		}
		nox_gui_gamemode_load_457410()
		sub_457B60(int32(uintptr(unsafe.Pointer(v9))))
		sub_457A10()
		if noxflags.HasGame(noxflags.GameHost) {
			sub_4165F0(0, 1)
			sub_4165D0(1)
		} else {
			v10 = (*libc.WChar)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1046492)))).ChildByID(0x27AF)))
			v11 = strMan.GetStringInFile("servopts.wnd:teams", "C:\\NoxPost\\src\\client\\Gui\\ServOpts\\guiserv.c")
			(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v10)))))).Func94(asWindowEvent(0x4001, int32(uintptr(unsafe.Pointer(v11))), -1))
			v12 = strMan.GetStringInFile("servopts.wnd:TeamTT", "C:\\NoxPost\\src\\client\\Gui\\ServOpts\\guiserv.c")
			nox_xxx_wndWddSetTooltip_46B000((*nox_window_data)(unsafe.Pointer((*libc.WChar)(unsafe.Add(unsafe.Pointer(v10), unsafe.Sizeof(libc.WChar(0))*18)))), v12)
			v13 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1046492)))).ChildByID(0x27A5)))
			(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v13)))))).Show()
			(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1046508))))).Hide()
			sub_46ACE0(*(**uint32)(unsafe.Pointer(&dword_5d4594_1046492)), 0x27A1, 0x27A2, 1)
		}
		if dword_587000_129656 != 0 {
			if sub_4D6F30() != 0 || nox_xxx_isQuest_4D6F50() != 0 {
				nox_server_parseCmdText_443C80(libc.CWString("execrul OTQuest.rul"), 1)
			} else if nox_common_getEngineFlag(nox_engine_flag(NOX_ENGINE_FLAG_DISABLE_GRAPHICS_RENDERING)) {
				nox_server_parseCmdText_443C80(libc.CWString("execrul server.rul"), 1)
			}
		}
		dword_587000_129656 = 0
	}
	return 1
}
func nox_xxx_options_457AA0(a1 int32, a2 *uint8) int32 {
	var v2 *libc.WChar
	if int32(*a2)&4 != 0 {
		v2 = strMan.GetStringInFile("AutoAssignOnTT", "C:\\NoxPost\\src\\client\\Gui\\ServOpts\\guiserv.c")
	} else {
		v2 = strMan.GetStringInFile("AutoAssignOffTT", "C:\\NoxPost\\src\\client\\Gui\\ServOpts\\guiserv.c")
	}
	nox_xxx_cursorSetTooltip_4776B0(v2)
	return 1
}
func nox_xxx_options_457B00(a1 int32, a2 *uint8) int32 {
	var v2 *libc.WChar
	if int32(*a2)&4 != 0 {
		v2 = strMan.GetStringInFile("TeamDamageOnTT", "C:\\NoxPost\\src\\client\\Gui\\ServOpts\\guiserv.c")
	} else {
		v2 = strMan.GetStringInFile("TeamDamageOffTT", "C:\\NoxPost\\src\\client\\Gui\\ServOpts\\guiserv.c")
	}
	nox_xxx_cursorSetTooltip_4776B0(v2)
	return 1
}
func sub_457FE0() int32 {
	var (
		v0 *uint32
		v1 *libc.WChar
		v2 *libc.WChar
		v3 *libc.WChar
		v5 *libc.WChar
		v6 *libc.WChar
	)
	v0 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1046492)))).ChildByID(10210)))
	if noxflags.HasGame(noxflags.GameFlag15) {
		v5 = strMan.GetStringInFile("Servopts.wnd:Individual", "C:\\NoxPost\\src\\client\\Gui\\ServOpts\\guiserv.c")
		v1 = strMan.GetStringInFile("Servopts.wnd:Ladder", "C:\\NoxPost\\src\\client\\Gui\\ServOpts\\guiserv.c")
		nox_swprintf((*libc.WChar)(memmap.PtrOff(6112660, 1045700)), libc.CWString("%s %s"), v1, v5)
	} else if noxflags.HasGame(noxflags.GameFlag16) {
		v6 = strMan.GetStringInFile("Servopts.wnd:Clan", "C:\\NoxPost\\src\\client\\Gui\\ServOpts\\guiserv.c")
		v2 = strMan.GetStringInFile("Servopts.wnd:Ladder", "C:\\NoxPost\\src\\client\\Gui\\ServOpts\\guiserv.c")
		nox_swprintf((*libc.WChar)(memmap.PtrOff(6112660, 1045700)), libc.CWString("%s %s"), v2, v6)
	} else {
		v3 = strMan.GetStringInFile("WindowDir:Empty", "C:\\NoxPost\\src\\client\\Gui\\ServOpts\\guiserv.c")
		nox_swprintf((*libc.WChar)(memmap.PtrOff(6112660, 1045700)), v3)
	}
	return (*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v0)))))).Func94(asWindowEvent(0x4001, int32(uintptr(memmap.PtrOff(6112660, 1045700))), -1))
}
func sub_4580E0(a1 int32) int32 {
	var (
		v1  *libc.WChar
		v2  *libc.WChar
		v3  int16
		v4  *libc.WChar
		v5  *libc.WChar
		v6  *uint32
		v7  *uint32
		v8  *uint32
		v10 int32
	)
	v1 = strMan.GetStringInFile("SettingsMsg", "C:\\NoxPost\\src\\client\\Gui\\ServOpts\\guiserv.c")
	nox_swprintf((*libc.WChar)(memmap.PtrOff(6112660, 1045968)), v1, a1)
	if noxflags.HasGame(noxflags.GameModeChat) {
		v2 = strMan.GetStringInFile("GameType", "C:\\NoxPost\\src\\client\\Gui\\ServOpts\\guiserv.c")
		nox_swprintf((*libc.WChar)(memmap.PtrOff(6112660, 1046096)), v2)
	} else {
		v3 = int16(uint16(nox_common_gameFlags_getVal_40A5B0()))
		v10 = int32(uintptr(unsafe.Pointer(nox_xxx_guiServerOptionsGetGametypeName_4573C0(v3))))
		v4 = strMan.GetStringInFile("GameTypeIs", (*byte)(memmap.PtrOff(0x587000, 131072)))
		nox_swprintf((*libc.WChar)(memmap.PtrOff(6112660, 1046096)), v4, v10)
	}
	if noxflags.HasGame(noxflags.GameHost) {
		v5 = strMan.GetStringInFile("GoMessage", "C:\\NoxPost\\src\\client\\Gui\\ServOpts\\guiserv.c")
	} else {
		v5 = strMan.GetStringInFile("OptsMessage", "C:\\NoxPost\\src\\client\\Gui\\ServOpts\\guiserv.c")
	}
	nox_swprintf((*libc.WChar)(memmap.PtrOff(6112660, 1046224)), v5)
	v6 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1046492)))).ChildByID(0x2789)))
	(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v6)))))).Func94(asWindowEvent(0x4001, int32(uintptr(memmap.PtrOff(6112660, 1045968))), -1))
	v7 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1046492)))).ChildByID(0x2786)))
	(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v7)))))).Func94(asWindowEvent(0x4001, int32(uintptr(memmap.PtrOff(6112660, 1046096))), -1))
	v8 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1046492)))).ChildByID(0x2785)))
	return (*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v8)))))).Func94(asWindowEvent(0x4001, int32(uintptr(memmap.PtrOff(6112660, 1046224))), -1))
}
func nox_client_guiserv_updateMapList_458230(mode int32, current *byte, a3 bool) {
	var (
		v19 [58]byte
		v20 [58]byte
		v21 [100]libc.WChar
	)
	(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1046496))))).Func94(asWindowEvent(0x400F, 0, 0))
	var v18 int32 = -1
	*memmap.PtrUint32(6112660, 1046552) = uint32(mode)
	var v17 int32 = 0
	for it := (*nox_map_list_item)(nox_common_maplist_first_4D09B0()); it != nil; it = nox_common_maplist_next_4D09C0(it) {
		if it.field_6 == 0 {
			continue
		}
		if (sub_4CFFC0(int32(uintptr(unsafe.Pointer(it)))) & mode) == 0 {
			continue
		}
		libc.StrCpy(&v19[0], &it.name[0])
		libc.MemCpy(unsafe.Pointer(&v20[0]), unsafe.Pointer(&v19[0]), 56)
		*(*uint16)(unsafe.Pointer(&v20[56])) = *(*uint16)(unsafe.Pointer(&v19[56]))
		sub_57A1E0((*int32)(unsafe.Pointer(&v19[0])), nil, nil, 1, int16(mode))
		sub_57A1E0((*int32)(unsafe.Pointer(&v20[0])), CString("user.rul"), nil, 3, int16(mode))
		var v6 int32 = -1
		for i := int32(0); i < 20; i += 4 {
			if *(*uint32)(unsafe.Pointer(&v19[i+24])) != *(*uint32)(unsafe.Pointer(&v20[i+24])) {
				v6 = 6
			}
		}
		if v6 == -1 {
			for j := int32(0); j < 4; j++ {
				if v19[j+44] != v20[j+44] {
					v6 = 6
				}
			}
			if v6 == -1 && *(*uint32)(unsafe.Pointer(&v19[48])) != *(*uint32)(unsafe.Pointer(&v20[48])) {
				v6 = 6
			}
		}
		var v16 int32 = int32(it.field_8_1)
		var v15 int32 = int32(it.field_8_0)
		var v9 *libc.WChar = strMan.GetStringInFile("RecPlayers", "C:\\NoxPost\\src\\client\\Gui\\ServOpts\\guiserv.c")
		nox_swprintf(&v21[0], v9, &it.name[0], v15, v16)
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1046496))))).Func94(asWindowEvent(0x400D, int32(uintptr(unsafe.Pointer(&v21[0]))), v6))
		if libc.StrCaseCmp(current, &it.name[0]) == 0 {
			v18 = v17
			(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1046496))))).Func94(asWindowEvent(0x4013, v17, 0))
			(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1046496))))).Func94(asWindowEvent(0x401C, v17, 0))
		}
		v17++
	}
	if v18 < 0 {
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1046496))))).Func94(asWindowEvent(0x4013, 0, 0))
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1046496))))).Func94(asWindowEvent(0x401C, 0, 0))
	}
	if !a3 {
		return
	}
	var v11 *byte = sub_4165B0()
	var v12 int32 = (*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1046496))))).Func94(asWindowEvent(0x4014, 0, 0))
	if v12 < 0 {
		*v11 = 0
	} else {
		var v13 int32 = (*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1046496))))).Func94(asWindowEvent(0x4016, v12, 0))
		nox_sprintf((*byte)(unsafe.Pointer(&v21[0])), CString("%S"), v13)
		var v14 *byte = libc.StrTok((*byte)(unsafe.Pointer(&v21[0])), CString("\t"))
		if v14 == nil {
			*v11 = 0
		} else {
			libc.StrCpy(v11, v14)
		}
	}
	sub_57A1E0((*int32)(unsafe.Pointer(v11)), CString("user.rul"), nil, 7, int16(mode))
	sub_459880(int32(uintptr(unsafe.Pointer(v11))))
}
func nox_xxx_guiServerOptionsProcPre_4585D0(a1 int32, a2 uint32, a3 int32, a4 int32) int32 {
	var (
		v4          *uint32
		v5          *byte
		result      int32
		v7          int32
		v8          int32
		v9          int32
		v10         *int32
		v11         int32
		v12         *uint32
		v13         *uint32
		v14         *uint32
		v15         *uint32
		v16         int32
		v17         *uint32
		v18         *byte
		v19         int32
		v20         int16
		v21         bool
		v22         *libc.WChar
		v23         *libc.WChar
		v24         *libc.WChar
		v25         *byte
		v26         *uint32
		v27         int32
		v28         int32
		v29         *byte
		v30         *uint32
		v31         int32
		v32         *libc.WChar
		v33         int32
		v34         uint32
		v35         int32
		v36         *byte
		v37         *byte
		v38         *libc.WChar
		v39         *libc.WChar
		WideCharStr [4]libc.WChar
		v41         [100]byte
		v42         [100]byte
	)
	if a2 > 0x4007 {
		if a2 == 16400 {
			v28 = (*nox_window)(unsafe.Pointer(uintptr(a3))).ID() - 0x2782
			if v28 != 0 {
				if v28 == 6 {
					v29 = sub_4165B0()
					v30 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1046492)))).ChildByID(0x2787)))
					v31 = (*nox_window)(unsafe.Pointer(uintptr(a3))).Func94(asWindowEvent(0x4014, 0, 0))
					if v31 >= 0 && v31 < int32(*(*int16)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a3 + 32))) + 44)))) {
						v32 = (*libc.WChar)(unsafe.Pointer(uintptr((*nox_window)(unsafe.Pointer(uintptr(a3))).Func94(asWindowEvent(0x4016, a4, 0)))))
						(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v30)))))).Func94(asWindowEvent(0x4001, int32(uintptr(unsafe.Pointer(v32))), -1))
						v33 = int32(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v29))), unsafe.Sizeof(uint16(0))*26)))) & 6128
						*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v29))), unsafe.Sizeof(uint16(0))*26))) &= 0xE80F
						*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v29))), unsafe.Sizeof(uint16(0))*26))) |= uint16(int16(sub_459650(v32)))
						if v33 != (int32(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v29))), unsafe.Sizeof(uint16(0))*26)))) & 6128) {
							nox_client_guiserv_updateMapList_458230(int32(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v29))), unsafe.Sizeof(uint16(0))*26)))), (*byte)(memmap.PtrOff(6112660, 1046556)), true)
						}
						(*nox_window)(unsafe.Pointer(uintptr(a3))).Hide()
						nox_xxx_wndClearCaptureMain((*nox_window)(unsafe.Pointer(uintptr(a3))))
						*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v29))), unsafe.Sizeof(uint16(0))*27))) = uint16(nox_xxx_servGamedataGet_40A020(int16(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v29))), unsafe.Sizeof(uint16(0))*26))))))
						*(*byte)(unsafe.Add(unsafe.Pointer(v29), 56)) = byte(sub_40A180(int16(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v29))), unsafe.Sizeof(uint16(0))*26))))))
						sub_457460(int32(uintptr(unsafe.Pointer(v29))))
						sub_459D50(1)
						if *(*byte)(unsafe.Add(unsafe.Pointer(v29), 53))&16 != 0 {
							nox_xxx_spellEnableAll_424BD0()
							sub_4537F0()
							sub_459700()
						}
						*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v34))), unsafe.Sizeof(uint16(0))*0)) = *((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v29))), unsafe.Sizeof(uint16(0))*26)))
						sub_457F30(int32((v34 >> 12) & 1))
						return 1
					}
				}
			} else if (*nox_window)(unsafe.Pointer(uintptr(a3))).Func94(asWindowEvent(0x4014, 0, 0)) >= 0 {
				v35 = (*nox_window)(unsafe.Pointer(uintptr(a3))).Func94(asWindowEvent(0x4016, a4, 0))
				nox_sprintf(&v42[0], CString("%S"), v35)
				v36 = libc.StrTok(&v42[0], CString("\t"))
				v37 = sub_4165B0()
				if v36 == nil {
					v36 = CString("")
				}
				libc.StrCpy(v37, v36)
				sub_57A1E0((*int32)(unsafe.Pointer(v37)), CString("user.rul"), nil, 7, int16(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v37))), unsafe.Sizeof(uint16(0))*26)))))
				sub_459880(int32(uintptr(unsafe.Pointer(v37))))
				sub_459D50(1)
			}
		}
		return 1
	}
	if a2 == 0x4007 {
		v9 = (*nox_window)(unsafe.Pointer(uintptr(a3))).ID()
		clientPlaySoundSpecial(766, 100)
		switch v9 {
		case 0x2787:
			v26 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1046492)))).ChildByID(10120)))
			(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v26)))))).Show()
			nox_xxx_wndSetCaptureMain((*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v26)))))))
			nox_xxx_wndShowModalMB((*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v26)))))))
			nox_xxx_windowServerOptionsFillGametypeList_4596A0()
			return 1
		case 0x278A:
			v27 = bool2int((int32(*(*uint8)(unsafe.Pointer(uintptr(a3 + 36)))) & 4) != 0)
			if !noxflags.HasGame(noxflags.GameFlag15 | noxflags.GameFlag16) {
				nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1046500))))), v27)
				nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1046504))))), v27)
			}
			sub_459D50(1)
			return 1
		case 0x279D:
			sub_459700()
			return 1
		case 0x27A1:
			v18 = nox_xxx_cliGamedataGet_416590(1)
			sub_459AA0(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v18))))))
			sub_4165F0(1, 0)
			v19 = int32(nox_xxx_getTeamCounter_417DD0())
			if noxflags.HasGame(noxflags.GameModeChat) && (func() int32 {
				v20 = int16(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v18))), unsafe.Sizeof(uint16(0))*26))))
				return int32(v20) & 96
			}()) != 0 {
				if int32(v20) < 0 {
					v21 = v19 <= 2
				} else {
					v21 = v19 <= 2
					if v19 < 2 {
						v22 = strMan.GetStringInFile("NeedTeams", "C:\\NoxPost\\src\\client\\Gui\\ServOpts\\guiserv.c")
						NewDialogWindow((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1046492))))), (*libc.WChar)(unsafe.Pointer(uintptr(int32(uintptr(memmap.PtrOff(6112660, 1046560)))))), (*libc.WChar)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v22)))))), 56, unsafe.Pointer(cgoFuncAddr(sub_459150)), nil)
						sub_44A360(1)
						return 1
					}
				}
				if !v21 {
					v38 = strMan.GetStringInFile("TooManyTeams", "C:\\NoxPost\\src\\client\\Gui\\ServOpts\\guiserv.c")
					v23 = strMan.GetStringInFile("Notice", "C:\\NoxPost\\src\\client\\Gui\\ServOpts\\guiserv.c")
					NewDialogWindow((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1046492))))), (*libc.WChar)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v23)))))), (*libc.WChar)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v38)))))), 33, nil, nil)
					sub_44A360(1)
					return 1
				}
			} else if noxflags.HasGame(noxflags.GameModeChat) && noxflags.HasGame(noxflags.GameModeKOTR) && nox_xxx_CheckGameplayFlags_417DA0(4) && v19 > 2 {
				v39 = strMan.GetStringInFile("TooManyTeams", "C:\\NoxPost\\src\\client\\Gui\\ServOpts\\guiserv.c")
				v24 = strMan.GetStringInFile("Notice", "C:\\NoxPost\\src\\client\\Gui\\ServOpts\\guiserv.c")
				NewDialogWindow((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1046492))))), (*libc.WChar)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v24)))))), (*libc.WChar)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v39)))))), 33, nil, nil)
				sub_44A360(1)
				return 1
			}
			sub_459150()
			result = 1
		case 0x27A2:
			v25 = sub_4165B0()
			if noxflags.HasGame(noxflags.GameModeChat) {
				if int32(*(*byte)(unsafe.Add(unsafe.Pointer(v25), 53))) < 0 {
					nox_server_teamsZzz_419030(1)
				}
				*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v25))), unsafe.Sizeof(uint16(0))*26))) &= 0x3FFF
			}
			goto LABEL_58
		case 0x27A5:
		LABEL_58:
			nox_xxx_guiServerOptionsHide_4597E0(0)
			return 1
		case 0x27A8:
			v10 = (*int32)(unsafe.Pointer(nox_xxx_cliGamedataGet_416590(1)))
			nox_xxx_loadAdvancedWnd_4BDC10(v10)
			return 1
		case 0x27AF:
			v11 = int32(*(*uint32)(unsafe.Pointer(uintptr(a3 + 396))))
			sub_46B120((*nox_window)(unsafe.Pointer(uintptr(a3))), nil)
			sub_46B120((*nox_window)(unsafe.Pointer(uintptr(a3))), (*nox_window)(unsafe.Pointer(uintptr(v11))))
			v12 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1046492)))).ChildByID(0x27D4)))
			if noxflags.HasGame(noxflags.GameHost) {
				(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v12)))))).Show()
				sub_46ACE0(*(**uint32)(unsafe.Pointer(&dword_5d4594_1046492)), 0x27B1, 0x27B3, 0)
				sub_46AD20(*(**uint32)(unsafe.Pointer(&dword_5d4594_1046492)), 0x27B1, 0x27B3, 1)
				sub_4593B0(0)
				v13 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1046492)))).ChildByID(0x27B3)))
				*(*uint32)(unsafe.Add(unsafe.Pointer(v13), unsafe.Sizeof(uint32(0))*9)) |= 4
				v14 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1046492)))).ChildByID(0x27B1)))
				*(*uint32)(unsafe.Add(unsafe.Pointer(v14), unsafe.Sizeof(uint32(0))*9)) &= 0xFFFFFFFB
				v15 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1046492)))).ChildByID(0x27B2)))
				*(*uint32)(unsafe.Add(unsafe.Pointer(v15), unsafe.Sizeof(uint32(0))*9)) &= 0xFFFFFFFB
				sub_46ACE0(*(**uint32)(unsafe.Pointer(&dword_5d4594_1046492)), 0x279D, 0x279D, 1)
			} else {
				(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v12)))))).Hide()
				dword_5d4594_1046532 = uint32(nox_xxx_guiServerPlayersLoad_456270(*(*int32)(unsafe.Pointer(&dword_5d4594_1046492))))
				sub_46ACE0(*(**uint32)(unsafe.Pointer(&dword_5d4594_1046492)), 0x279D, 0x279D, 1)
			}
			(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1046524))))).Hide()
			return 1
		case 0x27B0:
			v16 = int32(*(*uint32)(unsafe.Pointer(uintptr(a3 + 396))))
			sub_46B120((*nox_window)(unsafe.Pointer(uintptr(a3))), nil)
			sub_46B120((*nox_window)(unsafe.Pointer(uintptr(a3))), (*nox_window)(unsafe.Pointer(uintptr(v16))))
			if dword_5d4594_1046532 != 0 {
				sub_456D60(1)
				dword_5d4594_1046532 = 0
			}
			if dword_5d4594_1046528 != 0 {
				sub_4557D0(1)
				dword_5d4594_1046528 = 0
			}
			if dword_5d4594_1046536 != 0 {
				sub_4AD820()
				dword_5d4594_1046540 = 0
			}
			sub_46ACE0(*(**uint32)(unsafe.Pointer(&dword_5d4594_1046492)), 0x27B1, 0x27B3, 1)
			if noxflags.HasGame(noxflags.GameHost) {
				sub_46ACE0(*(**uint32)(unsafe.Pointer(&dword_5d4594_1046492)), 0x279D, 0x279D, 0)
			}
			v17 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1046492)))).ChildByID(0x27D4)))
			(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v17)))))).Hide()
			(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1046524))))).Show()
			return 1
		case 0x27B1:
			sub_4593B0(1)
			return 1
		case 0x27B2:
			sub_4593B0(2)
			return 1
		case 0x27B3:
			sub_4593B0(0)
			return 1
		case 10330:
			if nox_xxx_CheckGameplayFlags_417DA0(4) {
				nox_xxx_toggleAllTeamFlags_418690(0)
				nox_server_teamsZzz_419030(1)
				sub_46AD20(*(**uint32)(unsafe.Pointer(&dword_5d4594_1046508)), 0x285B, 0x285D, 0)
				result = 1
			} else {
				nox_xxx_wndGuiTeamCreate_4185B0()
				if nox_xxx_CheckGameplayFlags_417DA0(2) {
					sub_4181F0(0)
				} else {
					nox_xxx_toggleAllTeamFlags_418690(1)
				}
				sub_46AD20(*(**uint32)(unsafe.Pointer(&dword_5d4594_1046508)), 0x285B, 0x285D, 1)
				result = 1
			}
			return result
		case 0x285B:
			if nox_xxx_CheckGameplayFlags_417DA0(2) {
				nox_xxx_UnsetGameplayFlags_417D70(2)
				nox_xxx_toggleAllTeamFlags_418690(1)
			} else {
				nox_xxx_toggleAllTeamFlags_418690(0)
				sub_418390()
			}
			return 1
		case 0x285C:
			sub_4181F0(1)
			return 1
		case 0x285D:
			if nox_xxx_CheckGameplayFlags_417DA0(1) {
				nox_xxx_UnsetGameplayFlags_417D70(1)
			} else {
				nox_xxx_SetGameplayFlag_417D50(1)
			}
			return 1
		default:
			return 1
		}
		return result
	}
	if a2 == 23 {
		return 0
	}
	if a2 != 0x4003 {
		return 1
	}
	v4 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1046492)))).ChildByID(a4)))
	v5 = sub_4165B0()
	if v4 == nil {
		return 0
	}
	if int32(uint16(int16(a3))) == 1 {
		if a4 == 0x2775 {
			*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*26)) = nox_color_white_2523948
			return 1
		}
		return 1
	}
	v7 = (*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v4)))))).Func94(asWindowEvent(0x401D, 0, 0))
	nox_sprintf(&v41[0], CString("%S"), v7)
	if v41[0] == 0 {
		return 1
	}
	v8 = int32(libc.Atoi(libc.GoString(&v41[0])))
	if v8 < 0 {
		v8 = 0
	}
	if a4 == 0x2775 {
		*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*26)) = nox_color_rgb_4344A0(230, 165, 65)
		libc.StrCpy((*byte)(unsafe.Add(unsafe.Pointer(v5), 9)), &v41[0])
		nox_xxx_gameSetServername_40A440(&v41[0])
		return 1
	}
	if a4 == 0x2796 {
		*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v5))), unsafe.Sizeof(uint16(0))*27))) = uint16(int16(v8))
		sub_459D50(1)
		return 1
	}
	if a4 != 0x2797 {
		return 1
	}
	if v8 > math.MaxUint8 {
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v8))), 0)) = math.MaxUint8
		compat_itow(math.MaxUint8, &WideCharStr[0], 10)
		(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v4)))))).Func94(asWindowEvent(0x401E, int32(uintptr(unsafe.Pointer(&WideCharStr[0]))), -1))
	}
	*(*byte)(unsafe.Add(unsafe.Pointer(v5), 56)) = byte(int8(v8))
	sub_459D50(1)
	return 1
}
func sub_459880(a1 int32) int32 {
	var (
		v1 int16
		v2 *libc.WChar
		v4 *uint32
		v5 int32
		v6 uint32
	)
	v1 = int16(*(*uint16)(unsafe.Pointer(uintptr(a1 + 52))))
	if int32(v1)&32 != 0 {
		if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 57)))) == 0 && noxflags.HasGame(noxflags.GameHost) {
			nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1046516))))), 1)
			nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1046520))))), 1)
		}
		v2 = strMan.GetStringInFile("Servopts.wnd:CaptureLimit", "C:\\NoxPost\\src\\client\\Gui\\ServOpts\\guiserv.c")
	} else {
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v1))), 0)) = *(*uint8)(unsafe.Pointer(uintptr(a1 + 57)))
		if int32(v1)&1024 != 0 {
			if int32(*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v1))), 0))) == 0 && noxflags.HasGame(noxflags.GameHost) {
				nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1046516))))), 1)
				nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1046520))))), 1)
			}
			v2 = strMan.GetStringInFile("Servopts.wnd:DeathLimit", "C:\\NoxPost\\src\\client\\Gui\\ServOpts\\guiserv.c")
		} else {
			if int32(*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v1))), 0))) == 0 && noxflags.HasGame(noxflags.GameHost) && !noxflags.HasGame(noxflags.GameFlag15|noxflags.GameFlag16) {
				nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1046516))))), 1)
				nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1046520))))), 1)
			}
			v2 = strMan.GetStringInFile("Servopts.wnd:KillLimit", "C:\\NoxPost\\src\\client\\Gui\\ServOpts\\guiserv.c")
		}
	}
	nox_wcscpy((*libc.WChar)(unsafe.Pointer(uintptr(dword_5d4594_1046516+108))), v2)
	sub_4580E0(a1)
	sub_459A40((*byte)(unsafe.Pointer(uintptr(a1 + 9))))
	if noxflags.HasGame(noxflags.GameHost) && !noxflags.HasGame(noxflags.GameFlag15|noxflags.GameFlag16) {
		nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1046500))))), bool2int(int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 57)))) == 0))
		nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1046504))))), bool2int(int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 57)))) == 0))
	}
	v4 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1046492)))).ChildByID(0x278A)))
	v5 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*9)))
	if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 57)))) != 0 {
		v6 = uint32(v5 | 4)
	} else {
		v6 = uint32(v5) & 0xFFFFFFFB
	}
	*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*9)) = v6
	sub_453F70(unsafe.Pointer(uintptr(a1 + 24)))
	sub_4535E0((*int32)(unsafe.Pointer(uintptr(a1 + 44))))
	return sub_4535F0(int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 48)))))
}
func sub_459CD0() int32 {
	var (
		result int32
		v1     uint8
		v2     *libc.WChar
		v3     *uint32
		v4     int32
	)
	result = int32(dword_5d4594_1046492)
	if dword_5d4594_1046492 != 0 {
		if sub_40A740() != 0 || noxflags.HasGame(noxflags.GameFlag16) {
			v1 = uint8(sub_417DE0())
		} else {
			v1 = nox_xxx_getTeamCounter_417DD0()
		}
		v4 = int32(v1)
		v2 = strMan.GetStringInFile("NumTeamsMsg", "C:\\NoxPost\\src\\client\\Gui\\ServOpts\\guiserv.c")
		nox_swprintf((*libc.WChar)(memmap.PtrOff(6112660, 1046364)), v2, v4)
		v3 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1046492)))).ChildByID(10110)))
		result = (*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v3)))))).Func94(asWindowEvent(0x4001, int32(uintptr(memmap.PtrOff(6112660, 1046364))), -1))
	}
	return result
}
