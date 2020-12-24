package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	lines := readFile("./input.txt")

	tiles := make(map[string]string)
	for _, line := range lines {
		directions := getDirections(line)
		tileX, tileY := calculateEndTile(directions)

		color, found := tiles[key(tileX, tileY)]
		if found == false || color == "white" {
			tiles[key(tileX, tileY)] = "black"
		} else {
			tiles[key(tileX, tileY)] = "white"
		}
	}

	fmt.Println("Solution 1", countBlackTiles(&tiles))

	for day := 1; day <= 100; day++ {
		nextDayTile := make(map[string]string)
		for key := range tiles {
			x, y := keyToCoordinates(key)
			neighboursCoordinats := getNeighbours(x, y)
			for i := 0; i < len(neighboursCoordinats); i = i + 2 {
				calculateNextDayColor(neighboursCoordinats[i], neighboursCoordinats[i+1], &tiles, &nextDayTile)
			}
		}

		tiles = nextDayTile
	}
	fmt.Println("Solution 2", countBlackTiles(&tiles))
}

func getDirections(line string) []string {
	var directions []string
	var lastChar string
	for _, char := range line {
		if char == 'n' || char == 's' {
			lastChar = string(char)
			continue
		}
		directions = append(directions, lastChar+string(char))
		lastChar = ""
	}
	return directions
}

func calculateEndTile(directions []string) (int, int) {
	x := 0
	y := 0
	for _, direction := range directions {
		if direction == "e" {
			x = x + 2
		}
		if direction == "se" {
			x++
			y--
		}
		if direction == "sw" {
			x--
			y--
		}
		if direction == "w" {
			x = x - 2
		}
		if direction == "nw" {
			x--
			y++
		}
		if direction == "ne" {
			x++
			y++
		}
	}
	return x, y
}

func key(x int, y int) string {
	return strconv.Itoa(x) + "," + strconv.Itoa(y)
}

func keyToCoordinates(key string) (int, int) {
	parts := strings.Split(key, ",")
	return stringToInt(parts[0]), stringToInt(parts[1])
}

func calculateNextDayColor(x int, y int, currentTiles *(map[string]string), nextTiles *(map[string]string)) {
	neighbours := getNeighbours(x, y)
	blackNeighbours := 0
	for i := 0; i < len(neighbours); i = i + 2 {
		neighbourColor := (*currentTiles)[key(neighbours[i], neighbours[i+1])]
		if neighbourColor == "black" {
			blackNeighbours++
		}
	}
	tileColor := (*currentTiles)[key(x, y)]
	if tileColor == "black" && blackNeighbours > 0 && blackNeighbours < 3 {
		(*nextTiles)[key(x, y)] = "black"
	}
	if tileColor != "black" && blackNeighbours == 2 {
		(*nextTiles)[key(x, y)] = "black"
	}
}

func getNeighbours(x int, y int) []int {
	var result []int
	result = append(result, x+2)
	result = append(result, y)

	result = append(result, x+1)
	result = append(result, y-1)

	result = append(result, x-1)
	result = append(result, y-1)

	result = append(result, x-2)
	result = append(result, y)

	result = append(result, x-1)
	result = append(result, y+1)

	result = append(result, x+1)
	result = append(result, y+1)
	return result
}

func countBlackTiles(tiles *map[string]string) int {
	blackTiles := 0
	for _, tile := range *tiles {
		if tile == "black" {
			blackTiles++
		}
	}
	return blackTiles
}
