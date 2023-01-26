package maybebig

import (
	"fmt"
	"math/big"
)

type Float struct {
	simple   float64
	big      *big.Float
	bigFn    *bigFn
	lossyOps uint
}

func (f *Float) String() string {
	b := "nil"
	if f.big != nil {
		b = fmt.Sprintf("%f", bigToFloat64(f.big))
	}
	return fmt.Sprintf(`F{simple:%f big:%s lops:%d}`, f.simple, b, f.lossyOps)
}

type bigFn struct {
	op   floatOperation
	args []*Float
}

func (f *Float) GetBig() *big.Float {
	if f.big == nil {
		f.big = f.bigFn.Eval()
		f.bigFn = nil
		if !doubleCheckMode {
			f.simple = bigToFloat64(f.big)
			f.lossyOps = 0
		}
	}
	return f.big
}

func (f *Float) hardRecompute() {
	if f.big == nil && doubleCheckMode {
		recordHardRecompute()
	}
	b := f.GetBig()
	f.simple = bigToFloat64(b)
	f.lossyOps = 0
}

func newFloatBig() *big.Float {
	b := big.NewFloat(0)
	b.SetPrec(GetPrecision())
	b.SetMode(big.ToZero)
	return b
}

func newFloatBigF(f float64) *big.Float {
	b := big.NewFloat(f)
	b.SetPrec(GetPrecision())
	b.SetMode(big.ToZero)
	return b
}

func NewFloat(f float64) *Float {
	return &Float{
		simple:   f,
		big:      newFloatBigF(f),
		lossyOps: 1,
	}
}

func NewFloatZ() *Float {
	return &Float{
		simple:   0,
		big:      newFloatBig(),
		lossyOps: 0,
	}
}

func (a *Float) Set(b *Float) *Float {
	a.simple = b.simple
	a.big = b.big
	a.lossyOps = b.lossyOps
	a.bigFn = b.bigFn
	return a
}

func (f *Float) SetInt64(i int64) *Float {
	f.simple = float64(i)
	f.big = newFloatBig().SetInt64(i)
	f.bigFn = nil
	f.lossyOps = 1
	return f
}

func (f *Float) SetFloat64(f64 float64) *Float {
	f.simple = f64
	f.big = newFloatBig().SetFloat64(f64)
	f.bigFn = nil
	f.lossyOps = 1
	return f
}
