package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	"github.com/gotranspile/cxgo/runtime/stdio"
	"github.com/noxworld-dev/opennox/v1/common/alloc"
	noxflags "github.com/noxworld-dev/opennox/v1/common/flags"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"math"
	"unicode"
	"unsafe"
)

var nox_xxx_wallSounds_2386840 uint32 = 0

func sub_5098A0() int32 {
	var (
		v0     int32
		v1     int32
		v2     int32
		i      *byte
		v4     int32
		v5     int32
		v6     int32
		v7     int32
		v8     int32
		result int32
		v10    *byte
		v11    int32
	)
	v0 = 0
	v11 = 0
	v10 = nil
	v1 = 0
	v2 = math.MinInt32
	for i = (*byte)(unsafe.Pointer(nox_server_teamFirst_418B10())); i != nil; i = (*byte)(unsafe.Pointer(nox_server_teamNext_418B60((*nox_team_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(i))))))))) {
		v4 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(i))), unsafe.Sizeof(uint32(0))*13))))
		if v4 >= v2 {
			v1 = bool2int(v4 == v2 && v0 != 0)
			v2 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(i))), unsafe.Sizeof(uint32(0))*13))))
			v10 = i
			v0 = int32(uintptr(unsafe.Pointer(i)))
		}
	}
	v5 = int32(uintptr(unsafe.Pointer(nox_xxx_getFirstPlayerUnit_4DA7C0())))
	if v5 != 0 {
		for {
			v6 = int32(*(*uint32)(unsafe.Pointer(uintptr(v5 + 748))))
			if nox_xxx_servObjectHasTeam_419130(v5+48) == 0 {
				v7 = int32(*(*uint32)(unsafe.Pointer(uintptr(v6 + 276))))
				if (int32(*(*uint8)(unsafe.Pointer(uintptr(v7 + 3680)))) & 1) == 0 {
					v8 = int32(*(*uint32)(unsafe.Pointer(uintptr(v7 + 2136))))
					if v8 >= v2 {
						v1 = bool2int(v8 == v2 && v11 != 0)
						v2 = v8
						v11 = v5
					}
				}
			}
			v5 = int32(uintptr(unsafe.Pointer(nox_xxx_getNextPlayerUnit_4DA7F0((*nox_object_t)(unsafe.Pointer(uintptr(v5)))))))
			if v5 == 0 {
				break
			}
		}
		v0 = int32(uintptr(unsafe.Pointer(v10)))
	}
	noxflags.SetGame(noxflags.GameFlag4)
	if v1 != 0 {
		return nox_xxx_netSendDMTeamWinner_4D8BF0(0, 1)
	}
	result = v11
	if v11 != 0 {
		return nox_xxx_netSendDMWinner_4D8B90(v11, 1)
	}
	if v0 != 0 {
		result = nox_xxx_netSendDMTeamWinner_4D8BF0(v0, 1)
	}
	return result
}
func sub_5099B0() int32 {
	var (
		v0     int32
		v1     int32
		v2     int32
		i      *byte
		v4     int32
		result int32
	)
	v0 = 0
	v1 = 0
	v2 = -1
	for i = (*byte)(unsafe.Pointer(nox_server_teamFirst_418B10())); i != nil; i = (*byte)(unsafe.Pointer(nox_server_teamNext_418B60((*nox_team_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(i))))))))) {
		v4 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(i))), unsafe.Sizeof(uint32(0))*13))))
		if v4 >= v2 {
			v1 = bool2int(v4 == v2 && v0 != 0)
			v2 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(i))), unsafe.Sizeof(uint32(0))*13))))
			v0 = int32(uintptr(unsafe.Pointer(i)))
		}
	}
	noxflags.SetGame(noxflags.GameFlag4)
	if v0 == 0 || v1 != 0 {
		if noxflags.HasGame(noxflags.GameModeFlagBall) {
			result = nox_xxx_netFlagballWinner_4D8C40(0)
		} else {
			result = nox_xxx_netFlagWinner_4D8C40_4D8C80(0, 1)
		}
	} else if noxflags.HasGame(noxflags.GameModeFlagBall) {
		result = nox_xxx_netFlagballWinner_4D8C40(v0)
	} else {
		result = nox_xxx_netFlagWinner_4D8C40_4D8C80(v0, 1)
	}
	return result
}
func nox_server_checkVictory_509A60() {
	if noxflags.HasGame(noxflags.GameModeElimination) {
		var (
			v6 *byte  = nil
			v7 int16  = int16(uint16(nox_common_gameFlags_getVal_40A5B0()))
			v0 uint16 = uint16(nox_xxx_servGamedataGet_40A020(v7))
			v8 int32  = int32(v0)
		)
		if v8 < 1 {
			return
		}
		var v9 int32 = 0
		for i := int32(int32(uintptr(unsafe.Pointer(nox_xxx_getFirstPlayerUnit_4DA7C0())))); i != 0; i = int32(uintptr(unsafe.Pointer(nox_xxx_getNextPlayerUnit_4DA7F0((*nox_object_t)(unsafe.Pointer(uintptr(i))))))) {
			var v11 int32 = int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(i + 748))) + 276))))
			if (int32(*(*uint8)(unsafe.Pointer(uintptr(v11 + 3680))))&1) == 0 && *(*uint32)(unsafe.Pointer(uintptr(v11 + 2140))) < uint32(v8) {
				if nox_xxx_servObjectHasTeam_419130(i+48) != 0 {
					var v0 *byte = (*byte)(unsafe.Pointer(nox_xxx_clientGetTeamColor_418AB0(int32(*(*uint8)(unsafe.Pointer(uintptr(i + 52)))))))
					if v6 != nil {
						if v6 != v0 {
							return
						}
					} else {
						v6 = v0
					}
				} else {
					if v9 != 0 || v6 != nil {
						return
					}
					v9 = i
				}
			}
		}
		if nox_xxx_gamePlayIsAnyPlayers_40A8A0() == 0 {
			return
		}
		noxflags.SetGame(noxflags.GameFlag4)
		if v6 != nil {
			nox_xxx_netSendDMTeamWinner_4D8BF0(int32(uintptr(unsafe.Pointer(v6))), 0)
		} else if v9 != 0 {
			nox_xxx_netSendDMWinner_4D8B90(v9, 0)
		} else {
			nox_xxx_netSendDMWinner_4D8B90(0, 0)
		}
		return
	}
	if noxflags.HasGame(noxflags.GameModeSolo10) {
		return
	}
	var v1 int16 = int16(uint16(nox_common_gameFlags_getVal_40A5B0()))
	var v0a uint16 = uint16(nox_xxx_servGamedataGet_40A020(v1))
	var v2 int32 = int32(v0a)
	if int32(v0a) == 0 {
		return
	}
	for v3 := (*byte)((*byte)(unsafe.Pointer(nox_server_teamFirst_418B10()))); v3 != nil; v3 = (*byte)(unsafe.Pointer(nox_server_teamNext_418B60((*nox_team_t)(unsafe.Pointer(v3))))) {
		if *((*int32)(unsafe.Add(unsafe.Pointer((*int32)(unsafe.Pointer(v3))), unsafe.Sizeof(int32(0))*13))) >= v2 {
			noxflags.SetGame(noxflags.GameFlag4)
			nox_xxx_netSendDMTeamWinner_4D8BF0(int32(uintptr(unsafe.Pointer(v3))), 0)
			return
		}
	}
	for v4 := int32(int32(uintptr(unsafe.Pointer(nox_xxx_getFirstPlayerUnit_4DA7C0())))); v4 != 0; v4 = int32(uintptr(unsafe.Pointer(nox_xxx_getNextPlayerUnit_4DA7F0((*nox_object_t)(unsafe.Pointer(uintptr(v4))))))) {
		var v5 int32 = int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v4 + 748))) + 276))))
		if (int32(*(*uint8)(unsafe.Pointer(uintptr(v5 + 3680))))&1) == 0 && *(*int32)(unsafe.Pointer(uintptr(v5 + 2136))) >= v2 {
			noxflags.SetGame(noxflags.GameFlag4)
			nox_xxx_netSendDMWinner_4D8B90(v4, 0)
			break
		}
	}
}
func sub_509C30(pl *nox_playerInfo) {
	var (
		a1 int32 = int32(uintptr(unsafe.Pointer(pl)))
		v1 *byte
	)
	if dword_5d4594_1599688 == 0 {
		nox_common_list_clear_425760((*nox_list_item_t)(memmap.PtrOff(6112660, 1599676)))
		dword_5d4594_1599688 = 1
	}
	v1 = (*byte)(alloc.Calloc(1, 32))
	*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v1))), unsafe.Sizeof(uint32(0))*6))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 2068)))
	libc.StrCpy((*byte)(unsafe.Add(unsafe.Pointer(v1), 12)), (*byte)(unsafe.Pointer(uintptr(a1+2096))))
	*(*byte)(unsafe.Add(unsafe.Pointer(v1), 28)) = byte(*(*uint8)(unsafe.Pointer(uintptr(a1 + 2251))))
	sub_4257F0(memmap.PtrInt32(6112660, 1599676), (*uint32)(unsafe.Pointer(v1)))
}
func sub_509CB0() *int32 {
	var (
		result *int32
		v1     *int32
		v2     *int32
	)
	result = *(**int32)(unsafe.Pointer(&dword_5d4594_1599688))
	if dword_5d4594_1599688 != 0 {
		result = (*int32)(unsafe.Pointer(nox_common_list_getFirstSafe_425890((*nox_list_item_t)(unsafe.Pointer(memmap.PtrInt32(6112660, 1599676))))))
		v1 = result
		if result != nil {
			for {
				v2 = (*int32)(unsafe.Pointer(nox_common_list_getNextSafe_4258A0((*nox_list_item_t)(unsafe.Pointer(v1)))))
				sub_425920((**uint32)(unsafe.Pointer(v1)))
				alloc.Free(unsafe.Pointer(v1))
				v1 = v2
				if v2 == nil {
					break
				}
			}
		}
	}
	return result
}
func sub_509CF0(a1 *byte, a2 int8, a3 int32) int32 {
	var v3 *int32
	v3 = (*int32)(unsafe.Pointer(nox_common_list_getFirstSafe_425890((*nox_list_item_t)(unsafe.Pointer(memmap.PtrInt32(6112660, 1599676))))))
	if v3 == nil {
		return 1
	}
	for libc.StrCmp((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(v3))), 12)), a1) != 0 || int32(a2) == int32(*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(v3))), 28)))) && a3 == *(*int32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(int32(0))*6)) {
		v3 = (*int32)(unsafe.Pointer(nox_common_list_getNextSafe_4258A0((*nox_list_item_t)(unsafe.Pointer(v3)))))
		if v3 == nil {
			return 1
		}
	}
	return 0
}
func sub_509D80(a1 int32) int32 {
	var v1 *int32
	v1 = (*int32)(unsafe.Pointer(nox_common_list_getFirstSafe_425890((*nox_list_item_t)(unsafe.Pointer(memmap.PtrInt32(6112660, 1599676))))))
	if v1 == nil {
		return 0
	}
	for libc.StrCmp((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(v1))), 12)), (*byte)(unsafe.Pointer(uintptr(a1+2096)))) != 0 {
		v1 = (*int32)(unsafe.Pointer(nox_common_list_getNextSafe_4258A0((*nox_list_item_t)(unsafe.Pointer(v1)))))
		if v1 == nil {
			return 0
		}
	}
	return 1
}
func nox_xxx_xferDirectionToAngle_509E00(a1 *uint32) int32 {
	return int32(*memmap.PtrUint32(0x587000, (*a1+*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*1))*3)*4+0x382B8))
}
func nox_xxx_xferIndexedDirection_509E20(a1 int32, a2 *int2) int32 {
	var (
		v2     int32
		v3     int32
		result int32
	)
	v2 = *memmap.PtrInt32(0x587000, uint32(a1*8)+0x2EE58)
	if v2 <= *(*int32)(unsafe.Pointer(&dword_587000_230092)) {
		a2.field_0 = bool2int(v2 >= -*(*int32)(unsafe.Pointer(&dword_587000_230092))) - 1
	} else {
		a2.field_0 = 1
	}
	v3 = *memmap.PtrInt32(0x587000, uint32(a1*8)+0x2EE5C)
	result = int32(dword_587000_230092)
	if v3 <= *(*int32)(unsafe.Pointer(&dword_587000_230092)) {
		result = -*(*int32)(unsafe.Pointer(&dword_587000_230092))
		if v3 >= -*(*int32)(unsafe.Pointer(&dword_587000_230092)) {
			a2.field_4 = 0
		} else {
			a2.field_4 = -1
		}
	} else {
		a2.field_4 = 1
	}
	return result
}
func nox_xxx_mathDirection4ToAngle_509E90(a1 int32) int32 {
	return int32(*memmap.PtrUint32(0x587000, uint32(a1*4)+0x382A8))
}
func nox_xxx_math_509EA0(a1 int32) int32 {
	var a2 int2
	nox_xxx_xferIndexedDirection_509E20(a1, &a2)
	return a2.field_4 + a2.field_0 + a2.field_4*2 + 4
}
func nox_xxx_math_509ED0(a1 *float2) int32 {
	var (
		result int32
		v2     float32
	)
	v2 = float32((math.Atan2(float64(a1.field_4), float64(a1.field_0))+6.2831855)*40.743664 + 0.5)
	result = int(v2)
	if result < 0 {
		result += int32(uint32(math.MaxUint8-result) >> 8 << 8)
	}
	if result >= 256 {
		result += int32((uint32(result) >> 8) * 0xFFFFFF00)
	}
	return result
}
func nox_xxx_utilNormalizeVector_509F20(a1 *float2) {
	var v1 float64
	v1 = float64(nox_double2float(math.Sqrt(float64(a1.field_0*a1.field_0 + a1.field_4*a1.field_4))))
	a1.field_0 = float32(float64(a1.field_0) / v1)
	a1.field_4 = float32(float64(a1.field_4) / v1)
}
func sub_509FF0(a1 int32) int32 {
	var result int32
	result = a1
	if int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1))) + 16))))&32 != 0 {
		*(*uint32)(unsafe.Pointer(uintptr(a1))) = 0
	}
	return result
}
func nox_xxx_monsterActionIsCondition_50A010(a1 int32) int32 {
	var result int32
	result = bool2int(a1 < 39)
	*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&result))), 0)) = uint8(int8(bool2int(a1 > 39)))
	return result
}
func nox_xxx_mobActionGet_50A020(a1 int32) int32 {
	return int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))) + uint32((*(*byte)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))) + 544)))+23)*24)))))
}
func sub_50A040(a1 int32) int32 {
	var (
		v1 int32
		v2 int32
		i  *int32
		v4 int32
	)
	v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
	v2 = int32(*(*byte)(unsafe.Pointer(uintptr(v1 + 544))) - 1)
	if v2 < 0 {
		return 38
	}
	for i = (*int32)(unsafe.Pointer(uintptr(v1 + (v2*3+69)*8))); ; i = (*int32)(unsafe.Add(unsafe.Pointer(i), -int(unsafe.Sizeof(int32(0))*6))) {
		v4 = nox_xxx_monsterActionIsCondition_50A010(*i)
		if v4 == 0 {
			break
		}
		if func() int32 {
			p := &v2
			*p--
			return *p
		}() < 0 {
			return 38
		}
	}
	return int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + (v2*3+69)*8))))
}
func nox_xxx_monsterIsActionScheduled_50A090(a1 int32, a2 int32) int32 {
	var (
		v2 int32
		v3 int32
		i  *uint32
	)
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
	v3 = int32(*(*byte)(unsafe.Pointer(uintptr(v2 + 544))) - 1)
	if v3 < 0 {
		return 0
	}
	for i = (*uint32)(unsafe.Pointer(uintptr(v2 + (v3*3+69)*8))); *i != uint32(a2); i = (*uint32)(unsafe.Add(unsafe.Pointer(i), -int(unsafe.Sizeof(uint32(0))*6))) {
		if func() int32 {
			p := &v3
			*p--
			return *p
		}() < 0 {
			return 0
		}
	}
	return 1
}
func nox_xxx_checkMobAction_50A0D0(a1p *nox_object_t, a2 int32) int32 {
	var (
		a1 int32 = int32(uintptr(unsafe.Pointer(a1p)))
		v2 int32
		v3 int32
		i  *uint32
	)
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
	v3 = int32(*(*byte)(unsafe.Pointer(uintptr(v2 + 544))))
	if v3 < 0 {
		return 0
	}
	for i = (*uint32)(unsafe.Pointer(uintptr(v2 + (v3*3+69)*8))); *i != uint32(a2); i = (*uint32)(unsafe.Add(unsafe.Pointer(i), -int(unsafe.Sizeof(uint32(0))*6))) {
		if func() int32 {
			p := &v3
			*p--
			return *p
		}() < 0 {
			return 0
		}
	}
	return 1
}
func nox_xxx_monsterActionReset_50A110(a1p *nox_object_t) {
	var (
		a1     int32 = int32(uintptr(unsafe.Pointer(a1p)))
		result int32
	)
	result = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
	*(*uint8)(unsafe.Pointer(uintptr(result + 481))) = 0
	*(*uint8)(unsafe.Pointer(uintptr(result + 482))) = 0
	*(*uint8)(unsafe.Pointer(uintptr(result + 483))) = 0
	*(*uint32)(unsafe.Pointer(uintptr(result + 8))) = 0
	*(*uint32)(unsafe.Pointer(uintptr(result + 268))) = 0
	*(*uint32)(unsafe.Pointer(uintptr(result + 296))) = 0
	*(*uint32)(unsafe.Pointer(uintptr(result + 364))) = 0
	*(*uint32)(unsafe.Pointer(uintptr(result + 548))) = nox_frame_xxx_2598000
	*(*uint32)(unsafe.Pointer(uintptr(result + 496))) = nox_frame_xxx_2598000
}
func nox_xxx_monsterPopAction_50A160(a1p *nox_object_t) int8 {
	var (
		a1     int32 = int32(uintptr(unsafe.Pointer(a1p)))
		v1     int32
		v2     int32
		v3     int32
		v4     int8
		v5     *uint32
		v6     int32
		v7     func(int32)
		v8     int8
		v9     int32
		v10    *int32
		v11    int32
		result int8
		v13    int32
		v14    int32
		v15    int32
	)
	v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
	v2 = int32(*(*byte)(unsafe.Pointer(uintptr(v1 + 544))))
	v15 = v2
	v14 = int32(*memmap.PtrUint32(0x587000, *(*uint32)(unsafe.Pointer(uintptr(v1 + (v2*3+69)*8)))*4+0x382D0))
	v13 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 36))))
	v3 = nox_xxx_getUnitName_4E39D0(a1)
	nox_ai_debug_printf_5341A0(CString("%d: PopActionStack( %s(#%d) ) = %s@%d:\n"), nox_frame_xxx_2598000, v3, v13, v14, v15)
	v4 = int8(*(*uint8)(unsafe.Pointer(uintptr(v1 + 544))))
	if int32(v4) >= 0 {
		v5 = (*uint32)(unsafe.Pointer(uintptr(v1 + (int32(v4)+23)*24)))
		v6 = nox_xxx_monsterActionIsCondition_50A010(int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + (int32(v4)+23)*24)))))
		if v6 == 0 {
			if *(*uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(uint32(0))*5)) != 0 {
				v7 = cgoAsFunc(*(*uint32)(memmap.PtrOff(0x587000, *v5*16+0x38ED0)), (*func(int32))(nil))
				if v7 != nil {
					v7(a1)
				}
			}
		}
	}
	dword_5d4594_1599692 = 1
	v8 = int8(int32(*(*uint8)(unsafe.Pointer(uintptr(v1 + 544)))) - 1)
	v9 = int32(v8)
	*(*uint8)(unsafe.Pointer(uintptr(v1 + 544))) = uint8(v8)
	if int32(v8) >= 0 {
		v10 = (*int32)(unsafe.Pointer(uintptr(v1 + (int32(v8)*3+69)*8)))
		for {
			v11 = nox_xxx_monsterActionIsCondition_50A010(*v10)
			if v11 == 0 {
				break
			}
			v10 = (*int32)(unsafe.Add(unsafe.Pointer(v10), -int(unsafe.Sizeof(int32(0))*6)))
			v9--
			*(*uint8)(unsafe.Pointer(uintptr(v1 + 544)))--
			if v9 < 0 {
				break
			}
		}
	}
	nox_xxx_monsterActionReset_50A110((*nox_object_t)(unsafe.Pointer(uintptr(a1))))
	result = int8(*(*uint8)(unsafe.Pointer(uintptr(v1 + 544))))
	if int32(result) < 0 {
		*(*uint8)(unsafe.Pointer(uintptr(v1 + 544))) = 0
		*(*uint32)(unsafe.Pointer(uintptr(v1 + 552))) = 0
	}
	return result
}
func nox_xxx_monsterPushAction_50A260_impl(a1p *nox_object_t, a2 int32, file *byte, line int32) *int32 {
	var (
		a1     int32 = int32(uintptr(unsafe.Pointer(a1p)))
		v2     int32
		result *int32
		v4     int8
		v5     *uint32
		v6     int32
		v7     func(int32)
		v8     int8
		v9     *int32
		v10    int32
		v11    int32
		v12    int32
	)
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
	if (int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8)))) & 2) == 0 {
		return nil
	}
	v4 = int8(*(*uint8)(unsafe.Pointer(uintptr(v2 + 544))))
	if int32(v4) == 23 {
		return nil
	}
	if int32(v4) < 0 {
		goto LABEL_19
	}
	if *(*uint32)(unsafe.Pointer(uintptr(v2 + 552))) == 31 && a2 != 30 {
		return nil
	}
	v5 = (*uint32)(unsafe.Pointer(uintptr(v2 + 552 + int32(v4)*24)))
	if *(*uint32)(unsafe.Pointer(uintptr(v2 + 552 + int32(v4)*24))) != 0 || int32(v4) != 0 {
		v6 = nox_xxx_monsterActionIsCondition_50A010(int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 552 + int32(v4)*24)))))
		if v6 == 0 && *(*uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(uint32(0))*5)) != 0 {
			v7 = cgoAsFunc(*(*uint32)(memmap.PtrOff(0x587000, *v5*16+0x38ED4)), (*func(int32))(nil))
			if v7 != nil {
				v7(a1)
			}
		}
	} else {
	LABEL_19:
		*(*uint8)(unsafe.Pointer(uintptr(v2 + 544))) = math.MaxUint8
	}
	v8 = int8(int32(*(*uint8)(unsafe.Pointer(uintptr(v2 + 544)))) + 1)
	*(*uint8)(unsafe.Pointer(uintptr(v2 + 544))) = uint8(v8)
	v9 = (*int32)(unsafe.Pointer(uintptr(v2 + 552 + int32(v8)*24)))
	*v9 = a2
	*(*int32)(unsafe.Add(unsafe.Pointer(v9), unsafe.Sizeof(int32(0))*5)) = 0
	nox_xxx_monsterActionReset_50A110((*nox_object_t)(unsafe.Pointer(uintptr(a1))))
	v12 = int32(*memmap.PtrUint32(0x587000, uint32(a2*4)+0x382D0))
	v11 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 36))))
	v10 = nox_xxx_getUnitName_4E39D0(a1)
	nox_ai_debug_printf_5341A0(CString("%d: PushActionStack( %s(#%d), %s ), result: (%s:%d)\n"), nox_frame_xxx_2598000, v10, v11, v12, file, line)
	result = v9
	dword_5d4594_1599692 = 1
	return result
}
func nox_xxx_monsterAction_50A360(a1 int32, a2 int32) *int32 {
	var result *int32
	if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))&2 != 0 && *(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))) + uint32((*(*byte)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))) + 544)))+23)*24)))) != uint32(a2) {
		result = nox_xxx_monsterPushAction_50A260_impl((*nox_object_t)(unsafe.Pointer(uintptr(a1))), a2, CString(__FILE__), __LINE__)
	} else {
		result = nil
	}
	return result
}
func nox_xxx_monsterClearActionStack_50A3A0(a1 int32) {
	if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))&2 != 0 {
		for sub_5341F0((*nox_object_t)(unsafe.Pointer(uintptr(a1)))) == 0 {
			nox_xxx_monsterPopAction_50A160((*nox_object_t)(unsafe.Pointer(uintptr(a1))))
		}
	}
}
func nox_xxx_monsterCallDieFn_50A3D0(a1 *uint32) int32 {
	var (
		v1     int32
		i      int32
		result int32
		v4     int32
		v5     int32
		v6     int32
		v7     int32
		v8     int32
		v9     int32
		v10    int32
		v11    int32
		v12    int32
	)
	v1 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*187)))
	if noxflags.HasGame(noxflags.GameModeQuest) {
		sub_50E1E0(int32(uintptr(unsafe.Pointer(a1))))
	}
	for i = int32(uintptr(unsafe.Pointer(nox_xxx_getFirstPlayerUnit_4DA7C0()))); i != 0; i = int32(uintptr(unsafe.Pointer(nox_xxx_getNextPlayerUnit_4DA7F0((*nox_object_t)(unsafe.Pointer(uintptr(i))))))) {
		if (*uint32)(unsafe.Pointer(uintptr(nox_xxx_playerGetPossess_4DDF30(i)))) == a1 {
			nox_xxx_playerObserveClear_4DDEF0(i)
		}
	}
	nox_xxx_monsterClearActionStack_50A3A0(int32(uintptr(unsafe.Pointer(a1))))
	nox_xxx_monsterPushAction_50A260_impl((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1)))))), 31, CString(__FILE__), __LINE__)
	nox_xxx_monsterPushAction_50A260_impl((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1)))))), 30, CString(__FILE__), __LINE__)
	result = nox_xxx_unitIsZombie_534A40(int32(uintptr(unsafe.Pointer(a1))))
	if result == 0 {
		v4 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*4)))
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v4))), 0)) = uint8(int8(v4 & math.MaxInt8))
		*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*4)) = uint32(v4)
		nox_xxx_action_4DA9F0(a1)
		nox_xxx_unitClearBuffs_4FF580(int32(uintptr(unsafe.Pointer(a1))))
		if int32(int8(*(*uint8)(unsafe.Pointer(uintptr(v1 + 1440))))) >= 0 {
			if !noxflags.HasGame(noxflags.GameModeQuest) {
				goto LABEL_13
			}
			v5 = noxRndCounter1.IntClamp(5, 8)
		} else {
			v5 = noxRndCounter1.IntClamp(10, 20)
		}
		nox_xxx_unitSetDecayTime_511660((*nox_object_t)(unsafe.Pointer(a1)), int32(nox_gameFPS*uint32(v5)))
	LABEL_13:
		v6 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*math.MaxInt8)))
		if v6 != 0 && int32(*(*uint8)(unsafe.Pointer(uintptr(v6 + 8))))&4 != 0 {
			v7 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*3)))
			v8 = int32(*(*uint32)(unsafe.Pointer(uintptr(v6 + 748))))
			*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v7))), 0)) = uint8(int8(v7 & math.MaxInt8))
			*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*3)) = uint32(v7)
			nox_xxx_netFxShield_0_4D9200(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v8 + 276))) + 2064)))), int32(uintptr(unsafe.Pointer(a1))))
			nox_xxx_netUnmarkMinimapObj_417300(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v8 + 276))) + 2064)))), int32(uintptr(unsafe.Pointer(a1))), 1)
		}
		v9 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*3)))
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v9))), 1)) &= 254
		*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*3)) = uint32(v9)
		nox_xxx_unitTransferSlaves_4EC4B0(int32(uintptr(unsafe.Pointer(a1))))
		nox_xxx_unitClearOwner_4EC300((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1)))))))
		v10 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*3)))
		if (v10 & 8192) == 0 {
			nox_xxx_dropAllItems_4EDA40(a1)
		}
		if !noxflags.HasGame(noxflags.GameModeCoop) && *(*uint32)(unsafe.Pointer(uintptr(v1 + 2188))) == 2 && *(*uint32)(unsafe.Pointer(uintptr(v1 + 2184))) == 2 {
			if *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*130)) != 0 {
				v11 = int32(uintptr(unsafe.Pointer(nox_xxx_findParentChainPlayer_4EC580((*nox_object_t)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*130)))))))))
				if int32(*(*uint8)(unsafe.Pointer(uintptr(v11 + 8))))&4 != 0 {
					sub_4FC0B0(v11, 1)
				}
			}
		}
		result = bool2int(noxflags.HasGame(noxflags.GameModeQuest))
		if result != 0 {
			v12 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*130)))
			if v12 != 0 {
				result = int32(uintptr(unsafe.Pointer(nox_xxx_findParentChainPlayer_4EC580((*nox_object_t)(unsafe.Pointer(uintptr(v12)))))))
				if int32(*(*uint8)(unsafe.Pointer(uintptr(result + 8))))&4 != 0 {
					result = sub_4D6170(result)
				}
			}
		}
	}
	return result
}
func nox_xxx_unitUpdateMonster_50A5C0(a1 float32) int32 {
	var (
		v1     int32
		v2     int32
		v3     int8
		v4     int32
		result int32
		v6     int32
		v7     int32
		v8     int32
		v9     *uint16
		v10    int32
		v11    uint16
		v12    int32
		v13    int32
		v14    int32
		v15    int32
		v16    int32
		v17    *uint32
	)
	_ = v17
	var v18 func(int32)
	var v19 int32
	var v20 uint8
	var v21 int32
	v1 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&a1))), unsafe.Sizeof(uint32(0))*0)))
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr((*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&a1))), unsafe.Sizeof(uint32(0))*0))) + 748))))
	v3 = int8(*(*uint8)(unsafe.Pointer(uintptr(v2 + 2094))))
	if int32(v3) != 0 {
		*(*uint8)(unsafe.Pointer(uintptr(v2 + 2094))) = uint8(int8(int32(v3) - 1))
	}
	v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 2192))))
	if v4 != 0 && *(*uint32)(unsafe.Pointer(uintptr(v4 + 16)))&32800 != 0 {
		*(*uint32)(unsafe.Pointer(uintptr(v2 + 2192))) = 0
	}
	nox_xxx_mobAction_50A910(*(*int32)(unsafe.Add(unsafe.Pointer((*int32)(unsafe.Pointer(&a1))), unsafe.Sizeof(int32(0))*0)))
	nox_xxx_unitListenRoutine_50CDD0(*(*int32)(unsafe.Add(unsafe.Pointer((*int32)(unsafe.Pointer(&a1))), unsafe.Sizeof(int32(0))*0)))
	result = int32(*(*uint32)(unsafe.Pointer(uintptr((*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&a1))), unsafe.Sizeof(uint32(0))*0))) + 16))))
	if uint32(result)&0x1000000 != 0 {
		dword_5d4594_1599692 = 0
		result = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 484))))
		if result != 0 {
			if (*(*uint32)(unsafe.Pointer(uintptr((*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&a1))), unsafe.Sizeof(uint32(0))*0))) + 16))) & 32800) == 0 {
				v6 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 1440))))
				if v6&512 != 0 {
					v7 = nox_xxx_monsterGetSoundSet_424300(*(*int32)(unsafe.Add(unsafe.Pointer((*int32)(unsafe.Pointer(&a1))), unsafe.Sizeof(int32(0))*0)))
					if v7 != 0 {
						nox_xxx_aud_501960(int32(*(*uint32)(unsafe.Pointer(uintptr(v7 + 64)))), (*nox_object_t)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer((*int32)(unsafe.Pointer(&a1))), unsafe.Sizeof(int32(0))*0))))), 0, 0)
					}
					nox_xxx_scriptCallByEventBlock_502490((*int32)(unsafe.Pointer(uintptr(v2+1248))), int32(*(*uint32)(unsafe.Pointer(uintptr((*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&a1))), unsafe.Sizeof(uint32(0))*0))) + 520)))), *(*int32)(unsafe.Add(unsafe.Pointer((*int32)(unsafe.Pointer(&a1))), unsafe.Sizeof(int32(0))*0)), 17)
					nox_ai_debug_printf_5341A0(CString("%d: HP = %d/%d\n"), nox_frame_xxx_2598000, **(**uint16)(unsafe.Pointer(uintptr((*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&a1))), unsafe.Sizeof(uint32(0))*0))) + 556))), *(*uint16)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr((*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&a1))), unsafe.Sizeof(uint32(0))*0))) + 556))) + 4))))
				}
				v8 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 520))))
				if v8 != 0 && (nox_frame_xxx_2598000-uint32(v8)) >= uint32(int32(nox_gameFPS)) {
					nox_xxx_monsterPlayHurtSound_532800(*(*int32)(unsafe.Add(unsafe.Pointer((*int32)(unsafe.Pointer(&a1))), unsafe.Sizeof(int32(0))*0)))
					*(*uint32)(unsafe.Pointer(uintptr(v2 + 520))) = 0
				}
				nox_xxx_mobAction_5469B0(*(*int32)(unsafe.Add(unsafe.Pointer((*int32)(unsafe.Pointer(&a1))), unsafe.Sizeof(int32(0))*0)))
			}
			v9 = *(**uint16)(unsafe.Pointer(uintptr((*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&a1))), unsafe.Sizeof(uint32(0))*0))) + 556)))
			if v9 != nil {
				v10 = int32(*(*uint32)(unsafe.Pointer(uintptr((*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&a1))), unsafe.Sizeof(uint32(0))*0))) + 16))))
				if (v10&0x8000) == 0 && (nox_frame_xxx_2598000-*(*uint32)(unsafe.Pointer(uintptr((*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&a1))), unsafe.Sizeof(uint32(0))*0))) + 536)))) > uint32(int32(nox_gameFPS)) {
					v11 = *(*uint16)(unsafe.Add(unsafe.Pointer(v9), unsafe.Sizeof(uint16(0))*2))
					if int32(*v9) < int32(v11) && int32(v11) != 0 && (nox_frame_xxx_2598000%(nox_gameFPS*180/uint32(*(*uint16)(unsafe.Add(unsafe.Pointer(v9), unsafe.Sizeof(uint16(0))*2))))) == 0 {
						nox_xxx_unitAdjustHP_4EE460((*nox_object_t)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer((*int32)(unsafe.Pointer(&a1))), unsafe.Sizeof(int32(0))*0))))), 1)
					}
				}
			}
			nox_xxx_unitUpdateSightMB_5281F0(a1)
			nox_xxx_monsterMainAIFn_547210(*(*int32)(unsafe.Add(unsafe.Pointer((*int32)(unsafe.Pointer(&a1))), unsafe.Sizeof(int32(0))*0)))
			nox_xxx_mobActionDependency_546A70((*nox_object_t)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer((*int32)(unsafe.Pointer(&a1))), unsafe.Sizeof(int32(0))*0))))))
			nox_xxx_updateNPCAnimData_50A850(*(*int32)(unsafe.Add(unsafe.Pointer((*int32)(unsafe.Pointer(&a1))), unsafe.Sizeof(int32(0))*0)))
			v12 = int32(dword_5d4594_1599692)
			v21 = int32(dword_5d4594_1599692)
			for {
				v13 = int32(*(*byte)(unsafe.Pointer(uintptr(v2 + 544))))
				v14 = v13*3 + 69
				v13 *= 3
				v15 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + v14*8))))
				v16 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + v13*8 + 572))))
				v17 = (*uint32)(unsafe.Pointer(uintptr(v2 + v13*8 + 572)))
				if v16 != 0 {
					break
				}
				*v17 = 1
				dword_5d4594_1599692 = 0
				v18 = cgoAsFunc(*(*uint32)(memmap.PtrOff(0x587000, uint32(v15*16)+233160)), (*func(int32))(nil))
				if v18 != nil {
					v18(v1)
					v12 = int32(dword_5d4594_1599692)
					if dword_5d4594_1599692 != 0 {
						continue
					}
				}
				goto LABEL_30
			}
			if v12 != 0 {
				goto LABEL_31
			}
		LABEL_30:
			dword_5d4594_1599692 = uint32(v21)
		LABEL_31:
			(cgoAsFunc(*(*uint32)(memmap.PtrOff(0x587000, uint32(v15*16)+0x38ECC)), (*func(int32))(nil)))(v1)
			if dword_5d4594_1599692 != 0 {
				nox_ai_debug_printStack_509F60((*nox_object_t)(unsafe.Pointer(uintptr(v1))), CString("stack changed"))
			}
			v19 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 1440))))
			*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v19))), 1)) &= 253
			*(*uint32)(unsafe.Pointer(uintptr(v2 + 1440))) = uint32(v19)
			nox_xxx_monsterPolygonEnter_421FF0(v1)
			v20 = *(*uint8)(unsafe.Pointer(uintptr(v2 + 1128)))
			if int32(v20) < 100 {
				*(*uint8)(unsafe.Pointer(uintptr(v2 + 1128))) = uint8(uint32(v20) + 100/nox_gameFPS)
			}
			if nox_xxx_unitIsMimic_534840(v1) != 0 {
				nox_xxx_monsterMimicCheckMorph_534950(v1)
			}
			result = int32(nox_gameFPS)
			if nox_frame_xxx_2598000-*(*uint32)(unsafe.Pointer(uintptr(v1 + 536))) > (nox_gameFPS * 3) {
				*(*uint32)(unsafe.Pointer(uintptr(v2 + 1440))) &= 0xFFF7FFFF
			}
		}
	}
	return result
}
func nox_xxx_updateNPCAnimData_50A850(a1 int32) int8 {
	var (
		v1 *uint8
		v2 *uint8
		v3 uint8
		v4 uint8
		v5 uint8
		v6 uint8
	)
	v1 = *(**uint8)(unsafe.Pointer(uintptr(a1 + 748)))
	if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 12))))&16 != 0 {
		v2 = *(**uint8)(unsafe.Add(unsafe.Pointer(v1), int32(int8(*(*uint8)(unsafe.Add(unsafe.Pointer(v1), 544))))*24+552))
		if uintptr(unsafe.Pointer(v2)) == uintptr(16) || uintptr(unsafe.Pointer(v2)) == uintptr(17) {
			*(*uint8)(unsafe.Add(unsafe.Pointer(v1), 483)) = 0
			return int8(uintptr(unsafe.Pointer(v2)))
		}
	}
	*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v2))), 0)) = *(*uint8)(unsafe.Add(unsafe.Pointer(v1), 483))
	if int32(uint8(uintptr(unsafe.Pointer(v2)))) == 0 {
		v2 = nox_xxx_unitNPCActionToAnim_533D00(a1)
		if v2 != nil {
			*(*uint8)(unsafe.Add(unsafe.Pointer(v1), 480)) = *(*uint8)(unsafe.Add(unsafe.Pointer(v2), 9))
			if int32(*(*uint8)(unsafe.Add(unsafe.Pointer(v2), 9))) == 0 {
				goto LABEL_13
			}
			v3 = uint8(int8(int32(*(*uint8)(unsafe.Add(unsafe.Pointer(v1), 482))) + 1))
			*(*uint8)(unsafe.Add(unsafe.Pointer(v1), 482)) = v3
			if int32(v3) >= int32(*(*uint8)(unsafe.Add(unsafe.Pointer(v2), 10)))+1 {
				v4 = *(*uint8)(unsafe.Add(unsafe.Pointer(v1), 481))
				*(*uint8)(unsafe.Add(unsafe.Pointer(v1), 482)) = 0
				*(*uint8)(unsafe.Add(unsafe.Pointer(v1), 481)) = func() uint8 {
					p := &v4
					*p++
					return *p
				}()
				v5 = v4
				v6 = *(*uint8)(unsafe.Add(unsafe.Pointer(v2), 9))
				if int32(v5) >= int32(v6) {
					if *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), unsafe.Sizeof(uint32(0))*3))) != 0 {
						*(*uint8)(unsafe.Add(unsafe.Pointer(v1), 481)) = 0
						return int8(uintptr(unsafe.Pointer(v2)))
					}
					*(*uint8)(unsafe.Add(unsafe.Pointer(v1), 481)) = uint8(int8(int32(v6) - 1))
				LABEL_13:
					*(*uint8)(unsafe.Add(unsafe.Pointer(v1), 483)) = 1
					return int8(uintptr(unsafe.Pointer(v2)))
				}
			}
		}
	}
	return int8(uintptr(unsafe.Pointer(v2)))
}
func nox_xxx_mobAction_50A910(a1 int32) int32 {
	var (
		v1     int32
		v2     int32
		result int32
		v4     *uint32
		v5     int32
		v6     *uint32
		v7     int32
		v8     int32
		v9     int32
		v10    int32
		v11    int32
		v12    int32
		v13    int32
	)
	v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
	v12 = v1
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 1216))))
	if v2 != 0 && *(*uint32)(unsafe.Pointer(uintptr(v2 + 16)))&32800 != 0 {
		*(*uint32)(unsafe.Pointer(uintptr(v1 + 1216))) = 0
	}
	result = int32(*(*byte)(unsafe.Pointer(uintptr(v1 + 544))))
	if result >= 0 {
		v4 = (*uint32)(unsafe.Pointer(uintptr(v1 + (result*3+69)*8)))
		v13 = result + 1
		for {
			v5 = 0
			if *memmap.PtrUint32(0x587000, *v4*16+0x383F4) > 0 {
				v6 = (*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*1))
				for {
					if *memmap.PtrUint32(0x587000, (uint32(v5)+*v4*4)*4+0x383F8) == 1 && *v6 != 0 {
						sub_509FF0(int32(uintptr(unsafe.Pointer(v6))))
					}
					v5++
					v6 = (*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*2))
					if v5 >= *memmap.PtrInt32(0x587000, *v4*16+0x383F4) {
						break
					}
				}
				v1 = v12
			}
			switch *v4 {
			case 3:
				v7 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*3)))
				if v7 != 0 {
					goto LABEL_26
				}
			case 7:
				fallthrough
			case 8:
				v10 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*3)))
				if v10 != 0 {
					if nox_xxx_unitCanInteractWith_5370E0((*nox_object_t)(unsafe.Pointer(uintptr(a1))), (*nox_object_t)(unsafe.Pointer(uintptr(v10))), 0) != 0 || nox_xxx_checkMobAction_50A0D0((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 3) != 0 {
						v7 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*3)))
						v11 = int32(*(*uint32)(unsafe.Pointer(uintptr(v7 + 56))))
						goto LABEL_27
					}
					*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*3)) = 0
				}
			case 15:
				v7 = int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 1196))))
				if v7 != 0 {
				LABEL_26:
					v11 = int32(*(*uint32)(unsafe.Pointer(uintptr(v7 + 56))))
				LABEL_27:
					*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*1)) = uint32(v11)
					*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*2)) = *(*uint32)(unsafe.Pointer(uintptr(v7 + 60)))
				}
			case 17:
				v8 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*3)))
				if v8 != 0 && nox_xxx_unitCanInteractWith_5370E0((*nox_object_t)(unsafe.Pointer(uintptr(a1))), (*nox_object_t)(unsafe.Pointer(uintptr(v8))), 0) != 0 {
					v9 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*3)))
					*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*1)) = *(*uint32)(unsafe.Pointer(uintptr(v9 + 56)))
					*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*2)) = *(*uint32)(unsafe.Pointer(uintptr(v9 + 60)))
				}
			default:
			}
			v4 = (*uint32)(unsafe.Add(unsafe.Pointer(v4), -int(unsafe.Sizeof(uint32(0))*6)))
			result = func() int32 {
				p := &v13
				*p--
				return *p
			}()
			if v13 == 0 {
				break
			}
			v1 = v12
		}
	}
	return result
}
func sub_50AA90(a1 int32) int32 {
	var (
		v1 int32
		v2 int32
		i  *uint32
		v5 int32
	)
	v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
	v2 = int32(*(*byte)(unsafe.Pointer(uintptr(v1 + 544))))
	if v2 < 0 {
		return 0
	}
	for i = (*uint32)(unsafe.Pointer(uintptr(v1 + (v2*3+69)*8))); *i != 3; i = (*uint32)(unsafe.Add(unsafe.Pointer(i), -int(unsafe.Sizeof(uint32(0))*6))) {
		if func() int32 {
			p := &v2
			*p--
			return *p
		}() < 0 {
			return 0
		}
	}
	v5 = v1 + v2*24 + 564
	if *(*uint32)(unsafe.Pointer(uintptr(v5))) != 0 {
		sub_509FF0(v5)
	}
	return int32(*(*uint32)(unsafe.Pointer(uintptr(v5))))
}
func nox_xxx_minimapFirstMonster_50AAE0() int32 {
	var v0 int32
	v0 = int32(uintptr(unsafe.Pointer(noxServer.firstServerObject())))
	dword_5d4594_1599696 = uint32(v0)
	if v0 == 0 {
		return 0
	}
	for (int32(*(*uint8)(unsafe.Pointer(uintptr(v0 + 8)))) & 2) == 0 {
		v0 = int32(uintptr(unsafe.Pointer((*nox_object_t)(unsafe.Pointer(uintptr(v0))).Next())))
		dword_5d4594_1599696 = uint32(v0)
		if v0 == 0 {
			return 0
		}
	}
	return v0 + 56
}
func nox_xxx_minimapNextMonster_50AB10() int32 {
	var v0 int32
	v0 = int32(uintptr(unsafe.Pointer((*nox_object_t)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1599696))))).Next())))
	dword_5d4594_1599696 = uint32(v0)
	if v0 == 0 {
		return 0
	}
	for (int32(*(*uint8)(unsafe.Pointer(uintptr(v0 + 8)))) & 2) == 0 {
		v0 = int32(uintptr(unsafe.Pointer((*nox_object_t)(unsafe.Pointer(uintptr(v0))).Next())))
		dword_5d4594_1599696 = uint32(v0)
		if v0 == 0 {
			return 0
		}
	}
	return v0 + 56
}
func sub_50AB50(a1 int32, a2 int32) int16 {
	if a1 < 0 || a1 >= 256 || a2 < 0 || a2 >= 256 {
		return 0
	}
	return int16(nox_server_xxx_1599716[a2+(a1<<8)].field_8)
}
func nox_xxx_allocVisitNodesArray_50AB90() int32 {
	var result int32
	result = int32(uintptr(unsafe.Pointer(nox_new_alloc_class(CString("VisitNodes"), 16, 1024))))
	nox_alloc_visitNode_2386184 = unsafe.Pointer(uintptr(result))
	if result != 0 {
		dword_5d4594_2386176 = uint32(uintptr(alloc.Calloc(1, 8192)))
		if dword_5d4594_2386176 != 0 {
			dword_5d4594_2386172 = 0
			result = 1
		} else {
			nox_free_alloc_class((*nox_alloc_class)(nox_alloc_visitNode_2386184))
			result = 0
		}
	}
	return result
}
func sub_50ABF0() {
	*(*unsafe.Pointer)(unsafe.Pointer(&dword_5d4594_2386176)) = nil
	nox_free_alloc_class((*nox_alloc_class)(nox_alloc_visitNode_2386184))
	nox_alloc_visitNode_2386184 = nil
}
func sub_50AC20(a3 int32, a2 *uint16) int32 {
	var (
		v2     *uint16
		v3     uint16
		v4     uint16
		v5     uint16
		v6     int16
		v7     *uint32
		v8     *uint16
		result int32
		v10    int32
		v11    int32
		v12    int32
		v13    int16
		v14    *uint16
		v15    int16
		v16    *uint16
		v17    float32
		v18    float32
		v19    float32
		v20    float32
		v21    int32
		a1     float2
	)
	v2 = (*uint16)(unsafe.Pointer(uintptr(a3)))
	v3 = *(*uint16)(unsafe.Pointer(uintptr(a3 + 2)))
	a3 = int32(*(*uint16)(unsafe.Pointer(uintptr(a3))))
	v21 = int32(v3)
	if (int32(nox_server_xxx_1599716[int32(v3)+(a3<<8)].field_8) & 60) == 0 {
		return 0
	}
	dword_5d4594_2386152 = 0
	v4 = *v2
	v5 = *(*uint16)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint16(0))*1))
	a1.field_0 = float32(float64(a3) * 23.0)
	v6 = int16(nox_server_xxx_1599716[int32(v5)+(int32(v4)<<8)].field_8)
	a1.field_4 = float32(float64(v21) * 23.0)
	if int32(v6)&16 != 0 {
		a3 = 2048
		sub_517B70(&a1, func(arg1 uint32, arg2 int32) {
			sub_50AE80(int32(arg1), arg2)
		}, int32(uintptr(unsafe.Pointer(&a3))))
		if dword_5d4594_2386152 != 0 {
			v7 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_2386152 + 700))) + 8)))
			if *(*int32)(unsafe.Pointer(uintptr(dword_5d4594_2386152 + 700))) != -8 {
				if *(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_2386152 + 16)))&0x1000000 != 0 {
					v8 = a2
					*a2 = uint16(int16(int32(int16(*(*int32)(unsafe.Pointer(v7)))) / 23))
					result = 1
					*(*uint16)(unsafe.Add(unsafe.Pointer(v8), unsafe.Sizeof(uint16(0))*1)) = uint16(int16(int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v7), unsafe.Sizeof(uint32(0))*1))) / 23))
					return result
				}
			}
		}
		return 0
	}
	if int32(v6)&32 != 0 {
		a3 = 1024
		sub_517B70(&a1, func(arg1 uint32, arg2 int32) {
			sub_50AE80(int32(arg1), arg2)
		}, int32(uintptr(unsafe.Pointer(&a3))))
		v10 = int32(dword_5d4594_2386152)
		if dword_5d4594_2386152 == 0 {
			return 0
		}
		v11 = int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_2386152 + 748))) + 12))))
		goto LABEL_18
	}
	if int32(v6)&4 != 0 {
		a3 = 0x4000
		sub_517B70(&a1, func(arg1 uint32, arg2 int32) {
			sub_50AE80(int32(arg1), arg2)
		}, int32(uintptr(unsafe.Pointer(&a3))))
		if dword_5d4594_2386152 != 0 {
			v12 = int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_2386152 + 748))) + 4))))
			if v12 != 0 {
				if *(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_2386152 + 16)))&0x1000000 != 0 {
					v17 = float32(float64(*(*float32)(unsafe.Pointer(uintptr(v12 + 56)))) * 0.043478262)
					v13 = int16(int(v17))
					v14 = a2
					*a2 = uint16(v13)
					v18 = float32(float64(*(*float32)(unsafe.Pointer(uintptr(v12 + 60)))) * 0.043478262)
					*(*uint16)(unsafe.Add(unsafe.Pointer(v14), unsafe.Sizeof(uint16(0))*1)) = uint16(int16(int(v18)))
					return 1
				}
			}
		}
		return 0
	}
	if int32(v6)&8 != 0 {
		a3 = 0x8000
		sub_517B70(&a1, func(arg1 uint32, arg2 int32) {
			sub_50AE80(int32(arg1), arg2)
		}, int32(uintptr(unsafe.Pointer(&a3))))
		v10 = int32(dword_5d4594_2386152)
		if dword_5d4594_2386152 != 0 {
			v11 = int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_2386152 + 748))) + 4))))
		LABEL_18:
			if v11 != 0 && *(*uint32)(unsafe.Pointer(uintptr(v10 + 16)))&0x1000000 != 0 {
				v19 = float32(float64(*(*float32)(unsafe.Pointer(uintptr(v11 + 56)))) * 0.043478262)
				v15 = int16(int(v19))
				v16 = a2
				*a2 = uint16(v15)
				v20 = float32(float64(*(*float32)(unsafe.Pointer(uintptr(v11 + 60)))) * 0.043478262)
				*(*uint16)(unsafe.Add(unsafe.Pointer(v16), unsafe.Sizeof(uint16(0))*1)) = uint16(int16(int(v20)))
				return 1
			}
			return 0
		}
	}
	return 0
}
func sub_50AE80(a1 int32, a2 int32) {
	if *(*uint32)(unsafe.Pointer(uintptr(a1 + 8)))&*(*uint32)(unsafe.Pointer(uintptr(a2))) != 0 {
		dword_5d4594_2386152 = uint32(a1)
	}
}
func sub_50AEA0(a1 int32, a2 *float2, a3 *uint32) int32 {
	var (
		v3  *uint16
		v4  uint16
		v5  uint16
		v6  uint16
		v7  int16
		v9  int32
		a1a float2
	)
	v3 = (*uint16)(unsafe.Pointer(uintptr(a1)))
	v4 = *(*uint16)(unsafe.Pointer(uintptr(a1 + 2)))
	a1 = int32(*(*uint16)(unsafe.Pointer(uintptr(a1))))
	v9 = int32(v4)
	if (int32(nox_server_xxx_1599716[int32(v4)+(a1<<8)].field_8) & 60) == 0 {
		return 0
	}
	dword_5d4594_2386152 = 0
	v5 = *v3
	v6 = *(*uint16)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint16(0))*1))
	a1a.field_0 = float32(float64(a1) * 23.0)
	v7 = int16(nox_server_xxx_1599716[int32(v6)+(int32(v5)<<8)].field_8)
	a1a.field_4 = float32(float64(v9) * 23.0)
	if int32(v7)&16 != 0 {
		a1 = 2048
	} else if int32(v7)&32 != 0 {
		a1 = 1024
	} else if int32(v7)&4 != 0 {
		a1 = 0x4000
	} else if int32(v7)&8 != 0 {
		a1 = 0x8000
	}
	sub_517B70(&a1a, func(arg1 uint32, arg2 int32) {
		sub_50AE80(int32(arg1), arg2)
	}, int32(uintptr(unsafe.Pointer(&a1))))
	if dword_5d4594_2386152 == 0 || (*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_2386152 + 16)))&0x1000000) == 0 {
		return 0
	}
	*a2 = *(*float2)(unsafe.Pointer(uintptr(dword_5d4594_2386152 + 56)))
	*a3 = *(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_2386152 + 36)))
	return 1
}
func nox_xxx_aiTestUnitDangerous_50B2C0(a1 int32) int32 {
	var (
		v1     int32
		result int32
		v3     float64
		v4     int32
		v5     float64
		v6     int32
		v7     int32
		v8     int32
		v9     int32
		v10    *uint8
		v11    *float32
		v12    int32
		v13    int16
		v14    float32
		v15    float32
		v16    float32
		v17    float32
		v18    int32
		v19    int32
		v20    int32
		v21    float32
		v22    float32
		a2     float2
		v24    int32
		v25    [60]byte
		a1a    int32
	)
	v1 = a1
	result = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 8))))
	if (result&0xC080) == 0 && (result&8192 != 0 || (int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 16))))&73) == 0) && (result&8200 != 0 || int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 16))))&2 != 0) {
		if uint32(result)&0x10000 != 0 {
			libc.MemCpy(unsafe.Pointer(&v25[0]), unsafe.Pointer(uintptr(a1+172)), int(unsafe.Sizeof([60]byte{})))
			if *(*uint32)(unsafe.Pointer(uintptr(a1 + 172))) == 2 {
				v3 = float64(*(*float32)(unsafe.Pointer(&dword_587000_234176)) + *(*float32)(unsafe.Pointer(uintptr(a1 + 176))))
				*(*float32)(unsafe.Pointer(uintptr(a1 + 176))) = float32(v3)
				*(*float32)(unsafe.Pointer(uintptr(a1 + 180))) = float32(v3 * v3)
			} else if *(*uint32)(unsafe.Pointer(uintptr(a1 + 172))) == 3 {
				*(*float32)(unsafe.Pointer(uintptr(a1 + 184))) = *(*float32)(unsafe.Pointer(&dword_587000_234176)) + *(*float32)(unsafe.Pointer(&dword_587000_234176)) + *(*float32)(unsafe.Pointer(uintptr(a1 + 184)))
				*(*float32)(unsafe.Pointer(uintptr(a1 + 188))) = *(*float32)(unsafe.Pointer(&dword_587000_234176)) + *(*float32)(unsafe.Pointer(&dword_587000_234176)) + *(*float32)(unsafe.Pointer(uintptr(a1 + 188)))
				nox_shape_box_calc((*nox_shape)(unsafe.Pointer(uintptr(a1 + 172))))
			}
			nox_xxx_objectUnkUpdateCoords_4E7290(a1)
		}
		v14 = float32(float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 232)))) * 0.043478262)
		v4 = int(v14)
		v5 = float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 236)))) * 0.043478262
		v6 = v4
		v24 = v4
		v15 = float32(v5)
		v7 = int(v15)
		v16 = float32(float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 240)))) * 0.043478262)
		v19 = int(v16)
		v17 = float32(float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 244)))) * 0.043478262)
		v8 = int(v17)
		v9 = v7
		v20 = v8
		for a1a = v7; v9 <= v8; a1a = v9 {
			v18 = v6
			if v6 <= v19 {
				v10 = (*uint8)(unsafe.Pointer(&nox_server_xxx_1599716[v9+(v6<<8)].field_8))
				for {
					v11 = mem_getFloatPtr(0x587000, 0x3927C)
					for {
						v21 = float32(float64(v18) * 23.0)
						a2.field_0 = v21 + *((*float32)(unsafe.Add(unsafe.Pointer(v11), -int(unsafe.Sizeof(float32(0))*1))))
						v22 = float32(float64(a1a) * 23.0)
						a2.field_4 = v22 + *v11
						if sub_547DB0(v1, &a2) != 0 {
							break
						}
						v11 = (*float32)(unsafe.Add(unsafe.Pointer(v11), unsafe.Sizeof(float32(0))*2))
						if int32(uintptr(unsafe.Pointer(v11))) >= int32(uintptr(memmap.PtrOff(0x587000, 234180))) {
							goto LABEL_25
						}
					}
					*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v10))), -int(unsafe.Sizeof(uint32(0))*1)))) = dword_5d4594_2386164
					v12 = int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 8))))
					if v12&8 != 0 || int32(*(*uint8)(unsafe.Pointer(uintptr(v1 + 16))))&2 != 0 {
						v13 = int16(int32(*(*uint16)(unsafe.Pointer(v10))) | 256)
						*(*uint16)(unsafe.Pointer(v10)) = uint16(v13)
						if (int32(*(*uint8)(unsafe.Pointer(uintptr(v1 + 16)))) & 16) == 0 {
							*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v13))), unsafe.Sizeof(int16(0))-1)) |= 2
							*(*uint16)(unsafe.Pointer(v10)) = uint16(v13)
						}
					} else if v12&8192 != 0 {
						*(*uint8)(unsafe.Add(unsafe.Pointer(v10), 1)) |= 4
					}
				LABEL_25:
					v6++
					v10 = (*uint8)(unsafe.Add(unsafe.Pointer(v10), 3072))
					v18 = v6
					if v6 > v19 {
						break
					}
				}
				v9 = a1a
				v8 = v20
				v6 = v24
			}
			v9++
		}
		result = int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 8))))
		if uint32(result)&0x10000 != 0 {
			libc.MemCpy(unsafe.Pointer(uintptr(v1+172)), unsafe.Pointer(&v25[0]), 60)
			result = nox_xxx_objectUnkUpdateCoords_4E7290(v1)
		}
	}
	return result
}
func sub_50B500() {
	dword_5d4594_2386168 = 0
}
func sub_50B510() int32 {
	var result int32
	result = 0
	dword_5d4594_2386168 = 0
	dword_5d4594_2386172 = 0
	return result
}
func sub_50B520() int32 {
	var (
		result int32
		i      int32
	)
	result = int32(nox_frame_xxx_2598000)
	if (nox_frame_xxx_2598000 - dword_5d4594_2386172) >= 15 {
		dword_5d4594_2386172 = nox_frame_xxx_2598000
		dword_5d4594_2386164++
		result = int32(uintptr(unsafe.Pointer(noxServer.firstServerObject())))
		for i = result; result != 0; i = result {
			nox_xxx_aiTestUnitDangerous_50B2C0(i)
			result = int32(uintptr(unsafe.Pointer((*nox_object_t)(unsafe.Pointer(uintptr(i))).Next())))
		}
		dword_5d4594_2386168 = 1
	}
	return result
}
func sub_50B810(a1 int32, a2 *float32) int32 {
	var (
		v2     int32
		v3     int32
		result int32
		v5     float32
		v6     float32
	)
	v5 = float32(float64(*a2) * 0.043478262)
	v2 = int(v5)
	v6 = float32(float64(*(*float32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(float32(0))*1))) * 0.043478262)
	v3 = int(v6)
	if sub_50B870(a1, v2, v3) != 0 {
		result = 0
	} else {
		result = nox_xxx_pathfind_preCheckWalls2_50B8A0(a1, v2, v3)
	}
	return result
}
func sub_50B870(a1 int32, a2 int32, a3 int32) int32 {
	return bool2int(int32(sub_57B630(a1, a2, a3)) != -1)
}
func nox_xxx_pathfind_preCheckWalls2_50B8A0(a1 int32, a2 int32, a3 int32) int32 {
	var result int32
	if sub_50B950(a1, a2, a3) != 0 {
		result = 0
	} else {
		result = bool2int(sub_50B8E0(a1, a2, a3) == 0)
	}
	return result
}
func sub_50B8E0(a1 int32, a2 int32, a3 int32) uint32 {
	var (
		v4  uint32
		v5  int32
		ind int32 = (a3 + (a2 << 8))
	)
	if nox_server_xxx_1599716[ind].field_4 != dword_5d4594_2386164 {
		return 0
	}
	*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v4))), unsafe.Sizeof(uint16(0))*1)) = *(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&a1))), unsafe.Sizeof(uint16(0))*1))
	v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 16))))
	if v5&0x4000 != 0 {
		*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v4))), unsafe.Sizeof(uint16(0))*0)) = nox_server_xxx_1599716[ind].field_8
		return (v4 >> 9) & 1
	}
	if int32(nox_server_xxx_1599716[ind].field_8)&256 != 0 {
		return 1
	}
	if sub_534020(a1) != 0 || (int32(nox_server_xxx_1599716[ind].field_8)&1024) == 0 {
		return 0
	}
	return 1
}
func sub_50B950(a1 int32, a2 int32, a3 int32) int32 {
	var (
		v3     int32
		result int32
	)
	v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 16))))
	if v3&0x4000 != 0 {
		result = (int32(nox_server_xxx_1599716[a3+(a2<<8)].field_8) >> 1) & 1
	} else {
		result = int32(nox_server_xxx_1599716[a3+(a2<<8)].field_8) & 1
	}
	return result
}
func nox_xxx_genPathToPoint_50B9A0(a1 int32, a2 int32, a3 int32, a4 *float32) int32 {
	*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a3 + 748))) + 2172))) = 0
	nox_xxx_pathFind_50BA00(0, a3, (*float32)(unsafe.Pointer(uintptr(a3+56))), a4, nil, 0)
	if dword_5d4594_2386180 > uint32(a2) && dword_5d4594_1599712 == 0 {
		dword_5d4594_1599712 = 1
	}
	return nox_xxx_appendWorkPath_50C990(a1, 0, a2)
}
func nox_xxx_pathFind_50BA00(a1 int32, a2 int32, a3 *float32, a4 *float32, a5 func(int32, int32, int32) int32, a6 int32) unsafe.Pointer {
	var (
		v6     int32
		v7     float64
		v8     int32
		v9     float64
		result unsafe.Pointer
		v11    *uint16
		v12    uint16
		v13    uint16
		v14    *uint16
		v15    *uint16
		v16    int32
		v17    int32
		v18    int32
		v19    int32
		v20    int32
		v21    int32
		v22    uint16
		v23    int32
		v24    int32
		v25    uint16
		v26    int32
		v27    float32
		v28    float32
		v29    int32
		v30    int32
		v31    *uint16
		v32    *uint16
		v33    *uint16
		v34    *uint16
		v35    *uint16
		v36    int32
		v37    int32
		v38    bool
		v39    int32
		v40    int32
		v41    int32
		v42    int32
		v43    int32
		v44    int32
		v45    *uint16
		v46    *uint16
		v47    int32
		v48    int32
		v49    int32
		v50    int32
		v51    float32
		v52    float32
		v53    float32
		v54    float32
		v55    int32
		v56    int32
		v57    int32
		v58    int32
		v59    int32
		v60    int8
		v61    int32
		i      int32
		v63    int32
		v64    int32
		v65    int32
		v66    int32
		v67    int32
		v68    [2]int16
		v69    int32
		v70    int32
		v71    *uint8
		v72    *uint8
		v73    int32
		v74    int32
		a2a    int2
		v76    float4
		v77    *uint16
	)
	v67 = 0
	v64 = 0
	v65 = 0
	if a1 != 0 {
		v59 = 0
	} else {
		v59 = 0xF423F
	}
	dword_5d4594_1599712 = 0
	dword_5d4594_2386160++
	if dword_5d4594_2386168 == 0 {
		sub_50B520()
	}
	v51 = float32(float64(*a3) * 0.043478262)
	v6 = int(v51)
	v7 = float64(*(*float32)(unsafe.Add(unsafe.Pointer(a3), unsafe.Sizeof(float32(0))*1))) * 0.043478262
	a2a.field_0 = v6
	v52 = float32(v7)
	a2a.field_4 = int(v52)
	nox_xxx_pathfind_preCheckWalls_50C8D0(a2, &a2a)
	v61 = 0
	v63 = bool2int(nox_xxx_pathfind_preCheckWalls2_50B8A0(a2, a2a.field_0, a2a.field_4) == 0)
	if a5 != nil && a5(a2, a2a.field_0, a2a.field_4) == 0 {
		v61 = 1
	}
	v53 = float32(float64(*a4) * 0.043478262)
	v8 = int(v53)
	v9 = float64(*(*float32)(unsafe.Add(unsafe.Pointer(a4), unsafe.Sizeof(float32(0))*1))) * 0.043478262
	v73 = v8
	v54 = float32(v9)
	v74 = int(v54)
	result = nox_alloc_visitNode_2386184
	if nox_alloc_visitNode_2386184 == nil {
		dword_5d4594_2386180 = 0
		dword_5d4594_1599712 = 2
		return result
	}
	nox_alloc_class_free_all((*nox_alloc_class)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&nox_alloc_visitNode_2386184)))))
	v11 = (*uint16)(nox_alloc_class_new_obj_zero((*nox_alloc_class)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&nox_alloc_visitNode_2386184))))))
	*v11 = uint16(int16(a2a.field_0))
	*(*uint16)(unsafe.Add(unsafe.Pointer(v11), unsafe.Sizeof(uint16(0))*1)) = uint16(int16(a2a.field_4))
	v12 = *v11
	v13 = *(*uint16)(unsafe.Add(unsafe.Pointer(v11), unsafe.Sizeof(uint16(0))*1))
	*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v11))), unsafe.Sizeof(uint32(0))*1))) = 0
	*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v11))), unsafe.Sizeof(uint32(0))*2))) = 0
	v14 = v11
	v66 = 0
	nox_server_xxx_1599716[int32(v13)+(int32(v12)<<8)].field_0 = dword_5d4594_2386160
	for 2 != 0 {
		v15 = v14
		v14 = nil
		v77 = nil
		if v15 == nil {
			goto LABEL_40
		}
		for 2 != 0 {
			v16 = int32(*v15)
			v17 = int32(*(*uint16)(unsafe.Add(unsafe.Pointer(v15), unsafe.Sizeof(uint16(0))*1)))
			v18 = (int32(uint16(int16(v16)))-v73)*(int32(uint16(int16(v16)))-v73) + (int32(uint16(int16(v17)))-v74)*(int32(uint16(int16(v17)))-v74)
			if a1 != 0 {
				if v18 > v59 {
					v67 = int32(uintptr(unsafe.Pointer(v15)))
					v59 = (int32(uint16(int16(v16)))-v73)*(int32(uint16(int16(v16)))-v73) + (int32(uint16(int16(v17)))-v74)*(int32(uint16(int16(v17)))-v74)
				}
			} else {
				if v18 < v59 {
					v67 = int32(uintptr(unsafe.Pointer(v15)))
					v59 = (int32(uint16(int16(v16)))-v73)*(int32(uint16(int16(v16)))-v73) + (int32(uint16(int16(v17)))-v74)*(int32(uint16(int16(v17)))-v74)
				}
				if v63 == 0 && v61 == 0 && v16 == v73 && v17 == v74 {
					dword_5d4594_1599712 = 0
					return unsafe.Pointer(uintptr(sub_50C320(a2, int32(uintptr(unsafe.Pointer(v15))), (*uint32)(unsafe.Pointer(a4)))))
				}
			}
			v69 = noxRndCounter1.IntClamp(0, 7)
			for i = 0; i < 8; i++ {
				v19 = (i + v69) % 8
				v20 = int32(*memmap.PtrUint32(0x587000, uint32(v19*8)+0x392E8))
				v21 = int32(*memmap.PtrUint32(0x587000, uint32(v19*8)+234220))
				v22 = *v15
				v72 = (*uint8)(memmap.PtrOff(0x587000, uint32(v19*8)+0x392E8))
				v23 = v20 + int32(v22)
				v24 = v21 + int32(*(*uint16)(unsafe.Add(unsafe.Pointer(v15), unsafe.Sizeof(uint16(0))*1)))
				v71 = (*uint8)(memmap.PtrOff(0x587000, uint32(v19*8)+234220))
				if v23 < 0 || v23 >= 256 || v24 < 0 || v24 >= 256 || nox_server_xxx_1599716[v24+(v23<<8)].field_0 == dword_5d4594_2386160 {
					continue
				}
				nox_server_xxx_1599716[v24+(v23<<8)].field_0 = dword_5d4594_2386160
				if v23 == v73 && v24 == v74 {
					v25 = *(*uint16)(unsafe.Add(unsafe.Pointer(v15), unsafe.Sizeof(uint16(0))*1))
					v70 = int32(*v15)*23 + 11
					v76.field_0 = float32(float64(v70))
					v26 = int32(v25) * 23
					v27 = *a4
					v28 = *(*float32)(unsafe.Add(unsafe.Pointer(a4), unsafe.Sizeof(float32(0))*1))
					v70 = v26 + 11
					v76.field_8 = v27
					v76.field_4 = float32(float64(v26 + 11))
					v76.field_C = v28
					if nox_xxx_mapTraceObstacles_50B580((*nox_object_t)(unsafe.Pointer(uintptr(a2))), &v76) != 0 {
						v29 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 16))))
						if v29&0x4000 != 0 {
							if nox_xxx_mapTraceRay_535250(&v76, nil, nil, 5) != 0 {
								goto LABEL_31
							}
						} else if nox_xxx_mapTraceRay_535250(&v76, nil, nil, 1) != 0 {
						LABEL_31:
							v30 = a2
							goto LABEL_32
						}
					}
				}
				if v19 < 4 {
					v30 = a2
				} else {
					v60 = int8(int32(uint16(int16(^(int32(*(*uint16)(unsafe.Pointer(uintptr(a2 + 16))))>>8))))&216 | 152)
					sub_57B4D0(a2)
					switch v19 {
					case 4:
						v36 = int32(uint8(sub_57B500(int32(*v15), int32(*(*uint16)(unsafe.Add(unsafe.Pointer(v15), unsafe.Sizeof(uint16(0))*1)))+1, v60)))
						if v36 != 1 && v36 != 6 && v36 != 10 && v36 != 9 && v36 != math.MaxUint8 {
							continue
						}
						*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v37))), 0)) = uint8(sub_57B500(int32(*v15)+1, int32(*(*uint16)(unsafe.Add(unsafe.Pointer(v15), unsafe.Sizeof(uint16(0))*1))), v60))
						goto LABEL_54
					case 5:
						v39 = int32(uint8(sub_57B500(int32(*v15)-1, int32(*(*uint16)(unsafe.Add(unsafe.Pointer(v15), unsafe.Sizeof(uint16(0))*1))), v60)))
						if v39 != 1 && v39 != 6 && v39 != 10 && v39 != 9 && v39 != math.MaxUint8 {
							continue
						}
						*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v37))), 0)) = uint8(sub_57B500(int32(*v15), int32(*(*uint16)(unsafe.Add(unsafe.Pointer(v15), unsafe.Sizeof(uint16(0))*1)))-1, v60))
					LABEL_54:
						v37 = int32(uint8(int8(v37)))
						if int32(uint8(int8(v37))) == 1 || v37 == 4 || v37 == 7 {
							goto LABEL_86
						}
						v38 = v37 == 8
					case 6:
						v40 = int32(uint8(sub_57B500(int32(*v15), int32(*(*uint16)(unsafe.Add(unsafe.Pointer(v15), unsafe.Sizeof(uint16(0))*1)))+1, v60)))
						if v40 != 0 && v40 != 5 && v40 != 9 && v40 != 8 && v40 != math.MaxUint8 {
							continue
						}
						v37 = int32(uint8(sub_57B500(int32(*v15)-1, int32(*(*uint16)(unsafe.Add(unsafe.Pointer(v15), unsafe.Sizeof(uint16(0))*1))), v60)))
						if v37 == 0 || v37 == 3 || v37 == 10 {
							goto LABEL_86
						}
						v38 = v37 == 7
					case 7:
						v41 = int32(uint8(sub_57B500(int32(*v15)+1, int32(*(*uint16)(unsafe.Add(unsafe.Pointer(v15), unsafe.Sizeof(uint16(0))*1))), v60)))
						if v41 != 0 && v41 != 5 && v41 != 9 && v41 != 8 && v41 != math.MaxUint8 {
							continue
						}
						v42 = int32(uint8(sub_57B500(int32(*v15), int32(*(*uint16)(unsafe.Add(unsafe.Pointer(v15), unsafe.Sizeof(uint16(0))*1)))-1, v60)))
						if v42 != 0 {
							if v42 != 3 && v42 != 10 && v42 != 7 && v42 != math.MaxUint8 {
								continue
							}
						}
						goto LABEL_86
					default:
						goto LABEL_86
					}
					if !v38 && v37 != math.MaxUint8 {
						continue
					}
				LABEL_86:
					v30 = a2
					if sub_50B950(a2, int32(*v15), int32(*(*uint32)(unsafe.Pointer(v71))+uint32(*(*uint16)(unsafe.Add(unsafe.Pointer(v15), unsafe.Sizeof(uint16(0))*1))))) != 0 || sub_50B950(a2, int32(*(*uint32)(unsafe.Pointer(v72))+uint32(*v15)), int32(*(*uint16)(unsafe.Add(unsafe.Pointer(v15), unsafe.Sizeof(uint16(0))*1)))) != 0 {
						continue
					}
				}
				if sub_50B870(v30, v23, v24) != 0 {
					continue
				}
				if v73 != v23 || v74 != v24 {
					if sub_50C830(v30, v23, v24) == 0 {
						continue
					}
				} else if sub_50C830(v30, v23, v24) == 0 {
					continue
				}
				v43 = nox_xxx_pathfind_preCheckWalls2_50B8A0(v30, v23, v24)
				if v63 != 0 {
					if v43 != 0 {
						v64 = 1
					} else if sub_50B950(v30, int32(*v15), int32(*(*uint16)(unsafe.Add(unsafe.Pointer(v15), unsafe.Sizeof(uint16(0))*1)))) == 0 && sub_50B950(v30, v23, v24) != 0 {
						continue
					}
				} else {
					if v43 == 0 {
						continue
					}
					if a5 != nil {
						v44 = a5(v30, v23, v24)
						if v61 != 0 {
							if v44 != 0 {
								v65 = 1
							}
						} else if v44 == 0 {
							continue
						}
					}
				}
			LABEL_32:
				v31 = (*uint16)(nox_alloc_class_new_obj_zero((*nox_alloc_class)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&nox_alloc_visitNode_2386184))))))
				if v31 == nil {
					v55 = int32(*(*uint32)(unsafe.Pointer(uintptr(v30 + 36))))
					v47 = nox_xxx_getUnitName_4E39D0(v30)
					nox_ai_debug_printf_5341A0(CString("%d: %s(#%d), buildPath: Exhausted search storage\n"), nox_frame_xxx_2598000, v47, v55)
				LABEL_130:
					dword_5d4594_1599712 = 1
					return unsafe.Pointer(uintptr(sub_50C320(v30, v67, nil)))
				}
				*v31 = uint16(int16(v23))
				*(*uint16)(unsafe.Add(unsafe.Pointer(v31), unsafe.Sizeof(uint16(0))*1)) = uint16(int16(v24))
				*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v31))), unsafe.Sizeof(uint32(0))*1))) = uint32(uintptr(unsafe.Pointer(v15)))
				*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v31))), unsafe.Sizeof(uint32(0))*2))) = uint32(uintptr(unsafe.Pointer(v77)))
				v77 = v31
			}
			if sub_50AC20(int32(uintptr(unsafe.Pointer(v15))), (*uint16)(unsafe.Pointer(&v68[0]))) == 0 {
				goto LABEL_38
			}
			*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(v15))), 12))) |= 2
			v32 = (*uint16)(nox_alloc_class_new_obj_zero((*nox_alloc_class)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&nox_alloc_visitNode_2386184))))))
			if v32 == nil {
				v30 = a2
				v56 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 36))))
				v48 = nox_xxx_getUnitName_4E39D0(a2)
				nox_ai_debug_printf_5341A0(CString("%d: %s(#%d), buildPath: Exhausted search storage\n"), nox_frame_xxx_2598000, v48, v56)
				goto LABEL_130
			}
			*(*uint32)(unsafe.Pointer(v32)) = *(*uint32)(unsafe.Pointer(&v68[0]))
			*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v32))), unsafe.Sizeof(uint32(0))*1))) = uint32(uintptr(unsafe.Pointer(v15)))
			*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v32))), unsafe.Sizeof(uint32(0))*2))) = uint32(uintptr(unsafe.Pointer(v77)))
			v77 = v32
		LABEL_38:
			v15 = (*uint16)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v15))), unsafe.Sizeof(uint32(0))*2))))))
			if v15 != nil {
				continue
			}
			break
		}
		v14 = v77
	LABEL_40:
		v33 = nil
		if v64 != 0 {
			v34 = v14
			if v14 != nil {
				for {
					v35 = (*uint16)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v34))), unsafe.Sizeof(uint32(0))*2))))))
					if nox_xxx_pathfind_preCheckWalls2_50B8A0(a2, int32(*v34), int32(*(*uint16)(unsafe.Add(unsafe.Pointer(v34), unsafe.Sizeof(uint16(0))*1)))) != 0 {
						v33 = v34
					} else if v33 != nil {
						*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v33))), unsafe.Sizeof(uint32(0))*2))) = *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v34))), unsafe.Sizeof(uint32(0))*2)))
					} else {
						v14 = (*uint16)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v34))), unsafe.Sizeof(uint32(0))*2))))))
					}
					v34 = v35
					if v35 == nil {
						break
					}
				}
				v33 = nil
			}
			v63 = 0
			v64 = 0
		}
		if v65 != 0 {
			v45 = v14
			if v14 != nil {
				for {
					v46 = (*uint16)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v45))), unsafe.Sizeof(uint32(0))*2))))))
					if a5(a2, int32(*v45), int32(*(*uint16)(unsafe.Add(unsafe.Pointer(v45), unsafe.Sizeof(uint16(0))*1)))) != 0 {
						v33 = v45
					} else if v33 != nil {
						*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v33))), unsafe.Sizeof(uint32(0))*2))) = *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v45))), unsafe.Sizeof(uint32(0))*2)))
					} else {
						v14 = (*uint16)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v45))), unsafe.Sizeof(uint32(0))*2))))))
					}
					v45 = v46
					if v46 == nil {
						break
					}
				}
			}
			v61 = 0
			v65 = 0
		}
		if v14 != nil {
			if a6 != 0 && v66 >= a6 {
				v58 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 36))))
				v50 = nox_xxx_getUnitName_4E39D0(a2)
				nox_ai_debug_printf_5341A0(CString("%d: %s(#%d), buildPath: Reached search depth limit\n"), nox_frame_xxx_2598000, v50, v58)
				goto LABEL_134
			}
			v66++
			continue
		}
		break
	}
	v57 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 36))))
	v49 = nox_xxx_getUnitName_4E39D0(a2)
	nox_ai_debug_printf_5341A0(CString("%d: %s(#%d), buildPath: Exhausted search space\n"), nox_frame_xxx_2598000, v49, v57)
LABEL_134:
	v30 = a2
	dword_5d4594_1599712 = 2
	return unsafe.Pointer(uintptr(sub_50C320(v30, v67, nil)))
}
func sub_50C320(a1 int32, a2 int32, a3 *uint32) int32 {
	var (
		result int32
		v4     *uint16
		v5     int32
		v6     *uint32
		v7     *uint16
		v8     int32
		v9     int32
		v10    int8
		v11    int32
		v12    int32
		v13    uint16
		v14    int32
		v15    uint16
		v16    int32
		v17    uint16
		v18    float64
		v19    float64
		v20    float64
		v21    int32
		v22    *uint32
		v23    *uint32
		v24    int32
		v25    int32
		v26    float2
		v27    float4
	)
	result = a1
	v4 = (*uint16)(unsafe.Pointer(uintptr(a2)))
	v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
	v25 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
	if a2 == 0 {
		dword_5d4594_2386180 = 0
		return result
	}
	v6 = a3
	if a3 != nil {
		**(**uint32)(unsafe.Pointer(&dword_5d4594_2386176)) = *a3
		*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_2386176 + 4))) = *(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*1))
	} else if int32(*(*uint8)(unsafe.Pointer(uintptr(a2 + 12))))&2 != 0 && sub_50AEA0(a2, &v26, (*uint32)(unsafe.Pointer(&a1))) != 0 {
		**(**float32)(unsafe.Pointer(&dword_5d4594_2386176)) = v26.field_0
		*(*float32)(unsafe.Pointer(uintptr(dword_5d4594_2386176 + 4))) = v26.field_4
		if int32(*(*uint8)(unsafe.Pointer(uintptr(v5 + 2172)))) < 8 {
			*(*uint32)(unsafe.Pointer(uintptr(v5 + int32(func() uint8 {
				p := (*uint8)(unsafe.Pointer(uintptr(v5 + 2172)))
				x := *p
				*p++
				return x
			}())*4 + 2140))) = uint32(a1)
		}
	} else {
		a2 = int32(*v4)
		**(**float32)(unsafe.Pointer(&dword_5d4594_2386176)) = float32(float64(a2)*23.0 + 11.5)
		a2 = int32(*(*uint16)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint16(0))*1)))
		*(*float32)(unsafe.Pointer(uintptr(dword_5d4594_2386176 + 4))) = float32(float64(a2)*23.0 + 11.5)
	}
	v7 = (*uint16)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v4))), unsafe.Sizeof(uint32(0))*1))))))
	v8 = 1
	if v7 != nil {
		for {
			v9 = int32(*v7) - int32(*v4)
			v10 = int8(*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(v7))), 12))))
			v11 = int32(*(*uint16)(unsafe.Add(unsafe.Pointer(v7), unsafe.Sizeof(uint16(0))*1))) - int32(*(*uint16)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint16(0))*1)))
			a2 = int32(*v7) - int32(*v4)
			if (int32(v10) & 2) == 0 {
				goto LABEL_19
			}
			if sub_50AEA0(int32(uintptr(unsafe.Pointer(v7))), &v26, (*uint32)(unsafe.Pointer(&a1))) == 0 {
				break
			}
			v12 = v25
			*(*float32)(unsafe.Pointer(uintptr(dword_5d4594_2386176 + uint32(v8*8)))) = v26.field_0
			*(*float32)(unsafe.Pointer(uintptr(dword_5d4594_2386176 + uint32(v8*8) + 4))) = v26.field_4
			if int32(*(*uint8)(unsafe.Pointer(uintptr(v25 + 2172)))) < 8 {
				a2 = a2*a2 + v11*v11
				if float64(a2) > 4761.0 {
					goto LABEL_17
				}
				v13 = *(*uint16)(unsafe.Add(unsafe.Pointer(v7), unsafe.Sizeof(uint16(0))*1))
				a2 = int32(*v7)*23 + 11
				v14 = int32(v13) * 23
				v15 = *v4
				v27.field_0 = float32(float64(a2))
				a2 = v14 + 11
				v16 = int32(v15) * 23
				v17 = *(*uint16)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint16(0))*1))
				v27.field_4 = float32(float64(v14 + 11))
				v27.field_8 = float32(float64(v16 + 11))
				a2 = int32(v17)*23 + 11
				v27.field_C = float32(float64(a2))
				if nox_xxx_mapTraceRay_535250(&v27, nil, nil, 1) == 0 {
					v12 = v25
				LABEL_17:
					*(*uint32)(unsafe.Pointer(uintptr(v12 + int32(func() uint8 {
						p := (*uint8)(unsafe.Pointer(uintptr(v12 + 2172)))
						x := *p
						*p++
						return x
					}())*4 + 2140))) = uint32(a1)
					goto LABEL_38
				}
			}
		LABEL_38:
			v4 = v7
			v7 = (*uint16)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v7))), unsafe.Sizeof(uint32(0))*1))))))
			v8++
			if v7 == nil {
				goto LABEL_39
			}
		}
		v9 = a2
	LABEL_19:
		if v9 <= 0 {
			if v9 < 0 {
				if v11 > 0 {
					a2 = int32(*v7)
					*(*float32)(unsafe.Pointer(uintptr(dword_5d4594_2386176 + uint32(v8*8)))) = float32(float64(a2)*23.0 + 23.0 - 2.3)
					a2 = int32(*(*uint16)(unsafe.Add(unsafe.Pointer(v7), unsafe.Sizeof(uint16(0))*1)))
					v20 = float64(a2)
				LABEL_29:
					*(*float32)(unsafe.Pointer(uintptr(dword_5d4594_2386176 + uint32(v8*8) + 4))) = float32(v20*23.0 + 2.3)
					goto LABEL_38
				}
				if v11 < 0 {
					a2 = int32(*v7)
					*(*float32)(unsafe.Pointer(uintptr(dword_5d4594_2386176 + uint32(v8*8)))) = float32(float64(a2)*23.0 + 23.0 - 2.3)
					a2 = int32(*(*uint16)(unsafe.Add(unsafe.Pointer(v7), unsafe.Sizeof(uint16(0))*1)))
					v18 = float64(a2)
					goto LABEL_37
				}
				a2 = int32(*v7)
				*(*float32)(unsafe.Pointer(uintptr(dword_5d4594_2386176 + uint32(v8*8)))) = float32(float64(a2)*23.0 + 23.0 - 2.3)
				a2 = int32(*(*uint16)(unsafe.Add(unsafe.Pointer(v7), unsafe.Sizeof(uint16(0))*1)))
				v19 = float64(a2)
				goto LABEL_25
			}
			if v11 > 0 {
				a2 = int32(*v7)
				*(*float32)(unsafe.Pointer(uintptr(dword_5d4594_2386176 + uint32(v8*8)))) = float32(float64(a2)*23.0 + 11.5)
				a2 = int32(*(*uint16)(unsafe.Add(unsafe.Pointer(v7), unsafe.Sizeof(uint16(0))*1)))
				v20 = float64(a2)
				goto LABEL_29
			}
			if v11 < 0 {
				a2 = int32(*v7)
				*(*float32)(unsafe.Pointer(uintptr(dword_5d4594_2386176 + uint32(v8*8)))) = float32(float64(a2)*23.0 + 11.5)
				a2 = int32(*(*uint16)(unsafe.Add(unsafe.Pointer(v7), unsafe.Sizeof(uint16(0))*1)))
				v18 = float64(a2)
				goto LABEL_37
			}
		} else {
			if v11 <= 0 {
				if v11 < 0 {
					a2 = int32(*v7)
					*(*float32)(unsafe.Pointer(uintptr(dword_5d4594_2386176 + uint32(v8*8)))) = float32(float64(a2)*23.0 + 2.3)
					a2 = int32(*(*uint16)(unsafe.Add(unsafe.Pointer(v7), unsafe.Sizeof(uint16(0))*1)))
					v18 = float64(a2)
				LABEL_37:
					*(*float32)(unsafe.Pointer(uintptr(dword_5d4594_2386176 + uint32(v8*8) + 4))) = float32(v18*23.0 + 23.0 - 2.3)
					goto LABEL_38
				}
				a2 = int32(*v7)
				*(*float32)(unsafe.Pointer(uintptr(dword_5d4594_2386176 + uint32(v8*8)))) = float32(float64(a2)*23.0 + 2.3)
				a2 = int32(*(*uint16)(unsafe.Add(unsafe.Pointer(v7), unsafe.Sizeof(uint16(0))*1)))
				v19 = float64(a2)
			LABEL_25:
				*(*float32)(unsafe.Pointer(uintptr(dword_5d4594_2386176 + uint32(v8*8) + 4))) = float32(v19*23.0 + 11.5)
				goto LABEL_38
			}
			a2 = int32(*v7)
			*(*float32)(unsafe.Pointer(uintptr(dword_5d4594_2386176 + uint32(v8*8)))) = float32(float64(a2)*23.0 + 2.3)
			a2 = int32(*(*uint16)(unsafe.Add(unsafe.Pointer(v7), unsafe.Sizeof(uint16(0))*1)))
			*(*float32)(unsafe.Pointer(uintptr(dword_5d4594_2386176 + uint32(v8*8) + 4))) = float32(float64(a2)*23.0 + 2.3)
		}
		goto LABEL_38
	}
LABEL_39:
	v21 = 0
	dword_5d4594_2386180 = uint32(v8)
	for result = v8 / 2; v21 < *(*int32)(unsafe.Pointer(&dword_5d4594_2386180))/2; result = int32(dword_5d4594_2386180 / 2) {
		v22 = (*uint32)(unsafe.Pointer(uintptr(dword_5d4594_2386176 + uint32(v21*8))))
		v23 = (*uint32)(unsafe.Pointer(uintptr(dword_5d4594_2386176 + uint32((v8-v21)*8) - 8)))
		v26 = *(*float2)(unsafe.Pointer(uintptr(dword_5d4594_2386176 + uint32(v21*8))))
		*v22 = *v23
		*(*uint32)(unsafe.Add(unsafe.Pointer(v22), unsafe.Sizeof(uint32(0))*1)) = *(*uint32)(unsafe.Add(unsafe.Pointer(v23), unsafe.Sizeof(uint32(0))*1))
		v24 = int32(dword_5d4594_2386180 - uint32(func() int32 {
			p := &v21
			x := *p
			*p++
			return x
		}()))
		*(*float32)(unsafe.Pointer(uintptr(dword_5d4594_2386176 + uint32(v24*8) - 8))) = v26.field_0
		*(*float32)(unsafe.Pointer(uintptr(dword_5d4594_2386176 + uint32(v24*8) - 4))) = v26.field_4
		v8 = int32(dword_5d4594_2386180)
	}
	return result
}
func sub_50C830(a1 int32, a2 int32, a3 int32) int32 {
	var (
		v3 int32
		v4 *float32
		v6 float2
		v7 float32
		v8 float32
	)
	v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 16))))
	if v3&0x4000 != 0 || sub_534020(a1) != 0 {
		return 1
	}
	v4 = mem_getFloatPtr(0x587000, 0x392CC)
	v8 = float32(float64(a2 * 23))
	for {
		v6.field_0 = v8 + *((*float32)(unsafe.Add(unsafe.Pointer(v4), -int(unsafe.Sizeof(float32(0))*1))))
		v7 = float32(float64(a3 * 23))
		v6.field_4 = v7 + *v4
		if nox_xxx_tileNFromPoint_411160(&v6) == 6 {
			break
		}
		v4 = (*float32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(float32(0))*2))
		if int32(uintptr(unsafe.Pointer(v4))) >= int32(uintptr(memmap.PtrOff(0x587000, 234220))) {
			return 1
		}
	}
	return 0
}
func nox_xxx_pathfind_preCheckWalls_50C8D0(a1 int32, a2 *int2) {
	var (
		v2 int32
		v3 int32
		v4 float64
		v5 float32
		v6 float32
	)
	if sub_50B870(a1, a2.field_0, a2.field_4) != 0 {
		v2 = a2.field_0
		v3 = a2.field_4
		v5 = float32(float64(v3)*23.0 + 11.5)
		v4 = float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 56)))) - (float64(a2.field_0)*23.0 + 11.5)
		v6 = *(*float32)(unsafe.Pointer(uintptr(a1 + 60))) - v5
		if math.Abs(float64(v6)) >= math.Abs(v4) {
			if float64(v6) <= 0.0 {
				a2.field_4 = v3 - 1
			} else {
				a2.field_4 = v3 + 1
			}
		} else if v4 <= 0.0 {
			a2.field_0 = v2 - 1
		} else {
			a2.field_0 = v2 + 1
		}
	}
}
func nox_xxx_appendWorkPath_50C990(a1 int32, a2 int32, a3 int32) int32 {
	var (
		v3 int32
		v4 int32
		v5 int32
		v6 *uint32
	)
	v3 = 0
	if *(*int32)(unsafe.Pointer(&dword_5d4594_2386180)) <= 0 {
		return a2
	}
	v4 = a2
	v5 = a1 + a2*8
	for v4 != a3-1 {
		v4++
		v5 += 8
		v6 = (*uint32)(unsafe.Pointer(uintptr(dword_5d4594_2386176 + uint32(func() int32 {
			p := &v3
			x := *p
			*p++
			return x
		}()*8))))
		*(*uint32)(unsafe.Pointer(uintptr(v5 - 8))) = *v6
		*(*uint32)(unsafe.Pointer(uintptr(v5 - 4))) = *(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*1))
		if v3 >= *(*int32)(unsafe.Pointer(&dword_5d4594_2386180)) {
			return v4
		}
	}
	nox_ai_debug_printf_5341A0(CString("appendWorkPath: Path truncated.\n"))
	return v4
}
func nox_xxx_generateRetreatPath_50CA00(a1 int32, a2 int32, a3 int32, a4 *float32) int32 {
	*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a3 + 748))) + 2172))) = 0
	nox_xxx_pathFind_50BA00(1, a3, (*float32)(unsafe.Pointer(uintptr(a3+56))), a4, sub_50CA60, 6)
	if dword_5d4594_2386180 > uint32(a2) && dword_5d4594_1599712 == 0 {
		dword_5d4594_1599712 = 1
	}
	return nox_xxx_appendWorkPath_50C990(a1, 0, a2)
}
func sub_50CA60(a4 int32, a2 int32, a3 int32) int32 {
	var a1 float2
	dword_5d4594_1599708 = 0
	a1.field_0 = float32(float64(a2)*23.0 + 11.5)
	a1.field_4 = float32(float64(a3)*23.0 + 11.5)
	nox_xxx_unitsGetInCircle_517F90(&a1, 100.0, unsafe.Pointer(cgoFuncAddr(sub_50CAC0)), (*nox_object_t)(unsafe.Pointer(uintptr(a4))))
	return bool2int(dword_5d4594_1599708 == 0)
}
func sub_50CAC0(a1 int32, a2 int32) {
	if dword_5d4594_1599708 != 1 {
		if nox_xxx_unitIsEnemyTo_5330C0((*nox_object_t)(unsafe.Pointer(uintptr(a2))), (*nox_object_t)(unsafe.Pointer(uintptr(a1)))) != 0 {
			dword_5d4594_1599708 = 1
		}
	}
}
func nox_xxx_pathFindStatus_50CAF0() int32 {
	return int32(dword_5d4594_1599712)
}
func sub_50CB00() int32 {
	return int32(dword_5d4594_2386180)
}
func sub_50CB10() unsafe.Pointer {
	return *(*unsafe.Pointer)(unsafe.Pointer(&dword_5d4594_2386176))
}
func sub_50CB20(a1 int32, a2 *float32) int32 {
	var (
		v2  int32
		v3  float64
		v4  *uint16
		v5  int32
		v6  *uint16
		v7  int32
		v8  int32
		v9  *uint8
		v10 int32
		v11 int32
		v12 *uint16
		v14 float32
		v15 float32
		v16 int32
		a2a int2
		v18 float2
		v19 *uint16
	)
	dword_5d4594_2386160++
	v14 = float32(float64(*a2) * 0.043478262)
	v2 = int(v14)
	v3 = float64(*(*float32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(float32(0))*1))) * 0.043478262
	a2a.field_0 = v2
	v15 = float32(v3)
	a2a.field_4 = int(v15)
	nox_xxx_pathfind_preCheckWalls_50C8D0(a1, &a2a)
	nox_alloc_class_free_all((*nox_alloc_class)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&nox_alloc_visitNode_2386184)))))
	v4 = (*uint16)(nox_alloc_class_new_obj_zero((*nox_alloc_class)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&nox_alloc_visitNode_2386184))))))
	*v4 = uint16(int16(a2a.field_0))
	*(*uint16)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint16(0))*1)) = uint16(int16(a2a.field_4))
	v5 = int32(*(*uint16)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint16(0))*1))) + (int32(*v4) << 8)
	*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v4))), unsafe.Sizeof(uint32(0))*1))) = 0
	*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v4))), unsafe.Sizeof(uint32(0))*2))) = 0
	nox_server_xxx_1599716[v5].field_0 = dword_5d4594_2386160
	for {
		v6 = v4
		v19 = nil
		if v4 == nil {
			break
		}
		for {
			v7 = int32(*v6)
			v8 = int32(*(*uint16)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint16(0))*1)))
			if int32(nox_server_xxx_1599716[int32(uint16(int16(v8)))+(v7<<8)].field_8)&64 != 0 && sub_50B870(a1, v7, v8) == 0 {
				v18.field_0 = float32(float64(v7*23 + 11))
				v18.field_4 = float32(float64(v8*23 + 11))
				return sub_518740(&v18, 128)
			}
			v9 = (*uint8)(memmap.PtrOff(0x587000, 234284))
			for {
				v10 = int32(uint32(*v6) + *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v9))), -int(unsafe.Sizeof(uint32(0))*1)))))
				v11 = int32(*(*uint32)(unsafe.Pointer(v9)) + uint32(*(*uint16)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint16(0))*1))))
				if v10 >= 0 && v10 < 256 && v11 >= 0 && v11 < 256 && nox_server_xxx_1599716[v11+(v10<<8)].field_0 != dword_5d4594_2386160 {
					v16 = int32(*(*uint32)(unsafe.Pointer(v9)) + uint32(*(*uint16)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint16(0))*1))))
					nox_server_xxx_1599716[v11+(v10<<8)].field_0 = dword_5d4594_2386160
					if sub_50B870(a1, v10, v16) == 0 {
						if sub_50C830(a1, v10, v11) != 0 {
							v12 = (*uint16)(nox_alloc_class_new_obj_zero((*nox_alloc_class)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&nox_alloc_visitNode_2386184))))))
							if v12 != nil {
								*v12 = uint16(int16(v10))
								*(*uint16)(unsafe.Add(unsafe.Pointer(v12), unsafe.Sizeof(uint16(0))*1)) = uint16(int16(v11))
								*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v12))), unsafe.Sizeof(uint32(0))*1))) = uint32(uintptr(unsafe.Pointer(v6)))
								*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v12))), unsafe.Sizeof(uint32(0))*2))) = uint32(uintptr(unsafe.Pointer(v19)))
								v19 = v12
							}
						}
					}
				}
				v9 = (*uint8)(unsafe.Add(unsafe.Pointer(v9), 8))
				if int32(uintptr(unsafe.Pointer(v9))) >= int32(uintptr(memmap.PtrOff(0x587000, 234316))) {
					break
				}
			}
			v6 = (*uint16)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v6))), unsafe.Sizeof(uint32(0))*2))))))
			if v6 == nil {
				break
			}
		}
		v4 = v19
		if v19 == nil {
			break
		}
	}
	return 0
}
func sub_50CD30() *uint8 {
	return (*uint8)(memmap.PtrOff(6112660, 2386196))
}
func nox_xxx_audioAddAIInteresting_50CD40(a1 int32, a2 int32, a3 *uint32) {
	if sub_501900(a1) == 0 {
		return
	}
	if nox_alloc_monsterListen_2386188 == nil {
		nox_alloc_monsterListen_2386188 = unsafe.Pointer(nox_new_alloc_class(CString("MonsterListen"), 24, 128))
	}
	if nox_alloc_monsterListen_2386188 == nil {
		return
	}
	if sub_534160(a2) == nil {
		return
	}
	var p *uint32 = (*uint32)(nox_alloc_class_new_obj((*nox_alloc_class)(nox_alloc_monsterListen_2386188)))
	if p == nil {
		return
	}
	*((*uint32)(unsafe.Add(unsafe.Pointer(p), unsafe.Sizeof(uint32(0))*0))) = uint32(a1)
	*((*uint32)(unsafe.Add(unsafe.Pointer(p), unsafe.Sizeof(uint32(0))*1))) = uint32(a2)
	*((*uint32)(unsafe.Add(unsafe.Pointer(p), unsafe.Sizeof(uint32(0))*2))) = *a3
	*((*uint32)(unsafe.Add(unsafe.Pointer(p), unsafe.Sizeof(uint32(0))*3))) = *(*uint32)(unsafe.Add(unsafe.Pointer(a3), unsafe.Sizeof(uint32(0))*1))
	*((*uint32)(unsafe.Add(unsafe.Pointer(p), unsafe.Sizeof(uint32(0))*4))) = nox_frame_xxx_2598000
	*((*uint32)(unsafe.Add(unsafe.Pointer(p), unsafe.Sizeof(uint32(0))*5))) = uint32(uintptr(nox_monsterListen_2386192))
	nox_monsterListen_2386192 = unsafe.Pointer(p)
}
func nox_xxx_unitListenRoutine_50CDD0(unitA int32) {
	var (
		unit                  int32
		previousMonsterListen int32
		monsterListen         int32
		monsterListenFlags    uint32
		nextMonsterListen     int32
		monsterListenData     int32
		raycastResults        int32
		heardMonster          *int32
		flags                 int32
		raycastDistance       int32
	)
	unit = unitA
	previousMonsterListen = 0
	flags = int32(*(*uint32)(unsafe.Pointer(uintptr(unitA + 748))))
	heardMonster = nil
	raycastDistance = 0
	if nox_xxx_checkIsKillable_528190(unit) == 0 {
		return
	}
	if nox_xxx_unitIsAFish_534B10((*nox_object_t)(unsafe.Pointer(uintptr(unit)))) != 0 {
		return
	}
	if nox_xxx_unitIsAFrog_534B90((*nox_object_t)(unsafe.Pointer(uintptr(unit)))) != 0 {
		return
	}
	if nox_xxx_unitIsARat_534B60((*nox_object_t)(unsafe.Pointer(uintptr(unit)))) != 0 {
		return
	}
	monsterListen = int32(uintptr(nox_monsterListen_2386192))
	if nox_monsterListen_2386192 == nil {
		return
	}
	for {
		monsterListenFlags = *(*uint32)(unsafe.Pointer(uintptr(monsterListen + 16)))
		nextMonsterListen = int32(*(*uint32)(unsafe.Pointer(uintptr(monsterListen + 20))))
		if nox_frame_xxx_2598000 < monsterListenFlags || nox_frame_xxx_2598000-monsterListenFlags > 2 {
			if previousMonsterListen != 0 {
				*(*uint32)(unsafe.Pointer(uintptr(previousMonsterListen + 20))) = uint32(nextMonsterListen)
			} else {
				nox_monsterListen_2386192 = unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(monsterListen + 20)))))
			}
			nox_alloc_class_free_obj_first((*nox_alloc_class)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&nox_alloc_monsterListen_2386188)))), unsafe.Pointer(uintptr(monsterListen)))
		} else {
			monsterListenData = int32(*(*uint32)(unsafe.Pointer(uintptr(monsterListen + 4))))
			if monsterListenData != 0 && int32(*(*uint8)(unsafe.Pointer(uintptr(monsterListenData + 16))))&32 != 0 {
				*(*uint32)(unsafe.Pointer(uintptr(monsterListen + 4))) = 0
			}
			if *(*uint32)(unsafe.Pointer(uintptr(flags + 404))) <= uint32(*(*int32)(unsafe.Pointer(uintptr(monsterListen + 16)))) {
				if sub_nox_xxx_checkIfUnitShouldListenOther_50CF10(unit, monsterListen) != 0 {
					raycastResults = sub_50D000(unit, monsterListen)
					if raycastResults > 0 && raycastResults > raycastDistance {
						raycastDistance = raycastResults
						heardMonster = (*int32)(unsafe.Pointer(uintptr(monsterListen)))
					}
				}
			}
			previousMonsterListen = monsterListen
		}
		monsterListen = nextMonsterListen
		if nextMonsterListen == 0 {
			break
		}
	}
	if heardMonster != nil && (uint32(*(*int32)(unsafe.Add(unsafe.Pointer(heardMonster), unsafe.Sizeof(int32(0))*4))) > uint32(*(*int32)(unsafe.Pointer(uintptr(flags + 404)))) || raycastDistance > *(*int32)(unsafe.Pointer(uintptr(flags + 408)))) {
		nox_xxx_unitEmitHearEvent_50D110(unit, (*uint32)(unsafe.Pointer(heardMonster)), raycastDistance)
	}
}
func sub_nox_xxx_checkIfUnitShouldListenOther_50CF10(unitA int32, unitB int32) int32 {
	var (
		unit                 int32
		v3                   int32
		playerUnit           int32
		listenable           int32
		xpos                 int32
		ypos                 float32
		polygonQueryResponse *nox_player_polygon_check_data
		pos                  int2
		v12                  int8
	)
	unit = unitA
	v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(unitA + 748))))
	playerUnit = int32(uintptr(unsafe.Pointer(nox_xxx_findParentChainPlayer_4EC580((*nox_object_t)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(unitB + 4))))))))))
	listenable = sub_501900(int32(*(*uint32)(unsafe.Pointer(uintptr(unitB)))))
	if *(*uint32)(unsafe.Pointer(uintptr(v3 + 404))) > nox_frame_xxx_2598000 {
		return 0
	}
	if playerUnit == 0 {
		if listenable == 0 {
			return 0
		}
	} else {
		if nox_xxx_unitIsEnemyTo_5330C0((*nox_object_t)(unsafe.Pointer(uintptr(unit))), (*nox_object_t)(unsafe.Pointer(uintptr(playerUnit)))) != 0 {
			if listenable == 0 {
				return 0
			}
		} else if (listenable & 16) == 0 {
			return 0
		}
	}
	v12 = sub_501C00((*float32)(unsafe.Pointer(uintptr(unitB+8))), int32(*(*uint32)(unsafe.Pointer(uintptr(unitB + 4)))))
	if int32(v12) != 0 {
		xpos = int(*(*float32)(unsafe.Pointer(uintptr(unit + 56))))
		ypos = *(*float32)(unsafe.Pointer(uintptr(unit + 60)))
		pos.field_0 = xpos
		pos.field_4 = int(ypos)
		polygonQueryResponse = nox_xxx_polygonIsPlayerInPolygon_4217B0(&pos, 0)
		if polygonQueryResponse != nil {
			if int32(v12) != int32(*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&polygonQueryResponse.field_0[32]))), 2))) {
				return 0
			}
		}
	}
	return 1
}
func sub_50D000(a1 int32, a2 int32) int32 {
	var (
		v2     int8
		v3     int32
		result int32
		v5     float32
		v6     float32
		v7     float32
		v8     float32
		v9     float4
	)
	v2 = int8(sub_501900(int32(*(*uint32)(unsafe.Pointer(uintptr(a2))))))
	v3 = sub_501AF0(int32(*(*uint32)(unsafe.Pointer(uintptr(a2)))), (*float32)(unsafe.Pointer(uintptr(a2+8))), (*float32)(unsafe.Pointer(uintptr(a1+56))))
	if sub_50D0C0(v2, v3) == 0 {
		return -1
	}
	v5 = *(*float32)(unsafe.Pointer(uintptr(a1 + 56)))
	v6 = *(*float32)(unsafe.Pointer(uintptr(a2 + 8)))
	v9.field_4 = *(*float32)(unsafe.Pointer(uintptr(a1 + 60)))
	v9.field_0 = v5
	v7 = *(*float32)(unsafe.Pointer(uintptr(a2 + 12)))
	v9.field_8 = v6
	v9.field_C = v7
	if nox_xxx_mapTraceRay_535250(&v9, nil, nil, 5) == 0 {
		v8 = float32(float64(v3) * 0.5)
		v3 = int(v8)
	}
	if sub_50D0C0(v2, v3) != 0 {
		result = v3
	} else {
		result = -1
	}
	return result
}
func sub_50D0C0(a1 int8, a2 int32) int32 {
	if int32(a1)&32 != 0 {
		if a2 < *memmap.PtrInt32(0x587000, 234644) {
			return 0
		}
	} else if int32(a1)&64 != 0 {
		if a2 < *memmap.PtrInt32(0x587000, 234652) {
			return 0
		}
	} else if a2 < *memmap.PtrInt32(0x587000, 234648) {
		return 0
	}
	return 1
}
func nox_xxx_unitEmitHearEvent_50D110(unit int32, heardUnit *uint32, distance int32) *uint8 {
	var (
		v3 *uint32
		v4 int32
		v5 int32
	)
	v3 = *(**uint32)(unsafe.Pointer(uintptr(unit + 748)))
	*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*97)) = *heardUnit
	*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*101)) = nox_frame_xxx_2598000
	*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*102)) = uint32(distance)
	v4 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(heardUnit), unsafe.Sizeof(uint32(0))*1)))
	if v4 != 0 {
		*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*98)) = *(*uint32)(unsafe.Pointer(uintptr(v4 + 36)))
	} else {
		*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*98)) = 0
	}
	sub_50D190(unit, (*uint32)(unsafe.Add(unsafe.Pointer(heardUnit), unsafe.Sizeof(uint32(0))*2)), (*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*99)))
	v5 = int32(uintptr(unsafe.Pointer(nox_xxx_findParentChainPlayer_4EC580((*nox_object_t)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(heardUnit), unsafe.Sizeof(uint32(0))*1)))))))))
	return nox_xxx_scriptCallByEventBlock_502490((*int32)(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*320)))), v5, unit, 16)
}
func sub_50D190(a1 int32, a2 *uint32, a3 *uint32) *uint32 {
	var result *uint32
	result = a2
	*a3 = *a2
	*(*uint32)(unsafe.Add(unsafe.Pointer(a3), unsafe.Sizeof(uint32(0))*1)) = *(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*1))
	*memmap.PtrUint32(6112660, 2386196) = *a2
	*memmap.PtrUint32(6112660, 2386200) = *(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*1))
	return result
}
func sub_50D1C0() {
	nox_alloc_class_free_all((*nox_alloc_class)(nox_alloc_monsterListen_2386188))
	nox_alloc_monsterListen_2386188 = nil
	nox_monsterListen_2386192 = nil
}
func sub_50D1E0() {
	if nox_alloc_monsterListen_2386188 != nil {
		nox_free_alloc_class((*nox_alloc_class)(nox_alloc_monsterListen_2386188))
		nox_alloc_monsterListen_2386188 = nil
	}
	nox_monsterListen_2386192 = nil
}
func sub_50D210() int32 {
	return int32(*memmap.PtrUint32(6112660, 2386204))
}
func nox_xxx_creatureSetDetailedPath_50D220(a1 int32, a2 int32) int32 {
	var (
		result int32
		v3     int32
		v4     int32
	)
	result = 0
	v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
	v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(v3 + 280))))
	*(*uint32)(unsafe.Pointer(uintptr(v3 + 268))) = 0
	if (nox_frame_xxx_2598000 - uint32(v4)) >= 10 {
		*(*uint32)(unsafe.Pointer(uintptr(v3 + 272))) = *(*uint32)(unsafe.Pointer(uintptr(a2)))
		*(*uint32)(unsafe.Pointer(uintptr(v3 + 276))) = *(*uint32)(unsafe.Pointer(uintptr(a2 + 4)))
		*(*uint32)(unsafe.Pointer(uintptr(v3 + 8))) = uint32(nox_xxx_genPathToPoint_50B9A0(v3+12, 32, a1, (*float32)(unsafe.Pointer(uintptr(a2)))))
		result = nox_xxx_pathFindStatus_50CAF0()
		*(*uint8)(unsafe.Pointer(uintptr(v3 + 284))) = uint8(int8(result))
		*(*uint32)(unsafe.Pointer(uintptr(v3 + 280))) = nox_frame_xxx_2598000
	} else {
		*(*uint32)(unsafe.Pointer(uintptr(v3 + 8))) = 0
		*(*uint8)(unsafe.Pointer(uintptr(v3 + 284))) = 1
	}
	return result
}
func sub_50D2A0(a1 int32, a2 int32) int32 {
	var (
		v2     *uint32
		result int32
	)
	v2 = *(**uint32)(unsafe.Pointer(uintptr(a1 + 748)))
	*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*92)) = *(*uint32)(unsafe.Pointer(uintptr(a2)))
	*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*93)) = *(*uint32)(unsafe.Pointer(uintptr(a2 + 4)))
	result = sub_547F20(a1, (*float32)(unsafe.Pointer(uintptr(a2))))
	*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*74)) = uint32(result)
	*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*91)) = 0
	return result
}
func sub_50D2E0(a1 int32) int32 {
	var (
		v1 int32
		v2 int32
		v3 int32
		v4 int32
		v5 float64
		v6 float64
		v8 *uint32
		v9 int32
	)
	v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 296))))
	if v2 != 0 && *(*uint32)(unsafe.Pointer(uintptr(v1 + 8))) == 0 {
		v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 364))))
		v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + v3*4 + 300))))
		v5 = float64(*(*float32)(unsafe.Pointer(uintptr(v4 + 8))) - *(*float32)(unsafe.Pointer(uintptr(a1 + 56))))
		v6 = float64(*(*float32)(unsafe.Pointer(uintptr(v4 + 12))) - *(*float32)(unsafe.Pointer(uintptr(a1 + 60))))
		if v6*v6+v5*v5 < 64.0 {
			if v3 == v2-1 {
				*(*uint32)(unsafe.Pointer(uintptr(v1 + 296))) = 0
				return 1
			}
			*(*uint32)(unsafe.Pointer(uintptr(v1 + 364))) = uint32(v3 + 1)
			v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + (v3+1)*4 + 300))))
		}
		v8 = (*uint32)(unsafe.Pointer(uintptr(v4 + 8)))
		nox_xxx_creatureSetDetailedPath_50D220(a1, v4+8)
		if int32(*(*uint8)(unsafe.Pointer(uintptr(v1 + 284)))) == 0 {
			v9 = int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 8))))
			*(*uint32)(unsafe.Pointer(uintptr(v1 + v9*8 + 4))) = *v8
			*(*uint32)(unsafe.Pointer(uintptr(v1 + v9*8 + 8))) = *(*uint32)(unsafe.Add(unsafe.Pointer(v8), unsafe.Sizeof(uint32(0))*1))
		}
	}
	return bool2int(nox_xxx_creatureActuallyMove_50D3B0((*float32)(unsafe.Pointer(uintptr(a1)))) != 0 && int32(*(*uint8)(unsafe.Pointer(uintptr(v1 + 284)))) == 2)
}
func nox_xxx_creatureActuallyMove_50D3B0(a1 *float32) int32 {
	var (
		v1  int32
		v2  int32
		v3  float32
		v4  int32
		v5  *float32
		v6  float64
		v7  float64
		v8  float64
		v10 *float32
		v11 float64
		v12 float64
		v13 int16
		v14 int32
		v15 float64
		v16 *float32
		v17 float32
		v18 *float32
		v19 int32
		v20 float32
		v21 float32
		v22 float2
		v23 float4
		v24 float32
	)
	v1 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a1))), unsafe.Sizeof(uint32(0))*187))))
	v2 = 0
	if *(*uint32)(unsafe.Pointer(uintptr(v1 + 8))) != 0 {
		v3 = *(*float32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(float32(0))*15))
		v23.field_0 = *(*float32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(float32(0))*14))
		v23.field_4 = v3
		v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 268))))
		v16 = nil
		v18 = nil
		v19 = 0
		v17 = 1e+07
		if v4 >= *(*int32)(unsafe.Pointer(uintptr(v1 + 8))) {
		LABEL_14:
			*(*uint32)(unsafe.Pointer(uintptr(v1 + 8))) = 0
			return 0
		}
		v5 = (*float32)(unsafe.Pointer(uintptr(v1 + v4*8 + 12)))
		for {
			v23.field_8 = *v5
			v23.field_C = *(*float32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(float32(0))*1))
			if nox_xxx_mapTraceRay_535250(&v23, nil, nil, -124) != 0 {
				v6 = float64(v23.field_8 - *(*float32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(float32(0))*14)))
				v7 = float64(v23.field_C - *(*float32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(float32(0))*15)))
				v8 = v7*v7 + v6*v6
				if v8 > 64.0 {
					if v16 == nil || float64(v17) > v8 {
						v17 = float32(v8)
						v16 = v5
						v2 = v4
					}
				} else {
					if uint32(v4) == *(*uint32)(unsafe.Pointer(uintptr(v1 + 8)))-1 {
						*(*uint32)(unsafe.Pointer(uintptr(v1 + 8))) = 0
						return 1
					}
					v18 = v5
					v19 = v4
				}
			}
			v4++
			v5 = (*float32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(float32(0))*2))
			if v4 >= *(*int32)(unsafe.Pointer(uintptr(v1 + 8))) {
				break
			}
		}
		if v16 == nil {
			if v18 == nil {
				goto LABEL_14
			}
			v2 = v19
		}
		v10 = a1
		*memmap.PtrUint32(6112660, 2386204) = uint32(v2)
		*(*uint32)(unsafe.Pointer(uintptr(v1 + 268))) = uint32(v2)
		v20 = *(*float32)(unsafe.Pointer(uintptr(v1 + v2*8 + 12))) - *(*float32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(float32(0))*14))
		v11 = float64(*(*float32)(unsafe.Pointer(uintptr(v1 + v2*8 + 16))) - *(*float32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(float32(0))*15)))
		v21 = float32(v11)
		v24 = float32(math.Sqrt(v11*float64(v21)+float64(v20*v20)) + *mem_getDoublePtr(0x581450, 0x2830))
		if v2 <= 0 {
			v22.field_0 = *(*float32)(unsafe.Pointer(uintptr(v1 + 20))) - *(*float32)(unsafe.Pointer(uintptr(v1 + 12)))
			v12 = float64(*(*float32)(unsafe.Pointer(uintptr(v1 + 24))) - *(*float32)(unsafe.Pointer(uintptr(v1 + 16))))
		} else {
			v22.field_0 = *(*float32)(unsafe.Pointer(uintptr(v1 + v2*8 + 12))) - *(*float32)(unsafe.Pointer(uintptr(v1 + v2*8 + 4)))
			v12 = float64(*(*float32)(unsafe.Pointer(uintptr(v1 + v2*8 + 16))) - *(*float32)(unsafe.Pointer(uintptr(v1 + v2*8 + 8))))
		}
		v22.field_4 = float32(v12)
		v13 = int16(nox_xxx_math_509ED0(&v22))
		*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v10))), unsafe.Sizeof(uint16(0))*62))) = uint16(v13)
		*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v10))), unsafe.Sizeof(uint16(0))*63))) = uint16(v13)
		v14 = int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 1440))))
		if v14&0x4000 != 0 {
			v15 = float64(*(*float32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v1 + 484))) + 96))) * *(*float32)(unsafe.Add(unsafe.Pointer(v10), unsafe.Sizeof(float32(0))*136)))
		} else {
			v15 = float64(*(*float32)(unsafe.Add(unsafe.Pointer(v10), unsafe.Sizeof(float32(0))*136)))
		}
		*(*float32)(unsafe.Add(unsafe.Pointer(v10), unsafe.Sizeof(float32(0))*22)) = float32(v15 * float64(v20) / float64(v24))
		*(*float32)(unsafe.Add(unsafe.Pointer(v10), unsafe.Sizeof(float32(0))*23)) = float32(v15 * float64(v21) / float64(v24))
	}
	return 0
}
func nox_xxx_creatureSetMovePath_50D5A0(a1 int32) int32 {
	var (
		v1  int32
		v2  float64
		v3  *float32
		v4  float64
		v6  float32
		v7  int8
		v8  bool
		v9  int32
		v10 float64
		v11 float64
		v12 float64
		v13 float64
		v14 float4
	)
	v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
	v2 = float64(*(*float32)(unsafe.Pointer(uintptr(v1 + int32(*(*byte)(unsafe.Pointer(uintptr(v1 + 544)))*24) + 556))) - *(*float32)(unsafe.Pointer(uintptr(a1 + 56))))
	v3 = (*float32)(unsafe.Pointer(uintptr(v1 + int32(*(*byte)(unsafe.Pointer(uintptr(v1 + 544)))*24) + 556)))
	v4 = float64(*(*float32)(unsafe.Pointer(uintptr(v1 + int32(*(*byte)(unsafe.Pointer(uintptr(v1 + 544)))*24) + 560))) - *(*float32)(unsafe.Pointer(uintptr(a1 + 60))))
	if math.Sqrt(v4*v4+v2*v2)+*mem_getDoublePtr(0x581450, 0x2830) <= 8.0 {
		return 1
	}
	v6 = *(*float32)(unsafe.Pointer(uintptr(a1 + 60)))
	v14.field_0 = *(*float32)(unsafe.Pointer(uintptr(a1 + 56)))
	v14.field_4 = v6
	v14.field_8 = *v3
	v7 = int8(uint8((*(*uint32)(unsafe.Pointer(uintptr(a1 + 16))) >> 12) & 4))
	v14.field_C = *(*float32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(float32(0))*1))
	v8 = nox_xxx_mapTraceRay_535250(&v14, nil, nil, v7) == 0
	v9 = int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 8))))
	if v8 {
		if v9 == 0 {
			if *(*uint32)(unsafe.Pointer(uintptr(v1 + 296))) == 0 || (nox_frame_xxx_2598000-*(*uint32)(unsafe.Pointer(uintptr(v1 + 280)))) > 10 && (func() bool {
				v12 = float64(*(*float32)(unsafe.Pointer(uintptr(v1 + 368))) - *v3)
				v13 = float64(*(*float32)(unsafe.Pointer(uintptr(v1 + 372))) - *(*float32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(float32(0))*1)))
				return v13*v13+v12*v12 > 10000.0
			}()) {
				sub_50D2A0(a1, int32(uintptr(unsafe.Pointer(v3))))
			}
			if *(*uint32)(unsafe.Pointer(uintptr(v1 + 296))) != 0 {
			LABEL_16:
				if sub_50D2E0(a1) != 0 {
					*(*uint32)(unsafe.Pointer(uintptr(v1 + 8))) = 0
					*(*uint32)(unsafe.Pointer(uintptr(v1 + 296))) = 0
					return 1
				}
				return 0
			}
			nox_ai_debug_printf_5341A0(CString(" ** Waypoint path failed, doing detailed path\n"))
			nox_xxx_creatureSetDetailedPath_50D220(a1, int32(uintptr(unsafe.Pointer(v3))))
		}
	} else if v9 == 0 || (nox_frame_xxx_2598000-*(*uint32)(unsafe.Pointer(uintptr(v1 + 280)))) > 10 && (func() bool {
		v10 = float64(*(*float32)(unsafe.Pointer(uintptr(v1 + 272))) - *v3)
		v11 = float64(*(*float32)(unsafe.Pointer(uintptr(v1 + 276))) - *(*float32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(float32(0))*1)))
		return v11*v11+v10*v10 > 2500.0
	}()) {
		nox_xxx_creatureSetDetailedPath_50D220(a1, int32(uintptr(unsafe.Pointer(v3))))
	}
	if *(*uint32)(unsafe.Pointer(uintptr(v1 + 296))) != 0 {
		goto LABEL_16
	}
	if *(*uint32)(unsafe.Pointer(uintptr(v1 + 8))) != 0 && nox_xxx_creatureActuallyMove_50D3B0((*float32)(unsafe.Pointer(uintptr(a1)))) != 0 {
		*(*uint32)(unsafe.Pointer(uintptr(v1 + 8))) = 0
		return 1
	}
	return 0
}
func nox_xxx_allocMonsterRelatedArrays_50D780() int32 {
	var result int32
	result = int32(uintptr(unsafe.Pointer(nox_new_alloc_class(CString("SpawnClass"), 12, 96))))
	nox_alloc_spawn_2386216 = unsafe.Pointer(uintptr(result))
	if result != 0 {
		dword_5d4594_2386212 = 0
		result = int32(uintptr(unsafe.Pointer(nox_new_alloc_class(CString("MonsterListClass"), 148, 96))))
		nox_alloc_monsterList_2386220 = unsafe.Pointer(uintptr(result))
		if result != 0 {
			dword_5d4594_2386224 = 0
			dword_5d4594_2386228 = 0
			result = 1
		}
	}
	return result
}
func sub_50D7E0() {
	nox_alloc_class_free_all((*nox_alloc_class)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&nox_alloc_spawn_2386216)))))
	dword_5d4594_2386212 = 0
	nox_alloc_class_free_all((*nox_alloc_class)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&nox_alloc_monsterList_2386220)))))
	dword_5d4594_2386224 = 0
	dword_5d4594_2386228 = 0
}
func sub_50D820() {
	nox_free_alloc_class((*nox_alloc_class)(nox_alloc_spawn_2386216))
	nox_alloc_spawn_2386216 = nil
	dword_5d4594_2386212 = 0
	nox_free_alloc_class((*nox_alloc_class)(nox_alloc_monsterList_2386220))
	nox_alloc_monsterList_2386220 = nil
	dword_5d4594_2386224 = 0
	dword_5d4594_2386228 = 0
}
func sub_50D860(a1 int32, a2 int32) int32 {
	if *(*float32)(unsafe.Pointer(uintptr(a1 + 4))) == *(*float32)(unsafe.Pointer(uintptr(a2 + 4))) {
		return 0
	}
	if float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 4)))) <= float64(*(*float32)(unsafe.Pointer(uintptr(a2 + 4)))) {
		return -1
	}
	return 1
}
func sub_50D8D0() {
	var (
		v0 *int32
		v1 *int32
		v2 int32
		v3 int32
		v4 int32
	)
	v0 = *(**int32)(unsafe.Pointer(&dword_5d4594_2386212))
	if dword_5d4594_2386212 != 0 {
		for {
			v1 = (*int32)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(v0), unsafe.Sizeof(int32(0))*1)))))
			v2 = 1
			v3 = int32(uintptr(unsafe.Pointer(nox_xxx_getFirstPlayerUnit_4DA7C0())))
			if v3 == 0 {
				goto LABEL_13
			}
			for {
				v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(v3 + 748))))
				if *(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v4 + 276))) + 4792))) == 1 && *(*uint32)(unsafe.Pointer(uintptr(v4 + 312))) == 0 && nox_xxx_calcDistance_4E6C00((*nox_object_t)(unsafe.Pointer(uintptr(v3))), (*nox_object_t)(unsafe.Pointer(uintptr(*v0)))) < 700.0 {
					v2 = 0
				}
				v3 = int32(uintptr(unsafe.Pointer(nox_xxx_getNextPlayerUnit_4DA7F0((*nox_object_t)(unsafe.Pointer(uintptr(v3)))))))
				if v3 == 0 {
					break
				}
			}
			if v2 == 1 {
			LABEL_13:
				nox_xxx_delayedDeleteObject_4E5CC0((*nox_object_t)(unsafe.Pointer(uintptr(*v0))))
			}
			v0 = v1
			if v1 == nil {
				break
			}
		}
	}
}
func sub_50D960() int32 {
	var (
		v0  int32
		v1  int32
		v2  int32
		v3  *int32
		v4  int32
		v5  *uint32
		v6  *uint32
		v7  int32
		i   int32
		v9  float64
		v10 int32
		v11 *float32
		v12 int32
		j   int32
		k   int32
		v15 int32
		v16 int32
		v17 int32
		v18 int32
		v19 int32
		v20 int32
		v21 int32
		v22 int32
		v23 int32
		v24 int32
		v25 *uint8
		v26 *int32
		v27 *int32
		v28 int8
		v29 *uint8
		v30 int32
		v31 int32
		v32 int32
		v34 float32
		v35 int32
		v36 int32
		v37 int32
		a2  float4
		v39 float4
	)
	v34 = float32(nox_xxx_gamedataGetFloat_419D40(CString("MaxOnscreenMonsterCount")))
	v35 = int(v34)
	v36 = 0
	libc.MemSet(memmap.PtrOff(6112660, 2386232), 0, 128)
	v0 = int32(uintptr(unsafe.Pointer(nox_xxx_getFirstPlayerUnit_4DA7C0())))
	if v0 != 0 {
		for {
			v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(v0 + 748))))
			*memmap.PtrUint32(6112660, uint32(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v1 + 276))) + 2064))))*4)+0x246938) = 0
			v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 276))))
			if *(*uint32)(unsafe.Pointer(uintptr(v2 + 4792))) != 0 {
				v3 = *(**int32)(unsafe.Pointer(&dword_5d4594_2386212))
				a2.field_0 = float32(float64(*(*float32)(unsafe.Pointer(uintptr(v0 + 56)))) - float64(*(*uint16)(unsafe.Pointer(uintptr(v2 + 10)))) - 100.0)
				a2.field_4 = float32(float64(*(*float32)(unsafe.Pointer(uintptr(v0 + 60)))) - float64(*(*uint16)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v1 + 276))) + 12)))) - 100.0)
				a2.field_8 = float32(float64(*(*uint16)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v1 + 276))) + 10)))) + float64(*(*float32)(unsafe.Pointer(uintptr(v0 + 56)))) + 100.0)
				a2.field_C = float32(float64(*(*uint16)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v1 + 276))) + 12)))) + float64(*(*float32)(unsafe.Pointer(uintptr(v0 + 60)))) + 100.0)
				if dword_5d4594_2386212 != 0 {
					for {
						if sub_428220((*float2)(unsafe.Pointer(uintptr(*v3+56))), &a2) != 0 {
							v39.field_0 = *(*float32)(unsafe.Pointer(uintptr(v0 + 56)))
							v39.field_4 = *(*float32)(unsafe.Pointer(uintptr(v0 + 60)))
							v4 = *v3
							v39.field_8 = *(*float32)(unsafe.Pointer(uintptr(*v3 + 56)))
							v39.field_C = *(*float32)(unsafe.Pointer(uintptr(v4 + 60)))
							if nox_xxx_mapTraceRay_535250(&v39, nil, nil, 69) != 0 {
								v5 = sub_50DE60(*v3)
								if v5 == nil {
									v6 = (*uint32)(nox_alloc_class_new_obj_zero((*nox_alloc_class)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&nox_alloc_monsterList_2386220))))))
									v5 = v6
									if v6 == nil {
										break
									}
									*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*35)) = dword_5d4594_2386224
									if dword_5d4594_2386224 != 0 {
										*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_2386224 + 144))) = uint32(uintptr(unsafe.Pointer(v6)))
									}
									dword_5d4594_2386224 = uint32(uintptr(unsafe.Pointer(v6)))
									dword_5d4594_2386228++
									*v6 = uint32(*v3)
								}
								*(*uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(uint32(0))*2)) |= uint32(1 << int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v1 + 276))) + 2064)))))
								v7 = int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v1 + 276))) + 2064))))
								*memmap.PtrUint32(6112660, uint32(v7*4)+0x246938)++
								if *memmap.PtrUint32(6112660, uint32(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v1 + 276))) + 2064))))*4)+0x246938) > uint32(v35) {
									v36 = 1
								}
								*(*float32)(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(uint32(0))*uintptr(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v1 + 276))) + 2064))))+3))))) = float32(nox_xxx_calcDistance_4E6C00((*nox_object_t)(unsafe.Pointer(uintptr(*v3))), (*nox_object_t)(unsafe.Pointer(uintptr(v0)))))
							}
						}
						v3 = (*int32)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(int32(0))*1)))))
						if v3 == nil {
							break
						}
					}
				}
			}
			v0 = int32(uintptr(unsafe.Pointer(nox_xxx_getNextPlayerUnit_4DA7F0((*nox_object_t)(unsafe.Pointer(uintptr(v0)))))))
			if v0 == 0 {
				break
			}
		}
		if v36 != 0 {
			for i = int32(dword_5d4594_2386224); i != 0; i = int32(*(*uint32)(unsafe.Pointer(uintptr(i + 140)))) {
				v9 = 0.0
				v10 = 0
				v11 = (*float32)(unsafe.Pointer(uintptr(i + 12)))
				v12 = 32
				for {
					if float64(*v11) != 0.0 {
						v9 = v9 + float64(*v11)
						v10++
					}
					v11 = (*float32)(unsafe.Add(unsafe.Pointer(v11), unsafe.Sizeof(float32(0))*1))
					v12--
					if v12 == 0 {
						break
					}
				}
				*(*float32)(unsafe.Pointer(uintptr(i + 4))) = float32(v9 / float64(v10))
			}
			for j = int32(dword_5d4594_2386224); j != 0; j = int32(*(*uint32)(unsafe.Pointer(uintptr(j + 140)))) {
				for k = int32(*(*uint32)(unsafe.Pointer(uintptr(j + 140)))); k != 0; k = int32(*(*uint32)(unsafe.Pointer(uintptr(k + 140)))) {
					if float64(*(*float32)(unsafe.Pointer(uintptr(k + 4)))) > float64(*(*float32)(unsafe.Pointer(uintptr(j + 4)))) {
						v15 = int32(*(*uint32)(unsafe.Pointer(uintptr(j + 140))))
						if v15 == k {
							v16 = int32(*(*uint32)(unsafe.Pointer(uintptr(j + 144))))
							*(*uint32)(unsafe.Pointer(uintptr(k + 144))) = uint32(v16)
							if v16 != 0 {
								*(*uint32)(unsafe.Pointer(uintptr(v16 + 140))) = uint32(k)
							}
							v17 = int32(*(*uint32)(unsafe.Pointer(uintptr(k + 140))))
							*(*uint32)(unsafe.Pointer(uintptr(j + 140))) = uint32(v17)
							if v17 != 0 {
								*(*uint32)(unsafe.Pointer(uintptr(v17 + 144))) = uint32(j)
							}
							*(*uint32)(unsafe.Pointer(uintptr(j + 144))) = uint32(k)
							*(*uint32)(unsafe.Pointer(uintptr(k + 140))) = uint32(j)
						} else {
							v18 = int32(*(*uint32)(unsafe.Pointer(uintptr(j + 144))))
							v19 = int32(*(*uint32)(unsafe.Pointer(uintptr(k + 140))))
							v20 = int32(*(*uint32)(unsafe.Pointer(uintptr(k + 144))))
							if v18 != 0 {
								*(*uint32)(unsafe.Pointer(uintptr(v18 + 140))) = uint32(k)
							}
							if v20 != 0 {
								*(*uint32)(unsafe.Pointer(uintptr(v20 + 140))) = uint32(j)
							}
							if v15 != 0 {
								*(*uint32)(unsafe.Pointer(uintptr(v15 + 144))) = uint32(k)
							}
							if v19 != 0 {
								*(*uint32)(unsafe.Pointer(uintptr(v19 + 144))) = uint32(j)
							}
							*(*uint32)(unsafe.Pointer(uintptr(j + 144))) = uint32(v20)
							*(*uint32)(unsafe.Pointer(uintptr(j + 140))) = uint32(v19)
							*(*uint32)(unsafe.Pointer(uintptr(k + 144))) = uint32(v18)
							*(*uint32)(unsafe.Pointer(uintptr(k + 140))) = uint32(v15)
						}
						if uint32(j) == dword_5d4594_2386224 {
							dword_5d4594_2386224 = uint32(k)
						}
						v21 = k
						k = j
						j = v21
					}
				}
			}
			for {
				v22 = 0
				v23 = 0
				v24 = 0
				v25 = (*uint8)(memmap.PtrOff(6112660, 2386232))
				for {
					if *(*uint32)(unsafe.Pointer(v25)) > uint32(v22) {
						v22 = int32(*(*uint32)(unsafe.Pointer(v25)))
						v23 = v24
					}
					v25 = (*uint8)(unsafe.Add(unsafe.Pointer(v25), 4))
					v24++
					if int32(uintptr(unsafe.Pointer(v25))) >= int32(uintptr(memmap.PtrOff(6112660, 2386360))) {
						break
					}
				}
				if v22 <= v35 {
					break
				}
				v26 = *(**int32)(unsafe.Pointer(&dword_5d4594_2386224))
				v37 = 1
				if dword_5d4594_2386224 != 0 {
					for {
						if v37 != 1 {
							break
						}
						v27 = (*int32)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(v26), unsafe.Sizeof(int32(0))*35)))))
						if *(*int32)(unsafe.Add(unsafe.Pointer(v26), unsafe.Sizeof(int32(0))*2))&(1<<v23) != 0 {
							nox_xxx_delayedDeleteObject_4E5CC0((*nox_object_t)(unsafe.Pointer(uintptr(*v26))))
							v28 = 0
							v29 = (*uint8)(memmap.PtrOff(6112660, 2386232))
							for {
								if (1<<int32(v28))&*(*int32)(unsafe.Add(unsafe.Pointer(v26), unsafe.Sizeof(int32(0))*2)) != 0 {
									*(*uint32)(unsafe.Pointer(v29))--
								}
								v29 = (*uint8)(unsafe.Add(unsafe.Pointer(v29), 4))
								v28++
								if int32(uintptr(unsafe.Pointer(v29))) >= int32(uintptr(memmap.PtrOff(6112660, 2386360))) {
									break
								}
							}
							v30 = *(*int32)(unsafe.Add(unsafe.Pointer(v26), unsafe.Sizeof(int32(0))*35))
							if v30 != 0 {
								*(*uint32)(unsafe.Pointer(uintptr(v30 + 144))) = uint32(*(*int32)(unsafe.Add(unsafe.Pointer(v26), unsafe.Sizeof(int32(0))*36)))
							}
							v31 = *(*int32)(unsafe.Add(unsafe.Pointer(v26), unsafe.Sizeof(int32(0))*36))
							if v31 != 0 {
								*(*uint32)(unsafe.Pointer(uintptr(v31 + 140))) = uint32(*(*int32)(unsafe.Add(unsafe.Pointer(v26), unsafe.Sizeof(int32(0))*35)))
							} else {
								dword_5d4594_2386224 = uint32(*(*int32)(unsafe.Add(unsafe.Pointer(v26), unsafe.Sizeof(int32(0))*35)))
							}
							nox_alloc_class_free_obj_first((*nox_alloc_class)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&nox_alloc_monsterList_2386220)))), unsafe.Pointer(v26))
							v32 = int32(*memmap.PtrUint32(6112660, uint32(v23*4)+0x246938))
							dword_5d4594_2386228--
							if v32 <= v35 {
								v37 = 0
							}
						}
						v26 = v27
						if v27 == nil {
							break
						}
					}
				}
			}
		}
	}
	return sub_50DE10()
}
func sub_50DE10() int32 {
	var (
		result int32
		v1     int32
	)
	result = int32(dword_5d4594_2386224)
	if dword_5d4594_2386224 != 0 {
		for {
			v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(result + 140))))
			nox_alloc_class_free_obj_first((*nox_alloc_class)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&nox_alloc_monsterList_2386220)))), unsafe.Pointer(uintptr(result)))
			result = v1
			dword_5d4594_2386228--
			if v1 == 0 {
				break
			}
		}
		dword_5d4594_2386224 = 0
	} else {
		dword_5d4594_2386224 = 0
	}
	return result
}
func sub_50DE60(a1 int32) *uint32 {
	var result *uint32
	result = *(**uint32)(unsafe.Pointer(&dword_5d4594_2386224))
	if dword_5d4594_2386224 == 0 {
		return nil
	}
	for *result != uint32(a1) {
		result = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*35)))))
		if result == nil {
			return nil
		}
	}
	return result
}
func sub_50DE80(a1 int32, a2 *float32) int32 {
	var (
		v2  int32
		v3  int32
		v4  int32
		v5  int32
		v7  float32
		a2a float4
	)
	v7 = float32(nox_xxx_gamedataGetFloat_419D40(CString("MaxOnscreenMonsterCount")))
	v2 = int(v7)
	v3 = int32(uintptr(unsafe.Pointer(nox_xxx_getFirstPlayerUnit_4DA7C0())))
	if v3 == 0 {
		return 1
	}
	for {
		v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(v3 + 748))))
		v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(v4 + 276))))
		if *(*uint32)(unsafe.Pointer(uintptr(v5 + 4792))) != 0 {
			a2a.field_0 = float32(float64(*(*float32)(unsafe.Pointer(uintptr(v3 + 56)))) - float64(*(*uint16)(unsafe.Pointer(uintptr(v5 + 10)))) - 100.0)
			a2a.field_4 = float32(float64(*(*float32)(unsafe.Pointer(uintptr(v3 + 60)))) - float64(*(*uint16)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v4 + 276))) + 12)))) - 100.0)
			a2a.field_8 = float32(float64(*(*uint16)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v4 + 276))) + 10)))) + float64(*(*float32)(unsafe.Pointer(uintptr(v3 + 56)))) + 100.0)
			a2a.field_C = float32(float64(*(*uint16)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v4 + 276))) + 12)))) + float64(*(*float32)(unsafe.Pointer(uintptr(v3 + 60)))) + 100.0)
			if sub_428220((*float2)(unsafe.Pointer(a2)), &a2a) != 0 {
				*memmap.PtrUint32(6112660, 2386208) = 0
				nox_xxx_getUnitsInRect_517C10(&a2a, sub_50DFB0, unsafe.Pointer(uintptr(v3)))
				if *memmap.PtrUint32(6112660, 2386208) >= uint32(v2) {
					break
				}
			}
		}
		v3 = int32(uintptr(unsafe.Pointer(nox_xxx_getNextPlayerUnit_4DA7F0((*nox_object_t)(unsafe.Pointer(uintptr(v3)))))))
		if v3 == 0 {
			return 1
		}
	}
	return 0
}
func sub_50DFB0(a1 *float32, a2 int32) {
	var (
		v2 float32
		v3 float32
		v4 float32
		v5 float4
	)
	if int32(uint8(*(*float32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(float32(0))*2))))&2 != 0 && (uint32(*(*float32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(float32(0))*4)))&32) == 0 && ((uint32(*(*float32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(float32(0))*4)))&0x8000) != 0x8000 || nox_xxx_unitIsZombie_534A40(int32(uintptr(unsafe.Pointer(a1)))) != 0) {
		v2 = *(*float32)(unsafe.Pointer(uintptr(a2 + 56)))
		v3 = *(*float32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(float32(0))*14))
		v5.field_4 = *(*float32)(unsafe.Pointer(uintptr(a2 + 60)))
		v5.field_0 = v2
		v4 = *(*float32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(float32(0))*15))
		v5.field_8 = v3
		v5.field_C = v4
		if nox_xxx_mapTraceRay_535250(&v5, nil, nil, 69) == 1 {
			*memmap.PtrUint32(6112660, 2386208)++
		}
	}
}
func sub_50E030(a1 int32, a2 *uint32) int32 {
	var (
		v2     int32
		v3     *uint32
		result int32
		v5     *uint32
		v6     int32
		v7     *uint32
		v8     int32
		v9     *int32
		v10    *uint32
		v11    int32
		v12    int32
		v13    *uint32
		v14    int32
	)
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
	v3 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*187)))))
	if *(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*549)) == 0 {
		result = int32(uintptr(nox_alloc_class_new_obj_zero((*nox_alloc_class)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&nox_alloc_spawn_2386216)))))))
		v5 = (*uint32)(unsafe.Pointer(uintptr(result)))
		if result == 0 {
			return result
		}
		sub_50E110(result)
		*v5 = uint32(uintptr(unsafe.Pointer(a2)))
		*(*uint8)(unsafe.Pointer(uintptr(v2 + 86)))++
		*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*549)) = uint32(uintptr(unsafe.Pointer(v5)))
		*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*548)) = uint32(a1)
		v6 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*3)))
		if v6&8192 != 0 {
			v7 = (*uint32)(unsafe.Pointer(nox_xxx_newObjectByTypeID_4E3810(CString("Glyph"))))
			if v7 != nil {
				v8 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v7), unsafe.Sizeof(uint32(0))*173)))
				v9 = (*int32)(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*511))))
				v10 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v7), unsafe.Sizeof(uint32(0))*173)))))
				v11 = 3
				for {
					v12 = *v9
					v9 = (*int32)(unsafe.Add(unsafe.Pointer(v9), unsafe.Sizeof(int32(0))*1))
					*v10 = uint32(v12)
					v10 = (*uint32)(unsafe.Add(unsafe.Pointer(v10), unsafe.Sizeof(uint32(0))*1))
					v11--
					if v11 == 0 {
						break
					}
				}
				*(*uint32)(unsafe.Pointer(uintptr(v8 + 24))) = 0
				*(*uint32)(unsafe.Pointer(uintptr(v8 + 28))) = *(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*14))
				*(*uint32)(unsafe.Pointer(uintptr(v8 + 32))) = *(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*15))
				*(*uint8)(unsafe.Pointer(uintptr(v8 + 20))) = 0
				v13 = (*uint32)(unsafe.Pointer(uintptr(v8)))
				v14 = 3
				for {
					if *v13 != 0 {
						*(*uint8)(unsafe.Pointer(uintptr(v8 + 20)))++
					}
					v13 = (*uint32)(unsafe.Add(unsafe.Pointer(v13), unsafe.Sizeof(uint32(0))*1))
					v14--
					if v14 == 0 {
						break
					}
				}
			}
			nox_xxx_inventoryPutImpl_4F3070(int32(uintptr(unsafe.Pointer(a2))), (*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v7)))))), 1)
		}
	}
	return 1
}
func sub_50E110(a1 int32) int32 {
	var result int32
	result = a1
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 8))) = 0
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 4))) = dword_5d4594_2386212
	if dword_5d4594_2386212 != 0 {
		*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_2386212 + 8))) = uint32(a1)
	}
	dword_5d4594_2386212 = uint32(a1)
	return result
}
func sub_50E140(a1 int32) {
	var (
		v1 int32
		v2 int32
	)
	if a1 != 0 {
		v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
		v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 2192))))
		if v2 != 0 {
			*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 748))) + 86)))--
			*(*uint32)(unsafe.Pointer(uintptr(v1 + 2192))) = 0
		}
		if *(*uint32)(unsafe.Pointer(uintptr(v1 + 2196))) != 0 {
			sub_50E1B0(int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 2196)))))
			nox_alloc_class_free_obj_first((*nox_alloc_class)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&nox_alloc_spawn_2386216)))), unsafe.Pointer(*(**uint64)(unsafe.Pointer(uintptr(v1 + 2196)))))
			*(*uint32)(unsafe.Pointer(uintptr(v1 + 2196))) = 0
		}
	}
}
func sub_50E1B0(a1 int32) int32 {
	var (
		result int32
		v2     int32
		v3     int32
	)
	result = a1
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 4))))
	if v2 != 0 {
		*(*uint32)(unsafe.Pointer(uintptr(v2 + 8))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 8)))
	}
	v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 8))))
	if v3 != 0 {
		result = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 4))))
		*(*uint32)(unsafe.Pointer(uintptr(v3 + 4))) = uint32(result)
	} else {
		dword_5d4594_2386212 = *(*uint32)(unsafe.Pointer(uintptr(a1 + 4)))
	}
	return result
}
func sub_50E1E0(object int32) {
	if object != 0 {
		if nox_xxx_unitIsZombie_534A40(object) == 0 {
			sub_50E140(object)
		}
	}
}
func sub_50E210(a1 int32) {
	var (
		v1 int32
		v2 int32
		v3 int32
		v4 int32
	)
	v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
	if *memmap.PtrUint32(6112660, 2386360) == 0 {
		*memmap.PtrUint32(6112660, 2386360) = uint32(nox_xxx_getNameId_4E3AA0(CString("Glyph")))
	}
	if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))&2 != 0 {
		v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 12))))
		if v2&8192 != 0 {
			if *(*uint32)(unsafe.Pointer(uintptr(v1 + 2196))) != 0 {
				v3 = a1.FirstItem()
				if v3 != 0 {
					for {
						v4 = nox_xxx_inventoryGetNext_4E7990(v3)
						if uint32(*(*uint16)(unsafe.Pointer(uintptr(v3 + 4)))) == *memmap.PtrUint32(6112660, 2386360) {
							nox_xxx_delayedDeleteObject_4E5CC0((*nox_object_t)(unsafe.Pointer(uintptr(v3))))
						}
						v3 = v4
						if v4 == 0 {
							break
						}
					}
				}
			}
		}
	}
	sub_50E140(a1)
}
func nox_xxx_registerShopClasses_50E2A0() int32 {
	var result int32
	result = int32(uintptr(unsafe.Pointer(nox_new_alloc_class(CString("TradeSessions"), 64, 64))))
	nox_alloc_tradeSession_2386492 = unsafe.Pointer(uintptr(result))
	if result != 0 {
		nox_alloc_tradeItems_2386496 = unsafe.Pointer(nox_new_alloc_class(CString("TradeItems"), 16, 500))
		if nox_alloc_tradeItems_2386496 != nil {
			libc.MemSet(memmap.PtrOff(6112660, 2386364), 0, 128)
			dword_5d4594_2386500 = 0
			result = 1
		} else {
			nox_xxx_deleteShopInventories_50E300()
			result = 0
		}
	}
	return result
}
func nox_xxx_deleteShopInventories_50E300() int32 {
	var result int32
	if nox_alloc_tradeSession_2386492 != nil {
		nox_free_alloc_class((*nox_alloc_class)(nox_alloc_tradeSession_2386492))
	}
	nox_alloc_tradeSession_2386492 = nil
	if nox_alloc_tradeItems_2386496 != nil {
		nox_free_alloc_class((*nox_alloc_class)(nox_alloc_tradeItems_2386496))
	}
	result = 0
	libc.MemSet(memmap.PtrOff(6112660, 2386364), 0, 128)
	nox_alloc_tradeItems_2386496 = nil
	dword_5d4594_2386500 = 0
	return result
}
func sub_50E360() int32 {
	var (
		v0     *uint32
		i      *int32
		result int32
	)
	v0 = *(**uint32)(unsafe.Pointer(&dword_5d4594_2386500))
	if dword_5d4594_2386500 != 0 {
		for {
			if *(*uint32)(unsafe.Add(unsafe.Pointer(v0), unsafe.Sizeof(uint32(0))*4)) != 0 {
				for i = (*int32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v0), unsafe.Sizeof(uint32(0))*5))))); i != nil; i = (*int32)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(i), unsafe.Sizeof(int32(0))*2))))) {
					nox_xxx_objectFreeMem_4E38A0(*i)
				}
				*(*uint32)(unsafe.Add(unsafe.Pointer(v0), unsafe.Sizeof(uint32(0))*5)) = 0
			}
			v0 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v0), unsafe.Sizeof(uint32(0))*14)))))
			if v0 == nil {
				break
			}
		}
	}
	nox_alloc_class_free_all((*nox_alloc_class)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&nox_alloc_tradeSession_2386492)))))
	nox_alloc_class_free_all((*nox_alloc_class)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&nox_alloc_tradeItems_2386496)))))
	result = 0
	libc.MemSet(memmap.PtrOff(6112660, 2386364), 0, 128)
	dword_5d4594_2386500 = 0
	return result
}
func nox_xxx_shopGetItemCost_50E3D0(a1 int32, a2 int32, a3 float32) int32 {
	var (
		v3  float32
		v4  int32
		v5  int32
		v6  int32
		v7  int32
		v8  int32
		v9  float64
		v10 int32
		v11 int32
		v12 int32
		v13 float64
		v14 int32
		v15 *uint8
		v16 float64
		v17 int32
		v18 float64
		v19 int32
		v20 int32
		v21 *uint16
		v22 float64
		v23 int32
		v24 float64
		v26 float32
		v27 int32
		v28 float32
		v29 int32
		v30 float32
		v31 float32
	)
	v3 = a3
	v4 = 0
	v27 = 0
	v31 = float32(float64(*(*uint32)(unsafe.Pointer(uintptr((*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&a3))), unsafe.Sizeof(uint32(0))*0))) + 32)))))
	if *memmap.PtrUint32(6112660, 2386504) == 0 {
		*memmap.PtrUint32(6112660, 2386504) = uint32(nox_xxx_getNameId_4E3AA0(CString("Diamond")))
		*memmap.PtrUint32(6112660, 2386508) = uint32(nox_xxx_getNameId_4E3AA0(CString("Ruby")))
		*memmap.PtrUint32(6112660, 2386512) = uint32(nox_xxx_getNameId_4E3AA0(CString("Emerald")))
	}
	v5 = a2
	if a2 != 0 && *(*uint32)(unsafe.Pointer(uintptr(a2 + 16))) != 0 {
		v6 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 8))))
		if int32(*(*uint8)(unsafe.Pointer(uintptr(v6 + 8))))&4 != 0 {
			v6 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 12))))
		}
		v27 = int32(*(*uint32)(unsafe.Pointer(uintptr(v6 + 692))))
		v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(v6 + 692))))
	}
	v7 = int32(*(*uint32)(unsafe.Pointer(uintptr((*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v3))), unsafe.Sizeof(uint32(0))*0))) + 8))))
	if v7&256 != 0 {
		v8 = int32(*(*uint32)(unsafe.Pointer(uintptr((*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v3))), unsafe.Sizeof(uint32(0))*0))) + 12))))
		if v8&1 != 0 {
			v9 = float64(uint16(int16(nox_xxx_spellPrice_424C40(int32(**(**uint8)(unsafe.Pointer(uintptr((*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v3))), unsafe.Sizeof(uint32(0))*0))) + 736))))))))
		LABEL_16:
			v31 = float32(v9)
			goto LABEL_17
		}
		if v8&2 != 0 {
			v10 = int32(uintptr(unsafe.Pointer(nox_xxx_objectTypeByID_4E3B60(*(**byte)(unsafe.Pointer(uintptr((*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v3))), unsafe.Sizeof(uint32(0))*0))) + 736)))))))
			if v10 != 0 {
				v31 = float32(float64(*(*uint32)(unsafe.Pointer(uintptr(v10 + 48)))))
			}
			if noxflags.HasGame(noxflags.GameModeQuest) {
				v9 = nox_xxx_gamedataGetFloat_419D40(CString("QuestGuideWorthMultiplier")) * float64(v31)
				goto LABEL_16
			}
		}
	}
LABEL_17:
	if *(*uint32)(unsafe.Pointer(uintptr((*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v3))), unsafe.Sizeof(uint32(0))*0))) + 8)))&0x13001000 != 0 {
		v11 = int32(*(*uint32)(unsafe.Pointer(uintptr((*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v3))), unsafe.Sizeof(uint32(0))*0))) + 692))))
		v12 = 4
		for {
			if *(*uint32)(unsafe.Pointer(uintptr(v11))) != 0 {
				*(*float32)(unsafe.Pointer(&v29)) = float32(float64(*(*int32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v11))) + 20)))))
				if noxflags.HasGame(noxflags.GameModeQuest) {
					v13 = nox_xxx_gamedataGetFloat_419D40(CString("QuestModifierWorthMultiplier")) * float64(*(*float32)(unsafe.Pointer(&v29)))
				} else {
					v13 = float64(*(*float32)(unsafe.Pointer(&v29)))
				}
				v31 = float32(v13 + float64(v31))
			}
			v11 += 4
			v12--
			if v12 == 0 {
				break
			}
		}
		v4 = v27
	}
	v14 = int32(*(*uint32)(unsafe.Pointer(uintptr((*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v3))), unsafe.Sizeof(uint32(0))*0))) + 8))))
	v30 = v31
	if uint32(v14)&0x1000000 != 0 && int32(*(*uint8)(unsafe.Pointer(uintptr((*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v3))), unsafe.Sizeof(uint32(0))*0))) + 12))))&130 != 0 {
		v15 = *(**uint8)(unsafe.Pointer(uintptr((*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v3))), unsafe.Sizeof(uint32(0))*0))) + 736)))
		if int32(*(*uint8)(unsafe.Add(unsafe.Pointer(v15), 2))) == 0 && int32(*v15) != 0 {
			if noxflags.HasGame(noxflags.GameModeQuest) {
				v16 = nox_xxx_gamedataGetFloat_419D40(CString("DefaultAmmoAmountQuest"))
			} else {
				v16 = nox_xxx_gamedataGetFloat_419D40(CString("DefaultAmmoAmount"))
			}
			v26 = float32(v16)
			v17 = int(v26)
			v18 = float64(v17)
			v19 = int32(*(*uint8)(unsafe.Add(unsafe.Pointer(v15), 1)))
			if v19 > v17 {
				v28 = float32(float64(v31) / v18)
				v31 = float32(float64(v19-v17)*float64(v28) + float64(v31))
			} else {
				v31 = float32(float64(*(*uint8)(unsafe.Add(unsafe.Pointer(v15), 1))) / v18 * float64(v31))
			}
		}
	} else if v14&4096 != 0 {
		if *(*uint32)(unsafe.Pointer(uintptr((*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v3))), unsafe.Sizeof(uint32(0))*0))) + 12)))&0x47F0000 != 0 {
			v20 = int32(*(*uint32)(unsafe.Pointer(uintptr((*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v3))), unsafe.Sizeof(uint32(0))*0))) + 736))))
			if int32(*(*uint8)(unsafe.Pointer(uintptr(v20 + 109)))) != 0 {
				v4 = v27
				v31 = float32(float64(*(*uint8)(unsafe.Pointer(uintptr(v20 + 108)))) / float64(*(*uint8)(unsafe.Pointer(uintptr(v20 + 109)))) * float64(v31))
			}
		}
	}
	if v4 != 0 && v5 != 0 && *(*uint32)(unsafe.Pointer(uintptr(v5 + 16))) != 0 && a1 != 0 {
		v31 = v31 * *(*float32)(unsafe.Pointer(uintptr(v4 + 1716)))
		v30 = v30 * *(*float32)(unsafe.Pointer(uintptr(v4 + 1716)))
	}
	v21 = *(**uint16)(unsafe.Pointer(uintptr((*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v3))), unsafe.Sizeof(uint32(0))*0))) + 556)))
	if v21 != nil && int32(*(*uint16)(unsafe.Add(unsafe.Pointer(v21), unsafe.Sizeof(uint16(0))*2))) != 0 {
		v31 = float32(float64(*v21) / float64(*(*uint16)(unsafe.Add(unsafe.Pointer(v21), unsafe.Sizeof(uint16(0))*2))) * float64(v31))
	}
	if v4 != 0 && a1 == 0 {
		if noxflags.HasGame(noxflags.GameModeQuest) {
			v22 = nox_xxx_gamedataGetFloat_419D40(CString("QuestSellMultiplier"))
		} else {
			v22 = float64(*(*float32)(unsafe.Pointer(uintptr(v4 + 1720))))
		}
		v23 = int32(*(*uint16)(unsafe.Pointer(uintptr((*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v3))), unsafe.Sizeof(uint32(0))*0))) + 4))))
		if uint32(uint16(int16(v23))) != *memmap.PtrUint32(6112660, 2386504) && uint32(v23) != *memmap.PtrUint32(6112660, 2386512) && uint32(v23) != *memmap.PtrUint32(6112660, 2386508) {
			v31 = float32(v22 * float64(v31))
			v30 = float32(v22 * float64(v30))
		}
	}
	if noxflags.HasGame(noxflags.GameModeQuest) && a1 == 0 && *(*uint32)(unsafe.Pointer(uintptr((*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v3))), unsafe.Sizeof(uint32(0))*0))) + 8)))&0x3001000 != 0 && int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr((*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v3))), unsafe.Sizeof(uint32(0))*0))) + 748))) + 4))))&1 != 0 {
		v30 = 0.0
	} else if float64(v31) >= 1.0 {
		goto LABEL_64
	}
	v31 = 1.0
LABEL_64:
	if float64(v30) < 1.0 {
		v30 = 1.0
	}
	if a1 == 2 {
		v24 = nox_xxx_gamedataGetFloat_419D40(CString("RepairCoefficient")) * float64(v30-v31)
		v31 = float32(v24)
		if v24 < 1.0 {
			v31 = 1.0
		}
	}
	return int(v31)
}
func sub_50E820(a1 int32, a2 int32) int32 {
	var (
		v3 int32
		v5 int32
	)
	*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint16(0))*1)) = *(*uint16)(unsafe.Pointer(uintptr(a2 + 36)))
	v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
	*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v5))), unsafe.Sizeof(uint16(0))*0)) = 2505
	return nox_xxx_netSendPacket1_4E5390(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v3 + 276))) + 2064)))), int32(uintptr(unsafe.Pointer(&v5))), 4, 0, 1)
}
func nox_xxx_createShopStruct_50E870() *uint32 {
	var (
		v0 *uint32
		v1 *uint32
	)
	v0 = (*uint32)(nox_alloc_class_new_obj_zero((*nox_alloc_class)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&nox_alloc_tradeSession_2386492))))))
	v1 = v0
	if v0 == nil {
		return nil
	}
	*(*uint32)(unsafe.Add(unsafe.Pointer(v0), unsafe.Sizeof(uint32(0))*2)) = 0
	*(*uint32)(unsafe.Add(unsafe.Pointer(v0), unsafe.Sizeof(uint32(0))*3)) = 0
	*(*uint32)(unsafe.Add(unsafe.Pointer(v0), unsafe.Sizeof(uint32(0))*6)) = 0
	*(*uint32)(unsafe.Add(unsafe.Pointer(v0), unsafe.Sizeof(uint32(0))*7)) = 0
	*(*uint32)(unsafe.Add(unsafe.Pointer(v0), unsafe.Sizeof(uint32(0))*8)) = 0
	*(*uint32)(unsafe.Add(unsafe.Pointer(v0), unsafe.Sizeof(uint32(0))*9)) = 0
	*(*uint32)(unsafe.Add(unsafe.Pointer(v0), unsafe.Sizeof(uint32(0))*10)) = 0
	*(*uint32)(unsafe.Add(unsafe.Pointer(v0), unsafe.Sizeof(uint32(0))*11)) = 0
	*(*uint32)(unsafe.Add(unsafe.Pointer(v0), unsafe.Sizeof(uint32(0))*5)) = 0
	*(*uint32)(unsafe.Add(unsafe.Pointer(v0), unsafe.Sizeof(uint32(0))*12)) = uint32(uintptr(unsafe.Pointer(nox_xxx_newObjectByTypeID_4E3810(CString("Gold")))))
	*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*13)) = uint32(uintptr(unsafe.Pointer(nox_xxx_newObjectByTypeID_4E3810(CString("Gold")))))
	*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*15)) = 0
	*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*14)) = dword_5d4594_2386500
	if dword_5d4594_2386500 != 0 {
		*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_2386500 + 60))) = uint32(uintptr(unsafe.Pointer(v1)))
	}
	dword_5d4594_2386500 = uint32(uintptr(unsafe.Pointer(v1)))
	return v1
}
func nox_xxx_loadShopItems_50E970(a1 int32) {
	var (
		v1  int32
		v2  int32
		v3  int32
		v4  uint32
		v5  *uint32
		v6  **byte
		v7  *uint8
		v8  *uint32
		v9  int32
		v10 *uint32
		v11 int32
		v12 *uint32
	)
	_ = v12
	var v13 int32
	var v14 uint32
	var v15 *uint32
	var v16 *uint32
	var v17 *uint32
	var v18 *uint32
	var v19 *uint32
	var v20 *uint32
	var v21 *uint32
	var v22 *uint32
	var v23 *uint32
	var v24 *uint32
	var v25 *uint32
	var v26 *uint32
	var v27 *uint32
	var v28 *uint32
	var v29 *uint32
	var v30 *uint32
	var v31 *uint32
	var v32 *uint32
	var v33 *uint8
	var v34 *uint8
	var v35 *uint32
	var v36 int32
	var v37 int32
	var v38 int32
	var v39 int32
	var v40 float32
	var v41 int32
	var v42 int32
	var v43 *uint8
	var v44 [20]uint8
	v1 = a1
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 8))))
	if v2 != 0 && int32(*(*uint8)(unsafe.Pointer(uintptr(v2 + 8))))&2 != 0 && int32(*(*uint8)(unsafe.Pointer(uintptr(v2 + 12))))&8 != 0 || (func() int32 {
		v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 12))))
		return v2
	}()) != 0 && int32(*(*uint8)(unsafe.Pointer(uintptr(v2 + 8))))&2 != 0 && int32(*(*uint8)(unsafe.Pointer(uintptr(v2 + 12))))&8 != 0 {
		v3 = v2
		if noxflags.HasGame(noxflags.GameModeQuest) {
			v40 = float32(nox_xxx_gamedataGetFloat_419D40(CString("ShopAnkhCutoffStage")))
			v4 = uint32(int(v40))
			if uint32(nox_game_getQuestStage_4E3CC0()) < v4 {
				v5 = (*uint32)(unsafe.Pointer(nox_xxx_newObjectByTypeID_4E3810(CString("AnkhTradable"))))
				if v5 != nil {
					nox_xxx_addItemToShopSession_50EE00(a1, *(*float32)(unsafe.Pointer(&v5)))
				}
			}
			if *memmap.PtrUint32(0x587000, 234816) != 0 {
				v6 = (**byte)(memmap.PtrOff(0x587000, 234816))
				v7 = (*uint8)(memmap.PtrOff(0x587000, 234816))
				for {
					v8 = (*uint32)(unsafe.Pointer(nox_xxx_newObjectByTypeID_4E3810(*v6)))
					if v8 != nil {
						nox_xxx_addItemToShopSession_50EE00(a1, *(*float32)(unsafe.Pointer(&v8)))
					}
					v9 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v7))), unsafe.Sizeof(uint32(0))*1))))
					v7 = (*uint8)(unsafe.Add(unsafe.Pointer(v7), 4))
					v6 = (**byte)(unsafe.Pointer(v7))
					if v9 == 0 {
						break
					}
				}
			}
			v10 = (*uint32)(unsafe.Pointer(nox_xxx_newObjectByTypeID_4E3810(CString("RewardMarker"))))
			v11 = int32(uintptr(unsafe.Pointer(v10)))
			if v10 != nil {
				v12 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v10), unsafe.Sizeof(uint32(0))*173)))))
				v13 = nox_game_getQuestStage_4E3CC0()
				*v12 = 8
				v14 = uint32(v13 + 2)
				v15 = nox_server_rewardgen_activateMarker_4F0720(v11, uint32(v13+2))
				if v15 != nil {
					nox_xxx_addItemToShopSession_50EE00(a1, *(*float32)(unsafe.Pointer(&v15)))
				}
				*v12 = 8
				v16 = nox_server_rewardgen_activateMarker_4F0720(v11, v14)
				if v16 != nil {
					nox_xxx_addItemToShopSession_50EE00(a1, *(*float32)(unsafe.Pointer(&v16)))
				}
				*v12 = 8
				v17 = nox_server_rewardgen_activateMarker_4F0720(v11, v14)
				if v17 != nil {
					nox_xxx_addItemToShopSession_50EE00(a1, *(*float32)(unsafe.Pointer(&v17)))
				}
				*v12 = 8
				v18 = nox_server_rewardgen_activateMarker_4F0720(v11, v14)
				if v18 != nil {
					nox_xxx_addItemToShopSession_50EE00(a1, *(*float32)(unsafe.Pointer(&v18)))
				}
				*v12 = 16
				v19 = nox_server_rewardgen_activateMarker_4F0720(v11, v14)
				if v19 != nil {
					nox_xxx_addItemToShopSession_50EE00(a1, *(*float32)(unsafe.Pointer(&v19)))
				}
				*v12 = 16
				v20 = nox_server_rewardgen_activateMarker_4F0720(v11, v14)
				if v20 != nil {
					nox_xxx_addItemToShopSession_50EE00(a1, *(*float32)(unsafe.Pointer(&v20)))
				}
				*v12 = 16
				v21 = nox_server_rewardgen_activateMarker_4F0720(v11, v14)
				if v21 != nil {
					nox_xxx_addItemToShopSession_50EE00(a1, *(*float32)(unsafe.Pointer(&v21)))
				}
				*v12 = 16
				v22 = nox_server_rewardgen_activateMarker_4F0720(v11, v14)
				if v22 != nil {
					nox_xxx_addItemToShopSession_50EE00(a1, *(*float32)(unsafe.Pointer(&v22)))
				}
				*v12 = 1
				v23 = nox_server_rewardgen_activateMarker_4F0720(v11, v14)
				if v23 != nil {
					nox_xxx_addItemToShopSession_50EE00(a1, *(*float32)(unsafe.Pointer(&v23)))
				}
				*v12 = 1
				v24 = nox_server_rewardgen_activateMarker_4F0720(v11, v14)
				if v24 != nil {
					nox_xxx_addItemToShopSession_50EE00(a1, *(*float32)(unsafe.Pointer(&v24)))
				}
				*v12 = 1
				v25 = nox_server_rewardgen_activateMarker_4F0720(v11, v14)
				if v25 != nil {
					nox_xxx_addItemToShopSession_50EE00(a1, *(*float32)(unsafe.Pointer(&v25)))
				}
				*v12 = 1
				v26 = nox_server_rewardgen_activateMarker_4F0720(v11, v14)
				if v26 != nil {
					nox_xxx_addItemToShopSession_50EE00(a1, *(*float32)(unsafe.Pointer(&v26)))
				}
				*v12 = 1
				v27 = nox_server_rewardgen_activateMarker_4F0720(v11, v14)
				if v27 != nil {
					nox_xxx_addItemToShopSession_50EE00(a1, *(*float32)(unsafe.Pointer(&v27)))
				}
				*v12 = 1
				v28 = nox_server_rewardgen_activateMarker_4F0720(v11, v14)
				if v28 != nil {
					nox_xxx_addItemToShopSession_50EE00(a1, *(*float32)(unsafe.Pointer(&v28)))
				}
				*v12 = 4
				v29 = nox_server_rewardgen_activateMarker_4F0720(v11, v14)
				if v29 != nil {
					nox_xxx_addItemToShopSession_50EE00(a1, *(*float32)(unsafe.Pointer(&v29)))
				}
				*v12 = 4
				v30 = nox_server_rewardgen_activateMarker_4F0720(v11, v14)
				if v30 != nil {
					nox_xxx_addItemToShopSession_50EE00(a1, *(*float32)(unsafe.Pointer(&v30)))
				}
				*v12 = 4
				v31 = nox_server_rewardgen_activateMarker_4F0720(v11, v14)
				if v31 != nil {
					nox_xxx_addItemToShopSession_50EE00(a1, *(*float32)(unsafe.Pointer(&v31)))
				}
				if noxRndCounter1.IntClamp(0, 100) > 90 {
					*v12 = 2
					v32 = nox_server_rewardgen_activateMarker_4F0720(v11, v14)
					if v32 != nil {
						nox_xxx_addItemToShopSession_50EE00(a1, *(*float32)(unsafe.Pointer(&v32)))
					}
				}
				nox_xxx_delayedDeleteObject_4E5CC0((*nox_object_t)(unsafe.Pointer(uintptr(v11))))
			}
		} else {
			v33 = *(**uint8)(unsafe.Pointer(uintptr(v3 + 692)))
			v43 = v33
			if v33 != nil {
				if int32(*v33) != 0 {
					v42 = 0
					if int32(*v33) != 0 {
						v34 = (*uint8)(unsafe.Add(unsafe.Pointer(v33), 8))
						for {
							v41 = 0
							if int32(*v34) != 0 {
								for {
									v35 = (*uint32)(unsafe.Pointer(nox_xxx_newObjectByTypeID_4E3810(*(**byte)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v34))), -int(unsafe.Sizeof(uint32(0))*1)))) + 4))))))
									v36 = int32(uintptr(unsafe.Pointer(v35)))
									if v35 != nil {
										if *(*uint32)(unsafe.Add(unsafe.Pointer(v35), unsafe.Sizeof(uint32(0))*2))&0x13001000 != 0 {
											v37 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v34))), unsafe.Sizeof(uint32(0))*2))))
											v38 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v34))), unsafe.Sizeof(uint32(0))*4))))
											*(*uint32)(unsafe.Pointer(&v44[4])) = *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v34))), unsafe.Sizeof(uint32(0))*3)))
											*(*uint32)(unsafe.Pointer(&v44[0])) = uint32(v37)
											v39 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v34))), unsafe.Sizeof(uint32(0))*5))))
											*(*uint32)(unsafe.Pointer(&v44[8])) = uint32(v38)
											*(*uint32)(unsafe.Pointer(&v44[12])) = uint32(v39)
											nox_xxx_modifSetItemAttrs_4E4990((*nox_object_t)(unsafe.Pointer(uintptr(v36))), (*int32)(unsafe.Pointer(&v44[0])))
										}
										if cgoFuncAddr(cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(v36 + 704))), (*func(*int32) int32)(nil))) == cgoFuncAddr(nox_xxx_XFerSpellReward_4F5F30) {
											**(**uint8)(unsafe.Pointer(uintptr(v36 + 736))) = *(*uint8)(unsafe.Add(unsafe.Pointer(v34), 4))
										}
										if cgoFuncAddr(cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(v36 + 704))), (*func(*int32) int32)(nil))) == cgoFuncAddr(nox_xxx_XFerAbilityReward_4F6240) {
											**(**uint8)(unsafe.Pointer(uintptr(v36 + 736))) = *(*uint8)(unsafe.Add(unsafe.Pointer(v34), 4))
										}
										if cgoFuncAddr(cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(v36 + 704))), (*func(*int32) int32)(nil))) == cgoFuncAddr(nox_xxx_XFerFieldGuide_4F6390) {
											libc.StrCpy(*(**byte)(unsafe.Pointer(uintptr(v36 + 736))), (*byte)(unsafe.Pointer(uintptr(nox_xxx_getUnitNameByThingType_4E3A80(int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v34))), unsafe.Sizeof(uint32(0))*1)))))))))
											v1 = a1
											v33 = v43
										}
										nox_xxx_addItemToShopSession_50EE00(v1, *(*float32)(unsafe.Pointer(&v36)))
									}
									v41++
									if v41 >= int32(*v34) {
										break
									}
								}
							}
							v34 = (*uint8)(unsafe.Add(unsafe.Pointer(v34), 28))
							v42++
							if v42 >= int32(*v33) {
								break
							}
						}
					}
				}
			}
		}
	}
}
func nox_xxx_addItemToShopSession_50EE00(a1 int32, a2 float32) *float32 {
	var (
		result *float32
		v3     *float32
		v4     uint32
		v5     int32
		i      *uint32
		v7     int32
	)
	result = (*float32)(nox_alloc_class_new_obj_zero((*nox_alloc_class)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&nox_alloc_tradeItems_2386496))))))
	v3 = result
	if result != nil {
		*result = a2
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(result))), unsafe.Sizeof(uint32(0))*1))) = uint32(nox_xxx_shopGetItemCost_50E3D0(1, a1, a2))
		v4 = uint32(sub_50EEC0((*uint32)(unsafe.Pointer(v3))))
		if *(*uint32)(unsafe.Pointer(uintptr(a1 + 20))) != 0 {
			if v4 <= uint32(sub_50EEC0(*(**uint32)(unsafe.Pointer(uintptr(a1 + 20))))) {
				*(*float32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(float32(0))*3)) = 0.0
				*(*float32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(float32(0))*2)) = *(*float32)(unsafe.Pointer(uintptr(a1 + 20)))
				*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 20))) + 12))) = uint32(uintptr(unsafe.Pointer(v3)))
				*(*uint32)(unsafe.Pointer(uintptr(a1 + 20))) = uint32(uintptr(unsafe.Pointer(v3)))
				result = v3
			} else {
				v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 20))))
				for i = *(**uint32)(unsafe.Pointer(uintptr(v5 + 8))); i != nil; i = *(**uint32)(unsafe.Pointer(uintptr(v5 + 8))) {
					if v4 <= uint32(sub_50EEC0(i)) {
						break
					}
					v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(v5 + 8))))
				}
				*(*float32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(float32(0))*2)) = *(*float32)(unsafe.Pointer(uintptr(v5 + 8)))
				v7 = int32(*(*uint32)(unsafe.Pointer(uintptr(v5 + 8))))
				if v7 != 0 {
					*(*uint32)(unsafe.Pointer(uintptr(v7 + 12))) = uint32(uintptr(unsafe.Pointer(v3)))
				}
				*(*uint32)(unsafe.Pointer(uintptr(v5 + 8))) = uint32(uintptr(unsafe.Pointer(v3)))
				*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v3))), unsafe.Sizeof(uint32(0))*3))) = uint32(v5)
				result = v3
			}
		} else {
			*(*float32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(float32(0))*2)) = 0.0
			*(*float32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(float32(0))*3)) = 0.0
			*(*uint32)(unsafe.Pointer(uintptr(a1 + 20))) = uint32(uintptr(unsafe.Pointer(v3)))
			result = v3
		}
	}
	return result
}
func sub_50EEC0(a1 *uint32) int32 {
	var (
		v1 int32
		v2 int32
		v3 *uint8
	)
	v1 = 0x1001000
	v2 = math.MaxUint8
	v3 = (*uint8)(memmap.PtrOff(0x581450, 10308))
	for {
		if *(*uint32)(unsafe.Pointer(uintptr(*a1 + 8)))&uint32(v1) != 0 && (*(*uint32)(unsafe.Pointer(v3)) == 0 || *(*uint32)(unsafe.Pointer(v3))&*(*uint32)(unsafe.Pointer(uintptr(*a1 + 12))) != 0) {
			break
		}
		v1 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v3))), unsafe.Sizeof(uint32(0))*1))))
		v3 = (*uint8)(unsafe.Add(unsafe.Pointer(v3), 8))
		v2--
		if v1 == 0 {
			break
		}
	}
	return int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*1)) | uint32(v2<<24))
}
func sub_50F2B0(a1 int32, a2 *uint32) int32 {
	var (
		v2 int32
		v3 int32
		v4 int32
		v5 int32
		v6 *uint16
		i  int32
		v9 [18]byte
	)
	v2 = int32(*a2)
	v3 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*1)))
	v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
	v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(*a2 + 692))))
	v9[0] = 201
	v9[1] = 8
	*(*uint16)(unsafe.Pointer(&v9[4])) = *(*uint16)(unsafe.Pointer(uintptr(v2 + 36)))
	*(*uint16)(unsafe.Pointer(&v9[2])) = *(*uint16)(unsafe.Pointer(uintptr(v2 + 4)))
	*(*uint32)(unsafe.Pointer(&v9[6])) = uint32(v3)
	v6 = *(**uint16)(unsafe.Pointer(uintptr(v2 + 556)))
	if v6 != nil {
		*(*uint32)(unsafe.Pointer(&v9[10])) = uint32(*v6)
	} else {
		*(*uint32)(unsafe.Pointer(&v9[10])) = 0
	}
	if *(*uint32)(unsafe.Pointer(uintptr(v2 + 8)))&0x13001000 != 0 {
		for i = 0; i < 4; i++ {
			if *(*uint32)(unsafe.Pointer(uintptr(v5))) != 0 {
				v9[i+14] = byte(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v5))) + 4))))
			} else {
				v9[i+14] = math.MaxUint8
			}
			v5 += 4
		}
	} else {
		*(*uint32)(unsafe.Pointer(&v9[14])) = math.MaxUint32
	}
	return nox_xxx_netSendPacket1_4E5390(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v4 + 276))) + 2064)))), int32(uintptr(unsafe.Pointer(&v9[0]))), 18, 0, 1)
}
func sub_50F3A0(a1 *uint32) {
	var (
		i *int32
		j *int32
	)
	for i = (*int32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*8))))); i != nil; i = (*int32)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(i), unsafe.Sizeof(int32(0))*2))))) {
		nox_xxx_inventoryPutImpl_4F3070(int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*2))), (*nox_object_t)(unsafe.Pointer(uintptr(*i))), 1)
	}
	nox_xxx_playerAddGold_4FA590(int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*2))), int32(**(**uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*12)) + 692)))))
	sub_50F450(int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*2))))
	for j = (*int32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*9))))); j != nil; j = (*int32)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(j), unsafe.Sizeof(int32(0))*2))))) {
		nox_xxx_inventoryPutImpl_4F3070(int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*3))), (*nox_object_t)(unsafe.Pointer(uintptr(*j))), 1)
	}
	nox_xxx_playerAddGold_4FA590(int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*3))), int32(**(**uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*13)) + 692)))))
	sub_50F450(int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*3))))
	sub_50F490(a1, int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*2))))
	sub_50F490(a1, int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*3))))
	sub_510000(int32(uintptr(unsafe.Pointer(a1))))
}
func sub_50F450(a1 int32) int32 {
	var (
		v2 int32
		v4 [2]byte
	)
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
	*(*uint16)(unsafe.Pointer(&v4[0])) = 457
	return nox_xxx_netSendPacket0_4E5420(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 276))) + 2064)))), unsafe.Pointer(&v4[0]), 2, 0, 1)
}
func sub_50F490(a1 *uint32, a2 int32) int32 {
	var result int32
	result = a2
	*a1 = 0
	if int32(*(*uint8)(unsafe.Pointer(uintptr(a2 + 8))))&4 != 0 {
		result = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 748))))
		if *(**uint32)(unsafe.Pointer(uintptr(result + 280))) == a1 {
			*(*uint32)(unsafe.Pointer(uintptr(result + 280))) = 0
		}
	}
	return result
}
func nox_xxx_shopExit_50F4C0(a1 *uint32) {
	var (
		v1 int32
		v2 int32
		v3 int32
	)
	sub_50F490(a1, int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*2))))
	sub_50F490(a1, int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*3))))
	v1 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*2)))
	if int32(*(*uint8)(unsafe.Pointer(uintptr(v1 + 8))))&4 != 0 {
		nox_xxx_unitUnFreeze_4E7A60((*nox_object_t)(unsafe.Pointer(uintptr(v1))), 0)
	} else {
		nox_xxx_unitUnFreeze_4E7A60((*nox_object_t)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*3))))), 0)
	}
	v2 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*2)))
	if int32(*(*uint8)(unsafe.Pointer(uintptr(v2 + 8))))&4 != 0 {
		nox_xxx_sendEndTradeMsg_50F560(v2)
		v3 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*2)))
	} else {
		nox_xxx_sendEndTradeMsg_50F560(int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*3))))
		v3 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*3)))
	}
	if noxflags.HasGame(noxflags.GameModeQuest) {
		*memmap.PtrUint32(6112660, uint32(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v3 + 748))) + 276))) + 2064))))*4)+0x2469BC) = uint32(uintptr(unsafe.Pointer(a1)))
	} else {
		sub_510000(int32(uintptr(unsafe.Pointer(a1))))
	}
}
func nox_xxx_sendEndTradeMsg_50F560(a1 int32) int32 {
	var (
		v2 int32
		v4 int16
	)
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
	v4 = 713
	return nox_xxx_netSendPacket1_4E5390(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 276))) + 2064)))), int32(uintptr(unsafe.Pointer(&v4))), 2, 0, 1)
}
func nox_xxx_tradeAccept_50F5A0(a1 int32, a2 int32) {
	var (
		v2 int32
		v3 int32
		v4 int32
		v5 int32
		v6 int32
		v7 int32
	)
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 8))))
	if v2 == a2 {
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 24))) = 1
	} else if *(*uint32)(unsafe.Pointer(uintptr(a1 + 12))) == uint32(a2) {
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 28))) = 1
	}
	if int32(*(*uint8)(unsafe.Pointer(uintptr(v2 + 8))))&4 != 0 {
		sub_50F720(v2, (*uint32)(unsafe.Pointer(uintptr(a1))))
	}
	v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 12))))
	if int32(*(*uint8)(unsafe.Pointer(uintptr(v3 + 8))))&4 != 0 {
		sub_50F720(v3, (*uint32)(unsafe.Pointer(uintptr(a1))))
	}
	if *(*uint32)(unsafe.Pointer(uintptr(a1 + 24))) == 1 && *(*uint32)(unsafe.Pointer(uintptr(a1 + 28))) == 1 {
		sub_50F790(a1, int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 12)))), *(**int32)(unsafe.Pointer(uintptr(a1 + 32))))
		sub_50F790(a1, int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 8)))), *(**int32)(unsafe.Pointer(uintptr(a1 + 36))))
		sub_50F7F0(int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 12)))), int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 48)))))
		sub_50F7F0(int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 8)))), int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 52)))))
		sub_50F6B0(int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 32)))))
		v7 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 36))))
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 32))) = 0
		sub_50F6B0(v7)
		v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 48))))
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 36))) = 0
		**(**uint32)(unsafe.Pointer(uintptr(v4 + 692))) = 0
		**(**uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 52))) + 692))) = 0
		v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 16))))
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 40))) = 0
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 44))) = 0
		if v5 == 1 {
			v6 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 8))))
			if int32(*(*uint8)(unsafe.Pointer(uintptr(v6 + 8))))&4 != 0 {
				sub_50F6E0(v6)
			} else {
				sub_50F6E0(int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 12)))))
			}
		} else {
			sub_50F3A0((*uint32)(unsafe.Pointer(uintptr(a1))))
		}
	}
}
func sub_50F6B0(a1 int32) int32 {
	var (
		result int32
		v2     int32
	)
	result = a1
	if a1 != 0 {
		for {
			v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(result + 8))))
			nox_alloc_class_free_obj_first((*nox_alloc_class)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&nox_alloc_tradeItems_2386496)))), unsafe.Pointer(uintptr(result)))
			result = v2
			if v2 == 0 {
				break
			}
		}
	}
	return result
}
func sub_50F6E0(a1 int32) int32 {
	var (
		v2 int32
		v4 int16
	)
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
	v4 = 1993
	return nox_xxx_netSendPacket1_4E5390(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 276))) + 2064)))), int32(uintptr(unsafe.Pointer(&v4))), 2, 0, 1)
}
func sub_50F720(a1 int32, a2 *uint32) int32 {
	var (
		v2 int32
		v3 int32
		v4 int32
		v5 int8
		v6 int8
	)
	v2 = a1
	v3 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*6)))
	v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
	v5 = 0
	*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&a1))), unsafe.Sizeof(uint16(0))*0)) = 969
	*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a1))), 2)) = 0
	if v3 != 0 {
		v5 = int8(bool2int(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*2)) != uint32(v2)) + 1)
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a1))), 2)) = uint8(int8(bool2int(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*2)) != uint32(v2)) + 1))
	}
	if *(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*7)) != 0 {
		if *(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*3)) == uint32(v2) {
			v6 = int8(int32(v5) | 1)
		} else {
			v6 = int8(int32(v5) | 2)
		}
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a1))), 2)) = uint8(v6)
	}
	return nox_xxx_netSendPacket1_4E5390(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v4 + 276))) + 2064)))), int32(uintptr(unsafe.Pointer(&a1))), 3, 0, 1)
}
func sub_50F790(a1 int32, a2 int32, a3 *int32) {
	var (
		i  *int32
		v4 *float32
		v5 int32
	)
	for i = a3; i != nil; i = (*int32)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(i), unsafe.Sizeof(int32(0))*2))))) {
		if int32(*(*uint8)(unsafe.Pointer(uintptr(a2 + 8))))&4 != 0 {
			nox_xxx_inventoryPutImpl_4F3070(a2, (*nox_object_t)(unsafe.Pointer(uintptr(*i))), 1)
		} else {
			v4 = nox_xxx_addItemToShopSession_50EE00(a1, *(*float32)(unsafe.Pointer(i)))
			if v4 != nil {
				v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 8))))
				if int32(*(*uint8)(unsafe.Pointer(uintptr(v5 + 8))))&4 != 0 {
					sub_50F2B0(v5, (*uint32)(unsafe.Pointer(v4)))
				} else {
					sub_50F2B0(int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 12)))), (*uint32)(unsafe.Pointer(v4)))
				}
			}
		}
	}
}
func sub_50F7F0(a1 int32, a2 int32) *uint32 {
	var result *uint32
	result = (*uint32)(unsafe.Pointer(uintptr(a1)))
	if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))&4 != 0 {
		result = nox_xxx_playerAddGold_4FA590(a1, int32(**(**uint32)(unsafe.Pointer(uintptr(a2 + 692)))))
	}
	return result
}
func nox_xxx_tradeP2PUpdStuff_50FA00(a1 int32, a2 *uint32) int32 {
	var (
		v2 int32
		v3 int32
		v4 int32
		v5 int32
		v7 [14]byte
	)
	v7[0] = 201
	v2 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*2)))
	v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
	v7[1] = 6
	if v2 == a1 {
		v4 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*3)))
		*(*uint32)(unsafe.Pointer(&v7[2])) = **(**uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*12)) + 692)))
		if int32(*(*uint8)(unsafe.Pointer(uintptr(v4 + 8))))&4 != 0 {
			v5 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*13)))
			*(*uint32)(unsafe.Pointer(&v7[6])) = 0
		} else {
			v5 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*13)))
			*(*uint32)(unsafe.Pointer(&v7[6])) = *(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*11)) - *(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*10))
		}
		*(*uint32)(unsafe.Pointer(&v7[10])) = **(**uint32)(unsafe.Pointer(uintptr(v5 + 692)))
	} else if *(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*3)) == uint32(a1) {
		*(*uint32)(unsafe.Pointer(&v7[2])) = **(**uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*13)) + 692)))
		if int32(*(*uint8)(unsafe.Pointer(uintptr(v2 + 8))))&4 != 0 {
			*(*uint32)(unsafe.Pointer(&v7[6])) = 0
		} else {
			*(*uint32)(unsafe.Pointer(&v7[6])) = *(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*10)) - *(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*11))
		}
		*(*uint32)(unsafe.Pointer(&v7[10])) = **(**uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*12)) + 692)))
	}
	return nox_xxx_netSendPacket1_4E5390(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v3 + 276))) + 2064)))), int32(uintptr(unsafe.Pointer(&v7[0]))), 14, 0, 1)
}
func sub_50FB90(a1 *uint32) *uint32 {
	var (
		v1     int32
		result *uint32
		v3     *int32
		v4     int32
		v5     *int32
		v6     int32
		v7     int32
		v8     uint32
		v9     uint32
		v10    uint32
		v11    int32
		v12    uint32
		v13    uint32
		v14    int32
	)
	v1 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*2)))
	result = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v1 + 8))) & 4)))
	if uintptr(unsafe.Pointer(result)) != uintptr(4) || (int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*3)) + 8))))&4) != 4 {
		v3 = *(**int32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*12)) + 692)))
		if result != nil {
			nox_xxx_playerAddGold_4FA590(v1, *v3)
		}
		*v3 = 0
		v4 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*3)))
		v5 = *(**int32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*13)) + 692)))
		if int32(*(*uint8)(unsafe.Pointer(uintptr(v4 + 8))))&4 != 0 {
			nox_xxx_playerAddGold_4FA590(v4, *v5)
		}
		*v5 = 0
		v6 = sub_50FD20(a1, int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*2))))
		v7 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*3)))
		*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*10)) = uint32(v6)
		result = (*uint32)(unsafe.Pointer(uintptr(sub_50FD20(a1, v7))))
		v8 = *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*10))
		*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*11)) = uint32(uintptr(unsafe.Pointer(result)))
		if uint32(uintptr(unsafe.Pointer(result))) <= v8 {
			if uint32(uintptr(unsafe.Pointer(result))) < v8 {
				v12 = v8 - uint32(uintptr(unsafe.Pointer(result)))
				if int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*3)) + 8))))&4 != 0 {
					v13 = uint32(nox_xxx_playerGetGold_4FA6B0(int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*3)))))
					if v13 < v12 {
						*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*11)) += v13
						**(**uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*13)) + 692))) += v13
						result = nox_xxx_playerSubGold_4FA5D0(int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*3))), v13)
					} else {
						v14 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*13)))
						*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*11)) += v12
						**(**uint32)(unsafe.Pointer(uintptr(v14 + 692))) += v12
						result = nox_xxx_playerSubGold_4FA5D0(int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*3))), v12)
					}
				} else {
					*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*11)) = v8
					result = *(**uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*13)) + 692)))
					*result += v12
				}
			}
		} else {
			v9 = uint32(uintptr(unsafe.Pointer(result))) - v8
			if int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*2)) + 8))))&4 != 0 {
				v10 = uint32(nox_xxx_playerGetGold_4FA6B0(int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*2)))))
				if v10 < v9 {
					*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*10)) += v10
					**(**uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*12)) + 692))) += v10
					result = nox_xxx_playerSubGold_4FA5D0(int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*2))), v10)
				} else {
					*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*10)) += v9
					**(**uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*12)) + 692))) += v9
					result = nox_xxx_playerSubGold_4FA5D0(int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*2))), v9)
				}
			} else {
				v11 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*12)))
				*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*10)) = v9 + v8
				result = *(**uint32)(unsafe.Pointer(uintptr(v11 + 692)))
				*result += v9
			}
		}
	}
	return result
}
func sub_50FD20(a1 *uint32, a2 int32) int32 {
	var (
		v2     int32
		v3     int32
		result int32
		v5     int32
	)
	if *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*2)) == uint32(a2) {
		v2 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*12)))
		v3 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*8)))
	} else {
		v2 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*13)))
		v3 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*9)))
	}
	for result = int32(**(**uint32)(unsafe.Pointer(uintptr(v2 + 692)))); v3 != 0; result += v5 {
		v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(v3 + 4))))
		v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(v3 + 8))))
	}
	return result
}
func nox_xxx_tradeP2PAddOfferMB_50FE20(a1 int32, a2 int32) int32 {
	var (
		v2     *int32
		v3     int32
		result int32
		v5     *float32
		v6     int32
		v7     int32
		v8     int32
		v9     int32
		v10    int32
	)
	v2 = (*int32)(unsafe.Pointer(sub_50FFE0(*(**uint32)(unsafe.Pointer(uintptr(a1 + 32))), a2)))
	if v2 != nil {
		v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 8))))
	} else {
		result = int32(uintptr(unsafe.Pointer(sub_50FFE0(*(**uint32)(unsafe.Pointer(uintptr(a1 + 36))), a2))))
		v2 = (*int32)(unsafe.Pointer(uintptr(result)))
		if result == 0 {
			return result
		}
		v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 12))))
	}
	if int32(*(*uint8)(unsafe.Pointer(uintptr(v3 + 8))))&4 != 0 {
		nox_xxx_inventoryPutImpl_4F3070(v3, (*nox_object_t)(unsafe.Pointer(uintptr(*v2))), 1)
	} else {
		v5 = nox_xxx_addItemToShopSession_50EE00(a1, *(*float32)(unsafe.Pointer(v2)))
		if v5 != nil {
			v6 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 8))))
			if (int32(*(*uint8)(unsafe.Pointer(uintptr(v6 + 8)))) & 4) == 0 {
				v6 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 12))))
			}
			sub_50F2B0(v6, (*uint32)(unsafe.Pointer(v5)))
		}
	}
	v7 = *(*int32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(int32(0))*2))
	if v7 != 0 {
		*(*uint32)(unsafe.Pointer(uintptr(v7 + 12))) = uint32(*(*int32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(int32(0))*3)))
	}
	v8 = *(*int32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(int32(0))*3))
	if v8 != 0 {
		*(*uint32)(unsafe.Pointer(uintptr(v8 + 8))) = uint32(*(*int32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(int32(0))*2)))
	}
	if *(*uint32)(unsafe.Pointer(uintptr(a1 + 8))) == uint32(v3) {
		if *(**int32)(unsafe.Pointer(uintptr(a1 + 32))) == v2 {
			*(*uint32)(unsafe.Pointer(uintptr(a1 + 32))) = uint32(*(*int32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(int32(0))*2)))
		}
	} else if *(**int32)(unsafe.Pointer(uintptr(a1 + 36))) == v2 {
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 36))) = uint32(*(*int32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(int32(0))*2)))
	}
	sub_50FB90((*uint32)(unsafe.Pointer(uintptr(a1))))
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 24))) = 0
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 28))) = 0
	v9 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 8))))
	if int32(*(*uint8)(unsafe.Pointer(uintptr(v9 + 8))))&4 != 0 {
		if (int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 12))) + 8))))&4) == 0 && *(*uint32)(unsafe.Pointer(uintptr(a1 + 44))) <= uint32(*(*int32)(unsafe.Pointer(uintptr(a1 + 40)))) {
			*(*uint32)(unsafe.Pointer(uintptr(a1 + 28))) = 1
		}
	} else if *(*uint32)(unsafe.Pointer(uintptr(a1 + 40))) <= uint32(*(*int32)(unsafe.Pointer(uintptr(a1 + 44)))) {
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 24))) = 1
	}
	if int32(*(*uint8)(unsafe.Pointer(uintptr(v9 + 8))))&4 != 0 {
		nox_xxx_tradeP2PUpdStuff_50FA00(v9, (*uint32)(unsafe.Pointer(uintptr(a1))))
		sub_50FF90(int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 8)))), a1, *v2)
		sub_50F720(int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 8)))), (*uint32)(unsafe.Pointer(uintptr(a1))))
	}
	v10 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 12))))
	if int32(*(*uint8)(unsafe.Pointer(uintptr(v10 + 8))))&4 != 0 {
		nox_xxx_tradeP2PUpdStuff_50FA00(v10, (*uint32)(unsafe.Pointer(uintptr(a1))))
		sub_50FF90(int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 12)))), a1, *v2)
		sub_50F720(int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 12)))), (*uint32)(unsafe.Pointer(uintptr(a1))))
	}
	nox_alloc_class_free_obj_first((*nox_alloc_class)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&nox_alloc_tradeItems_2386496)))), unsafe.Pointer(v2))
	return 1
}
func sub_50FF90(a1 int32, a2 int32, a3 int32) int32 {
	var (
		v4 int32
		v6 int32
	)
	*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v6))), unsafe.Sizeof(uint16(0))*1)) = *(*uint16)(unsafe.Pointer(uintptr(a3 + 36)))
	v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
	*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v6))), unsafe.Sizeof(uint16(0))*0)) = 1481
	return nox_xxx_netSendPacket1_4E5390(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v4 + 276))) + 2064)))), int32(uintptr(unsafe.Pointer(&v6))), 4, 0, 1)
}
func sub_50FFE0(a1 *uint32, a2 int32) *uint32 {
	var result *uint32
	result = a1
	if a1 == nil {
		return nil
	}
	for *(*uint32)(unsafe.Pointer(uintptr(*result + 36))) != uint32(a2) {
		result = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*2)))))
		if result == nil {
			return nil
		}
	}
	return result
}
func sub_510000(a1 int32) {
	var (
		v1 *int32
		v2 *int32
		v3 int32
		v4 int32
	)
	v1 = *(**int32)(unsafe.Pointer(uintptr(a1 + 20)))
	if v1 != nil {
		for {
			v2 = (*int32)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(int32(0))*2)))))
			nox_xxx_objectFreeMem_4E38A0(*v1)
			nox_alloc_class_free_obj_first((*nox_alloc_class)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&nox_alloc_tradeItems_2386496)))), unsafe.Pointer(v1))
			v1 = v2
			if v2 == nil {
				break
			}
		}
	}
	nox_xxx_objectFreeMem_4E38A0(int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 48)))))
	nox_xxx_objectFreeMem_4E38A0(int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 52)))))
	sub_50F6B0(int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 32)))))
	sub_50F6B0(int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 36)))))
	v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 56))))
	if v3 != 0 {
		*(*uint32)(unsafe.Pointer(uintptr(v3 + 60))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 60)))
	}
	v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 60))))
	if v4 != 0 {
		*(*uint32)(unsafe.Pointer(uintptr(v4 + 56))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 56)))
	}
	if uint32(a1) == dword_5d4594_2386500 {
		dword_5d4594_2386500 = *(*uint32)(unsafe.Pointer(uintptr(a1 + 56)))
	}
	nox_alloc_class_free_obj_first((*nox_alloc_class)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&nox_alloc_tradeSession_2386492)))), unsafe.Pointer(uintptr(a1)))
}
func nox_xxx_getSomeShopData_5103A0(a1 int32, a2 int32) int32 {
	var (
		v2 int32
		v3 *uint8
		v4 int32
		i  int32
	)
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 8))))
	if int32(*(*uint8)(unsafe.Pointer(uintptr(v2 + 8))))&4 != 0 {
		v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 12))))
	}
	v3 = *(**uint8)(unsafe.Pointer(uintptr(v2 + 692)))
	v4 = 0
	if int32(*v3) == 0 {
		return -1
	}
	for i = int32(uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v3), 4))))); sub_5103F0(a2, i) == 0; i += 28 {
		if func() int32 {
			p := &v4
			*p++
			return *p
		}() >= int32(*v3) {
			return -1
		}
	}
	return v4
}
func sub_5103F0(a1 int32, a2 int32) int32 {
	var (
		v2 int32
		v3 int32
		v4 *uint32
		v5 int32
	)
	v2 = 0
	if a1 == 0 || a2 == 0 {
		return 0
	}
	if int32(*(*uint16)(unsafe.Pointer(uintptr(a1 + 4)))) == int32(**(**uint16)(unsafe.Pointer(uintptr(a2)))) {
		v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 8))))
		v2 = 1
		if uint32(v3)&0x13001000 != 0 {
			v4 = *(**uint32)(unsafe.Pointer(uintptr(a1 + 692)))
			if *(*uint32)(unsafe.Pointer(uintptr(a2 + 12))) != *v4 {
				v2 = 0
			}
			if *(*uint32)(unsafe.Pointer(uintptr(a2 + 16))) != *(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*1)) {
				v2 = 0
			}
			if *(*uint32)(unsafe.Pointer(uintptr(a2 + 20))) != *(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*2)) {
				v2 = 0
			}
			if *(*uint32)(unsafe.Pointer(uintptr(a2 + 24))) != *(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*3)) {
				v2 = 0
			}
		}
		if v3&256 != 0 {
			v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 12))))
			if v5&1 != 0 || v5&4 != 0 {
				if uint32(**(**uint8)(unsafe.Pointer(uintptr(a1 + 736)))) != *(*uint32)(unsafe.Pointer(uintptr(a2 + 8))) {
					return 0
				}
			} else if libc.StrCmp(*(**byte)(unsafe.Pointer(uintptr(a1 + 736))), (*byte)(unsafe.Pointer(uintptr(nox_xxx_getUnitNameByThingType_4E3A80(int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 8))))))))) != 0 {
				v2 = 0
			}
		}
	}
	return v2
}
func sub_5108D0(a1 int32, a2 int32, a3 int32) *uint32 {
	var (
		v3     int32
		result *uint32
		v5     int32
		v6     int32
		v7     int32
		v8     uint8
		v9     *uint16
		v10    int16
		v11    [8]byte
	)
	v3 = 0
	result = *(**uint32)(unsafe.Pointer(uintptr(a1 + 504)))
	v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
	v11[0] = 201
	v11[1] = 31
	*(*uint16)(unsafe.Pointer(&v11[2])) = uint16(int16(a3))
	if result != nil {
		for *(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*9)) != uint32(a3) {
			result = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*124)))))
			if result == nil {
				return result
			}
		}
		v6 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*2)))
		if v6&4096 != 0 {
			if *(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*3))&0x47F0000 != 0 {
				v7 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*184)))
				v8 = *(*uint8)(unsafe.Pointer(uintptr(v7 + 109)))
				if int32(*(*uint8)(unsafe.Pointer(uintptr(v7 + 108)))) < int32(v8) {
					if int32(v8) != 0 {
						v3 = 1
					}
				}
			}
		}
		v9 = (*uint16)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*139)))))
		if v9 != nil && (func() bool {
			v10 = int16(*(*uint16)(unsafe.Add(unsafe.Pointer(v9), unsafe.Sizeof(uint16(0))*2)))
			return int32(*v9) != int32(v10)
		}()) && int32(v10) != 0 || v3 != 0 {
			*(*uint32)(unsafe.Pointer(&v11[4])) = uint32(nox_xxx_shopGetItemCost_50E3D0(2, a2, *(*float32)(unsafe.Pointer(&result))))
			result = (*uint32)(unsafe.Pointer(uintptr(nox_xxx_netSendPacket0_4E5420(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v5 + 276))) + 2064)))), unsafe.Pointer(&v11[0]), 8, 0, 1))))
		} else {
			result = nox_xxx_aud_501960(925, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
		}
	}
	return result
}
func sub_510AE0(a1 *int32, a2 int32, a3 *uint32) *uint32 {
	var (
		result *uint32
		v4     int32
		v5     int32
		v6     uint32
		v7     int32
		v8     int32
	)
	result = (*uint32)(unsafe.Pointer(uintptr(nox_xxx_playerGetGold_4FA6B0(int32(uintptr(unsafe.Pointer(a1)))))))
	v4 = *(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*126))
	v5 = *(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*187))
	if v4 != 0 {
		result = a3
		for *(**uint32)(unsafe.Pointer(uintptr(v4 + 36))) != a3 {
			v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(v4 + 496))))
			if v4 == 0 {
				return result
			}
		}
		v6 = uint32(nox_xxx_shopGetItemCost_50E3D0(2, a2, *(*float32)(unsafe.Pointer(&v4))))
		nox_xxx_playerSubGold_4FA5D0(int32(uintptr(unsafe.Pointer(a1))), v6)
		nox_xxx_unitSetHP_4E4560((*nox_object_t)(unsafe.Pointer(uintptr(v4))), *(*uint16)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v4 + 556))) + 4))))
		nox_xxx_itemReportHealth_4D87A0(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v5 + 276))) + 2064)))), (*nox_object_t)(unsafe.Pointer(uintptr(v4))))
		v7 = int32(*(*uint32)(unsafe.Pointer(uintptr(v4 + 8))))
		if v7&4096 != 0 && *(*uint32)(unsafe.Pointer(uintptr(v4 + 12)))&0x47F0000 != 0 {
			v8 = int32(*(*uint32)(unsafe.Pointer(uintptr(v4 + 736))))
			if nox_xxx_rechargeItem_53C520(v4, 100) != 0 {
				nox_xxx_netReportCharges_4D82B0(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v5 + 276))) + 2064)))), (*nox_object_t)(unsafe.Pointer(uintptr(v4))), int8(*(*uint8)(unsafe.Pointer(uintptr(v8 + 108)))), int8(*(*uint8)(unsafe.Pointer(uintptr(v8 + 109)))))
			}
		}
		sub_4D8870(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v5 + 276))) + 2064)))), int32(uintptr(unsafe.Pointer(a1))))
		result = nox_xxx_aud_501960(803, (*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1)))))), 2, *(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*9)))
	}
	return result
}
func sub_510D10(a1 *int32, a2 int32, a3 int32, a4 uint32) uint16 {
	var (
		result uint16
		v5     int32
		v6     int32
		v7     int32
		v8     int32
	)
	nox_xxx_playerGetGold_4FA6B0(int32(uintptr(unsafe.Pointer(a1))))
	result = uint16(a4)
	v5 = *(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*187))
	v6 = 0
	if a4 != 0 {
		for {
			v7 = *(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*126))
			if v7 == 0 {
				break
			}
			for {
				result = *(*uint16)(unsafe.Pointer(uintptr(v7 + 4)))
				if int32(result) == a3 {
					break
				}
				v7 = int32(*(*uint32)(unsafe.Pointer(uintptr(v7 + 496))))
				if v7 == 0 {
					return result
				}
			}
			sub_4ED0C0(int32(uintptr(unsafe.Pointer(a1))), (*nox_object_t)(unsafe.Pointer(uintptr(v7))))
			nox_xxx_delayedDeleteObject_4E5CC0((*nox_object_t)(unsafe.Pointer(uintptr(v7))))
			v8 = nox_xxx_shopGetItemCost_50E3D0(0, a2, *(*float32)(unsafe.Pointer(&v7)))
			nox_xxx_playerAddGold_4FA590(int32(uintptr(unsafe.Pointer(a1))), v8)
			sub_4D8870(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v5 + 276))) + 2064)))), int32(uintptr(unsafe.Pointer(a1))))
			result = uint16(a4)
			if uint32(func() int32 {
				p := &v6
				*p++
				return *p
			}()) >= a4 {
				return uint16(uint32(uintptr(unsafe.Pointer(nox_xxx_aud_501960(307, (*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1)))))), 2, *(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*9)))))))
			}
		}
	}
	return result
}
func nox_xxx_shopCancelSession_510DC0(a1 *uint32) {
	if *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*4)) != 0 {
		nox_xxx_shopExit_50F4C0(a1)
	} else {
		sub_50F3A0(a1)
	}
}
func sub_510DE0(a1 int32, a2 int32) int32 {
	var (
		v2 int32
		v3 *uint32
	)
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))) + 280))))
	if v2 == 0 {
		return 0
	}
	v3 = *(**uint32)(unsafe.Pointer(uintptr(v2 + 20)))
	if v3 == nil {
		return 0
	}
	for *(*uint32)(unsafe.Pointer(uintptr(*v3 + 36))) != uint32(a2) {
		v3 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*2)))))
		if v3 == nil {
			return 0
		}
	}
	return int32(*v3)
}
func sub_510E20(a1 int32) {
	if *memmap.PtrUint32(6112660, uint32(a1*4)+0x2469BC) != 0 {
		sub_510000(int32(*memmap.PtrUint32(6112660, uint32(a1*4)+0x2469BC)))
	}
	*memmap.PtrUint32(6112660, uint32(a1*4)+0x2469BC) = 0
}
func sub_510E50() {
	dword_5d4594_2386564 = 0
}
func nox_xxx_updateSentryGlobe_510E60(a1 int32) int32 {
	var (
		result int32
		v2     int32
		v3     float32
		v4     float32
		v5     float32
		v6     float32
		a2     float2
		v8     [2]int32
		a1a    float4
		v10    float4
	)
	result = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 16))))
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
	if result >= 0 {
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 500))) = 0
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 496))) = dword_5d4594_2386564
		if dword_5d4594_2386564 != 0 {
			*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_2386564 + 500))) = uint32(a1)
		}
		dword_5d4594_2386564 = uint32(a1)
		result = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 16))) | 0x80000000)
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 16))) = uint32(result)
	}
	if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 16))))&32 != 0 {
		result = int32(uintptr(unsafe.Pointer(nox_xxx_sentryUpdateList_510FD0((*uint32)(unsafe.Pointer(uintptr(a1)))))))
	}
	if *(*uint32)(unsafe.Pointer(uintptr(a1 + 16)))&0x1000000 != 0 {
		v3 = *(*float32)(unsafe.Pointer(uintptr(a1 + 60)))
		a1a.field_0 = *(*float32)(unsafe.Pointer(uintptr(a1 + 56)))
		a1a.field_4 = v3
		a1a.field_8 = float32(math.Cos(float64(*(*float32)(unsafe.Pointer(uintptr(v2)))))*600.0 + float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 56)))))
		a1a.field_C = float32(math.Sin(float64(*(*float32)(unsafe.Pointer(uintptr(v2)))))*600.0 + float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 60)))))
		if nox_xxx_mapTraceRay_535250(&a1a, &a2, nil, 5) != 0 {
			v4 = a1a.field_C
			*(*float32)(unsafe.Pointer(uintptr(a1 + 156))) = a1a.field_8
			*(*float32)(unsafe.Pointer(uintptr(a1 + 160))) = v4
		} else {
			v5 = a2.field_0
			v6 = a2.field_4
			*(*float32)(unsafe.Pointer(uintptr(a1 + 156))) = a2.field_0
			*(*float32)(unsafe.Pointer(uintptr(a1 + 160))) = v6
			a1a.field_8 = v5
			a1a.field_C = v6
		}
		*(*float32)(unsafe.Pointer(uintptr(v2))) = *(*float32)(unsafe.Pointer(uintptr(v2 + 8))) + *(*float32)(unsafe.Pointer(uintptr(v2)))
		if float64(a1a.field_0) >= float64(a1a.field_8) {
			v10.field_8 = a1a.field_0
			v10.field_0 = a1a.field_8
		} else {
			v10.field_0 = a1a.field_0
			v10.field_8 = a1a.field_8
		}
		if float64(a1a.field_4) >= float64(a1a.field_C) {
			v10.field_C = a1a.field_4
			v10.field_4 = a1a.field_C
		} else {
			v10.field_4 = a1a.field_4
			v10.field_C = a1a.field_C
		}
		v8[0] = a1
		v8[1] = int32(uintptr(unsafe.Pointer(&a1a)))
		nox_xxx_getUnitsInRect_517C10(&v10, func(arg1 *float32, arg2 int32) {
			nox_xxx_sentry_511020(int32(uintptr(unsafe.Pointer(arg1))), arg2)
		}, unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(&v8[0]))))))
	} else {
		*(*uint32)(unsafe.Pointer(uintptr(v2))) = *(*uint32)(unsafe.Pointer(uintptr(v2 + 4)))
	}
	return result
}
func nox_xxx_sentryUpdateList_510FD0(a1 *uint32) *uint32 {
	var (
		result *uint32
		v2     int32
		v3     int32
	)
	result = a1
	if *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*4)) < 0 {
		v2 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*125)))
		if v2 != 0 {
			*(*uint32)(unsafe.Pointer(uintptr(v2 + 496))) = *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*124))
		} else {
			dword_5d4594_2386564 = *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*124))
		}
		v3 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*124)))
		if v3 != 0 {
			*(*uint32)(unsafe.Pointer(uintptr(v3 + 500))) = *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*125))
		}
	}
	*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*4)) &= math.MaxInt32
	return result
}
func nox_xxx_sentry_511020(a1 int32, a2 int32) {
	var (
		v2 int32
		v3 float64
		v6 int32
		v7 int32
		a3 float2
	)
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 16))))
	if (v2&65) == 0 && ((v2&16) == 0 || noxflags.HasGame(noxflags.GameModeQuest) && int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))&2 != 0 && (*(*uint32)(unsafe.Pointer(uintptr(a1 + 16)))&0x8000) != 0x8000) && *(*uint32)(unsafe.Pointer(uintptr(a1 + 556))) != 0 && nox_xxx_mathPointOnTheLine_57C8A0(*(**float4)(unsafe.Pointer(uintptr(a2 + 4))), (*float2)(unsafe.Pointer(uintptr(a1+56))), &a3) != 0 {
		v3 = float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 60))) - a3.field_4)
		if float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 176)))**(*float32)(unsafe.Pointer(uintptr(a1 + 176)))) > v3*v3+float64((*(*float32)(unsafe.Pointer(uintptr(a1 + 56)))-a3.field_0)*(*(*float32)(unsafe.Pointer(uintptr(a1 + 56)))-a3.field_0)) {
			v7 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2))))
			v6 = int32(uintptr(unsafe.Pointer(nox_xxx_findParentChainPlayer_4EC580((*nox_object_t)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a2))))))))))
			(cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(a1 + 716))), (*func(int32, int32, int32, int32, int32))(nil)))(a1, v6, v7, 500, 16)
			nox_xxx_aud_501960(298, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
		}
	}
}
func sub_511100(a1 int32) {
	var (
		v1  int32
		v2  *byte
		v3  int32
		v4  float64
		v5  float64
		v6  float32
		v7  float32
		v8  float32
		v9  float32
		v10 float32
		v11 float32
		i   float32
		v13 float32
	)
	v1 = a1
	v2 = (*byte)(unsafe.Pointer(nox_common_playerInfoFromNum_417090(a1)))
	v3 = int32(dword_5d4594_2386564)
	v4 = float64(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v2))), unsafe.Sizeof(uint16(0))*5))))
	v9 = float32(float64(*((*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(v2))), unsafe.Sizeof(float32(0))*908)))) - v4)
	v13 = float32(float64(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v2))), unsafe.Sizeof(uint16(0))*6)))))
	v10 = *((*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(v2))), unsafe.Sizeof(float32(0))*909))) - v13
	v11 = float32(v4 + float64(*((*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(v2))), unsafe.Sizeof(float32(0))*908)))))
	for i = v13 + *((*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(v2))), unsafe.Sizeof(float32(0))*909))); v3 != 0; v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(v3 + 496)))) {
		if *(*uint32)(unsafe.Pointer(uintptr(v3 + 16)))&0x1000000 != 0 {
			if float64(*(*float32)(unsafe.Pointer(uintptr(v3 + 56)))) >= float64(*(*float32)(unsafe.Pointer(uintptr(v3 + 156)))) {
				v5 = float64(*(*float32)(unsafe.Pointer(uintptr(v3 + 56))))
				v6 = *(*float32)(unsafe.Pointer(uintptr(v3 + 156)))
			} else {
				v5 = float64(*(*float32)(unsafe.Pointer(uintptr(v3 + 156))))
				v6 = *(*float32)(unsafe.Pointer(uintptr(v3 + 56)))
			}
			if float64(*(*float32)(unsafe.Pointer(uintptr(v3 + 60)))) >= float64(*(*float32)(unsafe.Pointer(uintptr(v3 + 160)))) {
				v8 = *(*float32)(unsafe.Pointer(uintptr(v3 + 60)))
				v7 = *(*float32)(unsafe.Pointer(uintptr(v3 + 160)))
			} else {
				v7 = *(*float32)(unsafe.Pointer(uintptr(v3 + 60)))
				v8 = *(*float32)(unsafe.Pointer(uintptr(v3 + 160)))
			}
			if float64(v9) < v5 && float64(v11) > float64(v6) && float64(v10) < float64(v8) && float64(i) > float64(v7) {
				sub_511250(v1, (*float32)(unsafe.Pointer(uintptr(v3))))
			}
		} else {
			**(**uint32)(unsafe.Pointer(uintptr(v3 + 748))) = *(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v3 + 748))) + 4)))
		}
	}
}
func sub_511250(a1 int32, a2 *float32) int32 {
	var (
		v2 int16
		v3 float32
		v4 int16
		v5 float32
		v7 [9]byte
	)
	v7[0] = 149
	v2 = int16(uint16(sub_419A30(*(*float32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(float32(0))*14)))))
	v3 = *(*float32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(float32(0))*15))
	*(*uint16)(unsafe.Pointer(&v7[1])) = uint16(v2)
	v4 = int16(uint16(sub_419A30(v3)))
	v5 = *(*float32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(float32(0))*39))
	*(*uint16)(unsafe.Pointer(&v7[3])) = uint16(v4)
	*(*uint16)(unsafe.Pointer(&v7[5])) = uint16(sub_419A30(v5))
	*(*uint16)(unsafe.Pointer(&v7[7])) = uint16(sub_419A30(*(*float32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(float32(0))*40))))
	return nox_netlist_addToMsgListCli_40EBC0(a1, 1, (*uint8)(unsafe.Pointer(&v7[0])), 9)
}
func nox_xxx_allocSpringsArray_5112C0() int32 {
	nox_alloc_springs_2386568 = unsafe.Pointer(nox_new_alloc_class(CString("Springs"), 52, 256))
	return bool2int(nox_alloc_springs_2386568 != nil)
}
func sub_5112F0() {
	dword_5d4594_2386572 = 0
	nox_alloc_class_free_all((*nox_alloc_class)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&nox_alloc_springs_2386568)))))
}
func sub_511310() int32 {
	nox_free_alloc_class((*nox_alloc_class)(nox_alloc_springs_2386568))
	dword_5d4594_2386572 = 0
	return 1
}
func sub_511330(a1 int32) {
	var v1 int32
	v1 = int32(dword_5d4594_2386572)
	if dword_5d4594_2386572 != 0 {
		for *(*uint32)(unsafe.Pointer(uintptr(v1 + 8))) != uint32(a1) {
			v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 44))))
			if v1 == 0 {
				return
			}
		}
		if v1 != 0 {
			sub_511360(v1)
		}
	}
}
func sub_511360(a1 int32) {
	var (
		v1 int32
		v2 int32
	)
	v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 48))))
	if v1 != 0 {
		*(*uint32)(unsafe.Pointer(uintptr(v1 + 44))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 44)))
	} else {
		dword_5d4594_2386572 = *(*uint32)(unsafe.Pointer(uintptr(a1 + 44)))
	}
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 44))))
	if v2 != 0 {
		*(*uint32)(unsafe.Pointer(uintptr(v2 + 48))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 48)))
	}
	nox_alloc_class_free_obj_first((*nox_alloc_class)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&nox_alloc_springs_2386568)))), unsafe.Pointer(uintptr(a1)))
}
func nox_xxx_updateDebugObjects_5113A0() {
	var (
		v0  int32
		v1  int32
		v2  int32
		v3  int32
		v4  int32
		v5  int32
		v6  float64
		v7  float64
		v8  float64
		v9  float64
		v10 float32
		v11 float32
		v12 float32
		v13 float32
		v14 float32
		v15 float32
	)
	v0 = int32(dword_5d4594_2386572)
	if dword_5d4594_2386572 != 0 {
		for {
			v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(v0 + 8))))
			v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(v0 + 44))))
			if v1 == 0 {
				goto LABEL_16
			}
			v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(v0 + 12))))
			if v3 == 0 {
				goto LABEL_16
			}
			if (int32(*(*uint8)(unsafe.Pointer(uintptr(v1 + 16))))&32) == 0 && (int32(*(*uint8)(unsafe.Pointer(uintptr(v3 + 16))))&32) == 0 && (func() bool {
				v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(v0 + 12))))
				v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(v0 + 8))))
				v12 = *(*float32)(unsafe.Pointer(uintptr(v4 + 64))) - *(*float32)(unsafe.Pointer(uintptr(v5 + 64)))
				v6 = float64(*(*float32)(unsafe.Pointer(uintptr(v4 + 68))) - *(*float32)(unsafe.Pointer(uintptr(v5 + 68))))
				v14 = float32(v6)
				v7 = math.Sqrt(v6*float64(v14) + float64(v12*v12))
				*(*float32)(unsafe.Pointer(uintptr(v0 + 32))) = float32(v7)
				return v7 <= float64(*(*float32)(unsafe.Pointer(uintptr(v0 + 28))))
			}()) {
				v8 = float64(*(*float32)(unsafe.Pointer(uintptr(v0 + 32))) - *(*float32)(unsafe.Pointer(uintptr(v0 + 24))))
				if (int32(*(*uint8)(unsafe.Pointer(uintptr(v0 + 40))))&1) == 0 || v8 >= 0.0 {
					if v8 > float64(*(*float32)(unsafe.Pointer(uintptr(v0 + 24)))) {
						v8 = float64(*(*float32)(unsafe.Pointer(uintptr(v0 + 24))))
					}
					v9 = -(v8 * float64(*(*float32)(unsafe.Pointer(uintptr(v0 + 16)))))
					v15 = float32(v9 * float64(v14 / *(*float32)(unsafe.Pointer(uintptr(v0 + 32)))))
					v13 = float32(v9 * float64(v12 / *(*float32)(unsafe.Pointer(uintptr(v0 + 32)))))
					v11 = -v15
					v10 = -v13
					sub_548600(int32(*(*uint32)(unsafe.Pointer(uintptr(v0 + 8)))), v10, v11)
					nox_xxx_unitHasCollideOrUpdateFn_537610((*nox_object_t)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v0 + 8)))))))
					sub_548600(int32(*(*uint32)(unsafe.Pointer(uintptr(v0 + 12)))), v13, v15)
					nox_xxx_unitHasCollideOrUpdateFn_537610((*nox_object_t)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v0 + 12)))))))
					*(*uint32)(unsafe.Pointer(uintptr(v0 + 36))) = *(*uint32)(unsafe.Pointer(uintptr(v0 + 32)))
				}
			} else {
			LABEL_16:
				sub_511360(v0)
			}
			v0 = v2
			if v2 == 0 {
				break
			}
		}
	}
}
func sub_511560(a1 int32) int32 {
	var result int32
	result = a1
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 48))) = 0
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 44))) = dword_5d4594_2386572
	if dword_5d4594_2386572 != 0 {
		*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_2386572 + 48))) = uint32(a1)
	}
	dword_5d4594_2386572 = uint32(a1)
	return result
}
func nox_xxx_unitSetDecayTime_511660(a1p *nox_object_t, a2 int32) int32 {
	var (
		a1     *uint32 = (*uint32)(unsafe.Pointer(a1p))
		result int32
		v3     uint32
		v4     int32
		v5     int32
		v6     int32
	)
	result = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*4)))
	if (uint32(result) & 0x10000) == 0 {
		if uint32(result)&0x400000 != 0 {
			nox_xxx_decay_5116F0((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1)))))))
		}
		v3 = nox_frame_xxx_2598000 + uint32(a2)
		v4 = 0
		*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*34)) = nox_frame_xxx_2598000 + uint32(a2)
		v5 = int32(dword_5d4594_2386576)
		if dword_5d4594_2386576 == 0 {
			goto LABEL_14
		}
		for {
			if v3 < uint32(*(*int32)(unsafe.Pointer(uintptr(v5 + 136)))) {
				break
			}
			v4 = v5
			v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(v5 + 468))))
			if v5 == 0 {
				break
			}
		}
		if v4 != 0 {
			*(*uint32)(unsafe.Pointer(uintptr(v4 + 468))) = uint32(uintptr(unsafe.Pointer(a1)))
			if v5 == 0 {
				v6 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*4)))
				*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*117)) = 0
				result = int32(uint32(v6) | 0x400000)
				*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*4)) = uint32(result)
				return result
			}
		} else {
		LABEL_14:
			dword_5d4594_2386576 = uint32(uintptr(unsafe.Pointer(a1)))
		}
		*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*117)) = uint32(v5)
		result = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*4)) | 0x400000)
		*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*4)) = uint32(result)
	}
	return result
}
func nox_xxx_decay_5116F0(item *nox_object_t) int32 {
	var (
		result int32
		v2     int32
	)
	result = int32(item.obj_flags)
	if uint32(result)&0x400000 != 0 {
		v2 = 0
		item.obj_flags = uint32(result) & 0xFFBFFFFF
		result = int32(dword_5d4594_2386576)
		if dword_5d4594_2386576 != 0 {
			for unsafe.Pointer(uintptr(result)) != unsafe.Pointer(item) {
				v2 = result
				result = int32(*(*uint32)(unsafe.Pointer(uintptr(result + 468))))
				if result == 0 {
					return result
				}
			}
			if result != 0 {
				if v2 != 0 {
					result = int32(item.field_117)
					*(*uint32)(unsafe.Pointer(uintptr(v2 + 468))) = uint32(result)
				} else {
					dword_5d4594_2386576 = item.field_117
				}
			}
		}
	}
	return result
}
func nox_xxx_decay_511750() {
	var (
		v0 *uint32
		v1 *uint32
		v2 int32
	)
	v0 = *(**uint32)(unsafe.Pointer(&dword_5d4594_2386576))
	if dword_5d4594_2386576 != 0 {
		for {
			v1 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v0), unsafe.Sizeof(uint32(0))*117)))))
			if *(*uint32)(unsafe.Add(unsafe.Pointer(v0), unsafe.Sizeof(uint32(0))*123)) != 0 {
				nox_xxx_decay_5116F0((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v0)))))))
			} else {
				if *(*uint32)(unsafe.Add(unsafe.Pointer(v0), unsafe.Sizeof(uint32(0))*34)) > nox_frame_xxx_2598000 {
					return
				}
				nox_xxx_decay_5116F0((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v0)))))))
				v2 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v0), unsafe.Sizeof(uint32(0))*5)))
				*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v2))), 0)) = uint8(int8(v2 | 128))
				*(*uint32)(unsafe.Add(unsafe.Pointer(v0), unsafe.Sizeof(uint32(0))*5)) = uint32(v2)
				nox_xxx_delayedDeleteObject_4E5CC0((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v0)))))))
			}
			v0 = v1
			if v1 == nil {
				break
			}
		}
	}
}
func nox_xxx_decayDestroy_5117B0() int32 {
	var (
		result int32
		v1     int32
	)
	result = int32(dword_5d4594_2386576)
	if dword_5d4594_2386576 != 0 {
		for {
			v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(result + 468))))
			nox_xxx_decay_5116F0((*nox_object_t)(unsafe.Pointer(uintptr(result))))
			result = v1
			if v1 == 0 {
				break
			}
		}
		dword_5d4594_2386576 = 0
	} else {
		dword_5d4594_2386576 = 0
	}
	return result
}
func sub_5117F0(a1 int32) int8 {
	var result int8
	result = int8(a1)
	if (int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8)))) & 1) == 0 {
		result = nox_xxx_unitHasCollideOrUpdateFn_537610((*nox_object_t)(unsafe.Pointer(uintptr(a1))))
	}
	return result
}
func nox_xxx_unit_511810(a1 int32) {
	if (int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8)))) & 1) == 0 {
		if sub_537580(a1) != 0 {
			sub_5375A0(a1)
		}
	}
}
func nox_xxx_collisions_511850() {
	var (
		v0 int32
		v1 int32
	)
	nox_xxx_allocHitArray_5486D0()
	v0 = 5
	for {
		nox_xxx_updateObjectsVelocity_5118A0(0.2)
		sub_548B60()
		v0--
		if v0 == 0 {
			break
		}
	}
	for sub_537590() == 0 {
		v1 = sub_537700()
		if int32(int8(*(*uint8)(unsafe.Pointer(uintptr(v1 + 8))))) >= 0 {
			nox_xxx_unitAddToUpdatable_4DA8D0((*nox_object_t)(unsafe.Pointer(uintptr(v1))))
		}
	}
	nox_xxx_collide_548740()
}
func nox_xxx_updateObjectsVelocity_5118A0(step float32) int32 {
	var (
		i      int32
		result int32
		j      int32
		v4     float64
		v5     float32
		v6     *float32
		v7     *uint8
		v8     int8
		v9     int32
		v10    float64
		v11    float64
		v12    float32
		v13    float32
		v15    float2
		v16    float4
	)
	if *memmap.PtrUint32(6112660, 2386580) == 0 {
		*memmap.PtrUint32(6112660, 2386580) = uint32(nox_xxx_getNameId_4E3AA0(CString("SmallFlameCleanse")))
		*memmap.PtrUint32(6112660, 2386584) = uint32(nox_xxx_getNameId_4E3AA0(CString("SmallFlameCleanse")))
		*memmap.PtrUint32(6112660, 2386588) = uint32(nox_xxx_getNameId_4E3AA0(CString("MediumFlameCleanse")))
		*memmap.PtrUint32(6112660, 2386592) = uint32(nox_xxx_getNameId_4E3AA0(CString("FlameCleanse")))
		*memmap.PtrUint32(6112660, 2386596) = uint32(nox_xxx_getNameId_4E3AA0(CString("LargeFlameCleanse")))
		*memmap.PtrUint32(6112660, 2386600) = uint32(nox_xxx_getNameId_4E3AA0(CString("SmallBlueFlameCleanse")))
		*memmap.PtrUint32(6112660, 2386604) = uint32(nox_xxx_getNameId_4E3AA0(CString("SmallBlueFlameCleanse")))
		*memmap.PtrUint32(6112660, 2386608) = uint32(nox_xxx_getNameId_4E3AA0(CString("MediumBlueFlameCleanse")))
		*memmap.PtrUint32(6112660, 2386612) = uint32(nox_xxx_getNameId_4E3AA0(CString("BlueFlameCleanse")))
		*memmap.PtrUint32(6112660, 2386616) = uint32(nox_xxx_getNameId_4E3AA0(CString("LargeBlueFlameCleanse")))
	}
	for i = sub_537740(); i != 0; i = sub_537750(i) {
		sub_5481C0(i)
	}
	nox_xxx_updateDebugObjects_5113A0()
	result = sub_537740()
	for j = result; result != 0; j = result {
		if int32(*(*uint8)(unsafe.Pointer(uintptr(j + 16))))&2 != 0 || int32(*(*uint8)(unsafe.Pointer(uintptr(j + 8))))&2 != 0 && nox_xxx_checkMobAction_50A0D0((*nox_object_t)(unsafe.Pointer(uintptr(j))), 67) != 0 {
			*(*uint32)(unsafe.Pointer(uintptr(j + 100))) = 0
			*(*uint32)(unsafe.Pointer(uintptr(j + 96))) = 0
			*(*uint32)(unsafe.Pointer(uintptr(j + 84))) = 0
			*(*uint32)(unsafe.Pointer(uintptr(j + 80))) = 0
		} else {
			if nox_xxx_testUnitBuffs_4FF350((*nox_object_t)(unsafe.Pointer(uintptr(j))), 5) != 0 || nox_xxx_testUnitBuffs_4FF350((*nox_object_t)(unsafe.Pointer(uintptr(j))), 25) != 0 || nox_xxx_testUnitBuffs_4FF350((*nox_object_t)(unsafe.Pointer(uintptr(j))), 28) != 0 {
				v4 = float64(*(*float32)(unsafe.Pointer(uintptr(j + 96))))
				v13 = *(*float32)(unsafe.Pointer(uintptr(j + 100)))
			} else {
				v4 = float64(*(*float32)(unsafe.Pointer(uintptr(j + 96))) + *(*float32)(unsafe.Pointer(uintptr(j + 88))))
				v13 = *(*float32)(unsafe.Pointer(uintptr(j + 100))) + *(*float32)(unsafe.Pointer(uintptr(j + 92)))
			}
			v5 = *(*float32)(unsafe.Pointer(uintptr(j + 64)))
			v6 = (*float32)(unsafe.Pointer(uintptr(j + 64)))
			*(*float32)(unsafe.Pointer(uintptr(j + 80))) += float32((v4 - float64(*(*float32)(unsafe.Pointer(uintptr(j + 80)))**(*float32)(unsafe.Pointer(uintptr(j + 112))))) * float64(step))
			*(*float32)(unsafe.Pointer(uintptr(j + 84))) += (v13 - *(*float32)(unsafe.Pointer(uintptr(j + 84)))**(*float32)(unsafe.Pointer(uintptr(j + 112)))) * step
			v16.field_0 = v5
			v16.field_4 = *(*float32)(unsafe.Pointer(uintptr(j + 68)))
			v7 = (*uint8)(memmap.PtrOff(6112660, 2386580))
			v16.field_8 = step**(*float32)(unsafe.Pointer(uintptr(j + 80))) + *(*float32)(unsafe.Pointer(uintptr(j + 64)))
			v16.field_C = step**(*float32)(unsafe.Pointer(uintptr(j + 84))) + *(*float32)(unsafe.Pointer(uintptr(j + 68)))
			v8 = int8(uint8((*(*uint32)(unsafe.Pointer(uintptr(j + 16)))>>12)&4 | 1))
			for uint32(*(*uint16)(unsafe.Pointer(uintptr(j + 4)))) != *(*uint32)(unsafe.Pointer(v7)) {
				v7 = (*uint8)(unsafe.Add(unsafe.Pointer(v7), 4))
				if int32(uintptr(unsafe.Pointer(v7))) >= int32(uintptr(memmap.PtrOff(6112660, 2386620))) {
					goto LABEL_20
				}
			}
			v8 = int8(uint8((*(*uint32)(unsafe.Pointer(uintptr(j + 16)))>>12)&4 | 65))
		LABEL_20:
			if nox_xxx_mapTraceRay_535250(&v16, nil, nil, v8) != 0 {
				*(*float32)(unsafe.Pointer(uintptr(j + 64))) = v16.field_8
				*(*float32)(unsafe.Pointer(uintptr(j + 68))) = v16.field_C
			}
			v9 = int32(*(*uint32)(unsafe.Pointer(uintptr(j + 16))))
			if (v9&0x4000) == 0 && *(*uint32)(unsafe.Pointer(uintptr(j + 556))) != 0 && nox_xxx_tileNFromPoint_411160((*float2)(unsafe.Pointer(uintptr(j+64)))) == 6 {
				v15.field_0 = 0.0
				v15.field_4 = 0.0
				nox_xxx_collSysAddCollision_548630(j, 6, &v15)
			}
			v10 = float64(*v6 - *(*float32)(unsafe.Pointer(uintptr(j + 56))))
			if v10 < 0.0 {
				v10 = -v10
			}
			v11 = float64(*(*float32)(unsafe.Pointer(uintptr(j + 68))) - *(*float32)(unsafe.Pointer(uintptr(j + 60))))
			if v11 < 0.0 {
				v11 = -v11
			}
			if v10 > 0.0099999998 || (func() bool {
				v12 = float32(v11)
				return float64(v12) > 0.0099999998
			}()) {
				nox_xxx_unitNeedSync_4E44F0(j)
				nox_xxx_objectUnkUpdateCoords_4E7290(j)
				nox_xxx_moveUpdateSpecial_517970((*nox_object_t)(unsafe.Pointer(uintptr(j))))
			}
		}
		result = sub_537750(j)
	}
	return result
}
func nox_xxx_script_511C50(a1 int32) *nox_object_t {
	var v1 *uint32
	if dword_587000_237036 != 0 {
		sub_511D20()
	}
	v1 = *(**uint32)(memmap.PtrOff(6112660, 2386820))
	if *memmap.PtrUint32(6112660, 2386820) == 0 {
		return nil
	}
	for int32(*(*uint8)(unsafe.Pointer(uintptr(*v1 + 16))))&32 != 0 || *(*uint32)(unsafe.Pointer(uintptr(*v1 + 44))) != uint32(a1) {
		v1 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*2)))))
		if v1 == nil {
			return nil
		}
	}
	sub_511CE0((*uint32)(memmap.PtrOff(6112660, 2386820)), int32(uintptr(unsafe.Pointer(v1))))
	sub_511CB0((*uint32)(memmap.PtrOff(6112660, 2386820)), int32(uintptr(unsafe.Pointer(v1))))
	return (*nox_object_t)(unsafe.Pointer(uintptr(*v1)))
}
func sub_511CB0(a1 *uint32, a2 int32) int32 {
	var result int32
	result = a2
	*(*uint32)(unsafe.Pointer(uintptr(a2 + 4))) = 0
	*(*uint32)(unsafe.Pointer(uintptr(a2 + 8))) = *a1
	if *a1 != 0 {
		*(*uint32)(unsafe.Pointer(uintptr(*a1 + 4))) = uint32(a2)
	} else {
		*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*1)) = uint32(a2)
	}
	*a1 = uint32(a2)
	return result
}
func sub_511CE0(a1 *uint32, a2 int32) int32 {
	var (
		result int32
		v3     int32
		v4     int32
	)
	result = a2
	v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 8))))
	if v3 != 0 {
		*(*uint32)(unsafe.Pointer(uintptr(v3 + 4))) = *(*uint32)(unsafe.Pointer(uintptr(a2 + 4)))
	} else {
		*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*1)) = *(*uint32)(unsafe.Pointer(uintptr(a2 + 4)))
	}
	v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 4))))
	if v4 != 0 {
		*(*uint32)(unsafe.Pointer(uintptr(v4 + 8))) = *(*uint32)(unsafe.Pointer(uintptr(a2 + 8)))
	} else {
		result = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 8))))
		*a1 = uint32(result)
	}
	return result
}
func sub_511D20() int32 {
	var (
		v0     *uint8
		result int32
	)
	v0 = (*uint8)(memmap.PtrOff(6112660, 2386628))
	*memmap.PtrUint32(6112660, 2386820) = 0
	*memmap.PtrUint32(6112660, 2386824) = 0
	*memmap.PtrUint32(6112660, 2386620) = 0
	*memmap.PtrUint32(6112660, 2386624) = 0
	for {
		result = sub_511CB0((*uint32)(memmap.PtrOff(6112660, 2386620)), int32(uintptr(unsafe.Pointer(v0))))
		v0 = (*uint8)(unsafe.Add(unsafe.Pointer(v0), 12))
		if int32(uintptr(unsafe.Pointer(v0))) >= int32(uintptr(memmap.PtrOff(6112660, 2386820))) {
			break
		}
	}
	dword_587000_237036 = 0
	return result
}
func nox_xxx_scriptPrepareFoundUnit_511D70(obj *nox_object_t) int32 {
	var (
		a1     int32 = int32(uintptr(unsafe.Pointer(obj)))
		v1     *int32
		v2     int32
		result int32
		v4     int32
	)
	v1 = (*int32)(unsafe.Pointer(uintptr(sub_511DC0())))
	if v1 != nil {
		*v1 = a1
		result = sub_511CB0((*uint32)(memmap.PtrOff(6112660, 2386820)), int32(uintptr(unsafe.Pointer(v1))))
	} else {
		v2 = int32(*memmap.PtrUint32(6112660, 2386824))
		v4 = int32(*memmap.PtrUint32(6112660, 2386824))
		**(**uint32)(memmap.PtrOff(6112660, 2386824)) = uint32(a1)
		sub_511CE0((*uint32)(memmap.PtrOff(6112660, 2386820)), v4)
		result = sub_511CB0((*uint32)(memmap.PtrOff(6112660, 2386820)), v2)
	}
	return result
}
func sub_511DC0() int32 {
	var result int32
	result = int32(*memmap.PtrUint32(6112660, 2386620))
	if *memmap.PtrUint32(6112660, 2386620) == 0 {
		return 0
	}
	*memmap.PtrUint32(6112660, 2386620) = *(*uint32)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, 2386620) + 8)))
	return result
}
func sub_511DE0(a1 int32) int32 {
	var (
		result int32
		v2     *uint32
	)
	result = int32(dword_587000_237036)
	if dword_587000_237036 == 0 {
		v2 = *(**uint32)(memmap.PtrOff(6112660, 2386820))
		if *memmap.PtrUint32(6112660, 2386820) != 0 {
			result = a1
			for *v2 != uint32(a1) {
				v2 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*2)))))
				if v2 == nil {
					return result
				}
			}
			sub_511CE0((*uint32)(memmap.PtrOff(6112660, 2386820)), int32(uintptr(unsafe.Pointer(v2))))
			result = sub_511CB0((*uint32)(memmap.PtrOff(6112660, 2386620)), int32(uintptr(unsafe.Pointer(v2))))
		}
	}
	return result
}
func sub_511E20() int32 {
	var (
		result int32
		v1     int32
		v2     int32
	)
	result = int32(dword_587000_237036)
	if dword_587000_237036 == 0 {
		v1 = int32(*memmap.PtrUint32(6112660, 2386820))
		if *memmap.PtrUint32(6112660, 2386820) != 0 {
			for {
				v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 8))))
				sub_511CE0((*uint32)(memmap.PtrOff(6112660, 2386820)), v1)
				result = sub_511CB0((*uint32)(memmap.PtrOff(6112660, 2386620)), v1)
				v1 = v2
				if v2 == 0 {
					break
				}
			}
		}
	}
	return result
}
func nox_xxx_wallOpen_511F80(a1 int32) {
	var (
		v1 int32
		v2 int32
		v3 int8
		v4 int32
		v5 *byte
		v6 int32
		v7 float2
		v8 int32
	)
	v1 = a1
	if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 4))))&4 != 0 {
		v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 28))))
		v3 = int8(*(*uint8)(unsafe.Pointer(uintptr(v2 + 21))))
		if int32(v3) != 3 && int32(v3) != 4 {
			v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 4))))
			*(*uint8)(unsafe.Pointer(uintptr(v2 + 21))) = 4
			v8 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 8))) * 23)
			v7.field_0 = float32(float64(v4*23) + 11.5)
			v7.field_4 = float32(float64(v8) + 11.5)
			if nox_xxx_wallSounds_2386840 == 0 {
				v5 = nox_xxx_wallFindOpenSound_410EE0(int32(*(*uint8)(unsafe.Pointer(uintptr(v1 + 1)))))
				v6 = nox_xxx_utilFindSound_40AF50(v5)
				nox_xxx_audCreate_501A30(v6, &v7, 0, 0)
			}
		}
	}
}
func nox_xxx_wallClose_512070(a1 int32) {
	var (
		v1 int32
		v2 int32
		v3 int8
		v4 int32
		v5 *byte
		v6 int32
		v7 float2
		v8 int32
	)
	v1 = a1
	if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 4))))&4 != 0 {
		v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 28))))
		v3 = int8(*(*uint8)(unsafe.Pointer(uintptr(v2 + 21))))
		if int32(v3) != 1 && int32(v3) != 2 {
			v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 4))))
			*(*uint8)(unsafe.Pointer(uintptr(v2 + 21))) = 2
			v8 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 8))) * 23)
			v7.field_0 = float32(float64(v4*23) + 11.5)
			v7.field_4 = float32(float64(v8) + 11.5)
			if nox_xxx_wallSounds_2386840 == 0 {
				v5 = nox_xxx_wallFindCloseSound_410F20(int32(*(*uint8)(unsafe.Pointer(uintptr(v1 + 1)))))
				v6 = nox_xxx_utilFindSound_40AF50(v5)
				nox_xxx_audCreate_501A30(v6, &v7, 0, 0)
			}
		}
	}
}
func nox_xxx_wallToggle_512160(a1 int32) {
	var (
		v1 int32
		v2 int8
		v3 int32
		v4 *byte
		v5 int32
		v6 int32
		v7 float2
	)
	if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 4))))&4 != 0 {
		v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 28))))
		v2 = int8(*(*uint8)(unsafe.Pointer(uintptr(v1 + 21))))
		if int32(v2) == 1 || int32(v2) == 2 {
			v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 4))))
			*(*uint8)(unsafe.Pointer(uintptr(v1 + 21))) = 4
			v7.field_0 = float32(float64(v5*23) + 11.5)
			v7.field_4 = float32(float64(int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 8)))*23)) + 11.5)
			if nox_xxx_wallSounds_2386840 != 0 {
				return
			}
			v4 = nox_xxx_wallFindOpenSound_410EE0(int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 1)))))
		} else {
			v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 4))))
			*(*uint8)(unsafe.Pointer(uintptr(v1 + 21))) = 2
			v7.field_0 = float32(float64(v3*23) + 11.5)
			v7.field_4 = float32(float64(int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 8)))*23)) + 11.5)
			if nox_xxx_wallSounds_2386840 != 0 {
				return
			}
			v4 = nox_xxx_wallFindCloseSound_410F20(int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 1)))))
		}
		v6 = nox_xxx_utilFindSound_40AF50(v4)
		nox_xxx_audCreate_501A30(v6, &v7, 0, 0)
	}
}
func nox_xxx_wallPreDestroyByPtr_5122C0(a1 int32) int32 {
	var (
		v1 int32
		v3 int2
	)
	v1 = int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 6))))
	v3.field_0 = int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 5))))
	v3.field_4 = v1
	return nox_xxx_wallPreDestroy_534DA0(&v3.field_0)
}
func nox_xxx_monsterLookAt_5125A0(obj *nox_object_t, a2 int32) *float32 {
	var (
		a1     int32 = int32(uintptr(unsafe.Pointer(obj)))
		result *float32
		v3     int32
		v4     float32
		v5     float32
	)
	result = (*float32)(unsafe.Pointer(uintptr(nox_xxx_mathDirection4ToAngle_509E90(a2))))
	if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))&2 != 0 {
		v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 16))))
		if (v3 & 0x8000) == 0 {
			v4 = float32(float64(*mem_getFloatPtr(0x587000, uint32(uintptr(unsafe.Pointer(result)))*8+0x2F658))*10.0 + float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 56)))))
			v5 = float32(float64(*mem_getFloatPtr(0x587000, uint32(uintptr(unsafe.Pointer(result)))*8+194140))*10.0 + float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 60)))))
			result = (*float32)(unsafe.Pointer(nox_xxx_monsterPushAction_50A260_impl((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 25, CString(__FILE__), __LINE__)))
			if result != nil {
				*(*float32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(float32(0))*1)) = v4
				*(*float32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(float32(0))*2)) = v5
			}
		}
	}
	return result
}
func sub_5126C0(a1 int32) int8 {
	return nox_xxx_objectSetOn_4E75B0((*nox_object_t)(unsafe.Pointer(uintptr(a1))))
}
func sub_512720(a1 int32) int32 {
	var result int32
	result = a1
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 480))) |= 1
	return result
}
func sub_512780(a1 int32) int32 {
	return nox_xxx_objectSetOff_4E7600((*nox_object_t)(unsafe.Pointer(uintptr(a1))))
}
func sub_5127E0(a1 int32) int32 {
	var result int32
	result = a1
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 480))) &= 0xFFFFFFFE
	return result
}
func sub_512840(a1 int32) int8 {
	return nox_xxx_objectToggle_4E7650(a1)
}
func sub_5128A0(a1 int32) int32 {
	var result int32
	result = a1
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 480))) ^= 1
	return result
}
func sub_512900(a1 int32) {
	nox_xxx_delayedDeleteObject_4E5CC0((*nox_object_t)(unsafe.Pointer(uintptr(a1))))
}
func sub_512FE0(a1 int32, a2 *int32) {
	var (
		v2 int32
		v3 int32
		v4 int32
		v5 int32
	)
	if *(*int32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(int32(0))*1)) > 0 {
		v5 = *(*int32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(int32(0))*2))
		v4 = *(*int32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(int32(0))*1))
		v3 = *a2
		v2 = int32(uintptr(unsafe.Pointer(nox_xxx_findParentChainPlayer_4EC580((*nox_object_t)(unsafe.Pointer(uintptr(*a2)))))))
		(cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(a1 + 716))), (*func(int32, int32, int32, int32, int32))(nil)))(a1, v2, v3, v4, v5)
	}
}
func sub_5130E0(a1 int32, a2 *uint32) *uint32 {
	var (
		v2     *uint32
		v3     *uint32
		result *uint32
		v5     int32
	)
	v2 = (*uint32)(unsafe.Pointer(nox_xxx_newObjectByTypeID_4E3810(CString("Mover"))))
	v3 = v2
	if v2 != nil {
		nox_xxx_createAt_4DAA50((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v2)))))), nil, *(*float32)(unsafe.Pointer(uintptr(a1 + 56))), *(*float32)(unsafe.Pointer(uintptr(a1 + 60))))
		v5 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*187)))
		*(*uint32)(unsafe.Pointer(uintptr(v5 + 28))) = uint32(a1)
		*(*uint32)(unsafe.Pointer(uintptr(v5 + 8))) = *(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*1))
		*(*uint32)(unsafe.Pointer(uintptr(v5 + 4))) = *a2
		*(*uint8)(unsafe.Pointer(uintptr(v5))) = 0
		*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*20)) = 0
		*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*21)) = 0
		nox_xxx_objectSetOn_4E75B0((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v3)))))))
		nox_xxx_unitAddToUpdatable_4DA8D0((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v3)))))))
		result = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*11)))))
		*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*2)) = uint32(uintptr(unsafe.Pointer(result)))
	} else {
		result = a2
		*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*2)) = 0
	}
	return result
}
func sub_513280(a1 int32, a2 *int32) int32 {
	var v2 int32
	v2 = 0
	if noxflags.HasGame(noxflags.GameModeCoop) && int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))&4 != 0 && *(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))) + 276))) + uint32(*a2*4) + 3696))) == 0 {
		v2 = 1
	}
	return nox_xxx_spellGrantToPlayer_4FB550((*nox_object_t)(unsafe.Pointer(uintptr(a1))), *a2, 1, v2, 0)
}
func nox_xxx_enchantUnit_513390(a1 int32, a2 *int32) int16 {
	nox_xxx_buffApplyTo_4FF380((*nox_object_t)(unsafe.Pointer(uintptr(a1))), *a2, int16(*(*int32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(int32(0))*1))), 5)
	return 0
}
func nox_xxx_monsterWalkTo_514110(obj *nox_object_t, x float32, y float32) {
	var (
		a1     int32 = int32(uintptr(unsafe.Pointer(obj)))
		result *int32
		v4     *int32
	)
	result = *(**int32)(unsafe.Pointer(uintptr(a1 + 16)))
	if int32(*(*int8)(unsafe.Add(unsafe.Pointer((*int8)(unsafe.Pointer(&result))), 1))) >= 0 && int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))&2 != 0 {
		nox_xxx_monsterClearActionStack_50A3A0(a1)
		v4 = nox_xxx_monsterPushAction_50A260_impl((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 32, CString(__FILE__), __LINE__)
		if v4 != nil {
			*(*int32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(int32(0))*1)) = 8
		}
		result = nox_xxx_monsterPushAction_50A260_impl((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 8, CString(__FILE__), __LINE__)
		if result != nil {
			*(*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(result))), unsafe.Sizeof(float32(0))*1)) = x
			*(*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(result))), unsafe.Sizeof(float32(0))*2)) = y
			*(*int32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(int32(0))*3)) = 0
		}
	}
}
func nox_xxx_monsterGoPatrol_515680(a1 int32, a2 int32) {
	var (
		v2 int32
		v3 int32
		v4 *int32
		v5 float2
	)
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
	if a1 != 0 {
		if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))&2 != 0 {
			v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 16))))
			if (v3 & 0x8000) == 0 {
				v5.field_0 = *(*float32)(unsafe.Pointer(uintptr(a2 + 8))) - *(*float32)(unsafe.Pointer(uintptr(a2)))
				v5.field_4 = *(*float32)(unsafe.Pointer(uintptr(a2 + 12))) - *(*float32)(unsafe.Pointer(uintptr(a2 + 4)))
				nox_xxx_monsterClearActionStack_50A3A0(a1)
				v4 = nox_xxx_monsterPushAction_50A260_impl((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 4, CString(__FILE__), __LINE__)
				if v4 != nil {
					*(*int32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(int32(0))*1)) = int32(*(*uint32)(unsafe.Pointer(uintptr(a2))))
					*(*int32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(int32(0))*2)) = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 4))))
					*(*int32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(int32(0))*3)) = nox_xxx_math_509ED0(&v5)
				}
				*(*uint32)(unsafe.Pointer(uintptr(v2 + 1312))) = *(*uint32)(unsafe.Pointer(uintptr(a2 + 16)))
			}
		}
	}
}
func nox_xxx_unitHunt_5157A0(obj *nox_object_t) {
	var (
		a1     int32 = int32(uintptr(unsafe.Pointer(obj)))
		result int32
	)
	if a1 != 0 && int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))&2 != 0 {
		result = *(*int32)(unsafe.Pointer(uintptr(a1 + 16)))
		if int32(*(*int8)(unsafe.Add(unsafe.Pointer((*int8)(unsafe.Pointer(&result))), 1))) >= 0 {
			nox_xxx_monsterClearActionStack_50A3A0(a1)
			nox_xxx_monsterPushAction_50A260_impl((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 5, CString(__FILE__), __LINE__)
		}
	}
}
func nox_xxx_unitIdle_515820(obj *nox_object_t) {
	var (
		a1     int32 = int32(uintptr(unsafe.Pointer(obj)))
		result int32
	)
	if a1 != 0 && int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))&2 != 0 {
		result = *(*int32)(unsafe.Pointer(uintptr(a1 + 16)))
		if int32(*(*int8)(unsafe.Add(unsafe.Pointer((*int8)(unsafe.Pointer(&result))), 1))) >= 0 {
			nox_xxx_monsterClearActionStack_50A3A0(a1)
			nox_xxx_monsterPushAction_50A260_impl((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, CString(__FILE__), __LINE__)
		}
	}
}
func nox_xxx_unitSetFollow_5158C0(obj1 *nox_object_t, obj2 *nox_object_t) {
	var (
		a1 int32 = int32(uintptr(unsafe.Pointer(obj1)))
		a2 int32 = int32(uintptr(unsafe.Pointer(obj2)))
		v2 int32
		v3 *int32
	)
	if a1 != 0 {
		if a2 != 0 {
			if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))&2 != 0 {
				if a1 != a2 {
					v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 16))))
					if (v2 & 0x8000) == 0 {
						nox_xxx_monsterClearActionStack_50A3A0(a1)
						v3 = nox_xxx_monsterPushAction_50A260_impl((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 3, CString(__FILE__), __LINE__)
						if v3 != nil {
							*(*int32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(int32(0))*1)) = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 56))))
							*(*int32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(int32(0))*2)) = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 60))))
							*(*int32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(int32(0))*3)) = a2
						}
					}
				}
			}
		}
	}
}
func nox_xxx_monsterSetAgressiveness_515980(a1 int32, a2 *uint32) int32 {
	var result int32
	result = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
	if a1 != 0 {
		if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))&2 != 0 {
			*(*uint32)(unsafe.Pointer(uintptr(result + 1308))) = *a2
			*(*uint32)(unsafe.Pointer(uintptr(result + 1304))) = *a2
		}
	}
	return result
}
func nox_xxx_monsterActionMelee_515A30(a1 int32, a2 *float2) {
	var (
		v2 int32
		v3 *int32
		v4 *float32
		v5 *float32
	)
	if a1 != 0 {
		if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))&2 != 0 {
			v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 16))))
			if (v2 & 0x8000) == 0 {
				if nox_xxx_monsterCanMelee_534220(a1) != 0 {
					nox_xxx_monsterClearActionStack_50A3A0(a1)
					v3 = nox_xxx_monsterPushAction_50A260_impl((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 32, CString(__FILE__), __LINE__)
					if v3 != nil {
						*(*int32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(int32(0))*1)) = 16
					}
					nox_xxx_monsterPushAction_50A260_impl((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 16, CString(__FILE__), __LINE__)
					v4 = (*float32)(unsafe.Pointer(nox_xxx_monsterPushAction_50A260_impl((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 51, CString(__FILE__), __LINE__)))
					if v4 != nil {
						*(*float32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(float32(0))*1)) = float32(sub_534470(a1) + float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 176)))))
						*(*float32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(float32(0))*3)) = a2.field_0
						*(*float32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(float32(0))*4)) = a2.field_4
					}
					v5 = (*float32)(unsafe.Pointer(nox_xxx_monsterPushAction_50A260_impl((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 7, CString(__FILE__), __LINE__)))
					if v5 != nil {
						*(*float32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(float32(0))*1)) = a2.field_0
						*(*float32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(float32(0))*2)) = a2.field_4
						*(*float32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(float32(0))*3)) = 0
					}
				}
			}
		}
	}
}
func nox_xxx_monsterMissileAttack_515B80(a1 int32, a2 *uint32) {
	var (
		v2 int32
		v3 *int32
		v4 *int32
	)
	if a1 != 0 {
		if a2 != nil {
			if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))&2 != 0 {
				v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 16))))
				if (v2 & 0x8000) == 0 {
					if nox_xxx_monsterCanShoot_534280(a1) != 0 {
						nox_xxx_monsterClearActionStack_50A3A0(a1)
						v3 = nox_xxx_monsterPushAction_50A260_impl((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 32, CString(__FILE__), __LINE__)
						if v3 != nil {
							*(*int32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(int32(0))*1)) = 17
						}
						v4 = nox_xxx_monsterPushAction_50A260_impl((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 17, CString(__FILE__), __LINE__)
						if v4 != nil {
							*(*int32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(int32(0))*1)) = int32(*a2)
							*(*int32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(int32(0))*2)) = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*1)))
							*(*int32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(int32(0))*3)) = 0
						}
					}
				}
			}
		}
	}
}
func sub_515C80(a1 int32, a2 *uint8) int32 {
	var result int32
	result = a1
	if a1 != 0 {
		if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))&2 != 0 {
			result = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
			*(*uint8)(unsafe.Pointer(uintptr(result + 1332))) = *a2
		}
	}
	return result
}
func nox_xxx_mobSetFightTarg_515D30(a1 int32, a2 int32) {
	var (
		v2 int32
		v3 int32
		v4 *int32
		v5 *int32
	)
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
	if a1 != 0 {
		if a2 != 0 {
			if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))&2 != 0 {
				if a1 != a2 {
					v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 16))))
					if (v3 & 0x8000) == 0 {
						nox_xxx_monsterClearActionStack_50A3A0(a1)
						*(*uint32)(unsafe.Pointer(uintptr(v2 + 1216))) = uint32(a2)
						nox_xxx_frameCounterSetCopyToNextFrame_5281D0()
						v4 = nox_xxx_monsterPushAction_50A260_impl((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 32, CString(__FILE__), __LINE__)
						if v4 != nil {
							*(*int32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(int32(0))*1)) = 15
						}
						v5 = nox_xxx_monsterPushAction_50A260_impl((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 15, CString(__FILE__), __LINE__)
						if v5 != nil {
							*(*int32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(int32(0))*1)) = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 56))))
							*(*int32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(int32(0))*2)) = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 60))))
							*(*int32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(int32(0))*3)) = int32(nox_frame_xxx_2598000)
						}
					}
				}
			}
		}
	}
}
func sub_515E20(a1 int32, a2 *uint32) int32 {
	var result int32
	result = a1
	if a1 != 0 && a2 != nil {
		if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))&2 != 0 {
			result = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
			*(*uint32)(unsafe.Pointer(uintptr(result + 1336))) = *a2
		}
	}
	return result
}
func sub_515EB0(a1 int32, a2 *uint32) int32 {
	var result int32
	result = a1
	if a1 != 0 && a2 != nil {
		if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))&2 != 0 {
			result = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
			*(*uint32)(unsafe.Pointer(uintptr(result + 1344))) = *a2
		}
	}
	return result
}
func nox_server_scriptFleeFrom_515F70(a1 int32, a2 *uint32) {
	var (
		v2 int32
		v3 *int32
		v4 *int32
		v5 *int32
		v6 int32
		v7 int32
	)
	if a1 != 0 {
		if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))&2 != 0 {
			v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 16))))
			if (v2&0x8000) == 0 && nox_xxx_checkMobAction_50A0D0((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 24) == 0 {
				v3 = nox_xxx_monsterPushAction_50A260_impl((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 32, CString(__FILE__), __LINE__)
				if v3 != nil {
					*(*int32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(int32(0))*1)) = 24
				}
				v4 = nox_xxx_monsterPushAction_50A260_impl((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 41, CString(__FILE__), __LINE__)
				if v4 != nil {
					*(*int32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(int32(0))*1)) = int32(nox_frame_xxx_2598000 + *(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*1)))
				}
				v5 = nox_xxx_monsterPushAction_50A260_impl((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 24, CString(__FILE__), __LINE__)
				if v5 != nil {
					v6 = int32(*a2)
					*(*int32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(int32(0))*1)) = int32(*(*uint32)(unsafe.Pointer(uintptr(*a2 + 56))))
					v7 = int32(*(*uint32)(unsafe.Pointer(uintptr(v6 + 60))))
					*(*int32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(int32(0))*3)) = 0
					*(*int32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(int32(0))*2)) = v7
				}
			}
		}
	}
}
func sub_516090(a1 int32, a2 *uint32) {
	var (
		v2 int32
		v3 *int32
		v4 *int32
	)
	if a1 != 0 {
		if a2 != nil {
			if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))&2 != 0 {
				v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 16))))
				if (v2 & 0x8000) == 0 {
					v3 = nox_xxx_monsterPushAction_50A260_impl((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 32, CString(__FILE__), __LINE__)
					if v3 != nil {
						*(*int32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(int32(0))*1)) = 1
					}
					v4 = nox_xxx_monsterPushAction_50A260_impl((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 1, CString(__FILE__), __LINE__)
					if v4 != nil {
						*(*int32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(int32(0))*1)) = int32(nox_frame_xxx_2598000 + *a2)
					}
				}
			}
		}
	}
}
func sub_516570(this *uint8) int32 {
	var (
		v1 int32
		v2 *uint8
		v4 *uint8
	)
	v4 = this
	v1 = int32(uintptr(unsafe.Pointer(nox_xxx_getFirstPlayerUnit_4DA7C0())))
	if v1 != 0 {
		for {
			v2 = *(**uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v1 + 748))) + 276)))
			if int32(*(*uint8)(unsafe.Add(unsafe.Pointer(v2), 2064))) == 31 {
				break
			}
			v1 = int32(uintptr(unsafe.Pointer(nox_xxx_getNextPlayerUnit_4DA7F0((*nox_object_t)(unsafe.Pointer(uintptr(v1)))))))
			if v1 == 0 {
				break
			}
		}
	} else {
		v2 = v4
	}
	nox_gameDisableMapDraw_5d4594_2650672 = 1
	return nox_xxx_netSendChapterEnd_4D9560(int32(*(*uint8)(unsafe.Add(unsafe.Pointer(v2), 2064))), int8(*memmap.PtrUint8(6112660, 2386828)), *memmap.PtrInt32(6112660, 2386832))
}
func sub_516A80(a1 int32, a2 int32, a3 int32, a4 int32) {
	var (
		v4 int32
		v5 *uint32
		v6 int32
	)
	if *memmap.PtrUint32(6112660, 2386904) == 0 {
		*memmap.PtrUint32(6112660, 2386904) = uint32(nox_xxx_getNameId_4E3AA0(CString("Glyph")))
	}
	v4 = a1.FirstItem()
	if v4 != 0 {
		for uint32(*(*uint16)(unsafe.Pointer(uintptr(v4 + 4)))) != *memmap.PtrUint32(6112660, 2386904) {
			v4 = nox_xxx_inventoryGetNext_4E7990(v4)
			if v4 == 0 {
				goto LABEL_8
			}
		}
		nox_xxx_delayedDeleteObject_4E5CC0((*nox_object_t)(unsafe.Pointer(uintptr(v4))))
	}
LABEL_8:
	v5 = (*uint32)(unsafe.Pointer(nox_xxx_newObjectByTypeID_4E3810(CString("Glyph"))))
	if v5 != nil {
		v6 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(uint32(0))*173)))
		*(*uint32)(unsafe.Pointer(uintptr(v6))) = uint32(a2)
		*(*uint32)(unsafe.Pointer(uintptr(v6 + 4))) = uint32(a3)
		*(*uint32)(unsafe.Pointer(uintptr(v6 + 8))) = uint32(a4)
		*(*uint8)(unsafe.Pointer(uintptr(v6 + 20))) = 0
		if a2 != 0 {
			*(*uint8)(unsafe.Pointer(uintptr(v6 + 20))) = 1
		}
		if a3 != 0 {
			*(*uint8)(unsafe.Pointer(uintptr(v6 + 20)))++
		}
		if a4 != 0 {
			*(*uint8)(unsafe.Pointer(uintptr(v6 + 20)))++
		}
		*(*uint32)(unsafe.Pointer(uintptr(v6 + 24))) = 0
		*(*uint32)(unsafe.Pointer(uintptr(v6 + 28))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 56)))
		*(*uint32)(unsafe.Pointer(uintptr(v6 + 32))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 60)))
		nox_xxx_inventoryPutImpl_4F3070(a1, (*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v5)))))), 1)
	}
}
func nox_xxx_zombieSetStayDead_516C90(a1 int32) int32 {
	var result int32
	result = a1
	if a1 != 0 {
		if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))&2 != 0 {
			*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))) + 1440))) |= 0x100000
		}
	}
	return result
}
func sub_516D00(a1 int32) uint32 {
	var (
		result uint32
		v2     int32
	)
	result = uint32(a1)
	if a1 != 0 && int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))&2 != 0 {
		v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 16))))
		if (v2 & 0x8000) != 0 {
			*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))) + 1440))) &= 0xFFEFFFFF
			result = nox_xxx_mobRaiseZombie_534AB0(a1)
		}
	}
	return result
}
func nox_xxx_allocPendingOwnsArray_516EE0() int32 {
	dword_5d4594_2386920 = 0
	nox_alloc_pendingOwn_2386916 = unsafe.Pointer(nox_new_alloc_class(CString("PendingOwn"), 12, 512))
	return bool2int(nox_alloc_pendingOwn_2386916 != nil)
}
func sub_516F10() int32 {
	var result int32
	nox_free_alloc_class((*nox_alloc_class)(nox_alloc_pendingOwn_2386916))
	result = 0
	nox_alloc_pendingOwn_2386916 = nil
	dword_5d4594_2386920 = 0
	return result
}
func sub_516F30() {
	nox_alloc_class_free_all((*nox_alloc_class)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&nox_alloc_pendingOwn_2386916)))))
	dword_5d4594_2386920 = 0
}
func sub_516F90(a1 int32, a2 int32) *uint32 {
	var result *uint32
	result = (*uint32)(nox_alloc_class_new_obj_zero((*nox_alloc_class)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&nox_alloc_pendingOwn_2386916))))))
	if result != nil {
		*result = uint32(a1)
		*(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*1)) = uint32(a2)
		*(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*2)) = dword_5d4594_2386920
		dword_5d4594_2386920 = uint32(uintptr(unsafe.Pointer(result)))
	}
	return result
}
func sub_516FC0() {
	var (
		v0 *int32
		v1 int32
		v2 int32
	)
	v0 = *(**int32)(unsafe.Pointer(&dword_5d4594_2386920))
	if dword_5d4594_2386920 != 0 {
		for {
			v1 = sub_4ECF10(*v0)
			v2 = sub_4ECF10(*(*int32)(unsafe.Add(unsafe.Pointer(v0), unsafe.Sizeof(int32(0))*1)))
			if v1 != 0 && v2 != 0 {
				nox_xxx_unitSetOwner_4EC290((*nox_object_t)(unsafe.Pointer(uintptr(v1))), (*nox_object_t)(unsafe.Pointer(uintptr(v2))))
			}
			v0 = (*int32)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(v0), unsafe.Sizeof(int32(0))*2)))))
			if v0 == nil {
				break
			}
		}
	}
	sub_516F30()
}
func nox_xxx_loadMonsterBin_517010() int32 {
	var (
		result int32
		v1     *File
		v2     [256]byte
	)
	dword_5d4594_2386924 = 0
	result = int32(uintptr(unsafe.Pointer(nox_binfile_open_408CC0(CString("monster.bin"), 0))))
	v1 = (*File)(unsafe.Pointer(uintptr(result)))
	if result != 0 {
		result = nox_binfile_cryptSet_408D40((*File)(unsafe.Pointer(uintptr(result))), 23)
		if result != 0 {
			for nox_xxx_readStr_517090(v1, (*uint8)(unsafe.Pointer(&v2[0]))) != 0 && nox_xxx_servParseMonsterDef_517170(v1, &v2[0]) != 0 {
			}
			nox_binfile_close_408D90(v1)
			result = 1
		}
	}
	return result
}
func nox_xxx_readStr_517090(a1 *File, a2 *uint8) int32 {
	var (
		v2       *uint8
		v3       int32
		v4       int32
		v5       int32
		v6       int32
		CharType [2]uint16
	)
	v2 = a2
	v3 = 0
	*(*uint32)(unsafe.Pointer(&CharType[0])) = 0
	v4 = 1
	for {
		for {
			v5 = v3
			nox_binfile_fread_408E40((*byte)(unsafe.Pointer(&CharType[0])), 1, 1, a1)
			if nox_binfile_lastErr_409370(a1) == -1 {
				return 0
			}
			if *memmap.PtrUint32(0x587000, 1668) <= 1 {
				v3 = int32(*(*uint32)(unsafe.Pointer(&CharType[0])))
				v6 = bool2int(unicode.IsSpace(rune(CharType[0])))
			} else {
				v6 = bool2int(unicode.IsSpace(rune(CharType[0])))
				v3 = int32(*(*uint32)(unsafe.Pointer(&CharType[0])))
			}
			if v6 != 0 {
				break
			}
			v4 = 0
			if v3 != 47 || v5 != 47 {
				*func() *uint8 {
					p := &v2
					x := *p
					*p = (*uint8)(unsafe.Add(unsafe.Pointer(*p), 1))
					return x
				}() = uint8(int8(v3))
			} else {
				sub_517140(a1)
				v2 = a2
				v3 = int32(*(*uint32)(unsafe.Pointer(&CharType[0])))
				v4 = 1
			}
		}
		if v4 == 0 {
			break
		}
	}
	*v2 = 0
	return 1
}
func sub_517140(a1 *File) int32 {
	var (
		v1     *File
		result int32
	)
	v1 = a1
	for {
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a1))), 0)) = 0
		nox_binfile_fread_408E40((*byte)(unsafe.Pointer(&a1)), 1, 1, v1)
		result = nox_binfile_lastErr_409370(v1)
		if !(result != -1 && int32(uint8(uintptr(unsafe.Pointer(a1)))) != 10) {
			break
		}
	}
	return result
}
func nox_xxx_servParseMonsterDef_517170(a1 *File, a2 *byte) int32 {
	var (
		result int32
		v3     *uint32
		v4     *uint8
		v5     int32
		v6     *byte
		v7     int32
		v8     *uint8
		v9     int32
		v10    [256]byte
	)
	result = int32(uintptr(alloc.Calloc(1, 248)))
	v3 = (*uint32)(unsafe.Pointer(uintptr(result)))
	if result != 0 {
		libc.StrCpy((*byte)(unsafe.Pointer(uintptr(result))), a2)
		for {
			if nox_xxx_readStr_517090(a1, (*uint8)(unsafe.Pointer(&v10[0]))) == 0 || libc.StrCaseCmp(CString("END"), &v10[0]) == 0 {
				*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*61)) = dword_5d4594_2386924
				dword_5d4594_2386924 = uint32(uintptr(unsafe.Pointer(v3)))
				return 1
			}
			if noxflags.HasGame(noxflags.GameModeCoop) || noxflags.HasGame(noxflags.GameFlag22) {
				if libc.StrCaseCmp(CString("ARENA"), &v10[0]) != 0 {
					if libc.StrCaseCmp(CString("SOLO"), &v10[0]) != 0 {
						goto LABEL_10
					}
				} else {
					sub_517140(a1)
				}
			} else {
			LABEL_10:
				if !noxflags.HasGame(noxflags.GameOnline) {
					goto LABEL_14
				}
				if libc.StrCaseCmp(CString("SOLO"), &v10[0]) != 0 {
					if libc.StrCaseCmp(CString("ARENA"), &v10[0]) != 0 {
					LABEL_14:
						v4 = (*uint8)(memmap.PtrOff(0x587000, 248192))
						if *memmap.PtrUint32(0x587000, 248192) != 0 {
							for {
								if libc.StrCaseCmp(*(**byte)(unsafe.Pointer(v4)), &v10[0]) == 0 {
									break
								}
								v5 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v4))), unsafe.Sizeof(uint32(0))*3))))
								v4 = (*uint8)(unsafe.Add(unsafe.Pointer(v4), 12))
								if v5 == 0 {
									break
								}
							}
						}
						if *(*uint32)(unsafe.Pointer(v4)) == 0 {
							alloc.Free(unsafe.Pointer(v3))
							return 0
						}
						v6 = (*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(v3))), *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v4))), unsafe.Sizeof(uint32(0))*2)))))
						switch *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v4))), unsafe.Sizeof(uint32(0))*1))) {
						case 0:
							nox_xxx_readStr_517090(a1, (*uint8)(unsafe.Pointer(&v10[0])))
							*(*uint32)(unsafe.Pointer(v6)) = uint32(libc.Atoi(libc.GoString(&v10[0])))
							continue
						case 1:
							nox_xxx_readStr_517090(a1, (*uint8)(unsafe.Pointer(&v10[0])))
							*(*float32)(unsafe.Pointer(v6)) = float32(libc.Atof(libc.GoString(&v10[0])))
							continue
						case 2:
							nox_xxx_readStr_517090(a1, (*uint8)(unsafe.Pointer(&v10[0])))
							*(*uint32)(unsafe.Pointer(v6)) = uint32(nox_xxx_utilFindSound_40AF50(&v10[0]))
							continue
						case 3:
							nox_xxx_readStr_517090(a1, (*uint8)(unsafe.Pointer(&v10[0])))
							if nox_xxx_monsterLoadStrikeFn_549040(int32(uintptr(unsafe.Pointer(v3))), &v10[0]) != 0 {
								continue
							}
							return 0
						case 4:
							nox_xxx_readStr_517090(a1, (*uint8)(unsafe.Pointer(&v10[0])))
							if nox_xxx_monsterLoadDieFn_5490E0(int32(uintptr(unsafe.Pointer(v3))), &v10[0]) != 0 {
								continue
							}
							return 0
						case 5:
							nox_xxx_readStr_517090(a1, (*uint8)(unsafe.Pointer(&v10[0])))
							if nox_xxx_monsterLoadDeadFn_549180(int32(uintptr(unsafe.Pointer(v3))), &v10[0]) != 0 {
								continue
							}
							return 0
						case 6:
							v9 = 0
							nox_xxx_readStr_517090(a1, (*uint8)(unsafe.Pointer(&v10[0])))
							set_bitmask_flags_from_plus_separated_names_423930(&v10[0], (*uint32)(unsafe.Pointer(&v9)), (**byte)(memmap.PtrOff(0x587000, 247536)))
							*(*uint16)(unsafe.Pointer(v6)) = uint16(int16(v9))
							continue
						case 7:
							nox_xxx_readStr_517090(a1, (*uint8)(unsafe.Pointer(v6)))
							if libc.StrCmp(CString("NULL"), v6) == 0 {
								*v6 = 0
							}
							continue
						case 8:
							nox_xxx_readStr_517090(a1, (*uint8)(unsafe.Pointer(&v10[0])))
							*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*31)) = 18
							v7 = 0
							v8 = (*uint8)(memmap.PtrOff(0x587000, 247464))
						default:
							continue
						}
						for libc.StrCaseCmp(&v10[0], (*byte)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(v8))+7)))) != 0 {
							v8 = (*uint8)(unsafe.Add(unsafe.Pointer(v8), 4))
							v7++
							if int32(uintptr(unsafe.Pointer(v8))) >= int32(uintptr(memmap.PtrOff(0x587000, 247536))) {
								goto LABEL_27
							}
						}
						*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*31)) = uint32(v7)
					LABEL_27:
						if *(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*31)) == 18 {
							return 0
						}
					}
				} else {
					sub_517140(a1)
				}
			}
		}
	}
	return result
}
func nox_xxx_monsterListFree_5174F0() *uint32 {
	var (
		result *uint32
		v1     *uint32
	)
	result = *(**uint32)(unsafe.Pointer(&dword_5d4594_2386924))
	if dword_5d4594_2386924 != 0 {
		for {
			v1 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*61)))))
			alloc.Free(unsafe.Pointer(result))
			result = v1
			if v1 == nil {
				break
			}
		}
		dword_5d4594_2386924 = 0
	}
	return result
}
func nox_xxx_monsterList_517520() int32 {
	var (
		v0 int32
		v1 int32
	)
	v0 = int32(dword_5d4594_2386924)
	if dword_5d4594_2386924 == 0 {
		return 1
	}
	for {
		v1 = nox_xxx_getNameId_4E3AA0((*byte)(unsafe.Pointer(uintptr(v0))))
		*(*uint32)(unsafe.Pointer(uintptr(v0 + 240))) = uint32(v1)
		if v1 == 0 {
			break
		}
		v0 = int32(*(*uint32)(unsafe.Pointer(uintptr(v0 + 244))))
		if v0 == 0 {
			return 1
		}
	}
	nox_xxx_monsterListFree_5174F0()
	return 0
}
func nox_xxx_monsterDefByTT_517560(a1 int32) *uint32 {
	var result *uint32
	result = *(**uint32)(unsafe.Pointer(&dword_5d4594_2386924))
	if dword_5d4594_2386924 == 0 {
		return nil
	}
	for *(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*60)) != uint32(a1) {
		result = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*61)))))
		if result == nil {
			return nil
		}
	}
	return result
}
func sub_517590(a1 float32, a2 float32) int32 {
	var v4 int32
	nox_xxx_roundCoord_5175E0(a1, int32(uintptr(unsafe.Pointer(&a1))))
	nox_xxx_roundCoord_5175E0(a2, int32(uintptr(unsafe.Pointer(&v4))))
	return bool2int(float64(a1) >= 0.0 && (*(*int32)(unsafe.Add(unsafe.Pointer((*int32)(unsafe.Pointer(&a1))), unsafe.Sizeof(int32(0))*0))) < *(*int32)(unsafe.Pointer(&dword_5d4594_2386944)) && v4 >= 0 && v4 < *(*int32)(unsafe.Pointer(&dword_5d4594_2386944)))
}
func nox_xxx_roundCoord_5175E0(a1 float32, a2 int32) int32 {
	var result int32
	if float64(a1) >= 0.0 {
		result = a2
		*memmap.PtrUint32(6112660, 2386932) = uint32(uintptr(memmap.PtrOff(6112660, 2386952)))
		*mem_getFloatPtr(6112660, 0x246C08) = float32(float64(a1)*0.011764706 + 8.388608e+06)
		*(*uint32)(unsafe.Pointer(uintptr(a2))) = *memmap.PtrUint32(6112660, 2386952) & 0x7FFFFF
	} else {
		result = int(a1)
		*(*uint32)(unsafe.Pointer(uintptr(a2))) = uint32(result)
	}
	return result
}
func nox_xxx_unitCreateMissileSmth_517640(a1p *nox_object_t) int16 {
	var (
		a1  int32 = int32(uintptr(unsafe.Pointer(a1p)))
		v1  int32
		v2  int32
		v3  int32
		v4  int32
		v5  int32
		v6  int32
		v8  int32
		v9  int32
		v10 int32
		v11 int32
		v12 int32
	)
	v1 = a1
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 16))))
	if (v2&32) == 0 && (v2&512) == 0 && v2&4 != 0 {
		if (int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8)))) & 1) == 0 {
			sub_4E7350(a1)
			nox_xxx_roundCoord_5175E0(*(*float32)(unsafe.Pointer(uintptr(v1 + 232))), int32(uintptr(unsafe.Pointer(&v8))))
			nox_xxx_roundCoord_5175E0(*(*float32)(unsafe.Pointer(uintptr(v1 + 236))), int32(uintptr(unsafe.Pointer(&v10))))
			nox_xxx_roundCoord_5175E0(*(*float32)(unsafe.Pointer(uintptr(v1 + 240))), int32(uintptr(unsafe.Pointer(&a1))))
			nox_xxx_roundCoord_5175E0(*(*float32)(unsafe.Pointer(uintptr(v1 + 244))), int32(uintptr(unsafe.Pointer(&v9))))
			v3 = v8
			if v8 < 0 {
				v3 = 0
				v8 = 0
			}
			v4 = v10
			if v10 < 0 {
				v4 = 0
				v10 = 0
			}
			v5 = a1
			if a1 >= *(*int32)(unsafe.Pointer(&dword_5d4594_2386944)) {
				v5 = int32(dword_5d4594_2386944 - 1)
				a1 = int32(dword_5d4594_2386944 - 1)
			}
			v6 = v9
			if v9 >= *(*int32)(unsafe.Pointer(&dword_5d4594_2386944)) {
				v6 = int32(dword_5d4594_2386944 - 1)
				v9 = int32(dword_5d4594_2386944 - 1)
			}
			for *(*uint8)(unsafe.Pointer(uintptr(v1 + 336))) = 0; v4 <= v6; v4++ {
				if v3 <= v5 {
					for {
						nox_xxx_addObjToMapMB_517780(1, v3, v4, v1)
						v5 = a1
						v3++
						if v3 > a1 {
							break
						}
					}
					v3 = v8
					v6 = v9
				}
			}
		}
		nox_xxx_roundCoord_5175E0(*(*float32)(unsafe.Pointer(uintptr(v1 + 64))), int32(uintptr(unsafe.Pointer(&v12))))
		nox_xxx_roundCoord_5175E0(*(*float32)(unsafe.Pointer(uintptr(v1 + 68))), int32(uintptr(unsafe.Pointer(&v11))))
		nox_xxx_addObjToMapMB_517780(0, v12, v11, v1)
		v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 16))))
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v2))), 1)) |= 2
		*(*uint32)(unsafe.Pointer(uintptr(v1 + 16))) = uint32(v2)
	}
	return int16(v2)
}
func nox_xxx_addObjToMapMB_517780(a1 int32, a2 int32, a3 int32, a4 int32) int32 {
	if a2 < 0 || a3 < 0 {
		return 0
	}
	var result int32
	var v5 uint8
	var v6 int32
	var v7 int32
	if a1 != 0 {
		result = a1 - 1
		if a1 == 1 {
			v5 = *(*uint8)(unsafe.Pointer(uintptr(a4 + 336)))
			if int32(v5) < 4 {
				result = a4 + (int32(v5)+17)*16
				*(*uint8)(unsafe.Pointer(uintptr(a4 + 336))) = uint8(int8(int32(v5) + 1))
				*(*uint16)(unsafe.Pointer(uintptr(result))) = uint16(int16(a2))
				*(*uint16)(unsafe.Pointer(uintptr(result + 2))) = uint16(int16(a3))
				*(*uint32)(unsafe.Pointer(uintptr(result + 12))) = uint32(a4)
				*(*uint32)(unsafe.Pointer(uintptr(result + 8))) = 0
				*(*uint32)(unsafe.Pointer(uintptr(result + 4))) = *(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_2386940 + uint32(a2*4)))) + uint32(a3*16) + 4)))
				v6 = int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_2386940 + uint32(a2*4)))) + uint32(a3*16) + 4))))
				if v6 != 0 {
					*(*uint32)(unsafe.Pointer(uintptr(v6 + 8))) = uint32(result)
				}
				*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_2386940 + uint32(a2*4)))) + uint32(a3*16) + 4))) = uint32(result)
			}
		}
	} else {
		*(*uint32)(unsafe.Pointer(uintptr(a4 + 268))) = uint32(a4)
		result = a4 + 256
		*(*uint16)(unsafe.Pointer(uintptr(a4 + 256))) = uint16(int16(a2))
		*(*uint16)(unsafe.Pointer(uintptr(a4 + 258))) = uint16(int16(a3))
		*(*uint32)(unsafe.Pointer(uintptr(a4 + 264))) = 0
		*(*uint32)(unsafe.Pointer(uintptr(a4 + 260))) = *(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_2386940 + uint32(a2*4)))) + uint32(a3*16))))
		v7 = int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_2386940 + uint32(a2*4)))) + uint32(a3*16)))))
		if v7 != 0 {
			*(*uint32)(unsafe.Pointer(uintptr(v7 + 8))) = uint32(result)
		}
		*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_2386940 + uint32(a2*4)))) + uint32(a3*16)))) = uint32(result)
	}
	return result
}
func sub_517870(a1p *nox_object_t) int16 {
	var (
		a1 int32 = int32(uintptr(unsafe.Pointer(a1p)))
		v1 int32
		v2 int32
		v3 *uint16
	)
	v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 16))))
	if v1&512 != 0 {
		sub_5178E0(0, (*uint16)(unsafe.Pointer(uintptr(a1+256))))
		if (int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8)))) & 1) == 0 {
			v2 = 0
			if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 336)))) != 0 {
				v3 = (*uint16)(unsafe.Pointer(uintptr(a1 + 272)))
				for {
					sub_5178E0(1, v3)
					v2++
					v3 = (*uint16)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint16(0))*8))
					if v2 >= int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 336)))) {
						break
					}
				}
			}
			*(*uint8)(unsafe.Pointer(uintptr(a1 + 336))) = 0
		}
		v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 16))))
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v1))), 1)) &= 253
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 16))) = uint32(v1)
	}
	return int16(v1)
}
func sub_5178E0(a1 int32, a2 *uint16) *uint16 {
	var (
		result *uint16
		v3     int32
		v4     int32
		v5     int32
		v6     int32
	)
	if a1 != 0 {
		result = (*uint16)(unsafe.Pointer(uintptr(a1 - 1)))
		if a1 == 1 {
			result = a2
			v3 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a2))), unsafe.Sizeof(uint32(0))*2))))
			if v3 != 0 {
				*(*uint32)(unsafe.Pointer(uintptr(v3 + 4))) = *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a2))), unsafe.Sizeof(uint32(0))*1)))
			} else {
				*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_2386940 + uint32(int32(*a2)*4)))) + uint32(int32(*(*uint16)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint16(0))*1)))*16) + 4))) = *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a2))), unsafe.Sizeof(uint32(0))*1)))
			}
			v4 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a2))), unsafe.Sizeof(uint32(0))*1))))
			if v4 != 0 {
				*(*uint32)(unsafe.Pointer(uintptr(v4 + 8))) = *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a2))), unsafe.Sizeof(uint32(0))*2)))
			}
		}
	} else {
		result = a2
		v5 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a2))), unsafe.Sizeof(uint32(0))*2))))
		if v5 != 0 {
			*(*uint32)(unsafe.Pointer(uintptr(v5 + 4))) = *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a2))), unsafe.Sizeof(uint32(0))*1)))
		} else {
			*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_2386940 + uint32(int32(*a2)*4)))) + uint32(int32(*(*uint16)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint16(0))*1)))*16)))) = *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a2))), unsafe.Sizeof(uint32(0))*1)))
		}
		v6 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a2))), unsafe.Sizeof(uint32(0))*1))))
		if v6 != 0 {
			*(*uint32)(unsafe.Pointer(uintptr(v6 + 8))) = *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a2))), unsafe.Sizeof(uint32(0))*2)))
		}
	}
	return result
}
func nox_xxx_waypointMapRegister_5179B0(a1 int32) {
	var (
		v1 int32
		v2 int32
		v3 int32
		v4 int32
		v5 int32
		v6 int16
		v7 int32
		v8 int32
	)
	v1 = a1
	if (int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 480)))) & 2) == 0 {
		nox_xxx_roundCoord_5175E0(*(*float32)(unsafe.Pointer(uintptr(a1 + 8))), int32(uintptr(unsafe.Pointer(&a1))))
		nox_xxx_roundCoord_5175E0(*(*float32)(unsafe.Pointer(uintptr(v1 + 12))), int32(uintptr(unsafe.Pointer(&v8))))
		v2 = v8
		v3 = a1
		v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_2386940 + uint32(a1*4)))) + uint32(v8*16) + 8))))
		if v4 != 0 {
			*(*uint32)(unsafe.Pointer(uintptr(v4 + 500))) = uint32(v1)
			v3 = a1
			v2 = v8
		}
		*(*uint32)(unsafe.Pointer(uintptr(v1 + 500))) = 0
		v5 = v2 * 16
		*(*uint32)(unsafe.Pointer(uintptr(v1 + 496))) = *(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_2386940 + uint32(v3*4)))) + uint32(v5) + 8)))
		*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_2386940 + uint32(v3*4)))) + uint32(v5) + 8))) = uint32(v1)
		v6 = int16(a1)
		*(*uint16)(unsafe.Pointer(uintptr(v1 + 494))) = uint16(int16(v8))
		v7 = int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 480))))
		*(*uint16)(unsafe.Pointer(uintptr(v1 + 492))) = uint16(v6)
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v7))), 0)) = uint8(int8(v7 | 2))
		*(*uint32)(unsafe.Pointer(uintptr(v1 + 480))) = uint32(v7)
	}
}
func sub_517A70(a1 int32) int32 {
	var (
		result int32
		v2     int32
		v3     int32
	)
	result = a1
	if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 480))))&2 != 0 {
		v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 500))))
		if v2 != 0 {
			*(*uint32)(unsafe.Pointer(uintptr(v2 + 496))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 496)))
		} else {
			*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_2386940 + uint32(int32(*(*int16)(unsafe.Pointer(uintptr(a1 + 492))))*4)))) + uint32(int32(*(*int16)(unsafe.Pointer(uintptr(a1 + 494))))*16) + 8))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 496)))
		}
		v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 496))))
		if v3 != 0 {
			*(*uint32)(unsafe.Pointer(uintptr(v3 + 500))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 500)))
		}
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 480))) &= 0xFFFFFFFD
	}
	return result
}
func sub_517AE0() int32 {
	var (
		v0 uint32
		i  int32
	)
	dword_5d4594_2386944 = 70
	dword_5d4594_2386940 = uint32(uintptr(alloc.Calloc(1, 280)))
	v0 = dword_5d4594_2386944
	for i = 0; i < *(*int32)(unsafe.Pointer(&dword_5d4594_2386944)); i++ {
		*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_2386940 + uint32(i*4)))) = uint32(uintptr(alloc.Calloc(int(v0), 16)))
		v0 = dword_5d4594_2386944
	}
	return 1
}
func sub_517B30() {
	var i int32
	for i = 0; i < *(*int32)(unsafe.Pointer(&dword_5d4594_2386944)); i++ {
		*(*unsafe.Pointer)(unsafe.Pointer(uintptr(dword_5d4594_2386940 + uint32(i*4)))) = nil
	}
	*(*unsafe.Pointer)(unsafe.Pointer(&dword_5d4594_2386940)) = nil
}
func sub_517B70(a1 *float2, a2 func(uint32, int32), a3 int32) {
	var (
		v3 int32
		v4 int32
		i  int32
		v6 int32
		v7 int32
	)
	if a2 == nil {
		return
	}
	nox_xxx_roundCoord_5175E0(a1.field_0, int32(uintptr(unsafe.Pointer(&v6))))
	nox_xxx_roundCoord_5175E0(a1.field_4, int32(uintptr(unsafe.Pointer(&v7))))
	v3 = v6
	if v6 >= 0 {
		if v6 < *(*int32)(unsafe.Pointer(&dword_5d4594_2386944)) {
			goto LABEL_7
		}
		v3 = int32(dword_5d4594_2386944 - 1)
	} else {
		v3 = 0
	}
	v6 = v3
LABEL_7:
	v4 = v7
	if v7 >= 0 {
		if v7 < *(*int32)(unsafe.Pointer(&dword_5d4594_2386944)) {
			goto LABEL_12
		}
		v4 = int32(dword_5d4594_2386944 - 1)
	} else {
		v4 = 0
	}
	v7 = v4
LABEL_12:
	for i = int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_2386940 + uint32(v3*4)))) + uint32(v4*16) + 4)))); i != 0; i = int32(*(*uint32)(unsafe.Pointer(uintptr(i + 4)))) {
		a2(*(*uint32)(unsafe.Pointer(uintptr(i + 12))), a3)
	}
}
func nox_xxx_getUnitsInRect_517C10(a1 *float4, a2 func(*float32, int32), a3 unsafe.Pointer) {
	var (
		v3  int32
		v4  int32
		v5  func(*float32, int32)
		v6  int32
		v7  int32
		v8  *float4
		v9  int32
		v10 int32
		v11 int32
		v12 int32
		v13 int32
		v14 int32
		v15 int32
		v16 int32
		v17 int32
		v18 *uint32
		v19 int32
		v20 *float32
		v21 int32
		v22 int32
		v23 int32
		v24 int32
		v25 int32
		v26 int32
		v27 int32
		v28 int32
	)
	v4 = int32(nox_xxx_triggersCount_587000_249172)
	v3 = v4
	if v4 < 1 {
		v5 = a2
		if a2 != nil {
			v6 = v3 + 1
			v7 = int32(*memmap.PtrUint32(6112660, uint32(v6*4)+0x246C14))
			nox_xxx_triggersCount_587000_249172 = uint32(v6)
			*memmap.PtrUint32(6112660, uint32(v6*4)+0x246C14) = uint32(v7 + 1)
			v8 = a1
			nox_xxx_roundCoord_5175E0(a1.field_0, int32(uintptr(unsafe.Pointer(&v24))))
			nox_xxx_roundCoord_5175E0(a1.field_4, int32(uintptr(unsafe.Pointer(&v25))))
			nox_xxx_roundCoord_5175E0(a1.field_8, int32(uintptr(unsafe.Pointer(&v22))))
			nox_xxx_roundCoord_5175E0(a1.field_C, int32(uintptr(unsafe.Pointer(&v23))))
			v9 = v24
			if v24 < 0 {
				v24 = 0
				v9 = 0
			}
			v10 = v22
			if v22 >= *(*int32)(unsafe.Pointer(&dword_5d4594_2386944)) {
				v10 = int32(dword_5d4594_2386944 - 1)
				v22 = int32(dword_5d4594_2386944 - 1)
			}
			v11 = v25
			if v25 < 0 {
				v11 = 0
				v25 = 0
			}
			if v23 >= *(*int32)(unsafe.Pointer(&dword_5d4594_2386944)) {
				v23 = int32(dword_5d4594_2386944 - 1)
			}
			v12 = v11
			v26 = v11
			if v11 <= v23 {
				v13 = int32(uintptr(a3))
				for {
					v14 = v9
					v27 = v9
					if v9 <= v10 {
						v15 = v12 * 16
						v28 = v15
						for {
							v16 = int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_2386940 + uint32(v14*4)))) + uint32(v15) + 4))))
							if v16 != 0 {
								for {
									v17 = int32(nox_xxx_triggersCount_587000_249172)
									v18 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v16 + 12))) + uint32(v17*4) + 248)))
									v19 = int32(*memmap.PtrUint32(6112660, uint32(v17*4)+0x246C14))
									if *v18 != uint32(v19) {
										*v18 = uint32(v19)
										v20 = *(**float32)(unsafe.Pointer(uintptr(v16 + 12)))
										if float64(*(*float32)(unsafe.Add(unsafe.Pointer(v20), unsafe.Sizeof(float32(0))*58))) < float64(v8.field_8) && float64(*(*float32)(unsafe.Add(unsafe.Pointer(v20), unsafe.Sizeof(float32(0))*60))) > float64(v8.field_0) && float64(*(*float32)(unsafe.Add(unsafe.Pointer(v20), unsafe.Sizeof(float32(0))*59))) < float64(v8.field_C) && float64(*(*float32)(unsafe.Add(unsafe.Pointer(v20), unsafe.Sizeof(float32(0))*61))) > float64(v8.field_4) {
											v5(v20, v13)
										}
									}
									v16 = int32(*(*uint32)(unsafe.Pointer(uintptr(v16 + 4))))
									if v16 == 0 {
										break
									}
								}
								v15 = v28
								v10 = v22
								v14 = v27
							}
							v27 = func() int32 {
								p := &v14
								*p++
								return *p
							}()
							if v14 > v10 {
								break
							}
						}
						v12 = v26
						v9 = v24
					}
					v26 = func() int32 {
						p := &v12
						*p++
						return *p
					}()
					if v12 > v23 {
						break
					}
				}
			}
			v21 = int32(nox_xxx_triggersCount_587000_249172)
			nox_xxx_triggersCount_587000_249172 = uint32(v21 - 1)
		}
	}
}
func nox_xxx_getUnitsInRectAdvImpl_517DC0(a1 *float4, a2 func(int32, int32), a3 int32) {
	var (
		v3  func(int32, int32)
		v4  int32
		v5  int32
		v6  int32
		v7  int32
		v8  int32
		v9  int32
		v10 int32
		v11 int32
		v12 int32
		v13 int32
		v14 int32
		v15 *float4
		i   int32
	)
	v3 = a2
	if a2 != nil {
		nox_xxx_roundCoord_5175E0(a1.field_0, int32(uintptr(unsafe.Pointer(&v13))))
		nox_xxx_roundCoord_5175E0(a1.field_4, int32(uintptr(unsafe.Pointer(&v14))))
		nox_xxx_roundCoord_5175E0(a1.field_8, int32(uintptr(unsafe.Pointer(&v11))))
		nox_xxx_roundCoord_5175E0(a1.field_C, int32(uintptr(unsafe.Pointer(&v12))))
		v4 = v13
		if v13 < 0 {
			v13 = 0
			v4 = 0
		}
		v5 = v11
		if v11 >= *(*int32)(unsafe.Pointer(&dword_5d4594_2386944)) {
			v5 = int32(dword_5d4594_2386944 - 1)
			v11 = int32(dword_5d4594_2386944 - 1)
		}
		v6 = v14
		if v14 < 0 {
			v6 = 0
			v14 = 0
		}
		if v12 >= *(*int32)(unsafe.Pointer(&dword_5d4594_2386944)) {
			v12 = int32(dword_5d4594_2386944 - 1)
		}
		v7 = v6
		for i = v6; v7 <= v12; i = v7 {
			if v4 <= v5 {
				v8 = v7 * 16
				v15 = (*float4)(unsafe.Pointer(uintptr(v7 * 16)))
				for {
					v9 = int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_2386940 + uint32(v4*4)))) + uint32(v8)))))
					if v9 != 0 {
						for {
							v10 = int32(*(*uint32)(unsafe.Pointer(uintptr(v9 + 12))))
							if int32(*(*uint8)(unsafe.Pointer(uintptr(v10 + 8))))&1 != 0 {
								v3(v10, a3)
							}
							v9 = int32(*(*uint32)(unsafe.Pointer(uintptr(v9 + 4))))
							if v9 == 0 {
								break
							}
						}
						v5 = v11
						v8 = int32(uintptr(unsafe.Pointer(v15)))
					}
					v4++
					if v4 > v5 {
						break
					}
				}
				v4 = v13
				v7 = i
			}
			v7++
		}
	}
}
func nox_xxx_getUnitsInRectAdv_517ED0(a1 *float4, a2 func(*float32, int32), a3 int32) {
	nox_xxx_getUnitsInRect_517C10(a1, a2, unsafe.Pointer(uintptr(a3)))
	nox_xxx_getUnitsInRectAdvImpl_517DC0(a1, func(arg1 int32, arg2 int32) {
		a2((*float32)(unsafe.Pointer(uintptr(arg1))), arg2)
	}, a3)
}
func nox_xxx_secretWallCheckUnits_517F00(a1 *float32, a2 func(*int32, int32) int32, a3 int32) *int32 {
	var (
		result *int32
		i      *int32
		v5     float64
		v6     float64
	)
	result = (*int32)(nox_xxx_wallSecretGetFirstWall_410780())
	for i = result; result != nil; i = result {
		v5 = float64(*(*int32)(unsafe.Add(unsafe.Pointer(i), unsafe.Sizeof(int32(0))*1)) * 23)
		if v5 > float64(*a1) && v5 < float64(*(*float32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(float32(0))*2))) {
			v6 = float64(*(*int32)(unsafe.Add(unsafe.Pointer(i), unsafe.Sizeof(int32(0))*2)) * 23)
			if v6 > float64(*(*float32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(float32(0))*1))) && v6 < float64(*(*float32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(float32(0))*3))) {
				a2(i, a3)
			}
		}
		result = (*int32)(unsafe.Pointer(uintptr(nox_xxx_wallSecretNext_410790(i))))
	}
	return result
}
func nox_xxx_unitsGetInCircle_517F90(a1 *float2, r float32, a3 unsafe.Pointer, a4 *nox_object_t) {
	var (
		a3a [4]int32
		a1a float4
	)
	a3a[0] = int32(uintptr(unsafe.Pointer(a1)))
	a3a[2] = int32(uintptr(a3))
	*(*float32)(unsafe.Pointer(&a3a[1])) = r * r
	a3a[3] = int32(uintptr(unsafe.Pointer(a4)))
	a1a.field_0 = a1.field_0 - r
	a1a.field_4 = a1.field_4 - r
	a1a.field_8 = a1.field_0 + r
	a1a.field_C = a1.field_4 + r
	nox_xxx_getUnitsInRect_517C10(&a1a, nox_xxx_unitsGetInNotFarFn_518000, unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(&a3a[0]))))))
}
func nox_xxx_unitsGetInNotFarFn_518000(a1 *float32, a2 int32) {
	var (
		v2 float64
		v3 float64
	)
	v2 = float64(**(**float32)(unsafe.Pointer(uintptr(a2))) - *(*float32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(float32(0))*14)))
	v3 = float64(*(*float32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a2))) + 4))) - *(*float32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(float32(0))*15)))
	if v3*v3+v2*v2 < float64(*(*float32)(unsafe.Pointer(uintptr(a2 + 4)))) {
		(cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(a2 + 8))), (*func(*float32, uint32))(nil)))(a1, *(*uint32)(unsafe.Pointer(uintptr(a2 + 12))))
	}
}
func sub_518040(arg0 int32, a2 float32, arg8 int32, a4 int32) int32 {
	var (
		result int32
		v5     float64
		v6     float64
		a3     [4]int32
		a1     float4
	)
	result = arg0
	if arg0 != 0 {
		v5 = float64(*(*float32)(unsafe.Pointer(uintptr(arg0))) - a2)
		a3[2] = arg8
		a3[0] = arg0
		*(*float32)(unsafe.Pointer(&a3[1])) = a2
		a1.field_0 = float32(v5)
		v6 = float64(*(*float32)(unsafe.Pointer(uintptr(arg0 + 4))) - a2)
		a3[3] = a4
		a1.field_4 = float32(v6)
		a1.field_8 = a2 + *(*float32)(unsafe.Pointer(uintptr(arg0)))
		a1.field_C = a2 + *(*float32)(unsafe.Pointer(uintptr(arg0 + 4)))
		nox_xxx_getUnitsInRect_517C10(&a1, func(arg1 *float32, arg2 int32) {
			sub_5180B0(int32(uintptr(unsafe.Pointer(arg1))), arg2)
		}, unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(&a3[0]))))))
	}
	return result
}
func sub_5180B0(a1 int32, a2 int32) {
	var (
		v2 *float2
		v3 float64
		v4 float64
		v5 float64
		v6 *float2
		v7 float64
		v8 float2
	)
	if a1 != 0 && a2 != 0 {
		if *(*uint32)(unsafe.Pointer(uintptr(a1 + 172))) == 2 {
			v6 = *(**float2)(unsafe.Pointer(uintptr(a2)))
			v8.field_0 = **(**float32)(unsafe.Pointer(uintptr(a2))) - *(*float32)(unsafe.Pointer(uintptr(a1 + 56)))
			v7 = float64(v6.field_4 - *(*float32)(unsafe.Pointer(uintptr(a1 + 60))))
			v8.field_4 = float32(v7)
			v4 = math.Sqrt(v7*float64(v8.field_4)+float64(v8.field_0*v8.field_0)) - float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 176))))
		} else {
			if *(*uint32)(unsafe.Pointer(uintptr(a1 + 172))) == 3 {
				v5 = sub_54A990((*float2)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a2)))))), *(*float32)(unsafe.Pointer(uintptr(a2 + 4))), a1, &v8)
				goto LABEL_9
			}
			v2 = *(**float2)(unsafe.Pointer(uintptr(a2)))
			v8.field_0 = **(**float32)(unsafe.Pointer(uintptr(a2))) - *(*float32)(unsafe.Pointer(uintptr(a1 + 56)))
			v3 = float64(v2.field_4 - *(*float32)(unsafe.Pointer(uintptr(a1 + 60))))
			v8.field_4 = float32(v3)
			v4 = math.Sqrt(v3*float64(v8.field_4) + float64(v8.field_0*v8.field_0))
		}
		v5 = float64(*(*float32)(unsafe.Pointer(uintptr(a2 + 4)))) - v4
	LABEL_9:
		if v5 > 0.0 {
			(cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(a2 + 8))), (*func(int32, uint32))(nil)))(a1, *(*uint32)(unsafe.Pointer(uintptr(a2 + 12))))
		}
	}
}
func sub_518170(a1p unsafe.Pointer, a2 float32, a3 unsafe.Pointer, a4p *nox_object_t) {
	var (
		a1  int32 = int32(uintptr(a1p))
		a4  int32 = int32(uintptr(unsafe.Pointer(a4p)))
		v4  func(int32, int32)
		v5  int32
		v6  int32
		v7  int32
		v8  float64
		v9  int32
		v10 int32
		v11 int32
		v12 int32
		v13 int32
		v14 float64
		v15 float64
		v16 float32
		v17 float32
		v18 float32
		v19 float32
		v20 int32
		v21 int32
		v22 int32
		v23 int32
		v24 int32
		v25 int32
		v26 int32
	)
	if a3 == nil {
		return
	}
	v4 = cgoAsFunc(a3, (*func(int32, int32))(nil)).(func(int32, int32))
	v16 = *(*float32)(unsafe.Pointer(uintptr(a1))) - a2
	nox_xxx_roundCoord_5175E0(v16, int32(uintptr(unsafe.Pointer(&v22))))
	v17 = *(*float32)(unsafe.Pointer(uintptr(a1 + 4))) - a2
	nox_xxx_roundCoord_5175E0(v17, int32(uintptr(unsafe.Pointer(&v23))))
	v18 = a2 + *(*float32)(unsafe.Pointer(uintptr(a1)))
	nox_xxx_roundCoord_5175E0(v18, int32(uintptr(unsafe.Pointer(&v20))))
	v19 = a2 + *(*float32)(unsafe.Pointer(uintptr(a1 + 4)))
	nox_xxx_roundCoord_5175E0(v19, int32(uintptr(unsafe.Pointer(&v21))))
	v5 = v22
	if v22 < 0 {
		v22 = 0
		v5 = 0
	}
	v6 = v20
	if v20 >= *(*int32)(unsafe.Pointer(&dword_5d4594_2386944)) {
		v6 = int32(dword_5d4594_2386944 - 1)
		v20 = int32(dword_5d4594_2386944 - 1)
	}
	v7 = v23
	if v23 < 0 {
		v7 = 0
		v23 = 0
	}
	if v21 >= *(*int32)(unsafe.Pointer(&dword_5d4594_2386944)) {
		v21 = int32(dword_5d4594_2386944 - 1)
	}
	v8 = float64(a2 * a2)
	v9 = v7
	v24 = v7
	*(*float32)(unsafe.Pointer(&v25)) = float32(v8)
	if v7 > v21 {
		return
	}
	v10 = a4
	for {
		if v5 <= v6 {
			v11 = v9 * 16
			v26 = v9 * 16
			for {
				v12 = int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_2386940 + uint32(v5*4)))) + uint32(v11)))))
				if v12 != 0 {
					for {
						v13 = int32(*(*uint32)(unsafe.Pointer(uintptr(v12 + 12))))
						if int32(*(*uint8)(unsafe.Pointer(uintptr(v13 + 8))))&1 != 0 {
							v14 = float64(*(*float32)(unsafe.Pointer(uintptr(a1))) - *(*float32)(unsafe.Pointer(uintptr(v13 + 56))))
							v15 = float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 4))) - *(*float32)(unsafe.Pointer(uintptr(v13 + 60))))
							if v15*v15+v14*v14 <= float64(*(*float32)(unsafe.Pointer(&v25))) {
								v4(v13, v10)
							}
						}
						v12 = int32(*(*uint32)(unsafe.Pointer(uintptr(v12 + 4))))
						if v12 == 0 {
							break
						}
					}
					v11 = v26
					v6 = v20
				}
				v5++
				if v5 > v6 {
					break
				}
			}
			v5 = v22
			v9 = v24
		}
		v24 = func() int32 {
			p := &v9
			*p++
			return *p
		}()
		if v9 > v21 {
			break
		}
	}
}
func sub_518460(a1 *float2, a2 uint8, a3 int32) int32 {
	var (
		v3  int32
		v5  float32
		v6  float32
		v7  float32
		v8  float32
		v9  float32
		v10 int4
	)
	*memmap.PtrUint32(6112660, 2386960)++
	dword_5d4594_2386928 = 0x447A0000
	v9 = 0.0
	v3 = 0
	for {
		v5 = a1.field_0 - v9
		nox_xxx_roundCoord_5175E0(v5, int32(uintptr(unsafe.Pointer(&v10.field_0))))
		v6 = a1.field_4 - v9
		nox_xxx_roundCoord_5175E0(v6, int32(uintptr(unsafe.Pointer(&v10.field_4))))
		v7 = v9 + a1.field_0
		nox_xxx_roundCoord_5175E0(v7, int32(uintptr(unsafe.Pointer(&v10.field_8))))
		v8 = v9 + a1.field_4
		nox_xxx_roundCoord_5175E0(v8, int32(uintptr(unsafe.Pointer(&v10.field_C))))
		dword_5d4594_2386948 = 0
		sub_518550(&v10.field_0, (*int32)(unsafe.Pointer(a1)), a2, a3)
		if dword_5d4594_2386948 != 0 {
			v3 = int32(dword_5d4594_2386948)
			v9 = *(*float32)(unsafe.Pointer(&dword_5d4594_2386928))
		} else {
			if v3 != 0 {
				return v3
			}
			v9 = float32(float64(v9) + 85.0)
		}
		if float64(v9) >= 1000.0 {
			break
		}
	}
	return v3
}
func sub_518550(a1 *int32, a2 *int32, a3 uint8, a4 int32) int32 {
	var (
		v4     *int32
		v5     int32
		v6     float32
		result int32
		v8     int32
		v9     int32
		v10    int32
		v11    int32
		v12    int32
		v13    int32
		v14    int32
		v15    *int32
		v16    float64
		v17    float64
		v18    float64
		v19    int32
		v20    int32
		v21    int32
		v22    float32
		v23    float4
	)
	v4 = a1
	if *a1 < 0 {
		*a1 = 0
	}
	v5 = int32(dword_5d4594_2386944)
	if *(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*2)) >= *(*int32)(unsafe.Pointer(&dword_5d4594_2386944)) {
		*(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*2)) = int32(dword_5d4594_2386944 - 1)
		v5 = int32(dword_5d4594_2386944)
	}
	if *(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*1)) < 0 {
		*(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*1)) = 0
		v5 = int32(dword_5d4594_2386944)
	}
	if *(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*3)) >= v5 {
		*(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*3)) = v5 - 1
	}
	v6 = *((*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(a2))), unsafe.Sizeof(float32(0))*1)))
	result = *(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*1))
	*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v23.field_0))), unsafe.Sizeof(uint32(0))*0)) = uint32(*a2)
	v8 = *(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*3))
	v23.field_4 = v6
	v20 = result
	if result <= v8 {
		v9 = result * 16
		v21 = result * 16
		for {
			v10 = *v4
			v19 = *v4
			if *v4 <= *(*int32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(int32(0))*2)) {
				for {
					v11 = int32(uint32(v9) + *(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_2386940 + uint32(v10*4)))))
					if *(*uint32)(unsafe.Pointer(uintptr(v11 + 12))) != *memmap.PtrUint32(6112660, 2386960) {
						v12 = int32(*(*uint32)(unsafe.Pointer(uintptr(v11 + 8))))
						if v12 != 0 {
							for {
								if int32(*(*uint8)(unsafe.Pointer(uintptr(v12 + 480))))&1 != 0 && sub_579EE0(v12, a3) != 0 {
									if a4 == 0 {
										goto LABEL_35
									}
									v13 = 0
									v14 = 0
									if int32(*(*uint8)(unsafe.Pointer(uintptr(v12 + 476)))) != 0 {
										v15 = (*int32)(unsafe.Pointer(uintptr(v12 + 92)))
										for {
											if int32(*(*uint8)(unsafe.Pointer(uintptr(*v15 + 480))))&1 != 0 && sub_579EE0(*v15, a3) != 0 {
												v13++
											}
											v14++
											v15 = (*int32)(unsafe.Add(unsafe.Pointer(v15), unsafe.Sizeof(int32(0))*2))
											if v14 >= int32(*(*uint8)(unsafe.Pointer(uintptr(v12 + 476)))) {
												break
											}
										}
										v4 = a1
										if v13 != 0 {
										LABEL_35:
											v16 = float64(*(*float32)(unsafe.Pointer(a2)) - *(*float32)(unsafe.Pointer(uintptr(v12 + 8))))
											v17 = float64(*((*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(a2))), unsafe.Sizeof(float32(0))*1))) - *(*float32)(unsafe.Pointer(uintptr(v12 + 12))))
											v18 = math.Sqrt(v17*v17 + v16*v16)
											if v18 < float64(*(*float32)(unsafe.Pointer(&dword_5d4594_2386928))) {
												v23.field_8 = *(*float32)(unsafe.Pointer(uintptr(v12 + 8)))
												v23.field_C = *(*float32)(unsafe.Pointer(uintptr(v12 + 12)))
												if nox_xxx_mapTraceRay_535250(&v23, nil, nil, 1) != 0 {
													dword_5d4594_2386948 = uint32(v12)
													v22 = float32(v18)
													*(*float32)(unsafe.Pointer(&dword_5d4594_2386928)) = v22
												}
											}
										}
									}
								}
								v12 = int32(*(*uint32)(unsafe.Pointer(uintptr(v12 + 496))))
								if v12 == 0 {
									break
								}
							}
							v9 = v21
							v10 = v19
						}
						*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_2386940 + uint32(v10*4)))) + uint32(v9) + 12))) = *memmap.PtrUint32(6112660, 2386960)
					}
					v19 = func() int32 {
						p := &v10
						*p++
						return *p
					}()
					if v10 > *(*int32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(int32(0))*2)) {
						break
					}
				}
				result = v20
			}
			result++
			v9 += 16
			v20 = result
			v21 = v9
			if result > *(*int32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(int32(0))*3)) {
				break
			}
		}
	}
	return result
}
func sub_518740(a1 *float2, a2 uint8) int32 {
	return sub_518460(a1, a2, 1)
}
func sub_518770() int32 {
	var result int32
	result = nox_xxx_getNameId_4E3AA0(CString("TeleportPentagram"))
	*memmap.PtrUint32(6112660, 2386972) = uint32(result)
	if result == 0 {
		return 0
	}
	result = nox_xxx_getNameId_4E3AA0(CString("PressurePlate"))
	*memmap.PtrUint32(6112660, 2386976) = uint32(result)
	if result == 0 {
		return 0
	}
	result = nox_xxx_getNameId_4E3AA0(CString("Spike"))
	*memmap.PtrUint32(6112660, 2386980) = uint32(result)
	if result == 0 {
		return 0
	}
	result = nox_xxx_getNameId_4E3AA0(CString("PeriodicSpike"))
	*memmap.PtrUint32(6112660, 2386984) = uint32(result)
	if result == 0 {
		return 0
	}
	return 1
}
func nox_xxx_netSendPhantomPlrMb_5187E0(a1 int32, a2 int32) int32 {
	var (
		v2  int16
		v3  float32
		v4  int16
		v5  float32
		v6  int32
		v7  int8
		a2a int2
		v10 [11]byte
	)
	v10[0] = 48
	*(*uint16)(unsafe.Pointer(&v10[3])) = *(*uint16)(unsafe.Pointer(uintptr(a2 + 4)))
	v2 = int16(uint16(nox_xxx_netGetUnitCodeServ_578AC0((*nox_object_t)(unsafe.Pointer(uintptr(a2))))))
	v3 = *(*float32)(unsafe.Pointer(uintptr(a2 + 56)))
	*(*uint16)(unsafe.Pointer(&v10[1])) = uint16(v2)
	v4 = int16(int(v3))
	v5 = *(*float32)(unsafe.Pointer(uintptr(a2 + 60)))
	*(*uint16)(unsafe.Pointer(&v10[5])) = uint16(v4)
	v6 = int(v5)
	*(*uint16)(unsafe.Pointer(&v10[7])) = uint16(int16(v6))
	if int32(*(*uint8)(unsafe.Pointer(uintptr(a2 + 8))))&1 != 0 && int32(*(*uint8)(unsafe.Pointer(uintptr(a2 + 12))))&48 != 0 {
		*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v6))), unsafe.Sizeof(uint16(0))*0)) = *(*uint16)(unsafe.Pointer(uintptr(a2 + 124)))
		v10[10] = byte(int8(v6 >> 3))
	} else {
		v10[10] = math.MaxUint8
	}
	nox_xxx_xferIndexedDirection_509E20(int32(*(*int16)(unsafe.Pointer(uintptr(a2 + 124)))), &a2a)
	if int32(uint8(int8(int32(*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a2a.field_0))), 0)))+int32(*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a2a.field_4))), 0)))*3+4))) <= 3 {
		v7 = int8(int32(*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a2a.field_0))), 0))) + int32(*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a2a.field_4))), 0)))*3 + 4)
	} else {
		v7 = int8(int32(*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a2a.field_0))), 0))) + int32(*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a2a.field_4))), 0)))*3 + 3)
	}
	v10[9] = byte(int8(int32(v7) * 16))
	return bool2int(nox_netlist_addToMsgListSrv_40EF40(a1, (*uint8)(unsafe.Pointer(&v10[0])), 11))
}
func nox_xxx_netSendSimpleObj_5188A0(a1 int32, a2 int32) int32 {
	var (
		v2  *int16
		v3  int32
		v4  int16
		v5  int16
		v6  int16
		v7  int16
		v8  float32
		v10 [9]byte
	)
	if *(*uint32)(unsafe.Pointer(uintptr(a2 + 8)))&0x20000 != 0 {
		v2 = *(**int16)(unsafe.Pointer(uintptr(a2 + 556)))
		if v2 != nil {
			v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 748))))
			if (nox_frame_xxx_2598000 - *(*uint32)(unsafe.Pointer(uintptr(a2 + 536)))) > 2 {
				v4 = *v2
				v5 = int16(*(*uint16)(unsafe.Pointer(uintptr(v3 + a1*2 + 96))))
				if int32(v4) != int32(v5) {
					nox_xxx_netReportHealthDelta_4D8760(a1, int16(uint16(*(*uint32)(unsafe.Pointer(uintptr(a2 + 36))))), int16(int32(v4)-int32(v5)))
					*(*uint16)(unsafe.Pointer(uintptr(v3 + a1*2 + 96))) = **(**uint16)(unsafe.Pointer(uintptr(a2 + 556)))
				}
			}
		}
	}
	v6 = int16(*(*uint16)(unsafe.Pointer(uintptr(a2 + 4))))
	v10[0] = 47
	*(*uint16)(unsafe.Pointer(&v10[3])) = uint16(v6)
	*(*uint16)(unsafe.Pointer(&v10[1])) = uint16(nox_xxx_netGetUnitCodeServ_578AC0((*nox_object_t)(unsafe.Pointer(uintptr(a2)))))
	v7 = int16(int(*(*float32)(unsafe.Pointer(uintptr(a2 + 56)))))
	v8 = *(*float32)(unsafe.Pointer(uintptr(a2 + 60)))
	*(*uint16)(unsafe.Pointer(&v10[5])) = uint16(v7)
	*(*uint16)(unsafe.Pointer(&v10[7])) = uint16(int16(int(v8)))
	return bool2int(nox_netlist_addToMsgListSrv_40EF40(a1, (*uint8)(unsafe.Pointer(&v10[0])), 9))
}
func nox_xxx_netSendComplexObject_518960(a1 int32, a2 *uint32, a3 int32) int32 {
	var (
		v3  int32
		v4  *uint32
		v5  *int16
		v6  int32
		v7  int16
		v8  int16
		v9  int16
		v10 int16
		v11 float32
		v12 int16
		v13 int32
		v14 int8
		v15 int8
		a2a int2
		v18 [11]byte
	)
	v3 = a1
	v4 = a2
	if a3 != 0 {
		v5 = (*int16)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*139)))))
		if v5 != nil {
			v6 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*187)))
			if (nox_frame_xxx_2598000 - *(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*134))) > 2 {
				v7 = *v5
				v8 = int16(*(*uint16)(unsafe.Pointer(uintptr(v6 + a1*2 + 412))))
				if int32(v7) != int32(v8) {
					nox_xxx_netReportHealthDelta_4D8760(a1, int16(uint16(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*9)))), int16(int32(v7)-int32(v8)))
					*(*uint16)(unsafe.Pointer(uintptr(v6 + v3*2 + 412))) = *(*uint16)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*139)))))
				}
			}
		}
	}
	v9 = int16(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v4))), unsafe.Sizeof(uint16(0))*2))))
	v18[0] = 48
	*(*uint16)(unsafe.Pointer(&v18[3])) = uint16(v9)
	*(*uint16)(unsafe.Pointer(&v18[1])) = uint16(nox_xxx_netGetUnitCodeServ_578AC0((*nox_object_t)(unsafe.Pointer(v4))))
	v10 = int16(int(*((*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(v4))), unsafe.Sizeof(float32(0))*14)))))
	v11 = *((*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(v4))), unsafe.Sizeof(float32(0))*15)))
	*(*uint16)(unsafe.Pointer(&v18[5])) = uint16(v10)
	v12 = int16(int(v11))
	v13 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*187)))
	*(*uint16)(unsafe.Pointer(&v18[7])) = uint16(v12)
	v18[9] = byte(int8(nox_xxx_mobActionToAnimation_533790(int32(uintptr(unsafe.Pointer(v4))))))
	v14 = int8(*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(v4))), 12))))
	v18[10] = byte(*(*uint8)(unsafe.Pointer(uintptr(v13 + 481))))
	if int32(v14)&16 != 0 && int32(*(*uint8)(unsafe.Pointer(uintptr(v13 + 2094)))) != 0 && noxRndCounter1.IntClamp(0, 10) >= 8 {
		v18[9] = 14
		nox_xxx_animPlayerGetFrameRange_4F9F90(50, (*uint32)(unsafe.Pointer(&a3)), &a1)
		v18[10] = byte(int8(noxRndCounter1.IntClamp(0, a3)))
	}
	nox_xxx_xferIndexedDirection_509E20(int32(*((*int16)(unsafe.Add(unsafe.Pointer((*int16)(unsafe.Pointer(v4))), unsafe.Sizeof(int16(0))*62)))), &a2a)
	*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a3))), 0)) = uint8(int8(int32(*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a2a.field_0))), 0))) + int32(*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a2a.field_4))), 0)))*3 + 4))
	if int32(uint8(int8(a3))) <= 3 {
		v15 = int8(a3)
	} else {
		v15 = int8(a3 - 1)
	}
	v18[9] |= byte(int8(int32(v15) * 16))
	return bool2int(nox_netlist_addToMsgListSrv_40EF40(v3, (*uint8)(unsafe.Pointer(&v18[0])), 11))
}
func nox_xxx_netSpriteUpdate_518AE0(a1 int32, a2 int32, a3 *uint32) int32 {
	var (
		v3  *uint32
		v4  int16
		v5  int32
		v7  int32
		v8  int8
		v9  int32
		v10 int8
		v11 uint8
		v12 int32
		v13 int32
		v14 uint8
		v15 int8
	)
	v3 = a3
	v4 = int16(uint16(nox_xxx_netGetUnitCodeServ_578AC0((*nox_object_t)(unsafe.Pointer(a3)))))
	v5 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*2)))
	*(*uint16)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(&a3))), 1)))) = uint16(v4)
	if uint32(v5)&0x400000 != 0 && *(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*3))&24 != 0 {
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a3))), unsafe.Sizeof((*uint32)(nil))-1)) = *(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*187)))))
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a3))), 0)) = 179
		return nox_xxx_netSendPacket1_4E5390(a2, int32(uintptr(unsafe.Pointer(&a3))), 4, 0, 1)
	}
	v7 = int32(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v3))), unsafe.Sizeof(uint16(0))*2))))
	if uint32(uint16(int16(v7))) == *memmap.PtrUint32(6112660, 2386972) {
		v8 = int8(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*187)) + 20))))
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a3))), 0)) = 57
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a3))), unsafe.Sizeof((*uint32)(nil))-1)) = uint8(v8)
		return nox_netlist_addToMsgListCli_40EBC0(a2, 1, (*uint8)(unsafe.Pointer(&a3)), 4)
	}
	if uint32(v7) == *memmap.PtrUint32(6112660, 2386980) {
		v9 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*4)))
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a3))), 0)) = 57
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a3))), unsafe.Sizeof((*uint32)(nil))-1)) = uint8(int8(int32(^(*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v9))), unsafe.Sizeof(int32(0))-1)))) & 1))
		return nox_netlist_addToMsgListCli_40EBC0(a2, 1, (*uint8)(unsafe.Pointer(&a3)), 4)
	}
	if uint32(v7) == *memmap.PtrUint32(6112660, 2386976) {
		v10 = int8(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*187))))))
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a3))), 0)) = 180
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a3))), unsafe.Sizeof((*uint32)(nil))-1)) = uint8(int8(int32(v10) & 1))
		return nox_netlist_addToMsgListCli_40EBC0(a2, 1, (*uint8)(unsafe.Pointer(&a3)), 4)
	}
	if v5&0x4000 != 0 {
		v11 = *(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*187)) + 16)))
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a3))), 0)) = 57
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a3))), unsafe.Sizeof((*uint32)(nil))-1)) = uint8(int8(int32(v11) >> 2))
		return nox_netlist_addToMsgListCli_40EBC0(a2, 1, (*uint8)(unsafe.Pointer(&a3)), 4)
	}
	if (v5 & 0x8000) != 0 {
		v12 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*187)))
		if v12 != 0 && (func() int32 {
			v13 = int32(*(*uint32)(unsafe.Pointer(uintptr(v12 + 4))))
			return v13
		}()) != 0 {
			v14 = *(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v13 + 748))) + 16)))
			*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a3))), 0)) = 57
			*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a3))), unsafe.Sizeof((*uint32)(nil))-1)) = uint8(int8(int32(v14) >> 2))
		} else {
			*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a3))), unsafe.Sizeof((*uint32)(nil))-1)) = 0
			*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a3))), 0)) = 57
		}
		return nox_netlist_addToMsgListCli_40EBC0(a2, 1, (*uint8)(unsafe.Pointer(&a3)), 4)
	}
	if (v5 & 128) != 0 {
		v15 = int8(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*187)) + 12))))
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a3))), 0)) = 178
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a3))), unsafe.Sizeof((*uint32)(nil))-1)) = uint8(v15)
		return nox_netlist_addToMsgListCli_40EBC0(a2, 1, (*uint8)(unsafe.Pointer(&a3)), 4)
	}
	nox_xxx_netUpdateObjectSpecial_527E50(a1, v3)
	return 0
}
func nox_xxx_netPlayerObjSend_518C30(a1 int32, a2 *uint32, a3 int32, a4 int32) int32 {
	var (
		v4     int32
		v5     *uint32
		v6     int32
		v7     int32
		v8     *int16
		v9     int16
		v10    int16
		v11    int16
		v12    int16
		v13    float32
		v14    int16
		v15    float32
		v16    int8
		v17    int8
		v18    int32
		v19    int8
		result int32
		a2a    int2
		v22    [12]byte
		v23    int32
	)
	v4 = a1
	v5 = a2
	v23 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
	v6 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*187)))
	v7 = int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v23 + 276))) + 2064))))
	if a4 != 0 {
		v8 = (*int16)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*139)))))
		if v8 != nil {
			if (nox_frame_xxx_2598000 - *(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*134))) > 2 {
				v9 = *v8
				v10 = int16(*(*uint16)(unsafe.Pointer(uintptr(v6 + v7*2 + 12))))
				if int32(v9) != int32(v10) {
					nox_xxx_netReportHealthDelta_4D8760(v7, int16(uint16(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*9)))), int16(int32(v9)-int32(v10)))
					*(*uint16)(unsafe.Pointer(uintptr(v6 + v7*2 + 12))) = *(*uint16)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(uint32(0))*139)))))
				}
			}
		}
	}
	if (*uint32)(unsafe.Pointer(uintptr(v4))) == v5 {
		nox_xxx_playerReportAnything_4D9900(v4)
	}
	*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v23 + 276))) + uint32(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v6 + 276))) + 2064))))*4) + 4452))) = nox_frame_xxx_2598000
	v11 = int16(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v5))), unsafe.Sizeof(uint16(0))*2))))
	v22[0] = 195
	*(*uint16)(unsafe.Pointer(&v22[3])) = uint16(v11)
	v12 = int16(uint16(nox_xxx_netGetUnitCodeServ_578AC0((*nox_object_t)(unsafe.Pointer(v5)))))
	v13 = *((*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(v5))), unsafe.Sizeof(float32(0))*14)))
	*(*uint16)(unsafe.Pointer(&v22[1])) = uint16(v12)
	v14 = int16(int(v13))
	v15 = *((*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(v5))), unsafe.Sizeof(float32(0))*15)))
	*(*uint16)(unsafe.Pointer(&v22[5])) = uint16(v14)
	*(*uint16)(unsafe.Pointer(&v22[7])) = uint16(int16(int(v15)))
	v22[9] = 0
	v22[11] = byte(int8(nox_common_mapPlrActionToStateId_4FA2B0(int32(uintptr(unsafe.Pointer(v5))))))
	v16 = int8(*(*uint8)(unsafe.Pointer(uintptr(v6 + 88))))
	if int32(v16) == 1 || int32(v16) == 10 || int32(v16) == 2 || int32(v16) == 15 || int32(v16) == 16 || int32(v16) == 17 || int32(v16) == 14 || int32(v16) == 20 || int32(v16) == 18 || int32(v16) == 19 || int32(v16) == 21 || int32(v16) == 22 || int32(v16) == 24 || int32(v16) == 25 || int32(v16) == 27 || int32(v16) == 28 || int32(v16) == 29 || int32(v16) == 26 || int32(v16) == 30 || int32(v16) == 32 {
		v22[10] = byte(*(*uint8)(unsafe.Pointer(uintptr(v6 + 236))))
	} else {
		v22[10] = math.MaxUint8
	}
	if int32(*(*uint16)(unsafe.Pointer(uintptr(v6 + 160)))) != 0 && noxRndCounter1.IntClamp(0, 10) >= 8 {
		v22[11] = 50
		v22[10] = math.MaxUint8
	}
	v17 = int8(*(*uint8)(unsafe.Pointer(uintptr(v6 + 88))))
	if (int32(v17) == 3 || int32(v17) == 4) && *(*uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(uint32(0))*131)) == 16 {
		v22[11] = 51
		v22[10] = math.MaxUint8
	}
	if int32(*(*uint8)(unsafe.Pointer(uintptr(v6 + 88)))) != 30 && *(*uint32)(unsafe.Pointer(uintptr(v6 + 164))) != 0 {
		nox_xxx_animPlayerGetFrameRange_4F9F90(52, (*uint32)(unsafe.Pointer(&a4)), (*int32)(unsafe.Pointer(&a2)))
		v18 = int32((nox_frame_xxx_2598000 - *(*uint32)(unsafe.Pointer(uintptr(v6 + 164)))) / (uint32(uintptr(unsafe.Pointer(a2))) + 1))
		if v18 >= a4 || (nox_frame_xxx_2598000-*(*uint32)(unsafe.Pointer(uintptr(v6 + 164)))) >= 4 {
			*(*uint32)(unsafe.Pointer(uintptr(v6 + 164))) = 0
		} else {
			v22[11] = 52
			v22[10] = byte(int8(v18))
		}
	}
	nox_xxx_xferIndexedDirection_509E20(int32(*((*int16)(unsafe.Add(unsafe.Pointer((*int16)(unsafe.Pointer(v5))), unsafe.Sizeof(int16(0))*62)))), &a2a)
	*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a2))), 0)) = uint8(int8(int32(*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a2a.field_0))), 0))) + int32(*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a2a.field_4))), 0)))*3 + 4))
	if int32(uint8(uintptr(unsafe.Pointer(a2)))) <= 3 {
		v19 = int8(uintptr(unsafe.Pointer(a2)))
	} else {
		v19 = int8(int32(uint8(uintptr(unsafe.Pointer(a2)))) - 1)
	}
	v22[9] |= byte(int8(int32(v19) * 16))
	if a3 != 0 {
		result = bool2int(nox_netlist_addToMsgListSrv_40EF40(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v23 + 276))) + 2064)))), (*uint8)(unsafe.Pointer(&v22[0])), 12))
	} else {
		result = nox_netlist_addToMsgListCli_40EBC0(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v23 + 276))) + 2064)))), 1, (*uint8)(unsafe.Pointer(&v22[0])), 12)
	}
	return result
}
func nox_xxx_netUpdate_518EE0(obj *nox_object_t) int32 {
	var (
		v1     *uint32
		v2     int32
		v3     int32
		result int32
		v5     int32
		v6     int32
		v7     int32
		v8     int32
		v9     int32
		v10    int32
		v11    int32
		v12    int32
		v14    int32
		v17    [3]byte
		v18    int32
		v19    [5]byte
		a1     float4
	)
	v1 = (*uint32)(unsafe.Pointer(obj))
	v2 = int32(uintptr(obj.data_update))
	v3 = int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 276))) + 2064))))
	nox_netlist_initPlayerBufs_40F020(v3)
	if v3 != 31 && ((nox_frame_xxx_2598000+uint32(v3))%(nox_gameFPS*15)) == 0 {
		nox_xxx_netReportUnitHeight_4D9020(v3, int32(uintptr(unsafe.Pointer(obj))))
	}
	if dword_5d4594_2650652 == 0 || (nox_frame_xxx_2598000%uint32(nox_xxx_rateGet_40A6C0())) == 0 || noxflags.HasGame(noxflags.GameFlag4) {
		if int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 276))) + 3680))))&64 != 0 {
			v19[0] = 40
			*(*uint32)(unsafe.Pointer(&v19[1])) = nox_frame_xxx_2598000
			nox_netlist_addToMsgListSrv_40EF40(v3, (*uint8)(unsafe.Pointer(&v19[0])), 5)
			nox_xxx_playerUnsetStatus_417530((*nox_playerInfo)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 276)))))), 64)
		} else {
			v17[0] = 39
			*(*uint16)(unsafe.Pointer(&v17[1])) = uint16(nox_frame_xxx_2598000)
			nox_netlist_addToMsgListSrv_40EF40(v3, (*uint8)(unsafe.Pointer(&v17[0])), 3)
		}
	}
	if dword_5d4594_2650652 == 0 || (*uint32)(unsafe.Pointer(obj)) == *(**uint32)(unsafe.Pointer(&nox_xxx_host_player_unit_3843628)) || noxflags.HasGame(noxflags.GameFlag4) || (nox_frame_xxx_2598000%uint32(nox_xxx_rateGet_40A6C0())) == 0 {
		if int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 276))) + 3680))))&3 != 0 || nox_common_getEngineFlag(nox_engine_flag(NOX_ENGINE_FLAG_REPLAY_READ)) {
			result = nox_xxx_netPlayerObjSendCamera_519330(int32(uintptr(unsafe.Pointer(obj))))
			if result == 0 {
				return result
			}
			if nox_common_getEngineFlag(nox_engine_flag(NOX_ENGINE_FLAG_REPLAY_READ)) {
				nox_xxx_netPlayerObjSend_518C30(int32(uintptr(unsafe.Pointer(obj))), (*uint32)(unsafe.Pointer(obj)), 1, 1)
			}
		} else {
			result = nox_xxx_netPlayerObjSend_518C30(int32(uintptr(unsafe.Pointer(obj))), (*uint32)(unsafe.Pointer(obj)), 1, 1)
			if result == 0 {
				return result
			}
		}
		v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 276))))
		v18 = int32(*(*uint16)(unsafe.Pointer(uintptr(v5 + 10))))
		a1.field_0 = float32(float64(*(*float32)(unsafe.Pointer(uintptr(v5 + 3632)))) - float64(v18) - 100.0)
		v6 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 276))))
		v18 = int32(*(*uint16)(unsafe.Pointer(uintptr(v6 + 12))))
		a1.field_4 = float32(float64(*(*float32)(unsafe.Pointer(uintptr(v6 + 3636)))) - float64(v18) - 100.0)
		v7 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 276))))
		v18 = int32(*(*uint16)(unsafe.Pointer(uintptr(v7 + 10))))
		a1.field_8 = float32(float64(v18) + float64(*(*float32)(unsafe.Pointer(uintptr(v7 + 3632)))) + 100.0)
		v8 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 276))))
		v18 = int32(*(*uint16)(unsafe.Pointer(uintptr(v8 + 12))))
		a1.field_C = float32(float64(v18) + float64(*(*float32)(unsafe.Pointer(uintptr(v8 + 3636)))) + 100.0)
		nox_xxx_getUnitsInRectAdv_517ED0(&a1, func(arg1 *float32, arg2 int32) {
			nox_xxx_unitAroundPlayerFn_5193B0((*uint32)(unsafe.Pointer(arg1)), arg2)
		}, int32(uintptr(unsafe.Pointer(obj))))
		v9 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 276))))
		v18 = int32(*(*uint16)(unsafe.Pointer(uintptr(v9 + 10))))
		a1.field_0 = float32(float64(*(*float32)(unsafe.Pointer(uintptr(v9 + 3632)))) - float64(v18) - 128.0)
		v10 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 276))))
		v18 = int32(*(*uint16)(unsafe.Pointer(uintptr(v10 + 12))))
		a1.field_4 = float32(float64(*(*float32)(unsafe.Pointer(uintptr(v10 + 3636)))) - float64(v18) - 128.0)
		v11 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 276))))
		v18 = int32(*(*uint16)(unsafe.Pointer(uintptr(v11 + 10))))
		a1.field_8 = float32(float64(v18) + float64(*(*float32)(unsafe.Pointer(uintptr(v11 + 3632)))) + 128.0)
		v12 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 276))))
		v18 = int32(*(*uint16)(unsafe.Pointer(uintptr(v12 + 12))))
		a1.field_C = float32(float64(v18) + float64(*(*float32)(unsafe.Pointer(uintptr(v12 + 3636)))) + 128.0)
		nox_xxx_secretWallCheckUnits_517F00(&a1.field_0, func(arg1 *int32, arg2 int32) int32 {
			return sub_519660(int32(uintptr(unsafe.Pointer(arg1))), arg2)
		}, int32(uintptr(unsafe.Pointer(obj))))
		if sub_519710(v2) != 0 {
			sub_519760(int32(uintptr(unsafe.Pointer(obj))), &a1.field_0)
		}
		if int32(uint8(nox_frame_xxx_2598000))&8 != 0 {
			v14 = 1 << int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 276))) + 2064))))
			for cur_obj := int32(int32(uintptr(unsafe.Pointer(noxServer.firstServerObject())))); cur_obj != 0; cur_obj = int32(uintptr(unsafe.Pointer((*nox_object_t)(unsafe.Pointer(uintptr(cur_obj))).Next()))) {
				if (*(*uint32)(unsafe.Pointer(uintptr(cur_obj + 8)))&0x20400000) == 0 && nox_xxx_playerMapTracksObj_4173D0(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 276))) + 2064)))), cur_obj) == 0 && (float64(*(*float32)(unsafe.Pointer(uintptr(cur_obj + 232)))) > float64(a1.field_8) || float64(*(*float32)(unsafe.Pointer(uintptr(cur_obj + 240)))) < float64(a1.field_0) || float64(*(*float32)(unsafe.Pointer(uintptr(cur_obj + 236)))) > float64(a1.field_C) || float64(*(*float32)(unsafe.Pointer(uintptr(cur_obj + 244)))) < float64(a1.field_4)) {
					if uint32(v14)&*(*uint32)(unsafe.Pointer(uintptr(cur_obj + 148))) != 0 {
						nox_xxx_netObjectOutOfSight_528A60(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 276))) + 2064)))), (*uint32)(unsafe.Pointer(uintptr(cur_obj))))
						*(*uint32)(unsafe.Pointer(uintptr(cur_obj + 152))) |= uint32(v14)
						*(*uint32)(unsafe.Pointer(uintptr(cur_obj + 148))) &= uint32(^v14)
					}
				}
			}
			for cur_obj := (*nox_object_t)(firstServerObjectUpdatable2()); cur_obj != nil; cur_obj = cur_obj.Next() {
				if (cur_obj.obj_class&0x20400000) == 0 && nox_xxx_playerMapTracksObj_4173D0(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 276))) + 2064)))), int32(uintptr(unsafe.Pointer(cur_obj)))) == 0 && (float64(cur_obj.collide_x1) > float64(a1.field_8) || float64(cur_obj.collide_x2) < float64(a1.field_0) || float64(cur_obj.collide_y1) > float64(a1.field_C) || float64(cur_obj.collide_y2) < float64(a1.field_4)) {
					if uint32(v14)&cur_obj.field_37 != 0 {
						nox_xxx_netObjectOutOfSight_528A60(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 276))) + 2064)))), (*uint32)(unsafe.Pointer(cur_obj)))
						*(*uint32)(unsafe.Pointer(&cur_obj.field_38)) |= uint32(v14)
						cur_obj.field_37 &= uint32(^v14)
					}
				}
			}
			v1 = (*uint32)(unsafe.Pointer(obj))
		}
	}
	if dword_5d4594_2650652 == 0 || (nox_frame_xxx_2598000%uint32(nox_xxx_rateGet_40A6C0())) == 0 || noxflags.HasGame(noxflags.GameFlag4) {
		sub_4FF7B0(int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 276)))))
		sub_511100(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 276))) + 2064)))))
	}
	return nox_xxx_netUpdateRemotePlr_501CA0(int32(uintptr(unsafe.Pointer(v1))))
}
func nox_xxx_netPlayerObjSendCamera_519330(a1 int32) int32 {
	var (
		v1 int32
		v3 [12]byte
	)
	v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
	v3[0] = 195
	*(*uint16)(unsafe.Pointer(&v3[3])) = 0
	*(*uint16)(unsafe.Pointer(&v3[1])) = 0
	*(*uint16)(unsafe.Pointer(&v3[5])) = uint16(int16(int64(*(*float32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v1 + 276))) + 3632))))))
	*(*uint16)(unsafe.Pointer(&v3[7])) = uint16(int16(int64(*(*float32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v1 + 276))) + 3636))))))
	v3[9] = 0
	v3[10] = math.MaxUint8
	v3[11] = 0
	return bool2int(nox_netlist_addToMsgListSrv_40EF40(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v1 + 276))) + 2064)))), (*uint8)(unsafe.Pointer(&v3[0])), 12))
}
func nox_xxx_unitAroundPlayerFn_5193B0(a1 *uint32, a2 int32) int8 {
	var (
		v2 int32
		v3 int32
	)
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 748))))
	if (*uint32)(unsafe.Pointer(uintptr(a2))) != a1 || (func() int32 {
		nox_xxx_netUpdateObjectSpecial_527E50(a2, a1)
		v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 276))))
		return int32(*(*uint8)(unsafe.Pointer(uintptr(v3 + 3680)))) & 1
	}()) != 0 {
		if !noxflags.HasGame(noxflags.GameOnline) || (func() bool {
			*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v3))), 0)) = uint8(nox_frame_xxx_2598000)
			return *(*uint32)(unsafe.Pointer(uintptr(v2 + 272))) != nox_frame_xxx_2598000
		}()) {
			*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v3))), 0)) = uint8(nox_xxx_netSendObjects2Plr_519410(a2, int32(uintptr(unsafe.Pointer(a1)))))
		}
	}
	return int8(v3)
}
func nox_xxx_netSendObjects2Plr_519410(a1 int32, a2 int32) int8 {
	var (
		v2  int32
		v3  int32
		v4  int32
		v5  int32
		v6  int32
		v7  int32
		v8  int32
		v9  float32
		v10 float32
		v11 int32
		v13 int32
		v14 float4
	)
	v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
	v13 = v4
	v2 = 0
	v3 = int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v4 + 276))) + 2064))))
	*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v4))), 0)) = *(*uint8)(unsafe.Pointer(uintptr(a2 + 16)))
	v5 = v3
	if (v4&32) == 0 && (*(*uint32)(unsafe.Pointer(uintptr(a2 + 8)))&0x40000000) == 0 {
		v6 = 1 << v3
		if a2 != a1 && uint32(v6)&*(*uint32)(unsafe.Pointer(uintptr(a2 + 140))) != 0 {
			nox_xxx_netFriendAddRemove_4D97A0(v3, (*uint32)(unsafe.Pointer(uintptr(a2))), bool2int((uint32(v6)&*(*uint32)(unsafe.Pointer(uintptr(a2 + 144)))) != 0))
			*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v4))), 0)) = ^uint8(int8(v6))
			*(*uint32)(unsafe.Pointer(uintptr(a2 + 140))) &= uint32(^v6)
		}
		if uint32(v6)&*(*uint32)(unsafe.Pointer(uintptr(a2 + 152))) != 0 || int32(*(*uint8)(unsafe.Pointer(uintptr(a2 + 8))))&1 != 0 {
			v7 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 16))))
			if (v7&2048) == 0 || nox_xxx_unitHasThatParent_4EC4F0(a2, a1) != 0 || (func() int32 {
				v4 = bool2int(noxflags.HasGame(noxflags.GameModeSolo10))
				return v4
			}()) != 0 {
				v8 = int32(*(*uint32)(unsafe.Pointer(uintptr(v13 + 276))))
				v14.field_0 = *(*float32)(unsafe.Pointer(uintptr(v8 + 3632)))
				v9 = *(*float32)(unsafe.Pointer(uintptr(a2 + 60)))
				v10 = *(*float32)(unsafe.Pointer(uintptr(v8 + 3636)))
				v14.field_8 = *(*float32)(unsafe.Pointer(uintptr(a2 + 56)))
				v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 8))))
				v14.field_4 = v10
				v14.field_C = v9
				if uint32(v4)&0x20400000 != 0 || (func() int32 {
					v4 = nox_xxx_playerMapTracksObj_4173D0(v5, a2)
					return v4
				}()) != 0 || (func() int32 {
					v4 = nox_xxx_mapTraceRay_535250(&v14, nil, nil, 69)
					return v4
				}()) != 0 {
					if uint32(v6)&*(*uint32)(unsafe.Pointer(uintptr(a2 + 148))) != 0 {
						if int32(*(*uint8)(unsafe.Pointer(uintptr(a2 + 20))))&32 != 0 {
							return int8(v4)
						}
					} else if *(*uint32)(unsafe.Pointer(uintptr(a2 + v5*4 + 560)))&4095 != 0 {
						*(*uint32)(unsafe.Pointer(uintptr(a2 + v5*4 + 560))) |= (*(*uint32)(unsafe.Pointer(uintptr(a2 + v5*4 + 560))) & 4095) << 16
					}
					v11 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 8))))
					if uint32(v11)&0x400000 != 0 {
						v2 = nox_xxx_netSpriteUpdate_518AE0(a1, v5, (*uint32)(unsafe.Pointer(uintptr(a2))))
					} else if uint32(v11)&0x200000 != 0 {
						if v11&2 != 0 {
							v2 = nox_xxx_netSendComplexObject_518960(v5, (*uint32)(unsafe.Pointer(uintptr(a2))), 1)
						} else if v11&4 != 0 {
							v2 = nox_xxx_netPlayerObjSend_518C30(a1, (*uint32)(unsafe.Pointer(uintptr(a2))), 1, 1)
						} else {
							v2 = nox_xxx_netSendPhantomPlrMb_5187E0(v5, a2)
						}
					} else if uint32(v11)&0x100000 != 0 {
						v2 = nox_xxx_netSendSimpleObj_5188A0(v5, a2)
					} else {
						*(*uint32)(unsafe.Pointer(uintptr(a2 + 152))) = 0
					}
					*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v4))), 0)) = uint8(int8(nox_xxx_netUpdateObjectSpecial_527E50(a1, (*uint32)(unsafe.Pointer(uintptr(a2))))))
					if v2 != 0 {
						v4 = int32(uint32(v6) | *(*uint32)(unsafe.Pointer(uintptr(a2 + 148))))
						*(*uint32)(unsafe.Pointer(uintptr(a2 + 152))) &= uint32(^v6)
						*(*uint32)(unsafe.Pointer(uintptr(a2 + 148))) = uint32(v4)
					}
				} else if uint32(v6)&*(*uint32)(unsafe.Pointer(uintptr(a2 + 148))) != 0 {
					if int32(*(*uint8)(unsafe.Pointer(uintptr(a2 + 8))))&6 != 0 {
						nox_xxx_netObjectOutOfSight_528A60(v5, (*uint32)(unsafe.Pointer(uintptr(a2))))
					} else {
						nox_xxx_netObjectInShadows_528A90(v5, (*uint32)(unsafe.Pointer(uintptr(a2))))
					}
					v4 = int32(uint32(^v6) & *(*uint32)(unsafe.Pointer(uintptr(a2 + 148))))
					*(*uint32)(unsafe.Pointer(uintptr(a2 + 152))) |= uint32(v6)
					*(*uint32)(unsafe.Pointer(uintptr(a2 + 148))) = uint32(v4)
				}
			}
		}
	}
	return int8(v4)
}
func sub_519660(a1 int32, a2 int32) int32 {
	var (
		v2     int32
		result int32
		v4     int32
		v5     int32
	)
	v2 = 1 << int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a2 + 748))) + 276))) + 2064))))
	result = bool2int((uint32(v2) & *(*uint32)(unsafe.Pointer(uintptr(a1 + 28)))) != 0)
	switch *(*uint8)(unsafe.Pointer(uintptr(a1 + 21))) {
	case 1:
		fallthrough
	case 4:
		v4 = 0
	case 2:
		fallthrough
	case 3:
		v4 = 1
	default:
		v4 = a2
	}
	if result != v4 {
		v5 = int32(uintptr(nox_server_getWallAtGrid_410580(int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 4)))), int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 8)))))))
		if v4 != 0 {
			sub_4DF120(v5)
			result = int32(uint32(v2) | *(*uint32)(unsafe.Pointer(uintptr(a1 + 28))))
		} else {
			sub_4DF180(v5)
			result = int32(uint32(^v2) & *(*uint32)(unsafe.Pointer(uintptr(a1 + 28))))
		}
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 28))) = uint32(result)
	}
	return result
}
func sub_519710(a1 int32) int32 {
	var result int32
	result = sub_417270(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 276))) + 2064)))))
	if result != 0 {
		result = bool2int(result > 60 || nox_frame_xxx_2598000-*(*uint32)(unsafe.Pointer(uintptr(a1 + 268))) > uint32(60/result))
	}
	return result
}
func sub_519760(a1 int32, a2 *float32) {
	var (
		v2 int32
		v3 int32
		v4 int32
		v5 int32
	)
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
	v3 = int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 276))) + 2064))))
	v4 = sub_4172C0(v3)
	v5 = v4
	if v4 != 0 {
		if int32(*(*uint8)(unsafe.Pointer(uintptr(v4 + 16))))&32 != 0 {
			nox_xxx_netMinimapUnmark4All_417430(v4)
		} else if float64(*(*float32)(unsafe.Pointer(uintptr(v4 + 56)))) < float64(*a2) || float64(*(*float32)(unsafe.Pointer(uintptr(v4 + 56)))) > float64(*(*float32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(float32(0))*2))) || float64(*(*float32)(unsafe.Pointer(uintptr(v4 + 60)))) < float64(*(*float32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(float32(0))*1))) || float64(*(*float32)(unsafe.Pointer(uintptr(v4 + 60)))) > float64(*(*float32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(float32(0))*3))) {
			*(*uint32)(unsafe.Pointer(uintptr(v4 + 152))) |= uint32(1 << v3)
			nox_xxx_netSendObjects2Plr_519410(a1, v4)
			nox_xxx_netReportUnitHeight_4D9020(v3, v5)
			*(*uint32)(unsafe.Pointer(uintptr(v2 + 268))) = nox_frame_xxx_2598000
		}
	}
}
func nox_xxx_netMapSendClear_519830(a1 int32, a2 int8) int32 {
	var result int32
	result = a1
	if a1 != 0 {
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 40))) = 0
		*(*uint8)(unsafe.Pointer(uintptr(a1))) = uint8(a2)
		*(*uint16)(unsafe.Pointer(uintptr(a1 + 2))) = 0
		*(*uint16)(unsafe.Pointer(uintptr(a1 + 4))) = 0
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 8))) = 0
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 12))) = 0
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 16))) = 1
		*(*uint16)(unsafe.Pointer(uintptr(a1 + 20))) = 512
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 24))) = 0
		*(*uint8)(unsafe.Pointer(uintptr(a1 + 28))) = 2
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 32))) = 0
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 44))) = 0
	}
	return result
}
func nox_xxx_netMapSendStop_519870() int32 {
	var (
		v0     int8
		v1     *uint8
		result int32
	)
	v0 = 0
	v1 = (*uint8)(memmap.PtrOff(6112660, 2387148))
	*memmap.PtrUint16(6112660, 2388636) = 0
	*memmap.PtrUint16(6112660, 2388638) = 0
	dword_5d4594_2388640 = 0
	dword_5d4594_2388648 = 0
	for {
		result = nox_xxx_netMapSendClear_519830(int32(uintptr(unsafe.Pointer(v1))), v0)
		v1 = (*uint8)(unsafe.Add(unsafe.Pointer(v1), 48))
		v0++
		if int32(uintptr(unsafe.Pointer(v1))) >= int32(uintptr(memmap.PtrOff(6112660, 2388636))) {
			break
		}
	}
	return result
}
func nox_xxx_mapSendCancelAll_5198B0(a1 uint8) {
	var (
		v1 int8
		v2 *uint8
	)
	v1 = 0
	v2 = (*uint8)(memmap.PtrOff(6112660, 2387150))
	for {
		if int32(*(*uint16)(unsafe.Pointer(v2))) == 1 {
			*(*uint16)(unsafe.Pointer(v2)) = 0
			nullsub_27(uint32(uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v2), -2))))))
			nox_xxx_netSendMapAbort_519C80_net_mapsend((*uint8)(unsafe.Add(unsafe.Pointer(v2), -2)), a1)
			*memmap.PtrUint16(6112660, 2388638)++
			if int32(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v2))), unsafe.Sizeof(uint16(0))*1)))) == 1 {
				if *(*uint32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v2), 6)))) != 0 {
					*(*unsafe.Pointer)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v2), 6)))) = nil
				}
			}
			nox_xxx_netMapSendClear_519830(int32(uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v2), -2))))), v1)
		}
		v2 = (*uint8)(unsafe.Add(unsafe.Pointer(v2), 48))
		v1++
		if int32(uintptr(unsafe.Pointer(v2))) >= int32(uintptr(memmap.PtrOff(6112660, 2388638))) {
			break
		}
	}
}
func sub_519930(a1 int32) int32 {
	var (
		v1 int32
		v2 *uint8
		i  *byte
		v4 int8
		v6 int32
	)
	v1 = 0
	v6 = 0
	v2 = (*uint8)(memmap.PtrOff(6112660, uint32(a1*48)+0x246CCC))
	if v2 != nil {
		if a1 < 32 {
			for i = (*byte)(unsafe.Pointer(uintptr(sub_555250(uint32(*v2), (*uint32)(unsafe.Pointer(&v6)))))); i != nil; i = (*byte)(unsafe.Pointer(uintptr(sub_555290(uint32(*v2), (*uint32)(unsafe.Pointer(&v6)))))) {
				v4 = int8(*i)
				if int32(v4) == -72 || int32(v4) == -71 {
					v1++
				}
			}
		}
	}
	return v1
}
func sub_519E80(a1 int32) *uint32 {
	nox_common_playerInfoFromNum_417090(a1)
	return sub_4DE410(int32(*memmap.PtrUint8(6112660, uint32(a1*48)+0x246CCC)))
}
func sub_51A100() unsafe.Pointer {
	var result unsafe.Pointer
	if dword_5d4594_2388648 != 0 {
		nox_xxx_mapSendCancelAll_5198B0(0)
	}
	result = *(*unsafe.Pointer)(unsafe.Pointer(&dword_5d4594_2388640))
	if dword_5d4594_2388640 != 0 {
		*(*unsafe.Pointer)(unsafe.Pointer(&dword_5d4594_2388640)) = nil
	}
	*memmap.PtrUint32(6112660, 2388644) = 0
	return result
}
func sub_51A130() int32 {
	return int32(dword_5d4594_2388648)
}
func sub_51A140(a1 int32) int32 {
	var v1 *uint8
	v1 = (*uint8)(memmap.PtrOff(6112660, uint32(a1*48)+0x246CCC))
	return bool2int(v1 != nil && a1 < 32 && int32(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v1))), unsafe.Sizeof(uint16(0))*1)))) == 1)
}
func sub_51A170(a1 *byte) int32 {
	var (
		v1 int32
		v2 **byte
	)
	if a1 == nil {
		return -1
	}
	v1 = 0
	v2 = (**byte)(memmap.PtrOff(0x587000, 249880))
	for libc.StrCmp(*v2, a1) != 0 {
		v2 = (**byte)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof((*byte)(nil))*1))
		v1++
		if int32(uintptr(unsafe.Pointer(v2))) >= int32(uintptr(memmap.PtrOff(0x587000, 249892))) {
			return -1
		}
	}
	return v1
}
func sub_51A1D0(a1 int32) *byte {
	var result *byte
	if a1 <= -1 || a1 >= 3 {
		result = nil
	} else {
		result = *(**byte)(memmap.PtrOff(0x587000, uint32(a1*4)+249880))
	}
	return result
}
func sub_51A500(a1 int32) int32 {
	var (
		v1 int32
		v2 int32
		i  *uint8
		v4 int32
	)
	if *memmap.PtrUint32(6112660, 2388664) == 0 {
		sub_51A550()
	}
	if a1 == 0 {
		return 0
	}
	v1 = 0
	if *memmap.PtrUint32(0x587000, 249904) == 0 {
		return 0
	}
	*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v2))), unsafe.Sizeof(uint16(0))*1)) = 0
	for i = (*uint8)(memmap.PtrOff(0x587000, 249904)); ; i = (*uint8)(unsafe.Add(unsafe.Pointer(i), 16)) {
		*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v2))), unsafe.Sizeof(uint16(0))*0)) = *(*uint16)(unsafe.Pointer(uintptr(a1 + 4)))
		if *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(i))), -int(unsafe.Sizeof(uint32(0))*1)))) == uint32(v2) {
			break
		}
		v4 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(i))), unsafe.Sizeof(uint32(0))*4))))
		v1++
		if v4 == 0 {
			return 0
		}
	}
	return int32(*memmap.PtrUint32(0x587000, uint32(v1*16)+0x3D034))
}
func sub_51A550() *byte {
	var (
		result *byte
		v1     *uint8
	)
	result = *(**byte)(memmap.PtrOff(0x587000, 249904))
	if *memmap.PtrUint32(0x587000, 249904) != 0 {
		v1 = (*uint8)(memmap.PtrOff(0x587000, 249896))
		for {
			*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v1))), unsafe.Sizeof(uint32(0))*3))) = uint32(nox_xxx_getNameId_4E3AA0(result))
			*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v1))), unsafe.Sizeof(uint32(0))*1))) = uint32(nox_xxx_getNameId_4E3AA0(*(**byte)(unsafe.Pointer(v1))))
			result = (*byte)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v1))), unsafe.Sizeof(uint32(0))*6))))))
			v1 = (*uint8)(unsafe.Add(unsafe.Pointer(v1), 16))
			if result == nil {
				break
			}
		}
		*memmap.PtrUint32(6112660, 2388664) = 1
	} else {
		*memmap.PtrUint32(6112660, 2388664) = 1
	}
	return result
}
func nox_xxx_spawnHecubahQuest_51A5A0(a1 *int32) {
	var (
		v1  *uint32
		v2  *uint32
		v3  int32
		v4  int32
		v5  float64
		v6  int32
		v7  *uint32
		v8  int32
		v9  *uint32
		v10 int32
		v11 *uint32
		v12 int32
		v13 *uint32
		v14 int32
		v15 *uint32
		v16 float32
		v17 float32
	)
	v1 = (*uint32)(unsafe.Pointer(nox_xxx_newObjectByTypeID_4E3810(CString("Hecubah"))))
	v16 = float32(sub_4E40F0())
	if v1 != nil {
		v2 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*187)))))
		v3 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*121)))
		if v3 != 0 {
			v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(v3 + 72))))
		} else {
			v4 = int32(*(*uint16)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer((*nox_objectType_t)(unsafe.Add(unsafe.Pointer(nox_xxx_objectTypeByInd_4E3B70(int32(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v1))), unsafe.Sizeof(uint16(0))*2)))))), unsafe.Sizeof(nox_objectType_t{})*136)))) + 4))))
		}
		if float64(v16) < 1.0 {
			v16 = 1.0
		}
		v5 = float64(v4) * float64(v16)
		v17 = float32(v5)
		nox_xxx_unitSetHP_4E4560((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v1)))))), uint16(int16(int64(v5))))
		*(*uint16)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*139)) + 4))) = uint16(int16(int(v17)))
		if int32(*(*uint16)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*139)))))) == 0 {
			nox_xxx_unitSetHP_4E4560((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v1)))))), 1)
		}
		v6 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*139)))
		if int32(*(*uint16)(unsafe.Pointer(uintptr(v6 + 4)))) == 0 {
			*(*uint16)(unsafe.Pointer(uintptr(v6 + 4))) = 1
		}
		*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*411)) = 0x10000000
		*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*423)) = 0x10000000
		*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*340)) = 4
		*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*326)) = 0x3F547AE1
		*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*510)) = 3
		*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*410)) = 0x8000000
		*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*444)) = 0x20000000
		*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*388)) = 0x40000000
		*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*415)) = 0x40000000
		nox_xxx_gamedataGetFloat_419D40(CString("HecubahQuestSkill"))
		*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*330)) = 0x3F59999A
		nox_xxx_createAt_4DAA50((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v1)))))), nil, *(*float32)(unsafe.Pointer(a1)), *((*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(a1))), unsafe.Sizeof(float32(0))*1))))
		v7 = (*uint32)(unsafe.Pointer(nox_xxx_newObjectByTypeID_4E3810(CString("RewardMarker"))))
		if v7 != nil {
			v8 = nox_game_getQuestStage_4E3CC0()
			v9 = nox_server_rewardgen_activateMarker_4F0720(int32(uintptr(unsafe.Pointer(v7))), uint32(v8+2))
			if v9 != nil {
				nox_xxx_inventoryPutImpl_4F3070(int32(uintptr(unsafe.Pointer(v1))), (*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v9)))))), 0)
			}
			v10 = nox_game_getQuestStage_4E3CC0()
			v11 = nox_server_rewardgen_activateMarker_4F0720(int32(uintptr(unsafe.Pointer(v7))), uint32(v10+2))
			if v11 != nil {
				nox_xxx_inventoryPutImpl_4F3070(int32(uintptr(unsafe.Pointer(v1))), (*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v11)))))), 0)
			}
			v12 = nox_game_getQuestStage_4E3CC0()
			v13 = nox_server_rewardgen_activateMarker_4F0720(int32(uintptr(unsafe.Pointer(v7))), uint32(v12+2))
			if v13 != nil {
				nox_xxx_inventoryPutImpl_4F3070(int32(uintptr(unsafe.Pointer(v1))), (*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v13)))))), 0)
			}
			v14 = nox_game_getQuestStage_4E3CC0()
			v15 = nox_server_rewardgen_activateMarker_4F0720(int32(uintptr(unsafe.Pointer(v7))), uint32(v14+2))
			if v15 != nil {
				nox_xxx_inventoryPutImpl_4F3070(int32(uintptr(unsafe.Pointer(v1))), (*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v15)))))), 0)
			}
			nox_xxx_objectFreeMem_4E38A0(int32(uintptr(unsafe.Pointer(v7))))
		}
	}
}
func nox_xxx_spawnNecroQuest_51A7A0(a1 *int32) {
	var (
		v1  *uint32
		v2  *uint32
		v3  int32
		v4  int32
		v5  float64
		v6  int32
		v7  *uint32
		v8  int32
		v9  *uint32
		v10 float32
		v11 float32
	)
	v1 = (*uint32)(unsafe.Pointer(nox_xxx_newObjectByTypeID_4E3810(CString("Necromancer"))))
	v10 = float32(sub_4E40F0())
	if v1 != nil {
		v2 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*187)))))
		v3 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*121)))
		if v3 != 0 {
			v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(v3 + 72))))
		} else {
			v4 = int32(*(*uint16)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer((*nox_objectType_t)(unsafe.Add(unsafe.Pointer(nox_xxx_objectTypeByInd_4E3B70(int32(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v1))), unsafe.Sizeof(uint16(0))*2)))))), unsafe.Sizeof(nox_objectType_t{})*136)))) + 4))))
		}
		if float64(v10) < 1.0 {
			v10 = 1.0
		}
		v5 = float64(v4) * float64(v10)
		v11 = float32(v5)
		nox_xxx_unitSetHP_4E4560((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v1)))))), uint16(int16(int64(v5))))
		*(*uint16)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*139)) + 4))) = uint16(int16(int(v11)))
		if int32(*(*uint16)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*139)))))) == 0 {
			nox_xxx_unitSetHP_4E4560((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v1)))))), 1)
		}
		v6 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*139)))
		if int32(*(*uint16)(unsafe.Pointer(uintptr(v6 + 4)))) == 0 {
			*(*uint16)(unsafe.Pointer(uintptr(v6 + 4))) = 1
		}
		*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*340)) = 4
		*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*411)) = 0x10000000
		*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*423)) = 0x10000000
		*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*326)) = 0x3F547AE1
		*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*510)) = 1
		*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*410)) = 0x8000000
		*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*444)) = 0x20000000
		*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*415)) = 0x40000000
		nox_xxx_createAt_4DAA50((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v1)))))), nil, *(*float32)(unsafe.Pointer(a1)), *((*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(a1))), unsafe.Sizeof(float32(0))*1))))
		v7 = (*uint32)(unsafe.Pointer(nox_xxx_newObjectByTypeID_4E3810(CString("RewardMarker"))))
		if v7 != nil {
			v8 = nox_game_getQuestStage_4E3CC0()
			v9 = nox_server_rewardgen_activateMarker_4F0720(int32(uintptr(unsafe.Pointer(v7))), uint32(v8+2))
			if v9 != nil {
				nox_xxx_inventoryPutImpl_4F3070(int32(uintptr(unsafe.Pointer(v1))), (*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v9)))))), 0)
			}
			nox_xxx_objectFreeMem_4E38A0(int32(uintptr(unsafe.Pointer(v7))))
		}
	}
}
func nox_xxx_getQuestStage_51A930() int32 {
	return int32(*memmap.PtrUint32(6112660, 2388660))
}
func sub_51A940(a1 int32) int32 {
	var result int32
	result = a1
	*memmap.PtrUint32(6112660, 2388656) = uint32(a1)
	return result
}
func sub_51A950() int32 {
	return int32(*memmap.PtrUint32(6112660, 2388656))
}

var nox_players_controlBuffer_2388676 [32]int32 = [32]int32{}
var nox_players_controlBuffer_2388804 [32]int32 = [32]int32{}
var nox_players_controlBuffer_2388932 [32][128]nox_player_ctrl_t = [32][128]nox_player_ctrl_t{}

func sub_51AA20(a1 int32) int32 {
	var (
		v1     int32
		v2     int32
		v3     int32
		result int32
		v5     int32
		v6     *uint8
	)
	v1 = 0
	v2 = 0
	v3 = 0
	result = nox_players_controlBuffer_2388804[a1] - 1
	if result >= 0 {
		v5 = nox_players_controlBuffer_2388804[a1]
		v6 = (*uint8)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(&nox_players_controlBuffer_2388932[a1][result].field_4))))))
		for {
			if *(*uint32)(unsafe.Pointer(v6)) != 0 {
				result = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v6))), -int(unsafe.Sizeof(uint32(0))*2)))) - 2)
				if *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v6))), -int(unsafe.Sizeof(uint32(0))*2)))) == 2 {
					if v3 != 0 {
						*(*uint32)(unsafe.Pointer(v6)) = 0
					} else {
						v3 = 1
					}
				} else {
					result = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v6))), -int(unsafe.Sizeof(uint32(0))*2)))) - 4)
					if *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v6))), -int(unsafe.Sizeof(uint32(0))*2)))) == 4 {
						if v1 != 0 {
							*(*uint32)(unsafe.Pointer(v6)) = 0
						} else {
							v1 = 1
						}
					} else {
						result = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v6))), -int(unsafe.Sizeof(uint32(0))*2)))) - 5)
						if *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v6))), -int(unsafe.Sizeof(uint32(0))*2)))) == 5 {
							if v2 != 0 {
								*(*uint32)(unsafe.Pointer(v6)) = 0
							} else {
								v2 = 1
							}
						}
					}
				}
			}
			v6 = (*uint8)(unsafe.Add(unsafe.Pointer(v6), -24))
			v5--
			if v5 == 0 {
				break
			}
		}
	}
	return result
}
func nox_xxx_playerControlBufferFirst_51AB50(a1 int32) *nox_player_ctrl_t {
	nox_players_controlBuffer_2388676[a1] = 0
	if nox_players_controlBuffer_2388804[a1] <= 0 {
		return nil
	}
	for {
		var v2 int32 = nox_players_controlBuffer_2388676[a1]
		if nox_players_controlBuffer_2388932[a1][v2].field_4 != 0 {
			break
		}
		nox_players_controlBuffer_2388676[a1] = v2 + 1
		if v2+1 >= nox_players_controlBuffer_2388804[a1] {
			return nil
		}
	}
	var ind int32 = nox_players_controlBuffer_2388676[a1]
	return &nox_players_controlBuffer_2388932[a1][ind]
}
func nox_xxx_playerGetControlBufferNext_51ABC0(a1 int32) *nox_player_ctrl_t {
	var v1 int32 = nox_players_controlBuffer_2388676[a1] + 1
	nox_players_controlBuffer_2388676[a1] = v1
	var v2 int32 = v1
	if v1 >= nox_players_controlBuffer_2388804[a1] {
		return nil
	}
	for nox_players_controlBuffer_2388932[a1][v2].field_4 == 0 {
		nox_players_controlBuffer_2388676[a1] = func() int32 {
			p := &v2
			*p++
			return *p
		}()
		if v2 >= nox_players_controlBuffer_2388804[a1] {
			return nil
		}
	}
	var ind int32 = nox_players_controlBuffer_2388676[a1]
	return &nox_players_controlBuffer_2388932[a1][ind]
}
func nox_xxx_playerCmd_51AC30(a1 int32) {
	nox_players_controlBuffer_2388804[a1] = 0
}
func nox_xxx_playerCmdGet_51AC40(a1 int32) int32 {
	return bool2int(nox_players_controlBuffer_2388804[a1] == 0)
}
func sub_51B810(a1p *nox_object_t) {
	var (
		a1 int32 = int32(uintptr(unsafe.Pointer(a1p)))
		v1 *float32
		v2 float64
		v3 int32
		v4 float64
		v6 int32
		v7 float32
	)
	v1 = (*float32)(unsafe.Pointer(uintptr(a1)))
	v6 = a1
	v2 = float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 88))) + *(*float32)(unsafe.Pointer(uintptr(a1 + 80))))
	v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 68))))
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 56))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 64)))
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 60))) = uint32(v3)
	v4 = v2 * float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 112))))
	*(*float32)(unsafe.Pointer(uintptr(a1 + 80))) = float32(v4)
	v7 = (*(*float32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(float32(0))*23)) + *(*float32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(float32(0))*21))) * *(*float32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(float32(0))*28))
	*(*float32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(float32(0))*21)) = v7
	*(*float32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(float32(0))*16)) = float32(v4 + float64(*(*float32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(float32(0))*16))))
	*(*float32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(float32(0))*17)) = v7 + *(*float32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(float32(0))*17))
	nox_xxx_objectUnkUpdateCoords_4E7290(v6)
}
func sub_51B860(a1 int32) int8 {
	return nox_xxx_unitHasCollideOrUpdateFn_537610((*nox_object_t)(unsafe.Pointer(uintptr(a1))))
}
func nox_xxx_updateFallLogic_51B870(a1p *nox_object_t) {
	var (
		a1  int32 = int32(uintptr(unsafe.Pointer(a1p)))
		v1  int32
		v2  int32
		v3  float64
		v4  float64
		v5  float64
		v6  float64
		v7  float64
		v8  float64
		v9  float64
		v10 float32
		v11 float32
		v12 float32
		v13 float32
		v14 int32
	)
	v1 = a1
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 16))))
	v3 = float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 104))))
	if uint32(v2)&0x40000 != 0 {
		v10 = float32(v3 + float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 108)))))
		nox_xxx_unitRaise_4E46F0((*nox_object_t)(unsafe.Pointer(uintptr(a1))), v10)
		v4 = float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 108)))) - 1.0
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 88))) = 0
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 92))) = 0
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 80))) = 0
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 84))) = 0
		*(*float32)(unsafe.Pointer(uintptr(a1 + 108))) = float32(v4)
		v5 = float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 56))) - *(*float32)(unsafe.Pointer(uintptr(a1 + 156))))
		v6 = float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 60))) - *(*float32)(unsafe.Pointer(uintptr(a1 + 160))))
		v13 = float32(v6)
		v7 = math.Sqrt(v6*float64(v13) + v5*v5)
		*(*float32)(unsafe.Pointer(&v14)) = float32(v7)
		if v7 > 0.0 {
			*(*float32)(unsafe.Pointer(uintptr(v1 + 88))) = float32(v5 * (-3.0) / float64(*(*float32)(unsafe.Pointer(&v14))))
			*(*float32)(unsafe.Pointer(uintptr(v1 + 92))) = float32(float64(v13) * (-3.0) / float64(*(*float32)(unsafe.Pointer(&v14))))
		}
		if float64(*(*float32)(unsafe.Pointer(uintptr(v1 + 104)))) < -50.0 {
			nox_xxx_unitRaise_4E46F0((*nox_object_t)(unsafe.Pointer(uintptr(v1))), 90.0)
			*(*uint32)(unsafe.Pointer(uintptr(v1 + 16))) &= 0xFFFBFFFF
			nox_xxx_unitMove_4E7010((*nox_object_t)(unsafe.Pointer(uintptr(v1))), (*float2)(unsafe.Pointer(uintptr(v1+164))))
		}
	} else if v3 != 0.0 || float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 108)))) != 0.0 {
		if uint32(v2)&0x800000 != 0 {
			v11 = *(*float32)(unsafe.Pointer(uintptr(a1 + 104))) + *(*float32)(unsafe.Pointer(uintptr(a1 + 108)))
			nox_xxx_unitRaise_4E46F0((*nox_object_t)(unsafe.Pointer(uintptr(a1))), v11)
			if float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 104)))) >= 0.0 {
				*(*float32)(unsafe.Pointer(uintptr(a1 + 108))) = float32(float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 108)))) - 0.5)
			} else {
				nox_xxx_unitRaise_4E46F0((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0.0)
				v8 = float64(-*(*float32)(unsafe.Pointer(uintptr(a1 + 108)))**(*float32)(unsafe.Pointer(uintptr(a1 + 116)))) * 0.1
				*(*float32)(unsafe.Pointer(uintptr(a1 + 108))) = float32(v8)
				if v8 < 2.0 {
					nox_xxx_unitRaise_4E46F0((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0.0)
					*(*uint32)(unsafe.Pointer(uintptr(a1 + 108))) = 0
				}
			}
		} else if (uint32(v2) & 0x100000) == 0 {
			if float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 104)))) > 0.0 {
				if float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 108)))) <= 0.0 {
					*(*uint32)(unsafe.Pointer(uintptr(a1 + 16))) = uint32(v2) | 0x20000
				}
				v12 = *(*float32)(unsafe.Pointer(uintptr(a1 + 104))) + *(*float32)(unsafe.Pointer(uintptr(a1 + 108)))
				nox_xxx_unitRaise_4E46F0((*nox_object_t)(unsafe.Pointer(uintptr(a1))), v12)
				*(*float32)(unsafe.Pointer(uintptr(a1 + 108))) = float32(float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 108)))) - 1.0)
			}
			if float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 104)))) <= 0.0 {
				v9 = float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 108))))
				*(*uint32)(unsafe.Pointer(uintptr(a1 + 16))) &= 0xFFFDFFFF
				if v9 < 0.0 && (int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))&1) == 0 {
					nox_xxx_unitHasCollideOrUpdateFn_537610((*nox_object_t)(unsafe.Pointer(uintptr(a1))))
					if float64(*(*float32)(unsafe.Pointer(uintptr(a1 + 108)))) < -10.0 {
						if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))&4 != 0 {
							nox_xxx_aud_501960(280, (*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0, 0)
						}
					}
				}
				nox_xxx_unitRaise_4E46F0((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 0.0)
				*(*uint32)(unsafe.Pointer(uintptr(a1 + 108))) = 0
			}
		}
	}
}
func sub_51D0E0() {
	dword_5d4594_2487244 = 0
}
func sub_51D0F0(a1 int8) int32 {
	*memmap.PtrUint8(0x973F18, 35972) = uint8(a1)
	return 1
}
func sub_51D100(a1 int32) int32 {
	if a1 != 1 && a1 != 0 {
		return 0
	}
	*memmap.PtrUint32(0x973F18, 35976) = uint32(a1)
	return 1
}
func sub_51D120(a1 *float32) *uint32 {
	var (
		result *uint32
		v2     *uint32
		v3     float2
	)
	result = (*uint32)(unsafe.Pointer(uintptr(nox_xxx_mapGenFixCoords_4D3D90((*float2)(unsafe.Pointer(a1)), &v3))))
	if result != nil {
		result = nox_xxx_waypointNew_5798F0(v3.field_0, v3.field_4)
		v2 = result
		if result != nil {
			if dword_5d4594_2487244 != 0 {
				if *memmap.PtrUint32(0x973F18, 35976) == 1 {
					sub_51D300(*(*int32)(unsafe.Pointer(&dword_5d4594_2487244)), int32(uintptr(unsafe.Pointer(result))), int8(*memmap.PtrUint8(0x973F18, 35972)))
					sub_51D300(int32(uintptr(unsafe.Pointer(v2))), *(*int32)(unsafe.Pointer(&dword_5d4594_2487244)), int8(*memmap.PtrUint8(0x973F18, 35972)))
				}
			}
			dword_5d4594_2487244 = uint32(uintptr(unsafe.Pointer(v2)))
			result = v2
		}
	}
	return result
}
func sub_51D1A0(a1 *float2) *float32 {
	var (
		result *float32
		a2     float2
	)
	result = (*float32)(unsafe.Pointer(uintptr(nox_xxx_mapGenFixCoords_4D3D90(a1, &a2))))
	if result != nil {
		result = sub_579AD0(a2.field_0, a2.field_4)
	}
	return result
}
func sub_51D1E0(lpMem unsafe.Pointer) int32 {
	var (
		v1 unsafe.Pointer
		i  *byte
		v4 int32
		v5 *byte
	)
	if lpMem == nil {
		return 0
	}
	v1 = nox_xxx_waypointGetList_579860()
	if v1 == nil {
		return 0
	}
	for v1 != lpMem {
		v1 = unsafe.Pointer(uintptr(nox_xxx_waypointNext_579870(int32(uintptr(v1)))))
		if v1 == nil {
			return 0
		}
	}
	for i = (*byte)(nox_xxx_waypointGetList_579860()); i != nil; i = (*byte)(unsafe.Pointer(uintptr(nox_xxx_waypointNext_579870(int32(uintptr(unsafe.Pointer(i))))))) {
		v4 = 0
		if *(*byte)(unsafe.Add(unsafe.Pointer(i), 476)) != 0 {
			v5 = (*byte)(unsafe.Add(unsafe.Pointer(i), 96))
			for {
				if *((*unsafe.Pointer)(unsafe.Add(unsafe.Pointer((*unsafe.Pointer)(unsafe.Pointer(v5))), -int(unsafe.Sizeof(unsafe.Pointer(nil))*1)))) == lpMem {
					sub_51D370(int32(uintptr(unsafe.Pointer(i))), int32(uintptr(lpMem)), int8(*v5))
				}
				v4++
				v5 = (*byte)(unsafe.Add(unsafe.Pointer(v5), 8))
				if v4 >= int32(uint8(*(*byte)(unsafe.Add(unsafe.Pointer(i), 476)))) {
					break
				}
			}
		}
	}
	sub_579B30(lpMem)
	return 1
}
func sub_51D270(a1 *float32) *float32 {
	var (
		result *float32
		v2     float2
	)
	result = (*float32)(unsafe.Pointer(uintptr(nox_xxx_mapGenFixCoords_4D3D90((*float2)(unsafe.Pointer(a1)), &v2))))
	if result != nil {
		result = sub_579AD0(v2.field_0, v2.field_0)
		if result != nil {
			result = (*float32)(unsafe.Pointer(uintptr(sub_51D1E0(unsafe.Pointer(result)))))
		}
	}
	return result
}
func sub_51D2C0(a1 int32, a2 int32) int32 {
	return sub_51D300(a1, a2, int8(*memmap.PtrUint8(0x973F18, 35972)))
}
func sub_51D2E0(a1 int32, a2 int32) int32 {
	return sub_51D370(a1, a2, int8(*memmap.PtrUint8(0x973F18, 35972)))
}
func sub_51D300(a1 int32, a2 int32, a3 int8) int32 {
	var (
		v3 uint8
		v4 int32
		v5 int32
		v6 *uint8
	)
	v3 = *(*uint8)(unsafe.Pointer(uintptr(a1 + 476)))
	if int32(v3) >= 31 || a1 == a2 {
		return 0
	}
	v4 = 0
	v5 = int32(v3)
	if int32(v3) > 0 {
		v6 = (*uint8)(unsafe.Pointer(uintptr(a1 + 96)))
		for {
			if *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v6))), -int(unsafe.Sizeof(uint32(0))*1)))) == uint32(a2) && int32(*v6) == int32(a3) {
				break
			}
			v4++
			v6 = (*uint8)(unsafe.Add(unsafe.Pointer(v6), 8))
			if v4 >= v5 {
				break
			}
		}
	}
	if v4 != v5 {
		return 0
	}
	*(*uint32)(unsafe.Pointer(uintptr(a1 + v5*8 + 92))) = uint32(a2)
	*(*uint8)(unsafe.Pointer(uintptr(a1 + int32(func() uint8 {
		p := (*uint8)(unsafe.Pointer(uintptr(a1 + 476)))
		x := *p
		*p++
		return x
	}())*8 + 96))) = uint8(a3)
	return 1
}
func sub_51D370(a1 int32, a2 int32, a3 int8) int32 {
	var (
		v3 int32
		v4 int32
		i  *uint8
		v7 int32
		v8 *uint32
	)
	if a1 == a2 {
		return 0
	}
	v3 = 0
	v4 = int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 476))))
	if v4 <= 0 {
		return 0
	}
	for i = (*uint8)(unsafe.Pointer(uintptr(a1 + 96))); *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(i))), -int(unsafe.Sizeof(uint32(0))*1)))) != uint32(a2) || int32(*i) != int32(a3); i = (*uint8)(unsafe.Add(unsafe.Pointer(i), 8)) {
		if func() int32 {
			p := &v3
			*p++
			return *p
		}() >= v4 {
			return 0
		}
	}
	v7 = v3
	if v3 < v4-1 {
		v8 = (*uint32)(unsafe.Pointer(uintptr(a1 + v3*8 + 92)))
		for {
			v7++
			*v8 = *(*uint32)(unsafe.Add(unsafe.Pointer(v8), unsafe.Sizeof(uint32(0))*2))
			*(*uint32)(unsafe.Add(unsafe.Pointer(v8), unsafe.Sizeof(uint32(0))*1)) = *(*uint32)(unsafe.Add(unsafe.Pointer(v8), unsafe.Sizeof(uint32(0))*3))
			v8 = (*uint32)(unsafe.Add(unsafe.Pointer(v8), unsafe.Sizeof(uint32(0))*2))
			if v7 >= int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 476))))-1 {
				break
			}
		}
	}
	*(*uint8)(unsafe.Pointer(uintptr(a1 + 476)))--
	return 1
}
func sub_51D3F0(a1 *float2, a2 *float2) *float2 {
	var (
		result *float2
		v3     *float2
	)
	result = a1
	if a1 != nil {
		if a2 != nil {
			result = (*float2)(unsafe.Pointer(sub_51D1A0(a1)))
			v3 = result
			if result != nil {
				result = (*float2)(unsafe.Pointer(sub_51D1A0(a2)))
				if result != nil {
					result = (*float2)(unsafe.Pointer(uintptr(sub_51D2C0(int32(uintptr(unsafe.Pointer(v3))), int32(uintptr(unsafe.Pointer(result)))))))
				}
			}
		} else {
			result = nil
		}
	}
	return result
}
func sub_51D440(a1 *float32, a2 *float32) *float32 {
	var (
		result *float32
		v3     *float32
	)
	result = a1
	if a1 != nil {
		if a2 != nil {
			result = sub_51D1A0((*float2)(unsafe.Pointer(a1)))
			v3 = result
			if result != nil {
				result = sub_51D1A0((*float2)(unsafe.Pointer(a2)))
				if result != nil {
					result = (*float32)(unsafe.Pointer(uintptr(sub_51D2E0(int32(uintptr(unsafe.Pointer(v3))), int32(uintptr(unsafe.Pointer(result)))))))
				}
			}
		} else {
			result = nil
		}
	}
	return result
}
func sub_51D490(a1 *byte) int32 {
	var (
		v1 int32
		v2 *byte
	)
	if a1 == nil {
		return -1
	}
	v1 = 0
	v2 = (*byte)(memmap.PtrOff(0x85B3FC, 32484))
	for libc.StrCaseCmp(v2, a1) != 0 {
		v2 = (*byte)(unsafe.Add(unsafe.Pointer(v2), 60))
		v1++
		if int32(uintptr(unsafe.Pointer(v2))) >= int32(uintptr(memmap.PtrOff(0x85B3FC, 43044))) {
			return -1
		}
	}
	return v1
}
func nox_xxx_tileGetDefByName_51D4D0(a1 *byte) int32 {
	var (
		v1 int32
		v2 int32
		v3 *byte
	)
	v1 = 0
	v2 = 0
	v3 = (*byte)(memmap.PtrOff(0x85B3FC, 32484))
	for {
		if libc.StrCaseCmp(v3, a1) == 0 {
			v1 = 1
			*memmap.PtrUint32(0x973F18, 35912) = uint32(v2)
		}
		v3 = (*byte)(unsafe.Add(unsafe.Pointer(v3), 60))
		v2++
		if int32(uintptr(unsafe.Pointer(v3))) >= int32(uintptr(memmap.PtrOff(0x85B3FC, 43044))) {
			break
		}
	}
	if libc.StrCaseCmp(a1, CString("NONE")) == 0 {
		*memmap.PtrUint32(0x973F18, 35912) = math.MaxUint8
		return 1
	}
	if v1 != 0 {
		return 1
	}
	*memmap.PtrUint32(0x973F18, 35912) = 0
	return 0
}
func nox_xxx_tileCheckImage_51D540(a1 int32) int32 {
	var result int32
	if a1 < 0 || a1 >= 176 {
		*memmap.PtrUint32(0x973F18, 35912) = 0
		result = 0
	} else {
		*memmap.PtrUint32(0x973F18, 35912) = uint32(a1)
		result = 1
	}
	return result
}
func nox_xxx_tileCheckImageVari_51D570(a1 int32) int32 {
	var result int32
	if a1 <= int32(*memmap.PtrUint8(0x85B3FC, *memmap.PtrUint32(0x973F18, 35912)*60+0x7F18))*int32(*memmap.PtrUint8(0x85B3FC, *memmap.PtrUint32(0x973F18, 35912)*60+0x7F19))-1 {
		dword_5d4594_3835348 = uint32(a1)
		result = 1
	} else {
		dword_5d4594_3835348 = 0
		result = 0
	}
	return result
}
func nox_xxx_tile_51D5C0(a1 int32) int32 {
	if a1 != 1 && a1 != 0 {
		return 0
	}
	*memmap.PtrUint32(0x973F18, 35916) = uint32(a1)
	return 1
}
func sub_51D5E0(a1 *float32) *float32 {
	var (
		result *float32
		v2     float2
	)
	result = a1
	if a1 != nil {
		if *memmap.PtrUint32(0x973F18, 35912) == math.MaxUint8 || (func() bool {
			nox_xxx_mapGenFixCoords_4D3D90((*float2)(unsafe.Pointer(a1)), &v2)
			return (func() *float32 {
				result = (*float32)(unsafe.Pointer(uintptr(sub_51D8F0(&v2))))
				return result
			}()) != nil
		}()) && (dword_5d4594_3835364 != 1 || (func() *float32 {
			result = (*float32)(unsafe.Pointer(uintptr(sub_527380(&v2.field_0))))
			return result
		}()) != nil) {
			result = (*float32)(unsafe.Pointer(uintptr(1)))
		}
	}
	return result
}
func sub_51D8B0(a1 int32) int32 {
	if a1 == 0 {
		return -1
	}
	if *(*uint32)(unsafe.Pointer(uintptr(a1))) != 0 {
		return int32(**(**uint32)(unsafe.Pointer(uintptr(a1))))
	}
	return -1
}
func sub_51D8D0(a1 int32) int32 {
	if a1 == 0 {
		return -1
	}
	if *(*uint32)(unsafe.Pointer(uintptr(a1))) != 0 {
		return int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1))) + 4))))
	}
	return -1
}
func sub_51D8F0(a1 *float2) int32 {
	var (
		v1     float64
		v2     float64
		v3     int32
		v4     int32
		v5     int32
		v6     int32
		result int32
		v8     int32
		v9     int32
	)
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int8
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 float32
	v9 = int32(dword_5d4594_3835348)
	v1 = float64(a1.field_0) + 11.5
	v8 = int32(*memmap.PtrUint32(0x973F18, 35912))
	v10 = 0
	v13 = 0
	v11 = -1
	v12 = -1
	v14 = math.MaxUint8
	v15 = 0
	v2 = float64(a1.field_4) + 11.5
	v3 = int32(int64(v1 * 0.021739131))
	v16 = float32(v2)
	v4 = int32(int64(v2 * 0.021739131))
	v5 = int32(int64(v1)) % 46
	v6 = int32(int64(v16) % 46)
	if *memmap.PtrUint32(0x973F18, 35912) == math.MaxUint8 {
		result = sub_51D9C0(v3, v4, v5, v6, 0)
	} else {
		result = sub_51D9C0(v3, v4, v5, v6, int32(uintptr(unsafe.Pointer(&v8))))
	}
	return result
}
func sub_51D9C0(a1 int32, a2 int32, a3 int32, a4 int32, a5 int32) int32 {
	var result int32
	if a1-1 <= 0 || a1 >= math.MaxInt8 || a2-1 <= 0 || a2 >= math.MaxInt8 {
		return 0
	}
	if a3 <= a4 {
		if 46-a3 <= a4 {
			result = sub_51DA70(a1, a2, a5, 2, 0)
		} else {
			result = sub_51DA70(a1-1, a2, a5, 1, 0)
		}
	} else if 46-a3 <= a4 {
		result = sub_51DA70(a1, a2, a5, 1, 0)
	} else {
		result = sub_51DA70(a1, a2-1, a5, 2, 0)
	}
	return result
}
