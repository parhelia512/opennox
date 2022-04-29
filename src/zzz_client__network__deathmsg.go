package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"unsafe"
)

func sub_4952E0(a1 *uint16) int32 {
	var (
		v1  *byte
		v2  *libc.WChar
		v3  *byte
		v4  *libc.WChar
		v5  *byte
		v6  *libc.WChar
		v8  int32
		v9  *libc.WChar
		v10 int32
		v11 [32]libc.WChar
		v12 [32]libc.WChar
		v13 [64]libc.WChar
	)
	if int32(*(*uint16)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint16(0))*1))) != 0 && (func() *byte {
		v1 = (*byte)(unsafe.Pointer(noxServer.getPlayerByID(int32(*(*uint16)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint16(0))*1))))))
		return v1
	}()) != nil {
		v8 = int32(uintptr(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v1), 4704)))))
		v2 = strMan.GetStringInFile("die.c:LocalizeAttacker", "C:\\NoxPost\\src\\client\\Network\\deathmsg.c")
		nox_swprintf(&v13[0], v2, v8)
		if int32(*(*uint16)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint16(0))*2))) != 0 {
			v3 = (*byte)(unsafe.Pointer(noxServer.getPlayerByID(int32(*(*uint16)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint16(0))*2))))))
			if v3 != nil {
				nox_swprintf(&v12[0], libc.CWString(" + %s"), (*byte)(unsafe.Add(unsafe.Pointer(v3), 4704)))
				nox_wcscat(&v13[0], &v12[0])
			}
		}
	} else {
		v9 = strMan.GetStringInFile("die.c:AttackerNasty", "C:\\NoxPost\\src\\client\\Network\\deathmsg.c")
		v4 = strMan.GetStringInFile("die.c:LocalizeAttacker", "C:\\NoxPost\\src\\client\\Network\\deathmsg.c")
		nox_swprintf(&v13[0], v4, v9)
	}
	if int32(*(*uint16)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint16(0))*3))) != 0 {
		v5 = (*byte)(unsafe.Pointer(noxServer.getPlayerByID(int32(*(*uint16)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint16(0))*3))))))
		if v5 != nil {
			v10 = int32(uintptr(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v5), 4704)))))
			v6 = strMan.GetStringInFile("die.c:LocalizeVictim", "C:\\NoxPost\\src\\client\\Network\\deathmsg.c")
			nox_swprintf(&v11[0], v6, v10)
		}
	}
	return nox_gui_console_Printf_450C00(uint8(int8(NOX_CONSOLE_WHITE)), (*libc.WChar)(memmap.PtrOff(0x587000, 161668)), &v11[0], &v13[0])
}
