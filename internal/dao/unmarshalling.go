package dao

import "cinema-search/internal/domain/movie"

func unmarshallMovie(m Movie) *movie.Movie {
	return movie.UnmarshallMovieByDaoModel(m)
}
