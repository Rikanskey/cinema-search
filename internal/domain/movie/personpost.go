package movie

import "cinema-search/internal/dao"

type PersonPost struct {
	id     uint64
	person *Person
	post   *Post
	//movies []Movie
}

func UnmarshallPersonPostByDaoModel(personPost dao.PersonPost) *PersonPost {
	return &PersonPost{
		id:     personPost.Id,
		person: UnmarshallPersonByDaoModel(personPost.Person),
		post:   UnmarshallPostByDaoModel(personPost.Post),
	}
}
