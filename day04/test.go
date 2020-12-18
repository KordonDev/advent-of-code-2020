package main

import (
	"fmt"
)

func test() {
	fmt.Println("false", isHeightValid("2022"))
	fmt.Println("false", isHeightValid(""))
	fmt.Println("false", isHeightValid("90cm"))
	fmt.Println("false", isHeightValid("200cm"))
	fmt.Println("false", isHeightValid("58in"))
	fmt.Println("false", isHeightValid("77in"))
	fmt.Println("true", isHeightValid("60in"))
	fmt.Println("true", isHeightValid("150cm"))

	fmt.Println("false", isHairColorValid(""))
	fmt.Println("false", isHairColorValid("#123"))
	fmt.Println("false", isHairColorValid("121aa3#"))
	fmt.Println("false", isHairColorValid("#12g1aa3"))
	fmt.Println("true", isHairColorValid("#123456"))
	fmt.Println("true", isHairColorValid("#abacab"))
	fmt.Println("true", isHairColorValid("#aba123"))

	fmt.Println("false", isPassportIDValid(""))
	fmt.Println("false", isPassportIDValid("12554"))
	fmt.Println("false", isPassportIDValid("125ae1548"))
	fmt.Println("false", isPassportIDValid("0001235"))
	fmt.Println("true", isPassportIDValid("000123459"))
	fmt.Println("true", isPassportIDValid("459875615"))

	fmt.Println("false", isYearValid("", 1920, 2002))
	fmt.Println("false", isYearValid("123", 1920, 2002))
	fmt.Println("true", isYearValid("1920", 1920, 2002))
	fmt.Println("true", isYearValid("1990", 1920, 2002))
	fmt.Println("true", isYearValid("2002", 1920, 2002))
	fmt.Println("false", isYearValid("2022", 1920, 2002))

	fmt.Println("false", isEyeColorValid(""))
	fmt.Println("false", isEyeColorValid("aairen"))
	fmt.Println("true", isEyeColorValid("amb"))
	fmt.Println("true", isEyeColorValid("blu"))
	fmt.Println("true", isEyeColorValid("brn"))
	fmt.Println("true", isEyeColorValid("gry"))
	fmt.Println("true", isEyeColorValid("grn"))
	fmt.Println("true", isEyeColorValid("hzl"))
	fmt.Println("true", isEyeColorValid("oth"))
}
