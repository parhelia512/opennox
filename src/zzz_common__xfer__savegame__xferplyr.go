package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	"github.com/gotranspile/cxgo/runtime/stdio"
	noxflags "github.com/noxworld-dev/opennox/v1/common/flags"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"unsafe"
)

func nox_xxx_savePlayerMB_41C8F0(a1 *byte, a2 uint32) int32 {
	var (
		v2       *libc.WChar
		v3       *libc.WChar
		result   int32
		v5       *byte
		v6       *File
		v7       *byte
		v8       int32
		v9       *libc.WChar
		v10      *libc.WChar
		PathName [1024]byte
	)
	if !noxflags.HasGame(noxflags.GameHost) && a2 < 216 {
		v2 = strMan.GetStringInFile("Wol.c:WolApierror", "C:\\NoxPost\\src\\common\\Xfer\\SaveGame\\XferPlyr.c")
		nox_gui_console_Printf_450C00(uint8(int8(NOX_CONSOLE_BLUE)), v2)
		v3 = strMan.GetStringInFile("Wol.c:Wolapierror", "C:\\NoxPost\\src\\common\\Xfer\\SaveGame\\XferPlyr.c")
		nox_xxx_printCentered_445490(v3)
		nox_xxx_networkLog_printf_413D30(CString("SavePlayerOnClient: Error - character file too small '%s' (%d vs %d)\n"), memmap.PtrOff(0x85B3FC, 10984), a2, *memmap.PtrInt32(0x587000, 55984))
		return 0
	}
	v5 = nox_fs_root()
	nox_sprintf(&PathName[0], CString("%s\\Save\\"), v5)
	nox_fs_mkdir(&PathName[0])
	v6 = nox_binfile_open_408CC0((*byte)(memmap.PtrOff(0x85B3FC, 10984)), 1)
	nox_file_2 = v6
	if v6 == nil {
		nox_xxx_networkLog_printf_413D30(CString("SavePlayerOnClient: Unable to open file '%s'\n"), memmap.PtrOff(0x85B3FC, 10984))
		return 0
	}
	if nox_binfile_cryptSet_408D40((*File)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v6)))))), 27) == 0 {
		nox_xxx_networkLog_printf_413D30(CString("SavePlayerOnClient: Unable to key file '%s'\n"), memmap.PtrOff(0x85B3FC, 10984))
		return 0
	}
	v7 = a1
	if int32(uint16(a2)) != 0 {
		v8 = int32(uint16(a2))
		for {
			nox_binfile_fwrite_409200(func() *byte {
				p := &v7
				x := *p
				*p = (*byte)(unsafe.Add(unsafe.Pointer(*p), 1))
				return x
			}(), 1, 1, nox_file_2)
			v8--
			if v8 == 0 {
				break
			}
		}
	}
	nox_binfile_close_408D90(nox_file_2)
	if noxflags.HasGame(noxflags.GameModeQuest) {
		*memmap.PtrUint8(0x85B3FC, 12257) = uint8(int8(sub_465DF0()))
	} else {
		*memmap.PtrUint8(0x85B3FC, 12257) = 0
	}
	if nox_xxx_mapSavePlayerDataMB_41A230((*byte)(memmap.PtrOff(0x85B3FC, 10984))) != 0 {
		v9 = strMan.GetStringInFile("CharacterSaved", "C:\\NoxPost\\src\\common\\Xfer\\SaveGame\\XferPlyr.c")
		nox_gui_console_Printf_450C00(uint8(int8(NOX_CONSOLE_BLUE)), v9)
		v10 = strMan.GetStringInFile("CharacterSaved", "C:\\NoxPost\\src\\common\\Xfer\\SaveGame\\XferPlyr.c")
		nox_xxx_printCentered_445490(v10)
		result = 1
	} else {
		nox_xxx_networkLog_printf_413D30(CString("SavePlayerOnClient: Unable to save client data to file\n"))
		result = 0
	}
	return result
}
