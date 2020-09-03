package memstore

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/Crang25/json_api_service/cmd/models"
)

// MemStore ...
type MemStore struct {
	mu     sync.Mutex
	ads    models.AdList
	lastID int64
}

// New ...
func New() *MemStore {
	return &MemStore{}
}

// GetAdList ...
func (ms *MemStore) GetAdList(ctx context.Context) (models.AdList, error) {

	ms.mu.Lock()
	defer ms.mu.Unlock()

	ads := models.AdList{
		Ads: append([]models.Ad(nil), ms.ads.Ads...),
	}

	return ads, nil

}

// GetAd ...
func (ms *MemStore) GetAd(ctx context.Context, id string) (models.Ad, error) {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	SearchAd := models.Ad{Price: -1}
	for _, ad := range ms.ads.Ads {
		if ad.ID == id {
			SearchAd = ad
			break
		}
	}
	return SearchAd, nil
}

// CreateAd ...
func (ms *MemStore) CreateAd(ctx context.Context, name, description string, photoLinks []string,
	price float64) (id string, err error) {

	ms.mu.Lock()
	defer ms.mu.Unlock()

	ms.lastID++
	var newAd models.Ad
	newAd.Date = time.Now()
	newAd.ID = fmt.Sprintf("%v", ms.lastID)
	newAd.Price = price

	// Validation
	if len(photoLinks) > 3 {
		newAd.PhotoLinks = append(newAd.PhotoLinks, photoLinks[0], photoLinks[1], photoLinks[2])
	} else {
		newAd.PhotoLinks = append(newAd.PhotoLinks, photoLinks...)
	}
	if len([]byte(name)) <= 200 {
		newAd.Name = name
	}
	if len([]byte(description)) <= 1000 {
		newAd.Description = description
	}

	ms.ads.Ads = append(ms.ads.Ads, newAd)

	return newAd.ID, nil
}
