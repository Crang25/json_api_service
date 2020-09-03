package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Crang25/json_api_service/cmd/storages/memstore"

	"github.com/joho/godotenv"

	"github.com/Crang25/json_api_service/cmd/router"
)

var (
	port string
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("failed to load enviromental variables: %v", err)
	}

	var isExists bool
	if port, isExists = os.LookupEnv("PORT"); !isExists {
		log.Fatalf("failed to load enviromental variable PORT")
	}
}

func main() {
	r := router.New(memstore.New())
	log.Fatalf(fmt.Sprintf("%v", http.ListenAndServe(port, r.NewHandler())))
}
