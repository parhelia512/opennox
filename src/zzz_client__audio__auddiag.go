package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"unsafe"
)

func nox_xxx_WorkerHurt_44D810() int32 {
	var v1 int32
	if dword_5d4594_831076 == 0 {
		dword_5d4594_831092 = uint32(sub_43F130())
		dword_587000_122848 = uint32(bool2int(dword_5d4594_831092 != 0))
		sub_4864A0((*uint32)(memmap.PtrOff(6112660, 830876)))
		sub_486380((*uint32)(memmap.PtrOff(6112660, 830876)), 500, 0, 0x4000)
		dword_5d4594_830864 = 0
		dword_5d4594_830972 = 0
		dword_5d4594_830872 = 0
		*memmap.PtrUint32(6112660, 831080) = 0
		dword_5d4594_831084 = 0
		dword_5d4594_831076 = 1
		nox_strman_loadString_40F1D0(CString("Con03B.scr:WorkerHurt"), (**byte)(unsafe.Pointer(&v1)), CString("C:\\NoxPost\\src\\client\\Audio\\AudDiag.c"), 279)
		if v1 != 0 {
			nox_xxx_playDialogFile_44D900(v1, 0)
		}
	}
	return 1
}
