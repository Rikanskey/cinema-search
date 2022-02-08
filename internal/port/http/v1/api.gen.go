package v1

import (
	"fmt"
	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/go-chi/chi"
	"net/http"
)

type ServerInterface interface {

	// (GET /movie/{movieId})
	GetMovie(w http.ResponseWriter, r *http.Request, movieId string)
}

func (siw *ServerInterfaceWrapper) GetMovie(w http.ResponseWriter, r *http.Request) {
	var err error
	var movieId string

	ctx := r.Context()

	err = runtime.BindStyledParameter("simple", false, "movieId",
		chi.URLParam(r, "movieId"), &movieId)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid parameter for movieId: %s", err), http.StatusBadRequest)
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetMovie(w, r, movieId)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{})
}

type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
}

type ChiServerOptions struct {
	BaseUrl     string
	Router      chi.Router
	Middlewares []MiddlewareFunc
}

type MiddlewareFunc func(http.HandlerFunc) http.HandlerFunc

func HandlerWithOptions(serverInterface ServerInterface, options ChiServerOptions) http.Handler {
	router := options.Router

	if router == nil {
		router = chi.NewRouter()
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            serverInterface,
		HandlerMiddlewares: options.Middlewares,
	}

	router.Group(func(r chi.Router) {
		r.Get(options.BaseUrl+"/movie", wrapper.GetMovie)
	})

	return router
}

func HandlerFromMux(si ServerInterface, router chi.Router) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		Router: router,
	})
}
