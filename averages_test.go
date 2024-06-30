package rollssim

import (
	"fmt"
	"log"
	"strings"
	"testing"
)

func TestGenshinAveragesChart(t *testing.T) {
	iterations := 500_000
	t.Run("Genshin Impact Char Banner", func(t *testing.T) {
		t.Parallel()
		auxTestWeaponAveragesChart(iterations, "Genshin Impact Char Banner", "avgs_genshin_chars.csv", &GenshinCharRoller{})
	})
	t.Run("Genshin Impact Weapon Banner", func(t *testing.T) {
		t.Parallel()
		auxTestWeaponAveragesChart(iterations, "Genshin Impact Weapon Banner", "avgs_genshin_weapon.csv", &GenshinWeaponRoller{})
	})
}

func TestHSRAveragesChart(t *testing.T) {
	iterations := 500_000
	t.Run("HSR Char Banner", func(t *testing.T) {
		t.Parallel()
		auxTestWeaponAveragesChart(iterations, "HSR Char Banner", "avgs_hsr_chars.csv", &StarRailCharRoller{})
	})
	t.Run("HSR Weapon Banner", func(t *testing.T) {
		t.Parallel()
		auxTestWeaponAveragesChart(iterations, "HSR Weapon Banner", "avgs_hsr_weapons.csv", &StarRailLCRoller{})
	})
}

func TestZenlessAveragesChart(t *testing.T) {
	iterations := 500_000
	t.Run("Zenless Char Banner", func(t *testing.T) {
		t.Parallel()
		auxTestWeaponAveragesChart(iterations, "Zenless Char Banner", "avgs_zenless_chars.csv", &ZenlessCharRoller{})
	})
	t.Run("Zenless Weapon Banner", func(t *testing.T) {
		t.Parallel()
		auxTestWeaponAveragesChart(iterations, "Zenless Weapon Banner", "avgs_zenless_weapons.csv", &ZenlessEngineRoller{})
	})
}

func auxTestWeaponAveragesChart(iterations int, header, filename string, roller Roller) {
	rollsAmounts := []int{100, 1000, 10000}

	rateUpSRWeaponsPerRollAmount := [3]int{}
	stdSRWeaponsPerRollAmount := [3]int{}
	rateUpRaresPerRollAmount := [3]int{}
	stdRaresPerRollAmount := [3]int{}

	for it := 0; it < iterations; it++ {
		for i, rollAmount := range rollsAmounts {
			roller.Reset()
			result := CalcWeaponBannerRolls(rollAmount, roller)

			rateUpSRWeaponsPerRollAmount[i] += result.WeaponBannerRateUpSRCount
			stdSRWeaponsPerRollAmount[i] += result.WeaponBannerStdSRCount
			rateUpRaresPerRollAmount[i] += result.WeaponBannerRateUpRareCount
			stdRaresPerRollAmount[i] += result.WeaponBannerStdRareCount
		}
	}

	csv := fmt.Sprintln(header)
	csv += fmt.Sprintf("iterations:	%d\n", iterations)
	csv += "rollAmount	rateUp5*	std5*	rateUp4*	std4*\n"
	for i, rollAmount := range rollsAmounts {
		csv += fmt.Sprintf("%d	%d	%d	%d	%d\n", rollAmount, rateUpSRWeaponsPerRollAmount[i], stdSRWeaponsPerRollAmount[i], rateUpRaresPerRollAmount[i], stdRaresPerRollAmount[i])
	}
	csv = strings.ReplaceAll(csv, ".", ",")
	err := makeFile(filename, csv)
	if err != nil {
		log.Fatal(err)
	}
}
