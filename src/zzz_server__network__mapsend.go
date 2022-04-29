package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	"github.com/gotranspile/cxgo/runtime/stdio"
	"github.com/noxworld-dev/opennox/v1/common/alloc"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"unsafe"
)

func nox_xxx_netSendMap_5199F0_net_mapsend(a1 *uint8) {
	var (
		v1  *byte
		v2  *libc.WChar
		v3  *uint16
		v4  uint32
		v5  int32
		v6  int32
		v7  uint32
		v8  uint32
		v9  uint32
		v10 int32
		v11 *byte
		v12 *libc.WChar
		v13 int32
		v14 *libc.WChar
		v15 int32
		v16 int32
		v17 int32
		v18 int32
		v19 [88]byte
		v20 [128]libc.WChar
	)
	if a1 != nil {
		if *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a1))), unsafe.Sizeof(uint32(0))*6))) == 0 {
			v1 = (*byte)(unsafe.Pointer(nox_common_playerInfoFromNum_417090(int32(*a1))))
			if v1 != nil {
				v15 = int32(uintptr(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v1), 4704)))))
				v2 = strMan.GetStringInFile("Downloadingmap", "C:\\NoxPost\\src\\Server\\Network\\mapsend.c")
				nox_swprintf(&v20[0], v2, memmap.PtrOff(6112660, 2387068), v15)
				nox_gui_console_Print_450B90(uint8(int8(NOX_CONSOLE_RED)), (*libc.WChar)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(&v20[0])))))))
			}
			*(*uint32)(unsafe.Pointer(&v19[4])) = *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a1))), unsafe.Sizeof(uint32(0))*3)))
			v19[0] = 184
			libc.StrCpy(&v19[8], (*byte)(memmap.PtrOff(6112660, 2387068)))
			nox_xxx_netSendSock_552640(uint32(int32(*a1)+1), &v19[0], 88, int8(NOX_NET_SEND_NO_LOCK|NOX_NET_SEND_FLAG2))
		}
		v3 = (*uint16)(alloc.Calloc(1, int(int32(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(a1))), unsafe.Sizeof(uint16(0))*10))))+6)))
		if v3 != nil {
			v4 = uint32(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(a1))), unsafe.Sizeof(uint16(0))*10))))
			if uint32(uint16(v4)) >= (*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a1))), unsafe.Sizeof(uint32(0))*3))) - *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a1))), unsafe.Sizeof(uint32(0))*6)))) {
				v17 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a1))), unsafe.Sizeof(uint32(0))*3))) - *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a1))), unsafe.Sizeof(uint32(0))*6))))
				v4 = *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a1))), unsafe.Sizeof(uint32(0))*3))) - *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a1))), unsafe.Sizeof(uint32(0))*6)))
			} else {
				v17 = int32(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(a1))), unsafe.Sizeof(uint16(0))*10))))
			}
			*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v18))), 0)) = 185
			*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v18))), unsafe.Sizeof(uint16(0))*1)) = *((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(a1))), unsafe.Sizeof(uint16(0))*8)))
			*(*uint32)(unsafe.Pointer(v3)) = uint32(v18)
			*(*uint16)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint16(0))*2)) = uint16(v4)
			if int32(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(a1))), unsafe.Sizeof(uint16(0))*2)))) != 0 {
				v5 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a1))), unsafe.Sizeof(uint32(0))*2))))
				v6 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a1))), unsafe.Sizeof(uint32(0))*6))))
			} else {
				v5 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a1))), unsafe.Sizeof(uint32(0))*6))))
				v6 = int32(dword_5d4594_2388640)
			}
			libc.MemCpy(unsafe.Pointer((*uint16)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint16(0))*3))), unsafe.Pointer(uintptr(v6+v5)), int(v4))
			nox_xxx_netSendSock_552640(uint32(int32(*a1)+1), (*byte)(unsafe.Pointer(v3)), int32(v4+6), int8(NOX_NET_SEND_NO_LOCK|NOX_NET_SEND_FLAG2))
			alloc.Free(unsafe.Pointer(v3))
			v7 = *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a1))), unsafe.Sizeof(uint32(0))*3)))
			v8 = uint32(v17) + *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a1))), unsafe.Sizeof(uint32(0))*6)))
			*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a1))), unsafe.Sizeof(uint32(0))*4)))++
			*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a1))), unsafe.Sizeof(uint32(0))*6))) = v8
			if v8 < v7 {
				*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(a1))), unsafe.Sizeof(uint16(0))*1))) = 1
			} else {
				v9 = uint32((uint64(platformTicks()) - *((*uint64)(unsafe.Add(unsafe.Pointer((*uint64)(unsafe.Pointer(a1))), unsafe.Sizeof(uint64(0))*5)))) / 1000)
				if v9 != 0 {
					v10 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a1))), unsafe.Sizeof(uint32(0))*3))) / v9)
				} else {
					v10 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a1))), unsafe.Sizeof(uint32(0))*3))))
				}
				v11 = (*byte)(unsafe.Pointer(nox_common_playerInfoFromNum_417090(int32(*a1))))
				if v11 != nil {
					v12 = strMan.GetStringInFile("downloaddone", "C:\\NoxPost\\src\\Server\\Network\\mapsend.c")
					nox_swprintf(&v20[0], v12, *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a1))), unsafe.Sizeof(uint32(0))*6))), *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a1))), unsafe.Sizeof(uint32(0))*4)))-1, v10, v9, (*byte)(unsafe.Add(unsafe.Pointer(v11), 4704)))
					nox_gui_console_Print_450B90(uint8(int8(NOX_CONSOLE_RED)), (*libc.WChar)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(&v20[0])))))))
				}
				if int32(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(a1))), unsafe.Sizeof(uint16(0))*2)))) == 1 && *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a1))), unsafe.Sizeof(uint32(0))*2))) != 0 {
					*((*unsafe.Pointer)(unsafe.Add(unsafe.Pointer((*unsafe.Pointer)(unsafe.Pointer(a1))), unsafe.Sizeof(unsafe.Pointer(nil))*2))) = nil
				}
				v13 = int32(dword_5d4594_2388648)
				if dword_5d4594_2388648 != 0 {
					v13 = int32(func() uint32 {
						p := &dword_5d4594_2388648
						*p--
						return *p
					}())
				}
				v16 = v13
				v14 = strMan.GetStringInFile("InProgress", "C:\\NoxPost\\src\\Server\\Network\\mapsend.c")
				nox_swprintf(&v20[0], v14, v16)
				nox_gui_console_Print_450B90(uint8(int8(NOX_CONSOLE_RED)), (*libc.WChar)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(&v20[0])))))))
				nox_xxx_netMapSendClear_519830(int32(uintptr(unsafe.Pointer(a1))), int8(*a1))
			}
		}
	}
}
func nox_xxx_netSendMapAbort_519C80_net_mapsend(a1 *uint8, a2 uint8) {
	var (
		v2 int32
		v3 *libc.WChar
		v4 int32
		v5 [2]byte
		v6 [128]libc.WChar
	)
	if a1 != nil {
		v2 = int32(*a1)
		v5[0] = 186
		v5[1] = byte(a2)
		nox_xxx_netSendSock_552640(uint32(v2+1), &v5[0], 2, int8(NOX_NET_SEND_NO_LOCK|NOX_NET_SEND_FLAG2))
		if dword_5d4594_2388648 != 0 {
			dword_5d4594_2388648--
		}
		v4 = int32(*a1)
		v3 = strMan.GetStringInFile("downloadaborted", "C:\\NoxPost\\src\\Server\\Network\\mapsend.c")
		nox_swprintf(&v6[0], v3, v4, a2)
		nox_gui_console_Print_450B90(uint8(int8(NOX_CONSOLE_RED)), (*libc.WChar)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(&v6[0])))))))
	}
}
func nox_xxx_netMapSend_519D20(a1 int32) *byte {
	var (
		v1     *uint8
		v2     *libc.WChar
		result *byte
	)
	v1 = (*uint8)(memmap.PtrOff(6112660, uint32(a1*48)+0x246CCC))
	dword_5d4594_2388648++
	if int32(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v1))), unsafe.Sizeof(uint16(0))*1)))) != 0 {
		if int32(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v1))), unsafe.Sizeof(uint16(0))*1)))) == 1 {
			v2 = strMan.GetStringInFile("Sending", "C:\\NoxPost\\src\\Server\\Network\\mapsend.c")
		} else {
			v2 = strMan.GetStringInFile("BadState", "C:\\NoxPost\\src\\Server\\Network\\mapsend.c")
		}
		result = (*byte)(unsafe.Pointer(uintptr(nox_gui_console_Print_450B90(uint8(int8(NOX_CONSOLE_RED)), (*libc.WChar)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v2))))))))))
	} else {
		*memmap.PtrUint16(6112660, 2388636)++
		*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v1))), unsafe.Sizeof(uint16(0))*1))) = 1
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v1))), unsafe.Sizeof(uint32(0))*6))) = 0
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v1))), unsafe.Sizeof(uint32(0))*4))) = 1
		*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v1))), unsafe.Sizeof(uint16(0))*10))) = 512
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v1))), unsafe.Sizeof(uint32(0))*3))) = *memmap.PtrUint32(6112660, 2388644)
		*((*uint64)(unsafe.Add(unsafe.Pointer((*uint64)(unsafe.Pointer(v1))), unsafe.Sizeof(uint64(0))*5))) = uint64(platformTicks())
		result = (*byte)(unsafe.Pointer(nox_common_playerInfoFromNum_417090(a1)))
	}
	return result
}
func nox_xxx_netMapSendCancelMap_519DE0_net_mapsend(a1 int32) int32 {
	var (
		v1     *uint8
		v2     *byte
		result int32
		v4     *libc.WChar
	)
	v1 = (*uint8)(memmap.PtrOff(6112660, uint32(a1*48)+0x246CCC))
	v2 = (*byte)(unsafe.Pointer(nox_common_playerInfoFromNum_417090(a1)))
	*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v1))), unsafe.Sizeof(uint16(0))*1))) = 0
	nullsub_27(uint32(uintptr(unsafe.Pointer(v1))))
	*memmap.PtrUint16(6112660, 2388638)++
	if int32(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v1))), unsafe.Sizeof(uint16(0))*2)))) == 1 && *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v1))), unsafe.Sizeof(uint32(0))*2))) != 0 {
		*((*unsafe.Pointer)(unsafe.Add(unsafe.Pointer((*unsafe.Pointer)(unsafe.Pointer(v1))), unsafe.Sizeof(unsafe.Pointer(nil))*2))) = nil
	}
	nox_xxx_netMapSendClear_519830(int32(uintptr(unsafe.Pointer(v1))), int8(a1))
	result = int32(dword_5d4594_2388648)
	if dword_5d4594_2388648 != 0 {
		result = int32(func() uint32 {
			p := &dword_5d4594_2388648
			*p--
			return *p
		}())
	}
	if v2 != nil {
		v4 = strMan.GetStringInFile("downloadcancelled", "C:\\NoxPost\\src\\Server\\Network\\mapsend.c")
		result = nox_gui_console_Printf_450C00(uint8(int8(NOX_CONSOLE_RED)), v4, (*byte)(unsafe.Add(unsafe.Pointer(v2), 4704)))
	}
	return result
}
func nox_xxx_netMapSendPrepair_519EB0_net_mapsend() int32 {
	var (
		v0  *byte
		v1  *uint8
		v2  unsafe.Pointer
		v3  unsafe.Pointer
		v4  uint32
		v5  *libc.WChar
		v6  *byte
		v7  uint8
		v8  *File
		v10 *libc.WChar
		v11 *byte
		v12 [256]byte
	)
	v0 = nox_server_currentMapGetFilename_409B30()
	v11 = v0
	if dword_5d4594_2388648 != 0 {
		v1 = (*uint8)(memmap.PtrOff(6112660, 2387156))
		for {
			if int32(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v1))), -int(unsafe.Sizeof(uint16(0))*3))))) == 1 {
				v2 = alloc.Calloc(1, int(*(*uint32)(memmap.PtrOff(6112660, 2388644))))
				*(*uint32)(unsafe.Pointer(v1)) = uint32(uintptr(v2))
				if v2 != nil {
					v3 = *(*unsafe.Pointer)(unsafe.Pointer(&dword_5d4594_2388640))
					v4 = *memmap.PtrUint32(6112660, 2388644)
					*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v1))), -int(unsafe.Sizeof(uint16(0))*2)))) = 1
					libc.MemCpy(v2, v3, int(v4))
					*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v1))), unsafe.Sizeof(uint32(0))*1))) = *memmap.PtrUint32(6112660, 2388644)
					v5 = strMan.GetStringInFile("ForceCopy", "C:\\NoxPost\\src\\Server\\Network\\mapsend.c")
					nox_gui_console_Print_450B90(uint8(int8(NOX_CONSOLE_RED)), (*libc.WChar)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v5)))))))
				} else {
					nox_xxx_netSendMapAbort_519C80_net_mapsend((*uint8)(unsafe.Add(unsafe.Pointer(v1), -8)), 1)
				}
			}
			v1 = (*uint8)(unsafe.Add(unsafe.Pointer(v1), 48))
			if int32(uintptr(unsafe.Pointer(v1))) >= int32(uintptr(memmap.PtrOff(6112660, 2388644))) {
				break
			}
		}
		v0 = v11
	}
	if v0 == nil {
		v10 = strMan.GetStringInFile("CompressFail", "C:\\NoxPost\\src\\Server\\Network\\mapsend.c")
		nox_gui_console_Print_450B90(uint8(int8(NOX_CONSOLE_RED)), (*libc.WChar)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v10)))))))
		return 0
	}
	if libc.MemCmp(unsafe.Pointer(v0), memmap.PtrOff(6112660, 2388652), 1) == 0 {
		v10 = strMan.GetStringInFile("CompressFail", "C:\\NoxPost\\src\\Server\\Network\\mapsend.c")
		nox_gui_console_Print_450B90(uint8(int8(NOX_CONSOLE_RED)), (*libc.WChar)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v10)))))))
		return 0
	}
	libc.StrCpy(&v12[0], CString("maps\\"))
	libc.StrCat(&v12[0], nox_xxx_mapGetMapName_409B40())
	*(*uint16)(unsafe.Pointer(&v12[libc.StrLen(&v12[0])])) = *memmap.PtrUint16(0x587000, 249808)
	libc.StrCat(&v12[0], nox_xxx_mapGetMapName_409B40())
	v6 = &v12[libc.StrLen(&v12[0])+1]
	v7 = *memmap.PtrUint8(0x587000, 249816)
	*(*uint32)(unsafe.Pointer(func() *byte {
		p := &v6
		*p = (*byte)(unsafe.Add(unsafe.Pointer(*p), -1))
		return *p
	}())) = *memmap.PtrUint32(0x587000, 249812)
	*(*byte)(unsafe.Add(unsafe.Pointer(v6), 4)) = byte(v7)
	v8 = nox_fs_open(&v12[0])
	if v8 == nil {
		v10 = strMan.GetStringInFile("CompressFail", "C:\\NoxPost\\src\\Server\\Network\\mapsend.c")
		nox_gui_console_Print_450B90(uint8(int8(NOX_CONSOLE_RED)), (*libc.WChar)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v10)))))))
		return 0
	}
	libc.StrCpy((*byte)(memmap.PtrOff(6112660, 2386988)), &v12[0])
	libc.StrCpy((*byte)(memmap.PtrOff(6112660, 2387068)), v11)
	*memmap.PtrUint32(6112660, 2388644) = uint32(nox_fs_fsize(v8))
	dword_5d4594_2388640 = uint32(uintptr(alloc.Calloc(1, int(*(*uint32)(memmap.PtrOff(6112660, 2388644))))))
	nox_binfile_fread_raw_40ADD0(*(**byte)(unsafe.Pointer(&dword_5d4594_2388640)), 1, *(*uint32)(memmap.PtrOff(6112660, 2388644)), v8)
	nox_fs_close(v8)
	return 1
}
