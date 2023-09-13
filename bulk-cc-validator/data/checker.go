package data

import (
	"fmt"
	"regexp"
	"strconv"
)

// Luhn algorithm
// https://en.wikipedia.org/wiki/Luhn_algorithm
func checkSum(cc int) int {
	var check_cc int
	for i := 0; cc > 0; i++ {
		enter_cc := cc % 10
		if i%2 == 0 {
			enter_cc = enter_cc * 2
			if enter_cc > 9 {
				enter_cc = enter_cc%10 + enter_cc/10
			}
		}
		check_cc += enter_cc
		cc = cc / 10
	}
	return check_cc % 10
}

// CalculateLuhn returns the check number
func calculateLuhn(cc int) int {
	checkNumber := checkSum(cc)
	if checkNumber == 0 {
		return 0
	}
	return 10 - checkNumber
}

// Verifies Credit Card number is valid or not based on Luhn algorithm
func luhnValidation(cc int) bool {
	return (cc%10+checkSum(cc/10))%10 == 0
}

func CCTest(cNums []int, vInfo map[vendor]int) {
	fmt.Println("here we GO")
	for _, cc := range cNums {
		cclength := len(strconv.FormatInt(int64(cc), 10))
		if cclength < 14 || cclength > 19 { // card length estimate based on google search =)
			fmt.Printf("%v possibly a fake card?\n", cc)
		} else {
			if !luhnValidation(cc) {
				fmt.Printf("Card Number %v is not allowed.\n", cc)
			}
			if calculateLuhn(cc/10) != cc%10 {
				fmt.Printf("%v's check number should be %v, but got %v", cc, cc%10, calculateLuhn(cc/10))
			}
			// Searches through company regular expressions and matches card type.
			ccString := strconv.FormatInt(int64(cc), 10)
			match := false
			for vendor := range vInfo {
				ccReg := regexp.MustCompile(vendor.ccRegex)
				match = ccReg.MatchString(ccString)
				if match == true {
					fmt.Printf("%v is a VALID %v number\n", cc, vendor.ccType)
					break
				}
			}
			if match == false {
				fmt.Printf("%v is not a valid credit card type.\n", cc)
			}
		}
	}
}
