package main

import "fmt"

func main() {

	fmt.Println("Вычисление налоговых выплат сотрудника")
	fmt.Println("--------------------------------------")

	var income, totalTaxFee int
	var minTaxPercent = 13
	var midTaxPercent = 20
	var maxTaxPercent = 30

	var minCutIncome = 10000
	var maxCutIncome = 50000

	fmt.Printf("Укажите ваш доход: ")
	fmt.Scan(&income)

	if income > maxCutIncome {
		totalTaxFee = (income-maxCutIncome)*maxTaxPercent/100 + (maxCutIncome-minCutIncome)*midTaxPercent/100 + minCutIncome*minTaxPercent/100
	} else if income <= maxCutIncome && income > minCutIncome {
		totalTaxFee = (income-minCutIncome)*midTaxPercent/100 + minCutIncome*minTaxPercent/100
	} else {
		totalTaxFee = income * minTaxPercent / 100
	}

	if totalTaxFee == 0 {
		fmt.Println("Вы ничего не заработали, потому что занимались ерундой")
	} else {
		fmt.Printf("Итоговый налоговый вычет составит - %d \n", totalTaxFee)
	}
}
