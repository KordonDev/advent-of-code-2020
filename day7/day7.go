package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type BagMultiplyer struct {
	bagColor   string
	multiplier int
}

func main() {
	texts := readFile("./input.txt")

	bagsAreInBag := make(map[string][]string)
	numberOfBagsAreInBag := make(map[string][]BagMultiplyer)
	for _, text := range texts {
		cleanedText1 := strings.ReplaceAll(text, "bags", "bag")
		cleanedText2 := strings.ReplaceAll(cleanedText1, "contain", "")
		cleanedText3 := strings.ReplaceAll(cleanedText2, ",", "")
		cleanedText4 := strings.ReplaceAll(cleanedText3, ".", "")
		cleanedText5 := strings.ReplaceAll(cleanedText4, " ", "")
		parts := strings.Split(cleanedText5, "bag")
		if parts[1] != "noother" {
			for i := 1; i < len(parts); i++ {
				if len(parts[i]) > 0 {
					bagsAreInBag[parts[i][1:]] = append(bagsAreInBag[parts[i][1:]], parts[0])
					number, err := strconv.Atoi(string(parts[i][0]))
					if err != nil {
						log.Fatal(err)
					}
					bagMulti := BagMultiplyer{bagColor: parts[i][1:], multiplier: number}
					numberOfBagsAreInBag[parts[0]] = append(numberOfBagsAreInBag[parts[0]], bagMulti)
				}
			}
		}

	}

	checkedBags := make(map[string]string)
	bagsToCheck := []string{"shinygold"}
	countBagColors := -1
	for len(bagsToCheck) > 0 {
		color := bagsToCheck[0]
		bagsToCheck = bagsToCheck[1:]
		_, checked := checkedBags[color]
		if !checked {
			countBagColors = countBagColors + 1
			checkedBags[color] = "1"
			bagsToCheck = append(bagsToCheck, bagsAreInBag[color]...)
		}
	}

	fmt.Println("Number of possible bags for shiny gold", countBagColors)

	bagCount := 0
	var bagsToCheck2 []BagMultiplyer
	bagsToCheck2 = append(bagsToCheck2, BagMultiplyer{bagColor: "shinygold", multiplier: 1})

	for len(bagsToCheck2) > 0 {
		bag := bagsToCheck2[0]
		bagsToCheck2 = bagsToCheck2[1:]
		for _, nextBag := range numberOfBagsAreInBag[bag.bagColor] {
			currentBagsMultiplier := nextBag.multiplier * bag.multiplier
			bagCount = bagCount + currentBagsMultiplier
			bagsToCheck2 = append(bagsToCheck2, BagMultiplyer{bagColor: nextBag.bagColor, multiplier: currentBagsMultiplier})
		}
	}
	fmt.Println("Bag count loop", bagCount)

	fmt.Println("Bag count recursive", nextBag2("shinygold", &numberOfBagsAreInBag))
}

func nextBag(bagColor string, numberOfBagsAreInBag *map[string][]BagMultiplyer) int {
	nextBags, found := (*numberOfBagsAreInBag)[bagColor]
	// fmt.Println(nextBags, found)
	if !found {
		fmt.Println(1, bagColor)
		return 1
	}
	allBagsInBag := 1
	fmt.Println("add 1", bagColor)
	for _, bag := range nextBags {
		next := nextBag(bag.bagColor, numberOfBagsAreInBag)
		allBagsInBag = allBagsInBag + bag.multiplier*next
		fmt.Println(bag.bagColor, bag.multiplier, "*", next, "=", allBagsInBag)
	}
	fmt.Println("return ", bagColor, allBagsInBag)
	return allBagsInBag
}

func nextBag2(bagColor string, numberOfBagsAreInBag *map[string][]BagMultiplyer) int {
	nextBags, found := (*numberOfBagsAreInBag)[bagColor]
	if !found {
		return 1
	}
	allBagsInBag := 1
	for _, bag := range nextBags {
		allBagsInBag = allBagsInBag + bag.multiplier*nextBag2(bag.bagColor, numberOfBagsAreInBag)
	}
	return allBagsInBag
}
