package main

import "fmt"

func main() {

	mainValue := getUserInput()
	total := getTotal()
	newValue := getUserInput()

	result := calculateResult(total, mainValue, newValue)

	fmt.Println("Итоговая сумма:", result)

}

func getUserInput() string {

	var value string
	for {

		fmt.Print("Введите валюту (USD, EUR, RUB): ")
		fmt.Scan(&value)

		if checkValue(value) {
			break
		}
	}

	return value

}

func checkValue(value string) bool {

	return value == "USD" || value == "EUR" || value == "RUB"

}

func getTotal() float64 {

	var total float64
	for {

		fmt.Print("Введите сумму: ")
		fmt.Scan(&total)

		if checktotal(total) {
			break
		}
	}

	return total

}

func checktotal(total float64) bool {

	return total > 0

}

func calculateResult(total float64, mainValue, newValue string) float64 {

	course := setCourse(mainValue, newValue)
	return total / course
}

func setCourse(mainValue, newValue string) (course float64) {

	switch {
	case mainValue == "USD" && newValue == "RUB":
		course = 78.20
	case mainValue == "EUR" && newValue == "RUB":
		course = 98.77
	case mainValue == "RUB" && newValue == "USD":
		course = 1 / 78.20 // ≈ 0.01279
	case mainValue == "RUB" && newValue == "EUR":
		course = 1 / 98.77 // ≈ 0.01012
	case mainValue == "USD" && newValue == "EUR":
		course = 78.20 / 98.77 // ≈ 0.7917
	case mainValue == "EUR" && newValue == "USD":
		course = 98.77 / 78.20 // ≈ 1.2631
	default:
		course = 1 // одинаковые валюты
	}
	return course

}
