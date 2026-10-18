package main

import (
	"errors"
	"fmt"
)

func main() {

	defer func() {

		r := recover()
		if r != nil {
			fmt.Println("Ошибка", r)
		}

	}()

	Value := getUserInput()
	course, err := checkValue(Value)

	if err != nil {
		panic("Неверно указана валюта")
	}

	total := getTotal()

	result := calculateResult(total, course)

	fmt.Println("Сумма в выбранной валюте:", result)

}

func getUserInput() string {

	var value string

	fmt.Print("Введите валюту (USD, EUR, RUB): ")
	fmt.Scan(&value)

	return value
}

func getTotal() float64 {

	var total float64

	fmt.Print("Введите сумму в рублях: ")
	fmt.Scan(&total)

	return total

}

func checkValue(value string) (float64, error) {

	var course float64

	switch {

	case value == "USD":

		course = 68.90

	case value == "EUR":

		course = 90.70

	case value == "RUB":

		course = 1

	default:

		course = 0

	}

	if course == 0 {

		return 0, errors.New("NO_PARAMS_ERROR")

	} else {

		return course, nil
	}
}

func calculateResult(total, course float64) float64 {

	result := total / course

	return result
}
