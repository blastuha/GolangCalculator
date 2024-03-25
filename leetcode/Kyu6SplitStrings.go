package main

//https://www.codewars.com/kata/515de9ae9dcfc28eb6000001/train/go
//* 'abc' =>  ['ab', 'c_']
//* 'abcdef' => ['ab', 'cd', 'ef']

func Solution(str string) []string {
	var result []string
	for i := 0; i < len(str); i += 2 {
		// Проверяем, не является ли текущая итерация последней парой символов
		if i+1 < len(str) {
			result = append(result, string(str[i:i+2]))
		}
		// Условие, если  i - это последний символ строки, и он является нечетным
		if i+1 == len(str) {
			result = append(result, string(str[i:]+"_"))
			// [i:] === [i:len(str)
		}

	}

	return result
}

func Solution2(str string) []string {
	var result []string
	// Добавляем подчеркивание, если количество символов в строке нечетное
	if len(str)%2 != 0 {
		str += "_"
	}
	// Проходим по строке, создавая пары символов
	for i := 0; i < len(str); i += 2 {
		result = append(result, str[i:i+2])
	}
	return result
}
