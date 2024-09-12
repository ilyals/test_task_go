package main

import (
	"bufio"
	"calculator/calculator" // Импортируем пакет калькулятора
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите выражение (например: 3 + 5 или IV * II):")

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	// Разбираем и вычисляем результат с помощью пакета калькулятора
	result, err := calculator.CalculateExpression(input)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	fmt.Println("Результат:", result)
}
