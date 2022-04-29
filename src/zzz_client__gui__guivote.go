package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"unsafe"
)

func sub_48CB10(a1 int32) *uint32 {
	var (
		result *uint32
		v2     *libc.WChar
		v3     *uint32
		v4     *libc.WChar
		v5     *libc.WChar
		v6     *libc.WChar
		v7     *libc.WChar
		v8     *libc.WChar
		v9     *uint32
		v10    *libc.WChar
		v11    *libc.WChar
		v12    *libc.WChar
		v13    *uint32
		v14    *libc.WChar
		v15    *uint32
		v16    *byte
		i      *byte
		v18    *uint32
		v19    int32
		v20    *libc.WChar
		j      *byte
		v22    int32
		v23    *libc.WChar
		v24    *libc.WChar
		v25    int32
		v26    *byte
		v27    [256]libc.WChar
	)
	v25 = 0
	(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1197316))))).Func94(asWindowEvent(0x400F, 0, 0))
	(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1197320))))).Func94(asWindowEvent(0x400F, 0, 0))
	result = (*uint32)(unsafe.Pointer(uintptr(a1)))
	dword_5d4594_1197308 = uint32(a1)
	switch a1 {
	case 4:
		if *memmap.PtrUint32(0x8531A0, 2576) != 0 && *(*uint32)(unsafe.Pointer(uintptr(*memmap.PtrUint32(0x8531A0, 2576) + 4792))) == 0 {
			v24 = strMan.GetStringInFile("GUIVote.c:NotAllowedVote", "C:\\NoxPost\\src\\client\\Gui\\GUIVote.c")
			v2 = strMan.GetStringInFile("guiquit.c:Vote", "C:\\NoxPost\\src\\client\\Gui\\GUIVote.c")
			return (*uint32)(NewDialogWindow(nil, (*libc.WChar)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v2)))))), (*libc.WChar)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v24)))))), 33, nil, nil))
		}
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1197316))))).Hide()
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1197320))))).Show()
		v3 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1197312)))).ChildByID(4301)))
		v4 = strMan.GetStringInFile("SelectVoteTopic", "C:\\NoxPost\\src\\client\\Gui\\GUIVote.c")
		sub_46AEE0(int32(uintptr(unsafe.Pointer(v3))), int32(uintptr(unsafe.Pointer(v4))))
		v5 = strMan.GetStringInFile("VoteTopicLabel", "C:\\NoxPost\\src\\client\\Gui\\GUIVote.c")
		nox_wcscpy(&v27[0], v5)
		nox_wcscat(&v27[0], libc.CWString(" "))
		v6 = strMan.GetStringInFile("VoteResetServer", "C:\\NoxPost\\src\\client\\Gui\\GUIVote.c")
		nox_wcscat(&v27[0], v6)
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1197320))))).Func94(asWindowEvent(0x400D, int32(uintptr(unsafe.Pointer(&v27[0]))), 4))
		v7 = strMan.GetStringInFile("VoteTopicLabel", "C:\\NoxPost\\src\\client\\Gui\\GUIVote.c")
		nox_wcscpy(&v27[0], v7)
		nox_wcscat(&v27[0], libc.CWString(" "))
		v8 = strMan.GetStringInFile("VoteKickPlayer", "C:\\NoxPost\\src\\client\\Gui\\GUIVote.c")
		nox_wcscat(&v27[0], v8)
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1197320))))).Func94(asWindowEvent(0x400D, int32(uintptr(unsafe.Pointer(&v27[0]))), 4))
		return (*uint32)(unsafe.Pointer(uintptr(nox_xxx_wndShowModalMB((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1197312)))))))))
	case 2:
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1197316))))).Hide()
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1197320))))).Show()
		v9 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1197312)))).ChildByID(4301)))
		v10 = strMan.GetStringInFile("Vote:ResetQuest", "C:\\NoxPost\\src\\client\\Gui\\GUIVote.c")
		sub_46AEE0(int32(uintptr(unsafe.Pointer(v9))), int32(uintptr(unsafe.Pointer(v10))))
		v11 = strMan.GetStringInFile("WindowDir:Yes", "C:\\NoxPost\\src\\client\\Gui\\GUIVote.c")
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1197320))))).Func94(asWindowEvent(0x400D, int32(uintptr(unsafe.Pointer(v11))), 4))
		v12 = strMan.GetStringInFile("WindowDir:No", "C:\\NoxPost\\src\\client\\Gui\\GUIVote.c")
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1197320))))).Func94(asWindowEvent(0x400D, int32(uintptr(unsafe.Pointer(v12))), 4))
		if dword_5d4594_1197332 == 1 {
			(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1197320))))).Func94(asWindowEvent(0x4013, 0, 0))
		} else {
			(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1197320))))).Func94(asWindowEvent(0x4013, 1, 0))
		}
		return (*uint32)(unsafe.Pointer(uintptr(nox_xxx_wndShowModalMB((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1197312)))))))))
	case 0:
		fallthrough
	case 1:
		fallthrough
	case 3:
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1197316))))).Show()
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1197320))))).Hide()
		v13 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1197312)))).ChildByID(4301)))
		v14 = strMan.GetStringInFile("VoteKickPlayer", "C:\\NoxPost\\src\\client\\Gui\\GUIVote.c")
		sub_46AEE0(int32(uintptr(unsafe.Pointer(v13))), int32(uintptr(unsafe.Pointer(v14))))
		if int32(nox_xxx_getTeamCounter_417DD0()) != 0 {
			v15 = nox_xxx_objGetTeamByNetCode_418C80(int32(nox_player_netCode_85319C))
			v16 = (*byte)(unsafe.Pointer(nox_xxx_clientGetTeamColor_418AB0(int32(*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(v15))), 4)))))))
			v26 = v16
			if v16 != nil {
				for i = (*byte)(unsafe.Pointer(nox_common_playerInfoGetFirst_416EA0())); i != nil; i = (*byte)(unsafe.Pointer(nox_common_playerInfoGetNext_416EE0((*nox_playerInfo)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(i))))))))) {
					if i != *(**byte)(memmap.PtrOff(0x8531A0, 2576)) {
						v18 = nox_xxx_objGetTeamByNetCode_418C80(int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(i))), unsafe.Sizeof(uint32(0))*515)))))
						if v18 != nil {
							if nox_xxx_teamCompare2_419180(int32(uintptr(unsafe.Pointer(v18))), uint8(*(*byte)(unsafe.Add(unsafe.Pointer(v16), 57)))) != 0 {
								(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1197316))))).Func94(asWindowEvent(0x400D, int32(uintptr(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(i), 4704))))), int32(*memmap.PtrUint32(0x587000, uint32((int32(uint8(*(*byte)(unsafe.Add(unsafe.Pointer(v16), 57))))%10)*8)+156400))))
								v19 = 0
								if dword_5d4594_1197324 > 0 {
									v20 = (*libc.WChar)(memmap.PtrOff(6112660, 1193720))
									for {
										if nox_wcscmp(v20, (*libc.WChar)(unsafe.Add(unsafe.Pointer((*libc.WChar)(unsafe.Pointer(i))), unsafe.Sizeof(libc.WChar(0))*2352))) == 0 {
											(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1197316))))).Func94(asWindowEvent(0x4015, v25, 0))
										}
										v19++
										v20 = (*libc.WChar)(unsafe.Add(unsafe.Pointer(v20), unsafe.Sizeof(libc.WChar(0))*28))
										if v19 >= *(*int32)(unsafe.Pointer(&dword_5d4594_1197324)) {
											break
										}
									}
									v16 = v26
								}
								v25++
							}
						}
					}
				}
			}
		} else {
			for j = (*byte)(unsafe.Pointer(nox_common_playerInfoGetFirst_416EA0())); j != nil; j = (*byte)(unsafe.Pointer(nox_common_playerInfoGetNext_416EE0((*nox_playerInfo)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(j))))))))) {
				if j != *(**byte)(memmap.PtrOff(0x8531A0, 2576)) {
					(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1197316))))).Func94(asWindowEvent(0x400D, int32(uintptr(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(j), 4704))))), 4))
					v22 = 0
					if dword_5d4594_1197324 > 0 {
						v23 = (*libc.WChar)(memmap.PtrOff(6112660, 1193720))
						for {
							if nox_wcscmp(v23, (*libc.WChar)(unsafe.Add(unsafe.Pointer((*libc.WChar)(unsafe.Pointer(j))), unsafe.Sizeof(libc.WChar(0))*2352))) == 0 {
								(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1197316))))).Func94(asWindowEvent(0x4015, v25, 0))
							}
							v22++
							v23 = (*libc.WChar)(unsafe.Add(unsafe.Pointer(v23), unsafe.Sizeof(libc.WChar(0))*28))
							if v22 >= *(*int32)(unsafe.Pointer(&dword_5d4594_1197324)) {
								break
							}
						}
					}
					v25++
				}
			}
		}
		return (*uint32)(unsafe.Pointer(uintptr(nox_xxx_wndShowModalMB((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1197312)))))))))
	}
	return result
}
