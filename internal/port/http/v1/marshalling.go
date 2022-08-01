package v1

import (
	"cinema-search/internal/app"
	"github.com/go-chi/render"
	"net/http"
)

func marshalMovie(w http.ResponseWriter, r *http.Request, movie app.Movie) {
	response := marshalMovieToMovieResponse(movie)

	render.Respond(w, r, response)
}

func marshalMovieToMovieResponse(movie app.Movie) Movie {
	m := Movie{
		Id:       movie.Id,
		Name:     movie.Name,
		Release:  movie.Release,
		Duration: movie.Duration,
	}

	for i := 0; i < len(movie.Genres); i++ {
		m.Genres = append(m.Genres, Genre{
			Id:   movie.Genres[i].Id,
			Name: movie.Genres[i].Name,
		})
	}

	for i := 0; i < len(movie.Countries); i++ {
		m.Countries = append(m.Countries, Country{
			Id:   movie.Countries[i].Id,
			Name: movie.Countries[i].Name,
		})
	}

	for i := 0; i < len(movie.PersonsPost); i++ {
		m.PersonsPost = append(m.PersonsPost, PersonPost{
			Id: movie.PersonsPost[i].Id,
			Person: Person{
				Id:   movie.PersonsPost[i].Person.Id,
				Name: movie.PersonsPost[i].Person.Name,
			},
			Post: Post{
				Id:   movie.PersonsPost[i].Post.Id,
				Name: movie.PersonsPost[i].Post.Name,
			},
		})
	}

	return m
}
