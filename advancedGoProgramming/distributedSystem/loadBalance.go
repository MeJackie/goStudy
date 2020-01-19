package distributedSystem

import "math/rand"

func Shuffle(slice []int)  {
	for i := len(slice); i > 0; i-- {
		lastIdx := i - 1
		idx := rand.Intn(i)
		slice[idx], slice[lastIdx] = slice[lastIdx], slice[idx]
	}
}

