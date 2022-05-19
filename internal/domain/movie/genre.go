package movie

import "cinema-search/internal/dao"

type Genre struct {
	id   uint64
	name string
	//movies []Movie
}

func UnmarshallGenreByDaoModel(genre dao.Genre) *Genre {
	return &Genre{
		id:   genre.Id,
		name: genre.Name,
	}
}
