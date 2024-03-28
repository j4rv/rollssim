package rollssim

const RateUpSR = "Rate Up SR"
const StandardSR = "Standard SR"
const RateUpRare = "Rate Up Rare"
const StandardRare = "Standard Rare"
const Fodder = "Fodder"

var ThreeRateUpRares = []string{"Rate Up Rare 1", "Rate Up Rare 2", "Rate Up Rare 3"}

type Rollable struct {
	Name   string
	Type   string
	Rarity int
}

type Roller interface {
	Roll() Rollable
	MultiRoll(int) []Rollable
}

type WantedRollsResult struct {
	RateUpSRCharCount int
	StdSRCharCount    int
	FodderCharCount   int
	RateUpSRLCCount   int
	StdLCSRCount      int
	FodderLCCount     int
}

func CalcWantedRolls(rollCount, wantedCharCount, wantedLCCount int, chars, lcs Roller) WantedRollsResult {
	result := WantedRollsResult{}
	for i := 0; i < rollCount; i++ {
		if result.RateUpSRCharCount < wantedCharCount {
			c := chars.Roll()
			if c.Type == RateUpSR {
				result.RateUpSRCharCount++
			} else if c.Type == StandardSR {
				result.StdSRCharCount++
			} else {
				result.FodderCharCount++
			}
		} else if result.RateUpSRLCCount < wantedLCCount {
			lc := lcs.Roll()
			if lc.Type == RateUpSR {
				result.RateUpSRLCCount++
			} else if lc.Type == StandardSR {
				result.StdLCSRCount++
			} else {
				result.FodderLCCount++
			}
		} else {
			break
		}
	}
	return result
}
