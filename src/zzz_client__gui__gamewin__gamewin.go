package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"unsafe"
)

func nox_xxx_clientPickup_46C140(a1p *nox_drawable) {
	var (
		a1 int32 = int32(uintptr(unsafe.Pointer(a1p)))
		v2 int32
		v3 *libc.WChar
		v4 int32
	)
	if *memmap.PtrUint32(6112660, 1064928) == 0 {
		*memmap.PtrUint32(6112660, 1064928) = uint32(nox_xxx_getTTByNameSpriteMB_44CFC0("Gold"))
		*memmap.PtrUint32(6112660, 1064932) = uint32(nox_xxx_getTTByNameSpriteMB_44CFC0("QuestGoldPile"))
		*memmap.PtrUint32(6112660, 1064936) = uint32(nox_xxx_getTTByNameSpriteMB_44CFC0("QuestGoldChest"))
	}
	if a1 != 0 {
		v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 108))))
		if uint32(v2) == *memmap.PtrUint32(6112660, 1064928) || uint32(v2) == *memmap.PtrUint32(6112660, 1064932) || uint32(v2) == *memmap.PtrUint32(6112660, 1064936) || sub_467B00(v2, 1) != 0 {
			*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v4))), 0)) = 115
			*(*uint16)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(&v4))), 1)))) = uint16(nox_xxx_netGetUnitCodeCli_578B00(a1))
			nox_netlist_addToMsgListCli_40EBC0(31, 0, (*uint8)(unsafe.Pointer(&v4)), 3)
		} else {
			clientPlaySoundSpecial(925, 100)
			v3 = strMan.GetStringInFile("pickup.c:CarryingTooMuch", "C:\\NoxPost\\src\\Client\\Gui\\GameWin\\gamewin.c")
			nox_xxx_printCentered_445490(v3)
		}
	}
}
