package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	noxflags "github.com/noxworld-dev/opennox/v1/common/flags"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"unsafe"
)

func sub_4738D0() int32 {
	nox_xxx_bookHideMB_45ACA0(1)
	return 1
}
func nox_xxx_saveMakePlayerLocation_4DB600(a1 int32) int32 {
	var (
		v1 *byte
		v2 *float32
		v3 *uint32
		v4 float32
		v5 float32
		v6 int32
		v7 int32
		v8 int32
	)
	v1 = (*byte)(unsafe.Pointer(nox_common_playerInfoFromNum_417090(31)))
	if v1 == nil {
		return 0
	}
	v2 = (*float32)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v1))), unsafe.Sizeof(uint32(0))*514))))))
	if v2 == nil {
		return 0
	}
	v3 = (*uint32)(unsafe.Pointer(nox_xxx_newObjectByTypeID_4E3810(CString("SaveGameLocation"))))
	if v3 == nil {
		return 0
	}
	v4 = *(*float32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(float32(0))*14))
	v5 = *(*float32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(float32(0))*15))
	if a1 != 0 {
		v6 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 700))))
		v4 = *(*float32)(unsafe.Pointer(uintptr(v6 + 80)))
		v5 = *(*float32)(unsafe.Pointer(uintptr(v6 + 84)))
	}
	nox_xxx_createAt_4DAA50((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v3)))))), nil, v4, v5)
	nox_xxx_unitsNewAddToList_4DAC00()
	*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*11)) = *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), unsafe.Sizeof(uint32(0))*11)))
	v7 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), unsafe.Sizeof(uint32(0))*129))))
	if v7 != 0 {
		for {
			v8 = int32(*(*uint32)(unsafe.Pointer(uintptr(v7 + 512))))
			if int32(*(*uint8)(unsafe.Pointer(uintptr(v7 + 16))))&4 != 0 {
				nox_xxx_unitSetOwner_4EC290((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v3)))))), (*nox_object_t)(unsafe.Pointer(uintptr(v7))))
			}
			v7 = v8
			if v8 == 0 {
				break
			}
		}
	}
	return 1
}
func nox_xxx_monstersAllBelongToHost_4DB6A0() {
	var (
		v0 *byte
		v1 *byte
		v2 int32
		v3 int32
		v4 int32
		v5 int32
	)
	v0 = (*byte)(unsafe.Pointer(nox_common_playerInfoFromNum_417090(31)))
	v1 = v0
	if v0 != nil && *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v0))), unsafe.Sizeof(uint32(0))*514))) != 0 {
		if *memmap.PtrUint32(6112660, 1563124) == 0 {
			*memmap.PtrUint32(6112660, 1563124) = uint32(nox_xxx_getNameId_4E3AA0(CString("SaveGameLocation")))
		}
		v2 = int32(uintptr(unsafe.Pointer(noxServer.firstServerObject())))
		if v2 != 0 {
			for uint32(*(*uint16)(unsafe.Pointer(uintptr(v2 + 4)))) != *memmap.PtrUint32(6112660, 1563124) {
				v2 = int32(uintptr(unsafe.Pointer((*nox_object_t)(unsafe.Pointer(uintptr(v2))).Next())))
				if v2 == 0 {
					return
				}
			}
			v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 516))))
			if v3 != 0 {
				for {
					v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(v3 + 512))))
					nox_xxx_unitSetOwner_4EC290((*nox_object_t)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v1))), unsafe.Sizeof(uint32(0))*514)))))), (*nox_object_t)(unsafe.Pointer(uintptr(v3))))
					if int32(*(*uint8)(unsafe.Pointer(uintptr(v3 + 8))))&2 != 0 {
						if int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v3 + 748))) + 1440))))&128 != 0 {
							v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(v3 + 12))))
							*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v5))), 0)) = uint8(int8(v5 | 128))
							*(*uint32)(unsafe.Pointer(uintptr(v3 + 12))) = uint32(v5)
							nox_xxx_netReportAcquireCreature_4D91A0(int32(uint8(*(*byte)(unsafe.Add(unsafe.Pointer(v1), 2064)))), v3)
							nox_xxx_netMarkMinimapObject_417190(int32(uint8(*(*byte)(unsafe.Add(unsafe.Pointer(v1), 2064)))), v3, 1)
						}
					}
					v3 = v4
					if v4 == 0 {
						break
					}
				}
			}
			*(*uint32)(unsafe.Pointer(uintptr(v2 + 44))) = 0
			nox_xxx_delayedDeleteObject_4E5CC0((*nox_object_t)(unsafe.Pointer(uintptr(v2))))
		}
	}
}
func sub_4DB9C0() {
	var (
		v0 int32
		v1 int32
	)
	v0 = int32(uintptr(unsafe.Pointer(noxServer.firstServerObject())))
	if v0 != 0 {
		for {
			v1 = int32(uintptr(unsafe.Pointer((*nox_object_t)(unsafe.Pointer(uintptr(v0))).Next())))
			if nox_xxx_isUnit_4E5B50((*nox_object_t)(unsafe.Pointer(uintptr(v0)))) != 0 {
				nox_xxx_delayedDeleteObject_4E5CC0((*nox_object_t)(unsafe.Pointer(uintptr(v0))))
			}
			v0 = v1
			if v1 == 0 {
				break
			}
		}
	}
	var obj *nox_object_t = firstServerObjectUpdatable2()
	if obj != nil {
		for {
			{
				var v3 *nox_object_t = obj.Next()
				if sub_4E5B80(obj) != 0 {
					nox_xxx_delayedDeleteObject_4E5CC0(obj)
				}
				obj = v3
			}
			if obj == nil {
				break
			}
		}
	}
}
func nox_xxx_saveDoAutosaveMB_4DB370_savegame(a1 *byte) int32 {
	var (
		v1  *byte
		v2  *byte
		v3  int32
		v4  int32
		v5  *byte
		v6  *byte
		v7  *byte
		v8  uint32
		v9  uint32
		v10 *libc.WChar
		v12 *byte
		v13 *byte
		v14 [1024]byte
		v15 [1024]byte
	)
	sub_478000()
	nox_xxx_quickBarClose_4606B0()
	v1 = (*byte)(unsafe.Pointer(nox_common_playerInfoFromNum_417090(31)))
	v2 = v1
	if v1 == nil {
		return 0
	}
	v3 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v1))), unsafe.Sizeof(uint32(0))*514))))
	if v3 == 0 {
		return 0
	}
	if a1 == nil {
		return 0
	}
	v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(v3 + 16))))
	if (v4 & 0x8000) != 0 {
		return 0
	}
	if !nox_xxx_saveMakeFolder_0_4DB1D0() {
		return 0
	}
	if !nox_client_makeSaveDir_4DB540(CString("WORKING")) {
		return 0
	}
	v5 = nox_xxx_mapGetMapName_409B40()
	if !nox_client_makeSaveMapDir_4DB5A0(CString("WORKING"), (*byte)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v5))))))) {
		return 0
	}
	if nox_xxx_saveMakePlayerLocation_4DB600(int32(dword_5d4594_1563084)) == 0 {
		return 0
	}
	v13 = nox_xxx_mapGetMapName_409B40()
	v12 = nox_xxx_mapGetMapName_409B40()
	v6 = nox_fs_root()
	nox_sprintf(&v15[0], CString("%s\\Save\\%s\\%s\\%s.map"), v6, "WORKING", v12, v13)
	if nox_xxx_mapSaveMap_51E010(&v15[0], 0) == 0 {
		return 0
	}
	nox_xxx_monstersAllBelongToHost_4DB6A0()
	v7 = nox_fs_root()
	nox_sprintf(&v14[0], CString("%s\\Save\\%s\\Player.plr"), v7, "WORKING")
	v8 = *memmap.PtrUint32(0x85B3FC, 10980) & 0xFFFFFFF7
	*memmap.PtrUint32(0x85B3FC, 10980) &= 0xFFFFFFF7
	if *memmap.PtrUint32(6112660, 1563076) != 0 {
		v9 = v8
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v9))), 0)) = uint8(v8 | 8)
		*memmap.PtrUint32(0x85B3FC, 10980) = v9
	}
	*memmap.PtrUint8(0x85B3FC, 12257) = sub_450750()
	if nox_xxx_playerSaveToFile_41A140(&v14[0], int32(uint8(*(*byte)(unsafe.Add(unsafe.Pointer(v2), 2064))))) == 0 {
		return 0
	}
	if nox_xxx_mapSavePlayerDataMB_41A230(&v14[0]) == 0 {
		return 0
	}
	if libc.StrCmp(a1, CString("WORKING")) != 0 {
		v10 = strMan.GetStringInFile("AutoSaveComplete", "C:\\NoxPost\\src\\Server\\Xfer\\SaveGame\\SaveGame.c")
		nox_xxx_printToAll_4D9FD0(0, v10)
		if nox_client_copySave_4DC100(CString("WORKING"), a1) == 0 {
			return 0
		}
	}
	dword_5d4594_1563092 = 0
	dword_5d4594_1563088 = 0
	return 1
}
func nox_xxx_soloLoadGame_4DB7E0_savegame(a1 *byte) int32 {
	var (
		v1       *byte
		result   *byte
		v3       *byte
		v4       *libc.WChar
		v5       int32
		v6       *byte
		v7       *libc.WChar
		v8       [20]byte
		FileName [1024]byte
	)
	v1 = (*byte)(unsafe.Pointer(nox_common_playerInfoFromNum_417090(31)))
	if a1 == nil {
		return 0
	}
	if !noxflags.HasGame(noxflags.GameClient) || (func() *byte {
		result = (*byte)(unsafe.Pointer(uintptr(sub_4738D0())))
		return result
	}()) != nil {
		sub_4DB9C0()
		if libc.StrCmp(a1, CString("WORKING")) == 0 || (func() *byte {
			result = (*byte)(unsafe.Pointer(uintptr(nox_client_copySave_4DC100(a1, CString("WORKING")))))
			return result
		}()) != nil {
			v3 = nox_fs_root()
			nox_sprintf(&FileName[0], CString("%s\\Save\\%s\\Player.plr"), v3, memmap.PtrOff(0x587000, 199424))
			if nox_fs_access(&FileName[0], 0) == -1 {
				v4 = strMan.GetStringInFile("AutoSaveNotFound", "C:\\NoxPost\\src\\Server\\Xfer\\SaveGame\\SaveGame.c")
				nox_xxx_printToAll_4D9FD0(0, v4)
				result = nil
			} else {
				v5 = sub_41D090(&FileName[0])
				nox_server_SetFirstObjectScriptID_4E3C60(v5)
				nox_server_ResetObjectGIDs_4E3C70()
				nox_xxx_gameSetSwitchSolo_4DB220(1)
				nox_xxx_gameSetNoMPFlag_4DB230(1)
				result = nox_xxx_cliPlrInfoLoadFromFile_41A2E0(&FileName[0], 31)
				if result != nil {
					nox_xxx_cliPrepareGameplay1_460E60()
					nox_xxx_cliPrepareGameplay2_4721D0()
					nox_sprintf(&v8[0], CString("%s.map"), (*byte)(unsafe.Add(unsafe.Pointer(v1), 4760)))
					nox_xxx_gameSetMapPath_409D70(&v8[0])
					v6 = nox_fs_root()
					nox_sprintf((*byte)(memmap.PtrOff(6112660, 1559960)), CString("%s\\Save\\%s\\%s\\%s.map"), v6, memmap.PtrOff(0x587000, 199532), (*byte)(unsafe.Add(unsafe.Pointer(v1), 4760)), (*byte)(unsafe.Add(unsafe.Pointer(v1), 4760)))
					nox_xxx_mapLoad_4D2450(&v8[0])
					nox_xxx_cliPlayMapIntro_44E0B0(0)
					result = (*byte)(unsafe.Pointer(uintptr(nox_xxx_plrLoad_41A480(&FileName[0]))))
					if result != nil {
						nox_gui_console_Clear_450B70()
						sub_445450()
						nox_xxx_destroyEveryChatMB_528D60()
						v7 = strMan.GetStringInFile("GameLoaded", "C:\\NoxPost\\src\\Server\\Xfer\\SaveGame\\SaveGame.c")
						nox_xxx_printToAll_4D9FD0(0, v7)
						result = (*byte)(unsafe.Pointer(uintptr(1)))
					}
				}
			}
		}
	}
	return int32(uintptr(unsafe.Pointer(result)))
}
