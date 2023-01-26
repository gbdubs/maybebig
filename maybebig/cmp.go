package maybebig

import (
	"math"
	"math/big"
)

var zero = NewFloatZ().SetInt64(0)

func Eq(a, b *Float) bool {
	hardRecomputeIfWithinCheckThreshold(a, b)
	simpleAns := a.simple == b.simple
	if !doubleCheckMode {
		return simpleAns
	}
	bigAns := eqBig(a.GetBig(), b.GetBig())
	totalLossyOps := a.lossyOps + b.lossyOps
	if simpleAns != bigAns {
		recordDoubleCheckError(totalLossyOps)
	}
	recordDoubleCheckCount()
	return bigAns
}

func eqBig(a, b *big.Float) bool {
	return a.Cmp(b) == 0
}

func Eqz(a *Float) bool {
	return Eq(a, zero)
}

func IsZero(a *Float) bool {
	return Eqz(a)
}

func ltBig(a, b *big.Float) bool {
	return a.Cmp(b) < 0
}

func Lt(a, b *Float) bool {
	hardRecomputeIfWithinCheckThreshold(a, b)
	simpleAns := a.simple < b.simple
	if !doubleCheckMode {
		return simpleAns
	}
	bigAns := ltBig(a.GetBig(), b.GetBig())
	totalLossyOps := a.lossyOps + b.lossyOps
	if simpleAns != bigAns {
		recordDoubleCheckError(totalLossyOps)
	}
	recordDoubleCheckCount()
	return bigAns
}

func Ltz(a *Float) bool {
	return Lt(a, zero)
}

func gtBig(a, b *big.Float) bool {
	return a.Cmp(b) > 0
}

func Gt(a, b *Float) bool {
	hardRecomputeIfWithinCheckThreshold(a, b)
	simpleAns := a.simple > b.simple
	if !doubleCheckMode {
		return simpleAns
	}
	bigAns := gtBig(a.GetBig(), b.GetBig())
	totalLossyOps := a.lossyOps + b.lossyOps
	if simpleAns != bigAns {
		recordDoubleCheckError(totalLossyOps)
	}
	recordDoubleCheckCount()
	return bigAns
}

func Gtz(a *Float) bool {
	return Gt(a, zero)
}

func lteqBig(a, b *big.Float) bool {
	return a.Cmp(b) <= 0
}

func Lteq(a, b *Float) bool {
	hardRecomputeIfWithinCheckThreshold(a, b)
	simpleAns := a.simple <= b.simple
	if !doubleCheckMode {
		return simpleAns
	}
	bigAns := lteqBig(a.GetBig(), b.GetBig())
	totalLossyOps := a.lossyOps + b.lossyOps
	if simpleAns != bigAns {
		recordDoubleCheckError(totalLossyOps)
	}
	recordDoubleCheckCount()
	return bigAns
}

func Lteqz(a *Float) bool {
	return Lteq(a, zero)
}

func GteqBig(a, b *big.Float) bool {
	return a.Cmp(b) >= 0
}

func Gteq(a, b *Float) bool {
	hardRecomputeIfWithinCheckThreshold(a, b)
	simpleAns := a.simple >= b.simple
	if !doubleCheckMode {
		return simpleAns
	}
	bigAns := GteqBig(a.GetBig(), b.GetBig())
	totalLossyOps := a.lossyOps + b.lossyOps
	if simpleAns != bigAns {
		recordDoubleCheckError(totalLossyOps)
	}
	recordDoubleCheckCount()
	return bigAns
}

func Gteqz(a *Float) bool {
	return Gteq(a, zero)
}

func (f *Float) IsInf() bool {
	simpleAns := math.IsInf(f.simple, 0)
	if !doubleCheckMode {
		return simpleAns
	}
	b := f.GetBig()
	bigAns := b.IsInf()
	if simpleAns != bigAns {
		recordDoubleCheckError(f.lossyOps)
	}
	recordDoubleCheckCount()
	return bigAns
}

func EqualsEquateInf(a, b *Float) bool {
	hardRecomputeIfWithinCheckThreshold(a, b)
	simpleAns := false
	if math.IsInf(a.simple, 0) && math.IsInf(b.simple, 0) {
		simpleAns = true
	} else {
		simpleAns = a.simple == b.simple
	}
	if !doubleCheckMode {
		return simpleAns
	}
	bigAns := equalsEquateInfBig(a.GetBig(), b.GetBig())
	totalLossyOps := a.lossyOps + b.lossyOps
	if simpleAns != bigAns {
		recordDoubleCheckError(totalLossyOps)
	}
	recordDoubleCheckCount()
	return bigAns
}

func equalsEquateInfBig(a, b *big.Float) bool {
	if a.IsInf() && b.IsInf() {
		return true
	}
	return eqBig(a, b)
}

func hardRecomputeIfWithinCheckThreshold(a, b *Float) {
	diff := a.simple - b.simple
	if negCheckThreshold < diff && diff < posCheckThreshold {
		a.hardRecompute()
		b.hardRecompute()
	}
}
