package movie

import "cinema-search/internal/dao"

type Person struct {
	id   uint64
	name string
	//personPosts []PersonPosts
}

func UnmarshallPersonByDaoModel(person dao.Person) *Person {
	return &Person{
		id:   person.Id,
		name: person.Name,
	}
}
