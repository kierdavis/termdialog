package termdialog

import (
	"github.com/nsf/termbox-go"
)

type DialogStack struct {
	dialogs []Dialog
}

func NewDialogStack() (dialogStack *DialogStack) {
	return &DialogStack{
		dialogs: nil,
	}
}

func (dialogStack *DialogStack) Open(dialog Dialog) {
	dialog.Open()
	dialog.SetLastDialogStack(dialogStack)
	dialogStack.dialogs = append(dialogStack.dialogs, dialog)
	//return dialog
}

func (dialogStack *DialogStack) Close(dialog Dialog) {
	dialog.Close()

	for i, d := range dialogStack.dialogs {
		if d == dialog {
			dialogStack.dialogs = append(dialogStack.dialogs[:i], dialogStack.dialogs[i+1:]...)
		}
	}
}

func (dialogStack *DialogStack) CloseTop() {
	dialog := dialogStack.dialogs[len(dialogStack.dialogs)-1]
	dialog.Close()
	dialogStack.dialogs = dialogStack.dialogs[:len(dialogStack.dialogs)-1]
	//return dialog
}

func (dialogStack *DialogStack) Run() {
	_, windowHeight := termbox.Size()

	//Fill(0, 0, windowWidth, windowHeight, ' ', DefaultTheme.Screen.FG, DefaultTheme.Screen.BG)
	termbox.Clear(DefaultTheme.Screen.FG, DefaultTheme.Screen.BG)
	DrawString(0, windowHeight-1, "F1: TermDialog help", DefaultTheme.InactiveItem)

	for len(dialogStack.dialogs) > 0 {
		for _, dialog := range dialogStack.dialogs {
			dialog.Open()
		}
		termbox.Flush()

		activeDialog := dialogStack.dialogs[len(dialogStack.dialogs)-1]
		event := termbox.PollEvent()

		handled, shouldClose := activeDialog.HandleEvent(event)
		if !handled {
			for i := len(dialogStack.dialogs) - 1; i >= 0; i-- {
				handled, shouldClose = dialogStack.dialogs[i].HandleGlobalEvent(event)
				if handled {
					break
				}
			}
		}

		if shouldClose {
			dialogStack.Close(activeDialog)
		}
	}
}

func (dialogStack *DialogStack) Stop() {
	// To stop the application, we just need to remove all dialogs.
	// Therefore, on the next iteration of Run, the test "len(dialogStack.dialogs) > 0"
	// will fail and the loop will exit.

	dialogStack.dialogs = nil
}
