package main

import (
	"fmt"
	"math"
)

func main() {
	lines := readFile("./input.txt")

	preamble := 25
	var numbers []int
	for _, line := range lines {
		numbers = append(numbers, stringToInt(line))
	}

	for i := preamble; i < len(numbers); i++ {
		firstAdditionIndex := i - preamble
		lastAdditionIndex := i
		if !canBeAdded(numbers[i], numbers[firstAdditionIndex:lastAdditionIndex]) {
			fmt.Println(numbers[i])
			sumUpContigousList(numbers[i], numbers)
		}
	}
	/*
		resultPart1, _ := executeProgramm(lines)
		fmt.Println(resultPart1)

		for index := range lines {
			oldLine := lines[index]
			parts := strings.Split(lines[index], " ")
			if parts[0] == "nop" {
				lines[index] = "jmp " + parts[1]
			}
			if parts[0] == "jmp" {
				lines[index] = "nop " + parts[1]
			}
			acc, finished := executeProgramm(lines)
			if finished {
				fmt.Println("Part2", acc)
			}
			lines[index] = oldLine
		}*/
}

func canBeAdded(sum int, terms []int) bool {
	for _, term1 := range terms {
		for _, term2 := range terms {
			if term1 != term2 && term1+term2 == sum {
				return true
			}
		}
	}
	return false
}

func sumUpContigousList(sum int, terms []int) {
	for index, term1 := range terms {
		currentSum := term1
		for i := index + 1; currentSum < sum; i++ {
			currentSum = currentSum + terms[i]
			if currentSum == sum {
				addSmallestAndLargestTerms(terms[index:i])
			}

		}
	}
}

func addSmallestAndLargestTerms(terms []int) {
	max := 0
	min := math.MaxInt32
	for _, term := range terms {
		if term < min {
			min = term
		}
		if term > max {
			max = term
		}
	}
	fmt.Println(min, "+", max)
	fmt.Println("Addittion", min+max)
}

/*func executeProgramm(lines []string) (int, bool) {
	visitedLines := make(map[int]string)
	accumulator := 0
	executedLine := 0
	_, found := visitedLines[executedLine]
	for !found && executedLine != len(lines) {
		visitedLines[executedLine] = "1"
		parts := strings.Split(lines[executedLine], " ")
		if parts[0] == "nop" {
			executedLine = executedLine + 1
		}
		if parts[0] == "jmp" {
			lineJumps := stringToInt(string(parts[1]))
			executedLine = executedLine + lineJumps
		}
		if parts[0] == "acc" {
			addValue := stringToInt(string(parts[1]))
			accumulator = accumulator + addValue
			executedLine = executedLine + 1
		}
		_, found = visitedLines[executedLine]
	}
	return accumulator, executedLine == len(lines)
}
*/
