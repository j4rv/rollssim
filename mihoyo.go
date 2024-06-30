package rollssim

import (
	"math/rand"
)

const MIHOYO_CHAR_BANNER_BASE_CHANCE = 0.006
const MIHOYO_RARE_WEAPON_BANNER_SOFT_PITY = 8

type MihoyoRoller struct {
	CurrSRPity           int
	CurrRarePity         int
	GuaranteedRateUpSR   bool
	GuaranteedRateUpRare bool
}

func (s *MihoyoRoller) Reset() {
	s.CurrSRPity = 0
	s.CurrRarePity = 0
	s.GuaranteedRateUpSR = false
	s.GuaranteedRateUpRare = false
}

func (s *MihoyoRoller) roll(srChance, rareChance float64, rateUpSRChance float64, rateUpRareChance float64, rateUpSRItems, rateUpRareItems []Rollable) Rollable {
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
			return RandomRollable(rateUpSRItems)
		}
	}

	// Check if we get a Rare
	if rand.Float64() <= rareChance {
		s.CurrRarePity = 0
		if s.GuaranteedRateUpRare || rand.Float64() <= rateUpRareChance {
			s.GuaranteedRateUpRare = false
			return RandomRollable(rateUpRareItems)
		} else {
			s.GuaranteedRateUpRare = true
			return Rollable{Name: Rare, Type: Rare, Rarity: 4}
		}
	}

	return Rollable{Name: Fodder, Type: Fodder, Rarity: 3}
}

// --- Genshin ---

const GENSHIN_CHAR_BANNER_SOFT_PITY = 74
const GENSHIN_CHAR_BANNER_RARE_SOFT_PITY = 9
const GENSHIN_WEAPON_BANNER_SOFT_PITY = 63

const GENSHIN_CHAR_BANNER_RARE_BASE_CHANCE = 0.051
const GENSHIN_WEAPON_BANNER_BASE_CHANCE = 0.007
const GENSHIN_WEAPON_BANNER_RARE_BASE_CHANCE = 0.060

type GenshinCharRoller struct {
	MihoyoRoller
}

var StandardGenshinSRChars = []Rollable{
	{Name: "Standard 5*", Type: SuperRare, Rarity: 5, IsRateUp: false},
}

func (s *GenshinCharRoller) Roll() Rollable {
	srChance := MIHOYO_CHAR_BANNER_BASE_CHANCE
	if s.CurrSRPity+1 >= GENSHIN_CHAR_BANNER_SOFT_PITY {
		srChance += MIHOYO_CHAR_BANNER_BASE_CHANCE * 10 * float64(s.CurrSRPity+1-(GENSHIN_CHAR_BANNER_SOFT_PITY-1))
	}
	rareChance := GENSHIN_CHAR_BANNER_RARE_BASE_CHANCE
	if s.CurrRarePity+1 >= GENSHIN_CHAR_BANNER_RARE_SOFT_PITY {
		rareChance += GENSHIN_CHAR_BANNER_RARE_BASE_CHANCE * 10 * float64(s.CurrRarePity+1-(GENSHIN_CHAR_BANNER_RARE_SOFT_PITY-1))
	}
	return s.MihoyoRoller.roll(srChance, rareChance, 0.5, 0.5, StandardGenshinSRChars, ThreeRateUpRares)
}

type GenshinWeaponRoller struct {
	MihoyoRoller
	FatePoints int
}

var StandardGenshinSRWeapons = []Rollable{
	{Name: "Standard 5*", Type: SuperRare, Rarity: 5, IsRateUp: false},
}

func (s *GenshinWeaponRoller) Roll() Rollable {
	s.CurrSRPity++
	s.CurrRarePity++

	srChance := GENSHIN_WEAPON_BANNER_BASE_CHANCE
	if s.CurrSRPity >= GENSHIN_WEAPON_BANNER_SOFT_PITY {
		srChance += GENSHIN_WEAPON_BANNER_BASE_CHANCE * 10 * float64(s.CurrSRPity-(GENSHIN_WEAPON_BANNER_SOFT_PITY-1))
	}
	rareChance := GENSHIN_WEAPON_BANNER_RARE_BASE_CHANCE
	if s.CurrRarePity >= MIHOYO_RARE_WEAPON_BANNER_SOFT_PITY {
		rareChance += GENSHIN_WEAPON_BANNER_RARE_BASE_CHANCE * 10 * float64(s.CurrRarePity-(MIHOYO_RARE_WEAPON_BANNER_SOFT_PITY-1))
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
			return RandomRollable(StandardGenshinSRWeapons)
		}
	}

	// Check if we get a Rare
	if rand.Float64() <= rareChance {
		s.CurrRarePity = 0
		if s.GuaranteedRateUpRare || rand.Float64() <= 0.75 {
			s.GuaranteedRateUpRare = false
			return RandomRollable(FiveRateUpRares)
		} else {
			s.GuaranteedRateUpRare = true
			return Rollable{Name: Rare, Type: Rare, Rarity: 4}
		}
	}

	return Rollable{Name: Fodder, Type: Fodder, Rarity: 3}
}

// --- Zenless Zone Zero ---

const ZZZ_CHAR_BANNER_SOFT_PITY = 74
const ZZZ_CHAR_BANNER_RARE_SOFT_PITY = 10
const ZZZ_ENGINE_SOFT_PITY = 66
const ZZZ_ENGINE_BANNER_RARE_SOFT_PITY = 10

const ZZZ_SR_CHAR_BANNER_BASE_CHANCE = 0.006
const ZZZ_SR_ENGINE_BANNER_BASE_CHANCE = 0.010
const ZZZ_RARE_CHAR_BANNER_BASE_CHANCE = 0.094
const ZZZ_RARE_ENGINE_BANNER_BASE_CHANCE = 0.150

type ZenlessCharRoller struct {
	MihoyoRoller
}

var StandardZenlessSRChars = []Rollable{
	{Name: "Standard 5*", Type: SuperRare, Rarity: 5, IsRateUp: false},
}

func (s *ZenlessCharRoller) Roll() Rollable {
	srChance := ZZZ_SR_CHAR_BANNER_BASE_CHANCE
	if s.CurrSRPity+1 >= ZZZ_CHAR_BANNER_SOFT_PITY {
		srChance += ZZZ_SR_CHAR_BANNER_BASE_CHANCE * 10 * float64(s.CurrSRPity+1-(ZZZ_CHAR_BANNER_SOFT_PITY-1))
	}
	rareChance := ZZZ_RARE_CHAR_BANNER_BASE_CHANCE
	if s.CurrRarePity+1 >= ZZZ_CHAR_BANNER_RARE_SOFT_PITY {
		rareChance += ZZZ_RARE_CHAR_BANNER_BASE_CHANCE * 10 * float64(s.CurrRarePity+1-(ZZZ_CHAR_BANNER_RARE_SOFT_PITY-1))
	}
	return s.MihoyoRoller.roll(srChance, rareChance, 0.5, 0.5, StandardZenlessSRChars, ThreeRateUpRares)
}

type ZenlessEngineRoller struct {
	MihoyoRoller
}

var StandardZenlessSREngines = []Rollable{
	{Name: "Standard 5*", Type: SuperRare, Rarity: 5, IsRateUp: false},
}

func (s *ZenlessEngineRoller) Roll() Rollable {
	srChance := ZZZ_SR_ENGINE_BANNER_BASE_CHANCE
	if s.CurrSRPity+1 >= ZZZ_ENGINE_SOFT_PITY {
		srChance += ZZZ_SR_ENGINE_BANNER_BASE_CHANCE * 10 * float64(s.CurrSRPity+1-(ZZZ_ENGINE_SOFT_PITY-1))
	}
	rareChance := ZZZ_RARE_ENGINE_BANNER_BASE_CHANCE
	if s.CurrRarePity+1 >= ZZZ_ENGINE_BANNER_RARE_SOFT_PITY {
		rareChance += ZZZ_RARE_ENGINE_BANNER_BASE_CHANCE * 10 * float64(s.CurrRarePity+1-(ZZZ_ENGINE_BANNER_RARE_SOFT_PITY-1))
	}
	return s.MihoyoRoller.roll(srChance, rareChance, 0.75, 0.75, StandardZenlessSREngines, ThreeRateUpRares)
}

// --- Star Rail ---

const HSR_RARE_CHAR_BANNER_BASE_CHANCE = 0.051
const HSR_RARE_LIGHT_CONE_BANNER_BASE_CHANCE = 0.066
const HSR_SR_LIGHT_CONE_BANNER_BASE_CHANCE = 0.008

type StarRailCharRoller struct {
	MihoyoRoller
}

var StandardStarRailSRChars = []Rollable{
	{Name: "RateUp 5*", Type: SuperRare, Rarity: 5, IsRateUp: true},
	{Name: "Standard 5*", Type: SuperRare, Rarity: 5, IsRateUp: false},
	{Name: "Standard 5*", Type: SuperRare, Rarity: 5, IsRateUp: false},
	{Name: "Standard 5*", Type: SuperRare, Rarity: 5, IsRateUp: false},
	{Name: "Standard 5*", Type: SuperRare, Rarity: 5, IsRateUp: false},
	{Name: "Standard 5*", Type: SuperRare, Rarity: 5, IsRateUp: false},
	{Name: "Standard 5*", Type: SuperRare, Rarity: 5, IsRateUp: false},
	{Name: "Standard 5*", Type: SuperRare, Rarity: 5, IsRateUp: false},
}

func (s *StarRailCharRoller) Roll() Rollable {
	srChance := MIHOYO_CHAR_BANNER_BASE_CHANCE
	if s.CurrSRPity+1 >= 74 {
		srChance += MIHOYO_CHAR_BANNER_BASE_CHANCE * 10 * float64(s.CurrSRPity+1-73)
	}
	rareChance := HSR_RARE_CHAR_BANNER_BASE_CHANCE
	if s.CurrRarePity+1 >= 9 {
		rareChance += HSR_RARE_CHAR_BANNER_BASE_CHANCE * 10 * float64(s.CurrRarePity+1-8)
	}
	return s.MihoyoRoller.roll(srChance, rareChance, 0.50, 0.5, StandardStarRailSRChars, ThreeRateUpRares)
}

type StarRailLCRoller struct {
	MihoyoRoller
}

var StandardStarRailSRLCs = []Rollable{
	{Name: "RateUp 5*", Type: SuperRare, Rarity: 5, IsRateUp: true},
	{Name: "Standard 5*", Type: SuperRare, Rarity: 5, IsRateUp: false},
	{Name: "Standard 5*", Type: SuperRare, Rarity: 5, IsRateUp: false},
	{Name: "Standard 5*", Type: SuperRare, Rarity: 5, IsRateUp: false},
	{Name: "Standard 5*", Type: SuperRare, Rarity: 5, IsRateUp: false},
	{Name: "Standard 5*", Type: SuperRare, Rarity: 5, IsRateUp: false},
	{Name: "Standard 5*", Type: SuperRare, Rarity: 5, IsRateUp: false},
	{Name: "Standard 5*", Type: SuperRare, Rarity: 5, IsRateUp: false},
}

func (s *StarRailLCRoller) Roll() Rollable {
	srChance := HSR_SR_LIGHT_CONE_BANNER_BASE_CHANCE
	if s.CurrSRPity+1 >= 66 {
		srChance += HSR_SR_LIGHT_CONE_BANNER_BASE_CHANCE * 10 * float64(s.CurrSRPity+1-65)
	}
	rareChance := HSR_RARE_LIGHT_CONE_BANNER_BASE_CHANCE
	if s.CurrRarePity+1 >= MIHOYO_RARE_WEAPON_BANNER_SOFT_PITY {
		rareChance += HSR_RARE_LIGHT_CONE_BANNER_BASE_CHANCE * 10 * float64(s.CurrRarePity+1-(MIHOYO_RARE_WEAPON_BANNER_SOFT_PITY-1))
	}
	return s.MihoyoRoller.roll(srChance, rareChance, 0.75, 0.75, StandardStarRailSRLCs, ThreeRateUpRares)
}
