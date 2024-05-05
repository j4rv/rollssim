package rollssim

import "math/rand"

func RandomRollable(r []Rollable) Rollable {
	if len(r) == 0 {
		return Rollable{}
	}
	randomIndex := rand.Intn(len(r))
	return r[randomIndex]
}
