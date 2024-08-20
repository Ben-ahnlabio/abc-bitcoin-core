package container

import (
	"github.com/ahnlabio/bitcoin-core/electrum-api/config"
	"github.com/ahnlabio/bitcoin-core/electrum-api/electrum"
	"github.com/ahnlabio/bitcoin-core/electrum-api/service"
)

var container *Container

type Container struct {
	AppConfig      *config.Config
	Electrum       *electrum.Electrum
	BitcoinService *service.BitcoinApiService
}

func GetInstnace() *Container {
	if container == nil {
		appConfig := config.GetConfig()
		electrum := electrum.NewElectrum(appConfig.ElectrumHost, appConfig.ElectrumPort)
		bitcoinService := service.NewBitcoinApiService(electrum)
		container = &Container{
			AppConfig:      appConfig,
			BitcoinService: bitcoinService,
		}
	}
	return container
}

func (c Container) GetBitcoinApiService() *service.BitcoinApiService {
	return c.BitcoinService
}
