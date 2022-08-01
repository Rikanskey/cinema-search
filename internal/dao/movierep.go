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

func (m *MovieRepository) GetMovieById(ctx context.Context, movieId string) (app.Movie, error) {
	stmt, err := m.movies.Prepare("SELECT * FROM movie WHERE id=$1")
	defer stmt.Close()
	if err != nil {
		logrus.WithError(err).Println("Failed with \"Get movie by id\" query")
		return app.Movie{}, err
	}

	result, err := stmt.QueryContext(ctx, movieId)
	if err != nil {
		logrus.WithError(err).Println("Failed with \"Get movie by id\" query")
		return app.Movie{}, err
	}
	movie := Movie{}
	result.Next()
	if err := result.Scan(&movie.Id, &movie.Name, &movie.Budget, &movie.BoxOffice, &movie.Release, &movie.Duration,
		&movie.Synopsis, &movie.Rating, &movie.RatingNum, &movie.TrailerLink, &movie.PosterPath); err != nil {
		logrus.WithError(err).Println("Failed with \"Get movie by id\" query")
		return unmarshallAppMovie(movie), err
	}

	genres, err := m.getMovieGenres(ctx, movieId)
	if err != nil {
		logrus.WithError(err).Println("Failed with \"Get movie by id\" query")
		return unmarshallAppMovie(movie), err
	}
	movie.Genres = genres

	countries, err := m.getMovieCountries(ctx, movieId)
	if err != nil {
		logrus.WithError(err).Println("Failed with \"Get movie by id\" query")
		return unmarshallAppMovie(movie), err
	}
	movie.Countries = countries

	personPosts, err := m.getMoviePersonPosts(ctx, movieId)
	if err != nil {
		logrus.WithError(err).Println("Failed with \"Get movie by id\" query")
		return unmarshallAppMovie(movie), err
	}
	movie.PersonPosts = personPosts

	return unmarshallAppMovie(movie), nil
}

func (m *MovieRepository) getMovieGenres(ctx context.Context, movieId string) ([]Genre, error) {
	var genres []Genre

	stmt, err := m.movies.Prepare("SELECT genre.id, name FROM genre JOIN movie_genre ON movie_genre.genre_id=genre.id WHERE movie_genre.movie_id=$1")
	defer stmt.Close()

	if err != nil {
		return genres, err
	}

	result, err := stmt.QueryContext(ctx, movieId)
	if err != nil {
		return genres, err
	}

	for result.Next() {
		genre := Genre{}
		if err := result.Scan(&genre.Id, &genre.Name); err != nil {
			return genres, err
		}
		genres = append(genres, genre)
	}

	return genres, nil
}

func (m *MovieRepository) getMovieCountries(ctx context.Context, movieId string) ([]Country, error) {
	var countries []Country

	stmt, err := m.movies.Prepare("SELECT country.id, name FROM country JOIN movie_country ON movie_country.country_id=country.id WHERE movie_country.movie_id=$1")
	defer stmt.Close()

	if err != nil {
		return countries, err
	}

	result, err := stmt.QueryContext(ctx, movieId)
	if err != nil {
		return countries, err
	}

	for result.Next() {
		country := Country{}
		if err := result.Scan(&country.Id, &country.Name); err != nil {
			return countries, err
		}
		countries = append(countries, country)
	}

	return countries, nil
}

func (m *MovieRepository) getMoviePersonPosts(ctx context.Context, movieId string) ([]PersonPost, error) {
	var personPosts []PersonPost

	stmt, err := m.movies.Prepare("SELECT person_post_movie.id, p.id, p.name, p2.id, p2.name FROM person_post_movie JOIN person_post pp on person_post_movie.person_post_id = pp.id JOIN person p on pp.person_id = p.id JOIN post p2 on pp.post_id = p2.id WHERE person_post_movie.movie_id = $1")
	defer stmt.Close()

	if err != nil {
		return personPosts, err
	}

	result, err := stmt.QueryContext(ctx, movieId)
	if err != nil {
		return personPosts, err
	}

	for result.Next() {
		person := Person{}
		post := Post{}
		personPost := PersonPost{
			Person: person,
			Post:   post,
		}
		if err := result.Scan(&personPost.Id, &personPost.Person.Id, &personPost.Person.Name,
			&personPost.Post.Id, &personPost.Post.Name); err != nil {
			return personPosts, err
		}
		personPosts = append(personPosts, personPost)
	}

	return personPosts, nil
}
