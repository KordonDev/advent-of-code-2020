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
	var intervals []int
	for index, intervalString := range intervalTimes {
		if intervalString != "x" {
			interval := stringToInt(intervalString)
			busIntervalsList[interval] = index
			intervals = append(intervals, interval)
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

	sortMaxToMin(&intervals)

	nextGab := 0
	base := 0
	nextBase := 0
	gab := intervals[0]
	search := intervals[0]
	for i := 1; i <= len(intervals); i++ {
		currentInterval := intervals[0:i]
		for nextGab == 0 {
			if areDeparturesValidForTime(search, &currentInterval, &busIntervalsList) {
				if nextBase == 0 {
					nextBase = search
				} else {
					nextGab = search - nextBase
				}
			}
			search = search + gab
		}
		gab = nextGab
		base = nextBase
		search = base
		nextBase = 0
		nextGab = 0
	}
	fmt.Println("Generic Solution 2", base+busIntervalsList[stringToInt(intervalTimes[0])])

	// solutionForMyInput(&busIntervalsList, highestIntervalIndex)

}
func areDeparturesValidForTime(checkTime int, intervals *[]int, intervalTimes *(map[int]int)) bool {
	for _, check := range *intervals {
		if (checkTime+(*intervalTimes)[check])%check != 0 {
			return false
		}
	}
	return true
}

func isListOfDepartures(checkTime int, intervalTimes *(map[int]int)) bool {
	for interval, offset := range *intervalTimes {
		if (checkTime+offset)%interval != 0 {
			return false
		}
	}
	return true
}

func solutionForMyInput(busIntervalsList *(map[int]int), highestIntervalIndex int) {
	gabBetweenBiggestNumbers := 0
	firstFound := 0
	search := 593
	for gabBetweenBiggestNumbers == 0 {
		if search%593 == 0 && (search-31)%433 == 0 && (search-41)%41 == 0 && (search-68)%37 == 0 && (search-2)%29 == 0 {
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

	for smallestTimestamp == 0 {
		if isListOfDepartures(checkTime, busIntervalsList) {
			smallestTimestamp = checkTime
		}
		checkTime = checkTime + gabBetweenBiggestNumbers
	}

	fmt.Println("Solution 2", smallestTimestamp-highestIntervalIndex)

}
