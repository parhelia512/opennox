package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	noxflags "github.com/noxworld-dev/opennox/v1/common/flags"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"math"
	"unsafe"
)

var nox_console_playerWhoSent_823692 *nox_playerInfo = nil
var nox_client_consoleIsServer_823684 uint32 = 0

func nox_cmd_set_sysop(tokInd int32, tokCnt int32, tokens **libc.WChar) int32 {
	if tokCnt != 3 {
		return 0
	}
	nox_xxx_sysopSetPass_40A610(*(**libc.WChar)(unsafe.Add(unsafe.Pointer(tokens), unsafe.Sizeof((*libc.WChar)(nil))*uintptr(tokInd))))
	var s *libc.WChar = strMan.GetStringInFile("sysoppasswordset", "C:\\NoxPost\\src\\Client\\System\\parsecmd.c")
	nox_gui_console_Printf_450C00(uint8(int8(NOX_CONSOLE_RED)), s)
	return 1
}
func nox_cmd_show_game(tokInd int32, tokCnt int32, tokens **libc.WChar) int32 {
	var (
		v2  *byte
		v4  int32
		v5  int32
		v6  *libc.WChar
		v7  int16
		v8  *libc.WChar
		v9  *libc.WChar
		v11 *byte
		v12 *byte
		v13 int32
		v14 int32
		v15 *byte
		v16 int32
		v17 int32
		v18 [32]libc.WChar
	)
	v2 = sub_4165B0()
	if tokCnt != 2 {
		return 0
	}
	v4 = nox_common_playerInfoCount_416F40()
	if nox_common_getEngineFlag(nox_engine_flag(NOX_ENGINE_FLAG_DISABLE_GRAPHICS_RENDERING)) {
		v4--
	}
	v5 = int32(nox_client_getVersionBuild_409AC0())
	nox_gui_console_Printf_450C00(uint8(int8(NOX_CONSOLE_RED)), (*libc.WChar)(memmap.PtrOff(0x587000, 102952)), &nox_version_string_102944[0], v5)
	if noxflags.HasGame(noxflags.GameOnline) {
		v15 = nox_xxx_serverOptionsGetServername_40A4C0()
		v6 = strMan.GetStringInFile("Name", "C:\\NoxPost\\src\\Client\\System\\parsecmd.c")
		nox_gui_console_Printf_450C00(uint8(int8(NOX_CONSOLE_RED)), (*libc.WChar)(memmap.PtrOff(0x587000, 103028)), v6, v15)
		v7 = int16(uint16(nox_common_gameFlags_getVal_40A5B0()))
		v16 = int32(uintptr(unsafe.Pointer(nox_xxx_guiServerOptionsGetGametypeName_4573C0(v7))))
		v8 = strMan.GetStringInFile("Type", "C:\\NoxPost\\src\\Client\\System\\parsecmd.c")
		nox_gui_console_Printf_450C00(uint8(int8(NOX_CONSOLE_RED)), (*libc.WChar)(memmap.PtrOff(0x587000, 103088)), v8, v16)
		v17 = int32(sub_40A180(int16(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v2))), unsafe.Sizeof(uint16(0))*26))))))
		v14 = int32(uint16(nox_xxx_servGamedataGet_40A020(int16(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v2))), unsafe.Sizeof(uint16(0))*26)))))))
		v13 = nox_xxx_servGetPlrLimit_409FA0()
		v12 = nox_server_currentMapGetFilename_409B30()
		v9 = strMan.GetStringInFile("GameInfo", "C:\\NoxPost\\src\\Client\\System\\parsecmd.c")
		nox_gui_console_Printf_450C00(uint8(int8(NOX_CONSOLE_RED)), v9, v12, v4, v13, v14, v17)
		v11 = nox_net_ip2str(nox_net_in_addr(nox_xxx_net_getIP_554200(0)))
		nox_swprintf(&v18[0], libc.CWString("%S"), v11)
		nox_gui_console_Printf_450C00(uint8(int8(NOX_CONSOLE_RED)), (*libc.WChar)(memmap.PtrOff(0x587000, 103160)), &v18[0])
	}
	return 1
}
func sub_440A20(a1 *libc.WChar, _rest ...interface{}) {
	var va libc.ArgList
	va.Start(a1, _rest)
	nox_vswprintf((*libc.WChar)(memmap.PtrOff(6112660, 822660)), a1, va)
	nox_xxx_printCentered_445490((*libc.WChar)(memmap.PtrOff(6112660, 822660)))
}
func nox_cmd_set_cycle(tokInd int32, tokCnt int32, tokens **libc.WChar) int32 {
	var v3 *libc.WChar
	if tokCnt == 3 {
		if _nox_wcsicmp(*(**libc.WChar)(unsafe.Add(unsafe.Pointer(tokens), unsafe.Sizeof((*libc.WChar)(nil))*uintptr(tokInd))), libc.CWString("on")) == 0 {
			sub_4D0D90(1)
			v3 = strMan.GetStringInFile("MapCycleOn", "C:\\NoxPost\\src\\Client\\System\\parsecmd.c")
			nox_gui_console_Printf_450C00(uint8(int8(NOX_CONSOLE_RED)), v3)
			sub_4AD840()
			return 1
		}
		if _nox_wcsicmp(*(**libc.WChar)(unsafe.Add(unsafe.Pointer(tokens), unsafe.Sizeof((*libc.WChar)(nil))*uintptr(tokInd))), libc.CWString("off")) == 0 {
			sub_4D0D90(0)
			v3 = strMan.GetStringInFile("MapCycleOff", "C:\\NoxPost\\src\\Client\\System\\parsecmd.c")
			nox_gui_console_Printf_450C00(uint8(int8(NOX_CONSOLE_RED)), v3)
			sub_4AD840()
			return 1
		}
	}
	return 0
}
func nox_cmd_set_weapons(tokInd int32, tokCnt int32, tokens **libc.WChar) int32 {
	var (
		v3 *libc.WChar
		v5 *libc.WChar
	)
	if tokCnt == 3 {
		if _nox_wcsicmp(*(**libc.WChar)(unsafe.Add(unsafe.Pointer(tokens), unsafe.Sizeof((*libc.WChar)(nil))*uintptr(tokInd))), libc.CWString("on")) == 0 {
			sub_409E70(1)
			nox_server_gameSettingsUpdated_40A670()
			v5 = strMan.GetStringInFile("cmd_token:on", "C:\\NoxPost\\src\\Client\\System\\parsecmd.c")
			v3 = strMan.GetStringInFile("weapons", "C:\\NoxPost\\src\\Client\\System\\parsecmd.c")
			nox_gui_console_Printf_450C00(uint8(int8(NOX_CONSOLE_RED)), v3, v5)
			return 1
		}
		if _nox_wcsicmp(*(**libc.WChar)(unsafe.Add(unsafe.Pointer(tokens), unsafe.Sizeof((*libc.WChar)(nil))*uintptr(tokInd))), libc.CWString("off")) == 0 {
			sub_409EC0(1)
			nox_server_gameSettingsUpdated_40A670()
			v5 = strMan.GetStringInFile("cmd_token:off", "C:\\NoxPost\\src\\Client\\System\\parsecmd.c")
			v3 = strMan.GetStringInFile("weapons", "C:\\NoxPost\\src\\Client\\System\\parsecmd.c")
			nox_gui_console_Printf_450C00(uint8(int8(NOX_CONSOLE_RED)), v3, v5)
			return 1
		}
	}
	return 0
}
func nox_cmd_set_staffs(tokInd int32, tokCnt int32, tokens **libc.WChar) int32 {
	var (
		v3 *libc.WChar
		v5 *libc.WChar
	)
	if tokCnt == 3 {
		if _nox_wcsicmp(*(**libc.WChar)(unsafe.Add(unsafe.Pointer(tokens), unsafe.Sizeof((*libc.WChar)(nil))*uintptr(tokInd))), libc.CWString("on")) == 0 {
			sub_409E70(16)
			nox_server_gameSettingsUpdated_40A670()
			v5 = strMan.GetStringInFile("cmd_token:on", "C:\\NoxPost\\src\\Client\\System\\parsecmd.c")
			v3 = strMan.GetStringInFile("staffs", "C:\\NoxPost\\src\\Client\\System\\parsecmd.c")
			nox_gui_console_Printf_450C00(uint8(int8(NOX_CONSOLE_RED)), v3, v5)
			return 1
		}
		if _nox_wcsicmp(*(**libc.WChar)(unsafe.Add(unsafe.Pointer(tokens), unsafe.Sizeof((*libc.WChar)(nil))*uintptr(tokInd))), libc.CWString("off")) == 0 {
			sub_409EC0(16)
			nox_server_gameSettingsUpdated_40A670()
			v5 = strMan.GetStringInFile("cmd_token:off", "C:\\NoxPost\\src\\Client\\System\\parsecmd.c")
			v3 = strMan.GetStringInFile("staffs", "C:\\NoxPost\\src\\Client\\System\\parsecmd.c")
			nox_gui_console_Printf_450C00(uint8(int8(NOX_CONSOLE_RED)), v3, v5)
			return 1
		}
	}
	return 0
}
func nox_cmd_set_name(tokInd int32, tokCnt int32, tokens **libc.WChar) int32 {
	var (
		v3  int32
		v5  uint32
		v6  *byte
		v7  *byte
		v8  *byte
		v9  int8
		v10 int32
		v11 *libc.WChar
		v12 *uint32
		v13 [128]byte
		v14 [128]byte
	)
	v3 = tokInd
	if tokCnt < 3 {
		return 0
	}
	v13[0] = 0
	if tokInd < tokCnt {
		v12 = (*uint32)(unsafe.Pointer((**libc.WChar)(unsafe.Add(unsafe.Pointer(tokens), unsafe.Sizeof((*libc.WChar)(nil))*uintptr(tokInd)))))
		for {
			nox_sprintf(&v14[0], CString("%S"), *v12)
			v5 = uint32(libc.StrLen(&v14[0]) + 1)
			v6 = &v13[libc.StrLen(&v13[0])]
			libc.MemCpy(unsafe.Pointer(v6), unsafe.Pointer(&v14[0]), int((v5>>2)*4))
			v8 = &v14[(v5>>2)*4]
			v7 = (*byte)(unsafe.Add(unsafe.Pointer(v6), (v5>>2)*4))
			v9 = int8(uint8(v5))
			v10 = v3 + 1
			libc.MemCpy(unsafe.Pointer(v7), unsafe.Pointer(v8), int(int32(v9)&3))
			if v3+1 < tokCnt {
				*(*uint16)(unsafe.Pointer(&v13[libc.StrLen(&v13[0])])) = *memmap.PtrUint16(0x587000, 104484)
			}
			v3++
			v12 = (*uint32)(unsafe.Add(unsafe.Pointer(v12), unsafe.Sizeof(uint32(0))*1))
			if v10 >= tokCnt {
				break
			}
		}
		if v13[0] != 0 {
			nox_xxx_gameSetServername_40A440(&v13[0])
			v11 = strMan.GetStringInFile("setgamename", "C:\\NoxPost\\src\\Client\\System\\parsecmd.c")
			nox_gui_console_Printf_450C00(uint8(int8(NOX_CONSOLE_RED)), v11, &v13[0])
		}
	}
	return 1
}
func nox_cmd_set_mnstrs(tokInd int32, tokCnt int32, tokens **libc.WChar) int32 {
	var (
		v3 *libc.WChar
		v5 *libc.WChar
	)
	if tokCnt == 3 || tokCnt == 4 {
		if _nox_wcsicmp(*(**libc.WChar)(unsafe.Add(unsafe.Pointer(tokens), unsafe.Sizeof((*libc.WChar)(nil))*2)), *(**libc.WChar)(memmap.PtrOff(0x587000, 94468+4*6))) != 0 {
			if _nox_wcsicmp(*(**libc.WChar)(unsafe.Add(unsafe.Pointer(tokens), unsafe.Sizeof((*libc.WChar)(nil))*uintptr(tokInd))), libc.CWString("on")) == 0 {
				sub_409E70(4)
				nox_server_gameSettingsUpdated_40A670()
				v5 = strMan.GetStringInFile("cmd_token:on", "C:\\NoxPost\\src\\Client\\System\\parsecmd.c")
				v3 = strMan.GetStringInFile("monsters", "C:\\NoxPost\\src\\Client\\System\\parsecmd.c")
				nox_gui_console_Printf_450C00(uint8(int8(NOX_CONSOLE_RED)), v3, v5)
				return 1
			}
			if _nox_wcsicmp(*(**libc.WChar)(unsafe.Add(unsafe.Pointer(tokens), unsafe.Sizeof((*libc.WChar)(nil))*uintptr(tokInd))), libc.CWString("off")) == 0 {
				sub_409EC0(4)
				nox_server_gameSettingsUpdated_40A670()
				v5 = strMan.GetStringInFile("cmd_token:off", "C:\\NoxPost\\src\\Client\\System\\parsecmd.c")
				v3 = strMan.GetStringInFile("monsters", "C:\\NoxPost\\src\\Client\\System\\parsecmd.c")
				nox_gui_console_Printf_450C00(uint8(int8(NOX_CONSOLE_RED)), v3, v5)
				return 1
			}
		} else if tokCnt != 3 {
			if _nox_wcsicmp(*(**libc.WChar)(unsafe.Add(unsafe.Pointer(tokens), unsafe.Sizeof((*libc.WChar)(nil))*3)), libc.CWString("on")) == 0 {
				sub_409E70(8)
				nox_server_gameSettingsUpdated_40A670()
				v5 = strMan.GetStringInFile("cmd_token:on", "C:\\NoxPost\\src\\Client\\System\\parsecmd.c")
				v3 = strMan.GetStringInFile("monsterrespawn", "C:\\NoxPost\\src\\Client\\System\\parsecmd.c")
				nox_gui_console_Printf_450C00(uint8(int8(NOX_CONSOLE_RED)), v3, v5)
				return 1
			}
			if _nox_wcsicmp(*(**libc.WChar)(unsafe.Add(unsafe.Pointer(tokens), unsafe.Sizeof((*libc.WChar)(nil))*3)), libc.CWString("off")) == 0 {
				sub_409E70(8)
				nox_server_gameSettingsUpdated_40A670()
				v5 = strMan.GetStringInFile("cmd_token:off", "C:\\NoxPost\\src\\Client\\System\\parsecmd.c")
				v3 = strMan.GetStringInFile("monsterrespawn", "C:\\NoxPost\\src\\Client\\System\\parsecmd.c")
				nox_gui_console_Printf_450C00(uint8(int8(NOX_CONSOLE_RED)), v3, v5)
				return 1
			}
		}
	}
	return 0
}
func nox_cmd_set_spell(tokInd int32, tokCnt int32, tokens **libc.WChar) int32 {
	var (
		v3  *libc.WChar
		v4  int32
		v5  *byte
		v7  *libc.WChar
		v8  int32
		v9  int32
		v10 [100]byte
	)
	if tokCnt == 4 {
		if noxflags.HasGame(noxflags.GameModeChat) {
			v8 = int32(uintptr(unsafe.Pointer(*(**libc.WChar)(unsafe.Add(unsafe.Pointer(tokens), unsafe.Sizeof((*libc.WChar)(nil))*2)))))
			v3 = strMan.GetStringInFile("NotInChat", "C:\\NoxPost\\src\\Client\\System\\parsecmd.c")
			nox_gui_console_Printf_450C00(uint8(int8(NOX_CONSOLE_RED)), v3, v8)
			return 1
		}
		v4 = nox_xxx_spellByTitle_424960(*(**libc.WChar)(unsafe.Add(unsafe.Pointer(tokens), unsafe.Sizeof((*libc.WChar)(nil))*2)))
		if v4 != 0 || (func() bool {
			nox_sprintf(&v10[0], CString("%S"), *(**libc.WChar)(unsafe.Add(unsafe.Pointer(tokens), unsafe.Sizeof((*libc.WChar)(nil))*2)))
			return (func() int32 {
				v4 = things.SpellIndex(&v10[0])
				return v4
			}()) != 0
		}()) {
			if _nox_wcsicmp(*(**libc.WChar)(unsafe.Add(unsafe.Pointer(tokens), unsafe.Sizeof((*libc.WChar)(nil))*3)), libc.CWString("on")) == 0 {
				v5 = sub_4165B0()
				if (noxflags.HasGame(noxflags.GameModeFlagBall) || *(*byte)(unsafe.Add(unsafe.Pointer(v5), 52))&64 != 0) && v4 == 132 {
					return 1
				}
				if nox_xxx_spellIsEnabled_424B70(v4) {
					return 1
				}
				nox_xxx_spellEnable_424B90(v4)
				nox_server_gameSettingsUpdated_40A670()
				v8 = int32(uintptr(unsafe.Pointer(*(**libc.WChar)(unsafe.Add(unsafe.Pointer(tokens), unsafe.Sizeof((*libc.WChar)(nil))*2)))))
				v3 = strMan.GetStringInFile("spellenabled", "C:\\NoxPost\\src\\Client\\System\\parsecmd.c")
				nox_gui_console_Printf_450C00(uint8(int8(NOX_CONSOLE_RED)), v3, v8)
				return 1
			}
			if _nox_wcsicmp(*(**libc.WChar)(unsafe.Add(unsafe.Pointer(tokens), unsafe.Sizeof((*libc.WChar)(nil))*3)), libc.CWString("off")) == 0 {
				if !nox_xxx_spellIsEnabled_424B70(v4) {
					return 1
				}
				nox_xxx_spellDisable_424BB0(v4)
				nox_server_gameSettingsUpdated_40A670()
				v8 = int32(uintptr(unsafe.Pointer(*(**libc.WChar)(unsafe.Add(unsafe.Pointer(tokens), unsafe.Sizeof((*libc.WChar)(nil))*2)))))
				v3 = strMan.GetStringInFile("spelldisabled", "C:\\NoxPost\\src\\Client\\System\\parsecmd.c")
				nox_gui_console_Printf_450C00(uint8(int8(NOX_CONSOLE_RED)), v3, v8)
				return 1
			}
		} else {
			v9 = int32(uintptr(unsafe.Pointer(*(**libc.WChar)(unsafe.Add(unsafe.Pointer(tokens), unsafe.Sizeof((*libc.WChar)(nil))*2)))))
			v7 = strMan.GetStringInFile("invalidspell", "C:\\NoxPost\\src\\Client\\System\\parsecmd.c")
			nox_gui_console_Printf_450C00(uint8(int8(NOX_CONSOLE_RED)), v7, v9)
		}
	}
	return 0
}
func nox_xxx_deleteAllObjectsOfType_4E5DB0(a1 int32) {
	var v1 *nox_object_t = noxServer.firstServerObject()
	if v1 != nil {
		var v2 *nox_object_t
		for {
			{
				v2 = v1.Next()
				var v3 int32 = int32(*(*uint32)(unsafe.Pointer(&v1.field_126)))
				if v3 != 0 {
					var v4 int32
					for {
						v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(v3 + 496))))
						if int32(*(*uint16)(unsafe.Pointer(uintptr(v3 + 4)))) == a1 {
							nox_xxx_delayedDeleteObject_4E5CC0((*nox_object_t)(unsafe.Pointer(uintptr(v3))))
						}
						v3 = v4
						if v4 == 0 {
							break
						}
					}
				}
				if int32(v1.typ_ind) == a1 {
					nox_xxx_delayedDeleteObject_4E5CC0(v1)
				}
				v1 = v2
			}
			if v2 == nil {
				break
			}
		}
	}
}
func sub_4E3BC0(a1 int32) {
	dword_5d4594_1563664 ^= *(*uint32)(unsafe.Pointer(uintptr(a1 + 16)))
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 16))) = 0
}
func sub_415A60(a1 *libc.WChar) int32 {
	var (
		v1 *byte
		v2 int32
	)
	v1 = (*byte)(unsafe.Pointer(uintptr(sub_415960(a1))))
	if v1 != nil && (func() int32 {
		v2 = sub_415840(v1)
		return v2
	}()) != 0 {
		return int32(uintptr(unsafe.Pointer(nox_xxx_objectTypeByInd_4E3B70(v2))))
	}
	return 0
}
func nox_cmd_set_weapon(tokInd int32, tokCnt int32, tokens **libc.WChar) int32 {
	var (
		v3  *libc.WChar
		v4  *libc.WChar
		v6  *uint16
		v7  *libc.WChar
		v8  int32
		v9  int32
		v10 int32
		v11 [100]byte
	)
	if tokCnt == 4 {
		if noxflags.HasGame(noxflags.GameModeChat) {
			v8 = int32(uintptr(unsafe.Pointer(*(**libc.WChar)(unsafe.Add(unsafe.Pointer(tokens), unsafe.Sizeof((*libc.WChar)(nil))*2)))))
			v3 = strMan.GetStringInFile("NotInChat", "C:\\NoxPost\\src\\Client\\System\\parsecmd.c")
			nox_gui_console_Printf_450C00(uint8(int8(NOX_CONSOLE_RED)), v3, v8)
			return 1
		}
		if _nox_wcsicmp(*(**libc.WChar)(unsafe.Add(unsafe.Pointer(tokens), unsafe.Sizeof((*libc.WChar)(nil))*2)), *(**libc.WChar)(memmap.PtrOff(0x587000, 94468+4*6))) != 0 {
			v6 = (*uint16)(unsafe.Pointer(uintptr(sub_415A60(*(**libc.WChar)(unsafe.Add(unsafe.Pointer(tokens), unsafe.Sizeof((*libc.WChar)(nil))*2))))))
			if v6 != nil || (func() bool {
				nox_sprintf(&v11[0], CString("%S"), *(**libc.WChar)(unsafe.Add(unsafe.Pointer(tokens), unsafe.Sizeof((*libc.WChar)(nil))*2)))
				return (func() *uint16 {
					v6 = (*uint16)(unsafe.Pointer(uintptr(sub_415A30(&v11[0]))))
					return v6
				}()) != nil
			}()) {
				if _nox_wcsicmp(*(**libc.WChar)(unsafe.Add(unsafe.Pointer(tokens), unsafe.Sizeof((*libc.WChar)(nil))*3)), libc.CWString("on")) == 0 {
					if nox_xxx_getUnitDefDd10_4E3BA0(int32(*v6)) != 0 {
						return 1
					}
					sub_4E3BF0(int32(uintptr(unsafe.Pointer(v6))))
					nox_server_gameSettingsUpdated_40A670()
					v8 = int32(uintptr(unsafe.Pointer(*(**libc.WChar)(unsafe.Add(unsafe.Pointer(tokens), unsafe.Sizeof((*libc.WChar)(nil))*2)))))
					v3 = strMan.GetStringInFile("weaponEnabled", "C:\\NoxPost\\src\\Client\\System\\parsecmd.c")
					nox_gui_console_Printf_450C00(uint8(int8(NOX_CONSOLE_RED)), v3, v8)
					return 1
				}
				if _nox_wcsicmp(*(**libc.WChar)(unsafe.Add(unsafe.Pointer(tokens), unsafe.Sizeof((*libc.WChar)(nil))*3)), libc.CWString("off")) == 0 {
					if nox_xxx_getUnitDefDd10_4E3BA0(int32(*v6)) == 0 || nox_xxx_ammoCheck_415880(uint16(uintptr(unsafe.Pointer((*byte)(unsafe.Pointer(uintptr(*v6))))))) == 1 {
						return 1
					}
					sub_4E3BC0(int32(uintptr(unsafe.Pointer(v6))))
					nox_xxx_deleteAllObjectsOfType_4E5DB0(int32(*v6))
					nox_server_gameSettingsUpdated_40A670()
					v8 = int32(uintptr(unsafe.Pointer(*(**libc.WChar)(unsafe.Add(unsafe.Pointer(tokens), unsafe.Sizeof((*libc.WChar)(nil))*2)))))
					v3 = strMan.GetStringInFile("weaponDisabled", "C:\\NoxPost\\src\\Client\\System\\parsecmd.c")
					nox_gui_console_Printf_450C00(uint8(int8(NOX_CONSOLE_RED)), v3, v8)
					return 1
				}
			} else {
				v10 = int32(uintptr(unsafe.Pointer(*(**libc.WChar)(unsafe.Add(unsafe.Pointer(tokens), unsafe.Sizeof((*libc.WChar)(nil))*2)))))
				v7 = strMan.GetStringInFile("invalidweapon", "C:\\NoxPost\\src\\Client\\System\\parsecmd.c")
				nox_gui_console_Printf_450C00(uint8(int8(NOX_CONSOLE_RED)), v7, v10)
			}
		} else {
			if _nox_wcsicmp(*(**libc.WChar)(unsafe.Add(unsafe.Pointer(tokens), unsafe.Sizeof((*libc.WChar)(nil))*3)), libc.CWString("on")) == 0 {
				sub_409E70(2)
				v9 = int32(uintptr(unsafe.Pointer(*(**libc.WChar)(unsafe.Add(unsafe.Pointer(tokens), unsafe.Sizeof((*libc.WChar)(nil))*3)))))
				v4 = strMan.GetStringInFile("weaponsrespawn", "C:\\NoxPost\\src\\Client\\System\\parsecmd.c")
				nox_gui_console_Printf_450C00(uint8(int8(NOX_CONSOLE_RED)), v4, v9)
				sub_4AD840()
				return 1
			}
			if _nox_wcsicmp(*(**libc.WChar)(unsafe.Add(unsafe.Pointer(tokens), unsafe.Sizeof((*libc.WChar)(nil))*3)), libc.CWString("off")) == 0 {
				sub_409E70(2)
				v9 = int32(uintptr(unsafe.Pointer(*(**libc.WChar)(unsafe.Add(unsafe.Pointer(tokens), unsafe.Sizeof((*libc.WChar)(nil))*3)))))
				v4 = strMan.GetStringInFile("weaponsrespawn", "C:\\NoxPost\\src\\Client\\System\\parsecmd.c")
				nox_gui_console_Printf_450C00(uint8(int8(NOX_CONSOLE_RED)), v4, v9)
				sub_4AD840()
				return 1
			}
		}
	}
	return 0
}
func sub_415EF0(a1 *libc.WChar) int32 {
	var (
		v1 *byte
		v2 int32
	)
	v1 = (*byte)(unsafe.Pointer(uintptr(sub_415DA0(a1))))
	if v1 != nil && (func() int32 {
		v2 = sub_415CD0(v1)
		return v2
	}()) != 0 {
		return int32(uintptr(unsafe.Pointer(nox_xxx_objectTypeByInd_4E3B70(v2))))
	}
	return 0
}
func nox_cmd_set_armor(tokInd int32, tokCnt int32, tokens **libc.WChar) int32 {
	var (
		v3 *libc.WChar
		v4 *uint16
		v6 *libc.WChar
		v7 int32
		v8 int32
		v9 [100]byte
	)
	if tokCnt == 4 {
		if noxflags.HasGame(noxflags.GameModeChat) {
			v7 = int32(uintptr(unsafe.Pointer(*(**libc.WChar)(unsafe.Add(unsafe.Pointer(tokens), unsafe.Sizeof((*libc.WChar)(nil))*2)))))
			v3 = strMan.GetStringInFile("NotInChat", "C:\\NoxPost\\src\\Client\\System\\parsecmd.c")
			nox_gui_console_Printf_450C00(uint8(int8(NOX_CONSOLE_RED)), v3, v7)
			return 1
		}
		v4 = (*uint16)(unsafe.Pointer(uintptr(sub_415EF0(*(**libc.WChar)(unsafe.Add(unsafe.Pointer(tokens), unsafe.Sizeof((*libc.WChar)(nil))*2))))))
		if v4 != nil || (func() bool {
			nox_sprintf(&v9[0], CString("%S"), *(**libc.WChar)(unsafe.Add(unsafe.Pointer(tokens), unsafe.Sizeof((*libc.WChar)(nil))*2)))
			return (func() *uint16 {
				v4 = (*uint16)(unsafe.Pointer(uintptr(sub_415EC0(&v9[0]))))
				return v4
			}()) != nil
		}()) {
			if _nox_wcsicmp(*(**libc.WChar)(unsafe.Add(unsafe.Pointer(tokens), unsafe.Sizeof((*libc.WChar)(nil))*3)), libc.CWString("on")) == 0 {
				if nox_xxx_getUnitDefDd10_4E3BA0(int32(*v4)) == 0 {
					sub_4E3BF0(int32(uintptr(unsafe.Pointer(v4))))
					nox_server_gameSettingsUpdated_40A670()
					v7 = int32(uintptr(unsafe.Pointer(*(**libc.WChar)(unsafe.Add(unsafe.Pointer(tokens), unsafe.Sizeof((*libc.WChar)(nil))*2)))))
					v3 = strMan.GetStringInFile("armorEnabled", "C:\\NoxPost\\src\\Client\\System\\parsecmd.c")
					nox_gui_console_Printf_450C00(uint8(int8(NOX_CONSOLE_RED)), v3, v7)
					return 1
				}
				return 1
			}
			if _nox_wcsicmp(*(**libc.WChar)(unsafe.Add(unsafe.Pointer(tokens), unsafe.Sizeof((*libc.WChar)(nil))*3)), libc.CWString("off")) == 0 {
				if nox_xxx_getUnitDefDd10_4E3BA0(int32(*v4)) != 0 {
					sub_4E3BC0(int32(uintptr(unsafe.Pointer(v4))))
					nox_xxx_deleteAllObjectsOfType_4E5DB0(int32(*v4))
					nox_server_gameSettingsUpdated_40A670()
					v7 = int32(uintptr(unsafe.Pointer(*(**libc.WChar)(unsafe.Add(unsafe.Pointer(tokens), unsafe.Sizeof((*libc.WChar)(nil))*2)))))
					v3 = strMan.GetStringInFile("armorDisabled", "C:\\NoxPost\\src\\Client\\System\\parsecmd.c")
					nox_gui_console_Printf_450C00(uint8(int8(NOX_CONSOLE_RED)), v3, v7)
					return 1
				}
				return 1
			}
		} else {
			v8 = int32(uintptr(unsafe.Pointer(*(**libc.WChar)(unsafe.Add(unsafe.Pointer(tokens), unsafe.Sizeof((*libc.WChar)(nil))*2)))))
			v6 = strMan.GetStringInFile("invalidarmor", "C:\\NoxPost\\src\\Client\\System\\parsecmd.c")
			nox_gui_console_Printf_450C00(uint8(int8(NOX_CONSOLE_RED)), v6, v8)
		}
	}
	return 0
}
func nox_cmd_set_staff(tokInd int32, tokCnt int32, tokens **libc.WChar) int32 {
	var (
		v3 int32
		v4 *libc.WChar
		v6 *libc.WChar
		v7 int32
		v8 int32
		v9 [128]byte
	)
	if tokCnt == 4 {
		nox_sprintf(&v9[0], CString("%S"), *(**libc.WChar)(unsafe.Add(unsafe.Pointer(tokens), unsafe.Sizeof((*libc.WChar)(nil))*2)))
		v3 = int32(uintptr(unsafe.Pointer(nox_xxx_objectTypeByID_4E3B60(&v9[0]))))
		if v3 != 0 {
			if _nox_wcsicmp(*(**libc.WChar)(unsafe.Add(unsafe.Pointer(tokens), unsafe.Sizeof((*libc.WChar)(nil))*3)), libc.CWString("on")) == 0 {
				sub_4E3BF0(v3)
				v7 = int32(uintptr(unsafe.Pointer(*(**libc.WChar)(unsafe.Add(unsafe.Pointer(tokens), unsafe.Sizeof((*libc.WChar)(nil))*2)))))
				v4 = strMan.GetStringInFile("staffEnabled", "C:\\NoxPost\\src\\Client\\System\\parsecmd.c")
				nox_gui_console_Printf_450C00(uint8(int8(NOX_CONSOLE_RED)), v4, v7)
				return 1
			}
			if _nox_wcsicmp(*(**libc.WChar)(unsafe.Add(unsafe.Pointer(tokens), unsafe.Sizeof((*libc.WChar)(nil))*3)), libc.CWString("off")) == 0 {
				sub_4E3BC0(v3)
				v7 = int32(uintptr(unsafe.Pointer(*(**libc.WChar)(unsafe.Add(unsafe.Pointer(tokens), unsafe.Sizeof((*libc.WChar)(nil))*2)))))
				v4 = strMan.GetStringInFile("staffDisabled", "C:\\NoxPost\\src\\Client\\System\\parsecmd.c")
				nox_gui_console_Printf_450C00(uint8(int8(NOX_CONSOLE_RED)), v4, v7)
				return 1
			}
		} else {
			v8 = int32(uintptr(unsafe.Pointer(*(**libc.WChar)(unsafe.Add(unsafe.Pointer(tokens), unsafe.Sizeof((*libc.WChar)(nil))*2)))))
			v6 = strMan.GetStringInFile("invalidstaff", "C:\\NoxPost\\src\\Client\\System\\parsecmd.c")
			nox_gui_console_Printf_450C00(uint8(int8(NOX_CONSOLE_RED)), v6, v8)
		}
	}
	return 0
}
func nox_cmd_ban(tokInd int32, tokCnt int32, tokens **libc.WChar) int32 {
	var (
		v4 *libc.WChar
		v5 *byte
		v6 *byte
		v7 *libc.WChar
		v8 *libc.WChar
		v9 int32
	)
	if tokCnt != 2 {
		return 0
	}
	v4 = *(**libc.WChar)(unsafe.Add(unsafe.Pointer(tokens), unsafe.Sizeof((*libc.WChar)(nil))*uintptr(tokInd)))
	v5 = nox_xxx_playerByName_4170D0(*(**libc.WChar)(unsafe.Add(unsafe.Pointer(tokens), unsafe.Sizeof((*libc.WChar)(nil))*uintptr(tokInd))))
	v6 = v5
	if v5 != nil {
		if *(*byte)(unsafe.Add(unsafe.Pointer(v5), 2064)) == 31 {
			v8 = strMan.GetStringInFile("cantbanyourself", "C:\\NoxPost\\src\\Client\\System\\parsecmd.c")
			sub_440A20(v8)
			return 1
		}
		if noxflags.HasGame(noxflags.GameModeQuest) {
			sub_4DCFB0((*nox_object_t)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v6))), unsafe.Sizeof(uint32(0))*514)))))))
		} else {
			nox_xxx_playerDisconnByPlrID_4DEB00(int32(uint8(*(*byte)(unsafe.Add(unsafe.Pointer(v6), 2064)))))
		}
		sub_416770(0, v4, (*byte)(unsafe.Add(unsafe.Pointer(v6), 2112)))
		v9 = int32(uintptr(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v6), 4704)))))
		v7 = strMan.GetStringInFile("banned", "C:\\NoxPost\\src\\Client\\System\\parsecmd.c")
	} else {
		sub_416770(0, v4, nil)
		v9 = int32(uintptr(unsafe.Pointer(v4)))
		v7 = strMan.GetStringInFile("banDisallow", "C:\\NoxPost\\src\\Client\\System\\parsecmd.c")
	}
	sub_440A20(v7, v9)
	return 1
}
func nox_cmd_allow_user(tokInd int32, tokCnt int32, tokens **libc.WChar) int32 {
	var s *libc.WChar = strMan.GetStringInFile("notyetimplemented", "C:\\NoxPost\\src\\Client\\System\\parsecmd.c")
	nox_gui_console_Printf_450C00(uint8(int8(NOX_CONSOLE_RED)), s)
	return 1
}
func nox_cmd_allow_ip(tokInd int32, tokCnt int32, tokens **libc.WChar) int32 {
	var s *libc.WChar = strMan.GetStringInFile("notyetimplemented", "C:\\NoxPost\\src\\Client\\System\\parsecmd.c")
	nox_gui_console_Printf_450C00(uint8(int8(NOX_CONSOLE_RED)), s)
	return 1
}
func nox_cmd_kick(tokInd int32, tokCnt int32, tokens **libc.WChar) int32 {
	var (
		v4 *byte
		v5 *byte
		v6 *libc.WChar
		v7 *libc.WChar
	)
	if tokCnt != 2 {
		return 0
	}
	v4 = nox_xxx_playerByName_4170D0(*(**libc.WChar)(unsafe.Add(unsafe.Pointer(tokens), unsafe.Sizeof((*libc.WChar)(nil))*uintptr(tokInd))))
	v5 = v4
	if v4 == nil {
		return 1
	}
	if *(*byte)(unsafe.Add(unsafe.Pointer(v4), 2064)) == 31 {
		v7 = strMan.GetStringInFile("cantkickyourself", "C:\\NoxPost\\src\\Client\\System\\parsecmd.c")
		sub_440A20(v7)
		return 1
	}
	if noxflags.HasGame(noxflags.GameModeQuest) {
		sub_4DCFB0((*nox_object_t)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v5))), unsafe.Sizeof(uint32(0))*514)))))))
	} else {
		nox_xxx_playerCallDisconnect_4DEAB0(int32(uint8(*(*byte)(unsafe.Add(unsafe.Pointer(v5), 2064)))), 4)
		v6 = strMan.GetStringInFile("kicked", "C:\\NoxPost\\src\\Client\\System\\parsecmd.c")
		sub_440A20(v6, (*byte)(unsafe.Add(unsafe.Pointer(v5), 4704)))
	}
	return 1
}
func nox_cmd_set_players(tokInd int32, tokCnt int32, tokens **libc.WChar) int32 {
	var (
		v3 int32
		v5 *libc.WChar
		v6 int32
		v7 *libc.WChar
	)
	v3 = 0
	if tokCnt != 3 {
		return 0
	}
	v5 = *(**libc.WChar)(unsafe.Add(unsafe.Pointer(tokens), unsafe.Sizeof((*libc.WChar)(nil))*2))
	if v5 != nil {
		v6 = nox_wcstol(v5, nil, 10)
		v3 = v6
		if v6 >= 0 {
			if v6 > 999 {
				v3 = 999
			}
		} else {
			v3 = 0
		}
		if nox_xxx_servGetPlrLimit_409FA0() == v3 {
			return 1
		}
		nox_xxx_servSetPlrLimit_409F80(v3)
		sub_455800()
	}
	v7 = strMan.GetStringInFile("playersset", "C:\\NoxPost\\src\\Client\\System\\parsecmd.c")
	sub_440A20(v7, v3)
	return 1
}
func nox_cmd_set_spellpts(tokInd int32, tokCnt int32, tokens **libc.WChar) int32 {
	var s *libc.WChar = strMan.GetStringInFile("notyetimplemented", "C:\\NoxPost\\src\\Client\\System\\parsecmd.c")
	nox_gui_console_Printf_450C00(uint8(int8(NOX_CONSOLE_RED)), s)
	return 1
}
func nox_cmd_list_users(tokInd int32, tokCnt int32, tokens **libc.WChar) int32 {
	var (
		v0 *libc.WChar
		i  *byte
		v2 *libc.WChar
		v3 *libc.WChar
		v5 [128]libc.WChar
	)
	v0 = strMan.GetStringInFile("userslist", "C:\\NoxPost\\src\\Client\\System\\parsecmd.c")
	nox_gui_console_Printf_450C00(uint8(int8(NOX_CONSOLE_RED)), v0)
	for i = (*byte)(unsafe.Pointer(nox_common_playerInfoGetFirst_416EA0())); i != nil; i = (*byte)(unsafe.Pointer(nox_common_playerInfoGetNext_416EE0((*nox_playerInfo)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(i))))))))) {
		v5[0] = 0
		nox_wcscat(&v5[0], (*libc.WChar)(unsafe.Add(unsafe.Pointer((*libc.WChar)(unsafe.Pointer(i))), unsafe.Sizeof(libc.WChar(0))*2352)))
		if nox_client_consoleIsServer_823684 != 0 && *(*byte)(unsafe.Add(unsafe.Pointer(i), 3680))&4 != 0 {
			nox_wcscat(&v5[0], libc.CWString(", "))
			v2 = strMan.GetStringInFile("SysMuted", "C:\\NoxPost\\src\\Client\\System\\parsecmd.c")
			nox_wcscat(&v5[0], v2)
		}
		if *(*byte)(unsafe.Add(unsafe.Pointer(i), 3680))&8 != 0 {
			nox_wcscat(&v5[0], libc.CWString(", "))
			v3 = strMan.GetStringInFile("ClientMuted", "C:\\NoxPost\\src\\Client\\System\\parsecmd.c")
			nox_wcscat(&v5[0], v3)
		}
		nox_gui_console_Printf_450C00(uint8(int8(NOX_CONSOLE_RED)), (*libc.WChar)(memmap.PtrOff(0x587000, 106604)), &v5[0])
	}
	return 1
}
func sub_57A0F0(a1 *libc.WChar) int32 {
	if !noxflags.HasGame(noxflags.GameClient) {
		return 0
	}
	if a1 == nil {
		return 0
	}
	var v1 *byte = nox_xxx_playerByName_4170D0(a1)
	if v1 == nil {
		return 0
	}
	nox_xxx_playerUnsetStatus_417530((*nox_playerInfo)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v1)))))), 8)
	return 1
}
func sub_57A130(a1 *libc.WChar) int32 {
	if a1 == nil {
		return 0
	}
	var v1 *byte = nox_xxx_playerByName_4170D0(a1)
	if v1 == nil {
		return 0
	}
	nox_xxx_playerUnsetStatus_417530((*nox_playerInfo)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v1)))))), 4)
	return 1
}
func nox_cmd_unmute(tokInd int32, tokCnt int32, tokens **libc.WChar) int32 {
	var (
		v3 **libc.WChar
		v4 int32
		v5 int32
		v6 *libc.WChar
		v8 *libc.WChar
	)
	if tokCnt < 2 || tokCnt > 3 {
		return 0
	}
	v3 = (**libc.WChar)(unsafe.Add(unsafe.Pointer(tokens), unsafe.Sizeof((*libc.WChar)(nil))*uintptr(tokInd)))
	if nox_client_consoleIsServer_823684 == 0 {
		v5 = sub_57A0F0(*v3)
	} else {
		if _nox_wcsicmp(*(**libc.WChar)(unsafe.Add(unsafe.Pointer(tokens), unsafe.Sizeof((*libc.WChar)(nil))*uintptr(tokInd))), *(**libc.WChar)(memmap.PtrOff(0x587000, 94468+4*7))) != 0 {
			v5 = sub_57A0F0(*v3)
		} else {
			v4 = tokInd + 1
			if tokInd+1 != tokCnt-1 {
				return 0
			}
			v3 = (**libc.WChar)(unsafe.Add(unsafe.Pointer(tokens), unsafe.Sizeof((*libc.WChar)(nil))*uintptr(v4)))
			v5 = sub_57A130(*v3)
		}
	}
	v8 = *v3
	if v5 != 0 {
		v6 = strMan.GetStringInFile("UnMuted", "C:\\NoxPost\\src\\Client\\System\\parsecmd.c")
	} else {
		v6 = strMan.GetStringInFile("UserNotFound", "C:\\NoxPost\\src\\Client\\System\\parsecmd.c")
	}
	nox_gui_console_Printf_450C00(uint8(int8(NOX_CONSOLE_RED)), v6, v8)
	return 1
}
func sub_57A080(a1 *libc.WChar) int32 {
	if !noxflags.HasGame(noxflags.GameClient) {
		return 0
	}
	if a1 == nil {
		return 0
	}
	var v1 *byte = nox_xxx_playerByName_4170D0(a1)
	if v1 == nil || *(*byte)(unsafe.Add(unsafe.Pointer(v1), 2064)) == 31 {
		return 0
	}
	nox_xxx_netNeedTimestampStatus_4174F0(int32(uintptr(unsafe.Pointer(v1))), 8)
	return 1
}
func sub_57A0C0(a1 *libc.WChar) int32 {
	if a1 == nil {
		return 0
	}
	var v1 *byte = nox_xxx_playerByName_4170D0(a1)
	if v1 == nil {
		return 0
	}
	nox_xxx_netNeedTimestampStatus_4174F0(int32(uintptr(unsafe.Pointer(v1))), 4)
	return 1
}
func nox_cmd_mute(tokInd int32, tokCnt int32, tokens **libc.WChar) int32 {
	var (
		v3 **libc.WChar
		v4 int32
		v5 int32
		v6 *libc.WChar
		v8 *libc.WChar
	)
	if tokCnt < 2 || tokCnt > 3 {
		return 0
	}
	v3 = (**libc.WChar)(unsafe.Add(unsafe.Pointer(tokens), unsafe.Sizeof((*libc.WChar)(nil))*uintptr(tokInd)))
	if nox_client_consoleIsServer_823684 == 0 {
		v5 = sub_57A080(*v3)
	} else {
		if _nox_wcsicmp(*(**libc.WChar)(unsafe.Add(unsafe.Pointer(tokens), unsafe.Sizeof((*libc.WChar)(nil))*uintptr(tokInd))), *(**libc.WChar)(memmap.PtrOff(0x587000, 94468+4*7))) != 0 {
			v5 = sub_57A080(*v3)
		} else {
			v4 = tokInd + 1
			if tokInd+1 != tokCnt-1 {
				return 0
			}
			v3 = (**libc.WChar)(unsafe.Add(unsafe.Pointer(tokens), unsafe.Sizeof((*libc.WChar)(nil))*uintptr(v4)))
			v5 = sub_57A0C0(*v3)
		}
	}
	v8 = *v3
	if v5 != 0 {
		v6 = strMan.GetStringInFile("Muted", "C:\\NoxPost\\src\\Client\\System\\parsecmd.c")
	} else {
		v6 = strMan.GetStringInFile("UserNotFound", "C:\\NoxPost\\src\\Client\\System\\parsecmd.c")
	}
	nox_gui_console_Printf_450C00(uint8(int8(NOX_CONSOLE_RED)), v6, v8)
	return 1
}
func nox_cmd_exec_rul(tokInd int32, tokCnt int32, tokens **libc.WChar) int32 {
	if tokCnt != 2 {
		return 0
	}
	var buf [128]libc.WChar
	nox_wcscpy(&buf[0], *(**libc.WChar)(unsafe.Add(unsafe.Pointer(tokens), unsafe.Sizeof((*libc.WChar)(nil))*1)))
	if nox_wcschr(&buf[0], 46) == nil {
		nox_wcscat(&buf[0], libc.CWString(".rul"))
	}
	var s *libc.WChar = strMan.GetStringInFile("ExecutingRul", "C:\\NoxPost\\src\\Client\\System\\parsecmd.c")
	nox_gui_console_Printf_450C00(uint8(int8(NOX_CONSOLE_RED)), s, &buf[0])
	nox_xxx_doExecrul_4438A0(int32(uintptr(unsafe.Pointer(&buf[0]))))
	return 1
}
func nox_xxx_serverHandleClientConsole_443E90(pl *nox_playerInfo, a2 int8, a3 *libc.WChar) int32 {
	var (
		v3     int32
		result int32
		v5     int32
		v6     *libc.WChar
		v7     int32
		v8     int32
		v9     *libc.WChar
		v10    *libc.WChar
		v11    int32
		v12    *libc.WChar
		v13    *libc.WChar
		v14    *byte
		v15    int32
		v16    *libc.WChar
		v17    *byte
		v18    *libc.WChar
		v19    int32
		v20    int32
		v21    *libc.WChar
		v22    [128]byte
	)
	if pl == nil || pl.playerUnit == nil {
		return 0
	}
	if a3 != nil {
		nox_wcscpy((*libc.WChar)(memmap.PtrOff(6112660, 818228)), a3)
	} else {
		*memmap.PtrUint16(6112660, 818228) = 0
	}
	v3 = int32(uintptr(unsafe.Pointer(pl)))
	nox_console_playerWhoSent_823692 = pl
	if int32(a2) != 4 && int32(a2) != 5 && int32(a2) != 0 {
		if noxflags.HasGame(noxflags.GameFlag15 | noxflags.GameFlag16) {
			return 1
		}
	}
	switch a2 {
	case 0:
		if noxflags.HasGame(noxflags.GameFlag4) || noxflags.HasGame(noxflags.GameModeQuest) || int32(*(*uint8)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(pl))) + 3680))))&1 != 0 {
			nox_console_playerWhoSent_823692 = nil
			return 1
		}
		v5 = bool2int(int32(*memmap.PtrInt16(6112660, 818228)) == -4083 && int32(*memmap.PtrInt16(6112660, 818230)) == -3923 && int32(*memmap.PtrUint16(6112660, 818232)) == 0)
		if nox_xxx_playerGoObserver_4E6860(pl, v5, 0) != 1 {
			nox_console_playerWhoSent_823692 = nil
			return 1
		}
		v6 = strMan.GetStringInFile("set", "C:\\NoxPost\\src\\Client\\System\\parsecmd.c")
		v7 = nox_xxx_gamePlayIsAnyPlayers_40A8A0()
		if v7 != 0 {
			if v5 == 0 {
				nox_xxx_netNeedTimestampStatus_4174F0(int32(uintptr(unsafe.Pointer(pl))), 256)
			}
			if noxflags.HasGame(noxflags.GameModeElimination) && sub_40A770() == 1 {
				sub_5095E0()
			}
		}
		v8 = int32(uintptr(unsafe.Pointer(pl.playerUnit)))
		if v8 != 0 {
			nox_xxx_netChangeTeamMb_419570(v8+48, int32(pl.netCode))
		}
		v21 = v6
		v9 = strMan.GetStringInFile("observermode", "C:\\NoxPost\\src\\Client\\System\\parsecmd.c")
		nox_gui_console_Printf_450C00(uint8(int8(NOX_CONSOLE_RED)), v9, v21)
		nox_console_playerWhoSent_823692 = nil
		return 1
	case 1:
		if sub_4D12A0(int32(pl.playerInd)) == 0 && uint32(pl.playerInd) != nox_player_netCode_85319C && !noxflags.HasGame(noxflags.GameModeCoop) {
			nox_console_playerWhoSent_823692 = nil
			return 1
		}
		nox_wcstok((*libc.WChar)(memmap.PtrOff(6112660, 818228)), libc.CWString(" "))
		v10 = nox_wcstok(nil, libc.CWString(" "))
		nox_sprintf(&v22[0], CString("%S"), v10)
		v11 = nox_script_indexByEvent(&v22[0])
		if v11 != -1 && nox_console_playerWhoSent_823692 != nil {
			v12 = strMan.GetStringInFile("ExecutingFunction", "C:\\NoxPost\\src\\Client\\System\\parsecmd.c")
			nox_gui_console_Printf_450C00(uint8(int8(NOX_CONSOLE_RED)), v12, v10)
			nox_script_callByIndex_507310(v11, unsafe.Pointer(nox_console_playerWhoSent_823692.playerUnit), unsafe.Pointer(nox_console_playerWhoSent_823692.playerUnit))
			nox_console_playerWhoSent_823692 = nil
		} else {
			v21 = v10
			v9 = strMan.GetStringInFile("InvalidFunction", "C:\\NoxPost\\src\\Client\\System\\parsecmd.c")
			nox_gui_console_Printf_450C00(uint8(int8(NOX_CONSOLE_RED)), v9, v21)
			nox_console_playerWhoSent_823692 = nil
		}
		return 1
	case 2:
		if sub_4D12A0(int32(pl.playerInd)) == 0 && uint32(pl.playerInd) != nox_player_netCode_85319C && !noxflags.HasGame(noxflags.GameModeCoop) {
			nox_console_playerWhoSent_823692 = nil
			return 1
		}
		v19 = int32(uintptr(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(nox_console_playerWhoSent_823692))), 4704)))))
		v13 = strMan.GetStringInFile("RemoteSysop", "C:\\NoxPost\\src\\Client\\System\\parsecmd.c")
		nox_gui_console_Printf_450C00(uint8(int8(NOX_CONSOLE_RED)), v13, v19, a3)
		nox_server_parseCmdText_443C80(a3, 0)
		nox_console_playerWhoSent_823692 = nil
		return 1
	case 3:
		nox_xxx_printToAll_4D9FD0(0, a3)
		nox_console_playerWhoSent_823692 = nil
		return 1
	case 4:
		if (int32(*(*uint8)(unsafe.Pointer(uintptr(v3 + 3680))))&1) == 0 && !nox_common_getEngineFlag(nox_engine_flag(NOX_ENGINE_FLAG_REPLAY_READ)) {
			if noxflags.HasGame(noxflags.GameHost) {
				v16 = strMan.GetStringInFile("notinobserver", "C:\\NoxPost\\src\\Client\\System\\parsecmd.c")
				nox_gui_console_Printf_450C00(uint8(int8(NOX_CONSOLE_RED)), v16)
				nox_console_playerWhoSent_823692 = nil
				return 1
			}
			nox_console_playerWhoSent_823692 = nil
			return 1
		}
		if *a3 == 0 {
			nox_xxx_playerCameraUnlock_4E6040(int32(*(*uint32)(unsafe.Pointer(uintptr(v3 + 2056)))))
			nox_console_playerWhoSent_823692 = nil
			return 1
		}
		v17 = (*byte)(unsafe.Pointer(nox_common_playerInfoGetFirst_416EA0()))
		if v17 == nil {
			nox_console_playerWhoSent_823692 = nil
			return 1
		}
		for {
			if _nox_wcsicmp(a3, (*libc.WChar)(unsafe.Add(unsafe.Pointer((*libc.WChar)(unsafe.Pointer(v17))), unsafe.Sizeof(libc.WChar(0))*2352))) == 0 {
				nox_xxx_playerCameraFollow_4E6060(int32(uintptr(unsafe.Pointer(nox_console_playerWhoSent_823692.playerUnit))), int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v17))), unsafe.Sizeof(uint32(0))*514)))))
			}
			v17 = (*byte)(unsafe.Pointer(nox_common_playerInfoGetNext_416EE0((*nox_playerInfo)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v17)))))))))
			if v17 == nil {
				break
			}
		}
		nox_console_playerWhoSent_823692 = nil
		return 1
	case 5:
		nox_xxx_printToAll_4D9FD0(16, a3)
		v14 = (*byte)(unsafe.Pointer(nox_common_playerInfoGetFirst_416EA0()))
		if v14 == nil {
			nox_console_playerWhoSent_823692 = nil
			return 1
		}
		for {
			v15 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v14))), unsafe.Sizeof(uint32(0))*514))))
			if v15 != 0 {
				nox_xxx_aud_501960(902, (*nox_object_t)(unsafe.Pointer(uintptr(v15))), 0, 0)
			}
			v14 = (*byte)(unsafe.Pointer(nox_common_playerInfoGetNext_416EE0((*nox_playerInfo)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v14)))))))))
			if v14 == nil {
				break
			}
		}
		nox_console_playerWhoSent_823692 = nil
		return 1
	default:
		v20 = v3 + 4704
		v18 = strMan.GetStringInFile("invalidattempt", "C:\\NoxPost\\src\\Client\\System\\parsecmd.c")
		nox_gui_console_Printf_450C00(uint8(int8(NOX_CONSOLE_RED)), v18, v20, a3)
		nox_console_playerWhoSent_823692 = nil
		return 1
	}
	return result
}
func nox_cmd_offonly1(tokInd int32, tokCnt int32, tokens **libc.WChar) int32 {
	if tokCnt != 2 {
		return 0
	}
	nox_xxx_wndGuiTeamCreate_4185B0()
	return 1
}
func nox_cmd_offonly2(tokInd int32, tokCnt int32, tokens **libc.WChar) int32 {
	var (
		result int32
		v4     *byte
		v5     *libc.WChar
		v6     int32
		v7     *uint8
	)
	if tokCnt != 3 {
		return 0
	}
	result = int32(uintptr(memmap.PtrOff(6112660, 822660)))
	if true {
		v4 = nox_xxx_cliGamedataGet_416590(1)
		v5 = *(**libc.WChar)(memmap.PtrOff(0x587000, 94400))
		v6 = 0
		if *memmap.PtrUint32(0x587000, 94400) != 0 {
			v7 = (*uint8)(memmap.PtrOff(0x587000, 94400))
			for _nox_wcsicmp(v5, *(**libc.WChar)(unsafe.Add(unsafe.Pointer(tokens), unsafe.Sizeof((*libc.WChar)(nil))*uintptr(tokInd)))) != 0 {
				v5 = (*libc.WChar)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v7))), unsafe.Sizeof(uint32(0))*2))))))
				v7 = (*uint8)(unsafe.Add(unsafe.Pointer(v7), 8))
				v6++
				if v5 == nil {
					return 1
				}
			}
			*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v4))), unsafe.Sizeof(uint16(0))*26))) &= 0xE80F
			*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v4))), unsafe.Sizeof(uint16(0))*26))) |= *memmap.PtrUint16(0x587000, uint32(v6*8)+94404)
		}
		result = 1
	}
	return result
}
func nox_cmd_set_fr(tokInd int32, tokCnt int32, tokens **libc.WChar) int32 {
	if tokCnt != 2 {
		return 0
	}
	nox_xxx_setFrameLimit_43DDE0(1)
	return 1
}
func nox_cmd_unset_fr(tokInd int32, tokCnt int32, tokens **libc.WChar) int32 {
	if tokCnt != 2 {
		return 0
	}
	nox_xxx_setFrameLimit_43DDE0(0)
	return 1
}
func nox_cmd_set_net_debug(tokInd int32, tokCnt int32, tokens **libc.WChar) int32 {
	if tokCnt != 2 {
		return 0
	}
	noxflags.SetEngine(nox_engine_flag(NOX_ENGINE_FLAG_ENABLE_NET_DEBUG))
	return 1
}
func nox_cmd_unset_net_debug(tokInd int32, tokCnt int32, tokens **libc.WChar) int32 {
	if tokCnt != 2 {
		return 0
	}
	noxflags.UnsetEngine(nox_engine_flag(NOX_ENGINE_FLAG_ENABLE_NET_DEBUG))
	return 1
}
func nox_cmd_show_gui(tokInd int32, tokCnt int32, tokens **libc.WChar) int32 {
	var v0 int32 = int32(nox_client_renderGUI_80828 ^ 1)
	nox_client_renderGUI_80828 = uint32(v0)
	nox_xxx_xxxRenderGUI_587000_80832 = uint32(v0)
	return 1
}
func nox_cmd_show_netstat(tokInd int32, tokCnt int32, tokens **libc.WChar) int32 {
	sub_470A60()
	return 1
}
func nox_cmd_show_info(tokInd int32, tokCnt int32, tokens **libc.WChar) int32 {
	clientPlaySoundSpecial(921, 100)
	sub_435F60()
	return 1
}
func nox_cmd_show_mem(tokInd int32, tokCnt int32, tokens **libc.WChar) int32 {
	nox_server_currentMapGetFilename_409B30()
	nox_xxx_gameLoopMemDump_413E30()
	return 1
}
func nox_cmd_show_rank(tokInd int32, tokCnt int32, tokens **libc.WChar) int32 {
	if noxflags.HasGame(noxflags.GameOnline) {
		sub_4703F0()
	}
	return 1
}
func nox_cmd_show_motd(tokInd int32, tokCnt int32, tokens **libc.WChar) int32 {
	if tokCnt != 2 {
		return 0
	}
	nox_xxx_motd_4467F0()
	return 1
}
func nox_cmd_show_seq(tokInd int32, tokCnt int32, tokens **libc.WChar) int32 {
	if tokCnt != 2 {
		return 0
	}
	sub_48D7B0()
	return 1
}
func nox_cmd_list_maps(tokInd int32, tokCnt int32, tokens **libc.WChar) int32 {
	var (
		v0 *int32
		i  int32
		v2 bool
		v3 int32
	)
	libc.MemSet(memmap.PtrOff(6112660, 822404), 0, 256)
	v0 = (*int32)(unsafe.Pointer(nox_common_maplist_first_4D09B0()))
	for i = 1; v0 != nil; i++ {
		sub_4417E0((*libc.WChar)(memmap.PtrOff(6112660, 822404)), (*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(v0))), 12)))
		if (i % 4) == 0 {
			nox_gui_console_Printf_450C00(uint8(int8(NOX_CONSOLE_RED)), (*libc.WChar)(memmap.PtrOff(0x587000, 103276)), memmap.PtrOff(6112660, 822404))
			*memmap.PtrUint16(6112660, 822404) = 0
		}
		v0 = (*int32)(unsafe.Pointer(nox_common_maplist_next_4D09C0((*nox_map_list_item)(unsafe.Pointer(v0)))))
	}
	v3 = int32(uint32(i-1) & 0x80000003)
	v2 = v3 == 0
	if v3 < 0 {
		v2 = (uint32(int32(uint8(int8(v3)))-1) | 0xFFFFFFFC) == math.MaxUint32
	}
	if !v2 {
		nox_gui_console_Printf_450C00(uint8(int8(NOX_CONSOLE_RED)), (*libc.WChar)(memmap.PtrOff(0x587000, 103284)), memmap.PtrOff(6112660, 822404))
	}
	return 1
}
func nox_cmd_log_file(tokInd int32, tokCnt int32, tokens **libc.WChar) int32 {
	if tokCnt == 3 {
		if *(**libc.WChar)(unsafe.Add(unsafe.Pointer(tokens), unsafe.Sizeof((*libc.WChar)(nil))*2)) != nil {
			noxflags.SetEngine(nox_engine_flag(NOX_ENGINE_FLAG_LOG_TO_FILE))
			var v4 [256]byte
			nox_sprintf(&v4[0], CString("%S"), *(**libc.WChar)(unsafe.Add(unsafe.Pointer(tokens), unsafe.Sizeof((*libc.WChar)(nil))*2)))
			return nox_xxx_log_4_reopen_413A80(&v4[0])
		}
	}
	return 0
}
func nox_cmd_log_console(tokInd int32, tokCnt int32, tokens **libc.WChar) int32 {
	if tokCnt != 2 {
		return 0
	}
	noxflags.SetEngine(nox_engine_flag(NOX_ENGINE_FLAG_LOG_TO_CONSOLE))
	return 1
}
func nox_cmd_log_stop(tokInd int32, tokCnt int32, tokens **libc.WChar) int32 {
	if tokCnt != 2 {
		return 0
	}
	nox_xxx_log_4_close_413C00()
	return 1
}
func nox_cmd_set(tokInd int32, tokCnt int32, tokens **libc.WChar) int32 {
	return 1
}
func nox_cmd_cheat_ability(tokInd int32, tokCnt int32, tokens **libc.WChar) int32 {
	var i *byte
	if !noxflags.HasGame(noxflags.GameOnline) {
		for i = (*byte)(unsafe.Pointer(nox_common_playerInfoGetFirst_416EA0())); i != nil; i = (*byte)(unsafe.Pointer(nox_common_playerInfoGetNext_416EE0((*nox_playerInfo)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(i))))))))) {
			if *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(i))), unsafe.Sizeof(uint32(0))*514))) != 0 {
				nox_xxx_playerCancelAbils_4FC180(int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(i))), unsafe.Sizeof(uint32(0))*514)))))
			}
		}
	}
	return 1
}
func nox_cmd_cheat_level(tokInd int32, tokCnt int32, tokens **libc.WChar) int32 {
	var (
		i  *byte
		v5 uint8
	)
	if !noxflags.HasGame(noxflags.GameOnline) {
		if tokCnt < 3 {
			return 0
		}
		for i = (*byte)(unsafe.Pointer(nox_common_playerInfoGetFirst_416EA0())); i != nil; i = (*byte)(unsafe.Pointer(nox_common_playerInfoGetNext_416EE0((*nox_playerInfo)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(i))))))))) {
			if *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(i))), unsafe.Sizeof(uint32(0))*514))) != 0 {
				v5 = uint8(int8(nox_wcstol(*(**libc.WChar)(unsafe.Add(unsafe.Pointer(tokens), unsafe.Sizeof((*libc.WChar)(nil))*2)), nil, 10)))
				sub_4EF410(int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(i))), unsafe.Sizeof(uint32(0))*514)))), v5)
			}
		}
	}
	return 1
}
func nox_cmd_gamma(tokInd int32, tokCnt int32, tokens **libc.WChar) int32 {
	var (
		v3 *libc.WChar
		v4 int32
		v5 int32
		v6 int32
		v8 int32
	)
	if tokCnt <= 1 {
		return 1
	}
	v3 = *(**libc.WChar)(unsafe.Add(unsafe.Pointer(tokens), unsafe.Sizeof((*libc.WChar)(nil))*uintptr(tokInd)))
	if *v3 == 43 {
		v4 = nox_wcstol(v3, nil, 10)
		v8 = nox_video_getGammaSetting_434B00() + v4
		nox_video_setGammaSetting_434B30(v8)
	} else {
		if *v3 == 45 {
			v5 = nox_wcstol(v3, nil, 10)
			v6 = nox_video_getGammaSetting_434B00() - v5
		} else {
			v6 = nox_wcstol(v3, nil, 10)
		}
		nox_video_setGammaSetting_434B30(v6)
	}
	sub_434B60()
	return 1
}
func nox_cmd_window(tokInd int32, tokCnt int32, tokens **libc.WChar) int32 {
	var (
		v3 *libc.WChar
		v4 int32
		v6 int32
	)
	if tokCnt > 1 {
		v3 = *(**libc.WChar)(unsafe.Add(unsafe.Pointer(tokens), unsafe.Sizeof((*libc.WChar)(nil))*uintptr(tokInd)))
		if *v3 != 43 && *v3 != 45 {
			v4 = nox_wcstol(v3, nil, 10)
			nox_draw_setCutSize_476700(v4, 0)
			return 1
		}
		v6 = nox_wcstol(v3, nil, 10)
		nox_draw_setCutSize_476700(0, v6)
	}
	return 1
}
func nox_cmd_set_qual_modem(tokInd int32, tokCnt int32, tokens **libc.WChar) int32 {
	dword_5d4594_2650652 = 1
	var v0 int32 = sub_40A710(4)
	nox_xxx_rateUpdate_40A6D0(v0)
	nox_server_connectionType_3596 = 4
	sub_4AD840()
	return 1
}
func nox_cmd_set_qual_isdn(tokInd int32, tokCnt int32, tokens **libc.WChar) int32 {
	dword_5d4594_2650652 = 1
	var v0 int32 = sub_40A710(3)
	nox_xxx_rateUpdate_40A6D0(v0)
	nox_server_connectionType_3596 = 3
	sub_4AD840()
	return 1
}
func nox_cmd_set_qual_cable(tokInd int32, tokCnt int32, tokens **libc.WChar) int32 {
	dword_5d4594_2650652 = 1
	var v0 int32 = sub_40A710(2)
	nox_xxx_rateUpdate_40A6D0(v0)
	nox_server_connectionType_3596 = 2
	sub_4AD840()
	return 1
}
func nox_cmd_set_qual_t1(tokInd int32, tokCnt int32, tokens **libc.WChar) int32 {
	dword_5d4594_2650652 = 1
	var v0 int32 = sub_40A710(1)
	nox_xxx_rateUpdate_40A6D0(v0)
	nox_server_connectionType_3596 = 1
	sub_4AD840()
	return 1
}
func nox_cmd_set_qual_lan(tokInd int32, tokCnt int32, tokens **libc.WChar) int32 {
	dword_5d4594_2650652 = 0
	var v0 int32 = sub_40A710(1)
	nox_xxx_rateUpdate_40A6D0(v0)
	nox_server_connectionType_3596 = 1
	sub_4AD840()
	return 1
}
func nox_cmd_set_time(tokInd int32, tokCnt int32, tokens **libc.WChar) int32 {
	var (
		v4 *byte
		v5 *libc.WChar
		v6 uint8
	)
	if tokCnt != 3 {
		return 0
	}
	v4 = sub_4165B0()
	v5 = *(**libc.WChar)(unsafe.Add(unsafe.Pointer(tokens), unsafe.Sizeof((*libc.WChar)(nil))*2))
	if v5 != nil {
		v6 = uint8(int8(nox_wcstol(v5, nil, 10)))
		sub_40A040_settings(int16(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v4))), unsafe.Sizeof(uint16(0))*26)))), v6)
	}
	return 1
}
func nox_cmd_set_lessons(tokInd int32, tokCnt int32, tokens **libc.WChar) int32 {
	var (
		v4 *libc.WChar
		v5 *byte
		v6 uint16
	)
	if tokCnt != 3 {
		return 0
	}
	v4 = *(**libc.WChar)(unsafe.Add(unsafe.Pointer(tokens), unsafe.Sizeof((*libc.WChar)(nil))*2))
	v5 = sub_4165B0()
	if v4 != nil {
		v6 = uint16(int16(nox_wcstol(v4, nil, 10)))
		sub_409FB0_settings(int16(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v5))), unsafe.Sizeof(uint16(0))*26)))), v6)
	}
	return 1
}
func nox_cmd_menu_options(tokInd int32, tokCnt int32, tokens **libc.WChar) int32 {
	if !noxflags.HasGame(noxflags.GameFlag4) && noxflags.HasGame(noxflags.GameOnline) {
		nox_xxx_guiServerOptsLoad_457500()
	}
	return 1
}
func nox_cmd_menu_vidopt(tokInd int32, tokCnt int32, tokens **libc.WChar) int32 {
	sub_4ADA40()
	return 1
}
func nox_cmd_reenter(tokInd int32, tokCnt int32, tokens **libc.WChar) int32 {
	if !noxflags.HasGame(noxflags.GameOnline) {
		sub_40AA60(1)
	}
	return 1
}
