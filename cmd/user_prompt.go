package cmd

import "github.com/manifoldco/promptui"

func PromptUser(message string) error {
	prompt := promptui.Prompt{
		Label:     message,
		IsConfirm: true,
	}
	_, err := prompt.Run()
	return err
}
