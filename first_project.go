package main

import (
	"bufio"   // это для чтения строки
	"errors"  // обработка ошибок
	"fmt"     // реализует форматированный ввод/вывод
	"os"      // платформонезависимый интерфейс к операционной системе
	"strconv" // преобразования в строковые представления и из них основных типов данных
	"strings" // реализует простые функции для работы со строками в кодировке UTF-8
)

// Карта для преобразования римских цифр в арабские
var romanToArabic = map[string]int{
	"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
	"VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
}

// Функция для выполнения арифметических операций
func calculate(a int, b int, operator string) (int, error) {
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
	case "%":
		if b == 0 {
			return 0, errors.New("на ноль делить нельзя")
		}
		return a % b, nil
	default:
		return 0, errors.New("неизвестная операция")
	}
}

// Функция для проверки, является ли строка римским числом
func isRoman(s string) bool {
	_, exists := romanToArabic[s] // Первое значение — это значение по ключу, но оно здесь не используется (поэтому вместо переменной _).
	// Второе значение, exists, — это булево значение, указывающее, существует ли ключ s в карте.
	return exists
}

// Преобразование римских чисел в арабские
func romanToInt(s string) (int, error) {
	if val, exists := romanToArabic[s]; exists {
		return val, nil
	}
	return 0, errors.New("некорректное римское число")
}

// Преобразование арабских чисел в римские (динамическая функция)
func intToRoman(num int) (string, error) {
	if num <= 0 {
		return "", errors.New("римские числа не могут быть меньше 1")
	}

	// Последовательно вычитаем из числа самое большое возможное значение, добавляя при этом соответствующую римскую цифру к результату, пока число не станет равным нулю (значения засунем в массивы)

	vals := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}                        // содержит арабские числа в убывающем порядке, от 1000 до 1
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"} // содержит соответствующие римские цифры для каждого значения в массиве vals

	var result strings.Builder       // более производительный способ создания строк по сравнению с обычной конкатенацией
	for i := 0; i < len(vals); i++ { // цикл - проходит по всем элементам массива vals и соответствующим им римским символам из symbols
		for num >= vals[i] { // вложенный цикл - проверяет, можно ли вычесть текущее значение vals[i] из числа num. Если да, то:
			num -= vals[i]                 // Вычитает это значение из num
			result.WriteString(symbols[i]) // Добавляет соответствующий римский символ к результату
		}
	}
	return result.String(), nil // после того как все значения из массива vals были вычтены из числа, и соответствующие символы были добавлены к строке, возвращаем римское число в виде строки
}

// Функция для извлечения чисел и оператора из строки
func parseExpression(input string) (string, string, string, error) {
	// Удаляем пробелы вокруг строки
	input = strings.TrimSpace(input)

	// Возможные операторы
	operators := []string{"+", "-", "*", "/", "%"}

	// Найдем позицию первого оператора
	var operator string
	var index int
	for _, op := range operators {
		if i := strings.Index(input, op); i != -1 {
			operator = op
			index = i
			break
		}
	}

	// Если оператор не найден, возвращаем ошибку
	if operator == "" {
		return "", "", "", errors.New("некорректный формат выражения. Ожидается формат: число оператор число")
	}

	// Разбиваем строку на левую и правую части вокруг оператора
	left := strings.TrimSpace(input[:index])
	right := strings.TrimSpace(input[index+1:])

	// Если одна из частей пустая, возвращаем ошибку
	if left == "" || right == "" {
		return "", "", "", errors.New("некорректный формат выражения. Ожидается формат: число оператор число")
	}

	// Возвращаем левую часть, оператор и правую часть
	return left, operator, right, nil
}

func main() {
	// Используем bufio для чтения строки с пробелами
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите выражение (например: 3 + 5 или IV * II или 10 % 3):")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	// Преобразуем ввод в верхний регистр
	input = strings.ToUpper(input)

	// Разбираем выражение
	num1, operator, num2, err := parseExpression(input)
	if err != nil {
		panic(err)
	}

	// Определяем, римские или арабские числа переданы
	isRomanNumeral := isRoman(num1)

	// Проверяем корректность чисел
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
		// Преобразуем арабские числа с использованием strconv.Atoi
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
			panic("Результат римских чисел не может быть меньше 1. Римские числа не поддерживают ноль или отрицательные значения.")
		}
		romanResult, err := intToRoman(result)
		if err != nil {
			panic(err)
		}
		fmt.Println("Результат:", romanResult)
	} else {
		// Арабские числа: используем strconv.Itoa для вывода
		fmt.Println("Результат:", strconv.Itoa(result))
	}
}
