package app

import (
	"fmt"
	"os"
	"strconv"
)

// var DATABASE_URL string = "dbname=local_v8 user=postgres password=postgres host=127.0.0.1 port=5432"

type AppConfig struct {
	Port        string
	PG_Host     string
	PG_Port     string
	PG_DBName   string
	PG_User     string
	PG_Password string
	HashCost    int
}

func NewAppConfig() (*AppConfig, error) {

	port := os.Getenv("PORT")
	hash_cost, err := strconv.Atoi(os.Getenv("HASH_COST"))
	if err != nil {
		return &AppConfig{}, fmt.Errorf("error during creating AppConfig, could not read HASH_COST, either empty or non int; err: %w", err)
	}

	return &AppConfig{
		Port:        port,
		PG_Host:     os.Getenv("PG_HOST"),
		PG_Port:     os.Getenv("PG_PORT"),
		PG_DBName:   os.Getenv("PG_DBNAME"),
		PG_User:     os.Getenv("PG_USER"),
		PG_Password: os.Getenv("PG_PASSWORD"),
		HashCost:    hash_cost,
	}, nil
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
