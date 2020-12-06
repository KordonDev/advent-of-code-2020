package main

import (
	"fmt"
)

func main() {
	texts := readFile("./input.txt")

	answeredQuestion := make(map[rune]string)
	answerCount := 0
	var allAnswerCounts []int

	for _, text := range texts {
		if len(string(text)) == 0 {
			allAnswerCounts = append(allAnswerCounts, answerCount)
			answeredQuestion = make(map[rune]string)
			answerCount = 0
		}
		for _, answer := range text {
			if len(answeredQuestion[answer]) == 0 {
				answerCount = answerCount + 1
				answeredQuestion[answer] = "exists"
			}
		}
	}
	allAnswerCounts = append(allAnswerCounts, answerCount)

	totalCount := 0
	for _, count := range allAnswerCounts {
		totalCount = totalCount + count
	}
	fmt.Println("One answered yes count", totalCount)

	totalCount = 0
	answeredQuestionCount := make(map[rune]int)
	personAnsweredInGroup := 0
	for _, text := range texts {
		if len(string(text)) == 0 {
			for _, answerCount := range answeredQuestionCount {
				if answerCount == personAnsweredInGroup {
					totalCount = totalCount + 1
				}
			}
			answeredQuestionCount = make(map[rune]int)
			personAnsweredInGroup = 0
		}
		for _, answer := range text {
			answeredQuestionCount[answer] = answeredQuestionCount[answer] + 1
		}
		if len(string(text)) > 0 {
			personAnsweredInGroup = personAnsweredInGroup + 1
		}
	}
	for _, answerCount := range answeredQuestionCount {
		if answerCount == personAnsweredInGroup {
			totalCount = totalCount + 1
		}
	}

	fmt.Println("All answered yes count", totalCount)

}
