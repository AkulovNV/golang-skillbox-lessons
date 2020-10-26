package main

import (
	"fmt"
	"math"
)

func main() {
	var deposit uint64
	var mounts, years uint16
	const mountsOfYear = 12
	var rate, currentProfit, totalDeposit, creditProfit, debetProfit float64
	fmt.Println("Расчет дивидендов от вклада")
	fmt.Printf("Введите сумму вклада (валюта): ")
	fmt.Scan(&deposit)
	fmt.Printf("Укажите ежемесячный процент по дивидендам (проценты): ")
	fmt.Scan(&rate)
	fmt.Printf("Установите срок вклада (в годах): ")
	fmt.Scan(&years)
	totalDeposit = float64(deposit)
	for mounts = 1; mounts <= years*mountsOfYear; mounts++ {
		// Вычисляем обший процент на остаток по счету
		currentProfit = totalDeposit * rate / 100
		// Получаем заисляемую на депозит сумму
		creditProfit = math.Trunc(currentProfit*100) / 100
		// Остаток отправляем в банк
		debetProfit += currentProfit - creditProfit
		// Итого по счету на текущий период
		totalDeposit += creditProfit
	}
	fmt.Printf("Итого за %v лет при сумме вклада %v совокупный доход от вклада составит: %.2f\n", years, deposit, totalDeposit)
	fmt.Printf("Сумма, зачисленная в пользу банка: %.10f\n", debetProfit)
}
