package termdialog

import (
	"github.com/nsf/termbox-go"
)

// Type Style represents a pair of foreground and background attributes.
type Style struct {
	FG termbox.Attribute
	BG termbox.Attribute
}

func (style Style) Reverse() (newStyle Style) {
	return Style{style.BG, style.FG}
}

// Type Theme represents a GUI theme.
type Theme struct {
	Screen       Style // The style for the empty background region.
	Shadow       Style // The style for the shadow of dialogs (if enabled).
	Border       Style // The style for the border of dialogs.
	Dialog       Style // The style for the empty background of dialogs.
	Title        Style // The style for the title text of dialogs.
	InactiveItem Style // The style for inactive items and static text on dialogs.
	ActiveItem   Style // The style for active items and widgets that can be interacted with.

	HasShadow     bool // Whether to display a shadow behind dialogs. (keep this false, shadow rendering looks horrible at the moment)
	ShadowOffsetX int  // The X offset of the shadow, relative to the dialog's coordinates.
	ShadowOffsetY int  // The Y offset of the shadow, relative to the dialog's coordinates.
}

// Variable PlainTheme theme is the standard black/white theme
var PlainTheme = &Theme{
	Screen:       Style{termbox.ColorBlack, termbox.ColorBlack},
	Shadow:       Style{termbox.ColorBlack, termbox.ColorBlack},
	Border:       Style{termbox.ColorWhite, termbox.ColorBlack},
	Dialog:       Style{termbox.ColorWhite, termbox.ColorWhite},
	Title:        Style{termbox.ColorBlack | termbox.AttrUnderline, termbox.ColorWhite},
	InactiveItem: Style{termbox.ColorBlack, termbox.ColorWhite},
	ActiveItem:   Style{termbox.ColorWhite, termbox.ColorRed},

	HasShadow:     false,
	ShadowOffsetX: 2,
	ShadowOffsetY: 1,
}

// Variable WhiptailTheme is a theme that mimics whiptail dialogs
var WhiptailTheme = &Theme{
	Screen:       Style{termbox.ColorBlack, termbox.ColorBlue},
	Shadow:       Style{termbox.ColorBlack, termbox.ColorBlack},
	Border:       Style{termbox.ColorBlack, termbox.ColorWhite},
	Dialog:       Style{termbox.ColorBlack, termbox.ColorWhite},
	Title:        Style{termbox.ColorBlack | termbox.AttrUnderline, termbox.ColorWhite},
	InactiveItem: Style{termbox.ColorBlack, termbox.ColorWhite},
	ActiveItem:   Style{termbox.ColorWhite, termbox.ColorRed},

	HasShadow:     true,
	ShadowOffsetX: 1,
	ShadowOffsetY: 1,
}

var DefaultTheme = PlainTheme
