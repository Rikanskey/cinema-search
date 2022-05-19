package v1

import (
	"cinema-search/internal/dao"
	"github.com/go-chi/render"
	"net/http"
)

//func marshalMovie(w http.ResponseWriter, r *http.Request, movie app.MovieResponse) {
//	response := marshalMovieToMovieResponse(movie)
//
//	render.Respond(w, r, response)
//}

//func marshalMovieToMovieResponse(movie app.MovieResponse) Movie {
//	return Movie{
//		Id: movie.Id,
//	}
//}

func marshalMovie(w http.ResponseWriter, r *http.Request, movie dao.Movie) {
	response := marshalMovieToMovieResponse(movie)

	render.Respond(w, r, response)
}

func marshalMovieToMovieResponse(movie dao.Movie) Movie {
	m := Movie{
		Id:       movie.Id,
		Name:     movie.Name,
		Release:  movie.Release,
		Duration: movie.Duration,
	}

	//m.Genres = make(map[uint64]Genre)
	for i := 0; i < len(movie.Genres); i++ {
		//m.Genres[movie.Genres[i].Id] = Genre{
		//	Name: movie.Genres[i].Name,
		//}
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

	for i := 0; i < len(movie.PersonPosts); i++ {
		m.PersonPost = append(m.PersonPost, PersonPost{
			Id: movie.PersonPosts[i].Id,
			Person: Person{
				Id:   movie.PersonPosts[i].Person.Id,
				Name: movie.PersonPosts[i].Person.Name,
			},
			Post: Post{
				Id:   movie.PersonPosts[i].Post.Id,
				Name: movie.PersonPosts[i].Post.Name,
			},
		})
	}

	return m
}
