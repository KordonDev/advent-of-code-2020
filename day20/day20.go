package main

import (
	"fmt"
	"strings"
)

type Tile struct {
	id      string
	right   *Tile
	left    *Tile
	top     *Tile
	bottom  *Tile
	content []string
}

func main() {
	lines := readFile("./input.txt")

	var unusedTiles []*Tile
	var nextTileID string
	var nextContent []string
	for _, line := range lines {
		if len(line) == 0 {
			nextTile := Tile{
				id:      nextTileID,
				content: nextContent,
			}
			unusedTiles = append(unusedTiles, &nextTile)
			nextContent = nil
		} else if strings.HasPrefix(line, "Tile") {
			parts := strings.Split(line, " ")
			nextTileID = strings.ReplaceAll(parts[1], ":", "")
		} else {
			nextContent = append(nextContent, line)
		}
	}
	nextTile := Tile{
		id:      nextTileID,
		content: nextContent,
	}
	unusedTiles = append(unusedTiles, &nextTile)

	startTile, unusedTiles := unusedTiles[0], unusedTiles[1:]

	// Add to right
	addedRight := true
	for addedRight {
		next := startTile
		for next.right != nil {
			next = next.right
		}
		addedRight, unusedTiles = addTileRight(next, unusedTiles)
	}

	// Add to left
	addedLeft := true
	for addedLeft {
		next := startTile
		for next.left != nil {
			next = next.left
		}
		addedLeft, unusedTiles = addTileLeft(next, unusedTiles)
	}
	leftTile := startTile
	for leftTile.left != nil {
		leftTile = leftTile.left
	}

	iter := leftTile
	for iter != nil {
		// Add to bottom
		bottomAdded := true
		for bottomAdded {
			next := iter
			for next.bottom != nil {
				next = next.bottom
			}
			bottomAdded, unusedTiles = addTileBottom(next, unusedTiles)
		}

		topAdded := true
		for topAdded {
			next := iter
			for next.top != nil {
				next = next.top
			}
			topAdded, unusedTiles = addTileTop(next, unusedTiles)
		}
		iter = iter.right
	}

	fmt.Println("")
	bigTile := createBigTile(startTile)
	printTile(&bigTile)
	seaMonster := searchSeaMonster(&bigTile)
	fmt.Println(seaMonster, "seaMonster found")
	seaMonsterHashes := 15

	fmt.Println(countHash(&bigTile) - seaMonster*seaMonsterHashes)
}

func addTileRight(tile *Tile, unusedTiles []*Tile) (bool, []*Tile) {
	topTile := tile.top
	if topTile != nil {
		topTile = topTile.right
	}
	bottobTile := tile.bottom
	if bottobTile != nil {
		bottobTile = bottobTile.right
	}

	for index, possibleTile := range unusedTiles {

		for flip := 0; flip < 3; flip++ {
			for side := 0; side < 4; side++ {

				fits := true
				if getRightBorder(tile) != getLeftBorder(possibleTile) {
					fits = false
				}
				if topTile != nil && getBottomBorder(topTile) != getTopBorder(possibleTile) {
					fits = false
				}
				if bottobTile != nil && getTopBorder(bottobTile) != getBottomBorder(possibleTile) {
					fits = false
				}
				if fits {
					(*tile).right = possibleTile
					(*possibleTile).left = tile
					(*possibleTile).top = topTile
					(*possibleTile).bottom = bottobTile
					if topTile != nil {
						(*topTile).bottom = possibleTile
					}
					if bottobTile != nil {
						(*bottobTile).top = possibleTile
					}
					unusedTiles = removeTile(&unusedTiles, index)
					return true, unusedTiles
				}

				rotate(possibleTile)
			}
			if flip == 0 {
				flipHorizontal(possibleTile)
			}
			if flip == 1 {
				flipVertical(possibleTile)
			}
		}

	}
	return false, unusedTiles
}

func addTileLeft(tile *Tile, unusedTiles []*Tile) (bool, []*Tile) {
	topTile := tile.top
	if topTile != nil {
		topTile = topTile.left
	}
	bottobTile := tile.bottom
	if bottobTile != nil {
		bottobTile = bottobTile.left
	}

	for index, possibleTile := range unusedTiles {

		for flip := 0; flip < 3; flip++ {
			for side := 0; side < 4; side++ {

				fits := true
				if getLeftBorder(tile) != getRightBorder(possibleTile) {
					fits = false
				}
				if topTile != nil && getBottomBorder(topTile) != getTopBorder(possibleTile) {
					fits = false
				}
				if bottobTile != nil && getTopBorder(bottobTile) != getBottomBorder(possibleTile) {
					fits = false
				}
				if fits {
					(*tile).left = possibleTile
					(*possibleTile).right = tile
					(*possibleTile).top = topTile
					(*possibleTile).bottom = bottobTile
					if topTile != nil {
						(*topTile).bottom = possibleTile
					}
					if bottobTile != nil {
						(*bottobTile).top = possibleTile
					}
					unusedTiles = removeTile(&unusedTiles, index)
					return true, unusedTiles
				}

				rotate(possibleTile)
			}
			if flip == 0 {
				flipHorizontal(possibleTile)
			}
			if flip == 1 {
				flipVertical(possibleTile)
			}
		}

	}
	return false, unusedTiles
}

func addTileBottom(tile *Tile, unusedTiles []*Tile) (bool, []*Tile) {
	rightTile := tile.right
	if rightTile != nil {
		rightTile = rightTile.bottom
	}
	leftTile := tile.left
	if leftTile != nil {
		leftTile = leftTile.bottom
	}

	for index, possibleTile := range unusedTiles {

		for flip := 0; flip < 3; flip++ {
			for side := 0; side < 4; side++ {

				fits := true
				if getBottomBorder(tile) != getTopBorder(possibleTile) {
					fits = false
				}
				if rightTile != nil && getLeftBorder(rightTile) != getRightBorder(possibleTile) {
					fits = false
				}
				if leftTile != nil && getRightBorder(leftTile) != getLeftBorder(possibleTile) {
					fits = false
				}
				if fits {
					(*tile).bottom = possibleTile
					(*possibleTile).top = tile
					(*possibleTile).right = rightTile
					(*possibleTile).left = leftTile
					if rightTile != nil {
						(*rightTile).left = possibleTile
					}
					if leftTile != nil {
						(*leftTile).right = possibleTile
					}
					unusedTiles = removeTile(&unusedTiles, index)
					return true, unusedTiles
				}

				rotate(possibleTile)
			}
			if flip == 0 {
				flipHorizontal(possibleTile)
			}
			if flip == 1 {
				flipVertical(possibleTile)
			}
		}

	}
	return false, unusedTiles
}

func addTileTop(tile *Tile, unusedTiles []*Tile) (bool, []*Tile) {
	rightTile := tile.right
	if rightTile != nil {
		rightTile = rightTile.top
	}
	leftTile := tile.left
	if leftTile != nil {
		leftTile = leftTile.top
	}

	for index, possibleTile := range unusedTiles {

		for flip := 0; flip < 3; flip++ {
			for side := 0; side < 4; side++ {

				fits := true
				if getTopBorder(tile) != getBottomBorder(possibleTile) {
					fits = false
				}
				if rightTile != nil && getLeftBorder(rightTile) != getRightBorder(possibleTile) {
					fits = false
				}
				if leftTile != nil && getRightBorder(leftTile) != getLeftBorder(possibleTile) {
					fits = false
				}
				if fits {
					(*tile).top = possibleTile
					(*possibleTile).bottom = tile
					(*possibleTile).right = rightTile
					(*possibleTile).left = leftTile
					if rightTile != nil {
						(*rightTile).left = possibleTile
					}
					if leftTile != nil {
						(*leftTile).right = possibleTile
					}
					unusedTiles = removeTile(&unusedTiles, index)
					return true, unusedTiles
				}

				rotate(possibleTile)
			}
			if flip == 0 {
				flipHorizontal(possibleTile)
			}
			if flip == 1 {
				flipVertical(possibleTile)
			}
		}

	}
	return false, unusedTiles
}
func searchSeaMonster(tile *Tile) int {
	for flip := 0; flip < 3; flip++ {
		for side := 0; side < 4; side++ {
			foundSeaMonsters := findMonster(tile)
			if foundSeaMonsters > 0 {
				return foundSeaMonsters
			}

			rotate(tile)
		}
		if flip == 0 {
			flipHorizontal(tile)
		}
		if flip == 1 {
			flipVertical(tile)
		}
	}
	return 0
}

func findMonster(tile *Tile) int {
	seaMonsterWidth := 20
	seaMonsterHeight := 3
	content := (*tile).content
	monsters := 0
	for row := 0; row+seaMonsterHeight <= len(content); row++ {
		for column := 0; column+seaMonsterWidth <= len(content[0]); column++ {
			row1 := content[row][column:]
			row2 := content[row+1][column:]
			row3 := content[row+2][column:]
			if isMonster(row1, row2, row3) {
				fmt.Println("Found monster at", column, row)
				monsters++
			}
		}
	}
	return monsters
}

func isMonster(row1 string, row2 string, row3 string) bool {
	if row1[18] != '#' {
		return false
	}
	row2Indices := []int{0, 5, 6, 11, 12, 17, 18, 19}
	for _, index := range row2Indices {
		if row2[index] != '#' {
			return false
		}
	}
	row3Indices := []int{1, 4, 7, 10, 13, 16}
	for _, index := range row3Indices {
		if row3[index] != '#' {
			return false
		}
	}
	return true
}

func createBigTile(startTile *Tile) Tile {
	topLeftTile := startTile
	for topLeftTile.left != nil {
		topLeftTile = topLeftTile.left
	}
	for topLeftTile.top != nil {
		topLeftTile = topLeftTile.top
	}
	iter := topLeftTile
	var content []string
	for iter != nil {
		nextRows := getFullRow(iter)
		content = append(content, nextRows[1:len(nextRows)-1]...)
		iter = iter.bottom
	}
	return Tile{
		id:      "full",
		content: content,
	}
}

func getFullRow(leftTile *Tile) []string {
	var content []string
	for i := 0; i < len((*leftTile).content); i++ {
		content = append(content, "")
	}
	iter := leftTile
	for iter != nil {
		for index, row := range (*iter).content {
			content[index] = content[index] + row[1:len(row)-1]
		}
		iter = iter.right
	}
	return content
}

func countHash(tile *Tile) int {
	result := 0
	for _, row := range (*tile).content {
		result = result + strings.Count(row, "#")
	}
	return result
}

func removeTile(list *[]*Tile, index int) []*Tile {
	(*list)[index] = (*list)[len(*list)-1]
	_, newList := (*list)[len(*list)-1], (*list)[:len(*list)-1]
	return newList
}

func getTopBorder(tile *Tile) string {
	return (*tile).content[0]
}

func getBottomBorder(tile *Tile) string {
	return (*tile).content[len((*tile).content)-1]
}

func getLeftBorder(tile *Tile) string {
	return getColumn(tile, 0)
}

func getRightBorder(tile *Tile) string {
	return getColumn(tile, len((*tile).content[0])-1)
}

func rotate(tile *Tile) {
	var newContent []string
	for i := 1; i <= len((*tile).content[0]); i++ {
		newContent = append(newContent, getColumn(tile, len((*tile).content[0])-i))
	}
	(*tile).content = newContent
}

func getColumn(tile *Tile, columnNumber int) string {
	column := ""
	for _, row := range (*tile).content {
		column = column + string(row[columnNumber])
	}
	return column
}

func flipHorizontal(tile *Tile) {
	var newContent []string
	for i := 0; i < len((*tile).content); i++ {
		newContent = append(newContent, (*tile).content[len((*tile).content)-1-i])
	}
	(*tile).content = newContent
}

func flipVertical(tile *Tile) {
	var newContent []string
	for _, row := range (*tile).content {
		newContent = append(newContent, Reverse(row))
	}
	(*tile).content = newContent
}

func Reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

func printTile(tile *Tile) {
	fmt.Println((*tile).id)
	for _, row := range (*tile).content {
		fmt.Println(row)
	}
}
