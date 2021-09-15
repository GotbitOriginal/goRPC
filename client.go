package gorpc

import (
	"fmt"

	"github.com/GotbitOriginal/goRPC/factorycontract"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetClient(swap string) (client *ethclient.Client, factory *factorycontract.Factorycontract, err error) {
	switch {
	case swap == "cake1":
		client, err = ethclient.Dial("https://bsc-dataseed.binance.org/")
		if err != nil {
			return nil, nil, fmt.Errorf("Error with client: %s", err.Error())
		}

		factory, err = factorycontract.NewFactorycontract(common.HexToAddress("0xBCfCcbde45cE874adCB698cC183deBcF17952812"), client)
		if err != nil {
			return nil, nil, fmt.Errorf("Error with factory: %s", err)
		}

		return client, factory, nil
	case swap == "cake2":
		client, err = ethclient.Dial("https://bsc-dataseed.binance.org/")
		if err != nil {
			return nil, nil, fmt.Errorf("Error with client: %s", err.Error())
		}

		factory, err = factorycontract.NewFactorycontract(common.HexToAddress("0xcA143Ce32Fe78f1f7019d7d551a6402fC5350c73"), client)
		if err != nil {
			return nil, nil, fmt.Errorf("Error with factory: %s", err)
		}

		return client, factory, nil
	case swap == "uni":
		client, err = ethclient.Dial("https://mainnet.infura.io/v3/cefa7de205f543888138627880fab9cb")
		if err != nil {
			return nil, nil, fmt.Errorf("Error with client: %s", err.Error())
		}

		factory, err = factorycontract.NewFactorycontract(common.HexToAddress("0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f"), client)
		if err != nil {
			return nil, nil, fmt.Errorf("Error with factory: %s", err)
		}

		return client, factory, nil
	case swap == "quick":
		client, err = ethclient.Dial("https://rpc-mainnet.maticvigil.com/v1/4b7b61cd4277cd3f43ac6291aafa4b629a1a66b5")
		if err != nil {
			return nil, nil, fmt.Errorf("Error with client: %s", err.Error())
		}

		factory, err = factorycontract.NewFactorycontract(common.HexToAddress("0x5757371414417b8C6CAad45bAeF941aBc7d3Ab32"), client)
		if err != nil {
			return nil, nil, fmt.Errorf("Error with factory: %s", err)
		}

		return client, factory, nil
	case swap == "pangolin":
		client, err = ethclient.Dial("https://api.avax.network/ext/bc/C/rpc")
		if err != nil {
			return nil, nil, fmt.Errorf("Error with client: %s", err.Error())
		}

		factory, err = factorycontract.NewFactorycontract(common.HexToAddress("0xefa94DE7a4656D787667C749f7E1223D71E9FD88"), client)
		if err != nil {
			return nil, nil, fmt.Errorf("Error with factory: %s", err)
		}

		return client, factory, nil
	default:
		return nil, nil, fmt.Errorf("wrong Swap")
	}
}
