package storage

import "time"

// Sector database table implementation.
type Sector struct {
	ID        int64
	SectorID  int64
	CreatedAt time.Time
	Active    bool
}
