package v1

import "time"

//Movie defines model for Movie
type Movie struct {
	Id          uint64       `json:"id"`
	Name        string       `json:"name"`
	Budget      uint         `json:"budget,omitempty"`
	BoxOffice   uint         `json:"boxOffice,omitempty"`
	Release     time.Time    `json:"release,omitempty"`
	Duration    time.Time    `json:"duration,omitempty"`
	Synopsys    string       `json:"synopsys,omitempty"`
	Rating      float32      `json:"rating,omitempty"`
	RatingNum   uint         `json:"ratingNum,omitempty"`
	TrailerLink string       `json:"trailerLink,omitempty"`
	PosterPath  string       `json:"posterPath,omitempty"`
	Genres      []Genre      `json:"genres,omitempty"`
	Countries   []Country    `json:"countries,omitempty"`
	PersonsPost []PersonPost `json:"personPost,omitempty"`
}

//Genre defines model for Genre
type Genre struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
}

//Country defines model for Country
type Country struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
}

//Person defines model for Person
type Person struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
	//Firstname string `json:"firstname"`
	//Lastname  string `json:"lastname"`
}

//Post defines model for Post
type Post struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
}

//PersonPost defines model for PersonPost
type PersonPost struct {
	Id     uint64 `json:"id"`
	Person Person `json:"person"`
	Post   Post   `json:"post"`
}
