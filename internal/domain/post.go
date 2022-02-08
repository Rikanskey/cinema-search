package domain

type Post struct {
	id      uint64
	name    string
	persons []PersonPost
}
