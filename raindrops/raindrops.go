package raindrops

import "strconv"

const testVersion = 3

//Convert integers into raindrops
func Convert(input int) string {
	response := ""
	if input%3 == 0 {
		response += "Pling"
	}
	if input%5 == 0 {
		response += "Plang"
	}
	if input%7 == 0 {
		response += "Plong"
	}
	if response == "" {
		response = strconv.Itoa(input)
	}
	return response
}
