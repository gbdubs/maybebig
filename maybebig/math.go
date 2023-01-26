package maybebig

import (
	"math"
	"math/big"
)

func addBig(fs ...*big.Float) *big.Float {
	r := newFloatBig()
	for _, f := range fs {
		r.Add(r, f)
	}
	return r
}

func Add(fs ...*Float) *Float {
	simple := 0.0
	var lops uint = 0
	for _, f := range fs {
		simple += f.simple
		lops += f.lossyOps
	}
	return &Float{
		simple: simple,
		big:    nil,
		bigFn: &bigFn{
			op:   floatOperationAdd,
			args: fs,
		},
		lossyOps: lops,
	}
}

func mulBig(fs ...*big.Float) *big.Float {
	result := newFloatBig().SetInt64(1)
	for _, f := range fs {
		result = productBig(result, f)
	}
	return result
}

func Mul(fs ...*Float) *Float {
	simple := 1.0
	var lops uint = 0
	for _, f := range fs {
		simple *= f.simple
		lops += f.lossyOps
	}
	return &Float{
		simple: simple,
		big:    nil,
		bigFn: &bigFn{
			op:   floatOperationMul,
			args: fs,
		},
		lossyOps: lops,
	}
}

func productBig(a, b *big.Float) *big.Float {
	return newFloatBig().Mul(a, b)
}

func subBig(a, b *big.Float) *big.Float {
	return newFloatBig().Sub(a, b)
}

func Sub(a, b *Float) *Float {
	return &Float{
		simple:   a.simple - b.simple,
		lossyOps: a.lossyOps + b.lossyOps,
		bigFn: &bigFn{
			op:   floatOperationSub,
			args: []*Float{a, b},
		},
		big: nil,
	}
}

func divBig(a, b *big.Float) *big.Float {
	return newFloatBig().Quo(a, b)
}

func Div(a, b *Float) *Float {
	return &Float{
		simple:   a.simple / b.simple,
		lossyOps: a.lossyOps + b.lossyOps,
		bigFn: &bigFn{
			op:   floatOperationDiv,
			args: []*Float{a, b},
		},
		big: nil,
	}
}

func Sq(a *Float) *Float {
	return Mul(a, a)
}

func sqrtBig(a *big.Float) *big.Float {
	return newFloatBig().Sqrt(a)
}

func Sqrt(a *Float) *Float {
	return &Float{
		simple:   math.Sqrt(a.simple),
		lossyOps: a.lossyOps + 1,
		bigFn: &bigFn{
			op:   floatOperationSqrt,
			args: []*Float{a},
		},
		big: nil,
	}
}

func absBig(a *big.Float) *big.Float {
	return newFloatBig().Abs(a)
}

func Abs(a *Float) *Float {
	return &Float{
		simple:   math.Abs(a.simple),
		lossyOps: a.lossyOps + 1,
		bigFn: &bigFn{
			op:   floatOperationAbs,
			args: []*Float{a},
		},
		big: nil,
	}
}
