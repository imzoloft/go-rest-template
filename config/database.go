package config

import "fmt"

type Database struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
}

func (d Database) FormatDSN() string {
	// change the dsn if you need
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s",
		d.Host, d.Port, d.User, d.Password, d.DbName,
	)
	if Env.Mode == "development" {
		return dsn + " sslmode=disable"
	}
	return dsn
}
