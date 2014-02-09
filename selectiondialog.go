package termdialog

import (
	"fmt"
	"github.com/nsf/termbox-go"
)

/*
  +---------+
  |         |
  |  Title  |
  |         |
  |  xxx    |
  |  yyy    |
  |  zzz    |
  |         |
  +---------+
*/

// Type Option represents an option in a selection dialog.
type Option struct {
	Text     string             // The text of the option.
	Callback func(*Option) bool // The callback.
	Data     interface{}        // Arbitary associated data that can be accessed by the callback.
}

// Type SelectionDialog represents a dialog with a number of selectable options.
type SelectionDialog struct {
	BaseDialog
	options           []*Option
	selectedIndex     int
	maxVisibleOptions int // if > 0, scrolling activated, window height limited to the value
	topIndex          int
}

// Function NewSelectionDialog creates and returns a new selection dialog. The options argument can
// be nil, and an empty slice will be created (options can be added later with AddOption).
// Set maxVisibleOptions to get scrolling feature. Designed for compatibility.
func NewSelectionDialog(title string, options []*Option, maxVisibleOptions ...int) (dialog *SelectionDialog) {
	if options == nil {
		options = make([]*Option, 0)
	}

	dialog = &SelectionDialog{
		options: options,
	}

	if maxVisibleOptions != nil {
		dialog.maxVisibleOptions = maxVisibleOptions[0]
	}

	dialog.BaseDialog.title = title
	dialog.BaseDialog.metricsDirty = true
	dialog.BaseDialog.theme = DefaultTheme
	return dialog
}

// Function NOptions returns the number of options attached to this dialog.
func (dialog *SelectionDialog) NOptions() (num int) {
	return len(dialog.options)
}

func (dialog *SelectionDialog) GetOption(n int) (option *Option) {
	return dialog.options[n]
}

func (dialog *SelectionDialog) SetOption(n int, option *Option) {
	dialog.options[n] = option
	dialog.metricsDirty = true
}

func (dialog *SelectionDialog) AddOption(option *Option) (theSameOption *Option) {
	dialog.options = append(dialog.options, option)
	dialog.metricsDirty = true
	return option
}

func (dialog *SelectionDialog) RemoveOption(n int) {
	dialog.options = append(dialog.options[:n], dialog.options[n+1:]...)
	dialog.metricsDirty = true
}

func (dialog *SelectionDialog) FindOption(option *Option) (n int) {
	for n, o := range dialog.options {
		if o == option {
			return n
		}
	}
	return -1
}

func (dialog *SelectionDialog) ClearOptions() {
	dialog.options = make([]*Option, 0)
	dialog.topIndex = 0
	dialog.selectedIndex = 0
	dialog.metricsDirty = true
}

func (dialog *SelectionDialog) GetSelectedIndex() (selectedIndex int) {
	return dialog.selectedIndex
}

func (dialog *SelectionDialog) GetSelectedOption() (option *Option) {
	return dialog.options[dialog.selectedIndex]
}

func (dialog *SelectionDialog) SetSelectedIndex(selectedIndex int) {
	dialog.selectedIndex = selectedIndex
	dialog.metricsDirty = true // ?
}

func (dialog *SelectionDialog) SetSelectionOption(option *Option) {
	dialog.selectedIndex = dialog.FindOption(option)
	dialog.metricsDirty = true //?
}

func (dialog *SelectionDialog) CalcMetrics() {
	windowWidth, windowHeight := termbox.Size()

	maxWidth := 0
	for _, option := range dialog.options {
		if len(option.Text) > maxWidth {
			maxWidth = len(option.Text)
		}
	}

	maxWidth += 2 // Add the "* "
	if len(dialog.title) > maxWidth {
		maxWidth = len(dialog.title)
	}

	dialog.width = 6 + maxWidth // 6 = "|  " + "  |"
	dialog.height += 6          // 6 = Top border, Top padding, Title, Under-title padding, Bottom padding, Bottom border

	if dialog.maxVisibleOptions > 0 {
		dialog.height += dialog.maxVisibleOptions
	} else {
		dialog.height += len(dialog.options)
	}

	dialog.x = (windowWidth / 2) - (dialog.width / 2)
	dialog.y = (windowHeight / 2) - (dialog.height / 2)

	dialog.metricsDirty = false
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func (dialog *SelectionDialog) Open() {
	BaseDialogOpen(dialog)

	if len(dialog.options) <= dialog.maxVisibleOptions {
		dialog.topIndex = 0
	}

	cnt := len(dialog.options)

	if dialog.maxVisibleOptions > 0 {
		cnt = min(cnt, dialog.topIndex+dialog.maxVisibleOptions)
	}

	k := 0
	for i := dialog.topIndex; i < cnt; i++ {
		style := dialog.theme.InactiveItem

		if i == dialog.selectedIndex {
			style = dialog.theme.ActiveItem
		}

		DrawString(dialog.x+3, dialog.y+4+k, fmt.Sprintf("* %s", dialog.options[i].Text), style)
		k++
	}
}

func (dialog *SelectionDialog) HandleEvent(event termbox.Event) (handled bool, shouldClose bool) {
	handled, shouldClose = BaseDialogHandleEvent(dialog, event)
	if handled {
		return
	}

	maxOption := len(dialog.options) - 1

	switch event.Type {
	case termbox.EventKey:
		switch event.Key {
		case termbox.KeyArrowUp:
			if dialog.selectedIndex > 0 {
				dialog.selectedIndex--
				if (dialog.maxVisibleOptions > 0) && (dialog.selectedIndex < dialog.topIndex) {
					dialog.topIndex--
				}

			}

			return true, false

		case termbox.KeyArrowDown:
			if dialog.selectedIndex < maxOption {
				dialog.selectedIndex++

				if (dialog.maxVisibleOptions > 0) && (dialog.selectedIndex == dialog.topIndex+dialog.maxVisibleOptions) {
					dialog.topIndex++
				}
			}

			return true, false

		case termbox.KeyHome:
			dialog.selectedIndex = 0
			dialog.topIndex = 0
			return true, false

		case termbox.KeyEnd:
			dialog.selectedIndex = maxOption
			if dialog.maxVisibleOptions > 0 {
				dialog.topIndex = maxOption - dialog.maxVisibleOptions + 1
			}
			return true, false

		case termbox.KeyEnter, termbox.KeySpace:
			option := dialog.options[dialog.selectedIndex]
			shouldClose = true
			if option.Callback != nil {
				shouldClose = option.Callback(option)
			}

			return true, shouldClose
		}
	}

	return false, false
}
