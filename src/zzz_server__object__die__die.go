package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	"unsafe"
)

func nox_xxx_dieArmor_54E170_obj_die(a1 int32) {
	var (
		v1  int32
		v2  int32
		v3  *uint32
		v4  int16
		v5  int32
		v6  int16
		v7  *libc.WChar
		v8  int32
		v9  *libc.WChar
		v10 *float2
	)
	v1 = a1
	v2 = 0
	if strMan.Lang() == 0 || strMan.Lang() == 1 {
		v3 = nox_xxx_equipClothFindDefByTT_413270(int32(*(*uint16)(unsafe.Pointer(uintptr(a1 + 4)))))
		if v3 != nil {
			v4 = int16(*(*uint16)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*2)) + nox_wcslen((*libc.WChar)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*2))))))*2 - 2))))
			if int32(v4) == 83 || int32(v4) == 115 {
				v2 = 1
			}
		}
	}
	v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 492))))
	if v5 != 0 {
		v10 = (*float2)(unsafe.Pointer(uintptr(v5 + 56)))
	} else {
		v10 = (*float2)(unsafe.Pointer(uintptr(a1 + 56)))
	}
	v6 = int16(*(*uint16)(unsafe.Pointer(uintptr(v1 + 24))))
	if int32(v6)&16 != 0 {
		if v2 != 0 {
			v7 = strMan.GetStringInFile("ArmorDieMetalPlural", "C:\\NoxPost\\src\\Server\\Object\\die\\Die.c")
		} else {
			v7 = strMan.GetStringInFile("ArmorDieMetal", "C:\\NoxPost\\src\\Server\\Object\\die\\Die.c")
		}
		v8 = 806
	} else if int32(v6)&8 != 0 {
		if v2 != 0 {
			v7 = strMan.GetStringInFile("ArmorDieWoodPlural", "C:\\NoxPost\\src\\Server\\Object\\die\\Die.c")
		} else {
			v7 = strMan.GetStringInFile("ArmorDieWood", "C:\\NoxPost\\src\\Server\\Object\\die\\Die.c")
		}
		v8 = 812
	} else if int32(v6)&4 != 0 {
		if v2 != 0 {
			v7 = strMan.GetStringInFile("ArmorDieHidePlural", "C:\\NoxPost\\src\\Server\\Object\\die\\Die.c")
		} else {
			v7 = strMan.GetStringInFile("ArmorDieHide", "C:\\NoxPost\\src\\Server\\Object\\die\\Die.c")
		}
		v8 = 809
	} else if int32(v6)&2 != 0 {
		if v2 != 0 {
			v7 = strMan.GetStringInFile("ArmorDieClothPlural", "C:\\NoxPost\\src\\Server\\Object\\die\\Die.c")
		} else {
			v7 = strMan.GetStringInFile("ArmorDieCloth", "C:\\NoxPost\\src\\Server\\Object\\die\\Die.c")
		}
		v8 = 815
	} else {
		v8 = int32(uintptr(unsafe.Pointer(v10)))
		v7 = strMan.GetStringInFile("ArmorDieGeneric", "C:\\NoxPost\\src\\Server\\Object\\die\\Die.c")
	}
	v9 = nox_xxx_itemGetName_4E77E0_obj_util(v1)
	nox_xxx_netSendLineMessage_4D9EB0(v5, v7, v9)
	nox_xxx_audCreate_501A30(v8, v10, 0, 0)
	nox_xxx_delayedDeleteObject_4E5CC0((*nox_object_t)(unsafe.Pointer(uintptr(v1))))
}
func nox_xxx_dieWeapon_54E370_obj_die(a1 int32) {
	var (
		v1 int32
		v2 *float2
		v3 int16
		v4 *libc.WChar
		v5 *libc.WChar
		v6 *libc.WChar
		v7 *libc.WChar
		v8 *libc.WChar
		v9 *libc.WChar
	)
	v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 492))))
	v2 = (*float2)(unsafe.Pointer(uintptr(v1 + 56)))
	if v1 == 0 {
		v2 = (*float2)(unsafe.Pointer(uintptr(a1 + 56)))
	}
	v3 = int16(*(*uint16)(unsafe.Pointer(uintptr(a1 + 24))))
	if int32(v3)&16 != 0 {
		v7 = nox_xxx_itemGetName_4E77E0_obj_util(a1)
		v4 = strMan.GetStringInFile("WeaponDieMetal", "C:\\NoxPost\\src\\Server\\Object\\die\\Die.c")
		nox_xxx_netSendLineMessage_4D9EB0(v1, v4, v7)
		nox_xxx_audCreate_501A30(818, v2, 0, 0)
		nox_xxx_delayedDeleteObject_4E5CC0((*nox_object_t)(unsafe.Pointer(uintptr(a1))))
	} else {
		if int32(v3)&8 != 0 {
			v8 = nox_xxx_itemGetName_4E77E0_obj_util(a1)
			v5 = strMan.GetStringInFile("WeaponDieWood", "C:\\NoxPost\\src\\Server\\Object\\die\\Die.c")
			nox_xxx_netSendLineMessage_4D9EB0(v1, v5, v8)
			nox_xxx_audCreate_501A30(819, v2, 0, 0)
		} else {
			v9 = sub_415B60(a1)
			v6 = strMan.GetStringInFile("WeaponDieGeneric", "C:\\NoxPost\\src\\Server\\Object\\die\\Die.c")
			nox_xxx_netSendLineMessage_4D9EB0(v1, v6, v9)
		}
		nox_xxx_delayedDeleteObject_4E5CC0((*nox_object_t)(unsafe.Pointer(uintptr(a1))))
	}
}
