package main

import "testing"

var (
	persegi            Persegi = Persegi{6}
	luasSeharusnya     float64 = 36
	kelilingSeharusnya float64 = 24
)

func TestHitungLuas(t *testing.T) {
	t.Logf("Luas : %.2f", persegi.Luas())

	if persegi.Luas() != luasSeharusnya {
		t.Errorf("SALAH! Seharusnya %.2f", luasSeharusnya)
	}
}

func TestHitungKeliling(t *testing.T) {
	t.Logf("Keliling : %.2f", persegi.Keliling())

	if persegi.Keliling() != kelilingSeharusnya {
		t.Errorf("SALAH! Seharusnya %.2f", kelilingSeharusnya)
	}
}
