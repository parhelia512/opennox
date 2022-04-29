package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	"unsafe"
)

func nox_client_handlePacketInform_4C9BF0(a1 int32) int32 {
	var (
		result int32
		v2     *libc.WChar
		v3     *libc.WChar
		v4     *libc.WChar
		v5     *libc.WChar
		v6     *libc.WChar
		v7     *byte
		v8     *libc.WChar
		v9     *libc.WChar
		v10    *libc.WChar
		v11    *libc.WChar
		v12    *libc.WChar
		v13    *byte
		v14    *libc.WChar
		v15    *byte
		v16    *libc.WChar
		v17    *libc.WChar
		v18    *byte
		v19    *libc.WChar
		v20    *byte
		v21    *libc.WChar
		v22    *byte
		v23    *libc.WChar
		v24    *byte
		v25    *libc.WChar
		v26    *byte
		v27    *libc.WChar
		v28    *libc.WChar
		v29    *libc.WChar
		v30    *byte
		v31    *byte
		v32    *libc.WChar
		v33    *libc.WChar
		v34    *byte
		v35    *byte
		v36    *libc.WChar
		v37    *libc.WChar
		v38    *byte
		v39    *byte
		v40    *libc.WChar
		v41    *libc.WChar
		v42    uint32
		v43    *libc.WChar
		v44    *libc.WChar
		v45    *libc.WChar
		v46    int32
		v47    *libc.WChar
		v48    int32
		v49    *libc.WChar
		v50    int32
		v51    int32
		v52    int32
		v53    *libc.WChar
		v54    *libc.WChar
		v55    *libc.WChar
		v56    *libc.WChar
		v57    int32
		v58    int32
		v59    int32
		v60    *byte
		v61    *byte
		v62    [256]libc.WChar
		v63    [256]libc.WChar
	)
	result = 0
	switch *(*uint8)(unsafe.Pointer(uintptr(a1 + 1))) {
	case 0:
		nox_xxx_abilGetError_4FB0B0_magic_plyrspel(int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 2)))))
		return 6
	case 1:
		v50 = int32(uintptr(unsafe.Pointer(nox_xxx_spellTitle_424930(int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 2))))))))
		v10 = strMan.GetStringInFile("plyrspel.c:SpellCastSuccess", "C:\\NoxPost\\src\\client\\Network\\inform.c")
		nox_swprintf(&v62[0], v10, v50)
		nox_xxx_printCentered_445490(&v62[0])
		return 6
	case 2:
		nox_xxx_abilGetSuccess_4FB960_ability(int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 2)))))
		return 6
	case 3:
		v13 = (*byte)(unsafe.Pointer(noxServer.getPlayerByID(int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 2)))))))
		if v13 == nil {
			return 6
		}
		v51 = int32(uintptr(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v13), 4704)))))
		v14 = strMan.GetStringInFile("netserv.c:PlayerTimeout", "C:\\NoxPost\\src\\client\\Network\\inform.c")
		nox_swprintf(&v62[0], v14, v51)
		nox_xxx_printCentered_445490(&v62[0])
		return 6
	case 4:
		v15 = (*byte)(unsafe.Pointer(noxServer.getPlayerByID(int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 2)))))))
		if v15 == nil {
			return 6
		}
		v52 = int32(uintptr(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v15), 4704)))))
		v16 = strMan.GetStringInFile("objcoll.c:FlagRetrieveNotice", "C:\\NoxPost\\src\\client\\Network\\inform.c")
		nox_swprintf(&v62[0], v16, v52)
		goto LABEL_22
	case 5:
		v18 = (*byte)(unsafe.Pointer(noxServer.getPlayerByID(int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 2)))))))
		if v18 == nil {
			return 10
		}
		v54 = nox_server_teamTitle_418C20(int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 6)))))
		v19 = strMan.GetStringInFile("objcoll.c:FlagCaptureNotice", "C:\\NoxPost\\src\\client\\Network\\inform.c")
		nox_swprintf(&v62[0], v19, (*byte)(unsafe.Add(unsafe.Pointer(v18), 4704)), v54)
		nox_xxx_printCentered_445490(&v62[0])
		clientPlaySoundSpecial(306, 100)
		return 10
	case 6:
		v20 = (*byte)(unsafe.Pointer(noxServer.getPlayerByID(int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 2)))))))
		if v20 == nil {
			return 10
		}
		v55 = nox_server_teamTitle_418C20(int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 6)))))
		v21 = strMan.GetStringInFile("objcoll.c:FlagPickupNotice", "C:\\NoxPost\\src\\client\\Network\\inform.c")
		nox_swprintf(&v62[0], v21, (*byte)(unsafe.Add(unsafe.Pointer(v20), 4704)), v55)
		nox_xxx_printCentered_445490(&v62[0])
		clientPlaySoundSpecial(303, 100)
		return 10
	case 7:
		v22 = (*byte)(unsafe.Pointer(noxServer.getPlayerByID(int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 2)))))))
		if v22 != nil {
			v56 = nox_server_teamTitle_418C20(int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 6)))))
			v23 = strMan.GetStringInFile("drop.c:FlagDropNotice", "C:\\NoxPost\\src\\client\\Network\\inform.c")
			nox_swprintf(&v62[0], v23, (*byte)(unsafe.Add(unsafe.Pointer(v22), 4704)), v56)
			nox_xxx_printCentered_445490(&v62[0])
			clientPlaySoundSpecial(304, 100)
		}
		return 10
	case 8:
		v53 = nox_server_teamTitle_418C20(int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 2)))))
		v17 = strMan.GetStringInFile("update.c:FlagRespawnNotice", "C:\\NoxPost\\src\\client\\Network\\inform.c")
		nox_swprintf(&v62[0], v17, v53)
	LABEL_22:
		nox_xxx_printCentered_445490(&v62[0])
		clientPlaySoundSpecial(305, 100)
		return 6
	case 9:
		v38 = (*byte)(unsafe.Pointer(nox_xxx_clientGetTeamColor_418AB0(int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 6)))))))
		v39 = (*byte)(unsafe.Pointer(noxServer.getPlayerByID(int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 2)))))))
		if v39 == nil {
			if v38 != nil {
				v41 = strMan.GetStringInFile("objcoll.c:FlagBallUnknownNotice", "C:\\NoxPost\\src\\client\\Network\\inform.c")
				nox_swprintf(&v62[0], v41, v38)
				nox_xxx_printCentered_445490(&v62[0])
			}
			return 10
		}
		if v38 == nil {
			return 10
		}
		v46 = int32(uintptr(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v39), 4704)))))
		v40 = strMan.GetStringInFile("objcoll.c:FlagBallNotice", "C:\\NoxPost\\src\\client\\Network\\inform.c")
		nox_swprintf(&v62[0], v40, v46, v38)
		nox_xxx_printCentered_445490(&v62[0])
		return 10
	case 10:
		v34 = (*byte)(unsafe.Pointer(noxServer.getPlayerByID(int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 2)))))))
		v35 = (*byte)(unsafe.Pointer(nox_xxx_clientGetTeamColor_418AB0(int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 6)))))))
		if v35 != nil {
			if v34 == nil {
				return 10
			}
			v61 = v35
			v36 = strMan.GetStringInFile("pickup.c:PickUpTeamCrown", "C:\\NoxPost\\src\\client\\Network\\inform.c")
			nox_swprintf(&v62[0], v36, (*byte)(unsafe.Add(unsafe.Pointer(v34), 4704)), v61)
			nox_xxx_printCentered_445490(&v62[0])
			result = 10
		} else {
			if v34 == nil {
				return 10
			}
			v37 = strMan.GetStringInFile("pickup.c:PickUpCrown", "C:\\NoxPost\\src\\client\\Network\\inform.c")
			nox_swprintf(&v62[0], v37, (*byte)(unsafe.Add(unsafe.Pointer(v34), 4704)))
			nox_xxx_printCentered_445490(&v62[0])
			result = 10
		}
		return result
	case 11:
		v30 = (*byte)(unsafe.Pointer(noxServer.getPlayerByID(int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 2)))))))
		v31 = (*byte)(unsafe.Pointer(nox_xxx_clientGetTeamColor_418AB0(int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 6)))))))
		if v31 != nil {
			if v30 != nil {
				v60 = v31
				v32 = strMan.GetStringInFile("drop.c:DropTeamCrown", "C:\\NoxPost\\src\\client\\Network\\inform.c")
				nox_swprintf(&v62[0], v32, (*byte)(unsafe.Add(unsafe.Pointer(v30), 4704)), v60)
				nox_xxx_printCentered_445490(&v62[0])
				return 10
			}
		} else if v30 != nil {
			v33 = strMan.GetStringInFile("drop.c:DropCrown", "C:\\NoxPost\\src\\client\\Network\\inform.c")
			nox_swprintf(&v62[0], v33, (*byte)(unsafe.Add(unsafe.Pointer(v30), 4704)))
			nox_xxx_printCentered_445490(&v62[0])
			return 10
		}
		return 10
	case 12:
		v11 = strMan.GetStringInFile("Netserv.c:InObservationMode", "C:\\NoxPost\\src\\client\\Network\\inform.c")
		nox_xxx_printCentered_445490(v11)
		if *(*uint32)(unsafe.Pointer(uintptr(a1 + 2))) == 0 {
			return 6
		}
		v12 = strMan.GetStringInFile("Netserv.c:PressJump", "C:\\NoxPost\\src\\client\\Network\\inform.c")
		nox_xxx_printCentered_445490(v12)
		return 6
	case 13:
		switch *(*uint32)(unsafe.Pointer(uintptr(a1 + 2))) {
		case 0:
			v2 = strMan.GetStringInFile("atckexec.c:PlayerStunned", "C:\\NoxPost\\src\\client\\Network\\inform.c")
			nox_xxx_printCentered_445490(v2)
			result = 6
		case 1:
			v3 = strMan.GetStringInFile("atckexec.c:PlayerConfused", "C:\\NoxPost\\src\\client\\Network\\inform.c")
			nox_xxx_printCentered_445490(v3)
			result = 6
		case 2:
			v4 = strMan.GetStringInFile("atckexec.c:PlayerPoisoned", "C:\\NoxPost\\src\\client\\Network\\inform.c")
			nox_xxx_printCentered_445490(v4)
			result = 6
		case 3:
			v5 = strMan.GetStringInFile("player.c:TooHeavy", "C:\\NoxPost\\src\\client\\Network\\inform.c")
			nox_xxx_printCentered_445490(v5)
			result = 6
		default:
			return 6
		}
		return result
	case 14:
		sub_495210(a1)
		return 11
	case 15:
		v42 = uint32(libc.StrLen((*byte)(unsafe.Pointer(uintptr(a1+3)))) + 1)
		if nox_xxx_gameGetPlayState_4356B0() != 3 {
			goto LABEL_57
		}
		if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 2)))) != 0 {
			v47 = strMan.GetStringInFile((*byte)(unsafe.Pointer(uintptr(a1+3))), "C:\\NoxPost\\src\\client\\Network\\inform.c")
			v43 = strMan.GetStringInFile("use.c:SignSays", "C:\\NoxPost\\src\\client\\Network\\inform.c")
			nox_swprintf(&v63[0], v43, v47)
			nox_xxx_printCentered_445490(&v63[0])
			result = int32(v42 + 3)
		} else {
			v44 = strMan.GetStringInFile((*byte)(unsafe.Pointer(uintptr(a1+3))), "C:\\NoxPost\\src\\client\\Network\\inform.c")
			nox_xxx_printCentered_445490(v44)
		LABEL_57:
			result = int32(v42 + 3)
		}
		return result
	case 16:
		v49 = nox_server_teamTitle_418C20(int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 2)))))
		v9 = strMan.GetStringInFile("pickup.c:WrongTeam", "C:\\NoxPost\\src\\client\\Network\\inform.c")
		nox_swprintf(&v62[0], v9, v49)
		nox_xxx_printCentered_445490(&v62[0])
		return 6
	case 17:
		v45 = strMan.GetStringInFile("Noxworld.c:ErrChangedClass", "C:\\NoxPost\\src\\client\\Network\\inform.c")
		v6 = strMan.GetStringInFile("guiserv.c:Notice", "C:\\NoxPost\\src\\client\\Network\\inform.c")
		NewDialogWindow(nil, (*libc.WChar)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v6)))))), (*libc.WChar)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v45)))))), 33, nil, nil)
		return 2
	case 18:
		v24 = (*byte)(unsafe.Pointer(noxServer.getPlayerByID(int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 2)))))))
		if v24 == nil {
			return 6
		}
		v57 = int32(uintptr(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v24), 4704)))))
		v25 = strMan.GetStringInFile("objcoll.c:PlayerExited", "C:\\NoxPost\\src\\client\\Network\\inform.c")
		nox_swprintf(&v62[0], v25, v57)
		nox_xxx_printCentered_445490(&v62[0])
		return 6
	case 19:
		v26 = (*byte)(unsafe.Pointer(noxServer.getPlayerByID(int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 2)))))))
		if v26 != nil {
			v58 = int32(uintptr(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v26), 4704)))))
			v27 = strMan.GetStringInFile("objcoll.c:PlayerExitedWarp", "C:\\NoxPost\\src\\client\\Network\\inform.c")
			nox_swprintf(&v62[0], v27, v58)
			nox_xxx_printCentered_445490(&v62[0])
		}
		return 6
	case 20:
		v7 = (*byte)(unsafe.Pointer(noxServer.getPlayerByID(int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 2)))))))
		if v7 == nil {
			return 6
		}
		v48 = int32(uintptr(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v7), 4704)))))
		v8 = strMan.GetStringInFile("GeneralPrint:SecretFoundOther", "C:\\NoxPost\\src\\client\\Network\\inform.c")
		nox_swprintf(&v62[0], v8, v48)
		nox_xxx_printCentered_445490(&v62[0])
		return 6
	case 21:
		v59 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 2))))
		v28 = strMan.GetStringInFile("GeneralPrint:AdvanceToStage1", "C:\\NoxPost\\src\\client\\Network\\inform.c")
		nox_swprintf(&v62[0], v28, v59)
		v29 = strMan.GetStringInFile("use.c:SignSays", "C:\\NoxPost\\src\\client\\Network\\inform.c")
		nox_swprintf(&v63[0], v29, &v62[0])
		nox_xxx_printCentered_445490(&v63[0])
		return 6
	default:
		return result
	}
}
