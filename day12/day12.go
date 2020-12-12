package main

import (
	"fmt"
	"strings"
)

func main() {
	lines := readFile("./input.txt")

	northSouth := 0
	easthWest := 0
	shipDirection := "E"
	turnLeft := []string{"E", "N", "W", "S"}
	leftIndex := map[string]int{"E": 0, "N": 1, "W": 2, "S": 3}
	turnRight := []string{"E", "S", "W", "N"}
	rightIndex := map[string]int{"E": 0, "S": 1, "W": 2, "N": 3}

	for _, navigation := range lines {
		direction, steps := directionWithSteps(navigation)
		if direction == "F" {
			direction = shipDirection
		}
		if direction == "E" {
			easthWest = easthWest + steps
		}
		if direction == "W" {
			easthWest = easthWest - steps
		}
		if direction == "N" {
			northSouth = northSouth + steps
		}
		if direction == "S" {
			northSouth = northSouth - steps
		}
		if direction == "L" {
			shipDirection = turnLeft[(leftIndex[shipDirection]+steps/90)%len(turnLeft)]
		}
		if direction == "R" {
			shipDirection = turnRight[(rightIndex[shipDirection]+steps/90)%len(turnRight)]
		}
	}
	fmt.Println("Solution 1", abs(northSouth)+abs(easthWest))

	shipNorthSouth := 0
	shipEasthWest := 0
	targetNorthSouth := 1
	targetEasthWest := 10

	for _, navigation := range lines {
		direction, steps := directionWithSteps(navigation)
		if direction == "F" {
			shipNorthSouth = shipNorthSouth + targetNorthSouth*steps
			shipEasthWest = shipEasthWest + targetEasthWest*steps
		}
		if direction == "E" {
			targetEasthWest = targetEasthWest + steps
		}
		if direction == "W" {
			targetEasthWest = targetEasthWest - steps
		}
		if direction == "N" {
			targetNorthSouth = targetNorthSouth + steps
		}
		if direction == "S" {
			targetNorthSouth = targetNorthSouth - steps
		}
		if direction == "L" {
			if (steps / 90 % 4) == 1 {
				oldNorthSouth := targetNorthSouth
				targetNorthSouth = targetEasthWest * 1
				targetEasthWest = oldNorthSouth * -1
			}
			if (steps / 90 % 4) == 2 {
				targetNorthSouth = targetNorthSouth * -1
				targetEasthWest = targetEasthWest * -1
			}
			if (steps / 90 % 4) == 3 {
				oldNorthSouth := targetNorthSouth
				targetNorthSouth = targetEasthWest * -1
				targetEasthWest = oldNorthSouth * 1
			}
		}
		if direction == "R" {
			if (steps / 90 % 4) == 1 {
				oldNorthSouth := targetNorthSouth
				targetNorthSouth = targetEasthWest * -1
				targetEasthWest = oldNorthSouth * 1
			}
			if (steps / 90 % 4) == 2 {
				targetNorthSouth = targetNorthSouth * -1
				targetEasthWest = targetEasthWest * -1
			}
			if (steps / 90 % 4) == 3 {
				oldNorthSouth := targetNorthSouth
				targetNorthSouth = targetEasthWest * 1
				targetEasthWest = oldNorthSouth * -1
			}
		}
	}
	fmt.Println("Solution 2", abs(shipNorthSouth)+abs(shipEasthWest))
}

func directionWithSteps(s string) (string, int) {
	split := strings.SplitN(s, "", 2)
	direction := split[0]
	steps := stringToInt(split[1])
	return direction, steps
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
