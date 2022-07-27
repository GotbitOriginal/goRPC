package pair

import (
	"git.gotbit.io/gotbit/stroke"
	"github.com/GotbitOriginal/goRPC/internal/logic/v3/token"
)

type PairManager interface {
	SetUp(quote, base token.Token, fee string) (err error)
	GetSlot0() (Slot0, error)
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

func (m *pairManager) SetUp(quote, base token.Token, fee string) (err error) {
	m.Pair.Quote = quote
	m.Pair.Base = base

	err = validateFee(fee)
	if err != nil {
		return err
	}
	m.Pair.Fee = fee

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

func (m *pairManager) GetSlot0() (Slot0, error) {
	return m.repo.GetSlot0(m.Pair)
}

func (m *pairManager) GetPair() Pair {
	return m.Pair
}
