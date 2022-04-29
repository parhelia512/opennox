package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"math"
	"unsafe"
)

func nox_xxx_updDrawManabombCharge_4CCAC0(a1 int32, a2 *uint32) int32 {
	var (
		v2  uint32
		v3  *uint32
		v4  int32
		v5  int32
		v6  int32
		v7  int32
		v8  int32
		v9  int16
		v10 uint32
		v11 *uint8
		v12 int16
		v14 float32
		i   int32
		v16 int16
	)
	_ = v16
	var v17 int16
	_ = v17
	var v18 uint32
	var v19 int32
	v14 = float32(nox_xxx_gamedataGetFloat_419D40(CString("ManaBombOutRadius")))
	v2 = uint32(int(v14))
	if dword_5d4594_1522956 == 0 {
		dword_5d4594_1522956 = uint32(nox_xxx_getTTByNameSpriteMB_44CFC0("ManaBombOrb"))
		*memmap.PtrUint32(6112660, 1522960) = uint32(nox_xxx_getTTByNameSpriteMB_44CFC0("VioletSpark"))
	}
	v3 = a2
	v19 = 20
	v4 = int32(v2 >> 2)
	v18 = v2 >> 4
	for i = int32(v2 >> 2); ; v4 = i {
		v5 = v4 + randomIntMinMax(0, int32(v2))
		if v5 > int32(v2) {
			v5 = int32(v2)
		}
		v6 = randomIntMinMax(0, math.MaxUint8)
		v7 = int32(uintptr(unsafe.Pointer(nox_xxx_spriteLoadAdd_45A360_drawable(*memmap.PtrInt32(6112660, 1522960), int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*3))+uint32(v5**memmap.PtrInt32(0x587000, uint32(v6*8)+0x2EE58)/16)), int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*4))+uint32(v5**memmap.PtrInt32(0x587000, uint32(v6*8)+0x2EE5C)/16))))))
		v8 = v7
		if v7 != 0 {
			*(*uint32)(unsafe.Pointer(uintptr(v7 + 432))) = *(*uint32)(unsafe.Pointer(uintptr(v7 + 12))) << 12
			*(*uint32)(unsafe.Pointer(uintptr(v7 + 436))) = *(*uint32)(unsafe.Pointer(uintptr(v7 + 16))) << 12
			*(*uint8)(unsafe.Pointer(uintptr(v7 + 299))) = 0
			*(*uint32)(unsafe.Pointer(uintptr(v7 + 440))) = 0
			*(*uint32)(unsafe.Pointer(uintptr(v7 + 448))) = nox_frame_xxx_2598000 + uint32(randomIntMinMax(10, 30))
			*(*uint32)(unsafe.Pointer(uintptr(v8 + 444))) = nox_frame_xxx_2598000
			*(*uint16)(unsafe.Pointer(uintptr(v8 + 104))) = 0
			*(*uint8)(unsafe.Pointer(uintptr(v8 + 296))) = uint8(int8(randomIntMinMax(2, 8)))
			nox_xxx_sprite_45A110_drawable((*nox_drawable)(unsafe.Pointer(uintptr(v8))))
		}
		if func() int32 {
			p := &v19
			*p--
			return *p
		}() == 0 {
			break
		}
	}
	if int32(uint8(nox_frame_xxx_2598000))&1 != 0 {
		if (nox_frame_xxx_2598000 - *(*uint32)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint32(0))*80))) < 10 {
			v9 = int16(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v3))), unsafe.Sizeof(uint16(0))*8))))
			*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&i))), unsafe.Sizeof(uint16(0))*0)) = *((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v3))), unsafe.Sizeof(uint16(0))*6)))
			*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&i))), unsafe.Sizeof(uint16(0))*1)) = uint16(v9)
			v10 = nox_frame_xxx_2598000 % 51
			if int32(nox_frame_xxx_2598000%51) < 256 {
				v11 = (*uint8)(memmap.PtrOff(0x587000, v10*8+0x2EE5C))
				for {
					v12 = int16(uint16(v18 * uint32(*(*uint16)(unsafe.Pointer(v11)))))
					v16 = int16(uint16(uint32(i) + v18*uint32(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v11))), -int(unsafe.Sizeof(uint16(0))*2)))))))
					v17 = int16(int32(*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&i))), unsafe.Sizeof(uint16(0))*1))) + int32(v12))
					sub_499520(*(*int32)(unsafe.Pointer(&dword_5d4594_1522956)), (*uint16)(unsafe.Pointer(&i)), int16(uint16(v10)), 0, 0)
					sub_499520(*(*int32)(unsafe.Pointer(&dword_5d4594_1522956)), (*uint16)(unsafe.Pointer(&i)), int16(uint16(v10)), 1, 0)
					v11 = (*uint8)(unsafe.Add(unsafe.Pointer(v11), 408))
					*(*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(&v10))), unsafe.Sizeof(uint16(0))*0)) = uint16(v10 + 51)
					if int32(uintptr(unsafe.Pointer(v11))) >= int32(uintptr(memmap.PtrOff(0x587000, 194140))) {
						break
					}
				}
			}
		}
	}
	return 1
}
