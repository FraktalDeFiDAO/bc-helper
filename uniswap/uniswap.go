package uniswap

import (
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type IUniswap interface {
	Swap() ([]*big.Int, error)
	Liquidity() (*[]big.Int, error)
}

func SortTokens(tokenA common.Address, tokenB common.Address) (common.Address, common.Address) {
	tA := new(big.Int)
	tB := new(big.Int)

	tA.SetString(tokenA.Hex(), 16)
	tB.SetString(tokenA.Hex(), 16)

	if tA.Cmp(tB) == 0 {
		log.Fatal("ERROR: Can not sort a token and itself")
	} else if tA.Cmp(tB) < 0 {
		return tokenA, tokenB
	} else {
		return tokenB, tokenA
	}
}
