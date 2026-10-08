package main

import "fmt"

func main() {

	const usd_eur float64 = 1.36
	const usd_rub float64 = 87

	const eur_rub = usd_rub / usd_eur

	fmt.Println("Константа конвертации EUR-RUB", eur_rub)

}
