package dao

import "time"

type (
	Movie struct {
		Id          uint64    `db:"id"`
		Name        string    `db:"name"`
		Budget      uint      `db:"budget"`
		BoxOffice   uint      `db:"box_office"`
		Release     time.Time `db:"release"`
		Duration    time.Time `db:"duration"`
		Synopsis    string    `db:"synopsis"`
		Rating      float32   `db:"rating"`
		RatingNum   uint      `db:"rating_num"`
		TrailerLink string    `db:"trailer_link"`
		PosterPath  string    `db:"poster"`
		Genres      []Genre
		Countries   []Country
		PersonPosts []PersonPost
	}

	User struct {
		Id        uint64 `db:"id"`
		Username  string `db:"username"`
		Password  string `db:"password"`
		Avatar    string `db:"avatar"`
		Authority []Authority
	}

	Genre struct {
		Id   uint64 `db:"id"`
		Name string `db:"name"`
	}

	Country struct {
		Id   uint64 `db:"id"`
		Name string `db:"name"`
	}

	Person struct {
		Id   uint64 `db:"id"`
		Name string `db:"name"`
	}

	Post struct {
		Id   uint64 `db:"id"`
		Name string `db:"name"`
	}

	PersonPost struct {
		Id     uint64 `db:"id"`
		Person Person `db:"person"`
		Post   Post   `db:"post"`
	}

	Authority struct {
		Id   uint64 `db:"id"`
		Name string `db:"name"`
	}
)
