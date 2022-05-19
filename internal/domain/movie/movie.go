package movie

import (
	"cinema-search/internal/dao"
	"time"
)

type Movie struct {
	id          uint64
	name        string
	budget      uint64
	boxOffice   uint64
	release     time.Time
	duration    time.Time
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
	Duration    time.Time
	Synopsys    string
	Rating      float32
	RatingNum   uint
	TrailerLink string
	PosterPath  string
}

func UnmarshallMovieByDaoModel(movie dao.Movie) *Movie {
	return &Movie{
		id:          movie.Id,
		name:        movie.Name,
		budget:      movie.Budget,
		boxOffice:   movie.BoxOffice,
		release:     movie.Release,
		duration:    movie.Duration,
		synopsys:    movie.Synopsis,
		rating:      movie.Rating,
		ratingNum:   movie.RatingNum,
		trailerLink: movie.TrailerLink,
		posterPath:  movie.PosterPath,
	}
}

func InitMovie(movie CreateMovie) *Movie {
	return &Movie{id: movie.Id}
}
