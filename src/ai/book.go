package ai

//https://chessprogramming.wikispaces.com/Forsyth-Edwards+Notation
//http://www.365chess.com/

var Book = map[string][]string {
	// Initial position
	"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR b": []string{"pd2-d4", "pe2-e4"},

	// A10 English Opening
	"rnbqkbnr/pppppppp/8/8/2P5/8/PP1PPPPP/RNBQKBNR b": []string{"ng8-f6", "pe7-e5"},

	// A15 English Opening, Anglo-Indian Defense
	"rnbqkb1r/pppppppp/5n2/8/2P5/8/PP1PPPPP/RNBQKBNR w": []string{"ng1-f3", "pd2-d4", "nb1-c3"},

	// A16 English Opening
	"rnbqkb1r/pppppppp/5n2/8/2P5/2N5/PP1PPPPP/R1BQKBNR b": []string{"pe7-e6"},

	// A17 English Opening
	"rnbqkb1r/pppp1ppp/4pn2/8/2P5/2N5/PP1PPPPP/R1BQKBNR w":   []string{"pd2-d4", "pe2-e4", "ng1-f3"},
	"rnbqkb1r/pppp1ppp/4pn2/8/2P5/2N2N2/PP1PPPPP/R1BQKB1R b": []string{"pd7-d5"},

	// A18 English Opening, Mikenas-Carls Variation
	"rnbqkb1r/pppp1ppp/4pn2/8/2P1P3/2N5/PP1P1PPP/R1BQKBNR b":  []string{"pd7-d5"},
	"rnbqkb1r/ppp2ppp/4pn2/3p4/2P1P3/2N5/PP1P1PPP/R1BQKBNR w": []string{"pe4-e5"},
	"rnbqkb1r/ppp2ppp/4pn2/3pP3/2P5/2N5/PP1P1PPP/R1BQKBNR b":  []string{"pd5-d4"},
	"rnbqkb1r/ppp2ppp/4pn2/4P3/2Pp4/2N5/PP1P1PPP/R1BQKBNR w":  []string{"pe5-f6"},
	"rnbqkb1r/ppp2ppp/4pP2/8/2Pp4/2N5/PP1P1PPP/R1BQKBNR b":    []string{"pd4-c3"},
	"nbqkb1r/ppp2ppp/4pP2/8/2P5/2p5/PP1P1PPP/R1BQKBNR w":      []string{"pb2-c3"},

	// A20 English Opening
	"rnbqkbnr/pppp1ppp/8/4p3/2P5/8/PP1PPPPP/RNBQKBNR w": []string{"nb1-c3", "pg2-g3"},
}