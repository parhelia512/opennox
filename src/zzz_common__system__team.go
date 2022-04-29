package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	noxflags "github.com/noxworld-dev/opennox/v1/common/flags"
	"unsafe"
)

const NOX_TEAMS_MAX = 17

type nox_team_t struct {
	name     [21]libc.WChar
	field_42 uint16
	field_44 uint32
	field_48 uint32
	field_52 uint32
	def_ind  uint8
	field_57 uint8
	ind      uint8
	field_59 uint8
	field_60 uint32
	active   uint32
	field_68 uint32
	field_72 unsafe.Pointer
	field_76 uint32
}
type nox_team_info_t struct {
	name  *byte
	title *libc.WChar
	code  int32
	color *uint32
}

var nox_color_white_2523948 uint32 = 0
var nox_color_red_2589776 uint32 = 0
var nox_color_blue_2650684 uint32 = 0
var nox_color_green_2614268 uint32 = 0
var nox_color_cyan_2649820 uint32 = 0
var nox_color_yellow_2589772 uint32 = 0
var nox_color_violet_2598268 uint32 = 0
var nox_color_black_2650656 uint32 = 0
var nox_color_orange_2614256 uint32 = 0
var nox_team_table [10]nox_team_info_t = [10]nox_team_info_t{{name: CString("advserv.wnd:None"), title: nil, code: 0, color: &nox_color_white_2523948}, {name: CString("modifier.db:MaterialTeamRedDesc"), title: nil, code: 1, color: &nox_color_red_2589776}, {name: CString("modifier.db:MaterialTeamBlueDesc"), title: nil, code: 2, color: &nox_color_blue_2650684}, {name: CString("modifier.db:MaterialTeamGreenDesc"), title: nil, code: 3, color: &nox_color_green_2614268}, {name: CString("modifier.db:MaterialTeamCyanDesc"), title: nil, code: 4, color: &nox_color_cyan_2649820}, {name: CString("modifier.db:MaterialTeamYellowDesc"), title: nil, code: 5, color: &nox_color_yellow_2589772}, {name: CString("modifier.db:MaterialTeamVioletDesc"), title: nil, code: 6, color: &nox_color_violet_2598268}, {name: CString("modifier.db:MaterialTeamBlackDesc"), title: nil, code: 7, color: &nox_color_black_2650656}, {name: CString("modifier.db:MaterialTeamWhiteDesc"), title: nil, code: 8, color: &nox_color_white_2523948}, {name: CString("modifier.db:MaterialTeamOrangeDesc"), title: nil, code: 9, color: &nox_color_orange_2614256}}
var nox_team_table_cnt int32 = int32(unsafe.Sizeof([10]nox_team_info_t{}) / unsafe.Sizeof(nox_team_info_t{}))

func nox_xxx_createAtImpl_4191D0(a1 uint8, a2p unsafe.Pointer, a3 int32, a4 int32, a5 int32) {
	var (
		a2     int32 = int32(uintptr(a2p))
		result *byte
		v6     *byte
		v7     int32
		v8     *byte
		v9     *byte
		v10    *libc.WChar
		v11    *libc.WChar
		v12    int16
		v13    int32
		v14    *uint32
		v15    *byte
		v16    *libc.WChar
		v17    *libc.WChar
		v18    *byte
		v19    int32
		i      int32
		v21    [3]int32
		v22    int16
	)
	_ = v22
	var v23 *byte
	var v24 [128]libc.WChar
	result = *(**byte)(unsafe.Pointer(&dword_5d4594_527660))
	if dword_5d4594_527660 == 0 {
		result = (*byte)(unsafe.Pointer(uintptr(nox_xxx_getNameId_4E3AA0(CString("GameBall")))))
		dword_5d4594_527660 = uint32(uintptr(unsafe.Pointer(result)))
	}
	if a2 == 0 {
		return
	}
	v6 = (*byte)(unsafe.Pointer(nox_xxx_clientGetTeamColor_418AB0(int32(a1))))
	if v6 != nil {
		result = (*byte)(unsafe.Pointer(uintptr(nox_xxx_teamCompare2_419180(a2, a1))))
		if result != nil {
			return
		}
	} else {
		v6 = (*byte)(unsafe.Pointer(nox_xxx_teamCreate_4186D0(int8(a1))))
	}
	*(*uint8)(unsafe.Pointer(uintptr(a2 + 4))) = uint8(*(*byte)(unsafe.Add(unsafe.Pointer(v6), 57)))
	*(*uint32)(unsafe.Pointer(uintptr(a2))) = *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v6))), unsafe.Sizeof(uint32(0))*11)))
	*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v6))), unsafe.Sizeof(uint32(0))*11))) = uint32(a2)
	if uint32(a4) == nox_player_netCode_85319C {
		sub_455E70(uint8(*(*byte)(unsafe.Add(unsafe.Pointer(v6), 57))))
	}
	if noxflags.HasGame(noxflags.GameHost) {
		if noxflags.HasGame(noxflags.GameOnline) {
			v7 = nox_server_getObjectFromNetCode_4ECCB0(a4)
			v8 = (*byte)(unsafe.Pointer(noxServer.getPlayerByID(a4)))
			v23 = v8
			if v8 != nil {
				if noxflags.HasGame(noxflags.GameFlag16) {
					sub_425ED0(int32(uintptr(unsafe.Pointer(v8))), 1)
				}
				if v7 != 0 && int32(*(*uint8)(unsafe.Pointer(uintptr(v7 + 8))))&4 != 0 {
					if a5 == 1 && !nox_xxx_CheckGameplayFlags_417DA0(2) && noxflags.HasGame(noxflags.GameModeChat) {
						sub_4ED970(50.0, (*float2)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v6))), unsafe.Sizeof(uint32(0))*18)))+56))), (*float2)(unsafe.Pointer(&v21[0])))
						nox_xxx_unitMove_4E7010((*nox_object_t)(unsafe.Pointer(uintptr(v7))), (*float2)(unsafe.Pointer(&v21[0])))
					}
					v9 = (*byte)(unsafe.Pointer(noxServer.getPlayerByID(a4)))
					if v9 != nil {
						if noxflags.HasGame(noxflags.GameModeQuest) {
							v10 = strMan.GetStringInFile("GeneralPrint:PlayerJoinQuest", "C:\\NoxPost\\src\\common\\System\\team.c")
							nox_swprintf(&v24[0], v10, (*byte)(unsafe.Add(unsafe.Pointer(v9), 4704)))
						} else {
							v11 = strMan.GetStringInFile("NewMember", "C:\\NoxPost\\src\\common\\System\\team.c")
							nox_swprintf(&v24[0], v11, (*byte)(unsafe.Add(unsafe.Pointer(v9), 4704)), v6)
						}
						nox_xxx_printCentered_445490(&v24[0])
					}
					v8 = v23
				}
			}
			if a3 != 0 && v7 != 0 && (v8 != nil || uint32(*(*uint16)(unsafe.Pointer(uintptr(v7 + 4)))) == dword_5d4594_527660) {
				v12 = int16(*(*uint16)(unsafe.Pointer(uintptr(v7 + 4))))
				v13 = int32(uint8(*(*byte)(unsafe.Add(unsafe.Pointer(v6), 57))))
				*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v21[0]))), unsafe.Sizeof(uint16(0))*0)) = 452
				*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v21[1]))), unsafe.Sizeof(uint16(0))*1)) = uint16(int16(a4))
				*(*uint32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(&v21[0]))), 2)))) = uint32(v13)
				v22 = v12
				sub_4571A0(a4, v13)
				nox_xxx_netSendPacket1_4E5390(159, int32(uintptr(unsafe.Pointer(&v21[0]))), 10, 0, 1)
			}
		}
	} else {
		v14 = nox_xxx_netSpriteByCodeDynamic_45A6F0(a4)
		if v14 != nil && *(*uint32)(unsafe.Add(unsafe.Pointer(v14), unsafe.Sizeof(uint32(0))*28))&4 != 0 {
			v15 = (*byte)(unsafe.Pointer(noxServer.getPlayerByID(a4)))
			if v15 != nil {
				if noxflags.HasGame(noxflags.GameModeQuest) {
					v16 = strMan.GetStringInFile("GeneralPrint:PlayerJoinQuest", "C:\\NoxPost\\src\\common\\System\\team.c")
					nox_swprintf(&v24[0], v16, (*byte)(unsafe.Add(unsafe.Pointer(v15), 4704)))
				} else {
					v17 = strMan.GetStringInFile("NewMember", "C:\\NoxPost\\src\\common\\System\\team.c")
					nox_swprintf(&v24[0], v17, (*byte)(unsafe.Add(unsafe.Pointer(v15), 4704)), v6)
				}
				nox_xxx_printCentered_445490(&v24[0])
			}
		}
	}
	*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v6))), unsafe.Sizeof(uint32(0))*12)))++
	result = (*byte)(unsafe.Pointer(nox_xxx_getFirstPlayerUnit_4DA7C0()))
	v18 = result
	if result == nil {
		return
	}
	for *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v18))), unsafe.Sizeof(uint32(0))*9))) != uint32(a4) {
		result = (*byte)(unsafe.Pointer(nox_xxx_getNextPlayerUnit_4DA7F0((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v18)))))))))
		v18 = result
		if result == nil {
			return
		}
	}
	v19 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v18))), unsafe.Sizeof(uint32(0))*187))))
	sub_4D97E0(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v19 + 276))) + 2064)))))
	sub_4E8110(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v19 + 276))) + 2064)))))
	result = nox_xxx_monsterMarkUpdate_4E8020(int32(uintptr(unsafe.Pointer(v18))))
	for i = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v18))), unsafe.Sizeof(uint32(0))*129)))); i != 0; i = int32(*(*uint32)(unsafe.Pointer(uintptr(i + 512)))) {
		if int32(*(*uint8)(unsafe.Pointer(uintptr(i + 8))))&6 != 0 {
			result = nox_xxx_monsterMarkUpdate_4E8020(i)
		}
	}
}
