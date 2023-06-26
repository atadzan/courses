package main

import "fmt"

func main() {
	printNumSum()
}

/*
	Задание 1

	Дано трехзначное число. Найдите сумму его цифр.

	Формат входных данных
	На вход дается трехзначное число.

	Формат выходных данных
	Выведите одно целое число - сумму цифр введенного числа.

	Sample Input:
	745

	Sample Output:
	16
*/

func printNumSum() {
	var i int
	fmt.Scan(&i)
	fmt.Println((i / 100) + (i / 10 % 10) + (i % 10))
}
