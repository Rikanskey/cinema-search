package v1

import "net/http"

func (h handler) GetMovie(w http.ResponseWriter, r *http.Request, movieId string) {
	movie, err := h.app.Queries.GetMovie.Handle(r.Context(), movieId)

	if err == nil {
		return
	}
}
