package db

import "github.com/jackc/pgx/v5/pgxpool"

// 處理 Data func 的 interface
// For Extend 用
// 防止 再次用sqlc時會洗走新加的部分
type Store interface {
	Querier
}

type SQLStore struct {
	connPool *pgxpool.Pool
	*Queries
}

func NewStore(connPool *pgxpool.Pool) Store {
	return &SQLStore{
		connPool: connPool,
		Queries:  New(connPool),
	}
}
