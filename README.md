# Mostly Small `big.Float`

Ever have a situation in Golang where you want the precision of `big.Float`, but most
operations don't need it? This library may be for you!

## Overview

This library's main object is `maybebig.Float`, which wraps a `float64` which is always
evaluated, and a `big.Float` which may or may not be evaluated. The `maybefloat.Float` 
retains pointers to it's dependencies in such a way that if you ever want to compute the
precice value, you can do so by calling `myMaybeFloat.GetBig()`, which will compute and
return the full precision result.

The library automatically does the `big.Float` computations in a couple of circumstances:

- when the accumulated number of floating point operations reaches a certain size
(configurable)
- when a comparison yields a difference that is within some tolerance (configurable)
- when a user asks for it via `maybebig.Float.ForceComputation()`

## Support 

This library is a WIP, used in a few of my projects, but only supports the functions I've
built thus far. If you need some functions that aren't supported, just ask, I'm happy to
build them.

## Bugs

Please feel free to file a bug if you find one, I'm motivated to keep this library well supported.
