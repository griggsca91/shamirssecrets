package shamirssecrets

import (
	"crypto/rand"
	"math"
	mathrand "math/rand/v2"
)

func Split(secret []byte, shares, threshold int) (parts [][]byte) {
	out := make([][]byte, shares)

	xCoordinates := mathrand.Perm(255) // generating a slice of x points

	for i := range shares {
		out[i] = make([]byte, len(secret)+1)
		// the x coordinates that we'll use to calculate
		out[i][len(secret)] = uint8(xCoordinates[i]) + 1
	}

	// in f(x) = x^3 + x^2 + x + b
	// b is the value from the secret
	for idx := range secret {
		for i, x := range xCoordinates {
			coefficients := make([]uint8, threshold-1)
			// this is the constant of the polynomial being set
			rand.Read(coefficients)

			// for c := range coefficients {
			// }
			y := uint8(x)
			out[i][idx] = y
		}
	}

	return out
}

func Combine(parts []int) {
}

// f(x)=1234+166x+94x^{2}
// 1 == 1494
func evaluatePolynomial(x int, coefficients []int) int {
	result := 0
	for i, c := range coefficients {
		result += c * int(math.Pow(float64(x), float64(i)))
	}
	return result
}
