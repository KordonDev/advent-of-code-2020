package main

import (
	"bufio"
	"log"
	"os"
	"sort"
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

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

// sort max to min
func sortMaxToMin(ints *[]int) {
	sort.Ints(*ints)
	sort.SliceStable(*ints, func(i, j int) bool {
		return j < i
	})
}

func sortMixToMax(ints *[]int) {
	sort.Ints(*ints)
}
