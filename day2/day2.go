package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type PasswordData struct {
	min      int
	max      int
	pattern  string
	password string
}

func createPasswordData(text string) PasswordData {
	split := strings.Split(text, " ")

	min, err := strconv.Atoi(strings.Split(split[0], "-")[0])
	if err != nil {
		log.Fatal(err)
	}
	max, err := strconv.Atoi(strings.Split(split[0], "-")[1])
	if err != nil {
		log.Fatal(err)
	}
	return PasswordData{
		min:      min,
		max:      max,
		pattern:  strings.TrimSuffix(split[1], ":"),
		password: split[2],
	}
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	var texts []string
	for scanner.Scan() {
		texts = append(texts, scanner.Text())
	}
	file.Close()

	var correctPasswords = 0
	for _, text := range texts {
		pass := createPasswordData(text)
		numberOfOccurrence := strings.Count(pass.password, pass.pattern)
		if numberOfOccurrence >= pass.min && numberOfOccurrence <= pass.max {
			correctPasswords++
		}
	}
	fmt.Println("Number of correct passwords for counting", correctPasswords)

	correctPasswords = 0
	for _, text := range texts {
		pass := createPasswordData(text)
		firstCorrect := string(pass.password[pass.min-1]) == pass.pattern
		secondCorrect := string(pass.password[pass.max-1]) == pass.pattern
		if firstCorrect != secondCorrect {
			correctPasswords++
		}
	}
	fmt.Println("Number of correct passwords for position", correctPasswords)

}
