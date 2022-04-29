package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	noxflags "github.com/noxworld-dev/opennox/v1/common/flags"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"image"
	"math"
	"unsafe"
)

var dword_5d4594_815104 int32 = 0
var nox_wnd_xxx_815040 *nox_gui_animation = nil

func nox_game_showGameSel_4379F0() int32 {
	var (
		v2  *uint32
		v3  *uint32
		v4  *uint32
		v5  *libc.WChar
		v6  *uint32
		v7  *libc.WChar
		v8  *uint32
		v9  *uint32
		v10 *uint32
		v11 *uint32
		v12 *uint32
		v13 *uint32
		v14 *uint32
		v15 *uint32
		v16 int32
		v17 *libc.WChar
		v18 *libc.WChar
		v19 *uint32
		v20 *uint32
		v21 *uint32
		v22 *uint32
		v23 *uint32
		v24 *uint32
		v25 *uint32
		v26 *uint32
		v27 *libc.WChar
		v28 *libc.WChar
	)
	nox_xxx_setQuest_4D6F60(0)
	sub_4D6F80(0)
	if !noxflags.HasGame(noxflags.GameFlag26) && sub_4D6F30() == 0 {
		nox_game_createOrJoin_815048 = 0
		dword_587000_87412 = math.MaxUint32
	}
	*memmap.PtrUint32(6112660, 815076) = 0
	*memmap.PtrUint32(6112660, 815080) = 0
	*memmap.PtrUint32(6112660, 815084) = 0
	gameAddStateCode(10000)
	sub_4A24C0(1)
	if *(*int32)(unsafe.Pointer(&dword_587000_87404)) != 1 || (func() bool {
		sub_4A1BE0(0)
		return *(*int32)(unsafe.Pointer(&dword_587000_87412)) != -1
	}()) {
		sub_49FDB0(0)
	}
	if nox_wol_wnd_world_814980 != nil {
		dword_5d4594_815044 = 0
		nox_wnd_xxx_815040.state = nox_gui_anim_state(NOX_GUI_ANIM_IN)
		nox_wnd_xxx_815040.fnc_done_out = sub_438330
		sub_43BE40(3)
		clientPlaySoundSpecial(922, 100)
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&nox_wol_wnd_world_814980))))).Show()
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815000))))).Show()
		nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_814984))))), 1)
		if nox_game_createOrJoin_815048 != 0 {
			sub_4375C0(0)
		}
		sub_46ACE0(*(**uint32)(unsafe.Pointer(&nox_wol_wnd_world_814980)), 0x273F, 0x2743, bool2int(dword_587000_87408 == 0))
		if sub_4D6FA0() == 2 {
			v25 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&nox_wol_wnd_world_814980)))).ChildByID(10010)))
			(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&nox_wol_wnd_world_814980))))).Func94(asWindowEvent(0x4007, int32(uintptr(unsafe.Pointer(v25))), 0))
			v26 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&nox_wol_wnd_world_814980)))).ChildByID(10010)))
			(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&nox_wol_wnd_world_814980))))).Func94(asWindowEvent(0x4007, int32(uintptr(unsafe.Pointer(v26))), 0))
		}
		return 1
	}
	nox_wol_wnd_world_814980 = newWindowFromFile("noxworld.wnd", unsafe.Pointer(cgoFuncAddr(nox_xxx_windowMultiplayerSub_439E70)))
	if nox_wol_wnd_world_814980 == nil {
		return 0
	}
	sub_49FF20()
	nox_xxx_wndShowModalMB((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&nox_wol_wnd_world_814980))))))
	(*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&nox_wol_wnd_world_814980)))).SetAllFuncs(func(arg1 int32, arg2 int32, arg3 int32, arg4 int32) int32 {
		return sub_439D00((*int32)(unsafe.Pointer(uintptr(arg1))), arg2, uint32(arg3), arg4)
	}, nil, nil)
	nox_wnd_xxx_815040 = nox_gui_makeAnimation((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&nox_wol_wnd_world_814980)))), 0, 0, 0, -480, 0, 20, 0, -40)
	if nox_wnd_xxx_815040 == nil {
		return 0
	}
	nox_wnd_xxx_815040.field_0 = 10000
	nox_wnd_xxx_815040.field_12 = sub_438370
	nox_wnd_xxx_815040.fnc_done_out = sub_438330
	dword_5d4594_814984 = uint32(uintptr(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&nox_wol_wnd_world_814980)))).ChildByID(10020))))
	dword_5d4594_814988 = uint32(uintptr(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&nox_wol_wnd_world_814980)))).ChildByID(0x2725))))
	(*(*int32)(unsafe.Pointer(&dword_5d4594_814984))).setFunc93(func(arg1 int32, arg2 int32, arg3 int32, arg4 int32) int32 {
		return sub_439D00((*int32)(unsafe.Pointer(uintptr(arg1))), arg2, uint32(arg3), arg4)
	})
	(*(*int32)(unsafe.Pointer(&dword_5d4594_814988))).setFunc93(func(arg1 int32, arg2 int32, arg3 int32, arg4 int32) int32 {
		return sub_439D00((*int32)(unsafe.Pointer(uintptr(arg1))), arg2, uint32(arg3), arg4)
	})
	(*(*int32)(unsafe.Pointer(&dword_5d4594_814988))).setFunc94(func(arg1 int32, arg2 int32, arg3 int32, arg4 int32) int32 {
		return nox_xxx_windowMultiplayerSub_439E70(arg1, uint32(arg2), (*int32)(unsafe.Pointer(uintptr(arg3))), arg4)
	})
	dword_5d4594_814996 = uint32(uintptr(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&nox_wol_wnd_world_814980)))).ChildByID(0x271B))))
	nox_wol_wnd_gameList_815012 = uint32(uintptr(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&nox_wol_wnd_world_814980)))).ChildByID(0x2735))))
	dword_5d4594_815004 = (*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&nox_wol_wnd_world_814980)))).ChildByID(0x2732)
	dword_5d4594_815000 = uint32(uintptr(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&nox_wol_wnd_world_814980)))).ChildByID(0x2731))))
	dword_5d4594_814992 = uint32(uintptr(unsafe.Pointer(sub_489B80(*(*int32)(unsafe.Pointer(&nox_wol_wnd_world_814980))))))
	*memmap.PtrUint32(6112660, 815008) = uint32(uintptr(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&nox_wol_wnd_world_814980)))).ChildByID(0x2711))))
	dword_5d4594_815016 = uint32(uintptr(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&nox_wol_wnd_world_814980)))).ChildByID(0x2736))))
	dword_5d4594_815020 = uint32(uintptr(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&nox_wol_wnd_world_814980)))).ChildByID(0x2737))))
	dword_5d4594_815024 = uint32(uintptr(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&nox_wol_wnd_world_814980)))).ChildByID(10040))))
	dword_5d4594_815028 = uint32(uintptr(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&nox_wol_wnd_world_814980)))).ChildByID(0x2739))))
	dword_5d4594_815032 = uint32(uintptr(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&nox_wol_wnd_world_814980)))).ChildByID(0x273A))))
	v2 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&nox_wol_wnd_world_814980)))).ChildByID(0x2712)))
	if noxflags.HasGame(noxflags.GameFlag25) {
		nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v2)))))), 0)
	}
	v3 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&nox_wol_wnd_world_814980)))).ChildByID(0x2713)))
	*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*9)) &= 0xFFFFFFFB
	v4 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&nox_wol_wnd_world_814980)))).ChildByID(0x273E)))
	int32(uintptr(unsafe.Pointer(v4))).setDraw(sub_438C80)
	(*(*int32)(unsafe.Pointer(&nox_wol_wnd_gameList_815012))).setFunc94(func(arg1 int32, arg2 int32, arg3 int32, arg4 int32) int32 {
		return nox_xxx_windowMultiplayerSub_439E70(arg1, uint32(arg2), (*int32)(unsafe.Pointer(uintptr(arg3))), arg4)
	})
	(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&nox_wol_wnd_gameList_815012))))).Hide()
	(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815000))))).Hide()
	nox_xxx_wnd_46B280(*memmap.PtrInt32(6112660, 815008), *(*int32)(unsafe.Pointer(&nox_wol_wnd_world_814980)))
	if dword_587000_87404 != 0 {
		gameSetCliDrawFunc(unsafe.Pointer(cgoFuncAddr(sub_41E210)))
		v6 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&nox_wol_wnd_world_814980)))).ChildByID(0x2717)))
		(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v6)))))).Show()
		dword_587000_87412 = math.MaxUint32
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_814984))))).Hide()
		v7 = strMan.GetStringInFile("ChooseArea", "C:\\NoxPost\\src\\client\\shell\\noxworld.c")
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_814996))))).Func94(asWindowEvent(0x4001, int32(uintptr(unsafe.Pointer(v7))), 0))
	} else {
		dword_587000_87412 = 0
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_814988))))).Hide()
		v5 = strMan.GetStringInFile("JoinServer", "C:\\NoxPost\\src\\client\\shell\\noxworld.c")
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_814996))))).Func94(asWindowEvent(0x4001, int32(uintptr(unsafe.Pointer(v5))), 0))
		sub_46ACE0(*(**uint32)(unsafe.Pointer(&dword_5d4594_814984)), 10620, 0x2987, 1)
		sub_49FDB0(0)
	}
	*memmap.PtrUint32(6112660, 814556) = uint32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg("NWGameIconLargeGreen"))))
	*memmap.PtrUint32(6112660, 814560) = uint32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg("NWGameIconLargeGreenLit"))))
	*memmap.PtrUint32(6112660, 814564) = uint32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg("NWGameIconSmallGreen"))))
	*memmap.PtrUint32(6112660, 814568) = uint32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg("NWGameIconSmallGreenLit"))))
	*memmap.PtrUint32(6112660, 814572) = uint32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg("NWGameIconLargeYellow"))))
	*memmap.PtrUint32(6112660, 814576) = uint32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg("NWGameIconLargeYellowLit"))))
	*memmap.PtrUint32(6112660, 814580) = uint32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg("NWGameIconSmallYellow"))))
	*memmap.PtrUint32(6112660, 814584) = uint32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg("NWGameIconSmallYellowLit"))))
	*memmap.PtrUint32(6112660, 814588) = uint32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg("NWGameIconLargeRed"))))
	*memmap.PtrUint32(6112660, 814592) = uint32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg("NWGameIconLargeRedLit"))))
	*memmap.PtrUint32(6112660, 814596) = uint32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg("NWGameIconSmallRed"))))
	*memmap.PtrUint32(6112660, 814600) = uint32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg("NWGameIconSmallRedLit"))))
	*memmap.PtrUint32(6112660, 814900) = uint32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg("NWMapULLg"))))
	*memmap.PtrUint32(6112660, 814904) = uint32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg("NWMapURLg"))))
	*memmap.PtrUint32(6112660, 814908) = uint32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg("NWMapLLLg"))))
	*memmap.PtrUint32(6112660, 814912) = uint32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg("NWMapLRLg"))))
	v8 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&nox_wol_wnd_gameList_815012)))).ChildByID(0x2745)))
	v9 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&nox_wol_wnd_gameList_815012)))).ChildByID(0x273B)))
	v10 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&nox_wol_wnd_gameList_815012)))).ChildByID(0x273C)))
	v11 = *(**uint32)(unsafe.Pointer(uintptr(nox_wol_wnd_gameList_815012 + 32)))
	nox_xxx_wnd_46B280(int32(uintptr(unsafe.Pointer(v8))), *(*int32)(unsafe.Pointer(&nox_wol_wnd_gameList_815012)))
	nox_xxx_wnd_46B280(int32(uintptr(unsafe.Pointer(v9))), *(*int32)(unsafe.Pointer(&nox_wol_wnd_gameList_815012)))
	nox_xxx_wnd_46B280(int32(uintptr(unsafe.Pointer(v10))), *(*int32)(unsafe.Pointer(&nox_wol_wnd_gameList_815012)))
	*(*uint32)(unsafe.Add(unsafe.Pointer(v11), unsafe.Sizeof(uint32(0))*9)) = uint32(uintptr(unsafe.Pointer(v8)))
	*(*uint32)(unsafe.Add(unsafe.Pointer(v11), unsafe.Sizeof(uint32(0))*7)) = uint32(uintptr(unsafe.Pointer(v9)))
	*(*uint32)(unsafe.Add(unsafe.Pointer(v11), unsafe.Sizeof(uint32(0))*8)) = uint32(uintptr(unsafe.Pointer(v10)))
	*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v8), unsafe.Sizeof(uint32(0))*100)) + 8))) = 16
	*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v8), unsafe.Sizeof(uint32(0))*100)) + 12))) = 12
	nox_xxx_wndSetOffsetMB_46AE40(int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v8), unsafe.Sizeof(uint32(0))*100))), 0, -15)
	v12 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_815000)))).ChildByID(0x2730)))
	v13 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_815000)))).ChildByID(0x2733)))
	v14 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_815000)))).ChildByID(0x2734)))
	v15 = *(**uint32)(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer(dword_5d4594_815004))) + 32)))
	nox_xxx_wnd_46B280(int32(uintptr(unsafe.Pointer(v12))), *(*int32)(unsafe.Pointer(&dword_5d4594_815004)))
	nox_xxx_wnd_46B280(int32(uintptr(unsafe.Pointer(v13))), *(*int32)(unsafe.Pointer(&dword_5d4594_815004)))
	nox_xxx_wnd_46B280(int32(uintptr(unsafe.Pointer(v14))), *(*int32)(unsafe.Pointer(&dword_5d4594_815004)))
	*(*uint32)(unsafe.Add(unsafe.Pointer(v15), unsafe.Sizeof(uint32(0))*9)) = uint32(uintptr(unsafe.Pointer(v12)))
	*(*uint32)(unsafe.Add(unsafe.Pointer(v15), unsafe.Sizeof(uint32(0))*7)) = uint32(uintptr(unsafe.Pointer(v13)))
	*(*uint32)(unsafe.Add(unsafe.Pointer(v15), unsafe.Sizeof(uint32(0))*8)) = uint32(uintptr(unsafe.Pointer(v14)))
	*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v12), unsafe.Sizeof(uint32(0))*100)) + 8))) = 16
	*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v12), unsafe.Sizeof(uint32(0))*100)) + 12))) = 10
	sub_438480()
	v16 = nox_xxx_servGetPort_40A430()
	sub_40A3C0(uint32(v16))
	nox_game_createOrJoin_815048 = 0
	dword_5d4594_815044 = 0
	dword_5d4594_815052 = 0
	dword_5d4594_814548 = 0
	qword_5d4594_815068 = uint64(platformTicks() + 1000)
	if dword_587000_87408 == 1 {
		sub_4383A0()
	}
	if dword_587000_87404 == 0 {
		nox_xxx_createSocketLocal_554B40(0)
	}
	nox_wol_server_result_cnt_815088 = 0
	nox_xxx_loadModifyers_4158C0()
	nox_xxx_loadLook_415D50()
	sub_430C30_set_video_max(noxMaxWidth, noxMaxHeight)
	nox_client_setCursorType_477610(0)
	if dword_5d4594_815096 != 0 {
		v27 = strMan.GetStringInFile("Kicked", "C:\\NoxPost\\src\\client\\shell\\noxworld.c")
		v17 = strMan.GetStringInFile("Notification", "C:\\NoxPost\\src\\client\\shell\\noxworld.c")
		NewDialogWindow(nil, (*libc.WChar)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v17)))))), (*libc.WChar)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v27)))))), 33, nil, nil)
		sub_44A360(1)
		dword_5d4594_815096 = 0
	} else if dword_5d4594_815100 != 0 {
		v28 = strMan.GetStringInFile("Timeout", "C:\\NoxPost\\src\\client\\shell\\noxworld.c")
		v18 = strMan.GetStringInFile("Notification", "C:\\NoxPost\\src\\client\\shell\\noxworld.c")
		NewDialogWindow(nil, (*libc.WChar)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v18)))))), (*libc.WChar)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v28)))))), 33, nil, nil)
		sub_44A360(1)
		dword_5d4594_815100 = 0
	}
	sub_43DE40(unsafe.Pointer(cgoFuncAddr(sub_438770)))
	if sub_44A4A0() != 0 {
		sub_44A4B0()
	}
	(*(*int32)(unsafe.Pointer(&dword_5d4594_814984))).setDraw(func(arg1 int32, arg2 int32) int32 {
		return sub_438E30((*uint32)(unsafe.Pointer(uintptr(arg1))), arg2)
	})
	v19 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&nox_wol_wnd_world_814980)))).ChildByID(0x2746)))
	int32(uintptr(unsafe.Pointer(v19))).setDraw(func(arg1 int32, arg2 int32) int32 {
		return sub_438E30((*uint32)(unsafe.Pointer(uintptr(arg1))), arg2)
	})
	v20 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&nox_wol_wnd_world_814980)))).ChildByID(0x2747)))
	int32(uintptr(unsafe.Pointer(v20))).setDraw(func(arg1 int32, arg2 int32) int32 {
		return sub_438E30((*uint32)(unsafe.Pointer(uintptr(arg1))), arg2)
	})
	v21 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&nox_wol_wnd_world_814980)))).ChildByID(0x2748)))
	int32(uintptr(unsafe.Pointer(v21))).setDraw(func(arg1 int32, arg2 int32) int32 {
		return sub_438E30((*uint32)(unsafe.Pointer(uintptr(arg1))), arg2)
	})
	v22 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&nox_wol_wnd_world_814980)))).ChildByID(0x2749)))
	int32(uintptr(unsafe.Pointer(v22))).setDraw(func(arg1 int32, arg2 int32) int32 {
		return sub_438E30((*uint32)(unsafe.Pointer(uintptr(arg1))), arg2)
	})
	if dword_587000_87408 == 0 {
		sub_46ACE0(*(**uint32)(unsafe.Pointer(&nox_wol_wnd_world_814980)), 0x273F, 0x2743, 1)
	}
	if sub_4D6FA0() == 1 {
		v23 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&nox_wol_wnd_world_814980)))).ChildByID(0x2712)))
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&nox_wol_wnd_world_814980))))).Func94(asWindowEvent(0x4007, int32(uintptr(unsafe.Pointer(v23))), 0))
		nox_xxx_setQuest_4D6F60(1)
		v24 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&nox_wol_wnd_world_814980)))).ChildByID(10020)))
		(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v24)))))).Func93(asWindowEvent(0x5, 0xEF0198, 0))
	}
	return 1
}
func sub_4383A0() int32 {
	var (
		v0     *libc.WChar
		result int32
	)
	if (*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_814992))))).Flags().IsHidden() == 0 {
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_814992))))).Hide()
		sub_489870()
	}
	(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&nox_wol_wnd_gameList_815012))))).Show()
	(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_814984))))).Hide()
	(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_814988))))).Hide()
	sub_46AD20(*(**uint32)(unsafe.Pointer(&nox_wol_wnd_world_814980)), 0x2716, 0x2717, 1)
	sub_46ACE0(*(**uint32)(unsafe.Pointer(&nox_wol_wnd_world_814980)), 0x273F, 0x2743, 0)
	v0 = strMan.GetStringInFile("ListJoinServer", "C:\\NoxPost\\src\\client\\shell\\noxworld.c")
	(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_814996))))).Func94(asWindowEvent(0x4001, int32(uintptr(unsafe.Pointer(v0))), 0))
	result = int32(dword_587000_87404)
	dword_587000_87408 = 1
	if dword_587000_87404 == 1 {
		dword_587000_87412 = math.MaxUint32
	}
	return result
}
func sub_438770() int32 {
	var (
		result int32
		v1     int64
		v2     *libc.WChar
		v3     *libc.WChar
		v4     *libc.WChar
		v5     *libc.WChar
		v6     *uint16
		v7     *uint32
	)
	if dword_5d4594_814548 != 0 {
		switch dword_5d4594_814548 {
		case 2:
			sub_438BD0()
			sub_43AF90(1)
			result = 1
		case 3:
			if uint64(platformTicks()) < qword_5d4594_814956 {
				goto LABEL_29
			}
			nox_client_setConnError_43AFA0(8)
			result = 1
		case 4:
			sub_43AF90(3)
			v2 = strMan.GetStringInFile("TestCon", "C:\\NoxPost\\src\\client\\shell\\noxworld.c")
			sub_449E30((*libc.WChar)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v2)))))))
			qword_5d4594_814956 = uint64(platformTicks() + 20000)
			result = 1
		case 5:
			v3 = strMan.GetStringInFile("Password", "C:\\NoxPost\\src\\client\\shell\\noxworld.c")
			sub_449E00((*libc.WChar)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v3)))))))
			v4 = strMan.GetStringInFile("PasswordRequired", "C:\\NoxPost\\src\\client\\shell\\noxworld.c")
			sub_449E30((*libc.WChar)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v4)))))))
			sub_449EA0(7)
			sub_44A360(0)
			sub_43AF90(6)
			sub_4A24C0(1)
			result = 1
		case 7:
			sub_44A360(1)
			v5 = strMan.GetStringInFile("Connected", "C:\\NoxPost\\src\\client\\shell\\noxworld.c")
			sub_449E30((*libc.WChar)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v5)))))))
			sub_449EA0(0)
			gameSetCliDrawFunc(unsafe.Pointer(cgoFuncAddr(nox_xxx_cliDrawConnectedLoop_43B360)))
			sub_43AF90(1)
			result = 1
		case 8:
			v1 = int64(platformTicks())
			dword_5d4594_814548 = 9
			*memmap.PtrUint64(6112660, 814972) = uint64(v1 + 1000)
			result = 1
		case 9:
			if uint64(platformTicks()) <= *memmap.PtrUint64(6112660, 814972) {
				goto LABEL_29
			}
			nox_client_joinGame_438A90()
			result = 1
		case 10:
			v6 = (*uint16)(unsafe.Pointer(uintptr(sub_449E60(4))))
			v7 = (*uint32)(unsafe.Pointer(GUIChildByID(4001)))
			if v6 != nil && int32(*v6) != 0 {
				nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v7)))))), 1)
				result = 1
			} else {
				nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v7)))))), 0)
				result = 1
			}
		default:
			goto LABEL_29
		}
	} else {
		if ((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_814984))))).Flags().IsHidden() == 0 || (*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_814988))))).Flags().IsHidden() == 0 || (*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&nox_wol_wnd_gameList_815012))))).Flags().IsHidden() == 0) && nox_game_createOrJoin_815048 == 0 && dword_5d4594_815044 == 0 && dword_5d4594_815052 == 0 && (*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815000))))).Flags().IsHidden() != 0 {
			if uint64(platformTicks()) > qword_5d4594_815068 {
				nox_client_refreshServerList_4378B0()
				return 1
			}
			sub_438770_waitList()
		}
	LABEL_29:
		result = 1
	}
	return result
}
func sub_438BD0() int32 {
	var (
		v0 int32
		v1 *libc.WChar
		v2 *libc.WChar
	)
	if sub_44A4A0() == 0 {
		NewDialogWindow((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&nox_wol_wnd_world_814980))))), "", nil, 0, nil, nil)
	}
	v0 = int32(nox_client_connError_814552)
	if nox_client_connError_814552 != 8 && nox_client_connError_814552 != 9 && nox_client_connError_814552 != 10 {
		v1 = strMan.GetStringInFile("ConnError", "C:\\NoxPost\\src\\client\\shell\\noxworld.c")
		sub_449E00((*libc.WChar)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v1)))))))
		v0 = int32(nox_client_connError_814552)
	}
	v2 = strMan.GetStringInFile(*(**byte)(memmap.PtrOff(0x587000, uint32(v0*4)+87416)), "C:\\NoxPost\\src\\client\\shell\\noxworld.c")
	sub_449E30((*libc.WChar)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v2)))))))
	dword_5d4594_815044 = 0
	sub_449EA0(1)
	sub_44A360(1)
	if dword_587000_87404 == 1 {
		sub_40D380()
	}
	return sub_4A24C0(1)
}
func sub_439370(a1 *int2, a2 int32) {
	var (
		v2 int32
		v3 *libc.WChar
	)
	if (int32(*(*uint8)(unsafe.Pointer(uintptr(a2 + 164))))&16) == 0 || (func() int32 {
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v2))), 0)) = uint8(int8(bool2int(nox_client_checkQuestExp_SKU2_4D7700())))
		return v2
	}()) != 0 {
		sub_439450(a1.field_0, a1.field_4, (*uint32)(unsafe.Pointer(a1)))
		sub_46B120((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_815000)))), nil)
		nox_xxx_wndShowModalMB((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815000))))))
		sub_46C690((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815000))))))
		(*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_815000)))).SetPos(image.Pt(a1.field_0, a1.field_4))
		sub_4394D0(a2)
		dword_5d4594_815056 = 1
		*memmap.PtrUint16(6112660, 814604) = *(*uint16)(unsafe.Pointer(uintptr(a2 + 109)))
		if nox_common_getEngineFlag(nox_engine_flag(NOX_ENGINE_FLAG_DISABLE_GRAPHICS_RENDERING)) {
			nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(*memmap.PtrInt32(6112660, 815008)))), 0)
		}
	} else {
		v3 = strMan.GetStringInFile("GeneralPrint:InformExpansion", "C:\\NoxPost\\src\\client\\shell\\noxworld.c")
		NewDialogWindow((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&nox_wol_wnd_world_814980))))), "", (*libc.WChar)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v3)))))), 33, nil, nil)
		sub_44A360(0)
		sub_44A4B0()
	}
}
func sub_4394D0(a1 int32) int32 {
	var (
		v1     int32
		v2     *libc.WChar
		v3     *libc.WChar
		v4     *libc.WChar
		v5     *libc.WChar
		v6     *libc.WChar
		v7     *libc.WChar
		v8     *libc.WChar
		v9     int16
		v10    *libc.WChar
		v11    *libc.WChar
		result int32
		v13    *libc.WChar
		v14    *libc.WChar
		v15    int32
		v16    int32
		v17    int32
		v18    int32
		v19    int32
		v20    *libc.WChar
		v21    int32
		v22    *libc.WChar
		v23    int32
		v24    int32
		v25    int32
		v26    int32
		v27    *libc.WChar
		v28    int32
		v29    *libc.WChar
		v30    int32
		v31    int32
		v32    *libc.WChar
		v33    *libc.WChar
		v34    uint8
		v35    uint8
		v36    int32
		v37    [16]byte
		v38    [100]libc.WChar
	)
	v36 = a1 + 111
	v1 = int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 102)))) & math.MaxInt8
	(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815004))))).Func94(asWindowEvent(0x400F, 0, 0))
	v2 = strMan.GetStringInFile("Name", "C:\\NoxPost\\src\\client\\shell\\noxworld.c")
	(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815004))))).Func94(asWindowEvent(0x400D, int32(uintptr(unsafe.Pointer(v2))), 14))
	if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 120)))) != 0 {
		libc.StrNCpy(&v37[0], (*byte)(unsafe.Pointer(uintptr(a1+120))), 15)
		v37[15] = 0
	} else {
		nox_sprintAddrPort_43BC80((*byte)(unsafe.Pointer(uintptr(a1+12))), *(*uint16)(unsafe.Pointer(uintptr(a1 + 109))), &v37[0])
	}
	nox_swprintf(&v38[0], libc.CWString("%S"), &v37[0])
	(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815004))))).Func94(asWindowEvent(0x400D, int32(uintptr(unsafe.Pointer(&v38[0]))), -1))
	if dword_587000_87404 == 1 {
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815004))))).Func94(asWindowEvent(0x400D, int32(uintptr(memmap.PtrOff(0x587000, 89332))), -1))
		v3 = strMan.GetStringInFile("GameHost", "C:\\NoxPost\\src\\client\\shell\\noxworld.c")
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815004))))).Func94(asWindowEvent(0x400D, int32(uintptr(unsafe.Pointer(v3))), 14))
		sub_439CC0(int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 32)))), &v37[0])
		nox_swprintf(&v38[0], libc.CWString("%S"), &v37[0])
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815004))))).Func94(asWindowEvent(0x400D, int32(uintptr(unsafe.Pointer(&v38[0]))), -1))
	}
	(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815004))))).Func94(asWindowEvent(0x400D, int32(uintptr(memmap.PtrOff(0x587000, 89396))), -1))
	v4 = strMan.GetStringInFile("Ping", "C:\\NoxPost\\src\\client\\shell\\noxworld.c")
	(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815004))))).Func94(asWindowEvent(0x400D, int32(uintptr(unsafe.Pointer(v4))), 14))
	if *(*uint32)(unsafe.Pointer(uintptr(a1 + 96))) == 9999 {
		nox_swprintf(&v38[0], libc.CWString("--"))
	} else {
		nox_swprintf(&v38[0], libc.CWString("%d"), *(*uint32)(unsafe.Pointer(uintptr(a1 + 96))))
	}
	(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815004))))).Func94(asWindowEvent(0x400D, int32(uintptr(unsafe.Pointer(&v38[0]))), -1))
	(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815004))))).Func94(asWindowEvent(0x400D, int32(uintptr(memmap.PtrOff(0x587000, 89464))), -1))
	v5 = strMan.GetStringInFile("GameType", "C:\\NoxPost\\src\\client\\shell\\noxworld.c")
	(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815004))))).Func94(asWindowEvent(0x400D, int32(uintptr(unsafe.Pointer(v5))), 14))
	v6 = nox_gui_wol_gameModeString_43BCB0(int16(*(*uint16)(unsafe.Pointer(uintptr(a1 + 163)))))
	(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815004))))).Func94(asWindowEvent(0x400D, int32(uintptr(unsafe.Pointer(v6))), -1))
	if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 164))))&16 != 0 {
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815004))))).Func94(asWindowEvent(0x400D, int32(uintptr(memmap.PtrOff(0x587000, 89520))), -1))
		v7 = strMan.GetStringInFile("Stage", "C:\\NoxPost\\src\\client\\shell\\noxworld.c")
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815004))))).Func94(asWindowEvent(0x400D, int32(uintptr(unsafe.Pointer(v7))), 14))
		nox_swprintf(&v38[0], libc.CWString("%d"), *(*uint16)(unsafe.Pointer(uintptr(a1 + 165))))
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815004))))).Func94(asWindowEvent(0x400D, int32(uintptr(unsafe.Pointer(&v38[0]))), -1))
	}
	(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815004))))).Func94(asWindowEvent(0x400D, int32(uintptr(memmap.PtrOff(0x587000, 89580))), -1))
	v8 = strMan.GetStringInFile("Map", "C:\\NoxPost\\src\\client\\shell\\noxworld.c")
	(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815004))))).Func94(asWindowEvent(0x400D, int32(uintptr(unsafe.Pointer(v8))), 14))
	libc.StrCpy(&v37[0], (*byte)(unsafe.Pointer(uintptr(a1+111))))
	nox_swprintf(&v38[0], libc.CWString("%S"), &v37[0])
	(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815004))))).Func94(asWindowEvent(0x400D, int32(uintptr(unsafe.Pointer(&v38[0]))), -1))
	(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815004))))).Func94(asWindowEvent(0x400D, int32(uintptr(memmap.PtrOff(0x587000, 89636))), -1))
	v9 = int16(*(*uint16)(unsafe.Pointer(uintptr(a1 + 163))))
	if int32(v9)&0xC000 != 0 {
		if int32(v9)&0x4000 != 0 {
			v33 = strMan.GetStringInFile("Individual", "C:\\NoxPost\\src\\client\\shell\\noxworld.c")
		} else {
			v33 = strMan.GetStringInFile("Clan", "C:\\NoxPost\\src\\client\\shell\\noxworld.c")
		}
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815004))))).Func94(asWindowEvent(0x400D, int32(uintptr(unsafe.Pointer(v33))), 6))
		v10 = strMan.GetStringInFile("Ladder", "C:\\NoxPost\\src\\client\\shell\\noxworld.c")
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815004))))).Func94(asWindowEvent(0x400D, int32(uintptr(unsafe.Pointer(v10))), 6))
	}
	(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815004))))).Func94(asWindowEvent(0x400D, int32(uintptr(memmap.PtrOff(0x587000, 89788))), -1))
	v11 = strMan.GetStringInFile("Occupancy", "C:\\NoxPost\\src\\client\\shell\\noxworld.c")
	(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815004))))).Func94(asWindowEvent(0x400D, int32(uintptr(unsafe.Pointer(v11))), 14))
	nox_swprintf(&v38[0], libc.CWString("%d/%d\n"), *(*uint8)(unsafe.Pointer(uintptr(a1 + 103))), *(*uint8)(unsafe.Pointer(uintptr(a1 + 104))))
	result = (*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815004))))).Func94(asWindowEvent(0x400D, int32(uintptr(unsafe.Pointer(&v38[0]))), -1))
	if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 164))))&32 != 0 {
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815004))))).Func94(asWindowEvent(0x400D, int32(uintptr(memmap.PtrOff(0x587000, 89860))), -1))
		v13 = strMan.GetStringInFile("Resolution", "C:\\NoxPost\\src\\client\\shell\\noxworld.c")
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815004))))).Func94(asWindowEvent(0x400D, int32(uintptr(unsafe.Pointer(v13))), 14))
		var rstr int32 = int32(uintptr(unsafe.Pointer(get_video_mode_string(v1))))
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815004))))).Func94(asWindowEvent(0x400D, rstr, -1))
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815004))))).Func94(asWindowEvent(0x400D, int32(uintptr(memmap.PtrOff(0x587000, 89916))), -1))
		v14 = strMan.GetStringInFile("DisabledSpells", "C:\\NoxPost\\src\\client\\shell\\noxworld.c")
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815004))))).Func94(asWindowEvent(0x400D, int32(uintptr(unsafe.Pointer(v14))), 14))
		v15 = 0
		v16 = 1
		v34 = 0
		v17 = 1
		v18 = 136
		for {
			if uint32(v16) == 0x80000000 {
				v16 = 1
				v34++
			} else {
				v16 *= 2
			}
			if nox_xxx_spellIsValid_424B50(v17) && (uint32(v16)&*(*uint32)(unsafe.Pointer(uintptr(v36 + int32(v34)*4 + 24)))) == 0 && nox_xxx_spellFlags_424A70(v17)&0x7000000 != 0 {
				v15 = 1
				v19 = int32(uintptr(unsafe.Pointer(nox_xxx_spellTitle_424930(v17))))
				(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815004))))).Func94(asWindowEvent(0x400D, v19, 4))
			}
			v17++
			v18--
			if v18 == 0 {
				break
			}
		}
		if v15 == 0 {
			v20 = strMan.GetStringInFile("None", "C:\\NoxPost\\src\\client\\shell\\noxworld.c")
			(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815004))))).Func94(asWindowEvent(0x400D, int32(uintptr(unsafe.Pointer(v20))), 4))
		}
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815004))))).Func94(asWindowEvent(0x400D, int32(uintptr(memmap.PtrOff(0x587000, 90024))), -1))
		v21 = 0
		v22 = strMan.GetStringInFile("DisabledWeapons", "C:\\NoxPost\\src\\client\\shell\\noxworld.c")
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815004))))).Func94(asWindowEvent(0x400D, int32(uintptr(unsafe.Pointer(v22))), 14))
		v23 = 1
		v35 = 0
		v24 = 1
		v25 = 27
		for {
			if (int32(uint8(int8(v23))) & int32(*(*uint8)(unsafe.Pointer(uintptr(int32(v35) + v36 + 44))))) == 0 {
				v26 = sub_4159F0(v24)
				if v26 != 0 {
					(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815004))))).Func94(asWindowEvent(0x400D, v26, -1))
					v21 = 1
				}
			}
			if v23 == 128 {
				v23 = 1
				v35++
			} else {
				v23 *= 2
			}
			v24 *= 2
			v25--
			if v25 == 0 {
				break
			}
		}
		if v21 == 0 {
			v27 = strMan.GetStringInFile("None", "C:\\NoxPost\\src\\client\\shell\\noxworld.c")
			(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815004))))).Func94(asWindowEvent(0x400D, int32(uintptr(unsafe.Pointer(v27))), 4))
		}
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815004))))).Func94(asWindowEvent(0x400D, int32(uintptr(memmap.PtrOff(0x587000, 90132))), -1))
		v28 = 0
		v29 = strMan.GetStringInFile("DisabledArmor", "C:\\NoxPost\\src\\client\\shell\\noxworld.c")
		result = (*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815004))))).Func94(asWindowEvent(0x400D, int32(uintptr(unsafe.Pointer(v29))), 14))
		v30 = 1
		v31 = 26
		for {
			if (uint32(v30) & *(*uint32)(unsafe.Pointer(uintptr(a1 + 159)))) == 0 {
				result = sub_415E80(v30)
				if result != 0 {
					v28 = 1
					result = (*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815004))))).Func94(asWindowEvent(0x400D, result, -1))
				}
			}
			v30 *= 2
			v31--
			if v31 == 0 {
				break
			}
		}
		if v28 == 0 {
			v32 = strMan.GetStringInFile("None", "C:\\NoxPost\\src\\client\\shell\\noxworld.c")
			result = (*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815004))))).Func94(asWindowEvent(0x400D, int32(uintptr(unsafe.Pointer(v32))), 4))
		}
	}
	return result
}
func nox_xxx_windowMultiplayerSub_439E70(a1 int32, a2 uint32, a3 *int32, a4 int32) int32 {
	var (
		v4     int32
		v5     int32
		v6     *int32
		v7     *int32
		v8     int32
		v9     int32
		v10    int32
		result int32
		v12    *libc.WChar
		v13    int32
		v14    int32
		v15    *libc.WChar
		v16    int32
		v17    *uint32
		v18    *uint32
		v19    int32
		v20    int32
		v21    *libc.WChar
		v22    int32
		v23    *libc.WChar
		v24    *byte
		v25    *byte
		v26    *byte
		v27    int32
		v28    uint16
		v30    int32
		v31    int32
	)
	_ = v31
	var buf [4]byte
	var v33 [10]libc.WChar
	var v34 [36]byte
	var v35 [72]byte
	if a2 > 16400 {
		if a2 == 0x4013 || a2 == 0x401C {
			(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815016))))).Func94(asWindowEvent(int32(a2), int32(uintptr(unsafe.Pointer(a3))), 0))
			(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815020))))).Func94(asWindowEvent(int32(a2), int32(uintptr(unsafe.Pointer(a3))), 0))
			(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815024))))).Func94(asWindowEvent(int32(a2), int32(uintptr(unsafe.Pointer(a3))), 0))
			(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815028))))).Func94(asWindowEvent(int32(a2), int32(uintptr(unsafe.Pointer(a3))), 0))
			(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815032))))).Func94(asWindowEvent(int32(a2), int32(uintptr(unsafe.Pointer(a3))), 0))
		}
		return 0
	}
	switch a2 {
	case 16400:
		if (*nox_window)(unsafe.Pointer(a3)).ID() == 0x274D {
			var mpos nox_point = getMousePos()
			dword_5d4594_814624 = unsafe.Pointer(uintptr(sub_4A28C0(a4)))
			sub_439370((*int2)(unsafe.Pointer(&mpos)), int32(uintptr(dword_5d4594_814624)))
			return 0
		}
		return 0
	case 23:
		return 1
	case 0x4000:
		v27 = (*nox_window)(unsafe.Pointer(a3)).ID()
		if v27 >= 0x273B && v27 <= 0x273C {
			(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815016))))).Func94(asWindowEvent(0x4000, int32(uintptr(unsafe.Pointer(a3))), 0))
			(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815020))))).Func94(asWindowEvent(0x4000, int32(uintptr(unsafe.Pointer(a3))), 0))
			(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815024))))).Func94(asWindowEvent(0x4000, int32(uintptr(unsafe.Pointer(a3))), 0))
			(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815028))))).Func94(asWindowEvent(0x4000, int32(uintptr(unsafe.Pointer(a3))), 0))
			(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815032))))).Func94(asWindowEvent(0x4000, int32(uintptr(unsafe.Pointer(a3))), 0))
			return 0
		}
		return 0
	}
	if a2 != 0x4007 {
		return 0
	}
	v4 = (*nox_window)(unsafe.Pointer(a3)).ID()
	v5 = v4
	if v4 != 0x273B && v4 != 0x273C && v4 != 0x2733 && v4 != 0x2734 {
		clientPlaySoundSpecial(766, 100)
	}
	if v5 >= 10070 {
		var mpos nox_point = getMousePos()
		v30 = mpos.x - 216
		v31 = mpos.y - 27
		v6 = (*int32)(unsafe.Pointer(sub_4A0020()))
		if sub_4A25C0((*uint32)(unsafe.Pointer(&v30)), v6) >= 2 {
			v7 = (*int32)(unsafe.Pointer(sub_4A0020()))
			*memmap.PtrUint32(6112660, 815036) = uint32(sub_4A2610(*(*int32)(unsafe.Pointer(&nox_wol_wnd_world_814980)), (*uint32)(unsafe.Pointer(&v30)), v7))
		} else {
			dword_5d4594_814624 = unsafe.Pointer(sub_4A0490(v5 - 10070))
			sub_439370((*int2)(unsafe.Pointer(&mpos)), int32(uintptr(dword_5d4594_814624)))
		}
		nox_xxx_cursorSetTooltip_4776B0((*libc.WChar)(memmap.PtrOff(6112660, 815112)))
	}
	if v5 > 0x2716 {
		switch v5 {
		case 0x2717:
			if nox_game_createOrJoin_815048 != 0 {
				sub_4373A0()
			}
			if dword_587000_87404 != 0 {
				sub_43A980()
			}
			return 0
		case 10010:
			sub_4373A0()
			return 0
		case 0x273F:
			fallthrough
		case 0x2740:
			fallthrough
		case 0x2741:
			fallthrough
		case 10050:
			fallthrough
		case 0x2743:
			sub_4379C0()
			nox_wol_servers_sortBtnHandler_4A0290(v5)
			sub_4A0390()
			return 0
		case 0x2746:
			dword_587000_87412 = 0
			sub_43A810()
			nox_client_refreshServerList_4378B0()
			if noxflags.HasGame(noxflags.GameFlag26) {
				v24 = sub_4A7EF0()
				sub_439D90(*(*uint32)(unsafe.Pointer(v24))+216, *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v24))), unsafe.Sizeof(uint32(0))*1)))+27)
			} else {
				sub_439D90(408, 239)
			}
			result = 0
		case 0x2747:
			dword_587000_87412 = 1
			sub_43A810()
			nox_client_refreshServerList_4378B0()
			if !noxflags.HasGame(noxflags.GameFlag26) {
				sub_439D90(408, 239)
				result = 0
			} else {
				v25 = sub_4A7EF0()
				sub_439D90(*(*uint32)(unsafe.Pointer(v25))+216, *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v25))), unsafe.Sizeof(uint32(0))*1)))+27)
				return 0
			}
		case 0x2748:
			dword_587000_87412 = 2
			sub_43A810()
			nox_client_refreshServerList_4378B0()
			if !noxflags.HasGame(noxflags.GameFlag26) {
				sub_439D90(408, 239)
			} else {
				v24 = sub_4A7EF0()
				sub_439D90(*(*uint32)(unsafe.Pointer(v24))+216, *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v24))), unsafe.Sizeof(uint32(0))*1)))+27)
			}
			result = 0
		case 0x2749:
			dword_587000_87412 = 3
			sub_43A810()
			nox_client_refreshServerList_4378B0()
			if noxflags.HasGame(noxflags.GameFlag26) {
				v26 = sub_4A7EF0()
				sub_439D90(*(*uint32)(unsafe.Pointer(v26))+216, *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v26))), unsafe.Sizeof(uint32(0))*1)))+27)
			} else {
				sub_439D90(408, 239)
			}
			result = 0
		default:
			return 0
		}
		return result
	}
	if v5 == 0x2716 {
		nox_game_createOrJoin_815048 = 0
		nox_client_refreshServerList_4378B0()
		return 0
	}
	if v5 > 0x2712 {
		v19 = v5 - 0x2713
		if v5 != 0x2713 {
			v20 = v19 - 1
			if v20 == 0 {
				nox_game_createOrJoin_815048 = 0
				sub_4383A0()
				nox_client_refreshServerList_4378B0()
				return 0
			}
			if v20 == 1 {
				nox_game_createOrJoin_815048 = 0
				(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&nox_wol_wnd_gameList_815012))))).Hide()
				(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_814984))))).Hide()
				(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_814988))))).Hide()
				(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_814992))))).Show()
				sub_46AD20(*(**uint32)(unsafe.Pointer(&nox_wol_wnd_world_814980)), 0x2716, 0x2717, 0)
				sub_46ACE0(*(**uint32)(unsafe.Pointer(&nox_wol_wnd_world_814980)), 0x273F, 0x2743, 1)
				v21 = strMan.GetStringInFile("FilterMsg", "C:\\NoxPost\\src\\client\\shell\\noxworld.c")
				(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_814996))))).Func94(asWindowEvent(0x4001, int32(uintptr(unsafe.Pointer(v21))), 0))
				return 0
			}
			return 0
		}
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v22))), 0)) = uint8(int8(bool2int(nox_client_checkQuestExp_SKU2_4D7700())))
		if v22 == 0 {
			v23 = strMan.GetStringInFile("GeneralPrint:InformExpansion", "C:\\NoxPost\\src\\client\\shell\\noxworld.c")
			NewDialogWindow((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&nox_wol_wnd_world_814980))))), "", (*libc.WChar)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v23)))))), 33, nil, nil)
			sub_44A360(0)
			sub_44A4B0()
			return 0
		}
		nox_game_createOrJoin_815048 = 1
		noxflags.UnsetGame(noxflags.GameNotQuest)
		nox_client_setMousePos_430B10(408, 239)
		sub_4375C0(0)
		if (*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_814992))))).Flags().IsHidden() == 0 {
			sub_489870()
		}
		nox_xxx_setQuest_4D6F60(1)
		nox_xxx_cliShowHideTubes_470AA0(1)
		if dword_587000_87404 == 1 {
			if *(*int32)(unsafe.Pointer(&dword_587000_87412)) != -1 {
				return 0
			}
			v16 = noxRndCounter1.IntClamp(0x2746, 0x2749)
			v17 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&nox_wol_wnd_world_814980)))).ChildByID(v16)))
			(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&nox_wol_wnd_world_814980))))).Func94(asWindowEvent(0x4007, int32(uintptr(unsafe.Pointer(v17))), 0))
			return 0
		}
		dword_587000_87412 = 0
		v18 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&nox_wol_wnd_world_814980)))).ChildByID(10020)))
		(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v18)))))).Func93(asWindowEvent(0x5, 0xEF0198, 0))
		return 0
	}
	if v5 == 0x2712 {
		if noxflags.HasGame(noxflags.GameFlag25) {
			return 0
		}
		nox_game_createOrJoin_815048 = 1
		noxflags.SetGame(noxflags.GameNotQuest)
		nox_xxx_cliShowHideTubes_470AA0(0)
		nox_client_setMousePos_430B10(408, 239)
		sub_4375C0(0)
		if (*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_814992))))).Flags().IsHidden() == 0 {
			sub_489870()
		}
		if noxflags.HasGame(noxflags.GameFlag26) || sub_4D6F30() != 0 {
			return 0
		}
		if dword_587000_87404 == 1 {
			if *(*int32)(unsafe.Pointer(&dword_587000_87412)) != -1 {
				return 0
			}
			v16 = noxRndCounter1.IntClamp(0x2746, 0x2749)
			v17 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&nox_wol_wnd_world_814980)))).ChildByID(v16)))
			(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&nox_wol_wnd_world_814980))))).Func94(asWindowEvent(0x4007, int32(uintptr(unsafe.Pointer(v17))), 0))
			return 0
		}
		dword_587000_87412 = 0
		v18 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&nox_wol_wnd_world_814980)))).ChildByID(10020)))
		(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v18)))))).Func93(asWindowEvent(0x5, 0xEF0198, 0))
		return 0
	}
	v8 = v5 - 4001
	if v8 == 0 {
		if sub_43AF80() == 6 {
			v12 = (*libc.WChar)(unsafe.Pointer(uintptr(sub_449E60(4))))
			nox_wcsncpy(&v33[0], v12, 9)
			v33[8] = 0
			v28 = uint16(int16(clientGetServerPort()))
			v13 = int32(nox_client_getServerAddr_43B300())
			sub_5550D0(v13, v28, &buf[0])
			sub_43AF90(3)
			qword_5d4594_814956 = uint64(platformTicks() + 20000)
			sub_449EA0(0)
			return 0
		}
		if sub_43AF80() == 10 {
			v14 = sub_449E60(4)
			nox_sprintf(&v35[0], CString("%S"), v14)
			v15 = strMan.GetStringInFile("Finding", "C:\\NoxPost\\src\\client\\shell\\noxworld.c")
			sub_449E30((*libc.WChar)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v15)))))))
			sub_43AF90(11)
			sub_40D740(int32(uintptr(unsafe.Pointer(&v34[0]))))
			sub_449EA0(0)
			return 0
		}
		if sub_43AF80() == 1 {
			if dword_587000_87404 == 1 {
				sub_40D380()
			}
			sub_43A920()
			nox_game_showGameSel_4379F0()
			return 0
		}
		if dword_5d4594_815104 != 0 {
			qword_5d4594_815068 = uint64(platformTicks() + 1000)
			return 0
		}
		return 0
	}
	v9 = v8 - 1
	if v9 == 0 {
		if dword_587000_87404 == 1 {
			sub_40D380()
		}
		sub_43A920()
		nox_game_showGameSel_4379F0()
		return 0
	}
	if v9 != 5999 || dword_5d4594_815044 != 0 {
		return 0
	}
	v10 = sub_43B340()
	if v10&4096 != 0 {
		sub_4D6F80(1)
		nox_xxx_cliShowHideTubes_470AA0(1)
	}
	sub_43B460()
	if sub_4D6F70() != 0 {
		if nox_client_countPlayerFiles04_4DC7D0() != 0 {
			sub_4A7A70(1)
			nox_game_showSelChar_4A4DB0()
			dword_5d4594_815044 = 1
			sub_4A2890()
			return 0
		}
	} else if nox_client_countPlayerFiles02_4DC630() != 0 {
		sub_4A7A70(1)
		nox_game_showSelChar_4A4DB0()
		dword_5d4594_815044 = 1
		sub_4A2890()
		return 0
	}
	sub_4A7A70(0)
	nox_game_showSelClass_4A4840()
	dword_5d4594_815044 = 1
	sub_4A2890()
	return 0
}
func sub_43A810() {
	var (
		v0 int32
		v1 *int32
		v2 *libc.WChar
	)
	(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_814984))))).Show()
	(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_814988))))).Hide()
	nox_xxx_wndSetIcon_46AE60(*(*int32)(unsafe.Pointer(&dword_5d4594_814984)), int32(*memmap.PtrUint32(6112660, dword_587000_87412*4+814900)))
	if nox_game_createOrJoin_815048 == 1 {
		v2 = strMan.GetStringInFile("CreateMsg", "C:\\NoxPost\\src\\client\\shell\\noxworld.c")
	} else {
		v2 = strMan.GetStringInFile("JoinServer", "C:\\NoxPost\\src\\client\\shell\\noxworld.c")
	}
	(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_814996))))).Func94(asWindowEvent(0x4001, int32(uintptr(unsafe.Pointer(v2))), 0))
	if dword_587000_87404 == 1 {
		sub_46ACE0(*(**uint32)(unsafe.Pointer(&dword_5d4594_814984)), 10600, 0x2973, 1)
		v0 = 0
		v1 = memmap.PtrInt32(0x587000, 87560)
		for {
			sub_46ACE0(*(**uint32)(unsafe.Pointer(&dword_5d4594_814984)), *v1, *(*int32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(int32(0))*1)), bool2int(uint32(v0) != dword_587000_87412))
			v1 = (*int32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(int32(0))*2))
			v0++
			if int32(uintptr(unsafe.Pointer(v1))) >= int32(uintptr(memmap.PtrOff(0x587000, 87592))) {
				break
			}
		}
	}
	sub_49FDB0(*(*int32)(unsafe.Pointer(&dword_587000_87412)))
}
func sub_43A980() int32 {
	var (
		v0 *libc.WChar
		v2 *libc.WChar
	)
	v2 = strMan.GetStringInFile("FindText", "C:\\NoxPost\\src\\client\\shell\\noxworld.c")
	v0 = strMan.GetStringInFile("Find", "C:\\NoxPost\\src\\client\\shell\\noxworld.c")
	NewDialogWindow((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&nox_wol_wnd_world_814980))))), (*libc.WChar)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v0)))))), (*libc.WChar)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v2)))))), 39, nil, nil)
	return sub_43AF90(10)
}
func sub_43B630() *uint32 {
	var v0 *libc.WChar
	nox_xxx_wnd_46C6E0((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815000))))))
	v0 = strMan.GetStringInFile("AttemptingConn", "C:\\NoxPost\\src\\client\\shell\\noxworld.c")
	return (*uint32)(NewDialogWindow((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&nox_wol_wnd_world_814980))))), "", (*libc.WChar)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v0)))))), 34, nil, nil))
}
func sub_43B6E0() {
	var (
		v1 *libc.WChar
		v2 *libc.WChar
	)
	if nox_wol_wnd_world_814980 != nil {
		v2 = strMan.GetStringInFile("Kicked", "C:\\NoxPost\\src\\client\\shell\\noxworld.c")
		v1 = strMan.GetStringInFile("Notification", "C:\\NoxPost\\src\\client\\shell\\noxworld.c")
		NewDialogWindow(nil, (*libc.WChar)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v1)))))), (*libc.WChar)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v2)))))), 33, nil, nil)
		sub_44A360(1)
		dword_5d4594_815096 = 0
	} else {
		dword_5d4594_815096 = 1
	}
}
func sub_43B750() {
	var (
		v1 *libc.WChar
		v2 *libc.WChar
	)
	if nox_wol_wnd_world_814980 != nil {
		v2 = strMan.GetStringInFile("Timeout", "C:\\NoxPost\\src\\client\\shell\\noxworld.c")
		v1 = strMan.GetStringInFile("Notification", "C:\\NoxPost\\src\\client\\shell\\noxworld.c")
		NewDialogWindow(nil, (*libc.WChar)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v1)))))), (*libc.WChar)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v2)))))), 33, nil, nil)
		sub_44A360(1)
		dword_5d4594_815100 = 0
	} else {
		dword_5d4594_815100 = 1
	}
}
func nox_gui_wol_newServerLine_43B7C0(srv *nox_gui_server_ent_t) {
	var (
		wbuf [128]libc.WChar
		buf  [332]byte
	)
	if dword_587000_87408 != 0 {
		(*nox_window)(unsafe.Pointer(uintptr(nox_wol_wnd_gameList_815012))).Func94(asWindowEvent(0x400D, int32(uintptr(memmap.PtrOff(0x587000, 91164))), 4))
		libc.StrNCpy(&buf[0], &srv.server_name[0], 15)
		buf[15] = 0
		if libc.MemCmp(unsafe.Pointer(&buf[0]), memmap.PtrOff(6112660, 815120), 1) == 0 {
			nox_sprintAddrPort_43BC80(&srv.addr[0], srv.port, &buf[0])
		}
		nox_swprintf(&wbuf[0], libc.CWString("%S"), &buf[0])
		sub_43BC10(&wbuf[0], 100)
		(*nox_window)(unsafe.Pointer(uintptr(dword_5d4594_815016))).Func94(asWindowEvent(0x400D, int32(uintptr(unsafe.Pointer(&wbuf[0]))), 4))
		nox_swprintf(&wbuf[0], libc.CWString("%d/%d"), srv.players, srv.max_players)
		(*nox_window)(unsafe.Pointer(uintptr(dword_5d4594_815020))).Func94(asWindowEvent(0x400D, int32(uintptr(unsafe.Pointer(&wbuf[0]))), 4))
		var v6 *libc.WChar = nox_gui_wol_gameModeString_43BCB0(int16(srv.flags))
		if int32(srv.flags)&4096 != 0 {
			nox_swprintf((*libc.WChar)(memmap.PtrOff(6112660, 814772)), libc.CWString("%s %d"), v6, srv.quest_level)
			(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815024))))).Func94(asWindowEvent(0x400D, int32(uintptr(memmap.PtrOff(6112660, 814772))), 4))
		} else {
			(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815024))))).Func94(asWindowEvent(0x400D, int32(uintptr(unsafe.Pointer(v6))), 4))
		}
		if srv.ping == 9999 {
			nox_swprintf(&wbuf[0], libc.CWString("--"))
		} else {
			compat_itow(srv.ping, &wbuf[0], 10)
		}
		(*nox_window)(unsafe.Pointer(uintptr(dword_5d4594_815028))).Func94(asWindowEvent(0x400D, int32(uintptr(unsafe.Pointer(&wbuf[0]))), 4))
		wbuf[0] = 0
		if int32(srv.status)&32 != 0 {
			var v9 *libc.WChar = strMan.GetStringInFile("Noxworld.wnd:private", "C:\\NoxPost\\src\\client\\shell\\noxworld.c")
			nox_wcscat(&wbuf[0], v9)
		}
		if int32(srv.status)&16 != 0 {
			if wbuf[0] != 0 {
				nox_wcscat(&wbuf[0], libc.CWString("+"))
			}
			var v10 *libc.WChar = strMan.GetStringInFile("Noxworld.wnd:closed", "C:\\NoxPost\\src\\client\\shell\\noxworld.c")
			nox_wcscat(&wbuf[0], v10)
		}
		if wbuf[0] == 0 {
			var v11 *libc.WChar
			if int32(srv.players) < int32(srv.max_players) {
				v11 = strMan.GetStringInFile("Open", "C:\\NoxPost\\src\\client\\shell\\noxworld.c")
			} else {
				v11 = strMan.GetStringInFile("Full", "C:\\NoxPost\\src\\client\\shell\\noxworld.c")
			}
			nox_wcscat(&wbuf[0], v11)
		}
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_815032))))).Func94(asWindowEvent(0x400D, int32(uintptr(unsafe.Pointer(&wbuf[0]))), 4))
	} else {
		var a1 int32 = int32(uintptr(unsafe.Pointer(srv)))
		*(*[332]byte)(unsafe.Pointer(&buf[0])) = [332]byte{}
		*(*uint32)(unsafe.Pointer(&buf[8])) = 257
		*(*uint32)(unsafe.Pointer(&buf[16])) = uint32(uintptr(unsafe.Pointer(nox_wol_wnd_world_814980)))
		var v4 int32
		if *(*int32)(unsafe.Pointer(&dword_587000_87412)) == -1 {
			var (
				v1 int32 = sub_437860(int32(*(*int16)(unsafe.Pointer(uintptr(a1 + 44)))), int32(*(*int16)(unsafe.Pointer(uintptr(a1 + 46)))))
				v2 int16 = int16(*memmap.PtrUint16(0x587000, uint32(v1*8)+87528))
			)
			v1 += 0x2746
			*(*uint16)(unsafe.Pointer(uintptr(a1 + 44))) -= uint16(v2)
			*(*uint16)(unsafe.Pointer(uintptr(a1 + 46))) -= *memmap.PtrUint16(0x587000, uint32(v1*8+7098))
			*(*uint16)(unsafe.Pointer(uintptr(a1 + 44))) >>= 1
			*(*uint16)(unsafe.Pointer(uintptr(a1 + 46))) >>= 1
			var v3 *uint32 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_814988)))).ChildByID(v1)))
			v4 = int32(uintptr(unsafe.Pointer(nox_gui_newButtonOrCheckbox_4A91A0((*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v3)))))), 1185, int32(*(*int16)(unsafe.Pointer(uintptr(a1 + 44))))-5, int32(*(*int16)(unsafe.Pointer(uintptr(a1 + 46))))-5, 10, 10, (*nox_window_data)(unsafe.Pointer(&buf[0]))))))
		} else {
			*(*uint16)(unsafe.Pointer(uintptr(a1 + 44))) -= *memmap.PtrUint16(0x587000, dword_587000_87412*8+87528)
			*(*uint16)(unsafe.Pointer(uintptr(a1 + 46))) -= *memmap.PtrUint16(0x587000, dword_587000_87412*8+87530)
			v4 = int32(uintptr(unsafe.Pointer(nox_gui_newButtonOrCheckbox_4A91A0((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_814984))))), 1192, int32(*(*int16)(unsafe.Pointer(uintptr(a1 + 44))))-10, int32(*(*int16)(unsafe.Pointer(uintptr(a1 + 46))))-10, 20, 20, (*nox_window_data)(unsafe.Pointer(&buf[0]))))))
		}
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 28))) = uint32(v4)
		sub_437320(a1)
		var v13 [32]byte
		if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 120)))) != 0 {
			libc.StrNCpy(&v13[0], (*byte)(unsafe.Pointer(uintptr(a1+120))), 15)
			v13[15] = 0
		} else {
			nox_sprintAddrPort_43BC80((*byte)(unsafe.Pointer(uintptr(a1+12))), *(*uint16)(unsafe.Pointer(uintptr(a1 + 109))), &v13[0])
		}
		if *(*uint32)(unsafe.Pointer(uintptr(a1 + 96))) == 9999 {
			nox_swprintf(&wbuf[0], libc.CWString("%S -- ms"), &v13[0])
		} else {
			nox_swprintf(&wbuf[0], libc.CWString("%S %dms"), &v13[0], *(*uint32)(unsafe.Pointer(uintptr(a1 + 96))))
		}
		nox_xxx_wndWddSetTooltip_46B000((*nox_window_data)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 28)))+36))), &wbuf[0])
		int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 28)))).setFunc94(func(arg1 int32, arg2 int32, arg3 int32, arg4 int32) int32 {
			return nox_xxx_windowMultiplayerSub_439E70(arg1, uint32(arg2), (*int32)(unsafe.Pointer(uintptr(arg3))), arg4)
		})
		var result *uint32 = *(**uint32)(unsafe.Pointer(uintptr(a1 + 28)))
		_ = result
		*result = *(*uint32)(unsafe.Pointer(uintptr(a1 + 36))) + 10070
	}
}
func nox_gui_wol_gameModeString_43BCB0(a1 int16) *libc.WChar {
	if int32(a1)&4096 != 0 {
		return strMan.GetStringInFile("Quest", "C:\\NoxPost\\src\\client\\shell\\noxworld.c")
	}
	if int32(a1)&32 != 0 {
		return strMan.GetStringInFile("CTF", "C:\\NoxPost\\src\\client\\shell\\noxworld.c")
	}
	if int32(a1)&1024 != 0 {
		return strMan.GetStringInFile("Highlander", "C:\\NoxPost\\src\\client\\shell\\noxworld.c")
	}
	if int32(a1)&16 != 0 {
		return strMan.GetStringInFile("KotR", "C:\\NoxPost\\src\\client\\shell\\noxworld.c")
	}
	if int32(a1)&64 != 0 {
		return strMan.GetStringInFile("Flagball", "C:\\NoxPost\\src\\client\\shell\\noxworld.c")
	}
	if (int32(a1) & 128) == 0 {
		return strMan.GetStringInFile("Arena", "C:\\NoxPost\\src\\client\\shell\\noxworld.c")
	}
	return strMan.GetStringInFile("Chat", "C:\\NoxPost\\src\\client\\shell\\noxworld.c")
}
