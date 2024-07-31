package uniswapV2

import (
	"fmt"
	"math/big"
)

func GetAmountsOut(reserveIn, reserveOut, amountIn, fee *big.Int) []*big.Int {
	var amountOut *big.Int = new(big.Int)

	amountInWithFee := new(big.Int)
	feeConv := new(big.Float)
	_fee, _ := fee.Float64()
	ffee := big.NewFloat(_fee)

	feeConv.Mul(
		feeConv.Sub(
			big.NewFloat(1),
			feeConv.Quo(ffee, big.NewFloat(1000000)),
		),
		big.NewFloat(1000),
	)
	_feeConv, _ := feeConv.Int64()
	amountInWithFee.Mul(amountIn, big.NewInt(_feeConv))

	numerator := new(big.Int)
	denominator := new(big.Int)

	numerator.Mul(amountInWithFee, reserveOut)

	// reserveIn.mul(1000).add(amountInWithFee)
	denominator.Mul(reserveIn, big.NewInt(1000))

	denominator.Add(denominator, amountInWithFee)

	amountOut.Quo(numerator, denominator)
	return []*big.Int{
		amountIn,
		amountOut,
	}
}

func GetAmountsIn(reserveIn, reserveOut, amountOut, fee *big.Int) []*big.Int {
	// 	  uint numerator = reserveIn.mul(amountOut).mul(1000);
	//    uint denominator = reserveOut.sub(amountOut).mul(997);
	//    amountIn = (numerator / denominator).add(1);

	var (
		amountIn    *big.Int   = new(big.Int)
		numerator   *big.Int   = new(big.Int)
		denominator *big.Int   = new(big.Int)
		feeConv     *big.Float = new(big.Float)
	)
	_fee, _ := fee.Float64()
	ffee := big.NewFloat(_fee)

	feeConv.Mul(
		feeConv.Sub(
			big.NewFloat(1),
			feeConv.Quo(ffee, big.NewFloat(1000000)),
		),
		big.NewFloat(1000),
	)
	_feeConv, _ := feeConv.Int64()
	fmt.Println("_feeConv", _feeConv)

	numerator.Mul(reserveIn, amountOut)
	numerator.Mul(numerator, big.NewInt(1000))
	denominator.Mul(
		denominator.Sub(reserveOut, amountOut),
		big.NewInt(_feeConv),
	)

	amountIn.Add(
		denominator.Quo(numerator, denominator),
		big.NewInt(1),
	)
	return []*big.Int{amountIn, amountOut}
}
func FloatToBigInt(val float64) *big.Int {
	bigval := new(big.Float)
	bigval.SetFloat64(val)
	// Set precision if required.
	// bigval.SetPrec(64)

	coin := new(big.Float)
	coin.SetInt(big.NewInt(1000000000000000000))

	bigval.Mul(bigval, coin)

	result := new(big.Int)
	bigval.Int(result) // store converted number in result

	return result
}

// TODO: values may or may not need to be swap... Testing needed
func GetSpotPrice(reserveIn, reserveOut *big.Int) *big.Float {
	var (
		price *big.Float
	)
	rIn := new(big.Float)
	rOut := new(big.Float)

	rIn.SetInt(reserveIn)
	rOut.SetInt(reserveOut)

	price.Quo(rIn, rOut)
	return price
}
