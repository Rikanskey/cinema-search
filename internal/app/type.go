package app

import "time"

type (
	Movie struct {
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
		PersonsPost []PersonPost
		Countries   []Country
	}
	Country struct {
		Id   uint64
		Name string
	}
	Genre struct {
		Id   uint64
		Name string
	}
	Person struct {
		Id   uint64
		Name string
	}
	Post struct {
		Id   uint64
		Name string
	}
	PersonPost struct {
		Id     uint64
		Person Person
		Post   Post
	}
)
