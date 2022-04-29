package opennox

import "github.com/gotranspile/cxgo/runtime/libc"

func nox_gui_parseWindowTooltip_4A0800(data *nox_window_data, buf *byte) int32 {
	var str *libc.WChar = strMan.GetStringInFile(buf, "C:\\NoxPost\\src\\Client\\Gui\\GameWin\\psscript.c")
	nox_xxx_wndWddSetTooltip_46B000(data, str)
	return 1
}
func nox_gui_parseWindowText_4A0A90(data *nox_window_data, buf *byte) int32 {
	var str *libc.WChar
	if libc.StrCmp(buf, CString("Options.wnd:8BitColor")) == 0 {
		str = libc.CWString("\tWindowed")
	} else if libc.StrCmp(buf, CString("Options.wnd:16BitColor")) == 0 {
		str = libc.CWString("\tFullscreen")
	} else {
		str = strMan.GetStringInFile(buf, "C:\\NoxPost\\src\\Client\\Gui\\GameWin\\psscript.c")
	}
	nox_wcsncpy(&data.text[0], str, 63)
	data.text[63] = 0
	return 1
}
