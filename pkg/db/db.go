package DB

import (
	"config_service/internal/config"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose/v3"
	"time"
)

func MigrateDB(db *sqlx.DB) {
	goose.SetBaseFS(nil)
	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}

	if err := goose.Up(db.DB, "/scripts"); err != nil {
		panic(err)
	}
}

func NewDB(cfg *config.Config) (*sqlx.DB, error) {
	dbCfg := cfg.Connections.Postgres
	dbURL := fmt.Sprintf(
		"postgresql://%s:%s@%s:%d/%s?sslmode=%s",
		dbCfg.Username,
		dbCfg.Password,
		dbCfg.Host,
		dbCfg.Port,
		dbCfg.DBName,
		dbCfg.SSLEnabled,
	)

	db, err := sqlx.Connect("postgres", dbURL)
	if err != nil {
		return nil, err
	}

	MigrateDB(db)

	return db, nil
}

func NewRedis(cfg *config.Config) (*redis.Client, error) {
	var err error

	rdbCfg := cfg.Connections.Redis
	redisHost := rdbCfg.Host
	redisPort := rdbCfg.Port

	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", redisHost, redisPort),
		DB:   0,
	})
	for range 3 {
		_, err = rdb.Ping(context.Background()).Result()
		if err == nil {
			break
		}

		time.Sleep(1 * time.Second)
	}
	if err != nil {
		return nil, err
	}

	return rdb, nil
}
