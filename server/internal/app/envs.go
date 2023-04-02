package app

import (
	"fmt"
	"os"
)

// var DATABASE_URL string = "dbname=local_v8 user=postgres password=postgres host=127.0.0.1 port=5432"

type AppConfig struct {
	Port        string
	PG_Host     string
	PG_Port     string
	PG_DBName   string
	PG_User     string
	PG_Password string
}

func NewAppConfig() *AppConfig {

	port := os.Getenv("PORT")

	return &AppConfig{
		Port:        port,
		PG_Host:     os.Getenv("PG_HOST"),
		PG_Port:     os.Getenv("PG_PORT"),
		PG_DBName:   os.Getenv("PG_DBNAME"),
		PG_User:     os.Getenv("PG_USER"),
		PG_Password: os.Getenv("PG_PASSWORD"),
	}
}

func CreateDabaseURL(cfg *AppConfig) string {
	DATABASE_URL := fmt.Sprintf("dbname=%s user=%s password=%s host=%s port=%s",
		cfg.PG_DBName,
		cfg.PG_User,
		cfg.PG_Password,
		cfg.PG_Host,
		cfg.PG_Port,
	)
	return DATABASE_URL

}
