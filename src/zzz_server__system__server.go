package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	noxflags "github.com/noxworld-dev/opennox/v1/common/flags"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"math"
	"unsafe"
)

var nox_xxx_resetMapInit_1569652 uint32 = 0
var nox_server_xxx_1599716 [65536]nox_server_xxx = [65536]nox_server_xxx{}

func nullsub_21() {
}
func nullsub_25(a1 uint32) {
}
func sub_426060() {
	var (
		v0     *byte
		i      *byte
		v2     *byte
		j      *byte
		result unsafe.Pointer
		v5     *byte
		v6     *byte
		v7     int32
	)
	dword_5d4594_608316 = 0
	dword_5d4594_600116 = uint32(libc.GetTime(nil))
	v7 = int32(sub_5545A0())
	v0 = sub_554230()
	sub_4282D0(v0, v7)
	if !noxflags.HasGame(noxflags.GameOnline) || noxflags.HasGame(noxflags.GameModeQuest) {
		result = unsafe.Pointer(uintptr(bool2int(noxflags.HasGame(noxflags.GameModeQuest))))
		if result != nil {
			result = unsafe.Pointer(uintptr(bool2int(noxflags.HasGame(noxflags.GameModeQuest))))
			if result != nil {
				v5 = (*byte)(sub_416640())
				v6 = nox_xxx_cliGamedataGet_416590(0)
				*memmap.PtrUint16(6112660, 739396) = uint16(int16(sub_40A770()))
				*memmap.PtrUint32(6112660, 739400) = *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v5))), unsafe.Sizeof(uint32(0))*10)))
				*memmap.PtrUint32(6112660, 739404) = uint32(sub_4200E0())
				*memmap.PtrUint8(6112660, 739412) = uint8(int8(bool2int((*(*byte)(unsafe.Add(unsafe.Pointer(v6), 53)) & 192) != 0)))
				*memmap.PtrUint32(6112660, 739408) = 5
				libc.StrNCpy((*byte)(memmap.PtrOff(6112660, 739676)), (*byte)(unsafe.Add(unsafe.Pointer(v6), 9)), 15)
				*memmap.PtrUint8(6112660, 739691) = 0
				libc.StrNCpy((*byte)(memmap.PtrOff(6112660, 739420)), v6, 8)
				*memmap.PtrUint8(6112660, 739428) = 0
				sub_4289D0((*unsafe.Pointer)(memmap.PtrOff(6112660, 739396)))
			}
		}
	} else {
		for i = (*byte)(unsafe.Pointer(nox_common_playerInfoGetFirst_416EA0())); i != nil; i = (*byte)(unsafe.Pointer(nox_common_playerInfoGetNext_416EE0((*nox_playerInfo)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(i))))))))) {
			*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(i))), unsafe.Sizeof(uint32(0))*1162))) = math.MaxUint32
		}
		v2 = (*byte)(unsafe.Pointer(nox_common_playerInfoFromNum_417090(31)))
		if v2 != nil {
			sub_425F10((*nox_playerInfo)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v2)))))))
		}
		for j = (*byte)(unsafe.Pointer(nox_common_playerInfoGetFirst_416EA0())); j != nil; j = (*byte)(unsafe.Pointer(nox_common_playerInfoGetNext_416EE0((*nox_playerInfo)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(j))))))))) {
			if *(*byte)(unsafe.Add(unsafe.Pointer(j), 2064)) != 31 {
				sub_425F10((*nox_playerInfo)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(j)))))))
			}
		}
		sub_426150()
		sub_428810(int32(uintptr(memmap.PtrOff(6112660, 599476))), 0)
		*memmap.PtrUint16(6112660, 599482) = 0
	}
}
func sub_446040() int32 {
	return int32(dword_5d4594_825768)
}
func sub_46DCB0() int32 {
	return int32(uintptr(unsafe.Pointer(sub_46DCC0())))
}
func nox_xxx_getSomeMapName_4D0CF0() *byte {
	var (
		v0 int32
		v1 int32
		v2 int32
		v4 int32
	)
	v0 = int32(nox_common_gameFlags_getVal_40A5B0())
	v1 = sub_4D0D50(v0)
	v2 = int32(*memmap.PtrUint32(6112660, uint32(v1*4)+0x17A08C))
	if v2 == 0 {
		return nil
	}
	v4 = int32(*memmap.PtrUint32(6112660, uint32(v1*4)+0x17A0A4))
	if v4 > v2 {
		*memmap.PtrUint32(6112660, uint32(v1*4)+0x17A0A4) = 0
		v4 = 0
	}
	*memmap.PtrUint32(6112660, uint32(v1*4)+0x17A0A4) = (*memmap.PtrUint32(6112660, uint32(v1*4)+0x17A0A4) + 1) % uint32(v2)
	return (*byte)(memmap.PtrOff(6112660, uint32((v4+v1*20+v1*5)*128)+0x17558C))
}
func sub_4D10F0(a1 *byte) {
	var (
		v1 int32
		i  *byte
		v3 int32
		v4 int32
		v5 int32
		v6 int32
		v7 *uint8
	)
	if a1 != nil {
		v1 = 0
		if dword_5d4594_1548476 > 0 {
			for i = (*byte)(memmap.PtrOff(6112660, 1525136)); libc.StrCaseCmp(i, a1) != 0; i = (*byte)(unsafe.Add(unsafe.Pointer(i), 32)) {
				if func() int32 {
					p := &v1
					*p++
					return *p
				}() >= *(*int32)(unsafe.Pointer(&dword_5d4594_1548476)) {
					return
				}
			}
			v3 = int32(dword_5d4594_1548476)
			v4 = v1 * 32
			v5 = int32(dword_5d4594_1548480)
			*memmap.PtrUint32(0x587000, 191880) = uint32(v1)
			*memmap.PtrUint32(6112660, uint32(v4)+0x1745A4) = *memmap.PtrUint32(6112660, uint32(v1*32)+0x1745A4) + 1
			v6 = 0
			*memmap.PtrUint32(6112660, uint32(v4)+1525160) = uint32(v5)
			if v3 > 0 {
				v7 = (*uint8)(memmap.PtrOff(6112660, 1525156))
				for {
					if *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v7))), -int(unsafe.Sizeof(uint32(0))*6)))) == *memmap.PtrUint32(6112660, uint32(v1*32)+0x17458C) && v1 != v6 {
						*(*uint32)(unsafe.Pointer(v7))++
						*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v7))), unsafe.Sizeof(uint32(0))*1))) = uint32(v5)
					}
					v6++
					v7 = (*uint8)(unsafe.Add(unsafe.Pointer(v7), 32))
					if v6 >= v3 {
						break
					}
				}
			}
			dword_5d4594_1548480 = uint32(v5 + 1)
		}
	}
}
func nox_xxx____setargv_14_4D15F0() {
	*memmap.PtrUint32(6112660, 1548520) = 1
}
func sub_4D2160() {
	var (
		result *byte
		i      *byte
		v2     int32
		j      *byte
		v4     int32
		v5     int32
	)
	v5 = nox_common_playerInfoCount_416F40()
	result = (*byte)(unsafe.Pointer(nox_server_teamFirst_418B10()))
	for i = result; result != nil; i = result {
		if *(*uint32)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(i))), unsafe.Sizeof(uint32(0))*19))) + 492))) == 0 && sub_418BC0(int32(uintptr(unsafe.Pointer(i)))) != 0 {
			for {
				v2 = noxRndCounter1.IntClamp(0, v5-1)
				for j = (*byte)(unsafe.Pointer(nox_common_playerInfoGetFirst_416EA0())); j != nil; j = (*byte)(unsafe.Pointer(nox_common_playerInfoGetNext_416EE0((*nox_playerInfo)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(j))))))))) {
					if v2 == 0 {
						break
					}
					v2--
				}
				v4 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(j))), unsafe.Sizeof(uint32(0))*514))))
				if !(v4 == 0 || nox_xxx_teamCompare2_419180(v4+48, uint8(*(*byte)(unsafe.Add(unsafe.Pointer(i), 57)))) == 0) {
					break
				}
			}
			sub_4F3400(int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(j))), unsafe.Sizeof(uint32(0))*514)))), int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(i))), unsafe.Sizeof(uint32(0))*19)))), 1)
		}
		result = (*byte)(unsafe.Pointer(nox_server_teamNext_418B60((*nox_team_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(i)))))))))
	}
}
func sub_4D2230() {
	var (
		v0 int16
		v1 int16
		v2 int16
	)
	if noxflags.HasGame(noxflags.GameFlag15 | noxflags.GameFlag16) {
		v0 = int16(uint16(nox_common_gameFlags_getVal_40A5B0()))
		if int32(nox_xxx_servGamedataGet_40A020(v0)) == 0 {
			v1 = int16(uint16(nox_common_gameFlags_getVal_40A5B0()))
			if int32(sub_40A180(v1)) == 0 {
				v2 = int16(uint16(nox_common_gameFlags_getVal_40A5B0()))
				sub_409FB0_settings(v2, 15)
			}
		}
	} else if noxflags.HasGame(noxflags.GameModeElimination) && int32(nox_xxx_servGamedataGet_40A020(1024)) == 0 {
		sub_409FB0_settings(1024, 15)
	}
}
func sub_4D22B0() {
	var (
		result *byte
		i      int32
		v2     int32
		v3     int32
		v4     int32
		v5     int32
		v6     int32
	)
	result = (*byte)(unsafe.Pointer(nox_common_playerInfoGetFirst_416EA0()))
	for i = int32(uintptr(unsafe.Pointer(result))); result != nil; i = int32(uintptr(unsafe.Pointer(result))) {
		v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(i + 2056))))
		if v2 != 0 {
			v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 748))))
			if *(*uint32)(unsafe.Pointer(uintptr(v3 + 280))) != 0 {
				nox_xxx_shopCancelSession_510DC0(*(**uint32)(unsafe.Pointer(uintptr(v3 + 280))))
			}
			*(*uint32)(unsafe.Pointer(uintptr(v3 + 280))) = 0
			nox_xxx_playerCmd_51AC30(int32(*(*uint8)(unsafe.Pointer(uintptr(i + 2064)))))
			if noxflags.HasGame(noxflags.GameModeQuest) {
				v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(i + 4676))))
				*(*uint32)(unsafe.Pointer(uintptr(i + 4676))) = 0
				*(*uint32)(unsafe.Pointer(uintptr(i + 4680))) = uint32(v4)
			}
			if nox_xxx_playerSetState_4FA020((*nox_object_t)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(uintptr(i + 2056))))), 13) == 0 || !noxflags.HasGame(noxflags.GameModeSolo10) {
				*(*uint32)(unsafe.Pointer(uintptr(i + 4700))) = 0
				v6 = int32((uint32(sub_4CFE00()) >> 1) & 1)
				v5 = nox_xxx_gameIsSwitchToSolo_4DB240()
				nox_xxx_playerMakeDefItems_4EF7D0(int32(*(*uint32)(unsafe.Pointer(uintptr(i + 2056)))), bool2int(v5 == 0), v6)
			}
			if int32(*(*uint8)(unsafe.Pointer(uintptr(i + 3680))))&32 != 0 {
				nox_xxx_playerLeaveObserver_0_4E6AA0((*nox_playerInfo)(unsafe.Pointer(uintptr(i))))
				nox_xxx_playerCameraUnlock_4E6040(int32(*(*uint32)(unsafe.Pointer(uintptr(i + 2056)))))
			}
			*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(i + 2056))) + 136))) = nox_frame_xxx_2598000
		}
		result = (*byte)(unsafe.Pointer(nox_common_playerInfoGetNext_416EE0((*nox_playerInfo)(unsafe.Pointer(uintptr(i))))))
	}
}
func nox_xxx_updateServer_4D2DA0(a1 int64) int32 {
	var (
		result int32
		v2     *byte
		v3     int32
		v4     *byte
		v5     float64
		v6     int32
		v7     int32
		v8     *byte
		v9     float32
		v10    *uint8
	)
	if dword_5d4594_528252 == 1 && nox_frame_xxx_2598000 == dword_5d4594_528260 {
		nox_xxx_reconAttempt_41E390()
	}
	sub_5096F0()
	result = bool2int(noxflags.HasGame(noxflags.GameFlag4))
	if result != 0 {
		return result
	}
	if noxflags.HasGame(noxflags.GameOnline) {
		sub_416720()
		if !noxflags.HasGame(noxflags.GameModeChat) {
			if sub_409F40(8192) != 0 {
				sub_4D7CC0()
			}
		}
	}
	if !noxflags.HasGame(noxflags.GameModeElimination) {
		goto LABEL_31
	}
	if sub_40AA00() != 0 {
		if dword_5d4594_1548704 == 0 {
			sub_4D2FF0()
		}
	} else {
		dword_5d4594_1548704 = 0
	}
	if noxflags.HasGame(noxflags.GameSuddenDeath) {
		goto LABEL_31
	}
	if sub_40A300() != 0 {
		goto LABEL_31
	}
	v2 = (*byte)(unsafe.Pointer(nox_common_playerInfoGetFirst_416EA0()))
	if v2 == nil {
		goto LABEL_31
	}
	for *((*int32)(unsafe.Add(unsafe.Pointer((*int32)(unsafe.Pointer(v2))), unsafe.Sizeof(int32(0))*535))) <= 0 {
		v2 = (*byte)(unsafe.Pointer(nox_common_playerInfoGetNext_416EE0((*nox_playerInfo)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v2)))))))))
		if v2 == nil {
			goto LABEL_31
		}
	}
	if sub_40AA00() == 0 {
		goto LABEL_31
	}
	if nox_xxx_CheckGameplayFlags_417DA0(4) {
		v3 = int32(nox_xxx_getTeamCounter_417DD0())
		if v3 >= sub_40AA40() {
			goto LABEL_31
		}
		v4 = (*byte)(unsafe.Pointer(nox_server_teamFirst_418B10()))
		if v4 != nil {
			for nox_xxx_countNonEliminatedPlayersInTeam_40A830(int32(uintptr(unsafe.Pointer(v4)))) != 1 {
				v4 = (*byte)(unsafe.Pointer(nox_server_teamNext_418B60((*nox_team_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v4)))))))))
				if v4 == nil {
					goto LABEL_31
				}
			}
			v10 = (*uint8)(memmap.PtrOff(0x587000, 197200))
			v5 = nox_xxx_gamedataGetFloat_419D40(CString("SuddenDeathCountdown"))
			v9 = float32(v5)
			v7 = int(v9)
			nox_xxx_servStartCountdown_40A2A0(v7, (*byte)(unsafe.Pointer(v10)))
		}
	} else {
		v6 = sub_40A770()
		if v6 < sub_40AA40() {
			v10 = (*uint8)(memmap.PtrOff(0x587000, 197256))
			v5 = nox_xxx_gamedataGetFloat_419D40(CString("SuddenDeathCountdown"))
			v9 = float32(v5)
			v7 = int(v9)
			nox_xxx_servStartCountdown_40A2A0(v7, (*byte)(unsafe.Pointer(v10)))
		}
	}
LABEL_31:
	if sub_40A6B0() != 0 {
		v8 = (*byte)(sub_416640())
		sub_4D9700(int32(*(*uint32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v8), 66))))))
		sub_40A6A0(0)
	}
	if (uint64(a1) - *memmap.PtrUint64(6112660, 1548692)) > 500 {
		nox_xxx_netReportAllLatency_4D3050()
		*memmap.PtrUint64(6112660, 1548692) = uint64(a1)
	}
	if noxflags.HasGame(noxflags.GameModeChat) && sub_40A740() == 0 && int32(nox_xxx_getTeamCounter_417DD0()) != 0 && !nox_xxx_CheckGameplayFlags_417DA0(2) {
		sub_4183C0()
	}
	result = bool2int(noxflags.HasGame(noxflags.GameModeQuest))
	if result != 0 {
		sub_4D7150()
		sub_4D71F0()
		nox_server_checkWarpGate_4D7600()
		sub_4DCE00()
		result = sub_4D7A80()
	}
	return result
}
func sub_4D7140(a1 int32) int32 {
	var result int32
	result = a1
	dword_5d4594_1556144 = uint32(a1)
	return result
}
func sub_4DB160() *byte {
	return (*byte)(memmap.PtrOff(6112660, 1557900))
}
func sub_4DBA30(a1 int32) {
	var (
		result *byte
		v2     *byte
		v3     int32
		v4     int32
		v5     int32
		v6     int32
		v7     int32
		v8     int32
		v9     *int32
		v10    int32
		v11    int32
		v12    int32
		v13    int32
		v14    *int32
		v15    int32
		v16    int32
		v17    int32
		v18    int32
		v19    int32
		i      int32
		v23    int32
		v24    *byte
		v25    int32
		v26    int32
	)
	result = (*byte)(unsafe.Pointer(nox_common_playerInfoFromNum_417090(31)))
	v2 = result
	v3 = 0
	v24 = result
	if result != nil {
		v4 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(result))), unsafe.Sizeof(uint32(0))*514))))
		v25 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(result))), unsafe.Sizeof(uint32(0))*514))))
		if v4 != 0 {
			if *memmap.PtrUint32(6112660, 1563128) == 0 {
				*memmap.PtrUint32(6112660, 1563128) = uint32(nox_xxx_getNameId_4E3AA0(CString("SaveGameLocation")))
				*memmap.PtrUint32(6112660, 1563132) = uint32(nox_xxx_getNameId_4E3AA0(CString("Glyph")))
			}
			if a1 == 1 {
				v5 = int32(uintptr(unsafe.Pointer(noxServer.firstServerObject())))
				if v5 != 0 {
					for {
						v26 = int32(uintptr(unsafe.Pointer((*nox_object_t)(unsafe.Pointer(uintptr(v5))).Next())))
						if uint32(*(*uint16)(unsafe.Pointer(uintptr(v5 + 4)))) == *memmap.PtrUint32(6112660, 1563128) {
							nox_xxx_unitMove_4E7010((*nox_object_t)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), unsafe.Sizeof(uint32(0))*514)))))), (*float2)(unsafe.Pointer(uintptr(v5+56))))
							*(*uint32)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), unsafe.Sizeof(uint32(0))*514))) + 44))) = *(*uint32)(unsafe.Pointer(uintptr(v5 + 44)))
							*(*uint32)(unsafe.Pointer(uintptr(v5 + 44))) = 0
							nox_xxx_delayedDeleteObject_4E5CC0((*nox_object_t)(unsafe.Pointer(uintptr(v5))))
						} else {
							v6 = int32(*(*uint32)(unsafe.Pointer(uintptr(v5 + 8))))
							if v6&2 != 0 {
								v7 = int32(*(*uint32)(unsafe.Pointer(uintptr(v5 + 748))))
								v8 = 0
								if int32(*(*uint8)(unsafe.Pointer(uintptr(v7 + 1129)))) != 0 {
									v9 = (*int32)(unsafe.Pointer(uintptr(v7 + 1132)))
									for {
										v10 = sub_4ECF10(*v9)
										*v9 = v10
										if v10 == 0 {
											*(*uint8)(unsafe.Pointer(uintptr(v7 + 1129))) = 0
										}
										v8++
										v9 = (*int32)(unsafe.Add(unsafe.Pointer(v9), unsafe.Sizeof(int32(0))*1))
										if v8 >= int32(*(*uint8)(unsafe.Pointer(uintptr(v7 + 1129)))) {
											break
										}
									}
									v3 = 0
								}
								*(*uint32)(unsafe.Pointer(uintptr(v7 + 1196))) = uint32(sub_4ECF10(int32(*(*uint32)(unsafe.Pointer(uintptr(v7 + 1196))))))
								if int32(*(*uint32)(unsafe.Pointer(uintptr(v5 + 16)))) >= 0 {
									sub_52BAF0(v5)
								}
								v11 = sub_4ECF10(int32(*(*uint32)(unsafe.Pointer(uintptr(v7 + 392)))))
								if v11 != 0 {
									*(*uint32)(unsafe.Pointer(uintptr(v7 + 392))) = *(*uint32)(unsafe.Pointer(uintptr(v11 + 36)))
								} else {
									*(*uint32)(unsafe.Pointer(uintptr(v7 + 392))) = 0
								}
								v12 = sub_4ECF10(int32(*(*uint32)(unsafe.Pointer(uintptr(v7 + 1200)))))
								if v12 != 0 {
									*(*uint32)(unsafe.Pointer(uintptr(v7 + 1200))) = *(*uint32)(unsafe.Pointer(uintptr(v12 + 36)))
								} else {
									*(*uint32)(unsafe.Pointer(uintptr(v7 + 1200))) = 0
								}
								*(*uint32)(unsafe.Pointer(uintptr(v7 + 1216))) = uint32(sub_4ECF10(int32(*(*uint32)(unsafe.Pointer(uintptr(v7 + 1216))))))
								v13 = 0
								if int32(*(*uint8)(unsafe.Pointer(uintptr(v7 + 2172)))) != 0 {
									v14 = (*int32)(unsafe.Pointer(uintptr(v7 + 2140)))
									for {
										v15 = sub_4ECF10(*v14)
										if v15 != 0 {
											*v14 = int32(*(*uint32)(unsafe.Pointer(uintptr(v15 + 36))))
										} else {
											*(*uint8)(unsafe.Pointer(uintptr(v7 + 2172))) = 0
										}
										v13++
										v14 = (*int32)(unsafe.Add(unsafe.Pointer(v14), unsafe.Sizeof(int32(0))*1))
										if v13 >= int32(*(*uint8)(unsafe.Pointer(uintptr(v7 + 2172)))) {
											break
										}
									}
								}
								v2 = v24
							} else if v6&0x4000 != 0 {
								if *(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v5 + 748))) + 16))) != 0 {
									nox_xxx_unitNeedSync_4E44F0(v5)
								}
							} else if (v6 & 0x8000) == 0 {
								if (v6&128) != 0 && *(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v5 + 748))) + 12))) != *(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v5 + 748))) + 4))) {
									nox_xxx_unitAddToUpdatable_4DA8D0((*nox_object_t)(unsafe.Pointer(uintptr(v5))))
								}
							} else {
								v23 = int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v5 + 748))) + 4))))
								if v23 != 0 && *(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v23 + 748))) + 16))) != 0 {
									nox_xxx_unitNeedSync_4E44F0(v5)
								}
							}
						}
						v5 = v26
						if v26 == 0 {
							break
						}
					}
				}
				nox_script_activatorResolveObjs_51B0C0()
				sub_516FC0()
				if dword_5d4594_1563096 != 0 {
					v16 = int32(uintptr(unsafe.Pointer(noxServer.firstServerObject())))
					if v16 != 0 {
						for {
							v3 = int32(uintptr(unsafe.Pointer((*nox_object_t)(unsafe.Pointer(uintptr(v16))).Next())))
							if int32(*(*uint32)(unsafe.Pointer(uintptr(v16 + 16)))) >= 0 && nox_xxx_isUnit_4E5B50((*nox_object_t)(unsafe.Pointer(uintptr(v16)))) != 0 {
								if int32(*(*uint8)(unsafe.Pointer(uintptr(v16 + 8))))&2 != 0 {
									v17 = int32(*(*uint32)(unsafe.Pointer(uintptr(v16 + 12))))
									if v17&8192 != 0 {
										v18 = v16.FirstItem()
										if v18 != 0 {
											for {
												v19 = nox_xxx_inventoryGetNext_4E7990(v18)
												if uint32(*(*uint16)(unsafe.Pointer(uintptr(v18 + 4)))) == *memmap.PtrUint32(6112660, 1563132) {
													nox_xxx_delayedDeleteObject_4E5CC0((*nox_object_t)(unsafe.Pointer(uintptr(v18))))
												}
												v18 = v19
												if v19 == 0 {
													break
												}
											}
										}
									}
								}
								nox_xxx_delayedDeleteObject_4E5CC0((*nox_object_t)(unsafe.Pointer(uintptr(v16))))
							}
							v16 = v3
							if v3 == 0 {
								break
							}
						}
					}
					var obj *nox_object_t = firstServerObjectUpdatable2()
					if unsafe.Pointer(obj) != unsafe.Pointer(uintptr(v3)) {
						for {
							{
								var v21 *nox_object_t = obj.Next()
								if int32(obj.obj_flags) >= 0 {
									if sub_4E5B80(obj) != 0 {
										nox_xxx_delayedDeleteObject_4E5CC0(obj)
									}
								}
								obj = v21
							}
							if unsafe.Pointer(obj) == unsafe.Pointer(uintptr(v3)) {
								break
							}
						}
					}
					v2 = v24
				}
				v4 = v25
			}
			for i = int32(*(*uint32)(unsafe.Pointer(uintptr(v4 + 516)))); i != v3; i = int32(*(*uint32)(unsafe.Pointer(uintptr(i + 512)))) {
				if int32(*(*uint8)(unsafe.Pointer(uintptr(i + 8))))&2 != 0 {
					if int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(i + 748))) + 1440))))&128 != 0 {
						nox_xxx_netReportAcquireCreature_4D91A0(int32(uint8(*(*byte)(unsafe.Add(unsafe.Pointer(v2), 2064)))), i)
						nox_xxx_netMarkMinimapObject_417190(int32(uint8(*(*byte)(unsafe.Add(unsafe.Pointer(v2), 2064)))), i, 1)
						continue
					}
					if (int32(*(*uint8)(unsafe.Pointer(uintptr(i + 12)))) & 128) != 0 {
						nox_xxx_netMonitorCreature_4D9250(int32(uint8(*(*byte)(unsafe.Add(unsafe.Pointer(v2), 2064)))), i)
						nox_xxx_netMarkMinimapObject_417190(int32(uint8(*(*byte)(unsafe.Add(unsafe.Pointer(v2), 2064)))), i, 1)
						continue
					}
				}
			}
			nox_xxx_gameSetSwitchSolo_4DB220(v3)
			dword_5d4594_1563096 = uint32(v3)
			nox_ticks_reset_416D40()
		}
	}
}
func sub_4DCC00() int32 {
	return int32(dword_5d4594_1563064)
}
func nox_xxx_mapLoadRequired_4DCC80() int32 {
	return int32(*memmap.PtrUint32(6112660, 1563072))
}
func nox_xxx_netUseMap_4DEE00(a1 *byte, a2 int32) *uint32 {
	var (
		v2     int32
		result *uint32
		i      *uint32
		v5     int32
		v6     *uint8
		v7     [41]byte
	)
	v7[0] = 43
	v2 = int32(nox_frame_xxx_2598000)
	libc.StrCpy(&v7[1], a1)
	*(*uint32)(unsafe.Pointer(&v7[37])) = uint32(v2)
	*(*uint32)(unsafe.Pointer(&v7[33])) = uint32(a2)
	result = (*uint32)(unsafe.Pointer(nox_xxx_getFirstPlayerUnit_4DA7C0()))
	for i = result; result != nil; i = result {
		v5 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(i), unsafe.Sizeof(uint32(0))*187)))
		nox_netlist_addToMsgListCli_40EBC0(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v5 + 276))) + 2064)))), 1, (*uint8)(unsafe.Pointer(&v7[0])), 41)
		nox_xxx_netPlayerObjSend_518C30(int32(uintptr(unsafe.Pointer(i))), i, 0, 0)
		if !noxflags.HasGame(noxflags.GameClient) || int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v5 + 276))) + 2064)))) != 31 {
			v6 = nox_netlist_copyPacketList_40ED60(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v5 + 276))) + 2064)))), 1, (*uint32)(unsafe.Pointer(&a1)))
			if v6 != nil {
				nox_xxx_netSendSock_552640(uint32(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v5 + 276))) + 2064))))+1), (*byte)(unsafe.Pointer(v6)), int32(uintptr(unsafe.Pointer(a1))), int8(NOX_NET_SEND_NO_LOCK|NOX_NET_SEND_FLAG2))
			}
		}
		result = (*uint32)(unsafe.Pointer(nox_xxx_getNextPlayerUnit_4DA7F0((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(i)))))))))
	}
	return result
}
func sub_4E4170() {
	var v0 float32
	if (nox_frame_xxx_2598000 % (nox_gameFPS * 5)) == 0 {
		v0 = float32(sub_4E3CA0())
		sub_4E3D50()
		if sub_4E3CA0() != float64(v0) {
			sub_4E3DD0()
		}
	}
}
func nox_xxx_unitsUpdateDeletedList_4E5E20() {
	var (
		result *uint32
		v1     *uint32
		v2     *uint32
	)
	result = *(**uint32)(memmap.PtrOff(6112660, 1565588))
	v1 = nil
	if *memmap.PtrUint32(6112660, 1565588) != 0 {
		for {
			v2 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*113)))))
			if *(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*114)) == nox_frame_xxx_2598000 {
				*(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*113)) = uint32(uintptr(unsafe.Pointer(v1)))
				v1 = result
				nox_xxx_unitRemoveFromUpdatable_4DA920(result)
			} else {
				nox_xxx_unitDeleteFinish_4E5E80(result)
			}
			result = v2
			if v2 == nil {
				break
			}
		}
		*memmap.PtrUint32(6112660, 1565588) = uint32(uintptr(unsafe.Pointer(v1)))
	} else {
		*memmap.PtrUint32(6112660, 1565588) = 0
	}
}
func sub_4E76C0() int32 {
	var v0 int32
	_ = v0
	var result int32
	var i int32
	v0 = 0
	result = int32(uintptr(unsafe.Pointer(noxServer.firstServerObject())))
	for i = result; result != 0; i = result {
		nullsub_25(uint32(i))
		v0 ^= sub_4E7700(i)
		result = int32(uintptr(unsafe.Pointer((*nox_object_t)(unsafe.Pointer(uintptr(i))).Next())))
	}
	return result
}
func sub_4EC720() {
	var (
		v1 int32
		v2 int32
		v3 int32
		v4 int32
		v5 *uint32
		v6 int32
		v7 int16
		v8 *uint8
	)
	_ = v8
	var v9 int32
	var v10 int32
	var v11 float64
	var v12 float64
	var v13 int32
	var v14 int32
	var v15 *uint8
	_ = v15
	var v16 int32
	var v17 int32
	var v18 int32
	if dword_5d4594_1568028 == 0 {
		dword_5d4594_1568028 = uint32(nox_xxx_getNameId_4E3AA0(CString("Crown")))
	}
	if noxflags.HasGame(noxflags.GameModeSolo10 | noxflags.GameModeQuest) {
		return
	}
	nox_xxx_respawnAllow_587000_205200 = 0
	for v0 := int32(int32(dword_5d4594_1568024)); v0 != 0; v0 = int32(*(*uint32)(unsafe.Pointer(uintptr(v0 + 52)))) {
		if *(*uint32)(unsafe.Pointer(uintptr(v0 + 24))) == 0 {
			v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(v0 + 4))))
			v2 = 0
			if v1 == 0 {
				*(*uint32)(unsafe.Pointer(uintptr(v0 + 24))) = 1
				*(*uint32)(unsafe.Pointer(uintptr(v0 + 20))) = nox_frame_xxx_2598000 + nox_gameFPS*30
			} else {
				v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 8))))
				if v3&2 != 0 {
					v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 16))))
					if v4&32 != 0 {
						*(*uint32)(unsafe.Pointer(uintptr(v0 + 4))) = 0
						*(*uint32)(unsafe.Pointer(uintptr(v0 + 24))) = 1
						*(*uint32)(unsafe.Pointer(uintptr(v0 + 20))) = nox_frame_xxx_2598000 + nox_gameFPS*30
					} else if (v4 & 0x8000) != 0 {
						*(*uint32)(unsafe.Pointer(uintptr(v0 + 24))) = 1
						*(*uint32)(unsafe.Pointer(uintptr(v0 + 20))) = nox_frame_xxx_2598000 + nox_gameFPS*30
					}
				} else if int32(*(*uint8)(unsafe.Pointer(uintptr(v1 + 16))))&32 != 0 {
					if nox_xxx_getUnitDefDd10_4E3BA0(int32(*(*uint16)(unsafe.Pointer(uintptr(v1 + 4))))) != 0 {
						v2 = 1
					}
					*(*uint32)(unsafe.Pointer(uintptr(v0 + 4))) = 0
					if v2 != 0 {
						*(*uint32)(unsafe.Pointer(uintptr(v0 + 24))) = 1
						*(*uint32)(unsafe.Pointer(uintptr(v0 + 20))) = nox_frame_xxx_2598000 + nox_gameFPS*30
					}
				} else if uint32(v3)&0x3001000 != 0 || uint32(*(*uint16)(unsafe.Pointer(uintptr(v1 + 4)))) == dword_5d4594_1568028 {
					if *(*uint32)(unsafe.Pointer(uintptr(v1 + 492))) != 0 || nox_xxx_getUnitDefDd10_4E3BA0(int32(*(*uint16)(unsafe.Pointer(uintptr(v1 + 4))))) == 0 {
						v18 = int32(*(*uint32)(unsafe.Pointer(uintptr(v0 + 4))))
						if *(*uint32)(unsafe.Pointer(uintptr(v18 + 492))) != 0 && nox_xxx_getUnitDefDd10_4E3BA0(int32(*(*uint16)(unsafe.Pointer(uintptr(v18 + 4))))) != 0 && uint32(*(*uint16)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v0 + 4))) + 4)))) != dword_5d4594_1568028 && sub_409F40(2) != 0 {
							*(*uint32)(unsafe.Pointer(uintptr(v0 + 24))) = 1
							*(*uint32)(unsafe.Pointer(uintptr(v0 + 20))) = nox_frame_xxx_2598000 + nox_gameFPS*30
						}
					} else {
						v10 = int32(*(*uint32)(unsafe.Pointer(uintptr(v0 + 4))))
						if nox_frame_xxx_2598000 > (nox_gameFPS*5 + *(*uint32)(unsafe.Pointer(uintptr(v10 + 128)))) {
							v11 = float64(*(*float32)(unsafe.Pointer(uintptr(v0 + 8))) - *(*float32)(unsafe.Pointer(uintptr(v10 + 56))))
							v12 = float64(*(*float32)(unsafe.Pointer(uintptr(v0 + 12))) - *(*float32)(unsafe.Pointer(uintptr(v10 + 60))))
							if v12*v12+v11*v11 > 2500.0 {
								nox_xxx_netSendPointFx_522FF0(-127, (*float2)(unsafe.Pointer(uintptr(v10+56))))
								nox_xxx_audCreate_501A30(283, (*float2)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v0 + 4)))+56))), 0, 0)
								nox_xxx_unitMove_4E7010((*nox_object_t)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v0 + 4)))))), (*float2)(unsafe.Pointer(uintptr(v0+8))))
								v13 = int32(*(*uint32)(unsafe.Pointer(uintptr(v0 + 4))))
								v14 = int32(*(*uint32)(unsafe.Pointer(uintptr(v13 + 8))))
								if v14&4096 != 0 {
									nox_xxx_rechargeItem_53C520(v13, 100)
								} else if uint32(v14)&0x1000000 != 0 && nox_xxx_weaponInventoryEquipFlags_415820((*nox_object_t)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v0 + 4)))))))&130 != 0 {
									v15 = *(**uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v0 + 4))) + 736)))
									*(*uint8)(unsafe.Add(unsafe.Pointer(v15), 1)) = *(*uint8)(unsafe.Pointer(uintptr(v0 + 48)))
									*v15 = *(*uint8)(unsafe.Pointer(uintptr(v0 + 49)))
								}
								v16 = int32(*(*uint32)(unsafe.Pointer(uintptr(v0 + 4))))
								v17 = int32(*(*uint32)(unsafe.Pointer(uintptr(v16 + 556))))
								if v17 != 0 {
									nox_xxx_unitSetHP_4E4560((*nox_object_t)(unsafe.Pointer(uintptr(v16))), *(*uint16)(unsafe.Pointer(uintptr(v17 + 4))))
								}
								nox_xxx_netSendPointFx_522FF0(-127, (*float2)(unsafe.Pointer(uintptr(v0+8))))
								nox_xxx_audCreate_501A30(283, (*float2)(unsafe.Pointer(uintptr(v0+8))), 0, 0)
							}
						}
					}
				} else if *(*uint32)(unsafe.Pointer(uintptr(v1 + 492))) != 0 {
					*(*uint32)(unsafe.Pointer(uintptr(v0 + 24))) = 1
					*(*uint32)(unsafe.Pointer(uintptr(v0 + 20))) = nox_frame_xxx_2598000 + nox_gameFPS*30
				}
			}
			if *(*uint32)(unsafe.Pointer(uintptr(v0 + 24))) == 0 {
				continue
			}
		}
		if !(nox_frame_xxx_2598000 >= uint32(*(*int32)(unsafe.Pointer(uintptr(v0 + 20)))) && nox_xxx_getUnitDefDd10_4E3BA0(int32(*(*uint32)(unsafe.Pointer(uintptr(v0))))) != 0) {
			continue
		}
		v5 = (*uint32)(unsafe.Pointer(nox_xxx_newObjectWithTypeInd_4E3450(int32(*(*uint32)(unsafe.Pointer(uintptr(v0)))))))
		v6 = int32(uintptr(unsafe.Pointer(v5)))
		if v5 != nil {
			nox_xxx_createAt_4DAA50((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v5)))))), nil, *(*float32)(unsafe.Pointer(uintptr(v0 + 8))), *(*float32)(unsafe.Pointer(uintptr(v0 + 12))))
			nox_xxx_netSendPointFx_522FF0(-127, (*float2)(unsafe.Pointer(uintptr(v0+8))))
			v7 = int16(*(*uint16)(unsafe.Pointer(uintptr(v0 + 16))))
			*(*uint16)(unsafe.Pointer(uintptr(v6 + 124))) = uint16(v7)
			*(*uint16)(unsafe.Pointer(uintptr(v6 + 126))) = uint16(v7)
			if *(*uint32)(unsafe.Pointer(uintptr(v6 + 8)))&0x13001000 != 0 {
				nox_xxx_modifSetItemAttrs_4E4990((*nox_object_t)(unsafe.Pointer(uintptr(v6))), (*int32)(unsafe.Pointer(uintptr(v0+28))))
			}
			if *(*uint32)(unsafe.Pointer(uintptr(v6 + 8)))&0x1000000 != 0 && nox_xxx_weaponInventoryEquipFlags_415820((*nox_object_t)(unsafe.Pointer(uintptr(v6))))&130 != 0 {
				v8 = *(**uint8)(unsafe.Pointer(uintptr(v6 + 736)))
				*(*uint8)(unsafe.Add(unsafe.Pointer(v8), 1)) = *(*uint8)(unsafe.Pointer(uintptr(v0 + 48)))
				*v8 = *(*uint8)(unsafe.Pointer(uintptr(v0 + 49)))
			}
			nox_xxx_aud_501960(283, (*nox_object_t)(unsafe.Pointer(uintptr(v6))), 0, 0)
		}
		v9 = int32(*(*uint32)(unsafe.Pointer(uintptr(v0 + 4))))
		if v9 != 0 {
			if int32(*(*uint8)(unsafe.Pointer(uintptr(v9 + 8))))&2 != 0 {
				nox_xxx_delayedDeleteObject_4E5CC0((*nox_object_t)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v0 + 4)))))))
			}
		}
		*(*uint32)(unsafe.Pointer(uintptr(v0 + 24))) = 0
		*(*uint32)(unsafe.Pointer(uintptr(v0 + 4))) = uint32(v6)
	}
}
func sub_4EDD70() int32 {
	var (
		result int32
		i      int32
		v2     *uint32
		v3     *uint32
		a3     float2
	)
	result = int32(uintptr(unsafe.Pointer(nox_xxx_getFirstPlayerUnit_4DA7C0())))
	for i = result; result != 0; i = result {
		v2 = *(**uint32)(unsafe.Pointer(uintptr(i + 504)))
		if v2 != nil {
			for {
				v3 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*124)))))
				if *(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*2))&0x10000000 != 0 {
					sub_4ED970(50.0, (*float2)(unsafe.Pointer(uintptr(i+56))), &a3)
					nox_xxx_drop_4ED790((*nox_object_t)(unsafe.Pointer(uintptr(i))), (*nox_object_t)(unsafe.Pointer(v2)), &a3)
				}
				v2 = v3
				if v3 == nil {
					break
				}
			}
		}
		result = int32(uintptr(unsafe.Pointer(nox_xxx_getNextPlayerUnit_4DA7F0((*nox_object_t)(unsafe.Pointer(uintptr(i)))))))
	}
	return result
}
func sub_4EF660(a1p *nox_object_t) int32 {
	var (
		a1 int32 = int32(uintptr(unsafe.Pointer(a1p)))
		v1 int32
		i  int32
	)
	v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
	*(*uint32)(unsafe.Pointer(uintptr(v1 + 116))) = 0
	*(*uint32)(unsafe.Pointer(uintptr(v1 + 120))) = 0
	*(*uint32)(unsafe.Pointer(uintptr(v1 + 124))) = 0
	*(*uint32)(unsafe.Pointer(uintptr(v1 + 128))) = 0
	*(*uint32)(unsafe.Pointer(uintptr(v1 + 308))) = 0
	for i = 4796; i < 4816; *(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v1 + 276))) + uint32(i) - 4))) = 0 {
		i += 4
	}
	if !noxflags.HasGame(noxflags.GameModeCoop) {
		*(*uint8)(unsafe.Pointer(uintptr(v1 + 244))) = uint8(int8(sub_4EF6F0(a1)))
	}
	*(*uint32)(unsafe.Pointer(uintptr(v1 + 264))) = 0
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 520))) = 0
	return sub_422140(int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 276)))))
}
func sub_4F1F20() {
	var (
		v0  int32
		v1  int32
		v2  int32
		v3  int32
		v4  int32
		v5  *uint32
		v6  *uint32
		v7  *uint32
		v8  *uint32
		v9  int32
		v10 int32
		a3  float2
	)
	v9 = nox_game_getQuestStage_4E3CC0()
	if dword_5d4594_1568300 == 0 {
		dword_5d4594_1568300 = uint32(nox_xxx_getNameId_4E3AA0(CString("RewardMarker")))
		*memmap.PtrUint32(6112660, 1568304) = uint32(nox_xxx_getNameId_4E3AA0(CString("RewardMarkerPlus")))
	}
	sub_4F2110()
	sub_4F2210()
	v0 = int32(uintptr(unsafe.Pointer(noxServer.firstServerObject())))
	if v0 != 0 {
		for {
			v10 = int32(uintptr(unsafe.Pointer((*nox_object_t)(unsafe.Pointer(uintptr(v0))).Next())))
			v1 = int32(*(*uint16)(unsafe.Pointer(uintptr(v0 + 4))))
			if uint32(uint16(int16(v1))) == dword_5d4594_1568300 || uint32(v1) == *memmap.PtrUint32(6112660, 1568304) {
				if int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v0 + 692))) + 216))))&128 != 0 {
					v7 = nox_server_rewardgen_activateMarker_4F0720(v0, uint32(v9))
					if v7 != nil {
						nox_xxx_createAt_4DAA50((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v7)))))), nil, *(*float32)(unsafe.Pointer(uintptr(v0 + 56))), *(*float32)(unsafe.Pointer(uintptr(v0 + 60))))
						if *(*uint32)(unsafe.Add(unsafe.Pointer(v7), unsafe.Sizeof(uint32(0))*2))&0x1000000 != 0 {
							if *(*uint32)(unsafe.Add(unsafe.Pointer(v7), unsafe.Sizeof(uint32(0))*3))&12 != 0 {
								v8 = (*uint32)(unsafe.Pointer(nox_xxx_newObjectByTypeID_4E3810(CString("Quiver"))))
								if v8 != nil {
									a3 = *((*float2)(unsafe.Add(unsafe.Pointer((*float2)(unsafe.Pointer(v7))), unsafe.Sizeof(float2{})*7)))
									sub_4ED970(30.0, (*float2)(unsafe.Add(unsafe.Pointer((*float2)(unsafe.Pointer(v7))), unsafe.Sizeof(float2{})*7)), &a3)
									nox_xxx_createAt_4DAA50((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v8)))))), nil, a3.field_0, a3.field_4)
								}
							}
						}
					}
				}
				nox_xxx_delayedDeleteObject_4E5CC0((*nox_object_t)(unsafe.Pointer(uintptr(v0))))
			} else if cgoFuncAddr(cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(v0 + 688))), (*func(int32) *int32)(nil))) == cgoFuncAddr(nox_xxx_initChest_4F0400) {
				v2 = v0.FirstItem()
				if v2 != 0 {
					for {
						v3 = nox_xxx_inventoryGetNext_4E7990(v2)
						v4 = int32(*(*uint16)(unsafe.Pointer(uintptr(v2 + 4))))
						if uint32(uint16(int16(v4))) == dword_5d4594_1568300 || uint32(v4) == *memmap.PtrUint32(6112660, 1568304) {
							v5 = nox_server_rewardgen_activateMarker_4F0720(v2, uint32(v9+1))
							sub_4ED0C0(v0, (*nox_object_t)(unsafe.Pointer(uintptr(v2))))
							nox_xxx_delayedDeleteObject_4E5CC0((*nox_object_t)(unsafe.Pointer(uintptr(v2))))
							if v5 != nil {
								nox_xxx_inventoryPutImpl_4F3070(v0, (*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v5)))))), 0)
								if *(*uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(uint32(0))*2))&0x1000000 != 0 {
									if *(*uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(uint32(0))*3))&12 != 0 {
										v6 = (*uint32)(unsafe.Pointer(nox_xxx_newObjectByTypeID_4E3810(CString("Quiver"))))
										if v6 != nil {
											nox_xxx_inventoryPutImpl_4F3070(v0, (*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v6)))))), 0)
										}
									}
								}
							}
						}
						v2 = v3
						if v3 == 0 {
							break
						}
					}
				}
			}
			v0 = v10
			if v10 == 0 {
				break
			}
		}
	}
}
func nox_xxx_abilUpdateMB_4FBEE0() {
	for p := (*byte)((*byte)(unsafe.Pointer(nox_common_playerInfoGetFirst_416EA0()))); p != nil; p = (*byte)(unsafe.Pointer(nox_common_playerInfoGetNext_416EE0((*nox_playerInfo)(unsafe.Pointer(p))))) {
		if *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(p))), unsafe.Sizeof(uint32(0))*514))) != 0 && *(*byte)(unsafe.Add(unsafe.Pointer(p), 2251)) == 0 {
			for i := int32(0); i < 6; i++ {
				var (
					v2 int32 = i + int32(uint8(*(*byte)(unsafe.Add(unsafe.Pointer(p), 2064))))*6
					v3 int32 = int32(*memmap.PtrUint32(6112660, uint32(v2*4)+0x17F06C))
				)
				if v3 != 0 {
					*memmap.PtrUint32(6112660, uint32(v2*4)+0x17F06C) = uint32(v3 - 1)
					if *memmap.PtrUint32(6112660, uint32((i+int32(uint8(*(*byte)(unsafe.Add(unsafe.Pointer(p), 2064))))*6)*4)+0x17F06C) == 0 {
						nox_xxx_netAbilRepotState_4D8100(int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(p))), unsafe.Sizeof(uint32(0))*514)))), int8(i), 1)
					}
				}
			}
		}
	}
	var next *int32
	for p := (*int32)(*(**int32)(memmap.PtrOff(6112660, 1569648))); p != nil; p = next {
		next = (*int32)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(p), unsafe.Sizeof(int32(0))*4)))))
		if (*(*uint32)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(p), unsafe.Sizeof(int32(0))*1)) + 16))) & 32800) == 0 {
			if nox_frame_xxx_2598000 <= uint32(*(*int32)(unsafe.Add(unsafe.Pointer(p), unsafe.Sizeof(int32(0))*2))) {
				continue
			}
			var v11 int32 = *(*int32)(unsafe.Add(unsafe.Pointer(p), unsafe.Sizeof(int32(0))*1))
			var v8 int32 = sub_425230(*p, 2)
			nox_xxx_aud_501960(v8, (*nox_object_t)(unsafe.Pointer(uintptr(v11))), 0, 0)
			sub_4FC3C0(*(*int32)(unsafe.Add(unsafe.Pointer(p), unsafe.Sizeof(int32(0))*1)), int8(*p), 0)
			if *p == 1 {
				nox_xxx_playerSetState_4FA020((*nox_object_t)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(p), unsafe.Sizeof(int32(0))*1))))), 13)
			}
		}
		var v5 int32 = *(*int32)(unsafe.Add(unsafe.Pointer(p), unsafe.Sizeof(int32(0))*4))
		if v5 != 0 {
			*(*uint32)(unsafe.Pointer(uintptr(v5 + 20))) = uint32(*(*int32)(unsafe.Add(unsafe.Pointer(p), unsafe.Sizeof(int32(0))*5)))
		}
		var v7 int32 = *(*int32)(unsafe.Add(unsafe.Pointer(p), unsafe.Sizeof(int32(0))*5))
		if v7 != 0 {
			*(*uint32)(unsafe.Pointer(uintptr(v7 + 16))) = uint32(*(*int32)(unsafe.Add(unsafe.Pointer(p), unsafe.Sizeof(int32(0))*4)))
		} else {
			*memmap.PtrUint32(6112660, 1569648) = uint32(*(*int32)(unsafe.Add(unsafe.Pointer(p), unsafe.Sizeof(int32(0))*4)))
		}
		nox_alloc_class_free_obj_first((*nox_alloc_class)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&nox_alloc_execAbil_1569644)))), unsafe.Pointer(p))
	}
}
func sub_4FC680() {
	var v0 int32
	if noxflags.HasGame(noxflags.GameModeCoop) && !noxflags.HasGame(noxflags.GameFlag20) && dword_5d4594_1569660 != 0 {
		v0 = int32(uintptr(unsafe.Pointer(nox_xxx_getFirstPlayerUnit_4DA7C0())))
		if v0 != 0 {
			nox_xxx_playerExecuteAbil_4FBB70(v0, *(*int32)(unsafe.Pointer(&dword_5d4594_1569660)))
			dword_5d4594_1569660 = 0
		}
	}
}
func nox_server_xxxInitPlayerUnits_4FC6D0() {
	var (
		v1       *byte
		i        int32
		v3       int32
		v4       int32
		v5       int32
		j        int32
		v7       int32
		FileName [1024]byte
	)
	if nox_xxx_resetMapInit_1569652 != 1 && dword_5d4594_1569656 != 1 {
		return
	}
	if nox_xxx_getFirstPlayerUnit_4DA7C0() == nil {
		return
	}
	if noxflags.HasGame(noxflags.GameModeQuest) {
		if nox_game_getQuestStage_4E3CC0() == 1 {
			nox_game_sendQuestStage_4D6960(math.MaxUint8)
			sub_4D7440(1)
			sub_4D60B0()
		} else if sub_4D6F30() == 0 || sub_4D7430() != 0 {
			if sub_4D76F0() == 1 {
				sub_4D6880(math.MaxUint8, 1)
				sub_4D76E0(0)
				sub_4D7440(1)
				sub_4D60B0()
			} else {
				v1 = nox_fs_root()
				nox_sprintf(&FileName[0], CString("%s\\Save\\_temp_.dat"), v1)
				for i = int32(uintptr(unsafe.Pointer(nox_xxx_getFirstPlayerUnit_4DA7C0()))); i != 0; i = int32(uintptr(unsafe.Pointer(nox_xxx_getNextPlayerUnit_4DA7F0((*nox_object_t)(unsafe.Pointer(uintptr(i))))))) {
					v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(i + 748))))
					v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(v3 + 276))))
					if *(*uint32)(unsafe.Pointer(uintptr(v4 + 4792))) == 1 && *(*uint32)(unsafe.Pointer(uintptr(v3 + 552))) == 0 && nox_xxx_playerSaveToFile_41A140(&FileName[0], int32(*(*uint8)(unsafe.Pointer(uintptr(v4 + 2064))))) != 0 {
						v5 = sub_419EE0(int8(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v3 + 276))) + 2064)))))
						nox_xxx_sendGauntlet_4DCF80(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v3 + 276))) + 2064)))), 1)
						if sub_41CFA0(&FileName[0], int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v3 + 276))) + 2064))))) == 0 && v5 == 0 {
							nox_xxx_sendGauntlet_4DCF80(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v3 + 276))) + 2064)))), 0)
						}
						nox_fs_remove(&FileName[0])
					}
					sub_4D6770(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v3 + 276))) + 2064)))))
				}
				sub_4D6880(math.MaxUint8, 0)
				sub_4D7440(1)
				sub_4D60B0()
			}
		} else {
			nox_game_sendQuestStage_4D6960(math.MaxUint8)
			sub_4D7440(1)
			sub_4D60B0()
		}
	} else {
		nox_xxx_netMsgFadeBegin_4D9800(1, 1)
	}
	if noxflags.HasGame(noxflags.GameOnline) {
		if !noxflags.HasGame(noxflags.GameModeChat) {
			for j = int32(uintptr(unsafe.Pointer(nox_xxx_getFirstPlayerUnit_4DA7C0()))); j != 0; j = int32(uintptr(unsafe.Pointer(nox_xxx_getNextPlayerUnit_4DA7F0((*nox_object_t)(unsafe.Pointer(uintptr(j))))))) {
				v7 = int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(j + 748))) + 276))))
				if int32(*(*uint8)(unsafe.Pointer(uintptr(v7 + 2064)))) != 31 && (int32(*(*uint8)(unsafe.Pointer(uintptr(v7 + 3680))))&1) == 0 {
					nox_xxx_buffApplyTo_4FF380((*nox_object_t)(unsafe.Pointer(uintptr(j))), 23, 0, 5)
				}
			}
		}
	}
}
func nox_xxx_playerSomeWallsUpdate_5003B0(obj *nox_object_t) int32 {
	var (
		a1     int32 = int32(uintptr(unsafe.Pointer(obj)))
		v1     int32
		result int32
		v3     int32
		v4     *uint8
		v5     *uint8
	)
	if *(*int32)(unsafe.Pointer(&dword_5d4594_1569756)) <= 0 {
		v1 = a1
	} else {
		v1 = nox_xxx_spellCastedFirst_4FE930()
		if v1 == 0 {
			return 0
		}
		for *(*uint32)(unsafe.Pointer(uintptr(v1 + 4))) != 132 || *(*uint32)(unsafe.Pointer(uintptr(v1 + 16))) != uint32(a1) || int32(*(*uint8)(unsafe.Pointer(uintptr(v1 + 88))))&1 != 0 {
			result = nox_xxx_spellCastedNext_4FE940(v1)
			v1 = result
			if result == 0 {
				return result
			}
		}
		if v1 == 0 {
			return 0
		}
		*(*uint32)(unsafe.Pointer(uintptr(v1 + 92))) = uint32(cgoFuncAddr(nox_xxx_spellWallCreate_4FFA90))
		*(*uint32)(unsafe.Pointer(uintptr(v1 + 96))) = uint32(cgoFuncAddr(nox_xxx_spellWallUpdate_500070))
		*(*uint32)(unsafe.Pointer(uintptr(v1 + 100))) = uint32(cgoFuncAddr(nox_xxx_spellWallDestroy_500080))
		*(*uint32)(unsafe.Pointer(uintptr(v1 + 48))) = 0
	}
	v3 = 0
	if dword_5d4594_1569756 > 0 {
		v4 = (*uint8)(memmap.PtrOff(6112660, 1569764))
		for {
			v5 = (*uint8)(nox_server_getWallAtGrid_410580(int32(*v4), int32(*(*uint8)(unsafe.Add(unsafe.Pointer(v4), 1)))))
			if v5 != nil {
				*v5 = *(*uint8)(unsafe.Add(unsafe.Pointer(v4), 13))
			} else {
				v5 = (*uint8)(nox_xxx_wallCreateAt_410250(int32(*v4), int32(*(*uint8)(unsafe.Add(unsafe.Pointer(v4), 1)))))
				if v5 == nil {
					return 0
				}
				*(*uint8)(unsafe.Add(unsafe.Pointer(v5), 4)) |= 8
				*(*uint8)(unsafe.Add(unsafe.Pointer(v5), 1)) = *(*uint8)(unsafe.Add(unsafe.Pointer(v4), 11))
				*(*uint8)(unsafe.Add(unsafe.Pointer(v5), 2)) = *(*uint8)(unsafe.Add(unsafe.Pointer(v4), 12))
				*v5 = *(*uint8)(unsafe.Add(unsafe.Pointer(v4), 13))
				*(*uint8)(unsafe.Add(unsafe.Pointer(v5), 7)) = *(*uint8)(unsafe.Add(unsafe.Pointer(v4), 14))
			}
			nox_xxx_netWallCreate_4FFE80(v1, v5, int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v4))), unsafe.Sizeof(uint32(0))*1)))), int8(*(*uint8)(unsafe.Add(unsafe.Pointer(v4), 8))), int8(*(*uint8)(unsafe.Add(unsafe.Pointer(v4), 9))), int8(*(*uint8)(unsafe.Add(unsafe.Pointer(v4), 10))))
			v3++
			v4 = (*uint8)(unsafe.Add(unsafe.Pointer(v4), 16))
			if v3 >= *(*int32)(unsafe.Pointer(&dword_5d4594_1569756)) {
				break
			}
		}
	}
	dword_5d4594_1569756 = 0
	return 1
}
func sub_500510(a1 *byte) {
	if a1 != nil {
		libc.StrCpy((*byte)(memmap.PtrOff(6112660, 1570008)), a1)
	}
}
func sub_502100() {
	nox_alloc_class_free_all((*nox_alloc_class)(unsafe.Pointer(*(**uint32)(memmap.PtrOff(6112660, 1599056)))))
	dword_5d4594_1599060 = 0
}
func nox_xxx_voteUptate_506F30() {
	var (
		result *uint32
		v1     *uint32
	)
	result = *(**uint32)(unsafe.Pointer(&dword_5d4594_1599656))
	if dword_5d4594_1599656 != 0 {
		for {
			v1 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*11)))))
			switch *result {
			case 0:
				fallthrough
			case 1:
				sub_506F80(int32(uintptr(unsafe.Pointer(result))))
			case 2:
				sub_507090(int32(uintptr(unsafe.Pointer(result))))
			case 3:
				sub_507100(int32(uintptr(unsafe.Pointer(result))))
			default:
			}
			result = v1
			if v1 == nil {
				break
			}
		}
	}
}
func sub_50AFA0() {
	var (
		i      int32
		v1     int32
		v2     int32
		v3     int32
		v4     int32
		v5     int32
		v6     int32
		v7     int32
		v8     int32
		v9     int32
		v10    int32
		v11    float64
		v12    int32
		v13    int32
		v14    int32
		v15    int32
		v16    *uint8
		v17    *float32
		v18    int16
		result *float32
		k      *float32
		v21    int32
		v22    int32
		v23    float32
		v24    float32
		v25    float32
		v26    float32
		v27    float32
		v28    float32
		v29    float32
		v30    float32
		v31    float32
		v32    float32
		v33    float32
		v34    float32
		v35    float32
		v36    float32
		j      int32
		v38    int32
		v39    int32
		v40    int32
		v41    float32
		v42    float32
		a2     float2
		v44    int32
	)
	dword_5d4594_2386160 = 0
	*(*[65536]nox_server_xxx)(unsafe.Pointer(&nox_server_xxx_1599716[0])) = [65536]nox_server_xxx{}
	dword_5d4594_2386164 = 0
	for i = int32(uintptr(unsafe.Pointer(noxServer.firstServerObject()))); i != 0; i = int32(uintptr(unsafe.Pointer((*nox_object_t)(unsafe.Pointer(uintptr(i))).Next()))) {
		nox_xxx_aiTestUnitDangerous_50B2C0(i)
		v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(i + 8))))
		if (v1 & 128) == 0 {
			if v1&2048 != 0 {
				v23 = float32(float64(*(*float32)(unsafe.Pointer(uintptr(i + 56)))) * 0.043478262)
				v2 = int(v23)
				v24 = float32(float64(*(*float32)(unsafe.Pointer(uintptr(i + 60)))) * 0.043478262)
				v3 = int(v24)
				nox_server_xxx_1599716[v3+(v2<<8)].field_8 |= 16
			} else if v1&1024 != 0 {
				v25 = float32(float64(*(*float32)(unsafe.Pointer(uintptr(i + 56)))) * 0.043478262)
				v4 = int(v25)
				v26 = float32(float64(*(*float32)(unsafe.Pointer(uintptr(i + 60)))) * 0.043478262)
				v5 = int(v26)
				nox_server_xxx_1599716[v5+(v4<<8)].field_8 |= 32
			} else if v1&0x4000 != 0 {
				v27 = float32(float64(*(*float32)(unsafe.Pointer(uintptr(i + 56)))) * 0.043478262)
				v6 = int(v27)
				v28 = float32(float64(*(*float32)(unsafe.Pointer(uintptr(i + 60)))) * 0.043478262)
				v7 = int(v28)
				nox_server_xxx_1599716[v7+(v6<<8)].field_8 |= 4
			} else if (v1 & 0x8000) == 0 {
				if (int32(*(*uint8)(unsafe.Pointer(uintptr(i + 16)))) & 73) == 0 {
					if uint32(v1)&0x400000 != 0 {
						v31 = float32(float64(*(*float32)(unsafe.Pointer(uintptr(i + 232)))) * 0.043478262)
						v10 = int(v31)
						v11 = float64(*(*float32)(unsafe.Pointer(uintptr(i + 236)))) * 0.043478262
						v12 = v10
						v44 = v10
						v32 = float32(v11)
						v13 = int(v32)
						v33 = float32(float64(*(*float32)(unsafe.Pointer(uintptr(i + 240)))) * 0.043478262)
						v39 = int(v33)
						v34 = float32(float64(*(*float32)(unsafe.Pointer(uintptr(i + 244)))) * 0.043478262)
						v14 = int(v34)
						v15 = v13
						v40 = v14
						for j = v13; v15 <= v14; j = v15 {
							v38 = v12
							if v12 <= v39 {
								v16 = (*uint8)(unsafe.Pointer(&nox_server_xxx_1599716[v15+(v12<<8)].field_8))
								for {
									v17 = mem_getFloatPtr(0x587000, 0x3927C)
									for {
										v41 = float32(float64(v38) * 23.0)
										a2.field_0 = v41 + *((*float32)(unsafe.Add(unsafe.Pointer(v17), -int(unsafe.Sizeof(float32(0))*1))))
										v42 = float32(float64(j) * 23.0)
										a2.field_4 = v42 + *v17
										if sub_547DB0(i, &a2) != 0 {
											break
										}
										v17 = (*float32)(unsafe.Add(unsafe.Pointer(v17), unsafe.Sizeof(float32(0))*2))
										if int32(uintptr(unsafe.Pointer(v17))) >= int32(uintptr(memmap.PtrOff(0x587000, 234180))) {
											goto LABEL_22
										}
									}
									v18 = int16(*(*uint16)(unsafe.Pointer(v16)))
									*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v18))), 0)) = uint8(int8(int32(*(*uint16)(unsafe.Pointer(v16))) | 1))
									*(*uint16)(unsafe.Pointer(v16)) = uint16(v18)
									if (int32(*(*uint8)(unsafe.Pointer(uintptr(i + 16)))) & 16) == 0 {
										*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v18))), 0)) = uint8(int8(int32(v18) | 2))
										*(*uint16)(unsafe.Pointer(v16)) = uint16(v18)
									}
								LABEL_22:
									v12++
									v16 = (*uint8)(unsafe.Add(unsafe.Pointer(v16), NOX_SERVER_XXX_SIZE*unsafe.Sizeof(nox_server_xxx{})))
									v38 = v12
									if v12 > v39 {
										break
									}
								}
								v15 = j
								v14 = v40
								v12 = v44
							}
							v15++
						}
					}
				}
			} else {
				v29 = float32(float64(*(*float32)(unsafe.Pointer(uintptr(i + 56)))) * 0.043478262)
				v8 = int(v29)
				v30 = float32(float64(*(*float32)(unsafe.Pointer(uintptr(i + 60)))) * 0.043478262)
				v9 = int(v30)
				nox_server_xxx_1599716[v9+(v8<<8)].field_8 |= 8
			}
		}
	}
	result = (*float32)(nox_xxx_waypointGetList_579860())
	for k = result; result != nil; k = result {
		if sub_579EE0(int32(uintptr(unsafe.Pointer(k))), 128) != 0 {
			v35 = float32(float64(*(*float32)(unsafe.Add(unsafe.Pointer(k), unsafe.Sizeof(float32(0))*2))) * 0.043478262)
			v21 = int(v35)
			v36 = float32(float64(*(*float32)(unsafe.Add(unsafe.Pointer(k), unsafe.Sizeof(float32(0))*3))) * 0.043478262)
			v22 = int(v36)
			nox_server_xxx_1599716[v22+(v21<<8)].field_8 |= 64
		}
		result = (*float32)(unsafe.Pointer(uintptr(nox_xxx_waypointNext_579870(int32(uintptr(unsafe.Pointer(k)))))))
	}
}
func sub_50D890() uint32 {
	var result uint32
	if (nox_frame_xxx_2598000 % (nox_gameFPS * 5)) == 0 {
		sub_50D8D0()
	}
	result = nox_frame_xxx_2598000 / 15
	if (nox_frame_xxx_2598000 % 15) == 0 {
		result = uint32(sub_50D960())
	}
	return result
}
func nox_xxx_addDebugEntry_511590(a1p unsafe.Pointer, a2p unsafe.Pointer) {
	var (
		a1 int32 = int32(uintptr(a1p))
		a2 int32 = int32(uintptr(a2p))
		v2 int32
		v3 *uint32
		v4 int32
		v5 float64
		v6 float64
		v7 int8
		v8 float64
		v9 float32
	)
	v2 = a1
	if a1 != 0 {
		if a2 != 0 {
			if (int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 16))))&32) == 0 && (int32(*(*uint8)(unsafe.Pointer(uintptr(a2 + 16))))&32) == 0 {
				v3 = (*uint32)(nox_alloc_class_new_obj_zero((*nox_alloc_class)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&nox_alloc_springs_2386568))))))
				v4 = int32(uintptr(unsafe.Pointer(v3)))
				if v3 != nil {
					*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*2)) = uint32(a1)
					*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*3)) = uint32(a2)
					*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*4)) = 0x42700000
					v9 = float32(nox_xxx_objectGetMass_4E4A70(a1))
					if nox_xxx_objectGetMass_4E4A70(a2) >= float64(v9) {
						nox_xxx_objectGetMass_4E4A70(a2)
					} else {
						nox_xxx_objectGetMass_4E4A70(v2)
					}
					*(*uint32)(unsafe.Pointer(uintptr(v4 + 20))) = 0
					v5 = float64(*(*float32)(unsafe.Pointer(uintptr(a2 + 56))) - *(*float32)(unsafe.Pointer(uintptr(v2 + 56))))
					v6 = float64(*(*float32)(unsafe.Pointer(uintptr(a2 + 60))) - *(*float32)(unsafe.Pointer(uintptr(v2 + 60))))
					v7 = int8(int32(*(*uint8)(unsafe.Pointer(uintptr(v4 + 40)))) | 1)
					*(*uint32)(unsafe.Pointer(uintptr(v4 + 28))) = 0x43800000
					*(*uint8)(unsafe.Pointer(uintptr(v4 + 40))) = uint8(v7)
					v8 = math.Sqrt(v6*v6 + v5*v5)
					*(*float32)(unsafe.Pointer(uintptr(v4 + 32))) = float32(v8)
					*(*float32)(unsafe.Pointer(uintptr(v4 + 36))) = float32(v8)
					*(*float32)(unsafe.Pointer(uintptr(v4 + 24))) = float32(v8)
					sub_511560(v4)
				}
			}
		}
	}
}
func nox_xxx_serverLoopSendMap_519990() {
	var (
		v0 int32
		v1 *uint8
	)
	if dword_5d4594_2388648 != 0 {
		if *memmap.PtrUint32(6112660, 2388644) == 0 {
			nox_xxx_netMapSendPrepair_519EB0_net_mapsend()
		}
		v0 = 0
		v1 = (*uint8)(memmap.PtrOff(6112660, 2387180))
		for {
			if int32(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v1))), -int(unsafe.Sizeof(uint16(0))*15))))) == 1 {
				if sub_519930(v0+1) > int32(*((*uint8)(unsafe.Add(unsafe.Pointer(v1), -4)))) {
					*(*uint32)(unsafe.Pointer(v1))++
				} else {
					nox_xxx_netSendMap_5199F0_net_mapsend((*uint8)(unsafe.Add(unsafe.Pointer(v1), -32)))
				}
			}
			v1 = (*uint8)(unsafe.Add(unsafe.Pointer(v1), 48))
			v0++
			if int32(uintptr(unsafe.Pointer(v1))) >= int32(uintptr(memmap.PtrOff(6112660, 2388668))) {
				break
			}
		}
	}
}
func sub_51A1F0(a1 int32) {
	var (
		v1  int32
		v2  int32
		v3  int32
		v4  int32
		v5  int32
		v6  int32
		v7  float64
		v8  int32
		v9  int32
		v10 int32
		v11 int32
		v12 int32
		v13 int32
		v14 int32
		v15 int32
		i   int32
		v17 int32
		v18 int32
		v19 float32
		v20 float32
		v21 uint32
		v22 int32
		v23 int32
		v24 uint32
		v25 int32
	)
	v21 = 0
	v23 = 0
	v1 = nox_game_getQuestStage_4E3CC0()
	v22 = v1
	v19 = float32(nox_xxx_gamedataGetFloat_419D40(CString("QuestHardcoreStage")))
	v24 = uint32(int(v19))
	if *memmap.PtrUint32(6112660, 2388668) == 0 {
		*memmap.PtrUint32(6112660, 2388668) = uint32(nox_xxx_getNameId_4E3AA0(CString("HecubahMarker")))
		*memmap.PtrUint32(6112660, 2388672) = uint32(nox_xxx_getNameId_4E3AA0(CString("NecromancerMarker")))
	}
	v2 = int32(uintptr(unsafe.Pointer(noxServer.firstServerObject())))
	if v2 != 0 {
		v3 = a1
		for {
			v4 = int32(uintptr(unsafe.Pointer((*nox_object_t)(unsafe.Pointer(uintptr(v2))).Next())))
			v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 8))))
			v25 = v4
			if v5&32 != 0 && int32(*(*uint8)(unsafe.Pointer(uintptr(v2 + 12))))&1 != 0 {
				v21++
			} else if uint32(*(*uint16)(unsafe.Pointer(uintptr(v2 + 4)))) == *memmap.PtrUint32(6112660, 2388668) {
				v23++
			}
			if uint32(v5)&0x20000 != 0 {
				v6 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 748))))
				if *(*uint32)(unsafe.Pointer(uintptr(v3*16 + v6))) != 0 {
					switch *(*uint8)(unsafe.Pointer(uintptr(v6 + v3 + 83))) {
					case 0:
						v7 = nox_xxx_gamedataGetFloat_419D40(CString("GeneratorMaxActiveCreaturesHigh"))
						*(*uint8)(unsafe.Pointer(uintptr(v6 + 87))) = uint8(int8(int64(v7)))
					case 1:
						v7 = nox_xxx_gamedataGetFloat_419D40(CString("GeneratorMaxActiveCreaturesNormal"))
						*(*uint8)(unsafe.Pointer(uintptr(v6 + 87))) = uint8(int8(int64(v7)))
					case 2:
						v7 = nox_xxx_gamedataGetFloat_419D40(CString("GeneratorMaxActiveCreaturesLow"))
						*(*uint8)(unsafe.Pointer(uintptr(v6 + 87))) = uint8(int8(int64(v7)))
					case 3:
						v7 = nox_xxx_gamedataGetFloat_419D40(CString("GeneratorMaxActiveCreaturesSingular"))
						*(*uint8)(unsafe.Pointer(uintptr(v6 + 87))) = uint8(int8(int64(v7)))
					default:
					}
					if uint32(nox_game_getQuestStage_4E3CC0()) >= v24 && int32(*(*uint8)(unsafe.Pointer(uintptr(v6 + v3 + 83)))) != 3 {
						*(*uint8)(unsafe.Pointer(uintptr(v6 + 87))) *= 2
					}
					v8 = sub_51A500(int32(*(*uint32)(unsafe.Pointer(uintptr(v3*16 + v6)))))
					if v8 != 0 {
						*(*uint16)(unsafe.Pointer(uintptr(v2 + 4))) = uint16(int16(v8))
					}
				} else {
					nox_xxx_delayedDeleteObject_4E5CC0((*nox_object_t)(unsafe.Pointer(uintptr(v2))))
				}
			}
			v2 = v25
			if v25 == 0 {
				break
			}
		}
		if v21 > 1 {
			v9 = noxRndCounter1.IntClamp(0, int32(v21-1))
			v10 = 0
			v11 = int32(uintptr(unsafe.Pointer(noxServer.firstServerObject())))
			if v11 != 0 {
				for {
					v12 = int32(uintptr(unsafe.Pointer((*nox_object_t)(unsafe.Pointer(uintptr(v11))).Next())))
					if int32(*(*uint8)(unsafe.Pointer(uintptr(v11 + 8))))&32 != 0 && int32(*(*uint8)(unsafe.Pointer(uintptr(v11 + 12))))&1 != 0 {
						if v10 != v9 {
							nox_xxx_delayedDeleteObject_4E5CC0((*nox_object_t)(unsafe.Pointer(uintptr(v11))))
						}
						v10++
					}
					v11 = v12
					if v12 == 0 {
						break
					}
				}
			}
		}
		v1 = v22
	}
	sub_51A940(0)
	if v1 >= 5 {
		v20 = float32(nox_xxx_gamedataGetFloat_419D40(CString("MinionsAlwaysStage")))
		v13 = int(v20)
		if v1 == 5 || v1 >= v13 || v1&1 != 0 && noxRndCounter1.IntClamp(1, 100) >= 50 {
			sub_51A940(1)
			if v23 != 0 {
				v14 = noxRndCounter1.IntClamp(1, v23)
				v15 = 0
				for i = int32(uintptr(unsafe.Pointer(noxServer.firstServerObject()))); i != 0; i = int32(uintptr(unsafe.Pointer((*nox_object_t)(unsafe.Pointer(uintptr(i))).Next()))) {
					if uint32(*(*uint16)(unsafe.Pointer(uintptr(i + 4)))) == *memmap.PtrUint32(6112660, 2388668) && func() int32 {
						p := &v15
						*p++
						return *p
					}() == v14 {
						nox_xxx_spawnHecubahQuest_51A5A0((*int32)(unsafe.Pointer(uintptr(i + 56))))
					}
					if uint32(*(*uint16)(unsafe.Pointer(uintptr(i + 4)))) == *memmap.PtrUint32(6112660, 2388672) && noxRndCounter1.IntClamp(1, 100) >= 50 {
						nox_xxx_spawnNecroQuest_51A7A0((*int32)(unsafe.Pointer(uintptr(i + 56))))
					}
				}
			}
		}
	}
	v17 = int32(uintptr(unsafe.Pointer(noxServer.firstServerObject())))
	if v17 != 0 {
		for {
			v18 = int32(uintptr(unsafe.Pointer((*nox_object_t)(unsafe.Pointer(uintptr(v17))).Next())))
			if uint32(*(*uint16)(unsafe.Pointer(uintptr(v17 + 4)))) == *memmap.PtrUint32(6112660, 2388668) {
				nox_xxx_delayedDeleteObject_4E5CC0((*nox_object_t)(unsafe.Pointer(uintptr(v17))))
			}
			if uint32(*(*uint16)(unsafe.Pointer(uintptr(v17 + 4)))) == *memmap.PtrUint32(6112660, 2388672) {
				nox_xxx_delayedDeleteObject_4E5CC0((*nox_object_t)(unsafe.Pointer(uintptr(v17))))
			}
			v17 = v18
			if v18 == 0 {
				break
			}
		}
	}
}
func sub_51A920(a1 int32) int32 {
	var result int32
	result = a1
	*memmap.PtrUint32(6112660, 2388660) = uint32(a1)
	return result
}
func nox_xxx_updateUnits_51B100_B() {
	var v2 int32 = int32(uintptr(unsafe.Pointer(nox_xxx_getFirstUpdatableObject_4DA8A0())))
	if v2 != 0 {
		var v3 int32 = 0
		for {
			v3 = int32(uintptr(unsafe.Pointer(nox_xxx_getNextUpdatableObject_4DA8B0((*nox_object_t)(unsafe.Pointer(uintptr(v2)))))))
			if (int32(*(*uint8)(unsafe.Pointer(uintptr(v2 + 16)))) & 32) == 0 {
				var v4 int32 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 520))))
				if v4 != 0 && int32(*(*uint8)(unsafe.Pointer(uintptr(v4 + 16))))&32 != 0 {
					var v5 int32 = int32(*(*uint32)(unsafe.Pointer(uintptr(v4 + 508))))
					if v5 == 0 || int32(*(*uint8)(unsafe.Pointer(uintptr(v5 + 16))))&32 != 0 {
						*(*uint32)(unsafe.Pointer(uintptr(v2 + 520))) = 0
					} else {
						*(*uint32)(unsafe.Pointer(uintptr(v2 + 520))) = uint32(v5)
					}
				}
				var v6 func(int32)
				v6 = cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(v2 + 744))), (*func(int32))(nil))
				if v6 != nil {
					if int32(*(*uint8)(unsafe.Pointer(uintptr(v2 + 8))))&4 != 0 {
						var v7 int32 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 748))))
						if (int32(*(*uint8)(unsafe.Pointer(uintptr(v2 + 16)))) & 2) == 0 {
							v6(v2)
						}
						nox_xxx_playerCmd_51AC30(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v7 + 276))) + 2064)))))
					} else if (int32(*(*uint8)(unsafe.Pointer(uintptr(v2 + 16)))) & 2) == 0 {
						v6(v2)
					}
				}
				if int32(*(*uint8)(unsafe.Pointer(uintptr(v2 + 8))))&2 != 0 {
					nox_xxx_unitNeedSync_4E44F0(v2)
				}
				nox_xxx_updateFallLogic_51B870((*nox_object_t)(unsafe.Pointer(uintptr(v2))))
				var v8 *uint16 = *(**uint16)(unsafe.Pointer(uintptr(v2 + 556)))
				if v8 != nil {
					*(*uint16)(unsafe.Add(unsafe.Pointer(v8), unsafe.Sizeof(uint16(0))*1)) = *v8
				}
				var v9 float64 = float64(*(*float32)(unsafe.Pointer(uintptr(v2 + 64))) - *(*float32)(unsafe.Pointer(uintptr(v2 + 56))))
				*(*uint32)(unsafe.Pointer(uintptr(v2 + 16))) &= 0xFFEFFFFF
				var v58 float32 = float32(v9)
				var v61 float32 = float32(sub_419A10(v58))
				var v59 float32 = *(*float32)(unsafe.Pointer(uintptr(v2 + 68))) - *(*float32)(unsafe.Pointer(uintptr(v2 + 60)))
				var v62 float32 = float32(sub_419A10(v59))
				if float64(*(*float32)(unsafe.Pointer(uintptr(v2 + 88)))) <= *(*float64)(unsafe.Pointer(&qword_581450_10392)) || float64(*(*float32)(unsafe.Pointer(uintptr(v2 + 88)))) >= *(*float64)(unsafe.Pointer(&qword_581450_10256)) || float64(*(*float32)(unsafe.Pointer(uintptr(v2 + 92)))) <= *(*float64)(unsafe.Pointer(&qword_581450_10392)) || float64(*(*float32)(unsafe.Pointer(uintptr(v2 + 92)))) >= *(*float64)(unsafe.Pointer(&qword_581450_10256)) || float64(*(*float32)(unsafe.Pointer(uintptr(v2 + 80)))) <= *(*float64)(unsafe.Pointer(&qword_581450_10392)) || float64(*(*float32)(unsafe.Pointer(uintptr(v2 + 80)))) >= *(*float64)(unsafe.Pointer(&qword_581450_10256)) || float64(*(*float32)(unsafe.Pointer(uintptr(v2 + 84)))) <= *(*float64)(unsafe.Pointer(&qword_581450_10392)) || float64(*(*float32)(unsafe.Pointer(uintptr(v2 + 84)))) >= *(*float64)(unsafe.Pointer(&qword_581450_10256)) || float64(v61) <= *(*float64)(unsafe.Pointer(&qword_581450_10392)) || float64(v61) >= *(*float64)(unsafe.Pointer(&qword_581450_10256)) || float64(v62) <= *(*float64)(unsafe.Pointer(&qword_581450_10392)) || float64(v62) >= *(*float64)(unsafe.Pointer(&qword_581450_10256)) || float64(*(*float32)(unsafe.Pointer(uintptr(v2 + 104)))) <= *(*float64)(unsafe.Pointer(&qword_581450_10392)) || float64(*(*float32)(unsafe.Pointer(uintptr(v2 + 104)))) >= *(*float64)(unsafe.Pointer(&qword_581450_10256)) || float64(*(*float32)(unsafe.Pointer(uintptr(v2 + 108)))) <= *(*float64)(unsafe.Pointer(&qword_581450_10392)) || float64(*(*float32)(unsafe.Pointer(uintptr(v2 + 108)))) >= *(*float64)(unsafe.Pointer(&qword_581450_10256)) {
					nox_xxx_unitHasCollideOrUpdateFn_537610((*nox_object_t)(unsafe.Pointer(uintptr(v2))))
					*(*uint32)(unsafe.Pointer(uintptr(v2 + 16))) &= 0xF7FFFFFF
				} else {
					*(*uint32)(unsafe.Pointer(uintptr(v2 + 88))) = 0
					*(*uint32)(unsafe.Pointer(uintptr(v2 + 92))) = 0
					*(*uint32)(unsafe.Pointer(uintptr(v2 + 80))) = 0
					*(*uint32)(unsafe.Pointer(uintptr(v2 + 84))) = 0
					nox_xxx_unitRaise_4E46F0((*nox_object_t)(unsafe.Pointer(uintptr(v2))), 0.0)
					var v10 int32 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 56))))
					*(*uint32)(unsafe.Pointer(uintptr(v2 + 68))) = *(*uint32)(unsafe.Pointer(uintptr(v2 + 60)))
					var v11 int32 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 16))))
					*(*uint32)(unsafe.Pointer(uintptr(v2 + 64))) = uint32(v10)
					var v12 int32 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 744))))
					*(*uint32)(unsafe.Pointer(uintptr(v2 + 108))) = 0
					*(*uint32)(unsafe.Pointer(uintptr(v2 + 16))) = uint32(v11) | 0x8000000
					if v12 == 0 {
						nox_xxx_unitRemoveFromUpdatable_4DA920((*uint32)(unsafe.Pointer(uintptr(v2))))
					}
				}
			}
			v2 = v3
			if v3 == 0 {
				break
			}
		}
	}
}
func nox_xxx_updateUnits_51B100_C() {
	for k := int32(int32(uintptr(unsafe.Pointer(nox_xxx_getFirstUpdatableObject_4DA8A0())))); k != 0; k = int32(uintptr(unsafe.Pointer(nox_xxx_getNextUpdatableObject_4DA8B0((*nox_object_t)(unsafe.Pointer(uintptr(k))))))) {
		var (
			v20 int32 = int32(*(*uint32)(unsafe.Pointer(uintptr(k + 64))))
			v21 int32 = int32(*(*uint32)(unsafe.Pointer(uintptr(k + 60))))
		)
		*(*uint32)(unsafe.Pointer(uintptr(k + 72))) = *(*uint32)(unsafe.Pointer(uintptr(k + 56)))
		var v22 int32 = int32(*(*uint32)(unsafe.Pointer(uintptr(k + 68))))
		*(*uint32)(unsafe.Pointer(uintptr(k + 56))) = uint32(v20)
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v20))), 0)) = *(*uint8)(unsafe.Pointer(uintptr(k + 541)))
		*(*uint32)(unsafe.Pointer(uintptr(k + 76))) = uint32(v21)
		*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v21))), unsafe.Sizeof(uint16(0))*0)) = *(*uint16)(unsafe.Pointer(uintptr(k + 126)))
		*(*uint32)(unsafe.Pointer(uintptr(k + 60))) = uint32(v22)
		*(*uint32)(unsafe.Pointer(uintptr(k + 88))) = 0
		*(*uint32)(unsafe.Pointer(uintptr(k + 92))) = 0
		*(*uint16)(unsafe.Pointer(uintptr(k + 124))) = uint16(int16(v21))
		if int32(uint8(int8(v20))) > 4 {
			*(*uint8)(unsafe.Pointer(uintptr(k + 541))) = 4
		}
		*(*float32)(unsafe.Pointer(uintptr(k + 544))) = float32(float64(*(*float32)(unsafe.Pointer(uintptr(k + 552)))+*(*float32)(unsafe.Pointer(uintptr(k + 548)))) * (5.0 - float64(*(*uint8)(unsafe.Pointer(uintptr(k + 541))))) * 0.2)
		if nox_xxx_testUnitBuffs_4FF350((*nox_object_t)(unsafe.Pointer(uintptr(k))), 4) != 0 {
			*(*float32)(unsafe.Pointer(uintptr(k + 544))) = float32(float64(*(*float32)(unsafe.Pointer(uintptr(k + 544)))) * 0.5)
		}
		if int32(*(*uint8)(unsafe.Pointer(uintptr(k + 541)))) != 0 || int32(*(*uint8)(unsafe.Pointer(uintptr(k + 540)))) != 0 {
			var v23 int16 = int16(*(*uint16)(unsafe.Pointer(uintptr(k + 542))))
			if int32(v23) > 0 {
				var v24 int16 = int16(int32(v23) - 1)
				*(*uint16)(unsafe.Pointer(uintptr(k + 542))) = uint16(v24)
				if int32(v24) == 0 {
					var v25 int8 = int8(*(*uint8)(unsafe.Pointer(uintptr(k + 541))))
					if int32(v25) != 0 {
						*(*uint8)(unsafe.Pointer(uintptr(k + 541))) = uint8(int8(int32(v25) - 1))
					}
					if int32(*(*uint8)(unsafe.Pointer(uintptr(k + 540)))) != 0 {
						var v26 int32 = int32(*(*uint32)(unsafe.Pointer(uintptr(k + 16))))
						if (v26 & 0x8000) == 0 {
							nox_xxx_updatePoison_4EE8F0(k, 1)
						}
					}
					*(*uint16)(unsafe.Pointer(uintptr(k + 542))) = 1000
				}
			}
		}
		nox_xxx_updateUnitBuffs_4FF620(k)
		if int32(*(*uint8)(unsafe.Pointer(uintptr(k + 540)))) != 0 {
			var v27 *uint16 = *(**uint16)(unsafe.Pointer(uintptr(k + 556)))
			if v27 != nil {
				if int32(*(*uint16)(unsafe.Add(unsafe.Pointer(v27), unsafe.Sizeof(uint16(0))*2))) > 0 && int32(*v27) > 0 {
					var (
						v28 int32 = bool2int(noxflags.HasGame(noxflags.GameModeQuest))
						v29 int32 = int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(k + 556))) + 16))))
						v30 int32 = v28 + 1
					)
					if v29 == 0 || (nox_frame_xxx_2598000-uint32(v29)) > 60 {
						var v31 uint8 = *(*uint8)(unsafe.Pointer(uintptr(k + 540)))
						if int32(v31) > 8 || (nox_frame_xxx_2598000%uint32(128>>(int32(v31)-1))) == 0 {
							(cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(k + 716))), (*func(int32, uint32, uint32, int32, int32))(nil)))(k, 0, 0, v30, 5)
						}
					}
				}
			}
		}
	}
}
func nox_xxx_updateUnits_51B100_D() {
	var v32 *uint8 = (*uint8)(nox_xxx_wallSecretGetFirstWall_410780())
	if v32 != nil {
		var v33 int8 = 0
		for ; v32 != nil; v32 = (*uint8)(unsafe.Pointer(uintptr(nox_xxx_wallSecretNext_410790((*int32)(unsafe.Pointer(v32)))))) {
			switch *(*uint8)(unsafe.Add(unsafe.Pointer(v32), 21)) {
			case 1:
				var v47 int8 = int8(*(*uint8)(unsafe.Add(unsafe.Pointer(v32), 20)))
				if !((int32(v47)&4) != 0 && (int32(v47)&8) == 8) {
					v33 = 0
					break
				}
				var v48 int32 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v32))), unsafe.Sizeof(uint32(0))*6))) - 1)
				*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v32))), unsafe.Sizeof(uint32(0))*6))) = uint32(v48)
				if v48 != 0 {
					v33 = 0
					break
				}
				var v49 int32 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v32))), unsafe.Sizeof(uint32(0))*1))))
				*(*uint8)(unsafe.Add(unsafe.Pointer(v32), 21)) = 4
				var v50 int32 = v49 * 23
				var v51 int32 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v32))), unsafe.Sizeof(uint32(0))*2))))
				var v52 float64 = float64(v50 + 11)
				var v53 int32 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v32))), unsafe.Sizeof(uint32(0))*3))))
				var v64 float2
				v64.field_0 = float32(v52)
				v64.field_4 = float32(float64(v51*23 + 11))
				var v54 *byte = nox_xxx_wallFindOpenSound_410EE0(int32(*(*uint8)(unsafe.Pointer(uintptr(v53 + 1)))))
				var v55 int32 = nox_xxx_utilFindSound_40AF50(v54)
				nox_xxx_audCreate_501A30(v55, &v64, 0, 0)
				v33 = 0
			case 2:
				var v36 int8 = int8(int32(*(*uint8)(unsafe.Add(unsafe.Pointer(v32), 22))) - 1)
				*(*uint8)(unsafe.Add(unsafe.Pointer(v32), 22)) = uint8(v36)
				if int32(v36) == 0 {
					var v37 int32 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v32))), unsafe.Sizeof(uint32(0))*4))))
					*(*uint8)(unsafe.Add(unsafe.Pointer(v32), 21)) = 1
					*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v32))), unsafe.Sizeof(uint32(0))*6))) = nox_gameFPS * uint32(v37)
				}
				v33 = 1
			case 3:
				var (
					v38 int8  = int8(*(*uint8)(unsafe.Add(unsafe.Pointer(v32), 20)))
					v39 int32 = 0
				)
				if !((int32(v38)&4) == 0 || int32(v38)&8 != 0 || (func() bool {
					v39 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v32))), unsafe.Sizeof(uint32(0))*6))) - 1)
					return (func() uint32 {
						p := ((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v32))), unsafe.Sizeof(uint32(0))*6)))
						*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v32))), unsafe.Sizeof(uint32(0))*6))) = uint32(v39)
						return *p
					}()) != 0
				}())) {
					var v40 int32 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v32))), unsafe.Sizeof(uint32(0))*1))))
					*(*uint8)(unsafe.Add(unsafe.Pointer(v32), 21)) = 2
					var v41 int32 = v40 * 23
					var v42 int32 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v32))), unsafe.Sizeof(uint32(0))*2))))
					var v43 float64 = float64(v41 + 11)
					var v44 int32 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v32))), unsafe.Sizeof(uint32(0))*3))))
					var v63 float2
					v63.field_0 = float32(v43)
					v63.field_4 = float32(float64(v42*23 + 11))
					var v45 *byte = nox_xxx_wallFindCloseSound_410F20(int32(*(*uint8)(unsafe.Pointer(uintptr(v44 + 1)))))
					var v46 int32 = nox_xxx_utilFindSound_40AF50(v45)
					nox_xxx_audCreate_501A30(v46, &v63, 0, 0)
				}
				v33 = 0
			case 4:
				var v34 int8 = int8(int32(*(*uint8)(unsafe.Add(unsafe.Pointer(v32), 22))) + 1)
				*(*uint8)(unsafe.Add(unsafe.Pointer(v32), 22)) = uint8(v34)
				if int32(v34) == 23 {
					var v35 int32 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v32))), unsafe.Sizeof(uint32(0))*4))))
					*(*uint8)(unsafe.Add(unsafe.Pointer(v32), 21)) = 3
					*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v32))), unsafe.Sizeof(uint32(0))*6))) = nox_gameFPS * uint32(v35)
				}
				v33 = 1
			default:
			}
			if int32(v33) != 0 {
				var (
					v56 float64 = float64(int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v32))), unsafe.Sizeof(uint32(0))*1)))*23)) + 11.5
					v65 float64 = float64(int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v32))), unsafe.Sizeof(uint32(0))*2)))*23)) + 11.5
					v66 float4
				)
				v66.field_0 = float32(v56 - 42.5)
				v66.field_4 = float32(v65 - 42.5)
				v66.field_8 = float32(v56 + 42.5)
				v66.field_C = float32(v65 + 42.5)
				nox_xxx_getUnitsInRect_517C10(&v66, func(arg1 *float32, arg2 int32) {
					sub_51B860(int32(uintptr(unsafe.Pointer(arg1))))
				}, nil)
			}
		}
	}
}
func sub_5524C0() {
	dword_5d4594_2495920 = platformTicks()
	for i := int32(0); i < NOX_NET_STRUCT_MAX; i++ {
		var ns *nox_net_struct_t = nox_net_struct_arr[i]
		if ns != nil && ns.field_38 == 1 {
			if ns.field_40+300 < nox_frame_xxx_2598000 {
				nox_xxx_netStructReadPackets_5545B0(uint32(i))
			}
		}
	}
}
func sub_57B140() bool {
	var (
		v0     uint64
		result int32
	)
	result = 0
	if *memmap.PtrUint64(6112660, 2523796) != 0 {
		v0 = uint64(nox_xxx___Getcvt_57B180() + 5000)
		if v0 < uint64(platformTicks()) {
			result = 1
		}
	}
	return result != 0
}
