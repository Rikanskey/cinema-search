package movie

type PersonPost struct {
	id     uint64
	person *Person
	post   *Post
	//movies []Movie
}

type CreatePersonPost struct {
	Id     uint64
	Person CreatePerson
	Post   CreatePost
}

func UnmarshallPersonPostByDaoModel(personPost CreatePersonPost) *PersonPost {
	return &PersonPost{
		id:     personPost.Id,
		person: UnmarshallPersonByDaoModel(personPost.Person),
		post:   UnmarshallPostByDaoModel(personPost.Post),
	}
}
