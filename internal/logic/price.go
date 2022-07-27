package logic

import (
	"fmt"
	"os"

	"git.gotbit.io/gotbit/stroke"
	"github.com/GotbitOriginal/goRPC/internal/entity"
	"github.com/GotbitOriginal/goRPC/internal/logger"

	pairv2 "github.com/GotbitOriginal/goRPC/internal/logic/v2/pair"
	pricev2 "github.com/GotbitOriginal/goRPC/internal/logic/v2/price"
	tokenv2 "github.com/GotbitOriginal/goRPC/internal/logic/v2/token"

	pairv3 "github.com/GotbitOriginal/goRPC/internal/logic/v3/pair"
	pricev3 "github.com/GotbitOriginal/goRPC/internal/logic/v3/price"
	tokenv3 "github.com/GotbitOriginal/goRPC/internal/logic/v3/token"

	"github.com/ethereum/go-ethereum/ethclient"
)

type PriceManagerV interface {
	pricev2.PriceManager
	pricev3.PriceManager
}

type PriceManager interface {
	SetUp(quote, base, swap string) error
	GetPrice() (float64, error)
}

type priceManager struct {
	logger        stroke.Logger
	priceManagerV PriceManagerV
}

func NewPriceManager(logger stroke.Logger) PriceManager {
	return &priceManager{
		logger: logger,
	}
}

func (m *priceManager) SetUp(quote, base, swap string) error {
	var err error
	logger := logger.New(os.Stdout)

	config, err := entity.GetConfig()
	if err != nil {
		return err
	}

	confguration, ok := config[swap]
	if !ok {
		return fmt.Errorf("swap %s is not supported", swap)
	}

	client, err := ethclient.Dial(confguration.RPC)
	if err != nil {
		return err
	}

	switch confguration.V {
	case 2:
		tokenRepo := tokenv2.NewRepo(client)

		tokenManagerQuote := tokenv2.NewTokenManager(logger, tokenRepo)
		err = tokenManagerQuote.SetUp(quote)
		if err != nil {
			return err
		}

		tokenManagerBase := tokenv2.NewTokenManager(logger, tokenRepo)
		err = tokenManagerBase.SetUp(base)
		if err != nil {
			return err
		}

		pairRep := pairv2.NewRepo(client, confguration.Factory)
		pairManager := pairv2.NewPairManager(logger, pairRep)
		err = pairManager.SetUp(tokenManagerQuote.GetToken(), tokenManagerBase.GetToken())
		if err != nil {
			return err
		}

		m.priceManagerV = pricev2.NewPriceManager(logger, pairManager)
	case 3:
		tokenRepo := tokenv3.NewRepo(client)

		tokenManagerQuote := tokenv3.NewTokenManager(logger, tokenRepo)
		err = tokenManagerQuote.SetUp(quote)
		if err != nil {
			return err
		}

		tokenManagerBase := tokenv3.NewTokenManager(logger, tokenRepo)
		err = tokenManagerBase.SetUp(base)
		if err != nil {
			return err
		}

		pairRep := pairv3.NewRepo(client, confguration.Factory)
		pairManager := pairv3.NewPairManager(logger, pairRep)
		err = pairManager.SetUp(tokenManagerQuote.GetToken(), tokenManagerBase.GetToken(), confguration.Fee)
		if err != nil {
			return err
		}

		m.priceManagerV = pricev3.NewPriceManager(logger, pairManager)
	default:
		return fmt.Errorf("version %d is not supported", confguration.V)
	}

	return nil
}

func (m *priceManager) GetPrice() (float64, error) {
	return m.priceManagerV.GetPrice()
}
