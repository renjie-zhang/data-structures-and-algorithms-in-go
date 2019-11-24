package PrimeNumber

import (
	"math"
)

func Prime(number int) int {
	for i := 2; i <= int(math.Sqrt(float64(number))); i++ {
		if number%i == 0 {
			return 0
		}
	}
	return 1
}
