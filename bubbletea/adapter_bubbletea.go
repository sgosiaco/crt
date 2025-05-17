package bubbletea

import (
	"unicode"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/sgosiaco/crt"
)

type teaKey struct {
	key  tea.KeyType
	rune []rune
}

func repeatingKeyPressed(key ebiten.Key) bool {
	const (
		delay    = 30
		interval = 3
	)
	d := inpututil.KeyPressDuration(key)
	if d == 1 {
		return true
	}
	if d >= delay && (d-delay)%interval == 0 {
		return true
	}
	return false
}

var ebitenToTeaKeys = map[ebiten.Key]teaKey{
	ebiten.KeyEnter:      {tea.KeyEnter, []rune{'\n'}},
	ebiten.KeyTab:        {tea.KeyTab, []rune{}},
	ebiten.KeyBackspace:  {tea.KeyBackspace, []rune{}},
	ebiten.KeyInsert:     {tea.KeyInsert, []rune{}},
	ebiten.KeyDelete:     {tea.KeyDelete, []rune{}},
	ebiten.KeyHome:       {tea.KeyHome, []rune{}},
	ebiten.KeyEnd:        {tea.KeyEnd, []rune{}},
	ebiten.KeyPageUp:     {tea.KeyPgUp, []rune{}},
	ebiten.KeyPageDown:   {tea.KeyPgDown, []rune{}},
	ebiten.KeyArrowUp:    {tea.KeyUp, []rune{}},
	ebiten.KeyArrowDown:  {tea.KeyDown, []rune{}},
	ebiten.KeyArrowLeft:  {tea.KeyLeft, []rune{}},
	ebiten.KeyArrowRight: {tea.KeyRight, []rune{}},
	ebiten.KeyEscape:     {tea.KeyEscape, []rune{}},
	ebiten.KeyF1:         {tea.KeyF1, []rune{}},
	ebiten.KeyF2:         {tea.KeyF2, []rune{}},
	ebiten.KeyF3:         {tea.KeyF3, []rune{}},
	ebiten.KeyF4:         {tea.KeyF4, []rune{}},
	ebiten.KeyF5:         {tea.KeyF5, []rune{}},
	ebiten.KeyF6:         {tea.KeyF6, []rune{}},
	ebiten.KeyF7:         {tea.KeyF7, []rune{}},
	ebiten.KeyF8:         {tea.KeyF8, []rune{}},
	ebiten.KeyF9:         {tea.KeyF9, []rune{}},
	ebiten.KeyF10:        {tea.KeyF10, []rune{}},
	ebiten.KeyF11:        {tea.KeyF11, []rune{}},
	ebiten.KeyF12:        {tea.KeyF12, []rune{}},
	ebiten.KeyShift:      {tea.KeyShiftLeft, []rune{}},
	ebiten.KeyShiftLeft:  {tea.KeyShiftLeft, []rune{}},
	ebiten.KeyShiftRight: {tea.KeyShiftRight, []rune{}},
}

var ebitenToCtrlKeys = map[ebiten.Key]tea.KeyType{
	ebiten.KeyA:            tea.KeyCtrlA,
	ebiten.KeyB:            tea.KeyCtrlB,
	ebiten.KeyC:            tea.KeyCtrlC,
	ebiten.KeyD:            tea.KeyCtrlD,
	ebiten.KeyE:            tea.KeyCtrlE,
	ebiten.KeyF:            tea.KeyCtrlF,
	ebiten.KeyG:            tea.KeyCtrlG,
	ebiten.KeyH:            tea.KeyCtrlH,
	ebiten.KeyI:            tea.KeyCtrlI,
	ebiten.KeyJ:            tea.KeyCtrlJ,
	ebiten.KeyK:            tea.KeyCtrlK,
	ebiten.KeyL:            tea.KeyCtrlL,
	ebiten.KeyM:            tea.KeyCtrlM,
	ebiten.KeyN:            tea.KeyCtrlN,
	ebiten.KeyO:            tea.KeyCtrlO,
	ebiten.KeyP:            tea.KeyCtrlP,
	ebiten.KeyQ:            tea.KeyCtrlQ,
	ebiten.KeyR:            tea.KeyCtrlR,
	ebiten.KeyS:            tea.KeyCtrlS,
	ebiten.KeyT:            tea.KeyCtrlT,
	ebiten.KeyU:            tea.KeyCtrlU,
	ebiten.KeyV:            tea.KeyCtrlV,
	ebiten.KeyW:            tea.KeyCtrlW,
	ebiten.KeyX:            tea.KeyCtrlX,
	ebiten.KeyY:            tea.KeyCtrlY,
	ebiten.KeyZ:            tea.KeyCtrlZ,
	ebiten.KeyLeftBracket:  tea.KeyCtrlOpenBracket,
	ebiten.KeyBackslash:    tea.KeyCtrlBackslash,
	ebiten.KeyRightBracket: tea.KeyCtrlCloseBracket,
	ebiten.KeyApostrophe:   tea.KeyCtrlCaret,
}

var ebitenToShiftKeys = map[ebiten.Key]tea.KeyType{
	ebiten.KeyTab:   tea.KeyShiftTab,
	ebiten.KeyUp:    tea.KeyShiftUp,
	ebiten.KeyDown:  tea.KeyShiftDown,
	ebiten.KeyRight: tea.KeyShiftRight,
	ebiten.KeyLeft:  tea.KeyShiftLeft,
	ebiten.KeyHome:  tea.KeyShiftHome,
	ebiten.KeyEnd:   tea.KeyShiftEnd,
}

var ebitenToTeaMouseNew = map[ebiten.MouseButton]tea.MouseButton{
	ebiten.MouseButtonLeft:   tea.MouseButtonLeft,
	ebiten.MouseButtonMiddle: tea.MouseButtonMiddle,
	ebiten.MouseButtonRight:  tea.MouseButtonRight,

	// TODO: is this right?
	ebiten.MouseButton3: tea.MouseButtonBackward,
	ebiten.MouseButton4: tea.MouseButtonForward,
}

// Options are used to configure the adapter.
type Options func(*Adapter)

// WithFilterMousePressed filters the MousePressed event and only emits MouseReleased events.
func WithFilterMousePressed(filter bool) Options {
	return func(b *Adapter) {
		b.filterMousePressed = filter
	}
}

// Adapter represents a bubbletea adapter for the crt package.
type Adapter struct {
	prog               *tea.Program
	runeBuffer         []rune
	keyBuffer          []ebiten.Key
	filterMousePressed bool
}

// NewAdapter creates a new bubbletea adapter.
func NewAdapter(prog *tea.Program, options ...Options) *Adapter {
	b := &Adapter{
		prog:               prog,
		runeBuffer:         make([]rune, 100),       // TODO: Determine best sizes, but for now 100 should be ok
		keyBuffer:          make([]ebiten.Key, 100), // TODO: Determine best sizes, but for now 100 should be ok
		filterMousePressed: true,
	}

	for i := range options {
		options[i](b)
	}

	return b
}

func (b *Adapter) HandleMouseMotion(motion crt.MouseMotion) {
	b.prog.Send(tea.MouseMsg{
		X:      motion.X,
		Y:      motion.Y,
		Shift:  motion.Shift,
		Alt:    motion.Alt,
		Ctrl:   motion.Ctrl,
		Action: tea.MouseActionMotion,
	})
}

func (b *Adapter) HandleMouseButton(button crt.MouseButton) {
	// Filter this event or two events will be sent for one click in the current bubbletea version.
	// if b.filterMousePressed && button.JustPressed {
	// 	return
	// }

	msg := tea.MouseMsg{
		X:      button.X,
		Y:      button.Y,
		Shift:  button.Shift,
		Alt:    button.Alt,
		Ctrl:   button.Ctrl,
		Button: ebitenToTeaMouseNew[button.Button],
	}

	if button.JustReleased {
		msg.Action = tea.MouseActionRelease
	} else if button.JustPressed {
		msg.Action = tea.MouseActionPress
	}

	b.prog.Send(msg)
}

func (b *Adapter) HandleMouseWheel(wheel crt.MouseWheel) {
	direction := tea.MouseButtonNone
	if wheel.DY > 0 {
		direction = tea.MouseButtonWheelUp
	} else if wheel.DY < 0 {
		direction = tea.MouseButtonWheelUp
	}

	if direction == tea.MouseButtonNone {
		return
	}

	b.prog.Send(tea.MouseMsg{
		X:      wheel.X,
		Y:      wheel.Y,
		Shift:  wheel.Shift,
		Alt:    wheel.Alt,
		Ctrl:   wheel.Ctrl,
		Button: direction,
	})
}

func (b *Adapter) HandleKeyPress() {
	// "reset" buffer before using
	b.runeBuffer = ebiten.AppendInputChars(b.runeBuffer[:0])
	for _, v := range b.runeBuffer {
		switch v {
		case ' ':
			b.prog.Send(tea.KeyMsg{
				Type:  tea.KeySpace,
				Runes: []rune{v},
				Alt:   ebiten.IsKeyPressed(ebiten.KeyAlt),
			})
		default:
			b.prog.Send(tea.KeyMsg{
				Type:  tea.KeyRunes,
				Runes: []rune{v},
				Alt:   ebiten.IsKeyPressed(ebiten.KeyAlt),
			})
		}
	}

	// "reset" buffer before using
	b.keyBuffer = inpututil.AppendJustPressedKeys(b.keyBuffer[:0])
	repeatedBackspace := repeatingKeyPressed(ebiten.KeyBackspace)

	if repeatedBackspace {
		b.prog.Send(tea.KeyMsg{
			Type:  tea.KeyBackspace,
			Runes: []rune{},
			Alt:   ebiten.IsKeyPressed(ebiten.KeyAlt),
		})
	}

	for _, k := range b.keyBuffer {
		if ebiten.IsKeyPressed(ebiten.KeyControl) {
			if tk, ok := ebitenToCtrlKeys[k]; ok {
				b.prog.Send(tea.KeyMsg{
					Type:  tk,
					Runes: []rune{},
					Alt:   ebiten.IsKeyPressed(ebiten.KeyAlt),
				})
				continue
			}
		}

		if ebiten.IsKeyPressed(ebiten.KeyShift) {
			if tk, ok := ebitenToShiftKeys[k]; ok {
				b.prog.Send(tea.KeyMsg{
					Type:  tk,
					Runes: []rune{},
					Alt:   ebiten.IsKeyPressed(ebiten.KeyAlt),
				})
				continue
			}
		}

		if repeatedBackspace && k == ebiten.KeyBackspace {
			continue
		}

		if val, ok := ebitenToTeaKeys[k]; ok {
			runes := make([]rune, len(val.rune))
			copy(runes, val.rune)

			if ebiten.IsKeyPressed(ebiten.KeyShift) {
				for i := range runes {
					runes[i] = unicode.ToUpper(runes[i])
				}
			}

			b.prog.Send(tea.KeyMsg{
				Type:  val.key,
				Runes: runes,
				Alt:   ebiten.IsKeyPressed(ebiten.KeyAlt),
			})
		}
	}
}

func (b *Adapter) HandleWindowSize(size crt.WindowSize) {
	b.prog.Send(tea.WindowSizeMsg{
		Width:  size.Width,
		Height: size.Height,
	})
}
