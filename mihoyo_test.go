package rollssim

import (
	"fmt"
	"log"
	"os"
	"sort"
	"testing"
)

func TestGenshinWantedRolls(t *testing.T) {
	warps := 791 - 9
	wantedChars := 11
	wantedWeapons := 0
	iterations := 100000

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
			MihoyoRoller{},
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
	log.Printf("Average Chosen Weapon count: %.4f", float64(chosenWeaponCount)/float64(iterations))
	log.Printf("Average Standard Weapon count: %.4f", float64(standardLCCount)/float64(iterations))
	log.Printf("Average Weapon banner roll count: %.4f", float64(weaponRollCount)/float64(iterations))
}

func TestStarRailWantedRolls(t *testing.T) {
	warps := 300
	wantedChars := 3
	wantedLCs := 1
	iterations := 1000

	successCount := 0
	failureCount := 0
	rateUpCharCount := 0
	rateUpLCCount := 0
	standardCharCount := 0
	standardLCCount := 0

	for i := 0; i < iterations; i++ {
		result := CalcStarRailWantedRolls(warps, wantedChars, wantedLCs, &StarRailCharRoller{
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
	warps := 28_258_375
	fiveStarCount := 0
	rateUpCount := 0
	wins := 0
	losses := 0

	roller := StarRailCharRoller{}
	guarantee := false
	for i := 0; i < warps; i++ {
		char := roller.Roll()
		if char.Rarity == 5 {
			fiveStarCount++
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

	log.Println("Five star char count:", fiveStarCount)
	log.Println("rateUpCount:", rateUpCount)
	log.Println("Rate up rate:", float64(rateUpCount)/float64(fiveStarCount))
	log.Println("wins:", wins, "losses:", losses, "rate:", float64(wins)/float64(wins+losses))
	log.Printf("Five star char consolidated rate: %.5f%%", float64(fiveStarCount)/float64(warps))
}

func TestStarRailLCRollerRates(t *testing.T) {
	warps := 100_000_000
	fiveStarCount := 0
	rateUpCount := 0
	allNeededRolls := []int{}

	roller := StarRailLCRoller{}
	for i := 0; i < warps; i++ {
		neededRolls := roller.CurrSRPity + 1
		char := roller.Roll()
		if char.Rarity == 5 {
			fiveStarCount++
			if char.IsRateUp {
				rateUpCount++
			}
			allNeededRolls = append(allNeededRolls, neededRolls)
		}
	}

	log.Println("Five star LC count:", fiveStarCount)
	log.Println("Rate up count:", rateUpCount)
	log.Println("Rate up rate:", float64(rateUpCount)/float64(fiveStarCount))
	log.Printf("Five star LC consolidated rate: %.5f%%", float64(fiveStarCount)/float64(warps))

	/*err := makeFile("lc_sr_needed_rolls.csv", allNeededRolls)
	if err != nil {
		log.Fatal(err)
	}*/
}

func TestGenshinWeaponRollerRates(t *testing.T) {
	pulls := 1000000
	fiveStarCount := 0
	allNeededRolls := []int{}
	pityToFiveStarCount := make([]int, 80)
	csv := ""

	roller := GenshinWeaponRoller{}
	for i := 0; i < pulls; i++ {
		pityBeforeRolling := roller.CurrSRPity
		rolled := roller.Roll()
		if rolled.Rarity == 5 {
			fiveStarCount++
			allNeededRolls = append(allNeededRolls, pityBeforeRolling+1)
		}
	}

	for _, neededRolls := range allNeededRolls {
		// index 0 for neededRolls = 1
		pityToFiveStarCount[neededRolls-1]++
	}

	for i := 0; i < 80; i++ {
		csv += fmt.Sprintf("%d	%d\n", i+1, pityToFiveStarCount[i])
	}

	log.Println("First pull chance:", float64(pityToFiveStarCount[0])/float64(len(allNeededRolls)))

	log.Println("Iteration count:", pulls)
	log.Println("Five star Weapon count:", fiveStarCount)
	log.Printf("Five star Weapon consolidated rate: %.5f%%", float64(fiveStarCount*100)/float64(pulls))
	err := makeFile("genshin_weapon_needed_rolls.csv", csv)
	if err != nil {
		log.Fatal(err)
	}
}

func TestStarRailRareCharRollerRates(t *testing.T) {
	warps := 10000_000
	rareCount := 0

	roller := StarRailCharRoller{}
	for i := 0; i < warps; i++ {
		char := roller.Roll()
		if char.Rarity == 4 {
			rareCount++
		}
	}

	log.Println("Rare char count:", rareCount)
	log.Printf("Rare char consolidated rate: %.5f%%", float64(rareCount)/float64(warps))
}

func TestGenshinCharLuckChart(t *testing.T) {
	iterations := 500_000
	neededRollsPerConstellationStopPoint := map[int][]int{}
	pPoints := []float32{0.1, 0.5, 1, 5, 10, 15, 20, 25, 30, 35, 40, 45, 50, 55, 60, 65, 70, 75, 80, 85, 90, 95, 99, 99.5, 99.9}

	for cons := 0; cons < 7; cons++ {
		wantedChars := cons + 1
		maxNeededRolls := wantedChars * 180
		neededRollsPerConstellationStopPoint[cons] = make([]int, iterations)
		for i := 0; i < iterations; i++ {
			roller := GenshinCharRoller{}
			result := CalcGenshinWantedRolls(maxNeededRolls, wantedChars, 0, &roller, nil)
			neededRollsPerConstellationStopPoint[cons][i] = result.CharacterBannerRollCount
		}
		sort.Ints(neededRollsPerConstellationStopPoint[cons])
	}

	csv := ""
	for cons := 0; cons < 7; cons++ {
		rolls := neededRollsPerConstellationStopPoint[cons]
		for _, p := range pPoints {
			pRolls := rolls[int(float32(len(rolls))*p/100)]
			csv += fmt.Sprintf("%.1f	%d	%d\n", p, rolls[0], pRolls)
		}
	}
	err := makeFile("genshin_char_luck_chart.csv", csv)
	if err != nil {
		log.Fatal(err)
	}
}

func TestStarRailCharLuckChart(t *testing.T) {
	iterations := 500_000
	neededRollsPerConstellationStopPoint := map[int][]int{}
	pPoints := []float32{0.1, 0.5, 1, 5, 10, 15, 20, 25, 30, 35, 40, 45, 50, 55, 60, 65, 70, 75, 80, 85, 90, 95, 99, 99.5, 99.9}

	for cons := 0; cons < 7; cons++ {
		wantedChars := cons + 1
		maxNeededRolls := wantedChars * 180
		neededRollsPerConstellationStopPoint[cons] = make([]int, iterations)
		for i := 0; i < iterations; i++ {
			roller := StarRailCharRoller{}
			result := CalcStarRailWantedRolls(maxNeededRolls, wantedChars, 0, &roller, nil)
			neededRollsPerConstellationStopPoint[cons][i] = result.CharacterBannerRollCount
		}
		sort.Ints(neededRollsPerConstellationStopPoint[cons])
	}

	csv := ""
	for cons := 0; cons < 7; cons++ {
		rolls := neededRollsPerConstellationStopPoint[cons]
		for _, p := range pPoints {
			pRolls := rolls[int(float32(len(rolls))*p/100)]
			csv += fmt.Sprintf("%.1f	%d	%d\n", p, rolls[0], pRolls)
		}
	}
	err := makeFile("star_rail_char_luck_chart.csv", csv)
	if err != nil {
		log.Fatal(err)
	}
}

func TestGenshinWeaponLuckChart(t *testing.T) {
	iterations := 500_000
	neededRollsPerRefineStopPoint := map[int][]int{}
	pPoints := []float32{0.1, 0.5, 1, 5, 10, 15, 20, 25, 30, 35, 40, 45, 50, 55, 60, 65, 70, 75, 80, 85, 90, 95, 99, 99.5, 99.9}

	for refines := 0; refines < 5; refines++ {
		wantedWeapons := refines + 1
		maxNeededRolls := wantedWeapons * 78 * 3
		neededRollsPerRefineStopPoint[refines] = make([]int, iterations)
		for i := 0; i < iterations; i++ {
			roller := GenshinWeaponRoller{}
			result := CalcGenshinWantedRolls(maxNeededRolls, 0, wantedWeapons, nil, &roller)
			neededRollsPerRefineStopPoint[refines][i] = result.WeaponBannerRollCount
		}
		sort.Ints(neededRollsPerRefineStopPoint[refines])
	}

	csv := ""
	for refines := 0; refines < 5; refines++ {
		rolls := neededRollsPerRefineStopPoint[refines]
		for _, p := range pPoints {
			pRolls := rolls[int(float32(len(rolls))*p/100)]
			csv += fmt.Sprintf("%.1f	%d	%d\n", p, rolls[0], pRolls)
		}
	}
	err := makeFile("genshin_weapon_luck_chart.csv", csv)
	if err != nil {
		log.Fatal(err)
	}
}

func TestStarRailLightConeLuckChart(t *testing.T) {
	iterations := 500_000
	neededRollsPerRefineStopPoint := map[int][]int{}
	pPoints := []float32{0.1, 0.5, 1, 5, 10, 15, 20, 25, 30, 35, 40, 45, 50, 55, 60, 65, 70, 75, 80, 85, 90, 95, 99, 99.5, 99.9}

	for refines := 0; refines < 5; refines++ {
		wantedWeapons := refines + 1
		maxNeededRolls := wantedWeapons * 78 * 2
		neededRollsPerRefineStopPoint[refines] = make([]int, iterations)
		for i := 0; i < iterations; i++ {
			roller := StarRailLCRoller{}
			result := CalcStarRailWantedRolls(maxNeededRolls, 0, wantedWeapons, nil, &roller)
			neededRollsPerRefineStopPoint[refines][i] = result.WeaponBannerRollCount
		}
		sort.Ints(neededRollsPerRefineStopPoint[refines])
	}

	csv := ""
	for refines := 0; refines < 5; refines++ {
		rolls := neededRollsPerRefineStopPoint[refines]
		for _, p := range pPoints {
			pRolls := rolls[int(float32(len(rolls))*p/100)]
			csv += fmt.Sprintf("%.1f	%d	%d\n", p, rolls[0], pRolls)
		}
	}
	err := makeFile("star_rail_lc_luck_chart.csv", csv)
	if err != nil {
		log.Fatal(err)
	}
}

func makeFile(filename string, data string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(data)
	return err
}
