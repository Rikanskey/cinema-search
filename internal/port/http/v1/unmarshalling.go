package v1

import (
	"cinema-search/internal/app"
	"net/http"
)

func unmarshalMovie(
	w http.ResponseWriter, r *http.Request,
	movieID string) (q app.GetMovieQuery, ok bool) {

	return app.GetMovieQuery{
		MovieId: movieID,
	}, true
}
