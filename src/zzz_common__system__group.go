package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	"unsafe"
)

func nox_server_scriptGetGroupId_57C2D0(a1 **int32) int32 {
	var (
		v1 **int32
		v2 int8
		v4 *libc.WChar
	)
	v1 = a1
	v2 = int8(*(*uint8)(unsafe.Pointer(a1)))
	if int32(*(*uint8)(unsafe.Pointer(a1))) != 0 {
		for int32(v2) != 1 && int32(v2) != 2 {
			if int32(v2) != 3 || v1 == nil {
				v4 = strMan.GetStringInFile("Undefined", "C:\\NoxPost\\src\\Common\\System\\Group.c")
				nox_gui_console_PrintOrError_450C30(uint8(int8(NOX_CONSOLE_RED)), (*libc.WChar)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v4)))))))
				return 0
			}
			v1 = (**int32)(unsafe.Pointer(uintptr(nox_server_scriptGetGroup_57C0A0(**(**int32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof((*int32)(nil))*21))))))
			v2 = int8(*(*uint8)(unsafe.Pointer(v1)))
			if int32(*(*uint8)(unsafe.Pointer(v1))) == 0 {
				return int32(*(*uint8)(unsafe.Pointer(v1)))
			}
		}
	}
	return int32(*(*uint8)(unsafe.Pointer(v1)))
}
