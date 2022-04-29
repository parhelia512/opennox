package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	noxflags "github.com/noxworld-dev/opennox/v1/common/flags"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"unsafe"
)

func sub_4A4A20(a1 int32, a2 int32, a3 *int32, a4 int32) int32 {
	var (
		v4 int32
		v5 int32
		v6 int32
		v7 int32
		v8 *uint32
		v9 *libc.WChar
	)
	if a2 != 0x4005 {
		if a2 != 0x4007 {
			return 0
		}
		v4 = (*nox_window)(unsafe.Pointer(a3)).ID()
		if v4 >= 601 {
			if v4 <= 603 {
				return 1
			}
			if v4 == 610 {
				if noxflags.HasGame(noxflags.GameOnline) && !noxflags.HasGame(noxflags.GameModeQuest) {
					if nox_xxx_isQuest_4D6F50() != 0 || (func() int32 {
						v5 = sub_4D6F70()
						return v5
					}()) != 0 {
						v5 = 1
					}
					sub_4A4B70(v5)
				}
				sub_4A4970()
				nox_wnd_xxx_1307732.field_13 = nox_game_showSelColor_4A5D00
			}
		}
		clientPlaySoundSpecial(921, 100)
		return 1
	}
	v6 = (*nox_window)(unsafe.Pointer(a3)).ID()
	v7 = v6
	if v6 >= 601 && v6 <= 603 {
		nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(*memmap.PtrInt32(6112660, 1307728)))), 1)
		v8 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1307736)))).ChildByID(605)))
		*(*uint8)(unsafe.Pointer(uintptr(dword_5d4594_1307724 + 66))) = uint8(int8(v7 - 89))
		v9 = strMan.GetStringInFile(*(**byte)(memmap.PtrOff(0x587000, uint32(int32(uint8(int8(v7-89)))*4)+0x298E0)), "C:\\NoxPost\\src\\client\\shell\\SelClass.c")
		(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v8)))))).Func94(asWindowEvent(0x4001, int32(uintptr(unsafe.Pointer(v9))), 0))
		*memmap.PtrUint32(6112660, 1307740) = uint32(v7)
	}
	clientPlaySoundSpecial(920, 100)
	return 1
}
func sub_4A4B70(a1 int32) unsafe.Pointer {
	var (
		v1     uint8
		result unsafe.Pointer
		v3     *uint8
		v4     int32
		v5     int32
		v6     int32
		v7     int32
		v8     bool
		v9     uint8
		v10    int32
		v11    int32
	)
	v1 = 0
	result = unsafe.Pointer(uintptr(*(*uint8)(unsafe.Pointer(uintptr(dword_5d4594_1307724 + 66)))))
	v3 = *(**uint8)(memmap.PtrOff(0x587000, uint32(uintptr(result))*4+0x298AC))
	if int32(*v3) != 0 {
		for {
			result = unsafe.Pointer(uintptr(func() uint8 {
				p := &v1
				*p++
				return *p
			}()))
			if int32(*(*uint8)(unsafe.Add(unsafe.Pointer(v3), int32(v1)*4+int32(v1)))) == 0 {
				break
			}
		}
		if int32(v1) != 0 {
			v4 = 0
			v9 = uint8(int8(randomIntMinMax(0, int32(v1)-1)))
			if int32(*(*uint8)(unsafe.Pointer(uintptr(dword_5d4594_1307724 + 66)))) != 0 {
				v10 = 0
				v11 = 5
				for {
					nox_xxx_clientUpdateButtonRow_45E110(v10)
					v6 = 0
					v7 = 5
					for {
						if a1 == 1 {
							nox_xxx_book_45DBE0(unsafe.Pointer(uintptr(2)), 0, v6)
						} else {
							nox_xxx_book_45DBE0(unsafe.Pointer(uintptr(2)), int32(*(*uint8)(unsafe.Add(unsafe.Pointer(v3), int32(v9)*4+v6+int32(v9)))), v6)
						}
						v6++
						v7--
						if v7 == 0 {
							break
						}
					}
					v8 = v11 == 1
					v10++
					v11--
					if v8 {
						break
					}
				}
				result = unsafe.Pointer(uintptr(nox_xxx_clientUpdateButtonRow_45E110(0)))
			} else {
				nox_xxx_clientUpdateButtonRow_45E110(0)
				v5 = 5
				for {
					if a1 == 1 {
						result = nox_xxx_book_45DBE0(unsafe.Pointer(uintptr(3)), 0, v4)
					} else {
						result = nox_xxx_book_45DBE0(unsafe.Pointer(uintptr(3)), int32(*(*uint8)(unsafe.Add(unsafe.Pointer(v3), int32(v9)*4+v4+int32(v9)))), v4)
					}
					v4++
					v5--
					if v5 == 0 {
						break
					}
				}
			}
		}
	}
	return result
}
