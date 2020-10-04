package main

import "fmt"

func main() {

	var fistDep, secondDep, thirdDep int

	fmt.Println("Определение наиболее выгодных депозитных ставок")
	fmt.Printf("Введите первую ставку: ")
	fmt.Scan(&fistDep)
	fmt.Printf("Введите вторую ставку: ")
	fmt.Scan(&secondDep)
	fmt.Printf("Введите третью ставку: ")
	fmt.Scan(&thirdDep)

	if fistDep > secondDep || fistDep > thirdDep {
		if secondDep > thirdDep {
			fmt.Println(fistDep, secondDep)
		} else {
			fmt.Println(fistDep, thirdDep)
		}
	} else if secondDep > fistDep || secondDep > thirdDep {
		if fistDep > thirdDep {
			fmt.Println(secondDep, fistDep)
		} else {
			fmt.Println(secondDep, thirdDep)
		}
	} else if fistDep == secondDep {
		fmt.Println(fistDep, thirdDep)
	} else if fistDep == thirdDep || secondDep == thirdDep {
		fmt.Println(fistDep, secondDep)
	}

}
