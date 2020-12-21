package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	lines := readFile("./input.txt")

	allergies := make(map[string][]string)
	var allIngredients []string
	for _, line := range lines {
		parts := strings.Split(line, " (contains ")
		allergie := strings.Split(strings.ReplaceAll(parts[1], ")", ""), ", ")
		ingreditens := strings.Split(parts[0], " ")
		allIngredients = append(allIngredients, ingreditens...)
		for _, all := range allergie {
			prevIngredients := allergies[all]
			allergies[all] = findIngredientsInBothLists(&prevIngredients, &ingreditens)
		}
	}

	changed := true
	for changed {
		changed = false
		for allergie, ingredients := range allergies {
			if len(ingredients) == 1 {
				for otherAllergie, otherIngredients := range allergies {
					if allergie != otherAllergie {
						var change bool
						allergies[otherAllergie], change = removeItems(&otherIngredients, &ingredients)
						changed = changed || change
					}
				}
			}
		}
	}

	for _, ingredients := range allergies {
		allIngredients, _ = removeItems(&allIngredients, &ingredients)
	}

	fmt.Println("Solution 1:", len(allIngredients))

	var allergiesList []string
	for allergie := range allergies {
		allergiesList = append(allergiesList, allergie)
	}
	sort.Strings(allergiesList)
	var dangerousIngredients []string
	for _, allergie := range allergiesList {
		dangerousIngredients = append(dangerousIngredients, allergies[allergie][0])
	}
	fmt.Println("Solution 2: ", strings.Join(dangerousIngredients, ","))

}

func removeItems(list1 *[]string, list2 *[]string) ([]string, bool) {
	var result []string
	change := false
	for _, item1 := range *list1 {
		contains := false
		for _, item2 := range *list2 {
			if item1 == item2 {
				contains = true
				break
			}
		}
		if !contains {
			result = append(result, item1)
		} else {
			change = true
		}
	}
	return result, change
}

func findIngredientsInBothLists(list1 *[]string, list2 *[]string) []string {
	var result []string
	if len(*list1) == 0 {
		return *list2
	}
	for _, item1 := range *list1 {
		for _, item2 := range *list2 {
			if item1 == item2 {
				result = append(result, item1)
				break
			}
		}
	}
	return result
}
