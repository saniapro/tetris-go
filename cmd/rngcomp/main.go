package main

import (
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/saniapro/tetris/pkg/tetris"
)

func genPieces(n int, gen tetris.Nexter) []int {
	counts := make([]int, 7)
	for i := 0; i < n; i++ {
		piece := gen.Next()
		counts[piece]++
	}
	return counts
}
func main() {
	n := 1000 // number of pieces to generate

	fmt.Printf("=== Tetris Piece RNG Distribution Comparison ===\n")
	fmt.Printf("Sample size: %d pieces\n\n", n)

	rngSystems := []string{"7-Bag", "14-Bag", "ClassicNES", "rand/v2"}
	rngNexters := []func(int64) tetris.Nexter{
		tetris.NewBagGenerator,
		tetris.NewFourteenBagGenerator,
		tetris.NewClassicNexter,
		tetris.NewV2Nexter,
	}
	countsList := make([][]int, len(rngSystems))
	for i := range rngSystems {
		gen := rngNexters[i](42) // fixed seed for reproducibility
		countsList[i] = genPieces(n, gen)
	}
	// Print comparison table
	pieceNames := []string{"T (0)", "J (1)", "Z (2)", "O (3)", "S (4)", "L (5)", "I (6)"}

	fmt.Println("--- Distribution by Piece ---")
	fmt.Printf("%-10s ", "Piese")
	for i := range rngSystems {
		fmt.Printf("| %-14s ", rngSystems[i])
	}
	fmt.Print("\n")
	fmt.Println(strings.Repeat("-", 12+len(rngSystems)*16))

	for i := 0; i < 7; i++ {
		fmt.Printf("%-10s ", pieceNames[i])
		for j := range rngSystems {
			count := countsList[j][i]
			pct := 100.0 * float64(count) / float64(n)
			fmt.Printf("| %5d (%.2f%%) ", count, pct)
		}
		fmt.Print("\n")
	}

	// Chi-square test
	fmt.Println("\n--- Statistical Tests ---")
	expected := float64(n) / 7.0
	// fmt.Printf("Chi-square (lower = more uniform, critical ~12.6 at α=0.05):\n")
	for i, counts := range countsList {
		chi2 := chiSquare(counts, expected)
		fmt.Printf("   %s: %.4f\n", rngSystems[i], chi2)
	}

	// Standard deviation of frequencies
	fmt.Println("\n--- Uniformity (std dev of frequencies) ---")
	for i, counts := range countsList {
		std := stdDev(counts)
		fmt.Printf("   %s: %.2f\n", rngSystems[i], std)
	}

	// Drought analysis (intervals between I pieces)
	fmt.Println("\n--- Drought Analysis (intervals between I = piece 6) ---")

	for i, nexterFunc := range rngNexters {
		gen := nexterFunc(42)
		droughts := analyzeDroughts(gen, n)
		fmt.Printf("%s:\n", rngSystems[i])
		printDroughtStats(droughts)
	}

	// Min/max frequency analysis
	fmt.Println("\n--- Min/Max Frequency ---")
	fmt.Printf("%-13s | %-8s | %-8s | Range\n", "Method", "Min", "Max")
	fmt.Println(strings.Repeat("-", 40))
	for i, counts := range countsList {
		min, max := minMax(counts)
		fmt.Printf("   %-10s | %-8d | %-8d | %d\n", rngSystems[i], min, max, max-min)
	}

}

// chiSquare computes chi-square statistic for goodness of fit
func chiSquare(observed []int, expected float64) float64 {
	chi2 := 0.0
	for _, o := range observed {
		diff := float64(o) - expected
		chi2 += (diff * diff) / expected
	}
	return chi2
}

// stdDev computes standard deviation of frequencies
func stdDev(counts []int) float64 {
	mean := 0.0
	for _, c := range counts {
		mean += float64(c)
	}
	mean /= float64(len(counts))

	variance := 0.0
	for _, c := range counts {
		diff := float64(c) - mean
		variance += diff * diff
	}
	variance /= float64(len(counts))

	return math.Sqrt(variance)
}

// minMax finds min and max in slice
func minMax(counts []int) (int, int) {
	min, max := counts[0], counts[0]
	for _, c := range counts {
		if c < min {
			min = c
		}
		if c > max {
			max = c
		}
	}
	return min, max
}

// analyzeDroughts analyzes drought lengths using any Nexter implementation
func analyzeDroughts(gen tetris.Nexter, n int) []int {
	var droughts []int
	currentDrought := 0

	for i := 0; i < n; i++ {
		piece := gen.Next()
		if piece == 6 { // I piece
			droughts = append(droughts, currentDrought)
			currentDrought = 0
		} else {
			currentDrought++
		}
	}

	if currentDrought > 0 {
		droughts = append(droughts, currentDrought)
	}

	return droughts
}

// printDroughtStats prints summary statistics for droughts
func printDroughtStats(droughts []int) {
	if len(droughts) == 0 {
		fmt.Println("  No droughts recorded")
		return
	}

	sum := 0
	for _, d := range droughts {
		sum += d
	}
	mean := float64(sum) / float64(len(droughts))

	sorted := append([]int(nil), droughts...)
	sort.Ints(sorted)

	median := 0.0
	if len(sorted)%2 == 1 {
		median = float64(sorted[len(sorted)/2])
	} else {
		median = float64(sorted[len(sorted)/2-1]+sorted[len(sorted)/2]) / 2.0
	}

	max := sorted[len(sorted)-1]

	count20plus := 0
	for _, d := range droughts {
		if d >= 20 {
			count20plus++
		}
	}

	fmt.Printf("  Samples: %d\n", len(droughts))
	fmt.Printf("  Mean: %.2f, Median: %.2f, Max: %d\n", mean, median, max)
	fmt.Printf("  Droughts >=20: %d (%.2f%%)\n", count20plus, 100.0*float64(count20plus)/float64(len(droughts)))
}
