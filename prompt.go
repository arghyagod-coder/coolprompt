package coolprompt

import (
	"strconv"

	"github.com/manifoldco/promptui"
)

func Input(label string)(string){
	// label is the question and returns an integer
	prompt := promptui.Prompt{
		Label:    label,
	}

	result, _ := prompt.Run()
	return result
}

func IntInput(label string)(int){
	// label is the question and returns an integer
	prompt := promptui.Prompt{
		Label:    label,
	}

	result, _ := prompt.Run()
	intr, _ := strconv.Atoi(result)
	return intr
}

func Select(label string, options []string)(string){
	// Label is the question
	// Recieves a string array and return the selected item in string type 
	prompt := promptui.Select{
		Label: label,
		Items: options,
	}

	_, result, _ := prompt.Run()

	return result

}

func SelectYesNo(label string, options [2]string)(bool){
	// Label is the question
	// Recieve two options as Argument (for e.g. "Yes" and "No") in a string array and return true or false
	prompt := promptui.Select{
		Label: label,
		Items: options,
	}

	_, result, _ := prompt.Run()

	if result == options[0]{
		return true
	}else{
		return false
	}
}