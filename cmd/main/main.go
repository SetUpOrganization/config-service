package main

import (
	"config_service/internal/config"
	"config_service/internal/repo"
	"config_service/internal/service"
	DB "config_service/pkg/db"
	"log/slog"
	"os"
)

func main() {
	var err error

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	// TODO: normal logger

	cfg, err := config.LoadConfig("config.ini")
	if err != nil {
		logger.Error("Fatal load config: ", err.Error())
		os.Exit(1)
	}

	db, err := DB.NewDB(cfg)
	if err != nil {
		logger.Error("Fatal connect to database: ", err.Error())
		os.Exit(1)
	}
	defer db.Close()

	rdb, err := DB.NewRedis(cfg)
	if err != nil {
		logger.Error("Fatal connect to redis: ", err.Error())
	}
	defer rdb.Close()

	repository := repo.NewRepo(db, rdb)
	services, err := service.NewService(repository, cfg.Connections.Services.Users, cfg.Connections.Services.TimeOut)
	if err != nil {
		logger.Error("Fatal create service: ", err.Error())
	}

	//TODO: NewServer() + Start
}
