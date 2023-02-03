package cmd

import "github.com/manifoldco/promptui"

// PromptUser prompts the user to confirm an action and returns an error if the user does not confirm
func PromptUser(message string) error {
	prompt := promptui.Prompt{
		Label:     message,
		IsConfirm: true,
	}
	_, err := prompt.Run()
	return err
}
