package rollssim

const RateUpSR = "Rate Up SR"
const ChosenRateUpSR = "Chosen Rate Up SR"
const NotChosenRateUpSR = "Not Chosen Rate Up SR"
const StandardSR = "Standard SR"
const RateUpRare = "Rate Up Rare"
const StandardRare = "Standard Rare"
const Fodder = "Fodder"

var ThreeRateUpRares = []string{"Rate Up Rare 1", "Rate Up Rare 2", "Rate Up Rare 3"}
var FiveRateUpRares = []string{"Rate Up Rare 1", "Rate Up Rare 2", "Rate Up Rare 3", "Rate Up Rare 4", "Rate Up Rare 5"}

type Rollable struct {
	Name   string
	Type   string
	Rarity int
}

type Roller interface {
	Roll() Rollable
}

type WantedRollsResult struct {
	RateUpSRCharCount int
	StdSRCharCount    int
	FodderCharCount   int

	RateUpSRWeaponCount int
	StdSRWeaponCount    int
	FodderWeaponCount   int

	ChosenSRWeaponCount    int
	NotChosenSRWeaponCount int
}

func CalcWantedRolls(rollCount, wantedCharCount, wantedLCCount int, chars, lcs Roller) WantedRollsResult {
	result := WantedRollsResult{}
	for i := 0; i < rollCount; i++ {
		if result.RateUpSRCharCount < wantedCharCount {
			c := chars.Roll()
			switch c.Type {
			case RateUpSR:
				result.RateUpSRCharCount++
			case StandardSR:
				result.StdSRCharCount++
			default:
				result.FodderCharCount++
			}
		} else if result.RateUpSRWeaponCount < wantedLCCount {
			lc := lcs.Roll()

			switch lc.Type {
			case RateUpSR:
				result.RateUpSRWeaponCount++
			case StandardSR:
				result.StdSRWeaponCount++
			default:
				result.FodderWeaponCount++
			}

			switch lc.Name {
			case ChosenRateUpSR:
				result.ChosenSRWeaponCount++
			case NotChosenRateUpSR:
				result.NotChosenSRWeaponCount++
			}
		} else {
			break
		}
	}
	return result
}
