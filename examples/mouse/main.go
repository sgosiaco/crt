package main

// A simple program that opens the alternate screen buffer and displays mouse
// coordinates and events.

import (
	"image/color"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/sgosiaco/crt"
	bubbleadapter "github.com/sgosiaco/crt/bubbletea"
)

const (
	Width  = 1000
	Height = 600
)

func main() {
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		fmt.Println(err)
	// 		debug.PrintStack()
	// 	}
	// }()

	fonts, err := crt.LoadFaces("./fonts/IosevkaTermNerdFontMono-Regular.ttf", "./fonts/IosevkaTermNerdFontMono-Bold.ttf", "./fonts/IosevkaTermNerdFontMono-Italic.ttf", crt.GetFontDPI(), 16.0)
	if err != nil {
		panic(err)
	}

	win, _, err := bubbleadapter.Window(Width, Height, fonts, model{}, color.Black)
	if err != nil {
		panic(err)
	}

	// prog.Send(tea.ShowCursor())
	// win.SetCursorChar("$")

	if err := win.Run("Mouse"); err != nil {
		panic(err)
	}

	// p := tea.NewProgram(model{}, tea.WithMouseAllMotion())
	// if _, err := p.Run(); err != nil {
	// 	log.Fatal(err)
	// }
}

type model struct {
	mouseEvent tea.MouseEvent
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if s := msg.String(); s == "ctrl+c" || s == "q" || s == "esc" {
			return m, tea.Quit
		}

	case tea.MouseMsg:
		return m, tea.Printf("(X: %d, Y: %d) %s", msg.X, msg.Y, tea.MouseEvent(msg))
	}

	return m, nil
}

func (m model) View() string {
	s := "Do mouse stuff. When you're done press q to quit.\n"
	//s := "Do mouse stuff"
	return s
	//var sb strings.Builder
	//for i := range 20 {
	//	sb.WriteString(strings.Repeat(string(rune(65+i)), i+1))
	//}

	// return sb.String()
}
