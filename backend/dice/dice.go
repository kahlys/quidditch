package dice

import "math/rand"

func Roll(stat int, stamina int, nb int) bool {
	for i := 0; i < nb; i++ {
		if rand.Intn(101) > stat-(100-stamina) {
			return false
		}
	}
	return true
}
