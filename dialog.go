package termdialog

import (
	"github.com/nsf/termbox-go"
)

type Dialog interface {
	// Stuff implemented by BaseDialog
	GetTitle() string
	SetTitle(string)
	GetMetricsDirty() bool
	SetMetricsDirty(bool)
	GetWidth() int
	SetWidth(int)
	GetHeight() int
	SetHeight(int)
	GetX() int
	SetX(int)
	GetY() int
	SetY(int)
	GetTheme() *Theme
	SetTheme(*Theme)
	GetLastDialogStack() *DialogStack
	SetLastDialogStack(*DialogStack)
	Close()

	// Stuff implemented by subtypes
	CalcMetrics()
	Open()
	HandleEvent(termbox.Event) (bool, bool)
	HandleGlobalEvent(termbox.Event) (bool, bool)
}

type BaseDialog struct {
	title           string
	metricsDirty    bool
	width           int
	height          int
	x               int
	y               int
	theme           *Theme
	lastDialogStack *DialogStack
}

func (dialog *BaseDialog) GetTitle() (title string) {
	return dialog.title
}

func (dialog *BaseDialog) SetTitle(title string) {
	dialog.title = title
	dialog.metricsDirty = true
}

func (dialog *BaseDialog) GetMetricsDirty() (dirty bool) {
	return dialog.metricsDirty
}

func (dialog *BaseDialog) SetMetricsDirty(dirty bool) {
	dialog.metricsDirty = dirty
}

func (dialog *BaseDialog) GetWidth() (width int) {
	return dialog.width
}

func (dialog *BaseDialog) SetWidth(width int) {
	dialog.width = width
	dialog.metricsDirty = true
}

func (dialog *BaseDialog) GetHeight() (height int) {
	return dialog.height
}

func (dialog *BaseDialog) SetHeight(height int) {
	dialog.height = height
	dialog.metricsDirty = true
}

func (dialog *BaseDialog) GetX() (x int) {
	return dialog.x
}

func (dialog *BaseDialog) SetX(x int) {
	dialog.x = x
}

func (dialog *BaseDialog) GetY() (y int) {
	return dialog.y
}

func (dialog *BaseDialog) SetY(y int) {
	dialog.y = y
}

func (dialog *BaseDialog) GetTheme() (theme *Theme) {
	return dialog.theme
}

func (dialog *BaseDialog) SetTheme(theme *Theme) {
	dialog.theme = theme
}

func (dialog *BaseDialog) GetLastDialogStack() (lastDialogStack *DialogStack) {
	return dialog.lastDialogStack
}

func (dialog *BaseDialog) SetLastDialogStack(lastDialogStack *DialogStack) {
	dialog.lastDialogStack = lastDialogStack
}

func (dialog *BaseDialog) Close() {
	if dialog.metricsDirty {
		dialog.CalcMetrics()
	}

	Fill(dialog.x, dialog.y, dialog.width, dialog.height, ' ', dialog.theme.Screen)
}

func (dialog *BaseDialog) CalcMetrics() {

}

func (dialog *BaseDialog) HandleGlobalEvent(event termbox.Event) (handled bool, shouldClose bool) {
	return false, false
}

func BaseDialogOpen(dialog Dialog) {
	if dialog.GetMetricsDirty() {
		dialog.CalcMetrics()
	}

	title := dialog.GetTitle()
	x := dialog.GetX()
	y := dialog.GetY()
	width := dialog.GetWidth()
	height := dialog.GetHeight()
	theme := dialog.GetTheme()

	if theme.HasShadow {
		Fill(x+theme.ShadowOffsetX, y+theme.ShadowOffsetY, width, height, ' ', theme.Shadow)
	}

	DrawBox(x, y, width, height, theme.Border)
	Fill(x+1, y+1, width-2, height-2, ' ', theme.Dialog)

	DrawString(x+3, y+2, title, theme.Title)
}

func BaseDialogHandleEvent(dialog Dialog, event termbox.Event) (handled bool, shouldClose bool) {
	switch event.Type {
	case termbox.EventKey:
		if event.Ch == 0 {
			switch event.Key {
			case termbox.KeyEsc:
				return true, true

			case termbox.KeyF1:
				dialog.GetLastDialogStack().Open(HelpDialog)
				return true, false
			}

		} else {

		}
	}

	return false, false
}
