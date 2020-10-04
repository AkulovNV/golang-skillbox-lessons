package main

import "fmt"

func main() {

	var fistCoin, secondCoin, thirdCoin, sum int

	fmt.Println("Определить будет ли сдача с указанной суммы при оплате монетами")
	fmt.Printf("Сумма: ")
	fmt.Scan(&sum)
	fmt.Printf("Первая монета: ")
	fmt.Scan(&fistCoin)
	fmt.Printf("Вторая монета: ")
	fmt.Scan(&secondCoin)
	fmt.Printf("Третья монета: ")
	fmt.Scan(&thirdCoin)

	if fistCoin+secondCoin+thirdCoin == sum || fistCoin == sum || secondCoin == sum || thirdCoin == sum || fistCoin+secondCoin == sum || fistCoin+thirdCoin == sum || secondCoin+thirdCoin == sum {
		fmt.Println("Оплата без сдачи")
	} else {
		fmt.Println("Будет получена сдача")
	}
}
