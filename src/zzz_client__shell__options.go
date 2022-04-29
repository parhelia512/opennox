package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"unsafe"
)

func sub_4AA650() {
	var (
		v2 *byte
		v3 unsafe.Pointer
	)
	if sub_44D930() == 0 {
		v2 = *(**byte)(memmap.PtrOff(0x587000, func() uint32 {
			p := memmap.PtrUint32(6112660, 1309744)
			x := *p
			*p++
			return x
		}()*4+0x2A35C))
		nox_strman_loadString_40F1D0(v2, (**byte)(unsafe.Pointer(&v3)), CString("C:\\NoxPost\\src\\client\\shell\\Options.c"), 131)
		*memmap.PtrUint32(6112660, 1309744) %= 3
		if v3 != nil {
			nox_xxx_playDialogFile_44D900(int32(uintptr(v3)), 100)
		}
	}
}
