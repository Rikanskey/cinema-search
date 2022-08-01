package dao

import (
	"cinema-search/internal/app"
	"cinema-search/internal/domain/movie"
)

func unmarshallMovie(m Movie) *movie.Movie {
	return movie.UnmarshallMovieByDaoModel(movie.CreateMovie{
		Id:          m.Id,
		Name:        m.Name,
		Budget:      m.Budget,
		BoxOffice:   m.BoxOffice,
		Release:     m.Release,
		Duration:    m.Duration,
		Synopsis:    m.Synopsis,
		Rating:      m.Rating,
		RatingNum:   m.RatingNum,
		TrailerLink: m.TrailerLink,
		PosterPath:  m.PosterPath,
		Genres:      unmarshallGenres(m.Genres),
		Countries:   unmarshallCountries(m.Countries),
		PersonPosts: unmarshallPersonPosts(m.PersonPosts),
	})
}

func unmarshallGenres(genres []Genre) []movie.Genre {
	movieGenres := make([]movie.Genre, 0, len(genres))

	for _, genre := range genres {
		movieGenres = append(movieGenres, *movie.UnmarshallGenreByDaoModel(movie.CreateGenre{
			Id:   genre.Id,
			Name: genre.Name,
		}))
	}

	return movieGenres
}

func unmarshallCountries(countries []Country) []movie.Country {
	movieCountries := make([]movie.Country, 0, len(countries))

	for _, country := range countries {
		movieCountries = append(movieCountries, *movie.UnmarshallCountryByDaoModel(movie.CreateCountry{
			Id:   country.Id,
			Name: country.Name,
		}))
	}

	return movieCountries
}

func unmarshallPersonPosts(personPosts []PersonPost) []movie.PersonPost {
	moviePersonPosts := make([]movie.PersonPost, 0, len(personPosts))

	for _, personPost := range personPosts {
		moviePersonPosts = append(moviePersonPosts, *movie.UnmarshallPersonPostByDaoModel(movie.CreatePersonPost{
			Id: personPost.Id,
			Person: movie.CreatePerson{
				Id:   personPost.Person.Id,
				Name: personPost.Person.Name,
			},
			Post: movie.CreatePost{
				Id:   personPost.Post.Id,
				Name: personPost.Post.Name,
			},
		}))
	}

	return moviePersonPosts
}

func unmarshallAppMovie(m Movie) app.Movie {
	return app.Movie{
		Id:          m.Id,
		Name:        m.Name,
		Budget:      m.Budget,
		BoxOffice:   m.BoxOffice,
		Release:     m.Release,
		Duration:    m.Duration,
		Synopsis:    m.Synopsis,
		Rating:      m.Rating,
		RatingNum:   m.RatingNum,
		TrailerLink: m.TrailerLink,
		PosterPath:  m.PosterPath,
		Genres:      unmarshallAppGenres(m.Genres),
		PersonsPost: unmarshallAppPersonPosts(m.PersonPosts),
		Countries:   unmarshallAppCountries(m.Countries),
	}
}

func unmarshallAppGenres(genres []Genre) []app.Genre {
	movieGenres := make([]app.Genre, 0, len(genres))

	for _, genre := range genres {
		movieGenres = append(movieGenres, unmarshallAppGenre(genre))
	}

	return movieGenres
}

func unmarshallAppGenre(genre Genre) app.Genre {
	return app.Genre{
		Id:   genre.Id,
		Name: genre.Name,
	}
}

func unmarshallAppPersonPosts(personPosts []PersonPost) []app.PersonPost {
	moviePersonPost := make([]app.PersonPost, 0, len(personPosts))

	for _, personPost := range personPosts {
		moviePersonPost = append(moviePersonPost, unmarshallAppPersonPost(personPost))
	}

	return moviePersonPost
}

func unmarshallAppPersonPost(personPost PersonPost) app.PersonPost {
	return app.PersonPost{
		Id:     personPost.Id,
		Post:   unmarshallAppPost(personPost.Post),
		Person: unmarshallAppPerson(personPost.Person),
	}
}

func unmarshallAppPost(post Post) app.Post {
	return app.Post{
		Id:   post.Id,
		Name: post.Name,
	}
}

func unmarshallAppPerson(person Person) app.Person {
	return app.Person{
		Id:   person.Id,
		Name: person.Name,
	}
}

func unmarshallAppCountries(countries []Country) []app.Country {
	movieCountries := make([]app.Country, 0, len(countries))

	for _, country := range countries {
		movieCountries = append(movieCountries, unmarshallAppCountry(country))
	}

	return movieCountries
}

func unmarshallAppCountry(country Country) app.Country {
	return app.Country{
		Id:   country.Id,
		Name: country.Name,
	}
}
