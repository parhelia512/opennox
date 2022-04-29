package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"unsafe"
)

type nox_window_flags int32

const (
	NOX_WIN_HIDDEN      nox_window_flags = 16
	NOX_WIN_LAYER_FRONT nox_window_flags = 32
	NOX_WIN_LAYER_BACK  nox_window_flags = 64
)

type nox_window struct {
	id           int32
	flags        nox_window_flags
	width        int32
	height       int32
	off_x        int32
	off_y        int32
	end_x        int32
	end_y        int32
	widget_data  unsafe.Pointer
	draw_data    nox_window_data
	field_92     uint32
	field_93     func(*nox_window, int32, int32, int32) int32
	field_94     func(*nox_window, int32, int32, int32) int32
	draw_func    func(*nox_window, *nox_window_data) int32
	tooltip_func func(*nox_window, *nox_window_data, int32) int32
	prev         *nox_window
	next         *nox_window
	parent       *nox_window
	field_100    *nox_window
}
type nox_window_data struct {
	field_0    uint32
	group      int32
	style      int32
	status     int32
	win        *nox_window
	bg_color   uint32
	bg_image   unsafe.Pointer
	en_color   uint32
	en_image   unsafe.Pointer
	hl_color   uint32
	hl_image   unsafe.Pointer
	dis_color  uint32
	dis_image  unsafe.Pointer
	sel_color  uint32
	sel_image  unsafe.Pointer
	img_px     int32
	img_py     int32
	text_color uint32
	text       [64]libc.WChar
	font       unsafe.Pointer
	tooltip    [64]libc.WChar
}
type nox_window_ref struct {
	win  *nox_window
	next *nox_window_ref
}

var dword_5d4594_3799524 int32 = 0
var dword_5d4594_1309696 uint32 = 0
var dword_5d4594_1309704 uint32 = 0
var nox_win_1064912 *nox_window_ref = nil

func sub_46ACE0(a1 *uint32, a2 int32, a3 int32, a4 int32) {
	for i := int32(a2); i <= a3; i++ {
		var v5 *uint32 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(a1)).ChildByID(i)))
		(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v5)))))).SetHidden(a4)
	}
}
func sub_46AD20(a1 *uint32, a2 int32, a3 int32, a4 int32) {
	var (
		i  int32
		v5 *uint32
	)
	for i = a2; i <= a3; i++ {
		v5 = (*uint32)(unsafe.Pointer((*nox_window)(unsafe.Pointer(a1)).ChildByID(i)))
		nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v5)))))), a4)
	}
}
func nox_gui_getWindowOffs_46AA20(win *nox_window, px *uint32, py *uint32) int32 {
	if win == nil {
		*px = 0
		*py = 0
		return -2
	}
	*px = uint32(win.off_x)
	*py = uint32(win.off_y)
	return 0
}
func nox_client_wndGetPosition_46AA60(win *nox_window, px *uint32, py *uint32) int32 {
	if win == nil {
		return -2
	}
	*px = uint32(win.off_x)
	*py = uint32(win.off_y)
	for i := (*nox_window)(win.parent); i != nil; i = i.parent {
		*px += uint32(i.off_x)
		*py += uint32(i.off_y)
	}
	return 0
}
func nox_xxx_wndPointInWnd_46AAB0(a1 *uint32, a2 int32, a3 int32) bool {
	var (
		v3 *uint32
		v5 int32
		v6 int32 = 0
		v7 int32 = 0
	)
	v3 = a1
	nox_client_wndGetPosition_46AA60((*nox_window)(unsafe.Pointer(a1)), (*uint32)(unsafe.Pointer(&a1)), (*uint32)(unsafe.Pointer(&v6)))
	nox_window_get_size((*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v3)))))), &v5, &v7)
	return a2 >= int32(uintptr(unsafe.Pointer(a1))) && a2 <= int32(uintptr(unsafe.Pointer(a1)))+v5 && a3 >= v6 && a3 <= v6+v7
}
func sub_46AB20(a1 *uint32, a2 int32, a3 int32) int32 {
	var v4 int32
	if a1 == nil {
		return -2
	}
	*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*6)) = uint32(a2) + *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*4))
	v4 = int32(uint32(a3) + *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*5)))
	*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*2)) = uint32(a2)
	*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*3)) = uint32(a3)
	*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*7)) = uint32(v4)
	(*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1)))))).Func94(asWindowEvent(0x4004, a2, a3))
	return 0
}
func nox_window_get_size(win *nox_window, outW *int32, outH *int32) int32 {
	if win == nil {
		*outW = 0
		*outH = 0
		return -2
	}
	*outW = win.width
	*outH = win.height
	return 0
}
func nox_xxx_wnd_46ABB0(win *nox_window, a2 int32) int32 {
	var (
		a1 int32 = int32(uintptr(unsafe.Pointer(win)))
		v3 int32
		v4 uint32
		v5 int32
	)
	if a1 == 0 {
		return -2
	}
	v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 4))))
	if a2 != 0 {
		v4 = uint32(v3 | 8)
	} else {
		v4 = uint32(v3) & 0xFFFFFFF7
	}
	v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 400))))
	for *(*uint32)(unsafe.Pointer(uintptr(a1 + 4))) = v4; v5 != 0; v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(v5 + 388)))) {
		nox_xxx_wnd_46ABB0((*nox_window)(unsafe.Pointer(uintptr(v5))), a2)
	}
	return 0
}
func sub_4A9E90(a1 int32, a2 int32, a3 int32, a4 int32) {
	var (
		v4  int32
		v5  int32
		v6  int32
		v7  int32
		v8  int32
		v9  int32
		v10 int32
		v11 int32
	)
	_ = v11
	var v12 int32
	var v15 int32
	var v16 int32
	var v18 int32
	var v19 int32
	var v20 int32
	var v21 int32
	v4 = a1 + a3
	v18 = a1
	v5 = a1 + 10
	v6 = a2 - 10
	v7 = a1 + a3 - 30
	v19 = a1 + a3
	v21 = a2 - 10
	v8 = a2 + a4 - 10
	if v5 <= v7 {
		for {
			nox_client_drawImageAt_47D2C0(unsafe.Pointer(uintptr(*memmap.PtrInt32(6112660, 1309692))), v5, v6)
			nox_client_drawImageAt_47D2C0(unsafe.Pointer(uintptr(*memmap.PtrInt32(6112660, 1309692))), v5, v8)
			v5 += 20
			if v5 > v7 {
				break
			}
		}
		v4 = v19
	}
	v9 = v4 - 10
	if v4-10-v5 >= 10 {
		nox_client_drawImageAt_47D2C0(unsafe.Pointer(uintptr(dword_5d4594_1309696)), v5, v6)
		nox_client_drawImageAt_47D2C0(unsafe.Pointer(uintptr(dword_5d4594_1309696)), v5, v8)
		v5 += 10
	}
	if v5 < v9 {
		v10 = int32(uint32(v5) + (uint32(v9-v5+1) & 0xFFFFFFFE) - 10)
		nox_client_drawImageAt_47D2C0(unsafe.Pointer(uintptr(dword_5d4594_1309696)), v10, v6)
		nox_client_drawImageAt_47D2C0(unsafe.Pointer(uintptr(dword_5d4594_1309696)), v10, v8)
	}
	v11 = a2 + a4 - 30
	v12 = a2 + 10
	v15 = v18 - 10
	v20 = a2 + a4 - 30
	for v12 <= v20 {
		nox_client_drawImageAt_47D2C0(unsafe.Pointer(uintptr(*memmap.PtrInt32(6112660, 1309700))), v15, v12)
		nox_client_drawImageAt_47D2C0(unsafe.Pointer(uintptr(*memmap.PtrInt32(6112660, 1309700))), v9, v12)
		v12 += 20
	}
	if v8-v12 >= 10 {
		nox_client_drawImageAt_47D2C0(unsafe.Pointer(uintptr(dword_5d4594_1309704)), v15, v12)
		nox_client_drawImageAt_47D2C0(unsafe.Pointer(uintptr(dword_5d4594_1309704)), v9, v12)
		v12 += 10
	}
	if v12 < v8 {
		v16 = int32(uint32(v12) + (uint32(v8-v12+1) & 0xFFFFFFFE) - 10)
		nox_client_drawImageAt_47D2C0(unsafe.Pointer(uintptr(dword_5d4594_1309704)), v15, v16)
		nox_client_drawImageAt_47D2C0(unsafe.Pointer(uintptr(dword_5d4594_1309704)), v9, v16)
	}
	nox_client_drawImageAt_47D2C0(unsafe.Pointer(uintptr(*memmap.PtrInt32(6112660, 1309676))), v15, v21)
	nox_client_drawImageAt_47D2C0(unsafe.Pointer(uintptr(*memmap.PtrInt32(6112660, 1309680))), v9, v21)
	nox_client_drawImageAt_47D2C0(unsafe.Pointer(uintptr(*memmap.PtrInt32(6112660, 1309684))), v15, v8)
	nox_client_drawImageAt_47D2C0(unsafe.Pointer(uintptr(*memmap.PtrInt32(6112660, 1309688))), v9, v8)
}
func sub_4AA030(win *nox_window, data *nox_window_data) {
	var (
		a1  *uint32 = (*uint32)(unsafe.Pointer(win))
		a2  int32   = int32(uintptr(unsafe.Pointer(data)))
		v2  *uint32
		v3  int32
		v4  int32
		v5  uint32
		v6  int32
		v7  int32
		v8  int32
		v9  *uint32
		v10 int32
		v11 bool
		v12 int32
		v13 *uint32
		v14 int16
		v16 int32 = 0
		v17 uint32
		v18 int32
		v19 int32
	)
	v2 = a1
	v3 = 0
	nox_client_wndGetPosition_46AA60((*nox_window)(unsafe.Pointer(a1)), (*uint32)(unsafe.Pointer(&a1)), (*uint32)(unsafe.Pointer(&v16)))
	v5 = 0
	v17 = 0
	for v3 == 0 {
		v4 = int32((1 << v5) & *(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*11)))
		if v4 != 0 {
			if v4 <= 32 {
				if v4 != 32 {
					switch func() int32 {
						p := &v4
						*p--
						return *p
					}() {
					case 0:
						fallthrough
					case 1:
						goto LABEL_24
					case 3:
						fallthrough
					case 7:
						fallthrough
					case 15:
						goto LABEL_25
					default:
						goto LABEL_26
					}
					goto LABEL_26
				}
				v6 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*8)))
				v7 = 0
				v8 = 0
				if *(*uint32)(unsafe.Pointer(uintptr(v6 + 12))) != 0 {
					v7 = int32(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v6 + 36))) + 400))) + 12))))
				}
				if int32(*((*uint16)(unsafe.Add(unsafe.Pointer((*uint16)(unsafe.Pointer(v2))), unsafe.Sizeof(uint16(0))*54)))) != 0 {
					v8 = 4
				}
				sub_4A9E90(int32(uintptr(unsafe.Pointer(a1)))-3, v16-v8-3, int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*2))-uint32(v7)+3), int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*3))+6))
				goto LABEL_25
			}
			if v4 > 4096 {
				if v4 == 8192 {
					goto LABEL_24
				}
			} else {
				switch v4 {
				case 4096:
					goto LABEL_24
				case 128:
					v9 = a1
					v10 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*8)))
					v11 = int32(*(*uint16)(unsafe.Pointer(uintptr(a2 + 72)))) == 0
					v12 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*2)))
					v19 = v16
					v13 = a1
					if !v11 {
						nox_xxx_drawGetStringSize_43F840(int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 200)))), (*uint16)(unsafe.Pointer(uintptr(a2+72))), &v18, nil, 0)
						v13 = (*uint32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(v13))), v18))), 6))))
						v9 = a1
						v12 += int32(-6 - v18)
					}
					v14 = int16(*(*uint16)(unsafe.Pointer(uintptr(v10 + 1042))))
					if int32(v14) > 0 && v12 > int32(v14) {
						v12 = int32(v14)
						v13 = (*uint32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(v9))), *(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*2))))), -v14))))
					}
					sub_4A9E90(int32(uintptr(unsafe.Pointer(v13))), v19, v12, int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*3))))
					v5 = v17
					v3 = 1
				case 2048:
				LABEL_24:
					sub_4A9E90(int32(uintptr(unsafe.Pointer(a1))), v16, int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*2))), int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(uint32(0))*3))))
				LABEL_25:
					v3 = 1
				}
			}
		}
	LABEL_26:
		v17 = func() uint32 {
			p := &v5
			*p++
			return *p
		}()
		if v5 >= 32 {
			break
		}
	}
}
func nox_xxx_wndDrawFnDefault_46B370(a1 int32, a2 *int32) int32 {
	var (
		v2     int32
		v3     int32
		result int32
		xLeft  int32 = 0
		yTop   int32 = 0
	)
	nox_client_wndGetPosition_46AA60((*nox_window)(unsafe.Pointer(uintptr(a1))), (*uint32)(unsafe.Pointer(&xLeft)), (*uint32)(unsafe.Pointer(&yTop)))
	if int32(int8(*(*uint8)(unsafe.Pointer(uintptr(a1 + 4))))) >= 0 {
		if uint32(*(*int32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(int32(0))*5))) != 0x80000000 {
			nox_client_drawSetColor_434460(*(*int32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(int32(0))*5)))
			nox_client_drawRectFilledOpaque_49CE30(xLeft, yTop, int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 8)))), int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 12)))))
		}
		result = 1
	} else {
		v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 100))))
		xLeft += int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 96))))
		v3 = v2 + yTop
		yTop = v3
		if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 36))))&2 != 0 {
			nox_client_drawImageAt_47D2C0(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(int32(0))*10)))), xLeft, v3)
		} else {
			nox_client_drawImageAt_47D2C0(unsafe.Pointer(uintptr(*(*int32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(int32(0))*6)))), xLeft, v3)
		}
		result = 1
	}
	return result
}
func nox_xxx_wnd_46AD60(a1 int32, a2 int32) int32 {
	var result int32
	if a1 == 0 {
		return -2
	}
	result = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 4))))
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 4))) = uint32(a2 | result)
	return result
}
func nox_xxx_wndClearFlag_46AD80(a1 int32, a2 int32) int32 {
	var result int32
	if a1 == 0 {
		return -2
	}
	result = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 4))))
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 4))) = uint32(result & ^a2)
	return result
}
func nox_xxx_wndGetFlags_46ADA0(a1 int32) int32 {
	var result int32
	if a1 != 0 {
		result = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 4))))
	} else {
		result = -2
	}
	return result
}
func nox_window_is_child(a1 *nox_window, a2 *nox_window) int32 {
	if a1 == nil {
		return 0
	}
	if a2 == nil {
		return 0
	}
	var cur *nox_window = a2
	for {
		cur = cur.parent
		if a1 == cur {
			break
		}
		if cur == nil {
			return 0
		}
	}
	return 1
}
func nox_xxx_wnd_46B280(a1 int32, a2 int32) int32 {
	if a1 == 0 {
		return -2
	}
	if a2 != 0 {
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 52))) = uint32(a2)
	} else {
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 52))) = uint32(a1)
	}
	return 0
}
func sub_46B2B0(a1 int32) int32 {
	var result int32
	result = a1
	if a1 != 0 {
		result = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 52))))
	}
	return result
}
func nox_client_inWindowByPos_46B5B0(root *nox_window, x int32, y int32) *nox_window {
	if root == nil {
		return nil
	}
	var cur *nox_window = root
LOOP:
	for win := (*nox_window)(cur.field_100); win != nil; win = win.prev {
		var (
			px int32 = win.off_x
			py int32 = win.off_y
		)
		for win2 := (*nox_window)(win.parent); win2 != nil; win2 = win2.parent {
			px += win2.off_x
			py += win2.off_y
		}
		if x >= px && x <= px+win.width && y >= py && y <= py+win.height {
			if win.flags&8 != 0 && (win.flags&16) == 0 {
				cur = win
				goto LOOP
			}
		}
	}
	return cur
}
func nox_xxx_wndLoadBorder_4AA1F0() *byte {
	var result *byte
	*memmap.PtrUint32(6112660, 1309676) = uint32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg("BorderCornerUL"))))
	*memmap.PtrUint32(6112660, 1309680) = uint32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg("BorderCornerUR"))))
	*memmap.PtrUint32(6112660, 1309684) = uint32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg("BorderCornerLL"))))
	*memmap.PtrUint32(6112660, 1309688) = uint32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg("BorderCornerLR"))))
	*memmap.PtrUint32(6112660, 1309692) = uint32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg("BorderHorizontal"))))
	dword_5d4594_1309696 = uint32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg("BorderHorizontalShort"))))
	*memmap.PtrUint32(6112660, 1309700) = uint32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg("BorderVertical"))))
	result = (*byte)(unsafe.Pointer(nox_xxx_gLoadImg("BorderVerticalShort")))
	dword_5d4594_1309704 = uint32(uintptr(unsafe.Pointer(result)))
	return result
}
