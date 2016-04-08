package ai

//https://chessprogramming.wikispaces.com/Forsyth-Edwards+Notation
//http://www.365chess.com/

var Book = map[string][]string {
	"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR b": []string{"pd2-d4", "pe2-e4"},
	"rnbqkbnr/pppppppp/8/8/2P5/8/PP1PPPPP/RNBQKBNR b": []string{"ng8-f6", "pe7-e5"},
	"rnbqkb1r/pppppppp/5n2/8/2P5/8/PP1PPPPP/RNBQKBNR w": []string{"ng1-f3", "pd2-d4", "nb1-c3"},
	"rnbqkb1r/pppppppp/5n2/8/2P5/2N5/PP1PPPPP/R1BQKBNR b": []string{"pe7-e6"},
	"rnbqkb1r/pppp1ppp/4pn2/8/2P5/2N5/PP1PPPPP/R1BQKBNR w":   []string{"pd2-d4", "pe2-e4", "ng1-f3"},
	"rnbqkb1r/pppp1ppp/4pn2/8/2P5/2N2N2/PP1PPPPP/R1BQKB1R b": []string{"pd7-d5"},
	"rnbqkb1r/pppp1ppp/4pn2/8/2P1P3/2N5/PP1P1PPP/R1BQKBNR b":  []string{"pd7-d5"},
	"rnbqkb1r/ppp2ppp/4pn2/3p4/2P1P3/2N5/PP1P1PPP/R1BQKBNR w": []string{"pe4-e5"},
	"rnbqkb1r/ppp2ppp/4pn2/3pP3/2P5/2N5/PP1P1PPP/R1BQKBNR b":  []string{"pd5-d4"},
	"rnbqkb1r/ppp2ppp/4pn2/4P3/2Pp4/2N5/PP1P1PPP/R1BQKBNR w":  []string{"pe5-f6"},
	"rnbqkb1r/ppp2ppp/4pP2/8/2Pp4/2N5/PP1P1PPP/R1BQKBNR b":    []string{"pd4-c3"},
	"nbqkb1r/ppp2ppp/4pP2/8/2P5/2p5/PP1P1PPP/R1BQKBNR w":      []string{"pb2-c3"},
	"rnbqkbnr/pppp1ppp/8/4p3/2P5/8/PP1PPPPP/RNBQKBNR w": []string{"nb1-c3", "pg2-g3"},
	"rnbqkbnr/pppppppp/8/8/3P4/8/PPP1PPPP/RNBQKBNR b":   []string{"pd7-d5", "ng8-f6", "pe7-e6", "pf7-f5"},
	"rnbqkbnr/pppp1ppp/4p3/8/3P4/8/PPP1PPPP/RNBQKBNR w": []string{"pc2-c4", "pe2-e4"},
	"rnbqkbnr/pppp1ppp/4p3/8/2PP4/8/PP2PPPP/RNBQKBNR b": []string{"ng8-f6", "pd7-d5"},
	"rnbqkb1r/pppppppp/5n2/8/3P4/8/PPP1PPPP/RNBQKBNR w": []string{"pc2-c4", "ng1-f3"},
	"rnbqkb1r/pppppppp/5n2/8/3P4/5N2/PPP1PPPP/RNBQKB1R b":  []string{"pe7-e6"},
	"rnbqkb1r/pppp1ppp/4pn2/8/3P4/5N2/PPP1PPPP/RNBQKB1R w": []string{"pc2-c4"},
	"rnbqkb1r/pppppppp/5n2/8/2PP4/8/PP2PPPP/RNBQKBNR b": []string{"pe7-e6"},
	"rnbqkbnr/ppppp1pp/8/5p2/3P4/8/PPP1PPPP/RNBQKBNR w": []string{"pg2-g3", "pc2-c4"},
	"rnbqkbnr/ppppp1pp/8/5p2/3P4/6P1/PPP1PP1P/RNBQKBNR b":    []string{"ng8-f6"},
	"rnbqkb1r/ppppp1pp/5n2/5p2/3P4/6P1/PPP1PP1P/RNBQKBNR w":  []string{"bf1-g2"},
	"rnbqkb1r/ppppp1pp/5n2/5p2/3P4/6P1/PPP1PPBP/RNBQK1NR b":  []string{"pg7-g6", "pe7-e6"},
	"rnbqkb1r/ppppp2p/5np1/5p2/3P4/6P1/PPP1PPBP/RNBQK1NR w":  []string{"ng1-f3", "pc2-c4"},
	"rnbqkb1r/ppppp2p/5np1/5p2/3P4/5NP1/PPP1PPBP/RNBQK2R b":  []string{"bf8-g7"},
	"rnbqk2r/ppppp1bp/5np1/5p2/3P4/5NP1/PPP1PPBP/RNBQK2R w":  []string{"ke1-g1"},
	"rnbqk2r/ppppp1bp/5np1/5p2/3P4/5NP1/PPP1PPBP/RNBQ1RK1 b": []string{"ke8-g8"},
	"rnbqkb1r/ppppp2p/5np1/5p2/2PP4/6P1/PP2PPBP/RNBQK1NR b":   []string{"bf8-g7"},
	"rnbqk2r/ppppp1bp/5np1/5p2/2PP4/6P1/PP2PPBP/RNBQK1NR w":   []string{"nb1-c3"},
	"rnbqk2r/ppppp1bp/5np1/5p2/2PP4/2N3P1/PP2PPBP/R1BQK1NR b": []string{"ke8-g8"},
	"rnbqkb1r/pppp2pp/4pn2/5p2/3P4/6P1/PPP1PPBP/RNBQK1NR w":  []string{"ng1-f3", "pc2-c4"},
	"rnbqkb1r/pppp2pp/4pn2/5p2/3P4/5NP1/PPP1PPBP/RNBQK2R b":  []string{"bf8-e7", "pd7-d5"},
	"rnbqk2r/ppppb1pp/4pn2/5p2/3P4/5NP1/PPP1PPBP/RNBQK2R w":  []string{"ke1-g1"},
	"rnbqk2r/ppppb1pp/4pn2/5p2/3P4/5NP1/PPP1PPBP/RNBQ1RK1 b": []string{"ke8-g8"},
	"rnbqkb1r/ppp3pp/4pn2/3p1p2/3P4/5NP1/PPP1PPBP/RNBQK2R w":  []string{"ke1-g1"},
	"rnbqkb1r/ppp3pp/4pn2/3p1p2/3P4/5NP1/PPP1PPBP/RNBQ1RK1 b": []string{"bf8-d6"},
	"rnbqkbnr/ppppp1pp/8/5p2/2PP4/8/PP2PPPP/RNBQKBNR b":   []string{"ng8-f6"},
	"rnbqkb1r/ppppp1pp/5n2/5p2/2PP4/8/PP2PPPP/RNBQKBNR w": []string{"nb1-c3", "pg2-g3"},
	"rnbqkb1r/ppppp1pp/5n2/5p2/2PP4/2N5/PP2PPPP/R1BQKBNR b":   []string{"pg7-g6", "pe7-e6"},
	"rnbqkb1r/ppppp2p/5np1/5p2/2PP4/2N5/PP2PPPP/R1BQKBNR w":   []string{"pg2-g3", "ng1-f3"},
	"rnbqkb1r/ppppp2p/5np1/5p2/2PP4/2N3P1/PP2PP1P/R1BQKBNR b": []string{"bf8-g7"},
	"rnbqk2r/ppppp1bp/5np1/5p2/2PP4/2N3P1/PP2PP1P/R1BQKBNR w": []string{"bf1-g2"},
	"rnbqkb1r/ppppp2p/5np1/5p2/2PP4/2N2N2/PP2PPPP/R1BQKB1R b":  []string{"bf8-g7"},
	"rnbqk2r/ppppp1bp/5np1/5p2/2PP4/2N2N2/PP2PPPP/R1BQKB1R w":  []string{"pg2-g3"},
	"rnbqk2r/ppppp1bp/5np1/5p2/2PP4/2N2NP1/PP2PP1P/R1BQKB1R b": []string{"ke8-g8"},

	"rnbqkb1r/pppp2pp/4pn2/5p2/2PP4/2N5/PP2PPPP/R1BQKBNR w":   []string{"ng1-f3", "pg2-g3"},
	"rnbqkb1r/pppp2pp/4pn2/5p2/2PP4/2N2N2/PP2PPPP/R1BQKB1R b": []string{"bf8-b4", "pd7-d5", "bf8-e7"},

	"rnbqkb1r/ppppp1pp/5n2/5p2/2PP4/6P1/PP2PP1P/RNBQKBNR b": []string{"pg7-g6"},
	"rnbqkb1r/ppppp2p/5np1/5p2/2PP4/6P1/PP2PP1P/RNBQKBNR w": []string{"bf1-g2"},

	"rnbqkb1r/pppp2pp/4pn2/5p2/2PP4/6P1/PP2PPBP/RNBQK1NR b":  []string{"bf8-e7", "pd7-d5", "pc7-c6"},
	"rnbqkb1r/ppp3pp/4pn2/3p1p2/2PP4/6P1/PP2PPBP/RNBQK1NR w": []string{"ng1-f3"},
	"rnbqkb1r/ppp3pp/4pn2/3p1p2/2PP4/5NP1/PP2PPBP/RNBQK2R b": []string{"pc7-c6"},
	
	"rnbqk2r/ppppb1pp/4pn2/5p2/2PP4/6P1/PP2PPBP/RNBQK1NR w": []string{"ng1-f3"},
	"rnbqk2r/ppppb1pp/4pn2/5p2/2PP4/5NP1/PP2PPBP/RNBQK2R b": []string{"ke8-g8"},

	"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b": []string{"pe7-e5", "pc7-c5", "pe7-e6", "pc7-c6"},

	"rnbqkbnr/pp1ppppp/2p5/8/4P3/8/PPPP1PPP/RNBQKBNR w": []string{"pd2-d4"},

	"rnbqkbnr/pp1ppppp/2p5/8/3PP3/8/PPP2PPP/RNBQKBNR b":     []string{"pd7-d5"},
	"rnbqkbnr/pp2pppp/2p5/3p4/3PP3/8/PPP2PPP/RNBQKBNR w":    []string{"nb1-c3", "pe4-d5", "pe4-e5"},
	"rnbqkbnr/pp2pppp/2p5/3pP3/3P4/8/PPP2PPP/RNBQKBNR b":    []string{"bc8-f5"},
	"rn1qkbnr/pp2pppp/2p5/3pPb2/3P4/8/PPP2PPP/RNBQKBNR w":   []string{"nb1-c3", "ng1-f3"},
	"rn1qkbnr/pp2pppp/2p5/3pPb2/3P4/2N5/PPP2PPP/R1BQKBNR b": []string{"pe7-e6"},
	"rn1qkbnr/pp2pppp/2p5/3pPb2/3P4/5N2/PPP2PPP/RNBQKB1R b": []string{"pe7-e6"},

	// B13 Caro-Kann, Exchange Variation
	"rnbqkbnr/pp2pppp/2p5/3P4/3P4/8/PPP2PPP/RNBQKBNR b":   []string{"pc6-d5"},
	"rnbqkbnr/pp2pppp/8/3p4/3P4/8/PPP2PPP/RNBQKBNR w":     []string{"pc2-c4"},
	"rnbqkbnr/pp2pppp/8/3p4/2PP4/8/PP3PPP/RNBQKBNR b":     []string{"ng8-f6"},
	"rnbqkb1r/pp2pppp/5n2/3p4/2PP4/8/PP3PPP/RNBQKBNR w":   []string{"nb1-c3"},
	"rnbqkb1r/pp2pppp/5n2/3p4/2PP4/2N5/PP3PPP/R1BQKBNR b": []string{"pe7-e6", "nb8-c6"},

	// B15 Caro-Kann Defense
	"rnbqkbnr/pp2pppp/2p5/3p4/3PP3/2N5/PPP2PPP/R1BQKBNR b": []string{"pd5-e4"},
	"rnbqkbnr/pp2pppp/2p5/8/3Pp3/2N5/PPP2PPP/R1BQKBNR w":   []string{"nc3-e4"},
	"rnbqkbnr/pp2pppp/2p5/8/3PN3/8/PPP2PPP/R1BQKBNR b":     []string{"bc8-f5"},

	// B18 Caro-Kann, Classical Variation
	"rn1qkbnr/pp2pppp/2p5/5b2/3PN3/8/PPP2PPP/R1BQKBNR w":  []string{"ne4-g3"},
	"rn1qkbnr/pp2pppp/2p5/5b2/3P4/6N1/PPP2PPP/R1BQKBNR b": []string{"bf5-g6"},

	// B20 Sicilian Defense
	"rnbqkbnr/pp1ppppp/8/2p5/4P3/8/PPPP1PPP/RNBQKBNR w": []string{"ng1-f3", "nb1-c3", "pc2-c3"},

	// B22 Sicilian, Alapin Variation
	"rnbqkbnr/pp1ppppp/8/2p5/4P3/2P5/PP1P1PPP/RNBQKBNR b":   []string{"ng8-f6", "pd7-d5"},
	"rnbqkb1r/pp1ppppp/5n2/2p5/4P3/2P5/PP1P1PPP/RNBQKBNR w": []string{"pe4-e5"},
	"rnbqkb1r/pp1ppppp/5n2/2p1P3/8/2P5/PP1P1PPP/RNBQKBNR b": []string{"nf6-d5"},
	"rnbqkb1r/pp1ppppp/8/2pnP3/8/2P5/PP1P1PPP/RNBQKBNR w":   []string{"pd2-d4", "ng1-f3"},
	"rnbqkb1r/pp1ppppp/8/2pnP3/3P4/2P5/PP3PPP/RNBQKBNR b":   []string{"pc5-d4"},
	"rnbqkb1r/pp1ppppp/8/3nP3/3p4/2P5/PP3PPP/RNBQKBNR w":    []string{"ng1-f3"},
	"rnbqkb1r/pp1ppppp/8/3nP3/3p4/2P2N2/PP3PPP/RNBQKB1R b":  []string{"nb8-c6", "pe7-e6"},

	"rnbqkb1r/pp1ppppp/8/2pnP3/8/2P2N2/PP1P1PPP/RNBQKB1R b": []string{"nb8-c6", "pe7-e6"},

	"rnbqkbnr/pp2pppp/8/2pp4/4P3/2P5/PP1P1PPP/RNBQKBNR w":   []string{"pe4-d5"},
	"rnbqkbnr/pp2pppp/8/2pP4/8/2P5/PP1P1PPP/RNBQKBNR b":     []string{"qd8-d5"},
	"rnb1kbnr/pp2pppp/8/2pq4/8/2P5/PP1P1PPP/RNBQKBNR w":     []string{"pd2-d4"},
	"rnb1kbnr/pp2pppp/8/2pq4/3P4/2P5/PP3PPP/RNBQKBNR b":     []string{"ng8-f6"},
	"rnb1kb1r/pp2pppp/5n2/2pq4/3P4/2P5/PP3PPP/RNBQKBNR w":   []string{"ng1-f3"},
	"rnb1kb1r/pp2pppp/5n2/2pq4/3P4/2P2N2/PP3PPP/RNBQKB1R b": []string{"bc8-g4", "pe7-e6"},

	// B23 Sicilian, Closed
	"rnbqkbnr/pp1ppppp/8/2p5/4P3/2N5/PPPP1PPP/R1BQKBNR b":       []string{"nb8-c6"},
	"r1bqkbnr/pp1ppppp/2n5/2p5/4P3/2N5/PPPP1PPP/R1BQKBNR w":     []string{"pf2-f4", "pg2-g3"},
	"r1bqkbnr/pp1ppppp/2n5/2p5/4PP2/2N5/PPPP2PP/R1BQKBNR b":     []string{"pg7-g6"},
	"r1bqkbnr/pp1ppp1p/2n3p1/2p5/4PP2/2N5/PPPP2PP/R1BQKBNR w":   []string{"ng1-f3"},
	"r1bqkbnr/pp1ppp1p/2n3p1/2p5/4PP2/2N2N2/PPPP2PP/R1BQKB1R b": []string{"bf8-g7"},

	// B24 Sicilian, Closed
	"r1bqkbnr/pp1ppppp/2n5/2p5/4P3/2N3P1/PPPP1P1P/R1BQKBNR b":   []string{"pg7-g6"},
	"r1bqkbnr/pp1ppp1p/2n3p1/2p5/4P3/2N3P1/PPPP1P1P/R1BQKBNR w": []string{"bf1-g2"},
	"r1bqkbnr/pp1ppp1p/2n3p1/2p5/4P3/2N3P1/PPPP1PBP/R1BQK1NR b": []string{"bf8-g7"},
	"r1bqk1nr/pp1pppbp/2n3p1/2p5/4P3/2N3P1/PPPP1PBP/R1BQK1NR w": []string{"pd2-d3"},
	"r1bqk1nr/pp1pppbp/2n3p1/2p5/4P3/2NP2P1/PPP2PBP/R1BQK1NR b": []string{"pd7-d6"},

	// B27 Sicilian Defense
	"rnbqkbnr/pp1ppppp/8/2p5/4P3/5N2/PPPP1PPP/RNBQKB1R b": []string{"pd7-d6", "pe7-e6"},

	// B40 Sicilian Defense
	"rnbqkbnr/pp1p1ppp/4p3/2p5/4P3/5N2/PPPP1PPP/RNBQKB1R w": []string{"pd2-d4"},
	"rnbqkbnr/pp1p1ppp/4p3/2p5/3PP3/5N2/PPP2PPP/RNBQKB1R b": []string{"pc5-d4"},
	"rnbqkbnr/pp1p1ppp/4p3/8/3pP3/5N2/PPP2PPP/RNBQKB1R w":   []string{"nf3-d4"},
	"rnbqkbnr/pp1p1ppp/4p3/8/3NP3/8/PPP2PPP/RNBQKB1R b":     []string{"pa7-a6", "nb8-c6", "ng8-f6"},

	// B50 Sicilian
	"rnbqkbnr/pp2pppp/3p4/2p5/4P3/5N2/PPPP1PPP/RNBQKB1R w": []string{"pd2-d4"},
	"rnbqkbnr/pp2pppp/3p4/2p5/3PP3/5N2/PPP2PPP/RNBQKB1R b": []string{"pc5-d4"},
	"rnbqkbnr/pp2pppp/3p4/8/3pP3/5N2/PPP2PPP/RNBQKB1R w":   []string{"nf3-d4"},

	// B54 Sicilian
	"rnbqkbnr/pp2pppp/3p4/8/3NP3/8/PPP2PPP/RNBQKB1R b":   []string{"ng8-f6"},
	"rnbqkb1r/pp2pppp/3p1n2/8/3NP3/8/PPP2PPP/RNBQKB1R w": []string{"nb1-c3"},

	// C00 French Defense
	"rnbqkbnr/pppp1ppp/4p3/8/4P3/8/PPPP1PPP/RNBQKBNR w":  []string{"pd2-d4"},
	"rnbqkbnr/pppp1ppp/4p3/8/3PP3/8/PPP2PPP/RNBQKBNR b":  []string{"pd7-d5"},
	"rnbqkbnr/ppp2ppp/4p3/3p4/3PP3/8/PPP2PPP/RNBQKBNR w": []string{"pe4-e5", "nb1-c3"},

	// C02 French, Advance Variation
	"rnbqkbnr/ppp2ppp/4p3/3pP3/3P4/8/PPP2PPP/RNBQKBNR b":      []string{"pc7-c5"},
	"rnbqkbnr/pp3ppp/4p3/2ppP3/3P4/8/PPP2PPP/RNBQKBNR w":      []string{"pc2-c3"},
	"rnbqkbnr/pp3ppp/4p3/2ppP3/3P4/2P5/PP3PPP/RNBQKBNR b":     []string{"nb8-c6"},
	"r1bqkbnr/pp3ppp/2n1p3/2ppP3/3P4/2P5/PP3PPP/RNBQKBNR w":   []string{"ng1-f3"},
	"r1bqkbnr/pp3ppp/2n1p3/2ppP3/3P4/2P2N2/PP3PPP/RNBQKB1R b": []string{"bc8-d7", "qd8-b6"},

	// C10 French, Paulsen Variation
	"rnbqkbnr/ppp2ppp/4p3/3p4/3PP3/2N5/PPP2PPP/R1BQKBNR b": []string{"bf8-b4", "ng8-f6"},

	// C11 French Defense
	"rnbqkb1r/ppp2ppp/4pn2/3p4/3PP3/2N5/PPP2PPP/R1BQKBNR w":   []string{"bc1-g5", "pe4-e5"},
	"rnbqkb1r/ppp2ppp/4pn2/3p2B1/3PP3/2N5/PPP2PPP/R2QKBNR b":  []string{"pd5-e4", "bf8-e7"},
	"rnbqkb1r/ppp2ppp/4pn2/3pP3/3P4/2N5/PPP2PPP/R1BQKBNR b":   []string{"nf6-d7"},
	"rnbqkb1r/pppn1ppp/4p3/3pP3/3P4/2N5/PPP2PPP/R1BQKBNR w":   []string{"pf2-f4"},
	"rnbqkb1r/pppn1ppp/4p3/3pP3/3P1P2/2N5/PPP3PP/R1BQKBNR b":  []string{"pc7-c5"},
	"rnbqkb1r/pp1n1ppp/4p3/2ppP3/3P1P2/2N5/PPP3PP/R1BQKBNR w": []string{"ng1-f3"},

	// C15 French, Winawer, Nimzovich Variation
	"rnbqk1nr/ppp2ppp/4p3/3p4/1b1PP3/2N5/PPP2PPP/R1BQKBNR w": []string{"pe4-e5"},

	// C16 French, Winawer, Advance Variation
	"rnbqk1nr/ppp2ppp/4p3/3pP3/1b1P4/2N5/PPP2PPP/R1BQKBNR b": []string{"pc7-c5"},

	// C20 King's Pawn Game
	"rnbqkbnr/pppp1ppp/8/4p3/4P3/8/PPPP1PPP/RNBQKBNR w": []string{"ng1-f3"},

	// C40 King's Knight Opening
	"rnbqkbnr/pppp1ppp/8/4p3/4P3/5N2/PPPP1PPP/RNBQKB1R b": []string{"nb8-c6", "ng8-f6"},

	// C42 Petrov's Defense
	"rnbqkb1r/pppp1ppp/5n2/4p3/4P3/5N2/PPPP1PPP/RNBQKB1R w": []string{"nf3-e5", "pd2-d4", "nb1-c3"},
	"rnbqkb1r/pppp1ppp/5n2/4N3/4P3/8/PPPP1PPP/RNBQKB1R b":   []string{"pd7-d6"},
	"rnbqkb1r/ppp2ppp/3p1n2/4N3/4P3/8/PPPP1PPP/RNBQKB1R w":  []string{"ne5-f3"},
	"rnbqkb1r/ppp2ppp/3p1n2/8/4P3/5N2/PPPP1PPP/RNBQKB1R b":  []string{"nf6-e4"},
	"rnbqkb1r/ppp2ppp/3p4/8/4n3/5N2/PPPP1PPP/RNBQKB1R w":    []string{"pd2-d4"},
	"rnbqkb1r/ppp2ppp/3p4/8/3Pn3/5N2/PPP2PPP/RNBQKB1R b":    []string{"pd6-d5"},

	"rnbqkb1r/pppp1ppp/5n2/4p3/4P3/2N2N2/PPPP1PPP/R1BQKB1R b": []string{"nb8-c6"},

	// C43 Petrov's Defense, Modern Attack
	"rnbqkb1r/pppp1ppp/5n2/4p3/3PP3/5N2/PPP2PPP/RNBQKB1R b": []string{"nf6-e4"},
	"rnbqkb1r/pppp1ppp/8/4p3/3Pn3/5N2/PPP2PPP/RNBQKB1R w":   []string{"bf1-d3"},
	"rnbqkb1r/pppp1ppp/8/4p3/3Pn3/3B1N2/PPP2PPP/RNBQK2R b":  []string{"pd7-d5"},
	"rnbqkb1r/ppp2ppp/8/3pp3/3Pn3/3B1N2/PPP2PPP/RNBQK2R w":  []string{"nf3-e5"},
	"rnbqkb1r/ppp2ppp/8/3pN3/3Pn3/3B4/PPP2PPP/RNBQK2R b":    []string{"nb8-d7"},
	"r1bqkb1r/pppn1ppp/8/3pN3/3Pn3/3B4/PPP2PPP/RNBQK2R w":   []string{"ne5-d7"},
	"r1bqkb1r/pppN1ppp/8/3p4/3Pn3/3B4/PPP2PPP/RNBQK2R b":    []string{"bc8-d7"},

	// C44 King's Pawn Game
	"r1bqkbnr/pppp1ppp/2n5/4p3/4P3/5N2/PPPP1PPP/RNBQKB1R w": []string{"bf1-c5", "bf1-b5"},

	// C46 Four Knight's Game
	"r1bqkb1r/pppp1ppp/2n2n2/4p3/4P3/2N2N2/PPPP1PPP/R1BQKB1R w": []string{"bf1-b5", "pd2-d4"},

	// C48 Four Knights, Spanish Variation
	"r1bqkb1r/pppp1ppp/2n2n2/1B2p3/4P3/2N2N2/PPPP1PPP/R1BQK2R b": []string{"bf8-b4", "nc6-d4"},
	"r1bqkb1r/pppp1ppp/5n2/1B2p3/3nP3/2N2N2/PPPP1PPP/R1BQK2R w":  []string{"bb5-a4", "nf3-d4"},
	"r1bqkb1r/pppp1ppp/5n2/4p3/B2nP3/2N2N2/PPPP1PPP/R1BQK2R b":   []string{"bf8-c5"},

	"r1bqkb1r/pppp1ppp/5n2/1B2p3/3NP3/2N5/PPPP1PPP/R1BQK2R b": []string{"pe5-d4"},
	"r1bqkb1r/pppp1ppp/5n2/1B6/3pP3/2N5/PPPP1PPP/R1BQK2R w":   []string{"pe4-e5"},

	// C49 Four Knights, Double Ruy Lopez
	"r1bqk2r/pppp1ppp/2n2n2/1B2p3/1b2P3/2N2N2/PPPP1PPP/R1BQK2R w": []string{"ke1-g1"},
	"1bqk2r/pppp1ppp/2n2n2/1B2p3/1b2P3/2N2N2/PPPP1PPP/R1BQ1RK1 b": []string{"ke8-g8"},

	// C50 King's Pawn Game
	"r1bqkbnr/pppp1ppp/2n5/4p3/2B1P3/5N2/PPPP1PPP/RNBQK2R b": []string{"ng8-f6", "bf8-c5"},

	// C60 Ruy Lopez
	"r1bqkbnr/pppp1ppp/2n5/1B2p3/4P3/5N2/PPPP1PPP/RNBQK2R b":  []string{"pa7-a6", "ng8-f6"},
	"r1bqkbnr/1ppp1ppp/p1n5/1B2p3/4P3/5N2/PPPP1PPP/RNBQK2R w": []string{"bb5-c6", "bb5-a4"},
	
	"r1bqkb1r/pppp1ppp/2n2n2/1B2p3/4P3/5N2/PPPP1PPP/RNBQK2R w": []string{"ke1-g1"},
	
	"r1bqkbnr/1ppp1ppp/p1B5/4p3/4P3/5N2/PPPP1PPP/RNBQK2R b": []string{"pd7-c6"},
	"r1bqkbnr/1pp2ppp/p1p5/4p3/4P3/5N2/PPPP1PPP/RNBQK2R w":  []string{"ke1-g1"},

	
	"r1bqkbnr/1pp2ppp/p1p5/4p3/4P3/5N2/PPPP1PPP/RNBQ1RK1 b": []string{"pf7-f6", "bc8-g4"},
	// D00 Queen's Pawn Game
	"rnbqkbnr/ppp1pppp/8/3p4/3P4/8/PPP1PPPP/RNBQKBNR w": []string{"pc2-c4"},

	"rnbqkbnr/ppp1pppp/8/3p4/2PP4/8/PP2PPPP/RNBQKBNR b": []string{"pe7-e6", "pc7-c6"},

	"rnbqkbnr/pp2pppp/2p5/3p4/2PP4/8/PP2PPPP/RNBQKBNR w":     []string{"nb1-c3", "ng1-f3"},
	"rnbqkbnr/pp2pppp/2p5/3p4/2PP4/2N5/PP2PPPP/R1BQKBNR b":   []string{"ng8-f6"},
	"rnbqkb1r/pp2pppp/2p2n2/3p4/2PP4/2N5/PP2PPPP/R1BQKBNR w": []string{"ng1-f3"},

	"rnbqkbnr/pp2pppp/2p5/3p4/2PP4/5N2/PP2PPPP/RNBQKB1R b":   []string{"ng8-f6"},
	"rnbqkb1r/pp2pppp/2p2n2/3p4/2PP4/5N2/PP2PPPP/RNBQKB1R w": []string{"nb1-c3"},

	"rnbqkb1r/pp2pppp/2p2n2/3p4/2PP4/2N2N2/PP2PPPP/R1BQKB1R b": []string{"pe7-e6", "pd5-c4"},
	"rnbqkb1r/pp2pppp/2p2n2/8/2pP4/2N2N2/PP2PPPP/R1BQKB1R w":   []string{"pa2-a4"},

	"rnbqkb1r/pp2pppp/2p2n2/8/P1pP4/2N2N2/1P2PPPP/R1BQKB1R b": []string{"bc8-f5"},

	"rn1qkb1r/pp2pppp/2p2n2/5b2/P1pP4/2N2N2/1P2PPPP/R1BQKB1R w": []string{"pe2-e3", "nf3-e5"},

	"rnbqkbnr/ppp1pppp/8/8/2pP4/8/PP2PPPP/RNBQKBNR w":    []string{"ng1-f3", "pe2-e3"},
	"rnbqkbnr/ppp1pppp/8/8/2pP4/4P3/PP3PPP/RNBQKBNR b":   []string{"ng8-f6"},
	"rnbqkb1r/ppp1pppp/5n2/8/2pP4/4P3/PP3PPP/RNBQKBNR w": []string{"bf1-c4"},
	"rnbqkb1r/ppp1pppp/5n2/8/2BP4/4P3/PP3PPP/RNBQK1NR b": []string{"pe7-e6"},
	"rnbqkb1r/ppp2ppp/4pn2/8/2BP4/4P3/PP3PPP/RNBQK1NR w": []string{"ng1-f3"},
	"rnbqkbnr/ppp1pppp/8/8/2pP4/5N2/PP2PPPP/RNBQKB1R b": []string{"ng8-f6"},
	"rnbqkb1r/ppp1pppp/5n2/8/2pP4/5N2/PP2PPPP/RNBQKB1R w": []string{"pe2-e3", "nb1-c3"},
	"rnbqkb1r/ppp1pppp/5n2/8/2pP4/2N2N2/PP2PPPP/R1BQKB1R b": []string{"pc7-c6"},
	"rnbqkb1r/ppp1pppp/5n2/8/2pP4/4PN2/PP3PPP/RNBQKB1R b": []string{"pe7-e6"},
	"rnbqkb1r/ppp2ppp/4pn2/8/2pP4/4PN2/PP3PPP/RNBQKB1R w": []string{"bf1-c4"},
	"rnbqkb1r/ppp2ppp/4pn2/8/2BP4/4PN2/PP3PPP/RNBQK2R b":  []string{"pc7-c5"},
	"rnbqkbnr/ppp2ppp/4p3/3p4/2PP4/8/PP2PPPP/RNBQKBNR w":   []string{"nb1-c3", "ng1-f3"},
	"rnbqkbnr/ppp2ppp/4p3/3p4/2PP4/5N2/PP2PPPP/RNBQKB1R b": []string{"ng8-f6", "pc7-c6"},
	"rnbqkbnr/ppp2ppp/4p3/3p4/2PP4/2N5/PP2PPPP/R1BQKBNR b":    []string{"ng8-f6", "pc7-c6"},
	"rnbqkbnr/pp3ppp/2p1p3/3p4/2PP4/2N5/PP2PPPP/R1BQKBNR w":   []string{"ng1-f3", "pe2-e3"},
	"rnbqkbnr/pp3ppp/2p1p3/3p4/2PP4/2N2N2/PP2PPPP/R1BQKB1R b": []string{"ng8-f6"},
	"rnbqkbnr/pp3ppp/2p1p3/3p4/2PP4/2N1P3/PP3PPP/R1BQKBNR b":  []string{"ng8-f6"},
	"rnbqkb1r/ppp2ppp/4pn2/3p4/2PP4/2N5/PP2PPPP/R1BQKBNR w": []string{"ng1-f3", "pc4-d5"},
	"rnbqkb1r/ppp2ppp/4pn2/3P4/3P4/2N5/PP2PPPP/R1BQKBNR b":  []string{"pe6-d5"},
	"rnbqkb1r/ppp2ppp/5n2/3p4/3P4/2N5/PP2PPPP/R1BQKBNR w":   []string{"bc1-g5"},
	"rnbqkb1r/ppp2ppp/5n2/3p2B1/3P4/2N5/PP2PPPP/R2QKBNR b":  []string{"pc7-c6", "bf8-e7"},
	"rnbqkb1r/ppp2ppp/4pn2/3p4/2PP4/2N2N2/PP2PPPP/R1BQKB1R b":  []string{"Bf8-e7", "pc7-c6"},
	"rnbqk2r/ppp1bppp/4pn2/3p4/2PP4/2N2N2/PP2PPPP/R1BQKB1R w":  []string{"Bc1-g5"},
	"rnbqk2r/ppp1bppp/4pn2/3p2B1/2PP4/2N2N2/PP2PPPP/R2QKB1R b": []string{"ke8-g8"},
	"rnbqkb1r/pp3ppp/2p1pn2/3p4/2PP4/2N2N2/PP2PPPP/R1BQKB1R w":    []string{"bc1-g5", "pe2-e3"},
	"rnbqkb1r/pp3ppp/2p1pn2/3p2B1/2PP4/2N2N2/PP2PPPP/R2QKB1R b":   []string{"ph7-h6", "nb8-d7"},
	"rnbqkb1r/pp3pp1/2p1pn1p/3p2B1/2PP4/2N2N2/PP2PPPP/R2QKB1R w":  []string{"bg5-f6", "bg5-h4"},
	"r1bqkb1r/pp1n1ppp/2p1pn2/3p2B1/2PP4/2N2N2/PP2PPPP/R2QKB1R w": []string{"pe2-e3"},
	"rnbqkb1r/pp3ppp/2p1pn2/3p4/2PP4/2N1PN2/PP3PPP/R1BQKB1R b": []string{"nb8-d7"},
	"rnbqkb1r/pppp1ppp/4pn2/8/2PP4/8/PP2PPPP/RNBQKBNR w":     []string{"nb1-c3", "ng1-f3"},
	"rnbqkb1r/pppp1ppp/4pn2/8/2PP4/2N5/PP2PPPP/R1BQKBNR b":   []string{"bf8-b4", "pc7-c5", "pd7-d5"},
	"rnbqkb1r/pp1p1ppp/4pn2/2p5/2PP4/2N5/PP2PPPP/R1BQKBNR w": []string{"pd4-d5"},
	"rnbqkb1r/pp1p1ppp/4pn2/2pP4/2P5/2N5/PP2PPPP/R1BQKBNR b": []string{"pe6-d5"},
	"rnbqkb1r/pp1p1ppp/5n2/2pp4/2P5/2N5/PP2PPPP/R1BQKBNR w":  []string{"pc4-d5"},
	"rnbqkb1r/pp1p1ppp/5n2/2pP4/8/2N5/PP2PPPP/R1BQKBNR b":    []string{"pd7-d6"},
	"rnbqkb1r/pppp1ppp/4pn2/8/2PP4/5N2/PP2PPPP/RNBQKB1R b":   []string{"pc7-c5"},
	"rnbqkb1r/ppp2ppp/4pn2/3p4/2PP4/5N2/PP2PPPP/RNBQKB1R w":  []string{"nb1-c3"},
	"rnbqkb1r/pp1p1ppp/4pn2/2p5/2PP4/5N2/PP2PPPP/RNBQKB1R w": []string{"pd4-d5"},
	"rnbqkb1r/pp1p1ppp/4pn2/2pP4/2P5/5N2/PP2PPPP/RNBQKB1R b": []string{"pe6-d5"},
	"rnbqkb1r/pp1p1ppp/5n2/2pp4/2P5/5N2/PP2PPPP/RNBQKB1R w":  []string{"pc4-d5"},
	"rnbqkb1r/pp1p1ppp/5n2/2pP4/8/5N2/PP2PPPP/RNBQKB1R b":    []string{"pd7-d6"},
	"rnbqk2r/pppp1ppp/4pn2/8/1bPP4/2N5/PP2PPPP/R1BQKBNR w": []string{"ng1-f3", "qd1-c2", "pe2-e3"},
	"rnbqk2r/pppp1ppp/4pn2/8/1bPP4/2N2N2/PP2PPPP/R1BQKB1R b": []string{"pc7-c5", "pd7-d5"},
	"rnbqk2r/pppp1ppp/4pn2/8/1bPP4/2N5/PPQ1PPPP/R1B1KBNR b":   []string{"ke8-g8", "pc7-c5"},
	"rnbq1rk1/pppp1ppp/4pn2/8/1bPP4/2N5/PPQ1PPPP/R1B1KBNR w":  []string{"pa2-a3"},
	"rnbq1rk1/pppp1ppp/4pn2/8/1bPP4/P1N5/1PQ1PPPP/R1B1KBNR b": []string{"bb4-c3"},
	"rnbqk2r/pp1p1ppp/4pn2/2p5/1bPP4/2N5/PPQ1PPPP/R1B1KBNR w": []string{"pd4-c5"},
	"rnbqk2r/pp1p1ppp/4pn2/2P5/1bP5/2N5/PPQ1PPPP/R1B1KBNR b":  []string{"ke8-g8", "bb4-c5"},
	"rnbqk2r/pppp1ppp/4pn2/8/1bPP4/2N1P3/PP3PPP/R1BQKBNR b": []string{"ke8-g8", "pc7-c5"},
	"rnbq1rk1/pppp1ppp/4pn2/8/1bPP4/2N1P3/PP3PPP/R1BQKBNR w": []string{"bf1-d3"},
	"rnbq1rk1/pppp1ppp/4pn2/8/1bPP4/2NBP3/PP3PPP/R1BQK1NR b": []string{"pc7-c5", "pd7-d5"},
}