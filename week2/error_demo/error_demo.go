package main

import (
	"errors"
	"fmt"
)

func divide(a int, b int) int {
	return a / b
}

func main() {

	var input int
	_, err := fmt.Scan(&input) // функция Scan возвращает два параметра, но нам сейчас важно проверить только ошибку
	if err != nil {
		fmt.Println("Проверьте типы входных параметров")
	} else {
		fmt.Println(divide(input, 5)) //Выведем результат, если ошибок нет
	}

	var usrErr = errors.New("my error")
	fmt.Println("", usrErr)
}
