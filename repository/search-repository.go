package repository

type SearchDB struct {
	DB PgxIface
}

func NewSearchRepository(ds PgxIface) SearchDB {
	return SearchDB{DB: ds}
}

func (s *SearchDB) Check(places string) (bool, error) {
	return false, nil
}

func (s *SearchDB) Create(places string) error {
	return nil
}

func (s *SearchDB) CheckGeo(places string) (bool, error) {
	return false, nil
}

func (s *SearchDB) CreateGeo(places string) error {
	return nil
}
