package utils

import "math/big"

func _getAmountOut(myMaxBuy, reserveOut, reserveIn *big.Int) *big.Int ***REMOVED***

	var myMaxBuy9975 = new(big.Int)
	var z = new(big.Int)
	num := big.NewInt(9975)
	myMaxBuy9975.Mul(num, myMaxBuy)
	num.Mul(myMaxBuy9975, reserveOut)

	den := big.NewInt(10000)
	den.Mul(den, reserveIn)
	den.Add(den, myMaxBuy9975)
	z.Div(num, den)
	return z
***REMOVED***
