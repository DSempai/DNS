package storage

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
)

type Service struct {
	logger       *logrus.Logger
	pool         *pgxpool.Pool
	dsn          string
	queryBuilder squirrel.StatementBuilderType
}

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

func (s Service) Close() {
	if s.pool == nil {
		return
	}
	s.pool.Close()
}
