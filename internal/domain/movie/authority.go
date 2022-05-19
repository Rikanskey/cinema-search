package movie

import "cinema-search/internal/dao"

type Authority struct {
	id   uint64
	name string
}

func UnmarshallAuthorityByDaoModel(authority dao.Authority) *Authority {
	return &Authority{
		id:   authority.Id,
		name: authority.Name,
	}
}
