package mock

import "DNS/repository/storage"

type Sectors struct{
	Values storage.Sector
	Err error
}
func(m Sectors) RetrieveBySectorID(id int64) (storage.Sector, error){
	return m.Values, m.Err
}