package main

import (
	"fmt"
	"strconv"
)

func main() {
	lines := readFile("./input.txt")

	z := 0
	initialWidth := len(lines)
	activePositions := make(map[string]bool)
	activePositions4D := make(map[string]bool)

	for y, line := range lines {
		for x, char := range line {
			if char == '#' {
				activePositions[key(x, y, z)] = true
				activePositions4D[key4(x, y, z, 0)] = true
			}
		}
	}

	goalGenerations := 6
	for generation := 1; generation <= goalGenerations; generation++ {
		nextActivePositions := make(map[string]bool)

		for z := generation * -1; z < generation+1; z++ {
			for y := generation * -1; y < generation+initialWidth; y++ {
				for x := generation * -1; x < generation+initialWidth; x++ {
					activeNeighours := countNeighbours(x, y, z, &activePositions)
					currentPositionActive := activePositions[key(x, y, z)]
					if currentPositionActive && (activeNeighours == 2 || activeNeighours == 3) {
						nextActivePositions[key(x, y, z)] = true
					}
					if !currentPositionActive && activeNeighours == 3 {
						nextActivePositions[key(x, y, z)] = true
					}
				}

			}
		}
		activePositions = nextActivePositions
	}

	solution := 0
	for z := goalGenerations * -1; z < goalGenerations+1; z++ {
		for y := goalGenerations * -1; y < goalGenerations+initialWidth; y++ {
			for x := goalGenerations * -1; x < goalGenerations+initialWidth; x++ {
				if active := (activePositions)[key(x, y, z)]; active {
					solution++
				} else {
				}
			}
		}
	}
	fmt.Println("Solution 1", solution)

	for generation := 1; generation <= goalGenerations; generation++ {
		nextActivePositions4D := make(map[string]bool)

		for z := generation * -1; z < generation+1; z++ {
			for y := generation * -1; y < generation+initialWidth; y++ {
				for x := generation * -1; x < generation+initialWidth; x++ {
					for w := generation * -1; w <= generation+1; w++ {
						activeNeighours := 0
						activeNeighours = activeNeighours + countNeighbours4D(x, y, z, w, &activePositions4D)
						currentPositionActive := activePositions4D[key4(x, y, z, w)]
						if currentPositionActive && (activeNeighours == 2 || activeNeighours == 3) {
							nextActivePositions4D[key4(x, y, z, w)] = true
						}
						if !currentPositionActive && activeNeighours == 3 {
							nextActivePositions4D[key4(x, y, z, w)] = true
						}
					}

				}

			}
		}
		activePositions4D = nextActivePositions4D
	}
	fmt.Println(len(activePositions4D))

	solution = 0
	for w := goalGenerations * -1; w <= goalGenerations+1; w++ {
		for z := goalGenerations * -1; z < goalGenerations+1; z++ {
			for y := goalGenerations * -1; y < goalGenerations+initialWidth; y++ {
				for x := goalGenerations * -1; x < goalGenerations+initialWidth; x++ {
					if active := activePositions4D[key4(x, y, z, w)]; active {
						solution++
					}
				}
			}
		}
	}
	fmt.Println("Solution 2", solution)

}

func countNeighbours(x int, y int, z int, activePositions *(map[string]bool)) int {
	numberOfNeibours := 0
	numberOfNeibours = numberOfNeibours + countNeighbours2D(x, y, z-1, activePositions)
	numberOfNeibours = numberOfNeibours + countNeighbours2D(x, y, z, activePositions)
	if x == 0 && y == 1 && z == -1 {
	}
	numberOfNeibours = numberOfNeibours + countNeighbours2D(x, y, z+1, activePositions)
	if x == 0 && y == 1 && z == -1 {
	}
	if active := (*activePositions)[key(x, y, z)]; active {
		numberOfNeibours = numberOfNeibours - 1
	}
	return numberOfNeibours
}

func countNeighbours2D(x int, y int, z int, activePositions *(map[string]bool)) int {
	numberOfNeibours := 0
	if active := (*activePositions)[key(x-1, y, z)]; active {
		numberOfNeibours++
	}
	if active := (*activePositions)[key(x-1, y-1, z)]; active {
		numberOfNeibours++
	}
	if active := (*activePositions)[key(x-1, y+1, z)]; active {
		numberOfNeibours++
	}
	if active := (*activePositions)[key(x+1, y, z)]; active {
		numberOfNeibours++
	}
	if active := (*activePositions)[key(x+1, y-1, z)]; active {
		numberOfNeibours++
	}
	if active := (*activePositions)[key(x+1, y+1, z)]; active {
		numberOfNeibours++
	}
	if active := (*activePositions)[key(x, y-1, z)]; active {
		numberOfNeibours++
	}
	if active := (*activePositions)[key(x, y+1, z)]; active {
		numberOfNeibours++
	}
	if active := (*activePositions)[key(x, y, z)]; active {
		numberOfNeibours++
	}
	return numberOfNeibours
}

func key(x int, y int, z int) string {
	return strconv.Itoa(x) + "," + strconv.Itoa(y) + "," + strconv.Itoa(z)
}

func key4(x int, y int, z int, w int) string {
	return strconv.Itoa(x) + "," + strconv.Itoa(y) + "," + strconv.Itoa(z) + "," + strconv.Itoa(w)
}

func countNeighbours4D(x int, y int, z int, w int, activePositions *(map[string]bool)) int {
	numberOfNeibours := 0
	numberOfNeibours = numberOfNeibours + countNeighbours3D4(x, y, z, w-1, activePositions)
	numberOfNeibours = numberOfNeibours + countNeighbours3D4(x, y, z, w, activePositions)
	numberOfNeibours = numberOfNeibours + countNeighbours3D4(x, y, z, w+1, activePositions)
	if active := (*activePositions)[key4(x, y, z, w)]; active {
		numberOfNeibours = numberOfNeibours - 1
	}

	return numberOfNeibours
}

func countNeighbours3D4(x int, y int, z int, w int, activePositions *(map[string]bool)) int {
	numberOfNeibours := 0
	numberOfNeibours = numberOfNeibours + countNeighbours2D4(x, y, z-1, w, activePositions)
	numberOfNeibours = numberOfNeibours + countNeighbours2D4(x, y, z, w, activePositions)
	numberOfNeibours = numberOfNeibours + countNeighbours2D4(x, y, z+1, w, activePositions)
	return numberOfNeibours
}

func countNeighbours2D4(x int, y int, z int, w int, activePositions *(map[string]bool)) int {
	numberOfNeibours := 0
	if active := (*activePositions)[key4(x-1, y, z, w)]; active {
		numberOfNeibours++
	}
	if active := (*activePositions)[key4(x-1, y-1, z, w)]; active {
		numberOfNeibours++
	}
	if active := (*activePositions)[key4(x-1, y+1, z, w)]; active {
		numberOfNeibours++
	}
	if active := (*activePositions)[key4(x+1, y, z, w)]; active {
		numberOfNeibours++
	}
	if active := (*activePositions)[key4(x+1, y-1, z, w)]; active {
		numberOfNeibours++
	}
	if active := (*activePositions)[key4(x+1, y+1, z, w)]; active {
		numberOfNeibours++
	}
	if active := (*activePositions)[key4(x, y-1, z, w)]; active {
		numberOfNeibours++
	}
	if active := (*activePositions)[key4(x, y+1, z, w)]; active {
		numberOfNeibours++
	}
	if active := (*activePositions)[key4(x, y, z, w)]; active {
		numberOfNeibours++
	}
	return numberOfNeibours
}
