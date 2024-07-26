package main

import (
	"fmt"
)

func spaceGroupReverse(input string) string {
	size := len(input)
	temp := make([]rune, size)
	concat := make([]rune, size)

	pos := size
	count := 0
	for i := 0; i < len(input); i++ {
		if input[i] == ' ' {
			copy(concat[i-count:i], temp[pos:size])
			pos = size - 1
			count = 0
			concat[i] = ' '
			continue
		}
		pos--
		temp[pos] = rune(input[i])
		count++
	}

	if count > 0 {
		copy(concat[size-count:size], temp[pos:size])
	}

	return string(concat)
}

func buzzFizz(n uint8) {
	for i := uint8(1); i < n; i++ {
		if i%3 == 0 {
			fmt.Print("Fizz")
		}

		if i%5 == 0 {
			fmt.Print("Buzz")
		}

		fmt.Println(i)
	}
}

func countSerial(n uint) {
	prev, next := uint64(0), uint64(1)
	fmt.Print("(")
	for i := uint(0); i < n; i++ {
		if i > 0 {
			fmt.Print(",")
		}
		fmt.Print(prev)
		prev, next = next, prev+next
	}
	fmt.Print(")")
}

func getLowest(data []uint) uint {

	size := len(data)
	lowest := data[0]
	for i := 1; i < size; i++ {
		if lowest > data[i] {
			lowest = data[i]
		}
	}
	return lowest
}

const FIRST_DIGIT rune = '0'
const LAST_DIGIT = '9'

func countDigits(chars []rune) int {
	count := 0
	for i := 0; i < len(chars); i++ {
		if chars[i] <= LAST_DIGIT && chars[i] >= FIRST_DIGIT {
			count++
		}
	}
	return count
}

func task1() {
	mirroredTexts := []string{
		"italem irad irigayaj",
		"iadab itsap ulalreb",
		"nalub kusutret gnalali",
	}
	fmt.Println("Task #1")
	for i, mirroredText := range mirroredTexts {
		fmt.Printf("%d. input  : %s\n   output: %s\n", i+1, mirroredText, spaceGroupReverse(mirroredText))
	}
}

func task2() {
	fmt.Println("Task #2")
	buzzFizz(100)
}

func task3() {
	fmt.Println("Task #3")
	countSerial(9)
	fmt.Println()
}

func task4() {
	fmt.Println("Task #4")
	stockPrices := [][]uint{{7, 8, 3, 10, 8},
		{5, 12, 11, 12, 10},
		{7, 18, 27, 10, 29},
		{20, 17, 15, 14, 10}}
	for i, stockPrice := range stockPrices {
		fmt.Printf("%d. input : %v\n   output: %d\n", i+1, stockPrice, getLowest(stockPrice))
	}
}

func task5() {
	fmt.Println("Task #5")
	mixChars := []string{
		"b7h6hki5g78",
		"7b8569nfy69",
		"uhbn7651g79",
	}
	for i, mixChar := range mixChars {
		fmt.Printf("%d. input : %v\n   output: %d\n", i+1, string(mixChar), countDigits([]rune(mixChar)))
	}
}

func main() {
	task1()
	fmt.Println()
	task2()
	fmt.Println()
	task3()
	fmt.Println()
	task4()
	fmt.Println()
	task5()

}
