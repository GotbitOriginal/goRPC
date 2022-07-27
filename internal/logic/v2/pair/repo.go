package pair

import (
	"fmt"

	"github.com/GotbitOriginal/goRPC/internal/logic/v2/pair/factorycontract"
	"github.com/GotbitOriginal/goRPC/internal/logic/v2/pair/paircontract"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Repo interface {
	GetPairAddress(pair Pair) (common.Address, error)
	GetPairContract(address common.Address) (*paircontract.Paircontract, error)
	GetToken0Quote(pair Pair) (flag bool, err error)
	GetReserves(pair Pair) (res Reserve, errr error)
}

type repo struct {
	client  *ethclient.Client
	factory *factorycontract.Factorycontract
}

func NewRepo(client *ethclient.Client, factory string) Repo {
	factoryContract, _ := factorycontract.NewFactorycontract(common.HexToAddress(factory), client)
	return &repo{
		client:  client,
		factory: factoryContract,
	}
}

func (r *repo) GetPairAddress(pair Pair) (common.Address, error) {
	pairAddress, err := r.factory.GetPair(nil, pair.Quote.Contract, pair.Base.Contract)
	if err != nil {
		return common.Address{}, err
	}

	if pairAddress.Hex() == common.HexToAddress("0x0000000000000000000000000000000000000000").Hex() {
		return common.Address{}, fmt.Errorf("Pair does not exist")
	}

	return pairAddress, nil
}

func (r *repo) GetPairContract(address common.Address) (*paircontract.Paircontract, error) {
	return paircontract.NewPaircontract(address, r.client)
}

func (r *repo) GetToken0Quote(pair Pair) (flag bool, err error) {
	token0, err := pair.Contract.Token0(nil)
	if err != nil {
		return flag, err
	}

	if token0.Hex() == pair.Quote.Contract.Hex() {
		flag = true
	}

	return flag, nil
}

func (r *repo) GetReserves(pair Pair) (res Reserve, errr error) {
	reserves, err := pair.Contract.GetReserves(nil)
	if err != nil {
		return res, err
	}

	res.R0 = reserves.Reserve0
	res.R1 = reserves.Reserve1

	return res, nil
}
