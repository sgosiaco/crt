package fonts

import (
	_ "embed"
)

//go:embed IosevkaTermNerdFontMono-Regular.ttf
var Regular []byte

//go:embed IosevkaTermNerdFontMono-Bold.ttf
var Bold []byte

//go:embed IosevkaTermNerdFontMono-Regular.ttf
var Italic []byte
