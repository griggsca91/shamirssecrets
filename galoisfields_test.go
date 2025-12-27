package shamirssecrets

import (
	"math/rand/v2"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGaloisAdd(t *testing.T) {
	assert.Equal(t, GaloisAdd(16, 16), uint8(0))
}

func TestGaloisMult(t *testing.T) {
	assert.Equal(t, GaloisMult(3, 7), uint8(9))
}

func TestCompareVaultWithMineMult(t *testing.T) {
	for range 20 {
		a := rand.N[uint8](255)
		b := rand.N[uint8](255)
		assert.Equal(t,
			GaloisMult(a, b),
			VaultMult(a, b),
		)
	}
}

func BenchmarkMultGalois(b *testing.B) {
	for b.Loop() {
		GaloisMult(rand.N[uint8](255), rand.N[uint8](255))
	}
}

func BenchmarkMultVault(b *testing.B) {
	for b.Loop() {
		VaultMult(rand.N[uint8](255), rand.N[uint8](255))
	}
}
