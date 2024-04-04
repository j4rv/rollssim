package rollssim

import (
	"math/rand"
)

const SR_CHAR_BANNER_BASE_CHANCE = 0.006
const SR_WEAPON_BANNER_BASE_CHANCE = 0.007
const SR_LIGHT_CONE_BANNER_BASE_CHANCE = 0.008
const RARE_CHAR_BANNER_BASE_CHANCE = 0.051
const RARE_WEAPON_BANNER_BASE_CHANCE = 0.060
const RARE_LIGHT_CONE_BANNER_BASE_CHANCE = 0.066

type MihoyoRoller struct {
	CurrSRPity           int
	CurrRarePity         int
	GuaranteedRateUpSR   bool
	GuaranteedRateUpRare bool
}

func (s *MihoyoRoller) roll(srChance, rareChance float64, rateUpSRChance float64, rateUpSRItems, rateUpRareItems []string) Rollable {
	s.CurrSRPity++
	s.CurrRarePity++

	// Check if we get a SR
	if rand.Float64() <= srChance {
		s.CurrSRPity = 0
		if s.GuaranteedRateUpSR || rand.Float64() <= rateUpSRChance {
			s.GuaranteedRateUpSR = false
			return Rollable{Name: SuperRare, Type: SuperRare, Rarity: 5, IsRateUp: true}
		} else {
			s.GuaranteedRateUpSR = true
			return Rollable{Name: RandomString(rateUpSRItems), Type: SuperRare, Rarity: 5}
		}
	}

	// Check if we get a Rare
	if rand.Float64() <= rareChance {
		s.CurrRarePity = 0
		if s.GuaranteedRateUpRare || rand.Float64() <= 0.5 {
			s.GuaranteedRateUpRare = false
			return Rollable{Name: RandomString(rateUpRareItems), Type: Rare, Rarity: 4, IsRateUp: true}
		} else {
			s.GuaranteedRateUpRare = true
			return Rollable{Name: Rare, Type: Rare, Rarity: 4}
		}
	}

	return Rollable{Name: Fodder, Type: Fodder, Rarity: 3}
}

type GenshinCharRoller struct {
	MihoyoRoller
}

var StandardGenshinSRChars = []string{"Diluc", "Jean", "Qiqi", "Keqing", "Mona", "Tighnari", "Dehya"}

func (s *GenshinCharRoller) Roll() Rollable {
	srChance := SR_CHAR_BANNER_BASE_CHANCE
	if s.CurrSRPity+1 >= 74 {
		srChance += SR_CHAR_BANNER_BASE_CHANCE * 10 * float64(s.CurrSRPity+1-73)
	}
	rareChance := RARE_CHAR_BANNER_BASE_CHANCE
	if s.CurrRarePity+1 >= 9 {
		rareChance += RARE_CHAR_BANNER_BASE_CHANCE * 10 * float64(s.CurrRarePity+1-8)
	}
	return s.MihoyoRoller.roll(srChance, rareChance, 0.5, StandardGenshinSRChars, ThreeRateUpRares)
}

type GenshinWeaponRoller struct {
	MihoyoRoller
	FatePoints int
}

var StandardGenshinSRWeapons = []string{"TODO"}

func (s *GenshinWeaponRoller) Roll() Rollable {
	s.CurrSRPity++
	s.CurrRarePity++

	srChance := SR_WEAPON_BANNER_BASE_CHANCE
	if s.CurrSRPity >= 63 {
		srChance += SR_WEAPON_BANNER_BASE_CHANCE * 10 * float64(s.CurrSRPity-62)
	}
	rareChance := RARE_WEAPON_BANNER_BASE_CHANCE
	if s.CurrRarePity >= 9 {
		rareChance += RARE_WEAPON_BANNER_BASE_CHANCE * 10 * float64(s.CurrRarePity-8)
	}

	// Check if we get a SR
	if rand.Float64() <= srChance {
		s.CurrSRPity = 0
		r := rand.Float64()
		if r <= 0.375 || s.FatePoints >= 2 {
			s.FatePoints = 0
			return Rollable{Name: ChosenWeapon, Type: SuperRare, Rarity: 5, IsRateUp: true}
		} else if r <= 0.75 {
			s.FatePoints += 1
			return Rollable{Name: NotChosenWeapon, Type: SuperRare, Rarity: 5, IsRateUp: true}
		} else {
			s.FatePoints += 1
			return Rollable{Name: RandomString(StandardGenshinSRWeapons), Type: SuperRare, Rarity: 5}
		}
	}

	// Check if we get a Rare
	if rand.Float64() <= rareChance {
		s.CurrRarePity = 0
		if s.GuaranteedRateUpRare || rand.Float64() <= 0.5 {
			s.GuaranteedRateUpRare = false
			return Rollable{Name: RandomString(FiveRateUpRares), Type: Rare, Rarity: 4, IsRateUp: true}
		} else {
			s.GuaranteedRateUpRare = true
			return Rollable{Name: Rare, Type: Rare, Rarity: 4}
		}
	}

	return Rollable{Name: Fodder, Type: Fodder, Rarity: 3}
}

// --- Star Rail ---

type StarRailCharRoller struct {
	MihoyoRoller
}

var StandardStarRailSRChars = []string{"Bailu", "Yanqing", "Clara", "Gepard", "Bronya", "Welt", "Himeko"}

func (s *StarRailCharRoller) Roll() Rollable {
	srChance := SR_CHAR_BANNER_BASE_CHANCE
	if s.CurrSRPity+1 >= 74 {
		srChance += SR_CHAR_BANNER_BASE_CHANCE * 10 * float64(s.CurrSRPity+1-73)
	}
	rareChance := RARE_CHAR_BANNER_BASE_CHANCE
	if s.CurrRarePity+1 >= 9 {
		rareChance += RARE_CHAR_BANNER_BASE_CHANCE * 10 * float64(s.CurrRarePity+1-8)
	}
	return s.MihoyoRoller.roll(srChance, rareChance, 0.5, StandardStarRailSRChars, ThreeRateUpRares)
}

type StarRailLCRoller struct {
	MihoyoRoller
}

var StandardStarRailSRLCs = []string{"Bailu LC", "Yanqing LC", "Clara LC", "Gepard LC", "Bronya LC", "Welt LC", "Himeko LC"}

func (s *StarRailLCRoller) Roll() Rollable {
	srChance := SR_LIGHT_CONE_BANNER_BASE_CHANCE
	if s.CurrSRPity+1 >= 66 {
		srChance += SR_LIGHT_CONE_BANNER_BASE_CHANCE * 10 * float64(s.CurrSRPity+1-65)
	}
	rareChance := RARE_LIGHT_CONE_BANNER_BASE_CHANCE
	if s.CurrRarePity+1 >= 9 {
		rareChance += RARE_LIGHT_CONE_BANNER_BASE_CHANCE * 10 * float64(s.CurrRarePity+1-8)
	}
	return s.MihoyoRoller.roll(srChance, rareChance, 0.75, StandardStarRailSRLCs, ThreeRateUpRares)
}
