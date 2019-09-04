package dbs

import (
	"database/sql"
	"fmt"
)

type DBSource struct {
	Driver     string `json:"driver"`
	ServerName string `json:"server_name"`
	Name       string `json:"name"`
	User       string `json:"user"`
	Password   string `json:"password"`
}

func (dbSource *DBSource) Connection() (*sql.DB, error) {
	db, err := sql.Open(dbSource.Driver, dbSource.Source())
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (dbSource *DBSource) Source() string {
	if dbSource.Driver == "mysql" {
		return fmt.Sprintf("%s:%s@tcp(%s)/%s", dbSource.User, dbSource.Password, dbSource.ServerName, dbSource.Name)
	}

	if dbSource.Driver == "sqlite3" {
		return dbSource.Name
	}

	return ""
}
