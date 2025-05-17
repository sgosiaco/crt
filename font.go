package crt

import (
	"bytes"
	"os"

	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/sgosiaco/crt/fonts"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/text/language"
)

type Fonts struct {
	NFace  font.Face
	Normal *text.GoTextFace
	Bold   *text.GoTextFace
	Italic *text.GoTextFace
}

// LoadFaceBytesGo loads a font face from bytes. The dpi and size are used to generate the font face.
// The normal, bold, and italic files must be provided. Supports ttf and otf.
func LoadFaceBytesGo(file []byte, dpi float64, size float64) (font.Face, error) {
	tt, err := opentype.Parse(file)
	if err != nil {
		panic(err)
	}

	face, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    size,
		DPI:     dpi,
		Hinting: font.HintingNone,
	})
	if err != nil {
		return nil, err
	}

	return face, nil
}

// LoadFaceBytes loads a font face from bytes. The dpi and size are used to generate the font face.
// The normal, bold, and italic files must be provided. Supports ttf and otf.
func LoadFaceBytes(file []byte, dpi float64, size float64) (*text.GoTextFace, error) {
	source, err := text.NewGoTextFaceSource(bytes.NewReader(file))
	if err != nil {
		return nil, err
	}

	return &text.GoTextFace{
		Source:    source,
		Direction: text.DirectionLeftToRight,
		Size:      size,
		Language:  language.English,
	}, nil
}

// LoadFacesBytes loads a set of fonts from bytes. The normal, bold, and italic files
// must be provided. The dpi and size are used to generate the font faces. Supports ttf and otf.
func LoadFacesBytes(normal []byte, bold []byte, italic []byte, dpi float64, size float64) (Fonts, error) {
	nFace, err := LoadFaceBytesGo(normal, dpi, size)
	if err != nil {
		return Fonts{}, err
	}

	normalFace, err := LoadFaceBytes(normal, dpi, size)
	if err != nil {
		return Fonts{}, err
	}

	boldFace, err := LoadFaceBytes(bold, dpi, size)
	if err != nil {
		return Fonts{}, err
	}

	italicFace, err := LoadFaceBytes(italic, dpi, size)
	if err != nil {
		return Fonts{}, err
	}

	return Fonts{
		NFace:  nFace,
		Normal: normalFace,
		Bold:   boldFace,
		Italic: italicFace,
	}, nil
}

// LoadFace loads a font face from a file. The dpi and size are used to generate the font face. Supports ttf and otf.
//
// Example: LoadFace("./fonts/Mono-Regular.ttf", 72.0, 16.0)
func LoadFace(file string, dpi float64, size float64) (*text.GoTextFace, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	return LoadFaceBytes(data, dpi, size)
}

// LoadFace loads a font face from a file. The dpi and size are used to generate the font face. Supports ttf and otf.
//
// Example: LoadFace("./fonts/Mono-Regular.ttf", 72.0, 16.0)
func LoadFaceGo(file string, dpi float64, size float64) (font.Face, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	return LoadFaceBytesGo(data, dpi, size)
}

// LoadFaces loads a set of fonts from files. The normal, bold, and italic files
// must be provided. The dpi and size are used to generate the font faces. Supports ttf and otf.
func LoadFaces(normal string, bold string, italic string, dpi float64, size float64) (Fonts, error) {
	nFace, err := LoadFaceGo(normal, dpi, size)
	if err != nil {
		return Fonts{}, err
	}

	normalFace, err := LoadFace(normal, dpi, size)
	if err != nil {
		return Fonts{}, err
	}

	boldFace, err := LoadFace(bold, dpi, size)
	if err != nil {
		return Fonts{}, err
	}

	italicFace, err := LoadFace(italic, dpi, size)
	if err != nil {
		return Fonts{}, err
	}

	return Fonts{
		NFace:  nFace,
		Normal: normalFace,
		Bold:   boldFace,
		Italic: italicFace,
	}, nil
}

// LoadDefaultFaces loads the default fonts embebbed in the library
func LoadDefaultFaces(dpi float64, size float64) (Fonts, error) {
	return LoadFacesBytes(fonts.Regular, fonts.Bold, fonts.Italic, dpi, size)
}
