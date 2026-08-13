package main

import (
	"fmt"
	"math/rand"
	"sort"
)

// Симуляція генератора фігур NES Tetris (1989)
// Правила:
// - генеруємо число 0..7
// - якщо 7 -> повторюємо (кидаємо знову) до отримання 0..6
// - якщо candidate == last_piece -> робимо ще одну перегенерацію (і ця повторна спроба може теж дорівнювати last_piece)
// Відповідність чисел фігурам:
// 0: T
// 1: J
// 2: Z
// 3: O
// 4: S
// 5: L
// 6: I

func main() {
	n := 100000 // кількість згенерованих фігур

	counts := make([]int, 7) // лічильник по кожній фігурі 0..6

	// для аналізу голодувань на I
	var droughts []int
	currentDrought := 0

	last := -1 // остання прийнята фігура (-1 означає, що ще не було)

	for i := 0; i < n; i++ {
		candidate := genPiece()

		// анти-дублювання: якщо candidate == last -> робимо ще одну спробу
		if last != -1 && candidate == last {
			candidate = genPiece()
		}

		counts[candidate]++

		if candidate == 6 { // I
			// закінчили голодування
			droughts = append(droughts, currentDrought)
			currentDrought = 0
		} else {
			currentDrought++
		}

		last = candidate
	}

	// Якщо остання серія без I не закінчилась I — можна врахувати як незавершене голодування;
	// тут додамо її до списку, щоб не втратити статистику (але позначимо це виводом).
	if currentDrought > 0 {
		droughts = append(droughts, currentDrought)
	}

	// Статистика по фігурам
	fmt.Printf("Симуляція %d фігур (алгоритм NES Tetris)\n\n", n)
	pieceNames := []string{"T", "J", "Z", "O", "S", "L", "I"}
	total := 0
	for _, c := range counts {
		total += c
	}

	fmt.Println("--- Поширеність фігур ---")
	for i, c := range counts {
		pct := 100.0 * float64(c) / float64(total)
		fmt.Printf("%s: %6d (%.3f%%)\n", pieceNames[i], c, pct)
	}

	// Статистика голодувань I
	if len(droughts) == 0 {
		fmt.Println("\nНе знайдено жодного I у симуляції.")
		return
	}

	sorted := append([]int(nil), droughts...)
	sort.Ints(sorted)

	// обчислимо середнє, медіану, макс, частку >=20
	sum := 0
	ge20 := 0
	max := 0
	for _, d := range droughts {
		sum += d
		if d >= 20 {
			ge20++
		}
		if d > max {
			max = d
		}
	}

	mean := float64(sum) / float64(len(droughts))
	median := 0.0
	m := len(sorted)
	if m%2 == 1 {
		median = float64(sorted[m/2])
	} else {
		median = float64(sorted[m/2-1]+sorted[m/2]) / 2.0
	}

	fmt.Println("\n--- Статистика голодувань I (інтервали між I) ---")
	fmt.Printf("Кількість зафіксованих інтервалів: %d\n", len(droughts))
	fmt.Printf("Середня довжина: %.4f\n", mean)
	fmt.Printf("Медіана: %.4f\n", median)
	fmt.Printf("Максимальна довжина: %d\n", max)
	fmt.Printf("Частка інтервалів >=20: %d (%.4f%%)\n", ge20, 100.0*float64(ge20)/float64(len(droughts)))

	// Гістограма (підрахунок частот для невеликого діапазону)
	fmt.Println("\n--- Гістограма голодувань (довжина -> частота) ---")
	hist := make(map[int]int)
	for _, d := range droughts {
		hist[d]++
	}

	// Виведемо перші 30 бінів (0..30), а решта зведемо в "31+"
	maxBin := 30
	for i := 0; i <= maxBin; i++ {
		if cnt, ok := hist[i]; ok {
			fmt.Printf("%2d: %6d\n", i, cnt)
		} else {
			fmt.Printf("%2d: %6d\n", i, 0)
		}
	}
	// сума для >30
	sumGt := 0
	for k, v := range hist {
		if k > maxBin {
			sumGt += v
		}
	}
	fmt.Printf("%2d+: %6d\n", maxBin+1, sumGt)

	// Додаткова інформація: найпоширеніші довжини
	fmt.Println("\n--- Топ-10 найчастіших довжин голодувань ---")
	type kv struct{ k, v int }
	var list []kv
	for k, v := range hist {
		list = append(list, kv{k, v})
	}
	sort.Slice(list, func(i, j int) bool { return list[i].v > list[j].v })
	for i := 0; i < 10 && i < len(list); i++ {
		fmt.Printf("%2d. Довжина %2d -> %6d раз(и)\n", i+1, list[i].k, list[i].v)
	}

	// Заключні зауваги
	fmt.Println("\n(Примітка: через відкидання значення 7 й анти-дублювання точні теоретичні ймовірності трохи відрізняються від простого 1/7. Рівномірність і частоти будуть зближуватись при великому n.)")
}

// genPiece генерує число 0..6 за алгоритмом з "викиданням" 7
func genPiece() int {
	for {
		r := rand.Intn(8) // 0..7
		if r == 7 {
			continue
		}
		return r
	}
}
