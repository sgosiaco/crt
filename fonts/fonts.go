package fonts

import (
	_ "embed"
)

var (
	//go:embed IosevkaTermNerdFontMono/IosevkaTermNerdFontMono-Regular.ttf
	IosevkaTermNerdFontMonoRegular []byte
	//go:embed IosevkaTermNerdFontMono/IosevkaTermNerdFontMono-Bold.ttf
	IosevkaTermNerdFontMonoBold []byte
	//go:embed IosevkaTermNerdFontMono/IosevkaTermNerdFontMono-Regular.ttf
	IosevkaTermNerdFontMonoItalic []byte

	//go:embed UbuntuMono/UbuntuMonoNerdFontMono-Regular.ttf
	UbuntuMonoRegular []byte
	//go:embed UbuntuMono/UbuntuMonoNerdFontMono-Bold.ttf
	UbuntuMonoBold []byte
	//go:embed UbuntuMono/UbuntuMonoNerdFontMono-Italic.ttf
	UbuntuMonoItalic []byte
)
