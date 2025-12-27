// Copyright IBM Corp. 2016, 2025
// SPDX-License-Identifier: MPL-2.0

package shamirssecrets

import (
	"bytes"
	"testing"

	"github.com/hashicorp/vault/shamir"
)

func TestSplit_invalid(t *testing.T) {
	secret := []byte("test")

	if _, err := Split(secret, 0, 0); err == nil {
		t.Fatalf("expect error")
	}

	if _, err := Split(secret, 2, 3); err == nil {
		t.Fatalf("expect error")
	}

	if _, err := Split(secret, 1000, 3); err == nil {
		t.Fatalf("expect error")
	}

	if _, err := Split(secret, 10, 1); err == nil {
		t.Fatalf("expect error")
	}

	if _, err := Split(nil, 3, 2); err == nil {
		t.Fatalf("expect error")
	}
}

func TestSplit(t *testing.T) {
	secret := []byte("test")

	out, err := Split(secret, 5, 3)
	if err != nil {
		t.Fatalf("err: %v", err)
	}

	if len(out) != 5 {
		t.Fatalf("bad: %v", out)
	}

	for _, share := range out {
		if len(share) != len(secret)+1 {
			t.Fatalf("bad: %v", out)
		}
	}
}

func TestCombine_invalid(t *testing.T) {
	// Not enough parts
	if _, err := Combine(nil); err == nil {
		t.Fatalf("should err")
	}

	// Mis-match in length
	parts := [][]byte{
		[]byte("foo"),
		[]byte("ba"),
	}
	if _, err := Combine(parts); err == nil {
		t.Fatalf("should err")
	}

	// Too short
	parts = [][]byte{
		[]byte("f"),
		[]byte("b"),
	}
	if _, err := Combine(parts); err == nil {
		t.Fatalf("should err")
	}

	parts = [][]byte{
		[]byte("foo"),
		[]byte("foo"),
	}
	if _, err := Combine(parts); err == nil {
		t.Fatalf("should err")
	}
}

func TestCombine(t *testing.T) {
	secret := []byte("test")

	out, err := Split(secret, 5, 3)
	if err != nil {
		t.Fatalf("err: %v", err)
	}

	// There is 5*4*3 possible choices,
	// we will just brute force try them all
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if j == i {
				continue
			}
			for k := 0; k < 5; k++ {
				if k == i || k == j {
					continue
				}
				parts := [][]byte{out[i], out[j], out[k]}
				recomb, err := Combine(parts)
				if err != nil {
					t.Fatalf("err: %v", err)
				}

				if !bytes.Equal(recomb, secret) {
					t.Errorf("parts: (i:%d, j:%d, k:%d) %v", i, j, k, parts)
					t.Fatalf("bad: %v %v", recomb, secret)
				}
			}
		}
	}
}

func BenchmarkCombineMine(b *testing.B) {
	for b.Loop() {
		secret := []byte("test")

		out, err := Split(secret, 5, 3)
		if err != nil {
			b.Fatalf("err: %v", err)
		}

		// There is 5*4*3 possible choices,
		// we will just brute force try them all
		for i := 0; i < 5; i++ {
			for j := 0; j < 5; j++ {
				if j == i {
					continue
				}
				for k := 0; k < 5; k++ {
					if k == i || k == j {
						continue
					}
					parts := [][]byte{out[i], out[j], out[k]}
					recomb, err := Combine(parts)
					if err != nil {
						b.Fatalf("err: %v", err)
					}

					if !bytes.Equal(recomb, secret) {
						b.Errorf("parts: (i:%d, j:%d, k:%d) %v", i, j, k, parts)
						b.Fatalf("bad: %v %v", recomb, secret)
					}
				}
			}
		}
	}
}

func BenchmarkCombineVaults(b *testing.B) {
	for b.Loop() {
		secret := []byte("test")

		out, err := shamir.Split(secret, 5, 3)
		if err != nil {
			b.Fatalf("err: %v", err)
		}

		// There is 5*4*3 possible choices,
		// we will just brute force try them all
		for i := 0; i < 5; i++ {
			for j := 0; j < 5; j++ {
				if j == i {
					continue
				}
				for k := 0; k < 5; k++ {
					if k == i || k == j {
						continue
					}
					parts := [][]byte{out[i], out[j], out[k]}
					recomb, err := shamir.Combine(parts)
					if err != nil {
						b.Fatalf("err: %v", err)
					}

					if !bytes.Equal(recomb, secret) {
						b.Errorf("parts: (i:%d, j:%d, k:%d) %v", i, j, k, parts)
						b.Fatalf("bad: %v %v", recomb, secret)
					}
				}
			}
		}
	}
}

func TestField_Add(t *testing.T) {
	if out := add(16, 16); out != 0 {
		t.Fatalf("Bad: %v 16", out)
	}

	if out := add(3, 4); out != 7 {
		t.Fatalf("Bad: %v 7", out)
	}
}

func TestField_Mult(t *testing.T) {
	if out := mult(3, 7); out != 9 {
		t.Fatalf("Bad: %v 9", out)
	}

	if out := mult(3, 0); out != 0 {
		t.Fatalf("Bad: %v 0", out)
	}

	if out := mult(0, 3); out != 0 {
		t.Fatalf("Bad: %v 0", out)
	}
}

func TestField_Divide(t *testing.T) {
	if out := div(0, 7); out != 0 {
		t.Fatalf("Bad: %v 0", out)
	}

	if out := div(3, 3); out != 1 {
		t.Fatalf("Bad: %v 1", out)
	}

	if out := div(6, 3); out != 2 {
		t.Fatalf("Bad: %v 2", out)
	}
}
