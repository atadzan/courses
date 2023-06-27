package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	checkString()
}

/*
	Task 1

	На вход подается строка. Нужно определить, является ли она правильной или нет.
	Правильная строка начинается с заглавной буквы и заканчивается точкой. Если строка правильная - вывести Right иначе - вывести Wrong

	Маленькая подсказка: fmt.Scan() считывает строку до первого пробела, вы можете считать полностью строку за раз с помощью bufio:

	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	Sample Input: Быть или не быть.
	Sample Output: Right
*/

func checkString() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	textRunes := []rune(text)
	if strings.HasSuffix(text, ".") && unicode.IsUpper(textRunes[0]) {
		fmt.Println("Right")
	}
}

/*
Task 2
На вход подается строка. Нужно определить, является ли она палиндромом. Если строка является палиндромом - вывести Палиндром иначе - вывести Нет.
Все входные строчки в нижнем регистре.

Палиндром — буквосочетание, слово или текст, одинаково читающееся в обоих направлениях (например, "топот", "око", "заказ").

Sample Input: топот
Sample Output: Палиндром
*/
func checkPolindrom() {
	var word string
	fmt.Scan(&word)

}
