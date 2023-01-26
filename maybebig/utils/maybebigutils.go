/*
This file is useful to bring into a project that uses the maybebig.Float extensively
You just copy it into your package and all of the functions are now available in a
readable yet minimal infix syntax.
*/
package shims // your package name here

import (
	"math/big"

	"github.com/gbdubs/maybebig"
)

// Input

func newFloat() *maybebig.Float             { return maybebig.NewFloatZ() }
func fromInt(i int) *maybebig.Float         { return maybebig.FromInt(i) }
func fromFloat64(f float64) *maybebig.Float { return maybebig.FromFloat64(f) }
func fromBig(b *big.Float) *maybebig.Float  { return maybebig.FromBig(b) }

// Math
func add(a ...*maybebig.Float) *maybebig.Float { return maybebig.Add(a...) }
func sub(a, b *maybebig.Float) *maybebig.Float { return maybebig.Sub(a, b) }
func mul(a ...*maybebig.Float) *maybebig.Float { return maybebig.Mul(a...) }
func div(a, b *maybebig.Float) *maybebig.Float { return maybebig.Div(a, b) }
func abs(a *maybebig.Float) *maybebig.Float    { return maybebig.Abs(a) }
func sq(a *maybebig.Float) *maybebig.Float     { return maybebig.Sq(a) }
func sqrt(a *maybebig.Float) *maybebig.Float   { return maybebig.Sqrt(a) }

// Trigonometry

func pi() *maybebig.Float                        { return maybebig.Pi() }
func sin(a *maybebig.Float) *maybebig.Float      { return maybebig.Sin(a) }
func cos(a *maybebig.Float) *maybebig.Float      { return maybebig.Cos(a) }
func atan(a *maybebig.Float) *maybebig.Float     { return maybebig.Atan(a) }
func atan2(a, b *maybebig.Float) *maybebig.Float { return maybebig.Atan2(a, b) }

// Comparisons

func eq(a, b *maybebig.Float) bool          { return maybebig.Eq(a, b) }
func eqz(a *maybebig.Float) bool            { return maybebig.Eqz(a) }
func lt(a, b *maybebig.Float) bool          { return maybebig.Lt(a, b) }
func ltz(a *maybebig.Float) bool            { return maybebig.Ltz(a) }
func lteq(a, b *maybebig.Float) bool        { return maybebig.Lteq(a, b) }
func lteqz(a *maybebig.Float) bool          { return maybebig.Lteqz(a) }
func gt(a, b *maybebig.Float) bool          { return maybebig.Gt(a, b) }
func gtz(a *maybebig.Float) bool            { return maybebig.Gtz(a) }
func gteq(a, b *maybebig.Float) bool        { return maybebig.Gteq(a, b) }
func gteqz(a *maybebig.Float) bool          { return maybebig.Gteqz(a) }
func isInf(a *maybebig.Float) bool          { return a.IsInf() }
func eqEquateInf(a, b *maybebig.Float) bool { return maybebig.EqualsEquateInf(a, b) }

// Sequences

func min(a ...*maybebig.Float) *maybebig.Float { return maybebig.Min(a...) }
func max(a ...*maybebig.Float) *maybebig.Float { return maybebig.Max(a...) }
