package rollssim

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"testing"
)

func TestGenshinWantedRolls(t *testing.T) {
	warps := 1000
	wantedChars := 7
	wantedWeapons := 5
	iterations := 1000

	successCount := 0
	failureCount := 0
	rateUpCharCount := 0
	chosenWeaponCount := 0
	charRollCount := 0
	standardCharCount := 0
	standardLCCount := 0
	weaponRollCount := 0

	for i := 0; i < iterations; i++ {
		result := CalcGenshinWantedRolls(warps, wantedChars, wantedWeapons, &GenshinCharRoller{
			MihoyoRoller{
				CurrSRPity:         50,
				GuaranteedRateUpSR: true,
			},
		}, &GenshinWeaponRoller{})
		rateUpCharCount += result.CharacterBannerRateUpSRCount
		chosenWeaponCount += result.WeaponBannerChosenRateUpCount
		standardCharCount += result.CharacterBannerStdSRCount
		standardLCCount += result.WeaponBannerStdSRCount
		charRollCount += result.CharacterBannerRollCount
		weaponRollCount += result.WeaponBannerRollCount

		if result.CharacterBannerRateUpSRCount >= wantedChars && result.WeaponBannerChosenRateUpCount >= wantedWeapons {
			successCount++
		} else {
			failureCount++
		}
	}

	log.Println("Success:", successCount, "Failure:", failureCount)
	log.Printf("Success rate: %.4f%%", float64(successCount*100)/float64(iterations))
	log.Printf("Average Rate Up char count: %.4f", float64(rateUpCharCount)/float64(iterations))
	log.Printf("Average Standard char count: %.4f", float64(standardCharCount)/float64(iterations))
	log.Printf("Average Character banner roll count: %.4f", float64(charRollCount)/float64(iterations))
	log.Printf("Average Rate Up Weapon count: %.4f", float64(chosenWeaponCount)/float64(iterations))
	log.Printf("Average Standard Weapon count: %.4f", float64(standardLCCount)/float64(iterations))
	log.Printf("Average Weapon banner roll count: %.4f", float64(weaponRollCount)/float64(iterations))
}

func TestStarRailWantedRolls(t *testing.T) {
	warps := 300
	wantedChars := 7
	wantedLCs := 0
	iterations := 1000

	successCount := 0
	failureCount := 0
	rateUpCharCount := 0
	rateUpLCCount := 0
	standardCharCount := 0
	standardLCCount := 0

	for i := 0; i < iterations; i++ {
		result := CalcGenshinWantedRolls(warps, wantedChars, wantedLCs, &StarRailCharRoller{
			MihoyoRoller{
				CurrSRPity:         50,
				GuaranteedRateUpSR: true,
			},
		}, &StarRailLCRoller{})
		rateUpCharCount += result.CharacterBannerRateUpSRCount
		rateUpLCCount += result.WeaponBannerRateUpSRCount
		standardCharCount += result.CharacterBannerStdSRCount
		standardLCCount += result.WeaponBannerStdSRCount

		if result.CharacterBannerRateUpSRCount >= wantedChars && result.WeaponBannerRateUpSRCount >= wantedLCs {
			successCount++
		} else {
			failureCount++
		}
	}

	log.Println("Success:", successCount, "Failure:", failureCount)
	log.Printf("Success rate: %.4f%%", float64(successCount)/float64(iterations))
	log.Printf("Average Rate Up char count: %.4f", float64(rateUpCharCount)/float64(iterations))
	log.Printf("Average Rate Up LC count: %.4f", float64(rateUpLCCount)/float64(iterations))
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

func TestGenshinWeaponRollerRates(t *testing.T) {
	warps := 1_000_000
	fiveStarCount := 0

	roller := GenshinWeaponRoller{}
	for i := 0; i < warps; i++ {
		rolled := roller.Roll()
		if rolled.Rarity == 5 {
			fiveStarCount++
		}
	}

	log.Println("Five star LC count:", fiveStarCount)
	log.Printf("Five star LC consolidated rate: %.5f%%", float64(fiveStarCount*100)/float64(warps))
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
