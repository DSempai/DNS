package storage

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
)

// Service provides all the necessary tools to work with the database entity
// and implement interface of the database tables : storage.SectorsInterface, ...
type Service struct {
	logger       *logrus.Logger
	pool         *pgxpool.Pool
	dsn          string
	queryBuilder squirrel.StatementBuilderType
}

// Initialize connect to database with provided configuration, invoke query builder
// and return Service entity. Provide instruments to work with database.
func Initialize(logger *logrus.Logger, dsn string, maxConn int32) (Service, error) {
	conf, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return Service{}, err
	}
	conf.MaxConns = maxConn
	pg, err := pgxpool.ConnectConfig(context.Background(), conf)
	if err != nil {
		return Service{}, fmt.Errorf("create db connection failed. Error: %w", err)
	}
	statementBuilder := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	return Service{
		logger:       logger,
		pool:         pg,
		dsn:          dsn,
		queryBuilder: statementBuilder,
	}, nil
}

// Close connection pool.
func (s Service) Close() {
	if s.pool == nil {
		return
	}
	s.pool.Close()
}
