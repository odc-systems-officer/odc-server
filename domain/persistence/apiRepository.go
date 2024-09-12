package persistence

import "odcserver/domain/models"

type ApiRepository interface {
	GetApiProfile(apiKey string) (*models.ApiProfile, error)
	// CreateApiKey(apiKey string, profile models.ApiKeyProfile) error
	// SaveApiKey(apiKey string, profile models.ApiKeyProfile) error
}