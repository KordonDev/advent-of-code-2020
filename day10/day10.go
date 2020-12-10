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
	fmt.Println(numbers)

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

	cutIndex := 64

	possibilitiesFirstHalf := 0
	firstHalf := numbers[0:cutIndex]
	addNextCharger(0, -1, &firstHalf, &possibilitiesFirstHalf)

	possibilitiesSecondHalf := 0
	secondHalf := numbers[cutIndex:]
	addNextCharger(numbers[cutIndex-1], -1, &secondHalf, &possibilitiesSecondHalf)

	fmt.Println(firstHalf, secondHalf)
	fmt.Println(possibilitiesFirstHalf, "*", possibilitiesSecondHalf)
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
