package config

import "os"

type Config struct {
	AppName      string `env:"APP_NAME"`
	AppVersion   string `env:"APP_VERSION"`
	BuildType    string `env:"BUILD_TYPE"`
	ElectrumHost string `env:"ELECTRUM_HOST"`
	ElectrumPort string `env:"ELECTRUM_PORT"`
}

func GetConfig() *Config {
	return &Config{
		AppName:      os.Getenv("APP_NAME"),
		AppVersion:   os.Getenv("APP_VERSION"),
		BuildType:    os.Getenv("BUILD_TYPE"),
		ElectrumHost: os.Getenv("ELECTRUM_HOST"),
		ElectrumPort: os.Getenv("ELECTRUM_PORT"),
	}
}
