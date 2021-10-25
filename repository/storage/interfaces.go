package storage

type SectorsInterface interface {
	RetrieveBySectorID(id int64) (Sector, error)
}

