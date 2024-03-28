package rollssim

import "math/rand"

func RandomString(strings []string) string {
	if len(strings) == 0 {
		return ""
	}
	randomIndex := rand.Intn(len(strings))
	return strings[randomIndex]
}
