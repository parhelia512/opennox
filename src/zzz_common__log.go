package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	"github.com/gotranspile/cxgo/runtime/stdio"
	noxflags "github.com/noxworld-dev/opennox/v1/common/flags"
)

var nox_log_buf [512]byte = [512]byte{}
var nox_file_log_4 *File = nil
var nox_file_net_log *File = nil

func nox_asctime_413B00() *byte {
	var t libc.Time
	libc.GetTime(&t)
	return libc.AscTime(libc.LocalTime(&t))
}
func nox_xxx_log_open_413B20(f **File, path *byte) int32 {
	if path == nil {
		return 0
	}
	*f = nox_fs_create_rw(path)
	if f == nil {
		return 0
	}
	nox_fs_fprintf(*f, CString("Log opened: %s"), nox_asctime_413B00())
	nox_fs_flush(*f)
	return 1
}
func nox_xxx_log_close_413AD0(f *File) {
	if f != nil {
		nox_fs_fprintf(f, CString("Log closed: %s"), nox_asctime_413B00())
		nox_fs_flush(f)
		nox_fs_close(f)
	}
}
func nox_xxx_log_4_reopen_413A80(path *byte) int32 {
	if !nox_common_getEngineFlag(nox_engine_flag(NOX_ENGINE_FLAG_LOG_TO_FILE)) {
		return 0
	}
	nox_xxx_log_close_413AD0(nox_file_log_4)
	nox_file_log_4 = nil
	nox_xxx_log_open_413B20(&nox_file_log_4, path)
	return 1
}
func nox_xxx_log_4_printf_413B70(fmt *byte, _rest ...interface{}) {
	var va libc.ArgList
	va.Start(fmt, _rest)
	nox_vsprintf(&nox_log_buf[0], fmt, va)
	if nox_common_getEngineFlag(nox_engine_flag(NOX_ENGINE_FLAG_LOG_TO_FILE)) {
		nox_fs_fputs_sync(nox_file_log_4, &nox_log_buf[0])
	}
	if nox_common_getEngineFlag(nox_engine_flag(NOX_ENGINE_FLAG_LOG_TO_CONSOLE)) {
		nox_gui_console_Printf_450C00(uint8(int8(NOX_CONSOLE_GREEN)), libc.CWString("%S"), &nox_log_buf[0])
	}
}
func nox_xxx_log_4_close_413C00() {
	noxflags.UnsetEngine(nox_engine_flag(NOX_ENGINE_FLAG_LOG_TO_FILE | NOX_ENGINE_FLAG_LOG_TO_CONSOLE))
	if nox_file_log_4 != nil {
		nox_xxx_log_close_413AD0(nox_file_log_4)
		nox_file_log_4 = nil
	}
}
func nox_xxx_networkLog_printf_413D30(fmt *byte, _rest ...interface{}) {
	if !noxflags.HasGame(noxflags.GameFlag3) {
		return
	}
	var va libc.ArgList
	va.Start(fmt, _rest)
	var v6 libc.Time
	libc.GetTime(&v6)
	var v1 *libc.TimeInfo = libc.LocalTime(&v6)
	nox_vsprintf(&nox_log_buf[0], fmt, va)
	nox_sprintf(&nox_log_buf[0], CString("%s%c("), &nox_log_buf[0], 240)
	libc.StrCat(&nox_log_buf[0], libc.AscTime(v1))
	nox_log_buf[libc.StrLen(&nox_log_buf[0])-1] = 0
	libc.StrCat(&nox_log_buf[0], CString(")\n"))
	if nox_file_net_log != nil {
		nox_fs_fputs_sync(nox_file_net_log, &nox_log_buf[0])
	}
	nox_gui_console_Printf_450C00(uint8(int8(NOX_CONSOLE_GREEN)), libc.CWString("%S"), &nox_log_buf[0])
}
func nox_xxx_networkLog_init_413CC0() {
	nox_file_net_log = nox_fs_append_text(CString("network.log"))
	if nox_file_net_log != nil {
		nox_xxx_networkLog_printf_413D30(CString("StartLog%c%s"), 240, "1.0")
	}
}
func nox_xxx_networkLog_close_413D00() {
	nox_xxx_networkLog_printf_413D30(CString("EndLog"))
	nox_fs_close(nox_file_net_log)
	nox_file_net_log = nil
}
