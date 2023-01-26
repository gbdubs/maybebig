package maybebig

import (
	"math"
	"math/big"
)

func Pi() *Float {
	return FromFloat64(math.Pi)
}

func sinBig(b *big.Float) *big.Float {
	return newFloatBig().SetFloat64(math.Sin(bigToFloat64(b))) // TODO(grady) better trig
}

func Sin(a *Float) *Float {
	return &Float{
		simple:   math.Sin(a.simple),
		lossyOps: a.lossyOps + 1,
		bigFn: &bigFn{
			op:   floatOperationSin,
			args: []*Float{a},
		},
		big: nil,
	}
}

func cosBig(b *big.Float) *big.Float {
	return newFloatBig().SetFloat64(math.Cos(bigToFloat64(b))) // TODO(grady) better trig
}

func Cos(a *Float) *Float {
	return &Float{
		simple:   math.Cos(a.simple),
		lossyOps: a.lossyOps + 1,
		bigFn: &bigFn{
			op:   floatOperationCos,
			args: []*Float{a},
		},
		big: nil,
	}
}

func atanBig(b *big.Float) *big.Float {
	return newFloatBig().SetFloat64(math.Atan(bigToFloat64(b))) // TODO(grady) better trig
}

func Atan(a *Float) *Float {
	return &Float{
		simple:   math.Atan(a.simple),
		lossyOps: a.lossyOps + 1,
		bigFn: &bigFn{
			op:   floatOperationAtan,
			args: []*Float{a},
		},
		big: nil,
	}
}

func atan2Big(y, x *big.Float) *big.Float {
	return newFloatBig().SetFloat64(math.Atan2(bigToFloat64(y), bigToFloat64(x))) // TODO(grady) better trig
}

func Atan2(y, x *Float) *Float {
	return &Float{
		simple:   math.Atan2(y.simple, x.simple),
		lossyOps: y.lossyOps + x.lossyOps + 1,
		bigFn: &bigFn{
			op:   floatOperationAtan2,
			args: []*Float{y, x},
		},
		big: nil,
	}
}
