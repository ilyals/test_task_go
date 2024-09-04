package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// Карта для преобразования римских цифр в арабские
var romanToArabic = map[string]int{
	"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
	"VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
}

// Карта для преобразования арабских цифр в римские
var arabicToRoman = []string{
	"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X",
	"XI", "XII", "XIII", "XIV", "XV", "XVI", "XVII", "XVIII", "XIX", "XX",
}

// Функция для выполнения арифметических операций
func calculate(a, b int, operator string) (int, error) {
	switch operator {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, errors.New("на ноль делить нельзя")
		}
		return a / b, nil
	default:
		return 0, errors.New("неизвестная операция")
	}
}

// Функция для проверки, является ли строка римским числом
func isRoman(s string) bool {
	_, exists := romanToArabic[s]
	return exists
}

// Преобразование римских чисел в арабские
func romanToInt(s string) (int, error) {
	if val, exists := romanToArabic[s]; exists {
		return val, nil
	}
	return 0, errors.New("некорректное римское число")
}

// Преобразование арабских чисел в римские
func intToRoman(num int) (string, error) {
	if num <= 0 || num >= len(arabicToRoman) {
		return "", errors.New("римские числа не могут быть меньше 1")
	}
	return arabicToRoman[num], nil
}

// Функция для извлечения чисел и оператора из строки
func parseExpression(input string) (string, string, string, error) {
	var num1, num2, operator string
	// Временные строки для сборки числа
	var temp strings.Builder

	for _, ch := range input {
		if unicode.IsDigit(ch) || unicode.IsLetter(ch) {
			// Добавляем цифру или букву к текущему числу
			temp.WriteRune(ch)
		} else if ch == '+' || ch == '-' || ch == '*' || ch == '/' {
			if operator != "" {
				return "", "", "", errors.New("некорректное выражение")
			}
			// Сохраняем первое число и оператор
			num1 = temp.String()
			operator = string(ch)
			temp.Reset()
		} else {
			return "", "", "", errors.New("некорректный символ в выражении")
		}
	}

	// Сохраняем второе число
	num2 = temp.String()

	// Проверка на корректность
	if num1 == "" || num2 == "" || operator == "" {
		return "", "", "", errors.New("некорректный формат выражения. Ожидается формат: число оператор число")
	}

	return num1, operator, num2, nil
}

func main() {
	var input string
	fmt.Println("Введите выражение (например: 3 + 5 или IV * II):")
	fmt.Scanln(&input)

	// Убираем пробелы
	input = strings.TrimSpace(input)

	// Разбираем выражение
	num1, operator, num2, err := parseExpression(input)
	if err != nil {
		panic(err)
	}

	// Определяем, римские или арабские числа переданы
	isRomanNumeral := isRoman(num1) && isRoman(num2)
	isArabicNumeral := !isRoman(num1) && !isRoman(num2)

	if !isRomanNumeral && !isArabicNumeral {
		panic("Нельзя смешивать арабские и римские числа")
	}

	var a, b int

	if isRomanNumeral {
		// Преобразуем римские числа в арабские
		a, err = romanToInt(num1)
		if err != nil {
			panic(err)
		}
		b, err = romanToInt(num2)
		if err != nil {
			panic(err)
		}
	} else {
		// Преобразуем арабские числа из строк
		a, err = strconv.Atoi(num1)
		if err != nil {
			panic("Некорректное арабское число")
		}
		b, err = strconv.Atoi(num2)
		if err != nil {
			panic("Некорректное арабское число")
		}
	}

	// Проверка диапазона чисел
	if a < 1 || a > 10 || b < 1 || b > 10 {
		panic("Числа должны быть в диапазоне от 1 до 10")
	}

	// Выполняем вычисление
	result, err := calculate(a, b, operator)
	if err != nil {
		panic(err)
	}

	// Вывод результата
	if isRomanNumeral {
		// Римские числа: результат должен быть положительным
		if result < 1 {
			panic("Результат римских чисел не может быть меньше 1")
		}
		romanResult, err := intToRoman(result)
		if err != nil {
			panic(err)
		}
		fmt.Println("Результат:", romanResult)
	} else {
		// Арабские числа: выводим результат напрямую, включая отрицательные значения
		fmt.Println("Результат:", result)
	}
}
