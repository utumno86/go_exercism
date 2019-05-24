package hamming

import "errors"

//Distance function for calculating hamming distance
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("Incorrect values")
	}

	distance := 0

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			distance++
		}
	}

	return distance, nil
}
