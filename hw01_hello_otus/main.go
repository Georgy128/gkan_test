package main

import (
	"fmt"
	//"github.com/golang.org/x/example/stringutil"
)

func main() {
	// Place your code here.

	var s string // Переменная ввода
	// Ввод строки
	fmt.Println("Введите строку:")
	fmt.Scanf("%s\n", &s)
	// Проверяем на ""
	if len(s) > 1 {
		fmt.Println(Reverse(s)) //stringutil.Reverse(s))
	}

}

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
