package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	"github.com/noxworld-dev/opennox/v1/common/alloc"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"unsafe"
)

func nox_xxx_parseWeapDescription_411B90(a1 *byte, a2 *byte, obj *obj_412ae0_t) int32 {
	var stringName *byte = libc.StrTok(a2, CString(" =\n\r\t"))
	if stringName == nil {
		return 0
	}
	var str1 *libc.WChar = strMan.GetStringInFile(stringName, "C:\\NoxPost\\src\\common\\Object\\Modifier.c")
	var str1len uint32 = nox_wcslen(str1)
	obj.field_2 = (*libc.WChar)(alloc.Calloc(int(str1len+1), 2))
	if obj.field_2 == nil {
		return 0
	}
	var str2 *libc.WChar = strMan.GetStringInFile(stringName, "C:\\NoxPost\\src\\common\\Object\\Modifier.c")
	nox_wcsncpy(obj.field_2, str2, str1len)
	*(*libc.WChar)(unsafe.Add(unsafe.Pointer(obj.field_2), unsafe.Sizeof(libc.WChar(0))*uintptr(str1len))) = 0
	return 1
}
func nox_xxx_parseEnchDesc_412100_parse_desc(a1 *byte, a2 *byte, obj *obj_412ae0_t) int32 {
	var stringName *byte = libc.StrTok(a2, CString(" =\n\r\t"))
	if stringName == nil {
		return 0
	}
	var str1 *libc.WChar = strMan.GetStringInFile(stringName, "C:\\NoxPost\\src\\common\\Object\\Modifier.c")
	var str1len uint32 = nox_wcslen(str1)
	obj.field_2 = (*libc.WChar)(alloc.Calloc(int(str1len+1), 2))
	if obj.field_2 == nil {
		return 0
	}
	var str2 *libc.WChar = strMan.GetStringInFile(stringName, "C:\\NoxPost\\src\\common\\Object\\Modifier.c")
	nox_wcsncpy(obj.field_2, str2, str1len)
	*(*libc.WChar)(unsafe.Add(unsafe.Pointer(obj.field_2), unsafe.Sizeof(libc.WChar(0))*uintptr(str1len))) = 0
	return 1
}
func nox_xxx_parseEnchDescSecondary_4121B0_parse_second_desc(a1 *byte, a2 *byte, obj *obj_412ae0_t) int32 {
	var stringName *byte = libc.StrTok(a2, CString(" =\n\r\t"))
	if stringName == nil {
		return 0
	}
	var str1 *libc.WChar = strMan.GetStringInFile(stringName, "C:\\NoxPost\\src\\common\\Object\\Modifier.c")
	var str1len uint32 = nox_wcslen(str1)
	obj.field_3 = (*libc.WChar)(alloc.Calloc(int(str1len+1), 2))
	if obj.field_3 == nil {
		return 0
	}
	var str2 *libc.WChar = strMan.GetStringInFile(stringName, "C:\\NoxPost\\src\\common\\Object\\Modifier.c")
	nox_wcsncpy(obj.field_3, str2, str1len)
	*(*libc.WChar)(unsafe.Add(unsafe.Pointer(obj.field_3), unsafe.Sizeof(libc.WChar(0))*uintptr(str1len))) = 0
	return 1
}
func nox_xxx_parseIdentifyDesc_412260_parse_ident_desc(a1 *byte, a2 *byte, obj *obj_412ae0_t) int32 {
	var stringName *byte = libc.StrTok(a2, CString(" =\n\r\t"))
	if stringName == nil {
		return 0
	}
	var str1 *libc.WChar = strMan.GetStringInFile(stringName, "C:\\NoxPost\\src\\common\\Object\\Modifier.c")
	var str1len uint32 = nox_wcslen(str1)
	obj.field_4 = (*libc.WChar)(alloc.Calloc(int(str1len+1), 2))
	if obj.field_4 == nil {
		return 0
	}
	var str2 *libc.WChar = strMan.GetStringInFile(stringName, "C:\\NoxPost\\src\\common\\Object\\Modifier.c")
	nox_wcsncpy(obj.field_4, str2, str1len)
	*(*libc.WChar)(unsafe.Add(unsafe.Pointer(obj.field_4), unsafe.Sizeof(libc.WChar(0))*uintptr(str1len))) = 0
	return 1
}
func sub_413480(a1 int8) *libc.WChar {
	var (
		v1 int32
		v2 *uint8
	)
	v1 = 0
	v2 = (*uint8)(memmap.PtrOff(0x587000, 27332))
	for int32(a1) != int32(*v2) {
		v2 = (*uint8)(unsafe.Add(unsafe.Pointer(v2), 20))
		v1++
		if int32(uintptr(unsafe.Pointer(v2))) >= int32(uintptr(memmap.PtrOff(0x587000, 27452))) {
			return nil
		}
	}
	return strMan.GetStringInFile(*(**byte)(memmap.PtrOff(0x587000, uint32(v1*20+0x6AD0))), "C:\\NoxPost\\src\\common\\Object\\Modifier.c")
}
