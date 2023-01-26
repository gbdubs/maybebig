package maybebig

import "math/big"

func (f *Float) Float64() float64 {
	// TODO(grady) something with precision here
	return f.simple
}

func bigToFloat64(f *big.Float) float64 {
	r, _ := f.Float64()
	return r
}
