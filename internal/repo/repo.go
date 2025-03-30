package repo

import (
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
	"sync"
)

type Repo struct {
	db  *sqlx.DB
	rdb *redis.Client
	mu  *sync.Mutex
}

func NewRepo(db *sqlx.DB, rdb *redis.Client) *Repo {
	return &Repo{
		db:  db,
		rdb: rdb,
		mu:  &sync.Mutex{},
	}
}
