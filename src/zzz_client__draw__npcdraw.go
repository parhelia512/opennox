package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	noxflags "github.com/noxworld-dev/opennox/v1/common/flags"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"math"
	"unsafe"
)

func nox_thing_npc_draw(a1 *int32, dr *nox_drawable) int32 {
	var (
		v2  int32
		v4  int32
		v5  *int32
		v6  int32
		v7  int32
		v8  int32
		v9  int32
		v10 int32
		v11 int32
		v12 int32
		v13 *int32
		v14 int32
		v15 int32
		v16 *int32
		v17 int32
		v18 int32
		v19 int32
		v20 int32
		v21 *int32
		v22 int8
		v23 int32
		v24 *func(*byte, int32)
		v25 *byte
		v26 int32
		v27 *func(*byte, int32)
		v28 *byte
		i   int32
		v30 int32
		v31 int32
		v32 int32
		v33 *int32
		v34 int32
		v35 int32
		v36 int32
		a2  int32 = int32(uintptr(unsafe.Pointer(dr)))
	)
	if noxflags.HasGame(noxflags.GameFlag22) {
		v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(sub_44D330((*byte)(memmap.PtrOff(0x587000, 177440))) + 92))) + 4)
		if int32(*(*uint16)(unsafe.Pointer(uintptr(v2 + 40)))) != 0 {
			nox_xxx_drawObject_4C4770_draw(a1, (*nox_drawable)(unsafe.Pointer(uintptr(a2))), int32(**(**uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 48))) + uint32(int32(*(*uint8)(unsafe.Pointer(uintptr(a2 + 297))))*4) + 4)))))
			return 1
		}
	} else {
		v4 = a2
		v30 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 128))))
		v36 = int32(*(*uint32)(unsafe.Pointer(uintptr(*memmap.PtrUint32(0x852978, 8) + 304))))
		v5 = (*int32)(unsafe.Pointer(nox_npc_by_id(v30)))
		v33 = v5
		if v36 != 0 && v5 != nil {
			if nox_xxx_spriteCheckFlag31_4356C0((*nox_drawable)(unsafe.Pointer(uintptr(v4))), 23) {
				if int32(uint8(nox_frame_xxx_2598000))&1 != 0 {
					v6 = 0
					for {
						v7 = v6 + 1
						nox_xxx_drawPlayer_4341D0(v6+1, int32(nox_color_white_2523948))
						v6 = v7
						if v7 >= 6 {
							break
						}
					}
				} else {
					v8 = 0
					for {
						v9 = v8 + 1
						nox_xxx_drawPlayer_4341D0(v8+1, int32(nox_color_blue_2650684))
						v8 = v9
						if v9 >= 6 {
							break
						}
					}
				}
			} else if nox_xxx_spriteCheckFlag31_4356C0((*nox_drawable)(unsafe.Pointer(uintptr(v4))), 25) {
				v10 = 0
				for {
					v11 = v10 + 1
					nox_xxx_drawPlayer_4341D0(v10+1, int32(nox_color_blue_2650684))
					v10 = v11
					if v11 >= 6 {
						break
					}
				}
			} else if *(*int32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(int32(0))*328)) == 1 {
				v12 = 0
				v13 = (*int32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(int32(0))*2))
				for {
					sub_434480(*v13, (*int32)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(&v35)))))), (*int32)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(&a2)))))), (*int32)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(&v34)))))))
					if int32(uint8(int8(a2))) >= 155 {
						*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a2))), 0)) = math.MaxUint8
					} else {
						*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a2))), 0)) = uint8(int8(a2 + 100))
					}
					v12++
					v14 = int32(nox_color_rgb_4344A0(v35, a2, v34))
					nox_xxx_drawPlayer_4341D0(v12, v14)
					v13 = (*int32)(unsafe.Add(unsafe.Pointer(v13), unsafe.Sizeof(int32(0))*1))
					if v12 >= 6 {
						break
					}
				}
			} else {
				v15 = 0
				v16 = (*int32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(int32(0))*2))
				for {
					v17 = v15 + 1
					nox_xxx_drawPlayer_4341D0(v15+1, *v16)
					v15 = v17
					v16 = (*int32)(unsafe.Add(unsafe.Pointer(v16), unsafe.Sizeof(int32(0))*1))
					if v17 >= 6 {
						break
					}
				}
			}
			v18 = nox_xxx_spriteNPCInfo_49A4B0((*uint32)(unsafe.Pointer(uintptr(v4))), *(*int32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(int32(0))*326)), *(*int32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(int32(0))*327)))
			v19 = v36 + v18*264 + 4
			if int32(*(*uint16)(unsafe.Pointer(uintptr(v36 + v18*264 + 44)))) != 0 {
				v20 = sub_4BC5D0((*nox_drawable)(unsafe.Pointer(uintptr(v4))), v36+v18*264+4)
				if v20 < 0 {
					return 0
				}
				v21 = a1
				nox_xxx_drawObject_4C4770_draw(a1, (*nox_drawable)(unsafe.Pointer(uintptr(v4))), int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v19 + 48))) + uint32(int32(*(*uint8)(unsafe.Pointer(uintptr(v4 + 297))))*4) + 4))) + uint32(v20*4))))))
				v22 = int8(*(*uint8)(unsafe.Pointer(uintptr(v4 + 297))))
				if int32(v22) != 1 && int32(v22) != 0 && int32(v22) != 2 && int32(v22) != 3 && int32(v22) != 6 || *(*uint32)(unsafe.Pointer(uintptr(v4 + 276))) == 37 {
					sub_4B8960(v21, (*nox_drawable)(unsafe.Pointer(uintptr(v4))), *(*int32)(unsafe.Add(unsafe.Pointer(v33), unsafe.Sizeof(int32(0))*327)), (*uint32)(unsafe.Pointer((*int32)(unsafe.Add(unsafe.Pointer(v33), unsafe.Sizeof(int32(0))*170)))), v19, v20)
					sub_4B8D40(v21, (*nox_drawable)(unsafe.Pointer(uintptr(v4))), *(*int32)(unsafe.Add(unsafe.Pointer(v33), unsafe.Sizeof(int32(0))*326)), (*uint32)(unsafe.Pointer((*int32)(unsafe.Add(unsafe.Pointer(v33), unsafe.Sizeof(int32(0))*8)))), v19, v20)
				} else {
					sub_4B8D40(v21, (*nox_drawable)(unsafe.Pointer(uintptr(v4))), *(*int32)(unsafe.Add(unsafe.Pointer(v33), unsafe.Sizeof(int32(0))*326)), (*uint32)(unsafe.Pointer((*int32)(unsafe.Add(unsafe.Pointer(v33), unsafe.Sizeof(int32(0))*8)))), v19, v20)
					sub_4B8960(v21, (*nox_drawable)(unsafe.Pointer(uintptr(v4))), *(*int32)(unsafe.Add(unsafe.Pointer(v33), unsafe.Sizeof(int32(0))*327)), (*uint32)(unsafe.Pointer((*int32)(unsafe.Add(unsafe.Pointer(v33), unsafe.Sizeof(int32(0))*170)))), v19, v20)
				}
				if nox_xxx_spriteCheckFlag31_4356C0((*nox_drawable)(unsafe.Pointer(uintptr(v4))), 16) {
					if dword_5d4594_1313796 == 0 {
						v23 = nox_xxx_getTTByNameSpriteMB_44CFC0("SpinningSkull")
						dword_5d4594_1313796 = uint32(uintptr(unsafe.Pointer(nox_new_drawable_for_thing(v23))))
						*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1313796 + 120))) |= 0x1000000
					}
					*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1313796 + 12))) = uint32(*v21) + *(*uint32)(unsafe.Pointer(uintptr(v4 + 12))) - uint32(*(*int32)(unsafe.Add(unsafe.Pointer(v21), unsafe.Sizeof(int32(0))*4)))
					*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1313796 + 16))) = uint32(*(*int32)(unsafe.Add(unsafe.Pointer(v21), unsafe.Sizeof(int32(0))*1))-*(*int32)(unsafe.Add(unsafe.Pointer(v21), unsafe.Sizeof(int32(0))*5))) + *(*uint32)(unsafe.Pointer(uintptr(v4 + 16))) - 50
					v31 = int32(dword_5d4594_1313796)
					v24 = (*func(*byte, int32))(unsafe.Pointer(uintptr(dword_5d4594_1313796 + 300)))
					v25 = (*byte)(unsafe.Pointer(nox_draw_getViewport_437250()))
					(*v24)(v25, v31)
				}
				if nox_xxx_spriteCheckFlag31_4356C0((*nox_drawable)(unsafe.Pointer(uintptr(v4))), 30) {
					if dword_5d4594_1313800 == 0 {
						v26 = nox_xxx_getTTByNameSpriteMB_44CFC0("SpinningCrown")
						dword_5d4594_1313800 = uint32(uintptr(unsafe.Pointer(nox_new_drawable_for_thing(v26))))
						*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1313800 + 120))) |= 0x1000000
					}
					*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1313800 + 12))) = uint32(*v21) + *(*uint32)(unsafe.Pointer(uintptr(v4 + 12))) - uint32(*(*int32)(unsafe.Add(unsafe.Pointer(v21), unsafe.Sizeof(int32(0))*4)))
					*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1313800 + 16))) = uint32(*(*int32)(unsafe.Add(unsafe.Pointer(v21), unsafe.Sizeof(int32(0))*1))-*(*int32)(unsafe.Add(unsafe.Pointer(v21), unsafe.Sizeof(int32(0))*5))) + *(*uint32)(unsafe.Pointer(uintptr(v4 + 16))) - 50
					v32 = int32(dword_5d4594_1313800)
					v27 = (*func(*byte, int32))(unsafe.Pointer(uintptr(dword_5d4594_1313800 + 300)))
					v28 = (*byte)(unsafe.Pointer(nox_draw_getViewport_437250()))
					(*v27)(v28, v32)
				}
				if nox_xxx_unitSpriteCheckAlly_4951F0(int32(*(*uint32)(unsafe.Pointer(uintptr(v4 + 128))))) != 0 {
					v33 = nil
					a1 = nil
					*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a2))), 0)) = 0
					sub_495180(int32(*(*uint32)(unsafe.Pointer(uintptr(v4 + 128)))), (*uint16)(unsafe.Pointer(&v33)), (*uint16)(unsafe.Pointer(&a1)), (*uint8)(unsafe.Pointer(&a2)))
					nox_xxx_spriteDrawMonsterHP_4BC080((*uint32)(unsafe.Pointer(v21)), v4, uint16(uintptr(unsafe.Pointer(v33))), uint16(uintptr(unsafe.Pointer(a1))), int8(a2))
				}
				for i = 0; i < 6; i++ {
					nox_xxx_drawPlayer_4341D0(i, int32(nox_color_white_2523948))
				}
			}
		}
	}
	return 1
}
