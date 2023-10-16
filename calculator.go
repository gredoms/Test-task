package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	signs := [4]string{"+", "-", "*", "/"}

	incorrectSigns := [2]string{".", ","}

	var expression string

	fmt.Println("Введите")
	fmt.Scanln(&expression)

	q := -1

	for i := 0; i < len(signs); i++ {
		contains := strings.Contains(expression, signs[i])
		if contains == true {
			q = i
			break

		}
	}

	if q == -1 {
		fmt.Print("Ошибка, неккоретные данные")
		return
	}

	splitArray := strings.Split(expression, signs[q])

	if len(splitArray) > 2 {
		fmt.Println("Ошибка, некорректные данные")
		return
	}

	for i := 0; i < len(incorrectSigns); i++ {
		if strings.Contains(splitArray[0], incorrectSigns[i]) || strings.Contains(splitArray[1], incorrectSigns[i]) == true {
			fmt.Println("Ошибка, некорректные данные")
			return
		}
	}

	var a, b int
	variable1 := systemDefinition(splitArray[0])
	variable2 := systemDefinition(splitArray[1])
	if variable1 != variable2 == true {
		fmt.Println("Ошибка, некорректные данные")
	} else if variable1 && variable2 == true {
		a = fromRomantoarabic(splitArray[0])
		b = fromRomantoarabic(splitArray[1])
	} else {
		a, _ = strconv.Atoi(splitArray[0])
		b, _ = strconv.Atoi(splitArray[1])
	}

	if (10 < a || a < 1) || (10 < b || b < 1) {
		fmt.Println("Ошибка, некорректные данные")
		return
	}

	var result int

	switch signs[q] {
	case "+":
		result = a + b
		break
	case "-":
		result = a - b
		break
	case "*":
		result = a * b
		break
	case "/":
		result = a / b
		break
	}

	if systemDefinition(splitArray[0]) && systemDefinition(splitArray[1]) == true {
		if result <= 0 {
			panic("Ответ <=0, что невозможно в римской системе счисления")
		}
		fmt.Println(fromArabictoroman(result))
	} else {
		fmt.Println(result)
	}
}

func systemDefinition(r string) bool {
	k := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
	}
	for i := 0; i < len(r); i++ {
		_, contains := k[r[i]]
		if contains == false {
			return false
		}
	}
	return true
}

func fromRomantoarabic(x string) int {
	k := map[int32]int{
		'I': 1,
		'V': 5,
		'X': 10,
	}

	array := []int32{'X', 'V', 'I'}

	var result int

	for _, arrayValue := range array {
		for _, y := range x {
			if y == arrayValue {
				if result <= k[int32(x[0])]+1 {
					result += k[arrayValue]
				} else {
					result -= k[arrayValue]
				}
			}
		}
	}
	return result
}

func fromArabictoroman(x int) string {
	karta := map[int]string{
		1:   "I",
		4:   "IV",
		5:   "V",
		9:   "IX",
		10:  "X",
		40:  "XL",
		50:  "L",
		90:  "XC",
		100: "C",
	}

	array := []int{100, 90, 50, 40, 10, 9, 5, 4, 1}

	var result string

	for _, arrayValue := range array {
		for x >= arrayValue {
			x = x - arrayValue
			result = result + karta[arrayValue]
		}
	}
	return result
}
