/*
Написать программу, которая на вход получала бы строку, введенную пользователем, а в файл писала: № строки, дата и сообщение в формате:
2020-02-10 15:00:00 продам гараж
при вводе слова “выход” производился бы выход из программы
Далее программа должна полностью вычитать полученный  с использованием ioutil
*/
package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func main() {
	var b bytes.Buffer
	var userString string
	var strNum int
	fileName := "answers.txt"

	for {
		fmt.Println("Введите строку")
		fmt.Scan(&userString)
		if userString == "выход" {
			break
		}
		strNum += 1
		b.WriteString(fmt.Sprintf("%d: %v %s\n", strNum, time.Now().Format(time.RFC3339), userString))
	}

	if err := ioutil.WriteFile(fileName, b.Bytes(), 0666); err != nil {
		panic(err)
	}

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	readFile, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	fmt.Println("=================")
	fmt.Printf("Содержимое файла\n\n%s", readFile)
}
