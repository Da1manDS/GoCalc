package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func sum(input []int, alpha bool) string {
	result := input[0] + input[1]
	if alpha {
		return numToAlpha(result)
	}
	return strconv.Itoa(result)
}

func sub(input []int, alpha bool) string {
	result := input[0] - input[1]
	if alpha {
		if result < 0 {
			return "Вывод ошибки, так как в римской системе нет отрицательных чисел."
		} else {
			return numToAlpha(result)
		}
	}
	return strconv.Itoa(result)
}

func multi(input []int, alpha bool) string {
	result := input[0] * input[1]
	if alpha {
		return numToAlpha(result)
	}
	return strconv.Itoa(result)
}

func div(input []int, alpha bool) string {
	result := input[0] / input[1]
	if alpha {
		return numToAlpha(result)
	}
	return strconv.Itoa(result)
}

func alphaToDec(input []string) []int {
	decs := []int{}
	for i := 0; i <= 2; i += 2 {
		switch input[i] {
		case "I":
			decs = append(decs, 1)
		case "II":
			decs = append(decs, 2)
		case "III":
			decs = append(decs, 3)
		case "IV":
			decs = append(decs, 4)
		case "V":
			decs = append(decs, 5)
		case "VI":
			decs = append(decs, 6)
		case "VII":
			decs = append(decs, 7)
		case "VIII":
			decs = append(decs, 8)
		case "IX":
			decs = append(decs, 9)
		case "X":
			decs = append(decs, 10)
		}
	}
	return decs
}

func prepare(input string) []string {
	input = strings.TrimSpace(input)
	pre := strings.Split(input, " ")
	return pre
}

func choiceOperator(input []int, operator string, alpha bool) string {
	switch operator {
	case "+":
		return sum(input, alpha)
	case "-":
		return sub(input, alpha)
	case "/":
		return div(input, alpha)
	case "*":
		return multi(input, alpha)
	}
	return ""
}

func checkError(input string) string {
	pre := prepare(input)
	if cap(pre) != 3 {
		err := errors.New("Вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
		fmt.Println(err)
		return ""
	} else if cap(pre) < 3 {
		err := errors.New("Вывод ошибки, так как строка не является математической операцией.")
		fmt.Println(err)
		return ""
	}

	if pre[1] != "+" && pre[1] != "-" && pre[1] != "/" && pre[1] != "*" {
		err := errors.New("Вывод ошибки, неверный оператор.")
		fmt.Println(err)
		return ""
	}

	a, err1 := strconv.Atoi(pre[0])
	b, err2 := strconv.Atoi(pre[2])
	if (err1 == nil && err2 != nil) || (err2 == nil && err1 != nil) {
		err := errors.New("Вывод ошибки, так как используются одновременно разные системы счисления.")
		fmt.Println(err)
		return ""
	}

	result := ""
	decs := []int{}
	if err1 != nil && err2 != nil {
		decs = alphaToDec(pre)
		if pre == nil {
			err := errors.New("Вывод ошибки, не римский знак.")
			fmt.Println(err)
			return ""
		}
		result = choiceOperator(decs, pre[1], true)
	} else {
		decs = append(decs, a)
		decs = append(decs, b)
		result = choiceOperator(decs, pre[1], false)
	}

	return result
}

func calc(input string) string {
	result := checkError(input)
	if result == "" {
		return "Ошибка"
	}
	return result
}

func numToAlpha(x int) string {
	alphas := []string{"I", "V", "X", "L", "C"}
	multi := []int{x % 10, (x / 10) % 10, (x / 100) % 10}
	result := [3]string{}
	for i := 0; i < cap(multi); i++ {
		if multi[i] == 9 {
			result[i] += alphas[i*2] + alphas[i*2+2]
		} else if multi[i] >= 1 {
			if multi[i] >= 5 {
				result[i] += alphas[i*2+1]
			}
			multi[i] = multi[i] % 5
			if multi[i] == 4 {
				result[i] += alphas[i*2] + alphas[i*2+1]
			} else {
				for multi[i] > 0 {
					result[i] += alphas[i*2]
					multi[i]--
				}
			}
		}
	}
	return result[2] + result[1] + result[0]
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите операцию:")
	for {
		text, _ := reader.ReadString('\n')
		fmt.Println("Результат:", calc(text))
	}
}
