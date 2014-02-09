package termdialog

var (
	HelpDialog     *SelectionDialog
	HelpExitDialog *SelectionDialog

	HelpGeneralDialog   *MessageDialog
	HelpMessageDialog   *MessageDialog
	HelpSelectionDialog *MessageDialog
	HelpInputDialog     *MessageDialog
)

func OpenDialogCallback(option *Option) (shouldClose bool) {
	HelpDialog.GetLastDialogStack().Open(option.Data.(Dialog))
	return false
}

func ExitCallback(option *Option) (shouldClose bool) {
	HelpDialog.GetLastDialogStack().Stop()
	return true
}

func init() {
	HelpDialog = NewSelectionDialog("TermDialog Help", nil)
	HelpExitDialog = NewSelectionDialog("Are you sure you want to exit the application?", nil)

	HelpGeneralDialog = NewMessageDialog("General help", "* Any dialog can be closed by pressing the escape key.")
	HelpMessageDialog = NewMessageDialog("Message dialogs", "* Message dialogs display a simple text message.\r\n* Pressing <Enter> or <Space> will close the dialog.")
	HelpSelectionDialog = NewMessageDialog("Selection dialogs", "* Selection dialogs offer a choice of options for the user to select.\r\n* Use the up and down arrow keys to select an option.\r\n* Press the <Enter> or <Space> key to choose the selected option.\r\n* You can also use the <Home> and <End> keys to navigate to the start and end of the list respectively.")
	HelpInputDialog = NewMessageDialog("Input dialogs", "* Input dialogs allow the user to enter a line of text.\r\n* The <Backspace> key can be used as one would expect.\r\n* Pressing <Enter> will return the entered text to the application and close the dialog.")

	HelpDialog.AddOption(&Option{"General", OpenDialogCallback, HelpGeneralDialog})
	HelpDialog.AddOption(&Option{"Message dialogs", OpenDialogCallback, HelpMessageDialog})
	HelpDialog.AddOption(&Option{"Selection dialogs", OpenDialogCallback, HelpSelectionDialog})
	HelpDialog.AddOption(&Option{"Input dialogs", OpenDialogCallback, HelpInputDialog})
	HelpDialog.AddOption(&Option{"Exit the application", OpenDialogCallback, HelpExitDialog})

	HelpExitDialog.AddOption(&Option{"No", nil, nil})
	HelpExitDialog.AddOption(&Option{"Yes", ExitCallback, nil})
}
