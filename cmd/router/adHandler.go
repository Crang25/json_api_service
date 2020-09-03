package router

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Crang25/json_api_service/cmd/models"

	"github.com/Crang25/json_api_service/cmd/storages"
)

type adHandler struct {
	store storages.Store
}

func newAdHandler(store storages.Store) adHandler {
	return adHandler{
		store: store,
	}
}

func (adH adHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet: // Getting specific ad
		var ads models.Ad

		head, tail := shiftPath(r.URL.Path)
		_, err := strconv.Atoi(head)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		ads, err = adH.store.GetAd(r.Context(), head)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Fatalf("failed to getting ad: %v", err)
		}
		if ads.Price == -1 {
			http.NotFound(w, r)
			return
		}

		if tail == "/fields" {
			getAd := struct {
				Name        string   `json:"name"`
				Description string   `json:"description"`
				Price       float64  `json:"price"`
				PhotoLink   []string `json:"photo_links"`
			}{
				Name:        ads.Name,
				Description: ads.Description,
				Price:       ads.Price,
				PhotoLink:   ads.PhotoLinks,
			}

			writeJSON(w, getAd)
		} else {
			getAd := struct {
				Name      string  `json:"name"`
				Price     float64 `json:"Price"`
				PhotoLink string  `json:"photo_links"`
			}{
				Name:      ads.Name,
				Price:     ads.Price,
				PhotoLink: ads.PhotoLinks[0],
			}

			getAdJSON, err := json.Marshal(&getAd)
			if err != nil {
				log.Fatalf("failed to marshal ad: %v", err)
			}
			w.Write(getAdJSON)
		}

	case http.MethodDelete: // Deleting specific ad
		id, _ := shiftPath(r.URL.Path)
		log.Println("(deleting)id is", id)
		// Code ...
	case http.MethodPost: // Creating new ad
		var ad models.Ad
		err := json.NewDecoder(r.Body).Decode(&ad)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Fatalf("failed to decode json: %v", err)
		}

		id, err := adH.store.CreateAd(r.Context(), ad.Name, ad.Description, ad.PhotoLinks, ad.Price)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Fatalf("failed to create ad: %v", err)
		}

		writeJSON(w, id)
	case http.MethodPut: // Updating ad
		log.Println("updating ad\nurl is", r.URL.Path)
		// code ...
	}
}
