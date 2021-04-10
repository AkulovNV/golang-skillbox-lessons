/*
Создать файл только для чтения и проверить, что в него нельзя записать данные
*/
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	fileName := "readonly.txt"
	var userString string
	file, err := os.Create(fileName)
	if err := os.Chmod(fileName, 0444); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Введите что-либо")
	fmt.Scan(&userString)

	w := bufio.NewWriter(file)
	if err != nil {
		fmt.Println("Файл только на чтение")
		log.Fatal(err)
	}
	defer file.Close()
	w.WriteString(userString)
}
