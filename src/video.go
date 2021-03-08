package main

import (
	"image"

	"nox/common/types"
)

var noxPresentTicks uint

func callPresent() {
	presentFrame()
	noxPresentTicks++
}

func changeWindowedOrFullscreen() {
	winSize := noxGetWinSize1()
	dispSize := getDisplayDim()
	centeredPos := image.Point{
		X: dispSize[2] + (dispSize[0]-winSize.W)/2,
		Y: dispSize[3] + (dispSize[1]-winSize.H)/2,
	}

	// Init all sizes

	// Windowed
	windowedSize := winSize
	windowedPos := centeredPos

	// Fullscreen
	fullscreenSize := types.Size{W: dispSize[0], H: dispSize[1]}
	fullscreenPos := image.Point{X: dispSize[2], Y: dispSize[3]}

	switch noxFullScreen {
	case -1, 1:
		// Normal fullscreen
		setFullScreenMode(fullscreenSize, fullscreenPos)
	case -2, 2:
		// Borderless fullscreen
		setFullScreenBorderlessMode(fullscreenSize, fullscreenPos)
	default:
		// Windowed
		setWindowedMode(windowedSize, windowedPos)
	}
}
