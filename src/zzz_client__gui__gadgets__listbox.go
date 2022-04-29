package opennox

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	"github.com/noxworld-dev/opennox/v1/common/alloc"
	"math"
	"unsafe"
)

func nox_gui_newScrollListBox_4A4310(a1p *nox_window, a2 int32, a3 int32, a4 int32, a5 int32, a6 int32, a7 int32, opts *nox_scrollListBox_data) *nox_window {
	var (
		a1  int32  = int32(uintptr(unsafe.Pointer(a1p)))
		a8  *int16 = (*int16)(unsafe.Pointer(opts))
		v8  *uint32
		v9  unsafe.Pointer
		v10 int32
		v11 int32
		v12 unsafe.Pointer
		v14 int32
		v15 int32
		v16 *libc.WChar
		v17 int32
		v18 *libc.WChar
		v19 int32
		v20 *int16
		v21 int32
		v22 int32
		v23 int32
		v24 [4]int32
		v25 [332]byte
		v26 uint32
	)
	v21 = 0
	if int32(opts.line_height) < nox_xxx_guiFontHeightMB_43F320(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a7 + 200)))))) {
		opts.line_height = uint16(int16(nox_xxx_guiFontHeightMB_43F320(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a7 + 200))))))))
	}
	if int32(*(*uint16)(unsafe.Pointer(uintptr(a7 + 72)))) != 0 {
		v21 = 1
	}
	if (int32(*(*uint8)(unsafe.Pointer(uintptr(a7 + 8)))) & 32) == 0 {
		return nil
	}
	v8 = (*uint32)(unsafe.Pointer(nox_window_new((*nox_window)(unsafe.Pointer(uintptr(a1))), a2, a3, a4, a5, a6, func(arg1 int32, arg2 int32, arg3 int32, arg4 int32) int32 {
		return nox_xxx_wndListboxProcPre_4A30D0((*nox_window)(unsafe.Pointer(uintptr(arg1))), uint32(arg2), uint32(arg3), arg4)
	})))
	nox_xxx_wndListboxInit_4A3C00(int32(uintptr(unsafe.Pointer(v8))), int32(uintptr(unsafe.Pointer(opts))))
	if v8 == nil {
		return nil
	}
	if *(*uint32)(unsafe.Pointer(uintptr(a7 + 16))) == 0 {
		*(*uint32)(unsafe.Pointer(uintptr(a7 + 16))) = uint32(uintptr(unsafe.Pointer(v8)))
	}
	nox_gui_windowCopyDrawData_46AF80((*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v8)))))), unsafe.Pointer(uintptr(a7)))
	v9 = unsafe.Pointer(&make([]nox_scrollListBox_item, int(opts.count))[0])
	*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a8))), unsafe.Sizeof(uint32(0))*6))) = uint32(uintptr(v9))
	if v9 == nil {
		return nil
	}
	libc.MemSet(v9, 0, int(int32(opts.count)*524))
	v10 = a6
	*(*int16)(unsafe.Add(unsafe.Pointer(a8), unsafe.Sizeof(int16(0))*26)) = int16(a6)
	if v21 != 0 {
		*(*int16)(unsafe.Add(unsafe.Pointer(a8), unsafe.Sizeof(int16(0))*26)) -= int16(nox_xxx_guiFontHeightMB_43F320(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a7 + 200)))))))
	}
	v11 = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a8))), unsafe.Sizeof(uint32(0))*4))))
	*(*int16)(unsafe.Add(unsafe.Pointer(a8), unsafe.Sizeof(int16(0))*27)) = 0
	*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a8))), unsafe.Sizeof(uint32(0))*12))) = math.MaxUint32
	*(*int16)(unsafe.Add(unsafe.Pointer(a8), unsafe.Sizeof(int16(0))*23)) = 0
	*(*int16)(unsafe.Add(unsafe.Pointer(a8), unsafe.Sizeof(int16(0))*22)) = 0
	*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a8))), unsafe.Sizeof(uint32(0))*10))) = 0
	if v11 != 0 {
		v12 = alloc.Calloc(int(opts.count), 4)
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a8))), unsafe.Sizeof(uint32(0))*12))) = uint32(uintptr(v12))
		if v12 == nil {
			*((*unsafe.Pointer)(unsafe.Add(unsafe.Pointer((*unsafe.Pointer)(unsafe.Pointer(a8))), unsafe.Sizeof(unsafe.Pointer(nil))*6))) = nil
			return nil
		}
		libc.MemSet(v12, math.MaxUint8, int(int32(opts.count)*4))
		v10 = a6
	}
	if *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a8))), unsafe.Sizeof(uint32(0))*3))) != 0 {
		v24[0] = 0
		v14 = int32(*(*uint32)(unsafe.Pointer(uintptr(a7 + 200))))
		v24[1] = 0
		v24[2] = 0
		v24[3] = 0
		v26 = uint32(a2) & 0xFFFFEFEF
		v15 = nox_xxx_guiFontHeightMB_43F320(unsafe.Pointer(uintptr(v14)))
		if v21 != 0 {
			v22 = v15 + 1
			v10 = v10 - v15 - 1
		} else {
			v22 = 0
		}
		*(*[332]byte)(unsafe.Pointer(&v25[0])) = [332]byte{}
		if int32(int8(*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(v8))), 4))))) >= 0 {
			*(*uint32)(unsafe.Pointer(&v25[20])) = nox_color_black_2650656
			*(*uint32)(unsafe.Pointer(&v25[44])) = nox_color_black_2650656
			*(*uint32)(unsafe.Pointer(&v25[28])) = nox_color_orange_2614256
			*(*uint32)(unsafe.Pointer(&v25[36])) = nox_color_white_2523948
			*(*uint32)(unsafe.Pointer(&v25[52])) = nox_color_yellow_2589772
			*(*uint32)(unsafe.Pointer(&v25[68])) = nox_color_orange_2614256
			v16 = strMan.GetStringInFile("WindowDir:Up", "C:\\NoxPost\\src\\Client\\Gui\\Gadgets\\listbox.c")
			nox_wcscpy((*libc.WChar)(unsafe.Pointer(&v25[72])), v16)
			v23 = 10
		} else {
			*(*uint32)(unsafe.Pointer(&v25[24])) = uint32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg("DefaultLBUpButton"))))
			*(*uint32)(unsafe.Pointer(&v25[40])) = uint32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg("DefaultLBUpButtonLit"))))
			*(*uint32)(unsafe.Pointer(&v25[48])) = uint32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg("DefaultLBUpButtonDis"))))
			*(*uint32)(unsafe.Pointer(&v25[56])) = uint32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg("DefaultLBUpButtonLit"))))
			*(*uint32)(unsafe.Pointer(&v25[32])) = 0
			v23 = 13
		}
		v17 = int32(v26 | 9)
		*(*uint32)(unsafe.Pointer(&v25[16])) = uint32(uintptr(unsafe.Pointer(v8)))
		*(*uint32)(unsafe.Pointer(&v25[8])) = 1
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a8))), unsafe.Sizeof(uint32(0))*7))) = uint32(uintptr(unsafe.Pointer(nox_gui_newButtonOrCheckbox_4A91A0((*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v8)))))), int32(v26|9), a5-10, v22, 10, v23, (*nox_window_data)(unsafe.Pointer(&v25[0]))))))
		*(*[332]byte)(unsafe.Pointer(&v25[0])) = [332]byte{}
		if int32(int8(*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(v8))), 4))))) >= 0 {
			*(*uint32)(unsafe.Pointer(&v25[20])) = nox_color_black_2650656
			*(*uint32)(unsafe.Pointer(&v25[44])) = nox_color_black_2650656
			*(*uint32)(unsafe.Pointer(&v25[28])) = nox_color_orange_2614256
			*(*uint32)(unsafe.Pointer(&v25[36])) = nox_color_white_2523948
			*(*uint32)(unsafe.Pointer(&v25[52])) = nox_color_yellow_2589772
			*(*uint32)(unsafe.Pointer(&v25[68])) = nox_color_orange_2614256
			v18 = strMan.GetStringInFile("WindowDir:Down", "C:\\NoxPost\\src\\Client\\Gui\\Gadgets\\listbox.c")
			nox_wcscpy((*libc.WChar)(unsafe.Pointer(&v25[72])), v18)
		} else {
			*(*uint32)(unsafe.Pointer(&v25[24])) = uint32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg("DefaultLBDownButton"))))
			*(*uint32)(unsafe.Pointer(&v25[40])) = uint32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg("DefaultLBDownButtonLit"))))
			*(*uint32)(unsafe.Pointer(&v25[48])) = uint32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg("DefaultLBDownButtonDis"))))
			*(*uint32)(unsafe.Pointer(&v25[56])) = uint32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg("DefaultLBDownButtonLit"))))
			*(*uint32)(unsafe.Pointer(&v25[32])) = 0
		}
		*(*uint32)(unsafe.Pointer(&v25[8])) = 1
		*(*uint32)(unsafe.Pointer(&v25[16])) = uint32(uintptr(unsafe.Pointer(v8)))
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a8))), unsafe.Sizeof(uint32(0))*8))) = uint32(uintptr(unsafe.Pointer(nox_gui_newButtonOrCheckbox_4A91A0((*nox_window)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v8)))))), v17, a5-10, v22+v10-v23, 10, v23, (*nox_window_data)(unsafe.Pointer(&v25[0]))))))
		*(*[332]byte)(unsafe.Pointer(&v25[0])) = [332]byte{}
		if int32(int8(*((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Pointer(v8))), 4))))) >= 0 {
			*(*uint32)(unsafe.Pointer(&v25[20])) = nox_color_black_2650656
			*(*uint32)(unsafe.Pointer(&v25[44])) = nox_color_black_2650656
			*(*uint32)(unsafe.Pointer(&v25[36])) = nox_color_black_2650656
			*(*uint32)(unsafe.Pointer(&v25[28])) = nox_color_orange_2614256
			*(*uint32)(unsafe.Pointer(&v25[52])) = nox_color_orange_2614256
			v19 = 10
		} else {
			*(*uint32)(unsafe.Pointer(&v25[24])) = uint32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg("DefaultSliderThumb"))))
			*(*uint32)(unsafe.Pointer(&v25[40])) = uint32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg("DefaultSliderThumbLit"))))
			*(*uint32)(unsafe.Pointer(&v25[48])) = uint32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg("DefaultSliderThumbDis"))))
			*(*uint32)(unsafe.Pointer(&v25[56])) = uint32(uintptr(unsafe.Pointer(nox_xxx_gLoadImg("DefaultSliderThumbLit"))))
			*(*uint32)(unsafe.Pointer(&v25[32])) = 0
			v19 = 9
		}
		v24[0] = 0
		v24[1] = 0
		v24[2] = 0
		*(*uint32)(unsafe.Pointer(&v25[8])) = 8
		v24[3] = 0
		*(*uint32)(unsafe.Pointer(&v25[16])) = uint32(uintptr(unsafe.Pointer(v8)))
		*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(a8))), unsafe.Sizeof(uint32(0))*9))) = uint32(uintptr(unsafe.Pointer(nox_gui_newSlider_4B4EE0(int32(uintptr(unsafe.Pointer(v8))), v17, a5-v19, v22+v23, v19, v10-v23*2, (*uint32)(unsafe.Pointer(&v25[0])), (*float32)(unsafe.Pointer(&v24[0]))))))
	}
	v20 = (*int16)(unsafe.Pointer(new(nox_scrollListBox_data)))
	libc.MemCpy(unsafe.Pointer(v20), unsafe.Pointer(opts), int(unsafe.Sizeof(nox_scrollListBox_data{})))
	*(*uint32)(unsafe.Add(unsafe.Pointer(v8), unsafe.Sizeof(uint32(0))*8)) = uint32(uintptr(unsafe.Pointer(v20)))
	return (*nox_window)(unsafe.Pointer(v8))
}
