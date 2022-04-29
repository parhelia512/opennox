package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"unsafe"
)

func nox_xxx_netDrawRays_49BDD0(a1 *uint8) int32 {
	var (
		result int32
		v2     *uint8
		v3     *uint8
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
		v14    int32
		v15    int32
		v16    int8
		v17    int8
		v18    int8
		v19    [4]uint16
		v20    int32
	)
	result = int32(*memmap.PtrUint32(6112660, 1304308))
	if *memmap.PtrInt32(6112660, 1304308) < 96 {
		if *memmap.PtrUint32(6112660, 1304316) == 0 {
			*memmap.PtrUint32(6112660, 1304316) = uint32(nox_xxx_getTTByNameSpriteMB_44CFC0("DynamicLightning"))
			*memmap.PtrUint32(6112660, 1304320) = uint32(nox_xxx_getTTByNameSpriteMB_44CFC0("DynamicChainLightning"))
			*memmap.PtrUint32(6112660, 1304324) = uint32(nox_xxx_getTTByNameSpriteMB_44CFC0("DynamicEnergyBolt"))
			*memmap.PtrUint32(6112660, 1304348) = uint32(nox_xxx_getTTByNameSpriteMB_44CFC0("GreenZap"))
			dword_5d4594_1304328 = uint32(nox_xxx_getTTByNameSpriteMB_44CFC0("OrbRay"))
			*memmap.PtrUint32(6112660, 1304332) = uint32(nox_xxx_getTTByNameSpriteMB_44CFC0("PlasmaRay"))
			*memmap.PtrUint32(6112660, 1304336) = uint32(nox_xxx_getTTByNameSpriteMB_44CFC0("DrainManaOrb"))
			*memmap.PtrUint32(6112660, 1304340) = uint32(nox_xxx_getTTByNameSpriteMB_44CFC0("HealOrb"))
			*memmap.PtrUint32(6112660, 1304344) = uint32(nox_xxx_getTTByNameSpriteMB_44CFC0("CharmOrb"))
		}
		v2 = a1
		v3 = (*uint8)(unsafe.Add(unsafe.Pointer(a1), 1))
		v4 = int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(a1), 1))))) + (int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(a1), 5)))))-int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(a1), 1))))))/2
		result = int32(*a1) - 125
		v5 = int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(a1), 3))))) + (int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(a1), 7)))))-int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(a1), 3))))))/2
		switch *a1 {
		case 125:
			v20 = int32(*memmap.PtrUint32(6112660, 1304332))
			goto LABEL_17
		case 140:
			v20 = int32(*memmap.PtrUint32(6112660, 1304316))
			goto LABEL_17
		case 141:
			v20 = int32(*memmap.PtrUint32(6112660, 1304324))
			goto LABEL_17
		case 142:
			v20 = int32(*memmap.PtrUint32(6112660, 1304320))
			goto LABEL_17
		case 143:
			v18 = int8(randomIntMinMax(6, 12))
			v20 = int32(dword_5d4594_1304328)
			if randomIntMinMax(0, 100) < 50 {
				v15 = randomIntMinMax(-20, 20)
				v10 = randomIntMinMax(-20, 20)
				sub_499490(*memmap.PtrInt32(6112660, 1304336), (*uint16)(unsafe.Pointer(v3)), v10, v15, v18, 0)
			}
			goto LABEL_17
		case 144:
			v16 = int8(randomIntMinMax(6, 12))
			v20 = int32(dword_5d4594_1304328)
			if randomIntMinMax(0, 100) < 50 {
				v12 = randomIntMinMax(-20, 20)
				v6 = randomIntMinMax(-20, 20)
				sub_499490(*memmap.PtrInt32(6112660, 1304344), (*uint16)(unsafe.Pointer(v3)), v6, v12, v16, 0)
			}
			v7 = int32(*(*uint32)(unsafe.Pointer(v3)))
			*(*uint32)(unsafe.Pointer(&v19[0])) = *(*uint32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v2), 5))))
			*(*uint32)(unsafe.Pointer(&v19[2])) = uint32(v7)
			if randomIntMinMax(0, 100) < 50 {
				v13 = randomIntMinMax(-20, 20)
				v8 = randomIntMinMax(-20, 20)
				sub_499490(*memmap.PtrInt32(6112660, 1304344), &v19[0], v8, v13, v16, 0)
			}
			goto LABEL_17
		case 145:
			v17 = int8(randomIntMinMax(6, 12))
			v20 = int32(dword_5d4594_1304328)
			if randomIntMinMax(0, 100) < 50 {
				v14 = randomIntMinMax(-20, 20)
				v9 = randomIntMinMax(-20, 20)
				sub_499490(*memmap.PtrInt32(6112660, 1304340), (*uint16)(unsafe.Pointer(v3)), v9, v14, v17, 0)
			}
		LABEL_17:
			result = int32(uintptr(unsafe.Pointer(nox_xxx_spriteLoadAdd_45A360_drawable(v20, v4, v5))))
			if result != 0 {
				*(*uint8)(unsafe.Pointer(uintptr(result + 432))) = 0
				*(*uint32)(unsafe.Pointer(uintptr(result + 437))) = *(*uint32)(unsafe.Pointer(v3))
				*(*uint32)(unsafe.Pointer(uintptr(result + 441))) = *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v3))), unsafe.Sizeof(uint32(0))*1)))
				v11 = int32(*memmap.PtrUint32(6112660, 1304308))
				*memmap.PtrUint32(6112660, *memmap.PtrUint32(6112660, 1304308)*4+1303540) = uint32(result)
				*memmap.PtrUint32(6112660, 1304308) = uint32(v11 + 1)
			}
		default:
			return result
		}
	}
	return result
}
