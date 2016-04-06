package rules

type Square struct {
	X, Y int
}

var (
	Files = []byte{'a','b','c','d','e','f','g','h'}
	Ranks = []byte{'1','2','3','4','5','6','7','8'}
)

// Takes a Square struct and convert it to common chess notation
func (s *Square) ToString() string {
	byteArray := [2]byte{Files[s.X - 1], Ranks[s.Y -1]}
	
	return string(byteArray[:])
}