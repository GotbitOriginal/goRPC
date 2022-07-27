package price

import (
	"git.gotbit.io/gotbit/stroke"
	"github.com/GotbitOriginal/goRPC/internal/logic/v3/pair"
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
	Slot0, err := m.pairManager.GetSlot0()
	if err != nil {
		return 0.0, err
	}
	decimalsDiff := int64(0)

	if m.pairManager.GetPair().Flag {
		decimalsDiff = m.pairManager.GetPair().Base.Decimal - m.pairManager.GetPair().Quote.Decimal
	} else {
		decimalsDiff = m.pairManager.GetPair().Quote.Decimal - m.pairManager.GetPair().Base.Decimal
	}

	price := convertPriceToTokens(Slot0.SqrtPriceX96, decimalsDiff)

	if m.pairManager.GetPair().Flag {
		price = invertPrice(price)
	}

	price.SetPrec(uint(18))
	priceFloat64, _ := price.Float64()

	return priceFloat64, err
}
