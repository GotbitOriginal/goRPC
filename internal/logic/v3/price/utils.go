package price

import (
	"fmt"
	"math/big"

	"github.com/ALTree/bigfloat"
	"github.com/ethereum/go-ethereum/common/math"
)

func floatify(intString string, precision int) (float64, error) {
	if len(intString) <= precision {
		toAdd := precision - len(intString) + 1
		for i := 0; i < toAdd; i++ {
			intString = "0" + intString
		}
	}
	index := len(intString) - precision
	q := intString[:index] + "." + intString[index:]
	f, ok := new(big.Float).SetString(q)
	if !ok {
		return 0, fmt.Errorf("can't parse this resulting string to float: %s" + q)
	}
	res, _ := f.Float64()
	return res, nil
}

func convertPriceToTokens(sqrtPriceX96 *big.Int, decimalsDiff int64) *big.Float {
	//	sqrtPriceX96 doc: https://docs.uniswap.org/sdk/guides/fetching-prices
	price := math.Exp(sqrtPriceX96, big.NewInt(2))      //	sqrtPriceX96^2
	divider := math.Exp(big.NewInt(2), big.NewInt(192)) // 2^192

	priceFloat := new(big.Float).SetInt(price)
	dividerFloat := new(big.Float).SetInt(divider)

	priceFloat = priceFloat.Quo(priceFloat, dividerFloat) //	priceX96 / 2^192

	decimalBalancer := bigfloat.Pow(
		big.NewFloat(10),
		big.NewFloat(float64(decimalsDiff))) //	const 10^decimal

	priceFloat = priceFloat.Mul(priceFloat, decimalBalancer) //	decimals difference eliminated here
	return priceFloat
}

func invertPrice(price *big.Float) *big.Float {
	floatOne := new(big.Float).SetInt(big.NewInt(1))
	return price.Quo(floatOne, price) //	1 / price
}
