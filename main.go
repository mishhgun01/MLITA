package main

import (
	"fmt"
	"log"
)

var (
	operations = []byte{'+', '-', '*', '/', ' '}
	letters    = []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p',
		'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	output   = []byte{}
	iterator = 0
	char     byte
	ERR      = 0
)

func main() {
	fmt.Println("Правильная скобочная запись арифметических выражений с двумя видами скобок. \n" +
		"Между двумя круглыми скобками всегда должна стоять квадратная. \n" +
		"Знак умножения может не писаться. Могут быть “лишние” скобки, одна буква может браться в скобки.")
	var ans string = "1"
	for ans == "1" {
		fmt.Scan(&output)
		char = read()
		CorrectEntry()
		if ERR == 0 {
			fmt.Println("Правильная скобочная запись! Хотите попробовать снова? Если да, введите 1")
		} else {
			fmt.Println("Неправильная скобочная запись!")
		}
		iterator = 0
		ERR = 0
		fmt.Println("Хотите попробовать снова? Если да, введите 1")
		fmt.Scan(&ans)
	}
}

func read() byte {
	checkInput()
	if iterator < len(output) {
		char = output[iterator]
		iterator++
		return char
	}
	return 0
}

func CorrectEntry() {
	if contains(char, letters) {
		Letters()
		Action()
	} else if char == '(' {
		RoundedStaplesAction()
		HelpRSAction()
	} else if char == '[' {
		SquaredStaplesAction()
		HelpSSAction()
	} else {
		ERR = 1
	}
}

func RoundedStaplesAction() {
	if char == '(' {
		char = read()
		CorrectEntry()
	} else {
		ERR = 1
	}
}

func SquaredStaplesAction() {
	if char == '[' {
		char = read()
		CorrectEntry()
	} else {
		ERR = 1
	}
}

func HelpRSAction() {
	if contains(char, operations) {
		Operation()
		NextRSAction()
	} else {
		char = read()
	}
}

func NextRSAction() {
	if contains(char, letters) {
		Letters()
		HelpRSAction()
	} else if char == '[' {
		SquaredStaplesAction()
		HelpSSAction()
	} else {
		ERR = 1
	}
}

func HelpSSAction() {
	if contains(char, operations) {
		Operation()
		NextSSAction()
	} else {
		char = read()
	}
}

func NextSSAction() {
	if contains(char, letters) {
		Letters()
		HelpSSAction()
	} else if char == '[' {
		SquaredStaplesAction()
		HelpSSAction()
	} else if char == '(' {
		RoundedStaplesAction()
		HelpRSAction()
	} else {
		ERR = 1
	}
}

func Action() {
	if contains(char, operations) {
		Operation()
		CorrectEntry()
	} else if contains(char, letters) {
		CorrectEntry()
	} else if char == ')' {
		HelpRSAction()
	} else if char == ']' {
		HelpSSAction()
	} else if char == '[' {
		CorrectEntry()
	} else if char == '(' {
		CorrectEntry()
	} else {
		char = read()
	}
}
func Operation() {
	char = read()
}

func Letters() {
	if contains(char, letters) {
		char = read()
	} else {
		ERR = 1
		log.Fatal("Неправильная скобочная запись!")
	}
}

func contains(symbol byte, array []byte) bool {
	steps := 0
	for _, symb := range array {
		if symb == symbol {
			continue
		}
		steps += 1
	}
	return !(steps == len(array))
}

func checkInput() {
	if iterator < len(output)-1 {
		if output[iterator] == ')' && output[iterator+1] == '(' {
			ERR = 1
		}
	}
}

func checkInput_() bool {
	return contains(char, letters) || contains(char, operations)
}
