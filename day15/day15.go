package main

import (
	"fmt"
	"strings"
)

func main() {
	lines := readFile("./input.txt")

	var saidNumbers []int
	for _, char := range strings.Split(lines[0], ",") {
		saidNumbers = append(saidNumbers, stringToInt(string(char)))
	}

	for i := len(saidNumbers); i < 2020; i++ {
		searchedNumber := saidNumbers[i-1]
		searchList := saidNumbers[0 : len(saidNumbers)-1]

		lastIndex := findLastIndexOfElement(searchedNumber, &searchList)
		saidNumbers = append(saidNumbers, i-1-lastIndex)
	}
	fmt.Println("Solution 1:", saidNumbers[len(saidNumbers)-1])

	numbersWithPosition := make(map[int][]int)
	numbers := strings.Split(lines[0], ",")
	lastSaidNumber := 0
	for i := 0; i < len(numbers); i++ {
		saidNumber := stringToInt(numbers[i])
		numbersWithPosition[saidNumber] = append(numbersWithPosition[saidNumber], i)
		lastSaidNumber = saidNumber
	}

	for i := len(numbers); i < 30000000; i++ {
		numberSaidIndex := numbersWithPosition[lastSaidNumber]
		numberToSay := 0
		if len(numberSaidIndex) > 1 {
			numberToSay = i - 1 - numberSaidIndex[len(numberSaidIndex)-2]
		}
		numbersWithPosition[numberToSay] = append(numbersWithPosition[numberToSay], i)
		lastSaidNumber = numberToSay
	}

	fmt.Println("Solution 2:", lastSaidNumber)

}

func findLastIndexOfElement(element int, list *([]int)) int {
	for j := len(*list) - 1; j >= 0; j-- {
		if (*list)[j] == element {
			return j
		}
	}
	return len(*list)
}
