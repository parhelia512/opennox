package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	"unsafe"
)

func nox_client_screenParticleDraw_489700(a1p unsafe.Pointer, p *nox_screenParticle) int32 {
	var (
		a1    int32 = int32(uintptr(a1p))
		v7    int32
		v10   int32
		v11   int8
		v12   int8
		xLeft int2
	)
	xLeft.field_0 = int32(p.field_24 >> 16)
	xLeft.field_4 = int32(p.field_28 >> 16)
	if xLeft.field_0 <= 0 || xLeft.field_4 <= 0 || xLeft.field_0 >= *(*int32)(unsafe.Pointer(uintptr(a1 + 32))) || xLeft.field_4 >= *(*int32)(unsafe.Pointer(uintptr(a1 + 36))) {
		sub_431700((*uint64)(unsafe.Pointer(p)))
		return 0
	}
	sub_4B6720(&xLeft, int32(p.field_8), int32(p.field_40[0]), int8(p.field_40[0]))
	nox_client_drawSetColor_434460(int32(p.field_12))
	nox_xxx_drawPointMB_499B70(xLeft.field_0, xLeft.field_4, int32(p.field_40[0])>>1)
	var v3 int8 = int8(p.field_40[1])
	p.field_20 += p.field_36
	if int32(v3) != 0 {
		var v4 int8 = int8(int32(v3) - 1)
		p.field_40[1] = uint8(v4)
		if int32(v4) == 0 {
			if int32(p.field_40[2]) == 1 {
				var v5 uint8 = uint8(int8(int32(p.field_40[0]) + 1))
				p.field_40[0] = v5
				if int32(v5) >= 4 {
					p.field_40[2] = 2
				}
			} else {
				var v6 int8 = int8(int32(p.field_40[0]) - 1)
				p.field_40[0] = uint8(v6)
				if int32(v6) == 0 {
					sub_431700((*uint64)(unsafe.Pointer(p)))
					return 0
				}
			}
			p.field_40[1] = p.field_40[3]
		}
	}
	if int32(*(*uint8)(unsafe.Pointer(&p.field_32))) == 1 && randomIntMinMax(0, 10) >= 8 {
		v12 = int8(randomIntMinMax(3, 5))
		v11 = int8(randomIntMinMax(2, 3))
		v10 = xLeft.field_4
		v7 = randomIntMinMax(-2, 2)
		nox_client_newScreenParticle_431540(int32(p.field_4), xLeft.field_0+v7, v10, 0, 0, 1, v11, v12, 2, 2)
	}
	p.field_24 += p.field_16
	p.field_28 += p.field_20
	return 1
}
