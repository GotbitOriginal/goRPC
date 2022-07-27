package token

import "github.com/ethereum/go-ethereum/common"

type Token struct {
	Decimal  int64
	Contract common.Address
}
