package app

import (
	"fmt"
	"os"
	"strconv"
)

// var DATABASE_URL string = "dbname=local_v8 user=postgres password=postgres host=127.0.0.1 port=5432"

type AppConfig struct {
	Port               string
	PG_Host            string
	PG_Port            string
	PG_DBName          string
	PG_User            string
	PG_Password        string
	HashCost           int
	PublicKeyPath      string
	PrivateKeyPath     string
}

func NewAppConfig() (*AppConfig, error) {

	port := os.Getenv("PORT")
	hash_cost, err := strconv.Atoi(os.Getenv("HASH_COST"))
	if err != nil {
		return &AppConfig{}, fmt.Errorf("error during creating AppConfig, could not read HASH_COST, either empty or non int; err: %w", err)
	}

	// public/private keys
	private_key_path := os.Getenv("PRIVATE_KEY_PATH")
	public_key_path := os.Getenv("PUBLIC_KEY_PATH")
	return &AppConfig{
		PublicKeyPath:      public_key_path,
		PrivateKeyPath:     private_key_path,
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
