package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	"github.com/gotranspile/cxgo/runtime/stdio"
	"github.com/noxworld-dev/opennox/v1/common/alloc"
	noxflags "github.com/noxworld-dev/opennox/v1/common/flags"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"math"
	"unicode"
	"unsafe"
)

func nox_xxx_unknown_libname_86_57F634() {
}
func nullsub_45() {
}
func nullsub_46() {
}
func nullsub_47() {
}
func nullsub_48() {
}
func nullsub_62() {
}
func nullsub_65() {
}
func nullsub_66() {
}
func nullsub_67() {
}
func nullsub_68() {
}
func nullsub_69() {
}
func nullsub_70() {
}

var nox_enable_audio int32 = 1
var nox_enable_threads int32 = 1
var nox_video_dxFullScreen uint32 = 0
var nox_video_dxUnlockSurface uint32 = 0
var nox_video_cursorDrawThreadHandle *HANDLE
var nox_server_gameSettingsUpdated int32
var g_v20 int32
var g_v21 int32
var g_argc2 int32
var g_argv2 **byte
var dword_5D4594_251544 *uint32 = nil
var byte_5D4594_251584 [3]*obj_412ae0_t = [3]*obj_412ae0_t{}
var byte_5D4594_251596 int32 = 0

func nox_get_audio_enabled() int32 {
	return nox_enable_audio
}
func sub_401040() int32 {
	var result int32
	result = bool2int(noxflags.HasGame(noxflags.GameHost))
	nox_frame_xxx_2598000 = uint32(result)
	return result
}
func sub_401060() int32 {
	return int32(*memmap.PtrUint32(6112660, 264))
}
func nox_xxx_parseRead_4093E0(a1 *File, a2 *byte, a3 int32) {
	var (
		v3 int32
		v4 int32
		v5 *byte
		v6 int32
	)
	v3 = 0
	v4 = 1
	if a3 <= 0 {
	LABEL_12:
		*(*byte)(unsafe.Add(unsafe.Pointer(a2), a3-1)) = 0
		return
	}
	v5 = a2
	for {
		nox_binfile_fread_408E40(v5, 1, 1, a1)
		v6 = bool2int(unicode.IsSpace(rune(*v5)))
		if v6 == 0 {
			break
		}
		if v4 == 0 {
			*v5 = 32
		LABEL_10:
			v3++
			v5 = (*byte)(unsafe.Add(unsafe.Pointer(v5), 1))
		}
		if v3 >= a3 {
			goto LABEL_12
		}
	}
	v4 = 0
	if *v5 != 59 {
		goto LABEL_10
	}
	*(*byte)(unsafe.Add(unsafe.Pointer(a2), v3)) = 0
}
func nox_xxx_parseString_409470(a1 *File, a2 *uint8) int32 {
	var (
		v2       *uint8
		v3       int32
		v4       int32
		v5       int32
		v6       int32
		CharType [2]uint16
	)
	v2 = a2
	v3 = 0
	*(*uint32)(unsafe.Pointer(&CharType[0])) = 0
	v4 = 1
	for {
		for {
			v5 = v3
			nox_binfile_fread_408E40((*byte)(unsafe.Pointer(&CharType[0])), 1, 1, a1)
			if nox_binfile_lastErr_409370(a1) == -1 {
				return 0
			}
			if *memmap.PtrUint32(0x587000, 1668) <= 1 {
				v3 = int32(*(*uint32)(unsafe.Pointer(&CharType[0])))
				v6 = bool2int(unicode.IsSpace(rune(CharType[0])))
			} else {
				v6 = bool2int(unicode.IsSpace(rune(CharType[0])))
				v3 = int32(*(*uint32)(unsafe.Pointer(&CharType[0])))
			}
			if v6 != 0 {
				break
			}
			v4 = 0
			if v3 != int32('/') || v5 != int32('/') {
				*func() *uint8 {
					p := &v2
					x := *p
					*p = (*uint8)(unsafe.Add(unsafe.Pointer(*p), 1))
					return x
				}() = uint8(int8(v3))
			} else {
				nox_binfile_skipLine_409520(a1)
				v2 = a2
				v3 = int32(*(*uint32)(unsafe.Pointer(&CharType[0])))
				v4 = 1
			}
		}
		if v4 == 0 {
			break
		}
	}
	*v2 = 0
	return 1
}
func sub_409A70(a1 int16) int32 {
	var (
		result int32
		v2     *uint8
	)
	result = 0
	v2 = (*uint8)(memmap.PtrOff(0x587000, 4704))
	for *(*uint32)(unsafe.Pointer(v2)) != uint32(int32(a1)&6128) {
		v2 = (*uint8)(unsafe.Add(unsafe.Pointer(v2), 4))
		result++
		if int32(uintptr(unsafe.Pointer(v2))) >= int32(uintptr(memmap.PtrOff(0x587000, 4728))) {
			return 0
		}
	}
	return result
}

var nox_version_code uint32 = 0

func nox_client_getVersionBuild_409AC0() uint16 {
	return uint16(nox_version_code & math.MaxUint16)
}
func nox_client_getVersionMajor_409AA0() uint8 {
	return uint8((nox_version_code >> 24) & math.MaxUint8)
}
func nox_client_getVersionMinor_409AB0() uint8 {
	return uint8((nox_version_code >> 16) & math.MaxUint8)
}
func nox_client_getVersionCode_409AD0() uint32 {
	return nox_version_code
}
func nox_client_setVersion_409AE0(vers uint32) {
	if nox_version_code != vers {
		nox_version_code = vers
		nox_server_gameSettingsUpdated = 1
	}
}
func nox_xxx_mapCrcGetMB_409B00() int32 {
	return int32(*memmap.PtrUint32(6112660, 3604))
}
func nox_xxx_mapSetCrcMB_409B10(a1 int32) int32 {
	var result int32
	result = a1
	if *memmap.PtrUint32(6112660, 3604) != uint32(a1) {
		*memmap.PtrUint32(6112660, 3604) = uint32(a1)
	}
	return result
}
func nox_server_currentMapGetFilename_409B30() *byte {
	return (*byte)(memmap.PtrOff(6112660, 2598188))
}
func nox_xxx_mapGetMapName_409B40() *byte {
	return (*byte)(memmap.PtrOff(0x85B3FC, 36))
}
func sub_409B50(a1 *byte) uint32 {
	var result uint32
	result = uint32(libc.StrLen(a1) + 1)
	libc.MemCpy(memmap.PtrOff(6112660, 3452), unsafe.Pointer(a1), int(result))
	return result
}
func sub_409B80() *byte {
	return (*byte)(memmap.PtrOff(6112660, 3452))
}
func nox_server_get_current_map_path_409B90() *byte {
	var data_path *byte = nox_fs_root()
	if data_path == nil {
		return nil
	}
	var v1 int16 = int16(*memmap.PtrUint16(0x587000, 4736))
	libc.StrCpy((*byte)(memmap.PtrOff(6112660, 2424)), data_path)
	var v2 uint8 = (*memmap.PtrUint8(0x587000, 4738))
	var v3 *uint8 = (*uint8)(memmap.PtrOff(6112660, uint32(libc.StrLen((*byte)(memmap.PtrOff(6112660, 2424)))+2424)))
	*(*uint32)(unsafe.Pointer(v3)) = *memmap.PtrUint32(0x587000, 4732)
	*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v3))), unsafe.Sizeof(uint16(0))*2))) = uint16(v1)
	*(*uint8)(unsafe.Add(unsafe.Pointer(v3), 6)) = v2
	libc.StrCat((*byte)(memmap.PtrOff(6112660, 2424)), nox_xxx_mapGetMapName_409B40())
	*memmap.PtrUint16(6112660, uint32(libc.StrLen((*byte)(memmap.PtrOff(6112660, 2424)))+2424)) = *memmap.PtrUint16(0x587000, 4740)
	libc.StrCat((*byte)(memmap.PtrOff(6112660, 2424)), nox_server_currentMapGetFilename_409B30())
	return (*byte)(memmap.PtrOff(6112660, 2424))
}
func sub_409C70() *byte {
	var (
		v0     *byte
		v1     int16
		v2     uint8
		v3     *uint8
		v4     *uint8
		v5     uint8
		result *byte
	)
	v0 = nox_fs_root()
	if v0 == nil {
		return nil
	}
	v1 = int16(*memmap.PtrUint16(0x587000, 4748))
	libc.StrCpy((*byte)(memmap.PtrOff(6112660, 2424)), v0)
	v2 = *memmap.PtrUint8(0x587000, 4750)
	v3 = (*uint8)(memmap.PtrOff(6112660, uint32(libc.StrLen((*byte)(memmap.PtrOff(6112660, 2424)))+2424)))
	*(*uint32)(unsafe.Pointer(v3)) = *memmap.PtrUint32(0x587000, 4744)
	*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v3))), unsafe.Sizeof(uint16(0))*2))) = uint16(v1)
	*(*uint8)(unsafe.Add(unsafe.Pointer(v3), 6)) = v2
	libc.StrCat((*byte)(memmap.PtrOff(6112660, 2424)), nox_xxx_mapGetMapName_409B40())
	*memmap.PtrUint16(6112660, uint32(libc.StrLen((*byte)(memmap.PtrOff(6112660, 2424)))+2424)) = *memmap.PtrUint16(0x587000, 4752)
	libc.StrCat((*byte)(memmap.PtrOff(6112660, 2424)), nox_xxx_mapGetMapName_409B40())
	v4 = (*uint8)(memmap.PtrOff(6112660, uint32(libc.StrLen((*byte)(memmap.PtrOff(6112660, 2424)))+2425)))
	v5 = *memmap.PtrUint8(0x587000, 4760)
	*(*uint32)(unsafe.Pointer(func() *uint8 {
		p := &v4
		*p = (*uint8)(unsafe.Add(unsafe.Pointer(*p), -1))
		return *p
	}())) = *memmap.PtrUint32(0x587000, 4756)
	result = (*byte)(memmap.PtrOff(6112660, 2424))
	*(*uint8)(unsafe.Add(unsafe.Pointer(v4), 4)) = v5
	return result
}
func nox_xxx_gameSetMapPath_409D70(a1 *byte) *byte {
	var (
		result *byte
		v2     *byte
		v3     int32
	)
	result = (*byte)(unsafe.Pointer(uintptr(libc.StrCaseCmp((*byte)(memmap.PtrOff(6112660, 2598188)), a1))))
	if result != nil {
		libc.StrNCpy((*byte)(memmap.PtrOff(6112660, 2598188)), a1, 80)
		*memmap.PtrUint8(6112660, 2598267) = 0
		v2 = libc.StrRChr(a1, 92)
		if v2 != nil {
			v3 = int32(libc.StrLen((*byte)(unsafe.Add(unsafe.Pointer(v2), 1))) - 4)
			result = libc.StrNCpy((*byte)(memmap.PtrOff(0x85B3FC, 36)), (*byte)(unsafe.Add(unsafe.Pointer(v2), 1)), int(v3))
		} else {
			v3 = int32(libc.StrLen((*byte)(memmap.PtrOff(6112660, 2598188))) - 4)
			if v3 < 0 {
				v3 = 0
			}
			result = libc.StrNCpy((*byte)(memmap.PtrOff(0x85B3FC, 36)), (*byte)(memmap.PtrOff(6112660, 2598188)), int(v3))
		}
		*memmap.PtrUint8(0x85B3FC, uint32(v3+36)) = 0
		nox_server_gameSettingsUpdated = 1
	}
	return result
}
func sub_409E40(a1 int32) int32 {
	var result int32
	result = a1
	if dword_5d4594_3484 != uint32(a1) {
		dword_5d4594_3484 = uint32(a1)
		nox_server_gameSettingsUpdated = 1
	}
	return result
}
func nox_xxx_getServerSubFlags_409E60() int32 {
	return int32(dword_5d4594_3484)
}
func sub_409E70(a1 int32) int32 {
	var result int32
	result = int32(uint32(a1) & dword_5d4594_3484)
	if (uint32(a1) & dword_5d4594_3484) != uint32(a1) {
		dword_5d4594_3484 |= uint32(a1)
		result = bool2int(noxflags.HasGame(noxflags.GameHost))
		if result != 0 {
			if a1&8192 != 0 {
				result = sub_4D7EA0()
			}
		}
		nox_server_gameSettingsUpdated = 1
	}
	return result
}
func sub_409EC0(a1 int32) int32 {
	var result int32
	result = a1
	if dword_5d4594_3484&uint32(a1) != 0 {
		result = ^a1
		nox_server_gameSettingsUpdated = 1
		dword_5d4594_3484 &= uint32(^a1)
	}
	return result
}
func sub_409EF0(a1 int32) int32 {
	var result int32
	dword_5d4594_3484 ^= uint32(a1)
	result = bool2int(noxflags.HasGame(noxflags.GameHost))
	if result != 0 {
		if a1&8192 != 0 {
			result = int32(dword_5d4594_3484)
			if result&8192 != 0 {
				result = sub_4D7EA0()
			}
		}
	}
	nox_server_gameSettingsUpdated = 1
	return result
}
func sub_409F40(a1 int32) int32 {
	var result int32
	if a1 == 8192 && noxflags.HasGame(noxflags.GameModeCTF|noxflags.GameModeElimination) {
		result = 1
	} else {
		result = bool2int((uint32(a1) & dword_5d4594_3484) != 0)
	}
	return result
}
func nox_xxx_servSetPlrLimit_409F80(a1 int32) int32 {
	var result int32
	result = a1
	if *memmap.PtrUint32(6112660, 3464) != uint32(a1) {
		*memmap.PtrUint32(6112660, 3464) = uint32(a1)
		nox_server_gameSettingsUpdated = 1
	}
	return result
}
func nox_xxx_servGetPlrLimit_409FA0() int32 {
	return int32(*memmap.PtrUint32(6112660, 3464))
}
func nox_xxx_servGamedataGet_40A020(a1 int16) int16 {
	return int16(*memmap.PtrUint16(6112660, uint32(sub_409A70(a1)*2+3488)))
}
func sub_40A180(a1 int16) uint8 {
	return *memmap.PtrUint8(6112660, uint32(sub_409A70(a1)+3500))
}
func sub_40A1A0() int32 {
	var v0 int16
	v0 = int16(uint16(nox_common_gameFlags_getVal_40A5B0()))
	return bool2int((int32(sub_40A180(v0)) != 0 || dword_5d4594_3592 != 0) && *memmap.PtrUint32(0x587000, 4660) != 0 && uint64(platformTicks()) > *memmap.PtrUint64(6112660, 3468))
}
func sub_40A1F0(a1 int32) int32 {
	var result int32
	*memmap.PtrUint32(0x587000, 4660) = uint32(a1)
	result = bool2int(noxflags.HasGame(noxflags.GameHost))
	if result != 0 {
		result = nox_xxx_netTimerStatus_4D8F50(159, a1)
	}
	return result
}
func sub_40A220() int32 {
	return int32(*memmap.PtrUint32(0x587000, 4660))
}
func sub_40A230() int32 {
	return int32(uint32(uint64(*memmap.PtrUint32(6112660, 3468)) - uint64(platformTicks())))
}
func sub_40A250() int64 {
	var (
		v0     int16
		v1     int32
		v2     int64
		result int64
	)
	v0 = int16(uint16(nox_common_gameFlags_getVal_40A5B0()))
	v1 = sub_409A70(v0)
	v2 = int64(platformTicks())
	result = int64(int32(*memmap.PtrUint8(6112660, uint32(v1+3500))) * 60000)
	*memmap.PtrUint64(6112660, 3468) = uint64(result + v2)
	return result
}
func nox_xxx_servStartCountdown_40A2A0(a1 int32, a2 *byte) *byte {
	var result *byte
	*memmap.PtrUint64(6112660, 3468) = uint64(uint32(a1*1000) + platformTicks())
	sub_40A1F0(1)
	result = a2
	if a2 != nil {
		result = (*byte)(unsafe.Pointer(uintptr(nox_xxx_netPrintLineToAll_4DA390(a2))))
	}
	dword_5d4594_3592 = 1
	return result
}
func sub_40A300() int32 {
	return int32(dword_5d4594_3592)
}
func sub_40A310(a1 int32) int64 {
	var (
		v1     int64
		result int64
	)
	v1 = int64(platformTicks())
	result = int64(a1)
	*memmap.PtrUint64(6112660, 3468) = uint64(int64(a1) + v1)
	return result
}
func nox_xxx_set3512_40A340(a1 int32) {
	*memmap.PtrUint32(6112660, 3512) = uint32(a1)
}
func nox_xxx_get3512_40A350() int32 {
	return int32(*memmap.PtrUint32(6112660, 3512))
}
func sub_40A3C0(a1 uint32) uint32 {
	var result uint32
	result = a1
	if a1 > 2048 && a1 < 0x8000 {
		*memmap.PtrUint32(0x587000, 4652) = a1
	}
	return result
}
func nox_xxx_setClientNetPort_40A410(a1 int32) int32 {
	var result int32
	result = a1
	*memmap.PtrUint32(6112660, 3528) = uint32(a1)
	return result
}
func nox_client_getClientPort_40A420() int32 {
	return int32(*memmap.PtrUint32(6112660, 3528))
}
func nox_xxx_gameSetServername_40A440(a1 *byte) *byte {
	var (
		result *byte
		v2     [16]byte
	)
	result = a1
	if a1 != nil {
		libc.StrNCpy(&v2[0], a1, 15)
		v2[15] = 0
		result = (*byte)(unsafe.Pointer(uintptr(libc.StrCaseCmp((*byte)(memmap.PtrOff(6112660, 1324)), &v2[0]))))
		if result != nil {
			result = libc.StrNCpy((*byte)(memmap.PtrOff(6112660, 1324)), &v2[0], 15)
			*memmap.PtrUint8(6112660, 1339) = 0
			nox_server_gameSettingsUpdated = 1
		}
	} else {
		*memmap.PtrUint8(6112660, 1324) = 0
		nox_server_gameSettingsUpdated = 1
	}
	return result
}
func nox_xxx_serverOptionsGetServername_40A4C0() *byte {
	return (*byte)(memmap.PtrOff(6112660, 1324))
}
func nox_xxx_ruleSetNoRespawn_40A5E0(a1 int32) int32 {
	var result int32
	result = a1
	*memmap.PtrUint32(6112660, 3584) = uint32(a1)
	return result
}
func nox_server_doPlayersAutoRespawn_40A5F0() int32 {
	if !noxflags.HasGame(noxflags.GameModeQuest) {
		return int32(*memmap.PtrUint32(6112660, 3584))
	}
	return 0
}
func nox_xxx_sysopSetPass_40A610(a1 *libc.WChar) *libc.WChar {
	return nox_wcscpy((*libc.WChar)(memmap.PtrOff(6112660, 3540)), a1)
}
func nox_xxx_sysopGetPass_40A630() *libc.WChar {
	return (*libc.WChar)(memmap.PtrOff(6112660, 3540))
}
func sub_40A640(a1 *libc.WChar) *libc.WChar {
	var result *libc.WChar
	result = nox_wcsncpy((*libc.WChar)(memmap.PtrOff(6112660, 3560)), a1, 8)
	*memmap.PtrUint16(6112660, 3576) = 0
	return result
}
func sub_40A660() *libc.WChar {
	return (*libc.WChar)(memmap.PtrOff(6112660, 3560))
}
func nox_server_gameSettingsUpdated_40A670() {
	nox_server_gameSettingsUpdated = 1
}
func nox_server_gameDoSwitchMap_40A680() int32 {
	return nox_server_gameSettingsUpdated
}
func nox_server_gameUnsetMapLoad_40A690() {
	nox_server_gameSettingsUpdated = 0
}
func sub_40A6A0(a1 int32) int32 {
	var result int32
	result = a1
	*memmap.PtrUint32(6112660, 3588) = uint32(a1)
	return result
}
func sub_40A6B0() int32 {
	return int32(*memmap.PtrUint32(6112660, 3588))
}
func nox_xxx_rateGet_40A6C0() int32 {
	return int32(*memmap.PtrUint32(0x587000, 4728))
}
func nox_xxx_rateUpdate_40A6D0(a1 int32) int32 {
	var result int32
	result = int32(*memmap.PtrUint32(0x587000, 4728))
	if uint32(a1) != *memmap.PtrUint32(0x587000, 4728) {
		result = bool2int(noxflags.HasGame(noxflags.GameFlag18))
		if result == 1 {
			result = nox_xxx_netNotifyRate_4D7F10(159)
		}
	}
	*memmap.PtrUint32(0x587000, 4728) = uint32(a1)
	return result
}
func sub_40A710(a1 int32) int32 {
	var (
		v1 int32
		v2 *uint8
	)
	v1 = 0
	v2 = (*uint8)(memmap.PtrOff(0x587000, 4664))
	for *(*uint32)(unsafe.Pointer(v2)) != uint32(a1) {
		v2 = (*uint8)(unsafe.Add(unsafe.Pointer(v2), 8))
		v1++
		if int32(uintptr(unsafe.Pointer(v2))) >= int32(uintptr(memmap.PtrOff(0x587000, 4704))) {
			return 1
		}
	}
	return int32(*memmap.PtrUint32(0x587000, uint32(v1*8+4668)))
}
func sub_40A740() int32 {
	var v0 *byte
	v0 = sub_4165B0()
	return bool2int(noxflags.HasGame(noxflags.GameModeChat) && *(*byte)(unsafe.Add(unsafe.Pointer(v0), 53)) < 0)
}
func sub_40A770() int32 {
	var (
		v0 int32
		v1 *byte
		v2 int32
		v3 int32
		i  *byte
	)
	v0 = 0
	if !nox_xxx_CheckGameplayFlags_417DA0(4) {
		for i = (*byte)(unsafe.Pointer(nox_common_playerInfoGetFirst_416EA0())); i != nil; i = (*byte)(unsafe.Pointer(nox_common_playerInfoGetNext_416EE0((*nox_playerInfo)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(i))))))))) {
			if (*(*byte)(unsafe.Add(unsafe.Pointer(i), 3680))&1) == 0 && (*(*byte)(unsafe.Add(unsafe.Pointer(i), 2064)) != 31 || !nox_common_getEngineFlag(nox_engine_flag(NOX_ENGINE_FLAG_DISABLE_GRAPHICS_RENDERING))) {
				v0++
			}
		}
		return v0
	}
	v1 = (*byte)(unsafe.Pointer(nox_server_teamFirst_418B10()))
	if v1 == nil {
		return v0
	}
	for {
		v2 = int32(uintptr(unsafe.Pointer(nox_xxx_getFirstPlayerUnit_4DA7C0())))
		if v2 != 0 {
			for {
				v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 748))) + 276))))
				if (int32(*(*uint8)(unsafe.Pointer(uintptr(v3 + 3680))))&1) == 0 && (int32(*(*uint8)(unsafe.Pointer(uintptr(v3 + 2064)))) != 31 || !nox_common_getEngineFlag(nox_engine_flag(NOX_ENGINE_FLAG_DISABLE_GRAPHICS_RENDERING))) {
					break
				}
				v2 = int32(uintptr(unsafe.Pointer(nox_xxx_getNextPlayerUnit_4DA7F0((*nox_object_t)(unsafe.Pointer(uintptr(v2)))))))
				if v2 == 0 {
					goto LABEL_10
				}
			}
			v0++
		}
	LABEL_10:
		v1 = (*byte)(unsafe.Pointer(nox_server_teamNext_418B60((*nox_team_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v1)))))))))
		if v1 == nil {
			break
		}
	}
	return v0
}
func nox_xxx_countNonEliminatedPlayersInTeam_40A830(a1 int32) int32 {
	var (
		v1 int32
		v2 int32
		v3 int32
	)
	v1 = 0
	v2 = int32(uintptr(unsafe.Pointer(nox_xxx_getFirstPlayerUnit_4DA7C0())))
	if v2 == 0 {
		return 0
	}
	for {
		if nox_xxx_teamCompare2_419180(v2+48, *(*uint8)(unsafe.Pointer(uintptr(a1 + 57)))) != 0 {
			v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 748))) + 276))))
			if (int32(*(*uint8)(unsafe.Pointer(uintptr(v3 + 3680))))&1) == 0 && (int32(*(*uint8)(unsafe.Pointer(uintptr(v3 + 2064)))) != 31 || !nox_common_getEngineFlag(nox_engine_flag(NOX_ENGINE_FLAG_DISABLE_GRAPHICS_RENDERING))) {
				v1++
			}
		}
		v2 = int32(uintptr(unsafe.Pointer(nox_xxx_getNextPlayerUnit_4DA7F0((*nox_object_t)(unsafe.Pointer(uintptr(v2)))))))
		if v2 == 0 {
			break
		}
	}
	return v1
}
func nox_xxx_gamePlayIsAnyPlayers_40A8A0() int32 {
	var (
		v0     int32
		v1     *byte
		v2     int32
		v3     int32
		result int32
		i      *byte
		v6     int32
	)
	v0 = 0
	if !nox_xxx_CheckGameplayFlags_417DA0(4) {
		for i = (*byte)(unsafe.Pointer(nox_common_playerInfoGetFirst_416EA0())); i != nil; i = (*byte)(unsafe.Pointer(nox_common_playerInfoGetNext_416EE0((*nox_playerInfo)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(i))))))))) {
			if *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(i))), unsafe.Sizeof(uint32(0))*514))) != 0 {
				v6 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(i))), unsafe.Sizeof(uint32(0))*920))))
				if (v6&1) == 0 || v6&32 != 0 {
					v0++
				}
			}
		}
		goto LABEL_18
	}
	v1 = (*byte)(unsafe.Pointer(nox_server_teamFirst_418B10()))
	if v1 == nil {
	LABEL_18:
		result = bool2int(v0 < 1)
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&result))), 0)) = uint8(int8(bool2int(v0 > 1)))
		return result
	}
	for {
		v2 = int32(uintptr(unsafe.Pointer(nox_xxx_getFirstPlayerUnit_4DA7C0())))
		if v2 != 0 {
			for {
				if nox_xxx_teamCompare2_419180(v2+48, uint8(*(*byte)(unsafe.Add(unsafe.Pointer(v1), 57)))) != 0 {
					v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 748))) + 276))) + 3680))))
					if (v3&1) == 0 || v3&32 != 0 {
						break
					}
				}
				v2 = int32(uintptr(unsafe.Pointer(nox_xxx_getNextPlayerUnit_4DA7F0((*nox_object_t)(unsafe.Pointer(uintptr(v2)))))))
				if v2 == 0 {
					goto LABEL_10
				}
			}
			v0++
		}
	LABEL_10:
		v1 = (*byte)(unsafe.Pointer(nox_server_teamNext_418B60((*nox_team_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v1)))))))))
		if v1 == nil {
			break
		}
	}
	return bool2int(v0 > 1)
}
func sub_40A970() {
	var (
		i  *byte
		v1 int32
		v3 float32
		v4 float32
	)
	*memmap.PtrUint32(6112660, 3520) = nox_frame_xxx_2598000
	*memmap.PtrUint32(6112660, 3536) = 0
	v3 = float32(nox_xxx_gamedataGetFloat_419D40(CString("SuddenDeathPlayerThreshold")))
	*memmap.PtrUint32(6112660, 3476) = uint32(int(v3))
	v4 = float32(nox_xxx_gamedataGetFloat_419D40(CString("SuddenDeathLifeTime")))
	*memmap.PtrUint32(6112660, 1392) = uint32(int(v4))
	for i = (*byte)(unsafe.Pointer(nox_common_playerInfoGetFirst_416EA0())); i != nil; i = (*byte)(unsafe.Pointer(nox_common_playerInfoGetNext_416EE0((*nox_playerInfo)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(i))))))))) {
		v1 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(i))), unsafe.Sizeof(uint32(0))*920))))
		if v1&256 != 0 {
			nox_xxx_playerUnsetStatus_417530((*nox_playerInfo)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(i)))))), 256)
		}
	}
	noxflags.UnsetGame(noxflags.GameSuddenDeath)
}
func sub_40AA00() int32 {
	return bool2int(nox_gameFPS*20 < (nox_frame_xxx_2598000 - *memmap.PtrUint32(6112660, 3520)))
}
func sub_40AA20() int32 {
	return int32(*memmap.PtrUint32(6112660, 3536))
}
func sub_40AA30(a1 int32) int32 {
	var result int32
	result = a1
	*memmap.PtrUint32(6112660, 3536) = uint32(a1)
	return result
}
func sub_40AA40() int32 {
	return int32(*memmap.PtrUint32(6112660, 3476))
}
func sub_40AA60(a1 int32) int32 {
	var result int32
	result = a1
	*memmap.PtrUint32(6112660, 3508) = uint32(a1)
	return result
}
func sub_40AA70(pl *nox_playerInfo) int32 {
	var (
		a1     int32 = int32(uintptr(unsafe.Pointer(pl)))
		v1     *byte
		result int32
		v3     int32
		v4     int32
		v5     int32
		v6     *byte
	)
	v1 = (*byte)(sub_416640())
	if a1 == 0 {
		goto LABEL_31
	}
	if noxflags.HasGame(noxflags.GameModeQuest) {
		return bool2int(sub_40A770() < 6)
	}
	v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 3680))))
	if v3&256 != 0 && *memmap.PtrUint32(6112660, 3508) == 0 {
		return 0
	}
	if sub_40A740() == 0 && !noxflags.HasGame(noxflags.GameFlag16) {
		goto LABEL_31
	}
	result = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 2068))))
	if result == 0 {
		return result
	}
	if nox_server_teamByXxx_418AE0(int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 2068))))) != nil {
		goto LABEL_31
	}
	v4 = int32(uint8(*(*byte)(unsafe.Add(unsafe.Pointer(v1), 52))))
	if (noxflags.HasGame(noxflags.GameModeCTF|noxflags.GameModeFlagBall) || noxflags.HasGame(noxflags.GameModeKOTR) && nox_xxx_CheckGameplayFlags_417DA0(4)) && v4 > 2 {
		v4 = 2
	}
	if int32(uint8(sub_417DE0())) >= v4 {
		return 0
	}
	if noxflags.HasGame(noxflags.GameModeCTF | noxflags.GameModeFlagBall) {
		v5 = int32(uint8(sub_417DE0()))
		if v5 >= sub_417DC0() {
			return 0
		}
	}
LABEL_31:
	if noxflags.HasGame(noxflags.GameModeChat) {
		goto LABEL_26
	}
	if !noxflags.HasGame(noxflags.GameModeElimination) {
		goto LABEL_26
	}
	v6 = (*byte)(unsafe.Pointer(nox_common_playerInfoGetFirst_416EA0()))
	if v6 == nil {
		goto LABEL_26
	}
	for *((*int32)(unsafe.Add(unsafe.Pointer((*int32)(unsafe.Pointer(v6))), unsafe.Sizeof(int32(0))*535))) <= 0 {
		v6 = (*byte)(unsafe.Pointer(nox_common_playerInfoGetNext_416EE0((*nox_playerInfo)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v6)))))))))
		if v6 == nil {
			goto LABEL_26
		}
	}
	if sub_40AA00() != 0 {
		result = 0
	} else {
	LABEL_26:
		result = 1
	}
	return result
}
func nox_xxx_PtFuncCompare_40AE90(a1 unsafe.Pointer, a2 unsafe.Pointer) int32 {
	return int32(libc.StrCaseCmp(*((**byte)(unsafe.Add(unsafe.Pointer((**byte)(a1)), unsafe.Sizeof((*byte)(nil))*1))), *((**byte)(unsafe.Add(unsafe.Pointer((**byte)(a2)), unsafe.Sizeof((*byte)(nil))*1)))))
}
func sub_40AEB0(a1 unsafe.Pointer, a2 unsafe.Pointer) int32 {
	return int32(libc.StrCaseCmp((*byte)(a1), *((**byte)(unsafe.Add(unsafe.Pointer((**byte)(a2)), unsafe.Sizeof((*byte)(nil))*1)))))
}
func nox_init_sound_index_40AED0() {
	if dword_5d4594_3616 != 0 {
		return
	}
	var v0 *uint32 = (*uint32)(alloc.Calloc(1023, 8))
	dword_5d4594_3616 = uint32(uintptr(unsafe.Pointer(v0)))
	for i := int32(0); i < 1023; i++ {
		*(*uint32)(unsafe.Add(unsafe.Pointer(v0), unsafe.Sizeof(uint32(0))*uintptr(i*2+0))) = uint32(i)
		*(*uint32)(unsafe.Add(unsafe.Pointer(v0), unsafe.Sizeof(uint32(0))*uintptr(i*2+1))) = uint32(uintptr(unsafe.Pointer(table_5184[i])))
	}
	libc.Sort(*(*unsafe.Pointer)(unsafe.Pointer(&dword_5d4594_3616)), 1023, 8, nox_xxx_PtFuncCompare_40AE90)
}
func sub_40AF30() {
	if dword_5d4594_3616 != 0 {
		*(*unsafe.Pointer)(unsafe.Pointer(&dword_5d4594_3616)) = nil
		dword_5d4594_3616 = 0
	}
}
func nox_xxx_utilFindSound_40AF50(a1 *byte) int32 {
	var (
		v1     *int32
		result int32
	)
	v1 = (*int32)(libc.Search(unsafe.Pointer(a1), *(*unsafe.Pointer)(unsafe.Pointer(&dword_5d4594_3616)), 1023, 8, sub_40AEB0))
	if v1 != nil {
		result = *v1
	} else {
		result = 0
	}
	return result
}
func nox_xxx_getSndName_40AF80(a1 int32) *byte {
	return table_5184[a1]
}
func nox_xxx_soloGameEscMenuCallback_40AF90(a1 int32, a2 int32, a3 int8, a4 int32, a5 *uint8, a6 uint32) {
	var (
		v6       *byte
		v7       *byte
		v8       *byte
		v9       int32
		FileName [1024]byte
	)
	switch a3 {
	case 1:
		sub_446520(1, unsafe.Pointer(a5), int32(a6))
	case 2:
		nox_xxx_savePlayerMB_41C8F0((*byte)(unsafe.Pointer(a5)), a6)
		if noxflags.HasGame(noxflags.GameModeQuest) {
			if sub_4460B0() != 0 {
				sub_446140()
			} else {
				sub_41D170()
			}
		} else if sub_446030() != 0 && sub_446090() != 0 {
			nox_xxx_setContinueMenuOrHost_43DDD0(0)
			nox_game_exit_xxx_43DE60()
			sub_446060()
		}
	case 3:
		v6 = nox_fs_root()
		nox_sprintf(&FileName[0], CString("%s\\Save\\_temp_.dat"), v6)
		if nox_xxx_SavePlayerDataFromClient_41CD70(&FileName[0], a5, int32(a6)) != 0 {
			if nox_xxx_isQuest_4D6F50() != 0 && a1 == 31 {
				sub_4DCEE0(&FileName[0])
			} else {
				v7 = nox_xxx_cliPlrInfoLoadFromFile_41A2E0(&FileName[0], a1)
				if noxflags.HasGame(noxflags.GameModeQuest) {
					if v7 != nil {
						v8 = (*byte)(unsafe.Pointer(nox_common_playerInfoFromNum_417090(a1)))
						if v8 != nil {
							v9 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v8))), unsafe.Sizeof(uint32(0))*514))))
							if v9 != 0 {
								*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v9 + 748))) + 552))) = 0
							}
						}
					} else {
						nox_xxx_playerCallDisconnect_4DEAB0(a1, 4)
					}
				}
				nox_fs_remove(&FileName[0])
			}
		} else if noxflags.HasGame(noxflags.GameModeQuest) && a1 != 31 {
			nox_xxx_playerCallDisconnect_4DEAB0(a1, 4)
		}
	default:
		return
	}
}
func sub_40B170(a1 int32) int32 {
	var (
		v1     int32
		result int32
	)
	v1 = a1
	if a1 >= 0 {
		if a1 > 256 {
			v1 = 256
		}
	} else {
		v1 = 16
	}
	dword_5d4594_3624 = uint32(v1)
	dword_5d4594_3620 = uint32(uintptr(alloc.Calloc(int(v1), 168)))
	result = int32(dword_5d4594_3624)
	*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a1))), 0)) = 0
	if dword_5d4594_3624 > 0 {
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&result))), 0)) = 0
		for {
			sub_40B1F0(uint8(int8(result)))
			*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a1))), 0)) = uint8(int8(a1 + 1))
			result = a1
			if int32(uint8(int8(a1))) >= *(*int32)(unsafe.Pointer(&dword_5d4594_3624)) {
				break
			}
		}
	}
	return result
}
func sub_40B1F0(a1 uint8) *byte {
	var result *byte
	result = (*byte)(unsafe.Pointer(uintptr(dword_5d4594_3620 + uint32(int32(a1)*168))))
	*(*byte)(unsafe.Add(unsafe.Pointer(result), 4)) = 0
	*(*byte)(unsafe.Add(unsafe.Pointer(result), 5)) = byte(a1)
	*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(result))), unsafe.Sizeof(uint32(0))*4))) = 0
	*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(result))), unsafe.Sizeof(uint32(0))*3))) = 0
	*(*byte)(unsafe.Add(unsafe.Pointer(result), 24)) = 0
	*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(result))), unsafe.Sizeof(uint32(0))*40))) = 0
	*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(result))), unsafe.Sizeof(uint32(0))*41))) = 0
	*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(result))), unsafe.Sizeof(uint32(0))*5))) = 0
	*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(result))), unsafe.Sizeof(uint32(0))*38))) = 0
	*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(result))), unsafe.Sizeof(uint32(0))*2))) = 1
	*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(result))), unsafe.Sizeof(uint32(0))*39))) = 0
	return result
}
func sub_40B250(a1 int32, a2 uint8, a3 uint16, a4 unsafe.Pointer, a5 uint32) {
	var (
		v5  int32
		v6  int32
		v7  *uint32
		v8  unsafe.Pointer
		v9  int32
		v10 int32
		v11 *uint16
		v12 *uint16
		v13 int32
		v14 int32
		v15 int32
		v16 int32
		v17 int32
		v18 int32
		v19 uint32
	)
	if a4 == nil {
		return
	}
	if a5 == 0 {
		return
	}
	if int32(a2) >= *(*int32)(unsafe.Pointer(&dword_5d4594_3624)) {
		return
	}
	v5 = int32(dword_5d4594_3620 + uint32(int32(a2)*168))
	if *(*uint32)(unsafe.Pointer(uintptr(v5 + 16))) == 0 {
		return
	}
	*(*uint32)(unsafe.Pointer(uintptr(v5 + 156))) = nox_frame_xxx_2598000
	if uint32(a3) == *(*uint32)(unsafe.Pointer(uintptr(v5 + 8))) {
		libc.MemCpy(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v5 + 12)))+*(*uint32)(unsafe.Pointer(uintptr(v5 + 20))))), a4, int(a5))
		v6 = int32(*(*uint32)(unsafe.Pointer(uintptr(v5 + 8))) + 1)
		*(*uint32)(unsafe.Pointer(uintptr(v5 + 12))) += a5
		*(*uint32)(unsafe.Pointer(uintptr(v5 + 8))) = uint32(v6)
	} else {
		v7 = (*uint32)(alloc.Calloc(1, 28))
		if v7 != nil {
			v8 = alloc.Calloc(1, int(a5))
			*(*uint32)(unsafe.Add(unsafe.Pointer(v7), unsafe.Sizeof(uint32(0))*1)) = uint32(uintptr(v8))
			if v8 != nil {
				libc.MemCpy(v8, a4, int(a5))
				*(*uint16)(unsafe.Pointer(v7)) = a3
				*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v7))), unsafe.Sizeof(uint16(0))*4))) = uint16(a5)
				*(*uint32)(unsafe.Add(unsafe.Pointer(v7), unsafe.Sizeof(uint32(0))*5)) = 0
				*(*uint32)(unsafe.Add(unsafe.Pointer(v7), unsafe.Sizeof(uint32(0))*6)) = *(*uint32)(unsafe.Pointer(uintptr(v5 + 164)))
				v9 = int32(*(*uint32)(unsafe.Pointer(uintptr(v5 + 164))))
				if v9 != 0 {
					*(*uint32)(unsafe.Pointer(uintptr(v9 + 20))) = uint32(uintptr(unsafe.Pointer(v7)))
				}
				v10 = int32(*(*uint32)(unsafe.Pointer(uintptr(v5 + 160))))
				*(*uint32)(unsafe.Pointer(uintptr(v5 + 164))) = uint32(uintptr(unsafe.Pointer(v7)))
				if v10 == 0 {
					*(*uint32)(unsafe.Pointer(uintptr(v5 + 160))) = uint32(uintptr(unsafe.Pointer(v7)))
				}
			}
		}
	}
	v11 = *(**uint16)(unsafe.Pointer(uintptr(v5 + 160)))
	if v11 != nil {
		for {
			v12 = (*uint16)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v11))), unsafe.Sizeof(uint32(0))*5))))))
			v19 = *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v11))), unsafe.Sizeof(uint32(0))*5)))
			if *(*uint32)(unsafe.Pointer(uintptr(v5 + 8))) == uint32(*v11) {
				libc.MemCpy(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v5 + 12)))+*(*uint32)(unsafe.Pointer(uintptr(v5 + 20))))), *((*unsafe.Pointer)(unsafe.Add(unsafe.Pointer((*unsafe.Pointer)(unsafe.Pointer(v11))), unsafe.Sizeof(unsafe.Pointer(nil))*1))), int(int16(*(*uint16)(unsafe.Add(unsafe.Pointer(v11), unsafe.Sizeof(uint16(0))*4)))))
				v13 = int32(*(*uint32)(unsafe.Pointer(uintptr(v5 + 8))) + 1)
				*(*uint32)(unsafe.Pointer(uintptr(v5 + 12))) += uint32(int16(*(*uint16)(unsafe.Add(unsafe.Pointer(v11), unsafe.Sizeof(uint16(0))*4))))
				*(*uint32)(unsafe.Pointer(uintptr(v5 + 8))) = uint32(v13)
				v14 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v11))), unsafe.Sizeof(uint32(0))*6))))
				if v14 != 0 {
					*(*uint32)(unsafe.Pointer(uintptr(v14 + 20))) = *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v11))), unsafe.Sizeof(uint32(0))*5)))
				} else {
					*(*uint32)(unsafe.Pointer(uintptr(v5 + 160))) = *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v11))), unsafe.Sizeof(uint32(0))*5)))
				}
				v15 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v11))), unsafe.Sizeof(uint32(0))*5))))
				if v15 != 0 {
					*(*uint32)(unsafe.Pointer(uintptr(v15 + 24))) = *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v11))), unsafe.Sizeof(uint32(0))*6)))
				} else {
					*(*uint32)(unsafe.Pointer(uintptr(v5 + 164))) = *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v11))), unsafe.Sizeof(uint32(0))*6)))
				}
				*((*unsafe.Pointer)(unsafe.Add(unsafe.Pointer((*unsafe.Pointer)(unsafe.Pointer(v11))), unsafe.Sizeof(unsafe.Pointer(nil))*1))) = nil
				alloc.Free(unsafe.Pointer(v11))
				v12 = (*uint16)(unsafe.Pointer(uintptr(v19)))
			}
			v11 = v12
			if v12 == nil {
				break
			}
		}
	}
	v17 = int32(*(*uint32)(unsafe.Pointer(uintptr(v5 + 16))))
	v16 = int32(*(*uint32)(unsafe.Pointer(uintptr(v5 + 12))))
	*(*float32)(unsafe.Pointer(uintptr(v5 + 152))) = float32(float64(*(*uint32)(unsafe.Pointer(uintptr(v5 + 12)))) / float64(v17) * 100.0)
	if v16 == v17 {
		nox_xxx_netXfer_40B4B0(*(*uint32)(unsafe.Pointer(uintptr(v5))), int8(*(*uint8)(unsafe.Pointer(uintptr(v5 + 5)))))
		v18 = int32(*(*uint32)(unsafe.Pointer(uintptr(v5 + 20))))
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v18))), 0)) = *(*uint8)(unsafe.Pointer(uintptr(v5 + 5)))
		nox_xxx_soloGameEscMenuCallback_40AF90(a1-1, v18, int8(*(*uint8)(unsafe.Pointer(uintptr(v5 + 4)))), v5+24, *(**uint8)(unsafe.Pointer(uintptr(v5 + 20))), *(*uint32)(unsafe.Pointer(uintptr(v5 + 16))))
		if int32(*memmap.PtrUint8(6112660, 3628)) != 0 {
			*memmap.PtrUint8(6112660, 3628)--
		}
		*(*unsafe.Pointer)(unsafe.Pointer(uintptr(v5 + 20))) = nil
		sub_40B4E0(a2)
		sub_40B1F0(a2)
	}
}
func nox_xxx_netXfer_40B4B0(a1 uint32, a2 int8) int32 {
	var v4 [3]byte
	v4[0] = 194
	v4[1] = 4
	v4[2] = byte(a2)
	return nox_xxx_netSendSock_552640(a1, &v4[0], 3, int8(NOX_NET_SEND_NO_LOCK|NOX_NET_SEND_FLAG2))
}
func sub_40B4E0(a1 uint8) int32 {
	var (
		result int32
		v2     int32
		v3     int32
	)
	result = int32(a1) * 21
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_3620 + uint32(int32(a1)*168) + 160))))
	if v2 != 0 {
		for {
			v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 20))))
			*(*unsafe.Pointer)(unsafe.Pointer(uintptr(v2 + 4))) = nil
			alloc.Free(unsafe.Pointer(uintptr(v2)))
			v2 = v3
			if v3 == 0 {
				break
			}
		}
	}
	return result
}
func sub_40B530(a1 uint8, a2 int8) *byte {
	var result *byte
	result = (*byte)(unsafe.Pointer(uintptr(dword_5d4594_3620 + uint32(int32(a1)*168))))
	if *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(result))), unsafe.Sizeof(uint32(0))*4))) != 0 {
		nox_xxx_neXfer_40B590(*(*uint32)(unsafe.Pointer(result)), int8(*(*byte)(unsafe.Add(unsafe.Pointer(result), 5))), a2)
		if int32(*memmap.PtrUint8(6112660, 3628)) != 0 {
			*memmap.PtrUint8(6112660, 3628)--
		}
		sub_40B4E0(a1)
		result = sub_40B1F0(a1)
	}
	return result
}
func nox_xxx_neXfer_40B590(a1 uint32, a2 int8, a3 int8) int32 {
	var v4 [4]byte
	v4[2] = byte(a2)
	v4[0] = 194
	v4[1] = 6
	v4[3] = byte(a3)
	return nox_xxx_netSendSock_552640(a1, &v4[0], 4, int8(NOX_NET_SEND_NO_LOCK|NOX_NET_SEND_FLAG2))
}
func sub_40B5D0(a1 uint32, a2 int8, a3 *byte, a4 uint32, a5 int8) int32 {
	var (
		v5 uint32
		v6 *byte
		v7 *byte
		v9 int32
	)
	v5 = a4
	if a4 == 0 {
		return 0
	}
	v6 = sub_40B6D0((*uint8)(unsafe.Pointer(&v9)))
	v7 = v6
	if v6 == nil {
		return 0
	}
	*(*uint32)(unsafe.Pointer(v6)) = a1
	*(*byte)(unsafe.Add(unsafe.Pointer(v6), 4)) = byte(a2)
	if *a3 != 0 {
		libc.StrCpy((*byte)(unsafe.Add(unsafe.Pointer(v6), 24)), a3)
		v5 = a4
	}
	*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v6))), unsafe.Sizeof(uint32(0))*4))) = v5
	*(*byte)(unsafe.Add(unsafe.Pointer(v6), 5)) = byte(int8(v9))
	*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v6))), unsafe.Sizeof(uint32(0))*5))) = uint32(uintptr(alloc.Calloc(int(v5), 1)))
	*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v7))), unsafe.Sizeof(uint32(0))*39))) = nox_frame_xxx_2598000
	*memmap.PtrUint8(6112660, 3628)++
	sub_40B690(a1, int8(v9), a5)
	return 1
}
func sub_40B690(a1 uint32, a2 int8, a3 int8) int32 {
	var v4 [4]byte
	v4[2] = byte(a2)
	v4[0] = 194
	v4[1] = 1
	v4[3] = byte(a3)
	return nox_xxx_netSendSock_552640(a1, &v4[0], 4, int8(NOX_NET_SEND_NO_LOCK|NOX_NET_SEND_FLAG2))
}
func sub_40B6D0(a1 *uint8) *byte {
	var (
		v1     int32
		v2     *uint32
		result *byte
	)
	v1 = 0
	if *(*int32)(unsafe.Pointer(&dword_5d4594_3624)) <= 0 {
	LABEL_5:
		result = nil
		*a1 = 0
	} else {
		v2 = (*uint32)(unsafe.Pointer(uintptr(dword_5d4594_3620 + 16)))
		for *v2 != 0 {
			v1++
			v2 = (*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*42))
			if v1 >= *(*int32)(unsafe.Pointer(&dword_5d4594_3624)) {
				goto LABEL_5
			}
		}
		*a1 = uint8(int8(v1))
		result = (*byte)(unsafe.Pointer(uintptr(dword_5d4594_3620 + uint32(v1*168))))
	}
	return result
}
func sub_40B720(a1 int32, a2 uint8) *byte {
	sub_40B4E0(a2)
	return sub_40B1F0(a2)
}
func sub_40B740() {
	var (
		v0 uint8
		v2 uint8
	)
	v2 = 0
	if dword_5d4594_3624 > 0 {
		v0 = 0
		for {
			sub_40B4E0(v0)
			v0 = func() uint8 {
				p := &v2
				*p++
				return *p
			}()
			if int32(v2) >= *(*int32)(unsafe.Pointer(&dword_5d4594_3624)) {
				break
			}
		}
	}
	*(*unsafe.Pointer)(unsafe.Pointer(&dword_5d4594_3620)) = nil
}
func sub_40B790() int32 {
	var (
		result int32
		v1     uint8
		v2     int32
		v3     int32
		v4     uint8
	)
	result = int32(dword_5d4594_3624)
	v4 = 0
	if dword_5d4594_3624 > 0 {
		v1 = 0
		v2 = 0
		for {
			v3 = int32(dword_5d4594_3620 + uint32(v2*168))
			if *(*uint32)(unsafe.Pointer(uintptr(v3 + 16))) != 0 {
				if nox_frame_xxx_2598000 > (*(*uint32)(unsafe.Pointer(uintptr(v3 + 156))) + 900) {
					sub_40B530(v1, 3)
				}
			}
			result = int32(dword_5d4594_3624)
			v1 = func() uint8 {
				p := &v4
				*p++
				return *p
			}()
			v2 = int32(v4)
			if int32(v4) >= *(*int32)(unsafe.Pointer(&dword_5d4594_3624)) {
				break
			}
		}
	}
	return result
}
func sub_40B810(a1 int32, a2 int32, a3 int32) {
	nox_xxx_soloGameEscMenuCallback_40AF90(31, 0, int8(a1), int32(uintptr(memmap.PtrOff(6112660, 4664))), (*uint8)(unsafe.Pointer(uintptr(a2))), uint32(a3))
	sub_40B850(0, int8(a1))
}
func sub_40B850(a1 int32, a2 int8) {
	if int32(a2) == 2 && sub_446030() != 0 {
		sub_446070()
		if sub_446090() != 0 {
			nox_xxx_setContinueMenuOrHost_43DDD0(0)
			nox_game_exit_xxx_43DE60()
			sub_446060()
		}
	}
}
func sub_40B890(a1 int32) unsafe.Pointer {
	var (
		result unsafe.Pointer
		v2     int32
		v3     int32
	)
	if a1 >= 0 {
		*memmap.PtrUint16(6112660, 4660) = 256
		if a1 <= 256 {
			*memmap.PtrUint16(6112660, 4660) = uint16(int16(a1))
		}
	} else {
		*memmap.PtrUint16(6112660, 4660) = 16
	}
	result = alloc.Calloc(int(*memmap.PtrUint16(6112660, 4660)), 160)
	v2 = 0
	dword_5d4594_3632 = uint32(uintptr(result))
	if int32(*memmap.PtrUint16(6112660, 4660)) != 0 {
		v3 = 0
		for {
			sub_40B930(int32(uint32(v3) + dword_5d4594_3632))
			v2++
			result = unsafe.Pointer(uintptr(*memmap.PtrUint16(6112660, 4660)))
			v3 += 160
			if v2 >= int32(*memmap.PtrUint16(6112660, 4660)) {
				break
			}
		}
		*memmap.PtrUint16(6112660, 4662) = 0
	} else {
		*memmap.PtrUint16(6112660, 4662) = 0
	}
	return result
}
func sub_40B930(a1 int32) int32 {
	var result int32
	result = a1
	*(*uint32)(unsafe.Pointer(uintptr(a1))) = 0
	*(*uint8)(unsafe.Pointer(uintptr(a1 + 4))) = 0
	*(*uint16)(unsafe.Pointer(uintptr(a1 + 6))) = 0
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 8))) = 0
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 12))) = 1
	*(*uint16)(unsafe.Pointer(uintptr(a1 + 20))) = 0
	*(*uint8)(unsafe.Pointer(uintptr(a1 + 22))) = 0
	*(*uint8)(unsafe.Pointer(uintptr(a1 + 23))) = 0
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 16))) = 0
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 152))) = 0
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 156))) = 0
	return result
}
func sub_40B970() {
	var (
		v0 int32
		v1 int32
		v2 int32
		v3 int32
		v4 int32
		v5 int32
		v6 int32
		v7 int32
	)
	v0 = 0
	if int32(*memmap.PtrUint16(6112660, 4662)) != 0 {
		v7 = 0
		if int32(*memmap.PtrUint16(6112660, 4660)) > 0 {
			for {
				v1 = int32(dword_5d4594_3632 + uint32(v0))
				if int32(*(*uint16)(unsafe.Pointer(uintptr(dword_5d4594_3632 + uint32(v0) + 6)))) == 2 {
					v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 152))))
					v3 = 0
					if v2 != 0 {
						for v3 < 2 {
							v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 12))))
							if v4 != 0 {
								v5 = int32(nox_frame_xxx_2598000)
								if uint32(v4+90) < nox_frame_xxx_2598000 {
									if int32(*(*uint16)(unsafe.Pointer(uintptr(v2 + 16)))) >= 20 {
										if int32(*(*uint16)(unsafe.Pointer(uintptr(v1 + 6)))) == 2 {
											*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v5))), 0)) = *(*uint8)(unsafe.Pointer(uintptr(v1 + 4)))
											sub_40BB20(*(*uint32)(unsafe.Pointer(uintptr(v1))), v5, 2)
											break
										}
									} else {
										sub_40BA90(*(*uint32)(unsafe.Pointer(uintptr(v1))), int8(*(*uint8)(unsafe.Pointer(uintptr(v1 + 4)))), int16(*(*uint16)(unsafe.Pointer(uintptr(v2)))), int16(*(*uint16)(unsafe.Pointer(uintptr(v2 + 8)))), *(*unsafe.Pointer)(unsafe.Pointer(uintptr(v2 + 4))))
										v6 = int32(nox_frame_xxx_2598000)
										*(*uint16)(unsafe.Pointer(uintptr(v2 + 16)))++
										*(*uint32)(unsafe.Pointer(uintptr(v2 + 12))) = uint32(v6)
									}
								}
							} else {
								sub_40BA90(*(*uint32)(unsafe.Pointer(uintptr(v1))), int8(*(*uint8)(unsafe.Pointer(uintptr(v1 + 4)))), int16(*(*uint16)(unsafe.Pointer(uintptr(v2)))), int16(*(*uint16)(unsafe.Pointer(uintptr(v2 + 8)))), *(*unsafe.Pointer)(unsafe.Pointer(uintptr(v2 + 4))))
								*(*uint32)(unsafe.Pointer(uintptr(v2 + 12))) = nox_frame_xxx_2598000
								*(*uint32)(unsafe.Pointer(uintptr(v1 + 12)))++
							}
							v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 20))))
							v3++
							if v2 == 0 {
								break
							}
						}
					}
				}
				v0 += 160
				v7++
				if v7 >= int32(*memmap.PtrUint16(6112660, 4660)) {
					break
				}
			}
		}
	}
}
func sub_40BA90(a1 uint32, a2 int8, a3 int16, a4 int16, a5 unsafe.Pointer) int32 {
	var (
		v6 int32
		v7 int32
	)
	*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v6))), 2)) = uint8(a2)
	*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v7))), unsafe.Sizeof(uint16(0))*1)) = uint16(a4)
	*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v7))), unsafe.Sizeof(uint16(0))*0)) = uint16(a3)
	*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v6))), unsafe.Sizeof(uint16(0))*0)) = 706
	*memmap.PtrUint32(6112660, 3640) = uint32(v7)
	*memmap.PtrUint32(6112660, 3636) = uint32(v6)
	libc.MemCpy(memmap.PtrOff(6112660, 3644), a5, int(a4))
	nox_xxx_netSendSock_552640(a1, (*byte)(memmap.PtrOff(6112660, 3636)), int32(a4)+8, 0)
	return nox_xxx_netSendReadPacket_5528B0(a1, 1)
}
func sub_40BB20(a1 uint32, a2 int32, a3 int8) *byte {
	var (
		v3     int8
		result *byte
		v5     *byte
		v6     int32
		v7     int32
	)
	v3 = int8(a2)
	result = sub_40BC10(int32(a1), int8(a2))
	v5 = result
	if result != nil {
		*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&a2))), unsafe.Sizeof(uint16(0))*0)) = 1474
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a2))), 2)) = uint8(v3)
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a2))), unsafe.Sizeof(int32(0))-1)) = uint8(a3)
		nox_xxx_netSendSock_552640(a1, (*byte)(unsafe.Pointer(&a2)), 4, int8(NOX_NET_SEND_NO_LOCK|NOX_NET_SEND_FLAG2))
		sub_40BBC0(int32(*(*uint32)(unsafe.Pointer(v5))), int8(*(*byte)(unsafe.Add(unsafe.Pointer(v5), 22))))
		if int32(*memmap.PtrUint16(6112660, 4662)) != 0 {
			*memmap.PtrUint16(6112660, 4662)--
		}
		v6 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v5))), unsafe.Sizeof(uint32(0))*38))))
		if v6 != 0 {
			for {
				v7 = int32(*(*uint32)(unsafe.Pointer(uintptr(v6 + 20))))
				*(*unsafe.Pointer)(unsafe.Pointer(uintptr(v6 + 4))) = nil
				alloc.Free(unsafe.Pointer(uintptr(v6)))
				v6 = v7
				if v7 == 0 {
					break
				}
			}
		}
		result = (*byte)(unsafe.Pointer(uintptr(sub_40B930(int32(uintptr(unsafe.Pointer(v5)))))))
	}
	return result
}
func sub_40BBC0(a1 int32, a2 int8) {
	if int32(a2) == 2 {
		if sub_446030() != 0 {
			sub_446070()
			if sub_446090() != 0 {
				nox_xxx_setContinueMenuOrHost_43DDD0(0)
				nox_game_exit_xxx_43DE60()
				sub_446060()
			}
		}
	} else if int32(a2) == 3 {
		nox_xxx_setContinueMenuOrHost_43DDD0(0)
		nox_game_exit_xxx_43DE60()
	}
}
func sub_40BC10(a1 int32, a2 int8) *byte {
	var (
		v2 int32
		i  int32
	)
	v2 = 0
	if int32(*memmap.PtrUint16(6112660, 4660)) <= 0 {
		return nil
	}
	for i = int32(dword_5d4594_3632); *(*uint32)(unsafe.Pointer(uintptr(i))) != uint32(a1) || int32(*(*uint8)(unsafe.Pointer(uintptr(i + 4)))) != int32(a2); i += 160 {
		if func() int32 {
			p := &v2
			*p++
			return *p
		}() >= int32(*memmap.PtrUint16(6112660, 4660)) {
			return nil
		}
	}
	return (*byte)(unsafe.Pointer(uintptr(dword_5d4594_3632 + uint32(v2*160))))
}
func sub_40BC60(a1 int32, a2 int32, a3 *byte, a4 int32, a5 int32, a6 int32) int32 {
	var (
		v6  *byte
		v7  int32
		v8  int32
		v10 *uint16
		v11 *uint16
		v12 int16
		v13 *byte
		v14 uint32
		v15 int32
		v16 int32
		v17 int32
		v18 int32
		v19 uint32
		v20 int32
		v21 int32
		v22 int32
		v23 *byte
	)
	v6 = sub_40BF10((*uint8)(unsafe.Pointer(&v20)))
	if v6 == nil {
		return 0
	}
	v7 = a5
	if a5 == 0 || a4 == 0 {
		return 0
	}
	if !noxflags.HasGame(noxflags.GameHost) {
		v8 = nox_xxx_netGet_43C750()
		goto LABEL_12
	}
	if !noxflags.HasGame(noxflags.GameClient) {
		return 0
	}
	if a6 == 0 {
		sub_40B810(a2, a4, a5)
		return 1
	}
	if a1 == 31 {
		sub_40B810(a2, a4, a5)
		return 1
	}
	v8 = a1 + 1
LABEL_12:
	*memmap.PtrUint16(6112660, 4662)++
	v19 = uint32(v8)
	v21 = a5
	v23 = (*byte)(unsafe.Pointer(uintptr(a4)))
	v22 = 0
	if int32((uint32(a5-1)>>9)+1) > 0 {
		for {
			v10 = (*uint16)(alloc.Calloc(1, 28))
			v11 = v10
			if v10 == nil {
				return 0
			}
			*v10 = uint16(int16(v22 + 1))
			v12 = 512
			if v7 <= 512 {
				v12 = int16(v7)
			}
			*(*uint16)(unsafe.Add(unsafe.Pointer(v11), unsafe.Sizeof(uint16(0))*4)) = uint16(v12)
			v13 = (*byte)(alloc.Calloc(1, int(v12)))
			v14 = uint32(int16(*(*uint16)(unsafe.Add(unsafe.Pointer(v11), unsafe.Sizeof(uint16(0))*4))))
			*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v11))), unsafe.Sizeof(uint32(0))*3))) = 0
			*(*uint16)(unsafe.Add(unsafe.Pointer(v11), unsafe.Sizeof(uint16(0))*8)) = 0
			*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v11))), unsafe.Sizeof(uint32(0))*1))) = uint32(uintptr(unsafe.Pointer(v13)))
			libc.MemCpy(unsafe.Pointer(v13), unsafe.Pointer(v23), int(v14))
			v15 = int32(int16(*(*uint16)(unsafe.Add(unsafe.Pointer(v11), unsafe.Sizeof(uint16(0))*4))))
			v16 = v21 - v15
			v17 = int32(uintptr(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v23), v15)))))
			*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v11))), unsafe.Sizeof(uint32(0))*6))) = *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v6))), unsafe.Sizeof(uint32(0))*39)))
			*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v11))), unsafe.Sizeof(uint32(0))*5))) = 0
			v18 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v6))), unsafe.Sizeof(uint32(0))*39))))
			v23 = (*byte)(unsafe.Pointer(uintptr(v17)))
			v21 = v16
			if v18 != 0 {
				*(*uint32)(unsafe.Pointer(uintptr(v18 + 20))) = uint32(uintptr(unsafe.Pointer(v11)))
			} else {
				*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v6))), unsafe.Sizeof(uint32(0))*38))) = uint32(uintptr(unsafe.Pointer(v11)))
			}
			*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v6))), unsafe.Sizeof(uint32(0))*39))) = uint32(uintptr(unsafe.Pointer(v11)))
			if func() int32 {
				p := &v22
				*p++
				return *p
			}() >= int32((uint32(a5-1)>>9)+1) {
				v7 = a5
				break
			}
			v7 = v16
		}
	}
	*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v6))), unsafe.Sizeof(uint32(0))*4))) = uint32(v22)
	*(*byte)(unsafe.Add(unsafe.Pointer(v6), 4)) = 0
	*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v6))), unsafe.Sizeof(uint32(0))*3))) = 1
	*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v6))), unsafe.Sizeof(uint16(0))*10))) = 512
	*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v6))), unsafe.Sizeof(uint32(0))*2))) = uint32(v7)
	*(*uint32)(unsafe.Pointer(v6)) = v19
	*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v6))), unsafe.Sizeof(uint16(0))*3))) = 1
	if a3 != nil {
		libc.StrCpy((*byte)(unsafe.Add(unsafe.Pointer(v6), 23)), a3)
		v7 = a5
	}
	*(*byte)(unsafe.Add(unsafe.Pointer(v6), 22)) = byte(int8(a2))
	nox_xxx_netXferMsg_40BE80(v19, int8(a2), a3, v7, int8(v20))
	return 1
}
func nox_xxx_netXferMsg_40BE80(a1 uint32, a2 int8, a3 *byte, a4 int32, a5 int8) int32 {
	var v6 [140]byte
	v6[0] = 194
	v6[1] = 0
	v6[2] = byte(a2)
	*(*uint32)(unsafe.Pointer(&v6[4])) = uint32(a4)
	if a3 != nil {
		libc.StrCpy(&v6[8], a3)
	} else {
		v6[8] = 0
	}
	v6[136] = byte(a5)
	return nox_xxx_netSendSock_552640(a1, &v6[0], 140, int8(NOX_NET_SEND_NO_LOCK|NOX_NET_SEND_FLAG2))
}
func sub_40BF10(a1 *uint8) *byte {
	var (
		v1 int32
		i  *uint32
	)
	v1 = 0
	if int32(*memmap.PtrUint16(6112660, 4660)) <= 0 {
		return nil
	}
	for i = (*uint32)(unsafe.Pointer(uintptr(dword_5d4594_3632 + 8))); int32(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(i))), -int(unsafe.Sizeof(uint16(0))*1))))) != 0 || *i != 0; i = (*uint32)(unsafe.Add(unsafe.Pointer(i), unsafe.Sizeof(uint32(0))*40)) {
		if func() int32 {
			p := &v1
			*p++
			return *p
		}() >= int32(*memmap.PtrUint16(6112660, 4660)) {
			return nil
		}
	}
	*a1 = uint8(int8(v1))
	return (*byte)(unsafe.Pointer(uintptr(dword_5d4594_3632 + uint32(v1*160))))
}
func sub_40BF60(a1 int32, a2 int8, a3 int32) *byte {
	var (
		result *byte
		v4     *uint16
		v5     int32
		v6     int32
	)
	result = sub_40BC10(a1, a2)
	if result != nil {
		v4 = (*uint16)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(result))), unsafe.Sizeof(uint32(0))*38))))))
		if v4 != nil {
			for int32(*v4) != a3 {
				v4 = (*uint16)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v4))), unsafe.Sizeof(uint32(0))*5))))))
				if v4 == nil {
					return result
				}
			}
			v5 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v4))), unsafe.Sizeof(uint32(0))*5))))
			if v5 != 0 {
				*(*uint32)(unsafe.Pointer(uintptr(v5 + 24))) = *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v4))), unsafe.Sizeof(uint32(0))*6)))
			} else {
				*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(result))), unsafe.Sizeof(uint32(0))*39))) = *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v4))), unsafe.Sizeof(uint32(0))*6)))
			}
			v6 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v4))), unsafe.Sizeof(uint32(0))*6))))
			if v6 != 0 {
				*(*uint32)(unsafe.Pointer(uintptr(v6 + 20))) = *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v4))), unsafe.Sizeof(uint32(0))*5)))
			} else {
				*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(result))), unsafe.Sizeof(uint32(0))*38))) = *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v4))), unsafe.Sizeof(uint32(0))*5)))
			}
			*((*unsafe.Pointer)(unsafe.Add(unsafe.Pointer((*unsafe.Pointer)(unsafe.Pointer(v4))), unsafe.Sizeof(unsafe.Pointer(nil))*1))) = nil
			alloc.Free(unsafe.Pointer(v4))
		}
	}
	return result
}
func sub_40BFF0(a1 int32, a2 int8, a3 *byte) *byte {
	var result *byte
	result = a3
	if uint32(uint8(uintptr(unsafe.Pointer(a3)))) < uint32(*memmap.PtrUint16(6112660, 4660)) {
		result = (*byte)(unsafe.Pointer(uintptr(dword_5d4594_3632 + uint32(int32(uint8(uintptr(unsafe.Pointer(a3))))*160))))
		*(*byte)(unsafe.Add(unsafe.Pointer(result), 4)) = byte(a2)
		*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(result))), unsafe.Sizeof(uint16(0))*3))) = 2
	}
	return result
}
func sub_40C030(a1 int32, a2 int8) *byte {
	var (
		result *byte
		v3     *byte
	)
	result = sub_40BC10(a1, a2)
	v3 = result
	if result != nil {
		sub_40B850(a1, int8(*(*byte)(unsafe.Add(unsafe.Pointer(result), 22))))
		*memmap.PtrUint16(6112660, 4662)--
		result = (*byte)(unsafe.Pointer(uintptr(sub_40B930(int32(uintptr(unsafe.Pointer(v3)))))))
	}
	return result
}
func sub_40C070(a1 int32, a2 int32, a3 int8) *byte {
	var (
		result *byte
		v4     *byte
		v5     int32
		v6     int32
	)
	result = sub_40BC10(a1, a3)
	v4 = result
	if result != nil {
		sub_40BBC0(int32(*(*uint32)(unsafe.Pointer(result))), int8(*(*byte)(unsafe.Add(unsafe.Pointer(result), 22))))
		v5 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v4))), unsafe.Sizeof(uint32(0))*38))))
		if v5 != 0 {
			for {
				v6 = int32(*(*uint32)(unsafe.Pointer(uintptr(v5 + 20))))
				*(*unsafe.Pointer)(unsafe.Pointer(uintptr(v5 + 4))) = nil
				alloc.Free(unsafe.Pointer(uintptr(v5)))
				v5 = v6
				if v6 == 0 {
					break
				}
			}
		}
		result = (*byte)(unsafe.Pointer(uintptr(sub_40B930(int32(uintptr(unsafe.Pointer(v4)))))))
	}
	return result
}
func sub_40C0D0() {
	*(*unsafe.Pointer)(unsafe.Pointer(&dword_5d4594_3632)) = nil
}
func sub_40C0E0(a1 int32) {
	var (
		v1 int32
		v2 int32
		v3 int32
	)
	v1 = 0
	if int32(*memmap.PtrUint16(6112660, 4660)) != 0 {
		v2 = int32(dword_5d4594_3632)
		v3 = 0
		for {
			if int32(*(*uint16)(unsafe.Pointer(uintptr(v3 + v2 + 6)))) == 2 && *(*uint32)(unsafe.Pointer(uintptr(v3 + v2))) == uint32(a1) {
				sub_40BB20(*(*uint32)(unsafe.Pointer(uintptr(v3 + v2))), int32(*(*uint8)(unsafe.Pointer(uintptr(v3 + v2 + 4)))), 1)
				v2 = int32(dword_5d4594_3632)
			}
			v1++
			v3 += 160
			if v1 >= int32(*memmap.PtrUint16(6112660, 4660)) {
				break
			}
		}
	}
}
func sub_40C140(a1 int32, a2 int32, a3 *uint32) int32 {
	var (
		v3     int32
		result int32
		v5     int32
		v6     int32
	)
	if a2 == -2147221396 || a2 != 0 {
		if sub_41E300(11) != 0 {
			dword_5d4594_2660652 = uint32(a2)
			v5 = sub_41E2F0()
			sub_41DA70(v5, 5)
			v6 = sub_41E2F0()
			sub_41DA70(v6, 9)
		}
		result = 0
	} else {
		sub_41E4D0(a3)
		v3 = sub_41E2F0()
		sub_41DA70(v3, 2)
		result = 0
	}
	return result
}
func sub_40C1B0(a1 int32, a2 int32, a3 int32) int32 {
	var v3 int32
	if a2 == -2147221402 && (sub_41E2F0() == 7 || sub_41E2F0() == 5) {
		dword_5d4594_2660652 = 0x80040066
		v3 = sub_41E2F0()
		sub_41DA70(v3, 5)
	}
	return 0
}
func sub_40C1F0(a1 int32, a2 int32, a3 unsafe.Pointer) int32 {
	if a2 != 0 || sub_41E2F0() != 7 || a3 == nil {
		return 0
	}
	sub_446520(0, a3, int32(libc.StrLen((*byte)(a3))+1))
	sub_446D50()
	return 0
}
func sub_40C240(a1 int32, a2 int32, a3 int32) int32 {
	if a2 != 0 {
		if dword_5d4594_528252 != 0 && dword_5d4594_528256 != 0 {
			sub_41E470()
			return 0
		}
	} else {
		if sub_41E2F0() == 7 {
			sub_41ECB0(a3)
			sub_41EBB0(a3)
			sub_447470_wol_chat()
			sub_448640()
		} else if sub_41E2F0() == 8 {
			sub_416690()
		}
		sub_41F230(0, 1)
		if dword_5d4594_528252 != 0 && dword_5d4594_528256 != 0 {
			sub_41E370()
			sub_41E4B0(0)
		}
		*memmap.PtrUint32(0x587000, 25736) = 0
	}
	return 0
}
func sub_40C2E0(a1 int32, a2 int32, a3 int32, a4 *byte) int32 {
	var (
		result int32
		v5     int32
		v6     int32
		v7     int32
		v8     int32
		v9     int32
	)
	if sub_41E2F0() != 7 && sub_41E2F0() != 8 {
		return 0
	}
	if a2 <= -2147221390 {
		if a2 >= -2147221392 {
		LABEL_10:
			dword_5d4594_2660652 = uint32(a2)
			v7 = sub_41E2F0()
			sub_41DA70(v7, 5)
			return 0
		}
		if a2 != -2147221396 {
			if a2 != -2147221394 {
				return 0
			}
			goto LABEL_10
		}
		if sub_41E300(11) != 0 {
			dword_5d4594_2660652 = 2147745900
			v5 = sub_41E2F0()
			sub_41DA70(v5, 5)
			v6 = sub_41E2F0()
			sub_41DA70(v6, 9)
			return 0
		}
		return 0
	}
	if a2 != 0 {
		return 0
	}
	v8 = int32(*(*uint32)(unsafe.Pointer(a4)))
	if (v8 & 0x8000) == 0 {
	LABEL_15:
		sub_41F230(0, 1)
		sub_41F520(a4)
		return 0
	}
	sub_41EBB0(a3)
	if sub_41E2F0() == 7 {
		sub_447470_wol_chat()
		sub_448640()
		goto LABEL_15
	}
	if sub_41E2F0() != 8 {
		goto LABEL_15
	}
	libc.MemCpy(memmap.PtrOff(0x85B3FC, 10308), unsafe.Pointer(a4), 108)
	if *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a4))), unsafe.Sizeof(uint32(0))*5))) != 0 {
		sub_40DB20(int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a4))), unsafe.Sizeof(uint32(0))*5)))))
		result = 1
	} else {
		if (*a4 & 1) == 0 {
			v9 = sub_41E2F0()
			sub_41DA70(v9, 19)
		}
		result = 1
	}
	return result
}
func sub_40C420(a1 int32, a2 int32, a3 int32, a4 *int32) int32 {
	var (
		v4 int32
		v6 int32
	)
	if sub_41E2F0() == 7 || sub_41E2F0() == 8 {
		if a2 == -2147220994 {
			dword_5d4594_2660652 = 0x800401FE
			v6 = sub_41E2F0()
			sub_41DA70(v6, 5)
		} else if a2 == 0 {
			sub_41F620(int32(uintptr(unsafe.Pointer(a4))))
			if sub_41F9E0(int32(uintptr(unsafe.Pointer(a4)))) != 0 {
				sub_41F4B0()
				if sub_41E2F0() == 7 {
					sub_447590_wol_chat()
				}
				sub_41F230(0, -1)
				sub_41EBB0(0)
			} else {
				sub_41F230(0, -1)
			}
			if sub_41E2F0() != 7 {
				if *a4&1 != 0 {
					v4 = *a4
					if (v4 & 0x8000) == 0 {
						sub_40D380()
						return 0
					}
				}
			}
		}
	}
	return 0
}
func sub_40C4E0(a1 int32, a2 int32, a3 int32, a4 int32) int32 {
	return 0
}
func sub_40C4F0(a1 int32, a2 int32, a3 int32) int32 {
	return 0
}
func sub_40C500(a1 int32, a2 int32, a3 int32) int32 {
	return 0
}
func sub_40C510(a1 int32, a2 int32, a3 int32) int32 {
	return 0
}
func sub_40C520(a1 int32, a2 int32, a3 int32) int32 {
	return 0
}
func sub_40C530(a1 int32, a2 int32, a3 int32, a4 *uint8) int32 {
	if sub_41E2F0() == 7 && a2 == 0 && sub_41F9E0(a3) == 0 {
		sub_448650(a4, (*libc.WChar)(memmap.PtrOff(6112660, 8908)))
		sub_447250((*byte)(unsafe.Pointer(uintptr(a3+36))), int32(uintptr(memmap.PtrOff(6112660, 8908))), 1, 0)
	}
	return 0
}
func sub_40C580(a1 int32, a2 int32, a3 int32, a4 int32, a5 *uint8) int32 {
	if sub_41E2F0() == 7 && a2 == 0 {
		sub_448650(a5, (*libc.WChar)(memmap.PtrOff(6112660, 8908)))
		sub_447250((*byte)(unsafe.Pointer(uintptr(a4+36))), int32(uintptr(memmap.PtrOff(6112660, 8908))), 0, 0)
	}
	return 0
}
func sub_40C5C0(a1 int32, a2 int32, a3 int32, a4 int32, a5 *uint8) int32 {
	if sub_41E2F0() == 7 && a2 == 0 {
		sub_448650(a5, (*libc.WChar)(memmap.PtrOff(6112660, 8908)))
		sub_4471A0((*byte)(unsafe.Pointer(uintptr(a4+36))), int32(uintptr(memmap.PtrOff(6112660, 8908))), 0, 0)
	}
	return 0
}
func sub_40C600(a1 int32, a2 int32, a3 int32, a4 *uint8) int32 {
	if sub_41E2F0() == 7 && a2 == 0 && sub_41F9E0(a3) == 0 {
		sub_448650(a4, (*libc.WChar)(memmap.PtrOff(6112660, 8908)))
		sub_4471A0((*byte)(unsafe.Pointer(uintptr(a3+36))), int32(uintptr(memmap.PtrOff(6112660, 8908))), 1, 0)
	}
	return 0
}
func sub_40C650(a1 int32, a2 int32, a3 int32, a4 int32) int32 {
	return 0
}
func sub_40C660(a1 int32, a2 int32, a3 int32, a4 int32, a5 int32) int32 {
	return 0
}
func sub_40C670(a1 int32, a2 int32, a3 int32, a4 int32, a5 int32) int32 {
	if a2 == 0 {
		sub_416670(a5)
	}
	return 0
}
func sub_40C690(a1 int32, a2 int32, a3 int32, a4 int32, a5 int32) int32 {
	if sub_41E2F0() == 7 && a2 == 0 {
		sub_41F620(a4)
		if sub_41F9E0(a4) != 0 {
			sub_41F4B0()
			sub_41F230(0, -1)
			sub_41EBB0(0)
			sub_4474C0_wol_chat(0, a5+36)
			return 0
		}
		sub_4474C0_wol_chat(a4+36, a5+36)
		sub_41F230(0, -1)
	}
	return 0
}
func sub_40C710(a1 int32, a2 int32, a3 int32) int32 {
	return 0
}
func sub_40C720(a1 int32, a2 int32, a3 *uint32) int32 {
	if sub_41E2F0() == 7 || sub_41E2F0() == 8 {
		if a2 != 0 {
			if uint32(a2) <= 0x4012F || uint32(a2) > 262450 {
				sub_4490C0_wol_dialogs(0)
				return 0
			}
		} else if sub_41E2F0() == 8 {
			if *a3 == uint32(sub_420100()) {
				sub_4490C0_wol_dialogs(int32(uintptr(unsafe.Pointer(a3))))
				return 0
			}
		} else if *a3 == 0 {
			sub_4490C0_wol_dialogs(int32(uintptr(unsafe.Pointer(a3))))
			return 0
		}
		sub_4491B0_wol_dialogs()
	}
	return 0
}
func sub_40C7A0(a1 int32, a2 int32) int32 {
	if sub_41E2F0() == 7 {
		if a2 != 0 {
			sub_449240_wol_dialogs()
			return 0
		}
		sub_449200_wol_dialogs()
	}
	return 0
}
func sub_40C7D0(a1 int32, a2 int32, a3 int32, a4 *uint8) int32 {
	if sub_41E2F0() == 7 {
		if a2 == 0 {
			sub_448650(a4, (*libc.WChar)(memmap.PtrOff(6112660, 8908)))
			sub_447380((*byte)(unsafe.Pointer(uintptr(a3+36))), (*byte)(memmap.PtrOff(6112660, 8908)))
			return 0
		}
	} else if sub_43C710() != 0 {
		sub_43D260(a3+36, int32(uintptr(unsafe.Pointer(a4))))
	}
	return 0
}
func sub_40C830(a1 int32, a2 int32, a3 int32) int32 {
	var (
		v3 int32
		v4 int32
	)
	if sub_41E2F0() == 7 && a2 == 0 && sub_41E300(11) != 0 {
		dword_5d4594_2660652 = 1
		v3 = sub_41E2F0()
		sub_41DA70(v3, 5)
		v4 = sub_41E2F0()
		sub_41DA70(v4, 9)
	}
	return 0
}
func sub_40C880(a1 int32, a2 int32, a3 *byte, a4 int32, a5 int32) int32 {
	var v5 int32
	if sub_41E2F0() == 7 && a2 == 0 {
		sub_41F840(a3, a4)
		v5 = sub_41E2F0()
		sub_41DA70(v5, 11)
	}
	return 0
}
func sub_40C8C0(a1 int32, a2 int32, a3 int32, a4 int32) int32 {
	if sub_41E2F0() == 7 && a2 == 0 && a4 != 0 {
		sub_447540_wol_chat(a3)
	}
	return 0
}
func sub_40C8F0(a1 int32, a2 int32, a3 int32, a4 int32) int32 {
	var v5 int32
	if sub_41E2F0() == 8 && a2 == 0 {
		dword_5d4594_2660032 = *(*uint32)(unsafe.Pointer(uintptr(a4)))
		libc.StrNCpy((*byte)(memmap.PtrOff(0x85B3FC, 10395)), (*byte)(unsafe.Pointer(uintptr(a4+401))), 9)
		*memmap.PtrUint8(0x85B3FC, 10404) = 0
		if *memmap.PtrUint32(6112660, 10980) != 0 {
			*memmap.PtrUint32(6112660, 10980) = 0
			sub_41D5E0()
			return 0
		}
		v5 = sub_41E2F0()
		sub_41DA70(v5, 19)
	}
	return 0
}
func sub_40C960(a1 int32, a2 int32, a3 *uint32) int32 {
	var (
		v3 int32
		v4 int32
		v6 int16
	)
	if sub_41E2F0() == 2 && sub_40E0B0() != 0 {
		if a2 != 0 {
			if sub_41E300(11) == 0 {
				return 0
			}
			dword_5d4594_2660652 = uint32(a2)
			v3 = sub_41E2F0()
			sub_41DA70(v3, 5)
			v6 = 9
		} else {
			sub_4201B0(a3)
			v6 = 1
		}
		v4 = sub_41E2F0()
		sub_41DA70(v4, v6)
		return 0
	}
	return 0
}
func sub_40C9D0(a1 int32, a2 int32, a3 *int32) int32 {
	var v4 int32
	if a2 != 0 {
		return 0
	}
	if sub_41E2F0() != 7 {
		if *a3&1 != 0 {
			v4 = *a3
			if (v4 & 0x8000) == 0 {
				sub_40D380()
			}
		}
		return 0
	}
	sub_41F620(int32(uintptr(unsafe.Pointer(a3))))
	if sub_41F9E0(int32(uintptr(unsafe.Pointer(a3)))) != 0 {
		sub_41EBB0(0)
		sub_447590_wol_chat()
	}
	sub_41F230(0, -1)
	return 0
}
func sub_40CA40(a1 int32, a2 int32) int32 {
	return 0
}
func sub_40CA50(a1 int32, a2 int32) int32 {
	return 0
}
func sub_40CA60(a1 int32, a2 int32, a3 unsafe.Pointer) int32 {
	var (
		result int32
		v4     int32
		v5     int32
		v6     int32
		v7     int32
		v8     int32
	)
	result = sub_40E0B0()
	if result != 0 {
		if a2 != 0 {
			if dword_5d4594_528252 != 0 && dword_5d4594_528256 != 0 {
				sub_41E470()
				return 0
			}
			if sub_41E300(11) != 0 {
				dword_5d4594_2660652 = uint32(a2)
				v4 = sub_41E2F0()
				sub_41DA70(v4, 6)
				v5 = sub_41E2F0()
				sub_41DA70(v5, 9)
				return 0
			}
		} else {
			if sub_41E2F0() == 5 {
				sub_41FF30(a3)
				sub_41E300(7)
				v6 = sub_41E2F0()
				sub_41DA70(v6, 7)
				sub_4207F0(3)
				v7 = sub_41FC40()
				sub_40DB50(v7+1, int32(uintptr(unsafe.Pointer(&a3))))
				sub_40D670(int32(uintptr(a3)))
				return 0
			}
			if sub_41E2F0() == 3 {
				sub_446A90()
				sub_446AD0(nox_game_showWolChat_447620)
				sub_41FF30(a3)
				v8 = sub_41E2F0()
				sub_41DA70(v8, 7)
				return 0
			}
			if dword_5d4594_528252 != 0 && dword_5d4594_528256 != 0 {
				sub_43AA70()
			}
		}
		result = 0
	}
	return result
}
func sub_40CB90(a1 int32, a2 int32, a3 int32) int32 {
	return 0
}
func sub_40CBA0(a1 int32, a2 int32, a3 int32) int32 {
	if sub_41E2F0() == 7 && a2 == 0 {
		nox_swprintf((*libc.WChar)(memmap.PtrOff(6112660, 8908)), libc.CWString("%S"), a3)
		sub_447310(0, int32(uintptr(memmap.PtrOff(6112660, 8908))))
	}
	return 0
}
func sub_40CBE0(a1 int32, a2 int32) int32 {
	var (
		v2     int32
		v3     int32
		result int32
	)
	if uint32(a2) > 0x4012F {
		return 0
	}
	if uint32(a2) == 0x4012F {
		if sub_40CD00() != 0 {
			if sub_43C710() != 0 {
				sub_44B010()
				sub_40E0A0()
			} else {
				sub_41E300(9)
				sub_41F4B0()
				sub_41EC30()
				sub_446490(0)
				nox_xxx____setargv_4_44B000()
				sub_40E0A0()
				nox_game_checkStateWol_43C260()
				sub_44A400()
			}
			return 0
		}
		if sub_41E2F0() == 9 {
			sub_40E0A0()
			dword_5d4594_10988 = 0
			return 0
		}
		if sub_41E2F0() == 3 {
			nox_xxx_officialStringCmp_41FDE0()
		}
		dword_5d4594_10988 = 0
		return 0
	}
	switch uint32(a2) + 0x7FFBFF99 {
	case 0:
		fallthrough
	case 1:
		fallthrough
	case 2:
		fallthrough
	case 3:
		fallthrough
	case 4:
		fallthrough
	case 14:
		fallthrough
	case 15:
		if sub_41E300(11) == 0 {
			return 0
		}
		dword_5d4594_2660652 = uint32(a2)
		v2 = sub_41E2F0()
		sub_41DA70(v2, 5)
		v3 = sub_41E2F0()
		sub_41DA70(v3, 9)
		result = 0
	default:
		return 0
	}
	return result
}
func sub_40CD00() int32 {
	var result int32
	if sub_41E2F0() == 2 || sub_41E2F0() == 4 || sub_41E2F0() == 11 || sub_41E2F0() == 5 || sub_41E2F0() == 3 || (func() bool {
		result = 1
		return dword_5d4594_10988 == 1
	}()) {
		result = 0
	}
	return result
}
func sub_40CD50(a1 int32, a2 int32, a3 int32) int32 {
	return 0
}
func nox_xxx_somejoinproc_40CD60(a1 int32, a2 int32, a3 int32) int32 {
	var v4 int32
	if sub_41E2F0() == 7 || sub_41E2F0() == 8 {
		if a2 != 0 {
			if uint32(a2) == 0x4013B {
				nox_xxx_failconn_43B0E0(a3)
				return 0
			}
		} else {
			sub_448640()
			sub_41EFB0(a3)
			v4 = sub_41E2F0()
			sub_41DA70(v4, 10)
		}
	}
	return 0
}
func sub_40CDC0(a1 int32, a2 int32, a3 int32, a4 int32) int32 {
	var (
		v4 *int32
		v5 int32
		v7 int32
	)
	if a2 == 0 {
		v4 = (*int32)(unsafe.Pointer(uintptr(a4)))
		sub_41F6F0(a4)
		if sub_41E2F0() == 7 {
			v5 = sub_41E2F0()
			sub_41DA70(v5, 11)
			return 0
		}
		if sub_41E2F0() == 8 && a4 != 0 {
			for {
				v7 = *v4
				if (v7 & 0x8000) != 0 {
					break
				}
				v4 = (*int32)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(int32(0))*8)))))
				if v4 == nil {
					return 0
				}
			}
			if int32(*(*uint8)(unsafe.Pointer(v4)))&1 != 0 {
				if *(*int32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(int32(0))*5)) != 0 {
					sub_40DB20(*(*int32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(int32(0))*5)))
					*memmap.PtrUint32(6112660, 10980) = 1
				}
				if sub_416650() != 0 {
					sub_40D530(int32(uintptr(unsafe.Pointer(v4))))
				}
			}
		}
	}
	return 0
}
func sub_40CE60() int32 {
	return 0
}
func sub_40D0F0() {
	var (
		v0 *uint8
		v1 int32
		v2 int32
		v3 *uint32
		v4 *uint32
		v5 uint32
	)
	*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v0))), 0)) = *memmap.PtrUint8(6112660, 10976)
	if *memmap.PtrUint32(6112660, 10976) != 0 {
		if dword_5d4594_10956 != 0 {
			sub_40E320(*(***func(uint32, unsafe.Pointer, *int32) int32)(unsafe.Pointer(&dword_5d4594_10956)), int32(uintptr(memmap.PtrOff(0x581450, 5704))), *memmap.PtrInt32(6112660, 10968))
			(cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(**(**uint32)(unsafe.Pointer(&dword_5d4594_10956)) + 8))), (*func(uint32))(nil)))(dword_5d4594_10956)
		}
		if dword_5d4594_4808 != 0 {
			sub_40E320(*(***func(uint32, unsafe.Pointer, *int32) int32)(unsafe.Pointer(&dword_5d4594_4808)), int32(uintptr(memmap.PtrOff(0x581450, 5672))), *memmap.PtrInt32(6112660, 10964))
			(cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(**(**uint32)(unsafe.Pointer(&dword_5d4594_4808)) + 8))), (*func(uint32))(nil)))(dword_5d4594_4808)
		}
		if dword_5d4594_4668 != 0 {
			sub_40E320(*(***func(uint32, unsafe.Pointer, *int32) int32)(unsafe.Pointer(&dword_5d4594_4668)), int32(uintptr(memmap.PtrOff(0x581450, 5640))), *memmap.PtrInt32(6112660, 10960))
			(cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(**(**uint32)(unsafe.Pointer(&dword_5d4594_4668)) + 8))), (*func(uint32))(nil)))(dword_5d4594_4668)
		}
		v0 = (*uint8)(memmap.PtrOff(6112660, 4676))
		if true {
			v1 = int32(*memmap.PtrUint32(6112660, 4692))
			if *memmap.PtrUint32(6112660, 4692) != 0 && **(**uint32)(memmap.PtrOff(6112660, 4692)) != 0 {
				for {
					v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 16))))
					if v2 != 0 {
						(cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2))) + 8))), (*func(uint32))(nil)))(*(*uint32)(unsafe.Pointer(uintptr(v1 + 16))))
					}
					*(*uint32)(unsafe.Pointer(uintptr(v1 + 16))) = 0
					(cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(v1 + 32))), (*func(uint32))(nil)))(0)
					if *memmap.PtrUint32(6112660, 4676) == 100 {
						v1 += 28
					} else {
						v1 += 36
					}
					if *(*uint32)(unsafe.Pointer(uintptr(v1))) == 0 {
						break
					}
				}
			}
			nox_mutex_free((*nox_mutex_t)(memmap.PtrOff(6112660, 4704)))
			nox_mutex_free((*nox_mutex_t)(memmap.PtrOff(6112660, 4728)))
			nox_mutex_free((*nox_mutex_t)(memmap.PtrOff(6112660, 4752)))
			v3 = *(**uint32)(memmap.PtrOff(6112660, 4804))
			if *memmap.PtrUint32(6112660, 4804) != 0 {
				for {
					(cgoAsFunc(*v3, (*func(uint32))(nil)).(func(uint32)))(*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*1)))
					v4 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*2)))))
					operator_delete(unsafe.Pointer(v3))
					v3 = v4
					if v4 == nil {
						break
					}
				}
			}
			*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v0))), 0)) = *memmap.PtrUint8(6112660, 4700)
			if *memmap.PtrUint32(6112660, 4700) != 0 {
				*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v0))), 0)) = *memmap.PtrUint8(6112660, 4784)
				if int32(*memmap.PtrUint8(6112660, 4784)) != 0 {
					if *memmap.PtrUint32(6112660, 4796) != 0 {
						v5 = 0
						for {
							compatHeapDestroy(*(*HANDLE)(unsafe.Pointer(uintptr(*memmap.PtrUint32(6112660, 4796) + func() uint32 {
								p := &v5
								x := *p
								*p++
								return x
							}()*4))))
							if v5 > uint32(*memmap.PtrInt32(6112660, 4792)) {
								break
							}
						}
					}
					*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v0))), 0)) = uint8(int8(compatHeapDestroy(*(*HANDLE)(memmap.PtrOff(6112660, 4700)))))
				}
			}
		}
		*memmap.PtrUint32(6112660, 10976) = 0
	}
}
func sub_40D250() int32 {
	return bool2int(*memmap.PtrUint32(6112660, 10976) != 0 && dword_5d4594_10984 == 0 && (*(*funcuint32(int32))(unsafe.Pointer(uintptr(**(**uint32)(unsafe.Pointer(&dword_5d4594_4668)) + 12))))(dword_5d4594_4668) >= 0)
}
func sub_40D280(a1 int32, a2 int32, a3 int32, a4 int32, a5 int32) int32 {
	var result int32
	if *memmap.PtrUint32(6112660, 10976) != 0 {
		result = bool2int((cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(**(**uint32)(unsafe.Pointer(&dword_5d4594_4668)) + 16))), (*func(uint32, int32, int32, int32, int32, int32) int32)(nil)))(dword_5d4594_4668, a1, a2, a3, a4, a5) >= 0)
	} else {
		result = 0
	}
	return result
}
func sub_40D2C0(a1 int32, a2 int32, a3 int32) int32 {
	var result int32
	if *memmap.PtrUint32(6112660, 10976) != 0 {
		result = bool2int((cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(**(**uint32)(unsafe.Pointer(&dword_5d4594_4668)) + 20))), (*func(uint32, int32, int32, int32) int32)(nil)))(dword_5d4594_4668, a1, a2, a3) >= 0)
	} else {
		result = 0
	}
	return result
}
func sub_40D2F0(a1 int32, a2 int32) int32 {
	var result int32
	if *memmap.PtrUint32(6112660, 10976) != 0 {
		result = bool2int((cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(**(**uint32)(unsafe.Pointer(&dword_5d4594_4668)) + 24))), (*func(uint32, int32, int32) int32)(nil)))(dword_5d4594_4668, a1, a2) >= 0)
	} else {
		result = 0
	}
	return result
}
func sub_40D320(a1 int32) int32 {
	var result int32
	if *memmap.PtrUint32(6112660, 10976) != 0 {
		result = bool2int((cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(**(**uint32)(unsafe.Pointer(&dword_5d4594_4668)) + 28))), (*func(uint32, int32) int32)(nil)))(dword_5d4594_4668, a1) >= 0)
	} else {
		result = 0
	}
	return result
}
func sub_40D350(a1 int32) int32 {
	var result int32
	if *memmap.PtrUint32(6112660, 10976) != 0 {
		result = bool2int((cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(**(**uint32)(unsafe.Pointer(&dword_5d4594_4668)) + 32))), (*func(uint32, int32) int32)(nil)))(dword_5d4594_4668, a1) >= 0)
	} else {
		result = 0
	}
	return result
}
func sub_40D380() {
	if *memmap.PtrUint32(6112660, 10976) != 0 {
		(*(*funcuint32(int32))(unsafe.Pointer(uintptr(**(**uint32)(unsafe.Pointer(&dword_5d4594_4668)) + 36))))(dword_5d4594_4668)
	}
}
func sub_40D3B0() int32 {
	var result int32
	if *memmap.PtrUint32(6112660, 10976) != 0 {
		result = bool2int((*(*funcuint32(int32))(unsafe.Pointer(uintptr(**(**uint32)(unsafe.Pointer(&dword_5d4594_4668)) + 40))))(dword_5d4594_4668) >= 0)
	} else {
		result = 0
	}
	return result
}
func sub_40D3E0(a1 int32) int32 {
	var result int32
	if *memmap.PtrUint32(6112660, 10976) != 0 {
		result = bool2int((cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(**(**uint32)(unsafe.Pointer(&dword_5d4594_4668)) + 44))), (*func(uint32, int32) int32)(nil)))(dword_5d4594_4668, a1) >= 0)
	} else {
		result = 0
	}
	return result
}
func sub_40D410(a1 int32, a2 int32) int32 {
	var result int32
	if *memmap.PtrUint32(6112660, 10976) != 0 {
		result = bool2int((cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(**(**uint32)(unsafe.Pointer(&dword_5d4594_4668)) + 48))), (*func(uint32, int32, int32) int32)(nil)))(dword_5d4594_4668, a1, a2) >= 0)
	} else {
		result = 0
	}
	return result
}
func sub_40D440() int32 {
	var result int32
	if *memmap.PtrUint32(6112660, 10976) == 0 || (*(*funcuint32(int32))(unsafe.Pointer(uintptr(**(**uint32)(unsafe.Pointer(&dword_5d4594_4668)) + 52))))(dword_5d4594_4668) < 0 {
		return 0
	}
	result = 1
	dword_5d4594_10988 = 1
	return result
}
func sub_40D470(a1 int32, a2 int32) int32 {
	var result int32
	if *memmap.PtrUint32(6112660, 10976) != 0 {
		result = bool2int((cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(**(**uint32)(unsafe.Pointer(&dword_5d4594_4668)) + 56))), (*func(uint32, int32, int32) int32)(nil)))(dword_5d4594_4668, a1, a2) >= 0)
	} else {
		result = 0
	}
	return result
}
func sub_40D4A0(a1 int32) int32 {
	var result int32
	if *memmap.PtrUint32(6112660, 10976) != 0 {
		result = bool2int((cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(**(**uint32)(unsafe.Pointer(&dword_5d4594_4668)) + 60))), (*func(uint32, int32) int32)(nil)))(dword_5d4594_4668, a1) >= 0)
	} else {
		result = 0
	}
	return result
}
func sub_40D4D0(a1 int32) int32 {
	var result int32
	if *memmap.PtrUint32(6112660, 10976) != 0 {
		result = bool2int((cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(**(**uint32)(unsafe.Pointer(&dword_5d4594_4668)) + 64))), (*func(uint32, int32) int32)(nil)))(dword_5d4594_4668, a1) >= 0)
	} else {
		result = 0
	}
	return result
}
func sub_40D500(a1 int32, a2 int32) int32 {
	var result int32
	if *memmap.PtrUint32(6112660, 10976) != 0 {
		result = bool2int((cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(**(**uint32)(unsafe.Pointer(&dword_5d4594_4668)) + 68))), (*func(uint32, int32, int32) int32)(nil)))(dword_5d4594_4668, a1, a2) >= 0)
	} else {
		result = 0
	}
	return result
}
func sub_40D530(a1 int32) int32 {
	var result int32
	if *memmap.PtrUint32(6112660, 10976) != 0 {
		result = bool2int((cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(**(**uint32)(unsafe.Pointer(&dword_5d4594_4668)) + 72))), (*func(uint32, int32) int32)(nil)))(dword_5d4594_4668, a1) >= 0)
	} else {
		result = 0
	}
	return result
}
func sub_40D560(a1 int32) int32 {
	var v2 int32
	if *memmap.PtrUint32(6112660, 10976) == 0 {
		return 0
	}
	v2 = (cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(**(**uint32)(unsafe.Pointer(&dword_5d4594_4668)) + 76))), (*func(uint32, int32) int32)(nil)))(dword_5d4594_4668, a1)
	if v2 <= -2147220995 {
		return 0
	}
	if v2 != 0 {
		return 0
	}
	return 1
}
func sub_40D5B0(a1 int32) int32 {
	var result int32
	if *memmap.PtrUint32(6112660, 10976) != 0 {
		result = bool2int((cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(**(**uint32)(unsafe.Pointer(&dword_5d4594_4668)) + 80))), (*func(uint32, int32) int32)(nil)))(dword_5d4594_4668, a1) >= 0)
	} else {
		result = 0
	}
	return result
}
func sub_40D5E0(a1 int32) int32 {
	var result int32
	if *memmap.PtrUint32(6112660, 10976) != 0 {
		result = bool2int((cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(**(**uint32)(unsafe.Pointer(&dword_5d4594_4668)) + 84))), (*func(uint32, int32) int32)(nil)))(dword_5d4594_4668, a1) >= 0)
	} else {
		result = 0
	}
	return result
}
func sub_40D610(a1 int32) int32 {
	var result int32
	if *memmap.PtrUint32(6112660, 10976) != 0 {
		result = bool2int((cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(**(**uint32)(unsafe.Pointer(&dword_5d4594_4668)) + 88))), (*func(uint32, int32) int32)(nil)))(dword_5d4594_4668, a1) >= 0)
	} else {
		result = 0
	}
	return result
}
func sub_40D640(a1 int32) int32 {
	var result int32
	if *memmap.PtrUint32(6112660, 10976) != 0 {
		result = bool2int((cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(**(**uint32)(unsafe.Pointer(&dword_5d4594_4668)) + 188))), (*func(uint32, int32) int32)(nil)))(dword_5d4594_4668, a1) >= 0)
	} else {
		result = 0
	}
	return result
}
func sub_40D670(a1 int32) int32 {
	var result int32
	if *memmap.PtrUint32(6112660, 10976) != 0 {
		result = bool2int((cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(**(**uint32)(unsafe.Pointer(&dword_5d4594_4668)) + 184))), (*func(uint32, int32) int32)(nil)))(dword_5d4594_4668, a1) >= 0)
	} else {
		result = 0
	}
	return result
}
func sub_40D6A0(a1 int32) int32 {
	var result int32
	if *memmap.PtrUint32(6112660, 10976) != 0 {
		result = bool2int((cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(**(**uint32)(unsafe.Pointer(&dword_5d4594_4668)) + 192))), (*func(uint32, int32) int32)(nil)))(dword_5d4594_4668, a1) >= 0)
	} else {
		result = 0
	}
	return result
}
func sub_40D6D0(a1 int32) int32 {
	var result int32
	if *memmap.PtrUint32(6112660, 10976) != 0 {
		result = bool2int((cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(**(**uint32)(unsafe.Pointer(&dword_5d4594_4668)) + 180))), (*func(uint32, int32) int32)(nil)))(dword_5d4594_4668, a1) >= 0)
	} else {
		result = 0
	}
	return result
}
func sub_40D700(a1 int32, a2 int32, a3 int32, a4 int32, a5 int32, a6 int32) int32 {
	var result int32
	if *memmap.PtrUint32(6112660, 10976) != 0 {
		result = bool2int((cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(**(**uint32)(unsafe.Pointer(&dword_5d4594_4668)) + 92))), (*func(uint32, int32, int32, int32, int32, int32, int32) int32)(nil)))(dword_5d4594_4668, a1, a2, a3, a4, a5, a6) >= 0)
	} else {
		result = 0
	}
	return result
}
func sub_40D740(a1 int32) int32 {
	var result int32
	if *memmap.PtrUint32(6112660, 10976) != 0 {
		result = bool2int((cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(**(**uint32)(unsafe.Pointer(&dword_5d4594_4668)) + 96))), (*func(uint32, int32) int32)(nil)))(dword_5d4594_4668, a1) >= 0)
	} else {
		result = 0
	}
	return result
}
func sub_40D770(a1 int32, a2 int32) int32 {
	var result int32
	if *memmap.PtrUint32(6112660, 10976) != 0 {
		result = bool2int((cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(**(**uint32)(unsafe.Pointer(&dword_5d4594_4668)) + 100))), (*func(uint32, int32, int32) int32)(nil)))(dword_5d4594_4668, a1, a2) >= 0)
	} else {
		result = 0
	}
	return result
}
func sub_40D7A0(a1 int32, a2 int32) int32 {
	var result int32
	if *memmap.PtrUint32(6112660, 10976) != 0 {
		result = bool2int((cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(**(**uint32)(unsafe.Pointer(&dword_5d4594_4668)) + 104))), (*func(uint32, int32, int32) int32)(nil)))(dword_5d4594_4668, a1, a2) >= 0)
	} else {
		result = 0
	}
	return result
}
func sub_40D7D0(a1 int32, a2 int32) int32 {
	var result int32
	if *memmap.PtrUint32(6112660, 10976) != 0 {
		result = bool2int((cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(**(**uint32)(unsafe.Pointer(&dword_5d4594_4668)) + 108))), (*func(uint32, int32, int32) int32)(nil)))(dword_5d4594_4668, a1, a2) >= 0)
	} else {
		result = 0
	}
	return result
}
func sub_40D800(a1 int32) int32 {
	var result int32
	if *memmap.PtrUint32(6112660, 10976) != 0 {
		result = bool2int((cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(**(**uint32)(unsafe.Pointer(&dword_5d4594_4668)) + 112))), (*func(uint32, int32) int32)(nil)))(dword_5d4594_4668, a1) >= 0)
	} else {
		result = 0
	}
	return result
}
func sub_40D830(a1 int32) int32 {
	var result int32
	if *memmap.PtrUint32(6112660, 10976) != 0 {
		result = bool2int((cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(**(**uint32)(unsafe.Pointer(&dword_5d4594_4668)) + 116))), (*func(uint32, int32) int32)(nil)))(dword_5d4594_4668, a1) >= 0)
	} else {
		result = 0
	}
	return result
}
func sub_40D860() int32 {
	var result int32
	if *memmap.PtrUint32(6112660, 10976) != 0 {
		result = bool2int((*(*funcuint32(int32))(unsafe.Pointer(uintptr(**(**uint32)(unsafe.Pointer(&dword_5d4594_4668)) + 120))))(dword_5d4594_4668) >= 0)
	} else {
		result = 0
	}
	return result
}
func sub_40D890(a1 int32) int32 {
	var result int32
	if *memmap.PtrUint32(6112660, 10976) != 0 {
		result = bool2int((cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(**(**uint32)(unsafe.Pointer(&dword_5d4594_4668)) + 124))), (*func(uint32, int32) int32)(nil)))(dword_5d4594_4668, a1) >= 0)
	} else {
		result = 0
	}
	return result
}
func sub_40D8C0(a1 int32, a2 int32) int32 {
	var result int32
	if *memmap.PtrUint32(6112660, 10976) != 0 {
		result = bool2int((cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(**(**uint32)(unsafe.Pointer(&dword_5d4594_4668)) + 128))), (*func(uint32, int32, int32) int32)(nil)))(dword_5d4594_4668, a1, a2) >= 0)
	} else {
		result = 0
	}
	return result
}
func sub_40D8F0(a1 int32) int32 {
	var result int32
	if *memmap.PtrUint32(6112660, 10976) != 0 {
		result = bool2int((cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(**(**uint32)(unsafe.Pointer(&dword_5d4594_4668)) + 132))), (*func(uint32, int32) int32)(nil)))(dword_5d4594_4668, a1) >= 0)
	} else {
		result = 0
	}
	return result
}
func sub_40D920(a1 int32) int32 {
	var result int32
	if *memmap.PtrUint32(6112660, 10976) != 0 {
		result = bool2int((cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(**(**uint32)(unsafe.Pointer(&dword_5d4594_4668)) + 136))), (*func(uint32, int32) int32)(nil)))(dword_5d4594_4668, a1) >= 0)
	} else {
		result = 0
	}
	return result
}
func sub_40D950(a1 int32) int32 {
	var result int32
	if *memmap.PtrUint32(6112660, 10976) != 0 {
		result = bool2int((cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(**(**uint32)(unsafe.Pointer(&dword_5d4594_4668)) + 140))), (*func(uint32, int32) int32)(nil)))(dword_5d4594_4668, a1) >= 0)
	} else {
		result = 0
	}
	return result
}
func sub_40D980(a1 int32, a2 int32, a3 int32) int32 {
	var result int32
	if *memmap.PtrUint32(6112660, 10976) != 0 {
		result = bool2int((cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(**(**uint32)(unsafe.Pointer(&dword_5d4594_4668)) + 144))), (*func(uint32, int32, int32, int32) int32)(nil)))(dword_5d4594_4668, a1, a2, a3) >= 0)
	} else {
		result = 0
	}
	return result
}
func sub_40D9C0(a1 int32, a2 int32, a3 int32, a4 int32) int32 {
	var result int32
	if *memmap.PtrUint32(6112660, 10976) != 0 {
		result = bool2int((cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(**(**uint32)(unsafe.Pointer(&dword_5d4594_4668)) + 148))), (*func(uint32, int32, int32, int32, int32) int32)(nil)))(dword_5d4594_4668, a1, a2, a3, a4) >= 0)
	} else {
		result = 0
	}
	return result
}
func sub_40DA00(a1 int32) int32 {
	var result int32
	if *memmap.PtrUint32(6112660, 10976) != 0 {
		result = bool2int((cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(**(**uint32)(unsafe.Pointer(&dword_5d4594_4668)) + 152))), (*func(uint32, int32) int32)(nil)))(dword_5d4594_4668, a1) >= 0)
	} else {
		result = 0
	}
	return result
}
func sub_40DA30(a1 int32) int32 {
	var result int32
	if *memmap.PtrUint32(6112660, 10976) != 0 {
		result = bool2int((cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(**(**uint32)(unsafe.Pointer(&dword_5d4594_4668)) + 156))), (*func(uint32, int32) int32)(nil)))(dword_5d4594_4668, a1) >= 0)
	} else {
		result = 0
	}
	return result
}
func sub_40DA60(a1 int32, a2 int32) int32 {
	var result int32
	if *memmap.PtrUint32(6112660, 10976) != 0 {
		result = bool2int((cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(**(**uint32)(unsafe.Pointer(&dword_5d4594_4668)) + 160))), (*func(uint32, int32, int32) int32)(nil)))(dword_5d4594_4668, a1, a2) >= 0)
	} else {
		result = 0
	}
	return result
}
func sub_40DA90(a1 int32, a2 int32) int32 {
	var result int32
	if *memmap.PtrUint32(6112660, 10976) != 0 {
		result = bool2int((cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(**(**uint32)(unsafe.Pointer(&dword_5d4594_4668)) + 164))), (*func(uint32, int32, int32) int32)(nil)))(dword_5d4594_4668, a1, a2) >= 0)
	} else {
		result = 0
	}
	return result
}
func sub_40DAC0(a1 int32) int32 {
	var result int32
	if *memmap.PtrUint32(6112660, 10976) != 0 {
		result = bool2int((cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(**(**uint32)(unsafe.Pointer(&dword_5d4594_4668)) + 168))), (*func(uint32, int32) int32)(nil)))(dword_5d4594_4668, a1) >= 0)
	} else {
		result = 0
	}
	return result
}
func sub_40DAF0() int32 {
	var result int32
	if *memmap.PtrUint32(6112660, 10976) != 0 {
		result = bool2int((*(*funcuint32(int32))(unsafe.Pointer(uintptr(**(**uint32)(unsafe.Pointer(&dword_5d4594_4668)) + 172))))(dword_5d4594_4668) >= 0)
	} else {
		result = 0
	}
	return result
}
func sub_40DB20(a1 int32) int32 {
	var result int32
	if *memmap.PtrUint32(6112660, 10976) != 0 {
		result = bool2int((cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(**(**uint32)(unsafe.Pointer(&dword_5d4594_4668)) + 176))), (*func(uint32, int32) int32)(nil)))(dword_5d4594_4668, a1) >= 0)
	} else {
		result = 0
	}
	return result
}
func sub_40DB50(a1 int32, a2 int32) int32 {
	var result int32
	if *memmap.PtrUint32(6112660, 10976) != 0 {
		result = bool2int((cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(**(**uint32)(unsafe.Pointer(&dword_5d4594_4668)) + 196))), (*func(uint32, int32, int32) int32)(nil)))(dword_5d4594_4668, a1, a2) >= 0)
	} else {
		result = 0
	}
	return result
}
func sub_40DB80(a1 int32, a2 int32) int32 {
	var result int32
	if *memmap.PtrUint32(6112660, 10976) != 0 {
		result = bool2int((cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(**(**uint32)(unsafe.Pointer(&dword_5d4594_4668)) + 200))), (*func(uint32, int32, int32) int32)(nil)))(dword_5d4594_4668, a1, a2) >= 0)
	} else {
		result = 0
	}
	return result
}
func sub_40DBB0(a1 int32) int32 {
	var v1 int32
	sub_41E560(0, 0, 0, 0)
	v1 = sub_41E2F0()
	sub_41DA70(v1, 16)
	return 0
}
func sub_40DBE0(a1 int32, a2 int32) int32 {
	var v2 int32
	sub_41E560(0, 0, 0, 0)
	sub_41E5C0(a2)
	v2 = sub_41E2F0()
	sub_41DA70(v2, 17)
	return 0
}
func sub_40DC10(a1 int32, a2 int32, a3 int32, a4 int32, a5 int32) int32 {
	var v5 int32
	if sub_40E0B0() != 0 {
		sub_41E560(a2, a3, a4, a5)
		v5 = sub_41E2F0()
		sub_41DA70(v5, 18)
	}
	return 0
}
func sub_40DC50(a1 int32, a2 int32) int32 {
	sub_40E0B0()
	return 0
}
func sub_40DC60(a1 int32) int32 {
	return 0
}
func sub_40DC70() int32 {
	return bool2int((*(*funcuint32(int32))(unsafe.Pointer(uintptr(**(**uint32)(unsafe.Pointer(&dword_5d4594_4808)) + 16))))(dword_5d4594_4808) >= 0)
}
func sub_40DC90(a1 int32, a2 int32, a3 int32, a4 int32, a5 int32, a6 int32) int32 {
	return bool2int((cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(**(**uint32)(unsafe.Pointer(&dword_5d4594_4808)) + 12))), (*func(uint32, int32, int32, int32, int32, int32, int32) int32)(nil)))(dword_5d4594_4808, a1, a2, a3, a4, a5, a6) >= 0)
}
func sub_40DCD0() int32 {
	return bool2int(dword_5d4594_10984 == 0 && (*(*funcuint32(int32))(unsafe.Pointer(uintptr(**(**uint32)(unsafe.Pointer(&dword_5d4594_4808)) + 20))))(dword_5d4594_4808) >= 0)
}
func sub_40DD00(a1 int32, a2 int32, a3 int32, a4 int32, a5 int32) int32 {
	return 0
}
func sub_40DD10(a1 int32, a2 int32, a3 int32, a4 int32, a5 int32, a6 int32) int32 {
	var i int32
	if sub_41E2F0() != 8 {
		return 0
	}
	if a2 == 0 {
		for i = a3; i != 0; i = int32(*(*uint32)(unsafe.Pointer(uintptr(i + 60)))) {
			if int32(*(*uint8)(unsafe.Pointer(uintptr(i + 64)))) != 0 {
				sub_417C00((*byte)(unsafe.Pointer(uintptr(i+64))), int32(*(*uint32)(unsafe.Pointer(uintptr(i + 28)))))
			}
		}
	}
	return 0
}
func sub_40DD60(a1 int32, a2 int32) int32 {
	return 0
}
func sub_40DD70(a1 int32, a2 int32, a3 int32, a4 *byte, a5 *byte) int32 {
	var (
		v5     int32
		result int32
		v7     int32
	)
	if a2 != 0 {
		v5 = sub_41E2F0()
		sub_41DA70(v5, 5)
		result = a3
		dword_5d4594_2660652 = uint32(a2)
		if a3 != 0 {
			nox_swprintf((*libc.WChar)(memmap.PtrOff(0x85B3FC, 10436)), libc.CWString("%S"), a3)
			result = 0
		} else {
			*memmap.PtrUint16(0x85B3FC, 10436) = 0
		}
	} else {
		v7 = sub_41E2F0()
		sub_41DA70(v7, 22)
		sub_41FCA0(a4, a5)
		result = a3
		if a3 != 0 {
			nox_swprintf((*libc.WChar)(memmap.PtrOff(0x85B3FC, 10436)), libc.CWString("%S"), a3)
			result = 0
		} else {
			*memmap.PtrUint16(0x85B3FC, 10436) = 0
		}
	}
	return result
}
func sub_40DE10(a1 int32, a2 int32, a3 int32, a4 int32) int32 {
	var (
		v4 int32
		v6 int32
		v7 int32
		v8 int32
	)
	dword_5d4594_2660652 = uint32(a2)
	if a2 != 0 {
		*memmap.PtrUint16(0x85B3FC, 10436) = 0
		v4 = sub_41E2F0()
		sub_41DA70(v4, 5)
		return 0
	}
	if a3 >= 13 {
		if a3 < 18 {
			v7 = sub_41E2F0()
			sub_41DA70(v7, 21)
			return 0
		}
	} else if a4 == 0 {
		v6 = sub_41E2F0()
		sub_41DA70(v6, 5)
		return 0
	}
	v8 = sub_41E2F0()
	sub_41DA70(v8, 20)
	return 0
}
func sub_40DE90(a1 int32, a2 int32, a3 int32, a4 int32) int32 {
	dword_5d4594_2660652 = uint32(a2)
	return 0
}
func sub_40DEA0(a1 int32, a2 int32, a3 int32, a4 int32) int32 {
	return bool2int((cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(**(**uint32)(unsafe.Pointer(&dword_5d4594_10956)) + 12))), (*func(uint32, int32, int32, int32, int32) int32)(nil)))(dword_5d4594_10956, a1, a2, a3, a4) >= 0)
}
func sub_40DED0(a1 int32, a2 int32, a3 int32, a4 int32, a5 int32, a6 int32, a7 int32, a8 int32, a9 int32) int32 {
	return bool2int((cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(**(**uint32)(unsafe.Pointer(&dword_5d4594_10956)) + 16))), (*func(uint32, int32, int32, int32, int32, int32, int32, int32, int32, int32) int32)(nil)))(dword_5d4594_10956, a1, a2, a3, a4, a5, a6, a7, a8, a9) >= 0)
}
func sub_40DF20(a1 int32, a2 int32, a3 int32, a4 int32, a5 int32, a6 int32, a7 int32) int32 {
	return bool2int((cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(**(**uint32)(unsafe.Pointer(&dword_5d4594_10956)) + 20))), (*func(uint32, int32, int32, int32, int32, int32, int32, int32) int32)(nil)))(dword_5d4594_10956, a1, a2, a3, a4, a5, a6, a7) >= 0)
}
func sub_40DF60(a1 int32, a2 int32, a3 int32) int32 {
	return bool2int((cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(**(**uint32)(unsafe.Pointer(&dword_5d4594_10956)) + 24))), (*func(uint32, int32, int32, int32) int32)(nil)))(dword_5d4594_10956, a1, a2, a3) >= 0)
}
func sub_40DF90() int32 {
	return bool2int(dword_5d4594_10984 == 0 && (*(*funcuint32(int32))(unsafe.Pointer(uintptr(**(**uint32)(unsafe.Pointer(&dword_5d4594_10956)) + 28))))(dword_5d4594_10956) >= 0)
}
func sub_40DFC0(a1 int32, a2 int32) int32 {
	return bool2int((cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(**(**uint32)(unsafe.Pointer(&dword_5d4594_10956)) + 32))), (*func(uint32, int32, int32) int32)(nil)))(dword_5d4594_10956, a1, a2) >= 0)
}
func sub_40DFE0(a1 int32, a2 int32, a3 int32, a4 int32, a5 int32, a6 int32) int32 {
	return bool2int((cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(**(**uint32)(unsafe.Pointer(&dword_5d4594_10956)) + 36))), (*func(uint32, int32, int32, int32, int32, int32, int32) int32)(nil)))(dword_5d4594_10956, a1, a2, a3, a4, a5, a6) >= 0)
}
func sub_40E020(a1 int32, a2 int32, a3 int32, a4 int32) int32 {
	return bool2int((cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(**(**uint32)(unsafe.Pointer(&dword_5d4594_10956)) + 40))), (*func(uint32, int32, int32, int32, int32) int32)(nil)))(dword_5d4594_10956, a1, a2, a3, a4) >= 0)
}
func sub_40E050(a1 int32, a2 int32) int32 {
	return bool2int((cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(**(**uint32)(unsafe.Pointer(&dword_5d4594_10956)) + 44))), (*func(uint32, int32, int32) int32)(nil)))(dword_5d4594_10956, a1, a2) >= 0)
}
func sub_40E070() {
}
func sub_40E090() {
	dword_5d4594_10984 = 0
}
func sub_40E0A0() {
	dword_5d4594_10984 = 1
}
func sub_40E0B0() int32 {
	return bool2int(dword_5d4594_10984 == 0)
}
func sub_40E0C0() int32 {
	return int32(*memmap.PtrUint32(6112660, 10976))
}
func nox_common_getInstallPath_40E0D0(dst *byte, lpSubKey *byte, a3 int32) int32 {
	var (
		result    int32
		i         int32
		v5        *byte
		v6        int32
		v7        int32
		v8        *File
		v9        int32
		v10       int32
		v11       int32
		phkResult HKEY
		v13       *File
		cbData    uint32
		Type      uint32
		Data      [516]byte
	)
	if compatRegOpenKeyExA((HKEY)(unsafe.Pointer(uintptr(1))), lpSubKey, 0, 0x20019, &phkResult) != 0 {
		return 0
	}
	Data[0] = 0
	cbData = 513
	if compatRegQueryValueExA(phkResult, CString("InstallPath"), nil, &Type, (*uint8)(unsafe.Pointer(&Data[0])), &cbData) != 0 {
		compatRegCloseKey(phkResult)
		result = 0
	} else {
		compatRegCloseKey(phkResult)
		for i = int32(libc.StrLen(&Data[0]) - 1); i >= 0; Data[func() int32 {
			p := &i
			x := *p
			*p--
			return x
		}()] = 0 {
			if Data[i] == 92 {
				break
			}
		}
		v5 = &Data[libc.StrLen(&Data[0])+1]
		v6 = int32(*memmap.PtrUint32(0x587000, 25812))
		v7 = int32(*memmap.PtrUint32(0x587000, 25816))
		*(*uint32)(unsafe.Pointer(func() *byte {
			p := &v5
			*p = (*byte)(unsafe.Add(unsafe.Pointer(*p), -1))
			return *p
		}())) = *memmap.PtrUint32(0x587000, 25808)
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v5))), unsafe.Sizeof(uint32(0))*1))) = uint32(v6)
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v5))), unsafe.Sizeof(uint32(0))*2))) = uint32(v7)
		result = int32(uintptr(unsafe.Pointer(nox_fs_open_text(&Data[0]))))
		v8 = (*File)(unsafe.Pointer(uintptr(result)))
		v13 = (*File)(unsafe.Pointer(uintptr(result)))
		if result != 0 {
			v9 = 0
			v10 = 1
			if a3 == 0 {
				v10 = -1
			}
			v11 = nox_fs_fgetc((*File)(unsafe.Pointer(uintptr(result))))
			if v11 != -1 {
				for {
					*(*uint8)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(dst), v9)))) = uint8(int8((int32((*((*byte)(unsafe.Add(unsafe.Pointer(dst), v9)))-48)%10)+v10*v11+1000)%10 + 48))
					if func() int32 {
						p := &v9
						*p++
						return *p
					}() == int32(libc.StrLen(dst)) {
						v9 = 0
					}
					v11 = nox_fs_fgetc(v13)
					if v11 == -1 {
						break
					}
				}
				v8 = v13
			}
			nox_fs_close(v8)
			result = 1
		}
	}
	return result
}
func sub_40E260(a1 **func(uint32, unsafe.Pointer, *int32) int32, a2 int32, a3 int32, a4 int32) int32 {
	var (
		v4 *func(uint32, unsafe.Pointer, *int32) int32
		v5 int32
		v7 int32
		v8 [3]int32
		v9 int32
	)
	v8[0] = 0
	v9 = 0
	v7 = 0
	v4 = *a1
	*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v9))), 0)) = 1
	v5 = (*v4)(uint32(uintptr(unsafe.Pointer(a1))), memmap.PtrOff(0x581450, 11152), &v8[0])
	if v5 >= 0 {
		v5 = (cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v8[0]))) + 16))), (*func(int32, int32, *int32) int32)(nil)))(v8[0], a3, &v7)
		if v5 >= 0 {
			v5 = (cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v7))) + 20))), (*func(int32, int32, int32) int32)(nil)))(v7, a2, a4)
		}
	}
	*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v9))), 0)) = 0
	if v7 != 0 {
		(cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v7))) + 8))), (*func(int32))(nil)))(v7)
	}
	v9 = -1
	if v8[0] != 0 {
		(cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v8[0]))) + 8))), (*func(int32))(nil)))(v8[0])
	}
	return v5
}
func sub_40E320(a1 **func(uint32, unsafe.Pointer, *int32) int32, a2 int32, a3 int32) int32 {
	var (
		v3 *func(uint32, unsafe.Pointer, *int32) int32
		v4 int32
		v6 int32
		v7 [3]int32
		v8 int32
	)
	v7[0] = 0
	v8 = 0
	v6 = 0
	v3 = *a1
	*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v8))), 0)) = 1
	v4 = (*v3)(uint32(uintptr(unsafe.Pointer(a1))), memmap.PtrOff(0x581450, 11152), &v7[0])
	if v4 >= 0 {
		v4 = (cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v7[0]))) + 16))), (*func(int32, int32, *int32) int32)(nil)))(v7[0], a2, &v6)
		if v4 >= 0 {
			v4 = (cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v6))) + 24))), (*func(int32, int32) int32)(nil)))(v6, a3)
		}
	}
	*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v8))), 0)) = 0
	if v6 != 0 {
		(cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v6))) + 8))), (*func(int32))(nil)))(v6)
	}
	v8 = -1
	if v7[0] != 0 {
		(cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v7[0]))) + 8))), (*func(int32))(nil)))(v7[0])
	}
	return v4
}
func sub_40E3D0(this *int32) int32 {
	var result int32
	result = *this
	if *this != 0 {
		result = (cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(result))) + 8))), (*func(int32) int32)(nil)))(result)
	}
	return result
}
func sub_40E3E0(a1 *uint32) int32 {
	var (
		v1     unsafe.Pointer
		v2     *uint32
		result int32
	)
	v1 = operator_new(32)
	v2 = (*uint32)(v1)
	if v1 != nil {
		sub_40E470((*byte)(unsafe.Add(unsafe.Pointer((*byte)(v1)), 4)))
		*v2 = uint32(uintptr(memmap.PtrOff(0x581450, 5920)))
		compatInterlockedIncrement((*int32)(memmap.PtrOff(6112660, 4696)))
		result = 0
		*a1 = uint32(uintptr(unsafe.Pointer(v2)))
	} else {
		*a1 = 0
		result = -2147024882
	}
	return result
}
func sub_40E470(this *byte) *byte {
	var v1 *byte
	v1 = this
	*(*uint32)(unsafe.Pointer(this)) = 0
	nox_mutex_init((*nox_mutex_t)(unsafe.Add(unsafe.Pointer(this), 4)))
	return v1
}
func sub_40E490(a1 int32) int32 {
	return compatInterlockedIncrement((*int32)(unsafe.Pointer(uintptr(a1 + 4))))
}
func sub_40E4B0(lpMem *int32) int32 {
	var v1 int32
	v1 = compatInterlockedDecrement((*int32)(unsafe.Add(unsafe.Pointer(lpMem), unsafe.Sizeof(int32(0))*1)))
	if v1 == 0 && lpMem != nil {
		sub_40E5F0((*uint32)(unsafe.Pointer(lpMem)))
		operator_delete(unsafe.Pointer(lpMem))
	}
	return v1
}
func sub_40E4F0(a1 int32, a2 *uint32, a3 *int32) int32 {
	var (
		v3     *uint8
		result int32
		v5     *uint32
		v6     int32
		v7     func(int32, *uint32, *int32, uint32) int32
		v8     int32
		v9     int32
	)
	v3 = (*uint8)(memmap.PtrOff(0x581450, 5896))
	if a3 == nil {
		return -2147467261
	}
	*a3 = 0
	if *a2 == 0 && *(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*1)) == 0 && *(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*2)) == 192 && *(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*3)) == 0x46000000 {
		(cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1))) + 4))), (*func(int32))(nil)))(a1)
		*a3 = a1
		return 0
	}
	for {
		v5 = *(**uint32)(unsafe.Pointer(v3))
		v6 = bool2int(*(*uint32)(unsafe.Pointer(v3)) == 0)
		if *(*uint32)(unsafe.Pointer(v3)) == 0 || *v5 == *a2 && *(*uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(uint32(0))*1)) == *(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*1)) && *(*uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(uint32(0))*2)) == *(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*2)) && *(*uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(uint32(0))*3)) == *(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*3)) {
			break
		}
	LABEL_17:
		v8 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v3))), unsafe.Sizeof(uint32(0))*5))))
		v3 = (*uint8)(unsafe.Add(unsafe.Pointer(v3), 12))
		if v8 == 0 {
			return -2147467262
		}
	}
	v7 = cgoAsFunc(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v3))), unsafe.Sizeof(uint32(0))*2))), (*func(int32, *uint32, *int32, uint32) int32)(nil)).(func(int32, *uint32, *int32, uint32) int32)
	if int32(cgoFuncAddr(v7)) != 1 {
		result = v7(a1, a2, a3, *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v3))), unsafe.Sizeof(uint32(0))*1))))
		if result == 0 || v6 == 0 && result < 0 {
			return result
		}
		goto LABEL_17
	}
	v9 = int32(uint32(a1) + *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v3))), unsafe.Sizeof(uint32(0))*1))))
	(cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v9))) + 4))), (*func(int32))(nil)))(v9)
	*a3 = v9
	return 0
}
func sub_40E5F0(a1 *uint32) {
	var (
		v1 *uint32
		v2 *uint32
	)
	v1 = a1
	v2 = (*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*1))
	*a1 = uint32(uintptr(memmap.PtrOff(0x581450, 5920)))
	*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*1)) = 1
	compatInterlockedDecrement((*int32)(memmap.PtrOff(6112660, 4696)))
	nox_mutex_free((*nox_mutex_t)(unsafe.Pointer(func() *uint32 {
		if v1 != nil {
			return (*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*1))
		}
		return (*uint32)(unsafe.Pointer(uintptr(4)))
	}())))
}
func sub_40E630(a1 *uint32) int32 {
	var (
		v1     unsafe.Pointer
		v2     *uint32
		result int32
	)
	v1 = operator_new(32)
	v2 = (*uint32)(v1)
	if v1 != nil {
		sub_40E470((*byte)(unsafe.Add(unsafe.Pointer((*byte)(v1)), 4)))
		*v2 = uint32(uintptr(memmap.PtrOff(0x581450, 6084)))
		compatInterlockedIncrement((*int32)(memmap.PtrOff(6112660, 4696)))
		result = 0
		*a1 = uint32(uintptr(unsafe.Pointer(v2)))
	} else {
		*a1 = 0
		result = -2147024882
	}
	return result
}
func sub_40E680(lpMem *int32) int32 {
	var v1 int32
	v1 = compatInterlockedDecrement((*int32)(unsafe.Add(unsafe.Pointer(lpMem), unsafe.Sizeof(int32(0))*1)))
	if v1 == 0 && lpMem != nil {
		sub_40E7C0((*uint32)(unsafe.Pointer(lpMem)))
		operator_delete(unsafe.Pointer(lpMem))
	}
	return v1
}
func sub_40E6C0(a1 int32, a2 *uint32, a3 *int32) int32 {
	var (
		v3     *uint8
		result int32
		v5     *uint32
		v6     int32
		v7     func(int32, *uint32, *int32, uint32) int32
		v8     int32
		v9     int32
	)
	v3 = (*uint8)(memmap.PtrOff(0x581450, 5872))
	if a3 == nil {
		return -2147467261
	}
	*a3 = 0
	if *a2 == 0 && *(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*1)) == 0 && *(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*2)) == 192 && *(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*3)) == 0x46000000 {
		(cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1))) + 4))), (*func(int32))(nil)))(a1)
		*a3 = a1
		return 0
	}
	for {
		v5 = *(**uint32)(unsafe.Pointer(v3))
		v6 = bool2int(*(*uint32)(unsafe.Pointer(v3)) == 0)
		if *(*uint32)(unsafe.Pointer(v3)) == 0 || *v5 == *a2 && *(*uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(uint32(0))*1)) == *(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*1)) && *(*uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(uint32(0))*2)) == *(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*2)) && *(*uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(uint32(0))*3)) == *(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*3)) {
			break
		}
	LABEL_17:
		v8 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v3))), unsafe.Sizeof(uint32(0))*5))))
		v3 = (*uint8)(unsafe.Add(unsafe.Pointer(v3), 12))
		if v8 == 0 {
			return -2147467262
		}
	}
	v7 = cgoAsFunc(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v3))), unsafe.Sizeof(uint32(0))*2))), (*func(int32, *uint32, *int32, uint32) int32)(nil)).(func(int32, *uint32, *int32, uint32) int32)
	if int32(cgoFuncAddr(v7)) != 1 {
		result = v7(a1, a2, a3, *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v3))), unsafe.Sizeof(uint32(0))*1))))
		if result == 0 || v6 == 0 && result < 0 {
			return result
		}
		goto LABEL_17
	}
	v9 = int32(uint32(a1) + *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v3))), unsafe.Sizeof(uint32(0))*1))))
	(cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v9))) + 4))), (*func(int32))(nil)))(v9)
	*a3 = v9
	return 0
}
func sub_40E7C0(a1 *uint32) {
	var (
		v1 *uint32
		v2 *uint32
	)
	v1 = a1
	v2 = (*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*1))
	*a1 = uint32(uintptr(memmap.PtrOff(0x581450, 6084)))
	*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*1)) = 1
	compatInterlockedDecrement((*int32)(memmap.PtrOff(6112660, 4696)))
	nox_mutex_free((*nox_mutex_t)(unsafe.Pointer(func() *uint32 {
		if v1 != nil {
			return (*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*1))
		}
		return (*uint32)(unsafe.Pointer(uintptr(4)))
	}())))
}
func sub_40E800(a1 *uint32) int32 {
	var (
		v1     unsafe.Pointer
		v2     *uint32
		result int32
	)
	v1 = operator_new(32)
	v2 = (*uint32)(v1)
	if v1 != nil {
		sub_40E470((*byte)(unsafe.Add(unsafe.Pointer((*byte)(v1)), 4)))
		*v2 = uint32(uintptr(memmap.PtrOff(0x581450, 6116)))
		compatInterlockedIncrement((*int32)(memmap.PtrOff(6112660, 4696)))
		result = 0
		*a1 = uint32(uintptr(unsafe.Pointer(v2)))
	} else {
		*a1 = 0
		result = -2147024882
	}
	return result
}
func sub_40E850(lpMem *int32) int32 {
	var v1 int32
	v1 = compatInterlockedDecrement((*int32)(unsafe.Add(unsafe.Pointer(lpMem), unsafe.Sizeof(int32(0))*1)))
	if v1 == 0 && lpMem != nil {
		sub_40E990((*uint32)(unsafe.Pointer(lpMem)))
		operator_delete(unsafe.Pointer(lpMem))
	}
	return v1
}
func sub_40E890(a1 int32, a2 *uint32, a3 *int32) int32 {
	var (
		v3     *uint8
		result int32
		v5     *uint32
		v6     int32
		v7     func(int32, *uint32, *int32, uint32) int32
		v8     int32
		v9     int32
	)
	v3 = (*uint8)(memmap.PtrOff(0x581450, 5848))
	if a3 == nil {
		return -2147467261
	}
	*a3 = 0
	if *a2 == 0 && *(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*1)) == 0 && *(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*2)) == 192 && *(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*3)) == 0x46000000 {
		(cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1))) + 4))), (*func(int32))(nil)))(a1)
		*a3 = a1
		return 0
	}
	for {
		v5 = *(**uint32)(unsafe.Pointer(v3))
		v6 = bool2int(*(*uint32)(unsafe.Pointer(v3)) == 0)
		if *(*uint32)(unsafe.Pointer(v3)) == 0 || *v5 == *a2 && *(*uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(uint32(0))*1)) == *(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*1)) && *(*uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(uint32(0))*2)) == *(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*2)) && *(*uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(uint32(0))*3)) == *(*uint32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(uint32(0))*3)) {
			break
		}
	LABEL_17:
		v8 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v3))), unsafe.Sizeof(uint32(0))*5))))
		v3 = (*uint8)(unsafe.Add(unsafe.Pointer(v3), 12))
		if v8 == 0 {
			return -2147467262
		}
	}
	v7 = cgoAsFunc(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v3))), unsafe.Sizeof(uint32(0))*2))), (*func(int32, *uint32, *int32, uint32) int32)(nil)).(func(int32, *uint32, *int32, uint32) int32)
	if int32(cgoFuncAddr(v7)) != 1 {
		result = v7(a1, a2, a3, *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v3))), unsafe.Sizeof(uint32(0))*1))))
		if result == 0 || v6 == 0 && result < 0 {
			return result
		}
		goto LABEL_17
	}
	v9 = int32(uint32(a1) + *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v3))), unsafe.Sizeof(uint32(0))*1))))
	(cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v9))) + 4))), (*func(int32))(nil)))(v9)
	*a3 = v9
	return 0
}
func sub_40E990(a1 *uint32) {
	var (
		v1 *uint32
		v2 *uint32
	)
	v1 = a1
	v2 = (*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*1))
	*a1 = uint32(uintptr(memmap.PtrOff(0x581450, 6116)))
	*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*1)) = 1
	compatInterlockedDecrement((*int32)(memmap.PtrOff(6112660, 4696)))
	nox_mutex_free((*nox_mutex_t)(unsafe.Pointer(func() *uint32 {
		if v1 != nil {
			return (*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*1))
		}
		return (*uint32)(unsafe.Pointer(uintptr(4)))
	}())))
}
func nox_xxx_netSendBySock_40EE10(a1 uint32, a2 int32, a3 uint32) int32 {
	var v3 *uint8
	v3 = nox_netlist_copyPacketList_40ED60(a2, int32(a3), &a3)
	if v3 != nil {
		nox_xxx_netSendSock_552640(a1, (*byte)(unsafe.Pointer(v3)), int32(a3), 0)
		nox_xxx_netSendReadPacket_5528B0(a1, 1)
	}
	return 1
}
func nox_xxx_cliCanTalkMB_4100F0(a1 *int16) int32 {
	var (
		v1 int16
		v2 *int16
	)
	v1 = *a1
	v2 = (*int16)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int16(0))*1))
	if int32(*a1) == 0 {
		return 1
	}
	for int32(*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v1))), unsafe.Sizeof(int16(0))-1))) == 0 {
		v1 = *v2
		v2 = (*int16)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(int16(0))*1))
		if int32(v1) == 0 {
			return 1
		}
	}
	return 0
}
func sub_410120() int32 {
	var (
		v0     *byte
		result int32
	)
	v0 = *(**byte)(memmap.PtrOff(0x587000, uint32(strMan.Lang()*16+25900)))
	if v0 != nil {
		result = bool2int(libc.StrCaseCmp(v0, CString("e")) == 0)
	} else {
		result = 1
	}
	return result
}
func nox_xxx_wall_410160() int32 {
	for i := int32(0); i < 8192; i++ {
		var v1 int32 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(dword_5D4594_251544), unsafe.Sizeof(uint32(0))*uintptr(i))))
		if v1 != 0 {
			var (
				v2 int32 = int32(dword_5d4594_251548)
				v3 int32 = 0
			)
			for {
				v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 16))))
				*(*uint32)(unsafe.Pointer(uintptr(v1 + 20))) = uint32(v2)
				v2 = v1
				v1 = v3
				dword_5d4594_251548 = uint32(v2)
				if v3 == 0 {
					break
				}
			}
		}
		*(*uint32)(unsafe.Add(unsafe.Pointer(dword_5D4594_251544), unsafe.Sizeof(uint32(0))*uintptr(i))) = 0
	}
	dword_5d4594_251552 = 0
	var j int32 = 0
	for j = 0; j < 1024; j += 4 {
		*(*uint32)(unsafe.Pointer(uintptr(uint32(j) + dword_5d4594_251556))) = 0
	}
	return j
}
func nox_xxx_mapAlloc_4101D0() int32 {
	dword_5D4594_251544 = (*uint32)(alloc.Calloc(8192, 4))
	if dword_5D4594_251544 == nil {
		return 0
	}
	dword_5d4594_251556 = uint32(uintptr(alloc.Calloc(256, 4)))
	if dword_5d4594_251556 == 0 {
		return 0
	}
	dword_5d4594_251552 = 0
	var v1 int32 = 0
	for {
		var v2 *uint32 = (*uint32)(alloc.Calloc(1, 36))
		if v2 == nil {
			break
		}
		v1++
		*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*5)) = dword_5d4594_251548
		dword_5d4594_251548 = uint32(uintptr(unsafe.Pointer(v2)))
		if v1 >= 8192 {
			nox_xxx_wall_410160()
			return 1
		}
	}
	return 0
}
func nox_xxx_wallCreateAt_410250(a1 int32, a2 int32) unsafe.Pointer {
	var (
		result unsafe.Pointer
		v3     int32
		v4     int32
		v5     int32
		v6     int32
	)
	if a1 < 0 || a1 >= 256 || a2 < 0 || a2 >= 256 {
		return nil
	}
	result = nox_server_getWallAtGrid_410580(a1, a2)
	if result != nil {
		return result
	}
	v3 = int32(dword_5d4594_251548)
	if dword_5d4594_251548 == 0 {
		return nil
	}
	dword_5d4594_251548 = *(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_251548 + 20)))
	libc.MemSet(unsafe.Pointer(uintptr(v3)), 0, 36)
	*(*uint8)(unsafe.Pointer(uintptr(v3 + 6))) = uint8(int8(a2))
	*(*uint8)(unsafe.Pointer(uintptr(v3 + 5))) = uint8(int8(a1))
	v4 = (int32(uint16(int16(a2))) + (int32(uint16(int16(a1))) << 8)) & 8191
	*(*uint32)(unsafe.Pointer(uintptr(v3 + 16))) = *(*uint32)(unsafe.Add(unsafe.Pointer(dword_5D4594_251544), unsafe.Sizeof(uint32(0))*uintptr(v4)))
	*(*uint32)(unsafe.Add(unsafe.Pointer(dword_5D4594_251544), unsafe.Sizeof(uint32(0))*uintptr(v4))) = uint32(v3)
	*(*uint32)(unsafe.Pointer(uintptr(v3 + 20))) = dword_5d4594_251552
	dword_5d4594_251552 = uint32(v3)
	v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_251556 + uint32(a2*4)))))
	if v5 != 0 {
		v6 = 0
		for int32(*(*uint8)(unsafe.Pointer(uintptr(v3 + 5)))) >= int32(*(*uint8)(unsafe.Pointer(uintptr(v5 + 5)))) {
			v6 = v5
			v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(v5 + 24))))
			if v5 == 0 {
				*(*uint32)(unsafe.Pointer(uintptr(v6 + 24))) = uint32(v3)
				*(*uint32)(unsafe.Pointer(uintptr(v3 + 24))) = 0
				return unsafe.Pointer(uintptr(v3))
			}
		}
		if v6 != 0 {
			*(*uint32)(unsafe.Pointer(uintptr(v6 + 24))) = uint32(v3)
		} else {
			*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_251556 + uint32(a2*4)))) = uint32(v3)
		}
		*(*uint32)(unsafe.Pointer(uintptr(v3 + 24))) = uint32(v5)
		result = unsafe.Pointer(uintptr(v3))
	} else {
		*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_251556 + uint32(a2*4)))) = uint32(v3)
		*(*uint32)(unsafe.Pointer(uintptr(v3 + 24))) = 0
		result = unsafe.Pointer(uintptr(v3))
	}
	return result
}
func nox_xxx_doorAttachWall_410360(a1 int32, a2 int32, a3 int32) *uint8 {
	var (
		result *uint8
		v4     int8
	)
	result = (*uint8)(nox_xxx_wallCreateAt_410250(a2, a3))
	if result != nil {
		v4 = int8(int32(*(*uint8)(unsafe.Add(unsafe.Pointer(result), 4))) | 16)
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(result))), unsafe.Sizeof(uint32(0))*7))) = uint32(a1)
		*(*uint8)(unsafe.Add(unsafe.Pointer(result), 4)) = uint8(v4)
	}
	return result
}
func sub_410390(a1 int32, a2 int32, a3 int32) *uint32 {
	var (
		v3     *uint32
		result *uint32
		v5     int32
		v6     int8
		v7     [2]int32
	)
	v3 = (*uint32)(unsafe.Pointer(uintptr(nox_xxx_wall_4105E0(a2, a3))))
	if v3 != nil || (func() bool {
		result = (*uint32)(nox_xxx_wallCreateAt_410250(a2, a3))
		return (func() *uint32 {
			v3 = result
			return v3
		}()) != nil
	}()) {
		v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 16))))
		v7[0] = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 12))))
		v6 = int8(*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(v3))), 4))))
		*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*8)) = uint32(a1)
		v7[1] = v5
		*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(v3))), 4))) = uint8(int8(int32(v6) | 16))
		result = (*uint32)(unsafe.Pointer(nox_xxx_polygonIsPlayerInPolygon_4217B0((*int2)(unsafe.Pointer(&v7[0])), 0)))
		if result != nil || (func() *uint32 {
			result = (*uint32)(unsafe.Pointer(sub_421990((*int2)(unsafe.Pointer(&v7[0])), 10.0, 0)))
			return result
		}()) != nil {
			*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(v3))), 8))) = *((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(result))), 130)))
		} else {
			*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(v3))), 8))) = 1
		}
	}
	return result
}
func nox_xxx_mapDelWallAtPt_410430(a1 int32, a2 int32) *uint32 {
	var (
		v2     int32
		v3     *uint32
		result *uint32
		v5     *uint32
		v6     *uint32
		v7     int32
		v8     int32
	)
	v2 = (int32(uint16(int16(a2))) + (int32(uint16(int16(a1))) << 8)) & 8191
	v3 = nil
	result = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(dword_5D4594_251544), unsafe.Sizeof(uint32(0))*uintptr(v2))))))
	if result != nil {
		for int32(*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(result))), 5)))) != a1 || int32(*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(result))), 6)))) != a2 {
			v3 = result
			result = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*4)))))
			if result == nil {
				return result
			}
		}
		if v3 != nil {
			*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*4)) = *(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*4))
		} else {
			*(*uint32)(unsafe.Add(unsafe.Pointer(dword_5D4594_251544), unsafe.Sizeof(uint32(0))*uintptr(v2))) = *(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*4))
		}
		v5 = *(**uint32)(unsafe.Pointer(&dword_5d4594_251552))
		v6 = nil
		if dword_5d4594_251552 == 0 {
			goto LABEL_24
		}
		for {
			if v5 == result {
				break
			}
			v6 = v5
			v5 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(uint32(0))*5)))))
			if v5 == nil {
				break
			}
		}
		if v6 != nil {
			*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*5)) = *(*uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(uint32(0))*5))
		} else {
		LABEL_24:
			dword_5d4594_251552 = *(*uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(uint32(0))*5))
		}
		v7 = 0
		v8 = int32(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_251556 + uint32(a2*4)))))
		if v8 == 0 {
			goto LABEL_25
		}
		for {
			if int32(*(*uint8)(unsafe.Pointer(uintptr(v8 + 5)))) == a1 && int32(*(*uint8)(unsafe.Pointer(uintptr(v8 + 6)))) == a2 {
				break
			}
			v7 = v8
			v8 = int32(*(*uint32)(unsafe.Pointer(uintptr(v8 + 24))))
			if v8 == 0 {
				break
			}
		}
		if v7 != 0 {
			*(*uint32)(unsafe.Pointer(uintptr(v7 + 24))) = *(*uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(uint32(0))*6))
			result = *(**uint32)(unsafe.Pointer(&dword_5d4594_251548))
			*(*uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(uint32(0))*5)) = dword_5d4594_251548
			dword_5d4594_251548 = uint32(uintptr(unsafe.Pointer(v5)))
		} else {
		LABEL_25:
			*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_251556 + uint32(a2*4)))) = *(*uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(uint32(0))*6))
			result = *(**uint32)(unsafe.Pointer(&dword_5d4594_251548))
			*(*uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(uint32(0))*5)) = dword_5d4594_251548
			dword_5d4594_251548 = uint32(uintptr(unsafe.Pointer(v5)))
		}
	}
	return result
}
func nox_xxx_wallDestroyedByWallid_410520(a1 int16) *int32 {
	var (
		result *int32
		v2     int32
	)
	result = (*int32)(nox_xxx_wallGetFirstBreakableCli_410870())
	if result != nil {
		for {
			v2 = *(*int32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(int32(0))*1))
			if int32(*(*uint16)(unsafe.Pointer(uintptr(v2 + 10)))) == int32(a1) {
				break
			}
			result = (*int32)(unsafe.Pointer(uintptr(nox_xxx_wallGetNextBreakableCli_410880(result))))
			if result == nil {
				return result
			}
		}
		*(*uint8)(unsafe.Pointer(uintptr(v2 + 4))) |= 32
	}
	return result
}
func sub_410550(a1 int16) int32 {
	var v1 *int32
	v1 = (*int32)(nox_xxx_wallSecretGetFirstWall_410780())
	if v1 == nil {
		return 0
	}
	for int32(*(*uint16)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(int32(0))*3)) + 10)))) != int32(a1) {
		v1 = (*int32)(unsafe.Pointer(uintptr(nox_xxx_wallSecretNext_410790(v1))))
		if v1 == nil {
			return 0
		}
	}
	return *(*int32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(int32(0))*3))
}
func nox_server_getWallAtGrid_410580(a1 int32, a2 int32) unsafe.Pointer {
	var (
		result int32
		v3     int32
	)
	if (int32(uint8(int8(a1)))+int32(uint8(int8(a2))))&1 != 0 {
		return nil
	}
	v3 = (int32(uint16(int16(a2))) + (int32(uint16(int16(a1))) << 8)) & 8191
	*memmap.PtrUint32(6112660, 251528) = uint32(v3)
	result = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(dword_5D4594_251544), unsafe.Sizeof(uint32(0))*uintptr(v3))))
	for *memmap.PtrUint32(6112660, 251524) = uint32(result); result != 0; *memmap.PtrUint32(6112660, 251524) = uint32(result) {
		if int32(*(*uint8)(unsafe.Pointer(uintptr(result + 5)))) == a1 && int32(*(*uint8)(unsafe.Pointer(uintptr(result + 6)))) == a2 && (int32(*(*uint8)(unsafe.Pointer(uintptr(result + 4))))&48) == 0 {
			break
		}
		result = int32(*(*uint32)(unsafe.Pointer(uintptr(result + 16))))
	}
	return unsafe.Pointer(uintptr(result))
}
func nox_xxx_wall_4105E0(a1 int32, a2 int32) int32 {
	var (
		v2     int32
		result int32
	)
	v2 = (int32(uint16(int16(a2))) + (int32(uint16(int16(a1))) << 8)) & 8191
	*memmap.PtrUint32(6112660, 251536) = uint32(v2)
	result = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(dword_5D4594_251544), unsafe.Sizeof(uint32(0))*uintptr(v2))))
	*memmap.PtrUint32(6112660, 251532) = uint32(result)
	if result == 0 {
		return 0
	}
	for int32(*(*uint8)(unsafe.Pointer(uintptr(result + 5)))) != a1 || int32(*(*uint8)(unsafe.Pointer(uintptr(result + 6)))) != a2 || int32(*(*uint8)(unsafe.Pointer(uintptr(result + 4))))&32 != 0 {
		result = int32(*(*uint32)(unsafe.Pointer(uintptr(result + 16))))
		*memmap.PtrUint32(6112660, 251532) = uint32(result)
		if result == 0 {
			return 0
		}
	}
	return result
}
func nox_xxx_wallForeachFn_410640(a1 func(int32, int32), a2 int32) int32 {
	var (
		result int32
		v3     int32
	)
	result = int32(dword_5d4594_251552)
	if dword_5d4594_251552 != 0 {
		for {
			v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(result + 20))))
			if (int32(*(*uint8)(unsafe.Pointer(uintptr(result + 4)))) & 48) == 0 {
				a1(result, a2)
			}
			result = v3
			if v3 == 0 {
				break
			}
		}
	}
	return result
}
func sub_410670(a1 func(int32, int32), a2 int32) int32 {
	var (
		result int32
		v3     int32
	)
	result = int32(dword_5d4594_251552)
	if dword_5d4594_251552 != 0 {
		for {
			v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(result + 20))))
			if int32(*(*uint8)(unsafe.Pointer(uintptr(result + 4))))&32 != 0 {
				a1(result, a2)
			}
			result = v3
			if v3 == 0 {
				break
			}
		}
	}
	return result
}
func sub_4106A0(a1 int32) int32 {
	var result int32
	if a1 < 0 || a1 >= 256 {
		result = 0
	} else {
		result = int32(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_251556 + uint32(a1*4)))))
	}
	return result
}
func sub_4106C0() {
	var (
		v1 *uint32
		v2 *uint32
		v3 *uint32
		v4 *uint32
	)
	for i := int32(0); i < 8192; i++ {
		v1 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(dword_5D4594_251544), unsafe.Sizeof(uint32(0))*uintptr(i))))))
		if v1 != nil {
			for {
				v2 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*4)))))
				alloc.Free(unsafe.Pointer(v1))
				v1 = v2
				if v2 == nil {
					break
				}
			}
		}
	}
	v3 = *(**uint32)(unsafe.Pointer(&dword_5d4594_251548))
	if dword_5d4594_251548 != 0 {
		for {
			v4 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*5)))))
			alloc.Free(unsafe.Pointer(v3))
			v3 = v4
			if v4 == nil {
				break
			}
		}
	}
	alloc.Free(unsafe.Pointer(dword_5D4594_251544))
	*(*unsafe.Pointer)(unsafe.Pointer(&dword_5d4594_251556)) = nil
}
func sub_410730() *uint32 {
	var (
		result *uint32
		v1     *uint32
	)
	result = *(**uint32)(unsafe.Pointer(&dword_5d4594_251560))
	if dword_5d4594_251560 != 0 {
		for {
			v1 = (*uint32)(unsafe.Pointer(uintptr(*result)))
			alloc.Free(unsafe.Pointer(result))
			result = v1
			if v1 == nil {
				break
			}
		}
		dword_5d4594_251560 = 0
	} else {
		dword_5d4594_251560 = 0
	}
	return result
}
func nox_xxx_wallSecretBlock_410760(a1 *uint32) *uint32 {
	var result *uint32
	result = a1
	*a1 = dword_5d4594_251560
	dword_5d4594_251560 = uint32(uintptr(unsafe.Pointer(a1)))
	return result
}
func nox_xxx_wallSecretGetFirstWall_410780() unsafe.Pointer {
	return *(*unsafe.Pointer)(unsafe.Pointer(&dword_5d4594_251560))
}
func nox_xxx_wallSecretNext_410790(a1 *int32) int32 {
	var result int32
	if a1 != nil {
		result = *a1
	} else {
		result = 0
	}
	return result
}
func sub_4107A0(lpMem unsafe.Pointer) *int32 {
	var (
		result *int32
		v2     *int32
	)
	_ = v2
	result = *(**int32)(unsafe.Pointer(&dword_5d4594_251560))
	v2 = nil
	if dword_5d4594_251560 != 0 {
		for unsafe.Pointer(result) != lpMem {
			v2 = result
			result = (*int32)(unsafe.Pointer(uintptr(nox_xxx_wallSecretNext_410790(result))))
			if result == nil {
				return result
			}
		}
		if result == *(**int32)(unsafe.Pointer(&dword_5d4594_251560)) {
			dword_5d4594_251560 = uint32(nox_xxx_wallSecretNext_410790(result))
		} else {
			*v2 = nox_xxx_wallSecretNext_410790(result)
		}
		lpMem = nil
	}
	return result
}
func nox_xxx_wallBreackableListClear_410810() *uint32 {
	var (
		result *uint32
		v1     *uint32
	)
	result = *(**uint32)(unsafe.Pointer(&dword_5d4594_251564))
	if dword_5d4594_251564 != 0 {
		for {
			v1 = (*uint32)(unsafe.Pointer(uintptr(*result)))
			alloc.Free(unsafe.Pointer(result))
			result = v1
			if v1 == nil {
				break
			}
		}
		dword_5d4594_251564 = 0
	} else {
		dword_5d4594_251564 = 0
	}
	return result
}
func nox_xxx_wallBreackableListAdd_410840(a1 int32) *uint32 {
	var result *uint32
	result = (*uint32)(alloc.Calloc(1, 8))
	*(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*1)) = uint32(a1)
	*result = dword_5d4594_251564
	dword_5d4594_251564 = uint32(uintptr(unsafe.Pointer(result)))
	return result
}
func nox_xxx_wallGetFirstBreakableCli_410870() unsafe.Pointer {
	return *(*unsafe.Pointer)(unsafe.Pointer(&dword_5d4594_251564))
}
func nox_xxx_wallGetNextBreakableCli_410880(a1 *int32) int32 {
	var result int32
	if a1 != nil {
		result = *a1
	} else {
		result = 0
	}
	return result
}
func nox_xxx_wallBreakableRemove_410890(lpMem *int32) *int32 {
	var (
		result *int32
		v2     *int32
	)
	_ = v2
	result = *(**int32)(unsafe.Pointer(&dword_5d4594_251564))
	v2 = nil
	if dword_5d4594_251564 != 0 {
		for result != lpMem {
			v2 = result
			result = (*int32)(unsafe.Pointer(uintptr(nox_xxx_wallGetNextBreakableCli_410880(result))))
			if result == nil {
				return result
			}
		}
		if result == *(**int32)(unsafe.Pointer(&dword_5d4594_251564)) {
			dword_5d4594_251564 = uint32(nox_xxx_wallGetNextBreakableCli_410880(result))
		} else {
			*v2 = nox_xxx_wallGetNextBreakableCli_410880(result)
		}
		alloc.Free(unsafe.Pointer(lpMem))
	}
	return result
}
func nox_thing_read_WALL_410900(f *nox_memfile, a2 *byte) int32 {
	var (
		v4  *byte
		v7  int32
		v9  int32
		v11 int32
		v14 int32
		v15 int8
		v16 int32
		v22 uint8
		v23 int32
		v24 int32
		v25 int32
		v26 int32
		v27 uint32
		v29 int32
		v31 int32
		v33 int32
		v39 int32
		v40 int32
		v41 int32
		i   int32
		v43 uint8
		v44 uint8
		v45 uint8
		v46 uint8
		v47 uint8
		v48 uint32
		v49 uint32
	)
	if dword_5d4594_251540 >= 80 {
		return 0
	}
	v4 = a2
	f.Skip(4)
	v43 = f.ReadU8()
	nox_memfile_read(unsafe.Pointer(a2), 1, int32(v43), f)
	*(*byte)(unsafe.Add(unsafe.Pointer(a2), v43)) = 0
	libc.StrNCpy((*byte)(memmap.PtrOff(0x85B3FC, dword_5d4594_251540*0x302C+0xA824)), a2, 31)
	v7 = f.ReadI32()
	*memmap.PtrUint32(0x85B3FC, dword_5d4594_251540*0x302C+0xA844) = uint32(v7)
	v9 = f.ReadI32()
	*memmap.PtrUint16(0x85B3FC, dword_5d4594_251540*0x302C+43080) = uint16(int16(v9))
	v11 = f.ReadI32()
	*memmap.PtrUint16(0x85B3FC, dword_5d4594_251540*0x302C+0xA84A) = uint16(int16(v11))
	*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v11))), 0)) = f.ReadU8()
	*memmap.PtrUint8(0x85B3FC, dword_5d4594_251540*0x302C+0xA84C) = uint8(int8(v11))
	*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v11))), 0)) = f.ReadU8()
	*memmap.PtrUint8(0x85B3FC, dword_5d4594_251540*0x302C+0xA84D) = uint8(int8(v11))
	nox_memfile_read64align_40AD60((*byte)(unsafe.Pointer(&v39)), 1, 1, f)
	v15 = int8(v39)
	v16 = int32(dword_5d4594_251540 * 0x302C)
	*memmap.PtrUint8(0x85B3FC, uint32(v16+0xA84E)) = uint8(int8(v39))
	libc.MemSet(memmap.PtrOff(0x85B3FC, uint32(v16+0xA84F)), 0, 512)
	if int32(v15) != 0 {
		v14 = 0
		v48 = 0
		for {
			if int32(v48) >= 512 {
				return 0
			}
			v44 = f.ReadU8()
			nox_memfile_read(unsafe.Pointer(v4), 1, int32(v44), f)
			*(*byte)(unsafe.Add(unsafe.Pointer(v4), v44)) = 0
			libc.StrNCpy((*byte)(memmap.PtrOff(0x85B3FC, dword_5d4594_251540*0x302C+0xA84F+v48)), v4, 64)
			v14++
			v48 += 64
			if v14 >= int32(uint8(int8(v39))) {
				break
			}
		}
	}
	v45 = f.ReadU8()
	nox_memfile_read(unsafe.Pointer(v4), 1, int32(v45), f)
	*(*byte)(unsafe.Add(unsafe.Pointer(v4), v45)) = 0
	libc.StrNCpy((*byte)(memmap.PtrOff(0x85B3FC, dword_5d4594_251540*0x302C+0xAA4F)), v4, 64)
	v46 = f.ReadU8()
	nox_memfile_read(unsafe.Pointer(v4), 1, int32(v46), f)
	*(*byte)(unsafe.Add(unsafe.Pointer(v4), v46)) = 0
	libc.StrNCpy((*byte)(memmap.PtrOff(0x85B3FC, dword_5d4594_251540*0x302C+0xAA8F)), v4, 64)
	v47 = f.ReadU8()
	nox_memfile_read(unsafe.Pointer(v4), 1, int32(v47), f)
	*(*byte)(unsafe.Add(unsafe.Pointer(v4), v47)) = 0
	libc.StrNCpy((*byte)(memmap.PtrOff(0x85B3FC, dword_5d4594_251540*0x302C+0xAACF)), v4, 64)
	v22 = f.ReadU8()
	v23 = int32(dword_5d4594_251540 * 0x302C)
	*memmap.PtrUint8(0x85B3FC, uint32(v23+0xAB11)) = v22
	v24 = 0
	libc.MemSet(memmap.PtrOff(0x85B3FC, uint32(v23+0xC914)), 0, 3840)
	v49 = 0
	for {
		nox_memfile_read64align_40AD60((*byte)(unsafe.Pointer(&v40)), 1, 1, f)
		v25 = 0
		i = int32(uint8(int8(v40)))
		for v41 = 0; v41 < i; v41++ {
			v26 = 0
			v27 = (uint32(v25) + v49) * 8
			for {
				*memmap.PtrUint8(0x85B3FC, dword_5d4594_251540*0x302C+0xD814+uint32(v26)+uint32(v24)) = uint8(int8(v40))
				v29 = f.ReadI32()
				*memmap.PtrUint32(0x85B3FC, dword_5d4594_251540*0x302C+0xAB14+v27) = uint32(v29)
				v31 = f.ReadI32()
				*memmap.PtrUint32(0x85B3FC, dword_5d4594_251540*0x302C+43800+v27) = uint32(v31)
				v33 = f.ReadI32()
				if v33 == -1 {
					f.Skip(1)
					var v uint8 = f.ReadU8()
					f.Skip(int32(v))
				}
				v26 += 15
				v27 += 1920
				if v26 >= 60 {
					break
				}
			}
			v25 = v41 + 1
		}
		v24++
		v49 += 16
		if int32(v49) >= 240 {
			break
		}
	}
	var v37 uint32 = uint32(f.ReadI32())
	if v37 != 0x454E4420 {
		return 0
	}
	dword_5d4594_251540++
	return 1
}
func sub_410D40(a1 int32) *byte {
	return (*byte)(memmap.PtrOff(0x85B3FC, uint32(a1*0x302C+0xA824)))
}
func nox_xxx_wallTileByName_410D60(a1 *byte) int32 {
	var (
		v1 int32
		i  *byte
	)
	v1 = 0
	if dword_5d4594_251540 <= 0 {
		return -1
	}
	for i = (*byte)(memmap.PtrOff(0x85B3FC, 43044)); libc.StrCmp(a1, i) != 0; i = (*byte)(unsafe.Add(unsafe.Pointer(i), 12332)) {
		if uint32(func() int32 {
			p := &v1
			*p++
			return *p
		}()) >= dword_5d4594_251540 {
			return -1
		}
	}
	return v1
}
func nox_xxx_mapWallMaxVariation_410DD0(a1 int32, a2 int32, a3 int32) uint8 {
	return *memmap.PtrUint8(0x85B3FC, uint32(a1*0x302C+0xD814+a3*12+a3*3+a2))
}
func nox_xxx_map_410E00(a1 int32) uint8 {
	return *memmap.PtrUint8(0x85B3FC, uint32(a1*0x302C+0xAB11))
}
func nox_xxx_mapWallGetHpByTile_410E20(a1 int32) uint8 {
	return *memmap.PtrUint8(0x85B3FC, uint32(a1*0x302C+0xA84D))
}
func nox_xxx_wallGetBrickTypeMB_410E40(a1 int32) uint8 {
	return *memmap.PtrUint8(0x85B3FC, uint32(a1*0x302C+0xA84E))
}
func nox_xxx_wallGetBrickObj_410E60(a1 int32, a2 int32) *byte {
	var result *byte
	result = (*byte)(memmap.PtrOff(0x587000, 26432))
	if libc.StrLen((*byte)(memmap.PtrOff(0x85B3FC, uint32(a1*0x302C+0xA84F+a2*64)))) != 0 {
		result = (*byte)(memmap.PtrOff(0x85B3FC, uint32(a1*0x302C+0xA84F+a2*64)))
	}
	return result
}
func nox_xxx_wallSoundByTile_410EA0(a1 int32) *byte {
	var result *byte
	result = (*byte)(memmap.PtrOff(0x587000, 26440))
	if libc.StrLen((*byte)(memmap.PtrOff(0x85B3FC, uint32(a1*0x302C+0xAACF)))) != 0 {
		result = (*byte)(memmap.PtrOff(0x85B3FC, uint32(a1*0x302C+0xAACF)))
	}
	return result
}
func nox_xxx_wallFindOpenSound_410EE0(a1 int32) *byte {
	var result *byte
	result = (*byte)(memmap.PtrOff(0x587000, 26456))
	if libc.StrLen((*byte)(memmap.PtrOff(0x85B3FC, uint32(a1*0x302C+0xAA4F)))) != 0 {
		result = (*byte)(memmap.PtrOff(0x85B3FC, uint32(a1*0x302C+0xAA4F)))
	}
	return result
}
func nox_xxx_wallFindCloseSound_410F20(a1 int32) *byte {
	var result *byte
	result = (*byte)(memmap.PtrOff(0x587000, 26472))
	if libc.StrLen((*byte)(memmap.PtrOff(0x85B3FC, uint32(a1*0x302C+0xAA8F)))) != 0 {
		result = (*byte)(memmap.PtrOff(0x85B3FC, uint32(a1*0x302C+0xAA8F)))
	}
	return result
}
func nox_xxx_tileAlloc_410F60_init() int32 {
	ptr_5D4594_2650668 = (**obj_5D4594_2650668_t)(unsafe.Pointer(&make([]unsafe.Pointer, int(ptr_5D4594_2650668_cap))[0]))
	if ptr_5D4594_2650668 == nil {
		return 0
	}
	for i := int32(0); i < ptr_5D4594_2650668_cap; i++ {
		*(**obj_5D4594_2650668_t)(unsafe.Add(unsafe.Pointer(ptr_5D4594_2650668), unsafe.Sizeof((*obj_5D4594_2650668_t)(nil))*uintptr(i))) = &make([]obj_5D4594_2650668_t, int(ptr_5D4594_2650668_cap))[0]
		if *(**obj_5D4594_2650668_t)(unsafe.Add(unsafe.Pointer(ptr_5D4594_2650668), unsafe.Sizeof((*obj_5D4594_2650668_t)(nil))*uintptr(i))) == nil {
			return 0
		}
	}
	return 1
}
func nox_xxx_tileFree_410FC0_free() {
	for i := int32(0); i < ptr_5D4594_2650668_cap; i++ {
		var ptr *obj_5D4594_2650668_t = *(**obj_5D4594_2650668_t)(unsafe.Add(unsafe.Pointer(ptr_5D4594_2650668), unsafe.Sizeof((*obj_5D4594_2650668_t)(nil))*uintptr(i)))
		if ptr != nil {
			alloc.Free(unsafe.Pointer(ptr))
		}
	}
}
func nox_xxx_tileUnusedMB_410FF0() int32 {
	var (
		v0 int32
		v3 int8
		v5 *uint32
		v6 int32
		v7 *uint32
		v8 int32
	)
	v0 = 0
	for i := int32(0); i < ptr_5D4594_2650668_cap; i++ {
		for j := int32(0); j < ptr_5D4594_2650668_cap; j++ {
			var obj *obj_5D4594_2650668_t = (*obj_5D4594_2650668_t)(unsafe.Add(unsafe.Pointer(*(**obj_5D4594_2650668_t)(unsafe.Add(unsafe.Pointer(ptr_5D4594_2650668), unsafe.Sizeof((*obj_5D4594_2650668_t)(nil))*uintptr(j)))), unsafe.Sizeof(obj_5D4594_2650668_t{})*uintptr(i)))
			v3 = int8(obj.field_0)
			if int32(v3)&1 != 0 {
				if obj.field_1 >= *(*int32)(unsafe.Pointer(&dword_5d4594_251568)) {
					v0 = 1
					obj.field_1 = 0
				}
				if obj.field_2 >= int32(*memmap.PtrUint16(0x85B3FC, uint32(obj.field_1*60+0x7F10))) {
					v0 = 1
					obj.field_2 = 0
				}
				v5 = (*uint32)(obj.field_5)
				for v5 != nil {
					v6 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(uint32(0))*2)))
					if v6 >= *(*int32)(unsafe.Pointer(&dword_5d4594_251572)) || *(*uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(uint32(0))*3)) >= uint32(int32(*memmap.PtrUint16(0x85B3FC, uint32(v6*60+0x7010)))) {
						v0 = 1
						obj.field_5 = unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(uint32(0))*4))))
						nox_xxx_tileFreeTileOne_4221E0(int32(uintptr(unsafe.Pointer(v5))))
						v5 = (*uint32)(obj.field_5)
					} else {
						v5 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(uint32(0))*4)))))
					}
				}
			}
			if int32(v3)&2 != 0 {
				if obj.field_6 >= *(*int32)(unsafe.Pointer(&dword_5d4594_251568)) {
					v0 = 1
					obj.field_6 = 0
				}
				if obj.field_7 >= int32(*memmap.PtrUint16(0x85B3FC, uint32(obj.field_6*60+0x7F10))) {
					v0 = 1
					obj.field_7 = 0
				}
				v7 = (*uint32)(obj.field_10)
				for v7 != nil {
					v8 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v7), unsafe.Sizeof(uint32(0))*2)))
					if v8 >= *(*int32)(unsafe.Pointer(&dword_5d4594_251572)) || *(*uint32)(unsafe.Add(unsafe.Pointer(v7), unsafe.Sizeof(uint32(0))*3)) >= uint32(int32(*memmap.PtrUint16(0x85B3FC, uint32(v8*60+0x7010)))) {
						v0 = 1
						obj.field_10 = unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v7), unsafe.Sizeof(uint32(0))*4))))
						nox_xxx_tileFreeTileOne_4221E0(int32(uintptr(unsafe.Pointer(v7))))
						v7 = (*uint32)(obj.field_10)
					} else {
						v7 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v7), unsafe.Sizeof(uint32(0))*4)))))
					}
				}
			}
		}
	}
	return v0
}
func nox_xxx_tileNFromPoint_411160(a1 *float2) int32 {
	var (
		v12 float32 = float32((float64(a1.field_0) + 11.5) * 0.021739131)
		v13 float32 = float32((float64(a1.field_4) + 11.5) * 0.021739131)
		i   int32   = int(v12)
		j   int32   = int(v13)
		v14 float32 = float32(float64(a1.field_0) + 11.5)
		v15 float32 = float32(float64(a1.field_4) + 11.5)
		v4  int32   = int(v14) % 46
		v5  int32   = int(v15) % 46
	)
	if i-1 <= 0 || i >= math.MaxInt8 || j-1 <= 0 || j >= math.MaxInt8 {
		return -1
	}
	var result int32 = 0
	var v16 [2]int32 = [2]int32{}
	if v4 <= v5 {
		if 46-v4 <= v5 {
			var obj *obj_5D4594_2650668_t = (*obj_5D4594_2650668_t)(unsafe.Add(unsafe.Pointer(*(**obj_5D4594_2650668_t)(unsafe.Add(unsafe.Pointer(ptr_5D4594_2650668), unsafe.Sizeof((*obj_5D4594_2650668_t)(nil))*uintptr(i)))), unsafe.Sizeof(obj_5D4594_2650668_t{})*uintptr(j)))
			result = obj.field_6
			if obj.field_10 != nil {
				v16[0] = v4
				v16[1] = v5 - 23
				result = sub_411350((*int32)(obj.field_10), &v16[0], result)
			}
		} else {
			var obj *obj_5D4594_2650668_t = (*obj_5D4594_2650668_t)(unsafe.Add(unsafe.Pointer(*(**obj_5D4594_2650668_t)(unsafe.Add(unsafe.Pointer(ptr_5D4594_2650668), unsafe.Sizeof((*obj_5D4594_2650668_t)(nil))*uintptr(i-1)))), unsafe.Sizeof(obj_5D4594_2650668_t{})*uintptr(j)))
			result = obj.field_1
			if obj.field_5 != nil {
				v16[1] = v5
				v16[0] = v4 + 23
				result = sub_411350((*int32)(obj.field_5), &v16[0], result)
			}
		}
	} else if 46-v4 <= v5 {
		var obj *obj_5D4594_2650668_t = (*obj_5D4594_2650668_t)(unsafe.Add(unsafe.Pointer(*(**obj_5D4594_2650668_t)(unsafe.Add(unsafe.Pointer(ptr_5D4594_2650668), unsafe.Sizeof((*obj_5D4594_2650668_t)(nil))*uintptr(i)))), unsafe.Sizeof(obj_5D4594_2650668_t{})*uintptr(j)))
		result = obj.field_1
		if obj.field_5 != nil {
			v16[1] = v5
			v16[0] = v4 - 23
			result = sub_411350((*int32)(obj.field_5), &v16[0], result)
		}
	} else {
		var obj *obj_5D4594_2650668_t = (*obj_5D4594_2650668_t)(unsafe.Add(unsafe.Pointer(*(**obj_5D4594_2650668_t)(unsafe.Add(unsafe.Pointer(ptr_5D4594_2650668), unsafe.Sizeof((*obj_5D4594_2650668_t)(nil))*uintptr(i)))), unsafe.Sizeof(obj_5D4594_2650668_t{})*uintptr(j-1)))
		result = obj.field_6
		if obj.field_10 != nil {
			v16[0] = v4
			v16[1] = v5 + 23
			result = sub_411350((*int32)(obj.field_10), &v16[0], result)
		}
	}
	return result
}
func sub_411350(a1 *int32, a2 *int32, a3 int32) int32 {
	var (
		v3 *int32
		v4 int32
		v5 int32
	)
	v3 = a1
	if a1 == nil {
		return a3
	}
	v4 = a3
	for {
		v5 = sub_411490(*(*int32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(int32(0))*2)), *(*int32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(int32(0))*3)))
		if sub_4113A0(a2, v5) != 0 {
			v4 = *v3
		}
		v3 = (*int32)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(int32(0))*4)))))
		if v3 == nil {
			break
		}
	}
	return v4
}
func sub_4113A0(a1 *int32, a2 int32) int32 {
	var (
		v2 int32 = *a1
		v3 int32 = *(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*1))
		v4 bool
		v5 bool
	)
	if *a1 == v3 {
		v4 = true
		v5 = true
	} else if *a1 >= v3 {
		v4 = true
		v5 = false
	} else {
		v4 = false
		v5 = true
	}
	var v6 bool
	var v7 bool
	if v2 == 46-v3 {
		v6 = true
		v7 = true
	} else if v2 >= 46-v3 {
		v6 = true
		v7 = false
	} else {
		v6 = false
		v7 = true
	}
	switch a2 {
	case 0:
		if !v5 || !v6 {
			return 0
		}
		return 1
	case 1:
		if !v5 {
			return 0
		}
		return 1
	case 2:
		if !v5 || !v7 {
			return 0
		}
		return 1
	case 3:
		if !v6 {
			return 0
		}
		return 1
	case 4:
		if !v7 {
			return 0
		}
		return 1
	case 5:
		if !v6 || !v4 {
			return 0
		}
		return 1
	case 6:
		if !v4 {
			return 0
		}
		return 1
	case 7:
		if !v4 || !v7 {
			return 0
		}
		return 1
	case 8:
		if v5 {
			return 1
		}
		if !v6 {
			return 0
		}
		return 1
	case 9:
		if !v5 && !v7 {
			return 0
		}
		return 1
	case 10:
		if !v4 && !v7 {
			return 0
		}
		return 1
	case 11:
		if v4 {
			return 1
		}
		if !v6 {
			return 0
		}
		return 1
	}
	return 0
}
func sub_411490(a1 int32, a2 int32) int32 {
	var (
		v2 int32
		v3 int32
		v5 int32
	)
	v2 = int32(*memmap.PtrUint8(0x85B3FC, uint32(a1*60+0x7018)))
	v3 = int32(*memmap.PtrUint8(0x85B3FC, uint32(a1*60+0x7019)))
	if v2 == 3 && v3 == 3 {
		return a2
	}
	if a2 == 0 {
		return 0
	}
	if a2 <= v2-2 {
		return 1
	}
	if a2 == v2-1 {
		return 2
	}
	v5 = v2 + v3*2 - 4
	if a2 < v5 {
		return bool2int(((int32(uint8(int8(v2)))^int32(uint8(int8(a2))))&1) != 0) + 3
	}
	if a2 == v5 {
		return 5
	} else if a2 > (v3+v2)*2-6 {
		if a2 == (v3+v2)*2-5 {
			return 7
		}
		return a2 + (6-v3-v2)*2
	}
	return 6
}
func nox_thing_read_FLOR_411540(f *nox_memfile, a2 *uint8) int32 {
	var (
		v5  int32
		v7  int8
		v8  int8
		v9  int8
		v22 int32
		v23 int32
		v24 int32
		v26 int32
		v29 [6]uint8
		v30 int32
		v31 [32]byte
		v32 uint8
		v33 uint8
	)
	if dword_5d4594_251568 >= 176 {
		return 0
	}
	f.Skip(4)
	v32 = f.ReadU8()
	nox_memfile_read(unsafe.Pointer(&v31[0]), 1, int32(v32), f)
	v31[v32] = 0
	v5 = int32(dword_5d4594_251568 * 3)
	libc.StrNCpy((*byte)(memmap.PtrOff(0x85B3FC, uint32(v5*20+0x7EE4))), &v31[0], 31)
	v7 = f.ReadI8()
	v29[0] = uint8(v7)
	v8 = f.ReadI8()
	v29[1] = uint8(v8)
	v9 = f.ReadI8()
	v29[2] = uint8(v9)
	if int32(v7) != -1 || int32(v29[1]) != 0 || int32(v9) != -1 {
		*memmap.PtrUint32(0x85B3FC, dword_5d4594_251568*60+0x7F14) = nox_color_rgb_4344A0(*(*int32)(unsafe.Pointer(&v29[0])), *(*int32)(unsafe.Pointer(&v29[1])), *(*int32)(unsafe.Pointer(&v29[2])))
	} else {
		*memmap.PtrUint32(0x85B3FC, dword_5d4594_251568*60+0x7F14) = 0x80000000
	}
	*memmap.PtrUint8(0x85B3FC, dword_5d4594_251568*60+0x7F1E) = 0
	if nox_xxx_checkFacade_4117E0((*byte)(memmap.PtrOff(0x85B3FC, dword_5d4594_251568*60+0x7EE4))) == 1 {
		*memmap.PtrUint8(0x85B3FC, dword_5d4594_251568*60+0x7F1E) |= 1
	}
	*memmap.PtrUint32(0x85B3FC, dword_5d4594_251568*60+32520) = uint32(f.ReadI32())
	*memmap.PtrUint32(0x85B3FC, dword_5d4594_251568*60+0x7F0C) = uint32(f.ReadI32())
	var v14 uint8 = f.ReadU8()
	*memmap.PtrUint8(0x85B3FC, dword_5d4594_251568*60+0x7F1D) = v14
	v14 = f.ReadU8()
	*memmap.PtrUint8(0x85B3FC, dword_5d4594_251568*60+0x7F19) = v14
	v14 = f.ReadU8()
	*memmap.PtrUint8(0x85B3FC, dword_5d4594_251568*60+0x7F18) = v14
	v14 = f.ReadU8()
	*memmap.PtrUint8(0x85B3FC, dword_5d4594_251568*60+0x7F1A) = v14
	*memmap.PtrUint16(0x85B3FC, dword_5d4594_251568*60+32530) = 0
	v14 = f.ReadU8()
	*memmap.PtrUint8(0x85B3FC, dword_5d4594_251568*60+0x7F1B) = v14
	*memmap.PtrUint8(0x85B3FC, dword_5d4594_251568*60+32540) = v14
	*memmap.PtrUint16(0x85B3FC, dword_5d4594_251568*60+0x7F10) = uint16(int16(int32(*memmap.PtrUint8(0x85B3FC, dword_5d4594_251568*60+0x7F18)) * int32(*memmap.PtrUint8(0x85B3FC, dword_5d4594_251568*60+0x7F19))))
	v22 = int32(*memmap.PtrUint8(0x85B3FC, dword_5d4594_251568*60+0x7F19)) * int32(*memmap.PtrUint8(0x85B3FC, dword_5d4594_251568*60+0x7F1A))
	v23 = int32(*memmap.PtrUint8(0x85B3FC, dword_5d4594_251568*60+0x7F18))
	*memmap.PtrUint32(0x85B3FC, dword_5d4594_251568*60+0x7F04) = 0
	v24 = v23 * v22
	if v24 > 0 {
		v30 = v24
		for {
			v26 = f.ReadI32()
			*a2 = *memmap.PtrUint8(6112660, 251576)
			if v26 == -1 {
				f.Skip(1)
				v33 = f.ReadU8()
				nox_memfile_read(unsafe.Pointer(a2), 1, int32(v33), f)
				*(*uint8)(unsafe.Add(unsafe.Pointer(a2), v33)) = 0
			}
			v30--
			if v30 == 0 {
				break
			}
		}
	}
	dword_5d4594_251568++
	return 1
}
func nox_xxx_checkFacade_4117E0(a1 *byte) int32 {
	var (
		v1 *byte
		v2 *uint8
	)
	v1 = *(**byte)(memmap.PtrOff(0x587000, 26488))
	if *memmap.PtrUint32(0x587000, 26488) == 0 {
		return 0
	}
	v2 = (*uint8)(memmap.PtrOff(0x587000, 26488))
	for libc.StrCmp(v1, a1) != 0 {
		v1 = (*byte)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), unsafe.Sizeof(uint32(0))*1))))))
		v2 = (*uint8)(unsafe.Add(unsafe.Pointer(v2), 4))
		if v1 == nil {
			return 0
		}
	}
	return 1
}
func nox_thing_read_EDGE_411850(f *nox_memfile, a2 *uint8) int32 {
	var (
		v5  int32
		v7  int32
		v9  int32
		v12 int32
		v14 int32
		v17 uint8
		v19 uint8
		v20 int32
		v21 int32
		v23 int32
		v26 int32
		v28 uint8
		v29 [32]byte
		v30 uint8
		v31 uint8
	)
	if dword_5d4594_251572 >= 64 {
		return 0
	}
	f.Skip(4)
	v30 = f.ReadU8()
	nox_memfile_read(unsafe.Pointer(&v29[0]), 1, int32(v30), f)
	v5 = int32(dword_5d4594_251572 * 3)
	v29[v30] = 0
	libc.StrCpy((*byte)(memmap.PtrOff(0x85B3FC, uint32(v5*20+0x6FE4))), &v29[0])
	v7 = f.ReadI32()
	*memmap.PtrUint32(0x85B3FC, dword_5d4594_251572*60+28680) = uint32(v7)
	v9 = f.ReadI32()
	*memmap.PtrUint32(0x85B3FC, dword_5d4594_251572*60+0x700C) = uint32(v9)
	*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v9))), 0)) = f.ReadU8()
	*memmap.PtrUint8(0x85B3FC, dword_5d4594_251572*60+0x701D) = uint8(int8(v9))
	*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v9))), 0)) = f.ReadU8()
	v28 = uint8(int8(v9))
	v12 = int32(dword_5d4594_251572 * 60)
	*memmap.PtrUint8(0x85B3FC, uint32(v12+0x701A)) = uint8(int8(v9))
	*memmap.PtrUint16(0x85B3FC, uint32(v12+28690)) = 0
	*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v9))), 0)) = f.ReadU8()
	v14 = int32(dword_5d4594_251572 * 60)
	*memmap.PtrUint8(0x85B3FC, uint32(v14+0x701B)) = uint8(int8(v9))
	*memmap.PtrUint8(0x85B3FC, uint32(v14+28700)) = uint8(int8(v9))
	*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v9))), 0)) = f.ReadU8()
	if int32(uint8(int8(v9))) == 1 {
		return 0
	}
	v17 = f.ReadU8()
	*memmap.PtrUint8(0x85B3FC, dword_5d4594_251572*60+0x7019) = v17
	v19 = f.ReadU8()
	v20 = int32(dword_5d4594_251572 * 60)
	*memmap.PtrUint8(0x85B3FC, uint32(v20+0x7018)) = v19
	*memmap.PtrUint16(0x85B3FC, uint32(v20+0x7010)) = uint16(int16((int32(v17) + int32(v19)) * 2))
	*memmap.PtrUint32(0x85B3FC, uint32(v20+0x7004)) = 0
	if int32(v28)*2*(int32(v17)+int32(v19)) > 0 {
		v21 = int32(v28) * 2 * (int32(v17) + int32(v19))
		for {
			v23 = f.ReadI32()
			*a2 = *memmap.PtrUint8(6112660, 251580)
			if v23 == -1 {
				f.Skip(1)
				v31 = f.ReadU8()
				nox_memfile_read(unsafe.Pointer(a2), 1, int32(v31), f)
				*(*uint8)(unsafe.Add(unsafe.Pointer(a2), v31)) = 0
			}
			v21--
			if v21 == 0 {
				break
			}
		}
	}
	v26 = f.ReadI32()
	if uint32(v26) != 0x454E4420 {
		return 0
	}
	dword_5d4594_251572++
	return 1
}
func nox_xxx_mapTileAllowTeleport_411A90(a1 *float2) int32 {
	var (
		v1 int32
		v2 *byte
		v3 int32
	)
	if *memmap.PtrInt32(0x587000, 26520) == -1 {
		v1 = 0
		v2 = (*byte)(memmap.PtrOff(0x85B3FC, 32484))
		for {
			if libc.StrCmp(v2, CString("WaterNoTeleport")) == 0 {
				*memmap.PtrUint32(0x587000, 26516) = uint32(v1)
			} else if libc.StrCmp(v2, CString("WaterDeepNoTeleport")) == 0 {
				*memmap.PtrUint32(0x587000, 26520) = uint32(v1)
			} else if libc.StrCmp(v2, CString("WaterShallowNoTeleport")) == 0 {
				*memmap.PtrUint32(0x587000, 26524) = uint32(v1)
			} else if libc.StrCmp(v2, CString("WaterSwampDeepNoTeleport")) == 0 {
				*memmap.PtrUint32(0x587000, 26528) = uint32(v1)
			} else if libc.StrCmp(v2, CString("WaterSwampShallowNoTeleport")) == 0 {
				*memmap.PtrUint32(0x587000, 26532) = uint32(v1)
			}
			v2 = (*byte)(unsafe.Add(unsafe.Pointer(v2), 60))
			v1++
			if int32(uintptr(unsafe.Pointer(v2))) >= int32(uintptr(memmap.PtrOff(0x85B3FC, 43044))) {
				break
			}
		}
	}
	v3 = nox_xxx_tileNFromPoint_411160(a1)
	return bool2int(uint32(v3) == *memmap.PtrUint32(0x587000, 26516) || uint32(v3) == *memmap.PtrUint32(0x587000, 26520) || uint32(v3) == *memmap.PtrUint32(0x587000, 26524) || uint32(v3) == *memmap.PtrUint32(0x587000, 26528) || uint32(v3) == *memmap.PtrUint32(0x587000, 26532))
}
func nox_xxx_parseWeapColor_411C40(a1 *byte, a2 *byte, a3 int32) int32 {
	var (
		v3     int32
		result int32
	)
	v3 = sub_411C80(a1)
	if v3 == -1 {
		result = 0
	} else {
		result = bool2int(nox_xxx_parseBinColor_411CF0(a2, (*uint8)(unsafe.Pointer(uintptr(a3+(v3+4)*2+v3+4)))) != nil)
	}
	return result
}
func sub_411C80(a1 *byte) int32 {
	var (
		v1 *byte
		v2 int32
		v3 *uint8
	)
	v1 = *(**byte)(memmap.PtrOff(0x587000, 31096))
	v2 = 0
	if *memmap.PtrUint32(0x587000, 31096) == 0 {
		return -1
	}
	v3 = (*uint8)(memmap.PtrOff(0x587000, 31096))
	for libc.StrCmp(v1, a1) != 0 {
		v1 = (*byte)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v3))), unsafe.Sizeof(uint32(0))*1))))))
		v3 = (*uint8)(unsafe.Add(unsafe.Pointer(v3), 4))
		v2++
		if v1 == nil {
			return -1
		}
	}
	return v2
}
func nox_xxx_parseBinColor_411CF0(a1 *byte, a2 *uint8) *byte {
	var (
		result *byte
		v3     int8
		v4     int8
		v5     [8]byte
		v6     int8
	)
	libc.StrCpy(&v5[0], CString(" =\t\n\r"))
	result = libc.StrTok(a1, &v5[0])
	if result != nil {
		v6 = int8(libc.Atoi(libc.GoString(result)))
		result = libc.StrTok(nil, &v5[0])
		if result != nil {
			v3 = int8(libc.Atoi(libc.GoString(result)))
			result = libc.StrTok(nil, &v5[0])
			if result != nil {
				v4 = int8(libc.Atoi(libc.GoString(result)))
				*(*uint8)(unsafe.Add(unsafe.Pointer(a2), 1)) = uint8(v3)
				*(*uint8)(unsafe.Add(unsafe.Pointer(a2), 2)) = uint8(v4)
				*a2 = uint8(v6)
				result = (*byte)(unsafe.Pointer(uintptr(1)))
			}
		}
	}
	return result
}
func sub_411D90(a1 *byte, a2 *byte, a3 int32) int32 {
	var (
		v3 *byte
		v4 int32
		v5 int32
		v7 [8]byte
	)
	libc.StrCpy(&v7[0], CString(" =\n\r\t"))
	v3 = libc.StrTok(a2, &v7[0])
	if v3 != nil {
		if libc.StrCmp(a1, CString("EFFECTIVENESS")) == 0 {
			v4 = 0
		LABEL_10:
			v5 = sub_411C80(v3)
			*(*uint32)(unsafe.Pointer(uintptr(a3 + v4*4 + 36))) = uint32(v5)
			return bool2int(v5 != -1)
		}
		if libc.StrCmp(a1, CString("MATERIAL")) == 0 {
			v4 = 1
			goto LABEL_10
		}
		if libc.StrCmp(a1, CString("PRIMARYENCHANTMENT")) == 0 {
			v4 = 2
			goto LABEL_10
		}
		if libc.StrCmp(a1, CString("SECONDARYENCHANTMENT")) == 0 {
			v4 = 3
			goto LABEL_10
		}
	}
	return 0
}
func sub_411E60(a1 int32, a2 *byte, a3 int32) int32 {
	var (
		v3     uint32
		result int32
		v5     int32
	)
	v5 = 0
	v3 = uint32(libc.StrSpn(a2, CString("= ")))
	if v3 >= uint32(libc.StrLen(a2)) {
		*(*uint8)(unsafe.Pointer(uintptr(a3 + 62))) = uint8(int8(v5))
		result = 1
	} else {
		set_bitmask_flags_from_plus_separated_names_423930((*byte)(unsafe.Add(unsafe.Pointer(a2), v3)), (*uint32)(unsafe.Pointer(&v5)), (**byte)(memmap.PtrOff(0x587000, 29456)))
		result = 1
		*(*uint8)(unsafe.Pointer(uintptr(a3 + 62))) = uint8(int8(v5))
	}
	return result
}
func sub_411ED0(a1 int32, a2 *byte, a3 int32) *byte {
	var (
		result *byte
		v4     [8]byte
	)
	libc.StrCpy(&v4[0], CString(" =\n\r\t"))
	result = libc.StrTok(a2, &v4[0])
	if result != nil {
		*(*uint32)(unsafe.Pointer(uintptr(a3 + 52))) = uint32(libc.Atoi(libc.GoString(result)))
		result = (*byte)(unsafe.Pointer(uintptr(1)))
	}
	return result
}
func sub_411F20(a1 int32, a2 *byte, a3 int32) *byte {
	var (
		result *byte
		v4     [8]byte
	)
	libc.StrCpy(&v4[0], CString(" =\n\r\t"))
	result = libc.StrTok(a2, &v4[0])
	if result != nil {
		*(*uint16)(unsafe.Pointer(uintptr(a3 + 60))) = uint16(int16(libc.Atoi(libc.GoString(result))))
		result = (*byte)(unsafe.Pointer(uintptr(1)))
	}
	return result
}
func sub_411F70(a1 int32, a2 *byte, a3 int32) *byte {
	var (
		result *byte
		v4     [8]byte
	)
	libc.StrCpy(&v4[0], CString(" =\n\r\t"))
	result = libc.StrTok(a2, &v4[0])
	if result != nil {
		*(*uint16)(unsafe.Pointer(uintptr(a3 + 72))) = uint16(int16(libc.Atoi(libc.GoString(result))))
		result = (*byte)(unsafe.Pointer(uintptr(1)))
	}
	return result
}
func sub_411FC0(a1 int32, a2 *byte, a3 int32) *byte {
	var (
		result *byte
		v4     float64
		v5     [8]byte
	)
	libc.StrCpy(&v5[0], CString(" =\n\r\t"))
	result = libc.StrTok(a2, &v5[0])
	if result != nil {
		v4 = libc.Atof(libc.GoString(result))
		result = (*byte)(unsafe.Pointer(uintptr(1)))
		*(*float32)(unsafe.Pointer(uintptr(a3 + 64))) = float32(v4)
	}
	return result
}
func nox_xxx_parseWeapDamageType_412010(a1 int32, a2 *byte, a3 int32) *byte {
	var (
		result *byte
		v4     int32
		v5     [8]byte
	)
	libc.StrCpy(&v5[0], CString(" =\n\r\t"))
	result = libc.StrTok(a2, &v5[0])
	if result != nil {
		v4 = nox_xxx_parseDamageTypeByName_4E0A00(result)
		*(*uint32)(unsafe.Pointer(uintptr(a3 + 76))) = uint32(v4)
		result = (*byte)(unsafe.Pointer(uintptr(bool2int(v4 != 18))))
	}
	return result
}
func sub_412060(a1 int32, a2 *byte, a3 int32) *byte {
	var (
		result *byte
		v4     float64
		v5     [8]byte
	)
	libc.StrCpy(&v5[0], CString(" =\n\r\t"))
	result = libc.StrTok(a2, &v5[0])
	if result != nil {
		v4 = libc.Atof(libc.GoString(result))
		result = (*byte)(unsafe.Pointer(uintptr(1)))
		*(*float32)(unsafe.Pointer(uintptr(a3 + 68))) = float32(v4)
	}
	return result
}
func sub_4120B0(a1 int32, a2 *byte, a3 int32) *byte {
	var (
		result *byte
		v4     float64
		v5     [8]byte
	)
	libc.StrCpy(&v5[0], CString(" =\n\r\t"))
	result = libc.StrTok(a2, &v5[0])
	if result != nil {
		v4 = libc.Atof(libc.GoString(result))
		result = (*byte)(unsafe.Pointer(uintptr(1)))
		*(*float32)(unsafe.Pointer(uintptr(a3 + 64))) = float32(v4)
	}
	return result
}
func nox_xxx_parseEnchWorth_412310_parse_worth(a1 *byte, a2 *byte, a3 *obj_412ae0_t) int32 {
	var (
		result *byte
		v4     [8]byte
	)
	libc.StrCpy(&v4[0], CString(" =\n\r\t"))
	result = libc.StrTok(a2, &v4[0])
	if result == nil {
		return 0
	}
	a3.field_5 = int32(libc.Atoi(libc.GoString(result)))
	return 1
}
func nox_xxx_parseEnchColor_412360_parse_color(a1 *byte, a2 *byte, a3 *obj_412ae0_t) int32 {
	return bool2int(nox_xxx_parseBinColor_411CF0(a2, (*uint8)(unsafe.Pointer(&a3.field_6))) != nil)
}
func nox_xxx_parseEnchEffect_412380_parse_attack_effect(a1 *byte, a2 *byte, obj *obj_412ae0_t) int32 {
	var v9 [8]byte
	libc.StrCpy(&v9[0], CString(" \n\r\t="))
	var v3 *byte = libc.StrTok(a2, &v9[0])
	if v3 == nil {
		return 0
	}
	var ent *table_26792_t = nil
	for i := int32(0); i < table_26792_cnt; i++ {
		if libc.StrCmp(table_26792[i].name, v3) == 0 {
			ent = &table_26792[i]
			break
		}
	}
	if ent == nil {
		return 0
	}
	if libc.StrCmp(a1, CString("ATTACKEFFECT")) == 0 {
		obj.field_10 = ent.fnc
	} else if libc.StrCmp(a1, CString("ATTACKPREHITEFFECT")) == 0 {
		obj.field_13 = ent.fnc
	} else {
		if libc.StrCmp(a1, CString("ATTACKPREDAMAGEEFFECT")) != 0 {
			return 0
		}
		obj.field_16 = ent.fnc
	}
	if ent.parse_fnc != nil && ent.parse_fnc(a1, a2, obj) == 0 {
		return 0
	}
	return 1
}
func nox_xxx_parseEnchDefend_412490_parse_defend_effect(a1 *byte, a2 *byte, obj *obj_412ae0_t) int32 {
	var v9 [8]byte
	libc.StrCpy(&v9[0], CString(" \n\r\t="))
	var v3 *byte = libc.StrTok(a2, &v9[0])
	if v3 == nil {
		return 0
	}
	var ent *table_27008_t = nil
	for i := int32(0); i < table_27008_cnt; i++ {
		if libc.StrCmp(table_27008[i].name, v3) == 0 {
			ent = &table_27008[i]
			break
		}
	}
	if ent == nil {
		return 0
	}
	if libc.StrCmp(a1, CString("DEFENDEFFECT")) == 0 {
		obj.field_19 = ent.fnc
	} else {
		if libc.StrCmp(a1, CString("DEFENDCOLLIDEEFFECT")) != 0 {
			return 0
		}
		obj.field_22 = ent.fnc
	}
	if ent.parse_fnc != nil && ent.parse_fnc(a1, a2, obj) == 0 {
		return 0
	}
	return 1
}
func nox_xxx_parseEnchEngageOrDisEng_412580_parse_engage_effect(a1 *byte, a2 *byte, obj *obj_412ae0_t) int32 {
	var v9 [8]byte
	libc.StrCpy(&v9[0], CString(" =\n\r\t"))
	var v3 *byte = libc.StrTok(a2, &v9[0])
	if v3 == nil {
		return 0
	}
	var ent *table_27168_t = nil
	for i := int32(0); i < table_27168_cnt; i++ {
		if libc.StrCmp(table_27168[i].name, v3) == 0 {
			ent = &table_27168[i]
			break
		}
	}
	if ent == nil {
		return 0
	}
	if libc.StrCmp(a1, CString("ENGAGEEFFECT")) == 0 {
		obj.field_28 = ent.fnc
	} else {
		if libc.StrCmp(a1, CString("DISENGAGEEFFECT")) != 0 {
			return 0
		}
		obj.field_29 = ent.fnc
	}
	if ent.parse_fnc != nil && ent.parse_fnc(a1, a2, obj) == 0 {
		return 0
	}
	return 1
}
func nox_xxx_parseEnchUpdate_412670_parse_update_effect(a1 *byte, a2 *byte, obj *obj_412ae0_t) int32 {
	var v8 [8]byte
	libc.StrCpy(&v8[0], CString(" =\n\r\t"))
	var v3 *byte = libc.StrTok(a2, &v8[0])
	if v3 == nil {
		return 0
	}
	var ent *table_27104_t = nil
	for i := int32(0); i < table_27104_cnt; i++ {
		if libc.StrCmp(table_27104[i].name, v3) == 0 {
			ent = &table_27104[i]
			break
		}
	}
	if ent == nil {
		return 0
	}
	if libc.StrCmp(a1, CString("UPDATEEFFECT")) != 0 {
		return 0
	}
	obj.field_25 = ent.fnc
	if ent.parse_fnc == nil || ent.parse_fnc(a1, a2, obj) != 0 {
		return 1
	}
	return 0
}
func nox_xxx_parseEnchAllowedWeapon_412740_parse_allowed_weapons(a1 *byte, a2 *byte, a3 *obj_412ae0_t) int32 {
	a3.field_7 = 0
	return nox_xxx_parseAllowedWeaponArmor_412760(a2, &a3.field_7)
}
func nox_xxx_parseAllowedWeaponArmor_412760(a1 *byte, a2 *int32) int32 {
	var (
		v2  *byte
		v4  *int32
		v5  *byte
		v6  int32
		v7  int32
		v8  *uint8
		v9  int32
		v10 *byte
		v11 int32
		v12 [8]byte
	)
	libc.StrCpy(&v12[0], CString(" =\r\n\t"))
	v11 = 1
	v2 = libc.StrTok(a1, &v12[0])
	if v2 != nil {
		v4 = a2
		for {
			v5 = *(**byte)(memmap.PtrOff(0x587000, 28904))
			v6 = 0
			v7 = 0
			if *memmap.PtrUint32(0x587000, 28904) != 0 {
				v8 = (*uint8)(memmap.PtrOff(0x587000, 28904))
				for libc.StrCmp(v5, v2) != 0 {
					v5 = (*byte)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v8))), unsafe.Sizeof(uint32(0))*2))))))
					v8 = (*uint8)(unsafe.Add(unsafe.Pointer(v8), 8))
					v7++
					if v5 == nil {
						v4 = a2
						v6 = 0
						goto LABEL_10
					}
				}
				v6 = int32(*memmap.PtrUint32(0x587000, uint32(v7*8+0x70EC)))
				v4 = a2
			}
		LABEL_10:
			if *memmap.PtrUint32(0x587000, uint32(v7*8+0x70E8)) == 0 {
				break
			}
			if v11 == 1 {
				v9 = v6 | *v4
			} else {
				v9 = ^v6 & *v4
			}
			*v4 = v9
			v10 = libc.StrTok(nil, &v12[0])
			if v10 != nil {
				if libc.StrCmp(v10, CString("+")) == 0 {
					v11 = 1
				} else {
					if libc.StrCmp(v10, CString("-")) != 0 {
						return 0
					}
					v11 = 0
				}
				v2 = libc.StrTok(nil, &v12[0])
				if v2 != nil {
					continue
				}
			}
			return 1
		}
	}
	return 0
}
func nox_xxx_parseEnchAllowedArmor_4128A0_parse_allowed_armor(a1 *byte, a2 *byte, a3 *obj_412ae0_t) int32 {
	a3.field_8 = 0
	return nox_xxx_parseAllowedWeaponArmor_412760(a2, &a3.field_8)
}
func nox_xxx_parseAllowedPosition_4128C0_parse_allowed_pos(a1 *byte, a2 *byte, a3 *obj_412ae0_t) int32 {
	var (
		v3 uint32
		v5 int32
	)
	v5 = 0
	v3 = uint32(libc.StrSpn(a2, CString("= ")))
	if v3 >= uint32(libc.StrLen(a2)) {
		*(*uint8)(unsafe.Pointer(&a3.field_9)) = uint8(int8(v5))
		return 1
	}
	set_bitmask_flags_from_plus_separated_names_423930((*byte)(unsafe.Add(unsafe.Pointer(a2), v3)), (*uint32)(unsafe.Pointer(&v5)), (**byte)(memmap.PtrOff(0x587000, 26776)))
	*(*uint8)(unsafe.Pointer(&a3.field_9)) = uint8(int8(v5))
	return 1
}
func nox_xxx_parseModifierBin_412930(a1 *byte, a2 *byte) int32 {
	var (
		v2  *File
		v3  *File
		v10 [256]byte
	)
	byte_5D4594_251584[0] = nil
	byte_5D4594_251584[1] = nil
	byte_5D4594_251584[2] = nil
	byte_5D4594_251596 = 0
	dword_5d4594_251600 = 0
	*memmap.PtrUint32(6112660, 251604) = 0
	dword_5d4594_251608 = 0
	*memmap.PtrUint32(6112660, 251612) = 0
	v2 = nox_binfile_open_408CC0(a1, 0)
	v3 = v2
	if v2 == nil {
		return 0
	}
	if nox_binfile_cryptSet_408D40((*File)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v2)))))), 13) == 0 {
		return 0
	}
	for nox_xxx_parseString_409470(v3, (*uint8)(unsafe.Pointer(&v10[0]))) != 0 {
		if libc.StrCmp(&v10[0], CString("WEAPON_DEFINITIONS")) == 0 {
			if nox_xxx_parseWeaponDef_412D40(int32(uintptr(unsafe.Pointer(&v10[0]))), v3, a2) == 0 {
				return 0
			}
		} else if libc.StrCmp(&v10[0], CString("ARMOR_DEFINITIONS")) == 0 {
			if sub_412ED0(int32(uintptr(unsafe.Pointer(&v10[0]))), v3, a2) == 0 {
				return 0
			}
		} else if libc.StrCmp(&v10[0], CString("EFFECTIVENESS")) != 0 && libc.StrCmp(&v10[0], CString("MATERIAL")) != 0 && libc.StrCmp(&v10[0], CString("ENCHANTMENT")) != 0 || nox_xxx_parseModifDesc_412AE0(&v10[0], v3, a2) == 0 {
			return 0
		}
	}
	nox_binfile_close_408D90(v3)
	for i := unsafe.Pointer(sub_413370()); i != nil; i = unsafe.Pointer(uintptr(sub_413380(int32(uintptr(i))))) {
	}
	for j := unsafe.Pointer(sub_413390()); j != nil; j = unsafe.Pointer(uintptr(sub_4133A0(int32(uintptr(j))))) {
	}
	var v7 int32 = 0
	for k := int32(0); k < 3; k++ {
		for l := (*obj_412ae0_t)(nox_xxx_modifGetModifListByType_4133B0(k)); l != nil; l = nox_xxx_modifNext_4133C0(l) {
			v7++
		}
	}
	return 1
}
func nox_xxx_parseModifDesc_412AE0(a1 *byte, a2 *File, a3 *byte) int32 {
	var v9 [256]byte
	for {
		nox_xxx_parseString_409470(a2, (*uint8)(unsafe.Pointer(&v9[0])))
		if libc.StrCmp(&v9[0], CString("END")) == 0 {
			return 1
		}
		var v3 *obj_412ae0_t = new(obj_412ae0_t)
		if v3 == nil {
			return 0
		}
		v3.field_1 = uint32(func() int32 {
			p := &byte_5D4594_251596
			x := *p
			*p++
			return x
		}())
		if nox_xxx_parseModifAddByType_412C60(a1, v3) == 0 {
			return 0
		}
		var v5 *byte = (*byte)(alloc.Calloc(libc.StrLen(&v9[0])+1, 1))
		if v5 == nil {
			return 0
		}
		libc.StrCpy(v5, &v9[0])
		v3.field_0 = v5
		for {
			nox_xxx_parseString_409470(a2, (*uint8)(unsafe.Pointer(&v9[0])))
			if libc.StrCmp(&v9[0], CString("END")) == 0 {
				break
			}
			var v6 *table_28760_t = &table_28760[0]
			if v6.parse_fnc != nil {
				for libc.StrCmp(v6.name, &v9[0]) != 0 {
					v6 = (*table_28760_t)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(table_28760_t{})*1))
					if v6.parse_fnc == nil {
						goto LABEL_12
					}
				}
				nox_xxx_parseRead_4093E0(a2, a3, 0x40000)
				if v6.parse_fnc(&v9[0], a3, v3) == 0 {
					return 0
				}
			}
		LABEL_12:
			if *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v6))), unsafe.Sizeof(uint32(0))*1))) == 0 {
				return 0
			}
		}
	}
}
func nox_xxx_parseModifAddByType_412C60(a1 *byte, a2 *obj_412ae0_t) int32 {
	if libc.StrCmp(a1, CString("EFFECTIVENESS")) == 0 {
		a2.field_35 = nil
		a2.field_34 = byte_5D4594_251584[0]
		if byte_5D4594_251584[0] != nil {
			byte_5D4594_251584[0].field_35 = a2
		}
		byte_5D4594_251584[0] = a2
	} else if libc.StrCmp(a1, CString("MATERIAL")) == 0 {
		a2.field_35 = nil
		a2.field_34 = byte_5D4594_251584[1]
		if byte_5D4594_251584[1] != nil {
			byte_5D4594_251584[1].field_35 = a2
		}
		byte_5D4594_251584[1] = a2
	} else if libc.StrCmp(a1, CString("ENCHANTMENT")) == 0 {
		a2.field_35 = nil
		a2.field_34 = byte_5D4594_251584[2]
		if byte_5D4594_251584[2] != nil {
			byte_5D4594_251584[2] = a2
		}
		byte_5D4594_251584[2] = a2
	} else {
		return 0
	}
	return 1
}
func nox_xxx_parseWeaponDef_412D40(a1 int32, a2 *File, a3 *byte) int32 {
	var (
		v3 **byte
		v4 **byte
		v5 *byte
		v6 *uint8
		v7 int32
		v9 [256]byte
	)
	for {
		nox_xxx_parseString_409470(a2, (*uint8)(unsafe.Pointer(&v9[0])))
		if libc.StrCmp(&v9[0], CString("END")) == 0 {
			return 1
		}
		v3 = (**byte)(alloc.Calloc(1, 88))
		v4 = v3
		if v3 == nil {
			return 0
		}
		*(**byte)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof((*byte)(nil))*21)) = nil
		*(**byte)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof((*byte)(nil))*20)) = *(**byte)(unsafe.Pointer(&dword_5d4594_251600))
		if dword_5d4594_251600 != 0 {
			*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_251600 + 84))) = uint32(uintptr(unsafe.Pointer(v3)))
		}
		*memmap.PtrUint32(6112660, 251604)++
		dword_5d4594_251600 = uint32(uintptr(unsafe.Pointer(v3)))
		v5 = (*byte)(alloc.Calloc(libc.StrLen(&v9[0])+1, 1))
		*v4 = v5
		if v5 == nil {
			return 0
		}
		libc.StrCpy(v5, &v9[0])
		for {
			nox_xxx_parseString_409470(a2, (*uint8)(unsafe.Pointer(&v9[0])))
			if libc.StrCmp(&v9[0], CString("END")) == 0 {
				break
			}
			v6 = (*uint8)(memmap.PtrOff(0x587000, 28600))
			if *memmap.PtrUint32(0x587000, 28604) != 0 {
				for libc.StrCmp(*(**byte)(unsafe.Pointer(v6)), &v9[0]) != 0 {
					v7 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v6))), unsafe.Sizeof(uint32(0))*3))))
					v6 = (*uint8)(unsafe.Add(unsafe.Pointer(v6), 8))
					if v7 == 0 {
						goto LABEL_13
					}
				}
				nox_xxx_parseRead_4093E0(a2, a3, 0x40000)
				if (*((*func(*byte, *byte, **byte) int32)(unsafe.Add(unsafe.Pointer((*func(*byte, *byte, **byte) int32)(unsafe.Pointer(v6))), unsafe.Sizeof(uintptr(0))*1))))(&v9[0], a3, v4) == 0 {
					return 0
				}
			}
		LABEL_13:
			if *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v6))), unsafe.Sizeof(uint32(0))*1))) == 0 {
				return 0
			}
		}
	}
}
func sub_412ED0(a1 int32, a2 *File, a3 *byte) int32 {
	var (
		v3 **byte
		v4 **byte
		v5 *byte
		v6 *uint8
		v7 int32
		v9 [256]byte
	)
	for {
		nox_xxx_parseString_409470(a2, (*uint8)(unsafe.Pointer(&v9[0])))
		if libc.StrCmp(&v9[0], CString("END")) == 0 {
			return 1
		}
		v3 = (**byte)(alloc.Calloc(1, 88))
		v4 = v3
		if v3 == nil {
			return 0
		}
		*(**byte)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof((*byte)(nil))*21)) = nil
		*(**byte)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof((*byte)(nil))*20)) = *(**byte)(unsafe.Pointer(&dword_5d4594_251608))
		if dword_5d4594_251608 != 0 {
			*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_251608 + 84))) = uint32(uintptr(unsafe.Pointer(v3)))
		}
		*memmap.PtrUint32(6112660, 251612)++
		dword_5d4594_251608 = uint32(uintptr(unsafe.Pointer(v3)))
		v5 = (*byte)(alloc.Calloc(libc.StrLen(&v9[0])+1, 1))
		*v4 = v5
		if v5 == nil {
			return 0
		}
		libc.StrCpy(v5, &v9[0])
		for {
			nox_xxx_parseString_409470(a2, (*uint8)(unsafe.Pointer(&v9[0])))
			if libc.StrCmp(&v9[0], CString("END")) == 0 {
				break
			}
			v6 = (*uint8)(memmap.PtrOff(0x587000, 28600))
			if *memmap.PtrUint32(0x587000, 28604) != 0 {
				for libc.StrCmp(*(**byte)(unsafe.Pointer(v6)), &v9[0]) != 0 {
					v7 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v6))), unsafe.Sizeof(uint32(0))*3))))
					v6 = (*uint8)(unsafe.Add(unsafe.Pointer(v6), 8))
					if v7 == 0 {
						goto LABEL_13
					}
				}
				nox_xxx_parseRead_4093E0(a2, a3, 0x40000)
				if (*((*func(*byte, *byte, **byte) int32)(unsafe.Add(unsafe.Pointer((*func(*byte, *byte, **byte) int32)(unsafe.Pointer(v6))), unsafe.Sizeof(uintptr(0))*1))))(&v9[0], a3, v4) == 0 {
					return 0
				}
			}
		LABEL_13:
			if *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v6))), unsafe.Sizeof(uint32(0))*1))) == 0 {
				return 0
			}
		}
	}
}
func nox_xxx_freeWeaponArmorDefAndModifs_413060() {
	sub_4130C0(*(*unsafe.Pointer)(unsafe.Pointer(&dword_5d4594_251600)))
	dword_5d4594_251600 = 0
	*memmap.PtrUint32(6112660, 251604) = 0
	sub_413100(*(*unsafe.Pointer)(unsafe.Pointer(&dword_5d4594_251608)))
	dword_5d4594_251608 = 0
	*memmap.PtrUint32(6112660, 251612) = 0
	for i := int32(0); i < 3; i++ {
		nox_xxx_modifFreeOne_413140(unsafe.Pointer(byte_5D4594_251584[i]))
		byte_5D4594_251584[i] = nil
	}
	byte_5D4594_251596 = 0
}
func sub_4130C0(lpMem unsafe.Pointer) {
	var (
		v1 unsafe.Pointer
		v2 unsafe.Pointer
	)
	v1 = lpMem
	if lpMem != nil {
		for {
			v2 = unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(v1)), unsafe.Sizeof(uint32(0))*20)))))
			if *(*uint32)(v1) != 0 {
				*(*unsafe.Pointer)(v1) = nil
			}
			if *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(v1)), unsafe.Sizeof(uint32(0))*2))) != 0 {
				*((*unsafe.Pointer)(unsafe.Add(unsafe.Pointer((*unsafe.Pointer)(v1)), unsafe.Sizeof(unsafe.Pointer(nil))*2))) = nil
			}
			v1 = nil
			v1 = v2
			if v2 == nil {
				break
			}
		}
	}
}
func sub_413100(lpMem unsafe.Pointer) {
	var (
		v1 unsafe.Pointer
		v2 unsafe.Pointer
	)
	v1 = lpMem
	if lpMem != nil {
		for {
			v2 = unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(v1)), unsafe.Sizeof(uint32(0))*20)))))
			if *(*uint32)(v1) != 0 {
				*(*unsafe.Pointer)(v1) = nil
			}
			if *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(v1)), unsafe.Sizeof(uint32(0))*2))) != 0 {
				*((*unsafe.Pointer)(unsafe.Add(unsafe.Pointer((*unsafe.Pointer)(v1)), unsafe.Sizeof(unsafe.Pointer(nil))*2))) = nil
			}
			v1 = nil
			v1 = v2
			if v2 == nil {
				break
			}
		}
	}
}
func nox_xxx_modifFreeOne_413140(lpMem unsafe.Pointer) {
	var (
		v1 unsafe.Pointer
		v2 unsafe.Pointer
	)
	v1 = lpMem
	if lpMem != nil {
		for {
			v2 = unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(v1)), unsafe.Sizeof(uint32(0))*34)))))
			if *(*uint32)(v1) != 0 {
				*(*unsafe.Pointer)(v1) = nil
			}
			if *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(v1)), unsafe.Sizeof(uint32(0))*2))) != 0 {
				*((*unsafe.Pointer)(unsafe.Add(unsafe.Pointer((*unsafe.Pointer)(v1)), unsafe.Sizeof(unsafe.Pointer(nil))*2))) = nil
			}
			if *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(v1)), unsafe.Sizeof(uint32(0))*3))) != 0 {
				*((*unsafe.Pointer)(unsafe.Add(unsafe.Pointer((*unsafe.Pointer)(v1)), unsafe.Sizeof(unsafe.Pointer(nil))*3))) = nil
			}
			v1 = nil
			v1 = v2
			if v2 == nil {
				break
			}
		}
	}
}
func nox_xxx_getProjectileClassById_413250(a1 int32) *uint32 {
	var result *uint32
	result = *(**uint32)(unsafe.Pointer(&dword_5d4594_251600))
	if dword_5d4594_251600 == 0 {
		return nil
	}
	for *(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*1)) != uint32(a1) {
		result = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*20)))))
		if result == nil {
			return nil
		}
	}
	return result
}
func nox_xxx_equipClothFindDefByTT_413270(a1 int32) *uint32 {
	var result *uint32
	result = *(**uint32)(unsafe.Pointer(&dword_5d4594_251608))
	if dword_5d4594_251608 == 0 {
		return nil
	}
	for *(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*1)) != uint32(a1) {
		result = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*20)))))
		if result == nil {
			return nil
		}
	}
	return result
}
func nox_xxx_modifGetIdByName_413290(a1 *byte) int32 {
	if a1 == nil {
		return math.MaxUint8
	}
	for i := int32(0); i < 3; i++ {
		for it := (*obj_412ae0_t)(byte_5D4594_251584[i]); it != nil; it = it.field_34 {
			if libc.StrCmp(it.field_0, a1) == 0 {
				return int32(it.field_1)
			}
		}
	}
	return math.MaxUint8
}
func sub_413300(a1 int32) int32 {
	for i := int32(0); i < 3; i++ {
		for it := (*obj_412ae0_t)(byte_5D4594_251584[i]); it != nil; it = it.field_34 {
			if it.field_1 == uint32(a1) {
				return int32(uintptr(unsafe.Pointer(it.field_0)))
			}
		}
	}
	return 0
}
func nox_xxx_modifGetDescById_413330(a1 int32) *obj_412ae0_t {
	if a1 == math.MaxUint8 {
		return nil
	}
	for i := int32(0); i < 3; i++ {
		for it := (*obj_412ae0_t)(byte_5D4594_251584[i]); it != nil; it = it.field_34 {
			if it.field_1 == uint32(a1) {
				return it
			}
		}
	}
	return nil
}
func sub_413370() unsafe.Pointer {
	return *(*unsafe.Pointer)(unsafe.Pointer(&dword_5d4594_251600))
}
func sub_413380(a1 int32) int32 {
	return int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 80))))
}
func sub_413390() unsafe.Pointer {
	return *(*unsafe.Pointer)(unsafe.Pointer(&dword_5d4594_251608))
}
func sub_4133A0(a1 int32) int32 {
	return int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 80))))
}
func nox_xxx_modifGetModifListByType_4133B0(a1 int32) *obj_412ae0_t {
	return byte_5D4594_251584[a1]
}
func nox_xxx_modifNext_4133C0(a1 *obj_412ae0_t) *obj_412ae0_t {
	return a1.field_34
}
func sub_4133D0(a1 int32) int32 {
	var (
		v1 int32
		v2 int32
		v3 int32
	)
	v1 = int32(*memmap.PtrUint32(6112660, 251620))
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 692))))
	if *memmap.PtrUint32(6112660, 251620) == 0 {
		v3 = nox_xxx_modifGetIdByName_413290(CString("Material7"))
		v1 = int32(uintptr(unsafe.Pointer(nox_xxx_modifGetDescById_413330(v3))))
		*memmap.PtrUint32(6112660, 251620) = uint32(v1)
	}
	return bool2int(*(*uint32)(unsafe.Pointer(uintptr(a1 + 8)))&0x13001000 != 0 && *(*uint32)(unsafe.Pointer(uintptr(v2 + 4))) == uint32(v1))
}
func sub_413420(a1 int8) int32 {
	var (
		v1 *uint8
		v2 int32
		v3 *uint8
	)
	if *memmap.PtrUint32(6112660, 251624) == 0 {
		v1 = (*uint8)(memmap.PtrOff(0x587000, 27340))
		for {
			*(*uint32)(unsafe.Pointer(v1)) = uint32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg(*((**byte)(unsafe.Add(unsafe.Pointer((**byte)(unsafe.Pointer(v1))), -int(unsafe.Sizeof((*byte)(nil))*1))))))))
			v1 = (*uint8)(unsafe.Add(unsafe.Pointer(v1), 20))
			if int32(uintptr(unsafe.Pointer(v1))) >= int32(uintptr(memmap.PtrOff(0x587000, 27460))) {
				break
			}
		}
		*memmap.PtrUint32(6112660, 251624) = 1
	}
	v2 = 0
	v3 = (*uint8)(memmap.PtrOff(0x587000, 27332))
	for int32(*v3) != int32(a1) {
		v3 = (*uint8)(unsafe.Add(unsafe.Pointer(v3), 20))
		v2++
		if int32(uintptr(unsafe.Pointer(v3))) >= int32(uintptr(memmap.PtrOff(0x587000, 27452))) {
			return 0
		}
	}
	return int32(*memmap.PtrUint32(0x587000, uint32(v2*20+27340)))
}
func sub_4134D0() int32 {
	dword_5d4594_251708 = 0
	dword_5d4594_251712 = 0
	dword_5d4594_251716 = 0
	dword_5d4594_251720 = 0
	return 0
}
func sub_4137A0() int32 {
	var result int32
	result = nox_xxx_checkGameFlagPause_413A50()
	if result == 0 {
		dword_5d4594_251716 = 1
	}
	return result
}
func sub_4137C0() int32 {
	var result int32
	result = int32(dword_5d4594_251720)
	if dword_5d4594_251720 == 0 {
		dword_5d4594_251720 = 1
		result = sub_43DBD0()
	}
	return result
}
func sub_4137E0() int32 {
	var result int32
	result = int32(dword_5d4594_251720)
	if dword_5d4594_251720 != 0 {
		dword_5d4594_251720 = 0
		result = sub_43DBE0()
	}
	return result
}
func sub_413870(a1 int32) int32 {
	var result int32
	dword_5d4594_251708 = 1
	result = sub_423CC0(a1)
	dword_5d4594_251712 = uint32(result)
	return result
}
func sub_413890() *byte {
	var (
		v0 *uint8
		v1 int32
		v2 int16
		v3 uint8
	)
	*memmap.PtrUint8(6112660, 251636) = sub_423EC0()
	*memmap.PtrUint8(6112660, 251637) = 0
	if int32(*memmap.PtrUint8(6112660, 251636)) == 0 {
		return nil
	}
	v0 = (*uint8)(memmap.PtrOff(6112660, uint32(libc.StrLen((*byte)(memmap.PtrOff(6112660, 251636)))+0x3D6F5)))
	v1 = int32(*memmap.PtrUint32(0x587000, 32316))
	v2 = int16(*memmap.PtrUint16(0x587000, 32320))
	*(*uint32)(unsafe.Pointer(func() *uint8 {
		p := &v0
		*p = (*uint8)(unsafe.Add(unsafe.Pointer(*p), -1))
		return *p
	}())) = *memmap.PtrUint32(0x587000, 32312)
	v3 = *memmap.PtrUint8(0x587000, 32322)
	*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v0))), unsafe.Sizeof(uint32(0))*1))) = uint32(v1)
	*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v0))), unsafe.Sizeof(uint16(0))*4))) = uint16(v2)
	*(*uint8)(unsafe.Add(unsafe.Pointer(v0), 10)) = v3
	return (*byte)(memmap.PtrOff(6112660, 251636))
}
func sub_4138E0(a1 int32) {
	*memmap.PtrUint32(6112660, 251740) = uint32(nox_xxx_checkGameFlagPause_413A50())
	sub_413A00(1)
}
func sub_413900(a1 int32) {
	if sub_44E0D0() == 0 {
		if *memmap.PtrUint32(6112660, 251740) == 0 {
			sub_413A00(0)
		}
	}
}
func sub_413920() int32 {
	sub_42EBB0(1, sub_413900, 0, CString("Pause"))
	sub_42EBB0(2, sub_4138E0, 0, CString("Pause"))
	dword_5d4594_251744 = 0
	return 1
}
func sub_413960() {
	dword_5d4594_251744 = 0
	sub_413A00(0)
}
func sub_4139B0() int32 {
	return bool2int(dword_5d4594_251744 != 0)
}
func sub_413E10() int32 {
	var result int32
	for result = int32(*memmap.PtrUint32(6112660, 338304)); result != 0; result = int32(*(*uint32)(unsafe.Pointer(uintptr(result + 4)))) {
	}
	return result
}
func nox_xxx_gameLoopMemDump_413E30() {
	var (
		v0 int32
		v1 int32
	)
	_ = v1
	var v2 *uint8
	var v3 int32
	var v4 int32
	var v5 *byte
	var v6 *uint8
	var v7 int32
	var v8 int32
	var v9 int32
	v0 = 0
	v1 = 0
	v9 = 0
	v2 = (*uint8)(memmap.PtrOff(6112660, 252364))
	for {
		*(*uint32)(unsafe.Pointer(v2)) = 0
		v2 = (*uint8)(unsafe.Add(unsafe.Pointer(v2), 84))
		if int32(uintptr(unsafe.Pointer(v2))) >= int32(uintptr(memmap.PtrOff(6112660, 338380))) {
			break
		}
	}
	v3 = int32(*memmap.PtrUint32(6112660, 338304))
	if *memmap.PtrUint32(6112660, 338304) != 0 {
		for {
			v4 = 0
			if v0 > 0 {
				v5 = (*byte)(memmap.PtrOff(6112660, 252284))
				for libc.StrCmp(v5, (*byte)(unsafe.Pointer(uintptr(v3+20)))) != 0 {
					v0 = v9
					v4++
					v5 = (*byte)(unsafe.Add(unsafe.Pointer(v5), 84))
					if v4 >= v9 {
						goto LABEL_10
					}
				}
				v0 = v9
				*memmap.PtrUint32(6112660, uint32(v4*84)+0x3D9CC) += *(*uint32)(unsafe.Pointer(uintptr(v3 + 16)))
			}
		LABEL_10:
			if v4 == v0 {
				*memmap.PtrUint32(6112660, uint32(v4*84)+0x3D9CC) = *(*uint32)(unsafe.Pointer(uintptr(v3 + 16)))
				libc.StrCpy((*byte)(memmap.PtrOff(6112660, uint32(v4*84)+0x3D97C)), (*byte)(unsafe.Pointer(uintptr(v3+20))))
				v9 = func() int32 {
					p := &v0
					*p++
					return *p
				}()
			}
			v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(v3 + 4))))
			if v3 == 0 {
				break
			}
		}
		v1 = 0
	}
	libc.Sort(memmap.PtrOff(6112660, 252284), uint32(v0), 84, sub_413F60)
	if v0 > 0 {
		v6 = (*uint8)(memmap.PtrOff(6112660, 252364))
		v7 = v0
		for {
			v8 = int32(*(*uint32)(unsafe.Pointer(v6)))
			v6 = (*uint8)(unsafe.Add(unsafe.Pointer(v6), 84))
			v1 += v8
			v7--
			if v7 == 0 {
				break
			}
		}
	}
}
func sub_413F60(a1 unsafe.Pointer, a2 unsafe.Pointer) int32 {
	return int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(a1)), unsafe.Sizeof(uint32(0))*20))) - *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(a2)), unsafe.Sizeof(uint32(0))*20))))
}
func nox_clone_str(a1 *byte) *byte {
	var result *byte
	result = (*byte)(alloc.Calloc(libc.StrLen(a1)+1, 1))
	if result != nil {
		libc.StrCpy(result, a1)
	}
	return result
}
func sub_414580() int32 {
	return int32(*memmap.PtrUint32(6112660, 338312))
}
func sub_414590() int32 {
	return int32(*memmap.PtrUint32(6112660, 338308))
}
func sub_4145A0() int32 {
	return int32(*memmap.PtrUint32(6112660, 338320))
}
func sub_4145B0() int32 {
	return int32(*memmap.PtrUint32(6112660, 338316))
}
func sub_4145C0(a1 uint32, a2 uint32) unsafe.Pointer {
	return alloc.Calloc(int(a1), int(a2))
}
func sub_4145E0(lpMem unsafe.Pointer) {
	lpMem = nil
}
func sub_414800() int32 {
	var (
		MultiByteStr [128]byte
		v2           [128]byte
		v3           [256]byte
		v4           [128]libc.WChar
		v5           [128]libc.WChar
	)
	if *memmap.PtrUint32(6112660, 338464) == 0 {
		return 0
	}
	nox_common_convertWideToMbString_414B00(*(**libc.WChar)(memmap.PtrOff(0x587000, 32584)), &MultiByteStr[0], 128)
	nox_common_convertWideToMbString_414B00(*(**libc.WChar)(memmap.PtrOff(0x587000, 32588)), &v3[0], 256)
	nox_wsprintfA(&v2[0], &MultiByteStr[0], *memmap.PtrUint32(6112660, 338464))
	nox_swprintf(&v5[0], libc.CWString("%S"), &v3[0])
	nox_swprintf(&v4[0], libc.CWString("%S"), &v2[0])
	nullsub_4(0, uint32(uintptr(unsafe.Pointer(&v5[0]))), uint32(uintptr(unsafe.Pointer(&v4[0]))), 0x40030)
	return 1
}
func sub_4148D0(lpFileName *byte) uint32 {
	var (
		v1 uint32
		v2 *File
		v3 *uint32
		i  *byte
		j  *uint8
		v6 int32
	)
	v1 = math.MaxUint32
	v2 = nox_fs_open(lpFileName)
	if uintptr(unsafe.Pointer(v2)) == uintptr(math.MaxUint32) {
		return math.MaxUint32
	}
	v3 = (*uint32)(alloc.Calloc(1, 8192))
	var n int32 = nox_fs_fread(v2, unsafe.Pointer(v3), 8192)
	*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*10)) = 0
	for i = (*byte)(unsafe.Pointer(uintptr(n))); n != 0; i = (*byte)(unsafe.Pointer(uintptr(n))) {
		for j = (*uint8)(unsafe.Pointer(v3)); i != nil; n = int32(uintptr(unsafe.Pointer(i))) {
			v6 = int32(*func() *uint8 {
				p := &j
				x := *p
				*p = (*uint8)(unsafe.Add(unsafe.Pointer(*p), 1))
				return x
			}())
			v1 = *memmap.PtrUint32(0x581450, uint32((v6^int32(uint8(v1)))*4+6160)) ^ v1>>8
			i = (*byte)(unsafe.Add(unsafe.Pointer(i), -1))
		}
		nox_fs_fread(v2, unsafe.Pointer(v3), 8192)
	}
	nox_fs_close(v2)
	alloc.Free(unsafe.Pointer(v3))
	return ^v1
}
func sub_4149A0(lpFileName *byte, lpBuffer unsafe.Pointer, a3 unsafe.Pointer, a4 unsafe.Pointer) int32 {
	var (
		v4 int32
		v5 *File
		v6 unsafe.Pointer
	)
	v4 = 0
	v5 = nox_fs_open(lpFileName)
	v6 = unsafe.Pointer(v5)
	if uintptr(unsafe.Pointer(v5)) == uintptr(math.MaxUint32) {
		return 0
	}
	var off int32 = nox_fs_fseek(v5, 40, stdio.SEEK_SET)
	if off > 0 && nox_fs_fread((*File)(v6), lpBuffer, 4) == 4 && nox_fs_fread((*File)(v6), a3, 4) == 4 {
		if nox_fs_fread((*File)(v6), a4, 4) == 4 {
			v4 = 1
		}
	}
	nox_fs_close((*File)(v6))
	return v4
}
func nox_common_getRegistryValue_414A40(lpSubKey *byte, lpValueName *byte) int32 {
	var v2 int32
	v2 = 0
	if compatRegOpenKeyExA((HKEY)(unsafe.Pointer(uintptr(1))), lpSubKey, 0, 0xF003F, PHKEY(unsafe.Pointer(&lpSubKey))) == 0 {
		if compatRegQueryValueExA(HKEY(unsafe.Pointer(lpSubKey)), lpValueName, nil, nil, nil, nil) == 0 {
			v2 = 1
		}
		compatRegCloseKey(HKEY(unsafe.Pointer(lpSubKey)))
	}
	return v2
}
func nox_common_setRegistryValue_414A90(lpSubKey *byte, lpValueName *byte, lpData *uint8) int32 {
	var (
		v3            int32
		dwDisposition uint32
	)
	v3 = 0
	if compatRegCreateKeyExA((HKEY)(unsafe.Pointer(uintptr(1))), lpSubKey, 0, nil, 0, 0xF003F, nil, PHKEY(unsafe.Pointer(&lpSubKey)), &dwDisposition) == 0 {
		if compatRegSetValueExA(HKEY(unsafe.Pointer(lpSubKey)), lpValueName, 0, 1, lpData, uint32(libc.StrLen((*byte)(unsafe.Pointer(lpData)))+1)) == 0 {
			v3 = 1
		}
		compatRegCloseKey(HKEY(unsafe.Pointer(lpSubKey)))
	}
	return v3
}
func nox_common_convertWideToMbString_414B00(lpWideCharStr *libc.WChar, lpMultiByteStr *byte, cbMultiByte int32) *byte {
	compatWideCharToMultiByte(0, 0, lpWideCharStr, -1, lpMultiByteStr, cbMultiByte, nil, nil)
	return compat_strrev(lpMultiByteStr)
}
func sub_414BA0(a1 int32) int32 {
	var v1 int32
	v1 = a1 + 4096
	if a1+4096 < 0 {
		return int32(*memmap.PtrUint32(6112660, 338472))
	}
	if v1 >= 8192 {
		v1 = 8191
	}
	return int32(*memmap.PtrUint32(6112660, uint32(v1*4)+0x52A28))
}
func sub_414BD0(a1 int32) int32 {
	var v1 int32
	v1 = a1
	if a1 < 0 {
		v1 = a1 + ((0x6487-a1)/0x6488)*0x6488
	}
	if v1 >= 0x6488 {
		v1 %= 0x6488
	}
	return int32(*memmap.PtrUint32(0x85B3FC, uint32(((v1<<12)/0x6488)*4+12260)))
}
func sub_414C50(a1 int32) int32 {
	var (
		v1     int32
		v2     int32
		result int32
	)
	v1 = sub_414BD0(a1)
	v2 = sub_414BD0(6434 - a1)
	if v2 != 0 {
		result = (v1 << 12) / v2
	} else {
		result = bool2int(v1 <= 0) + math.MaxInt32
	}
	return result
}
func nox_xxx_initSinCosTables_414C90() int8 {
	var (
		v0 int64
		v1 *uint8
		v2 *uint8
		v4 int32
		v5 int32
	)
	*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v0))), 0)) = *memmap.PtrUint8(6112660, 371240)
	if int32(*memmap.PtrUint8(6112660, 371240)) == 0 {
		*memmap.PtrUint8(6112660, 371240) = 1
		v4 = 0
		v1 = (*uint8)(memmap.PtrOff(0x85B3FC, 12260))
		for {
			*(*uint32)(unsafe.Pointer(v1)) = uint32(int32(int64(math.Sin(float64(v4)*0.0015339808) * *mem_getDoublePtr(0x581450, 7200))))
			v1 = (*uint8)(unsafe.Add(unsafe.Pointer(v1), 4))
			v4++
			if int32(uintptr(unsafe.Pointer(v1))) >= int32(uintptr(memmap.PtrOff(0x85B3FC, 28644))) {
				break
			}
		}
		v5 = 0
		v2 = (*uint8)(memmap.PtrOff(6112660, 338472))
		for {
			v0 = int64(math.Acos(float64(v5)*0.00024414062-1.0) * *mem_getDoublePtr(0x581450, 7184))
			*(*uint32)(unsafe.Pointer(v2)) = uint32(int32(v0))
			v2 = (*uint8)(unsafe.Add(unsafe.Pointer(v2), 4))
			v5++
			if int32(uintptr(unsafe.Pointer(v2))) >= int32(uintptr(memmap.PtrOff(6112660, 371240))) {
				break
			}
		}
	}
	return int8(v0)
}
func nox_thing_skip_AUD_414D40(f *nox_memfile) int32 {
	var (
		v2 int32
		v3 int32
		v5 uint8
	)
	v2 = f.ReadI32()
	if v2 <= 0 {
		return 1
	}
	v3 = v2
	for {
		v5 = f.ReadU8()
		f.Skip(int32(v5) + 9)
		for {
			v5 = f.ReadU8()
			if int32(v5) == 0 {
				break
			}
			f.Skip(int32(v5))
		}
		v3--
		if v3 == 0 {
			break
		}
	}
	return 1
}
func nox_thing_read_FLOR_414DB0(f *nox_memfile) int32 {
	var (
		v2  *uint8
		v3  *uint8
		v4  *uint8
		v5  int32
		v6  int32
		v7  *int32
		v8  int32
		v9  int32
		v10 *uint8
		v11 *int32
		v12 int32
		v14 uint8
		v15 uint8
	)
	v2 = (*uint8)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(f.cur), 4))))
	f.cur = (*byte)(unsafe.Pointer(v2))
	v3 = (*uint8)(unsafe.Add(unsafe.Pointer(v2), int32(*v2)+13))
	f.cur = (*byte)(unsafe.Pointer(v3))
	v4 = (*uint8)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(f.cur), 1))))
	v14 = *v3
	f.cur = (*byte)(unsafe.Pointer(v4))
	v15 = *v4
	f.cur = (*byte)(unsafe.Add(unsafe.Pointer(v4), 1))
	v4 = (*uint8)(unsafe.Add(unsafe.Pointer(v4), 3))
	v5 = int32(v14) * int32(v15)
	v14 = *((*uint8)(unsafe.Add(unsafe.Pointer(v4), -2)))
	f.cur = (*byte)(unsafe.Pointer(v4))
	v6 = v5 * int32(v14)
	if v6 > 0 {
		for {
			v7 = (*int32)(unsafe.Pointer(f.cur))
			v8 = *v7
			v9 = int32(uintptr(unsafe.Pointer((*int32)(unsafe.Add(unsafe.Pointer(v7), unsafe.Sizeof(int32(0))*1)))))
			f.cur = (*byte)(unsafe.Pointer(uintptr(v9)))
			if v8 == -1 {
				v10 = (*uint8)(unsafe.Pointer(uintptr(v9 + 1)))
				f.cur = (*byte)(unsafe.Pointer(v10))
				f.cur = (*byte)(unsafe.Add(unsafe.Pointer(f.cur), int32(*v10)+1))
			}
			v6--
			if v6 == 0 {
				break
			}
		}
	}
	v11 = (*int32)(unsafe.Pointer(f.cur))
	v12 = *v11
	f.cur = (*byte)(unsafe.Pointer((*int32)(unsafe.Add(unsafe.Pointer(v11), unsafe.Sizeof(int32(0))*1))))
	return bool2int(uint32(v12) == 0x454E4420)
}
func nox_thing_read_EDGE_414E70(a1 int32, a2 unsafe.Pointer) int32 {
	var (
		v2  int32
		v3  *uint8
		v4  *uint8
		v5  uint8
		v6  uint8
		v7  *uint8
		v9  uint8
		v10 *uint8
		v11 int32
		v12 *int32
		v13 int32
		v14 int32
		v15 *uint8
		v16 *int32
		v17 int32
		v18 int32
		v19 uint8
		v20 uint8
		v21 uint8
	)
	v2 = a1
	v18 = a1
	v3 = (*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 8))) + 4)))
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 8))) = uint32(uintptr(unsafe.Pointer(v3)))
	v20 = *v3
	*(*uint32)(unsafe.Pointer(uintptr(v2 + 8))) = uint32(uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v3), 1)))))
	nox_memfile_read(a2, 1, int32(v20), (*nox_memfile)(unsafe.Pointer(uintptr(v18))))
	v4 = (*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 8))) + 9)))
	*(*uint32)(unsafe.Pointer(uintptr(v2 + 8))) = uint32(uintptr(unsafe.Pointer(v4)))
	v5 = *v4
	v4 = (*uint8)(unsafe.Add(unsafe.Pointer(v4), 2))
	v19 = v5
	*(*uint32)(unsafe.Pointer(uintptr(v2 + 8))) = uint32(uintptr(unsafe.Pointer(v4)))
	v6 = *v4
	v7 = (*uint8)(unsafe.Add(unsafe.Pointer(v4), 1))
	*(*uint32)(unsafe.Pointer(uintptr(v2 + 8))) = uint32(uintptr(unsafe.Pointer(v7)))
	if int32(v6) == 1 {
		return 0
	}
	v9 = *v7
	v10 = (*uint8)(unsafe.Add(unsafe.Pointer(v7), 1))
	*(*uint32)(unsafe.Pointer(uintptr(v2 + 8))) = uint32(uintptr(unsafe.Pointer(v10)))
	v21 = *v10
	*(*uint32)(unsafe.Pointer(uintptr(v2 + 8))) = uint32(uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v10), 1)))))
	v11 = int32(v19) * 2 * (int32(v9) + int32(v21))
	if v11 > 0 {
		for {
			v12 = *(**int32)(unsafe.Pointer(uintptr(v2 + 8)))
			v13 = *v12
			v14 = int32(uintptr(unsafe.Pointer((*int32)(unsafe.Add(unsafe.Pointer(v12), unsafe.Sizeof(int32(0))*1)))))
			*(*uint32)(unsafe.Pointer(uintptr(v2 + 8))) = uint32(v14)
			if v13 == -1 {
				v15 = (*uint8)(unsafe.Pointer(uintptr(v14 + 1)))
				*(*uint32)(unsafe.Pointer(uintptr(v2 + 8))) = uint32(uintptr(unsafe.Pointer(v15)))
				*(*uint32)(unsafe.Pointer(uintptr(v2 + 8))) = uint32(uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v15), int32(*v15)+1)))))
			}
			v11--
			if v11 == 0 {
				break
			}
		}
	}
	v16 = *(**int32)(unsafe.Pointer(uintptr(v2 + 8)))
	v17 = *v16
	*(*uint32)(unsafe.Pointer(uintptr(v2 + 8))) = uint32(uintptr(unsafe.Pointer((*int32)(unsafe.Add(unsafe.Pointer(v16), unsafe.Sizeof(int32(0))*1)))))
	return bool2int(uint32(v17) == 0x454E4420)
}
func nox_thing_read_WALL_414F60(a1 *uint32, a2 unsafe.Pointer) int32 {
	var (
		v2     *uint32
		v3     unsafe.Pointer
		v4     *uint8
		v5     int32
		v6     *uint8
		v7     *uint8
		v8     *uint8
		v9     *uint8
		v10    int32
		v11    int32
		v12    int32
		v13    *int32
		v14    int32
		v15    int32
		v16    *uint8
		v17    *int32
		v18    int32
		result int32
		v20    *uint32
		v21    int32
		v22    uint8
		v23    uint8
		v24    uint8
		v25    uint8
		v26    uint8
	)
	v2 = a1
	v3 = a2
	v20 = a1
	v4 = (*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*2)) + 4)))
	*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*2)) = uint32(uintptr(unsafe.Pointer(v4)))
	v22 = *v4
	*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*2)) = uint32(uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v4), 1)))))
	nox_memfile_read(v3, 1, int32(v22), (*nox_memfile)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v20)))))))
	*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(v3)), v22))) = 0
	*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*2)) += 14
	nox_memfile_read64align_40AD60((*byte)(unsafe.Pointer(&a2)), 1, 1, (*nox_memfile)(unsafe.Pointer(v2)))
	v5 = 0
	if int32(uint8(uintptr(a2))) != 0 {
		for v5 < 8 {
			v6 = (*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*2)))))
			v23 = *v6
			*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*2)) = uint32(uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v6), 1)))))
			nox_memfile_read(v3, 1, int32(v23), (*nox_memfile)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v2)))))))
			*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(v3)), v23))) = 0
			if func() int32 {
				p := &v5
				*p++
				return *p
			}() >= int32(uint8(uintptr(a2))) {
				goto LABEL_4
			}
		}
		result = 0
	} else {
	LABEL_4:
		v7 = (*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*2)))))
		v24 = *v7
		*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*2)) = uint32(uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v7), 1)))))
		nox_memfile_read(v3, 1, int32(v24), (*nox_memfile)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v2)))))))
		*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(v3)), v24))) = 0
		v8 = (*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*2)))))
		v25 = *v8
		*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*2)) = uint32(uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v8), 1)))))
		nox_memfile_read(v3, 1, int32(v25), (*nox_memfile)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v2)))))))
		*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(v3)), v25))) = 0
		v9 = (*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*2)))))
		v26 = *v9
		*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*2)) = uint32(uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v9), 1)))))
		nox_memfile_read(v3, 1, int32(v26), (*nox_memfile)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v2)))))))
		*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(v3)), v26))) = 0
		*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*2))++
		v10 = 15
		for {
			nox_memfile_read64align_40AD60((*byte)(unsafe.Pointer(&v21)), 1, 1, (*nox_memfile)(unsafe.Pointer(v2)))
			if int32(uint8(int8(v21))) > 0 {
				v11 = int32(uint8(int8(v21)))
				for {
					v12 = 4
					for {
						v13 = (*int32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*2)) + 8)))
						*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*2)) = uint32(uintptr(unsafe.Pointer(v13)))
						v14 = *v13
						v15 = int32(uintptr(unsafe.Pointer((*int32)(unsafe.Add(unsafe.Pointer(v13), unsafe.Sizeof(int32(0))*1)))))
						*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*2)) = uint32(v15)
						if v14 == -1 {
							v16 = (*uint8)(unsafe.Pointer(uintptr(v15 + 1)))
							*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*2)) = uint32(uintptr(unsafe.Pointer(v16)))
							*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*2)) = uint32(uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v16), int32(*v16)+1)))))
						}
						v12--
						if v12 == 0 {
							break
						}
					}
					v11--
					if v11 == 0 {
						break
					}
				}
			}
			v10--
			if v10 == 0 {
				break
			}
		}
		v17 = (*int32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*2)))))
		v18 = *v17
		*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*2)) = uint32(uintptr(unsafe.Pointer((*int32)(unsafe.Add(unsafe.Pointer(v17), unsafe.Sizeof(int32(0))*1)))))
		result = bool2int(uint32(v18) == 0x454E4420)
	}
	return result
}
func nox_thing_skip_spells_415100(a1 int32) int32 {
	var (
		v1  *int32
		v2  int32
		v3  int32
		v4  *uint8
		v5  *int32
		v6  int32
		v7  int32
		v8  *uint8
		v9  *int32
		v10 int32
		v11 int32
		v12 *uint8
		v13 *uint8
		v14 *int16
		v15 *uint8
		v16 *uint8
		v17 *uint8
	)
	v1 = *(**int32)(unsafe.Pointer(uintptr(a1 + 8)))
	v2 = *v1
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 8))) = uint32(uintptr(unsafe.Pointer((*int32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(int32(0))*1)))))
	if v2 <= 0 {
		return 1
	}
	v3 = v2
	for {
		v4 = (*uint8)(unsafe.Pointer(uintptr(uint32(int32(**(**uint8)(unsafe.Pointer(uintptr(a1 + 8))))+4) + *(*uint32)(unsafe.Pointer(uintptr(a1 + 8))))))
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 8))) = uint32(uintptr(unsafe.Pointer(v4)))
		v5 = (*int32)(unsafe.Pointer(uintptr(uint32(int32(*v4)+1) + *(*uint32)(unsafe.Pointer(uintptr(a1 + 8))))))
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 8))) = uint32(uintptr(unsafe.Pointer(v5)))
		v6 = *v5
		v7 = int32(uintptr(unsafe.Pointer((*int32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(int32(0))*1)))))
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 8))) = uint32(v7)
		if v6 == -1 {
			v8 = (*uint8)(unsafe.Pointer(uintptr(v7 + 1)))
			*(*uint32)(unsafe.Pointer(uintptr(a1 + 8))) = uint32(uintptr(unsafe.Pointer(v8)))
			*(*uint32)(unsafe.Pointer(uintptr(a1 + 8))) += uint32(int32(*v8) + 1)
		}
		v9 = *(**int32)(unsafe.Pointer(uintptr(a1 + 8)))
		v10 = *v9
		v11 = int32(uintptr(unsafe.Pointer((*int32)(unsafe.Add(unsafe.Pointer(v9), unsafe.Sizeof(int32(0))*1)))))
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 8))) = uint32(v11)
		if v10 == -1 {
			v12 = (*uint8)(unsafe.Pointer(uintptr(v11 + 1)))
			*(*uint32)(unsafe.Pointer(uintptr(a1 + 8))) = uint32(uintptr(unsafe.Pointer(v12)))
			*(*uint32)(unsafe.Pointer(uintptr(a1 + 8))) += uint32(int32(*v12) + 1)
		}
		v13 = (*uint8)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 8))) + 4)))
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 8))) = uint32(uintptr(unsafe.Pointer(v13)))
		v14 = (*int16)(unsafe.Pointer(uintptr(uint32(int32(*v13)+1) + *(*uint32)(unsafe.Pointer(uintptr(a1 + 8))))))
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 8))) = uint32(uintptr(unsafe.Pointer(v14)))
		v15 = (*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(v14))), *v14))), 2))
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 8))) = uint32(uintptr(unsafe.Pointer(v15)))
		v16 = (*uint8)(unsafe.Pointer(uintptr(uint32(int32(*v15)+1) + *(*uint32)(unsafe.Pointer(uintptr(a1 + 8))))))
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 8))) = uint32(uintptr(unsafe.Pointer(v16)))
		v17 = (*uint8)(unsafe.Pointer(uintptr(uint32(int32(*v16)+1) + *(*uint32)(unsafe.Pointer(uintptr(a1 + 8))))))
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 8))) = uint32(uintptr(unsafe.Pointer(v17)))
		v3--
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 8))) += uint32(int32(*v17) + 1)
		if v3 == 0 {
			break
		}
	}
	return 1
}
func nox_thing_read_image_415240(a1 int32) int32 {
	var (
		v1     int32
		v2     *int32
		v3     int32
		v4     int32
		result int32
		v6     *uint8
		v7     *uint8
		v8     uint8
		v9     int32
		v10    *int32
		v11    int32
		v12    int32
		v13    *uint8
		v14    uint8
		v15    int8
	)
	v1 = a1
	v2 = *(**int32)(unsafe.Pointer(uintptr(a1 + 8)))
	v3 = *v2
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 8))) = uint32(uintptr(unsafe.Pointer((*int32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(int32(0))*1)))))
	if v3 > 0 {
		v4 = v3
		result = 1
		for {
			v6 = (*uint8)(unsafe.Pointer(uintptr(uint32(int32(**(**uint8)(unsafe.Pointer(uintptr(v1 + 8))))+1) + *(*uint32)(unsafe.Pointer(uintptr(v1 + 8))))))
			*(*uint32)(unsafe.Pointer(uintptr(v1 + 8))) = uint32(uintptr(unsafe.Pointer(v6)))
			v15 = int8(*v6)
			*(*uint32)(unsafe.Pointer(uintptr(v1 + 8))) = uint32(uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v6), 1)))))
			v7 = (*uint8)(unsafe.Add(unsafe.Pointer(v6), 1))
			if int32(v15) == 1 || int32(v15) != 2 {
				break
			}
			v8 = *v7
			*(*uint32)(unsafe.Pointer(uintptr(v1 + 8))) = uint32(uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v7), 2)))))
			v14 = v8
			*(*uint32)(unsafe.Pointer(uintptr(v1 + 8))) += uint32(int32(*(*uint8)(unsafe.Add(unsafe.Pointer(v7), 2))) + 1)
			if int32(v8) != 0 {
				goto LABEL_7
			}
		LABEL_11:
			if func() int32 {
				p := &v4
				*p--
				return *p
			}() == 0 {
				return result
			}
		}
		v14 = 1
	LABEL_7:
		v9 = int32(v14)
		for {
			v10 = *(**int32)(unsafe.Pointer(uintptr(v1 + 8)))
			v11 = *v10
			v12 = int32(uintptr(unsafe.Pointer((*int32)(unsafe.Add(unsafe.Pointer(v10), unsafe.Sizeof(int32(0))*1)))))
			*(*uint32)(unsafe.Pointer(uintptr(v1 + 8))) = uint32(v12)
			if v11 == -1 {
				v13 = (*uint8)(unsafe.Pointer(uintptr(v12 + 1)))
				*(*uint32)(unsafe.Pointer(uintptr(v1 + 8))) = uint32(uintptr(unsafe.Pointer(v13)))
				*(*uint32)(unsafe.Pointer(uintptr(v1 + 8))) += uint32(int32(*v13) + 1)
			}
			v9--
			if v9 == 0 {
				break
			}
		}
		goto LABEL_11
	}
	return 1
}
func nox_thing_read_ability_415320(a1 int32) int32 {
	var (
		v1  *int32
		v2  int32
		v3  int32
		v4  *int32
		v5  int32
		v6  int32
		v7  *uint8
		v8  *int32
		v9  int32
		v10 int32
		v11 *uint8
		v12 *int32
		v13 int32
		v14 int32
		v15 *uint8
		v16 *int16
		v17 *uint8
		v18 *uint8
		v19 *uint8
	)
	v1 = *(**int32)(unsafe.Pointer(uintptr(a1 + 8)))
	v2 = *v1
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 8))) = uint32(uintptr(unsafe.Pointer((*int32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(int32(0))*1)))))
	if v2 <= 0 {
		return 1
	}
	v3 = v2
	for {
		v4 = (*int32)(unsafe.Pointer(uintptr(uint32(int32(**(**uint8)(unsafe.Pointer(uintptr(a1 + 8))))+2) + *(*uint32)(unsafe.Pointer(uintptr(a1 + 8))))))
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 8))) = uint32(uintptr(unsafe.Pointer(v4)))
		v5 = *v4
		v6 = int32(uintptr(unsafe.Pointer((*int32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(int32(0))*1)))))
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 8))) = uint32(v6)
		if v5 == -1 {
			v7 = (*uint8)(unsafe.Pointer(uintptr(v6 + 1)))
			*(*uint32)(unsafe.Pointer(uintptr(a1 + 8))) = uint32(uintptr(unsafe.Pointer(v7)))
			*(*uint32)(unsafe.Pointer(uintptr(a1 + 8))) += uint32(int32(*v7) + 1)
		}
		v8 = *(**int32)(unsafe.Pointer(uintptr(a1 + 8)))
		v9 = *v8
		v10 = int32(uintptr(unsafe.Pointer((*int32)(unsafe.Add(unsafe.Pointer(v8), unsafe.Sizeof(int32(0))*1)))))
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 8))) = uint32(v10)
		if v9 == -1 {
			v11 = (*uint8)(unsafe.Pointer(uintptr(v10 + 1)))
			*(*uint32)(unsafe.Pointer(uintptr(a1 + 8))) = uint32(uintptr(unsafe.Pointer(v11)))
			*(*uint32)(unsafe.Pointer(uintptr(a1 + 8))) += uint32(int32(*v11) + 1)
		}
		v12 = *(**int32)(unsafe.Pointer(uintptr(a1 + 8)))
		v13 = *v12
		v14 = int32(uintptr(unsafe.Pointer((*int32)(unsafe.Add(unsafe.Pointer(v12), unsafe.Sizeof(int32(0))*1)))))
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 8))) = uint32(v14)
		if v13 == -1 {
			v15 = (*uint8)(unsafe.Pointer(uintptr(v14 + 1)))
			*(*uint32)(unsafe.Pointer(uintptr(a1 + 8))) = uint32(uintptr(unsafe.Pointer(v15)))
			*(*uint32)(unsafe.Pointer(uintptr(a1 + 8))) += uint32(int32(*v15) + 1)
		}
		v16 = (*int16)(unsafe.Pointer(uintptr(uint32(int32(**(**uint8)(unsafe.Pointer(uintptr(a1 + 8))))+1) + *(*uint32)(unsafe.Pointer(uintptr(a1 + 8))))))
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 8))) = uint32(uintptr(unsafe.Pointer(v16)))
		v17 = (*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(v16))), *v16))), 2))
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 8))) = uint32(uintptr(unsafe.Pointer(v17)))
		v18 = (*uint8)(unsafe.Pointer(uintptr(uint32(int32(*v17)+1) + *(*uint32)(unsafe.Pointer(uintptr(a1 + 8))))))
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 8))) = uint32(uintptr(unsafe.Pointer(v18)))
		v19 = (*uint8)(unsafe.Pointer(uintptr(uint32(int32(*v18)+1) + *(*uint32)(unsafe.Pointer(uintptr(a1 + 8))))))
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 8))) = uint32(uintptr(unsafe.Pointer(v19)))
		v3--
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 8))) += uint32(int32(*v19) + 1)
		if v3 == 0 {
			break
		}
	}
	return 1
}
func nox_thing_read_audio_415660(a1p *nox_memfile, a2 *byte) int32 {
	var (
		a1 int32 = int32(uintptr(unsafe.Pointer(a1p)))
		v2 int32
		v3 *int32
		v4 int32
	)
	v2 = 0
	v3 = *(**int32)(unsafe.Pointer(uintptr(a1 + 8)))
	v4 = *v3
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 8))) = uint32(uintptr(unsafe.Pointer((*int32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(int32(0))*1)))))
	if v4 <= 0 {
		return 1
	}
	for sub_452BD0(a1, a2) != 0 {
		if func() int32 {
			p := &v2
			*p++
			return *p
		}() >= v4 {
			return 1
		}
	}
	return 0
}
func nox_thing_read_ABIL_415750(f *nox_memfile, a2 unsafe.Pointer) int32 {
	var (
		v2 int32
		v4 int32
	)
	v2 = 0
	v4 = f.ReadI32()
	if v4 <= 0 {
		return 1
	}
	for nox_thing_read_ABIL_rec_424F00(f, a2) != 0 {
		if func() int32 {
			p := &v2
			*p++
			return *p
		}() >= v4 {
			return 1
		}
	}
	return 0
}
func nox_xxx_weaponInventoryEquipFlags_415820(item *nox_object_t) int32 {
	if item == nil {
		return 0
	}
	return nox_xxx_ammoCheck_415880(item.typ_ind)
}
func sub_415840(a1 *byte) int32 {
	var (
		v1 int32
		i  *uint8
		v3 int32
	)
	v1 = 0
	if *memmap.PtrUint32(0x587000, 33064) == 0 {
		return 0
	}
	for i = (*uint8)(memmap.PtrOff(0x587000, 33064)); *((**byte)(unsafe.Add(unsafe.Pointer((**byte)(unsafe.Pointer(i))), unsafe.Sizeof((*byte)(nil))*2))) != a1; i = (*uint8)(unsafe.Add(unsafe.Pointer(i), 12)) {
		v3 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(i))), unsafe.Sizeof(uint32(0))*3))))
		v1++
		if v3 == 0 {
			return 0
		}
	}
	return int32(*memmap.PtrUint32(0x587000, uint32(v1*12+0x812C)))
}
func nox_xxx_ammoCheck_415880(typ_ind uint16) int32 {
	var (
		v1 int32
		i  *uint8
		v3 int32
	)
	v1 = 0
	if *memmap.PtrUint32(0x587000, 33064) == 0 {
		return 0
	}
	for i = (*uint8)(memmap.PtrOff(0x587000, 33064)); unsafe.Pointer(*((**byte)(unsafe.Add(unsafe.Pointer((**byte)(unsafe.Pointer(i))), unsafe.Sizeof((*byte)(nil))*1)))) != unsafe.Pointer(uintptr(typ_ind)); i = (*uint8)(unsafe.Add(unsafe.Pointer(i), 12)) {
		v3 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(i))), unsafe.Sizeof(uint32(0))*3))))
		v1++
		if v3 == 0 {
			return 0
		}
	}
	return int32(*memmap.PtrUint32(0x587000, uint32(v1*12+0x8130)))
}
func sub_415910(a1 *byte) int32 {
	var (
		v1 int32
		v2 **byte
		v3 *uint8
		v4 int32
	)
	v1 = 0
	if *memmap.PtrUint32(0x587000, 33064) == 0 {
		return 0
	}
	v2 = (**byte)(memmap.PtrOff(0x587000, 33064))
	v3 = (*uint8)(memmap.PtrOff(0x587000, 33064))
	for libc.StrCaseCmp(a1, *v2) != 0 {
		v4 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v3))), unsafe.Sizeof(uint32(0))*3))))
		v3 = (*uint8)(unsafe.Add(unsafe.Pointer(v3), 12))
		v1++
		v2 = (**byte)(unsafe.Pointer(v3))
		if v4 == 0 {
			return 0
		}
	}
	return int32(*memmap.PtrUint32(0x587000, uint32(v1*12+0x8130)))
}
func sub_415960(a1 *libc.WChar) int32 {
	var (
		v1 int32
		v2 **libc.WChar
		v3 *uint8
		v4 int32
	)
	v1 = 0
	if *memmap.PtrUint32(0x587000, 33392) == 0 {
		return 0
	}
	v2 = (**libc.WChar)(memmap.PtrOff(0x587000, 33392))
	v3 = (*uint8)(memmap.PtrOff(0x587000, 33392))
	for _nox_wcsicmp(a1, *v2) != 0 {
		v4 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v3))), unsafe.Sizeof(uint32(0))*3))))
		v3 = (*uint8)(unsafe.Add(unsafe.Pointer(v3), 12))
		v1++
		v2 = (**libc.WChar)(unsafe.Pointer(v3))
		if v4 == 0 {
			return 0
		}
	}
	return int32(*memmap.PtrUint32(0x587000, uint32(v1*12+33400)))
}
func sub_4159B0(a1 *byte) *byte {
	var (
		v1 int32
		i  *uint8
		v3 int32
	)
	v1 = 0
	if *memmap.PtrUint32(0x587000, 33064) == 0 {
		return nil
	}
	for i = (*uint8)(memmap.PtrOff(0x587000, 33064)); *((**byte)(unsafe.Add(unsafe.Pointer((**byte)(unsafe.Pointer(i))), unsafe.Sizeof((*byte)(nil))*2))) != a1; i = (*uint8)(unsafe.Add(unsafe.Pointer(i), 12)) {
		v3 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(i))), unsafe.Sizeof(uint32(0))*3))))
		v1++
		if v3 == 0 {
			return nil
		}
	}
	return *(**byte)(memmap.PtrOff(0x587000, uint32(v1*12+0x8128)))
}
func sub_4159F0(a1 int32) int32 {
	var (
		v1 int32
		i  *uint8
		v3 int32
	)
	v1 = 0
	if *memmap.PtrUint32(0x587000, 33392) == 0 {
		return 0
	}
	for i = (*uint8)(memmap.PtrOff(0x587000, 33392)); *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(i))), unsafe.Sizeof(uint32(0))*2))) != uint32(a1); i = (*uint8)(unsafe.Add(unsafe.Pointer(i), 12)) {
		v3 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(i))), unsafe.Sizeof(uint32(0))*3))))
		v1++
		if v3 == 0 {
			return 0
		}
	}
	return int32(*memmap.PtrUint32(0x587000, uint32(v1*12+0x8270)))
}
func sub_415A30(a1 *byte) int32 {
	var (
		v1     *byte
		v2     int32
		result int32
	)
	v1 = (*byte)(unsafe.Pointer(uintptr(sub_415910(a1))))
	if v1 != nil && (func() int32 {
		v2 = sub_415840(v1)
		return v2
	}()) != 0 {
		result = int32(uintptr(unsafe.Pointer(nox_xxx_objectTypeByInd_4E3B70(v2))))
	} else {
		result = 0
	}
	return result
}
func sub_415B20(a1 *byte) int32 {
	var (
		v1 int32
		i  *uint8
		v3 int32
	)
	v1 = 0
	if *memmap.PtrUint32(0x587000, 34848) == 0 {
		return 0
	}
	for i = (*uint8)(memmap.PtrOff(0x587000, 34848)); *((**byte)(unsafe.Add(unsafe.Pointer((**byte)(unsafe.Pointer(i))), unsafe.Sizeof((*byte)(nil))*2))) != a1; i = (*uint8)(unsafe.Add(unsafe.Pointer(i), 24)) {
		v3 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(i))), unsafe.Sizeof(uint32(0))*6))))
		v1++
		if v3 == 0 {
			return 0
		}
	}
	return int32(*memmap.PtrUint32(0x587000, uint32(v1*24+0x8834)))
}
func sub_415BD0(a1 int32) float64 {
	var (
		v1     *float32
		result float64
	)
	if *(*uint32)(unsafe.Pointer(uintptr(a1 + 8)))&0x2000000 != 0 && (func() *float32 {
		v1 = (*float32)(unsafe.Pointer(nox_xxx_equipClothFindDefByTT_413270(int32(*(*uint16)(unsafe.Pointer(uintptr(a1 + 4)))))))
		return v1
	}()) != nil {
		result = float64(*(*float32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(float32(0))*16)))
	} else {
		result = 0.0
	}
	return result
}
func nox_xxx_itemApplyDefendEffect_415C00(a1 int32) float64 {
	var (
		v1 *int32
		v2 *float32
		v4 int32
		v5 func(int32, int32, uint32, int32, uint32, *float32)
		v6 float32
	)
	v6 = 0.0
	if (*(*uint32)(unsafe.Pointer(uintptr(a1 + 8))) & 0x2000000) == 0 {
		return 0.0
	}
	v1 = *(**int32)(unsafe.Pointer(uintptr(a1 + 692)))
	v2 = (*float32)(unsafe.Pointer(nox_xxx_equipClothFindDefByTT_413270(int32(*(*uint16)(unsafe.Pointer(uintptr(a1 + 4)))))))
	if v2 == nil {
		return float64(v6)
	}
	v6 = *(*float32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(float32(0))*16))
	v4 = *v1
	if *v1 != 0 {
		v5 = cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(v4 + 76))), (*func(int32, int32, uint32, int32, uint32, *float32))(nil))
		if v5 != nil {
			v5(v4, a1, 0, a1, 0, &v6)
		}
	}
	return float64(v6)
}
func nox_xxx_unitArmorInventoryEquipFlags_415C70(item *nox_object_t) int32 {
	var (
		v1     int32
		result int32
	)
	v1 = sub_415C90(item)
	if v1 < 0 {
		result = 0
	} else {
		result = int32(*memmap.PtrUint32(0x587000, uint32(v1*24+34860)))
	}
	return result
}
func sub_415C90(item *nox_object_t) int32 {
	var (
		result int32
		v2     int32
		i      *uint8
		v4     int32
	)
	if item == nil {
		return -1
	}
	result = 0
	if *memmap.PtrUint32(0x587000, 34848) == 0 {
		return -1
	}
	*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v2))), unsafe.Sizeof(uint16(0))*1)) = 0
	for i = (*uint8)(memmap.PtrOff(0x587000, 34848)); ; i = (*uint8)(unsafe.Add(unsafe.Pointer(i), 24)) {
		*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v2))), unsafe.Sizeof(uint16(0))*0)) = item.typ_ind
		if uint32(v2) == *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(i))), unsafe.Sizeof(uint32(0))*2))) {
			break
		}
		v4 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(i))), unsafe.Sizeof(uint32(0))*6))))
		result++
		if v4 == 0 {
			return -1
		}
	}
	return result
}
func sub_415CD0(a1 *byte) int32 {
	var (
		v1 int32
		i  *uint8
		v3 int32
	)
	v1 = 0
	if *memmap.PtrUint32(0x587000, 34848) == 0 {
		return 0
	}
	for i = (*uint8)(memmap.PtrOff(0x587000, 34848)); a1 != *((**byte)(unsafe.Add(unsafe.Pointer((**byte)(unsafe.Pointer(i))), unsafe.Sizeof((*byte)(nil))*3))); i = (*uint8)(unsafe.Add(unsafe.Pointer(i), 24)) {
		v3 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(i))), unsafe.Sizeof(uint32(0))*6))))
		v1++
		if v3 == 0 {
			return 0
		}
	}
	return int32(*memmap.PtrUint32(0x587000, uint32(v1*24+0x8828)))
}
func sub_415D10(a1 int32) int32 {
	var (
		v1 int32
		i  *uint8
		v3 int32
	)
	v1 = 0
	if *memmap.PtrUint32(0x587000, 34848) == 0 {
		return 0
	}
	for i = (*uint8)(memmap.PtrOff(0x587000, 34848)); unsafe.Pointer(uintptr(a1)) != unsafe.Pointer(*((**byte)(unsafe.Add(unsafe.Pointer((**byte)(unsafe.Pointer(i))), unsafe.Sizeof((*byte)(nil))*2)))); i = (*uint8)(unsafe.Add(unsafe.Pointer(i), 24)) {
		v3 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(i))), unsafe.Sizeof(uint32(0))*6))))
		v1++
		if v3 == 0 {
			return 0
		}
	}
	return int32(*memmap.PtrUint32(0x587000, uint32(v1*24+34860)))
}
func sub_415DA0(a1 *libc.WChar) int32 {
	var (
		v1 int32
		v2 **libc.WChar
		v3 *uint8
		v4 int32
	)
	v1 = 0
	if *memmap.PtrUint32(0x587000, 35496) == 0 {
		return 0
	}
	v2 = (**libc.WChar)(memmap.PtrOff(0x587000, 35496))
	v3 = (*uint8)(memmap.PtrOff(0x587000, 35496))
	for _nox_wcsicmp(a1, *v2) != 0 {
		v4 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v3))), unsafe.Sizeof(uint32(0))*3))))
		v3 = (*uint8)(unsafe.Add(unsafe.Pointer(v3), 12))
		v1++
		v2 = (**libc.WChar)(unsafe.Pointer(v3))
		if v4 == 0 {
			return 0
		}
	}
	return int32(*memmap.PtrUint32(0x587000, uint32(v1*12+0x8AB0)))
}
func sub_415DF0(a1 *byte) int32 {
	var (
		v1 int32
		v2 **byte
		v3 *uint8
		v4 int32
	)
	v1 = 0
	if *memmap.PtrUint32(0x587000, 34848) == 0 {
		return 0
	}
	v2 = (**byte)(memmap.PtrOff(0x587000, 34848))
	v3 = (*uint8)(memmap.PtrOff(0x587000, 34848))
	for libc.StrCaseCmp(a1, *v2) != 0 {
		v4 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v3))), unsafe.Sizeof(uint32(0))*6))))
		v3 = (*uint8)(unsafe.Add(unsafe.Pointer(v3), 24))
		v1++
		v2 = (**byte)(unsafe.Pointer(v3))
		if v4 == 0 {
			return 0
		}
	}
	return int32(*memmap.PtrUint32(0x587000, uint32(v1*24+34860)))
}
func sub_415E40(a1 *byte) *byte {
	var (
		v1 int32
		i  *uint8
		v3 int32
	)
	v1 = 0
	if *memmap.PtrUint32(0x587000, 34852) == 0 {
		return nil
	}
	for i = (*uint8)(memmap.PtrOff(0x587000, 34852)); *((**byte)(unsafe.Add(unsafe.Pointer((**byte)(unsafe.Pointer(i))), unsafe.Sizeof((*byte)(nil))*2))) != a1; i = (*uint8)(unsafe.Add(unsafe.Pointer(i), 24)) {
		v3 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(i))), unsafe.Sizeof(uint32(0))*6))))
		v1++
		if v3 == 0 {
			return nil
		}
	}
	return *(**byte)(memmap.PtrOff(0x587000, uint32(v1*24+0x8820)))
}
func sub_415E80(a1 int32) int32 {
	var (
		v1 int32
		i  *uint8
		v3 int32
	)
	v1 = 0
	if *memmap.PtrUint32(0x587000, 35496) == 0 {
		return 0
	}
	for i = (*uint8)(memmap.PtrOff(0x587000, 35496)); *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(i))), unsafe.Sizeof(uint32(0))*2))) != uint32(a1); i = (*uint8)(unsafe.Add(unsafe.Pointer(i), 12)) {
		v3 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(i))), unsafe.Sizeof(uint32(0))*3))))
		v1++
		if v3 == 0 {
			return 0
		}
	}
	return int32(*memmap.PtrUint32(0x587000, uint32(v1*12+0x8AA8)))
}
func sub_415EC0(a1 *byte) int32 {
	var (
		v1     *byte
		v2     int32
		result int32
	)
	v1 = (*byte)(unsafe.Pointer(uintptr(sub_415DF0(a1))))
	if v1 != nil && (func() int32 {
		v2 = sub_415CD0(v1)
		return v2
	}()) != 0 {
		result = int32(uintptr(unsafe.Pointer(nox_xxx_objectTypeByInd_4E3B70(v2))))
	} else {
		result = 0
	}
	return result
}
func sub_4161E0() int32 {
	var (
		v0     int32
		v1     int32
		v2     int8
		v3     uint16
		v4     int16
		v5     *byte
		v6     uint8
		v7     uint8
		result int32
		v9     uint8
		v10    uint8
		v11    [4]byte
		v12    int32
		v13    [12]byte
		v14    [16]byte
		v15    [5]int32
	)
	sub_454040((*uint32)(unsafe.Pointer(&v15[0])))
	sub_4536B0((*uint32)(unsafe.Pointer(&v11[0])))
	v12 = sub_453710()
	v0 = nox_common_playerInfoCount_416F40()
	v1 = nox_xxx_servGetPlrLimit_409FA0()
	if nox_common_getEngineFlag(nox_engine_flag(NOX_ENGINE_FLAG_DISABLE_GRAPHICS_RENDERING)) {
		v0--
		v1--
	}
	if int32(*memmap.PtrUint8(6112660, 371619)) != v0 {
		*memmap.PtrUint8(6112660, 371619) = uint8(int8(v0))
		dword_5d4594_371692 = 1
	}
	if int32(*memmap.PtrUint8(6112660, 371620)) != v1 {
		*memmap.PtrUint8(6112660, 371620) = uint8(int8(v1))
		dword_5d4594_371692 = 1
	}
	if (int32(*memmap.PtrUint8(6112660, 371618)) & 239) != sub_43BE50_get_video_mode_id() {
		v2 = int8(sub_43BE50_get_video_mode_id())
		dword_5d4594_371692 = 1
		*memmap.PtrUint8(6112660, 371618) = uint8(int8(int32(*memmap.PtrUint8(6112660, 371618))&128 | int32(v2)))
	}
	v3 = uint16(nox_common_gameFlags_getVal_40A5B0())
	v4 = int16(*memmap.PtrUint16(6112660, 371432))
	if (int32(*memmap.PtrUint16(6112660, 371432))^int32(v3))&0xFFF0 != 0 {
		*memmap.PtrUint16(6112660, 371432) = uint16(nox_common_gameFlags_getVal_40A5B0())
		v4 = int16(*memmap.PtrUint16(6112660, 371432))
		dword_5d4594_371692 = 1
	}
	if int32(*memmap.PtrUint16(6112660, 371434)) != int32(nox_xxx_servGamedataGet_40A020(v4)) {
		*memmap.PtrUint16(6112660, 371434) = uint16(nox_xxx_servGamedataGet_40A020(*memmap.PtrInt16(6112660, 371432)))
		dword_5d4594_371692 = 1
	}
	if int32(*memmap.PtrUint8(6112660, 371436)) != int32(sub_40A180(*memmap.PtrInt16(6112660, 371432))) {
		*memmap.PtrUint8(6112660, 371436) = sub_40A180(*memmap.PtrInt16(6112660, 371432))
		dword_5d4594_371692 = 1
	}
	libc.StrNCpy(&v14[0], (*byte)(memmap.PtrOff(6112660, 371389)), 15)
	v14[15] = 0
	if libc.StrCmp(&v14[0], nox_xxx_serverOptionsGetServername_40A4C0()) != 0 {
		v5 = nox_xxx_serverOptionsGetServername_40A4C0()
		libc.StrNCpy((*byte)(memmap.PtrOff(6112660, 371389)), v5, 15)
		dword_5d4594_371692 = 1
	}
	libc.StrCpy(&v13[0], (*byte)(memmap.PtrOff(6112660, 371380)))
	if libc.StrCmp(&v13[0], nox_xxx_mapGetMapName_409B40()) != 0 {
		libc.StrCpy((*byte)(memmap.PtrOff(6112660, 371380)), nox_xxx_mapGetMapName_409B40())
		dword_5d4594_371692 = 1
	}
	v6 = 0
	v9 = 0
	for {
		if *memmap.PtrUint32(6112660, uint32(int32(v9)*4)+0x5AACC) != uint32(v15[v9]) {
			break
		}
		v9 = func() uint8 {
			p := &v6
			*p++
			return *p
		}()
		if int32(v6) >= 5 {
			break
		}
	}
	if int32(v6) != 5 {
		dword_5d4594_371692 = 1
		libc.MemCpy(memmap.PtrOff(6112660, 371404), unsafe.Pointer(&v15[0]), 20)
	}
	if !noxflags.HasGame(noxflags.GameHost) {
		return int32(dword_5d4594_371692)
	}
	if *memmap.PtrUint32(6112660, 371428) != uint32(v12) {
		*memmap.PtrUint32(6112660, 371428) = uint32(v12)
		dword_5d4594_371692 = 1
	}
	v7 = 0
	v10 = 0
	for {
		if (*memmap.PtrUint8(6112660, uint32(v10)+0x5AAE0)) != uint8(v11[v10]) {
			break
		}
		v10 = func() uint8 {
			p := &v7
			*p++
			return *p
		}()
		if int32(v7) >= 4 {
			break
		}
	}
	if int32(v7) == 4 {
		return int32(dword_5d4594_371692)
	}
	*memmap.PtrUint32(6112660, 371424) = *(*uint32)(unsafe.Pointer(&v11[0]))
	result = 1
	dword_5d4594_371692 = 1
	return result
}
func sub_4164F0() {
	dword_5d4594_371692 = 0
}
func nox_xxx_mapSetDataDefault_416500() {
	libc.MemSet(memmap.PtrOff(6112660, 371380), 0, 116)
	libc.MemSet(memmap.PtrOff(6112660, 371516), 0, 168)
	*memmap.PtrUint8(6112660, 371684) = 0
	*memmap.PtrUint16(6112660, 371621) = math.MaxUint16
	*memmap.PtrUint16(6112660, 371623) = math.MaxUint16
	*memmap.PtrUint32(6112660, 371688) = 0
	*memmap.PtrUint8(6112660, 371617) = math.MaxUint8
	*memmap.PtrUint8(6112660, 371568) = 32
	*memmap.PtrUint8(6112660, 371569) = 32
	*memmap.PtrUint32(6112660, 371578) = 1
	*memmap.PtrUint32(6112660, 371574) = 0
	*memmap.PtrUint32(6112660, 371582) = 0
	*memmap.PtrUint32(6112660, 371586) = 0
	*memmap.PtrUint32(6112660, 371590) = 20
	*memmap.PtrUint32(6112660, 371696) = 0
}
func sub_416580() int32 {
	return int32(*memmap.PtrUint32(6112660, 371688))
}
func nox_xxx_cliGamedataGet_416590(a1 int32) *byte {
	return (*byte)(memmap.PtrOff(6112660, uint32(a1*58)+371380))
}
func sub_4165B0() *byte {
	return (*byte)(memmap.PtrOff(6112660, *memmap.PtrUint32(6112660, 371688)*58+371380))
}
func sub_4165D0(a1 int32) *byte {
	*memmap.PtrUint32(6112660, 371688) = uint32(a1)
	return (*byte)(memmap.PtrOff(6112660, uint32(a1*58)+371380))
}
func sub_4165F0(a1 int32, a2 int32) int32 {
	var result int32
	result = a2
	libc.MemCpy(memmap.PtrOff(6112660, uint32(a2*58)+371380), memmap.PtrOff(6112660, uint32(a1*58)+371380), 58)
	return result
}
func sub_416630() *byte {
	return (*byte)(memmap.PtrOff(6112660, 371616))
}
func sub_416640() unsafe.Pointer {
	return memmap.PtrOff(6112660, 371516)
}
func sub_416650() int32 {
	return int32(*memmap.PtrUint32(6112660, 371700))
}
func sub_416670(a1 int32) int32 {
	var result int32
	result = a1
	*memmap.PtrUint32(6112660, 371700) = 0
	*memmap.PtrUint32(6112660, 371556) = uint32(a1)
	return result
}
func sub_416690() {
	var (
		v0 *byte
		v1 *byte
		v2 int16
		v3 [84]byte
	)
	if nox_xxx_check_flag_aaa_43AF70() == 1 {
		v0 = nox_xxx_cliGamedataGet_416590(0)
		sub_4161E0()
		v1 = sub_416630()
		libc.MemCpy(unsafe.Add(unsafe.Pointer(v1), 11), unsafe.Pointer(v0), 58)
		if nox_xxx_isQuest_4D6F50() != 0 {
			v2 = int16(int32(*(*uint16)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v1), 63))))) & 0xFF7F)
			*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v2))), unsafe.Sizeof(int16(0))-1)) |= 16
			*(*uint16)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v1), 63)))) = uint16(v2)
			*(*uint16)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v1), 65)))) = uint16(int16(nox_game_getQuestStage_4E3CC0()))
		}
		v3[sub_425550((*uint8)(unsafe.Pointer(v1)), (*uint8)(unsafe.Pointer(&v3[0])), 552)] = 0
		sub_40D560(int32(uintptr(unsafe.Pointer(&v3[0]))))
		sub_4164F0()
	}
}
func sub_416720() {
	var (
		v0 int32 = 0
		v2 int32 = int32(uintptr(unsafe.Pointer(sub_416900())))
	)
	for v2 != 0 {
		var v3 *int32 = sub_416910((*int32)(unsafe.Pointer(uintptr(v2))))
		if *(*uint32)(unsafe.Pointer(uintptr(v2 + 68))) != 0 || (*(*uint32)(unsafe.Pointer(uintptr(v2 + 64)))) != 0 {
			if uint64(platformTicks()) > *(*uint64)(unsafe.Pointer(uintptr(v2 + 64))) {
				sub_416820(v0)
			}
		}
		v0++
		v2 = int32(uintptr(unsafe.Pointer(v3)))
	}
}
func sub_416770(a1 int32, a2 *libc.WChar, a3 *byte) *int32 {
	var v3 *uint32
	v3 = (*uint32)(alloc.Calloc(1, 96))
	sub_425770(unsafe.Pointer(v3))
	nox_wcscpy((*libc.WChar)(unsafe.Add(unsafe.Pointer((*libc.WChar)(unsafe.Pointer(v3))), unsafe.Sizeof(libc.WChar(0))*6)), a2)
	if a3 != nil {
		libc.StrCpy((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(v3))), 72)), a3)
	} else {
		*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(v3))), 72))) = 0
	}
	nox_common_list_append_4258E0((*nox_list_item_t)(unsafe.Pointer(uintptr(int32(uintptr(memmap.PtrOff(6112660, 371500)))))), (*nox_list_item_t)(unsafe.Pointer(v3)))
	if a1 != 0 {
		*((*uint64)(unsafe.Add(unsafe.Pointer((*uint64)(unsafe.Pointer(v3))), unsafe.Sizeof(uint64(0))*8))) = uint64(uint32(a1*60000) + platformTicks())
	} else {
		*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*16)) = 0
		*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*17)) = 0
	}
	return sub_455800()
}
func sub_416820(a1 int32) {
	var (
		v2 *int32
		v3 int32
		v4 int32
	)
	v2 = sub_416900()
	if v2 == nil {
		return
	}
	v3 = a1
	for {
		v4 = func() int32 {
			p := &v3
			x := *p
			*p--
			return x
		}()
		if v4 == 0 {
			break
		}
		v2 = sub_416910(v2)
		if v2 == nil {
			return
		}
	}
	sub_425920((**uint32)(unsafe.Pointer(v2)))
	alloc.Free(unsafe.Pointer(v2))
}
func sub_416860(a1 int32) *int32 {
	var (
		result *int32
		v2     *int32
		v3     int32
		v4     int32
	)
	result = sub_4168E0()
	v2 = result
	if result != nil {
		v3 = a1
		for {
			v4 = func() int32 {
				p := &v3
				x := *p
				*p--
				return x
			}()
			if v4 == 0 {
				break
			}
			result = sub_4168F0(v2)
			v2 = result
			if result == nil {
				return result
			}
		}
		sub_425920((**uint32)(unsafe.Pointer(v2)))
		alloc.Free(unsafe.Pointer(v2))
	}
	return result
}
func sub_4168A0(a1 *libc.WChar) *int32 {
	var v1 *libc.WChar
	v1 = (*libc.WChar)(alloc.Calloc(1, 64))
	sub_425770(unsafe.Pointer(v1))
	nox_wcscpy((*libc.WChar)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(libc.WChar(0))*6)), a1)
	nox_common_list_append_4258E0((*nox_list_item_t)(unsafe.Pointer(uintptr(int32(uintptr(memmap.PtrOff(6112660, 371364)))))), (*nox_list_item_t)(unsafe.Pointer(v1)))
	return sub_455800()
}
func sub_4168E0() *int32 {
	return (*int32)(unsafe.Pointer(nox_common_list_getFirstSafe_425890((*nox_list_item_t)(unsafe.Pointer(memmap.PtrInt32(6112660, 371364))))))
}
func sub_4168F0(a1 *int32) *int32 {
	return (*int32)(unsafe.Pointer(nox_common_list_getNextSafe_4258A0((*nox_list_item_t)(unsafe.Pointer(a1)))))
}
func sub_416900() *int32 {
	return (*int32)(unsafe.Pointer(nox_common_list_getFirstSafe_425890((*nox_list_item_t)(unsafe.Pointer(memmap.PtrInt32(6112660, 371500))))))
}
func sub_416910(a1 *int32) *int32 {
	return (*int32)(unsafe.Pointer(nox_common_list_getNextSafe_4258A0((*nox_list_item_t)(unsafe.Pointer(a1)))))
}
func sub_416920() int32 {
	nox_common_list_clear_425760((*nox_list_item_t)(memmap.PtrOff(6112660, 371364)))
	nox_common_list_clear_425760((*nox_list_item_t)(memmap.PtrOff(6112660, 371500)))
	return sub_4E41B0((*byte)(memmap.PtrOff(0x587000, 54280)))
}
func sub_416950() *int32 {
	var (
		v0     *int32
		v1     *int32
		result *int32
		v3     *int32
		v4     *int32
	)
	sub_4E43F0(CString("ban.txt"))
	v0 = sub_4168E0()
	if v0 != nil {
		for {
			v1 = sub_4168F0(v0)
			sub_425920((**uint32)(unsafe.Pointer(v0)))
			alloc.Free(unsafe.Pointer(v0))
			v0 = v1
			if v1 == nil {
				break
			}
		}
	}
	result = sub_416900()
	v3 = result
	if result != nil {
		for {
			v4 = sub_416910(v3)
			sub_425920((**uint32)(unsafe.Pointer(v3)))
			alloc.Free(unsafe.Pointer(v3))
			v3 = v4
			if v4 == nil {
				break
			}
		}
	}
	return result
}
func sub_4169C0() int32 {
	return int32(*memmap.PtrUint32(6112660, 371704))
}
func nox_xxx_cliSetSettingsAcquired_4169D0(a1 int32) int32 {
	var result int32
	result = a1
	*memmap.PtrUint32(6112660, 371704) = uint32(a1)
	return result
}
func sub_4169E0() *byte {
	var result *byte
	result = (*byte)(sub_416640())
	*(*byte)(unsafe.Add(unsafe.Pointer(result), 100)) |= 16
	return result
}
func sub_4169F0() *byte {
	var result *byte
	result = (*byte)(sub_416640())
	*(*byte)(unsafe.Add(unsafe.Pointer(result), 100)) &= 239
	return result
}
func sub_416A00() uint32 {
	var v0 *byte
	v0 = (*byte)(sub_416640())
	*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v0))), 0)) = uint8(*(*byte)(unsafe.Add(unsafe.Pointer(v0), 100)))
	return (uint32(uintptr(unsafe.Pointer(v0))) >> 4) & 1
}
func nox_check_tick_flags() int32 {
	return bool2int(nox_common_getEngineFlag(nox_engine_flag(NOX_ENGINE_FLAG_PAUSE)) && nox_common_getEngineFlag(nox_engine_flag(NOX_ENGINE_FLAG_DISABLE_GRAPHICS_RENDERING)) && nox_get_audio_enabled() == 0)
}
func nox_xxx_playerForceSendLessons_416E50(a1 int32) *byte {
	var (
		result *byte
		i      *int32
	)
	result = (*byte)(unsafe.Pointer(nox_common_playerInfoGetFirst_416EA0()))
	for i = (*int32)(unsafe.Pointer(result)); result != nil; i = (*int32)(unsafe.Pointer(result)) {
		*(*int32)(unsafe.Add(unsafe.Pointer(i), unsafe.Sizeof(int32(0))*534)) = 0
		*(*int32)(unsafe.Add(unsafe.Pointer(i), unsafe.Sizeof(int32(0))*535)) = 0
		if a1 != 0 {
			if *(*int32)(unsafe.Add(unsafe.Pointer(i), unsafe.Sizeof(int32(0))*514)) != 0 {
				nox_xxx_netReportLesson_4D8EF0((*nox_object_t)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(i), unsafe.Sizeof(int32(0))*514))))))
			}
		}
		result = (*byte)(unsafe.Pointer(nox_common_playerInfoGetNext_416EE0((*nox_playerInfo)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(i)))))))))
	}
	return result
}
func nox_xxx_playerByName_4170D0(a1 *libc.WChar) *byte {
	var v1 *byte
	if a1 == nil {
		return nil
	}
	v1 = (*byte)(unsafe.Pointer(nox_common_playerInfoGetFirst_416EA0()))
	if v1 == nil {
		return nil
	}
	for _nox_wcsicmp((*libc.WChar)(unsafe.Add(unsafe.Pointer((*libc.WChar)(unsafe.Pointer(v1))), unsafe.Sizeof(libc.WChar(0))*2352)), a1) != 0 {
		v1 = (*byte)(unsafe.Pointer(nox_common_playerInfoGetNext_416EE0((*nox_playerInfo)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v1)))))))))
		if v1 == nil {
			return nil
		}
	}
	return v1
}
func sub_417120(a1 int32) *libc.WChar {
	var (
		v1     *byte
		result *libc.WChar
	)
	v1 = (*byte)(unsafe.Pointer(noxServer.getPlayerByID(a1)))
	if v1 != nil {
		result = (*libc.WChar)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v1), 4704))))
	} else {
		result = strMan.GetStringInFile("UnknownPlayerName", (*byte)(memmap.PtrOff(0x587000, 54432)))
	}
	return result
}
func nox_xxx_netMarkMinimapObject_417190(a1 int32, a2 int32, a3 int32) int32 {
	var (
		v3  int32
		v4  *uint32
		v5  *uint8
		v6  *uint32
		v8  *uint32
		v9  int32
		v10 int32
	)
	v3 = 0
	if !(a1 >= 0 && a1 < NOX_PLAYERINFO_MAX && a2 != 0) {
		return v3
	}
	var pl *nox_playerInfo = nox_common_playerInfoFromNumRaw(a1)
	v4 = (*uint32)(unsafe.Pointer(uintptr(pl.field_4580)))
	v5 = (*uint8)(unsafe.Pointer(pl))
	if v4 != nil {
		if *(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*1)) == uint32(a2) {
			v6 = (*uint32)(unsafe.Pointer(uintptr(pl.field_4580)))
			*v6 |= uint32(a3)
			return 1
		}
		v6 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*2)))))
		for v6 != v4 {
			if *(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*1)) == uint32(a2) {
				*v6 |= uint32(a3)
				return 1
			}
			v6 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*2)))))
		}
	}
	v8 = (*uint32)(alloc.Calloc(1, 16))
	if v8 == nil {
		return v3
	}
	*(*uint32)(unsafe.Add(unsafe.Pointer(v8), unsafe.Sizeof(uint32(0))*1)) = uint32(a2)
	*v8 = uint32(a3)
	v9 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v5))), unsafe.Sizeof(uint32(0))*1145))))
	if v9 != 0 {
		*(*uint32)(unsafe.Add(unsafe.Pointer(v8), unsafe.Sizeof(uint32(0))*2)) = uint32(v9)
		v10 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v8), unsafe.Sizeof(uint32(0))*2)))
		*(*uint32)(unsafe.Add(unsafe.Pointer(v8), unsafe.Sizeof(uint32(0))*3)) = *(*uint32)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v5))), unsafe.Sizeof(uint32(0))*1145))) + 12)))
		*(*uint32)(unsafe.Pointer(uintptr(v10 + 12))) = uint32(uintptr(unsafe.Pointer(v8)))
		*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v8), unsafe.Sizeof(uint32(0))*3)) + 8))) = uint32(uintptr(unsafe.Pointer(v8)))
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v5))), unsafe.Sizeof(uint32(0))*1145))) = uint32(uintptr(unsafe.Pointer(v8)))
	} else {
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v5))), unsafe.Sizeof(uint32(0))*1145))) = uint32(uintptr(unsafe.Pointer(v8)))
		*(*uint32)(unsafe.Add(unsafe.Pointer(v8), unsafe.Sizeof(uint32(0))*3)) = uint32(uintptr(unsafe.Pointer(v8)))
		*(*uint32)(unsafe.Add(unsafe.Pointer(v8), unsafe.Sizeof(uint32(0))*2)) = uint32(uintptr(unsafe.Pointer(v8)))
	}
	nox_xxx_unitSetXStatus_4E4800(a2, (*int32)(unsafe.Pointer(uintptr(1))))
	v3 = 1
	return v3
}
func sub_417270(a1 int32) int32 {
	var (
		result int32
		v2     int32
		v3     int32
	)
	result = 0
	if a1 >= 0 && a1 < NOX_PLAYERINFO_MAX {
		var pl *nox_playerInfo = nox_common_playerInfoFromNumRaw(a1)
		v2 = int32(pl.field_4580)
		if v2 != 0 {
			v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 8))))
			for result = 1; v3 != v2; result++ {
				v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(v3 + 8))))
			}
		}
	}
	return result
}
func sub_4172C0(a1 int32) int32 {
	var (
		result int32
		v2     int32
	)
	result = 0
	if a1 >= 0 && a1 < NOX_PLAYERINFO_MAX {
		var pl *nox_playerInfo = nox_common_playerInfoFromNumRaw(a1)
		v2 = int32(pl.field_4580)
		if v2 != 0 {
			result = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 4))))
			pl.field_4580 = *(*uint32)(unsafe.Pointer(uintptr(v2 + 8)))
		}
	}
	return result
}
func nox_xxx_netUnmarkMinimapObj_417300(a1 int32, a2 int32, a3 int32) int32 {
	var (
		result int32
		v4     *uint8
		v5     *uint32
		v6     int32
		v7     bool
		v8     *uint32
		v9     int32
	)
	result = 0
	if a1 >= 0 && a1 < NOX_PLAYERINFO_MAX {
		if a2 != 0 {
			var pl *nox_playerInfo = nox_common_playerInfoFromNumRaw(a1)
			v4 = (*uint8)(unsafe.Pointer(pl))
			v5 = (*uint32)(unsafe.Pointer(uintptr(pl.field_4580)))
			if v5 != nil {
				for {
					v6 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(uint32(0))*2)))
					if *(*uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(uint32(0))*1)) == uint32(a2) {
						break
					}
					if uint32(v6) != pl.field_4580 {
						v5 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(uint32(0))*2)))))
						if v6 != 0 {
							continue
						}
					}
					return result
				}
				v7 = (uint32(^a3) & *v5) == 0
				*v5 &= uint32(^a3)
				if v7 {
					v8 = (*uint32)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v4))), unsafe.Sizeof(uint32(0))*1145))))))
					if (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v8), unsafe.Sizeof(uint32(0))*2))))) == v8 {
						*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v4))), unsafe.Sizeof(uint32(0))*1145))) = 0
					} else {
						if v8 == v5 {
							*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v4))), unsafe.Sizeof(uint32(0))*1145))) = *(*uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(uint32(0))*2))
						}
						*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(uint32(0))*2)) + 12))) = *(*uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(uint32(0))*3))
						*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(uint32(0))*3)) + 8))) = *(*uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(uint32(0))*2))
					}
					alloc.Free(unsafe.Pointer(v5))
					v9 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 20))))
					*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v9))), 0)) = uint8(int8(v9 & 254))
					*(*uint32)(unsafe.Pointer(uintptr(a2 + 20))) = uint32(v9)
					result = 1
				} else {
					result = 0
				}
			}
		}
	}
	return result
}
func nox_xxx_playerMapTracksObj_4173D0(a1 int32, a2 int32) int32 {
	var (
		result int32
		v3     int32
	)
	result = 0
	if a1 >= 0 && a1 < NOX_PLAYERINFO_MAX {
		if a2 != 0 {
			var pl *nox_playerInfo = nox_common_playerInfoFromNumRaw(a1)
			v3 = int32(pl.field_4580)
			if v3 != 0 {
				for *(*uint32)(unsafe.Pointer(uintptr(v3 + 4))) != uint32(a2) {
					v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(v3 + 8))))
					if uint32(v3) == pl.field_4580 || v3 == 0 {
						return result
					}
				}
				result = 1
			}
		}
	}
	return result
}
func nox_xxx_netMinimapUnmark4All_417430(a1 int32) *byte {
	var (
		result *byte
		i      int32
	)
	result = (*byte)(unsafe.Pointer(nox_common_playerInfoGetFirst_416EA0()))
	for i = int32(uintptr(unsafe.Pointer(result))); result != nil; i = int32(uintptr(unsafe.Pointer(result))) {
		nox_xxx_netUnmarkMinimapObj_417300(int32(*(*uint8)(unsafe.Pointer(uintptr(i + 2064)))), a1, 3)
		result = (*byte)(unsafe.Pointer(nox_common_playerInfoGetNext_416EE0((*nox_playerInfo)(unsafe.Pointer(uintptr(i))))))
	}
	return result
}
func nox_xxx_netUnmarkMinimapSpec_417470(a1 int32, a2 int32) *byte {
	var (
		result *byte
		i      int32
	)
	result = (*byte)(unsafe.Pointer(nox_common_playerInfoGetFirst_416EA0()))
	for i = int32(uintptr(unsafe.Pointer(result))); result != nil; i = int32(uintptr(unsafe.Pointer(result))) {
		nox_xxx_netUnmarkMinimapObj_417300(int32(*(*uint8)(unsafe.Pointer(uintptr(i + 2064)))), a1, a2)
		result = (*byte)(unsafe.Pointer(nox_common_playerInfoGetNext_416EE0((*nox_playerInfo)(unsafe.Pointer(uintptr(i))))))
	}
	return result
}
func nox_xxx_netMarkMinimapForAll_4174B0(a1 int32, a2 int32) *byte {
	var (
		result *byte
		i      int32
	)
	result = (*byte)(unsafe.Pointer(nox_common_playerInfoGetFirst_416EA0()))
	for i = int32(uintptr(unsafe.Pointer(result))); result != nil; i = int32(uintptr(unsafe.Pointer(result))) {
		nox_xxx_netMarkMinimapObject_417190(int32(*(*uint8)(unsafe.Pointer(uintptr(i + 2064)))), a1, a2)
		result = (*byte)(unsafe.Pointer(nox_common_playerInfoGetNext_416EE0((*nox_playerInfo)(unsafe.Pointer(uintptr(i))))))
	}
	return result
}
func nox_xxx_netNeedTimestampStatus_4174F0(a1 int32, a2 int32) int32 {
	var result int32
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 3680))) |= uint32(a2)
	result = bool2int(noxflags.HasGame(noxflags.GameHost))
	if result != 0 {
		if a2&1059 != 0 {
			result = nox_xxx_netReportPlayerStatus_417630((*nox_playerInfo)(unsafe.Pointer(uintptr(a1))))
		}
	}
	return result
}
func nox_xxx_playerUnsetStatus_417530(a1p *nox_playerInfo, a2 int32) int8 {
	var (
		a1 int32 = int32(uintptr(unsafe.Pointer(a1p)))
		v2 int32
		v3 int16
	)
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 3680))) &= uint32(^a2)
	v2 = bool2int(noxflags.HasGame(noxflags.GameHost))
	if v2 != 0 {
		if a2&1 != 0 {
			v2 = bool2int(noxflags.HasGame(noxflags.GameModeChat))
			if v2 == 0 {
				v2 = nox_xxx_gamePlayIsAnyPlayers_40A8A0()
				if v2 != 0 {
					v2 = sub_40A220()
					if v2 == 0 {
						v3 = int16(uint16(nox_common_gameFlags_getVal_40A5B0()))
						*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v2))), 0)) = sub_40A180(v3)
						if int32(uint8(int8(v2))) != 0 {
							sub_40A250()
							*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v2))), 0)) = uint8(int8(sub_40A1F0(1)))
						}
					}
				}
			}
		}
		if a2&1059 != 0 {
			*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v2))), 0)) = uint8(int8(nox_xxx_netReportPlayerStatus_417630((*nox_playerInfo)(unsafe.Pointer(uintptr(a1))))))
		}
	}
	return int8(v2)
}
func nox_xxx_sendAllClientStatus_4175C0(a1 int32) *byte {
	var (
		result *byte
		i      int32
		v3     int32
		v4     [7]byte
	)
	result = (*byte)(unsafe.Pointer(nox_common_playerInfoGetFirst_416EA0()))
	for i = int32(uintptr(unsafe.Pointer(result))); result != nil; i = int32(uintptr(unsafe.Pointer(result))) {
		v4[0] = 106
		*(*uint16)(unsafe.Pointer(&v4[1])) = *(*uint16)(unsafe.Pointer(uintptr(i + 2060)))
		v3 = int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 2064))))
		*(*uint32)(unsafe.Pointer(&v4[3])) = *(*uint32)(unsafe.Pointer(uintptr(i + 3680))) & 1059
		nox_xxx_netSendPacket1_4E5390(v3, int32(uintptr(unsafe.Pointer(&v4[0]))), 7, 0, 0)
		result = (*byte)(unsafe.Pointer(nox_common_playerInfoGetNext_416EE0((*nox_playerInfo)(unsafe.Pointer(uintptr(i))))))
	}
	return result
}
func nox_xxx_netReportPlayerStatus_417630(pl *nox_playerInfo) int32 {
	var (
		a1 int32 = int32(uintptr(unsafe.Pointer(pl)))
		v1 int16
		v2 int32
		v4 [7]byte
	)
	v1 = int16(*(*uint16)(unsafe.Pointer(uintptr(a1 + 2060))))
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 3680))) & 1059)
	v4[0] = 106
	*(*uint16)(unsafe.Pointer(&v4[1])) = uint16(v1)
	*(*uint32)(unsafe.Pointer(&v4[3])) = uint32(v2)
	return nox_xxx_netSendPacket1_4E5390(math.MaxUint8, int32(uintptr(unsafe.Pointer(&v4[0]))), 7, 0, 0)
}
func nox_xxx_cliPlayerRespawn_417680(a1 int32, a2 int8) {
	var (
		v2  int32
		v3  int32
		v4  int32
		v5  *uint32
		v6  int32
		v7  int32
		v8  *uint32
		v9  int32
		v10 int32
		v11 int32
		v12 int32
		v13 int8
		v14 int32
		v15 int8
		v16 int32
		v17 int32
		v18 int32
		v19 int32
		v20 int8
		v21 int32
		v22 int32
		v23 int8
		v24 int32
		v25 int32
		v26 int32
	)
	v2 = a1
	if a1 == 0 {
		return
	}
	if !noxflags.HasGame(noxflags.GameHost) {
		*(*uint32)(unsafe.Pointer(uintptr(v2 + 4))) = 0
	}
	v3 = v2 + 2328
	v4 = 27
	for {
		v5 = (*uint32)(unsafe.Pointer(uintptr(v3)))
		*(*uint32)(unsafe.Pointer(uintptr(v3 - 4))) = 0
		v3 += 24
		*v5 = 0
		v4--
		*(*uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(uint32(0))*1)) = 0
		*(*uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(uint32(0))*2)) = 0
		*(*uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(uint32(0))*3)) = 0
		if v4 == 0 {
			break
		}
	}
	if !noxflags.HasGame(noxflags.GameHost) {
		*(*uint32)(unsafe.Pointer(uintptr(v2))) = 0
	}
	v6 = v2 + 2976
	v7 = 26
	for {
		v8 = (*uint32)(unsafe.Pointer(uintptr(v6)))
		*(*uint32)(unsafe.Pointer(uintptr(v6 - 4))) = 0
		v6 += 24
		*v8 = 0
		v7--
		*(*uint32)(unsafe.Add(unsafe.Pointer(v8), unsafe.Sizeof(uint32(0))*1)) = 0
		*(*uint32)(unsafe.Add(unsafe.Pointer(v8), unsafe.Sizeof(uint32(0))*2)) = 0
		*(*uint32)(unsafe.Add(unsafe.Pointer(v8), unsafe.Sizeof(uint32(0))*3)) = 0
		if v7 == 0 {
			break
		}
	}
	v9 = nox_xxx_modifGetIdByName_413290(CString("UserColor1"))
	v10 = int32(uintptr(unsafe.Pointer(nox_xxx_modifGetDescById_413330(v9))))
	if v10 == 0 {
		return
	}
	if int32(*(*uint8)(unsafe.Pointer(uintptr(v2 + 2251)))) != 0 || noxflags.HasGame(noxflags.GameModeCoop) {
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a1))), 0)) = math.MaxUint8
		v11 = int32(uintptr(unsafe.Pointer(nox_xxx_modifGetDescById_413330(int32(*(*uint32)(unsafe.Pointer(uintptr(v10 + 4))) + uint32(*(*uint8)(unsafe.Pointer(uintptr(v2 + 2269)))))))))
		v12 = int32(*(*uint8)(unsafe.Pointer(uintptr(v2 + 2270))))
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a1))), 1)) = *(*uint8)(unsafe.Pointer(uintptr(v11 + 4)))
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a1))), 2)) = *((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(nox_xxx_modifGetDescById_413330(int32(*(*uint32)(unsafe.Pointer(uintptr(v10 + 4)))+uint32(v12)))))), 4)))
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a1))), unsafe.Sizeof(int32(0))-1)) = math.MaxUint8
		if int32(a2)&1 != 0 {
			nox_xxx_clientEquipWeaponArmor_417AA0(82, int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 2060)))), 1024, int32(uintptr(unsafe.Pointer(&a1))))
		}
	}
	*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a1))), 0)) = math.MaxUint8
	*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a1))), 1)) = *((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(nox_xxx_modifGetDescById_413330(int32(*(*uint32)(unsafe.Pointer(uintptr(v10 + 4)))+uint32(*(*uint8)(unsafe.Pointer(uintptr(v2 + 2268))))))))), 4)))
	*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&a1))), unsafe.Sizeof(uint16(0))*1)) = math.MaxUint16
	if int32(a2)&2 != 0 {
		nox_xxx_clientEquipWeaponArmor_417AA0(82, int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 2060)))), 4, int32(uintptr(unsafe.Pointer(&a1))))
	}
	v13 = int8(*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(nox_xxx_modifGetDescById_413330(int32(*(*uint32)(unsafe.Pointer(uintptr(v10 + 4)))+uint32(*(*uint8)(unsafe.Pointer(uintptr(v2 + 2272))))))))), 4))))
	v14 = int32(*(*uint8)(unsafe.Pointer(uintptr(v2 + 2271))))
	*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a1))), 0)) = uint8(v13)
	*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a1))), 1)) = *((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(nox_xxx_modifGetDescById_413330(int32(*(*uint32)(unsafe.Pointer(uintptr(v10 + 4)))+uint32(v14)))))), 4)))
	*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&a1))), unsafe.Sizeof(uint16(0))*1)) = math.MaxUint16
	if int32(a2)&4 != 0 {
		nox_xxx_clientEquipWeaponArmor_417AA0(82, int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 2060)))), 1, int32(uintptr(unsafe.Pointer(&a1))))
	}
	v15 = int8(*(*uint8)(unsafe.Pointer(uintptr(v2 + 2251))))
	a1 = -1
	if int32(v15) == 1 {
		if noxflags.HasGame(noxflags.GameModeCoop) {
			if int32(a2)&8 != 0 {
				v16 = nox_xxx_modifGetIdByName_413290(CString("ArmorQuality1"))
				*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a1))), 0)) = *((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(nox_xxx_modifGetDescById_413330(v16)))), 4)))
				nox_xxx_clientEquipWeaponArmor_417AA0(80, int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 2060)))), 0x8000, int32(uintptr(unsafe.Pointer(&a1))))
			}
		} else if noxflags.HasGame(noxflags.GameModeQuest) {
			a1 = -1
			v17 = nox_xxx_modifGetIdByName_413290(CString("Replenishment1"))
			*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a1))), 2)) = *((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(nox_xxx_modifGetDescById_413330(v17)))), 4)))
			nox_xxx_clientEquipWeaponArmor_417AA0(80, int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 2060)))), 0x10000, int32(uintptr(unsafe.Pointer(&a1))))
		} else if int32(a2)&16 != 0 {
			nox_xxx_clientEquipWeaponArmor_417AA0(79, int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 2060)))), 0x4000, int32(uintptr(unsafe.Pointer(&a1))))
		}
	}
	if int32(*(*uint8)(unsafe.Pointer(uintptr(v2 + 2251)))) == 0 {
		if noxflags.HasGame(noxflags.GameModeCoop) {
			if int32(a2)&32 != 0 {
				v18 = nox_xxx_modifGetIdByName_413290(CString("ArmorQuality1"))
				*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a1))), 0)) = *((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(nox_xxx_modifGetDescById_413330(v18)))), 4)))
				v19 = nox_xxx_modifGetIdByName_413290(CString("Material1"))
				v20 = int8(*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(nox_xxx_modifGetDescById_413330(v19)))), 4))))
				v21 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 2060))))
				*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a1))), 0)) = uint8(v20)
				nox_xxx_clientEquipWeaponArmor_417AA0(80, v21, 256, int32(uintptr(unsafe.Pointer(&a1))))
			}
		} else if noxflags.HasGame(noxflags.GameModeQuest) {
			v25 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 2060))))
			a1 = -1
			nox_xxx_clientEquipWeaponArmor_417AA0(80, v25, 256, int32(uintptr(unsafe.Pointer(&a1))))
		} else {
			if int32(a2)&64 != 0 {
				nox_xxx_clientEquipWeaponArmor_417AA0(80, int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 2060)))), 512, int32(uintptr(unsafe.Pointer(&a1))))
			}
			if int32(a2) < 0 {
				nox_xxx_clientEquipWeaponArmor_417AA0(79, int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 2060)))), 0x1000000, int32(uintptr(unsafe.Pointer(&a1))))
			}
		}
	}
	if int32(*(*uint8)(unsafe.Pointer(uintptr(v2 + 2251)))) == 2 {
		if noxflags.HasGame(noxflags.GameModeCoop) {
			if int32(a2)&8 != 0 {
				v22 = nox_xxx_modifGetIdByName_413290(CString("ArmorQuality1"))
				v23 = int8(*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(nox_xxx_modifGetDescById_413330(v22)))), 4))))
				v24 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 2060))))
				*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a1))), 0)) = uint8(v23)
				nox_xxx_clientEquipWeaponArmor_417AA0(80, v24, 0x8000, int32(uintptr(unsafe.Pointer(&a1))))
			}
		} else if noxflags.HasGame(noxflags.GameModeQuest) {
			v26 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 2060))))
			a1 = -1
			nox_xxx_clientEquipWeaponArmor_417AA0(80, v26, 4, int32(uintptr(unsafe.Pointer(&a1))))
		}
	}
}
func nox_xxx_clientEquipWeaponArmor_417AA0(a1 int8, a2 int32, a3 int32, a4 int32) *byte {
	var (
		result *byte
		v5     *byte
		v6     int32
		v7     int32
		v8     *byte
		v9     *byte
		v10    int32
		v11    int32
		v12    *byte
	)
	result = (*byte)(unsafe.Pointer(noxServer.getPlayerByID(a2)))
	if result != nil {
		if int32(a1) == 81 || int32(a1) == 80 {
			v9 = (*byte)(unsafe.Add(unsafe.Pointer(result), 2324))
			*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(result))), unsafe.Sizeof(uint32(0))*1))) |= uint32(a3)
			v10 = 0
			for *(*uint32)(unsafe.Pointer(v9)) != 0 {
				v10++
				v9 = (*byte)(unsafe.Add(unsafe.Pointer(v9), 24))
				if v10 >= 27 {
					return result
				}
			}
			*(*uint32)(unsafe.Add(unsafe.Pointer(result), v10*24+2324)) = uint32(a3)
			v11 = 0
			v12 = (*byte)(unsafe.Add(unsafe.Pointer(result), v10*24+2328))
			for {
				result = (*byte)(unsafe.Pointer(nox_xxx_modifGetDescById_413330(int32(*(*uint8)(unsafe.Pointer(uintptr(v11 + a4)))))))
				*(*uint32)(unsafe.Pointer(v12)) = uint32(uintptr(unsafe.Pointer(result)))
				v11++
				v12 = (*byte)(unsafe.Add(unsafe.Pointer(v12), 4))
				if v11 >= 4 {
					break
				}
			}
		} else {
			v5 = (*byte)(unsafe.Add(unsafe.Pointer(result), 2972))
			*(*uint32)(unsafe.Pointer(result)) |= uint32(a3)
			v6 = 0
			for *(*uint32)(unsafe.Pointer(v5)) != 0 {
				v6++
				v5 = (*byte)(unsafe.Add(unsafe.Pointer(v5), 24))
				if v6 >= 26 {
					return result
				}
			}
			*(*uint32)(unsafe.Add(unsafe.Pointer(result), v6*24+2972)) = uint32(a3)
			v7 = 0
			v8 = (*byte)(unsafe.Add(unsafe.Pointer(result), v6*24+2976))
			for {
				result = (*byte)(unsafe.Pointer(nox_xxx_modifGetDescById_413330(int32(*(*uint8)(unsafe.Pointer(uintptr(v7 + a4)))))))
				*(*uint32)(unsafe.Pointer(v8)) = uint32(uintptr(unsafe.Pointer(result)))
				v7++
				v8 = (*byte)(unsafe.Add(unsafe.Pointer(v8), 4))
				if v7 >= 4 {
					break
				}
			}
		}
	}
	return result
}
func sub_417B80(a1 int8, a2 int32, a3 int32) *byte {
	var (
		result *byte
		v4     int32
		v5     int32
		v6     int32
		i      *byte
		v8     int32
		v9     int32
		j      *byte
	)
	result = (*byte)(unsafe.Pointer(noxServer.getPlayerByID(a2)))
	if result != nil {
		v4 = ^a3
		if int32(a1) == 84 {
			v5 = int32(uint32(v4) & *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(result))), unsafe.Sizeof(uint32(0))*1))))
			v6 = 0
			*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(result))), unsafe.Sizeof(uint32(0))*1))) = uint32(v5)
			for i = (*byte)(unsafe.Add(unsafe.Pointer(result), 2324)); *(*uint32)(unsafe.Pointer(i)) != uint32(a3); i = (*byte)(unsafe.Add(unsafe.Pointer(i), 24)) {
				if func() int32 {
					p := &v6
					*p++
					return *p
				}() >= 27 {
					return result
				}
			}
			*(*uint32)(unsafe.Add(unsafe.Pointer(result), v6*24+2324)) = 0
		} else {
			v8 = int32(uint32(v4) & *(*uint32)(unsafe.Pointer(result)))
			v9 = 0
			*(*uint32)(unsafe.Pointer(result)) = uint32(v8)
			for j = (*byte)(unsafe.Add(unsafe.Pointer(result), 2972)); *(*uint32)(unsafe.Pointer(j)) != uint32(a3); j = (*byte)(unsafe.Add(unsafe.Pointer(j), 24)) {
				if func() int32 {
					p := &v9
					*p++
					return *p
				}() >= 26 {
					return result
				}
			}
			*(*uint32)(unsafe.Add(unsafe.Pointer(result), v9*24+2972)) = 0
		}
	}
	return result
}
func sub_417C00(a1 *byte, a2 int32) *byte {
	var (
		result *byte
		i      int32
	)
	result = (*byte)(unsafe.Pointer(nox_common_playerInfoGetFirst_416EA0()))
	for i = int32(uintptr(unsafe.Pointer(result))); result != nil; i = int32(uintptr(unsafe.Pointer(result))) {
		if libc.StrCaseCmp((*byte)(unsafe.Pointer(uintptr(i+2096))), a1) == 0 {
			*(*uint32)(unsafe.Pointer(uintptr(i + 2108))) = uint32(a2)
		}
		result = (*byte)(unsafe.Pointer(nox_common_playerInfoGetNext_416EE0((*nox_playerInfo)(unsafe.Pointer(uintptr(i))))))
	}
	return result
}
func sub_417CF0() int32 {
	return nox_server_teamsZzz_419030(0)
}
func sub_417DC0() int32 {
	return int32(dword_5d4594_526276)
}
func sub_417DE0() int8 {
	var (
		v0 int8
		i  *byte
	)
	v0 = 0
	for i = (*byte)(unsafe.Pointer(nox_server_teamFirst_418B10())); i != nil; i = (*byte)(unsafe.Pointer(nox_server_teamNext_418B60((*nox_team_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(i))))))))) {
		if *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(i))), unsafe.Sizeof(uint32(0))*15))) != 0 {
			v0++
		}
	}
	return v0
}
func nox_xxx_mapInfoSetCapflag_417EA0() int32 {
	var (
		v0 int32
		v1 int32
	)
	*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v0))), 0)) = uint8(int8(bool2int(sub_417EC0())))
	v1 = v0
	if v0 != 0 {
		sub_455A50(2)
	}
	return v1
}
func sub_417EC0() bool {
	var i int32
	dword_5d4594_526276 = 0
	for i = int32(uintptr(unsafe.Pointer(noxServer.firstServerObject()))); i != 0; i = int32(uintptr(unsafe.Pointer((*nox_object_t)(unsafe.Pointer(uintptr(i))).Next()))) {
		if *(*uint32)(unsafe.Pointer(uintptr(i + 8)))&0x10000000 != 0 {
			dword_5d4594_526276++
		}
	}
	if dword_5d4594_526276 > 0 && !noxflags.HasGame(noxflags.GameFlag16) {
		sub_4181F0(0)
	}
	return dword_5d4594_526276 > 0
}
func nox_xxx_mapInfoSetFlagball_417F30() int8 {
	var v0 int32
	*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v0))), 0)) = uint8(int8(bool2int(sub_417EC0())))
	if v0 != 0 {
		sub_455F60()
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v0))), 0)) = uint8(int8(sub_417F50(0)))
	}
	return int8(v0)
}
func sub_417F50(a1 int32) int32 {
	var (
		v1  int32
		v2  int32
		v3  int32
		v4  int32
		v6  *uint32
		v7  *uint32
		v8  int32
		i   int32
		v10 float32
	)
	if dword_5d4594_527656 == 0 {
		dword_5d4594_527656 = uint32(nox_xxx_getNameId_4E3AA0(CString("GameBallStart")))
	}
	v1 = 0
	v2 = int32(uintptr(unsafe.Pointer(noxServer.firstServerObject())))
	if v2 == 0 {
		return 0
	}
	for {
		if uint32(*(*uint16)(unsafe.Pointer(uintptr(v2 + 4)))) == dword_5d4594_527656 {
			v1++
		}
		v2 = int32(uintptr(unsafe.Pointer((*nox_object_t)(unsafe.Pointer(uintptr(v2))).Next())))
		if v2 == 0 {
			break
		}
	}
	if v1 == 0 {
		return 0
	}
	v3 = noxRndCounter1.IntClamp(0, v1-1)
	v4 = int32(uintptr(unsafe.Pointer(noxServer.firstServerObject())))
	if v4 == 0 {
		return 0
	}
	for uint32(*(*uint16)(unsafe.Pointer(uintptr(v4 + 4)))) != dword_5d4594_527656 {
	LABEL_12:
		v4 = int32(uintptr(unsafe.Pointer((*nox_object_t)(unsafe.Pointer(uintptr(v4))).Next())))
		if v4 == 0 {
			return 0
		}
	}
	if v3 != 0 {
		v3--
		goto LABEL_12
	}
	v6 = (*uint32)(unsafe.Pointer(nox_xxx_newObjectByTypeID_4E3810(CString("GameBall"))))
	v7 = v6
	if v6 == nil {
		return 0
	}
	v8 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*187)))
	*(*uint64)(unsafe.Pointer(uintptr(v8 + 8))) = uint64(platformTicks())
	v10 = float32(nox_xxx_gamedataGetFloat_419D40(CString("FlagballPossDuration")))
	*(*uint32)(unsafe.Pointer(uintptr(v8 + 20))) = uint32(int(v10))
	*(*float32)(unsafe.Pointer(uintptr(v8 + 24))) = float32(nox_xxx_gamedataGetFloat_419D40(CString("FlagballResetVel")))
	nox_xxx_netMarkMinimapForAll_4174B0(int32(uintptr(unsafe.Pointer(v7))), 1)
	nox_xxx_createAt_4DAA50((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v7)))))), nil, 0.0, 0.0)
	nox_xxx_unitClearOwner_4EC300((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v7)))))))
	sub_4EB9B0(int32(uintptr(unsafe.Pointer(v7))), 0)
	sub_4E8290(0, 0)
	nox_xxx_unitMove_4E7010((*nox_object_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v7)))))), (*float2)(unsafe.Pointer(uintptr(v4+56))))
	*(*uint32)(unsafe.Add(unsafe.Pointer(v7), unsafe.Sizeof(uint32(0))*20)) = 0
	*(*uint32)(unsafe.Add(unsafe.Pointer(v7), unsafe.Sizeof(uint32(0))*21)) = 0
	*(*uint32)(unsafe.Add(unsafe.Pointer(v7), unsafe.Sizeof(uint32(0))*22)) = 0
	*(*uint32)(unsafe.Add(unsafe.Pointer(v7), unsafe.Sizeof(uint32(0))*25)) = 0
	if a1 != 0 {
		for i = int32(uintptr(unsafe.Pointer(nox_xxx_getFirstPlayerUnit_4DA7C0()))); i != 0; i = int32(uintptr(unsafe.Pointer(nox_xxx_getNextPlayerUnit_4DA7F0((*nox_object_t)(unsafe.Pointer(uintptr(i))))))) {
			if *(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(i + 748))) + 276))) + 3628))) == uint32(a1) {
				nox_xxx_playerCameraUnlock_4E6040(i)
				nox_xxx_playerCameraFollow_4E6060(i, int32(uintptr(unsafe.Pointer(v7))))
			}
		}
		nox_xxx_delayedDeleteObject_4E5CC0((*nox_object_t)(unsafe.Pointer(uintptr(a1))))
	}
	return 1
}
func nox_xxx_mapInfoSetKotr_4180D0() int32 {
	var (
		v0 int32
		v1 int32
		i  *byte
		v3 int32
		v4 int32
		v5 int32
		v6 bool
		v7 int8
		v8 *byte
	)
	if *memmap.PtrUint32(6112660, 527652) == 0 {
		*memmap.PtrUint32(6112660, 527652) = uint32(nox_xxx_getNameId_4E3AA0(CString("Crown")))
	}
	v0 = int32(uintptr(unsafe.Pointer(noxServer.firstServerObject())))
	if v0 != 0 {
		for {
			v1 = int32(uintptr(unsafe.Pointer((*nox_object_t)(unsafe.Pointer(uintptr(v0))).Next())))
			if *(*uint32)(unsafe.Pointer(uintptr(v0 + 8)))&0x10000000 != 0 {
				nox_xxx_delayedDeleteObject_4E5CC0((*nox_object_t)(unsafe.Pointer(uintptr(v0))))
			}
			v0 = v1
			if v1 == 0 {
				break
			}
		}
	}
	for i = (*byte)(unsafe.Pointer(nox_server_teamFirst_418B10())); i != nil; i = (*byte)(unsafe.Pointer(nox_server_teamNext_418B60((*nox_team_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(i))))))))) {
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(i))), unsafe.Sizeof(uint32(0))*19))) = 0
	}
	v3 = 0
	v4 = int32(uintptr(unsafe.Pointer(noxServer.firstServerObject())))
	if v4 == 0 {
		return 0
	}
	for {
		v5 = int32(uintptr(unsafe.Pointer((*nox_object_t)(unsafe.Pointer(uintptr(v4))).Next())))
		if uint32(*(*uint16)(unsafe.Pointer(uintptr(v4 + 4)))) != *memmap.PtrUint32(6112660, 527652) {
			v4 = v5
			continue
		}
		v3++
		*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v4 + 748))) + 4))) = 0
		v6 = !nox_xxx_CheckGameplayFlags_417DA0(4)
		v7 = int8(*(*uint8)(unsafe.Pointer(uintptr(v4 + 52))))
		if v6 {
			if int32(v7) != 0 {
				nox_xxx_delayedDeleteObject_4E5CC0((*nox_object_t)(unsafe.Pointer(uintptr(v4))))
				sub_4EC6A0(v4)
				v4 = v5
				continue
			}
			nox_xxx_netMarkMinimapForAll_4174B0(v4, 1)
			v4 = v5
			continue
		}
		if int32(v7) == 0 {
			nox_xxx_delayedDeleteObject_4E5CC0((*nox_object_t)(unsafe.Pointer(uintptr(v4))))
			sub_4EC6A0(v4)
			v4 = v5
			continue
		}
		v8 = (*byte)(unsafe.Pointer(nox_xxx_clientGetTeamColor_418AB0(int32(*(*uint8)(unsafe.Pointer(uintptr(v4 + 52)))))))
		if v8 != nil {
			*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v8))), unsafe.Sizeof(uint32(0))*19))) = uint32(v4)
			nox_xxx_netMarkMinimapForAll_4174B0(v4, 1)
		}
		v4 = v5
		if v5 == 0 {
			break
		}
	}
	if v3 == 0 {
		return 0
	}
	return 1
}
func sub_4181F0(a1 int32) {
	var (
		v1     uint8
		i      *byte
		result *byte
		v4     *byte
		v5     int32
		v6     int32
		v7     int32
		v8     bool
		v9     uint8
		v10    int32
		v11    *int32
		v12    int32
		v13    *int32
		v14    *byte
		v15    int32
		v16    int32
		v17    int32
		v18    *byte
		v19    uint8
		v20    uint8
		v21    uint8
		v22    [32]int32
	)
	v1 = 0
	v19 = 0
	if a1 != 0 {
		for i = (*byte)(unsafe.Pointer(nox_server_teamFirst_418B10())); i != nil; i = (*byte)(unsafe.Pointer(nox_server_teamNext_418B60((*nox_team_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(i))))))))) {
			sub_418D80(int32(uintptr(unsafe.Pointer(i))))
		}
	}
	result = (*byte)(unsafe.Pointer(nox_common_playerInfoGetFirst_416EA0()))
	v4 = result
	if result == nil {
		return
	}
	for {
		v5 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v4))), unsafe.Sizeof(uint32(0))*514))))
		if v5 != 0 && (*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v4))), unsafe.Sizeof(uint32(0))*515))) != nox_player_netCode_85319C || !nox_common_getEngineFlag(nox_engine_flag(NOX_ENGINE_FLAG_DISABLE_GRAPHICS_RENDERING))) {
			v6 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v4))), unsafe.Sizeof(uint32(0))*920))))
			if ((v6&1) == 0 || v6&32 != 0) && nox_xxx_servObjectHasTeam_419130(v5+48) == 0 {
				v7 = int32(v19)
				v19 = func() uint8 {
					p := &v1
					*p++
					return *p
				}()
				v22[v7] = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v4))), unsafe.Sizeof(uint32(0))*514))))
			}
		}
		result = (*byte)(unsafe.Pointer(nox_common_playerInfoGetNext_416EE0((*nox_playerInfo)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v4)))))))))
		v4 = result
		if result == nil {
			break
		}
	}
	if int32(v1) == 0 {
		return
	}
	v8 = int32(v1) <= 1
	v9 = v19
	if !v8 {
		v10 = 50
		for {
			v20 = uint8(int8(noxRndCounter1.IntClamp(0, int32(v9)-1)))
			v21 = uint8(int8(noxRndCounter1.IntClamp(0, int32(v9)-1)))
			v11 = &v22[v20]
			v10--
			v12 = *v11
			*v11 = v22[v21]
			v22[v21] = v12
			if v10 == 0 {
				break
			}
		}
	}
	result = (*byte)(unsafe.Pointer(uintptr(v9)))
	if int32(v9) <= 0 {
		return
	}
	v13 = &v22[0]
	v14 = result
	for {
		v15 = *v13
		v16 = int32(*(*uint32)(unsafe.Pointer(uintptr(*v13 + 748))))
		if sub_40A740() != 0 {
			v17 = int32(*(*uint32)(unsafe.Pointer(uintptr(v16 + 276))))
			result = *(**byte)(unsafe.Pointer(uintptr(v17 + 2068)))
			if result != nil {
				result = (*byte)(unsafe.Pointer(nox_server_teamByXxx_418AE0(int32(*(*uint32)(unsafe.Pointer(uintptr(v17 + 2068)))))))
				if result != nil {
					nox_xxx_createAtImpl_4191D0(uint8(*(*byte)(unsafe.Add(unsafe.Pointer(result), 57))), unsafe.Pointer(uintptr(v15+48)), 1, int32(*(*uint32)(unsafe.Pointer(uintptr(v15 + 36)))), 0)
				}
			}
		} else {
			v18 = sub_4189D0()
			nox_xxx_createAtImpl_4191D0(uint8(*(*byte)(unsafe.Add(unsafe.Pointer(v18), 57))), unsafe.Pointer(uintptr(v15+48)), 1, int32(*(*uint32)(unsafe.Pointer(uintptr(v15 + 36)))), 1)
		}
		v13 = (*int32)(unsafe.Add(unsafe.Pointer(v13), unsafe.Sizeof(int32(0))*1))
		v14 = (*byte)(unsafe.Add(unsafe.Pointer(v14), -1))
		if v14 == nil {
			break
		}
	}
}
func sub_418390() int32 {
	if int32(nox_xxx_getTeamCounter_417DD0()) == 0 {
		return 0
	}
	nox_xxx_SetGameplayFlag_417D50(2)
	sub_4181F0(0)
	return 1
}
func sub_4183C0() int32 {
	var (
		result int32
		i      int32
		v2     int32
		v3     int32
		v4     int32
		v5     *byte
		v6     int32
		v7     *byte
		v8     int32
		v9     float64
		v10    float64
		v11    float64
		v12    float32
	)
	result = int32(uintptr(unsafe.Pointer(nox_xxx_getFirstPlayerUnit_4DA7C0())))
	for i = result; result != 0; i = result {
		v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(i + 748))))
		if v2 == 0 || ((func() bool {
			v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 276))))
			return *(*uint32)(unsafe.Pointer(uintptr(v3 + 2060))) != nox_player_netCode_85319C
		}()) || !nox_common_getEngineFlag(nox_engine_flag(NOX_ENGINE_FLAG_DISABLE_GRAPHICS_RENDERING))) && ((func() bool {
			v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(v3 + 3680))))
			return (v4 & 1) == 0
		}()) || v4&32 != 0) {
			v5 = (*byte)(unsafe.Pointer(nox_xxx_clientGetTeamColor_418AB0(int32(*(*uint8)(unsafe.Pointer(uintptr(i + 52)))))))
			v6 = 0
			v12 = 1e+09
			v7 = (*byte)(unsafe.Pointer(nox_server_teamFirst_418B10()))
			if v7 != nil {
				for {
					v8 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v7))), unsafe.Sizeof(uint32(0))*18))))
					if v8 != 0 {
						v9 = float64(*(*float32)(unsafe.Pointer(uintptr(v8 + 60))) - *(*float32)(unsafe.Pointer(uintptr(i + 60))))
						v10 = float64(*(*float32)(unsafe.Pointer(uintptr(v8 + 56))) - *(*float32)(unsafe.Pointer(uintptr(i + 56))))
						v11 = v10*v10 + v9*v9
						if v11 < float64(v12) {
							v12 = float32(v11)
							v6 = int32(uintptr(unsafe.Pointer(v7)))
						}
					}
					v7 = (*byte)(unsafe.Pointer(nox_server_teamNext_418B60((*nox_team_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v7)))))))))
					if v7 == nil {
						break
					}
				}
				if v6 != 0 {
					if v5 != nil {
						if (*byte)(unsafe.Pointer(uintptr(v6))) != v5 {
							sub_4196D0(i+48, v6, int32(*(*uint32)(unsafe.Pointer(uintptr(i + 36)))), 0)
						}
					} else {
						nox_xxx_createAtImpl_4191D0(*(*uint8)(unsafe.Pointer(uintptr(v6 + 57))), unsafe.Pointer(uintptr(i+48)), 1, int32(*(*uint32)(unsafe.Pointer(uintptr(i + 36)))), 0)
					}
				}
			}
		}
		result = int32(uintptr(unsafe.Pointer(nox_xxx_getNextPlayerUnit_4DA7F0((*nox_object_t)(unsafe.Pointer(uintptr(i)))))))
	}
	return result
}
func sub_4184D0(a1p *nox_team_t) {
	var (
		a1 *libc.WChar = (*libc.WChar)(unsafe.Pointer(a1p))
		v1 int32
		v2 *uint32
		v3 [18]byte
	)
	if a1 != nil {
		sub_457230(a1)
		*(*uint32)(unsafe.Pointer(&v3[6])) = *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a1))), unsafe.Sizeof(uint32(0))*15)))
		*(*uint32)(unsafe.Pointer(&v3[2])) = uint32(*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(a1))), 57))))
		*(*uint16)(unsafe.Pointer(&v3[0])) = 196
		*(*uint32)(unsafe.Pointer(&v3[10])) = *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a1))), unsafe.Sizeof(uint32(0))*13)))
		v3[14] = 0
		v3[16] = byte(*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(a1))), 56))))
		v3[17] = byte(*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(a1))), 68))))
		if noxflags.HasGame(noxflags.GameModeSolo10) {
			v3[14] = 1
		}
		v3[15] = byte(nox_wcslen(a1))
		v1 = int32(uint8(v3[15])) * 2
		v2 = (*uint32)(alloc.Calloc(1, int(v1+18)))
		*(*uint64)(unsafe.Pointer(v2)) = *(*uint64)(unsafe.Pointer(&v3[0]))
		*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*2)) = *(*uint32)(unsafe.Pointer(&v3[8]))
		*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*3)) = *(*uint32)(unsafe.Pointer(&v3[12]))
		*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v2))), unsafe.Sizeof(uint16(0))*8))) = *(*uint16)(unsafe.Pointer(&v3[16]))
		libc.MemCpy(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(v2))), 18), unsafe.Pointer(a1), int(int32(uint8(v3[15]))*2))
		nox_xxx_netSendPacket1_4E5390(159, int32(uintptr(unsafe.Pointer(v2))), v1+18, 0, 1)
		alloc.Free(unsafe.Pointer(v2))
	}
}
func nox_xxx_wndGuiTeamCreate_4185B0() int32 {
	var (
		result int32
		i      int32
		v2     *byte
		v3     int32
		v4     int8
		v5     *libc.WChar
	)
	nox_xxx_SetGameplayFlag_417D50(4)
	nox_xxx_teamCreate_4186D0(0)
	nox_xxx_teamCreate_4186D0(0)
	result = int32(uintptr(unsafe.Pointer(noxServer.firstServerObject())))
	for i = result; result != 0; i = result {
		if *(*uint32)(unsafe.Pointer(uintptr(i + 8)))&0x10000000 != 0 {
			v2 = (*byte)(unsafe.Pointer(nox_xxx_clientGetTeamColor_418AB0(int32(*(*uint8)(unsafe.Pointer(uintptr(i + 52)))))))
			if v2 != nil {
				v3 = sub_4ECBD0(i)
				v4 = int8(v3)
				v5 = nox_server_teamTitle_418C20(v3)
				if v5 != nil {
					sub_418800((*libc.WChar)(unsafe.Pointer(v2)), v5, 1)
				}
				*(*byte)(unsafe.Add(unsafe.Pointer(v2), 56)) = byte(v4)
				sub_4184D0((*nox_team_t)(unsafe.Pointer((*libc.WChar)(unsafe.Pointer(v2)))))
				*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v2))), unsafe.Sizeof(uint32(0))*18))) = uint32(i)
			}
		}
		result = int32(uintptr(unsafe.Pointer((*nox_object_t)(unsafe.Pointer(uintptr(i))).Next())))
	}
	return result
}
func nox_xxx_teamAssignFlags_418640() int32 {
	var (
		result int32
		i      int32
		v2     int8
		v3     *byte
	)
	result = int32(uintptr(unsafe.Pointer(noxServer.firstServerObject())))
	for i = result; result != 0; i = result {
		if *(*uint32)(unsafe.Pointer(uintptr(i + 8)))&0x10000000 != 0 {
			v2 = int8(sub_4ECBD0(i))
			v3 = (*byte)(unsafe.Pointer(nox_xxx_clientGetTeamColor_418AB0(int32(*(*uint8)(unsafe.Pointer(uintptr(i + 52)))))))
			if v3 != nil {
				*(*byte)(unsafe.Add(unsafe.Pointer(v3), 56)) = byte(v2)
				*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v3))), unsafe.Sizeof(uint32(0))*18))) = uint32(i)
			}
		}
		result = int32(uintptr(unsafe.Pointer((*nox_object_t)(unsafe.Pointer(uintptr(i))).Next())))
	}
	return result
}
func nox_xxx_toggleAllTeamFlags_418690(a1 int32) *byte {
	var (
		result *byte
		i      *byte
		v3     int32
	)
	result = (*byte)(unsafe.Pointer(nox_server_teamFirst_418B10()))
	for i = result; result != nil; i = result {
		if *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(i))), unsafe.Sizeof(uint32(0))*18))) != 0 {
			v3 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(i))), unsafe.Sizeof(uint32(0))*18))))
			if a1 != 0 {
				nox_xxx_objectSetOn_4E75B0((*nox_object_t)(unsafe.Pointer(uintptr(v3))))
			} else {
				nox_xxx_objectSetOff_4E7600((*nox_object_t)(unsafe.Pointer(uintptr(v3))))
			}
		}
		result = (*byte)(unsafe.Pointer(nox_server_teamNext_418B60((*nox_team_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(i)))))))))
	}
	return result
}
