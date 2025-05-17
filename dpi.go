package crt

import (
	"os"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
)

// DeviceScale returns the current device scale factor.
//
// If the environment variable CRT_DEVICE_SCALE is set, it will be used instead.
func DeviceScale() float64 {
	if os.Getenv("CRT_DEVICE_SCALE") != "" {
		s, err := strconv.ParseFloat(os.Getenv("CRT_DEVICE_SCALE"), 64)
		if err == nil {
			return s
		}
	}

	return ebiten.Monitor().DeviceScaleFactor() // Deprecated: ebiten.DeviceScaleFactor()
}

// GetFontDPI returns the recommended font DPI for the current device.
func GetFontDPI() float64 {
	return 72.0 * DeviceScale()
}
