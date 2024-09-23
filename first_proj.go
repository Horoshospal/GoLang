go
package main

import (
	"fmt"
	"strings"
	"strconv"
	"os"
)

// Функция проверки валидности арабских чисел
func isArabic(num string) bool {
	_, err := strconv.Atoi(num)
	return err == nil
}

// Функция проверки валидности римских чисел
func isRoman(num string) bool {
	romanNumerals := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	for _, r := range romanNumerals {
		if r == num {
			return true
		}
	}
	return false
}

// Функция для преобразования римского числа в арабское
func romanToArabic(roman string) int {
	romanMap := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
	}

	sum := 0
	for i := 0; i < len(roman); i++ {
		value := romanMap[rune(roman[i])]
		if i < len(roman)-1 && value < romanMap[rune(roman[i+1])] {
			sum -= value
		} else {
			sum += value
		}
	}
	return sum
}

// Функция для преобразования арабского числа в римское
func arabicToRoman(num int) string {
	if num < 1 {
		panic("Результат вычислений должен быть положительным.")
	}

	romanNumerals := []struct {
		Value int
		Symbol string
	}{
		{1000, "M"}, {900, "CM"}, {500, "D"}, {400, "CD"},
		{100, "C"}, {90, "XC"}, {50, "L"}, {40, "XL"},
		{10, "X"}, {9, "IX"}, {5, "V"}, {4, "IV"},
		{1, "I"},
	}

	result := ""
	for _, rn := range romanNumerals {
		for num >= rn.Value {
			result += rn.Symbol
			num -= rn.Value
		}
	}
	return result
}

// Функция для выполнения операций
func calculate(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		if b == 0 {
			panic("Деление на ноль.")
		}
		return a / b
	default:
		panic("Недопустимая операция.")
	}
}

func main() {
	var input string

	fmt.Println("Введите выражение (например, II + IV или 3 * 2):")
	fmt.Scanln(&input)

	// Разделяем входные данные
	parts := strings.Fields(input)
	if len(parts) != 3 {
		panic("Ввод должен содержать два числа и оператор.")
	}

	// Получаем числа и оператор
	xStr, op, yStr := parts[0], parts[1], parts[2]

	var x, y int
	var isRomanInput bool

	// Проверяем является ли ввод римскими или арабскими числами
	if isArabic(xStr) && isArabic(yStr) {
		x, _ = strconv.Atoi(xStr)
		y, _ = strconv.Atoi(yStr)
		if x < 1 || x > 10 || y < 1 || y > 10 {
			panic("Числа должны быть от 1 до 10 включительно.")
		}
	} else if isRoman(xStr) && isRoman(yStr) {
		x = romanToArabic(xStr)
		y = romanToArabic(yStr)
		isRomanInput = true
		if x < 1 || x > 10 || y < 1 || y > 10 {
			panic("Числа должны быть от I до X включительно.")
		}
	} else {
		panic("Все числа должны быть либо арабскими, либо римскими.")
	}

	// Выполняем операцию
	result := calculate(x, y, op)

	// Выводим результат
	if isRomanInput {
		if result < 1 {
			panic("Результат римских чисел должен быть положительным.")
		}
		fmt.Println("Результат:", arabicToRoman(result))
	} else {
		fmt.Println("Результат:", result)
	}
}
