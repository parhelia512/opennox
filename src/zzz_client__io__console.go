package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	"github.com/gotranspile/cxgo/runtime/stdio"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"unsafe"
)

var nox_file_log *File = nil

func sub_451610() int32 {
	var (
		result int32
		v1     *uint8
	)
	result = int32(*memmap.PtrUint32(6112660, 835876))
	v1 = (*uint8)(memmap.PtrOff(6112660, 835876))
	for {
		*(*uint32)(unsafe.Pointer(v1)) = *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v1))), -int(unsafe.Sizeof(uint32(0))*1))))
		v1 = (*uint8)(unsafe.Add(unsafe.Pointer(v1), -4))
		if int32(uintptr(unsafe.Pointer(v1))) <= int32(uintptr(memmap.PtrOff(6112660, 835800))) {
			break
		}
	}
	*memmap.PtrUint32(6112660, 835800) = uint32(result)
	return result
}
func sub_451630() *uint8 {
	var (
		v1     *uint8
		result *uint8
		v3     *libc.WChar
	)
	_ = v3
	var v4 *libc.WChar
	_ = v4
	nox_file_log = nox_fs_create_text(CString("log"))
	if nox_file_log == nil {
		v4 = strMan.GetStringInFile("FatalError", "C:\\NoxPost\\src\\Client\\Io\\Console.c")
		v3 = strMan.GetStringInFile("CantOpenLog", "C:\\NoxPost\\src\\Client\\Io\\Console.c")
		nox_exit(0)
	}
	v1 = (*uint8)(memmap.PtrOff(6112660, 835880))
	result = (*uint8)(memmap.PtrOff(6112660, 835800))
	for {
		*(*uint32)(unsafe.Pointer(result)) = uint32(uintptr(unsafe.Pointer(v1)))
		*(*uint16)(unsafe.Pointer(v1)) = 0
		result = (*uint8)(unsafe.Add(unsafe.Pointer(result), 4))
		v1 = (*uint8)(unsafe.Add(unsafe.Pointer(v1), 200))
		if int32(uintptr(unsafe.Pointer(result))) >= int32(uintptr(memmap.PtrOff(6112660, 835880))) {
			break
		}
	}
	return result
}
func sub_4516C0(a1 *libc.WChar, _rest ...interface{}) {
	var (
		v1 *uint16
		v3 *uint16
	)
	_ = v3
	var va libc.ArgList
	va.Start(a1, _rest)
	if nox_file_log == nil {
		sub_451630()
	}
	sub_451610()
	v1 = (*uint16)(unsafe.Pointer(strMan.GetStringInFile("FatalErrorHeader", (*byte)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(CString("C:\\NoxPost\\src\\Client\\Io\\Console.c"))))))))))
	nox_swprintf((*libc.WChar)(unsafe.Pointer(memmap.PtrUint16(6112660, 833752))), (*libc.WChar)(unsafe.Pointer(v1)))
	nox_vswprintf((*libc.WChar)(unsafe.Pointer(memmap.PtrUint16(6112660, 833778))), a1, va)
	nox_fs_fprintf(nox_file_log, CString("%S"), memmap.PtrOff(6112660, 833752))
	nox_fs_flush(nox_file_log)
	v3 = (*uint16)(unsafe.Pointer(strMan.GetStringInFile("FatalError", (*byte)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(CString("C:\\NoxPost\\src\\Client\\Io\\Console.c"))))))))))
	nox_fs_fprintf(nox_file_log, CString("exiting..\n"))
	nox_fs_close(nox_file_log)
	if dword_5d4594_823776 != 0 {
		nox_xxx_freeFloorBuffer_430EF0()
	}
	nox_input_pollEvents_4453A0()
	nox_input_pollEvents_4453A0()
	nox_exit(0)
}
