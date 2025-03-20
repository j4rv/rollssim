package rollssim

const REVERSE_BANNER_BASE_CHANCE = 0.015
const REVERSE_CHAR_BANNER_SOFT_PITY = 61
const REVERSE_BANNER_RARE_BASE_CHANCE = 0.085

var ReverseStandardSRChars = []Rollable{
	{Name: "Standard 5*", Type: SuperRare, Rarity: 5, IsRateUp: false},
}

type ReverseCharRoller struct {
	MihoyoRoller
}

// 6* are returned as Rarity 5 for simplicity, 5* are returned as Rarity 4
func (s *ReverseCharRoller) Roll() Rollable {
	srChance := REVERSE_BANNER_BASE_CHANCE
	if s.CurrSRPity+1 >= REVERSE_CHAR_BANNER_SOFT_PITY {
		srChance += 0.04
	}
	if s.CurrSRPity+1 > REVERSE_CHAR_BANNER_SOFT_PITY {
		srChance += 0.025 * float64(s.CurrSRPity+1-REVERSE_CHAR_BANNER_SOFT_PITY)
	}
	rareChance := REVERSE_BANNER_RARE_BASE_CHANCE
	return s.MihoyoRoller.roll(srChance, rareChance, 0.5, 0.5, 0, ReverseStandardSRChars, TwoRateUpRares)
}
