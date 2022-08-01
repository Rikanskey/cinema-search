package movie

type Genre struct {
	id   uint64
	name string
	//movies []Movie
}

type CreateGenre struct {
	Id   uint64
	Name string
}

func UnmarshallGenreByDaoModel(genre CreateGenre) *Genre {
	return &Genre{
		id:   genre.Id,
		name: genre.Name,
	}
}
