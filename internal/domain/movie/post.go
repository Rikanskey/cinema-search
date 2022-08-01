package movie

type Post struct {
	id      uint64
	name    string
	persons []PersonPost
}

type CreatePost struct {
	Id   uint64
	Name string
}

func UnmarshallPostByDaoModel(post CreatePost) *Post {
	return &Post{
		id:   post.Id,
		name: post.Name,
	}
}
