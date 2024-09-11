package ethers

import (
	"bytes"
	"context"
	"log"
	"math/big"
	"time"

	wallet "github.com/FraktalDeFiDAO/bc-helper/ethers/wallet"
	"github.com/ethereum/go-ethereum/event"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
)

type Account common.Address

type Network struct {
	Name    string
	ChainId *big.Int
}
type EthersHelper struct {
	Networks []Network
}
type Provider struct {
	ChainId *big.Int
	Url     string
	Client  ethclient.Client
}

func GetBalance(
	client *ethclient.Client,
	account common.Address,
	blockNumber *big.Int,
) (balance *big.Int) {
	balance, err := client.BalanceAt(context.Background(), account, blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	return balance
}

func GetPendingBalance(
	client *ethclient.Client,
	account common.Address,
	blockNumber *big.Int,
) (balance *big.Int) {
	balance, err := client.BalanceAt(context.Background(), account, blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	return balance
}

func CreateClient(providerUrl string) (client *ethclient.Client) {
	var err error
	client, err = ethclient.Dial(providerUrl)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func GetCuuerntBlock(client *ethclient.Client) (blockNumber *big.Int) {
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	blockNumber = header.Number
	return blockNumber
}

func GetNonce(client *ethclient.Client, account common.Address) (nonce uint64) {
	nonce, err := client.PendingNonceAt(context.Background(), account)
	if err != nil {
		log.Fatal(err)
	}
	return nonce
}

func GetChainId(client *ethclient.Client) (chainId *big.Int) {
	var err error
	chainId, err = client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	return chainId
}

func GetGasPrice(client *ethclient.Client) (gasPrice *big.Int) {
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	return gasPrice
}

func BuildSendETH(
	client *ethclient.Client,
	wallet wallet.Wallet,
	to common.Address,
	amount *big.Int,
) (signedTx *types.Transaction) {
	gasLimit := uint64(21000)
	gasPrice := GetGasPrice(client)

	nonce := GetNonce(client, wallet.Address)

	var data []byte
	var err error

	tx := types.NewTransaction(nonce, to, amount, gasLimit, gasPrice, data)
	chainId := GetChainId(client)

	signedTx, err = types.SignTx(tx, types.NewEIP155Signer(chainId), &wallet.PrivateKey)

	if err != nil {
		log.Fatal(err)
	}

	return signedTx
}
func SendETH(client *ethclient.Client, signedTx *types.Transaction) (tx *types.Transaction) {

	if err := client.SendTransaction(context.Background(), signedTx); err != nil {
		log.Fatal(err)
	}

	// fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
	return signedTx
}

func CreateRawTx(
	client *ethclient.Client,
	wallet wallet.Wallet,
	to common.Address,
	amount *big.Int,
	data []byte,
) (
	signedTx *types.Transaction,
	rawTxBytes []byte,
) {
	gasLimit := uint64(21000)
	gasPrice := GetGasPrice(client)

	nonce := GetNonce(client, wallet.Address)

	tx := types.NewTransaction(nonce, to, amount, gasLimit, gasPrice, data)
	chainId := GetChainId(client)

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainId), &wallet.PrivateKey)

	if err != nil {
		log.Fatal(err)
	}
	ts := types.Transactions{signedTx}
	b := new(bytes.Buffer)
	ts.EncodeIndex(0, b)
	rawTxBytes = b.Bytes()

	return signedTx, rawTxBytes
}
func SendRawTx(client *ethclient.Client, rawTxBytes []byte) (tx *types.Transaction) {

	tx = new(types.Transaction)
	rlp.DecodeBytes(rawTxBytes, &tx)

	err := client.SendTransaction(context.Background(), tx)
	if err != nil {
		log.Fatal(err)
	}

	return tx
}
func resubFunc(
	client *ethclient.Client,
	query ethereum.FilterQuery,
	logs chan types.Log,
) event.ResubscribeFunc {

	return func(ctx context.Context) (event.Subscription, error) {
		return client.SubscribeFilterLogs(context.Background(), query, logs)
	}
}
func Subscribe(
	client *ethclient.Client,
	query ethereum.FilterQuery,
	logs chan types.Log,
) event.Subscription {
	log.Println("Starting Subscription...")

	subTimeout := time.Second * 2

	sub := event.Resubscribe(subTimeout, resubFunc(client, query, logs))

	return sub
}
