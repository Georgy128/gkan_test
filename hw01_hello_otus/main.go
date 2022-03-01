package main

import (
	"fmt"

	"github.com/golang.org/x/example/stringutil"
)

func main() {
	// Place your code here.

	var who string // Переменная ввода
	// Ввод строки
	fmt.Println("Введите строку:")
	fmt.Scanf("%s\n", &who)
	// Проверяем на ""
	if len(who) > 1 {
		fmt.Println("Hello", who)
		fmt.Println(stringutil.Reverse(who))
		//who = strings.Join(os.Args[1:], " ")
	}

}
