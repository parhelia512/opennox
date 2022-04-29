package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	"github.com/gotranspile/cxgo/runtime/stdio"
	"github.com/noxworld-dev/opennox/v1/common/alloc"
	noxflags "github.com/noxworld-dev/opennox/v1/common/flags"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"math"
	"unsafe"
)

type nox_parse_thing_funcs_t struct {
	name       *byte
	parse_func func(*nox_thing, *nox_memfile, *byte) bool
}

var nox_xxx_objectTypes_head_1563660 *nox_objectType_t = nil
var nox_parse_thing_funcs [21]nox_parse_thing_funcs_t = [21]nox_parse_thing_funcs_t{{name: CString("FLAGS"), parse_func: nox_parse_thing_flags}, {name: CString("CLASS"), parse_func: nox_parse_thing_class}, {name: CString("SUBCLASS"), parse_func: nox_parse_thing_subclass}, {name: CString("EXTENT"), parse_func: nox_parse_thing_extent}, {name: CString("LIGHTINTENSITY"), parse_func: nox_parse_thing_light_intensity}, {name: CString("DRAW"), parse_func: nox_parse_thing_draw}, {name: CString("Z"), parse_func: nox_parse_thing_z}, {name: CString("ZSIZE"), parse_func: nox_parse_thing_zsize}, {name: CString("SIZE"), parse_func: nox_parse_thing_size}, {name: CString("MENUICON"), parse_func: nox_parse_thing_menu_icon}, {name: CString("LIGHTCOLOR"), parse_func: nox_parse_thing_light_color}, {name: CString("LIGHTDIRECTION"), parse_func: nox_parse_thing_light_dir}, {name: CString("LIGHTPENUMBRA"), parse_func: nox_parse_thing_light_penumbra}, {name: CString("AUDIOLOOP"), parse_func: nox_parse_thing_audio_loop}, {name: CString("CLIENTUPDATE"), parse_func: nox_parse_thing_client_update}, {name: CString("LIFETIME"), parse_func: nox_parse_thing_lifetime}, {name: CString("WEIGHT"), parse_func: nox_parse_thing_weight}, {name: CString("PRETTYNAME"), parse_func: nox_parse_thing_pretty_name}, {name: CString("DESCRIPTION"), parse_func: nox_parse_thing_desc}, {name: CString("PRETTYIMAGE"), parse_func: nox_parse_thing_pretty_image}, {name: CString("HEALTH"), parse_func: nox_parse_thing_health}}
var nox_parse_thing_funcs_cnt int32 = int32(uint32(unsafe.Sizeof(nox_parse_thing_funcs)) / uint32(unsafe.Sizeof(nox_parse_thing_funcs_t{})))

func nox_parse_thing_flags(obj *nox_thing, f *nox_memfile, attr_value *byte) bool {
	obj.flags = 0
	set_bitmask_flags_from_plus_separated_names_423930(attr_value, (*uint32)(unsafe.Pointer(&obj.flags)), (**byte)(memmap.PtrOff(0x587000, 114076)))
	return true
}
func nox_parse_thing_class(obj *nox_thing, f *nox_memfile, attr_value *byte) bool {
	obj.pri_class = 0
	set_bitmask_flags_from_plus_separated_names_423930(attr_value, &obj.pri_class, (**byte)(memmap.PtrOff(0x587000, 114208)))
	return true
}
func nox_parse_thing_subclass(obj *nox_thing, f *nox_memfile, attr_value *byte) bool {
	obj.sub_class = 0
	set_bitmask_flags_from_plus_separated_names_multiple_423A10(attr_value, &obj.sub_class)
	return true
}
func nox_parse_thing_extent(obj *nox_thing, f *nox_memfile, attr_value *byte) bool {
	var (
		s   nox_shape
		res bool = nox_parse_shape(&s, attr_value) != 0
	)
	obj.shape_kind = uint16(int16(s.kind))
	obj.shape_r = s.circle_r
	obj.shape_w = s.box_w
	obj.shape_h = s.box_h
	return res
}
func nox_parse_thing_light_intensity(obj *nox_thing, f *nox_memfile, attr_value *byte) bool {
	stdio.Sscanf(attr_value, "%f", &obj.light_intensity)
	return true
}
func nox_parse_thing_draw(obj *nox_thing, f *nox_memfile, attr_value *byte) bool {
	var (
		read_len uint8 = f.ReadU8()
		read_str [256]byte
	)
	nox_memfile_read(unsafe.Pointer(&read_str[0]), 1, int32(read_len), f)
	read_str[read_len] = 0
	var tmp uint32
	nox_memfile_read64align_40AD60((*byte)(unsafe.Pointer(&tmp)), int32(unsafe.Sizeof(uint32(0))), 1, f)
	if *(*uint32)(unsafe.Pointer(&nox_parse_thing_draw_funcs[0])) == 0 {
		return true
	}
	var item *nox_parse_thing_draw_funcs_t = nil
	for i := int32(0); i < nox_parse_thing_draw_funcs_cnt; i++ {
		var cur *nox_parse_thing_draw_funcs_t = &nox_parse_thing_draw_funcs[i]
		if libc.StrCmp(cur.name, &read_str[0]) == 0 {
			item = cur
			break
		}
	}
	if item == nil {
		return true
	}
	if item.parse_fnc != nil {
		item.parse_fnc(obj, f, attr_value)
	}
	obj.draw_func = cgoAsFunc(item.draw, (*func(*uint32, *nox_drawable) int32)(nil)).(func(*uint32, *nox_drawable) int32)
	return true
}
func nox_parse_thing_z(obj *nox_thing, f *nox_memfile, attr_value *byte) bool {
	var z int32 = 0
	stdio.Sscanf(attr_value, "%d", &z)
	obj.z = uint16(int16(z))
	return true
}
func nox_parse_thing_zsize(obj *nox_thing, f *nox_memfile, attr_value *byte) bool {
	var (
		min int32 = 0
		max int32 = 0
	)
	stdio.Sscanf(attr_value, "%d %d", &min, &max)
	if max < min {
		max = min
	}
	obj.zsize_min = float32(float64(min))
	obj.zsize_max = float32(float64(max))
	return true
}
func nox_parse_thing_size(obj *nox_thing, f *nox_memfile, attr_value *byte) bool {
	var (
		w int32 = 0
		h int32 = 0
	)
	stdio.Sscanf(attr_value, "%d %d", &w, &h)
	obj.hwidth = uint8(int8(w / 2))
	obj.hheight = uint8(int8(h / 2))
	return true
}
func nox_parse_thing_menu_icon(obj *nox_thing, f *nox_memfile, attr_value *byte) bool {
	obj.menuicon = f.ReadU32()
	if obj.menuicon == math.MaxUint32 {
		f.Skip(1)
		var n int32 = int32(f.ReadU8())
		f.Skip(n)
	}
	return true
}
func nox_parse_thing_light_color(obj *nox_thing, f *nox_memfile, attr_value *byte) bool {
	var (
		r int32
		g int32
		b int32
	)
	if stdio.Sscanf(attr_value, "%d %d %d", &r, &g, &b) != 3 {
		return false
	}
	obj.field_f = 2
	if r <= math.MaxUint8 {
		obj.light_color_r = r
	} else {
		obj.light_color_r = math.MaxUint8
	}
	if g <= math.MaxUint8 {
		obj.light_color_g = g
	} else {
		obj.light_color_g = math.MaxUint8
	}
	if b <= math.MaxUint8 {
		obj.light_color_b = b
	} else {
		obj.light_color_b = math.MaxUint8
	}
	return true
}
func nox_parse_thing_light_dir(obj *nox_thing, f *nox_memfile, attr_value *byte) bool {
	var deg int32 = 0
	if stdio.Sscanf(attr_value, "%d", &deg) != 1 {
		return false
	}
	if deg < 0 || deg >= 360 {
		return false
	}
	obj.light_dir = uint16(int16(int64(float64(deg)**mem_getDoublePtr(0x581450, 9560)**(*float64)(unsafe.Pointer(&qword_581450_9552)) + *(*float64)(unsafe.Pointer(&qword_581450_9544)))))
	obj.field_10 = 0
	return true
}
func nox_parse_thing_light_penumbra(obj *nox_thing, f *nox_memfile, attr_value *byte) bool {
	var deg int32 = 0
	if stdio.Sscanf(attr_value, "%d", &deg) != 1 {
		return false
	}
	if deg < 0 || deg >= 180 {
		return false
	}
	obj.light_penumbra = uint16(int16(int64(float64(deg)**mem_getDoublePtr(0x581450, 9560)**(*float64)(unsafe.Pointer(&qword_581450_9552)) + *(*float64)(unsafe.Pointer(&qword_581450_9544)))))
	return true
}
func nox_parse_thing_audio_loop(obj *nox_thing, f *nox_memfile, attr_value *byte) bool {
	obj.audio_loop = uint32(nox_xxx_utilFindSound_40AF50(attr_value))
	return true
}
func nox_parse_thing_client_update(obj *nox_thing, f *nox_memfile, attr_value *byte) bool {
	var (
		v3 *byte
		v4 *byte
		v5 int32
		v6 *uint8
	)
	v3 = libc.StrTok(attr_value, CString(" \t\n\r"))
	v4 = *(**byte)(memmap.PtrOff(0x587000, 175072))
	v5 = 0
	if *memmap.PtrUint32(0x587000, 175072) != 0 {
		v6 = (*uint8)(memmap.PtrOff(0x587000, 175072))
		for {
			if libc.StrCmp(v4, v3) == 0 {
				break
			}
			v4 = (*byte)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v6))), unsafe.Sizeof(uint32(0))*2))))))
			v6 = (*uint8)(unsafe.Add(unsafe.Pointer(v6), 8))
			v5++
			if v4 == nil {
				break
			}
		}
	}
	if *memmap.PtrUint32(0x587000, uint32(v5*8)+0x2ABE0) == 0 {
		return false
	}
	obj.client_update = *memmap.PtrUint32(0x587000, uint32(v5*8)+0x2ABE4)
	return true
}
func nox_parse_thing_lifetime(obj *nox_thing, f *nox_memfile, attr_value *byte) bool {
	var v int32 = 0
	stdio.Sscanf(attr_value, "%d", &v)
	obj.lifetime = v
	return true
}
func nox_parse_thing_weight(obj *nox_thing, f *nox_memfile, attr_value *byte) bool {
	var v int32 = 0
	stdio.Sscanf(attr_value, "%d", &v)
	obj.weight = uint8(int8(v))
	obj.pri_class |= 0x80000000
	return true
}
func nox_parse_thing_pretty_name(obj *nox_thing, f *nox_memfile, attr_value *byte) bool {
	obj.pretty_name = strMan.GetStringInFile(attr_value, "C:\\NoxPost\\src\\Client\\Drawable\\drawdb.c")
	return true
}
func nox_parse_thing_desc(obj *nox_thing, f *nox_memfile, attr_value *byte) bool {
	obj.desc = strMan.GetStringInFile(attr_value, "C:\\NoxPost\\src\\Client\\Drawable\\drawdb.c")
	return true
}
func nox_parse_thing_pretty_image(obj *nox_thing, f *nox_memfile, attr_value *byte) bool {
	var (
		v10       [128]byte
		known_idx uint32 = f.ReadU32()
	)
	if known_idx != math.MaxUint32 {
		obj.pretty_image = uint32(uintptr(unsafe.Pointer(nox_xxx_readImgMB_42FAA0(int32(known_idx), 0, &v10[0]))))
		return true
	}
	var v8 int32 = int32(f.ReadU8())
	var n int32 = int32(f.ReadU8())
	nox_memfile_read(unsafe.Pointer(&v10[0]), 1, n, f)
	obj.pretty_image = uint32(uintptr(unsafe.Pointer(nox_xxx_readImgMB_42FAA0(int32(known_idx), int8(v8), &v10[0]))))
	return true
}
func nox_parse_thing_health(obj *nox_thing, f *nox_memfile, attr_value *byte) bool {
	obj.health = uint16(int16(libc.Atoi(libc.GoString(attr_value))))
	return true
}
func nox_parse_thing(thing_file *nox_memfile, scratch_buffer *byte, thing *nox_thing) int32 {
	var entry_len uint8
	for int32(func() uint8 {
		entry_len = thing_file.ReadU8()
		return entry_len
	}()) != 0 {
		nox_memfile_read(unsafe.Pointer(scratch_buffer), 1, int32(entry_len), thing_file)
		*(*byte)(unsafe.Add(unsafe.Pointer(scratch_buffer), entry_len)) = 0
		var attr_name *byte = libc.StrTok(scratch_buffer, CString(" \t\n\r"))
		for i := int32(0); i < nox_parse_thing_funcs_cnt; i++ {
			var attr_parser *nox_parse_thing_funcs_t = &nox_parse_thing_funcs[i]
			if libc.StrCmp(attr_name, attr_parser.name) != 0 {
				continue
			}
			var attr_value *byte = libc.StrTok(nil, CString("="))
			if attr_value != nil {
				libc.MemMove(unsafe.Pointer(scratch_buffer), unsafe.Add(unsafe.Pointer(attr_value), 1), libc.StrLen((*byte)(unsafe.Add(unsafe.Pointer(attr_value), 1)))+1)
			}
			attr_parser.parse_func(thing, thing_file, scratch_buffer)
			break
		}
	}
	return 1
}
func nox_xxx_spriteDefByAlphabetClear_44CCA0() int32 {
	var result int32
	result = 0
	libc.MemSet(memmap.PtrOff(6112660, 830616), 0, 108)
	libc.MemSet(memmap.PtrOff(6112660, 830724), 0, 108)
	libc.MemSet(memmap.PtrOff(6112660, 830296), 0, 108)
	return result
}
func sub_485CF0() int32 {
	var (
		v0 int32
		v1 *unsafe.Pointer
	)
	v0 = 0
	if *(*int32)(unsafe.Pointer(&dword_5d4594_251568)) <= 0 {
		return 1
	}
	v1 = (*unsafe.Pointer)(memmap.PtrOff(0x85B3FC, 32516))
	for {
		if *v1 != nil {
			*v1 = nil
			*v1 = nil
		}
		v0++
		v1 = (*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(unsafe.Pointer(nil))*15))
		if v0 >= *(*int32)(unsafe.Pointer(&dword_5d4594_251568)) {
			break
		}
	}
	return 1
}
func sub_485F30() int32 {
	var (
		v0 int32
		v1 *unsafe.Pointer
	)
	v0 = 0
	if *(*int32)(unsafe.Pointer(&dword_5d4594_251572)) <= 0 {
		return 1
	}
	v1 = (*unsafe.Pointer)(memmap.PtrOff(0x85B3FC, 28676))
	for {
		if *v1 != nil {
			*v1 = nil
			*v1 = nil
		}
		v0++
		v1 = (*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(unsafe.Pointer(nil))*15))
		if v0 >= *(*int32)(unsafe.Pointer(&dword_5d4594_251572)) {
			break
		}
	}
	return 1
}
func sub_46A360() int32 {
	var (
		i  int32
		v1 int32
		v2 int32
		v3 *uint8
	)
	for i = 0; uint32(i) < 0xF0DC0; i += 0x302C {
		v1 = 0
		for {
			v2 = 15
			for {
				v3 = (*uint8)(memmap.PtrOff(0x85B3FC, uint32(i+0xC914+v1)))
				v1 += 64
				v2--
				libc.MemSet(unsafe.Pointer(v3), 0, 64)
				if v2 == 0 {
					break
				}
			}
			if v1 >= 3840 {
				break
			}
		}
	}
	return 1
}
func nox_xxx_spriteDefByAlphabetAlloc_44CCD0() uint32 {
	var (
		i      int32
		result uint32
	)
	for i = 0; i < 108; i += 4 {
		result = *memmap.PtrUint32(6112660, uint32(i)+0xCAC98)
		if result != 0 {
			result = uint32(uintptr(alloc.Calloc(1, int(result*8))))
			*memmap.PtrUint32(6112660, uint32(i)+0xCAB58) = result
		} else {
			*memmap.PtrUint32(6112660, uint32(i)+0xCAB58) = 0
		}
		*memmap.PtrUint32(6112660, uint32(i)+0xCAD04) = 0
	}
	return result
}
func nox_xxx_spriteDefByAlphabetAdd_0_44CD60(a1 *nox_thing, a2 int32) {
	if a1 == nil {
		return
	}
	var v2 int32 = nox_xxx_keyFirstLetterNumberCli_44CD30(a1.name)
	if v2 < 0 {
		return
	}
	var v3 int32 = int32(*memmap.PtrUint32(6112660, uint32(v2*4)+0xCAB58))
	if v3 == 0 {
		return
	}
	var v4 int32 = int32(*memmap.PtrUint32(6112660, uint32(v2*4)+0xCAD04))
	*(*uint32)(unsafe.Pointer(uintptr(v3 + v4*8))) = uint32(uintptr(unsafe.Pointer(a1)))
	*(*uint32)(unsafe.Pointer(uintptr(v3 + v4*8 + 4))) = uint32(a2)
	*memmap.PtrUint32(6112660, uint32(v2*4)+0xCAD04)++
}
func nox_xxx_spriteDefByAlphabetCompare_44CDE0(a1 unsafe.Pointer, a2 unsafe.Pointer) int32 {
	return int32(libc.StrCaseCmp(**(***byte)(a1), **(***byte)(a2)))
}
func nox_xxx_spriteDefByAlphabetSort_44CDB0() {
	var (
		i  int32
		v1 int32
	)
	for i = 0; i < 108; i += 4 {
		v1 = int32(*memmap.PtrUint32(6112660, uint32(i)+0xCAC98))
		if v1 > 1 {
			libc.Sort(*(*unsafe.Pointer)(memmap.PtrOff(6112660, uint32(i)+0xCAB58)), uint32(v1), 8, nox_xxx_spriteDefByAlphabetCompare_44CDE0)
		}
	}
}
func nox_xxx_equipWeapon_4131A0() *byte {
	var (
		result *byte
		i      int32
		j      int32
	)
	result = *(**byte)(memmap.PtrOff(6112660, 251616))
	if *memmap.PtrUint32(6112660, 251616) != 1 {
		for i = int32(dword_5d4594_251600); i != 0; i = int32(*(*uint32)(unsafe.Pointer(uintptr(i + 80)))) {
			if noxflags.HasGame(noxflags.GameHost | noxflags.GameFlag22) {
				result = (*byte)(unsafe.Pointer(uintptr(nox_xxx_getNameId_4E3AA0(*(**byte)(unsafe.Pointer(uintptr(i)))))))
			} else {
				result = (*byte)(unsafe.Pointer(uintptr(bool2int(noxflags.HasGame(noxflags.GameClient)))))
				if result == nil {
					return result
				}
				result = (*byte)(unsafe.Pointer(uintptr(nox_xxx_getTTByNameSpriteMB_44CFC0(*(**byte)(unsafe.Pointer(uintptr(i)))))))
			}
			*(*uint32)(unsafe.Pointer(uintptr(i + 4))) = uint32(uintptr(unsafe.Pointer(result)))
		}
		for j = int32(dword_5d4594_251608); j != 0; j = int32(*(*uint32)(unsafe.Pointer(uintptr(j + 80)))) {
			if noxflags.HasGame(noxflags.GameHost | noxflags.GameFlag22) {
				result = (*byte)(unsafe.Pointer(uintptr(nox_xxx_getNameId_4E3AA0(*(**byte)(unsafe.Pointer(uintptr(j)))))))
			} else {
				result = (*byte)(unsafe.Pointer(uintptr(bool2int(noxflags.HasGame(noxflags.GameClient)))))
				if result == nil {
					return result
				}
				result = (*byte)(unsafe.Pointer(uintptr(nox_xxx_getTTByNameSpriteMB_44CFC0(*(**byte)(unsafe.Pointer(uintptr(j)))))))
			}
			*(*uint32)(unsafe.Pointer(uintptr(j + 4))) = uint32(uintptr(unsafe.Pointer(result)))
		}
		*memmap.PtrUint32(6112660, 251616) = 1
	}
	return result
}
func nox_xxx_equipArmor_415AB0() {
	var (
		v0 *uint8
		v1 int32
		v2 int32
	)
	if *memmap.PtrUint32(6112660, 371252) != 1 {
		if *memmap.PtrUint32(0x587000, 34848) != 0 {
			v0 = (*uint8)(memmap.PtrOff(0x587000, 34864))
			for {
				if noxflags.HasGame(noxflags.GameHost | noxflags.GameFlag22) {
					*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v0))), -int(unsafe.Sizeof(uint32(0))*2)))) = uint32(nox_xxx_getNameId_4E3AA0(*((**byte)(unsafe.Add(unsafe.Pointer((**byte)(unsafe.Pointer(v0))), -int(unsafe.Sizeof((*byte)(nil))*4))))))
					v1 = nox_xxx_getNameId_4E3AA0(*(**byte)(unsafe.Pointer(v0)))
				} else {
					*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v0))), -int(unsafe.Sizeof(uint32(0))*2)))) = uint32(nox_xxx_getTTByNameSpriteMB_44CFC0(*((**byte)(unsafe.Add(unsafe.Pointer((**byte)(unsafe.Pointer(v0))), -int(unsafe.Sizeof((*byte)(nil))*4))))))
					v1 = nox_xxx_getTTByNameSpriteMB_44CFC0(*(**byte)(unsafe.Pointer(v0)))
				}
				*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v0))), unsafe.Sizeof(uint32(0))*1))) = uint32(v1)
				v2 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v0))), unsafe.Sizeof(uint32(0))*2))))
				v0 = (*uint8)(unsafe.Add(unsafe.Pointer(v0), 24))
				if v2 == 0 {
					break
				}
			}
		}
		*memmap.PtrUint32(6112660, 371252) = 1
	}
}
func nox_xxx_equipWeapon_4157C0() {
	var (
		v0 *uint8
		v1 int32
		v2 int32
	)
	if *memmap.PtrUint32(6112660, 371244) != 1 {
		if *memmap.PtrUint32(0x587000, 33064) != 0 {
			v0 = (*uint8)(memmap.PtrOff(0x587000, 33064))
			for {
				if noxflags.HasGame(noxflags.GameHost | noxflags.GameFlag22) {
					v1 = nox_xxx_getNameId_4E3AA0(*(**byte)(unsafe.Pointer(v0)))
				} else {
					v1 = nox_xxx_getTTByNameSpriteMB_44CFC0(*(**byte)(unsafe.Pointer(v0)))
				}
				*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v0))), unsafe.Sizeof(uint32(0))*1))) = uint32(v1)
				v2 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v0))), unsafe.Sizeof(uint32(0))*3))))
				v0 = (*uint8)(unsafe.Add(unsafe.Pointer(v0), 12))
				if v2 == 0 {
					break
				}
			}
		}
		*memmap.PtrUint32(6112660, 371244) = 1
	}
}
func nox_get_things_count() int32 {
	return nox_things_count
}
func sub_42BF10() int32 {
	var v1 int32
	if dword_5d4594_741676 != 0 {
		return 1
	}
	if noxflags.HasGame(noxflags.GameHost | noxflags.GameFlag22) {
		v1 = nox_xxx_unitDefGetCount_4E3AC0()
	} else {
		if !noxflags.HasGame(noxflags.GameClient) {
			return 0
		}
		v1 = nox_get_things_count()
	}
	dword_5d4594_741680 = uint32(v1)
	dword_5d4594_741676 = uint32(uintptr(alloc.Calloc(int(v1), 2)))
	if dword_5d4594_741676 != 0 {
		sub_42BFB0()
		return 1
	}
	return 0
}
func nox_xxx_parseThingBinClient_44C840_read_things() unsafe.Pointer {
	var scratch_buffer *byte = (*byte)(alloc.Calloc(256*1024, 1))
	nox_xxx_spriteDefByAlphabetClear_44CCA0()
	nox_things_free_44C580()
	nox_things_count = 1
	sub_485CF0()
	sub_485F30()
	sub_46A360()
	var things *nox_memfile = nox_open_thing_bin()
	if things == nil {
		return nil
	}
	nox_memfile_seek_40AD10(things, 0, 0)
	var entry_type int32
	for nox_memfile_read(unsafe.Pointer(&entry_type), 4, 1, things) != 0 {
		switch entry_type {
		case 0x5350454C:
			nox_thing_skip_spells_415100(int32(uintptr(unsafe.Pointer(things))))
		case 0x4142494C:
			nox_thing_read_ability_415320(int32(uintptr(unsafe.Pointer(things))))
		case 0x41554420:
			nox_thing_skip_AUD_414D40(things)
		case 0x41564E54:
			nox_thing_skip_AVNT_452B00(things)
		case 0x57414C4C:
			if nox_thing_read_wall_46A010((*uint32)(unsafe.Pointer(things)), scratch_buffer) == 0 {
				return nil
			}
		case 0x464C4F52:
			if nox_thing_read_floor_485B30(int32(uintptr(unsafe.Pointer(things))), scratch_buffer) == 0 {
				return nil
			}
		case 0x45444745:
			if nox_thing_read_edge_485D40(int32(uintptr(unsafe.Pointer(things))), scratch_buffer) == 0 {
				return nil
			}
		case 0x494D4147:
			nox_thing_read_image_415240(int32(uintptr(unsafe.Pointer(things))))
		case 0x54484E47:
			var obj *nox_thing = new(nox_thing)
			if obj == nil {
				return nil
			}
			var v28 uint8 = things.ReadU8()
			nox_memfile_read(unsafe.Pointer(scratch_buffer), 1, int32(v28), things)
			*(*byte)(unsafe.Add(unsafe.Pointer(scratch_buffer), v28)) = 0
			obj.name = nox_clone_str(scratch_buffer)
			obj.menuicon = math.MaxUint32
			obj.field_1c = func() int32 {
				p := &nox_things_count
				x := *p
				*p++
				return x
			}()
			obj.flags |= 0x1000000
			*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(obj))), 15))) = 0
			obj.field_10 = math.MaxUint16
			obj.audio_loop = 0
			obj.draw_func = func(arg1 *uint32, arg2 *nox_drawable) int32 {
				return nox_thing_debug_draw((*nox_draw_viewport_t)(unsafe.Pointer(arg1)), arg2)
			}
			obj.zsize_min = 0
			obj.zsize_max = 30.0
			if nox_parse_thing(things, scratch_buffer, obj) == 0 {
				return nil
			}
			obj.next = nox_things_head
			nox_things_head = obj
			nox_xxx_spriteDefByAlphabetAdd_44CD10(obj.name)
		}
	}
	*memmap.PtrUint32(0x85B3FC, 4) = 1
	if nox_loaded_thing_bin != nil {
		if noxflags.HasGame(noxflags.GameHost) && *memmap.PtrUint32(0x85B3FC, 960) != 0 {
			nox_free_thing_bin()
		}
	} else {
		nox_memfile_free(things)
	}
	var result unsafe.Pointer = unsafe.Pointer(&make([]*nox_thing, int(nox_things_count))[0])
	nox_things_array = (**nox_thing)(result)
	if result != nil {
		nox_xxx_spriteDefByAlphabetAlloc_44CCD0()
		var cur *nox_thing = nox_things_head
		for i := int32(1); i < nox_things_count; i++ {
			*(**nox_thing)(unsafe.Add(unsafe.Pointer(nox_things_array), unsafe.Sizeof((*nox_thing)(nil))*uintptr(nox_things_count-i))) = cur
			nox_xxx_spriteDefByAlphabetAdd_0_44CD60(cur, nox_things_count-i)
			if int32(*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(cur))), 14)))) != 0 {
				if cur.pretty_name == nil {
					libc.StrCpy((*byte)(memmap.PtrOff(6112660, 830404)), CString("thing.db:"))
					var cur_name *byte = cur.name
					var cur_name_len_plus_one uint32 = uint32(libc.StrLen(cur.name) + 1)
					var v11 *uint8 = (*uint8)(memmap.PtrOff(6112660, uint32(libc.StrLen((*byte)(memmap.PtrOff(6112660, 830404)))+0xCABC4)))
					libc.MemCpy(unsafe.Pointer(v11), unsafe.Pointer(cur.name), int((cur_name_len_plus_one>>2)*4))
					var v13 *byte = (*byte)(unsafe.Add(unsafe.Pointer(cur_name), (cur_name_len_plus_one>>2)*4))
					var v12 *uint8 = (*uint8)(unsafe.Add(unsafe.Pointer(v11), (cur_name_len_plus_one>>2)*4))
					var v14 int8 = int8(uint8(cur_name_len_plus_one))
					*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&cur_name_len_plus_one))), unsafe.Sizeof(uint16(0))*0)) = *memmap.PtrUint16(0x587000, 122728)
					libc.MemCpy(unsafe.Pointer(v12), unsafe.Pointer(v13), int(int32(v14)&3))
					var v15 *uint8 = (*uint8)(memmap.PtrOff(6112660, uint32(libc.StrLen((*byte)(memmap.PtrOff(6112660, 830404)))+0xCABC5)))
					var v16 int32 = int32(*memmap.PtrUint32(0x587000, 122724))
					*(*uint32)(unsafe.Pointer(func() *uint8 {
						p := &v15
						*p = (*uint8)(unsafe.Add(unsafe.Pointer(*p), -1))
						return *p
					}())) = *memmap.PtrUint32(0x587000, 122720)
					var v17 uint8 = (*memmap.PtrUint8(0x587000, 122730))
					*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v15))), unsafe.Sizeof(uint32(0))*1))) = uint32(v16)
					*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v15))), unsafe.Sizeof(uint16(0))*4))) = uint16(cur_name_len_plus_one)
					*(*uint8)(unsafe.Add(unsafe.Pointer(v15), 10)) = v17
					cur.pretty_name = strMan.GetStringInFile((*byte)(memmap.PtrOff(6112660, 830404)), "C:\\NoxPost\\src\\Client\\Drawable\\drawdb.c")
				}
				if cur.desc == nil {
					libc.StrCpy((*byte)(memmap.PtrOff(6112660, 830404)), CString("thing.db:"))
					var v18 *byte = cur.name
					var v19 uint32 = uint32(libc.StrLen(cur.name) + 1)
					var v20 *uint8 = (*uint8)(memmap.PtrOff(6112660, uint32(libc.StrLen((*byte)(memmap.PtrOff(6112660, 830404)))+0xCABC4)))
					libc.MemCpy(unsafe.Pointer(v20), unsafe.Pointer(cur.name), int((v19>>2)*4))
					var v22 *byte = (*byte)(unsafe.Add(unsafe.Pointer(v18), (v19>>2)*4))
					var v21 *uint8 = (*uint8)(unsafe.Add(unsafe.Pointer(v20), (v19>>2)*4))
					var v23 int8 = int8(uint8(v19))
					var v24 int32 = int32(*memmap.PtrUint32(0x587000, 122792))
					libc.MemCpy(unsafe.Pointer(v21), unsafe.Pointer(v22), int(int32(v23)&3))
					var v25 *uint8 = (*uint8)(memmap.PtrOff(6112660, uint32(libc.StrLen((*byte)(memmap.PtrOff(6112660, 830404)))+0xCABC5)))
					var v26 int32 = int32(*memmap.PtrUint32(0x587000, 122788))
					*(*uint32)(unsafe.Pointer(func() *uint8 {
						p := &v25
						*p = (*uint8)(unsafe.Add(unsafe.Pointer(*p), -1))
						return *p
					}())) = *memmap.PtrUint32(0x587000, 122784)
					*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v25))), unsafe.Sizeof(uint32(0))*1))) = uint32(v26)
					*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v25))), unsafe.Sizeof(uint32(0))*2))) = uint32(v24)
					cur.desc = strMan.GetStringInFile((*byte)(memmap.PtrOff(6112660, 830404)), "C:\\NoxPost\\src\\Client\\Drawable\\drawdb.c")
				}
			}
			cur = cur.next
		}
		nox_xxx_spriteDefByAlphabetSort_44CDB0()
		alloc.Free(unsafe.Pointer(scratch_buffer))
		nox_xxx_equipWeapon_4131A0()
		nox_xxx_equipArmor_415AB0()
		nox_xxx_equipWeapon_4157C0()
		result = unsafe.Pointer(uintptr(bool2int(sub_42BF10() != 0)))
	}
	return result
}
func sub_44C620_free() {
	var v0 *unsafe.Pointer
	v0 = (*unsafe.Pointer)(memmap.PtrOff(6112660, 830296))
	for {
		if *v0 != nil {
			*v0 = nil
		}
		*v0 = nil
		v0 = (*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(v0), unsafe.Sizeof(unsafe.Pointer(nil))*1))
		if int32(uintptr(unsafe.Pointer(v0))) >= int32(uintptr(memmap.PtrOff(6112660, 830404))) {
			break
		}
	}
}
func nox_xxx_unitDefByAlphabetFree_4E2B30() {
	var (
		v0     *unsafe.Pointer
		result unsafe.Pointer
	)
	_ = result
	v0 = (*unsafe.Pointer)(memmap.PtrOff(6112660, 1563348))
	for {
		result = *v0
		if *v0 != nil {
			*v0 = nil
		}
		*v0 = nil
		v0 = (*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(v0), unsafe.Sizeof(unsafe.Pointer(nil))*1))
		if int32(uintptr(unsafe.Pointer(v0))) >= int32(uintptr(memmap.PtrOff(6112660, 1563456))) {
			break
		}
	}
}
func nox_xxx_free_42BF80() {
	if dword_5d4594_741676 != 0 {
		*(*unsafe.Pointer)(unsafe.Pointer(&dword_5d4594_741676)) = nil
		dword_5d4594_741676 = 0
	}
	dword_5d4594_741680 = 0
}
func nox_xxx_freeObjectTypes_4E2A20() int32 {
	var next *nox_objectType_t = nil
	for typ := (*nox_objectType_t)(nox_xxx_objectTypes_head_1563660); typ != nil; typ = next {
		next = typ.next
		if typ.id != nil {
			alloc.Free(unsafe.Pointer(typ.id))
		}
		if typ.data_34 != nil {
			typ.data_34 = nil
		}
		if typ.collide_data != nil {
			typ.collide_data = nil
		}
		if typ.die_data != nil {
			typ.die_data = nil
		}
		if typ.init_data != nil {
			typ.init_data = nil
		}
		var v2 unsafe.Pointer = typ.data_update
		if v2 != nil {
			if typ.obj_class&2 != 0 {
				var v3 unsafe.Pointer = *(*unsafe.Pointer)(unsafe.Pointer(uintptr(uint32(uintptr(v2)) + 476)))
				if v3 != nil {
					v3 = nil
				}
			}
			typ.data_update = nil
		}
		if typ.use_data != nil {
			typ.use_data = nil
		}
		alloc.Free(unsafe.Pointer(typ))
	}
	nox_xxx_objectTypes_head_1563660 = nil
	if *memmap.PtrUint32(6112660, 1563456) != 0 {
		*(*unsafe.Pointer)(memmap.PtrOff(6112660, 1563456)) = nil
		*memmap.PtrUint32(6112660, 1563456) = 0
	}
	*memmap.PtrUint32(0x587000, 201384) = 1
	nox_xxx_unitDefByAlphabetFree_4E2B30()
	nox_xxx_free_42BF80()
	return 1
}
func sub_4E3010() int32 {
	var result int32
	result = 0
	libc.MemSet(memmap.PtrOff(6112660, 1563668), 0, 108)
	libc.MemSet(memmap.PtrOff(6112660, 1563776), 0, 108)
	libc.MemSet(memmap.PtrOff(6112660, 1563348), 0, 108)
	return result
}
func nox_xxx_unitDefByAlphabetAdd_4E3080(a1 *byte) *byte {
	var result *byte
	result = a1
	if a1 != nil {
		result = (*byte)(unsafe.Pointer(uintptr(nox_xxx_keyFirstLetterNumber_4E30A0(a1))))
		if int32(uintptr(unsafe.Pointer(result))) >= 0 {
			*memmap.PtrUint32(6112660, uint32(uintptr(unsafe.Pointer(result)))*4+0x17DC14)++
		}
	}
	return result
}
func nox_xxx_objectTypes_allFit_4E3110() int32 {
	var ret int32 = 1
	for typ := (*nox_objectType_t)(nox_xxx_getFirstObjectType_4E3B30()); typ != nil; typ = nox_xxx_objectType_next_4E3B40(typ) {
		if (typ.obj_flags & 64) != 0 {
			continue
		}
		var shape *nox_shape = &typ.shape
		if shape.kind == nox_shape_kind(NOX_SHAPE_CIRCLE) {
			if float64(shape.circle_r+shape.circle_r) >= 85.0 {
				ret = 0
			}
		} else if shape.kind == nox_shape_kind(NOX_SHAPE_BOX) {
			nox_shape_box_calc(shape)
			if float64(shape.box_right_top-shape.box_left_bottom_2) >= 85.0 || float64(shape.box_right_top_2-shape.box_left_bottom) >= 85.0 {
				ret = 0
			}
		}
	}
	return ret
}
func nox_xxx_unitDefByAlphabetInit_4E3040() {
	for i := int32(0); i < 27; i++ {
		var cnt int32 = int32(*memmap.PtrUint32(6112660, uint32(i*4)+0x17DC14))
		if cnt != 0 {
			*memmap.PtrUint32(6112660, uint32(i*4)+0x17DAD4) = uint32(uintptr(alloc.Calloc(1, int(cnt*4))))
		} else {
			*memmap.PtrUint32(6112660, uint32(i*4)+0x17DAD4) = 0
		}
		*memmap.PtrUint32(6112660, uint32(i*4)+0x17DC80) = 0
	}
}
func nox_xxx_objectTypeAddToNameInd_4E30D0(a1 int32) {
	var (
		v1 int32
		v2 int32
	)
	if a1 != 0 {
		v1 = nox_xxx_keyFirstLetterNumber_4E30A0(*(**byte)(unsafe.Pointer(uintptr(a1 + 4))))
		v2 = int32(*memmap.PtrUint32(6112660, uint32(v1*4)+0x17DAD4))
		if v2 != 0 {
			*(*uint32)(unsafe.Pointer(uintptr(uint32(v2) + *memmap.PtrUint32(6112660, uint32(v1*4)+0x17DC80)*4))) = uint32(a1)
			*memmap.PtrUint32(6112660, uint32(v1*4)+0x17DC80)++
		}
	}
}
func sub_4E2A00(a1 unsafe.Pointer, a2 unsafe.Pointer) int32 {
	return int32(libc.StrCaseCmp(*(**byte)(unsafe.Pointer(uintptr(*(*uint32)(a1) + 4))), *(**byte)(unsafe.Pointer(uintptr(*(*uint32)(a2) + 4)))))
}
func sub_4E29D0() {
	var (
		i  int32
		v1 int32
	)
	for i = 0; i < 108; i += 4 {
		v1 = int32(*memmap.PtrUint32(6112660, uint32(i)+0x17DC14))
		if v1 > 1 {
			libc.Sort(*(*unsafe.Pointer)(memmap.PtrOff(6112660, uint32(i)+0x17DAD4)), uint32(v1), 4, sub_4E2A00)
		}
	}
}
func sub_4F0640() *byte {
	var (
		v0     *byte
		v1     *uint8
		v2     int8
		v3     *byte
		v4     *byte
		v5     *uint8
		v6     int32
		v7     *byte
		v8     *uint8
		v9     int32
		v10    *byte
		v11    *uint8
		v12    int32
		result *byte
		v14    *uint8
		v15    int32
	)
	v0 = *(**byte)(memmap.PtrOff(0x587000, 208180))
	if *memmap.PtrUint32(0x587000, 208180) != 0 {
		v1 = (*uint8)(memmap.PtrOff(0x587000, 208180))
		for {
			v2 = int8(*v0)
			v3 = *(**byte)(unsafe.Pointer(v1))
			if int32(v2) == 35 {
				v3 = (*byte)(unsafe.Add(unsafe.Pointer(v3), 1))
			}
			*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v1))), unsafe.Sizeof(uint32(0))*1))) = uint32(nox_xxx_getNameId_4E3AA0(v3))
			v0 = (*byte)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v1))), unsafe.Sizeof(uint32(0))*5))))))
			v1 = (*uint8)(unsafe.Add(unsafe.Pointer(v1), 20))
			if v0 == nil {
				break
			}
		}
	}
	v4 = *(**byte)(memmap.PtrOff(0x587000, 210712))
	if *memmap.PtrUint32(0x587000, 210712) != 0 {
		v5 = (*uint8)(memmap.PtrOff(0x587000, 210712))
		for {
			v6 = nox_xxx_modifGetIdByName_413290(v4)
			*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v5))), -int(unsafe.Sizeof(uint32(0))*1)))) = uint32(uintptr(unsafe.Pointer(nox_xxx_modifGetDescById_413330(v6))))
			v4 = (*byte)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v5))), unsafe.Sizeof(uint32(0))*6))))))
			v5 = (*uint8)(unsafe.Add(unsafe.Pointer(v5), 24))
			if v4 == nil {
				break
			}
		}
	}
	v7 = *(**byte)(memmap.PtrOff(0x587000, 210856))
	if *memmap.PtrUint32(0x587000, 210856) != 0 {
		v8 = (*uint8)(memmap.PtrOff(0x587000, 210856))
		for {
			v9 = nox_xxx_modifGetIdByName_413290(v7)
			*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v8))), -int(unsafe.Sizeof(uint32(0))*1)))) = uint32(uintptr(unsafe.Pointer(nox_xxx_modifGetDescById_413330(v9))))
			v7 = (*byte)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v8))), unsafe.Sizeof(uint32(0))*6))))))
			v8 = (*uint8)(unsafe.Add(unsafe.Pointer(v8), 24))
			if v7 == nil {
				break
			}
		}
	}
	v10 = *(**byte)(memmap.PtrOff(0x587000, 211000))
	if *memmap.PtrUint32(0x587000, 211000) != 0 {
		v11 = (*uint8)(memmap.PtrOff(0x587000, 211000))
		for {
			v12 = nox_xxx_modifGetIdByName_413290(v10)
			*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v11))), -int(unsafe.Sizeof(uint32(0))*1)))) = uint32(uintptr(unsafe.Pointer(nox_xxx_modifGetDescById_413330(v12))))
			v10 = (*byte)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v11))), unsafe.Sizeof(uint32(0))*6))))))
			v11 = (*uint8)(unsafe.Add(unsafe.Pointer(v11), 24))
			if v10 == nil {
				break
			}
		}
	}
	result = *(**byte)(memmap.PtrOff(0x587000, 209344))
	if *memmap.PtrUint32(0x587000, 209344) != 0 {
		v14 = (*uint8)(memmap.PtrOff(0x587000, 209344))
		for {
			v15 = nox_xxx_modifGetIdByName_413290(result)
			*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v14))), -int(unsafe.Sizeof(uint32(0))*1)))) = uint32(uintptr(unsafe.Pointer(nox_xxx_modifGetDescById_413330(v15))))
			result = (*byte)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v14))), unsafe.Sizeof(uint32(0))*6))))))
			v14 = (*uint8)(unsafe.Add(unsafe.Pointer(v14), 24))
			if result == nil {
				break
			}
		}
	}
	return result
}
func nox_read_things_alternative_4E2B60() int32 {
	var (
		v1     unsafe.Pointer
		v3     unsafe.Pointer
		result unsafe.Pointer
		v17    int32
		v18    int32
		v19    int32
		i      int32
		v21    unsafe.Pointer
	)
	if nox_xxx_objectTypes_head_1563660 != nil {
		nox_xxx_freeObjectTypes_4E2A20()
	}
	sub_4E3010()
	dword_5d4594_1563664 = 0
	v1 = alloc.Calloc(1, 0x40000)
	v3 = v1
	v21 = v1
	var things *nox_memfile = nox_open_thing_bin()
	if things == nil {
		return 0
	}
	nox_memfile_seek_40AD10(things, 0, 0)
	for nox_memfile_read(unsafe.Pointer(&i), 4, 1, things) != 0 {
		switch i {
		case 0x5350454C:
			nox_thing_skip_spells_415100(int32(uintptr(unsafe.Pointer(things))))
		case 0x4142494C:
			nox_thing_read_ability_415320(int32(uintptr(unsafe.Pointer(things))))
		case 0x41554420:
			nox_thing_read_audio_502320(int32(uintptr(unsafe.Pointer(things))), v3)
		case 0x41564E54:
			nox_thing_read_AVNT_502120(int32(uintptr(unsafe.Pointer(things))), v3)
		case 0x57414C4C:
			if nox_thing_read_WALL_414F60((*uint32)(unsafe.Pointer(things)), v3) == 0 {
				nox_memfile_free(things)
				return 0
			}
		case 0x464C4F52:
			if nox_thing_read_FLOR_414DB0(things) == 0 {
				nox_memfile_free(things)
				return 0
			}
		case 0x45444745:
			if nox_thing_read_EDGE_414E70(int32(uintptr(unsafe.Pointer(things))), v3) == 0 {
				nox_memfile_free(things)
				return 0
			}
		case 0x494D4147:
			nox_thing_read_image_415240(int32(uintptr(unsafe.Pointer(things))))
		case 0x54484E47:
			var typ *nox_objectType_t = new(nox_objectType_t)
			if typ == nil {
				nox_memfile_free(things)
				return 0
			}
			var n int32 = int32(things.ReadU8())
			nox_memfile_read(v3, 1, n, things)
			*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(v3)), n))) = 0
			typ.id = nox_clone_str((*byte)(v3))
			typ.ind = *memmap.PtrUint16(0x587000, 201384)
			*memmap.PtrUint32(0x587000, 201384)++
			typ.field_2 = 0
			typ.menu_icon = -1
			typ.material = 0x4000
			typ.mass = 1.0
			typ.zsize1 = 0
			typ.zsize2 = 30.0
			typ.func_damage = unsafe.Pointer(cgoFuncAddr(nox_xxx_damageDefaultProc_4E0B30))
			typ.func_damage_sound = unsafe.Pointer(cgoFuncAddr(nox_xxx_soundDefaultDamageSound_532E20))
			typ.func_xfer = unsafe.Pointer(cgoFuncAddr(nox_xxx_XFerDefault_4F49A0))
			typ.weight = math.MaxUint8
			if nox_thing_read_xxx_4E3220(things, (*byte)(v3), typ) == 0 {
				nox_memfile_free(things)
				return 0
			}
			if typ.func_collide == nil {
				typ.obj_flags |= 64
			}
			*((*uint16)(unsafe.Add(unsafe.Pointer(&typ.ind), unsafe.Sizeof(uint16(0))*10))) = typ.ind
			typ.obj_flags |= 0x1000000
			if typ.obj_class&0x400000 != 0 {
				typ.mass = 1e+10
			}
			if typ.obj_class&1 != 0 {
				typ.field_13 = 1.0
				typ.speed *= 2
				typ.speed_2 *= 2
			} else {
				typ.field_13 = 0.5
			}
			if libc.StrCmp(typ.id, CString("Boulder")) == 0 || libc.StrCmp(typ.id, CString("RollingBoulder")) == 0 || libc.StrCmp(typ.id, CString("BoulderIndestructible")) == 0 {
				typ.field_13 = 0.01
				typ.mass = 100.0
			}
			if libc.StrCmp(typ.id, CString("Rock7")) == 0 {
				typ.mass = 0.25
			}
			if typ.obj_class&2 != 0 {
				var v13 *uint32 = (*uint32)(typ.data_update)
				*(*uint32)(unsafe.Add(unsafe.Pointer(v13), unsafe.Sizeof(uint32(0))*309)) = math.MaxUint32
				*(*uint32)(unsafe.Add(unsafe.Pointer(v13), unsafe.Sizeof(uint32(0))*307)) = math.MaxUint32
				*(*uint32)(unsafe.Add(unsafe.Pointer(v13), unsafe.Sizeof(uint32(0))*317)) = math.MaxUint32
			} else if typ.obj_class&512 != 0 {
				var v14 *uint32 = (*uint32)(typ.data_update)
				*(*uint32)(unsafe.Add(unsafe.Pointer(v14), unsafe.Sizeof(uint32(0))*6)) = math.MaxUint32
				*(*uint32)(unsafe.Add(unsafe.Pointer(v14), unsafe.Sizeof(uint32(0))*8)) = math.MaxUint32
				*(*uint32)(unsafe.Add(unsafe.Pointer(v14), unsafe.Sizeof(uint32(0))*4)) = math.MaxUint32
			} else if typ.obj_class&2048 != 0 {
				*(*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(typ.collide_data)) + 4))) = math.MaxUint32
			}
			if typ.obj_class&0x3000006 != 0 {
				if typ.data_34 == nil {
					typ.data_34 = alloc.Calloc(1, 20)
					if typ.data_34 == nil {
						return 0
					}
				}
			}
			typ.field_4 = 1
			typ.mass *= 10.0
			dword_5d4594_1563664 ^= uint32(nox_xxx_unitDefProtectMB_4E31A0(int32(uintptr(unsafe.Pointer(typ)))))
			typ.next = nox_xxx_objectTypes_head_1563660
			nox_xxx_objectTypes_head_1563660 = typ
			nox_xxx_unitDefByAlphabetAdd_4E3080(typ.id)
			v3 = v21
		}
	}
	*memmap.PtrUint32(0x85B3FC, 960) = 1
	if nox_loaded_thing_bin != nil {
		if noxflags.HasGame(noxflags.GameClient) && *memmap.PtrUint32(0x85B3FC, 4) != 0 {
			nox_free_thing_bin()
		}
	} else {
		nox_memfile_free(things)
	}
	if nox_xxx_objectTypes_allFit_4E3110() == 0 {
		return 0
	}
	nox_xxx_protectUnitDefUpdateMB_4E3C20()
	result = alloc.Calloc(int(*memmap.PtrUint32(0x587000, 201384)), 4)
	*memmap.PtrUint32(6112660, 1563456) = uint32(uintptr(result))
	if result == nil {
		return 0
	}
	nox_xxx_unitDefByAlphabetInit_4E3040()
	v17 = int32(*memmap.PtrUint32(0x587000, 201384))
	v18 = int32(uintptr(unsafe.Pointer(nox_xxx_objectTypes_head_1563660)))
	v19 = 1
	for i = 1; i < *memmap.PtrInt32(0x587000, 201384); i++ {
		*(*uint32)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, 1563456) + uint32((v17-v19)*4)))) = uint32(v18)
		nox_xxx_objectTypeAddToNameInd_4E30D0(v18)
		v17 = int32(*memmap.PtrUint32(0x587000, 201384))
		v18 = int32(*(*uint32)(unsafe.Pointer(uintptr(v18 + 220))))
		v19 = i + 1
	}
	sub_4E29D0()
	v3 = nil
	nox_xxx_equipWeapon_4131A0()
	nox_xxx_equipArmor_415AB0()
	nox_xxx_equipWeapon_4157C0()
	sub_4F0640()
	if sub_42BF10() == 0 {
		return 0
	}
	return 1
}
func nox_xxx_draw_44C780(a1 int32) unsafe.Pointer {
	var (
		i      int32
		v2     int32
		result unsafe.Pointer
	)
	for i = 0; i < 32; i += 4 {
		v2 = i
		if i >= 16 {
			v2 = i + 4
		}
		result = *(*unsafe.Pointer)(unsafe.Pointer(uintptr(v2 + a1)))
		if result != nil {
			result = nil
		}
	}
	return result
}
func sub_44C7B0(a1 int32) unsafe.Pointer {
	var (
		v1     *unsafe.Pointer
		v2     int32
		v3     *unsafe.Pointer
		v4     int32
		v5     *unsafe.Pointer
		v6     int32
		result unsafe.Pointer
	)
	v1 = (*unsafe.Pointer)(unsafe.Pointer(uintptr(a1 + 52)))
	v2 = 55
	for {
		if *v1 != nil {
			nox_xxx_draw_44C780(int32(uintptr(*v1)) + 4)
			*v1 = nil
		}
		v3 = (*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(unsafe.Pointer(nil))*1))
		v4 = 26
		for {
			if *v3 != nil {
				nox_xxx_draw_44C780(int32(uintptr(*v3)) + 4)
				*v3 = nil
			}
			v3 = (*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(unsafe.Pointer(nil))*1))
			v4--
			if v4 == 0 {
				break
			}
		}
		v5 = (*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(unsafe.Pointer(nil))*27))
		v6 = 27
		for {
			result = *v5
			if *v5 != nil {
				nox_xxx_draw_44C780(int32(uintptr(result)) + 4)
				*v5 = nil
			}
			v5 = (*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(unsafe.Pointer(nil))*1))
			v6--
			if v6 == 0 {
				break
			}
		}
		v1 = (*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(unsafe.Pointer(nil))*66))
		v2--
		if v2 == 0 {
			break
		}
	}
	return result
}
func nox_xxx_draw_44C650_free(lpMem unsafe.Pointer, draw unsafe.Pointer) {
	var kind int32 = 0
	if *(*uint32)(unsafe.Pointer(&nox_parse_thing_draw_funcs[0])) != 0 {
		var item *nox_parse_thing_draw_funcs_t = nil
		for i := int32(0); i < nox_parse_thing_draw_funcs_cnt; i++ {
			var cur *nox_parse_thing_draw_funcs_t = &nox_parse_thing_draw_funcs[i]
			if cur.name == nil {
				break
			}
			if cur.draw == draw {
				item = cur
				break
			}
		}
		if item != nil {
			kind = int32(item.kind)
		}
	}
	var v7 *unsafe.Pointer = nil
	var v8 int32 = 0
	var v9 *byte = nil
	var v10 int32 = 0
	var v11 *byte = nil
	var v12 int32 = 0
	switch kind {
	case 2:
		fallthrough
	case 3:
		if *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(lpMem)), unsafe.Sizeof(uint32(0))*1))) != 0 {
			*((*unsafe.Pointer)(unsafe.Add(unsafe.Pointer((*unsafe.Pointer)(lpMem)), unsafe.Sizeof(unsafe.Pointer(nil))*1))) = nil
		}
		lpMem = nil
	case 4:
		v7 = (*unsafe.Pointer)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(lpMem)), 4))))
		v8 = 5
		for {
			if *v7 != nil {
				*v7 = nil
			}
			v7 = (*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(v7), unsafe.Sizeof(unsafe.Pointer(nil))*1))
			v8--
			if v8 == 0 {
				break
			}
		}
		lpMem = nil
	case 5:
		nox_xxx_draw_44C780(int32(uintptr(lpMem)) + 4)
		lpMem = nil
	case 6:
		sub_44C7B0(int32(uintptr(lpMem)))
		lpMem = nil
	case 7:
		v9 = (*byte)(unsafe.Add(unsafe.Pointer((*byte)(lpMem)), 8))
		v10 = 16
		for {
			nox_xxx_draw_44C780(int32(uintptr(unsafe.Pointer(v9))))
			v9 = (*byte)(unsafe.Add(unsafe.Pointer(v9), 48))
			v10--
			if v10 == 0 {
				break
			}
		}
		lpMem = nil
	case 8:
		v11 = (*byte)(unsafe.Add(unsafe.Pointer((*byte)(lpMem)), 8))
		v12 = 3
		for {
			nox_xxx_draw_44C780(int32(uintptr(unsafe.Pointer(v11))))
			v11 = (*byte)(unsafe.Add(unsafe.Pointer(v11), 48))
			v12--
			if v12 == 0 {
				break
			}
		}
		lpMem = nil
	default:
		lpMem = nil
	}
}
func nox_things_free_44C580() {
	if nox_things_head != nil {
		var cur *nox_thing = nox_things_head
		for cur != nil {
			if cur.name != nil {
				alloc.Free(unsafe.Pointer(cur.name))
			}
			if cur.field_5c != nil {
				nox_xxx_draw_44C650_free(cur.field_5c, unsafe.Pointer(cgoFuncAddr(cur.draw_func)))
			}
			var next *nox_thing = cur.next
			alloc.Free(unsafe.Pointer(cur))
			cur = next
		}
	}
	nox_things_head = nil
	if nox_things_array != nil {
		alloc.Free(unsafe.Pointer(nox_things_array))
		nox_things_array = nil
	}
	nox_things_count = 0
	sub_44C620_free()
	if !noxflags.HasGame(noxflags.GameHost) {
		nox_xxx_free_42BF80()
	}
}
