package app

import (
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
		Handle(ctx context.Context, q GetMovieQuery) (Movie, error)
	}
)
