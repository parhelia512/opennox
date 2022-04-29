package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"unsafe"
)

func sub_4CBD30() *byte {
	var (
		v0     int32
		i      int32
		v2     *libc.WChar
		v3     int32
		v4     *libc.WChar
		v5     *byte
		v6     *libc.WChar
		v7     *byte
		v8     *libc.WChar
		v9     int32
		v10    *libc.WChar
		result *byte
		v12    [256]byte
	)
	v0 = int32(*(*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer(dword_5d4594_1522620))) + 32))))
	sub_42CD90()
	for i = 0; i < int32(*(*int16)(unsafe.Pointer(uintptr(v0 + 44)))); i++ {
		v2 = (*libc.WChar)(unsafe.Pointer(uintptr((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1522620))))).Func94(asWindowEvent(0x4016, i, 0)))))
		v3 = int32(uintptr(unsafe.Pointer(nox_xxx_bindevent_bindNameByTitle_42EA40(v2))))
		v4 = (*libc.WChar)(unsafe.Pointer(uintptr((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1522624))))).Func94(asWindowEvent(0x4016, i, 0)))))
		v5 = nox_xxx_keybind_nameByTitle_42E960(v4)
		v6 = (*libc.WChar)(unsafe.Pointer(uintptr((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1522628))))).Func94(asWindowEvent(0x4016, i, 0)))))
		v7 = nox_xxx_keybind_nameByTitle_42E960(v6)
		if v7 != nil {
			nox_sprintf(&v12[0], CString("%s = %s"), v7, v3)
			nox_client_parseConfigHotkeysLine_42CF50(&v12[0])
		}
		if v5 != nil {
			nox_sprintf(&v12[0], CString("%s = %s"), v5, v3)
			nox_client_parseConfigHotkeysLine_42CF50(&v12[0])
		}
	}
	v8 = strMan.GetStringInFile("bindevent:ToggleQuitMenu", "C:\\NoxPost\\src\\Client\\shell\\InputCfg\\inputcfg.c")
	v9 = int32(uintptr(unsafe.Pointer(nox_xxx_bindevent_bindNameByTitle_42EA40(v8))))
	v10 = strMan.GetStringInFile("keybind:Esc", "C:\\NoxPost\\src\\Client\\shell\\InputCfg\\inputcfg.c")
	result = nox_xxx_keybind_nameByTitle_42E960(v10)
	if result != nil {
		nox_sprintf(&v12[0], CString("%s = %s"), result, v9)
		result = (*byte)(unsafe.Pointer(uintptr(nox_client_parseConfigHotkeysLine_42CF50(&v12[0]))))
	}
	return result
}
func sub_4CBF60(a1 int32, a2 uint32, a3 int32, a4 int32) int32 {
	var (
		v5 int32
		v6 *libc.WChar
		v7 int32
		v8 int32
		v9 int32
	)
	if a2 > 0x4007 {
		if a2 == 0x4009 {
			v7 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 32))))
			nox_xxx_wndListboxProcPre_4A30D0((*nox_window)(unsafe.Pointer(uintptr(a1))), 0x4009, uint32(uintptr(unsafe.Pointer((*libc.WChar)(unsafe.Pointer(uintptr(a3)))))), a4)
			v8 = sub_4A4800(v7)
			(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1522620))))).Func94(asWindowEvent(0x401C, v8, 0))
			(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1522624))))).Func94(asWindowEvent(0x401C, v8, 0))
			(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1522628))))).Func94(asWindowEvent(0x401C, v8, 0))
		} else if a2 == 16400 {
			v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(a3 + 32))))
			if int32(*(*uint32)(unsafe.Pointer(uintptr(v5 + 48)))) >= 0 {
				dword_5d4594_1522632 = uint32(a3)
				v9 = (*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1522620))))).Func94(asWindowEvent(0x4016, int32(*(*uint32)(unsafe.Pointer(uintptr(v5 + 48)))), 0))
				v6 = strMan.GetStringInFile("InputCfg.wnd:PressKey", "C:\\NoxPost\\src\\Client\\shell\\InputCfg\\inputcfg.c")
				nox_swprintf((*libc.WChar)(memmap.PtrOff(6112660, 1522636)), libc.CWString("%s\n'%s'"), v6, v9)
				nox_xxx_wndShowModalMB((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1522612))))))
				guiFocus((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1522612))))))
				sub_46C690((*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1522612))))))
				return nox_xxx_wndListboxProcPre_4A30D0((*nox_window)(unsafe.Pointer(uintptr(a1))), 16400, uint32(uintptr(unsafe.Pointer((*libc.WChar)(unsafe.Pointer(uintptr(a3)))))), a4)
			}
		}
	} else {
		if a2 != 0x4007 {
			if a2 == 23 {
				return 1
			}
			if a2 != 0x4000 {
				return nox_xxx_wndListboxProcPre_4A30D0((*nox_window)(unsafe.Pointer(uintptr(a1))), a2, uint32(uintptr(unsafe.Pointer((*libc.WChar)(unsafe.Pointer(uintptr(a3)))))), a4)
			}
		}
		if unsafe.Pointer((*uint32)(unsafe.Pointer(uintptr(a3)))) == unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1522604)))).ChildByID(921)) || unsafe.Pointer((*uint32)(unsafe.Pointer(uintptr(a3)))) == unsafe.Pointer((*nox_window)(unsafe.Pointer(*(**uint32)(unsafe.Pointer(&dword_5d4594_1522604)))).ChildByID(922)) {
			(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1522620))))).Func94(asWindowEvent(int32(a2), a3, 0))
			(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1522624))))).Func94(asWindowEvent(int32(a2), a3, 0))
			(*nox_window)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&dword_5d4594_1522628))))).Func94(asWindowEvent(int32(a2), a3, 0))
			return nox_xxx_wndListboxProcPre_4A30D0((*nox_window)(unsafe.Pointer(uintptr(a1))), a2, uint32(uintptr(unsafe.Pointer((*libc.WChar)(unsafe.Pointer(uintptr(a3)))))), a4)
		}
	}
	return nox_xxx_wndListboxProcPre_4A30D0((*nox_window)(unsafe.Pointer(uintptr(a1))), a2, uint32(uintptr(unsafe.Pointer((*libc.WChar)(unsafe.Pointer(uintptr(a3)))))), a4)
}
