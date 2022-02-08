package app

import (
	"cinema-search/internal/dao"
	"context"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type (
	Commands struct {
	}
)

type (
	Queries struct {
		GetMovie getMovie
	}

	getMovie interface {
		Handle(ctx context.Context, movieId string) (dao.Movie, error)
	}
)
