package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Логгер для обработки ошибок
	log.SetPrefix("Калькулятор: ")
	log.SetFlags(0)

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n=== Калькулятор ===")
		showMenu()

		action := readMenuChoice(reader)
		if action == "0" {
			fmt.Println("Выход из программы. До свидания!")
			break
		}

		a := readFloat(reader, "Введите первое число: ")
		b := readFloat(reader, "Введите второе число: ")

		// Выполнение операции
		result, err := calculate(a, b, action)
		if err != nil {
			log.Println("Ошибка выполнения:", err)
			continue
		}

		fmt.Printf("\nРезультат: %.2f %s %.2f = %.2f\n", a, action, b, result)
	}
}

// Показ меню действий
func showMenu() {
	fmt.Println("Выберите действие:")
	fmt.Println("1. Сложение (+)")
	fmt.Println("2. Вычитание (-)")
	fmt.Println("3. Умножение (*)")
	fmt.Println("4. Деление (/)")
	fmt.Println("5. Возведение в степень (^)")
	fmt.Println("6. Остаток от деления (%)")
	fmt.Println("0. Выход")
}

// Чтение выбора действия
func readMenuChoice(reader *bufio.Reader) string {
	choices := map[string]string{
		"1": "+",
		"2": "-",
		"3": "*",
		"4": "/",
		"5": "^",
		"6": "%",
		"0": "0",
	}

	for {
		fmt.Print("Введите номер действия: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if action, exists := choices[input]; exists {
			return action
		}
		log.Println("Ошибка: неверный номер действия. Попробуйте снова.")
	}
}

// Чтение числа с проверкой
func readFloat(reader *bufio.Reader, prompt string) float64 {
	for {
		fmt.Print(prompt)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		value, err := strconv.ParseFloat(input, 64)
		if err == nil {
			return value
		}

		log.Println("Ошибка ввода числа:", err)
	}
}

// Основная функция вычислений
func calculate(a, b float64, action string) (float64, error) {
	operations := map[string]func(float64, float64) (float64, error){
		"+": func(x, y float64) (float64, error) { return x + y, nil },
		"-": func(x, y float64) (float64, error) { return x - y, nil },
		"*": func(x, y float64) (float64, error) { return x * y, nil },
		"/": func(x, y float64) (float64, error) {
			if y == 0 {
				return 0, errors.New("деление на ноль")
			}
			return x / y, nil
		},
		"^": func(x, y float64) (float64, error) { return math.Pow(x, y), nil },
		"%": func(x, y float64) (float64, error) {
			if y == 0 {
				return 0, errors.New("остаток от деления на ноль")
			}
			return float64(int(x) % int(y)), nil
		},
	}

	if operation, exists := operations[action]; exists {
		return operation(a, b)
	}
	return 0, errors.New("неизвестная операция")
}
