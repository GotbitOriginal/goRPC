package test

import (
	"os"
	"testing"

	"github.com/GotbitOriginal/goRPC/internal/logger"
	"github.com/GotbitOriginal/goRPC/internal/logic"
)

func TestGetPrice(t *testing.T) {
	logger := logger.New(os.Stdout)

	swap := "uni3/1"
	quote := "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" // WETH
	base := "0x21BfBDa47A0B4B5b1248c767Ee49F7caA9B23697"  // OVR

	priceManager := logic.NewPriceManager(logger)
	err := priceManager.SetUp(quote, base, swap)
	if err != nil {
		logger.Error(err)
		return
	}

	logger.Info(priceManager.GetPrice())
}
