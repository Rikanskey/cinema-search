package dao

import "time"

type Movie struct {
	Id          uint64        `json:"id"`
	Name        string        `json:"name"`
	Budget      uint64        `json:"budget"`
	BoxOffice   uint64        `json:"box_office"`
	Release     time.Time     `json:"release"`
	Duration    time.Duration `json:"duration"`
	Synopsis    string        `json:"synopsis"`
	Rating      float32       `json:"rating"`
	RatingNum   float32       `json:"rating_num"`
	TrailerLink string        `json:"trailer_link"`
	Poster      string        `json:"poster"`
}

type user struct {
	Id       uint64 `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Avatar   string `json:"avatar"`
}

type genre struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
}

type country struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
}

type person struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
}
