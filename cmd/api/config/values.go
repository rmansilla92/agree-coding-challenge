package config

import (
	"fmt"
	"os"
)

const ()

var (
	DBName           string
	DBUser           string
	DBPass           string
	DBHost           string
	DBPort           string
	ConnectionString string
)

func LoadProductionValues() {
	DBName = os.Getenv("PRODUCTION_DB_NAME")
	DBUser = os.Getenv("PRODUCTION_DB_USER")
	DBPass = os.Getenv("PRODUCTION_DB_PASS")
	DBHost = os.Getenv("PRODUCTION_DB_HOST")
	DBPort = os.Getenv("PRODUCTION_DB_PORT")
	ConnectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", DBUser, DBPass, DBHost, DBPort, DBName)
}

func LoadDevelopValues() {
	DBName = os.Getenv("LOCAL_DB_NAME")
	DBUser = os.Getenv("LOCAL_DB_USER")
	DBPass = os.Getenv("LOCAL_DB_PASS")
	DBHost = os.Getenv("LOCAL_DB_HOST")
	DBPort = os.Getenv("LOCAL_DB_PORT")
	ConnectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", DBUser, DBPass, DBHost, DBPort, DBName)
}
