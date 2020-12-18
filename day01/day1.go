package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	var numbers []int
	for scanner.Scan() {
		var inputStr = scanner.Text()
		num, err := strconv.Atoi(inputStr)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, num)
	}
	file.Close()

	for outerIndex, outerValue := range numbers {
		for _, innerValue := range numbers[outerIndex+1:] {
			if outerValue+innerValue == 2020 {
				fmt.Println("Result is for two numbers is ", outerValue*innerValue)
			}
		}
	}

	for firstIndex, firstValue := range numbers {
		for secondIndex, secondValue := range numbers[firstIndex+1:] {
			for _, thirdValue := range numbers[secondIndex+1:] {
				if firstValue+secondValue+thirdValue == 2020 {
					fmt.Println("Result is for three numbers is ", firstValue, secondValue, thirdValue)
					fmt.Println("Result is for three numbers is ", firstValue*secondValue*thirdValue)
				}
			}
		}
	}
}
