package rollssim

const SuperRare = "Super Rare"
const Rare = "Rare"
const Fodder = "Fodder"
const ChosenWeapon = "Chosen Weapon"
const NotChosenWeapon = "Not Chosen Weapon"

var TwoRateUpRares = []Rollable{
	{Tag: "RateUp Rare 1", Type: Rare, Rarity: 4, IsRateUp: true},
	{Tag: "RateUp Rare 2", Type: Rare, Rarity: 4, IsRateUp: true},
}

var ThreeRateUpRares = []Rollable{
	{Tag: "RateUp Rare 1", Type: Rare, Rarity: 4, IsRateUp: true},
	{Tag: "RateUp Rare 2", Type: Rare, Rarity: 4, IsRateUp: true},
	{Tag: "RateUp Rare 3", Type: Rare, Rarity: 4, IsRateUp: true},
}

var FiveRateUpRares = []Rollable{
	{Tag: "RateUp Rare 1", Type: Rare, Rarity: 4, IsRateUp: true},
	{Tag: "RateUp Rare 2", Type: Rare, Rarity: 4, IsRateUp: true},
	{Tag: "RateUp Rare 3", Type: Rare, Rarity: 4, IsRateUp: true},
	{Tag: "RateUp Rare 4", Type: Rare, Rarity: 4, IsRateUp: true},
	{Tag: "RateUp Rare 5", Type: Rare, Rarity: 4, IsRateUp: true},
}

type Rollable struct {
	Tag      string
	Type     string
	Rarity   int
	IsRateUp bool
}

type Roller interface {
	Roll() Rollable
	Reset()
}

type WantedRollsResult struct {
	CharacterBannerRateUpSRCount   int
	CharacterBannerStdSRCount      int
	CharacterBannerRateUpRareCount int
	CharacterBannerStdRareCount    int
	CharacterBannerFodderCount     int
	CharacterBannerRollCount       int

	WeaponBannerRateUpSRCount        int
	WeaponBannerStdSRCount           int
	WeaponBannerRateUpRareCount      int
	WeaponBannerStdRareCount         int
	WeaponBannerFodderCount          int
	WeaponBannerChosenRateUpCount    int
	WeaponBannerNotChosenRateUpCount int
	WeaponBannerRollCount            int
}

func (r *WantedRollsResult) Add(r2 WantedRollsResult) {
	r.CharacterBannerRateUpSRCount += r2.CharacterBannerRateUpSRCount
	r.CharacterBannerStdSRCount += r2.CharacterBannerStdSRCount
	r.CharacterBannerRateUpRareCount += r2.CharacterBannerRateUpRareCount
	r.CharacterBannerStdRareCount += r2.CharacterBannerStdRareCount
	r.CharacterBannerFodderCount += r2.CharacterBannerFodderCount
	r.CharacterBannerRollCount += r2.CharacterBannerRollCount

	r.WeaponBannerRateUpSRCount += r2.WeaponBannerRateUpSRCount
	r.WeaponBannerStdSRCount += r2.WeaponBannerStdSRCount
	r.WeaponBannerRateUpRareCount += r2.WeaponBannerRateUpRareCount
	r.WeaponBannerStdRareCount += r2.WeaponBannerStdRareCount
	r.WeaponBannerFodderCount += r2.WeaponBannerFodderCount
	r.WeaponBannerChosenRateUpCount += r2.WeaponBannerChosenRateUpCount
	r.WeaponBannerNotChosenRateUpCount += r2.WeaponBannerNotChosenRateUpCount
	r.WeaponBannerRollCount += r2.WeaponBannerRollCount
}

func CalcCharacterBannerRolls(rollCount int, chars Roller) WantedRollsResult {
	result := WantedRollsResult{}
	for i := 0; i < rollCount; i++ {
		c := chars.Roll()
		switch c.Type {
		case SuperRare:
			if c.IsRateUp {
				result.CharacterBannerRateUpSRCount++
			} else {
				result.CharacterBannerStdSRCount++
			}
		case Rare:
			if c.IsRateUp {
				result.CharacterBannerRateUpRareCount++
			} else {
				result.CharacterBannerStdRareCount++
			}
		default:
			result.CharacterBannerFodderCount++
		}
	}
	return result
}

func CalcWeaponBannerRolls(rollCount int, weapons Roller) WantedRollsResult {
	result := WantedRollsResult{}
	for i := 0; i < rollCount; i++ {
		w := weapons.Roll()
		switch w.Type {
		case SuperRare:
			if w.IsRateUp {
				result.WeaponBannerRateUpSRCount++
				// Only for genshin
				switch w.Tag {
				case ChosenWeapon:
					result.WeaponBannerChosenRateUpCount++
				case NotChosenWeapon:
					result.WeaponBannerNotChosenRateUpCount++
				}
			} else {
				result.WeaponBannerStdSRCount++
			}
		case Rare:
			if w.IsRateUp {
				result.WeaponBannerRateUpRareCount++
			} else {
				result.WeaponBannerStdRareCount++
			}
		default:
			result.WeaponBannerFodderCount++
		}
	}
	return result
}

func CalcGenshinWantedRolls(rollCount, wantedCharCount, chosenWeaponCount int, chars *GenshinCharRoller, weapons *GenshinWeaponRoller) WantedRollsResult {
	result := WantedRollsResult{}
	for i := 0; i < rollCount; i++ {
		if result.CharacterBannerRateUpSRCount < wantedCharCount {
			result.CharacterBannerRollCount++
			c := chars.Roll()
			switch c.Type {
			case SuperRare:
				if c.IsRateUp {
					result.CharacterBannerRateUpSRCount++
				} else {
					result.CharacterBannerStdSRCount++
				}
			case Rare:
				if c.IsRateUp {
					result.CharacterBannerRateUpRareCount++
				} else {
					result.CharacterBannerStdRareCount++
				}
			default:
				result.CharacterBannerFodderCount++
			}
		} else if result.WeaponBannerChosenRateUpCount < chosenWeaponCount {
			result.WeaponBannerRollCount++
			lc := weapons.Roll()

			switch lc.Type {
			case SuperRare:
				if lc.IsRateUp {
					result.WeaponBannerRateUpSRCount++
				} else {
					result.WeaponBannerStdSRCount++
				}
			case Rare:
				if lc.IsRateUp {
					result.WeaponBannerRateUpRareCount++
				} else {
					result.WeaponBannerStdRareCount++
				}
			default:
				result.WeaponBannerFodderCount++
			}

			switch lc.Tag {
			case ChosenWeapon:
				result.WeaponBannerChosenRateUpCount++
			case NotChosenWeapon:
				result.WeaponBannerNotChosenRateUpCount++
			}
		} else {
			break
		}
	}

	return result
}

func CalcZenlessWantedRolls(rollCount, wantedCharCount, wantedWeaponCount int, chars *GenshinCharRoller, weapons *ZenlessEngineRoller) WantedRollsResult {
	result := WantedRollsResult{}
	for i := 0; i < rollCount; i++ {
		if result.CharacterBannerRateUpSRCount < wantedCharCount {
			result.CharacterBannerRollCount++
			c := chars.Roll()
			switch c.Type {
			case SuperRare:
				if c.IsRateUp {
					result.CharacterBannerRateUpSRCount++
				} else {
					result.CharacterBannerStdSRCount++
				}
			case Rare:
				if c.IsRateUp {
					result.CharacterBannerRateUpRareCount++
				} else {
					result.CharacterBannerStdRareCount++
				}
			default:
				result.CharacterBannerFodderCount++
			}
		} else if result.WeaponBannerRateUpSRCount < wantedWeaponCount {
			result.WeaponBannerRollCount++
			lc := weapons.Roll()

			switch lc.Type {
			case SuperRare:
				if lc.IsRateUp {
					result.WeaponBannerRateUpSRCount++
				} else {
					result.WeaponBannerStdSRCount++
				}
			case Rare:
				if lc.IsRateUp {
					result.WeaponBannerRateUpRareCount++
				} else {
					result.WeaponBannerStdRareCount++
				}
			default:
				result.WeaponBannerFodderCount++
			}
		} else {
			break
		}
	}

	return result
}

func CalcStarRailWantedRolls(rollCount, wantedCharCount, rateUpLCCount int, chars *StarRailCharRoller, lcs *StarRailLCRoller) WantedRollsResult {
	result := WantedRollsResult{}
	for i := 0; i < rollCount; i++ {
		if result.CharacterBannerRateUpSRCount < wantedCharCount {
			result.CharacterBannerRollCount++
			c := chars.Roll()
			switch c.Type {
			case SuperRare:
				if c.IsRateUp {
					result.CharacterBannerRateUpSRCount++
				} else {
					result.CharacterBannerStdSRCount++
				}
			case Rare:
				if c.IsRateUp {
					result.CharacterBannerRateUpRareCount++
				} else {
					result.CharacterBannerStdRareCount++
				}
			default:
				result.CharacterBannerFodderCount++
			}
		} else if result.WeaponBannerRateUpSRCount < rateUpLCCount {
			result.WeaponBannerRollCount++
			lc := lcs.Roll()

			switch lc.Type {
			case SuperRare:
				if lc.IsRateUp {
					result.WeaponBannerRateUpSRCount++
				} else {
					result.WeaponBannerStdSRCount++
				}
			case Rare:
				if lc.IsRateUp {
					result.WeaponBannerRateUpRareCount++
				} else {
					result.WeaponBannerStdRareCount++
				}
			default:
				result.WeaponBannerFodderCount++
			}
		} else {
			break
		}
	}

	return result
}
