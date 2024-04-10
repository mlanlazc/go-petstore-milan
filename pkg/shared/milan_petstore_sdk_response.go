package shared

import (
	"github.com/mlanlazc/go-petstore-milan/internal/clients/rest/httptransport"
)

type MilanPetstoreSdkResponse[T any] struct {
	Data     T
	Metadata MilanPetstoreSdkResponseMetadata
}

type MilanPetstoreSdkResponseMetadata struct {
	Headers    map[string]string
	StatusCode int
}

func NewMilanPetstoreSdkResponse[T any](resp *httptransport.Response[T]) *MilanPetstoreSdkResponse[T] {
	return &MilanPetstoreSdkResponse[T]{
		Data: resp.Data,
		Metadata: MilanPetstoreSdkResponseMetadata{
			StatusCode: resp.StatusCode,
			Headers:    resp.Headers,
		},
	}
}
