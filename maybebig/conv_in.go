package maybebig

import (
	"math/big"
)

func FromInt(i int) *Float {
	return NewFloatZ().SetInt64(int64(i))
}

func FromFloat64(f float64) *Float {
	return NewFloatZ().SetFloat64(f)
}

func FromBig(b *big.Float) *Float {
	return &Float{
		simple:   bigToFloat64(b),
		big:      b,
		lossyOps: 0,
	}
}
