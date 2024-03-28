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

	WeaponBannerRateUpSRCount        int
	WeaponBannerStdSRCount           int
	WeaponBannerRateUpRareCount      int
	WeaponBannerStdRareCount         int
	WeaponBannerFodderCount          int
	WeaponBannerChosenRateUpCount    int
	WeaponBannerNotChosenRateUpCount int
}

func CalcWantedRolls(rollCount, wantedCharCount, wantedLCCount int, chars, lcs Roller) WantedRollsResult {
	result := WantedRollsResult{}
	for i := 0; i < rollCount; i++ {
		if result.CharacterBannerRateUpSRCount < wantedCharCount {
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
		} else if result.WeaponBannerRateUpSRCount < wantedLCCount {
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
