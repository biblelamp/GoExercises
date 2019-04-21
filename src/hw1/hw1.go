package main

/*
 * Syntax Go. Homework 1
 * Sergey Iryupin, dated Apr 21, 2019
 */

import (
	"fmt"
	"math"
)

const rate = 64.05

func main() {
	var amountRub, legOne, legTwo, amount, percent float64

	fmt.Println("1. Currency conversion. Enter the amount in rubles:")
	fmt.Scanln(&amountRub)
	fmt.Printf("%.2f RUB = %.2f USD\n", amountRub, convertRUBToUSD(amountRub, rate))

	fmt.Println("2. Calculation of a right triangle. Enter triangle legs:")
	fmt.Scanln(&legOne, &legTwo)
	hypotenuse, area, perimeter := calcTriangle(legOne, legTwo)
	fmt.Printf("Hypotenuse = %.2f Area = %.2f Perimeter = %.2f\n", hypotenuse, area, perimeter)

	fmt.Println("3. Calculation of bank deposit. Enter amount and percent:")
	fmt.Scanln(&amount, &percent)
	amount = calcContribution(amount, percent)
	amount = calcContribution(amount, percent)
	amount = calcContribution(amount, percent)
	amount = calcContribution(amount, percent)
	amount = calcContribution(amount, percent)
	fmt.Printf("Deposit amount after 5 years is %.2f", amount)
}

func convertRUBToUSD(amountRub, rate float64) float64 {
	return amountRub / rate
}

func calcTriangle(legOne, legTwo float64) (hypotenuse, area, perimeter float64) {
	hypotenuse = math.Sqrt(math.Pow(legOne, 2) + math.Pow(legTwo, 2))
	area = legOne * legTwo
	perimeter = legOne + legTwo + hypotenuse
	return
}

func calcContribution(amount, percent float64) float64 {
	return amount + amount/100*percent
}
