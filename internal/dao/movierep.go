package dao

import (
	"cinema-search/internal/app"
	"context"
	"database/sql"
	"github.com/sirupsen/logrus"
)

type MovieRepository struct {
	movies *sql.DB
	//getMovieById(ctx context.Context, movieId string) (*domain.Movie, error)
}

func NewMovieRepository(db *sql.DB) *MovieRepository {
	return &MovieRepository{movies: db}
}

func (m *MovieRepository) GetMovieById(ctx context.Context, movieId string) (Movie, error) {
	stmt, err := m.movies.Prepare("SELECT * FROM movie WHERE id=$1")
	if err != nil {
		logrus.WithError(err).Println("Failed with \"Get movie by id\" query")
		return Movie{}, err
	}

	defer stmt.Close()

	result, err := stmt.QueryContext(ctx, movieId)
	if err != nil {
		logrus.Println(err)
		return Movie{}, err
	}
	movie := Movie{}
	if err := result.Scan(&movie.Id, &movie.Name, &movie.Budget, &movie.BoxOffice, &movie.Release, &movie.Duration,
		&movie.Synopsis, &movie.Rating, &movie.RatingNum, &movie.TrailerLink, &movie.Poster); err != nil {
		return movie, app.ErrMovieDoesntExist
	}
	return movie, nil
}
