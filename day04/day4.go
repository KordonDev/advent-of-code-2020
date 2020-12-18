package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	texts := readFile()

	var allPassport [](map[string]string)
	currentPassport := make(map[string]string)

	for _, line := range texts {
		if len(line) == 0 {
			allPassport = append(allPassport, currentPassport)
			currentPassport = make(map[string]string)
		} else {
			parts := strings.Split(line, " ")
			for _, data := range parts {
				pair := strings.Split(data, ":")
				currentPassport[pair[0]] = pair[1]
			}

		}
	}
	allPassport = append(allPassport, currentPassport)

	validPassports := 0
	for _, passport := range allPassport {

		valid := true

		// byr (Birth Year) - four digits; at least 1920 and at most 2002.
		valid = valid && isYearValid(passport["byr"], 1920, 2002)
		// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
		valid = valid && isYearValid(passport["iyr"], 2010, 2020)
		// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
		valid = valid && isYearValid(passport["eyr"], 2020, 2030)
		// hgt (Height) - a number followed by either cm or in:
		valid = valid && isHeightValid(passport["hgt"])
		valid = valid && isHairColorValid(passport["hcl"])
		valid = valid && isEyeColorValid(passport["ecl"])
		valid = valid && isPassportIDValid(passport["pid"])

		if valid {
			validPassports = validPassports + 1
		}
	}

	fmt.Println("Total number of valid passports", validPassports)
}

func isYearValid(year string, min int, max int) bool {
	yearNumber, err := strconv.Atoi(year)
	if err != nil || yearNumber < min || yearNumber > max {
		return false
	}
	return true
}

// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
func isEyeColorValid(eyeColor string) bool {
	eyeColors := map[string]string{"amb": "1", "blu": "1", "brn": "1", "gry": "1", "grn": "1", "hzl": "1", "oth": "1"}
	if len(eyeColors[eyeColor]) == 0 {
		return false
	}
	return true
}

func isHeightValid(height string) bool {
	textIndexCm := strings.Index(height, "cm")
	// If cm, the number must be at least 150 and at most 193.
	if textIndexCm > -1 {
		cm, err := strconv.Atoi(height[:textIndexCm])
		if err != nil || cm < 150 || cm > 193 {
			return false
		}
	}
	textIndexIn := strings.Index(height, "in")
	// If in, the number must be at least 59 and at most 76.
	if textIndexIn > -1 {
		inch, err := strconv.Atoi(height[:textIndexIn])
		if err != nil || inch < 59 || inch > 76 {
			return false
		}
	}
	if textIndexCm == -1 && textIndexIn == -1 {
		return false
	}
	return true
}

// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
func isHairColorValid(hairColor string) bool {
	if len(hairColor) == 0 {
		return false
	}
	hashIndex := strings.Index(hairColor, "#")
	match, _ := regexp.MatchString("^([a-f0-9])*$", hairColor[1:])
	if hashIndex != 0 || len(hairColor) != 7 || !match {
		return false
	}
	return true
}

// pid (Passport ID) - a nine-digit number, including leading zeroes.
func isPassportIDValid(passportID string) bool {
	match, _ := regexp.MatchString(`^[0-9]*$`, passportID)
	if len(passportID) != 9 || !match {
		return false
	}
	return true
}
