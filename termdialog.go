// Package termdialog provides a dialog-based GUI system based on Termbox.
package termdialog

import (
	"fmt"
	"github.com/nsf/termbox-go"
)

const (
	BOX_HOZ       rune = 0x2500 // Horizontal line
	BOX_VERT      rune = 0x2502 // Vertical line
	BOX_CORNER_TL rune = 0x250C // Top-left corner
	BOX_CORNER_TR rune = 0x2510 // Top-right corner
	BOX_CORNER_BL rune = 0x2514 // Bottom-left corner
	BOX_CORNER_BR rune = 0x2518 // Bottom-right corner
	BOX_TEE_L     rune = 0x251C // Left tee (i.e. on the left side of a box)
	BOX_TEE_R     rune = 0x2524 // Right tee
	BOX_TEE_T     rune = 0x252C // Top tee
	BOX_TEE_B     rune = 0x2534 // Bottom tee
	BOX_CROSS     rune = 0x253C // Cross
)

// Function DrawBox draws a box on the screen.
func DrawBox(x int, y int, width int, height int, style Style) {
	fg := style.FG
	bg := style.BG

	xmax := x + width - 1
	ymax := y + height - 1

	termbox.SetCell(x, y, BOX_CORNER_TL, fg, bg)
	termbox.SetCell(xmax, y, BOX_CORNER_TR, fg, bg)
	termbox.SetCell(x, ymax, BOX_CORNER_BL, fg, bg)
	termbox.SetCell(xmax, ymax, BOX_CORNER_BR, fg, bg)

	for i := x + 1; i <= xmax-1; i++ {
		termbox.SetCell(i, y, BOX_HOZ, fg, bg)
		termbox.SetCell(i, ymax, BOX_HOZ, fg, bg)
	}

	for i := y + 1; i <= ymax-1; i++ {
		termbox.SetCell(x, i, BOX_VERT, fg, bg)
		termbox.SetCell(xmax, i, BOX_VERT, fg, bg)
	}
}

// Function DrawString draws the specified text onto the screen.
func DrawString(x int, y int, str string, style Style) {
	startX := x

	for _, c := range str {
		if c == '\r' {
			x = startX
		} else if c == '\n' {
			y++
		} else {
			termbox.SetCell(x, y, c, style.FG, style.BG)
			x++
		}
	}
}

// Function Fill fills a region of the screen with the specified character and attribute.
func Fill(x int, y int, width int, height int, ch rune, style Style) {
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			termbox.SetCell(x+i, y+j, ch, style.FG, style.BG)
		}
	}
}

func Debug(y int, s string, args ...interface{}) {
	DrawString(0, y, fmt.Sprintf("D: "+s, args...), DefaultTheme.InactiveItem)
}
