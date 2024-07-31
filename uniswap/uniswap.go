package uniswap

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

func SortTokens(tokenA common.Address, tokenB common.Address) (common.Address, common.Address) {
	tA := new(big.Int)
	tB := new(big.Int)

	tA.SetString(tokenA.Hex(), 16)
	tB.SetString(tokenA.Hex(), 16)

	if tA.Cmp(tB) < 0 {
		return tokenA, tokenB
	} else {
		return tokenB, tokenA
	}
}
