package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		var a, b float64
		var action string

		fmt.Println("\n=== Калькулятор ===")

		a = readFloat(reader, "Введите первое число: ")
		action = readAction(reader)
		b = readFloat(reader, "Введите второе число: ")

		result, err := calculate(a, b, action)
		if err != nil {
			fmt.Println("Ошибка:", err)
			continue
		}

		fmt.Printf("\nРезультат: %.2f %s %.2f = %.2f\n", a, action, b, result)

		// Спрашиваем о продолжении
		if !askToContinue(reader) {
			break
		}
	}
}

func readFloat(reader *bufio.Reader, prompt string) float64 {
	for {
		fmt.Print(prompt)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		value, err := strconv.ParseFloat(input, 64)
		if err == nil {
			return value
		}

		fmt.Println("Ошибка ввода:", err)
	}
}

func readAction(reader *bufio.Reader) string {
	for {
		fmt.Print("Введите действие (+, -, *, /): ")
		action, _ := reader.ReadString('\n')
		action = strings.TrimSpace(action)

		if isValidAction(action) {
			return action
		}

		fmt.Println("Ошибка: неизвестное действие. Пожалуйста, попробуйте снова.")
	}
}

func isValidAction(action string) bool {
	return action == "+" || action == "-" || action == "*" || action == "/"
}

func askToContinue(reader *bufio.Reader) bool {
	for {
		fmt.Print("\nХотите продолжить? (y/n): ")
		answer, _ := reader.ReadString('\n')
		answer = strings.TrimSpace(answer)

		if answer == "y" {
			return true
		} else if answer == "n" {
			return false
		}

		fmt.Println("Ошибка: пожалуйста, введите 'y' или 'n'.")
	}
}

func calculate(a, b float64, action string) (float64, error) {
	switch action {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, errors.New("деление на ноль невозможно")
		}
		return a / b, nil
	default:
		return 0, errors.New("неизвестное действие")
	}
}
