package main

import (
	"fmt"
	"strings"
)

type Rule struct {
	min int
	max int
}
type Desc struct {
	name  string
	rules []Rule
}

func main() {
	lines := readFile("./input.txt")

	var descs []Desc
	lineNumber := 0
	allRules := false
	for !allRules {
		line := lines[lineNumber]
		if len(line) == 0 {
			allRules = true
		} else {

			nameAndRules := strings.Split(line, ": ")
			ruleStrings := strings.Split(nameAndRules[1], " or ")
			var rules []Rule
			for _, ruleString := range ruleStrings {
				minMax := strings.Split(ruleString, "-")
				rule := Rule{min: stringToInt(minMax[0]), max: stringToInt(minMax[1])}
				rules = append(rules, rule)
			}
			descs = append(descs, Desc{name: nameAndRules[0], rules: rules})
		}
		lineNumber++
	}

	// skip your ticket
	var validTickets [][]int
	lineNumber = lineNumber + 1
	var ticket []int
	for _, numberString := range strings.Split(lines[lineNumber], ",") {
		ticket = append(ticket, stringToInt(numberString))
	}
	validTickets = append(validTickets, ticket)
	lineNumber = lineNumber + 3

	var invalidNumbers []int
	for i := lineNumber; i < len(lines); i++ {
		validTicket := true
		var ticket []int
		numbers := strings.Split(lines[i], ",")
		for _, number := range numbers {
			valid := false
			numberToTest := stringToInt(number)
			ticket = append(ticket, numberToTest)
			for _, desc := range descs {
				for _, rule := range desc.rules {
					if numberToTest >= rule.min && numberToTest <= rule.max {
						valid = true
					}
				}
			}
			if !valid {
				invalidNumbers = append(invalidNumbers, numberToTest)
				validTicket = false
			}
		}
		if validTicket {
			validTickets = append(validTickets, ticket)
		}
	}

	solution := 0
	for _, value := range invalidNumbers {
		solution = solution + value
	}
	fmt.Println("Solution 1", invalidNumbers, solution)

	rulesWithPosition := make(map[string][]int)

	for position := 0; position < len(validTickets[0]); position++ {
		validRules := make(map[string]bool)
		for _, desc := range descs {
			validRules[desc.name] = true
		}
		for _, ticket := range validTickets {
			for _, desc := range descs {
				validRule := false
				for _, rule := range desc.rules {
					if ticket[position] >= rule.min && ticket[position] <= rule.max {
						validRule = true
					}
				}
				if !validRule {
					validRules[desc.name] = false
				}
			}
		}
		for name, valid := range validRules {
			if valid {
				rulesWithPosition[name] = append(rulesWithPosition[name], position)
			}
		}
	}

	changed := true
	for changed {
		changed = false
		for _, validRules := range rulesWithPosition {
			if len(validRules) == 1 {
				for name, oldPositions := range rulesWithPosition {
					var newPositions []int
					for _, position := range oldPositions {
						if len(oldPositions) == 1 || position != validRules[0] {
							newPositions = append(newPositions, position)
						}
					}
					rulesWithPosition[name] = newPositions
					if len(oldPositions) != len(newPositions) {
						changed = true
					}
				}
			}
		}
	}
	fmt.Println("rulesWithPosition", rulesWithPosition)

	solution = 1
	for name, positions := range rulesWithPosition {
		if strings.HasPrefix(name, "departure") {
			solution = solution * validTickets[0][positions[0]]
		}
	}

	fmt.Println("Solution 2", solution)

}
