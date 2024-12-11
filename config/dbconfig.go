package config

import "os"

type dbConf struct {
	Host     string
	User     string
	Password string
	Dbname   string
	Port     string
	Sslmode  string
	Timezone string
}

func GetDbConf() *dbConf {
	return &dbConf{
		Host:     os.Getenv("DB_HOST"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Dbname:   os.Getenv("DB_NAME"),
		Port:     os.Getenv("DB_PORT"),
		Sslmode:  "disable",
		Timezone: "UTC",
	}
}
