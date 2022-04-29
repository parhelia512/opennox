package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	"github.com/noxworld-dev/opennox/v1/common/alloc"
	noxflags "github.com/noxworld-dev/opennox/v1/common/flags"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"image"
	"math"
	"unsafe"
)

func nox_thing_player_draw(a1 *uint32, dr *nox_drawable) int32 {
	var (
		v3  *byte
		v4  *uint32
		v5  *byte
		v6  *uint32
		v7  *uint32
		v8  *uint32
		v9  uint8
		v10 *int32
		v11 *uint32
		v12 int32
		v13 int32
		v14 int32
		v15 *uint32
		v16 int32
		v17 *int32
		v18 int32
		v20 int8
		v21 *int16
		v22 int32
		v23 *int32
		v24 int32
		v25 int32
		v26 int32
		v27 int32
		v28 int32
		v29 *byte
		v30 int32
		v31 int32
		v32 *byte
		v33 int32
		v34 int32
		v35 int32
		v36 [13]int32
	)
	v26 = int32(dr.field_32)
	v31 = 0
	v30 = 0
	v29 = nil
	v3 = (*byte)(unsafe.Pointer(noxServer.getPlayerByID(v26)))
	v32 = v3
	if v3 == nil {
		return 1
	}
	if *memmap.PtrUint32(0x8531A0, 2576) != 0 && int32(*(*uint8)(unsafe.Pointer(uintptr(*memmap.PtrUint32(0x8531A0, 2576) + 3680))))&1 != 0 {
		v30 = 1
	}
	if nox_player_netCode_85319C == dr.field_32 {
		v4 = nox_xxx_objGetTeamByNetCode_418C80(int32(nox_player_netCode_85319C))
		if v4 == nil {
			goto LABEL_15
		}
		v5 = (*byte)(unsafe.Pointer(nox_xxx_clientGetTeamColor_418AB0(int32(*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(v4))), 4)))))))
		goto LABEL_14
	}
	if *(*byte)(unsafe.Add(unsafe.Pointer(v3), 3680))&1 != 0 {
		return 1
	}
	v6 = nox_xxx_objGetTeamByNetCode_418C80(int32(nox_player_netCode_85319C))
	if v6 != nil {
		v7 = nox_xxx_objGetTeamByNetCode_418C80(int32(dr.field_32))
		v8 = v7
		if v7 != nil {
			if nox_xxx_servCompareTeams_419150(int32(uintptr(unsafe.Pointer(v6))), int32(uintptr(unsafe.Pointer(v7)))) != 0 {
				v31 = 1
			}
			v5 = (*byte)(unsafe.Pointer(nox_xxx_clientGetTeamColor_418AB0(int32(*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(v8))), 4)))))))
		LABEL_14:
			v29 = v5
			goto LABEL_15
		}
	}
LABEL_15:
	v9 = nox_xxx_getTeamCounter_417DD0()
	v10 = (*int32)(unsafe.Pointer(a1))
	v35 = bool2int(int32(v9) != 0)
	if (*(*byte)(unsafe.Add(unsafe.Pointer(v3), 4)) & 1) == 0 {
	LABEL_24:
		v12 = 0
		goto LABEL_25
	}
	v36[2] = nox_win_width
	v36[8] = nox_win_width
	v11 = *(**uint32)(unsafe.Pointer(&dword_5d4594_1313792))
	v12 = 0
	v36[0] = 0
	v36[1] = 0
	v36[3] = nox_win_height
	v36[9] = nox_win_height
	v36[4] = 0
	v36[5] = 0
	if dword_5d4594_1313792 != 0 || (func() bool {
		v13 = nox_xxx_getTTByNameSpriteMB_44CFC0("Flag")
		v11 = &nox_new_drawable_for_thing(v13).field_0
		return (func() uint32 {
			dword_5d4594_1313792 = uint32(uintptr(unsafe.Pointer(v11)))
			return dword_5d4594_1313792
		}()) != 0
	}()) {
		v14 = 0
		v15 = (*uint32)(unsafe.Add(unsafe.Pointer(v3), 2324))
		for *v15 != 1 {
			v14++
			v15 = (*uint32)(unsafe.Add(unsafe.Pointer(v15), unsafe.Sizeof(uint32(0))*6))
			if v14 >= 27 {
				goto LABEL_23
			}
		}
		libc.MemCpy(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer(v11), unsafe.Sizeof(uint32(0))*108))), unsafe.Add(unsafe.Pointer(v3), v14*24+2328), 20)
		v11 = *(**uint32)(unsafe.Pointer(&dword_5d4594_1313792))
		v3 = v32
	LABEL_23:
		*(*uint32)(unsafe.Add(unsafe.Pointer(v11), unsafe.Sizeof(uint32(0))*3)) = uint32(dr.pos.x - *(*int32)(unsafe.Add(unsafe.Pointer(v10), unsafe.Sizeof(int32(0))*4)) + *v10 + 15)
		*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1313792 + 16))) = uint32(*(*int32)(unsafe.Add(unsafe.Pointer(v10), unsafe.Sizeof(int32(0))*1)) - *(*int32)(unsafe.Add(unsafe.Pointer(v10), unsafe.Sizeof(int32(0))*5)) + dr.pos.y - 25)
		(cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1313792 + 300))), (*func(*int32, uint32))(nil)))(&v36[0], dword_5d4594_1313792)
		goto LABEL_24
	}
LABEL_25:
	if !nox_xxx_spriteCheckFlag31_4356C0(dr, 23) {
		if !nox_xxx_spriteCheckFlag31_4356C0(dr, 25) {
			*memmap.PtrUint32(0x973A20, 512) = *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v3))), unsafe.Sizeof(uint32(0))*574)))
			dword_5d4594_3798672 = *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v3))), unsafe.Sizeof(uint32(0))*576)))
			dword_5d4594_3798676 = *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v3))), unsafe.Sizeof(uint32(0))*578)))
			dword_5d4594_3798680 = *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v3))), unsafe.Sizeof(uint32(0))*577)))
			dword_5d4594_3798684 = *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v3))), unsafe.Sizeof(uint32(0))*575)))
			dword_5d4594_3798688 = *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v3))), unsafe.Sizeof(uint32(0))*573)))
			v16 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v3))), unsafe.Sizeof(uint32(0))*920))))
			if v16&1024 != 0 {
				v17 = memmap.PtrInt32(0x973A20, 512)
				for {
					{
						var a2 int32 = 0
						sub_434480(*v17, (*int32)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(&v33)))))), &a2, (*int32)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(&a1)))))))
						if int32(uint8(int8(a2))) >= 155 {
							*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a2))), 0)) = math.MaxUint8
						} else {
							*(*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(&a2))), 0)) = uint8(int8(a2 + 100))
						}
						*v17 = int32(nox_color_rgb_4344A0(v33, a2, int32(uintptr(unsafe.Pointer(a1)))))
						v17 = (*int32)(unsafe.Add(unsafe.Pointer(v17), unsafe.Sizeof(int32(0))*1))
					}
					if int32(uintptr(unsafe.Pointer(v17))) >= int32(uintptr(memmap.PtrOff(0x973A20, 536))) {
						break
					}
				}
				v12 = 0
			}
			goto LABEL_37
		}
		goto LABEL_29
	}
	if (int32(uint8(nox_frame_xxx_2598000)) & 1) == 0 {
	LABEL_29:
		*memmap.PtrUint32(0x973A20, 512) = nox_color_blue_2650684
		dword_5d4594_3798672 = nox_color_blue_2650684
		dword_5d4594_3798676 = nox_color_blue_2650684
		dword_5d4594_3798680 = nox_color_blue_2650684
		dword_5d4594_3798684 = nox_color_blue_2650684
		dword_5d4594_3798688 = nox_color_blue_2650684
		goto LABEL_37
	}
	*memmap.PtrUint32(0x973A20, 512) = nox_color_white_2523948
	dword_5d4594_3798672 = nox_color_white_2523948
	dword_5d4594_3798676 = nox_color_white_2523948
	dword_5d4594_3798680 = nox_color_white_2523948
	dword_5d4594_3798684 = nox_color_white_2523948
	dword_5d4594_3798688 = nox_color_white_2523948
LABEL_37:
	nox_xxx_drawPlayer_4341D0(1, *memmap.PtrInt32(0x973A20, 512))
	nox_xxx_drawPlayer_4341D0(2, *(*int32)(unsafe.Pointer(&dword_5d4594_3798672)))
	nox_xxx_drawPlayer_4341D0(3, *(*int32)(unsafe.Pointer(&dword_5d4594_3798676)))
	nox_xxx_drawPlayer_4341D0(4, *(*int32)(unsafe.Pointer(&dword_5d4594_3798680)))
	nox_xxx_drawPlayer_4341D0(5, *(*int32)(unsafe.Pointer(&dword_5d4594_3798684)))
	nox_xxx_drawPlayer_4341D0(6, *(*int32)(unsafe.Pointer(&dword_5d4594_3798688)))
	v18 = sub_4B8FA0(dr, &v28, &v27)
	if v18 == 0 {
		return 0
	}
	nox_xxx_drawObject_4C4770_draw(v10, dr, v18)
	v20 = int8(dr.field_74_2)
	if int32(v20) != 1 && int32(v20) != 0 && int32(v20) != 2 && int32(v20) != 3 && int32(v20) != 6 || dr.field_69 == 37 {
		sub_4B8D40(v10, dr, int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v3))), unsafe.Sizeof(uint32(0))*1)))&2), (*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v3))), unsafe.Sizeof(uint32(0))*581)), v28, v27)
		sub_4B8960(v10, dr, int32(*(*uint32)(unsafe.Pointer(v3))), (*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v3))), unsafe.Sizeof(uint32(0))*743)), v28, v27)
		sub_4B8D40(v10, dr, int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v3))), unsafe.Sizeof(uint32(0))*1)))&uint32(^int32(2))), (*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v3))), unsafe.Sizeof(uint32(0))*581)), v28, v27)
	} else {
		sub_4B8D40(v10, dr, int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v3))), unsafe.Sizeof(uint32(0))*1)))&uint32(^int32(2))), (*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v3))), unsafe.Sizeof(uint32(0))*581)), v28, v27)
		sub_4B8960(v10, dr, int32(*(*uint32)(unsafe.Pointer(v3))), (*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v3))), unsafe.Sizeof(uint32(0))*743)), v28, v27)
		sub_4B8D40(v10, dr, int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v3))), unsafe.Sizeof(uint32(0))*1)))&2), (*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v3))), unsafe.Sizeof(uint32(0))*581)), v28, v27)
	}
	if v30 != 0 || !nox_xxx_spriteCheckFlag31_4356C0(dr, 0) || dr.field_32 == nox_player_netCode_85319C || *memmap.PtrUint32(0x852978, 8) != 0 && (nox_xxx_spriteCheckFlag31_4356C0((*nox_drawable)(unsafe.Pointer(uintptr(*memmap.PtrInt32(0x852978, 8)))), 21) || v31 != 0) {
		a1 = (*uint32)(unsafe.Pointer(uintptr(nox_color_rgb_4344A0(155, 155, 155))))
		if sub_48D830(dr) == 0 && !noxflags.HasGame(noxflags.GameModeCoop) {
			v21 = (*int16)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer(v3), 4704))))
			nox_xxx_drawGetStringSize_43F840(nil, (*libc.WChar)(unsafe.Pointer((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v3))), unsafe.Sizeof(uint16(0))*2352)))), &v34, nil, 0)
			v22 = *v10 + dr.pos.x + v34/(-2) - *(*int32)(unsafe.Add(unsafe.Pointer(v10), unsafe.Sizeof(int32(0))*4))
			var a2 int32 = *(*int32)(unsafe.Add(unsafe.Pointer(v10), unsafe.Sizeof(int32(0))*1)) - *(*int32)(unsafe.Add(unsafe.Pointer(v10), unsafe.Sizeof(int32(0))*5)) + dr.pos.y - 64
			nox_xxx_drawSetTextColor_434390(*memmap.PtrInt32(0x852978, 4))
			noxrend.DrawString(nil, (*libc.WChar)(unsafe.Pointer(v21)), image.Pt(v22+1, a2+1))
			nox_xxx_drawSetTextColor_434390(int32(uintptr(unsafe.Pointer(a1))))
			if v35 != 0 {
				if v29 != nil {
					v23 = (*int32)(unsafe.Pointer(nox_xxx_materialGetTeamColor_418D50((*nox_team_t)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v29)))))))))
					if v23 != nil {
						nox_xxx_drawSetTextColor_434390(*v23)
					}
				}
			}
			noxrend.DrawString(nil, (*libc.WChar)(unsafe.Pointer(v21)), image.Pt(v22, a2))
			v3 = v32
			v12 = 0
		}
		if nox_xxx_spriteCheckFlag31_4356C0(dr, 16) {
			v36[2] = nox_win_width
			v36[8] = nox_win_width
			v36[0] = 0
			v36[1] = 0
			v36[3] = nox_win_height
			v36[9] = nox_win_height
			v36[4] = 0
			v36[5] = 0
			if dword_5d4594_1313796 == 0 {
				v24 = nox_xxx_getTTByNameSpriteMB_44CFC0("SpinningSkull")
				dword_5d4594_1313796 = uint32(uintptr(unsafe.Pointer(nox_new_drawable_for_thing(v24))))
				*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1313796 + 120))) |= 0x1000000
			}
			*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1313796 + 12))) = uint32(*v10 + dr.pos.x - *(*int32)(unsafe.Add(unsafe.Pointer(v10), unsafe.Sizeof(int32(0))*4)))
			*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1313796 + 16))) = uint32(*(*int32)(unsafe.Add(unsafe.Pointer(v10), unsafe.Sizeof(int32(0))*1)) - *(*int32)(unsafe.Add(unsafe.Pointer(v10), unsafe.Sizeof(int32(0))*5)) + dr.pos.y - 50)
			(cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1313796 + 300))), (*func(*int32, uint32))(nil)))(&v36[0], dword_5d4594_1313796)
		}
		if nox_xxx_spriteCheckFlag31_4356C0(dr, 30) {
			v36[2] = nox_win_width
			v36[8] = nox_win_width
			v36[0] = 0
			v36[1] = 0
			v36[3] = nox_win_height
			v36[9] = nox_win_height
			v36[4] = 0
			v36[5] = 0
			if dword_5d4594_1313800 == 0 {
				v25 = nox_xxx_getTTByNameSpriteMB_44CFC0("SpinningCrown")
				dword_5d4594_1313800 = uint32(uintptr(unsafe.Pointer(nox_new_drawable_for_thing(v25))))
				*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1313800 + 120))) |= 0x1000000
			}
			*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1313800 + 12))) = uint32(*v10 + dr.pos.x - *(*int32)(unsafe.Add(unsafe.Pointer(v10), unsafe.Sizeof(int32(0))*4)))
			*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1313800 + 16))) = uint32(*(*int32)(unsafe.Add(unsafe.Pointer(v10), unsafe.Sizeof(int32(0))*1)) - *(*int32)(unsafe.Add(unsafe.Pointer(v10), unsafe.Sizeof(int32(0))*5)) + dr.pos.y - 50)
			(cgoAsFunc(*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1313800 + 300))), (*func(*int32, uint32))(nil)))(&v36[0], dword_5d4594_1313800)
		}
		for {
			nox_xxx_drawPlayer_4341D0(func() int32 {
				p := &v12
				x := *p
				*p++
				return x
			}(), int32(nox_color_white_2523948))
			if v12 >= 6 {
				break
			}
		}
		if unsafe.Pointer(dr) != unsafe.Pointer(uintptr(*memmap.PtrUint32(0x852978, 8))) && nox_xxx_playerGet_470A90() != 0 {
			if noxflags.HasGame(noxflags.GameModeQuest) {
				nox_xxx_drawOtherPlayerHP_4B8EB0((*uint32)(unsafe.Pointer(v10)), dr, uint16(uint8(*(*byte)(unsafe.Add(unsafe.Pointer(v3), 2282)))), int8(uint8((*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v3))), unsafe.Sizeof(uint32(0))*920)))>>10)&1)))
			}
		}
	}
	return 1
}
func nox_thing_player_waypoint_draw(a1 int32, dr *nox_drawable) int32 {
	var (
		v2 int32
		v3 int32
		v4 int32
		v5 int32
		v7 int32
		a2 int32 = int32(uintptr(unsafe.Pointer(dr)))
	)
	v2 = int32(*memmap.PtrUint32(0x85B3FC, 940))
	nox_xxx_spriteDrawCircleMB_4C32A0(int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 12)))-*(*uint32)(unsafe.Pointer(uintptr(a1 + 16)))), int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 16)))-*(*uint32)(unsafe.Pointer(uintptr(a1 + 20)))), 10, *memmap.PtrInt32(0x85B3FC, 940))
	v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 16))) - *(*uint32)(unsafe.Pointer(uintptr(a1 + 20))))
	v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 12))) - *(*uint32)(unsafe.Pointer(uintptr(a1 + 16))))
	v5 = int32(uint8(int8(int32(uint8(nox_frame_xxx_2598000)) * 2)))
	nox_client_drawEnableAlpha_434560(1)
	nox_client_drawSetColor_434460(v2)
	v7 = 5
	for {
		nox_client_drawAddPoint_49F500(v4+*memmap.PtrInt32(0x587000, uint32(v5*8)+0x2EE58)*10/16, v3+*memmap.PtrInt32(0x587000, uint32(v5*8)+0x2EE5C)*10/16)
		nox_client_drawAddPoint_49F500(v4+*memmap.PtrInt32(0x587000, uint32(((v5+102)%256)*8)+0x2EE58)*10/16, v3+*memmap.PtrInt32(0x587000, uint32(((v5+102)%256)*8)+0x2EE5C)*10/16)
		nox_client_drawLineFromPoints_49E4B0()
		v5 = (v5 + 102) % 256
		v7--
		if v7 == 0 {
			break
		}
	}
	nox_client_drawEnableAlpha_434560(0)
	return 1
}
func nox_things_player_draw_parse(obj *nox_thing, f *nox_memfile, attr_value *byte) bool {
	var (
		a3  *uint8 = (*uint8)(unsafe.Pointer(attr_value))
		v3  *uint32
		v5  *uint32
		v7  int32
		v10 int32
		v11 int32
		v13 int32
		v14 *uint32
		v15 int32
		v16 int32
		v17 int32
		v18 *uint32
		v19 int32
		v20 *uint32
		v21 *uint32
		v22 uint8
		v23 uint8
	)
	v3 = (*uint32)(alloc.Calloc(1, 0x38BC))
	v5 = v3
	v21 = v3
	*v3 = 0x38BC
	v7 = int32(f.ReadU32())
	if uint32(v7) == 0x454E4420 {
		return false
	}
LABEL_3:
	v22 = f.ReadU8()
	nox_memfile_read(unsafe.Pointer(a3), 1, int32(v22), f)
	*((*uint8)(unsafe.Add(unsafe.Pointer(a3), v22))) = 0
	v10 = sub_44BB20((*byte)(unsafe.Pointer(a3)))
	if v10 < 0 {
		return false
	}
	v11 = int32(uintptr(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(uint32(0))*uintptr(v10*66+1))))))
	if nox_xxx_loadVectorAnimated_44B8B0(v11, f) == 0 {
		return false
	}
	for {
		v13 = int32(f.ReadU32())
		if uint32(v13) == 0x454E4420 {
			break
		}
		if uint32(v13) == 0x53544154 {
			v5 = v21
			goto LABEL_3
		}
		v23 = f.ReadU8()
		nox_memfile_read(unsafe.Pointer(a3), 1, int32(v23), f)
		*((*uint8)(unsafe.Add(unsafe.Pointer(a3), v23))) = 0
		if libc.StrCmp(CString("NAKED"), (*byte)(unsafe.Pointer(a3))) == 0 {
			v14 = (*uint32)(alloc.Calloc(1, 40))
			v15 = int32(*(*int16)(unsafe.Pointer(uintptr(v11 + 40))))
			*(*uint32)(unsafe.Pointer(uintptr(v11 + 48))) = uint32(uintptr(unsafe.Pointer(v14)))
			v16 = sub_44B940(v14, v15, f)
		} else {
			v17 = nox_xxx_parse_Armor_44BA60((*byte)(unsafe.Pointer(a3)))
			if v17 < 0 {
				v19 = sub_44BAC0((*byte)(unsafe.Pointer(a3)))
				if v19 < 0 {
					return false
				}
				v20 = (*uint32)(alloc.Calloc(1, 40))
				*(*uint32)(unsafe.Pointer(uintptr(v11 + v19*4 + 156))) = uint32(uintptr(unsafe.Pointer(v20)))
				v16 = sub_44B940(v20, int32(*(*int16)(unsafe.Pointer(uintptr(v11 + 40)))), f)
			} else {
				v18 = (*uint32)(alloc.Calloc(1, 40))
				*(*uint32)(unsafe.Pointer(uintptr(v11 + v17*4 + 52))) = uint32(uintptr(unsafe.Pointer(v18)))
				v16 = sub_44B940(v18, int32(*(*int16)(unsafe.Pointer(uintptr(v11 + 40)))), f)
			}
		}
		if v16 == 0 {
			return false
		}
	}
	obj.draw_func = nox_thing_player_draw
	obj.field_5c = unsafe.Pointer(v21)
	return true
}
