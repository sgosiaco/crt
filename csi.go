package crt

import (
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/muesli/termenv"
)

var csiMtx = &sync.Mutex{}
var csiCache = map[string]any{}

type CursorUpSeq struct {
	Count int
}

type CursorDownSeq struct {
	Count int
}

type CursorForwardSeq struct {
	Count int
}

type CursorBackSeq struct {
	Count int
}

type CursorNextLineSeq struct {
	Count int
}

type CursorPreviousLineSeq struct {
	Count int
}

type CursorHorizontalSeq struct {
	Count int
}

type CursorPositionSeq struct {
	Row int
	Col int
}

type EraseDisplaySeq struct {
	Type int
}

type EraseLineSeq struct {
	Type int
}

type ScrollUpSeq struct {
	Count int
}

type ScrollDownSeq struct {
	Count int
}

type SaveCursorPositionSeq struct{}

type RestoreCursorPositionSeq struct{}

type ChangeScrollingRegionSeq struct {
	Top    int
	Bottom int
}

type InsertLineSeq struct {
	Count int
}

type DeleteLineSeq struct {
	Count int
}

type CursorShowSeq struct{}

type CursorHideSeq struct{}

// extractCSI extracts a CSI sequence from the beginning of a string.
// It returns the sequence without any suffix, and a boolean indicating
// whether a sequence was found.
func extractCSI(s string) (string, bool) {
	if !strings.HasPrefix(s, termenv.CSI) {
		return "", false
	}

	s = s[len(termenv.CSI):]
	if len(s) == 0 {
		return "", false
	}

	for i, c := range s {
		if c >= '@' && c <= '~' {
			return termenv.CSI + s[:i+1], true
		}
	}

	return "", false
}

// parseCSI parses a CSI sequence and returns a struct representing the sequence.
func parseCSI(s string) (any, bool) {
	if !strings.HasPrefix(s, termenv.CSI) {
		return nil, false
	}

	s = s[len(termenv.CSI):]
	if len(s) == 0 {
		return nil, false
	}

	csiMtx.Lock()
	if cached, ok := csiCache[s]; ok {
		csiMtx.Unlock()
		return cached, true
	}
	csiMtx.Unlock()

	if val, ok := parseCSIStruct(s); ok {
		csiMtx.Lock()
		csiCache[s] = val
		csiMtx.Unlock()

		return val, true
	}

	return nil, false
}

func parseCSIStruct(s string) (any, bool) {
	switch s {
	case termenv.ShowCursorSeq:
		return CursorShowSeq{}, true
	case termenv.HideCursorSeq:
		return CursorHideSeq{}, true
	}

	fmt.Println("ESC", fmt.Sprintf(`"%s"`, s))

	switch s[len(s)-1] {
	case 'A':
		return CursorUpSeq{Count: getCount(s, 1)}, true
	case 'B':
		return CursorDownSeq{Count: getCount(s, 1)}, true
	case 'C':
		return CursorForwardSeq{Count: getCount(s, 1)}, true
	case 'D':
		return CursorBackSeq{Count: getCount(s, 1)}, true
	case 'E':
		return CursorNextLineSeq{Count: getCount(s, 1)}, true
	case 'F':
		return CursorPreviousLineSeq{Count: getCount(s, 1)}, true
	case 'G':
		return CursorHorizontalSeq{Count: getCount(s, 1)}, true
	case 'H', 'f': // Note: f is the same but diff
		// if no semicolon, then it should be just "H" (1,1)
		if !strings.Contains(s, ";") {
			return CursorPositionSeq{Row: 1, Col: 1}, true
		}

		// extract "prefix" ("2;2", ";2", "2;")
		prefix := s[:len(s)-1]
		parts := strings.Split(prefix, ";")

		switch len(parts) {
		case 1: // if we have exactly one, either it's "r;" or ";c"
			row := 1
			col := 1

			conv, err := strconv.Atoi(parts[0])
			if err != nil {
				return nil, false
			}

			// if the "prefix" starts with ; then we know column was provided, otherwise row
			if strings.HasPrefix(prefix, ";") {
				col = conv
			} else {
				row = conv
			}

			return CursorPositionSeq{Row: row, Col: col}, true
		case 2:
			row, err := strconv.Atoi(parts[0])
			if err != nil {
				return nil, false
			}
			col, err := strconv.Atoi(parts[1])
			if err != nil {
				return nil, false
			}
			return CursorPositionSeq{Row: row, Col: col}, true
		default:
			return nil, false
		}
	case 'J':
		// if it's just "J" then same as "0J"
		// so default 0
		return EraseDisplaySeq{Type: getCount(s, 0)}, true
	case 'K':
		// if it's just "K" then same as "0K"
		// so default 0
		return EraseLineSeq{Type: getCount(s, 0)}, true
	case 'S':
		return ScrollUpSeq{Count: getCount(s, 1)}, true
	case 'T':
		return ScrollDownSeq{Count: getCount(s, 1)}, true
	case 's':
		if len(s) == 1 {
			return SaveCursorPositionSeq{}, true
		}
	case 'u':
		if len(s) == 1 {
			return RestoreCursorPositionSeq{}, true
		}
	case 'r':
		// TODO: implement
	case 'L':
		return InsertLineSeq{Count: getCount(s, 1)}, true
	case 'M':
		return DeleteLineSeq{Count: getCount(s, 1)}, true
	}

	fmt.Println("UNKNOWN ESC", fmt.Sprintf(`"%s"`, s))

	return nil, false
}

func getCount(s string, defaultVal int) int {
	count, err := strconv.Atoi(s[:len(s)-1])
	if err == nil {
		return count
	}

	if len(s) == 1 {
		return defaultVal
	}

	return 0
}
