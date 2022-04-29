package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	"github.com/gotranspile/cxgo/runtime/stdio"
	noxflags "github.com/noxworld-dev/opennox/v1/common/flags"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"math"
	"unsafe"
)

var g_scaled_cfg int32 = 0
var g_fullscreen_cfg int32 = 0

func nox_common_parsecfg_all(a1 *File) int32 {
	var (
		v1 *byte
		v2 *byte
		v3 *byte
		v4 int32
		v5 *uint8
	)
	sub_486670(0x4000, 0)
	sub_486670(0x4000, 1)
	sub_486670(0x4000, 2)
LABEL_2:
	for nox_fs_fgets(a1, (*byte)(memmap.PtrOff(6112660, 806084)), 1024) {
		if int32(*memmap.PtrUint8(6112660, 806084)) == int32('#') {
			continue
		}
		v1 = libc.StrTok((*byte)(memmap.PtrOff(6112660, 806084)), CString(" \r\t\n"))
		if v1 == nil {
			continue
		}
		if libc.StrCmp(v1, CString("---")) == 0 {
			return 1
		}
		if libc.StrCmp(v1, CString("Version")) == 0 {
			v2 = (*byte)(unsafe.Pointer(uintptr(sub_432AD0(memmap.PtrInt32(0x587000, 81284)))))
			if v2 == nil {
				return 0
			}
		} else if libc.StrCmp(v1, CString("VideoMode")) == 0 {
			v2 = (*byte)(unsafe.Pointer(uintptr(nox_common_parsecfg_videomode())))
			if v2 == nil {
				return 0
			}
		} else if libc.StrCmp(v1, CString("VideoSize")) == 0 {
			v2 = (*byte)(unsafe.Pointer(uintptr(nox_common_parsecfg_videosize())))
			if v2 == nil {
				return 0
			}
		} else if libc.StrCmp(v1, CString("Gamma")) == 0 {
			v2 = (*byte)(unsafe.Pointer(uintptr(sub_432C00())))
			if v2 == nil {
				return 0
			}
		} else if libc.StrCmp(v1, CString("FXVolume")) == 0 {
			v2 = (*byte)(unsafe.Pointer(uintptr(sub_432CB0(0))))
			if v2 == nil {
				return 0
			}
		} else if libc.StrCmp(v1, CString("DialogVolume")) == 0 {
			v2 = (*byte)(unsafe.Pointer(uintptr(sub_432CB0(1))))
			if v2 == nil {
				return 0
			}
		} else if libc.StrCmp(v1, CString("MusicVolume")) == 0 {
			v2 = (*byte)(unsafe.Pointer(uintptr(sub_432CB0(2))))
			if v2 == nil {
				return 0
			}
		} else if libc.StrCmp(v1, CString("LastServer")) == 0 {
			v2 = (*byte)(unsafe.Pointer(uintptr(sub_431FC0())))
			if v2 == nil {
				return 0
			}
		} else if libc.StrCmp(v1, CString("ServerName")) == 0 {
			v2 = (*byte)(unsafe.Pointer(uintptr(sub_432010())))
			if v2 == nil {
				return 0
			}
		} else if libc.StrCmp(v1, CString("UnlockSurface")) == 0 {
			v2 = (*byte)(unsafe.Pointer(uintptr(sub_4320B0())))
			if v2 == nil {
				return 0
			}
		} else if libc.StrCmp(v1, CString("SoftShadowEdge")) == 0 {
			v2 = (*byte)(unsafe.Pointer(uintptr(sub_432100())))
			if v2 == nil {
				return 0
			}
		} else if libc.StrCmp(v1, CString("DrawFrontWalls")) == 0 {
			v2 = (*byte)(unsafe.Pointer(uintptr(sub_432150())))
			if v2 == nil {
				return 0
			}
		} else if libc.StrCmp(v1, CString("TranslucentFrontWalls")) == 0 {
			v2 = (*byte)(unsafe.Pointer(uintptr(sub_4321A0())))
			if v2 == nil {
				return 0
			}
		} else if libc.StrCmp(v1, CString("HighResFrontWalls")) == 0 {
			v2 = (*byte)(unsafe.Pointer(uintptr(sub_4321F0())))
			if v2 == nil {
				return 0
			}
		} else if libc.StrCmp(v1, CString("HighResFloors")) == 0 {
			v2 = (*byte)(unsafe.Pointer(uintptr(sub_432240())))
			if v2 == nil {
				return 0
			}
		} else if libc.StrCmp(v1, CString("LockHighResFloors")) == 0 {
			v2 = (*byte)(unsafe.Pointer(uintptr(sub_432290())))
			if v2 == nil {
				return 0
			}
		} else if libc.StrCmp(v1, CString("TexturedFloors")) == 0 {
			v2 = (*byte)(unsafe.Pointer(uintptr(sub_4322E0())))
			if v2 == nil {
				return 0
			}
		} else if libc.StrCmp(v1, CString("TranslucentConsole")) == 0 {
			v2 = (*byte)(unsafe.Pointer(uintptr(sub_432340())))
			if v2 == nil {
				return 0
			}
		} else if libc.StrCmp(v1, CString("RenderGlow")) == 0 {
			v2 = (*byte)(unsafe.Pointer(uintptr(sub_432390())))
			if v2 == nil {
				return 0
			}
		} else if libc.StrCmp(v1, CString("RenderGUI")) == 0 {
			v2 = (*byte)(unsafe.Pointer(uintptr(sub_4323E0())))
			if v2 == nil {
				return 0
			}
		} else if libc.StrCmp(v1, CString("FadeObjects")) == 0 {
			v2 = (*byte)(unsafe.Pointer(uintptr(sub_432430())))
			if v2 == nil {
				return 0
			}
		} else if libc.StrCmp(v1, CString("TrackData")) == 0 {
			v2 = (*byte)(unsafe.Pointer(uintptr(sub_433190())))
			if v2 == nil {
				return 0
			}
		} else if libc.StrCmp(v1, CString("RenderBubbles")) == 0 {
			v2 = (*byte)(unsafe.Pointer(uintptr(sub_432480())))
			if v2 == nil {
				return 0
			}
		} else if libc.StrCmp(v1, CString("SysopPassword")) == 0 {
			v2 = (*byte)(unsafe.Pointer(uintptr(sub_4324D0())))
			if v2 == nil {
				return 0
			}
		} else if libc.StrCmp(v1, CString("ServerPassword")) == 0 {
			v2 = (*byte)(unsafe.Pointer(uintptr(sub_432520())))
			if v2 == nil {
				return 0
			}
		} else if libc.StrCmp(v1, CString("Profiled")) == 0 {
			v2 = (*byte)(unsafe.Pointer(uintptr(sub_432590())))
			if v2 == nil {
				return 0
			}
		} else if libc.StrCmp(v1, CString("SanctuaryHelp")) == 0 {
			v2 = (*byte)(unsafe.Pointer(uintptr(sub_432620())))
			if v2 == nil {
				return 0
			}
		} else if libc.StrCmp(v1, CString("MaxPacketLossPct")) == 0 {
			v2 = (*byte)(unsafe.Pointer(uintptr(sub_4325D0())))
			if v2 == nil {
				return 0
			}
		} else if libc.StrCmp(v1, CString("SendMessageOfTheDay")) == 0 {
			v2 = (*byte)(unsafe.Pointer(uintptr(sub_432A90())))
			if v2 == nil {
				return 0
			}
		} else if libc.StrCmp(v1, CString("InternetUpdateRate")) == 0 {
			v2 = (*byte)(unsafe.Pointer(uintptr(sub_432DF0())))
			if v2 == nil {
				return 0
			}
		} else if libc.StrCmp(v1, CString("ConnectionType")) == 0 {
			v2 = (*byte)(unsafe.Pointer(uintptr(sub_433050())))
			if v2 == nil {
				return 0
			}
		} else if libc.StrCmp(v1, CString("MapCycle")) == 0 {
			v2 = (*byte)(unsafe.Pointer(uintptr(sub_432C30())))
			if v2 == nil {
				return 0
			}
		} else if libc.StrCmp(v1, CString("LANFilters")) == 0 {
			v2 = sub_432E50(0)
			if v2 == nil {
				return 0
			}
		} else if libc.StrCmp(v1, CString("INETFilters")) == 0 {
			v2 = sub_432E50(1)
			if v2 == nil {
				return 0
			}
		} else if libc.StrCmp(v1, CString("LessonLimit")) == 0 {
			if sub_432D10() == 0 {
				return 0
			}
		} else if libc.StrCmp(v1, CString("TimeLimit")) == 0 {
			if sub_432D80() == 0 {
				return 0
			}
		} else if libc.StrCmp(v1, CString("PlayerSkeletons")) == 0 {
			if sub_4327C0() == 0 {
				return 0
			}
		} else if libc.StrCmp(v1, CString("BroadcastGestures")) == 0 {
			if sub_432810() == 0 {
				return 0
			}
		} else if libc.StrCmp(v1, CString("LatencyCompensation")) == 0 {
			if sub_432970() == 0 {
				return 0
			}
		} else if libc.StrCmp(v1, CString("Closed")) == 0 {
			if sub_432860() == 0 {
				return 0
			}
		} else if libc.StrCmp(v1, CString("Private")) == 0 {
			if sub_4328C0() == 0 {
				return 0
			}
		} else if libc.StrCmp(v1, CString("AudioThreshold")) == 0 {
			if sub_433130() == 0 {
				return 0
			}
		} else if libc.StrCmp(v1, CString("MaxPlayers")) == 0 {
			if sub_4330C0() == 0 {
				return 0
			}
		} else if libc.StrCmp(v1, CString("RestrictedClasses")) == 0 {
			if sub_432920() == 0 {
				return 0
			}
		} else if libc.StrCmp(v1, CString("RestrictedPing")) == 0 {
			if sub_4329D0() == 0 {
				return 0
			}
		} else if libc.StrCmp(v1, CString("LimitMaxRes")) == 0 {
			if sub_432A30() == 0 {
				return 0
			}
		} else if libc.StrCmp(v1, CString("ItemRespawn")) == 0 {
			if sub_432660() == 0 {
				return 0
			}
		} else if libc.StrCmp(v1, CString("CamperAlarm")) == 0 {
			if sub_4326B0() == 0 {
				return 0
			}
		} else if libc.StrCmp(v1, CString("MinKickVotes")) == 0 {
			if sub_432700() == 0 {
				return 0
			}
		} else if libc.StrCmp(v1, CString("ResetQuestMinVotes")) == 0 {
			if sub_432740() == 0 {
				return 0
			}
		} else if libc.StrCmp(v1, CString("KickQuestPlayerMinVotes")) == 0 {
			if sub_432780() == 0 {
				return 0
			}
		} else if libc.StrCmp(v1, CString("Fullscreen")) == 0 {
			var token *byte
			libc.StrTok(nil, CString(" \r\t\n"))
			token = libc.StrTok(nil, CString(" \r\t\n"))
			if token != nil {
				g_fullscreen_cfg = int32(libc.Atoi(libc.GoString(token)))
				if nox_video_getFullScreen() <= -4 {
					nox_video_setFullScreen(g_fullscreen_cfg)
				}
			}
			change_windowed_fullscreen()
		} else if libc.StrCmp(v1, CString("Gamma2")) == 0 {
			var token *byte
			libc.StrTok(nil, CString(" \r\t\n"))
			token = libc.StrTok(nil, CString(" \r\t\n"))
			if token != nil {
				nox_video_setGamma(float32(libc.Atof(libc.GoString(token))))
			}
		} else if libc.StrCmp(v1, CString("Stretched")) == 0 {
			var token *byte
			libc.StrTok(nil, CString(" \r\t\n"))
			token = libc.StrTok(nil, CString(" \r\t\n"))
			if token != nil {
				g_scaled_cfg = int32(libc.Atoi(libc.GoString(token)))
				if nox_video_getScaled() >= 0 {
					nox_video_setScaled(g_scaled_cfg)
				}
			}
		} else if libc.StrCmp(v1, CString("InputSensitivity")) == 0 {
			var token *byte
			libc.StrTok(nil, CString(" \r\t\n"))
			token = libc.StrTok(nil, CString(" \r\t\n"))
			if token != nil {
				nox_input_setSensitivity(float32(libc.Atof(libc.GoString(token))))
			}
		} else {
			v3 = *(**byte)(memmap.PtrOff(0x587000, 81168))
			v4 = 0
			if *memmap.PtrUint32(0x587000, 81168) != 0 {
				v5 = (*uint8)(memmap.PtrOff(0x587000, 81168))
				for libc.StrCmp(v1, v3) != 0 {
					v3 = (*byte)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v5))), unsafe.Sizeof(uint32(0))*2))))))
					v5 = (*uint8)(unsafe.Add(unsafe.Pointer(v5), 8))
					v4++
					if v3 == nil {
						goto LABEL_2
					}
				}
				v2 = (*byte)(unsafe.Pointer(uintptr(sub_432C70(int32(*memmap.PtrUint32(0x587000, uint32(v4*8)+81172))))))
				if v2 == nil {
					return 0
				}
			}
		}
	}
	return 0
}
func sub_431FC0() int32 {
	var v0 *byte
	libc.StrTok(nil, CString(" \r\t\n"))
	v0 = libc.StrTok(nil, CString(" \r\t\n"))
	if v0 != nil {
		libc.StrNCpy((*byte)(memmap.PtrOff(6112660, 806060)), v0, 22)
	} else {
		*memmap.PtrUint8(6112660, 806060) = *memmap.PtrUint8(6112660, 807272)
	}
	return 1
}
func sub_432010() int32 {
	var (
		v0     *byte
		v1     *byte
		result int32
		v3     *byte
		v4     int8
		v5     [64]byte
	)
	libc.StrTok(nil, CString(" \r\t\n"))
	v0 = libc.StrTok(nil, CString(" \r\t\n"))
	v1 = v0
	if v0 != nil {
		if *memmap.PtrUint32(0x587000, 81284) == 0x10000 && libc.StrCaseCmp(v0, CString("NoxWorld")) == 0 {
			v1 = CString("User_Game")
		}
		libc.StrNCpy(&v5[0], v1, 63)
		v3 = &v5[0]
		if v5[0] != 0 {
			for {
				if *v3 == 95 {
					*v3 = 32
				}
				v4 = int8(*func() *byte {
					p := &v3
					*p = (*byte)(unsafe.Add(unsafe.Pointer(*p), 1))
					return *p
				}())
				if int32(v4) == 0 {
					break
				}
			}
		}
		nox_xxx_gameSetServername_40A440(&v5[0])
		result = 1
	} else {
		nox_xxx_gameSetServername_40A440(nil)
		result = 1
	}
	return result
}
func sub_4320B0() int32 {
	var (
		v0     *byte
		v1     int32
		v2     bool
		result int32
	)
	libc.StrTok(nil, CString(" \r\t\n"))
	v0 = libc.StrTok(nil, CString(" \r\t\n"))
	if v0 == nil {
		return 1
	}
	v1 = int32(libc.Atoi(libc.GoString(v0)))
	nox_video_dxUnlockSurface = uint32(v1)
	v2 = v1 == 0
	result = 1
	if !v2 {
		nox_video_dxUnlockSurface = 1
	}
	return result
}
func sub_432100() int32 {
	var (
		v0 *byte
		v1 bool
	)
	libc.StrTok(nil, CString(" \r\t\n"))
	v0 = libc.StrTok(nil, CString(" \r\t\n"))
	if v0 != nil {
		v1 = libc.Atoi(libc.GoString(v0)) == 0
		if !v1 {
			noxflags.SetEngine(nox_engine_flag(NOX_ENGINE_FLAG_ENABLE_SOFT_SHADOW_EDGE))
			return 1
		}
		noxflags.UnsetEngine(nox_engine_flag(NOX_ENGINE_FLAG_ENABLE_SOFT_SHADOW_EDGE))
	}
	return 1
}
func sub_432150() int32 {
	var (
		v0     *byte
		v1     int32
		v2     bool
		result int32
	)
	libc.StrTok(nil, CString(" \r\t\n"))
	v0 = libc.StrTok(nil, CString(" \r\t\n"))
	if v0 == nil {
		return 1
	}
	v1 = int32(libc.Atoi(libc.GoString(v0)))
	nox_client_drawFrontWalls_80812 = uint32(v1)
	v2 = v1 == 0
	result = 1
	if !v2 {
		nox_client_drawFrontWalls_80812 = 1
	}
	return result
}
func sub_4321A0() int32 {
	var (
		v0     *byte
		v1     int32
		v2     bool
		result int32
	)
	libc.StrTok(nil, CString(" \r\t\n"))
	v0 = libc.StrTok(nil, CString(" \r\t\n"))
	if v0 == nil {
		return 1
	}
	v1 = int32(libc.Atoi(libc.GoString(v0)))
	nox_client_translucentFrontWalls_805844 = uint32(v1)
	v2 = v1 == 0
	result = 1
	if !v2 {
		nox_client_translucentFrontWalls_805844 = 1
	}
	return result
}
func sub_4321F0() int32 {
	var (
		v0     *byte
		v1     int32
		v2     bool
		result int32
	)
	libc.StrTok(nil, CString(" \r\t\n"))
	v0 = libc.StrTok(nil, CString(" \r\t\n"))
	if v0 == nil {
		return 1
	}
	v1 = int32(libc.Atoi(libc.GoString(v0)))
	nox_client_highResFrontWalls_80820 = uint32(v1)
	v2 = v1 == 0
	result = 1
	if !v2 {
		nox_client_highResFrontWalls_80820 = 1
	}
	return result
}
func sub_432240() int32 {
	var (
		v0     *byte
		v1     int32
		v2     bool
		result int32
	)
	libc.StrTok(nil, CString(" \r\t\n"))
	v0 = libc.StrTok(nil, CString(" \r\t\n"))
	if v0 == nil {
		return 1
	}
	v1 = int32(libc.Atoi(libc.GoString(v0)))
	nox_client_highResFloors_154952 = uint32(v1)
	v2 = v1 == 0
	result = 1
	if !v2 {
		nox_client_highResFloors_154952 = 1
	}
	return result
}
func sub_432290() int32 {
	var (
		v0     *byte
		v1     int32
		v2     bool
		result int32
	)
	libc.StrTok(nil, CString(" \r\t\n"))
	v0 = libc.StrTok(nil, CString(" \r\t\n"))
	if v0 == nil {
		return 1
	}
	v1 = int32(libc.Atoi(libc.GoString(v0)))
	nox_client_lockHighResFloors_1193152 = uint32(v1)
	v2 = v1 == 0
	result = 1
	if !v2 {
		nox_client_lockHighResFloors_1193152 = 1
	}
	return result
}
func sub_4322E0() int32 {
	var v0 *byte
	libc.StrTok(nil, CString(" \r\t\n"))
	v0 = libc.StrTok(nil, CString(" \r\t\n"))
	if v0 != nil {
		nox_client_texturedFloors_154956 = uint32(libc.Atoi(libc.GoString(v0)))
		nox_client_texturedFloors_154956 = uint32(bool2int(nox_client_texturedFloors_154956 != 0))
		nox_xxx_tileSetDrawFn_481420()
		dword_5d4594_1193156 = 0
		nox_client_texturedFloors2_154960 = nox_client_texturedFloors_154956
	}
	return 1
}
func sub_432340() int32 {
	var (
		v0     *byte
		v1     int32
		v2     bool
		result int32
	)
	libc.StrTok(nil, CString(" \r\t\n"))
	v0 = libc.StrTok(nil, CString(" \r\t\n"))
	if v0 == nil {
		return 1
	}
	v1 = int32(libc.Atoi(libc.GoString(v0)))
	nox_gui_console_translucent_set(v1)
	v2 = v1 == 0
	result = 1
	if !v2 {
		nox_gui_console_translucent_set(1)
	}
	return result
}
func sub_432390() int32 {
	var (
		v0     *byte
		v1     int32
		v2     bool
		result int32
	)
	libc.StrTok(nil, CString(" \r\t\n"))
	v0 = libc.StrTok(nil, CString(" \r\t\n"))
	if v0 == nil {
		return 1
	}
	v1 = int32(libc.Atoi(libc.GoString(v0)))
	nox_client_renderGlow_805852 = uint32(v1)
	v2 = v1 == 0
	result = 1
	if !v2 {
		nox_client_renderGlow_805852 = 1
	}
	return result
}
func sub_4323E0() int32 {
	var (
		v0     *byte
		v1     int32
		v2     bool
		result int32
	)
	libc.StrTok(nil, CString(" \r\t\n"))
	v0 = libc.StrTok(nil, CString(" \r\t\n"))
	if v0 == nil {
		return 1
	}
	v1 = int32(libc.Atoi(libc.GoString(v0)))
	nox_client_renderGUI_80828 = uint32(v1)
	v2 = v1 == 0
	result = 1
	if !v2 {
		nox_client_renderGUI_80828 = 1
	}
	return result
}
func sub_432430() int32 {
	var (
		v0     *byte
		v1     int32
		v2     bool
		result int32
	)
	libc.StrTok(nil, CString(" \r\t\n"))
	v0 = libc.StrTok(nil, CString(" \r\t\n"))
	if v0 == nil {
		return 1
	}
	v1 = int32(libc.Atoi(libc.GoString(v0)))
	nox_client_fadeObjects_80836 = uint32(v1)
	v2 = v1 == 0
	result = 1
	if !v2 {
		nox_client_fadeObjects_80836 = 1
	}
	return result
}
func sub_432480() int32 {
	var (
		v0     *byte
		v1     int32
		v2     bool
		result int32
	)
	libc.StrTok(nil, CString(" \r\t\n"))
	v0 = libc.StrTok(nil, CString(" \r\t\n"))
	if v0 == nil {
		return 1
	}
	v1 = int32(libc.Atoi(libc.GoString(v0)))
	nox_client_renderBubbles_80844 = uint32(v1)
	v2 = v1 == 0
	result = 1
	if !v2 {
		nox_client_renderBubbles_80844 = 1
	}
	return result
}
func sub_4324D0() int32 {
	var (
		v0 *byte
		v2 [100]libc.WChar
	)
	libc.StrTok(nil, CString(" \r\t\n"))
	v0 = libc.StrTok(nil, CString("\r\t\n"))
	if v0 != nil {
		nox_swprintf(&v2[0], libc.CWString("%S"), v0)
		nox_xxx_sysopSetPass_40A610(&v2[0])
	}
	return 1
}
func sub_432520() int32 {
	var (
		v0 *byte
		v1 *byte
		v3 [100]libc.WChar
	)
	v0 = (*byte)(sub_416640())
	libc.StrTok(nil, CString(" \r\t\n"))
	v1 = libc.StrTok(nil, CString("\r\t\n"))
	if v1 != nil {
		nox_swprintf(&v3[0], libc.CWString("%S"), v1)
		nox_wcsncpy((*libc.WChar)(unsafe.Add(unsafe.Pointer((*libc.WChar)(unsafe.Pointer(v0))), unsafe.Sizeof(libc.WChar(0))*39)), &v3[0], 8)
		*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v0))), unsafe.Sizeof(uint16(0))*47))) = 0
	}
	return 1
}
func sub_432590() int32 {
	var v0 *byte
	libc.StrTok(nil, CString(" \r\t\n"))
	v0 = libc.StrTok(nil, CString(" \r\t\n"))
	if v0 != nil {
		nox_profiled_805856 = uint32(bool2int(libc.Atoi(libc.GoString(v0)) != 0))
	}
	return 1
}
func sub_4325D0() int32 {
	var v0 *byte
	libc.StrTok(nil, CString(" \r\t\n"))
	v0 = libc.StrTok(nil, CString(" \r\t\n"))
	if v0 != nil {
		*memmap.PtrUint32(0x587000, 81280) = uint32(libc.Atoi(libc.GoString(v0)))
		*memmap.PtrUint32(0x587000, 292940) = uint32(int32(int64(float64(*memmap.PtrInt32(0x587000, 81280)) * 0.0099999998 * 10.0)))
	}
	return 1
}
func sub_432620() int32 {
	var v0 *byte
	libc.StrTok(nil, CString(" \r\t\n"))
	v0 = libc.StrTok(nil, CString(" \r\t\n"))
	if v0 != nil {
		nox_server_sanctuaryHelp_54276 = uint32(bool2int(libc.Atoi(libc.GoString(v0)) != 0))
	}
	return 1
}
func sub_432660() int32 {
	var v0 *byte
	libc.StrTok(nil, CString(" \r\t\n"))
	v0 = libc.StrTok(nil, CString(" \r\t\n"))
	if v0 != nil {
		if libc.Atoi(libc.GoString(v0)) != 0 {
			sub_409E70(2)
			return 1
		}
		sub_409EC0(2)
	}
	return 1
}
func sub_4326B0() int32 {
	var v0 *byte
	libc.StrTok(nil, CString(" \r\t\n"))
	v0 = libc.StrTok(nil, CString(" \r\t\n"))
	if v0 != nil {
		if libc.Atoi(libc.GoString(v0)) != 0 {
			sub_409E70(8192)
			return 1
		}
		sub_409EC0(8192)
	}
	return 1
}
func sub_432700() int32 {
	var v0 *byte
	libc.StrTok(nil, CString(" \r\t\n"))
	v0 = libc.StrTok(nil, CString(" \r\t\n"))
	if v0 != nil {
		*memmap.PtrUint32(0x587000, 229980) = uint32(libc.Atoi(libc.GoString(v0)))
	}
	return 1
}
func sub_432740() int32 {
	var v0 *byte
	libc.StrTok(nil, CString(" \r\t\n"))
	v0 = libc.StrTok(nil, CString(" \r\t\n"))
	if v0 != nil {
		nox_server_resetQuestMinVotes_229988 = uint32(libc.Atoi(libc.GoString(v0)))
	}
	return 1
}
func sub_432780() int32 {
	var v0 *byte
	libc.StrTok(nil, CString(" \r\t\n"))
	v0 = libc.StrTok(nil, CString(" \r\t\n"))
	if v0 != nil {
		nox_server_kickQuestPlayerMinVotes_229992 = uint32(libc.Atoi(libc.GoString(v0)))
	}
	return 1
}
func sub_4327C0() int32 {
	var (
		v0 *byte
		v1 *byte
	)
	_ = v1
	libc.StrTok(nil, CString(" \r\t\n"))
	v0 = libc.StrTok(nil, CString(" \r\t\n"))
	if v0 != nil {
		v1 = (*byte)(sub_416640())
		*(*uint32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v1), 58)))) = uint32(bool2int(libc.Atoi(libc.GoString(v0)) > 0))
	}
	return 1
}
func sub_432810() int32 {
	var (
		v0 *byte
		v1 *byte
	)
	_ = v1
	libc.StrTok(nil, CString(" \r\t\n"))
	v0 = libc.StrTok(nil, CString(" \r\t\n"))
	if v0 != nil {
		v1 = (*byte)(sub_416640())
		*(*uint32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v1), 62)))) = uint32(bool2int(libc.Atoi(libc.GoString(v0)) > 0))
	}
	return 1
}
func sub_432860() int32 {
	var (
		v0 *byte
		v1 *byte
		v2 bool
		v3 int8
	)
	libc.StrTok(nil, CString(" \r\t\n"))
	v0 = libc.StrTok(nil, CString(" \r\t\n"))
	if v0 != nil {
		v1 = (*byte)(sub_416640())
		v2 = libc.Atoi(libc.GoString(v0)) <= 0
		v3 = int8(*(*byte)(unsafe.Add(unsafe.Pointer(v1), 100)))
		if !v2 {
			*(*byte)(unsafe.Add(unsafe.Pointer(v1), 100)) = byte(int8(int32(v3) | 16))
			return 1
		}
		*(*byte)(unsafe.Add(unsafe.Pointer(v1), 100)) = byte(int8(int32(v3) & 239))
	}
	return 1
}
func sub_4328C0() int32 {
	var (
		v0 *byte
		v1 *byte
		v2 bool
		v3 int8
	)
	libc.StrTok(nil, CString(" \r\t\n"))
	v0 = libc.StrTok(nil, CString(" \r\t\n"))
	if v0 != nil {
		v1 = (*byte)(sub_416640())
		v2 = libc.Atoi(libc.GoString(v0)) <= 0
		v3 = int8(*(*byte)(unsafe.Add(unsafe.Pointer(v1), 100)))
		if !v2 {
			*(*byte)(unsafe.Add(unsafe.Pointer(v1), 100)) = byte(int8(int32(v3) | 32))
			return 1
		}
		*(*byte)(unsafe.Add(unsafe.Pointer(v1), 100)) = byte(int8(int32(v3) & 223))
	}
	return 1
}
func sub_432920() int32 {
	var (
		v0 *byte
		v1 *byte
	)
	_ = v1
	libc.StrTok(nil, CString(" \r\t\n"))
	v0 = libc.StrTok(nil, CString(" \r\t\n"))
	if v0 != nil {
		v1 = (*byte)(sub_416640())
		*(*byte)(unsafe.Add(unsafe.Pointer(v1), 100)) |= byte(int8(libc.Atoi(libc.GoString(v0)) & 7))
	}
	return 1
}
func sub_432970() int32 {
	var (
		v0 *byte
		v1 *byte
	)
	_ = v1
	var v2 *byte
	libc.StrTok(nil, CString(" \r\t\n"))
	v0 = libc.StrTok(nil, CString(" \r\t\n"))
	if v0 != nil {
		v1 = (*byte)(sub_416640())
		*(*uint32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v1), 66)))) = uint32(libc.Atoi(libc.GoString(v0)))
		v2 = libc.StrTok(nil, CString(" \r\t\n"))
		if v2 != nil {
			*(*uint32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v1), 70)))) = uint32(libc.Atoi(libc.GoString(v2)))
		}
	}
	return 1
}
func sub_4329D0() int32 {
	var (
		v0 *byte
		v1 *byte
	)
	_ = v1
	var v2 *byte
	libc.StrTok(nil, CString(" \r\t\n"))
	v0 = libc.StrTok(nil, CString(" \r\t\n"))
	if v0 != nil {
		v1 = (*byte)(sub_416640())
		*(*uint16)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v1), 105)))) = uint16(int16(libc.Atoi(libc.GoString(v0))))
		v2 = libc.StrTok(nil, CString(" \r\t\n"))
		if v2 != nil {
			*(*uint16)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v1), 107)))) = uint16(int16(libc.Atoi(libc.GoString(v2))))
		}
	}
	return 1
}
func sub_432A30() int32 {
	var (
		v0 *byte
		v1 *byte
		v2 bool
		v3 int8
	)
	libc.StrTok(nil, CString(" \r\t\n"))
	v0 = libc.StrTok(nil, CString(" \r\t\n"))
	if v0 != nil {
		v1 = (*byte)(sub_416640())
		v2 = libc.Atoi(libc.GoString(v0)) <= 0
		v3 = int8(*(*byte)(unsafe.Add(unsafe.Pointer(v1), 102)))
		if !v2 {
			*(*byte)(unsafe.Add(unsafe.Pointer(v1), 102)) = byte(int8(int32(v3) | 128))
			return 1
		}
		*(*byte)(unsafe.Add(unsafe.Pointer(v1), 102)) = byte(int8(int32(v3) & math.MaxInt8))
	}
	return 1
}
func sub_432A90() int32 {
	var v0 *byte
	libc.StrTok(nil, CString(" \r\t\n"))
	v0 = libc.StrTok(nil, CString(" \r\t\n"))
	if v0 != nil {
		nox_server_sendMotd_108752 = uint32(bool2int(libc.Atoi(libc.GoString(v0)) != 0))
	}
	return 1
}
func sub_432AD0(a1 *int32) int32 {
	var v1 *byte
	libc.StrTok(nil, CString(" \r\t\n"))
	v1 = libc.StrTok(nil, CString(" \r\t\n"))
	*a1 = int32(libc.Atoi(libc.GoString(v1)))
	return 1
}
func nox_common_parsecfg_videomode() int32 {
	var (
		v0 *byte
		v2 *byte
		v4 *byte
		v5 int32
		v6 int32
	)
	libc.StrTok(nil, CString(" \r\t\n"))
	v0 = libc.StrTok(nil, CString(" \r\t\n"))
	var w int32 = int32(libc.Atoi(libc.GoString(v0)))
	v2 = libc.StrTok(nil, CString(" \r\t\n"))
	var h int32 = int32(libc.Atoi(libc.GoString(v2)))
	v4 = libc.StrTok(nil, CString(" \r\t\n"))
	v5 = int32(libc.Atoi(libc.GoString(v4)))
	v6 = v5
	v6 = 16
	nox_common_parsecfg_videomode_apply(w, h, v6)
	return 1
}
func nox_common_parsecfg_videosize() int32 {
	var (
		v0 *byte
		v1 int32
	)
	libc.StrTok(nil, CString(" \r\t\n"))
	v0 = libc.StrTok(nil, CString(" \r\t\n"))
	v1 = int32(libc.Atoi(libc.GoString(v0)))
	nox_video_setCutSize_4766A0(v1)
	return 1
}
func sub_432C00() int32 {
	var v0 *byte
	libc.StrTok(nil, CString(" \r\t\n"))
	v0 = libc.StrTok(nil, CString(" \r\t\n"))
	*memmap.PtrUint32(0x587000, 80852) = uint32(libc.Atoi(libc.GoString(v0)))
	return 1
}
func sub_432C30() int32 {
	var (
		v0 *byte
		v1 int32
	)
	libc.StrTok(nil, CString(" \r\t\n"))
	v0 = libc.StrTok(nil, CString(" \r\t\n"))
	if v0 != nil {
		v1 = int32(libc.Atoi(libc.GoString(v0)))
		sub_4D0D90(bool2int(v1 != 0))
	}
	return 1
}
func sub_432C70(a1 int32) int32 {
	var (
		v1 *byte
		v2 int32
	)
	libc.StrTok(nil, CString(" \r\t\n"))
	v1 = libc.StrTok(nil, CString(" \r\t\n"))
	if v1 != nil {
		v2 = int32(libc.Atoi(libc.GoString(v1)))
		sub_4D0DC0(a1, v2)
	}
	return 1
}
func sub_432CB0(a1 int32) int32 {
	var (
		v1     *byte
		v2     uint32
		result int32
	)
	libc.StrTok(nil, CString(" \r\t\n"))
	v1 = libc.StrTok(nil, CString(" \r\t\n"))
	v2 = uint32(libc.Atoi(libc.GoString(v1)))
	if (v2 & 0x80000000) == 0 {
		if v2 > 0x4000 {
			v2 = 0x4000
		}
		sub_486670(int32(v2), a1)
		result = 1
	} else {
		sub_486670(0, a1)
		result = 1
	}
	return result
}
func sub_432D10() int32 {
	var (
		v0 *byte
		v1 *byte
		v2 *uint8
		v3 uint16
	)
	libc.StrTok(nil, CString(" \r\t\n"))
	v0 = libc.StrTok(nil, CString(" \r\t\n"))
	if v0 == nil {
		return 1
	}
	v1 = libc.StrTok(v0, CString(",\r\n"))
	if v1 == nil {
		return 1
	}
	v2 = (*uint8)(memmap.PtrOff(0x587000, 81224))
	for {
		v3 = uint16(int16(libc.Atoi(libc.GoString(v1))))
		sub_409FB0_settings(int16(uint16(*(*uint32)(unsafe.Pointer(v2)))), v3)
		v2 = (*uint8)(unsafe.Add(unsafe.Pointer(v2), 4))
		v1 = libc.StrTok(nil, CString(",\r\n"))
		if v1 == nil {
			break
		}
	}
	return 1
}
func sub_432D80() int32 {
	var (
		v0 *byte
		v1 *byte
		v2 *uint8
		v3 uint8
	)
	libc.StrTok(nil, CString(" \r\t\n"))
	v0 = libc.StrTok(nil, CString(" \r\t\n"))
	if v0 == nil {
		return 1
	}
	v1 = libc.StrTok(v0, CString(",\r\n"))
	if v1 == nil {
		return 1
	}
	v2 = (*uint8)(memmap.PtrOff(0x587000, 81224))
	for {
		v3 = uint8(int8(libc.Atoi(libc.GoString(v1))))
		sub_40A040_settings(int16(uint16(*(*uint32)(unsafe.Pointer(v2)))), v3)
		v2 = (*uint8)(unsafe.Add(unsafe.Pointer(v2), 4))
		v1 = libc.StrTok(nil, CString(",\r\n"))
		if v1 == nil {
			break
		}
	}
	return 1
}
func sub_432DF0() int32 {
	var (
		v0 *byte
		v1 int32
	)
	libc.StrTok(nil, CString(" \r\t\n"))
	v0 = libc.StrTok(nil, CString(" \r\t\n"))
	if v0 != nil {
		v1 = int32(libc.Atoi(libc.GoString(v0)))
		nox_xxx_rateUpdate_40A6D0(v1)
		if nox_xxx_rateGet_40A6C0() <= 0 || nox_xxx_rateGet_40A6C0() > 3 {
			nox_xxx_rateUpdate_40A6D0(1)
		}
	}
	return 1
}
func sub_432E50(a1 int32) *byte {
	var (
		v1     *byte
		result *byte
		v3     int32
		v4     *byte
		v5     [11]int32
	)
	libc.StrTok(nil, CString(" \r\t\n"))
	v1 = libc.StrTok(nil, CString(" \r\t\n"))
	if v1 == nil {
		return (*byte)(unsafe.Pointer(uintptr(1)))
	}
	result = libc.StrTok(v1, CString(",\r\n"))
	if result != nil {
		v3 = int32(libc.Atoi(libc.GoString(result)))
		result = libc.StrTok(nil, CString(",\r\n"))
		if result != nil {
			v5[0] = int32(libc.Atoi(libc.GoString(result)))
			result = libc.StrTok(nil, CString(",\r\n"))
			if result != nil {
				v5[1] = bool2int(*result == 49)
				result = libc.StrTok(nil, CString(",\r\n"))
				if result != nil {
					v5[2] = bool2int(*result == 49)
					result = libc.StrTok(nil, CString(",\r\n"))
					if result != nil {
						v5[3] = int32(libc.Atoi(libc.GoString(result)))
						result = libc.StrTok(nil, CString(",\r\n"))
						if result != nil {
							v5[4] = int32(libc.Atoi(libc.GoString(result)))
							result = libc.StrTok(nil, CString(",\r\n"))
							if result != nil {
								v5[5] = bool2int(*result == 49)
								result = libc.StrTok(nil, CString(",\r\n"))
								if result != nil {
									v5[6] = bool2int(*result == 49)
									result = libc.StrTok(nil, CString(",\r\n"))
									if result != nil {
										v5[7] = bool2int(*result == 49)
										result = libc.StrTok(nil, CString(",\r\n"))
										if result != nil {
											v5[8] = bool2int(*result == 49)
											result = libc.StrTok(nil, CString(",\r\n"))
											if result != nil {
												v5[9] = bool2int(*result == 49)
												v4 = libc.StrTok(nil, CString(",\r\n"))
												if v4 != nil {
													v5[10] = bool2int(*v4 == 49)
												} else {
													v5[10] = 0
												}
												sub_489FF0(a1, v3, unsafe.Pointer(&v5[0]))
												return (*byte)(unsafe.Pointer(uintptr(1)))
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	return result
}
func sub_433050() int32 {
	var (
		v0 *byte
		v1 int32
		i  *uint8
		v3 int32
	)
	libc.StrTok(nil, CString(" \r\t\n"))
	v0 = libc.StrTok(nil, CString(" \r\t\n"))
	if v0 != nil {
		v1 = int32(libc.Atoi(libc.GoString(v0)))
		if *memmap.PtrUint32(0x587000, 81248) != 0 {
			for i = (*uint8)(memmap.PtrOff(0x587000, 81248)); *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(i))), unsafe.Sizeof(uint32(0))*1))) != uint32(v1); i = (*uint8)(unsafe.Add(unsafe.Pointer(i), 8)) {
				v3 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(i))), unsafe.Sizeof(uint32(0))*2))))
				if v3 == 0 {
					return 1
				}
			}
			if v1 < 0 && v1 > 4 {
				v1 = 4
			}
			nox_server_connectionType_3596 = uint32(v1)
		}
	}
	return 1
}
func sub_4330C0() int32 {
	var (
		v0 *byte
		v1 *byte
		v2 int8
	)
	libc.StrTok(nil, CString(" \r\t\n"))
	v0 = libc.StrTok(nil, CString(" \r\t\n"))
	if v0 == nil {
		return 1
	}
	v1 = (*byte)(sub_416640())
	v2 = int8(libc.Atoi(libc.GoString(v0)))
	*(*byte)(unsafe.Add(unsafe.Pointer(v1), 104)) = byte(v2)
	if int32(uint8(v2)) <= 32 {
		if int32(uint8(v2)) < 1 {
			*(*byte)(unsafe.Add(unsafe.Pointer(v1), 104)) = 1
		}
	} else {
		*(*byte)(unsafe.Add(unsafe.Pointer(v1), 104)) = 32
	}
	nox_xxx_servSetPlrLimit_409F80(int32(uint8(*(*byte)(unsafe.Add(unsafe.Pointer(v1), 104)))))
	return 1
}
func sub_433130() int32 {
	var (
		v0 *byte
		v1 *byte
	)
	_ = v1
	var v2 int32
	libc.StrTok(nil, CString(" \r\t\n"))
	v0 = libc.StrTok(nil, CString(" \r\t\n"))
	if v0 != nil {
		v1 = (*byte)(sub_416640())
		v2 = int32(libc.Atoi(libc.GoString(v0)))
		*(*uint32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v1), 74)))) = uint32(v2)
		if v2 > 100 {
			*(*uint32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v1), 74)))) = 100
			return 1
		}
		if v2 < 0 {
			*(*uint32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v1), 74)))) = 0
		}
	}
	return 1
}
func sub_433190() int32 {
	var (
		v0 *byte
		v2 int8
	)
	v2 = 0
	libc.StrTok(nil, CString(" \r\t\n"))
	v0 = libc.StrTok(nil, CString(" \r\t\n"))
	if v0 != nil {
		v2 = int8(libc.Atoi(libc.GoString(v0)))
	}
	sub_578DE0(v2)
	return 1
}
func nox_common_skipcfgfile_4331E0(a1 *File, a2 int32) int32 {
	var (
		v2     int32
		result int32
	)
	v2 = 0
	if a2 != 0 {
		for nox_fs_fgets(a1, (*byte)(memmap.PtrOff(6112660, 806084)), 1024) {
			if libc.StrNCmp(CString("---"), (*byte)(memmap.PtrOff(6112660, 806084)), 3) == 0 {
				goto LABEL_6
			}
		}
		result = 1
	} else {
	LABEL_6:
		for {
			for {
				if !nox_fs_fgets(a1, (*byte)(memmap.PtrOff(6112660, 806084)), 1024) {
					return v2
				}
				if int32(*memmap.PtrUint8(6112660, 806084)) != 35 {
					break
				}
			}
			if libc.StrNCmp(CString("---"), (*byte)(memmap.PtrOff(6112660, 806084)), 3) == 0 {
				return 1
			}
			if nox_client_parseConfigHotkeysLine_42CF50((*byte)(memmap.PtrOff(6112660, 806084))) == 0 {
				break
			}
		}
		result = 0
	}
	return result
}
func nox_common_writecfgfile(a1 *byte) {
	var f *File = nox_fs_create_text(a1)
	if f == nil {
		return
	}
	sub_4332E0(f)
	nox_client_writeConfigHotkeys_42CDF0(f)
	nox_fs_fprintf(f, CString("---\n"))
	nox_fs_close(f)
}
func sub_4332E0(a1 *File) int32 {
	var (
		v1  *byte
		v2  int32
		v3  int32
		v4  int32
		v5  int32
		v6  *byte
		v7  uint8
		v8  *libc.WChar
		v9  int32
		v10 int32
		v11 uint32
		v12 uint32
		v13 int32
		v14 uint32
		v15 int32
		v16 int32
		v17 int32
		v18 *uint8
		v19 *uint8
		v20 int32
		v21 int32
		v23 *uint32
	)
	v1 = (*byte)(sub_416640())
	nox_fs_fprintf(a1, CString("Version = %d\n"), 65540)
	nox_fs_fprintf(a1, CString("VideoMode = %d %d %d\n"), nox_win_width_game, nox_win_height_game, nox_win_depth_game)
	nox_fs_fprintf(a1, CString("Stretched = %d\n"), g_scaled_cfg)
	nox_fs_fprintf(a1, CString("Fullscreen = %d\n"), g_fullscreen_cfg)
	v2 = nox_video_getCutSize_4766D0()
	nox_fs_fprintf(a1, CString("VideoSize = %d\n"), v2)
	nox_fs_fprintf(a1, CString("Gamma2 = %f\n"), nox_video_getGamma())
	nox_fs_fprintf(a1, CString("InputSensitivity = %f\n"), nox_input_getSensitivity())
	if sub_453070() == 1 {
		v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(dword_587000_127004)) + 4))) >> 16)
	} else {
		v3 = 0
	}
	nox_fs_fprintf(a1, CString("FXVolume = %d\n"), v3)
	if sub_44D990() == 1 {
		v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(dword_587000_122852)) + 4))) >> 16)
	} else {
		v4 = 0
	}
	nox_fs_fprintf(a1, CString("DialogVolume = %d\n"), v4)
	if sub_43DC30() == 1 {
		v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(dword_587000_93164)) + 4))) >> 16)
	} else {
		v5 = 0
	}
	nox_fs_fprintf(a1, CString("MusicVolume = %d\n"), v5)
	nox_fs_fprintf(a1, CString("LastServer = %s\n"), memmap.PtrOff(6112660, 806060))
	v6 = sub_433890()
	nox_fs_fprintf(a1, CString("ServerName = %s\n"), v6)
	nox_fs_fprintf(a1, CString("UnlockSurface = %d\n"), nox_video_dxUnlockSurface)
	nox_fs_fprintf(a1, CString("SoftShadowEdge = %d\n"), func() int32 {
		if nox_common_getEngineFlag(nox_engine_flag(NOX_ENGINE_FLAG_ENABLE_SOFT_SHADOW_EDGE)) {
			return 1
		}
		return 0
	}())
	nox_fs_fprintf(a1, CString("DrawFrontWalls = %d\n"), nox_client_drawFrontWalls_80812)
	nox_fs_fprintf(a1, CString("TranslucentFrontWalls = %d\n"), nox_client_translucentFrontWalls_805844)
	nox_fs_fprintf(a1, CString("HighResFrontWalls = %d\n"), nox_client_highResFrontWalls_80820)
	nox_fs_fprintf(a1, CString("HighResFloors = %d\n"), nox_client_highResFloors_154952)
	nox_fs_fprintf(a1, CString("LockHighResFloors = %d\n"), nox_client_lockHighResFloors_1193152)
	nox_fs_fprintf(a1, CString("TexturedFloors = %d\n"), nox_client_texturedFloors_154956)
	nox_fs_fprintf(a1, CString("TranslucentConsole = %d\n"), nox_gui_console_translucent_get())
	nox_fs_fprintf(a1, CString("RenderGlow = %d\n"), nox_client_renderGlow_805852)
	nox_fs_fprintf(a1, CString("RenderGUI = %d\n"), nox_client_renderGUI_80828)
	nox_fs_fprintf(a1, CString("FadeObjects = %d\n"), nox_client_fadeObjects_80836)
	nox_fs_fprintf(a1, CString("RenderBubbles = %d\n"), nox_client_renderBubbles_80844)
	v7 = sub_578DF0()
	nox_fs_fprintf(a1, CString("TrackData = %d\n"), v7)
	v8 = nox_xxx_sysopGetPass_40A630()
	nox_fs_fprintf(a1, CString("SysopPassword = %S\n"), v8)
	nox_fs_fprintf(a1, CString("ServerPassword = %S\n"), (*byte)(unsafe.Add(unsafe.Pointer(v1), 78)))
	nox_fs_fprintf(a1, CString("Profiled = %d\n"), nox_profiled_805856 != 0)
	nox_fs_fprintf(a1, CString("SanctuaryHelp = %d\n"), nox_server_sanctuaryHelp_54276)
	nox_fs_fprintf(a1, CString("MaxPacketLossPct = %d\n"), *memmap.PtrUint32(0x587000, 81280))
	nox_fs_fprintf(a1, CString("SendMessageOfTheDay = %d\n"), nox_server_sendMotd_108752)
	v9 = sub_4D0D70()
	nox_fs_fprintf(a1, CString("MapCycle = %d\n"), v9)
	nox_fs_fprintf(a1, CString("ConnectionType = %d\n"), nox_server_connectionType_3596)
	v10 = nox_xxx_rateGet_40A6C0()
	nox_fs_fprintf(a1, CString("InternetUpdateRate = %d\n"), v10)
	nox_fs_fprintf(a1, CString("LessonLimit ="))
	sub_4337B0(a1)
	nox_fs_fprintf(a1, CString("TimeLimit ="))
	sub_433820(a1)
	nox_fs_fprintf(a1, CString("PlayerSkeletons = %d\n"), *(*uint32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v1), 58)))))
	nox_fs_fprintf(a1, CString("BroadcastGestures = %d\n"), *(*uint32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v1), 62)))))
	v11 = uint32(nox_fs_fprintf(a1, CString("LatencyCompensation = %d %d\n"), *(*uint32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v1), 66)))), *(*uint32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v1), 70))))))
	*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v11))), 0)) = uint8(*(*byte)(unsafe.Add(unsafe.Pointer(v1), 100)))
	nox_fs_fprintf(a1, CString("Closed = %d\n"), (v11>>4)&1)
	*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v12))), 0)) = uint8(*(*byte)(unsafe.Add(unsafe.Pointer(v1), 100)))
	nox_fs_fprintf(a1, CString("Private = %d\n"), (v12>>5)&1)
	nox_fs_fprintf(a1, CString("AudioThreshold = %d\n"), *(*uint32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v1), 74)))))
	v13 = nox_xxx_servGetPlrLimit_409FA0()
	nox_fs_fprintf(a1, CString("MaxPlayers = %d\n"), v13)
	nox_fs_fprintf(a1, CString("RestrictedClasses = %d\n"), *(*byte)(unsafe.Add(unsafe.Pointer(v1), 100))&7)
	nox_fs_fprintf(a1, CString("RestrictedPing = %d %d\n"), *(*uint16)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v1), 105)))), *(*uint16)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v1), 107)))))
	nox_fs_fprintf(a1, CString("LimitMaxRes = %d\n"), int32(uint8(*(*byte)(unsafe.Add(unsafe.Pointer(v1), 102))))>>7)
	v14 = uint32(nox_xxx_getServerSubFlags_409E60())
	nox_fs_fprintf(a1, CString("CamperAlarm = %d\n"), (v14>>13)&1)
	v15 = sub_409F40(2)
	nox_fs_fprintf(a1, CString("ItemRespawn = %d\n"), v15)
	nox_fs_fprintf(a1, CString("MinKickVotes = %d\n"), *memmap.PtrUint32(0x587000, 229980))
	nox_fs_fprintf(a1, CString("ResetQuestMinVotes = %d\n"), nox_server_resetQuestMinVotes_229988)
	nox_fs_fprintf(a1, CString("KickQuestPlayerMinVotes = %d\n"), nox_server_kickQuestPlayerMinVotes_229992)
	v16 = sub_48A020(0, (*uint32)(unsafe.Pointer(&v23)))
	nox_fs_fprintf(a1, CString("LANFilters = %d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d\n"), v16, *v23, *(*uint32)(unsafe.Add(unsafe.Pointer(v23), unsafe.Sizeof(uint32(0))*1)), *(*uint32)(unsafe.Add(unsafe.Pointer(v23), unsafe.Sizeof(uint32(0))*2)), *(*uint32)(unsafe.Add(unsafe.Pointer(v23), unsafe.Sizeof(uint32(0))*3)), *(*uint32)(unsafe.Add(unsafe.Pointer(v23), unsafe.Sizeof(uint32(0))*4)), *(*uint32)(unsafe.Add(unsafe.Pointer(v23), unsafe.Sizeof(uint32(0))*5)), *(*uint32)(unsafe.Add(unsafe.Pointer(v23), unsafe.Sizeof(uint32(0))*6)), *(*uint32)(unsafe.Add(unsafe.Pointer(v23), unsafe.Sizeof(uint32(0))*7)), *(*uint32)(unsafe.Add(unsafe.Pointer(v23), unsafe.Sizeof(uint32(0))*8)), *(*uint32)(unsafe.Add(unsafe.Pointer(v23), unsafe.Sizeof(uint32(0))*9)), *(*uint32)(unsafe.Add(unsafe.Pointer(v23), unsafe.Sizeof(uint32(0))*10)))
	v17 = sub_48A020(1, (*uint32)(unsafe.Pointer(&v23)))
	nox_fs_fprintf(a1, CString("INETFilters = %d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d\n"), v17, *v23, *(*uint32)(unsafe.Add(unsafe.Pointer(v23), unsafe.Sizeof(uint32(0))*1)), *(*uint32)(unsafe.Add(unsafe.Pointer(v23), unsafe.Sizeof(uint32(0))*2)), *(*uint32)(unsafe.Add(unsafe.Pointer(v23), unsafe.Sizeof(uint32(0))*3)), *(*uint32)(unsafe.Add(unsafe.Pointer(v23), unsafe.Sizeof(uint32(0))*4)), *(*uint32)(unsafe.Add(unsafe.Pointer(v23), unsafe.Sizeof(uint32(0))*5)), *(*uint32)(unsafe.Add(unsafe.Pointer(v23), unsafe.Sizeof(uint32(0))*6)), *(*uint32)(unsafe.Add(unsafe.Pointer(v23), unsafe.Sizeof(uint32(0))*7)), *(*uint32)(unsafe.Add(unsafe.Pointer(v23), unsafe.Sizeof(uint32(0))*8)), *(*uint32)(unsafe.Add(unsafe.Pointer(v23), unsafe.Sizeof(uint32(0))*9)), *(*uint32)(unsafe.Add(unsafe.Pointer(v23), unsafe.Sizeof(uint32(0))*10)))
	if *memmap.PtrUint32(0x587000, 81168) != 0 {
		v18 = (*uint8)(memmap.PtrOff(0x587000, 81168))
		v19 = (*uint8)(memmap.PtrOff(0x587000, 81168))
		for {
			v20 = sub_4D0DE0(int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v19))), unsafe.Sizeof(uint32(0))*1)))))
			nox_fs_fprintf(a1, CString("%s = %d\n"), *(*uint32)(unsafe.Pointer(v18)), v20)
			v21 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v19))), unsafe.Sizeof(uint32(0))*2))))
			v19 = (*uint8)(unsafe.Add(unsafe.Pointer(v19), 8))
			v18 = v19
			if v21 == 0 {
				break
			}
		}
	}
	return nox_fs_fprintf(a1, CString("---\n"))
}
func sub_4337B0(a1 *File) int32 {
	var (
		v1 uint16
		v2 int32
		v3 *uint8
		v4 uint16
	)
	v1 = uint16(nox_xxx_servGamedataGet_40A020(*memmap.PtrInt16(0x587000, 81224)))
	nox_fs_fprintf(a1, CString(" %d"), v1)
	*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v2))), unsafe.Sizeof(uint16(0))*0)) = *memmap.PtrUint16(0x587000, 81228)
	if *memmap.PtrUint32(0x587000, 81228) != 0 {
		v3 = (*uint8)(memmap.PtrOff(0x587000, 81228))
		for {
			v4 = uint16(nox_xxx_servGamedataGet_40A020(int16(v2)))
			nox_fs_fprintf(a1, CString(",%d"), v4)
			v2 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v3))), unsafe.Sizeof(uint32(0))*1))))
			v3 = (*uint8)(unsafe.Add(unsafe.Pointer(v3), 4))
			if v2 == 0 {
				break
			}
		}
	}
	return nox_fs_fprintf(a1, CString("\n"))
}
func sub_433820(a1 *File) int32 {
	var (
		v1 uint8
		v2 int32
		v3 *uint8
		v4 uint8
	)
	v1 = sub_40A180(*memmap.PtrInt16(0x587000, 81224))
	nox_fs_fprintf(a1, CString(" %d"), v1)
	*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v2))), unsafe.Sizeof(uint16(0))*0)) = *memmap.PtrUint16(0x587000, 81228)
	if *memmap.PtrUint32(0x587000, 81228) != 0 {
		v3 = (*uint8)(memmap.PtrOff(0x587000, 81228))
		for {
			v4 = sub_40A180(int16(v2))
			nox_fs_fprintf(a1, CString(",%d"), v4)
			v2 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v3))), unsafe.Sizeof(uint32(0))*1))))
			v3 = (*uint8)(unsafe.Add(unsafe.Pointer(v3), 4))
			if v2 == 0 {
				break
			}
		}
	}
	return nox_fs_fprintf(a1, CString("\n"))
}
func sub_433890() *byte {
	var (
		v0 *byte
		v1 *uint8
		v2 int8
	)
	v0 = nox_xxx_serverOptionsGetServername_40A4C0()
	libc.StrNCpy((*byte)(memmap.PtrOff(6112660, 807108)), v0, 63)
	v1 = (*uint8)(memmap.PtrOff(6112660, 807108))
	if int32(*memmap.PtrUint8(6112660, 807108)) != 0 {
		for {
			if int32(*v1) == 32 {
				*v1 = 95
			}
			v2 = int8(*func() *uint8 {
				p := &v1
				*p = (*uint8)(unsafe.Add(unsafe.Pointer(*p), 1))
				return *p
			}())
			if int32(v2) == 0 {
				break
			}
		}
	}
	return (*byte)(memmap.PtrOff(6112660, 807108))
}
