package wallet

import (
	"crypto/ecdsa"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

type IWallet interface {
	AddressToString() string
	PublicKeyToString() string
	PrivateKeyToString() string
	ShowInfo(wallet Wallet) (walletInfo WalletInfo)
}
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
	walletInfo = WalletInfo{
		PrivateKey: w.PrivateKeyToString(),
		PublicKey:  w.PrivateKeyToString(),
		Address:    w.AddressToString(),
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

func (w *Wallet) PrivateKeyToString() string {
	privateKeyBytes := crypto.FromECDSA(&w.PrivateKey)
	return hexutil.Encode(privateKeyBytes)[4:]
}
func (w *Wallet) PublicKeyToString() string {
	publicKey := w.PrivateKey.Public()

	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	return hexutil.Encode(publicKeyBytes)[4:]
}
func (w *Wallet) AddressToString() string {
	return w.Address.Hex()
}
