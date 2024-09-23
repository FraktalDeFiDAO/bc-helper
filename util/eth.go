package util

import (
	"math"
	"math/big"
)

func ParseEthers(amt *big.Int, decimals uint8) *big.Float {
	fAmt := new(big.Float)
	fAmt.SetPrec(256).SetString(amt.String())
	return new(big.Float).Quo(fAmt, big.NewFloat(math.Pow10(int(decimals))))

}

func FormatEthers(amt *big.Float, decimals uint8) *big.Int {
	r := new(big.Int)
	d := big.NewInt(int64(decimals))
	var prec uint = 256
	shifted := d.Exp(big.NewInt(10), d, nil)
	fShifted := new(big.Float)
	fShifted.SetPrec(prec).SetString(shifted.String())

	n := amt.Mul(amt, fShifted)
	r.SetString(n.Text('f', -1), 10)

	return r
}
