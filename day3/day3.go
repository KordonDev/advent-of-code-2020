package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func countTrees(xIncrement int, yIncrement int, texts []string) int {

	xPosition := 0
	numberOfTrees := 0
	for i := 0; i < len(texts); i = i + yIncrement {
		text := texts[i]
		currentPosition := string(text[xPosition%len(text)])
		if currentPosition == "#" {
			numberOfTrees = numberOfTrees + 1
		}
		xPosition = xPosition + xIncrement
	}

	fmt.Println("Number of trees for Right", xIncrement, ", down", yIncrement, ":", numberOfTrees)
	return numberOfTrees
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	var texts []string
	for scanner.Scan() {
		texts = append(texts, scanner.Text())
	}
	file.Close()

	totalCount := countTrees(1, 1, texts) * countTrees(3, 1, texts) * countTrees(5, 1, texts) * countTrees(7, 1, texts) * countTrees(1, 2, texts)
	fmt.Println("Total number of trees", totalCount)
}
