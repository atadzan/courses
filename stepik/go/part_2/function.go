package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	//var s = "hello"
	var g = "приве"
	fmt.Println(utf8.RuneCountInString(g))
	//minimumNumber()
}

/*
	Task #1

	Напишите функцию, находящую наименьшее из четырех введённых в этой же функции чисел.

	Входные данные

	Вводится четыре числа.

	Выходные данные

	Необходимо вернуть из функции наименьшее из 4-х данных чисел

	Sample Input: 4 5 6 7
	Sample Output: 4
*/

func minimumNumber() int {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)
	switch {
	case a < b:

	}
	return 0
}

/*
	Task #2

	Последовательность Фибоначчи определена следующим образом: φ1=1, φ2=1, φn=φn-1+φn-2 при n>1.
	Начало ряда Фибоначчи выглядит следующим образом: 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, ... Напишите функцию, которая по указанному натуральному n возвращает φn.

	Входные данные

	Вводится одно число n.

	Выходные данные

	Необходимо вывести  значение φn.

	Sample Input: 4
	Sample Output: 3
*/

func fibonacci(n int) int {
	//print your code here
	return 0
}
