package main

import (
	"fmt"
	"log"
	"time"
)

type Node struct {
	value int
	next  *Node
	last  *Node
}

func (n *Node) AddNode(value int) *Node {
	newNode := Node{value, nil, nil}
	iter := n
	for iter.next != nil {
		iter = iter.next
	}
	iter.next = &newNode
	newNode.last = iter
	return &newNode
}

func (n *Node) PrintNode(max int) {
	iter := n
	iteration := 0
	for iter != nil && iteration < max {
		fmt.Println(iter.value)
		iter = iter.next
		iteration++
	}
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
