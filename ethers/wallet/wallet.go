package wallet

import (
	"crypto/ecdsa"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

type Wallet struct {
	PrivateKey ecdsa.PrivateKey
	PublicKey  ecdsa.PublicKey
	Address    common.Address
}
type WalletInfo struct {
	PrivateKey string
	PublicKey  string
	Address    string
}

func GenerateWallet() (wallet Wallet) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	wallet = Wallet{
		PrivateKey: *privateKey,
		PublicKey:  *publicKeyECDSA,
		Address:    common.HexToAddress(address),
	}
	return wallet
}

func (w *Wallet) ShowInfo(wallet Wallet) (walletInfo WalletInfo) {
	privateKeyBytes := crypto.FromECDSA(&wallet.PrivateKey)
	publicKeyBytes := crypto.FromECDSAPub(&wallet.PublicKey)

	walletInfo = WalletInfo{
		PrivateKey: hexutil.Encode(privateKeyBytes),
		PublicKey:  hexutil.Encode(publicKeyBytes),
		Address:    wallet.Address.Hex(),
	}
	return walletInfo
}

func FromPrivateKeyString(privateKeyStr string) (wallet Wallet) {
	privateKey, err := crypto.HexToECDSA(privateKeyStr)

	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	wallet = Wallet{
		PrivateKey: *privateKey,
		PublicKey:  *publicKeyECDSA,
		Address:    common.HexToAddress(address),
	}
	return wallet

}
