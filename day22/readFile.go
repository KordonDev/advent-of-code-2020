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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
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

func slice(array []string, start int, end int, insert string) []string {
	a1 := append(array[:start], insert)
	if end >= len(array) {
		return a1
	}
	return append(a1, array[end:]...)
}

func sliceArray(array []string, start int, end int, insert []string) []string {
	a1 := append(array[:start], insert...)
	if end >= len(array) {
		return a1
	}
	return append(a1, array[end:]...)
}

func removeItems(list1 *[]string, list2 *[]string) ([]string, bool) {
	var result []string
	change := false
	for _, item1 := range *list1 {
		contains := false
		for _, item2 := range *list2 {
			if item1 == item2 {
				contains = true
				break
			}
		}
		if !contains {
			result = append(result, item1)
		} else {
			change = true
		}
	}
	return result, change
}
