package storage

import "time"

type Sector struct {
	ID        int64
	SectorID  int64
	CreatedAt time.Time
	Active    bool
}
