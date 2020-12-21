package main

import (
	"fmt"
	"strings"
)

type Tile struct {
	id               int
	borders          []string
	rotate           int
	xFlip            bool
	yFlip            bool
	topNeigbourID    int
	rightNeigbourID  int
	bottomNeigbourID int
	leftNeigbourID   int
}

func main() {
	lines := readFile("./test-input.txt")

	tiles := make(map[int]Tile)
	var nextTileID int
	var nextBorders []string
	for _, line := range lines {
		if len(line) == 0 {
			tiles[nextTileID] = Tile{
				id:               nextTileID,
				borders:          nextBorders,
				rotate:           0,
				xFlip:            false,
				yFlip:            false,
				topNeigbourID:    -1,
				rightNeigbourID:  -1,
				bottomNeigbourID: -1,
				leftNeigbourID:   -1,
			}
			nextBorders = nil
			nextTileID = 0
		}
		if strings.HasPrefix(line, "Tile") {
			nextTileID = stringToInt(line[5 : len(line)-1])
		}
		if strings.HasPrefix(line, ".") || strings.HasPrefix(line, "#") {
			if len(nextBorders) == 0 {
				nextBorders = append(nextBorders, line)
				nextBorders = append(nextBorders, "")
				nextBorders = append(nextBorders, "")
				nextBorders = append(nextBorders, "")
			}
			if len(nextBorders[1]) == len(line)-1 {
				nextBorders[2] = line
			}
			nextBorders[1] = nextBorders[1] + string(line[len(line)-1])
			nextBorders[3] = nextBorders[3] + string(line[0])
		}
	}
	tiles[nextTileID] = Tile{
		id:               nextTileID,
		borders:          nextBorders,
		rotate:           0,
		xFlip:            false,
		yFlip:            false,
		topNeigbourID:    -1,
		rightNeigbourID:  -1,
		bottomNeigbourID: -1,
		leftNeigbourID:   -1,
	}

	var tileCornerID int
	solution := 1
	for tileIDtoCheck, tileToCheck := range tiles {
		neighborIDs := make(map[int]bool)

		for tileID, tile := range tiles {
			if tileID != tileIDtoCheck {
				if compareWithRotation(&tileToCheck, &tile) {
					neighborIDs[tileID] = true
				}
			}
		}
		if len(neighborIDs) == 2 {
			fmt.Println("corner", tileIDtoCheck)
			solution = solution * tileIDtoCheck
			tileCornerID = tileIDtoCheck
		}
	}

	fmt.Println("Solution 1:", solution)

	fmt.Println("start", tiles[tileCornerID])
	/*
		for tileIDtoCheck, tileToCheck := range tiles {
			neighborIDs := make(map[int]bool)

			for tileID, tile := range tiles {
				if tileID != tileIDtoCheck {
					if compareWithRotation(&tileToCheck, &tile) {
						neighborIDs[tileID] = true
					}
				}
			}
			if len(neighborIDs) == 2 {
				fmt.Println("corner", tileIDtoCheck)
				solution = solution * tileIDtoCheck
			}
		}*/

}

func compareAllFourSides(fixedTile *Tile, tileToCompare *Tile) bool {
	for rotate := 0; rotate <= 3; rotate++ {
		if (*fixedTile).borders[0] == (*tileToCompare).borders[(rotate)%4] {
			(*tileToCompare).rotate = rotate
			(*fixedTile).topNeigbourID = 4 //(*tileToCompare).id
			(*tileToCompare).bottomNeigbourID = (*fixedTile).id
			fmt.Println((*tileToCompare).id, "0", *fixedTile)
			return true
		}
		if (*fixedTile).borders[1] == (*tileToCompare).borders[(1+rotate)%4] {
			(*tileToCompare).rotate = rotate
			(*fixedTile).rightNeigbourID = (*tileToCompare).id
			(*tileToCompare).leftNeigbourID = (*fixedTile).id
			fmt.Println((*tileToCompare).id, "1")
			return true
		}
		if (*fixedTile).borders[2] == (*tileToCompare).borders[(2+rotate)%4] {
			(*tileToCompare).rotate = rotate
			(*fixedTile).bottomNeigbourID = (*tileToCompare).id
			(*tileToCompare).topNeigbourID = (*fixedTile).id
			fmt.Println((*tileToCompare).id, "2")
			return true
		}
		if (*fixedTile).borders[3] == (*tileToCompare).borders[(3+rotate)%4] {
			(*tileToCompare).rotate = rotate
			(*fixedTile).leftNeigbourID = (*tileToCompare).id
			(*tileToCompare).rightNeigbourID = (*fixedTile).id
			fmt.Println((*tileToCompare).id, "3")
			return true
		}
	}
	return false
}

func compareWithRotation(fixedTile *Tile, tileToCompare *Tile) bool {
	borders := (*tileToCompare).borders
	tileToCompareIsConnected := hasNeighbour(tileToCompare)

	if compareAllFourSides(fixedTile, tileToCompare) {
		return true
	}

	if tileToCompareIsConnected {
		return false
	}

	xFlipedBorders := []string{Reverse(borders[0]), borders[3], Reverse(borders[2]), borders[1]}
	(*tileToCompare).borders = xFlipedBorders
	(*tileToCompare).xFlip = true
	if compareAllFourSides(fixedTile, tileToCompare) {
		return true
	}
	yFlipedBorders := []string{borders[2], Reverse(borders[1]), borders[0], Reverse(borders[3])}
	(*tileToCompare).borders = yFlipedBorders
	(*tileToCompare).xFlip = false
	(*tileToCompare).yFlip = true
	if compareAllFourSides(fixedTile, tileToCompare) {
		return true
	}
	(*tileToCompare).borders = borders
	(*tileToCompare).yFlip = false

	return false
}

func hasNeighbour(tile *Tile) bool {
	return (*tile).topNeigbourID != -1 || (*tile).rightNeigbourID != -1 || (*tile).bottomNeigbourID != -1 || (*tile).leftNeigbourID != -1
}

func Reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}
