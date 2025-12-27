package shamirssecrets

func GaloisAdd(a, b uint8) uint8 {
	return a ^ b
}

func GaloisMult(a, b uint8) uint8 {
	p := uint8(0)

	for range 8 {
		if b&1 == 1 {
			p = p ^ a
		}

		hiBitSet := (a & 0x80)
		a <<= 1

		if hiBitSet == 0x80 {
			a ^= 0x1b
		}

		b >>= 1
	}

	return p
}

// mult multiplies two numbers in GF(2^8)
// copied from https://github.com/hashicorp/vault/blob/34e38af5f08c9c3a8a353314d6f5217c385f0e92/shamir/shamir.go#L116
func VaultMult(a, b uint8) (out uint8) {
	var r uint8 = 0
	var i uint8 = 8

	for i > 0 {
		i--
		r = (-(b >> i & 1) & a) ^ (-(r >> 7) & 0x1B) ^ (r + r)
	}

	return r
}
