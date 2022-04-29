package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	noxflags "github.com/noxworld-dev/opennox/v1/common/flags"
	"unsafe"
)

func nox_xxx_netOnPacketRecvServ_51BAD0_net_sdecode(a1 int32, operationId *uint8, a3 int32) int32 {
	var (
		data     *uint8
		v5       *byte
		v6       *byte
		v7       int32
		v8       int32
		unit     int32
		v10      *int32
		v11      *byte
		v12      uint16
		v13      int32
		v14      int32
		v15      int32
		v16      int32
		v17      int32
		v18      int32
		v19      int32
		v20      int32
		v21      uint16
		v22      int32
		v23      int32
		v24      int32
		v26      int32
		v27      int32
		v28      int32
		v29      int32
		v30      func(int32, int32, uint32)
		v31      int32
		v32      int32
		v33      int32
		v34      int32
		v35      int32
		v36      *int32
		v37      int32
		v38      *uint32
		v39      int32
		v40      int32
		v41      int32
		v42      *libc.WChar
		v43      *libc.WChar
		v44      int32
		v45      int32
		v49      *byte
		v51      int32
		v52      *byte
		v53      int32
		v54      *byte
		v55      *byte
		v56      int32
		v57      int32
		v58      int32
		v59      int32
		v60      int32
		v61      int32
		v62      uint8
		v63      *uint32
		v64      int32
		v65      int32
		v66      int32
		v67      int32
		v68      *int32
		v69      int32
		v70      int32
		v71      int32
		v72      *uint32
		v73      int32
		v74      int32
		v75      int32
		v76      int32
		v77      int32
		v78      int32
		v79      uint8
		v80      int32
		v81      int32
		v82      uint32
		v83      float32
		v84      float32
		v85      int32
		v86      *byte
		v87      [2]byte
		v88      [4]byte
		v89      *byte
		v90      [6]byte
		v91      *uint8
		v92      [5]byte
		v93      int32
		v94      float2
		v96      float2
		v97      [256]libc.WChar
		FileName [1024]byte
	)
	if a3 != 0 {
	}
	if a3 <= 0 {
		*((*uint32)(unsafe.Add(unsafe.Pointer(&nox_common_playerInfoFromNum_417090(a1).field_0), unsafe.Sizeof(uint32(0))*899))) = nox_frame_xxx_2598000
		return 1
	}
	data = operationId
	if nox_common_getEngineFlag(nox_engine_flag(NOX_ENGINE_FLAG_REPLAY_WRITE)) {
		v5 = (*byte)(unsafe.Pointer(nox_common_playerInfoFromNum_417090(a1)))
		nox_xxx_replayWriteMSgMB_4D3450((*nox_playerInfo)(unsafe.Pointer(v5)), unsafe.Pointer(operationId), uint32(a3))
	}
	switch *operationId {
	case 32:
		if nox_xxx_playerNew_4DD320(a1, (*uint8)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(operationId), 1))))))))) == 0 {
			nox_xxx_netStructReadPackets_5545B0(uint32(a1 + 1))
		}
		return 1
	case 34:
		nox_xxx_playerForceDisconnect_4DE7C0(a1)
		return 1
	case 37:
		*((*uint32)(unsafe.Add(unsafe.Pointer(&nox_common_playerInfoFromNum_417090(a1).field_0), unsafe.Sizeof(uint32(0))*899))) = nox_frame_xxx_2598000
		return 1
	}
	v91 = (*uint8)(unsafe.Add(unsafe.Pointer(operationId), a3))
	v6 = (*byte)(unsafe.Pointer(nox_common_playerInfoFromNum_417090(a1)))
	v8 = int32(uintptr(unsafe.Pointer(v6)))
	v89 = v6
	if v6 == nil {
		return 1
	}
	v93 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v6))), unsafe.Sizeof(uint32(0))*514))))
	unit = v93
	if v93 == 0 {
		return 1
	}
	v10 = *(**int32)(unsafe.Pointer(uintptr(v93 + 748)))
	v85 = int32(*(*uint32)(unsafe.Pointer(uintptr(v93 + 748))))
	if uintptr(unsafe.Pointer(operationId)) >= uintptr(unsafe.Pointer(v91)) {
		*(*uint32)(unsafe.Pointer(uintptr(v8 + 3596))) = nox_frame_xxx_2598000
		return 1
	}
	for uintptr(unsafe.Pointer(data)) < uintptr(unsafe.Pointer(v91)) {
		var old *uint8 = data
		v11 = nil
		var op int32 = int32(*(*uint8)(unsafe.Add(unsafe.Pointer(data), 0)))
		switch op {
		case 38:
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 2))
		case 41:
			nox_xxx_netNeedTimestampStatus_4174F0(v8, 64)
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 1))
		case 63:
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), nox_xxx_playerSaveInput_51A960(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(v10), unsafe.Sizeof(int32(0))*69)) + 2064)))), (*uint8)(unsafe.Add(unsafe.Pointer(data), 1)))+1))
		case 64:
			v84 = float32(float64(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 5))))))
			v83 = float32(float64(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 3))))))
			nox_xxx_playerSetCustomWP_4F79A0(unit, *(*int32)(unsafe.Add(unsafe.Pointer((*int32)(unsafe.Pointer(&v83))), unsafe.Sizeof(int32(0))*0)), *(*int32)(unsafe.Add(unsafe.Pointer((*int32)(unsafe.Pointer(&v84))), unsafe.Sizeof(int32(0))*0)))
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 7))
		case 114:
			v19 = nox_xxx_packetDynamicUnitCode_578B40(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
			if nox_common_getEngineFlag(nox_engine_flag(NOX_ENGINE_FLAG_ENABLE_NET_DEBUG)) {
				nox_xxx_netTestHighBit_578B70(uint32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
			}
			v7 = int32(*(*uint32)(unsafe.Pointer(uintptr(v85 + 276))))
			if (int32(*(*uint8)(unsafe.Pointer(uintptr(v7 + 3680))))&3) == 0 && *(*uint32)(unsafe.Pointer(uintptr(v85 + 280))) == 0 && *(*uint32)(unsafe.Pointer(uintptr(v85 + 284))) == 0 && (int32(*(*uint8)(unsafe.Pointer(uintptr(unit + 16))))&2) == 0 {
				v20 = nox_xxx_equipedItemByCode_4F7920(unit, v19)
				if v20 != 0 {
					v21 = *(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 5))))
					v94.field_0 = float32(float64(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 3))))))
					v94.field_4 = float32(float64(v21))
					nox_xxx_drop_4ED810(unit, v20, &v94.field_0)
				}
			}
			v10 = (*int32)(unsafe.Pointer(uintptr(v85)))
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 7))
		case 115:
			v22 = nox_xxx_packetDynamicUnitCode_578B40(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
			if nox_common_getEngineFlag(nox_engine_flag(NOX_ENGINE_FLAG_ENABLE_NET_DEBUG)) {
				nox_xxx_netTestHighBit_578B70(uint32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
			}
			if nox_xxx_gameGet_4DB1B0() == 0 {
				v7 = int32(*(*uint32)(unsafe.Pointer(uintptr(v85 + 276))))
				if !((int32(*(*uint8)(unsafe.Pointer(uintptr(v7 + 3680))))&3) != 0 || (*(*uint32)(unsafe.Pointer(uintptr(v85 + 280)))) != 0 || (*(*uint32)(unsafe.Pointer(uintptr(v85 + 284)))) != 0 || (int32(*(*uint8)(unsafe.Pointer(uintptr(unit + 16))))&2) != 0) {
					v23 = nox_server_getObjectFromNetCode_4ECCB0(v22)
					if v23 != 0 {
						v24 = 0
						for i := int32(unit.FirstItem()); i != 0; i = nox_xxx_inventoryGetNext_4E7990(i) {
							v24 += int32(*(*uint8)(unsafe.Pointer(uintptr(i + 488))))
						}
						if v24+int32(*(*uint8)(unsafe.Pointer(uintptr(v23 + 488)))) <= int32(*(*uint16)(unsafe.Pointer(uintptr(unit + 490)))) {
							OnLibraryNotice_420(uint32(unit), uint32(v23), uint32(unit), uint32(v23))
						} else {
							nox_xxx_netPriMsgToPlayer_4DA2C0((*nox_object_t)(unsafe.Pointer(uintptr(unit))), CString("pickup.c:CarryingTooMuch"), 0)
						}
					}
				}
			}
			v8 = int32(uintptr(unsafe.Pointer(v89)))
			v10 = (*int32)(unsafe.Pointer(uintptr(v85)))
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 3))
		case 116:
			v26 = nox_xxx_packetDynamicUnitCode_578B40(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
			if nox_common_getEngineFlag(nox_engine_flag(NOX_ENGINE_FLAG_ENABLE_NET_DEBUG)) {
				nox_xxx_netTestHighBit_578B70(uint32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
			}
			if nox_xxx_gameGet_4DB1B0() == 0 {
				v7 = int32(*(*uint32)(unsafe.Pointer(uintptr(v85 + 276))))
				if (int32(*(*uint8)(unsafe.Pointer(uintptr(v7 + 3680)))) & 3) == 0 {
					v27 = nox_xxx_equipedItemByCode_4F7920(unit, v26)
					if v27 != 0 {
						nox_xxx_useByNetCode_53F8E0(unit, v27)
					}
				}
			}
			v10 = (*int32)(unsafe.Pointer(uintptr(v85)))
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 3))
		case 117:
			v15 = nox_xxx_packetDynamicUnitCode_578B40(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
			if nox_common_getEngineFlag(nox_engine_flag(NOX_ENGINE_FLAG_ENABLE_NET_DEBUG)) {
				nox_xxx_netTestHighBit_578B70(uint32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
			}
			if nox_xxx_gameGet_4DB1B0() == 0 {
				v7 = v85
				if (int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v85 + 276))) + 3680)))) & 3) == 0 {
					v16 = nox_xxx_equipedItemByCode_4F7920(unit, v15)
					if v16 != 0 {
						nox_xxx_playerTryEquip_4F2F70((*uint32)(unsafe.Pointer(uintptr(unit))), (*nox_object_t)(unsafe.Pointer(uintptr(v16))))
					}
				}
			}
			v10 = (*int32)(unsafe.Pointer(uintptr(v85)))
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 3))
		case 118:
			v17 = nox_xxx_packetDynamicUnitCode_578B40(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
			if nox_common_getEngineFlag(nox_engine_flag(NOX_ENGINE_FLAG_ENABLE_NET_DEBUG)) {
				nox_xxx_netTestHighBit_578B70(uint32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
			}
			if (int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v85 + 276))) + 3680)))) & 3) == 0 {
				v18 = nox_xxx_equipedItemByCode_4F7920(unit, v17)
				if v18 != 0 {
					v7 = v85
					if !(int32(*(*uint8)(unsafe.Pointer(uintptr(v85 + 88)))) == 1 && *(*uint32)(unsafe.Pointer(uintptr(v18 + 8)))&0x1000000 != 0 && int32(*(*uint8)(unsafe.Pointer(uintptr(v18 + 12))))&8 != 0) {
						nox_xxx_playerTryDequip_4F2FB0((*uint32)(unsafe.Pointer(uintptr(unit))), (*nox_object_t)(unsafe.Pointer(uintptr(v18))))
					}
				}
			}
			v10 = (*int32)(unsafe.Pointer(uintptr(v85)))
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 3))
		case 120:
			v31 = nox_xxx_packetDynamicUnitCode_578B40(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
			if nox_common_getEngineFlag(nox_engine_flag(NOX_ENGINE_FLAG_ENABLE_NET_DEBUG)) {
				nox_xxx_netTestHighBit_578B70(uint32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
			}
			v7 = v85
			if (int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v85 + 276))) + 3680)))) & 1) == 0 {
				if int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))) != 0 {
					v32 = nox_server_getObjectFromNetCode_4ECCB0(v31)
					if v32 != 0 {
						nox_xxx_orderUnit_533900(unit, v32, int32(*(*uint8)(unsafe.Add(unsafe.Pointer(data), 3))))
					}
				} else {
					nox_xxx_orderUnit_533900(unit, 0, int32(*(*uint8)(unsafe.Add(unsafe.Pointer(data), 3))))
				}
			}
			v10 = (*int32)(unsafe.Pointer(uintptr(v85)))
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 4))
		case 121:
			v34 = 1
			if int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v85 + 276))) + 3680))))&1 != 0 {
				nox_xxx_netPriMsgToPlayer_4DA2C0((*nox_object_t)(unsafe.Pointer(uintptr(unit))), CString("GeneralPrint:NoSpellWarningGeneral"), 0)
				v34 = 0
			}
			if int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v85 + 276))) + 3680))))&2 != 0 {
				nox_xxx_netPriMsgToPlayer_4DA2C0((*nox_object_t)(unsafe.Pointer(uintptr(unit))), CString("GeneralPrint:ConjureNoSpellWarning1"), 0)
				v34 = 0
			}
			if !noxflags.HasGame(noxflags.GameModeCoop) {
				v35 = int32(*(*uint32)(unsafe.Pointer(uintptr(unit + 16))))
				if v35&0x4000 != 0 {
					v34 = 0
				}
			}
			if !noxflags.HasGame(noxflags.GameModeChat) && v34 != 0 {
				v36 = (*int32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))
				v37 = 0
				v38 = (*uint32)(unsafe.Add(unsafe.Pointer(data), 1))
				v39 = 5
				for v39 != 0 {
					if *v38 != 0 {
						v37++
					}
					v38 = (*uint32)(unsafe.Add(unsafe.Pointer(v38), unsafe.Sizeof(uint32(0))*1))
					v39--
				}
				if (v37 != 1 || !nox_xxx_spellHasFlags_424A50(*v36, 32) || *(*uint32)(unsafe.Pointer(uintptr(v85 + 288))) == 0 || nox_xxx_unitIsEnemyTo_5330C0((*nox_object_t)(unsafe.Pointer(uintptr(unit))), (*nox_object_t)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v85 + 288))))))) != 0 || noxflags.HasGame(noxflags.GameModeQuest)) && nox_xxx_spellByBookInsert_4FE340(unit, (*int32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1)))), v37, 3, int32(*(*uint8)(unsafe.Add(unsafe.Pointer(data), 21)))) == 0 && v37 == 1 {
					v40 = 5
					for v40 != 0 {
						if *v36 != 0 {
							nox_xxx_netReportSpellStat_4D9630(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v85 + 276))) + 2064)))), *v36, 0)
						}
						v36 = (*int32)(unsafe.Add(unsafe.Pointer(v36), unsafe.Sizeof(int32(0))*1))
						v40--
					}
				}
			}
			v8 = int32(uintptr(unsafe.Pointer(v89)))
			v10 = (*int32)(unsafe.Pointer(uintptr(v85)))
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 22))
		case 122:
			v33 = bool2int((int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v85 + 276))) + 3680)))) & 3) == 0)
			if !noxflags.HasGame(noxflags.GameModeChat) && v33 != 0 {
				nox_xxx_playerExecuteAbil_4FBB70(unit, int32(*(*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))
			}
			v10 = (*int32)(unsafe.Pointer(uintptr(v85)))
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 2))
		case 123:
			v28 = nox_xxx_packetDynamicUnitCode_578B40(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
			if nox_common_getEngineFlag(nox_engine_flag(NOX_ENGINE_FLAG_ENABLE_NET_DEBUG)) {
				nox_xxx_netTestHighBit_578B70(uint32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
			}
			v7 = v85
			if (int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v85 + 276))) + 3680))))&3) == 0 && *(*uint32)(unsafe.Pointer(uintptr(v85 + 280))) == 0 && *(*uint32)(unsafe.Pointer(uintptr(v85 + 284))) == 0 {
				v29 = nox_server_getObjectFromNetCode_4ECCB0(v28)
				if v29 != 0 {
					v30 = cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(v29 + 696))), (*func(int32, int32, uint32))(nil))
					if v30 != nil {
						v30(v29, unit, 0)
					}
				}
			}
			v10 = (*int32)(unsafe.Pointer(uintptr(v85)))
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 3))
		case 165:
			*(*uint16)(unsafe.Pointer(uintptr(v8 + int32(*(*uint8)(unsafe.Add(unsafe.Pointer(data), 1)))*8 + 16))) = *((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(data))), unsafe.Sizeof(uint16(0))*1)))
			*(*uint16)(unsafe.Pointer(uintptr(v8 + int32(*(*uint8)(unsafe.Add(unsafe.Pointer(data), 1)))*8 + 18))) = *((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(data))), unsafe.Sizeof(uint16(0))*2)))
			v14 = int32(*(*uint8)(unsafe.Add(unsafe.Pointer(data), 1)))
			v7 = int32(*(*uint32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 6)))))
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 10))
			*(*uint32)(unsafe.Pointer(uintptr(v8 + v14*8 + 20))) = uint32(v7)
		case 168:
			if int32(*(*uint8)(unsafe.Add(unsafe.Pointer(data), 3)))&2 != 0 {
				nox_swprintf(&v97[0], libc.CWString("%S"), (*uint8)(unsafe.Add(unsafe.Pointer(data), 11)))
				v44 = 1
			} else {
				nox_wcscpy(&v97[0], (*libc.WChar)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 11)))))
				v44 = 2
			}
			v45 = v44*int32(*(*uint8)(unsafe.Add(unsafe.Pointer(data), 8))) + 11
			if nox_xxx_giant_57A190(v8) == 0 {
				if (int32(*(*uint8)(unsafe.Add(unsafe.Pointer(data), 3))) & 1) == 0 {
					for j := (*byte)((*byte)(unsafe.Pointer(nox_common_playerInfoGetFirst_416EA0()))); j != nil; j = (*byte)(unsafe.Pointer(nox_common_playerInfoGetNext_416EE0((*nox_playerInfo)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(j))))))))) {
						if noxflags.HasGame(noxflags.GameClient) && *(*byte)(unsafe.Add(unsafe.Pointer(j), 2064)) == 31 {
							nox_xxx_netOnPacketRecvCli_48EA70(31, (*uint8)(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer(data)))))), v45)
						} else {
							nox_xxx_netSendSock_552640(uint32(int32(uint8(*(*byte)(unsafe.Add(unsafe.Pointer(j), 2064))))+1), (*byte)(unsafe.Pointer(data)), v45, 0)
							nox_xxx_netSendReadPacket_5528B0(uint32(int32(uint8(*(*byte)(unsafe.Add(unsafe.Pointer(j), 2064))))+1), 1)
						}
					}
				} else {
					var v47 *uint32 = nox_xxx_objGetTeamByNetCode_418C80(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
					if v47 != nil {
						if nox_xxx_servObjectHasTeam_419130(int32(uintptr(unsafe.Pointer(v47)))) != 0 {
							v86 = (*byte)(unsafe.Pointer(nox_xxx_clientGetTeamColor_418AB0(int32(*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(v47))), 4)))))))
							if v86 != nil {
								for j := int32(int32(uintptr(unsafe.Pointer(nox_xxx_getFirstPlayerUnit_4DA7C0())))); j != 0; j = int32(uintptr(unsafe.Pointer(nox_xxx_getNextPlayerUnit_4DA7F0((*nox_object_t)(unsafe.Pointer(uintptr(j))))))) {
									v49 = (*byte)(unsafe.Pointer(noxServer.getPlayerByID(int32(*(*uint32)(unsafe.Pointer(uintptr(j + 36)))))))
									if v49 != nil && nox_xxx_teamCompare2_419180(j+48, uint8(*(*byte)(unsafe.Add(unsafe.Pointer(v86), 57)))) != 0 {
										if noxflags.HasGame(noxflags.GameClient) && *(*uint32)(unsafe.Pointer(uintptr(j + 36))) == nox_player_netCode_85319C {
											nox_xxx_netOnPacketRecvCli_48EA70(int32(uint8(*(*byte)(unsafe.Add(unsafe.Pointer(v49), 2064)))), (*uint8)(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer(data)))))), v45)
										} else {
											nox_xxx_netSendSock_552640(uint32(int32(uint8(*(*byte)(unsafe.Add(unsafe.Pointer(v49), 2064))))+1), (*byte)(unsafe.Pointer(data)), v45, 0)
											nox_xxx_netSendReadPacket_5528B0(uint32(int32(uint8(*(*byte)(unsafe.Add(unsafe.Pointer(v49), 2064))))+1), 1)
										}
									}
								}
							}
						}
					}
				}
			}
			unit = v93
			v8 = int32(uintptr(unsafe.Pointer(v89)))
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), v45))
			v10 = (*int32)(unsafe.Pointer(uintptr(v85)))
		case 170:
			v13 = *(*int32)(unsafe.Add(unsafe.Pointer(v10), unsafe.Sizeof(int32(0))*69))
			*(*uint32)(unsafe.Pointer(&v92[1])) = *(*uint32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))
			v92[0] = 171
			nox_netlist_addToMsgListCli_40EBC0(int32(*(*uint8)(unsafe.Pointer(uintptr(v13 + 2064)))), 1, (*uint8)(unsafe.Pointer(&v92[0])), 5)
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 5))
		case 171:
			sub_4E55A0(*(*uint8)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(v10), unsafe.Sizeof(int32(0))*69)) + 2064))), int32(*(*uint32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 5))
		case 172:
			v12 = *(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 5))
			*(*uint32)(unsafe.Pointer(uintptr(v8 + 2284))) = uint32(v12)
			*(*uint32)(unsafe.Pointer(uintptr(v8 + 2288))) = uint32(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(data))), -int(unsafe.Sizeof(uint16(0))*1)))))
		case 173:
			nox_xxx_netPlayerIncomingServ_4DDF60(int32(*(*uint8)(unsafe.Pointer(uintptr(v8 + 2064)))))
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 1))
		case 174:
			nox_xxx_playerDisconnFinish_4DE530(int32(*(*uint8)(unsafe.Pointer(uintptr(v8 + 2064)))), 2)
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 1))
		case 182:
			nox_xxx_playerGoObserver_4E6860((*nox_playerInfo)(unsafe.Pointer(uintptr(v8))), 1, 1)
			v51 = int32(*(*uint32)(unsafe.Pointer(uintptr(v8 + 2056))))
			if v51 != 0 {
				nox_xxx_netChangeTeamMb_419570(v51+48, int32(*(*uint32)(unsafe.Pointer(uintptr(v8 + 2060)))))
			}
			nox_xxx_netMapSend_519D20(int32(*(*uint8)(unsafe.Pointer(uintptr(v8 + 2064)))))
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 1))
		case 183:
			nox_xxx_netMapSendCancelMap_519DE0_net_mapsend(int32(*(*uint8)(unsafe.Pointer(uintptr(v8 + 2064)))))
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 1))
		case 187:
			nox_xxx_serverHandleClientConsole_443E90((*nox_playerInfo)(unsafe.Pointer(uintptr(v8))), int8(*(*uint8)(unsafe.Add(unsafe.Pointer(data), 1))), (*libc.WChar)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 5)))))
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), int32(*(*uint8)(unsafe.Add(unsafe.Pointer(data), 4)))*2+7))
		case 188:
			v87[0] = 189
			v42 = nox_xxx_sysopGetPass_40A630()
			if *v42 == 0 || _nox_wcsicmp((*libc.WChar)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1)))), v42) != 0 {
				v87[1] = 0
			} else {
				v87[1] = 1
				if sub_4D12A0(int32(*(*uint8)(unsafe.Pointer(uintptr(v8 + 2064))))) == 0 {
					sub_4D1210(int32(*(*uint8)(unsafe.Pointer(uintptr(v8 + 2064)))))
					v43 = strMan.GetStringInFile("sysopAccessGranted", "C:\\NoxPost\\src\\Server\\Network\\sdecode.c")
					nox_gui_console_Printf_450C00(uint8(int8(NOX_CONSOLE_RED)), v43, v8+4704)
				}
			}
			nox_xxx_netSendPacket0_4E5420(int32(*(*uint8)(unsafe.Pointer(uintptr(v8 + 2064)))), unsafe.Pointer(&v87[0]), 2, 0, 1)
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 21))
		case 190:
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 1))
		case 191:
			*(*uint8)(unsafe.Pointer(uintptr(v8 + 3676))) = 3
			sub_519E80(int32(*(*uint8)(unsafe.Pointer(uintptr(v8 + 2064)))))
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 1))
		case 192:
			nox_xxx_gameServerReadyMB_4DD180(int32(*(*uint8)(unsafe.Pointer(uintptr(v8 + 2064)))))
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 1))
		case 193:
			if noxflags.HasGame(noxflags.GameModeQuest) && int32(*(*uint8)(unsafe.Pointer(uintptr(v8 + 2064)))) != 31 && *(*uint32)(unsafe.Pointer(uintptr(v8 + 2092))) != 0 && *(*uint32)(unsafe.Pointer(uintptr(v8 + 2056))) != 0 && *(*int32)(unsafe.Add(unsafe.Pointer(v10), unsafe.Sizeof(int32(0))*138)) == 1 {
				nox_xxx_playerCallDisconnect_4DEAB0(int32(*(*uint8)(unsafe.Pointer(uintptr(v8 + 2064)))), 2)
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 3))
			} else {
				v52 = nox_fs_root()
				nox_sprintf(&FileName[0], CString("%s\\Save\\_temp_.dat"), v52)
				if nox_xxx_playerSaveToFile_41A140(&FileName[0], int32(*(*uint8)(unsafe.Pointer(uintptr(v8 + 2064))))) != 0 {
					sub_41CFA0(&FileName[0], int32(*(*uint8)(unsafe.Pointer(uintptr(v8 + 2064)))))
				}
				nox_fs_remove(&FileName[0])
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 3))
			}
		case 194:
			switch *(*uint8)(unsafe.Add(unsafe.Pointer(data), 1)) {
			case 0:
				sub_40B5D0(uint32(int32(*(*uint8)(unsafe.Pointer(uintptr(v8 + 2064))))+1), int8(*(*uint8)(unsafe.Add(unsafe.Pointer(data), 2))), (*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(data))), 8)), *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(data))), unsafe.Sizeof(uint32(0))*1))), int8(*(*uint8)(unsafe.Add(unsafe.Pointer(data), 136))))
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 140))
			case 1:
				*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v11))), 0)) = *(*uint8)(unsafe.Add(unsafe.Pointer(data), 3))
				sub_40BFF0(int32(*(*uint8)(unsafe.Pointer(uintptr(v8 + 2064))))+1, int8(*(*uint8)(unsafe.Add(unsafe.Pointer(data), 2))), v11)
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 4))
			case 2:
				v90[0] = 194
				v90[1] = 3
				v90[2] = byte(*(*uint8)(unsafe.Add(unsafe.Pointer(data), 2)))
				v82 = uint32(int32(*(*uint8)(unsafe.Pointer(uintptr(v8 + 2064)))) + 1)
				*(*uint16)(unsafe.Pointer(&v90[4])) = *((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(data))), unsafe.Sizeof(uint16(0))*2)))
				nox_xxx_netSendSock_552640(v82, &v90[0], 6, int8(NOX_NET_SEND_NO_LOCK|NOX_NET_SEND_FLAG2))
				sub_40B250(int32(*(*uint8)(unsafe.Pointer(uintptr(v8 + 2064))))+1, *(*uint8)(unsafe.Add(unsafe.Pointer(data), 2)), *((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(data))), unsafe.Sizeof(uint16(0))*2))), unsafe.Add(unsafe.Pointer(data), 8), uint32(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(data))), unsafe.Sizeof(uint16(0))*3)))))
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), int32(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(data))), unsafe.Sizeof(uint16(0))*3))))+8))
			case 3:
				sub_40BF60(int32(*(*uint8)(unsafe.Pointer(uintptr(v8 + 2064))))+1, int8(*(*uint8)(unsafe.Add(unsafe.Pointer(data), 2))), int32(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(data))), unsafe.Sizeof(uint16(0))*2)))))
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 6))
			case 4:
				sub_40C030(int32(*(*uint8)(unsafe.Pointer(uintptr(v8 + 2064))))+1, int8(*(*uint8)(unsafe.Add(unsafe.Pointer(data), 2))))
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 3))
			case 5:
				*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v7))), 0)) = *(*uint8)(unsafe.Add(unsafe.Pointer(data), 3))
				sub_40B720(v7, *(*uint8)(unsafe.Add(unsafe.Pointer(data), 2)))
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 4))
			case 6:
				*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v11))), 0)) = *(*uint8)(unsafe.Add(unsafe.Pointer(data), 3))
				sub_40C070(int32(*(*uint8)(unsafe.Pointer(uintptr(v8 + 2064))))+1, int32(uintptr(unsafe.Pointer(v11))), int8(*(*uint8)(unsafe.Add(unsafe.Pointer(data), 2))))
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 4))
			default:
			}
		case 196:
			if int32(*(*uint8)(unsafe.Add(unsafe.Pointer(data), 1))) == 10 {
				v55 = (*byte)(unsafe.Pointer(nox_xxx_clientGetTeamColor_418AB0(int32(*(*uint32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 2))))))))
				if v55 != nil {
					v56 = nox_server_getObjectFromNetCode_4ECCB0(int32(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(data))), unsafe.Sizeof(uint16(0))*3)))))
					nox_xxx_createAtImpl_4191D0(uint8(*(*byte)(unsafe.Add(unsafe.Pointer(v55), 57))), unsafe.Pointer(uintptr(v56+48)), 1, int32(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(data))), unsafe.Sizeof(uint16(0))*3)))), 1)
				}
				v10 = (*int32)(unsafe.Pointer(uintptr(v85)))
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 10))
			} else if int32(*(*uint8)(unsafe.Add(unsafe.Pointer(data), 1))) == 11 {
				v53 = nox_server_getObjectFromNetCode_4ECCB0(int32(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(data))), unsafe.Sizeof(uint16(0))*3)))))
				if v53 != 0 {
					if !noxflags.HasGame(noxflags.GameModeChat) {
						nox_xxx_mapFindPlayerStart_4F7AB0(&v96, (*nox_object_t)(unsafe.Pointer(uintptr(v53))))
						nox_xxx_unitMove_4E7010((*nox_object_t)(unsafe.Pointer(uintptr(v53))), &v96)
					}
					v54 = (*byte)(unsafe.Pointer(nox_xxx_clientGetTeamColor_418AB0(int32(*(*uint32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 2))))))))
					if v54 != nil {
						sub_4196D0(v53+48, int32(uintptr(unsafe.Pointer(v54))), int32(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(data))), unsafe.Sizeof(uint16(0))*3)))), 1)
					}
				}
				v10 = (*int32)(unsafe.Pointer(uintptr(v85)))
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 10))
			}
		case 200:
			sub_446070()
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 1))
		case 208:
			if int32(*(*uint8)(unsafe.Add(unsafe.Pointer(data), 1))) == 1 {
				if nox_xxx_gameGet_4DB1B0() != 0 || (func() int32 {
					v7 = *(*int32)(unsafe.Add(unsafe.Pointer(v10), unsafe.Sizeof(int32(0))*69))
					return int32(*(*uint8)(unsafe.Pointer(uintptr(v7 + 3680)))) & 3
				}()) != 0 || (func() int32 {
					v57 = nox_server_getObjectFromNetCode_4ECCB0(int32(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(data))), unsafe.Sizeof(uint16(0))*1)))))
					return v57
				}()) == 0 {
					data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 4))
				} else {
					nox_xxx_script_forcedialog_548CD0(unit, v57)
					data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 4))
				}
			} else if int32(*(*uint8)(unsafe.Add(unsafe.Pointer(data), 1))) == 2 {
				nox_xxx_scriptDialog_548D30(unit, int8(*(*uint8)(unsafe.Add(unsafe.Pointer(data), 2))))
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 3))
			}
		case 212:
			v41 = sub_40A220()
			nox_xxx_netTimerStatus_4D8F50(a1, v41)
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 1))
		case 224:
			v58 = nox_xxx_packetDynamicUnitCode_578B40(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
			if nox_common_getEngineFlag(nox_engine_flag(NOX_ENGINE_FLAG_ENABLE_NET_DEBUG)) {
				nox_xxx_netTestHighBit_578B70(uint32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
			}
			if int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))) != 0 {
				v59 = nox_server_getObjectFromNetCode_4ECCB0(v58)
				sub_53AB90(unit, v59)
			} else {
				sub_53AB90(unit, 0)
			}
			v10 = (*int32)(unsafe.Pointer(uintptr(v85)))
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 3))
		case 226:
			v60 = nox_xxx_packetDynamicUnitCode_578B40(int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
			if nox_common_getEngineFlag(nox_engine_flag(NOX_ENGINE_FLAG_ENABLE_NET_DEBUG)) {
				nox_xxx_netTestHighBit_578B70(uint32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1))))))
			}
			v61 = nox_xxx_equipedItemByCode_4F7920(unit, v60)
			if v61 != 0 || (func() int32 {
				v61 = sub_510DE0(unit, v60)
				return v61
			}()) != 0 || (func() int32 {
				v61 = nox_server_getObjectFromNetCode_4ECCB0(v60)
				return v61
			}()) != 0 {
				v88[0] = 226
				*(*uint16)(unsafe.Pointer(&v88[1])) = uint16(nox_xxx_netGetUnitCodeServ_578AC0((*nox_object_t)(unsafe.Pointer(uintptr(v61)))))
				v62 = *(*uint8)(unsafe.Add(unsafe.Pointer(data), 3))
				if int32(v62) == 1 {
					v88[3] = byte(**(**uint8)(unsafe.Pointer(uintptr(v61 + 736))))
				} else if int32(v62) == 2 {
					v88[3] = byte(int8(nox_xxx_guide_427010(*(**byte)(unsafe.Pointer(uintptr(v61 + 736))))))
				} else {
					v88[3] = byte(**(**uint8)(unsafe.Pointer(uintptr(v61 + 736))))
				}
				nox_xxx_netSendPacket0_4E5420(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v85 + 276))) + 2064)))), unsafe.Pointer(&v88[0]), 4, 0, 1)
			}
			v8 = int32(uintptr(unsafe.Pointer(v89)))
			v10 = (*int32)(unsafe.Pointer(uintptr(v85)))
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 4))
		case 227:
			switch *(*uint8)(unsafe.Add(unsafe.Pointer(data), 1)) {
			case 1:
				nox_xxx_playerSetState_4FA020((*nox_object_t)(unsafe.Pointer(uintptr(unit))), 26)
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 2))
			case 2:
				nox_xxx_playerSetState_4FA020((*nox_object_t)(unsafe.Pointer(uintptr(unit))), 25)
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 2))
			case 4:
				nox_xxx_playerSetState_4FA020((*nox_object_t)(unsafe.Pointer(uintptr(unit))), 27)
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 2))
			default:
			}
		case 238:
			switch *(*uint8)(unsafe.Add(unsafe.Pointer(data), 1)) {
			case 0:
				var c1 int32 = 0
				if noxflags.HasGame(noxflags.GameModeQuest) {
					c1 = 3
				}
				sub_506870(c1, unit, (*libc.WChar)(unsafe.Add(unsafe.Pointer((*libc.WChar)(unsafe.Pointer(data))), unsafe.Sizeof(libc.WChar(0))*1)))
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 52))
			case 1:
				sub_506870(1, unit, (*libc.WChar)(unsafe.Add(unsafe.Pointer((*libc.WChar)(unsafe.Pointer(data))), unsafe.Sizeof(libc.WChar(0))*1)))
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 52))
			case 2:
				var c2 int32 = 0
				if noxflags.HasGame(noxflags.GameModeQuest) {
					c2 = 3
				}
				sub_506C90(c2, unit, (*libc.WChar)(unsafe.Add(unsafe.Pointer((*libc.WChar)(unsafe.Pointer(data))), unsafe.Sizeof(libc.WChar(0))*1)))
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 52))
			case 3:
				sub_506C90(1, unit, (*libc.WChar)(unsafe.Add(unsafe.Pointer((*libc.WChar)(unsafe.Pointer(data))), unsafe.Sizeof(libc.WChar(0))*1)))
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 52))
			case 4:
				sub_506870(2, unit, nil)
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 2))
			case 5:
				sub_506C90(2, unit, nil)
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 2))
			default:
				return 0
			}
			fallthrough
		case 201:
			switch *(*uint8)(unsafe.Add(unsafe.Pointer(data), 1)) {
			case 14:
				if *(*int32)(unsafe.Add(unsafe.Pointer(v10), unsafe.Sizeof(int32(0))*70)) != 0 {
					sub_50F3A0((*uint32)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(v10), unsafe.Sizeof(int32(0))*70))))))
				}
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 2))
			case 15:
				v67 = nox_xxx_packetDynamicUnitCode_578B40(int32(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(data))), unsafe.Sizeof(uint16(0))*1)))))
				if nox_common_getEngineFlag(nox_engine_flag(NOX_ENGINE_FLAG_ENABLE_NET_DEBUG)) {
					nox_xxx_netTestHighBit_578B70(uint32(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(data))), unsafe.Sizeof(uint16(0))*1)))))
				}
				v68 = (*int32)(unsafe.Pointer(uintptr(nox_xxx_equipedItemByCode_4F7920(unit, v67))))
				if v68 != nil {
					v69 = int32(*(*uint32)(unsafe.Pointer(uintptr(v85 + 280))))
					if v69 != 0 {
						if nox_xxx_tradeP2PAddOffer2_50F820_trade(v69, unit, *(*float32)(unsafe.Pointer(&v68))) == 1 {
							sub_4ED0C0(unit, (*nox_object_t)(unsafe.Pointer(v68)))
						}
					}
				}
				v10 = (*int32)(unsafe.Pointer(uintptr(v85)))
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 4))
			case 16:
				v70 = nox_xxx_packetDynamicUnitCode_578B40(int32(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(data))), unsafe.Sizeof(uint16(0))*1)))))
				if nox_common_getEngineFlag(nox_engine_flag(NOX_ENGINE_FLAG_ENABLE_NET_DEBUG)) {
					nox_xxx_netTestHighBit_578B70(uint32(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(data))), unsafe.Sizeof(uint16(0))*1)))))
				}
				v71 = int32(*(*uint32)(unsafe.Pointer(uintptr(v85 + 280))))
				if v71 != 0 {
					nox_xxx_tradeP2PAddOfferMB_50FE20(v71, v70)
				}
				v10 = (*int32)(unsafe.Pointer(uintptr(v85)))
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 4))
			case 17:
				v66 = *(*int32)(unsafe.Add(unsafe.Pointer(v10), unsafe.Sizeof(int32(0))*70))
				if v66 != 0 {
					nox_xxx_tradeAccept_50F5A0(v66, unit)
				}
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 2))
			case 18:
				if *(*int32)(unsafe.Add(unsafe.Pointer(v10), unsafe.Sizeof(int32(0))*70)) != 0 {
					nox_xxx_shopExit_50F4C0((*uint32)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(v10), unsafe.Sizeof(int32(0))*70))))))
				}
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 2))
			case 21:
				if nox_xxx_gameGet_4DB1B0() == 0 && (int32(*(*uint8)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(v10), unsafe.Sizeof(int32(0))*69)) + 3680))))&3) == 0 {
					v64 = nox_xxx_packetDynamicUnitCode_578B40(int32(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(data))), unsafe.Sizeof(uint16(0))*1)))))
					v65 = nox_server_getObjectFromNetCode_4ECCB0(v64)
					if v65 != 0 {
						if int32(*(*uint8)(unsafe.Pointer(uintptr(v65 + 12))))&8 != 0 {
							nox_xxx_servShopStart_50EF10_trade(unit, v65)
						}
					}
				}
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 4))
			case 22:
				v72 = (*uint32)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(v10), unsafe.Sizeof(int32(0))*70)))))
				if v72 != nil {
					sub_5100C0_trade(unit, v72, int32(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(data))), unsafe.Sizeof(uint16(0))*1)))))
				}
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 4))
			case 23:
				v73 = *(*int32)(unsafe.Add(unsafe.Pointer(v10), unsafe.Sizeof(int32(0))*70))
				if v73 != 0 {
					sub_510640_trade(unit, v73, int32(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(data))), unsafe.Sizeof(uint16(0))*1)))), (*float32)(unsafe.Pointer(uintptr(*(*uint8)(unsafe.Add(unsafe.Pointer(data), 4))))))
				}
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 5))
			case 24:
				v74 = *(*int32)(unsafe.Add(unsafe.Pointer(v10), unsafe.Sizeof(int32(0))*70))
				if v74 != 0 {
					sub_510BE0_trade((*int32)(unsafe.Pointer(uintptr(unit))), v74, (*uint32)(unsafe.Pointer(uintptr(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(data))), unsafe.Sizeof(uint16(0))*1)))))))
				}
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 4))
			case 25:
				v76 = *(*int32)(unsafe.Add(unsafe.Pointer(v10), unsafe.Sizeof(int32(0))*70))
				if v76 != 0 {
					sub_510D10((*int32)(unsafe.Pointer(uintptr(unit))), v76, int32(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(data))), unsafe.Sizeof(uint16(0))*1)))), uint32(*(*uint8)(unsafe.Add(unsafe.Pointer(data), 4))))
				}
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 5))
			case 26:
				v75 = *(*int32)(unsafe.Add(unsafe.Pointer(v10), unsafe.Sizeof(int32(0))*70))
				if v75 != 0 {
					sub_510AE0((*int32)(unsafe.Pointer(uintptr(unit))), v75, (*uint32)(unsafe.Pointer(uintptr(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(data))), unsafe.Sizeof(uint16(0))*1)))))))
				}
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 4))
			case 28:
				v77 = *(*int32)(unsafe.Add(unsafe.Pointer(v10), unsafe.Sizeof(int32(0))*70))
				if v77 != 0 {
					sub_5109C0_trade((*int32)(unsafe.Pointer(uintptr(unit))), v77, (*uint32)(unsafe.Pointer(uintptr(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(data))), unsafe.Sizeof(uint16(0))*1)))))))
				}
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 4))
			case 30:
				v78 = *(*int32)(unsafe.Add(unsafe.Pointer(v10), unsafe.Sizeof(int32(0))*70))
				if v78 != 0 {
					sub_5108D0(unit, v78, int32(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(data))), unsafe.Sizeof(uint16(0))*1)))))
				}
				data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 4))
			default:
				return 0
			}
		case 240:
			v79 = *(*uint8)(unsafe.Add(unsafe.Pointer(data), 1))
			if int32(v79) == 3 {
				v80 = *(*int32)(unsafe.Add(unsafe.Pointer(v10), unsafe.Sizeof(int32(0))*69))
				v81 = int32(*(*uint32)(unsafe.Pointer(uintptr(v80 + 2056))))
				if v81 != 0 {
					v7 = int32(*(*uint32)(unsafe.Pointer(uintptr(v81 + 16))))
					if (v7 & 0x8000) != 0 {
						*(*int32)(unsafe.Add(unsafe.Pointer(v10), unsafe.Sizeof(int32(0))*137)) = 0
						nox_xxx_playerRespawn_4F7EF0(int32(*(*uint32)(unsafe.Pointer(uintptr(v80 + 2056)))))
					}
				}
			} else {
				if int32(v79) != 27 {
					break
				}
				sub_4DD0B0(unit)
			}
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 2))
		case 241:
			v63 = (*uint32)(unsafe.Pointer(uintptr(nox_xxx_equipedItemByCode_4F7920(unit, int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(data), 1)))))))))
			if v63 != nil {
				nox_xxx_drop_4ED790((*nox_object_t)(unsafe.Pointer(uintptr(unit))), (*nox_object_t)(unsafe.Pointer(v63)), (*float2)(unsafe.Pointer(uintptr(unit+56))))
				nox_xxx_netPriMsgToPlayer_4DA2C0((*nox_object_t)(unsafe.Pointer(uintptr(unit))), CString("pickup.c:CarryingTooMuch"), 0)
			}
			data = (*uint8)(unsafe.Add(unsafe.Pointer(data), 3))
		default:
			return 0
		}
		noxOnSrvPacketDebug(op, old, int32(uintptr(unsafe.Pointer(data))-uintptr(unsafe.Pointer(old))))
	}
	*(*uint32)(unsafe.Pointer(uintptr(v8 + 3596))) = nox_frame_xxx_2598000
	return 1
}
