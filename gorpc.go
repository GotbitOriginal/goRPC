package gorpc

import (
	"fmt"
	"gorpc/factorycontract"
	"gorpc/paircontract"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/ethclient"
)

type IGorpc interface {
	GetPrice() (float64, error)
}

type Gorpc struct {
	client       *ethclient.Client
	flag         bool
	decimalLeft  int64
	decimalRight int64
	pair         *paircontract.Paircontract
	factory      *factorycontract.Factorycontract
}

func NewGorpc(swap, left, right string) (Gorpc, error) {
	client, factory, err := GetClient(swap)
	if err != nil {
		time.Sleep(time.Second)
		client, factory, err = GetClient(swap)
		if err != nil {
			return Gorpc{}, err
		}
	}

	tokenA := common.HexToAddress(left)
	tokenB := common.HexToAddress(right)

	decimalLeft, err := getDecimals(client, tokenA)
	if err != nil {
		return Gorpc{}, fmt.Errorf("Erroro with get decimal left: %s", err.Error())
	}

	decimalRight, err := getDecimals(client, tokenB)
	if err != nil {
		return Gorpc{}, fmt.Errorf("Erroro with get decimal right: %s", err.Error())
	}

	pairAddress, err := getPairAddress(tokenA, tokenB, factory)
	if err != nil {
		return Gorpc{}, fmt.Errorf("Erroro with get pair: %s", err.Error())
	}

	cPair, err := paircontract.NewPaircontract(pairAddress, client)
	if err != nil {
		return Gorpc{}, fmt.Errorf("Error in bind pair contract: %s", err)
	}

	token0, err := cPair.Token0(nil)
	if err != nil {
		return Gorpc{}, fmt.Errorf("Error with get token0: %s", err)
	}

	flag := false
	if token0.Hex() == tokenA.Hex() {
		flag = true
	}

	return Gorpc{
		client:       client,
		flag:         flag,
		factory:      factory,
		decimalLeft:  decimalLeft,
		decimalRight: decimalRight,
		pair:         cPair,
	}, nil
}

func (G *Gorpc) GetPrice(prec int64) (float64, error) {
	reserves, err := G.pair.GetReserves(nil)
	if err != nil {
		return 0.0, fmt.Errorf("Error in get reserves: %s", err)
	}

	if G.flag {
		result, ok := new(big.Int).SetString(reserves.Reserve1.String(), 10)
		if !ok {
			return 0.0, fmt.Errorf("SetString(): conversion problem")
		}

		result.Mul(result, math.BigPow(10, G.decimalLeft+prec)).Div(result, math.BigPow(10, G.decimalRight)).Div(result, reserves.Reserve0)
		return floatify(result.String(), int(prec))
	} else {
		result, ok := new(big.Int).SetString(reserves.Reserve0.String(), 10)
		if !ok {
			return 0.0, fmt.Errorf("SetString(): conversion problem")
		}

		result.Mul(result, math.BigPow(10, G.decimalLeft+prec)).Div(result, math.BigPow(10, G.decimalRight)).Div(result, reserves.Reserve1)
		return floatify(result.String(), int(prec))
	}
}
