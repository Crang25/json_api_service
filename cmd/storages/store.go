package storages

import (
	"context"

	"github.com/Crang25/json_api_service/internal/models"
)

// Store ...
type Store interface {
	GetAdList(ctx context.Context) (models.AdList, error)
	GetAd(ctx context.Context, id string) (models.Ad, error)
	CreateAd(ctx context.Context, name, description string, photoLinks []string,
		price float64) (id string, err error)
	// UpdateAd(ctx context.Context, id string, newAd models.Ad) error
	// DeleteAd(ctx context.Context, id string) error
}
