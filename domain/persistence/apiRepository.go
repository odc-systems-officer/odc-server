package persistence

import "odcserver/domain/models"

type ApiRepository interface {
	GetApiProfile(apiKey string) (*models.ApiProfile, error)
	SaveApiProfile(apiKey string, profile *models.ApiProfile) error
}