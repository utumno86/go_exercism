package luhn

import (
	"strconv"
	"strings"
)

// Valid checks whether a series of numbers are valid as credit card numbers
func Valid(input string) bool {
	strippedString := strings.Replace(input, " ", "", -1)
	sum := checksum(strippedString)
	if sum < 1 && len(strippedString) < 2 {
		return false
	}
	return (sum % 10) == 0
}

func checksum(input string) int64 {
	evenNums := len(input)%2 == 0
	nums := strings.Split(input, "")
	sum := int64(0)
	for i, num := range nums {
		currentDigit, err := strconv.ParseInt(num, 10, 64)
		if err != nil {
			return -1
		}
		evenOddIndex := i
		if !evenNums {
			evenOddIndex++
		}
		sum = sum + currentDigit
		if (evenOddIndex % 2) == 0 {
			sum = sum + currentDigit
			if currentDigit*2 > 9 {
				sum = sum - 9
			}
		}
	}
	return sum
}
