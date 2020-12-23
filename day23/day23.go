package main

import (
	"fmt"
	"strconv"
)

func main() {
	lines := readFile("./input.txt")

	var numbers []int
	for _, number := range lines[0] {
		numbers = append(numbers, stringToInt(string(number)))
	}

	currentIndex := 0
	rounds := 1
	for rounds <= 100 {
		insertNumber := numbers[currentIndex] - 1
		var numbersToMove []int
		for i := 1; i < 4; i++ {
			numbersToMove = append(numbersToMove, numbers[(currentIndex+i)%len(numbers)])
		}
		otherNumbers := make([]int, len(numbers)-3)
		otherNumbers[0] = numbers[currentIndex]
		for i := 1; i < len(numbers)-3; i++ {
			otherNumbers[i] = numbers[(currentIndex+3+i)%len(numbers)]
		}

		insertNumber = findNumberToInsertAfter(insertNumber, &otherNumbers)

		nextOrder := make([]int, len(numbers))
		inserted := 0
		insert := false
		for i := 0; i < len(numbers); i++ {
			insertIndex := (i + currentIndex) % len(numbers)
			numberToInsert := otherNumbers[(i-inserted+len(otherNumbers))%len(otherNumbers)]

			if insert && inserted < 3 {
				nextOrder[insertIndex] = numbersToMove[inserted]
				inserted++
			} else {
				nextOrder[insertIndex] = numberToInsert
			}

			if insertNumber == numberToInsert {
				insert = true
			}
		}

		numbers = nextOrder
		currentIndex = (currentIndex + 1) % len(numbers)
		rounds++
	}

	fmt.Println("Solution 1", numbersAfterOne(&numbers), numbers)

}

func findNumberToInsertAfter(insertNumber int, list *[]int) int {
	for true {
		for _, currentNumber := range *list {
			if currentNumber == insertNumber {
				return currentNumber
			}
		}
		insertNumber--
		if insertNumber <= 0 {
			insertNumber = 9
		}
	}
	return -1
}

func numbersAfterOne(list *[]int) string {
	var oneIndex int
	for index, number := range *list {
		if number == 1 {
			oneIndex = index
		}
	}

	result := ""
	for i := 1; i < len(*list); i++ {
		result = result + strconv.Itoa((*list)[(i+oneIndex)%len(*list)])
	}
	return result
}

func subgame(player1 []int, player2 []int) (int, []int) {
	var prevCardsPlayer1 [][]int
	var prevCardsPlayer2 [][]int
	var card1 int
	var card2 int
	var winner int
	for player1[0] != player2[len(player2)-1] && player2[0] != player1[len(player1)-1] {
		// fmt.Println("round", player1, player2)
		if gameHasBeenPlayed(&prevCardsPlayer1, &prevCardsPlayer2, player1, player2) {
			fmt.Println("same game")
			return 1, player1
		}
		winner = -1
		sub := false
		card1, player1 = player1[0], player1[min(1, len(player1)-1):]
		card2, player2 = player2[0], player2[min(1, len(player2)-1):]
		if card1 <= len(player1) && card2 <= len(player2) {
			p1 := copyList(player1[0:card1])
			p2 := copyList(player2[0:card2])
			fmt.Println("round", player1, player2, card1, card2)
			fmt.Println("start subgame", p1, p2)
			winner, _ = subgame(p1, p2)
			sub = true
		}
		if winner == 1 || winner == -1 && card1 > card2 {
			player1 = append(player1, []int{card1, card2}...)
			winner = 1
		}
		if winner == 2 || winner == -1 && card2 > card1 {
			player2 = append(player2, []int{card2, card1}...)
			winner = 2
		}
		if sub {
			fmt.Println("after subgame", player1, player2, winner)
		}
	}
	fmt.Println("end", player1, player2, winner)
	if winner == 1 {
		return winner, player1
	}
	return winner, player2
}

func gameHasBeenPlayed(prevCardsPlayer1 *[][]int, prevCardsPlayer2 *[][]int, player1 []int, player2 []int) bool {
	for i := 0; i < len(*prevCardsPlayer2); i++ {
		bothSame := true
		currentPrevPlayer1 := (*prevCardsPlayer1)[i]
		currentPrevPlayer2 := (*prevCardsPlayer2)[i]
		if len(currentPrevPlayer1) != len(player1) || len(currentPrevPlayer2) != len(player2) {
			continue
		}
		for j := 0; j < len(currentPrevPlayer1); j++ {
			if currentPrevPlayer1[j] != player1[j] {
				bothSame = false
				break
			}
		}
		for j := 0; j < len(currentPrevPlayer2); j++ {
			if currentPrevPlayer2[j] != player2[j] {
				bothSame = false
				break
			}
		}
		if bothSame {
			return true
		}
	}
	*prevCardsPlayer1 = append(*prevCardsPlayer1, copyList(player1))
	*prevCardsPlayer2 = append(*prevCardsPlayer2, copyList(player2))
	return false
}

func copyList(list []int) []int {
	l1 := make([]int, len(list))
	copy(l1, list)
	return l1
}

func calculteScore(winnerCards *[]int) int {
	score := 0
	multiplyer := len(*winnerCards)
	for _, card := range *winnerCards {
		score = score + multiplyer*card
		multiplyer--
	}
	return score
}
