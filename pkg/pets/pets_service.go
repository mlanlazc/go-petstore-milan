package pets

import (
	"context"
	restClient "github.com/mlanlazc/go-petstore-milan/internal/clients/rest"
	"github.com/mlanlazc/go-petstore-milan/internal/clients/rest/httptransport"
	"github.com/mlanlazc/go-petstore-milan/internal/configmanager"
	"github.com/mlanlazc/go-petstore-milan/pkg/milanpetstoresdkconfig"
	"github.com/mlanlazc/go-petstore-milan/pkg/shared"
)

type PetsService struct {
	manager *configmanager.ConfigManager
}

func NewPetsService(manager *configmanager.ConfigManager) *PetsService {
	return &PetsService{
		manager: manager,
	}
}

func (api *PetsService) getConfig() *milanpetstoresdkconfig.Config {
	return api.manager.GetPets()
}

func (api *PetsService) SetBaseUrl(baseUrl string) {
	config := api.getConfig()
	config.SetBaseUrl(baseUrl)
}

func (api *PetsService) ListPets(ctx context.Context, params ListPetsRequestParams) (*shared.MilanPetstoreSdkResponse[[]Pet], *shared.MilanPetstoreSdkError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[[]Pet](config)

	request := httptransport.NewRequest(ctx, "GET", "/pets", config)

	request.Options = params

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewMilanPetstoreSdkError[[]Pet](err)
	}

	return shared.NewMilanPetstoreSdkResponse[[]Pet](resp), nil
}

func (api *PetsService) CreatePets(ctx context.Context, pet Pet) (*shared.MilanPetstoreSdkResponse[any], *shared.MilanPetstoreSdkError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[any](config)

	request := httptransport.NewRequest(ctx, "POST", "/pets", config)

	request.Body = pet

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewMilanPetstoreSdkError[any](err)
	}

	return shared.NewMilanPetstoreSdkResponse[any](resp), nil
}

func (api *PetsService) ShowPetById(ctx context.Context, petId string) (*shared.MilanPetstoreSdkResponse[Pet], *shared.MilanPetstoreSdkError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[Pet](config)

	request := httptransport.NewRequest(ctx, "GET", "/pets/{petId}", config)

	request.SetPathParam("petId", petId)

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewMilanPetstoreSdkError[Pet](err)
	}

	return shared.NewMilanPetstoreSdkResponse[Pet](resp), nil
}
