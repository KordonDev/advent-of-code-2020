package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func readFile(filePath string) []string {
	file, err := os.Open(filePath)
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

func stringToInt(s string) int {
	value, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return value
}
