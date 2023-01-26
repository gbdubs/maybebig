package maybebig

import (
	"fmt"
	"math/big"
)

type floatOperation int

const (
	floatOperationUnknown floatOperation = iota
	floatOperationAdd
	floatOperationSub
	floatOperationMul
	floatOperationDiv
	floatOperationAbs
	floatOperationSqrt
	floatOperationSin
	floatOperationCos
	floatOperationAtan
	floatOperationAtan2
)

func (b *bigFn) Eval() *big.Float {
	args := make([]*big.Float, len(b.args))
	for i, arg := range b.args {
		args[i] = arg.GetBig()
	}
	switch b.op {
	case floatOperationAdd:
		return addBig(args...)
	case floatOperationMul:
		return mulBig(args...)
	case floatOperationSub:
		return subBig(args[0], args[1])
	case floatOperationDiv:
		return divBig(args[0], args[1])
	case floatOperationAbs:
		return absBig(args[0])
	case floatOperationSqrt:
		return sqrtBig(args[0])
	case floatOperationSin:
		return sinBig(args[0])
	case floatOperationCos:
		return cosBig(args[0])
	case floatOperationAtan:
		return atanBig(args[0])
	case floatOperationAtan2:
		return atan2Big(args[0], args[1])
	}
	panic(fmt.Sprintf("unrecognized operation %q", b.op))
}
