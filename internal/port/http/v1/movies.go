package v1

import (
	"net/http"
)

func (h handler) GetMovie(w http.ResponseWriter, r *http.Request, movieId string) {

	q, ok := unmarshalMovie(w, r, movieId)

	if !ok {
		return
	}

	movie, err := h.app.Queries.GetMovie.Handle(r.Context(), q)

	if err == nil {
		marshalMovie(w, r, movie)

		return
	}

}
