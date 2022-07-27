package pair

import (
	"math/big"

	"github.com/GotbitOriginal/goRPC/internal/logic/v2/pair/paircontract"
	"github.com/GotbitOriginal/goRPC/internal/logic/v2/token"
	"github.com/ethereum/go-ethereum/common"
)

type Pair struct {
	Quote    token.Token
	Base     token.Token
	Address  common.Address
	Contract *paircontract.Paircontract
	Flag     bool
}

type Reserve struct {
	R0 *big.Int
	R1 *big.Int
}
