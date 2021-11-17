package entity

type PersonPost struct {
	id uint64
	person *Person
	post *Post
	movies []Movie
}
