package main

import "fmt"

func main() {
	lines := readFile("./input.txt")

	var currentNode *Node
	var lastNode *Node
	var oneNode *Node
	maxSize := 1000000
	lookup := make(map[int]*Node)
	for _, number := range lines[0] {
		num := stringToInt(string(number))
		if currentNode == nil {
			currentNode = &Node{num, nil, nil}
			lastNode = currentNode
		} else {
			lastNode = lastNode.AddNode(num)
		}
		lookup[num] = lastNode
		if num == 1 {
			oneNode = lastNode
		}
	}
	for i := len(lines[0]); i < maxSize; i++ {
		lastNode = lastNode.AddNode(i + 1)
		lookup[i+1] = lastNode
	}
	lastNode.next = currentNode
	currentNode.last = lastNode

	rounds := 1
	for rounds <= 10000000 {
		pickup := currentNode.next
		endPickup := pickup.next.next

		var insertAfter *Node
		searchValue := currentNode.value - 1
		for insertAfter == nil {
			if searchValue <= 0 {
				searchValue = maxSize
			}
			insertAfter = lookup[searchValue]
			if insertAfter == pickup || insertAfter == pickup.next || insertAfter == pickup.next.next {
				insertAfter = nil
			}
			searchValue--
		}

		currentNode.next = endPickup.next
		endPickup.next.last = currentNode

		endPickup.next = insertAfter.next
		insertAfter.next.last = endPickup

		insertAfter.next = pickup
		pickup.last = insertAfter

		currentNode = currentNode.next
		rounds++
	}

	fmt.Println("Solution", oneNode.next.value*oneNode.next.next.value)
}
