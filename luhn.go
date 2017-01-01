package goluhn

import (
	"regexp"
	"strconv"
	"strings"
)

//IsValid returns true if cardNumber is a valid credit card number or false otherwise.
func IsValid(cardNumStr string) bool {
	cardNumStr = reverse(strings.TrimSpace(cardNumStr))
	cardNum, err := strconv.Atoi(cardNumStr)
	if err != nil {
		return false
	}

	numLen := len(cardNumStr)
	sum := 0
	for i := 0; i <= numLen; i++ {
		digit := cardNum % 10
		cardNum /= 10
		if i%2 == 0 {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}

		sum += digit
	}

	sum += (sum * 9) % 10

	return sum%10 == 0
}

//CardType returns the type of credit card as a string i.e. visa, mastercard ...
//or an empty string. It calls the IsValid function.
func CardType(cardNumStr string) string {
	if !IsValid(cardNumStr) {
		return ""
	}

	for i, pattern := range cardTypes {
		matched, err := regexp.MatchString(pattern, cardNumStr)
		if err != nil {
			return ""
		}

		if matched {
			switch i {
			case 0:
				return "electron"
			case 1:
				return "amex"
			case 2:
				return "dankort"
			case 3:
				return "diners"
			case 4:
				return "discover"
			case 5:
				return "jcb"
			case 6:
				return "interpayment"
			case 7:
				return "maestro"
			case 8:
				return "mastercard"
			case 9:
				return "unionpay"
			case 10:
				return "visa"
			}
		}
	}

	return ""
}

//reverse reverses a string
func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

var cardTypes = []string{
	"^(4026|417500|4405|4508|4844|4913|4917)\\d+$", //visaelectron
	"^3[47][0-9]{5,}$",                             //Amex
	"^(5019)\\d+$",                                 //Dankort"
	"^3(?:0[0-5]|[68][0-9])[0-9]{11}$",             //Diners
	"^6(?:011|5[0-9]{2})[0-9]{3,}$",                //Discover
	"^(?:2131|1800|35[0-9]{3})[0-9]{3,}$",          //JCB
	"^(636)\\d+$",                                  //Interpayment
	"^(5018|5020|5038|5612|5893|6304|6759|6761|6762|6763|0604|6390)\\d+$", //Maestro
	"^5[1-5][0-9]{5,}$",                                                   //Mastercard
	"^(62|88)\\d+$",                                                       //UnionPay
	"^4[0-9]{12}(?:[0-9]{3})?$",                                           //Visa

}
