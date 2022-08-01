package query

import (
	"cinema-search/internal/app"
	"context"
	"github.com/pkg/errors"
)

type getMovieModel interface {
	GetMovieById(ctx context.Context, movieId string) (app.Movie, error)
}

type GetMovieHandler struct {
	getModel getMovieModel
}

func NewGetMovieHandler(getModel getMovieModel) GetMovieHandler {
	if getModel == nil {
		panic("get model is nil")
	}

	return GetMovieHandler{getModel: getModel}
}

func (h GetMovieHandler) Handle(ctx context.Context, query app.GetMovieQuery) (app.Movie, error) {
	movie, err := h.getModel.GetMovieById(ctx, query.MovieId)

	return movie, errors.Wrapf(err, "Getting movie %d with id %s", movie.Id, movie.Name)
}
