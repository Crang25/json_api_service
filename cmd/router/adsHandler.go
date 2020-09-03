package router

import (
	"log"
	"net/http"
	"sort"

	"github.com/Crang25/json_api_service/cmd/models"

	"github.com/Crang25/json_api_service/cmd/storages"
)

type adsHandler struct {
	store storages.Store
}

func newAdsHandler(store storages.Store) adsHandler {
	return adsHandler{
		store: store,
	}
}

func orderBy(ads models.AdList, isIncrease, isPrice bool) {

	if isPrice { // Order by price
		if isIncrease { // Order prices by increase
			sort.SliceStable(ads.Ads, func(i, j int) bool {
				return ads.Ads[i].Price < ads.Ads[j].Price
			})
		} else { // Order prices by decrease
			sort.SliceStable(ads.Ads, func(i, j int) bool {
				return ads.Ads[i].Price > ads.Ads[j].Price
			})
		}
	} else { // Order by date
		if isIncrease { // Order by date by increase
			sort.SliceStable(ads.Ads, func(i, j int) bool {
				return ads.Ads[i].Date.Before(ads.Ads[j].Date)
			})
		} else { // Order by prices by decrease
			sort.SliceStable(ads.Ads, func(i, j int) bool {
				return ads.Ads[i].Date.After(ads.Ads[j].Date)
			})
		}
	}
}

func gettingStrings(ads models.AdList, flag rune) []string {
	var strs []string
	switch flag {
	case 1:
		for _, ad := range ads.Ads {
			strs = append(strs, ad.Name)
		}
	case 2:
		for _, ad := range ads.Ads {
			strs = append(strs, ad.PhotoLinks[0])
		}
	}

	return strs
}

func (adsH adsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	head, _ := shiftPath(r.URL.Path)
	ads, err := adsH.store.GetAdList(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatalf("failed to get ad list: %v", err)
	}
	if head == "OrderByPriceIncrease" { // Sort by price by increase
		orderBy(ads, true, true)
	} else if head == "OrderByPriceDecrease" { // Sort by price decrease
		orderBy(ads, false, true)
	} else if head == "OrderByDateIncrease" { // Sort by date increase
		orderBy(ads, true, false)
	} else if head == "OrderByDateDecrease" { // Sort by date decrease
		orderBy(ads, false, false)
	}

	var prices []float64
	for _, ad := range ads.Ads {
		prices = append(prices, ad.Price)
	}

	getAds := struct {
		Name       []string  `json:"name"`
		PhotoLinks []string  `json:"photo_links"`
		Price      []float64 `json:"price"`
	}{
		Name:       gettingStrings(ads, 1),
		PhotoLinks: gettingStrings(ads, 2),
		Price:      prices,
	}

	writeJSON(w, getAds)
}
