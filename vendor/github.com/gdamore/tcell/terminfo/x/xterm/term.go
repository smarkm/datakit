// Generated automatically.  DO NOT HAND-EDIT.

package xterm

import "github.com/gdamore/tcell/terminfo"

func init() {

	// X11 terminal emulator
	terminfo.AddTerminfo(&terminfo.Terminfo{
		Name:            "xterm",
		Aliases:         []string{"xterm-debian"},
		Columns:         80,
		Lines:           24,
		Colors:          8,
		Bell:            "\a",
		Clear:           "\x1b[H\x1b[2J",
		EnterCA:         "\x1b[?1049h\x1b[22;0;0t",
		ExitCA:          "\x1b[?1049l\x1b[23;0;0t",
		ShowCursor:      "\x1b[?12l\x1b[?25h",
		HideCursor:      "\x1b[?25l",
		AttrOff:         "\x1b(B\x1b[m",
		Underline:       "\x1b[4m",
		Bold:            "\x1b[1m",
		Dim:             "\x1b[2m",
		Blink:           "\x1b[5m",
		Reverse:         "\x1b[7m",
		EnterKeypad:     "\x1b[?1h\x1b=",
		ExitKeypad:      "\x1b[?1l\x1b>",
		SetFg:           "\x1b[3%p1%dm",
		SetBg:           "\x1b[4%p1%dm",
		SetFgBg:         "\x1b[3%p1%d;4%p2%dm",
		AltChars:        "``aaffggiijjkkllmmnnooppqqrrssttuuvvwwxxyyzz{{||}}~~",
		EnterAcs:        "\x1b(0",
		ExitAcs:         "\x1b(B",
		Mouse:           "\x1b[M",
		MouseMode:       "%?%p1%{1}%=%t%'h'%Pa%e%'l'%Pa%;\x1b[?1000%ga%c\x1b[?1002%ga%c\x1b[?1003%ga%c\x1b[?1006%ga%c",
		SetCursor:       "\x1b[%i%p1%d;%p2%dH",
		CursorBack1:     "\b",
		CursorUp1:       "\x1b[A",
		KeyUp:           "\x1bOA",
		KeyDown:         "\x1bOB",
		KeyRight:        "\x1bOC",
		KeyLeft:         "\x1bOD",
		KeyInsert:       "\x1b[2~",
		KeyDelete:       "\x1b[3~",
		KeyBackspace:    "\u007f",
		KeyHome:         "\x1bOH",
		KeyEnd:          "\x1bOF",
		KeyPgUp:         "\x1b[5~",
		KeyPgDn:         "\x1b[6~",
		KeyF1:           "\x1bOP",
		KeyF2:           "\x1bOQ",
		KeyF3:           "\x1bOR",
		KeyF4:           "\x1bOS",
		KeyF5:           "\x1b[15~",
		KeyF6:           "\x1b[17~",
		KeyF7:           "\x1b[18~",
		KeyF8:           "\x1b[19~",
		KeyF9:           "\x1b[20~",
		KeyF10:          "\x1b[21~",
		KeyF11:          "\x1b[23~",
		KeyF12:          "\x1b[24~",
		KeyF13:          "\x1b[1;2P",
		KeyF14:          "\x1b[1;2Q",
		KeyF15:          "\x1b[1;2R",
		KeyF16:          "\x1b[1;2S",
		KeyF17:          "\x1b[15;2~",
		KeyF18:          "\x1b[17;2~",
		KeyF19:          "\x1b[18;2~",
		KeyF20:          "\x1b[19;2~",
		KeyF21:          "\x1b[20;2~",
		KeyF22:          "\x1b[21;2~",
		KeyF23:          "\x1b[23;2~",
		KeyF24:          "\x1b[24;2~",
		KeyF25:          "\x1b[1;5P",
		KeyF26:          "\x1b[1;5Q",
		KeyF27:          "\x1b[1;5R",
		KeyF28:          "\x1b[1;5S",
		KeyF29:          "\x1b[15;5~",
		KeyF30:          "\x1b[17;5~",
		KeyF31:          "\x1b[18;5~",
		KeyF32:          "\x1b[19;5~",
		KeyF33:          "\x1b[20;5~",
		KeyF34:          "\x1b[21;5~",
		KeyF35:          "\x1b[23;5~",
		KeyF36:          "\x1b[24;5~",
		KeyF37:          "\x1b[1;6P",
		KeyF38:          "\x1b[1;6Q",
		KeyF39:          "\x1b[1;6R",
		KeyF40:          "\x1b[1;6S",
		KeyF41:          "\x1b[15;6~",
		KeyF42:          "\x1b[17;6~",
		KeyF43:          "\x1b[18;6~",
		KeyF44:          "\x1b[19;6~",
		KeyF45:          "\x1b[20;6~",
		KeyF46:          "\x1b[21;6~",
		KeyF47:          "\x1b[23;6~",
		KeyF48:          "\x1b[24;6~",
		KeyF49:          "\x1b[1;3P",
		KeyF50:          "\x1b[1;3Q",
		KeyF51:          "\x1b[1;3R",
		KeyF52:          "\x1b[1;3S",
		KeyF53:          "\x1b[15;3~",
		KeyF54:          "\x1b[17;3~",
		KeyF55:          "\x1b[18;3~",
		KeyF56:          "\x1b[19;3~",
		KeyF57:          "\x1b[20;3~",
		KeyF58:          "\x1b[21;3~",
		KeyF59:          "\x1b[23;3~",
		KeyF60:          "\x1b[24;3~",
		KeyF61:          "\x1b[1;4P",
		KeyF62:          "\x1b[1;4Q",
		KeyF63:          "\x1b[1;4R",
		KeyBacktab:      "\x1b[Z",
		KeyShfLeft:      "\x1b[1;2D",
		KeyShfRight:     "\x1b[1;2C",
		KeyShfUp:        "\x1b[1;2A",
		KeyShfDown:      "\x1b[1;2B",
		KeyCtrlLeft:     "\x1b[1;5D",
		KeyCtrlRight:    "\x1b[1;5C",
		KeyCtrlUp:       "\x1b[1;5A",
		KeyCtrlDown:     "\x1b[1;5B",
		KeyMetaLeft:     "\x1b[1;9D",
		KeyMetaRight:    "\x1b[1;9C",
		KeyMetaUp:       "\x1b[1;9A",
		KeyMetaDown:     "\x1b[1;9B",
		KeyAltLeft:      "\x1b[1;3D",
		KeyAltRight:     "\x1b[1;3C",
		KeyAltUp:        "\x1b[1;3A",
		KeyAltDown:      "\x1b[1;3B",
		KeyAltShfLeft:   "\x1b[1;4D",
		KeyAltShfRight:  "\x1b[1;4C",
		KeyAltShfUp:     "\x1b[1;4A",
		KeyAltShfDown:   "\x1b[1;4B",
		KeyMetaShfLeft:  "\x1b[1;10D",
		KeyMetaShfRight: "\x1b[1;10C",
		KeyMetaShfUp:    "\x1b[1;10A",
		KeyMetaShfDown:  "\x1b[1;10B",
		KeyCtrlShfLeft:  "\x1b[1;6D",
		KeyCtrlShfRight: "\x1b[1;6C",
		KeyCtrlShfUp:    "\x1b[1;6A",
		KeyCtrlShfDown:  "\x1b[1;6B",
		KeyShfHome:      "\x1b[1;2H",
		KeyShfEnd:       "\x1b[1;2F",
		KeyCtrlHome:     "\x1b[1;5H",
		KeyCtrlEnd:      "\x1b[1;5F",
		KeyAltHome:      "\x1b[1;9H",
		KeyAltEnd:       "\x1b[1;9F",
		KeyCtrlShfHome:  "\x1b[1;6H",
		KeyCtrlShfEnd:   "\x1b[1;6F",
		KeyMetaShfHome:  "\x1b[1;10H",
		KeyMetaShfEnd:   "\x1b[1;10F",
		KeyAltShfHome:   "\x1b[1;4H",
		KeyAltShfEnd:    "\x1b[1;4F",
	})

	// xterm with 88 colors
	terminfo.AddTerminfo(&terminfo.Terminfo{
		Name:            "xterm-88color",
		Columns:         80,
		Lines:           24,
		Colors:          88,
		Bell:            "\a",
		Clear:           "\x1b[H\x1b[2J",
		EnterCA:         "\x1b[?1049h\x1b[22;0;0t",
		ExitCA:          "\x1b[?1049l\x1b[23;0;0t",
		ShowCursor:      "\x1b[?12l\x1b[?25h",
		HideCursor:      "\x1b[?25l",
		AttrOff:         "\x1b(B\x1b[m",
		Underline:       "\x1b[4m",
		Bold:            "\x1b[1m",
		Dim:             "\x1b[2m",
		Blink:           "\x1b[5m",
		Reverse:         "\x1b[7m",
		EnterKeypad:     "\x1b[?1h\x1b=",
		ExitKeypad:      "\x1b[?1l\x1b>",
		SetFg:           "\x1b[%?%p1%{8}%<%t3%p1%d%e%p1%{16}%<%t9%p1%{8}%-%d%e38;5;%p1%d%;m",
		SetBg:           "\x1b[%?%p1%{8}%<%t4%p1%d%e%p1%{16}%<%t10%p1%{8}%-%d%e48;5;%p1%d%;m",
		SetFgBg:         "\x1b[%?%p1%{8}%<%t3%p1%d%e%p1%{16}%<%t9%p1%{8}%-%d%e38;5;%p1%d%;;%?%p2%{8}%<%t4%p2%d%e%p2%{16}%<%t10%p2%{8}%-%d%e48;5;%p2%d%;m",
		AltChars:        "``aaffggiijjkkllmmnnooppqqrrssttuuvvwwxxyyzz{{||}}~~",
		EnterAcs:        "\x1b(0",
		ExitAcs:         "\x1b(B",
		Mouse:           "\x1b[M",
		MouseMode:       "%?%p1%{1}%=%t%'h'%Pa%e%'l'%Pa%;\x1b[?1000%ga%c\x1b[?1002%ga%c\x1b[?1003%ga%c\x1b[?1006%ga%c",
		SetCursor:       "\x1b[%i%p1%d;%p2%dH",
		CursorBack1:     "\b",
		CursorUp1:       "\x1b[A",
		KeyUp:           "\x1bOA",
		KeyDown:         "\x1bOB",
		KeyRight:        "\x1bOC",
		KeyLeft:         "\x1bOD",
		KeyInsert:       "\x1b[2~",
		KeyDelete:       "\x1b[3~",
		KeyBackspace:    "\u007f",
		KeyHome:         "\x1bOH",
		KeyEnd:          "\x1bOF",
		KeyPgUp:         "\x1b[5~",
		KeyPgDn:         "\x1b[6~",
		KeyF1:           "\x1bOP",
		KeyF2:           "\x1bOQ",
		KeyF3:           "\x1bOR",
		KeyF4:           "\x1bOS",
		KeyF5:           "\x1b[15~",
		KeyF6:           "\x1b[17~",
		KeyF7:           "\x1b[18~",
		KeyF8:           "\x1b[19~",
		KeyF9:           "\x1b[20~",
		KeyF10:          "\x1b[21~",
		KeyF11:          "\x1b[23~",
		KeyF12:          "\x1b[24~",
		KeyF13:          "\x1b[1;2P",
		KeyF14:          "\x1b[1;2Q",
		KeyF15:          "\x1b[1;2R",
		KeyF16:          "\x1b[1;2S",
		KeyF17:          "\x1b[15;2~",
		KeyF18:          "\x1b[17;2~",
		KeyF19:          "\x1b[18;2~",
		KeyF20:          "\x1b[19;2~",
		KeyF21:          "\x1b[20;2~",
		KeyF22:          "\x1b[21;2~",
		KeyF23:          "\x1b[23;2~",
		KeyF24:          "\x1b[24;2~",
		KeyF25:          "\x1b[1;5P",
		KeyF26:          "\x1b[1;5Q",
		KeyF27:          "\x1b[1;5R",
		KeyF28:          "\x1b[1;5S",
		KeyF29:          "\x1b[15;5~",
		KeyF30:          "\x1b[17;5~",
		KeyF31:          "\x1b[18;5~",
		KeyF32:          "\x1b[19;5~",
		KeyF33:          "\x1b[20;5~",
		KeyF34:          "\x1b[21;5~",
		KeyF35:          "\x1b[23;5~",
		KeyF36:          "\x1b[24;5~",
		KeyF37:          "\x1b[1;6P",
		KeyF38:          "\x1b[1;6Q",
		KeyF39:          "\x1b[1;6R",
		KeyF40:          "\x1b[1;6S",
		KeyF41:          "\x1b[15;6~",
		KeyF42:          "\x1b[17;6~",
		KeyF43:          "\x1b[18;6~",
		KeyF44:          "\x1b[19;6~",
		KeyF45:          "\x1b[20;6~",
		KeyF46:          "\x1b[21;6~",
		KeyF47:          "\x1b[23;6~",
		KeyF48:          "\x1b[24;6~",
		KeyF49:          "\x1b[1;3P",
		KeyF50:          "\x1b[1;3Q",
		KeyF51:          "\x1b[1;3R",
		KeyF52:          "\x1b[1;3S",
		KeyF53:          "\x1b[15;3~",
		KeyF54:          "\x1b[17;3~",
		KeyF55:          "\x1b[18;3~",
		KeyF56:          "\x1b[19;3~",
		KeyF57:          "\x1b[20;3~",
		KeyF58:          "\x1b[21;3~",
		KeyF59:          "\x1b[23;3~",
		KeyF60:          "\x1b[24;3~",
		KeyF61:          "\x1b[1;4P",
		KeyF62:          "\x1b[1;4Q",
		KeyF63:          "\x1b[1;4R",
		KeyBacktab:      "\x1b[Z",
		KeyShfLeft:      "\x1b[1;2D",
		KeyShfRight:     "\x1b[1;2C",
		KeyShfUp:        "\x1b[1;2A",
		KeyShfDown:      "\x1b[1;2B",
		KeyCtrlLeft:     "\x1b[1;5D",
		KeyCtrlRight:    "\x1b[1;5C",
		KeyCtrlUp:       "\x1b[1;5A",
		KeyCtrlDown:     "\x1b[1;5B",
		KeyMetaLeft:     "\x1b[1;9D",
		KeyMetaRight:    "\x1b[1;9C",
		KeyMetaUp:       "\x1b[1;9A",
		KeyMetaDown:     "\x1b[1;9B",
		KeyAltLeft:      "\x1b[1;3D",
		KeyAltRight:     "\x1b[1;3C",
		KeyAltUp:        "\x1b[1;3A",
		KeyAltDown:      "\x1b[1;3B",
		KeyAltShfLeft:   "\x1b[1;4D",
		KeyAltShfRight:  "\x1b[1;4C",
		KeyAltShfUp:     "\x1b[1;4A",
		KeyAltShfDown:   "\x1b[1;4B",
		KeyMetaShfLeft:  "\x1b[1;10D",
		KeyMetaShfRight: "\x1b[1;10C",
		KeyMetaShfUp:    "\x1b[1;10A",
		KeyMetaShfDown:  "\x1b[1;10B",
		KeyCtrlShfLeft:  "\x1b[1;6D",
		KeyCtrlShfRight: "\x1b[1;6C",
		KeyCtrlShfUp:    "\x1b[1;6A",
		KeyCtrlShfDown:  "\x1b[1;6B",
		KeyShfHome:      "\x1b[1;2H",
		KeyShfEnd:       "\x1b[1;2F",
		KeyCtrlHome:     "\x1b[1;5H",
		KeyCtrlEnd:      "\x1b[1;5F",
		KeyAltHome:      "\x1b[1;9H",
		KeyAltEnd:       "\x1b[1;9F",
		KeyCtrlShfHome:  "\x1b[1;6H",
		KeyCtrlShfEnd:   "\x1b[1;6F",
		KeyMetaShfHome:  "\x1b[1;10H",
		KeyMetaShfEnd:   "\x1b[1;10F",
		KeyAltShfHome:   "\x1b[1;4H",
		KeyAltShfEnd:    "\x1b[1;4F",
	})

	// xterm with 256 colors
	terminfo.AddTerminfo(&terminfo.Terminfo{
		Name:            "xterm-256color",
		Columns:         80,
		Lines:           24,
		Colors:          256,
		Bell:            "\a",
		Clear:           "\x1b[H\x1b[2J",
		EnterCA:         "\x1b[?1049h\x1b[22;0;0t",
		ExitCA:          "\x1b[?1049l\x1b[23;0;0t",
		ShowCursor:      "\x1b[?12l\x1b[?25h",
		HideCursor:      "\x1b[?25l",
		AttrOff:         "\x1b(B\x1b[m",
		Underline:       "\x1b[4m",
		Bold:            "\x1b[1m",
		Dim:             "\x1b[2m",
		Blink:           "\x1b[5m",
		Reverse:         "\x1b[7m",
		EnterKeypad:     "\x1b[?1h\x1b=",
		ExitKeypad:      "\x1b[?1l\x1b>",
		SetFg:           "\x1b[%?%p1%{8}%<%t3%p1%d%e%p1%{16}%<%t9%p1%{8}%-%d%e38;5;%p1%d%;m",
		SetBg:           "\x1b[%?%p1%{8}%<%t4%p1%d%e%p1%{16}%<%t10%p1%{8}%-%d%e48;5;%p1%d%;m",
		SetFgBg:         "\x1b[%?%p1%{8}%<%t3%p1%d%e%p1%{16}%<%t9%p1%{8}%-%d%e38;5;%p1%d%;;%?%p2%{8}%<%t4%p2%d%e%p2%{16}%<%t10%p2%{8}%-%d%e48;5;%p2%d%;m",
		AltChars:        "``aaffggiijjkkllmmnnooppqqrrssttuuvvwwxxyyzz{{||}}~~",
		EnterAcs:        "\x1b(0",
		ExitAcs:         "\x1b(B",
		Mouse:           "\x1b[M",
		MouseMode:       "%?%p1%{1}%=%t%'h'%Pa%e%'l'%Pa%;\x1b[?1000%ga%c\x1b[?1002%ga%c\x1b[?1003%ga%c\x1b[?1006%ga%c",
		SetCursor:       "\x1b[%i%p1%d;%p2%dH",
		CursorBack1:     "\b",
		CursorUp1:       "\x1b[A",
		KeyUp:           "\x1bOA",
		KeyDown:         "\x1bOB",
		KeyRight:        "\x1bOC",
		KeyLeft:         "\x1bOD",
		KeyInsert:       "\x1b[2~",
		KeyDelete:       "\x1b[3~",
		KeyBackspace:    "\u007f",
		KeyHome:         "\x1bOH",
		KeyEnd:          "\x1bOF",
		KeyPgUp:         "\x1b[5~",
		KeyPgDn:         "\x1b[6~",
		KeyF1:           "\x1bOP",
		KeyF2:           "\x1bOQ",
		KeyF3:           "\x1bOR",
		KeyF4:           "\x1bOS",
		KeyF5:           "\x1b[15~",
		KeyF6:           "\x1b[17~",
		KeyF7:           "\x1b[18~",
		KeyF8:           "\x1b[19~",
		KeyF9:           "\x1b[20~",
		KeyF10:          "\x1b[21~",
		KeyF11:          "\x1b[23~",
		KeyF12:          "\x1b[24~",
		KeyF13:          "\x1b[1;2P",
		KeyF14:          "\x1b[1;2Q",
		KeyF15:          "\x1b[1;2R",
		KeyF16:          "\x1b[1;2S",
		KeyF17:          "\x1b[15;2~",
		KeyF18:          "\x1b[17;2~",
		KeyF19:          "\x1b[18;2~",
		KeyF20:          "\x1b[19;2~",
		KeyF21:          "\x1b[20;2~",
		KeyF22:          "\x1b[21;2~",
		KeyF23:          "\x1b[23;2~",
		KeyF24:          "\x1b[24;2~",
		KeyF25:          "\x1b[1;5P",
		KeyF26:          "\x1b[1;5Q",
		KeyF27:          "\x1b[1;5R",
		KeyF28:          "\x1b[1;5S",
		KeyF29:          "\x1b[15;5~",
		KeyF30:          "\x1b[17;5~",
		KeyF31:          "\x1b[18;5~",
		KeyF32:          "\x1b[19;5~",
		KeyF33:          "\x1b[20;5~",
		KeyF34:          "\x1b[21;5~",
		KeyF35:          "\x1b[23;5~",
		KeyF36:          "\x1b[24;5~",
		KeyF37:          "\x1b[1;6P",
		KeyF38:          "\x1b[1;6Q",
		KeyF39:          "\x1b[1;6R",
		KeyF40:          "\x1b[1;6S",
		KeyF41:          "\x1b[15;6~",
		KeyF42:          "\x1b[17;6~",
		KeyF43:          "\x1b[18;6~",
		KeyF44:          "\x1b[19;6~",
		KeyF45:          "\x1b[20;6~",
		KeyF46:          "\x1b[21;6~",
		KeyF47:          "\x1b[23;6~",
		KeyF48:          "\x1b[24;6~",
		KeyF49:          "\x1b[1;3P",
		KeyF50:          "\x1b[1;3Q",
		KeyF51:          "\x1b[1;3R",
		KeyF52:          "\x1b[1;3S",
		KeyF53:          "\x1b[15;3~",
		KeyF54:          "\x1b[17;3~",
		KeyF55:          "\x1b[18;3~",
		KeyF56:          "\x1b[19;3~",
		KeyF57:          "\x1b[20;3~",
		KeyF58:          "\x1b[21;3~",
		KeyF59:          "\x1b[23;3~",
		KeyF60:          "\x1b[24;3~",
		KeyF61:          "\x1b[1;4P",
		KeyF62:          "\x1b[1;4Q",
		KeyF63:          "\x1b[1;4R",
		KeyBacktab:      "\x1b[Z",
		KeyShfLeft:      "\x1b[1;2D",
		KeyShfRight:     "\x1b[1;2C",
		KeyShfUp:        "\x1b[1;2A",
		KeyShfDown:      "\x1b[1;2B",
		KeyCtrlLeft:     "\x1b[1;5D",
		KeyCtrlRight:    "\x1b[1;5C",
		KeyCtrlUp:       "\x1b[1;5A",
		KeyCtrlDown:     "\x1b[1;5B",
		KeyMetaLeft:     "\x1b[1;9D",
		KeyMetaRight:    "\x1b[1;9C",
		KeyMetaUp:       "\x1b[1;9A",
		KeyMetaDown:     "\x1b[1;9B",
		KeyAltLeft:      "\x1b[1;3D",
		KeyAltRight:     "\x1b[1;3C",
		KeyAltUp:        "\x1b[1;3A",
		KeyAltDown:      "\x1b[1;3B",
		KeyAltShfLeft:   "\x1b[1;4D",
		KeyAltShfRight:  "\x1b[1;4C",
		KeyAltShfUp:     "\x1b[1;4A",
		KeyAltShfDown:   "\x1b[1;4B",
		KeyMetaShfLeft:  "\x1b[1;10D",
		KeyMetaShfRight: "\x1b[1;10C",
		KeyMetaShfUp:    "\x1b[1;10A",
		KeyMetaShfDown:  "\x1b[1;10B",
		KeyCtrlShfLeft:  "\x1b[1;6D",
		KeyCtrlShfRight: "\x1b[1;6C",
		KeyCtrlShfUp:    "\x1b[1;6A",
		KeyCtrlShfDown:  "\x1b[1;6B",
		KeyShfHome:      "\x1b[1;2H",
		KeyShfEnd:       "\x1b[1;2F",
		KeyCtrlHome:     "\x1b[1;5H",
		KeyCtrlEnd:      "\x1b[1;5F",
		KeyAltHome:      "\x1b[1;9H",
		KeyAltEnd:       "\x1b[1;9F",
		KeyCtrlShfHome:  "\x1b[1;6H",
		KeyCtrlShfEnd:   "\x1b[1;6F",
		KeyMetaShfHome:  "\x1b[1;10H",
		KeyMetaShfEnd:   "\x1b[1;10F",
		KeyAltShfHome:   "\x1b[1;4H",
		KeyAltShfEnd:    "\x1b[1;4F",
	})
}
