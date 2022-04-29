package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	noxflags "github.com/noxworld-dev/opennox/v1/common/flags"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"unsafe"
)

func nox_xxx_guiSpelllistLoad_453850(a1 int32) int32 {
	var (
		v9  *uint32
		v10 *uint32
	)
	dword_5d4594_1045484 = uint32(uintptr(unsafe.Pointer(newWindowFromFile("spelllst.wnd", unsafe.Pointer(cgoFuncAddr(sub_453C00))))))
	(*(*int32)(unsafe.Pointer(&dword_5d4594_1045484))).setDraw(sub_453B80)
	sub_46B120((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045484)))), (*nox_window)(unsafe.Pointer(uintptr(a1))))
	nox_xxx_wnd_46B280(*(*int32)(unsafe.Pointer(&dword_5d4594_1045484)), a1)
	dword_5d4594_1045480 = uint32(uintptr(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045484)))).ChildByID(1110))))
	dword_5d4594_1045508 = uint32(uintptr(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045484)))).ChildByID(1112))))
	sub_453B00()
	(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045480))))).Func94(asWindowEvent(0x400F, 0, 0))
	(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045508))))).Func94(asWindowEvent(0x400F, 0, 0))
	var wbuf [64]libc.WChar = [64]libc.WChar{}
	for i := int32(1); i < NOX_SPELLS_MAX; i++ {
		if !nox_xxx_spellIsValid_424B50(i) {
			continue
		}
		var flags int32 = int32(nox_xxx_spellFlags_424A70(i))
		if (uint32(flags) & 0x15000) != 0 {
			continue
		}
		if uint32(flags)&0x1000000 != 0 || uint32(flags)&0x2000000 != 0 && uint32(flags)&0x4000000 != 0 {
			var v7 *libc.WChar = strMan.GetStringInFile("Common", "C:\\NoxPost\\src\\client\\Gui\\ServOpts\\spelllst.c")
			nox_wcscpy(&wbuf[0], v7)
		} else {
			if (uint32(flags) & 0x6000000) == 0 {
				continue
			}
			wbuf[0] = 0
			if uint32(flags)&0x2000000 != 0 {
				var v5 *libc.WChar = strMan.GetStringInFile("SpellWizard", "C:\\NoxPost\\src\\client\\Gui\\ServOpts\\spelllst.c")
				nox_wcscat(&wbuf[0], v5)
			}
			if uint32(flags)&0x4000000 != 0 {
				var v6 *libc.WChar = strMan.GetStringInFile("SpellConjurer", "C:\\NoxPost\\src\\client\\Gui\\ServOpts\\spelllst.c")
				nox_wcscat(&wbuf[0], v6)
			}
		}
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045508))))).Func94(asWindowEvent(0x400D, int32(uintptr(unsafe.Pointer(&wbuf[0]))), -1))
		var title *libc.WChar = nox_xxx_spellTitle_424930(i)
		nox_wcsncpy(&wbuf[0], title, uint32(unsafe.Sizeof([64]libc.WChar{})/2-1))
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045480))))).Func94(asWindowEvent(0x400D, int32(uintptr(unsafe.Pointer(&wbuf[0]))), -1))
	}
	v9 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045484)))).ChildByID(1113)))
	(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045480))))).Func94(asWindowEvent(0x4018, int32(uintptr(unsafe.Pointer(v9))), 0))
	(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045508))))).Func94(asWindowEvent(0x4018, int32(uintptr(unsafe.Pointer(v9))), 0))
	v10 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045484)))).ChildByID(1114)))
	(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045480))))).Func94(asWindowEvent(0x4019, int32(uintptr(unsafe.Pointer(v10))), 0))
	(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045508))))).Func94(asWindowEvent(0x4019, int32(uintptr(unsafe.Pointer(v10))), 0))
	sub_454040((*uint32)(memmap.PtrOff(6112660, 1045488)))
	sub_454120()
	if !noxflags.HasGame(noxflags.GameHost) || noxflags.HasGame(noxflags.GameFlag15|noxflags.GameFlag16) {
		sub_46AD20(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045484)), 1115, 1133, 0)
	}
	return int32(dword_5d4594_1045484)
}
func sub_453C00(a1 int32, a2 int32, a3 *int32, a4 int32) int32 {
	var (
		v3     *int32
		v4     int32
		v5     *int16
		v6     int32
		v7     int32
		v8     int32
		i      int32
		result int32
		v11    int32
		v12    *libc.WChar
		v13    int32
		v14    *libc.WChar
		v15    *byte
		v16    *libc.WChar
		v17    *libc.WChar
		v18    *libc.WChar
		v19    [15]int32
		v20    *byte
		v21    int32
	)
	if a2 == 0x4000 {
		if unsafe.Pointer(a3) == unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045484)))).ChildByID(1113)) || unsafe.Pointer(a3) == unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045484)))).ChildByID(1114)) {
			(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045480))))).Func94(asWindowEvent(0x4000, int32(uintptr(unsafe.Pointer(a3))), 0))
			(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045508))))).Func94(asWindowEvent(0x4000, int32(uintptr(unsafe.Pointer(a3))), 0))
		LABEL_37:
			sub_454120()
		}
		return 0
	}
	if a2 != 0x4007 {
		return 0
	}
	v3 = a3
	v4 = (*nox_window)(unsafe.Pointer(a3)).ID()
	v21 = v4
	switch v4 {
	case 1113:
		fallthrough
	case 1114:
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045480))))).Func94(asWindowEvent(0x4000, int32(uintptr(unsafe.Pointer(v3))), 0))
		(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045508))))).Func94(asWindowEvent(0x4000, int32(uintptr(unsafe.Pointer(v3))), 0))
		goto LABEL_37
	case 1115:
		fallthrough
	case 1116:
		v5 = *(**int16)(unsafe.Pointer(uintptr(dword_5d4594_1045480 + 32)))
		v6 = 0
		v20 = sub_4165B0()
		if int32(*v5) > 0 {
			v7 = 0
			for {
				if *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v5))), unsafe.Sizeof(uint32(0))*6)))+uint32(v7) != 0xFFFFFFFC {
					v8 = nox_xxx_spellByTitle_424960((*libc.WChar)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v5))), unsafe.Sizeof(uint32(0))*6))) + uint32(v7) + 4))))
					if v8 != 0 {
						if v21 == 1115 {
							if !noxflags.HasGame(noxflags.GameModeFlagBall) && (*(*byte)(unsafe.Add(unsafe.Pointer(v20), 52))&64) == 0 || v8 != 132 {
								sub_453FA0(int32(uintptr(memmap.PtrOff(6112660, 1045488))), v8, 1)
							}
						} else {
							sub_453FA0(int32(uintptr(memmap.PtrOff(6112660, 1045488))), v8, 0)
						}
					}
				}
				v6++
				v7 += 524
				if v6 >= int32(*v5) {
					break
				}
			}
		}
		if dword_5d4594_2650652 != 0 {
			sub_57A1E0(&v19[0], nil, nil, 4, 6128)
			for i = 0; i < 5; i++ {
				*memmap.PtrUint32(6112660, uint32(i*4)+0xFF3F0) &= uint32(v19[i+6])
			}
		}
		sub_454120()
		goto LABEL_19
	case 1120:
		fallthrough
	case 1121:
		fallthrough
	case 1122:
		fallthrough
	case 1123:
		fallthrough
	case 1124:
		fallthrough
	case 1125:
		fallthrough
	case 1126:
		fallthrough
	case 1127:
		fallthrough
	case 1128:
		fallthrough
	case 1129:
		fallthrough
	case 1130:
		fallthrough
	case 1131:
		fallthrough
	case 1132:
		fallthrough
	case 1133:
		v11 = int32(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1045480 + 32))))
		v12 = (*libc.WChar)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v11 + 24))) + uint32((sub_4A4800(int32(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1045480 + 32)))))+v4-1120)*524) + 4)))
		if v12 == nil {
			goto LABEL_19
		}
		v13 = nox_xxx_spellByTitle_424960(v12)
		if v13 == 0 {
			goto LABEL_19
		}
		if dword_5d4594_2650652 == 0 || (func() int32 {
			sub_57A1E0(&v19[0], nil, nil, 4, 6128)
			return sub_454000(int32(uintptr(unsafe.Pointer(&v19[6]))), v13)
		}()) != 0 {
			v15 = sub_4165B0()
			if (noxflags.HasGame(noxflags.GameModeFlagBall) || *(*byte)(unsafe.Add(unsafe.Pointer(v15), 52))&64 != 0) && v13 == 132 {
				*(*int32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(int32(0))*9)) ^= 4
				v18 = strMan.GetStringInFile("plyrspel.c:Illegal", "C:\\NoxPost\\src\\client\\Gui\\ServOpts\\spelllst.c")
				v16 = strMan.GetStringInFile("Notice", "C:\\NoxPost\\src\\client\\Gui\\ServOpts\\spelllst.c")
				NewDialogWindow((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045484))))), (*libc.WChar)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v16)))))), (*libc.WChar)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v18)))))), 33, nil, nil)
				sub_44A360(1)
			} else {
				if *(*int32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(int32(0))*9))&4 != 0 {
					sub_453FA0(int32(uintptr(memmap.PtrOff(6112660, 1045488))), v13, 0)
				} else {
					sub_453FA0(int32(uintptr(memmap.PtrOff(6112660, 1045488))), v13, 1)
				}
			LABEL_19:
				sub_459D50(1)
			}
		} else {
			*(*int32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(int32(0))*9)) ^= 4
			v17 = strMan.GetStringInFile("NotInternet", "C:\\NoxPost\\src\\client\\Gui\\ServOpts\\spelllst.c")
			v14 = strMan.GetStringInFile("Notice", "C:\\NoxPost\\src\\client\\Gui\\ServOpts\\spelllst.c")
			NewDialogWindow((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1045484))))), (*libc.WChar)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v14)))))), (*libc.WChar)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v17)))))), 33, nil, nil)
			sub_44A360(1)
		}
	LABEL_20:
		clientPlaySoundSpecial(766, 100)
		result = 0
	default:
		goto LABEL_20
	}
	return result
}
