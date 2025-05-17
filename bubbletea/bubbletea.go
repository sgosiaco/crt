package bubbletea

import (
	"fmt"
	"image/color"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
	"github.com/sgosiaco/crt"
)

func init() {
	lipgloss.SetColorProfile(termenv.TrueColor)
}

// Window creates a new crt based bubbletea window with the given width, height, fonts, model and default background color.
// Additional options can be passed to the bubbletea program.
func Window(width int, height int, fonts crt.Fonts, model tea.Model, defaultBg color.Color, options ...tea.ProgramOption) (*crt.Window, *tea.Program, error) {
	// TODO: Maybe try feeding ebiten keys directly to gameInput instead of having to manually map everything
	//gameInput := crt.NewConcurrentRW()
	gameOutput := crt.NewConcurrentRW()

	//go gameInput.Run()
	go gameOutput.Run()

	prog := tea.NewProgram(
		model,
		append([]tea.ProgramOption{
			tea.WithMouseAllMotion(),
			tea.WithInput(nil), // (gameInput),
			tea.WithOutput(gameOutput),
			tea.WithANSICompressor(),
		}, options...)...,
	)

	go func() {
		if _, err := prog.Run(); err != nil {
			fmt.Printf("Alas, there's been an error: %v", err)
		}

		crt.SysKill()
	}()

	win, err := crt.NewGame(width, height, fonts, gameOutput, NewAdapter(prog), defaultBg)
	return win, prog, err
}
