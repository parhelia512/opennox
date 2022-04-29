package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	noxflags "github.com/noxworld-dev/opennox/v1/common/flags"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"unsafe"
)

var gameex_flags uint32 = 30
var DefaultPacket [4]int32 = [4]int32{0xA3F0301, 1, 347, 0x2A55B62}
var wndEntryNames [5][35]libc.WChar = [5][35]libc.WChar{{119, 97, 114, 114, 105, 111, 114, 115, 32, 108, 105, 107, 101, 32, 115, 104, 105, 101, 108, 100, 115, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, {119, 97, 108, 107, 105, 110, 103, 32, 119, 105, 116, 104, 32, 115, 119, 111, 114, 100, 32, 98, 108, 111, 99, 107, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, {119, 101, 97, 112, 111, 110, 32, 114, 111, 108, 108, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, {115, 104, 105, 101, 108, 100, 32, 38, 32, 98, 101, 114, 115, 101, 114, 107, 101, 114, 32, 98, 108, 111, 99, 107, 105, 110, 103, 0, 0, 0, 0, 0, 0, 0, 0}, {101, 120, 116, 101, 110, 115, 105, 111, 110, 32, 109, 101, 115, 115, 97, 103, 101, 115, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}}
var someSwitch int8 = 0
var isInvalidIp int8 = 0
var inputNewIpMsgBox [512]byte

func nox_CharToOemW(pSrc *libc.WChar, pDst *byte) int32 {
	return nox_sprintf(pDst, CString("%S"), pSrc)
}
func gameex_sendPacket(buf *byte, len_ int32, smth int32) int32 {
	if buf == nil || len_ == 0 || *memmap.PtrUint32(6112660, 815700) >= NOX_NET_STRUCT_MAX {
		return 0
	}
	var ns *nox_net_struct_t = nox_net_struct_arr[*memmap.PtrUint32(6112660, 815700)]
	if ns == nil {
		return 0
	}
	return nox_net_sendto(ns.sock, unsafe.Pointer(buf), uint32(len_), &ns.addr)
}
func gameex_makeExtensionPacket(ptr int32, opcode int16, needsPlayer bool) {
	*(*uint16)(unsafe.Pointer(uintptr(ptr + 0))) = 0xF13A
	*(*uint16)(unsafe.Pointer(uintptr(ptr + 2))) = uint16(opcode)
	if needsPlayer {
		*(*uint32)(unsafe.Pointer(uintptr(ptr + 4))) = nox_player_netCode_85319C
	}
}
func getPlayerClassFromObjPtr(a1 int32) int8 {
	return int8(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))) + 276))) + 2251))))
}
func playerInfoStructParser_0(a1 *byte) int8 {
	var (
		v1   *byte
		pDst int8
	)
	if uintptr(unsafe.Pointer(a1)) == uintptr(0xFFFFFFFE) {
		return 0
	}
	v1 = (*byte)(unsafe.Pointer(nox_common_playerInfoGetFirst_416EA0()))
	if v1 == nil {
		return 0
	}
	for {
		nox_CharToOemW((*libc.WChar)(unsafe.Add(unsafe.Pointer((*libc.WChar)(unsafe.Pointer(v1))), unsafe.Sizeof(libc.WChar(0))*2352)), (*byte)(unsafe.Pointer(&pDst)))
		if libc.StrCmp((*byte)(unsafe.Pointer(&pDst)), (*byte)(unsafe.Add(unsafe.Pointer(a1), 2))) == 0 {
			break
		}
		v1 = (*byte)(unsafe.Pointer(nox_common_playerInfoGetNext_416EE0((*nox_playerInfo)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v1)))))))))
		if v1 == nil {
			return 0
		}
	}
	*(*byte)(unsafe.Add(unsafe.Pointer(a1), 1)) = byte(*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(nox_xxx_objGetTeamByNetCode_418C80(int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v1))), unsafe.Sizeof(uint32(0))*515)))))))), 4))))
	*a1 = *(*byte)(unsafe.Add(unsafe.Pointer(v1), 2251))
	return 1
}
func playerInfoStructParser_1(a1 int32, a2 int32, a3 *int32) int8 {
	var (
		v3   *byte
		v4   *byte
		v6   *uint32
		pDst int8
	)
	if a1 == -2 {
		return 0
	}
	v3 = (*byte)(unsafe.Pointer(nox_common_playerInfoGetFirst_416EA0()))
	*(*uint32)(unsafe.Pointer(uintptr(a2))) = uint32(uintptr(unsafe.Pointer(v3)))
	if v3 == nil {
		return 0
	}
	for {
		nox_CharToOemW((*libc.WChar)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a2)))+4704))), (*byte)(unsafe.Pointer(&pDst)))
		if libc.StrCmp((*byte)(unsafe.Pointer(&pDst)), (*byte)(unsafe.Pointer(uintptr(a1+2)))) == 0 {
			break
		}
		v4 = (*byte)(unsafe.Pointer(nox_common_playerInfoGetNext_416EE0((*nox_playerInfo)(unsafe.Pointer(uintptr(a2))))))
		*(*uint32)(unsafe.Pointer(uintptr(a2))) = uint32(uintptr(unsafe.Pointer(v4)))
		if v4 == nil {
			return 0
		}
	}
	v6 = nox_xxx_objGetTeamByNetCode_418C80(int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a2))) + 2060)))))
	*a3 = int32(uintptr(unsafe.Pointer(v6)))
	*(*uint8)(unsafe.Pointer(uintptr(a1 + 1))) = *((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(v6))), 4)))
	*(*uint8)(unsafe.Pointer(uintptr(a1))) = *(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a2))) + 2251)))
	return 1
}
func mix_MouseKeyboardWeaponRoll(playerObj int32, a2 int8) int8 {
	var (
		v2        int32
		v4        int32
		i         int32
		v6        int32
		v8        int32
		weapFlags int32
		v11       int32
		v16       int8
	)
	v16 = 0
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(playerObj + 748))))
	if (int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 276))) + 3680))))&3) == 0 && int32(*(*uint8)(unsafe.Pointer(uintptr(v2 + 88)))) != 1 {
		var cur int32 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 104))))
		v4 = 496
		if int32(a2) == 0 {
			v4 = 500
		}
		if cur != 0 {
			var next int32 = cur
			for {
				next = int32(*(*uint32)(unsafe.Pointer(uintptr(next + v4))))
				if next == 0 {
					break
				}
				weapFlags = nox_xxx_weaponInventoryEquipFlags_415820((*nox_object_t)(unsafe.Pointer(uintptr(next))))
				if weapFlags != 0 && weapFlags != 2 {
					if nox_xxx_playerClassCanUseItem_57B3D0((*nox_object_t)(unsafe.Pointer(uintptr(next))), int8(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(playerObj + 748))) + 276))) + 2251))))) != 0 {
						v11 = bool2int(nox_xxx_playerCheckStrength_4F3180((*nox_object_t)(unsafe.Pointer(uintptr(playerObj))), (*nox_object_t)(unsafe.Pointer(uintptr(next)))))
						if v11 != 0 {
							if nox_xxx_playerTryDequip_4F2FB0((*uint32)(unsafe.Pointer(uintptr(playerObj))), (*nox_object_t)(unsafe.Pointer(uintptr(cur)))) != 0 && nox_xxx_playerTryEquip_4F2F70((*uint32)(unsafe.Pointer(uintptr(playerObj))), (*nox_object_t)(unsafe.Pointer(uintptr(next)))) != 0 {
								v16 = 1
							}
							return v16
						}
					}
				}
			}
		} else {
			for i = int32(*(*uint32)(unsafe.Pointer(uintptr(playerObj + 504)))); i != 0; i = int32(*(*uint32)(unsafe.Pointer(uintptr(i + 496)))) {
				v6 = nox_xxx_weaponInventoryEquipFlags_415820((*nox_object_t)(unsafe.Pointer(uintptr(i))))
				if v6 != 0 && v6 != 2 {
					if nox_xxx_playerClassCanUseItem_57B3D0((*nox_object_t)(unsafe.Pointer(uintptr(i))), int8(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(playerObj + 748))) + 276))) + 2251))))) != 0 {
						v8 = bool2int(nox_xxx_playerCheckStrength_4F3180((*nox_object_t)(unsafe.Pointer(uintptr(playerObj))), (*nox_object_t)(unsafe.Pointer(uintptr(i)))))
						if v8 != 0 {
							if nox_xxx_playerTryEquip_4F2F70((*uint32)(unsafe.Pointer(uintptr(playerObj))), (*nox_object_t)(unsafe.Pointer(uintptr(i)))) != 0 {
								v16 = 1
							}
							return v16
						}
					}
				}
			}
		}
	}
	return v16
}
func playerDropATrap(playerObj int32) int8 {
	var (
		v2  int32
		i   int32
		v7  int8
		v8  int8
		pos [2]float32 = [2]float32{}
	)
	v7 = 17
	if playerObj == 0 {
		return 0
	}
	v8 = 0
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(playerObj + 748))) + 276))))
	pos[0] = *(*float32)(unsafe.Pointer(uintptr(v2 + 3632)))
	pos[1] = *(*float32)(unsafe.Pointer(uintptr(v2 + 3636)))
	if (int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(playerObj + 748))) + 276))) + 3680))))&3) == 0 && int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(playerObj + 748))) + 88)))) != 1 {
		for i = int32(*(*uint32)(unsafe.Pointer(uintptr(playerObj + 504)))); i != 0; i = int32(*(*uint32)(unsafe.Pointer(uintptr(i + 496)))) {
			if int32(*(*uint8)(unsafe.Pointer(uintptr(i + 10)))) == int32(v7) {
				nox_xxx_drop_4ED810(playerObj, i, &pos[0])
				return 1
			}
		}
	}
	return v8
}
func playErrSoundClient() {
	clientPlaySoundSpecial(766, 100)
}
func MixRecvFromReplacer(s nox_socket_t, buf *byte, len_ int32, from *nox_net_sockaddr) int32 {
	var (
		v8  *uint32
		v9  int8
		v10 int32
		v11 int32
	)
	_ = v11
	var v13 *byte
	_ = v13
	var v18 *byte
	var v19 int8
	var v20 *byte
	var v21 *byte
	var v22 int8
	var v23 int32
	var v24 int32
	var v25 int32
	var v26 *byte
	var v27 *byte
	var v28 int8
	var v29 uint32
	var v30 *byte
	var v31 uint32
	var v32 *uint8
	var v33 int8
	var v35 *uint32
	var v36 int8
	var to *nox_net_sockaddr
	var v39 int32
	var v43 [6]int32
	var v44 [128]uint8
	to = from
	var op int32 = int32(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(buf))), unsafe.Sizeof(uint16(0))*1))))
	switch op {
	case 0:
		if (gameex_flags>>3)&1 != 0 {
			v8 = nox_xxx_objGetTeamByNetCode_418C80(int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(buf))), unsafe.Sizeof(uint32(0))*1)))))
			v9 = int8(*(*byte)(unsafe.Add(unsafe.Pointer(buf), 8)))
			v10 = int32(uintptr(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer(v8), -int(unsafe.Sizeof(uint32(0))*12))))))
			v36 = int8(uint8((uint32(v9) >> 4) & 1))
			if int32(v36) != 0 {
				*(*byte)(unsafe.Add(unsafe.Pointer(buf), 8)) = byte(int8(int32(v9) & 239))
			}
			if int32(getPlayerClassFromObjPtr(v10)) == 0 || int32(v36) != 0 {
				if int32(mix_MouseKeyboardWeaponRoll(v10, int8(*(*byte)(unsafe.Add(unsafe.Pointer(buf), 8))))) != 0 {
					*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(buf))), unsafe.Sizeof(uint16(0))*1))) = 2
					nox_net_sendto(s, unsafe.Pointer(buf), 4, (*nox_net_sockaddr_in)(unsafe.Pointer(from)))
				}
			}
		}
	case 1:
		fallthrough
	case 6:
	case 2:
		clientPlaySoundSpecial(895, 100)
	case 4:
		if noxflags.HasGame(noxflags.GameHost) {
			if (gameex_flags>>5)&1 != 0 {
				v18 = (*byte)(unsafe.Add(unsafe.Pointer(buf), 4))
				for {
					v19 = int8(*v18)
					*(*byte)(unsafe.Add(unsafe.Pointer(v18), int32(uintptr(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(&v43[0]))), 2))))-uintptr(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(buf), 4))))))) = *v18
					v18 = (*byte)(unsafe.Add(unsafe.Pointer(v18), 1))
					if int32(v19) == 0 {
						break
					}
				}
				if int32(playerInfoStructParser_0((*byte)(unsafe.Pointer(&v43[0])))) != 0 {
					*(*byte)(unsafe.Add(unsafe.Pointer(buf), 2)) = 6
					v20 = sub_433890()
					v21 = (*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(&v43[0]))), 2))
					for {
						v22 = int8(*v20)
						*func() *byte {
							p := &v21
							x := *p
							*p = (*byte)(unsafe.Add(unsafe.Pointer(*p), 1))
							return x
						}() = *func() *byte {
							p := &v20
							x := *p
							*p = (*byte)(unsafe.Add(unsafe.Pointer(*p), 1))
							return x
						}()
						if int32(v22) == 0 {
							break
						}
					}
					v23 = v43[1]
					v24 = v43[2]
					*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(buf))), unsafe.Sizeof(uint32(0))*1))) = uint32(v43[0])
					v25 = v43[3]
					*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(buf))), unsafe.Sizeof(uint32(0))*2))) = uint32(v23)
					*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v23))), unsafe.Sizeof(uint16(0))*0)) = uint16(int16(v43[4]))
					*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(buf))), unsafe.Sizeof(uint32(0))*3))) = uint32(v24)
					*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(buf))), unsafe.Sizeof(uint32(0))*4))) = uint32(v25)
					*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(buf))), unsafe.Sizeof(uint16(0))*10))) = uint16(int16(v23))
					nox_net_sendto(s, unsafe.Pointer(buf), 22, (*nox_net_sockaddr_in)(unsafe.Pointer(from)))
				}
			}
		}
	case 5:
		if (gameex_flags>>5)&1 != 0 {
			libc.MemSet(unsafe.Pointer(&v44[0]), 0, 128)
			libc.Mbstowcs((*libc.WChar)(unsafe.Pointer(&v44[0])), (*byte)(unsafe.Add(unsafe.Pointer(buf), 4)), uint32(libc.StrLen((*byte)(unsafe.Add(unsafe.Pointer(buf), 4)))))
			nox_xxx_printCentered_445490((*libc.WChar)(unsafe.Pointer(&v44[0])))
			clientPlaySoundSpecial(901, 100)
		}
	case 7:
		libc.MemSet(unsafe.Pointer(&v44[0]), 0, 128)
		libc.Mbstowcs((*libc.WChar)(unsafe.Pointer(&v44[0])), (*byte)(unsafe.Add(unsafe.Pointer(buf), 4)), uint32(libc.StrLen((*byte)(unsafe.Add(unsafe.Pointer(buf), 4)))))
		nox_xxx_printCentered_445490((*libc.WChar)(unsafe.Pointer(&v44[0])))
		clientPlaySoundSpecial(901, 100)
	case 8:
		if noxflags.HasGame(noxflags.GameHost) && (gameex_flags>>5)&1 != 0 {
			v26 = (*byte)(unsafe.Add(unsafe.Pointer(buf), 4))
			v27 = (*byte)(unsafe.Add(unsafe.Pointer(buf), 4))
			for {
				v28 = int8(*v27)
				*(*byte)(unsafe.Add(unsafe.Pointer(v27), int32(uintptr(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(&v43[0]))), 2))))-uintptr(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(buf), 4))))))) = *v27
				v27 = (*byte)(unsafe.Add(unsafe.Pointer(v27), 1))
				if int32(v28) == 0 {
					break
				}
			}
			if int32(playerInfoStructParser_1(int32(uintptr(unsafe.Pointer(&v43[0]))), int32(uintptr(unsafe.Pointer(&to))), &v39)) != 0 {
				*(*byte)(unsafe.Add(unsafe.Pointer(buf), 2)) = 7
				v29 = uint32(libc.StrLen((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(&v43[0]))), 2))))
				v30 = (*byte)(unsafe.Add(unsafe.Pointer(buf), v29+4))
				v31 = uint32(libc.StrLen(v30))
				if v31 != 0 {
					v32 = (*uint8)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(&v44[0])) - uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(v30))))))))
					for {
						v33 = int8(*v30)
						*(*byte)(unsafe.Add(unsafe.Pointer(v30), uint32(uintptr(unsafe.Pointer(v32))))) = *v30
						v30 = (*byte)(unsafe.Add(unsafe.Pointer(v30), 1))
						if int32(v33) == 0 {
							break
						}
					}
					libc.MemSet(unsafe.Pointer(v26), 0, int(v29+v31))
					libc.StrCpy(v26, (*byte)(unsafe.Pointer(&v44[0])))
				}
			}
		}
	case 9:
		if (gameex_flags>>3)&1 != 0 {
			v35 = nox_xxx_objGetTeamByNetCode_418C80(int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(buf))), unsafe.Sizeof(uint32(0))*1)))))
			playerDropATrap(int32(uintptr(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer(v35), -int(unsafe.Sizeof(uint32(0))*12)))))))
		}
	}
	*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(buf))), unsafe.Sizeof(uint32(0))*0))) = uint32(DefaultPacket[0])
	*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(buf))), unsafe.Sizeof(uint32(0))*1))) = uint32(DefaultPacket[1])
	*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(buf))), unsafe.Sizeof(uint32(0))*2))) = uint32(DefaultPacket[2])
	*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(buf))), unsafe.Sizeof(uint32(0))*3))) = uint32(DefaultPacket[3])
	return 16
}
func OnLibraryNotice_256() {
}
func OnLibraryNotice_258() {
	someSwitch = 1
}
func OnLibraryNotice_259(arg1 uint32) {
}
func OnLibraryNotice_262(arg1 uint32) {
}
func OnLibraryNotice_263(arg1 uint32) {
	noxflags.HasGame(noxflags.GameHost)
}
func OnLibraryNotice_264(arg1 uint32) {
	noxflags.HasGame(noxflags.GameHost)
}
func OnLibraryNotice_265(arg1 uint32, arg2 uint32, arg3 int32) {
	var a2a int8 = int8(bool2int(arg3 > 0))
	if arg2 == 2 && gameexSomeWeirdCheckFixmePlease() {
		if (gameex_flags>>3)&1 != 0 {
			if noxflags.HasGame(noxflags.GameFlag3 | noxflags.GameModeSolo10) {
				if noxflags.HasGame(noxflags.GameHost) {
					if nox_xxx_host_player_unit_3843628 != nil {
						if int32(getPlayerClassFromObjPtr(int32(uintptr(unsafe.Pointer(nox_xxx_host_player_unit_3843628))))) == 0 {
							if int32(mix_MouseKeyboardWeaponRoll(int32(uintptr(unsafe.Pointer(nox_xxx_host_player_unit_3843628))), a2a)) != 0 {
								clientPlaySoundSpecial(895, 100)
							}
						}
					}
				} else {
					var buf [10]byte
					gameex_makeExtensionPacket(int32(uintptr(unsafe.Pointer(&buf[0]))), 0, true)
					buf[8] = byte(a2a)
					gameex_sendPacket(&buf[0], 9, 0)
				}
			}
		}
	}
}
func OnLibraryNotice_420(arg1 uint32, arg2 uint32, arg3 uint32, arg4 uint32) {
	var (
		v23 int32   = int32(arg1)
		v19 int32   = int32(arg2)
		v16 *uint32 = (*uint32)(unsafe.Pointer(uintptr(getPlayerClassFromObjPtr(int32(arg1)))))
	)
	if int32(*(*uint8)(unsafe.Pointer(uintptr(v19 + 10)))) != 17 {
		nox_xxx_inventoryServPlace_4F36F0(v23, v19, 1, 1)
		return
	}
	var v17 int8 = int8(*(*uint8)(unsafe.Pointer(uintptr(v19 + 4))))
	if int32(v17) != 106 {
		if (int32(v17) == 107 || int32(v17) == 109) && int32(uint8(uintptr(unsafe.Pointer(v16)))) != 0 {
			goto ifIsWarrior
		}
		nox_xxx_inventoryServPlace_4F36F0(v23, v19, 1, 1)
		return
	}
	if int32(uint8(uintptr(unsafe.Pointer(v16)))) == 1 {
		nox_xxx_inventoryServPlace_4F36F0(v23, v19, 1, 1)
		return
	}
ifIsWarrior:
	nox_xxx_netPriMsgToPlayer_4DA2C0((*nox_object_t)(unsafe.Pointer(uintptr(v23))), (*byte)(memmap.PtrOff(0x587000, 215732)), 0)
	nox_xxx_aud_501960(925, (*nox_object_t)(unsafe.Pointer(uintptr(v23))), 2, int32(*(*uint32)(unsafe.Pointer(uintptr(v23 + 36)))))
}
func getFlagValueFromFlagIndex(a1 int32) int32 {
	var (
		v1     int32
		v2     uint32
		v3     int32
		result int32
	)
	v1 = 2
	v2 = uint32(a1)
	if a1 < 0 {
		v2 = uint32(-a1)
	}
	v3 = 1
	for {
		if v2&1 != 0 {
			v3 *= v1
		}
		v2 >>= 1
		if v2 == 0 {
			break
		}
		v1 *= v1
	}
	if a1 >= 0 {
		result = v3
	} else {
		result = 1 / v3
	}
	return result
}
