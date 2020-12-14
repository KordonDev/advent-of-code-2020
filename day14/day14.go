package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	lines := readFile("./input.txt")

	memory := make(map[int]string)
	mask := ""

	for _, line := range lines {
		if strings.HasPrefix(line, "mask") {
			mask = line[7:]
		} else {
			parts := strings.Split(line, " = ")
			memoryAddress := stringToInt(parts[0][4:(len(parts[0]) - 1)])
			addition := base10ToBase2String(parts[1])
			withMaskApplied := applyMask(addition, mask)
			memory[memoryAddress] = withMaskApplied
		}
	}

	var result int64
	result = 0
	for _, value := range memory {
		result = result + stringBase2ToInt(value)
	}
	fmt.Println("Solution part 1", result)

	mask = ""
	memory2 := make(map[string]string)
	for _, line := range lines {
		if strings.HasPrefix(line, "mask") {
			mask = line[7:]
		} else {
			parts := strings.Split(line, " = ")
			memoryAddress := base10ToBase2String(parts[0][4:(len(parts[0]) - 1)])
			addition := base10ToBase2String(parts[1])
			memoryAddressWithMaskApplied := applyMaskToAddress(memoryAddress, mask)
			memoryAddresses := getAllMemoryAddresses(memoryAddressWithMaskApplied)
			for _, memoryAddress := range memoryAddresses {
				memory2[memoryAddress] = addition
			}
		}
	}

	result = 0
	for _, value := range memory2 {
		result = result + stringBase2ToInt(value)
	}
	fmt.Println("Solution part 2", result)
}

func applyMask(number string, mask string) string {
	offset := len(mask) - len(number)
	result := ""
	for index, char := range mask {
		if char == 'X' {
			if index >= offset {
				result = result + string(number[index-offset])
			} else {
				result = result + "0"
			}
		}
		if char == '1' {
			result = result + "1"
		}
		if char == '0' {
			result = result + "0"
		}
	}
	return result
}
func applyMaskToAddress(number string, mask string) string {
	offset := len(mask) - len(number)
	result := ""
	for index, char := range mask {
		if char == '0' {
			if index >= offset {
				result = result + string(number[index-offset])
			} else {
				result = result + "0"
			}
		}
		if char == '1' {
			result = result + "1"
		}
		if char == 'X' {
			result = result + "X"
		}
	}
	return result
}

func getAllMemoryAddresses(memoryAddress string) []string {
	allAdresses := []string{memoryAddress}
	for index, char := range memoryAddress {
		if char == 'X' {
			var nextAdresses []string
			for _, address := range allAdresses {
				next1 := address[0:index] + "1" + address[index+1:]
				next0 := address[0:index] + "0" + address[index+1:]
				nextAdresses = append(nextAdresses, next1, next0)
			}
			allAdresses = nextAdresses
		}
	}
	return allAdresses
}

func base10ToBase2String(base10 string) string {
	return strconv.FormatInt(int64(stringToInt(base10)), 2)
}

func stringBase2ToInt(number string) int64 {
	result, error := strconv.ParseInt(number, 2, 64)
	if error != nil {
		log.Panicln(error)
	}
	return result
}
