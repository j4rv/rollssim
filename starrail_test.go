package rollssim

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"testing"
)

func TestStarRailWantedRolls(t *testing.T) {
	warps := 800
	wantedChars := 7
	wantedLCs := 1
	iterations := 100000

	successCount := 0
	failureCount := 0
	standardCharCount := 0
	standardLCCount := 0

	for i := 0; i < iterations; i++ {
		result := CalcWantedRolls(warps, wantedChars, wantedLCs, &StarRailCharRoller{}, &StarRailLCRoller{})
		standardCharCount += result.StdSRCharCount
		standardLCCount += result.StdLCSRCount

		if result.RateUpSRCharCount >= wantedChars && result.RateUpSRLCCount >= wantedLCs {
			successCount++
		} else {
			failureCount++
		}
	}

	log.Println("Success:", successCount, "Failure:", failureCount)
	log.Printf("Success rate: %.4f%%", float64(successCount)/float64(iterations))
	log.Printf("Average Standard char count: %.4f", float64(standardCharCount)/float64(iterations))
	log.Printf("Average Standard LC count: %.4f", float64(standardLCCount)/float64(iterations))
}

func TestStarRailCharRollerRates(t *testing.T) {
	warps := 1_000_000
	fiveStarCount := 0

	roller := StarRailCharRoller{}
	for i := 0; i < warps; i++ {
		char := roller.Roll()
		if char.Rarity == 5 {
			fiveStarCount++
		}
	}

	log.Println("Five star char count:", fiveStarCount)
	log.Printf("Five star char consolidated rate: %.5f%%", float64(fiveStarCount)/float64(warps))
}

func TestStarRailLCRollerRates(t *testing.T) {
	warps := 1_000_000
	fiveStarCount := 0
	allNeededRolls := []int{}

	roller := StarRailLCRoller{}
	for i := 0; i < warps; i++ {
		neededRolls := roller.CurrSRPity + 1
		char := roller.Roll()
		if char.Rarity == 5 {
			fiveStarCount++
			allNeededRolls = append(allNeededRolls, neededRolls)
		}
	}

	log.Println("Five star LC count:", fiveStarCount)
	log.Printf("Five star LC consolidated rate: %.5f%%", float64(fiveStarCount)/float64(warps))

	err := makeFile("lc_sr_needed_rolls.csv", allNeededRolls)
	if err != nil {
		log.Fatal(err)
	}
}

func TestStarRailRareCharRollerRates(t *testing.T) {
	warps := 200_000
	rareCount := 0
	allNeededRolls := []int{}

	roller := StarRailCharRoller{}
	for i := 0; i < warps; i++ {
		neededRolls := roller.CurrRarePity + 1
		char := roller.Roll()
		if char.Rarity == 4 {
			rareCount++
			allNeededRolls = append(allNeededRolls, neededRolls)
		}
	}

	log.Println("Rare char count:", rareCount)
	log.Printf("Rare char consolidated rate: %.5f%%", float64(rareCount)/float64(warps))

	err := makeFile("rare_char_needed_rolls.csv", allNeededRolls)
	if err != nil {
		log.Fatal(err)
	}
}

func makeFile[T any](filename string, data []T) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	for _, roll := range data {
		if err := w.Write([]string{fmt.Sprintf("%v", roll)}); err != nil {
			return err
		}
	}

	return nil
}
