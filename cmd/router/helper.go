package router

import (
	"encoding/json"
	"log"
	"net/http"
)

func writeJSON(w http.ResponseWriter, data interface{}) {
	dataJSON, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatalf("failed to marshal data: %v", err)
	}

	w.Write(dataJSON)
}
