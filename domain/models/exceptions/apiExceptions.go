package exceptions

import "errors"

const (
	ApiKeyNotFound = "API Key not found"
	ApiKeyRequired = "API Key required"
)

var (
	ErrApiKeyNotFound = errors.New(ApiKeyNotFound)
	ErrApiKeyRequired = errors.New(ApiKeyRequired)
)