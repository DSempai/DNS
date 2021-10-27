package storage

import (
	"context"
	"errors"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"
)

func (s Service) RetrieveBySectorID(id int64) (Sector, error) {
	var result Sector

	sectors := s.queryBuilder.
		Select("id", "sector_id", "created_at", "active").
		From("sectors").
		Where(squirrel.Eq{"sector_id": id})
	sql, args, err := sectors.ToSql()
	if err != nil {
		s.logger.Errorf("Create sql query failed. Error: %v", err)
		return result, err
	}

	row := s.pool.QueryRow(context.Background(), sql, args...)

	err = row.Scan(&result.ID, &result.SectorID, &result.CreatedAt, &result.Active)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			s.logger.Infof("Information for sector %d not found", id)
			return result, ErrSectorNotFound
		}
		s.logger.Errorf("Scan row failed. Error: %v", err)
		return result, err
	}

	return result, nil
}
