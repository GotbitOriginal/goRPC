package gorpc

import (
	"fmt"

	"github.com/GotbitOriginal/goRPC/factorycontract"
	"github.com/GotbitOriginal/goRPC/tokencontract"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func getDecimals(client *ethclient.Client, address common.Address) (int64, error) {
	cTokenA, err := tokencontract.NewTokencontract(address, client)
	if err != nil {
		return 0, err
	}
	decimals, err := cTokenA.Decimals(nil)

	return decimals.Int64(), err
}

func getPairAddress(tokenA common.Address, tokenB common.Address, factory *factorycontract.Factorycontract) (common.Address, error) {
	pairAddress, err := factory.GetPair(nil, tokenA, tokenB)
	if err != nil {
		return tokenA, err
	}

	if pairAddress.Hex() == common.HexToAddress("0x0000000000000000000000000000000000000000").Hex() {
		return tokenA, fmt.Errorf("Pair does not exist")
	}

	return pairAddress, nil
}
