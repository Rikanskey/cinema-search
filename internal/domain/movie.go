package domain

import "time"

type Movie struct {
	id          uint64
	name        string
	budget      uint
	boxOffice   uint
	release     time.Time
	duration    time.Duration
	synopsys    string
	rating      float32
	ratingNum   uint
	trailerLink string
	posterPath  string
	genres      []Genre
	countries   []Country
	personPost  []PersonPost
}

type CreateMovie struct {
	Id          uint64
	Name        string
	Budget      uint
	BoxOffice   uint
	Release     time.Time
	Duration    time.Duration
	Synopsys    string
	Rating      float32
	RatingNum   uint
	TrailerLink string
	PosterPath  string
}

func InitMovie(movie CreateMovie) *Movie {
	return &Movie{id: movie.Id}
}
