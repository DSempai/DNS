package storage

// SectorsInterface describes all the methods that we use when working with "sectors" table.
type SectorsInterface interface {
	RetrieveBySectorID(id int64) (Sector, error)
}
