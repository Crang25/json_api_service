package router

import (
	"net/http"

	"github.com/Crang25/json_api_service/cmd/storages"
)

// Router ...
type Router struct {
	rootHandler rootHandler
}

// New ...
func New(store storages.Store) *Router {
	return &Router{
		rootHandler: newRootHandler(store),
	}
}

// NewHandler ...
func (r *Router) NewHandler() http.Handler {
	return &r.rootHandler
}
