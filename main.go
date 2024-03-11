package main

import "fmt"

func main() {
	operatorSlice := [4]string{"+", "-", "*", "/"}
	inputValue1 := 0
	inputValue2 := 0
	inputOperator := ""
	isOperatorFound := false
	result := 0

	fmt.Println("Введите первое число: ")
	fmt.Scanln(&inputValue1)
	fmt.Println("Зафиксировано первое число: ", inputValue1)
	fmt.Println("Введите оператор: ")
	fmt.Scanln(&inputOperator)
	fmt.Println("Введите второе число: ")
	fmt.Scanln(&inputValue2)
	fmt.Println("Зафиксировано второе число: ", inputValue2)

	for _, operator := range operatorSlice {
		if operator == inputOperator {
			isOperatorFound = true
			break
		}
	}

	if inputOperator != "" && isOperatorFound == true {
		switch inputOperator {
		case "+":
			result = inputValue1 + inputValue2
		case "-":
			result = inputValue1 - inputValue2
		case "*":
			result = inputValue1 * inputValue2
		case "/":
			if inputValue1 != 0 && inputValue2 != 0 {
				result = inputValue1 / inputValue2
			} else {
				fmt.Println("Одно из значений = 0. На ноль делить нельзя!")
				return
			}
		}

	}

	fmt.Println("Ваш результат: ", result)
}
