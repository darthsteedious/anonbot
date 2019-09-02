package routing

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Router interface {
	RegisterRoute(method, path string, handler http.HandlerFunc) Router
	AsHandler() http.Handler
}

type router struct {
	inner *mux.Router
}

func (r *router) RegisterRoute(method, path string, handler http.HandlerFunc) Router {
	r.inner.
		Methods(method).
		Path(path).
		HandlerFunc(handler)

	return r
}

func (r *router) AsHandler() http.Handler {
	return r.inner
}

func NewRouter(basePath string) Router {
	return &router {
		inner: mux.NewRouter().PathPrefix(basePath).Subrouter(),
	}
}