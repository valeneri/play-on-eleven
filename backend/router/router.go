package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Router struct {
	*mux.Router
}

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// NewRouter creates a new router instance
func NewRouter() *Router {
	// new router
	r := mux.NewRouter().PathPrefix("/api/v1").Subrouter()
	r.StrictSlash(false)
	router := Router{r}

	return &router
}

func (r *Router) AddRoute(routes []Route, prefix string) *Router {
	for _, route := range routes {
		// logger.WithField("route", route).Debug("adding route to mux")
		r.
			Methods(route.Method).
			Path(prefix + route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	return r
}
