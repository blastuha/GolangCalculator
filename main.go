package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// bufio используется для считывания данных из ввода
	// os.Stdin используется для считывания данных с клавиатуры
	scanner := bufio.NewScanner(os.Stdin) // получаем сканер, используя пакет bufio + os.Stdin
	fmt.Println("Введите выражение (например, 5 + 3):")
	scanner.Scan()                // ожидаем ввод с клавиатуры
	inputString := scanner.Text() // записываем введеную строку с клавы в переменную

	valuesArray := strings.Fields(inputString) // пр
	if len(valuesArray) != 3 {
		fmt.Println("Неверный формат ввода.")
		return
	}

	inputValue1, err1 := strconv.Atoi(valuesArray[0])
	inputOperator := valuesArray[1]
	inputValue2, err2 := strconv.Atoi(valuesArray[2])

	if err1 != nil || err2 != nil || inputValue1 < 1 || inputValue1 > 10 || inputValue2 < 1 || inputValue2 > 10 {
		fmt.Println("Ошибка: числа должны быть в диапазоне от 1 до 10.")
		return
	}

	var result int
	switch inputOperator {
	case "+":
		result = inputValue1 + inputValue2
	case "-":
		result = inputValue1 - inputValue2
	case "*":
		result = inputValue1 * inputValue2
	case "/":
		if inputValue2 == 0 {
			fmt.Println("На ноль делить нельзя!")
			return
		} else {
			result = inputValue1 / inputValue2
		}
	default:
		fmt.Println("Неверный оператор.")
		return
	}

	fmt.Printf("Ваш результат: %d\n", result)
}
