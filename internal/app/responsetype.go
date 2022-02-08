package app

type (
	MovieResponse struct {
		Id   uint64
		Name string

		PersonsPost []PersonPostMovieResponse
		Countries   []CountryMovieResponse
	}
	CountryMovieResponse struct {
		Id   uint64
		Name string
	}
	PersonMovieResponse struct {
		Id   uint64
		Name string
	}
	PostMovieResponse struct {
		Id   uint64
		Name string
	}
	PersonPostMovieResponse struct {
		Id     uint64
		Person PersonMovieResponse
		Post   PostMovieResponse
	}
)
