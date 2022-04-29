package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	"github.com/gotranspile/cxgo/runtime/stdio"
	"github.com/noxworld-dev/opennox/v1/common/alloc"
	noxflags "github.com/noxworld-dev/opennox/v1/common/flags"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"image"
	"math"
	"unsafe"
)

var dword_5d4594_830236 unsafe.Pointer = nil
var dword_5d4594_830232 unsafe.Pointer = nil
var nox_win_width_game int32 = NOX_DEFAULT_WIDTH
var nox_win_height_game int32 = NOX_DEFAULT_HEIGHT
var nox_win_depth_game int32 = NOX_DEFAULT_DEPTH
var dword_5d4594_816364 uint32 = 0
var dword_5d4594_816376 uint32 = 0
var dword_5d4594_816092 uint32 = 0
var nox_things_head *nox_thing = nil
var nox_things_array **nox_thing = nil
var nox_things_count int32 = 0
var obj_5D4594_3799660 nox_render_data_t = nox_render_data_t{}
var obj_5D4594_3800716 nox_render_data_t = nox_render_data_t{}
var nox_wnd_xxx_829520 *nox_gui_animation = nil
var nox_wnd_xxx_830244 *nox_gui_animation = nil

func nox_client_xxx_switchChatMap_43B510() {
	var (
		v0     int16
		result *byte
		v2     *byte
		v5     int16
		v6     int32
		v7     [80]byte
	)
	nox_client_gui_flag_815132 = 0
	noxflags.SetGame(noxflags.GameHost | noxflags.GameFlag3)
	setMouseBounds(image.Rect(0, 0, nox_win_width-1, nox_win_height-1))
	v6 = int32(*memmap.PtrUint32(6112660, 814916))
	if dword_587000_87404 == 1 {
		v0 = int16(int32(*memmap.PtrUint16(6112660, 814916)) - int32(*memmap.PtrUint16(0x587000, dword_587000_87412*8+87528)))
		*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v6))), unsafe.Sizeof(uint16(0))*1)) -= *memmap.PtrUint16(0x587000, dword_587000_87412*8+87530)
		*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v6))), unsafe.Sizeof(uint16(0))*0)) = uint16(v0)
	}
	libc.StrCpy(&v7[0], nox_client_getChatMap_49FF40((*int16)(unsafe.Pointer(&v6))))
	if libc.StrChr(&v7[0], byte('.')) == nil {
		sub_409B50(&v7[0])
		v2 = &v7[libc.StrLen(&v7[0])+1]
		*(*uint32)(unsafe.Pointer(func() *byte {
			p := &v2
			*p = (*byte)(unsafe.Add(unsafe.Pointer(*p), -1))
			return *p
		}())) = *memmap.PtrUint32(0x587000, 90856)
		*(*byte)(unsafe.Add(unsafe.Pointer(v2), 4)) = 0
		if !nox_common_getEngineFlag(nox_engine_flag(NOX_ENGINE_FLAG_REPLAY_READ)) {
			nox_xxx_gameSetMapPath_409D70(&v7[0])
		}
		noxflags.UnsetGame(noxflags.GameModeKOTR | noxflags.GameModeCTF | noxflags.GameModeFlagBall | noxflags.GameModeChat | noxflags.GameModeArena | noxflags.GameModeSolo10 | noxflags.GameModeElimination | noxflags.GameModeQuest | noxflags.GameFlag15 | noxflags.GameFlag16)
		noxflags.SetGame(noxflags.GameModeChat)
		result = sub_4165D0(0)
		v5 = int16(int32(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(result))), unsafe.Sizeof(uint16(0))*26)))) & 0x280F)
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v5))), 0)) = uint8(*(*byte)(unsafe.Add(unsafe.Pointer(result), 52))&15 | 128)
		*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(result))), unsafe.Sizeof(uint16(0))*26))) = uint16(v5)
	}
}
func sub_43B6D0() int32 {
	return int32(dword_5d4594_815044)
}
func sub_43BC10(a1 *libc.WChar, a2 uint8) *uint16 {
	var (
		a1_len       uint32      = nox_wcslen(a1)
		a1_last_char *libc.WChar = (*libc.WChar)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(libc.WChar(0))*uintptr(a1_len)))
		a1v          int32       = 0
	)
	for {
		nox_xxx_drawGetStringSize_43F840(nil, a1, &a1v, nil, 0)
		*a1_last_char = 0
		a1_last_char = (*libc.WChar)(unsafe.Add(unsafe.Pointer(a1_last_char), -int(unsafe.Sizeof(libc.WChar(0))*1)))
		if a1v+5 <= int32(a2) {
			break
		}
	}
	return (*uint16)(unsafe.Pointer(a1))
}
func nox_sprintAddrPort_43BC80(addr *byte, port uint16, dst *byte) int32 {
	return nox_sprintf(dst, CString("%s:%d"), addr, port)
}
func sub_43BD90(a1 int32) int32 {
	var result int32
	result = a1
	*memmap.PtrUint32(6112660, 815092) |= uint32(a1)
	return result
}
func sub_43BDB0() int32 {
	return int32(*memmap.PtrUint32(6112660, 815092))
}
func sub_43C650() int32 {
	var (
		v0     int64
		v1     uint32
		v2     int32
		result int32
	)
	v0 = int64(platformTicks())
	v1 = *memmap.PtrUint32(6112660, 815756)
	if *memmap.PtrUint64(6112660, 815740) != 0 {
		v2 = int32(uint32(bool2int(uint32(int32(v0)) < uint32(*memmap.PtrInt32(6112660, 815740)))) + *memmap.PtrUint32(6112660, 815744))
		*memmap.PtrUint32(6112660, *memmap.PtrUint32(6112660, 815756)*8+815220) = uint32(int32(v0 - int64(*memmap.PtrUint32(6112660, 815740))))
		*memmap.PtrUint32(6112660, v1*8+0xC7078) = (*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v0))), unsafe.Sizeof(uint32(0))*1))) - uint32(v2)
	} else {
		*memmap.PtrUint32(6112660, *memmap.PtrUint32(6112660, 815756)*8+815220) = uint32(int32(v0))
		*memmap.PtrUint32(6112660, v1*8+0xC7078) = *(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v0))), unsafe.Sizeof(uint32(0))*1))
	}
	*memmap.PtrUint64(6112660, 815756) = ((((uint64(*memmap.PtrUint32(6112660, 815760))) << 32) | uint64(v1)) + 1) % 60
	result = int32(dword_5d4594_815748 + 1)
	*memmap.PtrUint64(6112660, 815740) = uint64(v0)
	dword_5d4594_815748++
	return result
}
func sub_43C6E0() int32 {
	return bool2int(dword_5d4594_815704 == 0 && dword_5d4594_815708 == 0)
}
func nox_xxx_netGet_43C750() int32 {
	return int32(*(*uint32)(memmap.PtrOff(6112660, 815700)))
}
func sub_43C790() int32 {
	return int32(*memmap.PtrUint32(0x587000, 91876))
}
func sub_43C7A0(a1 int32) int32 {
	var result int32
	result = a1
	*memmap.PtrUint32(0x587000, 91876) = uint32(a1)
	return result
}
func nox_xxx_netHandleCliPacket_43C860(a1 int32, a2 *uint8, a3 int32, a4 unsafe.Pointer) int32 {
	var (
		v3 uint8
		v4 int32
	)
	v3 = *a2
	*memmap.PtrUint32(6112660, 815712) = uint32(a3)
	if int32(v3) < 39 {
		v4 = int32(*a2)
		if v4 == 33 {
			sub_446380()
		} else if v4 == 36 {
			nox_perfmon_ping_2614264 = *(*uint32)(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(a2), 1))))
		}
	} else {
		nox_xxx_netOnPacketRecvCli_48EA70(31, (*uint8)(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer(a2)))))), a3)
		if nox_client_isConnected_43C700() != 0 {
			sub_48D660()
		}
	}
	*memmap.PtrUint64(6112660, 815716) = uint64(platformTicks())
	if dword_5d4594_815704 == 1 {
		sub_4AB4A0(0)
		dword_5d4594_815704 = 0
	}
	if dword_5d4594_815708 == 1 {
		sub_43CF40()
	}
	return 1
}
func sub_43CC80() int32 {
	var result int32
	result = sub_5549F0(*(*uint32)(memmap.PtrOff(6112660, 815700)))
	dword_5d4594_2649712 = 0
	return result
}
func sub_43CCA0() {
	var (
		v0 int32
		v1 int32
		v2 uint64
		v3 uint64
		v4 uint64
		v5 uint64
	)
	nox_xxx_spriteDeleteSomeList_49C4B0()
	v0 = int32(nox_frame_xxx_2598000)
	nox_xxx_servNetInitialPackets_552A80(*(*uint32)(memmap.PtrOff(6112660, 815700)), 1)
	if uint32(v0) != nox_frame_xxx_2598000 && dword_5d4594_2650652 == 1 && !noxflags.HasGame(noxflags.GameHost) {
		v1 = sub_40A710(1)
		if sub_43C790() > v1 {
			sub_43CEB0()
			v2 = *memmap.PtrUint64(6112660, 815740) + *memmap.PtrUint64(0x587000, 91880)/uint64(sub_43C790())
			if uint64(platformTicks()) >= v2 {
				var v7 [8]byte
				v7[0] = 40
				*(*uint32)(unsafe.Pointer(&v7[1])) = nox_frame_xxx_2598000 + 1
				nox_xxx_netOnPacketRecvCli_48EA70(31, (*uint8)(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer(&v7[0])))))), 5)
			}
		}
	}
	v3 = uint64(platformTicks()) - qword_5d4594_815724
	if v3 >= 2000 {
		qword_5d4594_815724 = uint64(platformTicks())
		sub_552E70(*(*uint32)(memmap.PtrOff(6112660, 815700)))
	}
	if !noxflags.HasGame(noxflags.GameHost) {
		nox_xxx_netImportant_4E5770(31, 0)
	}
	nox_xxx_netSendBySock_40EE10(*(*uint32)(memmap.PtrOff(6112660, 815700)), 31, 0)
	nox_netlist_resetByInd_40ED10(31, 0)
	nox_xxx_netMaybeSendAll_552460()
	if !(*memmap.PtrUint32(6112660, 815720) != 0 || *memmap.PtrUint32(6112660, 815716) != 0) {
		return
	}
	v5 = uint64(platformTicks()) - *memmap.PtrUint64(6112660, 815716)
	if v5 > 2000 && dword_5d4594_815704 == 0 {
		dword_5d4594_815704 = 1
		sub_4AB4A0(1)
		*memmap.PtrUint64(6112660, 815732) = uint64(platformTicks())
	}
	if !(*memmap.PtrUint32(6112660, 815720) != 0 || *memmap.PtrUint32(6112660, 815716) != 0) {
		return
	}
	v4 = uint64(platformTicks()) - *memmap.PtrUint64(6112660, 815716)
	if v4 <= 20000 {
		return
	}
	if dword_5d4594_815708 != 0 {
		return
	}
	v4 = uint64(platformTicks()) - *memmap.PtrUint64(6112660, 815732)
	if v4 > 20000 {
		sub_43CF70()
	}
}
func sub_43CEB0() {
	var (
		v1 uint32
		v2 uint32
		v3 uint32
		v4 uint64 = 0
		v5 uint32
		v6 int32
		v7 uint32
		v9 uint32
		v0 int32 = int32(dword_5d4594_815748)
	)
	if *(*int32)(unsafe.Pointer(&dword_5d4594_815748)) >= 60 {
		v0 = 60
	}
	v1 = 0
	v2 = 0
	v3 = 0
	if !(v0 != 0 && v0 > 10) {
		*memmap.PtrUint32(0x587000, 91884) = 0
		*memmap.PtrUint32(0x587000, 91880) = 33
		return
	}
	v5 = 0
	v9 = uint32(v0)
	for {
		for {
			v6 = int32(*memmap.PtrUint32(6112660, v5*8+815220) + v2)
			v3 = uint32(((((uint64(*memmap.PtrUint32(6112660, v5*8+0xC7078))) << 32) | uint64(*memmap.PtrUint32(6112660, v5*8+815220))) + (((uint64(v3)) << 32) | uint64(v2))) >> 32)
			v2 += *memmap.PtrUint32(6112660, v5*8+815220)
			v7 = v5 + 1
			v1 = uint32(((((uint64(v1)) << 32) | uint64(func() uint32 {
				p := &v5
				x := *p
				*p++
				return x
			}())) + 1) >> 32)
			if v1 >= (*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v4))), unsafe.Sizeof(uint32(0))*1))) {
				break
			}
		}
		*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v4))), unsafe.Sizeof(uint32(0))*0)) = v9
		if !(v1 <= (*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v4))), unsafe.Sizeof(uint32(0))*1))) && v7 < v9) {
			break
		}
	}
	var v0a int64 = int64((((uint64(v3)) << 32) | uint64(uint32(v6))) / v4)
	*memmap.PtrUint64(0x587000, 91880) = uint64(v0a)
}
func sub_43CF40() int32 {
	*memmap.PtrUint64(6112660, 815732) = uint64(platformTicks())
	dword_5d4594_815708 = 0
	return sub_4AB4D0(0)
}
func sub_43CF70() int32 {
	var (
		result int32
		v2     int8
	)
	result = int32(dword_5d4594_815708)
	if dword_5d4594_815708 == 0 {
		sub_4AB4D0(1)
		dword_5d4594_815708 = 1
		result = int32(*memmap.PtrUint32(0x8531A0, 2576))
		if *memmap.PtrUint32(0x8531A0, 2576) != 0 {
			nox_xxx_netNeedTimestampStatus_4174F0(*memmap.PtrInt32(0x8531A0, 2576), 64)
			v2 = 41
			result = nox_xxx_netClientSend2_4E53C0(31, unsafe.Pointer(&v2), 1, 0, 1)
		}
	}
	return result
}
func nox_xxx_servNetInitialPacketsCheck_43CFD0() int32 {
	return bool2int(nox_xxx_servNetInitialPackets_552A80(*(*uint32)(memmap.PtrOff(6112660, 815700)), 1) != -1)
}
func nox_xxx_servNetInitialPacketsUntilCRC_43CFF0() int32 {
	var start int64 = int64(platformTicks())
	for uint64(int64(platformTicks())-start) < 10000 {
		nox_xxx_servNetInitialPackets_552A80(*(*uint32)(memmap.PtrOff(6112660, 815700)), 1)
		nox_xxx_netSendBySock_40EE10(*(*uint32)(memmap.PtrOff(6112660, 815700)), 31, 0)
		nox_netlist_resetByInd_40ED10(31, 0)
		nox_xxx_netMaybeSendAll_552460()
		if nox_xxx_getMapCRC_40A370() != 0 {
			return 1
		}
	}
	return 0
}
func sub_43D2D0() {
	var v1 int32
	if dword_5d4594_816340 != 0 {
		if dword_5d4594_816364 != 0 {
			sub_486520(memmap.PtrUint32(6112660, 816244))
			sub_486520(memmap.PtrUint32(6112660, 816148))
			v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(dword_587000_81128)) + 4))) >> 16)
			if uint32(v1) == *memmap.PtrUint32(0x587000, 93168) {
				if (uint64(platformTicks()) - *memmap.PtrUint64(6112660, 816380)) > 50 {
					**(**uint32)(unsafe.Pointer(&dword_587000_81128)) &= 0xFFFFFFFD
					*(*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(dword_587000_81128)) + 32))) &= 0xFFFFFFFD
					*(*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(dword_587000_81128)) + 64))) &= 0xFFFFFFFD
				}
			} else {
				*memmap.PtrUint64(6112660, 816380) = uint64(platformTicks())
				*memmap.PtrUint32(0x587000, 93168) = uint32(v1)
			}
			if dword_5d4594_816364 != 0 {
				if int32(**(**uint8)(unsafe.Pointer(&dword_587000_81128)))&2 != 0 || int32(*memmap.PtrUint8(6112660, 816148))&2 != 0 || int32(*memmap.PtrUint8(6112660, 816244))&2 != 0 {
					sub_43D3C0(*(*int32)(unsafe.Pointer(&dword_5d4594_816364)), *memmap.PtrInt32(6112660, 816096))
				}
			}
		}
	}
}
func sub_43D3C0(a1 int32, a2 int32) {
	var v2 uint32
	if a1 != 0 {
		v2 = (*(*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(dword_587000_81128)) + 4))) >> 16) * ((uint32(*memmap.PtrUint16(6112660, 816154)) * ((uint32(a2) * uint32(*memmap.PtrUint16(6112660, 816250))) >> 14)) >> 14)
		*memmap.PtrUint32(6112660, 816148) &= 0xFFFFFFFD
		*memmap.PtrUint32(6112660, 816244) &= 0xFFFFFFFD
		AIL_set_stream_volume((HSTREAM)(unsafe.Pointer(uintptr(a1))), int32((v2>>14)*math.MaxInt8)/100)
	}
}
func sub_43D440() {
	var v0 *int32 = memmap.PtrInt32(6112660, *memmap.PtrUint32(6112660, 816352)*20+0xC73EC)
	if dword_5d4594_816340 == 0 {
		return
	}
	switch dword_5d4594_816348 {
	case 0:
		if dword_5d4594_816356 != 0 && dword_587000_93156 != 0 {
			dword_5d4594_816348 = 3
		} else {
			if *v0 != 0 {
				if dword_587000_93156 != 0 {
					sub_486320((*uint32)(memmap.PtrOff(6112660, 816148)), 0x4000)
					if nox_xxx_musicStartPlay_43D6C0(v0) != 0 {
						dword_5d4594_816348 = 1
						*(*int32)(unsafe.Add(unsafe.Pointer(v0), unsafe.Sizeof(int32(0))*4)) = 1
					} else {
						*v0 = 0
					}
				}
			}
		}
	case 1:
		if dword_587000_93156 != 0 && uint32(*v0) == dword_5d4594_816092 && dword_5d4594_816364 != 0 && AIL_stream_status((HSTREAM)(unsafe.Pointer(uintptr(dword_5d4594_816364)))) != 2 {
			if dword_5d4594_816356 != 0 {
				dword_5d4594_816348 = 4
				sub_486350(unsafe.Pointer(uintptr(int32(uintptr(memmap.PtrOff(6112660, 816148))))), 0)
			}
		} else {
			dword_5d4594_816348 = 2
			sub_486350(unsafe.Pointer(uintptr(int32(uintptr(memmap.PtrOff(6112660, 816148))))), 0)
		}
	case 2:
		if dword_5d4594_816364 == 0 {
			sub_43D650()
			dword_5d4594_816348 = 0
		} else {
			if AIL_stream_status((HSTREAM)(unsafe.Pointer(uintptr(dword_5d4594_816364)))) == 2 || (*memmap.PtrUint32(6112660, 816152)&0xFFFF0000) == 0 {
				sub_43D650()
				dword_5d4594_816348 = 0
			}
		}
	case 3:
		if dword_5d4594_816356 == 0 || dword_587000_93156 == 0 {
			if dword_587000_93156 == 0 || uint32(*v0) != dword_5d4594_816092 || dword_5d4594_816364 == 0 || AIL_stream_status((HSTREAM)(unsafe.Pointer(uintptr(dword_5d4594_816364)))) == 2 {
				sub_43D650()
				dword_5d4594_816348 = 0
			} else {
				sub_486350(unsafe.Pointer(uintptr(int32(uintptr(memmap.PtrOff(6112660, 816148))))), 0x4000)
				sub_43D6A0()
				dword_5d4594_816348 = 1
			}
		}
	case 4:
		if dword_5d4594_816364 != 0 && AIL_stream_status((HSTREAM)(unsafe.Pointer(uintptr(dword_5d4594_816364)))) != 2 {
			if (*memmap.PtrUint32(6112660, 816152) & 0xFFFF0000) == 0 {
				sub_43D680()
				dword_5d4594_816348 = 3
			}
		} else {
			sub_43D650()
			dword_5d4594_816348 = 0
		}
	default:
	}
}
func sub_43D8E0() int32 {
	var (
		result int32
		v1     int32
	)
	if dword_5d4594_816340 != 0 {
		return 1
	}
	v1 = sub_43F130()
	dword_5d4594_816376 = uint32(v1)
	dword_587000_93156 = uint32(bool2int(v1 != 0))
	sub_4864A0((*uint32)(memmap.PtrOff(6112660, 816148)))
	sub_486380((*uint32)(memmap.PtrOff(6112660, 816148)), 1000, 0, 0x4000)
	dword_5d4594_816348 = 0
	dword_5d4594_816092 = 0
	*memmap.PtrUint32(6112660, 816352) = 0
	*memmap.PtrUint32(6112660, 816108) = 0
	dword_5d4594_816356 = 0
	dword_5d4594_816372 = 0
	dword_5d4594_816368 = 0
	dword_5d4594_816364 = 0
	result = 1
	dword_5d4594_816340 = 1
	return result
}
func sub_43D990() {
	var v1 int4
	v1.field_0 = 0
	sub_43D9E0(&v1)
}
func sub_43D9B0(a1 int32, a2 int32) {
	var v3 int4
	v3.field_0 = a1
	v3.field_C = 0
	v3.field_8 = 0
	v3.field_4 = a2
	sub_43D9E0(&v3)
}
func sub_43D9E0(a1 *int4) {
	var (
		v1 *int4
		v2 int32
	)
	v1 = (*int4)(memmap.PtrOff(6112660, (*memmap.PtrUint32(6112660, 816352)^1)*20+0xC73EC))
	v2 = a1.field_4
	if v2 <= 100 {
		if v2 < 0 {
			a1.field_4 = 0
		}
	} else {
		a1.field_4 = 100
	}
	if *memmap.PtrUint32(6112660, 816344) != 0 {
		*(*int4)(memmap.PtrOff(6112660, 816060)) = *a1
	} else {
		*v1 = *a1
		(*(*int4)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(int4{})*1))).field_0 = 0
		*memmap.PtrUint32(6112660, 816352) ^= 1
	}
}
func sub_43DA80() int32 {
	var result int32
	if *(*int32)(unsafe.Pointer(&dword_5d4594_816368)) < 6 {
		sub_43DD10(memmap.PtrInt32(6112660, (dword_5d4594_816368+dword_5d4594_816372*6)*16+0xC729C))
		dword_5d4594_816368++
		result = 1
	} else {
		dword_5d4594_816368 = 6
		result = 0
	}
	return result
}
func sub_43DAD0() {
	if dword_5d4594_816368 > 0 {
		sub_43D9E0((*int4)(memmap.PtrOff(6112660, (func() uint32 {
			p := &dword_5d4594_816368
			*p--
			return *p
		}()+dword_5d4594_816372*6)*16+0xC729C)))
	}
	dword_5d4594_816368 = 0
}
func sub_43DB10() {
	sub_43D990()
	dword_5d4594_816368 = 0
}
func sub_43DB20() int32 {
	return int32(dword_5d4594_816368)
}
func sub_43DB30(a1 int32) int32 {
	var result int32
	result = a1
	dword_5d4594_816368 = uint32(a1)
	return result
}
func sub_43DB40(a1 int32) *byte {
	return (*byte)(memmap.PtrOff(6112660, (uint32(a1)+dword_5d4594_816372*6)*16+0xC729C))
}
func sub_43DB60() int32 {
	var (
		result int32
		v1     int32
		v2     int32
	)
	result = 3
	if *(*int32)(unsafe.Pointer(&dword_5d4594_816372)) < 3 {
		sub_43DA80()
		v1 = int32(dword_5d4594_816372)
		v2 = int32(dword_5d4594_816368)
		dword_5d4594_816368 = 0
		*memmap.PtrUint32(6112660, dword_5d4594_816372*4+0xC73CC) = uint32(v2)
		result = v1 + 1
		dword_5d4594_816372 = uint32(result)
	} else {
		dword_5d4594_816372 = 3
	}
	return result
}
func sub_43DBA0() {
	var v1 int32
	if dword_5d4594_816372 > 0 {
		v1 = int32(dword_5d4594_816372 - 1)
		dword_5d4594_816372 = uint32(v1)
		dword_5d4594_816368 = *memmap.PtrUint32(6112660, uint32(v1*4)+0xC73CC)
		sub_43DAD0()
	} else {
		dword_5d4594_816372 = 0
	}
}
func sub_43DBD0() int32 {
	dword_5d4594_816356++
	sub_43D440()
	return 0
}
func sub_43DBE0() int32 {
	if func() int32 {
		p := (*int32)(unsafe.Pointer(&dword_5d4594_816356))
		*p--
		return *p
	}() < 0 {
		dword_5d4594_816356 = 0
	}
	sub_43D440()
	return 0
}
func sub_43DC00() {
	dword_587000_93156 = 0
}
func sub_43DC10() int32 {
	var result int32
	result = int32(dword_5d4594_816376)
	if dword_5d4594_816376 != 0 {
		dword_587000_93156 = 1
	}
	return result
}
func sub_43DC30() int32 {
	return int32(dword_587000_93156)
}
func sub_43DC40() int32 {
	var result int32
	if dword_5d4594_816340 == 0 || (func() bool {
		result = 1
		return dword_5d4594_816348 != 1
	}()) && dword_5d4594_816348 != 4 && dword_5d4594_816348 != 2 {
		result = 0
	}
	return result
}
func sub_43DC80() int32 {
	var v0 *uint8
	if dword_587000_93156 != 0 && dword_5d4594_816340 != 0 {
		for {
			v0 = (*uint8)(memmap.PtrOff(6112660, *memmap.PtrUint32(6112660, 816352)*20+0xC73EC))
			if *(*uint32)(unsafe.Pointer(v0)) != 0 {
				if *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v0))), unsafe.Sizeof(uint32(0))*4))) != 0 {
					break
				}
			}
			sub_4312C0()
		}
	}
	return 1
}
func sub_43DD10(a1 *int32) int32 {
	var (
		v1     *uint8
		result int32
	)
	v1 = (*uint8)(memmap.PtrOff(6112660, *memmap.PtrUint32(6112660, 816352)*20+0xC73EC))
	*a1 = int32(*memmap.PtrUint32(6112660, *memmap.PtrUint32(6112660, 816352)*20+0xC73EC))
	*(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*1)) = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v1))), unsafe.Sizeof(uint32(0))*1))))
	*(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*2)) = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v1))), unsafe.Sizeof(uint32(0))*2))))
	*(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*3)) = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v1))), unsafe.Sizeof(uint32(0))*3))))
	result = int32(dword_5d4594_816364)
	if dword_5d4594_816364 != 0 {
		result = *a1
		if dword_5d4594_816092 == uint32(*a1) {
			result = AIL_stream_position((HSTREAM)(unsafe.Pointer(uintptr(dword_5d4594_816364))))
			*(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*2)) = result
		}
	}
	return result
}
func sub_43DD70(a1 int32, a2 int32) {
	sub_43DD10(memmap.PtrInt32(6112660, 816060))
	sub_43D9B0(a1, a2)
	*memmap.PtrUint32(6112660, 816344) = 1
}
func nox_xxx_gui_43E1A0(a1 int32) *uint32 {
	var result *uint32
	if a1 != 0 {
		result = (*uint32)(unsafe.Pointer(nox_window_new(nil, 552, 0, 0, nox_win_width, nox_win_height, nil)))
		dword_5d4594_816412 = uint32(uintptr(unsafe.Pointer(result)))
		*(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*14)) = nox_color_black_2650656
	} else {
		result = *(**uint32)(unsafe.Pointer(&dword_5d4594_816412))
		if dword_5d4594_816412 != 0 {
			result = (*uint32)(unsafe.Pointer(uintptr((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_816412)))).Destroy())))
			dword_5d4594_816412 = 0
		}
	}
	return result
}
func sub_43E8C0(a1 int32) int32 {
	var result int32
	result = a1
	*memmap.PtrUint32(6112660, 816408) = uint32(a1)
	return result
}
func sub_43E9C0() int32 {
	sub_43E9F0()
	return 0
}
func sub_43EDB0(a1 HSAMPLE) {
	var (
		v1     int32
		v2     int32
		result int32
	)
	v1 = AIL_sample_user_data(a1, 0)
	v2 = v1
	result = int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 28))))
	if result == 0 {
		result = (*(*funcuint32(int32))(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 4))) + 284))))(*(*uint32)(unsafe.Pointer(uintptr(v2 + 4))))
		*(*uint32)(unsafe.Pointer(uintptr(v2 + 28))) = 1
	}
}
func sub_43EE00(a1 int32) int32 {
	var (
		v1     int32
		result int32
		v3     int32
		v4     int32
		v5     *byte
		v6     uint32
		v7     int32
		v8     int32
		v9     *byte
		v10    int8
		v11    uint32
		v12    *byte
		v13    *byte
		v14    int32
		v15    int32
		v16    int32
		v17    *byte
		v18    int32
		v19    *byte
		v20    int32
		v21    int32
		v22    *byte
	)
	v1 = a1
	result = AIL_sample_buffer_ready((HSAMPLE)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 8)))))))
	v20 = result
	if result != -1 {
		for {
			v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 4))))
			v4 = 0
			v5 = nil
			v17 = *(**byte)(unsafe.Pointer(uintptr(v1 + v20*4 + 20)))
			v6 = 0
			v7 = 0
			v21 = 0
			v19 = nil
			v18 = 0
			if *(*uint32)(unsafe.Pointer(uintptr(a1 + 12))) == 0 {
				break
			}
		LABEL_27:
			AIL_load_sample_buffer((HSAMPLE)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 8)))))), uint32(v20), unsafe.Pointer(v19), uint32(v4))
			result = AIL_sample_buffer_ready((HSAMPLE)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 8)))))))
			v20 = result
			if result == -1 {
				return result
			}
			v1 = a1
		}
		for {
			if v4 >= 0x4000 {
				goto LABEL_27
			}
			v8 = int32(*(*uint32)(unsafe.Pointer(uintptr(v3 + 300))))
			if v8 == 0 {
				(cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(v3 + 276))), (*func(int32))(nil)))(v3)
				v8 = int32(*(*uint32)(unsafe.Pointer(uintptr(v3 + 300))))
				if v8 == 0 {
					(cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(v3 + 280))), (*func(int32))(nil)))(v3)
					v8 = int32(*(*uint32)(unsafe.Pointer(uintptr(v3 + 300))))
					if v8 == 0 {
						*(*uint32)(unsafe.Pointer(uintptr(a1 + 12))) = 1
						goto LABEL_27
					}
				}
				v6 = uint32(v18)
			}
			if v4 != 0 {
				break
			}
			v19 = *(**byte)(unsafe.Pointer(uintptr(v3 + 296)))
			if v8 < 0x4000 || v5 != nil {
				break
			}
			v7 = v8
		LABEL_24:
			v15 = int32(*(*uint32)(unsafe.Pointer(uintptr(v3 + 304))) - uint32(v7))
			*(*uint32)(unsafe.Pointer(uintptr(v3 + 300))) -= uint32(v7)
			v16 = int32(uint32(v7) + *(*uint32)(unsafe.Pointer(uintptr(v3 + 296))))
			*(*uint32)(unsafe.Pointer(uintptr(v3 + 304))) = uint32(v15)
			*(*uint32)(unsafe.Pointer(uintptr(v3 + 296))) = uint32(v16)
			v4 += v7
			v21 = v4
			if *(*uint32)(unsafe.Pointer(uintptr(a1 + 12))) != 0 {
				goto LABEL_27
			}
		}
		if v8+v4 > 0x4000 {
			v8 = 0x4000 - v4
		}
		if v8 == 0 {
			goto LABEL_24
		}
		if v5 != nil {
			if v6 != 0 {
				v9 = v5
				v10 = int8(uint8(v6))
				v11 = v6 >> 2
				libc.MemCpy(unsafe.Pointer(v22), unsafe.Pointer(v9), int(v11*4))
				v13 = (*byte)(unsafe.Add(unsafe.Pointer(v9), v11*4))
				v12 = (*byte)(unsafe.Add(unsafe.Pointer(v22), v11*4))
				v19 = v22
				v14 = int32(v10) & 3
				v5 = nil
				libc.MemCpy(unsafe.Pointer(v12), unsafe.Pointer(v13), int(v14))
			}
		} else if v6 == 0 {
			v5 = *(**byte)(unsafe.Pointer(uintptr(v3 + 296)))
			v22 = v17
			v18 = v8
		LABEL_23:
			v7 = v8
			v17 = (*byte)(unsafe.Add(unsafe.Pointer(v17), v8))
			v6 = uint32(v18)
			goto LABEL_24
		}
		libc.MemCpy(unsafe.Pointer(v17), *(*unsafe.Pointer)(unsafe.Pointer(uintptr(v3 + 296))), int(v8))
		v4 = v21
		goto LABEL_23
	}
	return result
}
func sub_43F0E0(a1 *uint32) int32 {
	var (
		v1     int32
		result int32
		v3     int32
	)
	v1 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*1)))
	if v1 != 0 {
		if v1 == 2 {
			v3 = -bool2int(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*3)) != 2)
			*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v3))), 0)) = uint8(int8(v3 & 254))
			result = v3 + 7
		} else {
			result = 0
		}
	} else if *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*3)) == 2 {
		result = bool2int(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*4)) == 2) + 2
	} else {
		result = bool2int(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*4)) == 2)
	}
	return result
}
func sub_4417E0(a1 *libc.WChar, a2 *byte) {
	var (
		v2 uint8
		v3 *byte
	)
	_ = v3
	var v4 [64]byte
	var v5 [64]libc.WChar
	if a2 != nil {
		libc.StrCpy(&v4[0], a2)
		v2 = *memmap.PtrUint8(0x587000, 103296)
		v3 = &v4[libc.StrLen(&v4[0])]
		*(*uint32)(unsafe.Pointer(v3)) = *memmap.PtrUint32(0x587000, 103292)
		*(*byte)(unsafe.Add(unsafe.Pointer(v3), 4)) = byte(v2)
		nox_swprintf(&v5[0], libc.CWString("%-20.20S\t\t"), &v4[0])
		nox_wcscat(a1, &v5[0])
	}
}
func nox_xxx_doExecrul_4438A0(a1 int32) int32 {
	var (
		v1 *File
		v2 *File
		v3 *byte
		v5 [256]byte
		v6 [256]byte
		v7 [128]libc.WChar
	)
	if a1 == 0 {
		return 0
	}
	v6[0] = 0
	nox_sprintf(&v6[0], CString("%S"), a1)
	v1 = nox_fs_open_text(&v6[0])
	v2 = v1
	if v1 == nil {
		return 0
	}
	if !nox_fs_feof(v1) {
		for {
			libc.MemSet(unsafe.Pointer(&v5[0]), 0, 252)
			*(*uint16)(unsafe.Pointer(&v5[252])) = 0
			v5[254] = 0
			nox_fs_fgets(v2, &v5[0], math.MaxUint8)
			v3 = libc.StrChr(&v5[0], 10)
			if v3 != nil {
				*v3 = 0
			}
			if v5[0] != 0 {
				nox_swprintf(&v7[0], libc.CWString("%S"), &v5[0])
				nox_gui_console_Printf_450C00(uint8(int8(NOX_CONSOLE_WHITE)), (*libc.WChar)(memmap.PtrOff(0x587000, 106956)), &v7[0])
				nox_server_parseCmdText_443C80(&v7[0], 1)
			}
			if nox_fs_feof(v2) {
				break
			}
		}
	}
	nox_fs_close(v2)
	return 1
}
func sub_444D30() {
	libc.MemCpy(unsafe.Pointer(nox_draw_curDrawData_3799572), unsafe.Pointer(&obj_5D4594_3799660), int(unsafe.Sizeof(nox_render_data_t{})))
}
func sub_444D50(a1 *nox_render_data_t) {
	libc.MemCpy(unsafe.Pointer(a1), unsafe.Pointer(nox_draw_curDrawData_3799572), int(unsafe.Sizeof(nox_render_data_t{})))
}
func sub_444D70(a1 *nox_render_data_t) {
	libc.MemCpy(unsafe.Pointer(nox_draw_curDrawData_3799572), unsafe.Pointer(a1), int(unsafe.Sizeof(nox_render_data_t{})))
}
func sub_445440() int32 {
	return 1
}
func sub_445450() *libc.WChar {
	var (
		v0     *uint8
		result *libc.WChar
	)
	v0 = (*uint8)(memmap.PtrOff(6112660, 824440))
	for {
		result = nox_wcscpy((*libc.WChar)(unsafe.Add(unsafe.Pointer((*libc.WChar)(unsafe.Pointer(v0))), -int(unsafe.Sizeof(libc.WChar(0))*318))), (*libc.WChar)(memmap.PtrOff(6112660, 825740)))
		*(*uint32)(unsafe.Pointer(v0)) = 0
		*(*uint8)(unsafe.Add(unsafe.Pointer(v0), 4)) = 0
		v0 = (*uint8)(unsafe.Add(unsafe.Pointer(v0), 644))
		if int32(uintptr(unsafe.Pointer(v0))) >= int32(uintptr(memmap.PtrOff(6112660, 826372))) {
			break
		}
	}
	dword_5d4594_825736 = 0
	return result
}
func nox_xxx_drawMessageLines_445530() int32 {
	var (
		v1     int32
		v2     int32
		v3     int32
		result int32
		v5     int32
		v6     *uint16
		v7     int32
		v8     *uint8
		v9     int32
		v10    int32
		v11    int32
		i      int32
		v13    int32
		v14    int32
	)
	v13 = 0
	var rdr *nox_draw_viewport_t = nox_draw_getViewport_437250()
	v1 = rdr.height*3/4 + rdr.y1 - 15
	v2 = int32(dword_5d4594_825736)
	for i = int32(dword_5d4594_825736); ; v2 = i {
		v3 = v2 * 161
		result = int32(nox_frame_xxx_2598000)
		v5 = v3 * 4
		if *memmap.PtrUint32(6112660, uint32(v5)+824440) < nox_frame_xxx_2598000 {
			break
		}
		nox_xxx_drawSetTextColor_434390(int32(nox_color_black_2650656))
		v6 = memmap.PtrUint16(6112660, uint32(v5)+0xC91FC)
		nox_xxx_drawGetStringSize_43F840(nil, (*libc.WChar)(unsafe.Pointer(v6)), &v14, nil, 0)
		v7 = nox_win_width - v14
		v8 = (*uint8)(memmap.PtrOff(0x587000, 107848))
		v9 = (nox_win_width - v14) / 2
		for {
			noxrend.DrawString(nil, (*libc.WChar)(unsafe.Pointer((*int16)(unsafe.Pointer(v6)))), image.Pt(int32(uint32(v9)+*(*uint32)(unsafe.Pointer(v8))), int32(uint32(v1)+*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v8))), unsafe.Sizeof(uint32(0))*1))))))
			v8 = (*uint8)(unsafe.Add(unsafe.Pointer(v8), 8))
			if int32(uintptr(unsafe.Pointer(v8))) >= int32(uintptr(memmap.PtrOff(0x587000, 107880))) {
				break
			}
		}
		v10 = v13
		if v13 != 0 {
			nox_xxx_drawSetTextColor_434390(*memmap.PtrInt32(6112660, 2597996))
		} else {
			nox_xxx_drawSetTextColor_434390(int32(nox_color_white_2523948))
		}
		noxrend.DrawString(nil, (*libc.WChar)(unsafe.Pointer((*int16)(unsafe.Pointer(v6)))), image.Pt(v7/2, v1))
		v11 = int32(-4 - nox_xxx_guiFontHeightMB_43F320(nil))
		result = i
		v1 += v11
		if i != 0 {
			i--
		} else {
			i = 2
		}
		v13++
		if v10+1 >= 3 {
			break
		}
	}
	return result
}
func nox_xxx_guiChatMode_4456E0(a1 *int32) int32 {
	var (
		v1 *int32
		v2 int32
		v4 int32
	)
	v1 = a1
	nox_client_wndGetPosition_46AA60((*nox_window)(unsafe.Pointer(a1)), (*uint32)(unsafe.Pointer(&a1)), (*uint32)(unsafe.Pointer(&v4)))
	v2 = *(*int32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(int32(0))*25))
	a1 = (*int32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(a1))), *(*int32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(int32(0))*24))))))
	nox_client_drawImageAt_47D2C0((*nox_video_bag_image_t)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(int32(0))*15))))), int32(uintptr(unsafe.Pointer(a1))), v2+v4)
	return 1
}
func nox_xxx_guiChatShowHide_445730(a1 int32) int32 {
	return (*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_825744))))).SetHidden(a1)
}
func sub_445750() int32 {
	return bool2int(dword_5d4594_825744 != 0 && (*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_825744))))).Flags().IsHidden() == 0)
}
func sub_445770() int32 {
	var result int32
	result = (*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_825744)))).Destroy()
	dword_5d4594_825744 = 0
	return result
}
func nox_xxx_quitDialogYes_445B20() {
	nox_client_quit_4460C0()
	sub_445C40()
}
func nox_xxx_quitDialogNo_445B30() int32 {
	return nox_xxx_wndSetCaptureMain(nox_wnd_quitMenu_825760)
}
func sub_445BA0() {
	sub_413A00(0)
}
func sub_445BB0() int32 {
	return 1
}
func nox_xxx_wndDrawQuitMenu_445BC0(a1 *uint32) int32 {
	var (
		xLeft int32
		yTop  int32
	)
	nox_client_wndGetPosition_46AA60((*nox_window)(unsafe.Pointer(a1)), (*uint32)(unsafe.Pointer(&xLeft)), (*uint32)(unsafe.Pointer(&yTop)))
	nox_client_drawRectFilledAlpha_49CF10(xLeft, yTop, int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*2))), int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*3))))
	return 1
}
func sub_445C00() {
	var v int32 = nox_xxx_wndGetFlags_46ADA0(int32(uintptr(unsafe.Pointer(nox_wnd_quitMenu_825760))))
	if v&16 != 0 {
		sub_445C40()
	}
}
func sub_445C20() {
	var v int32 = nox_xxx_wndGetFlags_46ADA0(int32(uintptr(unsafe.Pointer(nox_wnd_quitMenu_825760))))
	if (v & 16) == 0 {
		sub_445C40()
	}
}
func sub_445FF0() *uint32 {
	var (
		i      int32
		result *uint32
	)
	if nox_wnd_quitMenu_825760 != nil {
		nox_wnd_quitMenu_825760.draw_data.bg_color = nox_color_black_2650656
	}
	for i = 9001; i <= 9006; i++ {
		result = (*uint32)(unsafe.Pointer(nox_wnd_quitMenu_825760.ChildByID(i)))
		if result != nil {
			*(*uint32)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof(uint32(0))*26)) = nox_color_orange_2614256
		}
	}
	return result
}
func sub_446030() int32 {
	return bool2int(dword_5d4594_825768 != 0)
}
func sub_446050() int32 {
	var result int32
	result = int32(nox_frame_xxx_2598000)
	dword_5d4594_825768 = nox_frame_xxx_2598000
	return result
}
func sub_446060() {
	dword_5d4594_825768 = 0
}
func sub_446070() {
	if func() uint32 {
		p := &dword_5d4594_825752
		*p--
		return *p
	}() == 0 {
		sub_446380()
	}
}
func sub_446090() int32 {
	return bool2int(dword_5d4594_825752 == 0)
}
func sub_4460A0(a1 int32) int32 {
	var result int32
	result = a1
	*memmap.PtrUint32(6112660, 825756) = uint32(a1)
	return result
}
func sub_4460B0() int32 {
	return int32(*memmap.PtrUint32(6112660, 825756))
}
func nox_client_quit_4460C0() int32 {
	var result int32
	if noxflags.HasGame(noxflags.GameModeQuest) {
		if noxflags.HasGame(noxflags.GameHost) {
			result = sub_4DCD40()
		} else if sub_4460B0() != 0 {
			sub_4460A0(0)
			result = sub_446140()
		} else {
			nox_xxx_netSavePlayer_41CE00()
			result = sub_4460A0(1)
		}
	} else {
		dword_5d4594_825764 = 1
		result = bool2int(noxflags.HasGame(noxflags.GameModeQuest))
		if result != 0 {
			result = bool2int(noxflags.HasGame(noxflags.GameHost))
			if result != 0 {
				result = sub_4D6B10(0)
			}
		}
	}
	return result
}
func sub_446140() int32 {
	var result int32
	dword_5d4594_825764 = 1
	result = bool2int(noxflags.HasGame(noxflags.GameModeQuest))
	if result != 0 {
		result = bool2int(noxflags.HasGame(noxflags.GameHost))
		if result != 0 {
			result = sub_4D6B10(0)
		}
	}
	return result
}
func nox_xxx_serverIsClosing_446180() int32 {
	return int32(dword_5d4594_825764)
}
func nox_gui_xxx_check_446360() uint32 {
	var result uint32
	if nox_wnd_quitMenu_825760 != nil {
		result = (uint32(^nox_xxx_wndGetFlags_46ADA0(int32(uintptr(unsafe.Pointer(nox_wnd_quitMenu_825760))))) >> 4) & 1
	} else {
		result = 0
	}
	return result
}
func sub_446380() {
	sub_44A400()
	if noxflags.HasGame(noxflags.GameOnline) && !noxflags.HasGame(noxflags.GameModeQuest) && sub_4D6F30() == 0 {
		sub_41CEE0(int32(uintptr(memmap.PtrOff(0x85B3FC, 10980))), 1)
	}
	nox_xxx_guiServerOptionsHide_4597E0(1)
	nox_xxx_setContinueMenuOrHost_43DDD0(0)
	nox_game_exit_xxx_43DE60()
	sub_446060()
}
func nox_motd_4463E0(a1 int32) int32 {
	var (
		result *File
		v2     *File
		v3     *byte
	)
	dword_5d4594_826036 = 0
	*memmap.PtrUint32(6112660, uint32(a1*4)+826040) = 0
	result = nox_fs_open(CString("motd.txt"))
	v2 = result
	if result != nil {
		*memmap.PtrUint32(6112660, uint32(a1*4)+826040) = uint32(nox_fs_fsize(v2))
		v3 = (*byte)(alloc.Calloc(int(*memmap.PtrUint32(6112660, uint32(a1*4)+826040)+1), 1))
		dword_5d4594_826036 = uint32(uintptr(unsafe.Pointer(v3)))
		if v3 != nil {
			nox_binfile_fread_raw_40ADD0(v3, *memmap.PtrUint32(6112660, uint32(a1*4)+826040), 1, v2)
			v3 = *(**byte)(unsafe.Pointer(&dword_5d4594_826036))
		}
		*(*byte)(unsafe.Add(unsafe.Pointer(v3), func() uint32 {
			p := memmap.PtrUint32(6112660, uint32(a1*4)+826040)
			x := *p
			*p++
			return x
		}())) = 0
		nox_fs_close(v2)
	}
	return 0
}
func sub_446490(a1 int32) unsafe.Pointer {
	var result unsafe.Pointer
	result = *(*unsafe.Pointer)(unsafe.Pointer(&dword_5d4594_826036))
	if dword_5d4594_826036 != 0 {
		*(*unsafe.Pointer)(unsafe.Pointer(&dword_5d4594_826036)) = nil
		result = unsafe.Pointer(uintptr(a1))
	}
	dword_5d4594_826036 = 0
	*memmap.PtrUint32(6112660, uint32(a1)*4+826040) = 0
	return result
}
func sub_4464D0(a1 int32, a2 *uint32) unsafe.Pointer {
	*a2 = *memmap.PtrUint32(6112660, uint32(a1*4)+826040)
	return *(*unsafe.Pointer)(unsafe.Pointer(&dword_5d4594_826036))
}
func sub_4464F0(a1 int32, a2 *uint32) int32 {
	if *memmap.PtrUint32(6112660, uint32(a1*4)+0xC9AC8) == 0 {
		return 0
	}
	*a2 = *memmap.PtrUint32(6112660, uint32(a1*4)+0xC9AC0)
	return int32(*memmap.PtrUint32(6112660, uint32(a1*4)+0xC9AC8))
}
func sub_446520(a1 int32, a2 unsafe.Pointer, a3 int32) {
	var v3 unsafe.Pointer
	if a2 != nil {
		if a3 > 0 {
			v3 = alloc.Calloc(int(a3), 1)
			*memmap.PtrUint32(6112660, uint32(a1*4)+0xC9AC8) = uint32(uintptr(v3))
			libc.MemCpy(v3, a2, int(a3))
			*memmap.PtrUint32(6112660, uint32(a1*4)+0xC9AC0) = uint32(a3)
			*memmap.PtrUint32(6112660, uint32(a1*4)+0xC9AD0) = 1
		}
	}
}
func sub_446580(a1 int32) int32 {
	var result int32
	result = int32(*memmap.PtrUint32(6112660, uint32(a1*4)+0xC9AC8))
	if result != 0 {
		*(*unsafe.Pointer)(memmap.PtrOff(6112660, uint32(a1*4)+0xC9AC8)) = nil
		*memmap.PtrUint32(6112660, uint32(a1*4)+0xC9AC8) = 0
		*memmap.PtrUint32(6112660, uint32(a1*4)+0xC9AC0) = 0
		*memmap.PtrUint32(6112660, uint32(a1*4)+0xC9AD0) = 0
	}
	return result
}
func nox_xxx_guiMotdLoad_4465C0() int32 {
	var (
		v0 *uint32
		v1 *uint32
		v2 *uint32
		v3 *byte
		v4 *uint32
		v5 *uint32
		v7 *byte
		v8 *uint32
	)
	v0 = (*uint32)(unsafe.Pointer(newWindowFromFile("motd.wnd", unsafe.Pointer(cgoFuncAddr(sub_4466C0)))))
	dword_5d4594_826028 = uint32(uintptr(unsafe.Pointer(v0)))
	v1 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(v0)).ChildByID(4203)))
	dword_5d4594_826032 = uint32(uintptr(unsafe.Pointer(v1)))
	v2 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*8)))))
	v7 = (*byte)(unsafe.Pointer(nox_xxx_gLoadImg("UISlider")))
	v3 = (*byte)(unsafe.Pointer(nox_xxx_gLoadImg("UISliderLit")))
	v4 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_826028)))).ChildByID(4204)))
	v5 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_826028)))).ChildByID(4205)))
	v8 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_826028)))).ChildByID(4206)))
	*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*100)) + 8))) = 16
	*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*100)) + 12))) = 10
	sub_4B5700(int32(uintptr(unsafe.Pointer(v4))), 0, 0, int32(uintptr(unsafe.Pointer(v7))), int32(uintptr(unsafe.Pointer(v3))), int32(uintptr(unsafe.Pointer(v3))))
	nox_xxx_wnd_46B280(int32(uintptr(unsafe.Pointer(v4))), *(*int32)(unsafe.Pointer(&dword_5d4594_826032)))
	nox_xxx_wnd_46B280(int32(uintptr(unsafe.Pointer(v5))), *(*int32)(unsafe.Pointer(&dword_5d4594_826032)))
	nox_xxx_wnd_46B280(int32(uintptr(unsafe.Pointer(v8))), *(*int32)(unsafe.Pointer(&dword_5d4594_826032)))
	*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*9)) = uint32(uintptr(unsafe.Pointer(v4)))
	*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*7)) = uint32(uintptr(unsafe.Pointer(v5)))
	*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*8)) = uint32(uintptr(unsafe.Pointer(v8)))
	return int32(dword_5d4594_826028)
}
func sub_4466C0(a1 int32, a2 int32, a3 int32, a4 int32) int32 {
	if a2 == 0x4007 {
		clientPlaySoundSpecial(766, 100)
		sub_446780()
	}
	return 0
}
func sub_4466F0(a1 *byte, a2 *uint8) *byte {
	var (
		v2     *byte
		v3     *byte
		v4     int8
		result *byte
	)
	v2 = (*byte)(unsafe.Pointer(a2))
	v3 = a1
	*a2 = 0
	if *a1 == 0 {
		return nil
	}
	for {
		v4 = int8(*v3)
		if *v3 == 0 {
			*v2 = 0
			return nil
		}
		if int32(v4) == 10 {
			*v2 = 0
			return (*byte)(unsafe.Add(unsafe.Pointer(v3), 1))
		}
		if int32(v4) == 13 {
			break
		}
		*func() *byte {
			p := &v2
			x := *p
			*p = (*byte)(unsafe.Add(unsafe.Pointer(*p), 1))
			return x
		}() = byte(v4)
		v3 = (*byte)(unsafe.Add(unsafe.Pointer(v3), 1))
	}
	result = (*byte)(unsafe.Add(unsafe.Pointer(v3), 1))
	*v2 = 0
	if *result == 10 {
		result = (*byte)(unsafe.Add(unsafe.Pointer(result), 1))
	}
	return result
}
func nox_xxx_motdAddSomeTextMB_446730(a1 *uint8) *uint8 {
	var (
		result *uint8
		v2     [256]libc.WChar
	)
	result = a1
	if int32(*a1) != 0 {
		nox_swprintf(&v2[0], libc.CWString("%S"), a1)
		result = (*uint8)(unsafe.Pointer(uintptr((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_826032))))).Func94(asWindowEvent(0x400D, int32(uintptr(unsafe.Pointer(&v2[0]))), -1)))))
	}
	return result
}
func sub_446780() int32 {
	if (*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_826028))))).Flags().IsHidden() != 0 {
		return 0
	}
	guiFocus(nil)
	(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_826028))))).Hide()
	*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_826028 + 4))) &= 0xFFFFFFF7
	*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_826032 + 4))) &= 0xFFFFFFF7
	(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_826032))))).Func94(asWindowEvent(0x400F, 0, 0))
	return 1
}
func nox_xxx_motd_4467F0() {
	var (
		result uint32
		v1     *uint32
		v2     *byte
		v3     *uint32
		v4     [256]byte
	)
	result = nox_gui_xxx_check_446360()
	if result == 0 {
		result = uint32(nox_xxx_isQuest_4D6F50())
		if result == 0 || (func() uint32 {
			result = uint32(bool2int(noxflags.HasGame(noxflags.GameModeChat)))
			return result
		}()) == 0 {
			if !nox_common_getEngineFlag(nox_engine_flag(NOX_ENGINE_FLAG_DISABLE_GRAPHICS_RENDERING)) {
				result = uint32(nox_xxx_wndGetFlags_46ADA0(*(*int32)(unsafe.Pointer(&dword_5d4594_826028))))
				if result&16 != 0 {
					result = uint32(sub_44A4A0())
					if result == 0 {
						result = uint32(sub_49C810())
						if result == 0 {
							result = uint32(sub_49CB40())
							if result == 0 {
								nox_xxx_wndShowModalMB((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_826028))))))
								v1 = (*uint32)(unsafe.Pointer(GUIChildByID(4100)))
								if v1 != nil {
									nox_xxx_wndShowModalMB((*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v1)))))))
								}
								*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_826028 + 4))) |= 8
								*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_826032 + 4))) |= 8
								v2 = *(**byte)(memmap.PtrOff(6112660, 826060))
								if *memmap.PtrUint32(6112660, 826060) != 0 {
									for {
										v2 = sub_4466F0(v2, (*uint8)(unsafe.Pointer(&v4[0])))
										if v2 == nil {
											break
										}
										if v4[0] == 0 {
											libc.StrCpy(&v4[0], CString(" "))
										}
										nox_xxx_motdAddSomeTextMB_446730((*uint8)(unsafe.Pointer(&v4[0])))
									}
									if v4[0] != 0 {
										nox_xxx_motdAddSomeTextMB_446730((*uint8)(unsafe.Pointer(&v4[0])))
									}
								}
								v3 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_826028)))).ChildByID(4202)))
								guiFocus((*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v3)))))))
							}
						}
					}
				}
			}
			*memmap.PtrUint32(6112660, 826068) = 0
		}
	}
}
func sub_446940(a1 int32) int32 {
	return int32(*memmap.PtrUint32(6112660, uint32(a1*4)+0xC9AD0))
}
func sub_446950() int32 {
	return bool2int(dword_5d4594_826028 != 0 && (nox_xxx_wndGetFlags_46ADA0(*(*int32)(unsafe.Pointer(&dword_5d4594_826028)))&16) == 0)
}
func sub_4469D0(a1 int32) *libc.WChar {
	if uint32(a1) / *memmap.PtrUint32(6112660, 826092) + 1 <= 1 {
		return *(**libc.WChar)(memmap.PtrOff(0x587000, uint32(a1*8)+0x1A930))
	}
	nox_swprintf((*libc.WChar)(memmap.PtrOff(6112660, 827140)), libc.CWString("%s %d"), *memmap.PtrUint32(0x587000, (uint32(a1)%*memmap.PtrUint32(6112660, 826092))*8+0x1A930), uint32(a1) / *memmap.PtrUint32(6112660, 826092) + 1)
	return (*libc.WChar)(memmap.PtrOff(6112660, 827140))
}
func sub_446A90() int32 {
	var result int32
	nox_wnd_xxx_829520.state = nox_gui_anim_state(NOX_GUI_ANIM_OUT)
	sub_43BE40(2)
	clientPlaySoundSpecial(923, 100)
	result = 1
	nox_wnd_xxx_829520.field_13 = nox_game_showGameSel_4379F0
	return result
}
func sub_446AD0(fnc func() int32) {
	nox_wnd_xxx_829520.field_13 = fnc
}
func sub_446BC0(a1 int32) *byte {
	var (
		v1     int32
		result *byte
	)
	v1 = (*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_829488))))).Func94(asWindowEvent(0x4016, a1, 0))
	nox_sprintf((*byte)(memmap.PtrOff(6112660, 826072)), CString("%S"), v1)
	result = (*byte)(memmap.PtrOff(6112660, 826072))
	if int32(*memmap.PtrUint8(6112660, 826072)) == 64 || int32(*memmap.PtrUint8(6112660, 826072)) == 32 {
		result = (*byte)(memmap.PtrOff(6112660, 826073))
	}
	return result
}
func sub_446C10(a1 int32, a2 int32) int32 {
	var (
		v2     int32
		result int32
	)
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_829488 + 32))))
	result = a1
	if a1 < int32(*(*int16)(unsafe.Pointer(uintptr(v2 + 44)))) {
		result = a1 * 131
		if a2 != 0 {
			*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 24))) + uint32(a1*524) + 516))) = **(**uint32)(memmap.PtrOff(0x85B3FC, 140))
		} else {
			*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 24))) + uint32(a1*524) + 516))) = **(**uint32)(memmap.PtrOff(0x85B3FC, 188))
		}
	}
	return result
}
func sub_446C70() int32 {
	return (*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_829488))))).Func94(asWindowEvent(0x4014, 0, 0))
}
func sub_446C90() int32 {
	return bool2int(*(*int32)(unsafe.Pointer(uintptr((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_829488))))).Func94(asWindowEvent(0x4014, 0, 0))))) != -1)
}
func sub_446CC0(a1 int32) *libc.WChar {
	var (
		result *libc.WChar
		v2     *byte
		v3     *byte
	)
	result = (*libc.WChar)(unsafe.Pointer(uintptr(sub_41E2F0())))
	if uintptr(unsafe.Pointer(result)) == uintptr(7) {
		if a1 != 0 {
			v2 = sub_41FA40()
			nox_swprintf((*libc.WChar)(memmap.PtrOff(6112660, 829224)), *(**libc.WChar)(memmap.PtrOff(6112660, 829528)), v2, a1)
		} else {
			result = *(**libc.WChar)(unsafe.Pointer(&dword_5d4594_829532))
			if dword_5d4594_829532 == 0 {
				return result
			}
			v3 = sub_41FA40()
			nox_swprintf((*libc.WChar)(memmap.PtrOff(6112660, 829224)), *(**libc.WChar)(unsafe.Pointer(&dword_5d4594_829532)), v3)
			(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_829484))))).Func94(asWindowEvent(0x4013, -1, 0))
		}
		result = (*libc.WChar)(unsafe.Pointer(uintptr((*nox_window)(unsafe.Pointer(uintptr(*memmap.PtrInt32(6112660, 829496)))).Func94(asWindowEvent(0x4001, int32(uintptr(memmap.PtrOff(6112660, 829224))), 0)))))
	}
	return result
}
func sub_446D50() {
	var (
		v0 *byte
		v1 int32
	)
	v1 = 0
	v0 = (*byte)(unsafe.Pointer(uintptr(sub_4464F0(0, (*uint32)(unsafe.Pointer(&v1))))))
	sub_446D80(v0)
}
func sub_446D80(a1 *byte) {
	var (
		v1 *byte
		v2 [256]byte
		v3 [256]libc.WChar
	)
	v1 = a1
	if a1 != nil {
		for {
			v1 = sub_446E20(v1, (*uint8)(unsafe.Pointer(&v2[0])))
			if v1 == nil {
				break
			}
			if v2[0] == 0 {
				libc.StrCpy(&v2[0], CString(" "))
			}
			nox_swprintf(&v3[0], libc.CWString("%S"), &v2[0])
			sub_447310(0, int32(uintptr(unsafe.Pointer(&v3[0]))))
		}
		if v2[0] != 0 {
			nox_swprintf(&v3[0], libc.CWString("%S"), &v2[0])
			sub_447310(0, int32(uintptr(unsafe.Pointer(&v3[0]))))
		}
	}
}
func sub_446E20(a1 *byte, a2 *uint8) *byte {
	var (
		v2     *byte
		v3     *byte
		v4     int8
		result *byte
	)
	v2 = (*byte)(unsafe.Pointer(a2))
	v3 = a1
	*a2 = 0
	if *a1 == 0 {
		return nil
	}
	for {
		v4 = int8(*v3)
		if *v3 == 0 {
			*v2 = 0
			return nil
		}
		if int32(v4) == 10 {
			*v2 = 0
			return (*byte)(unsafe.Add(unsafe.Pointer(v3), 1))
		}
		if int32(v4) == 13 {
			break
		}
		*func() *byte {
			p := &v2
			x := *p
			*p = (*byte)(unsafe.Add(unsafe.Pointer(*p), 1))
			return x
		}() = byte(v4)
		v3 = (*byte)(unsafe.Add(unsafe.Pointer(v3), 1))
	}
	result = (*byte)(unsafe.Add(unsafe.Pointer(v3), 1))
	*v2 = 0
	if *result == 10 {
		result = (*byte)(unsafe.Add(unsafe.Pointer(result), 1))
	}
	return result
}
func sub_446E60() *libc.WChar {
	var result *libc.WChar
	(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_829484))))).Func94(asWindowEvent(0x400F, 0, 0))
	for result = sub_41F0A0(); result != nil; result = sub_41F070() {
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_829484))))).Func94(asWindowEvent(0x400D, int32(uintptr(unsafe.Pointer(result))), 14))
	}
	return result
}
func sub_446EB0() *byte {
	var (
		result *byte
		i      *byte
		v2     [256]libc.WChar
	)
	(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_829488))))).Func94(asWindowEvent(0x400F, 0, 0))
	result = sub_41F780()
	for i = result; result != nil; i = result {
		nox_swprintf(&v2[0], libc.CWString("%S"), i)
		if sub_427880(i) != 0 {
			(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_829488))))).Func94(asWindowEvent(0x400D, int32(uintptr(unsafe.Pointer(&v2[0]))), 2))
		} else {
			(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_829488))))).Func94(asWindowEvent(0x400D, int32(uintptr(unsafe.Pointer(&v2[0]))), 14))
		}
		result = sub_41F710()
	}
	return result
}
func sub_446F40(a1 int32) *byte {
	var v1 int32
	if a1 == -1 {
		return nil
	}
	v1 = (*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_829488))))).Func94(asWindowEvent(0x4016, a1, 0))
	nox_sprintf((*byte)(memmap.PtrOff(6112660, 827120)), CString("%S"), v1)
	return (*byte)(memmap.PtrOff(6112660, 827120))
}
func sub_4471A0(a1 *byte, a2 int32, a3 int32, a4 int32) *libc.WChar {
	var result *libc.WChar
	result = (*libc.WChar)(unsafe.Pointer(uintptr(sub_427880(a1))))
	if result == nil {
		result = *(**libc.WChar)(unsafe.Pointer(&dword_5d4594_829492))
		if dword_5d4594_829492 != 0 {
			result = *(**libc.WChar)(unsafe.Pointer(&dword_5d4594_829544))
			if dword_5d4594_829544 != 0 {
				if a2 != 0 {
					if a3 != 0 {
						if a4 == 1 {
							nox_swprintf((*libc.WChar)(memmap.PtrOff(6112660, 827176)), *(**libc.WChar)(memmap.PtrOff(6112660, 829552)), a1, a2)
							return (*libc.WChar)(unsafe.Pointer(uintptr((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_829492))))).Func94(asWindowEvent(0x400D, int32(uintptr(memmap.PtrOff(6112660, 827176))), 7)))))
						}
						result = *(**libc.WChar)(memmap.PtrOff(6112660, 829548))
					}
					nox_swprintf((*libc.WChar)(memmap.PtrOff(6112660, 827176)), result, a1, a2)
					result = (*libc.WChar)(unsafe.Pointer(uintptr((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_829492))))).Func94(asWindowEvent(0x400D, int32(uintptr(memmap.PtrOff(6112660, 827176))), bool2int(a3 != 0)+6)))))
				}
			}
		}
	}
	return result
}
func sub_447250(a1 *byte, a2 int32, a3 int32, a4 int32) *libc.WChar {
	var (
		result *libc.WChar
		v5     *libc.WChar
		v6     *byte
	)
	result = (*libc.WChar)(unsafe.Pointer(uintptr(sub_427880(a1))))
	if result == nil {
		result = *(**libc.WChar)(unsafe.Pointer(&dword_5d4594_829492))
		if dword_5d4594_829492 != 0 {
			if *memmap.PtrUint32(6112660, 829556) != 0 {
				result = *(**libc.WChar)(memmap.PtrOff(6112660, 829560))
				if *memmap.PtrUint32(6112660, 829560) != 0 {
					if a2 != 0 {
						if a3 != 0 {
							if a4 == 1 {
								v5 = *(**libc.WChar)(memmap.PtrOff(6112660, 829564))
								v6 = sub_41FA40()
								nox_swprintf((*libc.WChar)(memmap.PtrOff(6112660, 827176)), v5, a1, v6, a2)
								return (*libc.WChar)(unsafe.Pointer(uintptr((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_829492))))).Func94(asWindowEvent(0x400D, int32(uintptr(memmap.PtrOff(6112660, 827176))), 13)))))
							}
						} else {
							result = *(**libc.WChar)(memmap.PtrOff(6112660, 829556))
						}
						nox_swprintf((*libc.WChar)(memmap.PtrOff(6112660, 827176)), result, a1, a2)
						result = (*libc.WChar)(unsafe.Pointer(uintptr((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_829492))))).Func94(asWindowEvent(0x400D, int32(uintptr(memmap.PtrOff(6112660, 827176))), 13)))))
					}
				}
			}
		}
	}
	return result
}
func sub_447310(a1 int32, a2 int32) int32 {
	var result int32
	result = int32(dword_5d4594_829492)
	if dword_5d4594_829492 != 0 {
		if *memmap.PtrUint32(6112660, 829536) != 0 {
			result = a2
			if a2 != 0 {
				if a1 != 0 {
					nox_swprintf((*libc.WChar)(memmap.PtrOff(6112660, 827176)), *(**libc.WChar)(memmap.PtrOff(6112660, 829536)), a1, a2)
				} else {
					nox_swprintf((*libc.WChar)(memmap.PtrOff(6112660, 827176)), libc.CWString("%s"), a2)
				}
				result = (*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_829492))))).Func94(asWindowEvent(0x400D, int32(uintptr(memmap.PtrOff(6112660, 827176))), 9))
			}
		}
	}
	return result
}
func sub_447380(a1 *byte, a2 *byte) {
	var v3 *libc.WChar
	if sub_427880(a1) != 0 {
		return
	}
	if dword_5d4594_829492 == 0 {
		return
	}
	if dword_5d4594_829544 == 0 {
		return
	}
	if a2 == nil {
		return
	}
	v3 = *(**libc.WChar)(memmap.PtrOff(6112660, 829540))
	sub_447410(a2)
	if a1 != nil {
		nox_swprintf((*libc.WChar)(memmap.PtrOff(6112660, 827176)), v3, a1, a2)
	} else {
		nox_swprintf((*libc.WChar)(memmap.PtrOff(6112660, 827176)), libc.CWString("%S"), a2)
	}
	(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_829492))))).Func94(asWindowEvent(0x400D, int32(uintptr(memmap.PtrOff(6112660, 827176))), 4))
	clientPlaySoundSpecial(226, 100)
}
func sub_447410(a1 *byte) uint32 {
	var (
		v1     int32
		v2     *byte
		v3     int8
		result uint32
	)
	v1 = 0
	v2 = a1
	if *a1 != 0 {
		for {
			if int32(uint8(*v2)) < 128 {
				*memmap.PtrUint8(6112660, uint32(func() int32 {
					p := &v1
					x := *p
					*p++
					return x
				}())+0xC9AF0) = uint8(*v2)
			}
			v3 = int8(*func() *byte {
				p := &v2
				*p = (*byte)(unsafe.Add(unsafe.Pointer(*p), 1))
				return *p
			}())
			if int32(v3) == 0 {
				break
			}
		}
	}
	*memmap.PtrUint8(6112660, uint32(v1)+0xC9AF0) = 0
	result = uint32(libc.StrLen((*byte)(memmap.PtrOff(6112660, 826096))) + 1)
	libc.MemCpy(unsafe.Pointer(a1), memmap.PtrOff(6112660, 826096), int(result))
	return result
}
func sub_4475E0() int32 {
	return (*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_829488))))).Func94(asWindowEvent(0x400F, 0, 0))
}
func sub_447600() {
	if sub_44A4A0() == 0 {
		guiFocus((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_829500))))))
	}
}
func sub_447BD0() int32 {
	var v0 func() int32
	v0 = nox_wnd_xxx_829520.field_13
	nox_gui_freeAnimation_43C570(nox_wnd_xxx_829520)
	sub_448490()
	v0()
	return 1
}
func sub_447BF0(a1 *uint32, a2 int32, a3 uint32, a4 int32) int32 {
	var result int32
	if a2 == 19 {
		(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1)))))).Func94(asWindowEvent(0x4007, int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*8)) + 28)))), 0))
		result = 0
	} else if a2 == 20 {
		(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1)))))).Func94(asWindowEvent(0x4007, int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*8)) + 32)))), 0))
		result = 0
	} else if *(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*8)) + 16))) != 0 {
		result = nox_xxx_wndListboxProcWithData10_4A2DE0(int32(uintptr(unsafe.Pointer(a1))), a2, a3, a4)
	} else {
		result = nox_xxx_wndListboxProcWithoutData10_4A28E0(a1, a2, a3, a4)
	}
	return result
}
func sub_447C70(a1 int32, a2 int32, a3 int32, a4 int32) int32 {
	var v5 int32
	if a2 != 21 {
		return 0
	}
	if a3 != 1 {
		return 0
	}
	if a4 == 2 && sub_41E2F0() == 7 {
		sub_41E300(9)
		v5 = sub_41E2F0()
		sub_41DA70(v5, 8)
		sub_44A400()
		nox_game_checkStateWol_43C260()
	}
	return 1
}
func sub_447CC0(a1 int32, a2 uint32, a3 *int32, a4 int32) int32 {
	var (
		result int32
		v5     *libc.WChar
		v6     *libc.WChar
		v7     *libc.WChar
		v9     int32
		v10    int32
		v11    int32
		v12    *uint32
		v13    *int32
		v14    *libc.WChar
		v15    int32
		v16    int32
		v17    int32
		v18    int32
		v19    int32
		v20    *libc.WChar
		v21    *int32 = nil
		File   [256]byte
	)
	if a2 > 16390 {
		switch a2 {
		case 0x4007:
			switch (*nox_window)(unsafe.Pointer(a3)).ID() {
			case 1902:
				sub_448730_wol_dialogs()
			case 1903:
				v16 = sub_41E2F0()
				sub_41DA70(v16, 12)
				v17 = sub_41E2F0()
				sub_41DA70(v17, 13)
			case 1904:
				sub_449280()
			case 1905:
				sub_40D380()
				sub_41E300(8)
				sub_4207F0(5)
			case 1906:
				sub_41E300(9)
				v15 = sub_41E2F0()
				sub_41DA70(v15, 8)
				nox_game_checkStateWol_43C260()
			case 1913:
				v20 = (*libc.WChar)(unsafe.Pointer(uintptr((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_829500))))).Func94(asWindowEvent(0x401D, 0, 0)))))
				sub_447090_wol_chat(v20)
				(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_829500))))).Func94(asWindowEvent(0x401E, int32(uintptr(memmap.PtrOff(6112660, 829580))), 0))
			case 1914:
				nox_xxx_guiServerListLoad_449530()
			case 1926:
				v18 = sub_4200F0()
				nox_sprintf(&File[0], CString("http://battleclans.westwood.com/nox/index_%d.html"), v18)
				compatShellExecuteA(0, nil, &File[0], nil, (*byte)(memmap.PtrOff(0x587000, 111564)), 3)
			case 1927:
				v19 = sub_4200F0()
				nox_sprintf(&File[0], CString("http://www.westwood.com/westwoodonline/tournaments/nox/rankindex_%d.html"), v19)
				compatShellExecuteA(0, nil, &File[0], nil, (*byte)(memmap.PtrOff(0x587000, 111644)), 3)
			default:
			}
			clientPlaySoundSpecial(921, 100)
			return 1
		case 16400:
			if (*nox_window)(unsafe.Pointer(a3)).ID() != 1908 {
				return 0
			}
			if int32(*(*uint16)(unsafe.Pointer(uintptr((*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a3)))))).Func94(asWindowEvent(0x4016, a4, 0)))))) == 0 {
				goto LABEL_23
			}
			v5 = (*libc.WChar)(unsafe.Pointer(uintptr(sub_41EC00())))
			v6 = (*libc.WChar)(unsafe.Pointer(uintptr(sub_41F1E0(a4))))
			v7 = v6
			if v6 == nil {
				goto LABEL_23
			}
			if v5 != nil {
				if nox_wcscmp(v6, v5) != 0 {
					sub_40D380()
					sub_446A20_wol_chat(v7)
				}
			LABEL_23:
				sub_447600()
				result = 0
			} else {
				sub_446A20_wol_chat(nil)
				sub_447600()
				result = 0
			}
		case 0x4011:
			if (*nox_window)(unsafe.Pointer(a3)).ID() != 1909 {
				return 0
			}
			if a4 == -1 || dword_5d4594_829504 != 0 {
				dword_587000_109280 = math.MaxUint32
				result = 0
			} else {
				var mpos nox_point = getMousePos()
				dword_587000_109280 = uint32(a4)
				v9 = mpos.x
				v10 = mpos.y
				v11 = nox_xxx_guiFontHeightMB_43F320(nil)
				dword_5d4594_829504 = uint32(uintptr(unsafe.Pointer(nox_window_new(nil, 40, v9-40, v10-(v11+4)/2, 80, v11+4, nil))))
				(*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_829504)))).SetAllFuncs(nil, func(arg1 *nox_window, arg2 unsafe.Pointer) int32 {
					return sub_448140(int32(uintptr(unsafe.Pointer(arg1))))
				}, nil)
				nox_xxx_wndShowModalMB((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_829504))))))
				v12 = (*uint32)(unsafe.Pointer(nox_window_new((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_829504))))), 8, 0, v11/2, 80, v11+2, nil)))
				(*nox_window)(unsafe.Pointer(v12)).SetAllFuncs(func(arg1 int32, arg2 int32, arg3 int32, arg4 int32) int32 {
					return sub_448340(arg1, uint32(arg2))
				}, func(arg1 *nox_window, arg2 unsafe.Pointer) int32 {
					return sub_448240((*uint32)(unsafe.Pointer(arg1)), (*uint8)(arg2))
				}, nil)
				*(*uint32)(unsafe.Add(unsafe.Pointer(v12), unsafe.Sizeof(uint32(0))*8)) = 0
				result = 0
			}
			return result
		case 0x401F:
			v13 = a3
			(*nox_window)(unsafe.Pointer(a3)).ID()
			goto LABEL_31
		default:
			return 0
		}
		return result
	}
	if a2 == 16390 {
		if dword_5d4594_830112 != 0 {
			return 0
		}
		sub_41E2F0()
		return 0
	}
	if a2 > 0x4003 {
		if a2 != 0x4005 {
			return 0
		}
		clientPlaySoundSpecial(920, 100)
		return 1
	}
	if a2 == 0x4003 {
		if (*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_829480)))).ChildByID(a4) == nil {
			return 0
		}
		return 0
	}
	if a2 == 23 {
		return 1
	}
	if a2 != 1910 {
		return 0
	}
	v13 = v21
LABEL_31:
	v14 = (*libc.WChar)(unsafe.Pointer(uintptr((*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v13)))))).Func94(asWindowEvent(0x401D, 0, 0)))))
	sub_446F80_wol_chat(v14)
	(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v13)))))).Func94(asWindowEvent(0x401E, int32(uintptr(memmap.PtrOff(6112660, 829576))), 0))
	return 0
}
func sub_448140(yTop int32) int32 {
	var (
		v1    int32
		xLeft int32
		v4    int32
		v5    int32
	)
	v1 = yTop
	nox_client_wndGetPosition_46AA60((*nox_window)(unsafe.Pointer(uintptr(yTop))), (*uint32)(unsafe.Pointer(&xLeft)), (*uint32)(unsafe.Pointer(&yTop)))
	nox_window_get_size((*nox_window)(unsafe.Pointer(uintptr(v1))), &v4, &v5)
	nox_client_drawRectFilledAlpha_49CF10(xLeft, yTop, v4, v5)
	nox_client_drawSetColor_434460(int32(nox_color_white_2523948))
	nox_client_drawAddPoint_49F500(xLeft+1, yTop)
	nox_xxx_rasterPointRel_49F570(v4-2, 0)
	nox_client_drawLineFromPoints_49E4B0()
	nox_client_drawAddPoint_49F500(xLeft+1, yTop+v5)
	nox_xxx_rasterPointRel_49F570(v4-2, 0)
	nox_client_drawLineFromPoints_49E4B0()
	nox_client_drawAddPoint_49F500(xLeft, yTop+1)
	nox_xxx_rasterPointRel_49F570(0, v5-2)
	nox_client_drawLineFromPoints_49E4B0()
	nox_client_drawAddPoint_49F500(xLeft+v4, yTop+1)
	nox_xxx_rasterPointRel_49F570(0, v5-2)
	nox_client_drawLineFromPoints_49E4B0()
	return 1
}
func sub_448240(a1 *uint32, a2 *uint8) int32 {
	var (
		v2 *uint16
		v3 int32
		v5 int32
		v6 int32
	)
	v2 = *(**uint16)(memmap.PtrOff(0x587000, *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*8))*12+0x1A918))
	nox_client_wndGetPosition_46AA60((*nox_window)(unsafe.Pointer(a1)), (*uint32)(unsafe.Pointer(&v6)), (*uint32)(unsafe.Pointer(&a1)))
	nox_xxx_drawGetStringSize_43F840(nil, (*libc.WChar)(unsafe.Pointer(v2)), &v5, nil, 0)
	v3 = int32(nox_color_yellow_2589772)
	if (int32(*a2) & 2) == 0 {
		v3 = int32(nox_color_white_2523948)
	}
	sub_4482D0((80-v5)/2+v6+1, int32(uintptr(unsafe.Pointer(a1))), v3, int32(nox_color_black_2650656), v2)
	return 1
}
func sub_4482D0(a1 int32, a2 int32, a3 int32, a4 int32, a5 *uint16) int32 {
	var v5 *uint8
	nox_xxx_drawSetTextColor_434390(a4)
	v5 = (*uint8)(memmap.PtrOff(0x587000, 109288))
	for {
		nox_xxx_drawStringWrap_43FAF0(nil, (*libc.WChar)(unsafe.Pointer(a5)), int32(uint32(a1)+*(*uint32)(unsafe.Pointer(v5))), int32(uint32(a2)+*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v5))), unsafe.Sizeof(uint32(0))*1)))), 320, 0)
		v5 = (*uint8)(unsafe.Add(unsafe.Pointer(v5), 8))
		if int32(uintptr(unsafe.Pointer(v5))) >= int32(uintptr(memmap.PtrOff(0x587000, 109320))) {
			break
		}
	}
	nox_xxx_drawSetTextColor_434390(a3)
	return nox_xxx_drawStringWrap_43FAF0(nil, (*libc.WChar)(unsafe.Pointer(a5)), a1, a2, 320, 0)
}
func sub_448340(a1 int32, a2 uint32) int32 {
	if a2 == 5 {
		return 1
	}
	if a2 > 5 && a2 <= 7 {
		sub_448380()
		(cgoAsFunc(*(*uint32)(memmap.PtrOff(0x587000, *(*uint32)(unsafe.Pointer(uintptr(a1 + 32)))*12+0x1A920)), (*func(int32))(nil)))(*(*int32)(unsafe.Pointer(&dword_587000_109280)))
		return 1
	}
	return 0
}
func sub_448380() {
	(*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_829504)))).Destroy()
	dword_5d4594_829504 = 0
	sub_447600()
}
func sub_4483A0(a1 int32, a2 int32) int32 {
	var (
		result int32
		xLeft  int32
		yTop   int32
	)
	if dword_5d4594_829504 != 0 {
		var mpos nox_point = getMousePos()
		if !nox_xxx_wndPointInWnd_46AAB0(*(**uint32)(unsafe.Pointer(&dword_5d4594_829504)), mpos.x, mpos.y) {
			sub_448380()
		}
	}
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
func sub_448450() int32 {
	var (
		result int32
		v1     *uint32
	)
	result = bool2int(noxflags.HasGame(noxflags.GameFlag26))
	if result != 0 {
		sub_446970_wol_chat()
		v1 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_829480)))).ChildByID(1905)))
		result = (*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_829480))))).Func94(asWindowEvent(0x4007, int32(uintptr(unsafe.Pointer(v1))), 0))
	}
	return result
}
func sub_448490() int32 {
	var result int32
	result = int32(dword_5d4594_829480)
	if dword_5d4594_829480 != 0 {
		result = (*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_829480)))).Destroy()
		dword_5d4594_829480 = 0
		dword_5d4594_829484 = 0
		dword_5d4594_829488 = 0
		dword_5d4594_829492 = 0
		dword_5d4594_829500 = 0
	}
	return result
}
func sub_4484D0(a1 *libc.WChar, a2 int32, a3 int32) {
	var (
		v3 *libc.WChar
		v4 *libc.WChar
	)
	if a1 != nil {
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_829484))))).Func94(asWindowEvent(0x4012, a3, 0))
		v3 = sub_41F330(int32(uintptr(unsafe.Pointer(a1))), a2)
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_829484))))).Func94(asWindowEvent(0x400D, int32(uintptr(unsafe.Pointer(v3))), 14))
		v4 = (*libc.WChar)(unsafe.Pointer(uintptr(sub_41EC00())))
		if v4 != nil {
			if nox_wcscmp(v4, a1) == 0 {
				(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_829484))))).Func94(asWindowEvent(0x4013, a3, 0))
			}
		}
	}
}
func sub_448550(a1 int32) int32 {
	return (*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_829484))))).Func94(asWindowEvent(0x400E, a1, 0))
}
func sub_448570(a1 *byte, a2 int8, a3 int32) {
	var (
		v3 *byte
		v4 [256]libc.WChar
	)
	if a1 != nil {
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_829488))))).Func94(asWindowEvent(0x4012, a3, 0))
		v3 = sub_41F740(int32(uintptr(unsafe.Pointer(a1))), a2)
		nox_swprintf(&v4[0], libc.CWString("%S"), v3)
		if sub_427880(a1) != 0 {
			(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_829488))))).Func94(asWindowEvent(0x400D, int32(uintptr(unsafe.Pointer(&v4[0]))), 2))
		} else {
			(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_829488))))).Func94(asWindowEvent(0x400D, int32(uintptr(unsafe.Pointer(&v4[0]))), 14))
		}
	}
}
func sub_448620(a1 int32) int32 {
	var result int32
	result = a1
	if a1 >= 0 {
		result = (*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_829488))))).Func94(asWindowEvent(0x400E, a1, 0))
	}
	return result
}
func sub_448650(a1 *uint8, a2 *libc.WChar) int32 {
	var result int32
	if int32(*a1) == -16 {
		result = sub_425680((*uint8)(unsafe.Add(unsafe.Pointer(a1), 1)), (*uint8)(unsafe.Pointer(a2)))
	} else {
		result = nox_swprintf(a2, libc.CWString("%S"), a1)
	}
	return result
}
func sub_448680(a1 *libc.WChar, a2 *byte) int32 {
	var (
		result int32
		v3     uint32
		v4     [2048]byte
	)
	if nox_xxx_cliCanTalkMB_4100F0((*int16)(unsafe.Pointer(a1))) != 0 {
		return nox_sprintf(a2, CString("%S"), a1)
	}
	*(*[2048]byte)(unsafe.Pointer(&v4[0])) = [2048]byte{}
	libc.MemSet(unsafe.Pointer(a2), 0, 2048)
	v3 = nox_wcslen(a1)
	sub_425550((*uint8)(unsafe.Pointer(a1)), (*uint8)(unsafe.Pointer(&v4[0])), int32((v3+1)*16))
	result = 0
	*a2 = 240
	libc.MemCpy(unsafe.Add(unsafe.Pointer(a2), 1), unsafe.Pointer(&v4[0]), libc.StrLen(&v4[0]))
	return result
}
func sub_448F00(a1 *uint32, a2 int32, a3 int32, a4 int32) int32 {
	if a2 != 21 || a3 != 28 {
		return nox_xxx_wndEditProc_487D70((*nox_window)(unsafe.Pointer(a1)), a2, a3, a4)
	}
	if a4 == 2 {
		if dword_587000_111668 == 1 {
			if dword_5d4594_830204 != 1 {
				sub_448CF0_wol_dialogs()
				return 1
			}
		} else if dword_5d4594_830208 != 1 {
			sub_448CF0_wol_dialogs()
		}
	}
	return 1
}
func sub_448F60(a1 *uint32, a2 int32, a3 int32, a4 int32) int32 {
	if a2 != 21 || a3 != 28 {
		return nox_xxx_wndEditProc_487D70((*nox_window)(unsafe.Pointer(a1)), a2, a3, a4)
	}
	if a4 == 2 {
		if dword_587000_111668 == 1 {
			if dword_5d4594_830204 != 1 {
				sub_448CF0_wol_dialogs()
				return 1
			}
		} else if dword_5d4594_830208 != 1 {
			sub_448CF0_wol_dialogs()
		}
	}
	return 1
}
func sub_448FC0(a1 int32, a2 int32) int32 {
	var (
		xLeft int32
		yTop  int32
	)
	nox_client_wndGetPosition_46AA60((*nox_window)(unsafe.Pointer(uintptr(a1))), (*uint32)(unsafe.Pointer(&xLeft)), (*uint32)(unsafe.Pointer(&yTop)))
	if int32(int8(*(*uint8)(unsafe.Pointer(uintptr(a1 + 4))))) >= 0 {
		if *(*uint32)(unsafe.Pointer(uintptr(a2 + 20))) != 0x80000000 {
			nox_client_drawSetColor_434460(int32(nox_color_black_2650656))
			nox_client_drawRectFilledOpaque_49CE30(xLeft, yTop, int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 8)))), int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 12)))))
		}
	} else {
		nox_client_drawImageAt_47D2C0((*nox_video_bag_image_t)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a2 + 24)))))), xLeft, yTop)
	}
	if dword_587000_111668 != 0 {
		if dword_5d4594_830204 != 0 && nox_xxx_checkKeybTimeout_4160F0(18, nox_gameFPS*10) {
			nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_830148))))), 1)
			dword_5d4594_830204 = 0
			return 1
		}
	} else if dword_5d4594_830208 != 0 && nox_xxx_checkKeybTimeout_4160F0(19, nox_gameFPS*10) {
		nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_830148))))), 1)
		dword_5d4594_830208 = 0
	}
	return 1
}
func sub_449280() int32 {
	var (
		result int32
		v1     int32
		v2     int32
		v3     uint32
		v4     int32
		v5     int32
		v6     uint32
		v7     int32
		v8     int32
		v9     uint32
	)
	if dword_5d4594_830112 != 0 {
		return 0
	}
	if dword_5d4594_830116 != 0 {
		return 0
	}
	if dword_5d4594_830120 != 0 {
		return 0
	}
	result = int32(uintptr(unsafe.Pointer(newWindowFromFile("wolopt.wnd", unsafe.Pointer(cgoFuncAddr(sub_4493D0))))))
	dword_5d4594_830104 = uint32(result)
	if result != 0 {
		*(*uint32)(unsafe.Pointer(uintptr(result + 56))) = nox_color_black_2650656
		*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_830104 + 16))) = (uint32(nox_win_width) - *(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_830104 + 8)))) / 2
		*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_830104 + 24))) = *(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_830104 + 8))) + *(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_830104 + 16)))
		dword_5d4594_830192 = uint32(uintptr(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_830104)))).ChildByID(1991))))
		dword_5d4594_830196 = uint32(uintptr(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_830104)))).ChildByID(1992))))
		dword_5d4594_830200 = uint32(uintptr(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_830104)))).ChildByID(1993))))
		v1 = sub_41FF60()
		v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_830192 + 36))))
		if v1 != 0 {
			v3 = uint32(v2 | 4)
		} else {
			v3 = uint32(v2) & 0xFFFFFFFB
		}
		*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_830192 + 36))) = v3
		v4 = sub_41FF90()
		v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_830196 + 36))))
		if v4 != 0 {
			v6 = uint32(v5 | 4)
		} else {
			v6 = uint32(v5) & 0xFFFFFFFB
		}
		*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_830196 + 36))) = v6
		v7 = sub_41FFC0()
		v8 = int32(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_830200 + 36))))
		if v7 != 0 {
			v9 = uint32(v8 | 4)
		} else {
			v9 = uint32(v8) & 0xFFFFFFFB
		}
		*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_830200 + 36))) = v9
		dword_5d4594_830120 = 1
		sub_46B120((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_830104)))), nil)
		nox_xxx_wndShowModalMB((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_830104))))))
		sub_46C690((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_830104))))))
		guiFocus((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_830104))))))
		result = 1
	}
	return result
}
func sub_4493D0(a1 int32, a2 int32, a3 *int32, a4 int32) int32 {
	var (
		v3     int32
		result int32
		v5     int32
	)
	if a2 != 0x4007 {
		return 0
	}
	v3 = (*nox_window)(unsafe.Pointer(a3)).ID()
	clientPlaySoundSpecial(766, 100)
	switch v3 {
	case 1985:
		v5 = sub_41FC40()
		sub_40DB50(v5+1, int32(uintptr(unsafe.Pointer(&a2))))
		sub_40D670(a2)
		return 0
	case 1994:
		sub_4494A0()
		result = 0
	case 1995:
		nox_xxx_wnd_46C6E0((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_830104))))))
		(*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_830104)))).Destroy()
		sub_447600()
		dword_5d4594_830120 = 0
		result = 0
	case 1996:
		sub_4B5770_wol_locale(*(*int32)(unsafe.Pointer(&dword_5d4594_830104)))
		result = 0
	default:
		return 0
	}
	return result
}
func sub_4494A0() {
	sub_41FF70(bool2int((int32(*(*uint8)(unsafe.Pointer(uintptr(dword_5d4594_830192 + 36)))) & 4) == 4))
	sub_41FFA0(bool2int((int32(*(*uint8)(unsafe.Pointer(uintptr(dword_5d4594_830196 + 36)))) & 4) == 4))
	sub_41FFD0(bool2int((int32(*(*uint8)(unsafe.Pointer(uintptr(dword_5d4594_830200 + 36)))) & 4) == 4))
	nox_xxx_wnd_46C6E0((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_830104))))))
	(*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_830104)))).Destroy()
	dword_5d4594_830120 = 0
	sub_447600()
}
func nox_xxx_guiServerListLoad_449530() int32 {
	var (
		v0     *byte
		v1     int32
		result int32
		v3     *uint32
		v4     int32
		i      *byte
		v6     *byte
		v7     *libc.WChar
		v8     int32
		v9     *uint32
		v10    *uint32
		v11    *uint32
		v12    *uint32
		v13    *uint32
		v14    *byte
		v15    *uint32
		v16    *byte
		v17    int32
		v18    *byte
		v19    int32
		v20    [66]libc.WChar
	)
	v0 = nil
	v16 = nil
	v17 = -1
	v1 = 0
	result = int32(uintptr(unsafe.Pointer(newWindowFromFile("servlist.wnd", unsafe.Pointer(cgoFuncAddr(sub_4497D0_wol_dialogs))))))
	dword_5d4594_830108 = uint32(result)
	if result != 0 {
		*(*uint32)(unsafe.Pointer(uintptr(result + 56))) = nox_color_black_2650656
		sub_46B120((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_830108)))), nil)
		nox_xxx_wndShowModalMB((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_830108))))))
		sub_46C690((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_830108))))))
		guiFocus((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_830108))))))
		if noxflags.HasGame(noxflags.GameFlag26) {
			v16 = sub_4A7F10()
			v0 = v16
		}
		v3 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_830108)))).ChildByID(1961)))
		v4 = sub_4200F0()
		for i = sub_4205B0(v4); i != nil; i = (*byte)(unsafe.Pointer(uintptr(sub_4206B0(int32(uintptr(unsafe.Pointer(i))))))) {
			if noxflags.HasGame(noxflags.GameFlag26) {
				v6 = libc.StrChr((*byte)(unsafe.Add(unsafe.Pointer(i), 24)), 58)
				if v6 != nil {
					if v0 != nil && libc.StrCaseCmp(v0, (*byte)(unsafe.Add(unsafe.Pointer(v6), 1))) == 0 {
						v17 = v1
					}
				}
			}
			nox_swprintf(&v20[0], libc.CWString("%S"), (*byte)(unsafe.Add(unsafe.Pointer(i), 24)))
			v7 = nox_wcschr(&v20[0], 58)
			if v7 != nil {
				(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v3)))))).Func94(asWindowEvent(0x400D, int32(uintptr(unsafe.Pointer((*libc.WChar)(unsafe.Add(unsafe.Pointer(v7), unsafe.Sizeof(libc.WChar(0))*1))))), -1))
			}
			v1++
		}
		v8 = sub_4200F0()
		if sub_4207A0(v8) <= 0 {
			v9 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_830108)))).ChildByID(1962)))
			nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v9)))))), 0)
		} else {
			(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v3)))))).Func94(asWindowEvent(0x4013, 0, 0))
		}
		v10 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_830108)))).ChildByID(1964)))
		v11 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_830108)))).ChildByID(1965)))
		v12 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_830108)))).ChildByID(1966)))
		v13 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*8)))))
		v19 = int32(uintptr(unsafe.Pointer(v12)))
		v18 = (*byte)(unsafe.Pointer(nox_xxx_gLoadImg("UISlider")))
		v14 = (*byte)(unsafe.Pointer(nox_xxx_gLoadImg("UISliderLit")))
		sub_4B5700(int32(uintptr(unsafe.Pointer(v10))), 0, 0, int32(uintptr(unsafe.Pointer(v18))), int32(uintptr(unsafe.Pointer(v14))), int32(uintptr(unsafe.Pointer(v14))))
		nox_xxx_wnd_46B280(int32(uintptr(unsafe.Pointer(v10))), int32(uintptr(unsafe.Pointer(v3))))
		nox_xxx_wnd_46B280(int32(uintptr(unsafe.Pointer(v11))), int32(uintptr(unsafe.Pointer(v3))))
		nox_xxx_wnd_46B280(v19, int32(uintptr(unsafe.Pointer(v3))))
		*(*uint32)(unsafe.Add(unsafe.Pointer(v13), unsafe.Sizeof(uint32(0))*9)) = uint32(uintptr(unsafe.Pointer(v10)))
		*(*uint32)(unsafe.Add(unsafe.Pointer(v13), unsafe.Sizeof(uint32(0))*7)) = uint32(uintptr(unsafe.Pointer(v11)))
		*(*uint32)(unsafe.Add(unsafe.Pointer(v13), unsafe.Sizeof(uint32(0))*8)) = uint32(v19)
		*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v10), unsafe.Sizeof(uint32(0))*100)) + 8))) = 16
		*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v10), unsafe.Sizeof(uint32(0))*100)) + 12))) = 10
		if noxflags.HasGame(noxflags.GameFlag26) && v16 != nil && v17 != -1 {
			*(*uint32)(unsafe.Pointer(uintptr(int32(uintptr((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_830108)))).ChildByID(1961).widget_data)) + 48))) = uint32(v17)
			v15 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_830108)))).ChildByID(1962)))
			(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_830108))))).Func94(asWindowEvent(0x4007, int32(uintptr(unsafe.Pointer(v15))), 0))
		}
		result = 1
	}
	return result
}
func sub_44A520(a1 int32) *byte {
	var (
		v1 int32
		v2 int32
		v3 *uint8
	)
	v1 = int32(*memmap.PtrUint32(0x587000, 113136))
	v2 = 0
	if *memmap.PtrUint32(0x587000, 113136) == 0 {
		return *(**byte)(memmap.PtrOff(0x587000, 113140))
	}
	v3 = (*uint8)(memmap.PtrOff(0x587000, 113136))
	for v1 != a1 {
		v1 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v3))), unsafe.Sizeof(uint32(0))*2))))
		v3 = (*uint8)(unsafe.Add(unsafe.Pointer(v3), 8))
		v2++
		if v1 == 0 {
			return *(**byte)(memmap.PtrOff(0x587000, 113140))
		}
	}
	return *(**byte)(memmap.PtrOff(0x587000, uint32(v2*8)+113140))
}
func sub_44AA40() int32 {
	nox_wnd_xxx_830244.state = nox_gui_anim_state(NOX_GUI_ANIM_OUT)
	sub_43BE40(2)
	clientPlaySoundSpecial(923, 100)
	sub_4207F0(1)
	return 1
}
func sub_44AA70() int32 {
	var v0 func() int32
	v0 = nox_wnd_xxx_830244.field_13
	nox_gui_freeAnimation_43C570(nox_wnd_xxx_830244)
	if dword_5d4594_830248 != 0 {
		(*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_830248)))).Destroy()
		dword_5d4594_830248 = 0
	}
	sub_43BE40(1)
	v0()
	return 1
}
func sub_44AAC0() int32 {
	return 0
}
func sub_44AAD0(a1 *uint32, a2 int32, a3 int32, a4 int32) int32 {
	if a2 == 21 {
		if dword_5d4594_830268 != 0 {
			(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1)))))).Func94(asWindowEvent(0x401E, int32(uintptr(memmap.PtrOff(6112660, 830284))), 0))
			sub_41FAE0((*byte)(memmap.PtrOff(6112660, 830288)))
		}
		dword_5d4594_830268 = 0
	}
	return nox_xxx_wndEditProc_487D70((*nox_window)(unsafe.Pointer(a1)), a2, a3, a4)
}
func sub_44AB30(a1 int32, a2 uint32, a3 *int32, a4 int32) int32 {
	var (
		result int32
		v5     int32
		v6     int32
		v7     int32
		v8     int32
		v9     int32
		v10    int32
		v11    int32
		v12    int32
		v13    int32
		v14    *byte
		v15    int32
		v16    int32
		v17    *uint8
		v18    int32
		File   [128]libc.WChar
		v20    [256]byte
	)
	if a2 > 0x4005 {
		if a2 != 0x4007 {
			if a2 == 16400 {
				v6 = (*nox_window)(unsafe.Pointer(a3)).ID()
				clientPlaySoundSpecial(766, 100)
				if v6 == 1708 && a4 >= 0 && a4 < 128 {
					sub_41FC20(a4)
					sub_41FB90(a4, (*uint32)(unsafe.Pointer(&v18)), (*uint32)(unsafe.Pointer(&v17)))
					if v18 != 0 {
						nox_swprintf(&File[0], libc.CWString("%S"), v18)
						(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_830256))))).Func94(asWindowEvent(0x401E, int32(uintptr(unsafe.Pointer(&File[0]))), 0))
					}
					if v17 != nil {
						nox_swprintf(&File[0], libc.CWString("%S"), v17)
						(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_830260))))).Func94(asWindowEvent(0x401E, int32(uintptr(unsafe.Pointer(&File[0]))), 0))
						if v17 != nil {
							if int32(*v17) != 0 {
								*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_830252 + 36))) |= 4
								dword_5d4594_830276 = 1
								return 0
							}
						}
					}
					*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_830252 + 36))) &= 0xFFFFFFFB
					dword_5d4594_830276 = 0
				}
			}
			return 0
		}
		v7 = (*nox_window)(unsafe.Pointer(a3)).ID()
		if v7 > 1706 {
			if v7 != 1985 {
				if v7 == 4010 {
					v16 = sub_4200F0()
					v14 = sub_44A520(*memmap.PtrInt32(6112660, 830280))
					nox_sprintf((*byte)(unsafe.Pointer(&File[0])), CString("%s%d.html"), v14, v16)
					compatShellExecuteA(0, nil, (*byte)(unsafe.Pointer(&File[0])), nil, (*byte)(memmap.PtrOff(0x587000, 113732)), 3)
				}
				goto LABEL_42
			}
		} else {
			if v7 != 1706 {
				v8 = v7 - 1703
				if v8 != 0 {
					if v8 == 2 {
						sub_41E300(6)
						sub_4207F0(6)
						nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_830248))))), 0)
						sub_40E090()
						sub_4A1BE0(0)
					}
				} else {
					dword_5d4594_830276 ^= 1
					if dword_5d4594_830276 != 0 {
						v10 = sub_41F9A0()
						(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_830264))))).Func94(asWindowEvent(0x4013, v10, 0))
						sub_41FC20(v10)
					} else {
						v9 = sub_41FB60()
						sub_41FC20(128)
						sub_41FB70(v9)
					}
				}
				goto LABEL_42
			}
			v17 = nil
			if dword_5d4594_830276 == 0 {
				sub_41FC20(128)
			}
			v11 = (*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_830256))))).Func94(asWindowEvent(0x401D, 0, 0))
			nox_sprintf(&v20[0], CString("%S"), v11)
			sub_41FA80(&v20[0])
			v12 = (*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_830260))))).Func94(asWindowEvent(0x401D, 0, 0))
			nox_sprintf(&v20[0], CString("%S"), v12)
			sub_41FAE0(&v20[0])
			sub_41FC50()
			v13 = sub_41FC40()
			sub_40DB50(v13+1, int32(uintptr(unsafe.Pointer(&v17))))
			if v17 == nil {
				sub_4B5770_wol_locale(*(*int32)(unsafe.Pointer(&dword_5d4594_830248)))
			LABEL_42:
				clientPlaySoundSpecial(921, 100)
				return 1
			}
		}
		dword_5d4594_830276 = uint32(int32(uint8(int8(int32(*(*uint8)(unsafe.Pointer(uintptr(dword_5d4594_830252 + 36))))&4))) >> 2)
		sub_41E300(2)
		v15 = sub_41E2F0()
		sub_41DA70(v15, 0)
		sub_4207F0(2)
		sub_40E090()
		nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_830248))))), 0)
		sub_4A1BE0(0)
		goto LABEL_42
	}
	if a2 == 0x4005 {
		clientPlaySoundSpecial(920, 100)
		return 1
	}
	if a2 == 2 {
		dword_5d4594_830248 = 0
		return 0
	}
	if a2 != 0x4003 {
		return 0
	}
	result = int32(uintptr(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_830248)))).ChildByID(a4))))
	if result == 0 {
		return result
	}
	if int32(uint16(uintptr(unsafe.Pointer(a3)))) != 1 {
		return 0
	}
	if a4 == 1702 {
		v5 = (*nox_window)(unsafe.Pointer(uintptr(result))).Func94(asWindowEvent(0x401D, 0, 0))
		nox_sprintf(&v20[0], CString("%S"), v5)
		if v20[0] != 0 {
			dword_5d4594_830268 = 1
		}
	}
	return 0
}
func sub_44AF70() int32 {
	var result int32
	nox_wnd_xxx_830244.state = nox_gui_anim_state(NOX_GUI_ANIM_OUT)
	sub_43BE40(2)
	clientPlaySoundSpecial(923, 100)
	result = 1
	nox_wnd_xxx_830244.field_13 = nox_game_showWolChat_447620
	dword_5d4594_830248 = 0
	return result
}
func sub_44AFB0(a1 int32, a2 int32, a3 int32) *uint32 {
	var result *uint32
	if a3 < -2147221388 || a3 > -2147221386 {
		result = (*uint32)(NewDialogWindow((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_830248))))), (*libc.WChar)(unsafe.Pointer(uintptr(a1))), (*libc.WChar)(unsafe.Pointer(uintptr(a2))), 33, nil, nil))
	} else {
		result = (*uint32)(NewDialogWindow((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_830248))))), (*libc.WChar)(unsafe.Pointer(uintptr(a1))), (*libc.WChar)(unsafe.Pointer(uintptr(a2))), 545, nil, nil))
	}
	*memmap.PtrUint32(6112660, 830280) = uint32(a3)
	return result
}
func nox_xxx____setargv_4_44B000() {
	dword_5d4594_830272 = 1
}
func sub_44B0F0(a1 int32, a2 int32, a3 *int32, a4 int32) int32 {
	var v3 int32
	if a2 != 23 {
		if a2 == 0x4007 {
			v3 = (*nox_window)(unsafe.Pointer(a3)).ID() - 581
			if v3 == 0 {
				nox_xxx____setargv_4_44B000()
				(*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_830292)))).Destroy()
				goto LABEL_7
			}
			if v3 == 1 {
				nox_client_quit_4460C0()
				sub_44A400()
				(*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_830292)))).Destroy()
			LABEL_7:
				dword_5d4594_830292 = 0
				return 0
			}
		}
		return 0
	}
	return 1
}
func sub_44B940(a1 *uint32, a2 int32, f *nox_memfile) int32 {
	var (
		v3  int32
		v4  int32
		v5  unsafe.Pointer
		v6  int32
		v8  int32
		v10 int8
		v13 int32
		v14 int32
		v15 uint8
		v16 *byte
		v17 [128]byte
	)
	v3 = a2
	v4 = 0
	*a1 = 40
	v14 = 0
	for {
		if v4 >= 16 {
			v13 = v4 + 4
		} else {
			v13 = v4
		}
		v5 = alloc.Calloc(int(v3), 4)
		*(*uint32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(a1))), v13))), 4)))) = uint32(uintptr(v5))
		if v5 == nil {
			break
		}
		v6 = 0
		if v3 > 0 {
			for {
				v8 = f.ReadI32()
				v17[0] = byte(*memmap.PtrUint8(6112660, 830840))
				if v8 == -1 {
					v10 = f.ReadI8()
					v15 = f.ReadU8()
					*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v16))), 0)) = uint8(v10)
					nox_memfile_read(unsafe.Pointer(&v17[0]), 1, int32(v15), f)
					v17[v15] = 0
					v3 = a2
				}
				*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(a1))), v13))), 4)))) + uint32(func() int32 {
					p := &v6
					*p++
					return *p
				}()*4) - 4))) = uint32(uintptr(unsafe.Pointer(nox_xxx_readImgMB_42FAA0(v8, int8(uintptr(unsafe.Pointer(v16))), &v17[0]))))
				if v6 >= v3 {
					break
				}
			}
			v4 = v14
		}
		v4 += 4
		v14 = v4
		if v4 >= 32 {
			return 1
		}
	}
	return 0
}
func nox_xxx_parse_Armor_44BA60(a1 *byte) int32 {
	var (
		v1 int32
		v2 **byte
	)
	v1 = 0
	v2 = (**byte)(memmap.PtrOff(0x587000, 113856))
	for libc.StrCmp(*v2, a1) != 0 {
		v2 = (**byte)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof((*byte)(nil))*1))
		v1++
		if int32(uintptr(unsafe.Pointer(v2))) >= int32(uintptr(memmap.PtrOff(0x587000, 113960))) {
			return -1
		}
	}
	return v1
}
func sub_44BAC0(a1 *byte) int32 {
	var (
		v1 int32
		v2 **byte
	)
	v1 = 0
	v2 = (**byte)(memmap.PtrOff(0x587000, 113964))
	for libc.StrCmp(*v2, a1) != 0 {
		v2 = (**byte)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof((*byte)(nil))*1))
		v1++
		if int32(uintptr(unsafe.Pointer(v2))) >= int32(uintptr(memmap.PtrOff(0x587000, 114072))) {
			return -1
		}
	}
	return v1
}
func sub_44BB20(a1 *byte) int32 {
	var (
		v1 int32
		v2 **byte
	)
	v1 = 0
	v2 = (**byte)(memmap.PtrOff(0x587000, 115688))
	for libc.StrCmp(a1, *v2) != 0 {
		v2 = (**byte)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof((*byte)(nil))*1))
		v1++
		if int32(uintptr(unsafe.Pointer(v2))) >= int32(uintptr(memmap.PtrOff(0x587000, 115908))) {
			return -1
		}
	}
	return v1
}
