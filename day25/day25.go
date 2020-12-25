package main

import (
	"fmt"
)

func main() {
	lines := readFile("./input.txt")

	publicKey1 := stringToInt(lines[0])
	publicKey2 := stringToInt(lines[1])
	loopCount1 := findLoopSize(publicKey1)
	loopCount2 := findLoopSize(publicKey2)

	encrypitonKey1 := calculateEncryptionKey(publicKey1, loopCount2)
	encrypitonKey2 := calculateEncryptionKey(publicKey2, loopCount1)
	fmt.Println("Solution 1", loopCount1, loopCount2, encrypitonKey1, encrypitonKey2)
}

func findLoopSize(targetValue int) int {
	value := 1
	for loopCount := 1; ; loopCount++ {
		value = loop(value, 7)
		if value == targetValue {
			return loopCount
		}
	}
}

func loop(value int, multiplier int) int {
	value = value * multiplier
	return value % 20201227
}

func calculateEncryptionKey(publicKey int, numberOfLoops int) int {
	value := 1
	for i := 0; i < numberOfLoops; i++ {
		value = loop(value, publicKey)
	}
	return value
}
