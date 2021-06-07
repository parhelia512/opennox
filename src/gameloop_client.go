//+build !server

package main

/*
#include <fenv.h>

#include "proto.h"
#include "common__system__team.h"
#include "common__system__gamedisk.h"
#include "client__system__gameloop.h"

extern unsigned int nox_client_gui_flag_1556112;
extern unsigned int nox_client_gui_flag_815132;
extern int nox_win_width;
extern int nox_win_height;
extern BYTE** nox_pixbuffer_rows_3798784;
*/
import "C"
import (
	"image"
	"image/color"
	"math"
	"unsafe"

	"nox/v1/common/alloc"
	"nox/v1/common/memmap"
)

func mainloopDrawAndPresent() {
	C.sub_437180()
	if C.nox_client_gui_flag_1556112 == 0 {
		DrawGUI() // Draw game windows
	}
	DrawSparks()
	if !getEngineFlag(NOX_ENGINE_FLAG_DISABLE_GRAPHICS_RENDERING) || getEngineFlag(NOX_ENGINE_FLAG_9) || C.nox_client_gui_flag_815132 != 0 {
		C.nox_client_drawCursorAndTooltips_477830() // Draw cursor
	}
	C.sub_44D9F0(1)
	maybeScreenshot()
	if C.sub_409F40(4096) == 0 { // CheckRuleFlags and smth
		C.nox_xxx_screenshot_46D830()
	}
	if !getEngineFlag(NOX_ENGINE_FLAG_DISABLE_GRAPHICS_RENDERING) || getEngineFlag(NOX_ENGINE_FLAG_9) || C.nox_client_gui_flag_815132 != 0 {
		// C.nox_xxx_directDrawBlitMB_48A220() // does nothing
		nox_video_callCopyBackBuffer_4AD170()
		callPresent()
	}
}

func getGamePixBufferC() []unsafe.Pointer {
	return asPtrSlice(unsafe.Pointer(C.nox_pixbuffer_rows_3798784), int(C.nox_win_height))
}

func colorRGB15(cl uint16) color.RGBA {
	r := byte((cl & 0xfc00) >> 10)
	g := byte((cl & 0x03e0) >> 5)
	b := byte((cl & 0x001f) >> 0)
	r = byte((float64(r) / 31) * 0xff)
	g = byte((float64(g) / 31) * 0xff)
	b = byte((float64(b) / 31) * 0xff)
	return color.RGBA{
		R: r, G: g, B: b, A: 0xff,
	}
}

func copyGamePixBuffer() image.Image {
	sz := videoGetWindowSize()
	img := image.NewRGBA(image.Rect(0, 0, sz.W, sz.H))

	pixbuf := getGamePixBufferC()
	for y := 0; y < sz.H; y++ {
		row := asU16Slice(pixbuf[y], sz.W)
		for x := 0; x < sz.W; x++ {
			img.SetRGBA(x, y, colorRGB15(row[x]))
		}
	}
	return img
}

func DrawSparks() {
	if C.nox_client_gui_flag_815132 != 0 {
		sz := videoGetWindowSize()
		rdr := (*C.nox_draw_viewport_t)(alloc.Malloc(unsafe.Sizeof(C.nox_draw_viewport_t{})))
		rdr.x1 = 0
		rdr.y1 = 0
		rdr.x2 = C.int(sz.W)
		rdr.y2 = C.int(sz.H)
		rdr.width = C.int(sz.W)
		rdr.height = C.int(sz.H)
		C.nox_client_screenParticlesDraw_431720(rdr)
		alloc.Free(unsafe.Pointer(rdr))
	} else {
		rdr := C.nox_draw_getViewport_437250()
		C.nox_client_screenParticlesDraw_431720(rdr)
	}
}

func generateMouseSparks() {
	if memmap.Uint32(0x5D4594, 816408) != 0 {
		return
	}

	mouse := C.nox_client_getMouseState_4309F0()
	// emit sparks when passing a certain distance
	const distanceSparks = 0.25
	dx := int(mouse.pos.x) - int(memmap.Uint32(0x5D4594, 816420))
	dy := int(mouse.pos.y) - int(memmap.Uint32(0x5D4594, 816424))
	r2 := dx*dx + dy*dy
	if memmap.Uint32(0x5D4594, 816428) != 0 {
		cnt := (int)(math.Sqrt(float64(r2)) * distanceSparks)
		for i := cnt; i > 0; i-- {
			v6 := randomIntMinMax(0, 100)
			v7 := int(memmap.Uint32(0x5D4594, 816420)) + dx*v6/100
			v9 := int(memmap.Uint32(0x5D4594, 816424)) + dy*v6/100
			v23 := randomIntMinMax(2, 5)
			v22 := randomIntMinMax(2, 5)
			v21 := randomIntMinMax(-7, 2)
			v10 := randomIntMinMax(-5, 5)
			C.nox_client_newScreenParticle_431540(4, C.int(v7), C.int(v9), C.int(v10), C.int(v21), 1, C.char(v22), C.char(v23), 2, 1)
		}
		if r2 < 10 {
			*memmap.PtrUint32(0x5D4594, 816428) = 0
		}
		*memmap.PtrUint32(0x5D4594, 816420) = uint32(mouse.pos.x)
		*memmap.PtrUint32(0x5D4594, 816424) = uint32(mouse.pos.y)
	} else if r2 > 64 {
		*memmap.PtrUint32(0x5D4594, 816428) = 1
	}
	// explode with sparks when clicking
	const explosionSparks = 75
	if mouse.btn[C.NOX_MOUSE_LEFT].pressed == 1 {
		randomIntMinMax(0, 2)
		if memmap.Uint32(0x5D4594, 816416) == 0 {
			*memmap.PtrUint32(0x5D4594, 816416) = 1
			clientPlaySoundSpecial(924, 100)
			for i := explosionSparks; i > 0; i-- {
				v12 := randomIntMinMax(0, 255)
				v13 := randomIntMinMax(6, 12)
				v14 := v13 * int(memmap.Int32(0x587000, 192088+8*uintptr(v12)))
				v15 := v13*int(memmap.Int32(0x587000, 192092+8*uintptr(v12)))/16 - 6
				v24 := randomIntMinMax(2, 5)
				v16 := randomIntMinMax(2, 5)
				C.nox_client_newScreenParticle_431540(4, C.int(v14/16+int(mouse.pos.x)), C.int(int(mouse.pos.y)+v15), C.int(v14/16), C.int(v15), 1, C.char(v16), C.char(v24), 2, 1)
			}
		}
	} else {
		*memmap.PtrUint32(0x5D4594, 816416) = 0
	}
}
