package pair

import (
	"fmt"
	"math/big"

	"github.com/GotbitOriginal/goRPC/internal/logic/v3/pair/poolcontract"
	"github.com/GotbitOriginal/goRPC/internal/logic/v3/token"
	"github.com/ethereum/go-ethereum/common"
)

type Pair struct {
	Quote    token.Token
	Base     token.Token
	Address  common.Address
	Contract *poolcontract.Pool
	Flag     bool
	Fee      string
}

type Slot0 struct {
	SqrtPriceX96 *big.Int
}

const (
	Fee001 = int64(100)
	Fee005 = int64(500)
	Fee03  = int64(3000)
	Fee1   = int64(10000)
)

func getFee(f string) int64 {
	if f == "1" {
		return Fee1
	}

	if f == "0.3" {
		return Fee03
	}

	if f == "0.05" {
		return Fee005
	}

	if f == "0.01" {
		return Fee001
	}

	return -1
}

func validateFee(f string) error {
	if f == "1" {
		return nil
	}

	if f == "0.3" {
		return nil
	}

	if f == "0.05" {
		return nil
	}

	if f == "0.01" {
		return nil
	}

	return fmt.Errorf("fee is not supported")
}
