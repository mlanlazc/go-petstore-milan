package configmanager

import "github.com/mlanlazc/go-petstore-milan/pkg/milanpetstoresdkconfig"

type ConfigManager struct {
	Pets milanpetstoresdkconfig.Config
}

func NewConfigManager(config milanpetstoresdkconfig.Config) *ConfigManager {
	return &ConfigManager{
		Pets: config,
	}
}

func (c *ConfigManager) SetBaseUrl(baseUrl string) {
	c.Pets.SetBaseUrl(baseUrl)
}

func (c *ConfigManager) GetPets() *milanpetstoresdkconfig.Config {
	return &c.Pets
}
