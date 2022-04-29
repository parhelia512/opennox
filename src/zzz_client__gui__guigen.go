package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"unsafe"
)

func nox_xxx_mapGenClientText_4A9D00(a1 *uint8) int16 {
	var (
		v1 *uint8
		v2 int32
		v3 int32
		v4 int32
		v5 *libc.WChar
		v6 int32
		v7 int32
	)
	v1 = a1
	v2 = int32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(a1), 1)))))
	if uint32(uint16(int16(v2))) == *memmap.PtrUint32(6112660, 1309668) {
		return int16(v2)
	}
	*memmap.PtrUint32(6112660, 1309668) = uint32(*(*uint16)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(a1), 1)))))
	clientPlaySoundSpecial(897, 50)
	sub_430B50(0, 0, nox_win_width-1, nox_win_height-1)
	nox_client_clearScreen_440900()
	v3 = nox_win_height/2 - 120
	v4 = nox_win_width/2 - 160
	nox_client_drawImageAt_47D2C0((*nox_video_bag_image_t)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, (*memmap.PtrUint32(6112660, 1309672)%4)*4+0x13FBCC)))), v4, v3)
	nox_client_drawImageAt_47D2C0((*nox_video_bag_image_t)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, *memmap.PtrUint32(6112660, 1309672)*4+0x13FB4C)))), v4, v3)
	switch *v1 {
	case 155:
		v5 = strMan.GetStringInFile("Generating", "C:\\NoxPost\\src\\client\\Gui\\guigen.c")
		goto LABEL_8
	case 156:
		v5 = strMan.GetStringInFile("Assembling", "C:\\NoxPost\\src\\client\\Gui\\guigen.c")
		goto LABEL_8
	case 157:
		v5 = strMan.GetStringInFile("Populating", "C:\\NoxPost\\src\\client\\Gui\\guigen.c")
	LABEL_8:
		*memmap.PtrUint32(6112660, 1309660) = uint32(uintptr(unsafe.Pointer(v5)))
	}
	nox_xxx_drawGetStringSize_43F840(nil, (*libc.WChar)(unsafe.Pointer(*(**uint16)(memmap.PtrOff(6112660, 1309660)))), (*int32)(unsafe.Pointer(&a1)), nil, 0)
	v6 = int32(uint32(nox_win_width) - uint32(uintptr(unsafe.Pointer(a1))))
	v7 = nox_win_height/2 - (nox_xxx_guiFontHeightMB_43F320(nil)*2 + 70)
	nox_xxx_drawSetTextColor_434390(int32(nox_color_white_2523948))
	nox_xxx_drawSetColor_4343E0(*memmap.PtrInt32(0x852978, 4))
	nox_draw_drawStringHL_43F730(nil, (*libc.WChar)(unsafe.Pointer(*(**int16)(memmap.PtrOff(6112660, 1309660)))), v6/2, v7)
	if func() uint32 {
		p := memmap.PtrUint32(6112660, 1309672)
		*p++
		return *p
	}() >= 32 {
		*memmap.PtrUint32(6112660, 1309672) = 0
	}
	nox_video_callCopyBackBuffer_4AD170()
	return int16(v2)
}
