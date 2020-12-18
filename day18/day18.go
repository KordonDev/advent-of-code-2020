package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	lines := readFile("./input.txt")

	solution := 0
	for _, line := range lines {
		result, _ := calculateResult(strings.ReplaceAll(line, " ", ""))
		solution = solution + result
	}
	fmt.Println("Solution 1", solution)

	solution = 0
	for _, line := range lines {
		operations := getOperationArray(strings.ReplaceAll(line, " ", ""))
		result := calculateAddThanMultiply(operations)
		solution = solution + stringToInt(result[0])
	}
	fmt.Println("Solution 2", solution)

}

func calculateResult(formula string) (int, string) {
	result := 0
	multiply := false
	add := false
	for index := 0; index < len(formula); index++ {
		char := formula[index]
		// fmt.Println(string(char), formula, index)
		if char == ')' {
			return result, formula[index:]
		}
		if char == '(' {
			subResult, newForumlar := calculateResult(formula[index+1:])
			// fmt.Println("subresult:", subResult, result, multiply, add, newForumlar)
			if multiply {
				result = result * subResult
				multiply = false
			} else if add {
				result = result + subResult
				add = false
			} else {
				result = subResult
			}
			formula = newForumlar
			index = 0

		} else if char == '*' {
			multiply = true
		} else if char == '+' {
			add = true
		} else {
			number := stringToInt(string(char))
			if multiply {
				result = result * number
				multiply = false
			} else if add {
				result = result + number
				add = false
			} else {

				result = number
			}
		}

	}
	// fmt.Println("end", result)
	return result, ""
}

func getOperationArray(formula string) []string {
	var result []string
	for _, char := range formula {
		result = append(result, string(char))
	}
	return result
}

func calculateAddThanMultiply(formula []string) []string {
	allAdded := false
	for len(formula) > 1 {
		changed := false
		for i := 0; i < len(formula); i++ {
			//fmt.Println(formula)
			if formula[i] == ")" {
				return formula
			}
			if formula[i] == "(" {
				closingBracketIndex := findClosingBracket(formula[i:])
				result := calculateAddThanMultiply(formula[i+1 : i+closingBracketIndex])
				formula = slice(formula, i, i+closingBracketIndex+1, result[0])
				changed = true
			}
			if !allAdded && formula[i] == "+" && formula[i+1] != "(" {
				res := calc(formula[i-1], formula[i+1], allAdded)
				formula = slice(formula, i-1, i+2, res)
				changed = true
			}
			if allAdded && formula[i] == "*" && formula[i+1] != "(" {
				res := calc(formula[i-1], formula[i+1], allAdded)
				formula = slice(formula, i-1, i+2, res)
				changed = true
			}
		}
		if !changed {
			allAdded = true
		}
	}
	return formula
}

func calc(s1 string, s2 string, multiply bool) string {
	if multiply {
		res := stringToInt(s1) * stringToInt(s2)
		return strconv.Itoa(res)
	}
	res := stringToInt(s1) + stringToInt(s2)
	return strconv.Itoa(res)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func findClosingBracket(formula []string) int {
	closingBracket := 0
	for index, sign := range formula {
		if sign == "(" {
			closingBracket = closingBracket + 1
		}
		if sign == ")" {
			closingBracket = closingBracket - 1
			if closingBracket == 0 {
				return index
			}
		}
	}
	return len(formula)
}

func slice(array []string, start int, end int, insert string) []string {
	a1 := append(array[:start], insert)
	if end >= len(array) {
		return a1
	}
	return append(a1, array[end:]...)
}
