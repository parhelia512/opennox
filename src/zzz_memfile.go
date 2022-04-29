package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	"github.com/gotranspile/cxgo/runtime/stdio"
	"github.com/noxworld-dev/opennox/v1/common/alloc"
	"unsafe"
)

type nox_memfile struct {
	data *byte
	size int32
	cur  *byte
	end  *byte
}

func nox_memfile_load(path *byte, a2 int32) *nox_memfile {
	var nf *nox_memfile = new(nox_memfile)
	if nf == nil {
		return nil
	}
	var f *File = nox_binfile_open_408CC0(path, 0)
	if f == nil {
		nox_memfile_free(nf)
		return nil
	}
	if nox_binfile_cryptSet_408D40(f, a2) == 0 {
		nox_memfile_free(nf)
		nox_binfile_close_408D90(f)
		return nil
	}
	nox_binfile_fseek_409050(f, 0, stdio.SEEK_END)
	nf.size = nox_binfile_ftell_426A50(f)
	nox_binfile_fseek_409050(f, 0, stdio.SEEK_SET)
	nf.data = (*byte)(alloc.Calloc(1, int(nf.size)))
	if nf.data == nil {
		nox_memfile_free(nf)
		nox_binfile_close_408D90(f)
		return nil
	}
	var csize int32 = nf.size
	var x1 int32 = nox_binfile_fread_408E40(nf.data, 1, nf.size, f)
	if x1 != csize {
		nox_memfile_free(nf)
		nox_binfile_close_408D90(f)
		return nil
	}
	nf.cur = nf.data
	nf.end = (*byte)(unsafe.Add(unsafe.Pointer(nf.data), csize))
	nox_binfile_close_408D90(f)
	return nf
}
func nox_memfile_free(f *nox_memfile) {
	if f.data != nil {
		alloc.Free(unsafe.Pointer(f.data))
	}
	f.data = nil
	f.cur = nil
	f.end = nil
	f.size = 0
	alloc.Free(unsafe.Pointer(f))
}
func nox_memfile_read_i8(f *nox_memfile) int8 {
	if f.data == nil {
		return 0
	}
	var v int8 = *(*int8)(unsafe.Pointer(f.cur))
	f.cur = (*byte)(unsafe.Add(unsafe.Pointer(f.cur), 1))
	return v
}
func nox_memfile_read_u8(f *nox_memfile) uint8 {
	if f.data == nil {
		return 0
	}
	var v uint8 = *(*uint8)(unsafe.Pointer(f.cur))
	f.cur = (*byte)(unsafe.Add(unsafe.Pointer(f.cur), 1))
	return v
}
func nox_memfile_read_i16(f *nox_memfile) int16 {
	if f.data == nil {
		return 0
	}
	var v int16 = *(*int16)(unsafe.Pointer(f.cur))
	f.cur = (*byte)(unsafe.Add(unsafe.Pointer(f.cur), 2))
	return v
}
func nox_memfile_read_u16(f *nox_memfile) uint16 {
	if f.data == nil {
		return 0
	}
	var v uint16 = *(*uint16)(unsafe.Pointer(f.cur))
	f.cur = (*byte)(unsafe.Add(unsafe.Pointer(f.cur), 2))
	return v
}
func nox_memfile_read_i32(f *nox_memfile) int32 {
	if f.data == nil {
		return 0
	}
	var v int32 = *(*int32)(unsafe.Pointer(f.cur))
	f.cur = (*byte)(unsafe.Add(unsafe.Pointer(f.cur), 4))
	return v
}
func nox_memfile_read_u32(f *nox_memfile) uint32 {
	if f.data == nil {
		return 0
	}
	var v uint32 = *(*uint32)(unsafe.Pointer(f.cur))
	f.cur = (*byte)(unsafe.Add(unsafe.Pointer(f.cur), 4))
	return v
}
func nox_memfile_skip(f *nox_memfile, n int32) {
	if f.data == nil {
		return
	}
	f.cur = (*byte)(unsafe.Add(unsafe.Pointer(f.cur), n))
}
func nox_memfile_read(dst unsafe.Pointer, sz uint32, cnt int32, f *nox_memfile) uint32 {
	var n uint32 = uint32(cnt) * sz
	if uintptr(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(f.cur), n)))) > uintptr(unsafe.Pointer(f.end)) {
		n = uint32(int32(uintptr(unsafe.Pointer(f.end)) - uintptr(unsafe.Pointer(f.cur))))
	}
	libc.MemCpy(dst, unsafe.Pointer(f.cur), int(n))
	f.cur = (*byte)(unsafe.Add(unsafe.Pointer(f.cur), n))
	return n / sz
}
func nox_memfile_seek_40AD10(memfile *nox_memfile, offset int32, fromStartOrEnd int32) uint32 {
	if fromStartOrEnd == 0 {
		memfile.cur = (*byte)(unsafe.Add(unsafe.Pointer(memfile.data), offset))
	} else if fromStartOrEnd == 2 {
		memfile.cur = (*byte)(unsafe.Add(unsafe.Pointer(memfile.end), offset))
	}
	if uintptr(unsafe.Pointer(memfile.cur)) < uintptr(unsafe.Pointer(memfile.data)) {
		memfile.cur = memfile.data
		return 0
	}
	if uintptr(unsafe.Pointer(memfile.cur)) > uintptr(unsafe.Pointer(memfile.end)) {
		memfile.cur = memfile.end
	}
	return uint32(int32(uintptr(unsafe.Pointer(memfile.cur)) - uintptr(unsafe.Pointer(memfile.data))))
}
func nox_memfile_read64align_40AD60(dest *byte, sz int32, cnt int32, f *nox_memfile) uint32 {
	var (
		cur_offset uint32 = uint32(int32(uintptr(unsafe.Pointer(f.cur)) - uintptr(unsafe.Pointer(f.data))))
		over       uint8  = uint8(cur_offset % 8)
		buf        [8]byte
	)
	if int32(over) != 0 {
		nox_memfile_read(unsafe.Pointer(&buf[0]), uint32(8-int32(over)), 1, f)
	}
	var result uint32 = nox_memfile_read(unsafe.Pointer(&buf[0]), 8, 1, f)
	if result != 1 {
		return result
	}
	libc.MemCpy(unsafe.Pointer(dest), unsafe.Pointer(&buf[0]), int(cnt*sz))
	return 1
}
