package apiserver

import "os"

//Config  ...
type Config struct {
	BindAddr    string `toml:"bind_addr"`
	DatabaseURL string `toml:"database_url"`
}

var (
	bindaddr = os.Getenv("BINDADDRESS")
	dbURL    = os.Getenv("DATABASEURL")
)

//NewConfig ...
func NewConfig() *Config {
	return &Config{
		BindAddr:    bindaddr, //":8080",
		DatabaseURL: dbURL,    //"host=localhost user=postgres password=1234 dbname=restapi_dev sslmode=disable",
	}
}
