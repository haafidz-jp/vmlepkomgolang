package main

type Persegi struct {
	Sisi float64
}

func (k Persegi) Luas() float64 {
	return k.Sisi * k.Sisi
}
func (k Persegi) Keliling() float64 {
	return k.Sisi * 4
}
