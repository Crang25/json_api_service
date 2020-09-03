package router

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/maraino/testify/require"

	"github.com/Crang25/json_api_service/cmd/models"

	"github.com/Crang25/json_api_service/cmd/storages/memstore"
)

func TestGetAdList(t *testing.T) {
	r := New(memstore.New())
	srv := httptest.NewServer(r.NewHandler())
	defer srv.Close()

	// Create ad
	// id := checkCreatAd(t, srv.URL+"/ad", "BMW", "selling car",
	// 	[]string{"https://i.pinimg.com/originals/cd/3e/69/cd3e693e3ae9c74ef1480d4247840d32.jpg",
	// 		"https://i.pinimg.com/originals/99/9b/49/999b497965b5e2778bdf70427e27aaf8.jpg"}, 75630.92)

	// Checking getting ad by id

	// Checking getting ad list
}

func checkCreatAd(t *testing.T, url, name, description string, photoLinks []string, price float64) string {
	ad := models.Ad{
		Name:        name,
		Description: description,
		PhotoLinks:  photoLinks,
		Price:       price,
	}

	dataJSON, err := json.Marshal(&ad)
	require.NoError(t, err)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(dataJSON))
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)

	var id string
	err = json.NewDecoder(resp.Body).Decode(&id)
	require.NoError(t, err)

	return id
}

func checkGetAd(t *testing.T, url string, withFields bool) interface{} {
	resp, err := http.Get(url)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)

	if withFields {
		ad := struct {
			Name        string   `json:"name"`
			Description string   `json:"description"`
			Price       float64  `json:"price"`
			PhotoLink   []string `json:"photo_links"`
		}{}
		err = json.NewDecoder(resp.Body).Decode(&ad)
		require.NoError(t, err)
		return ad
	}

	ad := struct {
		Name      string  `json:"name"`
		Price     float64 `json:"Price"`
		PhotoLink string  `json:"photo_links"`
	}{}
	err = json.NewDecoder(resp.Body).Decode(&ad)
	require.NoError(t, err)
	return ad
}

func checkAdList(t *testing.T, url string, ads ...interface{}) {
	resp, err := http.Get(url)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)

	var respAds models.AdList
	err = json.NewDecoder(resp.Body).Decode(&ads)
	require.NoError(t, err)

	require.Equal(t, ads, respAds)
}
