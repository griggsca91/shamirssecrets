package shamirssecrets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvaluatePolynomial(t *testing.T) {
	assert.Equal(t, evaluatePolynomial(1, []int{1234, 166, 94}), 1494)
	assert.Equal(t, evaluatePolynomial(2, []int{1234, 166, 94}), 1942)
}

func TestSplit(t *testing.T) {
	// secret := []byte("1234")
	// shares := 6
	// threshold := 3

	// splits := Split(secret, shares, threshold)
	// assert.Equal(t, len(splits), shares)
	// fmt.Println(splits)
}

func TestCombine(t *testing.T) {
}
