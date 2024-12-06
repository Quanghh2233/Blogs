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
		User:     os.Getenv("DB_HOST"),
		Password: os.Getenv("DB_HOST"),
		Dbname:   os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_HOST"),
		Sslmode:  "disable",
		Timezone: "UTC",
	}
}
