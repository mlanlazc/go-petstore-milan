package shared

import (
	"github.com/mlanlazc/go-petstore-milan/internal/clients/rest/httptransport"
)

type MilanPetstoreSdkError struct {
	Err      error
	Body     []byte
	Metadata MilanPetstoreSdkErrorMetadata
}

type MilanPetstoreSdkErrorMetadata struct {
	Headers    map[string]string
	StatusCode int
}

func NewMilanPetstoreSdkError[T any](transportError *httptransport.ErrorResponse[T]) *MilanPetstoreSdkError {
	return &MilanPetstoreSdkError{
		Err:  transportError.GetError(),
		Body: transportError.GetBody(),
		Metadata: MilanPetstoreSdkErrorMetadata{
			StatusCode: transportError.GetStatusCode(),
			Headers:    transportError.GetHeaders(),
		},
	}
}

func (e *MilanPetstoreSdkError) Error() string {
	return e.Err.Error()
}
