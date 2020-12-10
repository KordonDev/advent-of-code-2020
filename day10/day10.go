package main

import (
	"fmt"
	"sort"
)

func main() {
	lines := readFile("./input.txt")

	var numbers []int
	for _, line := range lines {
		numbers = append(numbers, stringToInt(line))
	}
	sort.Ints(numbers)

	threeJoltsJumps := 1 // last jump
	oneJoltsJumps := 0
	for index, currentJolts := range numbers {
		lastJolts := 0
		if index != 0 {
			lastJolts = numbers[index-1]
		}
		if currentJolts-lastJolts == 1 {
			oneJoltsJumps++
		}
		if currentJolts-lastJolts == 3 {
			threeJoltsJumps++
		}
	}
	fmt.Println("result one", oneJoltsJumps*threeJoltsJumps)

	possibilities := 1
	for index, currentJolt := range numbers {
		possibileJolts := 0
		for _, nextJolt := range numbers[index+1 : rangeEnd(index, &numbers)] {
			if nextJolt-currentJolt <= 3 {
				possibileJolts++
			}
		}
		if possibileJolts != 0 {
			possibilities = possibilities * possibileJolts
		}
	}
	fmt.Println("Wrong Part2 without recursion", possibilities)

	sum := make(chan int)
	sliceStart := 0
	sliceStartJolt := 0
	slices := 1
	for index, currentJolt := range numbers {
		if index != len(numbers)-1 && numbers[index+1]-currentJolt == 3 {
			subSlice := numbers[sliceStart : index+1]
			go calculateSlicePosibilities(sliceStartJolt, &subSlice, sum)
			sliceStart = index + 1
			sliceStartJolt = currentJolt
			slices++
		}
	}
	subSlice := numbers[sliceStart:]
	go calculateSlicePosibilities(sliceStartJolt, &subSlice, sum)

	result := 1
	for i := 0; i < slices; i++ {
		next := <-sum
		result = result * next
	}
	fmt.Println("Part 2 fast", result)

	// part2Old(numbers)
}

func part2Old(numbers []int) {
	cutIndex := 64 // only works for this input data

	possibilitiesFirstHalf := 0
	firstHalf := numbers[0:cutIndex]
	addNextCharger(0, -1, &firstHalf, &possibilitiesFirstHalf)

	possibilitiesSecondHalf := 0
	secondHalf := numbers[cutIndex:]
	addNextCharger(numbers[cutIndex-1], -1, &secondHalf, &possibilitiesSecondHalf)

	fmt.Println("Second result", possibilitiesSecondHalf*possibilitiesFirstHalf)

}

func addNextCharger(currentJolts int, currentIndex int, adapterJolts *[]int, possibilities *int) {
	if currentJolts == (*adapterJolts)[len(*adapterJolts)-1] {
		(*possibilities) = (*possibilities) + 1
		return
	}

	for i, nextJolt := range (*adapterJolts)[currentIndex+1 : rangeEnd(currentIndex, adapterJolts)] {
		if nextJolt-currentJolts <= 3 {
			addNextCharger(nextJolt, currentIndex+i+1, adapterJolts, possibilities)
		}
	}
}
func rangeEnd(currentIndex int, adapterJolts *[]int) int {
	if (currentIndex + 4) < len(*adapterJolts) {
		return currentIndex + 4
	}
	return len(*adapterJolts)
}

func calculateSlicePosibilities(currentJolts int, adapterJolts *[]int, res chan int) {
	possibilities := 0
	addNextCharger(currentJolts, -1, adapterJolts, &possibilities)
	res <- possibilities
}
