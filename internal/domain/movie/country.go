package movie

type Country struct {
	id     uint64
	name   string
	movies []Movie
}

type CreateCountry struct {
	Id   uint64
	Name string
}

func UnmarshallCountryByDaoModel(country CreateCountry) *Country {
	return &Country{
		id:   country.Id,
		name: country.Name,
	}
}
