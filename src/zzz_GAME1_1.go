package opennox

import (
	"github.com/gotranspile/cxgo/runtime/cnet"
	"github.com/gotranspile/cxgo/runtime/libc"
	"github.com/gotranspile/cxgo/runtime/stdio"
	"github.com/noxworld-dev/opennox/v1/common/alloc"
	noxflags "github.com/noxworld-dev/opennox/v1/common/flags"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"math"
	"unicode"
	"unsafe"
)

var nox_file_2 *File = nil

func sub_418800(a1 *libc.WChar, a2 *libc.WChar, a3 int32) {
	if a1 != nil {
		nox_wcsncpy(a1, a2, 20)
		*(*libc.WChar)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(libc.WChar(0))*20)) = 0
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a1))), unsafe.Sizeof(uint32(0))*17))) = uint32(a3)
	}
}
func sub_418830(a1 int32, a2 int32) int32 {
	var result int32
	result = a1
	if a1 != 0 {
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 60))) = uint32(a2)
	}
	return result
}
func nox_xxx_unused_418840() int32 {
	var (
		v0  int32
		v1  *byte
		v2  *byte
		v3  int16
		i   *int32
		v5  *byte
		j   int32
		v7  *uint32
		v8  int32
		v9  int32
		k   int32
		v11 int32
	)
	v0 = 0
	v1 = (*byte)(sub_416640())
	v2 = sub_4165B0()
	nox_server_teamsZzz_419030(1)
	v3 = int16(int32(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v2))), unsafe.Sizeof(uint16(0))*26)))) & 0x3FFF)
	*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v3))), unsafe.Sizeof(int16(0))-1)) |= 128
	*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v2))), unsafe.Sizeof(uint16(0))*26))) = uint16(v3)
	for i = sub_425A50(); i != nil; i = sub_425A60(i) {
		if v0 >= int32(uint8(*(*byte)(unsafe.Add(unsafe.Pointer(v1), 52)))) {
			break
		}
		v5 = (*byte)(unsafe.Pointer(nox_xxx_teamCreate_4186D0(0)))
		*(*int32)(unsafe.Add(unsafe.Pointer(i), unsafe.Sizeof(int32(0))*9)) = int32(uintptr(unsafe.Pointer(v5)))
		if v5 != nil {
			sub_418800((*libc.WChar)(unsafe.Pointer(v5)), (*libc.WChar)(unsafe.Add(unsafe.Pointer((*libc.WChar)(unsafe.Pointer(i))), unsafe.Sizeof(libc.WChar(0))*6)), 0)
			sub_418830(*(*int32)(unsafe.Add(unsafe.Pointer(i), unsafe.Sizeof(int32(0))*9)), *(*int32)(unsafe.Add(unsafe.Pointer(i), unsafe.Sizeof(int32(0))*8)))
			sub_4184D0((*nox_team_t)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(i), unsafe.Sizeof(int32(0))*9))))))
			v0++
			for j = int32(uintptr(unsafe.Pointer(nox_xxx_getFirstPlayerUnit_4DA7C0()))); j != 0; j = int32(uintptr(unsafe.Pointer(nox_xxx_getNextPlayerUnit_4DA7F0((*nox_object_t)(unsafe.Pointer(uintptr(j))))))) {
				v7 = *(**uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(j + 748))) + 276)))
				if *(*uint32)(unsafe.Add(unsafe.Pointer(v7), unsafe.Sizeof(uint32(0))*515)) != nox_player_netCode_85319C || !nox_common_getEngineFlag(nox_engine_flag(NOX_ENGINE_FLAG_DISABLE_GRAPHICS_RENDERING)) {
					v8 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v7), unsafe.Sizeof(uint32(0))*920)))
					if (v8&1) == 0 || v8&32 != 0 {
						v9 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v7), unsafe.Sizeof(uint32(0))*517)))
						if v9 != 0 {
							if v9 == *(*int32)(unsafe.Add(unsafe.Pointer(i), unsafe.Sizeof(int32(0))*8)) {
								nox_xxx_createAtImpl_4191D0(*(*uint8)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(i), unsafe.Sizeof(int32(0))*9)) + 57))), unsafe.Pointer(uintptr(j+48)), 1, int32(*(*uint32)(unsafe.Pointer(uintptr(j + 36)))), 0)
							}
						}
					}
				}
			}
		}
	}
	for k = int32(uintptr(unsafe.Pointer(nox_xxx_getFirstPlayerUnit_4DA7C0()))); k != 0; k = int32(uintptr(unsafe.Pointer(nox_xxx_getNextPlayerUnit_4DA7F0((*nox_object_t)(unsafe.Pointer(uintptr(k))))))) {
		if nox_xxx_servObjectHasTeam_419130(k+48) == 0 {
			v11 = int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(k + 748))) + 276))))
			if (*(*uint32)(unsafe.Pointer(uintptr(v11 + 2060))) != nox_player_netCode_85319C || !nox_common_getEngineFlag(nox_engine_flag(NOX_ENGINE_FLAG_DISABLE_GRAPHICS_RENDERING))) && (int32(*(*uint8)(unsafe.Pointer(uintptr(v11 + 3680))))&1) == 0 {
				nox_xxx_playerGoObserver_4E6860((*nox_playerInfo)(unsafe.Pointer(uintptr(v11))), 0, 0)
			}
		}
	}
	nox_xxx_SetGameplayFlag_417D50(4)
	return 1
}
func sub_4189D0() *byte {
	var (
		v0 *byte
		v1 uint8
		i  *byte
		v3 uint8
	)
	v0 = nil
	v1 = 32
	for i = (*byte)(unsafe.Pointer(nox_server_teamFirst_418B10())); i != nil; i = (*byte)(unsafe.Pointer(nox_server_teamNext_418B60((*nox_team_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(i))))))))) {
		v3 = uint8(int8(sub_418BC0(int32(uintptr(unsafe.Pointer(i))))))
		if int32(v3) < int32(v1) {
			v1 = v3
			v0 = i
		}
	}
	return v0
}
func sub_418A10() *byte {
	var result *byte
	result = (*byte)(unsafe.Pointer(nox_server_teamFirst_418B10()))
	if result == nil {
		return (*byte)(unsafe.Pointer(nox_xxx_teamCreate_4186D0(0)))
	}
	for *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(result))), unsafe.Sizeof(uint32(0))*15))) != 0 {
		result = (*byte)(unsafe.Pointer(nox_server_teamNext_418B60((*nox_team_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(result)))))))))
		if result == nil {
			return (*byte)(unsafe.Pointer(nox_xxx_teamCreate_4186D0(0)))
		}
	}
	return result
}
func sub_418A40(a1 *libc.WChar) *byte {
	var v1 *byte
	v1 = (*byte)(unsafe.Pointer(nox_server_teamFirst_418B10()))
	if v1 == nil {
		return nil
	}
	for _nox_wcsicmp((*libc.WChar)(unsafe.Pointer(v1)), a1) != 0 {
		v1 = (*byte)(unsafe.Pointer(nox_server_teamNext_418B60((*nox_team_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v1)))))))))
		if v1 == nil {
			return nil
		}
	}
	return v1
}
func sub_418A80(a1 int32) *byte {
	var result *byte
	result = (*byte)(unsafe.Pointer(nox_server_teamFirst_418B10()))
	if result == nil {
		return nil
	}
	for int32(uint8(*(*byte)(unsafe.Add(unsafe.Pointer(result), 56)))) != a1 {
		result = (*byte)(unsafe.Pointer(nox_server_teamNext_418B60((*nox_team_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(result)))))))))
		if result == nil {
			return nil
		}
	}
	return result
}
func sub_418BC0(a1 int32) int32 {
	var (
		v1 int32
		i  *byte
		v4 *uint32
	)
	v1 = 0
	if a1 == 0 {
		return 0
	}
	for i = (*byte)(unsafe.Pointer(nox_common_playerInfoGetFirst_416EA0())); i != nil; i = (*byte)(unsafe.Pointer(nox_common_playerInfoGetNext_416EE0((*nox_playerInfo)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(i))))))))) {
		v4 = nox_xxx_objGetTeamByNetCode_418C80(int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(i))), unsafe.Sizeof(uint32(0))*515)))))
		if v4 != nil {
			if nox_xxx_teamCompare2_419180(int32(uintptr(unsafe.Pointer(v4))), *(*uint8)(unsafe.Pointer(uintptr(a1 + 57)))) != 0 {
				v1++
			}
		}
	}
	return v1
}
func nox_xxx_teamCheckSmth_418C60(a1 int32) int32 {
	var result int32
	result = a1
	if a1 != 0 {
		result = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 44))))
	}
	return result
}
func sub_418C70(a1 *uint32) *uint32 {
	var result *uint32
	result = a1
	if a1 != nil {
		result = (*uint32)(unsafe.Pointer(uintptr(*a1)))
	}
	return result
}
func nox_xxx_objGetTeamByNetCode_418C80(a1 int32) *uint32 {
	var (
		v1 int32
		v3 *uint32
	)
	if noxflags.HasGame(noxflags.GameHost) {
		v1 = nox_server_getObjectFromNetCode_4ECCB0(a1)
		if v1 != 0 {
			return (*uint32)(unsafe.Pointer(uintptr(v1 + 48)))
		}
	} else {
		v3 = nox_xxx_netSpriteByCodeDynamic_45A6F0(a1)
		if v3 != nil {
			return (*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*6))
		}
	}
	return nil
}
func nox_xxx_teamRenameMB_418CD0(a1 *libc.WChar, a2 *libc.WChar) {
	var (
		v2 int32
		v3 [46]byte
	)
	if a1 != nil {
		sub_457010(int32(uintptr(unsafe.Pointer(a1))), a2)
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a1))), unsafe.Sizeof(uint32(0))*17))) = 0
		if noxflags.HasGame(noxflags.GameHost) {
			v2 = int32(*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(a1))), 57))))
			*(*uint16)(unsafe.Pointer(&v3[0])) = 1220
			*(*uint32)(unsafe.Pointer(&v3[2])) = uint32(v2)
			nox_wcscpy((*libc.WChar)(unsafe.Pointer(&v3[6])), a2)
			nox_xxx_netSendPacket1_4E5390(159, int32(uintptr(unsafe.Pointer(&v3[0]))), 46, 0, 1)
		}
		nox_wcscpy(a1, a2)
	}
}
func sub_418D80(a1 int32) {
	var (
		v1 int32
		i  *byte
		v3 *uint32
		v4 int32
		v5 [6]byte
	)
	if a1 != 0 && sub_418BC0(a1) > 0 {
		if noxflags.HasGame(noxflags.GameHost) {
			v1 = int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 57))))
			v5[0] = 196
			v5[1] = 5
			*(*uint32)(unsafe.Pointer(&v5[2])) = uint32(v1)
			nox_xxx_netSendPacket1_4E5390(159, int32(uintptr(unsafe.Pointer(&v5[0]))), 6, 0, 1)
		}
		for i = (*byte)(unsafe.Pointer(nox_common_playerInfoGetFirst_416EA0())); i != nil; i = (*byte)(unsafe.Pointer(nox_common_playerInfoGetNext_416EE0((*nox_playerInfo)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(i))))))))) {
			v3 = nox_xxx_objGetTeamByNetCode_418C80(int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(i))), unsafe.Sizeof(uint32(0))*515)))))
			v4 = int32(uintptr(unsafe.Pointer(v3)))
			if v3 != nil {
				if int32(*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(v3))), 4)))) == int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 57)))) {
					sub_4571A0(int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(i))), unsafe.Sizeof(uint32(0))*515)))), 0)
					sub_418E40(a1, v4)
				}
			}
		}
	}
}
func sub_418E40(a1 int32, a2 int32) *uint32 {
	var (
		result *uint32
		v3     *uint32
		v4     int32
		i      int32
	)
	result = (*uint32)(unsafe.Pointer(uintptr(a1)))
	if a1 != 0 && a2 != 0 {
		if *(*uint32)(unsafe.Pointer(uintptr(a1 + 44))) == uint32(a2) {
			*(*uint32)(unsafe.Pointer(uintptr(a1 + 44))) = *(*uint32)(unsafe.Pointer(uintptr(a2)))
		} else {
			result = (*uint32)(unsafe.Pointer(uintptr(nox_xxx_teamCheckSmth_418C60(a1))))
			if result == nil {
				return result
			}
			for *result != uint32(a2) {
				result = sub_418C70(result)
				if result == nil {
					return result
				}
			}
			*result = *(*uint32)(unsafe.Pointer(uintptr(a2)))
		}
		*(*uint32)(unsafe.Pointer(uintptr(a2))) = 0
		*(*uint8)(unsafe.Pointer(uintptr(a2 + 4))) = 0
		result = (*uint32)(unsafe.Pointer(nox_xxx_getFirstPlayerUnit_4DA7C0()))
		v3 = result
		if result != nil {
			for (*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*12)) != (*uint32)(unsafe.Pointer(uintptr(a2))) {
				result = (*uint32)(unsafe.Pointer(nox_xxx_getNextPlayerUnit_4DA7F0((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v3)))))))))
				v3 = result
				if result == nil {
					return result
				}
			}
			v4 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*187)))
			sub_4D97E0(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v4 + 276))) + 2064)))))
			sub_4E8110(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v4 + 276))) + 2064)))))
			result = (*uint32)(unsafe.Pointer(nox_xxx_monsterMarkUpdate_4E8020(int32(uintptr(unsafe.Pointer(v3))))))
			for i = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*129))); i != 0; i = int32(*(*uint32)(unsafe.Pointer(uintptr(i + 512)))) {
				if int32(*(*uint8)(unsafe.Pointer(uintptr(i + 8))))&6 != 0 {
					result = (*uint32)(unsafe.Pointer(nox_xxx_monsterMarkUpdate_4E8020(i)))
				}
			}
		}
	}
	return result
}
func sub_418F20(a1p *nox_team_t, a2 int32) {
	var (
		a1 *libc.WChar = (*libc.WChar)(unsafe.Pointer(a1p))
		v2 int32
		i  *byte
		v4 *uint32
		v5 [6]byte
		v6 [20]libc.WChar
	)
	if a1 != nil && nox_xxx_clientGetTeamColor_418AB0(int32(*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(a1))), 57))))) != nil {
		nox_wcscpy(&v6[0], a1)
		if noxflags.HasGame(noxflags.GameHost) && a2 != 0 {
			v2 = int32(*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(a1))), 57))))
			v5[0] = 196
			v5[1] = 6
			*(*uint32)(unsafe.Pointer(&v5[2])) = uint32(v2)
			nox_xxx_netSendPacket1_4E5390(159, int32(uintptr(unsafe.Pointer(&v5[0]))), 6, 0, 1)
		}
		for i = (*byte)(unsafe.Pointer(nox_common_playerInfoGetFirst_416EA0())); i != nil; i = (*byte)(unsafe.Pointer(nox_common_playerInfoGetNext_416EE0((*nox_playerInfo)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(i))))))))) {
			v4 = nox_xxx_objGetTeamByNetCode_418C80(int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(i))), unsafe.Sizeof(uint32(0))*515)))))
			if v4 != nil && int32(*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(v4))), 4)))) == int32(*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(a1))), 57)))) {
				sub_418E40(int32(uintptr(unsafe.Pointer(a1))), int32(uintptr(unsafe.Pointer(v4))))
			}
		}
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a1))), unsafe.Sizeof(uint32(0))*12))) = 0
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a1))), unsafe.Sizeof(uint32(0))*16))) = 0
		*a1 = 0
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a1))), unsafe.Sizeof(uint32(0))*17))) = 0
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a1))), unsafe.Sizeof(uint32(0))*11))) = 0
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a1))), unsafe.Sizeof(uint32(0))*18))) = 0
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a1))), unsafe.Sizeof(uint32(0))*19))) = 0
		*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(a1))), 57))) = 0
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a1))), unsafe.Sizeof(uint32(0))*15))) = 0
		*memmap.PtrUint8(6112660, 526280)--
		sub_459CD0()
		if noxflags.HasGame(noxflags.GameHost) {
			sub_456EA0(&v6[0])
		}
	}
}
func nox_xxx_netChangeTeamID_419090(a1 int32, a2 int32) {
	var (
		v2 int32
		v3 int32
		v4 [10]byte
	)
	if a1 != 0 {
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 52))) = uint32(a2)
		if noxflags.HasGame(noxflags.GameHost) {
			v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 52))))
			v3 = int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 57))))
			v4[0] = 196
			v4[1] = 8
			*(*uint32)(unsafe.Pointer(&v4[2])) = uint32(v3)
			*(*uint32)(unsafe.Pointer(&v4[6])) = uint32(v2)
			nox_xxx_netSendPacket1_4E5390(159, int32(uintptr(unsafe.Pointer(&v4[0]))), 10, 0, 1)
		}
	}
}
func sub_4190F0(a1 *libc.WChar) int32 {
	var v1 *byte
	v1 = (*byte)(unsafe.Pointer(nox_server_teamFirst_418B10()))
	if v1 == nil {
		return 0
	}
	for _nox_wcsicmp((*libc.WChar)(unsafe.Pointer(v1)), a1) != 0 {
		v1 = (*byte)(unsafe.Pointer(nox_server_teamNext_418B60((*nox_team_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v1)))))))))
		if v1 == nil {
			return 0
		}
	}
	return 1
}
func nox_xxx_servObjectHasTeam_419130(a1 int32) int32 {
	var result int32
	result = a1
	if a1 != 0 {
		result = bool2int(int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 4)))) != 0)
	}
	return result
}
func nox_xxx_servCompareTeams_419150(a1 int32, a2 int32) int32 {
	var (
		v2     int8
		v3     int8
		result int32
	)
	result = 0
	if a1 != 0 {
		if a2 != 0 {
			v2 = int8(*(*uint8)(unsafe.Pointer(uintptr(a1 + 4))))
			if int32(v2) != 0 {
				v3 = int8(*(*uint8)(unsafe.Pointer(uintptr(a2 + 4))))
				if int32(v3) != 0 {
					if int32(v2) == int32(v3) {
						result = 1
					}
				}
			}
		}
	}
	return result
}
func nox_xxx_teamCompare2_419180(a1 int32, a2 uint8) int32 {
	var (
		v2 *byte
		v3 *uint32
	)
	if a1 == 0 {
		return 0
	}
	if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 4)))) != int32(a2) {
		return 0
	}
	v2 = (*byte)(unsafe.Pointer(nox_xxx_clientGetTeamColor_418AB0(int32(a2))))
	if v2 == nil {
		return 0
	}
	v3 = (*uint32)(unsafe.Pointer(uintptr(nox_xxx_teamCheckSmth_418C60(int32(uintptr(unsafe.Pointer(v2)))))))
	if v3 == nil {
		return 0
	}
	for v3 != (*uint32)(unsafe.Pointer(uintptr(a1))) {
		v3 = sub_418C70(v3)
		if v3 == nil {
			return 0
		}
	}
	return 1
}
func nox_xxx_netChangeTeamMb_419570(a1 int32, a2 int32) {
	var (
		v2 *byte
		v3 *byte
		v4 [6]byte
	)
	if a1 != 0 {
		v2 = (*byte)(unsafe.Pointer(nox_xxx_clientGetTeamColor_418AB0(int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 4)))))))
		if v2 != nil {
			if nox_xxx_teamCompare2_419180(a1, *(*uint8)(unsafe.Pointer(uintptr(a1 + 4)))) != 0 {
				if noxflags.HasGame(noxflags.GameHost) && noxflags.HasGame(noxflags.GameOnline) {
					v3 = (*byte)(unsafe.Pointer(noxServer.getPlayerByID(a2)))
					if v3 != nil && nox_xxx_check_flag_aaa_43AF70() == 1 && !noxflags.HasGame(noxflags.GameModeChat) {
						sub_425ED0(int32(uintptr(unsafe.Pointer(v3))), 0)
					}
					sub_4571A0(a2, 0)
					v4[0] = 196
					v4[1] = 2
					*(*uint32)(unsafe.Pointer(&v4[2])) = uint32(a2)
					nox_xxx_netSendPacket1_4E5390(159, int32(uintptr(unsafe.Pointer(&v4[0]))), 6, 0, 1)
				}
				sub_418E40(int32(uintptr(unsafe.Pointer(v2))), a1)
				*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), unsafe.Sizeof(uint32(0))*12)))--
				if (sub_40A740() != 0 || noxflags.HasGame(noxflags.GameFlag16)) && sub_418BC0(int32(uintptr(unsafe.Pointer(v2)))) == 0 {
					if noxflags.HasGame(noxflags.GameModeCTF|noxflags.GameModeFlagBall) || noxflags.HasGame(noxflags.GameModeKOTR) && nox_xxx_CheckGameplayFlags_417DA0(4) {
						*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), unsafe.Sizeof(uint32(0))*15))) = 0
						sub_418800((*libc.WChar)(unsafe.Pointer(v2)), (*libc.WChar)(memmap.PtrOff(6112660, 527664)), 0)
					} else {
						sub_418F20((*nox_team_t)(unsafe.Pointer((*libc.WChar)(unsafe.Pointer(v2)))), 1)
					}
				}
			}
		}
	}
}
func sub_4196D0(a1 int32, a2 int32, a3 int32, a4 int32) int32 {
	var (
		v4 int32
		v5 *byte
		v7 [10]byte
	)
	if a1 == 0 || a2 == 0 || nox_xxx_teamCompare2_419180(a1, *(*uint8)(unsafe.Pointer(uintptr(a1 + 4)))) == 0 {
		return 0
	}
	noxServer.getPlayerByID(a3)
	if noxflags.HasGame(noxflags.GameHost) && noxflags.HasGame(noxflags.GameOnline) {
		v4 = int32(*(*uint8)(unsafe.Pointer(uintptr(a2 + 57))))
		v7[0] = 196
		v7[1] = 3
		*(*uint32)(unsafe.Pointer(&v7[2])) = uint32(v4)
		*(*uint16)(unsafe.Pointer(&v7[6])) = uint16(int16(a3))
		nox_xxx_netSendPacket1_4E5390(159, int32(uintptr(unsafe.Pointer(&v7[0]))), 10, 0, 1)
		sub_4571A0(a3, int32(*(*uint8)(unsafe.Pointer(uintptr(a2 + 57)))))
	}
	v5 = (*byte)(unsafe.Pointer(nox_xxx_clientGetTeamColor_418AB0(int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 4)))))))
	*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v5))), unsafe.Sizeof(uint32(0))*12)))--
	sub_418E40(int32(uintptr(unsafe.Pointer(v5))), a1)
	nox_xxx_createAtImpl_4191D0(*(*uint8)(unsafe.Pointer(uintptr(a2 + 57))), unsafe.Pointer(uintptr(a1)), 0, a3, a4)
	if uint32(a3) == nox_player_netCode_85319C {
		sub_455E70(*(*uint8)(unsafe.Pointer(uintptr(a2 + 57))))
	}
	return 1
}
func sub_4197C0(a1 *libc.WChar, a2 int32) {
	var (
		v2 int32
		v3 *uint32
		v4 [18]byte
	)
	if a1 != nil {
		*(*uint32)(unsafe.Pointer(&v4[6])) = *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a1))), unsafe.Sizeof(uint32(0))*15)))
		*(*uint32)(unsafe.Pointer(&v4[2])) = uint32(*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(a1))), 57))))
		*(*uint16)(unsafe.Pointer(&v4[0])) = 196
		*(*uint32)(unsafe.Pointer(&v4[10])) = *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a1))), unsafe.Sizeof(uint32(0))*13)))
		v4[14] = 0
		v4[16] = byte(*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(a1))), 56))))
		v4[17] = byte(*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(a1))), 68))))
		if noxflags.HasGame(noxflags.GameModeSolo10) {
			v4[14] = 1
		}
		v4[15] = byte(nox_wcslen(a1))
		v2 = int32(uint8(v4[15])) * 2
		v3 = (*uint32)(alloc.Calloc(1, int(v2+18)))
		*(*uint64)(unsafe.Pointer(v3)) = *(*uint64)(unsafe.Pointer(&v4[0]))
		*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*2)) = *(*uint32)(unsafe.Pointer(&v4[8]))
		*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*3)) = *(*uint32)(unsafe.Pointer(&v4[12]))
		*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v3))), unsafe.Sizeof(uint16(0))*8))) = *(*uint16)(unsafe.Pointer(&v4[16]))
		libc.MemCpy(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(v3))), 18), unsafe.Pointer(a1), int(int32(uint8(v4[15]))*2))
		nox_xxx_netSendPacket1_4E5390(a2, int32(uintptr(unsafe.Pointer(v3))), v2+18, 0, 1)
		alloc.Free(unsafe.Pointer(v3))
	}
}
func sub_4198A0(a1 int32, a2 int32, a3 int32) {
	var (
		v3 int32
		v4 int32
		v5 [10]byte
	)
	if a1 != 0 {
		v3 = nox_server_getObjectFromNetCode_4ECCB0(a3)
		if v3 != 0 {
			v5[0] = 196
			v4 = int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 4))))
			v5[1] = 1
			*(*uint16)(unsafe.Pointer(&v5[6])) = uint16(int16(a3))
			*(*uint32)(unsafe.Pointer(&v5[2])) = uint32(v4)
			*(*uint16)(unsafe.Pointer(&v5[8])) = *(*uint16)(unsafe.Pointer(uintptr(v3 + 4)))
			nox_xxx_netSendPacket1_4E5390(a2, int32(uintptr(unsafe.Pointer(&v5[0]))), 10, 0, 1)
		}
	}
}
func sub_419900(a1 int32, a2 int32, a3 int16) int8 {
	var (
		result int8
		v4     [10]byte
	)
	result = int8(a1)
	if a1 != 0 && a2 != 0 {
		result = int8(*(*uint8)(unsafe.Pointer(uintptr(a1 + 57))))
		if int32(*(*uint8)(unsafe.Pointer(uintptr(a2 + 4)))) != int32(result) {
			*(*uint32)(unsafe.Pointer(&v4[2])) = uint32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 57))))
			v4[0] = 196
			v4[1] = 10
			*(*uint16)(unsafe.Pointer(&v4[6])) = uint16(a3)
			result = int8(nox_xxx_netClientSend2_4E53C0(31, unsafe.Pointer(&v4[0]), 10, 0, 1))
		}
	}
	return result
}
func sub_419960(a1 int32, a2 int32, a3 int16) int8 {
	var (
		result int8
		v4     [10]byte
	)
	result = int8(a1)
	if a1 != 0 && a2 != 0 {
		result = int8(*(*uint8)(unsafe.Pointer(uintptr(a1 + 57))))
		if int32(*(*uint8)(unsafe.Pointer(uintptr(a2 + 4)))) != int32(result) {
			*(*uint32)(unsafe.Pointer(&v4[2])) = uint32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 57))))
			v4[0] = 196
			v4[1] = 11
			*(*uint16)(unsafe.Pointer(&v4[6])) = uint16(a3)
			result = int8(nox_xxx_netClientSend2_4E53C0(31, unsafe.Pointer(&v4[0]), 10, 0, 1))
		}
	}
	return result
}
func sub_4199C0(a1 int32) int32 {
	var v1 int32
	if a1 == 0 {
		return 0
	}
	v1 = int32(uintptr(unsafe.Pointer(nox_xxx_getFirstPlayerUnit_4DA7C0())))
	if v1 == 0 {
		return 0
	}
	for int32(*(*uint8)(unsafe.Pointer(uintptr(v1 + 52)))) != int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 57)))) || nox_xxx_unitIsCrown_4E7BE0(v1) == 0 {
		v1 = int32(uintptr(unsafe.Pointer(nox_xxx_getNextPlayerUnit_4DA7F0((*nox_object_t)(unsafe.Pointer(uintptr(v1)))))))
		if v1 == 0 {
			return 0
		}
	}
	return 1
}
func sub_419A10(a1 float32) float64 {
	*mem_getFloatPtr(6112660, 0x80D38) = a1
	**(**uint32)(memmap.PtrOff(0x587000, 55744)) &= math.MaxInt32
	return float64(*mem_getFloatPtr(6112660, 0x80D38))
}
func sub_419A30(a1 float32) uint32 {
	var result uint32
	if float64(a1) < 0.0 {
		return 0
	}
	*memmap.PtrUint32(6112660, 527668) = uint32(uintptr(memmap.PtrOff(6112660, 527676)))
	*mem_getFloatPtr(6112660, 0x80D3C) = float32(float64(a1) + 8.388608e+06)
	result = *memmap.PtrUint32(6112660, 527676) & 0x7FFFFF
	*memmap.PtrUint32(6112660, 527680) = *memmap.PtrUint32(6112660, 527676) & 0x7FFFFF
	return result
}
func nox_float2int(a1 float32) int32 {
	return int32(a1)
}
func nox_float2int16(a1 float32) int16 {
	return int16(int32(a1))
}
func nox_float2int16_abs(a1 float32) int16 {
	return int16(int32(math.Abs(float64(a1))))
}
func nox_float2int8(a1 float32) int8 {
	return int8(int32(a1))
}
func nox_double2float(a1 float64) float32 {
	return float32(a1)
}
func nox_double2int(a1 float64) int32 {
	return int32(a1)
}
func sub_419DE0(a1 int32, lpMem *unsafe.Pointer) {
	*lpMem = nil
	alloc.Free(unsafe.Pointer(lpMem))
}
func sub_419E10(a1 int32, a2 int32) int32 {
	var result int32
	result = a1
	if a1 != 0 && (int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 16))))&32) == 0 {
		result = 1 << int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))) + 276))) + 2064))))
		if a2 != 0 {
			dword_5d4594_527712 |= uint32(result)
		} else {
			result = ^result
			dword_5d4594_527712 &= uint32(result)
		}
	}
	return result
}
func sub_419E60(a1 int32) int32 {
	var result int32
	if a1 != 0 && int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))&4 != 0 {
		result = bool2int((dword_5d4594_527712 & uint32(1<<int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))) + 276))) + 2064)))))) != 0)
	} else {
		result = 0
	}
	return result
}
func sub_419EA0() int32 {
	return bool2int(dword_5d4594_527712 != 0)
}
func sub_419EB0(a1 int8, a2 int32) int32 {
	var result int32
	result = 1 << int32(a1)
	if a2 == 1 {
		*memmap.PtrUint32(6112660, 527716) |= uint32(result)
	} else {
		result = ^result
		*memmap.PtrUint32(6112660, 527716) &= uint32(result)
	}
	return result
}
func sub_419EE0(a1 int8) int32 {
	return bool2int((*memmap.PtrUint32(6112660, 527716) & uint32(1<<int32(a1))) != 0)
}
func sub_419F00() int32 {
	return bool2int(*memmap.PtrUint32(6112660, 527716) != 0)
}
func sub_419F10(a1 *byte, a2 *byte) int32 {
	var (
		v2 uint8
		v3 *byte
	)
	_ = v3
	var v4 *byte
	var result int32
	var v6 [1024]byte
	libc.StrCpy(&v6[0], a1)
	if libc.StrChr(&v6[0], 46) == nil {
		v2 = *memmap.PtrUint8(0x587000, 56164)
		v3 = &v6[libc.StrLen(&v6[0])]
		*(*uint32)(unsafe.Pointer(v3)) = *memmap.PtrUint32(0x587000, 56160)
		*(*byte)(unsafe.Add(unsafe.Pointer(v3), 4)) = byte(v2)
	}
	v4 = nox_fs_root()
	nox_sprintf((*byte)(memmap.PtrOff(0x85B3FC, 10984)), CString("%s\\Save\\%s"), v4, &v6[0])
	libc.StrCpy((*byte)(memmap.PtrOff(0x85B3FC, 12008)), a2)
	compatGetLocalTime((LPSYSTEMTIME)(memmap.PtrOff(0x85B3FC, 12168)))
	result = bool2int(noxflags.HasGame(noxflags.GameOnline))
	if result != 0 {
		result = int32(*memmap.PtrUint32(0x85B3FC, 10980))
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&result))), 0)) = uint8(int8(int32(*memmap.PtrUint8(0x85B3FC, 10980))&254 | 2))
		*memmap.PtrUint32(0x85B3FC, 10980) = uint32(result)
	} else {
		*memmap.PtrUint32(0x85B3FC, 10980) = *memmap.PtrUint32(0x85B3FC, 10980)&0xFFFFFFFD | 1
	}
	return result
}
func sub_41A000(a1 *byte, sv *nox_savegame_xxx) int32 {
	var (
		result int32
		v3     int32
		v4     *uint8
		v5     int32
		v6     uint32
		v7     int8
		v8     *uint8
		v9     *byte
		v10    int32
		v11    int32
		v12    [1276]byte
		v13    int16
	)
	sv.field_1028[0] = *memmap.PtrUint8(6112660, 527728)
	result = cryptFileOpen(a1, 1, 27)
	if result == 0 {
		return 0
	}
	for {
		cryptFileReadWrite((*uint8)(unsafe.Pointer(&v10))[:4])
		if v10 == 0 {
			break
		}
		cryptFileReadMaybeAlign((*uint8)(unsafe.Pointer(&v11))[:4])
		if v10 == 1 {
			libc.MemCpy(unsafe.Pointer(&v12[0]), memmap.PtrOff(0x85B3FC, 10980), int(unsafe.Sizeof([1276]byte{})))
			v3 = 0
			v13 = int16(*memmap.PtrUint16(0x85B3FC, 12256))
			if *memmap.PtrUint32(0x587000, 55936) != 0 {
				v4 = (*uint8)(memmap.PtrOff(0x587000, 55936))
				for unsafe.Pointer(v4) != memmap.PtrOff(0x587000, 55948) {
					v5 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v4))), unsafe.Sizeof(uint32(0))*3))))
					v4 = (*uint8)(unsafe.Add(unsafe.Pointer(v4), 12))
					v3++
					if v5 == 0 {
						goto LABEL_10
					}
				}
				if (*(*funcuint32(int32))(memmap.PtrOff(0x587000, 55956)))(0) == 0 {
					cryptFileClose()
					return 0
				}
			}
		LABEL_10:
			libc.MemCpy(unsafe.Pointer(sv), memmap.PtrOff(0x85B3FC, 10980), int(unsafe.Sizeof(nox_savegame_xxx{})))
			libc.MemCpy(memmap.PtrOff(0x85B3FC, 10980), unsafe.Pointer(&v12[0]), 1276)
			*memmap.PtrUint16(0x85B3FC, 12256) = uint16(v13)
		} else {
			nox_xxx_cryptSeekCur(v11)
		}
	}
	cryptFileClose()
	v6 = uint32(libc.StrLen(a1) + 1)
	v7 = int8(uint8(v6))
	v6 >>= 2
	libc.MemCpy(unsafe.Pointer(&sv.path[0]), unsafe.Pointer(a1), int(v6*4))
	v9 = (*byte)(unsafe.Add(unsafe.Pointer(a1), v6*4))
	v8 = (*uint8)(unsafe.Pointer(&sv.path[v6*4]))
	*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v6))), 0)) = uint8(v7)
	libc.MemCpy(unsafe.Pointer(v8), unsafe.Pointer(v9), int(v6&3))
	return 1
}
func nox_xxx_playerSaveToFile_41A140(a1 *byte, a2 int32) int32 {
	var (
		result int32
		v3     int32
		v4     int32
		v5     *uint8
		v7     int32
		v9     int32
		v10    int32
	)
	result = int32(uintptr(unsafe.Pointer(nox_common_playerInfoFromNum_417090(a2))))
	if result == 0 {
		return 0
	}
	v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(result + 2056))))
	if v3 == 0 {
		nox_xxx_networkLog_printf_413D30(CString("SaveServerPlayerData: NULL player object\n"))
		return 0
	}
	v4 = result + 2185
	if cryptFileOpen(a1, 0, 27) == 0 {
		nox_xxx_networkLog_printf_413D30(CString("SavePlayerData: Can't open file '%s'\n"), a1)
		return 0
	}
	v10 = 0
	if *memmap.PtrUint32(0x587000, 55816) == 0 {
		cryptFileClose()
		return 1
	}
	v5 = (*uint8)(memmap.PtrOff(0x587000, 55824))
	for {
		cryptFileReadWrite((*uint8)(unsafe.Add(unsafe.Pointer(v5), -4))[:4])
		nox_xxx_crypt_426C90()
		v7 = (cgoAsFunc(*(*uint32)(unsafe.Pointer(v5)), (*func(int32, int32, uint32, uint32) int32)(nil)))(v3, v4, 0, 0)
		nox_xxx_crypt_426D40()
		if v7 == 0 {
			break
		}
		v9 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v5))), unsafe.Sizeof(uint32(0))*1))))
		v5 = (*uint8)(unsafe.Add(unsafe.Pointer(v5), 12))
		v10++
		if v9 == 0 {
			cryptFileClose()
			return 1
		}
	}
	nox_xxx_networkLog_printf_413D30(CString("SavePlayerData: Error saving player data '%s'\n"), *memmap.PtrUint32(0x587000, uint32(v10*12+0xDA08)))
	cryptFileClose()
	return 0
}
func nox_xxx_mapSavePlayerDataMB_41A230(a1 *byte) int32 {
	var (
		v1     *byte
		result int32
		v3     *uint8
		v5     int32
		v7     int32
	)
	v1 = a1
	result = cryptFileOpen(a1, 2, 27)
	if result != 0 {
		if sub_45D9B0() != 0 {
			sub_45D870()
		}
		if (*byte)(memmap.PtrOff(0x85B3FC, 10984)) != v1 {
			libc.StrCpy((*byte)(memmap.PtrOff(0x85B3FC, 10984)), v1)
		}
		if *memmap.PtrUint32(0x587000, 55936) != 0 {
			v3 = (*uint8)(memmap.PtrOff(0x587000, 55944))
			for {
				cryptFileReadWrite((*uint8)(unsafe.Add(unsafe.Pointer(v3), -4))[:4])
				nox_xxx_crypt_426C90()
				v5 = (*(*funcuint32(int32))(unsafe.Pointer(v3)))(0)
				nox_xxx_crypt_426D40()
				if v5 == 0 {
					break
				}
				v7 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v3))), unsafe.Sizeof(uint32(0))*1))))
				v3 = (*uint8)(unsafe.Add(unsafe.Pointer(v3), 12))
				if v7 == 0 {
					goto LABEL_8
				}
			}
			cryptFileClose()
			result = 0
		} else {
		LABEL_8:
			a1 = nil
			cryptFileReadWrite((*uint8)(unsafe.Pointer(&a1))[:4])
			cryptFileClose()
			result = 1
		}
	}
	return result
}
func nox_xxx_cliPlrInfoLoadFromFile_41A2E0(a1 *byte, a2 int32) *byte {
	var (
		result *byte
		v3     int32
		v4     *byte
		v5     int32
		v6     *uint8
		v7     int32
		v8     int32
		v9     int32
		v10    uint16
		v11    uint16
	)
	result = (*byte)(unsafe.Pointer(nox_common_playerInfoFromNum_417090(a2)))
	if result != nil {
		v3 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(result))), unsafe.Sizeof(uint32(0))*514))))
		if v3 != 0 {
			v4 = (*byte)(unsafe.Add(unsafe.Pointer(result), 2185))
			result = (*byte)(unsafe.Pointer(uintptr(cryptFileOpen(a1, 1, 27))))
			if result != nil {
				if noxflags.HasGame(noxflags.GameModeCoop) {
					nox_xxx_set_god_4EF500(0)
					nox_xxx_set_sage(0)
				}
				*memmap.PtrUint16(6112660, 527696) = uint16(nox_xxx_unitGetHP_4EE780((*nox_object_t)(unsafe.Pointer(uintptr(v3)))))
				*memmap.PtrUint32(6112660, 527696) = uint32(*memmap.PtrUint16(6112660, 527696))
				*memmap.PtrUint16(6112660, 527700) = uint16(nox_xxx_unitGetOldMana_4EEC80(v3))
				*memmap.PtrUint32(6112660, 527700) = uint32(*memmap.PtrUint16(6112660, 527700))
				sub_4EFF10(v3)
				sub_419E10(v3, 1)
				for {
					for {
						cryptFileReadWrite((*uint8)(unsafe.Pointer(&a2))[:4])
						if a2 == 0 {
							cryptFileClose()
							sub_4EF140(v3)
							v10 = uint16(nox_xxx_unitGetMaxHP_4EE7A0(v3))
							nox_xxx_unitDamageClear_4EE5E0((*nox_object_t)(unsafe.Pointer(uintptr(v3))), int32(uint32(v10)-*memmap.PtrUint32(6112660, 527696)))
							v11 = uint16(nox_xxx_playerGetMaxMana_4EECB0(v3))
							nox_xxx_playerManaSub_4EEBF0(v3, int32(uint32(v11)-*memmap.PtrUint32(6112660, 527700)))
							nox_xxx_playerHP_4EE730(v3)
							sub_419E10(v3, 0)
							return (*byte)(unsafe.Pointer(uintptr(1)))
						}
						cryptFileReadMaybeAlign((*uint8)(unsafe.Pointer(&a1))[:4])
						v5 = 0
						if *memmap.PtrUint32(0x587000, 55816) != 0 {
							break
						}
					LABEL_13:
						nox_xxx_cryptSeekCur(int32(uintptr(unsafe.Pointer(a1))))
					}
					v6 = (*uint8)(memmap.PtrOff(0x587000, 55816))
					for uint32(a2) != *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v6))), unsafe.Sizeof(uint32(0))*1))) {
						v7 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v6))), unsafe.Sizeof(uint32(0))*3))))
						v6 = (*uint8)(unsafe.Add(unsafe.Pointer(v6), 12))
						v5++
						if v7 == 0 {
							goto LABEL_13
						}
					}
					if (cgoAsFunc(*(*uint32)(memmap.PtrOff(0x587000, uint32(v5*12+0xDA10))), (*func(int32, int32) int32)(nil)))(v3, int32(uintptr(unsafe.Pointer(v4)))) == 0 {
						break
					}
				}
				v8 = int32(*(*uint32)(unsafe.Pointer(uintptr(v3 + 504))))
				if v8 != 0 {
					for {
						v9 = int32(*(*uint32)(unsafe.Pointer(uintptr(v8 + 496))))
						nox_xxx_delayedDeleteObject_4E5CC0((*nox_object_t)(unsafe.Pointer(uintptr(v8))))
						v8 = v9
						if v9 == 0 {
							break
						}
					}
				}
				sub_419E10(v3, 0)
				cryptFileClose()
				result = nil
			}
		} else {
			result = nil
		}
	}
	return result
}
func nox_xxx_plrLoad_41A480(a1 *byte) int32 {
	var (
		result int32
		v2     int32
		v3     *uint8
		v4     int32
		v5     int32
		v6     int32
		v7     [1024]byte
	)
	libc.StrCpy(&v7[0], a1)
	result = cryptFileOpen(a1, 1, 27)
	if result != 0 {
		sub_4602F0()
		for {
			cryptFileReadWrite((*uint8)(unsafe.Pointer(&v5))[:4])
			if v5 == 0 {
				break
			}
			cryptFileReadMaybeAlign((*uint8)(unsafe.Pointer(&v6))[:4])
			v2 = 0
			if *memmap.PtrUint32(0x587000, 55936) != 0 {
				v3 = (*uint8)(memmap.PtrOff(0x587000, 55936))
				for uint32(v5) != *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v3))), unsafe.Sizeof(uint32(0))*1))) {
					v4 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v3))), unsafe.Sizeof(uint32(0))*3))))
					v3 = (*uint8)(unsafe.Add(unsafe.Pointer(v3), 12))
					v2++
					if v4 == 0 {
						goto LABEL_8
					}
				}
				if (cgoAsFunc(*(*uint32)(memmap.PtrOff(0x587000, uint32(v2*12+0xDA88))), (*func(unsafe.Pointer) int32)(nil)))(nil) == 0 {
					cryptFileClose()
					return 0
				}
			} else {
			LABEL_8:
				nox_xxx_cryptSeekCur(v6)
			}
		}
		cryptFileClose()
		result = 1
		libc.StrCpy((*byte)(memmap.PtrOff(0x85B3FC, 10984)), &v7[0])
	}
	return result
}
func sub_41A590(a1 int32, a2 int32) int32 {
	var (
		v2  int32
		v3  int32
		v4  int32
		v6  uint32
		v7  int32
		v8  int32
		v9  *byte
		v10 int32
		v11 int32
		v12 float32
		v13 int32
		v14 int32
		v15 int32
	)
	v2 = a2
	v3 = 0
	v4 = 0
	v15 = 0
	if a2 == 0 {
		return 0
	}
	if a1 != 0 {
		v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
		if v3 != 0 {
			v15 = int32(*(*uint32)(unsafe.Pointer(uintptr(v3 + 276))))
			v4 = v15
		}
	}
	a2 = 5
	cryptFileReadWrite((*uint8)(unsafe.Pointer(&a2))[:2])
	if int32(int16(a2)) > 5 {
		return 0
	}
	if int32(int16(a2)) >= 5 {
		if noxflags.HasGame(noxflags.GameModeQuest) {
			v13 = 4
		} else {
			v13 = 2 - bool2int(noxflags.HasGame(noxflags.GameModeCoop))
		}
		cryptFileReadWrite((*uint8)(unsafe.Pointer(&v13))[:4])
		if v4 != 0 {
			if int32(*(*uint8)(unsafe.Pointer(uintptr(v4 + 2064)))) != 31 {
				if noxflags.HasGame(noxflags.GameModeQuest) {
					if v13 != 4 {
						nox_xxx_playerCallDisconnect_4DEAB0(int32(*(*uint8)(unsafe.Pointer(uintptr(v4 + 2064)))), 1)
						return 0
					}
				} else if v13 != 4 {
					goto LABEL_20
				}
				if !noxflags.HasGame(noxflags.GameModeQuest) {
					nox_xxx_playerCallDisconnect_4DEAB0(int32(*(*uint8)(unsafe.Pointer(uintptr(v4 + 2064)))), 1)
					return 0
				}
			}
		}
	}
LABEL_20:
	*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v14))), 0)) = uint8(nox_wcslen((*libc.WChar)(unsafe.Pointer(uintptr(v2)))))
	cryptFileReadWrite((*uint8)(unsafe.Pointer(&v14))[:1])
	if int32(uint8(int8(v14))) >= 25 {
		return 0
	}
	cryptFileReadWrite((*uint8)(unsafe.Pointer(uintptr(v2)))[:uint32(int32(uint8(int8(v14)))*2)])
	*(*uint16)(unsafe.Pointer(uintptr(v2 + int32(uint8(int8(v14)))*2))) = 0
	if v3 != 0 {
		nox_wcscpy((*libc.WChar)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v3 + 276)))+4704))), (*libc.WChar)(unsafe.Pointer(uintptr(v2))))
	}
	if nox_xxx_cryptGetXxx() == 1 && v3 != 0 {
		v6 = nox_wcslen((*libc.WChar)(unsafe.Pointer(uintptr(v2))))
		v7 = nox_xxx_protectionStringCRCLen_56FAE0((*int32)(unsafe.Pointer(uintptr(v2))), v6*2)
		nox_xxx_playerResetProtectionCRC_56F7D0(int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v3 + 276))) + 4628)))), v7)
	}
	cryptFileReadWrite((*uint8)(unsafe.Pointer(uintptr(v2 + 50)))[:4])
	cryptFileReadWrite((*uint8)(unsafe.Pointer(uintptr(v2 + 54)))[:4])
	if nox_xxx_cryptGetXxx() == 1 && v3 != 0 {
		sub_56F780(int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v3 + 276))) + 4624)))), int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 54)))))
		sub_56F780(int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v3 + 276))) + 4620)))), int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 50)))))
	}
	cryptFileReadWrite((*uint8)(unsafe.Pointer(uintptr(v2 + 58)))[:4])
	cryptFileReadWrite((*uint8)(unsafe.Pointer(uintptr(v2 + 62)))[:4])
	cryptFileReadWrite((*uint8)(unsafe.Pointer(uintptr(v2 + 66)))[:1])
	if nox_xxx_cryptGetXxx() == 1 && v3 != 0 {
		sub_56F820(int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v3 + 276))) + 4616)))), *(*uint8)(unsafe.Pointer(uintptr(v2 + 66))))
	}
	cryptFileReadWrite((*uint8)(unsafe.Pointer(uintptr(v2 + 67)))[:1])
	cryptFileReadWrite((*uint8)(unsafe.Pointer(uintptr(v2 + 68)))[:3])
	cryptFileReadWrite((*uint8)(unsafe.Pointer(uintptr(v2 + 71)))[:3])
	cryptFileReadWrite((*uint8)(unsafe.Pointer(uintptr(v2 + 74)))[:3])
	cryptFileReadWrite((*uint8)(unsafe.Pointer(uintptr(v2 + 77)))[:3])
	cryptFileReadWrite((*uint8)(unsafe.Pointer(uintptr(v2 + 80)))[:3])
	if int32(int16(a2)) >= 2 {
		cryptFileReadWrite((*uint8)(unsafe.Pointer(uintptr(v2 + 83)))[:1])
		cryptFileReadWrite((*uint8)(unsafe.Pointer(uintptr(v2 + 84)))[:1])
		cryptFileReadWrite((*uint8)(unsafe.Pointer(uintptr(v2 + 85)))[:1])
		cryptFileReadWrite((*uint8)(unsafe.Pointer(uintptr(v2 + 86)))[:1])
		cryptFileReadWrite((*uint8)(unsafe.Pointer(uintptr(v2 + 87)))[:1])
	}
	v8 = a1
	if nox_xxx_cryptGetXxx() == 1 {
		if a1 != 0 {
			v9 = (*byte)(unsafe.Pointer(noxServer.getPlayerByID(int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 36)))))))
			if v9 != nil {
				nox_xxx_playerInitColors_461460((*nox_playerInfo)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v9)))))))
			}
		}
	}
	cryptFileReadWrite((*uint8)(unsafe.Pointer(uintptr(v2 + 88)))[:1])
	if int32(int16(a2)) >= 3 {
		v13 = 0
		if v3 != 0 {
			v13 = int32(*(*uint32)(unsafe.Pointer(uintptr(v3 + 320))))
		}
		cryptFileReadWrite((*uint8)(unsafe.Pointer(&v13))[:4])
		if nox_xxx_cryptGetXxx() == 1 && v3 != 0 {
			*(*uint32)(unsafe.Pointer(uintptr(v3 + 320))) = uint32(v13)
		}
		v12 = float32(nox_xxx_gamedataGetFloat_419D40(CString("MaxExtraLives")))
		v10 = int(v12)
		if v3 != 0 && *(*uint32)(unsafe.Pointer(uintptr(v3 + 320))) > uint32(v10) {
			return 0
		}
		if int32(uint16(int16(a2))) == 3 {
			cryptFileReadWrite((*uint8)(unsafe.Pointer(&a1))[:4])
			cryptFileReadWrite((*uint8)(unsafe.Pointer(&a1))[:4])
			cryptFileReadWrite((*uint8)(unsafe.Pointer(&a1))[:4])
			cryptFileReadWrite((*uint8)(unsafe.Pointer(&a1))[:4])
			cryptFileReadWrite((*uint8)(unsafe.Pointer(&a1))[:4])
			cryptFileReadWrite((*uint8)(unsafe.Pointer(&a1))[:4])
			cryptFileReadWrite((*uint8)(unsafe.Pointer(&a1))[:4])
			cryptFileReadWrite((*uint8)(unsafe.Pointer(&a1))[:4])
			cryptFileReadWrite((*uint8)(unsafe.Pointer(&a1))[:4])
		}
	}
	if nox_xxx_cryptGetXxx() == 1 {
		sub_4D6000((*nox_object_t)(unsafe.Pointer(uintptr(v8))))
	}
	if int32(int16(a2)) >= 4 {
		v11 = v15
		a1 = 0
		if v15 != 0 {
			a1 = int32(*(*uint32)(unsafe.Pointer(uintptr(v15 + 4696))))
		}
		cryptFileReadWrite((*uint8)(unsafe.Pointer(&a1))[:4])
		if nox_xxx_cryptGetXxx() == 1 {
			if v11 == 0 || (func() bool {
				*(*uint32)(unsafe.Pointer(uintptr(v11 + 4696))) = uint32(a1)
				return nox_xxx_cryptGetXxx() == 1
			}()) {
				if v3 != 0 {
					sub_4D7450(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v3 + 276))) + 2064)))), int16(uint16(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v3 + 276))) + 4696))))))
				}
			}
		}
	}
	return 1
}
func sub_41AA30(a1 int32) *byte {
	var (
		v1     int32
		v2     int32
		result *byte
		v4     int32
		v5     int32
		v6     int32
		v7     int32
		v8     int32
	)
	v1 = a1
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
	result = (*byte)(unsafe.Pointer(noxServer.getPlayerByID(int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 36)))))))
	*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a1))), 0)) = 1
	if result != nil {
		v7 = 2
		cryptFileReadWrite((*uint8)(unsafe.Pointer(&v7))[:2])
		if int32(int16(v7)) > 2 {
			return nil
		}
		if !noxflags.HasGame(noxflags.GameModeCoop) {
			*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a1))), 0)) = 0
		}
		cryptFileReadWrite((*uint8)(unsafe.Pointer(&a1))[:1])
		if int32(uint8(int8(a1))) != 0 {
			result = (*byte)(unsafe.Pointer(uintptr(bool2int(noxflags.HasGame(noxflags.GameModeCoop)))))
			if result == nil {
				return result
			}
			*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v4))), unsafe.Sizeof(uint16(0))*0)) = uint16(nox_xxx_unitGetMaxHP_4EE7A0(v1))
			v6 = v4
			cryptFileReadWrite((*uint8)(unsafe.Pointer(&v6))[:2])
			if nox_xxx_cryptGetXxx() == 1 {
				nox_xxx_unitSetMaxHP_4EE7C0(v1, int16(v6))
				nox_xxx_unitSetHP_4E4560((*nox_object_t)(unsafe.Pointer(uintptr(v1))), uint16(int16(v6)))
			}
			*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint16(0))*0)) = uint16(nox_xxx_playerGetMaxMana_4EECB0(v1))
			v6 = v5
			cryptFileReadWrite((*uint8)(unsafe.Pointer(&v6))[:2])
			if nox_xxx_cryptGetXxx() == 1 {
				nox_xxx_playerSetMaxMana_4EECD0(v1, int16(v6))
				nox_xxx_playerManaRefresh_4EECF0(v1)
			}
			*memmap.PtrUint32(6112660, 527696) = uint32(**(**uint16)(unsafe.Pointer(uintptr(v1 + 556))))
			cryptFileReadWrite((*uint8)(memmap.PtrOff(6112660, 527696))[:2])
			*memmap.PtrUint32(6112660, 527700) = uint32(*(*uint16)(unsafe.Pointer(uintptr(v2 + 4))))
			cryptFileReadWrite((*uint8)(memmap.PtrOff(6112660, 527700))[:2])
			*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v8))), 0)) = *(*uint8)(unsafe.Pointer(uintptr(v1 + 540)))
			cryptFileReadWrite((*uint8)(unsafe.Pointer(&v8))[:1])
			if nox_xxx_cryptGetXxx() == 1 {
				nox_xxx_setSomePoisonData_4EEA90(v1, int32(uint8(int8(v8))))
			}
			cryptFileReadWrite((*uint8)(unsafe.Pointer(uintptr(v1 + 541)))[:1])
			cryptFileReadWrite((*uint8)(unsafe.Pointer(uintptr(v1 + 542)))[:2])
			cryptFileReadWrite((*uint8)(unsafe.Pointer(uintptr(v1 + 28)))[:4])
			if nox_xxx_cryptGetXxx() == 1 {
				sub_56F8C0(int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 276))) + 4604)))), *(*float32)(unsafe.Pointer(uintptr(v1 + 28))))
				sub_4D81A0(v1)
			}
			if int32(int16(v7)) >= 2 {
				cryptFileReadWrite((*uint8)(unsafe.Pointer(uintptr(v1 + 124)))[:2])
				if nox_xxx_cryptGetXxx() == 1 {
					*(*uint16)(unsafe.Pointer(uintptr(v1 + 126))) = *(*uint16)(unsafe.Pointer(uintptr(v1 + 124)))
				}
			}
		}
		result = (*byte)(unsafe.Pointer(uintptr(1)))
	}
	return result
}
func sub_41AC30(a1 *uint32) int32 {
	var (
		v1     *uint32
		result int32
		v3     int32
		v4     int32
		v5     int32
		v6     int32
		v7     int32
		v8     int32
		v9     *byte
		v10    int32
		v11    *int32
		v12    int32
		v13    int32
		v14    *uint8
		v15    int32
		v16    *uint8
		v17    int32
		m      int32
		v19    int32
		v20    int32
		v21    int32
		v22    int32
		v23    int32
		v24    *uint32
		v25    *uint32
		v26    int32
		v27    int32
		i      int32
		v29    *uint32
		v30    *uint32
		j      *uint32
		v32    *uint8
	)
	_ = v32
	var v33 uint32
	var v34 uint32
	var l int8
	var v36 int8
	var v37 int32
	var k int32
	var v39 int32
	var v40 int32
	var v41 int32
	var v42 int32
	var v43 int32
	var v44 int32
	var v45 int32
	var v46 [256]byte
	v1 = a1
	v36 = 1
	v40 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*187)))
	if *memmap.PtrUint32(6112660, 527704) == 0 {
		*memmap.PtrUint32(6112660, 527704) = uint32(nox_xxx_getNameId_4E3AA0(CString("Glyph")))
	}
	if nox_xxx_cryptGetXxx() == 1 {
		sub_4EF140(int32(uintptr(unsafe.Pointer(a1))))
	}
	v42 = 3
	cryptFileReadWrite((*uint8)(unsafe.Pointer(&v42))[:2])
	if int32(int16(v42)) > 3 {
		return 0
	}
	if noxflags.HasGame(noxflags.GameOnline) && !noxflags.HasGame(noxflags.GameModeQuest) {
		v36 = 0
	}
	cryptFileReadWrite((*uint8)(unsafe.Pointer(&v36))[:1])
	if int32(v36) == 0 {
		goto LABEL_115
	}
	if !noxflags.HasGame(noxflags.GameModeCoop) {
		result = bool2int(noxflags.HasGame(noxflags.GameModeQuest))
		if result == 0 {
			return result
		}
	}
	v44 = int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v40 + 276))) + 2164))))
	cryptFileReadWrite((*uint8)(unsafe.Pointer(&v44))[:4])
	if nox_xxx_cryptGetXxx() == 1 {
		v3 = nox_xxx_playerGetGold_4FA6B0(int32(uintptr(unsafe.Pointer(a1))))
		nox_xxx_playerSubGold_4FA5D0(int32(uintptr(unsafe.Pointer(a1))), uint32(v3))
		nox_xxx_playerAddGold_4FA590(int32(uintptr(unsafe.Pointer(a1))), v44)
	}
	if nox_xxx_cryptGetXxx() != 0 {
		v21 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*126)))
		if v21 != 0 {
			for {
				v22 = int32(*(*uint32)(unsafe.Pointer(uintptr(v21 + 496))))
				nox_xxx_delayedDeleteObject_4E5CC0((*nox_object_t)(unsafe.Pointer(uintptr(v21))))
				v21 = v22
				if v22 == 0 {
					break
				}
			}
		}
		cryptFileReadWrite((*uint8)(unsafe.Pointer(&k))[:4])
		if noxflags.HasGame(noxflags.GameModeQuest) && k > 2560 {
			return 0
		}
		v23 = 0
		if k > 0 {
			for {
				cryptFileReadWrite((*uint8)(unsafe.Pointer(&v39))[:1])
				cryptFileReadWrite((*uint8)(unsafe.Pointer(&v46[0]))[:uint32(uint8(int8(v39)))])
				v46[uint8(int8(v39))] = 0
				v24 = (*uint32)(unsafe.Pointer(nox_xxx_newObjectByTypeID_4E3810(&v46[0])))
				v25 = v24
				if v24 == nil {
					return 0
				}
				if (cgoAsFunc(*(*uint32)(unsafe.Add(unsafe.Pointer(v24), unsafe.Sizeof(uint32(0))*176)), (*func(*uint32, uint32) int32)(nil)).(func(*uint32, uint32) int32))(v24, 0) == 0 {
					return 0
				}
				*(*uint32)(unsafe.Add(unsafe.Pointer(v25), unsafe.Sizeof(uint32(0))*14)) = 0x45380000
				*(*uint32)(unsafe.Add(unsafe.Pointer(v25), unsafe.Sizeof(uint32(0))*15)) = 0x45380000
				if noxflags.HasGame(noxflags.GameModeQuest) {
					if sub_4F2590(int32(uintptr(unsafe.Pointer(v25)))) == 0 {
						return 0
					}
				}
				nox_xxx_servMapLoadPlaceObj_4F3F50((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v25)))))), int32(uintptr(unsafe.Pointer(a1))), nil)
				nox_xxx_unitsNewAddToList_4DAC00()
				if nox_xxx_inventoryServPlace_4F36F0(int32(uintptr(unsafe.Pointer(a1))), int32(uintptr(unsafe.Pointer(v25))), 1, 1) == 0 {
					if !noxflags.HasGame(noxflags.GameModeQuest) {
						return 0
					}
					nox_xxx_delayedDeleteObject_4E5CC0((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v25)))))))
				}
				v26 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v25), unsafe.Sizeof(uint32(0))*4)))
				if (v26&32) == 0 && v26&256 != 0 {
					nox_xxx_playerTryDequip_4F2FB0(a1, (*nox_object_t)(unsafe.Pointer(v25)))
				}
				if func() int32 {
					p := &v23
					*p++
					return *p
				}() >= k {
					break
				}
			}
		}
		cryptFileReadWrite((*uint8)(unsafe.Pointer(&v41))[:1])
		v27 = 0
		if int32(uint8(int8(v41))) != 0 {
			for {
				cryptFileReadWrite((*uint8)(unsafe.Pointer(&v37))[:4])
				for i = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*126))); i != 0; i = int32(*(*uint32)(unsafe.Pointer(uintptr(i + 496)))) {
					if *(*uint32)(unsafe.Pointer(uintptr(i + 44))) == uint32(v37) {
						nox_xxx_playerTryEquip_4F2F70(a1, (*nox_object_t)(unsafe.Pointer(uintptr(i))))
					}
				}
				v27++
				if v27 >= int32(uint8(int8(v41))) {
					break
				}
			}
		}
		if noxflags.HasGame(noxflags.GameModeCoop) {
			sub_467750(0, 0)
			sub_467740(0)
		}
		cryptFileReadWrite((*uint8)(unsafe.Pointer(&v37))[:4])
		if v37 != 0 {
			v29 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*126)))))
			if v29 != nil {
				for *(*uint32)(unsafe.Add(unsafe.Pointer(v29), unsafe.Sizeof(uint32(0))*11)) != uint32(v37) {
					v29 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v29), unsafe.Sizeof(uint32(0))*124)))))
					if v29 == nil {
						goto LABEL_99
					}
				}
				nox_xxx_netSendSecondaryWeapon_4D9670(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v40 + 276))) + 2064)))), v29, 0)
			}
		}
	LABEL_99:
		if int32(int16(v42)) >= 2 {
			cryptFileReadWrite((*uint8)(unsafe.Pointer(&v37))[:4])
			if v37 != 0 {
				v30 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*126)))))
				if v30 != nil {
					for *(*uint32)(unsafe.Add(unsafe.Pointer(v30), unsafe.Sizeof(uint32(0))*11)) != uint32(v37) {
						v30 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v30), unsafe.Sizeof(uint32(0))*124)))))
						if v30 == nil {
							goto LABEL_106
						}
					}
					nox_xxx_netMsgLastQuiver_4D96B0(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v40 + 276))) + 2064)))), v30)
				}
			}
		}
	LABEL_106:
		if noxflags.HasGame(noxflags.GameModeQuest) {
			for j = (*uint32)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1))).FirstItem()))); j != nil; j = (*uint32)(unsafe.Pointer(uintptr(nox_xxx_inventoryGetNext_4E7990(int32(uintptr(unsafe.Pointer(j))))))) {
				*(*uint32)(unsafe.Add(unsafe.Pointer(j), unsafe.Sizeof(uint32(0))*11)) = nox_server_NextObjectScriptID()
				*(*uint32)(unsafe.Add(unsafe.Pointer(j), unsafe.Sizeof(uint32(0))*10)) = *(*uint32)(unsafe.Add(unsafe.Pointer(j), unsafe.Sizeof(uint32(0))*9))
			}
		}
		goto LABEL_109
	}
	v4 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*126)))
	v5 = 0
	for k = 0; v4 != 0; v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(v4 + 496)))) {
		if !noxflags.HasGame(noxflags.GameModeQuest) || sub_41B3E0(v4) != 0 {
			k++
		}
	}
	cryptFileReadWrite((*uint8)(unsafe.Pointer(&k))[:4])
	if noxflags.HasGame(noxflags.GameModeCoop) {
		v5 = sub_41B3B0()
	}
	if k != v5 || !noxflags.HasGame(noxflags.GameModeCoop) {
		v15 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*126)))
		if v15 != 0 {
			for {
				if !noxflags.HasGame(noxflags.GameModeQuest) || sub_41B3E0(v15) != 0 {
					*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v39))), 0)) = uint8(int8(libc.StrLen((*byte)(unsafe.Pointer(uintptr(nox_xxx_getUnitName_4E39D0(v15)))))))
					cryptFileReadWrite((*uint8)(unsafe.Pointer(&v39))[:1])
					v34 = uint32(uint8(int8(v39)))
					v16 = (*uint8)(unsafe.Pointer(uintptr(nox_xxx_getUnitName_4E39D0(v15))))
					cryptFileReadWrite(v16[:v34])
					if (cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(v15 + 704))), (*func(int32, uint32) int32)(nil)))(v15, 0) == 0 {
						return 0
					}
				}
				v15 = int32(*(*uint32)(unsafe.Pointer(uintptr(v15 + 496))))
				if v15 == 0 {
					goto LABEL_42
				}
			}
		}
	} else {
		v6 = 0
		v41 = 0
		for 2 != 0 {
			v7 = 0
			v43 = 0
			for {
				v8 = sub_467810(v7, v6)
				v45 = v8
				if v8 != 0 {
					v9 = sub_467870(v7, v6)
					v10 = 0
					if v8 > 0 {
						v11 = (*int32)(unsafe.Pointer(v9))
						for {
							v12 = nox_xxx_equipedItemByCode_4F7920(int32(uintptr(unsafe.Pointer(a1))), *v11)
							v13 = v12
							if v12 == 0 {
								return 0
							}
							*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v39))), 0)) = uint8(int8(libc.StrLen((*byte)(unsafe.Pointer(uintptr(nox_xxx_getUnitName_4E39D0(v12)))))))
							cryptFileReadWrite((*uint8)(unsafe.Pointer(&v39))[:1])
							v33 = uint32(uint8(int8(v39)))
							v14 = (*uint8)(unsafe.Pointer(uintptr(nox_xxx_getUnitName_4E39D0(v13))))
							cryptFileReadWrite(v14[:v33])
							if (cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(v13 + 704))), (*func(int32, uint32) int32)(nil)))(v13, 0) == 0 {
								return 0
							}
							v10++
							v11 = (*int32)(unsafe.Add(unsafe.Pointer(v11), unsafe.Sizeof(int32(0))*1))
							if v10 >= v45 {
								v6 = v41
								v7 = v43
								break
							}
						}
					}
				}
				v43 = func() int32 {
					p := &v7
					*p++
					return *p
				}()
				if v7 >= 4 {
					break
				}
			}
			v41 = func() int32 {
				p := &v6
				*p++
				return *p
			}()
			if v6 < 20 {
				continue
			}
			break
		}
	LABEL_42:
		v1 = a1
	}
	v17 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*126)))
	for l = 0; v17 != 0; v17 = int32(*(*uint32)(unsafe.Pointer(uintptr(v17 + 496)))) {
		if *(*uint32)(unsafe.Pointer(uintptr(v17 + 16)))&256 != 0 {
			l++
		}
	}
	cryptFileReadWrite((*uint8)(unsafe.Pointer(&l))[:1])
	for m = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*126))); m != 0; m = int32(*(*uint32)(unsafe.Pointer(uintptr(m + 496)))) {
		if *(*uint32)(unsafe.Pointer(uintptr(m + 16)))&256 != 0 {
			cryptFileReadWrite((*uint8)(unsafe.Pointer(uintptr(m + 44)))[:4])
		}
	}
	v19 = sub_4678B0()
	v37 = v19
	if v19 == 0 {
	LABEL_57:
		if m != 0 {
			goto LABEL_59
		}
		goto LABEL_58
	}
	m = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*126)))
	if m != 0 {
		for *(*uint32)(unsafe.Pointer(uintptr(m + 36))) != uint32(v19) {
			m = int32(*(*uint32)(unsafe.Pointer(uintptr(m + 496))))
			if m == 0 {
				goto LABEL_58
			}
		}
		v37 = int32(*(*uint32)(unsafe.Pointer(uintptr(m + 44))))
		goto LABEL_57
	}
LABEL_58:
	v37 = 0
LABEL_59:
	cryptFileReadWrite((*uint8)(unsafe.Pointer(&v37))[:4])
	v20 = sub_4678C0()
	v37 = v20
	if v20 != 0 {
		m = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*126)))
		if m == 0 {
		LABEL_66:
			v37 = 0
			goto LABEL_67
		}
		for *(*uint32)(unsafe.Pointer(uintptr(m + 36))) != uint32(v20) {
			m = int32(*(*uint32)(unsafe.Pointer(uintptr(m + 496))))
			if m == 0 {
				goto LABEL_66
			}
		}
		v37 = int32(*(*uint32)(unsafe.Pointer(uintptr(m + 44))))
	}
	if m == 0 {
		goto LABEL_66
	}
LABEL_67:
	cryptFileReadWrite((*uint8)(unsafe.Pointer(&v37))[:4])
LABEL_109:
	v32 = (*uint8)(unsafe.Pointer(uintptr(v40 + 244)))
	if int32(int16(v42)) < 3 {
		*(*uint8)(unsafe.Pointer(uintptr(v40 + 244))) = 0
	} else {
		cryptFileReadWrite((*uint8)(unsafe.Pointer(uintptr(v40 + 244)))[:1])
	}
	if nox_xxx_cryptGetXxx() == 1 && noxflags.HasGame(noxflags.GameModeQuest) {
		*v32 = 0
	}
LABEL_115:
	if noxflags.HasGame(noxflags.GameModeQuest) && sub_4F2C30(int32(uintptr(unsafe.Pointer(v1)))) == 0 {
		return 0
	}
	nox_xxx_netMsgInventoryLoaded_4D96E0(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v40 + 276))) + 2064)))))
	return 1
}
func sub_41B3B0() int32 {
	var (
		v0 int32
		i  int32
		j  int32
	)
	v0 = 0
	for i = 0; i < 20; i++ {
		for j = 0; j < 4; j++ {
			v0 += sub_467810(j, i)
		}
	}
	return v0
}
func sub_41B3E0(a1 int32) int32 {
	var (
		v1     int32
		result int32
	)
	v1 = int32(*memmap.PtrUint32(6112660, 527724))
	if *memmap.PtrUint32(6112660, 527724) == 0 {
		v1 = nox_xxx_getNameId_4E3AA0(CString("Glyph"))
		*memmap.PtrUint32(6112660, 527724) = uint32(v1)
	}
	if (int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8)))) & 64) == 64 {
		result = 0
	} else {
		result = bool2int(int32(*(*uint16)(unsafe.Pointer(uintptr(a1 + 4)))) != v1)
	}
	return result
}
func nox_xxx_guiFieldbook_41B420(a1 int32) int32 {
	var (
		v1     *byte
		result int32
		v3     *uint32
		v4     *uint32
		v5     int32
		i      int32
		v7     *byte
		v8     int32
		v9     int32
		v10    int32
		v11    uint32
		v12    int8
		v13    int32
		v14    int32
		v15    int32
		v16    [256]byte
	)
	v1 = (*byte)(unsafe.Pointer(noxServer.getPlayerByID(int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 36)))))))
	v12 = 1
	if v1 == nil {
		return 0
	}
	v15 = 1
	cryptFileReadWrite((*uint8)(unsafe.Pointer(&v15))[:2])
	if int32(int16(v15)) > 1 {
		return 0
	}
	if noxflags.HasGame(noxflags.GameOnline) && !noxflags.HasGame(noxflags.GameModeQuest) {
		v12 = 0
	}
	cryptFileReadWrite((*uint8)(unsafe.Pointer(&v12))[:1])
	if int32(v12) == 0 {
		return 1
	}
	if !noxflags.HasGame(noxflags.GameModeCoop) {
		result = bool2int(noxflags.HasGame(noxflags.GameModeQuest))
		if result == 0 {
			return result
		}
	}
	if nox_xxx_cryptGetXxx() != 0 {
		cryptFileReadWrite((*uint8)(unsafe.Pointer(&v13))[:1])
		if int32(uint8(int8(v13))) > 41 {
			return 0
		}
		v8 = 0
		if int32(uint8(int8(v13))) > 0 {
			for {
				cryptFileReadWrite((*uint8)(unsafe.Pointer(&v14))[:1])
				if int32(uint8(int8(v14))) >= 256 {
					break
				}
				cryptFileReadWrite((*uint8)(unsafe.Pointer(&v16[0]))[:uint32(uint8(int8(v14)))])
				v16[uint8(int8(v14))] = 0
				if noxflags.HasGame(noxflags.GameModeQuest) {
					v9 = nox_xxx_guide_427010(&v16[0])
					if sub_4F2EF0(v9) == 0 {
						break
					}
				}
				v10 = nox_xxx_guide_427010(&v16[0])
				nox_xxx_awardBeastGuide_4FAE80_magic_plyrgide(a1, v10, 0)
				if func() int32 {
					p := &v8
					*p++
					return *p
				}() >= int32(uint8(int8(v13))) {
					return 1
				}
			}
			return 0
		}
		return 1
	}
	v3 = (*uint32)(unsafe.Add(unsafe.Pointer(v1), 4248))
	*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v13))), 0)) = 0
	v4 = (*uint32)(unsafe.Add(unsafe.Pointer(v1), 4248))
	v5 = 40
	for {
		if *v4 != 0 {
			*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v13))), 0)) = uint8(int8(v13 + 1))
		}
		v4 = (*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*1))
		v5--
		if v5 == 0 {
			break
		}
	}
	cryptFileReadWrite((*uint8)(unsafe.Pointer(&v13))[:1])
	for i = 1; i < 41; i++ {
		if *v3 != 0 {
			*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v14))), 0)) = uint8(int8(libc.StrLen(nox_xxx_guideNameByN_427230(i))))
			cryptFileReadWrite((*uint8)(unsafe.Pointer(&v14))[:1])
			v11 = uint32(uint8(int8(v14)))
			v7 = nox_xxx_guideNameByN_427230(i)
			cryptFileReadWrite((*uint8)(unsafe.Pointer(v7))[:v11])
		}
		v3 = (*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*1))
	}
	return 1
}
func nox_xxx_guiSpellbook_41B660(a1 int32) int32 {
	var (
		result int32
		v2     int32
		v3     *uint8
		v4     *uint32
		v5     int32
		v6     int32
		v7     int8
		v8     int32
		v9     *byte
		v10    *byte
		v11    int32
		v12    bool
		v13    int32
		v14    int32
		v15    int32
		v16    uint32
		v17    uint32
		v18    int32
		v19    int8
		v20    int32
		v21    int32
		v22    int32
		v23    int32
		v24    [256]byte
	)
	result = int32(uintptr(unsafe.Pointer(noxServer.getPlayerByID(int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 36))))))))
	v2 = result
	v19 = 1
	if result == 0 {
		return result
	}
	v22 = 3
	cryptFileReadWrite((*uint8)(unsafe.Pointer(&v22))[:2])
	if int32(int16(v22)) > 3 {
		return 0
	}
	if noxflags.HasGame(noxflags.GameOnline) && !noxflags.HasGame(noxflags.GameModeQuest) {
		v19 = 0
	}
	cryptFileReadWrite((*uint8)(unsafe.Pointer(&v19))[:1])
	if int32(v19) == 0 {
		return 1
	}
	if !noxflags.HasGame(noxflags.GameModeCoop) {
		result = bool2int(noxflags.HasGame(noxflags.GameModeQuest))
		if result == 0 {
			return result
		}
	}
	if nox_xxx_cryptGetXxx() != 0 {
		cryptFileReadWrite((*uint8)(unsafe.Pointer(&v20))[:1])
		if int32(uint8(int8(v20))) > 137 {
			return 0
		}
		v11 = 0
		if int32(uint8(int8(v20))) != 0 {
			for {
				cryptFileReadWrite((*uint8)(unsafe.Pointer(&v21))[:1])
				if int32(uint8(int8(v21))) >= 256 {
					break
				}
				cryptFileReadWrite((*uint8)(unsafe.Pointer(&v24[0]))[:uint32(uint8(int8(v21)))])
				v12 = int32(int16(v22)) < 2
				v23 = 3
				v24[uint8(int8(v21))] = 0
				if !v12 {
					cryptFileReadWrite((*uint8)(unsafe.Pointer(&v23))[:4])
				}
				if noxflags.HasGame(noxflags.GameModeQuest) && v23 > 3 && int32(*(*uint8)(unsafe.Pointer(uintptr(v2 + 2251)))) != 0 {
					break
				}
				if noxflags.HasGame(noxflags.GameModeQuest) {
					if int32(*(*uint8)(unsafe.Pointer(uintptr(v2 + 2251)))) != 0 {
						v13 = things.SpellIndex(&v24[0])
						if nox_xxx_spell_4F2E70(v13) == 0 {
							break
						}
					}
				}
				if int32(int16(v22)) < 3 || int32(*(*uint8)(unsafe.Pointer(uintptr(v2 + 2251)))) != 0 {
					v18 = v23
					v15 = things.SpellIndex(&v24[0])
					nox_xxx_spellGrantToPlayer_4FB550((*nox_object_t)(unsafe.Pointer(uintptr(a1))), v15, 0, 0, v18)
				} else {
					v14 = nox_xxx_abilityNameToN_424D80(&v24[0])
					nox_xxx_abilityRewardServ_4FB9C0_ability(a1, v14, 0)
				}
				if func() int32 {
					p := &v11
					*p++
					return *p
				}() >= int32(uint8(int8(v20))) {
					return 1
				}
			}
			return 0
		}
		return 1
	}
	*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v20))), 0)) = 0
	v3 = (*uint8)(unsafe.Pointer(uintptr(v2 + 3700)))
	v4 = (*uint32)(unsafe.Pointer(uintptr(v2 + 3700)))
	if int32(*(*uint8)(unsafe.Pointer(uintptr(v2 + 2251)))) != 0 {
		v6 = 136
		for {
			if *v4 != 0 {
				*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v20))), 0)) = uint8(int8(v20 + 1))
			}
			v4 = (*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*1))
			v6--
			if v6 == 0 {
				break
			}
		}
	} else {
		v5 = 5
		for {
			if *v4 != 0 {
				*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v20))), 0)) = uint8(int8(v20 + 1))
			}
			v4 = (*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*1))
			v5--
			if v5 == 0 {
				break
			}
		}
	}
	cryptFileReadWrite((*uint8)(unsafe.Pointer(&v20))[:1])
	v7 = int8(*(*uint8)(unsafe.Pointer(uintptr(v2 + 2251))))
	v8 = 1
	if int32(v7) != 0 {
		for {
			if *(*uint32)(unsafe.Pointer(v3)) != 0 {
				*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v21))), 0)) = uint8(int8(libc.StrLen(nox_xxx_spellNameByN_424870(v8))))
				cryptFileReadWrite((*uint8)(unsafe.Pointer(&v21))[:1])
				v17 = uint32(uint8(int8(v21)))
				v10 = nox_xxx_spellNameByN_424870(v8)
				cryptFileReadWrite((*uint8)(unsafe.Pointer(v10))[:v17])
				cryptFileReadWrite(v3[:4])
			}
			v8++
			v3 = (*uint8)(unsafe.Add(unsafe.Pointer(v3), 4))
			if v8 >= 137 {
				break
			}
		}
		result = 1
	} else {
		for {
			if *(*uint32)(unsafe.Pointer(v3)) != 0 {
				*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v21))), 0)) = uint8(int8(libc.StrLen(nox_xxx_abilityGetName_425250(v8))))
				cryptFileReadWrite((*uint8)(unsafe.Pointer(&v21))[:1])
				v16 = uint32(uint8(int8(v21)))
				v9 = nox_xxx_abilityGetName_425250(v8)
				cryptFileReadWrite((*uint8)(unsafe.Pointer(v9))[:v16])
				cryptFileReadWrite(v3[:4])
			}
			v8++
			v3 = (*uint8)(unsafe.Add(unsafe.Pointer(v3), 4))
			if v8 >= 6 {
				break
			}
		}
		result = 1
	}
	return result
}
func nox_xxx_guiEnchantment_41B9C0(a1 *uint32) int32 {
	var (
		result int32
		i      int32
		v3     *byte
		v4     int32
		v5     int32
		v6     int32
		v7     int32
		v8     int32
		v9     int16
		v10    int32
		j      int32
		k      int32
		v13    int32
		v14    int8
		v15    int8
		v16    int8
		v17    int32
		v18    int32
		v19    int32
		v20    int32
		v21    int32
		v22    int32
		v23    int32
		v24    [3]int32
		v27    [256]byte
	)
	v19 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*187)))
	v16 = 1
	v18 = 5
	cryptFileReadWrite((*uint8)(unsafe.Pointer(&v18))[:2])
	if int32(int16(v18)) > 5 {
		return 0
	}
	if nox_xxx_cryptGetXxx() == 1 && noxflags.HasGame(noxflags.GameModeCoop) {
		nox_xxx_spellCastByPlayer_4FEEF0()
		sub_4FE8A0(0)
	}
	if noxflags.HasGame(noxflags.GameOnline) {
		v16 = 0
	}
	cryptFileReadWrite((*uint8)(unsafe.Pointer(&v16))[:1])
	if int32(v16) != 0 {
		result = bool2int(noxflags.HasGame(noxflags.GameModeCoop))
		if result == 0 {
			return result
		}
		if nox_xxx_cryptGetXxx() != 0 {
			cryptFileReadWrite((*uint8)(unsafe.Pointer(&v22))[:1])
			v5 = 0
			if int32(uint8(int8(v22))) != 0 {
				for {
					cryptFileReadWrite((*uint8)(unsafe.Pointer(&v20))[:1])
					cryptFileReadWrite((*uint8)(unsafe.Pointer(&v27[0]))[:uint32(uint8(int8(v20)))])
					v27[uint8(int8(v20))] = 0
					v6 = nox_xxx_enchantByName_424880(&v27[0])
					cryptFileReadWrite((*uint8)(unsafe.Pointer(&v17))[:2])
					*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v21))), 0)) = 2
					if int32(int16(v18)) >= 2 {
						cryptFileReadWrite((*uint8)(unsafe.Pointer(&v21))[:1])
					}
					v7 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*15)))
					v24[1] = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*14)))
					v13 = int32(uint8(int8(v21)))
					v24[0] = int32(uintptr(unsafe.Pointer(a1)))
					v24[2] = v7
					v8 = nox_xxx_getEnchantSpell_424920(v6)
					nox_xxx_spellAccept_4FD400(v8, unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1))))), unsafe.Pointer(a1), unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1))))), unsafe.Pointer(&v24[0]), v13)
					v9 = int16(v17)
					if int32(uint16(int16(v17))) == 0 {
						v9 = int16(uint16(nox_gameFPS))
						v17 = int32(nox_gameFPS)
					}
					*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(a1))), unsafe.Sizeof(uint16(0))*uintptr(v6)))), unsafe.Sizeof(uint16(0))*172))) = uint16(v9)
					if v6 == 26 && int32(int16(v18)) >= 3 {
						cryptFileReadWrite((*uint8)(unsafe.Pointer(&v23))[:4])
						v10 = sub_4FF2D0(51, int32(uintptr(unsafe.Pointer(a1))))
						if v10 != 0 {
							*(*uint32)(unsafe.Pointer(uintptr(v10 + 72))) = uint32(v23)
						}
					}
					v5++
					if v5 >= int32(uint8(int8(v22))) {
						break
					}
				}
			}
		} else {
			*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v22))), 0)) = uint8(sub_424CB0(int32(uintptr(unsafe.Pointer(a1)))))
			cryptFileReadWrite((*uint8)(unsafe.Pointer(&v22))[:1])
			for i = sub_424D00(); i != -1; i = sub_424D20(i) {
				if nox_xxx_testUnitBuffs_4FF350((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1)))))), int8(i)) != 0 {
					v3 = sub_4248F0(i)
					*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v20))), 0)) = uint8(int8(libc.StrLen(v3)))
					cryptFileReadWrite((*uint8)(unsafe.Pointer(&v20))[:1])
					cryptFileReadWrite((*uint8)(unsafe.Pointer(v3))[:uint32(uint8(int8(v20)))])
					cryptFileReadWrite((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(a1))), i*2))), 344))[:2])
					*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v21))), 0)) = uint8(nox_xxx_buffGetPower_4FF570(int32(uintptr(unsafe.Pointer(a1))), i))
					cryptFileReadWrite((*uint8)(unsafe.Pointer(&v21))[:1])
					if i == 26 {
						v4 = sub_4FF2D0(51, int32(uintptr(unsafe.Pointer(a1))))
						if v4 != 0 {
							v17 = int32(*(*uint32)(unsafe.Pointer(uintptr(v4 + 72))))
						} else {
							v17 = 100
						}
						cryptFileReadWrite((*uint8)(unsafe.Pointer(&v17))[:4])
					}
				}
			}
		}
		if int32(int16(v18)) >= 5 && int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v19 + 276))) + 2251)))) == 0 {
			v14 = int8(nox_common_playerIsAbilityActive_4FC250(int32(uintptr(unsafe.Pointer(a1))), 1))
			cryptFileReadWrite((*uint8)(unsafe.Pointer(&v14))[:1])
			if nox_xxx_cryptGetXxx() == 1 && int32(v14) == 1 {
				sub_4FC670(1)
			}
			v15 = int8(nox_common_playerIsAbilityActive_4FC250(int32(uintptr(unsafe.Pointer(a1))), 4))
			cryptFileReadWrite((*uint8)(unsafe.Pointer(&v15))[:1])
			v23 = sub_4FC030(int32(uintptr(unsafe.Pointer(a1))), 4)
			cryptFileReadWrite((*uint8)(unsafe.Pointer(&v23))[:4])
			if nox_xxx_cryptGetXxx() == 1 && int32(v15) == 1 {
				nox_xxx_playerExecuteAbil_4FBB70(int32(uintptr(unsafe.Pointer(a1))), 4)
				sub_4FC070(int32(uintptr(unsafe.Pointer(a1))), 4, v23)
			}
			for j = 2 - bool2int(int32(v14) != 1); j < 6; j++ {
				v17 = sub_4FBE60(int32(uintptr(unsafe.Pointer(a1))), j)
				cryptFileReadWrite((*uint8)(unsafe.Pointer(&v17))[:4])
				if nox_xxx_cryptGetXxx() == 1 {
					sub_4FBEA0(int32(uintptr(unsafe.Pointer(a1))), j, v17)
					if v17 != 0 {
						nox_xxx_netAbilRepotState_4D8100(int32(uintptr(unsafe.Pointer(a1))), int8(j), 0)
					}
				}
			}
		}
	}
	if int32(uint16(int16(v18))) == 4 && int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v19 + 276))) + 2251)))) == 0 {
		v14 = int8(nox_common_playerIsAbilityActive_4FC250(int32(uintptr(unsafe.Pointer(a1))), 1))
		cryptFileReadWrite((*uint8)(unsafe.Pointer(&v14))[:1])
		if nox_xxx_cryptGetXxx() == 1 && int32(v14) == 1 {
			sub_4FC670(1)
		}
		v15 = int8(nox_common_playerIsAbilityActive_4FC250(int32(uintptr(unsafe.Pointer(a1))), 4))
		cryptFileReadWrite((*uint8)(unsafe.Pointer(&v15))[:1])
		v19 = sub_4FC030(int32(uintptr(unsafe.Pointer(a1))), 4)
		cryptFileReadWrite((*uint8)(unsafe.Pointer(&v19))[:4])
		if nox_xxx_cryptGetXxx() == 1 && int32(v15) == 1 {
			nox_xxx_playerExecuteAbil_4FBB70(int32(uintptr(unsafe.Pointer(a1))), 4)
			sub_4FC070(int32(uintptr(unsafe.Pointer(a1))), 4, v19)
		}
		for k = 2 - bool2int(int32(v14) != 1); k < 6; k++ {
			v17 = sub_4FBE60(int32(uintptr(unsafe.Pointer(a1))), k)
			cryptFileReadWrite((*uint8)(unsafe.Pointer(&v17))[:4])
			if nox_xxx_cryptGetXxx() == 1 {
				sub_4FBEA0(int32(uintptr(unsafe.Pointer(a1))), k, v17)
				if v17 != 0 {
					nox_xxx_netAbilRepotState_4D8100(int32(uintptr(unsafe.Pointer(a1))), int8(k), 0)
				}
			}
		}
	}
	return 1
}
func sub_41BEC0(a1 int32) int32 {
	var (
		v1     int32
		v2     int32
		result int32
		i      int32
		v5     int32
		v6     int32
		j      int32
		v8     int32
		v9     int32
		v10    int32
		v11    int32
		v12    [64]byte
	)
	v1 = a1
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))) + 276))))
	*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a1))), 0)) = 1
	v10 = 1
	cryptFileReadWrite((*uint8)(unsafe.Pointer(&v10))[:2])
	if int32(int16(v10)) > 1 {
		return 0
	}
	if noxflags.HasGame(noxflags.GameOnline) {
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a1))), 0)) = 0
	}
	cryptFileReadWrite((*uint8)(unsafe.Pointer(&a1))[:1])
	if int32(uint8(int8(a1))) == 0 {
		return 1
	}
	result = bool2int(noxflags.HasGame(noxflags.GameModeCoop))
	if result == 0 {
		return result
	}
	v8 = 0
	for i = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 3644)))); i != 0; i = int32(*(*uint32)(unsafe.Pointer(uintptr(i + 64)))) {
		v8++
	}
	cryptFileReadWrite((*uint8)(unsafe.Pointer(&v8))[:2])
	if nox_xxx_cryptGetXxx() != 1 {
		if int32(uint16(int16(v8))) > 0 {
			v6 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 3644))))
			for j = int32(*(*uint32)(unsafe.Pointer(uintptr(v6 + 64)))); j != 0; j = int32(*(*uint32)(unsafe.Pointer(uintptr(j + 64)))) {
				v6 = j
			}
			for ; v6 != 0; v6 = int32(*(*uint32)(unsafe.Pointer(uintptr(v6 + 68)))) {
				*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v9))), 0)) = uint8(int8(libc.StrLen((*byte)(unsafe.Pointer(uintptr(v6))))))
				cryptFileReadWrite((*uint8)(unsafe.Pointer(&v9))[:1])
				cryptFileReadWrite((*uint8)(unsafe.Pointer(uintptr(v6)))[:uint32(uint8(int8(v9)))])
				cryptFileReadWrite((*uint8)(unsafe.Pointer(uintptr(v6 + 72)))[:2])
			}
		}
		return 1
	}
	sub_4277B0(v1, math.MaxUint16)
	v5 = 0
	if int32(uint16(int16(v8))) <= 0 {
		return 1
	}
	for {
		cryptFileReadWrite((*uint8)(unsafe.Pointer(&v9))[:1])
		cryptFileReadWrite((*uint8)(unsafe.Pointer(&v12[0]))[:uint32(uint8(int8(v9)))])
		v12[uint8(int8(v9))] = 0
		cryptFileReadWrite((*uint8)(unsafe.Pointer(&v11))[:2])
		nox_xxx_comJournalEntryAdd_427500(v1, &v12[0], int16(v11))
		v5++
		if v5 >= int32(uint16(int16(v8))) {
			break
		}
	}
	return 1
}
func sub_41C080(a1 int32) int32 {
	var (
		v1     *uint32
		v2     int32
		result int32
		v4     uint32
		v5     int32
		v6     [4]byte
	)
	v1 = (*uint32)(unsafe.Pointer(uintptr(a1)))
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))) + 276))))
	if noxflags.HasGame(noxflags.GameOnline) {
		return 1
	}
	a1 = 5
	cryptFileReadWrite((*uint8)(unsafe.Pointer(&a1))[:2])
	if int32(int16(a1)) > 5 {
		return 0
	}
	if int32(int16(a1)) >= 5 {
		if nox_xxx_cryptGetXxx() != 0 {
			cryptFileReadWrite((*uint8)(unsafe.Pointer(&v6[0]))[:4])
		} else {
			var oid int32 = int32(nox_server_LastObjectScriptID())
			cryptFileReadWrite((*uint8)(unsafe.Pointer(&oid))[:4])
			nox_server_SetLastObjectScriptID(uint32(oid))
		}
	}
	if nox_xxx_cryptGetXxx() == 0 {
		libc.StrCpy((*byte)(unsafe.Pointer(uintptr(v2+4760))), nox_xxx_mapGetMapName_409B40())
	}
	v4 = uint32(libc.StrLen((*byte)(unsafe.Pointer(uintptr(v2 + 4760)))))
	cryptFileReadWrite((*uint8)(unsafe.Pointer(&v4))[:2])
	cryptFileReadWrite((*uint8)(unsafe.Pointer(uintptr(v2 + 4760)))[:uint32(int32(uint16(v4))*2)])
	*(*uint8)(unsafe.Pointer(uintptr(int32(uint16(v4)) + v2 + 4760))) = 0
	if int32(int16(a1)) < 2 || (func() int32 {
		if nox_xxx_cryptGetXxx() != 0 {
			result = sub_500B70()
		} else {
			result = sub_500A60()
		}
		return result
	}()) != 0 {
		if int32(int16(a1)) < 3 || (func() int32 {
			result = sub_5000B0(v1)
			return result
		}()) != 0 {
			if int32(int16(a1)) >= 4 {
				*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v5))), 0)) = sub_450750()
				cryptFileReadWrite((*uint8)(unsafe.Pointer(&v5))[:1])
				sub_450760(int8(v5))
				return 1
			}
			sub_450760(0)
			return 1
		}
	}
	return result
}
func sub_41C280(a1 unsafe.Pointer) int32 {
	var (
		result int32
		v2     *byte
		v3     *byte
		v4     int32
	)
	v4 = 3
	cryptFileReadWrite((*uint8)(unsafe.Pointer(&v4))[:2])
	if int32(int16(v4)) > 3 {
		return 0
	}
	result = sub_460940(a1)
	if result != 0 {
		if int32(int16(v4)) >= 2 {
			*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a1))), 0)) = uint8(int8(nox_xxx_buttonsGetSelectedRow_45E180()))
			cryptFileReadWrite((*uint8)(unsafe.Pointer(&a1))[:1])
			if nox_xxx_cryptGetXxx() == 1 {
				nox_xxx_clientUpdateButtonRow_45E110(int32(uint8(uintptr(a1))))
			}
			*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a1))), 0)) = uint8(int8(sub_4604E0()))
			cryptFileReadWrite((*uint8)(unsafe.Pointer(&a1))[:1])
			if nox_xxx_cryptGetXxx() == 1 {
				nox_client_trapSetSelect_4604B0(int32(uint8(uintptr(a1))))
			}
		}
		if int32(int16(v4)) >= 3 {
			v2 = (*byte)(unsafe.Pointer(noxServer.getPlayerByID(int32(nox_player_netCode_85319C))))
			v3 = v2
			if v2 != nil {
				*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a1))), 0)) = uint8(*(*byte)(unsafe.Add(unsafe.Pointer(v2), 3648)))
			} else {
				*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a1))), 0)) = 4
			}
			cryptFileReadWrite((*uint8)(unsafe.Pointer(&a1))[:1])
			if nox_xxx_cryptGetXxx() == 1 {
				if noxflags.HasGame(noxflags.GameModeCoop) {
					nox_xxx_orderUnitLocal_500C70(int32(uint8(*(*byte)(unsafe.Add(unsafe.Pointer(v3), 2064)))), int32(uint8(uintptr(a1))))
				}
			}
		}
		result = 1
	}
	return result
}
func nox_xxx_parseFileInfoData_41C3B0(a1 int32) int32 {
	var (
		v1 int32
		v2 int32
		v3 int32
		v4 uint32
		v5 int32
	)
	v5 = 12
	cryptFileReadWrite((*uint8)(unsafe.Pointer(&v5))[:2])
	if int32(int16(v5)) > 12 {
		return 0
	}
	if noxflags.HasGame(noxflags.GameOnline) {
		*memmap.PtrUint32(0x85B3FC, 10980) &= 0xFFFFFFFE
		if noxflags.HasGame(noxflags.GameModeQuest) || nox_xxx_isQuest_4D6F50() != 0 || sub_4D6F70() != 0 {
			v2 = int32(*memmap.PtrUint32(0x85B3FC, 10980))
			*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v2))), 0)) = uint8(int8(int32(*memmap.PtrUint8(0x85B3FC, 10980)) | 4))
			*memmap.PtrUint32(0x85B3FC, 10980) = uint32(v2)
		} else {
			v1 = int32(*memmap.PtrUint32(0x85B3FC, 10980))
			*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v1))), 0)) = uint8(int8(int32(*memmap.PtrUint8(0x85B3FC, 10980)) | 2))
			*memmap.PtrUint32(0x85B3FC, 10980) = uint32(v1)
		}
	} else {
		*memmap.PtrUint32(0x85B3FC, 10980) = *memmap.PtrUint32(0x85B3FC, 10980)&0xFFFFFFF9 | 1
	}
	cryptFileReadWrite((*uint8)(memmap.PtrOff(0x85B3FC, 10980))[:4])
	if nox_xxx_cryptGetXxx() != 0 {
		cryptFileReadWrite((*uint8)(unsafe.Pointer(&v4))[:2])
		cryptFileReadWrite((*uint8)(memmap.PtrOff(0x85B3FC, 10984))[:uint32(int16(uint16(v4)))])
		*memmap.PtrUint8(0x85B3FC, uint32(int32(int16(uint16(v4)))+0x2AE8)) = 0
		cryptFileReadWrite((*uint8)(unsafe.Pointer(&v3))[:1])
		cryptFileReadWrite((*uint8)(memmap.PtrOff(0x85B3FC, 12008))[:uint32(uint8(int8(v3)))])
		*memmap.PtrUint8(0x85B3FC, uint32(int32(uint8(int8(v3)))+0x2EE8)) = 0
	} else {
		v4 = uint32(libc.StrLen((*byte)(memmap.PtrOff(0x85B3FC, 10984))))
		cryptFileReadWrite((*uint8)(unsafe.Pointer(&v4))[:2])
		cryptFileReadWrite((*uint8)(memmap.PtrOff(0x85B3FC, 10984))[:uint32(int16(uint16(v4)))])
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v3))), 0)) = uint8(int8(libc.StrLen((*byte)(memmap.PtrOff(0x85B3FC, 12008)))))
		cryptFileReadWrite((*uint8)(unsafe.Pointer(&v3))[:1])
		cryptFileReadWrite((*uint8)(memmap.PtrOff(0x85B3FC, 12008))[:uint32(uint8(int8(v3)))])
	}
	compatGetLocalTime((LPSYSTEMTIME)(memmap.PtrOff(0x85B3FC, 12168)))
	cryptFileReadWrite((*uint8)(memmap.PtrOff(0x85B3FC, 12168))[:2])
	cryptFileReadWrite((*uint8)(memmap.PtrOff(0x85B3FC, 12170))[:2])
	cryptFileReadWrite((*uint8)(memmap.PtrOff(0x85B3FC, 12172))[:2])
	cryptFileReadWrite((*uint8)(memmap.PtrOff(0x85B3FC, 12174))[:2])
	cryptFileReadWrite((*uint8)(memmap.PtrOff(0x85B3FC, 12176))[:2])
	cryptFileReadWrite((*uint8)(memmap.PtrOff(0x85B3FC, 12178))[:2])
	cryptFileReadWrite((*uint8)(memmap.PtrOff(0x85B3FC, 12180))[:2])
	cryptFileReadWrite((*uint8)(memmap.PtrOff(0x85B3FC, 12182))[:2])
	cryptFileReadWrite((*uint8)(memmap.PtrOff(0x85B3FC, 12187))[:3])
	cryptFileReadWrite((*uint8)(memmap.PtrOff(0x85B3FC, 12184))[:3])
	cryptFileReadWrite((*uint8)(memmap.PtrOff(0x85B3FC, 12190))[:3])
	cryptFileReadWrite((*uint8)(memmap.PtrOff(0x85B3FC, 12193))[:3])
	cryptFileReadWrite((*uint8)(memmap.PtrOff(0x85B3FC, 12196))[:3])
	cryptFileReadWrite((*uint8)(memmap.PtrOff(0x85B3FC, 12199))[:1])
	cryptFileReadWrite((*uint8)(memmap.PtrOff(0x85B3FC, 12200))[:1])
	cryptFileReadWrite((*uint8)(memmap.PtrOff(0x85B3FC, 12201))[:1])
	cryptFileReadWrite((*uint8)(memmap.PtrOff(0x85B3FC, 12202))[:1])
	cryptFileReadWrite((*uint8)(memmap.PtrOff(0x85B3FC, 12203))[:1])
	if nox_xxx_cryptGetXxx() != 0 {
		cryptFileReadWrite((*uint8)(unsafe.Pointer(&v3))[:1])
		cryptFileReadWrite((*uint8)(memmap.PtrOff(0x85B3FC, 12204))[:uint32(int32(uint8(int8(v3)))*2)])
		*memmap.PtrUint16(0x85B3FC, uint32(int32(uint8(int8(v3)))*2+0x2FAC)) = 0
	} else {
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v3))), 0)) = uint8(nox_wcslen((*libc.WChar)(memmap.PtrOff(0x85B3FC, 12204))))
		cryptFileReadWrite((*uint8)(unsafe.Pointer(&v3))[:1])
		cryptFileReadWrite((*uint8)(memmap.PtrOff(0x85B3FC, 12204))[:uint32(int32(uint8(int8(v3)))*2)])
	}
	cryptFileReadWrite((*uint8)(memmap.PtrOff(0x85B3FC, 12254))[:1])
	*memmap.PtrUint8(0x85B3FC, 12255) = 0
	cryptFileReadWrite((*uint8)(memmap.PtrOff(0x85B3FC, 12255))[:1])
	if nox_xxx_cryptGetXxx() == 0 {
		*memmap.PtrUint8(0x85B3FC, 12256) = uint8(int8(sub_467590()))
	}
	cryptFileReadWrite((*uint8)(memmap.PtrOff(0x85B3FC, 12256))[:1])
	if int32(int16(v5)) >= 11 {
		libc.StrCpy((*byte)(memmap.PtrOff(0x85B3FC, 12136)), nox_xxx_mapGetMapName_409B40())
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v3))), 0)) = uint8(int8(libc.StrLen((*byte)(memmap.PtrOff(0x85B3FC, 12136)))))
		cryptFileReadWrite((*uint8)(unsafe.Pointer(&v3))[:1])
		cryptFileReadWrite((*uint8)(memmap.PtrOff(0x85B3FC, 12136))[:uint32(uint8(int8(v3)))])
		*memmap.PtrUint8(0x85B3FC, uint32(int32(uint8(int8(v3)))+0x2F68)) = 0
	}
	if int32(int16(v5)) < 12 {
		*memmap.PtrUint8(0x85B3FC, 12257) = 0
	} else {
		cryptFileReadWrite((*uint8)(memmap.PtrOff(0x85B3FC, 12257))[:1])
	}
	return 1
}
func sub_41C780(a1 int32) int32 {
	var (
		i  int32
		v2 *byte
		v3 bool
		v4 int32
		v5 int32
		v6 int4
	)
	v5 = 11
	cryptFileReadWrite((*uint8)(unsafe.Pointer(&v5))[:2])
	if int32(int16(v5)) > 11 {
		return 0
	}
	if int32(int16(v5)) < 11 || (func() bool {
		v3 = !noxflags.HasGame(noxflags.GameOnline)
		cryptFileReadWrite((*uint8)(unsafe.Pointer(&v3))[:1])
		return v3
	}()) {
		if nox_xxx_cryptGetXxx() != 0 {
			v6.field_0 = 0
			v6.field_4 = 0
			v6.field_8 = 0
			v6.field_C = 0
		} else {
			sub_43DD10(&v6.field_0)
			v4 = sub_43DB20()
		}
		cryptFileReadWrite((*uint8)(unsafe.Pointer(&v6.field_C))[:4])
		cryptFileReadWrite((*uint8)(unsafe.Pointer(&v6.field_8))[:4])
		cryptFileReadWrite((*uint8)(unsafe.Pointer(&v6))[:4])
		cryptFileReadWrite((*uint8)(unsafe.Pointer(&v6.field_4))[:4])
		cryptFileReadWrite((*uint8)(unsafe.Pointer(&v4))[:4])
		for i = 0; i < v4; i++ {
			v2 = sub_43DB40(i)
			if nox_xxx_cryptGetXxx() == 1 {
				*(*uint32)(unsafe.Pointer(v2)) = 0
				*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), unsafe.Sizeof(uint32(0))*1))) = 0
				*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), unsafe.Sizeof(uint32(0))*2))) = 0
				*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), unsafe.Sizeof(uint32(0))*3))) = 0
			}
			cryptFileReadWrite((*uint8)(unsafe.Add(unsafe.Pointer(v2), 12))[:4])
			cryptFileReadWrite((*uint8)(unsafe.Add(unsafe.Pointer(v2), 8))[:4])
			cryptFileReadWrite((*uint8)(unsafe.Pointer(v2))[:4])
			cryptFileReadWrite((*uint8)(unsafe.Add(unsafe.Pointer(v2), 4))[:4])
		}
		if nox_xxx_cryptGetXxx() == 1 && !noxflags.HasGame(noxflags.GameOnline) {
			sub_43D9E0(&v6)
			sub_43DB30(v4)
		}
	}
	return 1
}
func sub_41CAC0(a1 *byte, a2 **byte) *File {
	var (
		result *File
		v3     **byte
		v4     int32
		v5     int32
		v6     int32
		v7     *uint8
		v8     int32
		v9     int32
		v10    *uint32
		v11    int32
		v12    int32
		v13    *byte
	)
	result = nox_binfile_open_408CC0(a1, 0)
	nox_file_2 = result
	if result != nil {
		result = (*File)(unsafe.Pointer(uintptr(nox_binfile_cryptSet_408D40((*File)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(result)))))), 27))))
		if result != nil {
			v3 = a2
			v4 = 0
			for {
				nox_binfile_fread_408E40((*byte)(unsafe.Pointer(&v13)), 4, 1, nox_file_2)
				if v13 == nil {
					break
				}
				v5 = nox_binfile_ftell_426A50(nox_file_2)
				nox_binfile_fread_align_408FE0((*byte)(unsafe.Pointer(&v12)), 4, 1, nox_file_2)
				v6 = nox_binfile_ftell_426A50(nox_file_2)
				if *memmap.PtrUint32(0x587000, 55816) != 0 {
					v7 = (*uint8)(memmap.PtrOff(0x587000, 55816))
					for v13 != *((**byte)(unsafe.Add(unsafe.Pointer((**byte)(unsafe.Pointer(v7))), unsafe.Sizeof((*byte)(nil))*1))) {
						v8 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v7))), unsafe.Sizeof(uint32(0))*3))))
						v7 = (*uint8)(unsafe.Add(unsafe.Pointer(v7), 12))
						if v8 == 0 {
							goto LABEL_9
						}
					}
					v9 = v6 - v5
					*v3 = v13
					v10 = (*uint32)(unsafe.Pointer((**byte)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(v3))), v9))), -4))))))
					v4 += v9 + 4
					*v10 = uint32(v12)
					v3 = (**byte)(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer(v10), unsafe.Sizeof(uint32(0))*2))))
					if v12 != 0 {
						v11 = v12
						v4 += v12
						for {
							nox_binfile_fread_408E40((*byte)(unsafe.Pointer(&a1)), 1, 1, nox_file_2)
							*(*uint8)(unsafe.Pointer(v3)) = uint8(uintptr(unsafe.Pointer(a1)))
							v3 = (**byte)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(v3))), 1))))
							v11--
							if v11 == 0 {
								break
							}
						}
					}
				} else {
				LABEL_9:
					nox_binfile_fseek_409050(nox_file_2, v12, stdio.SEEK_CUR)
				}
			}
			*v3 = nil
			nox_binfile_close_408D90(nox_file_2)
			result = (*File)(unsafe.Pointer(uintptr(v4 + 4)))
		}
	}
	return result
}
func sub_41CC00(a1 *byte) int32 {
	var (
		v1     int32
		v2     int32
		result int32
		v4     int32
	)
	nullsub_5()
	v1 = nox_xxx_computeServerPlayerDataBufferSize_41CC50(a1)
	v2 = v1 + 4
	result = int32(uintptr(alloc.Calloc(1, int(v1+4))))
	v4 = result
	if result != 0 {
		sub_41CAC0(a1, (**byte)(unsafe.Pointer(uintptr(result))))
		sub_40BC60(31, 3, CString("SAVE_SERVER"), v4, v2, 0)
		result = 1
	}
	return result
}
func nox_xxx_computeServerPlayerDataBufferSize_41CC50(a1 *byte) int32 {
	var (
		v1     *byte
		v2     *File
		result int32
		v4     int32
		v5     int32
		v6     int32
		v7     *uint8
		v8     int32
		v9     int32
	)
	v1 = a1
	v2 = nox_binfile_open_408CC0(a1, 0)
	nox_file_2 = v2
	if v2 != nil {
		if nox_binfile_cryptSet_408D40((*File)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v2)))))), 27) != 0 {
			v4 = 0
			for {
				nox_binfile_fread_408E40((*byte)(unsafe.Pointer(&a1)), 4, 1, nox_file_2)
				if nox_binfile_lastErr_409370(nox_file_2) == -1 {
					break
				}
				if a1 == nil {
					v4 += 4
					break
				}
				v5 = nox_binfile_ftell_426A50(nox_file_2)
				nox_binfile_fread_align_408FE0((*byte)(unsafe.Pointer(&v9)), 4, 1, nox_file_2)
				v6 = nox_binfile_ftell_426A50(nox_file_2) - v5
				if *memmap.PtrUint32(0x587000, 55816) != 0 {
					v7 = (*uint8)(memmap.PtrOff(0x587000, 55816))
					for a1 != *((**byte)(unsafe.Add(unsafe.Pointer((**byte)(unsafe.Pointer(v7))), unsafe.Sizeof((*byte)(nil))*1))) {
						v8 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v7))), unsafe.Sizeof(uint32(0))*3))))
						v7 = (*uint8)(unsafe.Add(unsafe.Pointer(v7), 12))
						if v8 == 0 {
							goto LABEL_14
						}
					}
					v4 += v9 + v6 + 4
				}
			LABEL_14:
				nox_binfile_fseek_409050(nox_file_2, v9, stdio.SEEK_CUR)
			}
			nox_binfile_close_408D90(nox_file_2)
			result = v4
		} else {
			nox_xxx_networkLog_printf_413D30(CString("computeServerPlayerDataBufferSize: Can't key file '%s'\n"), v1)
			result = 0
		}
	} else {
		nox_xxx_networkLog_printf_413D30(CString("computeServerPlayerDataBufferSize: Can't open file '%s'\n"), v1)
		result = 0
	}
	return result
}
func nox_xxx_SavePlayerDataFromClient_41CD70(a1 *byte, a2 *uint8, a3 int32) int32 {
	var (
		v3     *byte
		v4     *File
		result int32
		v6     int32
		v7     *uint8
	)
	v3 = a1
	v4 = nox_binfile_open_408CC0(a1, 1)
	nox_file_2 = v4
	if v4 != nil {
		if nox_binfile_cryptSet_408D40((*File)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v4)))))), 27) != 0 {
			v6 = a3
			if a3 != 0 {
				v7 = a2
				for {
					*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a1))), 0)) = *func() *uint8 {
						p := &v7
						x := *p
						*p = (*uint8)(unsafe.Add(unsafe.Pointer(*p), 1))
						return x
					}()
					nox_binfile_fwrite_409200((*byte)(unsafe.Pointer(&a1)), 1, 1, nox_file_2)
					v6--
					if v6 == 0 {
						break
					}
				}
			}
			nox_binfile_close_408D90(nox_file_2)
			result = 1
		} else {
			nox_xxx_networkLog_printf_413D30(CString("SavePlayerDataFromClient: Can't key file '%s'\n"), v3)
			result = 0
		}
	} else {
		nox_xxx_networkLog_printf_413D30(CString("SavePlayerDataFromClient: Can't open file '%s'\n"))
		result = 0
	}
	return result
}
func nox_xxx_netSavePlayer_41CE00() int32 {
	var v2 [3]byte
	v2[0] = 193
	*(*uint16)(unsafe.Pointer(&v2[1])) = uint16(nox_player_netCode_85319C)
	nox_xxx_netClientSend2_4E53C0(31, unsafe.Pointer(&v2[0]), 3, 0, 1)
	return 1
}
func sub_41CEE0(a1 int32, a2 int32) int32 {
	var (
		result int32
		v3     int32
		v4     *uint8
		v6     int32
		v8     int32
	)
	libc.MemCpy(memmap.PtrOff(0x85B3FC, 10980), unsafe.Pointer(uintptr(a1)), int(unsafe.Sizeof(nox_savegame_xxx{})))
	result = cryptFileOpen((*byte)(memmap.PtrOff(0x85B3FC, 10984)), 0, 27)
	if result != 0 {
		v3 = a2
		if a2 == 0 {
			a1 = 1
		}
		if *memmap.PtrUint32(0x587000, 55936) != 0 {
			v4 = (*uint8)(memmap.PtrOff(0x587000, 55940))
			for {
				if v3 != 0 || *(*uint32)(unsafe.Pointer(v4)) == uint32(a1) {
					cryptFileReadWrite(v4[:4])
					nox_xxx_crypt_426C90()
					v6 = (*((*funcuint32(int32))(unsafe.Add(unsafe.Pointer((*funcuint32(int32))(unsafe.Pointer(v4))), unsafe.Sizeof(uintptr(0))*1))))(0)
					nox_xxx_crypt_426D40()
					if v6 == 0 {
						break
					}
				}
				v8 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v4))), unsafe.Sizeof(uint32(0))*2))))
				v4 = (*uint8)(unsafe.Add(unsafe.Pointer(v4), 12))
				if v8 == 0 {
					goto LABEL_10
				}
			}
			cryptFileClose()
			result = 0
		} else {
		LABEL_10:
			a1 = 0
			cryptFileReadWrite((*uint8)(unsafe.Pointer(&a1))[:4])
			cryptFileClose()
			result = 1
		}
	}
	return result
}
func sub_41CFA0(a1 *byte, a2 int32) int32 {
	var (
		v2     uint16
		v3     *File
		result int32
		v5     *byte
	)
	if sub_419EE0(int8(a2)) != 0 {
		return 0
	}
	v2 = uint16(int16(nox_xxx_computeServerPlayerDataBufferSize_41CC50(a1)))
	if int32(v2) == 0 {
		return 0
	}
	v3 = nox_binfile_open_408CC0(a1, 0)
	nox_file_2 = v3
	if v3 == nil {
		nox_xxx_networkLog_printf_413D30(CString("SendPlayerSaveDataToClient: Can't open file '%s'\n"), a1)
		return 0
	}
	if nox_binfile_cryptSet_408D40((*File)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v3)))))), 27) != 0 {
		v5 = (*byte)(alloc.Calloc(1, int(v2)))
		if v5 != nil {
			nox_binfile_fread_408E40(v5, int32(v2), 1, nox_file_2)
			nox_binfile_close_408D90(nox_file_2)
			sub_419EB0(int8(a2), 1)
			sub_40BC60(a2, 2, CString("SAVEDATA"), int32(uintptr(unsafe.Pointer(v5))), int32(v2), 1)
			alloc.Free(unsafe.Pointer(v5))
		} else {
			nox_binfile_close_408D90(nox_file_2)
		}
		result = 1
	} else {
		nox_xxx_networkLog_printf_413D30(CString("SavePlayerOnClient: Unable to key file '%s'\n"), memmap.PtrOff(0x85B3FC, 10984))
		result = 0
	}
	return result
}
func sub_41D170() int32 {
	var v1 [2]byte
	v1[0] = 240
	v1[1] = 27
	return nox_xxx_netClientSend2_4E53C0(31, unsafe.Pointer(&v1[0]), 2, 0, 0)
}
func sub_41D1A0(a1 int32) int32 {
	var result int32
	result = a1
	*memmap.PtrUint32(6112660, 527720) = uint32(a1)
	return result
}
func sub_41D1B0() int32 {
	return int32(*memmap.PtrUint32(6112660, 527720))
}
func sub_41D440() int32 {
	var (
		result int32
		v1     int32
		v2     [128]byte
	)
	sub_41E370()
	sub_41E300(0)
	result = sub_40CE60()
	if result != 0 {
		sub_41FCF0()
		gameSetCliDrawFunc(unsafe.Pointer(cgoFuncAddr(sub_41E210)))
		sub_41EAC0()
		libc.StrCpy(&v2[0], CString("SOFTWARE\\Westwood\\Nox"))
		sub_40DA90(int32(uintptr(memmap.PtrOff(0x587000, 58096))), int32(uintptr(unsafe.Pointer(&v2[0]))))
		v1 = sub_4200E0()
		sub_40D950(v1)
		result = 1
	}
	return result
}
func sub_41D530() int32 {
	var (
		v0     int32
		v1     int32
		result int32
		v3     int32
		v4     int32
		v5     int32
		v6     int32
		v7     int32
		v8     int32
	)
	if sub_41FBE0((*uint32)(unsafe.Pointer(&v8)), (*uint32)(unsafe.Pointer(&v7))) != 1 {
		return 0
	}
	v0 = sub_420020()
	if v0 == -1 {
		if sub_41E300(11) != 0 {
			dword_5d4594_2660652 = 2147745900
			v3 = sub_41E2F0()
			sub_41DA70(v3, 5)
			v4 = sub_41E2F0()
			sub_41DA70(v4, 9)
		}
		result = 1
	} else {
		sub_420580()
		v6 = v7
		v5 = v8
		v1 = sub_4200E0()
		sub_40D280(v1, v0, v5, v6, 30)
		result = 1
	}
	return result
}
func sub_41D5D0() int32 {
	return 1
}
func sub_41D5E0() {
	var (
		result *uint8
		v1     *uint8
		v2     unsafe.Pointer
	)
	result = (*uint8)(unsafe.Pointer(nox_common_playerInfoFromNum_417090(31)))
	v1 = result
	if result != nil {
		result = *(**uint8)(unsafe.Pointer(&dword_5d4594_2660032))
		if dword_5d4594_2660032 != 0 {
			*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v1))), unsafe.Sizeof(uint32(0))*517))) = dword_5d4594_2660032
			nox_swprintf((*libc.WChar)(unsafe.Add(unsafe.Pointer((*libc.WChar)(unsafe.Pointer(v1))), unsafe.Sizeof(libc.WChar(0))*1036)), libc.CWString("%S"), memmap.PtrOff(0x85B3FC, 10395))
			v2 = unsafe.Pointer(sub_425A70(int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v1))), unsafe.Sizeof(uint32(0))*517))))))
			if v2 == nil {
				v2 = unsafe.Pointer(sub_425AD0(int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v1))), unsafe.Sizeof(uint32(0))*517)))), (*libc.WChar)(unsafe.Add(unsafe.Pointer((*libc.WChar)(unsafe.Pointer(v1))), unsafe.Sizeof(libc.WChar(0))*1036))))
			}
			sub_425B30(unsafe.Pointer(uintptr(int32(uintptr(v2)))), int32(*(*uint8)(unsafe.Add(unsafe.Pointer(v1), 2064))))
		}
	}
}
func sub_41D650() int32 {
	var (
		v0     *byte
		result int32
	)
	v0 = sub_41FA40()
	result = sub_41F800(v0)
	if result != 0 {
		result = sub_40D530(result)
	}
	*memmap.PtrUint32(6112660, 371700) = 1
	return result
}
func sub_41D670(a1 *byte) int32 {
	var (
		v1 int32
		v3 int16
		v4 [72]byte
	)
	if sub_420230(&v4[0], (*uint16)(unsafe.Pointer(&v3))) == 0 {
		return 0
	}
	v1 = sub_4200E0()
	return sub_40DF20(int32(uintptr(unsafe.Pointer(&v4[0]))), int32(v3), int32(uintptr(unsafe.Pointer(a1))), v1, -1, 0, 0)
}
func sub_41D6C0() int32 {
	var (
		v0     int32
		v1     *byte
		v2     int32
		v3     int32
		v4     *byte
		v5     int32
		result int32
		v7     *byte
		v8     int32
		v9     int16
		v10    *byte
		v11    int32
		v12    int32
		v13    int32
		v14    [72]byte
		v15    [1024]byte
	)
	v0 = 1
	*(*[1024]byte)(unsafe.Pointer(&v15[0])) = [1024]byte{}
	v11 = 1
	v12 = 1
	if sub_420230(&v14[0], (*uint16)(unsafe.Pointer(&v9))) == 0 {
		return 0
	}
	if uint32(nox_common_playerInfoCount_416F40()) > 25 {
		v1 = (*byte)(unsafe.Pointer(nox_common_playerInfoGetFirst_416EA0()))
		v13 = 25
		for {
			if *(*byte)(unsafe.Add(unsafe.Pointer(v1), 2096)) != 0 {
				if v0 != 0 {
					libc.StrCat(&v15[0], (*byte)(unsafe.Add(unsafe.Pointer(v1), 2096)))
					v0 = 0
				} else {
					*(*uint16)(unsafe.Pointer(&v15[libc.StrLen(&v15[0])])) = *memmap.PtrUint16(0x587000, 58112)
					libc.StrCat(&v15[0], (*byte)(unsafe.Add(unsafe.Pointer(v1), 2096)))
				}
			}
			v10 = (*byte)(unsafe.Pointer(nox_common_playerInfoGetNext_416EE0((*nox_playerInfo)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v1)))))))))
			if func() int32 {
				p := &v13
				*p--
				return *p
			}() == 0 {
				break
			}
			v1 = v10
		}
		if v0 == 0 {
			v2 = sub_4200E0()
			v11 = sub_40DF20(int32(uintptr(unsafe.Pointer(&v14[0]))), int32(v9), int32(uintptr(unsafe.Pointer(&v15[0]))), v2, -1, 0, 0)
		}
		v3 = 1
		if v10 == nil {
			goto LABEL_37
		}
		v15[0] = 0
		v4 = (*byte)(unsafe.Pointer(nox_common_playerInfoGetNext_416EE0((*nox_playerInfo)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v10)))))))))
		if v4 == nil {
			goto LABEL_37
		}
		for {
			if *(*byte)(unsafe.Add(unsafe.Pointer(v4), 2096)) != 0 {
				if v3 != 0 {
					libc.StrCat(&v15[0], (*byte)(unsafe.Add(unsafe.Pointer(v4), 2096)))
					v3 = 0
				} else {
					*(*uint16)(unsafe.Pointer(&v15[libc.StrLen(&v15[0])])) = *memmap.PtrUint16(0x587000, 58116)
					libc.StrCat(&v15[0], (*byte)(unsafe.Add(unsafe.Pointer(v4), 2096)))
				}
			}
			v4 = (*byte)(unsafe.Pointer(nox_common_playerInfoGetNext_416EE0((*nox_playerInfo)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v4)))))))))
			if v4 == nil {
				break
			}
		}
		if v3 == 0 {
			v5 = sub_4200E0()
			result = sub_40DF20(int32(uintptr(unsafe.Pointer(&v14[0]))), int32(v9), int32(uintptr(unsafe.Pointer(&v15[0]))), v5, -1, 0, 0)
		} else {
		LABEL_37:
			result = v12
		}
		if v11 != 0 && result == 1 {
			return result
		}
		return 0
	}
	v7 = (*byte)(unsafe.Pointer(nox_common_playerInfoGetFirst_416EA0()))
	if v7 != nil {
		for {
			if *(*byte)(unsafe.Add(unsafe.Pointer(v7), 2096)) != 0 {
				if v0 != 0 {
					libc.StrCat(&v15[0], (*byte)(unsafe.Add(unsafe.Pointer(v7), 2096)))
					v0 = 0
				} else {
					*(*uint16)(unsafe.Pointer(&v15[libc.StrLen(&v15[0])])) = *memmap.PtrUint16(0x587000, 58120)
					libc.StrCat(&v15[0], (*byte)(unsafe.Add(unsafe.Pointer(v7), 2096)))
				}
			}
			v7 = (*byte)(unsafe.Pointer(nox_common_playerInfoGetNext_416EE0((*nox_playerInfo)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v7)))))))))
			if v7 == nil {
				break
			}
		}
		if v0 == 0 {
			v8 = sub_4200E0()
			v11 = sub_40DF20(int32(uintptr(unsafe.Pointer(&v14[0]))), int32(v9), int32(uintptr(unsafe.Pointer(&v15[0]))), v8, -1, 0, 0)
		}
	}
	return v11
}
func sub_41D9F0() unsafe.Pointer {
	var (
		v0     *uint8
		result unsafe.Pointer
	)
	v0 = (*uint8)(memmap.PtrOff(0x587000, 58132))
	for {
		result = alloc.Calloc(1, 46)
		*(*uint32)(unsafe.Pointer(v0)) = uint32(uintptr(result))
		v0 = (*uint8)(unsafe.Add(unsafe.Pointer(v0), 16))
		if int32(uintptr(unsafe.Pointer(v0))) >= int32(uintptr(memmap.PtrOff(0x587000, 58340))) {
			break
		}
	}
	return result
}
func sub_41DA10(a1 int32) int32 {
	var (
		v1     *uint16
		result int32 = 0
	)
	v1 = *(**uint16)(memmap.PtrOff(0x587000, uint32(a1*16+0xE314)))
	if v1 != nil {
		result = 0
		libc.MemSet(unsafe.Pointer(v1), 0, 44)
		*(*uint16)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint16(0))*22)) = 0
		*memmap.PtrUint32(0x587000, uint32(a1*16+0xE318)) = 0
	}
	return result
}
func sub_41DA40() unsafe.Pointer {
	var (
		v0     *unsafe.Pointer
		result unsafe.Pointer
	)
	v0 = (*unsafe.Pointer)(memmap.PtrOff(0x587000, 58132))
	for {
		result = *v0
		if *v0 != nil {
			*v0 = nil
			*v0 = nil
		}
		v0 = (*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(v0), unsafe.Sizeof(unsafe.Pointer(nil))*4))
		if int32(uintptr(unsafe.Pointer(v0))) >= int32(uintptr(memmap.PtrOff(0x587000, 58340))) {
			break
		}
	}
	return result
}
func sub_41DA70(a1 int32, a2 int16) int32 {
	var (
		result int32
		v3     int32
		v4     int32
	)
	result = a1 * 16
	v3 = int32(*memmap.PtrUint32(0x587000, uint32(a1*16+0xE314)))
	if v3 != 0 && (func() bool {
		v4 = int32(*memmap.PtrUint32(0x587000, uint32(a1*16+0xE318)))
		return v4 < 23
	}()) {
		*(*uint16)(unsafe.Pointer(uintptr(v3 + v4*2))) = uint16(a2)
		*memmap.PtrUint32(0x587000, uint32(a1*16+0xE318))++
	} else {
		result = sub_41E300(11)
		if result != 0 {
			dword_5d4594_2660652 = 0
		}
	}
	return result
}
func sub_41DB90() int32 {
	var (
		v0 int32
		v1 int32
		v2 int32
		v3 int32
		v5 int32
	)
	v0 = 1
	v1 = int32(*memmap.PtrUint32(0x587000, uint32(sub_41E2F0()*16+0xE318)))
	if v1 != 0 {
		v5 = v1
		for {
			v2 = sub_41E2F0()
			switch sub_41DCC0(v2) {
			case 0:
				v0 = sub_41D530()
			case 1:
				nox_xxx_guiServerListLoad_449530()
			case 2:
				v3 = sub_41E520()
				if v3 != 0 {
					sub_468110_wol_prog()
					nox_sprintf((*byte)(memmap.PtrOff(6112660, 527996)), CString("%s/%s"), v3+81, v3+337)
					nox_sprintf((*byte)(memmap.PtrOff(6112660, 527732)), CString("%s\\%s"), v3+468, v3+337)
					nox_fs_mkdir((*byte)(unsafe.Pointer(uintptr(v3 + 468))))
					v0 = sub_40DC90(v3+16, v3+370, v3+403, int32(uintptr(memmap.PtrOff(6112660, 527996))), int32(uintptr(memmap.PtrOff(6112660, 527732))), int32(uintptr(memmap.PtrOff(0x587000, 58684))))
					sub_41E300(4)
				} else {
					sub_41E300(5)
					sub_40DC70()
				}
			case 3:
				sub_468080()
				v0 = nox_xxx_officialStringCmp_41FDE0()
				sub_41E300(5)
			default:
			}
			v5--
			if v5 == 0 {
				break
			}
		}
	}
	sub_40D250()
	return v0
}
func sub_41DCC0(a1 int32) int32 {
	var (
		v1 *uint16
		v2 int32
		v3 int32
		v4 int32
	)
	v1 = *(**uint16)(memmap.PtrOff(0x587000, uint32(a1*16+0xE314)))
	if v1 == nil {
		return a1
	}
	v2 = int32(*memmap.PtrUint32(0x587000, uint32(a1*16+0xE318)))
	if v2 <= 0 {
		return a1
	}
	v3 = v2 - 1
	v4 = int32(*v1)
	*memmap.PtrUint32(0x587000, uint32(a1*16+0xE318)) = uint32(v3)
	if v3 != 0 {
		libc.MemMove(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer(v1))))), unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(v1))), 2), int(v3*2))
	}
	return v4
}
func sub_41DD10() int32 {
	var (
		i  int32
		v1 int32
	)
	for i = int32(*memmap.PtrUint32(0x587000, 58216)); i != 0; i-- {
		v1 = sub_41E2F0()
		sub_41DCC0(v1)
	}
	return 1
}
func sub_41DD40() int32 {
	var (
		v0 int32
		v1 int32
		v2 int32
		v3 int32
		v4 int32
	)
	if *memmap.PtrUint32(0x587000, 58200) != 0 {
		v0 = int32(*memmap.PtrUint32(0x587000, 58200))
		for {
			v1 = sub_41E2F0()
			v2 = sub_41DCC0(v1) - 16
			if v2 != 0 {
				v3 = v2 - 1
				if v3 != 0 {
					if v3 == 1 {
						sub_468170_wol_prog()
					}
				} else {
					sub_468340()
					sub_41E300(5)
				}
			} else {
				sub_4682B0_wol_prog()
				sub_41E300(2)
				v4 = sub_41E2F0()
				sub_41DA70(v4, 2)
			}
			v0--
			if v0 == 0 {
				break
			}
		}
	}
	sub_40DCD0()
	return 1
}
func sub_41DFC0() int32 {
	var (
		v0 int32
		v1 int32
	)
	if *memmap.PtrUint32(0x587000, 58248) != 0 {
		v0 = int32(*memmap.PtrUint32(0x587000, 58248))
		for {
			v1 = sub_41E2F0()
			switch sub_41DCC0(v1) {
			case 3:
				sub_41FF10()
			case 5:
				sub_41DAC0(*(*int32)(unsafe.Pointer(&dword_5d4594_2660652)))
			case 10:
				sub_446E60()
			case 11:
				sub_446EB0()
			case 12:
				sub_41F3A0(-99999, 1)
			case 13:
				sub_40D3B0()
			case 14:
				sub_446D50()
			case 15:
				sub_41F4B0()
				sub_4475E0()
			default:
			}
			v0--
			if v0 == 0 {
				break
			}
		}
	}
	sub_40D250()
	return 1
}
func sub_41E080() int32 {
	var (
		v0 int32
		v1 int32
		v2 int32
	)
	if *memmap.PtrUint32(0x587000, 58184) != 0 {
		v0 = int32(*memmap.PtrUint32(0x587000, 58184))
		for {
			v1 = sub_41E2F0()
			v2 = sub_41DCC0(v1)
			if v2 == 3 {
				sub_41FF10()
			} else if v2 == 7 {
				sub_41E300(7)
				sub_4207F0(3)
			}
			v0--
			if v0 == 0 {
				break
			}
		}
	}
	sub_40D250()
	return 1
}
func sub_41E170() int32 {
	var (
		v0 int32
		v1 int32
	)
	if *memmap.PtrUint32(0x587000, 58280) != 0 {
		v0 = int32(*memmap.PtrUint32(0x587000, 58280))
		for {
			v1 = sub_41E2F0()
			if sub_41DCC0(v1) == 8 {
				sub_41FF10()
			}
			v0--
			if v0 == 0 {
				break
			}
		}
	}
	sub_40D250()
	return 1
}
func sub_41E1B0() int32 {
	var (
		v0 int32
		v1 int32
		v2 int32
		v3 int32
	)
	if *memmap.PtrUint32(0x587000, 58312) == 0 {
		return 1
	}
	v0 = int32(*memmap.PtrUint32(0x587000, 58312))
	for {
		v1 = sub_41E2F0()
		v2 = sub_41DCC0(v1) - 5
		if v2 != 0 {
			v3 = v2 - 1
			if v3 != 0 {
				if v3 == 3 {
					sub_4207F0(1)
					nox_game_checkStateWol_43C260()
				}
			} else {
				nox_xxx_wolApiError_41D1D0(1)
			}
		} else {
			nox_xxx_wolApiError_41D1D0(0)
		}
		v0--
		if v0 == 0 {
			break
		}
	}
	return 1
}
func sub_41E210() int32 {
	var (
		v0     int32
		v1     int32
		v2     int32
		result int32
		v4     int32
	)
	switch sub_41E2F0() {
	case 0:
		sub_41D9F0()
		*memmap.PtrUint32(0x85B3FC, 10288) = 0
		sub_4207F0(1)
		sub_41E300(1)
		*memmap.PtrUint32(6112660, 527992) = uint32(sub_41E2F0())
		result = 1
	case 1:
		fallthrough
	case 2:
		fallthrough
	case 3:
		fallthrough
	case 4:
		fallthrough
	case 5:
		fallthrough
	case 6:
		fallthrough
	case 7:
		fallthrough
	case 8:
		fallthrough
	case 9:
		fallthrough
	case 11:
		if *memmap.PtrUint32(0x587000, uint32(sub_41E2F0()*16+58140)) != 0 {
			v0 = sub_41E2F0()
			v4 = int32(*memmap.PtrUint32(6112660, 527992))
			v1 = v0 * 16
			v2 = sub_41E2F0()
			(cgoAsFunc(*(*uint32)(memmap.PtrOff(0x587000, uint32(v1+58140))), (*func(uint32, int32))(nil)))(*memmap.PtrUint32(0x587000, uint32(v2*16+0xE314)), v4)
		}
		*memmap.PtrUint32(6112660, 527992) = uint32(sub_41E2F0())
		result = 1
	case 10:
		sub_41DA40()
		result = 1
	case 12:
		nox_xxx_wolApiError_41D1D0(0)
		result = 1
	default:
		result = 0
	}
	return result
}
func sub_41E2F0() int32 {
	return int32(dword_5d4594_527988)
}
func sub_41E370() int32 {
	var result int32
	result = 0
	dword_5d4594_528252 = 0
	dword_5d4594_528256 = 0
	dword_5d4594_528260 = 0
	dword_5d4594_528264 = 0
	return result
}
func nox_xxx_reconAttempt_41E390() int32 {
	var result int32
	if nox_frame_xxx_2598000-dword_5d4594_528264 <= (nox_gameFPS * 3600) {
		result = int32(dword_5d4594_528252)
		if dword_5d4594_528252 != 0 {
			result = int32(dword_5d4594_528256)
			if dword_5d4594_528256 == 0 {
				nox_xxx_networkLog_printf_413D30(CString("RECON: Attempting to re-login"))
				sub_40E090()
				result = nox_xxx_officialStringCmp_41FDE0()
				if result == 1 {
					dword_5d4594_528256 = 1
				} else {
					result = sub_41E470()
				}
			}
		}
	} else {
		sub_41E370()
		result = sub_41E4B0(1)
	}
	return result
}
func nox_xxx_reconStart_41E400() {
	if dword_5d4594_528252 != 1 && dword_5d4594_528256 != 1 {
		if dword_5d4594_528260 == 0 {
			if dword_5d4594_528264 == 0 {
				nox_xxx_networkLog_printf_413D30(CString("RECON: Starting reconnection process frame (%d)"), nox_frame_xxx_2598000)
				dword_5d4594_528252 = 1
				dword_5d4594_528256 = 0
				dword_5d4594_528264 = nox_frame_xxx_2598000
				dword_5d4594_528260 = nox_frame_xxx_2598000 + nox_gameFPS*120
			}
		}
	}
}
func sub_41E470() int32 {
	var result int32
	nox_xxx_networkLog_printf_413D30(CString("RECON: TryReconnectAgain called on frame (%d)"), nox_frame_xxx_2598000)
	dword_5d4594_528256 = 0
	result = int32(nox_frame_xxx_2598000 + nox_gameFPS*120)
	dword_5d4594_528260 = nox_frame_xxx_2598000 + nox_gameFPS*120
	return result
}
func sub_41E4B0(a1 int32) int32 {
	var result int32
	result = a1
	*memmap.PtrUint32(6112660, 528268) = uint32(a1)
	return result
}
func sub_41E4C0() int32 {
	return int32(*memmap.PtrUint32(6112660, 528268))
}
func sub_41E4D0(a1 *uint32) {
	var (
		v1 *uint32
		i  *uint32
		v3 *uint32
	)
	v1 = a1
	for i = nil; v1 != nil; v1 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*3))))) {
		v3 = (*uint32)(alloc.Calloc(1, 724))
		if v3 != nil {
			libc.MemCpy(unsafe.Pointer(v3), unsafe.Pointer(v1), 724)
			*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*3)) = 0
			if i != nil {
				*(*uint32)(unsafe.Add(unsafe.Pointer(i), unsafe.Sizeof(uint32(0))*3)) = uint32(uintptr(unsafe.Pointer(v3)))
			} else {
				dword_5d4594_528272 = uint32(uintptr(unsafe.Pointer(v3)))
			}
			i = v3
		}
	}
}
func sub_41E520() int32 {
	var result int32
	result = int32(dword_5d4594_528272)
	if dword_5d4594_528272 != 0 {
		dword_5d4594_528272 = *(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_528272 + 12)))
	}
	return result
}
func sub_41E540() int32 {
	return bool2int(dword_5d4594_528272 != 0)
}
func sub_41E560(a1 int32, a2 int32, a3 int32, a4 int32) int32 {
	var result int32
	*memmap.PtrUint32(6112660, 528276) = uint32(a1)
	result = a4
	*memmap.PtrUint32(6112660, 528280) = uint32(a2)
	*memmap.PtrUint32(6112660, 528284) = uint32(a3)
	*memmap.PtrUint32(6112660, 528288) = uint32(a4)
	return result
}
func sub_41E590(a1 *uint32, a2 *uint32, a3 *uint32, a4 *uint32) *uint32 {
	var result *uint32
	*a1 = *memmap.PtrUint32(6112660, 528276)
	*a2 = *memmap.PtrUint32(6112660, 528280)
	result = a4
	*a3 = *memmap.PtrUint32(6112660, 528284)
	*a4 = *memmap.PtrUint32(6112660, 528288)
	return result
}
func sub_41E5C0(a1 int32) int32 {
	var result int32
	result = a1
	*memmap.PtrUint32(0x587000, 59168) = uint32(a1)
	return result
}
func sub_41E6C0(a1 *byte) *uint32 {
	var (
		v1 int32
		v2 *uint32
	)
	v1 = int32(dword_5d4594_529316)
	if dword_5d4594_529316 == 0 {
		return nil
	}
	for {
		v2 = *(**uint32)(unsafe.Pointer(uintptr(v1 + 28)))
		if v2 != nil {
			break
		}
	LABEL_5:
		v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 36))))
		if v1 == 0 {
			return nil
		}
	}
	for libc.StrCaseCmp(a1, (*byte)(unsafe.Pointer(uintptr(*v2+52)))) != 0 {
		v2 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*10)))))
		if v2 == nil {
			goto LABEL_5
		}
	}
	return (*uint32)(unsafe.Pointer(uintptr(v1)))
}
func sub_41E710(a1 *libc.WChar, a2 int32) {
	var v2 *uint32
	v2 = sub_41E6C0((*byte)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(a1)) + 52))))
	if *(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*5)) == uint32(sub_41F360()) {
		sub_4484D0((*libc.WChar)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(libc.WChar(0))*2)), int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(a1)) + 12)))), a2)
	}
}
func sub_41E750(a1 *libc.WChar) int32 {
	var (
		v1 int32
		v2 int32
	)
	v1 = int32(dword_5d4594_529316)
	if dword_5d4594_529316 == 0 {
		return 0
	}
	for {
		v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 28))))
		if v2 != 0 {
			break
		}
	LABEL_5:
		v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 36))))
		if v1 == 0 {
			return 0
		}
	}
	for nox_wcscmp(a1, (*libc.WChar)(unsafe.Pointer(uintptr(v2+4)))) != 0 {
		v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 40))))
		if v2 == 0 {
			goto LABEL_5
		}
	}
	return v2
}
func sub_41E7A0(a1 *byte) *uint32 {
	var (
		v1 int32
		v2 *uint32
	)
	v1 = int32(dword_5d4594_529316)
	if dword_5d4594_529316 == 0 {
		return nil
	}
	for {
		v2 = *(**uint32)(unsafe.Pointer(uintptr(v1 + 28)))
		if v2 != nil {
			break
		}
	LABEL_5:
		v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 36))))
		if v1 == 0 {
			return nil
		}
	}
	for libc.StrCmp(a1, (*byte)(unsafe.Pointer(uintptr(*v2+52)))) != 0 {
		v2 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*10)))))
		if v2 == nil {
			goto LABEL_5
		}
	}
	return (*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*1))
}
func sub_41E810() *uint32 {
	var (
		v0     *uint32
		v1     *uint32
		result *uint32
		v3     int32
	)
	v0 = (*uint32)(sub_41E990(CString("Chat Channels")))
	v1 = v0
	result = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v0), unsafe.Sizeof(uint32(0))*7)))))
	if result != nil {
		for *(*uint32)(unsafe.Pointer(uintptr(*result + 12))) >= 25 {
			result = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*10)))))
			if result == nil {
				goto LABEL_4
			}
		}
	} else {
	LABEL_4:
		v3 = noxRndCounter1.IntClamp(0, int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*6))-1))
		result = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*7)))))
		if result != nil {
			for v3 != 0 {
				result = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*10)))))
				v3--
				if result == nil {
					goto LABEL_7
				}
			}
		} else {
		LABEL_7:
			result = nil
		}
	}
	return result
}
func sub_41E870(a1 *byte, a2 *uint32) *uint32 {
	var (
		v2     int32
		v3     *uint32
		v4     int32
		result *uint32
	)
	v2 = int32(dword_5d4594_529316)
	if dword_5d4594_529316 == 0 {
		return nil
	}
	for {
		v3 = *(**uint32)(unsafe.Pointer(uintptr(v2 + 28)))
		v4 = 0
		if v3 != nil {
			break
		}
	LABEL_5:
		v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 36))))
		if v2 == 0 {
			return nil
		}
	}
	for libc.StrCaseCmp(a1, (*byte)(unsafe.Pointer(uintptr(*v3+52)))) != 0 {
		v3 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*10)))))
		v4++
		if v3 == nil {
			goto LABEL_5
		}
	}
	result = v3
	*a2 = uint32(v4)
	return result
}
func sub_41E8D0(a1 int32, a2 int32) int32 {
	var (
		v2 uint32
		v3 uint32
	)
	if a1 == 0 || a2 == 0 {
		return 0
	}
	v2 = *(*uint32)(unsafe.Pointer(uintptr(a1 + 16)))
	v3 = *(*uint32)(unsafe.Pointer(uintptr(a2 + 16)))
	if v2 > v3 {
		return -1
	}
	if v2 >= v3 {
		return int32(libc.StrCaseCmp((*byte)(unsafe.Pointer(uintptr(a1+52))), (*byte)(unsafe.Pointer(uintptr(a2+52)))))
	}
	return 1
}
func sub_41E910(a1 *byte) *byte {
	var (
		v1 *byte
		v2 int32
		v3 int8
	)
	v1 = a1
	v2 = 0
	if *a1 != 0 {
		for {
			if int32(uint8(*v1)) < 128 {
				*memmap.PtrUint8(6112660, uint32(func() int32 {
					p := &v2
					x := *p
					*p++
					return x
				}())+0x811A4) = uint8(*v1)
			}
			v3 = int8(*func() *byte {
				p := &v1
				*p = (*byte)(unsafe.Add(unsafe.Pointer(*p), 1))
				return *p
			}())
			if int32(v3) == 0 {
				break
			}
		}
	}
	*memmap.PtrUint8(6112660, uint32(v2)+0x811A4) = 0
	return (*byte)(memmap.PtrOff(6112660, 528804))
}
func sub_41E940(a1 int32) *uint32 {
	var result *uint32
	result = *(**uint32)(unsafe.Pointer(&dword_5d4594_529316))
	if dword_5d4594_529316 == 0 {
		return nil
	}
	for *(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*2)) != 0 || *(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*1)) != uint32(a1) {
		result = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*9)))))
		if result == nil {
			return nil
		}
	}
	return result
}
func sub_41E970(a1 int32) *uint32 {
	var result *uint32
	result = *(**uint32)(unsafe.Pointer(&dword_5d4594_529316))
	if dword_5d4594_529316 == 0 {
		return nil
	}
	for *(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*2)) != uint32(a1) {
		result = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*9)))))
		if result == nil {
			return nil
		}
	}
	return result
}
func sub_41E990(a1 *byte) unsafe.Pointer {
	var v1 int32
	v1 = int32(dword_5d4594_529316)
	if dword_5d4594_529316 == 0 {
		return nil
	}
	for libc.StrCaseCmp(*(**byte)(unsafe.Pointer(uintptr(v1 + 12))), a1) != 0 {
		v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 36))))
		if v1 == 0 {
			return nil
		}
	}
	return unsafe.Pointer(uintptr(v1))
}
func sub_41E9D0(a1 int32, a2 *byte, a3 *byte, a4 int32, a5 int32) *int32 {
	var (
		result *int32
		v6     *int32
		v7     *byte
		v8     *byte
	)
	result = nil
	if a2 != nil {
		result = (*int32)(alloc.Calloc(1, 44))
		v6 = result
		if result != nil {
			v7 = (*byte)(alloc.Calloc(libc.StrLen(a2)+1, 1))
			*(*int32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(int32(0))*3)) = int32(uintptr(unsafe.Pointer(v7)))
			if v7 != nil {
				libc.StrCpy(v7, a2)
			} else {
				*(*int32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(int32(0))*3)) = 0
			}
			*(*int32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(int32(0))*4)) = 0
			if a3 != nil {
				v8 = (*byte)(alloc.Calloc(libc.StrLen(a3)+1, 1))
				*(*int32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(int32(0))*4)) = int32(uintptr(unsafe.Pointer(v8)))
				if v8 != nil {
					libc.StrCpy(v8, a3)
				} else {
					*(*int32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(int32(0))*4)) = 0
				}
			}
			*v6 = a5
			*(*int32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(int32(0))*6)) = 0
			*(*int32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(int32(0))*7)) = 0
			*(*int32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(int32(0))*8)) = 0
			*(*int32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(int32(0))*9)) = 0
			*(*int32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(int32(0))*10)) = 0
			*(*int32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(int32(0))*1)) = a1
			*(*int32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(int32(0))*2)) = a4
			result = v6
		}
	}
	return result
}
func sub_41EAC0() *int32 {
	var (
		v0     *int32
		v1     int32
		result *int32
	)
	dword_5d4594_529340 = 0
	*memmap.PtrUint32(0x587000, 59616) = math.MaxUint32
	v0 = sub_41E9D0(0, CString("Chat Channels"), nil, 0, 1)
	dword_5d4594_529316 = uint32(uintptr(unsafe.Pointer(v0)))
	*(*int32)(unsafe.Add(unsafe.Pointer(v0), unsafe.Sizeof(int32(0))*5)) = 0
	dword_5d4594_529324 = uint32(uintptr(unsafe.Pointer(v0)))
	v1 = sub_420100()
	result = sub_41E9D0(0, CString("Nox"), nil, v1, 0)
	if result != nil {
		*(*int32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(int32(0))*9)) = int32(dword_5d4594_529316)
		if dword_5d4594_529316 != 0 {
			*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_529316 + 40))) = uint32(uintptr(unsafe.Pointer(result)))
		}
		dword_5d4594_529316 = uint32(uintptr(unsafe.Pointer(result)))
		*(*int32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(int32(0))*5)) = -1
	}
	*memmap.PtrUint32(0x587000, 59616) = math.MaxUint32
	dword_5d4594_529340 = 0
	return result
}
func sub_41EB40() int32 {
	var (
		v0     int32
		result int32
		v2     int32
	)
	v0 = int32(dword_5d4594_529316)
	result = sub_41EC30()
	if v0 != 0 {
		for {
			v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(v0 + 36))))
			if *(*uint32)(unsafe.Pointer(uintptr(v0 + 12))) != 0 {
				*(*unsafe.Pointer)(unsafe.Pointer(uintptr(v0 + 12))) = nil
			}
			if *(*uint32)(unsafe.Pointer(uintptr(v0 + 16))) != 0 {
				*(*unsafe.Pointer)(unsafe.Pointer(uintptr(v0 + 16))) = nil
			}
			alloc.Free(unsafe.Pointer(uintptr(v0)))
			v0 = v2
			if v2 == 0 {
				break
			}
		}
		dword_5d4594_529316 = 0
		dword_5d4594_529340 = 0
	} else {
		dword_5d4594_529316 = 0
		dword_5d4594_529340 = 0
	}
	return result
}
func sub_41EBB0(a1 int32) unsafe.Pointer {
	var (
		result unsafe.Pointer
		v2     *byte
	)
	if a1 != 0 {
		v2 = (*byte)(unsafe.Pointer(uintptr(a1 + 52)))
		a1 = 0
		result = unsafe.Pointer(sub_41E870(v2, (*uint32)(unsafe.Pointer(&a1))))
		dword_5d4594_529340 = uint32(uintptr(result))
		*memmap.PtrUint32(0x587000, 59616) = uint32(a1)
		if result != nil {
			result = unsafe.Pointer(sub_446CC0(int32(uintptr(result)) + 4))
		}
	} else {
		dword_5d4594_529340 = 0
		result = unsafe.Pointer(sub_446CC0(0))
	}
	return result
}
func sub_41EC00() int32 {
	var result int32
	if dword_5d4594_529340 != 0 {
		result = int32(dword_5d4594_529340 + 4)
	} else {
		result = 0
	}
	return result
}
func sub_41EC10() int32 {
	var result int32
	if dword_5d4594_529340 != 0 {
		result = int32(*memmap.PtrUint32(0x587000, 59616))
	} else {
		result = -1
	}
	return result
}
func sub_41EC30() int32 {
	var (
		v0     *uint32
		v1     int32
		v2     int32
		result int32
		v4     int32
	)
	v0 = *(**uint32)(unsafe.Pointer(&dword_5d4594_529316))
	dword_5d4594_529332 = 0
	dword_5d4594_529336 = 0
	if dword_5d4594_529316 != 0 {
		for {
			v1 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v0), unsafe.Sizeof(uint32(0))*7)))
			if v1 != 0 {
				for {
					v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 40))))
					if *(*uint32)(unsafe.Pointer(uintptr(v1))) != 0 {
						*(*unsafe.Pointer)(unsafe.Pointer(uintptr(v1))) = nil
					}
					alloc.Free(unsafe.Pointer(uintptr(v1)))
					v1 = v2
					if v2 == 0 {
						break
					}
				}
			}
			*(*uint32)(unsafe.Add(unsafe.Pointer(v0), unsafe.Sizeof(uint32(0))*7)) = 0
			*(*uint32)(unsafe.Add(unsafe.Pointer(v0), unsafe.Sizeof(uint32(0))*8)) = 0
			*(*uint32)(unsafe.Add(unsafe.Pointer(v0), unsafe.Sizeof(uint32(0))*6)) = 0
			v0 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v0), unsafe.Sizeof(uint32(0))*9)))))
			if v0 == nil {
				break
			}
		}
	}
	result = sub_41E2F0()
	if result == 7 {
		v4 = sub_41E2F0()
		result = sub_41DA70(v4, 10)
	}
	dword_5d4594_529340 = 0
	return result
}
func sub_41ECB0(a1 int32) {
	var (
		v1  int32
		v2  *uint32
		v3  *uint32
		v4  *byte
		v5  *libc.WChar
		v6  int32
		v7  unsafe.Pointer
		v8  int32
		v9  *libc.WChar
		v10 *int32
		v11 int32
		v12 int32
		v13 int32
		v14 uint32 = 0
	)
	v1 = a1
	v13 = 0
	if a1 != 0 && (sub_41E2F0() == 8 && *(*uint32)(unsafe.Pointer(uintptr(a1))) == uint32(sub_420100()) || sub_41E2F0() == 7 && *(*uint32)(unsafe.Pointer(uintptr(a1))) == 0) {
		if sub_41E2F0() == 8 {
			v12 = sub_420100()
			v2 = sub_41E970(v12)
		} else {
			v2 = sub_41E970(0)
		}
		v3 = v2
		if v2 != nil {
			if *(*uint32)(unsafe.Pointer(uintptr(a1))) == 0 {
				v4 = sub_420110()
				v14 = uint32(libc.StrLen(v4))
				if libc.StrNCmp((*byte)(unsafe.Pointer(uintptr(a1+52))), v4, int(v14)) != 0 {
					return
				}
				nox_sprintf((*byte)(unsafe.Pointer(uintptr(a1+215))), CString("zot0blat"))
			}
			v5 = (*libc.WChar)(alloc.Calloc(1, 48))
			v6 = 0
			if v5 != nil {
				v7 = alloc.Calloc(1, 268)
				*(*uint32)(unsafe.Pointer(v5)) = uint32(uintptr(v7))
				if v7 != nil {
					libc.MemCpy(v7, unsafe.Pointer(uintptr(a1)), 268)
					v1 = a1
					v6 = 0
				} else {
					*(*uint32)(unsafe.Pointer(v5)) = 0
				}
				sub_41E910((*byte)(unsafe.Pointer(uintptr(v1 + 52))))
				if *(*uint32)(unsafe.Pointer(uintptr(v1))) == 0 {
					v8 = int32(libc.Atoi(libc.GoString((*byte)(unsafe.Pointer(uintptr(v14 + uint32(v1) + 52))))))
					v9 = sub_4469D0(v8)
					nox_swprintf((*libc.WChar)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(libc.WChar(0))*2)), v9)
					*(*libc.WChar)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(libc.WChar(0))*19)) = 0
				}
				v10 = (*int32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*7)))))
				if v10 != nil {
					for sub_41E8D0(a1, *v10) > 0 {
						v10 = (*int32)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(v10), unsafe.Sizeof(int32(0))*10)))))
						v6++
						if v10 == nil {
							v13 = v6
							*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*8)) + 40))) = uint32(uintptr(unsafe.Pointer(v5)))
							*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v5))), unsafe.Sizeof(uint32(0))*11))) = *(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*8))
							*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v5))), unsafe.Sizeof(uint32(0))*10))) = 0
							*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*8)) = uint32(uintptr(unsafe.Pointer(v5)))
							goto LABEL_24
						}
					}
					v13 = v6
					v11 = *(*int32)(unsafe.Add(unsafe.Pointer(v10), unsafe.Sizeof(int32(0))*11))
					if v11 != 0 {
						*(*uint32)(unsafe.Pointer(uintptr(v11 + 40))) = uint32(uintptr(unsafe.Pointer(v5)))
					} else {
						*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*7)) = uint32(uintptr(unsafe.Pointer(v5)))
					}
					*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v5))), unsafe.Sizeof(uint32(0))*11))) = uint32(*(*int32)(unsafe.Add(unsafe.Pointer(v10), unsafe.Sizeof(int32(0))*11)))
					*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v5))), unsafe.Sizeof(uint32(0))*10))) = uint32(uintptr(unsafe.Pointer(v10)))
					*(*int32)(unsafe.Add(unsafe.Pointer(v10), unsafe.Sizeof(int32(0))*11)) = int32(uintptr(unsafe.Pointer(v5)))
				} else {
					*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*8)) = uint32(uintptr(unsafe.Pointer(v5)))
					*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*7)) = uint32(uintptr(unsafe.Pointer(v5)))
					*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v5))), unsafe.Sizeof(uint32(0))*11))) = 0
					*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v5))), unsafe.Sizeof(uint32(0))*10))) = 0
				}
			LABEL_24:
				dword_5d4594_529332++
				dword_5d4594_529336 += *(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(v5)) + 12)))
				*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*6))++
				if sub_41E2F0() == 7 {
					sub_41E710(v5, v13)
				}
			}
		}
	}
}
func sub_41EEA0(a1 int32) {
	var (
		v1 *uint32
		v2 *unsafe.Pointer
		v3 *uint32
		v4 *uint32
		v5 int32
	)
	v5 = 0
	if a1 != 0 {
		v1 = sub_41E940(int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 16)))))
		if v1 != nil {
			v2 = (*unsafe.Pointer)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*7)))))
			if v2 != nil {
				for libc.StrCmp((*byte)(unsafe.Add(unsafe.Pointer((*byte)(*v2)), 52)), (*byte)(unsafe.Pointer(uintptr(a1+52)))) != 0 {
					v2 = (*unsafe.Pointer)(*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(unsafe.Pointer(nil))*10)))
					v5++
					if v2 == nil {
						return
					}
				}
				v3 = (*uint32)(*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(unsafe.Pointer(nil))*11)))
				if v3 != nil {
					*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*10)) = uint32(uintptr(*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(unsafe.Pointer(nil))*10))))
				} else {
					*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*7)) = uint32(uintptr(*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(unsafe.Pointer(nil))*10))))
				}
				v4 = (*uint32)(*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(unsafe.Pointer(nil))*10)))
				if v4 != nil {
					*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*11)) = uint32(uintptr(*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(unsafe.Pointer(nil))*11))))
				} else {
					*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*8)) = uint32(uintptr(*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(unsafe.Pointer(nil))*11))))
				}
				dword_5d4594_529332--
				dword_5d4594_529336 -= *(*uint32)(unsafe.Pointer(uintptr(a1 + 12)))
				*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*6))--
				if sub_41E2F0() == 7 {
					sub_448550(v5)
				}
				*v2 = nil
				alloc.Free(unsafe.Pointer(v2))
			}
		}
	}
}
func sub_41EFB0(a1 int32) *libc.WChar {
	var (
		i      int32
		result *libc.WChar
		v3     [20]byte
	)
	sub_446970_wol_chat()
	if dword_5d4594_529340 != 0 {
		libc.StrCpy(&v3[0], (*byte)(unsafe.Pointer(uintptr(**(**uint32)(unsafe.Pointer(&dword_5d4594_529340))+52))))
		dword_5d4594_529340 = 0
	} else {
		v3[0] = 0
	}
	sub_41EC30()
	for i = a1; i != 0; i = int32(*(*uint32)(unsafe.Pointer(uintptr(i + 48)))) {
		sub_41ECB0(i)
	}
	if v3[0] != 0 {
		result = (*libc.WChar)(unsafe.Pointer(sub_41E870(&v3[0], (*uint32)(memmap.PtrOff(0x587000, 59616)))))
		dword_5d4594_529340 = uint32(uintptr(unsafe.Pointer(result)))
	} else {
		result = (*libc.WChar)(unsafe.Pointer(uintptr(sub_41E2F0())))
		if uintptr(unsafe.Pointer(result)) == uintptr(7) {
			result = (*libc.WChar)(unsafe.Pointer(sub_41E810()))
			dword_5d4594_529340 = uint32(uintptr(unsafe.Pointer(result)))
			if result != nil {
				result = (*libc.WChar)(unsafe.Pointer(uintptr(sub_446A20_wol_chat((*libc.WChar)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(libc.WChar(0))*2))))))
			}
		}
	}
	return result
}
func sub_41F070() *libc.WChar {
	var result *libc.WChar
	if dword_5d4594_529328 == 0 {
		return nil
	}
	result = sub_41F330(int32(dword_5d4594_529328+4), int32(*(*uint32)(unsafe.Pointer(uintptr(**(**uint32)(unsafe.Pointer(&dword_5d4594_529328)) + 12)))))
	dword_5d4594_529328 = *(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_529328 + 40)))
	return result
}
func sub_41F0A0() *libc.WChar {
	dword_5d4594_529328 = *(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_529324 + 28)))
	return sub_41F070()
}
func sub_41F0C0() int32 {
	var result int32
	result = 0
	if dword_5d4594_529328 != 0 {
		result = int32(**(**uint32)(unsafe.Pointer(&dword_5d4594_529328)))
		dword_5d4594_529328 = *(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_529328 + 40)))
	}
	return result
}
func sub_41F0E0() int32 {
	if dword_5d4594_529324 == 0 {
		return 0
	}
	dword_5d4594_529328 = *(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_529324 + 28)))
	return sub_41F0C0()
}
func sub_41F100(a1 *byte) int32 {
	var (
		result int32
		v2     *uint32
	)
	if a1 != nil {
		v2 = (*uint32)(sub_41E990(a1))
		if v2 != nil {
			dword_5d4594_529324 = uint32(uintptr(unsafe.Pointer(v2)))
			dword_5d4594_529328 = *(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*7))
			result = 1
		} else {
			result = 0
		}
	} else {
		dword_5d4594_529324 = dword_5d4594_529316
		result = 1
	}
	return result
}
func sub_41F140(a1 *libc.WChar) *libc.WChar {
	var result *libc.WChar
	result = a1
	if a1 != nil {
		result = (*libc.WChar)(unsafe.Pointer(uintptr(sub_41E750(a1))))
		if result != nil {
			result = (*libc.WChar)(unsafe.Pointer(uintptr(sub_40D350(int32(*(*uint32)(unsafe.Pointer(result)))))))
		}
	}
	return result
}
func sub_41F1C0(a1 int32) int32 {
	var v1 *uint32
	v1 = *(**uint32)(unsafe.Pointer(&dword_5d4594_529316))
	if dword_5d4594_529316 == 0 {
		return 0
	}
	for uint32(a1) != *(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*5)) {
		v1 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*9)))))
		if v1 == nil {
			return 0
		}
	}
	return int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*3)))
}
func sub_41F1E0(a1 int32) int32 {
	var (
		v1 int32
		v2 int32
	)
	v1 = 0
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_529324 + 28))))
	if v2 == 0 {
		return 0
	}
	for a1 != v1 {
		v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 40))))
		v1++
		if v2 == 0 {
			return 0
		}
	}
	return v2 + 4
}
func sub_41F210() int32 {
	return int32(dword_5d4594_529336)
}
func sub_41F220() int32 {
	return int32(dword_5d4594_529332)
}
func sub_41F230(a1 int32, a2 int32) {
	var (
		v2 int32
		v3 *int32
		v4 *libc.WChar
		v5 int32
		v6 *byte
	)
	if a1 != 0 {
		v6 = (*byte)(unsafe.Pointer(uintptr(a1 + 52)))
		a1 = 0
		v3 = (*int32)(unsafe.Pointer(sub_41E870(v6, (*uint32)(unsafe.Pointer(&a1)))))
		v4 = (*libc.WChar)(unsafe.Pointer(v3))
		if v3 != nil {
			v5 = a2
			*(*uint32)(unsafe.Pointer(uintptr(*v3 + 12))) += uint32(a2)
			dword_5d4594_529336 += uint32(v5)
			if *(*uint32)(unsafe.Pointer(uintptr(*v3 + 12))) == 0 {
				sub_41EEA0(*v3)
			}
			if sub_41E2F0() == 7 {
				sub_41E710(v4, a1)
			}
		}
	} else if dword_5d4594_529340 != 0 {
		v2 = a2
		*(*uint32)(unsafe.Pointer(uintptr(**(**uint32)(unsafe.Pointer(&dword_5d4594_529340)) + 12))) += uint32(a2)
		dword_5d4594_529336 += uint32(v2)
		if *(*uint32)(unsafe.Pointer(uintptr(**(**uint32)(unsafe.Pointer(&dword_5d4594_529340)) + 12))) != 0 {
			if sub_41E2F0() == 7 {
				sub_448550(*memmap.PtrInt32(0x587000, 59616))
				sub_41E710(*(**libc.WChar)(unsafe.Pointer(&dword_5d4594_529340)), *memmap.PtrInt32(0x587000, 59616))
			}
		} else {
			sub_41EEA0(int32(**(**uint32)(unsafe.Pointer(&dword_5d4594_529340))))
			dword_5d4594_529340 = 0
			*memmap.PtrUint32(0x587000, 59616) = math.MaxUint32
		}
	}
}
func sub_41F330(a1 int32, a2 int32) *libc.WChar {
	if a1 == 0 {
		return nil
	}
	nox_swprintf((*libc.WChar)(memmap.PtrOff(6112660, 528292)), libc.CWString("%s\t%d"), a1, a2)
	return (*libc.WChar)(memmap.PtrOff(6112660, 528292))
}
func sub_41F360() int32 {
	return int32(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_529324 + 20))))
}
func sub_41F370(a1 int32) *uint32 {
	var result *uint32
	result = *(**uint32)(unsafe.Pointer(&dword_5d4594_529316))
	if dword_5d4594_529316 != 0 {
		for *(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*5)) != uint32(a1) {
			result = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*9)))))
			if result == nil {
				goto LABEL_4
			}
		}
		dword_5d4594_529324 = uint32(uintptr(unsafe.Pointer(result)))
	} else {
	LABEL_4:
		dword_5d4594_529324 = dword_5d4594_529316
	}
	return result
}
func sub_41F440(a1 int32, a2 int8, a3 int32) {
	sub_448570((*byte)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1)))+36))), a2, a3)
}
func sub_41F460(a1 int32, a2 int32) int32 {
	var v2 int32
	if a1 == 0 || a2 == 0 {
		return 0
	}
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2))) & 1)
	if (*(*uint32)(unsafe.Pointer(uintptr(a1))) & 1) == 1 {
		if v2 == 0 {
			return -1
		}
	} else if v2 == 1 {
		return 1
	}
	return int32(libc.StrCaseCmp((*byte)(unsafe.Pointer(uintptr(a1+36))), (*byte)(unsafe.Pointer(uintptr(a2+36)))))
}
func sub_41F4B0() int32 {
	var (
		v0     int32
		v1     int32
		result int32
		v3     int32
	)
	v0 = int32(dword_5d4594_531648)
	if dword_5d4594_531648 != 0 {
		for {
			v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(v0 + 20))))
			if *(*uint32)(unsafe.Pointer(uintptr(v0))) != 0 {
				*(*unsafe.Pointer)(unsafe.Pointer(uintptr(v0))) = nil
			}
			alloc.Free(unsafe.Pointer(uintptr(v0)))
			v0 = v1
			if v1 == 0 {
				break
			}
		}
	}
	dword_5d4594_531648 = 0
	dword_5d4594_531652 = 0
	dword_5d4594_531656 = 0
	result = sub_41E2F0()
	if result == 7 {
		v3 = sub_41E2F0()
		result = sub_41DA70(v3, 11)
	}
	return result
}
func sub_41F520(a1 *byte) {
	var (
		v1 *uint32
		v2 *byte
		v3 *int32
		v4 int32
		v5 int32
	)
	v1 = (*uint32)(alloc.Calloc(1, 28))
	if v1 != nil {
		v2 = (*byte)(alloc.Calloc(1, 108))
		*v1 = uint32(uintptr(unsafe.Pointer(v2)))
		if v2 != nil {
			libc.MemCpy(unsafe.Pointer(v2), unsafe.Pointer(a1), 108)
		} else {
			*v1 = 0
		}
		*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*4)) = uint32(sub_427880((*byte)(unsafe.Add(unsafe.Pointer(a1), 36))))
		v3 = *(**int32)(unsafe.Pointer(&dword_5d4594_531648))
		v4 = 0
		if dword_5d4594_531648 != 0 {
			for sub_41F460(int32(uintptr(unsafe.Pointer(a1))), *v3) > 0 {
				v3 = (*int32)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(int32(0))*5)))))
				v4++
				if v3 == nil {
					goto LABEL_9
				}
			}
			if v3 == nil {
			LABEL_9:
				*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_531652 + 20))) = uint32(uintptr(unsafe.Pointer(v1)))
				*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*6)) = dword_5d4594_531652
				*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*5)) = 0
				dword_5d4594_531652 = uint32(uintptr(unsafe.Pointer(v1)))
				goto LABEL_10
			}
			v5 = *(*int32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(int32(0))*6))
			if v5 != 0 {
				*(*uint32)(unsafe.Pointer(uintptr(v5 + 20))) = uint32(uintptr(unsafe.Pointer(v1)))
			} else {
				dword_5d4594_531648 = uint32(uintptr(unsafe.Pointer(v1)))
			}
			*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*6)) = uint32(*(*int32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(int32(0))*6)))
			*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*5)) = uint32(uintptr(unsafe.Pointer(v3)))
			*(*int32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(int32(0))*6)) = int32(uintptr(unsafe.Pointer(v1)))
		LABEL_10:
			if sub_41E2F0() == 7 {
				sub_41F440(int32(uintptr(unsafe.Pointer(v1))), int8(uint8(*(*uint32)(unsafe.Pointer(uintptr(*v1))))), v4)
			}
		} else {
			dword_5d4594_531652 = uint32(uintptr(unsafe.Pointer(v1)))
			dword_5d4594_531648 = uint32(uintptr(unsafe.Pointer(v1)))
			*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*6)) = 0
			*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*5)) = 0
		}
	}
}
func sub_41F6F0(a1 int32) {
	var i int32
	sub_41F4B0()
	for i = a1; i != 0; i = int32(*(*uint32)(unsafe.Pointer(uintptr(i + 32)))) {
		sub_41F520((*byte)(unsafe.Pointer(uintptr(i))))
	}
}
func sub_41F710() *byte {
	var result *byte
	if dword_5d4594_531656 == 0 {
		return nil
	}
	result = sub_41F740(int32(**(**uint32)(unsafe.Pointer(&dword_5d4594_531656))+36), int8(uint8(***(***uint32)(unsafe.Pointer(&dword_5d4594_531656)))))
	dword_5d4594_531656 = *(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_531656 + 20)))
	return result
}
func sub_41F740(a1 int32, a2 int8) *byte {
	if a1 == 0 {
		return nil
	}
	nox_sprintf((*byte)(memmap.PtrOff(6112660, 531392)), CString("%c%s"), func() int32 {
		if (int32(a2) & 1) != 0 {
			return 64
		}
		return 32
	}(), a1)
	return (*byte)(memmap.PtrOff(6112660, 531392))
}
func sub_41F780() *byte {
	dword_5d4594_531656 = dword_5d4594_531648
	return sub_41F710()
}
func sub_41F790(a1 *byte) *uint32 {
	var v1 *uint32
	v1 = *(**uint32)(unsafe.Pointer(&dword_5d4594_531648))
	if a1 == nil || dword_5d4594_531648 == 0 {
		return nil
	}
	for libc.StrCmp((*byte)(unsafe.Pointer(uintptr(*v1+36))), a1) != 0 {
		v1 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*5)))))
		if v1 == nil {
			return nil
		}
	}
	return v1
}
func sub_41F800(a1 *byte) int32 {
	var (
		v1     *int32
		result int32
	)
	v1 = (*int32)(unsafe.Pointer(sub_41F790(a1)))
	if v1 != nil {
		result = *v1
	} else {
		result = 0
	}
	return result
}
func sub_41F820(a1 *byte) {
	sub_41F620(int32(uintptr(unsafe.Pointer(a1))))
	sub_41F520(a1)
}
func sub_41F840(a1 *byte, a2 int32) *uint32 {
	var result *uint32
	result = (*uint32)(unsafe.Pointer(a1))
	if a1 != nil {
		result = sub_41F790(a1)
		if result != nil {
			result = (*uint32)(unsafe.Pointer(uintptr(*result)))
			*result = uint32(a2)
		}
	}
	return result
}
func sub_41F860(a1 *byte, a2 *libc.WChar) int32 {
	var (
		v2 *byte
		v3 *uint32
		v5 [108]byte
	)
	*(*[108]byte)(unsafe.Pointer(&v5[0])) = [108]byte{}
	v2 = a1
	if a1 == nil || a2 == nil {
		return 0
	}
	if *a1 == 64 || *a1 == 32 {
		v2 = (*byte)(unsafe.Add(unsafe.Pointer(a1), 1))
	}
	v3 = sub_41F790(v2)
	if v3 == nil {
		return 0
	}
	libc.StrCpy(&v5[36], (*byte)(unsafe.Pointer(uintptr(*v3+36))))
	sub_448680(a2, (*byte)(memmap.PtrOff(6112660, 529344)))
	return sub_40D410(int32(uintptr(unsafe.Pointer(&v5[0]))), int32(uintptr(memmap.PtrOff(6112660, 529344))))
}
func sub_41F8F0(a1 *byte, a2 *libc.WChar) int32 {
	var (
		v2 *byte
		v3 *uint32
		v5 [108]byte
	)
	*(*[108]byte)(unsafe.Pointer(&v5[0])) = [108]byte{}
	v2 = a1
	if a1 == nil || a2 == nil {
		return 0
	}
	if *a1 == 64 || *a1 == 32 {
		v2 = (*byte)(unsafe.Add(unsafe.Pointer(a1), 1))
	}
	v3 = sub_41F790(v2)
	if v3 == nil {
		return 0
	}
	libc.StrCpy(&v5[36], (*byte)(unsafe.Pointer(uintptr(*v3+36))))
	sub_448680(a2, (*byte)(memmap.PtrOff(6112660, 529344)))
	return sub_40D500(int32(uintptr(unsafe.Pointer(&v5[0]))), int32(uintptr(memmap.PtrOff(6112660, 529344))))
}
func sub_41F980(a1 *byte, a2 int32) *uint32 {
	var result *uint32
	result = sub_41F790(a1)
	if result != nil {
		*(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*4)) = uint32(a2)
	}
	return result
}
func sub_41F9A0() int32 {
	var (
		v0 int32
		v1 int32
		v2 int32
		v4 int8
		v5 int8
	)
	v0 = 0
	v1 = 0
	for {
		v2 = v1 + 1
		if sub_40D980(v1+1, int32(uintptr(unsafe.Pointer(&v5))), int32(uintptr(unsafe.Pointer(&v4)))) == 0 {
			break
		}
		v1 = v2
		v0++
		if v2 >= 128 {
			break
		}
	}
	return v0
}
func sub_41F9E0(a1 int32) int32 {
	return bool2int(libc.StrCmp((*byte)(memmap.PtrOff(6112660, 534756)), (*byte)(unsafe.Pointer(uintptr(a1+36)))) == 0)
}
func sub_41FA40() *byte {
	return (*byte)(memmap.PtrOff(6112660, 534756))
}
func sub_41FA50(a1 *byte) {
	if a1 != nil {
		libc.StrCpy((*byte)(memmap.PtrOff(6112660, 534756)), a1)
	}
}
func sub_41FA80(a1 *byte) int32 {
	var result int32
	if a1 == nil || libc.StrLen(a1) >= 10 || *(*int32)(unsafe.Pointer(&dword_587000_60044)) == -1 {
		return 0
	}
	result = 1
	libc.StrCpy((*byte)(memmap.PtrOff(6112660, dword_587000_60044*24+531660)), a1)
	return result
}
func sub_41FAE0(a1 *byte) int32 {
	var (
		v1     int32
		v2     int32
		result int32
	)
	v1 = 0
	if *a1 != 0 {
		if a1 != nil && libc.StrLen(a1) < 10 && *(*int32)(unsafe.Pointer(&dword_587000_60044)) != -1 {
			libc.StrCpy((*byte)(memmap.PtrOff(6112660, dword_587000_60044*24+531670)), a1)
			v1 = 1
		}
		result = v1
	} else {
		v2 = int32(dword_587000_60044 * 24)
		*memmap.PtrUint8(6112660, uint32(v2)+531670) = 0
		*memmap.PtrUint32(6112660, uint32(v2)+0x81CE0) = 0
		result = 1
	}
	return result
}
func sub_41FB60() int32 {
	return int32(*memmap.PtrUint32(6112660, dword_587000_60044*24+0x81CE0))
}
func sub_41FB70(a1 int32) int32 {
	var result int32
	result = int32(dword_587000_60044 * 3)
	*memmap.PtrUint32(6112660, dword_587000_60044*24+0x81CE0) = uint32(a1)
	return result
}
func sub_41FB90(a1 int32, a2 *uint32, a3 *uint32) int32 {
	var result int32
	if a1 < 0 || a1 >= 128 {
		*a2 = 0
		*a3 = 0
		result = 0
	} else {
		*a2 = uint32(uintptr(memmap.PtrOff(6112660, uint32(a1*24)+531660)))
		*a3 = uint32(uintptr(memmap.PtrOff(6112660, uint32(a1*24)+531670)))
		result = 1
	}
	return result
}
func sub_41FBE0(a1 *uint32, a2 *uint32) int32 {
	var result int32
	if *(*int32)(unsafe.Pointer(&dword_587000_60044)) == -1 {
		return 0
	}
	*a1 = uint32(uintptr(memmap.PtrOff(6112660, dword_587000_60044*24+531660)))
	result = 1
	*a2 = uint32(uintptr(memmap.PtrOff(6112660, dword_587000_60044*24+531670)))
	return result
}
func sub_41FC20(a1 int32) int32 {
	if a1 < 0 || a1 > 128 {
		return 0
	}
	dword_587000_60044 = uint32(a1)
	return 1
}
func sub_41FC40() int32 {
	return int32(dword_587000_60044)
}
func sub_41FC50() int32 {
	var result int32
	result = 0
	if *(*int32)(unsafe.Pointer(&dword_587000_60044)) != -1 {
		result = sub_40D9C0(int32(dword_587000_60044+1), int32(uintptr(memmap.PtrOff(6112660, dword_587000_60044*24+531660))), int32(uintptr(memmap.PtrOff(6112660, dword_587000_60044*24+531670))), bool2int(*memmap.PtrUint32(6112660, dword_587000_60044*24+0x81CE0) == 0))
		if result != 0 {
			result = sub_41FCF0()
		}
	}
	return result
}
func sub_41FCA0(a1 *byte, a2 *byte) int32 {
	var v2 int32
	v2 = sub_41F9A0()
	if sub_41FC20(v2) == 0 {
		return 0
	}
	sub_41FA80(a1)
	sub_41FAE0(a2)
	*memmap.PtrUint32(6112660, dword_587000_60044*24+0x81CE0) = 0
	sub_41FC50()
	return 1
}
func sub_41FCF0() int32 {
	var (
		v0 int32
		v1 *uint8
		v2 int32
		v3 *byte
		v4 int32
		v6 *byte
		v7 *byte
	)
	v0 = 0
	dword_5d4594_534768 = 0
	v1 = (*uint8)(memmap.PtrOff(6112660, 531670))
	for {
		v2 = v0 + 1
		if sub_40D980(v0+1, int32(uintptr(unsafe.Pointer(&v6))), int32(uintptr(unsafe.Pointer(&v7)))) != 0 {
			v3 = v7
			libc.StrCpy((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(v1))), -10)), v6)
			libc.StrCpy((*byte)(unsafe.Pointer(v1)), v3)
			v4 = int32(dword_5d4594_534768)
			if *v3 != 0 {
				*(*uint32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v1), 10)))) = 1
			} else {
				*(*uint32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v1), 10)))) = 0
			}
			dword_5d4594_534768 = uint32(v4 + 1)
		} else {
			*((*uint8)(unsafe.Add(unsafe.Pointer(v1), -10))) = 0
			*v1 = 0
			*(*uint32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v1), 10)))) = 0
		}
		v1 = (*uint8)(unsafe.Add(unsafe.Pointer(v1), 24))
		v0 = v2
		if int32(uintptr(unsafe.Pointer(v1))) > int32(uintptr(memmap.PtrOff(6112660, 534742))) {
			break
		}
	}
	if dword_5d4594_534768 == 0 {
		sub_41FC20(0)
	}
	return 1
}
func nox_xxx_officialStringCmp_41FDE0() int32 {
	var (
		v0 int32
		v1 uint32
		v3 *byte
		v4 *byte
	)
	libc.MemSet(memmap.PtrOff(0x85B3FC, 10308), 0, 108)
	if sub_41FBE0((*uint32)(unsafe.Pointer(&v3)), (*uint32)(unsafe.Pointer(&v4))) != 1 {
		return 0
	}
	v0 = sub_4207E0()
	if v0 != 0 {
		libc.StrCpy((*byte)(unsafe.Pointer(uintptr(v0+228))), v3)
		libc.StrCpy((*byte)(unsafe.Pointer(uintptr(v0+238))), v4)
		sub_40D2C0(v0, 30, 0)
		sub_41FA50(v3)
		v1 = uint32(libc.StrCSpn((*byte)(unsafe.Pointer(uintptr(v0+24))), CString(":")))
		if libc.StrNCmp(CString("Official"), (*byte)(unsafe.Pointer(uintptr(v1+uint32(v0)+25))), 8) == 0 {
			noxflags.SetGame(noxflags.GameFlag25)
		} else {
			noxflags.UnsetGame(noxflags.GameFlag25)
		}
	}
	if dword_5d4594_830276 == 0 {
		sub_41FEE0()
	}
	return 1
}
func sub_41FEE0() int32 {
	sub_40D9C0(int32(dword_587000_60044+1), int32(uintptr(memmap.PtrOff(6112660, 534776))), int32(uintptr(memmap.PtrOff(6112660, 534772))), 0)
	return sub_40DB80(int32(dword_587000_60044+1), 0)
}
func sub_41FF10() int32 {
	sub_40D440()
	sub_41F4B0()
	sub_41EC30()
	sub_446490(0)
	return 1
}
func sub_41FF30(a1 unsafe.Pointer) {
	if a1 != nil {
		sub_446520(0, a1, int32(libc.StrLen((*byte)(a1))+1))
	}
}
func sub_41FF60() int32 {
	return int32(*memmap.PtrUint32(6112660, 534800))
}
func sub_41FF70(a1 int32) int32 {
	*memmap.PtrUint32(6112660, 534800) = uint32(a1)
	return sub_40D890(a1)
}
func sub_41FF90() int32 {
	return int32(dword_587000_60064)
}
func sub_41FFA0(a1 int32) int32 {
	dword_587000_60064 = uint32(a1)
	return sub_40D7A0(a1, *(*int32)(unsafe.Pointer(&dword_587000_60068)))
}
func sub_41FFC0() int32 {
	return int32(dword_587000_60068)
}
func sub_41FFD0(a1 int32) int32 {
	dword_587000_60068 = uint32(a1)
	return sub_40D7A0(*(*int32)(unsafe.Pointer(&dword_587000_60064)), a1)
}
func sub_41FFF0() int32 {
	sub_40D890(*memmap.PtrInt32(6112660, 534800))
	return sub_40D7A0(*(*int32)(unsafe.Pointer(&dword_587000_60064)), *(*int32)(unsafe.Pointer(&dword_587000_60068)))
}
func sub_420020() int32 {
	return 65540
}
func sub_420030() int32 {
	var (
		result    int32
		phkResult HKEY
		Type      uint32
		Data      [4]uint8
		cbData    uint32
		SubKey    [128]byte
	)
	libc.StrCpy(&SubKey[0], CString("SOFTWARE\\Westwood\\Nox"))
	cbData = 4
	*memmap.PtrUint32(0x587000, 60072) = math.MaxUint32
	result = compatRegOpenKeyExA((HKEY)(unsafe.Pointer(uintptr(1))), &SubKey[0], 0, 0xF003F, &phkResult)
	if result == 0 {
		if compatRegQueryValueExA(phkResult, CString("SKU"), nil, &Type, &Data[0], &cbData) == 0 && Type == 4 {
			*memmap.PtrUint32(0x587000, 60072) = *(*uint32)(unsafe.Pointer(&Data[0]))
			nox_sprintf((*byte)(memmap.PtrOff(6112660, 534780)), CString("%s%d_"), "Lob_", *(*uint32)(unsafe.Pointer(&Data[0]))>>8)
		}
		result = compatRegCloseKey(phkResult)
	}
	return result
}
func sub_4200E0() int32 {
	return int32(*memmap.PtrUint32(0x587000, 60072))
}
func sub_4200F0() int32 {
	return int32(*memmap.PtrUint8(0x587000, 60072)) & 15
}
func sub_420100() int32 {
	return int32(*memmap.PtrUint32(0x587000, 60072) >> 8)
}
func sub_420110() *byte {
	return (*byte)(memmap.PtrOff(6112660, 534780))
}
func sub_4201B0(a1 *uint32) {
	var (
		i  *uint32
		v2 *uint32
	)
	for i = a1; i != nil; i = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(i), unsafe.Sizeof(uint32(0))*5))))) {
		v2 = (*uint32)(alloc.Calloc(1, 248))
		if v2 != nil {
			libc.MemCpy(unsafe.Pointer(v2), unsafe.Pointer(i), 248)
			*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*5)) = 0
			if dword_5d4594_534808 != 0 {
				*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*5)) = dword_5d4594_534808
			}
			dword_5d4594_534808 = uint32(uintptr(unsafe.Pointer(v2)))
		}
	}
}
func sub_420200() unsafe.Pointer {
	var v0 int32
	v0 = int32(dword_5d4594_534808)
	if dword_5d4594_534808 == 0 {
		return nil
	}
	for libc.StrCaseCmp((*byte)(unsafe.Pointer(uintptr(v0+24))), CString("chat server")) != 0 {
		v0 = int32(*(*uint32)(unsafe.Pointer(uintptr(v0 + 20))))
		if v0 == 0 {
			return nil
		}
	}
	return unsafe.Pointer(uintptr(v0))
}
func sub_420230(a1 *byte, a2 *uint16) int32 {
	var (
		v2 int32
		v3 *byte
		v4 *byte
		v6 [128]byte
	)
	if a1 == nil {
		return 0
	}
	if a2 == nil {
		return 0
	}
	v2 = int32(dword_5d4594_534808)
	if dword_5d4594_534808 == 0 {
		return 0
	}
	for {
		if libc.StrCaseCmp((*byte)(unsafe.Pointer(uintptr(v2+95))), CString("LAD")) == 0 {
			libc.StrCpy(&v6[0], (*byte)(unsafe.Pointer(uintptr(v2+100))))
			*a1 = 0
			*a2 = 0
			libc.StrTok(&v6[0], CString(";"))
			v3 = libc.StrTok(nil, CString(";"))
			if v3 != nil {
				libc.StrCpy(a1, v3)
			}
			v4 = libc.StrTok(nil, CString(";"))
			if v4 != nil {
				*a2 = uint16(int16(libc.Atoi(libc.GoString(v4))))
			}
			if *a1 != 0 && int32(*a2) != 0 {
				break
			}
		}
		v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 20))))
		if v2 == 0 {
			return 0
		}
	}
	return 1
}
func sub_420360(a1 *byte, a2 *uint16) int32 {
	var (
		v2 int32
		v3 *byte
		v4 *byte
		v5 *byte
		v7 [128]byte
	)
	if a1 == nil {
		return 0
	}
	if a2 == nil {
		return 0
	}
	*a1 = 0
	*a2 = 0
	v2 = int32(dword_5d4594_534808)
	if dword_5d4594_534808 == 0 {
		return 0
	}
	for libc.StrCaseCmp((*byte)(unsafe.Pointer(uintptr(v2+95))), CString("GAM")) != 0 {
	LABEL_19:
		v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 20))))
		if v2 == 0 {
			return 0
		}
	}
	if noxflags.HasGame(noxflags.GameModeQuest) {
		if libc.StrCaseCmp((*byte)(unsafe.Pointer(uintptr(v2+24))), CString("Quest gameres server")) == 0 {
			libc.StrCpy(&v7[0], (*byte)(unsafe.Pointer(uintptr(v2+100))))
			*a1 = 0
			*a2 = 0
			libc.StrTok(&v7[0], CString(";"))
			v3 = libc.StrTok(nil, CString(";"))
			if v3 != nil {
				libc.StrCpy(a1, v3)
			}
			v4 = libc.StrTok(nil, CString(";"))
			if v4 != nil {
			LABEL_16:
				*a2 = uint16(int16(libc.Atoi(libc.GoString(v4))))
				goto LABEL_17
			}
		}
	} else {
		if !noxflags.HasGame(noxflags.GameOnline) {
			return 0
		}
		if libc.StrCaseCmp((*byte)(unsafe.Pointer(uintptr(v2+24))), CString("Gameres server")) == 0 {
			libc.StrCpy(&v7[0], (*byte)(unsafe.Pointer(uintptr(v2+100))))
			*a1 = 0
			*a2 = 0
			libc.StrTok(&v7[0], CString(";"))
			v5 = libc.StrTok(nil, CString(";"))
			if v5 != nil {
				libc.StrCpy(a1, v5)
			}
			v4 = libc.StrTok(nil, CString(";"))
			if v4 != nil {
				goto LABEL_16
			}
		}
	}
LABEL_17:
	if *a1 == 0 || int32(*a2) == 0 {
		goto LABEL_19
	}
	return 1
}
func sub_420580() *uint32 {
	var (
		result *uint32
		v1     *uint32
	)
	result = *(**uint32)(unsafe.Pointer(&dword_5d4594_534808))
	if dword_5d4594_534808 != 0 {
		for {
			v1 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*5)))))
			alloc.Free(unsafe.Pointer(result))
			result = v1
			if v1 == nil {
				break
			}
		}
		dword_5d4594_534808 = 0
	} else {
		dword_5d4594_534808 = 0
	}
	return result
}
func sub_4205B0(a1 int32) *byte {
	var (
		v1 int32
		v2 *byte
		v3 int32
		v4 *byte
		v5 *byte
		v7 [68]byte
	)
	v1 = int32(dword_5d4594_534808)
	*memmap.PtrUint32(6112660, 534804) = uint32(a1)
	if dword_5d4594_534808 == 0 {
		return nil
	}
	for {
		if libc.StrCaseCmp((*byte)(unsafe.Pointer(uintptr(v1+95))), CString("IRC")) == 0 {
			v2 = (*byte)(unsafe.Pointer(uintptr(v1 + 24)))
			v3 = bool2int(unicode.IsDigit(rune(*(*uint8)(unsafe.Pointer(v2)))))
			if v3 != 0 {
				libc.StrCpy(&v7[0], v2)
				v4 = libc.StrTok(&v7[0], CString(":"))
				if v4 != nil {
					v5 = libc.StrTok(v4, CString(","))
					if v5 != nil {
						break
					}
				}
			}
		}
	LABEL_11:
		v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 20))))
		if v1 == 0 {
			return nil
		}
	}
	for libc.Atoi(libc.GoString(v5)) != int(a1) {
		v5 = libc.StrTok(nil, CString(","))
		if v5 == nil {
			goto LABEL_11
		}
	}
	return (*byte)(unsafe.Pointer(uintptr(v1)))
}
func sub_4206B0(a1 int32) int32 {
	var (
		v1 int32
		v2 *byte
		v3 int32
		v4 *byte
		v5 *byte
		v7 [68]byte
	)
	v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 20))))
	if v1 == 0 {
		return 0
	}
	for {
		if libc.StrCaseCmp((*byte)(unsafe.Pointer(uintptr(v1+95))), CString("IRC")) == 0 {
			v2 = (*byte)(unsafe.Pointer(uintptr(v1 + 24)))
			v3 = bool2int(unicode.IsDigit(rune(*(*uint8)(unsafe.Pointer(v2)))))
			if v3 != 0 {
				libc.StrCpy(&v7[0], v2)
				v4 = libc.StrTok(&v7[0], CString(":"))
				if v4 != nil {
					v5 = libc.StrTok(v4, CString(","))
					if v5 != nil {
						break
					}
				}
			}
		}
	LABEL_11:
		v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 20))))
		if v1 == 0 {
			return 0
		}
	}
	for libc.Atoi(libc.GoString(v5)) != int(*memmap.PtrUint32(6112660, 534804)) {
		v5 = libc.StrTok(nil, CString(","))
		if v5 == nil {
			goto LABEL_11
		}
	}
	return v1
}
func sub_4207A0(a1 int32) int32 {
	var (
		v1 int32
		i  *byte
	)
	v1 = 0
	for i = sub_4205B0(a1); i != nil; i = (*byte)(unsafe.Pointer(uintptr(sub_4206B0(int32(uintptr(unsafe.Pointer(i))))))) {
		v1++
	}
	return v1
}
func sub_4207D0(a1 int32) int32 {
	var result int32
	result = a1
	*memmap.PtrUint32(6112660, 534812) = uint32(a1)
	return result
}
func sub_4207E0() int32 {
	return int32(*memmap.PtrUint32(6112660, 534812))
}
func sub_4207F0(a1 int32) {
	var (
		v1 func()
		v2 func()
	)
	if a1 < 7 {
		v1 = cgoAsFunc(*(*uint32)(memmap.PtrOff(0x587000, *memmap.PtrUint32(0x85B3FC, 10288)*8+60280)), (*func())(nil))
		if v1 != nil {
			v1()
		}
		v2 = cgoAsFunc(*(*uint32)(memmap.PtrOff(0x587000, uint32(a1*8+0xEB7C))), (*func())(nil))
		*memmap.PtrUint32(0x85B3FC, 10288) = uint32(a1)
		if v2 != nil {
			v2()
		}
	}
}
func sub_420B70(a1 *uint32, a2 *uint32, a3 func(uint32, int32) int32, a4 int32) int32 {
	var v4 *uint32
	if a3 == nil {
		return 0
	}
	v4 = (*uint32)(unsafe.Pointer(uintptr(*a1)))
	if *a1 == 0 {
		return 0
	}
	for a3(*v4, a4) != 1 {
		v4 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*3)))))
		if v4 == nil {
			return 0
		}
	}
	*a2 = *(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*1))
	return int32(*v4)
}
func sub_420BE0(a1 int32, a2 *uint32) int32 {
	var (
		v2     *int32
		v3     bool
		result int32
	)
	v2 = *(**int32)(unsafe.Pointer(uintptr(a1)))
	v3 = *(*uint32)(unsafe.Pointer(uintptr(a1))) == 0
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 8))) = *(*uint32)(unsafe.Pointer(uintptr(a1)))
	if v3 {
		return 0
	}
	*a2 = uint32(*(*int32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(int32(0))*1)))
	result = *v2
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 8))) = uint32(*(*int32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(int32(0))*3)))
	return result
}
func sub_420C10(a1 int32, a2 *uint32) int32 {
	var (
		v2     *int32
		result int32
	)
	v2 = *(**int32)(unsafe.Pointer(uintptr(a1 + 8)))
	if v2 == nil {
		return 0
	}
	*a2 = uint32(*(*int32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(int32(0))*1)))
	result = *v2
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 8))) = uint32(*(*int32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(int32(0))*3)))
	return result
}
func sub_420C40(a1 int32, a2 int32) *uint32 {
	var result *uint32
	result = (*uint32)(alloc.Calloc(1, 12))
	if result != nil {
		*result = uint32(a1)
		*(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*1)) = uint32(a2)
		*(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*2)) = dword_5d4594_588068
		dword_5d4594_588068 = uint32(uintptr(unsafe.Pointer(result)))
	}
	return result
}
func sub_420C70() *uint32 {
	var (
		result *uint32
		v1     *uint32
	)
	result = *(**uint32)(unsafe.Pointer(&dword_5d4594_588068))
	if dword_5d4594_588068 != 0 {
		for {
			v1 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*2)))))
			alloc.Free(unsafe.Pointer(result))
			result = v1
			if v1 == nil {
				break
			}
		}
		dword_5d4594_588068 = 0
	} else {
		dword_5d4594_588068 = 0
	}
	return result
}
func nox_xxx_polygon_420CA0() *byte {
	var (
		v0 int32
		i  *uint8
	)
	v0 = 1
	if nox_xxx_polygonNextAngle_587000_60356 <= 1 {
		return nil
	}
	for i = (*uint8)(memmap.PtrOff(6112660, 535872)); *(*uint32)(unsafe.Pointer(i)) == 0; i = (*uint8)(unsafe.Add(unsafe.Pointer(i), 16)) {
		if uint32(func() int32 {
			p := &v0
			*p++
			return *p
		}()) >= uint32(*(*int32)(unsafe.Pointer(&nox_xxx_polygonNextAngle_587000_60356))) {
			return nil
		}
	}
	return (*byte)(memmap.PtrOff(6112660, uint32(v0*16)+0x82D24))
}
func nox_xxx_polygon_420CD0(a1 *uint32) *byte {
	var (
		v1 uint32
		i  *uint8
	)
	v1 = *a1 + 1
	if v1 >= uint32(*(*int32)(unsafe.Pointer(&nox_xxx_polygonNextAngle_587000_60356))) {
		return nil
	}
	for i = (*uint8)(memmap.PtrOff(6112660, v1*16+0x82D30)); *(*uint32)(unsafe.Pointer(i)) == 0; i = (*uint8)(unsafe.Add(unsafe.Pointer(i), 16)) {
		if func() uint32 {
			p := &v1
			*p++
			return *p
		}() >= uint32(*(*int32)(unsafe.Pointer(&nox_xxx_polygonNextAngle_587000_60356))) {
			return nil
		}
	}
	return (*byte)(memmap.PtrOff(6112660, v1*16+0x82D24))
}
func sub_420D10() int32 {
	var (
		result int32
		i      *uint8
	)
	result = 1
	if nox_xxx_polygonNextAngle_587000_60356 <= 1 {
		return int32(func() uint32 {
			p := &nox_xxx_polygonNextAngle_587000_60356
			x := *p
			*p++
			return x
		}())
	}
	for i = (*uint8)(memmap.PtrOff(6112660, 535872)); *(*uint32)(unsafe.Pointer(i)) != 0; i = (*uint8)(unsafe.Add(unsafe.Pointer(i), 16)) {
		if uint32(func() int32 {
			p := &result
			*p++
			return *p
		}()) >= uint32(*(*int32)(unsafe.Pointer(&nox_xxx_polygonNextAngle_587000_60356))) {
			return int32(func() uint32 {
				p := &nox_xxx_polygonNextAngle_587000_60356
				x := *p
				*p++
				return x
			}())
		}
	}
	return result
}
func nox_xxx_polygonSetAngle_420D40(a1 int32, a2 int32, a3 uint32, a4 int32) *uint32 {
	var (
		v4     bool
		result *uint32
	)
	if a4 != 0 {
		sub_420C40(int32(a3), a4)
	}
	v4 = a3 < uint32(*(*int32)(unsafe.Pointer(&nox_xxx_polygonNextAngle_587000_60356)))
	result = memmap.PtrUint32(6112660, a3*16+0x82D24)
	*result = a3
	if !v4 {
		nox_xxx_polygonNextAngle_587000_60356 = a3 + 1
	}
	*memmap.PtrUint32(6112660, a3*16+0x82D28) = uint32(a1)
	*memmap.PtrUint32(6112660, a3*16+0x82D2C) = uint32(a2)
	*memmap.PtrUint32(6112660, a3*16+0x82D30) = 1
	return result
}
func sub_420DA0(a1 float32, a2 float32) *uint32 {
	var (
		v2     *int32
		v3     *int32
		v4     int32
		result *uint32
		v6     int32
		v7     int32
		v8     int32
		v9     int32
		v10    int2
	)
	v10.field_0 = int(a1)
	v10.field_4 = int(a2)
	if noxflags.HasGame(noxflags.GameFlag22) && (func() bool {
		v2 = (*int32)(unsafe.Pointer(uintptr(sub_420E80(a1, a2, 900.0))))
		return (func() *int32 {
			v3 = v2
			return v3
		}()) != nil
	}()) {
		if sub_421B40((*uint32)(unsafe.Pointer(v2))) == 0 {
			v8 = *v3
			v9 = int32(*memmap.PtrUint16(6112660, 588072))
			*memmap.PtrUint16(6112660, 588072)++
			*memmap.PtrUint32(6112660, uint32(v9*4)+534820) = uint32(v8)
		}
	} else if !noxflags.HasGame(noxflags.GameFlag22) || nox_xxx_polygonIsPlayerInPolygon_4217B0(&v10, 0) == nil {
		v4 = sub_420D10()
		result = nox_xxx_polygonSetAngle_420D40(*(*int32)(unsafe.Add(unsafe.Pointer((*int32)(unsafe.Pointer(&a1))), unsafe.Sizeof(int32(0))*0)), *(*int32)(unsafe.Add(unsafe.Pointer((*int32)(unsafe.Pointer(&a2))), unsafe.Sizeof(int32(0))*0)), uint32(v4), 0)
		v6 = int32(*result)
		v7 = int32(*memmap.PtrUint16(6112660, 588072))
		*memmap.PtrUint16(6112660, 588072)++
		*memmap.PtrUint32(6112660, uint32(v7*4)+534820) = uint32(v6)
		return result
	}
	return nil
}
func sub_420E80(a1 float32, a2 float32, a3 float32) int32 {
	var (
		v3 int32
		v4 *uint8
		v5 int32
		v6 float64
		v7 float64
	)
	v3 = 0
	if nox_xxx_polygonNextAngle_587000_60356 > 1 {
		v4 = (*uint8)(memmap.PtrOff(6112660, 535864))
		v5 = int32(nox_xxx_polygonNextAngle_587000_60356 - 1)
		for {
			if *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v4))), unsafe.Sizeof(uint32(0))*2))) != 0 {
				v6 = float64(*((*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(v4))), unsafe.Sizeof(float32(0))*1))) - a2)
				v7 = v6*v6 + float64((*(*float32)(unsafe.Pointer(v4))-a1)*(*(*float32)(unsafe.Pointer(v4))-a1))
				if v7 < float64(a3) {
					a3 = float32(v7)
					v3 = int32(uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v4), -4)))))
				}
			}
			v4 = (*uint8)(unsafe.Add(unsafe.Pointer(v4), 16))
			v5--
			if v5 == 0 {
				break
			}
		}
	}
	return v3
}
func sub_420EE0(a1 int32) int32 {
	var result int32
	result = a1
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 12))) = 0
	return result
}
func sub_421430() unsafe.Pointer {
	var (
		v0     *uint8
		result unsafe.Pointer
	)
	v0 = (*uint8)(memmap.PtrOff(6112660, 552476))
	for i := int32(0); i < math.MaxUint8; i++ {
		if *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v0))), -int(unsafe.Sizeof(uint32(0))*27)))) != 0 {
			if *memmap.PtrUint32(6112660, 588076) != 0 {
				*((*unsafe.Pointer)(unsafe.Add(unsafe.Pointer((*unsafe.Pointer)(unsafe.Pointer(v0))), -int(unsafe.Sizeof(unsafe.Pointer(nil))*27)))) = nil
			}
			*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v0))), -int(unsafe.Sizeof(uint32(0))*27)))) = 0
		}
		result = *(*unsafe.Pointer)(unsafe.Pointer(v0))
		if *(*uint32)(unsafe.Pointer(v0)) != 0 {
			if *memmap.PtrUint32(6112660, 588076) != 0 {
				*(*unsafe.Pointer)(unsafe.Pointer(v0)) = nil
			}
			*(*uint32)(unsafe.Pointer(v0)) = 0
		}
		*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v0))), unsafe.Sizeof(uint16(0))*10))) = 0
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v0))), -int(unsafe.Sizeof(uint32(0))*7)))) = 0
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v0))), unsafe.Sizeof(uint32(0))*2))) = math.MaxUint32
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v0))), unsafe.Sizeof(uint32(0))*4))) = math.MaxUint32
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v0))), -int(unsafe.Sizeof(uint32(0))*6)))) = 0
		v0 = (*uint8)(unsafe.Add(unsafe.Pointer(v0), 140))
	}
	nox_xxx_polygonNextIdx_587000_60352 = 1
	return result
}
func sub_420EF0(a1 *uint32) int32 {
	var (
		v1 uint32
		v2 *uint8
		v3 int32
		v4 *uint32
		v5 int32
		i  *uint8
	)
	v1 = 1
	if nox_xxx_polygonNextIdx_587000_60352 > 1 {
		v2 = (*uint8)(memmap.PtrOff(6112660, 552496))
		for {
			if *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), -int(unsafe.Sizeof(uint32(0))*11)))) != 0 {
				v3 = 0
				if int32(*(*uint16)(unsafe.Pointer(v2))) != 0 {
					v4 = (*uint32)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), -int(unsafe.Sizeof(uint32(0))*5)))))))
					for *v4 != *a1 {
						v3++
						v4 = (*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*1))
						if v3 >= int32(*(*uint16)(unsafe.Pointer(v2))) {
							goto LABEL_12
						}
					}
					if int32(func() uint16 {
						p := (*uint16)(unsafe.Pointer(v2))
						*p--
						return *p
					}()) < 3 {
						sub_4212C0(int32(uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v2), -128))))))
					}
					libc.MemCpy(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), -int(unsafe.Sizeof(uint32(0))*5))))+uint32(v3*4))), unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), -int(unsafe.Sizeof(uint32(0))*5))))+uint32(v3*4)+4)), int((int32(*(*uint16)(unsafe.Pointer(v2)))-v3)*4))
				}
			}
		LABEL_12:
			v1++
			v2 = (*uint8)(unsafe.Add(unsafe.Pointer(v2), 140))
			if v1 >= uint32(*(*int32)(unsafe.Pointer(&nox_xxx_polygonNextIdx_587000_60352))) {
				break
			}
		}
	}
	v5 = 0
	if int32(*memmap.PtrUint16(6112660, 588072)) > 0 {
		for i = (*uint8)(memmap.PtrOff(6112660, 534820)); *(*uint32)(unsafe.Pointer(i)) != *a1; i = (*uint8)(unsafe.Add(unsafe.Pointer(i), 4)) {
			if func() int32 {
				p := &v5
				*p++
				return *p
			}() >= int32(*memmap.PtrUint16(6112660, 588072)) {
				return sub_420EE0(int32(uintptr(unsafe.Pointer(a1))))
			}
		}
		*memmap.PtrUint16(6112660, 588072)--
		libc.MemCpy(memmap.PtrOff(6112660, uint32(v5*4)+534820), memmap.PtrOff(6112660, uint32(v5*4)+0x82928), int((int32(*memmap.PtrUint16(6112660, 588072))-v5)*4))
	}
	return sub_420EE0(int32(uintptr(unsafe.Pointer(a1))))
}
func sub_421010() *byte {
	var result *byte
	result = (*byte)(memmap.PtrOff(6112660, 535872))
	for {
		*(*uint32)(unsafe.Pointer(result)) = 0
		result = (*byte)(unsafe.Add(unsafe.Pointer(result), 16))
		if int32(uintptr(unsafe.Pointer(result))) >= int32(uintptr(memmap.PtrOff(6112660, 552240))) {
			break
		}
	}
	nox_xxx_polygonNextAngle_587000_60356 = 1
	return result
}
func nox_xxx_polygonGetAngle_421030(a1 int32) *byte {
	return (*byte)(memmap.PtrOff(6112660, uint32(a1*16)+0x82D24))
}
func sub_421040(a1 int32) {
	var (
		v1 int32
		v2 *uint32
		v3 *uint32
		v4 int32
		v5 int32
		v6 *uint32
	)
	_ = v6
	v1 = 0
	if int32(*(*uint16)(unsafe.Pointer(uintptr(a1 + 128)))) != 0 {
		v2 = *(**uint32)(unsafe.Pointer(&dword_5d4594_588068))
		for {
			v3 = v2
			if v2 != nil {
				v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 108))))
				v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(v4 + v1*4))))
				v6 = (*uint32)(unsafe.Pointer(uintptr(v4 + v1*4)))
				for uint32(v5) != *(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*1)) {
					v3 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*2)))))
					if v3 == nil {
						goto LABEL_9
					}
				}
				*v6 = *v3
				v2 = *(**uint32)(unsafe.Pointer(&dword_5d4594_588068))
			}
		LABEL_9:
			v1++
			if v1 >= int32(*(*uint16)(unsafe.Pointer(uintptr(a1 + 128)))) {
				break
			}
		}
	}
}
func nox_xxx_polygonGetNext_4210A0() *byte {
	var (
		v0 int32
		i  *uint8
	)
	v0 = 1
	if nox_xxx_polygonNextIdx_587000_60352 <= 1 {
		return nil
	}
	for i = (*uint8)(memmap.PtrOff(6112660, 552452)); *(*uint32)(unsafe.Pointer(i)) == 0; i = (*uint8)(unsafe.Add(unsafe.Pointer(i), 140)) {
		if uint32(func() int32 {
			p := &v0
			*p++
			return *p
		}()) >= uint32(*(*int32)(unsafe.Pointer(&nox_xxx_polygonNextIdx_587000_60352))) {
			return nil
		}
	}
	return (*byte)(memmap.PtrOff(6112660, uint32(v0*140)+0x86D24))
}
func sub_4210E0(a1 int32) *byte {
	var (
		v1 uint32
		i  *uint8
	)
	v1 = *(*uint32)(unsafe.Pointer(uintptr(a1 + 80))) + 1
	if v1 >= uint32(*(*int32)(unsafe.Pointer(&nox_xxx_polygonNextIdx_587000_60352))) {
		return nil
	}
	for i = (*uint8)(memmap.PtrOff(6112660, v1*140+0x86D78)); *(*uint32)(unsafe.Pointer(i)) == 0; i = (*uint8)(unsafe.Add(unsafe.Pointer(i), 140)) {
		if func() uint32 {
			p := &v1
			*p++
			return *p
		}() >= uint32(*(*int32)(unsafe.Pointer(&nox_xxx_polygonNextIdx_587000_60352))) {
			return nil
		}
	}
	return (*byte)(memmap.PtrOff(6112660, v1*140+0x86D24))
}
func sub_421130() int32 {
	var (
		result int32
		i      *uint8
	)
	result = 1
	if nox_xxx_polygonNextIdx_587000_60352 <= 1 {
		return int32(func() uint32 {
			p := &nox_xxx_polygonNextIdx_587000_60352
			x := *p
			*p++
			return x
		}())
	}
	for i = (*uint8)(memmap.PtrOff(6112660, 552452)); *(*uint32)(unsafe.Pointer(i)) != 0; i = (*uint8)(unsafe.Add(unsafe.Pointer(i), 140)) {
		if uint32(func() int32 {
			p := &result
			*p++
			return *p
		}()) >= uint32(*(*int32)(unsafe.Pointer(&nox_xxx_polygonNextIdx_587000_60352))) {
			return int32(func() uint32 {
				p := &nox_xxx_polygonNextIdx_587000_60352
				x := *p
				*p++
				return x
			}())
		}
	}
	return result
}
func sub_421160(a1 int32) int32 {
	var result int32
	libc.StrCpy((*byte)(unsafe.Pointer(uintptr(a1+4))), (*byte)(memmap.PtrOff(0x587000, 60364)))
	*(*uint8)(unsafe.Pointer(uintptr(a1 + 104))) = *memmap.PtrUint8(0x587000, 60464)
	*(*uint8)(unsafe.Pointer(uintptr(a1 + 105))) = *memmap.PtrUint8(0x587000, 60465)
	*(*uint8)(unsafe.Pointer(uintptr(a1 + 106))) = *memmap.PtrUint8(0x587000, 60466)
	*(*uint8)(unsafe.Pointer(uintptr(a1 + 130))) = *memmap.PtrUint8(0x587000, 60490)
	result = 0
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 132))) = 0
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 136))) = 0
	return result
}
func sub_4211D0(a1 int32) int32 {
	var result int32
	libc.StrCpy((*byte)(memmap.PtrOff(0x587000, 60364)), (*byte)(unsafe.Pointer(uintptr(a1+4))))
	*memmap.PtrUint8(0x587000, 60464) = *(*uint8)(unsafe.Pointer(uintptr(a1 + 104)))
	*memmap.PtrUint8(0x587000, 60465) = *(*uint8)(unsafe.Pointer(uintptr(a1 + 105)))
	*memmap.PtrUint8(0x587000, 60466) = *(*uint8)(unsafe.Pointer(uintptr(a1 + 106)))
	*memmap.PtrUint8(0x587000, 60490) = *(*uint8)(unsafe.Pointer(uintptr(a1 + 130)))
	result = 0
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 132))) = 0
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 136))) = 0
	return result
}
func sub_421230() *uint8 {
	var (
		v0 int32
		v1 *uint8
		v2 *uint8
	)
	v0 = sub_421130()
	v1 = (*uint8)(memmap.PtrOff(6112660, uint32(v0*140)+0x86D24))
	*(*uint32)(unsafe.Pointer(v1)) = 0
	if noxflags.HasGame(noxflags.GameFlag22) {
		v2 = (*uint8)(alloc.Calloc(1, 256))
		*(*uint32)(unsafe.Pointer(v1)) = uint32(uintptr(unsafe.Pointer(v2)))
		if v2 == nil {
			alloc.Free(unsafe.Pointer(v1))
			return nil
		}
		*v2 = 0
		*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(v1)) + 128))) = 0
	}
	*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v1))), unsafe.Sizeof(uint32(0))*20))) = uint32(v0)
	compat_itoa(v0, (*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(v1))), 4)), 10)
	sub_421160(int32(uintptr(unsafe.Pointer(v1))))
	*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v1))), unsafe.Sizeof(uint32(0))*21))) = 1
	return (*uint8)(memmap.PtrOff(6112660, uint32(v0*140)+0x86D24))
}
func sub_4212C0(a1 int32) int32 {
	var (
		v1     int32
		v2     int32
		v3     uint32
		v4     int32
		v5     *uint8
		v6     int32
		v7     *uint32
		v8     int32
		v9     *uint8
		v10    *uint8
		result int32
		v12    int32
		v13    int32
	)
	v1 = a1
	v2 = 0
	v12 = 0
	if int32(*(*uint16)(unsafe.Pointer(uintptr(a1 + 128)))) != 0 {
		for {
			v3 = 1
			v4 = 0
			v13 = 0
			if nox_xxx_polygonNextIdx_587000_60352 <= 1 {
			LABEL_16:
				v8 = 0
				v9 = (*uint8)(memmap.PtrOff(6112660, *(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v1 + 108))) + uint32(v2*4))))*16+0x82D24))
				if int32(*memmap.PtrUint16(6112660, 588072)) > 0 {
					v10 = (*uint8)(memmap.PtrOff(6112660, 534820))
					for *(*uint32)(unsafe.Pointer(v10)) != *(*uint32)(unsafe.Pointer(v9)) {
						v8++
						v10 = (*uint8)(unsafe.Add(unsafe.Pointer(v10), 4))
						if v8 >= int32(*memmap.PtrUint16(6112660, 588072)) {
							goto LABEL_22
						}
					}
					*memmap.PtrUint16(6112660, 588072)--
					libc.MemCpy(memmap.PtrOff(6112660, uint32(v8*4)+534820), memmap.PtrOff(6112660, uint32(v8*4)+0x82928), int((int32(*memmap.PtrUint16(6112660, 588072))-v2)*4))
				}
			LABEL_22:
				sub_420EE0(int32(uintptr(unsafe.Pointer(v9))))
			} else {
				v5 = (*uint8)(memmap.PtrOff(6112660, 552496))
				for v4 == 0 {
					if *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v5))), -int(unsafe.Sizeof(uint32(0))*11)))) != 0 {
						if (*uint8)(unsafe.Add(unsafe.Pointer(v5), -128)) != (*uint8)(unsafe.Pointer(uintptr(v1))) {
							v6 = 0
							if int32(*(*uint16)(unsafe.Pointer(v5))) != 0 {
								v7 = (*uint32)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v5))), -int(unsafe.Sizeof(uint32(0))*5)))))))
								for *v7 != *(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v1 + 108))) + uint32(v2*4)))) {
									v6++
									v7 = (*uint32)(unsafe.Add(unsafe.Pointer(v7), unsafe.Sizeof(uint32(0))*1))
									if v6 >= int32(*(*uint16)(unsafe.Pointer(v5))) {
										v4 = v13
										goto LABEL_14
									}
								}
								v4 = 1
								v13 = 1
							}
						}
					}
				LABEL_14:
					v3++
					v5 = (*uint8)(unsafe.Add(unsafe.Pointer(v5), 140))
					if v3 >= uint32(*(*int32)(unsafe.Pointer(&nox_xxx_polygonNextIdx_587000_60352))) {
						v2 = v12
						if v4 != 0 {
							break
						}
						goto LABEL_16
					}
					v2 = v12
				}
			}
			v12 = func() int32 {
				p := &v2
				*p++
				return *p
			}()
			if v2 >= int32(*(*uint16)(unsafe.Pointer(uintptr(v1 + 128)))) {
				break
			}
		}
	}
	if *(*uint32)(unsafe.Pointer(uintptr(v1))) != 0 {
		*(*unsafe.Pointer)(unsafe.Pointer(uintptr(v1))) = nil
		*(*uint32)(unsafe.Pointer(uintptr(v1))) = 0
	}
	if *(*uint32)(unsafe.Pointer(uintptr(v1 + 108))) != 0 {
		*(*unsafe.Pointer)(unsafe.Pointer(uintptr(v1 + 108))) = nil
		*(*uint32)(unsafe.Pointer(uintptr(v1 + 108))) = 0
	}
	*(*uint16)(unsafe.Pointer(uintptr(v1 + 128))) = 0
	*(*uint32)(unsafe.Pointer(uintptr(v1 + 80))) = 0
	result = -1
	*(*uint32)(unsafe.Pointer(uintptr(v1 + 84))) = 0
	*(*uint32)(unsafe.Pointer(uintptr(v1 + 116))) = math.MaxUint32
	*(*uint32)(unsafe.Pointer(uintptr(v1 + 124))) = math.MaxUint32
	return result
}
func nox_xxx_polygonGetByIdx_4214A0(a1 int32) *byte {
	var result *byte
	if a1 == -559023410 {
		result = nil
	} else {
		result = (*byte)(memmap.PtrOff(6112660, uint32(a1*140)+0x86D24))
	}
	return result
}
func sub_4214D0() {
	var (
		v0 *uint8
		v1 int32
		v2 *uint8
		v3 int32
		v4 int32
		v5 bool
		v6 *uint8
	)
	if int32(*memmap.PtrUint16(6112660, 588072)) >= 3 {
		v0 = sub_421230()
		if v0 != nil {
			*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v0))), unsafe.Sizeof(uint32(0))*27))) = uint32(uintptr(alloc.Calloc(int(*memmap.PtrUint16(6112660, 588072)), 4)))
			v1 = 0
			*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v0))), unsafe.Sizeof(uint16(0))*64))) = *memmap.PtrUint16(6112660, 588072)
			if int32(*memmap.PtrUint16(6112660, 588072)) != 0 {
				for {
					*(*uint32)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v0))), unsafe.Sizeof(uint32(0))*27))) + uint32(v1*4)))) = *memmap.PtrUint32(6112660, uint32(v1*4)+534820)
					v1++
					if v1 >= int32(*memmap.PtrUint16(6112660, 588072)) {
						break
					}
				}
			}
			v2 = (*uint8)(memmap.PtrOff(6112660, **((**uint32)(unsafe.Add(unsafe.Pointer((**uint32)(unsafe.Pointer(v0))), unsafe.Sizeof((*uint32)(nil))*27)))*16+0x82D24))
			*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v0))), unsafe.Sizeof(uint32(0))*22))) = uint32(int(*((*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(v2))), unsafe.Sizeof(float32(0))*1)))))
			*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v0))), unsafe.Sizeof(uint32(0))*23))) = uint32(int(*((*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(v2))), unsafe.Sizeof(float32(0))*2)))))
			*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v0))), unsafe.Sizeof(uint32(0))*24))) = uint32(int(*((*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(v2))), unsafe.Sizeof(float32(0))*1)))))
			v3 = int(*((*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(v2))), unsafe.Sizeof(float32(0))*2))))
			v4 = 1
			v5 = int32(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v0))), unsafe.Sizeof(uint16(0))*64)))) <= 1
			*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v0))), unsafe.Sizeof(uint32(0))*25))) = uint32(v3)
			if !v5 {
				for {
					v6 = (*uint8)(memmap.PtrOff(6112660, *(*uint32)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v0))), unsafe.Sizeof(uint32(0))*27))) + uint32(v4*4))))*16+0x82D24))
					if float64(*((*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(v6))), unsafe.Sizeof(float32(0))*1)))) >= float64(*((*int32)(unsafe.Add(unsafe.Pointer((*int32)(unsafe.Pointer(v0))), unsafe.Sizeof(int32(0))*22)))) {
						if float64(*((*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(v6))), unsafe.Sizeof(float32(0))*1)))) > float64(*((*int32)(unsafe.Add(unsafe.Pointer((*int32)(unsafe.Pointer(v0))), unsafe.Sizeof(int32(0))*24)))) {
							*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v0))), unsafe.Sizeof(uint32(0))*24))) = uint32(int(*((*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(v6))), unsafe.Sizeof(float32(0))*1)))))
						}
					} else {
						*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v0))), unsafe.Sizeof(uint32(0))*22))) = uint32(int(*((*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(v6))), unsafe.Sizeof(float32(0))*1)))))
					}
					if float64(*((*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(v6))), unsafe.Sizeof(float32(0))*2)))) >= float64(*((*int32)(unsafe.Add(unsafe.Pointer((*int32)(unsafe.Pointer(v0))), unsafe.Sizeof(int32(0))*23)))) {
						if float64(*((*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(v6))), unsafe.Sizeof(float32(0))*2)))) > float64(*((*int32)(unsafe.Add(unsafe.Pointer((*int32)(unsafe.Pointer(v0))), unsafe.Sizeof(int32(0))*25)))) {
							*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v0))), unsafe.Sizeof(uint32(0))*25))) = uint32(int(*((*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(v6))), unsafe.Sizeof(float32(0))*2)))))
						}
					} else {
						*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v0))), unsafe.Sizeof(uint32(0))*23))) = uint32(int(*((*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(v6))), unsafe.Sizeof(float32(0))*2)))))
					}
					v4++
					if v4 >= int32(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v0))), unsafe.Sizeof(uint16(0))*64)))) {
						break
					}
				}
			}
			*memmap.PtrUint16(6112660, 588072) = 0
		}
	}
}
func nox_xxx_polygon_421660(a1 *int32, a2 int32) int32 {
	var (
		v2  int8
		v4  int32
		v5  int32
		v6  *uint32
		v7  *uint8
		v8  int32
		v9  uint16
		v10 uint16
		v11 *uint8
		v12 int4 = int4{}
		v16 int4
	)
	v2 = 0
	if a2 == 0 {
		return 0
	}
	v4 = *(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*1))
	v16.field_0 = *a1
	v16.field_4 = v4
	if (int32(*memmap.PtrUint8(6112660, 588080)) & 2) != 0 {
		v5 = 5888
	} else {
		v5 = 0
	}
	*memmap.PtrUint32(6112660, 588080)++
	v6 = *(**uint32)(unsafe.Pointer(uintptr(a2 + 108)))
	v16.field_8 = v5
	if (int32(*memmap.PtrUint8(6112660, 588080)) & 2) != 0 {
		v16.field_C = 5888
	} else {
		v16.field_C = 0
	}
	v7 = (*uint8)(memmap.PtrOff(6112660, *v6*16+0x82D24))
	v12.field_0 = int(*((*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(v7))), unsafe.Sizeof(float32(0))*1))))
	v8 = int(*((*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(v7))), unsafe.Sizeof(float32(0))*2))))
	v9 = *(*uint16)(unsafe.Pointer(uintptr(a2 + 128)))
	v10 = 1
	for v12.field_4 = v8; int32(v10) <= int32(v9); v10++ {
		v11 = (*uint8)(memmap.PtrOff(6112660, *(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a2 + 108))) + uint32((int32(v10)%int32(v9))*4))))*16+0x82D24))
		if int32(v10)&1 != 0 {
			v12.field_8 = int(*((*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(v11))), unsafe.Sizeof(float32(0))*1))))
			v12.field_C = int(*((*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(v11))), unsafe.Sizeof(float32(0))*2))))
		} else {
			v12.field_0 = int(*((*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(v11))), unsafe.Sizeof(float32(0))*1))))
			v12.field_4 = int(*((*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(v11))), unsafe.Sizeof(float32(0))*2))))
		}
		if sub_427C80(&v16, &v12) != 0 {
			v2++
		}
		v9 = *(*uint16)(unsafe.Pointer(uintptr(a2 + 128)))
	}
	return int32(v2) & 1
}
func nox_xxx_polygonGetIdxA_421790(a1 *int2, a2 int32) int32 {
	var (
		v2     *nox_player_polygon_check_data
		result int32
	)
	v2 = nox_xxx_polygonIsPlayerInPolygon_4217B0(a1, a2)
	if v2 != nil {
		result = v2.field_0[20]
	} else {
		result = 0
	}
	return result
}
func nox_xxx_polygonIsPlayerInPolygon_4217B0(a1 *int2, a2 int32) *nox_player_polygon_check_data {
	var (
		v2 int32
		v4 int32
		i  *uint8
		v6 int32
	)
	if a2 != 0 {
		if a2 != -559023410 {
			if *memmap.PtrUint32(6112660, uint32(a2*140)+0x86D78) != 0 {
				v2 = sub_4281F0(a1, (*int4)(memmap.PtrOff(6112660, uint32(a2*140)+0x86D7C)))
				if v2 != 0 {
					if nox_xxx_polygon_421660(&a1.field_0, int32(uintptr(memmap.PtrOff(6112660, uint32(a2*140)+0x86D24)))) != 0 {
						return (*nox_player_polygon_check_data)(memmap.PtrOff(6112660, uint32(a2*140)+0x86D24))
					}
				}
			}
		}
	}
	v4 = 1
	if nox_xxx_polygonNextIdx_587000_60352 <= 1 {
		return nil
	}
	for i = (*uint8)(memmap.PtrOff(6112660, 552448)); ; i = (*uint8)(unsafe.Add(unsafe.Pointer(i), 140)) {
		if *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(i))), unsafe.Sizeof(uint32(0))*1))) != 0 {
			if *(*uint32)(unsafe.Pointer(i)) != uint32(a2) {
				v6 = sub_4281F0(a1, (*int4)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(i), 8)))))
				if v6 != 0 {
					if nox_xxx_polygon_421660(&a1.field_0, int32(uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(i), -80)))))) != 0 {
						break
					}
				}
			}
		}
		if uint32(func() int32 {
			p := &v4
			*p++
			return *p
		}()) >= uint32(*(*int32)(unsafe.Pointer(&nox_xxx_polygonNextIdx_587000_60352))) {
			return nil
		}
	}
	return (*nox_player_polygon_check_data)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(i), -80))))
}
func sub_421880(a1 int32, a2 int32, a3 float32) int32 {
	var (
		v3 *uint8
		v4 int32
		v5 int32
		v6 int32
		v7 *uint8
		v9 int4 = int4{}
	)
	v3 = (*uint8)(memmap.PtrOff(6112660, **(**uint32)(unsafe.Pointer(uintptr(a2 + 108)))*16+0x82D24))
	v9.field_0 = int(*((*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(v3))), unsafe.Sizeof(float32(0))*1))))
	v4 = int(*((*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(v3))), unsafe.Sizeof(float32(0))*2))))
	v5 = 1
	v6 = int32(*(*uint16)(unsafe.Pointer(uintptr(a2 + 128))))
	v9.field_4 = v4
	if int32(uint16(int16(v6))) < 1 {
		return 0
	}
	for {
		v7 = (*uint8)(memmap.PtrOff(6112660, *(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a2 + 108))) + uint32((v5%v6)*4))))*16+0x82D24))
		if v5&1 != 0 {
			v9.field_8 = int(*((*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(v7))), unsafe.Sizeof(float32(0))*1))))
			v9.field_C = int(*((*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(v7))), unsafe.Sizeof(float32(0))*2))))
		} else {
			v9.field_0 = int(*((*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(v7))), unsafe.Sizeof(float32(0))*1))))
			v9.field_4 = int(*((*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(v7))), unsafe.Sizeof(float32(0))*2))))
		}
		if sub_427DF0(a1, &v9.field_0, a3) != 0 {
			break
		}
		v5++
		v6 = int32(*(*uint16)(unsafe.Pointer(uintptr(a2 + 128))))
		if v5 > int32(uint16(int16(v6))) {
			return 0
		}
	}
	return 1
}
func sub_421960(a1 int32, a2 float32, a3 int32) int32 {
	var (
		v3     *int32
		result int32
	)
	v3 = sub_421990((*int2)(unsafe.Pointer(uintptr(a1))), a2, a3)
	if v3 != nil {
		result = *(*int32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(int32(0))*20))
	} else {
		result = 0
	}
	return result
}
func sub_421990(a1 *int2, a2 float32, a3 int32) *int32 {
	var (
		i  *int32
		v5 int32
	)
	if a3 != 0 {
		if a3 != -559023410 {
			i = memmap.PtrInt32(6112660, uint32(a3*140)+0x86D24)
			if *memmap.PtrUint32(6112660, uint32(a3*140)+0x86D78) != 0 {
				if sub_421880(int32(uintptr(unsafe.Pointer(a1))), int32(uintptr(memmap.PtrOff(6112660, uint32(a3*140)+0x86D24))), a2) != 0 {
					return i
				}
			}
		}
	}
	v5 = 1
	if nox_xxx_polygonNextIdx_587000_60352 > 1 {
		for i = memmap.PtrInt32(6112660, 552368); *(*int32)(unsafe.Add(unsafe.Pointer(i), unsafe.Sizeof(int32(0))*21)) == 0 || *(*int32)(unsafe.Add(unsafe.Pointer(i), unsafe.Sizeof(int32(0))*20)) == a3 || sub_421880(int32(uintptr(unsafe.Pointer(a1))), int32(uintptr(unsafe.Pointer(i))), a2) == 0; i = (*int32)(unsafe.Add(unsafe.Pointer(i), unsafe.Sizeof(int32(0))*35)) {
			if uint32(func() int32 {
				p := &v5
				*p++
				return *p
			}()) >= uint32(*(*int32)(unsafe.Pointer(&nox_xxx_polygonNextIdx_587000_60352))) {
				return nil
			}
		}
		return i
	}
	return nil
}
func sub_421A30() int16 {
	var (
		result int16
		v1     uint32
		v2     int32
		v3     *uint8
		v4     int32
		v5     *uint32
		v6     bool
		v7     *uint8
		v8     int32
	)
	result = int16(*memmap.PtrUint16(6112660, 588072))
	if int32(*memmap.PtrUint16(6112660, 588072)) != 0 {
		v8 = 0
		if int32(*memmap.PtrUint16(6112660, 588072)) != 0 {
			v7 = (*uint8)(memmap.PtrOff(6112660, 534820))
			for {
				v1 = 1
				v2 = 0
				if nox_xxx_polygonNextIdx_587000_60352 <= 1 {
					sub_420EE0(int32(uintptr(memmap.PtrOff(6112660, *(*uint32)(unsafe.Pointer(v7))*16+0x82D24))))
				} else {
					v3 = (*uint8)(memmap.PtrOff(6112660, 552496))
					for v2 == 0 {
						if *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v3))), -int(unsafe.Sizeof(uint32(0))*11)))) != 0 {
							v4 = 0
							if int32(*(*uint16)(unsafe.Pointer(v3))) != 0 {
								v5 = (*uint32)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v3))), -int(unsafe.Sizeof(uint32(0))*5)))))))
								for *v5 != *(*uint32)(unsafe.Pointer(v7)) {
									v4++
									v5 = (*uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(uint32(0))*1))
									if v4 >= int32(*(*uint16)(unsafe.Pointer(v3))) {
										goto LABEL_14
									}
								}
								v2 = 1
							}
						}
					LABEL_14:
						v1++
						v3 = (*uint8)(unsafe.Add(unsafe.Pointer(v3), 140))
						if v1 >= uint32(*(*int32)(unsafe.Pointer(&nox_xxx_polygonNextIdx_587000_60352))) {
							if v2 != 0 {
								break
							}
							sub_420EE0(int32(uintptr(memmap.PtrOff(6112660, *(*uint32)(unsafe.Pointer(v7))*16+0x82D24))))
						}
					}
				}
				result = int16(v8 + 1)
				v6 = func() int32 {
					p := &v8
					*p++
					return *p
				}() < int32(*memmap.PtrUint16(6112660, 588072))
				v7 = (*uint8)(unsafe.Add(unsafe.Pointer(v7), 4))
				if !v6 {
					break
				}
			}
		}
		*memmap.PtrUint16(6112660, 588072) = 0
	}
	return result
}
func sub_421B10() *uint32 {
	var result *uint32
	sub_421430()
	sub_421010()
	result = sub_420C70()
	*memmap.PtrUint32(6112660, 588076) = 1
	*memmap.PtrUint16(6112660, 588072) = 0
	return result
}
func sub_421B40(a1 *uint32) int32 {
	var (
		v1 int32
		i  *uint8
	)
	v1 = 0
	if int32(*memmap.PtrUint16(6112660, 588072)) <= 0 {
		return 0
	}
	for i = (*uint8)(memmap.PtrOff(6112660, 534820)); *(*uint32)(unsafe.Pointer(i)) != *a1; i = (*uint8)(unsafe.Add(unsafe.Pointer(i), 4)) {
		if func() int32 {
			p := &v1
			*p++
			return *p
		}() >= int32(*memmap.PtrUint16(6112660, 588072)) {
			return 0
		}
	}
	return 1
}
func nox_xxx_polygonDrawColor_421B80() {
	var (
		v0 int32
		v1 *byte
		v2 *byte
		v3 *nox_player_polygon_check_data
		v4 *byte
		v5 int2
		v6 int2
	)
	v0 = int32(*memmap.PtrUint32(0x852978, 8))
	if *memmap.PtrUint32(0x852978, 8) != 0 {
		nox_xxx_getSomeCoods_435670(&v5)
		sub_435690((*uint32)(unsafe.Pointer(&v6)))
		v1 = (*byte)(unsafe.Pointer(noxServer.getPlayerByID(int32(*(*uint32)(unsafe.Pointer(uintptr(v0 + 128)))))))
		v2 = v1
		if v1 != nil {
			if nox_xxx_polygonNextIdx_587000_60352 <= 1 {
			LABEL_8:
				*(*byte)(unsafe.Add(unsafe.Pointer(v2), 3668)) = 1
				*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), unsafe.Sizeof(uint32(0))*915))) = 0
				v4 = nox_xxx_getAmbientColor_469BB0()
				sub_4349C0((*uint32)(unsafe.Pointer(v4)))
				return
			}
			if *((*int32)(unsafe.Add(unsafe.Pointer((*int32)(unsafe.Pointer(v1))), unsafe.Sizeof(int32(0))*915))) == -559023410 || v5.field_0 != v6.field_0 || v5.field_4 != v6.field_4 {
				v3 = nox_xxx_polygonIsPlayerInPolygon_4217B0(&v5, int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v1))), unsafe.Sizeof(uint32(0))*915)))))
				if v3 == nil {
					goto LABEL_8
				}
				if *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), unsafe.Sizeof(uint32(0))*915))) != uint32(v3.field_0[20]) {
					*(*byte)(unsafe.Add(unsafe.Pointer(v2), 3668)) = byte(*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v3.field_0[32]))), 2)))
					*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), unsafe.Sizeof(uint32(0))*915))) = uint32(v3.field_0[20])
					sub_434990(int32(*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v3.field_0[26]))), 0))), int32(*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v3.field_0[26]))), 1))), int32(*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v3.field_0[26]))), 2))))
				}
			}
		}
	}
}
func nox_xxx_questCheckSecretArea_421C70(a1p *nox_object_t) {
	var (
		a1  int32 = int32(uintptr(unsafe.Pointer(a1p)))
		v1  int32
		v2  int32
		v3  float32
		v4  *uint8
		v5  int32
		v6  int32
		v7  int32
		i   int32
		v9  int32
		v10 int32
		v11 int32
		v12 int32
		v13 int2
	)
	if a1 == 0 || (int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))&4) == 0 {
		return
	}
	v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))) + 276))))
	if int32(*(*uint8)(unsafe.Pointer(uintptr(v1 + 2064)))) == 31 {
		v11 = int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 3660))))
		if v11 == 0 {
			*(*uint32)(unsafe.Pointer(uintptr(v1 + 3664))) = 0
			goto LABEL_33
		}
		if v11 == -559023410 {
			goto LABEL_33
		}
		v4 = (*uint8)(memmap.PtrOff(6112660, uint32(v11*140)+0x86D24))
		goto LABEL_11
	}
	if *(*int32)(unsafe.Pointer(uintptr(v1 + 3664))) != -559023410 && *(*float32)(unsafe.Pointer(uintptr(a1 + 56))) == *(*float32)(unsafe.Pointer(uintptr(a1 + 72))) && *(*float32)(unsafe.Pointer(uintptr(a1 + 60))) == *(*float32)(unsafe.Pointer(uintptr(a1 + 76))) {
		return
	}
	v2 = int(*(*float32)(unsafe.Pointer(uintptr(a1 + 56))))
	v3 = *(*float32)(unsafe.Pointer(uintptr(a1 + 60)))
	v13.field_0 = v2
	v13.field_4 = int(v3)
	v4 = (*uint8)(unsafe.Pointer(nox_xxx_polygonIsPlayerInPolygon_4217B0(&v13, int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 3664)))))))
	if v4 != nil {
		goto LABEL_12
	}
	v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 3664))))
	if v5 != 0 && v5 != -559023410 {
		v4 = (*uint8)(memmap.PtrOff(6112660, uint32(v5*140)+0x86D24))
	LABEL_11:
		if v4 != nil {
		LABEL_12:
			v6 = int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 3664))))
			if uint32(v6) != *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v4))), unsafe.Sizeof(uint32(0))*20))) {
				if v6 != -559023410 {
					if v6 != 0 {
						v7 = v6 * 35
						if *memmap.PtrInt32(6112660, uint32(v7*4)+0x86DA0) != -1 {
							nox_xxx_scriptCallByEventBlock_502490(memmap.PtrInt32(6112660, uint32(v7*4)+0x86D9C), a1, 0, 29)
						}
					}
					if (uint32(1<<int32(*(*uint8)(unsafe.Pointer(uintptr(v1 + 2064)))))&*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v4))), unsafe.Sizeof(uint32(0))*34)))) == 0 && int32(*(*uint8)(unsafe.Add(unsafe.Pointer(v4), 132)))&1 != 0 && noxflags.HasGame(noxflags.GameModeQuest) {
						sub_4D61F0(a1)
						nox_xxx_netPriMsgToPlayer_4DA2C0((*nox_object_t)(unsafe.Pointer(uintptr(a1))), CString("GeneralPrint:SecretFound"), 0)
						nox_xxx_aud_501960(904, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
						for i = int32(uintptr(unsafe.Pointer(nox_xxx_getFirstPlayerUnit_4DA7C0()))); i != 0; i = int32(uintptr(unsafe.Pointer(nox_xxx_getNextPlayerUnit_4DA7F0((*nox_object_t)(unsafe.Pointer(uintptr(i))))))) {
							if i != a1 {
								nox_xxx_netInformTextMsg_4DA0F0(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(i + 748))) + 276))) + 2064)))), 20, (*int32)(unsafe.Pointer(uintptr(a1+36))))
							}
						}
						v9 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v4))), unsafe.Sizeof(uint32(0))*33))))
						*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v9))), 0)) = uint8(int8(v9 & 254))
						*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v4))), unsafe.Sizeof(uint32(0))*33))) = uint32(v9)
					}
					v10 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v4))), unsafe.Sizeof(uint32(0))*29))))
					*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v4))), unsafe.Sizeof(uint32(0))*34))) |= uint32(1 << int32(*(*uint8)(unsafe.Pointer(uintptr(v1 + 2064)))))
					if v10 != -1 {
						nox_xxx_scriptCallByEventBlock_502490((*int32)(unsafe.Add(unsafe.Pointer((*int32)(unsafe.Pointer(v4))), unsafe.Sizeof(int32(0))*28)), a1, 0, 28)
					}
				}
				*(*uint32)(unsafe.Pointer(uintptr(v1 + 3664))) = *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v4))), unsafe.Sizeof(uint32(0))*20)))
				*(*uint8)(unsafe.Pointer(uintptr(v1 + 3668))) = *(*uint8)(unsafe.Add(unsafe.Pointer(v4), 130))
			}
			return
		}
	}
LABEL_33:
	v12 = int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 3664))))
	if v12 != 0 && v12 != -559023410 {
		if *memmap.PtrInt32(6112660, uint32(v12*140)+0x86DA0) != -1 {
			nox_xxx_scriptCallByEventBlock_502490(memmap.PtrInt32(6112660, uint32(v12*140)+0x86D9C), a1, 0, 27)
		}
		*(*uint32)(unsafe.Pointer(uintptr(v1 + 3664))) = 0
		*(*uint8)(unsafe.Pointer(uintptr(v1 + 3668))) = 1
	}
}
func sub_421F10(a1 *int32, a2 int32) *uint8 {
	var (
		v2 *uint8
		v3 int32
		v5 int32
		i  *uint8
		v7 int32
	)
	if a2 != 0 {
		if uint32(a2) != 0xDEADFACE {
			v2 = (*uint8)(memmap.PtrOff(6112660, uint32(a2*140)+0x86D24))
			if *memmap.PtrUint32(6112660, uint32(a2*140)+0x86D78) != 0 {
				if *((*int32)(unsafe.Add(unsafe.Pointer((*int32)(unsafe.Pointer(v2))), unsafe.Sizeof(int32(0))*29))) != -1 || *((*int32)(unsafe.Add(unsafe.Pointer((*int32)(unsafe.Pointer(v2))), unsafe.Sizeof(int32(0))*31))) != -1 {
					v3 = sub_4281F0((*int2)(unsafe.Pointer(a1)), (*int4)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v2), 88)))))
					if v3 != 0 {
						if nox_xxx_polygon_421660(a1, int32(uintptr(memmap.PtrOff(6112660, uint32(a2*140)+0x86D24)))) != 0 {
							return (*uint8)(memmap.PtrOff(6112660, uint32(a2*140)+0x86D24))
						}
					}
				}
			}
		}
	}
	v5 = 1
	if nox_xxx_polygonNextIdx_587000_60352 <= 1 {
		return nil
	}
	for i = (*uint8)(memmap.PtrOff(6112660, 552448)); ; i = (*uint8)(unsafe.Add(unsafe.Pointer(i), 140)) {
		if *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(i))), unsafe.Sizeof(uint32(0))*1))) != 0 && *(*int32)(unsafe.Pointer(i)) != a2 && (*((*int32)(unsafe.Add(unsafe.Pointer((*int32)(unsafe.Pointer(i))), unsafe.Sizeof(int32(0))*9))) != -1 || *((*int32)(unsafe.Add(unsafe.Pointer((*int32)(unsafe.Pointer(i))), unsafe.Sizeof(int32(0))*11))) != -1) {
			v7 = sub_4281F0((*int2)(unsafe.Pointer(a1)), (*int4)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(i), 8)))))
			if v7 != 0 {
				if nox_xxx_polygon_421660(a1, int32(uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(i), -80)))))) != 0 {
					break
				}
			}
		}
		if uint32(func() int32 {
			p := &v5
			*p++
			return *p
		}()) >= uint32(*(*int32)(unsafe.Pointer(&nox_xxx_polygonNextIdx_587000_60352))) {
			return nil
		}
	}
	return (*uint8)(unsafe.Add(unsafe.Pointer(i), -80))
}
func nox_xxx_monsterPolygonEnter_421FF0(a1 int32) {
	var (
		v1  *int32
		v2  int32
		v3  float32
		v4  *uint32
		v5  int32
		v6  *uint32
		v7  int32
		v8  int32
		v9  *uint8
		v10 int2
	)
	v1 = *(**int32)(unsafe.Pointer(uintptr(a1 + 748)))
	if a1 != 0 && int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))&2 != 0 && (*v1 == -559023410 || *(*float32)(unsafe.Pointer(uintptr(a1 + 56))) != *(*float32)(unsafe.Pointer(uintptr(a1 + 72))) || *(*float32)(unsafe.Pointer(uintptr(a1 + 60))) != *(*float32)(unsafe.Pointer(uintptr(a1 + 76)))) {
		v2 = int(*(*float32)(unsafe.Pointer(uintptr(a1 + 56))))
		v3 = *(*float32)(unsafe.Pointer(uintptr(a1 + 60)))
		v10.field_0 = v2
		v10.field_4 = int(v3)
		if *v1 == -559023410 {
			v4 = (*uint32)(unsafe.Pointer(nox_xxx_polygonIsPlayerInPolygon_4217B0(&v10, 0)))
		} else {
			v4 = (*uint32)(unsafe.Pointer(sub_421F10(&v10.field_0, *v1)))
		}
		v5 = *v1
		v6 = v4
		if v4 != nil {
			if uint32(v5) != *(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*20)) {
				if v5 != -559023410 {
					if v5 != 0 {
						v7 = v5 * 35
						if *memmap.PtrInt32(6112660, uint32(v7*4)+0x86DA0) != -1 {
							nox_xxx_scriptCallByEventBlock_502490(memmap.PtrInt32(6112660, uint32(v7*4)+0x86D9C), a1, 0, 26)
						}
					}
					if *(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*29)) != math.MaxUint32 {
						nox_xxx_scriptCallByEventBlock_502490((*int32)(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*28)))), a1, 0, 25)
					}
				}
				*v1 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*20)))
			}
		} else if v5 != 0 && v5 != -559023410 {
			v8 = v5 * 35
			v9 = (*uint8)(memmap.PtrOff(6112660, uint32(v8*4)+0x86D24))
			if *memmap.PtrUint32(6112660, uint32(v8*4)+0x86D78) != 0 && *((*int32)(unsafe.Add(unsafe.Pointer((*int32)(unsafe.Pointer(v9))), unsafe.Sizeof(int32(0))*31))) != -1 {
				nox_xxx_scriptCallByEventBlock_502490((*int32)(unsafe.Add(unsafe.Pointer((*int32)(unsafe.Pointer(v9))), unsafe.Sizeof(int32(0))*30)), a1, 0, 24)
			}
			*v1 = 0
		}
	}
}
func sub_422140(a1 int32) int32 {
	var result int32
	result = a1
	if a1 != 0 {
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 3660))) = 0xDEADFACE
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 3664))) = 0xDEADFACE
	}
	return result
}
func nox_xxx_tileListAddNewSubtile_422160(a1 int32, a2 int32, a3 int32, a4 int32) *int32 {
	var (
		result *int32
		v5     *byte
		i      int32
	)
	result = *(**int32)(unsafe.Pointer(&dword_5d4594_588084))
	if dword_5d4594_588084 == 0 {
		v5 = (*byte)(alloc.Calloc(1, 200))
		dword_5d4594_588084 = uint32(uintptr(unsafe.Pointer(v5)))
		for i = 0; i < 180; i += 20 {
			*(*uint32)(unsafe.Add(unsafe.Pointer(v5), i+16)) = uint32(uintptr(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v5), i+20)))))
			v5 = *(**byte)(unsafe.Pointer(&dword_5d4594_588084))
		}
		*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_588084 + 196))) = 0
		result = *(**int32)(unsafe.Pointer(&dword_5d4594_588084))
	}
	dword_5d4594_588084 = uint32(*(*int32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(int32(0))*4)))
	*result = a1
	*(*int32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(int32(0))*1)) = a2
	*(*int32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(int32(0))*2)) = a3
	*(*int32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(int32(0))*3)) = a4
	*(*int32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(int32(0))*4)) = 0
	return result
}
func nox_xxx_tileFreeTileOne_4221E0(a1 int32) int32 {
	var result int32
	result = a1
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 16))) = dword_5d4594_588084
	dword_5d4594_588084 = uint32(a1)
	return result
}
func nox_xxx_tileFreeTile_422200(a1 int32) int32 {
	var (
		result int32
		v2     int32
	)
	result = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 16))))
	if result != 0 {
		for {
			v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(result + 16))))
			nox_xxx_tileFreeTileOne_4221E0(result)
			result = v2
			if v2 == 0 {
				break
			}
		}
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 16))) = 0
	} else {
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 16))) = 0
	}
	return result
}
func nox_server_mapRWFloorMap_422230(a1 int32) int32 {
	var (
		v2  int32
		v3  int8
		v4  int32
		l   int32
		v6  *uint32
		v7  int8
		v8  int32
		m   int32
		v10 int32
		v11 *uint32
		v12 int8
		v13 int32
		v14 **uint8
		v15 *uint8
		v16 int32
		v17 **uint8
		v18 int32
		v19 int8
		n   int32
		v21 *uint8
		v22 int32
		v23 int32
		v24 int32
		v25 int32
		ii  int32
		v27 int32
		v28 uint16
		v29 int32
		v30 int32
		v31 *uint8
		v32 *uint8
		v33 int32
		v34 int32
		v35 int8
		v36 int32
		v37 *uint8
		v38 *uint8
		v39 *uint8
		v40 uint16
		v41 int32
		v42 int32
		v43 uint16
		k   uint16
		v47 int16
		v48 int32
		v49 int32
		v50 int32
		v51 int32
		v52 int32
		v53 int32
		v54 *uint8
		v55 *uint8
	)
	_ = v55
	var v56 int32
	var v58 int32
	var v59 int32
	var v60 int32
	var v61 int32
	var v62 int32
	var v63 int32 = 0
	var v64 int32 = 0
	var v65 int32 = 0
	var v66 int32
	var v67 int32
	var v68 int32
	var v69 int2
	var v70 int4
	v62 = 4
	cryptFileReadWrite((*uint8)(unsafe.Pointer(&v62))[:2])
	if int32(int16(v62)) > 4 {
		return 0
	}
	if int32(int16(v62)) <= 3 {
		return nox_xxx_tile_422C10(v62, a1)
	}
	if nox_xxx_cryptGetXxx() != 0 {
		cryptFileReadWrite((*uint8)(unsafe.Pointer(&v60))[:4])
		cryptFileReadWrite((*uint8)(unsafe.Pointer(&v61))[:4])
		cryptFileReadWrite((*uint8)(unsafe.Pointer(&v66))[:4])
		cryptFileReadWrite((*uint8)(unsafe.Pointer(&v67))[:4])
		if a1 != 0 {
			sub_428170(unsafe.Pointer(uintptr(a1)), &v70)
			v41 = v70.field_0/23 - v60
			v42 = v70.field_4/23 - v61
			a1 = v70.field_4/23 - v61
			v61 = v70.field_4 / 23
			v60 = v70.field_0 / 23
			cryptFileReadWrite((*uint8)(unsafe.Pointer(&v59))[:2])
			v43 = uint16(int16(v59))
			if int32(uint16(int16(v59))) != math.MaxUint16 {
				for {
					v50 = v41 + (int32(v43) >> 8)
					v51 = v42 + int32(uint8(v43))
					if (int32(uint8(int8(v42)))+int32(uint8(v43)))&1 != 0 {
						v56 = (v50*23 - 23) / 46
						v58 = (v51*23 + 23) / 46
						if noxflags.HasGame(noxflags.GameFlag23) {
							v54 = (*uint8)(unsafe.Pointer(uintptr(*nox_xxx_tileAllocTileInCoordList_5040A0(v56, v58, COERCE_FLOAT(1)))))
						} else {
							(*(*obj_5D4594_2650668_t)(unsafe.Add(unsafe.Pointer(*(**obj_5D4594_2650668_t)(unsafe.Add(unsafe.Pointer(ptr_5D4594_2650668), unsafe.Sizeof((*obj_5D4594_2650668_t)(nil))*uintptr(v56)))), unsafe.Sizeof(obj_5D4594_2650668_t{})*uintptr(v58)))).field_0 |= 1
							v54 = (*uint8)(unsafe.Pointer(&(*(*obj_5D4594_2650668_t)(unsafe.Add(unsafe.Pointer(*(**obj_5D4594_2650668_t)(unsafe.Add(unsafe.Pointer(ptr_5D4594_2650668), unsafe.Sizeof((*obj_5D4594_2650668_t)(nil))*uintptr(v56)))), unsafe.Sizeof(obj_5D4594_2650668_t{})*uintptr(v58)))).field_1))
						}
					} else {
						v52 = v50 * 23 / 46
						v53 = v51 * 23 / 46
						if noxflags.HasGame(noxflags.GameFlag23) {
							v54 = (*uint8)(unsafe.Pointer(uintptr(*nox_xxx_tileAllocTileInCoordList_5040A0(v52, v51*23/46, COERCE_FLOAT(2)))))
						} else {
							v55 = &(*(*obj_5D4594_2650668_t)(unsafe.Add(unsafe.Pointer(*(**obj_5D4594_2650668_t)(unsafe.Add(unsafe.Pointer(ptr_5D4594_2650668), unsafe.Sizeof((*obj_5D4594_2650668_t)(nil))*uintptr(v52)))), unsafe.Sizeof(obj_5D4594_2650668_t{})*uintptr(v53)))).field_0
							*v55 |= 2
							v54 = (*uint8)(unsafe.Pointer(&(*(*obj_5D4594_2650668_t)(unsafe.Add(unsafe.Pointer(*(**obj_5D4594_2650668_t)(unsafe.Add(unsafe.Pointer(ptr_5D4594_2650668), unsafe.Sizeof((*obj_5D4594_2650668_t)(nil))*uintptr(v52)))), unsafe.Sizeof(obj_5D4594_2650668_t{})*uintptr(v53)))).field_6))
						}
					}
					nox_xxx_tileReadOne_422A40(v62, v54)
					cryptFileReadWrite((*uint8)(unsafe.Pointer(&v59))[:2])
					v43 = uint16(int16(v59))
					if int32(uint16(int16(v59))) == math.MaxUint16 {
						break
					}
					v42 = a1
				}
				return 1
			}
		} else {
			for i := int32(0); i < ptr_5D4594_2650668_cap; i++ {
				for j := int32(0); j < ptr_5D4594_2650668_cap; j++ {
					(*(*obj_5D4594_2650668_t)(unsafe.Add(unsafe.Pointer(*(**obj_5D4594_2650668_t)(unsafe.Add(unsafe.Pointer(ptr_5D4594_2650668), unsafe.Sizeof((*obj_5D4594_2650668_t)(nil))*uintptr(j)))), unsafe.Sizeof(obj_5D4594_2650668_t{})*uintptr(i)))).field_0 = 0
				}
			}
			cryptFileReadWrite((*uint8)(unsafe.Pointer(&v59))[:2])
			for k = uint16(int16(v59)); int32(uint16(int16(v59))) != math.MaxUint16; k = uint16(int16(v59)) {
				v47 = 0
				v48 = (int32(k) >> 8) & math.MaxInt8
				v49 = int32(k) & math.MaxInt8
				if (int32(k) & 0x8000) != 0 {
					v47 = math.MinInt16
				}
				if (int32(k) & 128) != 0 {
					*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v47))), 0)) = uint8(int8(int32(v47) | 128))
				}
				if int32(v47) < 0 {
					(*(*obj_5D4594_2650668_t)(unsafe.Add(unsafe.Pointer(*(**obj_5D4594_2650668_t)(unsafe.Add(unsafe.Pointer(ptr_5D4594_2650668), unsafe.Sizeof((*obj_5D4594_2650668_t)(nil))*uintptr(v48)))), unsafe.Sizeof(obj_5D4594_2650668_t{})*uintptr(v49)))).field_0 |= 1
					nox_xxx_tileReadOne_422A40(v62, (*uint8)(unsafe.Pointer(&(*(*obj_5D4594_2650668_t)(unsafe.Add(unsafe.Pointer(*(**obj_5D4594_2650668_t)(unsafe.Add(unsafe.Pointer(ptr_5D4594_2650668), unsafe.Sizeof((*obj_5D4594_2650668_t)(nil))*uintptr(v48)))), unsafe.Sizeof(obj_5D4594_2650668_t{})*uintptr(v49)))).field_1)))
				}
				if int32(v47)&128 != 0 {
					(*(*obj_5D4594_2650668_t)(unsafe.Add(unsafe.Pointer(*(**obj_5D4594_2650668_t)(unsafe.Add(unsafe.Pointer(ptr_5D4594_2650668), unsafe.Sizeof((*obj_5D4594_2650668_t)(nil))*uintptr(v48)))), unsafe.Sizeof(obj_5D4594_2650668_t{})*uintptr(v49)))).field_0 |= 2
					nox_xxx_tileReadOne_422A40(v62, (*uint8)(unsafe.Pointer(&(*(*obj_5D4594_2650668_t)(unsafe.Add(unsafe.Pointer(*(**obj_5D4594_2650668_t)(unsafe.Add(unsafe.Pointer(ptr_5D4594_2650668), unsafe.Sizeof((*obj_5D4594_2650668_t)(nil))*uintptr(v48)))), unsafe.Sizeof(obj_5D4594_2650668_t{})*uintptr(v49)))).field_6)))
				}
				cryptFileReadWrite((*uint8)(unsafe.Pointer(&v59))[:2])
			}
		}
		return 1
	}
	if a1 != 0 {
		sub_428170(unsafe.Pointer(uintptr(a1)), &v70)
		v60 = v70.field_0 / 23
		v18 = v70.field_8 / 23
		v64 = v70.field_8 / 23
		v63 = v70.field_C / 23
		v61 = v70.field_4 / 23
		v66 = v70.field_8/23 - v70.field_0/23 + 1
		v67 = v70.field_C/23 - v70.field_4/23 + 1
		goto LABEL_43
	}
	v2 = v65
	v3 = 0
	v4 = 0
	for l = 0; l < ptr_5D4594_2650668_cap*44; l += 44 {
		if int32(v3) != 0 {
			break
		}
		v2 = 0
		v65 = l
		v6 = (*uint32)(unsafe.Pointer(ptr_5D4594_2650668))
		for int32(*(*uint8)(unsafe.Pointer(uintptr(uint32(l) + *v6)))) == 0 {
			v2++
			v6 = (*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*1))
			if v2 >= ptr_5D4594_2650668_cap {
				goto LABEL_14
			}
		}
		v61 = v4
		v3 = 1
	LABEL_14:
		v4++
	}
	if v2 == ptr_5D4594_2650668_cap && v4 == ptr_5D4594_2650668_cap {
		return 1
	}
	v7 = 0
	v8 = math.MaxInt8
	for m = 5588; m >= 0; m -= 44 {
		if int32(v7) != 0 {
			break
		}
		v10 = 0
		v65 = m
		v11 = (*uint32)(unsafe.Pointer(ptr_5D4594_2650668))
		for int32(*(*uint8)(unsafe.Pointer(uintptr(uint32(m) + *v11)))) == 0 {
			v10++
			v11 = (*uint32)(unsafe.Add(unsafe.Pointer(v11), unsafe.Sizeof(uint32(0))*1))
			if v10 >= ptr_5D4594_2650668_cap {
				goto LABEL_24
			}
		}
		v63 = v8
		v7 = 1
	LABEL_24:
		v8--
	}
	v12 = 0
	v13 = 0
	v14 = (**uint8)(unsafe.Pointer(ptr_5D4594_2650668))
	for {
		if int32(v12) != 0 {
			break
		}
		v15 = *v14
		v16 = 0
		for int32(*v15) == 0 {
			v16++
			v15 = (*uint8)(unsafe.Add(unsafe.Pointer(v15), 44))
			if v16 >= ptr_5D4594_2650668_cap {
				goto LABEL_32
			}
		}
		v60 = v13
		v12 = 1
	LABEL_32:
		v13++
		v14 = (**uint8)(unsafe.Add(unsafe.Pointer(v14), unsafe.Sizeof((*uint8)(nil))*1))
		if v13 >= ptr_5D4594_2650668_cap {
			break
		}
	}
	v17 = (**uint8)(unsafe.Pointer((**obj_5D4594_2650668_t)(unsafe.Add(unsafe.Pointer(ptr_5D4594_2650668), unsafe.Sizeof((*obj_5D4594_2650668_t)(nil))*uintptr(ptr_5D4594_2650668_cap-1)))))
	v18 = v64
	v19 = 0
	for n = ptr_5D4594_2650668_cap - 1; n >= 0; n-- {
		if int32(v19) != 0 {
			break
		}
		v21 = *v17
		v22 = 0
		for int32(*v21) == 0 {
			v22++
			v21 = (*uint8)(unsafe.Add(unsafe.Pointer(v21), 44))
			if v22 >= ptr_5D4594_2650668_cap {
				goto LABEL_40
			}
		}
		v18 = n
		v19 = 1
	LABEL_40:
		v17 = (**uint8)(unsafe.Add(unsafe.Pointer(v17), -int(unsafe.Sizeof((*uint8)(nil))*1)))
	}
	v64 = v18
	v66 = v18 - v60 + 1
	v67 = v63 - v61 + 1
LABEL_43:
	cryptFileReadWrite((*uint8)(unsafe.Pointer(&v60))[:4])
	cryptFileReadWrite((*uint8)(unsafe.Pointer(&v61))[:4])
	cryptFileReadWrite((*uint8)(unsafe.Pointer(&v66))[:4])
	cryptFileReadWrite((*uint8)(unsafe.Pointer(&v67))[:4])
	if a1 != 0 {
		v33 = v61
		if v61 > v63 {
		LABEL_75:
			v59 = math.MaxUint16
			cryptFileReadWrite((*uint8)(unsafe.Pointer(&v59))[:2])
			return 1
		}
		v34 = v60
		for {
			v35 = int8(v34)
			v65 = v34
			if v34 <= v64 {
				break
			}
		LABEL_74:
			if func() int32 {
				p := &v33
				*p++
				return *p
			}() > v63 {
				goto LABEL_75
			}
		}
		v36 = v34*23 + 11
		for {
			if ((int32(uint8(int8(v33))) + int32(v35)) & 1) == 0 {
				if v33&1 != 0 {
					v39 = (*uint8)(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer(*(**obj_5D4594_2650668_t)(unsafe.Add(unsafe.Pointer(ptr_5D4594_2650668), unsafe.Sizeof((*obj_5D4594_2650668_t)(nil))*uintptr((v36-34)/46)))))) + uint32(((v33+1)*23/46)*44))))
					if int32(*v39)&1 != 0 {
						v38 = (*uint8)(unsafe.Add(unsafe.Pointer(v39), 4))
						goto LABEL_70
					}
				} else {
					v37 = (*uint8)(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer(*(**obj_5D4594_2650668_t)(unsafe.Add(unsafe.Pointer(ptr_5D4594_2650668), unsafe.Sizeof((*obj_5D4594_2650668_t)(nil))*uintptr((v36-11)/46)))))) + uint32((v33*23/46)*44))))
					if int32(*v37)&2 != 0 {
						v38 = (*uint8)(unsafe.Add(unsafe.Pointer(v37), 24))
					LABEL_70:
						v69.field_0 = v36
						v69.field_4 = v33*23 + 34
						if nox_xxx_wallMath_427F30(&v69, (*int32)(unsafe.Pointer(uintptr(a1)))) != 0 {
							*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v40))), 0)) = 0
							*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v40))), unsafe.Sizeof(uint16(0))-1)) = uint8(int8(v65))
							v68 = v33 | int32(v40)
							cryptFileReadWrite((*uint8)(unsafe.Pointer(&v68))[:2])
							nox_xxx_tileReadOne_422A40(v62, v38)
						}
						v34 = v60
						goto LABEL_73
					}
				}
			}
		LABEL_73:
			v35 = int8(v65 + 1)
			v36 += 23
			if func() int32 {
				p := &v65
				*p++
				return *p
			}() > v64 {
				goto LABEL_74
			}
		}
	}
	v23 = v61
	if v61 <= v63 {
		v24 = int32(uintptr(unsafe.Pointer(ptr_5D4594_2650668)))
		v25 = v61 * 44
		for {
			for ii = v60; ii <= v18; ii++ {
				v27 = int32(*(*uint32)(unsafe.Pointer(uintptr(v24 + ii*4))))
				if int32(*(*uint8)(unsafe.Pointer(uintptr(v25 + v27))))&3 != 0 {
					*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v28))), 0)) = 0
					*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v28))), unsafe.Sizeof(uint16(0))-1)) = uint8(int8(ii))
					a1 = v23 | int32(v28)
					if int32(*(*uint8)(unsafe.Pointer(uintptr(v25 + v27))))&1 != 0 {
						v29 = a1
						*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v29))), 1)) |= 128
						a1 = v29
					}
					if int32(*(*uint8)(unsafe.Pointer(uintptr(uint32(v25) + *(*uint32)(unsafe.Pointer(uintptr(v24 + ii*4)))))))&2 != 0 {
						v30 = a1
						*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v30))), 0)) = uint8(int8(a1 | 128))
						a1 = v30
					}
					cryptFileReadWrite((*uint8)(unsafe.Pointer(&a1))[:2])
					v24 = int32(uintptr(unsafe.Pointer(ptr_5D4594_2650668)))
				}
				v31 = (*uint8)(unsafe.Pointer(uintptr(uint32(v25) + *(*uint32)(unsafe.Pointer(uintptr(v24 + ii*4))))))
				if int32(*v31)&1 != 0 {
					nox_xxx_tileReadOne_422A40(v62, (*uint8)(unsafe.Add(unsafe.Pointer(v31), 4)))
					v24 = int32(uintptr(unsafe.Pointer(ptr_5D4594_2650668)))
				}
				v32 = (*uint8)(unsafe.Pointer(uintptr(uint32(v25) + *(*uint32)(unsafe.Pointer(uintptr(v24 + ii*4))))))
				if int32(*v32)&2 != 0 {
					nox_xxx_tileReadOne_422A40(v62, (*uint8)(unsafe.Add(unsafe.Pointer(v32), 24)))
					v24 = int32(uintptr(unsafe.Pointer(ptr_5D4594_2650668)))
				}
			}
			v23++
			v25 += 44
			if v23 > v63 {
				break
			}
		}
	}
	v59 = math.MaxUint16
	cryptFileReadWrite((*uint8)(unsafe.Pointer(&v59))[:2])
	return 1
}
func nox_xxx_tileReadOne_422A40(a1 int32, a2 *uint8) uint8 {
	var (
		v2     *uint8
		v3     uint32
		v4     int8
		v5     int8
		v6     int32
		result uint8
		i      int32
		v9     uint32
		v10    int32
		v11    bool
		v12    *int32
		v13    *int32
		v14    uint32
		v15    int32
	)
	v2 = a2
	*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a2))), 0)) = *a2
	v3 = cryptFileReadWrite((*uint8)(unsafe.Pointer(&a2))[:1])
	*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v3))), unsafe.Sizeof(uint16(0))*0)) = *((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v2))), unsafe.Sizeof(uint16(0))*2)))
	*(*uint32)(unsafe.Pointer(v2)) = uint32(uint8(uintptr(unsafe.Pointer(a2))))
	v14 = v3
	cryptFileReadWrite((*uint8)(unsafe.Pointer(&v14))[:2])
	v4 = int8(*(*uint8)(unsafe.Add(unsafe.Pointer(v2), 8)))
	*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), unsafe.Sizeof(uint32(0))*1))) = uint32(int16(uint16(v14)))
	*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a2))), 0)) = uint8(v4)
	cryptFileReadWrite((*uint8)(unsafe.Pointer(&a2))[:1])
	v5 = int8(*(*uint8)(unsafe.Add(unsafe.Pointer(v2), 12)))
	*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), unsafe.Sizeof(uint32(0))*2))) = uint32(uint8(uintptr(unsafe.Pointer(a2))))
	*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a2))), 0)) = uint8(v5)
	cryptFileReadWrite((*uint8)(unsafe.Pointer(&a2))[:1])
	v6 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), unsafe.Sizeof(uint32(0))*4))))
	*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), unsafe.Sizeof(uint32(0))*3))) = uint32(uint8(uintptr(unsafe.Pointer(a2))))
	for *(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v15))), 0)) = 0; v6 != 0; v6 = int32(*(*uint32)(unsafe.Pointer(uintptr(v6 + 16)))) {
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v15))), 0)) = uint8(int8(v15 + 1))
	}
	cryptFileReadWrite((*uint8)(unsafe.Pointer(&v15))[:1])
	result = uint8(int8(nox_xxx_cryptGetXxx()))
	if nox_xxx_cryptGetXxx() != 0 {
		result = uint8(int8(v15))
		v10 = 0
		v11 = int32(uint8(int8(v15))) == 0
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), unsafe.Sizeof(uint32(0))*4))) = 0
		v12 = (*int32)(unsafe.Pointer(v2))
		if !v11 {
			for {
				v13 = nox_xxx_tileListAddNewSubtile_422160(0, 0, 0, 0)
				cryptFileReadWrite((*uint8)(unsafe.Pointer(&a2))[:1])
				*v13 = int32(uint8(uintptr(unsafe.Pointer(a2))))
				cryptFileReadWrite((*uint8)(unsafe.Pointer(&v14))[:2])
				*(*int32)(unsafe.Add(unsafe.Pointer(v13), unsafe.Sizeof(int32(0))*1)) = int32(int16(uint16(v14)))
				cryptFileReadWrite((*uint8)(unsafe.Pointer(&a2))[:1])
				*(*int32)(unsafe.Add(unsafe.Pointer(v13), unsafe.Sizeof(int32(0))*2)) = int32(uint8(uintptr(unsafe.Pointer(a2))))
				result = uint8(cryptFileReadWrite((*uint8)(unsafe.Pointer(&a2))[:1]))
				v10++
				*(*int32)(unsafe.Add(unsafe.Pointer(v13), unsafe.Sizeof(int32(0))*3)) = int32(uint8(uintptr(unsafe.Pointer(a2))))
				*(*int32)(unsafe.Add(unsafe.Pointer(v12), unsafe.Sizeof(int32(0))*4)) = int32(uintptr(unsafe.Pointer(v13)))
				v12 = v13
				if v10 >= int32(uint8(int8(v15))) {
					break
				}
			}
		}
	} else {
		for i = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), unsafe.Sizeof(uint32(0))*4)))); i != 0; i = int32(*(*uint32)(unsafe.Pointer(uintptr(i + 16)))) {
			*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a2))), 0)) = *(*uint8)(unsafe.Pointer(uintptr(i)))
			v9 = cryptFileReadWrite((*uint8)(unsafe.Pointer(&a2))[:1])
			*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v9))), unsafe.Sizeof(uint16(0))*0)) = *(*uint16)(unsafe.Pointer(uintptr(i + 4)))
			v14 = v9
			cryptFileReadWrite((*uint8)(unsafe.Pointer(&v14))[:2])
			*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a2))), 0)) = *(*uint8)(unsafe.Pointer(uintptr(i + 8)))
			cryptFileReadWrite((*uint8)(unsafe.Pointer(&a2))[:1])
			*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a2))), 0)) = *(*uint8)(unsafe.Pointer(uintptr(i + 12)))
			result = uint8(cryptFileReadWrite((*uint8)(unsafe.Pointer(&a2))[:1]))
		}
	}
	return result
}
func nox_xxx_tile_422C10(a1 int32, a2 int32) int32 {
	var (
		v3  int32
		v4  int8
		v5  int32
		i   int32
		v7  *uint32
		v8  int8
		v9  int32
		j   int32
		v11 int32
		v12 *uint32
		v13 int8
		v14 int32
		v15 **uint8
		v16 *uint8
		v17 int32
		v18 **uint8
		v19 int32
		v20 int8
		k   int32
		v22 *uint8
		v23 int32
		v24 int32
		v26 int32
		v27 int32
		l   int32
		m   int32
		n   int32
		ii  int32
		v35 int32
		v36 int32
		v37 int32
		v38 int32
		v39 *uint8
		v40 int32
		v41 *uint8
		jj  int32
		kk  int32
		v44 int32
		v45 int32
		v46 int32
		v47 int32
		v48 int32
		v49 int32
		v51 int32
		v52 int32
		v55 int32
		v56 int32
		v57 int32
		v58 int32
		v59 int32
		v60 *uint8
		v61 int32
		v62 int32
		v63 *int32
		v64 int32
		v66 int32
		v67 int32
		v68 *int32
		v69 *int32
		v70 bool
		v71 *uint8
		v72 int32
		v73 int32
		v74 *int32
		v75 int32
		v76 int32
		v77 int32
		v78 int32
		v79 int32
		v80 int32 = 0
		v81 int32 = 0
		v82 int2
		v83 int4
	)
	if int32(int16(a1)) < 3 {
		return 0
	}
	if nox_xxx_cryptGetXxx() == 0 {
		if a2 != 0 {
			sub_428170(unsafe.Pointer(uintptr(a2)), &v83)
			v76 = v83.field_0 / 23
			v80 = v83.field_8 / 23
			v77 = v83.field_4 / 23
			v81 = v83.field_C / 23
			v24 = v81
			v78 = v83.field_8/23 - v83.field_0/23 + 1
			v79 = v81 - v83.field_4/23 + 1
		} else {
			v3 = 0
			v4 = 0
			v5 = 0
			for i = 0; i < ptr_5D4594_2650668_cap*44; i += 44 {
				if int32(v4) != 0 {
					break
				}
				v3 = 0
				a1 = i
				v7 = (*uint32)(unsafe.Pointer(ptr_5D4594_2650668))
				for int32(*(*uint8)(unsafe.Pointer(uintptr(uint32(i) + *v7)))) == 0 {
					v3++
					v7 = (*uint32)(unsafe.Add(unsafe.Pointer(v7), unsafe.Sizeof(uint32(0))*1))
					if v3 >= ptr_5D4594_2650668_cap {
						goto LABEL_12
					}
				}
				v77 = v5
				v4 = 1
			LABEL_12:
				v5++
			}
			if v3 == ptr_5D4594_2650668_cap && v5 == ptr_5D4594_2650668_cap {
				return 1
			}
			v8 = 0
			v9 = math.MaxInt8
			for j = 5588; j >= 0; j -= 44 {
				if int32(v8) != 0 {
					break
				}
				v11 = 0
				a1 = j
				v12 = (*uint32)(unsafe.Pointer(ptr_5D4594_2650668))
				for int32(*(*uint8)(unsafe.Pointer(uintptr(uint32(j) + *v12)))) == 0 {
					v11++
					v12 = (*uint32)(unsafe.Add(unsafe.Pointer(v12), unsafe.Sizeof(uint32(0))*1))
					if v11 >= ptr_5D4594_2650668_cap {
						goto LABEL_22
					}
				}
				v81 = v9
				v8 = 1
			LABEL_22:
				v9--
			}
			v13 = 0
			v14 = 0
			v15 = (**uint8)(unsafe.Pointer(ptr_5D4594_2650668))
			for {
				if int32(v13) != 0 {
					break
				}
				v16 = *v15
				v17 = 0
				for int32(*v16) == 0 {
					v17++
					v16 = (*uint8)(unsafe.Add(unsafe.Pointer(v16), 44))
					if v17 >= ptr_5D4594_2650668_cap {
						goto LABEL_30
					}
				}
				v76 = v14
				v13 = 1
			LABEL_30:
				v14++
				v15 = (**uint8)(unsafe.Add(unsafe.Pointer(v15), unsafe.Sizeof((*uint8)(nil))*1))
				if v14 >= 128 {
					break
				}
			}
			v18 = (**uint8)(unsafe.Pointer((**obj_5D4594_2650668_t)(unsafe.Add(unsafe.Pointer((**obj_5D4594_2650668_t)(unsafe.Add(unsafe.Pointer(ptr_5D4594_2650668), unsafe.Sizeof((*obj_5D4594_2650668_t)(nil))*uintptr(ptr_5D4594_2650668_cap)))), -int(unsafe.Sizeof((*obj_5D4594_2650668_t)(nil))*1)))))
			v19 = v80
			v20 = 0
			for k = ptr_5D4594_2650668_cap - 1; k >= 0; k-- {
				if int32(v20) != 0 {
					break
				}
				v22 = *v18
				v23 = 0
				for int32(*v22) == 0 {
					v23++
					v22 = (*uint8)(unsafe.Add(unsafe.Pointer(v22), 44))
					if v23 >= ptr_5D4594_2650668_cap {
						goto LABEL_38
					}
				}
				v19 = k
				v20 = 1
			LABEL_38:
				v18 = (**uint8)(unsafe.Add(unsafe.Pointer(v18), -int(unsafe.Sizeof((*uint8)(nil))*1)))
			}
			v24 = v81
			v80 = v19
			v78 = v19 - v76 + 1
			v79 = v81 - v77 + 1
		}
		cryptFileReadWrite((*uint8)(unsafe.Pointer(&v76))[:4])
		cryptFileReadWrite((*uint8)(unsafe.Pointer(&v77))[:4])
		cryptFileReadWrite((*uint8)(unsafe.Pointer(&v78))[:4])
		cryptFileReadWrite((*uint8)(unsafe.Pointer(&v79))[:4])
		if a2 == 0 {
			if v77 > v24 {
				return 1
			}
			v26 = v76
			v27 = v80
			var v28 int32 = v77
			a2 = v81 - v77 + 1
			for {
				if v26 <= v27 {
					for {
						{
							var obj *obj_5D4594_2650668_t = (*obj_5D4594_2650668_t)(unsafe.Add(unsafe.Pointer(*(**obj_5D4594_2650668_t)(unsafe.Add(unsafe.Pointer(ptr_5D4594_2650668), unsafe.Sizeof((*obj_5D4594_2650668_t)(nil))*uintptr(v26)))), unsafe.Sizeof(obj_5D4594_2650668_t{})*uintptr(v28)))
							cryptFileReadWrite((&obj.field_0)[:1])
							if int32(obj.field_0)&1 != 0 {
								cryptFileReadWrite((*uint8)(unsafe.Pointer(&obj.field_1))[:16])
								*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v75))), 0)) = 0
								for l = int32(uintptr(obj.field_5)); l != 0; l = int32(*(*uint32)(unsafe.Pointer(uintptr(l + 16)))) {
									*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v75))), 0)) = uint8(int8(v75 + 1))
								}
								cryptFileReadWrite((*uint8)(unsafe.Pointer(&v75))[:1])
								for m = int32(uintptr(obj.field_5)); m != 0; m = int32(*(*uint32)(unsafe.Pointer(uintptr(m + 16)))) {
									cryptFileReadWrite((*uint8)(unsafe.Pointer(uintptr(m)))[:16])
								}
							}
							if int32(obj.field_0)&2 != 0 {
								cryptFileReadWrite((*uint8)(unsafe.Pointer(&obj.field_6))[:16])
								*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v75))), 0)) = 0
								for n = int32(uintptr(obj.field_10)); n != 0; n = int32(*(*uint32)(unsafe.Pointer(uintptr(n + 16)))) {
									*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v75))), 0)) = uint8(int8(v75 + 1))
								}
								cryptFileReadWrite((*uint8)(unsafe.Pointer(&v75))[:1])
								for ii = int32(uintptr(obj.field_10)); ii != 0; ii = int32(*(*uint32)(unsafe.Pointer(uintptr(ii + 16)))) {
									cryptFileReadWrite((*uint8)(unsafe.Pointer(uintptr(ii)))[:16])
								}
							}
							v26++
						}
						if v26 > v27 {
							break
						}
					}
					v26 = v76
				}
				v28++
				a2--
				if a2 == 0 {
					break
				}
			}
			return 1
		}
		v35 = v77
		if v77 > v24 {
			return 1
		}
		v36 = v76
	LABEL_61:
		v37 = v36
		if v36 > v80 {
			goto LABEL_78
		}
		v38 = v36*23 + 11
		for {
			if ((int32(uint8(int8(v35))) + int32(uint8(int8(v37)))) & 1) == 0 {
				if v35&1 != 0 {
					v41 = (*uint8)(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer(*(**obj_5D4594_2650668_t)(unsafe.Add(unsafe.Pointer(ptr_5D4594_2650668), unsafe.Sizeof((*obj_5D4594_2650668_t)(nil))*uintptr((v38-34)/46)))))) + uint32(((v35+1)*23/46)*44))))
					if (int32(*v41) & 1) == 0 {
						goto LABEL_75
					}
					v40 = int32(uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v41), 4)))))
				} else {
					v39 = (*uint8)(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer(*(**obj_5D4594_2650668_t)(unsafe.Add(unsafe.Pointer(ptr_5D4594_2650668), unsafe.Sizeof((*obj_5D4594_2650668_t)(nil))*uintptr((v38-11)/46)))))) + uint32((v35*23/46)*44))))
					if (int32(*v39) & 2) == 0 {
						goto LABEL_75
					}
					v40 = int32(uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v39), 24)))))
				}
				v82.field_0 = v38
				v82.field_4 = v35*23 + 34
				if nox_xxx_wallMath_427F30(&v82, (*int32)(unsafe.Pointer(uintptr(a2)))) != 0 {
					*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a1))), 0)) = 1
					cryptFileReadWrite((*uint8)(unsafe.Pointer(&a1))[:1])
					cryptFileReadWrite((*uint8)(unsafe.Pointer(uintptr(v40)))[:16])
					*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v75))), 0)) = 0
					for jj = int32(*(*uint32)(unsafe.Pointer(uintptr(v40 + 16)))); jj != 0; jj = int32(*(*uint32)(unsafe.Pointer(uintptr(jj + 16)))) {
						*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v75))), 0)) = uint8(int8(v75 + 1))
					}
					cryptFileReadWrite((*uint8)(unsafe.Pointer(&v75))[:1])
					for kk = int32(*(*uint32)(unsafe.Pointer(uintptr(v40 + 16)))); kk != 0; kk = int32(*(*uint32)(unsafe.Pointer(uintptr(kk + 16)))) {
						cryptFileReadWrite((*uint8)(unsafe.Pointer(uintptr(kk)))[:16])
					}
					goto LABEL_76
				}
			}
		LABEL_75:
			*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a1))), 0)) = 0
			cryptFileReadWrite((*uint8)(unsafe.Pointer(&a1))[:1])
		LABEL_76:
			v37++
			v38 += 23
			if v37 > v80 {
				v36 = v76
			LABEL_78:
				if func() int32 {
					p := &v35
					*p++
					return *p
				}() > v81 {
					return 1
				}
				goto LABEL_61
			}
		}
	}
	cryptFileReadWrite((*uint8)(unsafe.Pointer(&v76))[:4])
	cryptFileReadWrite((*uint8)(unsafe.Pointer(&v77))[:4])
	cryptFileReadWrite((*uint8)(unsafe.Pointer(&v78))[:4])
	cryptFileReadWrite((*uint8)(unsafe.Pointer(&v79))[:4])
	if a2 == 0 {
		for ll := int32(0); ll < ptr_5D4594_2650668_cap; ll++ {
			for mm := int32(0); mm < ptr_5D4594_2650668_cap; mm++ {
				(*obj_5D4594_2650668_t)(unsafe.Add(unsafe.Pointer(*(**obj_5D4594_2650668_t)(unsafe.Add(unsafe.Pointer(ptr_5D4594_2650668), unsafe.Sizeof((*obj_5D4594_2650668_t)(nil))*uintptr(mm)))), unsafe.Sizeof(obj_5D4594_2650668_t{})*uintptr(ll))).field_0 = 0
			}
		}
		v55 = v77
		v81 = v77
		if v77 < v77+v79 {
			v56 = v78
			a2 = v77 * 44
			v57 = v76
			for {
				v58 = v57
				if v57 < v56+v57 {
					for {
						v59 = a2
						cryptFileReadWrite((*uint8)(unsafe.Pointer(uintptr(uint32(a2) + uint32(uintptr(unsafe.Pointer(*(**obj_5D4594_2650668_t)(unsafe.Add(unsafe.Pointer(ptr_5D4594_2650668), unsafe.Sizeof((*obj_5D4594_2650668_t)(nil))*uintptr(v58)))))))))[:1])
						v60 = (*uint8)(unsafe.Pointer(uintptr(uint32(v59) + uint32(uintptr(unsafe.Pointer(*(**obj_5D4594_2650668_t)(unsafe.Add(unsafe.Pointer(ptr_5D4594_2650668), unsafe.Sizeof((*obj_5D4594_2650668_t)(nil))*uintptr(v58)))))))))
						if int32(*v60)&1 != 0 {
							cryptFileReadWrite((*uint8)(unsafe.Add(unsafe.Pointer(v60), 4))[:16])
							cryptFileReadWrite((*uint8)(unsafe.Pointer(&v75))[:1])
							if int32(uint8(int8(v75))) != 0 {
								v61 = 0
								v62 = int32(uint32(uintptr(unsafe.Pointer(*(**obj_5D4594_2650668_t)(unsafe.Add(unsafe.Pointer(ptr_5D4594_2650668), unsafe.Sizeof((*obj_5D4594_2650668_t)(nil))*uintptr(v58)))))) + uint32(v59) + 4)
								if int32(uint8(int8(v75))) != 0 {
									for {
										v63 = nox_xxx_tileListAddNewSubtile_422160(0, 0, 0, 0)
										cryptFileReadWrite((*uint8)(unsafe.Pointer(v63))[:16])
										*(*uint32)(unsafe.Pointer(uintptr(v62 + 16))) = uint32(uintptr(unsafe.Pointer(v63)))
										v61++
										v62 = int32(uintptr(unsafe.Pointer(v63)))
										if v61 >= int32(uint8(int8(v75))) {
											break
										}
									}
								}
							} else {
								*(*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer(*(**obj_5D4594_2650668_t)(unsafe.Add(unsafe.Pointer(ptr_5D4594_2650668), unsafe.Sizeof((*obj_5D4594_2650668_t)(nil))*uintptr(v58)))))) + uint32(a2) + 20))) = 0
							}
						}
						v71 = (*uint8)(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer(*(**obj_5D4594_2650668_t)(unsafe.Add(unsafe.Pointer(ptr_5D4594_2650668), unsafe.Sizeof((*obj_5D4594_2650668_t)(nil))*uintptr(v58)))))) + uint32(a2))))
						if int32(*v71)&2 != 0 {
							cryptFileReadWrite((*uint8)(unsafe.Add(unsafe.Pointer(v71), 24))[:16])
							cryptFileReadWrite((*uint8)(unsafe.Pointer(&v75))[:1])
							if int32(uint8(int8(v75))) != 0 {
								v72 = 0
								v73 = int32(uint32(uintptr(unsafe.Pointer(*(**obj_5D4594_2650668_t)(unsafe.Add(unsafe.Pointer(ptr_5D4594_2650668), unsafe.Sizeof((*obj_5D4594_2650668_t)(nil))*uintptr(v58)))))) + uint32(a2) + 24)
								if int32(uint8(int8(v75))) != 0 {
									for {
										v74 = nox_xxx_tileListAddNewSubtile_422160(0, 0, 0, 0)
										cryptFileReadWrite((*uint8)(unsafe.Pointer(v74))[:16])
										*(*uint32)(unsafe.Pointer(uintptr(v73 + 16))) = uint32(uintptr(unsafe.Pointer(v74)))
										v72++
										v73 = int32(uintptr(unsafe.Pointer(v74)))
										if v72 >= int32(uint8(int8(v75))) {
											break
										}
									}
								}
							} else {
								*(*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer(*(**obj_5D4594_2650668_t)(unsafe.Add(unsafe.Pointer(ptr_5D4594_2650668), unsafe.Sizeof((*obj_5D4594_2650668_t)(nil))*uintptr(v58)))))) + uint32(a2) + 40))) = 0
							}
						}
						v56 = v78
						v57 = v76
						v58++
						if v58 >= v78+v76 {
							break
						}
					}
					v55 = v77
				}
				v70 = func() int32 {
					p := &v81
					*p++
					return *p
				}() < v55+v79
				a2 += 44
				if !v70 {
					break
				}
			}
		}
		return 1
	}
	sub_428170(unsafe.Pointer(uintptr(a2)), &v83)
	v44 = v83.field_0 / 23
	v76 = v83.field_0 / 23
	v45 = v83.field_4 / 23
	v77 = v83.field_4 / 23
	v46 = v83.field_4 / 23
	if v83.field_4/23 >= v83.field_4/23+v79 {
		return 1
	}
	v47 = v78
	for {
		a2 = v44
		if v44 < v47+v44 {
			v81 = v44 * 23
			v48 = v44 * 23
			for {
				cryptFileReadWrite((*uint8)(unsafe.Pointer(&a1))[:1])
				if int32(uint8(int8(a1))) != 0 {
					if v46&1 != 0 {
						v64 = (v48 - 23) / 46
						v66 = (v46 + 23) / 46
						if noxflags.HasGame(noxflags.GameFlag23) {
							v52 = int32(*nox_xxx_tileAllocTileInCoordList_5040A0(v64, v66, COERCE_FLOAT(1)))
						} else {
							*(*uint8)(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer(*(**obj_5D4594_2650668_t)(unsafe.Add(unsafe.Pointer(ptr_5D4594_2650668), unsafe.Sizeof((*obj_5D4594_2650668_t)(nil))*uintptr(v64)))))) + uint32(v66*44)))) |= 1
							v52 = int32(uint32(uintptr(unsafe.Pointer(*(**obj_5D4594_2650668_t)(unsafe.Add(unsafe.Pointer(ptr_5D4594_2650668), unsafe.Sizeof((*obj_5D4594_2650668_t)(nil))*uintptr(v64)))))) + uint32(v66*44) + 4)
						}
					} else {
						v49 = v48 / 46
						v51 = v46 / 46
						if noxflags.HasGame(noxflags.GameFlag23) {
							v52 = int32(*nox_xxx_tileAllocTileInCoordList_5040A0(v49, v51, COERCE_FLOAT(2)))
						} else {
							*(*uint8)(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer(*(**obj_5D4594_2650668_t)(unsafe.Add(unsafe.Pointer(ptr_5D4594_2650668), unsafe.Sizeof((*obj_5D4594_2650668_t)(nil))*uintptr(v49)))))) + uint32(v51*44)))) |= 2
							v52 = int32(uint32(uintptr(unsafe.Pointer(*(**obj_5D4594_2650668_t)(unsafe.Add(unsafe.Pointer(ptr_5D4594_2650668), unsafe.Sizeof((*obj_5D4594_2650668_t)(nil))*uintptr(v49)))))) + uint32(v51*44) + 24)
						}
					}
					cryptFileReadWrite((*uint8)(unsafe.Pointer(uintptr(v52)))[:16])
					cryptFileReadWrite((*uint8)(unsafe.Pointer(&v75))[:1])
					if int32(uint8(int8(v75))) != 0 {
						v67 = 0
						v68 = (*int32)(unsafe.Pointer(uintptr(v52)))
						if int32(uint8(int8(v75))) != 0 {
							for {
								v69 = nox_xxx_tileListAddNewSubtile_422160(0, 0, 0, 0)
								cryptFileReadWrite((*uint8)(unsafe.Pointer(v69))[:16])
								*(*int32)(unsafe.Add(unsafe.Pointer(v68), unsafe.Sizeof(int32(0))*4)) = int32(uintptr(unsafe.Pointer(v69)))
								v67++
								v68 = v69
								if v67 >= int32(uint8(int8(v75))) {
									break
								}
							}
						}
					} else {
						*(*uint32)(unsafe.Pointer(uintptr(v52 + 16))) = 0
					}
				}
				v47 = v78
				v44 = v76
				v48 = v81 + 23
				v70 = func() int32 {
					p := &a2
					*p++
					return *p
				}() < v78+v76
				v81 += 23
				if !v70 {
					break
				}
			}
			v45 = v77
		}
		v46++
		if v46 >= v45+v79 {
			break
		}
	}
	return 1
}
func nox_xxx_effectFloatValueLoad_4235C0(a1 *byte, a2 *byte, a3 *obj_412ae0_t) int32 {
	var v4 float32 = float32(nox_xxx_effectLoadFloat_423730_parse_float())
	nox_xxx_effectLoadFloatParam_4235F0(a1, v4, a3)
	return 1
}
func nox_xxx_effectLoadFloatParam_4235F0(a1 *byte, a2 float32, a3 *obj_412ae0_t) int32 {
	if libc.StrCmp(a1, CString("ATTACKEFFECT")) == 0 {
		a3.field_11 = a2
	} else if libc.StrCmp(a1, CString("ATTACKPREHITEFFECT")) == 0 {
		a3.field_14 = a2
	} else if libc.StrCmp(a1, CString("ATTACKPREDAMAGEEFFECT")) == 0 {
		a3.field_17 = a2
	} else if libc.StrCmp(a1, CString("DEFENDEFFECT")) == 0 {
		a3.field_20 = a2
	} else if libc.StrCmp(a1, CString("DEFENDCOLLIDEEFFECT")) == 0 {
		a3.field_23 = a2
	} else if libc.StrCmp(a1, CString("UPDATEEFFECT")) == 0 {
		a3.field_26 = a2
	} else if libc.StrCmp(a1, CString("ENGAGEEFFECT")) == 0 {
		a3.field_30 = a2
	} else if libc.StrCmp(a1, CString("DISENGAGEEFFECT")) == 0 {
		a3.field_32 = a2
	} else {
		return 0
	}
	return 1
}
func nox_xxx_effectLoadFloat_423730_parse_float() float64 {
	var v2 [6]byte
	libc.StrCpy(&v2[0], CString(" =\n\r\t"))
	var v0 *byte = libc.StrTok(nil, &v2[0])
	if v0 == nil {
		return -1.0
	}
	return libc.Atof(libc.GoString(v0))
}
func nox_xxx_effectDwordValueLoad_423780(a1 *byte, a2 *byte, a3 *obj_412ae0_t) int32 {
	var v3 int32 = nox_xxx_effectLoadDword_4238F0_parse_int()
	nox_xxx_effectLoadDwordParam_4237B0(a1, v3, a3)
	return 1
}
func nox_xxx_effectLoadDwordParam_4237B0(a1 *byte, a2 int32, a3 *obj_412ae0_t) int32 {
	if libc.StrCmp(a1, CString("ATTACKEFFECT")) == 0 {
		a3.field_12 = a2
	} else if libc.StrCmp(a1, CString("ATTACKPREHITEFFECT")) == 0 {
		a3.field_15 = a2
	} else if libc.StrCmp(a1, CString("ATTACKPREDAMAGEEFFECT")) == 0 {
		a3.field_18 = a2
	} else if libc.StrCmp(a1, CString("DEFENDEFFECT")) == 0 {
		a3.field_21 = a2
	} else if libc.StrCmp(a1, CString("DEFENDCOLLIDEEFFECT")) == 0 {
		a3.field_24 = a2
	} else if libc.StrCmp(a1, CString("UPDATEEFFECT")) == 0 {
		a3.field_27 = a2
	} else if libc.StrCmp(a1, CString("ENGAGEEFFECT")) == 0 {
		a3.field_31 = a2
	} else if libc.StrCmp(a1, CString("DISENGAGEEFFECT")) == 0 {
		a3.field_33 = a2
	} else {
		return 0
	}
	return 1
}
func nox_xxx_effectLoadDword_4238F0_parse_int() int32 {
	var v2 [6]byte
	libc.StrCpy(&v2[0], CString(" =\n\r\t"))
	var v0 *byte = libc.StrTok(nil, &v2[0])
	if v0 == nil {
		return -1
	}
	return int32(libc.Atoi(libc.GoString(v0)))
}
func set_bitmask_flags_from_plus_separated_names_423930(input *byte, bitmask *uint32, allowed_names **byte) {
	var input_copy [256]byte
	libc.StrCpy(&input_copy[0], input)
	if libc.StrNCmp(&input_copy[0], CString("NULL"), 4) == 0 {
		return
	}
	var cur_value *byte = libc.StrTok(&input_copy[0], CString("+"))
	for cur_value != nil {
		set_one_bitmask_flag_by_name_4239C0(cur_value, bitmask, allowed_names)
		cur_value = libc.StrTok(nil, CString("+"))
	}
}
func set_one_bitmask_flag_by_name_4239C0(name *byte, bitmask *uint32, allowed_names **byte) int32 {
	for i := int8(0); *(**byte)(unsafe.Add(unsafe.Pointer(allowed_names), unsafe.Sizeof((*byte)(nil))*uintptr(i))) != nil; i++ {
		if libc.StrCaseCmp(*(**byte)(unsafe.Add(unsafe.Pointer(allowed_names), unsafe.Sizeof((*byte)(nil))*uintptr(i))), name) == 0 {
			*bitmask |= uint32(1 << int32(i))
			return 1
		}
	}
	return 0
}
func set_bitmask_flags_from_plus_separated_names_multiple_423A10(input *byte, bitmask *uint32) {
	var input_copy [256]byte
	libc.StrCpy(&input_copy[0], input)
	if libc.StrNCmp(&input_copy[0], CString("NULL"), 4) == 0 {
		return
	}
	var cur_value *byte = libc.StrTok(&input_copy[0], CString("+"))
	for cur_value != nil {
		var cur_allowed_values *uint32 = memmap.PtrUint32(0x587000, 61096)
		for *cur_allowed_values != 0 {
			if set_one_bitmask_flag_by_name_4239C0(cur_value, bitmask, (**byte)(unsafe.Pointer(cur_allowed_values))) != 0 {
				break
			}
			cur_allowed_values = (*uint32)(unsafe.Add(unsafe.Pointer(cur_allowed_values), unsafe.Sizeof(uint32(0))*32))
		}
		cur_value = libc.StrTok(nil, CString("+"))
	}
}
func nox_parse_shape(s *nox_shape, buf *byte) int32 {
	if libc.StrNCmp(buf, CString("NULL"), 4) == 0 {
		s.kind = nox_shape_kind(NOX_SHAPE_NONE)
		return 1
	} else if libc.StrNCmp(buf, CString("CENTER"), 6) == 0 {
		s.kind = nox_shape_kind(NOX_SHAPE_CENTER)
		return 1
	} else if libc.StrNCmp(buf, CString("CIRCLE"), 6) == 0 {
		s.kind = nox_shape_kind(NOX_SHAPE_CIRCLE)
		stdio.Sscanf(buf, "%*s %f", &s.circle_r)
		s.circle_r2 = s.circle_r * s.circle_r
		return 1
	} else if libc.StrNCmp(buf, CString("BOX"), 3) == 0 {
		s.kind = nox_shape_kind(NOX_SHAPE_BOX)
		stdio.Sscanf(buf, "%*s %f %f", &s.box_w, &s.box_h)
		nox_shape_box_calc(s)
		return 1
	}
	return 0
}
func sub_423F80(lpFileName *byte, a2 int32, a3 int32, a4 int32) HANDLE {
	var (
		result       HANDLE
		v5           HANDLE
		v6           *byte
		v7           uint8
		v8           HANDLE
		FileName     [260]byte
		FindFileData _WIN32_FIND_DATAA
	)
	result = compatFindFirstFileA(lpFileName, &FindFileData)
	v5 = result
	v8 = result
	if uintptr(result) != uintptr(math.MaxUint32) {
		for {
			if FindFileData.dwFileAttributes&uint32(FILE_ATTRIBUTE_DIRECTORY) != 0 {
				*(*uint32)(unsafe.Pointer(uintptr(a2)))++
				if libc.StrCaseCmp(&FindFileData.cFileName[0], CString(".")) != 0 {
					if libc.StrCaseCmp(&FindFileData.cFileName[0], CString("..")) != 0 {
						libc.StrCpy(&FileName[0], lpFileName)
						*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v8))), libc.StrLen(&FileName[0])))), 3))) = 0
						libc.StrCat(&FileName[0], &FindFileData.cFileName[0])
						v6 = &FileName[libc.StrLen(&FileName[0])+1]
						v7 = *memmap.PtrUint8(0x587000, 64698)
						*(*uint16)(unsafe.Pointer(func() *byte {
							p := &v6
							*p = (*byte)(unsafe.Add(unsafe.Pointer(*p), -1))
							return *p
						}())) = *memmap.PtrUint16(0x587000, 64696)
						*(*byte)(unsafe.Add(unsafe.Pointer(v6), 2)) = byte(v7)
						sub_423F80(&FileName[0], a2, a3, a4)
						v5 = v8
					}
				}
			} else {
				*(*uint32)(unsafe.Pointer(uintptr(a3)))++
				*(*uint32)(unsafe.Pointer(uintptr(a4))) += FindFileData.nFileSizeLow
			}
			if compatFindNextFileA(v5, &FindFileData) == 0 {
				break
			}
		}
		result = unsafe.Pointer(uintptr(compatFindClose(v5)))
	}
	return result
}
func sub_4240F0(a1 int32, a2 *byte, a3 int32) int32 {
	var (
		v3 *byte
		v4 int32
		v5 *uint8
	)
	v3 = *(**byte)(memmap.PtrOff(0x587000, 64704))
	v4 = 0
	if *memmap.PtrUint32(0x587000, 64704) == 0 {
		return 0
	}
	v5 = (*uint8)(memmap.PtrOff(0x587000, 64704))
	for libc.StrCmp(a2, v3) != 0 {
		v3 = (*byte)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v5))), unsafe.Sizeof(uint32(0))*2))))))
		v5 = (*uint8)(unsafe.Add(unsafe.Pointer(v5), 8))
		v4++
		if v3 == nil {
			return 0
		}
	}
	*(*uint32)(unsafe.Pointer(uintptr(*memmap.PtrUint32(0x587000, uint32(v4*8+0xFCC4)) + uint32(a1)))) = uint32(a3)
	return 1
}
func nox_xxx_parseSoundSetBin_424170(a1 *byte) int32 {
	var (
		v2 *File
		v3 *uint32
		v4 *byte
		v5 int32
		v6 [256]byte
		v7 [256]byte
	)
	v2 = nox_binfile_open_408CC0(a1, 0)
	if v2 == nil {
		return 0
	}
	if nox_binfile_cryptSet_408D40(v2, 5) == 0 {
		return 0
	}
	for nox_xxx_parseString_409470(v2, (*uint8)(unsafe.Pointer(&v6[0]))) != 0 {
		v3 = (*uint32)(alloc.Calloc(1, 84))
		*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*19)) = dword_5d4594_588120
		*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*20)) = 0
		dword_5d4594_588120 = uint32(uintptr(unsafe.Pointer(v3)))
		v4 = (*byte)(alloc.Calloc(libc.StrLen(&v6[0])+1, 1))
		*v3 = uint32(uintptr(unsafe.Pointer(v4)))
		libc.StrCpy(v4, &v6[0])
		for nox_xxx_parseString_409470(v2, (*uint8)(unsafe.Pointer(&v6[0]))) != 0 && libc.StrCmp(&v6[0], CString("END")) != 0 && nox_xxx_parseString_409470(v2, (*uint8)(unsafe.Pointer(&v7[0]))) != 0 {
			v5 = nox_xxx_utilFindSound_40AF50(&v7[0])
			if sub_4240F0(int32(uintptr(unsafe.Pointer(v3))), &v6[0], v5) == 0 {
				return 0
			}
		}
	}
	nox_binfile_close_408D90(v2)
	return 1
}
func sub_4242C0() {
	var (
		v0 int32
		v1 int32
	)
	v0 = int32(dword_5d4594_588120)
	if dword_5d4594_588120 != 0 {
		for {
			v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(v0 + 76))))
			*(*unsafe.Pointer)(unsafe.Pointer(uintptr(v0))) = nil
			alloc.Free(unsafe.Pointer(uintptr(v0)))
			v0 = v1
			if v1 == 0 {
				break
			}
		}
		dword_5d4594_588120 = 0
	} else {
		dword_5d4594_588120 = 0
	}
}
func nox_xxx_monsterGetSoundSet_424300(a1 int32) int32 {
	var result int32
	result = a1
	if a1 != 0 {
		if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))&2 != 0 {
			result = int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))) + 488))))
		} else {
			result = 0
		}
	}
	return result
}
func nox_xxx_setNPCVoiceSet_424320(a1 int32, a2 int32) int32 {
	var result int32
	result = a1
	if a1 != 0 {
		if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))&2 != 0 {
			*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))) + 488))) = uint32(a2)
			result = 1
		} else {
			result = 0
		}
	}
	return result
}
func nox_xxx_getDefaultSoundSet_424350(a1 *byte) **byte {
	var v1 int32
	if a1 == nil {
		return nil
	}
	v1 = int32(dword_5d4594_588120)
	if dword_5d4594_588120 == 0 {
		return nil
	}
	for libc.StrCmp(*(**byte)(unsafe.Pointer(uintptr(v1))), a1) != 0 {
		v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 76))))
		if v1 == 0 {
			return nil
		}
	}
	return (**byte)(unsafe.Pointer(uintptr(v1)))
}
func sub_4243C0() unsafe.Pointer {
	return *(*unsafe.Pointer)(unsafe.Pointer(&dword_5d4594_588120))
}
func sub_4243D0(a1 int32) int32 {
	var result int32
	if a1 != 0 {
		result = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 76))))
	} else {
		result = 0
	}
	return result
}
func nox_xxx_updateSpellRelated_424830(a1 int32, a2 int32) int32 {
	var result int32
	if a1 != 0 && a2 >= 0 && a2 <= 8 {
		result = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + a2*4 + 4))))
	} else {
		result = 0
	}
	return result
}
func nox_xxx_enchantByName_424880(a1 *byte) int32 {
	if a1 == nil {
		return -1
	}
	var v1 *byte
	var v2 int32
	var v3 *uint8
	var result int32
	v1 = *(**byte)(memmap.PtrOff(0x587000, 65316))
	v2 = 0
	if *memmap.PtrUint32(0x587000, 65316) != 0 {
		v3 = (*uint8)(memmap.PtrOff(0x587000, 65316))
		for {
			if libc.StrCmp(v1, a1) == 0 {
				break
			}
			v1 = (*byte)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v3))), unsafe.Sizeof(uint32(0))*1))))))
			v3 = (*uint8)(unsafe.Add(unsafe.Pointer(v3), 4))
			v2++
			if v1 == nil {
				break
			}
		}
	}
	if *memmap.PtrUint32(0x587000, uint32(v2*4+0xFF24)) != 0 {
		result = v2
	} else {
		result = -1
	}
	return result
}
func sub_4248F0(a1 int32) *byte {
	return *(**byte)(memmap.PtrOff(0x587000, uint32(a1*4+0xFF24)))
}
func sub_424900(a1 int32) int32 {
	var (
		result int32
		i      *uint8
	)
	result = 0
	for i = (*uint8)(memmap.PtrOff(0x587000, 65188)); *(*uint32)(unsafe.Pointer(i)) != uint32(a1); i = (*uint8)(unsafe.Add(unsafe.Pointer(i), 4)) {
		result++
	}
	return result
}
func nox_xxx_getEnchantSpell_424920(a1 int32) int32 {
	return int32(*memmap.PtrUint32(0x587000, uint32(a1*4+0xFEA4)))
}
func sub_424CB0(a1 int32) int8 {
	var (
		v1 int8
		v2 int32
		v3 *uint8
	)
	v1 = 0
	v2 = 0
	if *(*int32)(unsafe.Pointer(&dword_587000_66116)) <= 0 {
		return 0
	}
	v3 = (*uint8)(memmap.PtrOff(0x587000, 66000))
	for {
		if nox_xxx_testUnitBuffs_4FF350((*nox_object_t)(unsafe.Pointer(uintptr(a1))), int8(uint8(*(*uint32)(unsafe.Pointer(v3))))) != 0 {
			v1++
		}
		v2++
		v3 = (*uint8)(unsafe.Add(unsafe.Pointer(v3), 4))
		if v2 >= *(*int32)(unsafe.Pointer(&dword_587000_66116)) {
			break
		}
	}
	return v1
}
func sub_424D00() int32 {
	var result int32
	if *(*int32)(unsafe.Pointer(&dword_587000_66116)) <= 0 {
		result = -1
	} else {
		result = int32(*memmap.PtrUint32(0x587000, 66000))
	}
	return result
}
func sub_424D20(a1 int32) int32 {
	var (
		v1 int32
		i  *uint8
	)
	v1 = 0
	if *(*int32)(unsafe.Pointer(&dword_587000_66116)) <= 0 {
		return -1
	}
	for i = (*uint8)(memmap.PtrOff(0x587000, 66000)); *(*int32)(unsafe.Pointer(i)) != a1 || v1 >= *(*int32)(unsafe.Pointer(&dword_587000_66116))-1; i = (*uint8)(unsafe.Add(unsafe.Pointer(i), 4)) {
		if func() int32 {
			p := &v1
			*p++
			return *p
		}() >= *(*int32)(unsafe.Pointer(&dword_587000_66116)) {
			return -1
		}
	}
	return int32(*memmap.PtrUint32(0x587000, uint32(v1*4)+66004))
}
func nox_xxx_abilityNameToN_424D80(a1 *byte) int32 {
	var (
		v1 *byte
		v2 int32
		v3 *uint8
	)
	v1 = *(**byte)(memmap.PtrOff(0x587000, 69736))
	v2 = 0
	if *memmap.PtrUint32(0x587000, 69736) == 0 {
		return 0
	}
	v3 = (*uint8)(memmap.PtrOff(0x587000, 69736))
	for libc.StrCmp(v1, a1) != 0 {
		v1 = (*byte)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v3))), unsafe.Sizeof(uint32(0))*1))))))
		v3 = (*uint8)(unsafe.Add(unsafe.Pointer(v3), 4))
		v2++
		if v1 == nil {
			return 0
		}
	}
	return v2
}
func nox_xxx_loadWariorParams_424DF0() int32 {
	var (
		result int32
		v1     float32
		v2     float32
		v3     float32
		v4     float32
		v5     float32
		v6     float32
		v7     float32
		v8     float32
		v9     float32
		v10    float32
	)
	v1 = float32(nox_xxx_gamedataGetFloat_419D40(CString("BerserkerChargeDelay")))
	*memmap.PtrUint32(6112660, 599212) = uint32(int(v1))
	v2 = float32(nox_xxx_gamedataGetFloat_419D40(CString("BerserkerChargeDuration")))
	*memmap.PtrUint32(6112660, 599216) = uint32(int(v2))
	v3 = float32(nox_xxx_gamedataGetFloat_419D40(CString("WarcryDelay")))
	*memmap.PtrUint32(6112660, 599264) = uint32(int(v3))
	v4 = float32(nox_xxx_gamedataGetFloat_419D40(CString("WarCryDuration")))
	*memmap.PtrUint32(6112660, 599268) = uint32(int(v4))
	v5 = float32(nox_xxx_gamedataGetFloat_419D40(CString("HarpoonDelay")))
	*memmap.PtrUint32(6112660, 599316) = uint32(int(v5))
	v6 = float32(nox_xxx_gamedataGetFloat_419D40(CString("HarpoonDuration")))
	*memmap.PtrUint32(6112660, 599320) = uint32(int(v6))
	v7 = float32(nox_xxx_gamedataGetFloat_419D40(CString("TreadLightlyDelay")))
	*memmap.PtrUint32(6112660, 599368) = uint32(int(v7))
	v8 = float32(nox_xxx_gamedataGetFloat_419D40(CString("TreadLightlyDuration")))
	*memmap.PtrUint32(6112660, 599372) = uint32(int(v8))
	v9 = float32(nox_xxx_gamedataGetFloat_419D40(CString("EyeOfTheWolfDelay")))
	*memmap.PtrUint32(6112660, 599420) = uint32(int(v9))
	v10 = float32(nox_xxx_gamedataGetFloat_419D40(CString("EyeOfTheWolfDuration")))
	result = int(v10)
	*memmap.PtrUint32(6112660, 599424) = uint32(result)
	return result
}
func sub_425230(a1 int32, a2 int32) int32 {
	return int32(*memmap.PtrUint32(6112660, uint32((a1+a2+a1*12)*4)+0x92484))
}
func nox_xxx_abilityGetName_425250(a1 int32) *byte {
	return *(**byte)(memmap.PtrOff(0x587000, uint32(a1*4)+69736))
}
func nox_xxx_abilityGetName_0_425260(a1 int32) int32 {
	var result int32
	if a1 > 0 && a1 < 6 && *memmap.PtrUint32(6112660, uint32(a1*52)+0x92474) != 0 {
		result = int32(*memmap.PtrUint32(6112660, uint32(a1*52)+0x9245C))
	} else {
		result = 0
	}
	return result
}
func nox_xxx_abil_425290(a1 *libc.WChar) int32 {
	var (
		v1 int32
		v2 *uint8
	)
	v1 = 1
	v2 = (*uint8)(memmap.PtrOff(6112660, 599184))
	for *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), unsafe.Sizeof(uint32(0))*6))) == 0 || nox_wcscmp(*(**libc.WChar)(unsafe.Pointer(v2)), a1) != 0 {
		v2 = (*uint8)(unsafe.Add(unsafe.Pointer(v2), 52))
		v1++
		if int32(uintptr(unsafe.Pointer(v2))) >= int32(uintptr(memmap.PtrOff(6112660, 599444))) {
			return 0
		}
	}
	return v1
}
func nox_xxx_abilityCooldown_4252D0(a1 int32) int32 {
	return int32(*memmap.PtrUint32(6112660, uint32(a1*52)+599160))
}
func sub_4252F0(a1 int32) int32 {
	return int32(*memmap.PtrUint32(6112660, uint32(a1*52)+0x92460))
}
func nox_xxx_spellGetAbilityIcon_425310(a1 int32, a2 int32) unsafe.Pointer {
	return unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, uint32((a1+a2+a1*12)*4)+599140)))
}
func nox_xxx_bookFirstKnownAbil_425330() int32 {
	var (
		result int32
		v1     *uint8
	)
	result = 1
	v1 = (*uint8)(memmap.PtrOff(6112660, 599208))
	for *(*uint32)(unsafe.Pointer(v1)) == 0 {
		v1 = (*uint8)(unsafe.Add(unsafe.Pointer(v1), 52))
		result++
		if int32(uintptr(unsafe.Pointer(v1))) >= int32(uintptr(memmap.PtrOff(6112660, 599468))) {
			return 0
		}
	}
	return result
}
func nox_xxx_bookNextKnownAbil_425350(a1 int32) int32 {
	var (
		result int32
		v2     *uint8
	)
	result = a1 + 1
	if a1+1 >= 6 {
		return 0
	}
	v2 = (*uint8)(memmap.PtrOff(6112660, uint32(result*52)+0x92474))
	for *(*uint32)(unsafe.Pointer(v2)) == 0 {
		v2 = (*uint8)(unsafe.Add(unsafe.Pointer(v2), 52))
		result++
		if int32(uintptr(unsafe.Pointer(v2))) >= int32(uintptr(memmap.PtrOff(6112660, 599468))) {
			return 0
		}
	}
	return result
}
func sub_425380(a1 int32) int32 {
	var (
		result int32
		i      *uint8
	)
	if a1 == 0 {
		return 0
	}
	result = a1 - 1
	if a1-1 <= 0 {
		return 0
	}
	for i = (*uint8)(memmap.PtrOff(6112660, uint32(result*52)+0x92474)); *(*uint32)(unsafe.Pointer(i)) == 0; i = (*uint8)(unsafe.Add(unsafe.Pointer(i), -52)) {
		if func() int32 {
			p := &result
			*p--
			return *p
		}() <= 0 {
			return 0
		}
	}
	return result
}
func sub_4253B0(a1 int32) int32 {
	return int32(*memmap.PtrUint32(6112660, uint32(a1*52)+0x92474))
}
func sub_4253D0(a1 int32) int32 {
	return int32(*memmap.PtrUint32(6112660, uint32(a1*52)+0x92470))
}
func sub_4253F0(a1 int32) int32 {
	var result int32
	result = a1
	*memmap.PtrUint32(6112660, uint32(a1*52)+0x92470) = 1
	return result
}
func sub_425410(a1 int32) int32 {
	var result int32
	result = a1
	*memmap.PtrUint32(6112660, uint32(a1*52)+0x92470) = 0
	return result
}
func sub_425430(a1 int32, a2 int32) int32 {
	return bool2int((uint32(a2) & *memmap.PtrUint32(6112660, uint32(a1*52)+0x92480)) != 0)
}
func sub_425450(a1 int32) int32 {
	return int32(*memmap.PtrUint32(6112660, uint32(a1*52)+0x92480))
}
func sub_425470(a1 int32) int32 {
	return int32(*memmap.PtrUint32(6112660, uint32(a1*52)+0x9247C))
}
func sub_4254A0(a1 int32, a2 *uint8) int32 {
	*(*uint32)(unsafe.Pointer(uintptr(a1))) = uint32(uintptr(unsafe.Pointer(a2)))
	*(*uint8)(unsafe.Pointer(uintptr(a1 + 4))) = 0
	return int32(*a2) & 1
}
func sub_4254C0(a1 **uint8) bool {
	var (
		v1 int8
		v2 *uint8
	)
	v1 = int8(int32(*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(a1))), 4)))) + 1)
	*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(a1))), 4))) = uint8(v1)
	if int32(v1) == 8 {
		v2 = *a1
		*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(a1))), 4))) = 0
		*a1 = (*uint8)(unsafe.Add(unsafe.Pointer(v2), 1))
	}
	return ((1 << int32(*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(a1))), 4))))) & int32(**a1)) > 0
}
func sub_425500(a1 int32, a2 *uint8, a3 int8) *uint8 {
	var result *uint8
	result = a2
	*(*uint32)(unsafe.Pointer(uintptr(a1))) = uint32(uintptr(unsafe.Pointer(a2)))
	*(*uint8)(unsafe.Pointer(uintptr(a1 + 4))) = 0
	*a2 = uint8(a3)
	return result
}
func sub_425520(a1 int32, a2 int8) int8 {
	var (
		v2     int8
		v3     *uint8
		result int8
	)
	v2 = int8(int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 4)))) + 1)
	*(*uint8)(unsafe.Pointer(uintptr(a1 + 4))) = uint8(v2)
	if int32(v2) == 8 {
		v3 = *(**uint8)(unsafe.Pointer(uintptr(a1)))
		*(*uint8)(unsafe.Pointer(uintptr(a1 + 4))) = 0
		*(*uint32)(unsafe.Pointer(uintptr(a1))) = uint32(uintptr(unsafe.Pointer(func() *uint8 {
			p := &v3
			*p = (*uint8)(unsafe.Add(unsafe.Pointer(*p), 1))
			return *p
		}())))
		*v3 = 0
	}
	result = int8(int32(a2) << int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 4)))))
	**(**uint8)(unsafe.Pointer(uintptr(a1))) |= uint8(result)
	return result
}
func sub_425550(a1 *uint8, a2 *uint8, a3 int32) int32 {
	var (
		v3  int32
		v4  int32
		v5  int8
		v6  int32
		v8  [8]byte
		v9  [8]byte
		v10 bool
	)
	v3 = 1
	v4 = 0
	v5 = int8(sub_4254A0(int32(uintptr(unsafe.Pointer(&v9[0]))), a1))
	sub_425500(int32(uintptr(unsafe.Pointer(&v8[0]))), a2, v5)
	if a3 == 1 {
		return 1
	}
	v6 = a3 - 1
	for {
		if (func() int32 {
			p := &v4
			*p++
			return *p
		}() % 7) == 0 {
			sub_425520(int32(uintptr(unsafe.Pointer(&v8[0]))), 1)
			v3++
		}
		v10 = sub_4254C0((**uint8)(unsafe.Pointer(&v9[0])))
		sub_425520(int32(uintptr(unsafe.Pointer(&v8[0]))), int8(bool2int(v10)))
		v6--
		if v6 == 0 {
			break
		}
	}
	return v3
}
func sub_4255F0(a1 *uint8, a2 *uint8, a3 int32) uint32 {
	var (
		v3  uint32
		v4  int8
		v5  int32
		v6  int32
		v8  [8]byte
		v9  [8]byte
		v10 bool
	)
	v3 = uint32(a3+7) >> 3
	v4 = int8(sub_4254A0(int32(uintptr(unsafe.Pointer(&v8[0]))), a1))
	sub_425500(int32(uintptr(unsafe.Pointer(&v9[0]))), a2, v4)
	v5 = 0
	v6 = a3 - 1
	if a3 != 1 {
		for {
			if (func() int32 {
				p := &v5
				*p++
				return *p
			}() % 7) == 0 {
				sub_4254C0((**uint8)(unsafe.Pointer(&v8[0])))
			}
			v10 = sub_4254C0((**uint8)(unsafe.Pointer(&v8[0])))
			sub_425520(int32(uintptr(unsafe.Pointer(&v9[0]))), int8(bool2int(v10)))
			v6--
			if v6 == 0 {
				break
			}
		}
	}
	return v3
}
func sub_425680(a1 *uint8, a2 *uint8) int32 {
	var (
		v2     int32
		result int32
		v4     [8]byte
		v5     [8]byte
		v6     int8
		v7     bool
	)
	v6 = int8(sub_4254A0(int32(uintptr(unsafe.Pointer(&v4[0]))), a1))
	sub_425500(int32(uintptr(unsafe.Pointer(&v5[0]))), a2, v6)
	v2 = 0
	for {
		if (func() int32 {
			p := &v2
			*p++
			return *p
		}() % 7) == 0 {
			sub_4254C0((**uint8)(unsafe.Pointer(&v4[0])))
		}
		v7 = sub_4254C0((**uint8)(unsafe.Pointer(&v4[0])))
		sub_425520(int32(uintptr(unsafe.Pointer(&v5[0]))), int8(bool2int(v7)))
		result = sub_425710(int32(uintptr(unsafe.Pointer(a2))), int32((*(*uint32)(unsafe.Pointer(&v5[0]))-uint32(int32(uintptr(unsafe.Pointer(a2)))))>>2))
		if result != 0 {
			break
		}
	}
	return result
}
func sub_425710(a1 int32, a2 int32) int32 {
	var (
		v2 int32
		v3 int32
	)
	v2 = 0
	v3 = 0
	if a2 <= 0 {
		return 0
	}
	for {
		if (v3 % 2) == 0 {
			v2 = 0
		}
		if int32(*(*uint8)(unsafe.Pointer(uintptr(v3 + a1)))) == 0 {
			break
		}
		v2 = 0
	LABEL_8:
		if func() int32 {
			p := &v3
			*p++
			return *p
		}() >= a2 {
			return 0
		}
	}
	if func() int32 {
		p := &v2
		*p++
		return *p
	}() != 2 {
		goto LABEL_8
	}
	return 1
}
func nox_common_list_clear_425760(list *nox_list_item_t) {
	list.field_0 = list
	list.field_1 = list
	list.field_2 = list
}
func sub_425770(a1p unsafe.Pointer) unsafe.Pointer {
	var (
		a1     *uint32 = (*uint32)(a1p)
		result *uint32
	)
	result = a1
	*a1 = uint32(uintptr(unsafe.Pointer(a1)))
	*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*1)) = uint32(uintptr(unsafe.Pointer(a1)))
	*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*2)) = 0
	return unsafe.Pointer(result)
}
func sub_425790(a1 *int32, a2 *uint32) int32 {
	var (
		v2     int32
		v3     int32
		v4     *int32
		result int32
	)
	v2 = 0
	v3 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*2)))
	v4 = (*int32)(unsafe.Pointer(nox_common_list_getNext_425940((*nox_list_item_t)(unsafe.Pointer(a1)))))
	if v4 != nil {
		for v3 > *(*int32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(int32(0))*2)) {
			v2++
			v4 = (*int32)(unsafe.Pointer(nox_common_list_getNext_425940((*nox_list_item_t)(unsafe.Pointer(v4)))))
			if v4 == nil {
				goto LABEL_4
			}
		}
		nox_common_list_append_4258E0((*nox_list_item_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v4)))))), (*nox_list_item_t)(unsafe.Pointer(a2)))
		result = v2
	} else {
	LABEL_4:
		nox_common_list_append_4258E0((*nox_list_item_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1)))))), (*nox_list_item_t)(unsafe.Pointer(a2)))
		result = v2
	}
	return result
}
func sub_4257F0(a1 *int32, a2 *uint32) {
	var (
		v2 int32
		v3 *int32
	)
	v2 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*2)))
	v3 = (*int32)(unsafe.Pointer(nox_common_list_getNext_425940((*nox_list_item_t)(unsafe.Pointer(a1)))))
	if v3 == nil {
		nox_common_list_append_4258E0((*nox_list_item_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1)))))), (*nox_list_item_t)(unsafe.Pointer(a2)))
		return
	}
	for *(*int32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(int32(0))*2)) > v2 {
		v3 = (*int32)(unsafe.Pointer(nox_common_list_getNext_425940((*nox_list_item_t)(unsafe.Pointer(v3)))))
		if v3 == nil {
			nox_common_list_append_4258E0((*nox_list_item_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1)))))), (*nox_list_item_t)(unsafe.Pointer(a2)))
			return
		}
	}
	nox_common_list_append_4258E0((*nox_list_item_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v3)))))), (*nox_list_item_t)(unsafe.Pointer(a2)))
}
func sub_425840(a1 *uint32, a2 int32) {
	var (
		result *uint32
		v3     *int32
		v4     *uint32
	)
	result = (*uint32)(unsafe.Pointer(uintptr(*a1)))
	v3 = (*int32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*1)))))
	if (*uint32)(unsafe.Pointer(uintptr(*a1))) != a1 {
		v4 = *(**uint32)(unsafe.Pointer(uintptr(a2 + 4)))
		*v4 = uint32(uintptr(unsafe.Pointer(result)))
		*(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*1)) = uint32(uintptr(unsafe.Pointer(v4)))
		*v3 = a2
		*(*uint32)(unsafe.Pointer(uintptr(a2 + 4))) = uint32(uintptr(unsafe.Pointer(v3)))
		nox_common_list_clear_425760((*nox_list_item_t)(unsafe.Pointer(a1)))
	}
}
func sub_425870(a1 **uint32) int32 {
	var (
		result int32
		v2     **uint32
	)
	result = 0
	v2 = (**uint32)(unsafe.Pointer(*a1))
	if unsafe.Pointer(*a1) != unsafe.Pointer(a1) {
		for {
			v2 = (**uint32)(unsafe.Pointer(*v2))
			result++
			if v2 == a1 {
				break
			}
		}
	}
	return result
}
func nox_common_list_getFirstSafe_425890(list *nox_list_item_t) *nox_list_item_t {
	return nox_common_list_getNextSafe_4258A0(list)
}
func nox_common_list_getNextSafe_4258A0(list *nox_list_item_t) *nox_list_item_t {
	if list == nil {
		return nil
	}
	return nox_common_list_getNext_425940(list)
}
func sub_4258C0(a1 **uint32, a2 int32) *uint32 {
	var (
		result *uint32
		v3     int32
		v4     int32
	)
	result = *a1
	if unsafe.Pointer(*a1) == unsafe.Pointer(a1) {
		return nil
	}
	v3 = a2
	for {
		v4 = func() int32 {
			p := &v3
			x := *p
			*p--
			return x
		}()
		if v4 == 0 {
			break
		}
		result = (*uint32)(unsafe.Pointer(uintptr(*result)))
		if unsafe.Pointer(result) == unsafe.Pointer(a1) {
			return nil
		}
	}
	return result
}
func nox_common_list_append_4258E0(list *nox_list_item_t, cur *nox_list_item_t) {
	if list == nil || cur == nil {
		panic("abort")
	}
	var it *nox_list_item_t = list.field_1
	cur.field_0 = list
	cur.field_1 = it
	list.field_1 = cur
	if it != nil {
		it.field_0 = cur
	}
}
func sub_425900(a1 *uint32, a2 *uint32) *uint32 {
	var result *uint32
	result = a2
	*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*1)) = uint32(uintptr(unsafe.Pointer(a1)))
	*a2 = *a1
	*a1 = uint32(uintptr(unsafe.Pointer(a2)))
	*(*uint32)(unsafe.Pointer(uintptr(*a2 + 4))) = uint32(uintptr(unsafe.Pointer(a2)))
	return result
}
func sub_425920(a1 **uint32) **uint32 {
	var result **uint32
	result = a1
	**(**uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof((*uint32)(nil))*1)) = uint32(uintptr(unsafe.Pointer(*a1)))
	*(*uint32)(unsafe.Add(unsafe.Pointer(*a1), unsafe.Sizeof(uint32(0))*1)) = uint32(uintptr(unsafe.Pointer(*(**uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof((*uint32)(nil))*1)))))
	*a1 = (*uint32)(unsafe.Pointer(a1))
	*(**uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof((*uint32)(nil))*1)) = (*uint32)(unsafe.Pointer(a1))
	return result
}
func nox_common_list_getNext_425940(list *nox_list_item_t) *nox_list_item_t {
	var it *nox_list_item_t = list.field_0
	if it != nil && it == it.field_2 {
		return nil
	}
	return it
}
func sub_425960(a1 int32) int32 {
	if *(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 4))) + 8))) != *(*uint32)(unsafe.Pointer(uintptr(a1 + 4))) {
		return int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 4))))
	}
	return 0
}
func sub_425980(a1 *uint32) *uint32 {
	var result *uint32
	result = (*uint32)(unsafe.Pointer(uintptr(*a1)))
	if result == (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*2))))) {
		result = (*uint32)(unsafe.Pointer(uintptr(*result)))
		if result == (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*2))))) {
			result = nil
		}
	}
	return result
}
func sub_4259A0(a1 int32) int32 {
	var result int32
	result = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 4))))
	if uint32(result) == *(*uint32)(unsafe.Pointer(uintptr(result + 8))) {
		result = int32(*(*uint32)(unsafe.Pointer(uintptr(result + 4))))
		if uint32(result) == *(*uint32)(unsafe.Pointer(uintptr(result + 8))) {
			result = 0
		}
	}
	return result
}
func sub_4259C0() {
	if *memmap.PtrUint32(6112660, 599472) == 0 {
		nox_common_list_clear_425760((*nox_list_item_t)(memmap.PtrOff(6112660, 599460)))
		*memmap.PtrUint32(6112660, 599472) = 1
	}
}
func sub_4259F0() *int32 {
	var (
		result *int32
		v1     *int32
		v2     *int32
		v3     *int32
		v4     *int32
	)
	result = sub_425A50()
	v1 = result
	if result != nil {
		for {
			v2 = sub_425A60(v1)
			v3 = sub_425BC0(int32(uintptr(unsafe.Pointer(v1))))
			if v3 != nil {
				for {
					v4 = sub_425BE0(v3)
					sub_425920((**uint32)(unsafe.Pointer(v3)))
					alloc.Free(unsafe.Pointer(v3))
					v3 = v4
					if v4 == nil {
						break
					}
				}
			}
			sub_425920((**uint32)(unsafe.Pointer(v1)))
			alloc.Free(unsafe.Pointer(v1))
			v1 = v2
			if v2 == nil {
				break
			}
		}
	}
	return result
}
func sub_425A50() *int32 {
	return (*int32)(unsafe.Pointer(nox_common_list_getFirstSafe_425890((*nox_list_item_t)(unsafe.Pointer(memmap.PtrInt32(6112660, 599460))))))
}
func sub_425A60(a1 *int32) *int32 {
	return (*int32)(unsafe.Pointer(nox_common_list_getNextSafe_4258A0((*nox_list_item_t)(unsafe.Pointer(a1)))))
}
func sub_425A70(a1 int32) *int32 {
	var result *int32
	result = sub_425A50()
	if result == nil {
		return nil
	}
	for *(*int32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(int32(0))*8)) != a1 {
		result = sub_425A60(result)
		if result == nil {
			return nil
		}
	}
	return result
}
func sub_425AA0(a1 int32) int32 {
	var v1 *int32
	v1 = sub_425A50()
	if v1 == nil {
		return 0
	}
	for *(*int32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(int32(0))*8)) != a1 {
		v1 = sub_425A60(v1)
		if v1 == nil {
			return 0
		}
	}
	return 1
}
func sub_425AD0(a1 int32, a2 *libc.WChar) *libc.WChar {
	var v2 *libc.WChar
	v2 = nil
	if sub_425AA0(a1) == 0 {
		v2 = (*libc.WChar)(alloc.Calloc(1, 52))
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), unsafe.Sizeof(uint32(0))*8))) = uint32(a1)
		nox_wcscpy((*libc.WChar)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(libc.WChar(0))*6)), a2)
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), unsafe.Sizeof(uint32(0))*9))) = 0
		sub_425770(unsafe.Pointer(v2))
		nox_common_list_clear_425760((*nox_list_item_t)(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), unsafe.Sizeof(uint32(0))*10)))))
		nox_common_list_append_4258E0((*nox_list_item_t)(unsafe.Pointer(uintptr(int32(uintptr(memmap.PtrOff(6112660, 599460)))))), (*nox_list_item_t)(unsafe.Pointer(v2)))
	}
	return v2
}
func sub_425B30(a1 unsafe.Pointer, a2 int32) {
	var (
		v2 *uint32
		v3 *uint32
	)
	v2 = (*uint32)(alloc.Calloc(1, 16))
	v3 = v2
	*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*3)) = uint32(a2)
	sub_425770(unsafe.Pointer(v2))
	nox_common_list_append_4258E0((*nox_list_item_t)(unsafe.Pointer(uintptr(int32(uintptr(a1))+40))), (*nox_list_item_t)(unsafe.Pointer(v3)))
}
func sub_425B60(lpMem unsafe.Pointer, a2 int32) *byte {
	var (
		v2     *int32
		result *byte
	)
	v2 = sub_425BC0(int32(uintptr(lpMem)))
	if v2 != nil {
		for *(*int32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(int32(0))*3)) != a2 {
			v2 = sub_425BE0(v2)
			if v2 == nil {
				goto LABEL_6
			}
		}
		sub_425920((**uint32)(unsafe.Pointer(v2)))
		alloc.Free(unsafe.Pointer(v2))
	}
LABEL_6:
	result = (*byte)(unsafe.Add(unsafe.Pointer((*byte)(lpMem)), 40))
	if *((*unsafe.Pointer)(unsafe.Add(unsafe.Pointer((*unsafe.Pointer)(lpMem)), unsafe.Sizeof(unsafe.Pointer(nil))*11))) == unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(lpMem)), 40))) {
		sub_425920((**uint32)(lpMem))
		lpMem = nil
	}
	return result
}
func sub_425BC0(a1 int32) *int32 {
	return (*int32)(unsafe.Pointer(nox_common_list_getFirstSafe_425890((*nox_list_item_t)(unsafe.Pointer(uintptr(a1 + 40))))))
}
func sub_425BE0(a1 *int32) *int32 {
	return (*int32)(unsafe.Pointer(nox_common_list_getNextSafe_4258A0((*nox_list_item_t)(unsafe.Pointer(a1)))))
}
func nox_xxx_countObserverPlayers_425BF0() int32 {
	var (
		v0 int32
		i  *byte
		v2 int32
	)
	v0 = 0
	if noxflags.HasGame(noxflags.GameFlag16) {
		for i = (*byte)(unsafe.Pointer(nox_common_playerInfoGetFirst_416EA0())); i != nil; i = (*byte)(unsafe.Pointer(nox_common_playerInfoGetNext_416EE0((*nox_playerInfo)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(i))))))))) {
			v2 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(i))), unsafe.Sizeof(uint32(0))*920))))
			if v2&1 != 0 && (v2&32) == 0 && *(*byte)(unsafe.Add(unsafe.Pointer(i), 2064)) != 31 {
				v0++
			}
		}
	}
	return v0
}
func nox_xxx_firstReplaceablePlayer_425C40() *byte {
	var result *byte
	result = (*byte)(unsafe.Pointer(nox_common_playerInfoGetFirst_416EA0()))
	if result == nil {
		return nil
	}
	for (*(*byte)(unsafe.Add(unsafe.Pointer(result), 3680))&1) == 0 || *(*byte)(unsafe.Add(unsafe.Pointer(result), 2064)) == math.MaxUint8 {
		result = (*byte)(unsafe.Pointer(nox_common_playerInfoGetNext_416EE0((*nox_playerInfo)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(result)))))))))
		if result == nil {
			return nil
		}
	}
	return result
}
func nox_xxx_nextReplaceablePlayer_425C70(a1 int32) *byte {
	var result *byte
	result = (*byte)(unsafe.Pointer(uintptr(a1)))
	if a1 == 0 {
		return nil
	}
	for (*(*byte)(unsafe.Add(unsafe.Pointer(result), 3680))&1) == 0 || *(*byte)(unsafe.Add(unsafe.Pointer(result), 2064)) == math.MaxUint8 {
		result = (*byte)(unsafe.Pointer(nox_common_playerInfoGetNext_416EE0((*nox_playerInfo)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(result)))))))))
		if result == nil {
			return nil
		}
	}
	return result
}
func sub_425CA0(a1 int32, a2 int32) *byte {
	var (
		result *byte
		v3     int32
		v4     int32
		v5     int32
		v6     int32
		v7     int32
		v8     int32
		v9     int32
		v10    int32
		v11    int32
		v12    int32
		v13    int32
	)
	result = (*byte)(unsafe.Pointer(uintptr(bool2int(noxflags.HasGame(noxflags.GameOnline)))))
	if result != nil {
		result = (*byte)(unsafe.Pointer(uintptr(bool2int(noxflags.HasGame(noxflags.GameModeQuest)))))
		if result == nil {
			if a1 != 0 {
				v3 = a2
				if *(*int32)(unsafe.Pointer(uintptr(a1 + 4648))) == -1 {
					v4 = int32(func() uint32 {
						p := &dword_5d4594_608316
						x := *p
						*p++
						return x
					}())
					v13 = v4
					v5 = v4 * 32
					libc.StrCpy((*byte)(memmap.PtrOff(6112660, uint32(v5)+0x9283C)), (*byte)(unsafe.Pointer(uintptr(a1+2096))))
					if int32(*(*uint8)(unsafe.Pointer(uintptr(v3 + 2064)))) == 31 {
						v6 = nox_xxx_net_getIP_554200(0)
					} else {
						v6 = nox_xxx_net_getIP_554200(uint32(int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 2064)))) + 1))
					}
					*memmap.PtrUint32(6112660, uint32(v5)+0x92848) = cnet.Htonl(uint32(v6))
					*memmap.PtrUint32(6112660, uint32(v5)+600140) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 2068)))
					*memmap.PtrUint8(6112660, uint32(v5)+0x92850) = *(*uint8)(unsafe.Pointer(uintptr(a1 + 2251)))
					*(*uint32)(unsafe.Pointer(uintptr(a1 + 4648))) = uint32(v13)
				} else {
					v13 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 4648))))
				}
				if v3 != 0 {
					if *(*int32)(unsafe.Pointer(uintptr(v3 + 4648))) == -1 {
						v12 = int32(func() uint32 {
							p := &dword_5d4594_608316
							x := *p
							*p++
							return x
						}())
						libc.StrCpy((*byte)(memmap.PtrOff(6112660, uint32(v13*32)+0x9283C)), (*byte)(unsafe.Pointer(uintptr(v3+2096))))
						if int32(*(*uint8)(unsafe.Pointer(uintptr(v3 + 2064)))) == 31 {
							v7 = nox_xxx_net_getIP_554200(0)
							v8 = v12
							*memmap.PtrUint32(6112660, uint32(v13*32)+0x92848) = cnet.Htonl(uint32(v7))
						} else {
							v9 = nox_xxx_net_getIP_554200(uint32(int32(*(*uint8)(unsafe.Pointer(uintptr(v3 + 2064)))) + 1))
							v8 = v12
							*memmap.PtrUint32(6112660, uint32(v12*32)+0x92848) = cnet.Htonl(uint32(v9))
						}
						v10 = v8 * 32
						*memmap.PtrUint32(6112660, uint32(v10)+600140) = *(*uint32)(unsafe.Pointer(uintptr(v3 + 2068)))
						*memmap.PtrUint8(6112660, uint32(v10)+0x92850) = *(*uint8)(unsafe.Pointer(uintptr(v3 + 2251)))
						*(*uint32)(unsafe.Pointer(uintptr(v3 + 4648))) = uint32(v8)
					} else {
						v12 = int32(*(*uint32)(unsafe.Pointer(uintptr(v3 + 4648))))
					}
				} else {
					*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v12))), 0)) = math.MaxUint8
				}
				v11 = int32(dword_5d4594_739392)
				*memmap.PtrUint8(6112660, dword_5d4594_739392*2+0x94840) = uint8(int8(v13))
				*memmap.PtrUint8(6112660, uint32(v11*2)+0x94841) = uint8(int8(v12))
				dword_5d4594_739392 = uint32(v11 + 1)
				result = *(**byte)(unsafe.Pointer(&dword_5d4594_608316))
				if dword_5d4594_608316 >= math.MaxUint8 {
					result = nox_xxx_net_4263C0()
				}
			}
		}
	}
	return result
}
func sub_425E90(a1 int32, a2 int8) int32 {
	var result int32
	result = bool2int(noxflags.HasGame(noxflags.GameOnline))
	if result != 0 {
		result = bool2int(noxflags.HasGame(noxflags.GameModeQuest))
		if result == 0 {
			result = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 4648))))
			if result != -1 {
				result *= 32
				*memmap.PtrUint8(6112660, uint32(result)+0x92851) = uint8(a2)
			}
		}
	}
	return result
}
func sub_425ED0(a1 int32, a2 int8) int32 {
	var result int32
	result = bool2int(noxflags.HasGame(noxflags.GameOnline))
	if result != 0 {
		result = bool2int(noxflags.HasGame(noxflags.GameModeQuest))
		if result == 0 {
			result = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 4648))))
			if result != -1 {
				result *= 32
				*memmap.PtrUint8(6112660, uint32(result)+0x92858) = uint8(a2)
			}
		}
	}
	return result
}
func sub_425F10(pl *nox_playerInfo) {
	var (
		a1     int32 = int32(uintptr(unsafe.Pointer(pl)))
		result int32
		v2     int32
		v3     int32
		v4     int32
		v5     *uint8
		v6     int32
		v7     bool
		v8     int32
		v9     int32
		v10    int32
		v11    int32
		v12    int32
	)
	result = bool2int(noxflags.HasGame(noxflags.GameOnline))
	if result != 0 {
		result = bool2int(noxflags.HasGame(noxflags.GameModeQuest))
		if result == 0 {
			v2 = a1
			if a1 != 0 {
				result = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 4648))))
				v3 = -1
				if result == -1 {
					v4 = int32(func() uint32 {
						p := &dword_5d4594_608316
						x := *p
						*p++
						return x
					}())
					v12 = v4
					v5 = (*uint8)(unsafe.Pointer(uintptr(v2 + 2096)))
					v6 = v4 * 32
					for {
						if v3 == 0 {
							break
						}
						v7 = int32(*func() *uint8 {
							p := &v5
							x := *p
							*p = (*uint8)(unsafe.Add(unsafe.Pointer(*p), 1))
							return x
						}()) == 0
						v3--
						if v7 {
							break
						}
					}
					libc.MemCpy(memmap.PtrOff(6112660, uint32(v6)+0x9283C), unsafe.Add(unsafe.Pointer(v5), v3+1), int(^v3))
					if int32(*(*uint8)(unsafe.Pointer(uintptr(v2 + 2064)))) == 31 {
						v8 = nox_xxx_net_getIP_554200(0)
						*memmap.PtrUint32(6112660, uint32(v6)+0x92848) = cnet.Htonl(uint32(v8))
						if nox_common_getEngineFlag(nox_engine_flag(NOX_ENGINE_FLAG_DISABLE_GRAPHICS_RENDERING)) {
							*memmap.PtrUint8(6112660, uint32(v6)+0x92858) = 0
							goto LABEL_18
						}
						v9 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 3680))))
						if v9&1 != 0 && (v9&32) == 0 {
							*memmap.PtrUint8(6112660, uint32(v6)+0x92858) = 0
							goto LABEL_18
						}
					} else {
						v10 = nox_xxx_net_getIP_554200(uint32(int32(*(*uint8)(unsafe.Pointer(uintptr(v2 + 2064)))) + 1))
						*memmap.PtrUint32(6112660, uint32(v6)+0x92848) = cnet.Htonl(uint32(v10))
						v11 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 3680))))
						if v11&1 != 0 && (v11&32) == 0 {
							*memmap.PtrUint8(6112660, uint32(v6)+0x92858) = 0
							goto LABEL_18
						}
					}
					*memmap.PtrUint8(6112660, uint32(v6)+0x92858) = 1
				LABEL_18:
					*memmap.PtrUint32(6112660, uint32(v6)+600140) = *(*uint32)(unsafe.Pointer(uintptr(v2 + 2068)))
					*memmap.PtrUint8(6112660, uint32(v6)+0x92850) = *(*uint8)(unsafe.Pointer(uintptr(v2 + 2251)))
					*memmap.PtrUint32(6112660, uint32(v6)+0x92854) = uint32(libc.GetTime(nil))
					result = v12
					*memmap.PtrUint8(6112660, uint32(v6)+0x92851) = 1
					*(*uint32)(unsafe.Pointer(uintptr(v2 + 4648))) = uint32(v12)
				}
			}
		}
	}
}
func sub_426150() *byte {
	var (
		result *byte
		v1     *byte
		v2     *byte
		v3     int16
		v4     int32
	)
	result = (*byte)(unsafe.Pointer(uintptr(bool2int(noxflags.HasGame(noxflags.GameOnline)))))
	if result != nil {
		result = (*byte)(unsafe.Pointer(uintptr(bool2int(noxflags.HasGame(noxflags.GameModeQuest)))))
		if result == nil {
			v1 = (*byte)(sub_416640())
			v2 = nox_xxx_cliGamedataGet_416590(0)
			*memmap.PtrUint16(6112660, 599482) = uint16(uint8(*(*byte)(unsafe.Add(unsafe.Pointer(v1), 103))))
			*memmap.PtrUint32(6112660, 599484) = uint32(uint8(*(*byte)(unsafe.Add(unsafe.Pointer(v1), 104))))
			*memmap.PtrUint32(6112660, 599488) = *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v1))), unsafe.Sizeof(uint32(0))*10)))
			*memmap.PtrUint32(6112660, 599492) = uint32(sub_4200E0())
			*memmap.PtrUint8(6112660, 599502) = uint8(int8(bool2int((*(*byte)(unsafe.Add(unsafe.Pointer(v2), 53)) & 192) != 0)))
			v3 = int16(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v2))), unsafe.Sizeof(uint16(0))*26))))
			if int32(v3)&256 != 0 {
				dword_5d4594_599496 = 0
			LABEL_14:
				*memmap.PtrUint8(6112660, 599500) = uint8(int8(bool2int((*(*byte)(unsafe.Add(unsafe.Pointer(v2), 53)) & 64) == 0)))
				*memmap.PtrUint32(6112660, 599508) = uint32(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v2))), unsafe.Sizeof(uint16(0))*27))))
				*memmap.PtrUint32(6112660, 599512) = uint32(uint8(*(*byte)(unsafe.Add(unsafe.Pointer(v2), 56))))
				*memmap.PtrUint8(6112660, 599516) = uint8(*(*byte)(unsafe.Add(unsafe.Pointer(v1), 100)))
				*memmap.PtrUint32(6112660, 599520) = uint32(*(*byte)(unsafe.Add(unsafe.Pointer(v1), 101)) & 15)
				*memmap.PtrUint32(6112660, 599524) = uint32(int32(uint8(*(*byte)(unsafe.Add(unsafe.Pointer(v1), 101)))) >> 4)
				*memmap.PtrUint32(6112660, 599528) = uint32(*(*uint16)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v1), 105)))))
				*memmap.PtrUint32(6112660, 599532) = uint32(*(*uint16)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v1), 107)))))
				*memmap.PtrUint8(6112660, 599536) = uint8(*(*byte)(unsafe.Add(unsafe.Pointer(v1), 102)))
				*memmap.PtrUint8(6112660, 599537) = uint8(*(*byte)(unsafe.Add(unsafe.Pointer(v1), 100)) & 48)
				*memmap.PtrUint8(6112660, 599501) = uint8(sub_417DE0())
				libc.StrNCpy((*byte)(memmap.PtrOff(6112660, 599828)), (*byte)(unsafe.Add(unsafe.Pointer(v2), 9)), 15)
				*memmap.PtrUint8(6112660, 599843) = 0
				libc.MemCpy(memmap.PtrOff(6112660, 599540), unsafe.Add(unsafe.Pointer(v2), 24), 100)
				*memmap.PtrUint32(6112660, 599564) = *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), unsafe.Sizeof(uint32(0))*12)))
				*memmap.PtrUint32(6112660, 599560) = *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), unsafe.Sizeof(uint32(0))*11)))
				result = libc.StrNCpy((*byte)(memmap.PtrOff(6112660, 599572)), v2, 8)
				*memmap.PtrUint8(6112660, 599580) = 0
				*memmap.PtrUint32(6112660, 600112) = 0
				*memmap.PtrUint32(6112660, 600084) = 0
				*memmap.PtrUint32(6112660, 600088) = 0
				*memmap.PtrUint32(6112660, 600092) = 0
				*memmap.PtrUint32(6112660, 600096) = 0
				return result
			}
			if int32(v3)&32 != 0 {
				v4 = 1
			} else {
				if int32(v3)&64 != 0 {
					dword_5d4594_599496 = 2
					goto LABEL_14
				}
				if int32(v3)&16 != 0 {
					dword_5d4594_599496 = 3
					goto LABEL_14
				}
				if (int32(v3) & 1024) == 0 {
					goto LABEL_14
				}
				v4 = 4
			}
			dword_5d4594_599496 = uint32(v4)
			goto LABEL_14
		}
	}
	return result
}
func nox_xxx_net_4263C0() *byte {
	var (
		result *byte
		i      *byte
		v2     *byte
		j      int32
	)
	result = (*byte)(unsafe.Pointer(uintptr(bool2int(noxflags.HasGame(noxflags.GameOnline)))))
	if result != nil {
		result = (*byte)(unsafe.Pointer(uintptr(bool2int(noxflags.HasGame(noxflags.GameModeQuest)))))
		if result == nil {
			sub_4282F0(int32(uintptr(memmap.PtrOff(6112660, 599476))), int32(uintptr(memmap.PtrOff(6112660, 600124))), dword_5d4594_608316)
			sub_428540(int32(uintptr(memmap.PtrOff(6112660, 599476))), (*byte)(memmap.PtrOff(6112660, 608320)), *(*int32)(unsafe.Pointer(&dword_5d4594_739392)))
			*memmap.PtrUint32(6112660, 599504) = uint32(libc.GetTime(nil) - libc.Time(dword_5d4594_600116))
			sub_428810(int32(uintptr(memmap.PtrOff(6112660, 599476))), 1)
			libc.MemSet(memmap.PtrOff(6112660, 600124), 0, 8192)
			libc.MemSet(memmap.PtrOff(6112660, 608320), 0, 0x20000)
			dword_5d4594_608316 = 0
			dword_5d4594_739392 = 0
			for i = (*byte)(unsafe.Pointer(nox_common_playerInfoGetFirst_416EA0())); i != nil; i = (*byte)(unsafe.Pointer(nox_common_playerInfoGetNext_416EE0((*nox_playerInfo)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(i))))))))) {
				*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(i))), unsafe.Sizeof(uint32(0))*1162))) = math.MaxUint32
			}
			v2 = (*byte)(unsafe.Pointer(nox_common_playerInfoFromNum_417090(31)))
			if v2 != nil {
				sub_425F10((*nox_playerInfo)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v2)))))))
			}
			result = (*byte)(unsafe.Pointer(nox_common_playerInfoGetFirst_416EA0()))
			for j = int32(uintptr(unsafe.Pointer(result))); result != nil; j = int32(uintptr(unsafe.Pointer(result))) {
				if int32(*(*uint8)(unsafe.Pointer(uintptr(j + 2064)))) != 31 {
					sub_425F10((*nox_playerInfo)(unsafe.Pointer(uintptr(j))))
				}
				result = (*byte)(unsafe.Pointer(nox_common_playerInfoGetNext_416EE0((*nox_playerInfo)(unsafe.Pointer(uintptr(j))))))
			}
		}
	}
	return result
}
func sub_4264D0() int32 {
	var i *byte
	if !noxflags.HasGame(noxflags.GameOnline) || noxflags.HasGame(noxflags.GameModeQuest) {
		if noxflags.HasGame(noxflags.GameModeQuest) {
			*memmap.PtrUint32(6112660, 739416) = uint32(libc.GetTime(nil) - libc.Time(dword_5d4594_600116))
		}
	} else {
		for i = (*byte)(unsafe.Pointer(nox_common_playerInfoGetFirst_416EA0())); i != nil; i = (*byte)(unsafe.Pointer(nox_common_playerInfoGetNext_416EE0((*nox_playerInfo)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(i))))))))) {
			sub_425E90(int32(uintptr(unsafe.Pointer(i))), 1)
		}
		*memmap.PtrUint32(6112660, 599504) = uint32(libc.GetTime(nil) - libc.Time(dword_5d4594_600116))
	}
	if !noxflags.HasGame(noxflags.GameModeQuest) {
		return sub_428810(int32(uintptr(memmap.PtrOff(6112660, 599476))), 2)
	}
	sub_4285C0(memmap.PtrInt16(6112660, 739396))
	return sub_428890(memmap.PtrInt16(6112660, 739396))
}
func sub_426590() unsafe.Pointer {
	return alloc.Calloc(1, 324)
}
func sub_4265A0(lpMem unsafe.Pointer) {
	var (
		v1 *unsafe.Pointer
		v2 int32
		i  int32
	)
	if lpMem != nil {
		v1 = (*unsafe.Pointer)(lpMem)
		v2 = 27
		for {
			if *v1 != nil {
				for i = 0; i < int32(uintptr(*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(unsafe.Pointer(nil))*1)))); i++ {
					*((*unsafe.Pointer)(unsafe.Add(unsafe.Pointer((*unsafe.Pointer)(*v1)), unsafe.Sizeof(unsafe.Pointer(nil))*uintptr(i*2)))) = nil
				}
				*v1 = nil
			}
			v1 = (*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(unsafe.Pointer(nil))*3))
			v2--
			if v2 == 0 {
				break
			}
		}
		lpMem = nil
	}
}
func sub_426600(a1 int32, a2 func(uint32, uint32)) int32 {
	var (
		v2     *int32
		v3     int32
		result int32
		i      int32
	)
	v2 = (*int32)(unsafe.Pointer(uintptr(a1 + 4)))
	v3 = 27
	for {
		result = *v2
		for i = 0; i < *v2; i++ {
			a2(*(*uint32)(unsafe.Pointer(uintptr(*((*int32)(unsafe.Add(unsafe.Pointer(v2), -int(unsafe.Sizeof(int32(0))*1)))) + i*8))), *(*uint32)(unsafe.Pointer(uintptr(*((*int32)(unsafe.Add(unsafe.Pointer(v2), -int(unsafe.Sizeof(int32(0))*1)))) + i*8 + 4))))
			result = *v2
		}
		v2 = (*int32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(int32(0))*3))
		v3--
		if v3 == 0 {
			break
		}
	}
	return result
}
func sub_4268B0(a1 int32) int32 {
	var (
		v1     int32
		result int32
	)
	v1 = 27
	result = a1 + 4
	for {
		result += 12
		v1--
		if v1 == 0 {
			break
		}
	}
	return result
}
func sub_426A20(a1 int32) {
	*memmap.PtrUint32(6112660, 739992) = uint32(a1)
}
func nox_xxx_wallGet_426A30() int32 {
	return int32(*memmap.PtrUint32(6112660, 739992))
}
func nox_xxx_mapGetWallSize_426A70() *byte {
	return (*byte)(memmap.PtrOff(6112660, 739980))
}
func nox_xxx_mapWall_426A80(a1 *int32) {
	*memmap.PtrUint32(6112660, 739980) = uint32(*(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*0)))
	*memmap.PtrUint32(6112660, 739984) = uint32(*(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*1)))
}
func nox_xxx_guide_427010(a1 *byte) int32 {
	var (
		v1 int32
		v2 **byte
	)
	v1 = 0
	v2 = (**byte)(memmap.PtrOff(0x587000, 70500))
	for libc.StrCmp(*v2, a1) != 0 {
		v2 = (**byte)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof((*byte)(nil))*1))
		v1++
		if int32(uintptr(unsafe.Pointer(v2))) >= int32(uintptr(memmap.PtrOff(0x587000, 70664))) {
			return 0
		}
	}
	return v1
}
func nox_xxx_guideNameByN_427230(a1 int32) *byte {
	return *(**byte)(memmap.PtrOff(0x587000, uint32(a1*4)+70500))
}
func nox_xxx_guiCreatureGetName_427240(a1 int32) int32 {
	var result int32
	if a1 > 0 && a1 < 41 && *memmap.PtrUint32(6112660, uint32(a1*28)+0xB4AF0) != 0 {
		result = int32(*memmap.PtrUint32(6112660, uint32(a1*28)+0xB4AEC))
	} else {
		result = 0
	}
	return result
}
func nox_xxx_guidesTTByName_427270(a1 *libc.WChar) int32 {
	var (
		v1 int32
		v2 *uint8
	)
	v1 = 1
	v2 = (*uint8)(memmap.PtrOff(6112660, 740104))
	for *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), unsafe.Sizeof(uint32(0))*1))) == 0 || nox_wcscmp(*(**libc.WChar)(unsafe.Pointer(v2)), a1) != 0 {
		v2 = (*uint8)(unsafe.Add(unsafe.Pointer(v2), 28))
		v1++
		if int32(uintptr(unsafe.Pointer(v2))) >= int32(uintptr(memmap.PtrOff(6112660, 741224))) {
			return 0
		}
	}
	return v1
}
func nox_xxx_creatureIsCharmableByTT_4272B0(a1 int32) int32 {
	var (
		result int32
		v2     *uint8
	)
	result = 1
	v2 = (*uint8)(memmap.PtrOff(6112660, 740108))
	for *(*uint32)(unsafe.Pointer(v2)) == 0 || *(*uint32)(unsafe.Pointer(v2)) != uint32(a1) {
		v2 = (*uint8)(unsafe.Add(unsafe.Pointer(v2), 28))
		result++
		if int32(uintptr(unsafe.Pointer(v2))) >= int32(uintptr(memmap.PtrOff(6112660, 741228))) {
			return 0
		}
	}
	return result
}
func nox_xxx_guideGetDescById_4272E0(a1 int32) int32 {
	return int32(*memmap.PtrUint32(6112660, uint32(a1*28)+0xB4AF4))
}
func nox_xxx_bookGetFirstCreMB_427300() int32 {
	var (
		result int32
		v1     *uint8
	)
	result = 1
	v1 = (*uint8)(memmap.PtrOff(6112660, 740108))
	for *(*uint32)(unsafe.Pointer(v1)) == 0 {
		v1 = (*uint8)(unsafe.Add(unsafe.Pointer(v1), 28))
		result++
		if int32(uintptr(unsafe.Pointer(v1))) >= int32(uintptr(memmap.PtrOff(6112660, 741228))) {
			return 0
		}
	}
	return result
}
func nox_xxx_bookGetNextCre_427320(a1 int32) int32 {
	var (
		result int32
		v2     *uint8
	)
	result = a1 + 1
	if a1+1 >= 41 {
		return 0
	}
	v2 = (*uint8)(memmap.PtrOff(6112660, uint32(result*28)+0xB4AF0))
	for *(*uint32)(unsafe.Pointer(v2)) == 0 {
		v2 = (*uint8)(unsafe.Add(unsafe.Pointer(v2), 28))
		result++
		if int32(uintptr(unsafe.Pointer(v2))) >= int32(uintptr(memmap.PtrOff(6112660, 741228))) {
			return 0
		}
	}
	return result
}
func sub_4273A0(a1 int32) int32 {
	return int32(*memmap.PtrUint32(6112660, uint32(a1*28)+0xB4B00))
}
func sub_4273C0(a1 int32) int32 {
	var result int32
	result = a1
	*memmap.PtrUint32(6112660, uint32(a1*28)+0xB4B00) = 1
	return result
}
func sub_4273E0(a1 int32) int32 {
	var result int32
	result = a1
	*memmap.PtrUint32(6112660, uint32(a1*28)+0xB4B00) = 0
	return result
}
func nox_xxx_bookGetCreatureImg_427400(a1 int32) int32 {
	var result int32
	if a1 <= 0 || a1 >= 41 {
		result = 0
	} else {
		result = int32(*memmap.PtrUint32(6112660, uint32(a1*28)+0xB4AFC))
	}
	return result
}
func sub_427430(a1 int32) int32 {
	var result int32
	if a1 <= 0 || a1 >= 41 {
		result = 0
	} else {
		result = int32(*memmap.PtrUint32(6112660, uint32(a1*28)+0xB4AF8))
	}
	return result
}
func nox_xxx_guideGetUnitSize_427460(a1 int32) uint8 {
	return *memmap.PtrUint8(6112660, uint32(a1*28)+740100)
}
func nox_xxx_journalEntryAdd_427490(a1 int32, a2 *byte, a3 int16) *uint8 {
	var (
		result *uint8
		v4     *uint8
		v5     int32
	)
	result = (*uint8)(alloc.Calloc(1, 76))
	v4 = result
	if result != nil {
		libc.MemSet(unsafe.Pointer(result), 0, 76)
		libc.StrNCpy((*byte)(unsafe.Pointer(result)), a2, 63)
		*(*uint8)(unsafe.Add(unsafe.Pointer(v4), 63)) = 0
		*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v4))), unsafe.Sizeof(uint16(0))*36))) = uint16(a3)
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v4))), unsafe.Sizeof(uint32(0))*17))) = 0
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v4))), unsafe.Sizeof(uint32(0))*16))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 3644)))
		v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 3644))))
		if v5 != 0 {
			*(*uint32)(unsafe.Pointer(uintptr(v5 + 68))) = uint32(uintptr(unsafe.Pointer(v4)))
		}
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 3644))) = uint32(uintptr(unsafe.Pointer(v4)))
		result = v4
	}
	return result
}
func nox_xxx_comJournalEntryAdd_427500(a1 int32, a2 *byte, a3 int16) {
	var (
		v3 int32
		v4 *uint8
	)
	v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))) + 276))))
	v4 = nox_xxx_journalEntryAdd_427490(v3, a2, a3)
	if v4 != nil {
		if int32(*(*uint8)(unsafe.Pointer(uintptr(v3 + 2064)))) == 31 {
			nox_xxx_cliBuildJournalString_469BC0()
		} else {
			nox_xxx_netSendJournalAdd_4D9440(int32(*(*uint8)(unsafe.Pointer(uintptr(v3 + 2064)))), int32(uintptr(unsafe.Pointer(v4))))
		}
	}
}
func nox_xxx_comAddEntryAll_427550(a1 *byte, a2 int16) int32 {
	var (
		result int32
		i      int32
	)
	result = int32(uintptr(unsafe.Pointer(nox_xxx_getFirstPlayerUnit_4DA7C0())))
	for i = result; result != 0; i = result {
		nox_xxx_comJournalEntryAdd_427500(i, a1, a2)
		result = int32(uintptr(unsafe.Pointer(nox_xxx_getNextPlayerUnit_4DA7F0((*nox_object_t)(unsafe.Pointer(uintptr(i)))))))
	}
	return result
}
func nox_xxx_journalEntryRemove_427590(a1 int32, a2 *byte) int32 {
	var (
		v2 int32
		v4 int32
		v5 int32
	)
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 3644))))
	if v2 == 0 {
		return 0
	}
	for libc.StrCmp((*byte)(unsafe.Pointer(uintptr(v2))), a2) != 0 {
		v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 64))))
		if v2 == 0 {
			return 0
		}
	}
	v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 68))))
	if v4 != 0 {
		*(*uint32)(unsafe.Pointer(uintptr(v4 + 64))) = *(*uint32)(unsafe.Pointer(uintptr(v2 + 64)))
	}
	v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 64))))
	if v5 != 0 {
		*(*uint32)(unsafe.Pointer(uintptr(v5 + 68))) = *(*uint32)(unsafe.Pointer(uintptr(v2 + 68)))
	}
	if uint32(v2) == *(*uint32)(unsafe.Pointer(uintptr(a1 + 3644))) {
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 3644))) = *(*uint32)(unsafe.Pointer(uintptr(v2 + 64)))
	}
	alloc.Free(unsafe.Pointer(uintptr(v2)))
	return 1
}
func nox_xxx_comJournalEntryRemove_427630(a1 int32, a2 *byte) {
	var v2 int32
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))) + 276))))
	if nox_xxx_journalEntryRemove_427590(v2, a2) != 0 {
		if int32(*(*uint8)(unsafe.Pointer(uintptr(v2 + 2064)))) == 31 {
			nox_xxx_cliBuildJournalString_469BC0()
		} else {
			nox_xxx_netSendJournalRemove_4D94A0(int32(*(*uint8)(unsafe.Pointer(uintptr(v2 + 2064)))), a2)
		}
	}
}
func nox_xxx_comRemoveEntryAll_427680(a1 *byte) int32 {
	var (
		result int32
		i      int32
	)
	result = int32(uintptr(unsafe.Pointer(nox_xxx_getFirstPlayerUnit_4DA7C0())))
	for i = result; result != 0; i = result {
		nox_xxx_comJournalEntryRemove_427630(i, a1)
		result = int32(uintptr(unsafe.Pointer(nox_xxx_getNextPlayerUnit_4DA7F0((*nox_object_t)(unsafe.Pointer(uintptr(i)))))))
	}
	return result
}
func nox_xxx_journalUpdateEntry_4276B0(a1 int32, a2 *byte, a3 int16) int32 {
	var (
		v3     int32
		result int32
	)
	v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 3644))))
	if v3 == 0 {
		return 0
	}
	for libc.StrCmp((*byte)(unsafe.Pointer(uintptr(v3))), a2) != 0 {
		v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(v3 + 64))))
		if v3 == 0 {
			return 0
		}
	}
	result = v3
	*(*uint16)(unsafe.Pointer(uintptr(v3 + 72))) = uint16(a3)
	return result
}
func nox_xxx_comJournalEntryUpdate_427720(a1 int32, a2 *byte, a3 int16) int32 {
	var (
		v3     int32
		result int32
	)
	v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))) + 276))))
	result = nox_xxx_journalUpdateEntry_4276B0(v3, a2, a3)
	if result != 0 {
		if int32(*(*uint8)(unsafe.Pointer(uintptr(v3 + 2064)))) != 31 {
			result = nox_xxx_netSendJournalUpdate_4D9500(int32(*(*uint8)(unsafe.Pointer(uintptr(v3 + 2064)))), result)
		}
	}
	return result
}
func nox_xxx_comUpdateEntryAll_427770(a1 *byte, a2 int16) int32 {
	var (
		result int32
		i      int32
	)
	result = int32(uintptr(unsafe.Pointer(nox_xxx_getFirstPlayerUnit_4DA7C0())))
	for i = result; result != 0; i = result {
		nox_xxx_comJournalEntryUpdate_427720(i, a1, a2)
		result = int32(uintptr(unsafe.Pointer(nox_xxx_getNextPlayerUnit_4DA7F0((*nox_object_t)(unsafe.Pointer(uintptr(i)))))))
	}
	return result
}
func sub_4277B0(a1 int32, a2 uint16) int32 {
	var (
		v2     int32
		result int32
		v4     int32
		v5     int32
		v6     int32
	)
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))) + 276))))
	result = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 3644))))
	if result != 0 {
		for {
			v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(result + 64))))
			if int32(a2)&int32(*(*uint16)(unsafe.Pointer(uintptr(result + 72)))) != 0 {
				v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(result + 68))))
				if v5 != 0 {
					*(*uint32)(unsafe.Pointer(uintptr(v5 + 64))) = *(*uint32)(unsafe.Pointer(uintptr(result + 64)))
				}
				v6 = int32(*(*uint32)(unsafe.Pointer(uintptr(result + 64))))
				if v6 != 0 {
					*(*uint32)(unsafe.Pointer(uintptr(v6 + 68))) = *(*uint32)(unsafe.Pointer(uintptr(result + 68)))
				}
				if uint32(result) == *(*uint32)(unsafe.Pointer(uintptr(v2 + 3644))) {
					*(*uint32)(unsafe.Pointer(uintptr(v2 + 3644))) = *(*uint32)(unsafe.Pointer(uintptr(result + 64)))
				}
				alloc.Free(unsafe.Pointer(uintptr(result)))
			}
			result = v4
			if v4 == 0 {
				break
			}
		}
	}
	return result
}
func sub_427820(a1 int32) *uint32 {
	var result *uint32
	sub_40D7D0(a1, 1)
	result = (*uint32)(unsafe.Pointer(uintptr(sub_41F800((*byte)(unsafe.Pointer(uintptr(a1 + 36)))))))
	if result != nil {
		*result |= 4
	}
	return result
}
func sub_427850(a1 int32) *uint32 {
	var result *uint32
	sub_40D7D0(a1, 0)
	result = (*uint32)(unsafe.Pointer(uintptr(sub_41F800((*byte)(unsafe.Pointer(uintptr(a1 + 36)))))))
	if result != nil {
		*result &= 0xFFFFFFFB
	}
	return result
}
func sub_427880(a1 *byte) int32 {
	var (
		v1 *byte
		v2 *uint8
	)
	v1 = a1
	if a1 != nil && (*a1 == 64 || *a1 == 32) {
		v1 = (*byte)(unsafe.Add(unsafe.Pointer(a1), 1))
	}
	v2 = (*uint8)(unsafe.Pointer(uintptr(sub_41F800(v1))))
	return bool2int(v2 != nil && int32(*v2)&4 != 0)
}
func sub_4278B0(a1 *float32, a2 *float32, a3 *float32) int32 {
	var (
		v3 float64
		v4 float64
	)
	*(*float32)(unsafe.Pointer(&dword_5d4594_741244)) = *(*float32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(float32(0))*2)) - *a1
	*(*float32)(unsafe.Pointer(&dword_5d4594_741248)) = *(*float32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(float32(0))*3)) - *(*float32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(float32(0))*1))
	*(*float32)(unsafe.Pointer(&dword_5d4594_741252)) = *(*float32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(float32(0))*2)) - *a2
	*(*float32)(unsafe.Pointer(&dword_5d4594_741256)) = *(*float32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(float32(0))*3)) - *(*float32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(float32(0))*1))
	*(*float32)(unsafe.Pointer(&dword_5d4594_741260)) = *a2 - *a1
	v3 = float64(*(*float32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(float32(0))*1)) - *(*float32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(float32(0))*1)))
	*mem_getFloatPtr(6112660, 0xB4F90) = float32(v3)
	*(*float32)(unsafe.Pointer(&dword_5d4594_741292)) = float32(v3*float64(*(*float32)(unsafe.Pointer(&dword_5d4594_741244))) - float64(*(*float32)(unsafe.Pointer(&dword_5d4594_741260))**(*float32)(unsafe.Pointer(&dword_5d4594_741248))))
	v4 = float64(*(*float32)(unsafe.Pointer(&dword_5d4594_741252))**(*float32)(unsafe.Pointer(&dword_5d4594_741248)) - *(*float32)(unsafe.Pointer(&dword_5d4594_741256))**(*float32)(unsafe.Pointer(&dword_5d4594_741244)))
	*(*float32)(unsafe.Pointer(&dword_5d4594_741284)) = float32(v4)
	if v4 == 0.0 {
		return 0
	}
	*a3 = *(*float32)(unsafe.Pointer(&dword_5d4594_741292))**(*float32)(unsafe.Pointer(&dword_5d4594_741252)) / *(*float32)(unsafe.Pointer(&dword_5d4594_741284)) + *a2
	*(*float32)(unsafe.Add(unsafe.Pointer(a3), unsafe.Sizeof(float32(0))*1)) = *(*float32)(unsafe.Pointer(&dword_5d4594_741292))**(*float32)(unsafe.Pointer(&dword_5d4594_741256)) / *(*float32)(unsafe.Pointer(&dword_5d4594_741284)) + *(*float32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(float32(0))*1))
	return 1
}
func sub_427980(a1 *float4, a2 *float4) int32 {
	var (
		v2     float64
		v3     float32
		v4     float32
		v5     float32
		result int32
		v7     float64
		v8     float32
		v9     float32
	)
	if float64(a1.field_0) >= float64(a1.field_8) {
		*mem_getFloatPtr(6112660, 0xB4F94) = a1.field_8
		v2 = float64(a1.field_0)
	} else {
		*mem_getFloatPtr(6112660, 0xB4F94) = a1.field_0
		v2 = float64(a1.field_8)
	}
	*mem_getFloatPtr(6112660, 0xB4F9C) = float32(v2)
	if float64(a1.field_4) >= float64(a1.field_C) {
		*mem_getFloatPtr(6112660, 0xB4F98) = a1.field_C
		v3 = a1.field_4
	} else {
		*mem_getFloatPtr(6112660, 0xB4F98) = a1.field_4
		v3 = a1.field_C
	}
	*mem_getFloatPtr(6112660, 0xB4FA0) = v3
	if float64(a2.field_0) >= float64(a2.field_8) {
		*mem_getFloatPtr(6112660, 0xB4F6C) = a2.field_8
		v4 = a2.field_0
	} else {
		*mem_getFloatPtr(6112660, 0xB4F6C) = a2.field_0
		v4 = a2.field_8
	}
	*mem_getFloatPtr(6112660, 0xB4F74) = v4
	if float64(a2.field_4) >= float64(a2.field_C) {
		*mem_getFloatPtr(6112660, 0xB4F70) = a2.field_C
		v5 = a2.field_4
	} else {
		*mem_getFloatPtr(6112660, 0xB4F70) = a2.field_4
		v5 = a2.field_C
	}
	*mem_getFloatPtr(6112660, 741240) = v5
	if v2 < float64(*mem_getFloatPtr(6112660, 0xB4F6C)) || float64(*mem_getFloatPtr(6112660, 0xB4F94)) > float64(*mem_getFloatPtr(6112660, 0xB4F74)) || float64(*mem_getFloatPtr(6112660, 0xB4FA0)) < float64(*mem_getFloatPtr(6112660, 0xB4F70)) || float64(*mem_getFloatPtr(6112660, 0xB4F98)) > float64(*mem_getFloatPtr(6112660, 741240)) {
		goto LABEL_36
	}
	if a1.field_4 == a1.field_C && a2.field_4 == a2.field_C {
		return 1
	}
	*(*float32)(unsafe.Pointer(&dword_5d4594_741244)) = a1.field_8 - a1.field_0
	*(*float32)(unsafe.Pointer(&dword_5d4594_741248)) = a1.field_C - a1.field_4
	*(*float32)(unsafe.Pointer(&dword_5d4594_741252)) = a2.field_8 - a2.field_0
	*(*float32)(unsafe.Pointer(&dword_5d4594_741256)) = a2.field_C - a2.field_4
	*(*float32)(unsafe.Pointer(&dword_5d4594_741260)) = a2.field_0 - a1.field_0
	v7 = float64(a2.field_4 - a1.field_4)
	*mem_getFloatPtr(6112660, 0xB4F90) = float32(v7)
	*(*float32)(unsafe.Pointer(&dword_5d4594_741292)) = float32(v7*float64(*(*float32)(unsafe.Pointer(&dword_5d4594_741244))) - float64(*(*float32)(unsafe.Pointer(&dword_5d4594_741260))**(*float32)(unsafe.Pointer(&dword_5d4594_741248))))
	*(*float32)(unsafe.Pointer(&dword_5d4594_741284)) = *(*float32)(unsafe.Pointer(&dword_5d4594_741252))**(*float32)(unsafe.Pointer(&dword_5d4594_741248)) - *(*float32)(unsafe.Pointer(&dword_5d4594_741256))**(*float32)(unsafe.Pointer(&dword_5d4594_741244))
	if float64(*(*float32)(unsafe.Pointer(&dword_5d4594_741292))) == 0.0 || float64(*(*float32)(unsafe.Pointer(&dword_5d4594_741284))) == 0.0 || float64(*(*float32)(unsafe.Pointer(&dword_5d4594_741292))) < 0.0 && float64(*(*float32)(unsafe.Pointer(&dword_5d4594_741284))) > 0.0 {
		goto LABEL_36
	}
	if float64(*(*float32)(unsafe.Pointer(&dword_5d4594_741292))) > 0.0 && float64(*(*float32)(unsafe.Pointer(&dword_5d4594_741284))) < 0.0 || (func() bool {
		v8 = float32(sub_419A10(*(*float32)(unsafe.Pointer(&dword_5d4594_741292))))
		return sub_419A10(*(*float32)(unsafe.Pointer(&dword_5d4594_741284))) < float64(v8)
	}()) || (func() bool {
		*memmap.PtrUint32(6112660, 741288) = dword_5d4594_741284
		*(*float32)(unsafe.Pointer(&dword_5d4594_741296)) = *mem_getFloatPtr(6112660, 0xB4F90)**(*float32)(unsafe.Pointer(&dword_5d4594_741252)) - *(*float32)(unsafe.Pointer(&dword_5d4594_741260))**(*float32)(unsafe.Pointer(&dword_5d4594_741256))
		return float64(*(*float32)(unsafe.Pointer(&dword_5d4594_741296))) < 0.0
	}()) && float64(*(*float32)(unsafe.Pointer(&dword_5d4594_741284))) > 0.0 || float64(*(*float32)(unsafe.Pointer(&dword_5d4594_741296))) > 0.0 && float64(*(*float32)(unsafe.Pointer(&dword_5d4594_741284))) < 0.0 || (func() bool {
		v9 = float32(sub_419A10(*(*float32)(unsafe.Pointer(&dword_5d4594_741296))))
		return sub_419A10(*mem_getFloatPtr(6112660, 0xB4FA8)) < float64(v9)
	}()) {
	LABEL_36:
		result = 0
	} else {
		result = 1
	}
	return result
}
func sub_427C80(a1 *int4, a2 *int4) int32 {
	var (
		v2  int32
		v3  int32
		v4  int32
		v5  int32
		v6  int32
		v7  int32
		v8  int32
		v10 int32
		v11 int32
		v12 int32
		v13 int32
		v14 int32
		v15 int32
		v16 int32
		v17 int32
		v18 int32
		v19 int32
		v20 int32
		v21 int32
		v22 *int4
		v23 *int4
	)
	v2 = a1.field_0
	v3 = a1.field_4
	v4 = a1.field_8
	v19 = a1.field_C
	v5 = a2.field_0
	v21 = a2.field_4
	v17 = a2.field_C
	v6 = a2.field_8
	v7 = v4 - v2
	v8 = a2.field_0 - v6
	v20 = a2.field_0
	if v4-v2 >= 0 {
		v6 = a2.field_8
		v22 = (*int4)(unsafe.Pointer(uintptr(v2)))
	} else {
		v22 = (*int4)(unsafe.Pointer(uintptr(v4)))
		v4 = v2
	}
	if v8 <= 0 {
		if v4 < v5 || v6 < int32(uintptr(unsafe.Pointer(v22))) {
			return 0
		}
	} else if v4 < v6 || v5 < int32(uintptr(unsafe.Pointer(v22))) {
		return 0
	}
	v10 = v19
	v11 = v21 - v17
	v18 = v19 - v3
	if v19-v3 >= 0 {
		v23 = (*int4)(unsafe.Pointer(uintptr(v3)))
	} else {
		v23 = (*int4)(unsafe.Pointer(uintptr(v19)))
		v10 = v3
	}
	if v11 <= 0 {
		if v10 < v21 || v17 < int32(uintptr(unsafe.Pointer(v23))) {
			return 0
		}
	} else {
		if v10 < v17 {
			return 0
		}
		if v21 < int32(uintptr(unsafe.Pointer(v23))) {
			return 0
		}
	}
	v12 = v2 - v20
	v13 = v3 - v21
	v14 = v11*(v2-v20) - v8*(v3-v21)
	v15 = v8*v18 - v7*v11
	if v15 <= 0 {
		if v14 > 0 || v14 < v15 {
			return 0
		}
	} else {
		if v14 < 0 {
			return 0
		}
		if v14 > v15 {
			return 0
		}
	}
	v16 = v7*v13 - v18*v12
	if v15 <= 0 {
		if v16 <= 0 && v16 >= v15 {
			return 1
		}
		return 0
	}
	if v16 < 0 {
		return 0
	}
	if v16 > v15 {
		return 0
	}
	return 1
}
func sub_427DF0(a1 int32, a2 *int32, a3 float32) int32 {
	var (
		v3  float64
		v4  float64
		v6  float64
		v7  float32
		v8  float32
		v9  float32
		v10 float32
		v11 float32
		v12 float32
		v13 float32
		v14 float32
		v15 float32
		v16 float32
		v17 float32
		v18 float32
	)
	v16 = float32(float64(*a2))
	v17 = float32(float64(*(*int32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(int32(0))*1))))
	v8 = float32(float64(*(*int32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(int32(0))*2))) - float64(v16))
	v3 = float64(*(*int32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(int32(0))*3))) - float64(v17)
	v11 = float32(v3)
	v7 = nox_double2float(math.Sqrt(v3*float64(v11) + float64(v8*v8)))
	v14 = v8 / v7
	v4 = float64(v11 / v7)
	v15 = float32(v4)
	v9 = float32(float64(*(*int32)(unsafe.Pointer(uintptr(a1)))) - float64(v16))
	v12 = float32(float64(*(*int32)(unsafe.Pointer(uintptr(a1 + 4)))) - float64(v17))
	if nox_double2float(math.Abs(float64(v9)*(-v4)+float64(v12*v14))) > a3 {
		return 0
	}
	v6 = float64(v12*v15 + v9*v14)
	v18 = float32(v6)
	if v6 > float64(v7) {
		return 0
	}
	if float64(v18) < 0.0 {
		return 0
	}
	v10 = v18*v14 + v16
	v13 = v18*v15 + v17
	*(*uint32)(unsafe.Pointer(uintptr(a1))) = uint32(int(v10))
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 4))) = uint32(int(v13))
	return 1
}
