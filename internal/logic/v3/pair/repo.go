package pair

import (
	"fmt"
	"math/big"

	"github.com/GotbitOriginal/goRPC/internal/logic/v3/pair/factorycontract"
	"github.com/GotbitOriginal/goRPC/internal/logic/v3/pair/poolcontract"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Repo interface {
	GetPairAddress(pair Pair) (common.Address, error)
	GetPairContract(address common.Address) (*poolcontract.Pool, error)
	GetToken0Quote(pair Pair) (flag bool, err error)
	GetSlot0(pair Pair) (slot Slot0, errr error)
}

type repo struct {
	client  *ethclient.Client
	factory *factorycontract.Factory
}

func NewRepo(client *ethclient.Client, factory string) Repo {
	factoryContract, _ := factorycontract.NewFactory(common.HexToAddress(factory), client)
	return &repo{
		client:  client,
		factory: factoryContract,
	}
}

func (r *repo) GetPairAddress(pair Pair) (common.Address, error) {
	pairAddress, err := r.factory.GetPool(nil, pair.Quote.Contract, pair.Base.Contract, big.NewInt(getFee(pair.Fee)))
	if err != nil {
		return common.Address{}, err
	}

	if pairAddress.Hex() == common.HexToAddress("0x0000000000000000000000000000000000000000").Hex() {
		return common.Address{}, fmt.Errorf("Pair does not exist")
	}

	return pairAddress, nil
}

func (r *repo) GetPairContract(address common.Address) (*poolcontract.Pool, error) {
	return poolcontract.NewPool(address, r.client)
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

func (r *repo) GetSlot0(pair Pair) (slot Slot0, errr error) {
	slotPool, err := pair.Contract.Slot0(nil)
	if err != nil {
		return slot, err
	}

	slot.SqrtPriceX96 = slotPool.SqrtPriceX96

	return slot, nil
}
