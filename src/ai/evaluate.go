package ai

import "rules"

const (
	
)

var (
	VALUES = map[byte]int{'p':1, 'n':3, 'b':3, 'r': 5, 'q': 9}
)

// default math package uses float64
func absInt(i int) int {
	if i > 0 {
		return i
	}
	
	return i*-1
}

