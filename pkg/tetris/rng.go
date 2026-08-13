package tetris

import rand "math/rand/v2"

// Nexter defines an interface for generating the next Tetris piece index.
type Nexter interface {
	Next() int
}

type V2Nexter struct {
	rng *rand.Rand
}

func NewV2Nexter(seed int64) Nexter {
	// Use PCG source from math/rand/v2 for a good seeded generator.
	s := uint64(seed)
	src := rand.NewPCG(s, s^0x9e3779b97f4a7c15)
	return &V2Nexter{
		rng: rand.New(src),
	}
}

func (vn *V2Nexter) Next() int {
	return vn.rng.IntN(7)
}

// ClassicNexter implements the classic NES Tetris piece generation algorithm
// with anti-duplication (no two consecutive pieces are the same).
type ClassicNexter struct {
	rng  *rand.Rand
	last int
}

func NewClassicNexter(seed int64) Nexter {
	// Use PCG source from math/rand/v2 for a good seeded generator.
	s := uint64(seed)
	src := rand.NewPCG(s, s^0x9e3779b97f4a7c15)
	return &ClassicNexter{
		rng:  rand.New(src),
		last: -1,
	}
}

func (cn *ClassicNexter) Next() int {
	piece := cn.genPieceClassic()
	// Anti-duplication: if same as last, try again
	if cn.last != -1 && piece == cn.last {
		piece = cn.genPieceClassic()
	}
	cn.last = piece
	return piece
}

func (cn *ClassicNexter) genPieceClassic() int {
	for {
		r := cn.rng.IntN(8)
		if r != 7 {
			return r
		}
	}
}
