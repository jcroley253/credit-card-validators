package main

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

type cardInfo struct {
	CardNumber, CVV, Month, Year int
	Issuer                       string
	Valid, Expired               bool
}

var (
	cardNumber int
	cvv        int
	month      int
	year       int
	issuer     string
	valid      bool
	expired    bool
)

func main() {
	fmt.Println("here we GO")

	userInput()
	VendorRegex()
}

// Captures card data the user inputs
func userInput() (int, int, int, int) {
	fmt.Println("Enter the credit card number")
	for {
		fmt.Scan(&cardNumber)
		cardNumString := strconv.FormatInt(int64(cardNumber), 10) // converts int to string so logic can make sure a number is entered
		_, err := strconv.Atoi(cardNumString)                     // performs check
		if err == nil {
			break
		} else {
			fmt.Printf("Supplied value %s is not a number. Please try again.\n", cardNumString)
		}
	}

	fmt.Println("Enter the CVV")
	for {
		fmt.Scan(&cvv)
		cvvString := strconv.FormatInt(int64(cvv), 10) // converts int to string so logic can make sure a number is entered
		_, err := strconv.Atoi(cvvString)              // performs check
		if err == nil {
			break
		} else {
			fmt.Printf("Supplied value %s is not a number. Please try again.\n", cvvString)
		}
	}

	fmt.Println("Enter 2 digit expriation Month (ex: 01, 03, 12)")
	for {
		fmt.Scan(&month)
		monthString := strconv.FormatInt(int64(month), 10) // converts int to string so logic can make sure a number is entered
		_, err := strconv.Atoi(monthString)                // performs check
		if err == nil {
			break
		} else {
			fmt.Printf("Supplied value %s is not a number. Please try again.\n", monthString)
		}
	}

	fmt.Println("Enter the 4 digit exiration year (ex: 1992, 2001, 1999)")
	for {
		fmt.Scan(&year)
		yearString := strconv.FormatInt(int64(year), 10) // converts int to string so logic can make sure a number is entered
		_, err := strconv.Atoi(yearString)               // performs check
		if err == nil {
			break
		} else {
			fmt.Printf("Supplied value %s is not a number. Please try again.\n", yearString)
		}
	}

	return cardNumber, cvv, month, year
}

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

func CCTest(vInfo map[vendor]int) {
	var cc = cardNumber
	ccLength := len(strconv.FormatInt(int64(cc), 10))
	cvvLength := len(strconv.FormatInt(int64(cvv), 10))
	if ccLength < 14 || ccLength > 19 { // card length estimate based on google search =)
		fmt.Printf("%v is possibly a fake card?\n", cc)
		valid = false
	} else {
		if !luhnValidation(cc) {
			fmt.Printf("Card Number %v is not allowed.\n", cc)
		}
		if calculateLuhn(cc/10) != cc%10 {
			fmt.Printf("%v's check number should be %v, but got %v\n", cc, cc%10, calculateLuhn(cc/10))
		}
		// Searches through company regular expressions and matches card type.
		ccString := strconv.FormatInt(int64(cc), 10)
		match := false
		for vendor := range vInfo {
			ccReg := regexp.MustCompile(vendor.ccRegex)
			match = ccReg.MatchString(ccString)
			if match == true {
				fmt.Printf("%v is a VALID %v number\n", cc, vendor.ccType)
				issuer = vendor.ccType
				valid = true
				break
			}
		}
		if match == false {
			fmt.Printf("%v is not a valid credit card type.\n", cc)
			issuer = "UNKOWN"
			valid = false
		}
	}
	if cvvLength < 3 || cvvLength > 4 {
		fmt.Printf("CVV number %v is not a standard CVV. Possibly a fake card?\n", cvv)
		valid = false
	}

	compareDate := time.Now().Year()
	if compareDate >= year {
		compareMonth := int(time.Now().Month())
		if compareMonth < month {
			expired = false
		} else {
			expired = true
		}
	} else {
		expired = true
	}

	cardInfo := cardInfo{
		CardNumber: cardNumber,
		CVV:        cvv,
		Month:      month,
		Year:       year,
		Issuer:     issuer,
		Valid:      valid,
		Expired:    expired,
	}

	fmt.Printf("%+v\n", cardInfo)

	return
}
