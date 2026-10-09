package main

import "fmt"

func main() {

	usd_eur, usd_rub := getUserInput()
	calculateResult(usd_eur, usd_rub)

}

func getUserInput() (float64, float64) {

	var usd_eur float64
	var usd_rub float64

	fmt.Print("Введите стоимость одного доллара в евро: ")
	fmt.Scan(&usd_eur)

	fmt.Print("Введите стоимость одного доллара в рублях: ")
	fmt.Scan(&usd_rub)

	return usd_eur, usd_eur
}

func calculateResult(usd_eur, usd_rub float64) {

	eur_rub := usd_rub / usd_eur

	fmt.Println("Константа конвертации EUR-RUB", eur_rub)
}
