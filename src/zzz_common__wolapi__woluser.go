package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	"github.com/noxworld-dev/opennox/v1/common/alloc"
	"unsafe"
)

func sub_41F620(a1 int32) {
	var (
		v1 *uint32
		v2 int32
		v3 *uint32
		v4 int32
		v5 int32
		v6 int32
		v7 *libc.WChar
	)
	v1 = *(**uint32)(unsafe.Pointer(&dword_5d4594_531648))
	v2 = 0
	if dword_5d4594_531648 != 0 {
		for {
			v3 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*5)))))
			if libc.StrCaseCmp((*byte)(unsafe.Pointer(uintptr(*v1+36))), (*byte)(unsafe.Pointer(uintptr(a1+36)))) == 0 {
				break
			}
			v2++
			v1 = v3
			if v3 == nil {
				return
			}
		}
		v4 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*6)))
		if v4 != 0 {
			*(*uint32)(unsafe.Pointer(uintptr(v4 + 20))) = *(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*5))
		} else {
			dword_5d4594_531648 = *(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*5))
		}
		v5 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*5)))
		if v5 != 0 {
			*(*uint32)(unsafe.Pointer(uintptr(v5 + 24))) = *(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*6))
		} else {
			dword_5d4594_531652 = *(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*6))
		}
		if sub_41E2F0() == 7 {
			v6 = sub_446C70()
			sub_448620(v2)
			if v6 != -1 && sub_446C70() == -1 {
				clientPlaySoundSpecial(226, 100)
				v7 = strMan.GetStringInFile("lostselection", "C:\\NoxPost\\src\\common\\WolAPI\\woluser.c")
				sub_447310(0, int32(uintptr(unsafe.Pointer(v7))))
			}
		}
		alloc.Free(unsafe.Pointer(v1))
	}
}
