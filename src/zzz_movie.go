package opennox

import "unsafe"

func rgb888Toargb1555(red int32, green int32, blue int32) uint16 {
	var (
		b uint16 = uint16(int16((blue >> 3) & (31 << 0)))
		g uint16 = uint16(int16(((green >> 3) & 31) << 5))
		r uint16 = uint16(int16(((red >> 3) & 31) << 10))
		a uint16 = 1 << 15
	)
	return uint16(int16(int32(r) | int32(g) | int32(b) | int32(a)))
}

var movieSurface unsafe.Pointer = nil

func DrawMovieFrame(frame *uint8, cx uint32, cy uint32) {
	var (
		pixels unsafe.Pointer = nil
		pitch  int32          = 0
	)
	nox_video_getSurfaceData_48A720(movieSurface, &pitch, &pixels)
	for i := int32(0); uint32(i) < cy; i++ {
		var (
			surfaceRow *uint16 = (*uint16)(unsafe.Add(unsafe.Pointer((*uint8)(pixels)), i*pitch))
			frameRow   *uint16 = (*uint16)(unsafe.Add(unsafe.Pointer(frame), uint32(i)*cx*2))
		)
		for j := int32(0); uint32(j) < cx; j++ {
			*(*uint16)(unsafe.Add(unsafe.Pointer(surfaceRow), unsafe.Sizeof(uint16(0))*uintptr(j))) = *(*uint16)(unsafe.Add(unsafe.Pointer(frameRow), unsafe.Sizeof(uint16(0))*uintptr(j)))
		}
	}
	nox_video_showMovieFrame(movieSurface)
}
func PlayMovieCallback(frame *uint8, cx uint32, cy uint32) int32 {
	var result int32 = nox_input_pollEventsMovie()
	DrawMovieFrame(frame, cx, cy)
	return result
}
func sub_555430(a1 *unsafe.Pointer) {
	return
}
func sub_555500(a1 int8) int8 {
	return 0
}
