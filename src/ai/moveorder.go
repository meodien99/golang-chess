package ai

import (
	"sort"
	"rules"
)

// The following defines a type and functions such that the sort package can order move by thier score
type ByScore []*rules.Move

func (s ByScore) Len() int {
	return len(s)
}

func (s ByScore) Swap(i, j int){
	s[i], s[j] = s[j], s[i]
}

// Reverse to order moves from greatest to least
func (s ByScore) Less(i, j int) bool {
	return s[i].Score < s[j].Score
}

// Roughly orders moves in order of most likely to be good to least
// Examines all checks first, followed by captures, followed by good moves
// "Good moves" are sorted by their board evaluation after they are played
// If quiescence is set to true, then only checks and captures are returned
func orderedMove(b *rules.Board, quiescence bool) []*rules.Move {
	checks := make([]*rules.Move, 0)
	captures := make([]*rules.Move, 0)
	rests := make([]*rules.Move, 0)
	
	for _, move := range b.AllLegalMoves(){
		b.ForceMove(move)
		
		if b.IsCheck(b.Turn) {
			checks = append(checks, move)
		} else if move.Capture != 0 {
			captures = append(captures, move)
		} else if !quiescence {
			childScore := EvalBoard(b)* float64(b.Turn*-1)
			move.Score = childScore
			rests = append(rests, move)
		}
		b.UndoMove(move)
	}
	
	if !quiescence{
		sort.Sort(sort.Reverse(ByScore(rests)))
	}
	
	orderedMoves := make([]*rules.Move, len(checks) + len(captures) + len(rests))
	index := 0
	
	for _, l := range [][]*rules.Move{checks, captures, rests} {
		for _, m := range l {
			m.Score = 0
			orderedMoves[index] = m
			index ++
		}
	}
	
	return orderedMoves
}