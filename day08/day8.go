package main

import (
	"fmt"
	"strings"
)

func main() {
	lines := readFile("./input.txt")

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
	}
}

func executeProgramm(lines []string) (int, bool) {
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
