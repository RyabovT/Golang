package main

import (
	"errors"
	"fmt"
)

func main() {
	for {
		var a, b float64
		var action string

		fmt.Println("\n=== Калькулятор ===")
		fmt.Print("Введите первое число: ")
		if _, err := fmt.Scan(&a); err != nil {
			fmt.Println("Ошибка ввода:", err)
			continue
		}

		fmt.Print("Введите действие (+, -, *, /): ")
		if _, err := fmt.Scan(&action); err != nil {
			fmt.Println("Ошибка ввода:", err)
			continue
		}

		fmt.Print("Введите второе число: ")
		if _, err := fmt.Scan(&b); err != nil {
			fmt.Println("Ошибка ввода:", err)
			continue
		}

		result, err := calculate(a, b, action)
		if err != nil {
			fmt.Println("Ошибка:", err)
			continue
		}

		fmt.Printf("\nРезультат: %.2f %s %.2f = %.2f\n", a, action, b, result)

		// Спрашиваем о продолжении
		fmt.Print("\nХотите продолжить? (y/n): ")
		var answer string
		fmt.Scan(&answer)
		if answer != "y" {
			break
		}
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
