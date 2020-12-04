package main

import (
	"bufio"
	"log"
	"os"
)

func readFile() []string {
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
	return texts
}
