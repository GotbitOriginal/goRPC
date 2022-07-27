package token

import (
	"fmt"

	"git.gotbit.io/gotbit/stroke"
	"github.com/ethereum/go-ethereum/common"
)

type TokenManager interface {
	SetUp(contract string) error
	GetToken() Token
}

type tokenManager struct {
	logger stroke.Logger
	repo   Repo
	Token
}

func NewTokenManager(logger stroke.Logger, repo Repo) TokenManager {
	return &tokenManager{
		logger: logger,
		repo:   repo,
	}
}

func (m *tokenManager) SetUp(contract string) error {
	m.Contract = common.HexToAddress(contract)

	decimalLeft, err := m.repo.GetDecimals(m.Contract)
	if err != nil {
		return fmt.Errorf("error with get decimal left: %s", err.Error())
	}

	m.Decimal = decimalLeft

	return nil
}

func (m *tokenManager) GetToken() Token {
	return m.Token
}
