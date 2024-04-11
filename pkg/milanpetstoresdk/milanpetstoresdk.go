package milanpetstoresdk

import (
	"github.com/mlanlazc/go-petstore-milan/internal/configmanager"
	"github.com/mlanlazc/go-petstore-milan/pkg/milanpetstoresdkconfig"
	"github.com/mlanlazc/go-petstore-milan/pkg/pets"
)

type MilanPetstoreSdk struct {
	Pets    *pets.PetsService
	manager *configmanager.ConfigManager
}

func NewMilanPetstoreSdk(config milanpetstoresdkconfig.Config) *MilanPetstoreSdk {
	manager := configmanager.NewConfigManager(config)
	return &MilanPetstoreSdk{
		Pets:    pets.NewPetsService(manager),
		manager: manager,
	}
}

func (m *MilanPetstoreSdk) SetBaseUrl(baseUrl string) {
	m.manager.SetBaseUrl(baseUrl)
}

// c029837e0e474b76bc487506e8799df5e3335891efe4fb02bda7a1441840310c
