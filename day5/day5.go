package main

import (
	"fmt"
)

func main() {
	texts := readFile("./input.txt")

	seatIds := make(map[int]string)
	highestSeatId := 0
	for _, text := range texts {
		minRow := 0
		maxRow := 127
		minColumn := 0
		maxColumn := 7
		for _, direction := range text {
			// Back
			if string(direction) == "B" {
				minRow = (maxRow-minRow)/2 + minRow + 1
			}
			if string(direction) == "F" {
				maxRow = (maxRow-minRow)/2 + minRow
			}
			if string(direction) == "R" {
				minColumn = (maxColumn-minColumn)/2 + minColumn + 1
			}
			if string(direction) == "L" {
				maxColumn = (maxColumn-minColumn)/2 + minColumn
			}
		}
		seatId := minRow*8 + minColumn
		if seatId > highestSeatId {
			highestSeatId = seatId
		}
		if minRow != 0 && minRow != 127 {
			seatIds[seatId] = "exists"
		}
	}
	fmt.Println("Highest seat number is", highestSeatId)

	for i := 0; i < highestSeatId; i++ {
		if len(seatIds[i-1]) > 0 && len(seatIds[i+1]) > 0 && len(seatIds[i]) == 0 {
			fmt.Println("Your seat number is", i)

		}

	}
}
