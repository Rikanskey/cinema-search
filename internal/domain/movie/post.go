package movie

import "cinema-search/internal/dao"

type Post struct {
	id      uint64
	name    string
	persons []PersonPost
}

func UnmarshallPostByDaoModel(post dao.Post) *Post {
	return &Post{
		id:   post.Id,
		name: post.Name,
	}
}
