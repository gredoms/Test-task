package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	znaki := [4]string{"+", "-", "*", "/"}

	nevernieznacheniya := [2]string{".", ","}

	var exp string

	fmt.Println("Введите")
	fmt.Scanln(&exp)

	oper := -1

	for i := 0; i < len(znaki); i++ {
		soderjit := strings.Contains(exp, znaki[i])
		if soderjit == true {
			oper = i
			break
		}
	}

	if oper == -1 {
		fmt.Print("Ошибка, неккоретные данные")
		return
	}

	massifrazdznazh := strings.Split(exp, znaki[oper])

	if len(massifrazdznazh) > 2 {
		fmt.Println("Ошибка, некорректные данные")
		return
	}

	for i := 0; i < len(nevernieznacheniya); i++ {
		if strings.Contains(massifrazdznazh[0], nevernieznacheniya[i]) || strings.Contains(massifrazdznazh[1], nevernieznacheniya[i]) == true {
			fmt.Println("Ошибка, некорректные данные")
			return
		}
	}

	var a, b int

	if sistemais(massifrazdznazh[0]) != sistemais(massifrazdznazh[1]) == true {
		fmt.Println("Ошибка, некорректные данные")
	} else if sistemais(massifrazdznazh[0]) && sistemais(massifrazdznazh[1]) == true {
		a = isrimskihvarabskie(massifrazdznazh[0])
		b = isrimskihvarabskie(massifrazdznazh[1])
	} else {
		a, _ = strconv.Atoi(massifrazdznazh[0])
		b, _ = strconv.Atoi(massifrazdznazh[1])
	}

	if (10 < a || a < 1) || (10 < b || b < 1) {
		fmt.Println("Ошибка, некорректные данные")
		return
	}

	var result int

	switch znaki[oper] {
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

	if sistemais(massifrazdznazh[0]) && sistemais(massifrazdznazh[1]) == true {
		if result <= 0 {
			panic("Ответ <=0, что невозможно в римской системе счисления")
		}
		fmt.Println(isarabskihvrimskie(result))
	} else {
		fmt.Println(result)
	}
}

func sistemais(r string) bool {
	karta := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
	}
	for i := 0; i < len(r); i++ {
		_, estbli := karta[r[i]]
		if estbli == false {
			return false
		}
	}
	return true
}

func isrimskihvarabskie(x string) int {
	karta := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
	}

	poryadki := []string{"X", "V", "I"}

	massivizx := strings.Split(x, "")

	var result int

	for _, poryadok := range poryadki {
		for _, y := range massivizx {
			if y == poryadok {
				if result <= karta[massivizx[0]]+1 {
					result += karta[poryadok]
				} else {
					result -= karta[poryadok]
				}
			}
		}
	}
	return result
}

func isarabskihvrimskie(x int) string {
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

	poryadoki := []int{100, 90, 50, 40, 10, 9, 5, 4, 1}

	var result string

	for _, poryadok := range poryadoki {
		for x >= poryadok {
			x = x - poryadok
			result = result + karta[poryadok]
		}
	}
	return result
}
