package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	"github.com/gotranspile/cxgo/runtime/stdio"
	"github.com/noxworld-dev/opennox/v1/common/alloc"
	noxflags "github.com/noxworld-dev/opennox/v1/common/flags"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"math"
	"unsafe"
)

var nox_script_caller_3821964 unsafe.Pointer = nil
var nox_script_trigger_3821968 unsafe.Pointer = nil
var nox_script_arr_xxx_1599636 *nox_script_xxx_t = nil
var nox_script_count_xxx_1599640 int32 = 0
var nox_script_strings [1024]*byte = [1024]*byte{}
var nox_script_strings_xxx uint32 = 0
var nox_script_strings_cnt uint32 = 0
var nox_script_stack [1024]int32 = [1024]int32{}
var nox_script_stack_top int32 = 0

func nox_script_stack_at(i int32) uint32 {
	if i < 0 || i >= 1024 {
		return 0
	}
	return uint32(nox_script_stack[i])
}
func nox_script_push(val int32) {
	if nox_script_stack_top < 1024 {
		nox_script_stack[nox_script_stack_top] = val
		nox_script_stack_top++
	}
}
func nox_script_pushf(val float32) {
	nox_script_push(*((*int32)(unsafe.Pointer(&val))))
}
func nox_script_pop() int32 {
	var i int32 = nox_script_stack_top
	if nox_script_stack_top > 0 {
		i = func() int32 {
			p := &nox_script_stack_top
			*p--
			return *p
		}()
	}
	return nox_script_stack[i]
}
func nox_script_popf() float32 {
	var v int32 = nox_script_pop()
	return *((*float32)(unsafe.Pointer(&v)))
}
func nox_script_nextInt(p **int32) int32 {
	var v int32 = **p
	*p = (*int32)(unsafe.Add(unsafe.Pointer(*p), unsafe.Sizeof(int32(0))*1))
	return v
}
func nox_script_nextFloat(p **float32) float64 {
	var v float64 = float64(**p)
	*p = (*float32)(unsafe.Add(unsafe.Pointer(*p), unsafe.Sizeof(float32(0))*1))
	return v
}
func nox_script_indexByEvent(a1 *byte) int32 {
	if nox_script_count_xxx_1599640 <= 0 {
		return -1
	}
	for i := int32(0); i < nox_script_count_xxx_1599640; i++ {
		if libc.StrCmp((*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*uintptr(i)))).field_0, a1) == 0 {
			return i
		}
	}
	return -1
}
func nox_script_get_caller() unsafe.Pointer {
	return nox_script_caller_3821964
}
func nox_script_get_trigger() unsafe.Pointer {
	return nox_script_trigger_3821968
}
func nox_script_addString_512E40(s *byte) int32 {
	if nox_script_strings_cnt >= 1024 {
		return int32(nox_script_strings_cnt - 1)
	}
	var cstr *byte = libc.StrDup(s)
	var i int32 = int32(nox_script_strings_cnt)
	nox_script_strings[i] = cstr
	nox_script_strings_cnt++
	return i
}
func nox_script_getString_512E40(i int32) *byte {
	if i < 0 || uint32(i) >= nox_script_strings_cnt {
		return nil
	}
	return nox_script_strings[i]
}
func nox_script_getField36(i int32) *byte {
	return (*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*uintptr(i)))).field_36
}
func nox_script_getField40(i int32) uint32 {
	return (*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*uintptr(i)))).field_40
}
func nox_script_getField44(i int32) uint32 {
	return (*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*uintptr(i)))).field_44
}
func nox_xxx_xferReadScriptHandler_4F5580(a1 int32, a2 *byte) int32 {
	var (
		v3 bool
		v4 int32
		v5 uint32
		v6 int32
		v7 [1024]byte
	)
	v6 = 1
	cryptFileReadWrite((*uint8)(unsafe.Pointer(&v6))[:2])
	if int32(int16(v6)) > 1 {
		return 0
	}
	if nox_xxx_cryptGetXxx() == 1 {
		cryptFileReadWrite((*uint8)(unsafe.Pointer(&v5))[:4])
		if v5 >= 1024 {
			return 0
		}
		cryptFileReadWrite((*uint8)(unsafe.Pointer(&v7[0]))[:v5])
		v3 = v5 == 0
		v7[v5] = 0
		if !v3 {
			if noxflags.HasGame(noxflags.GameFlag22 | noxflags.GameFlag23) {
				libc.StrCpy(a2, &v7[0])
			} else {
				*(*uint32)(unsafe.Pointer(uintptr(a1 + 4))) = uint32(nox_script_indexByEvent(&v7[0]))
			}
		}
	} else {
		if noxflags.HasGame(noxflags.GameFlag22 | noxflags.GameFlag23) {
			if a2 != nil {
				v5 = uint32(libc.StrLen(a2))
				cryptFileReadWrite((*uint8)(unsafe.Pointer(&v5))[:4])
				cryptFileReadWrite((*uint8)(unsafe.Pointer(a2))[:v5])
				goto LABEL_16
			}
		} else {
			v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 4))))
			if v4 != -1 {
				v5 = uint32(libc.StrLen((*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*uintptr(v4)))).field_0))
				cryptFileReadWrite((*uint8)(unsafe.Pointer(&v5))[:4])
				cryptFileReadWrite((*uint8)(unsafe.Pointer((*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 4))))))).field_0))[:v5])
				goto LABEL_16
			}
		}
		v5 = 0
		cryptFileReadWrite((*uint8)(unsafe.Pointer(&v5))[:4])
	}
LABEL_16:
	cryptFileReadWrite((*uint8)(unsafe.Pointer(uintptr(a1)))[:4])
	return 1
}
func nox_xxx_scriptCallByEventBlock_502490(a1 *int32, a2 int32, a3 int32, eventCode int32) *uint8 {
	nox_script_callByEvent_cgo(eventCode, unsafe.Pointer(uintptr(a2)), unsafe.Pointer(uintptr(a3)))
	*memmap.PtrUint32(6112660, 1599076) = 0
	var v3 int32 = *(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*0))
	if *a1&1 != 0 {
		return nil
	}
	var v4 int32 = *(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*1))
	if v4 == -1 {
		return nil
	}
	if v3&2 != 0 {
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v3))), 0)) = uint8(int8(v3 | 1))
		*(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*0)) = v3
	}
	if nox_script_stack_top != 0 {
		sub_5025A0(int32(uintptr(unsafe.Pointer(a1))), a2, a3)
		return (*uint8)(memmap.PtrOff(6112660, 1599076))
	}
	nox_script_callByIndex_507310(v4, unsafe.Pointer(uintptr(a2)), unsafe.Pointer(uintptr(a3)))
	if (*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(int32(0))*1)))))).stack_size != 0 {
		*memmap.PtrUint32(6112660, 1599076) = uint32(nox_script_pop())
	}
	nox_script_stack_top = 0
	if nox_script_strings_xxx < nox_script_strings_cnt {
		for i := int32(int32(nox_script_strings_xxx)); uint32(i) < nox_script_strings_cnt; i++ {
			alloc.Free(unsafe.Pointer(nox_script_strings[i]))
			nox_script_strings[i] = nil
		}
	}
	nox_script_strings_cnt = nox_script_strings_xxx
	sub_5025E0(int32(uintptr(unsafe.Pointer(a1))), a2, a3)
	if *memmap.PtrUint32(6112660, 1599468) > 0 {
		nox_xxx_scriptCallByEventBlock_502490(*(**int32)(memmap.PtrOff(6112660, 1599084)), *memmap.PtrInt32(6112660, 1599088), *memmap.PtrInt32(6112660, 1599092), 0)
	}
	return (*uint8)(memmap.PtrOff(6112660, 1599076))
}
func nox_server_mapRWScriptData_504F90() int32 {
	var (
		result int32
		v1     int8
		v2     int32
	)
	v2 = 1
	cryptFileReadWrite((*uint8)(unsafe.Pointer(&v2))[:2])
	if int32(int16(v2)) > 1 {
		return 0
	}
	v1 = 0
	if nox_script_arr_xxx_1599636 != nil && noxflags.HasGame(noxflags.GameHost) && !noxflags.HasGame(noxflags.GameFlag23) {
		v1 = 1
	}
	cryptFileReadWrite((*uint8)(unsafe.Pointer(&v1))[:1])
	if int32(v1) == 0 {
		return 1
	}
	cryptFileReadWrite((*uint8)(unsafe.Pointer((*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*1))).field_28))[:(*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*1))).field_16*4])
	if nox_xxx_cryptGetXxx() != 0 {
		result = nox_script_activatorLoad_51AF80()
		if result == 0 {
			return result
		}
		return 1
	}
	result = nox_script_activatorSave_51AEA0()
	if result != 0 {
		return 1
	}
	return result
}
func nox_script_ncobj_readInt_505800(f *File) int32 {
	var (
		val int32
		n   int32 = nox_fs_fread(f, unsafe.Pointer(&val), 4)
	)
	if n == 4 {
		return val
	}
	return 0
}
func nox_script_ncobj_readString_505830(f *File, sz int32, dst *byte) bool {
	var n int32 = nox_fs_fread(f, unsafe.Pointer(dst), sz)
	*(*byte)(unsafe.Add(unsafe.Pointer(dst), sz)) = 0
	return n == sz
}
func nox_script_ncobj_readStringExpect_505870(f *File, exp *byte) bool {
	var buf [256]byte
	nox_script_ncobj_readString_505830(f, int32(libc.StrLen(exp)), &buf[0])
	return libc.StrCmp(&buf[0], exp) == 0
}
func nox_script_ncobj_parse_505360() int32 {
	var f *File = nox_fs_open(CString("nc.obj"))
	if f == nil {
		return 0
	}
	if !nox_script_ncobj_readStringExpect_505870(f, CString("SCRIPT03")) {
		nox_fs_close(f)
		return 0
	}
	if !nox_script_ncobj_readStringExpect_505870(f, CString("STRG")) {
		nox_fs_close(f)
		return 0
	}
	nox_script_strings_xxx = uint32(nox_script_ncobj_readInt_505800(f))
	nox_script_strings_cnt = nox_script_strings_xxx
	if nox_script_strings_xxx > 0 {
		for i := int32(0); uint32(i) < nox_script_strings_xxx; i++ {
			var (
				n   int32 = nox_script_ncobj_readInt_505800(f)
				str *byte = (*byte)(alloc.Calloc(1, int(n+1)))
			)
			nox_script_strings[i] = str
			if !nox_script_ncobj_readString_505830(f, n, str) {
				nox_fs_close(f)
				return 0
			}
		}
	}
	if !nox_script_ncobj_readStringExpect_505870(f, CString("CODE")) {
		nox_fs_close(f)
		return 0
	}
	nox_script_count_xxx_1599640 = nox_script_ncobj_readInt_505800(f)
	if nox_script_count_xxx_1599640 < 0 {
		nox_script_count_xxx_1599640 = 0
	}
	if nox_script_count_xxx_1599640 != 0 {
		nox_script_arr_xxx_1599636 = (*nox_script_xxx_t)(alloc.Calloc(1, int(nox_script_count_xxx_1599640*int32(unsafe.Sizeof(nox_script_xxx_t{})))))
	}
	var buf [1024]byte
	for ind := int32(0); ind < nox_script_count_xxx_1599640; ind++ {
		if !nox_script_ncobj_readStringExpect_505870(f, CString("FUNC")) {
			nox_fs_close(f)
			return 0
		}
		var cur *nox_script_xxx_t = (*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*uintptr(ind)))
		var n int32 = nox_script_ncobj_readInt_505800(f)
		cur.field_0 = (*byte)(alloc.Calloc(1, int(n+1)))
		if !nox_script_ncobj_readString_505830(f, n, cur.field_0) {
			nox_fs_close(f)
			return 0
		} else if libc.StrLen(cur.field_0) >= 1024 {
			return 0
		}
		libc.StrCpy(&buf[0], cur.field_0)
		var v14 *uint8 = (*uint8)(unsafe.Pointer(libc.StrTok(&buf[0], CString("%"))))
		var v16 *uint8 = (*uint8)(unsafe.Pointer(cur.field_0))
		var v19 int32 = 0
		for {
			var (
				v15 uint8 = *(*uint8)(unsafe.Add(unsafe.Pointer(v14), 0))
				v17 bool  = int32(v15) < int32(*(*uint8)(unsafe.Add(unsafe.Pointer(v16), 0)))
			)
			if int32(v15) != int32(*(*uint8)(unsafe.Add(unsafe.Pointer(v16), 0))) {
				v19 = -bool2int(v17) - (bool2int(v17) - 1)
				break
			}
			if int32(v15) == 0 {
				v19 = 0
				break
			}
			var v18 uint8 = *(*uint8)(unsafe.Add(unsafe.Pointer(v14), 1))
			v17 = int32(v18) < int32(*(*uint8)(unsafe.Add(unsafe.Pointer(v16), 1)))
			if int32(v18) != int32(*(*uint8)(unsafe.Add(unsafe.Pointer(v16), 1))) {
				v19 = -bool2int(v17) - (bool2int(v17) - 1)
				break
			}
			if int32(v18) == 0 {
				v19 = 0
				break
			}
			v14 = (*uint8)(unsafe.Add(unsafe.Pointer(v14), 2))
			v16 = (*uint8)(unsafe.Add(unsafe.Pointer(v16), 2))
		}
		if v19 != 0 {
			nox_sprintf(&buf[0], CString("%%%s"), libc.StrTok(nil, CString("%")))
			cur.field_36 = (*byte)(alloc.Calloc(1, libc.StrLen(&buf[0])+1))
			libc.StrCpy(cur.field_36, &buf[0])
			cur.field_40 = uint32(libc.Atoi(libc.GoString(libc.StrTok(nil, CString("%")))))
			cur.field_44 = uint32(libc.Atoi(libc.GoString(libc.StrTok(nil, CString("%")))))
		} else {
			cur.field_36 = nil
			cur.field_40 = 0
			cur.field_44 = 0
		}
		cur.stack_size = uint32(nox_script_ncobj_readInt_505800(f))
		cur.size_28 = uint32(nox_script_ncobj_readInt_505800(f))
		if !nox_script_ncobj_readStringExpect_505870(f, CString("SYMB")) {
			nox_fs_close(f)
			return 0
		}
		var cntY int32 = nox_script_ncobj_readInt_505800(f)
		if ind == 0 {
			cntY++
		}
		cur.field_12 = uint32(cntY)
		nox_script_ncobj_readInt_505800(f)
		if cntY != 0 {
			cur.field_20 = (*uint32)(alloc.Calloc(1, int(cntY*int32(unsafe.Sizeof(uint32(0))))))
			cur.field_24 = (*uint32)(alloc.Calloc(1, int(cntY*int32(unsafe.Sizeof(uint32(0))))))
		} else {
			cur.field_20 = nil
			cur.field_24 = nil
		}
		var j1 int32 = 0
		if ind == 0 {
			j1 = 1
			*(*uint32)(unsafe.Add(unsafe.Pointer((*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*0))).field_20), unsafe.Sizeof(uint32(0))*0)) = 0
			*(*uint32)(unsafe.Add(unsafe.Pointer((*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*0))).field_24), unsafe.Sizeof(uint32(0))*0)) = 0
		}
		var sum int32 = 0
		for j := int32(j1); j < cntY; j++ {
			var v int32 = nox_script_ncobj_readInt_505800(f)
			*(*uint32)(unsafe.Add(unsafe.Pointer(cur.field_20), unsafe.Sizeof(uint32(0))*uintptr(j))) = uint32(v)
			*(*uint32)(unsafe.Add(unsafe.Pointer(cur.field_24), unsafe.Sizeof(uint32(0))*uintptr(j))) = uint32(sum)
			sum += v
		}
		cur.field_16 = uint32(sum)
		if sum != 0 {
			cur.field_28 = (*uint32)(alloc.Calloc(1, int(sum*4)))
		} else {
			cur.field_28 = nil
		}
		if !nox_script_ncobj_readStringExpect_505870(f, CString("DATA")) {
			nox_fs_close(f)
			return 0
		}
		n = nox_script_ncobj_readInt_505800(f)
		cur.data = alloc.Calloc(1, int(n))
		if nox_fs_fread(f, cur.data, n) != n {
			nox_fs_close(f)
			return 0
		}
	}
	if !nox_script_ncobj_readStringExpect_505870(f, CString("DONE")) {
		nox_fs_close(f)
		return 0
	}
	nox_fs_close(f)
	return 1
}
func nox_script_freeEverything_5058F0() {
	if nox_script_strings_xxx < nox_script_strings_cnt {
		for i := int32(int32(nox_script_strings_xxx)); uint32(i) < nox_script_strings_cnt; i++ {
			alloc.Free(unsafe.Pointer(nox_script_strings[i]))
			nox_script_strings[i] = nil
		}
	}
	if nox_script_strings_xxx > 0 {
		for i := int32(0); uint32(i) < nox_script_strings_xxx; i++ {
			alloc.Free(unsafe.Pointer(nox_script_strings[i]))
			nox_script_strings[i] = nil
		}
	}
	nox_script_strings_xxx = 0
	dword_5d4594_1599628 = 0
	nox_script_strings_cnt = 0
	if nox_script_arr_xxx_1599636 != nil {
		for i := int32(0); i < nox_script_count_xxx_1599640; i++ {
			if (*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*uintptr(i)))).field_0 != nil {
				alloc.Free(unsafe.Pointer((*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*uintptr(i)))).field_0))
			}
			if (*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*uintptr(i)))).field_36 != nil {
				alloc.Free(unsafe.Pointer((*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*uintptr(i)))).field_36))
			}
			if (*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*uintptr(i)))).field_20 != nil {
				alloc.Free(unsafe.Pointer((*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*uintptr(i)))).field_20))
			}
			if (*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*uintptr(i)))).field_24 != nil {
				alloc.Free(unsafe.Pointer((*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*uintptr(i)))).field_24))
			}
			if (*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*uintptr(i)))).field_28 != nil {
				alloc.Free(unsafe.Pointer((*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*uintptr(i)))).field_28))
			}
			if (*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*uintptr(i)))).data != nil {
				(*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*uintptr(i)))).data = nil
			}
		}
		alloc.Free(unsafe.Pointer(nox_script_arr_xxx_1599636))
		nox_script_arr_xxx_1599636 = nil
	}
	nox_script_count_xxx_1599640 = 0
}
func nox_xxx_scriptRunFirst_507290() {
	nox_script_resetBuiltin()
	if nox_script_arr_xxx_1599636 != nil {
		var f28 *byte = (*byte)(unsafe.Pointer((*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*1))).field_28))
		_ = f28
		*(*uint32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(f28), 0)))) = 0xFFFFFFFE
		*(*uint32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(f28), 4)))) = math.MaxUint32
		*(*uint32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(f28), 8)))) = 1
		*(*uint32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(f28), 12)))) = 0
		if nox_xxx_gameIsSwitchToSolo_4DB240() == 0 {
			nox_script_callByIndex_507310(1, nil, nil)
		}
	}
}
func nox_script_callByIndex_507310(index int32, a2 unsafe.Pointer, a3 unsafe.Pointer) {
	var (
		buf    [256]byte
		script *nox_script_xxx_t = (*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*uintptr(index)))
	)
	nox_script_caller_3821964 = a2
	nox_script_trigger_3821968 = a3
	for i := int32(0); uint32(i) < script.size_28; i++ {
		var v int32 = nox_script_pop()
		*(*uint32)(unsafe.Add(unsafe.Pointer(script.field_28), unsafe.Sizeof(uint32(0))*uintptr(i))) = uint32(v)
	}
	var bstack int32 = nox_script_stack_top
	var data unsafe.Pointer = script.data
	for {
		var (
			sa1 int32   = 0
			sa2 int32   = 0
			sa3 int32   = 0
			sa4 int32   = 0
			sa5 int32   = 0
			sf1 float32 = 0
			sf2 float32 = 0
		)
		switch nox_script_nextInt((**int32)(unsafe.Pointer(&data))) {
		case 0:
			fallthrough
		case 3:
			sa1 = nox_script_nextInt((**int32)(unsafe.Pointer(&data)))
			sa2 = nox_script_nextInt((**int32)(unsafe.Pointer(&data)))
			if sa1 != 0 {
				if (*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*1))).field_24 != nil && (*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*1))).field_28 != nil {
					sa2 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer((*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*1))).field_24), unsafe.Sizeof(uint32(0))*uintptr(sa2))))
					sa2 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer((*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*1))).field_28), unsafe.Sizeof(uint32(0))*uintptr(sa2))))
				} else {
					sa2 = 0
				}
				nox_script_push(sa2)
			} else if sa2 < 0 {
				sa2 = -sa2
				sa2 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer((*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*0))).field_24), unsafe.Sizeof(uint32(0))*uintptr(sa2))))
				sa2 = int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(int32(uintptr(a3)) + 760))) + uint32(sa2*4)))))
				nox_script_push(sa2)
			} else if script.field_24 != nil && script.field_28 != nil {
				sa2 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(script.field_24), unsafe.Sizeof(uint32(0))*uintptr(sa2))))
				sa2 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(script.field_28), unsafe.Sizeof(uint32(0))*uintptr(sa2))))
				nox_script_push(sa2)
			} else {
				nox_script_push(0)
			}
			continue
		case 1:
			sa1 = nox_script_nextInt((**int32)(unsafe.Pointer(&data)))
			sa2 = nox_script_nextInt((**int32)(unsafe.Pointer(&data)))
			if sa1 != 0 {
				if (*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*1))).field_24 != nil && (*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*1))).field_28 != nil {
					sa2 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer((*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*1))).field_24), unsafe.Sizeof(uint32(0))*uintptr(sa2))))
					sf1 = *(*float32)(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer((*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*1))).field_28), unsafe.Sizeof(uint32(0))*uintptr(sa2)))))
				} else {
					sf1 = 0
				}
			} else if sa2 < 0 {
				sa2 = -sa2
				sa2 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer((*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*0))).field_24), unsafe.Sizeof(uint32(0))*uintptr(sa2))))
				sf1 = *(*float32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(int32(uintptr(a3)) + 760))) + uint32(sa2*4))))
			} else if script.field_24 != nil && script.field_28 != nil {
				sa2 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(script.field_24), unsafe.Sizeof(uint32(0))*uintptr(sa2))))
				sf1 = *(*float32)(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer(script.field_28))) + uint32(sa2*4))))
			} else {
				sf1 = 0
			}
			nox_script_pushf(sf1)
			continue
		case 2:
			sa1 = nox_script_nextInt((**int32)(unsafe.Pointer(&data)))
			sa2 = nox_script_nextInt((**int32)(unsafe.Pointer(&data)))
			if sa1 != 0 {
				sa3 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer((*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*1))).field_20), unsafe.Sizeof(uint32(0))*uintptr(sa2))))
				sa4 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer((*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*1))).field_24), unsafe.Sizeof(uint32(0))*uintptr(sa2))))
			} else if sa2 < 0 {
				sa2 = -sa2
				sa3 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer((*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*0))).field_20), unsafe.Sizeof(uint32(0))*uintptr(sa2))))
				sa4 = int32(-*(*uint32)(unsafe.Add(unsafe.Pointer((*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*0))).field_24), unsafe.Sizeof(uint32(0))*uintptr(sa2))))
			} else if script.field_20 != nil && script.field_24 != nil {
				sa3 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(script.field_20), unsafe.Sizeof(uint32(0))*uintptr(sa2))))
				sa4 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(script.field_24), unsafe.Sizeof(uint32(0))*uintptr(sa2))))
			} else {
				sa3 = 0
				sa4 = 0
			}
			if sa3 > 1 {
				nox_script_push(sa3)
			}
			nox_script_push(sa1)
			nox_script_push(sa4)
			continue
		case 4:
			fallthrough
		case 6:
			sa1 = nox_script_nextInt((**int32)(unsafe.Pointer(&data)))
			nox_script_push(sa1)
			continue
		case 5:
			sf1 = float32(nox_script_nextFloat((**float32)(unsafe.Pointer(&data))))
			nox_script_pushf(sf1)
			continue
		case 7:
			sa1 = nox_script_pop()
			sa2 = nox_script_pop()
			nox_script_push(sa2 + sa1)
			continue
		case 8:
			sf1 = nox_script_popf()
			sf2 = nox_script_popf()
			sf1 += sf2
			nox_script_pushf(sf1)
			continue
		case 9:
			sa1 = nox_script_pop()
			sa2 = nox_script_pop()
			nox_script_push(sa2 - sa1)
			continue
		case 10:
			sf1 = nox_script_popf()
			sf2 = nox_script_popf()
			sf2 -= sf1
			nox_script_pushf(sf2)
			continue
		case 11:
			sa1 = nox_script_pop()
			sa2 = nox_script_pop()
			nox_script_push(sa1 * sa2)
			continue
		case 12:
			sf1 = nox_script_popf()
			sf2 = nox_script_popf()
			sf1 *= sf2
			nox_script_pushf(sf1)
			continue
		case 13:
			sa1 = nox_script_pop()
			sa2 = nox_script_pop()
			nox_script_push(sa2 / sa1)
			continue
		case 14:
			sf1 = nox_script_popf()
			sf2 = nox_script_popf()
			sf2 /= sf1
			nox_script_pushf(sf2)
			continue
		case 15:
			sa1 = nox_script_pop()
			sa2 = nox_script_pop()
			nox_script_push(sa2 % sa1)
			continue
		case 16:
			sa1 = nox_script_pop()
			sa2 = nox_script_pop()
			nox_script_push(sa2 & sa1)
			continue
		case 17:
			sa1 = nox_script_pop()
			sa2 = nox_script_pop()
			nox_script_push(sa2 | sa1)
			continue
		case 18:
			sa1 = nox_script_pop()
			sa2 = nox_script_pop()
			nox_script_push(sa2 ^ sa1)
			continue
		case 19:
			data = unsafe.Pointer(uintptr(uint32(uintptr(script.data)) + uint32(nox_script_nextInt((**int32)(unsafe.Pointer(&data)))*4)))
			continue
		case 20:
			sa1 = nox_script_nextInt((**int32)(unsafe.Pointer(&data)))
			if nox_script_pop() != 0 {
				data = unsafe.Pointer(uintptr(uint32(uintptr(script.data)) + uint32(sa1*4)))
			}
			continue
		case 21:
			sa1 = nox_script_nextInt((**int32)(unsafe.Pointer(&data)))
			if nox_script_pop() == 0 {
				data = unsafe.Pointer(uintptr(uint32(uintptr(script.data)) + uint32(sa1*4)))
			}
			continue
		case 22:
			fallthrough
		case 24:
			sa1 = nox_script_pop()
			sa2 = nox_script_pop()
			if nox_script_pop() != 0 {
				if (*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*1))).field_28 != nil {
					*(*uint32)(unsafe.Add(unsafe.Pointer((*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*1))).field_28), unsafe.Sizeof(uint32(0))*uintptr(sa2))) = uint32(sa1)
				}
			} else if sa2 < 0 {
				*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(int32(uintptr(a3)) + 760))) - uint32(sa2*4)))) = uint32(sa1)
			} else if script.field_28 != nil {
				*(*uint32)(unsafe.Add(unsafe.Pointer(script.field_28), unsafe.Sizeof(uint32(0))*uintptr(sa2))) = uint32(sa1)
			}
			nox_script_push(sa1)
			continue
		case 23:
			sf1 = nox_script_popf()
			sa2 = nox_script_pop()
			if nox_script_pop() != 0 {
				if (*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*1))).field_28 != nil {
					*(*float32)(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer((*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*1))).field_28), unsafe.Sizeof(uint32(0))*uintptr(sa2))))) = sf1
				}
			} else if sa2 < 0 {
				*(*float32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(int32(uintptr(a3)) + 760))) - uint32(sa2*4)))) = sf1
			} else if script.field_28 != nil {
				*(*float32)(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer(script.field_28), unsafe.Sizeof(uint32(0))*uintptr(sa2))))) = sf1
			}
			nox_script_pushf(sf1)
			continue
		case 25:
			sa1 = nox_script_pop()
			sa2 = nox_script_pop()
			if nox_script_pop() != 0 {
				if (*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*1))).field_28 != nil {
					*(*uint32)(unsafe.Add(unsafe.Pointer((*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*1))).field_28), unsafe.Sizeof(uint32(0))*uintptr(sa2))) *= uint32(sa1)
					nox_script_push(int32(*(*uint32)(unsafe.Add(unsafe.Pointer((*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*1))).field_28), unsafe.Sizeof(uint32(0))*uintptr(sa2)))))
				} else {
					nox_script_push(0)
				}
			} else if sa2 < 0 {
				*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(int32(uintptr(a3)) + 760))) - uint32(sa2*4)))) *= uint32(sa1)
				nox_script_push(int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(int32(uintptr(a3)) + 760))) - uint32(sa2*4))))))
			} else if script.field_28 != nil {
				*(*uint32)(unsafe.Add(unsafe.Pointer(script.field_28), unsafe.Sizeof(uint32(0))*uintptr(sa2))) *= uint32(sa1)
				nox_script_push(int32(*(*uint32)(unsafe.Add(unsafe.Pointer(script.field_28), unsafe.Sizeof(uint32(0))*uintptr(sa2)))))
			} else {
				nox_script_push(0)
			}
			continue
		case 26:
			sf1 = nox_script_popf()
			sa1 = nox_script_pop()
			if nox_script_pop() != 0 {
				if (*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*1))).field_28 != nil {
					*(*float32)(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer((*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*1))).field_28), unsafe.Sizeof(uint32(0))*uintptr(sa1))))) *= sf1
					nox_script_pushf(*(*float32)(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer((*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*1))).field_28), unsafe.Sizeof(uint32(0))*uintptr(sa1))))))
				} else {
					nox_script_pushf(0)
				}
			} else if sa1 < 0 {
				*(*float32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(int32(uintptr(a3)) + 760))) - uint32(sa1*4)))) *= sf1
				nox_script_pushf(*(*float32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(int32(uintptr(a3)) + 760))) - uint32(sa1*4)))))
			} else if script.field_28 != nil {
				*(*float32)(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer(script.field_28), unsafe.Sizeof(uint32(0))*uintptr(sa1))))) *= sf1
				nox_script_pushf(*(*float32)(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer(script.field_28), unsafe.Sizeof(uint32(0))*uintptr(sa1))))))
			} else {
				nox_script_pushf(0)
			}
			continue
		case 27:
			sa1 = nox_script_pop()
			sa2 = nox_script_pop()
			if nox_script_pop() != 0 {
				if (*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*1))).field_28 != nil {
					*(*uint32)(unsafe.Add(unsafe.Pointer((*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*1))).field_28), unsafe.Sizeof(uint32(0))*uintptr(sa2))) /= uint32(sa1)
					nox_script_push(int32(*(*uint32)(unsafe.Add(unsafe.Pointer((*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*1))).field_28), unsafe.Sizeof(uint32(0))*uintptr(sa2)))))
				} else {
					nox_script_push(0)
				}
			} else if sa2 < 0 {
				*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(int32(uintptr(a3)) + 760))) - uint32(sa2*4)))) /= uint32(sa1)
				nox_script_push(int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(int32(uintptr(a3)) + 760))) - uint32(sa2*4))))))
			} else if script.field_28 != nil {
				*(*uint32)(unsafe.Add(unsafe.Pointer(script.field_28), unsafe.Sizeof(uint32(0))*uintptr(sa2))) /= uint32(sa1)
				nox_script_push(int32(*(*uint32)(unsafe.Add(unsafe.Pointer(script.field_28), unsafe.Sizeof(uint32(0))*uintptr(sa2)))))
			} else {
				nox_script_push(0)
			}
			continue
		case 28:
			sf1 = nox_script_popf()
			sa1 = nox_script_pop()
			if nox_script_pop() != 0 {
				if (*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*1))).field_28 != nil {
					*(*float32)(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer((*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*1))).field_28), unsafe.Sizeof(uint32(0))*uintptr(sa1))))) /= sf1
					nox_script_pushf(*(*float32)(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer((*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*1))).field_28), unsafe.Sizeof(uint32(0))*uintptr(sa1))))))
				} else {
					nox_script_pushf(0)
				}
			} else if sa1 < 0 {
				*(*float32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(int32(uintptr(a3)) + 760))) - uint32(sa1*4)))) /= sf1
				nox_script_pushf(*(*float32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(int32(uintptr(a3)) + 760))) - uint32(sa1*4)))))
			} else if script.field_28 != nil {
				*(*float32)(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer(script.field_28), unsafe.Sizeof(uint32(0))*uintptr(sa1))))) /= sf1
				nox_script_pushf(*(*float32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(script.field_28), unsafe.Sizeof(uint32(0))*uintptr(sa1)))))))
			} else {
				nox_script_pushf(0)
			}
			continue
		case 29:
			sa1 = nox_script_pop()
			sa2 = nox_script_pop()
			if nox_script_pop() != 0 {
				if (*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*1))).field_28 != nil {
					*(*uint32)(unsafe.Add(unsafe.Pointer((*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*1))).field_28), unsafe.Sizeof(uint32(0))*uintptr(sa2))) += uint32(sa1)
					nox_script_push(int32(*(*uint32)(unsafe.Add(unsafe.Pointer((*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*1))).field_28), unsafe.Sizeof(uint32(0))*uintptr(sa2)))))
				} else {
					nox_script_push(0)
				}
			} else if sa2 < 0 {
				*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(int32(uintptr(a3)) + 760))) - uint32(sa2*4)))) += uint32(sa1)
				nox_script_push(int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(int32(uintptr(a3)) + 760))) - uint32(sa2*4))))))
			} else if script.field_28 != nil {
				*(*uint32)(unsafe.Add(unsafe.Pointer(script.field_28), unsafe.Sizeof(uint32(0))*uintptr(sa2))) += uint32(sa1)
				nox_script_push(int32(*(*uint32)(unsafe.Add(unsafe.Pointer(script.field_28), unsafe.Sizeof(uint32(0))*uintptr(sa2)))))
			} else {
				nox_script_push(0)
			}
			continue
		case 30:
			sf1 = nox_script_popf()
			sa1 = nox_script_pop()
			if nox_script_pop() != 0 {
				if (*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*1))).field_28 != nil {
					*(*float32)(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer((*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*1))).field_28), unsafe.Sizeof(uint32(0))*uintptr(sa1))))) += sf1
					nox_script_pushf(*(*float32)(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer((*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*1))).field_28), unsafe.Sizeof(uint32(0))*uintptr(sa1))))))
				} else {
					nox_script_pushf(0)
				}
			} else if sa1 < 0 {
				*(*float32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(int32(uintptr(a3)) + 760))) - uint32(sa1*4)))) += sf1
				nox_script_pushf(*(*float32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(int32(uintptr(a3)) + 760))) - uint32(sa1*4)))))
			} else if script.field_28 != nil {
				*(*float32)(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer(script.field_28), unsafe.Sizeof(uint32(0))*uintptr(sa1))))) += sf1
				nox_script_pushf(*(*float32)(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer(script.field_28), unsafe.Sizeof(uint32(0))*uintptr(sa1))))))
			} else {
				nox_script_pushf(0)
			}
			continue
		case 31:
			sa1 = nox_script_pop()
			sa2 = nox_script_pop()
			sa3 = nox_script_pop()
			if sa3 != 0 {
				if (*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*1))).field_28 != nil {
					sa4 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer((*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*1))).field_28), unsafe.Sizeof(uint32(0))*uintptr(sa2))))
				} else {
					sa4 = 0
				}
			} else if sa2 < 0 {
				sa4 = int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(int32(uintptr(a3)) + 760))) - uint32(sa2*4)))))
			} else {
				sa4 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(script.field_28), unsafe.Sizeof(uint32(0))*uintptr(sa2))))
			}
			nox_sprintf(&buf[0], CString("%s%s"), nox_script_getString_512E40(sa4), nox_script_getString_512E40(sa1))
			sa4 = nox_script_addString_512E40(&buf[0])
			if sa3 != 0 {
				if (*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*1))).field_28 != nil {
					*(*uint32)(unsafe.Add(unsafe.Pointer((*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*1))).field_28), unsafe.Sizeof(uint32(0))*uintptr(sa2))) = uint32(sa4)
				}
			} else if sa2 < 0 {
				*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(int32(uintptr(a3)) + 760))) - uint32(sa2*4)))) = uint32(sa4)
			} else if script.field_28 != nil {
				*(*uint32)(unsafe.Add(unsafe.Pointer(script.field_28), unsafe.Sizeof(uint32(0))*uintptr(sa2))) = uint32(sa4)
			}
			nox_script_push(sa4)
			continue
		case 32:
			sa1 = nox_script_pop()
			sa2 = nox_script_pop()
			if nox_script_pop() != 0 {
				if (*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*1))).field_28 != nil {
					*(*uint32)(unsafe.Add(unsafe.Pointer((*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*1))).field_28), unsafe.Sizeof(uint32(0))*uintptr(sa2))) -= uint32(sa1)
					nox_script_push(int32(*(*uint32)(unsafe.Add(unsafe.Pointer((*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*1))).field_28), unsafe.Sizeof(uint32(0))*uintptr(sa2)))))
				} else {
					nox_script_push(0)
				}
			} else if sa2 < 0 {
				*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(int32(uintptr(a3)) + 760))) - uint32(sa2*4)))) -= uint32(sa1)
				nox_script_push(int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(int32(uintptr(a3)) + 760))) - uint32(sa2*4))))))
			} else if script.field_28 != nil {
				*(*uint32)(unsafe.Add(unsafe.Pointer(script.field_28), unsafe.Sizeof(uint32(0))*uintptr(sa2))) -= uint32(sa1)
				nox_script_push(int32(*(*uint32)(unsafe.Add(unsafe.Pointer(script.field_28), unsafe.Sizeof(uint32(0))*uintptr(sa2)))))
			} else {
				nox_script_push(0)
			}
			continue
		case 33:
			sf1 = nox_script_popf()
			sa1 = nox_script_pop()
			if nox_script_pop() != 0 {
				if (*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*1))).field_28 != nil {
					*(*float32)(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer((*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*1))).field_28), unsafe.Sizeof(uint32(0))*uintptr(sa1))))) -= sf1
					nox_script_pushf(*(*float32)(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer((*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*1))).field_28), unsafe.Sizeof(uint32(0))*uintptr(sa1))))))
				} else {
					nox_script_pushf(0)
				}
			} else if sa1 < 0 {
				*(*float32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(int32(uintptr(a3)) + 760))) - uint32(sa1*4)))) -= sf1
				nox_script_push(int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(int32(uintptr(a3)) + 760))) - uint32(sa1*4))))))
			} else if script.field_28 != nil {
				*(*float32)(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer(script.field_28), unsafe.Sizeof(uint32(0))*uintptr(sa1))))) -= sf1
				nox_script_pushf(*(*float32)(unsafe.Pointer(uintptr(uint32(uintptr(unsafe.Pointer(script.field_28))) + uint32(sa1*4)))))
			} else {
				nox_script_pushf(0)
			}
			continue
		case 34:
			sa1 = nox_script_pop()
			sa2 = nox_script_pop()
			if nox_script_pop() != 0 {
				if (*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*1))).field_28 != nil {
					*(*uint32)(unsafe.Add(unsafe.Pointer((*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*1))).field_28), unsafe.Sizeof(uint32(0))*uintptr(sa2))) %= uint32(sa1)
					nox_script_push(int32(*(*uint32)(unsafe.Add(unsafe.Pointer((*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*1))).field_28), unsafe.Sizeof(uint32(0))*uintptr(sa2)))))
				} else {
					nox_script_push(0)
				}
			} else if sa2 < 0 {
				*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(int32(uintptr(a3)) + 760))) - uint32(sa2*4)))) %= uint32(sa1)
				nox_script_push(int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(int32(uintptr(a3)) + 760))) - uint32(sa2*4))))))
			} else if script.field_28 != nil {
				*(*uint32)(unsafe.Add(unsafe.Pointer(script.field_28), unsafe.Sizeof(uint32(0))*uintptr(sa2))) %= uint32(sa1)
				nox_script_push(int32(*(*uint32)(unsafe.Add(unsafe.Pointer(script.field_28), unsafe.Sizeof(uint32(0))*uintptr(sa2)))))
			} else {
				nox_script_push(0)
			}
			continue
		case 35:
			sa1 = nox_script_pop()
			sa2 = nox_script_pop()
			nox_script_push(bool2int(sa2 == sa1))
			continue
		case 36:
			sf1 = nox_script_popf()
			sf2 = nox_script_popf()
			nox_script_push(bool2int(sf2 == sf1))
			continue
		case 37:
			sa1 = nox_script_pop()
			sa2 = nox_script_pop()
			nox_script_push(bool2int(libc.StrCmp(nox_script_getString_512E40(sa2), nox_script_getString_512E40(sa1)) == 0))
			continue
		case 38:
			sa1 = nox_script_pop()
			sa2 = nox_script_pop()
			nox_script_push(sa2 << sa1)
			continue
		case 39:
			sa1 = nox_script_pop()
			sa2 = nox_script_pop()
			nox_script_push(sa2 >> sa1)
			continue
		case 40:
			sa1 = nox_script_pop()
			sa2 = nox_script_pop()
			nox_script_push(bool2int(sa2 < sa1))
			continue
		case 41:
			sf1 = nox_script_popf()
			sf2 = nox_script_popf()
			nox_script_push(bool2int(sf2 < sf1))
			continue
		case 42:
			sa1 = nox_script_pop()
			sa2 = nox_script_pop()
			sa3 = int32(libc.StrCmp(nox_script_getString_512E40(sa2), nox_script_getString_512E40(sa1)))
			if sa3 != 0 {
				nox_script_push(bool2int(-bool2int(sa3 < 0)-(bool2int(sa3 < 0)-1) < 0))
			} else {
				nox_script_push(0)
			}
			continue
		case 43:
			sa1 = nox_script_pop()
			sa2 = nox_script_pop()
			nox_script_push(bool2int(sa2 > sa1))
			continue
		case 44:
			sf1 = nox_script_popf()
			sf2 = nox_script_popf()
			nox_script_push(bool2int(sf2 > sf1))
			continue
		case 45:
			sa1 = nox_script_pop()
			sa2 = nox_script_pop()
			sa3 = int32(libc.StrCmp(nox_script_getString_512E40(sa2), nox_script_getString_512E40(sa1)))
			if sa3 != 0 {
				nox_script_push(bool2int(-bool2int(sa3 < 0)-(bool2int(sa3 < 0)-1) > 0))
			} else {
				nox_script_push(0)
			}
			continue
		case 46:
			sa1 = nox_script_pop()
			sa2 = nox_script_pop()
			nox_script_push(bool2int(sa2 <= sa1))
			continue
		case 47:
			sf1 = nox_script_popf()
			sf2 = nox_script_popf()
			nox_script_push(bool2int(sf2 <= sf1))
			continue
		case 48:
			sa1 = nox_script_pop()
			sa2 = nox_script_pop()
			sa3 = int32(libc.StrCmp(nox_script_getString_512E40(sa2), nox_script_getString_512E40(sa1)))
			if sa3 != 0 {
				nox_script_push(bool2int(-bool2int(sa3 < 0)-(bool2int(sa3 < 0)-1) <= 0))
			} else {
				nox_script_push(1)
			}
			continue
		case 49:
			sa1 = nox_script_pop()
			sa2 = nox_script_pop()
			nox_script_push(bool2int(sa2 >= sa1))
			continue
		case 50:
			sf1 = nox_script_popf()
			sf2 = nox_script_popf()
			nox_script_push(bool2int(sf2 >= sf1))
			continue
		case 51:
			sa1 = nox_script_pop()
			sa2 = nox_script_pop()
			sa3 = int32(libc.StrCmp(nox_script_getString_512E40(sa2), nox_script_getString_512E40(sa1)))
			if sa3 != 0 {
				nox_script_push(bool2int(-bool2int(sa3 < 0)-(bool2int(sa3 < 0)-1) >= 0))
			} else {
				nox_script_push(1)
			}
			continue
		case 52:
			sa1 = nox_script_pop()
			sa2 = nox_script_pop()
			nox_script_push(bool2int(sa2 != sa1))
			continue
		case 53:
			sf1 = nox_script_popf()
			sf2 = nox_script_popf()
			nox_script_push(bool2int(sf2 != sf1))
			continue
		case 54:
			sa1 = nox_script_pop()
			sa2 = nox_script_pop()
			nox_script_push(bool2int(libc.StrCmp(nox_script_getString_512E40(sa2), nox_script_getString_512E40(sa1)) != 0))
			continue
		case 55:
			sa1 = nox_script_pop()
			sa2 = nox_script_pop()
			nox_script_push(bool2int(sa2 != 0 && sa1 != 0))
			continue
		case 56:
			sa1 = nox_script_pop()
			sa2 = nox_script_pop()
			nox_script_push(bool2int(sa2 != 0 || sa1 != 0))
			continue
		case 57:
			sa1 = nox_script_pop()
			sa2 = nox_script_pop()
			if nox_script_pop() != 0 {
				if (*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*1))).field_28 != nil {
					*(*uint32)(unsafe.Add(unsafe.Pointer((*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*1))).field_28), unsafe.Sizeof(uint32(0))*uintptr(sa2))) <<= uint32(sa1)
					nox_script_push(int32(*(*uint32)(unsafe.Add(unsafe.Pointer((*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*1))).field_28), unsafe.Sizeof(uint32(0))*uintptr(sa2)))))
				} else {
					nox_script_push(0)
				}
			} else if sa2 < 0 {
				*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(int32(uintptr(a3)) + 760))) - uint32(sa2*4)))) <<= uint32(sa1)
				nox_script_push(int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(int32(uintptr(a3)) + 760))) - uint32(sa2*4))))))
			} else if script.field_28 != nil {
				*(*uint32)(unsafe.Add(unsafe.Pointer(script.field_28), unsafe.Sizeof(uint32(0))*uintptr(sa2))) <<= uint32(sa1)
				nox_script_push(int32(*(*uint32)(unsafe.Add(unsafe.Pointer(script.field_28), unsafe.Sizeof(uint32(0))*uintptr(sa2)))))
			} else {
				nox_script_push(0)
			}
			continue
		case 58:
			sa1 = nox_script_pop()
			sa2 = nox_script_pop()
			if nox_script_pop() != 0 {
				if (*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*1))).field_28 != nil {
					*(*uint32)(unsafe.Add(unsafe.Pointer((*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*1))).field_28), unsafe.Sizeof(uint32(0))*uintptr(sa2))) >>= uint32(sa1)
					nox_script_push(int32(*(*uint32)(unsafe.Add(unsafe.Pointer((*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*1))).field_28), unsafe.Sizeof(uint32(0))*uintptr(sa2)))))
				} else {
					nox_script_push(0)
				}
			} else if sa2 < 0 {
				*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(int32(uintptr(a3)) + 760))) - uint32(sa2*4)))) >>= uint32(sa1)
				nox_script_push(int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(int32(uintptr(a3)) + 760))) - uint32(sa2*4))))))
			} else if script.field_28 != nil {
				*(*uint32)(unsafe.Add(unsafe.Pointer(script.field_28), unsafe.Sizeof(uint32(0))*uintptr(sa2))) >>= uint32(sa1)
				nox_script_push(int32(*(*uint32)(unsafe.Add(unsafe.Pointer(script.field_28), unsafe.Sizeof(uint32(0))*uintptr(sa2)))))
			} else {
				nox_script_push(0)
			}
			continue
		case 59:
			sa1 = nox_script_pop()
			sa2 = nox_script_pop()
			if nox_script_pop() != 0 {
				if (*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*1))).field_28 != nil {
					*(*uint32)(unsafe.Add(unsafe.Pointer((*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*1))).field_28), unsafe.Sizeof(uint32(0))*uintptr(sa2))) &= uint32(sa1)
					nox_script_push(int32(*(*uint32)(unsafe.Add(unsafe.Pointer((*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*1))).field_28), unsafe.Sizeof(uint32(0))*uintptr(sa2)))))
				} else {
					nox_script_push(0)
				}
			} else if sa2 < 0 {
				*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(int32(uintptr(a3)) + 760))) - uint32(sa2*4)))) &= uint32(sa1)
				nox_script_push(int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(int32(uintptr(a3)) + 760))) - uint32(sa2*4))))))
			} else if script.field_28 != nil {
				*(*uint32)(unsafe.Add(unsafe.Pointer(script.field_28), unsafe.Sizeof(uint32(0))*uintptr(sa2))) &= uint32(sa1)
				nox_script_push(int32(*(*uint32)(unsafe.Add(unsafe.Pointer(script.field_28), unsafe.Sizeof(uint32(0))*uintptr(sa2)))))
			} else {
				nox_script_push(0)
			}
			continue
		case 60:
			sa1 = nox_script_pop()
			sa2 = nox_script_pop()
			if nox_script_pop() != 0 {
				if (*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*1))).field_28 != nil {
					*(*uint32)(unsafe.Add(unsafe.Pointer((*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*1))).field_28), unsafe.Sizeof(uint32(0))*uintptr(sa2))) |= uint32(sa1)
					nox_script_push(int32(*(*uint32)(unsafe.Add(unsafe.Pointer((*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*1))).field_28), unsafe.Sizeof(uint32(0))*uintptr(sa2)))))
				} else {
					nox_script_push(0)
				}
			} else if sa2 < 0 {
				*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(int32(uintptr(a3)) + 760))) - uint32(sa2*4)))) |= uint32(sa1)
				nox_script_push(int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(int32(uintptr(a3)) + 760))) - uint32(sa2*4))))))
			} else if script.field_28 != nil {
				*(*uint32)(unsafe.Add(unsafe.Pointer(script.field_28), unsafe.Sizeof(uint32(0))*uintptr(sa2))) |= uint32(sa1)
				nox_script_push(int32(*(*uint32)(unsafe.Add(unsafe.Pointer(script.field_28), unsafe.Sizeof(uint32(0))*uintptr(sa2)))))
			} else {
				nox_script_push(0)
			}
			continue
		case 61:
			sa1 = nox_script_pop()
			sa2 = nox_script_pop()
			if nox_script_pop() != 0 {
				if (*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*1))).field_28 != nil {
					*(*uint32)(unsafe.Add(unsafe.Pointer((*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*1))).field_28), unsafe.Sizeof(uint32(0))*uintptr(sa2))) ^= uint32(sa1)
					nox_script_push(int32(*(*uint32)(unsafe.Add(unsafe.Pointer((*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*1))).field_28), unsafe.Sizeof(uint32(0))*uintptr(sa2)))))
				} else {
					nox_script_push(0)
				}
			} else if sa2 < 0 {
				*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(int32(uintptr(a3)) + 760))) - uint32(sa2*4)))) ^= uint32(sa1)
				nox_script_push(int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(int32(uintptr(a3)) + 760))) - uint32(sa2*4))))))
			} else if script.field_28 != nil {
				*(*uint32)(unsafe.Add(unsafe.Pointer(script.field_28), unsafe.Sizeof(uint32(0))*uintptr(sa2))) ^= uint32(sa1)
				nox_script_push(int32(*(*uint32)(unsafe.Add(unsafe.Pointer(script.field_28), unsafe.Sizeof(uint32(0))*uintptr(sa2)))))
			} else {
				nox_script_push(0)
			}
			continue
		case 62:
			sa1 = nox_script_pop()
			nox_script_push(bool2int(sa1 == 0))
			continue
		case 63:
			sa1 = nox_script_pop()
			nox_script_push(^sa1)
			continue
		case 64:
			sa1 = nox_script_pop()
			nox_script_push(-sa1)
			continue
		case 65:
			sf1 = -nox_script_popf()
			nox_script_pushf(sf1)
			continue
		case 66:
			sa1 = nox_script_pop()
			sa2 = nox_script_pop()
			sa3 = nox_script_pop()
			sa4 = nox_script_pop()
			sa5 = 0
			if sa1 < 0 || sa1 >= sa4 {
				sa5 = 1
			}
			if sa3 != 0 {
				if (*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*1))).field_28 != nil {
					sa4 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer((*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*1))).field_28), unsafe.Sizeof(uint32(0))*uintptr(sa1+sa2))))
				} else {
					sa4 = 0
				}
			} else if sa2 < 0 {
				sa4 = int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(int32(uintptr(a3)) + 760))) + uint32((sa2-sa1)*4)))))
			} else if script.field_28 != nil {
				sa4 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(script.field_28), unsafe.Sizeof(uint32(0))*uintptr(sa1+sa2))))
			} else {
				sa4 = 0
			}
			nox_script_push(sa4)
			if sa5 != 0 {
				break
			}
			continue
		case 67:
			sa1 = nox_script_pop()
			sa2 = nox_script_pop()
			sa3 = nox_script_pop()
			sa4 = nox_script_pop()
			sa5 = 0
			if sa1 < 0 || sa1 > sa4 {
				sa5 = 1
			}
			if sa3 != 0 {
				if (*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*1))).field_28 != nil {
					sf1 = *(*float32)(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer((*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*1))).field_28), unsafe.Sizeof(uint32(0))*uintptr(sa1+sa2)))))
				} else {
					sf1 = 0
				}
			} else if sa2 < 0 {
				sf1 = *(*float32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(int32(uintptr(a3)) + 760))) + uint32((sa2-sa1)*4))))
			} else if script.field_28 != nil {
				sf1 = *(*float32)(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer(script.field_28), unsafe.Sizeof(uint32(0))*uintptr(sa1+sa2)))))
			} else {
				sf1 = 0
			}
			nox_script_pushf(sf1)
			if sa5 != 0 {
				break
			}
			continue
		case 68:
			sa1 = nox_script_pop()
			sa2 = nox_script_pop()
			sa3 = nox_script_pop()
			sa4 = nox_script_pop()
			sa5 = 0
			if sa1 < 0 || sa1 > sa4 {
				sa5 = 1
			}
			nox_script_push(sa3)
			if sa2 < 0 {
				nox_script_push(sa2 - sa1)
			} else {
				nox_script_push(sa2 + sa1)
			}
			if sa5 != 0 {
				break
			}
			continue
		case 69:
			sa1 = nox_script_nextInt((**int32)(unsafe.Pointer(&data)))
			if nox_script_callBuiltin_508B70(index, sa1) == 1 {
				break
			}
			continue
		case 70:
			sa1 = nox_script_nextInt((**int32)(unsafe.Pointer(&data)))
			nox_script_callByIndex_507310(sa1, a2, a3)
			continue
		case 73:
			sa1 = nox_script_pop()
			sa2 = nox_script_pop()
			nox_sprintf(&buf[0], CString("%s%s"), nox_script_getString_512E40(sa2), nox_script_getString_512E40(sa1))
			sa3 = nox_script_addString_512E40(&buf[0])
			nox_script_push(sa3)
			continue
		default:
		}
		if uint32(nox_script_stack_top) != script.stack_size+uint32(bstack) {
			if script.stack_size != 0 {
				if nox_script_stack_top != 0 {
					var v int32 = nox_script_pop()
					nox_script_stack_top = bstack
					nox_script_push(v)
				} else {
					nox_script_stack_top = bstack
					nox_script_push(0)
				}
			} else {
				nox_script_stack_top = bstack
			}
		}
		return
	}
}
func sub_508CB0(a1 *uint32, a2 int32) int32 {
	var (
		v2     int32
		result int32
		v4     int32
		v5     *uint32
		v6     *uint32
		v7     int32
		v8     *uint32
	)
	v2 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*189)))
	if v2 == 0 {
		return 0
	}
	if a2 == 14 {
		if noxflags.HasGame(noxflags.GameFlag22 | noxflags.GameFlag23) {
			result = v2
		} else {
			result = int32(uintptr(unsafe.Pointer((*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*192)))))).field_0)))
		}
		return result
	}
	v4 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*2)))
	if v4&512 != 0 {
		v5 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*187)))))
		if a2 != 0 {
			if a2 != 1 {
				if a2 == 2 {
					if noxflags.HasGame(noxflags.GameFlag22 | noxflags.GameFlag23) {
						result = v2 + 384
					} else {
						result = int32(uintptr(unsafe.Pointer((*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(uint32(0))*8)))))).field_0)))
					}
					return result
				}
				return 0
			}
			if noxflags.HasGame(noxflags.GameFlag22 | noxflags.GameFlag23) {
				result = v2 + 256
			} else {
				result = int32(uintptr(unsafe.Pointer((*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(uint32(0))*6)))))).field_0)))
			}
		} else if noxflags.HasGame(noxflags.GameFlag22 | noxflags.GameFlag23) {
			result = v2 + 512
		} else {
			result = int32(uintptr(unsafe.Pointer((*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(uint32(0))*4)))))).field_0)))
		}
	} else if v4&2 != 0 {
		v6 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*187)))))
		switch a2 {
		case 3:
			if noxflags.HasGame(noxflags.GameFlag22 | noxflags.GameFlag23) {
				result = v2 + 640
			} else {
				result = int32(uintptr(unsafe.Pointer((*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*309)))))).field_0)))
			}
		case 4:
			if noxflags.HasGame(noxflags.GameFlag22 | noxflags.GameFlag23) {
				result = v2 + 768
			} else {
				result = int32(uintptr(unsafe.Pointer((*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*307)))))).field_0)))
			}
		case 5:
			if noxflags.HasGame(noxflags.GameFlag22 | noxflags.GameFlag23) {
				result = v2 + 896
			} else {
				result = int32(uintptr(unsafe.Pointer((*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*317)))))).field_0)))
			}
		case 6:
			if noxflags.HasGame(noxflags.GameFlag22 | noxflags.GameFlag23) {
				result = v2 + 1024
			} else {
				result = int32(uintptr(unsafe.Pointer((*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*311)))))).field_0)))
			}
		case 7:
			if noxflags.HasGame(noxflags.GameFlag22 | noxflags.GameFlag23) {
				result = v2 + 1152
			} else {
				result = int32(uintptr(unsafe.Pointer((*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*313)))))).field_0)))
			}
		case 8:
			if noxflags.HasGame(noxflags.GameFlag22 | noxflags.GameFlag23) {
				result = v2 + 1280
			} else {
				result = int32(uintptr(unsafe.Pointer((*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*315)))))).field_0)))
			}
		case 9:
			if noxflags.HasGame(noxflags.GameFlag22 | noxflags.GameFlag23) {
				result = v2 + 1408
			} else {
				result = int32(uintptr(unsafe.Pointer((*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*319)))))).field_0)))
			}
		case 10:
			if noxflags.HasGame(noxflags.GameFlag22 | noxflags.GameFlag23) {
				result = v2 + 1536
			} else {
				result = int32(uintptr(unsafe.Pointer((*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*321)))))).field_0)))
			}
		case 11:
			if noxflags.HasGame(noxflags.GameFlag22 | noxflags.GameFlag23) {
				result = v2 + 1664
			} else {
				result = int32(uintptr(unsafe.Pointer((*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*323)))))).field_0)))
			}
		case 13:
			if noxflags.HasGame(noxflags.GameFlag22 | noxflags.GameFlag23) {
				result = v2 + 1792
			} else {
				result = int32(uintptr(unsafe.Pointer((*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v6), unsafe.Sizeof(uint32(0))*325)))))).field_0)))
			}
		default:
			return 0
		}
	} else if v4&2048 != 0 {
		v7 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*175)))
		if a2 != 12 {
			return 0
		}
		if noxflags.HasGame(noxflags.GameFlag22 | noxflags.GameFlag23) {
			result = v2 + 128
		} else {
			result = int32(uintptr(unsafe.Pointer((*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*uintptr(*(*uint32)(unsafe.Pointer(uintptr(v7 + 4))))))).field_0)))
		}
	} else {
		if (uint32(v4) & 0x20000) == 0 {
			return 0
		}
		v8 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*187)))))
		switch a2 {
		case 15:
			if noxflags.HasGame(noxflags.GameFlag22 | noxflags.GameFlag23) {
				result = v2 + 1920
			} else {
				result = int32(uintptr(unsafe.Pointer((*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v8), unsafe.Sizeof(uint32(0))*13)))))).field_0)))
			}
		case 16:
			if noxflags.HasGame(noxflags.GameFlag22 | noxflags.GameFlag23) {
				result = v2 + 2048
			} else {
				result = int32(uintptr(unsafe.Pointer((*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v8), unsafe.Sizeof(uint32(0))*15)))))).field_0)))
			}
		case 17:
			if noxflags.HasGame(noxflags.GameFlag22 | noxflags.GameFlag23) {
				result = v2 + 2304
			} else {
				result = int32(uintptr(unsafe.Pointer((*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v8), unsafe.Sizeof(uint32(0))*17)))))).field_0)))
			}
		case 18:
			if noxflags.HasGame(noxflags.GameFlag22 | noxflags.GameFlag23) {
				result = v2 + 2176
			} else {
				result = int32(uintptr(unsafe.Pointer((*(*nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(nox_script_arr_xxx_1599636), unsafe.Sizeof(nox_script_xxx_t{})*uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v8), unsafe.Sizeof(uint32(0))*19)))))).field_0)))
			}
		default:
			return 0
		}
	}
	return result
}
