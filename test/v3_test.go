package test

import (
	"fmt"
	"os"
	"testing"

	"github.com/GotbitOriginal/goRPC/internal/entity"
	"github.com/GotbitOriginal/goRPC/internal/logger"
	"github.com/GotbitOriginal/goRPC/internal/logic/v3/pair"
	"github.com/GotbitOriginal/goRPC/internal/logic/v3/price"
	"github.com/GotbitOriginal/goRPC/internal/logic/v3/token"

	"github.com/ethereum/go-ethereum/ethclient"
)

func TestGetPriceV3(t *testing.T) {
	var err error
	logger := logger.New(os.Stdout)

	swap := "uni"
	config, err := entity.GetConfig()
	if err != nil {
		logger.Error(err)
	}
	confguration, ok := config[swap]
	if !ok {
		logger.Error(fmt.Errorf("swap %s is not supported", swap))
		return
	}

	client, err := ethclient.Dial(confguration.RPC)
	if err != nil {
		logger.Error(err)
	}

	quote := "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" // WETH
	base := "0x21BfBDa47A0B4B5b1248c767Ee49F7caA9B23697"  // OVR

	tokenRepo := token.NewRepo(client)

	tokenManagerQuote := token.NewTokenManager(logger, tokenRepo)
	err = tokenManagerQuote.SetUp(quote)
	if err != nil {
		logger.Error(err)
	}

	tokenManagerBase := token.NewTokenManager(logger, tokenRepo)
	err = tokenManagerBase.SetUp(base)
	if err != nil {
		logger.Error(err)
	}

	pairRep := pair.NewRepo(client, "0x1F98431c8aD98523631AE4a59f267346ea31F984")
	pairManager := pair.NewPairManager(logger, pairRep)
	err = pairManager.SetUp(tokenManagerQuote.GetToken(), tokenManagerBase.GetToken(), "1")
	if err != nil {
		logger.Error(err)
	}

	priceManager := price.NewPriceManager(logger, pairManager)

	logger.Info(priceManager.GetPrice())
}
