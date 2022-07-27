package entity

import (
	"encoding/json"
)

type Configuration struct {
	RPC     string `json:"rpc"`
	Factory string `json:"factory"`
	V       int    `json:"v"`
	Fee     string `json:"fee"`
}

type Config map[string]Configuration

const configGlobal = "{\"cake1\":{\"rpc\":\"https://bsc-dataseed.binance.org/\",\"factory\":\"0xBCfCcbde45cE874adCB698cC183deBcF17952812\",\"v\":2},\"cake2\":{\"rpc\":\"https://bsc-dataseed.binance.org/\",\"factory\":\"0xcA143Ce32Fe78f1f7019d7d551a6402fC5350c73\",\"v\":2},\"uni\":{\"rpc\":\"https://mainnet.infura.io/v3/cefa7de205f543888138627880fab9cb\",\"factory\":\"0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f\",\"v\":2},\"quick\":{\"rpc\":\"https://polygon-rpc.com/\",\"factory\":\"0x5757371414417b8C6CAad45bAeF941aBc7d3Ab32\",\"v\":2},\"pangolin\":{\"rpc\":\"https://api.avax.network/ext/bc/C/rpc\",\"factory\":\"0xefa94DE7a4656D787667C749f7E1223D71E9FD88\",\"v\":2},\"joe\":{\"rpc\":\"https://api.avax.network/ext/bc/C/rpc\",\"factory\":\"0x9Ad6C38BE94206cA50bb0d90783181662f0Cfa10\",\"v\":2},\"ape\":{\"rpc\":\"https://bsc-dataseed.binance.org/\",\"factory\":\"0x0841BD0B734E4F5853f0dD8d7Ea041c241fb0Da6\",\"v\":2},\"baby\":{\"rpc\":\"https://bsc-dataseed.binance.org/\",\"factory\":\"0x86407bea2078ea5f5eb5a52b2caa963bc1f889da\",\"v\":2},\"spooky\":{\"rpc\":\"https://rpc.ftm.tools/\",\"factory\":\"0x152eE697f2E276fA89E96742e9bB9aB1F2E61bE3\",\"v\":2},\"sushibsc\":{\"rpc\":\"https://bsc-dataseed.binance.org/\",\"factory\":\"0xc35DADB65012eC5796536bD9864eD8773aBc74C4\",\"v\":2},\"sushimatic\":{\"rpc\":\"https://polygon-rpc.com/\",\"factory\":\"0xc35DADB65012eC5796536bD9864eD8773aBc74C4\",\"v\":2},\"mdexheco\":{\"rpc\":\"https://http-mainnet-node.huobichain.com/\",\"factory\":\"0xb0b670fc1F7724119963018DB0BfA86aDb22d941\",\"v\":2},\"biswap\":{\"rpc\":\"https://bsc-dataseed.binance.org/\",\"factory\":\"0x858E3312ed3A876947EA49d572A7C42DE08af7EE\",\"v\":2},\"mimo\":{\"rpc\":\"https://babel-api.mainnet.iotex.io/\",\"factory\":\"0xda257cBe968202Dea212bBB65aB49f174Da58b9D\",\"v\":2},\"uni3/1\":{\"rpc\":\"https://mainnet.infura.io/v3/cefa7de205f543888138627880fab9cb\",\"factory\":\"0x1F98431c8aD98523631AE4a59f267346ea31F984\",\"v\":3,\"fee\":\"1\"},\"uni3/0.3\":{\"rpc\":\"https://mainnet.infura.io/v3/cefa7de205f543888138627880fab9cb\",\"factory\":\"0x1F98431c8aD98523631AE4a59f267346ea31F984\",\"v\":3,\"fee\":\"0.3\"},\"uni3/0.05\":{\"rpc\":\"https://mainnet.infura.io/v3/cefa7de205f543888138627880fab9cb\",\"factory\":\"0x1F98431c8aD98523631AE4a59f267346ea31F984\",\"v\":3,\"fee\":\"0.05\"},\"uni3/0.01\":{\"rpc\":\"https://mainnet.infura.io/v3/cefa7de205f543888138627880fab9cb\",\"factory\":\"0x1F98431c8aD98523631AE4a59f267346ea31F984\",\"v\":3,\"fee\":\"0.01\"}}"

func GetConfig() (config Config, err error) {
	err = json.Unmarshal([]byte(configGlobal), &config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
