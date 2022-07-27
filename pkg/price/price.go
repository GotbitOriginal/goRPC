package price

import (
	"os"

	"github.com/GotbitOriginal/goRPC/internal/logger"
	"github.com/GotbitOriginal/goRPC/internal/logic"
)

type IGorpc interface {
	GetPrice() (float64, error)
}

type Gorpc struct {
	priceManager logic.PriceManager
}

func NewGorpc(swap, left, right string) (IGorpc, error) {
	logger := logger.New(os.Stdout)
	priceManager := logic.NewPriceManager(logger)
	priceManager.SetUp(left, right, swap)

	return &Gorpc{
		priceManager: priceManager,
	}, nil
}

func (G *Gorpc) GetPrice() (float64, error) {
	return G.priceManager.GetPrice()
}
