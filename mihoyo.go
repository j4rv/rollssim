package rollssim

import (
	"math/rand"
)

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
			return Rollable{Name: RateUpSR, Type: RateUpSR, Rarity: 5}
		} else {
			s.GuaranteedRateUpSR = true
			return Rollable{Name: RandomString(rateUpSRItems), Type: StandardSR, Rarity: 5}
		}
	}

	// Check if we get a Rare
	if rand.Float64() <= rareChance {
		s.CurrRarePity = 0
		if s.GuaranteedRateUpRare || rand.Float64() <= 0.5 {
			s.GuaranteedRateUpRare = false
			return Rollable{Name: RandomString(rateUpRareItems), Type: RateUpRare, Rarity: 4}
		} else {
			s.GuaranteedRateUpRare = true
			return Rollable{Name: StandardRare, Type: StandardRare, Rarity: 4}
		}
	}

	return Rollable{Name: Fodder, Type: Fodder, Rarity: 3}
}

type GenshinCharRoller struct {
	MihoyoRoller
}

var StandardGenshinSRChars = []string{"Diluc", "Jean", "Qiqi", "Keqing", "Mona", "Tighnari", "Dehya"}

func (s *GenshinCharRoller) Roll() Rollable {
	srChance := 0.006
	if s.CurrSRPity+1 >= 74 {
		srChance += 1.0 / (80.0 - 65.0) * float64(s.CurrSRPity+1-73)
	}
	rareChance := 0.051
	if s.CurrRarePity+1 >= 9 {
		rareChance += 0.5 * float64(s.CurrRarePity+1-8)
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

	srChance := 0.008
	if s.CurrSRPity >= 66 {
		srChance += 1.0 / (80.0 - 65.0) * float64(s.CurrSRPity-65)
	}
	rareChance := 0.051
	if s.CurrRarePity >= 9 {
		rareChance += 0.5 * float64(s.CurrRarePity-8)
	}

	// Check if we get a SR
	if rand.Float64() <= srChance {
		s.CurrSRPity = 0
		r := rand.Float64()
		if r <= 0.375 || s.FatePoints >= 2 {
			s.FatePoints = 0
			return Rollable{Name: ChosenRateUpSR, Type: RateUpSR, Rarity: 5}
		} else if r <= 0.75 {
			s.FatePoints += 1
			return Rollable{Name: NotChosenRateUpSR, Type: RateUpSR, Rarity: 5}
		} else {
			s.FatePoints += 1
			return Rollable{Name: RandomString(StandardGenshinSRWeapons), Type: StandardSR, Rarity: 5}
		}
	}

	// Check if we get a Rare
	if rand.Float64() <= rareChance {
		s.CurrRarePity = 0
		if s.GuaranteedRateUpRare || rand.Float64() <= 0.5 {
			s.GuaranteedRateUpRare = false
			return Rollable{Name: RandomString(FiveRateUpRares), Type: RateUpRare, Rarity: 4}
		} else {
			s.GuaranteedRateUpRare = true
			return Rollable{Name: StandardRare, Type: StandardRare, Rarity: 4}
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
	srChance := 0.006
	if s.CurrSRPity+1 >= 74 {
		srChance += 1.0 / (80.0 - 65.0) * float64(s.CurrSRPity+1-73)
	}
	rareChance := 0.051
	if s.CurrRarePity+1 >= 9 {
		rareChance += 0.5 * float64(s.CurrRarePity+1-8)
	}
	return s.MihoyoRoller.roll(srChance, rareChance, 0.5, StandardStarRailSRChars, ThreeRateUpRares)
}

type StarRailLCRoller struct {
	MihoyoRoller
}

var StandardStarRailSRLCs = []string{"Bailu LC", "Yanqing LC", "Clara LC", "Gepard LC", "Bronya LC", "Welt LC", "Himeko LC"}

func (s *StarRailLCRoller) Roll() Rollable {
	srChance := 0.008
	if s.CurrSRPity+1 >= 66 {
		srChance += 1.0 / (80.0 - 65.0) * float64(s.CurrSRPity+1-65)
	}
	rareChance := 0.051
	if s.CurrRarePity+1 >= 9 {
		rareChance += 0.5 * float64(s.CurrRarePity+1-8)
	}
	return s.MihoyoRoller.roll(srChance, rareChance, 0.75, StandardStarRailSRLCs, ThreeRateUpRares)
}
