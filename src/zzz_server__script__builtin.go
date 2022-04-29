package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	"github.com/gotranspile/cxgo/runtime/stdio"
	noxflags "github.com/noxworld-dev/opennox/v1/common/flags"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"math"
	"unsafe"
)

var dword_5d4594_3821636 uint32 = 0
var dword_5d4594_3821640 uint32 = 0
var dword_5d4594_1599628 uint32 = 0
var nox_script_objGold unsafe.Pointer = nil
var nox_script_objTelekinesisHand unsafe.Pointer = nil
var dword_5d4594_2386848 uint32 = 0
var dword_5d4594_2386852 uint32 = 0
var nox_xxx_imagCasterUnit_1569664 unsafe.Pointer = nil

func nox_script_builtinGetF40() int32 {
	return int32(dword_5d4594_3821636)
}
func nox_script_builtinGetF44() int32 {
	return int32(dword_5d4594_3821640)
}

var nox_script_builtin_buf_xxx [64]byte = [64]byte{}

func nox_script_resetBuiltin() {
	dword_5d4594_3821640 = 0
	nox_script_builtin_buf_xxx[0] = 0
	dword_5d4594_3821636 = 0
}
func sub_5165D0() int32 {
	*memmap.PtrUint32(6112660, 2386828) = uint32(nox_script_pop() - 1)
	sub_413A00(1)
	return nox_client_screenFadeTimeout_44DAB0(25, 1, sub_516570)
}
func sub_512E80(a1 int32) int32 {
	var v1 int32
	v1 = int32(dword_5d4594_1599628)
	if *(*int32)(unsafe.Pointer(&dword_5d4594_1599628)) < 1024 {
		*memmap.PtrUint32(0x973F18, dword_5d4594_1599628*4+0x6828) = uint32(a1)
		dword_5d4594_1599628 = uint32(func() int32 {
			p := &v1
			*p++
			return *p
		}())
	}
	return v1 - 1
}
func sub_511E60() {
	if dword_5d4594_2386836 == 0 {
		dword_5d4594_2386836 = uint32(nox_xxx_getNameId_4E3AA0(CString("Mover")))
	}
	nox_script_freeEverything_5058F0()
	nox_script_activatorCancelAll_51AC60()
	*memmap.PtrUint32(6112660, 2386844) = 0
	dword_5d4594_2386848 = 0
	dword_5d4594_2386852 = 0
}
func nox_xxx_playerCanCarryItem_513B00(a1 int32, a2 int32) {
	var (
		v2 *uint32
		v3 int32
		v4 int32
		v5 int32
		v6 int32
		v7 float2
	)
	if *memmap.PtrUint32(6112660, 2386856) == 0 {
		*memmap.PtrUint32(6112660, 2386856) = uint32(nox_xxx_getNameId_4E3AA0(CString("Glyph")))
	}
	if sub_467B00(int32(*(*uint16)(unsafe.Pointer(uintptr(a2 + 4)))), 1)-*(*int32)(unsafe.Pointer(&dword_5d4594_2386848)) <= 0 {
		v2 = nil
		v3 = 0xF423F
		v4 = a1.FirstItem()
		if v4 != 0 {
			for {
				if (int32(*(*uint8)(unsafe.Pointer(uintptr(v4 + 8)))) & 16) == 0 {
					v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(v4 + 16))))
					if (v5&256) == 0 && uint32(*(*uint16)(unsafe.Pointer(uintptr(v4 + 4)))) != *memmap.PtrUint32(6112660, 2386856) && nox_xxx_ItemIsDroppable_53EBF0(v4) == 0 {
						v6 = nox_xxx_shopGetItemCost_50E3D0(1, 0, *(*float32)(unsafe.Pointer(&v4)))
						if v6 < v3 {
							v3 = v6
							v2 = (*uint32)(unsafe.Pointer(uintptr(v4)))
						}
					}
				}
				v4 = nox_xxx_inventoryGetNext_4E7990(v4)
				if v4 == 0 {
					break
				}
			}
			if v2 != nil {
				sub_4ED970(50.0, (*float2)(unsafe.Pointer(uintptr(a1+56))), &v7)
				nox_xxx_drop_4ED790((*nox_object_t)(unsafe.Pointer(uintptr(a1))), (*nox_object_t)(unsafe.Pointer(v2)), &v7)
				if dword_5d4594_2386852 == 0 {
					nox_xxx_netPriMsgToPlayer_4DA2C0((*nox_object_t)(unsafe.Pointer(uintptr(a1))), CString("pickup.c:CarryingTooMuch"), 0)
					dword_5d4594_2386852 = 1
				}
			}
		}
	}
}
func nox_script_getWall_511EB0() int32 {
	var (
		v0 int32
		v1 int32
	)
	v0 = nox_script_pop()
	v1 = nox_script_pop()
	if nox_script_builtinGetF40() != 0 || nox_script_builtinGetF44() != 0 {
		v1 = (nox_script_builtinGetF40() + v1*23) / 23
		v0 = (nox_script_builtinGetF44() + v0*23) / 23
	}
	if nox_server_getWallAtGrid_410580(v1, v0) != nil {
		nox_script_push(v0 | v1<<16)
	} else {
		nox_script_push(0)
	}
	return 0
}
func nox_script_openSecretWall_511F50() int32 {
	var (
		v0 int32
		v1 int32
	)
	v0 = nox_script_pop()
	v1 = int32(uintptr(nox_server_getWallAtGrid_410580(v0>>16, int32(uint16(int16(v0))))))
	if v1 != 0 {
		nox_xxx_wallOpen_511F80(v1)
	}
	return 0
}
func nox_script_openWallGroup_512010() int32 {
	var (
		v0 int32
		v1 *uint8
	)
	v0 = nox_script_pop()
	v1 = (*uint8)(unsafe.Pointer(uintptr(nox_server_scriptGetGroup_57C0A0(v0))))
	nox_server_scriptExecuteFnForEachGroupObj_502670(v1, 2, func(arg1 int32, arg2 int32) {
		nox_xxx_wallOpen_511F80(arg1)
	}, 0)
	return 0
}
func nox_script_closeWall_512040() int32 {
	var (
		v0 int32
		v1 int32
	)
	v0 = nox_script_pop()
	v1 = int32(uintptr(nox_server_getWallAtGrid_410580(v0>>16, int32(uint16(int16(v0))))))
	if v1 != 0 {
		nox_xxx_wallClose_512070(v1)
	}
	return 0
}
func nox_script_closeWallGroup_512100() int32 {
	var (
		v0 int32
		v1 *uint8
	)
	v0 = nox_script_pop()
	v1 = (*uint8)(unsafe.Pointer(uintptr(nox_server_scriptGetGroup_57C0A0(v0))))
	nox_server_scriptExecuteFnForEachGroupObj_502670(v1, 2, func(arg1 int32, arg2 int32) {
		nox_xxx_wallClose_512070(arg1)
	}, 0)
	return 0
}
func nox_script_toggleWall_512130() int32 {
	var (
		v0 int32
		v1 int32
	)
	v0 = nox_script_pop()
	v1 = int32(uintptr(nox_server_getWallAtGrid_410580(v0>>16, int32(uint16(int16(v0))))))
	if v1 != 0 {
		nox_xxx_wallToggle_512160(v1)
	}
	return 0
}
func nox_script_toggleWallGroup_512260() int32 {
	var (
		v0 int32
		v1 *uint8
	)
	v0 = nox_script_pop()
	v1 = (*uint8)(unsafe.Pointer(uintptr(nox_server_scriptGetGroup_57C0A0(v0))))
	nox_server_scriptExecuteFnForEachGroupObj_502670(v1, 2, func(arg1 int32, arg2 int32) {
		nox_xxx_wallToggle_512160(arg1)
	}, 0)
	return 0
}
func nox_script_wallBreak_512290() int32 {
	var (
		v0 int32
		v1 int32
	)
	v0 = nox_script_pop()
	v1 = int32(uintptr(nox_server_getWallAtGrid_410580(v0>>16, int32(uint16(int16(v0))))))
	if v1 != 0 {
		nox_xxx_wallPreDestroyByPtr_5122C0(v1)
	}
	return 0
}
func nox_script_wallGroupBreak_5122F0() int32 {
	var (
		v0 int32
		v1 *uint8
	)
	v0 = nox_script_pop()
	v1 = (*uint8)(unsafe.Pointer(uintptr(nox_server_scriptGetGroup_57C0A0(v0))))
	nox_server_scriptExecuteFnForEachGroupObj_502670(v1, 2, func(arg1 int32, arg2 int32) {
		nox_xxx_wallPreDestroyByPtr_5122C0(arg1)
	}, 0)
	return 0
}
func nox_script_secondTimer_512320() int32 {
	var (
		v0 int32
		v1 int32
	)
	v0 = nox_script_pop()
	v1 = nox_script_pop()
	nox_script_activatorTimer_51ACA0(int32(uint32(v1)*nox_gameFPS), v0, 0)
	return 0
}
func nox_script_frameTimer_512350() int32 {
	var (
		v0 int32
		v1 int32
	)
	v0 = nox_script_pop()
	v1 = nox_script_pop()
	nox_script_activatorTimer_51ACA0(v1, v0, 0)
	return 0
}
func nox_script_moverOrMonsterGo_512370() int32 {
	var (
		v0 int32
		v1 int32
		v2 int32
		v3 int32
		v4 int32
		v5 *uint32
	)
	v0 = nox_script_pop()
	v1 = nox_script_pop()
	v2 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v1))))
	v3 = v2
	if v2 != 0 {
		v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 16))))
		if (v4 & 0x8000) == 0 {
			v5 = nox_server_getWaypointById_579C40(v0)
			if v5 != nil {
				nox_server_scriptMoveTo_5123C0(v3, int32(uintptr(unsafe.Pointer(v5))))
			}
		}
	}
	return 0
}
func nox_server_scriptMoveTo_5123C0(a1 int32, a2 int32) *int32 {
	var (
		result *int32
		v3     int32
		v4     *int32
		v5     *int32
		i      *int32
	)
	result = *(**int32)(unsafe.Pointer(uintptr(a1 + 16)))
	if int32(*(*int8)(unsafe.Add(unsafe.Pointer((*int8)(unsafe.Pointer(&result))), 1))) >= 0 {
		if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))&2 != 0 {
			v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
			nox_xxx_monsterClearActionStack_50A3A0(a1)
			v4 = nox_xxx_monsterPushAction_50A260_impl((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 32, CString(__FILE__), __LINE__)
			if v4 != nil {
				*(*int32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(int32(0))*1)) = 8
			}
			if int32(*(*uint8)(unsafe.Pointer(uintptr(a2 + 476)))) != 0 {
				v5 = nox_xxx_monsterPushAction_50A260_impl((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 10, CString(__FILE__), __LINE__)
				if v5 != nil {
					*(*int32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(int32(0))*1)) = a2
					*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(v5))), 12))) = *(*uint8)(unsafe.Pointer(uintptr(v3 + 1332)))
				}
			}
			result = nox_xxx_monsterPushAction_50A260_impl((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 8, CString(__FILE__), __LINE__)
			if result != nil {
				*(*int32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(int32(0))*1)) = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 8))))
				*(*int32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(int32(0))*2)) = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 12))))
				*(*int32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(int32(0))*3)) = 0
			}
		} else if uint32(*(*uint16)(unsafe.Pointer(uintptr(a1 + 4)))) == dword_5d4594_2386836 {
			result = (*int32)(unsafe.Pointer(uintptr(nox_xxx_moverGoTo_5124C0((*uint32)(unsafe.Pointer(uintptr(a1))), (*uint32)(unsafe.Pointer(uintptr(a2)))))))
		} else {
			result = (*int32)(unsafe.Pointer(noxServer.firstServerObject()))
			for i = result; result != nil; i = result {
				if uint32(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(i))), unsafe.Sizeof(uint16(0))*2)))) == dword_5d4594_2386836 && *(*uint32)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(i), unsafe.Sizeof(int32(0))*187)) + 32))) == *(*uint32)(unsafe.Pointer(uintptr(a1 + 40))) {
					nox_xxx_moverGoTo_5124C0((*uint32)(unsafe.Pointer(i)), (*uint32)(unsafe.Pointer(uintptr(a2))))
				}
				result = (*int32)(unsafe.Pointer((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(i)))))).Next()))
			}
		}
	}
	return result
}
func nox_xxx_moverGoTo_5124C0(a1 *uint32, a2 *uint32) int32 {
	var v2 int32
	v2 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*187)))
	nox_xxx_objectSetOn_4E75B0((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1)))))))
	*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*20)) = 0
	*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*21)) = 0
	*(*uint8)(unsafe.Pointer(uintptr(v2))) = 0
	*(*uint32)(unsafe.Pointer(uintptr(v2 + 8))) = *a2
	return int32(uintptr(unsafe.Pointer(nox_xxx_unitAddToUpdatable_4DA8D0((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1))))))))))
}
func nox_script_groupGoTo_512500() int32 {
	var (
		v0 int32
		v1 int32
		v2 int32
		v3 *uint32
		i  *int32
		v5 int32
	)
	v0 = nox_script_pop()
	v1 = nox_script_pop()
	v2 = nox_server_scriptGetGroup_57C0A0(v1)
	v3 = nox_server_getWaypointById_579C40(v0)
	if v3 != nil {
		if v2 != 0 {
			for i = *(**int32)(unsafe.Pointer(uintptr(v2 + 84))); i != nil; i = (*int32)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(i), unsafe.Sizeof(int32(0))*2))))) {
				v5 = nox_xxx_netGetUnitByExtent_4ED020(*i)
				if v5 != 0 {
					nox_server_scriptMoveTo_5123C0(v5, int32(uintptr(unsafe.Pointer(v3))))
				}
			}
		}
	}
	return 0
}
func nox_script_lookAtDirection_512560() int32 {
	var (
		v0 int32
		v1 int32
		v2 int32
		v3 int32
	)
	v0 = nox_script_pop()
	v1 = nox_script_pop()
	v2 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v1))))
	if v2 != 0 {
		if int32(*(*uint8)(unsafe.Pointer(uintptr(v2 + 8))))&2 != 0 {
			v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 16))))
			if (v3 & 0x8000) == 0 {
				nox_xxx_monsterLookAt_5125A0((*nox_object_t)(unsafe.Pointer(uintptr(v2))), v0)
			}
		}
	}
	return 0
}
func nox_script_groupLookAtDirection_512610() int32 {
	var (
		v0     int32
		v1     int32
		result int32
		i      *int32
		v4     int32
		v5     int32
	)
	v0 = nox_script_pop()
	v1 = nox_script_pop()
	result = nox_server_scriptGetGroup_57C0A0(v1)
	if result != 0 {
		for i = *(**int32)(unsafe.Pointer(uintptr(result + 84))); i != nil; i = (*int32)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(i), unsafe.Sizeof(int32(0))*2))))) {
			v4 = nox_xxx_netGetUnitByExtent_4ED020(*i)
			if v4 != 0 && int32(*(*uint8)(unsafe.Pointer(uintptr(v4 + 8))))&2 != 0 {
				v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(v4 + 16))))
				if (v5 & 0x8000) == 0 {
					nox_xxx_monsterLookAt_5125A0((*nox_object_t)(unsafe.Pointer(uintptr(v4))), v0)
				}
			}
		}
		result = 0
	}
	return result
}
func nox_script_objectOn_512670() int32 {
	var (
		v0 int32
		v1 int32
	)
	v0 = nox_script_pop()
	v1 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v0))))
	if v1 != 0 {
		nox_xxx_objectSetOn_4E75B0((*nox_object_t)(unsafe.Pointer(uintptr(v1))))
	}
	return 0
}
func nox_script_objGroupOn_512690() int32 {
	var (
		v0 int32
		v1 *uint8
	)
	v0 = nox_script_pop()
	v1 = (*uint8)(unsafe.Pointer(uintptr(nox_server_scriptGetGroup_57C0A0(v0))))
	nox_server_scriptExecuteFnForEachGroupObj_502670(v1, 0, func(arg1 int32, arg2 int32) {
		sub_5126C0(arg1)
	}, 0)
	return 0
}
func nox_script_waypointOn_5126D0() int32 {
	var (
		v0 int32
		v1 *uint32
	)
	v0 = nox_script_pop()
	v1 = nox_server_getWaypointById_579C40(v0)
	if v1 != nil {
		*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*120)) |= 1
	}
	return 0
}
func nox_script_waypointGroupOn_5126F0() int32 {
	var (
		v0 int32
		v1 *uint8
	)
	v0 = nox_script_pop()
	v1 = (*uint8)(unsafe.Pointer(uintptr(nox_server_scriptGetGroup_57C0A0(v0))))
	nox_server_scriptExecuteFnForEachGroupObj_502670(v1, 1, func(arg1 int32, arg2 int32) {
		sub_512720(arg1)
	}, 0)
	return 0
}
func nox_script_objectOff_512730() int32 {
	var (
		v0 int32
		v1 int32
	)
	v0 = nox_script_pop()
	v1 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v0))))
	if v1 != 0 {
		nox_xxx_objectSetOff_4E7600((*nox_object_t)(unsafe.Pointer(uintptr(v1))))
	}
	return 0
}
func nox_script_objGroupOff_512750() int32 {
	var (
		v0 int32
		v1 *uint8
	)
	v0 = nox_script_pop()
	v1 = (*uint8)(unsafe.Pointer(uintptr(nox_server_scriptGetGroup_57C0A0(v0))))
	nox_server_scriptExecuteFnForEachGroupObj_502670(v1, 0, func(arg1 int32, arg2 int32) {
		sub_512780(arg1)
	}, 0)
	return 0
}
func nox_script_waypointOff_512790() int32 {
	var (
		v0 int32
		v1 *uint32
	)
	v0 = nox_script_pop()
	v1 = nox_server_getWaypointById_579C40(v0)
	if v1 != nil {
		*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*120)) &= 0xFFFFFFFE
	}
	return 0
}
func nox_script_waypointGroupOff_5127B0() int32 {
	var (
		v0 int32
		v1 *uint8
	)
	v0 = nox_script_pop()
	v1 = (*uint8)(unsafe.Pointer(uintptr(nox_server_scriptGetGroup_57C0A0(v0))))
	nox_server_scriptExecuteFnForEachGroupObj_502670(v1, 1, func(arg1 int32, arg2 int32) {
		sub_5127E0(arg1)
	}, 0)
	return 0
}
func nox_script_toggleObject_5127F0() int32 {
	var (
		v0 int32
		v1 int32
	)
	v0 = nox_script_pop()
	v1 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v0))))
	if v1 != 0 {
		nox_xxx_objectToggle_4E7650(v1)
	}
	return 0
}
func nox_script_toggleObjectGroup_512810() int32 {
	var (
		v0 int32
		v1 *uint8
	)
	v0 = nox_script_pop()
	v1 = (*uint8)(unsafe.Pointer(uintptr(nox_server_scriptGetGroup_57C0A0(v0))))
	nox_server_scriptExecuteFnForEachGroupObj_502670(v1, 0, func(arg1 int32, arg2 int32) {
		sub_512840(arg1)
	}, 0)
	return 0
}
func nox_script_toggleWaypoint_512850() int32 {
	var (
		v0 int32
		v1 *uint32
	)
	v0 = nox_script_pop()
	v1 = nox_server_getWaypointById_579C40(v0)
	if v1 != nil {
		*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*120)) ^= 1
	}
	return 0
}
func nox_script_toggleWaypointGroup_512870() int32 {
	var (
		v0 int32
		v1 *uint8
	)
	v0 = nox_script_pop()
	v1 = (*uint8)(unsafe.Pointer(uintptr(nox_server_scriptGetGroup_57C0A0(v0))))
	nox_server_scriptExecuteFnForEachGroupObj_502670(v1, 1, func(arg1 int32, arg2 int32) {
		sub_5128A0(arg1)
	}, 0)
	return 0
}
func nox_script_deleteObject_5128B0() int32 {
	var (
		v0 int32
		v1 int32
	)
	v0 = nox_script_pop()
	v1 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v0))))
	if v1 != 0 {
		nox_xxx_delayedDeleteObject_4E5CC0((*nox_object_t)(unsafe.Pointer(uintptr(v1))))
	}
	return 0
}
func nox_script_deleteObjectGroup_5128D0() int32 {
	var (
		v0 int32
		v1 *uint8
	)
	v0 = nox_script_pop()
	v1 = (*uint8)(unsafe.Pointer(uintptr(nox_server_scriptGetGroup_57C0A0(v0))))
	nox_server_scriptExecuteFnForEachGroupObj_502670(v1, 0, func(arg1 int32, arg2 int32) {
		sub_512900(arg1)
	}, 0)
	return 0
}
func nox_script_followNearestWaypoint_512910() int32 {
	var (
		v0 int32
		v1 int32
	)
	v0 = nox_script_pop()
	v1 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v0))))
	if v1 != 0 {
		nox_xxx_scriptMonsterRoam_512930((*nox_object_t)(unsafe.Pointer(uintptr(v1))))
	}
	return 0
}
func nox_xxx_scriptMonsterRoam_512930(obj *nox_object_t) {
	var (
		a1 int32 = int32(uintptr(unsafe.Pointer(obj)))
		v1 int32
		v2 int32
		v3 *int32
		v4 *int32
	)
	if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 8))))&2 != 0 {
		v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 16))))
		if (v1 & 0x8000) == 0 {
			v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 748))))
			nox_xxx_monsterClearActionStack_50A3A0(a1)
			v3 = nox_xxx_monsterPushAction_50A260_impl((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 32, CString(__FILE__), __LINE__)
			if v3 != nil {
				*(*int32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(int32(0))*1)) = 10
			}
			v4 = nox_xxx_monsterPushAction_50A260_impl((*nox_object_t)(unsafe.Pointer(uintptr(a1))), 10, CString(__FILE__), __LINE__)
			if v4 != nil {
				*(*int32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(int32(0))*1)) = 0
				*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(v4))), 12))) = *(*uint8)(unsafe.Pointer(uintptr(v2 + 1332)))
			}
		}
	}
}
func nox_script_groupRoam_512990() int32 {
	var (
		v0 int32
		v1 *uint8
	)
	v0 = nox_script_pop()
	v1 = (*uint8)(unsafe.Pointer(uintptr(nox_server_scriptGetGroup_57C0A0(v0))))
	nox_server_scriptExecuteFnForEachGroupObj_502670(v1, 0, func(arg1 int32, arg2 int32) {
		nox_xxx_scriptMonsterRoam_512930((*nox_object_t)(unsafe.Pointer(uintptr(arg1))))
	}, 0)
	return 0
}
func nox_script_getObject2_5129C0() int32 {
	var v0 int32
	v0 = nox_script_pop()
	nox_server_scriptValToObjectPtr_511B60(v0)
	return 0
}
func nox_script_getObject3_5129E0() int32 {
	var v0 int32
	v0 = nox_script_pop()
	nox_server_scriptValToObjectPtr_511B60(v0)
	return 0
}
func nox_server_gotoHome(obj *nox_object_t) {
	var (
		v2 int32 = int32(uintptr(unsafe.Pointer(obj)))
		v3 int32
		v4 *uint32
		v5 *int32
		v6 *int32
		v7 *int32
	)
	if int32(*(*uint8)(unsafe.Pointer(uintptr(v2 + 8))))&2 != 0 {
		v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 16))))
		if (v3 & 0x8000) == 0 {
			v4 = *(**uint32)(unsafe.Pointer(uintptr(v2 + 748)))
			nox_xxx_monsterClearActionStack_50A3A0(v2)
			v5 = nox_xxx_monsterPushAction_50A260_impl((*nox_object_t)(unsafe.Pointer(uintptr(v2))), 32, CString(__FILE__), __LINE__)
			if v5 != nil {
				*(*int32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(int32(0))*1)) = 37
			}
			v6 = nox_xxx_monsterPushAction_50A260_impl((*nox_object_t)(unsafe.Pointer(uintptr(v2))), 25, CString(__FILE__), __LINE__)
			if v6 != nil {
				*((*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(v6))), unsafe.Sizeof(float32(0))*1))) = float32(float64(*mem_getFloatPtr(0x587000, *(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*94))*8+0x2F658))*10.0 + float64(*(*float32)(unsafe.Pointer(uintptr(v2 + 56)))))
				*((*float32)(unsafe.Add(unsafe.Pointer((*float32)(unsafe.Pointer(v6))), unsafe.Sizeof(float32(0))*2))) = float32(float64(*mem_getFloatPtr(0x587000, *(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*94))*8+194140))*10.0 + float64(*(*float32)(unsafe.Pointer(uintptr(v2 + 60)))))
			}
			v7 = nox_xxx_monsterPushAction_50A260_impl((*nox_object_t)(unsafe.Pointer(uintptr(v2))), 37, CString(__FILE__), __LINE__)
			if v7 != nil {
				*(*int32)(unsafe.Add(unsafe.Pointer(v7), unsafe.Sizeof(int32(0))*1)) = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*95)))
				*(*int32)(unsafe.Add(unsafe.Pointer(v7), unsafe.Sizeof(int32(0))*2)) = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*96)))
				*(*int32)(unsafe.Add(unsafe.Pointer(v7), unsafe.Sizeof(int32(0))*3)) = 0
			}
		}
	}
}
func nox_script_gotoHome_512A00() int32 {
	var (
		v0     int32
		result int32
		v2     int32
	)
	v0 = nox_script_pop()
	result = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v0))))
	v2 = result
	if v2 != 0 {
		nox_server_gotoHome((*nox_object_t)(unsafe.Pointer(uintptr(v2))))
		result = 0
	}
	return result
}
func nox_script_audioEven_512AC0() int32 {
	var (
		v0 int32
		v1 int32
		v2 *float2
		v3 int32
		v5 *float2
	)
	v0 = nox_script_pop()
	v1 = nox_script_pop()
	v2 = (*float2)(unsafe.Pointer(nox_server_getWaypointById_579C40(v0)))
	if v2 != nil {
		v5 = (*float2)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(float2{})*1))
		v3 = nox_xxx_utilFindSound_40AF50(nox_script_getString_512E40(v1))
		nox_xxx_audCreate_501A30(v3, v5, 0, 0)
	}
	return 0
}
func nox_script_returnOne_512C10() int32 {
	return 1
}
func nox_script_unlockDoor_512C20() int32 {
	var (
		v0 int32
		v1 int32
	)
	v0 = nox_script_pop()
	v1 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v0))))
	if v1 != 0 && (int32(*(*uint8)(unsafe.Pointer(uintptr(v1 + 8))))&128) != 0 {
		*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v1 + 748))) + 1))) = 0
		nox_xxx_aud_501960(234, (*nox_object_t)(unsafe.Pointer(uintptr(v1))), 0, 0)
	}
	return 0
}
func nox_script_lockDoor_512C60() int32 {
	var (
		v0 int32
		v1 int32
	)
	v0 = nox_script_pop()
	v1 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v0))))
	if v1 != 0 && (int32(*(*uint8)(unsafe.Pointer(uintptr(v1 + 8))))&128) != 0 {
		*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v1 + 748))) + 1))) = 5
		nox_xxx_aud_501960(233, (*nox_object_t)(unsafe.Pointer(uintptr(v1))), 0, 0)
	}
	return 0
}
func nox_script_isOn_512CA0() int32 {
	var (
		v0     int32
		v1     int32
		result int32
	)
	v0 = nox_script_pop()
	v1 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v0))))
	if v1 != 0 && *(*uint32)(unsafe.Pointer(uintptr(v1 + 16)))&0x1000000 != 0 {
		nox_script_push(1)
		result = 0
	} else {
		nox_script_push(0)
		result = 0
	}
	return result
}
func nox_script_wpIsEnabled_512CE0() int32 {
	var (
		v0     int32
		v1     *uint32
		result int32
	)
	v0 = nox_script_pop()
	v1 = nox_server_getWaypointById_579C40(v0)
	if v1 != nil && *(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*120))&1 != 0 {
		nox_script_push(1)
		result = 0
	} else {
		nox_script_push(0)
		result = 0
	}
	return result
}
func nox_script_doorIsLocked_512D20() int32 {
	var (
		v0     int32
		v1     int32
		result int32
	)
	v0 = nox_script_pop()
	v1 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v0))))
	if v1 != 0 && (int32(*(*uint8)(unsafe.Pointer(uintptr(v1 + 8))))&128) != 0 {
		nox_script_push(bool2int(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v1 + 748))) + 1)))) != 0))
		result = 0
	} else {
		nox_script_push(0)
		result = 0
	}
	return result
}
func nox_script_randomFloat_512D70() int32 {
	var (
		v0 int32
		v1 int32
		v3 int32
	)
	v0 = nox_script_pop()
	v1 = nox_script_pop()
	*(*float32)(unsafe.Pointer(&v3)) = float32(nox_common_randomFloat_416030(*(*float32)(unsafe.Pointer(&v1)), *(*float32)(unsafe.Pointer(&v0))))
	nox_script_push(v3)
	return 0
}
func nox_script_randomInt_512DB0() int32 {
	var (
		v0 int32
		v1 int32
		v2 int32
	)
	v0 = nox_script_pop()
	v1 = nox_script_pop()
	v2 = noxRndCounter1.IntClamp(v1, v0)
	nox_script_push(v2)
	return 0
}
func nox_script_timerSecSpecial_512DE0() int32 {
	var (
		v0 int32
		v1 int32
		v2 int32
	)
	v0 = nox_script_pop()
	v1 = nox_script_pop()
	v2 = nox_script_pop()
	nox_script_activatorTimer_51ACA0(int32(uint32(v2)*nox_gameFPS), v0, v1)
	return 0
}
func nox_script_specialTimer_512E10() int32 {
	var (
		v0 int32
		v1 int32
		v2 int32
	)
	v0 = nox_script_pop()
	v1 = nox_script_pop()
	v2 = nox_script_pop()
	nox_script_activatorTimer_51ACA0(v2, v0, v1)
	return 0
}
func nox_script_intToString_512EA0() int32 {
	var (
		v0 int32
		v1 int32
	)
	v0 = nox_script_pop()
	compat_itoa(v0, (*byte)(memmap.PtrOff(0x973F18, 22280)), 10)
	v1 = nox_script_addString_512E40((*byte)(memmap.PtrOff(0x973F18, 22280)))
	nox_script_push(v1)
	return 0
}
func nox_script_floatToString_512ED0() int32 {
	var (
		v0 float32
		v1 int32
	)
	v0 = nox_script_popf()
	nox_sprintf((*byte)(memmap.PtrOff(0x973F18, 22280)), CString("%f"), v0, v0)
	v1 = nox_script_addString_512E40((*byte)(memmap.PtrOff(0x973F18, 22280)))
	nox_script_push(v1)
	return 0
}
func nox_script_create_512F10() int32 {
	var (
		v0 int32
		v1 int32
		v2 *float32
		v3 *uint32
		v4 *uint32
	)
	v0 = nox_script_pop()
	v1 = nox_script_pop()
	v2 = (*float32)(unsafe.Pointer(nox_server_getWaypointById_579C40(v0)))
	if v2 != nil {
		v3 = (*uint32)(unsafe.Pointer(nox_xxx_newObjectByTypeID_4E3810(nox_script_getString_512E40(v1))))
		v4 = v3
		if v3 == nil {
			nox_script_push(0)
			return 0
		}
		nox_xxx_createAt_4DAA50((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v3)))))), nil, *(*float32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(float32(0))*2)), *(*float32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(float32(0))*3)))
		nox_script_push(int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*11))))
	}
	return 0
}
func nox_script_damage_512F80() int32 {
	var (
		v0 int32
		v1 int32
		v2 int32
		v3 int32
		v4 int32
		v6 [3]int32
	)
	v0 = nox_script_pop()
	v1 = nox_script_pop()
	v2 = nox_script_pop()
	v3 = nox_script_pop()
	v4 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v3))))
	if v4 != 0 {
		v6[0] = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v2))))
		v6[1] = v1
		v6[2] = v0
		sub_512FE0(v4, &v6[0])
	}
	return 0
}
func nox_script_groupDamage_513010() int32 {
	var (
		v0 int32
		v1 int32
		v2 int32
		v3 int32
		v4 *uint8
		v6 [3]int32
	)
	v0 = nox_script_pop()
	v1 = nox_script_pop()
	v2 = nox_script_pop()
	v3 = nox_script_pop()
	v6[0] = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v2))))
	v6[1] = v1
	v6[2] = v0
	v4 = (*uint8)(unsafe.Pointer(uintptr(nox_server_scriptGetGroup_57C0A0(v3))))
	nox_server_scriptExecuteFnForEachGroupObj_502670(v4, 0, func(arg1 int32, arg2 int32) {
		sub_512FE0(arg1, (*int32)(unsafe.Pointer(uintptr(arg2))))
	}, int32(uintptr(unsafe.Pointer(&v6[0]))))
	return 0
}
func nox_script_builtin_513070() int32 {
	var (
		v0 int32
		v1 int32
		v2 int32
		v4 float32
		v5 [3]int32
	)
	v4 = nox_script_popf()
	v0 = nox_script_pop()
	v1 = nox_script_pop()
	v2 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v1))))
	if v2 != 0 {
		v5[1] = v0
		v5[0] = int32(int64(v4))
		sub_5130E0(v2, (*uint32)(unsafe.Pointer(&v5[0])))
		nox_script_push(v5[2])
	} else {
		nox_script_push(0)
	}
	return 0
}
func nox_script_builtin_513160() int32 {
	var (
		v0 int32
		v1 int32
		v2 *uint8
		v4 float32
		v5 [3]int32
	)
	v4 = nox_script_popf()
	v0 = nox_script_pop()
	v1 = nox_script_pop()
	v5[1] = v0
	v5[0] = int32(int64(v4))
	v2 = (*uint8)(unsafe.Pointer(uintptr(nox_server_scriptGetGroup_57C0A0(v1))))
	nox_server_scriptExecuteFnForEachGroupObj_502670(v2, 0, func(arg1 int32, arg2 int32) {
		sub_5130E0(arg1, (*uint32)(unsafe.Pointer(uintptr(arg2))))
	}, int32(uintptr(unsafe.Pointer(&v5[0]))))
	return 0
}
func nox_script_awardSpell_5131C0() int32 {
	var (
		v0     int32
		v1     int32
		v2     int32
		result int32
		v4     int32
		v5     int32
	)
	v0 = nox_script_pop()
	v1 = nox_script_pop()
	v2 = things.SpellIndex(nox_script_getString_512E40(v0))
	if v2 != 0 {
		v4 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v1))))
		if v4 != 0 {
			v5 = nox_xxx_spellGrantToPlayer_4FB550((*nox_object_t)(unsafe.Pointer(uintptr(v4))), v2, 1, 0, 0)
			nox_script_push(v5)
		} else {
			nox_script_push(0)
		}
		result = 0
	} else {
		nox_script_push(0)
		result = 0
	}
	return result
}
func nox_script_awardSpellGroup_513230() int32 {
	var (
		v0 int32
		v1 int32
		v2 *uint8
		v4 int32
	)
	v0 = nox_script_pop()
	v1 = nox_script_pop()
	v4 = things.SpellIndex(nox_script_getString_512E40(v0))
	if v4 != 0 {
		v2 = (*uint8)(unsafe.Pointer(uintptr(nox_server_scriptGetGroup_57C0A0(v1))))
		nox_server_scriptExecuteFnForEachGroupObj_502670(v2, 0, func(arg1 int32, arg2 int32) {
			sub_513280(arg1, (*int32)(unsafe.Pointer(uintptr(arg2))))
		}, int32(uintptr(unsafe.Pointer(&v4))))
	}
	return 0
}
func nox_script_enchant_5132E0() int32 {
	var (
		v0 int32
		v1 int32
		v2 int32
		v3 int32
		v5 float32
		v6 [2]int32
	)
	v5 = nox_script_popf()
	v0 = nox_script_pop()
	v1 = nox_script_pop()
	var v8 *byte = nox_script_getString_512E40(v0)
	v2 = nox_xxx_enchantByName_424880(v8)
	if v2 != -1 {
		v3 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v1))))
		if v3 != 0 {
			v6[0] = v2
			v6[1] = int32(int64(float64(nox_gameFPS) * float64(v5)))
			nox_xxx_enchantUnit_513390(v3, &v6[0])
		}
	}
	return 0
}
func nox_script_groupEnchant_5133B0() int32 {
	var (
		v0 int32
		v1 int32
		v2 int32
		v3 *uint8
		v5 float32
		v6 [2]int32
	)
	v5 = nox_script_popf()
	v0 = nox_script_pop()
	v1 = nox_script_pop()
	var v7 *byte = nox_script_getString_512E40(v0)
	v2 = nox_xxx_enchantByName_424880(v7)
	if v2 != -1 {
		v6[0] = v2
		v6[1] = int32(int64(float64(nox_gameFPS) * float64(v5)))
		v3 = (*uint8)(unsafe.Pointer(uintptr(nox_server_scriptGetGroup_57C0A0(v1))))
		nox_server_scriptExecuteFnForEachGroupObj_502670(v3, 0, func(arg1 int32, arg2 int32) {
			nox_xxx_enchantUnit_513390(arg1, (*int32)(unsafe.Pointer(uintptr(arg2))))
		}, int32(uintptr(unsafe.Pointer(&v6[0]))))
	}
	return 0
}
func nox_script_getHost_513460() int32 {
	var v0 int32
	v0 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer(&nox_common_playerInfoFromNum_417090(31).field_0), unsafe.Sizeof(uint32(0))*514))))
	if v0 != 0 {
		nox_script_push(int32(*(*uint32)(unsafe.Pointer(uintptr(v0 + 44)))))
	} else {
		nox_script_push(0)
	}
	return 0
}
func nox_script_objectGet_513490() int32 {
	var (
		v0 int32
		v2 [76]byte
	)
	libc.StrNCpy(&v2[0], nox_script_getString_512E40(nox_script_pop()), int(unsafe.Sizeof([76]byte{})-1))
	libc.StrNCat(&v2[0], &nox_script_builtin_buf_xxx[0], int(unsafe.Sizeof([76]byte{})-1))
	v0 = nox_xxx_getObjectByScrName_4DA4F0(&v2[0])
	if v0 != 0 {
		nox_script_push(int32(*(*uint32)(unsafe.Pointer(uintptr(v0 + 44)))))
	} else {
		nox_script_push(0)
	}
	return 0
}
func nox_script_getObjectX_513530() int32 {
	var (
		v0 int32
		v1 int32
	)
	v0 = nox_script_pop()
	v1 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v0))))
	if v1 != 0 {
		nox_script_push(int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 56)))))
	} else {
		nox_script_push(0)
	}
	return 0
}
func nox_script_getWaypointX_513570() int32 {
	var (
		v0 int32
		v1 *uint32
	)
	v0 = nox_script_pop()
	v1 = nox_server_getWaypointById_579C40(v0)
	if v1 != nil {
		nox_script_push(int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*2))))
	} else {
		nox_script_push(0)
	}
	return 0
}
func nox_script_getObjectY_5135B0() int32 {
	var (
		v0 int32
		v1 int32
	)
	v0 = nox_script_pop()
	v1 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v0))))
	if v1 != 0 {
		nox_script_push(int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 60)))))
	} else {
		nox_script_push(0)
	}
	return 0
}
func nox_script_getWaypointY_5135F0() int32 {
	var (
		v0 int32
		v1 *uint32
	)
	v0 = nox_script_pop()
	v1 = nox_server_getWaypointById_579C40(v0)
	if v1 != nil {
		nox_script_push(int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*3))))
	} else {
		nox_script_push(0)
	}
	return 0
}
func nox_script_unitHeight_513630() int32 {
	var (
		v0 int32
		v1 int32
	)
	v0 = nox_script_pop()
	v1 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v0))))
	if v1 != 0 {
		nox_script_push(int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 104)))))
	} else {
		nox_script_push(0)
	}
	return 0
}
func nox_script_getUnitLook_513670() int32 {
	var (
		v0 int32
		v1 int32
	)
	v0 = nox_script_pop()
	v1 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v0))))
	if v1 != 0 {
		nox_script_push(int32(*(*int16)(unsafe.Pointer(uintptr(v1 + 124)))))
	} else {
		nox_script_push(0)
	}
	return 0
}
func nox_script_moveObject_5136A0() int32 {
	var (
		v0 int32
		v1 int32
		v3 float2
		v4 float32
	)
	v4 = nox_script_popf()
	*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v3.field_0))), unsafe.Sizeof(uint32(0))*0)) = uint32(nox_script_pop())
	v0 = nox_script_pop()
	v1 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v0))))
	if v1 != 0 {
		v3.field_0 = float32(float64(nox_script_builtinGetF40()) + float64(v3.field_0))
		v3.field_4 = float32(float64(nox_script_builtinGetF44()) + float64(v4))
		nox_xxx_unitMove_4E7010((*nox_object_t)(unsafe.Pointer(uintptr(v1))), &v3)
	}
	return 0
}
func nox_script_moveWaypoint_513700() int32 {
	var (
		v0 int32
		v1 *float32
		v3 float32
		v4 float32
	)
	v4 = nox_script_popf()
	v3 = nox_script_popf()
	v0 = nox_script_pop()
	v1 = (*float32)(unsafe.Pointer(nox_server_getWaypointById_579C40(v0)))
	if v1 != nil {
		*(*float32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(float32(0))*2)) = float32(float64(nox_script_builtinGetF40()) + float64(v3))
		*(*float32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(float32(0))*3)) = float32(float64(nox_script_builtinGetF44()) + float64(v4))
	}
	return 0
}
func nox_script_raise_513750() int32 {
	var (
		v0 int32
		v1 int32
		v2 int32
	)
	v0 = nox_script_pop()
	v1 = nox_script_pop()
	v2 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v1))))
	if v2 != 0 {
		nox_xxx_unitRaise_4E46F0((*nox_object_t)(unsafe.Pointer(uintptr(v2))), *(*float32)(unsafe.Pointer(&v0)))
	}
	return 0
}
func nox_script_faceAngle_513780() int32 {
	var (
		v0 int32
		v1 int32
		v2 int32
		v3 uint32
	)
	v0 = nox_script_pop()
	v1 = nox_script_pop()
	v2 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v1))))
	if v2 != 0 {
		v3 = uint32(v0)
		if v0 < 0 {
			v3 = uint32(v0) + (uint32(math.MaxUint8-v0) >> 8 << 8)
		}
		if int32(v3) >= 256 {
			*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v3))), unsafe.Sizeof(uint16(0))*0)) = uint16((v3>>8)*0xFFFFFF00 + v3)
		}
		*(*uint16)(unsafe.Pointer(uintptr(v2 + 126))) = uint16(v3)
		*(*uint16)(unsafe.Pointer(uintptr(v2 + 124))) = uint16(v3)
	}
	return 0
}
func nox_script_pushObject_5137D0() int32 {
	var (
		v0 int32
		v1 int32
		v3 float32
		v4 float32
	)
	v4 = nox_script_popf()
	v3 = nox_script_popf()
	v0 = nox_script_pop()
	v1 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v0))))
	if v1 != 0 {
		*(*float32)(unsafe.Pointer(uintptr(v1 + 88))) = v3 + *(*float32)(unsafe.Pointer(uintptr(v1 + 88)))
		*(*float32)(unsafe.Pointer(uintptr(v1 + 92))) = v4 + *(*float32)(unsafe.Pointer(uintptr(v1 + 92)))
	}
	return 0
}
func nox_script_pushObjectTo_513820() int32 {
	var (
		v0 int32
		v1 *float32
		v2 float64
		v3 float64
		v5 float32
		v6 float32
		v7 float32
		v8 float32
		v9 float32
	)
	v5 = nox_script_popf()
	v8 = nox_script_popf()
	v7 = nox_script_popf()
	v0 = nox_script_pop()
	v1 = (*float32)(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v0)))
	if v1 != nil {
		v2 = float64(*(*float32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(float32(0))*14))-v8) + float64(nox_script_builtinGetF40())
		v3 = float64(*(*float32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(float32(0))*15))-v5) + float64(nox_script_builtinGetF44())
		v6 = float32(v3)
		v9 = float32(math.Sqrt(v3*float64(v6) + v2*v2))
		*(*float32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(float32(0))*22)) = float32(float64(v7)*v2/float64(v9) + float64(*(*float32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(float32(0))*22))))
		*(*float32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(float32(0))*23)) = v7*v6/v9 + *(*float32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(float32(0))*23))
	}
	return 0
}
func nox_script_getFirstInvItem_5138B0() int32 {
	var (
		v0 int32
		v1 int32
		v2 int32
		v3 int32
	)
	v0 = nox_script_pop()
	v1 = 0
	v2 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v0))))
	if v2 != 0 {
		v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 504))))
		if v3 != 0 {
			v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(v3 + 44))))
		}
	}
	nox_script_push(v1)
	return 0
}
func nox_script_getNextInvItem_5138E0() int32 {
	var (
		v0 int32
		v1 int32
		v2 int32
		v3 int32
	)
	v0 = nox_script_pop()
	v1 = 0
	v2 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v0))))
	if v2 != 0 {
		v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 496))))
		if v3 != 0 {
			v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(v3 + 44))))
		}
	}
	nox_script_push(v1)
	return 0
}
func nox_script_hasItem_513910() int32 {
	var (
		v0 int32
		v1 int32
		v2 int32
		v3 int32
		v4 int32
	)
	v0 = nox_script_pop()
	v1 = nox_script_pop()
	v2 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v1))))
	v3 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v0))))
	v4 = 0
	if v2 != 0 && v3 != 0 && nox_xxx_unitInventoryContains_4F78E0(v2, v3) != 0 {
		v4 = 1
	}
	nox_script_push(v4)
	return 0
}
func nox_script_getInvHolder_513960() int32 {
	var (
		v0 int32
		v1 int32
	)
	v0 = nox_script_pop()
	v1 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v0))))
	if v1 != 0 {
		nox_script_push(int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v1 + 492))) + 44)))))
	} else {
		nox_script_push(0)
	}
	return 0
}
func nox_script_pickup_5139A0() int32 {
	var (
		v0 int32
		v1 int32
		v2 int32
		v3 int32
		v4 int32
		v5 int32
		v6 int32
	)
	if nox_script_objGold == nil {
		nox_script_objGold = unsafe.Pointer(uintptr(nox_xxx_getNameId_4E3AA0(CString("Gold"))))
		*memmap.PtrUint32(6112660, 2386864) = uint32(nox_xxx_getNameId_4E3AA0(CString("QuestGoldPile")))
		*memmap.PtrUint32(6112660, 2386868) = uint32(nox_xxx_getNameId_4E3AA0(CString("QuestGoldChest")))
	}
	v0 = nox_script_pop()
	v1 = nox_script_pop()
	v2 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v1))))
	v3 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v0))))
	v4 = v3
	if v2 != 0 && v3 != 0 {
		if noxflags.HasGame(noxflags.GameModeCoop) && (int32(*(*uint8)(unsafe.Pointer(uintptr(v2 + 8))))&4) == 4 && *memmap.PtrUint32(6112660, 2386844) != nox_frame_xxx_2598000 {
			*memmap.PtrUint32(6112660, 2386844) = nox_frame_xxx_2598000
			dword_5d4594_2386848 = 0
			dword_5d4594_2386852 = 0
		}
		if int32(*(*uint8)(unsafe.Pointer(uintptr(v2 + 8))))&4 != 0 {
			if noxflags.HasGame(noxflags.GameModeCoop) {
				v5 = int32(*(*uint16)(unsafe.Pointer(uintptr(v2 + 4))))
				if unsafe.Pointer(uintptr(uint16(int16(v5)))) != nox_script_objGold && uint32(v5) != *memmap.PtrUint32(6112660, 2386864) && uint32(v5) != *memmap.PtrUint32(6112660, 2386868) {
					nox_xxx_playerCanCarryItem_513B00(v2, v4)
				}
			}
		}
		v6 = nox_xxx_inventoryServPlace_4F36F0(v2, v4, 1, 1)
		if v6 == 1 && int32(*(*uint8)(unsafe.Pointer(uintptr(v2 + 8))))&4 != 0 && noxflags.HasGame(noxflags.GameModeCoop) && unsafe.Pointer(uintptr(*(*uint16)(unsafe.Pointer(uintptr(v4 + 4))))) != nox_script_objGold {
			dword_5d4594_2386848++
			nox_script_push(1)
			return 0
		}
	} else {
		v6 = 0
	}
	nox_script_push(v6)
	return 0
}
func nox_script_drop_513C10() int32 {
	var (
		v0     int32
		v1     int32
		v2     int32
		v3     *uint32
		v4     int32
		result int32
	)
	v0 = nox_script_pop()
	v1 = nox_script_pop()
	v2 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v1))))
	v3 = (*uint32)(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v0)))
	if v2 != 0 && v3 != nil {
		v4 = nox_xxx_invForceDropItem_4ED930(v2, v3)
		nox_script_push(v4)
		result = 0
	} else {
		nox_script_push(0)
		result = 0
	}
	return result
}
func nox_script_builtin_513C60() int32 {
	return 0
}
func nox_script_TestBuffs_513C70() int32 {
	var (
		v0     int32
		v1     int32
		v2     int32
		v3     int8
		v4     int32
		v5     int32
		result int32
	)
	v0 = nox_script_pop()
	v1 = nox_script_pop()
	var v7 *byte = nox_script_getString_512E40(v0)
	v2 = nox_xxx_enchantByName_424880(v7)
	v3 = int8(v2)
	if v2 != -1 && (func() int32 {
		v4 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v1))))
		return v4
	}()) != 0 {
		v5 = nox_xxx_testUnitBuffs_4FF350((*nox_object_t)(unsafe.Pointer(uintptr(v4))), v3)
		nox_script_push(v5)
		result = 0
	} else {
		nox_script_push(0)
		result = 0
	}
	return result
}
func nox_script_cancelBuff_513D00() int32 {
	var (
		v0 int32
		v1 int32
		v2 int32
		v3 int32
	)
	v0 = nox_script_pop()
	v1 = nox_script_pop()
	var v5 *byte = nox_script_getString_512E40(v0)
	v2 = nox_xxx_enchantByName_424880(v5)
	if v2 != -1 {
		v3 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v1))))
		if v3 != 0 {
			nox_xxx_spellBuffOff_4FF5B0((*nox_object_t)(unsafe.Pointer(uintptr(v3))), v2)
		}
	}
	return 0
}
func nox_script_getCurrentHP_513D70() int32 {
	var (
		v0     int32
		v1     int32
		v2     *uint16
		result int32
	)
	v0 = nox_script_pop()
	v1 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v0))))
	if v1 != 0 && (func() *uint16 {
		v2 = *(**uint16)(unsafe.Pointer(uintptr(v1 + 556)))
		return v2
	}()) != nil {
		nox_script_push(int32(*v2))
		result = 0
	} else {
		nox_script_push(0)
		result = 0
	}
	return result
}
func nox_script_getMaxHP_513DB0() int32 {
	var (
		v0     int32
		v1     int32
		v2     int32
		result int32
	)
	v0 = nox_script_pop()
	v1 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v0))))
	if v1 != 0 && (func() int32 {
		v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 556))))
		return v2
	}()) != 0 {
		nox_script_push(int32(*(*uint16)(unsafe.Pointer(uintptr(v2 + 4)))))
		result = 0
	} else {
		nox_script_push(0)
		result = 0
	}
	return result
}
func nox_script_restoreHP_513DF0() int32 {
	var (
		v0 int32
		v1 int32
		v2 int32
	)
	v0 = nox_script_pop()
	v1 = nox_script_pop()
	v2 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v1))))
	if v2 != 0 && v0 > 0 {
		nox_xxx_unitAdjustHP_4EE460((*nox_object_t)(unsafe.Pointer(uintptr(v2))), v0)
	}
	return 0
}
func nox_script_getDistance_513E20() int32 {
	var (
		v0 float32
		v2 float32
		v3 float32
		v4 float32
		v5 int32
	)
	v4 = nox_script_popf()
	v2 = nox_script_popf()
	v3 = nox_script_popf()
	v0 = nox_script_popf()
	*(*float32)(unsafe.Pointer(&v5)) = float32(math.Sqrt(float64((v3-v4)*(v3-v4) + (v0-v2)*(v0-v2))))
	nox_script_push(v5)
	return 0
}
func nox_script_canInteract_513E80() int32 {
	var (
		v0 int32
		v1 int32
		v2 int32
		v3 int32
		v4 int32
		v5 float64
		v6 float64
	)
	v0 = nox_script_pop()
	v1 = nox_script_pop()
	v2 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v1))))
	v3 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v0))))
	v4 = 0
	if v2 != 0 && v3 != 0 {
		v5 = float64(*(*float32)(unsafe.Pointer(uintptr(v2 + 56))) - *(*float32)(unsafe.Pointer(uintptr(v3 + 56))))
		if v5 < 0.0 {
			v5 = -v5
		}
		if v5 <= 512.0 {
			v6 = float64(*(*float32)(unsafe.Pointer(uintptr(v2 + 60))) - *(*float32)(unsafe.Pointer(uintptr(v3 + 60))))
			if v6 < 0.0 {
				v6 = -v6
			}
			if v6 <= 512.0 {
				v4 = nox_xxx_unitCanInteractWith_5370E0((*nox_object_t)(unsafe.Pointer(uintptr(v2))), (*nox_object_t)(unsafe.Pointer(uintptr(v3))), 0)
			}
		}
	}
	nox_script_push(v4)
	return 0
}
func nox_script_fn58_513F10() int32 {
	nox_script_pop()
	nox_script_pop()
	return 0
}
func nox_script_fn59_513F20() int32 {
	nox_script_pop()
	nox_script_pop()
	return 0
}
func nox_script_fn5A_513F30() int32 {
	nox_script_pop()
	nox_script_pop()
	return 0
}
func nox_script_fn5B_513F40() int32 {
	nox_script_pop()
	nox_script_pop()
	return 0
}
func nox_script_Fn5C_513F50() int32 {
	nox_script_pop()
	nox_script_pop()
	return 0
}
func nox_script_Fn5D_513F60() int32 {
	nox_script_pop()
	nox_script_pop()
	return 0
}
func nox_script_GetHostInfo_513FA0() int32 {
	var (
		v0     int32
		v1     *byte
		result int32
	)
	v0 = nox_script_pop()
	v1 = nox_xxx_getHostInfoPtr_431770()
	switch v0 {
	case 0:
		nox_script_push(int32(*(*uint32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v1), 50))))))
		result = 0
	case 1:
		nox_script_push(int32(*(*uint32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v1), 54))))))
		result = 0
	case 2:
		nox_script_push(int32(*(*uint32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v1), 58))))))
		result = 0
	case 3:
		nox_script_push(int32(*(*uint32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v1), 62))))))
		result = 0
	case 4:
		nox_script_push(int32(uint8(*(*byte)(unsafe.Add(unsafe.Pointer(v1), 66)))))
		result = 0
	case 5:
		nox_script_push(int32(uint8(*(*byte)(unsafe.Add(unsafe.Pointer(v1), 67)))))
		result = 0
	default:
		nox_script_push(0)
		result = 0
	}
	return result
}
func nox_script_FaceObject_514050() int32 {
	var (
		v0 int32
		v1 int32
		v2 int32
		v3 int32
		v4 int16
		v6 float2
	)
	v0 = nox_script_pop()
	v1 = nox_script_pop()
	v2 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v1))))
	v3 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v0))))
	if v2 != 0 && v3 != 0 {
		v6.field_0 = *(*float32)(unsafe.Pointer(uintptr(v3 + 56))) - *(*float32)(unsafe.Pointer(uintptr(v2 + 56)))
		v6.field_4 = *(*float32)(unsafe.Pointer(uintptr(v3 + 60))) - *(*float32)(unsafe.Pointer(uintptr(v2 + 60)))
		v4 = int16(nox_xxx_math_509ED0(&v6))
		*(*uint16)(unsafe.Pointer(uintptr(v2 + 126))) = uint16(v4)
		*(*uint16)(unsafe.Pointer(uintptr(v2 + 124))) = uint16(v4)
	}
	return 0
}
func nox_script_Walk_5140B0() int32 {
	var (
		v0 int32
		v1 int32
		v3 float32
		v4 float32
		v5 float32
		v6 float32
	)
	v5 = nox_script_popf()
	v6 = nox_script_popf()
	v0 = nox_script_pop()
	v1 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v0))))
	if v1 != 0 {
		v4 = float32(float64(nox_script_builtinGetF44()) + float64(v5))
		v3 = float32(float64(nox_script_builtinGetF40()) + float64(v6))
		nox_xxx_monsterWalkTo_514110((*nox_object_t)(unsafe.Pointer(uintptr(v1))), v3, v4)
	}
	return 0
}
func nox_script_GroupWalk_514170() int32 {
	var (
		v0     int32
		result int32
		i      *int32
		v3     int32
		v4     float32
		v5     float32
		v6     float32
		v7     float32
	)
	v6 = nox_script_popf()
	v7 = nox_script_popf()
	v0 = nox_script_pop()
	result = nox_server_scriptGetGroup_57C0A0(v0)
	if result != 0 {
		for i = *(**int32)(unsafe.Pointer(uintptr(result + 84))); i != nil; i = (*int32)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(i), unsafe.Sizeof(int32(0))*2))))) {
			v3 = nox_xxx_netGetUnitByExtent_4ED020(*i)
			if v3 != 0 {
				v5 = float32(float64(nox_script_builtinGetF44()) + float64(v6))
				v4 = float32(float64(nox_script_builtinGetF40()) + float64(v7))
				nox_xxx_monsterWalkTo_514110((*nox_object_t)(unsafe.Pointer(uintptr(v3))), v4, v5)
			}
		}
		result = 0
	}
	return result
}
func nox_script_CancelTimer_5141F0() int32 {
	var (
		v0 int32
		v1 int32
	)
	v0 = nox_script_pop()
	v1 = bool2int(nox_script_activatorCancel_51AD60(v0))
	nox_script_push(bool2int(v1 != 0))
	return 0
}
func nox_script_Effect_514210() int32 {
	var (
		v0     int32
		v1     *byte
		v2     *byte
		v3     uint32
		v4     *byte
		v5     *byte
		v6     *byte
		v7     int8
		v8     int8
		v9     int8
		result int32
		v11    int32
		v12    int8
		v13    uint8
		v14    float32
		v15    float32
		v16    int32
		v17    float2
		v19    [4]int32
		v20    [64]byte
	)
	v16 = nox_script_pop()
	v0 = nox_script_pop()
	v15 = nox_script_popf()
	v14 = nox_script_popf()
	v1 = nox_script_getString_512E40(nox_script_pop())
	libc.StrNCpy(&v20[0], CString("MSG_FX_"), int(unsafe.Sizeof([64]byte{})-1))
	v17.field_0 = float32(float64(nox_script_builtinGetF40()) + float64(v14))
	v17.field_4 = float32(float64(nox_script_builtinGetF44()) + float64(v15))
	v2 = v1
	v3 = uint32(libc.StrLen(v1) + 1)
	v4 = &v20[libc.StrLen(&v20[0])]
	libc.MemCpy(unsafe.Pointer(v4), unsafe.Pointer(v2), int((v3>>2)*4))
	v6 = (*byte)(unsafe.Add(unsafe.Pointer(v2), (v3>>2)*4))
	v5 = (*byte)(unsafe.Add(unsafe.Pointer(v4), (v3>>2)*4))
	v7 = int8(uint8(v3))
	v8 = 87
	v13 = 87
	libc.MemCpy(unsafe.Pointer(v5), unsafe.Pointer(v6), int(int32(v7)&3))
	for {
		if libc.StrCaseCmp(*(**byte)(memmap.PtrOff(0x587000, uint32(int32(v13)*4)+0x3A438)), &v20[0]) == 0 {
			break
		}
		v13 = uint8(func() int8 {
			p := &v8
			*p++
			return *p
		}())
		if int32(uint8(v8)) > 115 {
			break
		}
	}
	v9 = int8(int32(v8) + 39)
	switch uint8(int8(int32(v8) + 39)) {
	case 129:
		fallthrough
	case 130:
		fallthrough
	case 131:
		fallthrough
	case 132:
		fallthrough
	case 133:
		fallthrough
	case 134:
		fallthrough
	case 135:
		fallthrough
	case 136:
		fallthrough
	case 137:
		fallthrough
	case 138:
		fallthrough
	case 139:
		fallthrough
	case 150:
		fallthrough
	case 154:
		fallthrough
	case 160:
		fallthrough
	case 163:
		nox_xxx_netSendPointFx_522FF0(int8(int32(v8)+39), &v17)
		result = 0
	case 140:
		fallthrough
	case 143:
		fallthrough
	case 144:
		fallthrough
	case 145:
		fallthrough
	case 148:
		fallthrough
	case 149:
		v19[0] = nox_script_builtinGetF40() + int(v14)
		v19[1] = nox_script_builtinGetF44() + int(v15)
		v19[2] = nox_script_builtinGetF40() + int(*(*float32)(unsafe.Pointer(&v0)))
		v19[3] = nox_script_builtinGetF44() + int(*(*float32)(unsafe.Pointer(&v16)))
		nox_xxx_netSendRayFx_5232F0(v9, int32(uintptr(unsafe.Pointer(&v19[0]))))
		result = 0
	case 147:
		v12 = nox_float2int8(*(*float32)(unsafe.Pointer(&v0)))
		nox_xxx_netSparkExplosionFx_5231B0(&v17.field_0, v12)
		result = 0
	case 151:
		v11 = int(*(*float32)(unsafe.Pointer(&v0)))
		nox_xxx_earthquakeSend_4D9110(&v17.field_0, v11)
		result = 0
	case 152:
		nox_xxx_netSendFxGreenBolt_523790((*int4)(unsafe.Pointer(&v19[0])), 30)
		result = 0
	case 162:
		v19[0] = nox_script_builtinGetF40() + int(v14)
		v19[1] = nox_script_builtinGetF44() + int(v15)
		v19[2] = nox_script_builtinGetF40() + int(*(*float32)(unsafe.Pointer(&v0)))
		v19[3] = nox_script_builtinGetF44() + int(*(*float32)(unsafe.Pointer(&v16)))
		nox_xxx_netSendVampFx_523270(v9, (*int16)(unsafe.Pointer(&v19[0])), 100)
		result = 0
	default:
		result = 0
	}
	return result
}
func nox_script_SetOwner_514490() int32 {
	var (
		v0 int32
		v1 int32
		v2 int32
		v3 int32
	)
	v0 = nox_script_pop()
	v1 = nox_script_pop()
	v2 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v1))))
	v3 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v0))))
	nox_xxx_unitSetOwner_4EC290((*nox_object_t)(unsafe.Pointer(uintptr(v2))), (*nox_object_t)(unsafe.Pointer(uintptr(v3))))
	return 0
}
func nox_script_GroupSetOwner_5144C0() int32 {
	var (
		v0 int32
		v1 int32
		v2 int32
		v3 int32
		i  *int32
		v5 int32
	)
	v0 = nox_script_pop()
	v1 = nox_script_pop()
	v2 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v1))))
	v3 = nox_server_scriptGetGroup_57C0A0(v0)
	if v3 != 0 {
		for i = *(**int32)(unsafe.Pointer(uintptr(v3 + 84))); i != nil; i = (*int32)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(i), unsafe.Sizeof(int32(0))*2))))) {
			v5 = nox_xxx_netGetUnitByExtent_4ED020(*i)
			if v5 != 0 {
				nox_xxx_unitSetOwner_4EC290((*nox_object_t)(unsafe.Pointer(uintptr(v2))), (*nox_object_t)(unsafe.Pointer(uintptr(v5))))
			}
		}
	}
	return 0
}
func nox_script_SetOwnerGroup_514510() int32 {
	var (
		v0 int32
		v1 int32
		v2 int32
		v3 int32
		v4 *int32
		i  int32
		v6 int32
	)
	v0 = nox_script_pop()
	v1 = nox_script_pop()
	v2 = nox_server_scriptGetGroup_57C0A0(v1)
	if v2 != 0 {
		v3 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v0))))
		v4 = *(**int32)(unsafe.Pointer(uintptr(v2 + 84)))
		for i = v3; v4 != nil; v4 = (*int32)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(int32(0))*2))))) {
			v6 = nox_xxx_netGetUnitByExtent_4ED020(*v4)
			if v6 != 0 {
				nox_xxx_unitSetOwner_4EC290((*nox_object_t)(unsafe.Pointer(uintptr(v6))), (*nox_object_t)(unsafe.Pointer(uintptr(i))))
			}
		}
	}
	return 0
}
func nox_script_builtin_514570() int32 {
	var (
		v0 int32
		v1 int32
		v2 int32
		v3 int32
		v4 int32
		i  *int32
		v6 int32
		j  *int32
		v8 int32
	)
	v0 = nox_script_pop()
	v1 = nox_script_pop()
	v2 = nox_server_scriptGetGroup_57C0A0(v1)
	v3 = nox_server_scriptGetGroup_57C0A0(v0)
	v4 = v3
	if v2 != 0 {
		if v3 != 0 {
			for i = *(**int32)(unsafe.Pointer(uintptr(v2 + 84))); i != nil; i = (*int32)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(i), unsafe.Sizeof(int32(0))*2))))) {
				v6 = nox_xxx_netGetUnitByExtent_4ED020(*i)
				if v6 != 0 {
					for j = *(**int32)(unsafe.Pointer(uintptr(v4 + 84))); j != nil; j = (*int32)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(j), unsafe.Sizeof(int32(0))*2))))) {
						v8 = nox_xxx_netGetUnitByExtent_4ED020(*j)
						if v8 != 0 {
							nox_xxx_unitSetOwner_4EC290((*nox_object_t)(unsafe.Pointer(uintptr(v6))), (*nox_object_t)(unsafe.Pointer(uintptr(v8))))
						}
					}
				}
			}
		}
	}
	return 0
}
func nox_script_builtin_5145F0() int32 {
	var (
		v0 int32
		v1 int32
		v2 int32
		v3 int32
		v4 int32
	)
	v0 = nox_script_pop()
	v1 = nox_script_pop()
	v2 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v0))))
	v3 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v1))))
	v4 = nox_xxx_unitHasThatParent_4EC4F0(v3, v2)
	nox_script_push(v4)
	return 0
}
func nox_script_builtin_514630() int32 {
	var (
		v0     int32
		v1     int32
		result int32
		v3     int32
		v4     int32
		v5     *int32
		v6     int32
		v7     int32
		v8     int32
	)
	v0 = nox_script_pop()
	v1 = nox_script_pop()
	result = nox_server_scriptGetGroup_57C0A0(v0)
	v3 = result
	if result != 0 {
		v4 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v1))))
		v5 = *(**int32)(unsafe.Pointer(uintptr(v3 + 84)))
		v6 = v4
		v7 = 1
		if v5 != nil {
			for {
				v8 = nox_xxx_netGetUnitByExtent_4ED020(*v5)
				if v8 != 0 {
					if nox_xxx_unitHasThatParent_4EC4F0(v6, v8) == 0 {
						break
					}
				}
				v5 = (*int32)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(int32(0))*2)))))
				if v5 == nil {
					nox_script_push(1)
					return 0
				}
			}
			v7 = 0
		}
		nox_script_push(v7)
		result = 0
	}
	return result
}
func nox_script_builtin_5146B0() int32 {
	var (
		v0     int32
		v1     int32
		result int32
		v3     int32
		v4     int32
		v5     *int32
		v6     int32
		v7     int32
		v8     int32
	)
	v0 = nox_script_pop()
	v1 = nox_script_pop()
	result = nox_server_scriptGetGroup_57C0A0(v1)
	v3 = result
	if result != 0 {
		v4 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v0))))
		v5 = *(**int32)(unsafe.Pointer(uintptr(v3 + 84)))
		v6 = v4
		v7 = 1
		if v5 != nil {
			for {
				v8 = nox_xxx_netGetUnitByExtent_4ED020(*v5)
				if v8 != 0 {
					if nox_xxx_unitHasThatParent_4EC4F0(v8, v6) == 0 {
						break
					}
				}
				v5 = (*int32)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(int32(0))*2)))))
				if v5 == nil {
					nox_script_push(1)
					return 0
				}
			}
			v7 = 0
		}
		nox_script_push(v7)
		result = 0
	}
	return result
}
func nox_script_builtin_514730() int32 {
	var (
		v0  int32
		v1  int32
		v2  int32
		v3  int32
		v4  *int32
		i   int32
		v6  int32
		v7  *int32
		v8  int32
		v10 int32
	)
	v0 = nox_script_pop()
	v1 = nox_script_pop()
	v2 = nox_server_scriptGetGroup_57C0A0(v0)
	v3 = nox_server_scriptGetGroup_57C0A0(v1)
	v10 = v3
	if v2 == 0 || v3 == 0 {
		return 0
	}
	v4 = *(**int32)(unsafe.Pointer(uintptr(v2 + 84)))
	for i = 1; v4 != nil; v4 = (*int32)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(int32(0))*2))))) {
		if i == 0 {
			break
		}
		v6 = nox_xxx_netGetUnitByExtent_4ED020(*v4)
		if v6 != 0 {
			v7 = *(**int32)(unsafe.Pointer(uintptr(v10 + 84)))
			if v7 != nil {
				for {
					v8 = nox_xxx_netGetUnitByExtent_4ED020(*v7)
					if v8 != 0 {
						if nox_xxx_unitHasThatParent_4EC4F0(v8, v6) == 0 {
							break
						}
					}
					v7 = (*int32)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(v7), unsafe.Sizeof(int32(0))*2)))))
					if v7 == nil {
						goto LABEL_12
					}
				}
				i = 0
			}
		}
	LABEL_12:
	}
	nox_script_push(i)
	return 0
}
func nox_script_ClearOwner_5147E0() int32 {
	var (
		v0 int32
		v1 int32
	)
	v0 = nox_script_pop()
	v1 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v0))))
	nox_xxx_unitClearOwner_4EC300((*nox_object_t)(unsafe.Pointer(uintptr(v1))))
	return 0
}
func nox_script_Waypoint_514800() int32 {
	var (
		v0 *int32
		v2 [76]byte
	)
	libc.StrNCpy(&v2[0], nox_script_getString_512E40(nox_script_pop()), int(unsafe.Sizeof([76]byte{})-1))
	libc.StrNCat(&v2[0], &nox_script_builtin_buf_xxx[0], int(unsafe.Sizeof([76]byte{})-1))
	v0 = (*int32)(unsafe.Pointer(nox_xxx_waypointByName_579E30(&v2[0])))
	if v0 != nil {
		nox_script_push(*v0)
	} else {
		nox_script_push(0)
	}
	return 0
}
func nox_script_GetWaypointGroup_5148A0() int32 {
	var (
		v0 int32
		v2 [76]byte
	)
	libc.StrNCpy(&v2[0], nox_script_getString_512E40(nox_script_pop()), int(unsafe.Sizeof([76]byte{})-1))
	libc.StrNCat(&v2[0], &nox_script_builtin_buf_xxx[0], int(unsafe.Sizeof([76]byte{})-1))
	v0 = int32(uintptr(nox_server_scriptGetMapGroupByName_57C280(&v2[0], 1)))
	if v0 != 0 {
		nox_script_push(int32(*(*uint32)(unsafe.Pointer(uintptr(v0 + 4)))))
	} else {
		nox_script_push(0)
	}
	return 0
}
func nox_script_GetObjectGroup_514940() int32 {
	var (
		v0 int32
		v2 [76]byte
	)
	libc.StrNCpy(&v2[0], nox_script_getString_512E40(nox_script_pop()), int(unsafe.Sizeof([76]byte{})-1))
	libc.StrNCat(&v2[0], &nox_script_builtin_buf_xxx[0], int(unsafe.Sizeof([76]byte{})-1))
	v0 = int32(uintptr(nox_server_scriptGetMapGroupByName_57C280(&v2[0], 0)))
	if v0 != 0 {
		nox_script_push(int32(*(*uint32)(unsafe.Pointer(uintptr(v0 + 4)))))
	} else {
		nox_script_push(0)
	}
	return 0
}
func nox_script_GetWallGroup_5149E0() int32 {
	var (
		v0 int32
		v2 [76]byte
	)
	libc.StrNCpy(&v2[0], nox_script_getString_512E40(nox_script_pop()), int(unsafe.Sizeof([76]byte{})-1))
	libc.StrNCat(&v2[0], &nox_script_builtin_buf_xxx[0], int(unsafe.Sizeof([76]byte{})-1))
	v0 = int32(uintptr(nox_server_scriptGetMapGroupByName_57C280(&v2[0], 2)))
	if v0 != 0 {
		nox_script_push(int32(*(*uint32)(unsafe.Pointer(uintptr(v0 + 4)))))
	} else {
		nox_script_push(0)
	}
	return 0
}
func nox_script_Pop2_74_514BA0() int32 {
	nox_script_pop()
	nox_script_pop()
	return 0
}
func nox_script_RemoveChat_514BB0() int32 {
	var (
		v0 int32
		v1 *uint32
	)
	v0 = nox_script_pop()
	v1 = (*uint32)(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v0)))
	if v1 != nil {
		nox_xxx_netKillChat_528D00(v1)
	}
	return 0
}
func nox_script_NoChatAll_514BD0() int32 {
	nox_xxx_destroyEveryChatMB_528D60()
	return 0
}
func nox_script_JournalQuest_514BE0() int32 {
	var (
		v0 int32
		v1 int32
	)
	v0 = nox_script_pop()
	v1 = nox_script_pop()
	nox_xxx_journalQuestSet_500540(nox_script_getString_512E40(v0), v1)
	return 0
}
func nox_script_JournalQuestBool_514C10() int32 {
	var (
		v0 int32
		v1 int32
	)
	v0 = nox_script_pop()
	v1 = nox_script_pop()
	nox_xxx_journalQuestSetBool_5006B0(nox_script_getString_512E40(v0), v1)
	return 0
}
func nox_script_JournalGetQuest_514C40() int32 {
	var (
		v0 int32
		v1 int32
	)
	v0 = nox_script_pop()
	v1 = sub_500750(nox_script_getString_512E40(v0))
	nox_script_push(v1)
	return 0
}
func nox_script_JournalGetQuestBool_514C60() int32 {
	var (
		v0 int32
		v2 int32
	)
	v0 = nox_script_pop()
	*(*float32)(unsafe.Pointer(&v2)) = float32(sub_500770(nox_script_getString_512E40(v0)))
	nox_script_push(v2)
	return 0
}
func nox_script_RestoreQuestStatus_514C90() int32 {
	var v0 int32
	v0 = nox_script_pop()
	sub_5007E0(nox_script_getString_512E40(v0))
	return 0
}
func nox_script_ObjIsTrigger_514CB0() int32 {
	var (
		v0 int32
		v1 int32
		v2 int32
	)
	v0 = nox_script_pop()
	v1 = 0
	v2 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v0))))
	if v2 != 0 && nox_script_get_trigger() != nil && *(*uint32)(unsafe.Pointer(uintptr(v2 + 44))) == *(*uint32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(nox_script_get_trigger())), 44)))) {
		v1 = 1
	}
	nox_script_push(v1)
	return 0
}
func nox_script_builtin_514CF0() int32 {
	var (
		v0 int32
		v1 int32
		v2 int32
	)
	v0 = nox_script_pop()
	v1 = 0
	v2 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v0))))
	if v2 != 0 && nox_script_get_caller() != nil && *(*uint32)(unsafe.Pointer(uintptr(v2 + 44))) == *(*uint32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(nox_script_get_caller())), 44)))) {
		v1 = 1
	}
	nox_script_push(v1)
	return 0
}
func nox_script_GetTrigger_514D30() int32 {
	if nox_script_get_trigger() != nil {
		nox_script_push(int32(*(*uint32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(nox_script_get_trigger())), 44))))))
	} else {
		nox_script_push(0)
	}
	return 0
}
func nox_script_GetCaller_514D60() int32 {
	if nox_script_get_caller() != nil {
		nox_script_push(int32(*(*uint32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(nox_script_get_caller())), 44))))))
	} else {
		nox_script_push(0)
	}
	return 0
}
func nox_script_SetDialog_514D90() int32 {
	var (
		v0 int32
		v1 int32
		v2 int32
		v3 int32
		v4 int32
		v6 int8
	)
	v0 = nox_script_pop()
	v1 = nox_script_pop()
	v2 = nox_script_pop()
	v3 = nox_script_pop()
	v6 = nox_xxx_scriptGetDialogIdx_548F70(nox_script_getString_512E40(v2))
	v4 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v3))))
	if v4 != 0 {
		nox_xxx_scriptSetDialog_548C80(v4, v6, v1, v0)
	}
	return 0
}
func nox_script_CancelDialog_514DF0() int32 {
	var (
		v0 int32
		v1 int32
		v2 int32
	)
	v0 = nox_script_pop()
	v1 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v0))))
	if v1 != 0 {
		v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 748))))
		if int32(*(*uint8)(unsafe.Pointer(uintptr(v1 + 8))))&2 != 0 {
			*(*uint32)(unsafe.Pointer(uintptr(v2 + 2096))) = math.MaxUint32
			*(*uint32)(unsafe.Pointer(uintptr(v2 + 2100))) = math.MaxUint32
			nox_xxx_unitUnsetXStatus_4E4780(v1, 16)
		}
	}
	return 0
}
func nox_script_DialogPortrait_514E30() int32 {
	var (
		v0 int32
		v1 int32
		v2 int32
	)
	v0 = nox_script_pop()
	v1 = nox_script_pop()
	v2 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v1))))
	if v2 != 0 && int32(*(*uint8)(unsafe.Pointer(uintptr(v2 + 8))))&2 != 0 {
		libc.StrCpy((*byte)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 748)))+2106))), nox_script_getString_512E40(v0))
	}
	return 0
}
func nox_script_TellStory_514E90() int32 {
	var (
		v0 int32
		v1 int32
		v2 int32
		v4 *byte
	)
	v0 = nox_script_pop()
	v1 = nox_script_pop()
	v4 = nox_script_getString_512E40(v0)
	v2 = nox_xxx_utilFindSound_40AF50(nox_script_getString_512E40(v1))
	nox_xxx_startShopDialog_548DE0(int32(uintptr(nox_script_get_caller())), int32(uintptr(nox_script_get_trigger())), v2, v4)
	return 0
}
func nox_script_ForceDialog_514ED0() int32 {
	var (
		v0 int32
		v1 int32
		v2 int32
		v3 int32
	)
	v0 = nox_script_pop()
	v1 = nox_script_pop()
	v2 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v0))))
	v3 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v1))))
	if v2 != 0 && v3 != 0 {
		nox_xxx_script_forcedialog_548CD0(v2, v3)
	}
	return 0
}
func nox_script_builtin_514F10() int32 {
	var (
		v0  int32
		v1  int32
		v2  int32
		v3  int32
		v4  int32
		v5  int32
		v6  int32
		v7  int32
		v9  float2
		v10 [3]int32
	)
	v0 = nox_script_pop()
	v1 = nox_script_pop()
	v2 = nox_script_pop()
	v3 = things.SpellIndex(nox_script_getString_512E40(v2))
	if v3 != 0 {
		v4 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v1))))
		v5 = v4
		if v4 != 0 {
			if (*(*uint32)(unsafe.Pointer(uintptr(v4 + 16))) & 32800) == 0 {
				v6 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v0))))
				v7 = v6
				if v6 != 0 {
					v9.field_0 = *(*float32)(unsafe.Pointer(uintptr(v6 + 56))) - *(*float32)(unsafe.Pointer(uintptr(v5 + 56)))
					v9.field_4 = *(*float32)(unsafe.Pointer(uintptr(v6 + 60))) - *(*float32)(unsafe.Pointer(uintptr(v5 + 60)))
					*(*uint16)(unsafe.Pointer(uintptr(v5 + 124))) = uint16(int16(nox_xxx_math_509ED0(&v9)))
					v10[0] = v7
					v10[1] = int32(*(*uint32)(unsafe.Pointer(uintptr(v7 + 56))))
					v10[2] = int32(*(*uint32)(unsafe.Pointer(uintptr(v7 + 60))))
					nox_xxx_castSpellByUser_4FDD20(v3, (*nox_object_t)(unsafe.Pointer(uintptr(v5))), unsafe.Pointer(&v10[0]))
				}
			}
		}
	}
	return 0
}
func nox_script_builtin_514FC0() int32 {
	var (
		v0 int32
		v1 int32
		v2 int32
		v3 int32
		v4 *uint32
		v6 float32
		v7 float32
		v8 float2
		v9 [3]float32
	)
	v7 = nox_script_popf()
	v6 = nox_script_popf()
	v0 = nox_script_pop()
	v1 = nox_script_pop()
	v2 = things.SpellIndex(nox_script_getString_512E40(v1))
	if v2 != 0 {
		v3 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v0))))
		v4 = (*uint32)(unsafe.Pointer(uintptr(v3)))
		if v3 != 0 {
			v8.field_0 = v6 - *(*float32)(unsafe.Pointer(uintptr(v3 + 56)))
			v8.field_4 = v7 - *(*float32)(unsafe.Pointer(uintptr(v3 + 60)))
			*(*uint16)(unsafe.Pointer(uintptr(v3 + 124))) = uint16(int16(nox_xxx_math_509ED0(&v8)))
			v9[0] = 0.0
			v9[1] = v6
			v9[2] = v7
			nox_xxx_castSpellByUser_4FDD20(v2, (*nox_object_t)(unsafe.Pointer(v4)), unsafe.Pointer(&v9[0]))
		}
	}
	return 0
}
func nox_script_builtin_515060() int32 {
	var (
		v0 int32
		v1 int32
		v2 int32
		v3 int32
		v5 float2
		v6 float2
		v7 [3]int32
	)
	v0 = nox_script_pop()
	*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5.field_4))), unsafe.Sizeof(uint32(0))*0)) = uint32(nox_script_pop())
	*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v5.field_0))), unsafe.Sizeof(uint32(0))*0)) = uint32(nox_script_pop())
	v1 = nox_script_pop()
	v2 = things.SpellIndex(nox_script_getString_512E40(v1))
	if v2 != 0 {
		v3 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v0))))
		if v3 != 0 {
			v6 = v5
			nox_xxx_unitMove_4E7010((*nox_object_t)(nox_xxx_imagCasterUnit_1569664), &v6)
			v6.field_0 = *(*float32)(unsafe.Pointer(uintptr(v3 + 56))) - v5.field_0
			v6.field_4 = *(*float32)(unsafe.Pointer(uintptr(v3 + 60))) - v5.field_4
			*(*uint16)(unsafe.Pointer(uintptr(uint32(uintptr(nox_xxx_imagCasterUnit_1569664)) + 124))) = uint16(int16(nox_xxx_math_509ED0(&v6)))
			v7[0] = v3
			v7[1] = int32(*(*uint32)(unsafe.Pointer(uintptr(v3 + 56))))
			v7[2] = int32(*(*uint32)(unsafe.Pointer(uintptr(v3 + 60))))
			nox_xxx_castSpellByUser_4FDD20(v2, (*nox_object_t)(nox_xxx_imagCasterUnit_1569664), unsafe.Pointer(&v7[0]))
		}
	}
	return 0
}
func nox_script_cast_515130() int32 {
	var (
		v0 int32
		v1 int32
		v3 float2
		v4 float32
		v5 float32
		v6 float2
		v7 [3]float32
	)
	v5 = nox_script_popf()
	v4 = nox_script_popf()
	*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v3.field_4))), unsafe.Sizeof(uint32(0))*0)) = uint32(nox_script_pop())
	*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v3.field_0))), unsafe.Sizeof(uint32(0))*0)) = uint32(nox_script_pop())
	v0 = nox_script_pop()
	v1 = things.SpellIndex(nox_script_getString_512E40(v0))
	if v1 != 0 {
		v6 = v3
		nox_xxx_unitMove_4E7010((*nox_object_t)(nox_xxx_imagCasterUnit_1569664), &v6)
		v6.field_0 = v4 - v3.field_0
		v6.field_4 = v5 - v3.field_4
		*(*uint16)(unsafe.Pointer(uintptr(uint32(uintptr(nox_xxx_imagCasterUnit_1569664)) + 124))) = uint16(int16(nox_xxx_math_509ED0(&v6)))
		v7[1] = v4
		v7[0] = 0.0
		v7[2] = v5
		nox_xxx_castSpellByUser_4FDD20(v1, (*nox_object_t)(nox_xxx_imagCasterUnit_1569664), unsafe.Pointer(&v7[0]))
	}
	return 0
}
func nox_script_UnBlind_515200() int32 {
	nox_gameDisableMapDraw_5d4594_2650672 = 0
	nox_client_screenFadeXxx_44DB30(25, 0, nil)
	return 0
}
func nox_script_Blind_515220() int32 {
	nox_client_screenFadeTimeout_44DAB0(25, 0, sub_44E000)
	return 0
}
func nox_xxx_WideScreenDo_515240(enable bool) {
	var (
		v0 int32
		v3 int32
		v4 int32
		v5 int32
		v6 int32
		i  int32
		v9 int32
	)
	if nox_script_objTelekinesisHand == nil {
		nox_script_objTelekinesisHand = unsafe.Pointer(uintptr(nox_xxx_getNameId_4E3AA0(CString("TelekinesisHand"))))
	}
	if !enable {
		if sub_44DCA0(16, 30) != 0 {
			sub_477530(0)
		}
		for i = int32(uintptr(unsafe.Pointer(noxServer.firstServerObject()))); i != 0; i = int32(uintptr(unsafe.Pointer((*nox_object_t)(unsafe.Pointer(uintptr(i))).Next()))) {
			if unsafe.Pointer(uintptr(*(*uint16)(unsafe.Pointer(uintptr(i + 4))))) == nox_script_objTelekinesisHand {
				v9 = int32(*(*uint32)(unsafe.Pointer(uintptr(i + 16))))
				if v9&64 != 0 {
					*(*uint32)(unsafe.Pointer(uintptr(i + 16))) = uint32(v9) & 0xFFFFFFBF
				}
			}
		}
		return
	}
	v0 = sub_44DD00()
	if sub_44DC40(16, 10) != 0 {
		sub_477530(1)
	}
	if v0 != 0 {
		return
	}
	if *memmap.PtrUint32(6112660, 2386876) == 0 {
		*memmap.PtrUint32(6112660, 2386876) = uint32(nox_xxx_getNameId_4E3AA0(CString("ToxicCloud")))
		*memmap.PtrUint32(6112660, 2386880) = uint32(nox_xxx_getNameId_4E3AA0(CString("SmallToxicCloud")))
		*memmap.PtrUint32(6112660, 2386884) = uint32(nox_xxx_getNameId_4E3AA0(CString("Meteor")))
		*memmap.PtrUint32(6112660, 2386888) = uint32(nox_xxx_getNameId_4E3AA0(CString("SmallFist")))
		*memmap.PtrUint32(6112660, 2386892) = uint32(nox_xxx_getNameId_4E3AA0(CString("MediumFist")))
		*memmap.PtrUint32(6112660, 2386896) = uint32(nox_xxx_getNameId_4E3AA0(CString("LargeFist")))
		*memmap.PtrUint32(6112660, 2386900) = uint32(nox_xxx_getNameId_4E3AA0(CString("Pixie")))
	}
	var obj *nox_object_t = firstServerObjectUpdatable2()
	if obj != nil {
		for {
			{
				var v2 *nox_object_t = obj.Next()
				if uint32(obj.typ_ind) != *memmap.PtrUint32(6112660, 2386900) {
					nox_xxx_delayedDeleteObject_4E5CC0(obj)
				}
				obj = v2
			}
			if obj == nil {
				break
			}
		}
	}
	v3 = int32(uintptr(unsafe.Pointer(noxServer.firstServerObject())))
	if v3 == 0 {
		return
	}
	for {
		v4 = int32(uintptr(unsafe.Pointer((*nox_object_t)(unsafe.Pointer(uintptr(v3))).Next())))
		if *(*uint32)(unsafe.Pointer(uintptr(v3 + 508))) != 0 && ((func() bool {
			v5 = int32(*(*uint16)(unsafe.Pointer(uintptr(v3 + 4))))
			return uint32(uint16(int16(v5))) == *memmap.PtrUint32(6112660, 2386876)
		}()) || uint32(v5) == *memmap.PtrUint32(6112660, 2386880) || uint32(v5) == *memmap.PtrUint32(6112660, 2386884) || uint32(v5) == *memmap.PtrUint32(6112660, 2386888) || uint32(v5) == *memmap.PtrUint32(6112660, 2386892) || uint32(v5) == *memmap.PtrUint32(6112660, 2386896)) {
			nox_xxx_delayedDeleteObject_4E5CC0((*nox_object_t)(unsafe.Pointer(uintptr(v3))))
		} else {
			if unsafe.Pointer(uintptr(*(*uint16)(unsafe.Pointer(uintptr(v3 + 4))))) == nox_script_objTelekinesisHand {
				v6 = int32(*(*uint32)(unsafe.Pointer(uintptr(v3 + 16))))
				if (v6 & 64) == 0 {
					*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v6))), 0)) = uint8(int8(v6 | 64))
					*(*uint32)(unsafe.Pointer(uintptr(v3 + 16))) = uint32(v6)
				}
			}
			nox_xxx_spellCancelDurSpell_4FEB10(132, (*nox_object_t)(unsafe.Pointer(uintptr(v3))))
			nox_xxx_spellCancelDurSpell_4FEB10(56, (*nox_object_t)(unsafe.Pointer(uintptr(v3))))
			nox_xxx_spellCancelDurSpell_4FEB10(24, (*nox_object_t)(unsafe.Pointer(uintptr(v3))))
			nox_xxx_spellCancelDurSpell_4FEB10(43, (*nox_object_t)(unsafe.Pointer(uintptr(v3))))
			nox_xxx_spellCancelDurSpell_4FEB10(22, (*nox_object_t)(unsafe.Pointer(uintptr(v3))))
			nox_xxx_spellCancelDurSpell_4FEB10(31, (*nox_object_t)(unsafe.Pointer(uintptr(v3))))
			nox_xxx_spellCancelDurSpell_4FEB10(35, (*nox_object_t)(unsafe.Pointer(uintptr(v3))))
			nox_xxx_spellCancelDurSpell_4FEB10(8, (*nox_object_t)(unsafe.Pointer(uintptr(v3))))
			nox_xxx_spellCancelDurSpell_4FEB10(9, (*nox_object_t)(unsafe.Pointer(uintptr(v3))))
			nox_xxx_spellCancelDurSpell_4FEB10(4, (*nox_object_t)(unsafe.Pointer(uintptr(v3))))
			nox_xxx_spellCancelDurSpell_4FEB10(115, (*nox_object_t)(unsafe.Pointer(uintptr(v3))))
			nox_xxx_spellCancelDurSpell_4FEB10(21, (*nox_object_t)(unsafe.Pointer(uintptr(v3))))
			nox_xxx_spellCancelDurSpell_4FEB10(59, (*nox_object_t)(unsafe.Pointer(uintptr(v3))))
			nox_xxx_spellCancelDurSpell_4FEB10(75, (*nox_object_t)(unsafe.Pointer(uintptr(v3))))
		}
		v3 = v4
		if v4 == 0 {
			break
		}
	}
	return
}
func nox_script_WideScreen_515240() int32 {
	nox_xxx_WideScreenDo_515240(nox_script_pop() == 1)
	return 0
}
func nox_script_GetElevatorStat_5154A0() int32 {
	var (
		v0 int32
		v1 int32
		v2 int32
		v3 int32
	)
	v0 = -1
	v1 = nox_script_pop()
	v2 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v1))))
	if v2 != 0 {
		v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 8))))
		if v3&0x4000 != 0 {
			v0 = int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 748))) + 12))))
		}
	}
	nox_script_push(v0)
	return 0
}
func nox_script_builtin_5154E0() int32 {
	var (
		v0 int16
		v1 int32
		v2 int32
		v3 int32
	)
	v0 = int16(nox_script_pop())
	v1 = nox_script_pop()
	v2 = nox_script_pop()
	if v2 != 0 {
		v3 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v2))))
		if v3 != 0 {
			nox_xxx_comJournalEntryAdd_427500(v3, nox_script_getString_512E40(v1), v0)
			if int32(v0)&11 != 0 {
				nox_xxx_aud_501960(903, (*nox_object_t)(unsafe.Pointer(uintptr(v3))), 0, 0)
				return 0
			}
		}
	} else {
		nox_xxx_comAddEntryAll_427550(nox_script_getString_512E40(v1), v0)
	}
	return 0
}
func nox_script_DeleteJournalEntry_515550() int32 {
	var (
		v0 int32
		v1 int32
		v2 int32
	)
	v0 = nox_script_pop()
	v1 = nox_script_pop()
	if v1 != 0 {
		v2 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v1))))
		if v2 != 0 {
			nox_xxx_comJournalEntryRemove_427630(v2, nox_script_getString_512E40(v0))
			return 0
		}
	} else {
		nox_xxx_comRemoveEntryAll_427680(nox_script_getString_512E40(v0))
	}
	return 0
}
func nox_script_builtin_5155A0() int32 {
	var (
		v0 int16
		v1 int32
		v2 int32
		v3 int32
	)
	v0 = int16(nox_script_pop())
	v1 = nox_script_pop()
	v2 = nox_script_pop()
	if v2 != 0 {
		v3 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v2))))
		if v3 != 0 {
			nox_xxx_comJournalEntryUpdate_427720(v3, nox_script_getString_512E40(v1), v0)
			return 0
		}
	} else {
		nox_xxx_comUpdateEntryAll_427770(nox_script_getString_512E40(v1), v0)
	}
	return 0
}
func nox_script_SetGuardPoint_515600() int32 {
	var (
		v0 int32
		v1 int32
		v3 int32
		v4 int32
		v5 int32
		v6 int32
		v7 int32
		v8 [5]int32
	)
	v7 = nox_script_pop()
	v6 = nox_script_pop()
	v5 = nox_script_pop()
	v4 = nox_script_pop()
	v3 = nox_script_pop()
	v0 = nox_script_pop()
	v8[0] = v3
	v8[1] = v4
	v8[2] = v5
	v8[3] = v6
	v8[4] = v7
	v1 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v0))))
	if v1 != 0 {
		nox_xxx_monsterGoPatrol_515680(v1, int32(uintptr(unsafe.Pointer(&v8[0]))))
	}
	return 0
}
func nox_script_builtin_515700() int32 {
	var (
		v0 int32
		v1 *uint8
		v3 int32
		v4 int32
		v5 int32
		v6 int32
		v7 int32
		v8 [5]int32
	)
	v7 = nox_script_pop()
	v6 = nox_script_pop()
	v5 = nox_script_pop()
	v4 = nox_script_pop()
	v3 = nox_script_pop()
	v0 = nox_script_pop()
	v8[1] = v4
	v8[3] = v6
	v8[0] = v3
	v8[2] = v5
	v8[4] = v7
	v1 = (*uint8)(unsafe.Pointer(uintptr(nox_server_scriptGetGroup_57C0A0(v0))))
	nox_server_scriptExecuteFnForEachGroupObj_502670(v1, 0, nox_xxx_monsterGoPatrol_515680, int32(uintptr(unsafe.Pointer(&v8[0]))))
	return 0
}
func nox_script_MonsterHunt_515780() int32 {
	var (
		v0 int32
		v1 int32
	)
	v0 = nox_script_pop()
	v1 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v0))))
	if v1 != 0 {
		nox_xxx_unitHunt_5157A0((*nox_object_t)(unsafe.Pointer(uintptr(v1))))
	}
	return 0
}
func nox_script_builtin_5157D0() int32 {
	var (
		v0 int32
		v1 *uint8
	)
	v0 = nox_script_pop()
	v1 = (*uint8)(unsafe.Pointer(uintptr(nox_server_scriptGetGroup_57C0A0(v0))))
	nox_server_scriptExecuteFnForEachGroupObj_502670(v1, 0, func(arg1 int32, arg2 int32) {
		nox_xxx_unitHunt_5157A0((*nox_object_t)(unsafe.Pointer(uintptr(arg1))))
	}, 0)
	return 0
}
func nox_script_Idle_515800() int32 {
	var (
		v0 int32
		v1 int32
	)
	v0 = nox_script_pop()
	v1 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v0))))
	if v1 != 0 {
		nox_xxx_unitIdle_515820((*nox_object_t)(unsafe.Pointer(uintptr(v1))))
	}
	return 0
}
func nox_script_GroupIdle_515850() int32 {
	var (
		v0 int32
		v1 *uint8
	)
	v0 = nox_script_pop()
	v1 = (*uint8)(unsafe.Pointer(uintptr(nox_server_scriptGetGroup_57C0A0(v0))))
	nox_server_scriptExecuteFnForEachGroupObj_502670(v1, 0, func(arg1 int32, arg2 int32) {
		nox_xxx_unitIdle_515820((*nox_object_t)(unsafe.Pointer(uintptr(arg1))))
	}, 0)
	return 0
}
func nox_script_Follow_515880() int32 {
	var (
		v0 int32
		v1 int32
		v2 int32
		v3 int32
	)
	v0 = nox_script_pop()
	v1 = nox_script_pop()
	v2 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v1))))
	if v2 != 0 {
		v3 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v0))))
		if v3 != 0 {
			nox_xxx_unitSetFollow_5158C0((*nox_object_t)(unsafe.Pointer(uintptr(v2))), (*nox_object_t)(unsafe.Pointer(uintptr(v3))))
		}
	}
	return 0
}
func nox_script_builtin_515910() int32 {
	var (
		v0 int32
		v1 int32
		v2 int32
		v3 *uint8
		v5 int32
	)
	v0 = nox_script_pop()
	v1 = nox_script_pop()
	v2 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v0))))
	if v2 != 0 {
		v5 = v2
		v3 = (*uint8)(unsafe.Pointer(uintptr(nox_server_scriptGetGroup_57C0A0(v1))))
		nox_server_scriptExecuteFnForEachGroupObj_502670(v3, 0, func(arg1 int32, arg2 int32) {
			nox_xxx_unitSetFollow_5158C0((*nox_object_t)(unsafe.Pointer(uintptr(arg1))), (*nox_object_t)(unsafe.Pointer(uintptr(arg2))))
		}, v5)
	}
	return 0
}
func nox_script_AgresssionLevel_515950() int32 {
	var (
		v0 int32
		v1 int32
		v3 int32
	)
	v3 = nox_script_pop()
	v0 = nox_script_pop()
	v1 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v0))))
	if v1 != 0 {
		nox_xxx_monsterSetAgressiveness_515980(v1, (*uint32)(unsafe.Pointer(&v3)))
	}
	return 0
}
func nox_script_builtin_5159B0() int32 {
	var (
		v0 int32
		v1 *uint8
		v3 int32
	)
	v3 = nox_script_pop()
	v0 = nox_script_pop()
	v1 = (*uint8)(unsafe.Pointer(uintptr(nox_server_scriptGetGroup_57C0A0(v0))))
	nox_server_scriptExecuteFnForEachGroupObj_502670(v1, 0, func(arg1 int32, arg2 int32) {
		nox_xxx_monsterSetAgressiveness_515980(arg1, (*uint32)(unsafe.Pointer(uintptr(arg2))))
	}, int32(uintptr(unsafe.Pointer(&v3))))
	return 0
}
func nox_script_HitLocation_5159E0() int32 {
	var (
		v0 int32
		v1 int32
		v3 float32
		v4 float32
		v5 float2
	)
	v4 = nox_script_popf()
	v3 = nox_script_popf()
	v0 = nox_script_pop()
	v1 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v0))))
	if v1 != 0 {
		v5.field_0 = v3
		v5.field_4 = v4
		nox_xxx_monsterActionMelee_515A30(v1, &v5)
	}
	return 0
}
func nox_script_builtin_515AE0() int32 {
	var (
		v0 int32
		v1 *uint8
		v3 float2
		v4 float2
	)
	*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v3.field_4))), unsafe.Sizeof(uint32(0))*0)) = uint32(nox_script_pop())
	*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v3.field_0))), unsafe.Sizeof(uint32(0))*0)) = uint32(nox_script_pop())
	v0 = nox_script_pop()
	v4 = v3
	v1 = (*uint8)(unsafe.Pointer(uintptr(nox_server_scriptGetGroup_57C0A0(v0))))
	nox_server_scriptExecuteFnForEachGroupObj_502670(v1, 0, func(arg1 int32, arg2 int32) {
		nox_xxx_monsterActionMelee_515A30(arg1, (*float2)(unsafe.Pointer(uintptr(arg2))))
	}, int32(uintptr(unsafe.Pointer(&v4))))
	return 0
}
func nox_script_ShootLocation_515B30() int32 {
	var (
		v0 int32
		v1 int32
		v3 [2]int32
		v4 [2]int32
	)
	v3[1] = nox_script_pop()
	v3[0] = nox_script_pop()
	v0 = nox_script_pop()
	v1 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v0))))
	if v1 != 0 {
		*(*uint64)(unsafe.Pointer(&v4[0])) = *(*uint64)(unsafe.Pointer(&v3[0]))
		nox_xxx_monsterMissileAttack_515B80(v1, (*uint32)(unsafe.Pointer(&v4[0])))
	}
	return 0
}
func nox_script_builtin_515BF0() int32 {
	var (
		v0 int32
		v1 *uint8
		v3 [2]int32
		v4 [2]int32
	)
	v3[1] = nox_script_pop()
	v3[0] = nox_script_pop()
	v0 = nox_script_pop()
	*(*uint64)(unsafe.Pointer(&v4[0])) = *(*uint64)(unsafe.Pointer(&v3[0]))
	v1 = (*uint8)(unsafe.Pointer(uintptr(nox_server_scriptGetGroup_57C0A0(v0))))
	nox_server_scriptExecuteFnForEachGroupObj_502670(v1, 0, func(arg1 int32, arg2 int32) {
		nox_xxx_monsterMissileAttack_515B80(arg1, (*uint32)(unsafe.Pointer(uintptr(arg2))))
	}, int32(uintptr(unsafe.Pointer(&v4[0]))))
	return 0
}
func nox_script_MonsterWayFlag_515C40(a1 int32) int32 {
	var (
		v1 int8
		v2 int32
		v3 int32
		v5 int8
	)
	v5 = int8(*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a1))), unsafe.Sizeof(int32(0))-1)))
	v1 = int8(nox_script_pop())
	v2 = nox_script_pop()
	v3 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v2))))
	if v3 != 0 {
		v5 = v1
		sub_515C80(v3, (*uint8)(unsafe.Pointer(&v5)))
	}
	return 0
}
func nox_script_builtin_515CB0() int32 {
	var (
		v0 int8
		v1 int32
		v2 *uint8
		v4 int8
	)
	v0 = int8(nox_script_pop())
	v1 = nox_script_pop()
	v4 = v0
	v2 = (*uint8)(unsafe.Pointer(uintptr(nox_server_scriptGetGroup_57C0A0(v1))))
	nox_server_scriptExecuteFnForEachGroupObj_502670(v2, 0, func(arg1 int32, arg2 int32) {
		sub_515C80(arg1, (*uint8)(unsafe.Pointer(uintptr(arg2))))
	}, int32(uintptr(unsafe.Pointer(&v4))))
	return 0
}
func nox_script_Attack_515CF0() int32 {
	var (
		v0 int32
		v1 int32
		v2 int32
		v3 int32
	)
	v0 = nox_script_pop()
	v1 = nox_script_pop()
	v2 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v1))))
	if v2 != 0 {
		v3 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v0))))
		nox_xxx_mobSetFightTarg_515D30(v2, v3)
	}
	return 0
}
func nox_script_builtin_515DB0() int32 {
	var (
		v0 int32
		v1 int32
		v2 *uint8
		v4 int32
	)
	v0 = nox_script_pop()
	v1 = nox_script_pop()
	v4 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v0))))
	v2 = (*uint8)(unsafe.Pointer(uintptr(nox_server_scriptGetGroup_57C0A0(v1))))
	nox_server_scriptExecuteFnForEachGroupObj_502670(v2, 0, nox_xxx_mobSetFightTarg_515D30, v4)
	return 0
}
func nox_script_SetIdleLevel_515DF0() int32 {
	var (
		v0 int32
		v1 int32
		v3 int32
	)
	v3 = nox_script_pop()
	v0 = nox_script_pop()
	v1 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v0))))
	if v1 != 0 {
		sub_515E20(v1, (*uint32)(unsafe.Pointer(&v3)))
	}
	return 0
}
func nox_script_GroupSetIdleHP_515E50() int32 {
	var (
		v0 int32
		v1 *uint8
		v3 int32
	)
	v3 = nox_script_pop()
	v0 = nox_script_pop()
	v1 = (*uint8)(unsafe.Pointer(uintptr(nox_server_scriptGetGroup_57C0A0(v0))))
	nox_server_scriptExecuteFnForEachGroupObj_502670(v1, 0, func(arg1 int32, arg2 int32) {
		sub_515E20(arg1, (*uint32)(unsafe.Pointer(uintptr(arg2))))
	}, int32(uintptr(unsafe.Pointer(&v3))))
	return 0
}
func nox_script_SetResumeLevel_515E80() int32 {
	var (
		v0 int32
		v1 int32
		v3 int32
	)
	v3 = nox_script_pop()
	v0 = nox_script_pop()
	v1 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v0))))
	if v1 != 0 {
		sub_515EB0(v1, (*uint32)(unsafe.Pointer(&v3)))
	}
	return 0
}
func nox_script_builtin_515EE0() int32 {
	var (
		v0 int32
		v1 *uint8
		v3 int32
	)
	v3 = nox_script_pop()
	v0 = nox_script_pop()
	v1 = (*uint8)(unsafe.Pointer(uintptr(nox_server_scriptGetGroup_57C0A0(v0))))
	nox_server_scriptExecuteFnForEachGroupObj_502670(v1, 0, func(arg1 int32, arg2 int32) {
		sub_515EB0(arg1, (*uint32)(unsafe.Pointer(uintptr(arg2))))
	}, int32(uintptr(unsafe.Pointer(&v3))))
	return 0
}
func nox_script_RunAway_515F10() int32 {
	var (
		v0 int32
		v1 int32
		v2 int32
		v3 int32
		v5 [2]int32
	)
	v0 = nox_script_pop()
	v1 = nox_script_pop()
	v2 = nox_script_pop()
	v5[0] = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v1))))
	if v5[0] != 0 {
		v3 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v2))))
		if v3 != 0 {
			v5[1] = v0
			nox_server_scriptFleeFrom_515F70(v3, (*uint32)(unsafe.Pointer(&v5[0])))
		}
	}
	return 0
}
func nox_script_monsterGroupFleeFrom_516000() int32 {
	var (
		v0 int32
		v1 int32
		v2 int32
		v3 *uint8
		v5 [2]int32
	)
	v0 = nox_script_pop()
	v1 = nox_script_pop()
	v2 = nox_script_pop()
	v5[0] = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v1))))
	if v5[0] != 0 {
		v5[1] = v0
		v3 = (*uint8)(unsafe.Pointer(uintptr(nox_server_scriptGetGroup_57C0A0(v2))))
		nox_server_scriptExecuteFnForEachGroupObj_502670(v3, 0, func(arg1 int32, arg2 int32) {
			nox_server_scriptFleeFrom_515F70(arg1, (*uint32)(unsafe.Pointer(uintptr(arg2))))
		}, int32(uintptr(unsafe.Pointer(&v5[0]))))
	}
	return 0
}
func nox_script_MonsterWait_516060() int32 {
	var (
		v0 int32
		v1 int32
		v3 int32
	)
	v3 = nox_script_pop()
	v0 = nox_script_pop()
	v1 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v0))))
	if v1 != 0 {
		sub_516090(v1, (*uint32)(unsafe.Pointer(&v3)))
	}
	return 0
}
func nox_script_builtin_5160F0() int32 {
	var (
		v0 int32
		v1 *uint8
		v3 int32
	)
	v3 = nox_script_pop()
	v0 = nox_script_pop()
	v1 = (*uint8)(unsafe.Pointer(uintptr(nox_server_scriptGetGroup_57C0A0(v0))))
	nox_server_scriptExecuteFnForEachGroupObj_502670(v1, 0, func(arg1 int32, arg2 int32) {
		sub_516090(arg1, (*uint32)(unsafe.Pointer(uintptr(arg2))))
	}, int32(uintptr(unsafe.Pointer(&v3))))
	return 0
}
func nox_script_builtin_516120() int32 {
	var (
		v0 int32
		v1 int32
		v2 int32
	)
	v0 = nox_script_pop()
	v1 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v0))))
	if v1 != 0 {
		v2 = sub_4FA6D0(v1)
		nox_script_push(v2)
	} else {
		nox_script_push(0)
	}
	return 0
}
func nox_script_builtin_516160() int32 {
	var (
		v0 int32
		v1 int32
		v2 int32
	)
	v0 = nox_script_pop()
	v1 = nox_script_pop()
	v2 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v1))))
	if v2 != 0 {
		sub_4FA620(v2, v0)
	}
	return 0
}
func nox_script_GiveExp_516190() int32 {
	var (
		v0 int32
		v1 int32
		v2 int32
	)
	v0 = nox_script_pop()
	v1 = nox_script_pop()
	v2 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v1))))
	if v2 != 0 {
		nox_xxx_plyrGiveExp_4EF3A0_exp_level(v2, *(*float32)(unsafe.Pointer(&v0)))
	}
	return 0
}
func nox_script_builtin_5161C0() int32 {
	var (
		v0 int32
		v1 int32
		v2 int32
		v3 int32
		v4 int32
	)
	v0 = nox_script_pop()
	v1 = nox_script_pop()
	v2 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v1))))
	v3 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v0))))
	v4 = 0
	if v2 != 0 && v3 != 0 && nox_xxx_unitIsEnemyTo_5330C0((*nox_object_t)(unsafe.Pointer(uintptr(v2))), (*nox_object_t)(unsafe.Pointer(uintptr(v3)))) != 0 {
		v4 = 1
	}
	nox_script_push(v4)
	return 0
}
func nox_script_getObjectType_516210() int32 {
	var (
		v0     int32
		v1     int32
		result int32
		v3     *byte
		v4     int32
		v5     *byte
		v6     *uint8
		v7     int32
	)
	v0 = nox_script_pop()
	v1 = nox_script_pop()
	v7 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v1))))
	if v7 != 0 {
		v3 = *(**byte)(memmap.PtrOff(0x587000, 237172))
		v4 = 1
		if *memmap.PtrUint32(0x587000, 237172) != 0 {
			v5 = nox_script_getString_512E40(v0)
			v6 = (*uint8)(memmap.PtrOff(0x587000, 237172))
			for libc.StrCmp(v3, v5) != 0 {
				v3 = (*byte)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v6))), unsafe.Sizeof(uint32(0))*1))))))
				v6 = (*uint8)(unsafe.Add(unsafe.Pointer(v6), 4))
				v4 *= 2
				if v3 == nil {
					return 0
				}
			}
			nox_script_push(bool2int((uint32(v4) & *(*uint32)(unsafe.Pointer(uintptr(v7 + 8)))) != 0))
		}
		result = 0
	} else {
		nox_script_push(0)
		result = 0
	}
	return result
}
func nox_script_builtin_5162D0() int32 {
	var (
		v0     int32
		result int32
		v2     *uint8
		v3     int32
		v4     *byte
		v5     bool
		i      **byte
		v7     *uint8
		v8     int32
		v9     int32
	)
	v8 = nox_script_pop()
	v0 = nox_script_pop()
	v9 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v0))))
	if v9 != 0 {
		v2 = (*uint8)(memmap.PtrOff(0x587000, 237304))
		if true {
			v7 = (*uint8)(memmap.PtrOff(0x587000, 237304))
			for i = (**byte)(memmap.PtrOff(0x587000, 237304)); ; i = (**byte)(unsafe.Add(unsafe.Pointer(i), unsafe.Sizeof((*byte)(nil))*32)) {
				v3 = 1
				v4 = *i
				if *i != nil {
					break
				}
			LABEL_8:
				v5 = (*uint8)(unsafe.Add(unsafe.Pointer(v7), 128)) == nil
				v2 = (*uint8)(unsafe.Add(unsafe.Pointer(v7), 128))
				v7 = (*uint8)(unsafe.Add(unsafe.Pointer(v7), 128))
				if v5 {
					return 0
				}
			}
			for libc.StrCmp(v4, nox_script_getString_512E40(v8)) != 0 {
				v4 = (*byte)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), unsafe.Sizeof(uint32(0))*1))))))
				v2 = (*uint8)(unsafe.Add(unsafe.Pointer(v2), 4))
				v3 *= 2
				if v4 == nil {
					goto LABEL_8
				}
			}
			nox_script_push(bool2int((uint32(v3) & *(*uint32)(unsafe.Pointer(uintptr(v9 + 12)))) != 0))
		}
		result = 0
	} else {
		nox_script_push(0)
		result = 0
	}
	return result
}
func nox_script_DialogResult_5163C0() int32 {
	var (
		v0 int32
		v1 int32
		v2 int32
	)
	v0 = nox_script_pop()
	v1 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v0))))
	if v1 != 0 {
		v2 = sub_548F40(v1)
		nox_script_push(v2)
	} else {
		nox_script_push(0)
	}
	return 0
}
func nox_script_ForceAutosave_516400() int32 {
	if noxflags.HasGame(noxflags.GameModeCoop) {
		sub_4DB130(CString("AUTOSAVE"))
		sub_4DB170(1, 0, 0)
	}
	return 0
}
func nox_script_Music_516430() int32 {
	var (
		v0 int32
		v2 [3]byte
		v3 int32
	)
	v0 = nox_script_pop()
	v3 = nox_script_pop()
	if noxflags.HasGame(noxflags.GameModeCoop) {
		sub_43D9B0(v3, v0)
	} else {
		v2[1] = byte(int8(v3))
		v2[2] = byte(int8(v0))
		v2[0] = 229
		nox_xxx_netSendPacket1_4E5390(math.MaxUint8, int32(uintptr(unsafe.Pointer(&v2[0]))), 3, 0, 1)
	}
	return 0
}
func nox_script_builtin_5164A0(this unsafe.Pointer) int32 {
	var v2 unsafe.Pointer
	v2 = this
	if noxflags.HasGame(noxflags.GameModeCoop) {
		sub_43DA80()
	} else {
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v2))), 0)) = 230
		nox_xxx_netSendPacket1_4E5390(math.MaxUint8, int32(uintptr(unsafe.Pointer(&v2))), 3, 0, 1)
	}
	return 0
}
func nox_script_builtin_5164E0(this unsafe.Pointer) int32 {
	var v2 unsafe.Pointer
	v2 = this
	if noxflags.HasGame(noxflags.GameModeCoop) {
		sub_43DAD0()
	} else {
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v2))), 0)) = 231
		nox_xxx_netSendPacket1_4E5390(math.MaxUint8, int32(uintptr(unsafe.Pointer(&v2))), 3, 0, 1)
	}
	return 0
}
func nox_script_ClearMusic_516520(this unsafe.Pointer) int32 {
	var v2 unsafe.Pointer
	v2 = this
	if noxflags.HasGame(noxflags.GameModeCoop) {
		sub_43D9B0(0, 0)
	} else {
		*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v2))), unsafe.Sizeof(uint16(0))*0)) = 229
		nox_xxx_netSendPacket1_4E5390(math.MaxUint8, int32(uintptr(unsafe.Pointer(&v2))), 3, 0, 1)
	}
	return 0
}
func nox_script_builtin_516600() int32 {
	var (
		i  int32
		v1 int32
		v2 int32
	)
	for i = int32(uintptr(unsafe.Pointer(nox_xxx_getFirstPlayerUnit_4DA7C0()))); i != 0; i = int32(uintptr(unsafe.Pointer(nox_xxx_getNextPlayerUnit_4DA7F0((*nox_object_t)(unsafe.Pointer(uintptr(i))))))) {
		if int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(i + 748))) + 276))) + 2064)))) == 31 {
			break
		}
	}
	sub_4277B0(i, 14)
	v1 = i.FirstItem()
	if v1 != 0 {
		for {
			v2 = nox_xxx_inventoryGetNext_4E7990(v1)
			if int32(*(*uint8)(unsafe.Pointer(uintptr(v1 + 8))))&64 != 0 {
				nox_xxx_delayedDeleteObject_4E5CC0((*nox_object_t)(unsafe.Pointer(uintptr(v1))))
			}
			v1 = v2
			if v2 == 0 {
				break
			}
		}
	}
	*memmap.PtrUint32(6112660, 2386832) = 1
	sub_5165D0()
	return 0
}
func nox_script_builtin_516680() int32 {
	*memmap.PtrUint32(6112660, 2386832) = 0
	sub_5165D0()
	return 0
}
func nox_script_builtin_5166A0() int32 {
	var (
		v0 *byte
		v1 int32
	)
	v0 = (*byte)(unsafe.Pointer(nox_common_playerInfoFromNum_417090(31)))
	v1 = 0
	if v0 != nil && *(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v0))), unsafe.Sizeof(uint32(0))*514))) + 748))) + 284))) != 0 {
		v1 = 1
	}
	nox_script_push(v1)
	return 0
}
func nox_script_PlayerIsTrading_5166E0() int32 {
	var (
		v0 *byte
		v1 int32
	)
	v0 = (*byte)(unsafe.Pointer(nox_common_playerInfoFromNum_417090(31)))
	v1 = 0
	if v0 != nil && *(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v0))), unsafe.Sizeof(uint32(0))*514))) + 748))) + 280))) != 0 {
		v1 = 1
	}
	nox_script_push(v1)
	return 0
}
func nox_script_builtin_516720() int32 {
	var (
		v0 *byte
		v1 int32
		v2 int32
		v3 int32
		v4 int32
	)
	v0 = (*byte)(unsafe.Pointer(nox_common_playerInfoFromNum_417090(31)))
	v1 = nox_script_pop()
	v2 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v1))))
	if v2 != 0 {
		v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 12))))
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v3))), 1)) |= 1
		*(*uint32)(unsafe.Pointer(uintptr(v2 + 12))) = uint32(v3)
		v4 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v0))), unsafe.Sizeof(uint32(0))*514))))
		if v4 != 0 {
			nox_xxx_unitSetOwner_4EC290((*nox_object_t)(unsafe.Pointer(uintptr(v4))), (*nox_object_t)(unsafe.Pointer(uintptr(v2))))
		}
	}
	return 0
}
func nox_script_builtin_516760() int32 {
	var (
		v0 int32
		v1 int32
		v2 int32
	)
	v0 = nox_script_pop()
	v1 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v0))))
	if v1 != 0 {
		v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 12))))
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v2))), 1)) &= 254
		*(*uint32)(unsafe.Pointer(uintptr(v1 + 12))) = uint32(v2)
		nox_xxx_unitClearOwner_4EC300((*nox_object_t)(unsafe.Pointer(uintptr(v1))))
	}
	return 0
}
func nox_script_builtin_516790(this unsafe.Pointer) int32 {
	var (
		v1 int32
		v2 int32
		v4 unsafe.Pointer
	)
	v4 = this
	v1 = nox_script_pop()
	v2 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v1))))
	if v2 != 0 {
		nox_script_push(int32((*(*uint32)(unsafe.Pointer(uintptr(v2 + 12))) >> 8) & 1))
	} else {
		nox_script_push(int32(uintptr(v4)))
	}
	return 0
}
func nox_script_BecomePet_5167D0() int32 {
	var (
		v0 *byte
		v1 int32
		v2 int32
		v3 int32
	)
	v0 = (*byte)(unsafe.Pointer(nox_common_playerInfoFromNum_417090(31)))
	v1 = nox_script_pop()
	v2 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v1))))
	if v2 != 0 {
		v3 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v0))), unsafe.Sizeof(uint32(0))*514))))
		if v3 != 0 {
			nox_xxx_unitBecomePet_4E7B00(v3, v2)
		}
	}
	return 0
}
func nox_script_BecomeEnemy_516810() int32 {
	var (
		v0 *byte
		v1 int32
		v2 *uint32
		v3 int32
	)
	v0 = (*byte)(unsafe.Pointer(nox_common_playerInfoFromNum_417090(31)))
	v1 = nox_script_pop()
	v2 = (*uint32)(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v1)))
	if v2 != nil {
		v3 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v0))), unsafe.Sizeof(uint32(0))*514))))
		if v3 != 0 {
			nox_xxx_monsterRemoveMonitors_4E7B60(v3, v2)
		}
	}
	return 0
}
func nox_script_builtin_516850(this unsafe.Pointer) int32 {
	var (
		v1 int32
		v2 int32
		v4 unsafe.Pointer
	)
	v4 = this
	v1 = nox_script_pop()
	v2 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v1))))
	if v2 != 0 {
		nox_script_push(int32((*(*uint32)(unsafe.Pointer(uintptr(v2 + 12))) & math.MaxUint8) >> 7))
	} else {
		nox_script_push(int32(uintptr(v4)))
	}
	return 0
}
func nox_script_OblivionGive_516890() int32 {
	var (
		v0 *uint32
		v1 int32
		v2 int32
		v3 *uint32
		v4 *uint32
	)
	v0 = (*uint32)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer(&nox_common_playerInfoFromNum_417090(31).field_0), unsafe.Sizeof(uint32(0))*514))))))
	v1 = 0
	v2 = nox_script_pop()
	v3 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v0), unsafe.Sizeof(uint32(0))*126)))))
	if v3 != nil {
		for (*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*2))&0x1000000) == 0 || (*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*3))&0x7800000) == 0 {
			v3 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*124)))))
			if v3 == nil {
				goto LABEL_7
			}
		}
		v1 = int32((*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*4)) >> 8) & 1)
		nox_xxx_delayedDeleteObject_4E5CC0((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v3)))))))
	}
LABEL_7:
	v4 = (*uint32)(unsafe.Pointer(nox_xxx_playerRespawnItem_4EF750((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v0)))))), *(**byte)(memmap.PtrOff(0x587000, uint32(v2*4)+0x3C628)), nil, 1, 1)))
	if v1 != 0 {
		nox_xxx_playerTryEquip_4F2F70(v0, (*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v4)))))))
	}
	return 0
}
func nox_script_Frozen_516920() int32 {
	var (
		v0 int32
		v1 int32
		v2 int32
	)
	v0 = nox_script_pop()
	v1 = nox_script_pop()
	v2 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v1))))
	if v2 != 0 {
		if v0 == 1 {
			nox_xxx_unitFreeze_4E79C0((*nox_object_t)(unsafe.Pointer(uintptr(v2))), 1)
			return 0
		}
		nox_xxx_unitUnFreeze_4E7A60((*nox_object_t)(unsafe.Pointer(uintptr(v2))), 1)
	}
	return 0
}
func nox_script_builtin_516960() int32 {
	nox_xxx_wallSounds_2386840 = uint32(nox_script_pop())
	return 0
}
func nox_script_SetProperty_516970() int32 {
	var (
		v0     int32
		v1     int32
		v2     int32
		v3     int32
		v4     *uint32
		result int32
	)
	v0 = nox_script_pop()
	v1 = nox_script_pop()
	v2 = nox_script_pop()
	v3 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v2))))
	if v3 == 0 || (int32(*(*uint8)(unsafe.Pointer(uintptr(v3 + 8))))&2) == 0 {
		return 0
	}
	v4 = *(**uint32)(unsafe.Pointer(uintptr(v3 + 748)))
	switch v1 {
	case 3:
		*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*309)) = uint32(v0)
		result = 0
	case 4:
		*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*307)) = uint32(v0)
		result = 0
	case 5:
		*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*317)) = uint32(v0)
		result = 0
	case 6:
		*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*311)) = uint32(v0)
		result = 0
	case 7:
		*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*313)) = uint32(v0)
		result = 0
	case 8:
		*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*315)) = uint32(v0)
		result = 0
	case 9:
		*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*319)) = uint32(v0)
		result = 0
	case 10:
		*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*321)) = uint32(v0)
		result = 0
	case 11:
		*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*323)) = uint32(v0)
		result = 0
	case 13:
		*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*325)) = uint32(v0)
		return 0
	default:
		return 0
	}
	return result
}
func nox_script_DecayObject_516A50() int32 {
	var (
		v0 int32
		v1 int32
		v2 *uint32
	)
	v0 = nox_script_pop()
	v1 = nox_script_pop()
	v2 = (*uint32)(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v1)))
	if v2 != nil {
		nox_xxx_unitSetDecayTime_511660((*nox_object_t)(unsafe.Pointer(v2)), v0)
	}
	return 0
}
func nox_script_TrapSpells_516B40() int32 {
	var (
		v0 int32
		v1 int32
		v2 int32
		v3 int32
		v4 int32
		v5 int32
		v6 int32
		v7 int32
		v8 int32
		v9 int32
	)
	v0 = nox_script_pop()
	v1 = nox_script_pop()
	v2 = nox_script_pop()
	v3 = nox_script_pop()
	v4 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v3))))
	v5 = v4
	if v4 != 0 {
		if int32(*(*uint8)(unsafe.Pointer(uintptr(v4 + 8))))&2 != 0 {
			v6 = int32(*(*uint32)(unsafe.Pointer(uintptr(v4 + 12))))
			if v6&8192 != 0 {
				v7 = things.SpellIndex(nox_script_getString_512E40(v2))
				v8 = things.SpellIndex(nox_script_getString_512E40(v1))
				v9 = things.SpellIndex(nox_script_getString_512E40(v0))
				sub_516A80(v5, v7, v8, v9)
			}
		}
	}
	return 0
}
func nox_script_builtin_516BC0() int32 {
	var (
		v0 int32
		v1 int32
	)
	v0 = nox_script_pop()
	v1 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v0))))
	if v1 != 0 {
		nox_xxx_netScriptMessageKill_4D9760(v1)
	}
	return 0
}
func nox_script_SetShopkeeperGreet_516BE0() int32 {
	var (
		v0 int32
		v1 int32
		v2 int32
	)
	v0 = nox_script_pop()
	v1 = nox_script_pop()
	v2 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v1))))
	if v2 != 0 {
		sub_548FE0(v2, nox_script_getString_512E40(v0))
	}
	return 0
}
func nox_script_builtin_516C10() int32 {
	sub_44E040()
	return 0
}
func nox_script_StartGame_516C20() int32 {
	nox_xxx_cliPlayMapIntro_44E0B0(1)
	return 0
}
func nox_script_TestMobStatus_516C30() int32 {
	var (
		v0 int32
		v1 int32
		v2 int32
	)
	v0 = nox_script_pop()
	v1 = 0
	v2 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v0))))
	if v2 != 0 && int32(*(*uint8)(unsafe.Pointer(uintptr(v2 + 8))))&2 != 0 {
		v1 = int32((*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 748))) + 1440))) >> 7) & 1)
	}
	nox_script_push(v1)
	return 0
}
func nox_script_builtin_516C70() int32 {
	var (
		v0 int32
		v1 int32
	)
	v0 = nox_script_pop()
	v1 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v0))))
	if v1 != 0 {
		nox_xxx_zombieSetStayDead_516C90(v1)
	}
	return 0
}
func nox_script_builtin_516CB0() int32 {
	var (
		v0 int32
		v1 *uint8
	)
	v0 = nox_script_pop()
	v1 = (*uint8)(unsafe.Pointer(uintptr(nox_server_scriptGetGroup_57C0A0(v0))))
	nox_server_scriptExecuteFnForEachGroupObj_502670(v1, 0, func(arg1 int32, arg2 int32) {
		nox_xxx_zombieSetStayDead_516C90(arg1)
	}, 0)
	return 0
}
func nox_script_builtin_516CE0() int32 {
	var (
		v0 int32
		v1 int32
	)
	v0 = nox_script_pop()
	v1 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v0))))
	if v1 != 0 {
		sub_516D00(v1)
	}
	return 0
}
func nox_script_builtin_516D40() int32 {
	var (
		v0 int32
		v1 *uint8
	)
	v0 = nox_script_pop()
	v1 = (*uint8)(unsafe.Pointer(uintptr(nox_server_scriptGetGroup_57C0A0(v0))))
	nox_server_scriptExecuteFnForEachGroupObj_502670(v1, 0, func(arg1 int32, arg2 int32) {
		sub_516D00(arg1)
	}, 0)
	return 0
}
func nox_script_ObjIsGameball_516D70() int32 {
	var (
		v0 int32
		v1 int32
		v2 int32
	)
	v0 = nox_script_pop()
	v1 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v0))))
	if v1 != 0 {
		v2 = int32(*memmap.PtrUint32(6112660, 2386908))
		if *memmap.PtrUint32(6112660, 2386908) == 0 {
			v2 = nox_xxx_getNameId_4E3AA0(CString("GameBall"))
			*memmap.PtrUint32(6112660, 2386908) = uint32(v2)
		}
		nox_script_push(bool2int(int32(*(*uint16)(unsafe.Pointer(uintptr(v1 + 4)))) == v2))
	}
	return 0
}
func nox_script_ObjIsCrown_516DC0() int32 {
	var (
		v0 int32
		v1 int32
		v2 int32
	)
	v0 = nox_script_pop()
	v1 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v0))))
	if v1 != 0 {
		v2 = int32(*memmap.PtrUint32(6112660, 2386912))
		if *memmap.PtrUint32(6112660, 2386912) == 0 {
			v2 = nox_xxx_getNameId_4E3AA0(CString("Crown"))
			*memmap.PtrUint32(6112660, 2386912) = uint32(v2)
		}
		nox_script_push(bool2int(int32(*(*uint16)(unsafe.Pointer(uintptr(v1 + 4)))) == v2))
	}
	return 0
}
func nox_script_EndGame_516E10() int32 {
	var v0 int32
	v0 = nox_script_pop()
	sub_578C90(v0)
	return 0
}
func nox_script_PlayerScore_516E30() int32 {
	var (
		v0 int32
		v1 int32
		v2 int32
		v3 int32
		v4 *byte
	)
	v0 = nox_script_pop()
	v1 = nox_script_pop()
	v2 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v1))))
	v3 = v2
	if v2 != 0 && int32(*(*uint8)(unsafe.Pointer(uintptr(v2 + 8))))&4 != 0 {
		if v0 <= 0 {
			nox_xxx_playerSubLessons_4D8EC0(v2, -v0)
		} else {
			nox_xxx_changeScore_4D8E90(v2, v0)
		}
		v4 = (*byte)(unsafe.Pointer(nox_xxx_clientGetTeamColor_418AB0(int32(*(*uint8)(unsafe.Pointer(uintptr(v3 + 52)))))))
		if v4 != nil {
			nox_xxx_netChangeTeamID_419090(int32(uintptr(unsafe.Pointer(v4))), int32(uint32(v0)+*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v4))), unsafe.Sizeof(uint32(0))*13)))))
		}
		nox_xxx_netReportLesson_4D8EF0((*nox_object_t)(unsafe.Pointer(uintptr(v3))))
	}
	return 0
}
func nox_script_builtin_516EA0() int32 {
	var (
		v0 int32
		v1 int32
		v2 int32
	)
	v0 = nox_script_pop()
	v1 = 0
	v2 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v0))))
	if v2 != 0 && int32(*(*uint8)(unsafe.Pointer(uintptr(v2 + 8))))&4 != 0 {
		v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 748))) + 276))) + 2136))))
	}
	nox_script_push(v1)
	return 0
}
func nox_script_Fn5E_513F70() int32 {
	var (
		v0 int32
		v1 *libc.WChar
		v2 int32
	)
	v0 = nox_script_pop()
	v1 = strMan.GetStringInFile(nox_script_getString_512E40(v0), "C:\\NoxPost\\src\\Server\\System\\CScrFunc.c")
	v2 = sub_512E80(int32(uintptr(unsafe.Pointer(v1))))
	nox_script_push(v2)
	return 0
}
func nox_script_builtin_514A80() int32 {
	var (
		v0 int16
		v1 int32
		v2 int32
		v3 int32
		v4 *byte
		v5 *libc.WChar
		v7 libc.WChar
		v8 int32
	)
	v0 = int16(nox_script_pop())
	v1 = nox_script_pop()
	v2 = nox_script_pop()
	v3 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v2))))
	if v3 != 0 {
		v4 = nox_script_getString_512E40(v1)
		v7 = libc.WChar(int16(int32(v0) * int32(uint16(nox_gameFPS))))
		v8 = 0
		v5 = nox_strman_loadString_40F1D0(v4, (**byte)(unsafe.Pointer(&v8)), CString("C:\\NoxPost\\src\\Server\\System\\CScrFunc.c"), 3629)
		nox_xxx_netSendChat_528AC0(v3, v5, v7)
		if noxflags.HasGame(noxflags.GameModeCoop) {
			nox_xxx_playDialogFile_44D900(v8, 100)
		}
	}
	return 0
}
func nox_script_builtin_514B10() int32 {
	var (
		v0 libc.WChar
		v1 int32
		v2 int32
		v3 int32
		v4 *byte
		v5 *libc.WChar
		v7 int32
	)
	v0 = libc.WChar(int16(nox_script_pop()))
	v1 = nox_script_pop()
	v2 = nox_script_pop()
	v3 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v2))))
	if v3 != 0 {
		v4 = nox_script_getString_512E40(v1)
		v7 = 0
		v5 = nox_strman_loadString_40F1D0(v4, (**byte)(unsafe.Pointer(&v7)), CString("C:\\NoxPost\\src\\Server\\System\\CScrFunc.c"), 3670)
		nox_xxx_netSendChat_528AC0(v3, v5, v0)
		if v7 != 0 {
			if noxflags.HasGame(noxflags.GameModeCoop) {
				nox_xxx_playDialogFile_44D900(v7, 100)
			}
		}
	}
	return 0
}
func nox_script_printToCaller_512B10() int32 {
	var (
		v0 int32
		v1 *libc.WChar
	)
	v0 = nox_script_pop()
	if nox_script_get_caller() != nil && int32(*(*uint8)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(nox_script_get_caller())), 8)))))&4 != 0 {
		v1 = strMan.GetStringInFile(nox_script_getString_512E40(v0), "C:\\NoxPost\\src\\Server\\System\\CScrFunc.c")
		nox_xxx_netSendLineMessage_4D9EB0(int32(uintptr(nox_script_get_caller())), v1)
	}
	return 0
}
func nox_script_printToAll_512B60() int32 {
	var (
		v0 int32
		v1 *libc.WChar
	)
	v0 = nox_script_pop()
	v1 = strMan.GetStringInFile(nox_script_getString_512E40(v0), "C:\\NoxPost\\src\\Server\\System\\CScrFunc.c")
	nox_xxx_printToAll_4D9FD0(0, v1)
	return 0
}
func nox_script_sayChat_512B90() int32 {
	var (
		v0 int32
		v1 int32
		v2 int32
		v3 *byte
		v4 *libc.WChar
		v6 int32
	)
	v0 = nox_script_pop()
	v1 = nox_script_pop()
	v2 = int32(uintptr(unsafe.Pointer(nox_server_scriptValToObjectPtr_511B60(v1))))
	if v2 != 0 {
		v3 = nox_script_getString_512E40(v0)
		v6 = 0
		v4 = nox_strman_loadString_40F1D0(v3, (**byte)(unsafe.Pointer(&v6)), CString("C:\\NoxPost\\src\\Server\\System\\CScrFunc.c"), 1342)
		nox_xxx_netSendChat_528AC0(v2, v4, 0)
		if noxflags.HasGame(noxflags.GameModeCoop) {
			nox_xxx_playDialogFile_44D900(v6, 100)
		}
	}
	return 0
}

var nox_script_builtin [211]unsafe.Pointer = [211]unsafe.Pointer{unsafe.Pointer(cgoFuncAddr(nox_script_getWall_511EB0)), unsafe.Pointer(cgoFuncAddr(nox_script_openSecretWall_511F50)), unsafe.Pointer(cgoFuncAddr(nox_script_openWallGroup_512010)), unsafe.Pointer(cgoFuncAddr(nox_script_closeWall_512040)), unsafe.Pointer(cgoFuncAddr(nox_script_closeWallGroup_512100)), unsafe.Pointer(cgoFuncAddr(nox_script_toggleWall_512130)), unsafe.Pointer(cgoFuncAddr(nox_script_toggleWallGroup_512260)), unsafe.Pointer(cgoFuncAddr(nox_script_wallBreak_512290)), unsafe.Pointer(cgoFuncAddr(nox_script_wallGroupBreak_5122F0)), unsafe.Pointer(cgoFuncAddr(nox_script_secondTimer_512320)), unsafe.Pointer(cgoFuncAddr(nox_script_frameTimer_512350)), unsafe.Pointer(cgoFuncAddr(nox_script_moverOrMonsterGo_512370)), unsafe.Pointer(cgoFuncAddr(nox_script_groupGoTo_512500)), unsafe.Pointer(cgoFuncAddr(nox_script_lookAtDirection_512560)), unsafe.Pointer(cgoFuncAddr(nox_script_groupLookAtDirection_512610)), unsafe.Pointer(cgoFuncAddr(nox_script_objectOn_512670)), unsafe.Pointer(cgoFuncAddr(nox_script_objGroupOn_512690)), unsafe.Pointer(cgoFuncAddr(nox_script_waypointOn_5126D0)), unsafe.Pointer(cgoFuncAddr(nox_script_waypointGroupOn_5126F0)), unsafe.Pointer(cgoFuncAddr(nox_script_objectOff_512730)), unsafe.Pointer(cgoFuncAddr(nox_script_objGroupOff_512750)), unsafe.Pointer(cgoFuncAddr(nox_script_waypointOff_512790)), unsafe.Pointer(cgoFuncAddr(nox_script_waypointGroupOff_5127B0)), unsafe.Pointer(cgoFuncAddr(nox_script_toggleObject_5127F0)), unsafe.Pointer(cgoFuncAddr(nox_script_toggleObjectGroup_512810)), unsafe.Pointer(cgoFuncAddr(nox_script_toggleWaypoint_512850)), unsafe.Pointer(cgoFuncAddr(nox_script_toggleWaypointGroup_512870)), unsafe.Pointer(cgoFuncAddr(nox_script_deleteObject_5128B0)), unsafe.Pointer(cgoFuncAddr(nox_script_deleteObjectGroup_5128D0)), unsafe.Pointer(cgoFuncAddr(nox_script_followNearestWaypoint_512910)), unsafe.Pointer(cgoFuncAddr(nox_script_groupRoam_512990)), unsafe.Pointer(cgoFuncAddr(nox_script_getObject2_5129C0)), unsafe.Pointer(cgoFuncAddr(nox_script_getObject3_5129E0)), unsafe.Pointer(cgoFuncAddr(nox_script_gotoHome_512A00)), unsafe.Pointer(cgoFuncAddr(nox_script_audioEven_512AC0)), unsafe.Pointer(cgoFuncAddr(nox_script_printToCaller_512B10)), unsafe.Pointer(cgoFuncAddr(nox_script_printToAll_512B60)), unsafe.Pointer(cgoFuncAddr(nox_script_sayChat_512B90)), unsafe.Pointer(cgoFuncAddr(nox_script_returnOne_512C10)), unsafe.Pointer(cgoFuncAddr(nox_script_unlockDoor_512C20)), unsafe.Pointer(cgoFuncAddr(nox_script_lockDoor_512C60)), unsafe.Pointer(cgoFuncAddr(nox_script_isOn_512CA0)), unsafe.Pointer(cgoFuncAddr(nox_script_wpIsEnabled_512CE0)), unsafe.Pointer(cgoFuncAddr(nox_script_doorIsLocked_512D20)), unsafe.Pointer(cgoFuncAddr(nox_script_randomFloat_512D70)), unsafe.Pointer(cgoFuncAddr(nox_script_randomInt_512DB0)), unsafe.Pointer(cgoFuncAddr(nox_script_timerSecSpecial_512DE0)), unsafe.Pointer(cgoFuncAddr(nox_script_specialTimer_512E10)), unsafe.Pointer(cgoFuncAddr(nox_script_intToString_512EA0)), unsafe.Pointer(cgoFuncAddr(nox_script_floatToString_512ED0)), unsafe.Pointer(cgoFuncAddr(nox_script_create_512F10)), unsafe.Pointer(cgoFuncAddr(nox_script_damage_512F80)), unsafe.Pointer(cgoFuncAddr(nox_script_groupDamage_513010)), unsafe.Pointer(cgoFuncAddr(nox_script_builtin_513070)), unsafe.Pointer(cgoFuncAddr(nox_script_builtin_513160)), unsafe.Pointer(cgoFuncAddr(nox_script_awardSpell_5131C0)), unsafe.Pointer(cgoFuncAddr(nox_script_awardSpellGroup_513230)), unsafe.Pointer(cgoFuncAddr(nox_script_enchant_5132E0)), unsafe.Pointer(cgoFuncAddr(nox_script_groupEnchant_5133B0)), unsafe.Pointer(cgoFuncAddr(nox_script_getHost_513460)), unsafe.Pointer(cgoFuncAddr(nox_script_objectGet_513490)), unsafe.Pointer(cgoFuncAddr(nox_script_getObjectX_513530)), unsafe.Pointer(cgoFuncAddr(nox_script_getWaypointX_513570)), unsafe.Pointer(cgoFuncAddr(nox_script_getObjectY_5135B0)), unsafe.Pointer(cgoFuncAddr(nox_script_getWaypointY_5135F0)), unsafe.Pointer(cgoFuncAddr(nox_script_unitHeight_513630)), unsafe.Pointer(cgoFuncAddr(nox_script_getUnitLook_513670)), unsafe.Pointer(cgoFuncAddr(nox_script_moveObject_5136A0)), unsafe.Pointer(cgoFuncAddr(nox_script_moveWaypoint_513700)), unsafe.Pointer(cgoFuncAddr(nox_script_raise_513750)), unsafe.Pointer(cgoFuncAddr(nox_script_faceAngle_513780)), unsafe.Pointer(cgoFuncAddr(nox_script_pushObject_5137D0)), unsafe.Pointer(cgoFuncAddr(nox_script_pushObjectTo_513820)), unsafe.Pointer(cgoFuncAddr(nox_script_getFirstInvItem_5138B0)), unsafe.Pointer(cgoFuncAddr(nox_script_getNextInvItem_5138E0)), unsafe.Pointer(cgoFuncAddr(nox_script_hasItem_513910)), unsafe.Pointer(cgoFuncAddr(nox_script_getInvHolder_513960)), unsafe.Pointer(cgoFuncAddr(nox_script_pickup_5139A0)), unsafe.Pointer(cgoFuncAddr(nox_script_drop_513C10)), unsafe.Pointer(cgoFuncAddr(nox_script_getObjectType_516210)), unsafe.Pointer(cgoFuncAddr(nox_script_builtin_513C60)), unsafe.Pointer(cgoFuncAddr(nox_script_TestBuffs_513C70)), unsafe.Pointer(cgoFuncAddr(nox_script_cancelBuff_513D00)), unsafe.Pointer(cgoFuncAddr(nox_script_getCurrentHP_513D70)), unsafe.Pointer(cgoFuncAddr(nox_script_getMaxHP_513DB0)), unsafe.Pointer(cgoFuncAddr(nox_script_restoreHP_513DF0)), unsafe.Pointer(cgoFuncAddr(nox_script_getDistance_513E20)), unsafe.Pointer(cgoFuncAddr(nox_script_canInteract_513E80)), unsafe.Pointer(cgoFuncAddr(nox_script_fn58_513F10)), unsafe.Pointer(cgoFuncAddr(nox_script_fn59_513F20)), unsafe.Pointer(cgoFuncAddr(nox_script_fn5A_513F30)), unsafe.Pointer(cgoFuncAddr(nox_script_fn5B_513F40)), unsafe.Pointer(cgoFuncAddr(nox_script_Fn5C_513F50)), unsafe.Pointer(cgoFuncAddr(nox_script_Fn5D_513F60)), unsafe.Pointer(cgoFuncAddr(nox_script_Fn5E_513F70)), unsafe.Pointer(cgoFuncAddr(nox_script_GetHostInfo_513FA0)), unsafe.Pointer(cgoFuncAddr(nox_script_FaceObject_514050)), unsafe.Pointer(cgoFuncAddr(nox_script_Walk_5140B0)), unsafe.Pointer(cgoFuncAddr(nox_script_GroupWalk_514170)), unsafe.Pointer(cgoFuncAddr(nox_script_CancelTimer_5141F0)), unsafe.Pointer(cgoFuncAddr(nox_script_Effect_514210)), unsafe.Pointer(cgoFuncAddr(nox_script_SetOwner_514490)), unsafe.Pointer(cgoFuncAddr(nox_script_GroupSetOwner_5144C0)), unsafe.Pointer(cgoFuncAddr(nox_script_SetOwnerGroup_514510)), unsafe.Pointer(cgoFuncAddr(nox_script_builtin_514570)), unsafe.Pointer(cgoFuncAddr(nox_script_builtin_5145F0)), unsafe.Pointer(cgoFuncAddr(nox_script_builtin_514630)), unsafe.Pointer(cgoFuncAddr(nox_script_builtin_5146B0)), unsafe.Pointer(cgoFuncAddr(nox_script_builtin_514730)), unsafe.Pointer(cgoFuncAddr(nox_script_ClearOwner_5147E0)), unsafe.Pointer(cgoFuncAddr(nox_script_Waypoint_514800)), unsafe.Pointer(cgoFuncAddr(nox_script_GetWaypointGroup_5148A0)), unsafe.Pointer(cgoFuncAddr(nox_script_GetObjectGroup_514940)), unsafe.Pointer(cgoFuncAddr(nox_script_GetWallGroup_5149E0)), unsafe.Pointer(cgoFuncAddr(nox_script_builtin_514A80)), unsafe.Pointer(cgoFuncAddr(nox_script_builtin_514B10)), unsafe.Pointer(cgoFuncAddr(nox_script_Pop2_74_514BA0)), unsafe.Pointer(cgoFuncAddr(nox_script_RemoveChat_514BB0)), unsafe.Pointer(cgoFuncAddr(nox_script_NoChatAll_514BD0)), unsafe.Pointer(cgoFuncAddr(nox_script_JournalQuest_514BE0)), unsafe.Pointer(cgoFuncAddr(nox_script_JournalQuestBool_514C10)), unsafe.Pointer(cgoFuncAddr(nox_script_JournalGetQuest_514C40)), unsafe.Pointer(cgoFuncAddr(nox_script_JournalGetQuestBool_514C60)), unsafe.Pointer(cgoFuncAddr(nox_script_RestoreQuestStatus_514C90)), unsafe.Pointer(cgoFuncAddr(nox_script_ObjIsTrigger_514CB0)), unsafe.Pointer(cgoFuncAddr(nox_script_builtin_514CF0)), unsafe.Pointer(cgoFuncAddr(nox_script_SetDialog_514D90)), unsafe.Pointer(cgoFuncAddr(nox_script_CancelDialog_514DF0)), unsafe.Pointer(cgoFuncAddr(nox_script_DialogPortrait_514E30)), unsafe.Pointer(cgoFuncAddr(nox_script_TellStory_514E90)), unsafe.Pointer(cgoFuncAddr(nox_script_ForceDialog_514ED0)), unsafe.Pointer(cgoFuncAddr(nox_script_builtin_514F10)), unsafe.Pointer(cgoFuncAddr(nox_script_builtin_514FC0)), unsafe.Pointer(cgoFuncAddr(nox_script_builtin_515060)), unsafe.Pointer(cgoFuncAddr(nox_script_cast_515130)), unsafe.Pointer(cgoFuncAddr(nox_script_UnBlind_515200)), unsafe.Pointer(cgoFuncAddr(nox_script_Blind_515220)), unsafe.Pointer(cgoFuncAddr(nox_script_WideScreen_515240)), unsafe.Pointer(cgoFuncAddr(nox_script_GetElevatorStat_5154A0)), unsafe.Pointer(cgoFuncAddr(nox_script_SetGuardPoint_515600)), unsafe.Pointer(cgoFuncAddr(nox_script_builtin_515700)), unsafe.Pointer(cgoFuncAddr(nox_script_MonsterHunt_515780)), unsafe.Pointer(cgoFuncAddr(nox_script_builtin_5157D0)), unsafe.Pointer(cgoFuncAddr(nox_script_Idle_515800)), unsafe.Pointer(cgoFuncAddr(nox_script_GroupIdle_515850)), unsafe.Pointer(cgoFuncAddr(nox_script_Follow_515880)), unsafe.Pointer(cgoFuncAddr(nox_script_builtin_515910)), unsafe.Pointer(cgoFuncAddr(nox_script_AgresssionLevel_515950)), unsafe.Pointer(cgoFuncAddr(nox_script_builtin_5159B0)), unsafe.Pointer(cgoFuncAddr(nox_script_HitLocation_5159E0)), unsafe.Pointer(cgoFuncAddr(nox_script_builtin_515AE0)), unsafe.Pointer(cgoFuncAddr(nox_script_ShootLocation_515B30)), unsafe.Pointer(cgoFuncAddr(nox_script_builtin_515BF0)), unsafe.Pointer(cgoFuncAddr(nox_script_MonsterWayFlag_515C40)), unsafe.Pointer(cgoFuncAddr(nox_script_builtin_515CB0)), unsafe.Pointer(cgoFuncAddr(nox_script_Attack_515CF0)), unsafe.Pointer(cgoFuncAddr(nox_script_builtin_515DB0)), unsafe.Pointer(cgoFuncAddr(nox_script_builtin_5154E0)), unsafe.Pointer(cgoFuncAddr(nox_script_DeleteJournalEntry_515550)), unsafe.Pointer(cgoFuncAddr(nox_script_builtin_5155A0)), unsafe.Pointer(cgoFuncAddr(nox_script_SetIdleLevel_515DF0)), unsafe.Pointer(cgoFuncAddr(nox_script_GroupSetIdleHP_515E50)), unsafe.Pointer(cgoFuncAddr(nox_script_SetResumeLevel_515E80)), unsafe.Pointer(cgoFuncAddr(nox_script_builtin_515EE0)), unsafe.Pointer(cgoFuncAddr(nox_script_RunAway_515F10)), unsafe.Pointer(cgoFuncAddr(nox_script_monsterGroupFleeFrom_516000)), unsafe.Pointer(cgoFuncAddr(nox_script_MonsterWait_516060)), unsafe.Pointer(cgoFuncAddr(nox_script_builtin_5160F0)), unsafe.Pointer(cgoFuncAddr(nox_script_builtin_5161C0)), unsafe.Pointer(cgoFuncAddr(nox_script_builtin_516120)), unsafe.Pointer(cgoFuncAddr(nox_script_builtin_516160)), unsafe.Pointer(cgoFuncAddr(nox_script_DialogResult_5163C0)), unsafe.Pointer(cgoFuncAddr(nox_script_GiveExp_516190)), unsafe.Pointer(cgoFuncAddr(nox_script_builtin_5162D0)), unsafe.Pointer(cgoFuncAddr(nox_script_ForceAutosave_516400)), unsafe.Pointer(cgoFuncAddr(nox_script_Music_516430)), unsafe.Pointer(cgoFuncAddr(nox_script_builtin_516600)), unsafe.Pointer(cgoFuncAddr(nox_script_builtin_5166A0)), unsafe.Pointer(cgoFuncAddr(nox_script_GetTrigger_514D30)), unsafe.Pointer(cgoFuncAddr(nox_script_GetCaller_514D60)), unsafe.Pointer(cgoFuncAddr(nox_script_builtin_516720)), unsafe.Pointer(cgoFuncAddr(nox_script_builtin_516760)), unsafe.Pointer(cgoFuncAddr(nox_script_BecomePet_5167D0)), unsafe.Pointer(cgoFuncAddr(nox_script_BecomeEnemy_516810)), unsafe.Pointer(cgoFuncAddr(nox_script_builtin_516790)), unsafe.Pointer(cgoFuncAddr(nox_script_builtin_516850)), unsafe.Pointer(cgoFuncAddr(nox_script_OblivionGive_516890)), unsafe.Pointer(cgoFuncAddr(nox_script_builtin_516680)), unsafe.Pointer(cgoFuncAddr(nox_script_Frozen_516920)), unsafe.Pointer(cgoFuncAddr(nox_script_builtin_516960)), unsafe.Pointer(cgoFuncAddr(nox_script_SetProperty_516970)), unsafe.Pointer(cgoFuncAddr(nox_script_DecayObject_516A50)), unsafe.Pointer(cgoFuncAddr(nox_script_TrapSpells_516B40)), unsafe.Pointer(cgoFuncAddr(nox_script_PlayerIsTrading_5166E0)), unsafe.Pointer(cgoFuncAddr(nox_script_builtin_516BC0)), unsafe.Pointer(cgoFuncAddr(nox_script_SetShopkeeperGreet_516BE0)), unsafe.Pointer(cgoFuncAddr(nox_script_builtin_516C10)), unsafe.Pointer(cgoFuncAddr(nox_script_TestMobStatus_516C30)), unsafe.Pointer(cgoFuncAddr(nox_script_builtin_516C70)), unsafe.Pointer(cgoFuncAddr(nox_script_builtin_516CB0)), unsafe.Pointer(cgoFuncAddr(nox_script_builtin_516CE0)), unsafe.Pointer(cgoFuncAddr(nox_script_builtin_516D40)), unsafe.Pointer(cgoFuncAddr(nox_script_builtin_5164A0)), unsafe.Pointer(cgoFuncAddr(nox_script_builtin_5164E0)), unsafe.Pointer(cgoFuncAddr(nox_script_ClearMusic_516520)), unsafe.Pointer(cgoFuncAddr(nox_script_ObjIsGameball_516D70)), unsafe.Pointer(cgoFuncAddr(nox_script_ObjIsCrown_516DC0)), unsafe.Pointer(cgoFuncAddr(nox_script_EndGame_516E10)), unsafe.Pointer(cgoFuncAddr(nox_script_StartGame_516C20)), unsafe.Pointer(cgoFuncAddr(nox_script_PlayerScore_516E30)), unsafe.Pointer(cgoFuncAddr(nox_script_builtin_516EA0))}
var nox_script_builtin_cnt int32 = int32(unsafe.Sizeof([211]unsafe.Pointer{}) / unsafe.Sizeof(unsafe.Pointer(nil)))
var nox_script_builtin_xxx [7]unsafe.Pointer = [7]unsafe.Pointer{unsafe.Pointer(cgoFuncAddr(nox_script_objectGet_513490)), unsafe.Pointer(cgoFuncAddr(nox_script_Waypoint_514800)), unsafe.Pointer(cgoFuncAddr(nox_script_GetObjectGroup_514940)), unsafe.Pointer(cgoFuncAddr(nox_script_GetWaypointGroup_5148A0)), unsafe.Pointer(cgoFuncAddr(nox_script_GetWallGroup_5149E0)), nil, nil}
var nox_script_xxxCount_246800 int32 = 7

func nox_script_builtinNeedsField36_508C30(fi int32) int32 {
	for i := int32(0); i < nox_script_xxxCount_246800; i++ {
		if nox_script_builtin[fi] == nox_script_builtin_xxx[i] {
			return 1
		}
	}
	return 0
}

var nox_script_builtin_yyy [7]unsafe.Pointer = [7]unsafe.Pointer{unsafe.Pointer(cgoFuncAddr(nox_script_getWall_511EB0)), unsafe.Pointer(cgoFuncAddr(nox_script_moveObject_5136A0)), unsafe.Pointer(cgoFuncAddr(nox_script_moveWaypoint_513700)), unsafe.Pointer(cgoFuncAddr(nox_script_pushObjectTo_513820)), unsafe.Pointer(cgoFuncAddr(nox_script_Walk_5140B0)), unsafe.Pointer(cgoFuncAddr(nox_script_GroupWalk_514170)), unsafe.Pointer(cgoFuncAddr(nox_script_Effect_514210))}
var nox_script_yyyCount_246768 int32 = 5

func nox_script_builtinNeedsFields4044_508C30(fi int32) int32 {
	for i := int32(0); i < nox_script_yyyCount_246768; i++ {
		if nox_script_builtin[fi] == nox_script_builtin_yyy[i] {
			return 1
		}
	}
	return 0
}
func nox_script_callBuiltin(fi int32) int32 {
	var res int32 = 0
	if nox_script_panic_compiler_call(fi, &res) {
		return res
	}
	return (cgoAsFunc(*(*uint32)(unsafe.Pointer(&nox_script_builtin[fi])), (*func() int32)(nil)))()
}
func nox_script_callBuiltin_508B70(i int32, fi int32) int32 {
	if fi < 0 || fi >= nox_script_builtin_cnt {
		if !nox_script_panic_compiler_check(fi) {
			stdio.Printf("noxscript: invalid builtin index: %d (%x)\n", fi, fi)
			return 0
		}
		return 0
	}
	if nox_script_getField36(i) == nil {
		return nox_script_callBuiltin(fi)
	}
	if nox_script_builtinNeedsField36_508C30(fi) != 0 {
		libc.StrCpy(&nox_script_builtin_buf_xxx[0], nox_script_getField36(i))
	}
	if nox_script_builtinNeedsFields4044_508C30(fi) != 0 {
		dword_5d4594_3821636 = nox_script_getField40(i)
		dword_5d4594_3821640 = nox_script_getField44(i)
	}
	var res int32 = nox_script_callBuiltin(fi)
	nox_script_resetBuiltin()
	return res
}
func nox_script_shouldReadMoreXxx(fi int32) bool {
	var f unsafe.Pointer = nox_script_builtin[fi]
	return cgoFuncAddr(cgoAsFunc(f, (*func() int32)(nil)).(func() int32)) == cgoFuncAddr(nox_script_secondTimer_512320) || cgoFuncAddr(cgoAsFunc(f, (*func() int32)(nil)).(func() int32)) == cgoFuncAddr(nox_script_frameTimer_512350) || cgoFuncAddr(cgoAsFunc(f, (*func() int32)(nil)).(func() int32)) == cgoFuncAddr(nox_script_timerSecSpecial_512DE0) || cgoFuncAddr(cgoAsFunc(f, (*func() int32)(nil)).(func() int32)) == cgoFuncAddr(nox_script_specialTimer_512E10) || cgoFuncAddr(cgoAsFunc(f, (*func() int32)(nil)).(func() int32)) == cgoFuncAddr(nox_script_SetProperty_516970) || cgoFuncAddr(cgoAsFunc(f, (*func() int32)(nil)).(func() int32)) == cgoFuncAddr(nox_script_SetDialog_514D90)
}
func nox_script_shouldReadEvenMoreXxx(fi int32) bool {
	var f unsafe.Pointer = nox_script_builtin[fi]
	return cgoFuncAddr(cgoAsFunc(f, (*func() int32)(nil)).(func() int32)) == cgoFuncAddr(nox_script_SetDialog_514D90)
}
