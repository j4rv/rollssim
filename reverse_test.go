package rollssim

import (
	"fmt"
	"log"
	"sort"
	"strings"
	"testing"
)

func TestReverseCharRollerRates(t *testing.T) {
	warps := 100_000_000
	sixStarCount := 0
	rateUpCount := 0
	wins := 0
	losses := 0

	roller := ReverseCharRoller{}
	guarantee := false
	for i := 0; i < warps; i++ {
		char := roller.Roll()
		if char.Rarity == 5 {
			sixStarCount++
			if !guarantee {
				if char.IsRateUp {
					wins++
				} else {
					losses++
				}
			}
			if char.IsRateUp {
				rateUpCount++
				guarantee = false
			} else {
				guarantee = true
			}
		}
	}

	log.Println("6 star char count:", sixStarCount)
	log.Println("rateUpCount:", rateUpCount)
	log.Println("Rate up rate:", float64(rateUpCount)/float64(sixStarCount))
	log.Println("wins:", wins, "losses:", losses, "rate:", float64(wins)/float64(wins+losses))
	log.Println("Average roll count:", float64(warps)/float64(sixStarCount))
	log.Printf("Five star char consolidated rate: %.5f%%", float64(sixStarCount)/float64(warps))
}

func TestReverseCharLuckChart(t *testing.T) {
	iterations := 1_000_000
	maxDupes := 6
	neededRollsPerDupeStopPoint := map[int][]int{}
	pPoints := []float32{0.1, 0.5, 1, 5, 10, 15, 20, 25, 30, 35, 40, 45, 50, 55, 60, 65, 70, 75, 80, 85, 90, 95, 99, 99.5, 99.9}

	for cons := 0; cons < maxDupes; cons++ {
		wantedChars := cons + 1
		maxNeededRolls := wantedChars * 70 * 2
		neededRollsPerDupeStopPoint[cons] = make([]int, iterations)
		for i := 0; i < iterations; i++ {
			roller := ReverseCharRoller{}
			result := CalcReverseWantedRolls(maxNeededRolls, wantedChars, &roller)
			neededRollsPerDupeStopPoint[cons][i] = result.CharacterBannerRollCount
		}
		sort.Ints(neededRollsPerDupeStopPoint[cons])
	}

	csv := ""
	for cons := 0; cons < maxDupes; cons++ {
		rolls := neededRollsPerDupeStopPoint[cons]
		for _, p := range pPoints {
			pRolls := rolls[int(float32(len(rolls))*p/100)]
			csv += fmt.Sprintf("%.1f	%d	%d\n", p, rolls[0], pRolls)
		}
	}
	csv = strings.ReplaceAll(csv, ".", ",")
	err := makeFile("reverse_char_luck_chart.csv", csv)
	if err != nil {
		log.Fatal(err)
	}
}
