package rollssim

import (
	"log"
	"testing"
)

const consolidatedRatesRolls = 100_000_000

func TestGenshinCharConsolidatedRates(t *testing.T) {
	rolls := consolidatedRatesRolls
	result := CalcCharacterBannerRolls(rolls, &GenshinCharRoller{})

	srChars := float64(result.CharacterBannerStdSRCount) + float64(result.CharacterBannerRateUpSRCount)
	log.Printf("Character Banner 5* cumulative rates: %.2f%%", srChars/float64(rolls)*100)
	rareChars := float64(result.CharacterBannerStdRareCount) + float64(result.CharacterBannerRateUpRareCount)
	log.Printf("Character Banner 4* cumulative rates: %.2f%%", rareChars/float64(rolls)*100)
}

func TestGenshinWeaponConsolidatedRates(t *testing.T) {
	rolls := consolidatedRatesRolls
	result := CalcWeaponBannerRolls(rolls, &GenshinWeaponRoller{})

	srWeapons := float64(result.WeaponBannerStdSRCount) + float64(result.WeaponBannerRateUpSRCount)
	log.Printf("Weapon Banner 5* cumulative rates: %.2f%%", srWeapons/float64(rolls)*100)
	rareWeapons := float64(result.WeaponBannerStdRareCount) + float64(result.WeaponBannerRateUpRareCount)
	log.Printf("Weapon Banner 4* cumulative rates: %.2f%%", rareWeapons/float64(rolls)*100)
}

func TestStarRailCharConsolidatedRates(t *testing.T) {
	rolls := consolidatedRatesRolls
	result := CalcCharacterBannerRolls(rolls, &StarRailCharRoller{})

	srChars := float64(result.CharacterBannerStdSRCount) + float64(result.CharacterBannerRateUpSRCount)
	log.Printf("Character Banner 5* cumulative rates: %.2f%%", srChars/float64(rolls)*100)
	rareChars := float64(result.CharacterBannerStdRareCount) + float64(result.CharacterBannerRateUpRareCount)
	log.Printf("Character Banner 4* cumulative rates: %.2f%%", rareChars/float64(rolls)*100)
}

func TestStarRailWeaponConsolidatedRates(t *testing.T) {
	rolls := consolidatedRatesRolls
	result := CalcWeaponBannerRolls(rolls, &StarRailLCRoller{})

	srWeapons := float64(result.WeaponBannerStdSRCount) + float64(result.WeaponBannerRateUpSRCount)
	log.Printf("Weapon Banner 5* cumulative rates: %.2f%%", srWeapons/float64(rolls)*100)
	rareWeapons := float64(result.WeaponBannerStdRareCount) + float64(result.WeaponBannerRateUpRareCount)
	log.Printf("Weapon Banner 4* cumulative rates: %.2f%%", rareWeapons/float64(rolls)*100)
}

func TestZenlessCharConsolidatedRates(t *testing.T) {
	rolls := consolidatedRatesRolls
	result := CalcCharacterBannerRolls(rolls, &ZenlessCharRoller{})

	srChars := float64(result.CharacterBannerStdSRCount) + float64(result.CharacterBannerRateUpSRCount)
	log.Printf("Character Banner 5* cumulative rates: %.2f%%", srChars/float64(rolls)*100)
	rareChars := float64(result.CharacterBannerStdRareCount) + float64(result.CharacterBannerRateUpRareCount)
	log.Printf("Character Banner 4* cumulative rates: %.2f%%", rareChars/float64(rolls)*100)
}

func TestZenlessWeaponConsolidatedRates(t *testing.T) {
	rolls := consolidatedRatesRolls
	result := CalcWeaponBannerRolls(rolls, &ZenlessEngineRoller{})

	srWeapons := float64(result.WeaponBannerStdSRCount) + float64(result.WeaponBannerRateUpSRCount)
	log.Printf("Weapon Banner 5* cumulative rates: %.2f%%", srWeapons/float64(rolls)*100)
	rareWeapons := float64(result.WeaponBannerStdRareCount) + float64(result.WeaponBannerRateUpRareCount)
	log.Printf("Weapon Banner 4* cumulative rates: %.2f%%", rareWeapons/float64(rolls)*100)
}
