package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	"github.com/gotranspile/cxgo/runtime/stdio"
	"github.com/noxworld-dev/opennox/v1/common/alloc"
	noxflags "github.com/noxworld-dev/opennox/v1/common/flags"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"image"
	"math"
	"unicode"
	"unsafe"
)

var nox_win_unk1 *nox_window = nil
var dword_5d4594_1046512 *nox_window = nil
var nox_xxx_drawablePlayer_1046600 *nox_drawable = nil
var dword_587000_122856 uint32 = 1
var dword_5d4594_831088 uint32 = 0
var dword_5d4594_831092 uint32 = 0
var nox_player_netCode_85319C uint32 = 0

func nox_xxx_spriteDefByAlphabetAdd_44CD10(a1 *byte) *byte {
	var result *byte
	result = a1
	if a1 != nil {
		result = (*byte)(unsafe.Pointer(uintptr(nox_xxx_keyFirstLetterNumberCli_44CD30(a1))))
		if int32(uintptr(unsafe.Pointer(result))) >= 0 {
			*memmap.PtrUint32(6112660, uint32(uintptr(unsafe.Pointer(result)))*4+0xCAC98)++
		}
	}
	return result
}
func nox_xxx_keyFirstLetterNumberCli_44CD30(a1 *byte) int32 {
	var (
		result int32
		v2     int32
	)
	if a1 == nil {
		return -1
	}
	v2 = int32(unicode.ToUpper(rune(*a1)))
	if v2 < 65 || v2 > 90 {
		result = 26
	} else {
		result = v2 - 65
	}
	return result
}
func nox_get_thing_name(i int32) *byte {
	if i < 1 || i >= nox_things_count {
		return nil
	}
	return (*(**nox_thing)(unsafe.Add(unsafe.Pointer(nox_things_array), unsafe.Sizeof((*nox_thing)(nil))*uintptr(i)))).name
}
func nox_get_thing(i int32) *nox_thing {
	if i < 1 || i >= nox_things_count {
		return nil
	}
	return *(**nox_thing)(unsafe.Add(unsafe.Pointer(nox_things_array), unsafe.Sizeof((*nox_thing)(nil))*uintptr(i)))
}
func nox_get_thing_pretty_name(i int32) *libc.WChar {
	if i < 1 || i >= nox_things_count {
		return nil
	}
	return (*(**nox_thing)(unsafe.Add(unsafe.Pointer(nox_things_array), unsafe.Sizeof((*nox_thing)(nil))*uintptr(i)))).pretty_name
}
func nox_get_thing_desc(i int32) *libc.WChar {
	if i < 1 || i >= nox_things_count {
		return nil
	}
	return (*(**nox_thing)(unsafe.Add(unsafe.Pointer(nox_things_array), unsafe.Sizeof((*nox_thing)(nil))*uintptr(i)))).desc
}
func nox_get_thing_pretty_image(i int32) int32 {
	if i < 1 || i >= nox_things_count {
		return 0
	}
	return int32((*(**nox_thing)(unsafe.Add(unsafe.Pointer(nox_things_array), unsafe.Sizeof((*nox_thing)(nil))*uintptr(i)))).pretty_image)
}
func nox_xxx_getTTByNameSpriteMB_44CFC0(a1 *byte) int32 {
	return nox_xxx_getvalByName2Imp_44CFD0(a1)
}
func nox_xxx_getvalByName2Imp_44CFD0(a1 *byte) int32 {
	var (
		v1     int32
		v2     int32
		v3     unsafe.Pointer
		v4     *uint32
		result int32
	)
	if a1 != nil && (func() bool {
		v1 = nox_xxx_keyFirstLetterNumberCli_44CD30(a1)
		return v1 >= 0
	}()) && (func() bool {
		v2 = int32(*memmap.PtrUint32(6112660, uint32(v1*4)+0xCAC98))
		return v2 >= 0
	}()) && (func() unsafe.Pointer {
		v3 = *(*unsafe.Pointer)(memmap.PtrOff(6112660, uint32(v1*4)+0xCAB58))
		return v3
	}()) != nil && (func() *uint32 {
		v4 = (*uint32)(libc.Search(unsafe.Pointer(a1), v3, uint32(v2), 8, sub_44D020))
		return v4
	}()) != nil {
		result = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*1)))
	} else {
		result = 0
	}
	return result
}
func sub_44D020(a1 unsafe.Pointer, a2 unsafe.Pointer) int32 {
	return int32(libc.StrCaseCmp((*byte)(a1), **(***byte)(a2)))
}
func sub_44D040(i int32) int32 {
	var obj *nox_thing = nox_get_thing(i)
	if obj == nil {
		return 0
	}
	return int32((obj.pri_class >> 22) & 1)
}
func sub_44D060(a1 int32) int32 {
	var v1 int32
	v1 = int32(uintptr(unsafe.Pointer(nox_get_thing(a1))))
	return bool2int(v1 != 0 && *(*uint32)(unsafe.Pointer(uintptr(v1 + 32)))&0x400000 != 0 && int32(*(*uint8)(unsafe.Pointer(uintptr(v1 + 36))))&24 != 0)
}
func sub_44D090(a1 int32) int32 {
	var result int32
	result = int32(uintptr(unsafe.Pointer(nox_get_thing(a1))))
	if result != 0 {
		result = bool2int((*(*uint32)(unsafe.Pointer(uintptr(result + 32))) & 0x20400000) != 0)
	}
	return result
}
func nox_drawable_link_thing(a1 *nox_drawable, i int32) int32 {
	var (
		v2 unsafe.Pointer
		v3 *uint32
		v4 *uint32
	)
	_ = v4
	var v9 float32
	if i < 1 || i >= nox_things_count {
		return 0
	}
	v2 = unsafe.Pointer(a1)
	*a1 = nox_drawable{}
	a1.field_27 = uint32(i)
	v3 = &a1.field_34
	v4 = &a1.field_38
	var v5 int32 = int32(uintptr(unsafe.Pointer(&(*(**nox_thing)(unsafe.Add(unsafe.Pointer(nox_things_array), unsafe.Sizeof((*nox_thing)(nil))*uintptr(i)))).hwidth)))
	*((*uint8)(unsafe.Pointer(&a1.field_0))) = *(*uint8)(unsafe.Pointer(uintptr(v5)))
	*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a1.field_0))), 1))) = *(*uint8)(unsafe.Pointer(uintptr(v5 + 1)))
	*(*uint16)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(a1))), 104))), 2)))) = *(*uint16)(unsafe.Pointer(uintptr(v5 + 10)))
	a1.flags28 = *(*uint32)(unsafe.Pointer(uintptr(v5 + 20)))
	a1.flags29 = *(*uint32)(unsafe.Pointer(uintptr(v5 + 24)))
	a1.flags30 = *(*uint32)(unsafe.Pointer(uintptr(v5 + 28)))
	a1.flags70 = *(*uint32)(unsafe.Pointer(uintptr(v5 + 72)))
	*(*uint8)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(a1))), 298)))) = *(*uint8)(unsafe.Pointer(uintptr(v5 + 2)))
	a1.draw_func = cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(v5 + 76))), (*func(*uint32, *nox_drawable) int32)(nil)).(func(*uint32, *nox_drawable) int32)
	a1.field_76 = unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v5 + 80)))))
	a1.field_77 = *(*uint32)(unsafe.Pointer(uintptr(v5 + 84)))
	a1.field_116 = *(*uint32)(unsafe.Pointer(uintptr(v5 + 88)))
	a1.field_123 = *(*uint32)(unsafe.Pointer(uintptr(v5 + 92)))
	a1.field_34 = uint32(*(*uint8)(unsafe.Pointer(uintptr(v5 + 3))))
	a1.field_42 = *(*uint32)(unsafe.Pointer(uintptr(v5 + 4)))
	a1.field_38 = *(*uint32)(unsafe.Pointer(uintptr(v5 + 36)))
	a1.field_39 = *(*uint32)(unsafe.Pointer(uintptr(v5 + 40)))
	a1.field_40 = *(*uint32)(unsafe.Pointer(uintptr(v5 + 44)))
	a1.field_41_0 = *(*uint16)(unsafe.Pointer(uintptr(v5 + 12)))
	a1.field_41_1 = *(*uint16)(unsafe.Pointer(uintptr(v5 + 14)))
	a1.shape.kind = nox_shape_kind(*(*uint8)(unsafe.Pointer(uintptr(v5 + 8))))
	a1.shape.circle_r = *(*float32)(unsafe.Pointer(uintptr(v5 + 52)))
	a1.shape.circle_r2 = *(*float32)(unsafe.Pointer(uintptr(v5 + 52))) * *(*float32)(unsafe.Pointer(uintptr(v5 + 52)))
	a1.shape.box_w = *(*float32)(unsafe.Pointer(uintptr(v5 + 64)))
	a1.shape.box_h = *(*float32)(unsafe.Pointer(uintptr(v5 + 68)))
	if a1.shape.kind == nox_shape_kind(NOX_SHAPE_BOX) {
		nox_shape_box_calc(&a1.shape)
	}
	a1.field_24 = *(*float32)(unsafe.Pointer(uintptr(v5 + 56)))
	a1.field_25 = *(*float32)(unsafe.Pointer(uintptr(v5 + 60)))
	v9 = *(*float32)(unsafe.Pointer(uintptr(v5 + 32)))
	if float64(v9) >= *(*float64)(unsafe.Pointer(&qword_581450_9568)) {
		*(*uint32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(v2)), 172)))) = 0
	} else {
		v9 = -v9
		*(*uint32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(v2)), 172)))) = 1
	}
	*(*float32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(v2)), 140)))) = v9
	if float64(v9) != 0.0 {
		nox_xxx_spriteChangeIntensity_484D70_light_intensity(int32(uintptr(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(v2)), 136))))), v9)
		if *v3 == 0 {
			*v3 = 1
			*v4 = math.MaxUint8
			*(*uint32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(v2)), 156)))) = math.MaxUint8
			*(*uint32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(v2)), 160)))) = math.MaxUint8
		}
	}
	if *(*uint32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(v2)), 112))))&0x13001000 != 0 {
		*(*uint32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(v2)), 432)))) = 0
		*(*uint32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(v2)), 436)))) = 0
		*(*uint32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(v2)), 440)))) = 0
		*(*uint32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(v2)), 444)))) = 0
		*(*uint16)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(v2)), 448)))) = math.MaxUint16
		*(*uint16)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(v2)), 450)))) = math.MaxUint16
	}
	var v7 *nox_thing = *(**nox_thing)(unsafe.Add(unsafe.Pointer(nox_things_array), unsafe.Sizeof((*nox_thing)(nil))*uintptr(i)))
	if v7.lifetime != 0 {
		nox_xxx_spriteTransparentDecay_49B950((*uint32)(v2), v7.lifetime)
	}
	return 1
}
func sub_44D2C0() {
	for cur := (*nox_thing)(nox_things_head); cur != nil; cur = cur.next {
		if int32(cur.shape_kind) == NOX_SHAPE_BOX {
			var (
				r float64 = float64(cur.shape_r)
				s nox_shape
			)
			s.kind = nox_shape_kind(NOX_SHAPE_BOX)
			s.circle_r = float32(r)
			s.circle_r2 = float32(r * r)
			nox_shape_box_calc(&s)
		}
	}
}
func nox_get_things_head() *nox_thing {
	return nox_things_head
}
func sub_44D320(a1 int32) int32 {
	var result int32
	if a1 != 0 {
		result = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 108))))
	} else {
		result = 0
	}
	return result
}
func sub_44D330(a1 *byte) int32 {
	return sub_44D340(a1)
}
func sub_44D340(a1 *byte) int32 {
	var (
		v1     int32
		v2     int32
		v3     *int32
		result int32
	)
	if a1 != nil && (func() bool {
		v1 = nox_xxx_keyFirstLetterNumberCli_44CD30(a1)
		return v1 >= 0
	}()) && (func() bool {
		v2 = int32(*memmap.PtrUint32(6112660, uint32(v1*4)+0xCAC98))
		return v2 >= 0
	}()) && (func() *int32 {
		v3 = (*int32)(libc.Search(unsafe.Pointer(a1), *(*unsafe.Pointer)(memmap.PtrOff(6112660, uint32(v1*4)+0xCAB58)), uint32(v2), 8, sub_44D020))
		return v3
	}()) != nil {
		result = *v3
	} else {
		result = 0
	}
	return result
}
func sub_44D3A0() {
	if dword_5d4594_831076 != 0 {
		switch dword_5d4594_830864 {
		case 0:
			if dword_5d4594_830872 != 0 && dword_587000_122848 != 0 {
				if sub_44D660(*(**byte)(unsafe.Pointer(&dword_5d4594_830872))) == 0 {
					goto LABEL_29
				}
				if dword_587000_122856 == 0 || *memmap.PtrUint32(0x587000, 93160) == 0 || dword_5d4594_831084 != 0 {
					goto LABEL_16
				}
				dword_5d4594_831084 = 1
				sub_43DBD0()
				dword_5d4594_830864 = 1
			} else if dword_5d4594_831084 != 0 {
				sub_43DBE0()
				dword_5d4594_831084 = 0
			}
		case 1:
			if sub_43DC40() == 0 {
			LABEL_16:
				dword_5d4594_830864 = 2
			}
		case 2:
			sub_486320((*uint32)(memmap.PtrOff(6112660, 830876)), 0x4000)
			if sub_44D7E0(*memmap.PtrInt32(6112660, 830868)) == 0 {
				dword_5d4594_830864 = 0
				goto LABEL_29
			}
			dword_5d4594_830864 = 3
			dword_5d4594_830972 = dword_5d4594_830872
			*memmap.PtrUint32(6112660, 830860) = *memmap.PtrUint32(6112660, 830868)
		case 3:
			if dword_587000_122848 == 0 || dword_5d4594_830972 == 0 || dword_5d4594_830872 != dword_5d4594_830972 || dword_5d4594_831088 == 0 || AIL_stream_status((HSTREAM)(unsafe.Pointer(uintptr(dword_5d4594_831088)))) == 2 {
				dword_5d4594_830864 = 4
				sub_486350(unsafe.Pointer(uintptr(int32(uintptr(memmap.PtrOff(6112660, 830876))))), 0)
			}
		case 4:
			if dword_5d4594_831088 == 0 || AIL_stream_status((HSTREAM)(unsafe.Pointer(uintptr(dword_5d4594_831088)))) == 2 || (*memmap.PtrUint32(6112660, 830880)&0xFFFF0000) == 0 {
				sub_44D640()
				dword_5d4594_830864 = 0
				if dword_5d4594_831084 != 0 {
					sub_43DBE0()
					dword_5d4594_831084 = 0
				}
				if dword_5d4594_830972 == dword_5d4594_830872 {
				LABEL_29:
					dword_5d4594_830872 = 0
				}
			}
		default:
		}
		sub_486520(memmap.PtrUint32(6112660, 830980))
		sub_486520(memmap.PtrUint32(6112660, 830876))
		if dword_5d4594_831088 != 0 && (int32(**(**uint8)(unsafe.Pointer(&dword_587000_81128)))&2 != 0 || int32(*memmap.PtrUint8(6112660, 830876))&2 != 0 || int32(*memmap.PtrUint8(6112660, 830980))&2 != 0) {
			sub_44D5C0(*(*int32)(unsafe.Pointer(&dword_5d4594_831088)), *memmap.PtrInt32(6112660, 830860))
		}
	}
}
func sub_44D5C0(a1 int32, a2 int32) {
	var v2 uint32
	if a1 != 0 {
		v2 = (*(*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(dword_587000_81128)) + 4))) >> 16) * ((uint32(*memmap.PtrUint16(6112660, 830882)) * ((uint32(a2) * uint32(*memmap.PtrUint16(6112660, 830986))) >> 14)) >> 14)
		*memmap.PtrUint32(6112660, 830876) &= 0xFFFFFFFD
		**(**uint32)(unsafe.Pointer(&dword_587000_122852)) &= 0xFFFFFFFD
		AIL_set_stream_volume((HSTREAM)(unsafe.Pointer(uintptr(a1))), int32((v2>>14)*math.MaxInt8)/100)
	}
}
func sub_44D8C0() {
	if dword_5d4594_831076 != 0 {
		sub_44D8F0()
		sub_44D640()
		sub_44D3A0()
		dword_5d4594_831076 = 0
	}
}
func nox_xxx_playDialogFile_44D900(a1 int32, a2 int32) int32 {
	var v2 int32
	if dword_587000_122848 != 0 && a1 != 0 {
		v2 = a2
		if a2 > 100 {
			v2 = 100
		}
		dword_5d4594_830872 = uint32(a1)
		*memmap.PtrUint32(6112660, 830868) = uint32(v2)
	}
	return 1
}
func sub_44D930() int32 {
	if dword_587000_122848 == 0 {
		return 0
	}
	if dword_5d4594_830872 != 0 || dword_5d4594_831088 != 0 {
		return 1
	}
	return 0
}
func sub_44D960() {
	dword_587000_122848 = 0
}
func sub_44D970() int32 {
	var result int32
	result = int32(dword_5d4594_831092)
	if dword_5d4594_831092 != 0 {
		dword_587000_122848 = 1
	}
	return result
}
func sub_44D990() int32 {
	return int32(dword_587000_122848)
}
func nox_client_initFade2_44D9A0() int32 {
	if dword_5d4594_823772 != 0 {
		return 1
	}
	nox_client_initFade_44D9D0()
	return 1
}
func nox_client_initFade_44D9D0() {
	var arr *struc_36 = (*struc_36)(memmap.PtrOff(6112660, 831108))
	for i := int32(0); i < 4; i++ {
		(*(*struc_36)(unsafe.Add(unsafe.Pointer(arr), unsafe.Sizeof(struc_36{})*uintptr(i)))).field_0 = 0
	}
}
func nox_client_procFade_44D9F0(a1 int32) int32 {
	var v1 int32 = 1
	if a1 != 0 {
		v1 = 5
	}
	*memmap.PtrUint32(6112660, 831100) = platformTicks()
	for i := int32(0); i < 4; i++ {
		var it *struc_36 = (*struc_36)(memmap.PtrOff(6112660, uint32(i*int32(unsafe.Sizeof(struc_36{})))+0xCAE84))
		if (it.field_0 & 5) != uint32(v1) {
			continue
		}
		it.field_6(it)
		if it.field_1 <= 0 {
			if (it.field_0 & 2) == 0 {
				it.field_0 &= uint32(^v1)
				if it.field_5 != nil {
					it.field_5()
				}
			}
		} else {
			it.field_1--
		}
	}
	return 0
}
func nox_client_fadeXxx_44DA60(a1 int32) int32 {
	if nox_client_checkFade_44DFD0(unsafe.Pointer(cgoFuncAddr(nox_client_drawFadingScreen_44DD70))) != 0 {
		return 1
	}
	var obj *struc_36 = nox_client_newFade_44DF50()
	if obj == nil {
		return 0
	}
	var v1 int32 = 3
	if a1 != 0 {
		v1 = 7
	}
	obj.field_0 = uint32(v1)
	obj.field_1 = 0
	obj.field_2 = 0
	obj.field_3 = 0
	obj.field_5 = nil
	obj.field_6 = func(arg1 *struc_36) {
		nox_xxx_cliClearScreen_44DDC0()
	}
	return 1
}
func nox_client_screenFadeTimeout_44DAB0(a1 int32, a2 int32, a3 func()) int32 {
	if nox_gameDisableMapDraw_5d4594_2650672 == 1 {
		if a3 != nil {
			a3()
		}
		return 1
	}
	nox_client_setScreenFade_44DF90(unsafe.Pointer(cgoFuncAddr(nox_xxx_screenFadeEffect_44DD20)))
	var obj *struc_36 = nox_client_newFade_44DF50()
	if obj == nil {
		return 0
	}
	var v3 int32 = 1
	if a2 != 0 {
		v3 = 5
	}
	obj.field_0 = uint32(v3)
	obj.field_1 = a1
	obj.field_2 = 0
	obj.field_3 = 0xFF0000 / uint32(a1)
	obj.field_5 = a3
	obj.field_6 = func(arg1 *struc_36) {
		nox_xxx_screenFadeEffect_44DD20(int32(uintptr(unsafe.Pointer(arg1))))
	}
	return 1
}
func nox_client_screenFadeXxx_44DB30(a1 int32, a2 int32, a3 func()) int32 {
	nox_client_setScreenFade_44DF90(unsafe.Pointer(cgoFuncAddr(nox_client_drawFadingScreen_44DD70)))
	var obj *struc_36 = nox_client_newFade_44DF50()
	if obj == nil {
		return 0
	}
	var v3 int32 = 1
	if a2 != 0 {
		v3 = 5
	}
	obj.field_0 = uint32(v3)
	obj.field_1 = a1
	obj.field_2 = 0
	obj.field_3 = 0xFF0000 / uint32(a1)
	obj.field_5 = a3
	obj.field_6 = func(arg1 *struc_36) {
		nox_client_drawFadingScreen_44DD70(int32(uintptr(unsafe.Pointer(arg1))))
	}
	nox_client_setScreenFade_44DF90(unsafe.Pointer(cgoFuncAddr(nox_xxx_cliClearScreen_44DDC0)))
	return 1
}
func sub_44DBA0(a1 int32, a2 int32, a3 func()) int32 {
	var obj *struc_36 = nox_client_newFade_44DF50()
	if obj == nil {
		return 0
	}
	var v3 int32 = 1
	if a2 != 0 {
		v3 = 5
	}
	obj.field_0 = uint32(v3)
	obj.field_1 = a1
	obj.field_2 = 0
	obj.field_3 = 0xFF0000 / uint32(a1)
	obj.field_5 = a3
	obj.field_6 = func(arg1 *struc_36) {
		sub_44DDF0(int32(uintptr(unsafe.Pointer(arg1))))
	}
	return 1
}
func sub_44DBF0(a1 int32, a2 int32, a3 func()) int32 {
	var obj *struc_36 = nox_client_newFade_44DF50()
	if obj == nil {
		return 0
	}
	var v3 int32 = 1
	if a2 != 0 {
		v3 = 5
	}
	obj.field_0 = uint32(v3)
	obj.field_1 = a1
	obj.field_2 = 0
	obj.field_3 = 0xFF0000 / uint32(a1)
	obj.field_5 = a3
	obj.field_6 = func(arg1 *struc_36) {
		sub_44DE30(int32(uintptr(unsafe.Pointer(arg1))))
	}
	return 1
}
func sub_44DC40(a1 int32, a2 int32) int32 {
	var obj *struc_36 = nox_client_newFade_44DF50()
	if obj == nil {
		return 0
	}
	var v4 int32 = a1 * nox_getBackbufHeight() / 100
	obj.field_0 = 3
	obj.field_1 = a2
	obj.field_2 = 0
	obj.field_3 = uint32((v4 << 16) / a2)
	obj.field_5 = nil
	obj.field_6 = func(arg1 *struc_36) {
		sub_44DE80(&arg1.field_0)
	}
	return 1
}
func sub_44DCA0(a1 int32, a2 int32) int32 {
	var obj *struc_36 = nox_client_newFade_44DF50()
	if obj == nil {
		return 0
	}
	var v4 int32 = a1 * nox_getBackbufHeight() / 100
	v4 <<= 16
	obj.field_0 = 1
	obj.field_1 = a2
	obj.field_2 = uint32(v4)
	obj.field_3 = uint32(v4 / a2)
	obj.field_4 = 0
	obj.field_5 = nil
	obj.field_6 = func(arg1 *struc_36) {
		sub_44DEE0(&arg1.field_0)
	}
	return 1
}
func sub_44DD00() int32 {
	return bool2int(nox_client_checkFade_44DFD0(unsafe.Pointer(cgoFuncAddr(sub_44DE80))) != 0)
}
func nox_xxx_screenFadeEffect_44DD20(a1 int32) int32 {
	var result int32
	nox_xxx_drawMakeRGB_433F10(uint8(int8(*(*int32)(unsafe.Pointer(uintptr(a1 + 8)))>>16)), uint8(int8(*(*int32)(unsafe.Pointer(uintptr(a1 + 8)))>>16)), uint8(int8(*(*int32)(unsafe.Pointer(uintptr(a1 + 8)))>>16)))
	nox_client_drawRectFadingScreen_49D0F0(0, 0, nox_getBackbufWidth(), nox_getBackbufHeight())
	result = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 12))))
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 8))) += uint32(result)
	return result
}
func nox_client_drawFadingScreen_44DD70(a1 int32) int32 {
	var (
		v1     int32
		result int32
		v3     uint8
	)
	v1 = a1
	v3 = uint8(int8(int32(-1 - (*(*int32)(unsafe.Pointer(uintptr(a1 + 8))) >> 16))))
	nox_xxx_drawMakeRGB_433F10(v3, v3, v3)
	nox_client_drawRectFadingScreen_49D0F0(0, 0, nox_getBackbufWidth(), nox_getBackbufHeight())
	result = int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 12))) + *(*uint32)(unsafe.Pointer(uintptr(v1 + 8))))
	*(*uint32)(unsafe.Pointer(uintptr(v1 + 8))) = uint32(result)
	return result
}
func nox_xxx_cliClearScreen_44DDC0() {
	nox_client_drawSetColor_434460(int32(nox_color_black_2650656))
	nox_client_drawRectFilledOpaque_49CE30(0, 0, nox_getBackbufWidth(), nox_getBackbufHeight())
}
func sub_44DDF0(a1 int32) *int4 {
	var v1 int32
	v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 12))) + *(*uint32)(unsafe.Pointer(uintptr(a1 + 8))))
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 8))) = uint32(v1)
	nox_xxx_drawMakeRGB_433F10(*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v1))), 2)), *(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v1))), 2)), *(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v1))), 2)))
	return sub_49D050(0, 0, nox_getBackbufWidth(), nox_getBackbufHeight())
}
func sub_44DE30(a1 int32) *int4 {
	var v1 int32
	v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 12))) + *(*uint32)(unsafe.Pointer(uintptr(a1 + 8))))
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 8))) = uint32(v1)
	nox_xxx_drawMakeRGB_433F10(uint8(int8(int32(-1-int32(*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v1))), 2)))))), uint8(int8(int32(-1-int32(*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v1))), 2)))))), uint8(int8(int32(-1-int32(*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v1))), 2)))))))
	return sub_49D050(0, 0, nox_getBackbufWidth(), nox_getBackbufHeight())
}
func sub_44DE80(a1 *uint32) {
	var v1 int32
	if *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*1)) != 0 {
		*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*2)) += *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*3))
	}
	v1 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*2)) >> 16)
	nox_set_color_rgb_434430(0, 0, 0)
	nox_client_drawRectFilledOpaque_49CE30(0, 0, nox_getBackbufWidth(), v1)
	nox_client_drawRectFilledOpaque_49CE30(0, nox_getBackbufHeight()-v1, nox_getBackbufWidth(), v1)
}
func sub_44DEE0(a1 *uint32) {
	var (
		v1 int32
		v2 int32
	)
	if *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*4)) == 0 {
		nox_client_setScreenFade_44DF90(unsafe.Pointer(cgoFuncAddr(sub_44DE80)))
		*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*4)) = 1
	}
	v1 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*2)) - *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*3)))
	*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*2)) = uint32(v1)
	v2 = v1 >> 16
	nox_set_color_rgb_434430(0, 0, 0)
	nox_client_drawRectFilledOpaque_49CE30(0, 0, nox_getBackbufWidth(), v2)
	nox_client_drawRectFilledOpaque_49CE30(0, nox_getBackbufHeight()-v2, nox_getBackbufWidth(), v2)
}
func nox_client_newFade_44DF50() *struc_36 {
	var (
		it *struc_36 = (*struc_36)(memmap.PtrOff(6112660, 831108))
		v1 int32     = 0
	)
	for it.field_0&1 != 0 {
		it = (*struc_36)(unsafe.Add(unsafe.Pointer(it), unsafe.Sizeof(struc_36{})*1))
		if func() int32 {
			p := &v1
			*p++
			return *p
		}() >= 4 {
			return nil
		}
	}
	return it
}
func sub_44DF70() {
	var arr *struc_36 = (*struc_36)(memmap.PtrOff(6112660, 831108))
	for i := int32(0); i < 4; i++ {
		arr.field_0 &= 0xFFFFFFFE
		arr = (*struc_36)(unsafe.Add(unsafe.Pointer(arr), unsafe.Sizeof(struc_36{})*1))
	}
}
func nox_client_setScreenFade_44DF90(a1 unsafe.Pointer) int32 {
	var (
		result int32     = 0
		it     *struc_36 = (*struc_36)(memmap.PtrOff(6112660, 831108))
	)
	for i := int32(0); i < 4; i++ {
		if it.field_0&1 != 0 && cgoFuncAddr(it.field_6) == cgoFuncAddr(cgoAsFunc(a1, (*func(*struc_36))(nil)).(func(*struc_36))) {
			it.field_0 &= 0xFFFFFFFE
			if it.field_5 != nil {
				it.field_5()
			}
			result = 1
		}
		it = (*struc_36)(unsafe.Add(unsafe.Pointer(it), unsafe.Sizeof(struc_36{})*1))
	}
	return result
}
func nox_client_checkFade_44DFD0(a1 unsafe.Pointer) int32 {
	var (
		it *struc_36 = (*struc_36)(memmap.PtrOff(6112660, 831108))
		v2 int32     = 0
	)
	for (it.field_0&1) == 0 || cgoFuncAddr(it.field_6) != cgoFuncAddr(cgoAsFunc(a1, (*func(*struc_36))(nil)).(func(*struc_36))) {
		it = (*struc_36)(unsafe.Add(unsafe.Pointer(it), unsafe.Sizeof(struc_36{})*1))
		if func() int32 {
			p := &v2
			*p++
			return *p
		}() >= 4 {
			return 0
		}
	}
	return bool2int(it != nil)
}
func sub_44E110() *uint32 {
	var (
		v0     *uint32
		v1     int32
		v2     *uint32
		v3     int32
		v4     *uint32
		v5     int32
		v6     *uint32
		v7     int32
		v8     *uint32
		v9     int32
		v10    *uint32
		v11    int32
		v12    *uint32
		v13    int32
		v14    *uint32
		v15    int32
		v16    *uint32
		v17    int32
		v18    *uint32
		v19    int32
		v20    *uint32
		v21    int32
		result *uint32
		v23    int32
	)
	if dword_5d4594_832484 == 0 {
		dword_5d4594_832484 = uint32(uintptr(guiFontPtrByName(CString("default"))))
	}
	v0 = *(**uint32)(unsafe.Pointer(&dword_5d4594_832496))
	if dword_5d4594_832496 == 0 {
		v1 = nox_xxx_getTTByNameSpriteMB_44CFC0("GauntletExitB")
		v0 = &nox_new_drawable_for_thing(v1).field_0
		dword_5d4594_832496 = uint32(uintptr(unsafe.Pointer(v0)))
	}
	*(*uint32)(unsafe.Add(unsafe.Pointer(v0), unsafe.Sizeof(uint32(0))*30)) |= 0x1000000
	v2 = *(**uint32)(unsafe.Pointer(&dword_5d4594_832492))
	if dword_5d4594_832492 == 0 {
		v3 = nox_xxx_getTTByNameSpriteMB_44CFC0("BeholderGenerator")
		v2 = &nox_new_drawable_for_thing(v3).field_0
		dword_5d4594_832492 = uint32(uintptr(unsafe.Pointer(v2)))
	}
	*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*30)) |= 0x1000000
	v4 = *(**uint32)(unsafe.Pointer(&dword_5d4594_832500))
	if dword_5d4594_832500 == 0 {
		v5 = nox_xxx_getTTByNameSpriteMB_44CFC0("Ankh")
		v4 = &nox_new_drawable_for_thing(v5).field_0
		dword_5d4594_832500 = uint32(uintptr(unsafe.Pointer(v4)))
	}
	*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*30)) |= 0x1000000
	v6 = *(**uint32)(unsafe.Pointer(&dword_5d4594_832504))
	if dword_5d4594_832504 == 0 {
		v7 = nox_xxx_getTTByNameSpriteMB_44CFC0("SoulGate")
		v6 = &nox_new_drawable_for_thing(v7).field_0
		dword_5d4594_832504 = uint32(uintptr(unsafe.Pointer(v6)))
	}
	*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*30)) |= 0x1000000
	v8 = *(**uint32)(unsafe.Pointer(&dword_5d4594_832508))
	if dword_5d4594_832508 == 0 {
		v9 = nox_xxx_getTTByNameSpriteMB_44CFC0("SilverKey")
		v8 = &nox_new_drawable_for_thing(v9).field_0
		dword_5d4594_832508 = uint32(uintptr(unsafe.Pointer(v8)))
	}
	*(*uint32)(unsafe.Add(unsafe.Pointer(v8), unsafe.Sizeof(uint32(0))*30)) |= 0x1000000
	v10 = *(**uint32)(unsafe.Pointer(&dword_5d4594_832512))
	if dword_5d4594_832512 == 0 {
		v11 = nox_xxx_getTTByNameSpriteMB_44CFC0("GoldKey")
		v10 = &nox_new_drawable_for_thing(v11).field_0
		dword_5d4594_832512 = uint32(uintptr(unsafe.Pointer(v10)))
	}
	*(*uint32)(unsafe.Add(unsafe.Pointer(v10), unsafe.Sizeof(uint32(0))*30)) |= 0x1000000
	v12 = *(**uint32)(unsafe.Pointer(&dword_5d4594_832516))
	if dword_5d4594_832516 == 0 {
		v13 = nox_xxx_getTTByNameSpriteMB_44CFC0("QuestGoldChest")
		v12 = &nox_new_drawable_for_thing(v13).field_0
		dword_5d4594_832516 = uint32(uintptr(unsafe.Pointer(v12)))
	}
	*(*uint32)(unsafe.Add(unsafe.Pointer(v12), unsafe.Sizeof(uint32(0))*30)) |= 0x1000000
	v14 = *(**uint32)(unsafe.Pointer(&dword_5d4594_832520))
	if dword_5d4594_832520 == 0 {
		v15 = nox_xxx_getTTByNameSpriteMB_44CFC0("QuestGoldPile")
		v14 = &nox_new_drawable_for_thing(v15).field_0
		dword_5d4594_832520 = uint32(uintptr(unsafe.Pointer(v14)))
	}
	*(*uint32)(unsafe.Add(unsafe.Pointer(v14), unsafe.Sizeof(uint32(0))*30)) |= 0x1000000
	v16 = *(**uint32)(unsafe.Pointer(&dword_5d4594_832524))
	if dword_5d4594_832524 == 0 {
		v17 = nox_xxx_getTTByNameSpriteMB_44CFC0("DunMirChest4")
		v16 = &nox_new_drawable_for_thing(v17).field_0
		dword_5d4594_832524 = uint32(uintptr(unsafe.Pointer(v16)))
	}
	*(*uint32)(unsafe.Add(unsafe.Pointer(v16), unsafe.Sizeof(uint32(0))*30)) |= 0x1000000
	v18 = *(**uint32)(unsafe.Pointer(&dword_5d4594_832528))
	if dword_5d4594_832528 == 0 {
		v19 = nox_xxx_getTTByNameSpriteMB_44CFC0("WarHammer")
		v18 = &nox_new_drawable_for_thing(v19).field_0
		dword_5d4594_832528 = uint32(uintptr(unsafe.Pointer(v18)))
	}
	*(*uint32)(unsafe.Add(unsafe.Pointer(v18), unsafe.Sizeof(uint32(0))*30)) |= 0x1000000
	v20 = *(**uint32)(unsafe.Pointer(&dword_5d4594_832532))
	if dword_5d4594_832532 == 0 {
		v21 = nox_xxx_getTTByNameSpriteMB_44CFC0("HastePotion")
		v20 = &nox_new_drawable_for_thing(v21).field_0
		dword_5d4594_832532 = uint32(uintptr(unsafe.Pointer(v20)))
	}
	*(*uint32)(unsafe.Add(unsafe.Pointer(v20), unsafe.Sizeof(uint32(0))*30)) |= 0x1000000
	result = *(**uint32)(unsafe.Pointer(&dword_5d4594_832536))
	if dword_5d4594_832536 == 0 {
		v23 = nox_xxx_getTTByNameSpriteMB_44CFC0("ConjurerSpellBook")
		result = &nox_new_drawable_for_thing(v23).field_0
		dword_5d4594_832536 = uint32(uintptr(unsafe.Pointer(result)))
	}
	*(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*30)) |= 0x1000000
	return result
}
func sub_44E3C0() {
	sub_450580()
	dword_5d4594_831260 = 0
	sub_413A00(0)
}
func nox_xxx_playGMCAPsmth_44E3E0() int32 {
	var result int32
	result = int32(dword_5d4594_831224)
	*memmap.PtrUint32(6112660, 831248) = 1
	if dword_5d4594_831224 != 0 {
		result = nox_xxx_playDialogFile_44D900(*(*int32)(unsafe.Pointer(&dword_5d4594_831240)), 100)
		dword_5d4594_831244 = 1
	}
	return result
}
func sub_44E560() *uint32 {
	var (
		result *uint32
		v1     *uint32
	)
	*memmap.PtrUint32(6112660, 831284) = uint32((nox_win_width - NOX_DEFAULT_WIDTH) / 2)
	*memmap.PtrUint32(6112660, 831288) = uint32((nox_win_height - NOX_DEFAULT_HEIGHT) / 2)
	result = (*uint32)(unsafe.Pointer(nox_window_new(nil, 56, 0, 0, nox_win_width, nox_win_height, nox_xxx_wndProc_44E6E0)))
	dword_5d4594_831236 = uint32(uintptr(unsafe.Pointer(result)))
	if result != nil {
		int32(uintptr(unsafe.Pointer(result))).setFunc93(nox_client_wndQuestBriefProc_44E630)
		*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_831236 + 56))) = nox_color_black_2650656
		nox_wnd_briefing_831232 = uint32(uintptr(unsafe.Pointer(newWindowFromFile("Briefing.wnd", nil))))
		result = (*uint32)(unsafe.Pointer(uintptr(nox_wnd_briefing_831232)))
		if result != nil {
			sub_46B120((*nox_window)(unsafe.Pointer(result)), (*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_831236))))))
			(*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&nox_wnd_briefing_831232)))).SetPos(image.Pt(*memmap.PtrInt32(6112660, 831284), *memmap.PtrInt32(6112660, 831288)))
			v1 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&nox_wnd_briefing_831232)))).ChildByID(1010)))
			int32(uintptr(unsafe.Pointer(v1))).setDraw(func(arg1 int32, arg2 int32) int32 {
				return sub_44E6F0((*uint32)(unsafe.Pointer(uintptr(arg1))), arg2)
			})
			sub_44E410()
			result = *(**uint32)(unsafe.Pointer(&dword_5d4594_831236))
		}
	}
	return result
}
func nox_client_wndQuestBriefProc_44E630(a1 int32, a2 int32, a3 int32, a4 int32) int32 {
	var v2 int32
	if *memmap.PtrUint32(6112660, 831248) != 0 && a2 != 18 && a2 != 17 && sub_41D1B0() != 1 {
		sub_450580()
		if dword_5d4594_831220 != 0 {
			if dword_5d4594_831220 == math.MaxUint8 {
				nox_gameDisableMapDraw_5d4594_2650672 = 0
				v2 = nox_client_getIntroScreenDuration_44E3B0()
				nox_client_screenFadeTimeout_44DAB0(v2, 1, sub_44E320)
				nox_gameDisableMapDraw_5d4594_2650672 = 1
			}
		} else {
			nox_xxx_wndClearCaptureMain((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_831236))))))
			guiFocus(nil)
			dword_5d4594_831256 = 1
			nox_savegame_sub_46D580()
		}
		(*(*int32)(unsafe.Pointer(&dword_5d4594_831236))).setFunc93(nil)
	}
	return 1
}
func nox_xxx_wndProc_44E6E0(a1 int32, a2 int32, a3 int32, a4 int32) int32 {
	return bool2int(a2 == 23)
}
func sub_44E6F0(a1 *uint32, xLeft int32) int32 {
	var (
		v2 int32
		v3 int32
	)
	*(*float32)(unsafe.Pointer(&dword_5d4594_831276)) = float32(float64(*(*float32)(unsafe.Pointer(&dword_5d4594_831276))) - sub_44E8B0())
	v2 = int(*(*float32)(unsafe.Pointer(&dword_5d4594_831276)))
	(*nox_window)(unsafe.Pointer(a1)).SetPos(image.Pt(0, v2))
	nox_client_copyRect_49F6F0(*memmap.PtrInt32(6112660, 831284), *memmap.PtrInt32(6112660, 831288), NOX_DEFAULT_WIDTH, NOX_DEFAULT_HEIGHT)
	if int32(*memmap.PtrUint8(6112660, 832472))&1 != 0 {
		sub_44E8E0(int32(uintptr(unsafe.Pointer(a1))), xLeft)
	} else if int32(*memmap.PtrUint8(6112660, 832472))&2 != 0 {
		sub_44F0F0(int32(uintptr(unsafe.Pointer(a1))), xLeft)
	} else if int32(*memmap.PtrUint8(6112660, 832472))&4 != 0 {
		sub_44F300(int32(uintptr(unsafe.Pointer(a1))), xLeft)
	} else {
		nox_xxx_wndStaticDrawNoImage_488D00((*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1)))))), (*nox_window_data)(unsafe.Pointer(uintptr(xLeft))))
	}
	nox_client_copyRect_49F6F0(0, 0, nox_getBackbufWidth(), nox_getBackbufHeight())
	if int64(*(*float32)(unsafe.Pointer(&dword_5d4594_831276))) <= int64(*memmap.PtrInt32(6112660, 831280)) && dword_5d4594_831244 == 1 && sub_44D930() == 0 && (uint64(platformTicks())-*memmap.PtrUint64(6112660, 831292)) > uint64(nox_client_getBriefDuration()) && (*memmap.PtrUint32(6112660, 832488) == 1 || (int32(*memmap.PtrUint8(6112660, 832472))&5) == 0) {
		if int32(*memmap.PtrUint8(6112660, 832472))&2 != 0 && dword_5d4594_832480 != 0 {
			clientPlaySoundSpecial(582, 100)
		}
		sub_450580()
		if dword_5d4594_831220 != 0 {
			nox_gameDisableMapDraw_5d4594_2650672 = 0
			v3 = nox_client_getIntroScreenDuration_44E3B0()
			nox_client_screenFadeTimeout_44DAB0(v3, 1, sub_44E320)
			nox_gameDisableMapDraw_5d4594_2650672 = 1
		} else if dword_5d4594_831256 == 0 {
			nox_xxx_wndClearCaptureMain((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_831236))))))
			dword_5d4594_831256 = 1
			nox_savegame_sub_46D580()
			(*(*int32)(unsafe.Pointer(&dword_5d4594_831236))).setFunc93(nil)
		}
		int32(uintptr(unsafe.Pointer(a1))).setDraw(func(arg1 int32, arg2 int32) int32 {
			return sub_44E8D0()
		})
	}
	return 1
}
func sub_44E8B0() float64 {
	var result float64
	if dword_5d4594_831220 == math.MaxUint8 {
		result = 1.0
	} else {
		result = 0.0
	}
	return result
}
func sub_44E8D0() int32 {
	return 1
}
func nox_client_lockScreenBriefing_450160(a1 int32, a2 int32, a3 int8) int32 {
	var (
		v3     int32
		v4     int32
		v5     *uint32
		v6     int32
		v7     *byte
		v8     *byte
		v9     *byte
		v10    int32
		v11    int32
		v12    int32
		result int32
		v14    int32
		v15    **uint16
	)
	v3 = a1
	dword_5d4594_831260 = 1
	dword_5d4594_831244 = 1
	v14 = int32(*memmap.PtrUint32(0x8531A0, 2576))
	*memmap.PtrUint32(6112660, 832488) = 0
	*memmap.PtrUint32(6112660, 832472) = 0
	dword_5d4594_831256 = 0
	*memmap.PtrUint8(6112660, 831252) = uint8(int8(a1 + 1))
	if *memmap.PtrUint32(0x8531A0, 2576) != 0 {
		v4 = int32(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(0x8531A0, 2576) + 2251))))
	} else {
		v4 = a2
	}
	nox_xxx_wndShowModalMB((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_831236))))))
	nox_xxx_wndSetCaptureMain((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_831236))))))
	guiFocus((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_831236))))))
	(*(*int32)(unsafe.Pointer(&dword_5d4594_831236))).setFunc93(nox_client_wndQuestBriefProc_44E630)
	v5 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&nox_wnd_briefing_831232)))).ChildByID(1010)))
	int32(uintptr(unsafe.Pointer(v5))).setDraw(func(arg1 int32, arg2 int32) int32 {
		return sub_44E6F0((*uint32)(unsafe.Pointer(uintptr(arg1))), arg2)
	})
	dword_5d4594_831220 = uint32(a2)
	v15 = (**uint16)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(uint32(0))*8)))))
	if v3 == math.MaxUint8 {
		nox_xxx_wndSetIcon_46AE60(*(*int32)(unsafe.Pointer(&nox_wnd_briefing_831232)), 0)
		(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v5)))))).Func94(asWindowEvent(0x4001, *memmap.PtrInt32(6112660, 831268), 0))
		nox_xxx_wndClearFlag_46AD80(int32(uintptr(unsafe.Pointer(v5))), 8192)
		dword_5d4594_831240 = uint32(uintptr(unsafe.Pointer(nox_xxx_GetEndgameDialog_578D80())))
		v6 = 24
		dword_5d4594_831220 = math.MaxUint8
	} else if v3 == 254 {
		nox_server_currentMapGetFilename_409B30()
		v7 = (*byte)(unsafe.Pointer(nox_xxx_gLoadImg("GauntletStartMines")))
		nox_xxx_wndSetIcon_46AE60(*(*int32)(unsafe.Pointer(&nox_wnd_briefing_831232)), int32(uintptr(unsafe.Pointer(v7))))
		(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v5)))))).Func94(asWindowEvent(0x4001, int32(uintptr(memmap.PtrOff(6112660, 832540))), 0))
		v6 = 0
		if int32(a3)&1 != 0 {
			*memmap.PtrUint32(6112660, 832472) = 1
			v8 = (*byte)(unsafe.Pointer(nox_xxx_gLoadImg("MenuSystemBG")))
			nox_xxx_wndSetIcon_46AE60(*(*int32)(unsafe.Pointer(&nox_wnd_briefing_831232)), int32(uintptr(unsafe.Pointer(v8))))
			nox_client_resetScreenParticles_431510()
			nox_xxx_bookHideMB_45ACA0(1)
		} else if int32(a3)&4 != 0 {
			*memmap.PtrUint32(6112660, 832472) = 4
			v9 = (*byte)(unsafe.Pointer(nox_xxx_gLoadImg("GauntletInstructionBackground")))
			nox_xxx_wndSetIcon_46AE60(*(*int32)(unsafe.Pointer(&nox_wnd_briefing_831232)), int32(uintptr(unsafe.Pointer(v9))))
			nox_client_resetScreenParticles_431510()
			nox_xxx_bookHideMB_45ACA0(1)
		} else {
			*memmap.PtrUint32(6112660, 832472) = 2
			nox_xxx_wndSetIcon_46AE60(*(*int32)(unsafe.Pointer(&nox_wnd_briefing_831232)), *memmap.PtrInt32(6112660, 832460))
			if *memmap.PtrUint32(6112660, 832464) != 0 {
				(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v5)))))).Func94(asWindowEvent(0x4001, *memmap.PtrInt32(6112660, 832464), 0))
			}
			nox_client_resetScreenParticles_431510()
			nox_xxx_bookHideMB_45ACA0(1)
		}
		sub_446780()
		dword_5d4594_831240 = 0
		dword_5d4594_831244 = 1
		dword_5d4594_831224 = 0
	} else {
		v10 = (v3 + v4 + v4*10) * 32
		if a2 != 0 {
			nox_xxx_wndSetIcon_46AE60(*(*int32)(unsafe.Pointer(&nox_wnd_briefing_831232)), int32(*memmap.PtrUint32(6112660, uint32(v10)+831300)))
			(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v5)))))).Func94(asWindowEvent(0x4001, int32(*memmap.PtrUint32(6112660, uint32(v10)+0xCAF48)), 0))
			v11 = int32(*memmap.PtrUint32(6112660, uint32(v10)+0xCAF4C))
			v6 = int32(*memmap.PtrUint32(6112660, uint32(v10)+0xCAF50))
		} else {
			nox_xxx_wndSetIcon_46AE60(*(*int32)(unsafe.Pointer(&nox_wnd_briefing_831232)), int32(*memmap.PtrUint32(6112660, uint32(v10)+0xCAF54)))
			(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v5)))))).Func94(asWindowEvent(0x4001, int32(*memmap.PtrUint32(6112660, uint32(v10)+831320)), 0))
			v11 = int32(*memmap.PtrUint32(6112660, uint32(v10)+0xCAF5C))
			v6 = int32(*memmap.PtrUint32(6112660, uint32(v10)+0xCAF60))
		}
		dword_5d4594_831240 = uint32(v11)
	}
	nox_xxx_drawGetStringSize_43F840(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(uint32(0))*59)))), (*libc.WChar)(unsafe.Pointer(*v15)), nil, (*int32)(memmap.PtrOff(6112660, 831280)), NOX_DEFAULT_HEIGHT)
	sub_46AB20(v5, NOX_DEFAULT_WIDTH, *memmap.PtrInt32(6112660, 831280))
	if v3 == math.MaxUint8 {
		dword_5d4594_831276 = 0x43FA0000
		*memmap.PtrUint32(6112660, 831280) = 0xFFFFFFEC - *memmap.PtrUint32(6112660, 831280)
	} else {
		*memmap.PtrUint32(6112660, 831280) = uint32((480 - nox_xxx_guiFontHeightMB_43F320(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(uint32(0))*59))))) - *memmap.PtrInt32(6112660, 831280)) / 2)
		*(*float32)(unsafe.Pointer(&dword_5d4594_831276)) = float32(float64(*memmap.PtrInt32(6112660, 831280)))
	}
	sub_431290()
	dword_5d4594_831224 = 0
	if a2 != 0 || v3 == math.MaxUint8 {
		if v3 != 254 {
			dword_5d4594_831224 = 1
		}
	} else if v3 != 254 {
		if *(*uint32)(unsafe.Pointer(uintptr(v14 + v3*4 + 4408))) != 0 {
			nox_xxx_wndClearCaptureMain((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_831236))))))
			dword_5d4594_831256 = 1
			nox_savegame_sub_46D580()
			(*(*int32)(unsafe.Pointer(&dword_5d4594_831236))).setFunc93(nil)
		} else {
			dword_5d4594_831224 = 1
			*(*uint32)(unsafe.Pointer(uintptr(v14 + v3*4 + 4408))) = 1
		}
	}
	sub_43DD70(v6, 50)
	*memmap.PtrUint64(6112660, 831292) = uint64(platformTicks())
	*memmap.PtrUint32(6112660, 831248) = 0
	v12 = nox_client_getIntroScreenDuration_44E3B0()
	nox_client_screenFadeXxx_44DB30(v12, 1, func() {
		nox_xxx_playGMCAPsmth_44E3E0()
	})
	result = int32(dword_5d4594_831224)
	if dword_5d4594_831224 != 0 {
		dword_5d4594_831244 = 0
	}
	return result
}
func sub_450560() int32 {
	return int32(dword_5d4594_831260)
}
func sub_450570() int32 {
	return int32(dword_5d4594_831220)
}
func sub_4505E0() int32 {
	var result int32
	if dword_5d4594_831236 != 0 {
		(*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_831236)))).Destroy()
		dword_5d4594_831236 = 0
		nox_wnd_briefing_831232 = 0
	}
	dword_5d4594_831260 = 0
	dword_5d4594_832484 = 0
	if dword_5d4594_832504 != 0 {
		nox_xxx_spriteDelete_45A4B0((*nox_drawable)(unsafe.Pointer(*(**uint64)(unsafe.Pointer(&dword_5d4594_832504)))))
	}
	dword_5d4594_832504 = 0
	if dword_5d4594_832492 != 0 {
		nox_xxx_spriteDelete_45A4B0((*nox_drawable)(unsafe.Pointer(*(**uint64)(unsafe.Pointer(&dword_5d4594_832492)))))
	}
	dword_5d4594_832492 = 0
	if dword_5d4594_832500 != 0 {
		nox_xxx_spriteDelete_45A4B0((*nox_drawable)(unsafe.Pointer(*(**uint64)(unsafe.Pointer(&dword_5d4594_832500)))))
	}
	dword_5d4594_832500 = 0
	if dword_5d4594_832496 != 0 {
		nox_xxx_spriteDelete_45A4B0((*nox_drawable)(unsafe.Pointer(*(**uint64)(unsafe.Pointer(&dword_5d4594_832496)))))
	}
	dword_5d4594_832496 = 0
	if dword_5d4594_832508 != 0 {
		nox_xxx_spriteDelete_45A4B0((*nox_drawable)(unsafe.Pointer(*(**uint64)(unsafe.Pointer(&dword_5d4594_832508)))))
	}
	dword_5d4594_832508 = 0
	if dword_5d4594_832512 != 0 {
		nox_xxx_spriteDelete_45A4B0((*nox_drawable)(unsafe.Pointer(*(**uint64)(unsafe.Pointer(&dword_5d4594_832512)))))
	}
	dword_5d4594_832512 = 0
	if dword_5d4594_832516 != 0 {
		nox_xxx_spriteDelete_45A4B0((*nox_drawable)(unsafe.Pointer(*(**uint64)(unsafe.Pointer(&dword_5d4594_832516)))))
	}
	dword_5d4594_832516 = 0
	if dword_5d4594_832520 != 0 {
		nox_xxx_spriteDelete_45A4B0((*nox_drawable)(unsafe.Pointer(*(**uint64)(unsafe.Pointer(&dword_5d4594_832520)))))
	}
	dword_5d4594_832520 = 0
	if dword_5d4594_832524 != 0 {
		nox_xxx_spriteDelete_45A4B0((*nox_drawable)(unsafe.Pointer(*(**uint64)(unsafe.Pointer(&dword_5d4594_832524)))))
	}
	dword_5d4594_832524 = 0
	if dword_5d4594_832528 != 0 {
		nox_xxx_spriteDelete_45A4B0((*nox_drawable)(unsafe.Pointer(*(**uint64)(unsafe.Pointer(&dword_5d4594_832528)))))
	}
	dword_5d4594_832528 = 0
	if dword_5d4594_832532 != 0 {
		nox_xxx_spriteDelete_45A4B0((*nox_drawable)(unsafe.Pointer(*(**uint64)(unsafe.Pointer(&dword_5d4594_832532)))))
	}
	result = int32(dword_5d4594_832536)
	dword_5d4594_832532 = 0
	if dword_5d4594_832536 != 0 {
		result = nox_xxx_spriteDelete_45A4B0((*nox_drawable)(unsafe.Pointer(*(**uint64)(unsafe.Pointer(&dword_5d4594_832536)))))
	}
	dword_5d4594_832536 = 0
	return result
}
func sub_450750() uint8 {
	return *memmap.PtrUint8(6112660, 831252)
}
func sub_450760(a1 int8) int8 {
	var result int8
	result = a1
	*memmap.PtrUint8(6112660, 831252) = uint8(a1)
	return result
}
func sub_450960(a1 unsafe.Pointer, a2 unsafe.Pointer) int32 {
	var (
		v2     uint32
		v3     uint32
		result int32
	)
	v2 = *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(a1)), unsafe.Sizeof(uint32(0))*3)))
	v3 = *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(a2)), unsafe.Sizeof(uint32(0))*3)))
	if v2 == v3 {
		result = 0
	} else {
		if v2 < v3 {
			result = 1
		} else {
			result = -1
		}
	}
	return result
}
func sub_450AD0(a1 *byte) *byte {
	var result *byte
	result = a1
	if a1 == nil {
		result = (*byte)(unsafe.Pointer(nox_xxx_gLoadImg("WarriorChapterBegin8")))
	}
	*memmap.PtrUint32(6112660, 832460) = uint32(uintptr(unsafe.Pointer(result)))
	return result
}
func sub_450AF0(a1 int32) int32 {
	var result int32
	result = a1
	*memmap.PtrUint32(6112660, 832464) = uint32(a1)
	return result
}
func nox_gui_setQuestStage_450B00(a1 int32) {
	*memmap.PtrUint32(6112660, 832468) = uint32(a1)
}
func nox_gui_getQuestStage_450B10() int32 {
	return int32(*memmap.PtrUint32(6112660, 832468))
}
func sub_451850(a2 int32, a3 int32) int32 {
	var (
		v2     int32
		v3     *uint8
		result int32
	)
	v2 = 0
	v3 = (*uint8)(memmap.PtrOff(6112660, 840712))
	for {
		sub_451920((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v3))), -int(unsafe.Sizeof(uint32(0))*21))))
		*(*uint32)(unsafe.Pointer(v3)) = uint32(uintptr(unsafe.Pointer(nox_xxx_getSndName_40AF80(v2))))
		v3 = (*uint8)(unsafe.Add(unsafe.Pointer(v3), 200))
		v2++
		if int32(uintptr(unsafe.Pointer(v3))) >= int32(uintptr(memmap.PtrOff(6112660, 1045312))) {
			break
		}
	}
	dword_5d4594_1045420 = uint32(a3)
	dword_5d4594_1045428 = uint32(a2)
	if a3 != 0 {
		dword_5d4594_1045424 = uint32(uintptr(unsafe.Pointer(sub_4BD340(a3, 0x100000, 200, 8192))))
		dword_5d4594_1045436 = uint32(uintptr(unsafe.Pointer(sub_4BD280(200, 576))))
	}
	if dword_5d4594_1045424 == 0 || dword_5d4594_1045420 == 0 || dword_5d4594_1045428 == 0 || dword_5d4594_1045436 == 0 {
		return 0
	}
	nox_common_list_clear_425760((*nox_list_item_t)(memmap.PtrOff(6112660, 840612)))
	sub_4864A0((*uint32)(memmap.PtrOff(6112660, 1045228)))
	result = 1
	*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1045428 + 184))) = uint32(uintptr(memmap.PtrOff(6112660, 1045228)))
	dword_5d4594_1045432 = 1
	return result
}
func sub_451920(a2 *uint32) int32 {
	*a2 = 0
	*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*1)) = 0
	*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*2)) = 0
	*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*14)) = 0
	*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*15)) = 0
	*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*19)) = 0
	*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*20)) = 0
	*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*12)) = 1
	*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*48)) = 0
	*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*18)) = 0
	*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*17)) = 0
	*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*25)) = 0
	*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*26)) = 0
	*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*16)) = 600
	return sub_4862E0(int32(uintptr(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*4))))), 0x4000)
}
func sub_451970() {
	sub_4521F0()
	sub_452230()
	if dword_5d4594_1045424 != 0 {
		sub_4BD3C0(*(*unsafe.Pointer)(unsafe.Pointer(&dword_5d4594_1045424)))
		dword_5d4594_1045424 = 0
	}
	if dword_5d4594_1045436 != 0 {
		sub_4BD2D0(*(*unsafe.Pointer)(unsafe.Pointer(&dword_5d4594_1045436)))
		dword_5d4594_1045436 = 0
	}
	dword_5d4594_1045432 = 0
}
func sub_4519C0() int32 {
	var (
		result int32
		v1     int32
		v2     int32
		v3     int32
		v4     int32
		v5     *uint8
		v6     *uint8
		v7     *uint8
		v8     int32
		v9     int32
		v10    int32
	)
	result = int32(dword_5d4594_1045432)
	if dword_5d4594_1045432 != 0 {
		result = int32(*memmap.PtrUint32(6112660, 1045448))
		if *memmap.PtrUint32(6112660, 1045448) == 0 {
			*memmap.PtrUint32(6112660, 1045448) = 1
			sub_486520(*(**uint32)(unsafe.Pointer(&dword_587000_127004)))
			v1 = int32(*memmap.PtrUint32(6112660, 840612))
			*memmap.PtrUint32(6112660, 1045440)++
			if unsafe.Pointer(*(**uint8)(memmap.PtrOff(6112660, 840612))) != memmap.PtrOff(6112660, 840612) {
				for {
					v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 36))))
					if *(*uint32)(unsafe.Pointer(uintptr(v2 + 100))) != *memmap.PtrUint32(6112660, 1045440) {
						nox_common_list_clear_425760((*nox_list_item_t)(unsafe.Pointer(uintptr(v2 + 88))))
						*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v1 + 36))) + 52))) = 0
						*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v1 + 36))) + 100))) = *memmap.PtrUint32(6112660, 1045440)
					}
					sub_486520((*uint32)(unsafe.Pointer(uintptr(v1 + 184))))
					if *(*uint32)(unsafe.Pointer(uintptr(v1 + 28))) != 4 {
						sub_451BE0(v1)
					}
					v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(v1))))
					if unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(v1)))) == memmap.PtrOff(6112660, 840612) {
						break
					}
				}
				v1 = int32(*memmap.PtrUint32(6112660, 840612))
				if unsafe.Pointer(*(**uint8)(memmap.PtrOff(6112660, 840612))) != memmap.PtrOff(6112660, 840612) {
					for {
						sub_452510(v1)
						v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(v1))))
						if unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(v1)))) == memmap.PtrOff(6112660, 840612) {
							break
						}
					}
					v1 = int32(*memmap.PtrUint32(6112660, 840612))
				}
			}
			v3 = 0
			sub_452010()
			if unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(v1)))) == memmap.PtrOff(6112660, 840612) {
				goto LABEL_35
			}
			for {
				v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 176))))
				v5 = *(**uint8)(unsafe.Pointer(uintptr(v1)))
				if v4 == 0 || uint32(v1) != *(*uint32)(unsafe.Pointer(uintptr(v4 + 152))) {
					sub_4523D0((*uint32)(unsafe.Pointer(uintptr(v1))))
				}
				if int32(*(*uint8)(unsafe.Pointer(uintptr(v1 + 24))))&1 != 0 {
					sub_451FE0(v1)
				} else {
					v3 += int32(((*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v1 + 36))) + 20))) >> 16) * 33) >> 14)
					sub_452050((*uint32)(unsafe.Pointer(uintptr(v1))))
				}
				v1 = int32(uintptr(unsafe.Pointer(v5)))
				if unsafe.Pointer(v5) == memmap.PtrOff(6112660, 840612) {
					break
				}
			}
			if v3 <= 100 {
			LABEL_35:
				sub_486350(unsafe.Pointer(uintptr(int32(uintptr(memmap.PtrOff(6112660, 1045228))))), 0x4000)
			} else {
				sub_486350(unsafe.Pointer(uintptr(int32(uintptr(memmap.PtrOff(6112660, 1045228))))), int32(0x190000/uint32(v3)))
			}
			result = sub_486520(memmap.PtrUint32(6112660, 1045228))
			v6 = *(**uint8)(memmap.PtrOff(6112660, 840612))
			if unsafe.Pointer(*(**uint8)(memmap.PtrOff(6112660, 840612))) != memmap.PtrOff(6112660, 840612) {
				for {
					v7 = *(**uint8)(unsafe.Pointer(v6))
					result = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v6))), unsafe.Sizeof(uint32(0))*7))))
					if result == 1 {
						sub_451DC0(int32(uintptr(unsafe.Pointer(v6))))
						v8 = sub_451CA0((*uint32)(unsafe.Pointer(v6)))
						*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v6))), unsafe.Sizeof(uint32(0))*74))) = uint32(v8)
						if v8 == 0 {
							for {
								if sub_452120(int32(uintptr(unsafe.Pointer(v6)))) == nil {
									break
								}
								v7 = *(**uint8)(unsafe.Pointer(v6))
								sub_451DC0(int32(uintptr(unsafe.Pointer(v6))))
								v9 = sub_451CA0((*uint32)(unsafe.Pointer(v6)))
								*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v6))), unsafe.Sizeof(uint32(0))*74))) = uint32(v9)
								if v9 != 0 {
									break
								}
							}
						}
						v10 = sub_451CA0((*uint32)(unsafe.Pointer(v6)))
						*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v6))), unsafe.Sizeof(uint32(0))*74))) = uint32(v10)
						if v10 == 0 || (func() int32 {
							result = sub_452490((*uint32)(unsafe.Pointer(v6)))
							return result
						}()) == 0 {
							sub_4523D0((*uint32)(unsafe.Pointer(v6)))
							result = sub_451FE0(int32(uintptr(unsafe.Pointer(v6))))
						}
					}
					v6 = v7
					if unsafe.Pointer(v7) == memmap.PtrOff(6112660, 840612) {
						break
					}
				}
			}
			*memmap.PtrUint32(6112660, 1045448) = 0
		}
	}
	return result
}
func sub_451BE0(a1 int32) int32 {
	var (
		v1     int32
		v2     int32
		v3     uint32
		v4     *uint32
		v5     int32
		v6     int32
		v7     *uint32
		result int32
		v9     int32
		v10    *uint32
	)
	v1 = a1
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 36))))
	v3 = *(*uint32)(unsafe.Pointer(uintptr(a1 + 188))) >> 16
	v4 = *(**uint32)(unsafe.Pointer(uintptr(v2 + 88)))
	if v4 != (*uint32)(unsafe.Pointer(uintptr(v2+88))) {
		for {
			v5 = int32((*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*44)) >> 16) - v3)
			if v5 < 0 {
				v5 = int32(v3 - (*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*44)) >> 16))
			}
			if uint32(v5) >= (*(*uint32)(unsafe.Pointer(uintptr(v2 + 20)))>>16)/10 {
				if *(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*44))>>16 < v3 {
					break
				}
			} else {
				v6 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*4)))
				if int32(*(*uint8)(unsafe.Pointer(uintptr(v2 + 4))))&16 != 0 {
					if v6 != 0 {
						break
					}
				} else if v6 == 0 {
					break
				}
			}
			v4 = (*uint32)(unsafe.Pointer(uintptr(*v4)))
			if v4 == (*uint32)(unsafe.Pointer(uintptr(v2+88))) {
				break
			}
		}
		v1 = a1
	}
	v7 = (*uint32)(unsafe.Pointer(uintptr(v1 + 12)))
	sub_425770(unsafe.Pointer(uintptr(v1 + 12)))
	nox_common_list_append_4258E0((*nox_list_item_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v4)))))), (*nox_list_item_t)(unsafe.Pointer(v7)))
	result = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 56))))
	v9 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 52))) + 1)
	*(*uint32)(unsafe.Pointer(uintptr(v2 + 52))) = uint32(v9)
	if result != 0 {
		if v9 > result {
			v10 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 92))) - 12)))
			sub_425920(*(***uint32)(unsafe.Pointer(uintptr(v2 + 92))))
			sub_4523D0(v10)
			result = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 52))) - 1)
			*(*uint32)(unsafe.Pointer(uintptr(v2 + 52))) = uint32(result)
		}
	}
	return result
}
func sub_451CA0(a1 *uint32) int32 {
	var (
		v1 int32
		v3 int32
		v4 *uint32
	)
	v1 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*42)))
	*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*108)) = uint32(v1)
	if v1 == 0 {
		return 0
	}
	v3 = 0
	if v1 > 0 {
		v4 = (*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*76))
		for {
			*v4 = uint32(func() int32 {
				p := &v3
				x := *p
				*p++
				return x
			}())
			v4 = (*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*1))
			if uint32(v3) >= *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*108)) {
				break
			}
		}
	}
	*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*43)) = math.MaxUint32
	return sub_451CF0(a1)
}
func sub_451F30(a1 int32, a2 int32) int32 {
	var (
		v2     int32
		result int32
	)
	*(*uint32)(unsafe.Pointer(uintptr(uint32(a1) + *(*uint32)(unsafe.Pointer(uintptr(a1 + 168)))*4 + 40))) = uint32(uintptr(unsafe.Pointer(sub_4BD470(*(***uint32)(unsafe.Pointer(&dword_5d4594_1045424)), int32(*(*int16)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 36))) + uint32(a2*2) + 128))))))))
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 168))))
	result = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + v2*4 + 40))))
	if result != 0 {
		sub_4BD650(int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + v2*4 + 40)))))
		result = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 168))) + 1)
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 168))) = uint32(result)
	}
	return result
}
func sub_451F90(a1 int32) int32 {
	var (
		v1     int32
		result int32
		v3     *int32
	)
	v1 = 0
	result = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 168))))
	if result <= 0 {
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 168))) = 0
	} else {
		v3 = (*int32)(unsafe.Pointer(uintptr(a1 + 40)))
		for {
			sub_4BD660(*v3)
			*v3 = 0
			result = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 168))))
			v1++
			v3 = (*int32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(int32(0))*1))
			if v1 >= result {
				break
			}
		}
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 168))) = 0
	}
	return result
}
func sub_451FE0(a1 int32) int32 {
	sub_425920((**uint32)(unsafe.Pointer(uintptr(a1))))
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 280))) = 0
	return sub_4BD300(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045436)), a1)
}
func sub_452010() int32 {
	var (
		v0 *uint8
		v1 int32
		v2 int32
	)
	v0 = (*uint8)(memmap.PtrOff(6112660, 839892))
	v1 = 6
	for {
		v2 = 10
		for {
			nox_common_list_clear_425760((*nox_list_item_t)(unsafe.Pointer(v0)))
			v0 = (*uint8)(unsafe.Add(unsafe.Pointer(v0), 12))
			v2--
			if v2 == 0 {
				break
			}
		}
		v1--
		if v1 == 0 {
			break
		}
	}
	return int32(func() uint32 {
		p := memmap.PtrUint32(6112660, 1045444)
		*p++
		return *p
	}())
}
func sub_452050(a1 *uint32) {
	var (
		v1     *uint32
		v2     int32
		v3     uint32
		v4     *uint8
		result *uint32
		v6     **uint32
		v7     **uint32
		v8     *uint32
	)
	v1 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*9)))))
	v2 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*12)) + *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*75)))
	v3 = (*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*47)) >> 16) / 1638
	v4 = (*uint8)(memmap.PtrOff(6112660, uint32(v2*120)+0xCD0D4))
	if *(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*26)) == *memmap.PtrUint32(6112660, 1045444) {
		result = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*27)))))
		if v2 <= int32(uintptr(unsafe.Pointer(result))) {
			if (*uint32)(unsafe.Pointer(uintptr(v2))) == result && v3 > *(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*31)) {
				*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*31)) = v3
				v7 = (**uint32)(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*28))))
				sub_425920(v7)
				nox_common_list_append_4258E0((*nox_list_item_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v4), v3*12)))))))), (*nox_list_item_t)(unsafe.Pointer(v7)))
			}
		} else {
			*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*27)) = uint32(v2)
			*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*31)) = v3
			v6 = (**uint32)(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*28))))
			sub_425920(v6)
			nox_common_list_append_4258E0((*nox_list_item_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v4), v3*12)))))))), (*nox_list_item_t)(unsafe.Pointer(v6)))
		}
	} else {
		*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*26)) = *memmap.PtrUint32(6112660, 1045444)
		*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*27)) = uint32(v2)
		*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*31)) = v3
		v8 = (*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*28))
		sub_425770(unsafe.Pointer(v8))
		nox_common_list_append_4258E0((*nox_list_item_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v4), v3*12)))))))), (*nox_list_item_t)(unsafe.Pointer(v8)))
	}
}
func sub_452120(a1 int32) *int32 {
	var (
		v1     int32
		result *int32
		v3     *int32
		v4     *uint8
		v5     *uint8
	)
	v1 = 0
	result = sub_4521A0(int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 300))) + *(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 36))) + 48)))))
	v3 = result
	if result != nil {
		sub_452190(int32(uintptr(unsafe.Pointer(result))))
		v4 = *(**uint8)(memmap.PtrOff(6112660, 840612))
		if unsafe.Pointer(*(**uint8)(memmap.PtrOff(6112660, 840612))) != memmap.PtrOff(6112660, 840612) {
			for {
				v5 = *(**uint8)(unsafe.Pointer(v4))
				if *((**int32)(unsafe.Add(unsafe.Pointer((**int32)(unsafe.Pointer(v4))), unsafe.Sizeof((*int32)(nil))*9))) == v3 {
					sub_4523D0((*uint32)(unsafe.Pointer(v4)))
					sub_451FE0(int32(uintptr(unsafe.Pointer(v4))))
					v1 = 1
				}
				v4 = v5
				if unsafe.Pointer(v5) == memmap.PtrOff(6112660, 840612) {
					break
				}
			}
		}
		result = (*int32)(unsafe.Pointer(uintptr(v1)))
	}
	return result
}
func sub_452190(a1 int32) **uint32 {
	return sub_425920((**uint32)(unsafe.Pointer(uintptr(a1 + 112))))
}
func sub_4521A0(a1 int32) *int32 {
	var (
		v1 int32
		v2 *uint8
		v3 int32
		v4 *int32
		v5 *int32
	)
	v1 = 0
	v2 = (*uint8)(memmap.PtrOff(6112660, 839892))
	if a1 > 0 {
		for {
			v3 = 0
			v4 = (*int32)(unsafe.Pointer(v2))
			for {
				v5 = (*int32)(unsafe.Pointer(nox_common_list_getFirstSafe_425890((*nox_list_item_t)(unsafe.Pointer(v4)))))
				if v5 != nil {
					return (*int32)(unsafe.Add(unsafe.Pointer(v5), -int(unsafe.Sizeof(int32(0))*28)))
				}
				v3++
				v4 = (*int32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(int32(0))*3))
				if v3 >= 10 {
					break
				}
			}
			v1++
			v2 = (*uint8)(unsafe.Add(unsafe.Pointer(v2), 120))
			if v1 < a1 {
				continue
			}
			break
		}
	}
	return nil
}
func sub_4521F0() int32 {
	var (
		result int32
		v1     *uint8
		v2     *uint8
	)
	result = int32(dword_5d4594_1045432)
	if dword_5d4594_1045432 != 0 {
		v1 = *(**uint8)(memmap.PtrOff(6112660, 840612))
		if unsafe.Pointer(*(**uint8)(memmap.PtrOff(6112660, 840612))) != memmap.PtrOff(6112660, 840612) {
			for {
				v2 = *(**uint8)(unsafe.Pointer(v1))
				sub_4523D0((*uint32)(unsafe.Pointer(v1)))
				result = sub_451FE0(int32(uintptr(unsafe.Pointer(v1))))
				v1 = v2
				if unsafe.Pointer(v2) == memmap.PtrOff(6112660, 840612) {
					break
				}
			}
		}
	}
	return result
}
func sub_452230() *****int32 {
	var (
		result *****int32
		v1     ****int32
	)
	result = *(******int32)(unsafe.Pointer(&dword_5d4594_1045432))
	if dword_5d4594_1045432 != 0 {
		result = *(******int32)(memmap.PtrOff(6112660, 840612))
		if unsafe.Pointer(*(**uint8)(memmap.PtrOff(6112660, 840612))) != memmap.PtrOff(6112660, 840612) {
			for {
				v1 = *result
				if int32(uint8(uintptr(unsafe.Pointer(*(*****int32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof((****int32)(nil))*6))))))&1 != 0 {
					sub_451FE0(int32(uintptr(unsafe.Pointer(result))))
				}
				result = (*****int32)(unsafe.Pointer(v1))
				if v1 == (****int32)(memmap.PtrOff(6112660, 840612)) {
					break
				}
			}
		}
	}
	return result
}
func nox_xxx_draw_452270(a1 int32) *byte {
	var result *byte
	if dword_5d4594_1045432 != 0 && a1 >= 0 && a1 < 1023 {
		result = (*byte)(memmap.PtrOff(6112660, uint32(a1*200)+0xCD3B4))
	} else {
		result = nil
	}
	return result
}
func sub_4522A0(a1 int32) int32 {
	var result int32
	if dword_5d4594_1045432 != 0 {
		result = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 64))))
	} else {
		result = 0
	}
	return result
}
func sub_4522C0(a1 int32) int32 {
	var result int32
	if dword_5d4594_1045432 != 0 {
		result = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 8))))
	} else {
		result = 0
	}
	return result
}
func sub_4522E0(a1 int32) *byte {
	var result *byte
	if dword_5d4594_1045432 != 0 {
		result = *(**byte)(unsafe.Pointer(uintptr(a1 + 84)))
	} else {
		result = (*byte)(memmap.PtrOff(0x587000, 127128))
	}
	return result
}
func nox_xxx_draw_452300(a1 *uint32) *uint32 {
	var v1 *uint32
	if dword_5d4594_1045432 == 0 {
		return nil
	}
	if dword_587000_126996 == 0 {
		return nil
	}
	if *a1 == 0 {
		return nil
	}
	v1 = sub_4BD2E0(*(***uint32)(unsafe.Pointer(&dword_5d4594_1045436)))
	if v1 == nil {
		sub_452230()
		v1 = sub_4BD2E0(*(***uint32)(unsafe.Pointer(&dword_5d4594_1045436)))
		if v1 == nil {
			return nil
		}
	}
	libc.MemSet(unsafe.Pointer(v1), 0, 576)
	*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*9)) = uint32(uintptr(unsafe.Pointer(a1)))
	sub_425770(unsafe.Pointer(v1))
	*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*7)) = 0
	*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*75)) = 0
	*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*142)) = 0
	*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*108)) = 0
	*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*42)) = 0
	sub_4864A0((*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*46)))
	nox_common_list_append_4258E0((*nox_list_item_t)(unsafe.Pointer(uintptr(int32(uintptr(memmap.PtrOff(6112660, 840612)))))), (*nox_list_item_t)(unsafe.Pointer(v1)))
	*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*70)) = func() uint32 {
		p := memmap.PtrUint32(0x587000, 127000)
		x := *p
		*p++
		return x
	}()
	return v1
}
func sub_4523D0(a1 *uint32) int32 {
	var result int32 = 0
	if (*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*6)) & 1) == 0 {
		sub_452410(int32(uintptr(unsafe.Pointer(a1))))
		sub_451F90(int32(uintptr(unsafe.Pointer(a1))))
		*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*7)) = 4
		*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*70)) = 0
		result = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*6)))
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&result))), 0)) = uint8(int8(result | 1))
		*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*6)) = uint32(result)
	}
	return result
}
func sub_452410(a1 int32) int32 {
	var result int32
	result = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 176))))
	if result != 0 && uint32(a1) == *(*uint32)(unsafe.Pointer(uintptr(result + 152))) {
		if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 24))))&2 != 0 {
			sub_4BDA80(int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 176)))))
		}
		sub_4BDB30(int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 176)))))
		*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 176))) + 152))) = 0
		*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 176))) + 148))) = 0
		result = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 176))))
		*(*uint32)(unsafe.Pointer(uintptr(result + 140))) = 0
		*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 176))) + 144))) = 0
		*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 176))) + 112))) = 0
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 176))) = 0
	}
	return result
}
func sub_452490(a1 *uint32) int32 {
	var (
		v1 int32
		v3 int32
		v4 int32
	)
	v1 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*44)))
	if a1 != *(**uint32)(unsafe.Pointer(uintptr(v1 + 152))) {
		return 0
	}
	v3 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*74)))
	sub_4BDB90((*uint32)(unsafe.Pointer(uintptr(v1))), (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*74))))))
	*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*7)) = 3
	v4 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*6)))
	*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v4))), 0)) = uint8(int8(v4 | 2))
	*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*6)) = uint32(v4)
	*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*74)) = 0
	if sub_4BDB40(int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*44)))) == 0 {
		return 1
	}
	*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*7)) = 1
	*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*74)) = uint32(v3)
	*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*6)) &= 0xFFFFFFFD
	return 0
}
func sub_452510(a3 int32) {
	var (
		v1 int32
		v2 int32
	)
	if dword_587000_126996 == 0 {
		*(*uint32)(unsafe.Pointer(uintptr(a3 + 28))) = 4
	}
	for {
		v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(a3 + 28))))
		if v1 == 0 {
			break
		}
		v2 = v1 - 2
		if v2 != 0 {
			var v3 uint32 = uint32(v2 - 2)
			if v3 == 0 {
				sub_4523D0((*uint32)(unsafe.Pointer(uintptr(a3))))
			}
			return
		}
		if uint64(platformTicks()) <= *(*uint64)(unsafe.Pointer(uintptr(a3 + 288))) {
			return
		}
		*(*uint32)(unsafe.Pointer(uintptr(a3 + 28))) = *(*uint32)(unsafe.Pointer(uintptr(a3 + 32)))
	}
	if sub_452580((*uint32)(unsafe.Pointer(uintptr(a3)))) == 0 {
		sub_4523D0((*uint32)(unsafe.Pointer(uintptr(a3))))
	}
}
func sub_452690(a3 int32, a4 int64, a5 int32) int64 {
	var result int64
	*(*uint32)(unsafe.Pointer(uintptr(a3 + 32))) = uint32(a5)
	result = a4 + int64(platformTicks())
	*(*uint64)(unsafe.Pointer(uintptr(a3 + 288))) = uint64(result)
	*(*uint32)(unsafe.Pointer(uintptr(a3 + 28))) = 2
	return result
}
func sub_4526D0(a1 int32) int32 {
	*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 152))) + 28))) = 4
	return 0
}
func sub_4526F0(a1 int32) int32 {
	var (
		v1 *uint32
		v2 int32
	)
	v1 = *(**uint32)(unsafe.Pointer(uintptr(a1 + 152)))
	*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*6)) &= 0xFFFFFFFD
	v2 = 4
	if *(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*7)) != 4 {
		if *(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*74)) != 0 || *(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*142)) != 0 {
			v2 = 1
		} else {
			*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*71)) = 0
		}
		if *(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*71)) != 0 {
			sub_452690(int32(uintptr(unsafe.Pointer(v1))), int64(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*71))), v2)
			*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*71)) = 0
			return 0
		}
		*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*7)) = uint32(v2)
	}
	return 0
}
func sub_452810(a1 int32, a2 int8) *int32 {
	var (
		v2 *int32
		v3 *int32
	)
	v2 = nil
	if dword_5d4594_1045428 != 0 {
		v3 = sub_487810(*(*int32)(unsafe.Pointer(&dword_5d4594_1045428)), 1)
		v2 = v3
		if v3 != nil {
			if *(*int32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(int32(0))*31))&21 != 0 && *(*int32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(int32(0))*30)) > a1 {
				return nil
			}
			sub_4BDA80(int32(uintptr(unsafe.Pointer(v3))))
			*(*int32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(int32(0))*29)) = int32(uintptr(dword_587000_127004))
			*(*int32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(int32(0))*30)) = a1
			if int32(a2)&1 != 0 {
				*(*int32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(int32(0))*32)) = -1
			} else {
				*(*int32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(int32(0))*32)) = 0
			}
			sub_486320((*uint32)(unsafe.Pointer((*int32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(int32(0))*4)))), 0x4000)
		}
	}
	return v2
}
func nox_thing_read_AVNT_452890(a1p *nox_memfile, a2 unsafe.Pointer) int32 {
	var (
		a1     int32 = int32(uintptr(unsafe.Pointer(a1p)))
		v2     int32
		v3     *uint8
		v4     int32
		v5     *byte
		v6     *byte
		v7     int8
		v8     *byte
		v9     int32
		v10    *uint8
		v11    int16
		v12    int32
		v13    *int32
		v14    int32
		v15    int16
		v16    int16
		v17    int32
		v18    *byte
		v19    int8
		result int32
		v21    int32
		v22    *uint16
		v23    uint8
		v24    uint8
		v25    uint8
		v26    uint8
		v27    uint8
		v28    uint8
		v29    uint8
	)
	v2 = a1
	v3 = *(**uint8)(unsafe.Pointer(uintptr(a1 + 8)))
	v21 = a1
	v28 = *v3
	*(*uint32)(unsafe.Pointer(uintptr(v21 + 8))) = uint32(uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v3), 1)))))
	nox_memfile_read(a2, 1, int32(v28), (*nox_memfile)(unsafe.Pointer(uintptr(v21))))
	*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(a2)), v28))) = 0
	v4 = nox_xxx_utilFindSound_40AF50((*byte)(a2))
	if v4 != 0 && (func() *byte {
		v5 = nox_xxx_draw_452270(v4)
		return v5
	}()) != nil {
		for {
			v6 = *(**byte)(unsafe.Pointer(uintptr(v2 + 8)))
			v7 = int8(*v6)
			v8 = (*byte)(unsafe.Add(unsafe.Pointer(v6), 1))
			*(*uint32)(unsafe.Pointer(uintptr(v2 + 8))) = uint32(uintptr(unsafe.Pointer(v8)))
			switch v7 {
			case 0:
				*(*uint32)(unsafe.Pointer(v5)) = 1
				result = bool2int(int32(v7) == 0)
			case 1:
				v23 = uint8(*v8)
				*(*uint32)(unsafe.Pointer(uintptr(v2 + 8))) = uint32(uintptr(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v8), 1)))))
				*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v5))), unsafe.Sizeof(uint32(0))*12))) = uint32(v23)
				continue
			case 2:
				v26 = uint8(*v8)
				*(*uint32)(unsafe.Pointer(uintptr(v2 + 8))) = uint32(uintptr(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v8), 1)))))
				*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v5))), unsafe.Sizeof(uint32(0))*1))) = uint32(v26)
				continue
			case 3:
				v27 = uint8(*v8)
				*(*uint32)(unsafe.Pointer(uintptr(v2 + 8))) = uint32(uintptr(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v8), 1)))))
				sub_4862E0(int32(uintptr(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v5), 16))))), int32(v27)*163)
				continue
			case 4:
				v24 = uint8(*v8)
				*(*uint32)(unsafe.Pointer(uintptr(v2 + 8))) = uint32(uintptr(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v8), 1)))))
				*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v5))), unsafe.Sizeof(uint32(0))*14))) = uint32(v24)
				continue
			case 5:
				v25 = uint8(*v8)
				*(*uint32)(unsafe.Pointer(uintptr(v2 + 8))) = uint32(uintptr(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v8), 1)))))
				*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v5))), unsafe.Sizeof(uint32(0))*15))) = uint32(v25)
				continue
			case 6:
				v17 = int32(*v8)
				*(*uint32)(unsafe.Pointer(uintptr(v2 + 8))) = uint32(uintptr(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v8), 1)))))
				*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v5))), unsafe.Sizeof(uint32(0))*19))) = uint32(v17)
				v18 = *(**byte)(unsafe.Pointer(uintptr(v2 + 8)))
				v19 = int8(*v18)
				*(*uint32)(unsafe.Pointer(uintptr(v2 + 8))) = uint32(uintptr(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v18), 1)))))
				*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v5))), unsafe.Sizeof(uint32(0))*20))) = uint32(v19)
				continue
			case 7:
				v9 = 0
				v22 = (*uint16)(unsafe.Add(unsafe.Pointer(v5), 128))
				for {
					v10 = *(**uint8)(unsafe.Pointer(uintptr(v2 + 8)))
					v29 = *v10
					*(*uint32)(unsafe.Pointer(uintptr(v2 + 8))) = uint32(uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v10), 1)))))
					if int32(v29) == 0 {
						break
					}
					nox_memfile_read(a2, 1, int32(v29), (*nox_memfile)(unsafe.Pointer(uintptr(v2))))
					*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(a2)), v29))) = 0
					if v9 < 32 {
						v11 = int16(uint16(sub_486A10(*(*int32)(unsafe.Pointer(&dword_5d4594_1045420)), a2)))
						*v22 = uint16(v11)
						if int32(v11) != -1 {
							v9++
							v22 = (*uint16)(unsafe.Add(unsafe.Pointer(v22), unsafe.Sizeof(uint16(0))*1))
						}
					}
				}
				*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v5))), unsafe.Sizeof(uint32(0))*48))) = uint32(v9)
				continue
			case 8:
				v12 = int32(*(*uint32)(unsafe.Pointer(v8)))
				*(*uint32)(unsafe.Pointer(uintptr(v2 + 8))) = uint32(uintptr(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v8), 4)))))
				*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v5))), unsafe.Sizeof(uint32(0))*17))) = uint32(v12)
				v13 = *(**int32)(unsafe.Pointer(uintptr(v2 + 8)))
				v14 = *v13
				*(*uint32)(unsafe.Pointer(uintptr(v2 + 8))) = uint32(uintptr(unsafe.Pointer((*int32)(unsafe.Add(unsafe.Pointer(v13), unsafe.Sizeof(int32(0))*1)))))
				*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v5))), unsafe.Sizeof(uint32(0))*18))) = uint32(v14)
				continue
			case 9:
				v15 = int16(*(*uint16)(unsafe.Pointer(v8)))
				*(*uint32)(unsafe.Pointer(uintptr(v2 + 8))) = uint32(uintptr(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v8), 2)))))
				if int32(v15) > 0 {
					*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v5))), unsafe.Sizeof(uint32(0))*16))) = uint32(int32(v15) * 15)
				}
				continue
			case 10:
				v16 = int16(*(*uint16)(unsafe.Pointer(v8)))
				*(*uint32)(unsafe.Pointer(uintptr(v2 + 8))) = uint32(uintptr(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v8), 2)))))
				*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v5))), unsafe.Sizeof(uint32(0))*2))) = uint32(v16)
				continue
			default:
				result = 0
			}
			break
		}
	} else {
		nox_thing_skip_AVNT_inner_452B30((*nox_memfile)(unsafe.Pointer(uintptr(v2))))
		result = 1
	}
	return result
}
func nox_thing_skip_AVNT_452B00(f *nox_memfile) int32 {
	var sz int32 = int32(f.ReadU8())
	f.Skip(sz)
	return nox_thing_skip_AVNT_inner_452B30(f)
}
func nox_thing_skip_AVNT_inner_452B30(f *nox_memfile) int32 {
	var (
		v3     int8
		v5     uint8
		result int32
		v8     int8
	)
	for {
		v3 = f.ReadI8()
		v8 = v3
		switch v8 {
		case 0:
			result = bool2int(int32(v3) == 0)
		case 1:
			fallthrough
		case 2:
			fallthrough
		case 3:
			fallthrough
		case 4:
			fallthrough
		case 5:
			f.Skip(1)
			continue
		case 6:
			fallthrough
		case 9:
			fallthrough
		case 10:
			f.Skip(2)
			continue
		case 7:
			for {
				v5 = f.ReadU8()
				if int32(v5) == 0 {
					break
				}
				f.Skip(int32(v5))
			}
			continue
		case 8:
			f.Skip(8)
			continue
		default:
			result = 0
		}
		break
	}
	return result
}
func sub_452BD0(a1 int32, a2 *byte) int32 {
	var (
		v2     int32
		v3     *byte
		v4     *byte
		v5     int32
		v6     int32
		v7     *byte
		v8     *int16
		v9     int16
		v10    *uint8
		v11    *int16
		v12    int16
		v13    *byte
		v14    int8
		v15    *byte
		v16    int8
		v17    *byte
		v18    int32
		v19    *byte
		v20    int8
		result int32
		v22    *byte
		v23    int8
		v24    int32
		v25    *byte
		v26    int16
		v27    *byte
		v28    int8
		v29    int32
		v30    uint8
		v31    *uint16
		v32    *byte
	)
	v2 = a1
	v3 = a2
	v4 = *(**byte)(unsafe.Pointer(uintptr(a1 + 8)))
	v5 = int32(*v4)
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 8))) = uint32(uintptr(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v4), 1)))))
	nox_memfile_read(unsafe.Pointer(a2), 1, v5, (*nox_memfile)(unsafe.Pointer(uintptr(a1))))
	*(*byte)(unsafe.Add(unsafe.Pointer(a2), v5)) = 0
	v6 = nox_xxx_utilFindSound_40AF50(a2)
	if v6 != 0 && (func() *byte {
		v7 = nox_xxx_draw_452270(v6)
		return v7
	}()) != nil {
		v8 = *(**int16)(unsafe.Pointer(uintptr(a1 + 8)))
		v9 = *v8
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 8))) = uint32(uintptr(unsafe.Pointer((*int16)(unsafe.Add(unsafe.Pointer(v8), unsafe.Sizeof(int16(0))*1)))))
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v7))), unsafe.Sizeof(uint32(0))*1))) = 2
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v7))), unsafe.Sizeof(uint32(0))*2))) = uint32(v9)
		v10 = *(**uint8)(unsafe.Pointer(uintptr(a1 + 8)))
		v30 = *v10
		*(*uint32)(unsafe.Pointer(uintptr(v2 + 8))) = uint32(uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v10), 1)))))
		sub_4862E0(int32(uintptr(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v7), 16))))), int32(v30)*163)
		v11 = *(**int16)(unsafe.Pointer(uintptr(v2 + 8)))
		v12 = *v11
		*(*uint32)(unsafe.Pointer(uintptr(v2 + 8))) = uint32(uintptr(unsafe.Pointer((*int16)(unsafe.Add(unsafe.Pointer(v11), unsafe.Sizeof(int16(0))*1)))))
		if int32(v12) > 0 {
			*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v7))), unsafe.Sizeof(uint32(0))*16))) = uint32(int32(v12) * 15)
		}
		v13 = *(**byte)(unsafe.Pointer(uintptr(v2 + 8)))
		v14 = int8(*v13)
		*(*uint32)(unsafe.Pointer(uintptr(v2 + 8))) = uint32(uintptr(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v13), 1)))))
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v7))), unsafe.Sizeof(uint32(0))*14))) = uint32(v14)
		v15 = *(**byte)(unsafe.Pointer(uintptr(v2 + 8)))
		v16 = int8(*v15)
		*(*uint32)(unsafe.Pointer(uintptr(v2 + 8))) = uint32(uintptr(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v15), 1)))))
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v7))), unsafe.Sizeof(uint32(0))*19))) = uint32(v16)
		v17 = *(**byte)(unsafe.Pointer(uintptr(v2 + 8)))
		v18 = int32(*v17)
		*(*uint32)(unsafe.Pointer(uintptr(v2 + 8))) = uint32(uintptr(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v17), 1)))))
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v7))), unsafe.Sizeof(uint32(0))*20))) = uint32(v18)
		v19 = *(**byte)(unsafe.Pointer(uintptr(v2 + 8)))
		v20 = int8(*v19)
		*(*uint32)(unsafe.Pointer(uintptr(v2 + 8))) = uint32(uintptr(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v19), 1)))))
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v7))), unsafe.Sizeof(uint32(0))*12))) = uint32(v20)
		if int32(v20) < 3 {
			v32 = nil
			v31 = (*uint16)(unsafe.Add(unsafe.Pointer(v7), 128))
			for {
				v22 = *(**byte)(unsafe.Pointer(uintptr(v2 + 8)))
				v23 = int8(*v22)
				*(*uint32)(unsafe.Pointer(uintptr(v2 + 8))) = uint32(uintptr(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v22), 1)))))
				if int32(v23) == 0 {
					break
				}
				v24 = int32(v23)
				nox_memfile_read(unsafe.Pointer(v3), 1, int32(v23), (*nox_memfile)(unsafe.Pointer(uintptr(v2))))
				*(*byte)(unsafe.Add(unsafe.Pointer(v3), v24)) = 0
				v25 = libc.StrRChr(v3, 46)
				if v25 != nil {
					*v25 = 0
				}
				v26 = int16(uint16(sub_486A10(*(*int32)(unsafe.Pointer(&dword_5d4594_1045420)), unsafe.Pointer(v3))))
				*v31 = uint16(v26)
				if int32(v26) != -1 {
					v32 = (*byte)(unsafe.Add(unsafe.Pointer(v32), 1))
					v31 = (*uint16)(unsafe.Add(unsafe.Pointer(v31), unsafe.Sizeof(uint16(0))*1))
				}
			}
			*(*uint32)(unsafe.Pointer(v7)) = 1
			*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v7))), unsafe.Sizeof(uint32(0))*48))) = uint32(uintptr(unsafe.Pointer(v32)))
			result = 1
		} else {
			result = 0
		}
	} else {
		for *(*uint32)(unsafe.Pointer(uintptr(a1 + 8))) += 9; ; *(*uint32)(unsafe.Pointer(uintptr(a1 + 8))) = uint32(v29 + int32(v28)) {
			v27 = *(**byte)(unsafe.Pointer(uintptr(a1 + 8)))
			v28 = int8(*v27)
			v29 = int32(uintptr(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v27), 1)))))
			*(*uint32)(unsafe.Pointer(uintptr(a1 + 8))) = uint32(v29)
			if int32(v28) == 0 {
				break
			}
		}
		result = 1
	}
	return result
}
func nox_xxx_clientPlaySoundSpecial_452D80(a1 int32, a2 int32) {
	var (
		result *uint32
		v3     *uint32
	)
	result = (*uint32)(unsafe.Pointer(nox_xxx_draw_452270(a1)))
	if result == nil {
		return
	}
	result = nox_xxx_draw_452300(result)
	v3 = result
	if result == nil {
		return
	}
	sub_452EE0(int32(uintptr(unsafe.Pointer(result))), a2)
	sub_452510(int32(uintptr(unsafe.Pointer(v3))))
}
func sub_452DC0(a1 int32, a2 int32, a3 int32) {
	var (
		result *uint32
		v4     *uint32
	)
	result = (*uint32)(unsafe.Pointer(nox_xxx_draw_452270(a1)))
	if result == nil {
		return
	}
	result = nox_xxx_draw_452300(result)
	v4 = result
	if result == nil {
		return
	}
	sub_452EE0(int32(uintptr(unsafe.Pointer(result))), a2)
	sub_452F80(int32(uintptr(unsafe.Pointer(v4))), a3)
	sub_452510(int32(uintptr(unsafe.Pointer(v4))))
}
func sub_452E10(a1 int32, a2 int32, a3 int32) {
	var (
		result *uint32
		v4     *uint32
	)
	result = (*uint32)(unsafe.Pointer(nox_xxx_draw_452270(a1)))
	if result == nil {
		return
	}
	result = nox_xxx_draw_452300(result)
	v4 = result
	if result == nil {
		return
	}
	sub_452EE0(int32(uintptr(unsafe.Pointer(result))), a2)
	sub_452F80(int32(uintptr(unsafe.Pointer(v4))), a3)
	*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*75)) = 2
	sub_452510(int32(uintptr(unsafe.Pointer(v4))))
}
func sub_452E60(a1 *uint32) *uint32 {
	var result *uint32
	result = a1
	*a1 = 0
	*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*1)) = 0
	return result
}
func sub_452E80(a1 *uint32) *uint32 {
	var result *uint32
	result = a1
	*a1 = 0
	return result
}
func sub_452E90(a1 *uint32, a2 int32) int32 {
	var result int32
	result = a2
	*a1 = uint32(a2)
	if a2 != 0 {
		*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*1)) = *(*uint32)(unsafe.Pointer(uintptr(a2 + 280)))
		result = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 36))))
		*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*2)) = uint32(result)
	}
	return result
}
func sub_452EB0(a1 *int32) int32 {
	var result int32
	result = *a1
	if *a1 != 0 && (uint32(*(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*2))) != *(*uint32)(unsafe.Pointer(uintptr(result + 36))) || uint32(*(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*1))) != *(*uint32)(unsafe.Pointer(uintptr(result + 280)))) {
		result = 0
		*a1 = 0
	}
	return result
}
func sub_452EE0(a1 int32, a2 int32) int32 {
	var v2 int32
	v2 = int32(sub_452F10(a1, a2))
	sub_486320((*uint32)(unsafe.Pointer(uintptr(a1+184))), v2)
	return sub_4863B0((*uint32)(unsafe.Pointer(uintptr(a1 + 184))))
}
func sub_452F10(a1 int32, a2 int32) uint32 {
	var v2 int32
	v2 = a2
	if a2 <= 100 {
		if a2 < 0 {
			v2 = 0
		}
	} else {
		v2 = 100
	}
	return (uint32(v2*163) * (*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 36))) + 20))) >> 16)) >> 14
}
func sub_452F50(a1 int32, a2 int32) int32 {
	var v2 int32
	v2 = int32(sub_452F10(a1, a2))
	return sub_486350(unsafe.Pointer(uintptr(a1+184)), v2)
}
func sub_452F80(a1 int32, a2 int32) *uint32 {
	var v2 int32
	v2 = sub_452FA0(a2)
	return sub_486320((*uint32)(unsafe.Pointer(uintptr(a1+248))), v2)
}
func sub_452FA0(a1 int32) int32 {
	var v1 int32
	v1 = a1
	if a1 <= 50 {
		if a1 < -50 {
			v1 = -50
		}
	} else {
		v1 = 50
	}
	return (v1*8192)/50 + 8192
}
func sub_452FE0(a1 int32, a2 int32) int32 {
	var v2 int32
	v2 = sub_452FA0(a2)
	return sub_486350(unsafe.Pointer(uintptr(a1+248)), v2)
}
func sub_453000(a1 *uint32, a2 uint32, a3 int32) int32 {
	sub_486380((*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*46)), a2, a3, 0x4000)
	sub_486380((*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*54)), a2, a3, 0x4000)
	return sub_486380((*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*62)), a2, a3, 0x4000)
}
func sub_453050() {
	dword_587000_126996 = 0
}
func nox_xxx____setargv_9_453060() {
	dword_587000_126996 = 1
}
func sub_453070() int32 {
	return int32(dword_587000_126996)
}
func sub_453080(a1 int8) int32 {
	var result int32
	if dword_5d4594_1045460 != 0 {
		result = sub_453690(1 << int32(a1))
	} else {
		result = sub_453660(1 << int32(a1))
	}
	return result
}
func sub_4532E0() *uint32 {
	var (
		v0     int32
		v1     int32
		v2     int32
		result *uint32
	)
	v0 = nox_xxx_guiFontHeightMB_43F320(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1045464 + 236)))))) + 1
	sub_46AB20(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045464)), int32(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1045464 + 8)))), v0*15+2)
	v1 = 1520
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1045464 + 20))) + uint32(v0) + 2)
	for {
		result = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045468)))).ChildByID(v1)))
		*(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*5)) = uint32(v2)
		v2 += v0
		v1++
		if v1 >= 1534 {
			break
		}
	}
	return result
}
func sub_453350(a1 int32, a2 int32) int32 {
	var (
		result int32
		xLeft  int32
		yTop   int32
	)
	nox_client_wndGetPosition_46AA60((*nox_window)(unsafe.Pointer(uintptr(a1))), (*uint32)(unsafe.Pointer(&xLeft)), (*uint32)(unsafe.Pointer(&yTop)))
	if int32(int8(*(*uint8)(unsafe.Pointer(uintptr(a1 + 4))))) >= 0 {
		if *(*uint32)(unsafe.Pointer(uintptr(a2 + 20))) != 0x80000000 {
			nox_client_drawRectFilledAlpha_49CF10(xLeft, yTop, int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 8)))), int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 12)))))
		}
		result = 1
	} else {
		nox_client_drawImageAt_47D2C0((*nox_video_bag_image_t)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a2 + 24)))))), xLeft, yTop)
		result = 1
	}
	return result
}
func sub_4533D0(a1 int32, a2 int32, a3 int32, a4 int32) int32 {
	var (
		v3  int32
		v5  int32
		v6  int32
		v7  int32
		v8  int32
		v9  int32
		v10 int32
		v11 *libc.WChar
	)
	if a2 == 0x4000 {
		if unsafe.Pointer((*uint32)(unsafe.Pointer(uintptr(a3)))) == unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045468)))).ChildByID(1513)) || unsafe.Pointer((*uint32)(unsafe.Pointer(uintptr(a3)))) == unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045468)))).ChildByID(1514)) {
			(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045464))))).Func94(asWindowEvent(0x4000, a3, 0))
		LABEL_35:
			sub_453750()
		}
		return 0
	}
	if a2 != 0x4007 {
		return 0
	}
	v3 = (*nox_window)(unsafe.Pointer(uintptr(a3))).ID()
	switch v3 {
	case 1513:
		fallthrough
	case 1514:
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045464))))).Func94(asWindowEvent(0x4000, a3, 0))
		goto LABEL_35
	case 1515:
		if dword_5d4594_1045460 != 0 {
			*memmap.PtrUint32(6112660, 1045456) = math.MaxUint32
		} else {
			*memmap.PtrUint32(6112660, 1045452) = math.MaxUint32
		}
		goto LABEL_6
	case 1516:
		if dword_5d4594_1045460 != 0 {
			*memmap.PtrUint32(6112660, 1045456) = 0
		} else {
			*memmap.PtrUint32(6112660, 1045452) = 0
		}
	LABEL_6:
		sub_453750()
		goto LABEL_7
	case 1520:
		fallthrough
	case 1521:
		fallthrough
	case 1522:
		fallthrough
	case 1523:
		fallthrough
	case 1524:
		fallthrough
	case 1525:
		fallthrough
	case 1526:
		fallthrough
	case 1527:
		fallthrough
	case 1528:
		fallthrough
	case 1529:
		fallthrough
	case 1530:
		fallthrough
	case 1531:
		fallthrough
	case 1532:
		fallthrough
	case 1533:
		v5 = sub_4A4800(int32(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1045464 + 32)))))
		v11 = (*libc.WChar)(unsafe.Pointer(uintptr((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045464))))).Func94(asWindowEvent(0x4016, v5+v3-1520, 0)))))
		if dword_5d4594_1045460 != 0 {
			v6 = sub_415DA0(v11)
		} else {
			v6 = sub_415960(v11)
		}
		if int32(*(*uint8)(unsafe.Pointer(uintptr(a3 + 36))))&4 != 0 {
			if dword_5d4594_1045460 != 0 {
				sub_453640((*uint32)(memmap.PtrOff(6112660, 1045456)), v6, 0)
			} else {
				v7 = 0
				if v6 > 0 {
					for {
						v8 = v6 >> 8
						if v6>>8 > 0 {
							v6 >>= 8
						}
						v7++
						if v8 <= 0 {
							break
						}
					}
				}
				sub_453620((*uint8)(memmap.PtrOff(6112660, uint32(v7)+0xFF3CB)), int8(v6), 0)
			}
		} else if dword_5d4594_1045460 != 0 {
			sub_453640((*uint32)(memmap.PtrOff(6112660, 1045456)), v6, 1)
		} else {
			v9 = 0
			if v6 > 0 {
				for {
					v10 = v6 >> 8
					if v6>>8 > 0 {
						v6 >>= 8
					}
					v9++
					if v10 <= 0 {
						break
					}
				}
			}
			sub_453620((*uint8)(memmap.PtrOff(6112660, uint32(v9)+0xFF3CB)), int8(v6), 1)
		}
	LABEL_7:
		sub_459D50(1)
	default:
	}
	clientPlaySoundSpecial(766, 100)
	return 0
}
func sub_4535E0(a1 *int32) *int32 {
	var result *int32
	result = a1
	*memmap.PtrUint32(6112660, 1045452) = uint32(*a1)
	return result
}
func sub_4535F0(a1 int32) int32 {
	var result int32
	result = a1
	*memmap.PtrUint32(6112660, 1045456) = uint32(a1)
	return result
}
func sub_453600() *byte {
	return (*byte)(memmap.PtrOff(6112660, 1045452))
}
func sub_453610() int32 {
	return int32(*memmap.PtrUint32(6112660, 1045456))
}
func sub_453620(a1 *uint8, a2 int8, a3 int32) *uint8 {
	var result *uint8
	result = a1
	if a3 != 0 {
		*a1 |= uint8(a2)
	} else {
		*a1 &= uint8(int8(int32(^a2)))
	}
	return result
}
func sub_453640(a1 *uint32, a2 int32, a3 int32) *uint32 {
	var result *uint32
	result = a1
	if a3 != 0 {
		*a1 |= uint32(a2)
	} else {
		*a1 &= uint32(^a2)
	}
	return result
}
func sub_453660(a1 int32) int32 {
	var (
		v1 int32
		v2 int32
		v3 int32
	)
	v1 = a1
	v2 = 0
	if a1 > 0 {
		for {
			v3 = v1 >> 8
			if v1>>8 > 0 {
				v1 >>= 8
			}
			v2++
			if v3 <= 0 {
				break
			}
		}
	}
	return bool2int((int32(uint8(int8(v1))) & int32(*memmap.PtrUint8(6112660, uint32(v2)+0xFF3CB))) != 0)
}
func sub_453690(a1 int32) int32 {
	return bool2int((uint32(a1) & *memmap.PtrUint32(6112660, 1045456)) != 0)
}
func sub_4536B0(a1 *uint32) int32 {
	var (
		v1     int32
		v2     int32
		v3     *uint32
		v4     int32
		result int32
	)
	v1 = 1
	*a1 = math.MaxUint32
	v2 = 1
	v3 = a1
	v4 = 27
	for {
		result = sub_415840((*byte)(unsafe.Pointer(uintptr(v2))))
		if result != 0 {
			result = nox_xxx_getUnitDefDd10_4E3BA0(result)
			if result == 0 {
				result = int32(^uint8(int8(v1)))
				*(*uint8)(unsafe.Pointer(v3)) &= ^uint8(int8(v1))
			}
		}
		if v1 == 128 {
			v1 = 1
			v3 = (*uint32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(v3))), 1))))
		} else {
			v1 *= 2
		}
		v2 *= 2
		v4--
		if v4 == 0 {
			break
		}
	}
	return result
}
func sub_453710() int32 {
	var (
		v0 int32
		v1 int32
		v2 int32
		v3 int32
	)
	v0 = -1
	v1 = 1
	v2 = 26
	for {
		v3 = sub_415CD0((*byte)(unsafe.Pointer(uintptr(v1))))
		if v3 != 0 && nox_xxx_getUnitDefDd10_4E3BA0(v3) == 0 {
			v0 &= ^v1
		}
		v1 *= 2
		v2--
		if v2 == 0 {
			break
		}
	}
	return v0
}
func sub_453750() int8 {
	var (
		v0 int32
		i  int32
		j  *uint32
		v3 bool
		v4 int32
	)
	v0 = sub_4A4800(int32(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1045464 + 32)))))
	for i = 1520; i <= 1533; i++ {
		for j = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045468)))).ChildByID(i))); (1<<v0)&51 != 0; v0++ {
		}
		if v0 >= *memmap.PtrInt32(6112660, dword_5d4594_1045460*4+0xFF3E0) {
			*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v4))), 0)) = uint8(int8((*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(j)))))).Hide()))
		} else {
			(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(j)))))).Show()
			v3 = sub_453080(int8(v0)) == 0
			v4 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(j), unsafe.Sizeof(uint32(0))*9)))
			if v3 {
				*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v4))), 0)) = uint8(int8(v4 & 251))
			} else {
				*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v4))), 0)) = uint8(int8(v4 | 4))
			}
			*(*uint32)(unsafe.Add(unsafe.Pointer(j), unsafe.Sizeof(uint32(0))*9)) = uint32(v4)
		}
		v0++
	}
	return int8(v4)
}
func sub_4537F0() int32 {
	var (
		v0     int32
		v1     int32
		v2     int32
		v3     int32
		v4     int32
		v5     int32
		result int32
		v7     int32
	)
	v0 = 1
	v1 = 26
	for {
		v2 = sub_415CD0((*byte)(unsafe.Pointer(uintptr(v0))))
		if v2 != 0 {
			v3 = int32(uintptr(unsafe.Pointer(nox_xxx_objectTypeByInd_4E3B70(v2))))
			sub_4E3BF0(v3)
		}
		v0 *= 2
		v1--
		if v1 == 0 {
			break
		}
	}
	v4 = 1
	v5 = 27
	for {
		result = sub_415840((*byte)(unsafe.Pointer(uintptr(v4))))
		if result != 0 {
			v7 = int32(uintptr(unsafe.Pointer(nox_xxx_objectTypeByInd_4E3B70(result))))
			result = sub_4E3BF0(v7)
		}
		v4 *= 2
		v5--
		if v5 == 0 {
			break
		}
	}
	return result
}
func sub_453B00() *uint32 {
	var (
		v0     int32
		v1     int32
		v2     int32
		result *uint32
	)
	v0 = nox_xxx_guiFontHeightMB_43F320(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1045480 + 236)))))) + 1
	sub_46AB20(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045480)), int32(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1045480 + 8)))), v0*15+2)
	sub_46AB20(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045508)), int32(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1045508 + 8)))), v0*15+2)
	v1 = 1120
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1045480 + 20))) + uint32(v0) + 2)
	for {
		result = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045484)))).ChildByID(v1)))
		*(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*5)) = uint32(v2)
		v2 += v0
		v1++
		if v1 >= 1134 {
			break
		}
	}
	return result
}
func sub_453B80(a1 int32, a2 int32) int32 {
	var (
		result int32
		xLeft  int32
		yTop   int32
	)
	nox_client_wndGetPosition_46AA60((*nox_window)(unsafe.Pointer(uintptr(a1))), (*uint32)(unsafe.Pointer(&xLeft)), (*uint32)(unsafe.Pointer(&yTop)))
	if int32(int8(*(*uint8)(unsafe.Pointer(uintptr(a1 + 4))))) >= 0 {
		if *(*uint32)(unsafe.Pointer(uintptr(a2 + 20))) != 0x80000000 {
			nox_client_drawRectFilledAlpha_49CF10(xLeft, yTop, int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 8)))), int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 12)))))
		}
		result = 1
	} else {
		nox_client_drawImageAt_47D2C0((*nox_video_bag_image_t)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a2 + 24)))))), xLeft, yTop)
		result = 1
	}
	return result
}
func sub_453F70(a1 unsafe.Pointer) {
	libc.MemCpy(memmap.PtrOff(6112660, 1045488), a1, 20)
}
func sub_453F90() *byte {
	return (*byte)(memmap.PtrOff(6112660, 1045488))
}
func sub_453FA0(a1 int32, a2 int32, a3 int32) int32 {
	var (
		v3     bool
		v4     int8
		result int32
		v6     uint8
	)
	v4 = int8(a2 & 31)
	v3 = (uint32(a2) & 0x8000001F & 0x80000000) != 0
	v6 = uint8(int8(a2 / 32))
	if v3 {
		v4 = int8(((int32(v4) - 1) | 224) + 1)
	}
	result = 1 << int32(v4)
	if a3 != 0 {
		*(*uint32)(unsafe.Pointer(uintptr(a1 + int32(v6)*4))) |= uint32(result)
	} else {
		result = ^result
		*(*uint32)(unsafe.Pointer(uintptr(a1 + int32(v6)*4))) &= uint32(result)
	}
	return result
}
func sub_454000(a1 int32, a2 int32) int32 {
	return bool2int((*(*uint32)(unsafe.Pointer(uintptr(a1 + ((a2/32)&math.MaxUint8)*4))) & uint32(1<<(a2%32))) != 0)
}
func sub_454040(a1 *uint32) int32 {
	var (
		v1     int32
		v2     int32
		result int32
		v4     uint8
	)
	v1 = 1
	*a1 = math.MaxUint32
	v4 = 0
	v2 = 1
	*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*1)) = math.MaxUint32
	*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*2)) = math.MaxUint32
	*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*3)) = math.MaxUint32
	*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*4)) = math.MaxUint32
	for {
		if uint32(v1) == 0x80000000 {
			v1 = 1
			v4++
		} else {
			v1 *= 2
		}
		result = bool2int(nox_xxx_spellIsValid_424B50(v2))
		if result != 0 {
			result = int32(nox_xxx_spellFlags_424A70(v2))
			if uint32(result)&0x7000000 != 0 {
				result = bool2int(nox_xxx_spellIsEnabled_424B70(v2))
				if result == 0 {
					result = int32(uintptr(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*uintptr(v4))))))
					*(*uint32)(unsafe.Pointer(uintptr(result))) = uint32(^v1) & *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*uintptr(v4)))
				}
			}
		}
		v2++
		if v2 >= 137 {
			break
		}
	}
	return result
}
func sub_4540E0(a1 int32) int32 {
	var (
		v1     int32
		result int32
	)
	v1 = 1
	for {
		if sub_454000(a1, v1) != 0 {
			result = bool2int(nox_xxx_spellEnable_424B90(v1))
		} else {
			result = bool2int(nox_xxx_spellDisable_424BB0(v1))
		}
		v1++
		if v1 >= 137 {
			break
		}
	}
	return result
}
func sub_454120() int8 {
	var (
		v0 int32
		v1 int32
		v2 *uint32
		v3 int32
		v4 int32
		v5 bool
		v6 int32
		v8 int32
	)
	v8 = int32(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1045480 + 32))))
	v0 = 1120
	v1 = sub_4A4800(v8) * 524
	for {
		v2 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045484)))).ChildByID(v0)))
		v3 = 0
		v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(v8 + 24))))
		if v4+v1 != -4 && int32(*(*uint16)(unsafe.Pointer(uintptr(v4 + v1 + 4)))) != 0 {
			v3 = nox_xxx_spellByTitle_424960((*libc.WChar)(unsafe.Pointer(uintptr(v4 + v1 + 4))))
			(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v2)))))).Show()
		} else {
			(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v2)))))).Hide()
		}
		v1 += 524
		if v3 != 0 {
			v5 = sub_454000(int32(uintptr(memmap.PtrOff(6112660, 1045488))), v3) == 0
			v6 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*9)))
			if !v5 {
				*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v6))), 0)) = uint8(int8(v6 | 4))
				goto LABEL_11
			}
		} else {
			v6 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*9)))
		}
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v6))), 0)) = uint8(int8(v6 & 251))
	LABEL_11:
		v0++
		*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*9)) = uint32(v6)
		if v0 > 1133 {
			break
		}
	}
	return int8(v6)
}
func nox_xxx_guiServerAccessLoad_4541D0(a1 int32) int32 {
	var (
		v2  int32
		v3  *uint32
		v4  *byte
		v5  *byte
		v6  *uint32
		v7  *uint32
		v8  *uint32
		v9  *uint32
		v10 *uint32
		v11 *uint32
		v12 *uint32
		v13 *uint32
		v14 int32
		v15 *uint32
		v16 *uint32
		v17 *uint32
	)
	if dword_5d4594_1045516 != 0 {
		return 0
	}
	v2 = strMan.Lang()
	if nox_xxx_guiFontHeightMB_43F320(nil) > 10 {
		v2 = 2
	}
	dword_5d4594_1045516 = uint32(uintptr(unsafe.Pointer(newWindowFromFile(*(**byte)(memmap.PtrOff(0x587000, uint32(v2*4)+0x1F350)), unsafe.Pointer(cgoFuncAddr(nox_xxx_windowAccessProc_454BA0))))))
	nox_draw_setTabWidth_43FE20(100)
	sub_46B120((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045516)))), (*nox_window)(unsafe.Pointer(uintptr(a1))))
	dword_5d4594_1045520 = uint32(uintptr(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045516)))).ChildByID(0x2776))))
	*memmap.PtrUint32(6112660, 1045524) = uint32(uintptr(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045516)))).ChildByID(0x2777))))
	dword_5d4594_1045532 = uint32(uintptr(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045516)))).ChildByID(0x277D))))
	dword_5d4594_1045528 = uint32(uintptr(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045516)))).ChildByID(0x2779))))
	dword_5d4594_1045536 = uint32(uintptr(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045516)))).ChildByID(10200))))
	dword_5d4594_1045540 = uint32(uintptr(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045516)))).ChildByID(0x277F))))
	dword_5d4594_1045544 = uint32(uintptr(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045516)))).ChildByID(0x2780))))
	dword_5d4594_1045548 = uint32(uintptr(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045516)))).ChildByID(0x2781))))
	dword_5d4594_1045556 = uint32(uintptr(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045516)))).ChildByID(0x2778))))
	*memmap.PtrUint32(6112660, 1045560) = uint32(uintptr(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045516)))).ChildByID(0x278D))))
	*memmap.PtrUint32(6112660, 1045564) = uint32(uintptr(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045516)))).ChildByID(0x278F))))
	*memmap.PtrUint32(6112660, 1045568) = uint32(uintptr(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045516)))).ChildByID(0x2791))))
	*memmap.PtrUint32(6112660, 1045572) = uint32(uintptr(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045516)))).ChildByID(0x2793))))
	dword_5d4594_1045576 = uint32(uintptr(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045516)))).ChildByID(0x278E))))
	dword_5d4594_1045580 = uint32(uintptr(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045516)))).ChildByID(0x2790))))
	dword_5d4594_1045584 = uint32(uintptr(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045516)))).ChildByID(10130))))
	dword_5d4594_1045588 = uint32(uintptr(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045516)))).ChildByID(0x2794))))
	dword_5d4594_1045552 = uint32(uintptr(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045516)))).ChildByID(0x278B))))
	*memmap.PtrUint32(6112660, 1045592) = uint32(uintptr(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045516)))).ChildByID(0x2795))))
	v3 = *(**uint32)(unsafe.Pointer(uintptr(dword_5d4594_1045532 + 32)))
	v4 = (*byte)(unsafe.Pointer(nox_xxx_gLoadImg("UISlider")))
	v5 = (*byte)(unsafe.Pointer(nox_xxx_gLoadImg("UISliderLit")))
	v6 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045516)))).ChildByID(10190)))
	v15 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045516)))).ChildByID(0x27CC)))
	v12 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045516)))).ChildByID(0x27CD)))
	*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*100)) + 8))) = 16
	*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*100)) + 12))) = 10
	sub_4B5700(int32(uintptr(unsafe.Pointer(v6))), 0, 0, int32(uintptr(unsafe.Pointer(v4))), int32(uintptr(unsafe.Pointer(v5))), int32(uintptr(unsafe.Pointer(v5))))
	nox_xxx_wnd_46B280(int32(uintptr(unsafe.Pointer(v6))), *(*int32)(unsafe.Pointer(&dword_5d4594_1045532)))
	nox_xxx_wnd_46B280(int32(uintptr(unsafe.Pointer(v15))), *(*int32)(unsafe.Pointer(&dword_5d4594_1045532)))
	nox_xxx_wnd_46B280(int32(uintptr(unsafe.Pointer(v12))), *(*int32)(unsafe.Pointer(&dword_5d4594_1045532)))
	*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*9)) = uint32(uintptr(unsafe.Pointer(v6)))
	*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*7)) = uint32(uintptr(unsafe.Pointer(v15)))
	*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*8)) = uint32(uintptr(unsafe.Pointer(v12)))
	v7 = *(**uint32)(unsafe.Pointer(uintptr(dword_5d4594_1045528 + 32)))
	v8 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045516)))).ChildByID(0x27CB)))
	v16 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045516)))).ChildByID(0x27C9)))
	v13 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045516)))).ChildByID(0x27CA)))
	*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v8), unsafe.Sizeof(uint32(0))*100)) + 8))) = 16
	*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v8), unsafe.Sizeof(uint32(0))*100)) + 12))) = 10
	sub_4B5700(int32(uintptr(unsafe.Pointer(v8))), 0, 0, int32(uintptr(unsafe.Pointer(v4))), int32(uintptr(unsafe.Pointer(v5))), int32(uintptr(unsafe.Pointer(v5))))
	nox_xxx_wnd_46B280(int32(uintptr(unsafe.Pointer(v8))), *(*int32)(unsafe.Pointer(&dword_5d4594_1045528)))
	nox_xxx_wnd_46B280(int32(uintptr(unsafe.Pointer(v16))), *(*int32)(unsafe.Pointer(&dword_5d4594_1045528)))
	nox_xxx_wnd_46B280(int32(uintptr(unsafe.Pointer(v13))), *(*int32)(unsafe.Pointer(&dword_5d4594_1045528)))
	*(*uint32)(unsafe.Add(unsafe.Pointer(v7), unsafe.Sizeof(uint32(0))*9)) = uint32(uintptr(unsafe.Pointer(v8)))
	*(*uint32)(unsafe.Add(unsafe.Pointer(v7), unsafe.Sizeof(uint32(0))*7)) = uint32(uintptr(unsafe.Pointer(v16)))
	*(*uint32)(unsafe.Add(unsafe.Pointer(v7), unsafe.Sizeof(uint32(0))*8)) = uint32(uintptr(unsafe.Pointer(v13)))
	v9 = *(**uint32)(unsafe.Pointer(uintptr(dword_5d4594_1045536 + 32)))
	v10 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045516)))).ChildByID(0x27DB)))
	v17 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045516)))).ChildByID(0x27D9)))
	v11 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045516)))).ChildByID(0x27DA)))
	*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v10), unsafe.Sizeof(uint32(0))*100)) + 8))) = 16
	v14 = int32(uintptr(unsafe.Pointer(v11)))
	*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v10), unsafe.Sizeof(uint32(0))*100)) + 12))) = 10
	sub_4B5700(int32(uintptr(unsafe.Pointer(v10))), 0, 0, int32(uintptr(unsafe.Pointer(v4))), int32(uintptr(unsafe.Pointer(v5))), int32(uintptr(unsafe.Pointer(v5))))
	nox_xxx_wnd_46B280(int32(uintptr(unsafe.Pointer(v10))), *(*int32)(unsafe.Pointer(&dword_5d4594_1045536)))
	nox_xxx_wnd_46B280(int32(uintptr(unsafe.Pointer(v17))), *(*int32)(unsafe.Pointer(&dword_5d4594_1045536)))
	nox_xxx_wnd_46B280(v14, *(*int32)(unsafe.Pointer(&dword_5d4594_1045536)))
	*(*uint32)(unsafe.Add(unsafe.Pointer(v9), unsafe.Sizeof(uint32(0))*9)) = uint32(uintptr(unsafe.Pointer(v10)))
	*(*uint32)(unsafe.Add(unsafe.Pointer(v9), unsafe.Sizeof(uint32(0))*7)) = uint32(uintptr(unsafe.Pointer(v17)))
	*(*uint32)(unsafe.Add(unsafe.Pointer(v9), unsafe.Sizeof(uint32(0))*8)) = uint32(v14)
	sub_454740()
	sub_454640()
	if noxflags.HasGame(noxflags.GameHost) {
		(*(*int32)(unsafe.Pointer(&dword_5d4594_1045516))).setDraw(sub_454A90)
	}
	nox_xxx_wndRetNULL_46A8A0()
	return int32(dword_5d4594_1045516)
}
func sub_454A90(a1 int32, a2 int32) int32 {
	var (
		v2    *uint16
		v3    int32
		v4    int8
		xLeft int32
		yTop  int32
	)
	nox_client_wndGetPosition_46AA60((*nox_window)(unsafe.Pointer(uintptr(a1))), (*uint32)(unsafe.Pointer(&xLeft)), (*uint32)(unsafe.Pointer(&yTop)))
	if int32(int8(*(*uint8)(unsafe.Pointer(uintptr(a1 + 4))))) >= 0 {
		if *(*uint32)(unsafe.Pointer(uintptr(a2 + 20))) != 0x80000000 {
			nox_client_drawRectFilledAlpha_49CF10(xLeft, yTop, int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 8)))), int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 12)))))
		}
	} else {
		nox_client_drawImageAt_47D2C0((*nox_video_bag_image_t)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a2 + 24)))))), xLeft, yTop)
	}
	if int32(*(*uint8)(unsafe.Pointer(uintptr(dword_5d4594_1045540 + 4))))&8 != 0 {
		v2 = (*uint16)(unsafe.Pointer(uintptr((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045540))))).Func94(asWindowEvent(0x401D, 0, 0)))))
		if v2 != nil && int32(*v2) != 0 {
			if (int32(*(*uint8)(unsafe.Pointer(uintptr(dword_5d4594_1045544 + 4)))) & 8) == 0 {
				nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045544))))), 1)
			}
		} else if int32(*(*uint8)(unsafe.Pointer(uintptr(dword_5d4594_1045544 + 4))))&8 != 0 {
			nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045544))))), 0)
		}
	}
	v3 = (*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045596))))).Func94(asWindowEvent(0x4014, 0, 0))
	v4 = int8(*(*uint8)(unsafe.Pointer(uintptr(dword_5d4594_1045548 + 4))))
	if v3 < 0 {
		if int32(v4)&8 != 0 {
			nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045548))))), 0)
		}
	} else if (int32(v4) & 8) == 0 {
		nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045548))))), 1)
		return 1
	}
	return 1
}
func nox_xxx_windowAccessProc_454BA0(a1 int32, a2 int32, a3 *int32, a4 int32) int32 {
	var (
		v4     int32
		result int32
		v6     *byte
	)
	_ = v6
	var v7 int32
	var v8 *byte
	var v9 *libc.WChar
	var v10 *byte
	var v11 *libc.WChar
	var v12 *byte
	_ = v12
	var v13 *libc.WChar
	var v14 *byte
	_ = v14
	var v15 *libc.WChar
	var v16 *byte
	var v17 int8
	var v18 *uint32
	var v19 *uint32
	var v20 *uint32
	var v21 *byte
	_ = v21
	var v22 *libc.WChar
	var v23 int32
	var v24 int32
	var v25 *byte
	_ = v25
	var v26 *int32
	var v27 *int32
	var v28 int32
	var v29 *libc.WChar
	var v30 *libc.WChar
	var v31 *byte
	var v32 *byte
	var v33 *int32
	var v34 *int32
	var v35 int32
	var v36 *libc.WChar
	var v37 *byte
	var v38 *byte
	var v39 *byte
	var v40 *libc.WChar
	var v41 *libc.WChar
	var v42 int32
	var v43 int32
	var v44 *uint32
	var v45 *libc.WChar
	var v46 *libc.WChar
	var v47 int32
	var WideCharStr [8]libc.WChar
	var v49 *byte
	var v50 *byte
	var v51 *byte
	switch a2 {
	case 0x4003:
		v43 = a4
		v44 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045516)))).ChildByID(a4)))
		v51 = sub_416630()
		v49 = (*byte)(sub_416640())
		if v44 == nil || int32(uint16(uintptr(unsafe.Pointer(a3)))) == 1 {
			goto LABEL_101
		}
		v45 = (*libc.WChar)(unsafe.Pointer(uintptr((*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v44)))))).Func94(asWindowEvent(0x401D, 0, 0)))))
		v46 = v45
		if v45 == nil || *v45 == 0 || (func() bool {
			v47 = nox_wcstol(v45, nil, 10)
			return v47 < 0
		}()) {
			v47 = 0
		}
		switch v43 {
		case 0x2778:
			nox_wcscpy((*libc.WChar)(unsafe.Add(unsafe.Pointer((*libc.WChar)(unsafe.Pointer(v49))), unsafe.Sizeof(libc.WChar(0))*39)), v46)
			goto LABEL_101
		case 0x278E:
			if v47 > 14 {
				*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v47))), 0)) = 14
				compat_itow(14, &WideCharStr[0], 10)
				(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v44)))))).Func94(asWindowEvent(0x401E, int32(uintptr(unsafe.Pointer(&WideCharStr[0]))), -1))
			}
			*(*byte)(unsafe.Add(unsafe.Pointer(v51), 1)) = byte(int8(v47 | int32(*(*byte)(unsafe.Add(unsafe.Pointer(v51), 1))&240)))
			result = 0
		case 0x2790:
			if v47 > 14 {
				*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v47))), 0)) = 14
				compat_itow(14, &WideCharStr[0], 10)
				(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v44)))))).Func94(asWindowEvent(0x401E, int32(uintptr(unsafe.Pointer(&WideCharStr[0]))), -1))
			}
			*(*byte)(unsafe.Add(unsafe.Pointer(v51), 1)) = byte(int8((v47 * 16) | int32(*(*byte)(unsafe.Add(unsafe.Pointer(v51), 1))&15)))
			result = 0
		case 10130:
			*(*uint16)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v51), 5)))) = uint16(int16(v47))
			result = 0
		case 0x2794:
			*(*uint16)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v51), 7)))) = uint16(int16(v47))
			result = 0
		case 0x2795:
			if v47 <= 32 {
				if v47 < 1 {
					v47 = 1
					compat_itow(1, &WideCharStr[0], 10)
					(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v44)))))).Func94(asWindowEvent(0x401E, int32(uintptr(unsafe.Pointer(&WideCharStr[0]))), -1))
				}
			} else {
				v47 = 32
				compat_itow(32, &WideCharStr[0], 10)
				(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v44)))))).Func94(asWindowEvent(0x401E, int32(uintptr(unsafe.Pointer(&WideCharStr[0]))), -1))
			}
			nox_xxx_servSetPlrLimit_409F80(v47)
			result = 0
			*(*byte)(unsafe.Add(unsafe.Pointer(v51), 4)) = byte(int8(v47))
		case 0x2798:
			nox_xxx_sysopSetPass_40A610(v46)
			result = 0
		default:
			goto LABEL_101
		}
	case 0x4007:
		v7 = (*nox_window)(unsafe.Pointer(a3)).ID()
		clientPlaySoundSpecial(766, 100)
		switch v7 {
		case 0x2776:
			v16 = sub_416630()
			v17 = int8(*v16 ^ 16)
			*v16 = byte(v17)
			if int32(v17)&16 != 0 {
				v20 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045516)))).ChildByID(0x27DE)))
				nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v20)))))), 1)
			} else {
				v18 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045516)))).ChildByID(0x27DE)))
				nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v18)))))), 0)
				*(*uint32)(unsafe.Add(unsafe.Pointer(v18), unsafe.Sizeof(uint32(0))*9)) &= 0xFFFFFFFB
				v19 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045516)))).ChildByID(0x27DF)))
				*(*uint32)(unsafe.Add(unsafe.Pointer(v19), unsafe.Sizeof(uint32(0))*9)) |= 4
				(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045532))))).Hide()
				sub_46ACE0(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045516)), 0x27CC, 10190, 1)
				(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045528))))).Show()
				sub_46ACE0(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045516)), 0x27C9, 0x27CB, 0)
			}
			return 0
		case 0x2777:
			v21 = sub_416630()
			nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045556))))), int32((^*(*uint32)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, 1045524) + 36)))>>2)&1))
			*v21 ^= 32
			return 0
		case 0x2780:
			v22 = (*libc.WChar)(unsafe.Pointer(uintptr((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045540))))).Func94(asWindowEvent(0x401D, 0, 0)))))
			if (*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045528))))).Flags().IsHidden() != 0 {
				sub_4168A0(v22)
			} else {
				sub_416770(0, v22, nil)
			}
			(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045540))))).Func94(asWindowEvent(0x401E, int32(uintptr(memmap.PtrOff(6112660, 1045600))), 0))
			return 0
		case 0x2781:
			if int32(*(*uint8)(unsafe.Pointer(uintptr(dword_5d4594_1045520 + 36))))&4 != 0 {
				v23 = (*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045532))))).Func94(asWindowEvent(0x4014, 0, 0))
				(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045532))))).Func94(asWindowEvent(0x400E, v23, 0))
				sub_416860(v23)
			} else {
				v24 = (*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045528))))).Func94(asWindowEvent(0x4014, 0, 0))
				(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045528))))).Func94(asWindowEvent(0x400E, v24, 0))
				sub_416820(v24)
			}
			return 0
		case 0x278C:
			v25 = sub_416630()
			*(*byte)(unsafe.Add(unsafe.Pointer(v25), 2)) ^= 128
			return 0
		case 0x278D:
			v8 = sub_416630()
			if int32(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, 1045560) + 36))))&4 != 0 {
				nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045576))))), 0)
				*(*byte)(unsafe.Add(unsafe.Pointer(v8), 1)) |= 15
				return 0
			}
			nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045576))))), 1)
			v9 = (*libc.WChar)(unsafe.Pointer(uintptr((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045576))))).Func94(asWindowEvent(0x401D, 0, 0)))))
			if *v9 == 0 {
				goto LABEL_101
			}
			*(*byte)(unsafe.Add(unsafe.Pointer(v8), 1)) = byte(int8(int32(*(*byte)(unsafe.Add(unsafe.Pointer(v8), 1))&240) | nox_wcstol(v9, nil, 10)))
			return 0
		case 0x278F:
			v10 = sub_416630()
			if int32(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, 1045564) + 36))))&4 != 0 {
				nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045580))))), 0)
				*(*byte)(unsafe.Add(unsafe.Pointer(v10), 1)) |= 240
				return 0
			}
			nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045580))))), 1)
			v11 = (*libc.WChar)(unsafe.Pointer(uintptr((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045580))))).Func94(asWindowEvent(0x401D, 0, 0)))))
			if *v11 == 0 {
				goto LABEL_101
			}
			*(*byte)(unsafe.Add(unsafe.Pointer(v10), 1)) = byte(int8(int32(*(*byte)(unsafe.Add(unsafe.Pointer(v10), 1))&15) | nox_wcstol(v11, nil, 10)))
			return 0
		case 0x2791:
			v12 = sub_416630()
			if int32(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, 1045568) + 36))))&4 != 0 {
				nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045584))))), 0)
				*(*uint16)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v12), 5)))) = math.MaxUint16
				return 0
			}
			nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045584))))), 1)
			v13 = (*libc.WChar)(unsafe.Pointer(uintptr((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045584))))).Func94(asWindowEvent(0x401D, 0, 0)))))
			if *v13 == 0 {
				goto LABEL_101
			}
			*(*uint16)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v12), 5)))) = uint16(int16(nox_wcstol(v13, nil, 10)))
			return 0
		case 0x2793:
			v14 = sub_416630()
			if int32(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, 1045572) + 36))))&4 != 0 {
				nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045588))))), 0)
				*(*uint16)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v14), 7)))) = math.MaxUint16
				return 0
			}
			nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045588))))), 1)
			v15 = (*libc.WChar)(unsafe.Pointer(uintptr((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045588))))).Func94(asWindowEvent(0x401D, 0, 0)))))
			if *v15 == 0 {
				goto LABEL_101
			}
			*(*uint16)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v14), 7)))) = uint16(int16(nox_wcstol(v15, nil, 10)))
			result = 0
		case 0x27CF:
			v33 = (*int32)(unsafe.Pointer(uintptr((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045536))))).Func94(asWindowEvent(0x4014, 0, 0)))))
			v34 = v33
			v35 = *v33
			if v35 < 0 {
				goto LABEL_101
			}
			for {
				v36 = (*libc.WChar)(unsafe.Pointer(uintptr((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045536))))).Func94(asWindowEvent(0x4016, v35, 0)))))
				v37 = nox_xxx_playerByName_4170D0(v36)
				v38 = v37
				if v37 != nil && *(*byte)(unsafe.Add(unsafe.Pointer(v37), 2064)) != 31 {
					if noxflags.HasGame(noxflags.GameModeQuest) {
						sub_4DCFB0((*nox_object_t)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v38))), unsafe.Sizeof(uint32(0))*514)))))))
					} else {
						nox_xxx_playerCallDisconnect_4DEAB0(int32(uint8(*(*byte)(unsafe.Add(unsafe.Pointer(v38), 2064)))), 4)
					}
				}
				v35 = *(*int32)(unsafe.Add(unsafe.Pointer(v34), unsafe.Sizeof(int32(0))*1))
				v34 = (*int32)(unsafe.Add(unsafe.Pointer(v34), unsafe.Sizeof(int32(0))*1))
				if v35 < 0 {
					break
				}
			}
			return 0
		case 0x27D0:
			v26 = (*int32)(unsafe.Pointer(uintptr((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045536))))).Func94(asWindowEvent(0x4014, 0, 0)))))
			v27 = v26
			v28 = *v26
			if v28 < 0 {
				goto LABEL_101
			}
			for {
				v29 = (*libc.WChar)(unsafe.Pointer(uintptr((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045536))))).Func94(asWindowEvent(0x4016, v28, 0)))))
				v30 = v29
				v31 = nox_xxx_playerByName_4170D0(v29)
				v32 = v31
				if v31 != nil && *(*byte)(unsafe.Add(unsafe.Pointer(v31), 2064)) != 31 {
					if noxflags.HasGame(noxflags.GameModeQuest) {
						sub_4DCFB0((*nox_object_t)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v32))), unsafe.Sizeof(uint32(0))*514)))))))
					} else {
						nox_xxx_playerDisconnByPlrID_4DEB00(int32(uint8(*(*byte)(unsafe.Add(unsafe.Pointer(v32), 2064)))))
					}
					sub_416770(0, v30, (*byte)(unsafe.Add(unsafe.Pointer(v32), 2112)))
				}
				v28 = *(*int32)(unsafe.Add(unsafe.Pointer(v27), unsafe.Sizeof(int32(0))*1))
				v27 = (*int32)(unsafe.Add(unsafe.Pointer(v27), unsafe.Sizeof(int32(0))*1))
				if v28 < 0 {
					break
				}
			}
			return 0
		case 0x27DE:
			(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045532))))).Show()
			sub_46ACE0(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045516)), 0x27CC, 10190, 0)
			(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045528))))).Hide()
			sub_46ACE0(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045516)), 0x27C9, 0x27CB, 1)
			dword_5d4594_1045596 = dword_5d4594_1045532
			return 0
		case 0x27DF:
			(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045532))))).Hide()
			sub_46ACE0(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045516)), 0x27CC, 10190, 1)
			(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045528))))).Show()
			sub_46ACE0(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045516)), 0x27C9, 0x27CB, 0)
			dword_5d4594_1045596 = dword_5d4594_1045528
			return 0
		default:
			goto LABEL_101
		}
	case 16400:
		v4 = (*nox_window)(unsafe.Pointer(a3)).ID() - 0x278B
		if v4 != 0 {
			if v4 != 77 {
				goto LABEL_101
			}
			if sub_455770() != 0 {
				sub_46AD20(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045516)), 0x27CF, 0x27D0, 1)
			} else {
				sub_46AD20(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045516)), 0x27CF, 0x27D0, 0)
			}
			result = 0
		} else {
			v6 = (*byte)(sub_416640())
			*(*byte)(unsafe.Add(unsafe.Pointer(v6), 100)) ^= byte(int8(1 << a4))
			result = 0
		}
	case 0x401F:
		v39 = sub_416630()
		v50 = (*byte)(sub_416640())
		v40 = (*libc.WChar)(unsafe.Pointer(uintptr((*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a3)))))).Func94(asWindowEvent(0x401D, 0, 0)))))
		v41 = v40
		if v40 == nil || *v40 == 0 || (func() bool {
			v42 = nox_wcstol(v40, nil, 10)
			return v42 < 0
		}()) {
			v42 = 0
		}
		switch (*nox_window)(unsafe.Pointer(a3)).ID() {
		case 0x2778:
			nox_wcscpy((*libc.WChar)(unsafe.Add(unsafe.Pointer((*libc.WChar)(unsafe.Pointer(v50))), unsafe.Sizeof(libc.WChar(0))*39)), v41)
			goto LABEL_101
		case 0x278E:
			if v42 > 14 {
				*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v42))), 0)) = 14
				compat_itow(14, &WideCharStr[0], 10)
				(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a3)))))).Func94(asWindowEvent(0x401E, int32(uintptr(unsafe.Pointer(&WideCharStr[0]))), -1))
			}
			result = 0
			*(*byte)(unsafe.Add(unsafe.Pointer(v39), 1)) = byte(int8(v42 | int32(*(*byte)(unsafe.Add(unsafe.Pointer(v39), 1))&240)))
		case 0x2790:
			if v42 > 14 {
				*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v42))), 0)) = 14
				compat_itow(14, &WideCharStr[0], 10)
				(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a3)))))).Func94(asWindowEvent(0x401E, int32(uintptr(unsafe.Pointer(&WideCharStr[0]))), -1))
			}
			result = 0
			*(*byte)(unsafe.Add(unsafe.Pointer(v39), 1)) = byte(int8((v42 * 16) | int32(*(*byte)(unsafe.Add(unsafe.Pointer(v39), 1))&15)))
		case 10130:
			*(*uint16)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v39), 5)))) = uint16(int16(v42))
			result = 0
		case 0x2794:
			*(*uint16)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v39), 7)))) = uint16(int16(v42))
			result = 0
		case 0x2795:
			if v42 <= 32 {
				if v42 < 1 {
					v42 = 1
					compat_itow(1, &WideCharStr[0], 10)
					(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a3)))))).Func94(asWindowEvent(0x401E, int32(uintptr(unsafe.Pointer(&WideCharStr[0]))), -1))
				}
			} else {
				v42 = 32
				compat_itow(32, &WideCharStr[0], 10)
				(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a3)))))).Func94(asWindowEvent(0x401E, int32(uintptr(unsafe.Pointer(&WideCharStr[0]))), -1))
			}
			nox_xxx_servSetPlrLimit_409F80(v42)
			*(*byte)(unsafe.Add(unsafe.Pointer(v39), 4)) = byte(int8(v42))
			result = 0
		case 0x2798:
			nox_xxx_sysopSetPass_40A610(v41)
			result = 0
		default:
			goto LABEL_101
		}
	default:
	LABEL_101:
		result = 0
	}
	return result
}
func sub_455770() int32 {
	var (
		v0 *int32
		v1 *int32
		v2 int32
		v3 *libc.WChar
		v4 *byte
	)
	v0 = (*int32)(unsafe.Pointer(uintptr((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045536))))).Func94(asWindowEvent(0x4014, 0, 0)))))
	v1 = v0
	v2 = *v0
	if v2 < 0 {
		return 0
	}
	for {
		v3 = (*libc.WChar)(unsafe.Pointer(uintptr((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045536))))).Func94(asWindowEvent(0x4016, v2, 0)))))
		v4 = nox_xxx_playerByName_4170D0(v3)
		if v4 != nil {
			if *(*byte)(unsafe.Add(unsafe.Pointer(v4), 2064)) != 31 {
				break
			}
		}
		v2 = *(*int32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(int32(0))*1))
		v1 = (*int32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(int32(0))*1))
		if v2 < 0 {
			return 0
		}
	}
	return 1
}
func sub_4557D0(a1 int32) int32 {
	var result int32
	result = a1
	if a1 != 0 {
		result = int32(dword_5d4594_1045516)
		if dword_5d4594_1045516 != 0 {
			result = (*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045516)))).Destroy()
		}
	}
	dword_5d4594_1045516 = 0
	return result
}
func sub_455800() *int32 {
	var (
		result      *int32
		v1          int32
		i           *int32
		j           *int32
		WideCharStr [10]libc.WChar
		v5          [26]libc.WChar
	)
	result = *(**int32)(unsafe.Pointer(&dword_5d4594_1045516))
	if dword_5d4594_1045516 != 0 {
		v1 = nox_xxx_servGetPlrLimit_409FA0()
		compat_itow(v1, &WideCharStr[0], 10)
		(*nox_window)(unsafe.Pointer(uintptr(*memmap.PtrInt32(6112660, 1045592)))).Func94(asWindowEvent(0x401E, int32(uintptr(unsafe.Pointer(&WideCharStr[0]))), 0))
		result = (*int32)(unsafe.Pointer(uintptr(bool2int(noxflags.HasGame(noxflags.GameHost)))))
		if result != nil {
			(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045528))))).Func94(asWindowEvent(0x400F, 0, 0))
			(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045532))))).Func94(asWindowEvent(0x400F, 0, 0))
			for i = sub_4168E0(); i != nil; i = sub_4168F0(i) {
				(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045532))))).Func94(asWindowEvent(0x400D, int32(uintptr(unsafe.Pointer((*int32)(unsafe.Add(unsafe.Pointer(i), unsafe.Sizeof(int32(0))*3))))), -1))
			}
			result = sub_416900()
			for j = result; result != nil; j = result {
				if int32(*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(j))), 72)))) != 0 {
					(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045528))))).Func94(asWindowEvent(0x400D, int32(uintptr(unsafe.Pointer((*int32)(unsafe.Add(unsafe.Pointer(j), unsafe.Sizeof(int32(0))*3))))), -1))
				} else {
					nox_swprintf(&v5[0], libc.CWString("*%s"), (*int32)(unsafe.Add(unsafe.Pointer(j), unsafe.Sizeof(int32(0))*3)))
					(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045528))))).Func94(asWindowEvent(0x400D, int32(uintptr(unsafe.Pointer(&v5[0]))), -1))
				}
				result = sub_416910(j)
			}
		}
	}
	return result
}
func sub_455920(a1 int32) int32 {
	var (
		result int32
		v2     *uint32
	)
	result = int32(dword_5d4594_1045516)
	if dword_5d4594_1045516 != 0 {
		v2 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045516)))).ChildByID(10200)))
		result = (*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v2)))))).Func94(asWindowEvent(0x400D, a1, 3))
	}
	return result
}
func sub_455950(a1 *libc.WChar) {
	var v1 int32
	if dword_5d4594_1045516 != 0 {
		v1 = sub_4559B0(a1)
		if v1 != -1 {
			(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045536))))).Func94(asWindowEvent(0x400E, v1, 0))
			if sub_455770() == 0 {
				sub_46AD20(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045516)), 0x27CF, 0x27D0, 0)
			}
		}
	}
}
func sub_4559B0(a1 *libc.WChar) int32 {
	var (
		v1 int32
		v2 int32
		v3 *libc.WChar
	)
	v1 = 0
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1045536 + 32))))
	if int32(*(*int16)(unsafe.Pointer(uintptr(v2 + 44)))) <= 0 {
		return -1
	}
	for {
		v3 = (*libc.WChar)(unsafe.Pointer(uintptr((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045536))))).Func94(asWindowEvent(0x4016, v1, 0)))))
		if nox_wcscmp(v3, a1) == 0 {
			break
		}
		if func() int32 {
			p := &v1
			*p++
			return *p
		}() >= int32(*(*int16)(unsafe.Pointer(uintptr(v2 + 44)))) {
			return -1
		}
	}
	return v1
}
func sub_455A00(a1 int32) int32 {
	var result int32
	if a1 != 0 && *memmap.PtrUint32(6112660, 1045608) != 0 && nox_xxx_wndGetFlags_46ADA0(*(*int32)(unsafe.Pointer(&dword_5d4594_1045604)))&16 != 0 {
		result = nox_xxx_wndShowModalMB((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045604))))))
	} else {
		result = (*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045604))))).Hide()
	}
	return result
}
func sub_455A50(a1 int8) int8 {
	var (
		v1 *uint32
		v2 int32
		v3 int32
		v4 int32
	)
	v1 = nil
	nox_xxx_objGetTeamByNetCode_418C80(int32(nox_player_netCode_85319C))
	if !(dword_5d4594_1045604 != 0 || (func() int32 {
		v2 = sub_455C30()
		return v2
	}()) != 0) {
		return int8(v2)
	}
	var cur_w int32
	var cur_h int32
	var cur_d int32
	nox_xxx_gameGetScreenBoundaries_43BEB0_get_video_mode(&cur_w, &cur_h, &cur_d)
	var max_w int32
	var max_h int32
	nox_xxx_screenGetSize_430C50_get_video_max(&max_w, &max_h)
	if cur_w > max_w {
		cur_w = max_w
	}
	if cur_h > max_h {
		cur_h = max_h
	}
	*memmap.PtrUint32(6112660, 1045612) = 0
	*memmap.PtrUint32(6112660, 1045616) = 0
	*memmap.PtrUint32(6112660, 1045620) = 0
	*memmap.PtrUint32(6112660, 1045608) = 1
	*memmap.PtrUint32(6112660, 1045624) = 0
	sub_46ACE0(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045604)), 8811, 8826, 1)
	v3 = 0
	*memmap.PtrUint8(6112660, 1045628) = uint8(a1)
	if int32(a1) != 0 {
		for {
			v1 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045604)))).ChildByID(v3 + 8811)))
			(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v1)))))).Show()
			v3++
			if v3 >= int32(*memmap.PtrUint8(6112660, 1045628)) {
				break
			}
		}
	}
	if int32(*memmap.PtrUint8(6112660, 1045628)) <= 4 {
		v4 = int32(uint32(cur_w) - *(*uint32)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045604)) + 8)))/2)
	} else {
		v4 = int32(uint32(cur_w) - *(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1045604 + 8))))
	}
	*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1045604 + 16))) = uint32(v4 - 91)
	*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1045604 + 24))) = *(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1045604 + 8))) + *(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1045604 + 16)))
	*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v2))), 0)) = *memmap.PtrUint8(6112660, 1045628)
	if int32(*memmap.PtrUint8(6112660, 1045628)) <= 8 {
		if int32(*memmap.PtrUint8(6112660, 1045628)) <= 4 {
			if v1 == nil {
				*memmap.PtrUint32(6112660, 1045608) = 0
				return int8(v2)
			}
			*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1045604 + 20))) = uint32(cur_h - int32(*memmap.PtrUint8(6112660, 1045628))*40)
		} else {
			*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1045604 + 20))) = uint32(cur_h - *(*int32)(unsafe.Pointer(uintptr(dword_5d4594_1045604 + 12)))/2)
		}
	} else {
		*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1045604 + 20))) = uint32(cur_h) - *(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1045604 + 12)))
	}
	*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1045604 + 28))) = *(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1045604 + 12))) + *(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1045604 + 20)))
	*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v2))), 0)) = uint8(int8(nox_xxx_wndShowModalMB((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045604))))))))
	return int8(v2)
}
func sub_455C10() int32 {
	*memmap.PtrUint32(6112660, 1045608) = 0
	return sub_455A00(0)
}
func sub_455CD0(a1 *uint8, a2 *uint32) int32 {
	var (
		v2 *uint8
		v3 int32
		v5 *uint8
		v6 int32
		v7 int32
		v8 uint8
	)
	v2 = a1
	v5 = a1
	v8 = uint8(int8(int32(*a1) - 107))
	nox_client_wndGetPosition_46AA60((*nox_window)(unsafe.Pointer(v5)), (*uint32)(unsafe.Pointer(&v6)), (*uint32)(unsafe.Pointer(&v7)))
	if int32(*memmap.PtrUint8(6112660, uint32(v8)+0xFF46C)) == 1 {
		v3 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*8)))
		if v3 != 0 {
		LABEL_8:
			nox_client_drawImageAt_47D2C0((*nox_video_bag_image_t)(unsafe.Pointer(uintptr(v3))), v6, v7)
			goto LABEL_9
		}
	} else if int32(*memmap.PtrUint8(6112660, uint32(v8)+0xFF46C)) == 2 {
		v3 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*12)))
		if v3 != 0 {
			goto LABEL_8
		}
	} else {
		v3 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*6)))
		if v3 != 0 {
			goto LABEL_8
		}
	}
LABEL_9:
	if int32(*(*uint8)(unsafe.Add(unsafe.Pointer(v2), 4)))&32 != 0 && *memmap.PtrUint32(6112660, 1045632) != 0 {
		nox_client_drawImageAt_47D2C0((*nox_video_bag_image_t)(unsafe.Pointer(uintptr(*memmap.PtrInt32(6112660, 1045632)))), v6-4, v7-4)
	}
	return 1
}
func sub_455E70(a1 uint8) int32 {
	var (
		v2 *uint32
		v3 *uint32
	)
	for i := int32(0); i < int32(*memmap.PtrUint8(6112660, 1045628)); i++ {
		v2 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045604)))).ChildByID(i + 8811)))
		nox_xxx_wndClearFlag_46AD80(int32(uintptr(unsafe.Pointer(v2))), 32)
	}
	v3 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045604)))).ChildByID(int32(a1) + 8810)))
	return nox_xxx_wnd_46AD60(int32(uintptr(unsafe.Pointer(v3))), 32)
}
func sub_455EE0() int32 {
	var result int32
	result = int32(dword_5d4594_1045604)
	if dword_5d4594_1045604 != 0 {
		result = (*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045604)))).Destroy()
	}
	dword_5d4594_1045604 = 0
	*memmap.PtrUint32(6112660, 1045608) = 0
	return result
}
func sub_455F10(a1 int32) int32 {
	var result int32
	if a1 != 0 && dword_5d4594_1045640 != 0 && nox_xxx_wndGetFlags_46ADA0(*(*int32)(unsafe.Pointer(&dword_5d4594_1045636)))&16 != 0 {
		result = nox_xxx_wndShowModalMB((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045636))))))
	} else {
		result = (*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045636))))).Hide()
	}
	return result
}
func sub_455F60() int32 {
	var (
		result int32
		v1     int32
		v2     int32
		v3     int32
		v4     int32
		v5     int32
		v6     int32
	)
	nox_xxx_objGetTeamByNetCode_418C80(int32(nox_player_netCode_85319C))
	if dword_5d4594_1045636 != 0 || (func() int32 {
		result = sub_456070()
		return result
	}()) != 0 {
		nox_xxx_gameGetScreenBoundaries_43BEB0_get_video_mode(&v3, &v2, &v6)
		nox_xxx_screenGetSize_430C50_get_video_max(&v4, &v5)
		v1 = v3
		if v3 > v4 {
			v1 = v4
			v3 = v4
		}
		if v2 > v5 {
			v2 = v5
		}
		dword_5d4594_1045640 = 1
		*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1045636 + 16))) = uint32(v1) - *(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1045636 + 8)))/3 - 91
		*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1045636 + 24))) = *(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1045636 + 8))) + *(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1045636 + 16)))
		*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1045636 + 20))) = uint32(v2 - 120)
		*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1045636 + 28))) = *(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1045636 + 12))) + *(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1045636 + 20)))
		*memmap.PtrUint8(6112660, 1045644) = 0
		result = nox_xxx_wndShowModalMB((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045636))))))
	}
	return result
}
func sub_456050() int32 {
	dword_5d4594_1045640 = 0
	return sub_455F10(0)
}
func sub_456070() int32 {
	if dword_5d4594_1045636 == 0 {
		dword_5d4594_1045636 = uint32(uintptr(unsafe.Pointer(newWindowFromFile("gui_fb.wnd", nil))))
		(*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045636)))).SetAllFuncs(nil, func(arg1 *nox_window, arg2 unsafe.Pointer) int32 {
			return sub_4560D0(int32(uintptr(unsafe.Pointer(arg1))), int32(uintptr(arg2)))
		}, nil)
		if dword_5d4594_1045636 == 0 {
			return 0
		}
		sub_455F10(0)
		*memmap.PtrUint32(6112660, 1045648) = uint32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg("FlagTeamBorder"))))
	}
	return 1
}
func sub_4560D0(a1 int32, a2 int32) int32 {
	var (
		v2 int32
		v4 int32
		v5 int32
	)
	nox_client_wndGetPosition_46AA60((*nox_window)(unsafe.Pointer(uintptr(a1))), (*uint32)(unsafe.Pointer(&v4)), (*uint32)(unsafe.Pointer(&v5)))
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 24))))
	if v2 != 0 {
		nox_client_drawImageAt_47D2C0((*nox_video_bag_image_t)(unsafe.Pointer(uintptr(v2))), v4, v5)
	}
	if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 4))))&32 != 0 && *memmap.PtrUint32(6112660, 1045648) != 0 {
		nox_client_drawImageAt_47D2C0((*nox_video_bag_image_t)(unsafe.Pointer(uintptr(*memmap.PtrInt32(6112660, 1045648)))), v4-4, v5-4)
	}
	return 1
}
func sub_456240() int32 {
	var result int32
	result = int32(dword_5d4594_1045636)
	if dword_5d4594_1045636 != 0 {
		result = (*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045636)))).Destroy()
	}
	dword_5d4594_1045636 = 0
	dword_5d4594_1045640 = 0
	return result
}
func sub_456500() int32 {
	var (
		v0     *uint32
		v1     *uint32
		i      *byte
		j      *byte
		v4     *uint32
		result int32
	)
	v0 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045684)))).ChildByID(0x2905)))
	v1 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045684)))).ChildByID(0x2906)))
	(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v0)))))).Func94(asWindowEvent(0x400F, 0, 0))
	(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v1)))))).Func94(asWindowEvent(0x400F, 0, 0))
	nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045688))))), 0)
	nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045692))))), 0)
	for i = (*byte)(unsafe.Pointer(nox_server_teamFirst_418B10())); i != nil; i = (*byte)(unsafe.Pointer(nox_server_teamNext_418B60((*nox_team_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(i))))))))) {
		sub_457230((*libc.WChar)(unsafe.Pointer(i)))
	}
	for j = (*byte)(unsafe.Pointer(nox_common_playerInfoGetFirst_416EA0())); j != nil; j = (*byte)(unsafe.Pointer(nox_common_playerInfoGetNext_416EE0((*nox_playerInfo)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(j))))))))) {
		if *(*byte)(unsafe.Add(unsafe.Pointer(j), 2064)) != 31 || !nox_common_getEngineFlag(nox_engine_flag(NOX_ENGINE_FLAG_DISABLE_GRAPHICS_RENDERING)) {
			sub_457140(int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(j))), unsafe.Sizeof(uint32(0))*515)))), (*libc.WChar)(unsafe.Add(unsafe.Pointer((*libc.WChar)(unsafe.Pointer(j))), unsafe.Sizeof(libc.WChar(0))*2352)))
			v4 = nox_xxx_objGetTeamByNetCode_418C80(int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(j))), unsafe.Sizeof(uint32(0))*515)))))
			if nox_xxx_servObjectHasTeam_419130(int32(uintptr(unsafe.Pointer(v4)))) != 0 {
				sub_4571A0(int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(j))), unsafe.Sizeof(uint32(0))*515)))), int32(*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(v4))), 4)))))
			}
		}
	}
	if noxflags.HasGame(noxflags.GameModeQuest) {
		nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v0)))))), 0)
		result = nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v1)))))), 0)
	} else {
		nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v0)))))), 1)
		result = nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v1)))))), 1)
	}
	return result
}
func sub_456640(a1 int32, a2 int32) int32 {
	var (
		v2    *uint32
		v3    *uint32
		v4    *uint32
		v5    *uint32
		v6    *uint32
		v7    *uint32
		xLeft int32
		yTop  int32
	)
	nox_xxx_objGetTeamByNetCode_418C80(int32(nox_player_netCode_85319C))
	nox_client_wndGetPosition_46AA60((*nox_window)(unsafe.Pointer(uintptr(a1))), (*uint32)(unsafe.Pointer(&xLeft)), (*uint32)(unsafe.Pointer(&yTop)))
	if int32(int8(*(*uint8)(unsafe.Pointer(uintptr(a1 + 4))))) >= 0 {
		if *(*uint32)(unsafe.Pointer(uintptr(a2 + 20))) != 0x80000000 {
			nox_client_drawRectFilledAlpha_49CF10(xLeft, yTop, int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 8)))), int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 12)))))
		}
	} else {
		nox_client_drawImageAt_47D2C0((*nox_video_bag_image_t)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a2 + 24)))))), xLeft, yTop)
	}
	if noxflags.HasGame(noxflags.GameHost) && !noxflags.HasGame(noxflags.GameFlag16) && sub_40A740() == 0 {
		v2 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045684)))).ChildByID(0x2906)))
		if (*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v2)))))).Func94(asWindowEvent(0x4014, 0, 0)) < 0 {
			v6 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045684)))).ChildByID(0x2907)))
			nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v6)))))), 0)
		} else {
			v3 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045684)))).ChildByID(0x2905)))
			v4 = (*uint32)(unsafe.Pointer(uintptr((*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v3)))))).Func94(asWindowEvent(0x4014, 0, 0)))))
			v5 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045684)))).ChildByID(0x2907)))
			nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v5)))))), bool2int(*v4 >= 0))
		}
	}
	v7 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045684)))).ChildByID(0x2906)))
	if (*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v7)))))).Func94(asWindowEvent(0x4014, 0, 0)) < 0 {
		nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045688))))), 0)
		nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045692))))), 0)
	}
	return 1
}
func sub_456BB0(a1 int32) int8 {
	var (
		v1 *uint32
		v2 int32
		v3 int32
		v4 *uint32
		v6 float2
	)
	*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v1))), 0)) = *memmap.PtrUint8(0x8531A0, 2576)
	if *memmap.PtrUint32(0x8531A0, 2576) == 0 || (int32(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(0x8531A0, 2576) + 4))))&1) == 0 && (int32(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(0x8531A0, 2576) + 3680))))&1) == 0 {
		v2 = bool2int(noxflags.HasGame(noxflags.GameHost))
		v1 = nox_xxx_objGetTeamByNetCode_418C80(int32(nox_player_netCode_85319C))
		v3 = int32(uintptr(unsafe.Pointer(v1)))
		if v1 != nil {
			if nox_xxx_servObjectHasTeam_419130(int32(uintptr(unsafe.Pointer(v1)))) != 0 {
				if v2 == 0 {
					sub_419960(a1, v3, int16(uint16(nox_player_netCode_85319C)))
					*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v1))), 0)) = uint8(int8(int32(*memmap.PtrUint8(6112660, 1045696)) + 1))
					*memmap.PtrUint32(6112660, 1045696)++
					return int8(uintptr(unsafe.Pointer(v1)))
				}
				sub_4196D0(v3, a1, int32(nox_player_netCode_85319C), 1)
				v1 = (*uint32)(unsafe.Pointer(uintptr(bool2int(noxflags.HasGame(noxflags.GameModeChat)))))
				if v1 == nil {
					v1 = &noxServer.getPlayerByID(int32(nox_player_netCode_85319C)).field_0
					v4 = v1
					if v1 != nil {
						nox_xxx_mapFindPlayerStart_4F7AB0(&v6, (*nox_object_t)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*514))))))
						nox_xxx_unitMove_4E7010((*nox_object_t)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*514))))), &v6)
						*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v1))), 0)) = uint8(int8(int32(*memmap.PtrUint8(6112660, 1045696)) + 1))
						*memmap.PtrUint32(6112660, 1045696)++
						return int8(uintptr(unsafe.Pointer(v1)))
					}
				}
			} else {
				if v2 != 0 {
					nox_xxx_createAtImpl_4191D0(*(*uint8)(unsafe.Pointer(uintptr(a1 + 57))), unsafe.Pointer(uintptr(v3)), v2, int32(nox_player_netCode_85319C), 1)
					*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v1))), 0)) = uint8(int8(int32(*memmap.PtrUint8(6112660, 1045696)) + 1))
					*memmap.PtrUint32(6112660, 1045696)++
					return int8(uintptr(unsafe.Pointer(v1)))
				}
				*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v1))), 0)) = uint8(sub_419900(a1, v3, int16(uint16(nox_player_netCode_85319C))))
			}
			*memmap.PtrUint32(6112660, 1045696)++
			return int8(uintptr(unsafe.Pointer(v1)))
		}
	}
	return int8(uintptr(unsafe.Pointer(v1)))
}
func sub_456D00(a1 int32, a2 *libc.WChar) *libc.WChar {
	var (
		v2 *uint32
		v3 *libc.WChar
		v4 *libc.WChar
		v6 [56]libc.WChar
	)
	v2 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045684)))).ChildByID(0x2906)))
	v3 = (*libc.WChar)(unsafe.Pointer(uintptr((*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v2)))))).Func94(asWindowEvent(0x4016, a1, 0)))))
	nox_wcscpy(&v6[0], v3)
	v4 = nox_wcstok(&v6[0], libc.CWString("\t\n\r"))
	return nox_wcscpy(a2, v4)
}
func sub_456D60(a1 int32) *int32 {
	var (
		v1     *int32
		v2     *int32
		result *int32
		v4     *int32
		v5     *int32
	)
	if a1 != 0 && dword_5d4594_1045684 != 0 {
		(*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045684)))).Destroy()
	}
	v1 = (*int32)(unsafe.Pointer(nox_common_list_getFirstSafe_425890((*nox_list_item_t)(unsafe.Pointer(memmap.PtrInt32(6112660, 1045652))))))
	if v1 != nil {
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
	result = (*int32)(unsafe.Pointer(nox_common_list_getFirstSafe_425890((*nox_list_item_t)(unsafe.Pointer(memmap.PtrInt32(6112660, 1045668))))))
	v4 = result
	if result != nil {
		for {
			v5 = (*int32)(unsafe.Pointer(nox_common_list_getNextSafe_4258A0((*nox_list_item_t)(unsafe.Pointer(v4)))))
			sub_425920((**uint32)(unsafe.Pointer(v4)))
			alloc.Free(unsafe.Pointer(v4))
			v4 = v5
			if v5 == nil {
				break
			}
		}
	}
	dword_5d4594_1045684 = 0
	return result
}
func sub_456DF0(a1 int32) int32 {
	var (
		result int32
		v2     int32
		v3     *uint32
	)
	result = int32(dword_5d4594_1045684)
	if dword_5d4594_1045684 != 0 {
		result = sub_456E40(a1, 1)
		v2 = result
		if result != -1 {
			v3 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045684)))).ChildByID(0x2905)))
			result = (*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v3)))))).Func94(asWindowEvent(0x400E, v2, 0))
		}
	}
	return result
}
func sub_456E40(a1 int32, a2 int32) int32 {
	var (
		v2 *int32
		v3 int32
	)
	v2 = (*int32)(unsafe.Pointer(nox_common_list_getFirstSafe_425890((*nox_list_item_t)(unsafe.Pointer(memmap.PtrInt32(6112660, 1045652))))))
	v3 = 0
	if v2 == nil {
		return -1
	}
	for *(*int32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(int32(0))*15)) != a1 {
		v2 = (*int32)(unsafe.Pointer(nox_common_list_getNextSafe_4258A0((*nox_list_item_t)(unsafe.Pointer(v2)))))
		v3++
		if v2 == nil {
			return -1
		}
	}
	if a2 != 0 {
		sub_425920((**uint32)(unsafe.Pointer(v2)))
		alloc.Free(unsafe.Pointer(v2))
	}
	return v3
}
func sub_456EA0(a1 *libc.WChar) int32 {
	var (
		result int32
		v2     *uint32
		v3     int32
		v4     int32
	)
	result = int32(dword_5d4594_1045684)
	if dword_5d4594_1045684 != 0 {
		v2 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045684)))).ChildByID(0x2906)))
		v3 = sub_456F10(a1, 1)
		v4 = v3
		if v3 != -1 {
			sub_4258C0((**uint32)(memmap.PtrOff(6112660, 1045668)), v3)
			(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v2)))))).Func94(asWindowEvent(0x4014, 0, 0))
			(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v2)))))).Func94(asWindowEvent(0x400E, v4, 0))
		}
		result = sub_456500()
	}
	return result
}
func sub_456F10(a1 *libc.WChar, a2 int32) int32 {
	var (
		v2 *int32
		v3 int32
		v4 *libc.WChar
		v6 [56]libc.WChar
	)
	v2 = (*int32)(unsafe.Pointer(nox_common_list_getFirstSafe_425890((*nox_list_item_t)(unsafe.Pointer(memmap.PtrInt32(6112660, 1045668))))))
	v3 = 0
	if v2 == nil {
		return -1
	}
	for {
		nox_wcscpy(&v6[0], (*libc.WChar)(unsafe.Add(unsafe.Pointer((*libc.WChar)(unsafe.Pointer(v2))), unsafe.Sizeof(libc.WChar(0))*6)))
		v4 = nox_wcstok(&v6[0], libc.CWString(" \t\n\r"))
		if _nox_wcsicmp(v4, a1) == 0 {
			break
		}
		v2 = (*int32)(unsafe.Pointer(nox_common_list_getNextSafe_4258A0((*nox_list_item_t)(unsafe.Pointer(v2)))))
		v3++
		if v2 == nil {
			return -1
		}
	}
	if a2 != 0 {
		sub_425920((**uint32)(unsafe.Pointer(v2)))
		alloc.Free(unsafe.Pointer(v2))
	}
	return v3
}
func sub_456FA0() int32 {
	var (
		result int32
		v1     *int32
		v2     *int32
	)
	result = int32(dword_5d4594_1045684)
	if dword_5d4594_1045684 != 0 {
		v1 = (*int32)(unsafe.Pointer(nox_common_list_getFirstSafe_425890((*nox_list_item_t)(unsafe.Pointer(memmap.PtrInt32(6112660, 1045668))))))
		if v1 != nil {
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
		nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045688))))), 0)
		nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045692))))), 0)
		result = sub_456500()
	}
	return result
}
func sub_457120(a1 int32) uint8 {
	return *memmap.PtrUint8(0x587000, uint32((int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 56))))%10)*8)+0x1F7C8)
}
func sub_457140(a1 int32, a2 *libc.WChar) int32 {
	var (
		result int32
		v3     *libc.WChar
		v4     *uint32
	)
	result = int32(dword_5d4594_1045684)
	if dword_5d4594_1045684 != 0 {
		v3 = (*libc.WChar)(alloc.Calloc(1, 72))
		sub_425770(unsafe.Pointer(v3))
		nox_common_list_append_4258E0((*nox_list_item_t)(unsafe.Pointer(uintptr(int32(uintptr(memmap.PtrOff(6112660, 1045652)))))), (*nox_list_item_t)(unsafe.Pointer(v3)))
		nox_wcscpy((*libc.WChar)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(libc.WChar(0))*6)), a2)
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v3))), unsafe.Sizeof(uint32(0))*15))) = uint32(a1)
		v4 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045684)))).ChildByID(0x2905)))
		result = (*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v4)))))).Func94(asWindowEvent(0x400D, int32(uintptr(unsafe.Pointer(a2))), 3))
	}
	return result
}
func sub_4571A0(a1 int32, a2 int32) int32 {
	var (
		result int32
		v3     int32
		v4     int32
		v5     int32
		i      *int32
	)
	result = int32(dword_5d4594_1045684)
	if dword_5d4594_1045684 != 0 {
		v3 = int32(uintptr((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045684)))).ChildByID(0x2905).widget_data))
		result = sub_456E40(a1, 0)
		v4 = result
		if result != -1 {
			v5 = 3
			if a2 != 0 {
				for i = (*int32)(unsafe.Pointer(nox_common_list_getFirstSafe_425890((*nox_list_item_t)(unsafe.Pointer(memmap.PtrInt32(6112660, 1045668)))))); i != nil; i = (*int32)(unsafe.Pointer(nox_common_list_getNextSafe_4258A0((*nox_list_item_t)(unsafe.Pointer(i))))) {
					if *(*int32)(unsafe.Add(unsafe.Pointer(i), unsafe.Sizeof(int32(0))*15)) == a2 {
						v5 = int32(*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(i))), 64))))
					}
				}
			}
			result = int32(*(*uint32)(unsafe.Pointer(uintptr(v3 + 24))))
			*(*uint32)(unsafe.Pointer(uintptr(result + v4*524 + 516))) = **(**uint32)(memmap.PtrOff(0x85B3FC, uint32(v5*4+132)))
		}
	}
	return result
}
func sub_457350(a1 uint8, a2 uint8) *int32 {
	var result *int32
	for result = (*int32)(unsafe.Pointer(nox_common_list_getFirstSafe_425890((*nox_list_item_t)(unsafe.Pointer(memmap.PtrInt32(6112660, 1045668)))))); result != nil; result = (*int32)(unsafe.Pointer(nox_common_list_getNextSafe_4258A0((*nox_list_item_t)(unsafe.Pointer(result))))) {
		if *(*int32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(int32(0))*15)) == int32(a1) {
			*(*int32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(int32(0))*15)) = int32(a2)
		}
	}
	return result
}
func sub_4573A0() int32 {
	var result int32
	result = int32(dword_5d4594_1045684)
	if dword_5d4594_1045684 != 0 {
		result = sub_456500()
	}
	return result
}
func sub_4573B0() {
	*memmap.PtrUint32(6112660, 1045696) = 0
}
func sub_457460(a1 int32) int32 {
	var WideCharStr [16]libc.WChar
	compat_itow(int32(*(*uint16)(unsafe.Pointer(uintptr(a1 + 54)))), &WideCharStr[0], 10)
	(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1046516))))).Func94(asWindowEvent(0x401E, int32(uintptr(unsafe.Pointer(&WideCharStr[0]))), 0))
	compat_itow(int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 56)))), &WideCharStr[0], 10)
	return (*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1046520))))).Func94(asWindowEvent(0x401E, int32(uintptr(unsafe.Pointer(&WideCharStr[0]))), 0))
}
func nox_xxx_guiServerOptionsTryHide_4574D0() int32 {
	if dword_5d4594_1046492 == 0 {
		return 0
	}
	nox_xxx_guiServerOptionsHide_4597E0(0)
	dword_5d4594_1046492 = 0
	return 1
}
func sub_457B60(a1 int32) int32 {
	var (
		v1          *byte
		v2          *uint32
		v3          *byte
		v4          uint16
		v5          *byte
		v6          *uint32
		v7          uint16
		v8          uint8
		v9          *uint32
		v10         *uint32
		v11         *uint32
		v12         int32
		v13         *uint32
		v14         int32
		v15         *uint32
		v16         int32
		v17         uint32
		WideCharStr [16]libc.WChar
		v20         [100]libc.WChar
	)
	v1 = nox_xxx_serverOptionsGetServername_40A4C0()
	nox_swprintf(&v20[0], libc.CWString("%S"), v1)
	dword_5d4594_1046512.Func94(asWindowEvent(0x401E, int32(uintptr(unsafe.Pointer(&v20[0]))), 0))
	sub_459CD0()
	if noxflags.HasGame(noxflags.GameModeChat) {
		v2 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1046492)))).ChildByID(0x278A)))
	} else {
		v2 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1046492)))).ChildByID(0x2787)))
	}
	(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v2)))))).Hide()
	if noxflags.HasGame(noxflags.GameHost) {
		nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1046508))))), 1)
		if noxflags.HasGame(noxflags.GameModeChat) {
			v3 = nox_xxx_cliGamedataGet_416590(1)
			v4 = uint16(int16(int32(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v3))), unsafe.Sizeof(uint16(0))*26)))) & 6128))
			nox_client_guiserv_updateMapList_458230(int32(v4), v3, true)
			v5 = sub_4165B0()
			*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v5))), unsafe.Sizeof(uint16(0))*26))) = uint16(int16(int32(v4) | (int32(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v5))), unsafe.Sizeof(uint16(0))*26)))) & 0xE80F)))
			if sub_40A740() != 0 {
				nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1046508))))), 0)
			} else if nox_xxx_CheckGameplayFlags_417DA0(4) {
				v6 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1046508)))).ChildByID(10330)))
				*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*9)) |= 4
			} else {
				sub_46AD20(*(**uint32)(unsafe.Pointer(&dword_5d4594_1046508)), 0x285B, 0x285D, 0)
			}
		} else {
			if nox_xxx_CheckGameplayFlags_417DA0(4) {
				sub_46AD20(*(**uint32)(unsafe.Pointer(&dword_5d4594_1046508)), 10330, 0x285B, 0)
			} else {
				sub_46AD20(*(**uint32)(unsafe.Pointer(&dword_5d4594_1046508)), 10330, 0x285D, 0)
			}
			v5 = (*byte)(unsafe.Pointer(uintptr(a1)))
			nox_client_guiserv_updateMapList_458230(int32(*(*uint16)(unsafe.Pointer(uintptr(a1 + 52)))), (*byte)(unsafe.Pointer(uintptr(a1))), false)
		}
		v7 = uint16(nox_xxx_servGamedataGet_40A020(int16(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v5))), unsafe.Sizeof(uint16(0))*26))))))
		compat_itow(int32(v7), &WideCharStr[0], 10)
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1046516))))).Func94(asWindowEvent(0x401E, int32(uintptr(unsafe.Pointer(&WideCharStr[0]))), 0))
		v8 = sub_40A180(int16(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v5))), unsafe.Sizeof(uint16(0))*26)))))
		compat_itow(int32(v8), &WideCharStr[0], 10)
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1046520))))).Func94(asWindowEvent(0x401E, int32(uintptr(unsafe.Pointer(&WideCharStr[0]))), 0))
		sub_4580E0(int32(uintptr(unsafe.Pointer(v5))))
	} else {
		v9 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1046492)))).ChildByID(0x27B1)))
		nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v9)))))), 0)
		nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1046536))))), 0)
		*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1046536 + 4))) |= 8
		nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1046504))))), 1)
		sub_46AD20(*(**uint32)(unsafe.Pointer(&dword_5d4594_1046504)), 0x2796, 0x2797, 0)
		sub_46ACE0(*(**uint32)(unsafe.Pointer(&dword_5d4594_1046492)), 0x279D, 0x279D, 1)
		v5 = (*byte)(unsafe.Pointer(uintptr(a1)))
		nox_client_guiserv_updateMapList_458230(int32(*(*uint16)(unsafe.Pointer(uintptr(a1 + 52)))), (*byte)(unsafe.Pointer(uintptr(a1))), false)
		compat_itow(int32(*(*uint16)(unsafe.Pointer(uintptr(a1 + 54)))), &WideCharStr[0], 10)
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1046516))))).Func94(asWindowEvent(0x401E, int32(uintptr(unsafe.Pointer(&WideCharStr[0]))), 0))
		compat_itow(int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 56)))), &WideCharStr[0], 10)
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1046520))))).Func94(asWindowEvent(0x401E, int32(uintptr(unsafe.Pointer(&WideCharStr[0]))), 0))
		if nox_xxx_check_flag_aaa_43AF70() == 1 {
			v10 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1046492)))).ChildByID(10210)))
			(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v10)))))).Show()
			sub_457FE0()
		}
	}
	v11 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1046508)))).ChildByID(0x285B)))
	if nox_xxx_CheckGameplayFlags_417DA0(2) {
		v12 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v11), unsafe.Sizeof(uint32(0))*9)))
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v12))), 0)) = uint8(int8(v12 | 4))
		*(*uint32)(unsafe.Add(unsafe.Pointer(v11), unsafe.Sizeof(uint32(0))*9)) = uint32(v12)
	}
	v13 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1046508)))).ChildByID(0x285D)))
	if nox_xxx_CheckGameplayFlags_417DA0(1) {
		v14 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v13), unsafe.Sizeof(uint32(0))*9)))
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v14))), 0)) = uint8(int8(v14 | 4))
		*(*uint32)(unsafe.Add(unsafe.Pointer(v13), unsafe.Sizeof(uint32(0))*9)) = uint32(v14)
	}
	v15 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1046492)))).ChildByID(0x2787)))
	v16 = int32(uintptr(unsafe.Pointer(nox_xxx_guiServerOptionsGetGametypeName_4573C0(int16(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v5))), unsafe.Sizeof(uint16(0))*26))))))))
	(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v15)))))).Func94(asWindowEvent(0x4001, v16, 0))
	sub_459880(int32(uintptr(unsafe.Pointer(v5))))
	*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v17))), unsafe.Sizeof(uint16(0))*0)) = *((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v5))), unsafe.Sizeof(uint16(0))*26)))
	return sub_457F30(int32((v17 >> 12) & 1))
}
func sub_457F30(a1 int32) int32 {
	var (
		v1 int32
		v2 *uint32
		v3 *uint32
		v4 *uint32
		v5 *uint32
	)
	v1 = bool2int(a1 != 1)
	v2 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1046504)))).ChildByID(0x27A8)))
	nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v2)))))), v1)
	v3 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1046504)))).ChildByID(0x279D)))
	nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v3)))))), v1)
	nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1046516))))), v1)
	nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1046520))))), v1)
	if a1 == 1 {
		v4 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1046492)))).ChildByID(0x278A)))
		if (*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v4)))))).Flags().IsHidden() == 0 {
			(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v4)))))).Hide()
		}
	}
	v5 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1046492)))).ChildByID(0x27C7)))
	return (*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v5)))))).SetHidden(a1)
}
func nox_xxx_windowServerOptionsDrawProc_458500(a1 *uint32, a2 int32) int32 {
	var (
		v2    *uint32
		xLeft int32
		v5    int32
		mpos  nox_point = getMousePos()
	)
	nox_client_wndGetPosition_46AA60((*nox_window)(unsafe.Pointer(a1)), (*uint32)(unsafe.Pointer(&xLeft)), (*uint32)(unsafe.Pointer(&v5)))
	nox_client_drawRectFilledAlpha_49CF10(xLeft, v5+25, int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*2))), int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*3))-25))
	v2 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1046492)))).ChildByID(10120)))
	if (*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v2)))))).Flags().IsHidden() == 0 && !nox_xxx_wndPointInWnd_46AAB0(v2, mpos.x, mpos.y) {
		nox_xxx_wndClearCaptureMain((*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v2)))))))
		(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v2)))))).Hide()
	}
	return 1
}
func nox_xxx_guiServerOptionsProc_458590(a1 int32, a2 int32, a3 int32, a4 int32) int32 {
	if a2 == 21 {
		if a3 != 1 {
			return 0
		}
		if a4 == 2 {
			clientPlaySoundSpecial(231, 100)
			nox_xxx_guiServerOptionsHide_4597E0(0)
		}
	}
	return 1
}
func sub_459150() int8 {
	var (
		v0  *byte
		i   int32
		v2  *byte
		v3  int32
		v4  *libc.WChar
		v5  *byte
		v6  *byte
		v7  *byte
		v8  int32
		v9  *byte
		v10 uint8
		v11 *byte
	)
	_ = v11
	var v12 *byte
	var v13 *byte
	var v14 *int32
	var v16 *int32
	var v17 *int32
	var v18 int16
	var v19 [16]byte
	v0 = sub_4165B0()
	if *(*byte)(unsafe.Add(unsafe.Pointer(v0), 52))&96 != 0 {
		for i = int32(nox_xxx_getTeamCounter_417DD0()); i < 2; i++ {
			v2 = (*byte)(unsafe.Pointer(nox_xxx_teamCreate_4186D0(0)))
			v3 = bool2int(i != 0) + 1
			if sub_40A740() != 0 {
				sub_418800((*libc.WChar)(unsafe.Pointer(v2)), (*libc.WChar)(memmap.PtrOff(6112660, 1046564)), 0)
			} else {
				v4 = nox_server_teamTitle_418C20(v3)
				sub_418800((*libc.WChar)(unsafe.Pointer(v2)), v4, 1)
			}
			*(*byte)(unsafe.Add(unsafe.Pointer(v2), 56)) = byte(int8(v3))
			sub_4184D0((*nox_team_t)(unsafe.Pointer((*libc.WChar)(unsafe.Pointer(v2)))))
		}
		nox_xxx_SetGameplayFlag_417D50(4)
	}
	*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v5))), 0)) = uint8(*(*byte)(unsafe.Add(unsafe.Pointer(v0), 57)))
	if int32(uint8(uintptr(unsafe.Pointer(v5)))) != 0 {
		v5 = sub_409B80()
		v6 = v5
	} else {
		v6 = v0
	}
	if *v6 != 0 {
		v7 = nox_xxx_mapGetMapName_409B40()
		v8 = int32(libc.StrCaseCmp(v6, v7))
		sub_4165F0(1, 0)
		v9 = sub_4165D0(0)
		nox_xxx_gameSetServername_40A440((*byte)(unsafe.Add(unsafe.Pointer(v9), 9)))
		if ((int32(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v9))), unsafe.Sizeof(uint16(0))*26)))) >> 8) & 16) == 0 {
			sub_409FB0_settings(int16(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v9))), unsafe.Sizeof(uint16(0))*26)))), *((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v9))), unsafe.Sizeof(uint16(0))*27))))
			sub_40A040_settings(int16(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v9))), unsafe.Sizeof(uint16(0))*26)))), uint8(*(*byte)(unsafe.Add(unsafe.Pointer(v9), 56))))
		}
		if v8 != 0 && nox_xxx_check_flag_aaa_43AF70() == 1 {
			if !noxflags.HasGame(noxflags.GameModeChat) {
				nox_xxx_net_4263C0()
				sub_40DF90()
				sub_4264D0()
				sub_416150(15, 0)
			}
			if *(*byte)(unsafe.Add(unsafe.Pointer(v9), 57)) == 0 {
				sub_41D650()
			}
		}
		if noxflags.HasGame(noxflags.GameModeChat) {
			noxflags.UnsetGame(noxflags.GameFlag15 | noxflags.GameFlag16)
			noxflags.SetGame(uint32(int32(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v9))), unsafe.Sizeof(uint16(0))*26)))) & 0xC000))
		}
		*(*byte)(unsafe.Add(unsafe.Pointer(v9), 57)) = 0
		if v8 != 0 {
			v19[0] = 0
			*(*uint32)(unsafe.Pointer(&v19[1])) = 0
			*(*uint32)(unsafe.Pointer(&v19[5])) = 0
			*(*uint32)(unsafe.Pointer(&v19[9])) = 0
			v10 = *memmap.PtrUint8(0x587000, 131672)
			libc.StrCpy(&v19[0], v6)
			v11 = &v19[libc.StrLen(&v19[0])]
			*(*uint32)(unsafe.Pointer(v11)) = *memmap.PtrUint32(0x587000, 131668)
			*(*byte)(unsafe.Add(unsafe.Pointer(v11), 4)) = byte(v10)
			nox_xxx_mapLoad_4D2450(&v19[0])
			sub_416690()
			sub_4165D0(1)
			*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v5))), 0)) = uint8(uint32(uintptr(unsafe.Pointer(nox_xxx_guiServerOptionsHide_4597E0(0)))))
		} else {
			nox_xxx_spellEnableAll_424BD0()
			sub_4537F0()
			v17 = (*int32)(unsafe.Pointer(sub_459870()))
			v12 = nox_xxx_cliGamedataGet_416590(1)
			sub_57AAA0(CString("user.rul"), v12, v17)
			v13 = nox_server_currentMapGetFilename_409B30()
			sub_57A950(v13)
			v18 = int16(uint16(nox_common_gameFlags_getVal_40A5B0()))
			v16 = (*int32)(unsafe.Pointer(sub_459870()))
			v14 = (*int32)(unsafe.Pointer(nox_xxx_cliGamedataGet_416590(0)))
			sub_57A1E0(v14, CString("user.rul"), v16, 3, v18)
			*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v5))), 0)) = uint8(uint32(uintptr(unsafe.Pointer(nox_xxx_guiServerOptionsHide_4597E0(0)))))
		}
	}
	return int8(uintptr(unsafe.Pointer(v5)))
}
func sub_4593B0(a1 int32) int32 {
	var (
		v1     *uint32
		v2     *uint32
		v3     *uint32
		result int32
	)
	v1 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1046492)))).ChildByID(0x27B3)))
	v2 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1046492)))).ChildByID(0x27B1)))
	v3 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1046492)))).ChildByID(0x27B2)))
	if a1 != 0 {
		if a1 == 1 {
			*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*15)) = dword_5d4594_1046360
			*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*15)) = dword_5d4594_1046356
			if dword_5d4594_1046532 != 0 {
				sub_456D60(1)
				dword_5d4594_1046532 = 0
			}
			if dword_5d4594_1046540 != 0 {
				sub_4AD820()
				dword_5d4594_1046540 = 0
			}
			dword_5d4594_1046528 = uint32(nox_xxx_guiServerAccessLoad_4541D0(*(*int32)(unsafe.Pointer(&dword_5d4594_1046492))))
			result = sub_459560(0x27B1)
		} else {
			result = a1 - 2
			if a1 == 2 {
				*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*15)) = dword_5d4594_1046360
				if noxflags.HasGame(noxflags.GameHost) {
					*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*15)) = dword_5d4594_1046360
				} else {
					*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*21)) = dword_5d4594_1046360
				}
				if dword_5d4594_1046528 != 0 {
					sub_4557D0(1)
					dword_5d4594_1046528 = 0
				}
				if dword_5d4594_1046540 != 0 {
					sub_4AD820()
					dword_5d4594_1046540 = 0
				}
				dword_5d4594_1046532 = uint32(nox_xxx_guiServerPlayersLoad_456270(*(*int32)(unsafe.Pointer(&dword_5d4594_1046492))))
				result = sub_459560(0x27B2)
			}
		}
	} else {
		if noxflags.HasGame(noxflags.GameHost) {
			*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*15)) = dword_5d4594_1046356
		} else {
			*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*21)) = dword_5d4594_1046356
		}
		*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*15)) = dword_5d4594_1046356
		if dword_5d4594_1046532 != 0 {
			sub_456D60(1)
			dword_5d4594_1046532 = 0
		} else if dword_5d4594_1046528 != 0 {
			sub_4557D0(1)
			dword_5d4594_1046528 = 0
		}
		dword_5d4594_1046540 = uint32(nox_xxx_gui_4AD320(*(*int32)(unsafe.Pointer(&dword_5d4594_1046492))))
		result = sub_459560(0x27B3)
	}
	return result
}
func sub_459560(a1 int32) int32 {
	var (
		v1 *uint32
		v2 int32
		v3 *uint32
		v4 *uint32
		v5 *uint32
		v6 *uint32
		v7 *uint32
	)
	v1 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1046492)))).ChildByID(a1)))
	v2 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*99)))
	switch a1 {
	case 0x27B1:
		v6 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1046492)))).ChildByID(0x27B3)))
		sub_46B120((*nox_window)(unsafe.Pointer(v6)), nil)
		sub_46B120((*nox_window)(unsafe.Pointer(v6)), (*nox_window)(unsafe.Pointer(uintptr(v2))))
		v4 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1046492)))).ChildByID(0x27B2)))
		goto LABEL_7
	case 0x27B2:
		v5 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1046492)))).ChildByID(0x27B3)))
		sub_46B120((*nox_window)(unsafe.Pointer(v5)), nil)
		sub_46B120((*nox_window)(unsafe.Pointer(v5)), (*nox_window)(unsafe.Pointer(uintptr(v2))))
		v4 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1046492)))).ChildByID(0x27B1)))
		goto LABEL_7
	case 0x27B3:
		v3 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1046492)))).ChildByID(0x27B2)))
		sub_46B120((*nox_window)(unsafe.Pointer(v3)), nil)
		sub_46B120((*nox_window)(unsafe.Pointer(v3)), (*nox_window)(unsafe.Pointer(uintptr(v2))))
		v4 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1046492)))).ChildByID(0x27B1)))
	LABEL_7:
		v7 = v4
		sub_46B120((*nox_window)(unsafe.Pointer(v4)), nil)
		sub_46B120((*nox_window)(unsafe.Pointer(v7)), (*nox_window)(unsafe.Pointer(uintptr(v2))))
	}
	sub_46B120((*nox_window)(unsafe.Pointer(v1)), nil)
	return sub_46B120((*nox_window)(unsafe.Pointer(v1)), (*nox_window)(unsafe.Pointer(uintptr(v2))))
}
func sub_459700() int32 {
	var (
		v0 *byte
		v1 int32
		v2 int32
		v3 *libc.WChar
		v5 [256]libc.WChar = [256]libc.WChar{}
	)
	sub_416580()
	v0 = sub_4165B0()
	v1 = (*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1046496))))).Func94(asWindowEvent(0x4014, 0, 0))
	v2 = v1
	v3 = (*libc.WChar)(unsafe.Pointer(uintptr((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1046496))))).Func94(asWindowEvent(0x4016, v1, 0)))))
	if v3 != nil {
		nox_wcsncpy(&v5[0], v3, math.MaxUint8)
		v5[math.MaxUint8] = 0
	} else {
		v5[0] = 0
	}
	(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1046496))))).Func94(asWindowEvent(0x400E, v2, 0))
	(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1046496))))).Func94(asWindowEvent(0x4012, v2, 0))
	(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1046496))))).Func94(asWindowEvent(0x400D, int32(uintptr(unsafe.Pointer(&v5[0]))), -1))
	(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1046496))))).Func94(asWindowEvent(0x4013, v2, 0))
	sub_57A9F0(v0, CString("user.rul"))
	sub_57A1E0((*int32)(unsafe.Pointer(v0)), CString("user.rul"), nil, 5, int16(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v0))), unsafe.Sizeof(uint16(0))*26)))))
	sub_459880(int32(uintptr(unsafe.Pointer(v0))))
	return sub_459D50(1)
}
func nox_xxx_guiServerOptionsHide_4597E0(a1 int32) *int32 {
	var (
		v1     int32
		result *int32
	)
	if dword_5d4594_1046492 != 0 {
		v1 = int32(uintptr(unsafe.Pointer(nox_xxx_wndGetFocus_46B4F0())))
		if nox_window_is_child((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1046492))))), (*nox_window)(unsafe.Pointer(uintptr(v1)))) != 0 {
			guiFocus(nil)
		}
		(*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1046492)))).Destroy()
		dword_5d4594_1046492 = 0
		sub_456D60(0)
		sub_4BE610()
		dword_5d4594_1046532 = 0
		sub_4557D0(0)
		dword_5d4594_1046528 = 0
		sub_4AD820()
		dword_5d4594_1046536 = 0
	}
	result = (*int32)(unsafe.Pointer(uintptr(a1)))
	if a1 != 0 {
		result = sub_57ADF0(memmap.PtrInt32(6112660, 1045956))
		dword_587000_129656 = 1
	}
	return result
}
func sub_459870() *byte {
	return (*byte)(memmap.PtrOff(6112660, 1045956))
}
func sub_459A40(a1 *byte) int32 {
	var (
		v2 [100]byte
		v3 [100]libc.WChar
	)
	libc.StrNCpy(&v2[0], a1, 15)
	v2[15] = 0
	nox_swprintf(&v3[0], libc.CWString("%S"), &v2[0])
	return dword_5d4594_1046512.Func94(asWindowEvent(0x401E, int32(uintptr(unsafe.Pointer(&v3[0]))), 0))
}
func sub_459AA0(a1p unsafe.Pointer) *byte {
	var (
		a1     int32 = int32(uintptr(a1p))
		v1     int32
		v2     *libc.WChar
		v3     *libc.WChar
		result *byte
		v5     int32
		v6     *byte
		v7     [100]byte
	)
	v1 = dword_5d4594_1046512.Func94(asWindowEvent(0x401D, 0, 0))
	nox_sprintf(&v7[0], CString("%S"), v1)
	libc.StrNCpy((*byte)(unsafe.Pointer(uintptr(a1+9))), &v7[0], 15)
	*(*uint16)(unsafe.Pointer(uintptr(a1 + 52))) = uint16(int16(sub_459C10()))
	libc.MemCpy(unsafe.Pointer(uintptr(a1+24)), unsafe.Pointer(sub_453F90()), 20)
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 44))) = *(*uint32)(unsafe.Pointer(sub_453600()))
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 48))) = uint32(sub_453610())
	v2 = (*libc.WChar)(unsafe.Pointer(uintptr((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1046516))))).Func94(asWindowEvent(0x401D, 0, 0)))))
	if *v2 != 0 {
		*(*uint16)(unsafe.Pointer(uintptr(a1 + 54))) = uint16(int16(nox_wcstol(v2, nil, 10)))
	}
	v3 = (*libc.WChar)(unsafe.Pointer(uintptr((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1046520))))).Func94(asWindowEvent(0x401D, 0, 0)))))
	if *v3 != 0 {
		*(*uint8)(unsafe.Pointer(uintptr(a1 + 56))) = uint8(int8(nox_wcstol(v3, nil, 10)))
	}
	*(*uint8)(unsafe.Pointer(uintptr(a1 + 57))) = uint8(((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1046492)))).ChildByID(0x278A).draw_data.field_0 >> 2) & 1)
	result = (*byte)(unsafe.Pointer(uintptr((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1046496))))).Func94(asWindowEvent(0x4014, 0, 0)))))
	if int32(uintptr(unsafe.Pointer(result))) >= 0 && (func() bool {
		v5 = (*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1046496))))).Func94(asWindowEvent(0x4016, int32(uintptr(unsafe.Pointer(result))), 0))
		nox_sprintf(&v7[0], CString("%S"), v5)
		return (func() *byte {
			result = libc.StrTok(&v7[0], CString("\t"))
			return result
		}()) != nil
	}()) {
		v6 = result
		result = nil
		libc.StrCpy((*byte)(unsafe.Pointer(uintptr(a1))), v6)
	} else {
		*(*uint8)(unsafe.Pointer(uintptr(a1))) = 0
	}
	return result
}
func sub_459C30() int32 {
	var (
		v0     *byte
		result int32
		v2     *uint32
		v3     int32
	)
	v0 = nox_xxx_cliGamedataGet_416590(1)
	sub_453F70(unsafe.Add(unsafe.Pointer(v0), 24))
	sub_4535E0((*int32)(unsafe.Add(unsafe.Pointer((*int32)(unsafe.Pointer(v0))), unsafe.Sizeof(int32(0))*11)))
	sub_4535F0(int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v0))), unsafe.Sizeof(uint32(0))*12)))))
	result = int32(dword_5d4594_1046492)
	if dword_5d4594_1046492 != 0 {
		sub_459880(int32(uintptr(unsafe.Pointer(v0))))
		sub_4BDF70((*int32)(unsafe.Pointer(v0)))
		nox_client_guiserv_updateMapList_458230(int32(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v0))), unsafe.Sizeof(uint16(0))*26)))), v0, false)
		sub_457460(int32(uintptr(unsafe.Pointer(v0))))
		if nox_xxx_check_flag_aaa_43AF70() == 1 {
			sub_457FE0()
		}
		v2 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1046492)))).ChildByID(0x2787)))
		v3 = int32(uintptr(unsafe.Pointer(nox_xxx_guiServerOptionsGetGametypeName_4573C0(int16(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v0))), unsafe.Sizeof(uint16(0))*26))))))))
		result = (*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v2)))))).Func94(asWindowEvent(0x4001, v3, 0))
	}
	return result
}
func sub_459D50(a1 int32) int32 {
	var result int32
	result = a1
	*memmap.PtrUint32(6112660, 1046544) = uint32(a1)
	return result
}
func sub_459D60() int32 {
	return int32(*memmap.PtrUint32(6112660, 1046544))
}
func sub_459D70() int32 {
	var v0 int32
	v0 = -bool2int(dword_5d4594_1046492 != 0)
	*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v0))), 0)) = uint8(int8(v0 & 254))
	return v0 + 2
}
func sub_459D80(a1 int32) int32 {
	var result int32
	result = int32(dword_5d4594_1046492)
	if dword_5d4594_1046492 != 0 {
		result = (*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1046492))))).SetHidden(a1)
	}
	return result
}
func sub_459DA0() int32 {
	return bool2int(dword_5d4594_1046492 != 0)
}
func sub_459DB0(dr *nox_drawable) int32 {
	var a1 int32 = int32(uintptr(unsafe.Pointer(dr)))
	return bool2int(*(*uint32)(unsafe.Pointer(uintptr(a1 + 112)))&0x400000 != 0 && int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 116))))&8 != 0)
}
func sub_459DD0(dr *nox_drawable, a2 int8) {
	var a1 int32 = int32(uintptr(unsafe.Pointer(dr)))
	if a1 == 0 {
		return
	}
	var v3 int32 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 120))))
	*(*uint8)(unsafe.Pointer(uintptr(a1 + 284))) |= uint8(a2)
	if v3 < 0 {
		return
	}
	for it := int32(nox_xxx_cliFirstMinimapObj_459EB0()); it != 0; it = nox_xxx_cliNextMinimapObj_459EC0(it) {
		if a1 == it {
			return
		}
	}
	var v4 int32 = int32(dword_5d4594_1046596)
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 412))) = 0
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 408))) = uint32(v4)
	if dword_5d4594_1046596 != 0 {
		*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1046596 + 412))) = uint32(a1)
	}
	dword_5d4594_1046596 = uint32(a1)
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 120))) |= 0x80000000
}
func nox_xxx_cliRemoveHealthbar_459E30(dr *nox_drawable, a2 int8) int32 {
	var (
		a1     int32 = int32(uintptr(unsafe.Pointer(dr)))
		result int32
		v3     bool
		v4     int32
		v5     int32
	)
	result = a1
	if *(*int32)(unsafe.Pointer(uintptr(a1 + 120))) < 0 {
		v3 = (int32(uint8(int8(int32(^a2)))) & int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 284))))) == 0
		*(*uint8)(unsafe.Pointer(uintptr(a1 + 284))) &= uint8(int8(int32(^a2)))
		if v3 {
			v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 408))))
			if v4 != 0 {
				*(*uint32)(unsafe.Pointer(uintptr(v4 + 412))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 412)))
			}
			v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 412))))
			if v5 != 0 {
				*(*uint32)(unsafe.Pointer(uintptr(v5 + 408))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 408)))
			} else {
				dword_5d4594_1046596 = *(*uint32)(unsafe.Pointer(uintptr(a1 + 408)))
			}
			*(*uint32)(unsafe.Pointer(uintptr(a1 + 120))) &= math.MaxInt32
		}
	}
	return result
}
func nox_xxx_cliFirstMinimapObj_459EB0() int32 {
	return int32(dword_5d4594_1046596)
}
func nox_xxx_cliNextMinimapObj_459EC0(a1 int32) int32 {
	var next int32 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 408))))
	if a1 != 0 && a1 == next {
		stdio.Printf("nox_xxx_cliNextMinimapObj_459EC0: infinite loop!\n")
		panic("abort")
		return 0
	}
	return next
}
func sub_459ED0_drawable(dr *nox_drawable) {
	dr.field_104 = nox_xxx_drawablePlayer_1046600
	dr.field_105 = nil
	if nox_xxx_drawablePlayer_1046600 != nil {
		nox_xxx_drawablePlayer_1046600.field_105 = dr
	}
	nox_xxx_drawablePlayer_1046600 = dr
}
func sub_459F00(dr *nox_drawable) int32 {
	var (
		a1     int32 = int32(uintptr(unsafe.Pointer(dr)))
		result int32
		v2     int32
		v3     int32
	)
	result = a1
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 416))))
	if v2 != 0 {
		*(*uint32)(unsafe.Pointer(uintptr(v2 + 420))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 420)))
	}
	v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 420))))
	if v3 != 0 {
		result = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 416))))
		*(*uint32)(unsafe.Pointer(uintptr(v3 + 416))) = uint32(result)
	} else {
		nox_xxx_drawablePlayer_1046600 = (*nox_drawable)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 416))))))
	}
	return result
}
func sub_459F40_drawable(a1 int32) int32 {
	var result int32
	result = a1
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 424))) = dword_5d4594_1046576
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 428))) = 0
	if dword_5d4594_1046576 != 0 {
		*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1046576 + 428))) = uint32(a1)
	}
	dword_5d4594_1046576 = uint32(a1)
	return result
}
func sub_459F70(dr *nox_drawable) *uint32 {
	var (
		a1     *uint32 = &dr.field_0
		result *uint32 = nil
		v2     int32
		v3     int32
	)
	result = a1
	v2 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*106)))
	if v2 != 0 {
		*(*uint32)(unsafe.Pointer(uintptr(v2 + 428))) = *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*107))
	} else if *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*107)) == 0 && *(**uint32)(unsafe.Pointer(&dword_5d4594_1046576)) != a1 {
		return result
	}
	v3 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*107)))
	if v3 != 0 {
		*(*uint32)(unsafe.Pointer(uintptr(v3 + 424))) = *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*106))
	} else {
		dword_5d4594_1046576 = *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*106))
	}
	*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*106)) = 0
	*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*107)) = 0
	result = (*uint32)(unsafe.Pointer(uintptr(sub_452EB0((*int32)(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*124))))))))
	if result != nil {
		result = (*uint32)(unsafe.Pointer(uintptr(sub_4523D0(result))))
	}
	return result
}
func nox_xxx_cliGetSpritePlayer_45A000() *nox_drawable {
	return nox_xxx_drawablePlayer_1046600
}
func sub_45A010(dr *nox_drawable) *nox_drawable {
	return dr.field_104
}
func sub_45A040(a1 int32) int32 {
	var result int32
	if a1 != 0 {
		result = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 360))))
	} else {
		result = 0
	}
	return result
}
func sub_45A060() int32 {
	return int32(uintptr(unsafe.Pointer(nox_drawable_head_unk1)))
}
func sub_45A070(a1 int32) int32 {
	var result int32
	if a1 != 0 {
		result = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 368))))
	} else {
		result = 0
	}
	return result
}
func sub_45A090() int32 {
	return int32(dword_5d4594_1046576)
}
func sub_45A0A0(a1 int32) int32 {
	var result int32
	if a1 != 0 {
		result = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 424))))
	} else {
		result = 0
	}
	return result
}
func sub_45A0C0() int32 {
	var result int32
	result = int32(uintptr(unsafe.Pointer(nox_drawable_head_unk1)))
	if result == 0 {
		return 0
	}
	for (*(*uint32)(unsafe.Pointer(uintptr(result + 120))) & 0x2000000) == 0 {
		result = int32(*(*uint32)(unsafe.Pointer(uintptr(result + 368))))
		if result == 0 {
			return 0
		}
	}
	return result
}
func sub_45A0E0(a1 int32) int32 {
	var result int32
	result = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 368))))
	if result == 0 {
		return 0
	}
	for (*(*uint32)(unsafe.Pointer(uintptr(result + 120))) & 0x2000000) == 0 {
		result = int32(*(*uint32)(unsafe.Pointer(uintptr(result + 368))))
		if result == 0 {
			return 0
		}
	}
	return result
}
func nox_xxx_sprite_45A110_drawable(a1 *nox_drawable) *uint32 {
	var result *uint32
	result = &a1.field_0
	a1.field_98 = nil
	a1.field_97 = nox_drawable_head_unk3
	if nox_drawable_head_unk3 != nil {
		nox_drawable_head_unk3.field_98 = a1
	} else {
		nox_drawable_head_unk4 = a1
	}
	nox_drawable_head_unk3 = a1
	a1.flags30 |= 0x400000
	return result
}
func sub_45A160_drawable(a1 *nox_drawable) {
	if (a1.flags30 & 0x400000) == 0 {
		return
	}
	var v2 *nox_drawable = a1.field_98
	if v2 != nil {
		v2.field_97 = a1.field_97
	} else {
		nox_drawable_head_unk3 = a1.field_97
	}
	var v3 *nox_drawable = a1.field_97
	if v3 != nil {
		v3.field_98 = a1.field_98
	} else {
		nox_drawable_head_unk4 = a1.field_98
	}
	a1.flags30 &= 0xFFBFFFFF
}
func nox_xxx_spriteLoadAdd_45A360_drawable(thingInd int32, a2 int32, a3 int32) *nox_drawable {
	var dr *nox_drawable = nox_new_drawable_for_thing(thingInd)
	if dr == nil {
		return nil
	}
	sub_452E60(&dr.flags31)
	if dr.field_116 != 0 {
		nox_xxx_spriteToList_49BC80_drawable(&dr.field_0)
	}
	if dr.flags30&0x200000 != 0 {
		nox_xxx_spriteToSightDestroyList_49BAB0_drawable(&dr.field_0)
	}
	if dr.field_123 != 0 {
		sub_459F40_drawable(int32(uintptr(unsafe.Pointer(dr))))
	}
	dr.pos.x = a2
	dr.field_8 = uint32(a2)
	dr.pos.y = a3
	dr.field_9 = uint32(a3)
	dr.field_80 = nox_frame_xxx_2598000
	dr.field_92 = nox_drawable_head_unk1
	dr.field_93 = nil
	if nox_drawable_head_unk1 != nil {
		nox_drawable_head_unk1.field_93 = dr
	}
	nox_drawable_head_unk1 = dr
	nox_xxx_sprite_49AA00_drawable(dr)
	if dr.flags30&0x10000 != 0 {
		var v6 *nox_drawable = nox_drawable_head_unk2
		dr.field_91 = nil
		dr.field_90 = v6
		if v6 != nil {
			v6.field_91 = dr
		}
		nox_drawable_head_unk2 = dr
	}
	if int32(*(*uint8)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(dr))), 112)))))&4 != 0 {
		sub_459ED0_drawable(dr)
	}
	dr.flags30 |= 0x1000000
	nox_xxx_spriteSetActiveMB_45A990_drawable(int32(uintptr(unsafe.Pointer(dr))))
	dr.field_120 = 0
	dr.field_121 = 0
	nox_xxx_sprite_45A480_drawable(int32(uintptr(unsafe.Pointer(dr))))
	return dr
}
func nox_xxx_sprite_45A480_drawable(a1 int32) {
	var v1 int32
	if *(*uint32)(unsafe.Pointer(uintptr(a1 + 112)))&0x1000000 != 0 && int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 116))))&192 != 0 {
		v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 120))))
		if v1&0x4000 != 0 {
			sub_495F70(a1)
		}
	}
}
func nox_xxx_spriteDeleteStatic_45A4E0_drawable(dr *nox_drawable) {
	if dr.field_93 != nil {
		dr.field_93.field_92 = dr.field_92
	} else {
		nox_drawable_head_unk1 = dr.field_92
	}
	if dr.field_92 != nil {
		dr.field_92.field_93 = dr.field_93
	}
	nox_xxx_sprite_49A9B0_drawable(dr)
	nox_xxx_clientDeleteSprite_476F10_drawable(dr)
	if dr.flags30&0x10000 != 0 {
		if dr.field_91 != nil {
			dr.field_91.field_90 = dr.field_90
		} else {
			nox_drawable_head_unk2 = dr.field_90
		}
		if dr.field_90 != nil {
			dr.field_90.field_91 = dr.field_91
		}
	}
	sub_45A160_drawable(dr)
	sub_49BCD0(dr)
	sub_49BAF0(dr)
	nox_xxx_sprite_49BA10(dr)
	nox_xxx_cliRemoveHealthbar_459E30(dr, 3)
	sub_459F70(dr)
	if int32(*(*uint8)(unsafe.Pointer(&dr.flags28)))&4 != 0 {
		sub_459F00(dr)
	}
	if nox_xxx_servObjectHasTeam_419130(int32(uintptr(unsafe.Pointer(&dr.field_6)))) != 0 {
		nox_xxx_netChangeTeamMb_419570(int32(uintptr(unsafe.Pointer(&dr.field_6))), int32(dr.field_32))
	}
	nox_xxx_spriteDelete_45A4B0(dr)
}
func nox_xxx_spriteDeleteAll_45A5E0(a1 int32) {
	var (
		v1 int32
		v2 int32
	)
	v1 = int32(uintptr(unsafe.Pointer(nox_drawable_head_unk1)))
	if v1 != 0 {
		for {
			v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 368))))
			if (int32(*(*uint8)(unsafe.Pointer(uintptr(v1 + 112))))&4) == 0 || a1 == 0 || nox_xxx_spriteIsPlayerSprite_45A630(v1) == 0 {
				nox_xxx_spriteDeleteStatic_45A4E0_drawable((*nox_drawable)(unsafe.Pointer(uintptr(v1))))
			}
			v1 = v2
			if v2 == 0 {
				break
			}
		}
	}
}
func nox_xxx_spriteIsPlayerSprite_45A630(a1 int32) int32 {
	var v1 *byte
	v1 = (*byte)(unsafe.Pointer(nox_common_playerInfoGetFirst_416EA0()))
	if v1 == nil {
		return 0
	}
	for *(*uint32)(unsafe.Pointer(uintptr(a1 + 128))) != *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v1))), unsafe.Sizeof(uint32(0))*515))) {
		v1 = (*byte)(unsafe.Pointer(nox_common_playerInfoGetNext_416EE0((*nox_playerInfo)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v1)))))))))
		if v1 == nil {
			return 0
		}
	}
	return 1
}
func sub_45A670(a1 uint32) {
	var (
		result int32
		v2     *uint32
		v3     *uint32
	)
	result = int32(dword_5d4594_1046604)
	if dword_5d4594_1046604 == 0 {
		result = nox_xxx_getTTByNameSpriteMB_44CFC0("SummonEffect")
		dword_5d4594_1046604 = uint32(result)
	}
	v2 = &nox_drawable_head_unk1.field_0
	if v2 == nil {
		return
	}
	for {
		{
			var result int32 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*28)))
			v3 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*92)))))
			if (uint32(result) & 0x20400006) == 0 {
				if sub_49C520(int32(uintptr(unsafe.Pointer(v2)))) == 0 {
					if *(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*27)) != dword_5d4594_1046604 && *(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*80)) < a1 {
						nox_xxx_spriteDeleteStatic_45A4E0_drawable((*nox_drawable)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v2)))))))
					}
				}
			}
			v2 = v3
		}
		if v3 == nil {
			break
		}
	}
}
func nox_xxx_netSpriteByCodeDynamic_45A6F0(a1 int32) *uint32 {
	var result *uint32
	result = &nox_drawable_head_unk1.field_0
	if result == nil {
		return nil
	}
	for *(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*28))&0x20400000 != 0 || *(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*32)) != uint32(a1) {
		result = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*92)))))
		if result == nil {
			return nil
		}
	}
	return result
}
func nox_xxx_netSpriteByCodeStatic_45A720(a1 int32) *uint32 {
	var result *uint32
	result = &nox_drawable_head_unk1.field_0
	if result == nil {
		return nil
	}
	for (*(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*28))&0x20400000) == 0 || *(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*32)) != uint32(a1) {
		result = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*92)))))
		if result == nil {
			return nil
		}
	}
	return result
}
func nox_xxx_unused_45A750() int32 {
	var (
		result int32
		v1     *uint32
		v2     int32
		v3     int32
		v4     int32
		v5     int32
		v6     int32
		v7     int32
		v8     int32
		v9     int32
		v10    int32
	)
	result = int32(*memmap.PtrUint32(0x852978, 8))
	v1 = &nox_drawable_head_unk1.field_0
	v10 = int32(*memmap.PtrUint32(0x852978, 8))
	if v1 != nil {
		for {
			if *(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*28))&0x300000 != 0 {
				v2 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*8)))
				if v2 != 0 {
					v4 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*3)))
					v5 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*4)))
					v6 = int32((*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*3)) - uint32(v2)) >> 1)
					v7 = v6
					if v6 < 0 {
						v7 = -v6
					}
					if v7 > 1 {
						v4 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*3)) + uint32(v6))
					}
					v8 = int32((*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*4)) - *(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*9))) >> 1)
					v9 = v8
					if v8 < 0 {
						v9 = -v8
					}
					if v9 > 1 {
						v5 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*4)) + uint32(v8))
					}
					result = v6
					if v6 < 0 {
						result = -v6
					}
					if result <= 10 {
						result = int32((*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*4)) - *(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*9))) >> 1)
						if v8 < 0 {
							result = -v8
						}
						if result <= 10 {
							*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*72)) = nox_frame_xxx_2598000
							nox_xxx_updateSpritePosition_49AA90((*nox_drawable)(unsafe.Pointer(v1)), v4, v5)
							result = v10
							if v1 == (*uint32)(unsafe.Pointer(uintptr(v10))) {
								if v6 < 0 {
									v6 = -v6
								}
								if v6 > 1 {
									goto LABEL_30
								}
								if v8 < 0 {
									v8 = -v8
								}
								if v8 > 1 {
								LABEL_30:
									nox_xxx_cliUpdateCameraPos_435600(int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*3))), int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*4))))
								}
							}
						}
					}
				} else {
					v3 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*4)))
					*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*8)) = *(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*3))
					*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*9)) = uint32(v3)
				}
			}
			v1 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*92)))))
			if v1 == nil {
				break
			}
		}
	}
	return result
}
func sub_45A840(a1 *uint32) int32 {
	var (
		result int32
		v2     int32
		v3     int32
	)
	if nox_xxx_spriteCheckFlag31_4356C0((*nox_drawable)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1)))))), 23) {
		nox_xxx_spriteChangeLightColor_484BE0((*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*34)), 128, 128, math.MaxUint8)
		nox_xxx_spriteChangeIntensity_484D70_light_intensity(int32(uintptr(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*34))))), 300.0)
		result = 1
	} else if nox_xxx_spriteCheckFlag31_4356C0((*nox_drawable)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1)))))), 15) || a1 == *(**uint32)(memmap.PtrOff(0x852978, 8)) && int32(sub_467430())&8 != 0 {
		nox_xxx_spriteChangeLightColor_484BE0((*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*34)), math.MaxUint8, math.MaxUint8, math.MaxUint8)
		nox_xxx_spriteChangeIntensity_484D70_light_intensity(int32(uintptr(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*34))))), 200.0)
		result = 1
	} else {
		v2 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*28)))
		if (v2&2) == 0 || uint32(v2)&0x80000 != 0 {
			result = 0
		} else if *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*69)) == 10 {
			nox_xxx_spriteChangeIntensity_484D70_light_intensity(int32(uintptr(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*34))))), 0.0)
			result = 1
		} else {
			nox_xxx_spriteChangeLightColor_484BE0((*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*34)), math.MaxUint8, math.MaxUint8, math.MaxUint8)
			v3 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*29)))
			if v3&1 != 0 {
				nox_xxx_spriteChangeIntensity_484D70_light_intensity(int32(uintptr(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*34))))), 25.0)
				result = 1
			} else if v3&2 != 0 {
				nox_xxx_spriteChangeIntensity_484D70_light_intensity(int32(uintptr(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*34))))), 35.0)
				result = 1
			} else {
				if v3&4 != 0 {
					nox_xxx_spriteChangeIntensity_484D70_light_intensity(int32(uintptr(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*34))))), 45.0)
				}
				result = 1
			}
		}
	}
	return result
}
func nox_xxx_spriteSetActiveMB_45A990_drawable(a1 int32) int32 {
	var result int32
	result = a1
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 120))) |= 4
	return result
}
func nox_xxx_cliDestroyObj_45A9A0(a1 int32) int32 {
	var result int32
	result = a1
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 120))) &= 0xFFFFFFFB
	return result
}
func sub_45A9B0(a1 int32, a2 int32) *int32 {
	var (
		v2     int32
		v3     int32
		v4     *byte
		v5     *byte
		result *int32 = nil
		v7     int32
		v8     int32
		v9     int32
		v10    int32
		v11    int64
		v12    int32
		v13    *int32
		v14    *int32
		v15    *int32
		v16    int32
		v17    *byte
		v18    *int32
	)
	v2 = a1
	v3 = 0
	v16 = 0
	v4 = nox_xxx_draw_452270(int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 492)))))
	v5 = v4
	v17 = v4
	result = &nox_draw_getViewport_437250().x1
	v18 = result
	if v5 != nil && result != nil {
		if *(*uint32)(unsafe.Pointer(uintptr(a1 + 120)))&0x1000000 != 0 && (int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 280))))&12) == 0 {
			v7 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 12))) - *(*uint32)(unsafe.Pointer(uintptr(a1 + 12))))
			v8 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 16))) - *(*uint32)(unsafe.Pointer(uintptr(a1 + 16))))
			v9 = sub_4522A0(int32(uintptr(unsafe.Pointer(v17))))
			v10 = v9
			if v7 < v9 && v8 < v9 && v9 > 0 {
				v11 = int64(math.Sqrt(float64(v8*v8 + v7*v7 + 1)))
				if int32(v11) < v10 {
					v12 = (v10 - int32(v11)) * 100 / v10
					v3 = v12
					if v12 <= 100 {
						if v12 < 0 {
							v3 = 0
						}
					} else {
						v3 = 100
					}
					v16 = (*(*int32)(unsafe.Pointer(uintptr(a1 + 12))) - *(*int32)(unsafe.Add(unsafe.Pointer(v18), unsafe.Sizeof(int32(0))*6)) - *v18) * 50 / (nox_win_width / 2)
				}
			}
			v2 = a1
		}
		v13 = (*int32)(unsafe.Pointer(uintptr(v2 + 496)))
		result = (*int32)(unsafe.Pointer(uintptr(sub_452EB0(v13))))
		v14 = result
		if v3 != 0 {
			if result != nil {
				sub_452FE0(int32(uintptr(unsafe.Pointer(result))), v16)
				result = (*int32)(unsafe.Pointer(uintptr(sub_452F50(int32(uintptr(unsafe.Pointer(v14))), v3))))
			} else {
				result = (*int32)(unsafe.Pointer(nox_xxx_draw_452300((*uint32)(unsafe.Pointer(v17)))))
				v15 = result
				if result != nil {
					sub_452EE0(int32(uintptr(unsafe.Pointer(result))), v3)
					sub_452F80(int32(uintptr(unsafe.Pointer(v15))), v16)
					result = (*int32)(unsafe.Pointer(uintptr(sub_452E90((*uint32)(unsafe.Pointer(v13)), int32(uintptr(unsafe.Pointer(v15)))))))
				}
			}
		} else if result != nil {
			result = (*int32)(unsafe.Pointer(uintptr(sub_4523D0((*uint32)(unsafe.Pointer(result))))))
		}
	}
	return result
}
func sub_45AB40() int32 {
	var (
		result int32
		i      int32
	)
	result = sub_45A090()
	for i = result; result != 0; i = result {
		if *(*uint32)(unsafe.Pointer(uintptr(i + 492))) != 0 {
			sub_45A9B0(i, *memmap.PtrInt32(0x852978, 8))
		}
		result = sub_45A0A0(i)
	}
	return result
}
func nox_xxx_spriteSetFrameMB_45AB80(a1 int32, a2 int32) int32 {
	var result int32
	result = a1
	if (int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 112))))&2) == 0 || (*(*uint32)(unsafe.Pointer(uintptr(a1 + 116)))&0x40000) == 0 || *(*uint32)(unsafe.Pointer(uintptr(a1 + 276))) != 8 {
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 312))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 308)))
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 308))) = uint32(a2)
	}
	return result
}
func nox_xxx_guiSpellSortFn_45ABC0(a1 unsafe.Pointer, a2 unsafe.Pointer) int32 {
	var (
		v2     int32
		v3     *libc.WChar
		v4     *libc.WChar
		result int32
		v6     int32
	)
	v2 = int32(*(*uint32)(a2))
	if dword_5d4594_1046868 == 1 {
		v3 = (*libc.WChar)(unsafe.Pointer(uintptr(nox_xxx_guiCreatureGetName_427240(int32(*(*uint32)(a1))))))
		v4 = (*libc.WChar)(unsafe.Pointer(uintptr(nox_xxx_guiCreatureGetName_427240(v2))))
	} else {
		v6 = int32(*(*uint32)(a1))
		if int32(*(*uint8)(unsafe.Pointer(uintptr(dword_5d4594_1047516 + 2251)))) != 0 {
			v3 = nox_xxx_spellTitle_424930(v6)
			v4 = nox_xxx_spellTitle_424930(v2)
		} else {
			v3 = (*libc.WChar)(unsafe.Pointer(uintptr(nox_xxx_abilityGetName_0_425260(v6))))
			v4 = (*libc.WChar)(unsafe.Pointer(uintptr(nox_xxx_abilityGetName_0_425260(v2))))
		}
	}
	if v3 != nil && v4 != nil {
		result = _nox_wcsicmp(v3, v4)
	} else {
		result = 0
	}
	return result
}
func nox_xxx_bookSetColor_45AC40() int32 {
	var result int32
	*memmap.PtrUint32(6112660, 1046880) = nox_color_rgb_4344A0(15, 15, 15)
	result = int32(nox_color_rgb_4344A0(115, 100, 100))
	*memmap.PtrUint32(6112660, 1046884) = uint32(result)
	return result
}
func nox_client_toggleSpellbook_45AC70() {
	if nox_win_unk1.Flags().IsHidden() != 0 {
		nox_xxx_bookShowMB_45AD70(0)
	} else {
		nox_xxx_bookHideMB_45ACA0(0)
	}
}
func nox_xxx_bookHideMB_45ACA0(a1 int32) int32 {
	var (
		v1 int32
		v2 int32
	)
	if nox_win_unk1.Flags().IsHidden() != 0 || dword_5d4594_1047520 == 1 {
		return 0
	}
	dword_5d4594_1046864 = 0
	dword_5d4594_3799524 = 1
	dword_5d4594_1046868 = uint32(bool2int(dword_5d4594_1046872 != 0))
	v1 = int32(uintptr(unsafe.Pointer(nox_xxx_wndGetCaptureMain())))
	v2 = int32(uintptr(unsafe.Pointer(nox_win_unk1)))
	if unsafe.Pointer(uintptr(v1)) == unsafe.Pointer(nox_win_unk1) {
		nox_xxx_wndClearCaptureMain(nox_win_unk1)
		v2 = int32(uintptr(unsafe.Pointer(nox_win_unk1)))
	}
	(*nox_window)(unsafe.Pointer(uintptr(v2))).Hide()
	clientPlaySoundSpecial(787, 100)
	if a1 != 0 {
		nox_xxx_aNox_cfg_0_587000_132132 = 1
		dword_5d4594_1046936 = 0
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1046952))))).Hide()
	}
	if nox_xxx_bookGetSpellDnDType_477670() == 1 {
		nox_xxx_bookSpellDnDclear_477660()
	}
	return 1
}
func nox_xxx_guiSpellSortList_45ADF0(a1 int32) int32 {
	var (
		v1 int32
		i  int32
		v3 int32
		k  int32
		v5 int32
		j  int32
		v7 int32
		v9 int32
	)
	v1 = 0
	dword_5d4594_1046656 = uint32(nox_xxx_guiFontHeightMB_43F320(nil) + 2)
	*memmap.PtrUint32(6112660, 1047508) = 0
	dword_5d4594_1047512 = 0
	v9 = int32((141/dword_5d4594_1046656)*2 - 2)
	if dword_5d4594_1046868 == 1 {
		for i = nox_xxx_bookGetFirstCreMB_427300(); i != 0; i = nox_xxx_bookGetNextCre_427320(i) {
			if noxflags.HasGame(noxflags.GameOnline) && !noxflags.HasGame(noxflags.GameModeQuest) || *(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1047516 + uint32(i*4) + 4244))) != 0 {
				if nox_xxx_bookCreatureTest_4D70C0(i) != 0 {
					v3 = int32(*memmap.PtrUint32(6112660, 1047508))
					v1 = 1
					*memmap.PtrUint32(6112660, *memmap.PtrUint32(6112660, 1047508)*4+0xFF9B0) = uint32(i)
					*memmap.PtrUint32(6112660, 1047508) = uint32(v3 + 1)
				}
			}
		}
	} else if a1 != 0 {
		for j = nox_xxx_spellFirstValid_424AD0(); j != 0; j = nox_xxx_spellNextValid_424AF0(j) {
			if j != 34 && nox_xxx_playerCheckSpellClass_57AEA0(a1, j) == 0 && (noxflags.HasGame(noxflags.GameOnline) && !noxflags.HasGame(noxflags.GameModeQuest) || *(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1047516 + uint32(j*4) + 3696))) != 0) {
				if nox_xxx_spellHasFlags_424A50(j, 0x15000) {
					dword_5d4594_1047512++
				}
				if !nox_xxx_spellHasFlags_424A50(j, 8192) {
					v7 = int32(*memmap.PtrUint32(6112660, 1047508))
					v1 = 1
					*memmap.PtrUint32(6112660, *memmap.PtrUint32(6112660, 1047508)*4+0xFF9B0) = uint32(j)
					*memmap.PtrUint32(6112660, 1047508) = uint32(v7 + 1)
				}
			}
		}
	} else {
		for k = nox_xxx_bookFirstKnownAbil_425330(); k != 0; k = nox_xxx_bookNextKnownAbil_425350(k) {
			if noxflags.HasGame(noxflags.GameOnline) && !noxflags.HasGame(noxflags.GameModeQuest) || *(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1047516 + uint32(k*4) + 3696))) != 0 {
				v5 = int32(*memmap.PtrUint32(6112660, 1047508))
				v1 = 1
				*memmap.PtrUint32(6112660, *memmap.PtrUint32(6112660, 1047508)*4+0xFF9B0) = uint32(k)
				*memmap.PtrUint32(6112660, 1047508) = uint32(v5 + 1)
			}
		}
	}
	*memmap.PtrUint32(6112660, 1046940) = (*memmap.PtrUint32(6112660, 1047508)-dword_5d4594_1047512)/uint32(v9) + 1
	libc.Sort(memmap.PtrOff(6112660, 1046960), *memmap.PtrUint32(6112660, 1047508)-dword_5d4594_1047512, 4, nox_xxx_guiSpellSortFn_45ABC0)
	return v1
}
func nox_xxx_book_45B010(a1 int32) {
	var result uint32
	result = nox_gui_xxx_check_446360()
	if result != 0 {
		return
	}
	result = uint32(sub_4AE3D0())
	if result != 0 {
		return
	}
	if a1 != 0 || nox_xxx_get_57AF20() == 0 {
		dword_5d4594_1046864 = 1
		dword_5d4594_1046868 = uint32(bool2int(dword_5d4594_1046872 != 0))
		nox_xxx_wndShowModalMB(nox_win_unk1)
		clientPlaySoundSpecial(786, 100)
	}
}
func nox_xxx_bookWndProc_45B070(a1 int32, a2 int32) int32 {
	if dword_5d4594_1047520 == 1 {
		return 1
	}
	if a2 == 10 {
		nox_client_toggleSpellbook_45AC70()
		return 1
	}
	if *memmap.PtrUint32(0x852978, 8) != 0 && sub_478030() == 0 && sub_47A260() == 0 && (int32(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(0x852978, 8) + 120))))&2) == 2 {
		return 1
	}
	if a2 != 5 {
		return 0
	}
	(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1046944))))).Show()
	if nox_xxx_aNox_cfg_0_587000_132132 != 0 {
		if dword_5d4594_1046936+1 < uint32(*memmap.PtrInt32(6112660, 1046940)) {
			dword_5d4594_1046936++
		} else {
			dword_5d4594_1046932 = 0
			nox_xxx_aNox_cfg_0_587000_132132 = 0
			(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1046952))))).Show()
		}
		if dword_5d4594_1046872 != 0 {
			**(**uint32)(unsafe.Pointer(uintptr(dword_5d4594_1046924 + 96))) = uint32(cgoFuncAddr(nox_xxx_bookClickCreature_45B200))
		} else {
			**(**uint32)(unsafe.Pointer(uintptr(dword_5d4594_1046924 + 96))) = uint32(cgoFuncAddr(nox_xxx_bookClickSpell_45B1F0))
		}
		dword_5d4594_1046868 = 2
		*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1046924 + 96))) + 12))) = uint32(nox_xxx_bookGet_430B40_get_mouse_prev_seq())
	LABEL_23:
		clientPlaySoundSpecial(788, 100)
		return 1
	}
	if dword_5d4594_1046932 < uint32(*memmap.PtrInt32(6112660, 1047508))-dword_5d4594_1047512-1 {
		dword_5d4594_1046932++
		if dword_5d4594_1046872 != 0 {
			**(**uint32)(unsafe.Pointer(uintptr(dword_5d4594_1046924 + 96))) = uint32(cgoFuncAddr(nox_xxx_bookClickCreature_45B200))
		} else {
			**(**uint32)(unsafe.Pointer(uintptr(dword_5d4594_1046924 + 96))) = uint32(cgoFuncAddr(nox_xxx_bookClickSpell_45B1F0))
		}
		dword_5d4594_1046868 = 2
		*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1046924 + 96))) + 12))) = uint32(nox_xxx_bookGet_430B40_get_mouse_prev_seq())
		goto LABEL_23
	}
	return 1
}
func nox_xxx_bookClickSpell_45B1F0() int32 {
	var result int32
	result = 0
	dword_5d4594_1046868 = 0
	dword_5d4594_1046872 = 0
	return result
}
func nox_xxx_bookClickCreature_45B200() int32 {
	var result int32
	result = 1
	dword_5d4594_1046868 = 1
	dword_5d4594_1046872 = 1
	return result
}
func nox_xxx_book_45B210(a1 int32, a2 int32) int32 {
	if dword_5d4594_1047520 == 1 {
		return 1
	}
	if a2 == 10 {
		nox_client_toggleSpellbook_45AC70()
		return 1
	}
	if *memmap.PtrUint32(0x852978, 8) != 0 && sub_478030() == 0 && sub_47A260() == 0 && (int32(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(0x852978, 8) + 120))))&2) == 2 {
		return 1
	}
	if a2 != 5 {
		return 0
	}
	(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1046948))))).Show()
	if nox_xxx_aNox_cfg_0_587000_132132 == 0 {
		if *(*int32)(unsafe.Pointer(&dword_5d4594_1046932)) < *memmap.PtrInt32(6112660, 1047508)-*(*int32)(unsafe.Pointer(&dword_5d4594_1047512)) {
			if *(*int32)(unsafe.Pointer(&dword_5d4594_1046932)) <= 0 {
				nox_xxx_aNox_cfg_0_587000_132132 = 1
				dword_5d4594_1046936 = *memmap.PtrUint32(6112660, 1046940) - 1
				(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1046952))))).Hide()
			} else {
				dword_5d4594_1046932--
			}
		} else {
			nox_xxx_aNox_cfg_0_587000_132132 = 1
			dword_5d4594_1046936 = 0
			(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1046952))))).Hide()
		}
	LABEL_18:
		if dword_5d4594_1046872 != 0 {
			**(**uint32)(unsafe.Pointer(uintptr(dword_5d4594_1046928 + 96))) = uint32(cgoFuncAddr(nox_xxx_bookClickCreature_45B200))
		} else {
			**(**uint32)(unsafe.Pointer(uintptr(dword_5d4594_1046928 + 96))) = uint32(cgoFuncAddr(nox_xxx_bookClickSpell_45B1F0))
		}
		dword_5d4594_1046868 = 3
		*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1046928 + 96))) + 12))) = uint32(nox_xxx_bookGet_430B40_get_mouse_prev_seq())
		clientPlaySoundSpecial(788, 100)
		return 1
	}
	if dword_5d4594_1046936 != 0 {
		dword_5d4594_1046936--
		goto LABEL_18
	}
	return 1
}
func nox_xxx_bookChildWndProcMB_45B360(a1 *uint32, a2 uint32) int32 {
	var v2 int32
	v2 = 0
	if dword_5d4594_1047520 == 1 || *memmap.PtrUint32(0x852978, 8) != 0 && sub_478030() == 0 && sub_47A260() == 0 && (int32(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(0x852978, 8) + 120))))&2) == 2 {
		return 1
	}
	if a2 != 5 {
		if a2 <= 5 || a2 > 7 {
			return 0
		}
		return 1
	}
	if *a1 == 1310 {
		if dword_5d4594_1046872 != 0 {
			dword_5d4594_1046868 = 0
			dword_5d4594_1046872 = 0
			if nox_xxx_guiSpellSortList_45ADF0(int32(*(*uint8)(unsafe.Pointer(uintptr(dword_5d4594_1047516 + 2251))))) == 0 {
				dword_5d4594_1046868 = 1
				dword_5d4594_1046872 = 1
				nox_xxx_guiSpellSortList_45ADF0(int32(*(*uint8)(unsafe.Pointer(uintptr(dword_5d4594_1047516 + 2251)))))
				clientPlaySoundSpecial(925, 100)
				return 1
			}
			v2 = 1
		}
		if nox_xxx_aNox_cfg_0_587000_132132 == 0 || dword_5d4594_1046936 != 0 || v2 == 1 {
			nox_xxx_aNox_cfg_0_587000_132132 = 1
			dword_5d4594_1046936 = 0
			clientPlaySoundSpecial(788, 100)
			dword_5d4594_1046868 = 3
			*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1046928 + 96))) + 12))) = uint32(nox_xxx_bookGet_430B40_get_mouse_prev_seq())
			**(**uint32)(unsafe.Pointer(uintptr(dword_5d4594_1046928 + 96))) = uint32(cgoFuncAddr(nox_xxx_bookClickSpell_45B1F0))
			(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1046952))))).Hide()
			(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1046948))))).Show()
			nox_xxx_wndClearCaptureMain(nox_win_unk1)
			nox_xxx_bookSpellDnDclear_477660()
		}
		return 1
	}
	if *a1 != 1320 {
		return 1
	}
	if dword_5d4594_1046872 != 1 {
		dword_5d4594_1046868 = 1
		dword_5d4594_1046872 = 1
		if nox_xxx_guiSpellSortList_45ADF0(int32(*(*uint8)(unsafe.Pointer(uintptr(dword_5d4594_1047516 + 2251))))) == 0 {
			dword_5d4594_1046868 = 0
			dword_5d4594_1046872 = 0
			nox_xxx_guiSpellSortList_45ADF0(int32(*(*uint8)(unsafe.Pointer(uintptr(dword_5d4594_1047516 + 2251)))))
			clientPlaySoundSpecial(925, 100)
			return 1
		}
		v2 = 1
	}
	if nox_xxx_aNox_cfg_0_587000_132132 != 0 && dword_5d4594_1046936 == 0 && v2 != 1 {
		return 1
	}
	nox_xxx_aNox_cfg_0_587000_132132 = 1
	dword_5d4594_1046936 = 0
	clientPlaySoundSpecial(788, 100)
	if v2 == 1 {
		dword_5d4594_1046868 = 2
		*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1046924 + 96))) + 12))) = uint32(nox_xxx_bookGet_430B40_get_mouse_prev_seq())
		**(**uint32)(unsafe.Pointer(uintptr(dword_5d4594_1046924 + 96))) = uint32(cgoFuncAddr(nox_xxx_bookClickCreature_45B200))
	} else {
		dword_5d4594_1046868 = 3
		*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1046928 + 96))) + 12))) = uint32(nox_xxx_bookGet_430B40_get_mouse_prev_seq())
		**(**uint32)(unsafe.Pointer(uintptr(dword_5d4594_1046928 + 96))) = uint32(cgoFuncAddr(nox_xxx_bookClickCreature_45B200))
	}
	(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1046952))))).Hide()
	(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1046948))))).Show()
	nox_xxx_wndClearCaptureMain(nox_win_unk1)
	nox_xxx_bookSpellDnDclear_477660()
	return 1
}
func nox_xxx_bookListWndProc_45B5F0(a1 int32, a2 uint32, a3 uint32) int32 {
	var (
		v3     int32
		v4     int32
		v5     int32
		result int32
		v7     bool
		v8     int32
		v9     int32
		v10    int32
		v11    int8
		v12    int32
		v13    int32
		v14    int32
	)
	v3 = int32(a3 >> 16)
	v4 = int32(uint16(a3))
	if dword_5d4594_1047520 == 1 || *memmap.PtrUint32(0x852978, 8) != 0 && sub_478030() == 0 && sub_47A260() == 0 && (int32(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(0x852978, 8) + 120))))&2) == 2 {
		return 1
	}
	if nox_xxx_aNox_cfg_0_587000_132132 == 0 {
		if a2 >= 5 {
			if a2 <= 8 {
				return 1
			}
			if a2 == 11 {
			LABEL_47:
				nox_client_toggleSpellbook_45AC70()
				return 1
			}
		}
		return 0
	}
	dword_5d4594_1046656 = uint32(nox_xxx_guiFontHeightMB_43F320(nil) + 2)
	v5 = int32(141/dword_5d4594_1046656 - 1)
	nox_gui_getWindowOffs_46AA20((*nox_window)(unsafe.Pointer(uintptr(a1))), (*uint32)(unsafe.Pointer(&v14)), &a3)
	switch a2 {
	case 5:
		if *memmap.PtrUint32(0x852978, 8) != 0 && (int32(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(0x852978, 8) + 120))))&2) == 2 {
			return 1
		}
		dword_5d4594_1047536 = uint32(v3)
		v7 = int32(0xFFFFFFED-a3+uint32(v3)) < 0
		v8 = int32(0xFFFFFFED - a3 + uint32(v3))
		dword_5d4594_1047532 = uint32(v4)
		if v7 {
			nox_xxx_aNox_cfg_0_587000_132136 = math.MaxUint32
			return 1
		}
		v9 = v8 / *(*int32)(unsafe.Pointer(&dword_5d4594_1046656))
		if v8 / *(*int32)(unsafe.Pointer(&dword_5d4594_1046656)) >= v5 {
			goto LABEL_14
		}
		if v4-v14 > 145 {
			v9 += v5
		}
		nox_xxx_aNox_cfg_0_587000_132136 = uint32(v5*2)*dword_5d4594_1046936 + uint32(v9)
		nox_xxx_wndSetCaptureMain((*nox_window)(unsafe.Pointer(uintptr(a1))))
		v10 = int32(nox_xxx_aNox_cfg_0_587000_132136)
		if nox_xxx_aNox_cfg_0_587000_132136 >= uint32(*memmap.PtrInt32(6112660, 1047508))-dword_5d4594_1047512 {
			goto LABEL_14
		}
		v11 = int8(*(*uint8)(unsafe.Pointer(uintptr(dword_5d4594_1047516 + 2251))))
		if int32(v11) == 0 && dword_5d4594_1046868 == 0 {
			dword_5d4594_1047528 = *memmap.PtrUint32(6112660, nox_xxx_aNox_cfg_0_587000_132136*4+0xFF9B0)
			nox_xxx_bookSaveSpellForDragDrop_477640(*(*int32)(unsafe.Pointer(&dword_5d4594_1047528)), 1)
			clientPlaySoundSpecial(793, 100)
			return 1
		}
		if int32(v11) != 2 || dword_5d4594_1046868 != 1 {
			if dword_5d4594_1046868 != 0 {
				return 1
			}
			dword_5d4594_1047528 = *memmap.PtrUint32(6112660, nox_xxx_aNox_cfg_0_587000_132136*4+0xFF9B0)
			nox_xxx_bookSaveSpellForDragDrop_477640(*(*int32)(unsafe.Pointer(&dword_5d4594_1047528)), 1)
			clientPlaySoundSpecial(793, 100)
			result = 1
		} else {
			if *(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1047516 + 4232))) == 0 {
				if !noxflags.HasGame(noxflags.GameOnline) || noxflags.HasGame(noxflags.GameModeQuest) {
					nox_xxx_wndClearCaptureMain((*nox_window)(unsafe.Pointer(uintptr(a1))))
					nox_xxx_bookMoveToPage_45B930(*(*int32)(unsafe.Pointer(&nox_xxx_aNox_cfg_0_587000_132136)))
					dword_5d4594_1047528 = 0
				LABEL_14:
					nox_xxx_aNox_cfg_0_587000_132136 = math.MaxUint32
					return 1
				}
				v10 = int32(nox_xxx_aNox_cfg_0_587000_132136)
			}
			dword_5d4594_1047528 = *memmap.PtrUint32(6112660, uint32(v10*4)+0xFF9B0) + 74
			nox_xxx_bookSaveSpellForDragDrop_477640(*(*int32)(unsafe.Pointer(&dword_5d4594_1047528)), 1)
			clientPlaySoundSpecial(793, 100)
			result = 1
		}
	case 6:
		fallthrough
	case 7:
		if int32(nox_xxx_aNox_cfg_0_587000_132136) >= 0 {
			v12 = int32(dword_5d4594_1047532 - uint32(v4))
			if *(*int32)(unsafe.Pointer(&dword_5d4594_1047532))-v4 < 0 {
				v12 = int32(uint32(v4) - dword_5d4594_1047532)
			}
			if v12 >= 5 {
				goto LABEL_52
			}
			v13 = int32(dword_5d4594_1047536 - uint32(v3))
			if *(*int32)(unsafe.Pointer(&dword_5d4594_1047536))-v3 < 0 {
				v13 = int32(uint32(v3) - dword_5d4594_1047536)
			}
			if v13 >= 5 {
			LABEL_52:
				nox_xxx_bookSpellDrop_45DCA0(*(*int32)(unsafe.Pointer(&dword_5d4594_1047528)), 0, v4, v3, nil)
			} else {
				nox_xxx_bookMoveToPage_45B930(*(*int32)(unsafe.Pointer(&nox_xxx_aNox_cfg_0_587000_132136)))
			}
		}
		if nox_xxx_wndGetCaptureMain() != nil {
			nox_xxx_wndClearCaptureMain((*nox_window)(unsafe.Pointer(uintptr(a1))))
		}
		nox_xxx_bookSpellDnDclear_477660()
		return 1
	case 8:
		return 1
	case 11:
		goto LABEL_47
	default:
		return 0
	}
	return result
}
func nox_xxx_bookMoveToPage_45B930(a1 int32) int32 {
	nox_xxx_aNox_cfg_0_587000_132132 = 0
	dword_5d4594_1046932 = uint32(a1)
	dword_5d4594_1046936 = 99
	if dword_5d4594_1046872 != 0 {
		**(**uint32)(unsafe.Pointer(uintptr(dword_5d4594_1046924 + 96))) = uint32(cgoFuncAddr(nox_xxx_bookClickCreature_45B200))
	} else {
		**(**uint32)(unsafe.Pointer(uintptr(dword_5d4594_1046924 + 96))) = uint32(cgoFuncAddr(nox_xxx_bookClickSpell_45B1F0))
	}
	dword_5d4594_1046868 = 2
	*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1046924 + 96))) + 12))) = uint32(nox_xxx_bookGet_430B40_get_mouse_prev_seq())
	clientPlaySoundSpecial(788, 100)
	(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1046952))))).Show()
	(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1046944))))).Show()
	return (*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1046948))))).Show()
}
func nox_xxx_book_45BD30(a1 int32, a2 int32) int32 {
	return 1
}
func nox_xxx_bookInit_45B9D0() int32 {
	var (
		result int32
		v1     *uint32
	)
	_ = v1
	var v2 *uint32
	_ = v2
	var v3 *byte
	var v4 *byte
	dword_5d4594_1047516 = *memmap.PtrUint32(0x8531A0, 2576)
	result = int32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg("ArrowNW"))))
	*memmap.PtrUint32(6112660, 1046888) = uint32(result)
	if result == 0 {
		return 0
	}
	result = int32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg("ArrowN"))))
	*memmap.PtrUint32(6112660, 1046892) = uint32(result)
	if result == 0 {
		return 0
	}
	result = int32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg("ArrowNE"))))
	*memmap.PtrUint32(6112660, 1046896) = uint32(result)
	if result == 0 {
		return 0
	}
	result = int32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg("ArrowW"))))
	*memmap.PtrUint32(6112660, 1046900) = uint32(result)
	if result == 0 {
		return 0
	}
	result = int32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg("ArrowE"))))
	*memmap.PtrUint32(6112660, 1046908) = uint32(result)
	if result == 0 {
		return 0
	}
	result = int32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg("ArrowSW"))))
	*memmap.PtrUint32(6112660, 1046912) = uint32(result)
	if result == 0 {
		return 0
	}
	result = int32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg("ArrowS"))))
	*memmap.PtrUint32(6112660, 1046916) = uint32(result)
	if result == 0 {
		return 0
	}
	result = int32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg("ArrowSE"))))
	*memmap.PtrUint32(6112660, 1046920) = uint32(result)
	if result == 0 {
		return 0
	}
	result = int32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg("BookOfKnowledge"))))
	*memmap.PtrUint32(6112660, 1046856) = uint32(result)
	if result == 0 {
		return 0
	}
	result = int32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg("GuideTabLit"))))
	*memmap.PtrUint32(6112660, 1046660) = uint32(result)
	if result == 0 {
		return 0
	}
	result = int32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg("SpellTabLit"))))
	*memmap.PtrUint32(6112660, 1046644) = uint32(result)
	if result == 0 {
		return 0
	}
	result = int32(uintptr(unsafe.Pointer(nox_xxx_gLoadAnim_42FA20(CString("BookPageForward")))))
	dword_5d4594_1046924 = uint32(result)
	if result == 0 {
		return 0
	}
	**(**uint32)(unsafe.Pointer(uintptr(result + 96))) = uint32(cgoFuncAddr(nox_xxx_bookClickSpell_45B1F0))
	result = int32(uintptr(unsafe.Pointer(nox_xxx_gLoadAnim_42FA20(CString("BookPageBackward")))))
	dword_5d4594_1046928 = uint32(result)
	if result == 0 {
		return 0
	}
	**(**uint32)(unsafe.Pointer(uintptr(result + 96))) = uint32(cgoFuncAddr(nox_xxx_bookClickSpell_45B1F0))
	nox_win_unk1 = nox_window_new(nil, 1196, 5, nox_win_height-323, 285, 168, nil)
	nox_win_unk1.SetAllFuncs(func(arg1 int32, arg2 int32, arg3 int32, arg4 int32) int32 {
		return nox_xxx_bookListWndProc_45B5F0(arg1, uint32(arg2), uint32(arg3))
	}, func(arg1 *nox_window, arg2 unsafe.Pointer) int32 {
		return nox_xxx_bookDrawList_45BD40(int32(uintptr(unsafe.Pointer(arg1))))
	}, nil)
	nox_win_unk1.Hide()
	result = int32(uintptr(unsafe.Pointer(nox_window_new(nox_win_unk1, 8, 257, 15, 27, 40, nil))))
	v1 = (*uint32)(unsafe.Pointer(uintptr(result)))
	if result == 0 {
		return 0
	}
	(*nox_window)(unsafe.Pointer(uintptr(result))).SetAllFuncs(func(arg1 int32, arg2 int32, arg3 int32, arg4 int32) int32 {
		return nox_xxx_bookChildWndProcMB_45B360((*uint32)(unsafe.Pointer(uintptr(arg1))), uint32(arg2))
	}, func(arg1 *nox_window, arg2 unsafe.Pointer) int32 {
		return nox_xxx_book_45BD30(int32(uintptr(unsafe.Pointer(arg1))), int32(uintptr(arg2)))
	}, unsafe.Pointer(cgoFuncAddr(nox_xxx_book_45CF00)))
	*v1 = 1320
	result = int32(uintptr(unsafe.Pointer(nox_window_new(nox_win_unk1, 8, 253, 61, 27, 40, nil))))
	v2 = (*uint32)(unsafe.Pointer(uintptr(result)))
	if result == 0 {
		return 0
	}
	(*nox_window)(unsafe.Pointer(uintptr(result))).SetAllFuncs(func(arg1 int32, arg2 int32, arg3 int32, arg4 int32) int32 {
		return nox_xxx_bookChildWndProcMB_45B360((*uint32)(unsafe.Pointer(uintptr(arg1))), uint32(arg2))
	}, func(arg1 *nox_window, arg2 unsafe.Pointer) int32 {
		return nox_xxx_book_45BD30(int32(uintptr(unsafe.Pointer(arg1))), int32(uintptr(arg2)))
	}, unsafe.Pointer(cgoFuncAddr(nox_xxx_book_45CF00)))
	*v2 = 1310
	dword_5d4594_1046944 = uint32(uintptr(unsafe.Pointer(nox_window_new(nox_win_unk1, 136, 24, 138, 20, 20, nil))))
	(*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1046944)))).SetAllFuncs(func(arg1 int32, arg2 int32, arg3 int32, arg4 int32) int32 {
		return nox_xxx_book_45B210(arg1, arg2)
	}, nil, nil)
	v3 = (*byte)(unsafe.Pointer(nox_xxx_gLoadImg("ArrowW")))
	nox_xxx_wndSetIcon_46AE60(*(*int32)(unsafe.Pointer(&dword_5d4594_1046944)), int32(uintptr(unsafe.Pointer(v3))))
	dword_5d4594_1046948 = uint32(uintptr(unsafe.Pointer(nox_window_new(nox_win_unk1, 136, 233, 138, 20, 20, nil))))
	(*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1046948)))).SetAllFuncs(func(arg1 int32, arg2 int32, arg3 int32, arg4 int32) int32 {
		return nox_xxx_bookWndProc_45B070(arg1, arg2)
	}, nil, nil)
	v4 = (*byte)(unsafe.Pointer(nox_xxx_gLoadImg("ArrowE")))
	nox_xxx_wndSetIcon_46AE60(*(*int32)(unsafe.Pointer(&dword_5d4594_1046948)), int32(uintptr(unsafe.Pointer(v4))))
	dword_5d4594_1046952 = uint32(uintptr(unsafe.Pointer(nox_window_new(nox_win_unk1, 8, 63, 19, 30, 30, nil))))
	(*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1046952)))).SetAllFuncs(func(arg1 int32, arg2 int32, arg3 int32, arg4 int32) int32 {
		return nox_xxx_bookWndFn_45CC10((*uint32)(unsafe.Pointer(uintptr(arg1))), arg2, uint32(arg3))
	}, func(arg1 *nox_window, arg2 unsafe.Pointer) int32 {
		return nox_xxx_bookDrawIconFn_45CB30((*uint32)(unsafe.Pointer(arg1)))
	}, nil)
	(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1046952))))).Hide()
	dword_5d4594_1046956 = uint32(uintptr(unsafe.Pointer(nox_window_new(nil, 40, 0, 0, 30, 30, nil))))
	(*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1046956)))).SetAllFuncs(nil, func(arg1 *nox_window, arg2 unsafe.Pointer) int32 {
		return nox_xxx_bookDrawFn_45C7D0((*uint32)(unsafe.Pointer(arg1)))
	}, nil)
	(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1046956))))).Hide()
	dword_5d4594_1046868 = 0
	dword_5d4594_1046872 = 0
	nox_xxx_aNox_cfg_0_587000_132132 = 1
	dword_5d4594_1046936 = 0
	return 1
}
func nox_xxx_bookDrawIconFn_45CB30(a1 *uint32) int32 {
	var (
		v1 int32
		v2 int32
		v4 int32
		v5 int32
	)
	if dword_5d4594_1046868 != 0 {
		if dword_5d4594_1046868 != 1 || int32(*(*uint8)(unsafe.Pointer(uintptr(dword_5d4594_1047516 + 2251)))) != 2 || *(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1047516 + 4232))) == 0 && (!noxflags.HasGame(noxflags.GameOnline) || noxflags.HasGame(noxflags.GameModeQuest)) {
			return 1
		}
		v1 = int32(uintptr(nox_xxx_spellIcon_424A90(int32(*memmap.PtrUint32(6112660, dword_5d4594_1046932*4+0xFF9B0) + 74))))
	} else if int32(*(*uint8)(unsafe.Pointer(uintptr(dword_5d4594_1047516 + 2251)))) != 0 {
		v1 = int32(uintptr(nox_xxx_spellIcon_424A90(int32(*memmap.PtrUint32(6112660, dword_5d4594_1046932*4+0xFF9B0)))))
	} else {
		v1 = int32(uintptr(nox_xxx_spellGetAbilityIcon_425310(int32(*memmap.PtrUint32(6112660, dword_5d4594_1046932*4+0xFF9B0)), 0)))
	}
	v2 = v1
	if v1 != 0 {
		nox_client_wndGetPosition_46AA60((*nox_window)(unsafe.Pointer(a1)), (*uint32)(unsafe.Pointer(&v5)), (*uint32)(unsafe.Pointer(&v4)))
		nox_client_drawImageAt_47D2C0((*nox_video_bag_image_t)(unsafe.Pointer(uintptr(v2))), v5, v4)
	}
	return 1
}
func nox_xxx_bookWndFn_45CC10(a1 *uint32, a2 int32, a3 uint32) int32 {
	var (
		result int32
		v4     int32
		v5     int8
		v6     int8
	)
	if *memmap.PtrUint32(0x852978, 8) != 0 && (int32(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(0x852978, 8) + 120))))&2) == 2 {
		return 1
	}
	v4 = int32(a3 >> 16)
	if dword_5d4594_1046868 == 1 {
		v5 = int8(*(*uint8)(unsafe.Pointer(uintptr(dword_5d4594_1047516 + 2251))))
		if int32(v5) == 1 || int32(v5) == 0 || !noxflags.HasGame(noxflags.GameOnline) && !noxflags.HasGame(noxflags.GameModeQuest) && *(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1047516 + 4232))) == 0 {
			return 0
		}
	}
	switch a2 {
	case 5:
		if nox_xxx_wndGetCaptureMain() != nil || dword_5d4594_1047540 != 0 {
			goto LABEL_44
		}
		v6 = int8(*(*uint8)(unsafe.Pointer(uintptr(dword_5d4594_1047516 + 2251))))
		if int32(v6) == 0 && dword_5d4594_1046868 == 0 {
			dword_5d4594_1047540 = *memmap.PtrUint32(6112660, dword_5d4594_1046932*4+0xFF9B0)
			nox_xxx_bookSaveSpellForDragDrop_477640(*(*int32)(unsafe.Pointer(&dword_5d4594_1047540)), 1)
		LABEL_22:
			nox_xxx_wndSetCaptureMain((*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1)))))))
			clientPlaySoundSpecial(793, 100)
			return 1
		}
		if int32(v6) == 2 && dword_5d4594_1046868 == 1 {
			dword_5d4594_1047540 = *memmap.PtrUint32(6112660, dword_5d4594_1046932*4+0xFF9B0) + 74
			nox_xxx_bookSaveSpellForDragDrop_477640(*(*int32)(unsafe.Pointer(&dword_5d4594_1047540)), 1)
			goto LABEL_22
		}
		if dword_5d4594_1046868 != 0 {
			goto LABEL_44
		}
		dword_5d4594_1047540 = *memmap.PtrUint32(6112660, dword_5d4594_1046932*4+0xFF9B0)
		if !nox_xxx_spellHasFlags_424A50(*(*int32)(unsafe.Pointer(&dword_5d4594_1047540)), 0x15000) {
			nox_xxx_bookSaveSpellForDragDrop_477640(*(*int32)(unsafe.Pointer(&dword_5d4594_1047540)), 1)
			goto LABEL_22
		}
		dword_5d4594_1047540 = 0
		return 1
	case 6:
		fallthrough
	case 7:
		if dword_5d4594_1047540 == 0 {
			goto LABEL_44
		}
		if !nox_xxx_wndPointInWnd_46AAB0(a1, int32(uint16(a3)), v4) {
			if nox_xxx_guiSpell_460650() != 0 {
				nox_xxx_guiSpellTargetClickSet_45D9D0(*(*int32)(unsafe.Pointer(&dword_5d4594_1047540)))
			} else if sub_4611A0() != 0 {
				sub_4611B0()
			} else {
				nox_xxx_bookSpellDrop_45DCA0(*(*int32)(unsafe.Pointer(&dword_5d4594_1047540)), 0, int32(uint16(a3)), v4, nil)
			}
		LABEL_42:
			dword_5d4594_1047540 = 0
			nox_xxx_wndClearCaptureMain((*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1)))))))
			goto LABEL_43
		}
		if int32(*(*uint8)(unsafe.Pointer(uintptr(dword_5d4594_1047516 + 2251)))) != 0 {
			if nox_xxx_guiSpell_460650() != 0 {
				nox_xxx_guiSpellSetCursor_45DF60(0, 0)
				dword_5d4594_1047540 = 0
				nox_xxx_wndClearCaptureMain((*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1)))))))
				nox_xxx_bookSpellDnDclear_477660()
				return 1
			}
			if !nox_xxx_spellHasFlags_424A50(*(*int32)(unsafe.Pointer(&dword_5d4594_1047540)), 1536) {
				nox_xxx_guiSpellSetCursor_45DF60(*(*int32)(unsafe.Pointer(&dword_5d4594_1047540)), 0)
				nox_xxx_bookSpellDnDclear_477660()
				return 1
			}
			if nox_xxx_guiSpellSetCursor_45DF60(*(*int32)(unsafe.Pointer(&dword_5d4594_1047540)), 1) != 0 {
				goto LABEL_42
			}
		} else {
			if sub_4611A0() != 0 {
				sub_45DFC0(0)
				dword_5d4594_1047540 = 0
				nox_xxx_wndClearCaptureMain((*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1)))))))
				nox_xxx_bookSpellDnDclear_477660()
				return 1
			}
			if sub_45DFC0(*(*int32)(unsafe.Pointer(&dword_5d4594_1047540))) != 0 {
				goto LABEL_42
			}
		}
	LABEL_43:
		nox_xxx_bookSpellDnDclear_477660()
	LABEL_44:
		result = 1
	case 8:
		goto LABEL_44
	default:
		return 0
	}
	return result
}
func sub_45CFA0() int32 {
	var result int32
	if nox_xxx_aNox_cfg_0_587000_132132 != 0 {
		result = 0
	} else {
		result = int32(*memmap.PtrUint32(6112660, dword_5d4594_1046932*4+0xFF9B0))
	}
	return result
}
func sub_45CFC0() int32 {
	return (int32(uint8(int8(^nox_xxx_wndGetFlags_46ADA0(int32(uintptr(unsafe.Pointer(nox_win_unk1))))))) >> 4) & 1
}
func nox_xxx_netSpellRewardCli_45CFE0(a1 int32, a2 int32, a3 int32, a4 int32) {
	var (
		v4  int32
		v5  int32
		v6  int32
		v7  *int32
		v8  int32
		v9  int32
		v10 *uint8
		v11 int32
	)
	v4 = int32(*memmap.PtrUint32(0x8531A0, 2576))
	if *memmap.PtrUint32(0x8531A0, 2576) == 0 {
		return
	}
	v5 = a1
	if nox_xxx_playerCheckSpellClass_57AEA0(int32(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(0x8531A0, 2576) + 2251)))), a1) == 9 {
		return
	}
	*(*uint32)(unsafe.Pointer(uintptr(v4 + a1*4 + 3696))) = uint32(a2)
	if nox_xxx_spellHasFlags_424A50(a1, 4096) {
		v11 = 8192
		goto LABEL_9
	}
	if nox_xxx_spellHasFlags_424A50(a1, 0x4000) {
		v11 = 0x8000
		goto LABEL_9
	}
	if nox_xxx_spellHasFlags_424A50(a1, 0x10000) {
		v11 = 0x20000
	LABEL_9:
		v6 = 1
		v7 = (*int32)(unsafe.Pointer(uintptr(v4 + 3700)))
		for {
			if nox_xxx_spellHasFlags_424A50(v6, v11) && nox_xxx_spellIsValid_424B50(v6) {
				*v7 = a2
			}
			v6++
			v7 = (*int32)(unsafe.Add(unsafe.Pointer(v7), unsafe.Sizeof(int32(0))*1))
			if v6 >= 137 {
				break
			}
		}
	}
	dword_5d4594_1046868 = 0
	dword_5d4594_1046872 = 0
	nox_xxx_guiSpellSortList_45ADF0(int32(*(*uint8)(unsafe.Pointer(uintptr(v4 + 2251)))))
	if v5 == 34 {
		v8 = 1
		if *(*uint32)(unsafe.Pointer(uintptr(v4 + 3832))) != 1 || a3 == 0 {
			v8 = 0
		}
		nox_xxx_quickbarAddTrap_460EC0(v8)
	} else if a3 != 0 {
		v9 = 0
		v10 = (*uint8)(memmap.PtrOff(6112660, 1046960))
		for {
			if *(*uint32)(unsafe.Pointer(v10)) == uint32(v5) {
				break
			}
			v10 = (*uint8)(unsafe.Add(unsafe.Pointer(v10), 4))
			v9++
			if int32(uintptr(unsafe.Pointer(v10))) >= int32(uintptr(memmap.PtrOff(6112660, 1047508))) {
				break
			}
		}
		if v9 != 137 {
			nox_xxx_bookHideMB_45ACA0(0)
			nox_xxx_bookMoveToPage_45B930(v9)
			nox_xxx_book_45B010(0)
			nox_xxx_bookRewardCli_499CF0((*int32)(unsafe.Pointer(uintptr(2))), v5, a4)
		}
	}
}
func nox_xxx_netGuideRewardCli_45D140(a1 int32, a2 int32) {
	var (
		v2 int32
		v3 *uint8
		v4 *uint32
		v5 int32
		v6 *int32
		v7 int32
		v8 int32
		v9 *uint8
	)
	v2 = int32(*memmap.PtrUint32(0x8531A0, 2576))
	if *memmap.PtrUint32(0x8531A0, 2576) != 0 {
		v3 = (*uint8)(memmap.PtrOff(0x587000, 132124))
		*(*uint32)(unsafe.Pointer(uintptr(*memmap.PtrUint32(0x8531A0, 2576) + uint32(a1*4) + 4244))) = 1
		if *memmap.PtrUint32(0x587000, 132124) != 0 {
			for {
				v4 = *(**uint32)(unsafe.Pointer(v3))
				if uint32(a1) == **(**uint32)(unsafe.Pointer(v3)) {
					v5 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*1)))
					v6 = (*int32)(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*1))))
					if v5 != 0 {
						for {
							v6 = (*int32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(int32(0))*1))
							*(*uint32)(unsafe.Pointer(uintptr(v2 + v5*4 + 4244))) = 1
							v5 = *v6
							if *v6 == 0 {
								break
							}
						}
					}
				}
				v7 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v3))), unsafe.Sizeof(uint32(0))*1))))
				v3 = (*uint8)(unsafe.Add(unsafe.Pointer(v3), 4))
				if v7 == 0 {
					break
				}
			}
		}
		dword_5d4594_1046868 = 1
		dword_5d4594_1046872 = 1
		nox_xxx_guiSpellSortList_45ADF0(int32(*(*uint8)(unsafe.Pointer(uintptr(v2 + 2251)))))
		if a2 != 0 {
			v8 = 0
			v9 = (*uint8)(memmap.PtrOff(6112660, 1046960))
			for {
				if *(*uint32)(unsafe.Pointer(v9)) == uint32(a1) {
					break
				}
				v9 = (*uint8)(unsafe.Add(unsafe.Pointer(v9), 4))
				v8++
				if int32(uintptr(unsafe.Pointer(v9))) >= int32(uintptr(memmap.PtrOff(6112660, 1047124))) {
					break
				}
			}
			if v8 != 41 {
				nox_xxx_bookHideMB_45ACA0(0)
				nox_xxx_bookMoveToPage_45B930(v8)
				nox_xxx_book_45B010(0)
				nox_xxx_bookRewardCli_499CF0((*int32)(unsafe.Pointer(uintptr(4))), a1, 0)
			}
		}
	}
}
func nox_xxx_bookSetForward_45D200(a1 *int32, a2 int32, a3 *int2) *int32 {
	var (
		result *int32
		v4     int32
		v5     int32
	)
	nox_xxx_bookShowMB_45AD70(1)
	nox_win_unk1.SetPos(image.Pt(a3.field_0, a3.field_4))
	result = a1
	if uintptr(unsafe.Pointer(a1)) == uintptr(2) || uintptr(unsafe.Pointer(a1)) == uintptr(4) {
		v5 = 0
		result = memmap.PtrInt32(6112660, 1046960)
		for {
			if *result == a2 {
				break
			}
			result = (*int32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(int32(0))*1))
			v5++
			if int32(uintptr(unsafe.Pointer(result))) >= int32(uintptr(memmap.PtrOff(6112660, 1047508))) {
				break
			}
		}
		if v5 != 137 {
			result = (*int32)(unsafe.Pointer(uintptr(nox_xxx_bookMoveToPage_45B930(v5))))
		}
	} else if uintptr(unsafe.Pointer(a1)) == uintptr(3) {
		v4 = 0
		result = memmap.PtrInt32(6112660, 1046960)
		for {
			if *result == a2 {
				break
			}
			result = (*int32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(int32(0))*1))
			v4++
			if int32(uintptr(unsafe.Pointer(result))) >= int32(uintptr(memmap.PtrOff(6112660, 1046984))) {
				break
			}
		}
		if v4 != 6 {
			result = (*int32)(unsafe.Pointer(uintptr(nox_xxx_bookMoveToPage_45B930(v4))))
		}
	}
	return result
}
func nox_xxx_abilityReward_45D290(a1 int32, a2 *byte, a3 int32) {
	var (
		result *byte
		v4     int32
	)
	result = (*byte)(unsafe.Pointer(noxServer.getPlayerByID(int32(nox_player_netCode_85319C))))
	if result != nil {
		dword_5d4594_1046868 = 0
		dword_5d4594_1046872 = 0
		nox_xxx_guiSpellSortList_45ADF0(int32(uint8(*(*byte)(unsafe.Add(unsafe.Pointer(result), 2251)))))
		result = a2
		if a2 != nil {
			v4 = 0
			result = (*byte)(memmap.PtrOff(6112660, 1046960))
			for {
				if *(*uint32)(unsafe.Pointer(result)) == uint32(a1) {
					break
				}
				result = (*byte)(unsafe.Add(unsafe.Pointer(result), 4))
				v4++
				if int32(uintptr(unsafe.Pointer(result))) >= int32(uintptr(memmap.PtrOff(6112660, 1046984))) {
					break
				}
			}
			if v4 != 6 {
				nox_xxx_bookHideMB_45ACA0(0)
				nox_xxx_bookMoveToPage_45B930(v4)
				nox_xxx_book_45B010(0)
				result = (*byte)(unsafe.Pointer(uintptr(a3)))
				if a3 != 0 {
					nox_xxx_bookRewardCli_499CF0((*int32)(unsafe.Pointer(uintptr(3))), a1, a3)
				}
			}
		}
	}
}
func sub_45D320(a1 int32) int32 {
	var (
		v1     int32
		result int32
		v3     int32
		v4     int32
		v5     *uint32
	)
	v1 = int32(*memmap.PtrUint32(0x8531A0, 2576))
	result = nox_xxx_bookHideMB_45ACA0(1)
	if v1 != 0 {
		*(*uint32)(unsafe.Pointer(uintptr(v1 + a1*4 + 3696))) = 0
		sub_461360(a1)
		if nox_xxx_spellHasFlags_424A50(a1, 4096) {
			v3 = 8192
		} else if nox_xxx_spellHasFlags_424A50(a1, 0x4000) {
			v3 = 0x8000
		} else {
			if !nox_xxx_spellHasFlags_424A50(a1, 0x10000) {
				return nox_xxx_guiSpellSortList_45ADF0(int32(*(*uint8)(unsafe.Pointer(uintptr(v1 + 2251)))))
			}
			v3 = 0x20000
		}
		v4 = 1
		v5 = (*uint32)(unsafe.Pointer(uintptr(v1 + 3700)))
		for {
			if nox_xxx_spellHasFlags_424A50(v4, v3) {
				if nox_xxx_spellIsValid_424B50(v4) {
					*v5 = 0
					sub_461360(v4)
				}
			}
			v4++
			v5 = (*uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(uint32(0))*1))
			if v4 >= 137 {
				break
			}
		}
		return nox_xxx_guiSpellSortList_45ADF0(int32(*(*uint8)(unsafe.Pointer(uintptr(v1 + 2251)))))
	}
	return result
}
func sub_45D400(a1 int32) int32 {
	var (
		v1     int32
		result int32
		v3     *uint8
		v4     *uint32
		v5     int32
		i      *uint32
		v7     int32
	)
	v1 = int32(*memmap.PtrUint32(0x8531A0, 2576))
	result = nox_xxx_bookHideMB_45ACA0(1)
	if v1 != 0 {
		*(*uint32)(unsafe.Pointer(uintptr(v1 + a1*4 + 4244))) = 0
		sub_461360(a1 + 74)
		v3 = (*uint8)(memmap.PtrOff(0x587000, 132124))
		if *memmap.PtrUint32(0x587000, 132124) != 0 {
			for {
				v4 = *(**uint32)(unsafe.Pointer(v3))
				if uint32(a1) == **(**uint32)(unsafe.Pointer(v3)) {
					v5 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*1)))
					for i = (*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*1)); v5 != 0; i = (*uint32)(unsafe.Add(unsafe.Pointer(i), unsafe.Sizeof(uint32(0))*1)) {
						*(*uint32)(unsafe.Pointer(uintptr(v1 + v5*4 + 4244))) = 0
						sub_461360(int32(*i + 74))
						v5 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(i), unsafe.Sizeof(uint32(0))*1)))
					}
				}
				v7 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v3))), unsafe.Sizeof(uint32(0))*1))))
				v3 = (*uint8)(unsafe.Add(unsafe.Pointer(v3), 4))
				if v7 == 0 {
					break
				}
			}
		}
		result = nox_xxx_guiSpellSortList_45ADF0(int32(*(*uint8)(unsafe.Pointer(uintptr(v1 + 2251)))))
	}
	return result
}
func nox_xxx_clientQuestDisableAbility_45D4A0(a1 int32) *byte {
	var result *byte
	nox_xxx_bookHideMB_45ACA0(1)
	nox_xxx_netAbilityRewardCli_4611E0(a1, 0, nil)
	if dword_5d4594_1047516 != 0 {
		*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1047516 + uint32(a1*4) + 3696))) = 0
	}
	sub_461360(a1)
	result = (*byte)(unsafe.Pointer(noxServer.getPlayerByID(int32(nox_player_netCode_85319C))))
	if result != nil {
		result = (*byte)(unsafe.Pointer(uintptr(nox_xxx_guiSpellSortList_45ADF0(int32(uint8(*(*byte)(unsafe.Add(unsafe.Pointer(result), 2251))))))))
	}
	return result
}
func sub_45D500(a1 int32) int32 {
	var result int32
	result = int32(dword_5d4594_1046864)
	if a1 != 0 {
		if dword_5d4594_1046864 != 0 {
			result = nox_win_unk1.Show()
		}
	} else if dword_5d4594_1046864 != 0 {
		result = sub_47A260()
		if result == 0 {
			result = nox_win_unk1.Hide()
		}
	}
	return result
}
func sub_45D550(a1 *uint32) *uint32 {
	var result *uint32
	result = a1
	if a1 != nil {
		result = (*uint32)(unsafe.Pointer(uintptr(nox_client_wndGetPosition_46AA60((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1046952)))), a1, (*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*1))))))
	}
	return result
}
func nox_xxx_bookFillAll_45D570(a1 int32, a2 int32) {
	var (
		result int32
		v3     float64
		a4     int2
		a3     int2
	)
	result = int32(dword_5d4594_1047520)
	if dword_5d4594_1047520 != 1 {
		nox_xxx_quickBarClose_4606B0()
		*memmap.PtrUint32(6112660, 1046612) = uint32(nox_xxx_buttonsGetSelectedRow_45E180())
		if a1 != 2 || (func() int32 {
			result = bool2int(nox_xxx_spellHasFlags_424A50(a2, 0x15000))
			return result
		}()) == 0 {
			result = nox_xxx_buttonHaveSpellInBarMB_4612D0(a2)
			if result != 1 {
				if a1 == 2 || a1 == 4 {
					result = nox_xxx_buttonFindFirstEmptySlot_461250()
				} else {
					if a1 != 3 {
						return
					}
					result = sub_4612A0()
				}
				dword_5d4594_1046852 = uint32(result)
				if result != -1 {
					dword_5d4594_1047520 = 1
					if a1 == 2 || a1 == 3 {
						dword_5d4594_1046868 = 0
						dword_5d4594_1046872 = 0
					} else {
						dword_5d4594_1046868 = 1
						dword_5d4594_1046872 = 1
					}
					sub_45D550((*uint32)(memmap.PtrOff(6112660, 1046844)))
					*(*float32)(unsafe.Pointer(&dword_5d4594_1046636)) = float32(float64(*memmap.PtrInt32(6112660, 1046844)))
					*(*float32)(unsafe.Pointer(&dword_5d4594_1046640)) = float32(float64(*memmap.PtrInt32(6112660, 1046848)))
					(*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1046956)))).SetPos(image.Pt(*memmap.PtrInt32(6112660, 1046844), *memmap.PtrInt32(6112660, 1046848)))
					nox_xxx_buttonSetImgMB_461320(*(*int32)(unsafe.Pointer(&dword_5d4594_1046852)), (*uint32)(memmap.PtrOff(6112660, 1046668)))
					a3.field_0 = 400
					a4.field_4 = 400
					dword_5d4594_1047524 = uint32(a2)
					dword_5d4594_1046652 = uint32(bool2int(a1 == 3))
					*memmap.PtrUint32(6112660, 1046676) = uint32(a1)
					nox_client_renderGUI_80828 = 1
					a3.field_4 = -500
					a4.field_0 = 350
					*memmap.PtrUint32(6112660, 1046680) = 0
					sub_4BEDE0((*int2)(memmap.PtrOff(6112660, 1046844)), (*int2)(memmap.PtrOff(6112660, 1046668)), &a3, &a4, 19, 0.0, int32(cgoFuncAddr(sub_45D7D0)), 0)
					*memmap.PtrUint32(6112660, 1046628) = 0
					obj_5d4594_1046620.field_0 = *mem_getFloatPtr(6112660, 0xFF8A4) - *mem_getFloatPtr(6112660, 0xFF89C)
					obj_5d4594_1046620.field_4 = *mem_getFloatPtr(6112660, 0xFF8A8) - *mem_getFloatPtr(6112660, 0xFF8A0)
					nox_xxx_utilNormalizeVector_509F20(&obj_5d4594_1046620)
					if nox_win_width < 1000 {
						if nox_win_width < 750 {
							v3 = 6.0
						} else {
							v3 = 8.0
						}
					} else {
						v3 = 10.0
					}
					obj_5d4594_1046620.field_0 = float32(float64(obj_5d4594_1046620.field_0) * v3)
					obj_5d4594_1046620.field_4 = float32(float64(obj_5d4594_1046620.field_4) * v3)
					(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1046956))))).Show()
					nox_xxx_wndShowModalMB((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1046956))))))
					if noxflags.HasGame(noxflags.GameModeCoop) {
						sub_57AF30(0, a1)
					}
					dword_5d4594_1046648 = uint32(nox_xxx_bookGet_430B40_get_mouse_prev_seq())
					if !noxflags.HasGame(noxflags.GameModeCoop) || nox_gui_xxx_check_446360() == 1 || (func() bool {
						result = nox_xxx_gameGet_4DB1B0()
						return result == 1
					}()) {
						sub_45D870()
					}
				}
			}
		}
	}
}
func sub_45D7D0(a1 *int32, a2 *int32) int32 {
	var (
		result int32
		v3     float64
	)
	result = int32(*memmap.PtrUint32(6112660, 1046680))
	if *memmap.PtrInt32(6112660, 1046680) < 20 {
		result = int32(*memmap.PtrUint32(6112660, 1046680) + 1)
		*mem_getFloatPtr(6112660, uint32(result*8)+0xFF894) = float32(float64(*a1))
		*mem_getFloatPtr(6112660, uint32(result*8)+0xFF898) = float32(float64(*(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*1))))
		*mem_getFloatPtr(6112660, uint32(result*8)+0xFF89C) = float32(float64(*a2))
		v3 = float64(*(*int32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(int32(0))*1)))
		*memmap.PtrUint32(6112660, 1046680) = uint32(result)
		*mem_getFloatPtr(6112660, uint32(result*8)+0xFF8A0) = float32(v3)
	}
	return result
}
func sub_45D810() {
	if dword_5d4594_1047520 != 0 {
		dword_5d4594_1047520 = 0
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1046956))))).Hide()
		dword_5d4594_1046648 = 0
		nox_xxx_clientUpdateButtonRow_45E110(*memmap.PtrInt32(6112660, 1046612))
		if noxflags.HasGame(noxflags.GameModeCoop) {
			sub_57B0A0()
			sub_413A00(0)
		}
	}
}
func sub_45D9B0() int32 {
	return int32(dword_5d4594_1047520)
}
func nox_xxx_guiSpellTest_45D9C0() int32 {
	return bool2int(*memmap.PtrUint32(6112660, 1047916) != 0)
}
func nox_xxx_guiSpellTargetClickSet_45D9D0(a1 int32) int32 {
	if *memmap.PtrUint32(6112660, 1047916) != 0 {
		return 0
	}
	*memmap.PtrUint32(6112660, 1047916) = uint32(a1)
	*memmap.PtrUint8(6112660, 1047920) = 0
	nox_xxx_setKeybTimeout_4160D0(5)
	*memmap.PtrUint32(6112660, 1047928) = 0
	*memmap.PtrUint32(6112660, 1047924) = 0
	return 1
}
func nox_xxx_guiSpell_45DA10(a1 int32) int32 {
	var result int32
	if *memmap.PtrUint32(6112660, 1047916) != 0 {
		return 0
	}
	*memmap.PtrUint32(6112660, 1047916) = uint32(a1)
	*memmap.PtrUint8(6112660, 1047920) = 0
	nox_xxx_setKeybTimeout_4160D0(5)
	result = 1
	*memmap.PtrUint32(6112660, 1047924) = 1
	return result
}
func nox_client_invokeSpellSlot_45DA50(a1 int32) {
	if *memmap.PtrUint32(0x8531A0, 2576) != 0 {
		if nox_xxx_playerAnimCheck_4372B0() == 0 {
			if int32(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(0x8531A0, 2576) + 2251)))) != 0 {
				nox_xxx_clientSendSpell_45DB20((*byte)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(nox_xxx_aClosewoodengat_587000_133480)) + 204)))+uint32(a1*8)))), 1, int8(int32(*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(nox_xxx_aClosewoodengat_587000_133480)) + 204))) + uint32(a1*8) + 4))))&1))
				nox_xxx_clientStoreLastButton_45DAD0(a1)
			} else {
				nox_xxx_clientSendAbil_45DAF0(int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(nox_xxx_aClosewoodengat_587000_133480)) + 204))) + uint32(a1*8))))))
				nox_xxx_clientStoreLastButton_45DAD0(a1)
			}
		}
	}
}
func nox_xxx_clientStoreLastButton_45DAD0(a1 int32) {
	*memmap.PtrUint32(0x587000, 133484) = uint32(a1)
	*memmap.PtrUint32(6112660, 1049540) = nox_frame_xxx_2598000
}
func nox_xxx_clientSendAbil_45DAF0(a1 int32) int32 {
	var (
		result int32
		v3     int16
	)
	result = nox_xxx_guiCursor_477600()
	if result == 0 {
		result = a1
		if a1 != 0 {
			*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v3))), unsafe.Sizeof(int16(0))-1)) = uint8(int8(a1))
			*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v3))), 0)) = 122
			result = nox_netlist_addToMsgListCli_40EBC0(31, 0, (*uint8)(unsafe.Pointer(&v3)), 2)
		}
	}
	return result
}
func nox_xxx_clientSendSpell_45DB20(a1 *byte, a2 int32, a3 int8) int32 {
	var (
		result int32
		v4     int32
		v5     *byte
		v6     [22]byte
	)
	result = nox_xxx_guiCursor_477600()
	if result == 0 && *(*uint32)(unsafe.Pointer(a1)) != 0 {
		v4 = 0
		v6[0] = 121
		v5 = &v6[1]
		for {
			if v4 >= a2 {
				*(*uint32)(unsafe.Pointer(v5)) = 0
			} else {
				*(*uint32)(unsafe.Pointer(v5)) = *(*uint32)(unsafe.Add(unsafe.Pointer(v5), int32(uintptr(unsafe.Pointer(a1))-uintptr(unsafe.Pointer(&v6[1])))))
			}
			v4++
			v5 = (*byte)(unsafe.Add(unsafe.Pointer(v5), 4))
			if v4 >= 5 {
				break
			}
		}
		v6[21] = byte(a3)
		result = nox_netlist_addToMsgListCli_40EBC0(31, 0, (*uint8)(unsafe.Pointer(&v6[0])), 22)
	}
	return result
}
func sub_45DB90() int32 {
	var result int32
	result = 0
	libc.MemSet(memmap.PtrOff(6112660, 1049544), 0, 136)
	*memmap.PtrUint8(6112660, 1049680) = 0
	return result
}
func nox_xxx_guiSpellTargetClickCheckSend_45DBB0() {
	if *memmap.PtrUint32(6112660, 1047916) != 0 {
		nox_xxx_clientSendSpell_45DB20((*byte)(memmap.PtrOff(6112660, 1047916)), 1, int8(*memmap.PtrUint8(6112660, 1047924)))
		*memmap.PtrUint32(6112660, 1047916) = 0
	}
}
func nox_xxx_book_45DBE0(a1 unsafe.Pointer, a2 int32, a3 int32) unsafe.Pointer {
	var result unsafe.Pointer
	nox_xxx_spellKeyPackSetSpell_45DC40(*(*int32)(unsafe.Pointer(&nox_xxx_aClosewoodengat_587000_133480)), a2, a3)
	result = a1
	if uintptr(a1) == uintptr(2) {
		if a2 != 0 && (func() unsafe.Pointer {
			result = unsafe.Pointer(uintptr(bool2int(nox_xxx_spellHasFlags_424A50(a2, 1536))))
			return result
		}()) != nil {
			*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(nox_xxx_aClosewoodengat_587000_133480)) + 204))) + uint32(a3*8) + 4))) = 1
		} else {
			result = nox_xxx_aClosewoodengat_587000_133480
			*(*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(nox_xxx_aClosewoodengat_587000_133480)) + 204))) + uint32(a3*8) + 4))) = 0
		}
	}
	return result
}
func nox_xxx_spellKeyPackSetSpell_45DC40(a1 int32, a2 int32, a3 int32) {
	var v3 *libc.WChar
	*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 204))) + uint32(a3*8)))) = uint32(a2)
	if *(*uint32)(unsafe.Pointer(uintptr(a1 + a3*4 + 212))) != 0 {
		v3 = nox_xxx_spellTitle_424930(a2)
		nox_xxx_wndWddSetTooltip_46B000((*nox_window_data)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + a3*4 + 212)))+36))), v3)
	}
	if *memmap.PtrUint32(0x8531A0, 2576) != 0 {
		clientPlaySoundSpecial(794, 100)
	}
}
func nox_xxx_bookSpellDrop_45DCA0(a1 int32, a2 int8, a3 int32, a4 int32, a5 *int32) int32 {
	var (
		v5  uint16
		v6  int32
		v7  *int32
		v8  int32
		v9  int32
		v10 int32
		v11 int32
		v13 uint16
		v14 uint16
	)
	v5 = 0
	v14 = 137
	v13 = 0
	if ((*nox_window)(unsafe.Pointer(uintptr(*memmap.PtrInt32(6112660, 1048148)))).Flags().IsHidden() != 0 || nox_xxx_spellPutInBox_45DEB0(memmap.PtrInt32(6112660, 1047940), a1, a3, a4) == 0) && a5 != memmap.PtrInt32(6112660, 1047940) {
		if *memmap.PtrUint32(6112660, 1049476) != 0 {
			v6 = 0
		} else {
			v6 = 4
		}
		v7 = memmap.PtrInt32(6112660, uint32(v6*256)+0xFFE84)
		for {
			v8 = nox_xxx_spellBoxPointToWnd_45DE60(int32(uintptr(unsafe.Pointer(v7))), a3, a4)
			v9 = v8
			if v8 >= 0 {
				v10 = *(*int32)(unsafe.Add(unsafe.Pointer(v7), unsafe.Sizeof(int32(0))*51))
				v14 = *(*uint16)(unsafe.Pointer(uintptr(v10 + v8*8)))
				v13 = uint16(*(*uint8)(unsafe.Pointer(uintptr(v10 + v8*8 + 4))))
			}
			if nox_xxx_spellPutInBox_45DEB0(v7, a1, a3, a4) != 0 {
				break
			}
			v7 = (*int32)(unsafe.Add(unsafe.Pointer(v7), unsafe.Sizeof(int32(0))*64))
			v6++
			if int32(uintptr(unsafe.Pointer(v7))) >= int32(uintptr(memmap.PtrOff(6112660, 1049476))) {
				goto LABEL_16
			}
		}
		if nox_xxx_bookGetSpellDnDType_477670() == 1 {
			v11 = v6 << 8
			*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, uint32(v11)+0xFFF50) + uint32(v9*8) + 4))) = uint8(int8(bool2int(nox_xxx_spellHasFlags_424A50(a1, 1536))))
		} else {
			v11 = v6 << 8
			*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, uint32(v11)+0xFFF50) + uint32(v9*8) + 4))) = uint8(a2)
		}
		nox_xxx_updateSpellIcons_45DDF0(int32(uintptr(memmap.PtrOff(6112660, uint32(v11)+0xFFE84))))
	LABEL_16:
		v5 = v13
	}
	return int32(v5) | int32(v14)<<16
}
func nox_xxx_updateSpellIcons_45DDF0(a1 int32) int32 {
	var (
		i      int32
		result int32
	)
	for i = 0; i < 5; i++ {
		result = nox_xxx_updateSpellIconDir_45DE10(i, a1)
	}
	return result
}
func nox_xxx_updateSpellIconDir_45DE10(a1 int32, a2 int32) int32 {
	var (
		v2     int32
		v3     int8
		v4     *uint32
		result int32
	)
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 204))))
	v3 = int8(*(*uint8)(unsafe.Pointer(uintptr(v2 + a1*8 + 4))))
	v4 = (*uint32)(unsafe.Pointer(uintptr(v2 + a1*8)))
	if int32(v3)&1 != 0 && *v4 != 0 {
		result = sub_46AE10(int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + a1*4 + 232)))), 1)
	} else {
		result = sub_46AE10(int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + a1*4 + 232)))), 0)
	}
	return result
}
func nox_xxx_spellBoxPointToWnd_45DE60(a1 int32, a2 int32, a3 int32) int32 {
	var (
		v3 int32
		i  **uint32
	)
	v3 = 0
	for i = (**uint32)(unsafe.Pointer(uintptr(a1 + 212))); !nox_xxx_wndPointInWnd_46AAB0(*i, a2, a3); i = (**uint32)(unsafe.Add(unsafe.Pointer(i), unsafe.Sizeof((*uint32)(nil))*1)) {
		if func() int32 {
			p := &v3
			*p++
			return *p
		}() >= 5 {
			return -1
		}
	}
	return v3
}
func nox_xxx_guiSpellSetCursor_45DF60(a1 int32, a2 int8) int32 {
	var result int32
	if a1 != 0 {
		clientPlaySoundSpecial(766, 100)
		if (int32(a2)&1) == 0 || nox_xxx_spellHasFlags_424A50(a1, 8192) {
			*memmap.PtrUint32(6112660, 1047556) = uint32(a1)
			*memmap.PtrUint32(6112660, 1047928) = 1
			result = 0
		} else {
			nox_xxx_guiSpell_45DA10(a1)
			result = 1
		}
	} else {
		*memmap.PtrUint32(6112660, 1047928) = 0
		result = 0
	}
	return result
}
func sub_45DFC0(a1 int32) int32 {
	var (
		result int32
		v2     int32
		v3     int32
	)
	if a1 != 0 {
		v2 = int32(*memmap.PtrUint32(6112660, uint32(a1*24)+0xFFCD8))
		if v2 != 0 {
			dword_5d4594_1047932 = 0
			dword_5d4594_1047936 = 0
			nox_xxx_clientSendAbil_45DAF0(a1)
			clientPlaySoundSpecial(766, 100)
			result = v2
		} else {
			v3 = int32(*memmap.PtrUint32(6112660, uint32(a1*24)+0xFFCD4))
			dword_5d4594_1047932 = 1
			dword_5d4594_1047936 = uint32(v3)
			clientPlaySoundSpecial(766, 100)
			result = 0
		}
	} else {
		dword_5d4594_1047932 = 0
		dword_5d4594_1047936 = 0
		result = 0
	}
	return result
}
func nox_xxx_clientUpdateButtonRow_45E110(a1 int32) int32 {
	var result int32
	result = a1
	if a1 >= 0 && a1 < 5 && *memmap.PtrUint32(6112660, 1049476) == 0 && dword_5d4594_1049496 == 0 {
		*(*uint8)(unsafe.Pointer(uintptr(uint32(uintptr(nox_xxx_aClosewoodengat_587000_133480)) + 200))) = uint8(int8(a1))
		*(*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(nox_xxx_aClosewoodengat_587000_133480)) + 204))) = uint32(uintptr(nox_xxx_aClosewoodengat_587000_133480)) + uint32(int32(*(*uint8)(unsafe.Pointer(uintptr(uint32(uintptr(nox_xxx_aClosewoodengat_587000_133480)) + 200))))*40)
		if *memmap.PtrUint32(0x8531A0, 2576) != 0 {
			clientPlaySoundSpecial(798, 100)
		}
		result = nox_xxx_updateSpellIcons_45DDF0(*(*int32)(unsafe.Pointer(&nox_xxx_aClosewoodengat_587000_133480)))
	}
	return result
}
func nox_xxx_buttonsGetSelectedRow_45E180() int32 {
	return int32(*(*uint8)(unsafe.Pointer(uintptr(uint32(uintptr(nox_xxx_aClosewoodengat_587000_133480)) + 200))))
}
func nox_xxx_quickbarButtonBookDraw_45EF30() int32 {
	return 1
}
func sub_45EF40() int32 {
	return 0
}
func nox_xxx_quickBarWnd_45EF50(a1 int32, a2 int32, a3 uint32) int32 {
	var (
		v3  *int32
		v4  int32
		v5  *uint32
		v7  uint16
		v8  int32
		v9  *uint32
		v10 int32
		v11 int32
		v12 uint32
		v13 int32
		v14 int2
		v15 int32
	)
	v3 = *(**int32)(unsafe.Pointer(uintptr(a1 + 368)))
	if *memmap.PtrUint32(0x852978, 8) == 0 || (int32(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(0x852978, 8) + 120))))&2) != 2 {
		v4 = 0
		v5 = (*uint32)(unsafe.Pointer((*int32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(int32(0))*53))))
		for {
			if uint32(a1) == *v5 {
				break
			}
			v4++
			v5 = (*uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(uint32(0))*1))
			if v4 >= 5 {
				break
			}
		}
		if v4 == 5 {
			return 0
		}
		v7 = uint16(a3)
		v8 = v4 * 8
		v15 = int32(a3 >> 16)
		v9 = (*uint32)(unsafe.Pointer(uintptr(v4*8 + *(*int32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(int32(0))*51)))))
		v14.field_0 = v4*8 + *(*int32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(int32(0))*51))
		switch a2 {
		case 5:
			if *v9 == 0 || *memmap.PtrUint32(6112660, 1047928) != 0 || dword_5d4594_1047932 != 0 || nox_xxx_get_57AF20() != 0 || int32(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(0x8531A0, 2576) + 2251)))) != 0 && !nox_xxx_spellIsEnabled_424B70(int32(*(*uint32)(unsafe.Pointer(uintptr(v14.field_0))))) {
				return 1
			}
			nox_xxx_wndSetCaptureMain((*nox_window)(unsafe.Pointer(uintptr(a1))))
			nox_xxx_bookSaveSpellForDragDrop_477640(int32(*(*uint32)(unsafe.Pointer(uintptr(v14.field_0)))), 2)
			dword_5d4594_1049692 = uint32(uintptr(unsafe.Pointer(v3)))
			dword_5d4594_1049696 = uint32(*(*int32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(int32(0))*51)))
			clientPlaySoundSpecial(793, 100)
			return 1
		case 6:
			fallthrough
		case 7:
			if nox_xxx_wndGetCaptureMain() == nil {
				return 1
			}
			if nox_xxx_get_57AF20() != 0 {
				goto LABEL_51
			}
			v10 = nox_xxx_spellBoxPointToWnd_45DE60(int32(uintptr(unsafe.Pointer(v3))), int32(v7), v15)
			if dword_5d4594_1047932 != 0 {
				if v10 < 0 {
					sub_4611B0()
				}
				dword_5d4594_1047932 = 0
				goto LABEL_51
			}
			if *memmap.PtrUint32(6112660, 1047928) != 0 {
				if v10 < 0 {
					nox_xxx_guiSpellTargetClickSet_45D9D0(*memmap.PtrInt32(6112660, 1047556))
				}
				*memmap.PtrUint32(6112660, 1047928) = 0
				goto LABEL_52
			}
			if v10 < 0 {
				v12 = uint32(nox_xxx_bookSpellDrop_45DCA0(int32(*(*uint32)(unsafe.Pointer(uintptr(v8 + *(*int32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(int32(0))*51)))))), int8(*(*uint8)(unsafe.Pointer(uintptr(v8 + *(*int32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(int32(0))*51)) + 4)))), int32(v7), v15, v3))
				if v12>>16 == 137 {
					v14.field_0 = int32(v7)
					v14.field_4 = v15
					v13 = sub_4281F0(&v14, (*int4)(memmap.PtrOff(0x587000, 133656)))
					if v13 != 0 || dword_5d4594_1049696 != 0 && dword_5d4594_1049696 != uint32(*(*int32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(int32(0))*51))) {
						goto LABEL_51
					}
					*(*uint32)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(int32(0))*51)) + v4*8))) = 0
					*(*uint8)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(int32(0))*51)) + v8 + 4))) = 0
					if v3 == memmap.PtrInt32(6112660, 1047940) {
						goto LABEL_51
					}
				} else {
					*(*uint32)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(int32(0))*51)) + v4*8))) = v12 >> 16
					*(*uint8)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(int32(0))*51)) + v8 + 4))) = uint8(v12)
					if v3 == memmap.PtrInt32(6112660, 1047940) {
						goto LABEL_51
					}
				}
			} else {
				if dword_5d4594_1049692 != 0 {
					if *(**int32)(unsafe.Pointer(&dword_5d4594_1049692)) == v3 && unsafe.Pointer(uintptr(dword_5d4594_1049692)) == nox_xxx_aClosewoodengat_587000_133480 {
						if dword_5d4594_1049696 != 0 {
							v11 = *(*int32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(int32(0))*51))
							if dword_5d4594_1049696 != uint32(v11) {
								sub_45F390(*(*int32)(unsafe.Pointer(&dword_5d4594_1049696)), v11, v4, v10)
								clientPlaySoundSpecial(794, 100)
								nox_xxx_updateSpellIcons_45DDF0(int32(uintptr(unsafe.Pointer(v3))))
								goto LABEL_51
							}
						}
					}
				}
				if v10 == v4 {
					dword_5d4594_1049532 = uint32(a1)
					if *memmap.PtrUint32(0x8531A0, 2576) == 0 || v3 == memmap.PtrInt32(6112660, 1047940) {
						goto LABEL_53
					}
					if int32(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(0x8531A0, 2576) + 2251)))) != 0 {
						if nox_xxx_spellIsEnabled_424B70(int32(*(*uint32)(unsafe.Pointer(uintptr(v14.field_0))))) {
							nox_xxx_guiSpellSetCursor_45DF60(int32(*(*uint32)(unsafe.Pointer(uintptr(v14.field_0)))), int8(*(*uint8)(unsafe.Pointer(uintptr(v14.field_0 + 4)))))
							nox_xxx_clientStoreLastButton_45DAD0(v4)
						}
					} else {
						sub_45DFC0(int32(*(*uint32)(unsafe.Pointer(uintptr(v14.field_0)))))
						nox_xxx_clientStoreLastButton_45DAD0(v4)
					}
				LABEL_51:
					if *memmap.PtrUint32(6112660, 1047928) != 0 {
						goto LABEL_54
					}
				LABEL_52:
					if dword_5d4594_1047932 != 0 {
						goto LABEL_54
					}
				LABEL_53:
					nox_xxx_wndClearCaptureMain((*nox_window)(unsafe.Pointer(uintptr(a1))))
					dword_5d4594_1049532 = 0
				LABEL_54:
					nox_xxx_bookSpellDnDclear_477660()
					return 1
				}
				nox_xxx_clientSwapQuickbarKeys_45F300(int32(uintptr(unsafe.Pointer(v3))), v4, v10)
				clientPlaySoundSpecial(794, 100)
				if v3 == memmap.PtrInt32(6112660, 1047940) {
					goto LABEL_51
				}
			}
			nox_xxx_updateSpellIcons_45DDF0(int32(uintptr(unsafe.Pointer(v3))))
			goto LABEL_51
		case 8:
			fallthrough
		case 12:
			fallthrough
		case 16:
			return 0
		default:
			return 1
		}
	}
	return 1
}
func nox_xxx_clientSwapQuickbarKeys_45F300(a1 int32, a2 int32, a3 int32) {
	var (
		v3 int32
		v4 int32
		v5 int32
		v6 *uint32
		v7 *libc.WChar
		v8 *libc.WChar
	)
	v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 204))))
	v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(v3 + a2*8))))
	v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(v3 + a2*8 + 4))))
	*(*uint32)(unsafe.Pointer(uintptr(v3 + a2*8))) = *(*uint32)(unsafe.Pointer(uintptr(a3*8 + v3)))
	*(*uint32)(unsafe.Pointer(uintptr(v3 + a2*8 + 4))) = *(*uint32)(unsafe.Pointer(uintptr(a3*8 + v3 + 4)))
	v6 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 204))) + uint32(a3*8))))
	*v6 = uint32(v4)
	*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*1)) = uint32(v5)
	v7 = nox_xxx_spellTitle_424930(int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 204))) + uint32(a2*8))))))
	nox_xxx_wndWddSetTooltip_46B000((*nox_window_data)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + a2*4 + 212)))+36))), v7)
	v8 = nox_xxx_spellTitle_424930(int32(*(*uint32)(unsafe.Pointer(uintptr(uint32(a3*8) + *(*uint32)(unsafe.Pointer(uintptr(a1 + 204))))))))
	nox_xxx_wndWddSetTooltip_46B000((*nox_window_data)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + a3*4 + 212)))+36))), v8)
}
func sub_45F390(a1 int32, a2 int32, a3 int32, a4 int32) {
	var (
		v4 int32
		v5 int32
		v6 *libc.WChar
	)
	v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + a4*8))))
	v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + a4*8 + 4))))
	*(*uint32)(unsafe.Pointer(uintptr(a2 + a4*8))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + a3*8)))
	*(*uint32)(unsafe.Pointer(uintptr(a2 + a4*8 + 4))) = *(*uint32)(unsafe.Pointer(uintptr(a1 + a3*8 + 4)))
	*(*uint32)(unsafe.Pointer(uintptr(a1 + a3*8))) = uint32(v4)
	*(*uint32)(unsafe.Pointer(uintptr(a1 + a3*8 + 4))) = uint32(v5)
	v6 = nox_xxx_spellTitle_424930(int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + a4*8)))))
	nox_xxx_wndWddSetTooltip_46B000((*nox_window_data)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(nox_xxx_aClosewoodengat_587000_133480)) + uint32(a4*4) + 212)))+36))), v6)
}
func nox_xxx_quickbarButtonBookWnd_45F450(a1 int32, a2 uint32) int32 {
	if a2 >= 5 {
		if a2 <= 6 {
			return 1
		}
		if a2 == 7 {
			nox_client_toggleSpellbook_45AC70()
			return 1
		}
	}
	return 0
}
func sub_45F500(a1 int32, a2 int32) int32 {
	return int32((*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a2 + a1*4 + 232))) + 36))) >> 1) & 1)
}
func sub_45F520(a1 int32, a2 uint32) int32 {
	var (
		v2 int32
		v3 int32
		v4 int32
		v5 int32
		v6 int32
	)
	v2 = int32(uint16(*(*uint32)(unsafe.Pointer(uintptr(a1 + 368)))>>16)) << 8
	v3 = int32(*(*uint16)(unsafe.Pointer(uintptr(a1 + 368))))
	v4 = int32(*memmap.PtrUint32(6112660, uint32(v2)+0xFFF50))
	v5 = v4 + v3*8
	v6 = int32(*(*uint32)(unsafe.Pointer(uintptr(v4 + v3*8))))
	if v6 == 0 {
		return 0
	}
	if a2 != 5 {
		if a2 > 5 && a2 <= 7 {
			return 1
		}
		return 0
	}
	if nox_xxx_spellHasFlags_424A50(v6, 0x200400) {
		clientPlaySoundSpecial(925, 100)
	} else {
		*(*uint8)(unsafe.Pointer(uintptr(v5 + 4))) ^= 1
		nox_xxx_updateSpellIconDir_45DE10(v3, int32(uintptr(memmap.PtrOff(6112660, uint32(v2)+0xFFE84))))
		clientPlaySoundSpecial(921, 100)
	}
	return 1
}
func sub_45F5D0(a1 *uint32) int32 {
	var (
		v1 int32
		v2 *int16
		v4 int32
	)
	v1 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*92)))
	nox_client_wndGetPosition_46AA60((*nox_window)(unsafe.Pointer(a1)), (*uint32)(unsafe.Pointer(&v4)), (*uint32)(unsafe.Pointer(&a1)))
	v2 = (*int16)(unsafe.Pointer(sub_42E8E0(v1+28, 1)))
	if v2 != nil {
		nox_xxx_drawSetTextColor_434390(int32(nox_color_white_2523948))
		noxrend.DrawString(unsafe.Pointer(uintptr(*memmap.PtrInt32(6112660, 1049684))), (*libc.WChar)(unsafe.Pointer(v2)), image.Pt(v4, int32(uintptr(unsafe.Pointer(a1)))))
	}
	return 1
}
func nox_xxx_quickbarTrapUpDownProc_45F630(a1 int32, a2 uint32) int32 {
	var v2 int32
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 368))))
	if (*memmap.PtrUint32(0x852978, 8) == 0 || (int32(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(0x852978, 8) + 120))))&2) != 2) && sub_4AE3D0() == 0 {
		if a2 == 5 {
			switch v2 {
			case 0:
				*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1049508 + 60))) = 0
				nox_client_spellSetPrev_460540()
			case 1:
				*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1049508 + 60))) = 0
				nox_client_spellSetNext_4604F0()
			case 2:
				*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1049508 + 60))) = 0
				sub_460920()
			case 3:
				nox_client_trapSetPrev_4603F0()
			case 4:
				nox_client_trapSetNext_4603A0()
			default:
			}
			*memmap.PtrUint32(6112660, 1049700) = 2
			*memmap.PtrUint32(6112660, 1049704) = uint32(v2)
		} else if a2 <= 5 || a2 > 7 {
			return 0
		}
	}
	return 1
}
func nox_xxx_quickbarTrapUpDownDraw_45F6F0(a1 *uint32) int32 {
	var (
		v1 *uint32
		v2 int32
		v3 *byte
		v4 int32
		v5 int32
		v6 int32
		v8 int32
	)
	v1 = a1
	v2 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*92)))
	nox_client_wndGetPosition_46AA60((*nox_window)(unsafe.Pointer(a1)), (*uint32)(unsafe.Pointer(&a1)), (*uint32)(unsafe.Pointer(&v8)))
	v3 = (*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(a1))), *(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*24))))
	v4 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*25)) + uint32(v8))
	a1 = (*uint32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(a1))), *(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*24))))))
	v8 = v4
	if *memmap.PtrInt32(6112660, 1049700) <= 0 || *memmap.PtrInt32(6112660, 1049704) != v2 {
		v5 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*15)))
		if v5 != 0 {
			nox_client_drawImageAt_47D2C0((*nox_video_bag_image_t)(unsafe.Pointer(uintptr(v5))), int32(uintptr(unsafe.Pointer(v3))), v4)
		}
	} else {
		nox_client_drawImageAt_47D2C0((*nox_video_bag_image_t)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*19))))), int32(uintptr(unsafe.Pointer(v3))), v4)
		if func() uint32 {
			p := memmap.PtrUint32(6112660, 1049700)
			*p--
			return *p
		}() == 0 && *memmap.PtrUint32(6112660, 1049704) <= 2 {
			*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1049508 + 60))) = *(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1049508 + 76)))
		}
	}
	v6 = sub_45CFC0()
	sub_46AE10(*(*int32)(unsafe.Pointer(&dword_5d4594_1049524)), v6)
	return 1
}
func nox_xxx_quickbarTrapButtonProc_45F7A0(a1 *uint32, a2 uint32, a3 uint32) int32 {
	var result int32
	if a2 < 5 {
		goto LABEL_18
	}
	if a2 <= 6 {
		return 1
	}
	if a2 != 7 {
	LABEL_18:
		sub_46AE10(*(*int32)(unsafe.Pointer(&dword_5d4594_1049504)), 0)
		return 0
	}
	if *memmap.PtrUint32(6112660, 1047928) != 0 {
		if !nox_xxx_wndPointInWnd_46AAB0(a1, int32(uint16(a3)), int32(a3>>16)) {
			nox_client_buildTrap_45E040()
		}
		nox_xxx_wndClearCaptureMain((*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1)))))))
		dword_5d4594_1049532 = 0
		*memmap.PtrUint32(6112660, 1047928) = 0
		return 1
	}
	sub_46AE10(*(*int32)(unsafe.Pointer(&dword_5d4594_1049504)), 1)
	if *memmap.PtrUint32(0x8531A0, 2576) == 0 {
		return 1
	}
	if int32(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(0x8531A0, 2576) + 2251)))) == 1 && sub_45F890() != 0 {
		nox_xxx_wndSetCaptureMain((*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1)))))))
		dword_5d4594_1049532 = uint32(uintptr(unsafe.Pointer(a1)))
		*memmap.PtrUint32(6112660, 1047928) = 1
		result = 1
	} else {
		nox_client_buildTrap_45E040()
		result = 1
	}
	return result
}
func sub_45F890() int32 {
	var (
		v0 *int32
		v1 int32
	)
	v0 = *(**int32)(memmap.PtrOff(6112660, 1048144))
	v1 = 0
	for *v0 == 0 || (nox_xxx_spellFlags_424A70(*v0)&8) != 8 {
		v1++
		v0 = (*int32)(unsafe.Add(unsafe.Pointer(v0), unsafe.Sizeof(int32(0))*2))
		if v1 >= 3 {
			return 0
		}
	}
	return 1
}
func nox_xxx_quickbar_45F8D0(a1 int32, a2 int32, a3 int32, a4 int32) int32 {
	return bool2int(a2 != 8 && a2 != 12 && a2 != 16)
}
func sub_45F8F0() int32 {
	sub_45F900()
	return 1
}
func sub_45F900() {
	var (
		result *uint32
		i      int32
		v2     int32
		v3     int32
		v4     int32
	)
	if nox_xxx_aClosewoodengat_587000_133480 == nil {
		return
	}
	if dword_5d4594_1049496 != 0 {
		result = (*uint32)(unsafe.Pointer(uintptr(nox_xxx_bookGet_430B40_get_mouse_prev_seq())))
		if *(**uint32)(memmap.PtrOff(6112660, 1049492)) == result {
			return
		}
		dword_5d4594_1049496 = 0
	} else {
		result = (*uint32)(unsafe.Pointer(uintptr(nox_xxx_bookGet_430B40_get_mouse_prev_seq())))
		if *(**uint32)(memmap.PtrOff(6112660, 1049492)) != result {
			return
		}
		dword_5d4594_1049496 = 1
	}
	for i = 0; i < 5; i++ {
		v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(nox_xxx_aClosewoodengat_587000_133480)) + 204))))
		v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + i*8))))
		v4 = v2 + i*8
		if v3 != 0 {
			if !nox_xxx_spellHasFlags_424A50(v3, 0x200400) {
				*(*uint8)(unsafe.Pointer(uintptr(v4 + 4))) ^= 1
				nox_xxx_updateSpellIconDir_45DE10(i, *(*int32)(unsafe.Pointer(&nox_xxx_aClosewoodengat_587000_133480)))
			}
		}
	}
	nox_xxx_updateSpellIcons_45DDF0(*(*int32)(unsafe.Pointer(&nox_xxx_aClosewoodengat_587000_133480)))
	clientPlaySoundSpecial(921, 100)
}
func nox_xxx_quickbarTrapProc_45FB90(a1 int32, a2 uint32, a3 int32, a4 int32) int32 {
	if a2 >= 5 {
		if a2 <= 6 {
			return 1
		}
		if a2 == 7 {
			if (*nox_window)(unsafe.Pointer(uintptr(a1))).Flags().IsHidden() == 0 {
				sub_461060()
			}
			return 1
		}
	}
	return 0
}
func nox_xxx_quickbarDrawFn_460000() int32 {
	var (
		v0     int32
		v1     int32
		result int32
	)
	v0 = int32(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1049504 + 20))))
	v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1049504 + 20))) - dword_5d4594_1049536)
	if v1 >= 0 {
		if v1 <= 0 {
			if dword_5d4594_1049536 > uint32(nox_win_height) {
				sub_460070()
				dword_5d4594_1049536 = uint32(nox_win_height - 74)
			}
			result = 1
		} else {
			(*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1049504)))).SetPos(image.Pt(int32(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1049504 + 16)))), v0-1))
			result = 1
		}
	} else {
		(*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1049504)))).SetPos(image.Pt(int32(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1049504 + 16)))), v0+1))
		result = 1
	}
	return result
}
func nox_xxx_quickBarInitWindow_4601F0(a1 int32, a2 int32, a3 int32, a4 int32, a5 int32, a6 int32, a7 int32) int32 {
	var (
		v7     int32
		v8     *byte
		v9     int32
		v10    *uint8
		v11    *int32
		v12    *uint32
		v13    int32
		result int32
		v15    int32
	)
	v7 = a1
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 208))) = uint32(uintptr(unsafe.Pointer(nox_window_new(nil, 1224, a2-10, a3-15, 199, 59, nil))))
	v8 = (*byte)(unsafe.Pointer(nox_xxx_gLoadImg("QuickBarBase")))
	nox_xxx_wndSetIcon_46AE60(int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 208)))), int32(uintptr(unsafe.Pointer(v8))))
	nox_xxx_wndSetOffsetMB_46AE40(int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 208)))), -61, -18)
	int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 208)))).setFunc93(nox_xxx_quickbar_45F8D0)
	v9 = a5 + 10
	if a4 > 0 {
		v10 = (*uint8)(memmap.PtrOff(0x587000, 133488))
		v11 = (*int32)(unsafe.Pointer(uintptr(a1 + 212)))
		v15 = a4
		for {
			v12 = (*uint32)(unsafe.Pointer(nox_window_new((*nox_window)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v7 + 208)))))), 1032, v9, 15, 34, 34, nil)))
			*v11 = int32(uintptr(unsafe.Pointer(v12)))
			(*nox_window)(unsafe.Pointer(v12)).SetAllFuncs(cgoAsFunc(a6, (*func(int32, int32, int32, int32) int32)(nil)).(func(int32, int32, int32, int32) int32), cgoAsFunc(a7, (*func(*nox_window, unsafe.Pointer) int32)(nil)).(func(*nox_window, unsafe.Pointer) int32), nil)
			v13 = *v11
			v11 = (*int32)(unsafe.Add(unsafe.Pointer(v11), unsafe.Sizeof(int32(0))*1))
			v10 = (*uint8)(unsafe.Add(unsafe.Pointer(v10), 4))
			*(*uint32)(unsafe.Pointer(uintptr(v13 + 368))) = uint32(v7)
			v9 += int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v10))), -int(unsafe.Sizeof(uint32(0))*1)))))
			v15--
			if v15 == 0 {
				break
			}
		}
	}
	result = int32(*(*uint8)(unsafe.Pointer(uintptr(v7 + 200))))
	*(*uint32)(unsafe.Pointer(uintptr(v7 + 204))) = uint32(v7 + result*40)
	return result
}
func sub_4602F0() unsafe.Pointer {
	var (
		v0     int32
		v1     *uint8
		v2     int32
		v3     *uint32
		v4     *uint8
		v5     int32
		v6     *uint8
		v7     *uint8
		v8     int32
		result unsafe.Pointer
	)
	nox_xxx_quickBarClose_4606B0()
	v0 = 0
	v1 = (*uint8)(memmap.PtrOff(6112660, 1048196))
	for {
		v2 = 5
		v3 = (*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(nox_xxx_aClosewoodengat_587000_133480)) + uint32(v0))))
		v4 = (*uint8)(unsafe.Add(unsafe.Pointer(v1), 4))
		v5 = int32(uint32(uintptr(nox_xxx_aClosewoodengat_587000_133480)) + uint32(v0) - uint32(uintptr(unsafe.Pointer(v1))))
		for {
			*v3 = 0
			*(*uint8)(unsafe.Add(unsafe.Pointer(v4), v5)) = 0
			*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v4))), -int(unsafe.Sizeof(uint32(0))*1)))) = 0
			*v4 = 0
			v3 = (*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*2))
			v4 = (*uint8)(unsafe.Add(unsafe.Pointer(v4), 8))
			v2--
			if v2 == 0 {
				break
			}
		}
		v1 = (*uint8)(unsafe.Add(unsafe.Pointer(v1), 256))
		v0 += 40
		if int32(uintptr(unsafe.Pointer(v1))) >= int32(uintptr(memmap.PtrOff(6112660, 1049476))) {
			break
		}
	}
	v6 = (*uint8)(memmap.PtrOff(6112660, 1047940))
	for {
		v7 = v6
		v8 = 3
		for {
			*(*uint32)(unsafe.Pointer(v7)) = 0
			*(*uint8)(unsafe.Add(unsafe.Pointer(v7), 4)) = 0
			v7 = (*uint8)(unsafe.Add(unsafe.Pointer(v7), 8))
			v8--
			if v8 == 0 {
				break
			}
		}
		v6 = (*uint8)(unsafe.Add(unsafe.Pointer(v6), 40))
		if int32(uintptr(unsafe.Pointer(v6))) >= int32(uintptr(memmap.PtrOff(6112660, 1048060))) {
			break
		}
	}
	result = nox_xxx_aClosewoodengat_587000_133480
	if *(*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(nox_xxx_aClosewoodengat_587000_133480)) + 232))) != 0 {
		result = unsafe.Pointer(uintptr(nox_xxx_updateSpellIcons_45DDF0(*(*int32)(unsafe.Pointer(&nox_xxx_aClosewoodengat_587000_133480)))))
	}
	return result
}
func sub_460380() unsafe.Pointer {
	var v0 *uint8
	v0 = (*uint8)(memmap.PtrOff(6112660, 1047804))
	for {
		*(*uint32)(unsafe.Pointer(v0)) = 0
		v0 = (*uint8)(unsafe.Add(unsafe.Pointer(v0), 24))
		if int32(uintptr(unsafe.Pointer(v0))) >= int32(uintptr(memmap.PtrOff(6112660, 1047924))) {
			break
		}
	}
	return sub_4602F0()
}
func nox_client_trapSetNext_4603A0() {
	if int32(*memmap.PtrUint8(6112660, 1048140)) == 2 {
		*memmap.PtrUint8(6112660, 1048140) = 0
	} else {
		*memmap.PtrUint8(6112660, 1048140)++
	}
	*memmap.PtrUint32(6112660, 1048144) = uint32(uintptr(memmap.PtrOff(6112660, uint32(int32(*memmap.PtrUint8(6112660, 1048140))*40)+0xFFD84)))
	clientPlaySoundSpecial(798, 100)
}
func nox_client_trapSetPrev_4603F0() {
	if int32(*memmap.PtrUint8(6112660, 1048140)) != 0 {
		*memmap.PtrUint8(6112660, 1048140)--
	} else {
		*memmap.PtrUint8(6112660, 1048140) = 2
	}
	*memmap.PtrUint32(6112660, 1048144) = uint32(uintptr(memmap.PtrOff(6112660, uint32(int32(*memmap.PtrUint8(6112660, 1048140))*40)+0xFFD84)))
	clientPlaySoundSpecial(798, 100)
}
func sub_460440() {
	var result *uint32 = nil
	if nox_xxx_checkKeybTimeout_4160F0(6, nox_gameFPS>>1) {
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&result))), 0)) = 0
		*memmap.PtrUint32(6112660, 1049708) = 0
	} else {
		*memmap.PtrUint32(6112660, 1049708)++
		if *memmap.PtrInt32(6112660, 1049708) >= 3 {
			*memmap.PtrUint8(6112660, 1048140) = 2
			return
		}
	}
	*memmap.PtrUint8(6112660, 1048140) = uint8(uintptr(unsafe.Pointer(result)))
	*memmap.PtrUint32(6112660, 1048144) = uint32(uintptr(memmap.PtrOff(6112660, uint32(int32(uint8(uintptr(unsafe.Pointer(result))))*40)+0xFFD84)))
	nox_xxx_setKeybTimeout_4160D0(6)
	clientPlaySoundSpecial(798, 100)
}
func nox_client_trapSetSelect_4604B0(a1 int32) int32 {
	var result int32
	result = a1
	if a1 >= 0 && a1 < 3 {
		*memmap.PtrUint8(6112660, 1048140) = uint8(int8(a1))
		result = int32(uint8(int8(a1))) * 5
		*memmap.PtrUint32(6112660, 1048144) = uint32(uintptr(memmap.PtrOff(6112660, uint32(int32(uint8(int8(a1)))*40)+0xFFD84)))
	}
	return result
}
func sub_4604E0() int32 {
	return int32(*memmap.PtrUint8(6112660, 1048140))
}
func nox_client_spellSetNext_4604F0() int32 {
	var (
		result int32
		v1     int32
	)
	result = int32(*memmap.PtrUint32(6112660, 1049476))
	if *memmap.PtrUint32(6112660, 1049476) == 0 {
		result = int32(dword_5d4594_1049496)
		if dword_5d4594_1049496 == 0 {
			result = nox_xxx_get_57AF20()
			if result == 0 {
				nox_xxx_clientStoreLastButton_45DAD0(-1)
				v1 = int32(*(*uint8)(unsafe.Pointer(uintptr(uint32(uintptr(nox_xxx_aClosewoodengat_587000_133480)) + 200)))) + 1
				if v1 > 4 {
					v1 = 0
				}
				result = nox_xxx_clientUpdateButtonRow_45E110(v1)
			}
		}
	}
	return result
}
func nox_client_spellSetPrev_460540() int32 {
	var (
		result int32
		v1     int32
	)
	result = int32(*memmap.PtrUint32(6112660, 1049476))
	if *memmap.PtrUint32(6112660, 1049476) == 0 {
		result = int32(dword_5d4594_1049496)
		if dword_5d4594_1049496 == 0 {
			result = nox_xxx_get_57AF20()
			if result == 0 {
				nox_xxx_clientStoreLastButton_45DAD0(-1)
				v1 = int32(*(*uint8)(unsafe.Pointer(uintptr(uint32(uintptr(nox_xxx_aClosewoodengat_587000_133480)) + 200)))) - 1
				if v1 < 0 {
					v1 = 4
				}
				result = nox_xxx_clientUpdateButtonRow_45E110(v1)
			}
		}
	}
	return result
}
func nox_client_spellSetSelect_460590() {
	var result int32
	result = int32(*memmap.PtrUint32(6112660, 1049476))
	if *memmap.PtrUint32(6112660, 1049476) == 0 {
		result = int32(dword_5d4594_1049496)
		if dword_5d4594_1049496 == 0 {
			if nox_xxx_checkKeybTimeout_4160F0(7, nox_gameFPS>>1) {
				*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&result))), 0)) = 0
				*memmap.PtrUint32(6112660, 1049712) = 0
			} else {
				result = int32(func() uint32 {
					p := memmap.PtrUint32(6112660, 1049712)
					*p++
					return *p
				}())
				if *memmap.PtrInt32(6112660, 1049712) >= 5 {
					*(*uint8)(unsafe.Pointer(uintptr(uint32(uintptr(nox_xxx_aClosewoodengat_587000_133480)) + 200))) = 4
					return
				}
			}
			*(*uint8)(unsafe.Pointer(uintptr(uint32(uintptr(nox_xxx_aClosewoodengat_587000_133480)) + 200))) = uint8(int8(result))
			*(*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(nox_xxx_aClosewoodengat_587000_133480)) + 204))) = uint32(uintptr(nox_xxx_aClosewoodengat_587000_133480)) + uint32(int32(*(*uint8)(unsafe.Pointer(uintptr(uint32(uintptr(nox_xxx_aClosewoodengat_587000_133480)) + 200))))*40)
			nox_xxx_setKeybTimeout_4160D0(7)
			clientPlaySoundSpecial(798, 100)
			nox_xxx_updateSpellIcons_45DDF0(*(*int32)(unsafe.Pointer(&nox_xxx_aClosewoodengat_587000_133480)))
		}
	}
}
func sub_460630() {
	if *memmap.PtrUint32(6112660, 1049476) == 0 {
		var result int32 = nox_xxx_bookGet_430B40_get_mouse_prev_seq()
		*memmap.PtrUint32(6112660, 1049492) = uint32(result)
	}
}
func nox_xxx_guiSpell_460650() int32 {
	return int32(*memmap.PtrUint32(6112660, 1047928))
}
func sub_460660() int32 {
	var result int32
	result = int32(dword_5d4594_1049532)
	if dword_5d4594_1049532 != 0 || *memmap.PtrUint32(6112660, 1047928) != 0 || dword_5d4594_1047932 != 0 {
		nox_xxx_wndClearCaptureMain((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1049532))))))
		dword_5d4594_1049532 = 0
		*memmap.PtrUint32(6112660, 1047928) = 0
		dword_5d4594_1047932 = 0
		result = 1
	}
	return result
}
func nox_xxx_quickBarClose_4606B0() int32 {
	var (
		v0 *uint8
		v1 int32
		v2 *uint8
		v3 *int32
		v4 *uint8
		v6 int32
		v7 int32
		v8 int32
	)
	if *memmap.PtrUint32(6112660, 1049476) != 1 {
		return 0
	}
	v6 = 0
	v0 = (*uint8)(memmap.PtrOff(6112660, 1048196))
	for {
		v1 = int32(uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v0), uint32(-v6)-uint32(uintptr(nox_xxx_aClosewoodengat_587000_133480)))))))
		v2 = v0
		v3 = (*int32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v0), 232))))
		v4 = (*uint8)(unsafe.Pointer(uintptr(uint32(v6) + uint32(uintptr(nox_xxx_aClosewoodengat_587000_133480)) + 4)))
		v8 = int32(uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v0), uint32(-v6)-uint32(uintptr(nox_xxx_aClosewoodengat_587000_133480)))))))
		v7 = 5
		for {
			*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v4))), -int(unsafe.Sizeof(uint32(0))*1)))) = *(*uint32)(unsafe.Pointer(v2))
			*v4 = *(*uint8)(unsafe.Add(unsafe.Pointer(v4), v1))
			(*nox_window)(unsafe.Pointer(uintptr(*v3))).Hide()
			v3 = (*int32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(int32(0))*1))
			v2 = (*uint8)(unsafe.Add(unsafe.Pointer(v2), 8))
			v4 = (*uint8)(unsafe.Add(unsafe.Pointer(v4), 8))
			if func() int32 {
				p := &v7
				*p--
				return *p
			}() == 0 {
				break
			}
			v1 = v8
		}
		(*nox_window)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v0))), unsafe.Sizeof(uint32(0))*52)))))).Hide()
		v0 = (*uint8)(unsafe.Add(unsafe.Pointer(v0), 256))
		v6 += 40
		if int32(uintptr(unsafe.Pointer(v0))) >= int32(uintptr(memmap.PtrOff(6112660, 1049220))) {
			break
		}
	}
	*(*uint8)(unsafe.Pointer(uintptr(uint32(uintptr(nox_xxx_aClosewoodengat_587000_133480)) + 200))) = *memmap.PtrUint8(6112660, 1047908)
	*(*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(nox_xxx_aClosewoodengat_587000_133480)) + 204))) = uint32(uintptr(nox_xxx_aClosewoodengat_587000_133480)) + uint32(int32(*(*uint8)(unsafe.Pointer(uintptr(uint32(uintptr(nox_xxx_aClosewoodengat_587000_133480)) + 200))))*40)
	nox_xxx_updateSpellIcons_45DDF0(*(*int32)(unsafe.Pointer(&nox_xxx_aClosewoodengat_587000_133480)))
	(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1049512))))).Show()
	clientPlaySoundSpecial(800, 100)
	*memmap.PtrUint32(6112660, 1049476) = 0
	return 1
}
func sub_460920() {
	var (
		v0     *uint8
		v1     *int32
		v2     *uint8
		v3     *uint32
		v4     int32
		result *uint32
		v6     int32
		v7     int32
		i      int32
	)
	if *memmap.PtrUint32(6112660, 1049476) != 0 {
		nox_xxx_quickBarClose_4606B0()
		return
	}
	result = *(**uint32)(unsafe.Pointer(&dword_5d4594_1049496))
	if dword_5d4594_1049496 != 0 {
		return
	}
	result = (*uint32)(unsafe.Pointer(uintptr(nox_xxx_get_57AF20())))
	if result != nil {
		return
	}
	result = *(**uint32)(memmap.PtrOff(6112660, 1049476))
	if *memmap.PtrUint32(6112660, 1049476) != 0 {
		return
	}
	if dword_5d4594_1049484 == 1 {
		sub_461010()
	}
	v6 = 0
	v0 = (*uint8)(memmap.PtrOff(6112660, 1048428))
	for {
		v1 = (*int32)(unsafe.Pointer(v0))
		v2 = (*uint8)(unsafe.Add(unsafe.Pointer(v0), -228))
		v3 = (*uint32)(unsafe.Pointer(uintptr(uint32(v6) + uint32(uintptr(nox_xxx_aClosewoodengat_587000_133480)))))
		v7 = 5
		v4 = int32(uint32(v6) + uint32(uintptr(nox_xxx_aClosewoodengat_587000_133480)) - uint32(uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v0), -232))))))
		for i = int32(uint32(v6) + uint32(uintptr(nox_xxx_aClosewoodengat_587000_133480)) - uint32(uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v0), -232)))))); ; v4 = i {
			*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), -int(unsafe.Sizeof(uint32(0))*1)))) = *v3
			*v2 = *(*uint8)(unsafe.Add(unsafe.Pointer(v2), v4))
			(*nox_window)(unsafe.Pointer(uintptr(*v1))).Show()
			v1 = (*int32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(int32(0))*1))
			v3 = (*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*2))
			v2 = (*uint8)(unsafe.Add(unsafe.Pointer(v2), 8))
			if func() int32 {
				p := &v7
				*p--
				return *p
			}() == 0 {
				break
			}
		}
		(*nox_window)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v0))), -int(unsafe.Sizeof(uint32(0))*6))))))).Show()
		nox_xxx_updateSpellIcons_45DDF0(int32(uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v0), -232))))))
		v0 = (*uint8)(unsafe.Add(unsafe.Pointer(v0), 256))
		v6 += 40
		if int32(uintptr(unsafe.Pointer(v0))) >= int32(uintptr(memmap.PtrOff(6112660, 1049452))) {
			break
		}
	}
	*memmap.PtrUint32(6112660, 1047912) = uint32(*(*uint8)(unsafe.Pointer(uintptr(uint32(uintptr(nox_xxx_aClosewoodengat_587000_133480)) + 200))))
	*(*uint8)(unsafe.Pointer(uintptr(uint32(uintptr(nox_xxx_aClosewoodengat_587000_133480)) + 200))) = 4
	*(*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(nox_xxx_aClosewoodengat_587000_133480)) + 204))) = uint32(uintptr(nox_xxx_aClosewoodengat_587000_133480)) + uint32(int32(*(*uint8)(unsafe.Pointer(uintptr(uint32(uintptr(nox_xxx_aClosewoodengat_587000_133480)) + 200))))*40)
	nox_xxx_updateSpellIcons_45DDF0(*(*int32)(unsafe.Pointer(&nox_xxx_aClosewoodengat_587000_133480)))
	(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1049512))))).Hide()
	clientPlaySoundSpecial(799, 100)
	*memmap.PtrUint32(6112660, 1049476) = 1
}
func sub_460940(this unsafe.Pointer) int32 {
	var (
		result int32
		v2     unsafe.Pointer
	)
	v2 = this
	if nox_xxx_cryptGetXxx() == 0 && sub_461450() == 1 {
		sub_461400()
		sub_461440(0)
	}
	if *memmap.PtrUint32(0x8531A0, 2576) != 0 {
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v2))), 0)) = *(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(0x8531A0, 2576) + 2251)))
	} else {
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v2))), 0)) = uint8(*(*byte)(unsafe.Add(unsafe.Pointer(nox_xxx_getHostInfoPtr_431770()), 66)))
	}
	cryptFileReadWrite((*uint8)(unsafe.Pointer(&v2))[:1])
	if int32(uint8(uintptr(v2))) == 0 {
		result = sub_460A10(*(*int32)(unsafe.Pointer(&nox_xxx_aClosewoodengat_587000_133480)), 5, 5, 0)
		if result == 0 {
			return result
		}
		return 1
	}
	result = sub_460A10(*(*int32)(unsafe.Pointer(&nox_xxx_aClosewoodengat_587000_133480)), 5, 5, int8(uintptr(v2)))
	if result != 0 {
		result = sub_460A10(int32(uintptr(memmap.PtrOff(6112660, 1047940))), 3, 3, int8(uintptr(v2)))
		if result != 0 {
			if *(*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(nox_xxx_aClosewoodengat_587000_133480)) + 232))) != 0 {
				nox_xxx_updateSpellIcons_45DDF0(*(*int32)(unsafe.Pointer(&nox_xxx_aClosewoodengat_587000_133480)))
			}
			return 1
		}
	}
	return result
}
func sub_460A10(a1 int32, a2 int32, a3 int32, a4 int8) int32 {
	var (
		v4  int32
		v5  int32
		v6  *byte
		v7  bool
		v9  uint32
		v10 int32
		v11 int32
		v12 int32
		v13 [256]byte
	)
	if a2 <= 0 {
		return 1
	}
	v4 = a1
	v11 = a1
	v12 = a2
	for {
		v5 = a3
		if a3 > 0 {
			for {
				if nox_xxx_cryptGetXxx() == 1 {
					cryptFileReadWrite((*uint8)(unsafe.Pointer(&v10))[:1])
					cryptFileReadWrite((*uint8)(unsafe.Pointer(&v13[0]))[:uint32(uint8(int8(v10)))])
					v13[uint8(int8(v10))] = 0
					if int32(a4) != 0 {
						*(*uint32)(unsafe.Pointer(uintptr(v4))) = uint32(things.SpellIndex(&v13[0]))
					} else {
						*(*uint32)(unsafe.Pointer(uintptr(v4))) = uint32(nox_xxx_abilityNameToN_424D80(&v13[0]))
					}
				} else {
					if int32(a4) != 0 {
						*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v10))), 0)) = uint8(int8(libc.StrLen(nox_xxx_spellNameByN_424870(int32(*(*uint32)(unsafe.Pointer(uintptr(v4))))))))
						cryptFileReadWrite((*uint8)(unsafe.Pointer(&v10))[:1])
						v9 = uint32(uint8(int8(v10)))
						v6 = nox_xxx_spellNameByN_424870(int32(*(*uint32)(unsafe.Pointer(uintptr(v4)))))
					} else {
						*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v10))), 0)) = uint8(int8(libc.StrLen(nox_xxx_abilityGetName_425250(int32(*(*uint32)(unsafe.Pointer(uintptr(v4))))))))
						cryptFileReadWrite((*uint8)(unsafe.Pointer(&v10))[:1])
						v9 = uint32(uint8(int8(v10)))
						v6 = nox_xxx_abilityGetName_425250(int32(*(*uint32)(unsafe.Pointer(uintptr(v4)))))
					}
					cryptFileReadWrite((*uint8)(unsafe.Pointer(v6))[:v9])
				}
				cryptFileReadWrite((*uint8)(unsafe.Pointer(uintptr(v4 + 4)))[:1])
				v4 += 8
				v5--
				if v5 == 0 {
					break
				}
			}
		}
		v4 = v11 + 40
		v7 = v12 == 1
		v11 += 40
		v12--
		if v7 {
			break
		}
	}
	return 1
}
func sub_460B90(a1 int32) int32 {
	var (
		i      int32
		result int32
		j      int32
	)
	if a1 != 0 {
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1049500))))).Show()
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1049504))))).Show()
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1049520))))).Show()
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1049508))))).Show()
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1049512))))).Show()
		for i = 232; i < 252; i += 4 {
			(*nox_window)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(uint32(i) + uint32(uintptr(nox_xxx_aClosewoodengat_587000_133480)) - 20)))))).Show()
			(*nox_window)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(uint32(i) + uint32(uintptr(nox_xxx_aClosewoodengat_587000_133480)))))))).Show()
			if dword_5d4594_1049484 != 0 {
				(*nox_window)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, uint32(i)+0xFFD70)))).Show()
				(*nox_window)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, uint32(i)+0xFFD84)))).Show()
				(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1049512))))).Hide()
			}
		}
		if dword_5d4594_1049484 != 0 {
			(*nox_window)(unsafe.Pointer(uintptr(*memmap.PtrInt32(6112660, 1048148)))).Show()
		}
		result = (*nox_window)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(nox_xxx_aClosewoodengat_587000_133480)) + 208)))))).Show()
	} else {
		if *memmap.PtrUint32(6112660, 1049476) != 0 {
			sub_460920()
		}
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1049500))))).Hide()
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1049504))))).Hide()
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1049520))))).Hide()
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1049508))))).Hide()
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1049512))))).Hide()
		for j = 232; j < 252; j += 4 {
			(*nox_window)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(uint32(j) + uint32(uintptr(nox_xxx_aClosewoodengat_587000_133480)) - 20)))))).Hide()
			(*nox_window)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(uint32(j) + uint32(uintptr(nox_xxx_aClosewoodengat_587000_133480)))))))).Hide()
		}
		(*nox_window)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(nox_xxx_aClosewoodengat_587000_133480)) + 208)))))).Hide()
		result = (*nox_window)(unsafe.Pointer(uintptr(*memmap.PtrInt32(6112660, 1048148)))).Hide()
	}
	return result
}
