package grains

//Square returns the number of grains on a particular square of the chess board
func Square(input int) (uint64, error) {
	if input == 1 {
		return uint64(1), nil
	}
	return uint64(2) ^ (uint64(input) - uint64(1)), nil
}

//Total returns the total number of grains on a 64 square chess board
func Total() uint64 {
	sum, error := Square(64)
	if error == nil {
		return sum
	}
	return 0
}
