package util

import "github.com/ethereum/go-ethereum/common"

func StringToAddress(sAddress string) (address common.Address) {
	address = common.HexToAddress(sAddress)
	return address
}
