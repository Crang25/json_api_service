package router

import (
	"net/http"

	"github.com/Crang25/json_api_service/cmd/storages"
)

type rootHandler struct {
	adsHandler adsHandler
	adHandler  adHandler
}

func newRootHandler(store storages.Store) rootHandler {
	return rootHandler{
		adsHandler: newAdsHandler(store),
		adHandler:  newAdHandler(store),
	}
}

func (rH rootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var head string
	head, r.URL.Path = shiftPath(r.URL.Path)

	switch head {
	case "ads":
		rH.adsHandler.ServeHTTP(w, r)
	case "ad":
		rH.adHandler.ServeHTTP(w, r)
	default:
		http.NotFound(w, r)
	}
}
