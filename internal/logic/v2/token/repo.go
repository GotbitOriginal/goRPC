package token

import (
	"github.com/GotbitOriginal/goRPC/internal/logic/v2/token/tokencontract"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Repo interface {
	GetDecimals(address common.Address) (int64, error)
}

type repo struct {
	client *ethclient.Client
}

func NewRepo(client *ethclient.Client) Repo {
	return &repo{
		client: client,
	}
}

func (r *repo) GetDecimals(address common.Address) (int64, error) {
	cTokenA, err := tokencontract.NewTokencontract(address, r.client)
	if err != nil {
		return 0, err
	}
	decimals, err := cTokenA.Decimals(nil)

	return decimals.Int64(), err
}
