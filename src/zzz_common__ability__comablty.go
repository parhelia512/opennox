package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	noxflags "github.com/noxworld-dev/opennox/v1/common/flags"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"unsafe"
)

func nox_thing_read_ABIL_rec_424F00(f *nox_memfile, a2 unsafe.Pointer) int32 {
	var (
		v3  int32
		v5  *uint8
		v7  int32
		v9  int8
		v11 int32
		v13 int8
		v15 int32
		v17 int8
		v20 int32
		v22 uint8
		v23 int32
		v26 uint8
		v27 int32
		v28 uint8
		v29 uint8
		v30 uint8
		v31 uint8
		v32 uint8
		v33 uint8
		v34 *byte
		v35 [256]byte
	)
	v28 = f.ReadU8()
	nox_memfile_read(a2, 1, int32(v28), f)
	*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(a2)), v28))) = 0
	v3 = nox_xxx_abilityNameToN_424D80((*byte)(a2))
	if v3 == 0 {
		return 0
	}
	v5 = (*uint8)(memmap.PtrOff(6112660, uint32(v3*52)+0x9245C))
	*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v5))), unsafe.Sizeof(uint32(0))*9))) = uint32(f.ReadI8())
	v7 = f.ReadI32()
	*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v5))), unsafe.Sizeof(uint32(0))*2))) = 0
	v35[0] = byte(*memmap.PtrUint8(6112660, 599444))
	if v7 == -1 {
		v9 = f.ReadI8()
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v34))), 0)) = uint8(v9)
		v29 = uint8(f.ReadI8())
		nox_memfile_read(unsafe.Pointer(&v35[0]), 1, int32(v29), f)
		v35[v29] = 0
		v7 = -1
	}
	if noxflags.HasGame(noxflags.GameClient) {
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v5))), unsafe.Sizeof(uint32(0))*2))) = uint32(uintptr(unsafe.Pointer(nox_xxx_readImgMB_42FAA0(v7, int8(uintptr(unsafe.Pointer(v34))), &v35[0]))))
	}
	v11 = f.ReadI32()
	*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v5))), unsafe.Sizeof(uint32(0))*3))) = 0
	v35[0] = byte(*memmap.PtrUint8(6112660, 599448))
	if v11 == -1 {
		v13 = f.ReadI8()
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v34))), 0)) = uint8(v13)
		v30 = uint8(f.ReadI8())
		nox_memfile_read(unsafe.Pointer(&v35[0]), 1, int32(v30), f)
		v35[v30] = 0
		v11 = -1
	}
	if noxflags.HasGame(noxflags.GameClient) {
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v5))), unsafe.Sizeof(uint32(0))*3))) = uint32(uintptr(unsafe.Pointer(nox_xxx_readImgMB_42FAA0(v11, int8(uintptr(unsafe.Pointer(v34))), &v35[0]))))
	}
	v15 = f.ReadI32()
	*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v5))), unsafe.Sizeof(uint32(0))*4))) = 0
	v35[0] = byte(*memmap.PtrUint8(6112660, 599452))
	if v15 == -1 {
		v17 = f.ReadI8()
		*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&v34))), 0)) = uint8(v17)
		v31 = uint8(f.ReadI8())
		nox_memfile_read(unsafe.Pointer(&v35[0]), 1, int32(v31), f)
		v35[v31] = 0
		v15 = -1
	}
	if noxflags.HasGame(noxflags.GameClient) {
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v5))), unsafe.Sizeof(uint32(0))*4))) = uint32(uintptr(unsafe.Pointer(nox_xxx_readImgMB_42FAA0(v15, int8(uintptr(unsafe.Pointer(v34))), &v35[0]))))
	}
	v32 = f.ReadU8()
	nox_memfile_read(unsafe.Pointer(&v35[0]), 1, int32(v32), f)
	v35[v32] = 0
	*(*uint32)(unsafe.Pointer(v5)) = uint32(uintptr(unsafe.Pointer(strMan.GetStringInFile(&v35[0], "C:\\NoxPost\\src\\common\\Ability\\ComAblty.c"))))
	v20 = int32(f.ReadI16())
	nox_memfile_read(unsafe.Pointer(&v35[0]), 1, v20, f)
	v35[v20] = 0
	*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v5))), unsafe.Sizeof(uint32(0))*1))) = uint32(uintptr(unsafe.Pointer(strMan.GetStringInFile(&v35[0], "C:\\NoxPost\\src\\common\\Ability\\ComAblty.c"))))
	v22 = f.ReadU8()
	v23 = int32(v22)
	nox_memfile_read(unsafe.Pointer(&v35[0]), 1, int32(v22), f)
	v35[v23] = 0
	*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v5))), unsafe.Sizeof(uint32(0))*10))) = uint32(nox_xxx_utilFindSound_40AF50(&v35[0]))
	v33 = f.ReadU8()
	nox_memfile_read(unsafe.Pointer(&v35[0]), 1, int32(v33), f)
	v35[v33] = 0
	*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v5))), unsafe.Sizeof(uint32(0))*11))) = uint32(nox_xxx_utilFindSound_40AF50(&v35[0]))
	v26 = f.ReadU8()
	v27 = int32(v26)
	nox_memfile_read(unsafe.Pointer(&v35[0]), 1, int32(v26), f)
	v35[v27] = 0
	*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v5))), unsafe.Sizeof(uint32(0))*12))) = uint32(nox_xxx_utilFindSound_40AF50(&v35[0]))
	*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v5))), unsafe.Sizeof(uint32(0))*5))) = 1
	*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v5))), unsafe.Sizeof(uint32(0))*6))) = 1
	return 1
}
