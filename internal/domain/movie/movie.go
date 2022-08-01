package movie

import (
	"time"
)

type Movie struct {
	id          uint64
	name        string
	budget      uint
	boxOffice   uint
	release     time.Time
	duration    time.Time
	synopsis    string
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
	Duration    time.Time
	Synopsis    string
	Rating      float32
	RatingNum   uint
	TrailerLink string
	PosterPath  string
	Genres      []Genre
	Countries   []Country
	PersonPosts []PersonPost
}

func UnmarshallMovieByDaoModel(movie CreateMovie) *Movie {
	return &Movie{
		id:          movie.Id,
		name:        movie.Name,
		budget:      movie.Budget,
		boxOffice:   movie.BoxOffice,
		release:     movie.Release,
		duration:    movie.Duration,
		synopsis:    movie.Synopsis,
		rating:      movie.Rating,
		ratingNum:   movie.RatingNum,
		trailerLink: movie.TrailerLink,
		posterPath:  movie.PosterPath,
	}
}

func InitMovie(movie CreateMovie) *Movie {
	return &Movie{id: movie.Id}
}
