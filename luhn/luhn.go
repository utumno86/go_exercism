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
	nums := strings.Split(input, "")
	sum := int64(0)
	evenOddIndex := 1
	for i := len(nums) - 1; i > -1; i-- {
		currentDigit, err := strconv.ParseInt(nums[i], 10, 64)
		if err != nil {
			return -1
		}
		if (evenOddIndex % 2) == 0 {
			doubled := currentDigit * 2
			if doubled >= 10 {
				sum = sum + (doubled) - 9
			} else {
				sum = sum + doubled
			}
		} else {
			sum = sum + currentDigit
		}
		evenOddIndex++
	}
	return sum
}
