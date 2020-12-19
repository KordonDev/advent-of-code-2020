package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	lines := readFile("./input2.txt")

	rules := make(map[string]string)
	readRules := true
	lineNumber := 0
	for readRules {
		line := lines[lineNumber]
		if len(line) == 0 {
			readRules = false
		} else {
			metaDataSplit := strings.Split(line, ": ")
			rule := "( " + metaDataSplit[1] + " )"
			orIndex := strings.Index(rule, "|")
			if orIndex != -1 {
				rule = "( " + rule[:orIndex] + ") | (" + rule[orIndex+1:] + " )"
			}
			rules[metaDataSplit[0]] = rule
		}
		lineNumber++
	}

	for key, value := range rules {
		keyInValueIndex := strings.Index(value, " "+key+" ")
		if keyInValueIndex != -1 {
			initialTerms := strings.Split(value, "|")
			newEnding := initialTerms[1]
			lengthReplace := len(initialTerms[1]) - 1
			replaceTerm := initialTerms[1][:lengthReplace] + "? "
			for i := 0; i < 10; i++ {
				newEnding = strings.ReplaceAll(newEnding, key, replaceTerm)
			}
			newEnding = strings.ReplaceAll(newEnding, key, "")
			rules[key] = initialTerms[0] + " | " + newEnding
		}
	}

	regex := rules["0"]
	changed := true
	for changed {
		changed = false

		for key, value := range rules {
			parts := strings.Split(value, " ")
			hasOtherRules := false
			for _, part := range parts {
				_, isNumberError := strconv.Atoi(part)
				if isNumberError == nil {
					hasOtherRules = true
				}
			}
			if !hasOtherRules {
				changed = replaceInOtherRules(key, value, &rules) || changed
			}
		}
	}
	regex = "^" + strings.ReplaceAll(strings.ReplaceAll(rules["0"], " ", ""), "\"", "") + "$"
	matcher, _ := regexp.Compile(regex)
	fmt.Println(len(regex))

	solution := 0
	for ; lineNumber < len(lines); lineNumber++ {
		if matcher.MatchString(lines[lineNumber]) {
			solution++
		}
	}

	fmt.Println("Solution:", solution)
}

func replaceInOtherRules(replaceKey string, replaceValue string, rules *(map[string]string)) bool {
	hasChanged := false
	for key, value := range *rules {
		parts := strings.Split(value, " ")
		for i, part := range parts {
			if part == replaceKey {
				parts[i] = "xxx"
				hasChanged = true
			}
		}
		newPart := strings.Join(parts, " ")
		newPart = strings.ReplaceAll(newPart, "xxx", replaceValue)
		(*rules)[key] = newPart
	}
	return hasChanged
}
