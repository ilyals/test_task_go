package calculator

import (
	"errors"
	"strconv"
	"strings"
)

// Карта для преобразования римских цифр в арабские
var romanToArabic = map[string]int{
	"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
	"VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
}

// Карта для преобразования арабских цифр в римские
var arabicToRoman = []struct {
	Value  int
	Symbol string
}{
	{1000, "M"}, {900, "CM"}, {500, "D"}, {400, "CD"},
	{100, "C"}, {90, "XC"}, {50, "L"}, {40, "XL"},
	{10, "X"}, {9, "IX"}, {5, "V"}, {4, "IV"}, {1, "I"},
}

// Преобразование римских чисел в арабские
func romanToInt(s string) (int, error) {
	if val, exists := romanToArabic[s]; exists {
		return val, nil
	}
	panic("Некорректное римское число")
}

// Преобразование арабских чисел в римские
func intToRoman(num int) (string, error) {
	if num <= 0 {
		panic("Римские числа не могут быть меньше 1")
	}

	var result strings.Builder
	for _, entry := range arabicToRoman {
		for num >= entry.Value {
			result.WriteString(entry.Symbol)
			num -= entry.Value
		}
	}
	return result.String(), nil
}

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
	case "%":
		if b == 0 {
			return 0, errors.New("на ноль делить нельзя")
		}
		return a % b, nil
	default:
		return 0, errors.New("неизвестная операция")
	}
}

func isRoman(s string) bool {
	_, exists := romanToArabic[s]
	return exists
}

func CalculateExpression(input string) (string, error) {

	// Удаляем лишние пробелы
	input = strings.TrimSpace(input)
	// Переводим все в верхний регистр
	input = strings.ToUpper(input)

	// Проверяем на наличие нескольких операторов или пустую строку
	if input == "" {
		panic("Строка не является математической операцией")
	}

	// Подсчет количества операторов
	operatorCount := strings.Count(input, "+") + strings.Count(input, "-") +
		strings.Count(input, "*") + strings.Count(input, "/") + strings.Count(input, "%")

	// Проверка на наличие слишком большого числа операторов
	if operatorCount == 0 {
		panic("Строка не является математической операцией")
	}
	if operatorCount > 1 {
		panic("Формат математической операции не удовлетворяет заданию — два операнда и один оператор")
	}

	// Возможные операторы
	operators := []string{"+", "-", "*", "/", "%"}

	// Найдем оператор в строке
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
		panic("Некорректный формат выражения. Ожидается формат: число оператор число")
	}

	// Разделяем строку на левую и правую части вокруг оператора
	left := strings.TrimSpace(input[:index])
	right := strings.TrimSpace(input[index+1:])

	// Проверяем, являются ли оба числа римскими или арабскими
	isLeftRoman := isRoman(left)
	isRightRoman := isRoman(right)

	if isLeftRoman != isRightRoman {
		panic("Ошибка: нельзя смешивать арабские и римские цифры в одном выражении!")
		//return "", errors.New("нельзя смешивать римские и арабские числа")
	}

	// Выполняем вычисление
	var result int
	if isLeftRoman {
		// Если оба числа римские
		num1, err := romanToInt(left)
		if err != nil {
			return "", err
		}

		num2, err := romanToInt(right)
		if err != nil {
			return "", err
		}

		// Проверка диапазона для римских чисел
		if num1 < 1 || num1 > 10 || num2 < 1 || num2 > 10 {
			panic("Римские числа должны быть в диапазоне от 1 до 10")
		}

		result, err = calculate(num1, num2, operator)
		if err != nil {
			return "", err
		}

		// Преобразуем результат обратно в римское число
		if result < 1 {
			panic("Римские числа не поддерживают значения меньше 1")
		}

		romanResult, err := intToRoman(result)
		if err != nil {
			return "", err
		}
		return romanResult, nil
	} else {
		// Если оба числа арабские
		num1, err := strconv.Atoi(left)
		if err != nil {
			panic("Некорректное первое арабское число")
		}

		num2, err := strconv.Atoi(right)
		if err != nil {
			panic("Некорректное второе арабское число")
		}

		// Проверка диапазона для римских чисел
		if num1 < 1 || num1 > 10 || num2 < 1 || num2 > 10 {
			panic("Арабские числа должны быть в диапазоне от 1 до 10")
		}

		result, err = calculate(num1, num2, operator)
		if err != nil {
			return "", err
		}

		return strconv.Itoa(result), nil
	}
}
