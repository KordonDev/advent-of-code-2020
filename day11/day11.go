package main

import (
	"fmt"
	"strconv"
)

func main() {
	lines := readFile("./input.txt")

	floor := make(map[(string)]bool)

	for y, line := range lines {
		for x, sign := range line {
			if string(sign) == "." {
				floor[key(x, y)] = true
			}
		}
	}

	maxX := len(lines[0]) - 1
	maxY := len(lines) - 1

	numberOfChanges := 1
	seated := make(map[(string)]bool)
	for numberOfChanges > 0 {
		nextSeated := make(map[(string)]bool)
		numberOfChanges = 0

		for y := 0; y <= maxY; y++ {
			for x := 0; x <= maxX; x++ {
				_, isFloor := floor[key(x, y)]
				if !isFloor {
					neibours := calculateNeibours(x, y, &seated)
					occupiedSeat := seated[key(x, y)]
					if occupiedSeat && neibours >= 4 {
						numberOfChanges++
					} else if !occupiedSeat && neibours == 0 {
						nextSeated[key(x, y)] = true
						numberOfChanges++
					} else {
						nextSeated[key(x, y)] = seated[key(x, y)]
					}
				}
			}
		}
		seated = nextSeated
	}

	solution := 0
	for _, occupiedSeat := range seated {
		if occupiedSeat {
			solution++
		}
	}

	fmt.Println("Solution 1:", solution)

	numberOfChanges = 1
	seated = make(map[(string)]bool)
	for numberOfChanges > 0 {
		nextSeated := make(map[(string)]bool)
		numberOfChanges = 0

		for y := 0; y <= maxY; y++ {
			for x := 0; x <= maxX; x++ {
				_, isFloor := floor[key(x, y)]
				if !isFloor {
					neibours := calculateSeeingNeibours(x, y, &seated, &floor)
					occupiedSeat := seated[key(x, y)]
					if occupiedSeat && neibours >= 5 {
						numberOfChanges++
					} else if !occupiedSeat && neibours == 0 {
						nextSeated[key(x, y)] = true
						numberOfChanges++
					} else {
						nextSeated[key(x, y)] = seated[key(x, y)]
					}
				}
			}
		}
		seated = nextSeated
	}

	solution = 0
	for _, occupiedSeat := range seated {
		if occupiedSeat {
			solution++
		}
	}

	fmt.Println("Solution 2:", solution)
}

func key(x int, y int) string {
	return strconv.Itoa(x) + "," + strconv.Itoa(y)
}

func calculateNeibours(x int, y int, seats *(map[string]bool)) int {
	numberOfNeibours := 0
	if occupiedSeat := (*seats)[key(x-1, y)]; occupiedSeat {
		numberOfNeibours++
	}
	if occupiedSeat := (*seats)[key(x-1, y-1)]; occupiedSeat {
		numberOfNeibours++
	}
	if occupiedSeat := (*seats)[key(x-1, y+1)]; occupiedSeat {
		numberOfNeibours++
	}
	if occupiedSeat := (*seats)[key(x+1, y)]; occupiedSeat {
		numberOfNeibours++
	}
	if occupiedSeat := (*seats)[key(x+1, y-1)]; occupiedSeat {
		numberOfNeibours++
	}
	if occupiedSeat := (*seats)[key(x+1, y+1)]; occupiedSeat {
		numberOfNeibours++
	}
	if occupiedSeat := (*seats)[key(x, y-1)]; occupiedSeat {
		numberOfNeibours++
	}
	if occupiedSeat := (*seats)[key(x, y+1)]; occupiedSeat {
		numberOfNeibours++
	}
	return numberOfNeibours
}

func calculateSeeingNeibours(x int, y int, seats *(map[string]bool), floor *(map[string]bool)) int {
	numberOfNeibours := 0
	if seesNeigbour(x, y, addOne, keep, seats, floor) {
		numberOfNeibours++
	}
	if seesNeigbour(x, y, addOne, addOne, seats, floor) {
		numberOfNeibours++
	}
	if seesNeigbour(x, y, addOne, substractOne, seats, floor) {
		numberOfNeibours++
	}
	if seesNeigbour(x, y, substractOne, keep, seats, floor) {
		numberOfNeibours++
	}
	if seesNeigbour(x, y, substractOne, substractOne, seats, floor) {
		numberOfNeibours++
	}
	if seesNeigbour(x, y, substractOne, addOne, seats, floor) {
		numberOfNeibours++
	}
	if seesNeigbour(x, y, keep, substractOne, seats, floor) {
		numberOfNeibours++
	}
	if seesNeigbour(x, y, keep, addOne, seats, floor) {
		numberOfNeibours++
	}
	return numberOfNeibours
}

func seesNeigbour(x int, y int, xfunc change, yFunc change, seats *(map[string]bool), floor *(map[string]bool)) bool {
	neigbourFound := false
	var occupiedSeat bool
	nextX := x
	nextY := y
	for !neigbourFound {
		nextX = xfunc(nextX)
		nextY = yFunc(nextY)
		if _, found := (*floor)[key(nextX, nextY)]; !found {
			neigbourFound = true
			occupiedSeat = (*seats)[key(nextX, nextY)]
		}
	}
	return occupiedSeat
}

type change func(int) int

func addOne(i int) int {
	return i + 1
}
func keep(i int) int {
	return i
}
func substractOne(i int) int {
	return i - 1
}
