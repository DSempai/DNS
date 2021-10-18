package storage

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
)

type Storage struct {
	logger       *logrus.Logger
	pool         *pgxpool.Pool
	dsn          string
	queryBuilder squirrel.StatementBuilderType
}

func Initialize(logger *logrus.Logger, dsn string, maxConn int32) (Storage, error) {
	conf, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return Storage{}, err
	}
	conf.MaxConns = maxConn
	pg, err := pgxpool.ConnectConfig(context.Background(), conf)

	if err != nil {
		err = fmt.Errorf("error during creating storagedb connection: %w", err)
		return Storage{}, err
	}

	statementBuilder := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	return Storage{
		logger:       logger,
		pool:         pg,
		dsn:          dsn,
		queryBuilder: statementBuilder,
	}, nil
}

func (s Storage) Close() {
	if s.pool == nil {
		return
	}
	s.pool.Close()
}
