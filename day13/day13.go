package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	lines := readFile("./input.txt")

	arrivalTime := stringToInt(lines[0])
	earliesDeparture := math.MaxInt32
	result := 0

	intervalTimes := strings.Split(lines[1], ",")
	for _, intervalString := range intervalTimes {
		if intervalString != "x" {
			interval := stringToInt(intervalString)
			busIntervalDriven := math.Ceil(float64(arrivalTime) / float64(interval))
			busDepartureTime := int(busIntervalDriven) * interval
			if busDepartureTime < earliesDeparture {
				earliesDeparture = busDepartureTime
				result = (earliesDeparture - arrivalTime) * interval
			}
		}
	}

	fmt.Println("Solution 1", result)

	highestInterval := 0
	highestIntervalIndex := 0
	busIntervalsList := make(map[int]int)
	for index, intervalString := range intervalTimes {
		if intervalString != "x" {
			interval := stringToInt(intervalString)
			busIntervalsList[interval] = index
			if interval > highestInterval {
				highestInterval = interval
				highestIntervalIndex = index
			}
		}
	}

	// Fix index
	for interval, index := range busIntervalsList {
		busIntervalsList[interval] = index - highestIntervalIndex
	}
	fmt.Println(busIntervalsList)

	gabBetweenBiggestNumbers := 0
	firstFound := 0
	search := 593
	for gabBetweenBiggestNumbers == 0 {
		if search%593 == 0 && (search-31)%433 == 0 && (search-41)%41 == 0 && (search-68)%37 == 0 && (search-2)%29 == 0 { // && (search-23)%23 == 0 && (search-12)%19 == 0 {
			if firstFound == 0 {
				firstFound = search
			} else {
				gabBetweenBiggestNumbers = search - firstFound
			}
		}
		search = search + 593
	}
	fmt.Println("kkv", gabBetweenBiggestNumbers, (gabBetweenBiggestNumbers-31)%433)

	smallestTimestamp := 0
	checkTime := firstFound

	//58420380058973072 is too big
	for smallestTimestamp == 0 && checkTime < 100000000000000000 {
		if isListOfDepartures(checkTime, &busIntervalsList) {
			smallestTimestamp = checkTime
		}
		checkTime = checkTime + gabBetweenBiggestNumbers
	}

	fmt.Println("Solution 2", smallestTimestamp-highestIntervalIndex)
}

func isListOfDepartures(checkTime int, intervalTimes *(map[int]int)) bool {
	for interval, offset := range *intervalTimes {
		if (checkTime+offset)%interval != 0 {
			return false
		}
	}
	return true
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
