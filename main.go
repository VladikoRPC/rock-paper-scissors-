package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	options := []string{"камень", "ножницы", "бумага"}
	fmt.Println("Сыграйте в камень/ножницы/бумага")

	for {
		choice := userChoice()
		computer := options[r.Intn(len(options))]
		fmt.Println("Компьютер выбрал:", computer)

		if choice == computer {
			fmt.Println("Ничья!")
		} else if choice == "камень" && computer == "ножницы" {
			fmt.Println("Победа!")
		} else if choice == "ножницы" && computer == "бумага" {
			fmt.Println("Победа!")
		} else if choice == "бумага" && computer == "камень" {
			fmt.Println("Победа!")
		} else {
			fmt.Println("Поражение!")
		}

		repeat := isRepeat()
		if !repeat {
			break
		}
	}
	fmt.Println("Спасибо за игру! До встречи!")

}

func userChoice() string {
	var userChoice string
	fmt.Println("Введите свой выбор (камень, ножницы, бумага):")
	fmt.Scan(&userChoice)
	if userChoice != "камень" && userChoice != "ножницы" && userChoice != "бумага" {
		fmt.Println("Некорректный ввод, попробуйте снова.")

	}
	return userChoice
}

func isRepeat() bool {
	var userChoice string
	fmt.Println("Хотите повторить? (да/нет)")
	fmt.Scan(&userChoice)
	return userChoice == "да" || userChoice == "ДА"
}
