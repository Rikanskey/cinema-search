package movie

type Person struct {
	id   uint64
	name string
	//personPosts []PersonPosts
}

type CreatePerson struct {
	Id   uint64
	Name string
}

func UnmarshallPersonByDaoModel(person CreatePerson) *Person {
	return &Person{
		id:   person.Id,
		name: person.Name,
	}
}
