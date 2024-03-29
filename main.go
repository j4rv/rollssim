package rollssim

const SuperRare = "Super Rare"
const Rare = "Rare"
const Fodder = "Fodder"
const ChosenWeapon = "Chosen Weapon"
const NotChosenWeapon = "Not Chosen Weapon"

var ThreeRateUpRares = []string{"Rare 1", "Rare 2", "Rare 3"}
var FiveRateUpRares = []string{"Rare 1", "Rare 2", "Rare 3", "Rare 4", "Rare 5"}

type Rollable struct {
	Name     string
	Type     string
	Rarity   int
	IsRateUp bool
}

type Roller interface {
	Roll() Rollable
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

			switch lc.Name {
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
