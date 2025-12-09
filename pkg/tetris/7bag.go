package tetris

import rand "math/rand/v2"

// BagGenerator implements the 7-bag random tetromino selection algorithm.
// Ensures each of the 7 tetrominoes appears exactly once per bag before reshuffling.
// This prevents long droughts of specific pieces.
type BagGenerator struct {
	bag []int      // Current bag of piece indices
	i   int        // Current position in bag
	rng *rand.Rand // Random number generator
}

// NewBagGenerator creates a new 7-bag generator with the given seed.
// Immediately generates and shuffles the first bag.
func NewBagGenerator(seed int64) *BagGenerator {
	// Use PCG source from math/rand/v2 for a good seeded generator.
	s := uint64(seed)
	src := rand.NewPCG(s, s^0x9e3779b97f4a7c15)
	g := &BagGenerator{
		rng: rand.New(src),
	}
	g.refill()
	return g
}

// refill generates a new shuffled bag containing all 7 piece types (0-6).
func (g *BagGenerator) refill() {
	g.bag = []int{0, 1, 2, 3, 4, 5, 6}
	g.rng.Shuffle(len(g.bag), func(i, j int) {
		g.bag[i], g.bag[j] = g.bag[j], g.bag[i]
	})
	g.i = 0
}

// Next returns the next piece index from the current bag.
// Automatically refills the bag when exhausted.
func (g *BagGenerator) Next() int {
	if g.i >= len(g.bag) {
		g.refill()
	}
	p := g.bag[g.i]
	g.i++
	return p
}
