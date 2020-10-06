package main

import "fmt"

func main() {
	var i, j, k = 0, 0, 0
	for {
		if i < 10 {
			i += 1
			continue
		}
		if j < 100 {
			j += 1
			continue
		}
		if k < 1000 {
			k += 1
			continue
		} else {
			break
		}
	}

	fmt.Printf("i = %d, j = %d, k = %d\n", i, j, k)

}
