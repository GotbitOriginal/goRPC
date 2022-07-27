package pair

import (
	"git.gotbit.io/gotbit/stroke"
	"github.com/GotbitOriginal/goRPC/internal/logic/v2/token"
)

type PairManager interface {
	SetUp(quote, base token.Token) (err error)
	GetReserves() (Reserve, error)
	GetPair() Pair
}

type pairManager struct {
	logger stroke.Logger
	repo   Repo
	Pair
}

func NewPairManager(logger stroke.Logger, repo Repo) PairManager {
	return &pairManager{
		logger: logger,
		repo:   repo,
	}
}

func (m *pairManager) SetUp(quote, base token.Token) (err error) {
	m.Pair.Quote = quote
	m.Pair.Base = base

	m.Pair.Address, err = m.repo.GetPairAddress(m.Pair)
	if err != nil {
		return err
	}

	m.Pair.Contract, err = m.repo.GetPairContract(m.Pair.Address)
	if err != nil {
		return err
	}

	m.Pair.Flag, err = m.repo.GetToken0Quote(m.Pair)
	if err != nil {
		return err
	}

	return nil
}

func (m *pairManager) GetReserves() (Reserve, error) {
	return m.repo.GetReserves(m.Pair)
}

func (m *pairManager) GetPair() Pair {
	return m.Pair
}
