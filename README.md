# MilanPetstoreSdk Go SDK 3.5.2

A Go SDK for MilanPetstoreSdk.

- API version: 1.0.0
- SDK version: 3.5.2

## Table of Contents

- [Services](#services)

## Services

### PetsService

#### ListPets

List all pets

```GO
import (
  "fmt"
  "encoding/json"
  "github.com/mlanlazc/go-petstore-milan/pkg/milanpetstoresdkconfig"
  "github.com/mlanlazc/go-petstore-milan/pkg/milanpetstoresdk"
  "github.com/mlanlazc/go-petstore-milan/pkg/pets"
)

config := milanpetstoresdkconfig.NewConfig()
client := milanpetstoresdk.NewMilanPetstoreSdk(config)


params := pets.ListPetsRequestParams{}


response, err := client.Pets.ListPets(context.Background(), params)
if err != nil {
  panic(err)
}

fmt.Print(response)
```

#### CreatePets

Create a pet

```GO
import (
  "fmt"
  "encoding/json"
  "github.com/mlanlazc/go-petstore-milan/pkg/milanpetstoresdkconfig"
  "github.com/mlanlazc/go-petstore-milan/pkg/milanpetstoresdk"
  "github.com/mlanlazc/go-petstore-milan/pkg/pets"
)

config := milanpetstoresdkconfig.NewConfig()
client := milanpetstoresdk.NewMilanPetstoreSdk(config)


request := pets.Pet{}
request.SetId(int64(123))
request.SetName("Name")

response, err := client.Pets.CreatePets(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Print(response)
```

#### ShowPetById

Info for a specific pet

```GO
import (
  "fmt"
  "encoding/json"
  "github.com/mlanlazc/go-petstore-milan/pkg/milanpetstoresdkconfig"
  "github.com/mlanlazc/go-petstore-milan/pkg/milanpetstoresdk"
)

config := milanpetstoresdkconfig.NewConfig()
client := milanpetstoresdk.NewMilanPetstoreSdk(config)

response, err := client.Pets.ShowPetById(context.Background(), "petId")
if err != nil {
  panic(err)
}

fmt.Print(response)
```
