package container

import (
	"log"

	"github.com/ahnlabio/bitcoin-core/bitcoin-api/config"
	"github.com/ahnlabio/bitcoin-core/bitcoin-api/electrum"
	"github.com/ahnlabio/bitcoin-core/bitcoin-api/handlers"
	"github.com/ahnlabio/bitcoin-core/bitcoin-api/service"
)

var container *Container

type Container struct {
	AppConfig      *config.Config
	Electrum       *electrum.Electrum
	BitcoinService *service.BtcService
	Handler        *handlers.Handler
}

func GetInstnace() *Container {
	if container == nil {
		log.Print("Container is not initialized. Create new container.")

		appConfig := config.GetConfig()
		electrum := electrum.NewElectrum(appConfig.ElectrumHost, appConfig.ElectrumPort)
		bitcoinService := service.NewBitcoinApiService(electrum)
		handlers := handlers.NewHandler(bitcoinService)

		container = &Container{
			AppConfig:      appConfig,
			BitcoinService: bitcoinService,
			Handler:        handlers,
		}
	}
	return container
}

func (c *Container) GetHandler() *handlers.Handler {
	return c.Handler
}
