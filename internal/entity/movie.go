package entity

type Movie struct {
	id uint64
	name, synopsys, posterPath, trailerLink string
	rating float32
	budget, boxOffice, ratingNum uint
}
