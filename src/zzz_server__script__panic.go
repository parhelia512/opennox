package opennox

import (
	"github.com/gotranspile/cxgo/runtime/stdio"
	"math"
	"unsafe"
)

var nox_script_panic_compiler_is_enabled bool = false

func nox_script_panic_compiler_check(fnc int32) bool {
	var (
		off       uint32 = uint32(fnc*4) + (0x587000 + 245900)
		stack_off int32  = int32(off - (6112660 + 0x3A51AC))
	)
	if stack_off < 0 || stack_off+4 > 4096 || stack_off%4 != 0 {
		return false
	}
	var stack_ind int32 = stack_off / 4
	if nox_script_stack_at(stack_ind) != 0x97974C {
		return false
	}
	stack_ind++
	var stack_exp [116]uint8 = [116]uint8{86, 190, 144, 16, 117, 0, 199, 6, 104, 80, 114, 80, 199, 70, 4, 0, 104, 48, 114, 199, 70, 8, 80, 0, math.MaxUint8, 84, 199, 70, 12, 36, 4, math.MaxUint8, 48, 199, 70, 16, math.MaxUint8, 84, 36, 4, 199, 70, 20, 131, 196, 12, 49, 199, 70, 24, 192, 195, 144, 144, 199, 70, 28, 86, 104, 80, 114, 199, 70, 32, 80, 0, math.MaxUint8, 20, 199, 70, 36, 36, 139, 240, math.MaxUint8, 199, 70, 40, 20, 36, 137, 48, 199, 70, 44, 131, 196, 4, 94, 199, 70, 48, 195, 144, 144, 144, 137, 53, 112, 51, 92, 0, 199, 5, 240, 49, 92, 0, 172, 16, 117, 0, 94, 195, 144, 144}
	var body_len int32 = int32(unsafe.Sizeof([116]uint8{}) / 4)
	if stack_ind+body_len >= 1024 {
		return false
	}
	var stack_dword *uint32 = (*uint32)(unsafe.Pointer(&stack_exp[0]))
	for i := int32(0); i < body_len; i++ {
		if *(*uint32)(unsafe.Add(unsafe.Pointer(stack_dword), unsafe.Sizeof(uint32(0))*uintptr(i))) != nox_script_stack_at(stack_ind+i) {
			return false
		}
	}
	nox_script_panic_compiler_is_enabled = true
	stdio.Printf("noxscript: enabled Panic's compiler API\n")
	return true
}
func nox_script_panic_compiler_sub_751090() int32 {
	var addr uint32 = uint32(nox_script_pop())
	stdio.Printf("noxscript: mem read [0x%x]\n", addr)
	var val uint32 = 0
	nox_script_push(int32(val))
	return 0
}
func nox_script_panic_compiler_write_7510AC() int32 {
	var (
		val  uint32 = uint32(nox_script_pop())
		addr uint32 = uint32(nox_script_pop())
	)
	stdio.Printf("noxscript: mem write [0x%x] = %d\n", addr, val)
	return 0
}
func nox_script_panic_compiler_call(fi int32, ret *int32) bool {
	if !nox_script_panic_compiler_is_enabled {
		return false
	}
	if fi == 185 {
		*ret = nox_script_panic_compiler_sub_751090()
		return true
	} else if fi == 89 {
		*ret = nox_script_panic_compiler_write_7510AC()
		return true
	}
	return false
}
