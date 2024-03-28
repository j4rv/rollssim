package rollssim

import (
	"math/rand"
)

type StarRailCharRoller struct {
	CurrSRPity           int
	CurrRarePity         int
	GuaranteedRateUpSR   bool
	GuaranteedRateUpRare bool
}

var StandardStarRailSRChars = []string{"Bailu", "Yanqing", "Clara", "Gepard", "Bronya", "Welt", "Himeko"}
var StandardStarRailSRLCs = []string{"Bailu LC", "Yanqing LC", "Clara LC", "Gepard LC", "Bronya LC", "Welt LC", "Himeko LC"}

func (s *StarRailCharRoller) Roll() Rollable {
	s.CurrSRPity++
	s.CurrRarePity++

	// Check if we get a SR
	srChance := 0.006
	if s.CurrSRPity >= 74 {
		srChance += 1.0 / 17.0 * float64(s.CurrSRPity-73)
	}
	if rand.Float64() <= srChance {
		s.CurrSRPity = 0
		if s.GuaranteedRateUpSR || rand.Float64() <= 0.5 {
			s.GuaranteedRateUpSR = false
			return Rollable{Name: RateUpSR, Type: RateUpSR, Rarity: 5}
		} else {
			s.GuaranteedRateUpSR = true
			return Rollable{Name: RandomString(StandardStarRailSRChars), Type: StandardSR, Rarity: 5}
		}
	}

	// Check if we get a Rare
	rareChance := 0.051
	if s.CurrRarePity >= 9 {
		rareChance += 1.0 / 2.0 * float64(s.CurrRarePity-8)
	}
	if rand.Float64() <= rareChance {
		s.CurrRarePity = 0
		if s.GuaranteedRateUpRare || rand.Float64() <= 0.5 {
			s.GuaranteedRateUpRare = false
			return Rollable{Name: RandomString(ThreeRateUpRares), Type: RateUpRare, Rarity: 4}
		} else {
			s.GuaranteedRateUpRare = true
			return Rollable{Name: StandardRare, Type: StandardRare, Rarity: 4}
		}
	}

	return Rollable{Type: Fodder, Rarity: 3}
}

func (s *StarRailCharRoller) MultiRoll(n int) []Rollable {
	rolls := make([]Rollable, n)
	for i := 0; i < n; i++ {
		rolls[i] = s.Roll()
	}
	return rolls
}

type StarRailLCRoller struct {
	CurrSRPity           int
	CurrRarePity         int
	GuaranteedRateUpSR   bool
	GuaranteedRateUpRare bool
}

func (s *StarRailLCRoller) Roll() Rollable {
	s.CurrSRPity++
	s.CurrRarePity++

	// Check if we get a SR
	srChance := 0.008
	if s.CurrSRPity >= 66 {
		srChance += 1.0 / 15.0 * float64(s.CurrSRPity-65)
	}
	if rand.Float64() <= srChance {
		s.CurrSRPity = 0
		if s.GuaranteedRateUpSR || rand.Float64() <= 0.75 {
			s.GuaranteedRateUpSR = false
			return Rollable{Name: RateUpSR, Type: RateUpSR, Rarity: 5}
		} else {
			s.GuaranteedRateUpSR = true
			return Rollable{Name: RandomString(StandardStarRailSRLCs), Type: StandardSR, Rarity: 5}
		}
	}

	// Check if we get a Rare
	rareChance := 0.051
	if s.CurrRarePity >= 9 {
		rareChance += 1.0 / 2.0 * float64(s.CurrRarePity-8)
	}
	if rand.Float64() <= rareChance {
		s.CurrRarePity = 0
		if s.GuaranteedRateUpRare || rand.Float64() <= 0.5 {
			s.GuaranteedRateUpRare = false
			return Rollable{Name: RandomString(ThreeRateUpRares), Type: RateUpRare, Rarity: 4}
		} else {
			s.GuaranteedRateUpRare = true
			return Rollable{Name: StandardRare, Type: StandardRare, Rarity: 4}
		}
	}

	return Rollable{Type: Fodder, Rarity: 3}
}

func (s *StarRailLCRoller) MultiRoll(n int) []Rollable {
	rolls := make([]Rollable, n)
	for i := 0; i < n; i++ {
		rolls[i] = s.Roll()
	}
	return rolls
}
