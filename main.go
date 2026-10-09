package main

import "fmt"

func main() {

	var USD string = "USD"
	var RUB string = "RUB"

	Value := getUserInput()
	calculateResult(Value, USD, RUB)

}

func getUserInput() float64 {

	var value float64

	fmt.Print("Введите число: ")
	fmt.Scan(&value)

	return value
}

func calculateResult(Value float64, USD, RUB string) {

}
