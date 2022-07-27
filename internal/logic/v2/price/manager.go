package price

import (
	"git.gotbit.io/gotbit/stroke"
	"github.com/GotbitOriginal/goRPC/internal/logic/v2/pair"
	"github.com/ethereum/go-ethereum/common/math"
)

type PriceManager interface {
	GetPrice() (result float64, err error)
}

type priceManager struct {
	logger      stroke.Logger
	pairManager pair.PairManager
}

func NewPriceManager(logger stroke.Logger, pairManager pair.PairManager) PriceManager {
	return &priceManager{
		pairManager: pairManager,
		logger:      logger,
	}
}

func (m *priceManager) GetPrice() (float64, error) {
	prec := int64(18)
	reserves, err := m.pairManager.GetReserves()
	if err != nil {
		return 0.0, err
	}

	if m.pairManager.GetPair().Flag {
		result := reserves.R0

		result.Mul(result, math.BigPow(10, m.pairManager.GetPair().Base.Decimal+prec-m.pairManager.GetPair().Quote.Decimal)).Div(result, reserves.R1)
		return floatify(result.String(), int(prec))
	} else {
		result := reserves.R1

		result.Mul(result, math.BigPow(10, m.pairManager.GetPair().Quote.Decimal+prec-m.pairManager.GetPair().Base.Decimal)).Div(result, reserves.R0)
		return floatify(result.String(), int(prec))
	}
}
